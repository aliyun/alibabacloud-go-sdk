// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnGatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessStatus(v string) *DescribeVpnGatewaysRequest
	GetBusinessStatus() *string
	SetGatewayType(v string) *DescribeVpnGatewaysRequest
	GetGatewayType() *string
	SetIncludeReservationData(v bool) *DescribeVpnGatewaysRequest
	GetIncludeReservationData() *bool
	SetOwnerAccount(v string) *DescribeVpnGatewaysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpnGatewaysRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVpnGatewaysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnGatewaysRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVpnGatewaysRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeVpnGatewaysRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeVpnGatewaysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpnGatewaysRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeVpnGatewaysRequest
	GetStatus() *string
	SetTag(v []*DescribeVpnGatewaysRequestTag) *DescribeVpnGatewaysRequest
	GetTag() []*DescribeVpnGatewaysRequestTag
	SetVpcId(v string) *DescribeVpnGatewaysRequest
	GetVpcId() *string
	SetVpnGatewayId(v string) *DescribeVpnGatewaysRequest
	GetVpnGatewayId() *string
}

type DescribeVpnGatewaysRequest struct {
	// The payment status of the VPN gateway. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// VPN 网关类型，取值：
	//
	// Traditional：传统型VPN网关，覆盖IPsec功能和SSL功能
	//
	// Enhance.SiteToSite：增强型站点入云VPN，只覆盖IPsec功能
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// Specifies whether to return information about pending orders. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// true
	IncludeReservationData *bool   `json:"IncludeReservationData,omitempty" xml:"IncludeReservationData,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the VPN gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the VPN gateway belongs.
	//
	//  You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query the resource group list.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the VPN gateway. Valid values:
	//
	// 	- **init**
	//
	// 	- **provisioning**
	//
	// 	- **active**
	//
	// 	- **updating**
	//
	// 	- **deleting**
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the VPN gateway.
	Tag []*DescribeVpnGatewaysRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the VPN gateway belongs.
	//
	// example:
	//
	// vpc-bp1m3i0kn1nd4wiw9****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp17lofy9fd0dnvzv****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s DescribeVpnGatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysRequest) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeVpnGatewaysRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *DescribeVpnGatewaysRequest) GetIncludeReservationData() *bool {
	return s.IncludeReservationData
}

func (s *DescribeVpnGatewaysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpnGatewaysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpnGatewaysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnGatewaysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnGatewaysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpnGatewaysRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVpnGatewaysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpnGatewaysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpnGatewaysRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpnGatewaysRequest) GetTag() []*DescribeVpnGatewaysRequestTag {
	return s.Tag
}

func (s *DescribeVpnGatewaysRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpnGatewaysRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *DescribeVpnGatewaysRequest) SetBusinessStatus(v string) *DescribeVpnGatewaysRequest {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetGatewayType(v string) *DescribeVpnGatewaysRequest {
	s.GatewayType = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetIncludeReservationData(v bool) *DescribeVpnGatewaysRequest {
	s.IncludeReservationData = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetOwnerAccount(v string) *DescribeVpnGatewaysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetOwnerId(v int64) *DescribeVpnGatewaysRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetPageNumber(v int32) *DescribeVpnGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetPageSize(v int32) *DescribeVpnGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetRegionId(v string) *DescribeVpnGatewaysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetResourceGroupId(v string) *DescribeVpnGatewaysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetResourceOwnerAccount(v string) *DescribeVpnGatewaysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetResourceOwnerId(v int64) *DescribeVpnGatewaysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetStatus(v string) *DescribeVpnGatewaysRequest {
	s.Status = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetTag(v []*DescribeVpnGatewaysRequestTag) *DescribeVpnGatewaysRequest {
	s.Tag = v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetVpcId(v string) *DescribeVpnGatewaysRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) SetVpnGatewayId(v string) *DescribeVpnGatewaysRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *DescribeVpnGatewaysRequest) Validate() error {
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

type DescribeVpnGatewaysRequestTag struct {
	// The tag key.
	//
	// You can specify at most 20 tag keys at a time.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// Each tag key corresponds to one tag value. You can specify at most 20 tag values at a time.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVpnGatewaysRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVpnGatewaysRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVpnGatewaysRequestTag) SetKey(v string) *DescribeVpnGatewaysRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVpnGatewaysRequestTag) SetValue(v string) *DescribeVpnGatewaysRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeVpnGatewaysRequestTag) Validate() error {
	return dara.Validate(s)
}
