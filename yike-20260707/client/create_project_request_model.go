// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTitle(v string) *CreateProjectRequest
	GetTitle() *string
	SetWorkspaceId(v string) *CreateProjectRequest
	GetWorkspaceId() *string
}

type CreateProjectRequest struct {
	// The name of the project. The name can be up to 50 characters in length and cannot contain single quotation marks, semicolons, less than signs, greater than signs, equal signs, vertical lines, or parentheses.
	//
	// This parameter is required.
	//
	// example:
	//
	// Product Demo
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The ID of the workspace to which the project belongs. If this parameter is not specified, the default workspace of the current account is used.
	//
	// example:
	//
	// ws_zna577pdximv***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateProjectRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateProjectRequest) SetTitle(v string) *CreateProjectRequest {
	s.Title = &v
	return s
}

func (s *CreateProjectRequest) SetWorkspaceId(v string) *CreateProjectRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	return dara.Validate(s)
}
