// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPrivateAccessPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserPrivateAccessPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserPrivateAccessPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserPrivateAccessPoliciesResponseBody) *ListUserPrivateAccessPoliciesResponse
	GetBody() *ListUserPrivateAccessPoliciesResponseBody
}

type ListUserPrivateAccessPoliciesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserPrivateAccessPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserPrivateAccessPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserPrivateAccessPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListUserPrivateAccessPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserPrivateAccessPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserPrivateAccessPoliciesResponse) GetBody() *ListUserPrivateAccessPoliciesResponseBody {
	return s.Body
}

func (s *ListUserPrivateAccessPoliciesResponse) SetHeaders(v map[string]*string) *ListUserPrivateAccessPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponse) SetStatusCode(v int32) *ListUserPrivateAccessPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponse) SetBody(v *ListUserPrivateAccessPoliciesResponseBody) *ListUserPrivateAccessPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponse) Validate() error {
	return dara.Validate(s)
}
