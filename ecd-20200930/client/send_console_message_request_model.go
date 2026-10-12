// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendConsoleMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v string) *SendConsoleMessageRequest
	GetMessageId() *string
}

type SendConsoleMessageRequest struct {
	// The ID of the console message to be sent. You can obtain the ID from the return value of the message creation operation or the query result of DescribeConsoleMessages. Each call supports only one message ID, and the message must belong to the current Alibaba Cloud account. Only messages in the Published state can be sent.
	//
	// This parameter is required.
	//
	// example:
	//
	// msg-123456
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s SendConsoleMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendConsoleMessageRequest) GoString() string {
	return s.String()
}

func (s *SendConsoleMessageRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *SendConsoleMessageRequest) SetMessageId(v string) *SendConsoleMessageRequest {
	s.MessageId = &v
	return s
}

func (s *SendConsoleMessageRequest) Validate() error {
	return dara.Validate(s)
}
