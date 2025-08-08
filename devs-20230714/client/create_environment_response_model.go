// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvironmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnvironmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnvironmentResponse
	GetStatusCode() *int32
	SetBody(v *Environment) *CreateEnvironmentResponse
	GetBody() *Environment
}

type CreateEnvironmentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Environment       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnvironmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnvironmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnvironmentResponse) GetBody() *Environment {
	return s.Body
}

func (s *CreateEnvironmentResponse) SetHeaders(v map[string]*string) *CreateEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvironmentResponse) SetStatusCode(v int32) *CreateEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnvironmentResponse) SetBody(v *Environment) *CreateEnvironmentResponse {
	s.Body = v
	return s
}

func (s *CreateEnvironmentResponse) Validate() error {
	return dara.Validate(s)
}
