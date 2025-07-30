// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheCustomInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetChargeType() *string
	SetExpired(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetExpired() *string
	SetInstanceClass(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetInstanceClass() *string
	SetInstanceIds(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetInstanceIds() *string
	SetInstanceStatus(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetInstanceStatus() *string
	SetInstanceType(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetInstanceType() *string
	SetNetworkType(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTairKVCacheCustomInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeTairKVCacheCustomInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTairKVCacheCustomInstancesRequest
	GetPageSize() *int32
	SetPrivateIp(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetPrivateIp() *string
	SetRegionId(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTairKVCacheCustomInstancesRequest
	GetResourceOwnerId() *int64
	SetSearchKey(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetSearchKey() *string
	SetSecurityToken(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetSecurityToken() *string
	SetTag(v []*DescribeTairKVCacheCustomInstancesRequestTag) *DescribeTairKVCacheCustomInstancesRequest
	GetTag() []*DescribeTairKVCacheCustomInstancesRequestTag
	SetVSwitchId(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeTairKVCacheCustomInstancesRequest
	GetZoneId() *string
}

type DescribeTairKVCacheCustomInstancesRequest struct {
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// tair.gpu.test.16g
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// example:
	//
	// tc-bp16e70a4338****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// Normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// TairCustom
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// VPC
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// apitest
	SearchKey     *string                                         `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	SecurityToken *string                                         `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*DescribeTairKVCacheCustomInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeTairKVCacheCustomInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetExpired() *string {
	return s.Expired
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetTag() []*DescribeTairKVCacheCustomInstancesRequestTag {
	return s.Tag
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeTairKVCacheCustomInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetChargeType(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetExpired(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.Expired = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetInstanceClass(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.InstanceClass = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetInstanceIds(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetInstanceStatus(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetInstanceType(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetNetworkType(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetOwnerAccount(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetOwnerId(v int64) *DescribeTairKVCacheCustomInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetPageNumber(v int32) *DescribeTairKVCacheCustomInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetPageSize(v int32) *DescribeTairKVCacheCustomInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetPrivateIp(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.PrivateIp = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetRegionId(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetResourceGroupId(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetResourceOwnerAccount(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetResourceOwnerId(v int64) *DescribeTairKVCacheCustomInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetSearchKey(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetSecurityToken(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetTag(v []*DescribeTairKVCacheCustomInstancesRequestTag) *DescribeTairKVCacheCustomInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetVSwitchId(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetVpcId(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) SetZoneId(v string) *DescribeTairKVCacheCustomInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheCustomInstancesRequestTag struct {
	// example:
	//
	// key1_test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1_test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTairKVCacheCustomInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeTairKVCacheCustomInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeTairKVCacheCustomInstancesRequestTag) SetKey(v string) *DescribeTairKVCacheCustomInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequestTag) SetValue(v string) *DescribeTairKVCacheCustomInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
