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
	// The objects that you want to migrate. The value is a JSON string and can contain regular expressions.
	//
	// For more information, see [MigrationObject](https://help.aliyun.com/document_detail/141227.html).
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Specifies whether to perform incremental data migration. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// >  For more information about databases that support incremental data migration, see [Supported databases and migration types](https://help.aliyun.com/document_detail/26618.html).
	//
	// example:
	//
	// 111
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// system
	//
	// This parameter is required.
	//
	// example:
	//
	// The operation that you want to perform. Set the value to **ConfigureMigrationJob**.
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	// The ID of the region where the data migration instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// >  The region ID of the data migration instance is the same as that of the destination database.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL_TO_RDS
	MigrationJobName *string `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	// Specifies whether to perform schema migration. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// >  For more information about databases that support schema migration, see [Supported databases and migration types](https://help.aliyun.com/document_detail/26618.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"DBName":"dtstestdata","TableIncludes":[{"TableName":"customer"}]}]
	MigrationObject *string `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty"`
	// Specifies whether to perform full data migration. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// >  For more information about databases that support full data migration, see [Supported databases and migration types](https://help.aliyun.com/document_detail/26618.html).
	//
	// example:
	//
	// {     "autoStartModulesAfterConfig": "none",     "targetTableMode": 2 }
	MigrationReserved *string `json:"MigrationReserved,omitempty" xml:"MigrationReserved,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource GroupId
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
	// The region ID of the destination database.
	//
	// >  If the **DestinationEndpoint.InstanceType*	- parameter is set to **LocalInstance**, you can enter **cn-hangzhou*	- or the ID of the region closest to the self-managed database. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// dtstestdatabase
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
	// The authorized RAM role of the source instance. You must specify the RAM role only if the source instance and the destination instance belong to different Alibaba Cloud accounts. You can use the RAM role to allow the Alibaba Cloud account that owns the destination instance to access the source instance.
	//
	// >  For information about the permissions and authorization methods of the RAM role, see [Configure RAM authorization for cross-account data migration and synchronization](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The ID of the instance that hosts the destination database.
	//
	// >  You must specify the instance ID only if the **DestinationEndpoint.InstanceType*	- parameter is set to **RDS**, **ECS**, **MongoDB**, **Redis**, **DRDS**, **PetaData**, **OceanBase**, **POLARDB**, **PolarDB_o**, **AnalyticDB**, or **Greenplum**. For example, if the DestinationEndpoint.InstanceType parameter is set to **ECS**, you must specify the ID of the ECS instance.
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The ID of the Alibaba Cloud account to which the source instance belongs.
	//
	// >  You must specify this parameter only when you configure data migration across different Alibaba Cloud accounts.
	//
	// example:
	//
	// rm-bp1r46452ai50****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The password of the source database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet special requirements, for example, whether to automatically start a precheck. For more information, see [MigrationReserved](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The service port number of the destination database.
	//
	// >  You must specify the service port number only if the **DestinationEndpoint.InstanceType*	- parameter is set to **ECS**, **LocalInstance**, or **Express**.
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The engine type of the destination database. Valid values: **MySQL**, **DRDS**, **SQLServer**, **PostgreSQL**, **PPAS**, **MongoDB**, **Redis**, **POLARDB**, and **polardb_pg**.
	//
	// >  You must specify the engine type only if the **DestinationEndpoint.InstanceType*	- parameter is set to **RDS**, **POLARDB**, **ECS**, **LocalInstance**, or **Express**.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The instance type of the destination database. Valid values:
	//
	// 	- **ECS**: self-managed database that is hosted on Elastic Compute Service (ECS)
	//
	// 	- **LocalInstance**: self-managed database with a public IP address
	//
	// 	- **RDS**: ApsaraDB RDS instance
	//
	// 	- **DRDS**: PolarDB-X instance
	//
	// 	- **MongoDB**: ApsaraDB for MongoDB instance
	//
	// 	- **Redis**: ApsaraDB for Redis instance
	//
	// 	- **PetaData**: HybridDB for MySQL instance
	//
	// 	- **POLARDB**: PolarDB for MySQL cluster or PolarDB for PostgreSQL cluster
	//
	// 	- **PolarDB_o**: PolarDB O Edition cluster
	//
	// 	- **AnalyticDB**: AnalyticDB for MySQL cluster V3.0 or V2.0
	//
	// 	- **Greenplum**: AnalyticDB for PostgreSQL instance
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The endpoint of the destination database.
	//
	// >  You must specify the endpoint only if the **DestinationEndpoint.InstanceType*	- parameter is set to **LocalInstance*	- or **Express**.
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
	// The database account of the destination database.
	//
	// >  The permissions that are required for database accounts vary with the migration scenario. For more information, see [Overview of data migration scenarios](https://help.aliyun.com/document_detail/26618.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataIntialization *bool `json:"DataIntialization,omitempty" xml:"DataIntialization,omitempty"`
	// The password of the destination database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// The name of the destination database or the authentication database.
	//
	// >
	//
	// 	- You must specify the database name only if the **DestinationEndpoint.EngineName*	- parameter is set to **PostgreSQL**, **DRDS**, or **MongoDB**.
	//
	// 	- If the **DestinationEndpoint.EngineName*	- parameter is set to **PostgreSQL*	- or **DRDS**, specify the name of the destination database. If the DestinationEndpoint.EngineName parameter is set to **MongoDB**, specify the name of the authentication database.
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
	// The endpoint of the source database.
	//
	// >  You must specify the endpoint only if the **SourceEndpoint.InstanceType*	- parameter is set to **LocalInstance*	- or **Express**.
	//
	// example:
	//
	// dtstestdatabase
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The instance type of the source database. Valid values:
	//
	// 	- **RDS**: ApsaraDB RDS instance
	//
	// 	- **ECS**: self-managed database that is hosted on ECS
	//
	// 	- **LocalInstance**: self-managed database with a public IP address
	//
	// 	- **Express**: self-managed database that is connected over Express Connect, VPN Gateway, or Smart Access Gateway
	//
	// 	- **dg**: self-managed database that is connected over Database Gateway
	//
	// 	- **cen**: self-managed database that is connected over Cloud Enterprise Network (CEN)
	//
	// 	- **MongoDB**: ApsaraDB for MongoDB instance
	//
	// 	- **POLARDB**: PolarDB for MySQL cluster or PolarDB for PostgreSQL cluster
	//
	// 	- **PolarDB_o**: PolarDB O Edition cluster
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// rm-bp1i99e8l7913****
	//
	// example:
	//
	// 172.16.88.***
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// dtsl3m1213ye7l****
	//
	// example:
	//
	// The name of the data migration task. The name can be up to 32 characters in length. We recommend that you specify an informative name to identify the task. You do not need to use a unique task name.
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The ID of the data migration instance. You can call the **DescribeMigrationJobs*	- operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The region ID of the source database.
	//
	// >  If the **SourceEndpoint.InstanceType*	- parameter is set to **LocalInstance**, you can enter **cn-hangzhou*	- or the ID of the region closest to the self-managed database. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The name of the source database or the authentication database.
	//
	// >
	//
	// 	- You must specify the database name only if the **SourceEndpoint.EngineName*	- parameter is set to **PostgreSQL*	- or **MongoDB**.
	//
	// 	- If the **SourceEndpoint.EngineName*	- parameter is set to **PostgreSQL**, specify the name of the source database. If the SourceEndpoint.EngineName parameter is set to **MongoDB**, specify the name of the authentication database.
	//
	// example:
	//
	// 140692647406****
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The system ID (SID) of the Oracle database.
	//
	// >  You must specify this parameter only if the **SourceEndpoint.EngineName*	- parameter is set to **Oracle*	- and the **Oracle*	- database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The engine type of the source database. Valid values: **MySQL**, **TiDB**, **SQLServer**, **PostgreSQL**, **Oracle**, **MongoDB**, **Redis**, **POLARDB**, and **polardb_pg**.
	//
	// >  You must specify the engine type only if the **DestinationEndpoint.InstanceType*	- parameter is set to **RDS**, **POLARDB**, **ECS**, **LocalInstance**, or **Express**.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the instance that hosts the source database.
	//
	// >
	//
	// 	- You must specify the instance ID only if the **SourceEndpoint.InstanceType*	- parameter is set to **RDS**, **ECS**, **Express**, **MongoDB**, **POLARDB**, or **PolarDB_o**. For example, if the SourceEndpoint.InstanceType parameter is set to **ECS**, you must specify the ID of the ECS instance.
	//
	// 	- If the **SourceEndpoint.InstanceType*	- parameter is set to **Express**, you must specify the ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The database account of the source database.
	//
	// >  The permissions that are required for database accounts vary with the migration scenario. For more information, see [Overview of data migration scenarios](https://help.aliyun.com/document_detail/26618.html).
	//
	// example:
	//
	// ram-for-dts
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The service port number of the source database.
	//
	// >  You must specify the service port number only if the **SourceEndpoint.InstanceType*	- parameter is set to **ECS**, **LocalInstance**, or **Express**.
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
