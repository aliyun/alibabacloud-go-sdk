// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTaskConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataCheckTaskConfigResponseBodyData) *GetDataCheckTaskConfigResponseBody
	GetData() *GetDataCheckTaskConfigResponseBodyData
	SetErrCode(v string) *GetDataCheckTaskConfigResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetDataCheckTaskConfigResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetDataCheckTaskConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCheckTaskConfigResponseBody
	GetSuccess() *bool
}

type GetDataCheckTaskConfigResponseBody struct {
	// The data body returned by the operation. For the field structure, see the descriptions of child fields.
	Data *GetDataCheckTaskConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, check errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDataCheckTaskConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTaskConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCheckTaskConfigResponseBody) GetData() *GetDataCheckTaskConfigResponseBodyData {
	return s.Data
}

func (s *GetDataCheckTaskConfigResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDataCheckTaskConfigResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetDataCheckTaskConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCheckTaskConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCheckTaskConfigResponseBody) SetData(v *GetDataCheckTaskConfigResponseBodyData) *GetDataCheckTaskConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetDataCheckTaskConfigResponseBody) SetErrCode(v string) *GetDataCheckTaskConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBody) SetErrMessage(v string) *GetDataCheckTaskConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBody) SetRequestId(v string) *GetDataCheckTaskConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBody) SetSuccess(v bool) *GetDataCheckTaskConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataCheckTaskConfigResponseBodyData struct {
	// The batch ID that uniquely identifies a data validation batch.
	//
	// example:
	//
	// 20001
	BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// The global node parameter settings (built-in configuration of the data validation service).
	//
	// example:
	//
	// {}
	CheckGlobalParams *string `json:"checkGlobalParams,omitempty" xml:"checkGlobalParams,omitempty"`
	// The validation template ID.
	//
	// example:
	//
	// 1001
	CheckTemplateId *string `json:"checkTemplateId,omitempty" xml:"checkTemplateId,omitempty"`
	// The validation rule type. Valid values:
	//
	// - 0: data volume comparison.
	//
	// - 1: metric comparison.
	//
	// - 2: weak content comparison.
	//
	// - 3: custom comparison.
	//
	// - 4: full-text comparison.
	//
	// - 5: null value ratio comparison.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"checkType,omitempty" xml:"checkType,omitempty"`
	// The batch concurrency.
	//
	// example:
	//
	// 5
	Concurrency *int32 `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
	// The scheduling cycle expression (cron expression).
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	CronExp *string `json:"cronExp,omitempty" xml:"cronExp,omitempty"`
	// The task configuration table.
	DataCheckConfig []*GetDataCheckTaskConfigResponseBodyDataDataCheckConfig `json:"dataCheckConfig,omitempty" xml:"dataCheckConfig,omitempty" type:"Repeated"`
	// The destination data source ID.
	//
	// example:
	//
	// 2001
	DstDsId *string `json:"dstDsId,omitempty" xml:"dstDsId,omitempty"`
	// The destination data source name.
	//
	// example:
	//
	// ds_demo
	DstDsName *string `json:"dstDsName,omitempty" xml:"dstDsName,omitempty"`
	// The destination data source type.
	//
	// example:
	//
	// Hive
	DstDsType *string `json:"dstDsType,omitempty" xml:"dstDsType,omitempty"`
	// The ID of the destination verification engine.
	//
	// example:
	//
	// 2001
	DstEngineId *string `json:"dstEngineId,omitempty" xml:"dstEngineId,omitempty"`
	// The name of the destination verification engine.
	//
	// example:
	//
	// engine_demo
	DstEngineName *string `json:"dstEngineName,omitempty" xml:"dstEngineName,omitempty"`
	// The type of the destination verification engine.
	//
	// example:
	//
	// Tez
	DstEngineType *string `json:"dstEngineType,omitempty" xml:"dstEngineType,omitempty"`
	// The execution type. Valid values:
	//
	// - 0: immediate execution
	//
	// - 1: scheduled execution
	//
	// example:
	//
	// 0
	ExecuteType *int32 `json:"executeType,omitempty" xml:"executeType,omitempty"`
	// The count mode. Valid values:
	//
	// - 0: count by partition
	//
	// - 1: count the entire table
	//
	// example:
	//
	// 0
	FullTableCount *int32 `json:"fullTableCount,omitempty" xml:"fullTableCount,omitempty"`
	// The group data volume comparison threshold.
	//
	// example:
	//
	// 0.5
	GroupCountThreshold *float32 `json:"groupCountThreshold,omitempty" xml:"groupCountThreshold,omitempty"`
	// Indicates whether the template is a built-in template. Valid values:
	//
	// - 0: No.
	//
	// - 1: Yes.
	//
	// example:
	//
	// 0
	IsBuiltin *int32 `json:"isBuiltin,omitempty" xml:"isBuiltin,omitempty"`
	// Indicates whether scheduling is enabled. Valid values:
	//
	// - 0: No.
	//
	// - 1: Yes.
	//
	// example:
	//
	// 0
	IsScheduled *int32 `json:"isScheduled,omitempty" xml:"isScheduled,omitempty"`
	// Indicates whether the task is on the whitelist. Valid values:
	//
	// - 0: No.
	//
	// - 1: Yes.
	//
	// example:
	//
	// 0
	IsWhiteList *int32 `json:"isWhiteList,omitempty" xml:"isWhiteList,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the scheduled task (scheduling ID).
	//
	// example:
	//
	// 1001
	ScheduleId *int64 `json:"scheduleId,omitempty" xml:"scheduleId,omitempty"`
	// The scope filter JSON data.
	ScopeFilter *GetDataCheckTaskConfigResponseBodyDataScopeFilter `json:"scopeFilter,omitempty" xml:"scopeFilter,omitempty" type:"Struct"`
	// The source node parameter settings (source execute parameters).
	//
	// example:
	//
	// {}
	SourceGlobalParams *string `json:"sourceGlobalParams,omitempty" xml:"sourceGlobalParams,omitempty"`
	// The ID of the source data source.
	//
	// example:
	//
	// 1001
	SrcDsId *string `json:"srcDsId,omitempty" xml:"srcDsId,omitempty"`
	// The name of the source data source.
	//
	// example:
	//
	// ds_demo
	SrcDsName *string `json:"srcDsName,omitempty" xml:"srcDsName,omitempty"`
	// The type of the source data source.
	//
	// example:
	//
	// Hive
	SrcDsType *string `json:"srcDsType,omitempty" xml:"srcDsType,omitempty"`
	// The ID of the source verification engine.
	//
	// example:
	//
	// 1001
	SrcEngineId *string `json:"srcEngineId,omitempty" xml:"srcEngineId,omitempty"`
	// The name of the source verification engine.
	//
	// example:
	//
	// engine_demo
	SrcEngineName *string `json:"srcEngineName,omitempty" xml:"srcEngineName,omitempty"`
	// The type of the source verification engine.
	//
	// example:
	//
	// Tez
	SrcEngineType *string `json:"srcEngineType,omitempty" xml:"srcEngineType,omitempty"`
	// Indicates whether to start the task immediately. Valid values:
	//
	// - 0: No.
	//
	// - 1: Yes.
	//
	// example:
	//
	// 0
	StartImmediately *int32 `json:"startImmediately,omitempty" xml:"startImmediately,omitempty"`
	// The destination node parameter settings (destination execute parameters).
	//
	// example:
	//
	// {}
	TargetGlobalParams *string `json:"targetGlobalParams,omitempty" xml:"targetGlobalParams,omitempty"`
	// The regular expression information of the verification task.
	//
	// example:
	//
	// lhm|lhm_dw|*
	TaskConfigInfo *string `json:"taskConfigInfo,omitempty" xml:"taskConfigInfo,omitempty"`
	// The task description.
	//
	// example:
	//
	// Data validation node description
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// The task ID, which uniquely identifies a task.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The parameter creation mode. Valid values:
	//
	// - 0: fine-grained creation on a per-table basis
	//
	// - 1: batch creation with the same pattern
	//
	// example:
	//
	// 0
	TaskMode *int32 `json:"taskMode,omitempty" xml:"taskMode,omitempty"`
	// The task name. When used as a query condition, fuzzy matching with % is supported (SQL syntax).
	//
	// example:
	//
	// data_check_task_demo
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The name of the verification template.
	//
	// example:
	//
	// Data volume verification template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The total data volume comparison threshold.
	//
	// example:
	//
	// 0.5
	TotalCountThreshold *float32 `json:"totalCountThreshold,omitempty" xml:"totalCountThreshold,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 10001
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetDataCheckTaskConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTaskConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetBatchId() *int64 {
	return s.BatchId
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetCheckGlobalParams() *string {
	return s.CheckGlobalParams
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetCheckTemplateId() *string {
	return s.CheckTemplateId
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetCronExp() *string {
	return s.CronExp
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetDataCheckConfig() []*GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	return s.DataCheckConfig
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetDstDsId() *string {
	return s.DstDsId
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetDstDsName() *string {
	return s.DstDsName
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetDstDsType() *string {
	return s.DstDsType
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetDstEngineId() *string {
	return s.DstEngineId
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetDstEngineName() *string {
	return s.DstEngineName
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetDstEngineType() *string {
	return s.DstEngineType
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetExecuteType() *int32 {
	return s.ExecuteType
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetFullTableCount() *int32 {
	return s.FullTableCount
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetGroupCountThreshold() *float32 {
	return s.GroupCountThreshold
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetIsBuiltin() *int32 {
	return s.IsBuiltin
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetIsScheduled() *int32 {
	return s.IsScheduled
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetIsWhiteList() *int32 {
	return s.IsWhiteList
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetScheduleId() *int64 {
	return s.ScheduleId
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetScopeFilter() *GetDataCheckTaskConfigResponseBodyDataScopeFilter {
	return s.ScopeFilter
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetSourceGlobalParams() *string {
	return s.SourceGlobalParams
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetSrcDsId() *string {
	return s.SrcDsId
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetSrcDsName() *string {
	return s.SrcDsName
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetSrcDsType() *string {
	return s.SrcDsType
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetSrcEngineId() *string {
	return s.SrcEngineId
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetSrcEngineName() *string {
	return s.SrcEngineName
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetSrcEngineType() *string {
	return s.SrcEngineType
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetStartImmediately() *int32 {
	return s.StartImmediately
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetTargetGlobalParams() *string {
	return s.TargetGlobalParams
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetTaskConfigInfo() *string {
	return s.TaskConfigInfo
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetTaskMode() *int32 {
	return s.TaskMode
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetTotalCountThreshold() *float32 {
	return s.TotalCountThreshold
}

func (s *GetDataCheckTaskConfigResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetBatchId(v int64) *GetDataCheckTaskConfigResponseBodyData {
	s.BatchId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetCheckGlobalParams(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.CheckGlobalParams = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetCheckTemplateId(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.CheckTemplateId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetCheckType(v int32) *GetDataCheckTaskConfigResponseBodyData {
	s.CheckType = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetConcurrency(v int32) *GetDataCheckTaskConfigResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetCronExp(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.CronExp = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetDataCheckConfig(v []*GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) *GetDataCheckTaskConfigResponseBodyData {
	s.DataCheckConfig = v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetDstDsId(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.DstDsId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetDstDsName(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.DstDsName = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetDstDsType(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.DstDsType = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetDstEngineId(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.DstEngineId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetDstEngineName(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.DstEngineName = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetDstEngineType(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.DstEngineType = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetExecuteType(v int32) *GetDataCheckTaskConfigResponseBodyData {
	s.ExecuteType = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetFullTableCount(v int32) *GetDataCheckTaskConfigResponseBodyData {
	s.FullTableCount = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetGroupCountThreshold(v float32) *GetDataCheckTaskConfigResponseBodyData {
	s.GroupCountThreshold = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetIsBuiltin(v int32) *GetDataCheckTaskConfigResponseBodyData {
	s.IsBuiltin = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetIsScheduled(v int32) *GetDataCheckTaskConfigResponseBodyData {
	s.IsScheduled = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetIsWhiteList(v int32) *GetDataCheckTaskConfigResponseBodyData {
	s.IsWhiteList = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetRequestId(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetScheduleId(v int64) *GetDataCheckTaskConfigResponseBodyData {
	s.ScheduleId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetScopeFilter(v *GetDataCheckTaskConfigResponseBodyDataScopeFilter) *GetDataCheckTaskConfigResponseBodyData {
	s.ScopeFilter = v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetSourceGlobalParams(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.SourceGlobalParams = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetSrcDsId(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.SrcDsId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetSrcDsName(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.SrcDsName = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetSrcDsType(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.SrcDsType = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetSrcEngineId(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.SrcEngineId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetSrcEngineName(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.SrcEngineName = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetSrcEngineType(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.SrcEngineType = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetStartImmediately(v int32) *GetDataCheckTaskConfigResponseBodyData {
	s.StartImmediately = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetTargetGlobalParams(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.TargetGlobalParams = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetTaskConfigInfo(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.TaskConfigInfo = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetTaskDescription(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.TaskDescription = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetTaskId(v int64) *GetDataCheckTaskConfigResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetTaskMode(v int32) *GetDataCheckTaskConfigResponseBodyData {
	s.TaskMode = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetTaskName(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetTemplateName(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetTenantId(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetTotalCountThreshold(v float32) *GetDataCheckTaskConfigResponseBodyData {
	s.TotalCountThreshold = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) SetUid(v string) *GetDataCheckTaskConfigResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyData) Validate() error {
	if s.DataCheckConfig != nil {
		for _, item := range s.DataCheckConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScopeFilter != nil {
		if err := s.ScopeFilter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataCheckTaskConfigResponseBodyDataDataCheckConfig struct {
	// The validation algorithm.
	//
	// example:
	//
	// 1
	Algorithm *int32 `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// The batch ID.
	//
	// example:
	//
	// 20001
	BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// The batch size.
	//
	// example:
	//
	// 1000
	BatchSize *int32 `json:"batchSize,omitempty" xml:"batchSize,omitempty"`
	// The validation rule type. Valid values:
	//
	// - 0: data volume comparison.
	//
	// - 1: metric comparison.
	//
	// - 2: weak content comparison.
	//
	// - 3: custom comparison.
	//
	// - 4: full-text comparison.
	//
	// - 5: null value ratio comparison.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"checkType,omitempty" xml:"checkType,omitempty"`
	// The comparison type. Valid values: =, !=, >, <, >=, <=, contains, does not contain, and ==.
	//
	// example:
	//
	// =
	Comparator *string `json:"comparator,omitempty" xml:"comparator,omitempty"`
	// The reserved field.
	//
	// example:
	//
	// {}
	Extra interface{} `json:"extra,omitempty" xml:"extra,omitempty"`
	// The group data volume comparison threshold.
	//
	// example:
	//
	// 0.5
	GroupCountThreshold *float32 `json:"groupCountThreshold,omitempty" xml:"groupCountThreshold,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Specifies whether to perform a full table count.
	//
	// example:
	//
	// 1
	IsFullTableCount *int32 `json:"isFullTableCount,omitempty" xml:"isFullTableCount,omitempty"`
	// Specifies whether to skip the task. Valid values:
	//
	// - 0: No.
	//
	// - 1: Yes.
	//
	// example:
	//
	// 0
	IsSkipped *int32 `json:"isSkipped,omitempty" xml:"isSkipped,omitempty"`
	// The metric type. Valid values:
	//
	// - CUSTOM_METRIC_NUM: built-in NUM mode.
	//
	// - CUSTOM_METRIC_LEN: built-in LEN mode.
	//
	// - CUSTOM_METRIC_MIX: built-in MIX mode.
	//
	// example:
	//
	// CUSTOM_METRIC_MIX
	MetricType *string `json:"metricType,omitempty" xml:"metricType,omitempty"`
	// The source table columns. You can specify multiple columns separated by commas (,).
	//
	// example:
	//
	// col_a,col_b
	SourceColumns *string `json:"sourceColumns,omitempty" xml:"sourceColumns,omitempty"`
	// The source comparison key (the key field used for data comparison between the source and destination).
	//
	// example:
	//
	// id
	SourceCompareKey *string `json:"sourceCompareKey,omitempty" xml:"sourceCompareKey,omitempty"`
	// The source data source name.
	//
	// example:
	//
	// ds_demo
	SourceDataSource *string `json:"sourceDataSource,omitempty" xml:"sourceDataSource,omitempty"`
	// The GROUP BY clause for the source table.
	//
	// example:
	//
	// col_a,col_b
	SourceGroupClause *string `json:"sourceGroupClause,omitempty" xml:"sourceGroupClause,omitempty"`
	SourceHint        *string `json:"sourceHint,omitempty" xml:"sourceHint,omitempty"`
	// The source data source ID.
	//
	// example:
	//
	// 1001
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The source partition.
	//
	// example:
	//
	// ds=20260116
	SourcePartition *string `json:"sourcePartition,omitempty" xml:"sourcePartition,omitempty"`
	// The source SQL statement.
	//
	// example:
	//
	// SELECT 	- FROM t;
	SourceSql *string `json:"sourceSql,omitempty" xml:"sourceSql,omitempty"`
	// The source table.
	//
	// example:
	//
	// table_demo
	SourceTable *string `json:"sourceTable,omitempty" xml:"sourceTable,omitempty"`
	// The source data source type.
	//
	// example:
	//
	// Hive
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The WHERE clause for the source table.
	//
	// example:
	//
	// col_a > 0 and col_b = \\"x\\"
	SourceWhereClause *string `json:"sourceWhereClause,omitempty" xml:"sourceWhereClause,omitempty"`
	// The destination table columns. You can specify multiple columns separated by commas (,).
	//
	// example:
	//
	// col_a,col_b
	TargetColumns *string `json:"targetColumns,omitempty" xml:"targetColumns,omitempty"`
	// The destination comparison key (the key field used for data comparison between the source and destination).
	//
	// example:
	//
	// id
	TargetCompareKey *string `json:"targetCompareKey,omitempty" xml:"targetCompareKey,omitempty"`
	// The destination data source.
	//
	// example:
	//
	// ds_demo
	TargetDataSource *string `json:"targetDataSource,omitempty" xml:"targetDataSource,omitempty"`
	// The GROUP BY clause for the destination table.
	//
	// example:
	//
	// col_a,col_b
	TargetGroupClause *string `json:"targetGroupClause,omitempty" xml:"targetGroupClause,omitempty"`
	TargetHint        *string `json:"targetHint,omitempty" xml:"targetHint,omitempty"`
	// The destination ID.
	//
	// example:
	//
	// 2001
	TargetId *string `json:"targetId,omitempty" xml:"targetId,omitempty"`
	// The destination partition.
	//
	// example:
	//
	// ds=20260116
	TargetPartition *string `json:"targetPartition,omitempty" xml:"targetPartition,omitempty"`
	// The destination SQL statement.
	//
	// example:
	//
	// SELECT 	- FROM t;
	TargetSql *string `json:"targetSql,omitempty" xml:"targetSql,omitempty"`
	// The destination table.
	//
	// example:
	//
	// table_demo
	TargetTable *string `json:"targetTable,omitempty" xml:"targetTable,omitempty"`
	// The destination data source type.
	//
	// example:
	//
	// hive
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// The WHERE clause for the destination table.
	//
	// example:
	//
	// col_a > 0 and col_b = \\"x\\"
	TargetWhereClause *string `json:"targetWhereClause,omitempty" xml:"targetWhereClause,omitempty"`
	// The validation task configuration ID.
	//
	// example:
	//
	// 1001
	TaskConfigId *int64 `json:"taskConfigId,omitempty" xml:"taskConfigId,omitempty"`
	// The validation task configuration information (regular expression matching rules). This parameter takes effect only when taskMode is set to 1.
	//
	// example:
	//
	// lhm|lhm_dw|*
	TaskConfigInfo *string `json:"taskConfigInfo,omitempty" xml:"taskConfigInfo,omitempty"`
	// The validation batch token. Together with batchId, it identifies the result records generated by a validation batch.
	//
	// example:
	//
	// 9f2c7a1e4b8d****
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The total data volume comparison threshold.
	//
	// example:
	//
	// 0.5
	TotalCountThreshold *float32 `json:"totalCountThreshold,omitempty" xml:"totalCountThreshold,omitempty"`
}

func (s GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GoString() string {
	return s.String()
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetAlgorithm() *int32 {
	return s.Algorithm
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetBatchId() *int64 {
	return s.BatchId
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetBatchSize() *int32 {
	return s.BatchSize
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetComparator() *string {
	return s.Comparator
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetExtra() interface{} {
	return s.Extra
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetGroupCountThreshold() *float32 {
	return s.GroupCountThreshold
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetId() *int64 {
	return s.Id
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetIsFullTableCount() *int32 {
	return s.IsFullTableCount
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetIsSkipped() *int32 {
	return s.IsSkipped
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetMetricType() *string {
	return s.MetricType
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetSourceColumns() *string {
	return s.SourceColumns
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetSourceCompareKey() *string {
	return s.SourceCompareKey
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetSourceDataSource() *string {
	return s.SourceDataSource
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetSourceGroupClause() *string {
	return s.SourceGroupClause
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetSourceHint() *string {
	return s.SourceHint
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetSourceId() *string {
	return s.SourceId
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetSourcePartition() *string {
	return s.SourcePartition
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetSourceSql() *string {
	return s.SourceSql
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetSourceTable() *string {
	return s.SourceTable
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetSourceType() *string {
	return s.SourceType
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetSourceWhereClause() *string {
	return s.SourceWhereClause
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTargetColumns() *string {
	return s.TargetColumns
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTargetCompareKey() *string {
	return s.TargetCompareKey
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTargetDataSource() *string {
	return s.TargetDataSource
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTargetGroupClause() *string {
	return s.TargetGroupClause
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTargetHint() *string {
	return s.TargetHint
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTargetId() *string {
	return s.TargetId
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTargetPartition() *string {
	return s.TargetPartition
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTargetSql() *string {
	return s.TargetSql
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTargetTable() *string {
	return s.TargetTable
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTargetType() *string {
	return s.TargetType
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTargetWhereClause() *string {
	return s.TargetWhereClause
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTaskConfigId() *int64 {
	return s.TaskConfigId
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTaskConfigInfo() *string {
	return s.TaskConfigInfo
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetToken() *string {
	return s.Token
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) GetTotalCountThreshold() *float32 {
	return s.TotalCountThreshold
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetAlgorithm(v int32) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.Algorithm = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetBatchId(v int64) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.BatchId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetBatchSize(v int32) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.BatchSize = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetCheckType(v int32) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.CheckType = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetComparator(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.Comparator = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetExtra(v interface{}) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.Extra = v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetGroupCountThreshold(v float32) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.GroupCountThreshold = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetId(v int64) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.Id = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetIsFullTableCount(v int32) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.IsFullTableCount = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetIsSkipped(v int32) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.IsSkipped = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetMetricType(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.MetricType = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetSourceColumns(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.SourceColumns = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetSourceCompareKey(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.SourceCompareKey = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetSourceDataSource(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.SourceDataSource = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetSourceGroupClause(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.SourceGroupClause = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetSourceHint(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.SourceHint = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetSourceId(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.SourceId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetSourcePartition(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.SourcePartition = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetSourceSql(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.SourceSql = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetSourceTable(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.SourceTable = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetSourceType(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.SourceType = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetSourceWhereClause(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.SourceWhereClause = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTargetColumns(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TargetColumns = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTargetCompareKey(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TargetCompareKey = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTargetDataSource(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TargetDataSource = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTargetGroupClause(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TargetGroupClause = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTargetHint(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TargetHint = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTargetId(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TargetId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTargetPartition(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TargetPartition = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTargetSql(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TargetSql = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTargetTable(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TargetTable = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTargetType(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TargetType = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTargetWhereClause(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TargetWhereClause = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTaskConfigId(v int64) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TaskConfigId = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTaskConfigInfo(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TaskConfigInfo = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetToken(v string) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.Token = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) SetTotalCountThreshold(v float32) *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig {
	s.TotalCountThreshold = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataDataCheckConfig) Validate() error {
	return dara.Validate(s)
}

type GetDataCheckTaskConfigResponseBodyDataScopeFilter struct {
	// The end time.
	//
	// example:
	//
	// 2026-01-14 13:59:00
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// The last N parameter.
	//
	// example:
	//
	// 7
	LastN *int32 `json:"lastN,omitempty" xml:"lastN,omitempty"`
	// The filter type.
	//
	// example:
	//
	// 1
	ScopeFilterType *int32 `json:"scopeFilterType,omitempty" xml:"scopeFilterType,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2026-01-14 11:21:00
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetDataCheckTaskConfigResponseBodyDataScopeFilter) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTaskConfigResponseBodyDataScopeFilter) GoString() string {
	return s.String()
}

func (s *GetDataCheckTaskConfigResponseBodyDataScopeFilter) GetEnd() *string {
	return s.End
}

func (s *GetDataCheckTaskConfigResponseBodyDataScopeFilter) GetLastN() *int32 {
	return s.LastN
}

func (s *GetDataCheckTaskConfigResponseBodyDataScopeFilter) GetScopeFilterType() *int32 {
	return s.ScopeFilterType
}

func (s *GetDataCheckTaskConfigResponseBodyDataScopeFilter) GetStart() *string {
	return s.Start
}

func (s *GetDataCheckTaskConfigResponseBodyDataScopeFilter) SetEnd(v string) *GetDataCheckTaskConfigResponseBodyDataScopeFilter {
	s.End = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataScopeFilter) SetLastN(v int32) *GetDataCheckTaskConfigResponseBodyDataScopeFilter {
	s.LastN = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataScopeFilter) SetScopeFilterType(v int32) *GetDataCheckTaskConfigResponseBodyDataScopeFilter {
	s.ScopeFilterType = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataScopeFilter) SetStart(v string) *GetDataCheckTaskConfigResponseBodyDataScopeFilter {
	s.Start = &v
	return s
}

func (s *GetDataCheckTaskConfigResponseBodyDataScopeFilter) Validate() error {
	return dara.Validate(s)
}
