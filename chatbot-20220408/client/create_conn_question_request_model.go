// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateConnQuestionRequest
	GetAgentKey() *string
	SetConnQuestionId(v int64) *CreateConnQuestionRequest
	GetConnQuestionId() *int64
	SetKnowledgeId(v int64) *CreateConnQuestionRequest
	GetKnowledgeId() *int64
}

type CreateConnQuestionRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30002654628
	ConnQuestionId *int64 `json:"ConnQuestionId,omitempty" xml:"ConnQuestionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30002174773
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
}

func (s CreateConnQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConnQuestionRequest) GoString() string {
	return s.String()
}

func (s *CreateConnQuestionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateConnQuestionRequest) GetConnQuestionId() *int64 {
	return s.ConnQuestionId
}

func (s *CreateConnQuestionRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *CreateConnQuestionRequest) SetAgentKey(v string) *CreateConnQuestionRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateConnQuestionRequest) SetConnQuestionId(v int64) *CreateConnQuestionRequest {
	s.ConnQuestionId = &v
	return s
}

func (s *CreateConnQuestionRequest) SetKnowledgeId(v int64) *CreateConnQuestionRequest {
	s.KnowledgeId = &v
	return s
}

func (s *CreateConnQuestionRequest) Validate() error {
	return dara.Validate(s)
}
