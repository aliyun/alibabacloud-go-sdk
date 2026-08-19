// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizDate(v int64) *ListWorkflowInstancesShrinkRequest
	GetBizDate() *int64
	SetEnvType(v string) *ListWorkflowInstancesShrinkRequest
	GetEnvType() *string
	SetFilter(v string) *ListWorkflowInstancesShrinkRequest
	GetFilter() *string
	SetIdsShrink(v string) *ListWorkflowInstancesShrinkRequest
	GetIdsShrink() *string
	SetName(v string) *ListWorkflowInstancesShrinkRequest
	GetName() *string
	SetOwner(v string) *ListWorkflowInstancesShrinkRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListWorkflowInstancesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWorkflowInstancesShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListWorkflowInstancesShrinkRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListWorkflowInstancesShrinkRequest
	GetSortBy() *string
	SetTagsShrink(v string) *ListWorkflowInstancesShrinkRequest
	GetTagsShrink() *string
	SetType(v string) *ListWorkflowInstancesShrinkRequest
	GetType() *string
	SetUnifiedWorkflowInstanceId(v int64) *ListWorkflowInstancesShrinkRequest
	GetUnifiedWorkflowInstanceId() *int64
	SetWorkflowId(v int64) *ListWorkflowInstancesShrinkRequest
	GetWorkflowId() *int64
}

type ListWorkflowInstancesShrinkRequest struct {
	// The business date. This is typically 00:00:00 of the day before the scheduled time of the periodic instance. The value is a millisecond-level timestamp, such as 1743350400000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1710239005403
	BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The project environment. Valid values:
	//
	// - Prod: production
	//
	// - Dev: development
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The filter. The value is in JSON format. Multiple filter conditions have an AND relationship. Currently supported fields: `status, executionDate`.
	//
	// example:
	//
	// {
	//
	//     "status": "Success",
	//
	//     "executionDate": "1763481600000"
	//
	// }
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The list of workflow instance IDs. You can use this parameter to query information about multiple workflow instances in a batch.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The name. Fuzzy match is supported.
	//
	// example:
	//
	// WorkflowInstance1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The account ID of the owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of sort fields. Sorting by scheduled time, start time, and other fields is supported. The format is "sort field + sort order (Desc/Asc)". Asc is the default if omitted. Valid values for the sort field:
	//
	// - TriggerTime (Desc/Asc)
	//
	// - StartedTime (Desc/Asc)
	//
	// - FinishedTime (Desc/Asc)
	//
	// - CreateTime (Desc/Asc)
	//
	// - Id (Desc/Asc)
	//
	// Default value: Id Desc.
	//
	// example:
	//
	// Id Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The list of tags. Results are returned if any one of the specified tags matches.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the workflow instance. Valid values:
	//
	// - Normal: periodic scheduling
	//
	// - Manual: manual task
	//
	// - SmokeTest: test
	//
	// - SupplementData: data backfill
	//
	// - ManualWorkflow: manual workflow
	//
	// - TriggerWorkflow: trigger-based workflow
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The unified workflow instance ID. All workflow instances within the same business date under a specific trigger share the same value for this field.
	//
	// example:
	//
	// 1234
	UnifiedWorkflowInstanceId *int64 `json:"UnifiedWorkflowInstanceId,omitempty" xml:"UnifiedWorkflowInstanceId,omitempty"`
	// The ID of the workflow to which the instance belongs.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListWorkflowInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesShrinkRequest) GetBizDate() *int64 {
	return s.BizDate
}

func (s *ListWorkflowInstancesShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListWorkflowInstancesShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListWorkflowInstancesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *ListWorkflowInstancesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListWorkflowInstancesShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListWorkflowInstancesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkflowInstancesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkflowInstancesShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListWorkflowInstancesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListWorkflowInstancesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListWorkflowInstancesShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListWorkflowInstancesShrinkRequest) GetUnifiedWorkflowInstanceId() *int64 {
	return s.UnifiedWorkflowInstanceId
}

func (s *ListWorkflowInstancesShrinkRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListWorkflowInstancesShrinkRequest) SetBizDate(v int64) *ListWorkflowInstancesShrinkRequest {
	s.BizDate = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetEnvType(v string) *ListWorkflowInstancesShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetFilter(v string) *ListWorkflowInstancesShrinkRequest {
	s.Filter = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetIdsShrink(v string) *ListWorkflowInstancesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetName(v string) *ListWorkflowInstancesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetOwner(v string) *ListWorkflowInstancesShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetPageNumber(v int32) *ListWorkflowInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetPageSize(v int32) *ListWorkflowInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetProjectId(v int64) *ListWorkflowInstancesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetSortBy(v string) *ListWorkflowInstancesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetTagsShrink(v string) *ListWorkflowInstancesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetType(v string) *ListWorkflowInstancesShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetUnifiedWorkflowInstanceId(v int64) *ListWorkflowInstancesShrinkRequest {
	s.UnifiedWorkflowInstanceId = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) SetWorkflowId(v int64) *ListWorkflowInstancesShrinkRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
