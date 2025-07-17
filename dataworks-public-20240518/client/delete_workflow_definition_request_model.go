// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteWorkflowDefinitionRequest
	GetId() *int64
	SetProjectId(v int64) *DeleteWorkflowDefinitionRequest
	GetProjectId() *int64
}

type DeleteWorkflowDefinitionRequest struct {
	// The ID of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteWorkflowDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowDefinitionRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteWorkflowDefinitionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteWorkflowDefinitionRequest) SetId(v int64) *DeleteWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *DeleteWorkflowDefinitionRequest) SetProjectId(v int64) *DeleteWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteWorkflowDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
