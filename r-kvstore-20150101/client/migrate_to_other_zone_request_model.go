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
	// The ID of the Tair (Redis OSS-compatible) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The time when the database is switched after the instance is migrated. Valid values:
	//
	// 	- **Immediately**: The database is immediately switched after the instance is migrated.
	//
	// 	- **MaintainTime**: The database is switched within the maintenance window.
	//
	// >  Default value: Immediately.
	//
	// example:
	//
	// Immediately
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of read replicas in the primary zone.
	//
	// >
	//
	// 	- The **ReadOnlyCount*	- and **SlaveReadOnlyCount*	- parameters are applicable only to cloud-native instances for which read/write splitting is enabled. When you migrate an instance to multiple zones, you can use these parameters to adjust the distribution of read replicas in the primary and secondary zones of the instance. This operation does not allow you to increase or decrease the number of nodes. Therefore, the sum of the values of `ReadOnlyCount and SlaveReadOnlyCount` must be the same as that before the migration.
	//
	// 	- If you do not specify these parameters when you migrate an instance from a single zone to multiple zones, one read replica is migrated to the secondary zone, and all other read replicas remain in the primary zone.
	//
	// 	- If the instance is a cluster instance, the preceding parameters indicate the number of read replicas per shard in the primary and secondary zones of the instance.
	//
	// example:
	//
	// 1
	ReadOnlyCount *int32 `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	// The number of replica nodes in the primary zone.
	//
	// >
	//
	// 	- The **ReplicaCount*	- and **SlaveReplicaCount*	- parameters are applicable only to cloud-native instances. When you migrate an instance to multiple zones, you can use these parameters to adjust the distribution of replica nodes in the primary and secondary zones of the instance. This operation does not allow you to increase or decrease the number of nodes. Therefore, the sum of the values of `ReplicaCount and SlaveReplicaCount` must be the same as that before the migration.
	//
	// 	- If you do not specify these parameters when you migrate an instance from a single zone to multiple zones, one replica node is migrated to the secondary zone, and all other replica nodes remain in the primary zone.
	//
	// 	- If the instance is a cluster instance, the preceding parameters indicate the number of replica nodes per shard in the primary and secondary zones of the instance.
	//
	// example:
	//
	// 1
	ReplicaCount         *int32  `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the secondary zone to which you want to migrate the instance. You can call the [DescribeZones](https://help.aliyun.com/document_detail/473764.html) operation to query zone IDs.
	//
	// >  If you specify this parameter, the master node and replica node of the instance can be deployed in different zones and disaster recovery is implemented across zones. The instance can withstand failures in data centers.
	//
	// example:
	//
	// cn-hangzhou-h
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of read replicas in the secondary zone.
	//
	// example:
	//
	// 1
	SlaveReadOnlyCount *int32 `json:"SlaveReadOnlyCount,omitempty" xml:"SlaveReadOnlyCount,omitempty"`
	// The number of replica nodes in the secondary zone.
	//
	// example:
	//
	// 1
	SlaveReplicaCount *int32 `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	// The ID of the vSwitch.
	//
	// >
	//
	// 	- The zone where the vSwitch resides must be the same as the ID of the destination zone.
	//
	// 	- If the network type of the instance is VPC, this parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the destination primary zone. You can call the [DescribeZones](https://help.aliyun.com/document_detail/473764.html) operation to query zone IDs.
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
