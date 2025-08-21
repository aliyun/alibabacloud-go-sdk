// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *AssociateShrinkRequest
	GetAgentKey() *string
	SetInstanceId(v string) *AssociateShrinkRequest
	GetInstanceId() *string
	SetPerspectiveShrink(v string) *AssociateShrinkRequest
	GetPerspectiveShrink() *string
	SetRecommendNum(v int64) *AssociateShrinkRequest
	GetRecommendNum() *int64
	SetSessionId(v string) *AssociateShrinkRequest
	GetSessionId() *string
	SetUtterance(v string) *AssociateShrinkRequest
	GetUtterance() *string
}

type AssociateShrinkRequest struct {
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
	// ["qyzzVfyFfa"]
	PerspectiveShrink *string `json:"Perspective,omitempty" xml:"Perspective,omitempty"`
	// example:
	//
	// 8
	RecommendNum *int64 `json:"RecommendNum,omitempty" xml:"RecommendNum,omitempty"`
	// example:
	//
	// 7c3cec23cc8940bc9db4a318c8f4f0aa
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 公积金提取
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s AssociateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateShrinkRequest) GoString() string {
	return s.String()
}

func (s *AssociateShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *AssociateShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssociateShrinkRequest) GetPerspectiveShrink() *string {
	return s.PerspectiveShrink
}

func (s *AssociateShrinkRequest) GetRecommendNum() *int64 {
	return s.RecommendNum
}

func (s *AssociateShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *AssociateShrinkRequest) GetUtterance() *string {
	return s.Utterance
}

func (s *AssociateShrinkRequest) SetAgentKey(v string) *AssociateShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *AssociateShrinkRequest) SetInstanceId(v string) *AssociateShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateShrinkRequest) SetPerspectiveShrink(v string) *AssociateShrinkRequest {
	s.PerspectiveShrink = &v
	return s
}

func (s *AssociateShrinkRequest) SetRecommendNum(v int64) *AssociateShrinkRequest {
	s.RecommendNum = &v
	return s
}

func (s *AssociateShrinkRequest) SetSessionId(v string) *AssociateShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *AssociateShrinkRequest) SetUtterance(v string) *AssociateShrinkRequest {
	s.Utterance = &v
	return s
}

func (s *AssociateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
