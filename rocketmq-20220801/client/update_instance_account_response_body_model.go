// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateInstanceAccountResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateInstanceAccountResponseBody
	GetCode() *string
	SetData(v bool) *UpdateInstanceAccountResponseBody
	GetData() *bool
	SetDynamicCode(v string) *UpdateInstanceAccountResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateInstanceAccountResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *UpdateInstanceAccountResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateInstanceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceAccountResponseBody
	GetSuccess() *bool
}

type UpdateInstanceAccountResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// Instance.NotFound
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
	// instanceId
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
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateInstanceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAccountResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateInstanceAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstanceAccountResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateInstanceAccountResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateInstanceAccountResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateInstanceAccountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateInstanceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceAccountResponseBody) SetAccessDeniedDetail(v string) *UpdateInstanceAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetCode(v string) *UpdateInstanceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetData(v bool) *UpdateInstanceAccountResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetDynamicCode(v string) *UpdateInstanceAccountResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetDynamicMessage(v string) *UpdateInstanceAccountResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetMessage(v string) *UpdateInstanceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetRequestId(v string) *UpdateInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetSuccess(v bool) *UpdateInstanceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
