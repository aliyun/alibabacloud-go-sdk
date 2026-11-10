// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayQuotaRuleSubjectUsageExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody
	GetCode() *string
	SetData(v *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody
	GetData() *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData
	SetMessage(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody
	GetRequestId() *string
}

type GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID of this OpenAPI query call, which is used to troubleshoot call issues. This is different from the request_id of each consumption record in the exported file.
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) GetData() *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	return s.Data
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) SetCode(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) SetData(v *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) SetMessage(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) SetRequestId(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData struct {
	// The error summary when the task fails, which is used to troubleshoot issues such as disabled logs, query failures, exceeded detail limits, processing timeouts, or file generation failures. This field may be omitted when no errors occur.
	//
	// example:
	//
	// log delivery is not enabled
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The gateway ID recorded when the task was created.
	//
	// example:
	//
	// gw-123456
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The number of detail rows written to the exported file, excluding the header. This value is used to check the processing progress. Even if this value equals totalCount, you should determine whether the file is downloadable based on status=Success.
	//
	// example:
	//
	// 100
	ProcessedCount *int32 `json:"processedCount,omitempty" xml:"processedCount,omitempty"`
	// The export conditions and file results determined when the task was created. The time range, format, and metering dimension may be returned even in non-Success states. File and download information is available only after success. Fields with values of 0, false, or empty strings are omitted.
	Result *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// The quota rule ID recorded when the task was created.
	//
	// example:
	//
	// qr-123456
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The current status of the task. Valid values:
	//
	// - Pending: The task is submitted and waiting to be processed.
	//
	// - Running: The task is querying details, generating files, or uploading files.
	//
	// - Success: The file is generated and uploaded. You still need to check whether the file is expired.
	//
	// - Failed: The task failed. For the cause, see errorMessage.
	//
	// Tasks that are not completed within 10 minutes after creation are processed as Failed. After a file expires, the task remains in the Success state.
	//
	// example:
	//
	// Success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The consumer or consumer group ID recorded when the task was created.
	//
	// example:
	//
	// cs-123456
	SubjectId *string `json:"subjectId,omitempty" xml:"subjectId,omitempty"`
	// The subject name recorded when the task was created, which is used to match consumption logs. Subsequent changes to the subject name do not modify the name recorded for this task.
	//
	// example:
	//
	// consumer-name
	SubjectName *string `json:"subjectName,omitempty" xml:"subjectName,omitempty"`
	// The export task ID, which is the same as the taskId request parameter.
	//
	// example:
	//
	// finops-export-abcdef1234567890
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The task type. The value is fixed as FinOpsUsageDetailExport, which indicates the export of quota rule subject consumption details.
	//
	// example:
	//
	// FinOpsUsageDetailExport
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// The total number of detail rows to be exported as counted by backend queries, excluding the CSV header. This value may be 0 before the counting is completed. You cannot determine whether the result is empty or the task is completed based solely on this value.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetProcessedCount() *int32 {
	return s.ProcessedCount
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetResult() *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	return s.Result
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetRuleId() *string {
	return s.RuleId
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetSubjectId() *string {
	return s.SubjectId
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetSubjectName() *string {
	return s.SubjectName
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetErrorMessage(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetGatewayId(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetProcessedCount(v int32) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.ProcessedCount = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetResult(v *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.Result = v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetRuleId(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetStatus(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetSubjectId(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.SubjectId = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetSubjectName(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.SubjectName = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetTaskId(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetTaskType(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) SetTotalCount(v int32) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult struct {
	// The temporary signed download URL of the exported file. It is returned only when status=Success and the file is not expired. It is dynamically generated for each query and is valid for 10 minutes from the generation time. If the link expires but fileExpiredAt is not reached, you can call this operation again to obtain the link. You do not need to pass the internal storage path as a request parameter.
	//
	// example:
	//
	// https://example.oss-cn-hangzhou.aliyuncs.com/export.csv.gz
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// The expiration time of the downloadUrl returned in this response. It is a UNIX timestamp in seconds. It is provided only when a download link is returned and is approximately 10 minutes after the link generation time. It can be updated by querying again and cannot be used to replace fileExpiredAt.
	//
	// example:
	//
	// 1783938800
	DownloadUrlExpiredAt *int64 `json:"downloadUrlExpiredAt,omitempty" xml:"downloadUrlExpiredAt,omitempty"`
	// The start time of the actual export time range. It is a UNIX timestamp in seconds. The actual start time is also limited by the rule creation time, rule reset, or the effective start time of clearing history records. No details are available if the start time is not earlier than the end time.
	//
	// example:
	//
	// 1783931600
	EffectiveWindowStart *int64 `json:"effectiveWindowStart,omitempty" xml:"effectiveWindowStart,omitempty"`
	// Indicates whether the exported file has exceeded the server-side period for providing download links. A value of true indicates that the file has expired, downloadUrl and downloadUrlExpiredAt are not returned, and you need to create a new export task. This field may be omitted when the value is false. To determine whether the file is downloadable, check both status and downloadUrl.
	//
	// example:
	//
	// false
	FileExpired *bool `json:"fileExpired,omitempty" xml:"fileExpired,omitempty"`
	// The time when the file stops providing new download links. It is a UNIX timestamp in seconds and is usually 24 hours after the file is successfully uploaded. After this time, this operation returns fileExpired=true and no longer generates downloadUrl. The task remains in the Success state. This time is different from the downloadUrlExpiredAt of a single signed link.
	//
	// example:
	//
	// 1784021600
	FileExpiredAt *int64 `json:"fileExpiredAt,omitempty" xml:"fileExpiredAt,omitempty"`
	// The size of the successfully generated and uploaded GZIP compressed file, in bytes. This field may be omitted when the task is not successful.
	//
	// example:
	//
	// 1024
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// Specifies whether to filter out requests with failed or zero usage, which is specified when the task is created. A value of true filters out records with 0 usage and retains only records with usage. This field may be omitted when the value is false.
	//
	// example:
	//
	// true
	FilterFailedRequests *bool `json:"filterFailedRequests,omitempty" xml:"filterFailedRequests,omitempty"`
	// The format of the exported file. The value is currently fixed as CSV_GZIP. The downloaded file is a .csv.gz file, which becomes a CSV file after decompression.
	//
	// example:
	//
	// CSV_GZIP
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// The end time of the request time range determined when the task was created. It is a UNIX timestamp in seconds and does not include this moment. If no time range is specified, this value is the task creation time and is not updated when the task is queried.
	//
	// example:
	//
	// 1783935200
	QueryWindowEnd *int64 `json:"queryWindowEnd,omitempty" xml:"queryWindowEnd,omitempty"`
	// The start time of the request time range determined when the task was created. It is a UNIX timestamp in seconds and includes this moment. If only one time parameter is passed, this value may be derived from the rule period. For the actual export start time, see effectiveWindowStart.
	//
	// example:
	//
	// 1783931600
	QueryWindowStart *int64 `json:"queryWindowStart,omitempty" xml:"queryWindowStart,omitempty"`
	// The quota metering dimension determined when the task was created. token indicates token usage, and credit indicates credit usage. The CSV file outputs the corresponding columns based on this dimension. Credit usage retains decimal places.
	//
	// example:
	//
	// token
	QuotaDimension *string `json:"quotaDimension,omitempty" xml:"quotaDimension,omitempty"`
	// The number of successfully exported consumption detail rows, excluding the CSV header. This field may be omitted when the task is not successful or the export result contains 0 rows. If the task is successful but no matching details exist, the file contains only the header.
	//
	// example:
	//
	// 100
	RowCount *int64 `json:"rowCount,omitempty" xml:"rowCount,omitempty"`
}

func (s GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GetDownloadUrlExpiredAt() *int64 {
	return s.DownloadUrlExpiredAt
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GetEffectiveWindowStart() *int64 {
	return s.EffectiveWindowStart
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GetFileExpired() *bool {
	return s.FileExpired
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GetFileExpiredAt() *int64 {
	return s.FileExpiredAt
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GetFilterFailedRequests() *bool {
	return s.FilterFailedRequests
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GetFormat() *string {
	return s.Format
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GetQueryWindowEnd() *int64 {
	return s.QueryWindowEnd
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GetQueryWindowStart() *int64 {
	return s.QueryWindowStart
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GetQuotaDimension() *string {
	return s.QuotaDimension
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) GetRowCount() *int64 {
	return s.RowCount
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) SetDownloadUrl(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	s.DownloadUrl = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) SetDownloadUrlExpiredAt(v int64) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	s.DownloadUrlExpiredAt = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) SetEffectiveWindowStart(v int64) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	s.EffectiveWindowStart = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) SetFileExpired(v bool) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	s.FileExpired = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) SetFileExpiredAt(v int64) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	s.FileExpiredAt = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) SetFileSize(v int64) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	s.FileSize = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) SetFilterFailedRequests(v bool) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	s.FilterFailedRequests = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) SetFormat(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	s.Format = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) SetQueryWindowEnd(v int64) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	s.QueryWindowEnd = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) SetQueryWindowStart(v int64) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	s.QueryWindowStart = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) SetQuotaDimension(v string) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	s.QuotaDimension = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) SetRowCount(v int64) *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult {
	s.RowCount = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageExportTaskResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
