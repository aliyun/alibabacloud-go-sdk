// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateInstanceAclResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateInstanceAclResponseBody
	GetCode() *string
	SetData(v bool) *CreateInstanceAclResponseBody
	GetData() *bool
	SetDynamicCode(v string) *CreateInstanceAclResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateInstanceAclResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *CreateInstanceAclResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateInstanceAclResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstanceAclResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceAclResponseBody
	GetSuccess() *bool
}

type CreateInstanceAclResponseBody struct {
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
	// C7E8AE3A-219B-52EE-BE32-4036F5F88833
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateInstanceAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceAclResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceAclResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateInstanceAclResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceAclResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateInstanceAclResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateInstanceAclResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateInstanceAclResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateInstanceAclResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstanceAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceAclResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceAclResponseBody) SetAccessDeniedDetail(v string) *CreateInstanceAclResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetCode(v string) *CreateInstanceAclResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetData(v bool) *CreateInstanceAclResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetDynamicCode(v string) *CreateInstanceAclResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetDynamicMessage(v string) *CreateInstanceAclResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetHttpStatusCode(v int32) *CreateInstanceAclResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetMessage(v string) *CreateInstanceAclResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetRequestId(v string) *CreateInstanceAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetSuccess(v bool) *CreateInstanceAclResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceAclResponseBody) Validate() error {
	return dara.Validate(s)
}
