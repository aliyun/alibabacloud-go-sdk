// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteRequest
	GetAcceptLanguage() *string
	SetDescription(v string) *UpdateGatewayRouteRequest
	GetDescription() *string
	SetDestinationType(v string) *UpdateGatewayRouteRequest
	GetDestinationType() *string
	SetDirectResponseJSON(v *UpdateGatewayRouteRequestDirectResponseJSON) *UpdateGatewayRouteRequest
	GetDirectResponseJSON() *UpdateGatewayRouteRequestDirectResponseJSON
	SetDomainIdListJSON(v string) *UpdateGatewayRouteRequest
	GetDomainIdListJSON() *string
	SetEnableWaf(v bool) *UpdateGatewayRouteRequest
	GetEnableWaf() *bool
	SetFallback(v bool) *UpdateGatewayRouteRequest
	GetFallback() *bool
	SetFallbackServices(v []*UpdateGatewayRouteRequestFallbackServices) *UpdateGatewayRouteRequest
	GetFallbackServices() []*UpdateGatewayRouteRequestFallbackServices
	SetGatewayId(v int64) *UpdateGatewayRouteRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayRouteRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayRouteRequest
	GetId() *int64
	SetName(v string) *UpdateGatewayRouteRequest
	GetName() *string
	SetPredicates(v *UpdateGatewayRouteRequestPredicates) *UpdateGatewayRouteRequest
	GetPredicates() *UpdateGatewayRouteRequestPredicates
	SetRedirectJSON(v *UpdateGatewayRouteRequestRedirectJSON) *UpdateGatewayRouteRequest
	GetRedirectJSON() *UpdateGatewayRouteRequestRedirectJSON
	SetRouteOrder(v int32) *UpdateGatewayRouteRequest
	GetRouteOrder() *int32
	SetServices(v []*UpdateGatewayRouteRequestServices) *UpdateGatewayRouteRequest
	GetServices() []*UpdateGatewayRouteRequestServices
}

type UpdateGatewayRouteRequest struct {
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
	DirectResponseJSON *UpdateGatewayRouteRequestDirectResponseJSON `json:"DirectResponseJSON,omitempty" xml:"DirectResponseJSON,omitempty" type:"Struct"`
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
	FallbackServices []*UpdateGatewayRouteRequestFallbackServices `json:"FallbackServices,omitempty" xml:"FallbackServices,omitempty" type:"Repeated"`
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
	Predicates *UpdateGatewayRouteRequestPredicates `json:"Predicates,omitempty" xml:"Predicates,omitempty" type:"Struct"`
	// The information about redirection.
	RedirectJSON *UpdateGatewayRouteRequestRedirectJSON `json:"RedirectJSON,omitempty" xml:"RedirectJSON,omitempty" type:"Struct"`
	// The sequence number of the route.
	//
	// example:
	//
	// 1
	RouteOrder *int32 `json:"RouteOrder,omitempty" xml:"RouteOrder,omitempty"`
	// The information about destination services.
	Services []*UpdateGatewayRouteRequestServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s UpdateGatewayRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGatewayRouteRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *UpdateGatewayRouteRequest) GetDirectResponseJSON() *UpdateGatewayRouteRequestDirectResponseJSON {
	return s.DirectResponseJSON
}

func (s *UpdateGatewayRouteRequest) GetDomainIdListJSON() *string {
	return s.DomainIdListJSON
}

func (s *UpdateGatewayRouteRequest) GetEnableWaf() *bool {
	return s.EnableWaf
}

func (s *UpdateGatewayRouteRequest) GetFallback() *bool {
	return s.Fallback
}

func (s *UpdateGatewayRouteRequest) GetFallbackServices() []*UpdateGatewayRouteRequestFallbackServices {
	return s.FallbackServices
}

func (s *UpdateGatewayRouteRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteRequest) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayRouteRequest) GetPredicates() *UpdateGatewayRouteRequestPredicates {
	return s.Predicates
}

func (s *UpdateGatewayRouteRequest) GetRedirectJSON() *UpdateGatewayRouteRequestRedirectJSON {
	return s.RedirectJSON
}

func (s *UpdateGatewayRouteRequest) GetRouteOrder() *int32 {
	return s.RouteOrder
}

func (s *UpdateGatewayRouteRequest) GetServices() []*UpdateGatewayRouteRequestServices {
	return s.Services
}

func (s *UpdateGatewayRouteRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteRequest) SetDescription(v string) *UpdateGatewayRouteRequest {
	s.Description = &v
	return s
}

func (s *UpdateGatewayRouteRequest) SetDestinationType(v string) *UpdateGatewayRouteRequest {
	s.DestinationType = &v
	return s
}

func (s *UpdateGatewayRouteRequest) SetDirectResponseJSON(v *UpdateGatewayRouteRequestDirectResponseJSON) *UpdateGatewayRouteRequest {
	s.DirectResponseJSON = v
	return s
}

func (s *UpdateGatewayRouteRequest) SetDomainIdListJSON(v string) *UpdateGatewayRouteRequest {
	s.DomainIdListJSON = &v
	return s
}

func (s *UpdateGatewayRouteRequest) SetEnableWaf(v bool) *UpdateGatewayRouteRequest {
	s.EnableWaf = &v
	return s
}

func (s *UpdateGatewayRouteRequest) SetFallback(v bool) *UpdateGatewayRouteRequest {
	s.Fallback = &v
	return s
}

func (s *UpdateGatewayRouteRequest) SetFallbackServices(v []*UpdateGatewayRouteRequestFallbackServices) *UpdateGatewayRouteRequest {
	s.FallbackServices = v
	return s
}

func (s *UpdateGatewayRouteRequest) SetGatewayId(v int64) *UpdateGatewayRouteRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteRequest) SetId(v int64) *UpdateGatewayRouteRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteRequest) SetName(v string) *UpdateGatewayRouteRequest {
	s.Name = &v
	return s
}

func (s *UpdateGatewayRouteRequest) SetPredicates(v *UpdateGatewayRouteRequestPredicates) *UpdateGatewayRouteRequest {
	s.Predicates = v
	return s
}

func (s *UpdateGatewayRouteRequest) SetRedirectJSON(v *UpdateGatewayRouteRequestRedirectJSON) *UpdateGatewayRouteRequest {
	s.RedirectJSON = v
	return s
}

func (s *UpdateGatewayRouteRequest) SetRouteOrder(v int32) *UpdateGatewayRouteRequest {
	s.RouteOrder = &v
	return s
}

func (s *UpdateGatewayRouteRequest) SetServices(v []*UpdateGatewayRouteRequestServices) *UpdateGatewayRouteRequest {
	s.Services = v
	return s
}

func (s *UpdateGatewayRouteRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteRequestDirectResponseJSON struct {
	// The mock return value.
	//
	// example:
	//
	// hello
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The mock return code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateGatewayRouteRequestDirectResponseJSON) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRequestDirectResponseJSON) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequestDirectResponseJSON) GetBody() *string {
	return s.Body
}

func (s *UpdateGatewayRouteRequestDirectResponseJSON) GetCode() *int64 {
	return s.Code
}

func (s *UpdateGatewayRouteRequestDirectResponseJSON) SetBody(v string) *UpdateGatewayRouteRequestDirectResponseJSON {
	s.Body = &v
	return s
}

func (s *UpdateGatewayRouteRequestDirectResponseJSON) SetCode(v int64) *UpdateGatewayRouteRequestDirectResponseJSON {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteRequestDirectResponseJSON) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteRequestFallbackServices struct {
	// The type of the protocol. Valid values:
	//
	// example:
	//
	// DUBBO
	AgreementType *string `json:"AgreementType,omitempty" xml:"AgreementType,omitempty"`
	// The name of the group to which the service belongs.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace in which the service resides.
	//
	// example:
	//
	// Namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The weight in the form of a percentage value.
	//
	// example:
	//
	// 80
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 1
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service port number.
	//
	// example:
	//
	// 8848
	ServicePort *int32 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The source type.
	//
	// example:
	//
	// MSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateGatewayRouteRequestFallbackServices) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRequestFallbackServices) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequestFallbackServices) GetAgreementType() *string {
	return s.AgreementType
}

func (s *UpdateGatewayRouteRequestFallbackServices) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateGatewayRouteRequestFallbackServices) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayRouteRequestFallbackServices) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateGatewayRouteRequestFallbackServices) GetPercent() *int32 {
	return s.Percent
}

func (s *UpdateGatewayRouteRequestFallbackServices) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *UpdateGatewayRouteRequestFallbackServices) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *UpdateGatewayRouteRequestFallbackServices) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateGatewayRouteRequestFallbackServices) GetVersion() *string {
	return s.Version
}

func (s *UpdateGatewayRouteRequestFallbackServices) SetAgreementType(v string) *UpdateGatewayRouteRequestFallbackServices {
	s.AgreementType = &v
	return s
}

func (s *UpdateGatewayRouteRequestFallbackServices) SetGroupName(v string) *UpdateGatewayRouteRequestFallbackServices {
	s.GroupName = &v
	return s
}

func (s *UpdateGatewayRouteRequestFallbackServices) SetName(v string) *UpdateGatewayRouteRequestFallbackServices {
	s.Name = &v
	return s
}

func (s *UpdateGatewayRouteRequestFallbackServices) SetNamespace(v string) *UpdateGatewayRouteRequestFallbackServices {
	s.Namespace = &v
	return s
}

func (s *UpdateGatewayRouteRequestFallbackServices) SetPercent(v int32) *UpdateGatewayRouteRequestFallbackServices {
	s.Percent = &v
	return s
}

func (s *UpdateGatewayRouteRequestFallbackServices) SetServiceId(v int64) *UpdateGatewayRouteRequestFallbackServices {
	s.ServiceId = &v
	return s
}

func (s *UpdateGatewayRouteRequestFallbackServices) SetServicePort(v int32) *UpdateGatewayRouteRequestFallbackServices {
	s.ServicePort = &v
	return s
}

func (s *UpdateGatewayRouteRequestFallbackServices) SetSourceType(v string) *UpdateGatewayRouteRequestFallbackServices {
	s.SourceType = &v
	return s
}

func (s *UpdateGatewayRouteRequestFallbackServices) SetVersion(v string) *UpdateGatewayRouteRequestFallbackServices {
	s.Version = &v
	return s
}

func (s *UpdateGatewayRouteRequestFallbackServices) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteRequestPredicates struct {
	// The information about header matching.
	HeaderPredicates []*UpdateGatewayRouteRequestPredicatesHeaderPredicates `json:"HeaderPredicates,omitempty" xml:"HeaderPredicates,omitempty" type:"Repeated"`
	// The information about method matching.
	MethodPredicates []*string `json:"MethodPredicates,omitempty" xml:"MethodPredicates,omitempty" type:"Repeated"`
	// The information about path matching.
	PathPredicates *UpdateGatewayRouteRequestPredicatesPathPredicates `json:"PathPredicates,omitempty" xml:"PathPredicates,omitempty" type:"Struct"`
	// The information about parameter matching.
	QueryPredicates []*UpdateGatewayRouteRequestPredicatesQueryPredicates `json:"QueryPredicates,omitempty" xml:"QueryPredicates,omitempty" type:"Repeated"`
}

func (s UpdateGatewayRouteRequestPredicates) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRequestPredicates) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequestPredicates) GetHeaderPredicates() []*UpdateGatewayRouteRequestPredicatesHeaderPredicates {
	return s.HeaderPredicates
}

func (s *UpdateGatewayRouteRequestPredicates) GetMethodPredicates() []*string {
	return s.MethodPredicates
}

func (s *UpdateGatewayRouteRequestPredicates) GetPathPredicates() *UpdateGatewayRouteRequestPredicatesPathPredicates {
	return s.PathPredicates
}

func (s *UpdateGatewayRouteRequestPredicates) GetQueryPredicates() []*UpdateGatewayRouteRequestPredicatesQueryPredicates {
	return s.QueryPredicates
}

func (s *UpdateGatewayRouteRequestPredicates) SetHeaderPredicates(v []*UpdateGatewayRouteRequestPredicatesHeaderPredicates) *UpdateGatewayRouteRequestPredicates {
	s.HeaderPredicates = v
	return s
}

func (s *UpdateGatewayRouteRequestPredicates) SetMethodPredicates(v []*string) *UpdateGatewayRouteRequestPredicates {
	s.MethodPredicates = v
	return s
}

func (s *UpdateGatewayRouteRequestPredicates) SetPathPredicates(v *UpdateGatewayRouteRequestPredicatesPathPredicates) *UpdateGatewayRouteRequestPredicates {
	s.PathPredicates = v
	return s
}

func (s *UpdateGatewayRouteRequestPredicates) SetQueryPredicates(v []*UpdateGatewayRouteRequestPredicatesQueryPredicates) *UpdateGatewayRouteRequestPredicates {
	s.QueryPredicates = v
	return s
}

func (s *UpdateGatewayRouteRequestPredicates) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteRequestPredicatesHeaderPredicates struct {
	// The key of the request header.
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
	// The value of the request header.
	//
	// example:
	//
	// on
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateGatewayRouteRequestPredicatesHeaderPredicates) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRequestPredicatesHeaderPredicates) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequestPredicatesHeaderPredicates) GetKey() *string {
	return s.Key
}

func (s *UpdateGatewayRouteRequestPredicatesHeaderPredicates) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayRouteRequestPredicatesHeaderPredicates) GetValue() *string {
	return s.Value
}

func (s *UpdateGatewayRouteRequestPredicatesHeaderPredicates) SetKey(v string) *UpdateGatewayRouteRequestPredicatesHeaderPredicates {
	s.Key = &v
	return s
}

func (s *UpdateGatewayRouteRequestPredicatesHeaderPredicates) SetType(v string) *UpdateGatewayRouteRequestPredicatesHeaderPredicates {
	s.Type = &v
	return s
}

func (s *UpdateGatewayRouteRequestPredicatesHeaderPredicates) SetValue(v string) *UpdateGatewayRouteRequestPredicatesHeaderPredicates {
	s.Value = &v
	return s
}

func (s *UpdateGatewayRouteRequestPredicatesHeaderPredicates) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteRequestPredicatesPathPredicates struct {
	// Specifies whether to perform case-insensitive matching.
	//
	// example:
	//
	// true
	IgnoreCase *bool `json:"IgnoreCase,omitempty" xml:"IgnoreCase,omitempty"`
	// The path used for route matching.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The matching type.
	//
	// example:
	//
	// PRE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateGatewayRouteRequestPredicatesPathPredicates) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRequestPredicatesPathPredicates) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequestPredicatesPathPredicates) GetIgnoreCase() *bool {
	return s.IgnoreCase
}

func (s *UpdateGatewayRouteRequestPredicatesPathPredicates) GetPath() *string {
	return s.Path
}

func (s *UpdateGatewayRouteRequestPredicatesPathPredicates) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayRouteRequestPredicatesPathPredicates) SetIgnoreCase(v bool) *UpdateGatewayRouteRequestPredicatesPathPredicates {
	s.IgnoreCase = &v
	return s
}

func (s *UpdateGatewayRouteRequestPredicatesPathPredicates) SetPath(v string) *UpdateGatewayRouteRequestPredicatesPathPredicates {
	s.Path = &v
	return s
}

func (s *UpdateGatewayRouteRequestPredicatesPathPredicates) SetType(v string) *UpdateGatewayRouteRequestPredicatesPathPredicates {
	s.Type = &v
	return s
}

func (s *UpdateGatewayRouteRequestPredicatesPathPredicates) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteRequestPredicatesQueryPredicates struct {
	// The name of the parameter.
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
	// The value of the parameter.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateGatewayRouteRequestPredicatesQueryPredicates) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRequestPredicatesQueryPredicates) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequestPredicatesQueryPredicates) GetKey() *string {
	return s.Key
}

func (s *UpdateGatewayRouteRequestPredicatesQueryPredicates) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayRouteRequestPredicatesQueryPredicates) GetValue() *string {
	return s.Value
}

func (s *UpdateGatewayRouteRequestPredicatesQueryPredicates) SetKey(v string) *UpdateGatewayRouteRequestPredicatesQueryPredicates {
	s.Key = &v
	return s
}

func (s *UpdateGatewayRouteRequestPredicatesQueryPredicates) SetType(v string) *UpdateGatewayRouteRequestPredicatesQueryPredicates {
	s.Type = &v
	return s
}

func (s *UpdateGatewayRouteRequestPredicatesQueryPredicates) SetValue(v string) *UpdateGatewayRouteRequestPredicatesQueryPredicates {
	s.Value = &v
	return s
}

func (s *UpdateGatewayRouteRequestPredicatesQueryPredicates) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteRequestRedirectJSON struct {
	// The status code returned.
	//
	// example:
	//
	// 302
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The hostname to be redirected to.
	//
	// example:
	//
	// test.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The path to be redirected to.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s UpdateGatewayRouteRequestRedirectJSON) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRequestRedirectJSON) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequestRedirectJSON) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayRouteRequestRedirectJSON) GetHost() *string {
	return s.Host
}

func (s *UpdateGatewayRouteRequestRedirectJSON) GetPath() *string {
	return s.Path
}

func (s *UpdateGatewayRouteRequestRedirectJSON) SetCode(v int32) *UpdateGatewayRouteRequestRedirectJSON {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteRequestRedirectJSON) SetHost(v string) *UpdateGatewayRouteRequestRedirectJSON {
	s.Host = &v
	return s
}

func (s *UpdateGatewayRouteRequestRedirectJSON) SetPath(v string) *UpdateGatewayRouteRequestRedirectJSON {
	s.Path = &v
	return s
}

func (s *UpdateGatewayRouteRequestRedirectJSON) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteRequestServices struct {
	// The type of the protocol. Valid values:
	//
	// example:
	//
	// DUBBO
	AgreementType *string `json:"AgreementType,omitempty" xml:"AgreementType,omitempty"`
	// The name of the group to which the service belongs.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The transcoder of the Dubbo protocol.
	HttpDubboTranscoder *UpdateGatewayRouteRequestServicesHttpDubboTranscoder `json:"HttpDubboTranscoder,omitempty" xml:"HttpDubboTranscoder,omitempty" type:"Struct"`
	// The name.
	//
	// example:
	//
	// web
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace in which the service resides.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The percentage.
	//
	// example:
	//
	// 80
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 1
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The Dubbo port number.
	//
	// example:
	//
	// 20880
	ServicePort *int32 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// The source type.
	//
	// example:
	//
	// MSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateGatewayRouteRequestServices) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRequestServices) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequestServices) GetAgreementType() *string {
	return s.AgreementType
}

func (s *UpdateGatewayRouteRequestServices) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateGatewayRouteRequestServices) GetHttpDubboTranscoder() *UpdateGatewayRouteRequestServicesHttpDubboTranscoder {
	return s.HttpDubboTranscoder
}

func (s *UpdateGatewayRouteRequestServices) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayRouteRequestServices) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateGatewayRouteRequestServices) GetPercent() *int32 {
	return s.Percent
}

func (s *UpdateGatewayRouteRequestServices) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *UpdateGatewayRouteRequestServices) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *UpdateGatewayRouteRequestServices) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateGatewayRouteRequestServices) GetVersion() *string {
	return s.Version
}

func (s *UpdateGatewayRouteRequestServices) SetAgreementType(v string) *UpdateGatewayRouteRequestServices {
	s.AgreementType = &v
	return s
}

func (s *UpdateGatewayRouteRequestServices) SetGroupName(v string) *UpdateGatewayRouteRequestServices {
	s.GroupName = &v
	return s
}

func (s *UpdateGatewayRouteRequestServices) SetHttpDubboTranscoder(v *UpdateGatewayRouteRequestServicesHttpDubboTranscoder) *UpdateGatewayRouteRequestServices {
	s.HttpDubboTranscoder = v
	return s
}

func (s *UpdateGatewayRouteRequestServices) SetName(v string) *UpdateGatewayRouteRequestServices {
	s.Name = &v
	return s
}

func (s *UpdateGatewayRouteRequestServices) SetNamespace(v string) *UpdateGatewayRouteRequestServices {
	s.Namespace = &v
	return s
}

func (s *UpdateGatewayRouteRequestServices) SetPercent(v int32) *UpdateGatewayRouteRequestServices {
	s.Percent = &v
	return s
}

func (s *UpdateGatewayRouteRequestServices) SetServiceId(v int64) *UpdateGatewayRouteRequestServices {
	s.ServiceId = &v
	return s
}

func (s *UpdateGatewayRouteRequestServices) SetServicePort(v int32) *UpdateGatewayRouteRequestServices {
	s.ServicePort = &v
	return s
}

func (s *UpdateGatewayRouteRequestServices) SetSourceType(v string) *UpdateGatewayRouteRequestServices {
	s.SourceType = &v
	return s
}

func (s *UpdateGatewayRouteRequestServices) SetVersion(v string) *UpdateGatewayRouteRequestServices {
	s.Version = &v
	return s
}

func (s *UpdateGatewayRouteRequestServices) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteRequestServicesHttpDubboTranscoder struct {
	// The Dubbo service group.
	//
	// example:
	//
	// service name
	DubboServiceGroup *string `json:"DubboServiceGroup,omitempty" xml:"DubboServiceGroup,omitempty"`
	// The name of the Dubbo service.
	//
	// example:
	//
	// org.apache.dubbo.samples.basic.api.DemoService
	DubboServiceName *string `json:"DubboServiceName,omitempty" xml:"DubboServiceName,omitempty"`
	// The version of the Dubbo service.
	//
	// example:
	//
	// 0.0.0
	DubboServiceVersion *string `json:"DubboServiceVersion,omitempty" xml:"DubboServiceVersion,omitempty"`
	// The forwarding rules of the Dubbo service.
	MothedMapList []*UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList `json:"MothedMapList,omitempty" xml:"MothedMapList,omitempty" type:"Repeated"`
}

func (s UpdateGatewayRouteRequestServicesHttpDubboTranscoder) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRequestServicesHttpDubboTranscoder) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoder) GetDubboServiceGroup() *string {
	return s.DubboServiceGroup
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoder) GetDubboServiceName() *string {
	return s.DubboServiceName
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoder) GetDubboServiceVersion() *string {
	return s.DubboServiceVersion
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoder) GetMothedMapList() []*UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	return s.MothedMapList
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoder) SetDubboServiceGroup(v string) *UpdateGatewayRouteRequestServicesHttpDubboTranscoder {
	s.DubboServiceGroup = &v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoder) SetDubboServiceName(v string) *UpdateGatewayRouteRequestServicesHttpDubboTranscoder {
	s.DubboServiceName = &v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoder) SetDubboServiceVersion(v string) *UpdateGatewayRouteRequestServicesHttpDubboTranscoder {
	s.DubboServiceVersion = &v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoder) SetMothedMapList(v []*UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) *UpdateGatewayRouteRequestServicesHttpDubboTranscoder {
	s.MothedMapList = v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoder) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList struct {
	// The method name of the Dubbo service.
	//
	// example:
	//
	// sayHello
	DubboMothedName *string `json:"DubboMothedName,omitempty" xml:"DubboMothedName,omitempty"`
	// The HTTP method.
	//
	// > Valid values:
	//
	// 	- ALL_GET
	//
	// 	- ALL_POST
	//
	// 	- ALL_PUT
	//
	// 	- ALL_DELETE
	//
	// 	- ALL_PATCH
	//
	// example:
	//
	// ALL_GET
	HttpMothed *string `json:"HttpMothed,omitempty" xml:"HttpMothed,omitempty"`
	// The path that is used to match a method.
	//
	// example:
	//
	// /mytestzbk/sayhello
	Mothedpath *string `json:"Mothedpath,omitempty" xml:"Mothedpath,omitempty"`
	// The information of parameter mappings.
	ParamMapsList []*UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList `json:"ParamMapsList,omitempty" xml:"ParamMapsList,omitempty" type:"Repeated"`
	// The pass-through type of the header.
	//
	// > Valid values:
	//
	// 	- PASS_ALL: All headers are passed through.
	//
	// 	- PASS_NOT: All headers are not passed through.
	//
	// 	- PASS_ASSIGN: Specified headers are passed through.
	//
	// example:
	//
	// PASS_NOT
	PassThroughAllHeaders *string `json:"PassThroughAllHeaders,omitempty" xml:"PassThroughAllHeaders,omitempty"`
	// The list of headers to be passed through.
	PassThroughList []*string `json:"PassThroughList,omitempty" xml:"PassThroughList,omitempty" type:"Repeated"`
}

func (s UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GetDubboMothedName() *string {
	return s.DubboMothedName
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GetHttpMothed() *string {
	return s.HttpMothed
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GetMothedpath() *string {
	return s.Mothedpath
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GetParamMapsList() []*UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList {
	return s.ParamMapsList
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GetPassThroughAllHeaders() *string {
	return s.PassThroughAllHeaders
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) GetPassThroughList() []*string {
	return s.PassThroughList
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) SetDubboMothedName(v string) *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	s.DubboMothedName = &v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) SetHttpMothed(v string) *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	s.HttpMothed = &v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) SetMothedpath(v string) *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	s.Mothedpath = &v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) SetParamMapsList(v []*UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	s.ParamMapsList = v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) SetPassThroughAllHeaders(v string) *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	s.PassThroughAllHeaders = &v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) SetPassThroughList(v []*string) *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList {
	s.PassThroughList = v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapList) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList struct {
	// The key extracted from the input parameter.
	//
	// example:
	//
	// name
	ExtractKey *string `json:"ExtractKey,omitempty" xml:"ExtractKey,omitempty"`
	// The position of the input parameter.
	//
	// > Valid values:
	//
	// 	- `ALL_QUERY_PARAMETER`: request parameter
	//
	// 	- `ALL_HEADER`: request header
	//
	// 	- `ALL_PATH`: request path
	//
	// 	- `ALL_BODY`: request body
	//
	// example:
	//
	// ALL_QUERY_PARAMETER
	ExtractKeySpec *string `json:"ExtractKeySpec,omitempty" xml:"ExtractKeySpec,omitempty"`
	// The type of the backend service parameter.
	//
	// example:
	//
	// java.lang.String
	MappingType *string `json:"MappingType,omitempty" xml:"MappingType,omitempty"`
}

func (s UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) GetExtractKey() *string {
	return s.ExtractKey
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) GetExtractKeySpec() *string {
	return s.ExtractKeySpec
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) GetMappingType() *string {
	return s.MappingType
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) SetExtractKey(v string) *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList {
	s.ExtractKey = &v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) SetExtractKeySpec(v string) *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList {
	s.ExtractKeySpec = &v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) SetMappingType(v string) *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList {
	s.MappingType = &v
	return s
}

func (s *UpdateGatewayRouteRequestServicesHttpDubboTranscoderMothedMapListParamMapsList) Validate() error {
	return dara.Validate(s)
}
