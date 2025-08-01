// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthConsumerResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGatewayAuthConsumerResourceResponseBody
	GetCode() *int32
	SetData(v *ListGatewayAuthConsumerResourceResponseBodyData) *ListGatewayAuthConsumerResourceResponseBody
	GetData() *ListGatewayAuthConsumerResourceResponseBodyData
	SetDynamicCode(v string) *ListGatewayAuthConsumerResourceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListGatewayAuthConsumerResourceResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *ListGatewayAuthConsumerResourceResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *ListGatewayAuthConsumerResourceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGatewayAuthConsumerResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayAuthConsumerResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGatewayAuthConsumerResourceResponseBody
	GetSuccess() *bool
}

type ListGatewayAuthConsumerResourceResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *ListGatewayAuthConsumerResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code that is returned.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A73AC37C-C617-4E3A-8049-372CF49C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewayAuthConsumerResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthConsumerResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthConsumerResourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGatewayAuthConsumerResourceResponseBody) GetData() *ListGatewayAuthConsumerResourceResponseBodyData {
	return s.Data
}

func (s *ListGatewayAuthConsumerResourceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListGatewayAuthConsumerResourceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListGatewayAuthConsumerResourceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListGatewayAuthConsumerResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewayAuthConsumerResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayAuthConsumerResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayAuthConsumerResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewayAuthConsumerResourceResponseBody) SetCode(v int32) *ListGatewayAuthConsumerResourceResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBody) SetData(v *ListGatewayAuthConsumerResourceResponseBodyData) *ListGatewayAuthConsumerResourceResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBody) SetDynamicCode(v string) *ListGatewayAuthConsumerResourceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBody) SetDynamicMessage(v string) *ListGatewayAuthConsumerResourceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBody) SetErrorCode(v string) *ListGatewayAuthConsumerResourceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBody) SetHttpStatusCode(v int32) *ListGatewayAuthConsumerResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBody) SetMessage(v string) *ListGatewayAuthConsumerResourceResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBody) SetRequestId(v string) *ListGatewayAuthConsumerResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBody) SetSuccess(v bool) *ListGatewayAuthConsumerResourceResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListGatewayAuthConsumerResourceResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data structure.
	Result []*ListGatewayAuthConsumerResourceResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 11
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListGatewayAuthConsumerResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthConsumerResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthConsumerResourceResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayAuthConsumerResourceResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayAuthConsumerResourceResponseBodyData) GetResult() []*ListGatewayAuthConsumerResourceResponseBodyDataResult {
	return s.Result
}

func (s *ListGatewayAuthConsumerResourceResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ListGatewayAuthConsumerResourceResponseBodyData) SetPageNumber(v int32) *ListGatewayAuthConsumerResourceResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBodyData) SetPageSize(v int32) *ListGatewayAuthConsumerResourceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBodyData) SetResult(v []*ListGatewayAuthConsumerResourceResponseBodyDataResult) *ListGatewayAuthConsumerResourceResponseBodyData {
	s.Result = v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBodyData) SetTotalSize(v int64) *ListGatewayAuthConsumerResourceResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListGatewayAuthConsumerResourceResponseBodyDataResult struct {
	// The ID of the consumer.
	//
	// example:
	//
	// 2
	ConsumerId *int64 `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-5017305290e14centbrveca****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-07 18:07:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2022-01-07 18:07:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the authorized resource for the consumer.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The resource authorization status. Valid values:
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// example:
	//
	// true
	ResourceStatus *bool `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// 3091
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// test
	RouteName *string `json:"RouteName,omitempty" xml:"RouteName,omitempty"`
}

func (s ListGatewayAuthConsumerResourceResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthConsumerResourceResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) GetConsumerId() *int64 {
	return s.ConsumerId
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) GetResourceStatus() *bool {
	return s.ResourceStatus
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) GetRouteId() *int64 {
	return s.RouteId
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) GetRouteName() *string {
	return s.RouteName
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) SetConsumerId(v int64) *ListGatewayAuthConsumerResourceResponseBodyDataResult {
	s.ConsumerId = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) SetGatewayUniqueId(v string) *ListGatewayAuthConsumerResourceResponseBodyDataResult {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) SetGmtCreate(v string) *ListGatewayAuthConsumerResourceResponseBodyDataResult {
	s.GmtCreate = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) SetGmtModified(v string) *ListGatewayAuthConsumerResourceResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) SetId(v int64) *ListGatewayAuthConsumerResourceResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) SetResourceStatus(v bool) *ListGatewayAuthConsumerResourceResponseBodyDataResult {
	s.ResourceStatus = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) SetRouteId(v int64) *ListGatewayAuthConsumerResourceResponseBodyDataResult {
	s.RouteId = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) SetRouteName(v string) *ListGatewayAuthConsumerResourceResponseBodyDataResult {
	s.RouteName = &v
	return s
}

func (s *ListGatewayAuthConsumerResourceResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
