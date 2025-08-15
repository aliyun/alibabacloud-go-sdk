// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UntagResourcesResponseBody
	GetCode() *string
	SetData(v bool) *UntagResourcesResponseBody
	GetData() *bool
	SetDynamicCode(v string) *UntagResourcesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *UntagResourcesResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *UntagResourcesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UntagResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *UntagResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UntagResourcesResponseBody
	GetSuccess() *bool
}

type UntagResourcesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
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
	// Parameter deliveryOrderType is invalid.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07B41BD-6DD3-5349-9E76-00303DF04BBE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *UntagResourcesResponseBody) GetData() *bool {
	return s.Data
}

func (s *UntagResourcesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *UntagResourcesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *UntagResourcesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UntagResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UntagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UntagResourcesResponseBody) SetCode(v string) *UntagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBody) SetData(v bool) *UntagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *UntagResourcesResponseBody) SetDynamicCode(v string) *UntagResourcesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UntagResourcesResponseBody) SetDynamicMessage(v string) *UntagResourcesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UntagResourcesResponseBody) SetHttpStatusCode(v int32) *UntagResourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UntagResourcesResponseBody) SetMessage(v string) *UntagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *UntagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
