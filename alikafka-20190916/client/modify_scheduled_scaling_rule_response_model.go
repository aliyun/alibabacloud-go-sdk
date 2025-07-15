// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyScheduledScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyScheduledScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyScheduledScalingRuleResponseBody) *ModifyScheduledScalingRuleResponse
	GetBody() *ModifyScheduledScalingRuleResponseBody
}

type ModifyScheduledScalingRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScheduledScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScheduledScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyScheduledScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyScheduledScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyScheduledScalingRuleResponse) GetBody() *ModifyScheduledScalingRuleResponseBody {
	return s.Body
}

func (s *ModifyScheduledScalingRuleResponse) SetHeaders(v map[string]*string) *ModifyScheduledScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyScheduledScalingRuleResponse) SetStatusCode(v int32) *ModifyScheduledScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScheduledScalingRuleResponse) SetBody(v *ModifyScheduledScalingRuleResponseBody) *ModifyScheduledScalingRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyScheduledScalingRuleResponse) Validate() error {
	return dara.Validate(s)
}
