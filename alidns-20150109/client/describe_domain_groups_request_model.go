// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyWord(v string) *DescribeDomainGroupsRequest
	GetKeyWord() *string
	SetLang(v string) *DescribeDomainGroupsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeDomainGroupsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainGroupsRequest
	GetPageSize() *int64
}

type DescribeDomainGroupsRequest struct {
	// The keyword for the group name. The search uses the %KeyWord% pattern and is case-insensitive.
	//
	// example:
	//
	// Group
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// Default value: zh
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The start value is **1**. The default value is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is **100**. The default value is **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDomainGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainGroupsRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *DescribeDomainGroupsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainGroupsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainGroupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainGroupsRequest) SetKeyWord(v string) *DescribeDomainGroupsRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeDomainGroupsRequest) SetLang(v string) *DescribeDomainGroupsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainGroupsRequest) SetPageNumber(v int64) *DescribeDomainGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainGroupsRequest) SetPageSize(v int64) *DescribeDomainGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainGroupsRequest) Validate() error {
	return dara.Validate(s)
}
