// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Project) *CreateProjectRequest
	GetBody() *Project
}

type CreateProjectRequest struct {
	Body *Project `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetBody() *Project {
	return s.Body
}

func (s *CreateProjectRequest) SetBody(v *Project) *CreateProjectRequest {
	s.Body = v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	return dara.Validate(s)
}
