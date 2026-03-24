// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddressesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressLike(v string) *DescribeAddressesRequest
	GetAddressLike() *string
	SetInstanceId(v string) *DescribeAddressesRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeAddressesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAddressesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeAddressesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAddressesRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *DescribeAddressesRequest
	GetRuleId() *int64
}

type DescribeAddressesRequest struct {
	// The address to use for a fuzzy match. If you specify this parameter, only addresses that contain the specified string are returned.
	//
	// example:
	//
	// 1.2.3.3
	AddressLike *string `json:"AddressLike,omitempty" xml:"AddressLike,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 500. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next page of results. Set this parameter to the value of **NextToken*	- that is returned from the previous call. Do not specify this parameter for the first call.
	//
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm4gh****wela
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The ID of the address book to query.
	//
	// example:
	//
	// 12345678
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeAddressesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddressesRequest) GetAddressLike() *string {
	return s.AddressLike
}

func (s *DescribeAddressesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAddressesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAddressesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAddressesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAddressesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAddressesRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeAddressesRequest) SetAddressLike(v string) *DescribeAddressesRequest {
	s.AddressLike = &v
	return s
}

func (s *DescribeAddressesRequest) SetInstanceId(v string) *DescribeAddressesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAddressesRequest) SetMaxResults(v int32) *DescribeAddressesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAddressesRequest) SetNextToken(v string) *DescribeAddressesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAddressesRequest) SetRegionId(v string) *DescribeAddressesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAddressesRequest) SetResourceManagerResourceGroupId(v string) *DescribeAddressesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAddressesRequest) SetRuleId(v int64) *DescribeAddressesRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeAddressesRequest) Validate() error {
	return dara.Validate(s)
}
