// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *AddControlPolicyResponseBody) *AddControlPolicyResponse
	GetBody() *AddControlPolicyResponseBody
}

type AddControlPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s AddControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *AddControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddControlPolicyResponse) GetBody() *AddControlPolicyResponseBody {
	return s.Body
}

func (s *AddControlPolicyResponse) SetHeaders(v map[string]*string) *AddControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *AddControlPolicyResponse) SetStatusCode(v int32) *AddControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AddControlPolicyResponse) SetBody(v *AddControlPolicyResponseBody) *AddControlPolicyResponse {
	s.Body = v
	return s
}

func (s *AddControlPolicyResponse) Validate() error {
	return dara.Validate(s)
}
