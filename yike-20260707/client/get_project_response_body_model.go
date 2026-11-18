// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProject(v *GetProjectResponseBodyProject) *GetProjectResponseBody
	GetProject() *GetProjectResponseBodyProject
	SetRequestId(v string) *GetProjectResponseBody
	GetRequestId() *string
}

type GetProjectResponseBody struct {
	// The project details.
	Project *GetProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The request ID used to locate and troubleshoot the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) GetProject() *GetProjectResponseBodyProject {
	return s.Project
}

func (s *GetProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectResponseBody) SetProject(v *GetProjectResponseBodyProject) *GetProjectResponseBody {
	s.Project = v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) Validate() error {
	if s.Project != nil {
		if err := s.Project.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectResponseBodyProject struct {
	// The permissions of the current caller on the project. For example, Manage indicates management permissions.
	//
	// example:
	//
	// Manage
	Auth *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
	// The URL of the cover image.
	//
	// example:
	//
	// https://example.com/cover.png
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// The time when the project was created, in UTC and ISO 8601 format, ending with Z.
	//
	// example:
	//
	// 2026-01-07T02:21:36Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The username of the user who created the project.
	//
	// example:
	//
	// aliyun_183320****
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// The description of the project.
	//
	// example:
	//
	// Product demo project
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the project.
	//
	// example:
	//
	// pd_061716****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// Product Demo
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// ws_zna577pdximv***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetProjectResponseBodyProject) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyProject) GetAuth() *string {
	return s.Auth
}

func (s *GetProjectResponseBodyProject) GetCoverUrl() *string {
	return s.CoverUrl
}

func (s *GetProjectResponseBodyProject) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProjectResponseBodyProject) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *GetProjectResponseBodyProject) GetDescription() *string {
	return s.Description
}

func (s *GetProjectResponseBodyProject) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetProjectResponseBodyProject) GetTitle() *string {
	return s.Title
}

func (s *GetProjectResponseBodyProject) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetProjectResponseBodyProject) SetAuth(v string) *GetProjectResponseBodyProject {
	s.Auth = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetCoverUrl(v string) *GetProjectResponseBodyProject {
	s.CoverUrl = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetCreateTime(v string) *GetProjectResponseBodyProject {
	s.CreateTime = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetCreateUserName(v string) *GetProjectResponseBodyProject {
	s.CreateUserName = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetDescription(v string) *GetProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetProjectId(v string) *GetProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetTitle(v string) *GetProjectResponseBodyProject {
	s.Title = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetWorkspaceId(v string) *GetProjectResponseBodyProject {
	s.WorkspaceId = &v
	return s
}

func (s *GetProjectResponseBodyProject) Validate() error {
	return dara.Validate(s)
}
