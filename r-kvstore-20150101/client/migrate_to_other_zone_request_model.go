// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateToOtherZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *MigrateToOtherZoneRequest
	GetDBInstanceId() *string
	SetEffectiveTime(v string) *MigrateToOtherZoneRequest
	GetEffectiveTime() *string
	SetOwnerAccount(v string) *MigrateToOtherZoneRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *MigrateToOtherZoneRequest
	GetOwnerId() *int64
	SetReadOnlyCount(v int32) *MigrateToOtherZoneRequest
	GetReadOnlyCount() *int32
	SetReplicaCount(v int32) *MigrateToOtherZoneRequest
	GetReplicaCount() *int32
	SetResourceOwnerAccount(v string) *MigrateToOtherZoneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *MigrateToOtherZoneRequest
	GetResourceOwnerId() *int64
	SetSecondaryZoneId(v string) *MigrateToOtherZoneRequest
	GetSecondaryZoneId() *string
	SetSecurityToken(v string) *MigrateToOtherZoneRequest
	GetSecurityToken() *string
	SetSlaveReadOnlyCount(v int32) *MigrateToOtherZoneRequest
	GetSlaveReadOnlyCount() *int32
	SetSlaveReplicaCount(v int32) *MigrateToOtherZoneRequest
	GetSlaveReplicaCount() *int32
	SetVSwitchId(v string) *MigrateToOtherZoneRequest
	GetVSwitchId() *string
	SetZoneId(v string) *MigrateToOtherZoneRequest
	GetZoneId() *string
}

type MigrateToOtherZoneRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The time when the database switchover is performed after data migration. Valid values:
	//
	// 	- **Immediately**: The switchover is performed immediately after the migration.
	//
	// 	- **MaintainTime**: The switchover is performed within the maintenance window.
	//
	// > Default value: **Immediately**.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of read-only nodes in the primary zone.
	//
	// > 	- The ReadOnlyCount and SlaveReadOnlyCount parameters apply only to cloud-native instances with read/write splitting enabled. When you migrate an instance to multiple zones, you can use these parameters to adjust the distribution of read-only nodes across the primary and secondary zones. This operation does not support adding or removing nodes. Therefore, the total number of read-only nodes calculated by ReadOnlyCount + SlaveReadOnlyCount must remain the same as that before the migration.
	//
	// > 	- When you migrate an instance from a single zone to multiple zones and do not specify these parameters, one read-only node is migrated to the secondary zone by default, and the remaining read-only nodes stay in the primary zone.
	//
	// > 	- If the instance uses the cluster architecture, these parameters specify the number of read-only nodes per shard in the primary and secondary zones.
	//
	// example:
	//
	// 1
	ReadOnlyCount *int32 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	// The number of secondary nodes in the primary zone.
	//
	// > 	- The ReplicaCount and SlaveReplicaCount parameters apply only to cloud-native instances. When you migrate an instance to multiple zones, you can use these parameters to adjust the distribution of secondary nodes across the primary and secondary zones. This operation does not support adding or removing nodes. Therefore, the total number of secondary nodes calculated by ReplicaCount + SlaveReplicaCount must remain the same as that before the migration.
	//
	// > 	- When you migrate an instance from a single zone to multiple zones and do not specify these parameters, one secondary node is migrated to the secondary zone by default, and the remaining secondary nodes stay in the primary zone.
	//
	// > 	- If the instance uses the cluster architecture, these parameters specify the number of secondary nodes per shard in the primary and secondary zones.
	//
	// example:
	//
	// 1
	ReplicaCount         *int32  `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the destination secondary zone. You can call the [DescribeZones](https://help.aliyun.com/document_detail/473764.html) operation to query the zone ID.
	//
	// > If you specify this parameter, the primary and secondary data nodes are distributed across different zones. This implements cross-zone disaster recovery and tolerates data center-level failures.
	//
	// example:
	//
	// cn-hangzhou-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of read-only nodes in the secondary zone.
	//
	// example:
	//
	// 1
	SlaveReadOnlyCount *int32 `json:"SlaveReadOnlyCount,omitempty" xml:"SlaveReadOnlyCount,omitempty"`
	// The number of secondary nodes in the secondary zone.
	//
	// example:
	//
	// 1
	SlaveReplicaCount *int32 `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	// The vSwitch ID.
	//
	// > 	- The zone where the vSwitch resides must be the same as the destination zone specified by ZoneId.
	//
	// > 	- This parameter is required if the network type of the instance is VPC.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the destination primary zone. You can call the [DescribeZones](https://help.aliyun.com/document_detail/473764.html) operation to query the zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s MigrateToOtherZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateToOtherZoneRequest) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *MigrateToOtherZoneRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *MigrateToOtherZoneRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *MigrateToOtherZoneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MigrateToOtherZoneRequest) GetReadOnlyCount() *int32 {
	return s.ReadOnlyCount
}

func (s *MigrateToOtherZoneRequest) GetReplicaCount() *int32 {
	return s.ReplicaCount
}

func (s *MigrateToOtherZoneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *MigrateToOtherZoneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *MigrateToOtherZoneRequest) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *MigrateToOtherZoneRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *MigrateToOtherZoneRequest) GetSlaveReadOnlyCount() *int32 {
	return s.SlaveReadOnlyCount
}

func (s *MigrateToOtherZoneRequest) GetSlaveReplicaCount() *int32 {
	return s.SlaveReplicaCount
}

func (s *MigrateToOtherZoneRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *MigrateToOtherZoneRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *MigrateToOtherZoneRequest) SetDBInstanceId(v string) *MigrateToOtherZoneRequest {
	s.DBInstanceId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetEffectiveTime(v string) *MigrateToOtherZoneRequest {
	s.EffectiveTime = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetOwnerAccount(v string) *MigrateToOtherZoneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetOwnerId(v int64) *MigrateToOtherZoneRequest {
	s.OwnerId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetReadOnlyCount(v int32) *MigrateToOtherZoneRequest {
	s.ReadOnlyCount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetReplicaCount(v int32) *MigrateToOtherZoneRequest {
	s.ReplicaCount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetResourceOwnerAccount(v string) *MigrateToOtherZoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetResourceOwnerId(v int64) *MigrateToOtherZoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetSecondaryZoneId(v string) *MigrateToOtherZoneRequest {
	s.SecondaryZoneId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetSecurityToken(v string) *MigrateToOtherZoneRequest {
	s.SecurityToken = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetSlaveReadOnlyCount(v int32) *MigrateToOtherZoneRequest {
	s.SlaveReadOnlyCount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetSlaveReplicaCount(v int32) *MigrateToOtherZoneRequest {
	s.SlaveReplicaCount = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetVSwitchId(v string) *MigrateToOtherZoneRequest {
	s.VSwitchId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) SetZoneId(v string) *MigrateToOtherZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *MigrateToOtherZoneRequest) Validate() error {
	return dara.Validate(s)
}
