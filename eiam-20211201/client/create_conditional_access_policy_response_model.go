// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConditionalAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConditionalAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConditionalAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateConditionalAccessPolicyResponseBody) *CreateConditionalAccessPolicyResponse
	GetBody() *CreateConditionalAccessPolicyResponseBody
}

type CreateConditionalAccessPolicyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConditionalAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConditionalAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConditionalAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateConditionalAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConditionalAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConditionalAccessPolicyResponse) GetBody() *CreateConditionalAccessPolicyResponseBody {
	return s.Body
}

func (s *CreateConditionalAccessPolicyResponse) SetHeaders(v map[string]*string) *CreateConditionalAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateConditionalAccessPolicyResponse) SetStatusCode(v int32) *CreateConditionalAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConditionalAccessPolicyResponse) SetBody(v *CreateConditionalAccessPolicyResponseBody) *CreateConditionalAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateConditionalAccessPolicyResponse) Validate() error {
	return dara.Validate(s)
}
