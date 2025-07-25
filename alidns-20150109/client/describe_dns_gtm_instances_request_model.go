// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeDnsGtmInstancesRequest
	GetKeyword() *string
	SetLang(v string) *DescribeDnsGtmInstancesRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeDnsGtmInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDnsGtmInstancesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeDnsGtmInstancesRequest
	GetResourceGroupId() *string
}

type DescribeDnsGtmInstancesRequest struct {
	// The keyword that you use for the query. Fuzzy search by instance ID or instance name is supported.
	//
	// example:
	//
	// instance1
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of the values for specific response parameters. Default value: en. Valid values: en, zh, and ja.
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **100**. Default value: **20**.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-testgroupid
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDnsGtmInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeDnsGtmInstancesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDnsGtmInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDnsGtmInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDnsGtmInstancesRequest) SetKeyword(v string) *DescribeDnsGtmInstancesRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDnsGtmInstancesRequest) SetLang(v string) *DescribeDnsGtmInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmInstancesRequest) SetPageNumber(v int32) *DescribeDnsGtmInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmInstancesRequest) SetPageSize(v int32) *DescribeDnsGtmInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmInstancesRequest) SetResourceGroupId(v string) *DescribeDnsGtmInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDnsGtmInstancesRequest) Validate() error {
	return dara.Validate(s)
}
