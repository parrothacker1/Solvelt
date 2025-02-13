package routers

import (
	"net/http"
	"strings"
	"time"

	"github.com/parrothacker1/Solvelt/users/utils/loggers"
)

type Route struct {
  Method      string
  Path        string
  Middleware  []http.HandlerFunc
  Handler     http.HandlerFunc
}

type Router struct {
  routes []Route
}

func NewRouter() *Router {
  return &Router{}
}

func (r *Router) Handle(method, path string,handler http.HandlerFunc,middleware ...http.HandlerFunc) {
  r.routes = append(r.routes, Route{
    Method: method,
    Path: path,
    Middleware: middleware,
    Handler: handler,
  })
}

func (rr *Router) ServeHTTP(w http.ResponseWriter,r *http.Request) {
  start := time.Now()
  for _,route := range rr.routes {
    if strings.EqualFold(r.Method,route.Method) && r.URL.Path == route.Path {
      w.Header().Set("Content-Type", "application/json")
      if len(route.Middleware) != 0 {
        for _,middleware := range route.Middleware {
          middleware.ServeHTTP(w,r)
          stop,ok := r.Context().Value("stop").(bool)
          if stop == true && ok {
            break
          }
        }
        if r.Context().Value("stop") == true {
          return
        } else {
          route.Handler.ServeHTTP(w,r)
        }
      } else {
        route.Handler.ServeHTTP(w,r)
      }
      loggers.ServerLogger.Infof("execution time=%dμs, method=%s, path=%s, ip=%s",
        time.Since(start).Microseconds(),
        r.Method,
        r.URL.Path,
        r.URL.Host,
      )
      return
    }
  }
  http.NotFound(w,r)
  loggers.ServerLogger.Infof("execution time=%dμs, method=%s, path=%s, ip=%s",
  time.Since(start).Microseconds(),
  r.Method,
  r.URL.Path,
  r.URL.Host)
}
