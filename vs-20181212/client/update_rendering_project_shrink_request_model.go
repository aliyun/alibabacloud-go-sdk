// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRenderingProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateRenderingProjectShrinkRequest
	GetDescription() *string
	SetProjectId(v string) *UpdateRenderingProjectShrinkRequest
	GetProjectId() *string
	SetProjectName(v string) *UpdateRenderingProjectShrinkRequest
	GetProjectName() *string
	SetSessionAttribsShrink(v string) *UpdateRenderingProjectShrinkRequest
	GetSessionAttribsShrink() *string
}

type UpdateRenderingProjectShrinkRequest struct {
	// Project description
	//
	// example:
	//
	// 目录1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Project ID
	//
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Default naming rules:
	//
	// 1. Length 1-128
	//
	// 2. Lowercase letters, numbers, underscores (_), hyphens (-), and periods (.).
	//
	// 3. The first and last characters must be letters or digits. At least one of ProjectName, SessionAttribs, or Description must be specified.
	//
	// example:
	//
	// idata_content
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Session attributes
	SessionAttribsShrink *string `json:"SessionAttribs,omitempty" xml:"SessionAttribs,omitempty"`
}

func (s UpdateRenderingProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRenderingProjectShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRenderingProjectShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateRenderingProjectShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateRenderingProjectShrinkRequest) GetSessionAttribsShrink() *string {
	return s.SessionAttribsShrink
}

func (s *UpdateRenderingProjectShrinkRequest) SetDescription(v string) *UpdateRenderingProjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateRenderingProjectShrinkRequest) SetProjectId(v string) *UpdateRenderingProjectShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateRenderingProjectShrinkRequest) SetProjectName(v string) *UpdateRenderingProjectShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateRenderingProjectShrinkRequest) SetSessionAttribsShrink(v string) *UpdateRenderingProjectShrinkRequest {
	s.SessionAttribsShrink = &v
	return s
}

func (s *UpdateRenderingProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
