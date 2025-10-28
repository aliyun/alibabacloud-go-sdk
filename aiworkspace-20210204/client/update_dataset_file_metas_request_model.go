// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetFileMetasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetFileMetas(v []*DatasetFileMetaConentUpdate) *UpdateDatasetFileMetasRequest
	GetDatasetFileMetas() []*DatasetFileMetaConentUpdate
	SetDatasetVersion(v string) *UpdateDatasetFileMetasRequest
	GetDatasetVersion() *string
	SetTagJobId(v string) *UpdateDatasetFileMetasRequest
	GetTagJobId() *string
	SetWorkspaceId(v string) *UpdateDatasetFileMetasRequest
	GetWorkspaceId() *string
}

type UpdateDatasetFileMetasRequest struct {
	// The metadata records to be updated for the dataset files.
	//
	// This parameter is required.
	DatasetFileMetas []*DatasetFileMetaConentUpdate `json:"DatasetFileMetas,omitempty" xml:"DatasetFileMetas,omitempty" type:"Repeated"`
	// The dataset version.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The ID of the tagging job that is associated with the metadata tag of the dataset file.
	//
	// example:
	//
	// dsjob-hv0b1****u8taig3y
	TagJobId *string `json:"TagJobId,omitempty" xml:"TagJobId,omitempty"`
	// The ID of the workspace to which the dataset belongs. To obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 796**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDatasetFileMetasRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetFileMetasRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetFileMetasRequest) GetDatasetFileMetas() []*DatasetFileMetaConentUpdate {
	return s.DatasetFileMetas
}

func (s *UpdateDatasetFileMetasRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *UpdateDatasetFileMetasRequest) GetTagJobId() *string {
	return s.TagJobId
}

func (s *UpdateDatasetFileMetasRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDatasetFileMetasRequest) SetDatasetFileMetas(v []*DatasetFileMetaConentUpdate) *UpdateDatasetFileMetasRequest {
	s.DatasetFileMetas = v
	return s
}

func (s *UpdateDatasetFileMetasRequest) SetDatasetVersion(v string) *UpdateDatasetFileMetasRequest {
	s.DatasetVersion = &v
	return s
}

func (s *UpdateDatasetFileMetasRequest) SetTagJobId(v string) *UpdateDatasetFileMetasRequest {
	s.TagJobId = &v
	return s
}

func (s *UpdateDatasetFileMetasRequest) SetWorkspaceId(v string) *UpdateDatasetFileMetasRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDatasetFileMetasRequest) Validate() error {
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
