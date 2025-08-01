// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayFlowRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayFlowRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayFlowRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayFlowRuleResponseBody) *DeleteGatewayFlowRuleResponse
	GetBody() *DeleteGatewayFlowRuleResponseBody
}

type DeleteGatewayFlowRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayFlowRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayFlowRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayFlowRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayFlowRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayFlowRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayFlowRuleResponse) GetBody() *DeleteGatewayFlowRuleResponseBody {
	return s.Body
}

func (s *DeleteGatewayFlowRuleResponse) SetHeaders(v map[string]*string) *DeleteGatewayFlowRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayFlowRuleResponse) SetStatusCode(v int32) *DeleteGatewayFlowRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayFlowRuleResponse) SetBody(v *DeleteGatewayFlowRuleResponseBody) *DeleteGatewayFlowRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayFlowRuleResponse) Validate() error {
	return dara.Validate(s)
}
