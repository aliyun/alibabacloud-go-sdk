// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestoreTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoOpenDatabase(v string) *CreateRestoreTaskRequest
	GetAutoOpenDatabase() *string
	SetAutoShutdownDatabase(v string) *CreateRestoreTaskRequest
	GetAutoShutdownDatabase() *string
	SetBackupGatewayId(v int64) *CreateRestoreTaskRequest
	GetBackupGatewayId() *int64
	SetBackupPlanId(v string) *CreateRestoreTaskRequest
	GetBackupPlanId() *string
	SetBackupSetId(v string) *CreateRestoreTaskRequest
	GetBackupSetId() *string
	SetClientToken(v string) *CreateRestoreTaskRequest
	GetClientToken() *string
	SetCrossAliyunId(v string) *CreateRestoreTaskRequest
	GetCrossAliyunId() *string
	SetCrossRoleName(v string) *CreateRestoreTaskRequest
	GetCrossRoleName() *string
	SetDestDatabaseInstanceClass(v string) *CreateRestoreTaskRequest
	GetDestDatabaseInstanceClass() *string
	SetDestDatabaseInstanceDatabaseVersion(v string) *CreateRestoreTaskRequest
	GetDestDatabaseInstanceDatabaseVersion() *string
	SetDestDatabaseInstanceRegion(v string) *CreateRestoreTaskRequest
	GetDestDatabaseInstanceRegion() *string
	SetDestDatabaseInstanceStorageSize(v string) *CreateRestoreTaskRequest
	GetDestDatabaseInstanceStorageSize() *string
	SetDestDatabaseInstanceType(v string) *CreateRestoreTaskRequest
	GetDestDatabaseInstanceType() *string
	SetDestDatabaseInstanceVSwitch(v string) *CreateRestoreTaskRequest
	GetDestDatabaseInstanceVSwitch() *string
	SetDestDatabaseInstanceVpc(v string) *CreateRestoreTaskRequest
	GetDestDatabaseInstanceVpc() *string
	SetDestinationEndpointDatabaseName(v string) *CreateRestoreTaskRequest
	GetDestinationEndpointDatabaseName() *string
	SetDestinationEndpointIP(v string) *CreateRestoreTaskRequest
	GetDestinationEndpointIP() *string
	SetDestinationEndpointInstanceID(v string) *CreateRestoreTaskRequest
	GetDestinationEndpointInstanceID() *string
	SetDestinationEndpointInstanceType(v string) *CreateRestoreTaskRequest
	GetDestinationEndpointInstanceType() *string
	SetDestinationEndpointOracleSID(v string) *CreateRestoreTaskRequest
	GetDestinationEndpointOracleSID() *string
	SetDestinationEndpointPassword(v string) *CreateRestoreTaskRequest
	GetDestinationEndpointPassword() *string
	SetDestinationEndpointPort(v int32) *CreateRestoreTaskRequest
	GetDestinationEndpointPort() *int32
	SetDestinationEndpointRegion(v string) *CreateRestoreTaskRequest
	GetDestinationEndpointRegion() *string
	SetDestinationEndpointUserName(v string) *CreateRestoreTaskRequest
	GetDestinationEndpointUserName() *string
	SetDuplicateConflict(v string) *CreateRestoreTaskRequest
	GetDuplicateConflict() *string
	SetEnableDestinationEndpointSsl(v bool) *CreateRestoreTaskRequest
	GetEnableDestinationEndpointSsl() *bool
	SetOwnerId(v string) *CreateRestoreTaskRequest
	GetOwnerId() *string
	SetRestoreDestinationMode(v string) *CreateRestoreTaskRequest
	GetRestoreDestinationMode() *string
	SetRestoreDir(v string) *CreateRestoreTaskRequest
	GetRestoreDir() *string
	SetRestoreHome(v string) *CreateRestoreTaskRequest
	GetRestoreHome() *string
	SetRestoreObjects(v string) *CreateRestoreTaskRequest
	GetRestoreObjects() *string
	SetRestoreTaskName(v string) *CreateRestoreTaskRequest
	GetRestoreTaskName() *string
	SetRestoreTime(v int64) *CreateRestoreTaskRequest
	GetRestoreTime() *int64
	SetSslCaPem(v string) *CreateRestoreTaskRequest
	GetSslCaPem() *string
}

type CreateRestoreTaskRequest struct {
	// Specifies whether to automatically start the instance after an Oracle physical backup is restored.
	//
	// example:
	//
	// true
	AutoOpenDatabase *string `json:"AutoOpenDatabase,omitempty" xml:"AutoOpenDatabase,omitempty"`
	// Specifies whether to automatically shut down the instance after an Oracle physical backup is restored.
	//
	// example:
	//
	// true
	AutoShutdownDatabase *string `json:"AutoShutdownDatabase,omitempty" xml:"AutoShutdownDatabase,omitempty"`
	// The ID of the backup gateway.
	//
	// > This parameter is required if **DestinationEndpointInstanceType*	- is set to Agent.
	//
	// example:
	//
	// 4312****
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The ID of the backup plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbs1hvb0ww****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the full backup set used for restoration. This parameter and RestoreTime are mutually exclusive.
	//
	// example:
	//
	// dbs1hvb0w*****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The UID of the Alibaba Cloud account used for cross-account backup.
	//
	// example:
	//
	// 2749528728********
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// The name of the RAM role used for cross-account backup.
	//
	// example:
	//
	// test123
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	// The instance specifications of the new database instance to which data is restored.
	//
	// example:
	//
	// mysql.x4.large.2
	DestDatabaseInstanceClass *string `json:"DestDatabaseInstanceClass,omitempty" xml:"DestDatabaseInstanceClass,omitempty"`
	// The database engine version of the new database instance to which data is restored.
	//
	// example:
	//
	// 8.0
	DestDatabaseInstanceDatabaseVersion *string `json:"DestDatabaseInstanceDatabaseVersion,omitempty" xml:"DestDatabaseInstanceDatabaseVersion,omitempty"`
	// The region of the new database instance to which data is restored.
	//
	// example:
	//
	// cn-beijing
	DestDatabaseInstanceRegion *string `json:"DestDatabaseInstanceRegion,omitempty" xml:"DestDatabaseInstanceRegion,omitempty"`
	// The instance storage size of the new database to recover.
	//
	// example:
	//
	// 500
	DestDatabaseInstanceStorageSize *string `json:"DestDatabaseInstanceStorageSize,omitempty" xml:"DestDatabaseInstanceStorageSize,omitempty"`
	// The instance type of the new database to which data is restored.
	//
	// example:
	//
	// rds
	DestDatabaseInstanceType *string `json:"DestDatabaseInstanceType,omitempty" xml:"DestDatabaseInstanceType,omitempty"`
	// The vSwitch of the new database instance to which data is restored.
	//
	// example:
	//
	// vsw-xx
	DestDatabaseInstanceVSwitch *string `json:"DestDatabaseInstanceVSwitch,omitempty" xml:"DestDatabaseInstanceVSwitch,omitempty"`
	// The VPC of the new database instance to which data is restored.
	//
	// example:
	//
	// vpc-xx
	DestDatabaseInstanceVpc *string `json:"DestDatabaseInstanceVpc,omitempty" xml:"DestDatabaseInstanceVpc,omitempty"`
	// The name of the database.
	//
	// > This parameter is required if the database engine is PostgreSQL or MongoDB.
	//
	// example:
	//
	// test
	DestinationEndpointDatabaseName *string `json:"DestinationEndpointDatabaseName,omitempty" xml:"DestinationEndpointDatabaseName,omitempty"`
	// The database connection address.
	//
	// > This parameter is required if **DestinationEndpointInstanceType*	- is set to Express, Agent, or Other.
	//
	// example:
	//
	// rm-bp*****9jv8pxero.mysql.rds.aliyuncs.com
	DestinationEndpointIP *string `json:"DestinationEndpointIP,omitempty" xml:"DestinationEndpointIP,omitempty"`
	// The instance ID of the database.
	//
	// > This parameter is required if **DestinationEndpointInstanceType*	- is set to RDS, ECS, DDS, or Express.
	//
	// example:
	//
	// rm-bp1p8c29*****
	DestinationEndpointInstanceID *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	// The location of the database. Valid values:
	//
	// - **RDS**
	//
	// - **ECS**
	//
	// - **Express**: databases connected over Express Connect circuits, VPN gateways, or Smart Access Gateway (SAG)
	//
	// - **Agent**: databases connected over backup gateways
	//
	// - **DDS**: ApsaraDB for MongoDB
	//
	// - **Other**: databases connected by using IP addresses and port numbers
	//
	// - **dg**: self-managed databases that have no public IP addresses or port numbers and are connected over Database Gateway
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	DestinationEndpointInstanceType *string `json:"DestinationEndpointInstanceType,omitempty" xml:"DestinationEndpointInstanceType,omitempty"`
	// The Oracle SID.
	//
	// > This parameter is required if the database engine is Oracle.
	//
	// example:
	//
	// test
	DestinationEndpointOracleSID *string `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	// The password of the database account.
	//
	// > This parameter is not required if the database engine is Redis, or if the database is connected over a backup gateway and the database engine is SQL Server. In other scenarios, this parameter is required.
	//
	// example:
	//
	// Test
	DestinationEndpointPassword *string `json:"DestinationEndpointPassword,omitempty" xml:"DestinationEndpointPassword,omitempty"`
	// The port number of the database.
	//
	// > This parameter is required if **DestinationEndpointInstanceType*	- is set to Express, Agent, Other, or ECS.
	//
	// example:
	//
	// 3306
	DestinationEndpointPort *int32 `json:"DestinationEndpointPort,omitempty" xml:"DestinationEndpointPort,omitempty"`
	// The region of the database instance.
	//
	// >This parameter is required if **DestinationEndpointInstanceType*	- is set to RDS, ECS, DDS, Express, or Agent.
	//
	// example:
	//
	// cn-hangzhou
	DestinationEndpointRegion *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	// The username of the database account.
	//
	// > This parameter is not required if the database engine is Redis, or if the database is connected over a backup gateway and the database engine is SQL Server. In other scenarios, this parameter is required.
	//
	// example:
	//
	// test
	DestinationEndpointUserName *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	// The method used to handle conflicts between objects with the same name. Valid values:
	//
	// - **renamenew**: renames the object if an object with the same name exists.
	//
	// example:
	//
	// renamenew
	DuplicateConflict *string `json:"DuplicateConflict,omitempty" xml:"DuplicateConflict,omitempty"`
	// Specifies whether to connect to the destination by using an SSL-encrypted connection.
	//
	// example:
	//
	// true
	EnableDestinationEndpointSsl *bool   `json:"EnableDestinationEndpointSsl,omitempty" xml:"EnableDestinationEndpointSsl,omitempty"`
	OwnerId                      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The method used to restore data to the destination. Valid values:
	//
	// - **EXIST_INSTANCE**: restores data to an existing instance.
	//
	// - **NEW_INSTANCE**: restores data to a new instance.
	//
	// example:
	//
	// EXIST_INSTANCE
	RestoreDestinationMode *string `json:"RestoreDestinationMode,omitempty" xml:"RestoreDestinationMode,omitempty"`
	// This parameter is required if **DestinationEndpointInstanceType*	- is set to Agent and the backup plan is for MySQL.
	//
	// example:
	//
	// test
	RestoreDir *string `json:"RestoreDir,omitempty" xml:"RestoreDir,omitempty"`
	// The directory of the database program.
	//
	// example:
	//
	// test
	RestoreHome *string `json:"RestoreHome,omitempty" xml:"RestoreHome,omitempty"`
	// The objects to be restored.
	//
	// - For more information, see the definition of the **RestoreObjects*	- parameter below. This parameter is not required if the database is connected over a backup gateway. In other scenarios, this parameter is required.
	//
	// - Template: `[{ "DBName": "Name of the database to be restored", "NewDBName": "Name of the destination database" }]`
	//
	// > Restoring objects by calling this API operation only supports restoration at the database level. To restore specific tables, go to the console. For more information, see [Restore a database](https://help.aliyun.com/document_detail/85543.html).
	//
	// example:
	//
	// Example of MySQL table-level restoration:
	//
	// [{\\"DBName\\":\\"dbname\\", \\"NewDBName\\":\\"dbname1\\"}]
	RestoreObjects *string `json:"RestoreObjects,omitempty" xml:"RestoreObjects,omitempty"`
	// The name of the restore task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RestoreTaskName *string `json:"RestoreTaskName,omitempty" xml:"RestoreTaskName,omitempty"`
	// The point in time to which you want to restore the database. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1554560477000
	RestoreTime *int64 `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	// The credential for the SSL-encrypted connection to the destination.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- ... -----END CERTIFICATE-----
	SslCaPem *string `json:"SslCaPem,omitempty" xml:"SslCaPem,omitempty"`
}

func (s CreateRestoreTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRestoreTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRestoreTaskRequest) GetAutoOpenDatabase() *string {
	return s.AutoOpenDatabase
}

func (s *CreateRestoreTaskRequest) GetAutoShutdownDatabase() *string {
	return s.AutoShutdownDatabase
}

func (s *CreateRestoreTaskRequest) GetBackupGatewayId() *int64 {
	return s.BackupGatewayId
}

func (s *CreateRestoreTaskRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *CreateRestoreTaskRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *CreateRestoreTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRestoreTaskRequest) GetCrossAliyunId() *string {
	return s.CrossAliyunId
}

func (s *CreateRestoreTaskRequest) GetCrossRoleName() *string {
	return s.CrossRoleName
}

func (s *CreateRestoreTaskRequest) GetDestDatabaseInstanceClass() *string {
	return s.DestDatabaseInstanceClass
}

func (s *CreateRestoreTaskRequest) GetDestDatabaseInstanceDatabaseVersion() *string {
	return s.DestDatabaseInstanceDatabaseVersion
}

func (s *CreateRestoreTaskRequest) GetDestDatabaseInstanceRegion() *string {
	return s.DestDatabaseInstanceRegion
}

func (s *CreateRestoreTaskRequest) GetDestDatabaseInstanceStorageSize() *string {
	return s.DestDatabaseInstanceStorageSize
}

func (s *CreateRestoreTaskRequest) GetDestDatabaseInstanceType() *string {
	return s.DestDatabaseInstanceType
}

func (s *CreateRestoreTaskRequest) GetDestDatabaseInstanceVSwitch() *string {
	return s.DestDatabaseInstanceVSwitch
}

func (s *CreateRestoreTaskRequest) GetDestDatabaseInstanceVpc() *string {
	return s.DestDatabaseInstanceVpc
}

func (s *CreateRestoreTaskRequest) GetDestinationEndpointDatabaseName() *string {
	return s.DestinationEndpointDatabaseName
}

func (s *CreateRestoreTaskRequest) GetDestinationEndpointIP() *string {
	return s.DestinationEndpointIP
}

func (s *CreateRestoreTaskRequest) GetDestinationEndpointInstanceID() *string {
	return s.DestinationEndpointInstanceID
}

func (s *CreateRestoreTaskRequest) GetDestinationEndpointInstanceType() *string {
	return s.DestinationEndpointInstanceType
}

func (s *CreateRestoreTaskRequest) GetDestinationEndpointOracleSID() *string {
	return s.DestinationEndpointOracleSID
}

func (s *CreateRestoreTaskRequest) GetDestinationEndpointPassword() *string {
	return s.DestinationEndpointPassword
}

func (s *CreateRestoreTaskRequest) GetDestinationEndpointPort() *int32 {
	return s.DestinationEndpointPort
}

func (s *CreateRestoreTaskRequest) GetDestinationEndpointRegion() *string {
	return s.DestinationEndpointRegion
}

func (s *CreateRestoreTaskRequest) GetDestinationEndpointUserName() *string {
	return s.DestinationEndpointUserName
}

func (s *CreateRestoreTaskRequest) GetDuplicateConflict() *string {
	return s.DuplicateConflict
}

func (s *CreateRestoreTaskRequest) GetEnableDestinationEndpointSsl() *bool {
	return s.EnableDestinationEndpointSsl
}

func (s *CreateRestoreTaskRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateRestoreTaskRequest) GetRestoreDestinationMode() *string {
	return s.RestoreDestinationMode
}

func (s *CreateRestoreTaskRequest) GetRestoreDir() *string {
	return s.RestoreDir
}

func (s *CreateRestoreTaskRequest) GetRestoreHome() *string {
	return s.RestoreHome
}

func (s *CreateRestoreTaskRequest) GetRestoreObjects() *string {
	return s.RestoreObjects
}

func (s *CreateRestoreTaskRequest) GetRestoreTaskName() *string {
	return s.RestoreTaskName
}

func (s *CreateRestoreTaskRequest) GetRestoreTime() *int64 {
	return s.RestoreTime
}

func (s *CreateRestoreTaskRequest) GetSslCaPem() *string {
	return s.SslCaPem
}

func (s *CreateRestoreTaskRequest) SetAutoOpenDatabase(v string) *CreateRestoreTaskRequest {
	s.AutoOpenDatabase = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetAutoShutdownDatabase(v string) *CreateRestoreTaskRequest {
	s.AutoShutdownDatabase = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetBackupGatewayId(v int64) *CreateRestoreTaskRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetBackupPlanId(v string) *CreateRestoreTaskRequest {
	s.BackupPlanId = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetBackupSetId(v string) *CreateRestoreTaskRequest {
	s.BackupSetId = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetClientToken(v string) *CreateRestoreTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetCrossAliyunId(v string) *CreateRestoreTaskRequest {
	s.CrossAliyunId = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetCrossRoleName(v string) *CreateRestoreTaskRequest {
	s.CrossRoleName = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestDatabaseInstanceClass(v string) *CreateRestoreTaskRequest {
	s.DestDatabaseInstanceClass = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestDatabaseInstanceDatabaseVersion(v string) *CreateRestoreTaskRequest {
	s.DestDatabaseInstanceDatabaseVersion = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestDatabaseInstanceRegion(v string) *CreateRestoreTaskRequest {
	s.DestDatabaseInstanceRegion = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestDatabaseInstanceStorageSize(v string) *CreateRestoreTaskRequest {
	s.DestDatabaseInstanceStorageSize = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestDatabaseInstanceType(v string) *CreateRestoreTaskRequest {
	s.DestDatabaseInstanceType = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestDatabaseInstanceVSwitch(v string) *CreateRestoreTaskRequest {
	s.DestDatabaseInstanceVSwitch = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestDatabaseInstanceVpc(v string) *CreateRestoreTaskRequest {
	s.DestDatabaseInstanceVpc = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointDatabaseName(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointDatabaseName = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointIP(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointIP = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointInstanceID(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointInstanceID = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointInstanceType(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointInstanceType = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointOracleSID(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointOracleSID = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointPassword(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointPassword = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointPort(v int32) *CreateRestoreTaskRequest {
	s.DestinationEndpointPort = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointRegion(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointRegion = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDestinationEndpointUserName(v string) *CreateRestoreTaskRequest {
	s.DestinationEndpointUserName = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetDuplicateConflict(v string) *CreateRestoreTaskRequest {
	s.DuplicateConflict = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetEnableDestinationEndpointSsl(v bool) *CreateRestoreTaskRequest {
	s.EnableDestinationEndpointSsl = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetOwnerId(v string) *CreateRestoreTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetRestoreDestinationMode(v string) *CreateRestoreTaskRequest {
	s.RestoreDestinationMode = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetRestoreDir(v string) *CreateRestoreTaskRequest {
	s.RestoreDir = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetRestoreHome(v string) *CreateRestoreTaskRequest {
	s.RestoreHome = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetRestoreObjects(v string) *CreateRestoreTaskRequest {
	s.RestoreObjects = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetRestoreTaskName(v string) *CreateRestoreTaskRequest {
	s.RestoreTaskName = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetRestoreTime(v int64) *CreateRestoreTaskRequest {
	s.RestoreTime = &v
	return s
}

func (s *CreateRestoreTaskRequest) SetSslCaPem(v string) *CreateRestoreTaskRequest {
	s.SslCaPem = &v
	return s
}

func (s *CreateRestoreTaskRequest) Validate() error {
	return dara.Validate(s)
}
