// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateConsumerGroupResponseBody
	GetCode() *string
	SetData(v bool) *CreateConsumerGroupResponseBody
	GetData() *bool
	SetDynamicCode(v string) *CreateConsumerGroupResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *CreateConsumerGroupResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *CreateConsumerGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateConsumerGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateConsumerGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateConsumerGroupResponseBody
	GetSuccess() *bool
}

type CreateConsumerGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidConsumerGroupId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// consumerGroupId
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
	// Parameter consumerGroupId is invalid.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateConsumerGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateConsumerGroupResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *CreateConsumerGroupResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *CreateConsumerGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateConsumerGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConsumerGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateConsumerGroupResponseBody) SetCode(v string) *CreateConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetData(v bool) *CreateConsumerGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetDynamicCode(v string) *CreateConsumerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetDynamicMessage(v string) *CreateConsumerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetHttpStatusCode(v int32) *CreateConsumerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetMessage(v string) *CreateConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetRequestId(v string) *CreateConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetSuccess(v bool) *CreateConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
