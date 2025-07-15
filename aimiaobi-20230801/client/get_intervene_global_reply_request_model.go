// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterveneGlobalReplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *GetInterveneGlobalReplyRequest
	GetAgentKey() *string
}

type GetInterveneGlobalReplyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// fcb14f25c9ee41ccad33a049de8f941b_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s GetInterveneGlobalReplyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneGlobalReplyRequest) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *GetInterveneGlobalReplyRequest) SetAgentKey(v string) *GetInterveneGlobalReplyRequest {
	s.AgentKey = &v
	return s
}

func (s *GetInterveneGlobalReplyRequest) Validate() error {
	return dara.Validate(s)
}
