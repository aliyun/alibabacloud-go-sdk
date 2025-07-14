// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationScalingRuleResponseBody) *DeleteApplicationScalingRuleResponse
	GetBody() *DeleteApplicationScalingRuleResponseBody
}

type DeleteApplicationScalingRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationScalingRuleResponse) GetBody() *DeleteApplicationScalingRuleResponseBody {
	return s.Body
}

func (s *DeleteApplicationScalingRuleResponse) SetHeaders(v map[string]*string) *DeleteApplicationScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationScalingRuleResponse) SetStatusCode(v int32) *DeleteApplicationScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationScalingRuleResponse) SetBody(v *DeleteApplicationScalingRuleResponseBody) *DeleteApplicationScalingRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationScalingRuleResponse) Validate() error {
	return dara.Validate(s)
}
