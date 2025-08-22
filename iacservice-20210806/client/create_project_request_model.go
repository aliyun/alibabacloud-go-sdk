// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateProjectRequest
	GetClientToken() *string
	SetDescription(v string) *CreateProjectRequest
	GetDescription() *string
	SetName(v string) *CreateProjectRequest
	GetName() *string
}

type CreateProjectRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectRequest) SetClientToken(v string) *CreateProjectRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	return dara.Validate(s)
}
