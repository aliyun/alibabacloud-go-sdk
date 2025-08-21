// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRenderingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateRenderingProjectRequest
	GetDescription() *string
	SetProjectId(v string) *UpdateRenderingProjectRequest
	GetProjectId() *string
	SetProjectName(v string) *UpdateRenderingProjectRequest
	GetProjectName() *string
	SetSessionAttribs(v *UpdateRenderingProjectRequestSessionAttribs) *UpdateRenderingProjectRequest
	GetSessionAttribs() *UpdateRenderingProjectRequestSessionAttribs
}

type UpdateRenderingProjectRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// idata_content
	ProjectName    *string                                      `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SessionAttribs *UpdateRenderingProjectRequestSessionAttribs `json:"SessionAttribs,omitempty" xml:"SessionAttribs,omitempty" type:"Struct"`
}

func (s UpdateRenderingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateRenderingProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRenderingProjectRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateRenderingProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateRenderingProjectRequest) GetSessionAttribs() *UpdateRenderingProjectRequestSessionAttribs {
	return s.SessionAttribs
}

func (s *UpdateRenderingProjectRequest) SetDescription(v string) *UpdateRenderingProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateRenderingProjectRequest) SetProjectId(v string) *UpdateRenderingProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateRenderingProjectRequest) SetProjectName(v string) *UpdateRenderingProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *UpdateRenderingProjectRequest) SetSessionAttribs(v *UpdateRenderingProjectRequestSessionAttribs) *UpdateRenderingProjectRequest {
	s.SessionAttribs = v
	return s
}

func (s *UpdateRenderingProjectRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateRenderingProjectRequestSessionAttribs struct {
	// example:
	//
	// Async
	StartMode *string `json:"StartMode,omitempty" xml:"StartMode,omitempty"`
}

func (s UpdateRenderingProjectRequestSessionAttribs) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingProjectRequestSessionAttribs) GoString() string {
	return s.String()
}

func (s *UpdateRenderingProjectRequestSessionAttribs) GetStartMode() *string {
	return s.StartMode
}

func (s *UpdateRenderingProjectRequestSessionAttribs) SetStartMode(v string) *UpdateRenderingProjectRequestSessionAttribs {
	s.StartMode = &v
	return s
}

func (s *UpdateRenderingProjectRequestSessionAttribs) Validate() error {
	return dara.Validate(s)
}
