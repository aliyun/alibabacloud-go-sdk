// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContextDatabaseSyncRunsResponseBody
	GetCode() *string
	SetData(v *ListContextDatabaseSyncRunsResponseBodyData) *ListContextDatabaseSyncRunsResponseBody
	GetData() *ListContextDatabaseSyncRunsResponseBodyData
	SetHttpStatusCode(v int64) *ListContextDatabaseSyncRunsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListContextDatabaseSyncRunsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContextDatabaseSyncRunsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContextDatabaseSyncRunsResponseBody
	GetSuccess() *bool
}

type ListContextDatabaseSyncRunsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Sample value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListContextDatabaseSyncRunsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message.
	//
	// example:
	//
	// Sample value
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// id-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListContextDatabaseSyncRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContextDatabaseSyncRunsResponseBody) GetData() *ListContextDatabaseSyncRunsResponseBodyData {
	return s.Data
}

func (s *ListContextDatabaseSyncRunsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListContextDatabaseSyncRunsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContextDatabaseSyncRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseSyncRunsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContextDatabaseSyncRunsResponseBody) SetCode(v string) *ListContextDatabaseSyncRunsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBody) SetData(v *ListContextDatabaseSyncRunsResponseBodyData) *ListContextDatabaseSyncRunsResponseBody {
	s.Data = v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBody) SetHttpStatusCode(v int64) *ListContextDatabaseSyncRunsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBody) SetMessage(v string) *ListContextDatabaseSyncRunsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBody) SetRequestId(v string) *ListContextDatabaseSyncRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBody) SetSuccess(v bool) *ListContextDatabaseSyncRunsResponseBody {
	s.Success = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListContextDatabaseSyncRunsResponseBodyData struct {
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of runs.
	//
	// example:
	//
	// [{"RunId":"run-as1llqmkol****", "JobId":"job-as1llqmkol****", "JobName":"R&D team document import", "KnowledgeBaseId":"kb-as1llqmkol****"}]
	Runs []*ListContextDatabaseSyncRunsResponseBodyDataRuns `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListContextDatabaseSyncRunsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncRunsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncRunsResponseBodyData) GetRuns() []*ListContextDatabaseSyncRunsResponseBodyDataRuns {
	return s.Runs
}

func (s *ListContextDatabaseSyncRunsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListContextDatabaseSyncRunsResponseBodyData) SetPage(v int32) *ListContextDatabaseSyncRunsResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyData) SetPageSize(v int32) *ListContextDatabaseSyncRunsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyData) SetRuns(v []*ListContextDatabaseSyncRunsResponseBodyDataRuns) *ListContextDatabaseSyncRunsResponseBodyData {
	s.Runs = v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyData) SetTotal(v int32) *ListContextDatabaseSyncRunsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyData) Validate() error {
	if s.Runs != nil {
		for _, item := range s.Runs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseSyncRunsResponseBodyDataRuns struct {
	// The number of attempts.
	//
	// example:
	//
	// 1
	Attempt *int32 `json:"Attempt,omitempty" xml:"Attempt,omitempty"`
	// The time when the cancellation was requested.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	CancelRequestedAt *string `json:"CancelRequestedAt,omitempty" xml:"CancelRequestedAt,omitempty"`
	// Specifies whether the run is a dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The end time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	// The configuration version number of the import job.
	//
	// example:
	//
	// 3
	JobConfigVersion *int32 `json:"JobConfigVersion,omitempty" xml:"JobConfigVersion,omitempty"`
	// The job ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the import job.
	//
	// example:
	//
	// R&D team document import
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The ID of the destination knowledge base.
	//
	// example:
	//
	// kb-as1llqmkol****
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The latest error message.
	//
	// example:
	//
	// Data source authorization has expired
	LastError *string `json:"LastError,omitempty" xml:"LastError,omitempty"`
	// The plugin reference.
	//
	// example:
	//
	// dingtalk@1.0.0
	PluginRef *string `json:"PluginRef,omitempty" xml:"PluginRef,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1B2C3D4E-5F6A-4B7C-8D9E-0F1A2B3C4D5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The run ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The scheduled run time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	ScheduledAt *string `json:"ScheduledAt,omitempty" xml:"ScheduledAt,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	StartedAt *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	// The run status. Valid values: queued, running, cancelling, succeeded, partial, failed, cancelled, skipped, and dead.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The summary of the run results.
	//
	// example:
	//
	// {"uploaded":12,"unchanged":3,"skipped":1,"deferred":0,"failed":0,"processing":0}
	Summary *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
	// The trigger method of the run. Valid values: manual, schedule, and retry.
	//
	// example:
	//
	// manual
	Trigger *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
}

func (s ListContextDatabaseSyncRunsResponseBodyDataRuns) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunsResponseBodyDataRuns) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetAttempt() *int32 {
	return s.Attempt
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetCancelRequestedAt() *string {
	return s.CancelRequestedAt
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetJobConfigVersion() *int32 {
	return s.JobConfigVersion
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetJobId() *string {
	return s.JobId
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetJobName() *string {
	return s.JobName
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetLastError() *string {
	return s.LastError
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetPluginRef() *string {
	return s.PluginRef
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetRunId() *string {
	return s.RunId
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetScheduledAt() *string {
	return s.ScheduledAt
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetStartedAt() *string {
	return s.StartedAt
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetStatus() *string {
	return s.Status
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetSummary() *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary {
	return s.Summary
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) GetTrigger() *string {
	return s.Trigger
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetAttempt(v int32) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.Attempt = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetCancelRequestedAt(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.CancelRequestedAt = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetDryRun(v bool) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.DryRun = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetFinishedAt(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.FinishedAt = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetJobConfigVersion(v int32) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.JobConfigVersion = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetJobId(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.JobId = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetJobName(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.JobName = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetKnowledgeBaseId(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetLastError(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.LastError = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetPluginRef(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.PluginRef = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetRequestId(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetRunId(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.RunId = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetScheduledAt(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.ScheduledAt = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetStartedAt(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.StartedAt = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetStatus(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.Status = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetSummary(v *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.Summary = v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) SetTrigger(v string) *ListContextDatabaseSyncRunsResponseBodyDataRuns {
	s.Trigger = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRuns) Validate() error {
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListContextDatabaseSyncRunsResponseBodyDataRunsSummary struct {
	// The number of canceled items.
	//
	// example:
	//
	// 0
	Cancelled *int32 `json:"Cancelled,omitempty" xml:"Cancelled,omitempty"`
	// The number of deferred items.
	//
	// example:
	//
	// 0
	Deferred *int32 `json:"Deferred,omitempty" xml:"Deferred,omitempty"`
	// The number of failed items.
	//
	// example:
	//
	// 0
	Failed *int32 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of items with updated metadata.
	//
	// example:
	//
	// 2
	MetadataUpdated *int32 `json:"MetadataUpdated,omitempty" xml:"MetadataUpdated,omitempty"`
	// The number of items being processed.
	//
	// example:
	//
	// 1
	Processing *int32 `json:"Processing,omitempty" xml:"Processing,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 1
	Skipped *int32 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of unchanged items.
	//
	// example:
	//
	// 3
	Unchanged *int32 `json:"Unchanged,omitempty" xml:"Unchanged,omitempty"`
	// The number of accepted uploads.
	//
	// example:
	//
	// 12
	UploadAccepted *int32 `json:"UploadAccepted,omitempty" xml:"UploadAccepted,omitempty"`
	// The number of uploaded items.
	//
	// example:
	//
	// 12
	Uploaded *int32 `json:"Uploaded,omitempty" xml:"Uploaded,omitempty"`
}

func (s ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) GetCancelled() *int32 {
	return s.Cancelled
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) GetDeferred() *int32 {
	return s.Deferred
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) GetFailed() *int32 {
	return s.Failed
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) GetMetadataUpdated() *int32 {
	return s.MetadataUpdated
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) GetProcessing() *int32 {
	return s.Processing
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) GetSkipped() *int32 {
	return s.Skipped
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) GetUnchanged() *int32 {
	return s.Unchanged
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) GetUploadAccepted() *int32 {
	return s.UploadAccepted
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) GetUploaded() *int32 {
	return s.Uploaded
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) SetCancelled(v int32) *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary {
	s.Cancelled = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) SetDeferred(v int32) *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary {
	s.Deferred = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) SetFailed(v int32) *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary {
	s.Failed = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) SetMetadataUpdated(v int32) *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary {
	s.MetadataUpdated = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) SetProcessing(v int32) *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary {
	s.Processing = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) SetSkipped(v int32) *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary {
	s.Skipped = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) SetUnchanged(v int32) *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary {
	s.Unchanged = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) SetUploadAccepted(v int32) *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary {
	s.UploadAccepted = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) SetUploaded(v int32) *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary {
	s.Uploaded = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponseBodyDataRunsSummary) Validate() error {
	return dara.Validate(s)
}
