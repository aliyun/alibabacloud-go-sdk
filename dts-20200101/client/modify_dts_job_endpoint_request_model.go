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
	SetEndpointPrimaryVswId(v string) *ModifyDtsJobEndpointRequest
	GetEndpointPrimaryVswId() *string
	SetEndpointRegionId(v string) *ModifyDtsJobEndpointRequest
	GetEndpointRegionId() *string
	SetEndpointSecondaryVswId(v string) *ModifyDtsJobEndpointRequest
	GetEndpointSecondaryVswId() *string
	SetEndpointVpcId(v string) *ModifyDtsJobEndpointRequest
	GetEndpointVpcId() *string
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
	// The ID of the Alibaba Cloud account that owns the database instance.
	//
	// > Specifying this parameter indicates cross-account data synchronization. You must also specify the **RoleName*	- parameter.
	//
	// example:
	//
	// 150780020300****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The database name when the database type is **PostgreSQL**, **PolarDB for PostgreSQL**, or **AnalyticDB PostgreSQL**. The authentication database name when the database type is **MongoDB**.
	//
	// > This parameter is available and required only when the database type is **PostgreSQL**, **PolarDB for PostgreSQL**, **AnalyticDB PostgreSQL**, or **MongoDB**.
	//
	// example:
	//
	// admin
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - **true**: Yes. After the dry run succeeds, the instance is not modified.
	//
	// - **false*	- (default): No. After the dry run succeeds, the database instance of the DTS task is modified and the task runs.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the DTS instance.
	//
	// > If you do not specify this parameter, you must specify **DtsJobId**.
	//
	// example:
	//
	// dtsaw012y2g15q****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the DTS task. You can call [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) to query the task ID.
	//
	// > If you do not specify this parameter, you must specify **DtsInstanceId**.
	//
	// example:
	//
	// m4312mab158****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The database instance to be modified. Valid values:
	//
	// - **src**: source instance.
	//
	// - **dest**: destination instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// src
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The ID of the database instance.
	//
	// example:
	//
	// rm-bp10k50h8374w****
	EndpointInstanceId *string `json:"EndpointInstanceId,omitempty" xml:"EndpointInstanceId,omitempty"`
	// The type of the database instance. Valid values:
	//
	// - **rds**: ApsaraDB RDS for MySQL or ApsaraDB RDS for PostgreSQL.
	//
	// - **polardb**: PolarDB for MySQL or PolarDB for PostgreSQL.
	//
	// - **mongodb**: when used as the source, ApsaraDB for MongoDB (replica set architecture). When used as the destination, ApsaraDB for MongoDB (replica set or sharded cluster architecture).
	//
	// - **distributed_mongodb**: supported only as the source of a distributed instance. Indicates ApsaraDB for MongoDB (sharded cluster architecture).
	//
	// > The incremental node of a distributed instance must obtain data changes from the source through Oplog.
	//
	// - **greenplum**: cloud-native data warehouse AnalyticDB for PostgreSQL.
	//
	// - **kafka**: ApsaraMQ for Kafka.
	//
	// - **ecs**: self-managed database on an ECS instance (only supported database types).
	//
	// - **express**: database connected over Express Connect (only supported database types).
	//
	// - **other**: database connected over the Internet (only supported database types).
	//
	// > - Currently supported database types include **MySQL**, **PolarDB for MySQL**, **PostgreSQL**, **PolarDB for PostgreSQL**, **MongoDB**, **Kafka**, and **AnalyticDB PostgreSQL**.
	//
	// - If the database is MongoDB (sharded cluster), the number of shards in the new database must be the same as that in the original MongoDB (sharded cluster).
	//
	// - If the source instance is to be modified and the database type is **PostgreSQL**, make sure that the latency of the DTS instance is less than 30 seconds and stop writing data to the source. Otherwise, inconsistent data may occur.
	//
	// - The parameter values are case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	EndpointInstanceType *string `json:"EndpointInstanceType,omitempty" xml:"EndpointInstanceType,omitempty"`
	// The IP address of the database instance.
	//
	// example:
	//
	// 172.168.XX.XXX
	EndpointIp *string `json:"EndpointIp,omitempty" xml:"EndpointIp,omitempty"`
	// The port of the database instance.
	//
	// example:
	//
	// 3306
	EndpointPort *string `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	// The primary vSwitch for Express Connect access.
	//
	// example:
	//
	// vsw-bp1w7gscw7pky*******
	EndpointPrimaryVswId *string `json:"EndpointPrimaryVswId,omitempty" xml:"EndpointPrimaryVswId,omitempty"`
	// The region to which the database instance belongs.
	//
	// example:
	//
	// cn-hangzhou
	EndpointRegionId *string `json:"EndpointRegionId,omitempty" xml:"EndpointRegionId,omitempty"`
	// The secondary vSwitch for Express Connect access.
	//
	// example:
	//
	// vsw-bp1ud8e2mhw*****
	EndpointSecondaryVswId *string `json:"EndpointSecondaryVswId,omitempty" xml:"EndpointSecondaryVswId,omitempty"`
	// The VPC ID for Express Connect access.
	//
	// example:
	//
	// vpc-bp1q00qitocaem****
	EndpointVpcId *string `json:"EndpointVpcId,omitempty" xml:"EndpointVpcId,omitempty"`
	// Specifies whether to modify the account and password. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false*	- (default): No.
	//
	// example:
	//
	// false
	ModifyAccount *bool `json:"ModifyAccount,omitempty" xml:"ModifyAccount,omitempty"`
	// The database password.
	//
	// > This parameter takes effect only when **ModifyAccount*	- is set to **true**.
	//
	// example:
	//
	// DTStest****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region to which the DTS instance belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the RAM role for cross-account access.
	//
	// > Specify this parameter when performing cross-account data synchronization. For the required permissions and authorization method of this role, see [Configure RAM authorization for cross-account data migration or synchronization](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The password of the shard in the MongoDB sharded cluster instance.
	//
	// > - This parameter is available and required only when the source database instance is ApsaraDB for MongoDB (sharded cluster architecture).
	//
	// - This parameter takes effect only when **ModifyAccount*	- is set to **true**.
	//
	// example:
	//
	// DTStest****
	ShardPassword *string `json:"ShardPassword,omitempty" xml:"ShardPassword,omitempty"`
	// The account of the shard in the MongoDB sharded cluster instance.
	//
	// > - This parameter is available and required only when the source database instance is ApsaraDB for MongoDB (sharded cluster architecture).
	//
	// - This parameter takes effect only when **ModifyAccount*	- is set to **true**.
	//
	// example:
	//
	// shard
	ShardUsername *string `json:"ShardUsername,omitempty" xml:"ShardUsername,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward*	- (default): forward.
	//
	// - **Reverse**: reverse.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The database account.
	//
	// > This parameter takes effect only when **ModifyAccount*	- is set to **true**.
	//
	// example:
	//
	// dtstest
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// Specifies whether this is a seamless integration (zero-ETL) node. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
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

func (s *ModifyDtsJobEndpointRequest) GetEndpointPrimaryVswId() *string {
	return s.EndpointPrimaryVswId
}

func (s *ModifyDtsJobEndpointRequest) GetEndpointRegionId() *string {
	return s.EndpointRegionId
}

func (s *ModifyDtsJobEndpointRequest) GetEndpointSecondaryVswId() *string {
	return s.EndpointSecondaryVswId
}

func (s *ModifyDtsJobEndpointRequest) GetEndpointVpcId() *string {
	return s.EndpointVpcId
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

func (s *ModifyDtsJobEndpointRequest) SetEndpointPrimaryVswId(v string) *ModifyDtsJobEndpointRequest {
	s.EndpointPrimaryVswId = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetEndpointRegionId(v string) *ModifyDtsJobEndpointRequest {
	s.EndpointRegionId = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetEndpointSecondaryVswId(v string) *ModifyDtsJobEndpointRequest {
	s.EndpointSecondaryVswId = &v
	return s
}

func (s *ModifyDtsJobEndpointRequest) SetEndpointVpcId(v string) *ModifyDtsJobEndpointRequest {
	s.EndpointVpcId = &v
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
