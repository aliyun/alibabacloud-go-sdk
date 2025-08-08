// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProjectResponse
	GetStatusCode() *int32
	SetBody(v *Project) *CreateProjectResponse
	GetBody() *Project
}

type CreateProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Project           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProjectResponse) GetBody() *Project {
	return s.Body
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *Project) *CreateProjectResponse {
	s.Body = v
	return s
}

func (s *CreateProjectResponse) Validate() error {
	return dara.Validate(s)
}
