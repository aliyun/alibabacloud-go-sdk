// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteWafStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteWafStatusRequest
	GetAcceptLanguage() *string
	SetEnableWaf(v bool) *UpdateGatewayRouteWafStatusRequest
	GetEnableWaf() *bool
	SetGatewayUniqueId(v string) *UpdateGatewayRouteWafStatusRequest
	GetGatewayUniqueId() *string
	SetRouteId(v int64) *UpdateGatewayRouteWafStatusRequest
	GetRouteId() *int64
}

type UpdateGatewayRouteWafStatusRequest struct {
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
	// Specifies whether to activate Web Application Firewall (WAF).
	//
	// example:
	//
	// true
	EnableWaf *bool `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-86575c0bc9f04ecfbacb92b8e392****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// 645
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
}

func (s UpdateGatewayRouteWafStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteWafStatusRequest) GetEnableWaf() *bool {
	return s.EnableWaf
}

func (s *UpdateGatewayRouteWafStatusRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteWafStatusRequest) GetRouteId() *int64 {
	return s.RouteId
}

func (s *UpdateGatewayRouteWafStatusRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteWafStatusRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusRequest) SetEnableWaf(v bool) *UpdateGatewayRouteWafStatusRequest {
	s.EnableWaf = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteWafStatusRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusRequest) SetRouteId(v int64) *UpdateGatewayRouteWafStatusRequest {
	s.RouteId = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusRequest) Validate() error {
	return dara.Validate(s)
}
