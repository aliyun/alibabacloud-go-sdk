// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayQuotaRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGatewayQuotaRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGatewayQuotaRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddGatewayQuotaRuleResponseBody) *AddGatewayQuotaRuleResponse
	GetBody() *AddGatewayQuotaRuleResponseBody
}

type AddGatewayQuotaRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGatewayQuotaRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGatewayQuotaRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayQuotaRuleResponse) GoString() string {
	return s.String()
}

func (s *AddGatewayQuotaRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGatewayQuotaRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGatewayQuotaRuleResponse) GetBody() *AddGatewayQuotaRuleResponseBody {
	return s.Body
}

func (s *AddGatewayQuotaRuleResponse) SetHeaders(v map[string]*string) *AddGatewayQuotaRuleResponse {
	s.Headers = v
	return s
}

func (s *AddGatewayQuotaRuleResponse) SetStatusCode(v int32) *AddGatewayQuotaRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGatewayQuotaRuleResponse) SetBody(v *AddGatewayQuotaRuleResponseBody) *AddGatewayQuotaRuleResponse {
	s.Body = v
	return s
}

func (s *AddGatewayQuotaRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
