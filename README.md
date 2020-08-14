# cf-telemetry-to

A Spring Boot based app:
- consume cf-telemetry files
- parse/transform to TO metrics
- send to Tanzu Observability

### Plan
- Consume cf-telemetry JSON files
- Parse desired metrics
- Deliver metrics

### Spring Boot Integration

https://docs.wavefront.com/wavefront_springboot.html

### Uploading files

https://spring.io/guides/gs/uploading-files/

### Send metrics directly to TO:

```
curl 'https://<wavefront_instance>/api/v2/alert' --header 'Authorization: Bearer <wavefront_api_token>'

curl -X POST --header "Content-Type: application/octet-stream"
--header "Accept: application/json"
--header "Authorization: Bearer xxxxxxx"
-d "test.metric 100 source=test.source"
"https://longboard.wavefront.com/report?f=wavefront"
```