// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunUid(v string) *ModifyDtsJobEndpointRequest
	GetAliyunUid() *string
	SetDatabase(v string) *ModifyDtsJobEndpointRequest
	GetDatabase() *string
	SetDryRun(v bool) *ModifyDtsJobEndpointRequest
	GetDryRun() *bool
	SetDtsInstanceId(v string) *ModifyDtsJobEndpointRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ModifyDtsJobEndpointRequest
	GetDtsJobId() *string
	SetEndpoint(v string) *ModifyDtsJobEndpointRequest
	GetEndpoint() *string
	SetEndpointInstanceId(v string) *ModifyDtsJobEndpointRequest
	GetEndpointInstanceId() *string
	SetEndpointInstanceType(v string) *ModifyDtsJobEndpointRequest
	GetEndpointInstanceType() *string
	SetEndpointIp(v string) *ModifyDtsJobEndpointRequest
	GetEndpointIp() *string
	SetEndpointPort(v string) *ModifyDtsJobEndpointRequest
	GetEndpointPort() *string
	SetEndpointRegionId(v string) *ModifyDtsJobEndpointRequest
	GetEndpointRegionId() *string
	SetModifyAccount(v bool) *ModifyDtsJobEndpointRequest
	GetModifyAccount() *bool
	SetPassword(v string) *ModifyDtsJobEndpointRequest
	GetPassword() *string
	SetRegionId(v string) *ModifyDtsJobEndpointRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDtsJobEndpointRequest
	GetResourceGroupId() *string
	SetRoleName(v string) *ModifyDtsJobEndpointRequest
	GetRoleName() *string
	SetShardPassword(v string) *ModifyDtsJobEndpointRequest
	GetShardPassword() *string
	SetShardUsername(v string) *ModifyDtsJobEndpointRequest
	GetShardUsername() *string
	SetSynchronizationDirection(v string) *ModifyDtsJobEndpointRequest
	GetSynchronizationDirection() *string
	SetUsername(v string) *ModifyDtsJobEndpointRequest
	GetUsername() *string
	SetZeroEtlJob(v bool) *ModifyDtsJobEndpointRequest
	GetZeroEtlJob() *bool
}

type ModifyDtsJobEndpointRequest struct {
	// The ID of the Alibaba Cloud account (primary account) to which the database instance belongs.
	//
	// >  Passing this parameter indicates that cross-Alibaba Cloud account data synchronization will be performed, and you also need to pass the **RoleName*	- parameter.
	//
	// example:
	//
	// 150780020300****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// When the database type is **PostgreSQL**, **PolarDB for PostgreSQL**, or **AnalyticDB PostgreSQL**, it represents the database name; when the database type is **MongoDB**, it represents the authentication database name.
	//
	// > This parameter is only available and must be provided when the database type is **PostgreSQL**, **PolarDB for PostgreSQL**, **AnalyticDB PostgreSQL**, or **MongoDB**.
	//
	// example:
	//
	// admin
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// Specifies whether to perform only a precheck. Valid values:
	//
	// 	- **true**: Yes. After the precheck is passed, the database is not changed.
	//
	// 	- **false*	- (default): No. After the precheck is passed, the system changes the original database of the DTS task and runs the task.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the DTS instance. If this parameter is not provided, **DtsJobId*	- must be specified.
	//
	// example:
	//
	// dtsaw012y2g15q****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// DTS job ID, which can be queried by calling [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html).
	//
	// > If this parameter is not provided, **DtsInstanceId*	- must be filled in.
	//
	// example:
	//
	// m4312mab158****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The database instance to be modified, with values:
	//
	// - **src**: Source database instance. - **dest**: Target database instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// src
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// ID of the database instance.
	//
	// example:
	//
	// rm-bp10k50h8374w****
	EndpointInstanceId *string `json:"EndpointInstanceId,omitempty" xml:"EndpointInstanceId,omitempty"`
	// The type of the database. Valid values:
	//
	// 	- **rds**: ApsaraDB RDS for MySQL instance, ApsaraDB RDS for SQL Server instance, or ApsaraDB RDS for PostgreSQL instance.
	//
	// 	- **polardb**: PolarDB for MySQL cluster or PolarDB for PostgreSQL cluster.
	//
	// 	- **mongodb**: ApsaraDB for MongoDB replica set instance.
	//
	// 	- **distributed_mongodb**: ApsaraDB for MongoDB sharded cluster instance.
	//
	// 	- **greenplum**: AnalyticDB for PostgreSQL instance.
	//
	// 	- **kafka**: ApsaraMQ for Kafka instance.
	//
	// 	- **ecs**: self-managed database that is hosted on an Elastic Compute Service (ECS) instance. If you set this parameter to ecs, the database must be the supported one.
	//
	// 	- **express**: database that is connected over Express Connect. If you set this parameter to express, the database must be the supported one.
	//
	// 	- **other**: database that is connected over Internet. If you set this parameter to other, the database must be the supported one.
	//
	// >
	//
	// 	- The following types of databases are supported: **MySQL**, **PolarDB for MySQL**, **PostgreSQL**, **PolarDB for PostgreSQL**, **MongoDB**, **SQL Server**, **Kafka**, and **AnalyticDB for PostgreSQL**.
	//
	// 	- If the original database is an ApsaraDB for MongoDB sharded cluster instance, the new database must have the same number of shards as the original database.
	//
	// 	- If the database that you want to change is a source **PostgreSQL*	- database, you must make sure that the latency of the DTS instance is less than 30 seconds and no data is written to the source database during the change. Otherwise, data inconsistency may occur.
	//
	// 	- The value of this parameter is case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	EndpointInstanceType *string `json:"EndpointInstanceType,omitempty" xml:"EndpointInstanceType,omitempty"`
	// The IP of the database instance.
	//
	// example:
	//
	// 172.168.XX.XXX
	EndpointIp *string `json:"EndpointIp,omitempty" xml:"EndpointIp,omitempty"`
	// port of the database instance.
	//
	// example:
	//
	// 3306
	EndpointPort *string `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	// The ID of the region in which the database resides.
	//
	// example:
	//
	// cn-hangzhou
	EndpointRegionId *string `json:"EndpointRegionId,omitempty" xml:"EndpointRegionId,omitempty"`
	// Specifies whether to change the password of the database account. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ModifyAccount *bool `json:"ModifyAccount,omitempty" xml:"ModifyAccount,omitempty"`
	// The password of the database account.
	//
	// >  This parameter is valid only if **ModifyAccount*	- is set to **true**.
	//
	// example:
	//
	// DTStest****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the region in which the DTS instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Cross Alibaba Cloud account role name. When performing data synchronization across Alibaba Cloud accounts, this parameter must be passed. For the required permissions and authorization methods for this role, please refer to [How to Configure RAM Authorization for Cross-Account Data Migration or Synchronization](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The account password of the shard of the ApsaraDB for MongoDB sharded cluster instance.
	//
	// >
	//
	// 	- This parameter is valid and required only if the source database is an ApsaraDB for MongoDB sharded cluster instance.
	//
	// 	- This parameter is valid only if **ModifyAccount*	- is set to **true**.
	//
	// example:
	//
	// DTStest****
	ShardPassword *string `json:"ShardPassword,omitempty" xml:"ShardPassword,omitempty"`
	// The account username of the shard of the ApsaraDB for MongoDB sharded cluster instance.
	//
	// >
	//
	// 	- This parameter is valid and required only if the source database is an ApsaraDB for MongoDB sharded cluster instance.
	//
	// 	- This parameter is valid only if **ModifyAccount*	- is set to **true**.
	//
	// example:
	//
	// shard
	ShardUsername *string `json:"ShardUsername,omitempty" xml:"ShardUsername,omitempty"`
	// Synchronization direction, with values:
	//
	// - **Forward*	- (default): Forward. - **Reverse**: Reverse.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The database account.
	//
	// >  This parameter is valid only if **ModifyAccount*	- is set to **true**.
	//
	// example:
	//
	// dtstest
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
	ZeroEtlJob *bool   `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s ModifyDtsJobEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobEndpointRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobEndpointRequest) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *ModifyDtsJobEndpointRequest) GetDatabase() *string {
	return s.Database
}

func (s *ModifyDtsJobEndpointRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDtsJobEndpointRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ModifyDtsJobEndpointRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifyDtsJobEndpointRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ModifyDtsJobEndpointRequest) GetEndpointInstanceId() *string {
	return s.EndpointInstanceId
}

func (s *ModifyDtsJobEndpointRequest) GetEndpointInstanceType() *string {
	return s.EndpointInstanceType
}

func (s *ModifyDtsJobEndpointRequest) GetEndpointIp() *string {
	return s.EndpointIp
}

func (s *ModifyDtsJobEndpointRequest) GetEndpointPort() *string {
	return s.EndpointPort
}

func (s *ModifyDtsJobEndpointRequest) GetEndpointRegionId() *string {
	return s.EndpointRegionId
}

func (s *ModifyDtsJobEndpointRequest) GetModifyAccount() *bool {
	return s.ModifyAccount
}

func (s *ModifyDtsJobEndpointRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyDtsJobEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDtsJobEndpointRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDtsJobEndpointRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *ModifyDtsJobEndpointRequest) GetShardPassword() *string {
	return s.ShardPassword
}

func (s *ModifyDtsJobEndpointRequest) GetShardUsername() *string {
	return s.ShardUsername
}

func (s *ModifyDtsJobEndpointRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ModifyDtsJobEndpointRequest) GetUsername() *string {
	return s.Username
}

func (s *ModifyDtsJobEndpointRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *ModifyDtsJobEndpointRequest) SetAliyunUid(v string) *ModifyDtsJobEndpointRequest {
	s.AliyunUid = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetDatabase(v string) *ModifyDtsJobEndpointRequest {
	s.Database = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetDryRun(v bool) *ModifyDtsJobEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetDtsInstanceId(v string) *ModifyDtsJobEndpointRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetDtsJobId(v string) *ModifyDtsJobEndpointRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetEndpoint(v string) *ModifyDtsJobEndpointRequest {
	s.Endpoint = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetEndpointInstanceId(v string) *ModifyDtsJobEndpointRequest {
	s.EndpointInstanceId = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetEndpointInstanceType(v string) *ModifyDtsJobEndpointRequest {
	s.EndpointInstanceType = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetEndpointIp(v string) *ModifyDtsJobEndpointRequest {
	s.EndpointIp = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetEndpointPort(v string) *ModifyDtsJobEndpointRequest {
	s.EndpointPort = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetEndpointRegionId(v string) *ModifyDtsJobEndpointRequest {
	s.EndpointRegionId = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetModifyAccount(v bool) *ModifyDtsJobEndpointRequest {
	s.ModifyAccount = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetPassword(v string) *ModifyDtsJobEndpointRequest {
	s.Password = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetRegionId(v string) *ModifyDtsJobEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetResourceGroupId(v string) *ModifyDtsJobEndpointRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetRoleName(v string) *ModifyDtsJobEndpointRequest {
	s.RoleName = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetShardPassword(v string) *ModifyDtsJobEndpointRequest {
	s.ShardPassword = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetShardUsername(v string) *ModifyDtsJobEndpointRequest {
	s.ShardUsername = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetSynchronizationDirection(v string) *ModifyDtsJobEndpointRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetUsername(v string) *ModifyDtsJobEndpointRequest {
	s.Username = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetZeroEtlJob(v bool) *ModifyDtsJobEndpointRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) Validate() error {
	return dara.Validate(s)
}
