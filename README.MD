## Krakend Test

```
  docker run -d -p 8080:8080 --add-host=host.docker.internal:host-gateway krakend-test run --config "/etc/krakend/krakend.json"
  go run main.go
```

## Access Category List
  
  ```
    curl -X GET "http://localhost:8080/category/list" -H "accept: application/json"
  ```