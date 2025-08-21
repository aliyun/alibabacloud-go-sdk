// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSimQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListSimQuestionRequest
	GetAgentKey() *string
	SetKnowledgeId(v int64) *ListSimQuestionRequest
	GetKnowledgeId() *int64
}

type ListSimQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30002299537
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s ListSimQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSimQuestionRequest) GoString() string {
	return s.String()
}

func (s *ListSimQuestionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListSimQuestionRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *ListSimQuestionRequest) SetAgentKey(v string) *ListSimQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *ListSimQuestionRequest) SetKnowledgeId(v int64) *ListSimQuestionRequest {
	s.KnowledgeId = &v
	return s
}

func (s *ListSimQuestionRequest) Validate() error {
	return dara.Validate(s)
}
