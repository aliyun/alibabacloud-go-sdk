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
	SetBackupStorageEncryptMethod(v string) *ConfigureBackupPlanRequest
	GetBackupStorageEncryptMethod() *string
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
	SetEnableMysqlPhysicalBackupBinlog(v string) *ConfigureBackupPlanRequest
	GetEnableMysqlPhysicalBackupBinlog() *string
	SetEnableSourceEndpointSsl(v string) *ConfigureBackupPlanRequest
	GetEnableSourceEndpointSsl() *string
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
	SetSourceEndpointOracleHome(v string) *ConfigureBackupPlanRequest
	GetSourceEndpointOracleHome() *string
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
	SetSslCaPem(v string) *ConfigureBackupPlanRequest
	GetSslCaPem() *string
}

type ConfigureBackupPlanRequest struct {
	// Enable automatic backup. Valid values:
	//
	// - **true**: Enable
	//
	// - **false**: Disable
	//
	// example:
	//
	// false
	AutoStartBackup *bool `json:"AutoStartBackup,omitempty" xml:"AutoStartBackup,omitempty"`
	// The backup gateway ID. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) API to get this parameter\\"s value.
	//
	// > This parameter is required when **SourceEndpointInstanceType*	- is **agent**.
	//
	// example:
	//
	// 23313123312
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The incremental interval in seconds (s).
	//
	// > Only physical backup is supported.
	//
	// example:
	//
	// 1000
	BackupLogIntervalSeconds *int32 `json:"BackupLogIntervalSeconds,omitempty" xml:"BackupLogIntervalSeconds,omitempty"`
	// The backup objects. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) API to get this parameter\\"s value.
	//
	// example:
	//
	// [     {         "DBName":"待备份库名",         "SchemaName":"待备份 Schema 名",         "TableIncludes":[{             "TableName":"待备份表表名"         }],         "TableExcludes":[{             "TableName":"待备份库名不需要备份表的表名"         }]     } ]
	BackupObjects *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	// The full backup period. Valid values:
	//
	// - **Monday**: Monday
	//
	// - **Tuesday**: Tuesday
	//
	// - **Wednesday**: Wednesday
	//
	// - **Thursday**: Thursday
	//
	// - **Friday**: Friday
	//
	// - **Saturday**: Saturday
	//
	// - **Sunday**: Sunday
	//
	// example:
	//
	// Monday
	BackupPeriod *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// The backup plan ID. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) API to get this parameter\\"s value.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi0*******
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The custom backup plan name. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) API to get this parameter\\"s value.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi0*******
	BackupPlanName *string `json:"BackupPlanName,omitempty" xml:"BackupPlanName,omitempty"`
	// The network bandwidth throttling in KB/s. The maximum value is 10 GB.
	//
	// > This parameter is valid only for MySQL physical backup.
	//
	// example:
	//
	// 262144
	BackupRateLimit *int64 `json:"BackupRateLimit,omitempty" xml:"BackupRateLimit,omitempty"`
	// The retention period for backup data. Valid values: 0 to 1825. Default value: 730 days.
	//
	// example:
	//
	// 730
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The disk I/O limit in KB/s.
	//
	// > This parameter is valid only for MySQL physical backup.
	//
	// example:
	//
	// 262144
	BackupSpeedLimit *int64 `json:"BackupSpeedLimit,omitempty" xml:"BackupSpeedLimit,omitempty"`
	// The full backup start time in *HH:mm*Z (UTC) format. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) API to get this parameter\\"s value.
	//
	// example:
	//
	// 14:22
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// example:
	//
	// encrypted
	BackupStorageEncryptMethod *string `json:"BackupStorageEncryptMethod,omitempty" xml:"BackupStorageEncryptMethod,omitempty"`
	// The built-in storage type:
	//
	// - Empty (default): Backup data is stored on your OSS.
	//
	// - system: Backup data is stored on the built-in OSS of DBS.
	//
	// example:
	//
	// 无
	BackupStorageType *string `json:"BackupStorageType,omitempty" xml:"BackupStorageType,omitempty"`
	// The full backup period. Valid values:
	//
	// - **simple**: Periodic backup. Use this value with BackupPeriod and BackupStartTime.
	//
	// - **manual**: Manual backup.
	//
	// > Default value: **simple**.
	//
	// example:
	//
	// simple
	BackupStrategyType *string `json:"BackupStrategyType,omitempty" xml:"BackupStrategyType,omitempty"`
	// Ensure the idempotence of the request to prevent duplicate submissions. The client generates this parameter value. Ensure its uniqueness across different requests. The maximum length is 64 ASCII characters, and the value cannot contain non-ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The UID for cross-Alibaba Cloud account backup. Call the [DescribeRestoreTaskList](https://help.aliyun.com/document_detail/2869838.html) API to get this parameter\\"s value.
	//
	// example:
	//
	// 2xxx7778xxxxxxxxxx
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// The RAM role name for cross-Alibaba Cloud account backup. Call the [DescribeRestoreTaskList](https://help.aliyun.com/document_detail/2869838.html) API to get this parameter\\"s value.
	//
	// example:
	//
	// test123
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	// The period after which data is converted to archive cold storage. Default value: 365 days.
	//
	// example:
	//
	// 365
	DuplicationArchivePeriod *int32 `json:"DuplicationArchivePeriod,omitempty" xml:"DuplicationArchivePeriod,omitempty"`
	// The period after which data is converted to Infrequent Access storage. Default value: 180 days.
	//
	// example:
	//
	// 180
	DuplicationInfrequentAccessPeriod *int32 `json:"DuplicationInfrequentAccessPeriod,omitempty" xml:"DuplicationInfrequentAccessPeriod,omitempty"`
	// Enable incremental log backup. Valid values:
	//
	// - **true**: Enable
	//
	// - **false**: Disable
	//
	// example:
	//
	// true
	EnableBackupLog *bool `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// example:
	//
	// true
	EnableMysqlPhysicalBackupBinlog *string `json:"EnableMysqlPhysicalBackupBinlog,omitempty" xml:"EnableMysqlPhysicalBackupBinlog,omitempty"`
	// example:
	//
	// true
	EnableSourceEndpointSsl *string `json:"EnableSourceEndpointSsl,omitempty" xml:"EnableSourceEndpointSsl,omitempty"`
	// The OSS bucket name.
	//
	// > The system automatically generates a new name by default.
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
	// The database name. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) API to get this parameter\\"s value.
	//
	// > This parameter is required when the database type is **PostgreSQL*	- or **MongoDB**.
	//
	// example:
	//
	// testRDS
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// The database endpoint. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) API to get this parameter\\"s value.
	//
	// > This parameter is required when **SourceEndpointInstanceType*	- is **express**, **agent**, or **other**.
	//
	// example:
	//
	// rm-uf6wjk5*******.mysql.rds.aliyuncs.com
	SourceEndpointIP *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	// The database instance ID. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) API to get this parameter\\"s value.
	//
	// > This parameter is required when **SourceEndpoint**.**InstanceType*	- is **RDS**, **ECS**, **DDS**, or **Express**.
	//
	// example:
	//
	// rm-uf6wjk5*********
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The location of the database. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) API to get this parameter\\"s value. Valid values:
	//
	// - **RDS**
	//
	// - **ECS**
	//
	// - **Express**: A database connected through a leased line, VPN Gateway, or Smart Gateway.
	//
	// - **Agent**: A database connected through a backup gateway.
	//
	// - **DDS**: Cloud MongoDB.
	//
	// - **Other**: A database directly connected through IP:Port.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// example:
	//
	// /home/test
	SourceEndpointOracleHome *string `json:"SourceEndpointOracleHome,omitempty" xml:"SourceEndpointOracleHome,omitempty"`
	// The Oracle SID name.
	//
	// > This parameter is required when the database type is Oracle.
	//
	// example:
	//
	// test
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// The password.
	//
	// > This parameter is optional when the database type is **Redis**, or when the database location is **agent*	- and the database type is **SQL Server**. It is required in other scenarios.
	//
	// example:
	//
	// testPassword
	SourceEndpointPassword *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	// The database port. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) API to get this parameter\\"s value.
	//
	// > This parameter is required when **SourceEndpoint**.**InstanceType*	- is **express**, **agent**, **other**, or **ECS**.
	//
	// example:
	//
	// 3306
	SourceEndpointPort *int32 `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// The region of the database. Call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) API to get this parameter\\"s value.
	//
	// > This parameter is required when **SourceEndpointInstanceType*	- is RDS, ECS, DDS, Express, or Agent.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The database account.
	//
	// > This parameter is optional when the database type is **Redis**, or when the database location is **agent*	- and the database type is **SQL Server**. It is required in other scenarios.
	//
	// example:
	//
	// testRDS
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE----- ... -----END CERTIFICATE-----
	SslCaPem *string `json:"SslCaPem,omitempty" xml:"SslCaPem,omitempty"`
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

func (s *ConfigureBackupPlanRequest) GetBackupStorageEncryptMethod() *string {
	return s.BackupStorageEncryptMethod
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

func (s *ConfigureBackupPlanRequest) GetEnableMysqlPhysicalBackupBinlog() *string {
	return s.EnableMysqlPhysicalBackupBinlog
}

func (s *ConfigureBackupPlanRequest) GetEnableSourceEndpointSsl() *string {
	return s.EnableSourceEndpointSsl
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

func (s *ConfigureBackupPlanRequest) GetSourceEndpointOracleHome() *string {
	return s.SourceEndpointOracleHome
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

func (s *ConfigureBackupPlanRequest) GetSslCaPem() *string {
	return s.SslCaPem
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

func (s *ConfigureBackupPlanRequest) SetBackupStorageEncryptMethod(v string) *ConfigureBackupPlanRequest {
	s.BackupStorageEncryptMethod = &v
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

func (s *ConfigureBackupPlanRequest) SetEnableMysqlPhysicalBackupBinlog(v string) *ConfigureBackupPlanRequest {
	s.EnableMysqlPhysicalBackupBinlog = &v
	return s
}

func (s *ConfigureBackupPlanRequest) SetEnableSourceEndpointSsl(v string) *ConfigureBackupPlanRequest {
	s.EnableSourceEndpointSsl = &v
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

func (s *ConfigureBackupPlanRequest) SetSourceEndpointOracleHome(v string) *ConfigureBackupPlanRequest {
	s.SourceEndpointOracleHome = &v
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

func (s *ConfigureBackupPlanRequest) SetSslCaPem(v string) *ConfigureBackupPlanRequest {
	s.SslCaPem = &v
	return s
}

func (s *ConfigureBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
