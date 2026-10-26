// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSteerAgentConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccepted(v bool) *SteerAgentConversationResponseBody
	GetAccepted() *bool
	SetMessageId(v string) *SteerAgentConversationResponseBody
	GetMessageId() *string
}

type SteerAgentConversationResponseBody struct {
	// example:
	//
	// true
	Accepted *bool `json:"accepted,omitempty" xml:"accepted,omitempty"`
	// example:
	//
	// msg-20260819-0001
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
}

func (s SteerAgentConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SteerAgentConversationResponseBody) GoString() string {
	return s.String()
}

func (s *SteerAgentConversationResponseBody) GetAccepted() *bool {
	return s.Accepted
}

func (s *SteerAgentConversationResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *SteerAgentConversationResponseBody) SetAccepted(v bool) *SteerAgentConversationResponseBody {
	s.Accepted = &v
	return s
}

func (s *SteerAgentConversationResponseBody) SetMessageId(v string) *SteerAgentConversationResponseBody {
	s.MessageId = &v
	return s
}

func (s *SteerAgentConversationResponseBody) Validate() error {
	return dara.Validate(s)
}
