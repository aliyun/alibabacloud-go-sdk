// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoScalingRuleResponseBody) *GetAutoScalingRuleResponse
	GetBody() *GetAutoScalingRuleResponseBody
}

type GetAutoScalingRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *GetAutoScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoScalingRuleResponse) GetBody() *GetAutoScalingRuleResponseBody {
	return s.Body
}

func (s *GetAutoScalingRuleResponse) SetHeaders(v map[string]*string) *GetAutoScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *GetAutoScalingRuleResponse) SetStatusCode(v int32) *GetAutoScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoScalingRuleResponse) SetBody(v *GetAutoScalingRuleResponseBody) *GetAutoScalingRuleResponse {
	s.Body = v
	return s
}

func (s *GetAutoScalingRuleResponse) Validate() error {
	return dara.Validate(s)
}
