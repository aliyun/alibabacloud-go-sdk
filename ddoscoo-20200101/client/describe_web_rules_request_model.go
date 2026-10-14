// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCname(v string) *DescribeWebRulesRequest
	GetCname() *string
	SetDomain(v string) *DescribeWebRulesRequest
	GetDomain() *string
	SetInstanceIds(v []*string) *DescribeWebRulesRequest
	GetInstanceIds() []*string
	SetPageNumber(v int32) *DescribeWebRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeWebRulesRequest
	GetPageSize() *int32
	SetQueryDomainPattern(v string) *DescribeWebRulesRequest
	GetQueryDomainPattern() *string
	SetResourceGroupId(v string) *DescribeWebRulesRequest
	GetResourceGroupId() *string
}

type DescribeWebRulesRequest struct {
	// The CNAME address to query.
	//
	// example:
	//
	// kzmk7b8tt351****.aliyunddos1014****
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain name of the website to query.
	//
	// > The domain name must have a website forwarding rule configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names that have website forwarding rules configured.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IDs of the Anti-DDoS Pro or Anti-DDoS Premium instances to query.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The page number when you use paging. Set the current page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of forwarding rules per page when you use paging. Set the number of entries per page. Valid values: **1*	- to **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query match mode. Valid values:
	//
	// - **fuzzy*	- (default): fuzzy match.
	//
	// - **exact**: exact match.
	//
	// example:
	//
	// exact
	QueryDomainPattern *string `json:"QueryDomainPattern,omitempty" xml:"QueryDomainPattern,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Pro or Anti-DDoS Premium instance belongs in Resource Management.
	//
	// If you do not set this parameter to a value, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWebRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebRulesRequest) GetCname() *string {
	return s.Cname
}

func (s *DescribeWebRulesRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebRulesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeWebRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeWebRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebRulesRequest) GetQueryDomainPattern() *string {
	return s.QueryDomainPattern
}

func (s *DescribeWebRulesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWebRulesRequest) SetCname(v string) *DescribeWebRulesRequest {
	s.Cname = &v
	return s
}

func (s *DescribeWebRulesRequest) SetDomain(v string) *DescribeWebRulesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeWebRulesRequest) SetInstanceIds(v []*string) *DescribeWebRulesRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeWebRulesRequest) SetPageNumber(v int32) *DescribeWebRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeWebRulesRequest) SetPageSize(v int32) *DescribeWebRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebRulesRequest) SetQueryDomainPattern(v string) *DescribeWebRulesRequest {
	s.QueryDomainPattern = &v
	return s
}

func (s *DescribeWebRulesRequest) SetResourceGroupId(v string) *DescribeWebRulesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebRulesRequest) Validate() error {
	return dara.Validate(s)
}
