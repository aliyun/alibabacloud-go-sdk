// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationScalingRuleResponseBody) *UpdateApplicationScalingRuleResponse
	GetBody() *UpdateApplicationScalingRuleResponseBody
}

type UpdateApplicationScalingRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationScalingRuleResponse) GetBody() *UpdateApplicationScalingRuleResponseBody {
	return s.Body
}

func (s *UpdateApplicationScalingRuleResponse) SetHeaders(v map[string]*string) *UpdateApplicationScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationScalingRuleResponse) SetStatusCode(v int32) *UpdateApplicationScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponse) SetBody(v *UpdateApplicationScalingRuleResponseBody) *UpdateApplicationScalingRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationScalingRuleResponse) Validate() error {
	return dara.Validate(s)
}
