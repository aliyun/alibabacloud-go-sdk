// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPasswordPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPasswordPolicyResponse
	GetStatusCode() *int32
	SetBody(v *SetPasswordPolicyResponseBody) *SetPasswordPolicyResponse
	GetBody() *SetPasswordPolicyResponseBody
}

type SetPasswordPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPasswordPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPasswordPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPasswordPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPasswordPolicyResponse) GetBody() *SetPasswordPolicyResponseBody {
	return s.Body
}

func (s *SetPasswordPolicyResponse) SetHeaders(v map[string]*string) *SetPasswordPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetPasswordPolicyResponse) SetStatusCode(v int32) *SetPasswordPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPasswordPolicyResponse) SetBody(v *SetPasswordPolicyResponseBody) *SetPasswordPolicyResponse {
	s.Body = v
	return s
}

func (s *SetPasswordPolicyResponse) Validate() error {
	return dara.Validate(s)
}
