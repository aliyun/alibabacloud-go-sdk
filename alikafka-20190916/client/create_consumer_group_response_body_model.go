// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateConsumerGroupResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateConsumerGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateConsumerGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateConsumerGroupResponseBody
	GetSuccess() *bool
}

type CreateConsumerGroupResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E57A8862-DF68-4055-8E55-B80CB4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponseBody) GetCode() *int32 {
	return s.Code
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

func (s *CreateConsumerGroupResponseBody) SetCode(v int32) *CreateConsumerGroupResponseBody {
	s.Code = &v
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
