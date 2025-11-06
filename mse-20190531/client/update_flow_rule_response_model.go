// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFlowRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFlowRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFlowRuleResponseBody) *UpdateFlowRuleResponse
	GetBody() *UpdateFlowRuleResponseBody
}

type UpdateFlowRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFlowRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFlowRuleResponse) GetBody() *UpdateFlowRuleResponseBody {
	return s.Body
}

func (s *UpdateFlowRuleResponse) SetHeaders(v map[string]*string) *UpdateFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowRuleResponse) SetStatusCode(v int32) *UpdateFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowRuleResponse) SetBody(v *UpdateFlowRuleResponseBody) *UpdateFlowRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateFlowRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
