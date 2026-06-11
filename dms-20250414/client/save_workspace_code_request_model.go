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
	// The content of the file.
	//
	// This parameter is required.
	//
	// example:
	//
	// Description: \\"Example template, describe instances in some status\\"nFormatVersion: OOS-2019-06-01nTasks:n  - Name: SleepTaskn    Action: ACS::Sleepn    Properties:n      Duration: PT1Mn
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Specifies whether to overwrite the file if it already exists. The default value is `true`.
	//
	// example:
	//
	// True
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// Specifies whether the file is an IaC template.
	//
	// example:
	//
	// false
	Iac *bool `json:"Iac,omitempty" xml:"Iac,omitempty"`
	// The modification time of the file.
	//
	// The time must be in the ISO 8601 format: `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// example:
	//
	// 2026-01-01T10:11:12Z
	Mtime *string `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	// The publishing configuration, specified as a JSON string. The `repos` array specifies the target repository and branch. The `exclude` array specifies the directories to ignore.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"repos":[{"repo":"git@xxxx.git", "branch":"master"}], "exclude":["/.dms", "/username"]}
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Information about the repository.
	//
	// example:
	//
	// computing/ecs
	Repo *string `json:"Repo,omitempty" xml:"Repo,omitempty"`
	// The ID of the workspace.
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
