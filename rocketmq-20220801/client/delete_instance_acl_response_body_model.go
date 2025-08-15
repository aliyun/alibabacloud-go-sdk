// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteInstanceAclResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteInstanceAclResponseBody
	GetCode() *string
	SetData(v bool) *DeleteInstanceAclResponseBody
	GetData() *bool
	SetDynamicCode(v string) *DeleteInstanceAclResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteInstanceAclResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeleteInstanceAclResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteInstanceAclResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstanceAclResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstanceAclResponseBody
	GetSuccess() *bool
}

type DeleteInstanceAclResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied due to the reason that the Resource Access Management (RAM) user does not have the required permissions.
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
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7358418D-83BD-507A-8079-*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteInstanceAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceAclResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceAclResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteInstanceAclResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstanceAclResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteInstanceAclResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteInstanceAclResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteInstanceAclResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteInstanceAclResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstanceAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceAclResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstanceAclResponseBody) SetAccessDeniedDetail(v string) *DeleteInstanceAclResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetCode(v string) *DeleteInstanceAclResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetData(v bool) *DeleteInstanceAclResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetDynamicCode(v string) *DeleteInstanceAclResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetDynamicMessage(v string) *DeleteInstanceAclResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceAclResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetMessage(v string) *DeleteInstanceAclResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetRequestId(v string) *DeleteInstanceAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetSuccess(v bool) *DeleteInstanceAclResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) Validate() error {
	return dara.Validate(s)
}
