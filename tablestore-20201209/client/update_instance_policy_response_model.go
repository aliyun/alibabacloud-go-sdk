// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstancePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstancePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstancePolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstancePolicyResponseBody) *UpdateInstancePolicyResponse
	GetBody() *UpdateInstancePolicyResponseBody
}

type UpdateInstancePolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstancePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstancePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstancePolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstancePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstancePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstancePolicyResponse) GetBody() *UpdateInstancePolicyResponseBody {
	return s.Body
}

func (s *UpdateInstancePolicyResponse) SetHeaders(v map[string]*string) *UpdateInstancePolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstancePolicyResponse) SetStatusCode(v int32) *UpdateInstancePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstancePolicyResponse) SetBody(v *UpdateInstancePolicyResponseBody) *UpdateInstancePolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateInstancePolicyResponse) Validate() error {
	return dara.Validate(s)
}
