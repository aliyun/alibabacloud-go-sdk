// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAutoScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAutoScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateAutoScalingRuleResponseBody) *CreateAutoScalingRuleResponse
	GetBody() *CreateAutoScalingRuleResponseBody
}

type CreateAutoScalingRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutoScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAutoScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAutoScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAutoScalingRuleResponse) GetBody() *CreateAutoScalingRuleResponseBody {
	return s.Body
}

func (s *CreateAutoScalingRuleResponse) SetHeaders(v map[string]*string) *CreateAutoScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoScalingRuleResponse) SetStatusCode(v int32) *CreateAutoScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoScalingRuleResponse) SetBody(v *CreateAutoScalingRuleResponseBody) *CreateAutoScalingRuleResponse {
	s.Body = v
	return s
}

func (s *CreateAutoScalingRuleResponse) Validate() error {
	return dara.Validate(s)
}
