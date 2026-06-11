// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayQuotaRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGatewayQuotaRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGatewayQuotaRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGatewayQuotaRuleResponseBody) *DeleteGatewayQuotaRuleResponse
	GetBody() *DeleteGatewayQuotaRuleResponseBody
}

type DeleteGatewayQuotaRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGatewayQuotaRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGatewayQuotaRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayQuotaRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteGatewayQuotaRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGatewayQuotaRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGatewayQuotaRuleResponse) GetBody() *DeleteGatewayQuotaRuleResponseBody {
	return s.Body
}

func (s *DeleteGatewayQuotaRuleResponse) SetHeaders(v map[string]*string) *DeleteGatewayQuotaRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteGatewayQuotaRuleResponse) SetStatusCode(v int32) *DeleteGatewayQuotaRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGatewayQuotaRuleResponse) SetBody(v *DeleteGatewayQuotaRuleResponseBody) *DeleteGatewayQuotaRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteGatewayQuotaRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
