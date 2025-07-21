// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTerminalPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTerminalPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTerminalPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTerminalPolicyResponseBody) *UpdateTerminalPolicyResponse
	GetBody() *UpdateTerminalPolicyResponseBody
}

type UpdateTerminalPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTerminalPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTerminalPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTerminalPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateTerminalPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTerminalPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTerminalPolicyResponse) GetBody() *UpdateTerminalPolicyResponseBody {
	return s.Body
}

func (s *UpdateTerminalPolicyResponse) SetHeaders(v map[string]*string) *UpdateTerminalPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateTerminalPolicyResponse) SetStatusCode(v int32) *UpdateTerminalPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTerminalPolicyResponse) SetBody(v *UpdateTerminalPolicyResponseBody) *UpdateTerminalPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateTerminalPolicyResponse) Validate() error {
	return dara.Validate(s)
}
