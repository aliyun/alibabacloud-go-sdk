// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *DescribeLoadBalancersRequest
	GetAddress() *string
	SetAddressType(v string) *DescribeLoadBalancersRequest
	GetAddressType() *string
	SetInternetChargeType(v string) *DescribeLoadBalancersRequest
	GetInternetChargeType() *string
	SetLoadBalancerId(v string) *DescribeLoadBalancersRequest
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *DescribeLoadBalancersRequest
	GetLoadBalancerName() *string
	SetMasterZoneId(v string) *DescribeLoadBalancersRequest
	GetMasterZoneId() *string
	SetNetworkType(v string) *DescribeLoadBalancersRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *DescribeLoadBalancersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLoadBalancersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLoadBalancersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLoadBalancersRequest
	GetPageSize() *int32
	SetPayType(v string) *DescribeLoadBalancersRequest
	GetPayType() *string
	SetRegionId(v string) *DescribeLoadBalancersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeLoadBalancersRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeLoadBalancersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLoadBalancersRequest
	GetResourceOwnerId() *int64
	SetServerId(v string) *DescribeLoadBalancersRequest
	GetServerId() *string
	SetServerIntranetAddress(v string) *DescribeLoadBalancersRequest
	GetServerIntranetAddress() *string
	SetSlaveZoneId(v string) *DescribeLoadBalancersRequest
	GetSlaveZoneId() *string
	SetTags(v string) *DescribeLoadBalancersRequest
	GetTags() *string
	SetVSwitchId(v string) *DescribeLoadBalancersRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeLoadBalancersRequest
	GetVpcId() *string
	SetAccessKeyId(v string) *DescribeLoadBalancersRequest
	GetAccessKeyId() *string
}

type DescribeLoadBalancersRequest struct {
	Address            *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressType        *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	LoadBalancerId     *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName   *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	MasterZoneId       *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	NetworkType        *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PayType            *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// This parameter is required.
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServerId              *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	ServerIntranetAddress *string `json:"ServerIntranetAddress,omitempty" xml:"ServerIntranetAddress,omitempty"`
	SlaveZoneId           *string `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	Tags                  *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                 *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	AccessKeyId           *string `json:"access_key_id,omitempty" xml:"access_key_id,omitempty"`
}

func (s DescribeLoadBalancersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersRequest) GetAddress() *string {
	return s.Address
}

func (s *DescribeLoadBalancersRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeLoadBalancersRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeLoadBalancersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancersRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *DescribeLoadBalancersRequest) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *DescribeLoadBalancersRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeLoadBalancersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLoadBalancersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLoadBalancersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLoadBalancersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoadBalancersRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeLoadBalancersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadBalancersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLoadBalancersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLoadBalancersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLoadBalancersRequest) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeLoadBalancersRequest) GetServerIntranetAddress() *string {
	return s.ServerIntranetAddress
}

func (s *DescribeLoadBalancersRequest) GetSlaveZoneId() *string {
	return s.SlaveZoneId
}

func (s *DescribeLoadBalancersRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeLoadBalancersRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeLoadBalancersRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeLoadBalancersRequest) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *DescribeLoadBalancersRequest) SetAddress(v string) *DescribeLoadBalancersRequest {
	s.Address = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetAddressType(v string) *DescribeLoadBalancersRequest {
	s.AddressType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetInternetChargeType(v string) *DescribeLoadBalancersRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetLoadBalancerId(v string) *DescribeLoadBalancersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetLoadBalancerName(v string) *DescribeLoadBalancersRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetMasterZoneId(v string) *DescribeLoadBalancersRequest {
	s.MasterZoneId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetNetworkType(v string) *DescribeLoadBalancersRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetOwnerAccount(v string) *DescribeLoadBalancersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetOwnerId(v int64) *DescribeLoadBalancersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetPageNumber(v int32) *DescribeLoadBalancersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetPageSize(v int32) *DescribeLoadBalancersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetPayType(v string) *DescribeLoadBalancersRequest {
	s.PayType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetRegionId(v string) *DescribeLoadBalancersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetResourceGroupId(v string) *DescribeLoadBalancersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetServerId(v string) *DescribeLoadBalancersRequest {
	s.ServerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetServerIntranetAddress(v string) *DescribeLoadBalancersRequest {
	s.ServerIntranetAddress = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetSlaveZoneId(v string) *DescribeLoadBalancersRequest {
	s.SlaveZoneId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetTags(v string) *DescribeLoadBalancersRequest {
	s.Tags = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetVSwitchId(v string) *DescribeLoadBalancersRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetVpcId(v string) *DescribeLoadBalancersRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetAccessKeyId(v string) *DescribeLoadBalancersRequest {
	s.AccessKeyId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) Validate() error {
	return dara.Validate(s)
}
