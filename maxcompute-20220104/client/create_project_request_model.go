// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *CreateProjectRequest
	GetBody() *string
}

type CreateProjectRequest struct {
	// The request body parameters.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetBody() *string {
	return s.Body
}

func (s *CreateProjectRequest) SetBody(v string) *CreateProjectRequest {
	s.Body = &v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	return dara.Validate(s)
}
