// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstanceDtsMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *DescribeGadInstanceDtsMembersResponseBody
	GetCreateTime() *int64
	SetDbEngineType(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetDbEngineType() *string
	SetDtsJobList(v []*DescribeGadInstanceDtsMembersResponseBodyDtsJobList) *DescribeGadInstanceDtsMembersResponseBody
	GetDtsJobList() []*DescribeGadInstanceDtsMembersResponseBodyDtsJobList
	SetDynamicCode(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetHttpStatusCode() *string
	SetInstanceId(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetInstanceName() *string
	SetPageNumber(v int32) *DescribeGadInstanceDtsMembersResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeGadInstanceDtsMembersResponseBody
	GetPageRecordCount() *int32
	SetRegionId(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetStatus() *string
	SetSuccess(v string) *DescribeGadInstanceDtsMembersResponseBody
	GetSuccess() *string
	SetTotalRecordCount(v int32) *DescribeGadInstanceDtsMembersResponseBody
	GetTotalRecordCount() *int32
}

type DescribeGadInstanceDtsMembersResponseBody struct {
	// The creation time.
	//
	// example:
	//
	// 2021-06-28 17:34:53.0
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The database engine.
	//
	// example:
	//
	// MySQL
	DbEngineType *string `json:"DbEngineType,omitempty" xml:"DbEngineType,omitempty"`
	// The list of tasks and the execution details of each task.
	DtsJobList []*DescribeGadInstanceDtsMembersResponseBodyDtsJobList `json:"DtsJobList,omitempty" xml:"DtsJobList,omitempty" type:"Repeated"`
	// The dynamic error code. This parameter will be deprecated soon.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the **%s*	- in the **ErrMessage*	- response parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the specified request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// present environment is not support,so skip.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned when the call fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned when a call error occurs.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the active-active instance group.
	//
	// example:
	//
	// rm-bp1i99e8l7913****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the active-active instance group.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The region ID of the instance. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C166D79D-436B-45F0-B5A5-25E1959F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the active geo-redundancy instance group.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetDbEngineType() *string {
	return s.DbEngineType
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetDtsJobList() []*DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	return s.DtsJobList
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeGadInstanceDtsMembersResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetCreateTime(v int64) *DescribeGadInstanceDtsMembersResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetDbEngineType(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.DbEngineType = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetDtsJobList(v []*DescribeGadInstanceDtsMembersResponseBodyDtsJobList) *DescribeGadInstanceDtsMembersResponseBody {
	s.DtsJobList = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetDynamicCode(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetDynamicMessage(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetErrCode(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetErrMessage(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetHttpStatusCode(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetInstanceId(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetInstanceName(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetPageNumber(v int32) *DescribeGadInstanceDtsMembersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetPageRecordCount(v int32) *DescribeGadInstanceDtsMembersResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetRegionId(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetRequestId(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetResourceGroupId(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetSuccess(v string) *DescribeGadInstanceDtsMembersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) SetTotalRecordCount(v int32) *DescribeGadInstanceDtsMembersResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBody) Validate() error {
	if s.DtsJobList != nil {
		for _, item := range s.DtsJobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobList struct {
	// The application name.
	//
	// example:
	//
	// new
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The start time of the timestamp range saved in the change tracking instance, in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is in UTC.
	//
	// example:
	//
	// 2021-03-15T08:25:34Z
	BeginTimestamp *string `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	// The checkpoint.
	//
	// example:
	//
	// 1616899019
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The downstream consumption checkpoint of the change tracking instance, in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is in UTC.
	//
	// example:
	//
	// 2021-03-23T07:30:31Z
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	// The downstream consumption client information, which consists of the IP address and a random number generated by DTS.
	//
	// example:
	//
	// 114...:dts******
	ConsumptionClient *string `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	// The current CPU utilization of the instance. Unit: %.
	//
	// example:
	//
	// 1
	CpuUsage *string `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty"`
	// The task creation time, in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is in UTC.
	//
	// example:
	//
	// 2021-06-28 17:34:53.0
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The status of the physical migration to the cloud task.
	DataCloudStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus `json:"DataCloudStatus,omitempty" xml:"DataCloudStatus,omitempty" type:"Struct"`
	// The status of the ETL task.
	//
	// > This parameter collection is returned only after an ETL task is configured.
	DataEtlStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus `json:"DataEtlStatus,omitempty" xml:"DataEtlStatus,omitempty" type:"Struct"`
	// The status of full data migration or initial data synchronization.
	DataInitializationStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The status of incremental data migration or synchronization.
	DataSynchronizationStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The migration, synchronization, or change tracking objects. For more information about the definitions, see [Description of migration, synchronization, or change tracking objects](https://help.aliyun.com/document_detail/209545.html).
	//
	// example:
	//
	// {"dtstestdata": { "name": "dtstestdata", "all": true }}
	DbObject *string `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	// The ID of the dedicated DTS cluster.
	//
	// example:
	//
	// dtscluster_ft7y3**********
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The latency of incremental data migration or synchronization.
	//
	// > If you query a data migration task, the unit of this parameter is milliseconds. If you query a data synchronization task, the unit of this parameter is seconds.
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The connection information of the destination instance.
	DestinationEndpoint *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
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
	// The ID of the data migration, data synchronization, or subscribe instance.
	//
	// example:
	//
	// dtsi03e3zty16i****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The link specifications.
	//
	// > For more information about the specifications and performance tests of each link, see [Specifications for data migration links](https://help.aliyun.com/document_detail/26606.html) and [Specifications for data synchronization links](https://help.aliyun.com/document_detail/26605.html).
	//
	// example:
	//
	// large
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**: Forward.
	//
	// - **Reverse**: Reverse.
	//
	// > This parameter is returned only when the topology of the data synchronization instance is two-way synchronization.
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
	// The name of the DTS task.
	//
	// example:
	//
	// TestNg_polardb_cn-beijing_cn-beijing_polardb
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The used resources of the dedicated DTS cluster. Unit: DU.
	//
	// example:
	//
	// 15
	DuUsage *int64 `json:"DuUsage,omitempty" xml:"DuUsage,omitempty"`
	// The end backup timestamp.
	//
	// example:
	//
	// 2021-03-26T14:03:21Z
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The error details.
	ErrorDetails []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails `json:"ErrorDetails,omitempty" xml:"ErrorDetails,omitempty" type:"Repeated"`
	// The error message for a failed task.
	//
	// example:
	//
	// The deviceName is empty.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The safe checkpoint of the ETL task.
	//
	// example:
	//
	// 1610540493
	EtlSafeCheckpoint *string `json:"EtlSafeCheckpoint,omitempty" xml:"EtlSafeCheckpoint,omitempty"`
	// The expiration time of the instance, in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is in UTC.
	//
	// > This parameter is returned only when the **PayType*	- parameter is set to **PrePaid**.
	//
	// example:
	//
	// 2022-03-16T08:01:19Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The status information of full data verification.
	FullDataCheckStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus `json:"FullDataCheckStatus,omitempty" xml:"FullDataCheckStatus,omitempty" type:"Struct"`
	// The status information of incremental data verification.
	IncDataCheckStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus `json:"IncDataCheckStatus,omitempty" xml:"IncDataCheckStatus,omitempty" type:"Struct"`
	// The DTS task type.
	//
	// example:
	//
	// SYNC
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The DU upper limit of the serverless instance.
	//
	// example:
	//
	// 4.0
	MaxDu *float64 `json:"MaxDu,omitempty" xml:"MaxDu,omitempty"`
	// The used memory. Unit: MB.
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
	// The ID of the help document for the error code.
	//
	// example:
	//
	// 462133
	MigrationErrHelpDocId *string `json:"MigrationErrHelpDocId,omitempty" xml:"MigrationErrHelpDocId,omitempty"`
	// The key of the help document for the error code.
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
	// The error workaround.
	//
	// example:
	//
	// dts.retry.err.0046.workaround
	MigrationErrWorkaround *string `json:"MigrationErrWorkaround,omitempty" xml:"MigrationErrWorkaround,omitempty"`
	// The migration type or initial synchronization type.
	MigrationMode *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// The minimum number of DUs for the serverless instance.
	//
	// example:
	//
	// 1
	MinDu *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	// The source of the task. Valid values:
	//
	// - **PTS task**
	//
	// - **DMS task**
	//
	// - **DTS task**
	//
	// example:
	//
	// DTS
	OriginType *string `json:"OriginType,omitempty" xml:"OriginType,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PrePaid**: Subscription.
	//
	// - **PostPaid**: Pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The performance information of the migration or synchronization link.
	Performance *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck status.
	PrecheckStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The reserved parameter for DTS. The value is a JSON string. You can specify this parameter to meet special requirements, such as whether to automatically start the precheck. For more information, see [MigrationReserved parameter description](https://help.aliyun.com/document_detail/176470.html).
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
	// The retry state of DTS due to a task exception.
	RetryState *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState `json:"RetryState,omitempty" xml:"RetryState,omitempty" type:"Struct"`
	// The execution details of the reverse synchronization task.
	//
	// > This parameter collection is returned only when the task is a data synchronization task and the topology type is two-way.
	ReverseJob *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob `json:"ReverseJob,omitempty" xml:"ReverseJob,omitempty" type:"Struct"`
	// The connection information of the source instance.
	SourceEndpoint *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The DTS instance status. Valid values:
	//
	// Status of a data migration node:
	//
	// - **NotStarted**: The node is not started.
	//
	// - **Prechecking**: The precheck is in progress.
	//
	// - **PrecheckFailed**: The precheck failed.
	//
	// - **PreCheckPass**: The precheck passed.
	//
	// - **NotConfigured**: The node is not configured.
	//
	// - **Migrating**: The migration is in progress.
	//
	// - **Suspending**: The node is being paused.
	//
	// - **MigrationFailed**: The migration failed.
	//
	// - **Finished**: The node is completed.
	//
	// - **Retrying**: The node is being retried.
	//
	// - **Upgrade**: The node is being upgraded.
	//
	// - **Locked**: The node is locked.
	//
	// - **Downgrade**: The node is being downgraded.
	//
	// Status of a data synchronization node:
	//
	// - **NotStarted**: The node is not started.
	//
	// - **Prechecking**: The precheck is in progress.
	//
	// - **PrecheckFailed**: The precheck failed.
	//
	// - **PreCheckPass**: The precheck passed.
	//
	// - **NotConfigured**: The node is not configured.
	//
	// - **Initializing**: Initial synchronization is in progress.
	//
	// - **InitializeFailed**: Initial synchronization failed.
	//
	// - **Synchronizing**: Data synchronization is in progress.
	//
	// - **Failed**: The synchronization failed.
	//
	// - **Suspending**: The node is being paused.
	//
	// - **Modifying**: The synchronization objects are being modified.
	//
	// - **Finished**: The node is completed.
	//
	// - **Retrying**: The node is being retried.
	//
	// - **Upgrade**: The node is being upgraded.
	//
	// - **Locked**: The node is locked.
	//
	// - **Downgrade**: The node is being downgraded.
	//
	// Status of a subscribe node:
	//
	// - **NotConfigured**: The node is not configured.
	//
	// - **NotStarted**: The node is not started.
	//
	// - **Prechecking**: The precheck is in progress.
	//
	// - **PrecheckFailed**: The precheck failed.
	//
	// - **PreCheckPass**: The precheck passed.
	//
	// - **Starting**: The node is being started.
	//
	// - **Normal**: The node is running normally.
	//
	// - **Retrying**: The node is being retried.
	//
	// - **Abnormal**: The node is abnormal.
	//
	// - **Upgrade**: The node is being upgraded.
	//
	// - **Locked**: The node is locked.
	//
	// - **Downgrade**: The node is being downgraded.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of schema migration or initial schema synchronization.
	StructureInitializationStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	// The list of tags.
	TagList []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobList) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetAppName() *string {
	return s.AppName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetBeginTimestamp() *string {
	return s.BeginTimestamp
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetConsumptionCheckpoint() *string {
	return s.ConsumptionCheckpoint
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetConsumptionClient() *string {
	return s.ConsumptionClient
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetCpuUsage() *string {
	return s.CpuUsage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDataCloudStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus {
	return s.DataCloudStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDataEtlStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus {
	return s.DataEtlStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDataInitializationStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDataSynchronizationStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDbObject() *string {
	return s.DbObject
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDestinationEndpoint() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDtsBisLabel() *string {
	return s.DtsBisLabel
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDtsJobDirection() *string {
	return s.DtsJobDirection
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetDuUsage() *int64 {
	return s.DuUsage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetErrorDetails() []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails {
	return s.ErrorDetails
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetEtlSafeCheckpoint() *string {
	return s.EtlSafeCheckpoint
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetFullDataCheckStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus {
	return s.FullDataCheckStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetIncDataCheckStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus {
	return s.IncDataCheckStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetJobType() *string {
	return s.JobType
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetMaxDu() *float64 {
	return s.MaxDu
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetMemUsage() *string {
	return s.MemUsage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetMigrationErrCode() *string {
	return s.MigrationErrCode
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetMigrationErrHelpDocId() *string {
	return s.MigrationErrHelpDocId
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetMigrationErrHelpDocKey() *string {
	return s.MigrationErrHelpDocKey
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetMigrationErrMsg() *string {
	return s.MigrationErrMsg
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetMigrationErrType() *string {
	return s.MigrationErrType
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetMigrationErrWorkaround() *string {
	return s.MigrationErrWorkaround
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetMigrationMode() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode {
	return s.MigrationMode
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetMinDu() *float64 {
	return s.MinDu
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetOriginType() *string {
	return s.OriginType
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetPayType() *string {
	return s.PayType
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetPerformance() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance {
	return s.Performance
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetPrecheckStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetReserved() *string {
	return s.Reserved
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetResourceGroupDisplayName() *string {
	return s.ResourceGroupDisplayName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetRetryState() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	return s.RetryState
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetReverseJob() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	return s.ReverseJob
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetSourceEndpoint() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetStructureInitializationStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) GetTagList() []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList {
	return s.TagList
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetAppName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.AppName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetBeginTimestamp(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetCheckpoint(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.Checkpoint = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetConsumptionCheckpoint(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetConsumptionClient(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetCpuUsage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.CpuUsage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetCreateTime(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.CreateTime = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDataCloudStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DataCloudStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDataEtlStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DataEtlStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDataInitializationStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDataSynchronizationStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDbObject(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DbObject = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDedicatedClusterId(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDelay(v int64) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.Delay = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDestinationEndpoint(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDtsBisLabel(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DtsBisLabel = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDtsInstanceID(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDtsJobClass(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDtsJobDirection(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDtsJobId(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DtsJobId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDtsJobName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DtsJobName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetDuUsage(v int64) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.DuUsage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetEndTimestamp(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetErrorDetails(v []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.ErrorDetails = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetEtlSafeCheckpoint(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.EtlSafeCheckpoint = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetExpireTime(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetFullDataCheckStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.FullDataCheckStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetIncDataCheckStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.IncDataCheckStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetJobType(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.JobType = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetMaxDu(v float64) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.MaxDu = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetMemUsage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.MemUsage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetMigrationErrCode(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.MigrationErrCode = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetMigrationErrHelpDocId(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.MigrationErrHelpDocId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetMigrationErrHelpDocKey(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.MigrationErrHelpDocKey = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetMigrationErrMsg(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.MigrationErrMsg = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetMigrationErrType(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.MigrationErrType = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetMigrationErrWorkaround(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.MigrationErrWorkaround = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetMigrationMode(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.MigrationMode = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetMinDu(v float64) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.MinDu = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetOriginType(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.OriginType = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetPayType(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.PayType = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetPerformance(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.Performance = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetPrecheckStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetReserved(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.Reserved = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetResourceGroupDisplayName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.ResourceGroupDisplayName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetResourceGroupId(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetRetryState(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.RetryState = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetReverseJob(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.ReverseJob = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetSourceEndpoint(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetStructureInitializationStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) SetTagList(v []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList) *DescribeGadInstanceDtsMembersResponseBodyDtsJobList {
	s.TagList = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobList) Validate() error {
	if s.DataCloudStatus != nil {
		if err := s.DataCloudStatus.Validate(); err != nil {
			return err
		}
	}
	if s.DataEtlStatus != nil {
		if err := s.DataEtlStatus.Validate(); err != nil {
			return err
		}
	}
	if s.DataInitializationStatus != nil {
		if err := s.DataInitializationStatus.Validate(); err != nil {
			return err
		}
	}
	if s.DataSynchronizationStatus != nil {
		if err := s.DataSynchronizationStatus.Validate(); err != nil {
			return err
		}
	}
	if s.DestinationEndpoint != nil {
		if err := s.DestinationEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.ErrorDetails != nil {
		for _, item := range s.ErrorDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FullDataCheckStatus != nil {
		if err := s.FullDataCheckStatus.Validate(); err != nil {
			return err
		}
	}
	if s.IncDataCheckStatus != nil {
		if err := s.IncDataCheckStatus.Validate(); err != nil {
			return err
		}
	}
	if s.MigrationMode != nil {
		if err := s.MigrationMode.Validate(); err != nil {
			return err
		}
	}
	if s.Performance != nil {
		if err := s.Performance.Validate(); err != nil {
			return err
		}
	}
	if s.PrecheckStatus != nil {
		if err := s.PrecheckStatus.Validate(); err != nil {
			return err
		}
	}
	if s.RetryState != nil {
		if err := s.RetryState.Validate(); err != nil {
			return err
		}
	}
	if s.ReverseJob != nil {
		if err := s.ReverseJob.Validate(); err != nil {
			return err
		}
	}
	if s.SourceEndpoint != nil {
		if err := s.SourceEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.StructureInitializationStatus != nil {
		if err := s.StructureInitializationStatus.Validate(); err != nil {
			return err
		}
	}
	if s.TagList != nil {
		for _, item := range s.TagList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus struct {
	// The error message for a failed task.
	//
	// example:
	//
	// The specified product does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance specifications need to be upgraded. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The task progress. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables that have been migrated.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of the physical migration to the cloud task.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) SetNeedUpgrade(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) SetProgress(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus {
	s.Progress = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataCloudStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus struct {
	// The error message for a failed ETL task.
	//
	// example:
	//
	// The specified product does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ETL task progress. Unit: %.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of records for which the ETL task is completed.
	//
	// example:
	//
	// 75
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The log level of the log information. Multiple statuses are separated by commas (,). Valid values:
	//
	// - **NORMAL**: Normal.
	//
	// - **WARN**: Warning.
	//
	// - **ERROR**: Error.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus) SetProgress(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus {
	s.Progress = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataEtlStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus struct {
	// The error message for failed full data migration or initial data synchronization.
	//
	// example:
	//
	// The deviceName is empty.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of full data migration or initial data synchronization. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of records for which full data migration or initial data synchronization is completed.
	//
	// example:
	//
	// 35
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of full data migration or initial data synchronization.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus) SetProgress(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus struct {
	// The error message for failed incremental data migration or synchronization.
	//
	// example:
	//
	// The deviceName format is incorrect.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance specifications need to be upgraded. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// > To upgrade the instance specifications, call the [TransferInstanceClass](https://help.aliyun.com/document_detail/281093.html) operation.
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of incremental data migration or synchronization. Unit: %.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of records for which incremental data migration or synchronization is completed.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The log level of the log information. Multiple statuses are separated by commas (,). Valid values:
	//
	// - **NORMAL**: Normal.
	//
	// - **WARN**: Warning.
	//
	// - **ERROR**: Error.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) SetProgress(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint struct {
	// The name of the database to which the migration objects in the instance belong.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine type of the instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-bp1imrtn6fq7h****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The instance type.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The endpoint of the instance.
	//
	// example:
	//
	// 172.16.88.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only when the **EngineName*	- parameter of the destination instance is set to **Oracle*	- and the Oracle database is a non-RAC instance.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region where the instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Indicates whether an SSL-secured connection is used. Valid values:
	//
	// - **DISABLE**: SSL-secured connections are not used.
	//
	// - **ENABLE_WITH_CERTIFICATE**: SSL-secured connections are used, and a CA certificate is uploaded.
	//
	// - **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL-secured connections to AWS MongoDB Atlas are used.
	//
	// - **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to connect to Kafka.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The username.
	//
	// example:
	//
	// dts
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) SetDatabaseName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) SetEngineName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) SetInstanceID(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) SetInstanceType(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) SetIp(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) SetOracleSID(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) SetPort(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) SetRegion(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) SetUserName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails struct {
	// The error code.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The URL of the help document.
	//
	// example:
	//
	// https://**.ali**.com/**
	HelpUrl *string `json:"HelpUrl,omitempty" xml:"HelpUrl,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails) GetHelpUrl() *string {
	return s.HelpUrl
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails) SetErrorCode(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails {
	s.ErrorCode = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails) SetHelpUrl(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails {
	s.HelpUrl = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListErrorDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus struct {
	// The error message for a failed task.
	//
	// example:
	//
	// device status is offline
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of the full data verification task. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The completion progress of the full data verification task.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status information of full data verification.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus) SetProgress(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus {
	s.Progress = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListFullDataCheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus struct {
	// The error message for a failed task.
	//
	// example:
	//
	// voice intercom existed
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of the incremental data verification task. Unit: %.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The completion progress of the incremental data verification task.
	//
	// example:
	//
	// 5
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status information of incremental data verification.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus) SetProgress(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus {
	s.Progress = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListIncDataCheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode struct {
	// Indicates whether full data migration or initial full data synchronization is performed. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Indicates whether incremental data migration or synchronization is performed. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Indicates whether full data validation is performed. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	FullDataCheck *bool `json:"FullDataCheck,omitempty" xml:"FullDataCheck,omitempty"`
	// Indicates whether incremental data validation is performed. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	IncDataCheck *bool `json:"IncDataCheck,omitempty" xml:"IncDataCheck,omitempty"`
	// Indicates whether schema migration or initial schema synchronization is performed. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) GetFullDataCheck() *bool {
	return s.FullDataCheck
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) GetIncDataCheck() *bool {
	return s.IncDataCheck
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) SetDataInitialization(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) SetDataSynchronization(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) SetFullDataCheck(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode {
	s.FullDataCheck = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) SetIncDataCheck(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode {
	s.IncDataCheck = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) SetStructureInitialization(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance struct {
	// The data migration or synchronization traffic per second, in MB/s.
	//
	// example:
	//
	// 1
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The number of SQL statements migrated or synchronized per second, including BEGIN, COMMIT, DML statements (INSERT, DELETE, and UPDATE), and DDL statements.
	//
	// example:
	//
	// 100
	Rps *string `json:"Rps,omitempty" xml:"Rps,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance) GetFlow() *string {
	return s.Flow
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance) GetRps() *string {
	return s.Rps
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance) SetFlow(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance {
	s.Flow = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance) SetRps(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance {
	s.Rps = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPerformance) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus struct {
	// The execution status of each precheck item.
	Detail []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The error message for the precheck failure.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The overall progress of the precheck, in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck status.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus) GetDetail() []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus) SetDetail(v []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatus) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail struct {
	// The precheck item.
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
	// The check result. Valid values:
	//
	// - **Success**: The item passed the check.
	//
	// - **Failed**: The item failed the check.
	//
	// example:
	//
	// Success
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The error message for the failed precheck.
	//
	// > This parameter is returned only when the **CheckResult*	- parameter is set to **Failed**.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The solution for the failed precheck.
	//
	// > This parameter is returned only when the **CheckResult*	- parameter is set to **Failed**.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) GetCheckItemDescription() *string {
	return s.CheckItemDescription
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) GetCheckResult() *string {
	return s.CheckResult
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) GetFailedReason() *string {
	return s.FailedReason
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) SetCheckItem(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) SetCheckItemDescription(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) SetCheckResult(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) SetFailedReason(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail {
	s.FailedReason = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) SetRepairMethod(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState struct {
	// The error message returned when a call error occurs.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The task ID.
	//
	// example:
	//
	// bi6e22ay243****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum retry time, in seconds.
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
	// The ID of the help document for the error code.
	//
	// example:
	//
	// 462133
	MigrationErrHelpDocId *string `json:"MigrationErrHelpDocId,omitempty" xml:"MigrationErrHelpDocId,omitempty"`
	// The key of the help document for the error code.
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
	// The error workaround.
	//
	// example:
	//
	// dts.retry.err.0046.workaround
	MigrationErrWorkaround *string `json:"MigrationErrWorkaround,omitempty" xml:"MigrationErrWorkaround,omitempty"`
	// The instance progress.
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
	// The retry target. Valid values:
	//
	// - **srcDB**: The source database.
	//
	// - **destDB**: The destination database.
	//
	// - **inner_module**: The internal module of DTS.
	//
	// example:
	//
	// srcDB
	RetryTarget *string `json:"RetryTarget,omitempty" xml:"RetryTarget,omitempty"`
	// The elapsed retry time, in seconds.
	//
	// example:
	//
	// 3600
	RetryTime *int32 `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
	// Indicates whether the task is being retried. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// false
	Retrying *bool `json:"Retrying,omitempty" xml:"Retrying,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetJobId() *string {
	return s.JobId
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetMaxRetryTime() *int32 {
	return s.MaxRetryTime
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetMigrationErrCode() *string {
	return s.MigrationErrCode
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetMigrationErrHelpDocId() *string {
	return s.MigrationErrHelpDocId
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetMigrationErrHelpDocKey() *string {
	return s.MigrationErrHelpDocKey
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetMigrationErrMsg() *string {
	return s.MigrationErrMsg
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetMigrationErrType() *string {
	return s.MigrationErrType
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetMigrationErrWorkaround() *string {
	return s.MigrationErrWorkaround
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetModule() *string {
	return s.Module
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetRetryCount() *int32 {
	return s.RetryCount
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetRetryTarget() *string {
	return s.RetryTarget
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetRetryTime() *int32 {
	return s.RetryTime
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) GetRetrying() *bool {
	return s.Retrying
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetErrMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.ErrMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetJobId(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.JobId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetMaxRetryTime(v int32) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.MaxRetryTime = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetMigrationErrCode(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.MigrationErrCode = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetMigrationErrHelpDocId(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.MigrationErrHelpDocId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetMigrationErrHelpDocKey(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.MigrationErrHelpDocKey = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetMigrationErrMsg(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.MigrationErrMsg = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetMigrationErrType(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.MigrationErrType = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetMigrationErrWorkaround(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.MigrationErrWorkaround = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetModule(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.Module = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetRetryCount(v int32) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.RetryCount = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetRetryTarget(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.RetryTarget = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetRetryTime(v int32) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.RetryTime = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) SetRetrying(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState {
	s.Retrying = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListRetryState) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob struct {
	// The checkpoint for incremental data synchronization. The value is a UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1616980369
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The current CPU utilization of the instance. Unit: %.
	//
	// example:
	//
	// 90
	CpuUsage *string `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty"`
	// The task creation time, in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is in UTC.
	//
	// example:
	//
	// 2021-06-28 17:34:53.0
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The status of initial full data synchronization.
	DataInitializationStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	// The status of incremental data synchronization.
	DataSynchronizationStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	// The schema definition of the synchronized objects. The value is a JSON string and supports specific regular expressions. For more information and examples, see the object configuration instructions for migration, synchronization, or change tracking.
	//
	// example:
	//
	// {"dtstestdata": { "name": "dtstestdata", "all": true }}
	DbObject *string `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	// The ID of the dedicated DTS cluster.
	//
	// example:
	//
	// dtscluster_dpwl3**********
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The latency, in seconds.
	//
	// example:
	//
	// 0
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The connection information of the destination instance.
	DestinationEndpoint *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	// The instance ID of the data synchronization instance.
	//
	// example:
	//
	// dtsi03e3zty16i****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The specifications of the synchronization link.
	//
	// > For more information about the specifications and performance test results of each link, see [Specifications of data synchronization links](https://help.aliyun.com/document_detail/26605.html).
	//
	// example:
	//
	// large
	DtsJobClass *string `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	// The description of the DTS task.
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
	// The name of the DTS task.
	//
	// example:
	//
	// test2
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	// The used resources of the dedicated DTS cluster. Unit: DU.
	//
	// example:
	//
	// 15
	DuUsage *int64 `json:"DuUsage,omitempty" xml:"DuUsage,omitempty"`
	// The error details.
	ErrorDetails []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails `json:"ErrorDetails,omitempty" xml:"ErrorDetails,omitempty" type:"Repeated"`
	// The error message for a failed task.
	//
	// example:
	//
	// Instance not found.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The safe checkpoint of the ETL task.
	//
	// example:
	//
	// 1610540493
	EtlSafeCheckpoint *string `json:"EtlSafeCheckpoint,omitempty" xml:"EtlSafeCheckpoint,omitempty"`
	// The expiration time of the instance, in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is in UTC.
	//
	// > This parameter is returned only when the **PayType*	- parameter is set to **PrePaid**.
	//
	// example:
	//
	// 2023-03-16T08:01:19Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The status information of full data verification.
	FullDataCheckStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus `json:"FullDataCheckStatus,omitempty" xml:"FullDataCheckStatus,omitempty" type:"Struct"`
	// The status information of incremental data verification.
	IncDataCheckStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus `json:"IncDataCheckStatus,omitempty" xml:"IncDataCheckStatus,omitempty" type:"Struct"`
	// The DU upper limit of the serverless instance.
	//
	// example:
	//
	// 16
	MaxDu *float64 `json:"MaxDu,omitempty" xml:"MaxDu,omitempty"`
	// The amount of used memory, in MB.
	//
	// example:
	//
	// 500
	MemUsage *string `json:"MemUsage,omitempty" xml:"MemUsage,omitempty"`
	// The type of initial synchronization.
	MigrationMode *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	// The minimum number of DUs for the serverless instance.
	//
	// example:
	//
	// 1
	MinDu *float64 `json:"MinDu,omitempty" xml:"MinDu,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PrePaid**: Subscription.
	//
	// - **PostPaid**: Pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The performance information of the synchronization link.
	Performance *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	// The precheck status.
	PrecheckStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	// The reserved parameters of DTS. The value is a JSON string. You can specify this parameter to meet special requirements, such as specifying whether to automatically start the precheck. For more information, see [MigrationReserved parameter description](https://help.aliyun.com/document_detail/176470.html).
	//
	// example:
	//
	// {\\"srcHostPorts\\":\\"\\",\\"whitelist.dms.online.ddl.enable\\":false,\\"filterDDL\\":false,\\"sqlparser.dms.original.ddl\\":true,\\"srcOracleType\\":\\"sid\\",\\"maxRetryTime\\":43200,\\"destSSL\\":\\"0\\",\\"destOracleType\\":\\"sid\\",\\"srcSSL\\":\\"0\\",\\"dbListCaseChangeMode\\":\\"default\\",\\"SourceEngineVersion\\":\\"8.0.18\\",\\"srcNetType\\":\\"VPC\\",\\"destNetType\\":\\"VPC\\",\\"srcVpcNetMappingInst\\":\\"172.16.1**.**:10803\\",\\"destVpcNetMappingInst\\":\\"172.16.1**.**:11077\\",\\"useJobTask\\":\\"1\\"}
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The connection information of the source instance.
	SourceEndpoint *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	// The status of the reverse synchronization task.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of initial schema synchronization.
	StructureInitializationStatus *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetCpuUsage() *string {
	return s.CpuUsage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetDataInitializationStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus {
	return s.DataInitializationStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetDataSynchronizationStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	return s.DataSynchronizationStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetDbObject() *string {
	return s.DbObject
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetDestinationEndpoint() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint {
	return s.DestinationEndpoint
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetDtsJobClass() *string {
	return s.DtsJobClass
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetDtsJobDirection() *string {
	return s.DtsJobDirection
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetDtsJobName() *string {
	return s.DtsJobName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetDuUsage() *int64 {
	return s.DuUsage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetErrorDetails() []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails {
	return s.ErrorDetails
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetEtlSafeCheckpoint() *string {
	return s.EtlSafeCheckpoint
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetFullDataCheckStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus {
	return s.FullDataCheckStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetIncDataCheckStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus {
	return s.IncDataCheckStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetMaxDu() *float64 {
	return s.MaxDu
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetMemUsage() *string {
	return s.MemUsage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetMigrationMode() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode {
	return s.MigrationMode
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetMinDu() *float64 {
	return s.MinDu
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetPayType() *string {
	return s.PayType
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetPerformance() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance {
	return s.Performance
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetPrecheckStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus {
	return s.PrecheckStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetReserved() *string {
	return s.Reserved
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetSourceEndpoint() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint {
	return s.SourceEndpoint
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) GetStructureInitializationStatus() *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	return s.StructureInitializationStatus
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetCheckpoint(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.Checkpoint = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetCpuUsage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.CpuUsage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetCreateTime(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.CreateTime = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetDataInitializationStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetDataSynchronizationStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetDbObject(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.DbObject = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetDedicatedClusterId(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetDelay(v int64) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.Delay = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetDestinationEndpoint(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetDtsInstanceID(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetDtsJobClass(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetDtsJobDirection(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetDtsJobId(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.DtsJobId = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetDtsJobName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.DtsJobName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetDuUsage(v int64) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.DuUsage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetErrorDetails(v []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.ErrorDetails = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetEtlSafeCheckpoint(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.EtlSafeCheckpoint = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetExpireTime(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetFullDataCheckStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.FullDataCheckStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetIncDataCheckStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.IncDataCheckStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetMaxDu(v float64) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.MaxDu = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetMemUsage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.MemUsage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetMigrationMode(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.MigrationMode = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetMinDu(v float64) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.MinDu = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetPayType(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.PayType = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetPerformance(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.Performance = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetPrecheckStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetReserved(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.Reserved = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetSourceEndpoint(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) SetStructureInitializationStatus(v *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJob) Validate() error {
	if s.DataInitializationStatus != nil {
		if err := s.DataInitializationStatus.Validate(); err != nil {
			return err
		}
	}
	if s.DataSynchronizationStatus != nil {
		if err := s.DataSynchronizationStatus.Validate(); err != nil {
			return err
		}
	}
	if s.DestinationEndpoint != nil {
		if err := s.DestinationEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.ErrorDetails != nil {
		for _, item := range s.ErrorDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FullDataCheckStatus != nil {
		if err := s.FullDataCheckStatus.Validate(); err != nil {
			return err
		}
	}
	if s.IncDataCheckStatus != nil {
		if err := s.IncDataCheckStatus.Validate(); err != nil {
			return err
		}
	}
	if s.MigrationMode != nil {
		if err := s.MigrationMode.Validate(); err != nil {
			return err
		}
	}
	if s.Performance != nil {
		if err := s.Performance.Validate(); err != nil {
			return err
		}
	}
	if s.PrecheckStatus != nil {
		if err := s.PrecheckStatus.Validate(); err != nil {
			return err
		}
	}
	if s.SourceEndpoint != nil {
		if err := s.SourceEndpoint.Validate(); err != nil {
			return err
		}
	}
	if s.StructureInitializationStatus != nil {
		if err := s.StructureInitializationStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus struct {
	// The error message for the initial full data synchronization failure.
	//
	// example:
	//
	// The deviceName format is incorrect.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of initial full data synchronization, in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of records for which initial full data synchronization is completed.
	//
	// example:
	//
	// 43071
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of initial full data synchronization.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus) SetProgress(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus struct {
	// The error message for the incremental data synchronization failure.
	//
	// example:
	//
	// Instance not found.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the instance specifications need to be upgraded. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// > To upgrade the instance specifications, call the [TransferInstanceClass](https://help.aliyun.com/document_detail/281093.html) operation.
	//
	// example:
	//
	// true
	NeedUpgrade *bool `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	// The progress of incremental data synchronization, in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of records for which incremental data migration or synchronization is completed.
	//
	// example:
	//
	// 31
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of incremental data synchronization.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetProgress(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDataSynchronizationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint struct {
	// The name of the database to which the synchronized objects belong in the instance.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database engine type of the instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-bp1162kryivb8****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The instance type.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The endpoint of the instance.
	//
	// example:
	//
	// 172.16.88.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only when the **EngineName*	- parameter of the destination instance is set to **Oracle*	- and the Oracle database is a non-RAC instance.
	//
	// example:
	//
	// testsid
	OracleSID *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	// The database service port of the instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region where the instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Indicates whether an SSL-secured connection is used. Valid values:
	//
	// - **DISABLE**: SSL-secured connections are not used.
	//
	// - **ENABLE_WITH_CERTIFICATE**: SSL-secured connections are used, and a CA certificate is uploaded.
	//
	// - **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL-secured connections to AWS MongoDB Atlas are used.
	//
	// - **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to connect to Kafka.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The username.
	//
	// example:
	//
	// dts
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) SetDatabaseName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) SetEngineName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) SetInstanceID(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) SetInstanceType(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) SetIp(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) SetOracleSID(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) SetPort(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) SetRegion(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) SetUserName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobDestinationEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails struct {
	// The error code.
	//
	// example:
	//
	// DTS-31009
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The URL of the help document.
	//
	// example:
	//
	// https://**.ali**.com/**
	HelpUrl *string `json:"HelpUrl,omitempty" xml:"HelpUrl,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails) GetHelpUrl() *string {
	return s.HelpUrl
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails) SetErrorCode(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails {
	s.ErrorCode = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails) SetHelpUrl(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails {
	s.HelpUrl = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobErrorDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus struct {
	// The error message for a failed task.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of the full data verification task. Unit: %.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The completion progress of the full data verification task.
	//
	// example:
	//
	// 35
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status information of full data verification.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus) SetProgress(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus {
	s.Progress = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobFullDataCheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus struct {
	// The error message for a failed task.
	//
	// example:
	//
	// java.lang.NumberFormatException: For input string: ""
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of the incremental data verification task. Unit: %.
	//
	// example:
	//
	// 95
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The completion progress of the incremental data verification task.
	//
	// example:
	//
	// 35
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status information of incremental data verification.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus) SetProgress(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus {
	s.Progress = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobIncDataCheckStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode struct {
	// Specifies whether to perform initial full data synchronization. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Specifies whether to perform incremental data synchronization. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// Indicates whether full data validation is performed. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	FullDataCheck *bool `json:"FullDataCheck,omitempty" xml:"FullDataCheck,omitempty"`
	// Indicates whether incremental data validation is performed. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	IncDataCheck *bool `json:"IncDataCheck,omitempty" xml:"IncDataCheck,omitempty"`
	// Specifies whether to perform initial schema synchronization. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) GetFullDataCheck() *bool {
	return s.FullDataCheck
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) GetIncDataCheck() *bool {
	return s.IncDataCheck
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) SetDataInitialization(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) SetDataSynchronization(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) SetFullDataCheck(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode {
	s.FullDataCheck = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) SetIncDataCheck(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode {
	s.IncDataCheck = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) SetStructureInitialization(v bool) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobMigrationMode) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance struct {
	// The data synchronization traffic per second, in MB/s.
	//
	// example:
	//
	// 1
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The number of SQL statements synchronized per second, including BEGIN, COMMIT, DML statements (INSERT, DELETE, and UPDATE), and DDL statements.
	//
	// example:
	//
	// 100
	Rps *string `json:"Rps,omitempty" xml:"Rps,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance) GetFlow() *string {
	return s.Flow
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance) GetRps() *string {
	return s.Rps
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance) SetFlow(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance {
	s.Flow = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance) SetRps(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance {
	s.Rps = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPerformance) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus struct {
	// The execution status of each precheck item.
	Detail []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The error message for the precheck failure.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The overall progress of the precheck, in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The precheck status.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus) GetDetail() []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	return s.Detail
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus) SetDetail(v []*DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatus) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail struct {
	// The precheck item.
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
	// - **Success**: The item passed the precheck.
	//
	// - **Failed**: The item failed the precheck.
	//
	// example:
	//
	// Success
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// The error message returned when the precheck fails.
	//
	// example:
	//
	// Original error: Access denied for user \\"dtstest\\"@\\"100.104.***.**\\" (using password: YES)
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The solution provided when the precheck fails.
	//
	// example:
	//
	// CHECK_ERROR_DEST_CONN_REPAIR2
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GetCheckItem() *string {
	return s.CheckItem
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GetCheckItemDescription() *string {
	return s.CheckItemDescription
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GetCheckResult() *string {
	return s.CheckResult
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GetFailedReason() *string {
	return s.FailedReason
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetCheckItem(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetCheckItemDescription(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetCheckResult(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetFailedReason(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.FailedReason = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetRepairMethod(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobPrecheckStatusDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint struct {
	// The name of the database to which the migration objects in the instance belong.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database type of the source instance.
	//
	// example:
	//
	// MySQL
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The source instance ID.
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
	// The endpoint of the instance.
	//
	// example:
	//
	// 172.16.88.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The SID of the Oracle database.
	//
	// > This parameter is returned only when the **EngineName*	- parameter of the source instance is set to **Oracle*	- and the Oracle database is a non-RAC instance.
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
	// The region where the instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Indicates whether an SSL-secured connection is used. Valid values:
	//
	// - **DISABLE**: SSL-secured connections are not used.
	//
	// - **ENABLE_WITH_CERTIFICATE**: SSL-secured connections are used, and a CA certificate is uploaded.
	//
	// - **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL-secured connections to AWS MongoDB Atlas are used.
	//
	// - **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to connect to Kafka.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The username.
	//
	// example:
	//
	// dts
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) SetDatabaseName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) SetEngineName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) SetInstanceID(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) SetInstanceType(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) SetIp(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) SetOracleSID(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) SetPort(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) SetRegion(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) SetSslSolutionEnum(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) SetUserName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus struct {
	// The error message returned when initial schema synchronization fails.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: ERROR: type "geometry" does not exist;
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of initial schema synchronization, in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables for which initial schema synchronization is completed.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of initial schema synchronization.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus) SetProgress(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListReverseJobStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint struct {
	// The name of the database to which the migration objects in the source instance belong.
	//
	// example:
	//
	// dtstestdata
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database type of the source instance.
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
	// The SID of the Oracle database.
	//
	// > This parameter is returned only when the **EngineName*	- parameter of the source instance is set to **Oracle*	- and the Oracle database is a non-RAC instance.
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
	// The region where the source instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Indicates whether an SSL-secured connection is used. Valid values:
	//
	// - **DISABLE**: SSL-secured connections are not used.
	//
	// - **ENABLE_WITH_CERTIFICATE**: SSL-secured connections are used, and a CA certificate is uploaded.
	//
	// - **ENABLE_ONLY_4_MONGODB_ALTAS**: SSL-secured connections to AWS MongoDB Atlas are used.
	//
	// - **ENABLE_ONLY_4_KAFKA_SCRAM_SHA_256**: SCRAM-SHA-256 is used to connect to Kafka.
	//
	// example:
	//
	// DISABLE
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	// The username.
	//
	// example:
	//
	// dts
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) GetEngineName() *string {
	return s.EngineName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) GetIp() *string {
	return s.Ip
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) GetOracleSID() *string {
	return s.OracleSID
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) GetPort() *string {
	return s.Port
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) GetRegion() *string {
	return s.Region
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) GetSslSolutionEnum() *string {
	return s.SslSolutionEnum
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) GetUserName() *string {
	return s.UserName
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) SetDatabaseName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) SetEngineName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) SetInstanceID(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) SetInstanceType(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) SetIp(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) SetOracleSID(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) SetPort(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) SetRegion(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) SetSslSolutionEnum(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) SetUserName(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListSourceEndpoint) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus struct {
	// The error message returned when schema migration or initial schema synchronization fails.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: ERROR: type "geometry" does not exist;
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The progress of schema migration or initial schema synchronization, in percentage.
	//
	// example:
	//
	// 100
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The number of tables for which schema migration or initial schema synchronization is completed.
	//
	// example:
	//
	// 35
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of schema migration or initial schema synchronization.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus) GetPercent() *string {
	return s.Percent
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus) SetErrorMessage(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus) SetPercent(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus) SetProgress(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus) SetStatus(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListStructureInitializationStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList struct {
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

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList) GoString() string {
	return s.String()
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList) SetTagKey(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList) SetTagValue(v string) *DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList {
	s.TagValue = &v
	return s
}

func (s *DescribeGadInstanceDtsMembersResponseBodyDtsJobListTagList) Validate() error {
	return dara.Validate(s)
}
