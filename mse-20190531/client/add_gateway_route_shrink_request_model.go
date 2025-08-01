// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayRouteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddGatewayRouteShrinkRequest
	GetAcceptLanguage() *string
	SetDescription(v string) *AddGatewayRouteShrinkRequest
	GetDescription() *string
	SetDestinationType(v string) *AddGatewayRouteShrinkRequest
	GetDestinationType() *string
	SetDirectResponseJSONShrink(v string) *AddGatewayRouteShrinkRequest
	GetDirectResponseJSONShrink() *string
	SetDomainId(v int64) *AddGatewayRouteShrinkRequest
	GetDomainId() *int64
	SetDomainIdListJSON(v string) *AddGatewayRouteShrinkRequest
	GetDomainIdListJSON() *string
	SetEnableWaf(v bool) *AddGatewayRouteShrinkRequest
	GetEnableWaf() *bool
	SetFallback(v bool) *AddGatewayRouteShrinkRequest
	GetFallback() *bool
	SetFallbackServicesShrink(v string) *AddGatewayRouteShrinkRequest
	GetFallbackServicesShrink() *string
	SetGatewayId(v int64) *AddGatewayRouteShrinkRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *AddGatewayRouteShrinkRequest
	GetGatewayUniqueId() *string
	SetName(v string) *AddGatewayRouteShrinkRequest
	GetName() *string
	SetPolicies(v string) *AddGatewayRouteShrinkRequest
	GetPolicies() *string
	SetPredicatesShrink(v string) *AddGatewayRouteShrinkRequest
	GetPredicatesShrink() *string
	SetRedirectJSONShrink(v string) *AddGatewayRouteShrinkRequest
	GetRedirectJSONShrink() *string
	SetRouteOrder(v int32) *AddGatewayRouteShrinkRequest
	GetRouteOrder() *int32
	SetRouteType(v string) *AddGatewayRouteShrinkRequest
	GetRouteType() *string
	SetServicesShrink(v string) *AddGatewayRouteShrinkRequest
	GetServicesShrink() *string
}

type AddGatewayRouteShrinkRequest struct {
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
	// The type of the destination service. Valid values:
	//
	// 	- Single
	//
	// 	- Multiple
	//
	// 	- VersionOriented
	//
	// 	- Mock
	//
	// 	- Redirect
	//
	// example:
	//
	// Multiple
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The mock response configuration.
	DirectResponseJSONShrink *string `json:"DirectResponseJSON,omitempty" xml:"DirectResponseJSON,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// 20
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The domain IDs.
	//
	// example:
	//
	// [0,94]
	DomainIdListJSON *string `json:"DomainIdListJSON,omitempty" xml:"DomainIdListJSON,omitempty"`
	// Deprecated
	//
	// Specifies whether to activate Web Application Firewall (WAF).
	//
	// example:
	//
	// true
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
	// 526
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-492af9b04bb4474cae9d645be8*****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The routing policy in a JSON string.
	//
	// example:
	//
	// {"CORS":"{\\"allowMethods\\":\\"GET,POST,PUT,DELETE,HEAD,OPTIONS,PATCH\\",\\"allowHeaders\\":\\"*\\",\\"exposeHeaders\\":\\"*\\",\\"unitNum\\":12,\\"allowCredentials\\":true,\\"status\\":\\"off\\",\\"allowOrigins\\":\\"*\\",\\"timeUnit\\":\\"h\\"}","Timeout":"{\\"unitNum\\":10,\\"timeUnit\\":\\"s\\",\\"status\\":\\"off\\"}","Retry":"{\\"attempts\\":2,\\"retryOn\\":[\\"5xx\\"],\\"status\\":\\"off\\"}","HTTPRewrite":"{\\"pathType\\":\\"EQUAL\\",\\"path\\":\\"/o\\",\\"status\\":\\"off\\"}","Waf":"{\\"enabled\\":false}","HeaderOp":"{\\"status\\":\\"off\\",\\"headerOpItems\\":[{\\"directionType\\":\\"Request\\",\\"opType\\":\\"Add\\",\\"key\\":\\"kkk\\",\\"value\\":\\"ll\\"}]}"}
	Policies *string `json:"Policies,omitempty" xml:"Policies,omitempty"`
	// The matching rule.
	PredicatesShrink *string `json:"Predicates,omitempty" xml:"Predicates,omitempty"`
	// The configuration of the redirection.
	RedirectJSONShrink *string `json:"RedirectJSON,omitempty" xml:"RedirectJSON,omitempty"`
	// The sequence number of the route. (A small value indicates a high priority.)
	//
	// example:
	//
	// 1
	RouteOrder *int32 `json:"RouteOrder,omitempty" xml:"RouteOrder,omitempty"`
	// The route type. Valid values:
	//
	// Op: Manage routes.
	//
	// example:
	//
	// Op
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The list of services.
	ServicesShrink *string `json:"Services,omitempty" xml:"Services,omitempty"`
}

func (s AddGatewayRouteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddGatewayRouteShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *AddGatewayRouteShrinkRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *AddGatewayRouteShrinkRequest) GetDirectResponseJSONShrink() *string {
	return s.DirectResponseJSONShrink
}

func (s *AddGatewayRouteShrinkRequest) GetDomainId() *int64 {
	return s.DomainId
}

func (s *AddGatewayRouteShrinkRequest) GetDomainIdListJSON() *string {
	return s.DomainIdListJSON
}

func (s *AddGatewayRouteShrinkRequest) GetEnableWaf() *bool {
	return s.EnableWaf
}

func (s *AddGatewayRouteShrinkRequest) GetFallback() *bool {
	return s.Fallback
}

func (s *AddGatewayRouteShrinkRequest) GetFallbackServicesShrink() *string {
	return s.FallbackServicesShrink
}

func (s *AddGatewayRouteShrinkRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *AddGatewayRouteShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddGatewayRouteShrinkRequest) GetName() *string {
	return s.Name
}

func (s *AddGatewayRouteShrinkRequest) GetPolicies() *string {
	return s.Policies
}

func (s *AddGatewayRouteShrinkRequest) GetPredicatesShrink() *string {
	return s.PredicatesShrink
}

func (s *AddGatewayRouteShrinkRequest) GetRedirectJSONShrink() *string {
	return s.RedirectJSONShrink
}

func (s *AddGatewayRouteShrinkRequest) GetRouteOrder() *int32 {
	return s.RouteOrder
}

func (s *AddGatewayRouteShrinkRequest) GetRouteType() *string {
	return s.RouteType
}

func (s *AddGatewayRouteShrinkRequest) GetServicesShrink() *string {
	return s.ServicesShrink
}

func (s *AddGatewayRouteShrinkRequest) SetAcceptLanguage(v string) *AddGatewayRouteShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetDescription(v string) *AddGatewayRouteShrinkRequest {
	s.Description = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetDestinationType(v string) *AddGatewayRouteShrinkRequest {
	s.DestinationType = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetDirectResponseJSONShrink(v string) *AddGatewayRouteShrinkRequest {
	s.DirectResponseJSONShrink = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetDomainId(v int64) *AddGatewayRouteShrinkRequest {
	s.DomainId = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetDomainIdListJSON(v string) *AddGatewayRouteShrinkRequest {
	s.DomainIdListJSON = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetEnableWaf(v bool) *AddGatewayRouteShrinkRequest {
	s.EnableWaf = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetFallback(v bool) *AddGatewayRouteShrinkRequest {
	s.Fallback = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetFallbackServicesShrink(v string) *AddGatewayRouteShrinkRequest {
	s.FallbackServicesShrink = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetGatewayId(v int64) *AddGatewayRouteShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetGatewayUniqueId(v string) *AddGatewayRouteShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetName(v string) *AddGatewayRouteShrinkRequest {
	s.Name = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetPolicies(v string) *AddGatewayRouteShrinkRequest {
	s.Policies = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetPredicatesShrink(v string) *AddGatewayRouteShrinkRequest {
	s.PredicatesShrink = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetRedirectJSONShrink(v string) *AddGatewayRouteShrinkRequest {
	s.RedirectJSONShrink = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetRouteOrder(v int32) *AddGatewayRouteShrinkRequest {
	s.RouteOrder = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetRouteType(v string) *AddGatewayRouteShrinkRequest {
	s.RouteType = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) SetServicesShrink(v string) *AddGatewayRouteShrinkRequest {
	s.ServicesShrink = &v
	return s
}

func (s *AddGatewayRouteShrinkRequest) Validate() error {
	return dara.Validate(s)
}
