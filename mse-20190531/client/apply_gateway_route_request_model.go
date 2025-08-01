// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyGatewayRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ApplyGatewayRouteRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *ApplyGatewayRouteRequest
	GetGatewayUniqueId() *string
	SetRouteId(v string) *ApplyGatewayRouteRequest
	GetRouteId() *string
}

type ApplyGatewayRouteRequest struct {
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
	// gw-1a4ab101d5924b6f92c5ec98a841761f
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// 950
	RouteId *string `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
}

func (s ApplyGatewayRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyGatewayRouteRequest) GoString() string {
	return s.String()
}

func (s *ApplyGatewayRouteRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ApplyGatewayRouteRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ApplyGatewayRouteRequest) GetRouteId() *string {
	return s.RouteId
}

func (s *ApplyGatewayRouteRequest) SetAcceptLanguage(v string) *ApplyGatewayRouteRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ApplyGatewayRouteRequest) SetGatewayUniqueId(v string) *ApplyGatewayRouteRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ApplyGatewayRouteRequest) SetRouteId(v string) *ApplyGatewayRouteRequest {
	s.RouteId = &v
	return s
}

func (s *ApplyGatewayRouteRequest) Validate() error {
	return dara.Validate(s)
}
