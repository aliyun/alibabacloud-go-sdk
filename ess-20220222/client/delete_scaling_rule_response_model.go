// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScalingRuleResponseBody) *DeleteScalingRuleResponse
	GetBody() *DeleteScalingRuleResponseBody
}

type DeleteScalingRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScalingRuleResponse) GetBody() *DeleteScalingRuleResponseBody {
	return s.Body
}

func (s *DeleteScalingRuleResponse) SetHeaders(v map[string]*string) *DeleteScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteScalingRuleResponse) SetStatusCode(v int32) *DeleteScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScalingRuleResponse) SetBody(v *DeleteScalingRuleResponseBody) *DeleteScalingRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteScalingRuleResponse) Validate() error {
	return dara.Validate(s)
}
