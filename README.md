# TODO
simple todo api

## Using gin and gorm (sqlite)
```
protected.POST("/todos", handler.NewTask)
protected.GET("/todos", handler.List)
protected.DELETE("/todos/:id", handler.Remove)
protected.PUT("/todos/:id", handler.Update)
```

### How to send request
You must request for jwt token first ( /tokenz ) then copy the token's string and paste it after "bearer "
before sending protected requests.

<br/>
Basic liveness, readiness probes, and rate limiter are also provided.
