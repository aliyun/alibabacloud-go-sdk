// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConditionalAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConditionalAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConditionalAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetConditionalAccessPolicyResponseBody) *GetConditionalAccessPolicyResponse
	GetBody() *GetConditionalAccessPolicyResponseBody
}

type GetConditionalAccessPolicyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConditionalAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConditionalAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConditionalAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetConditionalAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConditionalAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConditionalAccessPolicyResponse) GetBody() *GetConditionalAccessPolicyResponseBody {
	return s.Body
}

func (s *GetConditionalAccessPolicyResponse) SetHeaders(v map[string]*string) *GetConditionalAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetConditionalAccessPolicyResponse) SetStatusCode(v int32) *GetConditionalAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConditionalAccessPolicyResponse) SetBody(v *GetConditionalAccessPolicyResponseBody) *GetConditionalAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *GetConditionalAccessPolicyResponse) Validate() error {
	return dara.Validate(s)
}
