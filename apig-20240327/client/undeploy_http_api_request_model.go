// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUndeployHttpApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *UndeployHttpApiRequest
	GetEnvironmentId() *string
	SetGatewayId(v string) *UndeployHttpApiRequest
	GetGatewayId() *string
	SetOperationId(v string) *UndeployHttpApiRequest
	GetOperationId() *string
	SetRouteId(v string) *UndeployHttpApiRequest
	GetRouteId() *string
}

type UndeployHttpApiRequest struct {
	// The environment ID.
	//
	// example:
	//
	// env-cqsmtellhtgvo***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Gateway ID.
	//
	// example:
	//
	// gw-cq7l5s5lhtg***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The operation ID.
	//
	// example:
	//
	// op-d4favr6m1hk***
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
	// The route ID. You must specify this parameter when you unpublish the route of an HTTP API.
	//
	// example:
	//
	// hr-cr82undlhtgrle***
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
}

func (s UndeployHttpApiRequest) String() string {
	return dara.Prettify(s)
}

func (s UndeployHttpApiRequest) GoString() string {
	return s.String()
}

func (s *UndeployHttpApiRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UndeployHttpApiRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *UndeployHttpApiRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *UndeployHttpApiRequest) GetRouteId() *string {
	return s.RouteId
}

func (s *UndeployHttpApiRequest) SetEnvironmentId(v string) *UndeployHttpApiRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UndeployHttpApiRequest) SetGatewayId(v string) *UndeployHttpApiRequest {
	s.GatewayId = &v
	return s
}

func (s *UndeployHttpApiRequest) SetOperationId(v string) *UndeployHttpApiRequest {
	s.OperationId = &v
	return s
}

func (s *UndeployHttpApiRequest) SetRouteId(v string) *UndeployHttpApiRequest {
	s.RouteId = &v
	return s
}

func (s *UndeployHttpApiRequest) Validate() error {
	return dara.Validate(s)
}
