// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListSkillRequest
	GetLanguage() *string
	SetPageNumber(v int64) *ListSkillRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSkillRequest
	GetPageSize() *int64
	SetQuery(v string) *ListSkillRequest
	GetQuery() *string
	SetScope(v string) *ListSkillRequest
	GetScope() *string
	SetSort(v string) *ListSkillRequest
	GetSort() *string
	SetSource(v string) *ListSkillRequest
	GetSource() *string
	SetWorkspaceId(v string) *ListSkillRequest
	GetWorkspaceId() *string
}

type ListSkillRequest struct {
	// The supported languages. Valid values:
	//
	// 	- zh-CN: Simplified Chinese
	//
	// 	- zh-TW: Traditional Chinese
	//
	// 	- en-US: English
	//
	// 	- ja-JP: Japanese
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search parameter for skill name or description. This filter is not currently enabled in the ContextDB management console.
	//
	// example:
	//
	// example
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ContextDB query scope. Valid values: all (default), private, and system.
	//
	// example:
	//
	// all
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The sort parameter for skills. This sorting is not currently enabled in the ContextDB management console.
	//
	// example:
	//
	// updated_at_desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The filter parameter for skill source. This filter is not currently enabled in the ContextDB management console.
	//
	// example:
	//
	// SYSTEM
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ContextDB workspace ID.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillRequest) GoString() string {
	return s.String()
}

func (s *ListSkillRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListSkillRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSkillRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSkillRequest) GetQuery() *string {
	return s.Query
}

func (s *ListSkillRequest) GetScope() *string {
	return s.Scope
}

func (s *ListSkillRequest) GetSort() *string {
	return s.Sort
}

func (s *ListSkillRequest) GetSource() *string {
	return s.Source
}

func (s *ListSkillRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListSkillRequest) SetLanguage(v string) *ListSkillRequest {
	s.Language = &v
	return s
}

func (s *ListSkillRequest) SetPageNumber(v int64) *ListSkillRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSkillRequest) SetPageSize(v int64) *ListSkillRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillRequest) SetQuery(v string) *ListSkillRequest {
	s.Query = &v
	return s
}

func (s *ListSkillRequest) SetScope(v string) *ListSkillRequest {
	s.Scope = &v
	return s
}

func (s *ListSkillRequest) SetSort(v string) *ListSkillRequest {
	s.Sort = &v
	return s
}

func (s *ListSkillRequest) SetSource(v string) *ListSkillRequest {
	s.Source = &v
	return s
}

func (s *ListSkillRequest) SetWorkspaceId(v string) *ListSkillRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListSkillRequest) Validate() error {
	return dara.Validate(s)
}
