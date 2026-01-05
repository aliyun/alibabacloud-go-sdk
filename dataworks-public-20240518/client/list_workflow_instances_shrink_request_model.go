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
	// The data timestamp. The value of this parameter is 00:00:00 of the day before the scheduling time of the instance. The value is a UNIX timestamp. Unit: milliseconds. Example: 1743350400000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1710239005403
	BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
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
	// The IDs of the workflow instances. You can query multiple instances at a time by instance ID.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The instance name. Fuzzy match is supported.
	//
	// example:
	//
	// WorkflowInstance1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The account ID of the workflow instance owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
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
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The fields used for sorting. Fields such as TriggerTime and StartedTime are supported. The value of this parameter is in the Sort field + Sort by (Desc/Asc) format. By default, results are sorted in ascending order. Valid values:
	//
	// 	- TriggerTime (Desc/Asc)
	//
	// 	- StartedTime (Desc/Asc)
	//
	// 	- FinishedTime (Desc/Asc)
	//
	// 	- CreateTime (Desc/Asc)
	//
	// 	- Id (Desc/Asc)
	//
	// Default value: Id Desc.
	//
	// example:
	//
	// Id Desc
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the workflow instance. Valid values:
	//
	// 	- Normal: Scheduled execution
	//
	// 	- Manual: Manually triggered node
	//
	// 	- SmokeTest: Smoke test
	//
	// 	- SupplementData: Data backfill
	//
	// 	- ManualWorkflow: Manually triggered workflow
	//
	// 	- TriggerWorkflow: Triggered Workflow
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
