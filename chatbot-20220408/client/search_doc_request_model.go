// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SearchDocRequest
	GetAgentKey() *string
	SetCategoryIds(v []*int64) *SearchDocRequest
	GetCategoryIds() []*int64
	SetCreateTimeBegin(v string) *SearchDocRequest
	GetCreateTimeBegin() *string
	SetCreateTimeEnd(v string) *SearchDocRequest
	GetCreateTimeEnd() *string
	SetCreateUserName(v string) *SearchDocRequest
	GetCreateUserName() *string
	SetEndTimeBegin(v string) *SearchDocRequest
	GetEndTimeBegin() *string
	SetEndTimeEnd(v string) *SearchDocRequest
	GetEndTimeEnd() *string
	SetKeyword(v string) *SearchDocRequest
	GetKeyword() *string
	SetModifyTimeBegin(v string) *SearchDocRequest
	GetModifyTimeBegin() *string
	SetModifyTimeEnd(v string) *SearchDocRequest
	GetModifyTimeEnd() *string
	SetModifyUserName(v string) *SearchDocRequest
	GetModifyUserName() *string
	SetPageNumber(v int32) *SearchDocRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchDocRequest
	GetPageSize() *int32
	SetProcessStatus(v int32) *SearchDocRequest
	GetProcessStatus() *int32
	SetSearchScope(v int32) *SearchDocRequest
	GetSearchScope() *int32
	SetStartTimeBegin(v string) *SearchDocRequest
	GetStartTimeBegin() *string
	SetStartTimeEnd(v string) *SearchDocRequest
	GetStartTimeEnd() *string
	SetStatus(v int32) *SearchDocRequest
	GetStatus() *int32
	SetTagIds(v []*int64) *SearchDocRequest
	GetTagIds() []*int64
}

type SearchDocRequest struct {
	// The workspace key. If omitted, the default workspace is used. You can obtain this key from the Business Management page of your main account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// A list of category IDs.
	CategoryIds []*int64 `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	// The start of the creation time range.
	//
	// example:
	//
	// 2022-04-02T03:09:30Z
	CreateTimeBegin *string `json:"CreateTimeBegin,omitempty" xml:"CreateTimeBegin,omitempty"`
	// The end of the creation time range.
	//
	// example:
	//
	// 2022-05-02T03:09:30Z
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The name of the creator.
	//
	// example:
	//
	// 张三
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// The start of the expiration time range.
	//
	// example:
	//
	// 2023-04-02T03:09:30Z
	EndTimeBegin *string `json:"EndTimeBegin,omitempty" xml:"EndTimeBegin,omitempty"`
	// The end of the expiration time range.
	//
	// example:
	//
	// 2023-05-02T03:09:30Z
	EndTimeEnd *string `json:"EndTimeEnd,omitempty" xml:"EndTimeEnd,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// 搜索关键字
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The start of the modification time range.
	//
	// example:
	//
	// 2023-04-02T03:09:30Z
	ModifyTimeBegin *string `json:"ModifyTimeBegin,omitempty" xml:"ModifyTimeBegin,omitempty"`
	// The end of the modification time range.
	//
	// example:
	//
	// 2023-05-02T03:09:30Z
	ModifyTimeEnd *string `json:"ModifyTimeEnd,omitempty" xml:"ModifyTimeEnd,omitempty"`
	// The name of the last modifier.
	//
	// example:
	//
	// 李四
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// The page number to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The task status.<br>Valid values:<br>• -1: Queued<br>• 0: Succeeded<br>• 1: Parsing<br>• 2: Processing<br>• 3: Failed<br><br><br><br><br><br>
	//
	// example:
	//
	// 0
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// The search scope.<br>Valid values:<br>• 1: Search titles<br>• 2: Search content<br><br><br>
	//
	// example:
	//
	// 1
	SearchScope *int32 `json:"SearchScope,omitempty" xml:"SearchScope,omitempty"`
	// The start of the effective time range.
	//
	// example:
	//
	// 2022-04-02T03:09:30Z
	StartTimeBegin *string `json:"StartTimeBegin,omitempty" xml:"StartTimeBegin,omitempty"`
	// The end of the effective time range.
	//
	// example:
	//
	// 2022-04-03T03:09:30Z
	StartTimeEnd *string `json:"StartTimeEnd,omitempty" xml:"StartTimeEnd,omitempty"`
	// The edit status.<br>Valid values:<br>• 1: Unpublished<br>• 2: Published<br>• 3: Updated but not published<br><br><br><br>
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// A list of tag IDs. If you provide multiple IDs, the operation returns documents that match at least one of the specified tags (logical OR).
	TagIds []*int64 `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s SearchDocRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchDocRequest) GoString() string {
	return s.String()
}

func (s *SearchDocRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SearchDocRequest) GetCategoryIds() []*int64 {
	return s.CategoryIds
}

func (s *SearchDocRequest) GetCreateTimeBegin() *string {
	return s.CreateTimeBegin
}

func (s *SearchDocRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *SearchDocRequest) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *SearchDocRequest) GetEndTimeBegin() *string {
	return s.EndTimeBegin
}

func (s *SearchDocRequest) GetEndTimeEnd() *string {
	return s.EndTimeEnd
}

func (s *SearchDocRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchDocRequest) GetModifyTimeBegin() *string {
	return s.ModifyTimeBegin
}

func (s *SearchDocRequest) GetModifyTimeEnd() *string {
	return s.ModifyTimeEnd
}

func (s *SearchDocRequest) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *SearchDocRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchDocRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchDocRequest) GetProcessStatus() *int32 {
	return s.ProcessStatus
}

func (s *SearchDocRequest) GetSearchScope() *int32 {
	return s.SearchScope
}

func (s *SearchDocRequest) GetStartTimeBegin() *string {
	return s.StartTimeBegin
}

func (s *SearchDocRequest) GetStartTimeEnd() *string {
	return s.StartTimeEnd
}

func (s *SearchDocRequest) GetStatus() *int32 {
	return s.Status
}

func (s *SearchDocRequest) GetTagIds() []*int64 {
	return s.TagIds
}

func (s *SearchDocRequest) SetAgentKey(v string) *SearchDocRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchDocRequest) SetCategoryIds(v []*int64) *SearchDocRequest {
	s.CategoryIds = v
	return s
}

func (s *SearchDocRequest) SetCreateTimeBegin(v string) *SearchDocRequest {
	s.CreateTimeBegin = &v
	return s
}

func (s *SearchDocRequest) SetCreateTimeEnd(v string) *SearchDocRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *SearchDocRequest) SetCreateUserName(v string) *SearchDocRequest {
	s.CreateUserName = &v
	return s
}

func (s *SearchDocRequest) SetEndTimeBegin(v string) *SearchDocRequest {
	s.EndTimeBegin = &v
	return s
}

func (s *SearchDocRequest) SetEndTimeEnd(v string) *SearchDocRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *SearchDocRequest) SetKeyword(v string) *SearchDocRequest {
	s.Keyword = &v
	return s
}

func (s *SearchDocRequest) SetModifyTimeBegin(v string) *SearchDocRequest {
	s.ModifyTimeBegin = &v
	return s
}

func (s *SearchDocRequest) SetModifyTimeEnd(v string) *SearchDocRequest {
	s.ModifyTimeEnd = &v
	return s
}

func (s *SearchDocRequest) SetModifyUserName(v string) *SearchDocRequest {
	s.ModifyUserName = &v
	return s
}

func (s *SearchDocRequest) SetPageNumber(v int32) *SearchDocRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchDocRequest) SetPageSize(v int32) *SearchDocRequest {
	s.PageSize = &v
	return s
}

func (s *SearchDocRequest) SetProcessStatus(v int32) *SearchDocRequest {
	s.ProcessStatus = &v
	return s
}

func (s *SearchDocRequest) SetSearchScope(v int32) *SearchDocRequest {
	s.SearchScope = &v
	return s
}

func (s *SearchDocRequest) SetStartTimeBegin(v string) *SearchDocRequest {
	s.StartTimeBegin = &v
	return s
}

func (s *SearchDocRequest) SetStartTimeEnd(v string) *SearchDocRequest {
	s.StartTimeEnd = &v
	return s
}

func (s *SearchDocRequest) SetStatus(v int32) *SearchDocRequest {
	s.Status = &v
	return s
}

func (s *SearchDocRequest) SetTagIds(v []*int64) *SearchDocRequest {
	s.TagIds = v
	return s
}

func (s *SearchDocRequest) Validate() error {
	return dara.Validate(s)
}
