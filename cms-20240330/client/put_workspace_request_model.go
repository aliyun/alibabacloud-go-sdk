// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PutWorkspaceRequest
	GetDescription() *string
	SetDisplayName(v string) *PutWorkspaceRequest
	GetDisplayName() *string
	SetSlsProject(v string) *PutWorkspaceRequest
	GetSlsProject() *string
}

type PutWorkspaceRequest struct {
	// 工作空间描述
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// workspace-test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 工作空间绑定的日志服务项目名称
	//
	// This parameter is required.
	//
	// example:
	//
	// sls-project-test-001
	SlsProject *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
}

func (s PutWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s PutWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *PutWorkspaceRequest) GetDescription() *string {
	return s.Description
}

func (s *PutWorkspaceRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *PutWorkspaceRequest) GetSlsProject() *string {
	return s.SlsProject
}

func (s *PutWorkspaceRequest) SetDescription(v string) *PutWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *PutWorkspaceRequest) SetDisplayName(v string) *PutWorkspaceRequest {
	s.DisplayName = &v
	return s
}

func (s *PutWorkspaceRequest) SetSlsProject(v string) *PutWorkspaceRequest {
	s.SlsProject = &v
	return s
}

func (s *PutWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
