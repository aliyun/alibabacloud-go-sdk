// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProjectsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProjectsResponseBody) *DeleteProjectsResponse
	GetBody() *DeleteProjectsResponseBody
}

type DeleteProjectsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectsResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProjectsResponse) GetBody() *DeleteProjectsResponseBody {
	return s.Body
}

func (s *DeleteProjectsResponse) SetHeaders(v map[string]*string) *DeleteProjectsResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectsResponse) SetStatusCode(v int32) *DeleteProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectsResponse) SetBody(v *DeleteProjectsResponseBody) *DeleteProjectsResponse {
	s.Body = v
	return s
}

func (s *DeleteProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
