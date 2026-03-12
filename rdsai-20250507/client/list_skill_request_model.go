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
}

type ListSkillRequest struct {
	// The languages supported by the skills.
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
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records to return on each page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListSkillRequest) Validate() error {
	return dara.Validate(s)
}
