// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentConversationRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAgentConversationRequest struct {
}

func (s GetAgentConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConversationRequest) GoString() string {
	return s.String()
}

func (s *GetAgentConversationRequest) Validate() error {
	return dara.Validate(s)
}
