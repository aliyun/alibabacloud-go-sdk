// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetGatewayQuotaRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetGatewayQuotaRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetGatewayQuotaRuleResponse
	GetStatusCode() *int32
	SetBody(v *ResetGatewayQuotaRuleResponseBody) *ResetGatewayQuotaRuleResponse
	GetBody() *ResetGatewayQuotaRuleResponseBody
}

type ResetGatewayQuotaRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetGatewayQuotaRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetGatewayQuotaRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetGatewayQuotaRuleResponse) GoString() string {
	return s.String()
}

func (s *ResetGatewayQuotaRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetGatewayQuotaRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetGatewayQuotaRuleResponse) GetBody() *ResetGatewayQuotaRuleResponseBody {
	return s.Body
}

func (s *ResetGatewayQuotaRuleResponse) SetHeaders(v map[string]*string) *ResetGatewayQuotaRuleResponse {
	s.Headers = v
	return s
}

func (s *ResetGatewayQuotaRuleResponse) SetStatusCode(v int32) *ResetGatewayQuotaRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetGatewayQuotaRuleResponse) SetBody(v *ResetGatewayQuotaRuleResponseBody) *ResetGatewayQuotaRuleResponse {
	s.Body = v
	return s
}

func (s *ResetGatewayQuotaRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
