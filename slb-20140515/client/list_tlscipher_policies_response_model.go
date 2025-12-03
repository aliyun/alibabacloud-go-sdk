// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTLSCipherPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTLSCipherPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTLSCipherPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListTLSCipherPoliciesResponseBody) *ListTLSCipherPoliciesResponse
	GetBody() *ListTLSCipherPoliciesResponseBody
}

type ListTLSCipherPoliciesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTLSCipherPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTLSCipherPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTLSCipherPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListTLSCipherPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTLSCipherPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTLSCipherPoliciesResponse) GetBody() *ListTLSCipherPoliciesResponseBody {
	return s.Body
}

func (s *ListTLSCipherPoliciesResponse) SetHeaders(v map[string]*string) *ListTLSCipherPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListTLSCipherPoliciesResponse) SetStatusCode(v int32) *ListTLSCipherPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTLSCipherPoliciesResponse) SetBody(v *ListTLSCipherPoliciesResponseBody) *ListTLSCipherPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListTLSCipherPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
