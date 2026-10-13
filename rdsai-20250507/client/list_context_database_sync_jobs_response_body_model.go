// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListContextDatabaseSyncJobsResponseBody
	GetCode() *string
	SetData(v *ListContextDatabaseSyncJobsResponseBodyData) *ListContextDatabaseSyncJobsResponseBody
	GetData() *ListContextDatabaseSyncJobsResponseBodyData
	SetHttpStatusCode(v int64) *ListContextDatabaseSyncJobsResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *ListContextDatabaseSyncJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContextDatabaseSyncJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContextDatabaseSyncJobsResponseBody
	GetSuccess() *bool
}

type ListContextDatabaseSyncJobsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Sample value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListContextDatabaseSyncJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 1
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
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

func (s ListContextDatabaseSyncJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListContextDatabaseSyncJobsResponseBody) GetData() *ListContextDatabaseSyncJobsResponseBodyData {
	return s.Data
}

func (s *ListContextDatabaseSyncJobsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListContextDatabaseSyncJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContextDatabaseSyncJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseSyncJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContextDatabaseSyncJobsResponseBody) SetCode(v string) *ListContextDatabaseSyncJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBody) SetData(v *ListContextDatabaseSyncJobsResponseBodyData) *ListContextDatabaseSyncJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBody) SetHttpStatusCode(v int64) *ListContextDatabaseSyncJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBody) SetMessage(v string) *ListContextDatabaseSyncJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBody) SetRequestId(v string) *ListContextDatabaseSyncJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBody) SetSuccess(v bool) *ListContextDatabaseSyncJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListContextDatabaseSyncJobsResponseBodyData struct {
	// The list of import tasks.
	//
	// example:
	//
	// [{"Id":"job-as1llqmkol****", "SourceId":"src-as1llqmkol****", "TargetId":"tg-as1llqmkol****", "KnowledgeBaseId":"kb-as1llqmkol****"}]
	Jobs []*ListContextDatabaseSyncJobsResponseBodyDataJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The page number. The value starts from 1.
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
	// The total number of entries.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListContextDatabaseSyncJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobsResponseBodyData) GetJobs() []*ListContextDatabaseSyncJobsResponseBodyDataJobs {
	return s.Jobs
}

func (s *ListContextDatabaseSyncJobsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListContextDatabaseSyncJobsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContextDatabaseSyncJobsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListContextDatabaseSyncJobsResponseBodyData) SetJobs(v []*ListContextDatabaseSyncJobsResponseBodyDataJobs) *ListContextDatabaseSyncJobsResponseBodyData {
	s.Jobs = v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyData) SetPage(v int32) *ListContextDatabaseSyncJobsResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyData) SetPageSize(v int32) *ListContextDatabaseSyncJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyData) SetTotal(v int32) *ListContextDatabaseSyncJobsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyData) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseSyncJobsResponseBodyDataJobs struct {
	// The reason for automatic pause.
	//
	// example:
	//
	// Failed 3 consecutive times
	AutoPauseReason *string `json:"AutoPauseReason,omitempty" xml:"AutoPauseReason,omitempty"`
	// Indicates whether the task is automatically paused.
	//
	// example:
	//
	// true
	AutoPaused *bool `json:"AutoPaused,omitempty" xml:"AutoPaused,omitempty"`
	// The configuration version number.
	//
	// example:
	//
	// 3
	ConfigVersion *int32 `json:"ConfigVersion,omitempty" xml:"ConfigVersion,omitempty"`
	// Indicates whether the task is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The filter conditions.
	//
	// example:
	//
	// {"include":["/研发文档/**"],"exclude":["/归档文档/**"],"excludedDocumentIds":[],"sizeLimit":2147483648,"typeAllowlist":["docx","pdf"]}
	Filters *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	// The task ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the target knowledge base.
	//
	// example:
	//
	// kb-as1llqmkol****
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The name.
	//
	// example:
	//
	// R&D Team Document Import
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time of the next run.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	NextRunAt *string `json:"NextRunAt,omitempty" xml:"NextRunAt,omitempty"`
	// The scheduling type. Valid values: interval, daily-times.
	//
	// example:
	//
	// daily-times
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The data source ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The target ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s ListContextDatabaseSyncJobsResponseBodyDataJobs) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobsResponseBodyDataJobs) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) GetAutoPauseReason() *string {
	return s.AutoPauseReason
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) GetAutoPaused() *bool {
	return s.AutoPaused
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) GetConfigVersion() *int32 {
	return s.ConfigVersion
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) GetFilters() *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters {
	return s.Filters
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) GetId() *string {
	return s.Id
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) GetName() *string {
	return s.Name
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) GetNextRunAt() *string {
	return s.NextRunAt
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) GetSourceId() *string {
	return s.SourceId
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) GetTargetId() *string {
	return s.TargetId
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) SetAutoPauseReason(v string) *ListContextDatabaseSyncJobsResponseBodyDataJobs {
	s.AutoPauseReason = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) SetAutoPaused(v bool) *ListContextDatabaseSyncJobsResponseBodyDataJobs {
	s.AutoPaused = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) SetConfigVersion(v int32) *ListContextDatabaseSyncJobsResponseBodyDataJobs {
	s.ConfigVersion = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) SetEnabled(v bool) *ListContextDatabaseSyncJobsResponseBodyDataJobs {
	s.Enabled = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) SetFilters(v *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) *ListContextDatabaseSyncJobsResponseBodyDataJobs {
	s.Filters = v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) SetId(v string) *ListContextDatabaseSyncJobsResponseBodyDataJobs {
	s.Id = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) SetKnowledgeBaseId(v string) *ListContextDatabaseSyncJobsResponseBodyDataJobs {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) SetName(v string) *ListContextDatabaseSyncJobsResponseBodyDataJobs {
	s.Name = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) SetNextRunAt(v string) *ListContextDatabaseSyncJobsResponseBodyDataJobs {
	s.NextRunAt = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) SetScheduleType(v string) *ListContextDatabaseSyncJobsResponseBodyDataJobs {
	s.ScheduleType = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) SetSourceId(v string) *ListContextDatabaseSyncJobsResponseBodyDataJobs {
	s.SourceId = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) SetTargetId(v string) *ListContextDatabaseSyncJobsResponseBodyDataJobs {
	s.TargetId = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobs) Validate() error {
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListContextDatabaseSyncJobsResponseBodyDataJobsFilters struct {
	// The list of exclusion rules.
	//
	// example:
	//
	// ["/Archived Documents"]
	Exclude []*string `json:"Exclude,omitempty" xml:"Exclude,omitempty" type:"Repeated"`
	// The list of excluded document IDs.
	//
	// example:
	//
	// ["doc-as1llqmkol****"]
	ExcludedDocumentIds []*string `json:"ExcludedDocumentIds,omitempty" xml:"ExcludedDocumentIds,omitempty" type:"Repeated"`
	// The list of inclusion rules.
	//
	// example:
	//
	// ["/R&D Documents"]
	Include []*string `json:"Include,omitempty" xml:"Include,omitempty" type:"Repeated"`
	// The size limit in bytes.
	//
	// example:
	//
	// 2147483648
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
	// The list of allowed document types.
	//
	// example:
	//
	// ["docx", "pdf"]
	TypeAllowlist []*string `json:"TypeAllowlist,omitempty" xml:"TypeAllowlist,omitempty" type:"Repeated"`
}

func (s ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) GetExclude() []*string {
	return s.Exclude
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) GetExcludedDocumentIds() []*string {
	return s.ExcludedDocumentIds
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) GetInclude() []*string {
	return s.Include
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) GetTypeAllowlist() []*string {
	return s.TypeAllowlist
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) SetExclude(v []*string) *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters {
	s.Exclude = v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) SetExcludedDocumentIds(v []*string) *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters {
	s.ExcludedDocumentIds = v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) SetInclude(v []*string) *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters {
	s.Include = v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) SetSizeLimit(v int64) *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters {
	s.SizeLimit = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) SetTypeAllowlist(v []*string) *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters {
	s.TypeAllowlist = v
	return s
}

func (s *ListContextDatabaseSyncJobsResponseBodyDataJobsFilters) Validate() error {
	return dara.Validate(s)
}
