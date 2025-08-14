// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageChatTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendMessageChatTextResponseBody
	GetRequestId() *string
}

type SendMessageChatTextResponseBody struct {
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendMessageChatTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendMessageChatTextResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageChatTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendMessageChatTextResponseBody) SetRequestId(v string) *SendMessageChatTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendMessageChatTextResponseBody) Validate() error {
	return dara.Validate(s)
}
