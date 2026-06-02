// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v string) *DeleteWorkspaceCodeRequest
	GetPath() *string
	SetRepo(v string) *DeleteWorkspaceCodeRequest
	GetRepo() *string
	SetSymlink(v bool) *DeleteWorkspaceCodeRequest
	GetSymlink() *bool
	SetWorkspaceId(v string) *DeleteWorkspaceCodeRequest
	GetWorkspaceId() *string
}

type DeleteWorkspaceCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// /Workspace/code/default/test.ipynb
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// database/adb
	Repo *string `json:"Repo,omitempty" xml:"Repo,omitempty"`
	// example:
	//
	// false
	Symlink *bool `json:"Symlink,omitempty" xml:"Symlink,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteWorkspaceCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceCodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceCodeRequest) GetPath() *string {
	return s.Path
}

func (s *DeleteWorkspaceCodeRequest) GetRepo() *string {
	return s.Repo
}

func (s *DeleteWorkspaceCodeRequest) GetSymlink() *bool {
	return s.Symlink
}

func (s *DeleteWorkspaceCodeRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteWorkspaceCodeRequest) SetPath(v string) *DeleteWorkspaceCodeRequest {
	s.Path = &v
	return s
}

func (s *DeleteWorkspaceCodeRequest) SetRepo(v string) *DeleteWorkspaceCodeRequest {
	s.Repo = &v
	return s
}

func (s *DeleteWorkspaceCodeRequest) SetSymlink(v bool) *DeleteWorkspaceCodeRequest {
	s.Symlink = &v
	return s
}

func (s *DeleteWorkspaceCodeRequest) SetWorkspaceId(v string) *DeleteWorkspaceCodeRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteWorkspaceCodeRequest) Validate() error {
	return dara.Validate(s)
}
