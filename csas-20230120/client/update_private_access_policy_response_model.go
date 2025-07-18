// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrivateAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrivateAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrivateAccessPolicyResponseBody) *UpdatePrivateAccessPolicyResponse
	GetBody() *UpdatePrivateAccessPolicyResponseBody
}

type UpdatePrivateAccessPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrivateAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrivateAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrivateAccessPolicyResponse) GetBody() *UpdatePrivateAccessPolicyResponseBody {
	return s.Body
}

func (s *UpdatePrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *UpdatePrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivateAccessPolicyResponse) SetStatusCode(v int32) *UpdatePrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivateAccessPolicyResponse) SetBody(v *UpdatePrivateAccessPolicyResponseBody) *UpdatePrivateAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdatePrivateAccessPolicyResponse) Validate() error {
	return dara.Validate(s)
}
