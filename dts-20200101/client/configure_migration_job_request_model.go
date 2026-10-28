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
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter. This parameter is pending deprecation.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The start checkpoint for incremental data migration. The value is a UNIX timestamp in seconds.
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
	// The name of the migration task. The name can be up to 32 characters in length. Specify a descriptive name for easy identification. The name does not need to be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL_TO_RDS
	MigrationJobName *string `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	// The objects to be migrated. The value is a JSON string and supports specific regular expressions. For more information, see [Configure migration objects](~141901~).
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"DBName":"dtstestdata","TableIncludes":[{"TableName":"customer"}]}]
	MigrationObject *string `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty"`
	// The reserved parameters of DTS. The value is a JSON string. You can specify this parameter to meet specific requirements, such as specifying whether to automatically start a precheck. For more information, see [MigrationReserved parameter description](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {     "autoStartModulesAfterConfig": "none",     "targetTableMode": 2 }
	MigrationReserved *string `json:"MigrationReserved,omitempty" xml:"MigrationReserved,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the data migration instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > The region ID must be the same as the region ID of the destination database.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
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
	// The name of the destination database or the authentication database.
	//
	// > - This parameter is available and required only when **DestinationEndpoint.EngineName*	- is set to **PostgreSQL**, **DRDS**, or **MongoDB**.
	//
	// - When **DestinationEndpoint.EngineName*	- is set to **PostgreSQL*	- or **DRDS**, specify the name of the database to be migrated. When it is set to **MongoDB**, specify the name of the authentication database for the database account.
	//
	// example:
	//
	// dtstestdatabase
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
	// The database engine of the destination database. Valid values: **MySQL**, **DRDS**, **SQLServer**, **PostgreSQL**, **PPAS**, **MongoDB**, **Redis**, **POLARDB**, and **polardb_pg**.
	//
	// > This parameter is available and required only when **DestinationEndpoint.InstanceType*	- is set to **RDS**, **POLARDB**, **ECS**, **LocalInstance**, or **Express**.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The endpoint of the destination database.
	//
	// > This parameter is available and required only when **DestinationEndpoint.InstanceType*	- is set to **LocalInstance*	- or **Express**.
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the destination instance.
	//
	// > This parameter is available and required only when **DestinationEndpoint.InstanceType*	- is set to **RDS**, **ECS**, **MongoDB**, **Redis**, **DRDS**, **PetaData**, **OceanBase**, **POLARDB**, **PolarDB_o**, **AnalyticDB**, or **Greenplum**. You must specify the corresponding instance ID. For example, if the value is **ECS**, specify the ID of the ECS instance.
	//
	// example:
	//
	// rm-bp1r46452ai50****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The instance type of the destination database. Valid values:
	//
	// - **ECS**: Self-managed database hosted on ECS.
	//
	// - **LocalInstance**: Self-managed database with a public IP address.
	//
	// - **RDS**: ApsaraDB RDS instance.
	//
	// - **DRDS**: PolarDB-X instance.
	//
	// - **MongoDB**: ApsaraDB for MongoDB instance.
	//
	// - **Redis**: ApsaraDB for Redis instance.
	//
	// - **PetaData**: HybridDB for MySQL instance.
	//
	// - **POLARDB**: PolarDB for MySQL or PolarDB for PostgreSQL cluster.
	//
	// - **PolarDB_o**: PolarDB for Oracle cluster.
	//
	// - **AnalyticDB**: AnalyticDB for MySQL 3.0 or 2.0 cluster.
	//
	// - **Greenplum**: AnalyticDB for PostgreSQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is available and required only when **DestinationEndpoint.EngineName*	- is set to **Oracle*	- and the Oracle database is a non-Real Application Cluster instance.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The password of the destination database account.
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The service port number of the destination database.
	//
	// > This parameter is available and required only when **DestinationEndpoint.InstanceType*	- is set to **ECS**, **LocalInstance**, or **Express**.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region where the destination database resides.
	//
	// > When **DestinationEndpoint.InstanceType*	- is set to **LocalInstance**, you can specify **cn-hangzhou*	- or the ID of the region that is closest to your self-managed database. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The database account of the destination database.
	//
	// Note: The required permissions vary based on the database engine. For more information, see the corresponding configuration examples in the migration overview.
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
	// Specifies whether to perform full data migration. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// > The support for full data migration varies based on the database engine. For more information, see [Supported databases and migration types](https://help.aliyun.com/document_detail/26618.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataIntialization *bool `json:"DataIntialization,omitempty" xml:"DataIntialization,omitempty"`
	// Specifies whether to perform incremental data migration. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// > The support for incremental data migration varies based on the database engine. For more information, see [Supported databases and migration types](https://help.aliyun.com/document_detail/26618.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Specifies whether to perform schema migration. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// > The support for schema migration varies based on the database engine. For more information, see [Supported databases and migration types](https://help.aliyun.com/document_detail/26618.html).
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
	// The name of the database to be migrated or the authentication database.
	//
	// > - This parameter is available and required only when **SourceEndpoint.EngineName*	- is set to **PostgreSQL*	- or **MongoDB**.
	//
	// - When **SourceEndpoint.EngineName*	- is set to **PostgreSQL**, specify the name of the database to be migrated. When it is set to **MongoDB**, specify the name of the authentication database for the database account.
	//
	// example:
	//
	// dtstestdatabase
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine of the source database. Valid values: **MySQL**, **TiDB**, **SQLServer**, **PostgreSQL**, **Oracle**, **MongoDB**, **Redis**, **POLARDB**, and **polardb_pg**.
	//
	// > This parameter is available and required only when **DestinationEndpoint.InstanceType*	- is set to **RDS**, **POLARDB**, **ECS**, **LocalInstance**, or **Express**.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The endpoint of the source database.
	//
	// > This parameter is available and required only when **SourceEndpoint.InstanceType*	- is set to **LocalInstance*	- or **Express**.
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the source instance.
	//
	// > - This parameter is available and required only when **SourceEndpoint.InstanceType*	- is set to **RDS**, **ECS**, **Express**, **MongoDB**, **POLARDB**, or **PolarDB_o**. You must specify the corresponding instance ID. For example, if the value is **ECS**, specify the ID of the ECS instance.
	//
	// - When **SourceEndpoint.InstanceType*	- is set to **Express**, specify the VPC ID.
	//
	// example:
	//
	// bp-rmxxxxxxxx
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The instance type of the source database. Valid values:
	//
	// - **RDS**: ApsaraDB RDS instance.
	//
	// - **ECS**: Self-managed database hosted on Elastic Compute Service (ECS).
	//
	// - **LocalInstance**: Self-managed database with a public IP address.
	//
	// - **Express**: Self-managed database connected over Express Connect, VPN Gateway, or Smart Access Gateway.
	//
	// - **dg**: Self-managed database connected over Database Gateway.
	//
	// - **cen**: Self-managed database connected over Cloud Enterprise Network (CEN).
	//
	// - **MongoDB**: ApsaraDB for MongoDB instance.
	//
	// - **POLARDB**: PolarDB for MySQL or PolarDB for PostgreSQL cluster.
	//
	// - **PolarDB_o**: PolarDB for Oracle cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The system ID (SID) of the Oracle database.
	//
	// > This parameter is available and required only when **SourceEndpoint.EngineName*	- is set to **Oracle*	- and the Oracle database is a non-Real Application Cluster instance.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The Alibaba Cloud account ID to which the source instance belongs.
	//
	// > This parameter is active and required only when you configure a cross-account data migration task.
	//
	// example:
	//
	// 140692647406****
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The password of the source database account.
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The service port number of the source database.
	//
	// > This parameter is available and required only when **SourceEndpoint.InstanceType*	- is set to **ECS**, **LocalInstance**, or **Express**.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region where the source database resides.
	//
	// > When **SourceEndpoint.InstanceType*	- is set to **LocalInstance**, you can specify **cn-hangzhou*	- or the ID of the region that is closest to your self-managed database. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The RAM role that is authorized to the destination Alibaba Cloud account to access the source instance. This parameter is required when the source and destination instances belong to different Alibaba Cloud accounts.
	//
	// > For more information about the required permissions and authorization method, see [Configure RAM authorization for cross-account data migration or synchronization](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The database account of the source database.
	//
	// Note: The required permissions vary based on the database engine. For more information, see the corresponding configuration examples in the migration overview.
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
