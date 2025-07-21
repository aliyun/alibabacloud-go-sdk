// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetKeyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetKeyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetKeyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *SetKeyPolicyResponseBody) *SetKeyPolicyResponse
	GetBody() *SetKeyPolicyResponseBody
}

type SetKeyPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetKeyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetKeyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s SetKeyPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetKeyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetKeyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetKeyPolicyResponse) GetBody() *SetKeyPolicyResponseBody {
	return s.Body
}

func (s *SetKeyPolicyResponse) SetHeaders(v map[string]*string) *SetKeyPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetKeyPolicyResponse) SetStatusCode(v int32) *SetKeyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetKeyPolicyResponse) SetBody(v *SetKeyPolicyResponseBody) *SetKeyPolicyResponse {
	s.Body = v
	return s
}

func (s *SetKeyPolicyResponse) Validate() error {
	return dara.Validate(s)
}
