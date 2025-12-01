// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoStartBackup(v bool) *ConfigureBackupPlanRequest
	GetAutoStartBackup() *bool
	SetBackupGatewayId(v int64) *ConfigureBackupPlanRequest
	GetBackupGatewayId() *int64
	SetBackupLogIntervalSeconds(v int32) *ConfigureBackupPlanRequest
	GetBackupLogIntervalSeconds() *int32
	SetBackupObjects(v string) *ConfigureBackupPlanRequest
	GetBackupObjects() *string
	SetBackupPeriod(v string) *ConfigureBackupPlanRequest
	GetBackupPeriod() *string
	SetBackupPlanId(v string) *ConfigureBackupPlanRequest
	GetBackupPlanId() *string
	SetBackupPlanName(v string) *ConfigureBackupPlanRequest
	GetBackupPlanName() *string
	SetBackupRateLimit(v int64) *ConfigureBackupPlanRequest
	GetBackupRateLimit() *int64
	SetBackupRetentionPeriod(v int32) *ConfigureBackupPlanRequest
	GetBackupRetentionPeriod() *int32
	SetBackupSpeedLimit(v int64) *ConfigureBackupPlanRequest
	GetBackupSpeedLimit() *int64
	SetBackupStartTime(v string) *ConfigureBackupPlanRequest
	GetBackupStartTime() *string
	SetBackupStorageType(v string) *ConfigureBackupPlanRequest
	GetBackupStorageType() *string
	SetBackupStrategyType(v string) *ConfigureBackupPlanRequest
	GetBackupStrategyType() *string
	SetClientToken(v string) *ConfigureBackupPlanRequest
	GetClientToken() *string
	SetCrossAliyunId(v string) *ConfigureBackupPlanRequest
	GetCrossAliyunId() *string
	SetCrossRoleName(v string) *ConfigureBackupPlanRequest
	GetCrossRoleName() *string
	SetDuplicationArchivePeriod(v int32) *ConfigureBackupPlanRequest
	GetDuplicationArchivePeriod() *int32
	SetDuplicationInfrequentAccessPeriod(v int32) *ConfigureBackupPlanRequest
	GetDuplicationInfrequentAccessPeriod() *int32
	SetEnableBackupLog(v bool) *ConfigureBackupPlanRequest
	GetEnableBackupLog() *bool
	SetOSSBucketName(v string) *ConfigureBackupPlanRequest
	GetOSSBucketName() *string
	SetOwnerId(v string) *ConfigureBackupPlanRequest
	GetOwnerId() *string
	SetResourceGroupId(v string) *ConfigureBackupPlanRequest
	GetResourceGroupId() *string
	SetSourceEndpointDatabaseName(v string) *ConfigureBackupPlanRequest
	GetSourceEndpointDatabaseName() *string
	SetSourceEndpointIP(v string) *ConfigureBackupPlanRequest
	GetSourceEndpointIP() *string
	SetSourceEndpointInstanceID(v string) *ConfigureBackupPlanRequest
	GetSourceEndpointInstanceID() *string
	SetSourceEndpointInstanceType(v string) *ConfigureBackupPlanRequest
	GetSourceEndpointInstanceType() *string
	SetSourceEndpointOracleSID(v string) *ConfigureBackupPlanRequest
	GetSourceEndpointOracleSID() *string
	SetSourceEndpointPassword(v string) *ConfigureBackupPlanRequest
	GetSourceEndpointPassword() *string
	SetSourceEndpointPort(v int32) *ConfigureBackupPlanRequest
	GetSourceEndpointPort() *int32
	SetSourceEndpointRegion(v string) *ConfigureBackupPlanRequest
	GetSourceEndpointRegion() *string
	SetSourceEndpointUserName(v string) *ConfigureBackupPlanRequest
	GetSourceEndpointUserName() *string
}

type ConfigureBackupPlanRequest struct {
	// Specifies whether to enable the automatic backup feature.
	//
	// 	- **true**: enables the automatic backup feature.
	//
	// 	- **false**: disables the automatic backup feature.
	//
	// example:
	//
	// false
	AutoStartBackup *bool `json:"AutoStartBackup,omitempty" xml:"AutoStartBackup,omitempty"`
	// The backup gateway ID. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain it.
	//
	// >  If you set **SourceEndpointInstanceType*	- to **Agent**, this parameter is required.
	//
	// example:
	//
	// 23313123312
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The interval at which you want to perform incremental log backups. Unit: seconds.
	//
	// >  Only physical backup supports this parameter.
	//
	// example:
	//
	// 1000
	BackupLogIntervalSeconds *int32 `json:"BackupLogIntervalSeconds,omitempty" xml:"BackupLogIntervalSeconds,omitempty"`
	// The objects to be backed up. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the objects.
	//
	// example:
	//
	// [ { "DBName":"Name of the database to be backed up", "SchemaName":"Name of the schema to be backed up", "TableIncludes":[{ "TableName":"Name of the table to be backed up" }], "TableExcludes":[{ "TableName":"Name of the table to be excluded during the backup" }] } ]
	BackupObjects *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	// The day of each week when the full backup task runs. Valid values:
	//
	// 	- **Monday**
	//
	// 	- **Tuesday**
	//
	// 	- **Wednesday**
	//
	// 	- **Thursday**
	//
	// 	- **Friday**
	//
	// 	- **Saturday**
	//
	// 	- **Sunday**
	//
	// example:
	//
	// Monday
	BackupPeriod *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain it.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi0*******
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The name of the backup schedule. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi0*******
	BackupPlanName *string `json:"BackupPlanName,omitempty" xml:"BackupPlanName,omitempty"`
	// The network bandwidth throttling. Unit: KB/s. DBS allows a maximum bandwidth of 10 GB/s.
	//
	// > This parameter takes effect only when physical backups for MySQL databases are performed.
	//
	// example:
	//
	// 262144
	BackupRateLimit *int64 `json:"BackupRateLimit,omitempty" xml:"BackupRateLimit,omitempty"`
	// The number of days for which the backup data is retained. Valid values: 0 to 1825. Default value: 730.
	//
	// example:
	//
	// 730
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The disk I/O limit. Unit: KB/s.
	//
	// >  This parameter takes effect only during the physical backup of a MySQL database.
	//
	// example:
	//
	// 262144
	BackupSpeedLimit *int64 `json:"BackupSpeedLimit,omitempty" xml:"BackupSpeedLimit,omitempty"`
	// The start time of the full backup. Specify the time in the *HH:mm*Z format. The time must be in UTC. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the start time of full backup tasks.
	//
	// example:
	//
	// 14:22
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The storage type. Valid values:
	//
	// 	- Empty: If you do not specify this parameter, the system stores backup data in your OSS bucket.
	//
	// 	- system: The system stores backup data in the built-in OSS bucket of DBS.
	//
	// example:
	//
	// N/A
	BackupStorageType *string `json:"BackupStorageType,omitempty" xml:"BackupStorageType,omitempty"`
	// The backup method that you want to use for full backups. Valid values:
	//
	// 	- **simple**: scheduled backup. If you specify this value for the BackupStrategyType parameter, you must also specify the BackupPeriod and BackupStartTime parameters.
	//
	// 	- **Manual**: manual backup.
	//
	// > Default value: **simple**.
	//
	// example:
	//
	// simple
	BackupStrategyType *string `json:"BackupStrategyType,omitempty" xml:"BackupStrategyType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The unique ID (UID) of the Alibaba Cloud account to which the backup schedule belongs. You can call the [DescribeRestoreTaskList](https://help.aliyun.com/document_detail/2869838.html) operation to obtain the UID.
	//
	// example:
	//
	// 2xxx7778xxxxxxxxxx
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// The name of the RAM role that is used to perform backup across Alibaba Cloud accounts. You can call the [DescribeRestoreTaskList](https://help.aliyun.com/document_detail/2869838.html) operation to obtain the RAM role.
	//
	// example:
	//
	// test123
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	// The number of days after which the storage class of the backup data is changed to Archive. Default value: 365.
	//
	// example:
	//
	// 365
	DuplicationArchivePeriod *int32 `json:"DuplicationArchivePeriod,omitempty" xml:"DuplicationArchivePeriod,omitempty"`
	// The number of days after which the storage class of the backup data is changed to Infrequent Access (IA). Default value: 180.
	//
	// example:
	//
	// 180
	DuplicationInfrequentAccessPeriod *int32 `json:"DuplicationInfrequentAccessPeriod,omitempty" xml:"DuplicationInfrequentAccessPeriod,omitempty"`
	// Specifies whether to enable the incremental log backup feature. Valid values:
	//
	// 	- **true**: enables the incremental log backup feature.
	//
	// 	- **false**: disables the incremental log backup feature.
	//
	// example:
	//
	// true
	EnableBackupLog *bool `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// The name of the OSS bucket.
	//
	// >  By default, the system automatically generates an OSS bucket name.
	//
	// example:
	//
	// TestOssBucket
	OSSBucketName *string `json:"OSSBucketName,omitempty" xml:"OSSBucketName,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzecovzti****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The source database name. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain it.
	//
	// >  If the source database runs the **PostgreSQL*	- database engine or **MongoDB*	- database engine, this parameter is required.
	//
	// example:
	//
	// testRDS
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// The source database endpoint. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain it.
	//
	// >  If you set **SourceEndpointInstanceType*	- to **Express**, **Agent**, or **Other**, this parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*******.mysql.rds.aliyuncs.com
	SourceEndpointIP *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	// The database instance ID. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the ID.
	//
	// >  If you set **SourceEndpoint****InstanceType*	- to **RDS**, **ECS**, **DDS**, or **Express**, this parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*********
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The location of the database. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the location. Valid values:
	//
	// 	- **RDS**
	//
	// 	- **ECS**
	//
	// 	- **Express**: The database is connected to Database Backup (DBS) via Express Connect, VPN Gateway, or Smart Access Gateway.
	//
	// 	- **Agent**: The database is connected over a DBS backup gateway.
	//
	// 	- **DDS**: The database is an ApsaraDB for MongoDB database.
	//
	// 	- **Other**: The database is connected to DBS by using the IP address and port of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// The system ID (SID) of the Oracle database.
	//
	// > This parameter is required if the database is an Oracle database.
	//
	// example:
	//
	// test
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// The password of the account that is used to connect to the database.
	//
	// > This parameter is required except that the database is an **SQL Server*	- database that is connected to DBS over a DBS backup gateway or a **Redis*	- database.
	//
	// example:
	//
	// testPassword
	SourceEndpointPassword *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	// The port that is used to connect to the source database. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the port.
	//
	// >  If you set **SourceEndpoint****InstanceType*	- to **Express**, **Agent**, **Other**, or **ECS**, this parameter is required.
	//
	// example:
	//
	// 3306
	SourceEndpointPort *int32 `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// The region in which the source database resides. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the region.
	//
	// >  If you set **SourceEndpointInstanceType*	- to RDS, ECS, DDS, Express, or Agent, this parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The username of the account that is used to connect to the database.
	//
	// > This parameter is required except that the database is an **SQL Server*	- database that is connected to DBS over a DBS backup gateway or a **Redis*	- database.
	//
	// example:
	//
	// testRDS
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
}

func (s ConfigureBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *ConfigureBackupPlanRequest) GetAutoStartBackup() *bool {
	return s.AutoStartBackup
}

func (s *ConfigureBackupPlanRequest) GetBackupGatewayId() *int64 {
	return s.BackupGatewayId
}

func (s *ConfigureBackupPlanRequest) GetBackupLogIntervalSeconds() *int32 {
	return s.BackupLogIntervalSeconds
}

func (s *ConfigureBackupPlanRequest) GetBackupObjects() *string {
	return s.BackupObjects
}

func (s *ConfigureBackupPlanRequest) GetBackupPeriod() *string {
	return s.BackupPeriod
}

func (s *ConfigureBackupPlanRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ConfigureBackupPlanRequest) GetBackupPlanName() *string {
	return s.BackupPlanName
}

func (s *ConfigureBackupPlanRequest) GetBackupRateLimit() *int64 {
	return s.BackupRateLimit
}

func (s *ConfigureBackupPlanRequest) GetBackupRetentionPeriod() *int32 {
	return s.BackupRetentionPeriod
}

func (s *ConfigureBackupPlanRequest) GetBackupSpeedLimit() *int64 {
	return s.BackupSpeedLimit
}

func (s *ConfigureBackupPlanRequest) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *ConfigureBackupPlanRequest) GetBackupStorageType() *string {
	return s.BackupStorageType
}

func (s *ConfigureBackupPlanRequest) GetBackupStrategyType() *string {
	return s.BackupStrategyType
}

func (s *ConfigureBackupPlanRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ConfigureBackupPlanRequest) GetCrossAliyunId() *string {
	return s.CrossAliyunId
}

func (s *ConfigureBackupPlanRequest) GetCrossRoleName() *string {
	return s.CrossRoleName
}

func (s *ConfigureBackupPlanRequest) GetDuplicationArchivePeriod() *int32 {
	return s.DuplicationArchivePeriod
}

func (s *ConfigureBackupPlanRequest) GetDuplicationInfrequentAccessPeriod() *int32 {
	return s.DuplicationInfrequentAccessPeriod
}

func (s *ConfigureBackupPlanRequest) GetEnableBackupLog() *bool {
	return s.EnableBackupLog
}

func (s *ConfigureBackupPlanRequest) GetOSSBucketName() *string {
	return s.OSSBucketName
}

func (s *ConfigureBackupPlanRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ConfigureBackupPlanRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigureBackupPlanRequest) GetSourceEndpointDatabaseName() *string {
	return s.SourceEndpointDatabaseName
}

func (s *ConfigureBackupPlanRequest) GetSourceEndpointIP() *string {
	return s.SourceEndpointIP
}

func (s *ConfigureBackupPlanRequest) GetSourceEndpointInstanceID() *string {
	return s.SourceEndpointInstanceID
}

func (s *ConfigureBackupPlanRequest) GetSourceEndpointInstanceType() *string {
	return s.SourceEndpointInstanceType
}

func (s *ConfigureBackupPlanRequest) GetSourceEndpointOracleSID() *string {
	return s.SourceEndpointOracleSID
}

func (s *ConfigureBackupPlanRequest) GetSourceEndpointPassword() *string {
	return s.SourceEndpointPassword
}

func (s *ConfigureBackupPlanRequest) GetSourceEndpointPort() *int32 {
	return s.SourceEndpointPort
}

func (s *ConfigureBackupPlanRequest) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *ConfigureBackupPlanRequest) GetSourceEndpointUserName() *string {
	return s.SourceEndpointUserName
}

func (s *ConfigureBackupPlanRequest) SetAutoStartBackup(v bool) *ConfigureBackupPlanRequest {
	s.AutoStartBackup = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupGatewayId(v int64) *ConfigureBackupPlanRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupLogIntervalSeconds(v int32) *ConfigureBackupPlanRequest {
	s.BackupLogIntervalSeconds = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupObjects(v string) *ConfigureBackupPlanRequest {
	s.BackupObjects = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupPeriod(v string) *ConfigureBackupPlanRequest {
	s.BackupPeriod = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupPlanId(v string) *ConfigureBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupPlanName(v string) *ConfigureBackupPlanRequest {
	s.BackupPlanName = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupRateLimit(v int64) *ConfigureBackupPlanRequest {
	s.BackupRateLimit = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupRetentionPeriod(v int32) *ConfigureBackupPlanRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupSpeedLimit(v int64) *ConfigureBackupPlanRequest {
	s.BackupSpeedLimit = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupStartTime(v string) *ConfigureBackupPlanRequest {
	s.BackupStartTime = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupStorageType(v string) *ConfigureBackupPlanRequest {
	s.BackupStorageType = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetBackupStrategyType(v string) *ConfigureBackupPlanRequest {
	s.BackupStrategyType = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetClientToken(v string) *ConfigureBackupPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetCrossAliyunId(v string) *ConfigureBackupPlanRequest {
	s.CrossAliyunId = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetCrossRoleName(v string) *ConfigureBackupPlanRequest {
	s.CrossRoleName = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetDuplicationArchivePeriod(v int32) *ConfigureBackupPlanRequest {
	s.DuplicationArchivePeriod = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetDuplicationInfrequentAccessPeriod(v int32) *ConfigureBackupPlanRequest {
	s.DuplicationInfrequentAccessPeriod = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetEnableBackupLog(v bool) *ConfigureBackupPlanRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetOSSBucketName(v string) *ConfigureBackupPlanRequest {
	s.OSSBucketName = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetOwnerId(v string) *ConfigureBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetResourceGroupId(v string) *ConfigureBackupPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointDatabaseName(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointIP(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointInstanceID(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointInstanceType(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointOracleSID(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointPassword(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointPort(v int32) *ConfigureBackupPlanRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointRegion(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetSourceEndpointUserName(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *ConfigureBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
