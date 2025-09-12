// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeLindormV2StreamEngineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomConfig(v string) *UpgradeLindormV2StreamEngineRequest
	GetCustomConfig() *string
	SetInstanceId(v string) *UpgradeLindormV2StreamEngineRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *UpgradeLindormV2StreamEngineRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpgradeLindormV2StreamEngineRequest
	GetOwnerId() *int64
	SetQuantity(v int32) *UpgradeLindormV2StreamEngineRequest
	GetQuantity() *int32
	SetResourceGroupName(v string) *UpgradeLindormV2StreamEngineRequest
	GetResourceGroupName() *string
	SetResourceOwnerAccount(v string) *UpgradeLindormV2StreamEngineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeLindormV2StreamEngineRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *UpgradeLindormV2StreamEngineRequest
	GetSecurityToken() *string
	SetSpec(v string) *UpgradeLindormV2StreamEngineRequest
	GetSpec() *string
	SetSpecId(v string) *UpgradeLindormV2StreamEngineRequest
	GetSpecId() *string
	SetUpgradeType(v string) *UpgradeLindormV2StreamEngineRequest
	GetUpgradeType() *string
}

type UpgradeLindormV2StreamEngineRequest struct {
	CustomConfig *string `json:"CustomConfig,omitempty" xml:"CustomConfig,omitempty"`
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	Quantity             *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ResourceGroupName    *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// This parameter is required.
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	// This parameter is required.
	UpgradeType *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
}

func (s UpgradeLindormV2StreamEngineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeLindormV2StreamEngineRequest) GoString() string {
	return s.String()
}

func (s *UpgradeLindormV2StreamEngineRequest) GetCustomConfig() *string {
	return s.CustomConfig
}

func (s *UpgradeLindormV2StreamEngineRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeLindormV2StreamEngineRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpgradeLindormV2StreamEngineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeLindormV2StreamEngineRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *UpgradeLindormV2StreamEngineRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *UpgradeLindormV2StreamEngineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeLindormV2StreamEngineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeLindormV2StreamEngineRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpgradeLindormV2StreamEngineRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpgradeLindormV2StreamEngineRequest) GetSpecId() *string {
	return s.SpecId
}

func (s *UpgradeLindormV2StreamEngineRequest) GetUpgradeType() *string {
	return s.UpgradeType
}

func (s *UpgradeLindormV2StreamEngineRequest) SetCustomConfig(v string) *UpgradeLindormV2StreamEngineRequest {
	s.CustomConfig = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetInstanceId(v string) *UpgradeLindormV2StreamEngineRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetOwnerAccount(v string) *UpgradeLindormV2StreamEngineRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetOwnerId(v int64) *UpgradeLindormV2StreamEngineRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetQuantity(v int32) *UpgradeLindormV2StreamEngineRequest {
	s.Quantity = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetResourceGroupName(v string) *UpgradeLindormV2StreamEngineRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetResourceOwnerAccount(v string) *UpgradeLindormV2StreamEngineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetResourceOwnerId(v int64) *UpgradeLindormV2StreamEngineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetSecurityToken(v string) *UpgradeLindormV2StreamEngineRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetSpec(v string) *UpgradeLindormV2StreamEngineRequest {
	s.Spec = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetSpecId(v string) *UpgradeLindormV2StreamEngineRequest {
	s.SpecId = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetUpgradeType(v string) *UpgradeLindormV2StreamEngineRequest {
	s.UpgradeType = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) Validate() error {
	return dara.Validate(s)
}
