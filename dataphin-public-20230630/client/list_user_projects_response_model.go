// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserProjectsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserProjectsResponseBody) *ListUserProjectsResponse
	GetBody() *ListUserProjectsResponseBody
}

type ListUserProjectsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListUserProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserProjectsResponse) GetBody() *ListUserProjectsResponseBody {
	return s.Body
}

func (s *ListUserProjectsResponse) SetHeaders(v map[string]*string) *ListUserProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListUserProjectsResponse) SetStatusCode(v int32) *ListUserProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserProjectsResponse) SetBody(v *ListUserProjectsResponseBody) *ListUserProjectsResponse {
	s.Body = v
	return s
}

func (s *ListUserProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
