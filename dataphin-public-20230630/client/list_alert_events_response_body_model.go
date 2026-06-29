// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAlertEventsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListAlertEventsResponseBody
	GetHttpStatusCode() *int32
	SetListResult(v *ListAlertEventsResponseBodyListResult) *ListAlertEventsResponseBody
	GetListResult() *ListAlertEventsResponseBodyListResult
	SetMessage(v string) *ListAlertEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAlertEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAlertEventsResponseBody
	GetSuccess() *bool
}

type ListAlertEventsResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The query result.
	ListResult *ListAlertEventsResponseBodyListResult `json:"ListResult,omitempty" xml:"ListResult,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAlertEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAlertEventsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAlertEventsResponseBody) GetListResult() *ListAlertEventsResponseBodyListResult {
	return s.ListResult
}

func (s *ListAlertEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlertEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAlertEventsResponseBody) SetCode(v string) *ListAlertEventsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlertEventsResponseBody) SetHttpStatusCode(v int32) *ListAlertEventsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAlertEventsResponseBody) SetListResult(v *ListAlertEventsResponseBodyListResult) *ListAlertEventsResponseBody {
	s.ListResult = v
	return s
}

func (s *ListAlertEventsResponseBody) SetMessage(v string) *ListAlertEventsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlertEventsResponseBody) SetRequestId(v string) *ListAlertEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertEventsResponseBody) SetSuccess(v bool) *ListAlertEventsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAlertEventsResponseBody) Validate() error {
	if s.ListResult != nil {
		if err := s.ListResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAlertEventsResponseBodyListResult struct {
	// The alert event query results.
	Data []*ListAlertEventsResponseBodyListResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlertEventsResponseBodyListResult) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResult) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResult) GetData() []*ListAlertEventsResponseBodyListResultData {
	return s.Data
}

func (s *ListAlertEventsResponseBodyListResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAlertEventsResponseBodyListResult) SetData(v []*ListAlertEventsResponseBodyListResultData) *ListAlertEventsResponseBodyListResult {
	s.Data = v
	return s
}

func (s *ListAlertEventsResponseBodyListResult) SetTotalCount(v int32) *ListAlertEventsResponseBodyListResult {
	s.TotalCount = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAlertEventsResponseBodyListResultData struct {
	// The alert frequency. Valid values:
	//
	// - ONCE: instant alert
	//
	// - PERIOD: periodic alert. Format: 1HOUR, 1MINUTE, 1SECOND.
	//
	// example:
	//
	// ONCE
	AlertFrequency *string `json:"AlertFrequency,omitempty" xml:"AlertFrequency,omitempty"`
	// The alert object.
	AlertObject *ListAlertEventsResponseBodyListResultDataAlertObject `json:"AlertObject,omitempty" xml:"AlertObject,omitempty" type:"Struct"`
	// The alert reason.
	AlertReason *ListAlertEventsResponseBodyListResultDataAlertReason `json:"AlertReason,omitempty" xml:"AlertReason,omitempty" type:"Struct"`
	// The list of alert receivers.
	AlertReceiverList []*ListAlertEventsResponseBodyListResultDataAlertReceiverList `json:"AlertReceiverList,omitempty" xml:"AlertReceiverList,omitempty" type:"Repeated"`
	// The project to which the alert belongs.
	BelongProject *ListAlertEventsResponseBodyListResultDataBelongProject `json:"BelongProject,omitempty" xml:"BelongProject,omitempty" type:"Struct"`
	// The expiration time of the do-not-disturb period.
	//
	// example:
	//
	// 2024-11-25 00:00:00
	DoNotDisturbEndTime *string `json:"DoNotDisturbEndTime,omitempty" xml:"DoNotDisturbEndTime,omitempty"`
	// The time of the first alert.
	//
	// example:
	//
	// 2024-11-25 10:02:47
	FirstAlertTime *string `json:"FirstAlertTime,omitempty" xml:"FirstAlertTime,omitempty"`
	// The alert event ID.
	//
	// example:
	//
	// 12345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time of the latest alert.
	//
	// example:
	//
	// 2024-11-25 10:02:47
	LatestAlertTime *string `json:"LatestAlertTime,omitempty" xml:"LatestAlertTime,omitempty"`
	// The alert status. Valid values:
	//
	// - ALERTING: alerting
	//
	// - DO_NOT_DISTURB: do not disturb
	//
	// - SILENCING: alerting (cool-down period)
	//
	// - FINISH: completed.
	//
	// example:
	//
	// FINISH
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of alert occurrences.
	//
	// example:
	//
	// 1
	TotalAlertTimes *int64 `json:"TotalAlertTimes,omitempty" xml:"TotalAlertTimes,omitempty"`
	// The URL information.
	UrlConfig *ListAlertEventsResponseBodyListResultDataUrlConfig `json:"UrlConfig,omitempty" xml:"UrlConfig,omitempty" type:"Struct"`
}

func (s ListAlertEventsResponseBodyListResultData) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultData) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultData) GetAlertFrequency() *string {
	return s.AlertFrequency
}

func (s *ListAlertEventsResponseBodyListResultData) GetAlertObject() *ListAlertEventsResponseBodyListResultDataAlertObject {
	return s.AlertObject
}

func (s *ListAlertEventsResponseBodyListResultData) GetAlertReason() *ListAlertEventsResponseBodyListResultDataAlertReason {
	return s.AlertReason
}

func (s *ListAlertEventsResponseBodyListResultData) GetAlertReceiverList() []*ListAlertEventsResponseBodyListResultDataAlertReceiverList {
	return s.AlertReceiverList
}

func (s *ListAlertEventsResponseBodyListResultData) GetBelongProject() *ListAlertEventsResponseBodyListResultDataBelongProject {
	return s.BelongProject
}

func (s *ListAlertEventsResponseBodyListResultData) GetDoNotDisturbEndTime() *string {
	return s.DoNotDisturbEndTime
}

func (s *ListAlertEventsResponseBodyListResultData) GetFirstAlertTime() *string {
	return s.FirstAlertTime
}

func (s *ListAlertEventsResponseBodyListResultData) GetId() *string {
	return s.Id
}

func (s *ListAlertEventsResponseBodyListResultData) GetLatestAlertTime() *string {
	return s.LatestAlertTime
}

func (s *ListAlertEventsResponseBodyListResultData) GetStatus() *string {
	return s.Status
}

func (s *ListAlertEventsResponseBodyListResultData) GetTotalAlertTimes() *int64 {
	return s.TotalAlertTimes
}

func (s *ListAlertEventsResponseBodyListResultData) GetUrlConfig() *ListAlertEventsResponseBodyListResultDataUrlConfig {
	return s.UrlConfig
}

func (s *ListAlertEventsResponseBodyListResultData) SetAlertFrequency(v string) *ListAlertEventsResponseBodyListResultData {
	s.AlertFrequency = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetAlertObject(v *ListAlertEventsResponseBodyListResultDataAlertObject) *ListAlertEventsResponseBodyListResultData {
	s.AlertObject = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetAlertReason(v *ListAlertEventsResponseBodyListResultDataAlertReason) *ListAlertEventsResponseBodyListResultData {
	s.AlertReason = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetAlertReceiverList(v []*ListAlertEventsResponseBodyListResultDataAlertReceiverList) *ListAlertEventsResponseBodyListResultData {
	s.AlertReceiverList = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetBelongProject(v *ListAlertEventsResponseBodyListResultDataBelongProject) *ListAlertEventsResponseBodyListResultData {
	s.BelongProject = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetDoNotDisturbEndTime(v string) *ListAlertEventsResponseBodyListResultData {
	s.DoNotDisturbEndTime = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetFirstAlertTime(v string) *ListAlertEventsResponseBodyListResultData {
	s.FirstAlertTime = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetId(v string) *ListAlertEventsResponseBodyListResultData {
	s.Id = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetLatestAlertTime(v string) *ListAlertEventsResponseBodyListResultData {
	s.LatestAlertTime = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetStatus(v string) *ListAlertEventsResponseBodyListResultData {
	s.Status = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetTotalAlertTimes(v int64) *ListAlertEventsResponseBodyListResultData {
	s.TotalAlertTimes = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) SetUrlConfig(v *ListAlertEventsResponseBodyListResultDataUrlConfig) *ListAlertEventsResponseBodyListResultData {
	s.UrlConfig = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultData) Validate() error {
	if s.AlertObject != nil {
		if err := s.AlertObject.Validate(); err != nil {
			return err
		}
	}
	if s.AlertReason != nil {
		if err := s.AlertReason.Validate(); err != nil {
			return err
		}
	}
	if s.AlertReceiverList != nil {
		for _, item := range s.AlertReceiverList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.BelongProject != nil {
		if err := s.BelongProject.Validate(); err != nil {
			return err
		}
	}
	if s.UrlConfig != nil {
		if err := s.UrlConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAlertEventsResponseBodyListResultDataAlertObject struct {
	// The name of the alert object.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source system. Valid values:
	//
	// - ALL: all
	//
	// - DQE: data quality
	//
	// - OS: data service
	//
	// - STREAM: real-time computing
	//
	// - VDM_BATCH: offline computing
	//
	// - SOP: O&M platform
	//
	// - REAL_TIME_PIPELINE: real-time integration
	//
	// - KGB: baseline monitoring
	//
	// and more.
	//
	// example:
	//
	// VDM_BATCH
	SourceSystemType *string `json:"SourceSystemType,omitempty" xml:"SourceSystemType,omitempty"`
	// The alert object type. Valid values:
	//
	// - OS_API: API operation
	//
	// - OS_APPLICATION_SERVICE: service application
	//
	// - STREAM_TASK: real-time computing
	//
	// - REAL_TIME_PIPELINE_TASK: real-time integration
	//
	// - VDM_BATCH_SHELL: SHELL
	//
	// - VDM_BATCH_PYTHON: PYTHON
	//
	// - VDM_BATCH_DATAX: DATAX
	//
	// - VDM_BATCH_DLINK: DLINK
	//
	// - VDM_BATCH_VIRTUAL: VIRTUAL
	//
	// - VDM_BATCH_PYTHON37: PYTHON37
	//
	// - VDM_BATCH_PYTHON311: PYTHON311
	//
	// - VDM_BATCH_MAX_COMPUTE_SQL: MAXCOMPUTE_SQL
	//
	// - VDM_BATCH_MAX_COMPUTE_MR: MAXCOMPUTE_MR
	//
	// - VDM_BATCH_SPARK_JAR_ON_MAX_COMPUTE: SPARK_JAR_ON_MAX_COMPUTE
	//
	// - VDM_BATCH_HIVE_SQL: HIVE_SQL
	//
	// - VDM_BATCH_HADOOP_MR: HADOOP_MR
	//
	// - VDM_BATCH_SPARK_JAR_ON_HIVE: SPARK_JAR_ON_HIVE
	//
	// - VDM_BATCH_SPARK_SQL_ON_HIVE: SPARK_SQL_ON_HIVE
	//
	// - VDM_BATCH_SPARK_SQL: VDM_BATCH_SPARK_SQL
	//
	// - DQE_LOGICAL_TABLE: logical table
	//
	// - DQE_PHYSICAL_TABLE: physical table
	//
	// - DQE_REALTIME_TABLE: real-time metadata table
	//
	// - DQE_DATA_SOURCE: data source
	//
	// - DQE_INDEX: metric
	//
	// - QD_DECISION_INVOKE: QD decision invocation
	//
	// - BASELINE: baseline
	//
	// and more.
	//
	// example:
	//
	// STREAM_TASK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAlertEventsResponseBodyListResultDataAlertObject) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataAlertObject) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) GetName() *string {
	return s.Name
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) GetSourceSystemType() *string {
	return s.SourceSystemType
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) GetType() *string {
	return s.Type
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) SetName(v string) *ListAlertEventsResponseBodyListResultDataAlertObject {
	s.Name = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) SetSourceSystemType(v string) *ListAlertEventsResponseBodyListResultDataAlertObject {
	s.SourceSystemType = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) SetType(v string) *ListAlertEventsResponseBodyListResultDataAlertObject {
	s.Type = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertObject) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResultDataAlertReason struct {
	// The list of alert reason parameters.
	AlertReasonParamList []*ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList `json:"AlertReasonParamList,omitempty" xml:"AlertReasonParamList,omitempty" type:"Repeated"`
	// The business date.
	//
	// example:
	//
	// 2024-11-25 10:02:47
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The alert reason type. Valid values:
	//
	// - DQE_COLUMN: field rule exception
	//
	// - DQE_DATA_SOURCE: data source rule exception
	//
	// - DQE_CUSTOMIZE: custom rule exception
	//
	// - DQE_TABLE: table rule exception
	//
	// - DQE_REALTIME_TABLE: real-time table rule exception
	//
	// - DQE_INDEX: metric rule exception
	//
	// - OS_AVG_RESPONSE: average response time exception
	//
	// - OS_CALL_TIMES: call count exception
	//
	// - OS_ERROR_RATE: error rate exception
	//
	// - OS_OFFLINE: Offline percentage exception
	//
	// - STREAM_BIZ_DELAY: business delay too high
	//
	// - STREAM_DATA_RETENTION: data retention exceeds configuration
	//
	// - STREAM_MORE_THAN_FAILURE: failure frequency exceeds configuration
	//
	// - STREAM_TPS_OUT_RANGE: TPS out of range
	//
	// - STREAM_CHECKPOINT_FAILURE: checkpoint failures exceed configuration
	//
	// - STREAM_BACKPRESSURE: backpressure duration exceeds configuration
	//
	// - STREAM_JOB_FAILURE: job failure
	//
	// - VDM_BATCH_ERROR: error
	//
	// - VDM_BATCH_FINISH: completed
	//
	// - VDM_BATCH_TIME_OUT: execution timeout
	//
	// - VDM_BATCH_UNDONE: incomplete
	//
	// - VDM_BATCH_LOGIC_DATA_DELAY: data delay
	//
	// - QD_DECISION_CALL_TIMES: decision call count exception
	//
	// - QD_DECISION_MAX_RESPONSE: maximum response time exception
	//
	// - QD_DECISION_ERROR_RATE: error rate exception
	//
	// - QD_DECISION_PARAM_COUNT: decision parameter count exception
	//
	// - QD_DECISION_PARAM_PERCENTAGE: decision parameter percentage exception
	//
	// - QD_DECISION_PARAM_SUM: decision parameter sum exception
	//
	// - QD_DECISION_PARAM_AVG: decision parameter average exception
	//
	// - LOGICAL_INSTANCE_GENERATION: logical instance generation monitoring
	//
	// - KGB_TASK_ERROR: baseline task error
	//
	// - KGB_TASK_SLOW_DOWN: baseline task slowdown
	//
	// - KGB_EARLY_WARNING: baseline early warning
	//
	// - KGB_BROKEN_LINE: baseline breach
	//
	// and more.
	//
	// example:
	//
	// VDM_BATCH_FINISH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The unique identifier.
	//
	// example:
	//
	// t_6340134343289405440_20241124_639873707610
	UniqueKey *string `json:"UniqueKey,omitempty" xml:"UniqueKey,omitempty"`
}

func (s ListAlertEventsResponseBodyListResultDataAlertReason) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataAlertReason) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) GetAlertReasonParamList() []*ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList {
	return s.AlertReasonParamList
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) GetBizDate() *string {
	return s.BizDate
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) GetType() *string {
	return s.Type
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) GetUniqueKey() *string {
	return s.UniqueKey
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) SetAlertReasonParamList(v []*ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) *ListAlertEventsResponseBodyListResultDataAlertReason {
	s.AlertReasonParamList = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) SetBizDate(v string) *ListAlertEventsResponseBodyListResultDataAlertReason {
	s.BizDate = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) SetType(v string) *ListAlertEventsResponseBodyListResultDataAlertReason {
	s.Type = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) SetUniqueKey(v string) *ListAlertEventsResponseBodyListResultDataAlertReason {
	s.UniqueKey = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReason) Validate() error {
	if s.AlertReasonParamList != nil {
		for _, item := range s.AlertReasonParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList struct {
	// The name of the alert reason parameter.
	//
	// example:
	//
	// biz_date
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the alert reason parameter.
	//
	// example:
	//
	// 2024-11-24 00:00:00
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) GetKey() *string {
	return s.Key
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) GetValue() *string {
	return s.Value
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) SetKey(v string) *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList {
	s.Key = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) SetValue(v string) *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList {
	s.Value = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReasonAlertReasonParamList) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResultDataAlertReceiverList struct {
	// The list of alert channel types.
	AlertChannelTypeList []*string `json:"AlertChannelTypeList,omitempty" xml:"AlertChannelTypeList,omitempty" type:"Repeated"`
	// The list of custom alert channel IDs.
	CustomAlertChannelIdList []*string `json:"CustomAlertChannelIdList,omitempty" xml:"CustomAlertChannelIdList,omitempty" type:"Repeated"`
	// The name of the on-call schedule.
	//
	// example:
	//
	// test
	OnCallTableName *string `json:"OnCallTableName,omitempty" xml:"OnCallTableName,omitempty"`
	// The type of the alert receiver. Valid values:
	//
	// - ON_CALL_TABLE: on-call schedule
	//
	// - USER_DEFINED: custom user
	//
	// - OWNER: owner.
	//
	// example:
	//
	// OWNER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The list of alert users.
	UserList []*ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s ListAlertEventsResponseBodyListResultDataAlertReceiverList) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataAlertReceiverList) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) GetAlertChannelTypeList() []*string {
	return s.AlertChannelTypeList
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) GetCustomAlertChannelIdList() []*string {
	return s.CustomAlertChannelIdList
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) GetOnCallTableName() *string {
	return s.OnCallTableName
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) GetType() *string {
	return s.Type
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) GetUserList() []*ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList {
	return s.UserList
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) SetAlertChannelTypeList(v []*string) *ListAlertEventsResponseBodyListResultDataAlertReceiverList {
	s.AlertChannelTypeList = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) SetCustomAlertChannelIdList(v []*string) *ListAlertEventsResponseBodyListResultDataAlertReceiverList {
	s.CustomAlertChannelIdList = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) SetOnCallTableName(v string) *ListAlertEventsResponseBodyListResultDataAlertReceiverList {
	s.OnCallTableName = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) SetType(v string) *ListAlertEventsResponseBodyListResultDataAlertReceiverList {
	s.Type = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) SetUserList(v []*ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList) *ListAlertEventsResponseBodyListResultDataAlertReceiverList {
	s.UserList = v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverList) Validate() error {
	if s.UserList != nil {
		for _, item := range s.UserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList struct {
	// The username.
	//
	// example:
	//
	// ADMIN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList) GetName() *string {
	return s.Name
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList) SetName(v string) *ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList {
	s.Name = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataAlertReceiverListUserList) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResultDataBelongProject struct {
	// The business unit name.
	//
	// example:
	//
	// biz_1
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The project name.
	//
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListAlertEventsResponseBodyListResultDataBelongProject) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataBelongProject) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataBelongProject) GetBizName() *string {
	return s.BizName
}

func (s *ListAlertEventsResponseBodyListResultDataBelongProject) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListAlertEventsResponseBodyListResultDataBelongProject) SetBizName(v string) *ListAlertEventsResponseBodyListResultDataBelongProject {
	s.BizName = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataBelongProject) SetProjectName(v string) *ListAlertEventsResponseBodyListResultDataBelongProject {
	s.ProjectName = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataBelongProject) Validate() error {
	return dara.Validate(s)
}

type ListAlertEventsResponseBodyListResultDataUrlConfig struct {
	// The URL of the alert configuration page.
	//
	// example:
	//
	// https://dataphin.com/ops/test3
	AlertConfigUrl *string `json:"AlertConfigUrl,omitempty" xml:"AlertConfigUrl,omitempty"`
	// The URL of the log page.
	//
	// example:
	//
	// https://dataphin.com/ops/test2
	LogUrl *string `json:"LogUrl,omitempty" xml:"LogUrl,omitempty"`
	// The URL of the alert object page.
	//
	// example:
	//
	// https://dataphin.com/ops/test1
	ObjectUrl *string `json:"ObjectUrl,omitempty" xml:"ObjectUrl,omitempty"`
}

func (s ListAlertEventsResponseBodyListResultDataUrlConfig) String() string {
	return dara.Prettify(s)
}

func (s ListAlertEventsResponseBodyListResultDataUrlConfig) GoString() string {
	return s.String()
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) GetAlertConfigUrl() *string {
	return s.AlertConfigUrl
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) GetLogUrl() *string {
	return s.LogUrl
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) GetObjectUrl() *string {
	return s.ObjectUrl
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) SetAlertConfigUrl(v string) *ListAlertEventsResponseBodyListResultDataUrlConfig {
	s.AlertConfigUrl = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) SetLogUrl(v string) *ListAlertEventsResponseBodyListResultDataUrlConfig {
	s.LogUrl = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) SetObjectUrl(v string) *ListAlertEventsResponseBodyListResultDataUrlConfig {
	s.ObjectUrl = &v
	return s
}

func (s *ListAlertEventsResponseBodyListResultDataUrlConfig) Validate() error {
	return dara.Validate(s)
}
