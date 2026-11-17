// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatRateLimitRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNatRateLimitRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNatRateLimitRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNatRateLimitRuleResponseBody) *DeleteNatRateLimitRuleResponse
	GetBody() *DeleteNatRateLimitRuleResponseBody
}

type DeleteNatRateLimitRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNatRateLimitRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNatRateLimitRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatRateLimitRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteNatRateLimitRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNatRateLimitRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNatRateLimitRuleResponse) GetBody() *DeleteNatRateLimitRuleResponseBody {
	return s.Body
}

func (s *DeleteNatRateLimitRuleResponse) SetHeaders(v map[string]*string) *DeleteNatRateLimitRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteNatRateLimitRuleResponse) SetStatusCode(v int32) *DeleteNatRateLimitRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNatRateLimitRuleResponse) SetBody(v *DeleteNatRateLimitRuleResponseBody) *DeleteNatRateLimitRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteNatRateLimitRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
