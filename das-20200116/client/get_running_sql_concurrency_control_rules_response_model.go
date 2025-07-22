// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningSqlConcurrencyControlRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRunningSqlConcurrencyControlRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRunningSqlConcurrencyControlRulesResponse
	GetStatusCode() *int32
	SetBody(v *GetRunningSqlConcurrencyControlRulesResponseBody) *GetRunningSqlConcurrencyControlRulesResponse
	GetBody() *GetRunningSqlConcurrencyControlRulesResponseBody
}

type GetRunningSqlConcurrencyControlRulesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRunningSqlConcurrencyControlRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRunningSqlConcurrencyControlRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRunningSqlConcurrencyControlRulesResponse) GoString() string {
	return s.String()
}

func (s *GetRunningSqlConcurrencyControlRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRunningSqlConcurrencyControlRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRunningSqlConcurrencyControlRulesResponse) GetBody() *GetRunningSqlConcurrencyControlRulesResponseBody {
	return s.Body
}

func (s *GetRunningSqlConcurrencyControlRulesResponse) SetHeaders(v map[string]*string) *GetRunningSqlConcurrencyControlRulesResponse {
	s.Headers = v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponse) SetStatusCode(v int32) *GetRunningSqlConcurrencyControlRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponse) SetBody(v *GetRunningSqlConcurrencyControlRulesResponseBody) *GetRunningSqlConcurrencyControlRulesResponse {
	s.Body = v
	return s
}

func (s *GetRunningSqlConcurrencyControlRulesResponse) Validate() error {
	return dara.Validate(s)
}
