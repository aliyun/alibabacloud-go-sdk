// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheInferInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *DescribeTairKVCacheInferInstancesRequest
	GetChargeType() *string
	SetExpired(v string) *DescribeTairKVCacheInferInstancesRequest
	GetExpired() *string
	SetInstanceClass(v string) *DescribeTairKVCacheInferInstancesRequest
	GetInstanceClass() *string
	SetInstanceIds(v string) *DescribeTairKVCacheInferInstancesRequest
	GetInstanceIds() *string
	SetInstanceStatus(v string) *DescribeTairKVCacheInferInstancesRequest
	GetInstanceStatus() *string
	SetNetworkType(v string) *DescribeTairKVCacheInferInstancesRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *DescribeTairKVCacheInferInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTairKVCacheInferInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeTairKVCacheInferInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTairKVCacheInferInstancesRequest
	GetPageSize() *int32
	SetPrivateIp(v string) *DescribeTairKVCacheInferInstancesRequest
	GetPrivateIp() *string
	SetRegionId(v string) *DescribeTairKVCacheInferInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeTairKVCacheInferInstancesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeTairKVCacheInferInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTairKVCacheInferInstancesRequest
	GetResourceOwnerId() *int64
	SetSearchKey(v string) *DescribeTairKVCacheInferInstancesRequest
	GetSearchKey() *string
	SetSecurityToken(v string) *DescribeTairKVCacheInferInstancesRequest
	GetSecurityToken() *string
	SetTag(v []*DescribeTairKVCacheInferInstancesRequestTag) *DescribeTairKVCacheInferInstancesRequest
	GetTag() []*DescribeTairKVCacheInferInstancesRequestTag
	SetVSwitchId(v string) *DescribeTairKVCacheInferInstancesRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeTairKVCacheInferInstancesRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeTairKVCacheInferInstancesRequest
	GetZoneId() *string
}

type DescribeTairKVCacheInferInstancesRequest struct {
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Expired        *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	InstanceClass  *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceIds    *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	NetworkType    *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrivateIp      *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// This parameter is required.
	RegionId             *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                         `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SearchKey            *string                                        `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	SecurityToken        *string                                        `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag                  []*DescribeTairKVCacheInferInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VSwitchId            *string                                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                *string                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeTairKVCacheInferInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetExpired() *string {
	return s.Expired
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetTag() []*DescribeTairKVCacheInferInstancesRequestTag {
	return s.Tag
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeTairKVCacheInferInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetChargeType(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetExpired(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.Expired = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetInstanceClass(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.InstanceClass = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetInstanceIds(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetInstanceStatus(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetNetworkType(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetOwnerAccount(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetOwnerId(v int64) *DescribeTairKVCacheInferInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetPageNumber(v int32) *DescribeTairKVCacheInferInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetPageSize(v int32) *DescribeTairKVCacheInferInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetPrivateIp(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.PrivateIp = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetRegionId(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetResourceGroupId(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetResourceOwnerAccount(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetResourceOwnerId(v int64) *DescribeTairKVCacheInferInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetSearchKey(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetSecurityToken(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetTag(v []*DescribeTairKVCacheInferInstancesRequestTag) *DescribeTairKVCacheInferInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetVSwitchId(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetVpcId(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) SetZoneId(v string) *DescribeTairKVCacheInferInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheInferInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTairKVCacheInferInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeTairKVCacheInferInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeTairKVCacheInferInstancesRequestTag) SetKey(v string) *DescribeTairKVCacheInferInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequestTag) SetValue(v string) *DescribeTairKVCacheInferInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
