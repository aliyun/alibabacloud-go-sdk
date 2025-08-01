// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayRouteDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetGatewayRouteDetailRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *GetGatewayRouteDetailRequest
	GetGatewayUniqueId() *string
	SetRouteId(v int64) *GetGatewayRouteDetailRequest
	GetRouteId() *int64
}

type GetGatewayRouteDetailRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-5d3a78a53ec947aa928212d671d400ac
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// 1050
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
}

func (s GetGatewayRouteDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetGatewayRouteDetailRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayRouteDetailRequest) GetRouteId() *int64 {
	return s.RouteId
}

func (s *GetGatewayRouteDetailRequest) SetAcceptLanguage(v string) *GetGatewayRouteDetailRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetGatewayRouteDetailRequest) SetGatewayUniqueId(v string) *GetGatewayRouteDetailRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayRouteDetailRequest) SetRouteId(v int64) *GetGatewayRouteDetailRequest {
	s.RouteId = &v
	return s
}

func (s *GetGatewayRouteDetailRequest) Validate() error {
	return dara.Validate(s)
}
