// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFaqRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateFaqRequest
	GetAgentKey() *string
	SetCategoryId(v int64) *CreateFaqRequest
	GetCategoryId() *int64
	SetEndDate(v string) *CreateFaqRequest
	GetEndDate() *string
	SetSolutionContent(v string) *CreateFaqRequest
	GetSolutionContent() *string
	SetSolutionType(v int32) *CreateFaqRequest
	GetSolutionType() *int32
	SetStartDate(v string) *CreateFaqRequest
	GetStartDate() *string
	SetTagIdList(v []*int64) *CreateFaqRequest
	GetTagIdList() []*int64
	SetTitle(v string) *CreateFaqRequest
	GetTitle() *string
}

type CreateFaqRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000053274
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2030-12-31T16:00:00Z
	EndDate         *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	SolutionContent *string `json:"SolutionContent,omitempty" xml:"SolutionContent,omitempty"`
	// example:
	//
	// 0
	SolutionType *int32 `json:"SolutionType,omitempty" xml:"SolutionType,omitempty"`
	// example:
	//
	// 2022-05-25T16:28:36Z
	StartDate *string  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TagIdList []*int64 `json:"TagIdList,omitempty" xml:"TagIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateFaqRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFaqRequest) GoString() string {
	return s.String()
}

func (s *CreateFaqRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateFaqRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *CreateFaqRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateFaqRequest) GetSolutionContent() *string {
	return s.SolutionContent
}

func (s *CreateFaqRequest) GetSolutionType() *int32 {
	return s.SolutionType
}

func (s *CreateFaqRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *CreateFaqRequest) GetTagIdList() []*int64 {
	return s.TagIdList
}

func (s *CreateFaqRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateFaqRequest) SetAgentKey(v string) *CreateFaqRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateFaqRequest) SetCategoryId(v int64) *CreateFaqRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateFaqRequest) SetEndDate(v string) *CreateFaqRequest {
	s.EndDate = &v
	return s
}

func (s *CreateFaqRequest) SetSolutionContent(v string) *CreateFaqRequest {
	s.SolutionContent = &v
	return s
}

func (s *CreateFaqRequest) SetSolutionType(v int32) *CreateFaqRequest {
	s.SolutionType = &v
	return s
}

func (s *CreateFaqRequest) SetStartDate(v string) *CreateFaqRequest {
	s.StartDate = &v
	return s
}

func (s *CreateFaqRequest) SetTagIdList(v []*int64) *CreateFaqRequest {
	s.TagIdList = v
	return s
}

func (s *CreateFaqRequest) SetTitle(v string) *CreateFaqRequest {
	s.Title = &v
	return s
}

func (s *CreateFaqRequest) Validate() error {
	return dara.Validate(s)
}
