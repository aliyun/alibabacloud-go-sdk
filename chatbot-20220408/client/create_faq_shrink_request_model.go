// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFaqShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateFaqShrinkRequest
	GetAgentKey() *string
	SetCategoryId(v int64) *CreateFaqShrinkRequest
	GetCategoryId() *int64
	SetEndDate(v string) *CreateFaqShrinkRequest
	GetEndDate() *string
	SetSolutionContent(v string) *CreateFaqShrinkRequest
	GetSolutionContent() *string
	SetSolutionType(v int32) *CreateFaqShrinkRequest
	GetSolutionType() *int32
	SetStartDate(v string) *CreateFaqShrinkRequest
	GetStartDate() *string
	SetTagIdListShrink(v string) *CreateFaqShrinkRequest
	GetTagIdListShrink() *string
	SetTitle(v string) *CreateFaqShrinkRequest
	GetTitle() *string
}

type CreateFaqShrinkRequest struct {
	// The agent key. If omitted, the default agent is used. Find this key on the Agent Management page.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The ID of the knowledge category.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000053274
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The knowledge end time. The time is in UTC and in ISO 8601 format.
	//
	// example:
	//
	// 2030-12-31T16:00:00Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The content of the default solution. Required if the fallback feature is enabled.
	//
	// example:
	//
	// 测试答案
	SolutionContent *string `json:"SolutionContent,omitempty" xml:"SolutionContent,omitempty"`
	// The type of the default solution. Valid values: `0` (plain text) and `1` (rich text).
	//
	// example:
	//
	// 0
	SolutionType *int32 `json:"SolutionType,omitempty" xml:"SolutionType,omitempty"`
	// The knowledge start time. The time is in UTC and in ISO 8601 format.
	//
	// example:
	//
	// 2022-05-25T16:28:36Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// A list of tag IDs to associate with the knowledge.
	TagIdListShrink *string `json:"TagIdList,omitempty" xml:"TagIdList,omitempty"`
	// The knowledge title. Max length: 120 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateFaqShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFaqShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFaqShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateFaqShrinkRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *CreateFaqShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateFaqShrinkRequest) GetSolutionContent() *string {
	return s.SolutionContent
}

func (s *CreateFaqShrinkRequest) GetSolutionType() *int32 {
	return s.SolutionType
}

func (s *CreateFaqShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *CreateFaqShrinkRequest) GetTagIdListShrink() *string {
	return s.TagIdListShrink
}

func (s *CreateFaqShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateFaqShrinkRequest) SetAgentKey(v string) *CreateFaqShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateFaqShrinkRequest) SetCategoryId(v int64) *CreateFaqShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateFaqShrinkRequest) SetEndDate(v string) *CreateFaqShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *CreateFaqShrinkRequest) SetSolutionContent(v string) *CreateFaqShrinkRequest {
	s.SolutionContent = &v
	return s
}

func (s *CreateFaqShrinkRequest) SetSolutionType(v int32) *CreateFaqShrinkRequest {
	s.SolutionType = &v
	return s
}

func (s *CreateFaqShrinkRequest) SetStartDate(v string) *CreateFaqShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *CreateFaqShrinkRequest) SetTagIdListShrink(v string) *CreateFaqShrinkRequest {
	s.TagIdListShrink = &v
	return s
}

func (s *CreateFaqShrinkRequest) SetTitle(v string) *CreateFaqShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateFaqShrinkRequest) Validate() error {
	return dara.Validate(s)
}
