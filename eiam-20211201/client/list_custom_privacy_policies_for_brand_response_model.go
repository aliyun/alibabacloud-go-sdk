// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomPrivacyPoliciesForBrandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomPrivacyPoliciesForBrandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomPrivacyPoliciesForBrandResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomPrivacyPoliciesForBrandResponseBody) *ListCustomPrivacyPoliciesForBrandResponse
	GetBody() *ListCustomPrivacyPoliciesForBrandResponseBody
}

type ListCustomPrivacyPoliciesForBrandResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomPrivacyPoliciesForBrandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomPrivacyPoliciesForBrandResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPrivacyPoliciesForBrandResponse) GoString() string {
	return s.String()
}

func (s *ListCustomPrivacyPoliciesForBrandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomPrivacyPoliciesForBrandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomPrivacyPoliciesForBrandResponse) GetBody() *ListCustomPrivacyPoliciesForBrandResponseBody {
	return s.Body
}

func (s *ListCustomPrivacyPoliciesForBrandResponse) SetHeaders(v map[string]*string) *ListCustomPrivacyPoliciesForBrandResponse {
	s.Headers = v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandResponse) SetStatusCode(v int32) *ListCustomPrivacyPoliciesForBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandResponse) SetBody(v *ListCustomPrivacyPoliciesForBrandResponseBody) *ListCustomPrivacyPoliciesForBrandResponse {
	s.Body = v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
