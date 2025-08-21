// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSolutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListSolutionRequest
	GetAgentKey() *string
	SetKnowledgeId(v int64) *ListSolutionRequest
	GetKnowledgeId() *int64
}

type ListSolutionRequest struct {
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

func (s ListSolutionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSolutionRequest) GoString() string {
	return s.String()
}

func (s *ListSolutionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListSolutionRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *ListSolutionRequest) SetAgentKey(v string) *ListSolutionRequest {
	s.AgentKey = &v
	return s
}

func (s *ListSolutionRequest) SetKnowledgeId(v int64) *ListSolutionRequest {
	s.KnowledgeId = &v
	return s
}

func (s *ListSolutionRequest) Validate() error {
	return dara.Validate(s)
}
