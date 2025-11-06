// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGatewayRouteResponseBody
	GetCode() *int32
	SetData(v *ListGatewayRouteResponseBodyData) *ListGatewayRouteResponseBody
	GetData() *ListGatewayRouteResponseBodyData
	SetHttpStatusCode(v int32) *ListGatewayRouteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGatewayRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGatewayRouteResponseBody
	GetSuccess() *bool
}

type ListGatewayRouteResponseBody struct {
	// The code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListGatewayRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 83F02EAB-ECDF-55C6-B332-8649E5E7AF2C
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

func (s ListGatewayRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGatewayRouteResponseBody) GetData() *ListGatewayRouteResponseBodyData {
	return s.Data
}

func (s *ListGatewayRouteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewayRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewayRouteResponseBody) SetCode(v int32) *ListGatewayRouteResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayRouteResponseBody) SetData(v *ListGatewayRouteResponseBodyData) *ListGatewayRouteResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayRouteResponseBody) SetHttpStatusCode(v int32) *ListGatewayRouteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayRouteResponseBody) SetMessage(v string) *ListGatewayRouteResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayRouteResponseBody) SetRequestId(v string) *ListGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayRouteResponseBody) SetSuccess(v bool) *ListGatewayRouteResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewayRouteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayRouteResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 11
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data structure.
	Result []*ListGatewayRouteResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 36
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListGatewayRouteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayRouteResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayRouteResponseBodyData) GetResult() []*ListGatewayRouteResponseBodyDataResult {
	return s.Result
}

func (s *ListGatewayRouteResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ListGatewayRouteResponseBodyData) SetPageNumber(v int32) *ListGatewayRouteResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayRouteResponseBodyData) SetPageSize(v int32) *ListGatewayRouteResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayRouteResponseBodyData) SetResult(v []*ListGatewayRouteResponseBodyDataResult) *ListGatewayRouteResponseBodyData {
	s.Result = v
	return s
}

func (s *ListGatewayRouteResponseBodyData) SetTotalSize(v int64) *ListGatewayRouteResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGatewayRouteResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayRouteResponseBodyDataResult struct {
	// The route comment (ingress).
	Comment *ListGatewayRouteResponseBodyDataResultComment `json:"Comment,omitempty" xml:"Comment,omitempty" type:"Struct"`
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
	DirectResponse *ListGatewayRouteResponseBodyDataResultDirectResponse `json:"DirectResponse,omitempty" xml:"DirectResponse,omitempty" type:"Struct"`
	// The domain ID.
	//
	// example:
	//
	// 265
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The domain IDs.
	DomainIdList []*int64 `json:"DomainIdList,omitempty" xml:"DomainIdList,omitempty" type:"Repeated"`
	// The domain name.
	//
	// example:
	//
	// y.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The domain names.
	DomainNameList []*string `json:"DomainNameList,omitempty" xml:"DomainNameList,omitempty" type:"Repeated"`
	DynamicRoute   *bool     `json:"DynamicRoute,omitempty" xml:"DynamicRoute,omitempty"`
	// Indicates whether Web Application Firewall (WAF) is activated.
	//
	// example:
	//
	// false
	EnableWaf *string `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	// Indicates whether the Fallback service is enabled.
	//
	// example:
	//
	// true
	Fallback *bool `json:"Fallback,omitempty" xml:"Fallback,omitempty"`
	// The information about the Fallback service.
	FallbackServices []*ListGatewayRouteResponseBodyDataResultFallbackServices `json:"FallbackServices,omitempty" xml:"FallbackServices,omitempty" type:"Repeated"`
	// The ID of the gateway.
	//
	// example:
	//
	// 496
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-cf0e7f52ecc7429dbc7ba4d5e3656100
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-25T07:14:01.817+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2022-01-07 18:07:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID.
	//
	// example:
	//
	// 47
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The matching rules.
	//
	// example:
	//
	// {\\"PathPredicates\\":{\\"Path\\":\\"/rpc/compute \\",\\"Type\\":\\"PRE\\",\\"IgnoreCase\\":true}}
	Predicates *string `json:"Predicates,omitempty" xml:"Predicates,omitempty"`
	// The information about redirection.
	Redirect *ListGatewayRouteResponseBodyDataResultRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The order.
	//
	// example:
	//
	// 1
	RouteOrder *int32 `json:"RouteOrder,omitempty" xml:"RouteOrder,omitempty"`
	// The matching rules.
	RoutePredicates *ListGatewayRouteResponseBodyDataResultRoutePredicates `json:"RoutePredicates,omitempty" xml:"RoutePredicates,omitempty" type:"Struct"`
	// The information about services.
	RouteServices []*ListGatewayRouteResponseBodyDataResultRouteServices `json:"RouteServices,omitempty" xml:"RouteServices,omitempty" type:"Repeated"`
	// The information about services.
	//
	// example:
	//
	// []
	Services *string `json:"Services,omitempty" xml:"Services,omitempty"`
	// The status.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The route type.
	//
	// example:
	//
	// Op
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayRouteResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResult) GetComment() *ListGatewayRouteResponseBodyDataResultComment {
	return s.Comment
}

func (s *ListGatewayRouteResponseBodyDataResult) GetDefaultServiceId() *int64 {
	return s.DefaultServiceId
}

func (s *ListGatewayRouteResponseBodyDataResult) GetDefaultServiceName() *string {
	return s.DefaultServiceName
}

func (s *ListGatewayRouteResponseBodyDataResult) GetDestinationType() *string {
	return s.DestinationType
}

func (s *ListGatewayRouteResponseBodyDataResult) GetDirectResponse() *ListGatewayRouteResponseBodyDataResultDirectResponse {
	return s.DirectResponse
}

func (s *ListGatewayRouteResponseBodyDataResult) GetDomainId() *int64 {
	return s.DomainId
}

func (s *ListGatewayRouteResponseBodyDataResult) GetDomainIdList() []*int64 {
	return s.DomainIdList
}

func (s *ListGatewayRouteResponseBodyDataResult) GetDomainName() *string {
	return s.DomainName
}

func (s *ListGatewayRouteResponseBodyDataResult) GetDomainNameList() []*string {
	return s.DomainNameList
}

func (s *ListGatewayRouteResponseBodyDataResult) GetDynamicRoute() *bool {
	return s.DynamicRoute
}

func (s *ListGatewayRouteResponseBodyDataResult) GetEnableWaf() *string {
	return s.EnableWaf
}

func (s *ListGatewayRouteResponseBodyDataResult) GetFallback() *bool {
	return s.Fallback
}

func (s *ListGatewayRouteResponseBodyDataResult) GetFallbackServices() []*ListGatewayRouteResponseBodyDataResultFallbackServices {
	return s.FallbackServices
}

func (s *ListGatewayRouteResponseBodyDataResult) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *ListGatewayRouteResponseBodyDataResult) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayRouteResponseBodyDataResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListGatewayRouteResponseBodyDataResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListGatewayRouteResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayRouteResponseBodyDataResult) GetName() *string {
	return s.Name
}

func (s *ListGatewayRouteResponseBodyDataResult) GetPredicates() *string {
	return s.Predicates
}

func (s *ListGatewayRouteResponseBodyDataResult) GetRedirect() *ListGatewayRouteResponseBodyDataResultRedirect {
	return s.Redirect
}

func (s *ListGatewayRouteResponseBodyDataResult) GetRouteOrder() *int32 {
	return s.RouteOrder
}

func (s *ListGatewayRouteResponseBodyDataResult) GetRoutePredicates() *ListGatewayRouteResponseBodyDataResultRoutePredicates {
	return s.RoutePredicates
}

func (s *ListGatewayRouteResponseBodyDataResult) GetRouteServices() []*ListGatewayRouteResponseBodyDataResultRouteServices {
	return s.RouteServices
}

func (s *ListGatewayRouteResponseBodyDataResult) GetServices() *string {
	return s.Services
}

func (s *ListGatewayRouteResponseBodyDataResult) GetStatus() *int32 {
	return s.Status
}

func (s *ListGatewayRouteResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *ListGatewayRouteResponseBodyDataResult) SetComment(v *ListGatewayRouteResponseBodyDataResultComment) *ListGatewayRouteResponseBodyDataResult {
	s.Comment = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetDefaultServiceId(v int64) *ListGatewayRouteResponseBodyDataResult {
	s.DefaultServiceId = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetDefaultServiceName(v string) *ListGatewayRouteResponseBodyDataResult {
	s.DefaultServiceName = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetDestinationType(v string) *ListGatewayRouteResponseBodyDataResult {
	s.DestinationType = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetDirectResponse(v *ListGatewayRouteResponseBodyDataResultDirectResponse) *ListGatewayRouteResponseBodyDataResult {
	s.DirectResponse = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetDomainId(v int64) *ListGatewayRouteResponseBodyDataResult {
	s.DomainId = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetDomainIdList(v []*int64) *ListGatewayRouteResponseBodyDataResult {
	s.DomainIdList = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetDomainName(v string) *ListGatewayRouteResponseBodyDataResult {
	s.DomainName = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetDomainNameList(v []*string) *ListGatewayRouteResponseBodyDataResult {
	s.DomainNameList = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetDynamicRoute(v bool) *ListGatewayRouteResponseBodyDataResult {
	s.DynamicRoute = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetEnableWaf(v string) *ListGatewayRouteResponseBodyDataResult {
	s.EnableWaf = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetFallback(v bool) *ListGatewayRouteResponseBodyDataResult {
	s.Fallback = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetFallbackServices(v []*ListGatewayRouteResponseBodyDataResultFallbackServices) *ListGatewayRouteResponseBodyDataResult {
	s.FallbackServices = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetGatewayId(v int64) *ListGatewayRouteResponseBodyDataResult {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetGatewayUniqueId(v string) *ListGatewayRouteResponseBodyDataResult {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetGmtCreate(v string) *ListGatewayRouteResponseBodyDataResult {
	s.GmtCreate = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetGmtModified(v string) *ListGatewayRouteResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetId(v int64) *ListGatewayRouteResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetName(v string) *ListGatewayRouteResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetPredicates(v string) *ListGatewayRouteResponseBodyDataResult {
	s.Predicates = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetRedirect(v *ListGatewayRouteResponseBodyDataResultRedirect) *ListGatewayRouteResponseBodyDataResult {
	s.Redirect = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetRouteOrder(v int32) *ListGatewayRouteResponseBodyDataResult {
	s.RouteOrder = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetRoutePredicates(v *ListGatewayRouteResponseBodyDataResultRoutePredicates) *ListGatewayRouteResponseBodyDataResult {
	s.RoutePredicates = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetRouteServices(v []*ListGatewayRouteResponseBodyDataResultRouteServices) *ListGatewayRouteResponseBodyDataResult {
	s.RouteServices = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetServices(v string) *ListGatewayRouteResponseBodyDataResult {
	s.Services = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetStatus(v int32) *ListGatewayRouteResponseBodyDataResult {
	s.Status = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) SetType(v string) *ListGatewayRouteResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResult) Validate() error {
	if s.Comment != nil {
		if err := s.Comment.Validate(); err != nil {
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
	if s.Redirect != nil {
		if err := s.Redirect.Validate(); err != nil {
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
	return nil
}

type ListGatewayRouteResponseBodyDataResultComment struct {
	// The status.
	//
	// example:
	//
	// error
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListGatewayRouteResponseBodyDataResultComment) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResultComment) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResultComment) GetStatus() *string {
	return s.Status
}

func (s *ListGatewayRouteResponseBodyDataResultComment) SetStatus(v string) *ListGatewayRouteResponseBodyDataResultComment {
	s.Status = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultComment) Validate() error {
	return dara.Validate(s)
}

type ListGatewayRouteResponseBodyDataResultDirectResponse struct {
	// The return value for service mocking.
	//
	// example:
	//
	// hello
	Body *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListGatewayRouteResponseBodyDataResultDirectResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResultDirectResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResultDirectResponse) GetBody() *string {
	return s.Body
}

func (s *ListGatewayRouteResponseBodyDataResultDirectResponse) GetCode() *int32 {
	return s.Code
}

func (s *ListGatewayRouteResponseBodyDataResultDirectResponse) SetBody(v string) *ListGatewayRouteResponseBodyDataResultDirectResponse {
	s.Body = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultDirectResponse) SetCode(v int32) *ListGatewayRouteResponseBodyDataResultDirectResponse {
	s.Code = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultDirectResponse) Validate() error {
	return dara.Validate(s)
}

type ListGatewayRouteResponseBodyDataResultFallbackServices struct {
	// The type of the protocol.
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
	// The namespace to which the service belongs.
	//
	// example:
	//
	// Namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The weight in the form of a percentage value.
	//
	// example:
	//
	// 100
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 353
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// service name
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
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

func (s ListGatewayRouteResponseBodyDataResultFallbackServices) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResultFallbackServices) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) GetAgreementType() *string {
	return s.AgreementType
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) GetName() *string {
	return s.Name
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) GetNamespace() *string {
	return s.Namespace
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) GetPercent() *int32 {
	return s.Percent
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) GetSourceType() *string {
	return s.SourceType
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) GetVersion() *string {
	return s.Version
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) SetAgreementType(v string) *ListGatewayRouteResponseBodyDataResultFallbackServices {
	s.AgreementType = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) SetGroupName(v string) *ListGatewayRouteResponseBodyDataResultFallbackServices {
	s.GroupName = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) SetName(v string) *ListGatewayRouteResponseBodyDataResultFallbackServices {
	s.Name = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) SetNamespace(v string) *ListGatewayRouteResponseBodyDataResultFallbackServices {
	s.Namespace = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) SetPercent(v int32) *ListGatewayRouteResponseBodyDataResultFallbackServices {
	s.Percent = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) SetServiceId(v int64) *ListGatewayRouteResponseBodyDataResultFallbackServices {
	s.ServiceId = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) SetServiceName(v string) *ListGatewayRouteResponseBodyDataResultFallbackServices {
	s.ServiceName = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) SetServicePort(v int32) *ListGatewayRouteResponseBodyDataResultFallbackServices {
	s.ServicePort = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) SetSourceType(v string) *ListGatewayRouteResponseBodyDataResultFallbackServices {
	s.SourceType = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) SetVersion(v string) *ListGatewayRouteResponseBodyDataResultFallbackServices {
	s.Version = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultFallbackServices) Validate() error {
	return dara.Validate(s)
}

type ListGatewayRouteResponseBodyDataResultRedirect struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The hostname to be redirected to.
	//
	// example:
	//
	// test.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The path.
	//
	// example:
	//
	// /test
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ListGatewayRouteResponseBodyDataResultRedirect) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResultRedirect) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResultRedirect) GetCode() *int32 {
	return s.Code
}

func (s *ListGatewayRouteResponseBodyDataResultRedirect) GetHost() *string {
	return s.Host
}

func (s *ListGatewayRouteResponseBodyDataResultRedirect) GetPath() *string {
	return s.Path
}

func (s *ListGatewayRouteResponseBodyDataResultRedirect) SetCode(v int32) *ListGatewayRouteResponseBodyDataResultRedirect {
	s.Code = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRedirect) SetHost(v string) *ListGatewayRouteResponseBodyDataResultRedirect {
	s.Host = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRedirect) SetPath(v string) *ListGatewayRouteResponseBodyDataResultRedirect {
	s.Path = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRedirect) Validate() error {
	return dara.Validate(s)
}

type ListGatewayRouteResponseBodyDataResultRoutePredicates struct {
	// The headers used for route matching.
	HeaderPredicates []*ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates `json:"HeaderPredicates,omitempty" xml:"HeaderPredicates,omitempty" type:"Repeated"`
	// The HTTP methods used for route matching.
	MethodPredicates []*string `json:"MethodPredicates,omitempty" xml:"MethodPredicates,omitempty" type:"Repeated"`
	// The path used for route matching.
	PathPredicates *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates `json:"PathPredicates,omitempty" xml:"PathPredicates,omitempty" type:"Struct"`
	// The parameters used for route matching.
	QueryPredicates []*ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates `json:"QueryPredicates,omitempty" xml:"QueryPredicates,omitempty" type:"Repeated"`
}

func (s ListGatewayRouteResponseBodyDataResultRoutePredicates) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResultRoutePredicates) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicates) GetHeaderPredicates() []*ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates {
	return s.HeaderPredicates
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicates) GetMethodPredicates() []*string {
	return s.MethodPredicates
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicates) GetPathPredicates() *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates {
	return s.PathPredicates
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicates) GetQueryPredicates() []*ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates {
	return s.QueryPredicates
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicates) SetHeaderPredicates(v []*ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates) *ListGatewayRouteResponseBodyDataResultRoutePredicates {
	s.HeaderPredicates = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicates) SetMethodPredicates(v []*string) *ListGatewayRouteResponseBodyDataResultRoutePredicates {
	s.MethodPredicates = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicates) SetPathPredicates(v *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates) *ListGatewayRouteResponseBodyDataResultRoutePredicates {
	s.PathPredicates = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicates) SetQueryPredicates(v []*ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates) *ListGatewayRouteResponseBodyDataResultRoutePredicates {
	s.QueryPredicates = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicates) Validate() error {
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

type ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates struct {
	// The header key.
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
	// The value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates) GetKey() *string {
	return s.Key
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates) GetType() *string {
	return s.Type
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates) GetValue() *string {
	return s.Value
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates) SetKey(v string) *ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates {
	s.Key = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates) SetType(v string) *ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates {
	s.Type = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates) SetValue(v string) *ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates {
	s.Value = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesHeaderPredicates) Validate() error {
	return dara.Validate(s)
}

type ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates struct {
	// Indicates whether case sensitivity is ignored.
	//
	// example:
	//
	// true
	IgnoreCase *bool `json:"IgnoreCase,omitempty" xml:"IgnoreCase,omitempty"`
	// The path of the node.
	//
	// example:
	//
	// /getIp
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The matching type.
	//
	// example:
	//
	// PRE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates) GetIgnoreCase() *bool {
	return s.IgnoreCase
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates) GetPath() *string {
	return s.Path
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates) GetType() *string {
	return s.Type
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates) SetIgnoreCase(v bool) *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates {
	s.IgnoreCase = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates) SetPath(v string) *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates {
	s.Path = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates) SetType(v string) *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates {
	s.Type = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesPathPredicates) Validate() error {
	return dara.Validate(s)
}

type ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates struct {
	// The key.
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
	// The value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates) GetKey() *string {
	return s.Key
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates) GetType() *string {
	return s.Type
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates) GetValue() *string {
	return s.Value
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates) SetKey(v string) *ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates {
	s.Key = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates) SetType(v string) *ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates {
	s.Type = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates) SetValue(v string) *ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates {
	s.Value = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRoutePredicatesQueryPredicates) Validate() error {
	return dara.Validate(s)
}

type ListGatewayRouteResponseBodyDataResultRouteServices struct {
	// The type of the protocol.
	//
	// example:
	//
	// DUBBO
	AgreementType *string `json:"AgreementType,omitempty" xml:"AgreementType,omitempty"`
	// The name of the group to which the service belongs.
	//
	// example:
	//
	// api
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// Health
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The transcoder of the Dubbo protocol.
	HttpDubboTranscoder *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder `json:"HttpDubboTranscoder,omitempty" xml:"HttpDubboTranscoder,omitempty" type:"Struct"`
	// The name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The weight in the form of a percentage value.
	//
	// example:
	//
	// 11
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 1563
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// hu
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
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
	SourceType         *string   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	UnhealthyEndpoints []*string `json:"UnhealthyEndpoints,omitempty" xml:"UnhealthyEndpoints,omitempty" type:"Repeated"`
	// The version of the service.
	//
	// example:
	//
	// v1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListGatewayRouteResponseBodyDataResultRouteServices) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResultRouteServices) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetAgreementType() *string {
	return s.AgreementType
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetHttpDubboTranscoder() *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder {
	return s.HttpDubboTranscoder
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetName() *string {
	return s.Name
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetNamespace() *string {
	return s.Namespace
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetPercent() *int32 {
	return s.Percent
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetServicePort() *int32 {
	return s.ServicePort
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetSourceType() *string {
	return s.SourceType
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetUnhealthyEndpoints() []*string {
	return s.UnhealthyEndpoints
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) GetVersion() *string {
	return s.Version
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetAgreementType(v string) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.AgreementType = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetGroupName(v string) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.GroupName = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetHealthStatus(v string) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.HealthStatus = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetHttpDubboTranscoder(v *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.HttpDubboTranscoder = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetName(v string) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.Name = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetNamespace(v string) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.Namespace = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetPercent(v int32) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.Percent = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetServiceId(v int64) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.ServiceId = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetServiceName(v string) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.ServiceName = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetServicePort(v int32) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.ServicePort = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetSourceType(v string) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.SourceType = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetUnhealthyEndpoints(v []*string) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.UnhealthyEndpoints = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) SetVersion(v string) *ListGatewayRouteResponseBodyDataResultRouteServices {
	s.Version = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServices) Validate() error {
	if s.HttpDubboTranscoder != nil {
		if err := s.HttpDubboTranscoder.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder struct {
	// The Dubbo service group.
	//
	// example:
	//
	// service group
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
	MothedMapList []*ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList `json:"MothedMapList,omitempty" xml:"MothedMapList,omitempty" type:"Repeated"`
}

func (s ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder) GetDubboServiceGroup() *string {
	return s.DubboServiceGroup
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder) GetDubboServiceName() *string {
	return s.DubboServiceName
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder) GetDubboServiceVersion() *string {
	return s.DubboServiceVersion
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder) GetMothedMapList() []*ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList {
	return s.MothedMapList
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder) SetDubboServiceGroup(v string) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder {
	s.DubboServiceGroup = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder) SetDubboServiceName(v string) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder {
	s.DubboServiceName = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder) SetDubboServiceVersion(v string) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder {
	s.DubboServiceVersion = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder) SetMothedMapList(v []*ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder {
	s.MothedMapList = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoder) Validate() error {
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

type ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList struct {
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
	// The path used for method matching.
	//
	// example:
	//
	// /mytestzbk/sayhello
	Mothedpath *string `json:"Mothedpath,omitempty" xml:"Mothedpath,omitempty"`
	// The information about parameter mappings.
	ParamMapsList []*ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList `json:"ParamMapsList,omitempty" xml:"ParamMapsList,omitempty" type:"Repeated"`
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

func (s ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) GetDubboMothedName() *string {
	return s.DubboMothedName
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) GetHttpMothed() *string {
	return s.HttpMothed
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) GetMothedpath() *string {
	return s.Mothedpath
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) GetParamMapsList() []*ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList {
	return s.ParamMapsList
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) GetPassThroughAllHeaders() *string {
	return s.PassThroughAllHeaders
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) GetPassThroughList() []*string {
	return s.PassThroughList
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) SetDubboMothedName(v string) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList {
	s.DubboMothedName = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) SetHttpMothed(v string) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList {
	s.HttpMothed = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) SetMothedpath(v string) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList {
	s.Mothedpath = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) SetParamMapsList(v []*ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList {
	s.ParamMapsList = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) SetPassThroughAllHeaders(v string) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList {
	s.PassThroughAllHeaders = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) SetPassThroughList(v []*string) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList {
	s.PassThroughList = v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapList) Validate() error {
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

type ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList struct {
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

func (s ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) GetExtractKey() *string {
	return s.ExtractKey
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) GetExtractKeySpec() *string {
	return s.ExtractKeySpec
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) GetMappingType() *string {
	return s.MappingType
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) SetExtractKey(v string) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList {
	s.ExtractKey = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) SetExtractKeySpec(v string) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList {
	s.ExtractKeySpec = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) SetMappingType(v string) *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList {
	s.MappingType = &v
	return s
}

func (s *ListGatewayRouteResponseBodyDataResultRouteServicesHttpDubboTranscoderMothedMapListParamMapsList) Validate() error {
	return dara.Validate(s)
}
