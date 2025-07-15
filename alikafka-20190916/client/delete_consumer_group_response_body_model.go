// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteConsumerGroupResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteConsumerGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteConsumerGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteConsumerGroupResponseBody
	GetSuccess() *bool
}

type DeleteConsumerGroupResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteConsumerGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConsumerGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteConsumerGroupResponseBody) SetCode(v int32) *DeleteConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetMessage(v string) *DeleteConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetRequestId(v string) *DeleteConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetSuccess(v bool) *DeleteConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
