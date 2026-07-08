// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRenderingProjectRequest
	GetDescription() *string
	SetProjectName(v string) *CreateRenderingProjectRequest
	GetProjectName() *string
	SetSessionAttribs(v *CreateRenderingProjectRequestSessionAttribs) *CreateRenderingProjectRequest
	GetSessionAttribs() *CreateRenderingProjectRequestSessionAttribs
}

type CreateRenderingProjectRequest struct {
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
	SessionAttribs *CreateRenderingProjectRequestSessionAttribs `json:"SessionAttribs,omitempty" xml:"SessionAttribs,omitempty" type:"Struct"`
}

func (s CreateRenderingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateRenderingProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRenderingProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateRenderingProjectRequest) GetSessionAttribs() *CreateRenderingProjectRequestSessionAttribs {
	return s.SessionAttribs
}

func (s *CreateRenderingProjectRequest) SetDescription(v string) *CreateRenderingProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateRenderingProjectRequest) SetProjectName(v string) *CreateRenderingProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateRenderingProjectRequest) SetSessionAttribs(v *CreateRenderingProjectRequestSessionAttribs) *CreateRenderingProjectRequest {
	s.SessionAttribs = v
	return s
}

func (s *CreateRenderingProjectRequest) Validate() error {
	if s.SessionAttribs != nil {
		if err := s.SessionAttribs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRenderingProjectRequestSessionAttribs struct {
	// The mode to start the cloud application service for the session. Valid values:
	//
	// 1. Async: asynchronous
	//
	// 2. Sync: synchronous
	//
	// example:
	//
	// Async
	StartMode *string `json:"StartMode,omitempty" xml:"StartMode,omitempty"`
}

func (s CreateRenderingProjectRequestSessionAttribs) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingProjectRequestSessionAttribs) GoString() string {
	return s.String()
}

func (s *CreateRenderingProjectRequestSessionAttribs) GetStartMode() *string {
	return s.StartMode
}

func (s *CreateRenderingProjectRequestSessionAttribs) SetStartMode(v string) *CreateRenderingProjectRequestSessionAttribs {
	s.StartMode = &v
	return s
}

func (s *CreateRenderingProjectRequestSessionAttribs) Validate() error {
	return dara.Validate(s)
}
