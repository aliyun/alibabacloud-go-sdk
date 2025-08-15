// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceIpWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteInstanceIpWhitelistResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteInstanceIpWhitelistResponseBody
	GetCode() *string
	SetData(v bool) *DeleteInstanceIpWhitelistResponseBody
	GetData() *bool
	SetDynamicCode(v string) *DeleteInstanceIpWhitelistResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteInstanceIpWhitelistResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeleteInstanceIpWhitelistResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteInstanceIpWhitelistResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstanceIpWhitelistResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstanceIpWhitelistResponseBody
	GetSuccess() *bool
}

type DeleteInstanceIpWhitelistResponseBody struct {
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
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16425867-C948-5A0C-9A24-5259727BE727
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteInstanceIpWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceIpWhitelistResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteInstanceIpWhitelistResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstanceIpWhitelistResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteInstanceIpWhitelistResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteInstanceIpWhitelistResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteInstanceIpWhitelistResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteInstanceIpWhitelistResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstanceIpWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceIpWhitelistResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetAccessDeniedDetail(v string) *DeleteInstanceIpWhitelistResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetCode(v string) *DeleteInstanceIpWhitelistResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetData(v bool) *DeleteInstanceIpWhitelistResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetDynamicCode(v string) *DeleteInstanceIpWhitelistResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetDynamicMessage(v string) *DeleteInstanceIpWhitelistResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceIpWhitelistResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetMessage(v string) *DeleteInstanceIpWhitelistResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetRequestId(v string) *DeleteInstanceIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetSuccess(v bool) *DeleteInstanceIpWhitelistResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
