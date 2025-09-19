// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyTypeScaItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePropertyTypeScaItemRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribePropertyTypeScaItemRequest
	GetLang() *string
	SetPageSize(v int32) *DescribePropertyTypeScaItemRequest
	GetPageSize() *int32
}

type DescribePropertyTypeScaItemRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePropertyTypeScaItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyTypeScaItemRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyTypeScaItemRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyTypeScaItemRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePropertyTypeScaItemRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyTypeScaItemRequest) SetCurrentPage(v int32) *DescribePropertyTypeScaItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyTypeScaItemRequest) SetLang(v string) *DescribePropertyTypeScaItemRequest {
	s.Lang = &v
	return s
}

func (s *DescribePropertyTypeScaItemRequest) SetPageSize(v int32) *DescribePropertyTypeScaItemRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyTypeScaItemRequest) Validate() error {
	return dara.Validate(s)
}
