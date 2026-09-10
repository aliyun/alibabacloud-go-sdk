// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckTaskHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataCheckTaskHistoryResponseBodyData) *ListDataCheckTaskHistoryResponseBody
	GetData() []*ListDataCheckTaskHistoryResponseBodyData
	SetErrCode(v string) *ListDataCheckTaskHistoryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListDataCheckTaskHistoryResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *ListDataCheckTaskHistoryResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDataCheckTaskHistoryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataCheckTaskHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataCheckTaskHistoryResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListDataCheckTaskHistoryResponseBody
	GetTotalCount() *int32
}

type ListDataCheckTaskHistoryResponseBody struct {
	// The data list returned by the operation. For the element structure, see the child field descriptions.
	Data []*ListDataCheckTaskHistoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// The page size, which is the number of records returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, check errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that match the query conditions. This value is used for pagination.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDataCheckTaskHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckTaskHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCheckTaskHistoryResponseBody) GetData() []*ListDataCheckTaskHistoryResponseBodyData {
	return s.Data
}

func (s *ListDataCheckTaskHistoryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListDataCheckTaskHistoryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListDataCheckTaskHistoryResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDataCheckTaskHistoryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCheckTaskHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCheckTaskHistoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataCheckTaskHistoryResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataCheckTaskHistoryResponseBody) SetData(v []*ListDataCheckTaskHistoryResponseBodyData) *ListDataCheckTaskHistoryResponseBody {
	s.Data = v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBody) SetErrCode(v string) *ListDataCheckTaskHistoryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBody) SetErrMessage(v string) *ListDataCheckTaskHistoryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBody) SetPageIndex(v int32) *ListDataCheckTaskHistoryResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBody) SetPageSize(v int32) *ListDataCheckTaskHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBody) SetRequestId(v string) *ListDataCheckTaskHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBody) SetSuccess(v bool) *ListDataCheckTaskHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBody) SetTotalCount(v int32) *ListDataCheckTaskHistoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBody) Validate() error {
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

type ListDataCheckTaskHistoryResponseBodyData struct {
	// The batch ID.
	//
	// example:
	//
	// 20001
	BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// The business field.
	//
	// example:
	//
	// lhm
	Biz *string `json:"biz,omitempty" xml:"biz,omitempty"`
	// The execution result. Valid values: no record, passed, or not passed.
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
	// The concurrency of the batch.
	//
	// example:
	//
	// 5
	Concurrency *int64 `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
	// The creator.
	//
	// example:
	//
	// user001
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The scheduling cycle expression.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	CronExp *string `json:"cronExp,omitempty" xml:"cronExp,omitempty"`
	// The end time, in the format of YYYY-MM-DD HH:MM:SS.
	//
	// example:
	//
	// 2026-01-16 12:00:00
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
	// The execution status. Valid values: pending, running, stopped, failed, or completed.
	//
	// example:
	//
	// 1
	ExecStatus *int32 `json:"execStatus,omitempty" xml:"execStatus,omitempty"`
	// The execution duration, in the format of HH:MM:SS.
	//
	// example:
	//
	// 2026-01-16 10:00:00
	ExecTime *string `json:"execTime,omitempty" xml:"execTime,omitempty"`
	// The reserved field.
	//
	// example:
	//
	// {}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-01-16 10:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2026-01-16 10:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The updater.
	//
	// example:
	//
	// user001
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The original batch ID.
	//
	// example:
	//
	// 1001
	OriginBatchId *int64 `json:"originBatchId,omitempty" xml:"originBatchId,omitempty"`
	// The check pass rate.
	//
	// example:
	//
	// 0.95
	PassProcess *float64 `json:"passProcess,omitempty" xml:"passProcess,omitempty"`
	// The pass rate (export report field), calculated by dividing the number of passed tables by the total number of checked tables. The value is returned as a string with a percent sign and two decimal places (for example, 100.00%). If no checked table data exists, the value is -.
	//
	// example:
	//
	// 100.00%
	PassProcessExport *string `json:"passProcessExport,omitempty" xml:"passProcessExport,omitempty"`
	// The task progress.
	//
	// example:
	//
	// 0.95
	Progress *float64 `json:"progress,omitempty" xml:"progress,omitempty"`
	// The check report time, which is the completion time of the last job.
	//
	// example:
	//
	// 2024-01-01 12:00:00
	ReportTime *string `json:"reportTime,omitempty" xml:"reportTime,omitempty"`
	// The check report title.
	//
	// example:
	//
	// Data Validation Report
	ReportTitle *string `json:"reportTitle,omitempty" xml:"reportTitle,omitempty"`
	// The scheduled task ID.
	//
	// example:
	//
	// 1001
	ScheduleId *int64 `json:"scheduleId,omitempty" xml:"scheduleId,omitempty"`
	// The task number.
	//
	// example:
	//
	// SEQ20260116001
	SeqId *string `json:"seqId,omitempty" xml:"seqId,omitempty"`
	// The number of skipped tables.
	//
	// example:
	//
	// 3
	SkipTableNum *int32 `json:"skipTableNum,omitempty" xml:"skipTableNum,omitempty"`
	// The start time, in the format of YYYY-MM-DD HH:MM:SS.
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
}

func (s ListDataCheckTaskHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckTaskHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetBatchId() *int64 {
	return s.BatchId
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetBiz() *string {
	return s.Biz
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetCheckResult() *int32 {
	return s.CheckResult
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetCheckTableNum() *int64 {
	return s.CheckTableNum
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetCronExp() *string {
	return s.CronExp
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetErrorTableNum() *int64 {
	return s.ErrorTableNum
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetExecStatus() *int32 {
	return s.ExecStatus
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetExecTime() *string {
	return s.ExecTime
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetOperator() *string {
	return s.Operator
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetOriginBatchId() *int64 {
	return s.OriginBatchId
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetPassProcess() *float64 {
	return s.PassProcess
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetPassProcessExport() *string {
	return s.PassProcessExport
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetProgress() *float64 {
	return s.Progress
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetReportTime() *string {
	return s.ReportTime
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetReportTitle() *string {
	return s.ReportTitle
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetScheduleId() *int64 {
	return s.ScheduleId
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetSeqId() *string {
	return s.SeqId
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetSkipTableNum() *int32 {
	return s.SkipTableNum
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDataCheckTaskHistoryResponseBodyData) GetSuccessfulTableNum() *int64 {
	return s.SuccessfulTableNum
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetBatchId(v int64) *ListDataCheckTaskHistoryResponseBodyData {
	s.BatchId = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetBiz(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.Biz = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetCheckResult(v int32) *ListDataCheckTaskHistoryResponseBodyData {
	s.CheckResult = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetCheckTableNum(v int64) *ListDataCheckTaskHistoryResponseBodyData {
	s.CheckTableNum = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetConcurrency(v int64) *ListDataCheckTaskHistoryResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetCreator(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.Creator = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetCronExp(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.CronExp = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetEndTime(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetErrorMsg(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetErrorTableNum(v int64) *ListDataCheckTaskHistoryResponseBodyData {
	s.ErrorTableNum = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetExecStatus(v int32) *ListDataCheckTaskHistoryResponseBodyData {
	s.ExecStatus = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetExecTime(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.ExecTime = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetExtra(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.Extra = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetGmtCreate(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetGmtModified(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetOperator(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.Operator = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetOriginBatchId(v int64) *ListDataCheckTaskHistoryResponseBodyData {
	s.OriginBatchId = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetPassProcess(v float64) *ListDataCheckTaskHistoryResponseBodyData {
	s.PassProcess = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetPassProcessExport(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.PassProcessExport = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetProgress(v float64) *ListDataCheckTaskHistoryResponseBodyData {
	s.Progress = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetReportTime(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.ReportTime = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetReportTitle(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.ReportTitle = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetScheduleId(v int64) *ListDataCheckTaskHistoryResponseBodyData {
	s.ScheduleId = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetSeqId(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.SeqId = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetSkipTableNum(v int32) *ListDataCheckTaskHistoryResponseBodyData {
	s.SkipTableNum = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetStartTime(v string) *ListDataCheckTaskHistoryResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) SetSuccessfulTableNum(v int64) *ListDataCheckTaskHistoryResponseBodyData {
	s.SuccessfulTableNum = &v
	return s
}

func (s *ListDataCheckTaskHistoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
