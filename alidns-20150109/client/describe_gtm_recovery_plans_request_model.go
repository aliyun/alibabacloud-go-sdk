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
	// The keyword for the query. Fuzzy match is supported by disaster recovery plan name.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language in which you want the values of some response parameters to be returned. These response parameters support multiple languages.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number to return. The page number starts from **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Maximum value: **100**. Default value: **20**.
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
