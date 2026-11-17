// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatRateLimitRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNatRateLimitRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNatRateLimitRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNatRateLimitRuleResponseBody) *ModifyNatRateLimitRuleResponse
	GetBody() *ModifyNatRateLimitRuleResponseBody
}

type ModifyNatRateLimitRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNatRateLimitRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNatRateLimitRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatRateLimitRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyNatRateLimitRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNatRateLimitRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNatRateLimitRuleResponse) GetBody() *ModifyNatRateLimitRuleResponseBody {
	return s.Body
}

func (s *ModifyNatRateLimitRuleResponse) SetHeaders(v map[string]*string) *ModifyNatRateLimitRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyNatRateLimitRuleResponse) SetStatusCode(v int32) *ModifyNatRateLimitRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNatRateLimitRuleResponse) SetBody(v *ModifyNatRateLimitRuleResponseBody) *ModifyNatRateLimitRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyNatRateLimitRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
