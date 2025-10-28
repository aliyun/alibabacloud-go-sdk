// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableApplicationScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableApplicationScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *DisableApplicationScalingRuleResponseBody) *DisableApplicationScalingRuleResponse
	GetBody() *DisableApplicationScalingRuleResponseBody
}

type DisableApplicationScalingRuleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableApplicationScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableApplicationScalingRuleResponse) GetBody() *DisableApplicationScalingRuleResponseBody {
	return s.Body
}

func (s *DisableApplicationScalingRuleResponse) SetHeaders(v map[string]*string) *DisableApplicationScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationScalingRuleResponse) SetStatusCode(v int32) *DisableApplicationScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationScalingRuleResponse) SetBody(v *DisableApplicationScalingRuleResponseBody) *DisableApplicationScalingRuleResponse {
	s.Body = v
	return s
}

func (s *DisableApplicationScalingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
