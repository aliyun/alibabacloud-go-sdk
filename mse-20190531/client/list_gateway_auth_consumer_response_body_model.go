// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayAuthConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGatewayAuthConsumerResponseBody
	GetCode() *int32
	SetData(v *ListGatewayAuthConsumerResponseBodyData) *ListGatewayAuthConsumerResponseBody
	GetData() *ListGatewayAuthConsumerResponseBodyData
	SetDynamicCode(v string) *ListGatewayAuthConsumerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListGatewayAuthConsumerResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *ListGatewayAuthConsumerResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *ListGatewayAuthConsumerResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGatewayAuthConsumerResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayAuthConsumerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGatewayAuthConsumerResponseBody
	GetSuccess() *bool
}

type ListGatewayAuthConsumerResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data structure.
	Data *ListGatewayAuthConsumerResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the **%s*	- variable in the **ErrMessage*	- parameter.
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
	// dc63-465d-8ef5-20dc18af****
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

func (s ListGatewayAuthConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthConsumerResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGatewayAuthConsumerResponseBody) GetData() *ListGatewayAuthConsumerResponseBodyData {
	return s.Data
}

func (s *ListGatewayAuthConsumerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListGatewayAuthConsumerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListGatewayAuthConsumerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListGatewayAuthConsumerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewayAuthConsumerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayAuthConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayAuthConsumerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewayAuthConsumerResponseBody) SetCode(v int32) *ListGatewayAuthConsumerResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBody) SetData(v *ListGatewayAuthConsumerResponseBodyData) *ListGatewayAuthConsumerResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayAuthConsumerResponseBody) SetDynamicCode(v string) *ListGatewayAuthConsumerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBody) SetDynamicMessage(v string) *ListGatewayAuthConsumerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBody) SetErrorCode(v string) *ListGatewayAuthConsumerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBody) SetHttpStatusCode(v int32) *ListGatewayAuthConsumerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBody) SetMessage(v string) *ListGatewayAuthConsumerResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBody) SetRequestId(v string) *ListGatewayAuthConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBody) SetSuccess(v bool) *ListGatewayAuthConsumerResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayAuthConsumerResponseBodyData struct {
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
	// The returned information.
	Result []*ListGatewayAuthConsumerResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of returned entries.
	//
	// example:
	//
	// 9
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListGatewayAuthConsumerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthConsumerResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthConsumerResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayAuthConsumerResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayAuthConsumerResponseBodyData) GetResult() []*ListGatewayAuthConsumerResponseBodyDataResult {
	return s.Result
}

func (s *ListGatewayAuthConsumerResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ListGatewayAuthConsumerResponseBodyData) SetPageNumber(v int32) *ListGatewayAuthConsumerResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyData) SetPageSize(v int32) *ListGatewayAuthConsumerResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyData) SetResult(v []*ListGatewayAuthConsumerResponseBodyDataResult) *ListGatewayAuthConsumerResponseBodyData {
	s.Result = v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyData) SetTotalSize(v int64) *ListGatewayAuthConsumerResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyData) Validate() error {
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

type ListGatewayAuthConsumerResponseBodyDataResult struct {
	// The status of the consumer. Valid values:
	//
	// 	- true: enabled
	//
	// 	- false: disabled
	//
	// example:
	//
	// true
	ConsumerStatus *bool `json:"ConsumerStatus,omitempty" xml:"ConsumerStatus,omitempty"`
	// The description of the consumer.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// 2021-09-13 19:24:23
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2022-01-07 18:07:57
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the consumer.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the consumer.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The creator.
	//
	// example:
	//
	// 123
	PrimaryUser *string `json:"PrimaryUser,omitempty" xml:"PrimaryUser,omitempty"`
	// The authentication type. Valid values:
	//
	// 	- JWT
	//
	// example:
	//
	// JWT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayAuthConsumerResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayAuthConsumerResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) GetConsumerStatus() *bool {
	return s.ConsumerStatus
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) GetName() *string {
	return s.Name
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) GetPrimaryUser() *string {
	return s.PrimaryUser
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) SetConsumerStatus(v bool) *ListGatewayAuthConsumerResponseBodyDataResult {
	s.ConsumerStatus = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) SetDescription(v string) *ListGatewayAuthConsumerResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) SetGatewayUniqueId(v string) *ListGatewayAuthConsumerResponseBodyDataResult {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) SetGmtCreate(v string) *ListGatewayAuthConsumerResponseBodyDataResult {
	s.GmtCreate = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) SetGmtModified(v string) *ListGatewayAuthConsumerResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) SetId(v int64) *ListGatewayAuthConsumerResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) SetName(v string) *ListGatewayAuthConsumerResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) SetPrimaryUser(v string) *ListGatewayAuthConsumerResponseBodyDataResult {
	s.PrimaryUser = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) SetType(v string) *ListGatewayAuthConsumerResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *ListGatewayAuthConsumerResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
