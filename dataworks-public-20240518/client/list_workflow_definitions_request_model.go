// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowDefinitionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListWorkflowDefinitionsRequest
	GetName() *string
	SetOwner(v string) *ListWorkflowDefinitionsRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListWorkflowDefinitionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWorkflowDefinitionsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListWorkflowDefinitionsRequest
	GetProjectId() *int64
	SetType(v string) *ListWorkflowDefinitionsRequest
	GetType() *string
}

type ListWorkflowDefinitionsRequest struct {
	// The name of the workflow. Fuzzy search is supported.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the owner, which is the UID of the workspace administrator account. You can log on to the Alibaba Cloud Management Console and view the account UID on the Security Settings page of Account Management.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Configuration page to obtain the workspace ID.
	//
	// This parameter specifies the DataWorks workspace for the current API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the workflow. This parameter is used as a filter condition. If you do not specify this parameter, the default value is CycleWorkflow. Valid values:
	//
	// - CycleWorkflow
	//
	// - ManualWorkflow
	//
	// example:
	//
	// CycleWorkflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWorkflowDefinitionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowDefinitionsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsRequest) GetName() *string {
	return s.Name
}

func (s *ListWorkflowDefinitionsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListWorkflowDefinitionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkflowDefinitionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkflowDefinitionsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListWorkflowDefinitionsRequest) GetType() *string {
	return s.Type
}

func (s *ListWorkflowDefinitionsRequest) SetName(v string) *ListWorkflowDefinitionsRequest {
	s.Name = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetOwner(v string) *ListWorkflowDefinitionsRequest {
	s.Owner = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetPageNumber(v int32) *ListWorkflowDefinitionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetPageSize(v int32) *ListWorkflowDefinitionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetProjectId(v int64) *ListWorkflowDefinitionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetType(v string) *ListWorkflowDefinitionsRequest {
	s.Type = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) Validate() error {
	return dara.Validate(s)
}
