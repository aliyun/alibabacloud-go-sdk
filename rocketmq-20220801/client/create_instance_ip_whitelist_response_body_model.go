// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceIpWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateInstanceIpWhitelistResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateInstanceIpWhitelistResponseBody
	GetCode() *string
	SetData(v bool) *CreateInstanceIpWhitelistResponseBody
	GetData() *bool
	SetDynamicCode(v string) *CreateInstanceIpWhitelistResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateInstanceIpWhitelistResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *CreateInstanceIpWhitelistResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateInstanceIpWhitelistResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstanceIpWhitelistResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceIpWhitelistResponseBody
	GetSuccess() *bool
}

type CreateInstanceIpWhitelistResponseBody struct {
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
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
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
	// A07B41BD-6DD3-5349-9E76-00303DF04BBE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateInstanceIpWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceIpWhitelistResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateInstanceIpWhitelistResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceIpWhitelistResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateInstanceIpWhitelistResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateInstanceIpWhitelistResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateInstanceIpWhitelistResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateInstanceIpWhitelistResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstanceIpWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceIpWhitelistResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceIpWhitelistResponseBody) SetAccessDeniedDetail(v string) *CreateInstanceIpWhitelistResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetCode(v string) *CreateInstanceIpWhitelistResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetData(v bool) *CreateInstanceIpWhitelistResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetDynamicCode(v string) *CreateInstanceIpWhitelistResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetDynamicMessage(v string) *CreateInstanceIpWhitelistResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetHttpStatusCode(v int32) *CreateInstanceIpWhitelistResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetMessage(v string) *CreateInstanceIpWhitelistResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetRequestId(v string) *CreateInstanceIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetSuccess(v bool) *CreateInstanceIpWhitelistResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
