// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAgentConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetStopped(v bool) *StopAgentConversationResponseBody
	GetStopped() *bool
}

type StopAgentConversationResponseBody struct {
	// example:
	//
	// true
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
}

func (s StopAgentConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAgentConversationResponseBody) GoString() string {
	return s.String()
}

func (s *StopAgentConversationResponseBody) GetStopped() *bool {
	return s.Stopped
}

func (s *StopAgentConversationResponseBody) SetStopped(v bool) *StopAgentConversationResponseBody {
	s.Stopped = &v
	return s
}

func (s *StopAgentConversationResponseBody) Validate() error {
	return dara.Validate(s)
}
