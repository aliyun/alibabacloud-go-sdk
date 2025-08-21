// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaqRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DeleteFaqRequest
	GetAgentKey() *string
	SetKnowledgeId(v int64) *DeleteFaqRequest
	GetKnowledgeId() *int64
}

type DeleteFaqRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s DeleteFaqRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaqRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaqRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DeleteFaqRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *DeleteFaqRequest) SetAgentKey(v string) *DeleteFaqRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteFaqRequest) SetKnowledgeId(v int64) *DeleteFaqRequest {
	s.KnowledgeId = &v
	return s
}

func (s *DeleteFaqRequest) Validate() error {
	return dara.Validate(s)
}
