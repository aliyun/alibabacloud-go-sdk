// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteInstanceAccountResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteInstanceAccountResponseBody
	GetCode() *string
	SetData(v bool) *DeleteInstanceAccountResponseBody
	GetData() *bool
	SetDynamicCode(v string) *DeleteInstanceAccountResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteInstanceAccountResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeleteInstanceAccountResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteInstanceAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstanceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstanceAccountResponseBody
	GetSuccess() *bool
}

type DeleteInstanceAccountResponseBody struct {
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
	// 157DF7D4-53FB-58C6-BEBC-A9400E7EF68A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteInstanceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceAccountResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteInstanceAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstanceAccountResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteInstanceAccountResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteInstanceAccountResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteInstanceAccountResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteInstanceAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstanceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstanceAccountResponseBody) SetAccessDeniedDetail(v string) *DeleteInstanceAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetCode(v string) *DeleteInstanceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetData(v bool) *DeleteInstanceAccountResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetDynamicCode(v string) *DeleteInstanceAccountResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetDynamicMessage(v string) *DeleteInstanceAccountResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetMessage(v string) *DeleteInstanceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetRequestId(v string) *DeleteInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetSuccess(v bool) *DeleteInstanceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
