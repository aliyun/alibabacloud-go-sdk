// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateProjectRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateProjectRequest
	GetDescription() *string
	SetName(v string) *UpdateProjectRequest
	GetName() *string
}

type UpdateProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProjectRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProjectRequest) SetClientToken(v string) *UpdateProjectRequest {
	s.ClientToken = &v
	return s
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
