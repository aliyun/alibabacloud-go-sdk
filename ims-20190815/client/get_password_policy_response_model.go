// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPasswordPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPasswordPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetPasswordPolicyResponseBody) *GetPasswordPolicyResponse
	GetBody() *GetPasswordPolicyResponseBody
}

type GetPasswordPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPasswordPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPasswordPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPasswordPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPasswordPolicyResponse) GetBody() *GetPasswordPolicyResponseBody {
	return s.Body
}

func (s *GetPasswordPolicyResponse) SetHeaders(v map[string]*string) *GetPasswordPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetPasswordPolicyResponse) SetStatusCode(v int32) *GetPasswordPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPasswordPolicyResponse) SetBody(v *GetPasswordPolicyResponseBody) *GetPasswordPolicyResponse {
	s.Body = v
	return s
}

func (s *GetPasswordPolicyResponse) Validate() error {
	return dara.Validate(s)
}
