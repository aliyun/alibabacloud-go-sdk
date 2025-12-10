// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderName(v string) *CreateFolderRequest
	GetFolderName() *string
	SetParentFolderId(v string) *CreateFolderRequest
	GetParentFolderId() *string
}

type CreateFolderRequest struct {
	// The name of the folder.
	//
	// The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.),and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// rdFolder
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// r-b1****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s CreateFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateFolderRequest) GetFolderName() *string {
	return s.FolderName
}

func (s *CreateFolderRequest) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *CreateFolderRequest) SetFolderName(v string) *CreateFolderRequest {
	s.FolderName = &v
	return s
}

func (s *CreateFolderRequest) SetParentFolderId(v string) *CreateFolderRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateFolderRequest) Validate() error {
	return dara.Validate(s)
}
