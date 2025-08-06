// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPoliciesForIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppPoliciesForIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppPoliciesForIdentityResponse
	GetStatusCode() *int32
	SetBody(v *ListAppPoliciesForIdentityResponseBody) *ListAppPoliciesForIdentityResponse
	GetBody() *ListAppPoliciesForIdentityResponseBody
}

type ListAppPoliciesForIdentityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppPoliciesForIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppPoliciesForIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppPoliciesForIdentityResponse) GoString() string {
	return s.String()
}

func (s *ListAppPoliciesForIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppPoliciesForIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppPoliciesForIdentityResponse) GetBody() *ListAppPoliciesForIdentityResponseBody {
	return s.Body
}

func (s *ListAppPoliciesForIdentityResponse) SetHeaders(v map[string]*string) *ListAppPoliciesForIdentityResponse {
	s.Headers = v
	return s
}

func (s *ListAppPoliciesForIdentityResponse) SetStatusCode(v int32) *ListAppPoliciesForIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponse) SetBody(v *ListAppPoliciesForIdentityResponseBody) *ListAppPoliciesForIdentityResponse {
	s.Body = v
	return s
}

func (s *ListAppPoliciesForIdentityResponse) Validate() error {
	return dara.Validate(s)
}
