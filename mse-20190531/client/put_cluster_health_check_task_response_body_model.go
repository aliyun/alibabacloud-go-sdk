// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutClusterHealthCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PutClusterHealthCheckTaskResponseBody
	GetCode() *int32
	SetData(v bool) *PutClusterHealthCheckTaskResponseBody
	GetData() *bool
	SetDynamicCode(v string) *PutClusterHealthCheckTaskResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *PutClusterHealthCheckTaskResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *PutClusterHealthCheckTaskResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *PutClusterHealthCheckTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PutClusterHealthCheckTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutClusterHealthCheckTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutClusterHealthCheckTaskResponseBody
	GetSuccess() *bool
}

type PutClusterHealthCheckTaskResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
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
	// 5B170A0D-2C5D-4CF8-B808-03966B86****
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

func (s PutClusterHealthCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutClusterHealthCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PutClusterHealthCheckTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PutClusterHealthCheckTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *PutClusterHealthCheckTaskResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *PutClusterHealthCheckTaskResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *PutClusterHealthCheckTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PutClusterHealthCheckTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PutClusterHealthCheckTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutClusterHealthCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutClusterHealthCheckTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutClusterHealthCheckTaskResponseBody) SetCode(v int32) *PutClusterHealthCheckTaskResponseBody {
	s.Code = &v
	return s
}

func (s *PutClusterHealthCheckTaskResponseBody) SetData(v bool) *PutClusterHealthCheckTaskResponseBody {
	s.Data = &v
	return s
}

func (s *PutClusterHealthCheckTaskResponseBody) SetDynamicCode(v string) *PutClusterHealthCheckTaskResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PutClusterHealthCheckTaskResponseBody) SetDynamicMessage(v string) *PutClusterHealthCheckTaskResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PutClusterHealthCheckTaskResponseBody) SetErrorCode(v string) *PutClusterHealthCheckTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PutClusterHealthCheckTaskResponseBody) SetHttpStatusCode(v int32) *PutClusterHealthCheckTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PutClusterHealthCheckTaskResponseBody) SetMessage(v string) *PutClusterHealthCheckTaskResponseBody {
	s.Message = &v
	return s
}

func (s *PutClusterHealthCheckTaskResponseBody) SetRequestId(v string) *PutClusterHealthCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutClusterHealthCheckTaskResponseBody) SetSuccess(v bool) *PutClusterHealthCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *PutClusterHealthCheckTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
