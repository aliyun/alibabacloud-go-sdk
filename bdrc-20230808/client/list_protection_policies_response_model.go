// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectionPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProtectionPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProtectionPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListProtectionPoliciesResponseBody) *ListProtectionPoliciesResponse
	GetBody() *ListProtectionPoliciesResponseBody
}

type ListProtectionPoliciesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProtectionPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProtectionPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListProtectionPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProtectionPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProtectionPoliciesResponse) GetBody() *ListProtectionPoliciesResponseBody {
	return s.Body
}

func (s *ListProtectionPoliciesResponse) SetHeaders(v map[string]*string) *ListProtectionPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListProtectionPoliciesResponse) SetStatusCode(v int32) *ListProtectionPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProtectionPoliciesResponse) SetBody(v *ListProtectionPoliciesResponseBody) *ListProtectionPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListProtectionPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
