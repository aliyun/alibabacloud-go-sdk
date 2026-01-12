// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsByDependencyIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectsByDependencyIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectsByDependencyIdResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectsByDependencyIdResponseBody) *ListProjectsByDependencyIdResponse
	GetBody() *ListProjectsByDependencyIdResponseBody
}

type ListProjectsByDependencyIdResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectsByDependencyIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectsByDependencyIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsByDependencyIdResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsByDependencyIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectsByDependencyIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectsByDependencyIdResponse) GetBody() *ListProjectsByDependencyIdResponseBody {
	return s.Body
}

func (s *ListProjectsByDependencyIdResponse) SetHeaders(v map[string]*string) *ListProjectsByDependencyIdResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsByDependencyIdResponse) SetStatusCode(v int32) *ListProjectsByDependencyIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsByDependencyIdResponse) SetBody(v *ListProjectsByDependencyIdResponseBody) *ListProjectsByDependencyIdResponse {
	s.Body = v
	return s
}

func (s *ListProjectsByDependencyIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
