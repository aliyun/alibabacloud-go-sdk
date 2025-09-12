// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAutoScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAutoScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAutoScalingRuleResponseBody) *ModifyAutoScalingRuleResponse
	GetBody() *ModifyAutoScalingRuleResponseBody
}

type ModifyAutoScalingRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAutoScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAutoScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAutoScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAutoScalingRuleResponse) GetBody() *ModifyAutoScalingRuleResponseBody {
	return s.Body
}

func (s *ModifyAutoScalingRuleResponse) SetHeaders(v map[string]*string) *ModifyAutoScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoScalingRuleResponse) SetStatusCode(v int32) *ModifyAutoScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoScalingRuleResponse) SetBody(v *ModifyAutoScalingRuleResponseBody) *ModifyAutoScalingRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyAutoScalingRuleResponse) Validate() error {
	return dara.Validate(s)
}
