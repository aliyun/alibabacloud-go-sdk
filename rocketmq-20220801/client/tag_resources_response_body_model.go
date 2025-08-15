// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TagResourcesResponseBody
	GetCode() *string
	SetData(v bool) *TagResourcesResponseBody
	GetData() *bool
	SetDynamicCode(v string) *TagResourcesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *TagResourcesResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *TagResourcesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *TagResourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *TagResourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TagResourcesResponseBody
	GetSuccess() *bool
}

type TagResourcesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
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
	// The error code.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B962390-D84B-5D44-8C11-79DF40299D41
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *TagResourcesResponseBody) GetData() *bool {
	return s.Data
}

func (s *TagResourcesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *TagResourcesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *TagResourcesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TagResourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagResourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TagResourcesResponseBody) SetCode(v string) *TagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBody) SetData(v bool) *TagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *TagResourcesResponseBody) SetDynamicCode(v string) *TagResourcesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *TagResourcesResponseBody) SetDynamicMessage(v string) *TagResourcesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *TagResourcesResponseBody) SetHttpStatusCode(v int32) *TagResourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TagResourcesResponseBody) SetMessage(v string) *TagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *TagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
