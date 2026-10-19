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
  // The instance type.
  // 
  // > This parameter is required when you evaluate resources for a replica set instance. For more information about instance types, see [Instance types](https://help.aliyun.com/document_detail/57141.html).
  // 
  // example:
  // 
  // dds.mongo.mid
  DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
  // The instance ID. This parameter is required when you evaluate resources for changing instance specifications.
  // 
  // example:
  // 
  // dds-bp14bf67a76d****
  DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
  // The database engine. Set the value to **MongoDB**.
  // 
  // example:
  // 
  // MongoDB
  Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
  // The database engine version.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 4.2
  EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The number of read-only nodes in the instance. Valid values: **1*	- to **5**.
  // 
  // > This parameter is not required for standalone instances<ph props ="china"> or serverless instances</ph>.
  // 
  // example:
  // 
  // 1
  ReadonlyReplicas *string `json:"ReadonlyReplicas,omitempty" xml:"ReadonlyReplicas,omitempty"`
  // The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) to query available regions.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The number of nodes in the instance.
  // 
  // - For a standalone instance, set the value to **1**.
  // 
  // - For an instance with shared storage, set the value to **2**.
  // 
  // - For a replica set instance, valid values are **3**, **5**, and **7**.
  // 
  // <props="china">
  // 
  // > This parameter is not required for serverless instances.
  // 
  // example:
  // 
  // 3
  ReplicationFactor *string `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The shard information of the sharded cluster instance. This parameter is required when you evaluate resources for a sharded cluster instance.
  // 
  // When you evaluate resources for a new sharded cluster instance, you must specify the specifications of each shard in a JSON-formatted character string. Example:
  // 
  // ```
  // 
  // {
  // 
  //      "ConfigSvrs":
  // 
  //          [{"Storage":20,"DBInstanceClass":"dds.cs.mid"}],
  // 
  //      "Mongos":
  // 
  //          [{"DBInstanceClass":"dds.mongos.standard"},{"DBInstanceClass":"dds.mongos.standard"}],
  // 
  //      "Shards":
  // 
  //          [{"Storage":50,"DBInstanceClass":"dds.shard.standard"},{"Storage":50,"DBInstanceClass":"dds.shard.standard"},   {"Storage":50,"DBInstanceClass":"dds.shard.standard"}]
  // 
  //  }
  // 
  // ```
  // 
  // The parameter descriptions in the example are as follows:
  // 
  //  - ConfigSvrs: the Configserver nodes.
  // 
  //  - Mongos: the mongos nodes.
  // 
  //  - Shards: the shard nodes.
  // 
  //  - Storage: the storage capacity of the target shard.
  // 
  //  - DBInstanceClass: the instance type of the target shard. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
  // 
  // When you evaluate resources for upgrading or downgrading a sharded cluster instance, you only need to specify the node information in a JSON-formatted character string. Example:
  // 
  // ```
  // 
  // {
  // 
  //      "NodeId": "d-bp147c4d9ca7****", "NodeClass": "dds.shard.standard"
  // 
  // } 
  // 
  // ```
  // 
  // The parameter descriptions in the example are as follows:
  // 
  // - NodeId: the ID of the target node.
  // 
  // - NodeClass: the instance type of the target node. For more information, see [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html).
  // 
  // example:
  // 
  // {"NodeId": "d-bp147c4d9ca7****", "NodeClass": "dds.shard.standard"}
  ShardsInfo *string `json:"ShardsInfo,omitempty" xml:"ShardsInfo,omitempty"`
  // The storage capacity of the replica set instance. Unit: GB.
  // 
  // > This parameter is required when the instance uses cloud disks.
  // 
  // example:
  // 
  // 10
  Storage *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
  // The zone ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) to query available zones.
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

