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

func (s *ListSkillRequest) SetWorkspaceId(v string) *ListSkillRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListSkillRequest) Validate() error {
	return dara.Validate(s)
}
