// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFaqShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SearchFaqShrinkRequest
	GetAgentKey() *string
	SetCategoryIdsShrink(v string) *SearchFaqShrinkRequest
	GetCategoryIdsShrink() *string
	SetCreateTimeBegin(v string) *SearchFaqShrinkRequest
	GetCreateTimeBegin() *string
	SetCreateTimeEnd(v string) *SearchFaqShrinkRequest
	GetCreateTimeEnd() *string
	SetCreateUserName(v string) *SearchFaqShrinkRequest
	GetCreateUserName() *string
	SetEndTimeBegin(v string) *SearchFaqShrinkRequest
	GetEndTimeBegin() *string
	SetEndTimeEnd(v string) *SearchFaqShrinkRequest
	GetEndTimeEnd() *string
	SetKeyword(v string) *SearchFaqShrinkRequest
	GetKeyword() *string
	SetModifyTimeBegin(v string) *SearchFaqShrinkRequest
	GetModifyTimeBegin() *string
	SetModifyTimeEnd(v string) *SearchFaqShrinkRequest
	GetModifyTimeEnd() *string
	SetModifyUserName(v string) *SearchFaqShrinkRequest
	GetModifyUserName() *string
	SetPageNumber(v int32) *SearchFaqShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchFaqShrinkRequest
	GetPageSize() *int32
	SetSearchScope(v int32) *SearchFaqShrinkRequest
	GetSearchScope() *int32
	SetStartTimeBegin(v string) *SearchFaqShrinkRequest
	GetStartTimeBegin() *string
	SetStartTimeEnd(v string) *SearchFaqShrinkRequest
	GetStartTimeEnd() *string
	SetStatus(v int32) *SearchFaqShrinkRequest
	GetStatus() *int32
}

type SearchFaqShrinkRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey          *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CategoryIdsShrink *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	CreateTimeBegin *string `json:"CreateTimeBegin,omitempty" xml:"CreateTimeBegin,omitempty"`
	// example:
	//
	// 2022-05-02T03:09:30Z
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// test01
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 2023-04-02T03:09:30Z
	EndTimeBegin *string `json:"EndTimeBegin,omitempty" xml:"EndTimeBegin,omitempty"`
	// example:
	//
	// 2023-05-02T03:09:30Z
	EndTimeEnd *string `json:"EndTimeEnd,omitempty" xml:"EndTimeEnd,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 2023-04-02T03:09:30Z
	ModifyTimeBegin *string `json:"ModifyTimeBegin,omitempty" xml:"ModifyTimeBegin,omitempty"`
	// example:
	//
	// 2023-05-02T03:09:30Z
	ModifyTimeEnd *string `json:"ModifyTimeEnd,omitempty" xml:"ModifyTimeEnd,omitempty"`
	// example:
	//
	// test01
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	SearchScope *int32 `json:"SearchScope,omitempty" xml:"SearchScope,omitempty"`
	// example:
	//
	// 2022-04-02T03:09:30Z
	StartTimeBegin *string `json:"StartTimeBegin,omitempty" xml:"StartTimeBegin,omitempty"`
	// example:
	//
	// 2022-04-03T03:09:30Z
	StartTimeEnd *string `json:"StartTimeEnd,omitempty" xml:"StartTimeEnd,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SearchFaqShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchFaqShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchFaqShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SearchFaqShrinkRequest) GetCategoryIdsShrink() *string {
	return s.CategoryIdsShrink
}

func (s *SearchFaqShrinkRequest) GetCreateTimeBegin() *string {
	return s.CreateTimeBegin
}

func (s *SearchFaqShrinkRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *SearchFaqShrinkRequest) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *SearchFaqShrinkRequest) GetEndTimeBegin() *string {
	return s.EndTimeBegin
}

func (s *SearchFaqShrinkRequest) GetEndTimeEnd() *string {
	return s.EndTimeEnd
}

func (s *SearchFaqShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchFaqShrinkRequest) GetModifyTimeBegin() *string {
	return s.ModifyTimeBegin
}

func (s *SearchFaqShrinkRequest) GetModifyTimeEnd() *string {
	return s.ModifyTimeEnd
}

func (s *SearchFaqShrinkRequest) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *SearchFaqShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchFaqShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchFaqShrinkRequest) GetSearchScope() *int32 {
	return s.SearchScope
}

func (s *SearchFaqShrinkRequest) GetStartTimeBegin() *string {
	return s.StartTimeBegin
}

func (s *SearchFaqShrinkRequest) GetStartTimeEnd() *string {
	return s.StartTimeEnd
}

func (s *SearchFaqShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *SearchFaqShrinkRequest) SetAgentKey(v string) *SearchFaqShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetCategoryIdsShrink(v string) *SearchFaqShrinkRequest {
	s.CategoryIdsShrink = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetCreateTimeBegin(v string) *SearchFaqShrinkRequest {
	s.CreateTimeBegin = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetCreateTimeEnd(v string) *SearchFaqShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetCreateUserName(v string) *SearchFaqShrinkRequest {
	s.CreateUserName = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetEndTimeBegin(v string) *SearchFaqShrinkRequest {
	s.EndTimeBegin = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetEndTimeEnd(v string) *SearchFaqShrinkRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetKeyword(v string) *SearchFaqShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetModifyTimeBegin(v string) *SearchFaqShrinkRequest {
	s.ModifyTimeBegin = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetModifyTimeEnd(v string) *SearchFaqShrinkRequest {
	s.ModifyTimeEnd = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetModifyUserName(v string) *SearchFaqShrinkRequest {
	s.ModifyUserName = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetPageNumber(v int32) *SearchFaqShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetPageSize(v int32) *SearchFaqShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetSearchScope(v int32) *SearchFaqShrinkRequest {
	s.SearchScope = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetStartTimeBegin(v string) *SearchFaqShrinkRequest {
	s.StartTimeBegin = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetStartTimeEnd(v string) *SearchFaqShrinkRequest {
	s.StartTimeEnd = &v
	return s
}

func (s *SearchFaqShrinkRequest) SetStatus(v int32) *SearchFaqShrinkRequest {
	s.Status = &v
	return s
}

func (s *SearchFaqShrinkRequest) Validate() error {
	return dara.Validate(s)
}
