// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFaqRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SearchFaqRequest
	GetAgentKey() *string
	SetCategoryIds(v []*int64) *SearchFaqRequest
	GetCategoryIds() []*int64
	SetCreateTimeBegin(v string) *SearchFaqRequest
	GetCreateTimeBegin() *string
	SetCreateTimeEnd(v string) *SearchFaqRequest
	GetCreateTimeEnd() *string
	SetCreateUserName(v string) *SearchFaqRequest
	GetCreateUserName() *string
	SetEndTimeBegin(v string) *SearchFaqRequest
	GetEndTimeBegin() *string
	SetEndTimeEnd(v string) *SearchFaqRequest
	GetEndTimeEnd() *string
	SetKeyword(v string) *SearchFaqRequest
	GetKeyword() *string
	SetModifyTimeBegin(v string) *SearchFaqRequest
	GetModifyTimeBegin() *string
	SetModifyTimeEnd(v string) *SearchFaqRequest
	GetModifyTimeEnd() *string
	SetModifyUserName(v string) *SearchFaqRequest
	GetModifyUserName() *string
	SetPageNumber(v int32) *SearchFaqRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchFaqRequest
	GetPageSize() *int32
	SetSearchScope(v int32) *SearchFaqRequest
	GetSearchScope() *int32
	SetStartTimeBegin(v string) *SearchFaqRequest
	GetStartTimeBegin() *string
	SetStartTimeEnd(v string) *SearchFaqRequest
	GetStartTimeEnd() *string
	SetStatus(v int32) *SearchFaqRequest
	GetStatus() *int32
}

type SearchFaqRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey    *string  `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CategoryIds []*int64 `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
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

func (s SearchFaqRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchFaqRequest) GoString() string {
	return s.String()
}

func (s *SearchFaqRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SearchFaqRequest) GetCategoryIds() []*int64 {
	return s.CategoryIds
}

func (s *SearchFaqRequest) GetCreateTimeBegin() *string {
	return s.CreateTimeBegin
}

func (s *SearchFaqRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *SearchFaqRequest) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *SearchFaqRequest) GetEndTimeBegin() *string {
	return s.EndTimeBegin
}

func (s *SearchFaqRequest) GetEndTimeEnd() *string {
	return s.EndTimeEnd
}

func (s *SearchFaqRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchFaqRequest) GetModifyTimeBegin() *string {
	return s.ModifyTimeBegin
}

func (s *SearchFaqRequest) GetModifyTimeEnd() *string {
	return s.ModifyTimeEnd
}

func (s *SearchFaqRequest) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *SearchFaqRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchFaqRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchFaqRequest) GetSearchScope() *int32 {
	return s.SearchScope
}

func (s *SearchFaqRequest) GetStartTimeBegin() *string {
	return s.StartTimeBegin
}

func (s *SearchFaqRequest) GetStartTimeEnd() *string {
	return s.StartTimeEnd
}

func (s *SearchFaqRequest) GetStatus() *int32 {
	return s.Status
}

func (s *SearchFaqRequest) SetAgentKey(v string) *SearchFaqRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchFaqRequest) SetCategoryIds(v []*int64) *SearchFaqRequest {
	s.CategoryIds = v
	return s
}

func (s *SearchFaqRequest) SetCreateTimeBegin(v string) *SearchFaqRequest {
	s.CreateTimeBegin = &v
	return s
}

func (s *SearchFaqRequest) SetCreateTimeEnd(v string) *SearchFaqRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchFaqRequest) SetCreateUserName(v string) *SearchFaqRequest {
	s.CreateUserName = &v
	return s
}

func (s *SearchFaqRequest) SetEndTimeBegin(v string) *SearchFaqRequest {
	s.EndTimeBegin = &v
	return s
}

func (s *SearchFaqRequest) SetEndTimeEnd(v string) *SearchFaqRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *SearchFaqRequest) SetKeyword(v string) *SearchFaqRequest {
	s.Keyword = &v
	return s
}

func (s *SearchFaqRequest) SetModifyTimeBegin(v string) *SearchFaqRequest {
	s.ModifyTimeBegin = &v
	return s
}

func (s *SearchFaqRequest) SetModifyTimeEnd(v string) *SearchFaqRequest {
	s.ModifyTimeEnd = &v
	return s
}

func (s *SearchFaqRequest) SetModifyUserName(v string) *SearchFaqRequest {
	s.ModifyUserName = &v
	return s
}

func (s *SearchFaqRequest) SetPageNumber(v int32) *SearchFaqRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchFaqRequest) SetPageSize(v int32) *SearchFaqRequest {
	s.PageSize = &v
	return s
}

func (s *SearchFaqRequest) SetSearchScope(v int32) *SearchFaqRequest {
	s.SearchScope = &v
	return s
}

func (s *SearchFaqRequest) SetStartTimeBegin(v string) *SearchFaqRequest {
	s.StartTimeBegin = &v
	return s
}

func (s *SearchFaqRequest) SetStartTimeEnd(v string) *SearchFaqRequest {
	s.StartTimeEnd = &v
	return s
}

func (s *SearchFaqRequest) SetStatus(v int32) *SearchFaqRequest {
	s.Status = &v
	return s
}

func (s *SearchFaqRequest) Validate() error {
	return dara.Validate(s)
}
