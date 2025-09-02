// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileId(v int64) *DeleteFileRequest
	GetFileId() *int64
	SetProjectId(v int64) *DeleteFileRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *DeleteFileRequest
	GetProjectIdentifier() *string
}

type DeleteFileRequest struct {
	// The file ID. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to view the workspace name.
	//
	// You must specify one of the ProjectId and ProjectIdentifier parameters to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the file. You can call the [ListFiles](https://help.aliyun.com/document_detail/173942.html) operation to query the ID of the file.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *DeleteFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteFileRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *DeleteFileRequest) SetFileId(v int64) *DeleteFileRequest {
	s.FileId = &v
	return s
}

func (s *DeleteFileRequest) SetProjectId(v int64) *DeleteFileRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFileRequest) SetProjectIdentifier(v string) *DeleteFileRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *DeleteFileRequest) Validate() error {
	return dara.Validate(s)
}
