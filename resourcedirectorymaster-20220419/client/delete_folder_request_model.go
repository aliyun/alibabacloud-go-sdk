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
}

type DeleteFolderRequest struct {
	// The ID of the folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-ae1in7****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
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

func (s *DeleteFolderRequest) SetFolderId(v string) *DeleteFolderRequest {
	s.FolderId = &v
	return s
}

func (s *DeleteFolderRequest) Validate() error {
	return dara.Validate(s)
}
