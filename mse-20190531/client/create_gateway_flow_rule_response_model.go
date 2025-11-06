// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayFlowRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGatewayFlowRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGatewayFlowRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateGatewayFlowRuleResponseBody) *CreateGatewayFlowRuleResponse
	GetBody() *CreateGatewayFlowRuleResponseBody
}

type CreateGatewayFlowRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayFlowRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayFlowRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGatewayFlowRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGatewayFlowRuleResponse) GetBody() *CreateGatewayFlowRuleResponseBody {
	return s.Body
}

func (s *CreateGatewayFlowRuleResponse) SetHeaders(v map[string]*string) *CreateGatewayFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayFlowRuleResponse) SetStatusCode(v int32) *CreateGatewayFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayFlowRuleResponse) SetBody(v *CreateGatewayFlowRuleResponseBody) *CreateGatewayFlowRuleResponse {
	s.Body = v
	return s
}

func (s *CreateGatewayFlowRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
