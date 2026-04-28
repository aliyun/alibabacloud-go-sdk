// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRateLimitPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRateLimitPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRateLimitPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRateLimitPolicyResponseBody) *ModifyRateLimitPolicyResponse
	GetBody() *ModifyRateLimitPolicyResponseBody
}

type ModifyRateLimitPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRateLimitPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRateLimitPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRateLimitPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyRateLimitPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRateLimitPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRateLimitPolicyResponse) GetBody() *ModifyRateLimitPolicyResponseBody {
	return s.Body
}

func (s *ModifyRateLimitPolicyResponse) SetHeaders(v map[string]*string) *ModifyRateLimitPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyRateLimitPolicyResponse) SetStatusCode(v int32) *ModifyRateLimitPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRateLimitPolicyResponse) SetBody(v *ModifyRateLimitPolicyResponseBody) *ModifyRateLimitPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyRateLimitPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
