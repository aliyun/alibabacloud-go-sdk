// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDhcpOptionsSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDhcpOptionsSetId(v []*string) *ListDhcpOptionsSetsRequest
	GetDhcpOptionsSetId() []*string
	SetDhcpOptionsSetName(v string) *ListDhcpOptionsSetsRequest
	GetDhcpOptionsSetName() *string
	SetDomainName(v string) *ListDhcpOptionsSetsRequest
	GetDomainName() *string
	SetMaxResults(v int32) *ListDhcpOptionsSetsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDhcpOptionsSetsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListDhcpOptionsSetsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListDhcpOptionsSetsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListDhcpOptionsSetsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListDhcpOptionsSetsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ListDhcpOptionsSetsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListDhcpOptionsSetsRequest
	GetResourceOwnerId() *int64
	SetTags(v []*ListDhcpOptionsSetsRequestTags) *ListDhcpOptionsSetsRequest
	GetTags() []*ListDhcpOptionsSetsRequestTags
}

type ListDhcpOptionsSetsRequest struct {
	// The ID of the DHCP options set. You can specify at most 20 IDs.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId []*string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty" type:"Repeated"`
	// The name of the DHCP options set.
	//
	// The name must be 1 to 128 characters in length and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// Valid values:
	//
	// 	- tf-testAccVpcDhcpOptionsSets-1585169790614573448
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     tf-testAccVpcDhcpOptionsSets-1585169790614573448
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// test
	DhcpOptionsSetName *string `json:"DhcpOptionsSetName,omitempty" xml:"DhcpOptionsSetName,omitempty"`
	// The root domain. For example, you can set the value to example.com.
	//
	// After a DHCP options set is associated with a virtual private cloud (VPC), the root domain in the DHCP options set is automatically synchronized with the ECS instances in the VPC.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of the **NextToken*	- parameter.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the DHCP options sets that you want to query.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the DHCP options set belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag list.
	Tags []*ListDhcpOptionsSetsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListDhcpOptionsSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDhcpOptionsSetsRequest) GoString() string {
	return s.String()
}

func (s *ListDhcpOptionsSetsRequest) GetDhcpOptionsSetId() []*string {
	return s.DhcpOptionsSetId
}

func (s *ListDhcpOptionsSetsRequest) GetDhcpOptionsSetName() *string {
	return s.DhcpOptionsSetName
}

func (s *ListDhcpOptionsSetsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ListDhcpOptionsSetsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDhcpOptionsSetsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDhcpOptionsSetsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListDhcpOptionsSetsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDhcpOptionsSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDhcpOptionsSetsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDhcpOptionsSetsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListDhcpOptionsSetsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListDhcpOptionsSetsRequest) GetTags() []*ListDhcpOptionsSetsRequestTags {
	return s.Tags
}

func (s *ListDhcpOptionsSetsRequest) SetDhcpOptionsSetId(v []*string) *ListDhcpOptionsSetsRequest {
	s.DhcpOptionsSetId = v
	return s
}

func (s *ListDhcpOptionsSetsRequest) SetDhcpOptionsSetName(v string) *ListDhcpOptionsSetsRequest {
	s.DhcpOptionsSetName = &v
	return s
}

func (s *ListDhcpOptionsSetsRequest) SetDomainName(v string) *ListDhcpOptionsSetsRequest {
	s.DomainName = &v
	return s
}

func (s *ListDhcpOptionsSetsRequest) SetMaxResults(v int32) *ListDhcpOptionsSetsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDhcpOptionsSetsRequest) SetNextToken(v string) *ListDhcpOptionsSetsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDhcpOptionsSetsRequest) SetOwnerAccount(v string) *ListDhcpOptionsSetsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListDhcpOptionsSetsRequest) SetOwnerId(v int64) *ListDhcpOptionsSetsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDhcpOptionsSetsRequest) SetRegionId(v string) *ListDhcpOptionsSetsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDhcpOptionsSetsRequest) SetResourceGroupId(v string) *ListDhcpOptionsSetsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDhcpOptionsSetsRequest) SetResourceOwnerAccount(v string) *ListDhcpOptionsSetsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListDhcpOptionsSetsRequest) SetResourceOwnerId(v int64) *ListDhcpOptionsSetsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListDhcpOptionsSetsRequest) SetTags(v []*ListDhcpOptionsSetsRequestTags) *ListDhcpOptionsSetsRequest {
	s.Tags = v
	return s
}

func (s *ListDhcpOptionsSetsRequest) Validate() error {
	return dara.Validate(s)
}

type ListDhcpOptionsSetsRequestTags struct {
	// The tag key. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The tag key must start with a letter but cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). The tag value must start with a letter but cannot start with `aliyun` or `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDhcpOptionsSetsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListDhcpOptionsSetsRequestTags) GoString() string {
	return s.String()
}

func (s *ListDhcpOptionsSetsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListDhcpOptionsSetsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListDhcpOptionsSetsRequestTags) SetKey(v string) *ListDhcpOptionsSetsRequestTags {
	s.Key = &v
	return s
}

func (s *ListDhcpOptionsSetsRequestTags) SetValue(v string) *ListDhcpOptionsSetsRequestTags {
	s.Value = &v
	return s
}

func (s *ListDhcpOptionsSetsRequestTags) Validate() error {
	return dara.Validate(s)
}
