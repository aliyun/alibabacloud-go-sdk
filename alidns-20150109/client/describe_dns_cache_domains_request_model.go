// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsCacheDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeDnsCacheDomainsRequest
	GetKeyword() *string
	SetLang(v string) *DescribeDnsCacheDomainsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeDnsCacheDomainsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDnsCacheDomainsRequest
	GetPageSize() *int64
}

type DescribeDnsCacheDomainsRequest struct {
	// The keyword for searches in "%KeyWord%" mode. The value is not case-sensitive.
	//
	// example:
	//
	// a\\"\\"
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDnsCacheDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsCacheDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsCacheDomainsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeDnsCacheDomainsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsCacheDomainsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDnsCacheDomainsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDnsCacheDomainsRequest) SetKeyword(v string) *DescribeDnsCacheDomainsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDnsCacheDomainsRequest) SetLang(v string) *DescribeDnsCacheDomainsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsCacheDomainsRequest) SetPageNumber(v int64) *DescribeDnsCacheDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsCacheDomainsRequest) SetPageSize(v int64) *DescribeDnsCacheDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsCacheDomainsRequest) Validate() error {
	return dara.Validate(s)
}
