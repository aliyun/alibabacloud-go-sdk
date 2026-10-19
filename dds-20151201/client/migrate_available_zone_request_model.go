// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateAvailableZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *MigrateAvailableZoneRequest
	GetDBInstanceId() *string
	SetEffectiveTime(v string) *MigrateAvailableZoneRequest
	GetEffectiveTime() *string
	SetHiddenZoneId(v string) *MigrateAvailableZoneRequest
	GetHiddenZoneId() *string
	SetOwnerAccount(v string) *MigrateAvailableZoneRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *MigrateAvailableZoneRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *MigrateAvailableZoneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MigrateAvailableZoneRequest
	GetResourceOwnerId() *int64
	SetSecondaryZoneId(v string) *MigrateAvailableZoneRequest
	GetSecondaryZoneId() *string
	SetVswitch(v string) *MigrateAvailableZoneRequest
	GetVswitch() *string
	SetZoneId(v string) *MigrateAvailableZoneRequest
	GetZoneId() *string
}

type MigrateAvailableZoneRequest struct {
	// The instance ID.
	//
	// > If the network type of the instance is VPC, you must also specify the **Vswitch*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1ece71ff2f****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The effective time for migrating the zone. The instance must be active. Valid values:
	//
	// - **Immediately**: Takes effect immediately.
	//
	// - **MaintainTime**: Takes effect during the maintenance window of the instance.
	//
	// Default value: **Immediately**.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The ID of destination secondary zone 2.
	//
	// example:
	//
	// cn-shanghai-n
	HiddenZoneId         *string `json:"HiddenZoneId,omitempty" xml:"HiddenZoneId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of destination secondary zone 1.
	//
	// example:
	//
	// cn-hangzhou-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The ID of the vSwitch in the destination zone.
	//
	// > This parameter is required when the network type of the instance is VPC.
	//
	// example:
	//
	// vsw-bp1buy0h9myt5i9e7****
	Vswitch *string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty"`
	// The ID of the destination zone.
	//
	// > - The destination zone and the current zone of the instance are in the same region.
	//
	// > - You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query zone IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s MigrateAvailableZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateAvailableZoneRequest) GoString() string {
	return s.String()
}

func (s *MigrateAvailableZoneRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MigrateAvailableZoneRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *MigrateAvailableZoneRequest) GetHiddenZoneId() *string {
	return s.HiddenZoneId
}

func (s *MigrateAvailableZoneRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *MigrateAvailableZoneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MigrateAvailableZoneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MigrateAvailableZoneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MigrateAvailableZoneRequest) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *MigrateAvailableZoneRequest) GetVswitch() *string {
	return s.Vswitch
}

func (s *MigrateAvailableZoneRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *MigrateAvailableZoneRequest) SetDBInstanceId(v string) *MigrateAvailableZoneRequest {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetEffectiveTime(v string) *MigrateAvailableZoneRequest {
	s.EffectiveTime = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetHiddenZoneId(v string) *MigrateAvailableZoneRequest {
	s.HiddenZoneId = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetOwnerAccount(v string) *MigrateAvailableZoneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetOwnerId(v int64) *MigrateAvailableZoneRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetResourceOwnerAccount(v string) *MigrateAvailableZoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetResourceOwnerId(v int64) *MigrateAvailableZoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetSecondaryZoneId(v string) *MigrateAvailableZoneRequest {
	s.SecondaryZoneId = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetVswitch(v string) *MigrateAvailableZoneRequest {
	s.Vswitch = &v
	return s
}

func (s *MigrateAvailableZoneRequest) SetZoneId(v string) *MigrateAvailableZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *MigrateAvailableZoneRequest) Validate() error {
	return dara.Validate(s)
}
