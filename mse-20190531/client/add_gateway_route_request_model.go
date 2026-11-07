// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddGatewayRouteRequest
	GetAcceptLanguage() *string
	SetDescription(v string) *AddGatewayRouteRequest
	GetDescription() *string
	SetDestinationType(v string) *AddGatewayRouteRequest
	GetDestinationType() *string
	SetDirectResponseJSON(v *AddGatewayRouteRequestDirectResponseJSON) *AddGatewayRouteRequest
	GetDirectResponseJSON() *AddGatewayRouteRequestDirectResponseJSON
	SetDomainId(v int64) *AddGatewayRouteRequest
	GetDomainId() *int64
	SetDomainIdListJSON(v string) *AddGatewayRouteRequest
	GetDomainIdListJSON() *string
	SetEnableWaf(v bool) *AddGatewayRouteRequest
	GetEnableWaf() *bool
	SetFallback(v bool) *AddGatewayRouteRequest
	GetFallback() *bool
	SetFallbackServices(v []*AddGatewayRouteRequestFallbackServices) *AddGatewayRouteRequest
	GetFallbackServices() []*AddGatewayRouteRequestFallbackServices
	SetGatewayId(v int64) *AddGatewayRouteRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *AddGatewayRouteRequest
	GetGatewayUniqueId() *string
	SetName(v string) *AddGatewayRouteRequest
	GetName() *string
	SetPolicies(v string) *AddGatewayRouteRequest
	GetPolicies() *string
	SetPredicates(v *AddGatewayRouteRequestPredicates) *AddGatewayRouteRequest
	GetPredicates() *AddGatewayRouteRequestPredicates
	SetRedirectJSON(v *AddGatewayRouteRequestRedirectJSON) *AddGatewayRouteRequest
	GetRedirectJSON() *AddGatewayRouteRequestRedirectJSON
	SetRouteOrder(v int32) *AddGatewayRouteRequest
	GetRouteOrder() *int32
	SetRouteType(v string) *AddGatewayRouteRequest
	GetRouteType() *string
	SetServices(v []*AddGatewayRouteRequestServices) *AddGatewayRouteRequest
	GetServices() []*AddGatewayRouteRequestServices
}

type AddGatewayRouteRequest struct {
	// The language of the returned information. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The route description.
	//
	// example:
	//
	// a route for xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination service type. Valid values:
	//
	// - Single: single service
	//
	// - Multiple: multiple services
	//
	// - VersionOriented: tag-based routing
	//
	// - Mock: mock
	//
	// - Redirect: redirect
	//
	// example:
	//
	// Multiple
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The mock response configuration.
	DirectResponseJSON *AddGatewayRouteRequestDirectResponseJSON `json:"DirectResponseJSON,omitempty" xml:"DirectResponseJSON,omitempty" type:"Struct"`
	// The domain name ID.
	//
	// example:
	//
	// 20
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The domain name list.
	//
	// example:
	//
	// [0,94]
	DomainIdListJSON *string `json:"DomainIdListJSON,omitempty" xml:"DomainIdListJSON,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- Use UpdateGatewayRouteWafStatus to update the WAF status of the route.
	//
	// example:
	//
	// true
	EnableWaf *bool `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	// Specifies whether to enable the fallback service.
	//
	// example:
	//
	// true
	Fallback *bool `json:"Fallback,omitempty" xml:"Fallback,omitempty"`
	// The fallback service.
	FallbackServices []*AddGatewayRouteRequestFallbackServices `json:"FallbackServices,omitempty" xml:"FallbackServices,omitempty" type:"Repeated"`
	// The gateway ID.
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
	// The route name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The JSON string of the route policy.
	//
	// example:
	//
	// {"CORS":"{\\"allowMethods\\":\\"GET,POST,PUT,DELETE,HEAD,OPTIONS,PATCH\\",\\"allowHeaders\\":\\"*\\",\\"exposeHeaders\\":\\"*\\",\\"unitNum\\":12,\\"allowCredentials\\":true,\\"status\\":\\"off\\",\\"allowOrigins\\":\\"*\\",\\"timeUnit\\":\\"h\\"}","Timeout":"{\\"unitNum\\":10,\\"timeUnit\\":\\"s\\",\\"status\\":\\"off\\"}","Retry":"{\\"attempts\\":2,\\"retryOn\\":[\\"5xx\\"],\\"status\\":\\"off\\"}","HTTPRewrite":"{\\"pathType\\":\\"EQUAL\\",\\"path\\":\\"/o\\",\\"status\\":\\"off\\"}","Waf":"{\\"enabled\\":false}","HeaderOp":"{\\"status\\":\\"off\\",\\"headerOpItems\\":[{\\"directionType\\":\\"Request\\",\\"opType\\":\\"Add\\",\\"key\\":\\"kkk\\",\\"value\\":\\"ll\\"}]}"}
	Policies *string `json:"Policies,omitempty" xml:"Policies,omitempty"`
	// The matching rules.
	Predicates *AddGatewayRouteRequestPredicates `json:"Predicates,omitempty" xml:"Predicates,omitempty" type:"Struct"`
	// The redirect configuration.
	RedirectJSON *AddGatewayRouteRequestRedirectJSON `json:"RedirectJSON,omitempty" xml:"RedirectJSON,omitempty" type:"Struct"`
	// The route order. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	RouteOrder *int32 `json:"RouteOrder,omitempty" xml:"RouteOrder,omitempty"`
	// The route type. Valid values:
	//
	// - Op: management and control route
	//
	// example:
	//
	// Op
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The service list. This parameter is required when the route points to a single service, multiple services, or tag-based routing.
	Services []*AddGatewayRouteRequestServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s AddGatewayRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteRequest) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddGatewayRouteRequest) GetDescription() *string {
	return s.Description
}

func (s *AddGatewayRouteRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *AddGatewayRouteRequest) GetDirectResponseJSON() *AddGatewayRouteRequestDirectResponseJSON {
	return s.DirectResponseJSON
}

func (s *AddGatewayRouteRequest) GetDomainId() *int64 {
	return s.DomainId
}

func (s *AddGatewayRouteRequest) GetDomainIdListJSON() *string {
	return s.DomainIdListJSON
}

func (s *AddGatewayRouteRequest) GetEnableWaf() *bool {
	return s.EnableWaf
}

func (s *AddGatewayRouteRequest) GetFallback() *bool {
	return s.Fallback
}

func (s *AddGatewayRouteRequest) GetFallbackServices() []*AddGatewayRouteRequestFallbackServices {
	return s.FallbackServices
}

func (s *AddGatewayRouteRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *AddGatewayRouteRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddGatewayRouteRequest) GetName() *string {
	return s.Name
}

func (s *AddGatewayRouteRequest) GetPolicies() *string {
	return s.Policies
}

func (s *AddGatewayRouteRequest) GetPredicates() *AddGatewayRouteRequestPredicates {
	return s.Predicates
}

func (s *AddGatewayRouteRequest) GetRedirectJSON() *AddGatewayRouteRequestRedirectJSON {
	return s.RedirectJSON
}

func (s *AddGatewayRouteRequest) GetRouteOrder() *int32 {
	return s.RouteOrder
}

func (s *AddGatewayRouteRequest) GetRouteType() *string {
	return s.RouteType
}

func (s *AddGatewayRouteRequest) GetServices() []*AddGatewayRouteRequestServices {
	return s.Services
}

func (s *AddGatewayRouteRequest) SetAcceptLanguage(v string) *AddGatewayRouteRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddGatewayRouteRequest) SetDescription(v string) *AddGatewayRouteRequest {
	s.Description = &v
	return s
}

func (s *AddGatewayRouteRequest) SetDestinationType(v string) *AddGatewayRouteRequest {
	s.DestinationType = &v
	return s
}

func (s *AddGatewayRouteRequest) SetDirectResponseJSON(v *AddGatewayRouteRequestDirectResponseJSON) *AddGatewayRouteRequest {
	s.DirectResponseJSON = v
	return s
}

func (s *AddGatewayRouteRequest) SetDomainId(v int64) *AddGatewayRouteRequest {
	s.DomainId = &v
	return s
}

func (s *AddGatewayRouteRequest) SetDomainIdListJSON(v string) *AddGatewayRouteRequest {
	s.DomainIdListJSON = &v
	return s
}

func (s *AddGatewayRouteRequest) SetEnableWaf(v bool) *AddGatewayRouteRequest {
	s.EnableWaf = &v
	return s
}

func (s *AddGatewayRouteRequest) SetFallback(v bool) *AddGatewayRouteRequest {
	s.Fallback = &v
	return s
}

func (s *AddGatewayRouteRequest) SetFallbackServices(v []*AddGatewayRouteRequestFallbackServices) *AddGatewayRouteRequest {
	s.FallbackServices = v
	return s
}

func (s *AddGatewayRouteRequest) SetGatewayId(v int64) *AddGatewayRouteRequest {
	s.GatewayId = &v
	return s
}

func (s *AddGatewayRouteRequest) SetGatewayUniqueId(v string) *AddGatewayRouteRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddGatewayRouteRequest) SetName(v string) *AddGatewayRouteRequest {
	s.Name = &v
	return s
}

func (s *AddGatewayRouteRequest) SetPolicies(v string) *AddGatewayRouteRequest {
	s.Policies = &v
	return s
}

func (s *AddGatewayRouteRequest) SetPredicates(v *AddGatewayRouteRequestPredicates) *AddGatewayRouteRequest {
	s.Predicates = v
	return s
}

func (s *AddGatewayRouteRequest) SetRedirectJSON(v *AddGatewayRouteRequestRedirectJSON) *AddGatewayRouteRequest {
	s.RedirectJSON = v
	return s
}

func (s *AddGatewayRouteRequest) SetRouteOrder(v int32) *AddGatewayRouteRequest {
	s.RouteOrder = &v
	return s
}

func (s *AddGatewayRouteRequest) SetRouteType(v string) *AddGatewayRouteRequest {
	s.RouteType = &v
	return s
}

func (s *AddGatewayRouteRequest) SetServices(v []*AddGatewayRouteRequestServices) *AddGatewayRouteRequest {
	s.Services = v
	return s
}

func (s *AddGatewayRouteRequest) Validate() error {
	if s.DirectResponseJSON != nil {
		if err := s.DirectResponseJSON.Validate(); err != nil {
			return err
		}
	}
	if s.FallbackServices != nil {
		for _, item := range s.FallbackServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Predicates != nil {
		if err := s.Predicates.Validate(); err != nil {
			return err
		}
	}
	if s.RedirectJSON != nil {
		if err := s.RedirectJSON.Validate(); err != nil {
			return err
		}
	}
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddGatewayRouteRequestDirectResponseJSON struct {
	// The mock response body.
	//
	// example:
	//
	// hello
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The mock response code.
	//
	// example:
	//
	// 403
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AddGatewayRouteRequestDirectResponseJSON) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteRequestDirectResponseJSON) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteRequestDirectResponseJSON) GetBody() *string {
	return s.Body
}

func (s *AddGatewayRouteRequestDirectResponseJSON) GetCode() *int64 {
	return s.Code
}

func (s *AddGatewayRouteRequestDirectResponseJSON) SetBody(v string) *AddGatewayRouteRequestDirectResponseJSON {
	s.Body = &v
	return s
}

func (s *AddGatewayRouteRequestDirectResponseJSON) SetCode(v int64) *AddGatewayRouteRequestDirectResponseJSON {
	s.Code = &v
	return s
}

func (s *AddGatewayRouteRequestDirectResponseJSON) Validate() error {
	return dara.Validate(s)
}

type AddGatewayRouteRequestFallbackServices struct {
	// The protocol type.
	//
	// example:
	//
	// DUBBO
	AgreementType *string `json:"AgreementType,omitempty" xml:"AgreementType,omitempty"`
	// The service group name.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name.
	//
	// example:
	//
	// user
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace where the service resides.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The weight percentage.
	//
	// example:
	//
	// 80
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The service ID.
	//
	// example:
	//
	// 353
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service port.
	//
	// example:
	//
	// 443
	ServicePort *int32 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The source type.
	//
	// example:
	//
	// MSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The service version.
	//
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s AddGatewayRouteRequestFallbackServices) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteRequestFallbackServices) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteRequestFallbackServices) GetAgreementType() *string {
	return s.AgreementType
}

func (s *AddGatewayRouteRequestFallbackServices) GetGroupName() *string {
	return s.GroupName
}

func (s *AddGatewayRouteRequestFallbackServices) GetName() *string {
	return s.Name
}

func (s *AddGatewayRouteRequestFallbackServices) GetNamespace() *string {
	return s.Namespace
}

func (s *AddGatewayRouteRequestFallbackServices) GetPercent() *int32 {
	return s.Percent
}

func (s *AddGatewayRouteRequestFallbackServices) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *AddGatewayRouteRequestFallbackServices) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *AddGatewayRouteRequestFallbackServices) GetSourceType() *string {
	return s.SourceType
}

func (s *AddGatewayRouteRequestFallbackServices) GetVersion() *string {
	return s.Version
}

func (s *AddGatewayRouteRequestFallbackServices) SetAgreementType(v string) *AddGatewayRouteRequestFallbackServices {
	s.AgreementType = &v
	return s
}

func (s *AddGatewayRouteRequestFallbackServices) SetGroupName(v string) *AddGatewayRouteRequestFallbackServices {
	s.GroupName = &v
	return s
}

func (s *AddGatewayRouteRequestFallbackServices) SetName(v string) *AddGatewayRouteRequestFallbackServices {
	s.Name = &v
	return s
}

func (s *AddGatewayRouteRequestFallbackServices) SetNamespace(v string) *AddGatewayRouteRequestFallbackServices {
	s.Namespace = &v
	return s
}

func (s *AddGatewayRouteRequestFallbackServices) SetPercent(v int32) *AddGatewayRouteRequestFallbackServices {
	s.Percent = &v
	return s
}

func (s *AddGatewayRouteRequestFallbackServices) SetServiceId(v int64) *AddGatewayRouteRequestFallbackServices {
	s.ServiceId = &v
	return s
}

func (s *AddGatewayRouteRequestFallbackServices) SetServicePort(v int32) *AddGatewayRouteRequestFallbackServices {
	s.ServicePort = &v
	return s
}

func (s *AddGatewayRouteRequestFallbackServices) SetSourceType(v string) *AddGatewayRouteRequestFallbackServices {
	s.SourceType = &v
	return s
}

func (s *AddGatewayRouteRequestFallbackServices) SetVersion(v string) *AddGatewayRouteRequestFallbackServices {
	s.Version = &v
	return s
}

func (s *AddGatewayRouteRequestFallbackServices) Validate() error {
	return dara.Validate(s)
}

type AddGatewayRouteRequestPredicates struct {
	// The request header matching rules.
	HeaderPredicates []*AddGatewayRouteRequestPredicatesHeaderPredicates `json:"HeaderPredicates,omitempty" xml:"HeaderPredicates,omitempty" type:"Repeated"`
	// The method matching rules.
	MethodPredicates []*string `json:"MethodPredicates,omitempty" xml:"MethodPredicates,omitempty" type:"Repeated"`
	// The path matching rules.
	PathPredicates *AddGatewayRouteRequestPredicatesPathPredicates `json:"PathPredicates,omitempty" xml:"PathPredicates,omitempty" type:"Struct"`
	// The URL parameter matching rules.
	QueryPredicates []*AddGatewayRouteRequestPredicatesQueryPredicates `json:"QueryPredicates,omitempty" xml:"QueryPredicates,omitempty" type:"Repeated"`
}

func (s AddGatewayRouteRequestPredicates) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteRequestPredicates) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteRequestPredicates) GetHeaderPredicates() []*AddGatewayRouteRequestPredicatesHeaderPredicates {
	return s.HeaderPredicates
}

func (s *AddGatewayRouteRequestPredicates) GetMethodPredicates() []*string {
	return s.MethodPredicates
}

func (s *AddGatewayRouteRequestPredicates) GetPathPredicates() *AddGatewayRouteRequestPredicatesPathPredicates {
	return s.PathPredicates
}

func (s *AddGatewayRouteRequestPredicates) GetQueryPredicates() []*AddGatewayRouteRequestPredicatesQueryPredicates {
	return s.QueryPredicates
}

func (s *AddGatewayRouteRequestPredicates) SetHeaderPredicates(v []*AddGatewayRouteRequestPredicatesHeaderPredicates) *AddGatewayRouteRequestPredicates {
	s.HeaderPredicates = v
	return s
}

func (s *AddGatewayRouteRequestPredicates) SetMethodPredicates(v []*string) *AddGatewayRouteRequestPredicates {
	s.MethodPredicates = v
	return s
}

func (s *AddGatewayRouteRequestPredicates) SetPathPredicates(v *AddGatewayRouteRequestPredicatesPathPredicates) *AddGatewayRouteRequestPredicates {
	s.PathPredicates = v
	return s
}

func (s *AddGatewayRouteRequestPredicates) SetQueryPredicates(v []*AddGatewayRouteRequestPredicatesQueryPredicates) *AddGatewayRouteRequestPredicates {
	s.QueryPredicates = v
	return s
}

func (s *AddGatewayRouteRequestPredicates) Validate() error {
	if s.HeaderPredicates != nil {
		for _, item := range s.HeaderPredicates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PathPredicates != nil {
		if err := s.PathPredicates.Validate(); err != nil {
			return err
		}
	}
	if s.QueryPredicates != nil {
		for _, item := range s.QueryPredicates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddGatewayRouteRequestPredicatesHeaderPredicates struct {
	// The request header key.
	//
	// example:
	//
	// debug
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching type.
	//
	// example:
	//
	// PRE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The corresponding value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddGatewayRouteRequestPredicatesHeaderPredicates) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteRequestPredicatesHeaderPredicates) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteRequestPredicatesHeaderPredicates) GetKey() *string {
	return s.Key
}

func (s *AddGatewayRouteRequestPredicatesHeaderPredicates) GetType() *string {
	return s.Type
}

func (s *AddGatewayRouteRequestPredicatesHeaderPredicates) GetValue() *string {
	return s.Value
}

func (s *AddGatewayRouteRequestPredicatesHeaderPredicates) SetKey(v string) *AddGatewayRouteRequestPredicatesHeaderPredicates {
	s.Key = &v
	return s
}

func (s *AddGatewayRouteRequestPredicatesHeaderPredicates) SetType(v string) *AddGatewayRouteRequestPredicatesHeaderPredicates {
	s.Type = &v
	return s
}

func (s *AddGatewayRouteRequestPredicatesHeaderPredicates) SetValue(v string) *AddGatewayRouteRequestPredicatesHeaderPredicates {
	s.Value = &v
	return s
}

func (s *AddGatewayRouteRequestPredicatesHeaderPredicates) Validate() error {
	return dara.Validate(s)
}

type AddGatewayRouteRequestPredicatesPathPredicates struct {
	// Specifies whether to ignore case.
	//
	// example:
	//
	// true
	IgnoreCase *bool `json:"IgnoreCase,omitempty" xml:"IgnoreCase,omitempty"`
	// The path.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The route matching type. Valid values:
	//
	// - PRE: prefix match
	//
	// - EQUAL: exact match
	//
	// - ERGULAR: regular expression match
	//
	// example:
	//
	// PRE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddGatewayRouteRequestPredicatesPathPredicates) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteRequestPredicatesPathPredicates) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteRequestPredicatesPathPredicates) GetIgnoreCase() *bool {
	return s.IgnoreCase
}

func (s *AddGatewayRouteRequestPredicatesPathPredicates) GetPath() *string {
	return s.Path
}

func (s *AddGatewayRouteRequestPredicatesPathPredicates) GetType() *string {
	return s.Type
}

func (s *AddGatewayRouteRequestPredicatesPathPredicates) SetIgnoreCase(v bool) *AddGatewayRouteRequestPredicatesPathPredicates {
	s.IgnoreCase = &v
	return s
}

func (s *AddGatewayRouteRequestPredicatesPathPredicates) SetPath(v string) *AddGatewayRouteRequestPredicatesPathPredicates {
	s.Path = &v
	return s
}

func (s *AddGatewayRouteRequestPredicatesPathPredicates) SetType(v string) *AddGatewayRouteRequestPredicatesPathPredicates {
	s.Type = &v
	return s
}

func (s *AddGatewayRouteRequestPredicatesPathPredicates) Validate() error {
	return dara.Validate(s)
}

type AddGatewayRouteRequestPredicatesQueryPredicates struct {
	// The parameter name.
	//
	// example:
	//
	// userid
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching type.
	//
	// example:
	//
	// PRE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The matching value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddGatewayRouteRequestPredicatesQueryPredicates) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteRequestPredicatesQueryPredicates) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteRequestPredicatesQueryPredicates) GetKey() *string {
	return s.Key
}

func (s *AddGatewayRouteRequestPredicatesQueryPredicates) GetType() *string {
	return s.Type
}

func (s *AddGatewayRouteRequestPredicatesQueryPredicates) GetValue() *string {
	return s.Value
}

func (s *AddGatewayRouteRequestPredicatesQueryPredicates) SetKey(v string) *AddGatewayRouteRequestPredicatesQueryPredicates {
	s.Key = &v
	return s
}

func (s *AddGatewayRouteRequestPredicatesQueryPredicates) SetType(v string) *AddGatewayRouteRequestPredicatesQueryPredicates {
	s.Type = &v
	return s
}

func (s *AddGatewayRouteRequestPredicatesQueryPredicates) SetValue(v string) *AddGatewayRouteRequestPredicatesQueryPredicates {
	s.Value = &v
	return s
}

func (s *AddGatewayRouteRequestPredicatesQueryPredicates) Validate() error {
	return dara.Validate(s)
}

type AddGatewayRouteRequestRedirectJSON struct {
	// The response code.
	//
	// example:
	//
	// 302
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The redirect hostname.
	//
	// example:
	//
	// test.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The redirect path.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s AddGatewayRouteRequestRedirectJSON) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteRequestRedirectJSON) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteRequestRedirectJSON) GetCode() *int32 {
	return s.Code
}

func (s *AddGatewayRouteRequestRedirectJSON) GetHost() *string {
	return s.Host
}

func (s *AddGatewayRouteRequestRedirectJSON) GetPath() *string {
	return s.Path
}

func (s *AddGatewayRouteRequestRedirectJSON) SetCode(v int32) *AddGatewayRouteRequestRedirectJSON {
	s.Code = &v
	return s
}

func (s *AddGatewayRouteRequestRedirectJSON) SetHost(v string) *AddGatewayRouteRequestRedirectJSON {
	s.Host = &v
	return s
}

func (s *AddGatewayRouteRequestRedirectJSON) SetPath(v string) *AddGatewayRouteRequestRedirectJSON {
	s.Path = &v
	return s
}

func (s *AddGatewayRouteRequestRedirectJSON) Validate() error {
	return dara.Validate(s)
}

type AddGatewayRouteRequestServices struct {
	// The protocol type.
	//
	// example:
	//
	// DUBBO
	AgreementType *string `json:"AgreementType,omitempty" xml:"AgreementType,omitempty"`
	// The service group name.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The Dubbo protocol transcoding configuration.
	HttpDubboTranscoder *AddGatewayRouteRequestServicesHttpDubboTranscoder `json:"HttpDubboTranscoder,omitempty" xml:"HttpDubboTranscoder,omitempty" type:"Struct"`
	// The name.
	//
	// example:
	//
	// user
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace where the service resides.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The traffic distribution percentage for the current service or service version. This parameter is required only when multiple services or tag-based routing is used.
	//
	// example:
	//
	// 80
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The service ID.
	//
	// example:
	//
	// 353
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service port.
	//
	// example:
	//
	// 443
	ServicePort *int32 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The source type.
	//
	// example:
	//
	// MSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The service version. This parameter is required only when tag-based routing is used.
	//
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s AddGatewayRouteRequestServices) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteRequestServices) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteRequestServices) GetAgreementType() *string {
	return s.AgreementType
}

func (s *AddGatewayRouteRequestServices) GetGroupName() *string {
	return s.GroupName
}

func (s *AddGatewayRouteRequestServices) GetHttpDubboTranscoder() *AddGatewayRouteRequestServicesHttpDubboTranscoder {
	return s.HttpDubboTranscoder
}

func (s *AddGatewayRouteRequestServices) GetName() *string {
	return s.Name
}

func (s *AddGatewayRouteRequestServices) GetNamespace() *string {
	return s.Namespace
}

func (s *AddGatewayRouteRequestServices) GetPercent() *int32 {
	return s.Percent
}

func (s *AddGatewayRouteRequestServices) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *AddGatewayRouteRequestServices) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *AddGatewayRouteRequestServices) GetSourceType() *string {
	return s.SourceType
}

func (s *AddGatewayRouteRequestServices) GetVersion() *string {
	return s.Version
}

func (s *AddGatewayRouteRequestServices) SetAgreementType(v string) *AddGatewayRouteRequestServices {
	s.AgreementType = &v
	return s
}

func (s *AddGatewayRouteRequestServices) SetGroupName(v string) *AddGatewayRouteRequestServices {
	s.GroupName = &v
	return s
}

func (s *AddGatewayRouteRequestServices) SetHttpDubboTranscoder(v *AddGatewayRouteRequestServicesHttpDubboTranscoder) *AddGatewayRouteRequestServices {
	s.HttpDubboTranscoder = v
	return s
}

func (s *AddGatewayRouteRequestServices) SetName(v string) *AddGatewayRouteRequestServices {
	s.Name = &v
	return s
}

func (s *AddGatewayRouteRequestServices) SetNamespace(v string) *AddGatewayRouteRequestServices {
	s.Namespace = &v
	return s
}

func (s *AddGatewayRouteRequestServices) SetPercent(v int32) *AddGatewayRouteRequestServices {
	s.Percent = &v
	return s
}

func (s *AddGatewayRouteRequestServices) SetServiceId(v int64) *AddGatewayRouteRequestServices {
	s.ServiceId = &v
	return s
}

func (s *AddGatewayRouteRequestServices) SetServicePort(v int32) *AddGatewayRouteRequestServices {
	s.ServicePort = &v
	return s
}

func (s *AddGatewayRouteRequestServices) SetSourceType(v string) *AddGatewayRouteRequestServices {
	s.SourceType = &v
	return s
}

func (s *AddGatewayRouteRequestServices) SetVersion(v string) *AddGatewayRouteRequestServices {
	s.Version = &v
	return s
}

func (s *AddGatewayRouteRequestServices) Validate() error {
	if s.HttpDubboTranscoder != nil {
		if err := s.HttpDubboTranscoder.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddGatewayRouteRequestServicesHttpDubboTranscoder struct {
	// The service group.
	//
	// example:
	//
	// group
	DubboServiceGroup *string `json:"DubboServiceGroup,omitempty" xml:"DubboServiceGroup,omitempty"`
	// The Dubbo service name.
	//
	// example:
	//
	// org.apache.dubbo.samples.basic.api.DemoService
	DubboServiceName *string `json:"DubboServiceName,omitempty" xml:"DubboServiceName,omitempty"`
	// The Dubbo service version.
	//
	// example:
	//
	// 0.0.0
	DubboServiceVersion *string `json:"DubboServiceVersion,omitempty" xml:"DubboServiceVersion,omitempty"`
	// The list of Dubbo forwarding rules.
	MothedMapList []*AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList `json:"MothedMapList,omitempty" xml:"MothedMapList,omitempty" type:"Repeated"`
}

func (s AddGatewayRouteRequestServicesHttpDubboTranscoder) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteRequestServicesHttpDubboTranscoder) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoder) GetDubboServiceGroup() *string {
	return s.DubboServiceGroup
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoder) GetDubboServiceName() *string {
	return s.DubboServiceName
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoder) GetDubboServiceVersion() *string {
	return s.DubboServiceVersion
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoder) GetMothedMapList() []*AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	return s.MothedMapList
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoder) SetDubboServiceGroup(v string) *AddGatewayRouteRequestServicesHttpDubboTranscoder {
	s.DubboServiceGroup = &v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoder) SetDubboServiceName(v string) *AddGatewayRouteRequestServicesHttpDubboTranscoder {
	s.DubboServiceName = &v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoder) SetDubboServiceVersion(v string) *AddGatewayRouteRequestServicesHttpDubboTranscoder {
	s.DubboServiceVersion = &v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoder) SetMothedMapList(v []*AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) *AddGatewayRouteRequestServicesHttpDubboTranscoder {
	s.MothedMapList = v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoder) Validate() error {
	if s.MothedMapList != nil {
		for _, item := range s.MothedMapList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList struct {
	// The Dubbo method name.
	//
	// example:
	//
	// sayHello
	DubboMothedName *string `json:"DubboMothedName,omitempty" xml:"DubboMothedName,omitempty"`
	// The HTTP method. Valid values:
	//
	// - ALL_GET: all GET methods
	//
	// - ALL_POST: all POST methods
	//
	// - ALL_PUT: all PUT methods
	//
	// - ALL_DELETE: all DELETE methods
	//
	// - ALL_PATCH: all PATCH methods
	//
	// example:
	//
	// ALL_GET
	HttpMothed *string `json:"HttpMothed,omitempty" xml:"HttpMothed,omitempty"`
	// The method matching path.
	//
	// example:
	//
	// /mytestzbk/sayhello
	Mothedpath *string `json:"Mothedpath,omitempty" xml:"Mothedpath,omitempty"`
	// The parameter mapping list.
	ParamMapsList []*AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList `json:"ParamMapsList,omitempty" xml:"ParamMapsList,omitempty" type:"Repeated"`
	// The header pass-through type. Valid values:
	//
	// - PASS_ALL: pass through all headers
	//
	// - PASS_NOT: do not pass through headers
	//
	// - PASS_ASSIGN: pass through specified headers
	//
	// example:
	//
	// PASS_NOT
	PassThroughAllHeaders *string `json:"PassThroughAllHeaders,omitempty" xml:"PassThroughAllHeaders,omitempty"`
	// The list of specified headers to pass through.
	PassThroughList []*string `json:"PassThroughList,omitempty" xml:"PassThroughList,omitempty" type:"Repeated"`
}

func (s AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GetDubboMothedName() *string {
	return s.DubboMothedName
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GetHttpMothed() *string {
	return s.HttpMothed
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GetMothedpath() *string {
	return s.Mothedpath
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GetParamMapsList() []*AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList {
	return s.ParamMapsList
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GetPassThroughAllHeaders() *string {
	return s.PassThroughAllHeaders
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GetPassThroughList() []*string {
	return s.PassThroughList
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) SetDubboMothedName(v string) *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	s.DubboMothedName = &v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) SetHttpMothed(v string) *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	s.HttpMothed = &v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) SetMothedpath(v string) *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	s.Mothedpath = &v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) SetParamMapsList(v []*AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	s.ParamMapsList = v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) SetPassThroughAllHeaders(v string) *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	s.PassThroughAllHeaders = &v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) SetPassThroughList(v []*string) *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	s.PassThroughList = v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) Validate() error {
	if s.ParamMapsList != nil {
		for _, item := range s.ParamMapsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList struct {
	// The key extracted from the input parameter location.
	//
	// example:
	//
	// name
	ExtractKey *string `json:"ExtractKey,omitempty" xml:"ExtractKey,omitempty"`
	// The input parameter location. Valid values:
	//
	// - ALL_QUERY_PARAMETER: request parameters
	//
	// - ALL_HEADER: request headers
	//
	// - ALL_PATH: URI of the request
	//
	// - ALL_BODY: request body
	//
	// example:
	//
	// ALL_QUERY_PARAMETER
	ExtractKeySpec *string `json:"ExtractKeySpec,omitempty" xml:"ExtractKeySpec,omitempty"`
	// The backend parameter type.
	//
	// example:
	//
	// java.lang.String
	MappingType *string `json:"MappingType,omitempty" xml:"MappingType,omitempty"`
}

func (s AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) GoString() string {
	return s.String()
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) GetExtractKey() *string {
	return s.ExtractKey
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) GetExtractKeySpec() *string {
	return s.ExtractKeySpec
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) GetMappingType() *string {
	return s.MappingType
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) SetExtractKey(v string) *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList {
	s.ExtractKey = &v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) SetExtractKeySpec(v string) *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList {
	s.ExtractKeySpec = &v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) SetMappingType(v string) *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList {
	s.MappingType = &v
	return s
}

func (s *AddGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) Validate() error {
	return dara.Validate(s)
}
