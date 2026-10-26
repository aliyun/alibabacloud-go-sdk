// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsNotForProjectAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectsNotForProjectAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectsNotForProjectAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectsNotForProjectAuthorizationResponseBody) *ListProjectsNotForProjectAuthorizationResponse
	GetBody() *ListProjectsNotForProjectAuthorizationResponseBody
}

type ListProjectsNotForProjectAuthorizationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectsNotForProjectAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectsNotForProjectAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsNotForProjectAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsNotForProjectAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectsNotForProjectAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectsNotForProjectAuthorizationResponse) GetBody() *ListProjectsNotForProjectAuthorizationResponseBody {
	return s.Body
}

func (s *ListProjectsNotForProjectAuthorizationResponse) SetHeaders(v map[string]*string) *ListProjectsNotForProjectAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsNotForProjectAuthorizationResponse) SetStatusCode(v int32) *ListProjectsNotForProjectAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsNotForProjectAuthorizationResponse) SetBody(v *ListProjectsNotForProjectAuthorizationResponseBody) *ListProjectsNotForProjectAuthorizationResponse {
	s.Body = v
	return s
}

func (s *ListProjectsNotForProjectAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
