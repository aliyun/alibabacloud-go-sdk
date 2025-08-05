// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsJobDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeDtsJobDetailResponseBody
	GetAppName() *string
	SetBeginTimestamp(v string) *DescribeDtsJobDetailResponseBody
	GetBeginTimestamp() *string
	SetBinlog(v string) *DescribeDtsJobDetailResponseBody
	GetBinlog() *string
	SetBinlogSite(v string) *DescribeDtsJobDetailResponseBody
	GetBinlogSite() *string
	SetBinlogTime(v string) *DescribeDtsJobDetailResponseBody
	GetBinlogTime() *string
	SetBootTime(v string) *DescribeDtsJobDetailResponseBody
	GetBootTime() *string
	SetCheckpoint(v int64) *DescribeDtsJobDetailResponseBody
	GetCheckpoint() *int64
	SetCode(v int32) *DescribeDtsJobDetailResponseBody
	GetCode() *int32
	SetConsumptionCheckpoint(v string) *DescribeDtsJobDetailResponseBody
	GetConsumptionCheckpoint() *string
	SetConsumptionClient(v string) *DescribeDtsJobDetailResponseBody
	GetConsumptionClient() *string
	SetCreateTime(v string) *DescribeDtsJobDetailResponseBody
	GetCreateTime() *string
	SetDataDeliveryChannelInfo(v *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) *DescribeDtsJobDetailResponseBody
	GetDataDeliveryChannelInfo() *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo
	SetDataSynchronizationStatus(v *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) *DescribeDtsJobDetailResponseBody
	GetDataSynchronizationStatus() *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus
	SetDatabaseCount(v int32) *DescribeDtsJobDetailResponseBody
	GetDatabaseCount() *int32
	SetDbObject(v string) *DescribeDtsJobDetailResponseBody
	GetDbObject() *string
	SetDedicatedClusterId(v string) *DescribeDtsJobDetailResponseBody
	GetDedicatedClusterId() *string
	SetDelay(v int64) *DescribeDtsJobDetailResponseBody
	GetDelay() *int64
	SetDemoJob(v bool) *DescribeDtsJobDetailResponseBody
	GetDemoJob() *bool
	SetDestNetType(v string) *DescribeDtsJobDetailResponseBody
	GetDestNetType() *string
	SetDestinationEndpoint(v *DescribeDtsJobDetailResponseBodyDestinationEndpoint) *DescribeDtsJobDetailResponseBody
	GetDestinationEndpoint() *DescribeDtsJobDetailResponseBodyDestinationEndpoint
	SetDtsBisLabel(v string) *DescribeDtsJobDetailResponseBody
	GetDtsBisLabel() *string
	SetDtsInstanceID(v string) *DescribeDtsJobDetailResponseBody
	GetDtsInstanceID() *string
	SetDtsJobClass(v string) *DescribeDtsJobDetailResponseBody
	GetDtsJobClass() *string
	SetDtsJobDirection(v string) *DescribeDtsJobDetailResponseBody
	GetDtsJobDirection() *string
	SetDtsJobId(v string) *DescribeDtsJobDetailResponseBody
	GetDtsJobId() *string
	SetDtsJobName(v string) *DescribeDtsJobDetailResponseBody
	GetDtsJobName() *string
	SetDynamicMessage(v string) *DescribeDtsJobDetailResponseBody
	GetDynamicMessage() *string
	SetEndTimestamp(v string) *DescribeDtsJobDetailResponseBody
	GetEndTimestamp() *string
	SetErrCode(v string) *DescribeDtsJobDetailResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDtsJobDetailResponseBody
	GetErrMessage() *string
	SetErrorMessage(v string) *DescribeDtsJobDetailResponseBody
	GetErrorMessage() *string
	SetEtlCalculator(v string) *DescribeDtsJobDetailResponseBody
	GetEtlCalculator() *string
	SetExpireTime(v string) *DescribeDtsJobDetailResponseBody
	GetExpireTime() *string
	SetFinishTime(v string) *DescribeDtsJobDetailResponseBody
	GetFinishTime() *string
	SetGroupId(v string) *DescribeDtsJobDetailResponseBody
	GetGroupId() *string
	SetHttpStatusCode(v int32) *DescribeDtsJobDetailResponseBody
	GetHttpStatusCode() *int32
	SetInitCheckpoint(v string) *DescribeDtsJobDetailResponseBody
	GetInitCheckpoint() *string
	SetJobType(v string) *DescribeDtsJobDetailResponseBody
	GetJobType() *string
	SetLastUpdateTime(v string) *DescribeDtsJobDetailResponseBody
	GetLastUpdateTime() *string
	SetMaxDu(v float64) *DescribeDtsJobDetailResponseBody
	GetMaxDu() *float64
	SetMigrationMode(v *DescribeDtsJobDetailResponseBodyMigrationMode) *DescribeDtsJobDetailResponseBody
	GetMigrationMode() *DescribeDtsJobDetailResponseBodyMigrationMode
	SetMinDu(v float64) *DescribeDtsJobDetailResponseBody
	GetMinDu() *float64
	SetPayType(v string) *DescribeDtsJobDetailResponseBody
	GetPayType() *string
	SetRequestId(v string) *DescribeDtsJobDetailResponseBody
	GetRequestId() *string
	SetReserved(v string) *DescribeDtsJobDetailResponseBody
	GetReserved() *string
	SetResourceGroupDisplayName(v string) *DescribeDtsJobDetailResponseBody
	GetResourceGroupDisplayName() *string
	SetResourceGroupId(v string) *DescribeDtsJobDetailResponseBody
	GetResourceGroupId() *string
	SetRetryState(v *DescribeDtsJobDetailResponseBodyRetryState) *DescribeDtsJobDetailResponseBody
	GetRetryState() *DescribeDtsJobDetailResponseBodyRetryState
	SetSourceEndpoint(v *DescribeDtsJobDetailResponseBodySourceEndpoint) *DescribeDtsJobDetailResponseBody
	GetSourceEndpoint() *DescribeDtsJobDetailResponseBodySourceEndpoint
	SetStatus(v string) *DescribeDtsJobDetailResponseBody
	GetStatus() *string
	SetSubDistributedJob(v []*DescribeDtsJobDetailResponseBodySubDistributedJob) *DescribeDtsJobDetailResponseBody
	GetSubDistributedJob() []*DescribeDtsJobDetailResponseBodySubDistributedJob
	SetSubSyncJob(v []*DescribeDtsJobDetailResponseBodySubSyncJob) *DescribeDtsJobDetailResponseBody
	GetSubSyncJob() []*DescribeDtsJobDetailResponseBodySubSyncJob
	SetSubscribeTopic(v string) *DescribeDtsJobDetailResponseBody
	GetSubscribeTopic() *string
	SetSubscriptionDataType(v *DescribeDtsJobDetailResponseBodySubscriptionDataType) *DescribeDtsJobDetailResponseBody
	GetSubscriptionDataType() *DescribeDtsJobDetailResponseBodySubscriptionDataType
	SetSubscriptionHost(v *DescribeDtsJobDetailResponseBodySubscriptionHost) *DescribeDtsJobDetailResponseBody
	GetSubscriptionHost() *DescribeDtsJobDetailResponseBodySubscriptionHost
	SetSuccess(v bool) *DescribeDtsJobDetailResponseBody
	GetSuccess() *bool
	SetSynchronizationDirection(v string) *DescribeDtsJobDetailResponseBody
	GetSynchronizationDirection() *string
	SetTaskType(v string) *DescribeDtsJobDetailResponseBody
	GetTaskType() *string
}

type DescribeDtsJobDetailResponseBody struct {
	// Indicates whether the new change tracking feature is used.
	//
	// > This parameter is returned only for change tracking instances of the new version.
	//
	// example:
	//
	// new
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The start of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-15T08:25:34Z
	BeginTimestamp *string `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	// The binary logs.
	//
	// example:
	//
	// ****
	Binlog *string `json:"Binlog,omitempty" xml:"Binlog,omitempty"`
	// The current offset.
	//
	// example:
	//
	// 156629109****
	BinlogSite *string `json:"BinlogSite,omitempty" xml:"BinlogSite,omitempty"`
	// The offset range.
	//
	// example:
	//
	// ****
	BinlogTime *string `json:"BinlogTime,omitempty" xml:"BinlogTime,omitempty"`
	// The time when the task was started. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-30T03:36:11.000
	BootTime *string `json:"BootTime,omitempty" xml:"BootTime,omitempty"`
	// The start offset of incremental data migration or data synchronization. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1616405159
	Checkpoint *int64 `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The consumption checkpoint of the change tracking instance. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-23T07:30:31Z
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	// The downstream client information in the following format: \\<IP address of the downstream client>:\\<Random ID generated by DTS>.
	//
	// example:
	//
	// 114.***.***.**:dts********
	ConsumptionClient *string `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about the data shipping channel.
	DataDeliveryChannelInfo *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo `json:"DataDeliveryChannelInfo,omitempty" xml:"DataDeliveryChannelInfo,omitempty" type:"Struct"`
	// The state of incremental data migration or synchronization.
	DataSynchronizationStatus *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The number of ApsaraDB RDS for MySQL instances that are attached to the source PolarDB-X 1.0 instance.
	//
	// example:
	//
	// 2
	DatabaseCount *int32 `json:"DatabaseCount,omitempty" xml:"DatabaseCount,omitempty"`
	// The objects of the data migration, data synchronization, or change tracking task. For more information, see [Objects of DTS tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// example:
	//
	// {\\"dtstestdata\\":{\\"all\\":true,\\"name\\":\\"dtstestdata\\",\\"state\\":\\"normal\\"}}
	DbObject *string `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	// The dedicated cluster ID.
	//
	// example:
	//
	// dtsxxxxx
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The latency of incremental data migration or synchronization. Unit: milliseconds.
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// Indicates whether the task is a subtask. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DemoJob *bool `json:"DemoJob,omitempty" xml:"DemoJob,omitempty"`
	// The network type of the consumer client. Valid values:
	//
	// 	- **CLASSIC**: classic network.
	//
	// 	- **VPC**: virtual private cloud (VPC).
	//
	// example:
	//
	// VPC
	DestNetType *string `json:"DestNetType,omitempty" xml:"DestNetType,omitempty"`
	// The connection settings of the destination instance.
	DestinationEndpoint *DescribeDtsJobDetailResponseBodyDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
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
	// The ID of the data migration, data synchronization, or change tracking instance.
	//
	// example:
	//
	// dtsi03e3zty16i****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The instance class.
	//
	// > For more information about the description and test performance of each instance class, see [Specifications of data migration instances](https://help.aliyun.com/document_detail/26606.html) and [Specifications of data synchronization instances](https://help.aliyun.com/document_detail/26605.html).
	//
	// example:
	//
	// xlarge
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// > This parameter is returned only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	DtsJobDirection *string `json:"DtsJobDirection,omitempty" xml:"DtsJobDirection,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking task.
	//
	// example:
	//
	// i03e3zty16i****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The name of the data migration, data synchronization, or change tracking task.
	//
	// example:
	//
	// api_test
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The dynamic part in the error message. The value of this parameter is used to replace **%s*	- in the value of **ErrMessage**.
	//
	// > For example, if the return value of **ErrMessage*	- is **The Value of Input Parameter %s is not valid*	- and the value of **DynamicMessage*	- is **DtsJobId**, the specified value of **DtsJobId*	- is invalid.
	//
	// example:
	//
	// DtsJobId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The end of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-26T14:03:21Z
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by [com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException:Could not create connection to database server. Attempted reconnect 3 times. Giving up.][com.mysql.jdbc.exceptions.jdbc4.CommunicationsException:Communications link failure\\n\\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.][java.net.ConnectException:Connection timed out (Connection timed out)] About more information in [https://yq.aliyun.com/articles/499178].
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The operator information of the ETL task.
	//
	// > This parameter is returned only if you query the details of an ETL task.
	//
	// example:
	//
	// { 	"cells ": [{\\"shape\\":\\"edge\\",\\"attrs\\":{\\"line\\":{\\"stroke\\":\\"#b1b1b1\\",\\"strokeWidth\\":1,\\"targetMarker\\":{\\"name\\":\\"block\\",\\"args\\":{\\"size\\":\\"8\\"}},\\"strokeDasharray\\":\\"\\"}},\\"id\\":\\"cd1ec473-f9b9-4e9b-a742-ac23f442****\\",\\"source\\":{\\"cell\\":\\"8b261182-bfab-4803-ad8e-6bb08e3e****\\",\\"port\\":\\"out1\\"},\\"target\\":{\\"cell\\":\\"b36770df-f48c-4d6b-9644-54c5e924****\\",\\"port\\":\\"in1\\"},\\"zIndex\\":7 	}] }
	EtlCalculator *string `json:"EtlCalculator,omitempty" xml:"EtlCalculator,omitempty"`
	// The time when the instance expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// > This parameter is returned only if the return value of **PayType*	- is **PrePaid**.
	//
	// example:
	//
	// 2023-06-16T08:01:19Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the task was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-16T10:34:17Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InitCheckpoint *string `json:"InitCheckpoint,omitempty" xml:"InitCheckpoint,omitempty"`
	// The type of the DTS task. Valid values:
	//
	// 	- **sync**: a data synchronization task.
	//
	// 	- **subSync**: a subtask generated when the objects to be synchronized are modified.
	//
	// > In most cases, this parameter is returned together with **TaskType**.
	//
	// example:
	//
	// sync
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The timestamp when the task was last updated.
	//
	// example:
	//
	// 156629109****
	LastUpdateTime *string `json:"LastUpdateTime,omitempty" xml:"LastUpdateTime,omitempty"`
	// The maximum number of DUs.
	//
	// > This parameter is supported only for serverless instances.
	//
	// example:
	//
	// 16
	MaxDu *float64 `json:"MaxDu,omitempty" xml:"MaxDu,omitempty"`
	// The migration types or initial synchronization types.
	MigrationMode *DescribeDtsJobDetailResponseBodyMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// The minimum number of DUs.
	//
	// > This parameter is supported only for serverless instances.
	//
	// example:
	//
	// 1
	MinDu *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	// The billing method of the DTS instance. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 29207299-7C41-493A-BA4F-2FAC5DE4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet specific requirements, such as whether to automatically start a precheck. For more information, see [MigrationReserved](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {\\"srcHostPorts\\":\\"\\",\\"whitelist.dms.online.ddl.enable\\":false,\\"filterDDL\\":false,\\"sqlparser.dms.original.ddl\\":true,\\"srcOracleType\\":\\"sid\\",\\"maxRetryTime\\":43200,\\"destSSL\\":\\"0\\",\\"destOracleType\\":\\"sid\\",\\"srcSSL\\":\\"0\\",\\"dbListCaseChangeMode\\":\\"default\\",\\"SourceEngineVersion\\":\\"8.0.18\\",\\"srcNetType\\":\\"VPC\\",\\"destNetType\\":\\"VPC\\",\\"srcVpcNetMappingInst\\":\\"172.16.1**.**:10803\\",\\"destVpcNetMappingInst\\":\\"172.16.1**.**:11077\\",\\"useJobTask\\":\\"1\\"}
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The resource group name.
	//
	// example:
	//
	// default resource group
	ResourceGroupDisplayName *string `json:"ResourceGroupDisplayName,omitempty" xml:"ResourceGroupDisplayName,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The information about the retries performed by DTS due to an exception.
	RetryState *DescribeDtsJobDetailResponseBodyRetryState `json:"RetryState,omitempty" xml:"RetryState,omitempty" type:"Struct"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeDtsJobDetailResponseBodySourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The state of the data migration or synchronization task. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **NotConfigured**: The task is not configured.
	//
	// 	- **Prechecking**: The task is in precheck.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **PreCheckPass**: The task passed the precheck.
	//
	// 	- **Initializing**: Initial data synchronization is in progress.
	//
	// 	- **InitializeFailed**: Initial data synchronization failed.
	//
	// 	- **synchronizing**: Data synchronization is in progress.
	//
	// 	- **Migrating**: Data migration is in progress.
	//
	// 	- **Failed**: Data synchronization failed.
	//
	// 	- **MigrationFailed**: Data migration failed.
	//
	// 	- **Suspending**: The task is paused.
	//
	// 	- **Modifying**: The objects of the task are being modified.
	//
	// 	- **Retrying**: The task is being retried.
	//
	// 	- **Upgrade**: The task is being upgraded.
	//
	// 	- **Downgrade**: The task is being downgraded.
	//
	// 	- **Locked**: The task is locked.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Migrating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the subtasks in the current distributed task. If the DTS task is not a distributed task, the value of this parameter is null.
	//
	// > This parameter is available only if the DTS task is a data synchronization task.
	SubDistributedJob []*DescribeDtsJobDetailResponseBodySubDistributedJob `json:"SubDistributedJob,omitempty" xml:"SubDistributedJob,omitempty" type:"Repeated"`
	// The information about the subtasks in the current data synchronization task.
	SubSyncJob []*DescribeDtsJobDetailResponseBodySubSyncJob `json:"SubSyncJob,omitempty" xml:"SubSyncJob,omitempty" type:"Repeated"`
	// The topic of the change tracking instance.
	//
	// > This parameter is returned only if your change tracking instances are of the new version and you have called the [CreateConsumerGroup](https://help.aliyun.com/document_detail/122863.html) operation to create a consumer group.
	//
	// example:
	//
	// cn_hangzhou_rm_bp1162kryivb8****_dtstest_version2
	SubscribeTopic *string `json:"SubscribeTopic,omitempty" xml:"SubscribeTopic,omitempty"`
	// The type of data for change tracking.
	SubscriptionDataType *DescribeDtsJobDetailResponseBodySubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	// The endpoint of the change tracking instance.
	SubscriptionHost *DescribeDtsJobDetailResponseBodySubscriptionHost `json:"SubscriptionHost,omitempty" xml:"SubscriptionHost,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// 	- The value **Reverse*	- takes effect only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The type of the task.
	//
	// > In most cases, this parameter is returned together with **JobType**.
	//
	// example:
	//
	// Distributed_xxx
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDtsJobDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *DescribeDtsJobDetailResponseBody) GetBeginTimestamp() *string {
	return s.BeginTimestamp
}

func (s *DescribeDtsJobDetailResponseBody) GetBinlog() *string {
	return s.Binlog
}

func (s *DescribeDtsJobDetailResponseBody) GetBinlogSite() *string {
	return s.BinlogSite
}

func (s *DescribeDtsJobDetailResponseBody) GetBinlogTime() *string {
	return s.BinlogTime
}

func (s *DescribeDtsJobDetailResponseBody) GetBootTime() *string {
	return s.BootTime
}

func (s *DescribeDtsJobDetailResponseBody) GetCheckpoint() *int64 {
	return s.Checkpoint
}

func (s *DescribeDtsJobDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeDtsJobDetailResponseBody) GetConsumptionCheckpoint() *string {
	return s.ConsumptionCheckpoint
}

func (s *DescribeDtsJobDetailResponseBody) GetConsumptionClient() *string {
	return s.ConsumptionClient
}

func (s *DescribeDtsJobDetailResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDtsJobDetailResponseBody) GetDataDeliveryChannelInfo() *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo {
	return s.DataDeliveryChannelInfo
}

func (s *DescribeDtsJobDetailResponseBody) GetDataSynchronizationStatus() *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeDtsJobDetailResponseBody) GetDatabaseCount() *int32 {
	return s.DatabaseCount
}

func (s *DescribeDtsJobDetailResponseBody) GetDbObject() *string {
	return s.DbObject
}

func (s *DescribeDtsJobDetailResponseBody) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeDtsJobDetailResponseBody) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeDtsJobDetailResponseBody) GetDemoJob() *bool {
	return s.DemoJob
}

func (s *DescribeDtsJobDetailResponseBody) GetDestNetType() *string {
	return s.DestNetType
}

func (s *DescribeDtsJobDetailResponseBody) GetDestinationEndpoint() *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeDtsJobDetailResponseBody) GetDtsBisLabel() *string {
	return s.DtsBisLabel
}

func (s *DescribeDtsJobDetailResponseBody) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeDtsJobDetailResponseBody) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeDtsJobDetailResponseBody) GetDtsJobDirection() *string {
	return s.DtsJobDirection
}

func (s *DescribeDtsJobDetailResponseBody) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsJobDetailResponseBody) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeDtsJobDetailResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDtsJobDetailResponseBody) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeDtsJobDetailResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDtsJobDetailResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDtsJobDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBody) GetEtlCalculator() *string {
	return s.EtlCalculator
}

func (s *DescribeDtsJobDetailResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDtsJobDetailResponseBody) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeDtsJobDetailResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDtsJobDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDtsJobDetailResponseBody) GetInitCheckpoint() *string {
	return s.InitCheckpoint
}

func (s *DescribeDtsJobDetailResponseBody) GetJobType() *string {
	return s.JobType
}

func (s *DescribeDtsJobDetailResponseBody) GetLastUpdateTime() *string {
	return s.LastUpdateTime
}

func (s *DescribeDtsJobDetailResponseBody) GetMaxDu() *float64 {
	return s.MaxDu
}

func (s *DescribeDtsJobDetailResponseBody) GetMigrationMode() *DescribeDtsJobDetailResponseBodyMigrationMode {
	return s.MigrationMode
}

func (s *DescribeDtsJobDetailResponseBody) GetMinDu() *float64 {
	return s.MinDu
}

func (s *DescribeDtsJobDetailResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDtsJobDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDtsJobDetailResponseBody) GetReserved() *string {
	return s.Reserved
}

func (s *DescribeDtsJobDetailResponseBody) GetResourceGroupDisplayName() *string {
	return s.ResourceGroupDisplayName
}

func (s *DescribeDtsJobDetailResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDtsJobDetailResponseBody) GetRetryState() *DescribeDtsJobDetailResponseBodyRetryState {
	return s.RetryState
}

func (s *DescribeDtsJobDetailResponseBody) GetSourceEndpoint() *DescribeDtsJobDetailResponseBodySourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeDtsJobDetailResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBody) GetSubDistributedJob() []*DescribeDtsJobDetailResponseBodySubDistributedJob {
	return s.SubDistributedJob
}

func (s *DescribeDtsJobDetailResponseBody) GetSubSyncJob() []*DescribeDtsJobDetailResponseBodySubSyncJob {
	return s.SubSyncJob
}

func (s *DescribeDtsJobDetailResponseBody) GetSubscribeTopic() *string {
	return s.SubscribeTopic
}

func (s *DescribeDtsJobDetailResponseBody) GetSubscriptionDataType() *DescribeDtsJobDetailResponseBodySubscriptionDataType {
	return s.SubscriptionDataType
}

func (s *DescribeDtsJobDetailResponseBody) GetSubscriptionHost() *DescribeDtsJobDetailResponseBodySubscriptionHost {
	return s.SubscriptionHost
}

func (s *DescribeDtsJobDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDtsJobDetailResponseBody) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeDtsJobDetailResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeDtsJobDetailResponseBody) SetAppName(v string) *DescribeDtsJobDetailResponseBody {
	s.AppName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetBeginTimestamp(v string) *DescribeDtsJobDetailResponseBody {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetBinlog(v string) *DescribeDtsJobDetailResponseBody {
	s.Binlog = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetBinlogSite(v string) *DescribeDtsJobDetailResponseBody {
	s.BinlogSite = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetBinlogTime(v string) *DescribeDtsJobDetailResponseBody {
	s.BinlogTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetBootTime(v string) *DescribeDtsJobDetailResponseBody {
	s.BootTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetCheckpoint(v int64) *DescribeDtsJobDetailResponseBody {
	s.Checkpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetCode(v int32) *DescribeDtsJobDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetConsumptionCheckpoint(v string) *DescribeDtsJobDetailResponseBody {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetConsumptionClient(v string) *DescribeDtsJobDetailResponseBody {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetCreateTime(v string) *DescribeDtsJobDetailResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDataDeliveryChannelInfo(v *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) *DescribeDtsJobDetailResponseBody {
	s.DataDeliveryChannelInfo = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDataSynchronizationStatus(v *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) *DescribeDtsJobDetailResponseBody {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDatabaseCount(v int32) *DescribeDtsJobDetailResponseBody {
	s.DatabaseCount = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDbObject(v string) *DescribeDtsJobDetailResponseBody {
	s.DbObject = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDedicatedClusterId(v string) *DescribeDtsJobDetailResponseBody {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDelay(v int64) *DescribeDtsJobDetailResponseBody {
	s.Delay = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDemoJob(v bool) *DescribeDtsJobDetailResponseBody {
	s.DemoJob = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDestNetType(v string) *DescribeDtsJobDetailResponseBody {
	s.DestNetType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDestinationEndpoint(v *DescribeDtsJobDetailResponseBodyDestinationEndpoint) *DescribeDtsJobDetailResponseBody {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDtsBisLabel(v string) *DescribeDtsJobDetailResponseBody {
	s.DtsBisLabel = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDtsInstanceID(v string) *DescribeDtsJobDetailResponseBody {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDtsJobClass(v string) *DescribeDtsJobDetailResponseBody {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDtsJobDirection(v string) *DescribeDtsJobDetailResponseBody {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDtsJobId(v string) *DescribeDtsJobDetailResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDtsJobName(v string) *DescribeDtsJobDetailResponseBody {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDynamicMessage(v string) *DescribeDtsJobDetailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetEndTimestamp(v string) *DescribeDtsJobDetailResponseBody {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetErrCode(v string) *DescribeDtsJobDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetErrMessage(v string) *DescribeDtsJobDetailResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetEtlCalculator(v string) *DescribeDtsJobDetailResponseBody {
	s.EtlCalculator = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetExpireTime(v string) *DescribeDtsJobDetailResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetFinishTime(v string) *DescribeDtsJobDetailResponseBody {
	s.FinishTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetGroupId(v string) *DescribeDtsJobDetailResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetHttpStatusCode(v int32) *DescribeDtsJobDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetInitCheckpoint(v string) *DescribeDtsJobDetailResponseBody {
	s.InitCheckpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetJobType(v string) *DescribeDtsJobDetailResponseBody {
	s.JobType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetLastUpdateTime(v string) *DescribeDtsJobDetailResponseBody {
	s.LastUpdateTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetMaxDu(v float64) *DescribeDtsJobDetailResponseBody {
	s.MaxDu = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetMigrationMode(v *DescribeDtsJobDetailResponseBodyMigrationMode) *DescribeDtsJobDetailResponseBody {
	s.MigrationMode = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetMinDu(v float64) *DescribeDtsJobDetailResponseBody {
	s.MinDu = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetPayType(v string) *DescribeDtsJobDetailResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetRequestId(v string) *DescribeDtsJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetReserved(v string) *DescribeDtsJobDetailResponseBody {
	s.Reserved = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetResourceGroupDisplayName(v string) *DescribeDtsJobDetailResponseBody {
	s.ResourceGroupDisplayName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetResourceGroupId(v string) *DescribeDtsJobDetailResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetRetryState(v *DescribeDtsJobDetailResponseBodyRetryState) *DescribeDtsJobDetailResponseBody {
	s.RetryState = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSourceEndpoint(v *DescribeDtsJobDetailResponseBodySourceEndpoint) *DescribeDtsJobDetailResponseBody {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetStatus(v string) *DescribeDtsJobDetailResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSubDistributedJob(v []*DescribeDtsJobDetailResponseBodySubDistributedJob) *DescribeDtsJobDetailResponseBody {
	s.SubDistributedJob = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSubSyncJob(v []*DescribeDtsJobDetailResponseBodySubSyncJob) *DescribeDtsJobDetailResponseBody {
	s.SubSyncJob = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSubscribeTopic(v string) *DescribeDtsJobDetailResponseBody {
	s.SubscribeTopic = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSubscriptionDataType(v *DescribeDtsJobDetailResponseBodySubscriptionDataType) *DescribeDtsJobDetailResponseBody {
	s.SubscriptionDataType = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSubscriptionHost(v *DescribeDtsJobDetailResponseBodySubscriptionHost) *DescribeDtsJobDetailResponseBody {
	s.SubscriptionHost = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSuccess(v bool) *DescribeDtsJobDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSynchronizationDirection(v string) *DescribeDtsJobDetailResponseBody {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetTaskType(v string) *DescribeDtsJobDetailResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo struct {
	// The number of partitions of the destination topic.
	//
	// example:
	//
	// 3
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The public endpoint of the data shipping channel.
	//
	// example:
	//
	// dts-****.aliyuncs.com:18***
	PublicDproxyUrl *string `json:"PublicDproxyUrl,omitempty" xml:"PublicDproxyUrl,omitempty"`
	// The region in which the data shipping channel resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The destination topic of the data shipping instance.
	//
	// example:
	//
	// cn_hangzhou_******_data_delivery_version2
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The VPC endpoint of the data shipping channel.
	//
	// example:
	//
	// dts-****.aliyuncs.com:18***
	VpcDproxyUrl *string `json:"VpcDproxyUrl,omitempty" xml:"VpcDproxyUrl,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) GetPublicDproxyUrl() *string {
	return s.PublicDproxyUrl
}

func (s *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) GetTopic() *string {
	return s.Topic
}

func (s *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) GetVpcDproxyUrl() *string {
	return s.VpcDproxyUrl
}

func (s *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) SetPartitionNum(v int32) *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo {
	s.PartitionNum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) SetPublicDproxyUrl(v string) *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo {
	s.PublicDproxyUrl = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) SetRegion(v string) *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) SetTopic(v string) *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo {
	s.Topic = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) SetVpcDproxyUrl(v string) *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo {
	s.VpcDproxyUrl = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDataDeliveryChannelInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodyDataSynchronizationStatus struct {
	// The error message returned if incremental data migration or synchronization failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by [com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException:Could not create connection to database server. Attempted reconnect 3 times. Giving up.][com.mysql.jdbc.exceptions.jdbc4.CommunicationsException:Communications link failure\\n\\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.][java.net.ConnectException:Connection timed out (Connection timed out)] About more information in [https://yq.aliyun.com/articles/499178].
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of incremental data migration or synchronization. Unit: percentage.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of rows and size of data that is synchronized or migrated to the destination table per second during incremental data synchronization or migration.
	//
	// example:
	//
	// 0.00RPS/(0.000MB/s)
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of incremental data migration or synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Suspending**: The task is paused.
	//
	// 	- **Checking**: The task is in precheck.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// 	- **Catched**: The task has no latency.
	//
	// example:
	//
	// Catched
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodyDestinationEndpoint struct {
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// Indicates whether the password can be modified. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CanModifyPassword *bool `json:"CanModifyPassword,omitempty" xml:"CanModifyPassword,omitempty"`
	// The name of the database to which the objects are migrated in the destination instance.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine of the destination instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The destination instance ID.
	//
	// example:
	//
	// rm-bp1imrtn6fq7h****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the destination instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The endpoint of the destination instance.
	//
	// example:
	//
	// 172.16.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the return value of **EngineName*	- of the destination instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the destination instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the destination instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **DISABLE**: SSL encryption is disabled.
	//
	// 	- **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// 	- **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection to an AWS MongoDB Altas database.
	//
	// 	- **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection to a Kafka cluster.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The database account of the destination instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodyDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodyDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetCanModifyPassword() *bool {
	return s.CanModifyPassword
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetAliyunUid(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.AliyunUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetCanModifyPassword(v bool) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.CanModifyPassword = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetDatabaseName(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetEngineName(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetInstanceID(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetInstanceType(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetIp(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetOracleSID(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetPort(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetRegion(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetRoleName(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.RoleName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetUserName(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodyMigrationMode struct {
	// Indicates whether data transformation is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DataExtractTransformLoad *bool `json:"DataExtractTransformLoad,omitempty" xml:"DataExtractTransformLoad,omitempty"`
	// Indicates whether full data migration or initial full data synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Indicates whether incremental data migration or synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Indicates whether schema migration or initial schema synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodyMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodyMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) GetDataExtractTransformLoad() *bool {
	return s.DataExtractTransformLoad
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) SetDataExtractTransformLoad(v bool) *DescribeDtsJobDetailResponseBodyMigrationMode {
	s.DataExtractTransformLoad = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) SetDataInitialization(v bool) *DescribeDtsJobDetailResponseBodyMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) SetDataSynchronization(v bool) *DescribeDtsJobDetailResponseBodyMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) SetStructureInitialization(v bool) *DescribeDtsJobDetailResponseBodyMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodyRetryState struct {
	// The error message returned if these retries failed.
	//
	// example:
	//
	// Unexpected error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The task ID.
	//
	// example:
	//
	// bi6e22ay243****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum duration of a retry. Unit: seconds.
	//
	// example:
	//
	// 7200
	MaxRetryTime *int32 `json:"MaxRetryTime,omitempty" xml:"MaxRetryTime,omitempty"`
	// The progress of the instance when DTS performs retries.
	//
	// example:
	//
	// 03
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 5
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// The object on which the retries are performed. Valid values:
	//
	// 	- **srcDB**: the source database.
	//
	// 	- **destDB**: the destination database.
	//
	// 	- **inner_module**: an internal module of DTS.
	//
	// example:
	//
	// srcDB
	RetryTarget *string `json:"RetryTarget,omitempty" xml:"RetryTarget,omitempty"`
	// The time that has elapsed from the point in time when the first retry starts. Unit: seconds.
	//
	// example:
	//
	// 3600
	RetryTime *int32 `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	// Indicates whether the task is being retried. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Retrying *bool `json:"Retrying,omitempty" xml:"Retrying,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodyRetryState) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodyRetryState) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) GetMaxRetryTime() *int32 {
	return s.MaxRetryTime
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) GetModule() *string {
	return s.Module
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) GetRetryTarget() *string {
	return s.RetryTarget
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) GetRetryTime() *int32 {
	return s.RetryTime
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) GetRetrying() *bool {
	return s.Retrying
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) SetErrMessage(v string) *DescribeDtsJobDetailResponseBodyRetryState {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) SetJobId(v string) *DescribeDtsJobDetailResponseBodyRetryState {
	s.JobId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) SetMaxRetryTime(v int32) *DescribeDtsJobDetailResponseBodyRetryState {
	s.MaxRetryTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) SetModule(v string) *DescribeDtsJobDetailResponseBodyRetryState {
	s.Module = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) SetRetryCount(v int32) *DescribeDtsJobDetailResponseBodyRetryState {
	s.RetryCount = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) SetRetryTarget(v string) *DescribeDtsJobDetailResponseBodyRetryState {
	s.RetryTarget = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) SetRetryTime(v int32) *DescribeDtsJobDetailResponseBodyRetryState {
	s.RetryTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) SetRetrying(v bool) *DescribeDtsJobDetailResponseBodyRetryState {
	s.Retrying = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyRetryState) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySourceEndpoint struct {
	// The ID of the Alibaba Cloud account to which the source instance belongs.
	//
	// example:
	//
	// 140692647406****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// Indicates whether the password can be modified. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CanModifyPassword *bool `json:"CanModifyPassword,omitempty" xml:"CanModifyPassword,omitempty"`
	// The name of the database from which the objects are migrated in the source instance.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine of the source instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The source instance ID.
	//
	// example:
	//
	// rm-bp1162kryivb8****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the source instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The endpoint of the source instance.
	//
	// example:
	//
	// 172.16.88.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The system ID (SID) of the Oracle database.
	//
	// > This parameter is returned only if the return value of **EngineName*	- of the source instance is **Oracle*	- and the Oracle database is deployed in a non-Real Application Cluster (RAC) architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the source instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the source instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the Resource Access Management (RAM) role configured for the Alibaba Cloud account to which the source instance belongs.
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **DISABLE**: SSL encryption is disabled.
	//
	// 	- **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// 	- **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection to an AWS MongoDB Altas database.
	//
	// 	- **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection to a Kafka cluster.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The database account of the source instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetCanModifyPassword() *bool {
	return s.CanModifyPassword
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetAliyunUid(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.AliyunUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetCanModifyPassword(v bool) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.CanModifyPassword = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetDatabaseName(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetEngineName(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetInstanceID(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetInstanceType(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetIp(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetOracleSID(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetPort(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetRegion(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetRoleName(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.RoleName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetUserName(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJob struct {
	// Indicates whether the new change tracking feature is used.
	//
	// > This parameter is returned only for change tracking instances of the new version.
	//
	// example:
	//
	// new
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The start of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-15T08:25:34Z
	BeginTimestamp *string `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	// The start offset of incremental data migration or data synchronization. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1616405159
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The consumption checkpoint of the change tracking instance. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-23T07:30:31Z
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	// The downstream client information in the following format: \\<IP address of the downstream client>:\\<Random ID generated by DTS>.
	//
	// example:
	//
	// 114.***.***.**:dts********
	ConsumptionClient *string `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-12T08:34:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The state of the ETL task.
	//
	// > This parameter collection is returned only if an ETL task is configured.
	DataEtlStatus *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus `json:"DataEtlStatus,omitempty" xml:"DataEtlStatus,omitempty" type:"Struct"`
	// The state of full data migration or initial full data synchronization.
	DataInitializationStatus *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The state of incremental data migration or synchronization.
	DataSynchronizationStatus *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The number of ApsaraDB RDS for MySQL instances that are attached to the source PolarDB-X 1.0 instance.
	//
	// example:
	//
	// 2
	DatabaseCount *int32 `json:"DatabaseCount,omitempty" xml:"DatabaseCount,omitempty"`
	// The objects of the data migration, data synchronization, or change tracking task. For more information, see [Objects of DTS tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// example:
	//
	// {\\"dtstestdata\\":{\\"all\\":true,\\"name\\":\\"dtstestdata\\",\\"state\\":\\"normal\\"}}
	DbObject *string `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	// The latency of incremental data migration or synchronization. Unit: milliseconds.
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The network type of the consumer client. Valid values:
	//
	// 	- **CLASSIC**: classic network.
	//
	// 	- **VPC**: VPC.
	//
	// example:
	//
	// VPC
	DestNetType *string `json:"DestNetType,omitempty" xml:"DestNetType,omitempty"`
	// The connection settings of the destination instance.
	DestinationEndpoint *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The DTS instance ID.
	//
	// example:
	//
	// dtsnjuc14kp12u****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The instance class.
	//
	// example:
	//
	// xlarge
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// > This parameter is returned only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	DtsJobDirection *string `json:"DtsJobDirection,omitempty" xml:"DtsJobDirection,omitempty"`
	// The DTS task ID.
	//
	// example:
	//
	// m06j1g92124****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The DTS instance name.
	//
	// example:
	//
	// dtstest****
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The end of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-26T14:03:21Z
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The operator information of the ETL task.
	//
	// > This parameter is returned only if you query the details of an ETL task.
	//
	// example:
	//
	// { "cells ": [{\\"shape\\":\\"edge\\",\\"attrs\\":{\\"line\\":{\\"stroke\\":\\"#b1b1b1\\",\\"strokeWidth\\":1,\\"targetMarker\\":{\\"name\\":\\"block\\",\\"args\\":{\\"size\\":\\"8\\"}},\\"strokeDasharray\\":\\"\\"}},\\"id\\":\\"cd1ec473-f9b9-4e9b-a742-ac23f442****\\",\\"source\\":{\\"cell\\":\\"8b261182-bfab-4803-ad8e-6bb08e3e****\\",\\"port\\":\\"out1\\"},\\"target\\":{\\"cell\\":\\"b36770df-f48c-4d6b-9644-54c5e924****\\",\\"port\\":\\"in1\\"},\\"zIndex\\":7 }] }
	EtlCalculator *string `json:"EtlCalculator,omitempty" xml:"EtlCalculator,omitempty"`
	// The time when the instance expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// > This parameter is returned only if the return value of **PayType*	- is **PrePaid**.
	//
	// example:
	//
	// 2023-06-16T08:01:19Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the task was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-16T10:34:17Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether the task is a subtask. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsDemoJob *bool `json:"IsDemoJob,omitempty" xml:"IsDemoJob,omitempty"`
	// The type of the DTS task. Valid values:
	//
	// 	- **online**: data migration task.
	//
	// 	- **SYNC**: data synchronization task.
	//
	// 	- **SUBSCRIBE**: change tracking task.
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
	// The migration types or initial synchronization types.
	MigrationMode *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// The minimum number of DUs.
	//
	// > This parameter is supported only for serverless instances.
	//
	// example:
	//
	// 1
	MinDu *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	// The source of the task. Valid values:
	//
	// 	- **PTS**
	//
	// 	- **DMS**
	//
	// 	- **DTS**
	//
	// example:
	//
	// DTS
	OriginType *string `json:"OriginType,omitempty" xml:"OriginType,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The performance of the data migration or synchronization instance.
	Performance *DescribeDtsJobDetailResponseBodySubDistributedJobPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck state.
	PrecheckStatus *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet specific requirements, such as whether to automatically start a precheck. For more information, see [MigrationReserved](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {\\"srcHostPorts\\":\\"\\",\\"whitelist.dms.online.ddl.enable\\":false,\\"filterDDL\\":false,\\"sqlparser.dms.original.ddl\\":true,\\"srcOracleType\\":\\"sid\\",\\"maxRetryTime\\":43200,\\"destSSL\\":\\"0\\",\\"destOracleType\\":\\"sid\\",\\"srcSSL\\":\\"0\\",\\"dbListCaseChangeMode\\":\\"default\\",\\"SourceEngineVersion\\":\\"8.0.18\\",\\"srcNetType\\":\\"VPC\\",\\"destNetType\\":\\"VPC\\",\\"srcVpcNetMappingInst\\":\\"172.16.1**.**:10803\\",\\"destVpcNetMappingInst\\":\\"172.16.1**.**:11077\\",\\"useJobTask\\":\\"1\\"}
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The information about the retries performed by DTS due to an exception.
	RetryState *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState `json:"RetryState,omitempty" xml:"RetryState,omitempty" type:"Struct"`
	// The details of the data synchronization task in the reverse direction.
	//
	// > This parameter is returned only for two-way data synchronization tasks.
	ReverseJob *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob `json:"ReverseJob,omitempty" xml:"ReverseJob,omitempty" type:"Struct"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The state of initial schema synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Initialization status of library table structure.
	StructureInitializationStatus *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	// The information about the subtasks in the current data synchronization task.
	SubSyncJob []interface{} `json:"SubSyncJob,omitempty" xml:"SubSyncJob,omitempty" type:"Repeated"`
	// The topic of the change tracking instance.
	//
	// > This parameter is returned only if your change tracking instances are of the new version and you have called the [CreateConsumerGroup](https://help.aliyun.com/document_detail/122863.html) operation to create a consumer group.
	//
	// example:
	//
	// cn_hangzhou_rm_bp1162kryivb8****_dtstest_version2
	SubscribeTopic *string `json:"SubscribeTopic,omitempty" xml:"SubscribeTopic,omitempty"`
	// The type of data for change tracking.
	SubscriptionDataType *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	// The endpoint of the change tracking instance.
	SubscriptionHost *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost `json:"SubscriptionHost,omitempty" xml:"SubscriptionHost,omitempty" type:"Struct"`
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
	// 	- The value **Reverse*	- takes effect only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The tags of the task.
	TagList []*DescribeDtsJobDetailResponseBodySubDistributedJobTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The task type.
	//
	// example:
	//
	// rds
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJob) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetAppName() *string {
	return s.AppName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetBeginTimestamp() *string {
	return s.BeginTimestamp
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetConsumptionCheckpoint() *string {
	return s.ConsumptionCheckpoint
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetConsumptionClient() *string {
	return s.ConsumptionClient
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDataEtlStatus() *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus {
	return s.DataEtlStatus
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDataInitializationStatus() *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDataSynchronizationStatus() *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDatabaseCount() *int32 {
	return s.DatabaseCount
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDbObject() *string {
	return s.DbObject
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDestNetType() *string {
	return s.DestNetType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDestinationEndpoint() *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDtsJobDirection() *string {
	return s.DtsJobDirection
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetEtlCalculator() *string {
	return s.EtlCalculator
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetIsDemoJob() *bool {
	return s.IsDemoJob
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetJobType() *string {
	return s.JobType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetMaxDu() *float64 {
	return s.MaxDu
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetMigrationMode() *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode {
	return s.MigrationMode
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetMinDu() *float64 {
	return s.MinDu
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetOriginType() *string {
	return s.OriginType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetPerformance() *DescribeDtsJobDetailResponseBodySubDistributedJobPerformance {
	return s.Performance
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetPrecheckStatus() *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetReserved() *string {
	return s.Reserved
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetRetryState() *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState {
	return s.RetryState
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetReverseJob() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	return s.ReverseJob
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetSourceEndpoint() *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetStructureInitializationStatus() *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetSubSyncJob() []interface{} {
	return s.SubSyncJob
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetSubscribeTopic() *string {
	return s.SubscribeTopic
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetSubscriptionDataType() *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType {
	return s.SubscriptionDataType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetSubscriptionHost() *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost {
	return s.SubscriptionHost
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetTagList() []*DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	return s.TagList
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetAppName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.AppName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetBeginTimestamp(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetCheckpoint(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.Checkpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetConsumptionCheckpoint(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetConsumptionClient(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetCreateTime(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDataEtlStatus(v *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.DataEtlStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDataInitializationStatus(v *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDataSynchronizationStatus(v *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDatabaseCount(v int32) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.DatabaseCount = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDbObject(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.DbObject = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDelay(v int64) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.Delay = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDestNetType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.DestNetType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDestinationEndpoint(v *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDtsInstanceID(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDtsJobClass(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDtsJobDirection(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDtsJobId(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetDtsJobName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetEndTimestamp(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetEtlCalculator(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.EtlCalculator = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetExpireTime(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetFinishTime(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.FinishTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetGroupId(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.GroupId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetIsDemoJob(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.IsDemoJob = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetJobType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.JobType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetMaxDu(v float64) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.MaxDu = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetMigrationMode(v *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.MigrationMode = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetMinDu(v float64) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.MinDu = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetOriginType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.OriginType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetPayType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.PayType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetPerformance(v *DescribeDtsJobDetailResponseBodySubDistributedJobPerformance) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.Performance = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetPrecheckStatus(v *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetReserved(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.Reserved = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetRetryState(v *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.RetryState = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetReverseJob(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.ReverseJob = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetSourceEndpoint(v *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetStructureInitializationStatus(v *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetSubSyncJob(v []interface{}) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.SubSyncJob = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetSubscribeTopic(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.SubscribeTopic = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetSubscriptionDataType(v *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.SubscriptionDataType = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetSubscriptionHost(v *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.SubscriptionHost = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetSynchronizationDirection(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetTagList(v []*DescribeDtsJobDetailResponseBodySubDistributedJobTagList) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.TagList = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) SetTaskType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJob {
	s.TaskType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJob) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus struct {
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of full data migration or initial full data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that are migrated or synchronized during full data migration or initial full data synchronization.
	//
	// example:
	//
	// 16
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of the ETL task. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// 	- **Catched**: The task has no latency.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataEtlStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus struct {
	// The error message returned if full data migration or initial full data synchronization failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of full data migration or initial full data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that are migrated or synchronized during full data migration or initial full data synchronization.
	//
	// example:
	//
	// 16
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of full data migration or initial full data synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus struct {
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of incremental data migration or synchronization.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of rows and size of data that is synchronized or migrated to the destination table per second during incremental data synchronization or migration.
	//
	// example:
	//
	// 0.00RPS/(0.000MB/s)
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of incremental data migration or synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Suspending**: The task is paused.
	//
	// 	- **Checking**: The task is in precheck.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// 	- **Catched**: The task has no latency.
	//
	// example:
	//
	// Catched
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint struct {
	// The ID of the Alibaba Cloud account to which the destination instance belongs.
	//
	// example:
	//
	// 140692647406****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The name of the database to which the objects are migrated in the destination instance.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine of the destination instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The destination instance ID.
	//
	// example:
	//
	// rm-bp1f9guj5rhzq****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the destination instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The endpoint of the destination instance.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the return value of **EngineName*	- of the destination instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the destination instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the destination instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account to which the destination instance belongs.
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **DISABLE**: SSL encryption is disabled.
	//
	// 	- **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// 	- **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection to an AWS MongoDB Altas database.
	//
	// 	- **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection to a Kafka cluster.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The database account of the destination instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) SetAliyunUid(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	s.AliyunUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) SetDatabaseName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) SetEngineName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) SetInstanceID(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) SetInstanceType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) SetIp(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) SetOracleSID(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) SetPort(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) SetRegion(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) SetRoleName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	s.RoleName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) SetUserName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode struct {
	// Indicates whether data transformation is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DataExtractTransformLoad *bool `json:"DataExtractTransformLoad,omitempty" xml:"DataExtractTransformLoad,omitempty"`
	// Indicates whether full data migration or initial full data synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Indicates whether incremental data migration or synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Indicates whether schema migration or initial schema synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode) GetDataExtractTransformLoad() *bool {
	return s.DataExtractTransformLoad
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode) SetDataExtractTransformLoad(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode {
	s.DataExtractTransformLoad = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode) SetDataInitialization(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode) SetDataSynchronization(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode) SetStructureInitialization(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobPerformance struct {
	// The size of data that is migrated or synchronized per second. Unit: Mbit/s.
	//
	// example:
	//
	// 1
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The number of times that SQL statements are migrated or synchronized per second, including BEGIN, COMMIT, DML, and DDL statements. DML statements include INSERT, DELETE, and UPDATE.
	//
	// example:
	//
	// 100
	Rps *string `json:"Rps,omitempty" xml:"Rps,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobPerformance) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPerformance) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPerformance) GetRps() *string {
	return s.Rps
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPerformance) SetFlow(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobPerformance {
	s.Flow = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPerformance) SetRps(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobPerformance {
	s.Rps = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPerformance) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus struct {
	// The result of each precheck item.
	Detail []*DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The precheck progress. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck state. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is in precheck.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus) GetDetail() []*DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus) SetDetail(v []*DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail struct {
	// The name of the precheck item.
	//
	// example:
	//
	// CHECK_CONN_SRC
	CheckItem *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	// The description of the precheck item.
	//
	// example:
	//
	// CHECK_CONN_SRC_DETAIL
	CheckItemDescription *string `json:"CheckItemDescription,omitempty" xml:"CheckItemDescription,omitempty"`
	// The precheck result. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Success
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The error message returned if the task failed to pass the precheck.
	//
	// > This parameter is returned only if the return value of **CheckResult*	- is **Failed**.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The method used to fix the precheck failure.
	//
	// > This parameter is returned only if the return value of **CheckResult*	- is **Failed**.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) GetCheckItemDescription() *string {
	return s.CheckItemDescription
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) GetCheckResult() *string {
	return s.CheckResult
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) GetFailedReason() *string {
	return s.FailedReason
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) SetCheckItem(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) SetCheckItemDescription(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) SetCheckResult(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) SetFailedReason(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail {
	s.FailedReason = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) SetRepairMethod(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobRetryState struct {
	// The error message returned.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ta7w132u12h****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum duration of a retry. Unit: seconds.
	//
	// example:
	//
	// 7200
	MaxRetryTime *int32 `json:"MaxRetryTime,omitempty" xml:"MaxRetryTime,omitempty"`
	// The progress of the instance when DTS performs retries.
	//
	// example:
	//
	// 03
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 5
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// The object on which the retries are performed. Valid values:
	//
	// 	- **srcDB**: the source database.
	//
	// 	- **destDB**: the destination database.
	//
	// 	- **inner_module**: an internal module of DTS.
	//
	// example:
	//
	// srcDB
	RetryTarget *string `json:"RetryTarget,omitempty" xml:"RetryTarget,omitempty"`
	// The time that has elapsed from the point in time when the first retry starts. Unit: seconds.
	//
	// example:
	//
	// 3600
	RetryTime *int32 `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	// Indicates whether the task is being retried. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Retrying *bool `json:"Retrying,omitempty" xml:"Retrying,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) GetMaxRetryTime() *int32 {
	return s.MaxRetryTime
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) GetModule() *string {
	return s.Module
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) GetRetryTarget() *string {
	return s.RetryTarget
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) GetRetryTime() *int32 {
	return s.RetryTime
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) GetRetrying() *bool {
	return s.Retrying
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) SetErrMsg(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState {
	s.ErrMsg = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) SetJobId(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState {
	s.JobId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) SetMaxRetryTime(v int32) *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState {
	s.MaxRetryTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) SetModule(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState {
	s.Module = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) SetRetryCount(v int32) *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState {
	s.RetryCount = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) SetRetryTarget(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState {
	s.RetryTarget = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) SetRetryTime(v int32) *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState {
	s.RetryTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) SetRetrying(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState {
	s.Retrying = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobRetryState) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob struct {
	// Indicates whether the new change tracking feature is used.
	//
	// > This parameter is returned only for change tracking instances of the new version.
	//
	// example:
	//
	// new
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The start of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-15T08:25:34Z
	BeginTimestamp *string `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	// The start offset of incremental data migration or data synchronization. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1616405159
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The consumption checkpoint of the change tracking instance. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-23T07:30:31Z
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	// The downstream client information in the following format: \\<IP address of the downstream client>:\\<Random ID generated by DTS>.
	//
	// example:
	//
	// 114.***.***.**:dts********
	ConsumptionClient *string `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-12T08:34:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The state of the ETL task.
	//
	// > This parameter collection is returned only if an ETL task is configured.
	DataEtlStatus *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus `json:"DataEtlStatus,omitempty" xml:"DataEtlStatus,omitempty" type:"Struct"`
	// The state of full data migration or initial full data synchronization.
	DataInitializationStatus *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The state of incremental data migration or synchronization.
	DataSynchronizationStatus *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The number of ApsaraDB RDS for MySQL instances that are attached to the source PolarDB-X 1.0 instance.
	//
	// example:
	//
	// 2
	DatabaseCount *int32 `json:"DatabaseCount,omitempty" xml:"DatabaseCount,omitempty"`
	// The objects of the data migration, data synchronization, or change tracking task. For more information, see [Objects of DTS tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// example:
	//
	// {\\"dtstestdata\\":{\\"all\\":true,\\"name\\":\\"dtstestdata\\",\\"state\\":\\"normal\\"}}
	DbObject *string `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	// The latency of incremental data migration or synchronization. Unit: milliseconds.
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The network type of the consumer client. Valid values:
	//
	// 	- **CLASSIC**: classic network.
	//
	// 	- **VPC**: VPC.
	//
	// example:
	//
	// VPC
	DestNetType *string `json:"DestNetType,omitempty" xml:"DestNetType,omitempty"`
	// The connection settings of the destination instance.
	DestinationEndpoint *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The DTS instance ID.
	//
	// example:
	//
	// dtsnjuc14kp12u****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The instance class.
	//
	// example:
	//
	// xlarge
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// > This parameter is returned only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	DtsJobDirection *string `json:"DtsJobDirection,omitempty" xml:"DtsJobDirection,omitempty"`
	// The DTS task ID.
	//
	// example:
	//
	// m06j1g92124****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The DTS instance name.
	//
	// example:
	//
	// dtstest****
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The end of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-26T14:03:21Z
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The operator information of the ETL task.
	//
	// > This parameter is returned only if you query the details of an ETL task.
	//
	// example:
	//
	// { "cells ": [{\\"shape\\":\\"edge\\",\\"attrs\\":{\\"line\\":{\\"stroke\\":\\"#b1b1b1\\",\\"strokeWidth\\":1,\\"targetMarker\\":{\\"name\\":\\"block\\",\\"args\\":{\\"size\\":\\"8\\"}},\\"strokeDasharray\\":\\"\\"}},\\"id\\":\\"cd1ec473-f9b9-4e9b-a742-ac23f442****\\",\\"source\\":{\\"cell\\":\\"8b261182-bfab-4803-ad8e-6bb08e3e****\\",\\"port\\":\\"out1\\"},\\"target\\":{\\"cell\\":\\"b36770df-f48c-4d6b-9644-54c5e924****\\",\\"port\\":\\"in1\\"},\\"zIndex\\":7 }] }
	EtlCalculator *string `json:"EtlCalculator,omitempty" xml:"EtlCalculator,omitempty"`
	// The time when the instance expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// > This parameter is returned only if the return value of **PayType*	- is **PrePaid**.
	//
	// example:
	//
	// 2023-06-16T08:01:19Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the task was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-16T10:34:17Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether the task is a subtask. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsDemoJob *bool `json:"IsDemoJob,omitempty" xml:"IsDemoJob,omitempty"`
	// The type of the DTS task. Valid values:
	//
	// 	- **online**: data migration task.
	//
	// 	- **SYNC**: data synchronization task.
	//
	// 	- **SUBSCRIBE**: change tracking task.
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
	// The migration types or initial synchronization types.
	MigrationMode *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// The minimum number of DTS Units (DUs).
	//
	// > This parameter is supported only for serverless instances.
	//
	// example:
	//
	// 1
	MinDu *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	// The source of the task. Valid values:
	//
	// 	- **PTS**
	//
	// 	- **DMS**
	//
	// 	- **DTS**
	//
	// example:
	//
	// DTS
	OriginType *string `json:"OriginType,omitempty" xml:"OriginType,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The performance of the data migration or synchronization instance.
	Performance *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck state.
	PrecheckStatus *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet specific requirements, such as whether to automatically start a precheck. For more information, see [MigrationReserved](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {\\"srcHostPorts\\":\\"\\",\\"whitelist.dms.online.ddl.enable\\":false,\\"filterDDL\\":false,\\"sqlparser.dms.original.ddl\\":true,\\"srcOracleType\\":\\"sid\\",\\"maxRetryTime\\":43200,\\"destSSL\\":\\"0\\",\\"destOracleType\\":\\"sid\\",\\"srcSSL\\":\\"0\\",\\"dbListCaseChangeMode\\":\\"default\\",\\"SourceEngineVersion\\":\\"8.0.18\\",\\"srcNetType\\":\\"VPC\\",\\"destNetType\\":\\"VPC\\",\\"srcVpcNetMappingInst\\":\\"172.16.1**.**:10803\\",\\"destVpcNetMappingInst\\":\\"172.16.1**.**:11077\\",\\"useJobTask\\":\\"1\\"}
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The information about the retries performed by DTS due to an exception.
	RetryState *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState `json:"RetryState,omitempty" xml:"RetryState,omitempty" type:"Struct"`
	// The details of the data synchronization task in the reverse direction.
	//
	// > This parameter is returned only for two-way data synchronization tasks.
	//
	// example:
	//
	// ****
	ReverseJob interface{} `json:"ReverseJob,omitempty" xml:"ReverseJob,omitempty"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The state of initial schema synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Initialization status of library table structure.
	StructureInitializationStatus *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	// The topic of the change tracking instance.
	//
	// > This parameter is returned only if your change tracking instances are of the new version and you have called the [CreateConsumerGroup](https://help.aliyun.com/document_detail/122863.html) operation to create a consumer group.
	//
	// example:
	//
	// cn_hangzhou_rm_bp1162kryivb8****_dtstest_version2
	SubscribeTopic *string `json:"SubscribeTopic,omitempty" xml:"SubscribeTopic,omitempty"`
	// The type of data for change tracking.
	SubscriptionDataType *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	// The endpoint of the change tracking instance.
	SubscriptionHost *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost `json:"SubscriptionHost,omitempty" xml:"SubscriptionHost,omitempty" type:"Struct"`
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
	// 	- The value **Reverse*	- takes effect only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The tags of the task.
	TagList []*DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The task type.
	//
	// example:
	//
	// rds
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetAppName() *string {
	return s.AppName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetBeginTimestamp() *string {
	return s.BeginTimestamp
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetConsumptionCheckpoint() *string {
	return s.ConsumptionCheckpoint
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetConsumptionClient() *string {
	return s.ConsumptionClient
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDataEtlStatus() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus {
	return s.DataEtlStatus
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDataInitializationStatus() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDataSynchronizationStatus() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDatabaseCount() *int32 {
	return s.DatabaseCount
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDbObject() *string {
	return s.DbObject
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDestNetType() *string {
	return s.DestNetType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDestinationEndpoint() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDtsJobDirection() *string {
	return s.DtsJobDirection
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetEtlCalculator() *string {
	return s.EtlCalculator
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetIsDemoJob() *bool {
	return s.IsDemoJob
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetJobType() *string {
	return s.JobType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetMaxDu() *float64 {
	return s.MaxDu
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetMigrationMode() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode {
	return s.MigrationMode
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetMinDu() *float64 {
	return s.MinDu
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetOriginType() *string {
	return s.OriginType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetPerformance() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance {
	return s.Performance
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetPrecheckStatus() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetReserved() *string {
	return s.Reserved
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetRetryState() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState {
	return s.RetryState
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetReverseJob() interface{} {
	return s.ReverseJob
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetSourceEndpoint() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetStructureInitializationStatus() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetSubscribeTopic() *string {
	return s.SubscribeTopic
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetSubscriptionDataType() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType {
	return s.SubscriptionDataType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetSubscriptionHost() *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost {
	return s.SubscriptionHost
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetTagList() []*DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	return s.TagList
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetAppName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.AppName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetBeginTimestamp(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetCheckpoint(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.Checkpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetConsumptionCheckpoint(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetConsumptionClient(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetCreateTime(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDataEtlStatus(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.DataEtlStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDataInitializationStatus(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDataSynchronizationStatus(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDatabaseCount(v int32) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.DatabaseCount = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDbObject(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.DbObject = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDelay(v int64) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.Delay = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDestNetType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.DestNetType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDestinationEndpoint(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDtsInstanceID(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDtsJobClass(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDtsJobDirection(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDtsJobId(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetDtsJobName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetEndTimestamp(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetEtlCalculator(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.EtlCalculator = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetExpireTime(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetFinishTime(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.FinishTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetGroupId(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.GroupId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetIsDemoJob(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.IsDemoJob = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetJobType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.JobType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetMaxDu(v float64) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.MaxDu = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetMigrationMode(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.MigrationMode = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetMinDu(v float64) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.MinDu = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetOriginType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.OriginType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetPayType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.PayType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetPerformance(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.Performance = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetPrecheckStatus(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetReserved(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.Reserved = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetRetryState(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.RetryState = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetReverseJob(v interface{}) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.ReverseJob = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetSourceEndpoint(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetStructureInitializationStatus(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetSubscribeTopic(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.SubscribeTopic = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetSubscriptionDataType(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.SubscriptionDataType = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetSubscriptionHost(v *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.SubscriptionHost = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetSynchronizationDirection(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetTagList(v []*DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.TagList = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) SetTaskType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob {
	s.TaskType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJob) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus struct {
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of full data migration or initial full data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that are migrated or synchronized during full data migration or initial full data synchronization.
	//
	// example:
	//
	// 16
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of the ETL task. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// 	- **Catched**: The task has no latency.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataEtlStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus struct {
	// The error message returned if full data migration or initial full data synchronization failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of full data migration or initial full data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that are migrated or synchronized during full data migration or initial full data synchronization.
	//
	// example:
	//
	// 16
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of full data migration or initial full data synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus struct {
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of incremental data migration or synchronization.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of rows and size of data that is synchronized or migrated to the destination table per second during incremental data synchronization or migration.
	//
	// example:
	//
	// 0.00RPS/(0.000MB/s)
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of incremental data migration or synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Suspending**: The task is paused.
	//
	// 	- **Checking**: The task is in precheck.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// 	- **Catched**: The task has no latency.
	//
	// example:
	//
	// Catched
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint struct {
	// The ID of the Alibaba Cloud account to which the destination instance belongs.
	//
	// example:
	//
	// 140692647406****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The name of the database to which the objects are migrated in the destination instance.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine of the destination instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The destination instance ID.
	//
	// example:
	//
	// rm-bp1f9guj5rhzq****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the destination instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The endpoint of the destination instance.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the return value of **EngineName*	- of the destination instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the destination instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the destination instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account to which the destination instance belongs.
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **DISABLE**: SSL encryption is disabled.
	//
	// 	- **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// 	- **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection to an AWS MongoDB Altas database.
	//
	// 	- **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection to a Kafka cluster.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The database account of the destination instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) SetAliyunUid(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	s.AliyunUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) SetDatabaseName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) SetEngineName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) SetInstanceID(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) SetInstanceType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) SetIp(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) SetOracleSID(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) SetPort(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) SetRegion(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) SetRoleName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	s.RoleName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) SetUserName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode struct {
	// Indicates whether data transformation is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DataExtractTransformLoad *bool `json:"DataExtractTransformLoad,omitempty" xml:"DataExtractTransformLoad,omitempty"`
	// Indicates whether full data migration or initial full data synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Indicates whether incremental data migration or synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Indicates whether schema migration or initial schema synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode) GetDataExtractTransformLoad() *bool {
	return s.DataExtractTransformLoad
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode) SetDataExtractTransformLoad(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode {
	s.DataExtractTransformLoad = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode) SetDataInitialization(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode) SetDataSynchronization(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode) SetStructureInitialization(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance struct {
	// The size of data that is migrated or synchronized per second. Unit: Mbit/s.
	//
	// example:
	//
	// 1
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The number of times that SQL statements are migrated or synchronized per second, including BEGIN, COMMIT, DML, and DDL statements. DML statements include INSERT, DELETE, and UPDATE.
	//
	// example:
	//
	// 100
	Rps *string `json:"Rps,omitempty" xml:"Rps,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance) GetRps() *string {
	return s.Rps
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance) SetFlow(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance {
	s.Flow = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance) SetRps(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance {
	s.Rps = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPerformance) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus struct {
	// The result of each precheck item.
	Detail []*DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The precheck progress. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck state. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is in precheck.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus) GetDetail() []*DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus) SetDetail(v []*DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail struct {
	// The name of the precheck item.
	//
	// example:
	//
	// CHECK_CONN_SRC
	CheckItem *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	// The description of the precheck item.
	//
	// example:
	//
	// CHECK_CONN_SRC_DETAIL
	CheckItemDescription *string `json:"CheckItemDescription,omitempty" xml:"CheckItemDescription,omitempty"`
	// The precheck result. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Success
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The error message returned if the task failed to pass the precheck.
	//
	// > This parameter is returned only if the return value of **CheckResult*	- is **Failed**.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The method used to fix the precheck failure.
	//
	// > This parameter is returned only if the return value of **CheckResult*	- is **Failed**.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) GetCheckItemDescription() *string {
	return s.CheckItemDescription
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) GetCheckResult() *string {
	return s.CheckResult
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) GetFailedReason() *string {
	return s.FailedReason
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) SetCheckItem(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) SetCheckItemDescription(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) SetCheckResult(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) SetFailedReason(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail {
	s.FailedReason = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) SetRepairMethod(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState struct {
	// The error message returned.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ta7w132u12h****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum duration of a retry. Unit: seconds.
	//
	// example:
	//
	// 7200
	MaxRetryTime *int32 `json:"MaxRetryTime,omitempty" xml:"MaxRetryTime,omitempty"`
	// The progress of the instance when DTS performs retries.
	//
	// example:
	//
	// 03
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 5
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// The object on which the retries are performed. Valid values:
	//
	// 	- **srcDB**: the source database.
	//
	// 	- **destDB**: the destination database.
	//
	// 	- **inner_module**: an internal module of DTS.
	//
	// example:
	//
	// srcDB
	RetryTarget *string `json:"RetryTarget,omitempty" xml:"RetryTarget,omitempty"`
	// The time that has elapsed from the point in time when the first retry starts. Unit: seconds.
	//
	// example:
	//
	// 3600
	RetryTime *int32 `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	// Indicates whether the task is being retried. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Retrying *bool `json:"Retrying,omitempty" xml:"Retrying,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) GetMaxRetryTime() *int32 {
	return s.MaxRetryTime
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) GetModule() *string {
	return s.Module
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) GetRetryTarget() *string {
	return s.RetryTarget
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) GetRetryTime() *int32 {
	return s.RetryTime
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) GetRetrying() *bool {
	return s.Retrying
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) SetErrMsg(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState {
	s.ErrMsg = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) SetJobId(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState {
	s.JobId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) SetMaxRetryTime(v int32) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState {
	s.MaxRetryTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) SetModule(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState {
	s.Module = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) SetRetryCount(v int32) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState {
	s.RetryCount = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) SetRetryTarget(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState {
	s.RetryTarget = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) SetRetryTime(v int32) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState {
	s.RetryTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) SetRetrying(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState {
	s.Retrying = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobRetryState) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint struct {
	// The ID of the Alibaba Cloud account to which the source instance belongs.
	//
	// example:
	//
	// 140692647406****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The name of the database from which the objects are migrated in the source instance.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine of the source instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The source instance ID.
	//
	// example:
	//
	// rm-bp2f3huj5rhzq****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the source instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The endpoint of the source instance.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the return value of **EngineName*	- of the destination instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the source instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the source instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account to which the source instance belongs.
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **DISABLE**: SSL encryption is disabled.
	//
	// 	- **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// 	- **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection to an AWS MongoDB Altas database.
	//
	// 	- **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection to a Kafka cluster.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The database account of the source instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) SetAliyunUid(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	s.AliyunUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) SetDatabaseName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) SetEngineName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) SetInstanceID(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) SetInstanceType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) SetIp(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) SetOracleSID(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) SetPort(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) SetRegion(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) SetRoleName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	s.RoleName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) SetUserName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus struct {
	// Error message indicating task failure.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Whether to display upgrade specifications, return value:
	//
	// - True: Yes.
	//
	// - False: No.
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// Initialization progress of library table structure, measured in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables that have completed library table structure initialization.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The initialization status of the library table structure includes:
	//
	// - NotStarted: Not started.
	//
	// - Migration: In the process of initialization.
	//
	// - Failed: Initialization failed.
	//
	// - Finished: Initialization completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType struct {
	// Indicates whether DDL statements are tracked. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Ddl *bool `json:"Ddl,omitempty" xml:"Ddl,omitempty"`
	// Indicates whether DML statements are tracked. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Dml *bool `json:"Dml,omitempty" xml:"Dml,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType) GetDdl() *bool {
	return s.Ddl
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType) GetDml() *bool {
	return s.Dml
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType) SetDdl(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType {
	s.Ddl = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType) SetDml(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType {
	s.Dml = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionDataType) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost struct {
	// The private endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****-internal.aliyuncs.com:18002
	PrivateHost *string `json:"PrivateHost,omitempty" xml:"PrivateHost,omitempty"`
	// The public endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****.aliyuncs.com:18001
	PublicHost *string `json:"PublicHost,omitempty" xml:"PublicHost,omitempty"`
	// The VPC endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****.aliyuncs.com:18001
	VpcHost *string `json:"VpcHost,omitempty" xml:"VpcHost,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost) GetPrivateHost() *string {
	return s.PrivateHost
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost) GetPublicHost() *string {
	return s.PublicHost
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost) GetVpcHost() *string {
	return s.VpcHost
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost) SetPrivateHost(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost {
	s.PrivateHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost) SetPublicHost(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost {
	s.PublicHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost) SetVpcHost(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost {
	s.VpcHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobSubscriptionHost) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 191448876515****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The operator of the tag.
	//
	// example:
	//
	// 191448876515****
	Creator *int64 `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the task was modified.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary key of the table.
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the region in which the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking instance.
	//
	// example:
	//
	// dtsnjuc14kp12u****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::DTS::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Indicates whether the tag is visible. Valid values:
	//
	// 	- **0**: The tag is public.
	//
	// 	- **1**: The tag is private.
	//
	// example:
	//
	// 0
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The ID of the region in which the DTS task resides.
	//
	// > In most cases, the ID of the region in which the destination instance resides is returned.
	//
	// example:
	//
	// cn-hangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The type of the tag. Valid values:
	//
	// 	- **System**: The tag was created by the system.
	//
	// 	- **Custom**: The tag was created by a user.
	//
	// > By default, if the parameter is left empty, custom tags and system tags are returned.
	//
	// example:
	//
	// System
	TagCategory *string `json:"TagCategory,omitempty" xml:"TagCategory,omitempty"`
	// The tag key.
	//
	// example:
	//
	// key1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetCreator() *int64 {
	return s.Creator
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetId() *int64 {
	return s.Id
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetScope() *string {
	return s.Scope
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetTagCategory() *string {
	return s.TagCategory
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetAliUid(v int64) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.AliUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetCreator(v int64) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.Creator = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetGmtCreate(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetGmtModified(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.GmtModified = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetId(v int64) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.Id = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetRegionId(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.RegionId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetResourceId(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.ResourceId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetResourceType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.ResourceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetScope(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.Scope = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetSrcRegion(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.SrcRegion = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetTagCategory(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.TagCategory = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetTagKey(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) SetTagValue(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList {
	s.TagValue = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobReverseJobTagList) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint struct {
	// The ID of the Alibaba Cloud account to which the source instance belongs.
	//
	// example:
	//
	// 140692647406****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The name of the database from which the objects are migrated in the source instance.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine of the source instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The source instance ID.
	//
	// example:
	//
	// dtsnjuc14kp12u****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the source instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The endpoint of the source instance.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the return value of **EngineName*	- of the destination instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the source instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the source instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account to which the source instance belongs.
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **DISABLE**: SSL encryption is disabled.
	//
	// 	- **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// 	- **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection to an AWS MongoDB Altas database.
	//
	// 	- **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection to a Kafka cluster.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The database account of the source instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) SetAliyunUid(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	s.AliyunUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) SetDatabaseName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) SetEngineName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) SetInstanceID(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) SetInstanceType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) SetIp(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) SetOracleSID(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) SetPort(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) SetRegion(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) SetRoleName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	s.RoleName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) SetUserName(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus struct {
	// Error message.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Whether to display upgrade specifications, return value:
	//
	// - True: Yes.
	//
	// - False: No.
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// Initialization progress of library table structure, measured in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables that have completed library table structure initialization.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The initialization status of the library table structure includes:
	//
	// - NotStarted: Not started.
	//
	// - Migration: In the process of initialization.
	//
	// - Failed: Initialization failed.
	//
	// - Finished: Initialization completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType struct {
	// Indicates whether DDL statements are tracked. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Ddl *bool `json:"Ddl,omitempty" xml:"Ddl,omitempty"`
	// Indicates whether DML statements are tracked. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Dml *bool `json:"Dml,omitempty" xml:"Dml,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType) GetDdl() *bool {
	return s.Ddl
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType) GetDml() *bool {
	return s.Dml
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType) SetDdl(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType {
	s.Ddl = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType) SetDml(v bool) *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType {
	s.Dml = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionDataType) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost struct {
	// The private endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****-internal.aliyuncs.com:18002
	PrivateHost *string `json:"PrivateHost,omitempty" xml:"PrivateHost,omitempty"`
	// The public endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****.aliyuncs.com:18001
	PublicHost *string `json:"PublicHost,omitempty" xml:"PublicHost,omitempty"`
	// The VPC endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****.aliyuncs.com:18001
	VpcHost *string `json:"VpcHost,omitempty" xml:"VpcHost,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost) GetPrivateHost() *string {
	return s.PrivateHost
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost) GetPublicHost() *string {
	return s.PublicHost
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost) GetVpcHost() *string {
	return s.VpcHost
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost) SetPrivateHost(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost {
	s.PrivateHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost) SetPublicHost(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost {
	s.PublicHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost) SetVpcHost(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost {
	s.VpcHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobSubscriptionHost) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubDistributedJobTagList struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 191448876515****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The operator of the tag.
	//
	// example:
	//
	// 191448876515****
	Creator *int64 `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the task was modified.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary key of the table.
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the region in which the DTS task resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking instance.
	//
	// example:
	//
	// dtsnjuc14kp12u****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::DTS::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Indicates whether the tag is visible. Valid values:
	//
	// 	- **0**: The tag is public.
	//
	// 	- **1**: The tag is private.
	//
	// example:
	//
	// 0
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The ID of the region in which the DTS task resides.
	//
	// > In most cases, the ID of the region in which the destination instance resides is returned.
	//
	// example:
	//
	// cn-hangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The type of the tag. Valid values:
	//
	// 	- **System**: The tag was created by the system.
	//
	// 	- **Custom**: The tag was created by a user.
	//
	// > By default, if the parameter is left empty, custom tags and system tags are returned.
	//
	// example:
	//
	// System
	TagCategory *string `json:"TagCategory,omitempty" xml:"TagCategory,omitempty"`
	// The tag key.
	//
	// example:
	//
	// key1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetCreator() *int64 {
	return s.Creator
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetId() *int64 {
	return s.Id
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetScope() *string {
	return s.Scope
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetTagCategory() *string {
	return s.TagCategory
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetAliUid(v int64) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.AliUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetCreator(v int64) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.Creator = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetGmtCreate(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetGmtModified(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.GmtModified = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetId(v int64) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.Id = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetRegionId(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.RegionId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetResourceId(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.ResourceId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetResourceType(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.ResourceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetScope(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.Scope = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetSrcRegion(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.SrcRegion = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetTagCategory(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.TagCategory = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetTagKey(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) SetTagValue(v string) *DescribeDtsJobDetailResponseBodySubDistributedJobTagList {
	s.TagValue = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubDistributedJobTagList) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJob struct {
	// Indicates whether the new change tracking feature is used.
	//
	// > This parameter is returned only for change tracking instances of the new version.
	//
	// example:
	//
	// new
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The start of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-15T08:25:34Z
	BeginTimestamp *string `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	// The start offset of incremental data migration or data synchronization. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1616405159
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The consumption checkpoint of the change tracking instance. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-23T07:30:31Z
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	// The downstream client information in the following format: \\<IP address of the downstream client>:\\<Random ID generated by DTS>.
	//
	// example:
	//
	// 114.***.***.**:dts********
	ConsumptionClient *string `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-12T08:34:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The state of the ETL task.
	//
	// > This parameter collection is returned only if an ETL task is configured.
	DataEtlStatus *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus `json:"DataEtlStatus,omitempty" xml:"DataEtlStatus,omitempty" type:"Struct"`
	// The state of full data migration or initial full data synchronization.
	DataInitializationStatus *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The state of incremental data migration or synchronization.
	DataSynchronizationStatus *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The number of ApsaraDB RDS for MySQL instances that are attached to the source PolarDB-X 1.0 instance.
	//
	// example:
	//
	// 2
	DatabaseCount *int32 `json:"DatabaseCount,omitempty" xml:"DatabaseCount,omitempty"`
	// The objects of the data migration, data synchronization, or change tracking task. For more information, see [Objects of DTS tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// example:
	//
	// {\\"dtstestdata\\":{\\"all\\":true,\\"name\\":\\"dtstestdata\\",\\"state\\":\\"normal\\"}}
	DbObject *string `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	// The latency of incremental data migration or synchronization. Unit: milliseconds.
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The network type of the consumer client. Valid values:
	//
	// 	- **CLASSIC**: classic network.
	//
	// 	- **VPC**: VPC.
	//
	// example:
	//
	// VPC
	DestNetType *string `json:"DestNetType,omitempty" xml:"DestNetType,omitempty"`
	// The connection settings of the destination instance.
	DestinationEndpoint *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The DTS instance ID.
	//
	// example:
	//
	// dtsnjuc14kp12u****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The instance class.
	//
	// example:
	//
	// xlarge
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// > This parameter is returned only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	DtsJobDirection *string `json:"DtsJobDirection,omitempty" xml:"DtsJobDirection,omitempty"`
	// The DTS task ID.
	//
	// example:
	//
	// m06j1g92124****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The DTS instance name.
	//
	// example:
	//
	// dtstest****
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The end of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-26T14:03:21Z
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The operator information of the ETL task.
	//
	// > This parameter is returned only if you query the details of an ETL task.
	//
	// example:
	//
	// { "cells ": [{\\"shape\\":\\"edge\\",\\"attrs\\":{\\"line\\":{\\"stroke\\":\\"#b1b1b1\\",\\"strokeWidth\\":1,\\"targetMarker\\":{\\"name\\":\\"block\\",\\"args\\":{\\"size\\":\\"8\\"}},\\"strokeDasharray\\":\\"\\"}},\\"id\\":\\"cd1ec473-f9b9-4e9b-a742-ac23f442****\\",\\"source\\":{\\"cell\\":\\"8b261182-bfab-4803-ad8e-6bb08e3e****\\",\\"port\\":\\"out1\\"},\\"target\\":{\\"cell\\":\\"b36770df-f48c-4d6b-9644-54c5e924****\\",\\"port\\":\\"in1\\"},\\"zIndex\\":7 }] }
	EtlCalculator *string `json:"EtlCalculator,omitempty" xml:"EtlCalculator,omitempty"`
	// The time when the instance expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// > This parameter is returned only if the return value of **PayType*	- is **PrePaid**.
	//
	// example:
	//
	// 2023-06-16T08:01:19Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the task was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-16T10:34:17Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether the task is a subtask. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsDemoJob *bool `json:"IsDemoJob,omitempty" xml:"IsDemoJob,omitempty"`
	// The type of the DTS task. Valid values:
	//
	// 	- **online**: data migration task.
	//
	// 	- **SYNC**: data synchronization task.
	//
	// 	- **SUBSCRIBE**: change tracking task.
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
	// The migration types or initial synchronization types.
	MigrationMode *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// The minimum number of DUs.
	//
	// > This parameter is supported only for serverless instances.
	//
	// example:
	//
	// 1
	MinDu *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	// The source of the task. Valid values:
	//
	// 	- **PTS**
	//
	// 	- **DMS**
	//
	// 	- **DTS**
	//
	// example:
	//
	// PTS
	OriginType *string `json:"OriginType,omitempty" xml:"OriginType,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The performance of the data migration or synchronization instance.
	Performance *DescribeDtsJobDetailResponseBodySubSyncJobPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck state.
	PrecheckStatus *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet specific requirements, such as whether to automatically start a precheck. For more information, see [MigrationReserved](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {\\"srcHostPorts\\":\\"\\",\\"whitelist.dms.online.ddl.enable\\":false,\\"filterDDL\\":false,\\"sqlparser.dms.original.ddl\\":true,\\"srcOracleType\\":\\"sid\\",\\"maxRetryTime\\":43200,\\"destSSL\\":\\"0\\",\\"destOracleType\\":\\"sid\\",\\"srcSSL\\":\\"0\\",\\"dbListCaseChangeMode\\":\\"default\\",\\"SourceEngineVersion\\":\\"8.0.18\\",\\"srcNetType\\":\\"VPC\\",\\"destNetType\\":\\"VPC\\",\\"srcVpcNetMappingInst\\":\\"172.16.1**.**:10803\\",\\"destVpcNetMappingInst\\":\\"172.16.1**.**:11077\\",\\"useJobTask\\":\\"1\\"}
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The information about the retries performed by DTS due to an exception.
	RetryState *DescribeDtsJobDetailResponseBodySubSyncJobRetryState `json:"RetryState,omitempty" xml:"RetryState,omitempty" type:"Struct"`
	// The details of the data synchronization task in the reverse direction.
	//
	// > This parameter is returned only for two-way data synchronization tasks.
	ReverseJob *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob `json:"ReverseJob,omitempty" xml:"ReverseJob,omitempty" type:"Struct"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The state of initial schema synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Initialization status of library table structure.
	StructureInitializationStatus *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	// The information about the subtasks in the current data synchronization task.
	SubSyncJob []interface{} `json:"SubSyncJob,omitempty" xml:"SubSyncJob,omitempty" type:"Repeated"`
	// The topic of the change tracking instance.
	//
	// > This parameter is returned only if your change tracking instances are of the new version and you have called the [CreateConsumerGroup](https://help.aliyun.com/document_detail/122863.html) operation to create a consumer group.
	//
	// example:
	//
	// cn_hangzhou_rm_bp1162kryivb8****_dtstest_version2
	SubscribeTopic *string `json:"SubscribeTopic,omitempty" xml:"SubscribeTopic,omitempty"`
	// The type of data for change tracking.
	SubscriptionDataType *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	// The endpoint of the change tracking instance.
	SubscriptionHost *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost `json:"SubscriptionHost,omitempty" xml:"SubscriptionHost,omitempty" type:"Struct"`
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
	// 	- The value **Reverse*	- takes effect only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The tags of the task.
	TagList []*DescribeDtsJobDetailResponseBodySubSyncJobTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The task type.
	//
	// example:
	//
	// rds
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJob) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetAppName() *string {
	return s.AppName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetBeginTimestamp() *string {
	return s.BeginTimestamp
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetConsumptionCheckpoint() *string {
	return s.ConsumptionCheckpoint
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetConsumptionClient() *string {
	return s.ConsumptionClient
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDataEtlStatus() *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus {
	return s.DataEtlStatus
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDataInitializationStatus() *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDataSynchronizationStatus() *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDatabaseCount() *int32 {
	return s.DatabaseCount
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDbObject() *string {
	return s.DbObject
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDestNetType() *string {
	return s.DestNetType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDestinationEndpoint() *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDtsJobDirection() *string {
	return s.DtsJobDirection
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetEtlCalculator() *string {
	return s.EtlCalculator
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetIsDemoJob() *bool {
	return s.IsDemoJob
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetJobType() *string {
	return s.JobType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetMaxDu() *float64 {
	return s.MaxDu
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetMigrationMode() *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode {
	return s.MigrationMode
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetMinDu() *float64 {
	return s.MinDu
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetOriginType() *string {
	return s.OriginType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetPerformance() *DescribeDtsJobDetailResponseBodySubSyncJobPerformance {
	return s.Performance
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetPrecheckStatus() *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetReserved() *string {
	return s.Reserved
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetRetryState() *DescribeDtsJobDetailResponseBodySubSyncJobRetryState {
	return s.RetryState
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetReverseJob() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	return s.ReverseJob
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetSourceEndpoint() *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetStructureInitializationStatus() *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetSubSyncJob() []interface{} {
	return s.SubSyncJob
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetSubscribeTopic() *string {
	return s.SubscribeTopic
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetSubscriptionDataType() *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType {
	return s.SubscriptionDataType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetSubscriptionHost() *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost {
	return s.SubscriptionHost
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetTagList() []*DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	return s.TagList
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetAppName(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.AppName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetBeginTimestamp(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetCheckpoint(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.Checkpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetConsumptionCheckpoint(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetConsumptionClient(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetCreateTime(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDataEtlStatus(v *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.DataEtlStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDataInitializationStatus(v *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDataSynchronizationStatus(v *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDatabaseCount(v int32) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.DatabaseCount = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDbObject(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.DbObject = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDelay(v int64) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.Delay = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDestNetType(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.DestNetType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDestinationEndpoint(v *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDtsInstanceID(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDtsJobClass(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDtsJobDirection(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDtsJobId(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetDtsJobName(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetEndTimestamp(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetEtlCalculator(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.EtlCalculator = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetExpireTime(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetFinishTime(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.FinishTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetGroupId(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.GroupId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetIsDemoJob(v bool) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.IsDemoJob = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetJobType(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.JobType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetMaxDu(v float64) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.MaxDu = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetMigrationMode(v *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.MigrationMode = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetMinDu(v float64) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.MinDu = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetOriginType(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.OriginType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetPayType(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.PayType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetPerformance(v *DescribeDtsJobDetailResponseBodySubSyncJobPerformance) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.Performance = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetPrecheckStatus(v *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetReserved(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.Reserved = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetRetryState(v *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.RetryState = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetReverseJob(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.ReverseJob = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetSourceEndpoint(v *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetStructureInitializationStatus(v *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetSubSyncJob(v []interface{}) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.SubSyncJob = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetSubscribeTopic(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.SubscribeTopic = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetSubscriptionDataType(v *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.SubscriptionDataType = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetSubscriptionHost(v *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.SubscriptionHost = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetSynchronizationDirection(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetTagList(v []*DescribeDtsJobDetailResponseBodySubSyncJobTagList) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.TagList = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) SetTaskType(v string) *DescribeDtsJobDetailResponseBodySubSyncJob {
	s.TaskType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJob) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus struct {
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of full data migration or initial full data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that are migrated or synchronized during full data migration or initial full data synchronization.
	//
	// example:
	//
	// 16
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of the ETL task. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// 	- **Catched**: The task has no latency.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataEtlStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus struct {
	// The error message returned if full data migration or initial full data synchronization failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of full data migration or initial full data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that are migrated or synchronized during full data migration or initial full data synchronization.
	//
	// example:
	//
	// 16
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of full data migration or initial full data synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus struct {
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of incremental data migration or synchronization.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of rows and size of data that is synchronized or migrated to the destination table per second during incremental data synchronization or migration.
	//
	// example:
	//
	// 0.00RPS/(0.000MB/s)
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of incremental data migration or synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Suspending**: The task is paused.
	//
	// 	- **Checking**: The task is in precheck.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// 	- **Catched**: The task has no latency.
	//
	// example:
	//
	// Catched
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint struct {
	// The ID of the Alibaba Cloud account to which the destination instance belongs.
	//
	// example:
	//
	// 140692647406****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The name of the database to which the objects are migrated in the destination instance.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine of the destination instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The destination instance ID.
	//
	// example:
	//
	// rm-bp1f9guj5rhzq****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the destination instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The endpoint of the destination instance.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the return value of **EngineName*	- of the destination instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the destination instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the destination instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account to which the destination instance belongs.
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **DISABLE**: SSL encryption is disabled.
	//
	// 	- **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// 	- **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection to an AWS MongoDB Altas database.
	//
	// 	- **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection to a Kafka cluster.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The database account of the destination instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) SetAliyunUid(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	s.AliyunUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) SetDatabaseName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) SetEngineName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) SetInstanceID(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) SetInstanceType(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) SetIp(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) SetOracleSID(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) SetPort(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) SetRegion(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) SetRoleName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	s.RoleName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) SetUserName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode struct {
	// Indicates whether data transformation is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DataExtractTransformLoad *bool `json:"DataExtractTransformLoad,omitempty" xml:"DataExtractTransformLoad,omitempty"`
	// Indicates whether full data migration or initial full data synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Indicates whether incremental data migration or synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Indicates whether schema migration or initial schema synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode) GetDataExtractTransformLoad() *bool {
	return s.DataExtractTransformLoad
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode) SetDataExtractTransformLoad(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode {
	s.DataExtractTransformLoad = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode) SetDataInitialization(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode) SetDataSynchronization(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode) SetStructureInitialization(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobPerformance struct {
	// The size of data that is migrated or synchronized per second. Unit: Mbit/s.
	//
	// example:
	//
	// 1
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The number of times that SQL statements are migrated or synchronized per second, including BEGIN, COMMIT, DML, and DDL statements. DML statements include INSERT, DELETE, and UPDATE.
	//
	// example:
	//
	// 100
	Rps *string `json:"Rps,omitempty" xml:"Rps,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobPerformance) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPerformance) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPerformance) GetRps() *string {
	return s.Rps
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPerformance) SetFlow(v string) *DescribeDtsJobDetailResponseBodySubSyncJobPerformance {
	s.Flow = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPerformance) SetRps(v string) *DescribeDtsJobDetailResponseBodySubSyncJobPerformance {
	s.Rps = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPerformance) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus struct {
	// The result of each precheck item.
	Detail []*DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The precheck progress. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck state. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is in precheck.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus) GetDetail() []*DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus) SetDetail(v []*DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail struct {
	// The name of the precheck item.
	//
	// example:
	//
	// CHECK_CONN_SRC
	CheckItem *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	// The description of the precheck item.
	//
	// example:
	//
	// CHECK_CONN_SRC_DETAIL
	CheckItemDescription *string `json:"CheckItemDescription,omitempty" xml:"CheckItemDescription,omitempty"`
	// The precheck result. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Success
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The error message returned if the task failed to pass the precheck.
	//
	// > This parameter is returned only if the return value of **CheckResult*	- is **Failed**.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The method used to fix the precheck failure.
	//
	// > This parameter is returned only if the return value of **CheckResult*	- is **Failed**.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) GetCheckItemDescription() *string {
	return s.CheckItemDescription
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) GetCheckResult() *string {
	return s.CheckResult
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) GetFailedReason() *string {
	return s.FailedReason
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) SetCheckItem(v string) *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) SetCheckItemDescription(v string) *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) SetCheckResult(v string) *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) SetFailedReason(v string) *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail {
	s.FailedReason = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) SetRepairMethod(v string) *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobRetryState struct {
	// The error message returned.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ta7w132u12h****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum duration of a retry. Unit: seconds.
	//
	// example:
	//
	// 7200
	MaxRetryTime *int32 `json:"MaxRetryTime,omitempty" xml:"MaxRetryTime,omitempty"`
	// The progress of the instance when DTS performs retries.
	//
	// example:
	//
	// 03
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 5
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// The object on which the retries are performed. Valid values:
	//
	// 	- **srcDB**: the source database.
	//
	// 	- **destDB**: the destination database.
	//
	// 	- **inner_module**: an internal module of DTS.
	//
	// example:
	//
	// srcDB
	RetryTarget *string `json:"RetryTarget,omitempty" xml:"RetryTarget,omitempty"`
	// The time that has elapsed from the point in time when the first retry starts. Unit: seconds.
	//
	// example:
	//
	// 3600
	RetryTime *int32 `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	// Indicates whether the task is being retried. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Retrying *bool `json:"Retrying,omitempty" xml:"Retrying,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobRetryState) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobRetryState) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) GetMaxRetryTime() *int32 {
	return s.MaxRetryTime
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) GetModule() *string {
	return s.Module
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) GetRetryTarget() *string {
	return s.RetryTarget
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) GetRetryTime() *int32 {
	return s.RetryTime
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) GetRetrying() *bool {
	return s.Retrying
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) SetErrMsg(v string) *DescribeDtsJobDetailResponseBodySubSyncJobRetryState {
	s.ErrMsg = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) SetJobId(v string) *DescribeDtsJobDetailResponseBodySubSyncJobRetryState {
	s.JobId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) SetMaxRetryTime(v int32) *DescribeDtsJobDetailResponseBodySubSyncJobRetryState {
	s.MaxRetryTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) SetModule(v string) *DescribeDtsJobDetailResponseBodySubSyncJobRetryState {
	s.Module = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) SetRetryCount(v int32) *DescribeDtsJobDetailResponseBodySubSyncJobRetryState {
	s.RetryCount = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) SetRetryTarget(v string) *DescribeDtsJobDetailResponseBodySubSyncJobRetryState {
	s.RetryTarget = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) SetRetryTime(v int32) *DescribeDtsJobDetailResponseBodySubSyncJobRetryState {
	s.RetryTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) SetRetrying(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobRetryState {
	s.Retrying = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobRetryState) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJob struct {
	// Indicates whether the new change tracking feature is used.
	//
	// > This parameter is returned only for change tracking instances of the new version.
	//
	// example:
	//
	// new
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The start of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-15T08:25:34Z
	BeginTimestamp *string `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	// The start offset of incremental data migration or data synchronization. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1616405159
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The consumption checkpoint of the change tracking instance. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-23T07:30:31Z
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	// The downstream client information in the following format: \\<IP address of the downstream client>:\\<Random ID generated by DTS>.
	//
	// example:
	//
	// 114.***.***.**:dts********
	ConsumptionClient *string `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-01-12T08:34:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The state of the ETL task.
	//
	// > This parameter collection is returned only if an ETL task is configured.
	DataEtlStatus *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus `json:"DataEtlStatus,omitempty" xml:"DataEtlStatus,omitempty" type:"Struct"`
	// The state of full data migration or initial full data synchronization.
	DataInitializationStatus *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The state of incremental data migration or synchronization.
	DataSynchronizationStatus *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The number of ApsaraDB RDS for MySQL instances that are attached to the source PolarDB-X 1.0 instance.
	//
	// example:
	//
	// 2
	DatabaseCount *int32 `json:"DatabaseCount,omitempty" xml:"DatabaseCount,omitempty"`
	// The objects of the data migration, data synchronization, or change tracking task. For more information, see [Objects of DTS tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// example:
	//
	// {\\"dtstestdata\\":{\\"all\\":true,\\"name\\":\\"dtstestdata\\",\\"state\\":\\"normal\\"}}
	DbObject *string `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	// The latency of incremental data migration or synchronization. Unit: milliseconds.
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The network type of the consumer client. Valid values:
	//
	// 	- **CLASSIC**: classic network.
	//
	// 	- **VPC**: VPC.
	//
	// example:
	//
	// VPC
	DestNetType *string `json:"DestNetType,omitempty" xml:"DestNetType,omitempty"`
	// The connection settings of the destination instance.
	DestinationEndpoint *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The DTS instance ID.
	//
	// example:
	//
	// dtsnjuc14kp12u****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The instance class.
	//
	// example:
	//
	// xlarge
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// > This parameter is returned only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	DtsJobDirection *string `json:"DtsJobDirection,omitempty" xml:"DtsJobDirection,omitempty"`
	// The DTS task ID.
	//
	// example:
	//
	// m06j1g92124****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The DTS instance name.
	//
	// example:
	//
	// dtstest****
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The end of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-26T14:03:21Z
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The operator information of the ETL task.
	//
	// > This parameter is returned only if you query the details of an ETL task.
	//
	// example:
	//
	// { "cells ": [{\\"shape\\":\\"edge\\",\\"attrs\\":{\\"line\\":{\\"stroke\\":\\"#b1b1b1\\",\\"strokeWidth\\":1,\\"targetMarker\\":{\\"name\\":\\"block\\",\\"args\\":{\\"size\\":\\"8\\"}},\\"strokeDasharray\\":\\"\\"}},\\"id\\":\\"cd1ec473-f9b9-4e9b-a742-ac23f442****\\",\\"source\\":{\\"cell\\":\\"8b261182-bfab-4803-ad8e-6bb08e3e****\\",\\"port\\":\\"out1\\"},\\"target\\":{\\"cell\\":\\"b36770df-f48c-4d6b-9644-54c5e924****\\",\\"port\\":\\"in1\\"},\\"zIndex\\":7 }] }
	EtlCalculator *string `json:"EtlCalculator,omitempty" xml:"EtlCalculator,omitempty"`
	// The time when the instance expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// > This parameter is returned only if the return value of **PayType*	- is **PrePaid**.
	//
	// example:
	//
	// 2023-06-16T08:01:19Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the task was complete. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-16T10:34:17Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Indicates whether the task is a subtask. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsDemoJob *bool `json:"IsDemoJob,omitempty" xml:"IsDemoJob,omitempty"`
	// The type of the DTS task. Valid values:
	//
	// 	- **online**: data migration task.
	//
	// 	- **SYNC**: data synchronization task.
	//
	// 	- **SUBSCRIBE**: change tracking task.
	//
	// example:
	//
	// SYNC
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The migration types or initial synchronization types.
	MigrationMode *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// The source of the task. Valid values:
	//
	// 	- **PTS**
	//
	// 	- **DMS**
	//
	// 	- **DTS**
	//
	// example:
	//
	// DTS
	OriginType *string `json:"OriginType,omitempty" xml:"OriginType,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The performance of the data migration or synchronization instance.
	Performance *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck state.
	PrecheckStatus *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet specific requirements, such as whether to automatically start a precheck. For more information, see [MigrationReserved](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {\\"srcHostPorts\\":\\"\\",\\"whitelist.dms.online.ddl.enable\\":false,\\"filterDDL\\":false,\\"sqlparser.dms.original.ddl\\":true,\\"srcOracleType\\":\\"sid\\",\\"maxRetryTime\\":43200,\\"destSSL\\":\\"0\\",\\"destOracleType\\":\\"sid\\",\\"srcSSL\\":\\"0\\",\\"dbListCaseChangeMode\\":\\"default\\",\\"SourceEngineVersion\\":\\"8.0.18\\",\\"srcNetType\\":\\"VPC\\",\\"destNetType\\":\\"VPC\\",\\"srcVpcNetMappingInst\\":\\"172.16.1**.**:10803\\",\\"destVpcNetMappingInst\\":\\"172.16.1**.**:11077\\",\\"useJobTask\\":\\"1\\"}
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The information about the retries performed by DTS due to an exception.
	RetryState *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState `json:"RetryState,omitempty" xml:"RetryState,omitempty" type:"Struct"`
	// The details of the data synchronization task in the reverse direction.
	//
	// > This parameter is returned only for two-way data synchronization tasks.
	//
	// example:
	//
	// ****
	ReverseJob interface{} `json:"ReverseJob,omitempty" xml:"ReverseJob,omitempty"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The state of initial schema synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Initialization status of library table structure.
	StructureInitializationStatus *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	// The topic of the change tracking instance.
	//
	// > This parameter is returned only if your change tracking instances are of the new version and you have called the [CreateConsumerGroup](https://help.aliyun.com/document_detail/122863.html) operation to create a consumer group.
	//
	// example:
	//
	// cn_hangzhou_rm_bp1162kryivb8****_dtstest_version2
	SubscribeTopic *string `json:"SubscribeTopic,omitempty" xml:"SubscribeTopic,omitempty"`
	// The type of data for change tracking.
	SubscriptionDataType *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	// The endpoint of the change tracking instance.
	SubscriptionHost *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost `json:"SubscriptionHost,omitempty" xml:"SubscriptionHost,omitempty" type:"Struct"`
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
	// 	- The value **Reverse*	- takes effect only if the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The tags of the task.
	TagList []*DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The task type.
	//
	// example:
	//
	// rds
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetAppName() *string {
	return s.AppName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetBeginTimestamp() *string {
	return s.BeginTimestamp
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetConsumptionCheckpoint() *string {
	return s.ConsumptionCheckpoint
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetConsumptionClient() *string {
	return s.ConsumptionClient
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDataEtlStatus() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus {
	return s.DataEtlStatus
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDataInitializationStatus() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDataSynchronizationStatus() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDatabaseCount() *int32 {
	return s.DatabaseCount
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDbObject() *string {
	return s.DbObject
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDestNetType() *string {
	return s.DestNetType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDestinationEndpoint() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDtsJobDirection() *string {
	return s.DtsJobDirection
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetEtlCalculator() *string {
	return s.EtlCalculator
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetIsDemoJob() *bool {
	return s.IsDemoJob
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetJobType() *string {
	return s.JobType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetMigrationMode() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode {
	return s.MigrationMode
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetOriginType() *string {
	return s.OriginType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetPerformance() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance {
	return s.Performance
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetPrecheckStatus() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetReserved() *string {
	return s.Reserved
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetRetryState() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState {
	return s.RetryState
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetReverseJob() interface{} {
	return s.ReverseJob
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetSourceEndpoint() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetStructureInitializationStatus() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetSubscribeTopic() *string {
	return s.SubscribeTopic
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetSubscriptionDataType() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType {
	return s.SubscriptionDataType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetSubscriptionHost() *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost {
	return s.SubscriptionHost
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetTagList() []*DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	return s.TagList
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetAppName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.AppName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetBeginTimestamp(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetCheckpoint(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.Checkpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetConsumptionCheckpoint(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetConsumptionClient(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetCreateTime(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDataEtlStatus(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.DataEtlStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDataInitializationStatus(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDataSynchronizationStatus(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDatabaseCount(v int32) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.DatabaseCount = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDbObject(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.DbObject = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDelay(v int64) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.Delay = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDestNetType(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.DestNetType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDestinationEndpoint(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDtsInstanceID(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDtsJobClass(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDtsJobDirection(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDtsJobId(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetDtsJobName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetEndTimestamp(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetEtlCalculator(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.EtlCalculator = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetExpireTime(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetFinishTime(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.FinishTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetGroupId(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.GroupId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetIsDemoJob(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.IsDemoJob = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetJobType(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.JobType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetMigrationMode(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.MigrationMode = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetOriginType(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.OriginType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetPayType(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.PayType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetPerformance(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.Performance = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetPrecheckStatus(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetReserved(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.Reserved = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetRetryState(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.RetryState = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetReverseJob(v interface{}) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.ReverseJob = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetSourceEndpoint(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetStructureInitializationStatus(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetSubscribeTopic(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.SubscribeTopic = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetSubscriptionDataType(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.SubscriptionDataType = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetSubscriptionHost(v *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.SubscriptionHost = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetSynchronizationDirection(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetTagList(v []*DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.TagList = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) SetTaskType(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob {
	s.TaskType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJob) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus struct {
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of full data migration or initial full data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that are migrated or synchronized during full data migration or initial full data synchronization.
	//
	// example:
	//
	// 16
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of the ETL task. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// 	- **Catched**: The task has no latency.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataEtlStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus struct {
	// The error message returned if full data migration or initial full data synchronization failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of full data migration or initial full data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that are migrated or synchronized during full data migration or initial full data synchronization.
	//
	// example:
	//
	// 16
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of full data migration or initial full data synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus struct {
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance class needs to be upgraded. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of incremental data migration or synchronization.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of rows and size of data that is synchronized or migrated to the destination table per second during incremental data synchronization or migration.
	//
	// example:
	//
	// 0.00RPS/(0.000MB/s)
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of incremental data migration or synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Suspending**: The task is paused.
	//
	// 	- **Checking**: The task is in precheck.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// 	- **Catched**: The task has no latency.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint struct {
	// The ID of the Alibaba Cloud account to which the destination instance belongs.
	//
	// example:
	//
	// 140692647406****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The name of the database to which the objects are migrated in the destination instance.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine of the destination instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The destination instance ID.
	//
	// example:
	//
	// rm-bp1f9guj5rhzq****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the destination instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The endpoint of the destination instance.
	//
	// example:
	//
	// 192.168.XX,XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the return value of **EngineName*	- of the destination instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the destination instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the destination instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account to which the destination instance belongs.
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **DISABLE**: SSL encryption is disabled.
	//
	// 	- **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// 	- **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection to an AWS MongoDB Altas database.
	//
	// 	- **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection to a Kafka cluster.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The database account of the destination instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) SetAliyunUid(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	s.AliyunUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) SetDatabaseName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) SetEngineName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) SetInstanceID(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) SetInstanceType(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) SetIp(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) SetOracleSID(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) SetPort(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) SetRegion(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) SetRoleName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	s.RoleName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) SetUserName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode struct {
	// Indicates whether data transformation is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DataExtractTransformLoad *bool `json:"DataExtractTransformLoad,omitempty" xml:"DataExtractTransformLoad,omitempty"`
	// Indicates whether full data migration or initial full data synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Indicates whether incremental data migration or synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Indicates whether schema migration or initial schema synchronization is performed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode) GetDataExtractTransformLoad() *bool {
	return s.DataExtractTransformLoad
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode) SetDataExtractTransformLoad(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode {
	s.DataExtractTransformLoad = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode) SetDataInitialization(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode) SetDataSynchronization(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode) SetStructureInitialization(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance struct {
	// The size of data that is migrated or synchronized per second. Unit: Mbit/s.
	//
	// example:
	//
	// 1
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The number of times that SQL statements are migrated or synchronized per second, including BEGIN, COMMIT, DML, and DDL statements. DML statements include INSERT, DELETE, and UPDATE.
	//
	// example:
	//
	// 100
	Rps *string `json:"Rps,omitempty" xml:"Rps,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance) GetRps() *string {
	return s.Rps
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance) SetFlow(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance {
	s.Flow = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance) SetRps(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance {
	s.Rps = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPerformance) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus struct {
	// The result of each precheck item.
	Detail []*DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The precheck progress. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck state. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is in precheck.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus) GetDetail() []*DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus) SetDetail(v []*DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail struct {
	// The name of the precheck item.
	//
	// example:
	//
	// CHECK_CONN_SRC
	CheckItem *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	// The description of the precheck item.
	//
	// example:
	//
	// CHECK_CONN_SRC_DETAIL
	CheckItemDescription *string `json:"CheckItemDescription,omitempty" xml:"CheckItemDescription,omitempty"`
	// The precheck result. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Success
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The error message returned if the task failed to pass the precheck.
	//
	// > This parameter is returned only if the return value of **CheckResult*	- is **Failed**.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The method used to fix the precheck failure.
	//
	// > This parameter is returned only if the return value of **CheckResult*	- is **Failed**.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) GetCheckItemDescription() *string {
	return s.CheckItemDescription
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) GetCheckResult() *string {
	return s.CheckResult
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) GetFailedReason() *string {
	return s.FailedReason
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) SetCheckItem(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) SetCheckItemDescription(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) SetCheckResult(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) SetFailedReason(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail {
	s.FailedReason = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) SetRepairMethod(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState struct {
	// The error message returned.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ta7w132u12h****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum duration of a retry. Unit: seconds.
	//
	// example:
	//
	// 7200
	MaxRetryTime *int32 `json:"MaxRetryTime,omitempty" xml:"MaxRetryTime,omitempty"`
	// The progress of the instance when DTS performs retries.
	//
	// example:
	//
	// 03
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 5
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// The object on which the retries are performed. Valid values:
	//
	// 	- **srcDB**: the source database.
	//
	// 	- **destDB**: the destination database.
	//
	// 	- **inner_module**: an internal module of DTS.
	//
	// example:
	//
	// srcDB
	RetryTarget *string `json:"RetryTarget,omitempty" xml:"RetryTarget,omitempty"`
	// The time that has elapsed from the point in time when the first retry starts. Unit: seconds.
	//
	// example:
	//
	// 3600
	RetryTime *int32 `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	// Indicates whether the task is being retried. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Retrying *bool `json:"Retrying,omitempty" xml:"Retrying,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) GetMaxRetryTime() *int32 {
	return s.MaxRetryTime
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) GetModule() *string {
	return s.Module
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) GetRetryTarget() *string {
	return s.RetryTarget
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) GetRetryTime() *int32 {
	return s.RetryTime
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) GetRetrying() *bool {
	return s.Retrying
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) SetErrMsg(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState {
	s.ErrMsg = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) SetJobId(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState {
	s.JobId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) SetMaxRetryTime(v int32) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState {
	s.MaxRetryTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) SetModule(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState {
	s.Module = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) SetRetryCount(v int32) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState {
	s.RetryCount = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) SetRetryTarget(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState {
	s.RetryTarget = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) SetRetryTime(v int32) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState {
	s.RetryTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) SetRetrying(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState {
	s.Retrying = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobRetryState) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint struct {
	// The ID of the Alibaba Cloud account to which the source instance belongs.
	//
	// example:
	//
	// 140692647406****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The name of the database from which the objects are migrated in the source instance.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine of the source instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The DTS instance ID.
	//
	// example:
	//
	// rm-bp2f3huj5rhzq****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the source instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The endpoint of the source instance.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the return value of **EngineName*	- of the destination instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the source instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the source instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account to which the source instance belongs.
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **DISABLE**: SSL encryption is disabled.
	//
	// 	- **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// 	- **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection to an AWS MongoDB Altas database.
	//
	// 	- **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection to a Kafka cluster.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The database account of the source instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) SetAliyunUid(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	s.AliyunUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) SetDatabaseName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) SetEngineName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) SetInstanceID(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) SetInstanceType(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) SetIp(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) SetOracleSID(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) SetPort(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) SetRegion(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) SetRoleName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	s.RoleName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) SetUserName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus struct {
	// Error message indicating task failure.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Whether to display upgrade specifications, return value:
	//
	// - True: Yes.
	//
	// - False: No.
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// Initialization progress of library table structure, measured in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables that have completed library table structure initialization.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The initialization status of the library table structure includes:
	//
	// - NotStarted: Not started.
	//
	// - Migration: In the process of initialization.
	//
	// - Failed: Initialization failed.
	//
	// - Finished: Initialization completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType struct {
	// Indicates whether DDL statements are tracked. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Ddl *bool `json:"Ddl,omitempty" xml:"Ddl,omitempty"`
	// Indicates whether DML statements are tracked. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Dml *bool `json:"Dml,omitempty" xml:"Dml,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType) GetDdl() *bool {
	return s.Ddl
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType) GetDml() *bool {
	return s.Dml
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType) SetDdl(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType {
	s.Ddl = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType) SetDml(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType {
	s.Dml = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionDataType) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost struct {
	// The private endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****-internal.aliyuncs.com:18002
	PrivateHost *string `json:"PrivateHost,omitempty" xml:"PrivateHost,omitempty"`
	// The public endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****.aliyuncs.com:18001
	PublicHost *string `json:"PublicHost,omitempty" xml:"PublicHost,omitempty"`
	// The VPC endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****.aliyuncs.com:18001
	VpcHost *string `json:"VpcHost,omitempty" xml:"VpcHost,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost) GetPrivateHost() *string {
	return s.PrivateHost
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost) GetPublicHost() *string {
	return s.PublicHost
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost) GetVpcHost() *string {
	return s.VpcHost
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost) SetPrivateHost(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost {
	s.PrivateHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost) SetPublicHost(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost {
	s.PublicHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost) SetVpcHost(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost {
	s.VpcHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobSubscriptionHost) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 191448876515****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The operator of the tag.
	//
	// example:
	//
	// 191448876515****
	Creator *int64 `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the task was modified.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary key of the table.
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the region in which the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking instance.
	//
	// example:
	//
	// dtsnjuc14kp12u****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::DTS::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Indicates whether the tag is visible. Valid values:
	//
	// 	- **0**: The tag is public.
	//
	// 	- **1**: The tag is private.
	//
	// example:
	//
	// 0
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The ID of the region in which the DTS task resides.
	//
	// > In most cases, the ID of the region in which the destination instance resides is returned.
	//
	// example:
	//
	// cn-hangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The type of the tag. Valid values:
	//
	// 	- **System**: The tag was created by the system.
	//
	// 	- **Custom**: The tag was created by a user.
	//
	// > By default, if the parameter is left empty, custom tags and system tags are returned.
	//
	// example:
	//
	// System
	TagCategory *string `json:"TagCategory,omitempty" xml:"TagCategory,omitempty"`
	// The tag key.
	//
	// example:
	//
	// key1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetCreator() *int64 {
	return s.Creator
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetId() *int64 {
	return s.Id
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetScope() *string {
	return s.Scope
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetTagCategory() *string {
	return s.TagCategory
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetAliUid(v int64) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.AliUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetCreator(v int64) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.Creator = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetGmtCreate(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetGmtModified(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.GmtModified = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetId(v int64) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.Id = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetRegionId(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.RegionId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetResourceId(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.ResourceId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetResourceType(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.ResourceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetScope(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.Scope = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetSrcRegion(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.SrcRegion = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetTagCategory(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.TagCategory = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetTagKey(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) SetTagValue(v string) *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList {
	s.TagValue = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobReverseJobTagList) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint struct {
	// The ID of the Alibaba Cloud account to which the source instance belongs.
	//
	// example:
	//
	// 140692647406****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The name of the database from which the objects are migrated in the source instance.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine of the source instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The source instance ID.
	//
	// example:
	//
	// rm-bp2f3huj5rhzq****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The type of the source instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The endpoint of the source instance.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the return value of **EngineName*	- of the destination instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the source instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the source instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the RAM role configured for the Alibaba Cloud account to which the source instance belongs.
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **DISABLE**: SSL encryption is disabled.
	//
	// 	- **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// 	- **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection to an AWS MongoDB Altas database.
	//
	// 	- **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection to a Kafka cluster.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The database account of the source instance.
	//
	// example:
	//
	// dtstest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GetRoleName() *string {
	return s.RoleName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) SetAliyunUid(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	s.AliyunUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) SetDatabaseName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) SetEngineName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) SetInstanceID(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) SetInstanceType(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) SetIp(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) SetOracleSID(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) SetPort(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) SetRegion(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) SetRoleName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	s.RoleName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) SetUserName(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus struct {
	// Error message indicating task failure.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Whether to display upgrade specifications, return value:
	//
	// - True: Yes.
	//
	// - False: No.
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// Initialization progress of library table structure, measured in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables that have completed library table structure initialization.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The initialization status of the library table structure includes:
	//
	// - NotStarted: Not started.
	//
	// - Migration: In the process of initialization.
	//
	// - Failed: Initialization failed.
	//
	// - Finished: Initialization completed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) SetPercent(v string) *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) SetProgress(v string) *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) SetStatus(v string) *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType struct {
	// Indicates whether DDL statements are tracked. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Ddl *bool `json:"Ddl,omitempty" xml:"Ddl,omitempty"`
	// Indicates whether DML statements are tracked. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Dml *bool `json:"Dml,omitempty" xml:"Dml,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType) GetDdl() *bool {
	return s.Ddl
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType) GetDml() *bool {
	return s.Dml
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType) SetDdl(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType {
	s.Ddl = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType) SetDml(v bool) *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType {
	s.Dml = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionDataType) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost struct {
	// The private endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****-internal.aliyuncs.com:18002
	PrivateHost *string `json:"PrivateHost,omitempty" xml:"PrivateHost,omitempty"`
	// The public endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****.aliyuncs.com:18001
	PublicHost *string `json:"PublicHost,omitempty" xml:"PublicHost,omitempty"`
	// The VPC endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****.aliyuncs.com:18001
	VpcHost *string `json:"VpcHost,omitempty" xml:"VpcHost,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost) GetPrivateHost() *string {
	return s.PrivateHost
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost) GetPublicHost() *string {
	return s.PublicHost
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost) GetVpcHost() *string {
	return s.VpcHost
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost) SetPrivateHost(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost {
	s.PrivateHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost) SetPublicHost(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost {
	s.PublicHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost) SetVpcHost(v string) *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost {
	s.VpcHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobSubscriptionHost) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubSyncJobTagList struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 191448876515****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The operator of the tag.
	//
	// example:
	//
	// 191448876515****
	Creator *int64 `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the task was modified.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary key of the table.
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the region in which the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking instance.
	//
	// example:
	//
	// dtsnjuc14kp12u****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::DTS::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Indicates whether the tag is visible. Valid values:
	//
	// 	- **0**: The tag is public.
	//
	// 	- **1**: The tag is private.
	//
	// example:
	//
	// 0
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The ID of the region in which the DTS task resides.
	//
	// > In most cases, the ID of the region in which the destination instance resides is returned.
	//
	// example:
	//
	// cn-hangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// The type of the tag. Valid values:
	//
	// 	- **System**: The tag was created by the system.
	//
	// 	- **Custom**: The tag was created by a user.
	//
	// > By default, if the parameter is left empty, custom tags and system tags are returned.
	//
	// example:
	//
	// System
	TagCategory *string `json:"TagCategory,omitempty" xml:"TagCategory,omitempty"`
	// The tag key.
	//
	// example:
	//
	// key1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubSyncJobTagList) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetCreator() *int64 {
	return s.Creator
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetId() *int64 {
	return s.Id
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetScope() *string {
	return s.Scope
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetTagCategory() *string {
	return s.TagCategory
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetAliUid(v int64) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.AliUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetCreator(v int64) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.Creator = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetGmtCreate(v string) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetGmtModified(v string) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.GmtModified = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetId(v int64) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.Id = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetRegionId(v string) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.RegionId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetResourceId(v string) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.ResourceId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetResourceType(v string) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.ResourceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetScope(v string) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.Scope = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetSrcRegion(v string) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.SrcRegion = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetTagCategory(v string) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.TagCategory = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetTagKey(v string) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) SetTagValue(v string) *DescribeDtsJobDetailResponseBodySubSyncJobTagList {
	s.TagValue = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubSyncJobTagList) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubscriptionDataType struct {
	// Indicates whether DDL statements are tracked. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Ddl *bool `json:"Ddl,omitempty" xml:"Ddl,omitempty"`
	// Indicates whether DML statements are tracked. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Dml *bool `json:"Dml,omitempty" xml:"Dml,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubscriptionDataType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubscriptionDataType) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionDataType) GetDdl() *bool {
	return s.Ddl
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionDataType) GetDml() *bool {
	return s.Dml
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionDataType) SetDdl(v bool) *DescribeDtsJobDetailResponseBodySubscriptionDataType {
	s.Ddl = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionDataType) SetDml(v bool) *DescribeDtsJobDetailResponseBodySubscriptionDataType {
	s.Dml = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionDataType) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobDetailResponseBodySubscriptionHost struct {
	// The private endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****-internal.aliyuncs.com:18002
	PrivateHost *string `json:"PrivateHost,omitempty" xml:"PrivateHost,omitempty"`
	// The public endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****.aliyuncs.com:18001
	PublicHost *string `json:"PublicHost,omitempty" xml:"PublicHost,omitempty"`
	// The VPC endpoint of the change tracking instance. The format is `<Address>:<Port number>`.
	//
	// example:
	//
	// dts-cn-****-vpc.aliyuncs.com:18003
	VpcHost *string `json:"VpcHost,omitempty" xml:"VpcHost,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubscriptionHost) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubscriptionHost) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionHost) GetPrivateHost() *string {
	return s.PrivateHost
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionHost) GetPublicHost() *string {
	return s.PublicHost
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionHost) GetVpcHost() *string {
	return s.VpcHost
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionHost) SetPrivateHost(v string) *DescribeDtsJobDetailResponseBodySubscriptionHost {
	s.PrivateHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionHost) SetPublicHost(v string) *DescribeDtsJobDetailResponseBodySubscriptionHost {
	s.PublicHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionHost) SetVpcHost(v string) *DescribeDtsJobDetailResponseBodySubscriptionHost {
	s.VpcHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionHost) Validate() error {
	return dara.Validate(s)
}
