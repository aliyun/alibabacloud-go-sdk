// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAbacPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAbacPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAbacPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAbacPolicyResponseBody) *UpdateAbacPolicyResponse
	GetBody() *UpdateAbacPolicyResponseBody
}

type UpdateAbacPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAbacPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAbacPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAbacPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateAbacPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAbacPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAbacPolicyResponse) GetBody() *UpdateAbacPolicyResponseBody {
	return s.Body
}

func (s *UpdateAbacPolicyResponse) SetHeaders(v map[string]*string) *UpdateAbacPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateAbacPolicyResponse) SetStatusCode(v int32) *UpdateAbacPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAbacPolicyResponse) SetBody(v *UpdateAbacPolicyResponseBody) *UpdateAbacPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateAbacPolicyResponse) Validate() error {
	return dara.Validate(s)
}
