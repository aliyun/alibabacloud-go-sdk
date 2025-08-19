// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateControlPolicyResponseBody) *UpdateControlPolicyResponse
	GetBody() *UpdateControlPolicyResponseBody
}

type UpdateControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateControlPolicyResponse) GetBody() *UpdateControlPolicyResponseBody {
	return s.Body
}

func (s *UpdateControlPolicyResponse) SetHeaders(v map[string]*string) *UpdateControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateControlPolicyResponse) SetStatusCode(v int32) *UpdateControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateControlPolicyResponse) SetBody(v *UpdateControlPolicyResponseBody) *UpdateControlPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateControlPolicyResponse) Validate() error {
	return dara.Validate(s)
}
