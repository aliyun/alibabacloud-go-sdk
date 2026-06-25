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
	// A list of file metadata content.
	//
	// This parameter is required.
	DatasetFileMetas []*DatasetFileMetaContentCreate `json:"DatasetFileMetas,omitempty" xml:"DatasetFileMetas,omitempty" type:"Repeated"`
	// The name of the dataset version.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The ID of the workspace where the dataset is located. For more information about how to obtain a workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
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
	if s.DatasetFileMetas != nil {
		for _, item := range s.DatasetFileMetas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
