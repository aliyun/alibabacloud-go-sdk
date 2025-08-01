// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderClusterHealthCheckRiskNoticeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OrderClusterHealthCheckRiskNoticeResponseBody
	GetCode() *int32
	SetData(v bool) *OrderClusterHealthCheckRiskNoticeResponseBody
	GetData() *bool
	SetDynamicCode(v string) *OrderClusterHealthCheckRiskNoticeResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *OrderClusterHealthCheckRiskNoticeResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *OrderClusterHealthCheckRiskNoticeResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *OrderClusterHealthCheckRiskNoticeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *OrderClusterHealthCheckRiskNoticeResponseBody
	GetMessage() *string
	SetRequestId(v string) *OrderClusterHealthCheckRiskNoticeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OrderClusterHealthCheckRiskNoticeResponseBody
	GetSuccess() *bool
}

type OrderClusterHealthCheckRiskNoticeResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the subscription was successful.
	//
	// example:
	//
	// null
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// > If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
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
	// The message returned.
	//
	// 	- If the request is successful, a success message is returned.
	//
	// 	- If the request fails, an error message is returned, such as the "TaskId not found" message.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AF21683A-29C7-4853-AC0F-B5ADEE4D****
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

func (s OrderClusterHealthCheckRiskNoticeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OrderClusterHealthCheckRiskNoticeResponseBody) GoString() string {
	return s.String()
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) GetData() *bool {
	return s.Data
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) SetCode(v int32) *OrderClusterHealthCheckRiskNoticeResponseBody {
	s.Code = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) SetData(v bool) *OrderClusterHealthCheckRiskNoticeResponseBody {
	s.Data = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) SetDynamicCode(v string) *OrderClusterHealthCheckRiskNoticeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) SetDynamicMessage(v string) *OrderClusterHealthCheckRiskNoticeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) SetErrorCode(v string) *OrderClusterHealthCheckRiskNoticeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) SetHttpStatusCode(v int32) *OrderClusterHealthCheckRiskNoticeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) SetMessage(v string) *OrderClusterHealthCheckRiskNoticeResponseBody {
	s.Message = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) SetRequestId(v string) *OrderClusterHealthCheckRiskNoticeResponseBody {
	s.RequestId = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) SetSuccess(v bool) *OrderClusterHealthCheckRiskNoticeResponseBody {
	s.Success = &v
	return s
}

func (s *OrderClusterHealthCheckRiskNoticeResponseBody) Validate() error {
	return dara.Validate(s)
}
