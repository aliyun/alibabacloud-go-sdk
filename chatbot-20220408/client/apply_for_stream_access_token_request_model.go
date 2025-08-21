// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyForStreamAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ApplyForStreamAccessTokenRequest
	GetAgentKey() *string
}

type ApplyForStreamAccessTokenRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ApplyForStreamAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyForStreamAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *ApplyForStreamAccessTokenRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ApplyForStreamAccessTokenRequest) SetAgentKey(v string) *ApplyForStreamAccessTokenRequest {
	s.AgentKey = &v
	return s
}

func (s *ApplyForStreamAccessTokenRequest) Validate() error {
	return dara.Validate(s)
}
