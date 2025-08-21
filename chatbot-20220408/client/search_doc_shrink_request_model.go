// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDocShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SearchDocShrinkRequest
	GetAgentKey() *string
	SetCategoryIdsShrink(v string) *SearchDocShrinkRequest
	GetCategoryIdsShrink() *string
	SetCreateTimeBegin(v string) *SearchDocShrinkRequest
	GetCreateTimeBegin() *string
	SetCreateTimeEnd(v string) *SearchDocShrinkRequest
	GetCreateTimeEnd() *string
	SetCreateUserName(v string) *SearchDocShrinkRequest
	GetCreateUserName() *string
	SetEndTimeBegin(v string) *SearchDocShrinkRequest
	GetEndTimeBegin() *string
	SetEndTimeEnd(v string) *SearchDocShrinkRequest
	GetEndTimeEnd() *string
	SetKeyword(v string) *SearchDocShrinkRequest
	GetKeyword() *string
	SetModifyTimeBegin(v string) *SearchDocShrinkRequest
	GetModifyTimeBegin() *string
	SetModifyTimeEnd(v string) *SearchDocShrinkRequest
	GetModifyTimeEnd() *string
	SetModifyUserName(v string) *SearchDocShrinkRequest
	GetModifyUserName() *string
	SetPageNumber(v int32) *SearchDocShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchDocShrinkRequest
	GetPageSize() *int32
	SetProcessStatus(v int32) *SearchDocShrinkRequest
	GetProcessStatus() *int32
	SetSearchScope(v int32) *SearchDocShrinkRequest
	GetSearchScope() *int32
	SetStartTimeBegin(v string) *SearchDocShrinkRequest
	GetStartTimeBegin() *string
	SetStartTimeEnd(v string) *SearchDocShrinkRequest
	GetStartTimeEnd() *string
	SetStatus(v int32) *SearchDocShrinkRequest
	GetStatus() *int32
	SetTagIdsShrink(v string) *SearchDocShrinkRequest
	GetTagIdsShrink() *string
}

type SearchDocShrinkRequest struct {
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
	CreateTimeEnd  *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
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
	ModifyTimeEnd  *string `json:"ModifyTimeEnd,omitempty" xml:"ModifyTimeEnd,omitempty"`
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
	// 0
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
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
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIdsShrink *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
}

func (s SearchDocShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchDocShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchDocShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SearchDocShrinkRequest) GetCategoryIdsShrink() *string {
	return s.CategoryIdsShrink
}

func (s *SearchDocShrinkRequest) GetCreateTimeBegin() *string {
	return s.CreateTimeBegin
}

func (s *SearchDocShrinkRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *SearchDocShrinkRequest) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *SearchDocShrinkRequest) GetEndTimeBegin() *string {
	return s.EndTimeBegin
}

func (s *SearchDocShrinkRequest) GetEndTimeEnd() *string {
	return s.EndTimeEnd
}

func (s *SearchDocShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchDocShrinkRequest) GetModifyTimeBegin() *string {
	return s.ModifyTimeBegin
}

func (s *SearchDocShrinkRequest) GetModifyTimeEnd() *string {
	return s.ModifyTimeEnd
}

func (s *SearchDocShrinkRequest) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *SearchDocShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchDocShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchDocShrinkRequest) GetProcessStatus() *int32 {
	return s.ProcessStatus
}

func (s *SearchDocShrinkRequest) GetSearchScope() *int32 {
	return s.SearchScope
}

func (s *SearchDocShrinkRequest) GetStartTimeBegin() *string {
	return s.StartTimeBegin
}

func (s *SearchDocShrinkRequest) GetStartTimeEnd() *string {
	return s.StartTimeEnd
}

func (s *SearchDocShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *SearchDocShrinkRequest) GetTagIdsShrink() *string {
	return s.TagIdsShrink
}

func (s *SearchDocShrinkRequest) SetAgentKey(v string) *SearchDocShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchDocShrinkRequest) SetCategoryIdsShrink(v string) *SearchDocShrinkRequest {
	s.CategoryIdsShrink = &v
	return s
}

func (s *SearchDocShrinkRequest) SetCreateTimeBegin(v string) *SearchDocShrinkRequest {
	s.CreateTimeBegin = &v
	return s
}

func (s *SearchDocShrinkRequest) SetCreateTimeEnd(v string) *SearchDocShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchDocShrinkRequest) SetCreateUserName(v string) *SearchDocShrinkRequest {
	s.CreateUserName = &v
	return s
}

func (s *SearchDocShrinkRequest) SetEndTimeBegin(v string) *SearchDocShrinkRequest {
	s.EndTimeBegin = &v
	return s
}

func (s *SearchDocShrinkRequest) SetEndTimeEnd(v string) *SearchDocShrinkRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *SearchDocShrinkRequest) SetKeyword(v string) *SearchDocShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *SearchDocShrinkRequest) SetModifyTimeBegin(v string) *SearchDocShrinkRequest {
	s.ModifyTimeBegin = &v
	return s
}

func (s *SearchDocShrinkRequest) SetModifyTimeEnd(v string) *SearchDocShrinkRequest {
	s.ModifyTimeEnd = &v
	return s
}

func (s *SearchDocShrinkRequest) SetModifyUserName(v string) *SearchDocShrinkRequest {
	s.ModifyUserName = &v
	return s
}

func (s *SearchDocShrinkRequest) SetPageNumber(v int32) *SearchDocShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchDocShrinkRequest) SetPageSize(v int32) *SearchDocShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchDocShrinkRequest) SetProcessStatus(v int32) *SearchDocShrinkRequest {
	s.ProcessStatus = &v
	return s
}

func (s *SearchDocShrinkRequest) SetSearchScope(v int32) *SearchDocShrinkRequest {
	s.SearchScope = &v
	return s
}

func (s *SearchDocShrinkRequest) SetStartTimeBegin(v string) *SearchDocShrinkRequest {
	s.StartTimeBegin = &v
	return s
}

func (s *SearchDocShrinkRequest) SetStartTimeEnd(v string) *SearchDocShrinkRequest {
	s.StartTimeEnd = &v
	return s
}

func (s *SearchDocShrinkRequest) SetStatus(v int32) *SearchDocShrinkRequest {
	s.Status = &v
	return s
}

func (s *SearchDocShrinkRequest) SetTagIdsShrink(v string) *SearchDocShrinkRequest {
	s.TagIdsShrink = &v
	return s
}

func (s *SearchDocShrinkRequest) Validate() error {
	return dara.Validate(s)
}
