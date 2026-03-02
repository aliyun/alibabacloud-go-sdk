// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *GetFolderRequest
	GetFolderId() *string
	SetRootType(v string) *GetFolderRequest
	GetRootType() *string
}

type GetFolderRequest struct {
	// example:
	//
	// 89097
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// example:
	//
	// DEVELOPMENT
	RootType *string `json:"rootType,omitempty" xml:"rootType,omitempty"`
}

func (s GetFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFolderRequest) GoString() string {
	return s.String()
}

func (s *GetFolderRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *GetFolderRequest) GetRootType() *string {
	return s.RootType
}

func (s *GetFolderRequest) SetFolderId(v string) *GetFolderRequest {
	s.FolderId = &v
	return s
}

func (s *GetFolderRequest) SetRootType(v string) *GetFolderRequest {
	s.RootType = &v
	return s
}

func (s *GetFolderRequest) Validate() error {
	return dara.Validate(s)
}
