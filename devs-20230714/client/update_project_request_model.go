// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Project) *UpdateProjectRequest
	GetBody() *Project
}

type UpdateProjectRequest struct {
	Body *Project `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) GetBody() *Project {
	return s.Body
}

func (s *UpdateProjectRequest) SetBody(v *Project) *UpdateProjectRequest {
	s.Body = v
	return s
}

func (s *UpdateProjectRequest) Validate() error {
	return dara.Validate(s)
}
