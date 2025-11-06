// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayIsolationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayIsolationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayIsolationRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayIsolationRuleResponseBody) *DeleteGatewayIsolationRuleResponse
	GetBody() *DeleteGatewayIsolationRuleResponseBody
}

type DeleteGatewayIsolationRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayIsolationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayIsolationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayIsolationRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIsolationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayIsolationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayIsolationRuleResponse) GetBody() *DeleteGatewayIsolationRuleResponseBody {
	return s.Body
}

func (s *DeleteGatewayIsolationRuleResponse) SetHeaders(v map[string]*string) *DeleteGatewayIsolationRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayIsolationRuleResponse) SetStatusCode(v int32) *DeleteGatewayIsolationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayIsolationRuleResponse) SetBody(v *DeleteGatewayIsolationRuleResponseBody) *DeleteGatewayIsolationRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayIsolationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
