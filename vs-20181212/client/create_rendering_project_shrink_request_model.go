// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRenderingProjectShrinkRequest
	GetDescription() *string
	SetProjectName(v string) *CreateRenderingProjectShrinkRequest
	GetProjectName() *string
	SetSessionAttribsShrink(v string) *CreateRenderingProjectShrinkRequest
	GetSessionAttribsShrink() *string
}

type CreateRenderingProjectShrinkRequest struct {
	// The description of the project. The description can be 0 to 255 characters in length.
	//
	// example:
	//
	// 项目概述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The custom name of the project. This name is the unique identifier for the project.
	//
	// The name must meet the following requirements:
	//
	// 1. Be 1 to 128 characters in length.
	//
	// 2. Contain only lowercase letters, digits, underscores (_), hyphens (-), and periods (.).
	//
	// 3. Start and end with a letter or a digit.
	//
	// This parameter is required.
	//
	// example:
	//
	// du_merchant_d
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The session properties.
	SessionAttribsShrink *string `json:"SessionAttribs,omitempty" xml:"SessionAttribs,omitempty"`
}

func (s CreateRenderingProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRenderingProjectShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRenderingProjectShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateRenderingProjectShrinkRequest) GetSessionAttribsShrink() *string {
	return s.SessionAttribsShrink
}

func (s *CreateRenderingProjectShrinkRequest) SetDescription(v string) *CreateRenderingProjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRenderingProjectShrinkRequest) SetProjectName(v string) *CreateRenderingProjectShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateRenderingProjectShrinkRequest) SetSessionAttribsShrink(v string) *CreateRenderingProjectShrinkRequest {
	s.SessionAttribsShrink = &v
	return s
}

func (s *CreateRenderingProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
