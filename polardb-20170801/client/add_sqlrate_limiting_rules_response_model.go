// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSQLRateLimitingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSQLRateLimitingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSQLRateLimitingRulesResponse
	GetStatusCode() *int32
	SetBody(v *AddSQLRateLimitingRulesResponseBody) *AddSQLRateLimitingRulesResponse
	GetBody() *AddSQLRateLimitingRulesResponseBody
}

type AddSQLRateLimitingRulesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSQLRateLimitingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSQLRateLimitingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSQLRateLimitingRulesResponse) GoString() string {
	return s.String()
}

func (s *AddSQLRateLimitingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSQLRateLimitingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSQLRateLimitingRulesResponse) GetBody() *AddSQLRateLimitingRulesResponseBody {
	return s.Body
}

func (s *AddSQLRateLimitingRulesResponse) SetHeaders(v map[string]*string) *AddSQLRateLimitingRulesResponse {
	s.Headers = v
	return s
}

func (s *AddSQLRateLimitingRulesResponse) SetStatusCode(v int32) *AddSQLRateLimitingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSQLRateLimitingRulesResponse) SetBody(v *AddSQLRateLimitingRulesResponseBody) *AddSQLRateLimitingRulesResponse {
	s.Body = v
	return s
}

func (s *AddSQLRateLimitingRulesResponse) Validate() error {
	return dara.Validate(s)
}
