// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocateStrategy(v string) *ModifyDBInstanceSpecRequest
	GetAllocateStrategy() *string
	SetAllowMajorVersionUpgrade(v bool) *ModifyDBInstanceSpecRequest
	GetAllowMajorVersionUpgrade() *bool
	SetAutoUseCoupon(v bool) *ModifyDBInstanceSpecRequest
	GetAutoUseCoupon() *bool
	SetBurstingEnabled(v bool) *ModifyDBInstanceSpecRequest
	GetBurstingEnabled() *bool
	SetCategory(v string) *ModifyDBInstanceSpecRequest
	GetCategory() *string
	SetColdDataEnabled(v bool) *ModifyDBInstanceSpecRequest
	GetColdDataEnabled() *bool
	SetCompressionMode(v string) *ModifyDBInstanceSpecRequest
	GetCompressionMode() *string
	SetDBInstanceClass(v string) *ModifyDBInstanceSpecRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *ModifyDBInstanceSpecRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *ModifyDBInstanceSpecRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *ModifyDBInstanceSpecRequest
	GetDBInstanceStorageType() *string
	SetDedicatedHostGroupId(v string) *ModifyDBInstanceSpecRequest
	GetDedicatedHostGroupId() *string
	SetDirection(v string) *ModifyDBInstanceSpecRequest
	GetDirection() *string
	SetEffectiveTime(v string) *ModifyDBInstanceSpecRequest
	GetEffectiveTime() *string
	SetEngineVersion(v string) *ModifyDBInstanceSpecRequest
	GetEngineVersion() *string
	SetIoAccelerationEnabled(v string) *ModifyDBInstanceSpecRequest
	GetIoAccelerationEnabled() *string
	SetOptimizedWrites(v string) *ModifyDBInstanceSpecRequest
	GetOptimizedWrites() *string
	SetOwnerAccount(v string) *ModifyDBInstanceSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceSpecRequest
	GetOwnerId() *int64
	SetPayType(v string) *ModifyDBInstanceSpecRequest
	GetPayType() *string
	SetPromotionCode(v string) *ModifyDBInstanceSpecRequest
	GetPromotionCode() *string
	SetReadOnlyDBInstanceClass(v string) *ModifyDBInstanceSpecRequest
	GetReadOnlyDBInstanceClass() *string
	SetResourceGroupId(v string) *ModifyDBInstanceSpecRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceSpecRequest
	GetResourceOwnerId() *int64
	SetServerlessConfiguration(v *ModifyDBInstanceSpecRequestServerlessConfiguration) *ModifyDBInstanceSpecRequest
	GetServerlessConfiguration() *ModifyDBInstanceSpecRequestServerlessConfiguration
	SetSourceBiz(v string) *ModifyDBInstanceSpecRequest
	GetSourceBiz() *string
	SetSwitchTime(v string) *ModifyDBInstanceSpecRequest
	GetSwitchTime() *string
	SetTargetMinorVersion(v string) *ModifyDBInstanceSpecRequest
	GetTargetMinorVersion() *string
	SetUsedTime(v int64) *ModifyDBInstanceSpecRequest
	GetUsedTime() *int64
	SetVSwitchId(v string) *ModifyDBInstanceSpecRequest
	GetVSwitchId() *string
	SetZoneId(v string) *ModifyDBInstanceSpecRequest
	GetZoneId() *string
	SetZoneIdSlave1(v string) *ModifyDBInstanceSpecRequest
	GetZoneIdSlave1() *string
}

type ModifyDBInstanceSpecRequest struct {
	AllocateStrategy         *string `json:"AllocateStrategy,omitempty" xml:"AllocateStrategy,omitempty"`
	AllowMajorVersionUpgrade *bool   `json:"AllowMajorVersionUpgrade,omitempty" xml:"AllowMajorVersionUpgrade,omitempty"`
	AutoUseCoupon            *bool   `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	BurstingEnabled          *bool   `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Category                 *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ColdDataEnabled          *bool   `json:"ColdDataEnabled,omitempty" xml:"ColdDataEnabled,omitempty"`
	CompressionMode          *string `json:"CompressionMode,omitempty" xml:"CompressionMode,omitempty"`
	DBInstanceClass          *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// This parameter is required.
	DBInstanceId            *string                                             `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStorage       *int32                                              `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceStorageType   *string                                             `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	DedicatedHostGroupId    *string                                             `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	Direction               *string                                             `json:"Direction,omitempty" xml:"Direction,omitempty"`
	EffectiveTime           *string                                             `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	EngineVersion           *string                                             `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	IoAccelerationEnabled   *string                                             `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	OptimizedWrites         *string                                             `json:"OptimizedWrites,omitempty" xml:"OptimizedWrites,omitempty"`
	OwnerAccount            *string                                             `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64                                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType                 *string                                             `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PromotionCode           *string                                             `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	ReadOnlyDBInstanceClass *string                                             `json:"ReadOnlyDBInstanceClass,omitempty" xml:"ReadOnlyDBInstanceClass,omitempty"`
	ResourceGroupId         *string                                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount    *string                                             `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64                                              `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ServerlessConfiguration *ModifyDBInstanceSpecRequestServerlessConfiguration `json:"ServerlessConfiguration,omitempty" xml:"ServerlessConfiguration,omitempty" type:"Struct"`
	SourceBiz               *string                                             `json:"SourceBiz,omitempty" xml:"SourceBiz,omitempty"`
	SwitchTime              *string                                             `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	TargetMinorVersion      *string                                             `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	UsedTime                *int64                                              `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VSwitchId               *string                                             `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                  *string                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneIdSlave1            *string                                             `json:"ZoneIdSlave1,omitempty" xml:"ZoneIdSlave1,omitempty"`
}

func (s ModifyDBInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecRequest) GetAllocateStrategy() *string {
	return s.AllocateStrategy
}

func (s *ModifyDBInstanceSpecRequest) GetAllowMajorVersionUpgrade() *bool {
	return s.AllowMajorVersionUpgrade
}

func (s *ModifyDBInstanceSpecRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyDBInstanceSpecRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyDBInstanceSpecRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyDBInstanceSpecRequest) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
}

func (s *ModifyDBInstanceSpecRequest) GetCompressionMode() *string {
	return s.CompressionMode
}

func (s *ModifyDBInstanceSpecRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *ModifyDBInstanceSpecRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceSpecRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *ModifyDBInstanceSpecRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *ModifyDBInstanceSpecRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *ModifyDBInstanceSpecRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyDBInstanceSpecRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBInstanceSpecRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *ModifyDBInstanceSpecRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *ModifyDBInstanceSpecRequest) GetOptimizedWrites() *string {
	return s.OptimizedWrites
}

func (s *ModifyDBInstanceSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceSpecRequest) GetPayType() *string {
	return s.PayType
}

func (s *ModifyDBInstanceSpecRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyDBInstanceSpecRequest) GetReadOnlyDBInstanceClass() *string {
	return s.ReadOnlyDBInstanceClass
}

func (s *ModifyDBInstanceSpecRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceSpecRequest) GetServerlessConfiguration() *ModifyDBInstanceSpecRequestServerlessConfiguration {
	return s.ServerlessConfiguration
}

func (s *ModifyDBInstanceSpecRequest) GetSourceBiz() *string {
	return s.SourceBiz
}

func (s *ModifyDBInstanceSpecRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyDBInstanceSpecRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *ModifyDBInstanceSpecRequest) GetUsedTime() *int64 {
	return s.UsedTime
}

func (s *ModifyDBInstanceSpecRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBInstanceSpecRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBInstanceSpecRequest) GetZoneIdSlave1() *string {
	return s.ZoneIdSlave1
}

func (s *ModifyDBInstanceSpecRequest) SetAllocateStrategy(v string) *ModifyDBInstanceSpecRequest {
	s.AllocateStrategy = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetAllowMajorVersionUpgrade(v bool) *ModifyDBInstanceSpecRequest {
	s.AllowMajorVersionUpgrade = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetAutoUseCoupon(v bool) *ModifyDBInstanceSpecRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetBurstingEnabled(v bool) *ModifyDBInstanceSpecRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetCategory(v string) *ModifyDBInstanceSpecRequest {
	s.Category = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetColdDataEnabled(v bool) *ModifyDBInstanceSpecRequest {
	s.ColdDataEnabled = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetCompressionMode(v string) *ModifyDBInstanceSpecRequest {
	s.CompressionMode = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceClass(v string) *ModifyDBInstanceSpecRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceId(v string) *ModifyDBInstanceSpecRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceStorage(v int32) *ModifyDBInstanceSpecRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDBInstanceStorageType(v string) *ModifyDBInstanceSpecRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDedicatedHostGroupId(v string) *ModifyDBInstanceSpecRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetDirection(v string) *ModifyDBInstanceSpecRequest {
	s.Direction = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetEffectiveTime(v string) *ModifyDBInstanceSpecRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetEngineVersion(v string) *ModifyDBInstanceSpecRequest {
	s.EngineVersion = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetIoAccelerationEnabled(v string) *ModifyDBInstanceSpecRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetOptimizedWrites(v string) *ModifyDBInstanceSpecRequest {
	s.OptimizedWrites = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetOwnerAccount(v string) *ModifyDBInstanceSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetOwnerId(v int64) *ModifyDBInstanceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetPayType(v string) *ModifyDBInstanceSpecRequest {
	s.PayType = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetPromotionCode(v string) *ModifyDBInstanceSpecRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetReadOnlyDBInstanceClass(v string) *ModifyDBInstanceSpecRequest {
	s.ReadOnlyDBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetResourceGroupId(v string) *ModifyDBInstanceSpecRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetServerlessConfiguration(v *ModifyDBInstanceSpecRequestServerlessConfiguration) *ModifyDBInstanceSpecRequest {
	s.ServerlessConfiguration = v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetSourceBiz(v string) *ModifyDBInstanceSpecRequest {
	s.SourceBiz = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetSwitchTime(v string) *ModifyDBInstanceSpecRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetTargetMinorVersion(v string) *ModifyDBInstanceSpecRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetUsedTime(v int64) *ModifyDBInstanceSpecRequest {
	s.UsedTime = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetVSwitchId(v string) *ModifyDBInstanceSpecRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetZoneId(v string) *ModifyDBInstanceSpecRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) SetZoneIdSlave1(v string) *ModifyDBInstanceSpecRequest {
	s.ZoneIdSlave1 = &v
	return s
}

func (s *ModifyDBInstanceSpecRequest) Validate() error {
	if s.ServerlessConfiguration != nil {
		if err := s.ServerlessConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDBInstanceSpecRequestServerlessConfiguration struct {
	// if can be null:
	// false
	AutoPause   *bool    `json:"AutoPause,omitempty" xml:"AutoPause,omitempty"`
	MaxCapacity *float64 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	MinCapacity *float64 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
	SwitchForce *bool    `json:"SwitchForce,omitempty" xml:"SwitchForce,omitempty"`
}

func (s ModifyDBInstanceSpecRequestServerlessConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSpecRequestServerlessConfiguration) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) GetAutoPause() *bool {
	return s.AutoPause
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) GetMaxCapacity() *float64 {
	return s.MaxCapacity
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) GetMinCapacity() *float64 {
	return s.MinCapacity
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) GetSwitchForce() *bool {
	return s.SwitchForce
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) SetAutoPause(v bool) *ModifyDBInstanceSpecRequestServerlessConfiguration {
	s.AutoPause = &v
	return s
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) SetMaxCapacity(v float64) *ModifyDBInstanceSpecRequestServerlessConfiguration {
	s.MaxCapacity = &v
	return s
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) SetMinCapacity(v float64) *ModifyDBInstanceSpecRequestServerlessConfiguration {
	s.MinCapacity = &v
	return s
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) SetSwitchForce(v bool) *ModifyDBInstanceSpecRequestServerlessConfiguration {
	s.SwitchForce = &v
	return s
}

func (s *ModifyDBInstanceSpecRequestServerlessConfiguration) Validate() error {
	return dara.Validate(s)
}
