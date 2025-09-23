// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLRateLimitingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSQLRateLimitingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSQLRateLimitingRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSQLRateLimitingRulesResponseBody) *DescribeSQLRateLimitingRulesResponse
	GetBody() *DescribeSQLRateLimitingRulesResponseBody
}

type DescribeSQLRateLimitingRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSQLRateLimitingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSQLRateLimitingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLRateLimitingRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLRateLimitingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSQLRateLimitingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSQLRateLimitingRulesResponse) GetBody() *DescribeSQLRateLimitingRulesResponseBody {
	return s.Body
}

func (s *DescribeSQLRateLimitingRulesResponse) SetHeaders(v map[string]*string) *DescribeSQLRateLimitingRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLRateLimitingRulesResponse) SetStatusCode(v int32) *DescribeSQLRateLimitingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLRateLimitingRulesResponse) SetBody(v *DescribeSQLRateLimitingRulesResponseBody) *DescribeSQLRateLimitingRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeSQLRateLimitingRulesResponse) Validate() error {
	return dara.Validate(s)
}
