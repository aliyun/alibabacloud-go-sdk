// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayQuotaRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayQuotaRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayQuotaRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayQuotaRuleResponseBody) *GetGatewayQuotaRuleResponse
	GetBody() *GetGatewayQuotaRuleResponseBody
}

type GetGatewayQuotaRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayQuotaRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayQuotaRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayQuotaRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayQuotaRuleResponse) GetBody() *GetGatewayQuotaRuleResponseBody {
	return s.Body
}

func (s *GetGatewayQuotaRuleResponse) SetHeaders(v map[string]*string) *GetGatewayQuotaRuleResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayQuotaRuleResponse) SetStatusCode(v int32) *GetGatewayQuotaRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayQuotaRuleResponse) SetBody(v *GetGatewayQuotaRuleResponseBody) *GetGatewayQuotaRuleResponse {
	s.Body = v
	return s
}

func (s *GetGatewayQuotaRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
