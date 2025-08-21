// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateSimQuestionRequest
	GetAgentKey() *string
	SetKnowledgeId(v int64) *CreateSimQuestionRequest
	GetKnowledgeId() *int64
	SetTitle(v string) *CreateSimQuestionRequest
	GetTitle() *string
}

type CreateSimQuestionRequest struct {
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
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateSimQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSimQuestionRequest) GoString() string {
	return s.String()
}

func (s *CreateSimQuestionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateSimQuestionRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *CreateSimQuestionRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateSimQuestionRequest) SetAgentKey(v string) *CreateSimQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateSimQuestionRequest) SetKnowledgeId(v int64) *CreateSimQuestionRequest {
	s.KnowledgeId = &v
	return s
}

func (s *CreateSimQuestionRequest) SetTitle(v string) *CreateSimQuestionRequest {
	s.Title = &v
	return s
}

func (s *CreateSimQuestionRequest) Validate() error {
	return dara.Validate(s)
}
