// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayIsolationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayIsolationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayIsolationRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayIsolationRuleResponseBody) *ListGatewayIsolationRuleResponse
	GetBody() *ListGatewayIsolationRuleResponseBody
}

type ListGatewayIsolationRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayIsolationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayIsolationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIsolationRuleResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayIsolationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayIsolationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayIsolationRuleResponse) GetBody() *ListGatewayIsolationRuleResponseBody {
	return s.Body
}

func (s *ListGatewayIsolationRuleResponse) SetHeaders(v map[string]*string) *ListGatewayIsolationRuleResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayIsolationRuleResponse) SetStatusCode(v int32) *ListGatewayIsolationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayIsolationRuleResponse) SetBody(v *ListGatewayIsolationRuleResponseBody) *ListGatewayIsolationRuleResponse {
	s.Body = v
	return s
}

func (s *ListGatewayIsolationRuleResponse) Validate() error {
	return dara.Validate(s)
}
