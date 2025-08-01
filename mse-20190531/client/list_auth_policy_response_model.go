// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthPolicyResponseBody) *ListAuthPolicyResponse
	GetBody() *ListAuthPolicyResponseBody
}

type ListAuthPolicyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListAuthPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthPolicyResponse) GetBody() *ListAuthPolicyResponseBody {
	return s.Body
}

func (s *ListAuthPolicyResponse) SetHeaders(v map[string]*string) *ListAuthPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListAuthPolicyResponse) SetStatusCode(v int32) *ListAuthPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthPolicyResponse) SetBody(v *ListAuthPolicyResponseBody) *ListAuthPolicyResponse {
	s.Body = v
	return s
}

func (s *ListAuthPolicyResponse) Validate() error {
	return dara.Validate(s)
}
