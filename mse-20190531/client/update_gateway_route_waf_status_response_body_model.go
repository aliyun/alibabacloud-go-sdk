// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteWafStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateGatewayRouteWafStatusResponseBody
	GetCode() *int32
	SetData(v *UpdateGatewayRouteWafStatusResponseBodyData) *UpdateGatewayRouteWafStatusResponseBody
	GetData() *UpdateGatewayRouteWafStatusResponseBodyData
	SetHttpStatusCode(v int32) *UpdateGatewayRouteWafStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateGatewayRouteWafStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayRouteWafStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGatewayRouteWafStatusResponseBody
	GetSuccess() *bool
}

type UpdateGatewayRouteWafStatusResponseBody struct {
	// The status code returned. A value of 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *UpdateGatewayRouteWafStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7466566F-F30F-4A29-965D-3E0AF21D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayRouteWafStatusResponseBody) GetData() *UpdateGatewayRouteWafStatusResponseBodyData {
	return s.Data
}

func (s *UpdateGatewayRouteWafStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateGatewayRouteWafStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayRouteWafStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayRouteWafStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGatewayRouteWafStatusResponseBody) SetCode(v int32) *UpdateGatewayRouteWafStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBody) SetData(v *UpdateGatewayRouteWafStatusResponseBodyData) *UpdateGatewayRouteWafStatusResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBody) SetHttpStatusCode(v int32) *UpdateGatewayRouteWafStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBody) SetMessage(v string) *UpdateGatewayRouteWafStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBody) SetRequestId(v string) *UpdateGatewayRouteWafStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBody) SetSuccess(v bool) *UpdateGatewayRouteWafStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyData struct {
	// The configuration for cross-origin resource sharing (CORS).
	Cors *UpdateGatewayRouteWafStatusResponseBodyDataCors `json:"Cors,omitempty" xml:"Cors,omitempty" type:"Struct"`
	// The default service ID.
	//
	// example:
	//
	// 1
	DefaultServiceId *int64 `json:"DefaultServiceId,omitempty" xml:"DefaultServiceId,omitempty"`
	// The default service name.
	//
	// example:
	//
	// test
	DefaultServiceName *string `json:"DefaultServiceName,omitempty" xml:"DefaultServiceName,omitempty"`
	// The destination service type.
	//
	// example:
	//
	// Single
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The information about service mocking.
	DirectResponse *UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse `json:"DirectResponse,omitempty" xml:"DirectResponse,omitempty" type:"Struct"`
	// The domain ID.
	//
	// example:
	//
	// 235
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The list of domain IDs.
	DomainIdList []*int64 `json:"DomainIdList,omitempty" xml:"DomainIdList,omitempty" type:"Repeated"`
	// The domain name.
	//
	// example:
	//
	// nbhamster.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The domain names.
	DomainNameList []*string `json:"DomainNameList,omitempty" xml:"DomainNameList,omitempty" type:"Repeated"`
	// Indicates whether WAF is activated.
	//
	// example:
	//
	// 0
	EnableWaf *bool `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 102
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597cd4a9****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-26T09:52:41.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2022-02-24T06:08:29.230+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The information about the rewrite policy.
	HTTPRewrite *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite `json:"HTTPRewrite,omitempty" xml:"HTTPRewrite,omitempty" type:"Struct"`
	// The header settings.
	HeaderOp *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp `json:"HeaderOp,omitempty" xml:"HeaderOp,omitempty" type:"Struct"`
	// The ID of the route.
	//
	// example:
	//
	// 12
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The matching rule.
	//
	// example:
	//
	// {\\"PathPredicates\\":{\\"Path\\":\\"/metas\\",\\"Type\\":\\"PRE\\",\\"IgnoreCase\\":false}}
	Predicates *string `json:"Predicates,omitempty" xml:"Predicates,omitempty"`
	// The configuration of the redirection.
	Redirect *UpdateGatewayRouteWafStatusResponseBodyDataRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The retry configuration.
	Retry *UpdateGatewayRouteWafStatusResponseBodyDataRetry `json:"Retry,omitempty" xml:"Retry,omitempty" type:"Struct"`
	// The sequence number of the route.
	//
	// example:
	//
	// 1
	RouteOrder *int32 `json:"RouteOrder,omitempty" xml:"RouteOrder,omitempty"`
	// The information about route matching.
	RoutePredicates *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates `json:"RoutePredicates,omitempty" xml:"RoutePredicates,omitempty" type:"Struct"`
	// The information about services.
	RouteServices []*UpdateGatewayRouteWafStatusResponseBodyDataRouteServices `json:"RouteServices,omitempty" xml:"RouteServices,omitempty" type:"Repeated"`
	// The information about services.
	//
	// example:
	//
	// [{\\"Percent\\":100,\\"ServiceId\\":126}]
	Services *string `json:"Services,omitempty" xml:"Services,omitempty"`
	// The status of the route.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The timeout configuration.
	Timeout *UpdateGatewayRouteWafStatusResponseBodyDataTimeout `json:"Timeout,omitempty" xml:"Timeout,omitempty" type:"Struct"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetCors() *UpdateGatewayRouteWafStatusResponseBodyDataCors {
	return s.Cors
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetDefaultServiceId() *int64 {
	return s.DefaultServiceId
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetDefaultServiceName() *string {
	return s.DefaultServiceName
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetDestinationType() *string {
	return s.DestinationType
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetDirectResponse() *UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse {
	return s.DirectResponse
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetDomainId() *int64 {
	return s.DomainId
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetDomainIdList() []*int64 {
	return s.DomainIdList
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetDomainNameList() []*string {
	return s.DomainNameList
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetEnableWaf() *bool {
	return s.EnableWaf
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetHTTPRewrite() *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite {
	return s.HTTPRewrite
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetHeaderOp() *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp {
	return s.HeaderOp
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetPredicates() *string {
	return s.Predicates
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetRedirect() *UpdateGatewayRouteWafStatusResponseBodyDataRedirect {
	return s.Redirect
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetRetry() *UpdateGatewayRouteWafStatusResponseBodyDataRetry {
	return s.Retry
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetRouteOrder() *int32 {
	return s.RouteOrder
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetRoutePredicates() *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates {
	return s.RoutePredicates
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetRouteServices() []*UpdateGatewayRouteWafStatusResponseBodyDataRouteServices {
	return s.RouteServices
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetServices() *string {
	return s.Services
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) GetTimeout() *UpdateGatewayRouteWafStatusResponseBodyDataTimeout {
	return s.Timeout
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetCors(v *UpdateGatewayRouteWafStatusResponseBodyDataCors) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.Cors = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetDefaultServiceId(v int64) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.DefaultServiceId = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetDefaultServiceName(v string) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.DefaultServiceName = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetDestinationType(v string) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.DestinationType = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetDirectResponse(v *UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.DirectResponse = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetDomainId(v int64) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.DomainId = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetDomainIdList(v []*int64) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.DomainIdList = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetDomainName(v string) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetDomainNameList(v []*string) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.DomainNameList = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetEnableWaf(v bool) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.EnableWaf = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetGatewayId(v int64) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetGatewayUniqueId(v string) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetGmtCreate(v string) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetGmtModified(v string) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetHTTPRewrite(v *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.HTTPRewrite = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetHeaderOp(v *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.HeaderOp = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetId(v int64) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetName(v string) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetPredicates(v string) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.Predicates = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetRedirect(v *UpdateGatewayRouteWafStatusResponseBodyDataRedirect) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.Redirect = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetRetry(v *UpdateGatewayRouteWafStatusResponseBodyDataRetry) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.Retry = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetRouteOrder(v int32) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.RouteOrder = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetRoutePredicates(v *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.RoutePredicates = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetRouteServices(v []*UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.RouteServices = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetServices(v string) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.Services = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetStatus(v int32) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) SetTimeout(v *UpdateGatewayRouteWafStatusResponseBodyDataTimeout) *UpdateGatewayRouteWafStatusResponseBodyData {
	s.Timeout = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataCors struct {
	// The credentials allowed.
	//
	// example:
	//
	// true
	AllowCredentials *bool `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	// The headers allowed.
	//
	// example:
	//
	// *
	AllowHeaders *string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty"`
	// The methods allowed.
	//
	// example:
	//
	// GET,POST,PUT,DELETE,HEAD,OPTIONS,PATCH
	AllowMethods *string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty"`
	// The origins allowed.
	//
	// example:
	//
	// *
	AllowOrigins *string `json:"AllowOrigins,omitempty" xml:"AllowOrigins,omitempty"`
	// The response headers.
	//
	// example:
	//
	// *
	ExposeHeaders *string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty"`
	// The status.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time unit.
	//
	// example:
	//
	// s
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	// The unit number.
	//
	// example:
	//
	// 1
	UnitNum *int64 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataCors) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataCors) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) GetAllowCredentials() *bool {
	return s.AllowCredentials
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) GetAllowHeaders() *string {
	return s.AllowHeaders
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) GetAllowMethods() *string {
	return s.AllowMethods
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) GetAllowOrigins() *string {
	return s.AllowOrigins
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) GetExposeHeaders() *string {
	return s.ExposeHeaders
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) GetStatus() *string {
	return s.Status
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) GetUnitNum() *int64 {
	return s.UnitNum
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) SetAllowCredentials(v bool) *UpdateGatewayRouteWafStatusResponseBodyDataCors {
	s.AllowCredentials = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) SetAllowHeaders(v string) *UpdateGatewayRouteWafStatusResponseBodyDataCors {
	s.AllowHeaders = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) SetAllowMethods(v string) *UpdateGatewayRouteWafStatusResponseBodyDataCors {
	s.AllowMethods = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) SetAllowOrigins(v string) *UpdateGatewayRouteWafStatusResponseBodyDataCors {
	s.AllowOrigins = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) SetExposeHeaders(v string) *UpdateGatewayRouteWafStatusResponseBodyDataCors {
	s.ExposeHeaders = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) SetStatus(v string) *UpdateGatewayRouteWafStatusResponseBodyDataCors {
	s.Status = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) SetTimeUnit(v string) *UpdateGatewayRouteWafStatusResponseBodyDataCors {
	s.TimeUnit = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) SetUnitNum(v int64) *UpdateGatewayRouteWafStatusResponseBodyDataCors {
	s.UnitNum = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataCors) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse struct {
	// The mock return value.
	//
	// example:
	//
	// [{\\"key\\":\\"h68d13466.sqa.eu95\\",\\"dims\\":\\"groupName=All}]
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The return value.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse) GetBody() *string {
	return s.Body
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse) SetBody(v string) *UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse {
	s.Body = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse) SetCode(v int32) *UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataDirectResponse) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite struct {
	// The domain name.
	//
	// example:
	//
	// aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The HTTP request path.
	//
	// example:
	//
	// /test/client
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The path type of the HTTP request.
	//
	// example:
	//
	// PRE
	PathType *string `json:"PathType,omitempty" xml:"PathType,omitempty"`
	// The matching pattern.
	//
	// example:
	//
	// /test
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// The HTTP status.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The replacement.
	//
	// example:
	//
	// test
	Substitution *string `json:"Substitution,omitempty" xml:"Substitution,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) GetHost() *string {
	return s.Host
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) GetPath() *string {
	return s.Path
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) GetPathType() *string {
	return s.PathType
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) GetPattern() *string {
	return s.Pattern
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) GetStatus() *string {
	return s.Status
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) GetSubstitution() *string {
	return s.Substitution
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) SetHost(v string) *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite {
	s.Host = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) SetPath(v string) *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite {
	s.Path = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) SetPathType(v string) *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite {
	s.PathType = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) SetPattern(v string) *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite {
	s.Pattern = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) SetStatus(v string) *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite {
	s.Status = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) SetSubstitution(v string) *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite {
	s.Substitution = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHTTPRewrite) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp struct {
	// The policy.
	HeaderOpItems []*UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems `json:"HeaderOpItems,omitempty" xml:"HeaderOpItems,omitempty" type:"Repeated"`
	// The status.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp) GetHeaderOpItems() []*UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems {
	return s.HeaderOpItems
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp) GetStatus() *string {
	return s.Status
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp) SetHeaderOpItems(v []*UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems) *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp {
	s.HeaderOpItems = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp) SetStatus(v string) *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp {
	s.Status = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOp) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems struct {
	// The request or response.
	//
	// example:
	//
	// Response
	DirectionType *string `json:"DirectionType,omitempty" xml:"DirectionType,omitempty"`
	// The header key.
	//
	// example:
	//
	// debug
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The operation type.
	//
	// example:
	//
	// Add
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	// The header value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems) GetDirectionType() *string {
	return s.DirectionType
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems) GetKey() *string {
	return s.Key
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems) GetOpType() *string {
	return s.OpType
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems) GetValue() *string {
	return s.Value
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems) SetDirectionType(v string) *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems {
	s.DirectionType = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems) SetKey(v string) *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems {
	s.Key = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems) SetOpType(v string) *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems {
	s.OpType = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems) SetValue(v string) *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems {
	s.Value = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataHeaderOpHeaderOpItems) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataRedirect struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The hostname.
	//
	// example:
	//
	// ww.al.c
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The path.
	//
	// example:
	//
	// /
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRedirect) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRedirect) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRedirect) GetCode() *int32 {
	return s.Code
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRedirect) GetHost() *string {
	return s.Host
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRedirect) GetPath() *string {
	return s.Path
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRedirect) SetCode(v int32) *UpdateGatewayRouteWafStatusResponseBodyDataRedirect {
	s.Code = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRedirect) SetHost(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRedirect {
	s.Host = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRedirect) SetPath(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRedirect {
	s.Path = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRedirect) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataRetry struct {
	// The number of retries allowed for a request.
	//
	// example:
	//
	// 1
	Attempts *int32 `json:"Attempts,omitempty" xml:"Attempts,omitempty"`
	// The HTTP status code.
	HttpCodes []*string `json:"HttpCodes,omitempty" xml:"HttpCodes,omitempty" type:"Repeated"`
	// The retry condition.
	RetryOn []*string `json:"RetryOn,omitempty" xml:"RetryOn,omitempty" type:"Repeated"`
	// The retry status.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRetry) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRetry) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRetry) GetAttempts() *int32 {
	return s.Attempts
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRetry) GetHttpCodes() []*string {
	return s.HttpCodes
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRetry) GetRetryOn() []*string {
	return s.RetryOn
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRetry) GetStatus() *string {
	return s.Status
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRetry) SetAttempts(v int32) *UpdateGatewayRouteWafStatusResponseBodyDataRetry {
	s.Attempts = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRetry) SetHttpCodes(v []*string) *UpdateGatewayRouteWafStatusResponseBodyDataRetry {
	s.HttpCodes = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRetry) SetRetryOn(v []*string) *UpdateGatewayRouteWafStatusResponseBodyDataRetry {
	s.RetryOn = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRetry) SetStatus(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRetry {
	s.Status = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRetry) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates struct {
	// The information about matching based on request headers.
	HeaderPredicates []*UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates `json:"HeaderPredicates,omitempty" xml:"HeaderPredicates,omitempty" type:"Repeated"`
	// The information about method matching.
	MethodPredicates []*string `json:"MethodPredicates,omitempty" xml:"MethodPredicates,omitempty" type:"Repeated"`
	// The information about route matching.
	PathPredicates *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates `json:"PathPredicates,omitempty" xml:"PathPredicates,omitempty" type:"Struct"`
	// The parameter matching rules.
	QueryPredicates []*UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates `json:"QueryPredicates,omitempty" xml:"QueryPredicates,omitempty" type:"Repeated"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates) GetHeaderPredicates() []*UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates {
	return s.HeaderPredicates
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates) GetMethodPredicates() []*string {
	return s.MethodPredicates
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates) GetPathPredicates() *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates {
	return s.PathPredicates
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates) GetQueryPredicates() []*UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates {
	return s.QueryPredicates
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates) SetHeaderPredicates(v []*UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates {
	s.HeaderPredicates = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates) SetMethodPredicates(v []*string) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates {
	s.MethodPredicates = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates) SetPathPredicates(v *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates {
	s.PathPredicates = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates) SetQueryPredicates(v []*UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates {
	s.QueryPredicates = v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicates) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates struct {
	// The key of the request header.
	//
	// example:
	//
	// alibo
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
	// 200
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates) GetKey() *string {
	return s.Key
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates) GetValue() *string {
	return s.Value
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates) SetKey(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates {
	s.Key = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates) SetType(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates {
	s.Type = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates) SetValue(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates {
	s.Value = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesHeaderPredicates) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates struct {
	// Indicates whether case sensitivity is ignored.
	//
	// example:
	//
	// true
	IgnoreCase *bool `json:"IgnoreCase,omitempty" xml:"IgnoreCase,omitempty"`
	// The path.
	//
	// example:
	//
	// /zookeeper/mmgw/unlogined/common.getBu
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The matching type.
	//
	// example:
	//
	// PRE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates) GetIgnoreCase() *bool {
	return s.IgnoreCase
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates) GetPath() *string {
	return s.Path
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates) SetIgnoreCase(v bool) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates {
	s.IgnoreCase = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates) SetPath(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates {
	s.Path = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates) SetType(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates {
	s.Type = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesPathPredicates) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates struct {
	// The name of the parameter.
	//
	// example:
	//
	// instanceId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The type.
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

func (s UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates) GetKey() *string {
	return s.Key
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates) GetValue() *string {
	return s.Value
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates) SetKey(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates {
	s.Key = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates) SetType(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates {
	s.Type = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates) SetValue(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates {
	s.Value = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRoutePredicatesQueryPredicates) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataRouteServices struct {
	// The name of the group to which the service belongs.
	//
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// test-aixue-gray
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The weight.
	//
	// example:
	//
	// 80
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 547
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// b-service
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The source type of the service.
	//
	// example:
	//
	// K8s
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) GetName() *string {
	return s.Name
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) GetPercent() *int32 {
	return s.Percent
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) GetVersion() *string {
	return s.Version
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) SetGroupName(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices {
	s.GroupName = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) SetName(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices {
	s.Name = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) SetNamespace(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices {
	s.Namespace = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) SetPercent(v int32) *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices {
	s.Percent = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) SetServiceId(v int64) *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices {
	s.ServiceId = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) SetServiceName(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices {
	s.ServiceName = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) SetSourceType(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices {
	s.SourceType = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) SetVersion(v string) *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices {
	s.Version = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataRouteServices) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteWafStatusResponseBodyDataTimeout struct {
	// The status.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time unit.
	//
	// example:
	//
	// s
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	// The unit number.
	//
	// example:
	//
	// 10
	UnitNum *int32 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataTimeout) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteWafStatusResponseBodyDataTimeout) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataTimeout) GetStatus() *string {
	return s.Status
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataTimeout) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataTimeout) GetUnitNum() *int32 {
	return s.UnitNum
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataTimeout) SetStatus(v string) *UpdateGatewayRouteWafStatusResponseBodyDataTimeout {
	s.Status = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataTimeout) SetTimeUnit(v string) *UpdateGatewayRouteWafStatusResponseBodyDataTimeout {
	s.TimeUnit = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataTimeout) SetUnitNum(v int32) *UpdateGatewayRouteWafStatusResponseBodyDataTimeout {
	s.UnitNum = &v
	return s
}

func (s *UpdateGatewayRouteWafStatusResponseBodyDataTimeout) Validate() error {
	return dara.Validate(s)
}
