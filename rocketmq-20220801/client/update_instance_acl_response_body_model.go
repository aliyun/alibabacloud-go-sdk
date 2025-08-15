// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateInstanceAclResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateInstanceAclResponseBody
	GetCode() *string
	SetData(v bool) *UpdateInstanceAclResponseBody
	GetData() *bool
	SetDynamicCode(v string) *UpdateInstanceAclResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UpdateInstanceAclResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *UpdateInstanceAclResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateInstanceAclResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceAclResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceAclResponseBody
	GetSuccess() *bool
}

type UpdateInstanceAclResponseBody struct {
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
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
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
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C115601B-8736-5BBF-AC99-7FEAE1245A80
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateInstanceAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAclResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAclResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateInstanceAclResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstanceAclResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateInstanceAclResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UpdateInstanceAclResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UpdateInstanceAclResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateInstanceAclResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceAclResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceAclResponseBody) SetAccessDeniedDetail(v string) *UpdateInstanceAclResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetCode(v string) *UpdateInstanceAclResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetData(v bool) *UpdateInstanceAclResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetDynamicCode(v string) *UpdateInstanceAclResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetDynamicMessage(v string) *UpdateInstanceAclResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceAclResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetMessage(v string) *UpdateInstanceAclResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetRequestId(v string) *UpdateInstanceAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetSuccess(v bool) *UpdateInstanceAclResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) Validate() error {
	return dara.Validate(s)
}
