// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFaqShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateFaqShrinkRequest
	GetAgentKey() *string
	SetCategoryId(v int64) *UpdateFaqShrinkRequest
	GetCategoryId() *int64
	SetEndDate(v string) *UpdateFaqShrinkRequest
	GetEndDate() *string
	SetKnowledgeId(v int64) *UpdateFaqShrinkRequest
	GetKnowledgeId() *int64
	SetStartDate(v string) *UpdateFaqShrinkRequest
	GetStartDate() *string
	SetTagIdListShrink(v string) *UpdateFaqShrinkRequest
	GetTagIdListShrink() *string
	SetTitle(v string) *UpdateFaqShrinkRequest
	GetTitle() *string
}

type UpdateFaqShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30000049006
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2030-12-31T16:00:00Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001905617
	KnowledgeId *int64 `json:"KnowledgeId,omitempty" xml:"KnowledgeId,omitempty"`
	// example:
	//
	// 2022-05-27T05:18:20Z
	StartDate       *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TagIdListShrink *string `json:"TagIdList,omitempty" xml:"TagIdList,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateFaqShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFaqShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaqShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateFaqShrinkRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdateFaqShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *UpdateFaqShrinkRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *UpdateFaqShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *UpdateFaqShrinkRequest) GetTagIdListShrink() *string {
	return s.TagIdListShrink
}

func (s *UpdateFaqShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateFaqShrinkRequest) SetAgentKey(v string) *UpdateFaqShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateFaqShrinkRequest) SetCategoryId(v int64) *UpdateFaqShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateFaqShrinkRequest) SetEndDate(v string) *UpdateFaqShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateFaqShrinkRequest) SetKnowledgeId(v int64) *UpdateFaqShrinkRequest {
	s.KnowledgeId = &v
	return s
}

func (s *UpdateFaqShrinkRequest) SetStartDate(v string) *UpdateFaqShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateFaqShrinkRequest) SetTagIdListShrink(v string) *UpdateFaqShrinkRequest {
	s.TagIdListShrink = &v
	return s
}

func (s *UpdateFaqShrinkRequest) SetTitle(v string) *UpdateFaqShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateFaqShrinkRequest) Validate() error {
	return dara.Validate(s)
}
