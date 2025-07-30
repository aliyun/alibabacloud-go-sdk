// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobList(v []*DescribeDtsJobsResponseBodyDtsJobList) *DescribeDtsJobsResponseBody
	GetDtsJobList() []*DescribeDtsJobsResponseBodyDtsJobList
	SetDynamicCode(v string) *DescribeDtsJobsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeDtsJobsResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeDtsJobsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDtsJobsResponseBody
	GetErrMessage() *string
	SetEtlDemoList(v []*DescribeDtsJobsResponseBodyEtlDemoList) *DescribeDtsJobsResponseBody
	GetEtlDemoList() []*DescribeDtsJobsResponseBodyEtlDemoList
	SetHttpStatusCode(v int32) *DescribeDtsJobsResponseBody
	GetHttpStatusCode() *int32
	SetPageNumber(v int32) *DescribeDtsJobsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeDtsJobsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeDtsJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDtsJobsResponseBody
	GetSuccess() *bool
	SetTotalRecordCount(v int32) *DescribeDtsJobsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDtsJobsResponseBody struct {
	// The Data Transmission Service (DTS) tasks and the details of each task.
	DtsJobList []*DescribeDtsJobsResponseBodyDtsJobList `json:"DtsJobList,omitempty" xml:"DtsJobList,omitempty" type:"Repeated"`
	// The dynamic error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. The value of this parameter is used to replace the **%s*	- variable in the value of the **ErrMessage*	- parameter.
	//
	// >  For example, if the value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the value of the **DynamicMessage*	- parameter is **Type**, the specified **Type*	- parameter is invalid.
	//
	// example:
	//
	// Type
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The DTS tasks and the details of each task.
	EtlDemoList []*DescribeDtsJobsResponseBodyEtlDemoList `json:"EtlDemoList,omitempty" xml:"EtlDemoList,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of DTS tasks that meet the query condition.
	//
	// example:
	//
	// 15
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDtsJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBody) GetDtsJobList() []*DescribeDtsJobsResponseBodyDtsJobList {
	return s.DtsJobList
}

func (s *DescribeDtsJobsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeDtsJobsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDtsJobsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDtsJobsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDtsJobsResponseBody) GetEtlDemoList() []*DescribeDtsJobsResponseBodyEtlDemoList {
	return s.EtlDemoList
}

func (s *DescribeDtsJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDtsJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDtsJobsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeDtsJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDtsJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDtsJobsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDtsJobsResponseBody) SetDtsJobList(v []*DescribeDtsJobsResponseBodyDtsJobList) *DescribeDtsJobsResponseBody {
	s.DtsJobList = v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetDynamicCode(v string) *DescribeDtsJobsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetDynamicMessage(v string) *DescribeDtsJobsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetErrCode(v string) *DescribeDtsJobsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetErrMessage(v string) *DescribeDtsJobsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetEtlDemoList(v []*DescribeDtsJobsResponseBodyEtlDemoList) *DescribeDtsJobsResponseBody {
	s.EtlDemoList = v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetHttpStatusCode(v int32) *DescribeDtsJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetPageNumber(v int32) *DescribeDtsJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetPageRecordCount(v int32) *DescribeDtsJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetRequestId(v string) *DescribeDtsJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetSuccess(v bool) *DescribeDtsJobsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetTotalRecordCount(v int32) *DescribeDtsJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobList struct {
	// Indicates whether the **new*	- change tracking feature is used.
	//
	// >  This parameter is returned only for change tracking instances of the new version.
	//
	// example:
	//
	// new
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The start of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-15T08:25:34Z
	BeginTimestamp *string `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	// The start offset of incremental data synchronization. The value is a UNIX timestamp representing the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1616899019
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The consumption checkpoint of the change tracking instance. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-23T07:30:31Z
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	// The downstream client information, in the following format: \\<IP address of the downstream client>:\\<Random ID generated by DTS>.
	//
	// example:
	//
	// 114...:dts******
	ConsumptionClient *string `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	// The CPU utilization of the instance. Unit: percentage.
	//
	// example:
	//
	// 1
	CpuUsage *string `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty"`
	// The point in time when the task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-16T08:01:19Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The state of the physical gateway-based migration task.
	DataCloudStatus *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus `json:"DataCloudStatus,omitempty" xml:"DataCloudStatus,omitempty" type:"Struct"`
	// The state of the extract, transform, and load (ETL) task. Valid values:
	//
	// >  This parameter collection is returned only if an ETL task is configured.
	DataEtlStatus *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus `json:"DataEtlStatus,omitempty" xml:"DataEtlStatus,omitempty" type:"Struct"`
	// The state of full data synchronization.
	DataInitializationStatus *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The state of incremental data migration or synchronization.
	DataSynchronizationStatus *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The objects that you want to synchronize. The value is a JSON string and can contain regular expressions. For more information, see "Objects of DTS tasks".
	//
	// example:
	//
	// {"dtstestdata": { "name": "dtstestdata", "all": true }}
	DbObject *string `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	// The ID of the DTS dedicated cluster on which a DTS task runs.
	//
	// example:
	//
	// dtscluster_ft7y3**********
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The latency of incremental data synchronization. Unit: seconds.
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The connection settings of the destination instance.
	DestinationEndpoint *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The environment tag of the DTS instance. Valid values:
	//
	// - **normal**
	//
	// - **online**
	//
	// example:
	//
	// normal
	DtsBisLabel *string `json:"DtsBisLabel,omitempty" xml:"DtsBisLabel,omitempty"`
	// The ID of the data synchronization instance.
	//
	// example:
	//
	// dtsi03e3zty16i****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The instance class.
	//
	// >  For more information about the test performance of each instance class, see [Specifications of data synchronization instances](https://help.aliyun.com/document_detail/26605.html).
	//
	// example:
	//
	// large
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The synchronization direction. The value is **Reverse**.
	//
	// example:
	//
	// Forward
	DtsJobDirection *string `json:"DtsJobDirection,omitempty" xml:"DtsJobDirection,omitempty"`
	// The ID of the data synchronization task.
	//
	// example:
	//
	// i03e3zty16i****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The name of the data synchronization task.
	//
	// example:
	//
	// RDS_TO_RDS_MIGRATION
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The DTS Units (DUs) usage of a task in a DTS dedicated cluster.
	//
	// example:
	//
	// 12.0%
	DuRealUsage *string `json:"DuRealUsage,omitempty" xml:"DuRealUsage,omitempty"`
	// The number of DUs that have been used.
	//
	// example:
	//
	// 15
	DuUsage *int64 `json:"DuUsage,omitempty" xml:"DuUsage,omitempty"`
	// The end of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-26T14:03:21Z
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The error message returned.
	ErrorDetails []*DescribeDtsJobsResponseBodyDtsJobListErrorDetails `json:"ErrorDetails,omitempty" xml:"ErrorDetails,omitempty" type:"Repeated"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by [com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException:Could not create connection to database server. Attempted reconnect 3 times. Giving up.][com.mysql.jdbc.exceptions.jdbc4.CommunicationsException:Communications link failure\\n\\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.][java.net.ConnectException:Connection timed out (Connection timed out)] About more information in [https://yq.aliyun.com/articles/499178].
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The checkpoint of the ETL task.
	//
	// example:
	//
	// 1610540493
	EtlSafeCheckpoint *string `json:"EtlSafeCheckpoint,omitempty" xml:"EtlSafeCheckpoint,omitempty"`
	// The point in time when the instance expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// >  This parameter is returned only if the value of the **PayType*	- parameter is **PrePaid**.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The state information about the full data verification task.
	FullDataCheckStatus *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus `json:"FullDataCheckStatus,omitempty" xml:"FullDataCheckStatus,omitempty" type:"Struct"`
	// The state information about the incremental data verification task.
	IncDataCheckStatus *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus `json:"IncDataCheckStatus,omitempty" xml:"IncDataCheckStatus,omitempty" type:"Struct"`
	// The type of the DTS task. Valid values:
	//
	// - **MIGRATION**: data migration task
	//
	// - **SYNC**: data synchronization task
	//
	// - **SUBSCRIBE**: change tracking task
	//
	// example:
	//
	// MIGRATION
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// Upper limit of DU.
	//
	// > Only supported by Serverless instances.
	//
	// example:
	//
	// 16
	MaxDu *float64 `json:"MaxDu,omitempty" xml:"MaxDu,omitempty"`
	// The memory that has been used. Unit: MB.
	//
	// example:
	//
	// 500
	MemUsage *string `json:"MemUsage,omitempty" xml:"MemUsage,omitempty"`
	// The error code.
	//
	// example:
	//
	// dts.retry.err.0046
	MigrationErrCode *string `json:"MigrationErrCode,omitempty" xml:"MigrationErrCode,omitempty"`
	// The ID of the error code-related documentation.
	//
	// example:
	//
	// 462133
	MigrationErrHelpDocId *string `json:"MigrationErrHelpDocId,omitempty" xml:"MigrationErrHelpDocId,omitempty"`
	// The key of the error code-related documentation.
	//
	// example:
	//
	// DTS-RETRY-ERR-0046
	MigrationErrHelpDocKey *string `json:"MigrationErrHelpDocKey,omitempty" xml:"MigrationErrHelpDocKey,omitempty"`
	// The error message.
	//
	// example:
	//
	// dts.retry.err.0046.msg
	MigrationErrMsg *string `json:"MigrationErrMsg,omitempty" xml:"MigrationErrMsg,omitempty"`
	// The type of the error code.
	//
	// example:
	//
	// ForeignKey
	MigrationErrType *string `json:"MigrationErrType,omitempty" xml:"MigrationErrType,omitempty"`
	// The solution to the error.
	//
	// example:
	//
	// dts.retry.err.0046.workaround
	MigrationErrWorkaround *string `json:"MigrationErrWorkaround,omitempty" xml:"MigrationErrWorkaround,omitempty"`
	// The migration or synchronization modes.
	MigrationMode *DescribeDtsJobsResponseBodyDtsJobListMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// Lower limit of DU.
	//
	// > Only supported by Serverless instances.
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
	// The billing method of the DTS instance. Valid values:
	//
	// 	- **PrePaid**: subscription
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The performance of the data migration or synchronization instance.
	Performance *DescribeDtsJobsResponseBodyDtsJobListPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck state.
	PrecheckStatus *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet specific requirements, for example, whether to automatically start a precheck. For more information, see [MigrationReserved](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {\\"srcHostPorts\\":\\"\\",\\"whitelist.dms.online.ddl.enable\\":false,\\"filterDDL\\":false,\\"sqlparser.dms.original.ddl\\":true,\\"srcOracleType\\":\\"sid\\",\\"maxRetryTime\\":43200,\\"destSSL\\":\\"0\\",\\"destOracleType\\":\\"sid\\",\\"srcSSL\\":\\"0\\",\\"dbListCaseChangeMode\\":\\"default\\",\\"SourceEngineVersion\\":\\"8.0.18\\",\\"srcNetType\\":\\"VPC\\",\\"destNetType\\":\\"VPC\\",\\"srcVpcNetMappingInst\\":\\"172.16.1**.**:10803\\",\\"destVpcNetMappingInst\\":\\"172.16.1**.**:11077\\",\\"useJobTask\\":\\"1\\"}
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The name of the resource group.
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
	RetryState *DescribeDtsJobsResponseBodyDtsJobListRetryState `json:"RetryState,omitempty" xml:"RetryState,omitempty" type:"Struct"`
	// The details of the data synchronization task in the reverse direction.
	//
	// > This parameter is returned only for two-way data synchronization tasks.
	ReverseJob *DescribeDtsJobsResponseBodyDtsJobListReverseJob `json:"ReverseJob,omitempty" xml:"ReverseJob,omitempty" type:"Struct"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The state of the DTS instance. For more information about the valid values, see the description of the request parameter **Status**.
	//
	// example:
	//
	// Migrating
	Status                   *string                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	StructureDataCheckStatus *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus `json:"StructureDataCheckStatus,omitempty" xml:"StructureDataCheckStatus,omitempty" type:"Struct"`
	// The state of schema migration or initial schema synchronization.
	StructureInitializationStatus *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	// The tags of the task.
	TagList []*DescribeDtsJobsResponseBodyDtsJobListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
}

func (s DescribeDtsJobsResponseBodyDtsJobList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobList) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetAppName() *string {
	return s.AppName
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetBeginTimestamp() *string {
	return s.BeginTimestamp
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetConsumptionCheckpoint() *string {
	return s.ConsumptionCheckpoint
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetConsumptionClient() *string {
	return s.ConsumptionClient
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetCpuUsage() *string {
	return s.CpuUsage
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDataCloudStatus() *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus {
	return s.DataCloudStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDataEtlStatus() *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus {
	return s.DataEtlStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDataInitializationStatus() *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDataSynchronizationStatus() *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDbObject() *string {
	return s.DbObject
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDestinationEndpoint() *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDtsBisLabel() *string {
	return s.DtsBisLabel
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDtsJobDirection() *string {
	return s.DtsJobDirection
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDuRealUsage() *string {
	return s.DuRealUsage
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetDuUsage() *int64 {
	return s.DuUsage
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetErrorDetails() []*DescribeDtsJobsResponseBodyDtsJobListErrorDetails {
	return s.ErrorDetails
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetEtlSafeCheckpoint() *string {
	return s.EtlSafeCheckpoint
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetFullDataCheckStatus() *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus {
	return s.FullDataCheckStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetIncDataCheckStatus() *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus {
	return s.IncDataCheckStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetJobType() *string {
	return s.JobType
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetMaxDu() *float64 {
	return s.MaxDu
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetMemUsage() *string {
	return s.MemUsage
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetMigrationErrCode() *string {
	return s.MigrationErrCode
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetMigrationErrHelpDocId() *string {
	return s.MigrationErrHelpDocId
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetMigrationErrHelpDocKey() *string {
	return s.MigrationErrHelpDocKey
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetMigrationErrMsg() *string {
	return s.MigrationErrMsg
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetMigrationErrType() *string {
	return s.MigrationErrType
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetMigrationErrWorkaround() *string {
	return s.MigrationErrWorkaround
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetMigrationMode() *DescribeDtsJobsResponseBodyDtsJobListMigrationMode {
	return s.MigrationMode
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetMinDu() *float64 {
	return s.MinDu
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetOriginType() *string {
	return s.OriginType
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetPerformance() *DescribeDtsJobsResponseBodyDtsJobListPerformance {
	return s.Performance
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetPrecheckStatus() *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetReserved() *string {
	return s.Reserved
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetResourceGroupDisplayName() *string {
	return s.ResourceGroupDisplayName
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetRetryState() *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	return s.RetryState
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetReverseJob() *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	return s.ReverseJob
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetSourceEndpoint() *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetStructureDataCheckStatus() *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus {
	return s.StructureDataCheckStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetStructureInitializationStatus() *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) GetTagList() []*DescribeDtsJobsResponseBodyDtsJobListTagList {
	return s.TagList
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetAppName(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.AppName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetBeginTimestamp(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetCheckpoint(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.Checkpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetConsumptionCheckpoint(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetConsumptionClient(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetCpuUsage(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.CpuUsage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetCreateTime(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDataCloudStatus(v *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DataCloudStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDataEtlStatus(v *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DataEtlStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDataInitializationStatus(v *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDataSynchronizationStatus(v *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDbObject(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DbObject = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDedicatedClusterId(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDelay(v int64) *DescribeDtsJobsResponseBodyDtsJobList {
	s.Delay = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDestinationEndpoint(v *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDtsBisLabel(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DtsBisLabel = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDtsInstanceID(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDtsJobClass(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDtsJobDirection(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDtsJobId(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDtsJobName(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDuRealUsage(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DuRealUsage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDuUsage(v int64) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DuUsage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetEndTimestamp(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetErrorDetails(v []*DescribeDtsJobsResponseBodyDtsJobListErrorDetails) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ErrorDetails = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetEtlSafeCheckpoint(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.EtlSafeCheckpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetExpireTime(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetFullDataCheckStatus(v *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.FullDataCheckStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetIncDataCheckStatus(v *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.IncDataCheckStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetJobType(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.JobType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetMaxDu(v float64) *DescribeDtsJobsResponseBodyDtsJobList {
	s.MaxDu = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetMemUsage(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.MemUsage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetMigrationErrCode(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.MigrationErrCode = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetMigrationErrHelpDocId(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.MigrationErrHelpDocId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetMigrationErrHelpDocKey(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.MigrationErrHelpDocKey = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetMigrationErrMsg(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.MigrationErrMsg = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetMigrationErrType(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.MigrationErrType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetMigrationErrWorkaround(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.MigrationErrWorkaround = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetMigrationMode(v *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) *DescribeDtsJobsResponseBodyDtsJobList {
	s.MigrationMode = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetMinDu(v float64) *DescribeDtsJobsResponseBodyDtsJobList {
	s.MinDu = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetOriginType(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.OriginType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetPayType(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.PayType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetPerformance(v *DescribeDtsJobsResponseBodyDtsJobListPerformance) *DescribeDtsJobsResponseBodyDtsJobList {
	s.Performance = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetPrecheckStatus(v *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetReserved(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.Reserved = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetResourceGroupDisplayName(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ResourceGroupDisplayName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetResourceGroupId(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetRetryState(v *DescribeDtsJobsResponseBodyDtsJobListRetryState) *DescribeDtsJobsResponseBodyDtsJobList {
	s.RetryState = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetReverseJob(v *DescribeDtsJobsResponseBodyDtsJobListReverseJob) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ReverseJob = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetSourceEndpoint(v *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) *DescribeDtsJobsResponseBodyDtsJobList {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetStructureDataCheckStatus(v *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.StructureDataCheckStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetStructureInitializationStatus(v *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetTagList(v []*DescribeDtsJobsResponseBodyDtsJobListTagList) *DescribeDtsJobsResponseBodyDtsJobList {
	s.TagList = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus struct {
	// The error message returned if the task failed.
	//
	// example:
	//
	// framework: DTS-31009: In process of processing data ****
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance needs to be upgraded. Valid values:
	//
	// - **true*	-
	//
	// - **false**
	//
	// example:
	//
	// false
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of the task. Unit: percentage.
	//
	// example:
	//
	// 85
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables that have been migrated.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of the task. For more information about the valid values, see the description of the request parameter **Status**.
	//
	// example:
	//
	// Migrating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) SetNeedUpgrade(v bool) *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataCloudStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus struct {
	// The error message returned if the task failed.
	//
	// example:
	//
	// The task has failed for a long time and cannot be recovered.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of the ETL task.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of records that have been processed by the ETL task.
	//
	// example:
	//
	// 0/0
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
	// 	- **Catched**: The task is not delayed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus struct {
	// The error message returned if full data synchronization failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of full data synchronization. This is expressed as a percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of records that have been synchronized during full data synchronization.
	//
	// example:
	//
	// 44755
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of full data synchronization. Valid values:
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

func (s DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus struct {
	// The error message returned if incremental data migration or synchronization failed.
	//
	// example:
	//
	// The task has failed for a long time and cannot be recovered.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance needs to be upgraded. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// >  To upgrade a DTS instance, call the [TransferInstanceClass](https://help.aliyun.com/document_detail/281093.html) operation.
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
	// The number of records that have been migrated or synchronized during incremental data migration or synchronization.
	//
	// example:
	//
	// 0/0
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of incremental data migration or synchronization. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// 	- **Catched**: The task is not delayed.
	//
	// example:
	//
	// Catched
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint struct {
	// The name of the database to which the migration object in the destination instance belongs.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database type of the destination instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The ID of the destination instance.
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
	// 172.16.88.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// >  This parameter is returned only if the **EngineName*	- parameter of the destination instance is set to **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
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
	// The ID of the region in which the destination instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **DISABLE**: SSL encryption is disabled.
	//
	// 	- **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled, and the CA certificate is uploaded.
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

func (s DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetDatabaseName(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetEngineName(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetInstanceID(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetInstanceType(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetIp(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetOracleSID(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetPort(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetRegion(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetUserName(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListErrorDetails struct {
	// The error code returned.
	//
	// example:
	//
	// DTS-31009
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The URL of the documentation.
	//
	// example:
	//
	// https://**.ali**.com/**
	HelpUrl *string `json:"HelpUrl,omitempty" xml:"HelpUrl,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListErrorDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListErrorDetails) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListErrorDetails) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDtsJobsResponseBodyDtsJobListErrorDetails) GetHelpUrl() *string {
	return s.HelpUrl
}

func (s *DescribeDtsJobsResponseBodyDtsJobListErrorDetails) SetErrorCode(v string) *DescribeDtsJobsResponseBodyDtsJobListErrorDetails {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListErrorDetails) SetHelpUrl(v string) *DescribeDtsJobsResponseBodyDtsJobListErrorDetails {
	s.HelpUrl = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListErrorDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus struct {
	CanSwitch *bool `json:"CanSwitch,omitempty" xml:"CanSwitch,omitempty"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of the full data verification task. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The progress of the full data verification task.
	//
	// example:
	//
	// 1 rows/s (row: 5/5, table: 1/1)
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of the full data verification task. Valid values:
	//
	// - **NotStarted**: The verification is not started.
	//
	// - **Checking**: The verification is in progress.
	//
	// - **Failed**: The verification failed.
	//
	// - **Finished**: The verification is complete.
	//
	// example:
	//
	// Checking
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) GetCanSwitch() *bool {
	return s.CanSwitch
}

func (s *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) SetCanSwitch(v bool) *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus {
	s.CanSwitch = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListFullDataCheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus struct {
	// The error message returned if the task failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of the incremental data verification task. Unit: percentage.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The progress of the incremental data verification task.
	//
	// example:
	//
	// 1 rows/s (row: 5/5, table: 1/1)
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of the incremental data verification task. Valid values:
	//
	// - **Catched**: The verification is delayed.
	//
	// - **NotStarted**: The verification is not started.
	//
	// - **Checking**: The verification is in progress.
	//
	// - **Failed**: The verification failed.
	//
	// example:
	//
	// Checking
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListIncDataCheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListMigrationMode struct {
	// Indicates whether full data migration or synchronization is performed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Indicates whether incremental data migration or synchronization is performed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Indicates whether full data verification is performed. Valid values:
	//
	// -  **true**: yes
	//
	// -  **false**: no
	//
	// example:
	//
	// true
	FullDataCheck *bool `json:"FullDataCheck,omitempty" xml:"FullDataCheck,omitempty"`
	// Indicates whether incremental data verification is performed. Valid values:
	//
	// -  **true**: yes
	//
	// -  **false**: no
	//
	// example:
	//
	// true
	IncDataCheck       *bool `json:"IncDataCheck,omitempty" xml:"IncDataCheck,omitempty"`
	StructureDataCheck *bool `json:"StructureDataCheck,omitempty" xml:"StructureDataCheck,omitempty"`
	// Indicates whether schema migration or schema synchronization is performed. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) GetFullDataCheck() *bool {
	return s.FullDataCheck
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) GetIncDataCheck() *bool {
	return s.IncDataCheck
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) GetStructureDataCheck() *bool {
	return s.StructureDataCheck
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) SetDataInitialization(v bool) *DescribeDtsJobsResponseBodyDtsJobListMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) SetDataSynchronization(v bool) *DescribeDtsJobsResponseBodyDtsJobListMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) SetFullDataCheck(v bool) *DescribeDtsJobsResponseBodyDtsJobListMigrationMode {
	s.FullDataCheck = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) SetIncDataCheck(v bool) *DescribeDtsJobsResponseBodyDtsJobListMigrationMode {
	s.IncDataCheck = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) SetStructureDataCheck(v bool) *DescribeDtsJobsResponseBodyDtsJobListMigrationMode {
	s.StructureDataCheck = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) SetStructureInitialization(v bool) *DescribeDtsJobsResponseBodyDtsJobListMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListPerformance struct {
	// The size of data that is migrated or synchronized per second. Unit: MB/s.
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

func (s DescribeDtsJobsResponseBodyDtsJobListPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListPerformance) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPerformance) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPerformance) GetRps() *string {
	return s.Rps
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPerformance) SetFlow(v string) *DescribeDtsJobsResponseBodyDtsJobListPerformance {
	s.Flow = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPerformance) SetRps(v string) *DescribeDtsJobsResponseBodyDtsJobListPerformance {
	s.Rps = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPerformance) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus struct {
	// The result of each precheck item.
	Detail []*DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The cause of the precheck failure.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The precheck progress. This is expressed as a percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck status. Valid values:
	//
	// 	- **NotStarted**
	//
	// 	- **Suspending**:
	//
	// 	- **Checking**
	//
	// 	- **Failed**
	//
	// 	- **Finished**
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) GetDetail() []*DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) SetDetail(v []*DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail struct {
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
	// >  This parameter is returned only if the value of the **CheckResult*	- parameter is **Failed**.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The method to fix the precheck failure.
	//
	// >  This parameter is returned only if the value of the **CheckResult*	- parameter is **Failed**.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) GetCheckItemDescription() *string {
	return s.CheckItemDescription
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) GetCheckResult() *string {
	return s.CheckResult
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) GetFailedReason() *string {
	return s.FailedReason
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) SetCheckItem(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) SetCheckItemDescription(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) SetCheckResult(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) SetFailedReason(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail {
	s.FailedReason = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) SetRepairMethod(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListRetryState struct {
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
	// The error code.
	//
	// example:
	//
	// dts.retry.err.0046
	MigrationErrCode *string `json:"MigrationErrCode,omitempty" xml:"MigrationErrCode,omitempty"`
	// The ID of the error code-related documentation.
	//
	// example:
	//
	// 462133
	MigrationErrHelpDocId *string `json:"MigrationErrHelpDocId,omitempty" xml:"MigrationErrHelpDocId,omitempty"`
	// The key of the error code-related documentation.
	//
	// example:
	//
	// DTS-RETRY-ERR-0046
	MigrationErrHelpDocKey *string `json:"MigrationErrHelpDocKey,omitempty" xml:"MigrationErrHelpDocKey,omitempty"`
	// The error message.
	//
	// example:
	//
	// dts.retry.err.0046.msg
	MigrationErrMsg *string `json:"MigrationErrMsg,omitempty" xml:"MigrationErrMsg,omitempty"`
	// The type of the error code.
	//
	// example:
	//
	// ForeignKey
	MigrationErrType *string `json:"MigrationErrType,omitempty" xml:"MigrationErrType,omitempty"`
	// The solution to the error.
	//
	// example:
	//
	// dts.retry.err.0046.workaround
	MigrationErrWorkaround *string `json:"MigrationErrWorkaround,omitempty" xml:"MigrationErrWorkaround,omitempty"`
	// The progress of the instance when DTS retries.
	//
	// example:
	//
	// 03
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The number of retries that have been performed.
	//
	// example:
	//
	// 5
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// The object on which these retries are performed. Valid values:
	//
	// - **srcDB**: the source database
	//
	// - **destDB**: the destination database
	//
	// - **inner_module**: an internal module of DTS
	//
	// example:
	//
	// srcDB
	RetryTarget *string `json:"RetryTarget,omitempty" xml:"RetryTarget,omitempty"`
	// The time that has elapsed from the time when the first retry starts. Unit: seconds.
	//
	// example:
	//
	// 3600
	RetryTime *int32 `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	// Indicates whether the task is being retried. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// false
	Retrying *bool `json:"Retrying,omitempty" xml:"Retrying,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListRetryState) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListRetryState) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetMaxRetryTime() *int32 {
	return s.MaxRetryTime
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetMigrationErrCode() *string {
	return s.MigrationErrCode
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetMigrationErrHelpDocId() *string {
	return s.MigrationErrHelpDocId
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetMigrationErrHelpDocKey() *string {
	return s.MigrationErrHelpDocKey
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetMigrationErrMsg() *string {
	return s.MigrationErrMsg
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetMigrationErrType() *string {
	return s.MigrationErrType
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetMigrationErrWorkaround() *string {
	return s.MigrationErrWorkaround
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetModule() *string {
	return s.Module
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetRetryTarget() *string {
	return s.RetryTarget
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetRetryTime() *int32 {
	return s.RetryTime
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) GetRetrying() *bool {
	return s.Retrying
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetErrMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetJobId(v string) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.JobId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetMaxRetryTime(v int32) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.MaxRetryTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetMigrationErrCode(v string) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.MigrationErrCode = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetMigrationErrHelpDocId(v string) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.MigrationErrHelpDocId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetMigrationErrHelpDocKey(v string) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.MigrationErrHelpDocKey = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetMigrationErrMsg(v string) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.MigrationErrMsg = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetMigrationErrType(v string) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.MigrationErrType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetMigrationErrWorkaround(v string) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.MigrationErrWorkaround = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetModule(v string) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.Module = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetRetryCount(v int32) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.RetryCount = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetRetryTarget(v string) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.RetryTarget = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetRetryTime(v int32) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.RetryTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetRetrying(v bool) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.Retrying = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJob struct {
	// The start offset of incremental data synchronization. This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1616980369
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The CPU utilization of the instance. Unit: percentage.
	//
	// example:
	//
	// 90
	CpuUsage *string `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The state of initial full data synchronization.
	DataInitializationStatus *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The state of incremental data synchronization.
	DataSynchronizationStatus *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The schema of the objects that you want to synchronize. The value is a JSON string and can contain regular expressions. For more information, see Objects of DTS tasks.
	//
	// example:
	//
	// {"dtstestdata": { "name": "dtstestdata", "all": true }}
	DbObject *string `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	// The ID of the DTS dedicated cluster on which a DTS task runs.
	//
	// example:
	//
	// dtscluster_dpwl3**********
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The latency of incremental data synchronization. Unit: seconds.
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The connection settings of the destination instance.
	DestinationEndpoint *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The ID of the data synchronization instance.
	//
	// example:
	//
	// dtsi03e3zty16i****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The instance class.
	//
	// > For more information about the test performance of each instance class, see [Specifications of data synchronization instances](https://help.aliyun.com/document_detail/26605.html).
	//
	// example:
	//
	// large
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The synchronization direction. **Reverse*	- is returned.
	//
	// example:
	//
	// Reverse
	DtsJobDirection *string `json:"DtsJobDirection,omitempty" xml:"DtsJobDirection,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// i03e3zty16i****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The name of the data synchronization task.
	//
	// example:
	//
	// RDS_TO_RDS_MIGRATION
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The number of DUs that have been used.
	//
	// example:
	//
	// 15
	DuUsage *int64 `json:"DuUsage,omitempty" xml:"DuUsage,omitempty"`
	// The error message returned.
	ErrorDetails []*DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails `json:"ErrorDetails,omitempty" xml:"ErrorDetails,omitempty" type:"Repeated"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by [com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException:Could not create connection to database server. Attempted reconnect 3 times. Giving up.][com.mysql.jdbc.exceptions.jdbc4.CommunicationsException:Communications link failure\\n\\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.][java.net.ConnectException:Connection timed out (Connection timed out)] About more information in [https://yq.aliyun.com/articles/499178].
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The checkpoint of the ETL task.
	//
	// example:
	//
	// 1610540493
	EtlSafeCheckpoint *string `json:"EtlSafeCheckpoint,omitempty" xml:"EtlSafeCheckpoint,omitempty"`
	// The time when the instance expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// > This parameter is returned only if the returned value of **PayType*	- is **PrePaid**.
	//
	// example:
	//
	// 2023-03-16T08:01:19Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The state information about the full data verification task.
	FullDataCheckStatus *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus `json:"FullDataCheckStatus,omitempty" xml:"FullDataCheckStatus,omitempty" type:"Struct"`
	// The state information about the incremental data verification task.
	IncDataCheckStatus *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus `json:"IncDataCheckStatus,omitempty" xml:"IncDataCheckStatus,omitempty" type:"Struct"`
	// Upper limit of DU.
	//
	// > Only supported by Serverless instances.
	//
	// example:
	//
	// 16
	MaxDu *float64 `json:"MaxDu,omitempty" xml:"MaxDu,omitempty"`
	// The memory that has been used. Unit: MB.
	//
	// example:
	//
	// 500
	MemUsage *string `json:"MemUsage,omitempty" xml:"MemUsage,omitempty"`
	// The initial synchronization types.
	MigrationMode *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// Lower limit of DU.
	//
	// > Only supported by Serverless instances.
	//
	// example:
	//
	// 1
	MinDu *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	// The billing method of the DTS instance. Valid values:
	//
	// - **PrePaid**: subscription
	//
	// - **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The performance of the data synchronization instance.
	Performance *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck state.
	PrecheckStatus *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet specific requirements, for example, whether to automatically start a precheck. For more information, see [MigrationReserved](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {\\"srcHostPorts\\":\\"\\",\\"whitelist.dms.online.ddl.enable\\":false,\\"filterDDL\\":false,\\"sqlparser.dms.original.ddl\\":true,\\"srcOracleType\\":\\"sid\\",\\"maxRetryTime\\":43200,\\"destSSL\\":\\"0\\",\\"destOracleType\\":\\"sid\\",\\"srcSSL\\":\\"0\\",\\"dbListCaseChangeMode\\":\\"default\\",\\"SourceEngineVersion\\":\\"8.0.18\\",\\"srcNetType\\":\\"VPC\\",\\"destNetType\\":\\"VPC\\",\\"srcVpcNetMappingInst\\":\\"172.16.1**.**:10803\\",\\"destVpcNetMappingInst\\":\\"172.16.1**.**:11077\\",\\"useJobTask\\":\\"1\\"}
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The state of the DTS instance. For more information about the valid values, see the description of the request parameter **Status**.
	//
	// example:
	//
	// Synchronizing
	Status                   *string                                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	StructureDataCheckStatus *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus `json:"StructureDataCheckStatus,omitempty" xml:"StructureDataCheckStatus,omitempty" type:"Struct"`
	// The state of initial schema synchronization.
	StructureInitializationStatus *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJob) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetCpuUsage() *string {
	return s.CpuUsage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetDataInitializationStatus() *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetDataSynchronizationStatus() *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetDbObject() *string {
	return s.DbObject
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetDestinationEndpoint() *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetDtsJobDirection() *string {
	return s.DtsJobDirection
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetDuUsage() *int64 {
	return s.DuUsage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetErrorDetails() []*DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails {
	return s.ErrorDetails
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetEtlSafeCheckpoint() *string {
	return s.EtlSafeCheckpoint
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetFullDataCheckStatus() *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus {
	return s.FullDataCheckStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetIncDataCheckStatus() *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus {
	return s.IncDataCheckStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetMaxDu() *float64 {
	return s.MaxDu
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetMemUsage() *string {
	return s.MemUsage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetMigrationMode() *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode {
	return s.MigrationMode
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetMinDu() *float64 {
	return s.MinDu
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetPerformance() *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance {
	return s.Performance
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetPrecheckStatus() *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetReserved() *string {
	return s.Reserved
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetSourceEndpoint() *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetStructureDataCheckStatus() *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus {
	return s.StructureDataCheckStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) GetStructureInitializationStatus() *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetCheckpoint(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.Checkpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetCpuUsage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.CpuUsage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetCreateTime(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDataInitializationStatus(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDataSynchronizationStatus(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDbObject(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DbObject = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDedicatedClusterId(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDelay(v int64) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.Delay = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDestinationEndpoint(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDtsInstanceID(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDtsJobClass(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDtsJobDirection(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDtsJobId(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDtsJobName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDuUsage(v int64) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DuUsage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetErrorDetails(v []*DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.ErrorDetails = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetEtlSafeCheckpoint(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.EtlSafeCheckpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetExpireTime(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetFullDataCheckStatus(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.FullDataCheckStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetIncDataCheckStatus(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.IncDataCheckStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetMaxDu(v float64) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.MaxDu = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetMemUsage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.MemUsage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetMigrationMode(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.MigrationMode = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetMinDu(v float64) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.MinDu = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetPayType(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.PayType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetPerformance(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.Performance = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetPrecheckStatus(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetReserved(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.Reserved = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetSourceEndpoint(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetStructureDataCheckStatus(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.StructureDataCheckStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetStructureInitializationStatus(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus struct {
	// The error message returned if initial full data synchronization failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of initial full data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that have been synchronized during initial full data synchronization.
	//
	// example:
	//
	// 43071
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of initial full data synchronization. Valid values:
	//
	// - **NotStarted**: The task is not started.
	//
	// - **Migrating**: The task is in progress.
	//
	// - **Failed**: The task failed.
	//
	// - **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus struct {
	// The error message returned if incremental data synchronization failed.
	//
	// example:
	//
	// The task has failed for a long time and cannot be recovered.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance needs to be upgraded. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// > To upgrade a DTS instance, call the [TransferInstanceClass](https://help.aliyun.com/document_detail/281093.html) operation.
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of incremental data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that have been migrated or synchronized during incremental data migration or synchronization.
	//
	// example:
	//
	// 20001
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of incremental data synchronization.
	//
	// example:
	//
	// Migrating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint struct {
	// The name of the database that contains the synchronized objects in the destination instance.
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
	// The ID of the destination instance.
	//
	// example:
	//
	// rm-bp1162kryivb8****
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
	// 172.16.88.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the returned value of **EngineName*	- of the destination instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The port number of the destination instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the region in which the destination instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// - **DISABLE**: SSL encryption is disabled.
	//
	// - **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// - **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection with an AWS MongoDB Altas database.
	//
	// - **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection with a Kafka cluster.
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

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetDatabaseName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetEngineName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetInstanceID(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetInstanceType(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetIp(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetOracleSID(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetPort(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetRegion(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetUserName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails struct {
	// The error code returned.
	//
	// example:
	//
	// DTS-31009
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The URL of the documentation.
	//
	// example:
	//
	// https://**.ali**.com/**
	HelpUrl *string `json:"HelpUrl,omitempty" xml:"HelpUrl,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails) GetHelpUrl() *string {
	return s.HelpUrl
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails) SetErrorCode(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails) SetHelpUrl(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails {
	s.HelpUrl = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobErrorDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus struct {
	CanSwitch *bool `json:"CanSwitch,omitempty" xml:"CanSwitch,omitempty"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of the full data verification task. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The progress of the full data verification task.
	//
	// example:
	//
	// 1 rows/s (row: 5/5, table: 1/1)
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of the full data verification task. Valid values:
	//
	// - **NotStarted**: The verification is not started.
	//
	// - **Checking**: The verification is in progress.
	//
	// - **Failed**: The verification failed.
	//
	// - **Finished**: The verification is complete.
	//
	// example:
	//
	// Checking
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) GetCanSwitch() *bool {
	return s.CanSwitch
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) SetCanSwitch(v bool) *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus {
	s.CanSwitch = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobFullDataCheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus struct {
	// The error message returned if the task failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of the incremental data verification task. Unit: percentage.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The progress of the incremental data verification task.
	//
	// example:
	//
	// 1 rows/s (row: 5/5, table: 1/1)
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of the incremental data verification task. Valid values:
	//
	// - **Catched**: The verification is delayed.
	//
	// - **NotStarted**: The verification is not started.
	//
	// - **Checking**: The verification is in progress.
	//
	// - **Failed**: The verification failed.
	//
	// example:
	//
	// Checking
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobIncDataCheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode struct {
	// Indicates whether initial full data synchronization is performed. Valid values:
	//
	// -  **true**
	//
	// -  **false**
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Indicates whether incremental data synchronization is performed. Valid values:
	//
	// -  **true**
	//
	// -  **false**
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Indicates whether full data verification is performed. Valid values:
	//
	// -  **true**: yes
	//
	// -  **false**: no
	//
	// example:
	//
	// true
	FullDataCheck *bool `json:"FullDataCheck,omitempty" xml:"FullDataCheck,omitempty"`
	// Indicates whether incremental data verification is performed. Valid values:
	//
	// -  **true**: yes
	//
	// -  **false**: no
	//
	// example:
	//
	// true
	IncDataCheck       *bool `json:"IncDataCheck,omitempty" xml:"IncDataCheck,omitempty"`
	StructureDataCheck *bool `json:"StructureDataCheck,omitempty" xml:"StructureDataCheck,omitempty"`
	// Indicates whether initial schema synchronization is performed. Valid values:
	//
	// -  **true**
	//
	// -  **false**
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) GetFullDataCheck() *bool {
	return s.FullDataCheck
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) GetIncDataCheck() *bool {
	return s.IncDataCheck
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) GetStructureDataCheck() *bool {
	return s.StructureDataCheck
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) SetDataInitialization(v bool) *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) SetDataSynchronization(v bool) *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) SetFullDataCheck(v bool) *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode {
	s.FullDataCheck = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) SetIncDataCheck(v bool) *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode {
	s.IncDataCheck = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) SetStructureDataCheck(v bool) *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode {
	s.StructureDataCheck = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) SetStructureInitialization(v bool) *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance struct {
	// The size of data that is synchronized per second. Unit: MB/s.
	//
	// example:
	//
	// 1
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The number of times that SQL statements are synchronized per second, including BEGIN, COMMIT, DML, and DDL statements. DML statements include INSERT, DELETE, and UPDATE.
	//
	// example:
	//
	// 100
	Rps *string `json:"Rps,omitempty" xml:"Rps,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) GetRps() *string {
	return s.Rps
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) SetFlow(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance {
	s.Flow = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) SetRps(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance {
	s.Rps = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus struct {
	// The result of each precheck item.
	Detail []*DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The error message returned if the precheck failed.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The precheck progress. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck state. Valid values:
	//
	// - **NotStarted**: The precheck is not started.
	//
	// - **Suspending**: The precheck is paused.
	//
	// - **Checking**: The precheck is in progress.
	//
	// - **Failed**: The precheck failed.
	//
	// - **Finished**: The precheck is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) GetDetail() []*DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) SetDetail(v []*DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail struct {
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
	// - **Success**
	//
	// - **Failed**
	//
	// example:
	//
	// Success
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The error message returned if the task failed to pass the precheck.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The method to fix a precheck failure.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GetCheckItemDescription() *string {
	return s.CheckItemDescription
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GetCheckResult() *string {
	return s.CheckResult
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GetFailedReason() *string {
	return s.FailedReason
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetCheckItem(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetCheckItemDescription(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetCheckResult(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetFailedReason(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.FailedReason = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetRepairMethod(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint struct {
	// The name of the database that contains the objects to be migrated from the source instance.
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
	// The ID of the source instance.
	//
	// example:
	//
	// rm-bp1imrtn6fq7h****
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
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the returned value of **EngineName*	- of the source instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The port number of the source instance.
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
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// - **DISABLE**: SSL encryption is disabled.
	//
	// - **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// - **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection with an AWS MongoDB Altas database.
	//
	// - **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection with a Kafka cluster.
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

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetDatabaseName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetEngineName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetInstanceID(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetInstanceType(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetIp(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetOracleSID(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetPort(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetRegion(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetUserName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureDataCheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus struct {
	// The error message returned if initial schema synchronization failed.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: ERROR: type "geometry" does not exist;
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of initial schema synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables that have been synchronized during initial schema synchronization.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of initial schema synchronization. Valid values:
	//
	// - **NotStarted**: The task is not started.
	//
	// - **Migrating**: The task is in progress.
	//
	// - **Failed**: The task failed.
	//
	// - **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint struct {
	// The name of the database that contains the objects to be migrated from the source instance.
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
	// The ID of the source instance.
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
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the returned value of **EngineName*	- of the source instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The port number of the source instance.
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
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// - **DISABLE**: SSL encryption is disabled.
	//
	// - **ENABLE_WITH_CERTIFICAT**E: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// - **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection with an AWS MongoDB Altas database.
	//
	// - **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection with a Kafka cluster.
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

func (s DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetDatabaseName(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetEngineName(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetInstanceID(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetInstanceType(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetIp(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetOracleSID(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetPort(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetRegion(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetUserName(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureDataCheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus struct {
	// The error message returned if schema migration or initial schema synchronization failed.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: ERROR: type "geometry" does not exist;
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of schema migration or initial schema synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables that have been migrated or synchronized during schema migration or initial schema synchronization.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of schema migration or initial schema synchronization. Valid values:
	//
	// - **NotStarted**: The task is not started.
	//
	// - **Migrating**: The task is in progress.
	//
	// - **Failed**: The task failed.
	//
	// - **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyDtsJobListTagList struct {
	// The key of the tag.
	//
	// example:
	//
	// testkey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// testvalue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListTagList) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDtsJobsResponseBodyDtsJobListTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDtsJobsResponseBodyDtsJobListTagList) SetTagKey(v string) *DescribeDtsJobsResponseBodyDtsJobListTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListTagList) SetTagValue(v string) *DescribeDtsJobsResponseBodyDtsJobListTagList {
	s.TagValue = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListTagList) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoList struct {
	// Indicates whether the **new*	- change tracking feature is used.
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
	// 2021-03-15T08:25:34Z
	BeginTimestamp *string `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	// The start offset of incremental data migration or data synchronization. This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1616899019
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The consumption checkpoint of the change tracking instance. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-23T07:30:31Z
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	// The downstream client information in the following format: <IP address of the downstream client>:<Random ID generated by DTS>.
	//
	// example:
	//
	// 114...:dts******
	ConsumptionClient *string `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:s*sZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-16T08:01:19Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The state of the ETL task.
	//
	// > This parameter collection is returned only if an ETL task is configured.
	DataEtlStatus *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus `json:"DataEtlStatus,omitempty" xml:"DataEtlStatus,omitempty" type:"Struct"`
	// The state of full data migration or initial full data synchronization.
	DataInitializationStatus *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The state of incremental data migration or synchronization.
	DataSynchronizationStatus *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The objects of the data migration, data synchronization, or change tracking task. For more information, see [Objects of DTS tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// example:
	//
	// {"dtstestdata": { "name": "dtstestdata", "all": true }}
	DbObject *string `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	// The latency of incremental data migration or synchronization.
	//
	// > If you query data migration tasks, the unit of this parameter is milliseconds. If you query data synchronization tasks, the unit of this parameter is seconds.
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The connection settings of the destination instance.
	DestinationEndpoint *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The ID of the data migration, data synchronization, or change tracking instance.
	//
	// example:
	//
	// dtsi03e3zty16i****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The instance class.
	//
	// > For more information about the test performance of each instance class, see [Specifications of data migration instances](https://help.aliyun.com/document_detail/26606.html) and [Specifications of data synchronization instances](https://help.aliyun.com/document_detail/26605.html).
	//
	// example:
	//
	// large
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**
	//
	// - **Reverse**
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
	// RDS_TO_RDS_MIGRATION
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The end of the time range for change tracking. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-26T14:03:21Z
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by [com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException:Could not create connection to database server. Attempted reconnect 3 times. Giving up.][com.mysql.jdbc.exceptions.jdbc4.CommunicationsException:Communications link failure\\n\\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.][java.net.ConnectException:Connection timed out (Connection timed out)] About more information in [https://yq.aliyun.com/articles/499178].
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The checkpoint of the ETL task.
	//
	// example:
	//
	// 1610540493
	EtlSafeCheckpoint *string `json:"EtlSafeCheckpoint,omitempty" xml:"EtlSafeCheckpoint,omitempty"`
	// The time when the instance expires. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// > This parameter is returned only if the returned value of **PayType*	- is **PrePaid**.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The type of the DTS task. Valid values:
	//
	// - **MIGRATION**: data migration task
	//
	// - **SYNC**: data synchronization task
	//
	// - **SUBSCRIBE**: change tracking task
	//
	// example:
	//
	// MIGRATION
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The migration types or initial synchronization types.
	MigrationMode *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// The source of the task.
	//
	// - **PTS**
	//
	// - **DMS**
	//
	// - **DTS**
	//
	// example:
	//
	// DTS
	OriginType *string `json:"OriginType,omitempty" xml:"OriginType,omitempty"`
	// The billing method of the DTS instance. Valid values:
	//
	// - **PrePaid**: subscription
	//
	// - **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The performance of the data migration or synchronization instance.
	Performance *DescribeDtsJobsResponseBodyEtlDemoListPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck state.
	PrecheckStatus *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet specific requirements, for example, whether to automatically start a precheck. For more information, see [MigrationReserved](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {\\"srcHostPorts\\":\\"\\",\\"whitelist.dms.online.ddl.enable\\":false,\\"filterDDL\\":false,\\"sqlparser.dms.original.ddl\\":true,\\"srcOracleType\\":\\"sid\\",\\"maxRetryTime\\":43200,\\"destSSL\\":\\"0\\",\\"destOracleType\\":\\"sid\\",\\"srcSSL\\":\\"0\\",\\"dbListCaseChangeMode\\":\\"default\\",\\"SourceEngineVersion\\":\\"8.0.18\\",\\"srcNetType\\":\\"VPC\\",\\"destNetType\\":\\"VPC\\",\\"srcVpcNetMappingInst\\":\\"172.16.1**.**:10803\\",\\"destVpcNetMappingInst\\":\\"172.16.1**.**:11077\\",\\"useJobTask\\":\\"1\\"}
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The name of the resource group.
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
	RetryState *DescribeDtsJobsResponseBodyEtlDemoListRetryState `json:"RetryState,omitempty" xml:"RetryState,omitempty" type:"Struct"`
	// The details of the data synchronization task in the reverse direction.
	//
	// > This parameter is returned only for two-way data synchronization tasks.
	ReverseJob *DescribeDtsJobsResponseBodyEtlDemoListReverseJob `json:"ReverseJob,omitempty" xml:"ReverseJob,omitempty" type:"Struct"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The state of the DTS instance. For more information about the valid values, see the description of the request parameter **Status**.
	//
	// example:
	//
	// Migrating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The state of schema migration or initial schema synchronization.
	StructureInitializationStatus *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	// The tags of the task.
	TagList []*DescribeDtsJobsResponseBodyEtlDemoListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoList) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetAppName() *string {
	return s.AppName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetBeginTimestamp() *string {
	return s.BeginTimestamp
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetConsumptionCheckpoint() *string {
	return s.ConsumptionCheckpoint
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetConsumptionClient() *string {
	return s.ConsumptionClient
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetDataEtlStatus() *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus {
	return s.DataEtlStatus
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetDataInitializationStatus() *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetDataSynchronizationStatus() *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetDbObject() *string {
	return s.DbObject
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetDestinationEndpoint() *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetDtsJobDirection() *string {
	return s.DtsJobDirection
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetEtlSafeCheckpoint() *string {
	return s.EtlSafeCheckpoint
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetJobType() *string {
	return s.JobType
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetMigrationMode() *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode {
	return s.MigrationMode
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetOriginType() *string {
	return s.OriginType
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetPerformance() *DescribeDtsJobsResponseBodyEtlDemoListPerformance {
	return s.Performance
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetPrecheckStatus() *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetReserved() *string {
	return s.Reserved
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetResourceGroupDisplayName() *string {
	return s.ResourceGroupDisplayName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetRetryState() *DescribeDtsJobsResponseBodyEtlDemoListRetryState {
	return s.RetryState
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetReverseJob() *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	return s.ReverseJob
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetSourceEndpoint() *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetStructureInitializationStatus() *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) GetTagList() []*DescribeDtsJobsResponseBodyEtlDemoListTagList {
	return s.TagList
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetAppName(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.AppName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetBeginTimestamp(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetCheckpoint(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.Checkpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetConsumptionCheckpoint(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetConsumptionClient(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetCreateTime(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetDataEtlStatus(v *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.DataEtlStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetDataInitializationStatus(v *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetDataSynchronizationStatus(v *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetDbObject(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.DbObject = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetDelay(v int64) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.Delay = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetDestinationEndpoint(v *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetDtsInstanceID(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetDtsJobClass(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetDtsJobDirection(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetDtsJobId(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetDtsJobName(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetEndTimestamp(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetEtlSafeCheckpoint(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.EtlSafeCheckpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetExpireTime(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetJobType(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.JobType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetMigrationMode(v *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.MigrationMode = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetOriginType(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.OriginType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetPayType(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.PayType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetPerformance(v *DescribeDtsJobsResponseBodyEtlDemoListPerformance) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.Performance = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetPrecheckStatus(v *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetReserved(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.Reserved = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetResourceGroupDisplayName(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.ResourceGroupDisplayName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetResourceGroupId(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetRetryState(v *DescribeDtsJobsResponseBodyEtlDemoListRetryState) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.RetryState = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetReverseJob(v *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.ReverseJob = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetSourceEndpoint(v *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetStatus(v string) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetStructureInitializationStatus(v *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) SetTagList(v []*DescribeDtsJobsResponseBodyEtlDemoListTagList) *DescribeDtsJobsResponseBodyEtlDemoList {
	s.TagList = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoList) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus struct {
	// The error message returned if the ETL task failed.
	//
	// example:
	//
	// The task has failed for a long time and cannot be recovered.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of the ETL task. Unit: percentage.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that have been processed by the ETL task.
	//
	// example:
	//
	// 0/0
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of the ETL task. Valid values:
	//
	// - **NotStarted**: The task is not started.
	//
	// - **Migrating**: The task is in progress.
	//
	// - **Failed**: The task failed.
	//
	// - **Finished**: The task is complete.
	//
	// - **Catched**: The task is not delayed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataEtlStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus struct {
	// The error message returned if full data migration or initial full data synchronization failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of full data migration or initial full data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that have been migrated or synchronized during full data migration or initial full data synchronization.
	//
	// example:
	//
	// 44755
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of full data migration or initial full data synchronization. Valid values:
	//
	// - **NotStarted**: The task is not started.
	//
	// - **Migrating**: The task is in progress.
	//
	// - **Failed**: The task failed.
	//
	// - **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus struct {
	// The error message returned if incremental data migration or synchronization failed.
	//
	// example:
	//
	// The task has failed for a long time and cannot be recovered.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance needs to be upgraded. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// > To upgrade a DTS instance, call the [TransferInstanceClass](https://help.aliyun.com/document_detail/281093.html) operation.
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
	// The number of entries that have been migrated or synchronized during incremental data migration or synchronization.
	//
	// example:
	//
	// 0/0
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of incremental data migration or synchronization. Valid values:
	//
	// - **NotStarted**: The task is not started.
	//
	// - **Migrating**: The task is in progress.
	//
	// - **Failed**: The task failed.
	//
	// - **Finished**: The task is complete.
	//
	// - **Catched**: The task is not delayed.
	//
	// example:
	//
	// Catched
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint struct {
	// The name of the database that contains the migrated objects in the destination instance.
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
	// The ID of the destination instance.
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
	// 172.16.88.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the returned value of **EngineName*	- of the destination instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The port number of the destination instance.
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
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// - **DISABLE**: SSL encryption is disabled.
	//
	// - **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// - **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection with an AWS MongoDB Altas database.
	//
	// - **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection with a Kafka cluster.
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

func (s DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) SetDatabaseName(v string) *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) SetEngineName(v string) *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) SetInstanceID(v string) *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) SetInstanceType(v string) *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) SetIp(v string) *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) SetOracleSID(v string) *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) SetPort(v string) *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) SetRegion(v string) *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) SetUserName(v string) *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListMigrationMode struct {
	// Indicates whether full data migration or initial full data synchronization is performed. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Indicates whether incremental data migration or synchronization is performed. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Indicates whether schema migration or initial schema synchronization is performed. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode) SetDataInitialization(v bool) *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode) SetDataSynchronization(v bool) *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode) SetStructureInitialization(v bool) *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListPerformance struct {
	// The size of data that is migrated or synchronized per second. Unit: MB/s.
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

func (s DescribeDtsJobsResponseBodyEtlDemoListPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListPerformance) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPerformance) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPerformance) GetRps() *string {
	return s.Rps
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPerformance) SetFlow(v string) *DescribeDtsJobsResponseBodyEtlDemoListPerformance {
	s.Flow = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPerformance) SetRps(v string) *DescribeDtsJobsResponseBodyEtlDemoListPerformance {
	s.Rps = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPerformance) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus struct {
	// The result of each precheck item.
	Detail []*DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The error message returned if the precheck failed.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The precheck progress. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck state. Valid values:
	//
	// - **NotStarted**: The precheck is not started.
	//
	// - **Suspending**: The precheck is paused.
	//
	// - **Checking**: The precheck is in progress.
	//
	// - **Failed**: The precheck failed.
	//
	// - **Finished**: The precheck is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus) GetDetail() []*DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus) SetDetail(v []*DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail struct {
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
	// - **Success**
	//
	// - **Failed**
	//
	// example:
	//
	// Success
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The error message returned if the task failed to pass the precheck.
	//
	// > This parameter is returned only if the returned value of **CheckResult*	- is **Failed**.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The method to fix a precheck failure.
	//
	// > This parameter is returned only if the returned value of **CheckResult*	- is **Failed**.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) GetCheckItemDescription() *string {
	return s.CheckItemDescription
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) GetCheckResult() *string {
	return s.CheckResult
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) GetFailedReason() *string {
	return s.FailedReason
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) SetCheckItem(v string) *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) SetCheckItemDescription(v string) *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) SetCheckResult(v string) *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) SetFailedReason(v string) *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail {
	s.FailedReason = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) SetRepairMethod(v string) *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListRetryState struct {
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
	// The progress of the instance when DTS retries.
	//
	// example:
	//
	// 03
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The number of retries that have been performed.
	//
	// example:
	//
	// 5
	RetryCount *int32 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// The object on which these retries are performed. Valid values:
	//
	// - **srcDB**: the source database
	//
	// - **destDB**: the destination database
	//
	// - **inner_module**: an internal module of DTS
	//
	// example:
	//
	// srcDB
	RetryTarget *string `json:"RetryTarget,omitempty" xml:"RetryTarget,omitempty"`
	// The time that has elapsed from the time when the first retry starts. Unit: seconds.
	//
	// example:
	//
	// 3600
	RetryTime *int32 `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	// Indicates whether the task is being retried. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// false
	Retrying *bool `json:"Retrying,omitempty" xml:"Retrying,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListRetryState) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListRetryState) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) GetMaxRetryTime() *int32 {
	return s.MaxRetryTime
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) GetModule() *string {
	return s.Module
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) GetRetryTarget() *string {
	return s.RetryTarget
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) GetRetryTime() *int32 {
	return s.RetryTime
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) GetRetrying() *bool {
	return s.Retrying
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) SetErrMessage(v string) *DescribeDtsJobsResponseBodyEtlDemoListRetryState {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) SetJobId(v string) *DescribeDtsJobsResponseBodyEtlDemoListRetryState {
	s.JobId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) SetMaxRetryTime(v int32) *DescribeDtsJobsResponseBodyEtlDemoListRetryState {
	s.MaxRetryTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) SetModule(v string) *DescribeDtsJobsResponseBodyEtlDemoListRetryState {
	s.Module = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) SetRetryCount(v int32) *DescribeDtsJobsResponseBodyEtlDemoListRetryState {
	s.RetryCount = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) SetRetryTarget(v string) *DescribeDtsJobsResponseBodyEtlDemoListRetryState {
	s.RetryTarget = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) SetRetryTime(v int32) *DescribeDtsJobsResponseBodyEtlDemoListRetryState {
	s.RetryTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) SetRetrying(v bool) *DescribeDtsJobsResponseBodyEtlDemoListRetryState {
	s.Retrying = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListRetryState) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListReverseJob struct {
	// The start offset of incremental data synchronization. This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1616980369
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-16T08:01:19Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The state of initial full data synchronization.
	DataInitializationStatus *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The state of incremental data synchronization.
	DataSynchronizationStatus *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The schema of the objects that you want to synchronize. The value is a JSON string and can contain regular expressions. For more information, see Objects of DTS tasks.
	//
	// example:
	//
	// {"dtstestdata": { "name": "dtstestdata", "all": true }}
	DbObject *string `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	// The latency of incremental data synchronization. Unit: seconds.
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The connection settings of the destination instance.
	DestinationEndpoint *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The ID of the data synchronization instance.
	//
	// example:
	//
	// dtsi03e3zty16i****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The instance class.
	//
	// > For more information about the test performance of each instance class, see [Specifications of data synchronization instances](https://help.aliyun.com/document_detail/26605.html).
	//
	// example:
	//
	// large
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The synchronization direction. **Reverse*	- is returned.
	//
	// example:
	//
	// Reverse
	DtsJobDirection *string `json:"DtsJobDirection,omitempty" xml:"DtsJobDirection,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// i03e3zty16i****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The name of the data synchronization task.
	//
	// example:
	//
	// RDS_TO_RDS_MIGRATION
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The error message returned if the task failed.
	//
	// example:
	//
	// DTS-070211: Connect Source DB failed. cause by [com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException:Could not create connection to database server. Attempted reconnect 3 times. Giving up.][com.mysql.jdbc.exceptions.jdbc4.CommunicationsException:Communications link failure\\n\\nThe last packet sent successfully to the server was 0 milliseconds ago. The driver has not received any packets from the server.][java.net.ConnectException:Connection timed out (Connection timed out)] About more information in [https://yq.aliyun.com/articles/499178].
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The checkpoint of the ETL task.
	//
	// example:
	//
	// 1610540493
	EtlSafeCheckpoint *string `json:"EtlSafeCheckpoint,omitempty" xml:"EtlSafeCheckpoint,omitempty"`
	// The time when the instance expires. The time follows the ISO 8601 standard in the	- yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// > This parameter is returned only if the returned value of **PayType*	- is **PrePaid**.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The migration types or initial synchronization types.
	MigrationMode *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// The billing method of the DTS instance. Valid values:
	//
	// - **PrePaid**: subscription
	//
	// - **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The performance of the data migration or synchronization instance.
	Performance *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck state.
	PrecheckStatus *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The reserved parameter of DTS. The value is a JSON string. You can specify this parameter to meet specific requirements, for example, whether to automatically start a precheck. For more information, see [MigrationReserved](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {\\"srcHostPorts\\":\\"\\",\\"whitelist.dms.online.ddl.enable\\":false,\\"filterDDL\\":false,\\"sqlparser.dms.original.ddl\\":true,\\"srcOracleType\\":\\"sid\\",\\"maxRetryTime\\":43200,\\"destSSL\\":\\"0\\",\\"destOracleType\\":\\"sid\\",\\"srcSSL\\":\\"0\\",\\"dbListCaseChangeMode\\":\\"default\\",\\"SourceEngineVersion\\":\\"8.0.18\\",\\"srcNetType\\":\\"VPC\\",\\"destNetType\\":\\"VPC\\",\\"srcVpcNetMappingInst\\":\\"172.16.1**.**:10803\\",\\"destVpcNetMappingInst\\":\\"172.16.1**.**:11077\\",\\"useJobTask\\":\\"1\\"}
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The connection settings of the source instance.
	SourceEndpoint *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The state of the DTS instance. For more information about the valid values, see the description of the request parameter **Status**.
	//
	// example:
	//
	// Synchronizing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The state of initial schema synchronization.
	StructureInitializationStatus *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetDataInitializationStatus() *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetDataSynchronizationStatus() *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetDbObject() *string {
	return s.DbObject
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetDestinationEndpoint() *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetDtsJobDirection() *string {
	return s.DtsJobDirection
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetEtlSafeCheckpoint() *string {
	return s.EtlSafeCheckpoint
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetMigrationMode() *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode {
	return s.MigrationMode
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetPerformance() *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance {
	return s.Performance
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetPrecheckStatus() *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetReserved() *string {
	return s.Reserved
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetSourceEndpoint() *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) GetStructureInitializationStatus() *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetCheckpoint(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.Checkpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetCreateTime(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetDataInitializationStatus(v *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetDataSynchronizationStatus(v *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetDbObject(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.DbObject = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetDelay(v int64) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.Delay = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetDestinationEndpoint(v *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetDtsInstanceID(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetDtsJobClass(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetDtsJobDirection(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetDtsJobId(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetDtsJobName(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetEtlSafeCheckpoint(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.EtlSafeCheckpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetExpireTime(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetMigrationMode(v *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.MigrationMode = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetPayType(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.PayType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetPerformance(v *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.Performance = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetPrecheckStatus(v *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetReserved(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.Reserved = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetSourceEndpoint(v *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetStatus(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) SetStructureInitializationStatus(v *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus) *DescribeDtsJobsResponseBodyEtlDemoListReverseJob {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJob) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus struct {
	// The error message returned if initial full data synchronization failed.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of initial full data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that have been synchronized during initial full data synchronization.
	//
	// example:
	//
	// 43071
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of initial full data synchronization. Valid values:
	//
	// - **NotStarted**: The task is not started.
	//
	// - **Migrating**: The task is in progress.
	//
	// - **Failed**: The task failed.
	//
	// - **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus struct {
	// The error message returned if incremental data synchronization failed.
	//
	// example:
	//
	// The task has failed for a long time and cannot be recovered.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance needs to be upgraded. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// > To upgrade a DTS instance, call the [TransferInstanceClass](https://help.aliyun.com/document_detail/281093.html) operation.
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of incremental data synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of entries that have been migrated or synchronized during incremental data migration or synchronization.
	//
	// example:
	//
	// 20001
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of incremental data synchronization.
	//
	// example:
	//
	// Migrating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint struct {
	// The name of the database that contains the synchronized objects in the destination instance.
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
	// The ID of the destination instance.
	//
	// example:
	//
	// rm-bp1162kryivb8****
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
	// 172.16.88.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the returned value of **EngineName*	- of the destination instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The port number of the destination instance.
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
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// - **DISABLE**: SSL encryption is disabled.
	//
	// - **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// - **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection with an AWS MongoDB Altas database.
	//
	// - **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection with a Kafka cluster.
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

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) SetDatabaseName(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) SetEngineName(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) SetInstanceID(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) SetInstanceType(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) SetIp(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) SetOracleSID(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) SetPort(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) SetRegion(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) SetUserName(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode struct {
	// Indicates whether full data migration or initial full data synchronization is performed. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Indicates whether incremental data migration or synchronization is performed. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Indicates whether schema migration or initial schema synchronization is performed. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode) SetDataInitialization(v bool) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode) SetDataSynchronization(v bool) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode) SetStructureInitialization(v bool) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance struct {
	// The size of data that is migrated or synchronized per second. Unit: MB/s.
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

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance) GetFlow() *string {
	return s.Flow
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance) GetRps() *string {
	return s.Rps
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance) SetFlow(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance {
	s.Flow = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance) SetRps(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance {
	s.Rps = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPerformance) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus struct {
	// The result of each precheck item.
	Detail []*DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The error message returned if the precheck failed.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The precheck progress. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck state. Valid values:
	//
	// - **NotStarted**: The precheck is not started.
	//
	// - **Suspending**: The precheck is paused.
	//
	// - **Checking**: The precheck is in progress.
	//
	// - **Failed**: The precheck failed.
	//
	// - **Finished**: The precheck is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus) GetDetail() []*DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus) SetDetail(v []*DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail struct {
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
	// - **Success**
	//
	// - **Failed**
	//
	// example:
	//
	// Success
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The error message returned if the task failed to pass the precheck.
	//
	// > This parameter is returned only if the returned value of **CheckResult*	- is **Failed**.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The method to fix a precheck failure.
	//
	// > This parameter is returned only if the returned value of **CheckResult*	- is **Failed**.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) GetCheckItemDescription() *string {
	return s.CheckItemDescription
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) GetCheckResult() *string {
	return s.CheckResult
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) GetFailedReason() *string {
	return s.FailedReason
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) SetCheckItem(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) SetCheckItemDescription(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) SetCheckResult(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) SetFailedReason(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail {
	s.FailedReason = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) SetRepairMethod(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint struct {
	// The name of the database that contains the objects to be migrated from the source instance.
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
	// The ID of the source instance.
	//
	// example:
	//
	// rm-bp1imrtn6fq7h****
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
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the returned value of **EngineName*	- of the source instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The port number of the source instance.
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
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// - **DISABLE**: SSL encryption is disabled.
	//
	// - **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// - **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection with an AWS MongoDB Altas database.
	//
	// - **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection with a Kafka cluster.
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

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) SetDatabaseName(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) SetEngineName(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) SetInstanceID(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) SetInstanceType(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) SetIp(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) SetOracleSID(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) SetPort(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) SetRegion(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) SetUserName(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus struct {
	// The error message returned if initial schema synchronization failed.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: ERROR: type "geometry" does not exist;
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of initial schema synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables that have been synchronized during initial schema synchronization.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of initial schema synchronization. Valid values:
	//
	// - **NotStarted**: The task is not started.
	//
	// - **Migrating**: The task is in progress.
	//
	// - **Failed**: The task failed.
	//
	// - **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListReverseJobStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint struct {
	// The name of the database that contains the objects to be migrated from the source instance.
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
	// The ID of the source instance.
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
	// The SID of the Oracle database.
	//
	// > This parameter is returned only if the returned value of **EngineName*	- of the source instance is **Oracle*	- and the Oracle database is deployed in a non-RAC architecture.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The port number of the source instance.
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
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// - **DISABLE**: SSL encryption is disabled.
	//
	// - **ENABLE_WITH_CERTIFICATE**: SSL encryption is enabled and the CA certificate is uploaded.
	//
	// - **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL encryption is enabled for the connection with an AWS MongoDB Altas database.
	//
	// - **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to encrypt the connection with a Kafka cluster.
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

func (s DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) SetDatabaseName(v string) *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) SetEngineName(v string) *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) SetInstanceID(v string) *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) SetInstanceType(v string) *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) SetIp(v string) *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) SetOracleSID(v string) *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) SetPort(v string) *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) SetRegion(v string) *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) SetUserName(v string) *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus struct {
	// The error message returned if schema migration or initial schema synchronization failed.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: ERROR: type "geometry" does not exist;
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of schema migration or initial schema synchronization. Unit: percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables that have been migrated or synchronized during schema migration or initial schema synchronization.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of schema migration or initial schema synchronization. Valid values:
	//
	// - **NotStarted**: The task is not started.
	//
	// - **Migrating**: The task is in progress.
	//
	// - **Failed**: The task failed.
	//
	// - **Finished**: The task is complete.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDtsJobsResponseBodyEtlDemoListTagList struct {
	// The tag key.
	//
	// example:
	//
	// testkey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testvalue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDtsJobsResponseBodyEtlDemoListTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyEtlDemoListTagList) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListTagList) SetTagKey(v string) *DescribeDtsJobsResponseBodyEtlDemoListTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListTagList) SetTagValue(v string) *DescribeDtsJobsResponseBodyEtlDemoListTagList {
	s.TagValue = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyEtlDemoListTagList) Validate() error {
	return dara.Validate(s)
}
