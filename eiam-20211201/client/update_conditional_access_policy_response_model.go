// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConditionalAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConditionalAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConditionalAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConditionalAccessPolicyResponseBody) *UpdateConditionalAccessPolicyResponse
	GetBody() *UpdateConditionalAccessPolicyResponseBody
}

type UpdateConditionalAccessPolicyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConditionalAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConditionalAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConditionalAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateConditionalAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConditionalAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConditionalAccessPolicyResponse) GetBody() *UpdateConditionalAccessPolicyResponseBody {
	return s.Body
}

func (s *UpdateConditionalAccessPolicyResponse) SetHeaders(v map[string]*string) *UpdateConditionalAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateConditionalAccessPolicyResponse) SetStatusCode(v int32) *UpdateConditionalAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConditionalAccessPolicyResponse) SetBody(v *UpdateConditionalAccessPolicyResponseBody) *UpdateConditionalAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateConditionalAccessPolicyResponse) Validate() error {
	return dara.Validate(s)
}
