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
	// Specifies whether to enable automatic backup. Valid values:
	//
	// - **true**: enables automatic backup
	//
	// - **false**: disables automatic backup
	//
	// example:
	//
	// false
	AutoStartBackup *bool `json:"AutoStartBackup,omitempty" xml:"AutoStartBackup,omitempty"`
	// The ID of the backup gateway. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the value of this parameter.
	//
	// > This parameter is required when **SourceEndpointInstanceType*	- is set to **agent**.
	//
	// example:
	//
	// 23313123312
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The interval for incremental backup, in seconds.
	//
	// > This parameter is supported only for physical backup.
	//
	// example:
	//
	// 1000
	BackupLogIntervalSeconds *int32 `json:"BackupLogIntervalSeconds,omitempty" xml:"BackupLogIntervalSeconds,omitempty"`
	// The objects to be backed up. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the value of this parameter.
	//
	// example:
	//
	// [     {         "DBName":"DatabaseToBackUp",         "SchemaName":"SchemaToBackUp",         "TableIncludes":[{             "TableName":"TableToBackUp"         }],         "TableExcludes":[{             "TableName":"TableToExclude"         }]     } ]
	BackupObjects *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	// The full backup cycle. Valid values:
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
	// The ID of the backup plan. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi0*******
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The custom name of the backup plan. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi0*******
	BackupPlanName *string `json:"BackupPlanName,omitempty" xml:"BackupPlanName,omitempty"`
	// The network bandwidth throttling. Unit: KB/s. Maximum value: 10 GB.
	//
	// > This parameter takes effect only for MySQL physical backup.
	//
	// example:
	//
	// 262144
	BackupRateLimit *int64 `json:"BackupRateLimit,omitempty" xml:"BackupRateLimit,omitempty"`
	// The retention period of backup data. Valid values: 0 to 1825. Default value: 730. Unit: days.
	//
	// example:
	//
	// 730
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The disk I/O limit. Unit: KB/s.
	//
	// > This parameter takes effect only for MySQL physical backup.
	//
	// example:
	//
	// 262144
	BackupSpeedLimit *int64 `json:"BackupSpeedLimit,omitempty" xml:"BackupSpeedLimit,omitempty"`
	// The start time of the full backup. The time is in the <i>HH:mm</i>Z format and displayed in UTC. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the value of this parameter.
	//
	// example:
	//
	// 14:22
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The encryption method for backup storage.
	//
	// example:
	//
	// encrypted
	BackupStorageEncryptMethod *string `json:"BackupStorageEncryptMethod,omitempty" xml:"BackupStorageEncryptMethod,omitempty"`
	// The built-in storage type. Valid values:
	//
	// - Empty string (default): Backup data is stored in your OSS bucket.
	//
	// - system: Backup data is stored in the built-in OSS bucket of DBS.
	//
	// example:
	//
	// None
	BackupStorageType *string `json:"BackupStorageType,omitempty" xml:"BackupStorageType,omitempty"`
	// The full backup epoch. Valid values:
	//
	// - **simple**: periodic backup, which is used together with BackupPeriod and BackupStartTime
	//
	// - **manual**: manual backup
	//
	// > Default value: **simple**.
	//
	// example:
	//
	// simple
	BackupStrategyType *string `json:"BackupStrategyType,omitempty" xml:"BackupStrategyType,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can use the client to generate the value, and you must ensure that the value is unique among different requests. The value can be up to 64 ASCII characters in length and cannot contain non-ASCII characters.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The UID of the Alibaba Cloud account for cross-account backup. You can call the [DescribeRestoreTaskList](https://help.aliyun.com/document_detail/2869838.html) operation to query the value of this parameter.
	//
	// example:
	//
	// 2xxx7778xxxxxxxxxx
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// The name of the RAM role for cross-account backup. You can call the [DescribeRestoreTaskList](https://help.aliyun.com/document_detail/2869838.html) operation to query the value of this parameter.
	//
	// example:
	//
	// test123
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	// The period after which backup data is transitioned to Archive storage. Default value: 365. Unit: days.
	//
	// example:
	//
	// 365
	DuplicationArchivePeriod *int32 `json:"DuplicationArchivePeriod,omitempty" xml:"DuplicationArchivePeriod,omitempty"`
	// The period after which backup data is transitioned to Infrequent Access (IA) storage. Default value: 180. Unit: days.
	//
	// example:
	//
	// 180
	DuplicationInfrequentAccessPeriod *int32 `json:"DuplicationInfrequentAccessPeriod,omitempty" xml:"DuplicationInfrequentAccessPeriod,omitempty"`
	// Specifies whether to enable incremental log backup. Valid values:
	//
	// - **true**: enables incremental log backup
	//
	// - **false**: disables incremental log backup
	//
	// example:
	//
	// true
	EnableBackupLog *bool `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
	// Specifies whether to enable log backup for MySQL physical backup.
	//
	// example:
	//
	// true
	EnableMysqlPhysicalBackupBinlog *string `json:"EnableMysqlPhysicalBackupBinlog,omitempty" xml:"EnableMysqlPhysicalBackupBinlog,omitempty"`
	// Specifies whether to enable an SSL-encrypted connection to the source database instance.
	//
	// example:
	//
	// true
	EnableSourceEndpointSsl *string `json:"EnableSourceEndpointSsl,omitempty" xml:"EnableSourceEndpointSsl,omitempty"`
	// The name of the OSS bucket.
	//
	// > By default, the system automatically generates a new name.
	//
	// example:
	//
	// TestOssBucket
	OSSBucketName *string `json:"OSSBucketName,omitempty" xml:"OSSBucketName,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzecovzti****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the database. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the value of this parameter.
	//
	// > This parameter is required when the database engine is **PostgreSQL*	- or **MongoDB**.
	//
	// example:
	//
	// testRDS
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// The database connection address. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the value of this parameter.
	//
	// > This parameter is required when **SourceEndpointInstanceType*	- is set to **express**, **agent**, or **other**.
	//
	// example:
	//
	// rm-uf6wjk5*******.mysql.rds.aliyuncs.com
	SourceEndpointIP *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	// The database instance ID. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the value of this parameter.
	//
	// > This parameter is required when **SourceEndpoint**.**InstanceType*	- is set to **RDS**, **ECS**, **DDS**, or **Express**.
	//
	// example:
	//
	// rm-uf6wjk5*********
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The location of the database. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the value of this parameter. Valid values:
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
	// This parameter is required.
	//
	// example:
	//
	// RDS
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// Oracle Home
	//
	// example:
	//
	// /home/test
	SourceEndpointOracleHome *string `json:"SourceEndpointOracleHome,omitempty" xml:"SourceEndpointOracleHome,omitempty"`
	// The Oracle system identifier (SID).
	//
	// > This parameter is required when the database engine is Oracle.
	//
	// example:
	//
	// test
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// The password of the database account.
	//
	// > This parameter is not required when the database engine is **Redis**, or when the database is connected over a backup gateway (**agent**) and the database engine is **SQL Server**. This parameter is required in other scenarios.
	//
	// example:
	//
	// testPassword
	SourceEndpointPassword *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	// The port number of the database. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the value of this parameter.
	//
	// > This parameter is required when **SourceEndpoint**.**InstanceType*	- is set to **express**, **agent**, **other**, or **ECS**.
	//
	// example:
	//
	// 3306
	SourceEndpointPort *int32 `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// The region of the database. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to query the value of this parameter.
	//
	// > This parameter is required when **SourceEndpointInstanceType*	- is set to RDS, ECS, DDS, Express, or Agent.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The username of the database account.
	//
	// > This parameter is not required when the database engine is **Redis**, or when the database is connected over a backup gateway (**agent**) and the database engine is **SQL Server**. This parameter is required in other scenarios.
	//
	// example:
	//
	// testRDS
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	// The PEM certificate for the SSL-encrypted connection to the source database.
	//
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
