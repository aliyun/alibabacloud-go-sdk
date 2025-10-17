// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQLRateLimitingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySQLRateLimitingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySQLRateLimitingRulesResponse
	GetStatusCode() *int32
	SetBody(v *ModifySQLRateLimitingRulesResponseBody) *ModifySQLRateLimitingRulesResponse
	GetBody() *ModifySQLRateLimitingRulesResponseBody
}

type ModifySQLRateLimitingRulesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySQLRateLimitingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySQLRateLimitingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySQLRateLimitingRulesResponse) GoString() string {
	return s.String()
}

func (s *ModifySQLRateLimitingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySQLRateLimitingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySQLRateLimitingRulesResponse) GetBody() *ModifySQLRateLimitingRulesResponseBody {
	return s.Body
}

func (s *ModifySQLRateLimitingRulesResponse) SetHeaders(v map[string]*string) *ModifySQLRateLimitingRulesResponse {
	s.Headers = v
	return s
}

func (s *ModifySQLRateLimitingRulesResponse) SetStatusCode(v int32) *ModifySQLRateLimitingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySQLRateLimitingRulesResponse) SetBody(v *ModifySQLRateLimitingRulesResponseBody) *ModifySQLRateLimitingRulesResponse {
	s.Body = v
	return s
}

func (s *ModifySQLRateLimitingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
