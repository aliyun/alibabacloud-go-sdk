// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetFileMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetVersion(v string) *GetDatasetFileMetaRequest
	GetDatasetVersion() *string
	SetWorkspaceId(v string) *GetDatasetFileMetaRequest
	GetWorkspaceId() *string
}

type GetDatasetFileMetaRequest struct {
	// The dataset version.
	//
	// example:
	//
	// v1
	DatasetVersion *string `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasetFileMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetFileMetaRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetFileMetaRequest) GetDatasetVersion() *string {
	return s.DatasetVersion
}

func (s *GetDatasetFileMetaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDatasetFileMetaRequest) SetDatasetVersion(v string) *GetDatasetFileMetaRequest {
	s.DatasetVersion = &v
	return s
}

func (s *GetDatasetFileMetaRequest) SetWorkspaceId(v string) *GetDatasetFileMetaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetDatasetFileMetaRequest) Validate() error {
	return dara.Validate(s)
}
