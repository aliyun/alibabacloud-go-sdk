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
	// This parameter is required.
	//
	// example:
	//
	// 2eb6f9****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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
