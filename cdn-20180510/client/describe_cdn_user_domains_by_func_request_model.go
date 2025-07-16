// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserDomainsByFuncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFuncId(v int32) *DescribeCdnUserDomainsByFuncRequest
	GetFuncId() *int32
	SetPageNumber(v int32) *DescribeCdnUserDomainsByFuncRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCdnUserDomainsByFuncRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeCdnUserDomainsByFuncRequest
	GetResourceGroupId() *string
}

type DescribeCdnUserDomainsByFuncRequest struct {
	// The ID of the feature.
	//
	// For example, the ID of the origin host feature (set_req_host_header) is 18.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	FuncId *int32 `json:"FuncId,omitempty" xml:"FuncId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// Valid values: **1*	- to **100000**.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of domain names to return on each page. Default value: **20**.
	//
	// Valid values: **1*	- to **50**.
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

func (s DescribeCdnUserDomainsByFuncRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserDomainsByFuncRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserDomainsByFuncRequest) GetFuncId() *int32 {
	return s.FuncId
}

func (s *DescribeCdnUserDomainsByFuncRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCdnUserDomainsByFuncRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCdnUserDomainsByFuncRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCdnUserDomainsByFuncRequest) SetFuncId(v int32) *DescribeCdnUserDomainsByFuncRequest {
	s.FuncId = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncRequest) SetPageNumber(v int32) *DescribeCdnUserDomainsByFuncRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncRequest) SetPageSize(v int32) *DescribeCdnUserDomainsByFuncRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncRequest) SetResourceGroupId(v string) *DescribeCdnUserDomainsByFuncRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCdnUserDomainsByFuncRequest) Validate() error {
	return dara.Validate(s)
}
