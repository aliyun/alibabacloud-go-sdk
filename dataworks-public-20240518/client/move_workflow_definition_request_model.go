// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveWorkflowDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *MoveWorkflowDefinitionRequest
	GetId() *int64
	SetPath(v string) *MoveWorkflowDefinitionRequest
	GetPath() *string
	SetProjectId(v int64) *MoveWorkflowDefinitionRequest
	GetProjectId() *int64
}

type MoveWorkflowDefinitionRequest struct {
	// The ID of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The path to which you want to move the workflow. You do not need to specify a workflow name in the path.
	//
	// For example, if you want to move the test workflow to root/demo/test, you must set this parameter to root/demo.
	//
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID. This parameter indicates the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveWorkflowDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *MoveWorkflowDefinitionRequest) GetId() *int64 {
	return s.Id
}

func (s *MoveWorkflowDefinitionRequest) GetPath() *string {
	return s.Path
}

func (s *MoveWorkflowDefinitionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *MoveWorkflowDefinitionRequest) SetId(v int64) *MoveWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *MoveWorkflowDefinitionRequest) SetPath(v string) *MoveWorkflowDefinitionRequest {
	s.Path = &v
	return s
}

func (s *MoveWorkflowDefinitionRequest) SetProjectId(v int64) *MoveWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *MoveWorkflowDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
