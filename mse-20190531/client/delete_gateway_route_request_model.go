// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteGatewayRouteRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *DeleteGatewayRouteRequest
	GetGatewayUniqueId() *string
	SetRouteId(v string) *DeleteGatewayRouteRequest
	GetRouteId() *string
}

type DeleteGatewayRouteRequest struct {
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
	// gw-492af9b04bb4474cae9d645be850e3d7
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// 252
	RouteId *string `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
}

func (s DeleteGatewayRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRouteRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteGatewayRouteRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayRouteRequest) GetRouteId() *string {
	return s.RouteId
}

func (s *DeleteGatewayRouteRequest) SetAcceptLanguage(v string) *DeleteGatewayRouteRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteGatewayRouteRequest) SetGatewayUniqueId(v string) *DeleteGatewayRouteRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayRouteRequest) SetRouteId(v string) *DeleteGatewayRouteRequest {
	s.RouteId = &v
	return s
}

func (s *DeleteGatewayRouteRequest) Validate() error {
	return dara.Validate(s)
}
