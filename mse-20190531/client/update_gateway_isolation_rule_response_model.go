// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayIsolationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayIsolationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayIsolationRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayIsolationRuleResponseBody) *UpdateGatewayIsolationRuleResponse
	GetBody() *UpdateGatewayIsolationRuleResponseBody
}

type UpdateGatewayIsolationRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayIsolationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayIsolationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayIsolationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayIsolationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayIsolationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayIsolationRuleResponse) GetBody() *UpdateGatewayIsolationRuleResponseBody {
	return s.Body
}

func (s *UpdateGatewayIsolationRuleResponse) SetHeaders(v map[string]*string) *UpdateGatewayIsolationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayIsolationRuleResponse) SetStatusCode(v int32) *UpdateGatewayIsolationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayIsolationRuleResponse) SetBody(v *UpdateGatewayIsolationRuleResponseBody) *UpdateGatewayIsolationRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayIsolationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
