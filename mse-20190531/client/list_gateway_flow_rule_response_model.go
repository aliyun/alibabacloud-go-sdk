// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayFlowRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayFlowRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayFlowRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayFlowRuleResponseBody) *ListGatewayFlowRuleResponse
	GetBody() *ListGatewayFlowRuleResponseBody
}

type ListGatewayFlowRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayFlowRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayFlowRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayFlowRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayFlowRuleResponse) GetBody() *ListGatewayFlowRuleResponseBody {
	return s.Body
}

func (s *ListGatewayFlowRuleResponse) SetHeaders(v map[string]*string) *ListGatewayFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayFlowRuleResponse) SetStatusCode(v int32) *ListGatewayFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayFlowRuleResponse) SetBody(v *ListGatewayFlowRuleResponseBody) *ListGatewayFlowRuleResponse {
	s.Body = v
	return s
}

func (s *ListGatewayFlowRuleResponse) Validate() error {
	return dara.Validate(s)
}
