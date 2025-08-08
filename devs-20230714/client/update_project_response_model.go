// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProjectResponse
	GetStatusCode() *int32
	SetBody(v *Project) *UpdateProjectResponse
	GetBody() *Project
}

type UpdateProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Project           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProjectResponse) GetBody() *Project {
	return s.Body
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetStatusCode(v int32) *UpdateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectResponse) SetBody(v *Project) *UpdateProjectResponse {
	s.Body = v
	return s
}

func (s *UpdateProjectResponse) Validate() error {
	return dara.Validate(s)
}
