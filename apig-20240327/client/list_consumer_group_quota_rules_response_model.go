// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupQuotaRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsumerGroupQuotaRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsumerGroupQuotaRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListConsumerGroupQuotaRulesResponseBody) *ListConsumerGroupQuotaRulesResponse
	GetBody() *ListConsumerGroupQuotaRulesResponseBody
}

type ListConsumerGroupQuotaRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumerGroupQuotaRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumerGroupQuotaRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupQuotaRulesResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupQuotaRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsumerGroupQuotaRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsumerGroupQuotaRulesResponse) GetBody() *ListConsumerGroupQuotaRulesResponseBody {
	return s.Body
}

func (s *ListConsumerGroupQuotaRulesResponse) SetHeaders(v map[string]*string) *ListConsumerGroupQuotaRulesResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponse) SetStatusCode(v int32) *ListConsumerGroupQuotaRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponse) SetBody(v *ListConsumerGroupQuotaRulesResponseBody) *ListConsumerGroupQuotaRulesResponse {
	s.Body = v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
