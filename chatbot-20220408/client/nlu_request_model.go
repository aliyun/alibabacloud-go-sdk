// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNluRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *NluRequest
	GetAgentKey() *string
	SetInstanceId(v string) *NluRequest
	GetInstanceId() *string
	SetUtterance(v string) *NluRequest
	GetUtterance() *string
}

type NluRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 北京的天气怎么样
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s NluRequest) String() string {
	return dara.Prettify(s)
}

func (s NluRequest) GoString() string {
	return s.String()
}

func (s *NluRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *NluRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *NluRequest) GetUtterance() *string {
	return s.Utterance
}

func (s *NluRequest) SetAgentKey(v string) *NluRequest {
	s.AgentKey = &v
	return s
}

func (s *NluRequest) SetInstanceId(v string) *NluRequest {
	s.InstanceId = &v
	return s
}

func (s *NluRequest) SetUtterance(v string) *NluRequest {
	s.Utterance = &v
	return s
}

func (s *NluRequest) Validate() error {
	return dara.Validate(s)
}
