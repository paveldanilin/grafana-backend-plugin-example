## Test resource plugin for Grafana 6.7.x

Just an example of usage [grafana-plugin-sdk](https://github.com/grafana/grafana-plugin-sdk-go) to build a resource plugin. 

### Build

`go build -o test-plugin_windows_amd64.exe ./pkg`

### Install
Copy `/dist` folder to a Grafana plugin directory and restart Grafana service.

### Test

`http://<your-grafana-host>/api/plugins/grafana-backend-plugin/resources`

The result 
```json
{
    "code": 3345,
    "message": "You are awesome"
}
```
