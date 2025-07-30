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
	// The start offset of incremental data migration or incremental data synchronization. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1610540493
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The parameters for data verification, including the configurations for data verification and alerts. The value is a JSON string. For more information, see [DataCheckConfigure parameter description](https://help.aliyun.com/document_detail/459023.html).
	//
	// example:
	//
	// {"fullCheckModel":1,"fullCheckRatio":20,"checkMaximumHourEnable":1,"checkMaximumHour":1,"fullCheckErrorNotice":true,"fullCheckValidFailNotice":true,"fullCheckNoticeValue":8,"incrementalCheckErrorNotice":true,"incrementalCheckValidFailNotice":true,"incrementalCheckValidFailNoticeTimes":2,"incrementalCheckValidFailNoticePeriod":1,"incrementalCheckValidFailNoticeValue":1,"incrementalCheckDelayNotice":true,"incrementalCheckDelayNoticeTimes":2,"incrementalCheckDelayNoticePeriod":1,"incrementalCheckDelayNoticeValue":60,"fullDataCheck":true,"incrementalDataCheck":true,"dataCheckNoticePhone":"13126800****","dataCheckDbList":{"dts":{"name":"dts","all":true}}}
	DataCheckConfigure *string `json:"DataCheckConfigure,omitempty" xml:"DataCheckConfigure,omitempty"`
	// Specifies whether to perform full data migration or full data synchronization. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// > If **JobType*	- is set to **CHECK**, set this parameter to **false**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Specifies whether to perform incremental data migration or incremental data synchronization. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// > If **JobType*	- is set to **CHECK**, set this parameter to **false**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// The objects that you want to migrate or synchronize. The value is a JSON string. For more information, see [Objects of DTS tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// example:
	//
	// {"dtstest":{"name":"dtstest","all":true}}
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The ID of the DTS dedicated cluster on which the task runs.
	//
	// > If this parameter is specified, the task is scheduled to the specified DTS dedicated cluster.
	//
	// example:
	//
	// dtscluster_atyl3b5214uk***
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// Specifies whether to monitor task latency. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DelayNotice *bool `json:"DelayNotice,omitempty" xml:"DelayNotice,omitempty"`
	// The mobile phone numbers to which latency-related alerts are sent. Separate multiple mobile phone numbers with commas (,).
	//
	// >
	//
	// 	- This parameter is available only for users of the China site (aliyun.com). Only mobile phone numbers in the Chinese mainland are supported. You can specify up to 10 mobile phone numbers.
	//
	// 	- Users of the international site (alibabacloud.com) cannot receive alerts by using mobile phone numbers, but can configure alert rules for DTS tasks in the CloudMonitor console. For more information, see [Configure alert rules for DTS tasks in the CloudMonitor console](https://help.aliyun.com/document_detail/175876.html).
	//
	// example:
	//
	// 1361234****,1371234****
	DelayPhone *string `json:"DelayPhone,omitempty" xml:"DelayPhone,omitempty"`
	// The threshold for latency alerts. Unit: seconds. The value must be an integer. You can set the threshold based on your business requirements. To prevent unstable latency caused by network and database overloads, we recommend that you set the threshold to more than 10 seconds.
	//
	// > If **DelayNotice*	- is set to **true**, this parameter is required.
	//
	// example:
	//
	// 10
	DelayRuleTime *int64 `json:"DelayRuleTime,omitempty" xml:"DelayRuleTime,omitempty"`
	// The path of the CA certificate that is used if the connection to the destination database is encrypted by using SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestCaCertificateOssUrl *string `json:"DestCaCertificateOssUrl,omitempty" xml:"DestCaCertificateOssUrl,omitempty"`
	// The key of the CA certificate that is used if the connection to the destination database is encrypted by using SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestCaCertificatePassword *string `json:"DestCaCertificatePassword,omitempty" xml:"DestCaCertificatePassword,omitempty"`
	// The path to the client certificate that is used if the connection to the destination database is encrypted by using SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestClientCertOssUrl *string `json:"DestClientCertOssUrl,omitempty" xml:"DestClientCertOssUrl,omitempty"`
	// The path to the private key of the client certificate that is used if the connection to the destination database is encrypted by using SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestClientKeyOssUrl *string `json:"DestClientKeyOssUrl,omitempty" xml:"DestClientKeyOssUrl,omitempty"`
	// The password of the private key of the client certificate that is used if the connection to the destination database is encrypted by using SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	DestClientPassword *string `json:"DestClientPassword,omitempty" xml:"DestClientPassword,omitempty"`
	// VPCNAT destination main VSW
	//
	// example:
	//
	// ****
	DestPrimaryVswId *string `json:"DestPrimaryVswId,omitempty" xml:"DestPrimaryVswId,omitempty"`
	// VPCNAT destination backup VSW
	//
	// example:
	//
	// ****
	DestSecondaryVswId *string `json:"DestSecondaryVswId,omitempty" xml:"DestSecondaryVswId,omitempty"`
	// The name of the database to which the objects are migrated or synchronized in the destination instance.
	//
	// >
	//
	// 	- This parameter is valid and required only if the destination database is a PolarDB for PostgreSQL (Compatible with Oracle) cluster, an AnalyticDB for PostgreSQL instance, a PostgreSQL database, a MaxCompute project, or a MongoDB database.
	//
	// 	- If the destination instance is a MaxCompute project, you must specify the MaxCompute project ID.
	//
	// example:
	//
	// dtstestdata
	DestinationEndpointDataBaseName *string `json:"DestinationEndpointDataBaseName,omitempty" xml:"DestinationEndpointDataBaseName,omitempty"`
	// The type of the destination database. Valid values:
	//
	// 	- **MYSQL**: ApsaraDB RDS for MySQL instance or self-managed MySQL database.
	//
	// 	- **MARIADB**: ApsaraDB RDS for MariaDB instance.
	//
	// 	- **PolarDB**: PolarDB for MySQL cluster.
	//
	// 	- **POLARDB_O**: PolarDB for PostgreSQL (Compatible with Oracle) cluster.
	//
	// 	- **POLARDBX10**: PolarDB-X 1.0 instance (formerly DRDS).
	//
	// 	- **POLARDBX20**: PolarDB-X 2.0 instance.
	//
	// 	- **ORACLE**: self-managed Oracle database.
	//
	// 	- **POSTGRESQL**: ApsaraDB RDS for PostgreSQL instance or self-managed PostgreSQL database.
	//
	// 	- **MSSQL**: ApsaraDB RDS for SQL Server instance or self-managed SQL Server database.
	//
	// 	- **ADS**: AnalyticDB for MySQL V2.0 cluster.
	//
	// 	- **ADB30**: AnalyticDB for MySQL V3.0 cluster.
	//
	// 	- **MONGODB**: ApsaraDB for MongoDB instance or self-managed MongoDB database.
	//
	// 	- **GREENPLUM**: AnalyticDB for PostgreSQL instance.
	//
	// 	- **KAFKA**: ApsaraMQ for Kafka instance or self-managed Kafka cluster.
	//
	// 	- **DATAHUB**: DataHub project.
	//
	// 	- **DB2**: self-managed Db2 for LUW database.
	//
	// 	- **AS400**: Db2 for i database.
	//
	// 	- **ODPS**: MaxCompute project.
	//
	// 	- **Tablestore**: Tablestore instance.
	//
	// 	- **ELK**: Elasticsearch cluster.
	//
	// 	- **REDIS**: ApsaraDB for Redis instance or self-managed Redis database.
	//
	// >
	//
	// 	- Default value: **MYSQL**.
	//
	// 	- If this parameter is set to **KAFKA**, **MONGODB**, or **PolarDB**, you must also specify the database information in Reserve. For more information, see [Reserve parameter](https://help.aliyun.com/document_detail/273111.html).
	//
	// example:
	//
	// MYSQL
	DestinationEndpointEngineName *string `json:"DestinationEndpointEngineName,omitempty" xml:"DestinationEndpointEngineName,omitempty"`
	// The IP address of the destination instance.
	//
	// > This parameter is valid and required only if **DestinationEndpointInstanceType*	- is set to **OTHER**, **EXPRESS**, **DG**, or **CEN**.
	//
	// example:
	//
	// ``172.16.**.**``*
	DestinationEndpointIP *string `json:"DestinationEndpointIP,omitempty" xml:"DestinationEndpointIP,omitempty"`
	// The destination instance ID.
	//
	// If the destination instance is an Alibaba Cloud database instance, you must specify the database instance ID. For example, if the destination instance is an ApsaraDB RDS for MySQL instance, you must specify the ID of the ApsaraDB RDS for MySQL instance.
	//
	// If the destination instance is a self-managed database, the value of this parameter varies with the value of **DestinationEndpointInstanceType**.****
	//
	// 	- If DestinationEndpointInstanceType is set to **ECS**, you must specify the ECS instance ID.
	//
	// 	- If DestinationEndpointInstanceType is set to **DG**, you must specify the database gateway ID.
	//
	// 	- If DestinationEndpointInstanceType is set to **EXPRESS*	- or **CEN**, you must specify the ID of the VPC that is connected to the source instance.
	//
	// > If DestinationEndpointInstanceType is set to **CEN**, you must also specify the ID of the CEN instance in Reserve. For more information, see [Reserve parameter](https://help.aliyun.com/document_detail/273111.html).
	//
	// example:
	//
	// vpc-bp1opxu1zkhn00gzv****
	DestinationEndpointInstanceID *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	// The type of the destination instance. Valid values:
	//
	// **Alibaba Cloud database instance**
	//
	// 	- **RDS**: ApsaraDB RDS for MySQL instance, ApsaraDB RDS for SQL Server instance, ApsaraDB RDS for PostgreSQL instance, or ApsaraDB RDS for MariaDB instance.
	//
	// 	- **PolarDB**: PolarDB for MySQL cluster.
	//
	// 	- **DISTRIBUTED_POLARDBX10**: PolarDB-X 1.0 instance (formerly DRDS).
	//
	// 	- **POLARDBX20**: PolarDB-X 2.0 instance.
	//
	// 	- **REDIS**: ApsaraDB for Redis instance.
	//
	// 	- **ADS**: AnalyticDB for MySQL V2.0 cluster or AnalyticDB for MySQL V3.0 cluster.
	//
	// 	- **MONGODB**: ApsaraDB for MongoDB instance.
	//
	// 	- **GREENPLUM**: AnalyticDB for PostgreSQL instance.
	//
	// 	- **DATAHUB**: DataHub project.
	//
	// 	- **ELK**: Elasticsearch cluster.
	//
	// 	- **Tablestore**: Tablestore instance.
	//
	// 	- **ODPS**: MaxCompute project.
	//
	// **Self-managed database**
	//
	// 	- **OTHER**: self-managed database with a public IP address.
	//
	// 	- **ECS**: self-managed database hosted on an ECS instance.
	//
	// 	- **EXPRESS**: self-managed database connected over Express Connect.
	//
	// 	- **CEN**: self-managed database connected over Cloud Enterprise Network (CEN).
	//
	// 	- **DG**: self-managed database connected over Database Gateway.
	//
	// >
	//
	// 	- If the destination instance is a PolarDB for PostgreSQL (Compatible with Oracle) cluster, you must connect the cluster to DTS as a self-managed database by using a public IP address or Express Connect and set this parameter to **OTHER*	- or **EXPRESS**.
	//
	// 	- If the destination instance is an ApsaraMQ for Kafka instance, you must connect the instance to DTS as a self-managed database by using ECS or Express Connect and set this parameter to **ECS*	- or **EXPRESS**.
	//
	// 	- For more information, see [Supported source and destination databases](https://help.aliyun.com/document_detail/176064.html).
	//
	// 	- If the destination instance is a self-managed database, you must deploy the network environment for the database. For more information, see [Preparation overview](https://help.aliyun.com/document_detail/146958.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// EXPRESS
	DestinationEndpointInstanceType *string `json:"DestinationEndpointInstanceType,omitempty" xml:"DestinationEndpointInstanceType,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is valid and required only if **DestinationEndpointEngineName*	- is set to **ORACLE*	- and the **Oracle*	- database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	DestinationEndpointOracleSID *string `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	// The ID of the Alibaba Cloud account to which the destination ApsaraDB RDS for MySQL instance belongs.
	//
	// >
	//
	// 	- This parameter is available only if the destination instance is an ApsaraDB RDS for MySQL instance.
	//
	// 	- You can specify this parameter to migrate or synchronize data across different Alibaba Cloud accounts. In this case, you must specify **DestinationEndpointRole**.
	//
	// example:
	//
	// 140692647406****
	DestinationEndpointOwnerID *string `json:"DestinationEndpointOwnerID,omitempty" xml:"DestinationEndpointOwnerID,omitempty"`
	// The password of the account that is used to log on to the destination database.
	//
	// > If the destination database is a MaxCompute project, you must specify the AccessKey secret of your Alibaba Cloud account. For information about how to obtain an AccessKey pair, see [Create an AccessKey pair](https://help.aliyun.com/document_detail/116401.html).
	//
	// example:
	//
	// Test123456
	DestinationEndpointPassword *string `json:"DestinationEndpointPassword,omitempty" xml:"DestinationEndpointPassword,omitempty"`
	// The port number of the destination instance.
	//
	// > This parameter is valid and required only if the destination instance is a self-managed database.
	//
	// example:
	//
	// 3306
	DestinationEndpointPort *string `json:"DestinationEndpointPort,omitempty" xml:"DestinationEndpointPort,omitempty"`
	// The ID of the region in which the destination instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > If the destination instance is an Alibaba Cloud database instance, this parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DestinationEndpointRegion *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account to which the destination instance belongs.
	//
	// > This parameter is required if you migrate or synchronize data across Alibaba Cloud accounts. For information about the permissions and authorization methods of the RAM role, see [Configure RAM authorization for cross-account DTS tasks](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	DestinationEndpointRole *string `json:"DestinationEndpointRole,omitempty" xml:"DestinationEndpointRole,omitempty"`
	// The username of the account that is used to log on to the destination database.
	//
	// >
	//
	// 	- In most cases, this parameter is required.
	//
	// 	- The permissions that are required for the database account vary with the migration or synchronization scenario. For more information, see [Prepare the database accounts for data migration](https://help.aliyun.com/document_detail/175878.html) or [Prepare the database accounts for data synchronization](https://help.aliyun.com/document_detail/213152.html).
	//
	// 	- If the destination database is a MaxCompute project, you must specify the AccessKey ID of your Alibaba Cloud account. For information about how to obtain an AccessKey pair, see [Create an AccessKey pair](https://help.aliyun.com/document_detail/116401.html).
	//
	// example:
	//
	// dtstest
	DestinationEndpointUserName *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	// Specifies whether the instance is a disaster recovery instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DisasterRecoveryJob *bool `json:"DisasterRecoveryJob,omitempty" xml:"DisasterRecoveryJob,omitempty"`
	// The environment tag of the DTS instance. Valid values:
	//
	// 	- **normal******
	//
	// 	- **online******
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
	// Specifies whether to monitor task status. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ErrorNotice *bool `json:"ErrorNotice,omitempty" xml:"ErrorNotice,omitempty"`
	// The mobile phone numbers to which status-related alerts are sent. Separate multiple mobile phone numbers with commas (,).
	//
	// >
	//
	// 	- This parameter is available only for users of the China site (aliyun.com). Only mobile phone numbers in the Chinese mainland are supported. You can specify up to 10 mobile phone numbers.
	//
	// 	- Users of the international site (alibabacloud.com) cannot receive alerts by using mobile phone numbers, but can configure alert rules for DTS tasks in the CloudMonitor console. For more information, see [Configure alert rules for DTS tasks in the CloudMonitor console](https://help.aliyun.com/document_detail/175876.html).
	//
	// example:
	//
	// 1361234****,1371234****
	ErrorPhone *string `json:"ErrorPhone,omitempty" xml:"ErrorPhone,omitempty"`
	// The URL of the Object Storage Service (OSS) bucket that stores the files related to the DTS task.
	//
	// example:
	//
	// http://db-list-os-file.oss-cn-shanghai.aliyuncs.com/8e42_121852**********_79dd3aeabe2f43cdb**************
	FileOssUrl *string `json:"FileOssUrl,omitempty" xml:"FileOssUrl,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **MIGRATION**: data migration task.
	//
	// 	- **SYNC**: data synchronization task.
	//
	// 	- **CHECK**: data verification task. You must separately purchase a data verification instance.
	//
	// > If you set this parameter to **MIGRATION*	- or **SYNC**, you can also enable data verification in the data migration or synchronization task.
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
	// The minimum number of DTS Units (DUs).
	//
	// > This parameter is supported only for serverless instances.
	//
	// example:
	//
	// 1
	MinDu   *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	OwnerId *string  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to add more configurations of the source or destination instance to the DTS task. For example, you can specify the data storage format of the destination Kafka database and the CEN instance ID. For more information, see [Reserve parameter](https://help.aliyun.com/document_detail/273111.html).
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
	// The name of the database from which the objects are migrated or synchronized in the source instance.
	//
	// > This parameter is valid and required only if the source instance is a PolarDB for PostgreSQL (Compatible with Oracle) cluster, a PostgreSQL database, or a MongoDB database.
	//
	// example:
	//
	// dtstestdatabase
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// The database type of the source instance.
	//
	// 	- **MYSQL**: ApsaraDB RDS for MySQL instance or self-managed MySQL database.
	//
	// 	- **MARIADB**: ApsaraDB RDS for MariaDB instance.
	//
	// 	- **PolarDB**: PolarDB for MySQL cluster.
	//
	// 	- **POLARDB_O**: PolarDB for PostgreSQL (Compatible with Oracle) cluster.
	//
	// 	- **POLARDBX10**: PolarDB-X 1.0 instance (formerly DRDS).
	//
	// 	- **POLARDBX20**: PolarDB-X 2.0 instance.
	//
	// 	- **ORACLE**: self-managed Oracle database.
	//
	// 	- **POSTGRESQL**: ApsaraDB RDS for PostgreSQL instance or self-managed PostgreSQL database.
	//
	// 	- **MSSQL**: ApsaraDB RDS for SQL Server instance or self-managed SQL Server database.
	//
	// 	- **MONGODB**: ApsaraDB for MongoDB instance or self-managed MongoDB database.
	//
	// 	- **DB2**: self-managed Db2 for LUW database.
	//
	// 	- **AS400**: self-managed Db2 for i database.
	//
	// 	- **DMSPOLARDB**: DMS logical database.
	//
	// 	- **HBASE**: self-managed HBase database.
	//
	// 	- **TERADATA**: Teradata database.
	//
	// 	- **TiDB**: TiDB database.
	//
	// 	- **REDIS**: ApsaraDB for Redis instance or self-managed Redis database.
	//
	// >
	//
	// 	- Default value: **MYSQL**.
	//
	// 	- If this parameter is set to **MONGODB**, you must also specify the architecture type of the MongoDB database in Reserve. For more information, see [Reserve parameter](https://help.aliyun.com/document_detail/273111.html).
	//
	// example:
	//
	// MYSQL
	SourceEndpointEngineName *string `json:"SourceEndpointEngineName,omitempty" xml:"SourceEndpointEngineName,omitempty"`
	// The IP address of the source instance.
	//
	// > This parameter is valid and required only if **SourceEndpointInstanceType*	- is set to **OTHER**, **EXPRESS**, **DG**, or **CEN**.
	//
	// example:
	//
	// ``172.16.**.**``*
	SourceEndpointIP *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	// The source instance ID.
	//
	// If the source instance is an Alibaba Cloud database instance, you must specify the database instance ID. For example, if the source instance is an ApsaraDB RDS for MySQL instance, you must specify the ID of the ApsaraDB RDS for MySQL instance.
	//
	// If the source instance is a self-managed database, the value of this parameter varies with the value of **SourceEndpointInstanceType**.****
	//
	// 	- If SourceEndpointInstanceType is set to **ECS**, you must specify the ECS instance ID.
	//
	// 	- If SourceEndpointInstanceType is set to **DG**, you must specify the database gateway ID.
	//
	// 	- If SourceEndpointInstanceType is set to **EXPRESS*	- or **CEN**, you must specify the ID of the virtual private cloud (VPC) that is connected to the source instance.
	//
	// > If SourceEndpointInstanceType is set to **CEN**, you must also specify the ID of the CEN instance in Reserve. For more information, see [Reserve parameter](https://help.aliyun.com/document_detail/273111.html).
	//
	// example:
	//
	// rm-bp1imrtn6fq7h****
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The type of the source instance. Valid values:
	//
	// **Alibaba Cloud database instance**
	//
	// 	- **RDS**: ApsaraDB RDS for MySQL instance, ApsaraDB RDS for SQL Server instance, ApsaraDB RDS for PostgreSQL instance, or ApsaraDB RDS for MariaDB instance
	//
	// 	- **PolarDB**: PolarDB for MySQL cluster.
	//
	// 	- **REDIS**: ApsaraDB for Redis instance.
	//
	// 	- **DISTRIBUTED_POLARDBX10**: PolarDB-X 1.0 instance (formerly DRDS).
	//
	// 	- **POLARDBX20**: PolarDB-X 2.0 instance.
	//
	// 	- **MONGODB**: ApsaraDB for MongoDB instance.
	//
	// 	- **DISTRIBUTED_DMSLOGICDB**: Data Management (DMS) logical database
	//
	// **Self-managed database**
	//
	// 	- **OTHER**: self-managed database with a public IP address.
	//
	// 	- **ECS**: self-managed database hosted on an ECS instance.
	//
	// 	- **EXPRESS**: self-managed database connected over Express Connect.
	//
	// 	- **CEN**: self-managed database connected over Cloud Enterprise Network (CEN).
	//
	// 	- **DG**: self-managed database connected over Database Gateway.
	//
	// >
	//
	// 	- If the source instance is a PolarDB for PostgreSQL (Compatible with Oracle) cluster, you must connect the cluster to DTS as a self-managed database by using a public IP address or Express Connect and set this parameter to **OTHER*	- or **EXPRESS**.
	//
	// 	- For more information, see [Supported sources and targets](https://help.aliyun.com/document_detail/176064.html).
	//
	// 	- If the source instance is a self-managed database, you must deploy the network environment for the database. For more information, see [Preparation overview](https://help.aliyun.com/document_detail/146958.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is valid and required only if **SourceEndpointEngineName*	- is set to **ORACLE*	- and the **Oracle*	- database is deployed in a non-Real Application Cluster (RAC) architecture.
	//
	// example:
	//
	// testsid
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// The ID of the Alibaba Cloud account to which the source database belongs.
	//
	// > You can specify this parameter to migrate or synchronize data across different Alibaba Cloud accounts. In this case, you must specify **SourceEndpointRole**.
	//
	// example:
	//
	// 140692647406****
	SourceEndpointOwnerID *string `json:"SourceEndpointOwnerID,omitempty" xml:"SourceEndpointOwnerID,omitempty"`
	// The password of the account that is used to log on to the source database.
	//
	// example:
	//
	// Test123456
	SourceEndpointPassword *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	// The port number of the source instance.
	//
	// > This parameter is required only if the source instance is a self-managed database.
	//
	// example:
	//
	// 3306
	SourceEndpointPort *string `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// The ID of the region in which the source instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// > If the source instance is an Alibaba Cloud database instance, this parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The name of the Resource Access Management (RAM) role configured for the Alibaba Cloud account to which the source instance belongs.
	//
	// > This parameter is required if you migrate or synchronize data across different Alibaba Cloud accounts. For information about the permissions and authorization methods of the RAM role, see [Configure RAM authorization for cross-account DTS tasks](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	SourceEndpointRole *string `json:"SourceEndpointRole,omitempty" xml:"SourceEndpointRole,omitempty"`
	// The username of the account that is used to log on to the source database.
	//
	// >
	//
	// 	- In most cases, this parameter is required.
	//
	// 	- The permissions that are required for the database account vary with the migration or synchronization scenario. For more information, see [Prepare the database accounts for data migration](https://help.aliyun.com/document_detail/175878.html) or [Prepare the database accounts for data synchronization](https://help.aliyun.com/document_detail/213152.html).
	//
	// example:
	//
	// dtstest
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	// The ID of the vSwitch that is used for data shipping.
	//
	// example:
	//
	// vsw-bp10df3mxae6lpmku****
	SourceEndpointVSwitchID *string `json:"SourceEndpointVSwitchID,omitempty" xml:"SourceEndpointVSwitchID,omitempty"`
	// The path of the certificate authority (CA) certificate that is used if the connection to the source database is encrypted by using SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcCaCertificateOssUrl *string `json:"SrcCaCertificateOssUrl,omitempty" xml:"SrcCaCertificateOssUrl,omitempty"`
	// The key of the CA certificate that is used if the connection to the source database is encrypted by using SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcCaCertificatePassword *string `json:"SrcCaCertificatePassword,omitempty" xml:"SrcCaCertificatePassword,omitempty"`
	// The path to the client certificate that is used if the connection to the source database is encrypted by using SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcClientCertOssUrl *string `json:"SrcClientCertOssUrl,omitempty" xml:"SrcClientCertOssUrl,omitempty"`
	// The path to the private key of the client certificate that is used if the connection to the source database is encrypted by using SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcClientKeyOssUrl *string `json:"SrcClientKeyOssUrl,omitempty" xml:"SrcClientKeyOssUrl,omitempty"`
	// The password of the private key of the client certificate that is used if the connection to the source database is encrypted by using SSL.
	//
	// > This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcClientPassword *string `json:"SrcClientPassword,omitempty" xml:"SrcClientPassword,omitempty"`
	// VPCNAT source end main VSW
	//
	// example:
	//
	// ****
	SrcPrimaryVswId *string `json:"SrcPrimaryVswId,omitempty" xml:"SrcPrimaryVswId,omitempty"`
	// VPCNAT source backup VSW
	//
	// example:
	//
	// ****
	SrcSecondaryVswId *string `json:"SrcSecondaryVswId,omitempty" xml:"SrcSecondaryVswId,omitempty"`
	// Specifies whether to perform schema migration or schema synchronization. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// > If **JobType*	- is set to **CHECK**, set this parameter to **false**.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// >
	//
	// 	- The default value is **Forward**.
	//
	// 	- The value **Reverse*	- takes effect only if the topology of the data synchronization task is two-way synchronization.
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
