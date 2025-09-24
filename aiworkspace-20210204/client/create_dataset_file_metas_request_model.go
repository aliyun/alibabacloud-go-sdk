// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetFileMetasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetFileMetas(v []*DatasetFileMetaContentCreate) *CreateDatasetFileMetasRequest
	GetDatasetFileMetas() []*DatasetFileMetaContentCreate
	SetDatasetVersion(v string) *CreateDatasetFileMetasRequest
	GetDatasetVersion() *string
	SetWorkspaceId(v string) *CreateDatasetFileMetasRequest
	GetWorkspaceId() *string
}

type CreateDatasetFileMetasRequest struct {
	// The metadata of the file.
	//
	// This parameter is required.
	DatasetFileMetas []*DatasetFileMetaContentCreate `json:"DatasetFileMetas,omitempty" xml:"DatasetFileMetas,omitempty" type:"Repeated"`
	// The dataset version name.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The ID of the workspace to which the dataset belongs. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDatasetFileMetasRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetFileMetasRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetFileMetasRequest) GetDatasetFileMetas() []*DatasetFileMetaContentCreate {
	return s.DatasetFileMetas
}

func (s *CreateDatasetFileMetasRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *CreateDatasetFileMetasRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDatasetFileMetasRequest) SetDatasetFileMetas(v []*DatasetFileMetaContentCreate) *CreateDatasetFileMetasRequest {
	s.DatasetFileMetas = v
	return s
}

func (s *CreateDatasetFileMetasRequest) SetDatasetVersion(v string) *CreateDatasetFileMetasRequest {
	s.DatasetVersion = &v
	return s
}

func (s *CreateDatasetFileMetasRequest) SetWorkspaceId(v string) *CreateDatasetFileMetasRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDatasetFileMetasRequest) Validate() error {
	return dara.Validate(s)
}
