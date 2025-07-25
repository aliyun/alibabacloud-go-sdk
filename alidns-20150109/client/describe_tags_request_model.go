// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeTagsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeTagsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeTagsRequest
	GetPageSize() *int64
	SetResourceType(v string) *DescribeTagsRequest
	GetResourceType() *string
}

type DescribeTagsRequest struct {
	// The language in which you want the values of some response parameters to be returned. These response parameters support multiple languages. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Default value: 200.
	//
	// example:
	//
	// 200
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- **DOMAIN**: domain name
	//
	// This parameter is required.
	//
	// example:
	//
	// DOMAIN
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTagsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeTagsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagsRequest) SetLang(v string) *DescribeTagsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTagsRequest) SetPageNumber(v int64) *DescribeTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTagsRequest) SetPageSize(v int64) *DescribeTagsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceType(v string) *DescribeTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagsRequest) Validate() error {
	return dara.Validate(s)
}
