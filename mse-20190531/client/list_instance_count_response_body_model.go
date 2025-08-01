// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListInstanceCountResponseBody
	GetCode() *int32
	SetData(v []*int32) *ListInstanceCountResponseBody
	GetData() []*int32
	SetDynamicCode(v string) *ListInstanceCountResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListInstanceCountResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *ListInstanceCountResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *ListInstanceCountResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstanceCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstanceCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstanceCountResponseBody
	GetSuccess() *bool
}

type ListInstanceCountResponseBody struct {
	// The status code. A value of 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data entries returned.
	Data []*int32 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// %s
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace `%s` in the `ErrMessage` parameter.
	//
	// >  If the return value of the `ErrMessage` parameter is `The Value of Input Parameter %s is not valid` and the return value of the `DynamicMessage` parameter is `DtsJobId`, the specified `DtsJobId` parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed. If the request failed, the ErrorCode parameter is returned. For more information, see the [Error codes](https://help.aliyun.com/document_detail/456441.html) section of this topic.
	//
	// example:
	//
	// mse-100-100
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request is successful, a success message is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 316F5F64-F73D-42DC-8632-01E308B6****
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

func (s ListInstanceCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceCountResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceCountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListInstanceCountResponseBody) GetData() []*int32 {
	return s.Data
}

func (s *ListInstanceCountResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListInstanceCountResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListInstanceCountResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListInstanceCountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstanceCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceCountResponseBody) SetCode(v int32) *ListInstanceCountResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceCountResponseBody) SetData(v []*int32) *ListInstanceCountResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceCountResponseBody) SetDynamicCode(v string) *ListInstanceCountResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListInstanceCountResponseBody) SetDynamicMessage(v string) *ListInstanceCountResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListInstanceCountResponseBody) SetErrorCode(v string) *ListInstanceCountResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListInstanceCountResponseBody) SetHttpStatusCode(v int32) *ListInstanceCountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstanceCountResponseBody) SetMessage(v string) *ListInstanceCountResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceCountResponseBody) SetRequestId(v string) *ListInstanceCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceCountResponseBody) SetSuccess(v bool) *ListInstanceCountResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceCountResponseBody) Validate() error {
	return dara.Validate(s)
}
