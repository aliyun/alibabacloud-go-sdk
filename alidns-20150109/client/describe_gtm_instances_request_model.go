// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeGtmInstancesRequest
	GetKeyword() *string
	SetLang(v string) *DescribeGtmInstancesRequest
	GetLang() *string
	SetNeedDetailAttributes(v bool) *DescribeGtmInstancesRequest
	GetNeedDetailAttributes() *bool
	SetPageNumber(v int32) *DescribeGtmInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGtmInstancesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeGtmInstancesRequest
	GetResourceGroupId() *string
}

type DescribeGtmInstancesRequest struct {
	// The keyword that you use for query. Exact match is supported by instance ID or instance name.
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
	// Specifies whether additional information is required. Default value: **false**.
	//
	// example:
	//
	// false
	NeedDetailAttributes *bool `json:"NeedDetailAttributes,omitempty" xml:"NeedDetailAttributes,omitempty"`
	// The page number to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeGtmInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstancesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeGtmInstancesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmInstancesRequest) GetNeedDetailAttributes() *bool {
	return s.NeedDetailAttributes
}

func (s *DescribeGtmInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGtmInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGtmInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGtmInstancesRequest) SetKeyword(v string) *DescribeGtmInstancesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeGtmInstancesRequest) SetLang(v string) *DescribeGtmInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmInstancesRequest) SetNeedDetailAttributes(v bool) *DescribeGtmInstancesRequest {
	s.NeedDetailAttributes = &v
	return s
}

func (s *DescribeGtmInstancesRequest) SetPageNumber(v int32) *DescribeGtmInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmInstancesRequest) SetPageSize(v int32) *DescribeGtmInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmInstancesRequest) SetResourceGroupId(v string) *DescribeGtmInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGtmInstancesRequest) Validate() error {
	return dara.Validate(s)
}
