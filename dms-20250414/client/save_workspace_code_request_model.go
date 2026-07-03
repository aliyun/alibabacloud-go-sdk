// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWorkspaceCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SaveWorkspaceCodeRequest
	GetContent() *string
	SetForce(v bool) *SaveWorkspaceCodeRequest
	GetForce() *bool
	SetIac(v bool) *SaveWorkspaceCodeRequest
	GetIac() *bool
	SetMtime(v string) *SaveWorkspaceCodeRequest
	GetMtime() *string
	SetPath(v string) *SaveWorkspaceCodeRequest
	GetPath() *string
	SetRepo(v string) *SaveWorkspaceCodeRequest
	GetRepo() *string
	SetWorkspaceId(v string) *SaveWorkspaceCodeRequest
	GetWorkspaceId() *string
}

type SaveWorkspaceCodeRequest struct {
	// The code content.
	//
	// This parameter is required.
	//
	// example:
	//
	// print(123)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Specifies whether to forcibly overwrite the file regardless of whether it has been modified by others.
	//
	// example:
	//
	// True
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// Specifies whether the file is an infrastructure as code template file. Set this parameter to true for YAML configuration files edited through the visual editor.
	//
	// example:
	//
	// false
	Iac *bool `json:"Iac,omitempty" xml:"Iac,omitempty"`
	// The file modification time. The GetWorkspaceCode operation returns this mtime value. When calling SaveWorkspaceCode, include this mtime to compare against the server-side value. If the mtime does not match, the save fails, indicating that the server-side version has changed.
	//
	// example:
	//
	// 2026-01-01T10:11:12Z
	Mtime *string `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	// The file path to save.
	//
	// This parameter is required.
	//
	// example:
	//
	// /Workspace/code/test.py
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The repository information. Pass this parameter when creating a git repo directory during the save action.
	//
	// example:
	//
	// git@codeup.aliyun.com:test/abc.git
	Repo *string `json:"Repo,omitempty" xml:"Repo,omitempty"`
	// The ID of the workspace to publish (numeric ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SaveWorkspaceCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveWorkspaceCodeRequest) GoString() string {
	return s.String()
}

func (s *SaveWorkspaceCodeRequest) GetContent() *string {
	return s.Content
}

func (s *SaveWorkspaceCodeRequest) GetForce() *bool {
	return s.Force
}

func (s *SaveWorkspaceCodeRequest) GetIac() *bool {
	return s.Iac
}

func (s *SaveWorkspaceCodeRequest) GetMtime() *string {
	return s.Mtime
}

func (s *SaveWorkspaceCodeRequest) GetPath() *string {
	return s.Path
}

func (s *SaveWorkspaceCodeRequest) GetRepo() *string {
	return s.Repo
}

func (s *SaveWorkspaceCodeRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SaveWorkspaceCodeRequest) SetContent(v string) *SaveWorkspaceCodeRequest {
	s.Content = &v
	return s
}

func (s *SaveWorkspaceCodeRequest) SetForce(v bool) *SaveWorkspaceCodeRequest {
	s.Force = &v
	return s
}

func (s *SaveWorkspaceCodeRequest) SetIac(v bool) *SaveWorkspaceCodeRequest {
	s.Iac = &v
	return s
}

func (s *SaveWorkspaceCodeRequest) SetMtime(v string) *SaveWorkspaceCodeRequest {
	s.Mtime = &v
	return s
}

func (s *SaveWorkspaceCodeRequest) SetPath(v string) *SaveWorkspaceCodeRequest {
	s.Path = &v
	return s
}

func (s *SaveWorkspaceCodeRequest) SetRepo(v string) *SaveWorkspaceCodeRequest {
	s.Repo = &v
	return s
}

func (s *SaveWorkspaceCodeRequest) SetWorkspaceId(v string) *SaveWorkspaceCodeRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SaveWorkspaceCodeRequest) Validate() error {
	return dara.Validate(s)
}
