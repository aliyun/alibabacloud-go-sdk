// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecurityPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecurityPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListSecurityPoliciesResponseBody) *ListSecurityPoliciesResponse
	GetBody() *ListSecurityPoliciesResponseBody
}

type ListSecurityPoliciesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecurityPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecurityPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecurityPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecurityPoliciesResponse) GetBody() *ListSecurityPoliciesResponseBody {
	return s.Body
}

func (s *ListSecurityPoliciesResponse) SetHeaders(v map[string]*string) *ListSecurityPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityPoliciesResponse) SetStatusCode(v int32) *ListSecurityPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecurityPoliciesResponse) SetBody(v *ListSecurityPoliciesResponseBody) *ListSecurityPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListSecurityPoliciesResponse) Validate() error {
	return dara.Validate(s)
}
