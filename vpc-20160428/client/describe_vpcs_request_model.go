// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDhcpOptionsSetId(v string) *DescribeVpcsRequest
	GetDhcpOptionsSetId() *string
	SetDryRun(v bool) *DescribeVpcsRequest
	GetDryRun() *bool
	SetEnableIpv6(v bool) *DescribeVpcsRequest
	GetEnableIpv6() *bool
	SetIsDefault(v bool) *DescribeVpcsRequest
	GetIsDefault() *bool
	SetOwnerAccount(v string) *DescribeVpcsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpcsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVpcsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpcsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVpcsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeVpcsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeVpcsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpcsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeVpcsRequestTag) *DescribeVpcsRequest
	GetTag() []*DescribeVpcsRequestTag
	SetVpcId(v string) *DescribeVpcsRequest
	GetVpcId() *string
	SetVpcName(v string) *DescribeVpcsRequest
	GetVpcName() *string
	SetVpcOwnerId(v int64) *DescribeVpcsRequest
	GetVpcOwnerId() *int64
}

type DescribeVpcsRequest struct {
	// The ID of the DHCP options set.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system prechecks whether your AccessKey pair is valid, whether the RAM user is authorized, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. If the request passes the check, a 2xx HTTP status code is returned and VPCs are queried.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Query for VPCs in the specified region that have enabled IPv6 CIDR blocks. The value is empty by default, which means no filtering based on IPv6 availability is conducted. Valid values:
	//
	// - false: disabled
	//
	// - true: enabled
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// Specifies whether to query the default VPC in the specified region. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsDefault    *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the VPC.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the VPC to be queried belongs.
	//
	// example:
	//
	// rg-acfmxvfvazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the resource.
	Tag []*DescribeVpcsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// You can specify up to 20 VPC IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// vpc-bp1b1xjllp3ve5yze****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// Vpc-1
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 253460731706911258
	VpcOwnerId *int64 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
}

func (s DescribeVpcsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcsRequest) GetDhcpOptionsSetId() *string {
	return s.DhcpOptionsSetId
}

func (s *DescribeVpcsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeVpcsRequest) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *DescribeVpcsRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpcsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpcsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpcsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpcsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVpcsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpcsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpcsRequest) GetTag() []*DescribeVpcsRequestTag {
	return s.Tag
}

func (s *DescribeVpcsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcsRequest) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcsRequest) GetVpcOwnerId() *int64 {
	return s.VpcOwnerId
}

func (s *DescribeVpcsRequest) SetDhcpOptionsSetId(v string) *DescribeVpcsRequest {
	s.DhcpOptionsSetId = &v
	return s
}

func (s *DescribeVpcsRequest) SetDryRun(v bool) *DescribeVpcsRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeVpcsRequest) SetEnableIpv6(v bool) *DescribeVpcsRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *DescribeVpcsRequest) SetIsDefault(v bool) *DescribeVpcsRequest {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsRequest) SetOwnerAccount(v string) *DescribeVpcsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpcsRequest) SetOwnerId(v int64) *DescribeVpcsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcsRequest) SetPageNumber(v int32) *DescribeVpcsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcsRequest) SetPageSize(v int32) *DescribeVpcsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcsRequest) SetRegionId(v string) *DescribeVpcsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcsRequest) SetResourceGroupId(v string) *DescribeVpcsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVpcsRequest) SetResourceOwnerAccount(v string) *DescribeVpcsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpcsRequest) SetResourceOwnerId(v int64) *DescribeVpcsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpcsRequest) SetTag(v []*DescribeVpcsRequestTag) *DescribeVpcsRequest {
	s.Tag = v
	return s
}

func (s *DescribeVpcsRequest) SetVpcId(v string) *DescribeVpcsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcsRequest) SetVpcName(v string) *DescribeVpcsRequest {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcsRequest) SetVpcOwnerId(v int64) *DescribeVpcsRequest {
	s.VpcOwnerId = &v
	return s
}

func (s *DescribeVpcsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcsRequestTag struct {
	// The key of tag N to add to the resource. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpcsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVpcsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVpcsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVpcsRequestTag) SetKey(v string) *DescribeVpcsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVpcsRequestTag) SetValue(v string) *DescribeVpcsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeVpcsRequestTag) Validate() error {
	return dara.Validate(s)
}
