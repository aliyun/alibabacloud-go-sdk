// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvType(v string) *ListWorkflowsShrinkRequest
	GetEnvType() *string
	SetIdsShrink(v string) *ListWorkflowsShrinkRequest
	GetIdsShrink() *string
	SetName(v string) *ListWorkflowsShrinkRequest
	GetName() *string
	SetOwner(v string) *ListWorkflowsShrinkRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListWorkflowsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWorkflowsShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListWorkflowsShrinkRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListWorkflowsShrinkRequest
	GetSortBy() *string
	SetTagsShrink(v string) *ListWorkflowsShrinkRequest
	GetTagsShrink() *string
	SetTriggerType(v string) *ListWorkflowsShrinkRequest
	GetTriggerType() *string
}

type ListWorkflowsShrinkRequest struct {
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
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
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
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s ListWorkflowsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowsShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListWorkflowsShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *ListWorkflowsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListWorkflowsShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListWorkflowsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkflowsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkflowsShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListWorkflowsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListWorkflowsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListWorkflowsShrinkRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListWorkflowsShrinkRequest) SetEnvType(v string) *ListWorkflowsShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *ListWorkflowsShrinkRequest) SetIdsShrink(v string) *ListWorkflowsShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *ListWorkflowsShrinkRequest) SetName(v string) *ListWorkflowsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListWorkflowsShrinkRequest) SetOwner(v string) *ListWorkflowsShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListWorkflowsShrinkRequest) SetPageNumber(v int32) *ListWorkflowsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowsShrinkRequest) SetPageSize(v int32) *ListWorkflowsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowsShrinkRequest) SetProjectId(v int64) *ListWorkflowsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowsShrinkRequest) SetSortBy(v string) *ListWorkflowsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListWorkflowsShrinkRequest) SetTagsShrink(v string) *ListWorkflowsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListWorkflowsShrinkRequest) SetTriggerType(v string) *ListWorkflowsShrinkRequest {
	s.TriggerType = &v
	return s
}

func (s *ListWorkflowsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
