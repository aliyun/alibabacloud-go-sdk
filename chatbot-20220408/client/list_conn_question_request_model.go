// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListConnQuestionRequest
	GetAgentKey() *string
	SetKnowledgeId(v int64) *ListConnQuestionRequest
	GetKnowledgeId() *int64
}

type ListConnQuestionRequest struct {
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

func (s ListConnQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConnQuestionRequest) GoString() string {
	return s.String()
}

func (s *ListConnQuestionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListConnQuestionRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *ListConnQuestionRequest) SetAgentKey(v string) *ListConnQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *ListConnQuestionRequest) SetKnowledgeId(v int64) *ListConnQuestionRequest {
	s.KnowledgeId = &v
	return s
}

func (s *ListConnQuestionRequest) Validate() error {
	return dara.Validate(s)
}
