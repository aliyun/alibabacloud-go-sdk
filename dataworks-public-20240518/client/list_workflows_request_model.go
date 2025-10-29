// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvType(v string) *ListWorkflowsRequest
	GetEnvType() *string
	SetIds(v []*int64) *ListWorkflowsRequest
	GetIds() []*int64
	SetName(v string) *ListWorkflowsRequest
	GetName() *string
	SetOwner(v string) *ListWorkflowsRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListWorkflowsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWorkflowsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListWorkflowsRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListWorkflowsRequest
	GetSortBy() *string
	SetTriggerType(v string) *ListWorkflowsRequest
	GetTriggerType() *string
}

type ListWorkflowsRequest struct {
	// The environment of the workspace. Valid values:
	//
	// 	- Prod: production environment
	//
	// 	- Dev: development environment
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The IDs of the workflows. You can query multiple workflows at a time by workflow ID.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The name of the workflow. Fuzzy match is supported.
	//
	// example:
	//
	// Workflow1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The account ID of the workflow owner.
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
	// The field used for sorting. Fields such as TriggerTime and StartedTime are supported. The value of this parameter is in the Sort field + Sort by (Desc/Asc) format. By default, results are sorted in ascending order. Valid values:
	//
	// 	- ModifyTime (Desc/Asc)
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
	// The trigger type.
	//
	// 	- Scheduler
	//
	// 	- Manual
	//
	// example:
	//
	// Scheduler
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ListWorkflowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowsRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListWorkflowsRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *ListWorkflowsRequest) GetName() *string {
	return s.Name
}

func (s *ListWorkflowsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListWorkflowsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkflowsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkflowsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListWorkflowsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListWorkflowsRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListWorkflowsRequest) SetEnvType(v string) *ListWorkflowsRequest {
	s.EnvType = &v
	return s
}

func (s *ListWorkflowsRequest) SetIds(v []*int64) *ListWorkflowsRequest {
	s.Ids = v
	return s
}

func (s *ListWorkflowsRequest) SetName(v string) *ListWorkflowsRequest {
	s.Name = &v
	return s
}

func (s *ListWorkflowsRequest) SetOwner(v string) *ListWorkflowsRequest {
	s.Owner = &v
	return s
}

func (s *ListWorkflowsRequest) SetPageNumber(v int32) *ListWorkflowsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowsRequest) SetPageSize(v int32) *ListWorkflowsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowsRequest) SetProjectId(v int64) *ListWorkflowsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowsRequest) SetSortBy(v string) *ListWorkflowsRequest {
	s.SortBy = &v
	return s
}

func (s *ListWorkflowsRequest) SetTriggerType(v string) *ListWorkflowsRequest {
	s.TriggerType = &v
	return s
}

func (s *ListWorkflowsRequest) Validate() error {
	return dara.Validate(s)
}
