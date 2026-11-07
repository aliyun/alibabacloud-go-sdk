// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayRouteDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetGatewayRouteDetailResponseBody
	GetCode() *int32
	SetData(v *GetGatewayRouteDetailResponseBodyData) *GetGatewayRouteDetailResponseBody
	GetData() *GetGatewayRouteDetailResponseBodyData
	SetHttpStatusCode(v int32) *GetGatewayRouteDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetGatewayRouteDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayRouteDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGatewayRouteDetailResponseBody
	GetSuccess() *bool
}

type GetGatewayRouteDetailResponseBody struct {
	// Status code. Note that 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *GetGatewayRouteDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// FDF2D16C-5D28-5FAA-A56B-30BDE3559880
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request result, with the following values:
	//
	// - `true`: The request is successful.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayRouteDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetGatewayRouteDetailResponseBody) GetData() *GetGatewayRouteDetailResponseBodyData {
	return s.Data
}

func (s *GetGatewayRouteDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGatewayRouteDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayRouteDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayRouteDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGatewayRouteDetailResponseBody) SetCode(v int32) *GetGatewayRouteDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBody) SetData(v *GetGatewayRouteDetailResponseBodyData) *GetGatewayRouteDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayRouteDetailResponseBody) SetHttpStatusCode(v int32) *GetGatewayRouteDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBody) SetMessage(v string) *GetGatewayRouteDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBody) SetRequestId(v string) *GetGatewayRouteDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBody) SetSuccess(v bool) *GetGatewayRouteDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayRouteDetailResponseBodyData struct {
	// Ahas status.
	//
	// example:
	//
	// 1
	AhasStatus *int32 `json:"AhasStatus,omitempty" xml:"AhasStatus,omitempty"`
	// Cross-domain configuration.
	Cors *GetGatewayRouteDetailResponseBodyDataCors `json:"Cors,omitempty" xml:"Cors,omitempty" type:"Struct"`
	// Default service ID.
	//
	// example:
	//
	// 3
	DefaultServiceId *int64 `json:"DefaultServiceId,omitempty" xml:"DefaultServiceId,omitempty"`
	// Default service name.
	//
	// example:
	//
	// test
	DefaultServiceName *string `json:"DefaultServiceName,omitempty" xml:"DefaultServiceName,omitempty"`
	// Route description
	//
	// example:
	//
	// a route for xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Target service type.
	//
	// example:
	//
	// Single
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// Mock information.
	DirectResponse *GetGatewayRouteDetailResponseBodyDataDirectResponse `json:"DirectResponse,omitempty" xml:"DirectResponse,omitempty" type:"Struct"`
	// Domain ID.
	//
	// example:
	//
	// 235
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// Domain list
	DomainIdList []*int64 `json:"DomainIdList,omitempty" xml:"DomainIdList,omitempty" type:"Repeated"`
	// Domain name.
	//
	// example:
	//
	// 123.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Domain name list
	DomainNameList []*string `json:"DomainNameList,omitempty" xml:"DomainNameList,omitempty" type:"Repeated"`
	// Whether to enable Waf.
	//
	// example:
	//
	// true
	EnableWaf *bool `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	// Whether to enable Fallback service.
	//
	// example:
	//
	// true
	Fallback *bool `json:"Fallback,omitempty" xml:"Fallback,omitempty"`
	// Fallback services.
	FallbackServices []*GetGatewayRouteDetailResponseBodyDataFallbackServices `json:"FallbackServices,omitempty" xml:"FallbackServices,omitempty" type:"Repeated"`
	// Traffic mirroring configuration.
	FlowMirror *GetGatewayRouteDetailResponseBodyDataFlowMirror `json:"FlowMirror,omitempty" xml:"FlowMirror,omitempty" type:"Struct"`
	// Gateway ID.
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// Gateway unique identifier ID.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597c****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 2022-01-07 18:07:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Last modification time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Rewrite policy information.
	HTTPRewrite *GetGatewayRouteDetailResponseBodyDataHTTPRewrite `json:"HTTPRewrite,omitempty" xml:"HTTPRewrite,omitempty" type:"Struct"`
	// Header settings information.
	HeaderOp *GetGatewayRouteDetailResponseBodyDataHeaderOp `json:"HeaderOp,omitempty" xml:"HeaderOp,omitempty" type:"Struct"`
	// ID.
	//
	// example:
	//
	// 1050
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Routing policy Json string.
	//
	// example:
	//
	// {
	//
	//       "CORS": "{\\"allowMethods\\":\\"GET,POST,PUT,DELETE,HEAD,OPTIONS,PATCH\\",\\"allowHeaders\\":\\"*\\",\\"exposeHeaders\\":\\"*\\",\\"unitNum\\":12,\\"allowCredentials\\":true,\\"status\\":\\"off\\",\\"allowOrigins\\":\\"*\\",\\"timeUnit\\":\\"h\\"}",
	//
	//       "Timeout": "{\\"unitNum\\":10,\\"timeUnit\\":\\"s\\",\\"status\\":\\"off\\"}",
	//
	//       "Retry": "{\\"attempts\\":2,\\"retryOn\\":[\\"5xx\\"],\\"status\\":\\"off\\"}",
	//
	//       "HTTPRewrite": "{\\"pathType\\":\\"EQUAL\\",\\"path\\":\\"/o\\",\\"status\\":\\"off\\"}",
	//
	//       "Waf": "{\\"enabled\\":false}",
	//
	//       "HeaderOp": "{\\"status\\":\\"off\\",\\"headerOpItems\\":[{\\"directionType\\":\\"Request\\",\\"opType\\":\\"Add\\",\\"key\\":\\"kkk\\",\\"value\\":\\"ll\\"}]}"
	//
	// }
	Policies *string `json:"Policies,omitempty" xml:"Policies,omitempty"`
	// Matching conditions.
	//
	// example:
	//
	// {}
	Predicates *string `json:"Predicates,omitempty" xml:"Predicates,omitempty"`
	// Redirection configuration
	Redirect *GetGatewayRouteDetailResponseBodyDataRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// Retry configuration.
	Retry *GetGatewayRouteDetailResponseBodyDataRetry `json:"Retry,omitempty" xml:"Retry,omitempty" type:"Struct"`
	// Route order.
	//
	// example:
	//
	// 1
	RouteOrder *int32 `json:"RouteOrder,omitempty" xml:"RouteOrder,omitempty"`
	// Route matching.
	RoutePredicates *GetGatewayRouteDetailResponseBodyDataRoutePredicates `json:"RoutePredicates,omitempty" xml:"RoutePredicates,omitempty" type:"Struct"`
	// Services.
	RouteServices []*GetGatewayRouteDetailResponseBodyDataRouteServices `json:"RouteServices,omitempty" xml:"RouteServices,omitempty" type:"Repeated"`
	// Service configuration information.
	//
	// example:
	//
	// [{\\"Percent\\":100,\\"ServiceId\\":126}]
	Services *string `json:"Services,omitempty" xml:"Services,omitempty"`
	// Route status.
	//
	// - Not published: 0
	//
	// - Publishing: 2
	//
	// - Published: 3
	//
	// - Editing (updates not published): 4
	//
	// - Unpublishing: 5
	//
	// - Unavailable: 6
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Timeout configuration.
	Timeout *GetGatewayRouteDetailResponseBodyDataTimeout `json:"Timeout,omitempty" xml:"Timeout,omitempty" type:"Struct"`
}

func (s GetGatewayRouteDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyData) GetAhasStatus() *int32 {
	return s.AhasStatus
}

func (s *GetGatewayRouteDetailResponseBodyData) GetCors() *GetGatewayRouteDetailResponseBodyDataCors {
	return s.Cors
}

func (s *GetGatewayRouteDetailResponseBodyData) GetDefaultServiceId() *int64 {
	return s.DefaultServiceId
}

func (s *GetGatewayRouteDetailResponseBodyData) GetDefaultServiceName() *string {
	return s.DefaultServiceName
}

func (s *GetGatewayRouteDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetGatewayRouteDetailResponseBodyData) GetDestinationType() *string {
	return s.DestinationType
}

func (s *GetGatewayRouteDetailResponseBodyData) GetDirectResponse() *GetGatewayRouteDetailResponseBodyDataDirectResponse {
	return s.DirectResponse
}

func (s *GetGatewayRouteDetailResponseBodyData) GetDomainId() *int64 {
	return s.DomainId
}

func (s *GetGatewayRouteDetailResponseBodyData) GetDomainIdList() []*int64 {
	return s.DomainIdList
}

func (s *GetGatewayRouteDetailResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *GetGatewayRouteDetailResponseBodyData) GetDomainNameList() []*string {
	return s.DomainNameList
}

func (s *GetGatewayRouteDetailResponseBodyData) GetEnableWaf() *bool {
	return s.EnableWaf
}

func (s *GetGatewayRouteDetailResponseBodyData) GetFallback() *bool {
	return s.Fallback
}

func (s *GetGatewayRouteDetailResponseBodyData) GetFallbackServices() []*GetGatewayRouteDetailResponseBodyDataFallbackServices {
	return s.FallbackServices
}

func (s *GetGatewayRouteDetailResponseBodyData) GetFlowMirror() *GetGatewayRouteDetailResponseBodyDataFlowMirror {
	return s.FlowMirror
}

func (s *GetGatewayRouteDetailResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GetGatewayRouteDetailResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayRouteDetailResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGatewayRouteDetailResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGatewayRouteDetailResponseBodyData) GetHTTPRewrite() *GetGatewayRouteDetailResponseBodyDataHTTPRewrite {
	return s.HTTPRewrite
}

func (s *GetGatewayRouteDetailResponseBodyData) GetHeaderOp() *GetGatewayRouteDetailResponseBodyDataHeaderOp {
	return s.HeaderOp
}

func (s *GetGatewayRouteDetailResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayRouteDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetGatewayRouteDetailResponseBodyData) GetPolicies() *string {
	return s.Policies
}

func (s *GetGatewayRouteDetailResponseBodyData) GetPredicates() *string {
	return s.Predicates
}

func (s *GetGatewayRouteDetailResponseBodyData) GetRedirect() *GetGatewayRouteDetailResponseBodyDataRedirect {
	return s.Redirect
}

func (s *GetGatewayRouteDetailResponseBodyData) GetRetry() *GetGatewayRouteDetailResponseBodyDataRetry {
	return s.Retry
}

func (s *GetGatewayRouteDetailResponseBodyData) GetRouteOrder() *int32 {
	return s.RouteOrder
}

func (s *GetGatewayRouteDetailResponseBodyData) GetRoutePredicates() *GetGatewayRouteDetailResponseBodyDataRoutePredicates {
	return s.RoutePredicates
}

func (s *GetGatewayRouteDetailResponseBodyData) GetRouteServices() []*GetGatewayRouteDetailResponseBodyDataRouteServices {
	return s.RouteServices
}

func (s *GetGatewayRouteDetailResponseBodyData) GetServices() *string {
	return s.Services
}

func (s *GetGatewayRouteDetailResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetGatewayRouteDetailResponseBodyData) GetTimeout() *GetGatewayRouteDetailResponseBodyDataTimeout {
	return s.Timeout
}

func (s *GetGatewayRouteDetailResponseBodyData) SetAhasStatus(v int32) *GetGatewayRouteDetailResponseBodyData {
	s.AhasStatus = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetCors(v *GetGatewayRouteDetailResponseBodyDataCors) *GetGatewayRouteDetailResponseBodyData {
	s.Cors = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetDefaultServiceId(v int64) *GetGatewayRouteDetailResponseBodyData {
	s.DefaultServiceId = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetDefaultServiceName(v string) *GetGatewayRouteDetailResponseBodyData {
	s.DefaultServiceName = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetDescription(v string) *GetGatewayRouteDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetDestinationType(v string) *GetGatewayRouteDetailResponseBodyData {
	s.DestinationType = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetDirectResponse(v *GetGatewayRouteDetailResponseBodyDataDirectResponse) *GetGatewayRouteDetailResponseBodyData {
	s.DirectResponse = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetDomainId(v int64) *GetGatewayRouteDetailResponseBodyData {
	s.DomainId = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetDomainIdList(v []*int64) *GetGatewayRouteDetailResponseBodyData {
	s.DomainIdList = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetDomainName(v string) *GetGatewayRouteDetailResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetDomainNameList(v []*string) *GetGatewayRouteDetailResponseBodyData {
	s.DomainNameList = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetEnableWaf(v bool) *GetGatewayRouteDetailResponseBodyData {
	s.EnableWaf = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetFallback(v bool) *GetGatewayRouteDetailResponseBodyData {
	s.Fallback = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetFallbackServices(v []*GetGatewayRouteDetailResponseBodyDataFallbackServices) *GetGatewayRouteDetailResponseBodyData {
	s.FallbackServices = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetFlowMirror(v *GetGatewayRouteDetailResponseBodyDataFlowMirror) *GetGatewayRouteDetailResponseBodyData {
	s.FlowMirror = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetGatewayId(v int64) *GetGatewayRouteDetailResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetGatewayUniqueId(v string) *GetGatewayRouteDetailResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetGmtCreate(v string) *GetGatewayRouteDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetGmtModified(v string) *GetGatewayRouteDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetHTTPRewrite(v *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) *GetGatewayRouteDetailResponseBodyData {
	s.HTTPRewrite = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetHeaderOp(v *GetGatewayRouteDetailResponseBodyDataHeaderOp) *GetGatewayRouteDetailResponseBodyData {
	s.HeaderOp = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetId(v int64) *GetGatewayRouteDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetName(v string) *GetGatewayRouteDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetPolicies(v string) *GetGatewayRouteDetailResponseBodyData {
	s.Policies = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetPredicates(v string) *GetGatewayRouteDetailResponseBodyData {
	s.Predicates = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetRedirect(v *GetGatewayRouteDetailResponseBodyDataRedirect) *GetGatewayRouteDetailResponseBodyData {
	s.Redirect = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetRetry(v *GetGatewayRouteDetailResponseBodyDataRetry) *GetGatewayRouteDetailResponseBodyData {
	s.Retry = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetRouteOrder(v int32) *GetGatewayRouteDetailResponseBodyData {
	s.RouteOrder = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetRoutePredicates(v *GetGatewayRouteDetailResponseBodyDataRoutePredicates) *GetGatewayRouteDetailResponseBodyData {
	s.RoutePredicates = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetRouteServices(v []*GetGatewayRouteDetailResponseBodyDataRouteServices) *GetGatewayRouteDetailResponseBodyData {
	s.RouteServices = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetServices(v string) *GetGatewayRouteDetailResponseBodyData {
	s.Services = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetStatus(v int32) *GetGatewayRouteDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) SetTimeout(v *GetGatewayRouteDetailResponseBodyDataTimeout) *GetGatewayRouteDetailResponseBodyData {
	s.Timeout = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyData) Validate() error {
	if s.Cors != nil {
		if err := s.Cors.Validate(); err != nil {
			return err
		}
	}
	if s.DirectResponse != nil {
		if err := s.DirectResponse.Validate(); err != nil {
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
	if s.FlowMirror != nil {
		if err := s.FlowMirror.Validate(); err != nil {
			return err
		}
	}
	if s.HTTPRewrite != nil {
		if err := s.HTTPRewrite.Validate(); err != nil {
			return err
		}
	}
	if s.HeaderOp != nil {
		if err := s.HeaderOp.Validate(); err != nil {
			return err
		}
	}
	if s.Redirect != nil {
		if err := s.Redirect.Validate(); err != nil {
			return err
		}
	}
	if s.Retry != nil {
		if err := s.Retry.Validate(); err != nil {
			return err
		}
	}
	if s.RoutePredicates != nil {
		if err := s.RoutePredicates.Validate(); err != nil {
			return err
		}
	}
	if s.RouteServices != nil {
		for _, item := range s.RouteServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Timeout != nil {
		if err := s.Timeout.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayRouteDetailResponseBodyDataCors struct {
	// Allow credentials.
	//
	// example:
	//
	// true
	AllowCredentials *bool `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	// Cross-domain Header.
	//
	// example:
	//
	// *
	AllowHeaders *string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty"`
	// Cross-domain methods.
	//
	// example:
	//
	// GET,POST,PUT,DELETE,HEAD,OPTIONS,PATCH
	AllowMethods *string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty"`
	// Cross-domain.
	//
	// example:
	//
	// *
	AllowOrigins *string `json:"AllowOrigins,omitempty" xml:"AllowOrigins,omitempty"`
	// Response header.
	//
	// example:
	//
	// *
	ExposeHeaders *string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty"`
	// Status.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Time unit.
	//
	// example:
	//
	// h
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	// Time number.
	//
	// example:
	//
	// 24
	UnitNum *int64 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataCors) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataCors) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) GetAllowCredentials() *bool {
	return s.AllowCredentials
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) GetAllowHeaders() *string {
	return s.AllowHeaders
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) GetAllowMethods() *string {
	return s.AllowMethods
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) GetAllowOrigins() *string {
	return s.AllowOrigins
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) GetExposeHeaders() *string {
	return s.ExposeHeaders
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) GetStatus() *string {
	return s.Status
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) GetUnitNum() *int64 {
	return s.UnitNum
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) SetAllowCredentials(v bool) *GetGatewayRouteDetailResponseBodyDataCors {
	s.AllowCredentials = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) SetAllowHeaders(v string) *GetGatewayRouteDetailResponseBodyDataCors {
	s.AllowHeaders = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) SetAllowMethods(v string) *GetGatewayRouteDetailResponseBodyDataCors {
	s.AllowMethods = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) SetAllowOrigins(v string) *GetGatewayRouteDetailResponseBodyDataCors {
	s.AllowOrigins = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) SetExposeHeaders(v string) *GetGatewayRouteDetailResponseBodyDataCors {
	s.ExposeHeaders = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) SetStatus(v string) *GetGatewayRouteDetailResponseBodyDataCors {
	s.Status = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) SetTimeUnit(v string) *GetGatewayRouteDetailResponseBodyDataCors {
	s.TimeUnit = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) SetUnitNum(v int64) *GetGatewayRouteDetailResponseBodyDataCors {
	s.UnitNum = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataCors) Validate() error {
	return dara.Validate(s)
}

type GetGatewayRouteDetailResponseBodyDataDirectResponse struct {
	// Mock return value.
	//
	// example:
	//
	// {}
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// Response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataDirectResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataDirectResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataDirectResponse) GetBody() *string {
	return s.Body
}

func (s *GetGatewayRouteDetailResponseBodyDataDirectResponse) GetCode() *int32 {
	return s.Code
}

func (s *GetGatewayRouteDetailResponseBodyDataDirectResponse) SetBody(v string) *GetGatewayRouteDetailResponseBodyDataDirectResponse {
	s.Body = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataDirectResponse) SetCode(v int32) *GetGatewayRouteDetailResponseBodyDataDirectResponse {
	s.Code = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataDirectResponse) Validate() error {
	return dara.Validate(s)
}

type GetGatewayRouteDetailResponseBodyDataFallbackServices struct {
	// Protocol Type.
	//
	// example:
	//
	// DUBBO
	AgreementType *string `json:"AgreementType,omitempty" xml:"AgreementType,omitempty"`
	// Service group name.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Name.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Service namespace.
	//
	// example:
	//
	// namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Weight percentage.
	//
	// example:
	//
	// 80
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// Service ID.
	//
	// example:
	//
	// 1
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Service name.
	//
	// example:
	//
	// name
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// Service port.
	//
	// example:
	//
	// 8848
	ServicePort *int32 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// Source type.
	//
	// example:
	//
	// MSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Service version.
	//
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataFallbackServices) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataFallbackServices) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) GetAgreementType() *string {
	return s.AgreementType
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) GetName() *string {
	return s.Name
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) GetNamespace() *string {
	return s.Namespace
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) GetPercent() *int32 {
	return s.Percent
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) GetSourceType() *string {
	return s.SourceType
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) GetVersion() *string {
	return s.Version
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) SetAgreementType(v string) *GetGatewayRouteDetailResponseBodyDataFallbackServices {
	s.AgreementType = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) SetGroupName(v string) *GetGatewayRouteDetailResponseBodyDataFallbackServices {
	s.GroupName = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) SetName(v string) *GetGatewayRouteDetailResponseBodyDataFallbackServices {
	s.Name = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) SetNamespace(v string) *GetGatewayRouteDetailResponseBodyDataFallbackServices {
	s.Namespace = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) SetPercent(v int32) *GetGatewayRouteDetailResponseBodyDataFallbackServices {
	s.Percent = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) SetServiceId(v int64) *GetGatewayRouteDetailResponseBodyDataFallbackServices {
	s.ServiceId = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) SetServiceName(v string) *GetGatewayRouteDetailResponseBodyDataFallbackServices {
	s.ServiceName = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) SetServicePort(v int32) *GetGatewayRouteDetailResponseBodyDataFallbackServices {
	s.ServicePort = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) SetSourceType(v string) *GetGatewayRouteDetailResponseBodyDataFallbackServices {
	s.SourceType = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) SetVersion(v string) *GetGatewayRouteDetailResponseBodyDataFallbackServices {
	s.Version = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFallbackServices) Validate() error {
	return dara.Validate(s)
}

type GetGatewayRouteDetailResponseBodyDataFlowMirror struct {
	// Traffic replication percentage (%), value range 0-100.
	//
	// example:
	//
	// 90
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// Target service port.
	//
	// example:
	//
	// 8790
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Enabling status, values:
	//
	// - on: enabled
	//
	// - off: disabled
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Target service ID.
	//
	// example:
	//
	// 21
	TargetServiceId *int64 `json:"TargetServiceId,omitempty" xml:"TargetServiceId,omitempty"`
	// Target service name.
	//
	// example:
	//
	// test
	TargetServiceName *string `json:"TargetServiceName,omitempty" xml:"TargetServiceName,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataFlowMirror) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataFlowMirror) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataFlowMirror) GetPercentage() *int32 {
	return s.Percentage
}

func (s *GetGatewayRouteDetailResponseBodyDataFlowMirror) GetPort() *int32 {
	return s.Port
}

func (s *GetGatewayRouteDetailResponseBodyDataFlowMirror) GetStatus() *string {
	return s.Status
}

func (s *GetGatewayRouteDetailResponseBodyDataFlowMirror) GetTargetServiceId() *int64 {
	return s.TargetServiceId
}

func (s *GetGatewayRouteDetailResponseBodyDataFlowMirror) GetTargetServiceName() *string {
	return s.TargetServiceName
}

func (s *GetGatewayRouteDetailResponseBodyDataFlowMirror) SetPercentage(v int32) *GetGatewayRouteDetailResponseBodyDataFlowMirror {
	s.Percentage = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFlowMirror) SetPort(v int32) *GetGatewayRouteDetailResponseBodyDataFlowMirror {
	s.Port = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFlowMirror) SetStatus(v string) *GetGatewayRouteDetailResponseBodyDataFlowMirror {
	s.Status = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFlowMirror) SetTargetServiceId(v int64) *GetGatewayRouteDetailResponseBodyDataFlowMirror {
	s.TargetServiceId = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFlowMirror) SetTargetServiceName(v string) *GetGatewayRouteDetailResponseBodyDataFlowMirror {
	s.TargetServiceName = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataFlowMirror) Validate() error {
	return dara.Validate(s)
}

type GetGatewayRouteDetailResponseBodyDataHTTPRewrite struct {
	// Instance address.
	//
	// example:
	//
	// aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Node path.
	//
	// example:
	//
	// /test/client
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Rewrite type.
	//
	// example:
	//
	// PRE
	PathType *string `json:"PathType,omitempty" xml:"PathType,omitempty"`
	// Matching Pattern.
	//
	// example:
	//
	// /test
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// Enabling status.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Replacement.
	//
	// example:
	//
	// test
	Substitution *string `json:"Substitution,omitempty" xml:"Substitution,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataHTTPRewrite) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataHTTPRewrite) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) GetHost() *string {
	return s.Host
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) GetPath() *string {
	return s.Path
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) GetPathType() *string {
	return s.PathType
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) GetPattern() *string {
	return s.Pattern
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) GetStatus() *string {
	return s.Status
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) GetSubstitution() *string {
	return s.Substitution
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) SetHost(v string) *GetGatewayRouteDetailResponseBodyDataHTTPRewrite {
	s.Host = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) SetPath(v string) *GetGatewayRouteDetailResponseBodyDataHTTPRewrite {
	s.Path = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) SetPathType(v string) *GetGatewayRouteDetailResponseBodyDataHTTPRewrite {
	s.PathType = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) SetPattern(v string) *GetGatewayRouteDetailResponseBodyDataHTTPRewrite {
	s.Pattern = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) SetStatus(v string) *GetGatewayRouteDetailResponseBodyDataHTTPRewrite {
	s.Status = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) SetSubstitution(v string) *GetGatewayRouteDetailResponseBodyDataHTTPRewrite {
	s.Substitution = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataHTTPRewrite) Validate() error {
	return dara.Validate(s)
}

type GetGatewayRouteDetailResponseBodyDataHeaderOp struct {
	// Header field objects.
	HeaderOpItems []*GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems `json:"HeaderOpItems,omitempty" xml:"HeaderOpItems,omitempty" type:"Repeated"`
	// Status.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataHeaderOp) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataHeaderOp) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOp) GetHeaderOpItems() []*GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems {
	return s.HeaderOpItems
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOp) GetStatus() *string {
	return s.Status
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOp) SetHeaderOpItems(v []*GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems) *GetGatewayRouteDetailResponseBodyDataHeaderOp {
	s.HeaderOpItems = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOp) SetStatus(v string) *GetGatewayRouteDetailResponseBodyDataHeaderOp {
	s.Status = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOp) Validate() error {
	if s.HeaderOpItems != nil {
		for _, item := range s.HeaderOpItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems struct {
	// Request or response.
	//
	// example:
	//
	// Response
	DirectionType *string `json:"DirectionType,omitempty" xml:"DirectionType,omitempty"`
	// Header Key.
	//
	// example:
	//
	// debug
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Operation type.
	//
	// example:
	//
	// Add
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	// Setting value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems) GetDirectionType() *string {
	return s.DirectionType
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems) GetKey() *string {
	return s.Key
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems) GetOpType() *string {
	return s.OpType
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems) GetValue() *string {
	return s.Value
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems) SetDirectionType(v string) *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems {
	s.DirectionType = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems) SetKey(v string) *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems {
	s.Key = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems) SetOpType(v string) *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems {
	s.OpType = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems) SetValue(v string) *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems {
	s.Value = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataHeaderOpHeaderOpItems) Validate() error {
	return dara.Validate(s)
}

type GetGatewayRouteDetailResponseBodyDataRedirect struct {
	// Response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Host.
	//
	// example:
	//
	// 16
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Path.
	//
	// example:
	//
	// 10111
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataRedirect) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataRedirect) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataRedirect) GetCode() *int32 {
	return s.Code
}

func (s *GetGatewayRouteDetailResponseBodyDataRedirect) GetHost() *string {
	return s.Host
}

func (s *GetGatewayRouteDetailResponseBodyDataRedirect) GetPath() *string {
	return s.Path
}

func (s *GetGatewayRouteDetailResponseBodyDataRedirect) SetCode(v int32) *GetGatewayRouteDetailResponseBodyDataRedirect {
	s.Code = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRedirect) SetHost(v string) *GetGatewayRouteDetailResponseBodyDataRedirect {
	s.Host = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRedirect) SetPath(v string) *GetGatewayRouteDetailResponseBodyDataRedirect {
	s.Path = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRedirect) Validate() error {
	return dara.Validate(s)
}

type GetGatewayRouteDetailResponseBodyDataRetry struct {
	// Number of retry attempts.
	//
	// example:
	//
	// 1
	Attempts *int32 `json:"Attempts,omitempty" xml:"Attempts,omitempty"`
	// HTTP status codes.
	HttpCodes []*string `json:"HttpCodes,omitempty" xml:"HttpCodes,omitempty" type:"Repeated"`
	// Retry conditions.
	RetryOn []*string `json:"RetryOn,omitempty" xml:"RetryOn,omitempty" type:"Repeated"`
	// Retry status.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataRetry) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataRetry) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataRetry) GetAttempts() *int32 {
	return s.Attempts
}

func (s *GetGatewayRouteDetailResponseBodyDataRetry) GetHttpCodes() []*string {
	return s.HttpCodes
}

func (s *GetGatewayRouteDetailResponseBodyDataRetry) GetRetryOn() []*string {
	return s.RetryOn
}

func (s *GetGatewayRouteDetailResponseBodyDataRetry) GetStatus() *string {
	return s.Status
}

func (s *GetGatewayRouteDetailResponseBodyDataRetry) SetAttempts(v int32) *GetGatewayRouteDetailResponseBodyDataRetry {
	s.Attempts = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRetry) SetHttpCodes(v []*string) *GetGatewayRouteDetailResponseBodyDataRetry {
	s.HttpCodes = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRetry) SetRetryOn(v []*string) *GetGatewayRouteDetailResponseBodyDataRetry {
	s.RetryOn = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRetry) SetStatus(v string) *GetGatewayRouteDetailResponseBodyDataRetry {
	s.Status = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRetry) Validate() error {
	return dara.Validate(s)
}

type GetGatewayRouteDetailResponseBodyDataRoutePredicates struct {
	// Header matching.
	HeaderPredicates []*GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates `json:"HeaderPredicates,omitempty" xml:"HeaderPredicates,omitempty" type:"Repeated"`
	// Method matching.
	MethodPredicates []*string `json:"MethodPredicates,omitempty" xml:"MethodPredicates,omitempty" type:"Repeated"`
	// Route matching.
	PathPredicates *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates `json:"PathPredicates,omitempty" xml:"PathPredicates,omitempty" type:"Struct"`
	// Parameter matching.
	QueryPredicates []*GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates `json:"QueryPredicates,omitempty" xml:"QueryPredicates,omitempty" type:"Repeated"`
}

func (s GetGatewayRouteDetailResponseBodyDataRoutePredicates) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataRoutePredicates) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicates) GetHeaderPredicates() []*GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates {
	return s.HeaderPredicates
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicates) GetMethodPredicates() []*string {
	return s.MethodPredicates
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicates) GetPathPredicates() *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates {
	return s.PathPredicates
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicates) GetQueryPredicates() []*GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates {
	return s.QueryPredicates
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicates) SetHeaderPredicates(v []*GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates) *GetGatewayRouteDetailResponseBodyDataRoutePredicates {
	s.HeaderPredicates = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicates) SetMethodPredicates(v []*string) *GetGatewayRouteDetailResponseBodyDataRoutePredicates {
	s.MethodPredicates = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicates) SetPathPredicates(v *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates) *GetGatewayRouteDetailResponseBodyDataRoutePredicates {
	s.PathPredicates = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicates) SetQueryPredicates(v []*GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates) *GetGatewayRouteDetailResponseBodyDataRoutePredicates {
	s.QueryPredicates = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicates) Validate() error {
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

type GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates struct {
	// Request header KEY.
	//
	// example:
	//
	// id
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Type.
	//
	// example:
	//
	// PRE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Matching value.
	//
	// example:
	//
	// 200
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates) GetKey() *string {
	return s.Key
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates) GetType() *string {
	return s.Type
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates) GetValue() *string {
	return s.Value
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates) SetKey(v string) *GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates {
	s.Key = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates) SetType(v string) *GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates {
	s.Type = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates) SetValue(v string) *GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates {
	s.Value = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesHeaderPredicates) Validate() error {
	return dara.Validate(s)
}

type GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates struct {
	// Whether to ignore case.
	//
	// example:
	//
	// true
	IgnoreCase *bool `json:"IgnoreCase,omitempty" xml:"IgnoreCase,omitempty"`
	// Path.
	//
	// example:
	//
	// /api
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Matching type.
	//
	// example:
	//
	// PRE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates) GetIgnoreCase() *bool {
	return s.IgnoreCase
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates) GetPath() *string {
	return s.Path
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates) GetType() *string {
	return s.Type
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates) SetIgnoreCase(v bool) *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates {
	s.IgnoreCase = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates) SetPath(v string) *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates {
	s.Path = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates) SetType(v string) *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates {
	s.Type = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesPathPredicates) Validate() error {
	return dara.Validate(s)
}

type GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates struct {
	// Parameter name.
	//
	// example:
	//
	// userid
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Type.
	//
	// example:
	//
	// PRE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Matching value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates) GetKey() *string {
	return s.Key
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates) GetType() *string {
	return s.Type
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates) GetValue() *string {
	return s.Value
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates) SetKey(v string) *GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates {
	s.Key = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates) SetType(v string) *GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates {
	s.Type = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates) SetValue(v string) *GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates {
	s.Value = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRoutePredicatesQueryPredicates) Validate() error {
	return dara.Validate(s)
}

type GetGatewayRouteDetailResponseBodyDataRouteServices struct {
	// Protocol Type.
	//
	// example:
	//
	// DUBBO
	AgreementType *string `json:"AgreementType,omitempty" xml:"AgreementType,omitempty"`
	// Service group name.
	//
	// example:
	//
	// api
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Health status
	//
	// - Health
	//
	// - Unhealthy
	//
	// - Checking
	//
	// - Unknown
	//
	// example:
	//
	// Health
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// Dubbo protocol conversion body.
	HttpDubboTranscoder *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder `json:"HttpDubboTranscoder,omitempty" xml:"HttpDubboTranscoder,omitempty" type:"Struct"`
	// Service name.
	//
	// example:
	//
	// dubbo4
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Namespace.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Weight.
	//
	// example:
	//
	// 80
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// Service ID.
	//
	// example:
	//
	// 782
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Service name.
	//
	// example:
	//
	// xkc-crm
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// Service port.
	//
	// example:
	//
	// 20880
	ServicePort *int32 `json:"ServicePort,omitempty" xml:"ServicePort,omitempty"`
	// Service source.
	//
	// example:
	//
	// MSE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Unhealthy Endpoint list
	UnhealthyEndpoints []*string `json:"UnhealthyEndpoints,omitempty" xml:"UnhealthyEndpoints,omitempty" type:"Repeated"`
	// Service version.
	//
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataRouteServices) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataRouteServices) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetAgreementType() *string {
	return s.AgreementType
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetHttpDubboTranscoder() *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder {
	return s.HttpDubboTranscoder
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetName() *string {
	return s.Name
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetNamespace() *string {
	return s.Namespace
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetPercent() *int32 {
	return s.Percent
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetSourceType() *string {
	return s.SourceType
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetUnhealthyEndpoints() []*string {
	return s.UnhealthyEndpoints
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) GetVersion() *string {
	return s.Version
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetAgreementType(v string) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.AgreementType = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetGroupName(v string) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.GroupName = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetHealthStatus(v string) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.HealthStatus = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetHttpDubboTranscoder(v *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.HttpDubboTranscoder = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetName(v string) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.Name = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetNamespace(v string) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.Namespace = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetPercent(v int32) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.Percent = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetServiceId(v int64) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.ServiceId = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetServiceName(v string) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.ServiceName = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetServicePort(v int32) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.ServicePort = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetSourceType(v string) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.SourceType = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetUnhealthyEndpoints(v []*string) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.UnhealthyEndpoints = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) SetVersion(v string) *GetGatewayRouteDetailResponseBodyDataRouteServices {
	s.Version = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServices) Validate() error {
	if s.HttpDubboTranscoder != nil {
		if err := s.HttpDubboTranscoder.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder struct {
	// Dubbo service group.
	//
	// example:
	//
	// service group
	DubboServiceGroup *string `json:"DubboServiceGroup,omitempty" xml:"DubboServiceGroup,omitempty"`
	// Dubbo service name.
	//
	// example:
	//
	// org.apache.dubbo.samples.basic.api.DemoService
	DubboServiceName *string `json:"DubboServiceName,omitempty" xml:"DubboServiceName,omitempty"`
	// Dubbo service version.
	//
	// example:
	//
	// 0.0.0
	DubboServiceVersion *string `json:"DubboServiceVersion,omitempty" xml:"DubboServiceVersion,omitempty"`
	// Dubbo forwarding rule list.
	MothedMapList []*GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList `json:"MothedMapList,omitempty" xml:"MothedMapList,omitempty" type:"Repeated"`
}

func (s GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder) GetDubboServiceGroup() *string {
	return s.DubboServiceGroup
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder) GetDubboServiceName() *string {
	return s.DubboServiceName
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder) GetDubboServiceVersion() *string {
	return s.DubboServiceVersion
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder) GetMothedMapList() []*GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList {
	return s.MothedMapList
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder) SetDubboServiceGroup(v string) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder {
	s.DubboServiceGroup = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder) SetDubboServiceName(v string) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder {
	s.DubboServiceName = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder) SetDubboServiceVersion(v string) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder {
	s.DubboServiceVersion = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder) SetMothedMapList(v []*GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder {
	s.MothedMapList = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoder) Validate() error {
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

type GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList struct {
	// Dubbo method name.
	//
	// example:
	//
	// sayHello
	DubboMothedName *string `json:"DubboMothedName,omitempty" xml:"DubboMothedName,omitempty"`
	// Http method.
	//
	// > Supported http method parameter values.
	//
	// >
	//
	// > - ALL_GET
	//
	// >
	//
	// > - ALL_POST
	//
	// >
	//
	// > - ALL_PUT
	//
	// >
	//
	// > - ALL_DELETE
	//
	// >
	//
	// > - ALL_PATCH
	//
	// example:
	//
	// ALL_GET
	HttpMothed *string `json:"HttpMothed,omitempty" xml:"HttpMothed,omitempty"`
	// Method matching path.
	//
	// example:
	//
	// /mytestzbk/sayhello
	Mothedpath *string `json:"Mothedpath,omitempty" xml:"Mothedpath,omitempty"`
	// Parameter mapping list.
	ParamMapsList []*GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList `json:"ParamMapsList,omitempty" xml:"ParamMapsList,omitempty" type:"Repeated"`
	// Header pass-through type.
	//
	// > Pass-through type enumeration values.
	//
	// >
	//
	// > - Pass through all: PASS_ALL
	//
	// >
	//
	// > - No pass-through: PASS_NOT
	//
	// >
	//
	// > - Specified parameters: PASS_ASSIGN
	//
	// example:
	//
	// PASS_NOT
	PassThroughAllHeaders *string `json:"PassThroughAllHeaders,omitempty" xml:"PassThroughAllHeaders,omitempty"`
	// Specified Header pass-through list.
	PassThroughList []*string `json:"PassThroughList,omitempty" xml:"PassThroughList,omitempty" type:"Repeated"`
}

func (s GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) GetDubboMothedName() *string {
	return s.DubboMothedName
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) GetHttpMothed() *string {
	return s.HttpMothed
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) GetMothedpath() *string {
	return s.Mothedpath
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) GetParamMapsList() []*GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList {
	return s.ParamMapsList
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) GetPassThroughAllHeaders() *string {
	return s.PassThroughAllHeaders
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) GetPassThroughList() []*string {
	return s.PassThroughList
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) SetDubboMothedName(v string) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList {
	s.DubboMothedName = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) SetHttpMothed(v string) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList {
	s.HttpMothed = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) SetMothedpath(v string) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList {
	s.Mothedpath = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) SetParamMapsList(v []*GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList {
	s.ParamMapsList = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) SetPassThroughAllHeaders(v string) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList {
	s.PassThroughAllHeaders = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) SetPassThroughList(v []*string) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList {
	s.PassThroughList = v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapList) Validate() error {
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

type GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList struct {
	// Input parameter extraction Key.
	//
	// example:
	//
	// name
	ExtractKey *string `json:"ExtractKey,omitempty" xml:"ExtractKey,omitempty"`
	// Input parameter position.
	//
	// > Input parameter position enumeration values.
	//
	// >
	//
	// > - Request parameter: `ALL_QUERY_PARAMETER`
	//
	// >
	//
	// > - Request header: `ALL_HEADER`
	//
	// >
	//
	// > - URI of the request: `ALL_PATH`
	//
	// >
	//
	// > - Request body: `ALL_BODY`
	//
	// example:
	//
	// ALL_QUERY_PARAMETER
	ExtractKeySpec *string `json:"ExtractKeySpec,omitempty" xml:"ExtractKeySpec,omitempty"`
	// Backend parameter type.
	//
	// example:
	//
	// java.lang.String
	MappingType *string `json:"MappingType,omitempty" xml:"MappingType,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) GetExtractKey() *string {
	return s.ExtractKey
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) GetExtractKeySpec() *string {
	return s.ExtractKeySpec
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) GetMappingType() *string {
	return s.MappingType
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) SetExtractKey(v string) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList {
	s.ExtractKey = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) SetExtractKeySpec(v string) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList {
	s.ExtractKeySpec = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) SetMappingType(v string) *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList {
	s.MappingType = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) Validate() error {
	return dara.Validate(s)
}

type GetGatewayRouteDetailResponseBodyDataTimeout struct {
	// Status.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Time unit.
	//
	// example:
	//
	// s
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	// Unit number.
	//
	// example:
	//
	// 10
	UnitNum *int32 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s GetGatewayRouteDetailResponseBodyDataTimeout) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayRouteDetailResponseBodyDataTimeout) GoString() string {
	return s.String()
}

func (s *GetGatewayRouteDetailResponseBodyDataTimeout) GetStatus() *string {
	return s.Status
}

func (s *GetGatewayRouteDetailResponseBodyDataTimeout) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *GetGatewayRouteDetailResponseBodyDataTimeout) GetUnitNum() *int32 {
	return s.UnitNum
}

func (s *GetGatewayRouteDetailResponseBodyDataTimeout) SetStatus(v string) *GetGatewayRouteDetailResponseBodyDataTimeout {
	s.Status = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataTimeout) SetTimeUnit(v string) *GetGatewayRouteDetailResponseBodyDataTimeout {
	s.TimeUnit = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataTimeout) SetUnitNum(v int32) *GetGatewayRouteDetailResponseBodyDataTimeout {
	s.UnitNum = &v
	return s
}

func (s *GetGatewayRouteDetailResponseBodyDataTimeout) Validate() error {
	return dara.Validate(s)
}
