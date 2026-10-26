// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAgentConversationRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StopAgentConversationRequest struct {
}

func (s StopAgentConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAgentConversationRequest) GoString() string {
	return s.String()
}

func (s *StopAgentConversationRequest) Validate() error {
	return dara.Validate(s)
}
