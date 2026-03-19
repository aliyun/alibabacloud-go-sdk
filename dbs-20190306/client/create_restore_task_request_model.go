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
	AutoOpenDatabase     *string `json:"AutoOpenDatabase,omitempty" xml:"AutoOpenDatabase,omitempty"`
	AutoShutdownDatabase *string `json:"AutoShutdownDatabase,omitempty" xml:"AutoShutdownDatabase,omitempty"`
	// backup gateway ID.
	//
	// > This parameter is required when **DestinationEndpointInstanceType*	- is agent.
	//
	// example:
	//
	// 4312****
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// backup plan ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbs1hvb0ww****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the full backup set used for restoration. Mutually exclusive with RestoreTime.
	//
	// example:
	//
	// dbs1hvb0w*****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// Ensures request idempotence and prevents duplicate submissions.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// UID for cross-Alibaba Cloud account backup.
	//
	// example:
	//
	// 2749528728********
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// RAM role name for cross-Alibaba Cloud account backup.
	//
	// example:
	//
	// test123
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	// example:
	//
	// mysql.x4.large.2
	DestDatabaseInstanceClass           *string `json:"DestDatabaseInstanceClass,omitempty" xml:"DestDatabaseInstanceClass,omitempty"`
	DestDatabaseInstanceDatabaseVersion *string `json:"DestDatabaseInstanceDatabaseVersion,omitempty" xml:"DestDatabaseInstanceDatabaseVersion,omitempty"`
	// example:
	//
	// cn-beijing
	DestDatabaseInstanceRegion *string `json:"DestDatabaseInstanceRegion,omitempty" xml:"DestDatabaseInstanceRegion,omitempty"`
	// example:
	//
	// 500
	DestDatabaseInstanceStorageSize *string `json:"DestDatabaseInstanceStorageSize,omitempty" xml:"DestDatabaseInstanceStorageSize,omitempty"`
	// example:
	//
	// rds
	DestDatabaseInstanceType    *string `json:"DestDatabaseInstanceType,omitempty" xml:"DestDatabaseInstanceType,omitempty"`
	DestDatabaseInstanceVSwitch *string `json:"DestDatabaseInstanceVSwitch,omitempty" xml:"DestDatabaseInstanceVSwitch,omitempty"`
	// example:
	//
	// vpc-xx
	DestDatabaseInstanceVpc *string `json:"DestDatabaseInstanceVpc,omitempty" xml:"DestDatabaseInstanceVpc,omitempty"`
	// database name.
	//
	// > This parameter is required when the database type is PostgreSQL or MongoDB.
	//
	// example:
	//
	// test
	DestinationEndpointDatabaseName *string `json:"DestinationEndpointDatabaseName,omitempty" xml:"DestinationEndpointDatabaseName,omitempty"`
	// database endpoint.
	//
	// > This parameter is required when **DestinationEndpointInstanceType*	- is express, agent, or other.
	//
	// example:
	//
	// rm-bp*****9jv8pxero.mysql.rds.aliyuncs.com
	DestinationEndpointIP *string `json:"DestinationEndpointIP,omitempty" xml:"DestinationEndpointIP,omitempty"`
	// database instance ID.
	//
	// > This parameter is required when **DestinationEndpointInstanceType*	- is RDS, ECS, DDS, or Express.
	//
	// example:
	//
	// rm-bp1p8c29*****
	DestinationEndpointInstanceID *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	// database location. Valid values:
	//
	// - **RDS**
	//
	// - **ECS**
	//
	// - **Express**: databases accessed via leased line/VPN Gateway/Smart Gateway
	//
	// - **Agent**: databases accessed via backup gateway
	//
	// - **DDS**: Cloud MongoDB
	//
	// - **Other**: databases directly connected via IP:Port
	//
	// - **dg**: self-managed databases without public IP:Port (accessed via Database Gateway DG)
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	DestinationEndpointInstanceType *string `json:"DestinationEndpointInstanceType,omitempty" xml:"DestinationEndpointInstanceType,omitempty"`
	// Oracle SID name.
	//
	// > This parameter is required when the database type is Oracle.
	//
	// example:
	//
	// test
	DestinationEndpointOracleSID *string `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	// password.
	//
	// > This parameter is optional when the database type is Redis, or when the database location is agent and the database type is MSSQL. It is required in all other scenarios.
	//
	// example:
	//
	// Test
	DestinationEndpointPassword *string `json:"DestinationEndpointPassword,omitempty" xml:"DestinationEndpointPassword,omitempty"`
	// database port.
	//
	// > This parameter is required when **DestinationEndpointInstanceType*	- is express, agent, other, or ECS.
	//
	// example:
	//
	// 3306
	DestinationEndpointPort *int32 `json:"DestinationEndpointPort,omitempty" xml:"DestinationEndpointPort,omitempty"`
	// region of the database instance.
	//
	// > This parameter is required when **DestinationEndpointInstanceType*	- is RDS, ECS, DDS, Express, or Agent.
	//
	// example:
	//
	// cn-hangzhou
	DestinationEndpointRegion *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	// database account.
	//
	// > This parameter is optional when the database type is Redis, or when the database location is agent and the database type is MSSQL. It is required in all other scenarios.
	//
	// example:
	//
	// test
	DestinationEndpointUserName *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	// Conflict handling for objects with the same name. Currently supports:
	//
	// **renamenew**: Rename objects if names conflict.
	//
	// example:
	//
	// renamenew
	DuplicateConflict            *string `json:"DuplicateConflict,omitempty" xml:"DuplicateConflict,omitempty"`
	EnableDestinationEndpointSsl *bool   `json:"EnableDestinationEndpointSsl,omitempty" xml:"EnableDestinationEndpointSsl,omitempty"`
	OwnerId                      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// exist_instance
	RestoreDestinationMode *string `json:"RestoreDestinationMode,omitempty" xml:"RestoreDestinationMode,omitempty"`
	// Required when **DestinationEndpointInstanceType*	- is agent and the backup plan is MySQL.
	//
	// example:
	//
	// test
	RestoreDir *string `json:"RestoreDir,omitempty" xml:"RestoreDir,omitempty"`
	// database program directory.
	//
	// example:
	//
	// test
	RestoreHome *string `json:"RestoreHome,omitempty" xml:"RestoreHome,omitempty"`
	// restore objects.
	//
	// - For details, see the **RestoreObjects*	- parameter definition below. This parameter is optional when the database location is agent, and required in all other scenarios.
	//
	// - Input template: `[{ "DBName": "database name to be restored", "NewDBName": "target database name to be restored" }]`
	//
	// > This API only supports restoring objects at the database level. To configure table-level restoration, use the console. For details, see [Recover databases](https://help.aliyun.com/document_detail/85543.html).
	//
	// example:
	//
	// MySQL表级别恢复示例如下：
	//
	// [{\\"DBName\\":\\"dbname\\", \\"NewDBName\\":\\"dbname1\\"}]
	RestoreObjects *string `json:"RestoreObjects,omitempty" xml:"RestoreObjects,omitempty"`
	// restore job name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RestoreTaskName *string `json:"RestoreTaskName,omitempty" xml:"RestoreTaskName,omitempty"`
	// restore time. Value: 1554560477000.
	//
	// example:
	//
	// 1554560477000
	RestoreTime *int64 `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
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
