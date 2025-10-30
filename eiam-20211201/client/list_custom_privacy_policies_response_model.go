// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomPrivacyPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomPrivacyPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomPrivacyPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomPrivacyPoliciesResponseBody) *ListCustomPrivacyPoliciesResponse
	GetBody() *ListCustomPrivacyPoliciesResponseBody
}

type ListCustomPrivacyPoliciesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomPrivacyPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomPrivacyPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPrivacyPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomPrivacyPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomPrivacyPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomPrivacyPoliciesResponse) GetBody() *ListCustomPrivacyPoliciesResponseBody {
	return s.Body
}

func (s *ListCustomPrivacyPoliciesResponse) SetHeaders(v map[string]*string) *ListCustomPrivacyPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomPrivacyPoliciesResponse) SetStatusCode(v int32) *ListCustomPrivacyPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomPrivacyPoliciesResponse) SetBody(v *ListCustomPrivacyPoliciesResponseBody) *ListCustomPrivacyPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListCustomPrivacyPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
