// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iConfigureDtsJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckpoint(v string) *ConfigureDtsJobAdvanceRequest
	GetCheckpoint() *string
	SetDataCheckConfigure(v string) *ConfigureDtsJobAdvanceRequest
	GetDataCheckConfigure() *string
	SetDataInitialization(v bool) *ConfigureDtsJobAdvanceRequest
	GetDataInitialization() *bool
	SetDataSynchronization(v bool) *ConfigureDtsJobAdvanceRequest
	GetDataSynchronization() *bool
	SetDbList(v string) *ConfigureDtsJobAdvanceRequest
	GetDbList() *string
	SetDedicatedClusterId(v string) *ConfigureDtsJobAdvanceRequest
	GetDedicatedClusterId() *string
	SetDelayNotice(v bool) *ConfigureDtsJobAdvanceRequest
	GetDelayNotice() *bool
	SetDelayPhone(v string) *ConfigureDtsJobAdvanceRequest
	GetDelayPhone() *string
	SetDelayRuleTime(v int64) *ConfigureDtsJobAdvanceRequest
	GetDelayRuleTime() *int64
	SetDestCaCertificateOssUrl(v string) *ConfigureDtsJobAdvanceRequest
	GetDestCaCertificateOssUrl() *string
	SetDestCaCertificatePassword(v string) *ConfigureDtsJobAdvanceRequest
	GetDestCaCertificatePassword() *string
	SetDestClientCertOssUrl(v string) *ConfigureDtsJobAdvanceRequest
	GetDestClientCertOssUrl() *string
	SetDestClientKeyOssUrl(v string) *ConfigureDtsJobAdvanceRequest
	GetDestClientKeyOssUrl() *string
	SetDestClientPassword(v string) *ConfigureDtsJobAdvanceRequest
	GetDestClientPassword() *string
	SetDestPrimaryVswId(v string) *ConfigureDtsJobAdvanceRequest
	GetDestPrimaryVswId() *string
	SetDestSecondaryVswId(v string) *ConfigureDtsJobAdvanceRequest
	GetDestSecondaryVswId() *string
	SetDestinationEndpointDataBaseName(v string) *ConfigureDtsJobAdvanceRequest
	GetDestinationEndpointDataBaseName() *string
	SetDestinationEndpointEngineName(v string) *ConfigureDtsJobAdvanceRequest
	GetDestinationEndpointEngineName() *string
	SetDestinationEndpointIP(v string) *ConfigureDtsJobAdvanceRequest
	GetDestinationEndpointIP() *string
	SetDestinationEndpointInstanceID(v string) *ConfigureDtsJobAdvanceRequest
	GetDestinationEndpointInstanceID() *string
	SetDestinationEndpointInstanceType(v string) *ConfigureDtsJobAdvanceRequest
	GetDestinationEndpointInstanceType() *string
	SetDestinationEndpointOracleSID(v string) *ConfigureDtsJobAdvanceRequest
	GetDestinationEndpointOracleSID() *string
	SetDestinationEndpointOwnerID(v string) *ConfigureDtsJobAdvanceRequest
	GetDestinationEndpointOwnerID() *string
	SetDestinationEndpointPassword(v string) *ConfigureDtsJobAdvanceRequest
	GetDestinationEndpointPassword() *string
	SetDestinationEndpointPort(v string) *ConfigureDtsJobAdvanceRequest
	GetDestinationEndpointPort() *string
	SetDestinationEndpointRegion(v string) *ConfigureDtsJobAdvanceRequest
	GetDestinationEndpointRegion() *string
	SetDestinationEndpointRole(v string) *ConfigureDtsJobAdvanceRequest
	GetDestinationEndpointRole() *string
	SetDestinationEndpointUserName(v string) *ConfigureDtsJobAdvanceRequest
	GetDestinationEndpointUserName() *string
	SetDisasterRecoveryJob(v bool) *ConfigureDtsJobAdvanceRequest
	GetDisasterRecoveryJob() *bool
	SetDtsBisLabel(v string) *ConfigureDtsJobAdvanceRequest
	GetDtsBisLabel() *string
	SetDtsInstanceId(v string) *ConfigureDtsJobAdvanceRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ConfigureDtsJobAdvanceRequest
	GetDtsJobId() *string
	SetDtsJobName(v string) *ConfigureDtsJobAdvanceRequest
	GetDtsJobName() *string
	SetErrorNotice(v bool) *ConfigureDtsJobAdvanceRequest
	GetErrorNotice() *bool
	SetErrorPhone(v string) *ConfigureDtsJobAdvanceRequest
	GetErrorPhone() *string
	SetFileOssUrlObject(v io.Reader) *ConfigureDtsJobAdvanceRequest
	GetFileOssUrlObject() io.Reader
	SetJobType(v string) *ConfigureDtsJobAdvanceRequest
	GetJobType() *string
	SetMaxDu(v float64) *ConfigureDtsJobAdvanceRequest
	GetMaxDu() *float64
	SetMinDu(v float64) *ConfigureDtsJobAdvanceRequest
	GetMinDu() *float64
	SetOwnerId(v string) *ConfigureDtsJobAdvanceRequest
	GetOwnerId() *string
	SetRegionId(v string) *ConfigureDtsJobAdvanceRequest
	GetRegionId() *string
	SetReserve(v string) *ConfigureDtsJobAdvanceRequest
	GetReserve() *string
	SetResourceGroupId(v string) *ConfigureDtsJobAdvanceRequest
	GetResourceGroupId() *string
	SetSourceEndpointDatabaseName(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointDatabaseName() *string
	SetSourceEndpointEngineName(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointEngineName() *string
	SetSourceEndpointIP(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointIP() *string
	SetSourceEndpointInstanceID(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointInstanceID() *string
	SetSourceEndpointInstanceType(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointInstanceType() *string
	SetSourceEndpointOracleSID(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointOracleSID() *string
	SetSourceEndpointOwnerID(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointOwnerID() *string
	SetSourceEndpointPassword(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointPassword() *string
	SetSourceEndpointPort(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointPort() *string
	SetSourceEndpointRegion(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointRegion() *string
	SetSourceEndpointRole(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointRole() *string
	SetSourceEndpointUserName(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointUserName() *string
	SetSourceEndpointVSwitchID(v string) *ConfigureDtsJobAdvanceRequest
	GetSourceEndpointVSwitchID() *string
	SetSrcCaCertificateOssUrl(v string) *ConfigureDtsJobAdvanceRequest
	GetSrcCaCertificateOssUrl() *string
	SetSrcCaCertificatePassword(v string) *ConfigureDtsJobAdvanceRequest
	GetSrcCaCertificatePassword() *string
	SetSrcClientCertOssUrl(v string) *ConfigureDtsJobAdvanceRequest
	GetSrcClientCertOssUrl() *string
	SetSrcClientKeyOssUrl(v string) *ConfigureDtsJobAdvanceRequest
	GetSrcClientKeyOssUrl() *string
	SetSrcClientPassword(v string) *ConfigureDtsJobAdvanceRequest
	GetSrcClientPassword() *string
	SetSrcPrimaryVswId(v string) *ConfigureDtsJobAdvanceRequest
	GetSrcPrimaryVswId() *string
	SetSrcSecondaryVswId(v string) *ConfigureDtsJobAdvanceRequest
	GetSrcSecondaryVswId() *string
	SetStructureInitialization(v bool) *ConfigureDtsJobAdvanceRequest
	GetStructureInitialization() *bool
	SetSynchronizationDirection(v string) *ConfigureDtsJobAdvanceRequest
	GetSynchronizationDirection() *string
}

type ConfigureDtsJobAdvanceRequest struct {
	// The start offset for incremental data migration or the synchronization checkpoint, in UNIX timestamp format. Unit: seconds.
	//
	// > If you specify the **Checkpoint*	- parameter, you must ensure that no other running DTS instances use the same source database as the destination DTS instance.
	//
	// example:
	//
	// 1610540493
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The parameters for the data validation task, in JSON character string format, such as parameter limits and alert configuration. For more information, see [DataCheckConfigure metric description](https://help.aliyun.com/document_detail/459023.html).
	//
	// example:
	//
	// {"fullCheckModel":1,"fullCheckRatio":20,"checkMaximumHourEnable":1,"checkMaximumHour":1,"fullCheckErrorNotice":true,"fullCheckValidFailNotice":true,"fullCheckNoticeValue":8,"incrementalCheckErrorNotice":true,"incrementalCheckValidFailNotice":true,"incrementalCheckValidFailNoticeTimes":2,"incrementalCheckValidFailNoticePeriod":1,"incrementalCheckValidFailNoticeValue":1,"incrementalCheckDelayNotice":true,"incrementalCheckDelayNoticeTimes":2,"incrementalCheckDelayNoticePeriod":1,"incrementalCheckDelayNoticeValue":60,"fullDataCheck":true,"incrementalDataCheck":true,"dataCheckNoticePhone":"13126800****","dataCheckDbList":{"dts":{"name":"dts","all":true}}}
	DataCheckConfigure *string `json:"DataCheckConfigure,omitempty" xml:"DataCheckConfigure,omitempty"`
	// Specifies whether to perform full data migration or initial full data synchronization. Valid values:
	//
	// - **true**: Yes. This is the default value.
	//
	// - **false**: No.
	//
	// > If **JobType*	- is set to **CHECK**, you can only set this parameter to **false**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Specifies whether to perform incremental data migration or synchronization. Valid values:
	//
	// - **false**: No. This is the default value.
	//
	// - **true**: Yes.
	//
	// > If **JobType*	- is set to **CHECK**, you can only set this parameter to **false**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// The objects to be migrated or synchronized, in JSON format. For more information, see [Object definitions for migration, synchronization, and change tracking](https://help.aliyun.com/document_detail/209545.html).
	//
	// - The DbList parameter supports up to 1 MB of data.
	//
	// - If DbList contains filter conditions, the total length of DbList including the filter conditions can be up to 1 MB.
	//
	// - For distributed tasks such as migration or synchronization tasks where the source is PolarDB-X 1.0, DbList is split based on physical table shards to generate multiple subtasks. The DbList of each subtask supports up to 1 MB of data.
	//
	// example:
	//
	// {"dtstest":{"name":"dtstest","all":true}}
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The ID of the dedicated DTS cluster.
	//
	// > When the ID of a dedicated cluster is specified, the task is scheduled to the corresponding cluster.
	//
	// example:
	//
	// dtscluster_atyl3b5214uk***
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// Specifies whether to monitor the latency status. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	DelayNotice *bool `json:"DelayNotice,omitempty" xml:"DelayNotice,omitempty"`
	// The mobile phone numbers of the contacts to receive latency alerts. Separate multiple mobile phone numbers with commas (,).
	//
	// > - This parameter is supported only on the China site and supports only mobile phone numbers in the Chinese mainland. You can specify up to 10 mobile phone numbers.
	//
	// - The international site does not support mobile phone alerts. You can only [set alert rules for DTS tasks through CloudMonitor](https://help.aliyun.com/document_detail/175876.html).
	//
	// example:
	//
	// 1361234****,1371234****
	DelayPhone *string `json:"DelayPhone,omitempty" xml:"DelayPhone,omitempty"`
	// The threshold for triggering latency alerts. Unit: seconds. The value must be an integer. You can set the threshold based on your business requirements. We recommend that you set the threshold to more than 10 seconds to avoid latency fluctuations caused by network or database loads.
	//
	// > This parameter is required when **DelayNotice*	- is set to **true**.
	//
	// example:
	//
	// 10
	DelayRuleTime *int64 `json:"DelayRuleTime,omitempty" xml:"DelayRuleTime,omitempty"`
	// The path of the CA certificate when the destination database is connected over SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestCaCertificateOssUrl *string `json:"DestCaCertificateOssUrl,omitempty" xml:"DestCaCertificateOssUrl,omitempty"`
	// The password of the CA certificate when the destination database is connected over SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestCaCertificatePassword *string `json:"DestCaCertificatePassword,omitempty" xml:"DestCaCertificatePassword,omitempty"`
	// The path of the client certificate when the destination database is connected over SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestClientCertOssUrl *string `json:"DestClientCertOssUrl,omitempty" xml:"DestClientCertOssUrl,omitempty"`
	// The path of the client certificate private key when the destination database is connected over SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestClientKeyOssUrl *string `json:"DestClientKeyOssUrl,omitempty" xml:"DestClientKeyOssUrl,omitempty"`
	// The password of the client certificate private key when the destination database is connected over SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestClientPassword *string `json:"DestClientPassword,omitempty" xml:"DestClientPassword,omitempty"`
	// The primary vSwitch at the destination end of the VPC NAT gateway.
	//
	// example:
	//
	// ****
	DestPrimaryVswId *string `json:"DestPrimaryVswId,omitempty" xml:"DestPrimaryVswId,omitempty"`
	// The secondary vSwitch at the destination end of the VPC NAT gateway.
	//
	// example:
	//
	// ****
	DestSecondaryVswId *string `json:"DestSecondaryVswId,omitempty" xml:"DestSecondaryVswId,omitempty"`
	// The name of the database to which the migration objects belong in the destination instance.
	//
	// > - This parameter is available and required only when the destination instance or the database engine of the destination instance is PolarDB for PostgreSQL (Compatible with Oracle), AnalyticDB for PostgreSQL, PostgreSQL, MaxCompute, or MongoDB.
	//
	// - If the destination database is MaxCompute, you must specify the name of the MaxCompute project.
	//
	// example:
	//
	// dtstestdata
	DestinationEndpointDataBaseName *string `json:"DestinationEndpointDataBaseName,omitempty" xml:"DestinationEndpointDataBaseName,omitempty"`
	// The database type of the destination instance. Valid values:
	//
	// - **MYSQL**: MySQL database, including ApsaraDB RDS for MySQL and self-managed MySQL.
	//
	// - **MARIADB**: ApsaraDB RDS for MariaDB.
	//
	// - **PolarDB**: PolarDB for MySQL.
	//
	// - **POLARDB_O**: PolarDB for PostgreSQL (Compatible with Oracle).
	//
	// - **POLARDBX10**: PolarDB-X 1.0 (formerly DRDS).
	//
	// - **POLARDBX20**: cloud-native distributed database PolarDB-X 2.0.
	//
	// - **ORACLE**: Self-managed Oracle database.
	//
	// - **PostgreSQL**: PostgreSQL database, including ApsaraDB RDS for PostgreSQL and self-managed PostgreSQL.
	//
	// - **MSSQL**: SQL Server database, including ApsaraDB RDS for SQL Server and self-managed SQL Server.
	//
	// - **ADS**: AnalyticDB for MySQL V2.0.
	//
	// - **ADB30**: AnalyticDB for MySQL V3.0.
	//
	// - **MONGODB**: MongoDB database, including self-managed MongoDB and ApsaraDB for MongoDB.
	//
	// - **ROCKETMQ**: MSMQ RocketMQ.
	//
	// - **GREENPLUM**: AnalyticDB for PostgreSQL.
	//
	// - **KAFKA**: Kafka database, including MSMQ for Kafka and self-managed Kafka.
	//
	// - **DATAHUB**: DataHub.
	//
	// - **DB2**: Self-managed Db2 for LUW database.
	//
	// - **AS400**: Self-managed Db2 for i database.
	//
	// - **ODPS**: MaxCompute.
	//
	// - **Tablestore**: Tablestore.
	//
	// - **ELK**: Alibaba Cloud retrieve and analysis service Elasticsearch.
	//
	// - **REDIS**: Redis database, including self-managed Redis and Tair (Redis OSS-Compatible).
	//
	// - **LINDORM**: cloud-native multi-model database Lindorm.
	//
	// > - Default value: **MYSQL**.
	//
	// - If the database type of the destination instance is set to **KAFKA**, **MONGODB**, or **PolarDB**, you must also specify some information in the reserved parameter Reserve. For more information about the configuration method, see [Reserve metric description](https://help.aliyun.com/document_detail/273111.html).
	//
	// example:
	//
	// MYSQL
	DestinationEndpointEngineName *string `json:"DestinationEndpointEngineName,omitempty" xml:"DestinationEndpointEngineName,omitempty"`
	// The IP address of the destination instance.
	//
	// > This parameter is available and required only when **DestinationEndpointInstanceType*	- is set to **OTHER**, **EXPRESS**, **DG**, or **CEN**.
	//
	// example:
	//
	// ``172.16.**.**``*
	DestinationEndpointIP *string `json:"DestinationEndpointIP,omitempty" xml:"DestinationEndpointIP,omitempty"`
	// The instance ID of the destination instance.
	//
	// If the destination instance is an ApsaraDB database such as ApsaraDB RDS for MySQL, you must specify the instance ID of the ApsaraDB database.
	//
	// If the destination instance is a self-managed database, the value of this parameter varies based on the value of **DestinationEndpointInstanceType**. For example:
	//
	// - If **DestinationEndpointInstanceType*	- is set to **ECS**, specify the instance ID of the ECS instance.
	//
	// - If **DestinationEndpointInstanceType*	- is set to **DG**, specify the instance ID of the database gateway.
	//
	// - If **DestinationEndpointInstanceType*	- is set to **EXPRESS*	- or **CEN**, specify the instance ID of the VPC that is connected to the source database.
	//
	// > If the value is **CEN**, you must also specify the instance ID of the CEN instance in the reserved parameter Reserve. For more information about the configuration method, see [Reserve metric description](https://help.aliyun.com/document_detail/273111.html).
	//
	// example:
	//
	// vpc-bp1opxu1zkhn00gzv****
	DestinationEndpointInstanceID *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	// The target instance type. Valid values:
	//
	// **Alibaba Cloud databases**
	//
	// - **RDS**: ApsaraDB RDS for MySQL, ApsaraDB RDS for SQL Server, ApsaraDB RDS for PostgreSQL, or ApsaraDB RDS for MariaDB.
	//
	// - **PolarDB**: PolarDB for MySQL.
	//
	// - **DISTRIBUTED_POLARDBX10**: PolarDB-X 1.0 (formerly DRDS).
	//
	// - **POLARDBX20**: PolarDB-X 2.0.
	//
	// - **REDIS**: Tair (Redis OSS-Compatible).
	//
	// - **ADS**: AnalyticDB for MySQL V2.0 or V3.0.
	//
	// - **MONGODB**: ApsaraDB for MongoDB.
	//
	// - **ROCKETMQ**: ApsaraDB for RocketMQ.
	//
	// - **GREENPLUM**: AnalyticDB for PostgreSQL.
	//
	// - **DATAHUB**: DataHub, the Alibaba Cloud streaming data processing platform.
	//
	// - **ELK**: Alibaba Cloud retrieve and analysis service Elasticsearch.
	//
	// - **Tablestore**: Tablestore.
	//
	// - **ODPS**: MaxCompute.
	//
	// - **LINDORM**: cloud-native multi-model database Lindorm.
	//
	// **Self-managed databases**
	//
	// - **OTHER**: Self-managed database with a public IP address.
	//
	// - **ECS**: Self-managed database hosted on ECS.
	//
	// - **EXPRESS**: Self-managed database connected over Express Connect.
	//
	// - **CEN**: Self-managed database connected over CEN.
	//
	// - **DG**: Self-managed database connected over Database Gateway.
	//
	// > - If the destination instance is a PolarDB for PostgreSQL (Compatible with Oracle) cluster, you must set this parameter to **OTHER*	- or **EXPRESS*	- to treat it as a self-managed database and connect it over a public IP address or Express Connect.
	//
	// - If the destination instance is MSMQ for Kafka, you must set this parameter to **ECS*	- or **EXPRESS*	- to treat it as a self-managed database and connect it over ECS or Express Connect.
	//
	// - For more information about the supported source and destination databases, see <props="china">[Supported databases](https://help.aliyun.com/document_detail/131497.html)<props="intl">[Supported source and destination databases](https://help.aliyun.com/document_detail/176064.html).
	//
	// - If the destination instance is a self-managed database, you must execute the corresponding preparations. For more information, see [Preparations overview](https://help.aliyun.com/document_detail/146958.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// EXPRESS
	DestinationEndpointInstanceType *string `json:"DestinationEndpointInstanceType,omitempty" xml:"DestinationEndpointInstanceType,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is available and required only when **DestinationEndpointEngineName*	- is set to **Oracle*	- and the **Oracle*	- database is a non-RAC instance.
	//
	// example:
	//
	// testsid
	DestinationEndpointOracleSID *string `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	// The Alibaba Cloud account ID to which the destination ApsaraDB RDS for MySQL instance belongs.
	//
	// > - This parameter can be configured only when the destination instance is ApsaraDB RDS for MySQL.
	//
	// - Specifying this parameter indicates that you execute cross-Alibaba Cloud account data migration or synchronization. You must also specify the **DestinationEndpointRole*	- parameter.
	//
	// example:
	//
	// 140692647406****
	DestinationEndpointOwnerID *string `json:"DestinationEndpointOwnerID,omitempty" xml:"DestinationEndpointOwnerID,omitempty"`
	// The password of the database account of the destination database.
	//
	// > If the destination database is MaxCompute, you must specify the AccessKey secret of the Alibaba Cloud account. For more information about how to obtain an AccessKey secret, see [Create an AccessKey pair](https://help.aliyun.com/document_detail/116401.html).
	//
	// example:
	//
	// Test123456
	DestinationEndpointPassword *string `json:"DestinationEndpointPassword,omitempty" xml:"DestinationEndpointPassword,omitempty"`
	// The database service port of the destination instance.
	//
	// > This parameter is available and required only when the destination instance is a self-managed database.
	//
	// example:
	//
	// 3306
	DestinationEndpointPort *string `json:"DestinationEndpointPort,omitempty" xml:"DestinationEndpointPort,omitempty"`
	// The area of the destination instance. For details, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > This parameter is required if the destination instance is an ApsaraDB database.
	//
	// example:
	//
	// cn-hangzhou
	DestinationEndpointRegion *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	// The name of the role configured for the Alibaba Cloud account to which the destination instance belongs.
	//
	// > This parameter is required when you migrate or synchronize data across Alibaba Cloud accounts. For more information about the permissions required for the role and the authorization method, see [Configure RAM authorization for cross-account data migration or synchronization](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	DestinationEndpointRole *string `json:"DestinationEndpointRole,omitempty" xml:"DestinationEndpointRole,omitempty"`
	// The database account of the destination database.
	//
	// > - In most cases, you must specify the database account of the destination database.
	//
	// - The permissions required for migrating or synchronizing data vary based on the database. For more information about the permission requirements, see [Prepare database accounts for data migration](https://help.aliyun.com/document_detail/175878.html) and [Prepare database accounts for data synchronization](https://help.aliyun.com/document_detail/213152.html).
	//
	// - If the destination database is MaxCompute, you must specify the AccessKey ID of the Alibaba Cloud account. For more information about how to obtain an AccessKey ID, see [Create an AccessKey pair](https://help.aliyun.com/document_detail/116401.html).
	//
	// example:
	//
	// dtstest
	DestinationEndpointUserName *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	// Specifies whether the instance is a disaster recovery instance. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	DisasterRecoveryJob *bool `json:"DisasterRecoveryJob,omitempty" xml:"DisasterRecoveryJob,omitempty"`
	// The environment tag of the DTS instance. Valid values:
	//
	// - **normal**: Normal.
	//
	// - **online**: Online.
	//
	// example:
	//
	// normal
	DtsBisLabel *string `json:"DtsBisLabel,omitempty" xml:"DtsBisLabel,omitempty"`
	// The ID of the data migration or synchronization instance.
	//
	// > You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the instance ID.
	//
	// example:
	//
	// dtsk2gm967v16f****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the data migration or synchronization task.
	//
	// > You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// example:
	//
	// k2gm967v16f****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The name of the DTS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rdsmysql_to_mysql
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// Specifies whether to monitor the exception status. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	ErrorNotice *bool `json:"ErrorNotice,omitempty" xml:"ErrorNotice,omitempty"`
	// The mobile phone numbers of the contacts to receive exception alerts. Separate multiple mobile phone numbers with commas (,).
	//
	// > - This parameter is supported only on the China site and supports only mobile phone numbers in the Chinese mainland. You can specify up to 10 mobile phone numbers.
	//
	// - The international site does not support mobile phone alerts. You can only [set alert rules for DTS tasks through CloudMonitor](https://help.aliyun.com/document_detail/175876.html).
	//
	// example:
	//
	// 1361234****,1371234****
	ErrorPhone *string `json:"ErrorPhone,omitempty" xml:"ErrorPhone,omitempty"`
	// The OSS URL of the task file.
	//
	// example:
	//
	// http://db-list-os-file.oss-cn-shanghai.aliyuncs.com/8e42_121852**********_79dd3aeabe2f43cdb**************
	FileOssUrlObject io.Reader `json:"FileOssUrl,omitempty" xml:"FileOssUrl,omitempty"`
	// The node type. Valid values:
	//
	// - **MIGRATION**: data migration.
	//
	// - **SYNC**: data synchronization.
	//
	// - **CHECK**: data validation (purchased separately).
	//
	// > - If the value is **MIGRATION*	- or **SYNC**, you can also configure a data validation node in the migration or synchronization instance.
	//
	// - To configure a data validation node, you must also configure the **DataCheckConfigure*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The maximum number of DUs.
	//
	// > This parameter is supported only for serverless instances.
	//
	// example:
	//
	// 16
	MaxDu *float64 `json:"MaxDu,omitempty" xml:"MaxDu,omitempty"`
	// The minimum number of DUs.
	//
	// > This parameter is supported only for serverless instances.
	//
	// example:
	//
	// 1
	MinDu   *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	OwnerId *string  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the DTS instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The reserved parameter of DTS, in JSON character string format. You can specify this parameter to supplement the information about the source and destination databases, such as the data storage format of the destination Kafka database, the instance ID of the CEN instance, and the ETL feature configuration. For more information, see [Reserve metric description](https://help.aliyun.com/document_detail/273111.html).
	//
	// example:
	//
	// {      "srcInstanceId": "cen-9kqshqum*******"  }
	Reserve *string `json:"Reserve,omitempty" xml:"Reserve,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the database to which the migration objects belong in the source instance.
	//
	// > This parameter is available and required only when the source instance or the database engine of the source instance is PolarDB for PostgreSQL (Compatible with Oracle), PostgreSQL, or MongoDB.
	//
	// example:
	//
	// dtstestdatabase
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// The database type of the source instance. Valid values:
	//
	// - **MYSQL**: MySQL database, including ApsaraDB RDS for MySQL and self-managed MySQL.
	//
	// - **MARIADB**: ApsaraDB RDS for MariaDB.
	//
	// - **PolarDB**: PolarDB for MySQL.
	//
	// - **POLARDB_O**: PolarDB for PostgreSQL (Compatible with Oracle).
	//
	// - **POLARDBX10**: PolarDB-X 1.0 (formerly DRDS).
	//
	// - **POLARDBX20**: cloud-native distributed database PolarDB-X 2.0.
	//
	// - **ADB30**: AnalyticDB for MySQL V3.0.
	//
	// - **ORACLE**: Self-managed Oracle database.
	//
	// - **POSTGRESQL**: PostgreSQL database, including ApsaraDB RDS for PostgreSQL and self-managed PostgreSQL.
	//
	// - **MSSQL**: SQL Server database, including ApsaraDB RDS for SQL Server and self-managed SQL Server.
	//
	// - **MONGODB**: MongoDB database, including self-managed MongoDB and ApsaraDB for MongoDB.
	//
	// - **DB2**: Self-managed Db2 for LUW database.
	//
	// - **AS400**: Self-managed Db2 for i database.
	//
	// - **DMSPOLARDB**: DMS logical database.
	//
	// - **HBASE**: Self-managed HBase database.
	//
	// - **TERADATA**: Teradata database.
	//
	// - **TiDB**: TiDB database.
	//
	// - **REDIS**: Redis database, including self-managed Redis and Tair (Redis OSS-Compatible).
	//
	// - **LINDORM**: cloud-native multi-model database Lindorm.
	//
	//
	// > - Default value: **MYSQL**.
	//
	// - If the database type of the source instance is set to **MONGODB**, you must also specify some information in the reserved parameter Reserve. For more information about the configuration method, see [Reserve metric description](https://help.aliyun.com/document_detail/273111.html).
	//
	// example:
	//
	// MYSQL
	SourceEndpointEngineName *string `json:"SourceEndpointEngineName,omitempty" xml:"SourceEndpointEngineName,omitempty"`
	// The IP address of the source instance.
	//
	// > This parameter is available and required only when **SourceEndpointInstanceType*	- is set to **OTHER**, **EXPRESS**, **DG**, or **CEN**.
	//
	// example:
	//
	// ``172.16.**.**``*
	SourceEndpointIP *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	// The instance ID of the source instance.
	//
	// If the source instance is an ApsaraDB database such as ApsaraDB RDS for MySQL, you must specify the instance ID of the ApsaraDB database.
	//
	// If the source instance is a self-managed database, the value of this parameter varies based on the value of **SourceEndpointInstanceType**. For example:
	//
	// - If **SourceEndpointInstanceType*	- is set to **ECS**, specify the instance ID of the ECS instance.
	//
	// - If **SourceEndpointInstanceType*	- is set to **DG**, specify the instance ID of the database gateway.
	//
	// - If **SourceEndpointInstanceType*	- is set to **EXPRESS*	- or **CEN**, specify the instance ID of the VPC that is connected to the source database.
	//
	// > If the value is **CEN**, you must also specify the instance ID of the CEN instance in the reserved parameter Reserve. For more information about the configuration method, see [Reserve metric description](https://help.aliyun.com/document_detail/273111.html).
	//
	// example:
	//
	// rm-bp1imrtn6fq7h****
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The type of the source instance. Valid values:
	//
	// **Alibaba Cloud databases**
	//
	// - **RDS**: ApsaraDB RDS for MySQL, ApsaraDB RDS for SQL Server, ApsaraDB RDS for PostgreSQL, or ApsaraDB RDS for MariaDB.
	//
	// - **PolarDB**: PolarDB for MySQL.
	//
	// - **ADS**: AnalyticDB for MySQL.
	//
	// - **REDIS**: Tair (Redis® OSS-Compatible).
	//
	// - **DISTRIBUTED_POLARDBX10**: PolarDB-X 1.0 (formerly DRDS).
	//
	// - **POLARDBX20**: PolarDB-X 2.0.
	//
	// - **MONGODB**: ApsaraDB for MongoDB.
	//
	// - **DISTRIBUTED_DMSLOGICDB**: Data Management (DMS) logical database.
	//
	// - **LINDORM**: Lindorm.
	//
	// **Self-managed databases**
	//
	// - **OTHER**: A self-managed database with a public IP address.
	//
	// - **ECS**: A self-managed database hosted on an ECS instance.
	//
	// - **EXPRESS**: A self-managed database connected over an Express Connect circuit.
	//
	// - **CEN**: A self-managed database connected over a Cloud Enterprise Network (CEN) instance.
	//
	// - **DG**: A self-managed database connected over a database gateway.
	//
	//
	// > - If the source instance is a PolarDB for PostgreSQL (Oracle-Compatible) cluster, set this parameter to **OTHER*	- or **EXPRESS**. The cluster is treated as a self-managed database and is connected by using a public IP address or an Express Connect circuit.
	//
	// - For more information about the supported source and destination databases, see [Supported databases](https://help.aliyun.com/document_detail/131497.html).
	//
	// - If the source instance is a self-managed database, you must also complete the corresponding preparations. For more information, see [Preparations overview](https://help.aliyun.com/document_detail/130607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// The system ID (SID) of the Oracle database.
	//
	// > This parameter is available and required only when **SourceEndpointEngineName*	- is set to **Oracle*	- and the **Oracle*	- database is a non-RAC instance.
	//
	// example:
	//
	// testsid
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// The Alibaba Cloud account ID to which the source instance belongs.
	//
	// > Specifying this parameter indicates that you execute cross-Alibaba Cloud account data migration or synchronization. You must also specify the **SourceEndpointRole*	- parameter.
	//
	// example:
	//
	// 140692647406****
	SourceEndpointOwnerID *string `json:"SourceEndpointOwnerID,omitempty" xml:"SourceEndpointOwnerID,omitempty"`
	// The password of the database account of the source database.
	//
	// example:
	//
	// Test123456
	SourceEndpointPassword *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	// The database service port of the source instance.
	//
	// > This parameter is available and required only when the source instance is a self-managed database.
	//
	// example:
	//
	// 3306
	SourceEndpointPort *string `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// The area of the source instance. For details, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > This parameter is required if the source instance is an ApsaraDB database.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The name of the role configured for the Alibaba Cloud account to which the source instance belongs.
	//
	// > This parameter is required when you migrate or synchronize data across Alibaba Cloud accounts. For more information about the permissions required for the role and the authorization method, see [Configure RAM authorization for cross-account data migration or synchronization](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	SourceEndpointRole *string `json:"SourceEndpointRole,omitempty" xml:"SourceEndpointRole,omitempty"`
	// The database account of the source database.
	//
	// > - In most cases, you must specify the database account of the source database.
	//
	// - The permissions required for migrating or synchronizing data vary based on the database. For more information about the permission requirements, see [Prepare database accounts for data migration](https://help.aliyun.com/document_detail/175878.html) and [Prepare database accounts for data synchronization](https://help.aliyun.com/document_detail/213152.html).
	//
	// example:
	//
	// dtstest
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	// The instance ID of the vSwitch for the data delivery link.
	//
	// example:
	//
	// vsw-bp10df3mxae6lpmku****
	SourceEndpointVSwitchID *string `json:"SourceEndpointVSwitchID,omitempty" xml:"SourceEndpointVSwitchID,omitempty"`
	// The path of the CA certificate when the source database is connected over SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcCaCertificateOssUrl *string `json:"SrcCaCertificateOssUrl,omitempty" xml:"SrcCaCertificateOssUrl,omitempty"`
	// The password of the CA certificate when the source database is connected over SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcCaCertificatePassword *string `json:"SrcCaCertificatePassword,omitempty" xml:"SrcCaCertificatePassword,omitempty"`
	// The path of the client certificate when the source database is connected over SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcClientCertOssUrl *string `json:"SrcClientCertOssUrl,omitempty" xml:"SrcClientCertOssUrl,omitempty"`
	// The path of the client certificate private key when the source database is connected over SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcClientKeyOssUrl *string `json:"SrcClientKeyOssUrl,omitempty" xml:"SrcClientKeyOssUrl,omitempty"`
	// The password of the client certificate private key when the source database is connected over SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcClientPassword *string `json:"SrcClientPassword,omitempty" xml:"SrcClientPassword,omitempty"`
	// The primary vSwitch at the source end of the VPC NAT gateway.
	//
	// example:
	//
	// ****
	SrcPrimaryVswId *string `json:"SrcPrimaryVswId,omitempty" xml:"SrcPrimaryVswId,omitempty"`
	// The secondary vSwitch at the source end of the VPC NAT gateway.
	//
	// example:
	//
	// ****
	SrcSecondaryVswId *string `json:"SrcSecondaryVswId,omitempty" xml:"SrcSecondaryVswId,omitempty"`
	// Specifies whether to perform schema migration or initial schema synchronization. Valid values:
	//
	// - **true**: Yes. This is the default value.
	//
	// - **false**: No.
	//
	// > If **JobType*	- is set to **CHECK**, you can only set this parameter to **false**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**: Forward.
	//
	// - **Reverse**: Reverse.
	//
	// > - Default value: **Forward**.
	//
	// - This parameter takes effect only when the synchronization topology of the synchronization task is two-way and this parameter is set to **Reverse**.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
}

func (s ConfigureDtsJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureDtsJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ConfigureDtsJobAdvanceRequest) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *ConfigureDtsJobAdvanceRequest) GetDataCheckConfigure() *string {
	return s.DataCheckConfigure
}

func (s *ConfigureDtsJobAdvanceRequest) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *ConfigureDtsJobAdvanceRequest) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *ConfigureDtsJobAdvanceRequest) GetDbList() *string {
	return s.DbList
}

func (s *ConfigureDtsJobAdvanceRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *ConfigureDtsJobAdvanceRequest) GetDelayNotice() *bool {
	return s.DelayNotice
}

func (s *ConfigureDtsJobAdvanceRequest) GetDelayPhone() *string {
	return s.DelayPhone
}

func (s *ConfigureDtsJobAdvanceRequest) GetDelayRuleTime() *int64 {
	return s.DelayRuleTime
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestCaCertificateOssUrl() *string {
	return s.DestCaCertificateOssUrl
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestCaCertificatePassword() *string {
	return s.DestCaCertificatePassword
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestClientCertOssUrl() *string {
	return s.DestClientCertOssUrl
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestClientKeyOssUrl() *string {
	return s.DestClientKeyOssUrl
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestClientPassword() *string {
	return s.DestClientPassword
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestPrimaryVswId() *string {
	return s.DestPrimaryVswId
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestSecondaryVswId() *string {
	return s.DestSecondaryVswId
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestinationEndpointDataBaseName() *string {
	return s.DestinationEndpointDataBaseName
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestinationEndpointEngineName() *string {
	return s.DestinationEndpointEngineName
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestinationEndpointIP() *string {
	return s.DestinationEndpointIP
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestinationEndpointInstanceID() *string {
	return s.DestinationEndpointInstanceID
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestinationEndpointInstanceType() *string {
	return s.DestinationEndpointInstanceType
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestinationEndpointOracleSID() *string {
	return s.DestinationEndpointOracleSID
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestinationEndpointOwnerID() *string {
	return s.DestinationEndpointOwnerID
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestinationEndpointPassword() *string {
	return s.DestinationEndpointPassword
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestinationEndpointPort() *string {
	return s.DestinationEndpointPort
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestinationEndpointRegion() *string {
	return s.DestinationEndpointRegion
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestinationEndpointRole() *string {
	return s.DestinationEndpointRole
}

func (s *ConfigureDtsJobAdvanceRequest) GetDestinationEndpointUserName() *string {
	return s.DestinationEndpointUserName
}

func (s *ConfigureDtsJobAdvanceRequest) GetDisasterRecoveryJob() *bool {
	return s.DisasterRecoveryJob
}

func (s *ConfigureDtsJobAdvanceRequest) GetDtsBisLabel() *string {
	return s.DtsBisLabel
}

func (s *ConfigureDtsJobAdvanceRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ConfigureDtsJobAdvanceRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ConfigureDtsJobAdvanceRequest) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *ConfigureDtsJobAdvanceRequest) GetErrorNotice() *bool {
	return s.ErrorNotice
}

func (s *ConfigureDtsJobAdvanceRequest) GetErrorPhone() *string {
	return s.ErrorPhone
}

func (s *ConfigureDtsJobAdvanceRequest) GetFileOssUrlObject() io.Reader {
	return s.FileOssUrlObject
}

func (s *ConfigureDtsJobAdvanceRequest) GetJobType() *string {
	return s.JobType
}

func (s *ConfigureDtsJobAdvanceRequest) GetMaxDu() *float64 {
	return s.MaxDu
}

func (s *ConfigureDtsJobAdvanceRequest) GetMinDu() *float64 {
	return s.MinDu
}

func (s *ConfigureDtsJobAdvanceRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ConfigureDtsJobAdvanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureDtsJobAdvanceRequest) GetReserve() *string {
	return s.Reserve
}

func (s *ConfigureDtsJobAdvanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointDatabaseName() *string {
	return s.SourceEndpointDatabaseName
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointEngineName() *string {
	return s.SourceEndpointEngineName
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointIP() *string {
	return s.SourceEndpointIP
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointInstanceID() *string {
	return s.SourceEndpointInstanceID
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointInstanceType() *string {
	return s.SourceEndpointInstanceType
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointOracleSID() *string {
	return s.SourceEndpointOracleSID
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointOwnerID() *string {
	return s.SourceEndpointOwnerID
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointPassword() *string {
	return s.SourceEndpointPassword
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointPort() *string {
	return s.SourceEndpointPort
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointRole() *string {
	return s.SourceEndpointRole
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointUserName() *string {
	return s.SourceEndpointUserName
}

func (s *ConfigureDtsJobAdvanceRequest) GetSourceEndpointVSwitchID() *string {
	return s.SourceEndpointVSwitchID
}

func (s *ConfigureDtsJobAdvanceRequest) GetSrcCaCertificateOssUrl() *string {
	return s.SrcCaCertificateOssUrl
}

func (s *ConfigureDtsJobAdvanceRequest) GetSrcCaCertificatePassword() *string {
	return s.SrcCaCertificatePassword
}

func (s *ConfigureDtsJobAdvanceRequest) GetSrcClientCertOssUrl() *string {
	return s.SrcClientCertOssUrl
}

func (s *ConfigureDtsJobAdvanceRequest) GetSrcClientKeyOssUrl() *string {
	return s.SrcClientKeyOssUrl
}

func (s *ConfigureDtsJobAdvanceRequest) GetSrcClientPassword() *string {
	return s.SrcClientPassword
}

func (s *ConfigureDtsJobAdvanceRequest) GetSrcPrimaryVswId() *string {
	return s.SrcPrimaryVswId
}

func (s *ConfigureDtsJobAdvanceRequest) GetSrcSecondaryVswId() *string {
	return s.SrcSecondaryVswId
}

func (s *ConfigureDtsJobAdvanceRequest) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *ConfigureDtsJobAdvanceRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ConfigureDtsJobAdvanceRequest) SetCheckpoint(v string) *ConfigureDtsJobAdvanceRequest {
	s.Checkpoint = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDataCheckConfigure(v string) *ConfigureDtsJobAdvanceRequest {
	s.DataCheckConfigure = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDataInitialization(v bool) *ConfigureDtsJobAdvanceRequest {
	s.DataInitialization = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDataSynchronization(v bool) *ConfigureDtsJobAdvanceRequest {
	s.DataSynchronization = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDbList(v string) *ConfigureDtsJobAdvanceRequest {
	s.DbList = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDedicatedClusterId(v string) *ConfigureDtsJobAdvanceRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDelayNotice(v bool) *ConfigureDtsJobAdvanceRequest {
	s.DelayNotice = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDelayPhone(v string) *ConfigureDtsJobAdvanceRequest {
	s.DelayPhone = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDelayRuleTime(v int64) *ConfigureDtsJobAdvanceRequest {
	s.DelayRuleTime = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestCaCertificateOssUrl(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestCaCertificateOssUrl = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestCaCertificatePassword(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestCaCertificatePassword = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestClientCertOssUrl(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestClientCertOssUrl = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestClientKeyOssUrl(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestClientKeyOssUrl = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestClientPassword(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestClientPassword = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestPrimaryVswId(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestPrimaryVswId = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestSecondaryVswId(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestSecondaryVswId = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestinationEndpointDataBaseName(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestinationEndpointDataBaseName = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestinationEndpointEngineName(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestinationEndpointEngineName = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestinationEndpointIP(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestinationEndpointIP = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestinationEndpointInstanceID(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestinationEndpointInstanceID = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestinationEndpointInstanceType(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestinationEndpointInstanceType = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestinationEndpointOracleSID(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestinationEndpointOracleSID = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestinationEndpointOwnerID(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestinationEndpointOwnerID = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestinationEndpointPassword(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestinationEndpointPassword = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestinationEndpointPort(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestinationEndpointPort = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestinationEndpointRegion(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestinationEndpointRegion = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestinationEndpointRole(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestinationEndpointRole = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDestinationEndpointUserName(v string) *ConfigureDtsJobAdvanceRequest {
	s.DestinationEndpointUserName = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDisasterRecoveryJob(v bool) *ConfigureDtsJobAdvanceRequest {
	s.DisasterRecoveryJob = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDtsBisLabel(v string) *ConfigureDtsJobAdvanceRequest {
	s.DtsBisLabel = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDtsInstanceId(v string) *ConfigureDtsJobAdvanceRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDtsJobId(v string) *ConfigureDtsJobAdvanceRequest {
	s.DtsJobId = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetDtsJobName(v string) *ConfigureDtsJobAdvanceRequest {
	s.DtsJobName = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetErrorNotice(v bool) *ConfigureDtsJobAdvanceRequest {
	s.ErrorNotice = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetErrorPhone(v string) *ConfigureDtsJobAdvanceRequest {
	s.ErrorPhone = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetFileOssUrlObject(v io.Reader) *ConfigureDtsJobAdvanceRequest {
	s.FileOssUrlObject = v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetJobType(v string) *ConfigureDtsJobAdvanceRequest {
	s.JobType = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetMaxDu(v float64) *ConfigureDtsJobAdvanceRequest {
	s.MaxDu = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetMinDu(v float64) *ConfigureDtsJobAdvanceRequest {
	s.MinDu = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetOwnerId(v string) *ConfigureDtsJobAdvanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetRegionId(v string) *ConfigureDtsJobAdvanceRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetReserve(v string) *ConfigureDtsJobAdvanceRequest {
	s.Reserve = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetResourceGroupId(v string) *ConfigureDtsJobAdvanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointDatabaseName(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointEngineName(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointEngineName = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointIP(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointInstanceID(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointInstanceType(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointOracleSID(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointOwnerID(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointOwnerID = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointPassword(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointPort(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointRegion(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointRole(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointRole = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointUserName(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSourceEndpointVSwitchID(v string) *ConfigureDtsJobAdvanceRequest {
	s.SourceEndpointVSwitchID = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSrcCaCertificateOssUrl(v string) *ConfigureDtsJobAdvanceRequest {
	s.SrcCaCertificateOssUrl = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSrcCaCertificatePassword(v string) *ConfigureDtsJobAdvanceRequest {
	s.SrcCaCertificatePassword = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSrcClientCertOssUrl(v string) *ConfigureDtsJobAdvanceRequest {
	s.SrcClientCertOssUrl = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSrcClientKeyOssUrl(v string) *ConfigureDtsJobAdvanceRequest {
	s.SrcClientKeyOssUrl = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSrcClientPassword(v string) *ConfigureDtsJobAdvanceRequest {
	s.SrcClientPassword = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSrcPrimaryVswId(v string) *ConfigureDtsJobAdvanceRequest {
	s.SrcPrimaryVswId = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSrcSecondaryVswId(v string) *ConfigureDtsJobAdvanceRequest {
	s.SrcSecondaryVswId = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetStructureInitialization(v bool) *ConfigureDtsJobAdvanceRequest {
	s.StructureInitialization = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) SetSynchronizationDirection(v string) *ConfigureDtsJobAdvanceRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ConfigureDtsJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
