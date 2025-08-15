// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateInstanceAccountResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateInstanceAccountResponseBody
	GetCode() *string
	SetData(v bool) *CreateInstanceAccountResponseBody
	GetData() *bool
	SetDynamicCode(v string) *CreateInstanceAccountResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateInstanceAccountResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *CreateInstanceAccountResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateInstanceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstanceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceAccountResponseBody
	GetSuccess() *bool
}

type CreateInstanceAccountResponseBody struct {
	// No permission details
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// 3AE0999C-8DBA-5CEE-8D9A-BE8D4A90DF8D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateInstanceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateInstanceAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceAccountResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateInstanceAccountResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateInstanceAccountResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateInstanceAccountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateInstanceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstanceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceAccountResponseBody) SetAccessDeniedDetail(v string) *CreateInstanceAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetCode(v string) *CreateInstanceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetData(v bool) *CreateInstanceAccountResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetDynamicCode(v string) *CreateInstanceAccountResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetDynamicMessage(v string) *CreateInstanceAccountResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetHttpStatusCode(v int32) *CreateInstanceAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetMessage(v string) *CreateInstanceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetRequestId(v string) *CreateInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetSuccess(v bool) *CreateInstanceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
