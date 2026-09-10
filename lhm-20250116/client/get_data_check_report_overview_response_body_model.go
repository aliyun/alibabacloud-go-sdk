// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckReportOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataCheckReportOverviewResponseBodyData) *GetDataCheckReportOverviewResponseBody
	GetData() *GetDataCheckReportOverviewResponseBodyData
	SetErrCode(v string) *GetDataCheckReportOverviewResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetDataCheckReportOverviewResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetDataCheckReportOverviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataCheckReportOverviewResponseBody
	GetSuccess() *bool
}

type GetDataCheckReportOverviewResponseBody struct {
	// The response data.
	Data *GetDataCheckReportOverviewResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault information code.
	//
	// example:
	//
	// None
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// not supported.pos 5459, line 196, column 14, token IDENTIFIER settings
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID. This value uniquely identifies the call. Provide this value when troubleshooting issues.
	//
	// example:
	//
	// FFF386FC-295C-5D2E-B2FE-410003095F06
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// False
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDataCheckReportOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckReportOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataCheckReportOverviewResponseBody) GetData() *GetDataCheckReportOverviewResponseBodyData {
	return s.Data
}

func (s *GetDataCheckReportOverviewResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDataCheckReportOverviewResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetDataCheckReportOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCheckReportOverviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataCheckReportOverviewResponseBody) SetData(v *GetDataCheckReportOverviewResponseBodyData) *GetDataCheckReportOverviewResponseBody {
	s.Data = v
	return s
}

func (s *GetDataCheckReportOverviewResponseBody) SetErrCode(v string) *GetDataCheckReportOverviewResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBody) SetErrMessage(v string) *GetDataCheckReportOverviewResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBody) SetRequestId(v string) *GetDataCheckReportOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBody) SetSuccess(v bool) *GetDataCheckReportOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataCheckReportOverviewResponseBodyData struct {
	// The ID of the validation job (batch).
	//
	// example:
	//
	// 20001
	BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// The number of validated fields.
	//
	// example:
	//
	// 10
	CheckColumnCount *int64 `json:"checkColumnCount,omitempty" xml:"checkColumnCount,omitempty"`
	// The number of validated partitions.
	//
	// example:
	//
	// 5
	CheckPtCount *int64 `json:"checkPtCount,omitempty" xml:"checkPtCount,omitempty"`
	// The validation result. Valid values:
	//
	// - 0: No record.
	//
	// - 1: Passed.
	//
	// - 2: Failed.
	//
	// example:
	//
	// 0
	CheckResult *int32 `json:"checkResult,omitempty" xml:"checkResult,omitempty"`
	// The number of validated data rows.
	//
	// example:
	//
	// 1000
	CheckRowCount *int64 `json:"checkRowCount,omitempty" xml:"checkRowCount,omitempty"`
	// The number of rows that passed validation.
	//
	// example:
	//
	// 1000
	CheckRowPassCount *int64 `json:"checkRowPassCount,omitempty" xml:"checkRowPassCount,omitempty"`
	// The row pass rate for the export report. This value is calculated by dividing the number of passed rows by the total number of validated rows. The value is returned as a percentage string with two decimal places.
	//
	// example:
	//
	// 100.00%
	CheckRowPassExport *string `json:"checkRowPassExport,omitempty" xml:"checkRowPassExport,omitempty"`
	// The number of validation SQL statements.
	//
	// example:
	//
	// 10
	CheckSqlNum *int64 `json:"checkSqlNum,omitempty" xml:"checkSqlNum,omitempty"`
	// The number of validated tables.
	//
	// example:
	//
	// 100
	CheckTableNum *int64 `json:"checkTableNum,omitempty" xml:"checkTableNum,omitempty"`
	// The validation template name. This field is available only for metric validation.
	//
	// example:
	//
	// 1001
	CheckTemplateId *string `json:"checkTemplateId,omitempty" xml:"checkTemplateId,omitempty"`
	// The validation template name. This field is available only for metric validation.
	//
	// example:
	//
	// Data volume comparison built-in template
	CheckTemplateName *string `json:"checkTemplateName,omitempty" xml:"checkTemplateName,omitempty"`
	// The validation type. Valid values:
	//
	// - 0: data volume comparison.
	//
	// - 1: metric comparison.
	//
	// - 2: weak content comparison.
	//
	// example:
	//
	// 0
	CheckType *int32 `json:"checkType,omitempty" xml:"checkType,omitempty"`
	// The name of the destination datasource.
	//
	// example:
	//
	// ds_demo
	DstDsName *string `json:"dstDsName,omitempty" xml:"dstDsName,omitempty"`
	// The type of the destination datasource.
	//
	// example:
	//
	// Hive
	DstDsType *string `json:"dstDsType,omitempty" xml:"dstDsType,omitempty"`
	// The number of tables with errors.
	//
	// example:
	//
	// 5
	ErrorTableNum *int64 `json:"errorTableNum,omitempty" xml:"errorTableNum,omitempty"`
	// The number of fields that passed validation.
	//
	// example:
	//
	// 8
	PassColumnCount *int64 `json:"passColumnCount,omitempty" xml:"passColumnCount,omitempty"`
	// The number of metrics that passed validation.
	//
	// example:
	//
	// 1.0
	PassColumnRate *float64 `json:"passColumnRate,omitempty" xml:"passColumnRate,omitempty"`
	// The pass rate.
	//
	// example:
	//
	// 0.95
	PassProcess *float64 `json:"passProcess,omitempty" xml:"passProcess,omitempty"`
	// The pass rate for the export report. This value is calculated by dividing the number of passed tables by the total number of validated tables. The value is returned as a percentage string with two decimal places (for example, 100.00%). A hyphen (-) is returned when no validated table data exists.
	//
	// example:
	//
	// 100.00%
	PassProcessExport *string `json:"passProcessExport,omitempty" xml:"passProcessExport,omitempty"`
	// The number of partitions that passed validation.
	//
	// example:
	//
	// 5
	PassPtNum *int64 `json:"passPtNum,omitempty" xml:"passPtNum,omitempty"`
	// The partition pass rate for the export report. This value is calculated by dividing the number of passed partitions by the total number of validated partitions. The value is returned as a percentage string with four decimal places. A hyphen (-) is returned when no partition data exists.
	//
	// example:
	//
	// 100.0000%
	PassPtProcessExport *string `json:"passPtProcessExport,omitempty" xml:"passPtProcessExport,omitempty"`
	// The number of tables that passed validation.
	//
	// example:
	//
	// 10
	PassTableNum *int64 `json:"passTableNum,omitempty" xml:"passTableNum,omitempty"`
	// The partition pass rate.
	//
	// example:
	//
	// 1.0
	PtPassProcess *float64 `json:"ptPassProcess,omitempty" xml:"ptPassProcess,omitempty"`
	// The report generation message.
	//
	// example:
	//
	// Validation report refresh completed
	ReportGenerateMessage *string `json:"reportGenerateMessage,omitempty" xml:"reportGenerateMessage,omitempty"`
	// The validation report status. Valid values:
	//
	// - 0: Not generated.
	//
	// - 1: Generating.
	//
	// - 2: Generated.
	//
	// example:
	//
	// 0
	ReportStatus *int32 `json:"reportStatus,omitempty" xml:"reportStatus,omitempty"`
	// The time when the report was generated.
	//
	// example:
	//
	// 2024-01-01 12:00:00
	ReportTime *string `json:"reportTime,omitempty" xml:"reportTime,omitempty"`
	// The title of the validation report.
	//
	// example:
	//
	// Data Validation Report
	ReportTitle *string `json:"reportTitle,omitempty" xml:"reportTitle,omitempty"`
	// The number of skipped partitions.
	//
	// example:
	//
	// 0
	SkipPtNum *int64 `json:"skipPtNum,omitempty" xml:"skipPtNum,omitempty"`
	// The number of skipped tables.
	//
	// example:
	//
	// 3
	SkipTableNum *int32 `json:"skipTableNum,omitempty" xml:"skipTableNum,omitempty"`
	// The name of the source datasource.
	//
	// example:
	//
	// ds_demo
	SrcDsName *string `json:"srcDsName,omitempty" xml:"srcDsName,omitempty"`
	// The type of the source datasource.
	//
	// example:
	//
	// Hive
	SrcDsType *string `json:"srcDsType,omitempty" xml:"srcDsType,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2024-01-01 11:00:00
	TaskCreateTime *string `json:"taskCreateTime,omitempty" xml:"taskCreateTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 551
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The time when the task was last modified.
	//
	// example:
	//
	// 2024-01-01 12:00:00
	TaskModifyTime *string `json:"taskModifyTime,omitempty" xml:"taskModifyTime,omitempty"`
	// The task name.
	//
	// example:
	//
	// data_check_task_demo
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s GetDataCheckReportOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckReportOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetBatchId() *int64 {
	return s.BatchId
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetCheckColumnCount() *int64 {
	return s.CheckColumnCount
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetCheckPtCount() *int64 {
	return s.CheckPtCount
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetCheckResult() *int32 {
	return s.CheckResult
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetCheckRowCount() *int64 {
	return s.CheckRowCount
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetCheckRowPassCount() *int64 {
	return s.CheckRowPassCount
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetCheckRowPassExport() *string {
	return s.CheckRowPassExport
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetCheckSqlNum() *int64 {
	return s.CheckSqlNum
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetCheckTableNum() *int64 {
	return s.CheckTableNum
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetCheckTemplateId() *string {
	return s.CheckTemplateId
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetCheckTemplateName() *string {
	return s.CheckTemplateName
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetDstDsName() *string {
	return s.DstDsName
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetDstDsType() *string {
	return s.DstDsType
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetErrorTableNum() *int64 {
	return s.ErrorTableNum
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetPassColumnCount() *int64 {
	return s.PassColumnCount
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetPassColumnRate() *float64 {
	return s.PassColumnRate
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetPassProcess() *float64 {
	return s.PassProcess
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetPassProcessExport() *string {
	return s.PassProcessExport
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetPassPtNum() *int64 {
	return s.PassPtNum
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetPassPtProcessExport() *string {
	return s.PassPtProcessExport
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetPassTableNum() *int64 {
	return s.PassTableNum
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetPtPassProcess() *float64 {
	return s.PtPassProcess
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetReportGenerateMessage() *string {
	return s.ReportGenerateMessage
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetReportStatus() *int32 {
	return s.ReportStatus
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetReportTime() *string {
	return s.ReportTime
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetReportTitle() *string {
	return s.ReportTitle
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetSkipPtNum() *int64 {
	return s.SkipPtNum
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetSkipTableNum() *int32 {
	return s.SkipTableNum
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetSrcDsName() *string {
	return s.SrcDsName
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetSrcDsType() *string {
	return s.SrcDsType
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetTaskCreateTime() *string {
	return s.TaskCreateTime
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetTaskModifyTime() *string {
	return s.TaskModifyTime
}

func (s *GetDataCheckReportOverviewResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetBatchId(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.BatchId = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetCheckColumnCount(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.CheckColumnCount = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetCheckPtCount(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.CheckPtCount = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetCheckResult(v int32) *GetDataCheckReportOverviewResponseBodyData {
	s.CheckResult = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetCheckRowCount(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.CheckRowCount = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetCheckRowPassCount(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.CheckRowPassCount = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetCheckRowPassExport(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.CheckRowPassExport = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetCheckSqlNum(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.CheckSqlNum = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetCheckTableNum(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.CheckTableNum = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetCheckTemplateId(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.CheckTemplateId = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetCheckTemplateName(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.CheckTemplateName = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetCheckType(v int32) *GetDataCheckReportOverviewResponseBodyData {
	s.CheckType = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetDstDsName(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.DstDsName = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetDstDsType(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.DstDsType = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetErrorTableNum(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.ErrorTableNum = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetPassColumnCount(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.PassColumnCount = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetPassColumnRate(v float64) *GetDataCheckReportOverviewResponseBodyData {
	s.PassColumnRate = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetPassProcess(v float64) *GetDataCheckReportOverviewResponseBodyData {
	s.PassProcess = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetPassProcessExport(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.PassProcessExport = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetPassPtNum(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.PassPtNum = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetPassPtProcessExport(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.PassPtProcessExport = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetPassTableNum(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.PassTableNum = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetPtPassProcess(v float64) *GetDataCheckReportOverviewResponseBodyData {
	s.PtPassProcess = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetReportGenerateMessage(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.ReportGenerateMessage = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetReportStatus(v int32) *GetDataCheckReportOverviewResponseBodyData {
	s.ReportStatus = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetReportTime(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.ReportTime = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetReportTitle(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.ReportTitle = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetSkipPtNum(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.SkipPtNum = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetSkipTableNum(v int32) *GetDataCheckReportOverviewResponseBodyData {
	s.SkipTableNum = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetSrcDsName(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.SrcDsName = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetSrcDsType(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.SrcDsType = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetTaskCreateTime(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.TaskCreateTime = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetTaskId(v int64) *GetDataCheckReportOverviewResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetTaskModifyTime(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.TaskModifyTime = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) SetTaskName(v string) *GetDataCheckReportOverviewResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *GetDataCheckReportOverviewResponseBodyData) Validate() error {
	return dara.Validate(s)
}
