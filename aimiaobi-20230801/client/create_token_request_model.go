// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateTokenRequest
	GetAgentKey() *string
}

type CreateTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2daaa2e0c209xb26acb97009ea77bd4b_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s CreateTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateTokenRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateTokenRequest) SetAgentKey(v string) *CreateTokenRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateTokenRequest) Validate() error {
	return dara.Validate(s)
}
