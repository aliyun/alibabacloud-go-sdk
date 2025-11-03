// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DescribeNatGatewaysRequest
	GetDryRun() *bool
	SetInstanceChargeType(v string) *DescribeNatGatewaysRequest
	GetInstanceChargeType() *string
	SetName(v string) *DescribeNatGatewaysRequest
	GetName() *string
	SetNatGatewayId(v string) *DescribeNatGatewaysRequest
	GetNatGatewayId() *string
	SetNatType(v string) *DescribeNatGatewaysRequest
	GetNatType() *string
	SetNetworkType(v string) *DescribeNatGatewaysRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *DescribeNatGatewaysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeNatGatewaysRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeNatGatewaysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNatGatewaysRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeNatGatewaysRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeNatGatewaysRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeNatGatewaysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeNatGatewaysRequest
	GetResourceOwnerId() *int64
	SetSpec(v string) *DescribeNatGatewaysRequest
	GetSpec() *string
	SetStatus(v string) *DescribeNatGatewaysRequest
	GetStatus() *string
	SetTag(v []*DescribeNatGatewaysRequestTag) *DescribeNatGatewaysRequest
	GetTag() []*DescribeNatGatewaysRequestTag
	SetVpcId(v string) *DescribeNatGatewaysRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeNatGatewaysRequest
	GetZoneId() *string
}

type DescribeNatGatewaysRequest struct {
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system prechecks whether your AccessKey pair is valid, whether the RAM user is authorized, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The billing method of the NAT gateway. Set the value to **PostPaid**, which specifies the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the NAT gateway.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// If this parameter is not set, the system automatically assigns a name to the NAT gateway.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The type of NAT gateway. Set the value to **Enhanced*	- (enhanced NAT gateway).
	//
	// example:
	//
	// Enhanced
	NatType *string `json:"NatType,omitempty" xml:"NatType,omitempty"`
	// The type of the NAT gateway. Valid values:
	//
	// 	- **internet**: an Internet NAT gateway
	//
	// 	- **intranet**: a VPC NAT gateway
	//
	// example:
	//
	// internet
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the NAT gateways that you want to query.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the NAT gateway belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The size of the NAT gateway. Ignore this parameter.
	//
	// example:
	//
	// Invalid parameter.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the NAT gateway. Valid values:
	//
	// 	- **Creating**: After you send a request to create a NAT gateway, the system creates the NAT gateway in the background. The NAT gateway remains in the **Creating*	- state until the operation is completed.
	//
	// 	- **Available**: The NAT gateway remains in a stable state after the NAT gateway is created.
	//
	// 	- **Modifying**: After you send a request to modify a NAT gateway, the system modifies the NAT gateway in the background. The NAT gateway remains in the **Modifying*	- state until the operation is completed.
	//
	// 	- **Deleting**: After you send a request to delete a NAT gateway, the system deletes the NAT gateway in the background. The NAT gateway remains in the **Deleting*	- state until the operation is completed.
	//
	// 	- **Converting**: After you send a request to upgrade a standard NAT gateway to an enhanced NAT gateway, the system upgrades the NAT gateway in the background. The NAT gateway remains in the **Converting*	- state until the operation is completed.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tag []*DescribeNatGatewaysRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC to which the NAT gateway belongs.
	//
	// example:
	//
	// vpc-bp15zckdt37pq72z****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone to which the NAT gateway belongs.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeNatGatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeNatGatewaysRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeNatGatewaysRequest) GetName() *string {
	return s.Name
}

func (s *DescribeNatGatewaysRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatGatewaysRequest) GetNatType() *string {
	return s.NatType
}

func (s *DescribeNatGatewaysRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeNatGatewaysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeNatGatewaysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNatGatewaysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNatGatewaysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNatGatewaysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNatGatewaysRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeNatGatewaysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNatGatewaysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeNatGatewaysRequest) GetSpec() *string {
	return s.Spec
}

func (s *DescribeNatGatewaysRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeNatGatewaysRequest) GetTag() []*DescribeNatGatewaysRequestTag {
	return s.Tag
}

func (s *DescribeNatGatewaysRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNatGatewaysRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeNatGatewaysRequest) SetDryRun(v bool) *DescribeNatGatewaysRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetInstanceChargeType(v string) *DescribeNatGatewaysRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetName(v string) *DescribeNatGatewaysRequest {
	s.Name = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetNatGatewayId(v string) *DescribeNatGatewaysRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetNatType(v string) *DescribeNatGatewaysRequest {
	s.NatType = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetNetworkType(v string) *DescribeNatGatewaysRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetOwnerAccount(v string) *DescribeNatGatewaysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetOwnerId(v int64) *DescribeNatGatewaysRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetPageNumber(v int32) *DescribeNatGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetPageSize(v int32) *DescribeNatGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetRegionId(v string) *DescribeNatGatewaysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetResourceGroupId(v string) *DescribeNatGatewaysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetResourceOwnerAccount(v string) *DescribeNatGatewaysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetResourceOwnerId(v int64) *DescribeNatGatewaysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetSpec(v string) *DescribeNatGatewaysRequest {
	s.Spec = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetStatus(v string) *DescribeNatGatewaysRequest {
	s.Status = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetTag(v []*DescribeNatGatewaysRequestTag) *DescribeNatGatewaysRequest {
	s.Tag = v
	return s
}

func (s *DescribeNatGatewaysRequest) SetVpcId(v string) *DescribeNatGatewaysRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetZoneId(v string) *DescribeNatGatewaysRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatGatewaysRequestTag struct {
	// The tag keys of the NAT gateway. You can specify up to 20 tag keys.
	//
	// Each tag key cannot exceed 64 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// KeyTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values of the NAT gateway. You can specify up to 20 tag values.
	//
	// The tag value cannot exceed 128 characters in length, and cannot start with `aliyun` or `acs:`. The value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// valueTest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNatGatewaysRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeNatGatewaysRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeNatGatewaysRequestTag) SetKey(v string) *DescribeNatGatewaysRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeNatGatewaysRequestTag) SetValue(v string) *DescribeNatGatewaysRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeNatGatewaysRequestTag) Validate() error {
	return dara.Validate(s)
}
