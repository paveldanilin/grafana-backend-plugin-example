package main

import (
	"encoding/json"
	"github.com/grafana/grafana-plugin-sdk-go/backend/plugin"
	"github.com/grafana/grafana-plugin-sdk-go/genproto/pluginv2"
	plugin2 "github.com/hashicorp/go-plugin"
)

type resourceServer struct {
}

func (r *resourceServer) CallResource(req *pluginv2.CallResourceRequest, srv pluginv2.Resource_CallResourceServer) error {
	data := map[string]interface{}{
		"message": "You are awesome",
		"code": 3345,
	}
	b, _ := json.Marshal(data)
	srv.Send(&pluginv2.CallResourceResponse{
		Code: 201,
		Body: b,
	})
	return nil
}

func main() {
	rr := &resourceServer{}
	plugin.Serve(plugin.ServeOpts{
		ResourceServer:    rr,
		GRPCServer:        plugin2.DefaultGRPCServer,
	})
}
