// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerOffsetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateConsumerOffsetResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateConsumerOffsetResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConsumerOffsetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateConsumerOffsetResponseBody
	GetSuccess() *bool
}

type UpdateConsumerOffsetResponseBody struct {
	// Return code. A return value of **200*	- indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 56729737-C428-4E1B-AC68-7A8C2D5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateConsumerOffsetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerOffsetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConsumerOffsetResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateConsumerOffsetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConsumerOffsetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConsumerOffsetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateConsumerOffsetResponseBody) SetCode(v int32) *UpdateConsumerOffsetResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConsumerOffsetResponseBody) SetMessage(v string) *UpdateConsumerOffsetResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConsumerOffsetResponseBody) SetRequestId(v string) *UpdateConsumerOffsetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConsumerOffsetResponseBody) SetSuccess(v bool) *UpdateConsumerOffsetResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConsumerOffsetResponseBody) Validate() error {
	return dara.Validate(s)
}
