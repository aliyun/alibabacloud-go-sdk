// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistrationPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRegistrationPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRegistrationPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRegistrationPoliciesResponseBody) *DeleteRegistrationPoliciesResponse
	GetBody() *DeleteRegistrationPoliciesResponseBody
}

type DeleteRegistrationPoliciesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRegistrationPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRegistrationPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistrationPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRegistrationPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRegistrationPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRegistrationPoliciesResponse) GetBody() *DeleteRegistrationPoliciesResponseBody {
	return s.Body
}

func (s *DeleteRegistrationPoliciesResponse) SetHeaders(v map[string]*string) *DeleteRegistrationPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DeleteRegistrationPoliciesResponse) SetStatusCode(v int32) *DeleteRegistrationPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRegistrationPoliciesResponse) SetBody(v *DeleteRegistrationPoliciesResponseBody) *DeleteRegistrationPoliciesResponse {
	s.Body = v
	return s
}

func (s *DeleteRegistrationPoliciesResponse) Validate() error {
	return dara.Validate(s)
}
