// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddProjectResponse
	GetStatusCode() *int32
	SetBody(v *AddProjectResponseBody) *AddProjectResponse
	GetBody() *AddProjectResponseBody
}

type AddProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s AddProjectResponse) GoString() string {
	return s.String()
}

func (s *AddProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddProjectResponse) GetBody() *AddProjectResponseBody {
	return s.Body
}

func (s *AddProjectResponse) SetHeaders(v map[string]*string) *AddProjectResponse {
	s.Headers = v
	return s
}

func (s *AddProjectResponse) SetStatusCode(v int32) *AddProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *AddProjectResponse) SetBody(v *AddProjectResponseBody) *AddProjectResponse {
	s.Body = v
	return s
}

func (s *AddProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
