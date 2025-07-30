// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckpoint(v string) *ConfigureSubscriptionRequest
	GetCheckpoint() *string
	SetDbList(v string) *ConfigureSubscriptionRequest
	GetDbList() *string
	SetDedicatedClusterId(v string) *ConfigureSubscriptionRequest
	GetDedicatedClusterId() *string
	SetDelayNotice(v bool) *ConfigureSubscriptionRequest
	GetDelayNotice() *bool
	SetDelayPhone(v string) *ConfigureSubscriptionRequest
	GetDelayPhone() *string
	SetDelayRuleTime(v int64) *ConfigureSubscriptionRequest
	GetDelayRuleTime() *int64
	SetDtsBisLabel(v string) *ConfigureSubscriptionRequest
	GetDtsBisLabel() *string
	SetDtsInstanceId(v string) *ConfigureSubscriptionRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ConfigureSubscriptionRequest
	GetDtsJobId() *string
	SetDtsJobName(v string) *ConfigureSubscriptionRequest
	GetDtsJobName() *string
	SetErrorNotice(v bool) *ConfigureSubscriptionRequest
	GetErrorNotice() *bool
	SetErrorPhone(v string) *ConfigureSubscriptionRequest
	GetErrorPhone() *string
	SetMaxDu(v float64) *ConfigureSubscriptionRequest
	GetMaxDu() *float64
	SetMinDu(v float64) *ConfigureSubscriptionRequest
	GetMinDu() *float64
	SetRegionId(v string) *ConfigureSubscriptionRequest
	GetRegionId() *string
	SetReserve(v string) *ConfigureSubscriptionRequest
	GetReserve() *string
	SetResourceGroupId(v string) *ConfigureSubscriptionRequest
	GetResourceGroupId() *string
	SetSourceEndpointDatabaseName(v string) *ConfigureSubscriptionRequest
	GetSourceEndpointDatabaseName() *string
	SetSourceEndpointEngineName(v string) *ConfigureSubscriptionRequest
	GetSourceEndpointEngineName() *string
	SetSourceEndpointIP(v string) *ConfigureSubscriptionRequest
	GetSourceEndpointIP() *string
	SetSourceEndpointInstanceID(v string) *ConfigureSubscriptionRequest
	GetSourceEndpointInstanceID() *string
	SetSourceEndpointInstanceType(v string) *ConfigureSubscriptionRequest
	GetSourceEndpointInstanceType() *string
	SetSourceEndpointOracleSID(v string) *ConfigureSubscriptionRequest
	GetSourceEndpointOracleSID() *string
	SetSourceEndpointOwnerID(v string) *ConfigureSubscriptionRequest
	GetSourceEndpointOwnerID() *string
	SetSourceEndpointPassword(v string) *ConfigureSubscriptionRequest
	GetSourceEndpointPassword() *string
	SetSourceEndpointPort(v string) *ConfigureSubscriptionRequest
	GetSourceEndpointPort() *string
	SetSourceEndpointRegion(v string) *ConfigureSubscriptionRequest
	GetSourceEndpointRegion() *string
	SetSourceEndpointRole(v string) *ConfigureSubscriptionRequest
	GetSourceEndpointRole() *string
	SetSourceEndpointUserName(v string) *ConfigureSubscriptionRequest
	GetSourceEndpointUserName() *string
	SetSrcCaCertificateOssUrl(v string) *ConfigureSubscriptionRequest
	GetSrcCaCertificateOssUrl() *string
	SetSrcCaCertificatePassword(v string) *ConfigureSubscriptionRequest
	GetSrcCaCertificatePassword() *string
	SetSrcClientCertOssUrl(v string) *ConfigureSubscriptionRequest
	GetSrcClientCertOssUrl() *string
	SetSrcClientKeyOssUrl(v string) *ConfigureSubscriptionRequest
	GetSrcClientKeyOssUrl() *string
	SetSrcClientPassword(v string) *ConfigureSubscriptionRequest
	GetSrcClientPassword() *string
	SetSubscriptionDataTypeDDL(v bool) *ConfigureSubscriptionRequest
	GetSubscriptionDataTypeDDL() *bool
	SetSubscriptionDataTypeDML(v bool) *ConfigureSubscriptionRequest
	GetSubscriptionDataTypeDML() *bool
	SetSubscriptionInstanceNetworkType(v string) *ConfigureSubscriptionRequest
	GetSubscriptionInstanceNetworkType() *string
	SetSubscriptionInstanceVPCId(v string) *ConfigureSubscriptionRequest
	GetSubscriptionInstanceVPCId() *string
	SetSubscriptionInstanceVSwitchId(v string) *ConfigureSubscriptionRequest
	GetSubscriptionInstanceVSwitchId() *string
}

type ConfigureSubscriptionRequest struct {
	// The UNIX timestamp that represents the start time of change tracking. Unit: seconds.
	//
	// >  You can use a search engine to obtain a UNIX timestamp converter.
	//
	// example:
	//
	// 1616902385
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The objects for which you want to track data changes. The value must be a JSON string. For more information, see [Objects of DTS tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"dtstest":{"name":"dtstest","all":true}}
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The ID of the DTS dedicated cluster on which the change tracking task is scheduled to run.
	//
	// example:
	//
	// dtscluster_atyl3b5214uk***
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// Specifies whether to monitor the task latency. Valid values:
	//
	// 	- **true**: monitors the task latency.
	//
	// 	- **false**: does not monitor the task latency.
	//
	// example:
	//
	// true
	DelayNotice *bool `json:"DelayNotice,omitempty" xml:"DelayNotice,omitempty"`
	// The mobile numbers to which latency-related alerts are sent. Separate multiple mobile numbers with commas (,).
	//
	// >
	//
	// 	- This parameter is available only for users of the China site (aliyun.com). Only mobile numbers in the Chinese mainland are supported. You can specify up to 10 mobile numbers.
	//
	// 	- Users of the international site (alibabacloud.com) cannot receive alerts by using mobile phones, but can [configure alert rules for DTS tasks in the CloudMonitor console](https://help.aliyun.com/document_detail/175876.html).
	//
	// example:
	//
	// 1361234****,1371234****
	DelayPhone *string `json:"DelayPhone,omitempty" xml:"DelayPhone,omitempty"`
	// The threshold for triggering latency-related alerts. Unit: seconds. The value must be an integer. You can set the threshold based on your business needs. To prevent jitters caused by network and database overloads, we recommend that you set the threshold to more than 10 seconds.
	//
	// >  If the **DelayNotice*	- parameter is set to **true**, this parameter is required.
	//
	// example:
	//
	// 10
	DelayRuleTime *int64 `json:"DelayRuleTime,omitempty" xml:"DelayRuleTime,omitempty"`
	// Environment label of the DTS instance, with values:
	//
	// - **normal**: **general*	- - **online**: **production**
	//
	// example:
	//
	// normal
	DtsBisLabel *string `json:"DtsBisLabel,omitempty" xml:"DtsBisLabel,omitempty"`
	// The ID of the change tracking instance. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the instance ID.
	//
	// example:
	//
	// dtsy0zz3t13h7d****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the change tracking task. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// example:
	//
	// y0zz3t13h7d****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The name of the change tracking task.
	//
	// >  We recommend that you specify a descriptive name for easy identification. You do not need to use a unique name.
	//
	// This parameter is required.
	//
	// example:
	//
	// for_test
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// Specifies whether to monitor the task status. Valid values:
	//
	// 	- **true**: monitors the task status.
	//
	// 	- **false**: does not monitor the task status.
	//
	// example:
	//
	// true
	ErrorNotice *bool `json:"ErrorNotice,omitempty" xml:"ErrorNotice,omitempty"`
	// The mobile numbers to which status-related alerts are sent. Separate multiple mobile numbers with commas (,).
	//
	// >
	//
	// 	- This parameter is available only for users of the China site (aliyun.com). Only mobile numbers in the Chinese mainland are supported. You can specify up to 10 mobile numbers.
	//
	// 	- Users of the international site (alibabacloud.com) cannot receive alerts by using mobile phones, but can [configure alert rules for DTS tasks in the CloudMonitor console](https://help.aliyun.com/document_detail/175876.html).
	//
	// example:
	//
	// 1361234****,1371234****
	ErrorPhone *string `json:"ErrorPhone,omitempty" xml:"ErrorPhone,omitempty"`
	// The DU upper limit of the Serverless instance, with values being: 2, 4, 8, 16.
	//
	// Currently, this feature is not supported, please do not pass in parameters.
	//
	// example:
	//
	// 16
	MaxDu *float64 `json:"MaxDu,omitempty" xml:"MaxDu,omitempty"`
	// The lower limit of DU for Serverless instances, with values being: 1, 2, 4, 8, 16.
	//
	// This feature is currently not supported, please do not pass in parameters.
	//
	// example:
	//
	// 1
	MinDu *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	// The ID of the region in which the Data Transmission Service (DTS) instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The reserved parameter of DTS. The value must be a JSON string. You can specify this parameter to add more configurations of the source or destination database to the DTS task. For example, you can specify the data storage format of the destination Kafka database and the ID of the CEN instance. For more information, see [MigrationReserved](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {      "srcInstanceId": "cen-9kqshqum*******"  }
	Reserve *string `json:"Reserve,omitempty" xml:"Reserve,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Name of the database to be subscribed.
	//
	// example:
	//
	// dtstestdata
	SourceEndpointDatabaseName *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	// The engine of the source database. Valid values: **MySQL**, **PostgreSQL**, and **Oracle**.
	//
	// >  If the source database is a self-managed database, you must specify this parameter.
	//
	// example:
	//
	// PostgreSQL
	SourceEndpointEngineName *string `json:"SourceEndpointEngineName,omitempty" xml:"SourceEndpointEngineName,omitempty"`
	// The endpoint of the source database.
	//
	// >  This parameter is required only when the source database is a self-managed database.
	//
	// example:
	//
	// 172.16.8*.***
	SourceEndpointIP *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	// The ID of the source database.
	//
	// >  This parameter is required only when the source database is an ApsaraDB RDS for MySQL instance, a PolarDB-X 1.0 instance, or a PolarDB for MySQL cluster.
	//
	// example:
	//
	// rm-bp1zc3iyqe3qw****
	SourceEndpointInstanceID *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	// The type of the source database. Valid values:
	//
	// 	- **RDS**: ApsaraDB RDS for MySQL instance
	//
	// 	- **PolarDB**: PolarDB for MySQL cluster
	//
	// 	- **DRDS**: PolarDB-X 1.0 instance
	//
	// 	- **LocalInstance**: self-managed database with a public IP address
	//
	// 	- **ECS**: self-managed database hosted on an Elastic Compute Service (ECS) instance
	//
	// 	- **Express**: self-managed database connected over Express Connect
	//
	// 	- **CEN**: self-managed database connected over Cloud Enterprise Network (CEN)
	//
	// 	- **dg**: self-managed database connected over Database Gateway
	//
	// example:
	//
	// RDS
	SourceEndpointInstanceType *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	// The system ID (SID) of the Oracle database.
	//
	// >  This parameter is required only when the source database is a self-managed Oracle database and is not deployed in the Real Application Clusters (RAC) architecture.
	//
	// example:
	//
	// testsid
	SourceEndpointOracleSID *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	// The ID of the Alibaba Cloud account to which the source database belongs.
	//
	// >  This parameter is required only when you track data changes across different Alibaba Cloud accounts.
	//
	// example:
	//
	// 140692647406****
	SourceEndpointOwnerID *string `json:"SourceEndpointOwnerID,omitempty" xml:"SourceEndpointOwnerID,omitempty"`
	// The password of the account that is used to connect to the source database.
	//
	// example:
	//
	// Test123456
	SourceEndpointPassword *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	// The service port number of the source database.
	//
	// >  This parameter is required only when the source database is a self-managed database.
	//
	// example:
	//
	// 3306
	SourceEndpointPort *string `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	// The ID of the region in which the source database resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// >  If the source database is a self-managed database with a public IP address, you can set the value of this parameter to **cn-hangzhou*	- or the ID of the region that is closest to the region in which the self-managed database resides.
	//
	// example:
	//
	// cn-hangzhou
	SourceEndpointRegion *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	// The RAM role that is authorized to access the source database. This parameter is required if the source database does not belong to the Alibaba Cloud account that you use to configure the change tracking task. In this case, you must authorize the Alibaba Cloud account to access the source database by using a RAM role.
	//
	// >  For more information about the permissions that are required for the RAM role and how to grant the permissions to the RAM role, see [Configure RAM authorization for cross-account data migration and synchronization](https://help.aliyun.com/document_detail/48468.html).
	//
	// example:
	//
	// ram-for-dts
	SourceEndpointRole *string `json:"SourceEndpointRole,omitempty" xml:"SourceEndpointRole,omitempty"`
	// The username of the account that is used to connect to the source database.
	//
	// >  The permissions that are required for the database account vary with the change tracking scenario. For more information, see [Prepare the source database account for change tracking](https://help.aliyun.com/document_detail/212653.html).
	//
	// example:
	//
	// dtstest
	SourceEndpointUserName *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	// The path of the certificate authority (CA) certificate that is used if the connection to the source database is encrypted by using the SSL protocol.
	//
	// >  This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcCaCertificateOssUrl *string `json:"SrcCaCertificateOssUrl,omitempty" xml:"SrcCaCertificateOssUrl,omitempty"`
	// The key of the CA certificate that is used if the connection to the source database is encrypted by using the SSL protocol.
	//
	// >  This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcCaCertificatePassword *string `json:"SrcCaCertificatePassword,omitempty" xml:"SrcCaCertificatePassword,omitempty"`
	// The path to the client certificate that is used if the connection to the source database is encrypted by using the SSL protocol.
	//
	// >  This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcClientCertOssUrl *string `json:"SrcClientCertOssUrl,omitempty" xml:"SrcClientCertOssUrl,omitempty"`
	// The path to the private key of the client certificate that is used if the connection to the source database is encrypted by using the SSL protocol.
	//
	// >  This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcClientKeyOssUrl *string `json:"SrcClientKeyOssUrl,omitempty" xml:"SrcClientKeyOssUrl,omitempty"`
	// The password of the private key of the client certificate that is used if the connection to the source database is encrypted by using the SSL protocol.
	//
	// >  This feature is not supported. Do not specify this parameter.
	//
	// example:
	//
	// ****
	SrcClientPassword *string `json:"SrcClientPassword,omitempty" xml:"SrcClientPassword,omitempty"`
	// Specifies whether to track DDL statements. Default value: true. Valid values:
	//
	// 	- **true**: tracks DDL statements.
	//
	// 	- **false**: does not track DDL statements.
	//
	// example:
	//
	// true
	SubscriptionDataTypeDDL *bool `json:"SubscriptionDataTypeDDL,omitempty" xml:"SubscriptionDataTypeDDL,omitempty"`
	// Specifies whether to track DML statements. Default value: true. Valid values:
	//
	// 	- **true**: tracks DML statements.
	//
	// 	- **false**: does not track DML statements.
	//
	// example:
	//
	// true
	SubscriptionDataTypeDML *bool `json:"SubscriptionDataTypeDML,omitempty" xml:"SubscriptionDataTypeDML,omitempty"`
	// The network type of the change tracking task. Set the value to **vpc**. A value of vpc indicates the Virtual Private Cloud (VPC) network type.
	//
	// >
	//
	// 	- To use the new version of the change tracking feature, you must specify the SubscriptionInstanceNetworkType parameter. You must also specify the **SubscriptionInstanceVPCId*	- and **SubscriptionInstanceVSwitchID*	- parameters. If you do not specify the SubscriptionInstanceNetworkType parameter, the previous version of the change tracking feature is used.
	//
	// 	- The previous version of the change tracking feature supports self-managed MySQL databases, ApsaraDB RDS for MySQL instances, and PolarDB-X 1.0 instances. The new version of the change tracking feature supports self-managed MySQL databases, ApsaraDB RDS for MySQL instances, PolarDB for MySQL clusters, and Oracle databases.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc
	SubscriptionInstanceNetworkType *string `json:"SubscriptionInstanceNetworkType,omitempty" xml:"SubscriptionInstanceNetworkType,omitempty"`
	// The ID of the VPC in which the change tracking instance is deployed.
	//
	// >  This parameter is required only when the **SubscriptionInstanceNetworkType*	- parameter is set to **vpc**.
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	SubscriptionInstanceVPCId *string `json:"SubscriptionInstanceVPCId,omitempty" xml:"SubscriptionInstanceVPCId,omitempty"`
	// The ID of the vSwitch in the specified VPC.
	//
	// >  This parameter is required only when the **SubscriptionInstanceNetworkType*	- parameter is set to **vpc**.
	//
	// example:
	//
	// vsw-bp10df3mxae6lpmku****
	SubscriptionInstanceVSwitchId *string `json:"SubscriptionInstanceVSwitchId,omitempty" xml:"SubscriptionInstanceVSwitchId,omitempty"`
}

func (s ConfigureSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionRequest) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *ConfigureSubscriptionRequest) GetDbList() *string {
	return s.DbList
}

func (s *ConfigureSubscriptionRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *ConfigureSubscriptionRequest) GetDelayNotice() *bool {
	return s.DelayNotice
}

func (s *ConfigureSubscriptionRequest) GetDelayPhone() *string {
	return s.DelayPhone
}

func (s *ConfigureSubscriptionRequest) GetDelayRuleTime() *int64 {
	return s.DelayRuleTime
}

func (s *ConfigureSubscriptionRequest) GetDtsBisLabel() *string {
	return s.DtsBisLabel
}

func (s *ConfigureSubscriptionRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ConfigureSubscriptionRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ConfigureSubscriptionRequest) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *ConfigureSubscriptionRequest) GetErrorNotice() *bool {
	return s.ErrorNotice
}

func (s *ConfigureSubscriptionRequest) GetErrorPhone() *string {
	return s.ErrorPhone
}

func (s *ConfigureSubscriptionRequest) GetMaxDu() *float64 {
	return s.MaxDu
}

func (s *ConfigureSubscriptionRequest) GetMinDu() *float64 {
	return s.MinDu
}

func (s *ConfigureSubscriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureSubscriptionRequest) GetReserve() *string {
	return s.Reserve
}

func (s *ConfigureSubscriptionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigureSubscriptionRequest) GetSourceEndpointDatabaseName() *string {
	return s.SourceEndpointDatabaseName
}

func (s *ConfigureSubscriptionRequest) GetSourceEndpointEngineName() *string {
	return s.SourceEndpointEngineName
}

func (s *ConfigureSubscriptionRequest) GetSourceEndpointIP() *string {
	return s.SourceEndpointIP
}

func (s *ConfigureSubscriptionRequest) GetSourceEndpointInstanceID() *string {
	return s.SourceEndpointInstanceID
}

func (s *ConfigureSubscriptionRequest) GetSourceEndpointInstanceType() *string {
	return s.SourceEndpointInstanceType
}

func (s *ConfigureSubscriptionRequest) GetSourceEndpointOracleSID() *string {
	return s.SourceEndpointOracleSID
}

func (s *ConfigureSubscriptionRequest) GetSourceEndpointOwnerID() *string {
	return s.SourceEndpointOwnerID
}

func (s *ConfigureSubscriptionRequest) GetSourceEndpointPassword() *string {
	return s.SourceEndpointPassword
}

func (s *ConfigureSubscriptionRequest) GetSourceEndpointPort() *string {
	return s.SourceEndpointPort
}

func (s *ConfigureSubscriptionRequest) GetSourceEndpointRegion() *string {
	return s.SourceEndpointRegion
}

func (s *ConfigureSubscriptionRequest) GetSourceEndpointRole() *string {
	return s.SourceEndpointRole
}

func (s *ConfigureSubscriptionRequest) GetSourceEndpointUserName() *string {
	return s.SourceEndpointUserName
}

func (s *ConfigureSubscriptionRequest) GetSrcCaCertificateOssUrl() *string {
	return s.SrcCaCertificateOssUrl
}

func (s *ConfigureSubscriptionRequest) GetSrcCaCertificatePassword() *string {
	return s.SrcCaCertificatePassword
}

func (s *ConfigureSubscriptionRequest) GetSrcClientCertOssUrl() *string {
	return s.SrcClientCertOssUrl
}

func (s *ConfigureSubscriptionRequest) GetSrcClientKeyOssUrl() *string {
	return s.SrcClientKeyOssUrl
}

func (s *ConfigureSubscriptionRequest) GetSrcClientPassword() *string {
	return s.SrcClientPassword
}

func (s *ConfigureSubscriptionRequest) GetSubscriptionDataTypeDDL() *bool {
	return s.SubscriptionDataTypeDDL
}

func (s *ConfigureSubscriptionRequest) GetSubscriptionDataTypeDML() *bool {
	return s.SubscriptionDataTypeDML
}

func (s *ConfigureSubscriptionRequest) GetSubscriptionInstanceNetworkType() *string {
	return s.SubscriptionInstanceNetworkType
}

func (s *ConfigureSubscriptionRequest) GetSubscriptionInstanceVPCId() *string {
	return s.SubscriptionInstanceVPCId
}

func (s *ConfigureSubscriptionRequest) GetSubscriptionInstanceVSwitchId() *string {
	return s.SubscriptionInstanceVSwitchId
}

func (s *ConfigureSubscriptionRequest) SetCheckpoint(v string) *ConfigureSubscriptionRequest {
	s.Checkpoint = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDbList(v string) *ConfigureSubscriptionRequest {
	s.DbList = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDedicatedClusterId(v string) *ConfigureSubscriptionRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDelayNotice(v bool) *ConfigureSubscriptionRequest {
	s.DelayNotice = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDelayPhone(v string) *ConfigureSubscriptionRequest {
	s.DelayPhone = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDelayRuleTime(v int64) *ConfigureSubscriptionRequest {
	s.DelayRuleTime = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDtsBisLabel(v string) *ConfigureSubscriptionRequest {
	s.DtsBisLabel = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDtsInstanceId(v string) *ConfigureSubscriptionRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDtsJobId(v string) *ConfigureSubscriptionRequest {
	s.DtsJobId = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDtsJobName(v string) *ConfigureSubscriptionRequest {
	s.DtsJobName = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetErrorNotice(v bool) *ConfigureSubscriptionRequest {
	s.ErrorNotice = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetErrorPhone(v string) *ConfigureSubscriptionRequest {
	s.ErrorPhone = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetMaxDu(v float64) *ConfigureSubscriptionRequest {
	s.MaxDu = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetMinDu(v float64) *ConfigureSubscriptionRequest {
	s.MinDu = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetRegionId(v string) *ConfigureSubscriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetReserve(v string) *ConfigureSubscriptionRequest {
	s.Reserve = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetResourceGroupId(v string) *ConfigureSubscriptionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointDatabaseName(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointEngineName(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointEngineName = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointIP(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointInstanceID(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointInstanceType(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointOracleSID(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointOwnerID(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointOwnerID = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointPassword(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointPort(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointRegion(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointRole(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointRole = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointUserName(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSrcCaCertificateOssUrl(v string) *ConfigureSubscriptionRequest {
	s.SrcCaCertificateOssUrl = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSrcCaCertificatePassword(v string) *ConfigureSubscriptionRequest {
	s.SrcCaCertificatePassword = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSrcClientCertOssUrl(v string) *ConfigureSubscriptionRequest {
	s.SrcClientCertOssUrl = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSrcClientKeyOssUrl(v string) *ConfigureSubscriptionRequest {
	s.SrcClientKeyOssUrl = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSrcClientPassword(v string) *ConfigureSubscriptionRequest {
	s.SrcClientPassword = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSubscriptionDataTypeDDL(v bool) *ConfigureSubscriptionRequest {
	s.SubscriptionDataTypeDDL = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSubscriptionDataTypeDML(v bool) *ConfigureSubscriptionRequest {
	s.SubscriptionDataTypeDML = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSubscriptionInstanceNetworkType(v string) *ConfigureSubscriptionRequest {
	s.SubscriptionInstanceNetworkType = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSubscriptionInstanceVPCId(v string) *ConfigureSubscriptionRequest {
	s.SubscriptionInstanceVPCId = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSubscriptionInstanceVSwitchId(v string) *ConfigureSubscriptionRequest {
	s.SubscriptionInstanceVSwitchId = &v
	return s
}

func (s *ConfigureSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
