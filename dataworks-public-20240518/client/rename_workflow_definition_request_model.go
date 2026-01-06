// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameWorkflowDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *RenameWorkflowDefinitionRequest
	GetId() *string
	SetName(v string) *RenameWorkflowDefinitionRequest
	GetName() *string
	SetProjectId(v int64) *RenameWorkflowDefinitionRequest
	GetProjectId() *int64
}

type RenameWorkflowDefinitionRequest struct {
	// The unique identifier of the Data Studio workflow.
	//
	// >  This field is of the Long type in SDK versions prior to 8.0.0, and of the String type in SDK versions 8.0.0 and later. This change does not affect normal SDK usage; the parameter will still be returned according to the type defined in the SDK. However, compilation failures may occur due to the type change only when upgrading the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the workflow.
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
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

func (s *RenameWorkflowDefinitionRequest) GetId() *string {
	return s.Id
}

func (s *RenameWorkflowDefinitionRequest) GetName() *string {
	return s.Name
}

func (s *RenameWorkflowDefinitionRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RenameWorkflowDefinitionRequest) SetId(v string) *RenameWorkflowDefinitionRequest {
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
