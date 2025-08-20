// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetStackPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetStackPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetStackPolicyResponse
	GetStatusCode() *int32
	SetBody(v *SetStackPolicyResponseBody) *SetStackPolicyResponse
	GetBody() *SetStackPolicyResponseBody
}

type SetStackPolicyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetStackPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetStackPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s SetStackPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetStackPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetStackPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetStackPolicyResponse) GetBody() *SetStackPolicyResponseBody {
	return s.Body
}

func (s *SetStackPolicyResponse) SetHeaders(v map[string]*string) *SetStackPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetStackPolicyResponse) SetStatusCode(v int32) *SetStackPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetStackPolicyResponse) SetBody(v *SetStackPolicyResponseBody) *SetStackPolicyResponse {
	s.Body = v
	return s
}

func (s *SetStackPolicyResponse) Validate() error {
	return dara.Validate(s)
}
