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
	SetNewFolderName(v string) *UpdateFolderRequest
	GetNewFolderName() *string
}

type UpdateFolderRequest struct {
	// The ID of the folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-u8B321****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The new name of the folder.
	//
	// The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// rdFolder
	NewFolderName *string `json:"NewFolderName,omitempty" xml:"NewFolderName,omitempty"`
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

func (s *UpdateFolderRequest) GetNewFolderName() *string {
	return s.NewFolderName
}

func (s *UpdateFolderRequest) SetFolderId(v string) *UpdateFolderRequest {
	s.FolderId = &v
	return s
}

func (s *UpdateFolderRequest) SetNewFolderName(v string) *UpdateFolderRequest {
	s.NewFolderName = &v
	return s
}

func (s *UpdateFolderRequest) Validate() error {
	return dara.Validate(s)
}
