// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntegrationPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntegrationPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListIntegrationPoliciesResponseBody) *ListIntegrationPoliciesResponse
	GetBody() *ListIntegrationPoliciesResponseBody
}

type ListIntegrationPoliciesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntegrationPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntegrationPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntegrationPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntegrationPoliciesResponse) GetBody() *ListIntegrationPoliciesResponseBody {
	return s.Body
}

func (s *ListIntegrationPoliciesResponse) SetHeaders(v map[string]*string) *ListIntegrationPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListIntegrationPoliciesResponse) SetStatusCode(v int32) *ListIntegrationPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegrationPoliciesResponse) SetBody(v *ListIntegrationPoliciesResponseBody) *ListIntegrationPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListIntegrationPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
