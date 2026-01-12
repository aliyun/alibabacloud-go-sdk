// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBuild(v bool) *UpdateProjectShrinkRequest
	GetAutoBuild() *bool
	SetExtShrink(v string) *UpdateProjectShrinkRequest
	GetExtShrink() *string
	SetIntro(v string) *UpdateProjectShrinkRequest
	GetIntro() *string
	SetJwtToken(v string) *UpdateProjectShrinkRequest
	GetJwtToken() *string
	SetProjectId(v string) *UpdateProjectShrinkRequest
	GetProjectId() *string
	SetTitle(v string) *UpdateProjectShrinkRequest
	GetTitle() *string
}

type UpdateProjectShrinkRequest struct {
	AutoBuild *bool   `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	ExtShrink *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Intro     *string `json:"Intro,omitempty" xml:"Intro,omitempty"`
	JwtToken  *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectShrinkRequest) GetAutoBuild() *bool {
	return s.AutoBuild
}

func (s *UpdateProjectShrinkRequest) GetExtShrink() *string {
	return s.ExtShrink
}

func (s *UpdateProjectShrinkRequest) GetIntro() *string {
	return s.Intro
}

func (s *UpdateProjectShrinkRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *UpdateProjectShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateProjectShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateProjectShrinkRequest) SetAutoBuild(v bool) *UpdateProjectShrinkRequest {
	s.AutoBuild = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetExtShrink(v string) *UpdateProjectShrinkRequest {
	s.ExtShrink = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetIntro(v string) *UpdateProjectShrinkRequest {
	s.Intro = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetJwtToken(v string) *UpdateProjectShrinkRequest {
	s.JwtToken = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetProjectId(v string) *UpdateProjectShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectShrinkRequest) SetTitle(v string) *UpdateProjectShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
