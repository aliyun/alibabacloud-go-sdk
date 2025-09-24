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
	// The dataset version name.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The dataset job description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
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
