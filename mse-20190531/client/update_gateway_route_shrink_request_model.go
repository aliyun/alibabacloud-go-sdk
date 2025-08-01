// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteShrinkRequest
	GetAcceptLanguage() *string
	SetDescription(v string) *UpdateGatewayRouteShrinkRequest
	GetDescription() *string
	SetDestinationType(v string) *UpdateGatewayRouteShrinkRequest
	GetDestinationType() *string
	SetDirectResponseJSONShrink(v string) *UpdateGatewayRouteShrinkRequest
	GetDirectResponseJSONShrink() *string
	SetDomainIdListJSON(v string) *UpdateGatewayRouteShrinkRequest
	GetDomainIdListJSON() *string
	SetEnableWaf(v bool) *UpdateGatewayRouteShrinkRequest
	GetEnableWaf() *bool
	SetFallback(v bool) *UpdateGatewayRouteShrinkRequest
	GetFallback() *bool
	SetFallbackServicesShrink(v string) *UpdateGatewayRouteShrinkRequest
	GetFallbackServicesShrink() *string
	SetGatewayId(v int64) *UpdateGatewayRouteShrinkRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayRouteShrinkRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayRouteShrinkRequest
	GetId() *int64
	SetName(v string) *UpdateGatewayRouteShrinkRequest
	GetName() *string
	SetPredicatesShrink(v string) *UpdateGatewayRouteShrinkRequest
	GetPredicatesShrink() *string
	SetRedirectJSONShrink(v string) *UpdateGatewayRouteShrinkRequest
	GetRedirectJSONShrink() *string
	SetRouteOrder(v int32) *UpdateGatewayRouteShrinkRequest
	GetRouteOrder() *int32
	SetServicesShrink(v string) *UpdateGatewayRouteShrinkRequest
	GetServicesShrink() *string
}

type UpdateGatewayRouteShrinkRequest struct {
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
	// example:
	//
	// a route for xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination service type.
	//
	// example:
	//
	// Mock
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The information about service mocking.
	DirectResponseJSONShrink *string `json:"DirectResponseJSON,omitempty" xml:"DirectResponseJSON,omitempty"`
	// The associated domain name.
	//
	// example:
	//
	// [90]
	DomainIdListJSON *string `json:"DomainIdListJSON,omitempty" xml:"DomainIdListJSON,omitempty"`
	// Deprecated
	//
	// Specifies whether to activate Web Application Firewall (WAF).
	//
	// example:
	//
	// false
	EnableWaf *bool `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	// Specifies whether to enable the Fallback service.
	//
	// example:
	//
	// true
	Fallback *bool `json:"Fallback,omitempty" xml:"Fallback,omitempty"`
	// The information about the Fallback service.
	FallbackServicesShrink *string `json:"FallbackServices,omitempty" xml:"FallbackServices,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 501
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-86575c0bc9f04ecfbacb92b8e392a2c4
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// 139
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Deprecated
	//
	// The name of the route.
	//
	// example:
	//
	// route-web
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The route matching conditions.
	PredicatesShrink *string `json:"Predicates,omitempty" xml:"Predicates,omitempty"`
	// The information about redirection.
	RedirectJSONShrink *string `json:"RedirectJSON,omitempty" xml:"RedirectJSON,omitempty"`
	// The sequence number of the route.
	//
	// example:
	//
	// 1
	RouteOrder *int32 `json:"RouteOrder,omitempty" xml:"RouteOrder,omitempty"`
	// The information about destination services.
	ServicesShrink *string `json:"Services,omitempty" xml:"Services,omitempty"`
}

func (s UpdateGatewayRouteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGatewayRouteShrinkRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *UpdateGatewayRouteShrinkRequest) GetDirectResponseJSONShrink() *string {
	return s.DirectResponseJSONShrink
}

func (s *UpdateGatewayRouteShrinkRequest) GetDomainIdListJSON() *string {
	return s.DomainIdListJSON
}

func (s *UpdateGatewayRouteShrinkRequest) GetEnableWaf() *bool {
	return s.EnableWaf
}

func (s *UpdateGatewayRouteShrinkRequest) GetFallback() *bool {
	return s.Fallback
}

func (s *UpdateGatewayRouteShrinkRequest) GetFallbackServicesShrink() *string {
	return s.FallbackServicesShrink
}

func (s *UpdateGatewayRouteShrinkRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayRouteShrinkRequest) GetPredicatesShrink() *string {
	return s.PredicatesShrink
}

func (s *UpdateGatewayRouteShrinkRequest) GetRedirectJSONShrink() *string {
	return s.RedirectJSONShrink
}

func (s *UpdateGatewayRouteShrinkRequest) GetRouteOrder() *int32 {
	return s.RouteOrder
}

func (s *UpdateGatewayRouteShrinkRequest) GetServicesShrink() *string {
	return s.ServicesShrink
}

func (s *UpdateGatewayRouteShrinkRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetDescription(v string) *UpdateGatewayRouteShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetDestinationType(v string) *UpdateGatewayRouteShrinkRequest {
	s.DestinationType = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetDirectResponseJSONShrink(v string) *UpdateGatewayRouteShrinkRequest {
	s.DirectResponseJSONShrink = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetDomainIdListJSON(v string) *UpdateGatewayRouteShrinkRequest {
	s.DomainIdListJSON = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetEnableWaf(v bool) *UpdateGatewayRouteShrinkRequest {
	s.EnableWaf = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetFallback(v bool) *UpdateGatewayRouteShrinkRequest {
	s.Fallback = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetFallbackServicesShrink(v string) *UpdateGatewayRouteShrinkRequest {
	s.FallbackServicesShrink = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetGatewayId(v int64) *UpdateGatewayRouteShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetId(v int64) *UpdateGatewayRouteShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetName(v string) *UpdateGatewayRouteShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetPredicatesShrink(v string) *UpdateGatewayRouteShrinkRequest {
	s.PredicatesShrink = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetRedirectJSONShrink(v string) *UpdateGatewayRouteShrinkRequest {
	s.RedirectJSONShrink = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetRouteOrder(v int32) *UpdateGatewayRouteShrinkRequest {
	s.RouteOrder = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) SetServicesShrink(v string) *UpdateGatewayRouteShrinkRequest {
	s.ServicesShrink = &v
	return s
}

func (s *UpdateGatewayRouteShrinkRequest) Validate() error {
	return dara.Validate(s)
}
