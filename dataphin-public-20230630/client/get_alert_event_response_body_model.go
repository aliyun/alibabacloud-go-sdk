// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertEventInfo(v *GetAlertEventResponseBodyAlertEventInfo) *GetAlertEventResponseBody
	GetAlertEventInfo() *GetAlertEventResponseBodyAlertEventInfo
	SetCode(v string) *GetAlertEventResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetAlertEventResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAlertEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAlertEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAlertEventResponseBody
	GetSuccess() *bool
}

type GetAlertEventResponseBody struct {
	// The alert event information.
	AlertEventInfo *GetAlertEventResponseBodyAlertEventInfo `json:"AlertEventInfo,omitempty" xml:"AlertEventInfo,omitempty" type:"Struct"`
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

func (s GetAlertEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBody) GetAlertEventInfo() *GetAlertEventResponseBodyAlertEventInfo {
	return s.AlertEventInfo
}

func (s *GetAlertEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAlertEventResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAlertEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAlertEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlertEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAlertEventResponseBody) SetAlertEventInfo(v *GetAlertEventResponseBodyAlertEventInfo) *GetAlertEventResponseBody {
	s.AlertEventInfo = v
	return s
}

func (s *GetAlertEventResponseBody) SetCode(v string) *GetAlertEventResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlertEventResponseBody) SetHttpStatusCode(v int32) *GetAlertEventResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAlertEventResponseBody) SetMessage(v string) *GetAlertEventResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlertEventResponseBody) SetRequestId(v string) *GetAlertEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlertEventResponseBody) SetSuccess(v bool) *GetAlertEventResponseBody {
	s.Success = &v
	return s
}

func (s *GetAlertEventResponseBody) Validate() error {
	if s.AlertEventInfo != nil {
		if err := s.AlertEventInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertEventResponseBodyAlertEventInfo struct {
	// The alert frequency. Valid values:
	//
	// - ONCE: Instant alert.
	//
	// - PERIOD: Periodic alert. Format: 1HOUR, 1MINUTE, 1SECOND.
	//
	// example:
	//
	// ONCE
	AlertFrequency *string `json:"AlertFrequency,omitempty" xml:"AlertFrequency,omitempty"`
	// The alert object.
	AlertObject *GetAlertEventResponseBodyAlertEventInfoAlertObject `json:"AlertObject,omitempty" xml:"AlertObject,omitempty" type:"Struct"`
	// The alert reason.
	AlertReason *GetAlertEventResponseBodyAlertEventInfoAlertReason `json:"AlertReason,omitempty" xml:"AlertReason,omitempty" type:"Struct"`
	// The list of alert receivers.
	AlertReceiverList []*GetAlertEventResponseBodyAlertEventInfoAlertReceiverList `json:"AlertReceiverList,omitempty" xml:"AlertReceiverList,omitempty" type:"Repeated"`
	// The project to which the alert event belongs.
	BelongProject *GetAlertEventResponseBodyAlertEventInfoBelongProject `json:"BelongProject,omitempty" xml:"BelongProject,omitempty" type:"Struct"`
	// The expiration time of the do-not-disturb period.
	//
	// example:
	//
	// 2024-11-05 00:00:00
	DoNotDisturbEndTime *string `json:"DoNotDisturbEndTime,omitempty" xml:"DoNotDisturbEndTime,omitempty"`
	// The time of the first alert.
	//
	// example:
	//
	// 2024-11-05 16:19:33
	FirstAlertTime *string `json:"FirstAlertTime,omitempty" xml:"FirstAlertTime,omitempty"`
	// The alert event ID.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time of the latest alert.
	//
	// example:
	//
	// 2024-11-05 16:19:33
	LatestAlertTime *string `json:"LatestAlertTime,omitempty" xml:"LatestAlertTime,omitempty"`
	// The alert status. Valid values:
	//
	// - ALERTING: Alerting.
	//
	// - DO_NOT_DISTURB: Do not disturb.
	//
	// - SILENCING: Alerting (cool-down period).
	//
	// - FINISH: Alert completed.
	//
	// example:
	//
	// FINISH
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of alerts.
	//
	// example:
	//
	// 1
	TotalAlertTimes *int64 `json:"TotalAlertTimes,omitempty" xml:"TotalAlertTimes,omitempty"`
	// The URL configuration.
	UrlConfig *GetAlertEventResponseBodyAlertEventInfoUrlConfig `json:"UrlConfig,omitempty" xml:"UrlConfig,omitempty" type:"Struct"`
}

func (s GetAlertEventResponseBodyAlertEventInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfo) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetAlertFrequency() *string {
	return s.AlertFrequency
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetAlertObject() *GetAlertEventResponseBodyAlertEventInfoAlertObject {
	return s.AlertObject
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetAlertReason() *GetAlertEventResponseBodyAlertEventInfoAlertReason {
	return s.AlertReason
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetAlertReceiverList() []*GetAlertEventResponseBodyAlertEventInfoAlertReceiverList {
	return s.AlertReceiverList
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetBelongProject() *GetAlertEventResponseBodyAlertEventInfoBelongProject {
	return s.BelongProject
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetDoNotDisturbEndTime() *string {
	return s.DoNotDisturbEndTime
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetFirstAlertTime() *string {
	return s.FirstAlertTime
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetId() *int64 {
	return s.Id
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetLatestAlertTime() *string {
	return s.LatestAlertTime
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetStatus() *string {
	return s.Status
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetTotalAlertTimes() *int64 {
	return s.TotalAlertTimes
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetUrlConfig() *GetAlertEventResponseBodyAlertEventInfoUrlConfig {
	return s.UrlConfig
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetAlertFrequency(v string) *GetAlertEventResponseBodyAlertEventInfo {
	s.AlertFrequency = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetAlertObject(v *GetAlertEventResponseBodyAlertEventInfoAlertObject) *GetAlertEventResponseBodyAlertEventInfo {
	s.AlertObject = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetAlertReason(v *GetAlertEventResponseBodyAlertEventInfoAlertReason) *GetAlertEventResponseBodyAlertEventInfo {
	s.AlertReason = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetAlertReceiverList(v []*GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) *GetAlertEventResponseBodyAlertEventInfo {
	s.AlertReceiverList = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetBelongProject(v *GetAlertEventResponseBodyAlertEventInfoBelongProject) *GetAlertEventResponseBodyAlertEventInfo {
	s.BelongProject = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetDoNotDisturbEndTime(v string) *GetAlertEventResponseBodyAlertEventInfo {
	s.DoNotDisturbEndTime = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetFirstAlertTime(v string) *GetAlertEventResponseBodyAlertEventInfo {
	s.FirstAlertTime = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetId(v int64) *GetAlertEventResponseBodyAlertEventInfo {
	s.Id = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetLatestAlertTime(v string) *GetAlertEventResponseBodyAlertEventInfo {
	s.LatestAlertTime = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetStatus(v string) *GetAlertEventResponseBodyAlertEventInfo {
	s.Status = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetTotalAlertTimes(v int64) *GetAlertEventResponseBodyAlertEventInfo {
	s.TotalAlertTimes = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetUrlConfig(v *GetAlertEventResponseBodyAlertEventInfoUrlConfig) *GetAlertEventResponseBodyAlertEventInfo {
	s.UrlConfig = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) Validate() error {
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

type GetAlertEventResponseBodyAlertEventInfoAlertObject struct {
	// The name of the alert object.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source system type. Valid values:
	//
	// - ALL: all systems
	//
	// - DQE: data quality
	//
	// - OS: data service
	//
	// - STREAM: real-time computing
	//
	// - VDM_BATCH: batch computing
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
	// ALL
	SourceSystemType *string `json:"SourceSystemType,omitempty" xml:"SourceSystemType,omitempty"`
	// The alerting object type. Valid values:
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
	// VDM_BATCH_PYTHON37
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertObject) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertObject) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) GetName() *string {
	return s.Name
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) GetSourceSystemType() *string {
	return s.SourceSystemType
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) GetType() *string {
	return s.Type
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) SetName(v string) *GetAlertEventResponseBodyAlertEventInfoAlertObject {
	s.Name = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) SetSourceSystemType(v string) *GetAlertEventResponseBodyAlertEventInfoAlertObject {
	s.SourceSystemType = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) SetType(v string) *GetAlertEventResponseBodyAlertEventInfoAlertObject {
	s.Type = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) Validate() error {
	return dara.Validate(s)
}

type GetAlertEventResponseBodyAlertEventInfoAlertReason struct {
	// The list of alert reason parameters.
	AlertReasonParamList []*GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList `json:"AlertReasonParamList,omitempty" xml:"AlertReasonParamList,omitempty" type:"Repeated"`
	// The business date.
	//
	// example:
	//
	// 2024-11-05 16:19:32
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The type of the alert reason. Valid values:
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
	// t_6340131422711644160_20241104_6340142
	UniqueKey *string `json:"UniqueKey,omitempty" xml:"UniqueKey,omitempty"`
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReason) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReason) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) GetAlertReasonParamList() []*GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList {
	return s.AlertReasonParamList
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) GetBizDate() *string {
	return s.BizDate
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) GetType() *string {
	return s.Type
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) GetUniqueKey() *string {
	return s.UniqueKey
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) SetAlertReasonParamList(v []*GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) *GetAlertEventResponseBodyAlertEventInfoAlertReason {
	s.AlertReasonParamList = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) SetBizDate(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReason {
	s.BizDate = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) SetType(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReason {
	s.Type = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) SetUniqueKey(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReason {
	s.UniqueKey = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) Validate() error {
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

type GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList struct {
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
	// 2024-11-04 00:00:00
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) GetKey() *string {
	return s.Key
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) GetValue() *string {
	return s.Value
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) SetKey(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList {
	s.Key = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) SetValue(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList {
	s.Value = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) Validate() error {
	return dara.Validate(s)
}

type GetAlertEventResponseBodyAlertEventInfoAlertReceiverList struct {
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
	UserList []*GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) GetAlertChannelTypeList() []*string {
	return s.AlertChannelTypeList
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) GetCustomAlertChannelIdList() []*string {
	return s.CustomAlertChannelIdList
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) GetOnCallTableName() *string {
	return s.OnCallTableName
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) GetType() *string {
	return s.Type
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) GetUserList() []*GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList {
	return s.UserList
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) SetAlertChannelTypeList(v []*string) *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList {
	s.AlertChannelTypeList = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) SetCustomAlertChannelIdList(v []*string) *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList {
	s.CustomAlertChannelIdList = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) SetOnCallTableName(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList {
	s.OnCallTableName = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) SetType(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList {
	s.Type = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) SetUserList(v []*GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList) *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList {
	s.UserList = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) Validate() error {
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

type GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList struct {
	// The username.
	//
	// example:
	//
	// Admin
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList) GetName() *string {
	return s.Name
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList) SetName(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList {
	s.Name = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList) Validate() error {
	return dara.Validate(s)
}

type GetAlertEventResponseBodyAlertEventInfoBelongProject struct {
	// The name of the business unit.
	//
	// example:
	//
	// biz_1
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetAlertEventResponseBodyAlertEventInfoBelongProject) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoBelongProject) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoBelongProject) GetBizName() *string {
	return s.BizName
}

func (s *GetAlertEventResponseBodyAlertEventInfoBelongProject) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetAlertEventResponseBodyAlertEventInfoBelongProject) SetBizName(v string) *GetAlertEventResponseBodyAlertEventInfoBelongProject {
	s.BizName = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoBelongProject) SetProjectName(v string) *GetAlertEventResponseBodyAlertEventInfoBelongProject {
	s.ProjectName = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoBelongProject) Validate() error {
	return dara.Validate(s)
}

type GetAlertEventResponseBodyAlertEventInfoUrlConfig struct {
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

func (s GetAlertEventResponseBodyAlertEventInfoUrlConfig) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoUrlConfig) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) GetAlertConfigUrl() *string {
	return s.AlertConfigUrl
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) GetLogUrl() *string {
	return s.LogUrl
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) GetObjectUrl() *string {
	return s.ObjectUrl
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) SetAlertConfigUrl(v string) *GetAlertEventResponseBodyAlertEventInfoUrlConfig {
	s.AlertConfigUrl = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) SetLogUrl(v string) *GetAlertEventResponseBodyAlertEventInfoUrlConfig {
	s.LogUrl = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) SetObjectUrl(v string) *GetAlertEventResponseBodyAlertEventInfoUrlConfig {
	s.ObjectUrl = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) Validate() error {
	return dara.Validate(s)
}
