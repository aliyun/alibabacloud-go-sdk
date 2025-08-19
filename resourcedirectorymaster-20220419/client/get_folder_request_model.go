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
}

type GetFolderRequest struct {
	// The ID of the folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-Jyl5U7****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
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

func (s *GetFolderRequest) SetFolderId(v string) *GetFolderRequest {
	s.FolderId = &v
	return s
}

func (s *GetFolderRequest) Validate() error {
	return dara.Validate(s)
}
