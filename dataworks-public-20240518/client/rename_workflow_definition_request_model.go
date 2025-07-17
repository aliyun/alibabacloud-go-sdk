// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameWorkflowDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *RenameWorkflowDefinitionRequest
	GetId() *int64
	SetName(v string) *RenameWorkflowDefinitionRequest
	GetName() *string
	SetProjectId(v int64) *RenameWorkflowDefinitionRequest
	GetProjectId() *int64
}

type RenameWorkflowDefinitionRequest struct {
	// The ID of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Rename
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID. You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameWorkflowDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *RenameWorkflowDefinitionRequest) GetId() *int64 {
	return s.Id
}

func (s *RenameWorkflowDefinitionRequest) GetName() *string {
	return s.Name
}

func (s *RenameWorkflowDefinitionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RenameWorkflowDefinitionRequest) SetId(v int64) *RenameWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *RenameWorkflowDefinitionRequest) SetName(v string) *RenameWorkflowDefinitionRequest {
	s.Name = &v
	return s
}

func (s *RenameWorkflowDefinitionRequest) SetProjectId(v int64) *RenameWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *RenameWorkflowDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
