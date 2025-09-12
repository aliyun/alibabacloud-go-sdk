// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAutoScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAutoScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAutoScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAutoScalingRuleResponseBody) *DeleteAutoScalingRuleResponse
	GetBody() *DeleteAutoScalingRuleResponseBody
}

type DeleteAutoScalingRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutoScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutoScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAutoScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAutoScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAutoScalingRuleResponse) GetBody() *DeleteAutoScalingRuleResponseBody {
	return s.Body
}

func (s *DeleteAutoScalingRuleResponse) SetHeaders(v map[string]*string) *DeleteAutoScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoScalingRuleResponse) SetStatusCode(v int32) *DeleteAutoScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoScalingRuleResponse) SetBody(v *DeleteAutoScalingRuleResponseBody) *DeleteAutoScalingRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteAutoScalingRuleResponse) Validate() error {
	return dara.Validate(s)
}
