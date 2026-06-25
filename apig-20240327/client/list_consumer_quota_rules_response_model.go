// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerQuotaRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsumerQuotaRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsumerQuotaRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListConsumerQuotaRulesResponseBody) *ListConsumerQuotaRulesResponse
	GetBody() *ListConsumerQuotaRulesResponseBody
}

type ListConsumerQuotaRulesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumerQuotaRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumerQuotaRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerQuotaRulesResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerQuotaRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsumerQuotaRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsumerQuotaRulesResponse) GetBody() *ListConsumerQuotaRulesResponseBody {
	return s.Body
}

func (s *ListConsumerQuotaRulesResponse) SetHeaders(v map[string]*string) *ListConsumerQuotaRulesResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerQuotaRulesResponse) SetStatusCode(v int32) *ListConsumerQuotaRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerQuotaRulesResponse) SetBody(v *ListConsumerQuotaRulesResponseBody) *ListConsumerQuotaRulesResponse {
	s.Body = v
	return s
}

func (s *ListConsumerQuotaRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
