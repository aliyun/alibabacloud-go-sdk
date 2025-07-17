// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizDate(v int64) *ListWorkflowInstancesRequest
	GetBizDate() *int64
	SetIds(v []*int64) *ListWorkflowInstancesRequest
	GetIds() []*int64
	SetName(v string) *ListWorkflowInstancesRequest
	GetName() *string
	SetOwner(v string) *ListWorkflowInstancesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListWorkflowInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWorkflowInstancesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListWorkflowInstancesRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListWorkflowInstancesRequest
	GetSortBy() *string
	SetType(v string) *ListWorkflowInstancesRequest
	GetType() *string
	SetWorkflowId(v int64) *ListWorkflowInstancesRequest
	GetWorkflowId() *int64
}

type ListWorkflowInstancesRequest struct {
	// The data timestamp. The value of this parameter is 00:00:00 of the day before the scheduling time of the instance. The value is a UNIX timestamp. Unit: milliseconds. Example: 1743350400000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1710239005403
	BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// The IDs of the workflow instances. You can query multiple instances at a time by instance ID.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
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
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The type of the workflow instance. Valid values:
	//
	// 	- Normal
	//
	// 	- Manual
	//
	// 	- SmokeTest
	//
	// 	- SupplementData
	//
	// 	- ManualWorkflow
	//
	// example:
	//
	// Normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the workflow to which the instance belongs.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListWorkflowInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesRequest) GetBizDate() *int64 {
	return s.BizDate
}

func (s *ListWorkflowInstancesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *ListWorkflowInstancesRequest) GetName() *string {
	return s.Name
}

func (s *ListWorkflowInstancesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListWorkflowInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkflowInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkflowInstancesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListWorkflowInstancesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListWorkflowInstancesRequest) GetType() *string {
	return s.Type
}

func (s *ListWorkflowInstancesRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListWorkflowInstancesRequest) SetBizDate(v int64) *ListWorkflowInstancesRequest {
	s.BizDate = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetIds(v []*int64) *ListWorkflowInstancesRequest {
	s.Ids = v
	return s
}

func (s *ListWorkflowInstancesRequest) SetName(v string) *ListWorkflowInstancesRequest {
	s.Name = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetOwner(v string) *ListWorkflowInstancesRequest {
	s.Owner = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetPageNumber(v int32) *ListWorkflowInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetPageSize(v int32) *ListWorkflowInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetProjectId(v int64) *ListWorkflowInstancesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetSortBy(v string) *ListWorkflowInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetType(v string) *ListWorkflowInstancesRequest {
	s.Type = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetWorkflowId(v int64) *ListWorkflowInstancesRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowInstancesRequest) Validate() error {
	return dara.Validate(s)
}
