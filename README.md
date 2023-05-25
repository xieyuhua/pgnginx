pgnginx  

-----------
> ./pgnginx  -fcgi /tmp/php-cgi-71.sock -http 0.0.0.0:8485 -root /www/wwwroot/demo/public/ 


```
./pgnginx  -h
Usage of ./pgnginx:
  -ext comma separated list
    	the fastcgi file extension(s) comma separated list (default "php")
  -fcgi string
    	the fcgi backend to connect to, you can pass more fcgi related params as query params (default "unix:///var/run/php/php7.0-fpm.sock")
  -http string
    	the http address to listen on (default ":6065")
  -index comma separated list
    	the default index file comma separated list (default "index.php,index.html")
  -listing
    	whether to allow directory listing or not
  -root string
    	the document root (default "./")
  -router string
    	the router filename incase of any 404 error (default "index.php")
  -rtimeout int
    	the read timeout, zero means unlimited
  -wtimeout int
    	the write timeout, zero means unlimited
    	
    
	
    	
func Serve(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	w.Write([]byte("bar"))
}

/prometheus metrics
func main() {
	e := prom_http_exporter.New()

	r := http.NewServeMux()
	r.Handle(e.Metric("/", Serve))
	r.Handle("/metrics", promhttp.Handler())

	s := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
	}
	fmt.Printf("  http server started on %s\n", *FlagHTTPAddr)
	log.Fatal(http.ListenAndServe(*FlagHTTPAddr, s.Handler))
	
// 	log.Fatal(http.ListenAndServe(*FlagHTTPAddr, http.HandlerFunc(Serve)))
}
```