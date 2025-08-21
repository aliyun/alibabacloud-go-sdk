// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListCategoryRequest
	GetAgentKey() *string
	SetKnowledgeType(v int32) *ListCategoryRequest
	GetKnowledgeType() *int32
	SetParentCategoryId(v int64) *ListCategoryRequest
	GetParentCategoryId() *int64
}

type ListCategoryRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	KnowledgeType *int32  `json:"KnowledgeType,omitempty" xml:"KnowledgeType,omitempty"`
	// example:
	//
	// -1
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s ListCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListCategoryRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListCategoryRequest) GetKnowledgeType() *int32 {
	return s.KnowledgeType
}

func (s *ListCategoryRequest) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *ListCategoryRequest) SetAgentKey(v string) *ListCategoryRequest {
	s.AgentKey = &v
	return s
}

func (s *ListCategoryRequest) SetKnowledgeType(v int32) *ListCategoryRequest {
	s.KnowledgeType = &v
	return s
}

func (s *ListCategoryRequest) SetParentCategoryId(v int64) *ListCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

func (s *ListCategoryRequest) Validate() error {
	return dara.Validate(s)
}
