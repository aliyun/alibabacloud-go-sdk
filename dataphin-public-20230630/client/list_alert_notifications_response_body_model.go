// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertNotificationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAlertNotificationsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListAlertNotificationsResponseBody
	GetHttpStatusCode() *int32
	SetListResult(v *ListAlertNotificationsResponseBodyListResult) *ListAlertNotificationsResponseBody
	GetListResult() *ListAlertNotificationsResponseBodyListResult
	SetMessage(v string) *ListAlertNotificationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAlertNotificationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAlertNotificationsResponseBody
	GetSuccess() *bool
}

type ListAlertNotificationsResponseBody struct {
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
	ListResult *ListAlertNotificationsResponseBodyListResult `json:"ListResult,omitempty" xml:"ListResult,omitempty" type:"Struct"`
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

func (s ListAlertNotificationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAlertNotificationsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAlertNotificationsResponseBody) GetListResult() *ListAlertNotificationsResponseBodyListResult {
	return s.ListResult
}

func (s *ListAlertNotificationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlertNotificationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertNotificationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAlertNotificationsResponseBody) SetCode(v string) *ListAlertNotificationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlertNotificationsResponseBody) SetHttpStatusCode(v int32) *ListAlertNotificationsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAlertNotificationsResponseBody) SetListResult(v *ListAlertNotificationsResponseBodyListResult) *ListAlertNotificationsResponseBody {
	s.ListResult = v
	return s
}

func (s *ListAlertNotificationsResponseBody) SetMessage(v string) *ListAlertNotificationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlertNotificationsResponseBody) SetRequestId(v string) *ListAlertNotificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertNotificationsResponseBody) SetSuccess(v bool) *ListAlertNotificationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAlertNotificationsResponseBody) Validate() error {
	if s.ListResult != nil {
		if err := s.ListResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAlertNotificationsResponseBodyListResult struct {
	// The list of push records.
	Data []*ListAlertNotificationsResponseBodyListResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlertNotificationsResponseBodyListResult) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsResponseBodyListResult) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsResponseBodyListResult) GetData() []*ListAlertNotificationsResponseBodyListResultData {
	return s.Data
}

func (s *ListAlertNotificationsResponseBodyListResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAlertNotificationsResponseBodyListResult) SetData(v []*ListAlertNotificationsResponseBodyListResultData) *ListAlertNotificationsResponseBodyListResult {
	s.Data = v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResult) SetTotalCount(v int32) *ListAlertNotificationsResponseBodyListResult {
	s.TotalCount = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResult) Validate() error {
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

type ListAlertNotificationsResponseBodyListResultData struct {
	// The alert event ID.
	//
	// example:
	//
	// 12345
	AlertEventId *string `json:"AlertEventId,omitempty" xml:"AlertEventId,omitempty"`
	// The alert object.
	AlertObject *ListAlertNotificationsResponseBodyListResultDataAlertObject `json:"AlertObject,omitempty" xml:"AlertObject,omitempty" type:"Struct"`
	// The alert reason.
	AlertReason *ListAlertNotificationsResponseBodyListResultDataAlertReason `json:"AlertReason,omitempty" xml:"AlertReason,omitempty" type:"Struct"`
	// The receiver information.
	AlertReceiver *ListAlertNotificationsResponseBodyListResultDataAlertReceiver `json:"AlertReceiver,omitempty" xml:"AlertReceiver,omitempty" type:"Struct"`
	// The alert sending information.
	AlertSend *ListAlertNotificationsResponseBodyListResultDataAlertSend `json:"AlertSend,omitempty" xml:"AlertSend,omitempty" type:"Struct"`
}

func (s ListAlertNotificationsResponseBodyListResultData) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsResponseBodyListResultData) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsResponseBodyListResultData) GetAlertEventId() *string {
	return s.AlertEventId
}

func (s *ListAlertNotificationsResponseBodyListResultData) GetAlertObject() *ListAlertNotificationsResponseBodyListResultDataAlertObject {
	return s.AlertObject
}

func (s *ListAlertNotificationsResponseBodyListResultData) GetAlertReason() *ListAlertNotificationsResponseBodyListResultDataAlertReason {
	return s.AlertReason
}

func (s *ListAlertNotificationsResponseBodyListResultData) GetAlertReceiver() *ListAlertNotificationsResponseBodyListResultDataAlertReceiver {
	return s.AlertReceiver
}

func (s *ListAlertNotificationsResponseBodyListResultData) GetAlertSend() *ListAlertNotificationsResponseBodyListResultDataAlertSend {
	return s.AlertSend
}

func (s *ListAlertNotificationsResponseBodyListResultData) SetAlertEventId(v string) *ListAlertNotificationsResponseBodyListResultData {
	s.AlertEventId = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultData) SetAlertObject(v *ListAlertNotificationsResponseBodyListResultDataAlertObject) *ListAlertNotificationsResponseBodyListResultData {
	s.AlertObject = v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultData) SetAlertReason(v *ListAlertNotificationsResponseBodyListResultDataAlertReason) *ListAlertNotificationsResponseBodyListResultData {
	s.AlertReason = v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultData) SetAlertReceiver(v *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) *ListAlertNotificationsResponseBodyListResultData {
	s.AlertReceiver = v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultData) SetAlertSend(v *ListAlertNotificationsResponseBodyListResultDataAlertSend) *ListAlertNotificationsResponseBodyListResultData {
	s.AlertSend = v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultData) Validate() error {
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
	if s.AlertReceiver != nil {
		if err := s.AlertReceiver.Validate(); err != nil {
			return err
		}
	}
	if s.AlertSend != nil {
		if err := s.AlertSend.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAlertNotificationsResponseBodyListResultDataAlertObject struct {
	// The object name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source system. Valid values:
	//
	// - ALL: all.
	//
	// - DQE: data quality.
	//
	// - OS: data service.
	//
	// - STREAM: real-time computing.
	//
	// - VDM_BATCH: offline computing.
	//
	// - SOP: O&M platform.
	//
	// - REAL_TIME_PIPELINE: real-time integration.
	//
	// - KGB: baseline monitoring.
	//
	// And more.
	//
	// example:
	//
	// VDM_BATCH
	SourceSystemType *string `json:"SourceSystemType,omitempty" xml:"SourceSystemType,omitempty"`
	// The alert object type. Valid values:
	//
	// - OS_API: API operation.
	//
	// - OS_APPLICATION_SERVICE: service application.
	//
	// - STREAM_TASK: real-time computing.
	//
	// - REAL_TIME_PIPELINE_TASK: real-time integration.
	//
	// - VDM_BATCH_SHELL: SHELL.
	//
	// - VDM_BATCH_PYTHON: PYTHON.
	//
	// - VDM_BATCH_DATAX: DATAX.
	//
	// - VDM_BATCH_DLINK: DLINK.
	//
	// - VDM_BATCH_VIRTUAL: VIRTUAL.
	//
	// - VDM_BATCH_PYTHON37: PYTHON37.
	//
	// - VDM_BATCH_PYTHON311: PYTHON311.
	//
	// - VDM_BATCH_MAX_COMPUTE_SQL: MAXCOMPUTE_SQL.
	//
	// - VDM_BATCH_MAX_COMPUTE_MR: MAXCOMPUTE_MR.
	//
	// - VDM_BATCH_SPARK_JAR_ON_MAX_COMPUTE: SPARK_JAR_ON_MAX_COMPUTE.
	//
	// - VDM_BATCH_HIVE_SQL: HIVE_SQL.
	//
	// - VDM_BATCH_HADOOP_MR: HADOOP_MR.
	//
	// - VDM_BATCH_SPARK_JAR_ON_HIVE: SPARK_JAR_ON_HIVE.
	//
	// - VDM_BATCH_SPARK_SQL_ON_HIVE: SPARK_SQL_ON_HIVE.
	//
	// - VDM_BATCH_SPARK_SQL: VDM_BATCH_SPARK_SQL.
	//
	// - DQE_LOGICAL_TABLE: logical table.
	//
	// - DQE_PHYSICAL_TABLE: physical table.
	//
	// - DQE_REALTIME_TABLE: real-time meta table.
	//
	// - DQE_DATA_SOURCE: data source.
	//
	// - DQE_INDEX: metric.
	//
	// - QD_DECISION_INVOKE: QD decision invocation.
	//
	// - BASELINE: baseline.
	//
	// And more.
	//
	// example:
	//
	// VDM_BATCH_SHELL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAlertNotificationsResponseBodyListResultDataAlertObject) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsResponseBodyListResultDataAlertObject) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertObject) GetName() *string {
	return s.Name
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertObject) GetSourceSystemType() *string {
	return s.SourceSystemType
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertObject) GetType() *string {
	return s.Type
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertObject) SetName(v string) *ListAlertNotificationsResponseBodyListResultDataAlertObject {
	s.Name = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertObject) SetSourceSystemType(v string) *ListAlertNotificationsResponseBodyListResultDataAlertObject {
	s.SourceSystemType = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertObject) SetType(v string) *ListAlertNotificationsResponseBodyListResultDataAlertObject {
	s.Type = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertObject) Validate() error {
	return dara.Validate(s)
}

type ListAlertNotificationsResponseBodyListResultDataAlertReason struct {
	// The list of alert parameters.
	AlertReasonParamList []*ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList `json:"AlertReasonParamList,omitempty" xml:"AlertReasonParamList,omitempty" type:"Repeated"`
	// The business date.
	//
	// example:
	//
	// 20241125
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The alert reason type. Valid values:
	//
	// - DQE_COLUMN: field rule exception.
	//
	// - DQE_DATA_SOURCE: data source rule exception.
	//
	// - DQE_CUSTOMIZE: custom rule exception.
	//
	// - DQE_TABLE: table rule exception.
	//
	// - DQE_REALTIME_TABLE: real-time table rule exception.
	//
	// - DQE_INDEX: metric rule exception.
	//
	// - OS_AVG_RESPONSE: average response time exception.
	//
	// - OS_CALL_TIMES: call count exception.
	//
	// - OS_ERROR_RATE: error rate exception.
	//
	// - OS_OFFLINE: Offline percentage exception.
	//
	// - STREAM_BIZ_DELAY: business delay too high.
	//
	// - STREAM_DATA_RETENTION: data retention exceeds configuration.
	//
	// - STREAM_MORE_THAN_FAILURE: failure frequency exceeds configuration.
	//
	// - STREAM_TPS_OUT_RANGE: TPS out of range.
	//
	// - STREAM_CHECKPOINT_FAILURE: checkpoint failures exceed configuration.
	//
	// - STREAM_BACKPRESSURE: backpressure duration exceeds configuration.
	//
	// - STREAM_JOB_FAILURE: job execution failed.
	//
	// - VDM_BATCH_ERROR: error.
	//
	// - VDM_BATCH_FINISH: completed.
	//
	// - VDM_BATCH_TIME_OUT: execution timed out.
	//
	// - VDM_BATCH_UNDONE: not completed.
	//
	// - VDM_BATCH_LOGIC_DATA_DELAY: data delay.
	//
	// - QD_DECISION_CALL_TIMES: decision call count exception.
	//
	// - QD_DECISION_MAX_RESPONSE: maximum response time exception.
	//
	// - QD_DECISION_ERROR_RATE: error rate exception.
	//
	// - QD_DECISION_PARAM_COUNT: decision parameter count exception.
	//
	// - QD_DECISION_PARAM_PERCENTAGE: decision parameter percentage exception.
	//
	// - QD_DECISION_PARAM_SUM: decision parameter sum exception.
	//
	// - QD_DECISION_PARAM_AVG: decision parameter average exception.
	//
	// - LOGICAL_INSTANCE_GENERATION: logical instance generation monitoring.
	//
	// - KGB_TASK_ERROR: baseline task error.
	//
	// - KGB_TASK_SLOW_DOWN: baseline task slowdown.
	//
	// - KGB_EARLY_WARNING: baseline early warning.
	//
	// - KGB_BROKEN_LINE: baseline broken line.
	//
	// And more.
	//
	// example:
	//
	// VDM_BATCH_FINISH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The unique identifier.
	//
	// example:
	//
	// 123456
	UniqueKey *string `json:"UniqueKey,omitempty" xml:"UniqueKey,omitempty"`
}

func (s ListAlertNotificationsResponseBodyListResultDataAlertReason) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsResponseBodyListResultDataAlertReason) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReason) GetAlertReasonParamList() []*ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList {
	return s.AlertReasonParamList
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReason) GetBizDate() *string {
	return s.BizDate
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReason) GetType() *string {
	return s.Type
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReason) GetUniqueKey() *string {
	return s.UniqueKey
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReason) SetAlertReasonParamList(v []*ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList) *ListAlertNotificationsResponseBodyListResultDataAlertReason {
	s.AlertReasonParamList = v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReason) SetBizDate(v string) *ListAlertNotificationsResponseBodyListResultDataAlertReason {
	s.BizDate = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReason) SetType(v string) *ListAlertNotificationsResponseBodyListResultDataAlertReason {
	s.Type = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReason) SetUniqueKey(v string) *ListAlertNotificationsResponseBodyListResultDataAlertReason {
	s.UniqueKey = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReason) Validate() error {
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

type ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList struct {
	// The alert parameter name.
	//
	// example:
	//
	// biz_date
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The alert parameter value.
	//
	// example:
	//
	// 2024-11-24 00:00:00
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList) GetKey() *string {
	return s.Key
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList) GetValue() *string {
	return s.Value
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList) SetKey(v string) *ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList {
	s.Key = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList) SetValue(v string) *ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList {
	s.Value = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReasonAlertReasonParamList) Validate() error {
	return dara.Validate(s)
}

type ListAlertNotificationsResponseBodyListResultDataAlertReceiver struct {
	// The push channel type. Valid values:
	//
	// - VOICE: phone call.
	//
	// - SMS: text message.
	//
	// - MAIL: email.
	//
	// - DINGTALK_ROBOT: DingTalk robot.
	//
	// - DINGDING: DingTalk work notification.
	//
	// - CUSTOM: custom message channel.
	//
	// - WECHAT: WeCom.
	//
	// - FEISHU: Lark.
	//
	// - SILENCE: do not send.
	//
	// example:
	//
	// SMS
	AlertChannelType *string `json:"AlertChannelType,omitempty" xml:"AlertChannelType,omitempty"`
	// The custom message channel ID.
	//
	// example:
	//
	// 123456
	CustomAlertChannelId *string `json:"CustomAlertChannelId,omitempty" xml:"CustomAlertChannelId,omitempty"`
	// The on-call schedule ID.
	//
	// example:
	//
	// 12345
	OnCallTableId *string `json:"OnCallTableId,omitempty" xml:"OnCallTableId,omitempty"`
	// The on-call schedule name.
	//
	// example:
	//
	// test
	OnCallTableName *string `json:"OnCallTableName,omitempty" xml:"OnCallTableName,omitempty"`
	// The alert receiver type. Valid values:
	//
	// - ON_CALL_TABLE: on-call schedule.
	//
	// - USER_DEFINED: custom user.
	//
	// - OWNER: owner.
	//
	// example:
	//
	// OWNER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user information.
	User *ListAlertNotificationsResponseBodyListResultDataAlertReceiverUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s ListAlertNotificationsResponseBodyListResultDataAlertReceiver) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsResponseBodyListResultDataAlertReceiver) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) GetAlertChannelType() *string {
	return s.AlertChannelType
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) GetCustomAlertChannelId() *string {
	return s.CustomAlertChannelId
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) GetOnCallTableId() *string {
	return s.OnCallTableId
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) GetOnCallTableName() *string {
	return s.OnCallTableName
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) GetType() *string {
	return s.Type
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) GetUser() *ListAlertNotificationsResponseBodyListResultDataAlertReceiverUser {
	return s.User
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) SetAlertChannelType(v string) *ListAlertNotificationsResponseBodyListResultDataAlertReceiver {
	s.AlertChannelType = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) SetCustomAlertChannelId(v string) *ListAlertNotificationsResponseBodyListResultDataAlertReceiver {
	s.CustomAlertChannelId = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) SetOnCallTableId(v string) *ListAlertNotificationsResponseBodyListResultDataAlertReceiver {
	s.OnCallTableId = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) SetOnCallTableName(v string) *ListAlertNotificationsResponseBodyListResultDataAlertReceiver {
	s.OnCallTableName = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) SetType(v string) *ListAlertNotificationsResponseBodyListResultDataAlertReceiver {
	s.Type = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) SetUser(v *ListAlertNotificationsResponseBodyListResultDataAlertReceiverUser) *ListAlertNotificationsResponseBodyListResultDataAlertReceiver {
	s.User = v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiver) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAlertNotificationsResponseBodyListResultDataAlertReceiverUser struct {
	// The name of the alert receiver.
	//
	// example:
	//
	// ADMIN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAlertNotificationsResponseBodyListResultDataAlertReceiverUser) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsResponseBodyListResultDataAlertReceiverUser) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiverUser) GetName() *string {
	return s.Name
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiverUser) SetName(v string) *ListAlertNotificationsResponseBodyListResultDataAlertReceiverUser {
	s.Name = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertReceiverUser) Validate() error {
	return dara.Validate(s)
}

type ListAlertNotificationsResponseBodyListResultDataAlertSend struct {
	// The alert reason.
	//
	// example:
	//
	// 不合法
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// The push content.
	//
	// example:
	//
	// test
	SendContent *string `json:"SendContent,omitempty" xml:"SendContent,omitempty"`
	// The push time.
	//
	// example:
	//
	// 2024-11-25 10:02:47
	SendTime *string `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	// The push status. Valid values:
	//
	// - SUCCESS: Sent successfully.
	//
	// - FAILE: Failed to send.
	//
	// - SENDING: Sending in progress.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAlertNotificationsResponseBodyListResultDataAlertSend) String() string {
	return dara.Prettify(s)
}

func (s ListAlertNotificationsResponseBodyListResultDataAlertSend) GoString() string {
	return s.String()
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertSend) GetFailReason() *string {
	return s.FailReason
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertSend) GetSendContent() *string {
	return s.SendContent
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertSend) GetSendTime() *string {
	return s.SendTime
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertSend) GetStatus() *string {
	return s.Status
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertSend) SetFailReason(v string) *ListAlertNotificationsResponseBodyListResultDataAlertSend {
	s.FailReason = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertSend) SetSendContent(v string) *ListAlertNotificationsResponseBodyListResultDataAlertSend {
	s.SendContent = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertSend) SetSendTime(v string) *ListAlertNotificationsResponseBodyListResultDataAlertSend {
	s.SendTime = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertSend) SetStatus(v string) *ListAlertNotificationsResponseBodyListResultDataAlertSend {
	s.Status = &v
	return s
}

func (s *ListAlertNotificationsResponseBodyListResultDataAlertSend) Validate() error {
	return dara.Validate(s)
}
