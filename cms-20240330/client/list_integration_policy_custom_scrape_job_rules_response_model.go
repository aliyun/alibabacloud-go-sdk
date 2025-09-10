// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyCustomScrapeJobRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntegrationPolicyCustomScrapeJobRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntegrationPolicyCustomScrapeJobRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListIntegrationPolicyCustomScrapeJobRulesResponseBody) *ListIntegrationPolicyCustomScrapeJobRulesResponse
	GetBody() *ListIntegrationPolicyCustomScrapeJobRulesResponseBody
}

type ListIntegrationPolicyCustomScrapeJobRulesResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntegrationPolicyCustomScrapeJobRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntegrationPolicyCustomScrapeJobRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyCustomScrapeJobRulesResponse) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponse) GetBody() *ListIntegrationPolicyCustomScrapeJobRulesResponseBody {
	return s.Body
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponse) SetHeaders(v map[string]*string) *ListIntegrationPolicyCustomScrapeJobRulesResponse {
	s.Headers = v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponse) SetStatusCode(v int32) *ListIntegrationPolicyCustomScrapeJobRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponse) SetBody(v *ListIntegrationPolicyCustomScrapeJobRulesResponseBody) *ListIntegrationPolicyCustomScrapeJobRulesResponse {
	s.Body = v
	return s
}

func (s *ListIntegrationPolicyCustomScrapeJobRulesResponse) Validate() error {
	return dara.Validate(s)
}
