// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateResourceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDBInstanceClass(v string) *EvaluateResourceRequest
  GetDBInstanceClass() *string 
  SetDBInstanceId(v string) *EvaluateResourceRequest
  GetDBInstanceId() *string 
  SetEngine(v string) *EvaluateResourceRequest
  GetEngine() *string 
  SetEngineVersion(v string) *EvaluateResourceRequest
  GetEngineVersion() *string 
  SetOwnerAccount(v string) *EvaluateResourceRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *EvaluateResourceRequest
  GetOwnerId() *int64 
  SetReadonlyReplicas(v string) *EvaluateResourceRequest
  GetReadonlyReplicas() *string 
  SetRegionId(v string) *EvaluateResourceRequest
  GetRegionId() *string 
  SetReplicationFactor(v string) *EvaluateResourceRequest
  GetReplicationFactor() *string 
  SetResourceOwnerAccount(v string) *EvaluateResourceRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *EvaluateResourceRequest
  GetResourceOwnerId() *int64 
  SetShardsInfo(v string) *EvaluateResourceRequest
  GetShardsInfo() *string 
  SetStorage(v string) *EvaluateResourceRequest
  GetStorage() *string 
  SetZoneId(v string) *EvaluateResourceRequest
  GetZoneId() *string 
}

type EvaluateResourceRequest struct {
  // The type of the instance.
  // 
  // > This parameter is required when you check whether resources are sufficient for creating or upgrading a replica set instance. For more information about instance types, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
  // 
  // example:
  // 
  // dds.mongo.mid
  DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
  // The ID of the instance. This parameter is required when you check whether resources are sufficient for upgrading an instance.
  // 
  // example:
  // 
  // dds-bp14bf67a76d****
  DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
  // The database engine of the instance. Set the value to **MongoDB**.
  // 
  // example:
  // 
  // MongoDB
  Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
  // The version of the database engine.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 4.0
  EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The number of read-only nodes in the instance. Valid values: **1*	- to **5**.
  // 
  // > This parameter is not required for standalone or serverless instances.
  // 
  // example:
  // 
  // 1
  ReadonlyReplicas *string `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
  // The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the region ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The number of nodes in the instance.
  // 
  // 	- Set the value to **1*	- for standalone instances.
  // 
  // 	- Valid values for replica set instances: **3**, **5**, and **7**
  // 
  // > This parameter is not required for serverless instances.
  // 
  // example:
  // 
  // 3
  ReplicationFactor *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The node information about the sharded cluster instance. This parameter is required when you check whether resources are sufficient for creating or upgrading a sharded cluster instance.
  // 
  // To check whether resources are sufficient for creating a sharded cluster instance, specify the specifications of each node in the instance. The value must be a JSON string. Example:
  // 
  //     {
  // 
  //          "ConfigSvrs":
  // 
  //              [{"Storage":20,"DBInstanceClass":"dds.cs.mid"}],
  // 
  //          "Mongos":
  // 
  //              [{"DBInstanceClass":"dds.mongos.standard"},{"DBInstanceClass":"dds.mongos.standard"}],
  // 
  //          "Shards":
  // 
  //              [{"Storage":50,"DBInstanceClass":"dds.shard.standard"},{"Storage":50,"DBInstanceClass":"dds.shard.standard"},   {"Storage":50,"DBInstanceClass":"dds.shard.standard"}]
  // 
  //      }
  // 
  // Parameters in the example:
  // 
  // 	- ConfigSvrs: the Configserver node.
  // 
  // 	- Mongos: the mongos node.
  // 
  // 	- Shards: the shard node.
  // 
  // 	- Storage: the storage space of the node.
  // 
  // 	- DBInstanceClass: the instance type of the node. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
  // 
  // To check whether resources are sufficient for upgrading a single node of a sharded cluster instance, specify only the information about the node to be upgraded. The value must be a JSON string. Example:
  // 
  //     {
  // 
  //          "NodeId": "d-bp147c4d9ca7****", "NodeClass": "dds.shard.standard"
  // 
  //     } 
  // 
  // Parameters in the example:
  // 
  // 	- NodeId: the ID of the node.
  // 
  // 	- NodeClass: the instance type of the node. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
  // 
  // example:
  // 
  // {"NodeId": "d-bp147c4d9ca7****", "NodeClass": "dds.shard.standard"}
  ShardsInfo *string `json:"ShardsInfo,omitempty" xml:"ShardsInfo,omitempty"`
  // The storage capacity of the replica set instance. Unit: GB.
  // 
  // > This parameter is required for the instances that use cloud disks.
  // 
  // example:
  // 
  // 10
  Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
  // The zone ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the zone ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou-h
  ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s EvaluateResourceRequest) String() string {
  return dara.Prettify(s)
}

func (s EvaluateResourceRequest) GoString() string {
  return s.String()
}

func (s *EvaluateResourceRequest) GetDBInstanceClass() *string  {
  return s.DBInstanceClass
}

func (s *EvaluateResourceRequest) GetDBInstanceId() *string  {
  return s.DBInstanceId
}

func (s *EvaluateResourceRequest) GetEngine() *string  {
  return s.Engine
}

func (s *EvaluateResourceRequest) GetEngineVersion() *string  {
  return s.EngineVersion
}

func (s *EvaluateResourceRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *EvaluateResourceRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EvaluateResourceRequest) GetReadonlyReplicas() *string  {
  return s.ReadonlyReplicas
}

func (s *EvaluateResourceRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EvaluateResourceRequest) GetReplicationFactor() *string  {
  return s.ReplicationFactor
}

func (s *EvaluateResourceRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *EvaluateResourceRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EvaluateResourceRequest) GetShardsInfo() *string  {
  return s.ShardsInfo
}

func (s *EvaluateResourceRequest) GetStorage() *string  {
  return s.Storage
}

func (s *EvaluateResourceRequest) GetZoneId() *string  {
  return s.ZoneId
}

func (s *EvaluateResourceRequest) SetDBInstanceClass(v string) *EvaluateResourceRequest {
  s.DBInstanceClass = &v
  return s
}

func (s *EvaluateResourceRequest) SetDBInstanceId(v string) *EvaluateResourceRequest {
  s.DBInstanceId = &v
  return s
}

func (s *EvaluateResourceRequest) SetEngine(v string) *EvaluateResourceRequest {
  s.Engine = &v
  return s
}

func (s *EvaluateResourceRequest) SetEngineVersion(v string) *EvaluateResourceRequest {
  s.EngineVersion = &v
  return s
}

func (s *EvaluateResourceRequest) SetOwnerAccount(v string) *EvaluateResourceRequest {
  s.OwnerAccount = &v
  return s
}

func (s *EvaluateResourceRequest) SetOwnerId(v int64) *EvaluateResourceRequest {
  s.OwnerId = &v
  return s
}

func (s *EvaluateResourceRequest) SetReadonlyReplicas(v string) *EvaluateResourceRequest {
  s.ReadonlyReplicas = &v
  return s
}

func (s *EvaluateResourceRequest) SetRegionId(v string) *EvaluateResourceRequest {
  s.RegionId = &v
  return s
}

func (s *EvaluateResourceRequest) SetReplicationFactor(v string) *EvaluateResourceRequest {
  s.ReplicationFactor = &v
  return s
}

func (s *EvaluateResourceRequest) SetResourceOwnerAccount(v string) *EvaluateResourceRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *EvaluateResourceRequest) SetResourceOwnerId(v int64) *EvaluateResourceRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EvaluateResourceRequest) SetShardsInfo(v string) *EvaluateResourceRequest {
  s.ShardsInfo = &v
  return s
}

func (s *EvaluateResourceRequest) SetStorage(v string) *EvaluateResourceRequest {
  s.Storage = &v
  return s
}

func (s *EvaluateResourceRequest) SetZoneId(v string) *EvaluateResourceRequest {
  s.ZoneId = &v
  return s
}

func (s *EvaluateResourceRequest) Validate() error {
  return dara.Validate(s)
}

