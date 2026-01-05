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
	// The unique identifier of the Data Studio workflow.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect normal SDK usage; the parameter will still be returned according to the type defined in the SDK. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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
