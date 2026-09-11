// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureDtsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckpoint(v string) *ConfigureDtsJobRequest
	GetCheckpoint() *string
	SetDataCheckConfigure(v string) *ConfigureDtsJobRequest
	GetDataCheckConfigure() *string
	SetDataInitialization(v bool) *ConfigureDtsJobRequest
	GetDataInitialization() *bool
	SetDataSynchronization(v bool) *ConfigureDtsJobRequest
	GetDataSynchronization() *bool
	SetDbList(v string) *ConfigureDtsJobRequest
	GetDbList() *string
	SetDedicatedClusterId(v string) *ConfigureDtsJobRequest
	GetDedicatedClusterId() *string
	SetDelayNotice(v bool) *ConfigureDtsJobRequest
	GetDelayNotice() *bool
	SetDelayPhone(v string) *ConfigureDtsJobRequest
	GetDelayPhone() *string
	SetDelayRuleTime(v int64) *ConfigureDtsJobRequest
	GetDelayRuleTime() *int64
	SetDestCaCertificateOssUrl(v string) *ConfigureDtsJobRequest
	GetDestCaCertificateOssUrl() *string
	SetDestCaCertificatePassword(v string) *ConfigureDtsJobRequest
	GetDestCaCertificatePassword() *string
	SetDestClientCertOssUrl(v string) *ConfigureDtsJobRequest
	GetDestClientCertOssUrl() *string
	SetDestClientKeyOssUrl(v string) *ConfigureDtsJobRequest
	GetDestClientKeyOssUrl() *string
	SetDestClientPassword(v string) *ConfigureDtsJobRequest
	GetDestClientPassword() *string
	SetDestPrimaryVswId(v string) *ConfigureDtsJobRequest
	GetDestPrimaryVswId() *string
	SetDestSecondaryVswId(v string) *ConfigureDtsJobRequest
	GetDestSecondaryVswId() *string
	SetDestinationEndpointDataBaseName(v string) *ConfigureDtsJobRequest
	GetDestinationEndpointDataBaseName() *string
	SetDestinationEndpointEngineName(v string) *ConfigureDtsJobRequest
	GetDestinationEndpointEngineName() *string
	SetDestinationEndpointIP(v string) *ConfigureDtsJobRequest
	GetDestinationEndpointIP() *string
	SetDestinationEndpointInstanceID(v string) *ConfigureDtsJobRequest
	GetDestinationEndpointInstanceID() *string
	SetDestinationEndpointInstanceType(v string) *ConfigureDtsJobRequest
	GetDestinationEndpointInstanceType() *string
	SetDestinationEndpointOracleSID(v string) *ConfigureDtsJobRequest
	GetDestinationEndpointOracleSID() *string
	SetDestinationEndpointOwnerID(v string) *ConfigureDtsJobRequest
	GetDestinationEndpointOwnerID() *string
	SetDestinationEndpointPassword(v string) *ConfigureDtsJobRequest
	GetDestinationEndpointPassword() *string
	SetDestinationEndpointPort(v string) *ConfigureDtsJobRequest
	GetDestinationEndpointPort() *string
	SetDestinationEndpointRegion(v string) *ConfigureDtsJobRequest
	GetDestinationEndpointRegion() *string
	SetDestinationEndpointRole(v string) *ConfigureDtsJobRequest
	GetDestinationEndpointRole() *string
	SetDestinationEndpointUserName(v string) *ConfigureDtsJobRequest
	GetDestinationEndpointUserName() *string
	SetDisasterRecoveryJob(v bool) *ConfigureDtsJobRequest
	GetDisasterRecoveryJob() *bool
	SetDtsBisLabel(v string) *ConfigureDtsJobRequest
	GetDtsBisLabel() *string
	SetDtsInstanceId(v string) *ConfigureDtsJobRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ConfigureDtsJobRequest
	GetDtsJobId() *string
	SetDtsJobName(v string) *ConfigureDtsJobRequest
	GetDtsJobName() *string
	SetErrorNotice(v bool) *ConfigureDtsJobRequest
	GetErrorNotice() *bool
	SetErrorPhone(v string) *ConfigureDtsJobRequest
	GetErrorPhone() *string
	SetFileOssUrl(v string) *ConfigureDtsJobRequest
	GetFileOssUrl() *string
	SetJobType(v string) *ConfigureDtsJobRequest
	GetJobType() *string
	SetMaxDu(v float64) *ConfigureDtsJobRequest
	GetMaxDu() *float64
	SetMinDu(v float64) *ConfigureDtsJobRequest
	GetMinDu() *float64
	SetOwnerId(v string) *ConfigureDtsJobRequest
	GetOwnerId() *string
	SetRegionId(v string) *ConfigureDtsJobRequest
	GetRegionId() *string
	SetReserve(v string) *ConfigureDtsJobRequest
	GetReserve() *string
	SetResourceGroupId(v string) *ConfigureDtsJobRequest
	GetResourceGroupId() *string
	SetSourceEndpointDatabaseName(v string) *ConfigureDtsJobRequest
	GetSourceEndpointDatabaseName() *string
	SetSourceEndpointEngineName(v string) *ConfigureDtsJobRequest
	GetSourceEndpointEngineName() *string
	SetSourceEndpointIP(v string) *ConfigureDtsJobRequest
	GetSourceEndpointIP() *string
	SetSourceEndpointInstanceID(v string) *ConfigureDtsJobRequest
	GetSourceEndpointInstanceID() *string
	SetSourceEndpointInstanceType(v string) *ConfigureDtsJobRequest
	GetSourceEndpointInstanceType() *string
	SetSourceEndpointOracleSID(v string) *ConfigureDtsJobRequest
	GetSourceEndpointOracleSID() *string
	SetSourceEndpointOwnerID(v string) *ConfigureDtsJobRequest
	GetSourceEndpointOwnerID() *string
	SetSourceEndpointPassword(v string) *ConfigureDtsJobRequest
	GetSourceEndpointPassword() *string
	SetSourceEndpointPort(v string) *ConfigureDtsJobRequest
	GetSourceEndpointPort() *string
	SetSourceEndpointRegion(v string) *ConfigureDtsJobRequest
	GetSourceEndpointRegion() *string
	SetSourceEndpointRole(v string) *ConfigureDtsJobRequest
	GetSourceEndpointRole() *string
	SetSourceEndpointUserName(v string) *ConfigureDtsJobRequest
	GetSourceEndpointUserName() *string
	SetSourceEndpointVSwitchID(v string) *ConfigureDtsJobRequest
	GetSourceEndpointVSwitchID() *string
	SetSrcCaCertificateOssUrl(v string) *ConfigureDtsJobRequest
	GetSrcCaCertificateOssUrl() *string
	SetSrcCaCertificatePassword(v string) *ConfigureDtsJobRequest
	GetSrcCaCertificatePassword() *string
	SetSrcClientCertOssUrl(v string) *ConfigureDtsJobRequest
	GetSrcClientCertOssUrl() *string
	SetSrcClientKeyOssUrl(v string) *ConfigureDtsJobRequest
	GetSrcClientKeyOssUrl() *string
	SetSrcClientPassword(v string) *ConfigureDtsJobRequest
	GetSrcClientPassword() *string
	SetSrcPrimaryVswId(v string) *ConfigureDtsJobRequest
	GetSrcPrimaryVswId() *string
	SetSrcSecondaryVswId(v string) *ConfigureDtsJobRequest
	GetSrcSecondaryVswId() *string
	SetStructureInitialization(v bool) *ConfigureDtsJobRequest
	GetStructureInitialization() *bool
	SetSynchronizationDirection(v string) *ConfigureDtsJobRequest
	GetSynchronizationDirection() *string
}

type ConfigureDtsJobRequest struct {
	// The start position for incremental data migration or the synchronization checkpoint, in the format of a UNIX timestamp. Unit: seconds.
	//
	// > If you specify the **Checkpoint*	- parameter, make sure that no other running DTS instance has the same source database as the destination DTS instance.
	//
	// example:
	//
	// 1610540493
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The parameters of the data validation node, in JSON character string format, such as parameter limits and alert configuration. For more information, see [DataCheckConfigure parameter description](https://help.aliyun.com/document_detail/459023.html).
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
	// > If **JobType*	- is set to **CHECK**, this parameter can only be set to **false**.
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
	// > If **JobType*	- is set to **CHECK**, this parameter can only be set to **false**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// The objects to be migrated or synchronized, in JSON format. For more information, see [Objects of migration, synchronization, or change tracking tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// - The maximum size of the DbList value is 1 MB.
	//
	// - If DbList contains filter conditions, the total length of DbList (including filter conditions) cannot exceed 1 MB.
	//
	// - For distributed tasks (such as migration or synchronization tasks with PolarDB-X 1.0 as the source), DbList is split based on physical shards and multiple subtasks are generated. The maximum size of DbList for each subtask is 1 MB.
	//
	// example:
	//
	// {"dtstest":{"name":"dtstest","all":true}}
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The ID of the DTS dedicated cluster.
	//
	// > If you specify the ID of a dedicated cluster, the task is scheduled to the corresponding cluster.
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
	// The mobile phone numbers for latency alerting of the contact. Separate multiple phone numbers with commas (,).
	//
	// > - This parameter is supported only on the China site. Only the Chinese mainland phone numbers are supported, and a maximum of 10 phone numbers can be specified.
	//
	// - The international site does not support phone alerting. You can only [configure alert rules for DTS tasks through the CloudMonitor platform to set alert rules](https://help.aliyun.com/document_detail/175876.html).
	//
	// example:
	//
	// 1361234****,1371234****
	DelayPhone *string `json:"DelayPhone,omitempty" xml:"DelayPhone,omitempty"`
	// The threshold for triggering latency alerts. Unit: seconds. The value must be an integer. Set the threshold based on your business requirements. To avoid alert fluctuations caused by network conditions or database loads, set the threshold to 10 seconds or more.
	//
	// > This parameter is required when **DelayNotice*	- is set to **true**.
	//
	// example:
	//
	// 10
	DelayRuleTime *int64 `json:"DelayRuleTime,omitempty" xml:"DelayRuleTime,omitempty"`
	// The path of the CA certificate for SSL connection to the destination database.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestCaCertificateOssUrl *string `json:"DestCaCertificateOssUrl,omitempty" xml:"DestCaCertificateOssUrl,omitempty"`
	// The password of the CA certificate for SSL connection to the destination database.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestCaCertificatePassword *string `json:"DestCaCertificatePassword,omitempty" xml:"DestCaCertificatePassword,omitempty"`
	// The path of the client certificate for SSL connection to the destination database.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestClientCertOssUrl *string `json:"DestClientCertOssUrl,omitempty" xml:"DestClientCertOssUrl,omitempty"`
	// The path of the client certificate private key for SSL connection to the destination database.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestClientKeyOssUrl *string `json:"DestClientKeyOssUrl,omitempty" xml:"DestClientKeyOssUrl,omitempty"`
	// The password of the client certificate private key for SSL connection to the destination database.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestClientPassword *string `json:"DestClientPassword,omitempty" xml:"DestClientPassword,omitempty"`
	// The primary vSwitch of the VPC NAT gateway on the destination side.
	//
	// example:
	//
	// ****
	DestPrimaryVswId *string `json:"DestPrimaryVswId,omitempty" xml:"DestPrimaryVswId,omitempty"`
	// The secondary vSwitch of the VPC NAT gateway on the destination side.
	//
	// example:
	//
	// ****
	DestSecondaryVswId *string `json:"DestSecondaryVswId,omitempty" xml:"DestSecondaryVswId,omitempty"`
	// The name of the database to which the objects to be migrated belong in the destination instance.
	//
	// > - This parameter is available and required only when the destination instance or destination database type is PolarDB for PostgreSQL (Compatible with Oracle), AnalyticDB for PostgreSQL, PostgreSQL, MaxCompute, or MongoDB.
	//
	// - If the destination database is MaxCompute, specify the project of the MaxCompute instance.
	//
	// example:
	//
	// dtstestdata
	DestinationEndpointDataBaseName *string `json:"DestinationEndpointDataBaseName,omitempty" xml:"DestinationEndpointDataBaseName,omitempty"`
	// The database type of the destination instance. Valid values:
	//
	// - **MYSQL**: MySQL database (including ApsaraDB RDS for MySQL and self-managed MySQL).
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
	// - **ORACLE**: self-managed Oracle.
	//
	// - **PostgreSQL**: PostgreSQL database (including ApsaraDB RDS for PostgreSQL and self-managed PostgreSQL).
	//
	// - **MSSQL**: SQL Server database (including ApsaraDB RDS for SQL Server and self-managed SQL Server).
	//
	// - **ADS**: AnalyticDB for MySQL 2.0.
	//
	// - **ADB30**: AnalyticDB for MySQL 3.0.
	//
	// - **MONGODB**: MongoDB database (including self-managed MongoDB and ApsaraDB for MongoDB).
	//
	// - **ROCKETMQ**: ApsaraMQ for RocketMQ.
	//
	// - **GREENPLUM**: AnalyticDB for PostgreSQL.
	//
	// - **KAFKA**: Kafka database (including MSMQ for Apache Kafka and self-managed Kafka).
	//
	// - **DATAHUB**: Alibaba Cloud DataHub.
	//
	// - **DB2**: self-managed Db2 for LUW.
	//
	// - **AS400**: Db2 for i.
	//
	// - **ODPS**: MaxCompute.
	//
	// - **Tablestore**: Tablestore.
	//
	// - **ELK**: Alibaba Cloud Elasticsearch.
	//
	// - **REDIS**: Redis database, including self-managed Redis and Tair (Redis® OSS-Compatible).
	//
	// - **LINDORM**: cloud-native multi-model database Lindorm.
	//
	// > - Default value: **MYSQL**.
	//
	// - If the database type of the destination instance is set to **KAFKA**, **MONGODB**, or **PolarDB**, you must also specify additional information in the Reserve parameter. For the metric description, see [Reserve parameter description](https://help.aliyun.com/document_detail/273111.html).
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
	// The ID of the destination instance.
	//
	//  If the destination instance is an Alibaba Cloud database (such as ApsaraDB RDS for MySQL), specify the ID of the Alibaba Cloud database instance (such as the ApsaraDB RDS for MySQL instance ID).
	//
	//  If the destination instance is a self-managed database, the value of this parameter varies based on the value of **DestinationEndpointInstanceType**. Example:
	//
	//
	// - **ECS**: Specify the ID of the ECS instance.
	//
	// - **DG**: Specify the ID of the database gateway.
	//
	// - **EXPRESS*	- or **CEN**: Specify the ID of the VPC that is connected to the source database.
	//
	// > If the value is **CEN**, you must also specify the CEN instance ID in the Reserve parameter. For the metric description, see [Reserve parameter description](https://help.aliyun.com/document_detail/273111.html).
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
	// - **REDIS**: Tair (Redis® OSS-Compatible).
	//
	// - **ADS**: AnalyticDB for MySQL 2.0 or 3.0.
	//
	// - **MONGODB**: ApsaraDB for MongoDB.
	//
	// - **ROCKETMQ**: ApsaraMQ for RocketMQ.
	//
	// - **GREENPLUM**: AnalyticDB for PostgreSQL.
	//
	// - **DATAHUB**: Alibaba Cloud DataHub platform.
	//
	// - **ELK**: Alibaba Cloud Elasticsearch.
	//
	// - **Tablestore**: Tablestore.
	//
	// - **ODPS**: MaxCompute.
	//
	// - **LINDORM**: cloud-native multi-model database Lindorm.
	//
	// **Self-managed databases**
	//
	// - **OTHER**: self-managed database with a public IP address.
	//
	// - **ECS**: self-managed database hosted on ECS.
	//
	// - **EXPRESS**: self-managed database connected over Express Connect.
	//
	// - **CEN**: self-managed database connected over Cloud Enterprise Network (CEN).
	//
	// - **DG**: self-managed database connected over Database Gateway.
	//
	// > - If the destination instance is a PolarDB for PostgreSQL (Compatible with Oracle) cluster, set this parameter to **OTHER*	- or **EXPRESS*	- to connect the cluster as a self-managed database over a public IP address or Express Connect.
	//
	// - If the destination instance is MSMQ for Apache Kafka, set this parameter to **ECS*	- or **EXPRESS*	- to connect the instance as a self-managed database over ECS or Express Connect.
	//
	// - For information about supported source and destination database combinations, see <props="china">[Supported databases](https://help.aliyun.com/document_detail/131497.html)<props="intl">[Supported source and destination databases](https://help.aliyun.com/document_detail/176064.html).
	//
	// - If the destination instance is a self-managed database, you must also execute the required preparations. For more information, see [Preparations overview](https://help.aliyun.com/document_detail/146958.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// EXPRESS
	DestinationEndpointInstanceType *string `json:"DestinationEndpointInstanceType,omitempty" xml:"DestinationEndpointInstanceType,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is available and required only when **DestinationEndpointEngineName*	- is set to **Oracle*	- and the Oracle database is a non-RAC instance.
	//
	// example:
	//
	// testsid
	DestinationEndpointOracleSID *string `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	// The Alibaba Cloud account ID to which the destination ApsaraDB RDS for MySQL instance belongs.
	//
	// > - This parameter can be configured only when the destination instance is ApsaraDB RDS for MySQL.
	//
	// - Specifying this parameter indicates you execute a cross-account data migration or synchronization. You must also specify the **DestinationEndpointRole*	- parameter.
	//
	// example:
	//
	// 140692647406****
	DestinationEndpointOwnerID *string `json:"DestinationEndpointOwnerID,omitempty" xml:"DestinationEndpointOwnerID,omitempty"`
	// The password of the destination database account.
	//
	// > If the destination database is MaxCompute, specify the AccessKey secret of the Alibaba Cloud account. For more information about how to obtain the AccessKey secret, see [Create an AccessKey pair](https://help.aliyun.com/document_detail/116401.html).
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
	// The region of the destination instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > If the destination instance is an Alibaba Cloud database, this parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DestinationEndpointRegion *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account to which the destination instance belongs.
	//
	// > This parameter is required for cross-account data migration or synchronization. For information about the permissions and authorization method required for this role, see [Configure RAM authorization for cross-account data migration or synchronization](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	DestinationEndpointRole *string `json:"DestinationEndpointRole,omitempty" xml:"DestinationEndpointRole,omitempty"`
	// The database account of the destination database.
	//
	// > - In most cases, you must specify the database account of the destination database.
	//
	// - The required permissions vary depending on the database being migrated or synchronized. For more information, see [Prepare database accounts for data migration](https://help.aliyun.com/document_detail/175878.html) and [Prepare database accounts for data synchronization](https://help.aliyun.com/document_detail/213152.html).
	//
	// - If the destination database is MaxCompute, specify the AccessKey ID of the Alibaba Cloud account. For more information about how to obtain the AccessKey ID, see [Create an AccessKey pair](https://help.aliyun.com/document_detail/116401.html).
	//
	// example:
	//
	// dtstest
	DestinationEndpointUserName *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	// Specifies whether this is a disaster recovery instance. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	DisasterRecoveryJob *bool `json:"DisasterRecoveryJob,omitempty" xml:"DisasterRecoveryJob,omitempty"`
	// The environment label of the DTS instance. Valid values:
	//
	// - **normal**: normal
	//
	// - **online**: online.
	//
	// example:
	//
	// normal
	DtsBisLabel *string `json:"DtsBisLabel,omitempty" xml:"DtsBisLabel,omitempty"`
	// The ID of the migration or synchronization instance.
	//
	// > You can call [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) to query the instance ID.
	//
	// example:
	//
	// dtsk2gm967v16f****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the migration or synchronization task.
	//
	// > You can call [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) to query the task ID.
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
	// Specifies whether to monitor the error status. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	ErrorNotice *bool `json:"ErrorNotice,omitempty" xml:"ErrorNotice,omitempty"`
	// The mobile phone numbers for error alerting of the contact. Separate multiple phone numbers with commas (,).
	//
	// > - This parameter is supported only on the China site. Only the Chinese mainland phone numbers are supported, and a maximum of 10 phone numbers can be specified.
	//
	// - The international site does not support phone alerting. You can only [configure alert rules for DTS tasks through the CloudMonitor platform to set alert rules](https://help.aliyun.com/document_detail/175876.html).
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
	FileOssUrl *string `json:"FileOssUrl,omitempty" xml:"FileOssUrl,omitempty"`
	// The type of the node. Valid values:
	//
	// - **MIGRATION**: data migration.
	//
	// - **SYNC**: data synchronization.
	//
	// - **CHECK**: data validation (purchased separately).
	//
	// > - If the value is **MIGRATION*	- or **SYNC**, you can also configure a data validation node within the migration or synchronization instance.
	//
	// - To configure a data validation node, you must also specify the **DataCheckConfigure*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// SYNC
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The maximum number of DTS Units (DUs).
	//
	// > This parameter is supported only for serverless instances.
	//
	// example:
	//
	// 16
	MaxDu *float64 `json:"MaxDu,omitempty" xml:"MaxDu,omitempty"`
	// The minimum number of DTS Units (DUs).
	//
	// > This parameter is supported only for serverless instances.
	//
	// example:
	//
	// 1
	MinDu   *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	OwnerId *string  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the DTS instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The reserved parameter of DTS, in JSON character string format. You can specify this parameter to add information about the source and destination databases (such as the data storage format of the destination Kafka database, the CEN instance ID, and ETL feature configurations). For more information, see [Reserve parameter description](https://help.aliyun.com/document_detail/273111.html).
	//
	// example:
	//
	// {      "srcInstanceId": "cen-9kqshqum*******"  }
	Reserve *string `json:"Reserve,omitempty" xml:"Reserve,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the database to which the objects to be migrated belong in the source instance.
	//
	// > This parameter is available and required only when the source instance or its database type is PolarDB for PostgreSQL (Compatible with Oracle), PostgreSQL, or MongoDB.
	//
	// example:
	//
	// dtstestdatabase
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// The database type of the source instance. Valid values:
	//
	// - **MYSQL**: MySQL database (including ApsaraDB RDS for MySQL and self-managed MySQL).
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
	// - **ADB30**: AnalyticDB for MySQL 3.0.
	//
	// - **ORACLE**: self-managed Oracle.
	//
	// - **POSTGRESQL**: PostgreSQL database (including ApsaraDB RDS for PostgreSQL and self-managed PostgreSQL).
	//
	// - **MSSQL**: SQL Server database (including ApsaraDB RDS for SQL Server and self-managed SQL Server).
	//
	// - **MONGODB**: MongoDB database (including self-managed MongoDB and ApsaraDB for MongoDB).
	//
	// - **DB2**: self-managed Db2 for LUW.
	//
	// - **AS400**: self-managed Db2 for i.
	//
	// - **DMSPOLARDB**: Data Management (DMS) logical database.
	//
	// - **HBASE**: self-managed HBase database.
	//
	// - **TERADATA**: Teradata database.
	//
	// - **TiDB**: TiDB database.
	//
	// - **REDIS**: Redis database, including self-managed Redis and Tair (Redis® OSS-Compatible).
	//
	// - **LINDORM**: Lindorm.
	//
	//
	// > - Default value: **MYSQL**.
	//
	//  - If the database type of the source instance is set to **MONGODB**, you must also specify additional information in the Reserve parameter. For the metric description, see [Reserve parameter description](https://help.aliyun.com/document_detail/273111.html).
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
	// The ID of the source instance.
	//
	// If the source instance is an Alibaba Cloud database (such as ApsaraDB RDS for MySQL), specify the ID of the Alibaba Cloud database instance (such as the ApsaraDB RDS for MySQL instance ID).
	//
	// If the source instance is a self-managed database, the value of this parameter varies based on the value of **SourceEndpointInstanceType**. Example:
	//
	// - **ECS**: Specify the ID of the ECS instance.
	//
	// - **DG**: Specify the ID of the database gateway.
	//
	// - **EXPRESS*	- or **CEN**: Specify the ID of the VPC that is connected to the source database.
	//
	// > If the value is **CEN**, you must also specify the CEN instance ID in the Reserve parameter. For the metric description, see [Reserve parameter description](https://help.aliyun.com/document_detail/273111.html).
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
	// - **OTHER**: self-managed database with a public IP address.
	//
	// - **ECS**: self-managed database hosted on ECS.
	//
	// - **EXPRESS**: self-managed database connected over Express Connect.
	//
	// - **CEN**: self-managed database connected over Cloud Enterprise Network (CEN).
	//
	// - **DG**: self-managed database connected over Database Gateway.
	//
	//
	// > - If the source instance is a PolarDB for PostgreSQL (Compatible with Oracle) cluster, set this parameter to **OTHER*	- or **EXPRESS*	- to connect the cluster as a self-managed database over a public IP address or Express Connect.
	//
	// - For information about supported source and destination database combinations, see [Supported databases](https://help.aliyun.com/document_detail/131497.html).
	//
	// - If the source instance is a self-managed database, you must complete the required preparations. For more information, see [Preparations overview](https://help.aliyun.com/document_detail/130607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is available and required only when **SourceEndpointEngineName*	- is set to **Oracle*	- and the Oracle database is a non-RAC instance.
	//
	// example:
	//
	// testsid
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// The Alibaba Cloud account ID to which the source instance belongs.
	//
	// > Specifying this parameter indicates you execute a cross-account data migration or synchronization. You must also specify the **SourceEndpointRole*	- parameter.
	//
	// example:
	//
	// 140692647406****
	SourceEndpointOwnerID *string `json:"SourceEndpointOwnerID,omitempty" xml:"SourceEndpointOwnerID,omitempty"`
	// The password of the source database account.
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
	// The region of the source instance. For details, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > If the source instance is an Alibaba Cloud database, this parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account to which the source instance belongs.
	//
	// > This parameter is required for cross-account data migration or synchronization. For information about the permissions and authorization method required for this role, see [Configure RAM authorization for cross-account data migration or synchronization](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	SourceEndpointRole *string `json:"SourceEndpointRole,omitempty" xml:"SourceEndpointRole,omitempty"`
	// The database account of the source database.
	//
	// > - In most cases, you must specify the database account of the source database.
	//
	// - The required permissions vary depending on the database being migrated or synchronized. For more information, see [Prepare database accounts for data migration](https://help.aliyun.com/document_detail/175878.html) and [Prepare database accounts for data synchronization](https://help.aliyun.com/document_detail/213152.html).
	//
	// example:
	//
	// dtstest
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	// The vSwitch instance ID for the data delivery link.
	//
	// example:
	//
	// vsw-bp10df3mxae6lpmku****
	SourceEndpointVSwitchID *string `json:"SourceEndpointVSwitchID,omitempty" xml:"SourceEndpointVSwitchID,omitempty"`
	// The path of the CA certificate for SSL connection to the source database.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcCaCertificateOssUrl *string `json:"SrcCaCertificateOssUrl,omitempty" xml:"SrcCaCertificateOssUrl,omitempty"`
	// The password of the CA certificate for SSL connection to the source database.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcCaCertificatePassword *string `json:"SrcCaCertificatePassword,omitempty" xml:"SrcCaCertificatePassword,omitempty"`
	// The path of the client certificate for SSL connection to the source database.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcClientCertOssUrl *string `json:"SrcClientCertOssUrl,omitempty" xml:"SrcClientCertOssUrl,omitempty"`
	// The path of the client certificate private key for SSL connection to the source database.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcClientKeyOssUrl *string `json:"SrcClientKeyOssUrl,omitempty" xml:"SrcClientKeyOssUrl,omitempty"`
	// The password of the client certificate private key for SSL connection to the source database.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcClientPassword *string `json:"SrcClientPassword,omitempty" xml:"SrcClientPassword,omitempty"`
	// The primary vSwitch of the VPC NAT gateway on the source side.
	//
	// example:
	//
	// ****
	SrcPrimaryVswId *string `json:"SrcPrimaryVswId,omitempty" xml:"SrcPrimaryVswId,omitempty"`
	// The secondary vSwitch of the VPC NAT gateway on the source side.
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
	// > If **JobType*	- is set to **CHECK**, this parameter can only be set to **false**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward.
	//
	// - **Reverse**: reverse.
	//
	// > - Default value: **Forward**.
	//
	// - The value **Reverse*	- takes effect only when the synchronization topology of the synchronization task is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
}

func (s ConfigureDtsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureDtsJobRequest) GoString() string {
	return s.String()
}

func (s *ConfigureDtsJobRequest) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *ConfigureDtsJobRequest) GetDataCheckConfigure() *string {
	return s.DataCheckConfigure
}

func (s *ConfigureDtsJobRequest) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *ConfigureDtsJobRequest) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *ConfigureDtsJobRequest) GetDbList() *string {
	return s.DbList
}

func (s *ConfigureDtsJobRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *ConfigureDtsJobRequest) GetDelayNotice() *bool {
	return s.DelayNotice
}

func (s *ConfigureDtsJobRequest) GetDelayPhone() *string {
	return s.DelayPhone
}

func (s *ConfigureDtsJobRequest) GetDelayRuleTime() *int64 {
	return s.DelayRuleTime
}

func (s *ConfigureDtsJobRequest) GetDestCaCertificateOssUrl() *string {
	return s.DestCaCertificateOssUrl
}

func (s *ConfigureDtsJobRequest) GetDestCaCertificatePassword() *string {
	return s.DestCaCertificatePassword
}

func (s *ConfigureDtsJobRequest) GetDestClientCertOssUrl() *string {
	return s.DestClientCertOssUrl
}

func (s *ConfigureDtsJobRequest) GetDestClientKeyOssUrl() *string {
	return s.DestClientKeyOssUrl
}

func (s *ConfigureDtsJobRequest) GetDestClientPassword() *string {
	return s.DestClientPassword
}

func (s *ConfigureDtsJobRequest) GetDestPrimaryVswId() *string {
	return s.DestPrimaryVswId
}

func (s *ConfigureDtsJobRequest) GetDestSecondaryVswId() *string {
	return s.DestSecondaryVswId
}

func (s *ConfigureDtsJobRequest) GetDestinationEndpointDataBaseName() *string {
	return s.DestinationEndpointDataBaseName
}

func (s *ConfigureDtsJobRequest) GetDestinationEndpointEngineName() *string {
	return s.DestinationEndpointEngineName
}

func (s *ConfigureDtsJobRequest) GetDestinationEndpointIP() *string {
	return s.DestinationEndpointIP
}

func (s *ConfigureDtsJobRequest) GetDestinationEndpointInstanceID() *string {
	return s.DestinationEndpointInstanceID
}

func (s *ConfigureDtsJobRequest) GetDestinationEndpointInstanceType() *string {
	return s.DestinationEndpointInstanceType
}

func (s *ConfigureDtsJobRequest) GetDestinationEndpointOracleSID() *string {
	return s.DestinationEndpointOracleSID
}

func (s *ConfigureDtsJobRequest) GetDestinationEndpointOwnerID() *string {
	return s.DestinationEndpointOwnerID
}

func (s *ConfigureDtsJobRequest) GetDestinationEndpointPassword() *string {
	return s.DestinationEndpointPassword
}

func (s *ConfigureDtsJobRequest) GetDestinationEndpointPort() *string {
	return s.DestinationEndpointPort
}

func (s *ConfigureDtsJobRequest) GetDestinationEndpointRegion() *string {
	return s.DestinationEndpointRegion
}

func (s *ConfigureDtsJobRequest) GetDestinationEndpointRole() *string {
	return s.DestinationEndpointRole
}

func (s *ConfigureDtsJobRequest) GetDestinationEndpointUserName() *string {
	return s.DestinationEndpointUserName
}

func (s *ConfigureDtsJobRequest) GetDisasterRecoveryJob() *bool {
	return s.DisasterRecoveryJob
}

func (s *ConfigureDtsJobRequest) GetDtsBisLabel() *string {
	return s.DtsBisLabel
}

func (s *ConfigureDtsJobRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ConfigureDtsJobRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ConfigureDtsJobRequest) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *ConfigureDtsJobRequest) GetErrorNotice() *bool {
	return s.ErrorNotice
}

func (s *ConfigureDtsJobRequest) GetErrorPhone() *string {
	return s.ErrorPhone
}

func (s *ConfigureDtsJobRequest) GetFileOssUrl() *string {
	return s.FileOssUrl
}

func (s *ConfigureDtsJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *ConfigureDtsJobRequest) GetMaxDu() *float64 {
	return s.MaxDu
}

func (s *ConfigureDtsJobRequest) GetMinDu() *float64 {
	return s.MinDu
}

func (s *ConfigureDtsJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ConfigureDtsJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureDtsJobRequest) GetReserve() *string {
	return s.Reserve
}

func (s *ConfigureDtsJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointDatabaseName() *string {
	return s.SourceEndpointDatabaseName
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointEngineName() *string {
	return s.SourceEndpointEngineName
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointIP() *string {
	return s.SourceEndpointIP
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointInstanceID() *string {
	return s.SourceEndpointInstanceID
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointInstanceType() *string {
	return s.SourceEndpointInstanceType
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointOracleSID() *string {
	return s.SourceEndpointOracleSID
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointOwnerID() *string {
	return s.SourceEndpointOwnerID
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointPassword() *string {
	return s.SourceEndpointPassword
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointPort() *string {
	return s.SourceEndpointPort
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointRole() *string {
	return s.SourceEndpointRole
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointUserName() *string {
	return s.SourceEndpointUserName
}

func (s *ConfigureDtsJobRequest) GetSourceEndpointVSwitchID() *string {
	return s.SourceEndpointVSwitchID
}

func (s *ConfigureDtsJobRequest) GetSrcCaCertificateOssUrl() *string {
	return s.SrcCaCertificateOssUrl
}

func (s *ConfigureDtsJobRequest) GetSrcCaCertificatePassword() *string {
	return s.SrcCaCertificatePassword
}

func (s *ConfigureDtsJobRequest) GetSrcClientCertOssUrl() *string {
	return s.SrcClientCertOssUrl
}

func (s *ConfigureDtsJobRequest) GetSrcClientKeyOssUrl() *string {
	return s.SrcClientKeyOssUrl
}

func (s *ConfigureDtsJobRequest) GetSrcClientPassword() *string {
	return s.SrcClientPassword
}

func (s *ConfigureDtsJobRequest) GetSrcPrimaryVswId() *string {
	return s.SrcPrimaryVswId
}

func (s *ConfigureDtsJobRequest) GetSrcSecondaryVswId() *string {
	return s.SrcSecondaryVswId
}

func (s *ConfigureDtsJobRequest) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *ConfigureDtsJobRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ConfigureDtsJobRequest) SetCheckpoint(v string) *ConfigureDtsJobRequest {
	s.Checkpoint = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDataCheckConfigure(v string) *ConfigureDtsJobRequest {
	s.DataCheckConfigure = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDataInitialization(v bool) *ConfigureDtsJobRequest {
	s.DataInitialization = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDataSynchronization(v bool) *ConfigureDtsJobRequest {
	s.DataSynchronization = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDbList(v string) *ConfigureDtsJobRequest {
	s.DbList = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDedicatedClusterId(v string) *ConfigureDtsJobRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDelayNotice(v bool) *ConfigureDtsJobRequest {
	s.DelayNotice = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDelayPhone(v string) *ConfigureDtsJobRequest {
	s.DelayPhone = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDelayRuleTime(v int64) *ConfigureDtsJobRequest {
	s.DelayRuleTime = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestCaCertificateOssUrl(v string) *ConfigureDtsJobRequest {
	s.DestCaCertificateOssUrl = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestCaCertificatePassword(v string) *ConfigureDtsJobRequest {
	s.DestCaCertificatePassword = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestClientCertOssUrl(v string) *ConfigureDtsJobRequest {
	s.DestClientCertOssUrl = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestClientKeyOssUrl(v string) *ConfigureDtsJobRequest {
	s.DestClientKeyOssUrl = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestClientPassword(v string) *ConfigureDtsJobRequest {
	s.DestClientPassword = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestPrimaryVswId(v string) *ConfigureDtsJobRequest {
	s.DestPrimaryVswId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestSecondaryVswId(v string) *ConfigureDtsJobRequest {
	s.DestSecondaryVswId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointDataBaseName(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointDataBaseName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointEngineName(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointEngineName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointIP(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointIP = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointInstanceID(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointInstanceID = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointInstanceType(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointInstanceType = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointOracleSID(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointOracleSID = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointOwnerID(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointOwnerID = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointPassword(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointPassword = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointPort(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointPort = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointRegion(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointRegion = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointRole(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointRole = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointUserName(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointUserName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDisasterRecoveryJob(v bool) *ConfigureDtsJobRequest {
	s.DisasterRecoveryJob = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDtsBisLabel(v string) *ConfigureDtsJobRequest {
	s.DtsBisLabel = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDtsInstanceId(v string) *ConfigureDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDtsJobId(v string) *ConfigureDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDtsJobName(v string) *ConfigureDtsJobRequest {
	s.DtsJobName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetErrorNotice(v bool) *ConfigureDtsJobRequest {
	s.ErrorNotice = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetErrorPhone(v string) *ConfigureDtsJobRequest {
	s.ErrorPhone = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetFileOssUrl(v string) *ConfigureDtsJobRequest {
	s.FileOssUrl = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetJobType(v string) *ConfigureDtsJobRequest {
	s.JobType = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetMaxDu(v float64) *ConfigureDtsJobRequest {
	s.MaxDu = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetMinDu(v float64) *ConfigureDtsJobRequest {
	s.MinDu = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetOwnerId(v string) *ConfigureDtsJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetRegionId(v string) *ConfigureDtsJobRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetReserve(v string) *ConfigureDtsJobRequest {
	s.Reserve = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetResourceGroupId(v string) *ConfigureDtsJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointDatabaseName(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointEngineName(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointEngineName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointIP(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointInstanceID(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointInstanceType(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointOracleSID(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointOwnerID(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointOwnerID = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointPassword(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointPort(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointRegion(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointRole(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointRole = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointUserName(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointVSwitchID(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointVSwitchID = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSrcCaCertificateOssUrl(v string) *ConfigureDtsJobRequest {
	s.SrcCaCertificateOssUrl = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSrcCaCertificatePassword(v string) *ConfigureDtsJobRequest {
	s.SrcCaCertificatePassword = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSrcClientCertOssUrl(v string) *ConfigureDtsJobRequest {
	s.SrcClientCertOssUrl = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSrcClientKeyOssUrl(v string) *ConfigureDtsJobRequest {
	s.SrcClientKeyOssUrl = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSrcClientPassword(v string) *ConfigureDtsJobRequest {
	s.SrcClientPassword = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSrcPrimaryVswId(v string) *ConfigureDtsJobRequest {
	s.SrcPrimaryVswId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSrcSecondaryVswId(v string) *ConfigureDtsJobRequest {
	s.SrcSecondaryVswId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetStructureInitialization(v bool) *ConfigureDtsJobRequest {
	s.StructureInitialization = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSynchronizationDirection(v string) *ConfigureDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ConfigureDtsJobRequest) Validate() error {
	return dara.Validate(s)
}
