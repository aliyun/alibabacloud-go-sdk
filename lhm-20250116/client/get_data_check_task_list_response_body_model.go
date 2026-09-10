// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetDataCheckTaskListResponseBodyData) *GetDataCheckTaskListResponseBody
	GetData() []*GetDataCheckTaskListResponseBodyData
	SetErrCode(v string) *GetDataCheckTaskListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetDataCheckTaskListResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *GetDataCheckTaskListResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *GetDataCheckTaskListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetDataCheckTaskListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCheckTaskListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetDataCheckTaskListResponseBody
	GetTotalCount() *int32
}

type GetDataCheckTaskListResponseBody struct {
	// The task list.
	Data []*GetDataCheckTaskListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The page size, which indicates the number of records returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues related to this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, check errCode and errMessage for details.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that match the query conditions. This parameter is used for pagination.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetDataCheckTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCheckTaskListResponseBody) GetData() []*GetDataCheckTaskListResponseBodyData {
	return s.Data
}

func (s *GetDataCheckTaskListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDataCheckTaskListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetDataCheckTaskListResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetDataCheckTaskListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDataCheckTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCheckTaskListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCheckTaskListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetDataCheckTaskListResponseBody) SetData(v []*GetDataCheckTaskListResponseBodyData) *GetDataCheckTaskListResponseBody {
	s.Data = v
	return s
}

func (s *GetDataCheckTaskListResponseBody) SetErrCode(v string) *GetDataCheckTaskListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDataCheckTaskListResponseBody) SetErrMessage(v string) *GetDataCheckTaskListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetDataCheckTaskListResponseBody) SetPageIndex(v int32) *GetDataCheckTaskListResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetDataCheckTaskListResponseBody) SetPageSize(v int32) *GetDataCheckTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetDataCheckTaskListResponseBody) SetRequestId(v string) *GetDataCheckTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCheckTaskListResponseBody) SetSuccess(v bool) *GetDataCheckTaskListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCheckTaskListResponseBody) SetTotalCount(v int32) *GetDataCheckTaskListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetDataCheckTaskListResponseBody) Validate() error {
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

type GetDataCheckTaskListResponseBodyData struct {
	// The check result. Valid values:
	//
	// - 0: no record.
	//
	// - 1: passed.
	//
	// - 2: failed.
	//
	// example:
	//
	// 1
	CheckResult *int32 `json:"checkResult,omitempty" xml:"checkResult,omitempty"`
	// The number of checked tables.
	//
	// example:
	//
	// 100
	CheckTableNum *int64 `json:"checkTableNum,omitempty" xml:"checkTableNum,omitempty"`
	// The check template ID.
	//
	// example:
	//
	// 1001
	CheckTemplateId *string `json:"checkTemplateId,omitempty" xml:"checkTemplateId,omitempty"`
	// The check type. Valid values:
	//
	// - 0: data volume comparison.
	//
	// - 1: metric comparison.
	//
	// - 2: weak content comparison.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"checkType,omitempty" xml:"checkType,omitempty"`
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
	// The destination check engine ID.
	//
	// example:
	//
	// 2001
	DstEngineId *string `json:"dstEngineId,omitempty" xml:"dstEngineId,omitempty"`
	// The destination check engine name.
	//
	// example:
	//
	// engine_demo
	DstEngineName *string `json:"dstEngineName,omitempty" xml:"dstEngineName,omitempty"`
	// The destination check engine type.
	//
	// example:
	//
	// Tez
	DstEngineType *string `json:"dstEngineType,omitempty" xml:"dstEngineType,omitempty"`
	// The end time.
	//
	// example:
	//
	// 2026-01-16 10:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The error message.
	//
	// example:
	//
	// connection timeout
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The number of error tables.
	//
	// example:
	//
	// 5
	ErrorTableNum *int64 `json:"errorTableNum,omitempty" xml:"errorTableNum,omitempty"`
	// The execution status. Valid values:
	//
	// - 0: pending.
	//
	// - 1: running.
	//
	// - 2: stopped.
	//
	// - 3: failed.
	//
	// - 4: completed.
	//
	// example:
	//
	// 1
	ExecStatus *int32 `json:"execStatus,omitempty" xml:"execStatus,omitempty"`
	// The execution duration.
	//
	// example:
	//
	// 2026-01-16 10:00:00
	ExecTime *string `json:"execTime,omitempty" xml:"execTime,omitempty"`
	// The execution type. Valid values:
	//
	// - 0: immediate execution.
	//
	// - 1: scheduled execution.
	//
	// example:
	//
	// 0
	ExecuteType *int32 `json:"executeType,omitempty" xml:"executeType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-01-16 10:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modified time.
	//
	// example:
	//
	// 2026-01-16 10:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether scheduling is enabled. Valid values:
	//
	// - 0: Disabled.
	//
	// - 1: Enabled.
	//
	// example:
	//
	// 0
	IsScheduled *int32 `json:"isScheduled,omitempty" xml:"isScheduled,omitempty"`
	// The latest batch ID.
	//
	// example:
	//
	// 20001
	LastBatchId *int64 `json:"lastBatchId,omitempty" xml:"lastBatchId,omitempty"`
	// The latest completed batch ID.
	//
	// example:
	//
	// 833
	LastFinishedId *int64 `json:"lastFinishedId,omitempty" xml:"lastFinishedId,omitempty"`
	// The check pass rate.
	PassProcess map[string]interface{} `json:"passProcess,omitempty" xml:"passProcess,omitempty"`
	// The execution progress (0-1).
	//
	// example:
	//
	// 1.0
	Process *float64 `json:"process,omitempty" xml:"process,omitempty"`
	// The report time.
	//
	// example:
	//
	// 2024-01-01 12:00:00
	ReportTime *string `json:"reportTime,omitempty" xml:"reportTime,omitempty"`
	// The report title.
	//
	// example:
	//
	// Data validation report
	ReportTitle *string `json:"reportTitle,omitempty" xml:"reportTitle,omitempty"`
	// The number of skipped tables.
	//
	// example:
	//
	// 3
	SkipTableNum *int64 `json:"skipTableNum,omitempty" xml:"skipTableNum,omitempty"`
	// The source data source ID.
	//
	// example:
	//
	// 1001
	SrcDsId *string `json:"srcDsId,omitempty" xml:"srcDsId,omitempty"`
	// The source data source name.
	//
	// example:
	//
	// ds_demo
	SrcDsName *string `json:"srcDsName,omitempty" xml:"srcDsName,omitempty"`
	// The source data source type.
	//
	// example:
	//
	// Hive
	SrcDsType *string `json:"srcDsType,omitempty" xml:"srcDsType,omitempty"`
	// The source check engine ID.
	//
	// example:
	//
	// 1001
	SrcEngineId *string `json:"srcEngineId,omitempty" xml:"srcEngineId,omitempty"`
	// The source check engine name.
	//
	// example:
	//
	// engine_demo
	SrcEngineName *string `json:"srcEngineName,omitempty" xml:"srcEngineName,omitempty"`
	// The source check engine type.
	//
	// example:
	//
	// Tez
	SrcEngineType *string `json:"srcEngineType,omitempty" xml:"srcEngineType,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2026-01-16 10:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The number of successful tables.
	//
	// example:
	//
	// 95
	SuccessfulTableNum *int64 `json:"successfulTableNum,omitempty" xml:"successfulTableNum,omitempty"`
	// The task description.
	//
	// example:
	//
	// Data validation task description
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// The creation mode. Valid values:
	//
	// - 0: table-by-table fine-grained mode.
	//
	// - 1: same-schema batch mode.
	//
	// example:
	//
	// 0
	TaskMode *int32 `json:"taskMode,omitempty" xml:"taskMode,omitempty"`
	// The task name.
	//
	// example:
	//
	// data_check_task_demo
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The check template name.
	//
	// example:
	//
	// Row Count Validation Template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
}

func (s GetDataCheckTaskListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTaskListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataCheckTaskListResponseBodyData) GetCheckResult() *int32 {
	return s.CheckResult
}

func (s *GetDataCheckTaskListResponseBodyData) GetCheckTableNum() *int64 {
	return s.CheckTableNum
}

func (s *GetDataCheckTaskListResponseBodyData) GetCheckTemplateId() *string {
	return s.CheckTemplateId
}

func (s *GetDataCheckTaskListResponseBodyData) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GetDataCheckTaskListResponseBodyData) GetDstDsId() *string {
	return s.DstDsId
}

func (s *GetDataCheckTaskListResponseBodyData) GetDstDsName() *string {
	return s.DstDsName
}

func (s *GetDataCheckTaskListResponseBodyData) GetDstDsType() *string {
	return s.DstDsType
}

func (s *GetDataCheckTaskListResponseBodyData) GetDstEngineId() *string {
	return s.DstEngineId
}

func (s *GetDataCheckTaskListResponseBodyData) GetDstEngineName() *string {
	return s.DstEngineName
}

func (s *GetDataCheckTaskListResponseBodyData) GetDstEngineType() *string {
	return s.DstEngineType
}

func (s *GetDataCheckTaskListResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetDataCheckTaskListResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetDataCheckTaskListResponseBodyData) GetErrorTableNum() *int64 {
	return s.ErrorTableNum
}

func (s *GetDataCheckTaskListResponseBodyData) GetExecStatus() *int32 {
	return s.ExecStatus
}

func (s *GetDataCheckTaskListResponseBodyData) GetExecTime() *string {
	return s.ExecTime
}

func (s *GetDataCheckTaskListResponseBodyData) GetExecuteType() *int32 {
	return s.ExecuteType
}

func (s *GetDataCheckTaskListResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDataCheckTaskListResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetDataCheckTaskListResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetDataCheckTaskListResponseBodyData) GetIsScheduled() *int32 {
	return s.IsScheduled
}

func (s *GetDataCheckTaskListResponseBodyData) GetLastBatchId() *int64 {
	return s.LastBatchId
}

func (s *GetDataCheckTaskListResponseBodyData) GetLastFinishedId() *int64 {
	return s.LastFinishedId
}

func (s *GetDataCheckTaskListResponseBodyData) GetPassProcess() map[string]interface{} {
	return s.PassProcess
}

func (s *GetDataCheckTaskListResponseBodyData) GetProcess() *float64 {
	return s.Process
}

func (s *GetDataCheckTaskListResponseBodyData) GetReportTime() *string {
	return s.ReportTime
}

func (s *GetDataCheckTaskListResponseBodyData) GetReportTitle() *string {
	return s.ReportTitle
}

func (s *GetDataCheckTaskListResponseBodyData) GetSkipTableNum() *int64 {
	return s.SkipTableNum
}

func (s *GetDataCheckTaskListResponseBodyData) GetSrcDsId() *string {
	return s.SrcDsId
}

func (s *GetDataCheckTaskListResponseBodyData) GetSrcDsName() *string {
	return s.SrcDsName
}

func (s *GetDataCheckTaskListResponseBodyData) GetSrcDsType() *string {
	return s.SrcDsType
}

func (s *GetDataCheckTaskListResponseBodyData) GetSrcEngineId() *string {
	return s.SrcEngineId
}

func (s *GetDataCheckTaskListResponseBodyData) GetSrcEngineName() *string {
	return s.SrcEngineName
}

func (s *GetDataCheckTaskListResponseBodyData) GetSrcEngineType() *string {
	return s.SrcEngineType
}

func (s *GetDataCheckTaskListResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetDataCheckTaskListResponseBodyData) GetSuccessfulTableNum() *int64 {
	return s.SuccessfulTableNum
}

func (s *GetDataCheckTaskListResponseBodyData) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *GetDataCheckTaskListResponseBodyData) GetTaskMode() *int32 {
	return s.TaskMode
}

func (s *GetDataCheckTaskListResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *GetDataCheckTaskListResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetDataCheckTaskListResponseBodyData) SetCheckResult(v int32) *GetDataCheckTaskListResponseBodyData {
	s.CheckResult = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetCheckTableNum(v int64) *GetDataCheckTaskListResponseBodyData {
	s.CheckTableNum = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetCheckTemplateId(v string) *GetDataCheckTaskListResponseBodyData {
	s.CheckTemplateId = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetCheckType(v int32) *GetDataCheckTaskListResponseBodyData {
	s.CheckType = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetDstDsId(v string) *GetDataCheckTaskListResponseBodyData {
	s.DstDsId = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetDstDsName(v string) *GetDataCheckTaskListResponseBodyData {
	s.DstDsName = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetDstDsType(v string) *GetDataCheckTaskListResponseBodyData {
	s.DstDsType = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetDstEngineId(v string) *GetDataCheckTaskListResponseBodyData {
	s.DstEngineId = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetDstEngineName(v string) *GetDataCheckTaskListResponseBodyData {
	s.DstEngineName = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetDstEngineType(v string) *GetDataCheckTaskListResponseBodyData {
	s.DstEngineType = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetEndTime(v string) *GetDataCheckTaskListResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetErrorMsg(v string) *GetDataCheckTaskListResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetErrorTableNum(v int64) *GetDataCheckTaskListResponseBodyData {
	s.ErrorTableNum = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetExecStatus(v int32) *GetDataCheckTaskListResponseBodyData {
	s.ExecStatus = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetExecTime(v string) *GetDataCheckTaskListResponseBodyData {
	s.ExecTime = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetExecuteType(v int32) *GetDataCheckTaskListResponseBodyData {
	s.ExecuteType = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetGmtCreate(v string) *GetDataCheckTaskListResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetGmtModified(v string) *GetDataCheckTaskListResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetId(v int64) *GetDataCheckTaskListResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetIsScheduled(v int32) *GetDataCheckTaskListResponseBodyData {
	s.IsScheduled = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetLastBatchId(v int64) *GetDataCheckTaskListResponseBodyData {
	s.LastBatchId = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetLastFinishedId(v int64) *GetDataCheckTaskListResponseBodyData {
	s.LastFinishedId = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetPassProcess(v map[string]interface{}) *GetDataCheckTaskListResponseBodyData {
	s.PassProcess = v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetProcess(v float64) *GetDataCheckTaskListResponseBodyData {
	s.Process = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetReportTime(v string) *GetDataCheckTaskListResponseBodyData {
	s.ReportTime = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetReportTitle(v string) *GetDataCheckTaskListResponseBodyData {
	s.ReportTitle = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetSkipTableNum(v int64) *GetDataCheckTaskListResponseBodyData {
	s.SkipTableNum = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetSrcDsId(v string) *GetDataCheckTaskListResponseBodyData {
	s.SrcDsId = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetSrcDsName(v string) *GetDataCheckTaskListResponseBodyData {
	s.SrcDsName = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetSrcDsType(v string) *GetDataCheckTaskListResponseBodyData {
	s.SrcDsType = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetSrcEngineId(v string) *GetDataCheckTaskListResponseBodyData {
	s.SrcEngineId = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetSrcEngineName(v string) *GetDataCheckTaskListResponseBodyData {
	s.SrcEngineName = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetSrcEngineType(v string) *GetDataCheckTaskListResponseBodyData {
	s.SrcEngineType = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetStartTime(v string) *GetDataCheckTaskListResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetSuccessfulTableNum(v int64) *GetDataCheckTaskListResponseBodyData {
	s.SuccessfulTableNum = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetTaskDescription(v string) *GetDataCheckTaskListResponseBodyData {
	s.TaskDescription = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetTaskMode(v int32) *GetDataCheckTaskListResponseBodyData {
	s.TaskMode = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetTaskName(v string) *GetDataCheckTaskListResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) SetTemplateName(v string) *GetDataCheckTaskListResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *GetDataCheckTaskListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
