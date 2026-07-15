// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteConsumerGroupResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteConsumerGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteConsumerGroupResponseBody
	GetRequestId() *string
}

type DeleteConsumerGroupResponseBody struct {
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

func (s DeleteConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteConsumerGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConsumerGroupResponseBody) SetCode(v string) *DeleteConsumerGroupResponseBody {
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

func (s *DeleteConsumerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
