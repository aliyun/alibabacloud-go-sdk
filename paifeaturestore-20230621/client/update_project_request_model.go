// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateProjectRequest
	GetDescription() *string
	SetName(v string) *UpdateProjectRequest
	GetName() *string
}

type UpdateProjectRequest struct {
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// project1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProjectRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectRequest) SetName(v string) *UpdateProjectRequest {
	s.Name = &v
	return s
}

func (s *UpdateProjectRequest) Validate() error {
	return dara.Validate(s)
}
