// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemSecurityPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSystemSecurityPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSystemSecurityPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListSystemSecurityPoliciesResponseBody) *ListSystemSecurityPoliciesResponse
	GetBody() *ListSystemSecurityPoliciesResponseBody
}

type ListSystemSecurityPoliciesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSystemSecurityPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSystemSecurityPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSystemSecurityPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSystemSecurityPoliciesResponse) GetBody() *ListSystemSecurityPoliciesResponseBody {
	return s.Body
}

func (s *ListSystemSecurityPoliciesResponse) SetHeaders(v map[string]*string) *ListSystemSecurityPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListSystemSecurityPoliciesResponse) SetStatusCode(v int32) *ListSystemSecurityPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponse) SetBody(v *ListSystemSecurityPoliciesResponseBody) *ListSystemSecurityPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListSystemSecurityPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
