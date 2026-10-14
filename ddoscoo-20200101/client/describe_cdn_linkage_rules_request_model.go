// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnLinkageRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeCdnLinkageRulesRequest
	GetDomain() *string
	SetPageNumber(v int32) *DescribeCdnLinkageRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCdnLinkageRulesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeCdnLinkageRulesRequest
	GetResourceGroupId() *string
}

type DescribeCdnLinkageRulesRequest struct {
	// The domain name of your website.
	//
	// > The domain name must be added to your Anti-DDoS instance. Call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of the page to return for a paged query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group to which the Anti-DDoS instance belongs in Resource Management. If you leave this parameter empty, the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeCdnLinkageRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnLinkageRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnLinkageRulesRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCdnLinkageRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCdnLinkageRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCdnLinkageRulesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCdnLinkageRulesRequest) SetDomain(v string) *DescribeCdnLinkageRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeCdnLinkageRulesRequest) SetPageNumber(v int32) *DescribeCdnLinkageRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnLinkageRulesRequest) SetPageSize(v int32) *DescribeCdnLinkageRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnLinkageRulesRequest) SetResourceGroupId(v string) *DescribeCdnLinkageRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCdnLinkageRulesRequest) Validate() error {
	return dara.Validate(s)
}
