// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureMigrationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationEndpoint(v *ConfigureMigrationJobRequestDestinationEndpoint) *ConfigureMigrationJobRequest
	GetDestinationEndpoint() *ConfigureMigrationJobRequestDestinationEndpoint
	SetMigrationMode(v *ConfigureMigrationJobRequestMigrationMode) *ConfigureMigrationJobRequest
	GetMigrationMode() *ConfigureMigrationJobRequestMigrationMode
	SetSourceEndpoint(v *ConfigureMigrationJobRequestSourceEndpoint) *ConfigureMigrationJobRequest
	GetSourceEndpoint() *ConfigureMigrationJobRequestSourceEndpoint
	SetAccountId(v string) *ConfigureMigrationJobRequest
	GetAccountId() *string
	SetCheckpoint(v string) *ConfigureMigrationJobRequest
	GetCheckpoint() *string
	SetMigrationJobId(v string) *ConfigureMigrationJobRequest
	GetMigrationJobId() *string
	SetMigrationJobName(v string) *ConfigureMigrationJobRequest
	GetMigrationJobName() *string
	SetMigrationObject(v string) *ConfigureMigrationJobRequest
	GetMigrationObject() *string
	SetMigrationReserved(v string) *ConfigureMigrationJobRequest
	GetMigrationReserved() *string
	SetOwnerId(v string) *ConfigureMigrationJobRequest
	GetOwnerId() *string
	SetRegionId(v string) *ConfigureMigrationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ConfigureMigrationJobRequest
	GetResourceGroupId() *string
}

type ConfigureMigrationJobRequest struct {
	DestinationEndpoint *ConfigureMigrationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationMode       *ConfigureMigrationJobRequestMigrationMode       `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	SourceEndpoint      *ConfigureMigrationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The Alibaba Cloud account ID. You do not need to specify this parameter because it will be deprecated.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The start position of incremental data migration. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 111
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The ID of the data migration instance. You can call the **DescribeMigrationJobs*	- operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsxxxxxxxx
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	// The name of the migration task. The name can be up to 32 characters in length. Specify a descriptive name for easy identification. Uniqueness is not required.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL_TO_RDS
	MigrationJobName *string `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	// The objects to be migrated. The value is a JSON string that supports regular expressions. For more information, see [Migration object configuration](~141901~).
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"DBName":"dtstestdata","TableIncludes":[{"TableName":"customer"}]}]
	MigrationObject *string `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet special requirements, such as whether to automatically start the precheck. For more information, see [MigrationReserved parameter description](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {     "autoStartModulesAfterConfig": "none",     "targetTableMode": 2 }
	MigrationReserved *string `json:"MigrationReserved,omitempty" xml:"MigrationReserved,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the data migration instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > The region ID must be the same as the region ID of the destination database.
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
}

func (s ConfigureMigrationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobRequest) GetDestinationEndpoint() *ConfigureMigrationJobRequestDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *ConfigureMigrationJobRequest) GetMigrationMode() *ConfigureMigrationJobRequestMigrationMode {
	return s.MigrationMode
}

func (s *ConfigureMigrationJobRequest) GetSourceEndpoint() *ConfigureMigrationJobRequestSourceEndpoint {
	return s.SourceEndpoint
}

func (s *ConfigureMigrationJobRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ConfigureMigrationJobRequest) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *ConfigureMigrationJobRequest) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *ConfigureMigrationJobRequest) GetMigrationJobName() *string {
	return s.MigrationJobName
}

func (s *ConfigureMigrationJobRequest) GetMigrationObject() *string {
	return s.MigrationObject
}

func (s *ConfigureMigrationJobRequest) GetMigrationReserved() *string {
	return s.MigrationReserved
}

func (s *ConfigureMigrationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ConfigureMigrationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureMigrationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigureMigrationJobRequest) SetDestinationEndpoint(v *ConfigureMigrationJobRequestDestinationEndpoint) *ConfigureMigrationJobRequest {
	s.DestinationEndpoint = v
	return s
}

func (s *ConfigureMigrationJobRequest) SetMigrationMode(v *ConfigureMigrationJobRequestMigrationMode) *ConfigureMigrationJobRequest {
	s.MigrationMode = v
	return s
}

func (s *ConfigureMigrationJobRequest) SetSourceEndpoint(v *ConfigureMigrationJobRequestSourceEndpoint) *ConfigureMigrationJobRequest {
	s.SourceEndpoint = v
	return s
}

func (s *ConfigureMigrationJobRequest) SetAccountId(v string) *ConfigureMigrationJobRequest {
	s.AccountId = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetCheckpoint(v string) *ConfigureMigrationJobRequest {
	s.Checkpoint = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetMigrationJobId(v string) *ConfigureMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetMigrationJobName(v string) *ConfigureMigrationJobRequest {
	s.MigrationJobName = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetMigrationObject(v string) *ConfigureMigrationJobRequest {
	s.MigrationObject = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetMigrationReserved(v string) *ConfigureMigrationJobRequest {
	s.MigrationReserved = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetOwnerId(v string) *ConfigureMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetRegionId(v string) *ConfigureMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetResourceGroupId(v string) *ConfigureMigrationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigureMigrationJobRequest) Validate() error {
	if s.DestinationEndpoint != nil {
		if err := s.DestinationEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.MigrationMode != nil {
		if err := s.MigrationMode.Validate(); err != nil {
			return err
		}
	}
	if s.SourceEndpoint != nil {
		if err := s.SourceEndpoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigureMigrationJobRequestDestinationEndpoint struct {
	// 待迁入的数据库名称或鉴权数据库名称。
	//
	// > - 当**DestinationEndpoint.EngineName**取值为**PostgreSQL**、**DRDS**或**MongoDB**时，本参数才可用且必须传入。
	//
	// - 当**DestinationEndpoint.EngineName**取值为**PostgreSQL**或**DRDS**时，传入待迁移的数据库名称；取值为**MongoDB**时，传入数据库账号的鉴权数据库名称。
	//
	// example:
	//
	// dtstestdatabase
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
	// 目标库的数据库类型。取值：**MySQL**、**DRDS**、**SQLServer**、**PostgreSQL**、**PPAS**、**MongoDB**、**Redis**、**POLARDB**、**polardb_pg**
	//
	// > 当**DestinationEndpoint.InstanceType**取值为**RDS**、**POLARDB**、**ECS**、**LocalInstance**或**Express**时，本参数才可用且必须传入。
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// 目标库的连接地址。
	//
	// > 当**DestinationEndpoint.InstanceType**取值为**LocalInstance**或**Express**时，本参数才可用且必须传入。
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// 目标实例ID。
	//
	// > 当**DestinationEndpoint.InstanceType**取值为**RDS**、**ECS**、**MongoDB**、**Redis**、**DRDS**、**PetaData**、**OceanBase**、**POLARDB**、**PolarDB_o**、**AnalyticDB**或**Greenplum**时，本参数才可用且必须传入对应的实例ID（例如取值为**ECS**，则需要传入ECS实例ID）。
	//
	// example:
	//
	// rm-bp1r46452ai50****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// 目标库的实例类型，取值：
	//
	// - **ECS**：ECS上的自建数据库。
	//
	// - **LocalInstance**：有公网IP的自建数据库。
	//
	// - **RDS**：阿里云RDS实例。
	//
	// - **DRDS**：阿里云PolarDB-X实例。
	//
	// - **MongoDB**：阿里云MongoDB实例。
	//
	// - **Redis**：阿里云Redis实例。
	//
	// - **PetaData**：阿里云HybridDB for MySQL实例。
	//
	// - **POLARDB**：阿里云PolarDB MySQL、PolarDB PostgreSQL。
	//
	// - **PolarDB_o**：阿里云PolarDB O引擎集群。
	//
	// - **AnalyticDB**：阿里云云原生数据仓库AnalyticDB MySQL 3.0和2.0版本。
	//
	// - **Greenplum**：阿里云云原生数据仓库AnalyticDB PostgreSQL。
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Oracle数据库的SID信息。
	//
	// > 当**DestinationEndpoint.EngineName**取值为**Oracle**，且**Oracle**数据库为非RAC实例时，本参数才可用且必须传入。
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// 目标库数据库账号的密码。
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 目标库的服务端口。
	//
	// > 当**DestinationEndpoint.InstanceType**取值为**ECS**、**LocalInstance**或**Express**时，本参数才可用且必须传入。
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 目标库所属的地域ID。
	//
	// > 当**DestinationEndpoint.InstanceType**取值为**LocalInstance**时，您可以传入**cn-hangzhou**或者离自建数据库地物理距离最近的地域ID，详情请参见[支持的地域列表](https://help.aliyun.com/document_detail/141033.html)。
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 目标库的数据库账号。
	//
	// 说明 迁移不同的数据库所需的权限有所差异，详情请参见迁移方案概览中对应的配置案例。
	//
	// example:
	//
	// dtstestaccount
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ConfigureMigrationJobRequestDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ConfigureMigrationJobRequestDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) GetDataBaseName() *string {
	return s.DataBaseName
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) GetIP() *string {
	return s.IP
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) GetPassword() *string {
	return s.Password
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) GetRegion() *string {
	return s.Region
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetDataBaseName(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.DataBaseName = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetEngineName(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetIP(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetInstanceID(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetInstanceType(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetOracleSID(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetPassword(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetPort(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetRegion(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetUserName(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type ConfigureMigrationJobRequestMigrationMode struct {
	// 是否进行全量数据迁移，取值：
	//
	// - **true**：是。
	//
	// - **false**：否。
	//
	// > DTS对全量数据迁移的支持情况因数据库类型不同而有所差异，详情请参见[支持的数据库和迁移类型](https://help.aliyun.com/document_detail/26618.html)。
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataIntialization *bool `json:"DataIntialization,omitempty" xml:"DataIntialization,omitempty"`
	// 是否进行增量数据迁移，取值：
	//
	// - **true**：是。
	//
	// - **false**：否。
	//
	// > DTS对增量数据迁移的支持情况因数据库类型不同而有所差异，详情请参见[支持的数据库和迁移类型](https://help.aliyun.com/document_detail/26618.html)。
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// 是否进行结构迁移，取值：
	//
	// - **true**：是。
	//
	// - **false**：否。
	//
	// > DTS对结构迁移的支持情况因数据库类型不同而有所差异，详情请参见[支持的数据库和迁移类型](https://help.aliyun.com/document_detail/26618.html)。
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	StructureIntialization *bool `json:"StructureIntialization,omitempty" xml:"StructureIntialization,omitempty"`
}

func (s ConfigureMigrationJobRequestMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s ConfigureMigrationJobRequestMigrationMode) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobRequestMigrationMode) GetDataIntialization() *bool {
	return s.DataIntialization
}

func (s *ConfigureMigrationJobRequestMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *ConfigureMigrationJobRequestMigrationMode) GetStructureIntialization() *bool {
	return s.StructureIntialization
}

func (s *ConfigureMigrationJobRequestMigrationMode) SetDataIntialization(v bool) *ConfigureMigrationJobRequestMigrationMode {
	s.DataIntialization = &v
	return s
}

func (s *ConfigureMigrationJobRequestMigrationMode) SetDataSynchronization(v bool) *ConfigureMigrationJobRequestMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *ConfigureMigrationJobRequestMigrationMode) SetStructureIntialization(v bool) *ConfigureMigrationJobRequestMigrationMode {
	s.StructureIntialization = &v
	return s
}

func (s *ConfigureMigrationJobRequestMigrationMode) Validate() error {
	return dara.Validate(s)
}

type ConfigureMigrationJobRequestSourceEndpoint struct {
	// 待迁移的数据库名称或鉴权数据库名称。
	//
	// > - 当**SourceEndpoint.EngineName**取值为**PostgreSQL**或**MongoDB**时，本参数才可用且必须传入。
	//
	// - 当**SourceEndpoint.EngineName**取值为**PostgreSQL**时，传入待迁移的数据库名称；取值为**MongoDB**时，传入数据库账号的鉴权数据库名称。
	//
	// example:
	//
	// dtstestdatabase
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// 源库的数据库类型，取值：**MySQL**、**TiDB**、**SQLServer**、**PostgreSQL**、**Oracle**、**MongoDB**、**Redis**、**POLARDB**、**polardb_pg**。
	//
	// > 当**DestinationEndpoint.InstanceType**取值为**RDS**、**POLARDB**、**ECS**、**LocalInstance**或**Express**时，本参数才可用且必须传入。
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// 源库的连接地址。
	//
	// > 当**SourceEndpoint.InstanceType**取值为**LocalInstance**或**Express**时，本参数才可用且必须传入。
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// 源库的实例ID。
	//
	// > - 当**SourceEndpoint.InstanceType**取值为**RDS**、**ECS**、**Express**、**MongoDB**、**POLARDB**或**PolarDB_o**时，本参数才可用且必须传入对应的实例ID（例如取值为**ECS**，则本参数传入ECS实例的ID）。
	//
	// - 当**SourceEndpoint.InstanceType**取值为**Express**时，本参数传入VPC ID（即专有网络ID）。
	//
	// example:
	//
	// bp-rmxxxxxxxx
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// 源库的实例类型，取值：
	//
	// - **RDS**：阿里云RDS实例。
	//
	// - **ECS**：ECS上的自建数据库。
	//
	// - **LocalInstance**：有公网IP的自建数据库。
	//
	// - **Express**：通过专线/VPN网关/智能接入网关接入的自建数据库。
	//
	// - **dg**：通过数据库网关DG接入的自建数据库。
	//
	// - **cen**：通过云企业网CEN接入的自建数据库。
	//
	// - **MongoDB**：阿里云MongoDB实例。
	//
	// - **POLARDB**：阿里云PolarDB MySQL、PolarDB PostgreSQL。
	//
	// - **PolarDB_o**：阿里云PolarDB O引擎集群。
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Oracle数据库的SID信息。
	//
	// > 当**SourceEndpoint.EngineName**取值为**Oracle**，且Oracle数据库为非RAC实例时，本参数才可用且必须传入。
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// 源实例所属的阿里云账号ID。
	//
	// > 仅在配置跨阿里云账号的数据迁移时本参数才可用，且必须传入。
	//
	// example:
	//
	// 140692647406****
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// 源库数据库账号对应的密码。
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 源库的服务端口。
	//
	// > 当**SourceEndpoint.InstanceType**取值为**ECS**、**LocalInstance**或**Express**时，本参数才可用且必须传入。
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 源库所属的地域ID。
	//
	// > 当**SourceEndpoint.InstanceType**取值为**LocalInstance**时，您可以传入**cn-hangzhou**或者离自建数据库地物理距离最近的地域ID，详情请参见[支持的地域列表](https://help.aliyun.com/document_detail/141033.html)。
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 当源实例与目标实例所属阿里云账号不同时，需传入该参数，来指定源实例的授权角色，以允许目标实例阿里云账号访问源实例的实例信息。
	//
	// > 角色所需的权限及授权方式，请参见[跨阿里云账号数据迁移或同步时如何配置RAM授权](https://help.aliyun.com/document_detail/48468.html)。
	//
	// example:
	//
	// ram-for-dts
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 源库的数据库账号。
	//
	// 说明 迁移不同的数据库所需的权限有所差异，详情请参见迁移方案概览中对应的配置案例。
	//
	// example:
	//
	// dtstestaccount
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ConfigureMigrationJobRequestSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ConfigureMigrationJobRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) GetIP() *string {
	return s.IP
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) GetOwnerID() *string {
	return s.OwnerID
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) GetPassword() *string {
	return s.Password
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) GetRegion() *string {
	return s.Region
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) GetRole() *string {
	return s.Role
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetDatabaseName(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetEngineName(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetIP(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetInstanceID(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetInstanceType(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetOracleSID(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetOwnerID(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.OwnerID = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetPassword(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetPort(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetRegion(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.Region = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetRole(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.Role = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetUserName(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) Validate() error {
	return dara.Validate(s)
}
