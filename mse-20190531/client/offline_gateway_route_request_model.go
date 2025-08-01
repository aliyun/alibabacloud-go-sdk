// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineGatewayRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *OfflineGatewayRouteRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *OfflineGatewayRouteRequest
	GetGatewayUniqueId() *string
	SetRouteId(v string) *OfflineGatewayRouteRequest
	GetRouteId() *string
}

type OfflineGatewayRouteRequest struct {
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
	// gw-77e1153db6e14c0a8b1fae20bcb89ca5
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// 645
	RouteId *string `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
}

func (s OfflineGatewayRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineGatewayRouteRequest) GoString() string {
	return s.String()
}

func (s *OfflineGatewayRouteRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *OfflineGatewayRouteRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *OfflineGatewayRouteRequest) GetRouteId() *string {
	return s.RouteId
}

func (s *OfflineGatewayRouteRequest) SetAcceptLanguage(v string) *OfflineGatewayRouteRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *OfflineGatewayRouteRequest) SetGatewayUniqueId(v string) *OfflineGatewayRouteRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *OfflineGatewayRouteRequest) SetRouteId(v string) *OfflineGatewayRouteRequest {
	s.RouteId = &v
	return s
}

func (s *OfflineGatewayRouteRequest) Validate() error {
	return dara.Validate(s)
}
