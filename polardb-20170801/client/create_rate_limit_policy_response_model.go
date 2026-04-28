// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRateLimitPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRateLimitPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRateLimitPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateRateLimitPolicyResponseBody) *CreateRateLimitPolicyResponse
	GetBody() *CreateRateLimitPolicyResponseBody
}

type CreateRateLimitPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRateLimitPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRateLimitPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRateLimitPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateRateLimitPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRateLimitPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRateLimitPolicyResponse) GetBody() *CreateRateLimitPolicyResponseBody {
	return s.Body
}

func (s *CreateRateLimitPolicyResponse) SetHeaders(v map[string]*string) *CreateRateLimitPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateRateLimitPolicyResponse) SetStatusCode(v int32) *CreateRateLimitPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRateLimitPolicyResponse) SetBody(v *CreateRateLimitPolicyResponseBody) *CreateRateLimitPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateRateLimitPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
