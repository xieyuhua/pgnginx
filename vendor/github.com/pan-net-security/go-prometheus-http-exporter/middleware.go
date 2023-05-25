// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package prom_http_exporter

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type exporter struct {
	inFlightGauge prometheus.Gauge
	counter       *prometheus.CounterVec
	duration      *prometheus.HistogramVec
	responseSize  *prometheus.HistogramVec
}

func New() *exporter {
	e := &exporter{
		inFlightGauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "http",
			Name:      "in_flight_requests",
			Help:      "A gauge of requests currently being served by the wrapped handler.",
		}),

		counter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "http",
				Name:      "api_requests_total",
				Help:      "A counter for requests to the wrapped handler.",
			},
			[]string{"code", "method", "handler"},
		),

		duration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "http",
				Name:      "request_duration_seconds",
				Help:      "A histogram of latencies for requests.",
				Buckets:   []float64{.25, .5, 1, 2.5, 5, 10},
			},
			[]string{"handler", "method"},
		),

		responseSize: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "http",
				Name:      "response_size_bytes",
				Help:      "A histogram of response sizes for requests.",
				Buckets:   []float64{200, 500, 900, 1500},
			},
			[]string{},
		),
	}
	prometheus.MustRegister(e.inFlightGauge, e.counter, e.duration, e.responseSize)
	return e
}

func (e *exporter) Metric(path string, h http.HandlerFunc) (string, http.Handler) {
	return path, promhttp.InstrumentHandlerInFlight(e.inFlightGauge,
		promhttp.InstrumentHandlerDuration(e.duration.MustCurryWith(prometheus.Labels{"handler": path}),
			promhttp.InstrumentHandlerCounter(e.counter.MustCurryWith(prometheus.Labels{"handler": path}),
				promhttp.InstrumentHandlerResponseSize(e.responseSize, h),
			),
		),
	)
}
