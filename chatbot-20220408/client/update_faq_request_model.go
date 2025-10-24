// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFaqRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateFaqRequest
	GetAgentKey() *string
	SetCategoryId(v int64) *UpdateFaqRequest
	GetCategoryId() *int64
	SetEndDate(v string) *UpdateFaqRequest
	GetEndDate() *string
	SetKnowledgeId(v int64) *UpdateFaqRequest
	GetKnowledgeId() *int64
	SetStartDate(v string) *UpdateFaqRequest
	GetStartDate() *string
	SetTagIdList(v []*int64) *UpdateFaqRequest
	GetTagIdList() []*int64
	SetTitle(v string) *UpdateFaqRequest
	GetTitle() *string
}

type UpdateFaqRequest struct {
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
	StartDate *string  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TagIdList []*int64 `json:"TagIdList,omitempty" xml:"TagIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateFaqRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFaqRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaqRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateFaqRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdateFaqRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *UpdateFaqRequest) GetKnowledgeId() *int64 {
	return s.KnowledgeId
}

func (s *UpdateFaqRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *UpdateFaqRequest) GetTagIdList() []*int64 {
	return s.TagIdList
}

func (s *UpdateFaqRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateFaqRequest) SetAgentKey(v string) *UpdateFaqRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateFaqRequest) SetCategoryId(v int64) *UpdateFaqRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateFaqRequest) SetEndDate(v string) *UpdateFaqRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateFaqRequest) SetKnowledgeId(v int64) *UpdateFaqRequest {
	s.KnowledgeId = &v
	return s
}

func (s *UpdateFaqRequest) SetStartDate(v string) *UpdateFaqRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateFaqRequest) SetTagIdList(v []*int64) *UpdateFaqRequest {
	s.TagIdList = v
	return s
}

func (s *UpdateFaqRequest) SetTitle(v string) *UpdateFaqRequest {
	s.Title = &v
	return s
}

func (s *UpdateFaqRequest) Validate() error {
	return dara.Validate(s)
}
