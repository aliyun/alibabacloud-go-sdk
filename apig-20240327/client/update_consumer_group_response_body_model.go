// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateConsumerGroupResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateConsumerGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateConsumerGroupResponseBody
	GetRequestId() *string
}

type UpdateConsumerGroupResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateConsumerGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConsumerGroupResponseBody) SetCode(v string) *UpdateConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetMessage(v string) *UpdateConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetRequestId(v string) *UpdateConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
