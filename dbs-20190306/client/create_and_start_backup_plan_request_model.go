// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndStartBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupGatewayId(v int64) *CreateAndStartBackupPlanRequest
	GetBackupGatewayId() *int64
	SetBackupLogIntervalSeconds(v int32) *CreateAndStartBackupPlanRequest
	GetBackupLogIntervalSeconds() *int32
	SetBackupMethod(v string) *CreateAndStartBackupPlanRequest
	GetBackupMethod() *string
	SetBackupObjects(v string) *CreateAndStartBackupPlanRequest
	GetBackupObjects() *string
	SetBackupPeriod(v string) *CreateAndStartBackupPlanRequest
	GetBackupPeriod() *string
	SetBackupPlanId(v string) *CreateAndStartBackupPlanRequest
	GetBackupPlanId() *string
	SetBackupPlanName(v string) *CreateAndStartBackupPlanRequest
	GetBackupPlanName() *string
	SetBackupRateLimit(v int64) *CreateAndStartBackupPlanRequest
	GetBackupRateLimit() *int64
	SetBackupRetentionPeriod(v int32) *CreateAndStartBackupPlanRequest
	GetBackupRetentionPeriod() *int32
	SetBackupSpeedLimit(v int64) *CreateAndStartBackupPlanRequest
	GetBackupSpeedLimit() *int64
	SetBackupStartTime(v string) *CreateAndStartBackupPlanRequest
	GetBackupStartTime() *string
	SetBackupStorageType(v string) *CreateAndStartBackupPlanRequest
	GetBackupStorageType() *string
	SetBackupStrategyType(v string) *CreateAndStartBackupPlanRequest
	GetBackupStrategyType() *string
	SetClientToken(v string) *CreateAndStartBackupPlanRequest
	GetClientToken() *string
	SetCrossAliyunId(v string) *CreateAndStartBackupPlanRequest
	GetCrossAliyunId() *string
	SetCrossRoleName(v string) *CreateAndStartBackupPlanRequest
	GetCrossRoleName() *string
	SetDatabaseRegion(v string) *CreateAndStartBackupPlanRequest
	GetDatabaseRegion() *string
	SetDatabaseType(v string) *CreateAndStartBackupPlanRequest
	GetDatabaseType() *string
	SetDuplicationArchivePeriod(v int32) *CreateAndStartBackupPlanRequest
	GetDuplicationArchivePeriod() *int32
	SetDuplicationInfrequentAccessPeriod(v int32) *CreateAndStartBackupPlanRequest
	GetDuplicationInfrequentAccessPeriod() *int32
	SetEnableBackupLog(v bool) *CreateAndStartBackupPlanRequest
	GetEnableBackupLog() *bool
	SetFromApp(v string) *CreateAndStartBackupPlanRequest
	GetFromApp() *string
	SetInstanceClass(v string) *CreateAndStartBackupPlanRequest
	GetInstanceClass() *string
	SetInstanceType(v string) *CreateAndStartBackupPlanRequest
	GetInstanceType() *string
	SetOSSBucketName(v string) *CreateAndStartBackupPlanRequest
	GetOSSBucketName() *string
	SetOwnerId(v string) *CreateAndStartBackupPlanRequest
	GetOwnerId() *string
	SetPayType(v string) *CreateAndStartBackupPlanRequest
	GetPayType() *string
	SetPeriod(v string) *CreateAndStartBackupPlanRequest
	GetPeriod() *string
	SetRegion(v string) *CreateAndStartBackupPlanRequest
	GetRegion() *string
	SetResourceGroupId(v string) *CreateAndStartBackupPlanRequest
	GetResourceGroupId() *string
	SetSourceEndpointDatabaseName(v string) *CreateAndStartBackupPlanRequest
	GetSourceEndpointDatabaseName() *string
	SetSourceEndpointIP(v string) *CreateAndStartBackupPlanRequest
	GetSourceEndpointIP() *string
	SetSourceEndpointInstanceID(v string) *CreateAndStartBackupPlanRequest
	GetSourceEndpointInstanceID() *string
	SetSourceEndpointInstanceType(v string) *CreateAndStartBackupPlanRequest
	GetSourceEndpointInstanceType() *string
	SetSourceEndpointOracleSID(v string) *CreateAndStartBackupPlanRequest
	GetSourceEndpointOracleSID() *string
	SetSourceEndpointPassword(v string) *CreateAndStartBackupPlanRequest
	GetSourceEndpointPassword() *string
	SetSourceEndpointPort(v int32) *CreateAndStartBackupPlanRequest
	GetSourceEndpointPort() *int32
	SetSourceEndpointRegion(v string) *CreateAndStartBackupPlanRequest
	GetSourceEndpointRegion() *string
	SetSourceEndpointUserName(v string) *CreateAndStartBackupPlanRequest
	GetSourceEndpointUserName() *string
	SetStorageRegion(v string) *CreateAndStartBackupPlanRequest
	GetStorageRegion() *string
	SetStorageType(v string) *CreateAndStartBackupPlanRequest
	GetStorageType() *string
	SetUsedTime(v int32) *CreateAndStartBackupPlanRequest
	GetUsedTime() *int32
}

type CreateAndStartBackupPlanRequest struct {
	// The backup gateway ID.
	//
	// >
	//
	// 	- If **SourceEndpointInstanceType*	- is set to **Agent**, this parameter is required.****
	//
	// 	- For more information about how to install a backup gateway, see [Install a backup gateway](https://help.aliyun.com/document_detail/93250.html).
	//
	// 	- You can query a list of existing backup gateways by calling the [DescribeBackupGatewayList](https://help.aliyun.com/document_detail/2869840.html) operation.
	//
	// example:
	//
	// 23313123312
	BackupGatewayId *int64 `json:"BackupGatewayId,omitempty" xml:"BackupGatewayId,omitempty"`
	// The interval at which you want to perform incremental log backups. Unit: seconds.
	//
	// >  This parameter is required only if you set BackupMethod to **physical**.
	//
	// example:
	//
	// 1000
	BackupLogIntervalSeconds *int32 `json:"BackupLogIntervalSeconds,omitempty" xml:"BackupLogIntervalSeconds,omitempty"`
	// The method that is used to generate the backup file. Valid values:
	//
	// 	- **logical**: logical backup
	//
	// 	- **physical**: physical backup
	//
	// 	- **duplication**: dump backup
	//
	// This parameter is required.
	//
	// example:
	//
	// logical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The object to be backed up.
	//
	// example:
	//
	// [ { "DBName":"Name of the database to be backed up", "SchemaName":"Name of the schema to be backed up", "TableIncludes":[{ "TableName":"Name of the table to be backed up" }], "TableExcludes":[{ "TableName":"Name of the table that you do not want to back up" }] } ]
	BackupObjects *string `json:"BackupObjects,omitempty" xml:"BackupObjects,omitempty"`
	// The day of the week on which you want to perform full backup. Valid values:
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
	// >  You can specify multiple values. Separate multiple values with commas (,).
	//
	// example:
	//
	// Monday
	BackupPeriod *string `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// The ID of the backup schedule.
	//
	// example:
	//
	// dbstooi0*******
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The name of the backup schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbstooi0*******
	BackupPlanName *string `json:"BackupPlanName,omitempty" xml:"BackupPlanName,omitempty"`
	// The network bandwidth throttling. Unit: KB/s. DBS allows a maximum bandwidth of 10 GB/s.
	//
	// >  This parameter takes effect only when physical backups for MySQL databases are performed.
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
	// The I/O limit for the disk. Unit: KB/s.
	//
	// >  This parameter takes effect only when physical backups for MySQL databases are performed.
	//
	// example:
	//
	// 262144
	BackupSpeedLimit *int64 `json:"BackupSpeedLimit,omitempty" xml:"BackupSpeedLimit,omitempty"`
	// The start time of full backup tasks. Specify the value in the *HH:mm	- format. The time must be in UTC.
	//
	// example:
	//
	// 14:22
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The storage type. Valid values:
	//
	// 	- Empty: If you do not specify this parameter, the system stores backup data in your OSS bucket.
	//
	// 	- system : The system stores backup data in the built-in OSS bucket of DBS.
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
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// DBS
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The unique ID (UID) of the Alibaba Cloud account to which the source database belongs.
	//
	// example:
	//
	// 2xxx7778xxxxxxxxxx
	CrossAliyunId *string `json:"CrossAliyunId,omitempty" xml:"CrossAliyunId,omitempty"`
	// The name of the RAM role that is used to perform backups across Alibaba Cloud accounts.
	//
	// example:
	//
	// test123
	CrossRoleName *string `json:"CrossRoleName,omitempty" xml:"CrossRoleName,omitempty"`
	// The region in which the database that you want to back up resides.
	//
	// > This parameter is required if the **PayType*	- parameter is set to **postpay**.
	//
	// example:
	//
	// cn-hangzhou
	DatabaseRegion *string `json:"DatabaseRegion,omitempty" xml:"DatabaseRegion,omitempty"`
	// The type of the source database. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **MSSQL**
	//
	// 	- **Oracle**
	//
	// 	- **MariaDB**
	//
	// 	- **PostgreSQL**
	//
	// 	- **DRDS**
	//
	// 	- **MongoDB**
	//
	// 	- **Redis**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
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
	// The request source. Default value: OpenApi. You do not need to set this parameter.
	//
	// example:
	//
	// OpenApi
	FromApp *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	// The type of the backup schedule. Valid values:
	//
	// 	- **micro**
	//
	// 	- **small**
	//
	// 	- **medium**
	//
	// 	- **large**
	//
	// 	- **xlarge**
	//
	// >  A backup schedule type with higher specifications offers higher backup and restoration performance. For more information, see [Select a backup schedule type](https://help.aliyun.com/document_detail/84372.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// micro
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The type of the source database instance. Valid values:
	//
	// 	- **RDS**: ApsaraDB RDS.
	//
	// 	- **PolarDB**: PolarDB.
	//
	// 	- **DDS**: ApsaraDB for MongoDB.
	//
	// 	- **Kvstore**: ApsaraDB for Redis.
	//
	// 	- **Other**: Database connected by using an IP address and a port number.
	//
	// 	- **dg**: Self-managed database that has no public IP address or port number and is connected over Database Gateway.
	//
	// >  If **PayType*	- is set to **postpay**, this parameter is required.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The name of the Object Storage Service (OSS) bucket used to store backup files. By default, the system automatically generates a name for the OSS bucket.
	//
	// example:
	//
	// TestOssBucket
	OSSBucketName *string `json:"OSSBucketName,omitempty" xml:"OSSBucketName,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **postpay**: pay-as-you-go.
	//
	// 	- **prepay**: subscription.
	//
	// >  The default value is **prepay**. You can set this parameter to **postpay*	- only if you set **BackupMethod*	- to **duplication**.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Specifies whether to use yearly subscription or monthly subscription for the instance. Valid values:
	//
	// 	- **Year**: yearly subscription
	//
	// 	- **Month**: monthly subscription
	//
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the region in which you want to store the backup data. You can query the supported regions of DBS by calling the [DescribeRegions](https://help.aliyun.com/document_detail/2869853.html) operation.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzecovzti****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the database.
	//
	// > This parameter is required if the DatabaseType parameter is set to **PostgreSQL*	- or **MongoDB**.
	//
	// example:
	//
	// testRDS
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// The endpoint of the database.
	//
	// > This parameter is required if the **SourceEndpointInstanceType*	- parameter is set to **Express**, **Agent**, or **Other**.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx.mysql.rds.aliyuncs.com
	SourceEndpointIP *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	// The ID of the database instance.
	//
	// > This parameter is required if the **SourceEndpointInstanceType*	- parameter is set to **RDS**, **ECS**, **DDS**, or **Express**.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The location of the source database. Valid values:
	//
	// 	- **RDS**: The database is on an ApsaraDB RDS instance.
	//
	// 	- **ECS**: The database is on an Elastic Compute Service (ECS) instance.
	//
	// 	- **Express**: The database is connected to DBS by using Express Connect, VPN Gateway, or Smart Access Gateway.
	//
	// 	- **Agent**: The database is connected to DBS over a DBS backup gateway.
	//
	// 	- **DDS**: The database is on an ApsaraDB for MongoDB instance.
	//
	// 	- **Other**: The database is connected to DBS by using an IP address and a port number.
	//
	// 	- **dg**: The database is a self-managed database that has no public IP address or port number and is connected to DBS over Database Gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// The system ID (SID) of the Oracle database. This parameter is required if the source database is an Oracle database.
	//
	// example:
	//
	// test
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// The password of the account that is used to connect to the database.
	//
	// > This parameter is required except that the database is an SQL Server database that is connected to DBS over a DBS backup gateway or a Redis database.
	//
	// example:
	//
	// testPassword
	SourceEndpointPassword *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	// The port of the database.
	//
	// > This parameter is required if the **SourceEndpointInstanceType*	- parameter is set to **Express**, **Agent**, **Other**, or **ECS**.
	//
	// example:
	//
	// 3306
	SourceEndpointPort *int32 `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// The region in which the database that you want to back up resides.
	//
	// > This parameter is required if the **SourceEndpointInstanceType*	- parameter is set to **RDS**, **ECS**, **DDS**, **Express**, or **Agent**.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The username of the account that is used to connect to the database.
	//
	// > This parameter is required except that the database is an SQL Server database that is connected to DBS over a DBS backup gateway or a Redis database.
	//
	// example:
	//
	// testRDS
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	// The region in which you want to store the backup data.
	//
	// > This parameter is required if the **PayType*	- parameter is set to **postpay**.
	//
	// example:
	//
	// cn-hangzhou
	StorageRegion *string `json:"StorageRegion,omitempty" xml:"StorageRegion,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// N/A
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The subscription duration. Valid values:
	//
	// 	- If **Period*	- is set to **Year**, the valid values of **UsedTime*	- range from 1 to 5.
	//
	// 	- If **Period*	- is set to **Month**, the valid values of **UsedTime*	- range from 1 to 11.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s CreateAndStartBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAndStartBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateAndStartBackupPlanRequest) GetBackupGatewayId() *int64 {
	return s.BackupGatewayId
}

func (s *CreateAndStartBackupPlanRequest) GetBackupLogIntervalSeconds() *int32 {
	return s.BackupLogIntervalSeconds
}

func (s *CreateAndStartBackupPlanRequest) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *CreateAndStartBackupPlanRequest) GetBackupObjects() *string {
	return s.BackupObjects
}

func (s *CreateAndStartBackupPlanRequest) GetBackupPeriod() *string {
	return s.BackupPeriod
}

func (s *CreateAndStartBackupPlanRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *CreateAndStartBackupPlanRequest) GetBackupPlanName() *string {
	return s.BackupPlanName
}

func (s *CreateAndStartBackupPlanRequest) GetBackupRateLimit() *int64 {
	return s.BackupRateLimit
}

func (s *CreateAndStartBackupPlanRequest) GetBackupRetentionPeriod() *int32 {
	return s.BackupRetentionPeriod
}

func (s *CreateAndStartBackupPlanRequest) GetBackupSpeedLimit() *int64 {
	return s.BackupSpeedLimit
}

func (s *CreateAndStartBackupPlanRequest) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *CreateAndStartBackupPlanRequest) GetBackupStorageType() *string {
	return s.BackupStorageType
}

func (s *CreateAndStartBackupPlanRequest) GetBackupStrategyType() *string {
	return s.BackupStrategyType
}

func (s *CreateAndStartBackupPlanRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAndStartBackupPlanRequest) GetCrossAliyunId() *string {
	return s.CrossAliyunId
}

func (s *CreateAndStartBackupPlanRequest) GetCrossRoleName() *string {
	return s.CrossRoleName
}

func (s *CreateAndStartBackupPlanRequest) GetDatabaseRegion() *string {
	return s.DatabaseRegion
}

func (s *CreateAndStartBackupPlanRequest) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *CreateAndStartBackupPlanRequest) GetDuplicationArchivePeriod() *int32 {
	return s.DuplicationArchivePeriod
}

func (s *CreateAndStartBackupPlanRequest) GetDuplicationInfrequentAccessPeriod() *int32 {
	return s.DuplicationInfrequentAccessPeriod
}

func (s *CreateAndStartBackupPlanRequest) GetEnableBackupLog() *bool {
	return s.EnableBackupLog
}

func (s *CreateAndStartBackupPlanRequest) GetFromApp() *string {
	return s.FromApp
}

func (s *CreateAndStartBackupPlanRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *CreateAndStartBackupPlanRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateAndStartBackupPlanRequest) GetOSSBucketName() *string {
	return s.OSSBucketName
}

func (s *CreateAndStartBackupPlanRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateAndStartBackupPlanRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateAndStartBackupPlanRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateAndStartBackupPlanRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateAndStartBackupPlanRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAndStartBackupPlanRequest) GetSourceEndpointDatabaseName() *string {
	return s.SourceEndpointDatabaseName
}

func (s *CreateAndStartBackupPlanRequest) GetSourceEndpointIP() *string {
	return s.SourceEndpointIP
}

func (s *CreateAndStartBackupPlanRequest) GetSourceEndpointInstanceID() *string {
	return s.SourceEndpointInstanceID
}

func (s *CreateAndStartBackupPlanRequest) GetSourceEndpointInstanceType() *string {
	return s.SourceEndpointInstanceType
}

func (s *CreateAndStartBackupPlanRequest) GetSourceEndpointOracleSID() *string {
	return s.SourceEndpointOracleSID
}

func (s *CreateAndStartBackupPlanRequest) GetSourceEndpointPassword() *string {
	return s.SourceEndpointPassword
}

func (s *CreateAndStartBackupPlanRequest) GetSourceEndpointPort() *int32 {
	return s.SourceEndpointPort
}

func (s *CreateAndStartBackupPlanRequest) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *CreateAndStartBackupPlanRequest) GetSourceEndpointUserName() *string {
	return s.SourceEndpointUserName
}

func (s *CreateAndStartBackupPlanRequest) GetStorageRegion() *string {
	return s.StorageRegion
}

func (s *CreateAndStartBackupPlanRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateAndStartBackupPlanRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateAndStartBackupPlanRequest) SetBackupGatewayId(v int64) *CreateAndStartBackupPlanRequest {
	s.BackupGatewayId = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetBackupLogIntervalSeconds(v int32) *CreateAndStartBackupPlanRequest {
	s.BackupLogIntervalSeconds = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetBackupMethod(v string) *CreateAndStartBackupPlanRequest {
	s.BackupMethod = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetBackupObjects(v string) *CreateAndStartBackupPlanRequest {
	s.BackupObjects = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetBackupPeriod(v string) *CreateAndStartBackupPlanRequest {
	s.BackupPeriod = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetBackupPlanId(v string) *CreateAndStartBackupPlanRequest {
	s.BackupPlanId = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetBackupPlanName(v string) *CreateAndStartBackupPlanRequest {
	s.BackupPlanName = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetBackupRateLimit(v int64) *CreateAndStartBackupPlanRequest {
	s.BackupRateLimit = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetBackupRetentionPeriod(v int32) *CreateAndStartBackupPlanRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetBackupSpeedLimit(v int64) *CreateAndStartBackupPlanRequest {
	s.BackupSpeedLimit = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetBackupStartTime(v string) *CreateAndStartBackupPlanRequest {
	s.BackupStartTime = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetBackupStorageType(v string) *CreateAndStartBackupPlanRequest {
	s.BackupStorageType = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetBackupStrategyType(v string) *CreateAndStartBackupPlanRequest {
	s.BackupStrategyType = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetClientToken(v string) *CreateAndStartBackupPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetCrossAliyunId(v string) *CreateAndStartBackupPlanRequest {
	s.CrossAliyunId = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetCrossRoleName(v string) *CreateAndStartBackupPlanRequest {
	s.CrossRoleName = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetDatabaseRegion(v string) *CreateAndStartBackupPlanRequest {
	s.DatabaseRegion = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetDatabaseType(v string) *CreateAndStartBackupPlanRequest {
	s.DatabaseType = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetDuplicationArchivePeriod(v int32) *CreateAndStartBackupPlanRequest {
	s.DuplicationArchivePeriod = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetDuplicationInfrequentAccessPeriod(v int32) *CreateAndStartBackupPlanRequest {
	s.DuplicationInfrequentAccessPeriod = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetEnableBackupLog(v bool) *CreateAndStartBackupPlanRequest {
	s.EnableBackupLog = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetFromApp(v string) *CreateAndStartBackupPlanRequest {
	s.FromApp = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetInstanceClass(v string) *CreateAndStartBackupPlanRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetInstanceType(v string) *CreateAndStartBackupPlanRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetOSSBucketName(v string) *CreateAndStartBackupPlanRequest {
	s.OSSBucketName = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetOwnerId(v string) *CreateAndStartBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetPayType(v string) *CreateAndStartBackupPlanRequest {
	s.PayType = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetPeriod(v string) *CreateAndStartBackupPlanRequest {
	s.Period = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetRegion(v string) *CreateAndStartBackupPlanRequest {
	s.Region = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetResourceGroupId(v string) *CreateAndStartBackupPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetSourceEndpointDatabaseName(v string) *CreateAndStartBackupPlanRequest {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetSourceEndpointIP(v string) *CreateAndStartBackupPlanRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetSourceEndpointInstanceID(v string) *CreateAndStartBackupPlanRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetSourceEndpointInstanceType(v string) *CreateAndStartBackupPlanRequest {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetSourceEndpointOracleSID(v string) *CreateAndStartBackupPlanRequest {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetSourceEndpointPassword(v string) *CreateAndStartBackupPlanRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetSourceEndpointPort(v int32) *CreateAndStartBackupPlanRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetSourceEndpointRegion(v string) *CreateAndStartBackupPlanRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetSourceEndpointUserName(v string) *CreateAndStartBackupPlanRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetStorageRegion(v string) *CreateAndStartBackupPlanRequest {
	s.StorageRegion = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetStorageType(v string) *CreateAndStartBackupPlanRequest {
	s.StorageType = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) SetUsedTime(v int32) *CreateAndStartBackupPlanRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateAndStartBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
