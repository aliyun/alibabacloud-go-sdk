// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAgentCreditQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAgentCreditQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAgentCreditQuotaResponse
	GetStatusCode() *int32
	SetBody(v *SetAgentCreditQuotaResponseBody) *SetAgentCreditQuotaResponse
	GetBody() *SetAgentCreditQuotaResponseBody
}

type SetAgentCreditQuotaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAgentCreditQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAgentCreditQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAgentCreditQuotaResponse) GoString() string {
	return s.String()
}

func (s *SetAgentCreditQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAgentCreditQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAgentCreditQuotaResponse) GetBody() *SetAgentCreditQuotaResponseBody {
	return s.Body
}

func (s *SetAgentCreditQuotaResponse) SetHeaders(v map[string]*string) *SetAgentCreditQuotaResponse {
	s.Headers = v
	return s
}

func (s *SetAgentCreditQuotaResponse) SetStatusCode(v int32) *SetAgentCreditQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAgentCreditQuotaResponse) SetBody(v *SetAgentCreditQuotaResponseBody) *SetAgentCreditQuotaResponse {
	s.Body = v
	return s
}

func (s *SetAgentCreditQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
