// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayQuotaRuleStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayQuotaRuleStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayQuotaRuleStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayQuotaRuleStatusResponseBody) *UpdateGatewayQuotaRuleStatusResponse
	GetBody() *UpdateGatewayQuotaRuleStatusResponseBody
}

type UpdateGatewayQuotaRuleStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayQuotaRuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayQuotaRuleStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayQuotaRuleStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayQuotaRuleStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayQuotaRuleStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayQuotaRuleStatusResponse) GetBody() *UpdateGatewayQuotaRuleStatusResponseBody {
	return s.Body
}

func (s *UpdateGatewayQuotaRuleStatusResponse) SetHeaders(v map[string]*string) *UpdateGatewayQuotaRuleStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayQuotaRuleStatusResponse) SetStatusCode(v int32) *UpdateGatewayQuotaRuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayQuotaRuleStatusResponse) SetBody(v *UpdateGatewayQuotaRuleStatusResponseBody) *UpdateGatewayQuotaRuleStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayQuotaRuleStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
