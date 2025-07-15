// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScheduledScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScheduledScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScheduledScalingRuleResponseBody) *DeleteScheduledScalingRuleResponse
	GetBody() *DeleteScheduledScalingRuleResponseBody
}

type DeleteScheduledScalingRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduledScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScheduledScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduledScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScheduledScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScheduledScalingRuleResponse) GetBody() *DeleteScheduledScalingRuleResponseBody {
	return s.Body
}

func (s *DeleteScheduledScalingRuleResponse) SetHeaders(v map[string]*string) *DeleteScheduledScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduledScalingRuleResponse) SetStatusCode(v int32) *DeleteScheduledScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduledScalingRuleResponse) SetBody(v *DeleteScheduledScalingRuleResponseBody) *DeleteScheduledScalingRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteScheduledScalingRuleResponse) Validate() error {
	return dara.Validate(s)
}
