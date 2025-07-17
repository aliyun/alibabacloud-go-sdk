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
	// This parameter is required.
	//
	// example:
	//
	// 2735c2c19d58
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MySecondFolder
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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
