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
	// The name of the dataset version.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a job description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The task operation.
	//
	// - SemanticIndex: semantic index
	//
	// - IntelligentTag: intelligent tagging
	//
	// - FileMetaExport: metadata export
	//
	// - FileMetaBuild: build and update metadata
	//
	// - IntelligentTagRevert: revoke intelligent tagging
	//
	// - FileMetaImport: metadata import
	//
	// This parameter is required.
	//
	// example:
	//
	// SemanticIndex
	JobAction *string `json:"JobAction,omitempty" xml:"JobAction,omitempty"`
	// The task type.
	//
	// - Full (default): forces the processing of all metadata. This task takes a long time to execute.
	//
	// - Increment: processes only changed or unsuccessfully processed metadata. The SemanticIndex and IntelligentTag tasks support Increment and Full. Other tasks support only Full.
	//
	// example:
	//
	// Full
	JobMode *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
	// The task details.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"modelId\\":\\"xxx\\"}
	JobSpec *string `json:"JobSpec,omitempty" xml:"JobSpec,omitempty"`
	// The workspace ID. For more information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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
