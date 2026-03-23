// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBProxyInstanceShrinkRequest
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *ModifyDBProxyInstanceShrinkRequest
	GetDBProxyEngineType() *string
	SetDBProxyInstanceNum(v string) *ModifyDBProxyInstanceShrinkRequest
	GetDBProxyInstanceNum() *string
	SetDBProxyInstanceType(v string) *ModifyDBProxyInstanceShrinkRequest
	GetDBProxyInstanceType() *string
	SetDBProxyNodesShrink(v string) *ModifyDBProxyInstanceShrinkRequest
	GetDBProxyNodesShrink() *string
	SetEffectiveSpecificTime(v string) *ModifyDBProxyInstanceShrinkRequest
	GetEffectiveSpecificTime() *string
	SetEffectiveTime(v string) *ModifyDBProxyInstanceShrinkRequest
	GetEffectiveTime() *string
	SetMigrateAZShrink(v string) *ModifyDBProxyInstanceShrinkRequest
	GetMigrateAZShrink() *string
	SetOwnerId(v int64) *ModifyDBProxyInstanceShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDBProxyInstanceShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDBProxyInstanceShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBProxyInstanceShrinkRequest
	GetResourceOwnerId() *int64
	SetVSwitchIds(v string) *ModifyDBProxyInstanceShrinkRequest
	GetVSwitchIds() *string
}

type ModifyDBProxyInstanceShrinkRequest struct {
	// This parameter is required.
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// This parameter is required.
	DBProxyInstanceNum *string `json:"DBProxyInstanceNum,omitempty" xml:"DBProxyInstanceNum,omitempty"`
	// This parameter is required.
	DBProxyInstanceType   *string `json:"DBProxyInstanceType,omitempty" xml:"DBProxyInstanceType,omitempty"`
	DBProxyNodesShrink    *string `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty"`
	EffectiveSpecificTime *string `json:"EffectiveSpecificTime,omitempty" xml:"EffectiveSpecificTime,omitempty"`
	EffectiveTime         *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	MigrateAZShrink       *string `json:"MigrateAZ,omitempty" xml:"MigrateAZ,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VSwitchIds            *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
}

func (s ModifyDBProxyInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetDBProxyInstanceNum() *string {
	return s.DBProxyInstanceNum
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetDBProxyInstanceType() *string {
	return s.DBProxyInstanceType
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetDBProxyNodesShrink() *string {
	return s.DBProxyNodesShrink
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetEffectiveSpecificTime() *string {
	return s.EffectiveSpecificTime
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetMigrateAZShrink() *string {
	return s.MigrateAZShrink
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBProxyInstanceShrinkRequest) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetDBInstanceId(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetDBProxyEngineType(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetDBProxyInstanceNum(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.DBProxyInstanceNum = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetDBProxyInstanceType(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.DBProxyInstanceType = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetDBProxyNodesShrink(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.DBProxyNodesShrink = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetEffectiveSpecificTime(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.EffectiveSpecificTime = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetEffectiveTime(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetMigrateAZShrink(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.MigrateAZShrink = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetOwnerId(v int64) *ModifyDBProxyInstanceShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetRegionId(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetResourceOwnerAccount(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetResourceOwnerId(v int64) *ModifyDBProxyInstanceShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) SetVSwitchIds(v string) *ModifyDBProxyInstanceShrinkRequest {
	s.VSwitchIds = &v
	return s
}

func (s *ModifyDBProxyInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
