// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeSkillsRequest
	GetKeyword() *string
	SetLanguage(v string) *DescribeSkillsRequest
	GetLanguage() *string
	SetPageNumber(v int32) *DescribeSkillsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSkillsRequest
	GetPageSize() *int32
	SetSkillId(v string) *DescribeSkillsRequest
	GetSkillId() *string
	SetStatusFilter(v string) *DescribeSkillsRequest
	GetStatusFilter() *string
	SetType(v string) *DescribeSkillsRequest
	GetType() *string
}

type DescribeSkillsRequest struct {
	// The keyword in the skill name or skill description.
	//
	// example:
	//
	// weather
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language type. The skill description is returned in this language.
	//
	// Valid values:
	//
	// - en: English.
	//
	// - zh-CN: Chinese.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The skill ID.
	//
	// example:
	//
	// sk-051j4pbwxzgol****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The status filter.
	//
	// example:
	//
	// UPLOADED
	StatusFilter *string `json:"StatusFilter,omitempty" xml:"StatusFilter,omitempty"`
	// The skill type.
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSkillsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSkillsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeSkillsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeSkillsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSkillsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSkillsRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *DescribeSkillsRequest) GetStatusFilter() *string {
	return s.StatusFilter
}

func (s *DescribeSkillsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeSkillsRequest) SetKeyword(v string) *DescribeSkillsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeSkillsRequest) SetLanguage(v string) *DescribeSkillsRequest {
	s.Language = &v
	return s
}

func (s *DescribeSkillsRequest) SetPageNumber(v int32) *DescribeSkillsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSkillsRequest) SetPageSize(v int32) *DescribeSkillsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSkillsRequest) SetSkillId(v string) *DescribeSkillsRequest {
	s.SkillId = &v
	return s
}

func (s *DescribeSkillsRequest) SetStatusFilter(v string) *DescribeSkillsRequest {
	s.StatusFilter = &v
	return s
}

func (s *DescribeSkillsRequest) SetType(v string) *DescribeSkillsRequest {
	s.Type = &v
	return s
}

func (s *DescribeSkillsRequest) Validate() error {
	return dara.Validate(s)
}
