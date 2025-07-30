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
	FileOssUrlObject io.Reader `json:"FileOssUrl,omitempty" xml:"FileOssUrl,omitempty"`
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
