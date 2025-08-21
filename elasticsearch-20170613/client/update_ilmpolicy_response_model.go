// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateILMPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateILMPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateILMPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateILMPolicyResponseBody) *UpdateILMPolicyResponse
	GetBody() *UpdateILMPolicyResponseBody
}

type UpdateILMPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateILMPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateILMPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateILMPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateILMPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateILMPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateILMPolicyResponse) GetBody() *UpdateILMPolicyResponseBody {
	return s.Body
}

func (s *UpdateILMPolicyResponse) SetHeaders(v map[string]*string) *UpdateILMPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateILMPolicyResponse) SetStatusCode(v int32) *UpdateILMPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateILMPolicyResponse) SetBody(v *UpdateILMPolicyResponseBody) *UpdateILMPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateILMPolicyResponse) Validate() error {
	return dara.Validate(s)
}
