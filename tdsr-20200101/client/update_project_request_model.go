// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v string) *UpdateProjectRequest
	GetBusinessId() *string
	SetId(v string) *UpdateProjectRequest
	GetId() *string
	SetName(v string) *UpdateProjectRequest
	GetName() *string
}

type UpdateProjectRequest struct {
	// example:
	//
	// 5432****
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) GetBusinessId() *string {
	return s.BusinessId
}

func (s *UpdateProjectRequest) GetId() *string {
	return s.Id
}

func (s *UpdateProjectRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProjectRequest) SetBusinessId(v string) *UpdateProjectRequest {
	s.BusinessId = &v
	return s
}

func (s *UpdateProjectRequest) SetId(v string) *UpdateProjectRequest {
	s.Id = &v
	return s
}

func (s *UpdateProjectRequest) SetName(v string) *UpdateProjectRequest {
	s.Name = &v
	return s
}

func (s *UpdateProjectRequest) Validate() error {
	return dara.Validate(s)
}
