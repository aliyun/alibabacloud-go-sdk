// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayQuotaRuleSubjectUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGatewayQuotaRuleSubjectUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGatewayQuotaRuleSubjectUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetGatewayQuotaRuleSubjectUsageResponseBody) *GetGatewayQuotaRuleSubjectUsageResponse
	GetBody() *GetGatewayQuotaRuleSubjectUsageResponseBody
}

type GetGatewayQuotaRuleSubjectUsageResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGatewayQuotaRuleSubjectUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGatewayQuotaRuleSubjectUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleSubjectUsageResponse) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleSubjectUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGatewayQuotaRuleSubjectUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGatewayQuotaRuleSubjectUsageResponse) GetBody() *GetGatewayQuotaRuleSubjectUsageResponseBody {
	return s.Body
}

func (s *GetGatewayQuotaRuleSubjectUsageResponse) SetHeaders(v map[string]*string) *GetGatewayQuotaRuleSubjectUsageResponse {
	s.Headers = v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponse) SetStatusCode(v int32) *GetGatewayQuotaRuleSubjectUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponse) SetBody(v *GetGatewayQuotaRuleSubjectUsageResponseBody) *GetGatewayQuotaRuleSubjectUsageResponse {
	s.Body = v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
