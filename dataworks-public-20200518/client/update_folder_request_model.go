// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *UpdateFolderRequest
	GetFolderId() *string
	SetFolderName(v string) *UpdateFolderRequest
	GetFolderName() *string
	SetProjectId(v int64) *UpdateFolderRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *UpdateFolderRequest
	GetProjectIdentifier() *string
}

type UpdateFolderRequest struct {
	// The ID of the folder. You can call the [ListFolders](https://help.aliyun.com/document_detail/173955.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2735c2c19d58
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The name of the folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySecondFolder
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
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

func (s UpdateFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFolderRequest) GoString() string {
	return s.String()
}

func (s *UpdateFolderRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *UpdateFolderRequest) GetFolderName() *string {
	return s.FolderName
}

func (s *UpdateFolderRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateFolderRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *UpdateFolderRequest) SetFolderId(v string) *UpdateFolderRequest {
	s.FolderId = &v
	return s
}

func (s *UpdateFolderRequest) SetFolderName(v string) *UpdateFolderRequest {
	s.FolderName = &v
	return s
}

func (s *UpdateFolderRequest) SetProjectId(v int64) *UpdateFolderRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateFolderRequest) SetProjectIdentifier(v string) *UpdateFolderRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *UpdateFolderRequest) Validate() error {
	return dara.Validate(s)
}
