// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestoreTaskRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetOwnerId(v string) *CreateRestoreTaskRequest
	GetOwnerId() *string
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
}

type CreateRestoreTaskRequest struct {
	// The ID of the backup gateway.
	//
	// > This parameter is required if the DestinationEndpointInstanceType parameter is set to Agent.
	//
	// example:
	//
	// 4312****
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The ID of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbs1hvb0ww****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The ID of the full backup set.
	//
	// example:
	//
	// dbs1hvb0w*****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC********
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The unique ID (UID) of the Alibaba Cloud account to which the source database belongs.
	//
	// example:
	//
	// 2749528728********
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// The name of the RAM role that is used to perform backups across Alibaba Cloud accounts.
	//
	// example:
	//
	// test123
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	// The name of the database.
	//
	//
	//
	// > This parameter is required if the database is a PostgreSQL database or a MongoDB database.
	//
	// example:
	//
	// test
	DestinationEndpointDatabaseName *string `json:"DestinationEndpointDatabaseName,omitempty" xml:"DestinationEndpointDatabaseName,omitempty"`
	// The endpoint that is used to connect to the database.
	//
	// > This parameter is required if the DestinationEndpointInstanceType parameter is set to Express, Agent, or Other.
	//
	// example:
	//
	// rm-bp*****9jv8pxero.mysql.rds.aliyuncs.com
	DestinationEndpointIP *string `json:"DestinationEndpointIP,omitempty" xml:"DestinationEndpointIP,omitempty"`
	// The ID of the database instance.
	//
	// > This parameter is required if the DestinationEndpointInstanceType parameter is set to RDS, ECS, DDS, or Express.
	//
	// example:
	//
	// rm-bp1p8c29*****
	DestinationEndpointInstanceID *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	// The location of the database. Valid values:
	//
	// 	- **RDS**: The database is deployed on an ApsaraDB RDS instance.
	//
	// 	- **ECS**: The database is deployed on an Elastic Compute Service (ECS) instance.
	//
	// 	- **Express**: The database is connected to Database Backup (DBS) by using Express Connect, VPN Gateway, or Smart Access Gateway.
	//
	// 	- **Agent**: The database is connected over a DBS backup gateway.
	//
	// 	- **DDS**: The database is an ApsaraDB for MongoDB database.
	//
	// 	- **Other**: The database is connected to DBS by using the IP address and port of the database.
	//
	// 	- **dg**: The database is a self-managed database that does not have public IP addresses or port numbers and is connected to DBS over Database Gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	DestinationEndpointInstanceType *string `json:"DestinationEndpointInstanceType,omitempty" xml:"DestinationEndpointInstanceType,omitempty"`
	// The system ID (SID) of the Oracle database.
	//
	//
	//
	// > This parameter is required if the source database is an Oracle database.
	//
	// example:
	//
	// test
	DestinationEndpointOracleSID *string `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	// The password of the account that is used to connect to the source database.
	//
	//
	//
	// > This parameter is required except that the database is an SQL Server database that is connected to DBS over a DBS backup gateway or a Redis database.
	//
	// example:
	//
	// Test
	DestinationEndpointPassword *string `json:"DestinationEndpointPassword,omitempty" xml:"DestinationEndpointPassword,omitempty"`
	// The port of the database.
	//
	// > This parameter is required if the DestinationEndpointInstanceType parameter is set to Express, Agent, Other, or ECS.
	//
	// example:
	//
	// 3306
	DestinationEndpointPort *int32 `json:"DestinationEndpointPort,omitempty" xml:"DestinationEndpointPort,omitempty"`
	// The region ID of the destination database instance.
	//
	// >  You must specify this parameter if **DestinationEndpointInstanceType*	- is set to RDS, ECS, DDS, Express, or Agent.
	//
	// example:
	//
	// cn-hangzhou
	DestinationEndpointRegion *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	// The username of the account that is used to connect to the database.
	//
	//
	// > This parameter is required except that the database is an SQL Server database that is connected to DBS over a DBS backup gateway or a Redis database.
	//
	// example:
	//
	// test
	DestinationEndpointUserName *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	// The method of processing objects with the same name. Valid values:
	//
	// - failure: The restore task fails if the system detects objects with the same name. This is the default value.
	//
	// - renamenew: The restore task renames objects with the same name starting from the second occurrence.
	//
	// example:
	//
	// renamenew
	DuplicateConflict *string `json:"DuplicateConflict,omitempty" xml:"DuplicateConflict,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required if the DestinationEndpointInstanceType parameter is set to Agent and the backup object of the backup schedule is a MySQL database.
	//
	// example:
	//
	// test
	RestoreDir *string `json:"RestoreDir,omitempty" xml:"RestoreDir,omitempty"`
	// The program directory of the database.
	//
	// example:
	//
	// test
	RestoreHome *string `json:"RestoreHome,omitempty" xml:"RestoreHome,omitempty"`
	// The objects to be restored.
	//
	//
	//
	// > This parameter is required except that the DestinationEndpointInstanceType parameter is set to Agent. For information about the parameter definition, see RestoreObjects.
	//
	// example:
	//
	// [ { "DBName":"Name of the database to be restored", "NewDBName":"Name of the database to which the objects will be restored", "SchemaName":"Schema name of the database to be restored", "NewSchemaName":"Schema name of the destination database to which the objects will be restored"}]
	RestoreObjects *string `json:"RestoreObjects,omitempty" xml:"RestoreObjects,omitempty"`
	// The name of the restore task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RestoreTaskName *string `json:"RestoreTaskName,omitempty" xml:"RestoreTaskName,omitempty"`
	// The time to run the restore task, such as 1554560477000.
	//
	// example:
	//
	// 1554560477000
	RestoreTime *int64 `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
}

func (s CreateRestoreTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRestoreTaskRequest) GoString() string {
	return s.String()
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

func (s *CreateRestoreTaskRequest) GetOwnerId() *string {
	return s.OwnerId
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

func (s *CreateRestoreTaskRequest) SetOwnerId(v string) *CreateRestoreTaskRequest {
	s.OwnerId = &v
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

func (s *CreateRestoreTaskRequest) Validate() error {
	return dara.Validate(s)
}
