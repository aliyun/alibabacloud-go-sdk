// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayFlowRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayFlowRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayFlowRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayFlowRuleResponseBody) *UpdateGatewayFlowRuleResponse
	GetBody() *UpdateGatewayFlowRuleResponseBody
}

type UpdateGatewayFlowRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayFlowRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayFlowRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayFlowRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayFlowRuleResponse) GetBody() *UpdateGatewayFlowRuleResponseBody {
	return s.Body
}

func (s *UpdateGatewayFlowRuleResponse) SetHeaders(v map[string]*string) *UpdateGatewayFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayFlowRuleResponse) SetStatusCode(v int32) *UpdateGatewayFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayFlowRuleResponse) SetBody(v *UpdateGatewayFlowRuleResponseBody) *UpdateGatewayFlowRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayFlowRuleResponse) Validate() error {
	return dara.Validate(s)
}
