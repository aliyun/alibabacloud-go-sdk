// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DescribeVSwitchesRequest
	GetDryRun() *bool
	SetEnableIpv6(v bool) *DescribeVSwitchesRequest
	GetEnableIpv6() *bool
	SetIsDefault(v bool) *DescribeVSwitchesRequest
	GetIsDefault() *bool
	SetOwnerAccount(v string) *DescribeVSwitchesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVSwitchesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVSwitchesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVSwitchesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVSwitchesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeVSwitchesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeVSwitchesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVSwitchesRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *DescribeVSwitchesRequest
	GetRouteTableId() *string
	SetTag(v []*DescribeVSwitchesRequestTag) *DescribeVSwitchesRequest
	GetTag() []*DescribeVSwitchesRequestTag
	SetVSwitchId(v string) *DescribeVSwitchesRequest
	GetVSwitchId() *string
	SetVSwitchName(v string) *DescribeVSwitchesRequest
	GetVSwitchName() *string
	SetVSwitchOwnerId(v int64) *DescribeVSwitchesRequest
	GetVSwitchOwnerId() *int64
	SetVpcId(v string) *DescribeVSwitchesRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeVSwitchesRequest
	GetZoneId() *string
}

type DescribeVSwitchesRequest struct {
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to query vSwitches with IPv6 enabled in the region. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// If you do not set this parameter, the system queries all vSwitches in the specified region by default.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	EnableIpv6 *bool `json:"EnableIpv6,omitempty" xml:"EnableIpv6,omitempty"`
	// Specifies whether to query the default vSwitches in the specified region. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// If you do not set this parameter, the system queries all vSwitches in the specified region by default.
	//
	// example:
	//
	// true
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
	// The region ID of the vSwitch. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// >  You must set at least one of **RegionId*	- and **VpcId**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the vSwitch belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the route table.
	//
	// example:
	//
	// vtb-bp145q7glnuzdvzu2****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The tags.
	Tag []*DescribeVSwitchesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch that you want to query.
	//
	// example:
	//
	// vsw-23dscddcffvf3****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The vSwitch name.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// vSwitch
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The ID of the Alibaba Cloud account to which the vSwitch belongs.
	//
	// example:
	//
	// 2546073170691****
	VSwitchOwnerId *int64 `json:"VSwitchOwnerId,omitempty" xml:"VSwitchOwnerId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the vSwitches belong.
	//
	// >  You must set at least one of **RegionId*	- and **VpcId**.
	//
	// example:
	//
	// vpc-25cdvfeq58pl****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone to which the vSwitches belong. You can call the [DescribeZones](https://help.aliyun.com/document_detail/36064.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-d
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVSwitchesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeVSwitchesRequest) GetEnableIpv6() *bool {
	return s.EnableIpv6
}

func (s *DescribeVSwitchesRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVSwitchesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVSwitchesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVSwitchesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVSwitchesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVSwitchesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVSwitchesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVSwitchesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVSwitchesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVSwitchesRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeVSwitchesRequest) GetTag() []*DescribeVSwitchesRequestTag {
	return s.Tag
}

func (s *DescribeVSwitchesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchesRequest) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchesRequest) GetVSwitchOwnerId() *int64 {
	return s.VSwitchOwnerId
}

func (s *DescribeVSwitchesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVSwitchesRequest) SetDryRun(v bool) *DescribeVSwitchesRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetEnableIpv6(v bool) *DescribeVSwitchesRequest {
	s.EnableIpv6 = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetIsDefault(v bool) *DescribeVSwitchesRequest {
	s.IsDefault = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetOwnerAccount(v string) *DescribeVSwitchesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetOwnerId(v int64) *DescribeVSwitchesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageNumber(v int32) *DescribeVSwitchesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetPageSize(v int32) *DescribeVSwitchesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetRegionId(v string) *DescribeVSwitchesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetResourceGroupId(v string) *DescribeVSwitchesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetResourceOwnerAccount(v string) *DescribeVSwitchesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetResourceOwnerId(v int64) *DescribeVSwitchesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetRouteTableId(v string) *DescribeVSwitchesRequest {
	s.RouteTableId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetTag(v []*DescribeVSwitchesRequestTag) *DescribeVSwitchesRequest {
	s.Tag = v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchId(v string) *DescribeVSwitchesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchName(v string) *DescribeVSwitchesRequest {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVSwitchOwnerId(v int64) *DescribeVSwitchesRequest {
	s.VSwitchOwnerId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetVpcId(v string) *DescribeVSwitchesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchesRequest) SetZoneId(v string) *DescribeVSwitchesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeVSwitchesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeVSwitchesRequestTag struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVSwitchesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVSwitchesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVSwitchesRequestTag) SetKey(v string) *DescribeVSwitchesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVSwitchesRequestTag) SetValue(v string) *DescribeVSwitchesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeVSwitchesRequestTag) Validate() error {
	return dara.Validate(s)
}
