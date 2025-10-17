// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSQLRateLimitingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSQLRateLimitingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSQLRateLimitingRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSQLRateLimitingRulesResponseBody) *DeleteSQLRateLimitingRulesResponse
	GetBody() *DeleteSQLRateLimitingRulesResponseBody
}

type DeleteSQLRateLimitingRulesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSQLRateLimitingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSQLRateLimitingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSQLRateLimitingRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteSQLRateLimitingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSQLRateLimitingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSQLRateLimitingRulesResponse) GetBody() *DeleteSQLRateLimitingRulesResponseBody {
	return s.Body
}

func (s *DeleteSQLRateLimitingRulesResponse) SetHeaders(v map[string]*string) *DeleteSQLRateLimitingRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteSQLRateLimitingRulesResponse) SetStatusCode(v int32) *DeleteSQLRateLimitingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSQLRateLimitingRulesResponse) SetBody(v *DeleteSQLRateLimitingRulesResponseBody) *DeleteSQLRateLimitingRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteSQLRateLimitingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
