// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocateStrategy(v string) *ModifyDBInstanceSpecShrinkRequest
	GetAllocateStrategy() *string
	SetAllowMajorVersionUpgrade(v bool) *ModifyDBInstanceSpecShrinkRequest
	GetAllowMajorVersionUpgrade() *bool
	SetAutoUseCoupon(v bool) *ModifyDBInstanceSpecShrinkRequest
	GetAutoUseCoupon() *bool
	SetBurstingEnabled(v bool) *ModifyDBInstanceSpecShrinkRequest
	GetBurstingEnabled() *bool
	SetCategory(v string) *ModifyDBInstanceSpecShrinkRequest
	GetCategory() *string
	SetColdDataEnabled(v bool) *ModifyDBInstanceSpecShrinkRequest
	GetColdDataEnabled() *bool
	SetCompressionMode(v string) *ModifyDBInstanceSpecShrinkRequest
	GetCompressionMode() *string
	SetDBInstanceClass(v string) *ModifyDBInstanceSpecShrinkRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *ModifyDBInstanceSpecShrinkRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *ModifyDBInstanceSpecShrinkRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *ModifyDBInstanceSpecShrinkRequest
	GetDBInstanceStorageType() *string
	SetDedicatedHostGroupId(v string) *ModifyDBInstanceSpecShrinkRequest
	GetDedicatedHostGroupId() *string
	SetDirection(v string) *ModifyDBInstanceSpecShrinkRequest
	GetDirection() *string
	SetEffectiveTime(v string) *ModifyDBInstanceSpecShrinkRequest
	GetEffectiveTime() *string
	SetEngineVersion(v string) *ModifyDBInstanceSpecShrinkRequest
	GetEngineVersion() *string
	SetIoAccelerationEnabled(v string) *ModifyDBInstanceSpecShrinkRequest
	GetIoAccelerationEnabled() *string
	SetOptimizedWrites(v string) *ModifyDBInstanceSpecShrinkRequest
	GetOptimizedWrites() *string
	SetOwnerAccount(v string) *ModifyDBInstanceSpecShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceSpecShrinkRequest
	GetOwnerId() *int64
	SetPayType(v string) *ModifyDBInstanceSpecShrinkRequest
	GetPayType() *string
	SetPromotionCode(v string) *ModifyDBInstanceSpecShrinkRequest
	GetPromotionCode() *string
	SetReadOnlyDBInstanceClass(v string) *ModifyDBInstanceSpecShrinkRequest
	GetReadOnlyDBInstanceClass() *string
	SetResourceGroupId(v string) *ModifyDBInstanceSpecShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceSpecShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceSpecShrinkRequest
	GetResourceOwnerId() *int64
	SetServerlessConfigurationShrink(v string) *ModifyDBInstanceSpecShrinkRequest
	GetServerlessConfigurationShrink() *string
	SetSourceBiz(v string) *ModifyDBInstanceSpecShrinkRequest
	GetSourceBiz() *string
	SetSwitchTime(v string) *ModifyDBInstanceSpecShrinkRequest
	GetSwitchTime() *string
	SetTargetMinorVersion(v string) *ModifyDBInstanceSpecShrinkRequest
	GetTargetMinorVersion() *string
	SetUsedTime(v int64) *ModifyDBInstanceSpecShrinkRequest
	GetUsedTime() *int64
	SetVSwitchId(v string) *ModifyDBInstanceSpecShrinkRequest
	GetVSwitchId() *string
	SetZoneId(v string) *ModifyDBInstanceSpecShrinkRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *ModifyDBInstanceSpecShrinkRequest
	GetZoneIdSlave1() *string
}

type ModifyDBInstanceSpecShrinkRequest struct {
	AllocateStrategy         *string `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	AllowMajorVersionUpgrade *bool   `json:"AllowMajorVersionUpgrade,omitempty" xml:"AllowMajorVersionUpgrade,omitempty"`
	AutoUseCoupon            *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	BurstingEnabled          *bool   `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Category                 *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ColdDataEnabled          *bool   `json:"ColdDataEnabled,omitempty" xml:"ColdDataEnabled,omitempty"`
	CompressionMode          *string `json:"CompressionMode,omitempty" xml:"CompressionMode,omitempty"`
	DBInstanceClass          *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// This parameter is required.
	DBInstanceId                  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStorage             *int32  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceStorageType         *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	DedicatedHostGroupId          *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	Direction                     *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	EffectiveTime                 *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	EngineVersion                 *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	IoAccelerationEnabled         *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	OptimizedWrites               *string `json:"OptimizedWrites,omitempty" xml:"OptimizedWrites,omitempty"`
	OwnerAccount                  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType                       *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PromotionCode                 *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	ReadOnlyDBInstanceClass       *string `json:"ReadOnlyDBInstanceClass,omitempty" xml:"ReadOnlyDBInstanceClass,omitempty"`
	ResourceGroupId               *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount          *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId               *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServerlessConfigurationShrink *string `json:"ServerlessConfiguration,omitempty" xml:"ServerlessConfiguration,omitempty"`
	SourceBiz                     *string `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
	SwitchTime                    *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	TargetMinorVersion            *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	UsedTime                      *int64  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VSwitchId                     *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                        *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneIdSlave1                  *string `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
}

func (s ModifyDBInstanceSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetAllocateStrategy() *string {
	return s.AllocateStrategy
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetAllowMajorVersionUpgrade() *bool {
	return s.AllowMajorVersionUpgrade
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetCompressionMode() *string {
	return s.CompressionMode
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetOptimizedWrites() *string {
	return s.OptimizedWrites
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetPayType() *string {
	return s.PayType
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetReadOnlyDBInstanceClass() *string {
	return s.ReadOnlyDBInstanceClass
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetServerlessConfigurationShrink() *string {
	return s.ServerlessConfigurationShrink
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetUsedTime() *int64 {
	return s.UsedTime
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBInstanceSpecShrinkRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetAllocateStrategy(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.AllocateStrategy = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetAllowMajorVersionUpgrade(v bool) *ModifyDBInstanceSpecShrinkRequest {
	s.AllowMajorVersionUpgrade = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetAutoUseCoupon(v bool) *ModifyDBInstanceSpecShrinkRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetBurstingEnabled(v bool) *ModifyDBInstanceSpecShrinkRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetCategory(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.Category = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetColdDataEnabled(v bool) *ModifyDBInstanceSpecShrinkRequest {
	s.ColdDataEnabled = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetCompressionMode(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.CompressionMode = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetDBInstanceClass(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetDBInstanceId(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetDBInstanceStorage(v int32) *ModifyDBInstanceSpecShrinkRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetDBInstanceStorageType(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetDedicatedHostGroupId(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetDirection(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.Direction = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetEffectiveTime(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetEngineVersion(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.EngineVersion = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetIoAccelerationEnabled(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetOptimizedWrites(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.OptimizedWrites = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetOwnerAccount(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetOwnerId(v int64) *ModifyDBInstanceSpecShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetPayType(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.PayType = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetPromotionCode(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetReadOnlyDBInstanceClass(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.ReadOnlyDBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetResourceGroupId(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceSpecShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetServerlessConfigurationShrink(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.ServerlessConfigurationShrink = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetSourceBiz(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.SourceBiz = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetSwitchTime(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetTargetMinorVersion(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetUsedTime(v int64) *ModifyDBInstanceSpecShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetVSwitchId(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetZoneId(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) SetZoneIdSlave1(v string) *ModifyDBInstanceSpecShrinkRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *ModifyDBInstanceSpecShrinkRequest) Validate() error {
	return dara.Validate(s)
}
