// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetVersion(v string) *CreateDatasetJobRequest
	GetDatasetVersion() *string
	SetDescription(v string) *CreateDatasetJobRequest
	GetDescription() *string
	SetJobAction(v string) *CreateDatasetJobRequest
	GetJobAction() *string
	SetJobMode(v string) *CreateDatasetJobRequest
	GetJobMode() *string
	SetJobSpec(v string) *CreateDatasetJobRequest
	GetJobSpec() *string
	SetWorkspaceId(v string) *CreateDatasetJobRequest
	GetWorkspaceId() *string
}

type CreateDatasetJobRequest struct {
	// The dataset version.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The job description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The job action.
	//
	// Valid values:
	//
	// 	- SemanticIndex
	//
	// 	- IntelligentTag
	//
	// 	- FileMetaExport
	//
	// This parameter is required.
	//
	// example:
	//
	// SemanticIndex
	JobAction *string `json:"JobAction,omitempty" xml:"JobAction,omitempty"`
	// The job mode.
	//
	// Valid values:
	//
	// 	- Full: full mode.
	//
	// example:
	//
	// Full
	JobMode *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
	// The job configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"modelId\\":\\"xxx\\"}
	JobSpec *string `json:"JobSpec,omitempty" xml:"JobSpec,omitempty"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDatasetJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetJobRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *CreateDatasetJobRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDatasetJobRequest) GetJobAction() *string {
	return s.JobAction
}

func (s *CreateDatasetJobRequest) GetJobMode() *string {
	return s.JobMode
}

func (s *CreateDatasetJobRequest) GetJobSpec() *string {
	return s.JobSpec
}

func (s *CreateDatasetJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDatasetJobRequest) SetDatasetVersion(v string) *CreateDatasetJobRequest {
	s.DatasetVersion = &v
	return s
}

func (s *CreateDatasetJobRequest) SetDescription(v string) *CreateDatasetJobRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetJobRequest) SetJobAction(v string) *CreateDatasetJobRequest {
	s.JobAction = &v
	return s
}

func (s *CreateDatasetJobRequest) SetJobMode(v string) *CreateDatasetJobRequest {
	s.JobMode = &v
	return s
}

func (s *CreateDatasetJobRequest) SetJobSpec(v string) *CreateDatasetJobRequest {
	s.JobSpec = &v
	return s
}

func (s *CreateDatasetJobRequest) SetWorkspaceId(v string) *CreateDatasetJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDatasetJobRequest) Validate() error {
	return dara.Validate(s)
}
