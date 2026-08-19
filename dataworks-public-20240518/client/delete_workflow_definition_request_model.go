// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteWorkflowDefinitionRequest
	GetId() *string
	SetProjectId(v int64) *DeleteWorkflowDefinitionRequest
	GetProjectId() *int64
}

type DeleteWorkflowDefinitionRequest struct {
	// The unique identifier of the data development workflow.
	//
	// 	Notice: This field was of the Long type in SDK versions earlier than 8.0.0 and is of the String type in SDK 8.0.0 and later. **This change does not affect normal SDK usage, and the parameter is still returned in the type defined in the SDK**. Only when you upgrade the SDK across version 8.0.0, the type change may cause project compilation failures, and you need to manually correct the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace configuration page to obtain the workspace ID.
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

func (s *DeleteWorkflowDefinitionRequest) GetId() *string {
	return s.Id
}

func (s *DeleteWorkflowDefinitionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteWorkflowDefinitionRequest) SetId(v string) *DeleteWorkflowDefinitionRequest {
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
