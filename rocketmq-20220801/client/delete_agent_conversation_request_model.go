// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentConversationRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteAgentConversationRequest struct {
}

func (s DeleteAgentConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentConversationRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentConversationRequest) Validate() error {
	return dara.Validate(s)
}
