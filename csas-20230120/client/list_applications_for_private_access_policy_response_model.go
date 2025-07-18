// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForPrivateAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationsForPrivateAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationsForPrivateAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationsForPrivateAccessPolicyResponseBody) *ListApplicationsForPrivateAccessPolicyResponse
	GetBody() *ListApplicationsForPrivateAccessPolicyResponseBody
}

type ListApplicationsForPrivateAccessPolicyResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForPrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForPrivateAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForPrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationsForPrivateAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationsForPrivateAccessPolicyResponse) GetBody() *ListApplicationsForPrivateAccessPolicyResponseBody {
	return s.Body
}

func (s *ListApplicationsForPrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *ListApplicationsForPrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponse) SetStatusCode(v int32) *ListApplicationsForPrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponse) SetBody(v *ListApplicationsForPrivateAccessPolicyResponseBody) *ListApplicationsForPrivateAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponse) Validate() error {
	return dara.Validate(s)
}
