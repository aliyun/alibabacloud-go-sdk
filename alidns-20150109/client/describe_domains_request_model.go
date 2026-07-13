// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeDomainsRequest
	GetGroupId() *string
	SetKeyWord(v string) *DescribeDomainsRequest
	GetKeyWord() *string
	SetLang(v string) *DescribeDomainsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeDomainsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainsRequest
	GetPageSize() *int64
	SetResourceGroupId(v string) *DescribeDomainsRequest
	GetResourceGroupId() *string
	SetSearchMode(v string) *DescribeDomainsRequest
	GetSearchMode() *string
	SetStarmark(v bool) *DescribeDomainsRequest
	GetStarmark() *bool
}

type DescribeDomainsRequest struct {
	// The ID of the domain name group.
	//
	// If you leave this parameter empty or pass an empty string, all domain names are queried.
	//
	// If you set this parameter to defaultGroup, domain names in the default group are queried.
	//
	// example:
	//
	// defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The keyword. The search is performed in the %KeyWord% pattern and is not case-sensitive.
	//
	// example:
	//
	// com
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// Default value: zh.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The value starts from **1**. The default value is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The maximum value is **100**. The default value is **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-re********
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The search mode. Valid values:
	//
	// - **LIKE**: fuzzy search
	//
	// - **EXACT**: exact search
	//
	// Default value: LIKE
	//
	// example:
	//
	// LIKE
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// Specifies whether to query starred domain names. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// Default value: true
	//
	// example:
	//
	// true
	Starmark *bool `json:"Starmark,omitempty" xml:"Starmark,omitempty"`
}

func (s DescribeDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDomainsRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *DescribeDomainsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainsRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *DescribeDomainsRequest) GetStarmark() *bool {
	return s.Starmark
}

func (s *DescribeDomainsRequest) SetGroupId(v string) *DescribeDomainsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainsRequest) SetKeyWord(v string) *DescribeDomainsRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeDomainsRequest) SetLang(v string) *DescribeDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageNumber(v int64) *DescribeDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainsRequest) SetPageSize(v int64) *DescribeDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsRequest) SetResourceGroupId(v string) *DescribeDomainsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainsRequest) SetSearchMode(v string) *DescribeDomainsRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeDomainsRequest) SetStarmark(v bool) *DescribeDomainsRequest {
	s.Starmark = &v
	return s
}

func (s *DescribeDomainsRequest) Validate() error {
	return dara.Validate(s)
}
