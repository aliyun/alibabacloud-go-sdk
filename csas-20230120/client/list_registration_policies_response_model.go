// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistrationPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRegistrationPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRegistrationPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListRegistrationPoliciesResponseBody) *ListRegistrationPoliciesResponse
	GetBody() *ListRegistrationPoliciesResponseBody
}

type ListRegistrationPoliciesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegistrationPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegistrationPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRegistrationPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRegistrationPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRegistrationPoliciesResponse) GetBody() *ListRegistrationPoliciesResponseBody {
	return s.Body
}

func (s *ListRegistrationPoliciesResponse) SetHeaders(v map[string]*string) *ListRegistrationPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListRegistrationPoliciesResponse) SetStatusCode(v int32) *ListRegistrationPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegistrationPoliciesResponse) SetBody(v *ListRegistrationPoliciesResponseBody) *ListRegistrationPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListRegistrationPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
