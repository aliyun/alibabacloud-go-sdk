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
	// The project description.
	//
	// example:
	//
	// Directory 1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The custom project name, which serves as the unique identifier of the project. Default naming conventions: 1. The length must be 1 to 128 characters. 2. The name can contain lowercase letters, digits, underscores (_), hyphens (-), and periods (.). 3. The first and last characters must be letters or digits. At least one of ProjectName, SessionAttribs, or Description must be specified.
	//
	// example:
	//
	// idata_content
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The session attributes.
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
	if s.SessionAttribs != nil {
		if err := s.SessionAttribs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRenderingProjectRequestSessionAttribs struct {
	// The mode in which the session starts the cloud application service. Valid values:
	//
	// - Async: asynchronous.
	//
	// - Sync: synchronous.
	//
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
