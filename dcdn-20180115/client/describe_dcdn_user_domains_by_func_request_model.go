// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserDomainsByFuncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnUserDomainsByFuncRequest
	GetDomainName() *string
	SetFuncFilter(v string) *DescribeDcdnUserDomainsByFuncRequest
	GetFuncFilter() *string
	SetFuncId(v int32) *DescribeDcdnUserDomainsByFuncRequest
	GetFuncId() *int32
	SetMatchType(v string) *DescribeDcdnUserDomainsByFuncRequest
	GetMatchType() *string
	SetPageNumber(v int32) *DescribeDcdnUserDomainsByFuncRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnUserDomainsByFuncRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeDcdnUserDomainsByFuncRequest
	GetResourceGroupId() *string
}

type DescribeDcdnUserDomainsByFuncRequest struct {
	// The accelerated domain name whose ICP filing status you want to update.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether the feature that is specified by the FuncId parameter is enabled.
	//
	// 	- **config**: enabled
	//
	// 	- **unconfig**: not enabled
	//
	// example:
	//
	// config
	FuncFilter *string `json:"FuncFilter,omitempty" xml:"FuncFilter,omitempty"`
	// The ID of the feature. For more information about how to query feature IDs, see [Parameters for configuring features for domain names](https://help.aliyun.com/document_detail/410622.html). For example, the ID of the origin host feature (set_req_host_header) is 18.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	FuncId *int32 `json:"FuncId,omitempty" xml:"FuncId,omitempty"`
	// The type of the search. Default value: exact_match. Valid values:
	//
	// 	- fuzzy_match: fuzzy search.
	//
	// 	- exact_match: exact search.
	//
	// example:
	//
	// exact_match
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The number of the page to return. Default value: **1**. Valid values: **1 to 100000**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **20**. Valid values: **1 to 500**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyuji4b6r4**
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDcdnUserDomainsByFuncRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnUserDomainsByFuncRequest) GetFuncFilter() *string {
	return s.FuncFilter
}

func (s *DescribeDcdnUserDomainsByFuncRequest) GetFuncId() *int32 {
	return s.FuncId
}

func (s *DescribeDcdnUserDomainsByFuncRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *DescribeDcdnUserDomainsByFuncRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnUserDomainsByFuncRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnUserDomainsByFuncRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetDomainName(v string) *DescribeDcdnUserDomainsByFuncRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetFuncFilter(v string) *DescribeDcdnUserDomainsByFuncRequest {
	s.FuncFilter = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetFuncId(v int32) *DescribeDcdnUserDomainsByFuncRequest {
	s.FuncId = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetMatchType(v string) *DescribeDcdnUserDomainsByFuncRequest {
	s.MatchType = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetPageNumber(v int32) *DescribeDcdnUserDomainsByFuncRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetPageSize(v int32) *DescribeDcdnUserDomainsByFuncRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) SetResourceGroupId(v string) *DescribeDcdnUserDomainsByFuncRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncRequest) Validate() error {
	return dara.Validate(s)
}
