// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRateLimitPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRateLimitPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRateLimitPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRateLimitPolicyResponseBody) *DeleteRateLimitPolicyResponse
	GetBody() *DeleteRateLimitPolicyResponseBody
}

type DeleteRateLimitPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRateLimitPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRateLimitPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRateLimitPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteRateLimitPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRateLimitPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRateLimitPolicyResponse) GetBody() *DeleteRateLimitPolicyResponseBody {
	return s.Body
}

func (s *DeleteRateLimitPolicyResponse) SetHeaders(v map[string]*string) *DeleteRateLimitPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteRateLimitPolicyResponse) SetStatusCode(v int32) *DeleteRateLimitPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRateLimitPolicyResponse) SetBody(v *DeleteRateLimitPolicyResponseBody) *DeleteRateLimitPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteRateLimitPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
