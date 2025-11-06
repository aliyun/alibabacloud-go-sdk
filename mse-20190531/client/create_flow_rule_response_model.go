// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFlowRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFlowRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateFlowRuleResponseBody) *CreateFlowRuleResponse
	GetBody() *CreateFlowRuleResponseBody
}

type CreateFlowRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFlowRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFlowRuleResponse) GetBody() *CreateFlowRuleResponseBody {
	return s.Body
}

func (s *CreateFlowRuleResponse) SetHeaders(v map[string]*string) *CreateFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowRuleResponse) SetStatusCode(v int32) *CreateFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowRuleResponse) SetBody(v *CreateFlowRuleResponseBody) *CreateFlowRuleResponse {
	s.Body = v
	return s
}

func (s *CreateFlowRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
