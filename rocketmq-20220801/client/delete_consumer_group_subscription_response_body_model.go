// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerGroupSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteConsumerGroupSubscriptionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteConsumerGroupSubscriptionResponseBody
	GetCode() *string
	SetData(v bool) *DeleteConsumerGroupSubscriptionResponseBody
	GetData() *bool
	SetDynamicCode(v string) *DeleteConsumerGroupSubscriptionResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DeleteConsumerGroupSubscriptionResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *DeleteConsumerGroupSubscriptionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteConsumerGroupSubscriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteConsumerGroupSubscriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteConsumerGroupSubscriptionResponseBody
	GetSuccess() *bool
}

type DeleteConsumerGroupSubscriptionResponseBody struct {
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
	// 200
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
	// 723CDA5C-E25C-5EAF-9601-******
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteConsumerGroupSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerGroupSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetAccessDeniedDetail(v string) *DeleteConsumerGroupSubscriptionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetCode(v string) *DeleteConsumerGroupSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetData(v bool) *DeleteConsumerGroupSubscriptionResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetDynamicCode(v string) *DeleteConsumerGroupSubscriptionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetDynamicMessage(v string) *DeleteConsumerGroupSubscriptionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetHttpStatusCode(v int32) *DeleteConsumerGroupSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetMessage(v string) *DeleteConsumerGroupSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetRequestId(v string) *DeleteConsumerGroupSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetSuccess(v bool) *DeleteConsumerGroupSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
