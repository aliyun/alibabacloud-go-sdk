// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BuildProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BuildProjectResponse
	GetStatusCode() *int32
	SetBody(v *BuildProjectResponseBody) *BuildProjectResponse
	GetBody() *BuildProjectResponseBody
}

type BuildProjectResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuildProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuildProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s BuildProjectResponse) GoString() string {
	return s.String()
}

func (s *BuildProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BuildProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BuildProjectResponse) GetBody() *BuildProjectResponseBody {
	return s.Body
}

func (s *BuildProjectResponse) SetHeaders(v map[string]*string) *BuildProjectResponse {
	s.Headers = v
	return s
}

func (s *BuildProjectResponse) SetStatusCode(v int32) *BuildProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *BuildProjectResponse) SetBody(v *BuildProjectResponseBody) *BuildProjectResponse {
	s.Body = v
	return s
}

func (s *BuildProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
