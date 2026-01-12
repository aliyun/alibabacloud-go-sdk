// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBuild(v bool) *UpdateProjectRequest
	GetAutoBuild() *bool
	SetExt(v map[string]interface{}) *UpdateProjectRequest
	GetExt() map[string]interface{}
	SetIntro(v string) *UpdateProjectRequest
	GetIntro() *string
	SetJwtToken(v string) *UpdateProjectRequest
	GetJwtToken() *string
	SetProjectId(v string) *UpdateProjectRequest
	GetProjectId() *string
	SetTitle(v string) *UpdateProjectRequest
	GetTitle() *string
}

type UpdateProjectRequest struct {
	AutoBuild *bool                  `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	Ext       map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Intro     *string                `json:"Intro,omitempty" xml:"Intro,omitempty"`
	JwtToken  *string                `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
	ProjectId *string                `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Title     *string                `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) GetAutoBuild() *bool {
	return s.AutoBuild
}

func (s *UpdateProjectRequest) GetExt() map[string]interface{} {
	return s.Ext
}

func (s *UpdateProjectRequest) GetIntro() *string {
	return s.Intro
}

func (s *UpdateProjectRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *UpdateProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateProjectRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateProjectRequest) SetAutoBuild(v bool) *UpdateProjectRequest {
	s.AutoBuild = &v
	return s
}

func (s *UpdateProjectRequest) SetExt(v map[string]interface{}) *UpdateProjectRequest {
	s.Ext = v
	return s
}

func (s *UpdateProjectRequest) SetIntro(v string) *UpdateProjectRequest {
	s.Intro = &v
	return s
}

func (s *UpdateProjectRequest) SetJwtToken(v string) *UpdateProjectRequest {
	s.JwtToken = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectId(v string) *UpdateProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectRequest) SetTitle(v string) *UpdateProjectRequest {
	s.Title = &v
	return s
}

func (s *UpdateProjectRequest) Validate() error {
	return dara.Validate(s)
}
