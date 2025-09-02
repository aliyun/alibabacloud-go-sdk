// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *DeleteFolderRequest
	GetFolderId() *string
	SetProjectId(v int64) *DeleteFolderRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *DeleteFolderRequest
	GetProjectIdentifier() *string
}

type DeleteFolderRequest struct {
	// The ID of the folder. You can call the [ListFolders](https://help.aliyun.com/document_detail/173955.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2eb6f9****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace Management page to obtain the ID. You must specify either this parameter or ProjectIdentifier to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the DataWorks console and go to the Workspace Settings panel to obtain the name. You must specify either this parameter or ProjectId to determine the DataWorks workspace to which the operation is applied.
	//
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
}

func (s DeleteFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFolderRequest) GoString() string {
	return s.String()
}

func (s *DeleteFolderRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *DeleteFolderRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteFolderRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *DeleteFolderRequest) SetFolderId(v string) *DeleteFolderRequest {
	s.FolderId = &v
	return s
}

func (s *DeleteFolderRequest) SetProjectId(v int64) *DeleteFolderRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFolderRequest) SetProjectIdentifier(v string) *DeleteFolderRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *DeleteFolderRequest) Validate() error {
	return dara.Validate(s)
}
