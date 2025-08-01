// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuthPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAuthPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAuthPolicyResponse
	GetStatusCode() *int32
	SetBody(v *AddAuthPolicyResponseBody) *AddAuthPolicyResponse
	GetBody() *AddAuthPolicyResponseBody
}

type AddAuthPolicyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAuthPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAuthPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAuthPolicyResponse) GoString() string {
	return s.String()
}

func (s *AddAuthPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAuthPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAuthPolicyResponse) GetBody() *AddAuthPolicyResponseBody {
	return s.Body
}

func (s *AddAuthPolicyResponse) SetHeaders(v map[string]*string) *AddAuthPolicyResponse {
	s.Headers = v
	return s
}

func (s *AddAuthPolicyResponse) SetStatusCode(v int32) *AddAuthPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAuthPolicyResponse) SetBody(v *AddAuthPolicyResponseBody) *AddAuthPolicyResponse {
	s.Body = v
	return s
}

func (s *AddAuthPolicyResponse) Validate() error {
	return dara.Validate(s)
}
