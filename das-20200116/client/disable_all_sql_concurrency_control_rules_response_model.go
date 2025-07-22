// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAllSqlConcurrencyControlRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAllSqlConcurrencyControlRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAllSqlConcurrencyControlRulesResponse
	GetStatusCode() *int32
	SetBody(v *DisableAllSqlConcurrencyControlRulesResponseBody) *DisableAllSqlConcurrencyControlRulesResponse
	GetBody() *DisableAllSqlConcurrencyControlRulesResponseBody
}

type DisableAllSqlConcurrencyControlRulesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAllSqlConcurrencyControlRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAllSqlConcurrencyControlRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAllSqlConcurrencyControlRulesResponse) GoString() string {
	return s.String()
}

func (s *DisableAllSqlConcurrencyControlRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAllSqlConcurrencyControlRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAllSqlConcurrencyControlRulesResponse) GetBody() *DisableAllSqlConcurrencyControlRulesResponseBody {
	return s.Body
}

func (s *DisableAllSqlConcurrencyControlRulesResponse) SetHeaders(v map[string]*string) *DisableAllSqlConcurrencyControlRulesResponse {
	s.Headers = v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponse) SetStatusCode(v int32) *DisableAllSqlConcurrencyControlRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponse) SetBody(v *DisableAllSqlConcurrencyControlRulesResponseBody) *DisableAllSqlConcurrencyControlRulesResponse {
	s.Body = v
	return s
}

func (s *DisableAllSqlConcurrencyControlRulesResponse) Validate() error {
	return dara.Validate(s)
}
