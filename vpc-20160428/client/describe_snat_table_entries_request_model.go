// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnatTableEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNatGatewayId(v string) *DescribeSnatTableEntriesRequest
	GetNatGatewayId() *string
	SetNetworkInterfaceIds(v []*string) *DescribeSnatTableEntriesRequest
	GetNetworkInterfaceIds() []*string
	SetOwnerAccount(v string) *DescribeSnatTableEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSnatTableEntriesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSnatTableEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnatTableEntriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSnatTableEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSnatTableEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSnatTableEntriesRequest
	GetResourceOwnerId() *int64
	SetSnatEntryId(v string) *DescribeSnatTableEntriesRequest
	GetSnatEntryId() *string
	SetSnatEntryName(v string) *DescribeSnatTableEntriesRequest
	GetSnatEntryName() *string
	SetSnatIp(v string) *DescribeSnatTableEntriesRequest
	GetSnatIp() *string
	SetSnatTableId(v string) *DescribeSnatTableEntriesRequest
	GetSnatTableId() *string
	SetSourceCIDR(v string) *DescribeSnatTableEntriesRequest
	GetSourceCIDR() *string
	SetSourceVSwitchId(v string) *DescribeSnatTableEntriesRequest
	GetSourceVSwitchId() *string
}

type DescribeSnatTableEntriesRequest struct {
	// The ID of the NAT gateway.
	//
	// >  You must specify at least one of **SnatTableId*	- and **NatGatewayId**.
	//
	// example:
	//
	// ngw-bp1uewa15k4iy5770****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The ID of the elastic network interface to be queried.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Repeated"`
	OwnerAccount        *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The ID of the region where you want to create the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SNAT entry.
	//
	// example:
	//
	// snat-8vbae8uqh7rjpk7d2****
	SnatEntryId *string `json:"SnatEntryId,omitempty" xml:"SnatEntryId,omitempty"`
	// The name of the SNAT entry.
	//
	// The name must be 2 to 128 characters in length, and can contain digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// SnatEntry-1
	SnatEntryName *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	// 	- When you query SNAT entries of Internet NAT gateways, this parameter specifies the EIP in an SNAT entry.
	//
	// 	- When you query SNAT entries of VPC NAT gateways, this parameter specifies the NAT IP address in an SNAT entry.
	//
	// example:
	//
	// 116.22.XX.XX
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
	// The ID of the SNAT table.
	//
	// >  You must specify at least one of **SnatTableId*	- and **NatGatewayId**.
	//
	// example:
	//
	// stb-8vbczigrhop8x5u3t****
	SnatTableId *string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty"`
	// The source CIDR block specified in the SNAT entry.
	//
	// example:
	//
	// 116.22.XX.XX/24
	SourceCIDR *string `json:"SourceCIDR,omitempty" xml:"SourceCIDR,omitempty"`
	// The ID of the vSwitch.
	//
	// 	- When you query SNAT entries of Internet NAT gateways, this parameter specifies that Elastic Compute Service (ECS) instances in the vSwitch can use SNAT entries to access the Internet.
	//
	// 	- When you query SNAT entries of virtual private cloud (VPC) NAT gateways, this parameter specifies that ECS instances in the vSwitch can use SNAT entries to access external networks.
	//
	// example:
	//
	// vsw-3xbjkhjshjdf****
	SourceVSwitchId *string `json:"SourceVSwitchId,omitempty" xml:"SourceVSwitchId,omitempty"`
}

func (s DescribeSnatTableEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatTableEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnatTableEntriesRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeSnatTableEntriesRequest) GetNetworkInterfaceIds() []*string {
	return s.NetworkInterfaceIds
}

func (s *DescribeSnatTableEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSnatTableEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSnatTableEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnatTableEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnatTableEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnatTableEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSnatTableEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSnatTableEntriesRequest) GetSnatEntryId() *string {
	return s.SnatEntryId
}

func (s *DescribeSnatTableEntriesRequest) GetSnatEntryName() *string {
	return s.SnatEntryName
}

func (s *DescribeSnatTableEntriesRequest) GetSnatIp() *string {
	return s.SnatIp
}

func (s *DescribeSnatTableEntriesRequest) GetSnatTableId() *string {
	return s.SnatTableId
}

func (s *DescribeSnatTableEntriesRequest) GetSourceCIDR() *string {
	return s.SourceCIDR
}

func (s *DescribeSnatTableEntriesRequest) GetSourceVSwitchId() *string {
	return s.SourceVSwitchId
}

func (s *DescribeSnatTableEntriesRequest) SetNatGatewayId(v string) *DescribeSnatTableEntriesRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetNetworkInterfaceIds(v []*string) *DescribeSnatTableEntriesRequest {
	s.NetworkInterfaceIds = v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetOwnerAccount(v string) *DescribeSnatTableEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetOwnerId(v int64) *DescribeSnatTableEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetPageNumber(v int32) *DescribeSnatTableEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetPageSize(v int32) *DescribeSnatTableEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetRegionId(v string) *DescribeSnatTableEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetResourceOwnerAccount(v string) *DescribeSnatTableEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetResourceOwnerId(v int64) *DescribeSnatTableEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSnatEntryId(v string) *DescribeSnatTableEntriesRequest {
	s.SnatEntryId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSnatEntryName(v string) *DescribeSnatTableEntriesRequest {
	s.SnatEntryName = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSnatIp(v string) *DescribeSnatTableEntriesRequest {
	s.SnatIp = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSnatTableId(v string) *DescribeSnatTableEntriesRequest {
	s.SnatTableId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSourceCIDR(v string) *DescribeSnatTableEntriesRequest {
	s.SourceCIDR = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) SetSourceVSwitchId(v string) *DescribeSnatTableEntriesRequest {
	s.SourceVSwitchId = &v
	return s
}

func (s *DescribeSnatTableEntriesRequest) Validate() error {
	return dara.Validate(s)
}
