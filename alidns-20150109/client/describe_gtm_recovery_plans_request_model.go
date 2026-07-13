// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmRecoveryPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeGtmRecoveryPlansRequest
	GetKeyword() *string
	SetLang(v string) *DescribeGtmRecoveryPlansRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeGtmRecoveryPlansRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGtmRecoveryPlansRequest
	GetPageSize() *int32
}

type DescribeGtmRecoveryPlansRequest struct {
	// The keyword for the query. This parameter supports a fuzzy search by disaster recovery plan name.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of the response. Valid values are `en` for English and `zh` for Chinese. The default value is `zh`.
	//
	// en: English.
	//
	// en: English
	//
	// Default value: zh.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return. Pages start from **1**. The default value is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is **100**. The default value is **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGtmRecoveryPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmRecoveryPlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmRecoveryPlansRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeGtmRecoveryPlansRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmRecoveryPlansRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGtmRecoveryPlansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGtmRecoveryPlansRequest) SetKeyword(v string) *DescribeGtmRecoveryPlansRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeGtmRecoveryPlansRequest) SetLang(v string) *DescribeGtmRecoveryPlansRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmRecoveryPlansRequest) SetPageNumber(v int32) *DescribeGtmRecoveryPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmRecoveryPlansRequest) SetPageSize(v int32) *DescribeGtmRecoveryPlansRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmRecoveryPlansRequest) Validate() error {
	return dara.Validate(s)
}
