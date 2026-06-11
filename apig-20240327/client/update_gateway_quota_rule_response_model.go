// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayQuotaRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayQuotaRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayQuotaRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayQuotaRuleResponseBody) *UpdateGatewayQuotaRuleResponse
	GetBody() *UpdateGatewayQuotaRuleResponseBody
}

type UpdateGatewayQuotaRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayQuotaRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayQuotaRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayQuotaRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayQuotaRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayQuotaRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayQuotaRuleResponse) GetBody() *UpdateGatewayQuotaRuleResponseBody {
	return s.Body
}

func (s *UpdateGatewayQuotaRuleResponse) SetHeaders(v map[string]*string) *UpdateGatewayQuotaRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayQuotaRuleResponse) SetStatusCode(v int32) *UpdateGatewayQuotaRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayQuotaRuleResponse) SetBody(v *UpdateGatewayQuotaRuleResponseBody) *UpdateGatewayQuotaRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayQuotaRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
