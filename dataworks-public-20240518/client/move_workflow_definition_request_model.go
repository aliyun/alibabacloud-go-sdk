// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveWorkflowDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *MoveWorkflowDefinitionRequest
	GetId() *string
	SetPath(v string) *MoveWorkflowDefinitionRequest
	GetPath() *string
	SetProjectId(v int64) *MoveWorkflowDefinitionRequest
	GetProjectId() *int64
}

type MoveWorkflowDefinitionRequest struct {
	// The unique identifier of the Data Studio workflow.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect normal SDK usage; the parameter will still be returned according to the type defined in the SDK. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the workflow.
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
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

func (s *MoveWorkflowDefinitionRequest) GetId() *string {
	return s.Id
}

func (s *MoveWorkflowDefinitionRequest) GetPath() *string {
	return s.Path
}

func (s *MoveWorkflowDefinitionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *MoveWorkflowDefinitionRequest) SetId(v string) *MoveWorkflowDefinitionRequest {
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
