// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsForProjectAuthorizationPrincipalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectsForProjectAuthorizationPrincipalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectsForProjectAuthorizationPrincipalResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectsForProjectAuthorizationPrincipalResponseBody) *ListProjectsForProjectAuthorizationPrincipalResponse
	GetBody() *ListProjectsForProjectAuthorizationPrincipalResponseBody
}

type ListProjectsForProjectAuthorizationPrincipalResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectsForProjectAuthorizationPrincipalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectsForProjectAuthorizationPrincipalResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsForProjectAuthorizationPrincipalResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponse) GetBody() *ListProjectsForProjectAuthorizationPrincipalResponseBody {
	return s.Body
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponse) SetHeaders(v map[string]*string) *ListProjectsForProjectAuthorizationPrincipalResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponse) SetStatusCode(v int32) *ListProjectsForProjectAuthorizationPrincipalResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponse) SetBody(v *ListProjectsForProjectAuthorizationPrincipalResponseBody) *ListProjectsForProjectAuthorizationPrincipalResponse {
	s.Body = v
	return s
}

func (s *ListProjectsForProjectAuthorizationPrincipalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
