// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuthPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuthPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuthPolicyResponseBody) *UpdateAuthPolicyResponse
	GetBody() *UpdateAuthPolicyResponseBody
}

type UpdateAuthPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuthPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuthPolicyResponse) GetBody() *UpdateAuthPolicyResponseBody {
	return s.Body
}

func (s *UpdateAuthPolicyResponse) SetHeaders(v map[string]*string) *UpdateAuthPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthPolicyResponse) SetStatusCode(v int32) *UpdateAuthPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthPolicyResponse) SetBody(v *UpdateAuthPolicyResponseBody) *UpdateAuthPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateAuthPolicyResponse) Validate() error {
	return dara.Validate(s)
}
