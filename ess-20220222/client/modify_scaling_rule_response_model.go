// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyScalingRuleResponseBody) *ModifyScalingRuleResponse
	GetBody() *ModifyScalingRuleResponseBody
}

type ModifyScalingRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyScalingRuleResponse) GetBody() *ModifyScalingRuleResponseBody {
	return s.Body
}

func (s *ModifyScalingRuleResponse) SetHeaders(v map[string]*string) *ModifyScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyScalingRuleResponse) SetStatusCode(v int32) *ModifyScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScalingRuleResponse) SetBody(v *ModifyScalingRuleResponseBody) *ModifyScalingRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyScalingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
