// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetContextDatabaseSyncJobResponseBody
	GetCode() *string
	SetData(v *GetContextDatabaseSyncJobResponseBodyData) *GetContextDatabaseSyncJobResponseBody
	GetData() *GetContextDatabaseSyncJobResponseBodyData
	SetHttpStatusCode(v int64) *GetContextDatabaseSyncJobResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetContextDatabaseSyncJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContextDatabaseSyncJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetContextDatabaseSyncJobResponseBody
	GetSuccess() *bool
}

type GetContextDatabaseSyncJobResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Sample value
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetContextDatabaseSyncJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetContextDatabaseSyncJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetContextDatabaseSyncJobResponseBody) GetData() *GetContextDatabaseSyncJobResponseBodyData {
	return s.Data
}

func (s *GetContextDatabaseSyncJobResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetContextDatabaseSyncJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContextDatabaseSyncJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextDatabaseSyncJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetContextDatabaseSyncJobResponseBody) SetCode(v string) *GetContextDatabaseSyncJobResponseBody {
	s.Code = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBody) SetData(v *GetContextDatabaseSyncJobResponseBodyData) *GetContextDatabaseSyncJobResponseBody {
	s.Data = v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBody) SetHttpStatusCode(v int64) *GetContextDatabaseSyncJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBody) SetMessage(v string) *GetContextDatabaseSyncJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBody) SetRequestId(v string) *GetContextDatabaseSyncJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBody) SetSuccess(v bool) *GetContextDatabaseSyncJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextDatabaseSyncJobResponseBodyData struct {
	// The reason for the automatic pause.
	//
	// example:
	//
	// 3 consecutive execution failures
	AutoPauseReason *string `json:"AutoPauseReason,omitempty" xml:"AutoPauseReason,omitempty"`
	// Indicates whether the job is automatically paused.
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
	// Indicates whether the job is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The filter conditions.
	//
	// example:
	//
	// {"include":["/R&D documents/**"],"exclude":["/Archived documents/**"],"excludedDocumentIds":[],"sizeLimit":2147483648,"typeAllowlist":["docx","pdf"]}
	Filters *GetContextDatabaseSyncJobResponseBodyDataFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	// The job ID, which is a system-generated UUID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The target knowledge base ID.
	//
	// example:
	//
	// kb-as1llqmkol****
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The name.
	//
	// example:
	//
	// R&D team document import
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The next execution time.
	//
	// example:
	//
	// 2026-09-23T10:00:00+08:00
	NextRunAt *string `json:"NextRunAt,omitempty" xml:"NextRunAt,omitempty"`
	// The schedule type. Valid values: interval and daily-times.
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

func (s GetContextDatabaseSyncJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobResponseBodyData) GetAutoPauseReason() *string {
	return s.AutoPauseReason
}

func (s *GetContextDatabaseSyncJobResponseBodyData) GetAutoPaused() *bool {
	return s.AutoPaused
}

func (s *GetContextDatabaseSyncJobResponseBodyData) GetConfigVersion() *int32 {
	return s.ConfigVersion
}

func (s *GetContextDatabaseSyncJobResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetContextDatabaseSyncJobResponseBodyData) GetFilters() *GetContextDatabaseSyncJobResponseBodyDataFilters {
	return s.Filters
}

func (s *GetContextDatabaseSyncJobResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetContextDatabaseSyncJobResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetContextDatabaseSyncJobResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetContextDatabaseSyncJobResponseBodyData) GetNextRunAt() *string {
	return s.NextRunAt
}

func (s *GetContextDatabaseSyncJobResponseBodyData) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *GetContextDatabaseSyncJobResponseBodyData) GetSourceId() *string {
	return s.SourceId
}

func (s *GetContextDatabaseSyncJobResponseBodyData) GetTargetId() *string {
	return s.TargetId
}

func (s *GetContextDatabaseSyncJobResponseBodyData) SetAutoPauseReason(v string) *GetContextDatabaseSyncJobResponseBodyData {
	s.AutoPauseReason = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyData) SetAutoPaused(v bool) *GetContextDatabaseSyncJobResponseBodyData {
	s.AutoPaused = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyData) SetConfigVersion(v int32) *GetContextDatabaseSyncJobResponseBodyData {
	s.ConfigVersion = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyData) SetEnabled(v bool) *GetContextDatabaseSyncJobResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyData) SetFilters(v *GetContextDatabaseSyncJobResponseBodyDataFilters) *GetContextDatabaseSyncJobResponseBodyData {
	s.Filters = v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyData) SetId(v string) *GetContextDatabaseSyncJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyData) SetKnowledgeBaseId(v string) *GetContextDatabaseSyncJobResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyData) SetName(v string) *GetContextDatabaseSyncJobResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyData) SetNextRunAt(v string) *GetContextDatabaseSyncJobResponseBodyData {
	s.NextRunAt = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyData) SetScheduleType(v string) *GetContextDatabaseSyncJobResponseBodyData {
	s.ScheduleType = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyData) SetSourceId(v string) *GetContextDatabaseSyncJobResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyData) SetTargetId(v string) *GetContextDatabaseSyncJobResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyData) Validate() error {
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextDatabaseSyncJobResponseBodyDataFilters struct {
	// The list of exclusion rules.
	//
	// example:
	//
	// ["/Archived documents"]
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
	// ["/R&D documents"]
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

func (s GetContextDatabaseSyncJobResponseBodyDataFilters) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobResponseBodyDataFilters) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobResponseBodyDataFilters) GetExclude() []*string {
	return s.Exclude
}

func (s *GetContextDatabaseSyncJobResponseBodyDataFilters) GetExcludedDocumentIds() []*string {
	return s.ExcludedDocumentIds
}

func (s *GetContextDatabaseSyncJobResponseBodyDataFilters) GetInclude() []*string {
	return s.Include
}

func (s *GetContextDatabaseSyncJobResponseBodyDataFilters) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *GetContextDatabaseSyncJobResponseBodyDataFilters) GetTypeAllowlist() []*string {
	return s.TypeAllowlist
}

func (s *GetContextDatabaseSyncJobResponseBodyDataFilters) SetExclude(v []*string) *GetContextDatabaseSyncJobResponseBodyDataFilters {
	s.Exclude = v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyDataFilters) SetExcludedDocumentIds(v []*string) *GetContextDatabaseSyncJobResponseBodyDataFilters {
	s.ExcludedDocumentIds = v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyDataFilters) SetInclude(v []*string) *GetContextDatabaseSyncJobResponseBodyDataFilters {
	s.Include = v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyDataFilters) SetSizeLimit(v int64) *GetContextDatabaseSyncJobResponseBodyDataFilters {
	s.SizeLimit = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyDataFilters) SetTypeAllowlist(v []*string) *GetContextDatabaseSyncJobResponseBodyDataFilters {
	s.TypeAllowlist = v
	return s
}

func (s *GetContextDatabaseSyncJobResponseBodyDataFilters) Validate() error {
	return dara.Validate(s)
}
