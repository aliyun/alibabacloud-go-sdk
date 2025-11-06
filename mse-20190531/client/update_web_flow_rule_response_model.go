// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebFlowRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWebFlowRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWebFlowRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWebFlowRuleResponseBody) *UpdateWebFlowRuleResponse
	GetBody() *UpdateWebFlowRuleResponseBody
}

type UpdateWebFlowRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWebFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWebFlowRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebFlowRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWebFlowRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWebFlowRuleResponse) GetBody() *UpdateWebFlowRuleResponseBody {
	return s.Body
}

func (s *UpdateWebFlowRuleResponse) SetHeaders(v map[string]*string) *UpdateWebFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebFlowRuleResponse) SetStatusCode(v int32) *UpdateWebFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebFlowRuleResponse) SetBody(v *UpdateWebFlowRuleResponseBody) *UpdateWebFlowRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateWebFlowRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
