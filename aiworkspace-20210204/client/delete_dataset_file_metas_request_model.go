// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetFileMetasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetFileMetaIds(v string) *DeleteDatasetFileMetasRequest
	GetDatasetFileMetaIds() *string
	SetDatasetVersion(v string) *DeleteDatasetFileMetasRequest
	GetDatasetVersion() *string
	SetWorkspaceId(v string) *DeleteDatasetFileMetasRequest
	GetWorkspaceId() *string
}

type DeleteDatasetFileMetasRequest struct {
	// The metadata ID of the dataset file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 07914c9534586e4e7aa6e9dbca5009082df******fd8a0d857b33296c59bf6
	DatasetFileMetaIds *string `json:"DatasetFileMetaIds,omitempty" xml:"DatasetFileMetaIds,omitempty"`
	// The dataset version.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The ID of the workspace to which the dataset belongs. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 132602
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteDatasetFileMetasRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetFileMetasRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetFileMetasRequest) GetDatasetFileMetaIds() *string {
	return s.DatasetFileMetaIds
}

func (s *DeleteDatasetFileMetasRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *DeleteDatasetFileMetasRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteDatasetFileMetasRequest) SetDatasetFileMetaIds(v string) *DeleteDatasetFileMetasRequest {
	s.DatasetFileMetaIds = &v
	return s
}

func (s *DeleteDatasetFileMetasRequest) SetDatasetVersion(v string) *DeleteDatasetFileMetasRequest {
	s.DatasetVersion = &v
	return s
}

func (s *DeleteDatasetFileMetasRequest) SetWorkspaceId(v string) *DeleteDatasetFileMetasRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteDatasetFileMetasRequest) Validate() error {
	return dara.Validate(s)
}
