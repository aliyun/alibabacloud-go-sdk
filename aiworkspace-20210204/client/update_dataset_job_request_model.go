// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetVersion(v string) *UpdateDatasetJobRequest
	GetDatasetVersion() *string
	SetDescription(v string) *UpdateDatasetJobRequest
	GetDescription() *string
	SetWorkspaceId(v string) *UpdateDatasetJobRequest
	GetWorkspaceId() *string
}

type UpdateDatasetJobRequest struct {
	// The name of the dataset version.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The description of the dataset job.
	//
	// example:
	//
	// This is a job description of a dataset.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the workspace. For more information about how to obtain a workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDatasetJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetJobRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *UpdateDatasetJobRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDatasetJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDatasetJobRequest) SetDatasetVersion(v string) *UpdateDatasetJobRequest {
	s.DatasetVersion = &v
	return s
}

func (s *UpdateDatasetJobRequest) SetDescription(v string) *UpdateDatasetJobRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasetJobRequest) SetWorkspaceId(v string) *UpdateDatasetJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDatasetJobRequest) Validate() error {
	return dara.Validate(s)
}
