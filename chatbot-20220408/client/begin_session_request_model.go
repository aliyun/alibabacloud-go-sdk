// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBeginSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *BeginSessionRequest
	GetAgentKey() *string
	SetInstanceId(v string) *BeginSessionRequest
	GetInstanceId() *string
}

type BeginSessionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s BeginSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s BeginSessionRequest) GoString() string {
	return s.String()
}

func (s *BeginSessionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *BeginSessionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BeginSessionRequest) SetAgentKey(v string) *BeginSessionRequest {
	s.AgentKey = &v
	return s
}

func (s *BeginSessionRequest) SetInstanceId(v string) *BeginSessionRequest {
	s.InstanceId = &v
	return s
}

func (s *BeginSessionRequest) Validate() error {
	return dara.Validate(s)
}
