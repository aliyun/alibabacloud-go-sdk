// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayQuotaRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayQuotaRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayQuotaRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayQuotaRulesResponseBody) *ListGatewayQuotaRulesResponse
	GetBody() *ListGatewayQuotaRulesResponseBody
}

type ListGatewayQuotaRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayQuotaRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayQuotaRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayQuotaRulesResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayQuotaRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayQuotaRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayQuotaRulesResponse) GetBody() *ListGatewayQuotaRulesResponseBody {
	return s.Body
}

func (s *ListGatewayQuotaRulesResponse) SetHeaders(v map[string]*string) *ListGatewayQuotaRulesResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayQuotaRulesResponse) SetStatusCode(v int32) *ListGatewayQuotaRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayQuotaRulesResponse) SetBody(v *ListGatewayQuotaRulesResponseBody) *ListGatewayQuotaRulesResponse {
	s.Body = v
	return s
}

func (s *ListGatewayQuotaRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
