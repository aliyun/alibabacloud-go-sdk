// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayIsolationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGatewayIsolationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGatewayIsolationRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateGatewayIsolationRuleResponseBody) *CreateGatewayIsolationRuleResponse
	GetBody() *CreateGatewayIsolationRuleResponseBody
}

type CreateGatewayIsolationRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGatewayIsolationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGatewayIsolationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayIsolationRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateGatewayIsolationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGatewayIsolationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGatewayIsolationRuleResponse) GetBody() *CreateGatewayIsolationRuleResponseBody {
	return s.Body
}

func (s *CreateGatewayIsolationRuleResponse) SetHeaders(v map[string]*string) *CreateGatewayIsolationRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateGatewayIsolationRuleResponse) SetStatusCode(v int32) *CreateGatewayIsolationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGatewayIsolationRuleResponse) SetBody(v *CreateGatewayIsolationRuleResponseBody) *CreateGatewayIsolationRuleResponse {
	s.Body = v
	return s
}

func (s *CreateGatewayIsolationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
