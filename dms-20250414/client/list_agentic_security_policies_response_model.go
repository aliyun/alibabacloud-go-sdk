// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticSecurityPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgenticSecurityPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgenticSecurityPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListAgenticSecurityPoliciesResponseBody) *ListAgenticSecurityPoliciesResponse
	GetBody() *ListAgenticSecurityPoliciesResponseBody
}

type ListAgenticSecurityPoliciesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgenticSecurityPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgenticSecurityPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticSecurityPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListAgenticSecurityPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgenticSecurityPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgenticSecurityPoliciesResponse) GetBody() *ListAgenticSecurityPoliciesResponseBody {
	return s.Body
}

func (s *ListAgenticSecurityPoliciesResponse) SetHeaders(v map[string]*string) *ListAgenticSecurityPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListAgenticSecurityPoliciesResponse) SetStatusCode(v int32) *ListAgenticSecurityPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponse) SetBody(v *ListAgenticSecurityPoliciesResponseBody) *ListAgenticSecurityPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListAgenticSecurityPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
