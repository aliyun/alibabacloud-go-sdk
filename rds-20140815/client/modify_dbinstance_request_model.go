// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUseCoupon(v bool) *ModifyDBInstanceRequest
	GetAutoUseCoupon() *bool
	SetBurstingEnabled(v bool) *ModifyDBInstanceRequest
	GetBurstingEnabled() *bool
	SetCategory(v string) *ModifyDBInstanceRequest
	GetCategory() *string
	SetColdDataEnabled(v bool) *ModifyDBInstanceRequest
	GetColdDataEnabled() *bool
	SetDBInstanceClass(v string) *ModifyDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *ModifyDBInstanceRequest
	GetDBInstanceId() *string
	SetDBInstanceStorage(v int32) *ModifyDBInstanceRequest
	GetDBInstanceStorage() *int32
	SetDBInstanceStorageType(v string) *ModifyDBInstanceRequest
	GetDBInstanceStorageType() *string
	SetDBNodes(v []*ModifyDBInstanceRequestDBNodes) *ModifyDBInstanceRequest
	GetDBNodes() []*ModifyDBInstanceRequestDBNodes
	SetDirection(v string) *ModifyDBInstanceRequest
	GetDirection() *string
	SetEffectiveTime(v string) *ModifyDBInstanceRequest
	GetEffectiveTime() *string
	SetIoAccelerationEnabled(v string) *ModifyDBInstanceRequest
	GetIoAccelerationEnabled() *string
	SetOwnerAccount(v string) *ModifyDBInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *ModifyDBInstanceRequest
	GetParameterGroupId() *string
	SetParameters(v map[string]*string) *ModifyDBInstanceRequest
	GetParameters() map[string]*string
	SetPromotionCode(v string) *ModifyDBInstanceRequest
	GetPromotionCode() *string
	SetResourceGroupId(v string) *ModifyDBInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *ModifyDBInstanceRequest
	GetSwitchTime() *string
	SetTargetMinorVersion(v string) *ModifyDBInstanceRequest
	GetTargetMinorVersion() *string
}

type ModifyDBInstanceRequest struct {
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// true
	ColdDataEnabled *bool `json:"ColdDataEnabled,omitempty" xml:"ColdDataEnabled,omitempty"`
	// example:
	//
	// pg.n4.2c.1m
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp15i4hn07r******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 500
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// example:
	//
	// cloud_essd
	DBInstanceStorageType *string                           `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	DBNodes               []*ModifyDBInstanceRequestDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	// example:
	//
	// Up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// Immediate
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// example:
	//
	// 0
	IoAccelerationEnabled *string `json:"IoAccelerationEnabled,omitempty" xml:"IoAccelerationEnabled,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// rpg-dp****
	ParameterGroupId *string            `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	Parameters       map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// aliwood-1688-mobile-promotion
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2019-10-17T18:50:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// example:
	//
	// rds_postgres_1200_20200830
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
}

func (s ModifyDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *ModifyDBInstanceRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyDBInstanceRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifyDBInstanceRequest) GetColdDataEnabled() *bool {
	return s.ColdDataEnabled
}

func (s *ModifyDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *ModifyDBInstanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceRequest) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *ModifyDBInstanceRequest) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *ModifyDBInstanceRequest) GetDBNodes() []*ModifyDBInstanceRequestDBNodes {
	return s.DBNodes
}

func (s *ModifyDBInstanceRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModifyDBInstanceRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBInstanceRequest) GetIoAccelerationEnabled() *string {
	return s.IoAccelerationEnabled
}

func (s *ModifyDBInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyDBInstanceRequest) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *ModifyDBInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyDBInstanceRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *ModifyDBInstanceRequest) SetAutoUseCoupon(v bool) *ModifyDBInstanceRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetBurstingEnabled(v bool) *ModifyDBInstanceRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetCategory(v string) *ModifyDBInstanceRequest {
	s.Category = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetColdDataEnabled(v bool) *ModifyDBInstanceRequest {
	s.ColdDataEnabled = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetDBInstanceClass(v string) *ModifyDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetDBInstanceId(v string) *ModifyDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetDBInstanceStorage(v int32) *ModifyDBInstanceRequest {
	s.DBInstanceStorage = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetDBInstanceStorageType(v string) *ModifyDBInstanceRequest {
	s.DBInstanceStorageType = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetDBNodes(v []*ModifyDBInstanceRequestDBNodes) *ModifyDBInstanceRequest {
	s.DBNodes = v
	return s
}

func (s *ModifyDBInstanceRequest) SetDirection(v string) *ModifyDBInstanceRequest {
	s.Direction = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetEffectiveTime(v string) *ModifyDBInstanceRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetIoAccelerationEnabled(v string) *ModifyDBInstanceRequest {
	s.IoAccelerationEnabled = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetOwnerAccount(v string) *ModifyDBInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetOwnerId(v int64) *ModifyDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetParameterGroupId(v string) *ModifyDBInstanceRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetParameters(v map[string]*string) *ModifyDBInstanceRequest {
	s.Parameters = v
	return s
}

func (s *ModifyDBInstanceRequest) SetPromotionCode(v string) *ModifyDBInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetResourceGroupId(v string) *ModifyDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetSwitchTime(v string) *ModifyDBInstanceRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDBInstanceRequest) SetTargetMinorVersion(v string) *ModifyDBInstanceRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *ModifyDBInstanceRequest) Validate() error {
	if s.DBNodes != nil {
		for _, item := range s.DBNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBInstanceRequestDBNodes struct {
	// example:
	//
	// 28542293
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// Master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// vsw-bp1g7uym6ia6yroes6dkm
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-shanghai-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyDBInstanceRequestDBNodes) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceRequestDBNodes) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceRequestDBNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ModifyDBInstanceRequestDBNodes) GetRole() *string {
	return s.Role
}

func (s *ModifyDBInstanceRequestDBNodes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBInstanceRequestDBNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBInstanceRequestDBNodes) SetNodeId(v string) *ModifyDBInstanceRequestDBNodes {
	s.NodeId = &v
	return s
}

func (s *ModifyDBInstanceRequestDBNodes) SetRole(v string) *ModifyDBInstanceRequestDBNodes {
	s.Role = &v
	return s
}

func (s *ModifyDBInstanceRequestDBNodes) SetVSwitchId(v string) *ModifyDBInstanceRequestDBNodes {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceRequestDBNodes) SetZoneId(v string) *ModifyDBInstanceRequestDBNodes {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBInstanceRequestDBNodes) Validate() error {
	return dara.Validate(s)
}
