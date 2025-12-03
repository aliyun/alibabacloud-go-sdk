// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatMessagesTaskStopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChatMessagesTaskStopResponseBody
	GetRequestId() *string
	SetResult(v string) *ChatMessagesTaskStopResponseBody
	GetResult() *string
}

type ChatMessagesTaskStopResponseBody struct {
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ChatMessagesTaskStopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatMessagesTaskStopResponseBody) GoString() string {
	return s.String()
}

func (s *ChatMessagesTaskStopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatMessagesTaskStopResponseBody) GetResult() *string {
	return s.Result
}

func (s *ChatMessagesTaskStopResponseBody) SetRequestId(v string) *ChatMessagesTaskStopResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatMessagesTaskStopResponseBody) SetResult(v string) *ChatMessagesTaskStopResponseBody {
	s.Result = &v
	return s
}

func (s *ChatMessagesTaskStopResponseBody) Validate() error {
	return dara.Validate(s)
}
