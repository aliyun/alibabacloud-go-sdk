// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCustomPrivacyPoliciesFromBrandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveCustomPrivacyPoliciesFromBrandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveCustomPrivacyPoliciesFromBrandResponse
	GetStatusCode() *int32
	SetBody(v *RemoveCustomPrivacyPoliciesFromBrandResponseBody) *RemoveCustomPrivacyPoliciesFromBrandResponse
	GetBody() *RemoveCustomPrivacyPoliciesFromBrandResponseBody
}

type RemoveCustomPrivacyPoliciesFromBrandResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveCustomPrivacyPoliciesFromBrandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveCustomPrivacyPoliciesFromBrandResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveCustomPrivacyPoliciesFromBrandResponse) GoString() string {
	return s.String()
}

func (s *RemoveCustomPrivacyPoliciesFromBrandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveCustomPrivacyPoliciesFromBrandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveCustomPrivacyPoliciesFromBrandResponse) GetBody() *RemoveCustomPrivacyPoliciesFromBrandResponseBody {
	return s.Body
}

func (s *RemoveCustomPrivacyPoliciesFromBrandResponse) SetHeaders(v map[string]*string) *RemoveCustomPrivacyPoliciesFromBrandResponse {
	s.Headers = v
	return s
}

func (s *RemoveCustomPrivacyPoliciesFromBrandResponse) SetStatusCode(v int32) *RemoveCustomPrivacyPoliciesFromBrandResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveCustomPrivacyPoliciesFromBrandResponse) SetBody(v *RemoveCustomPrivacyPoliciesFromBrandResponseBody) *RemoveCustomPrivacyPoliciesFromBrandResponse {
	s.Body = v
	return s
}

func (s *RemoveCustomPrivacyPoliciesFromBrandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
