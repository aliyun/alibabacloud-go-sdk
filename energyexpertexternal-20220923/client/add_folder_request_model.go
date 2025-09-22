// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderName(v string) *AddFolderRequest
	GetFolderName() *string
	SetParentFolderId(v string) *AddFolderRequest
	GetParentFolderId() *string
}

type AddFolderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// “abc” “1234”
	FolderName *string `json:"folderName,omitempty" xml:"folderName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "0" ：parent folder is root
	//
	// "4b648f6d11344f258df876cbdc24dc1f" ： folderId
	ParentFolderId *string `json:"parentFolderId,omitempty" xml:"parentFolderId,omitempty"`
}

func (s AddFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFolderRequest) GoString() string {
	return s.String()
}

func (s *AddFolderRequest) GetFolderName() *string {
	return s.FolderName
}

func (s *AddFolderRequest) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *AddFolderRequest) SetFolderName(v string) *AddFolderRequest {
	s.FolderName = &v
	return s
}

func (s *AddFolderRequest) SetParentFolderId(v string) *AddFolderRequest {
	s.ParentFolderId = &v
	return s
}

func (s *AddFolderRequest) Validate() error {
	return dara.Validate(s)
}
