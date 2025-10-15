// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableConditionalAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableConditionalAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableConditionalAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DisableConditionalAccessPolicyResponseBody) *DisableConditionalAccessPolicyResponse
	GetBody() *DisableConditionalAccessPolicyResponseBody
}

type DisableConditionalAccessPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableConditionalAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableConditionalAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableConditionalAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *DisableConditionalAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableConditionalAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableConditionalAccessPolicyResponse) GetBody() *DisableConditionalAccessPolicyResponseBody {
	return s.Body
}

func (s *DisableConditionalAccessPolicyResponse) SetHeaders(v map[string]*string) *DisableConditionalAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *DisableConditionalAccessPolicyResponse) SetStatusCode(v int32) *DisableConditionalAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableConditionalAccessPolicyResponse) SetBody(v *DisableConditionalAccessPolicyResponseBody) *DisableConditionalAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *DisableConditionalAccessPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
