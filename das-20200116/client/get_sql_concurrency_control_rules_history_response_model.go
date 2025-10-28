// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConcurrencyControlRulesHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSqlConcurrencyControlRulesHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSqlConcurrencyControlRulesHistoryResponse
	GetStatusCode() *int32
	SetBody(v *GetSqlConcurrencyControlRulesHistoryResponseBody) *GetSqlConcurrencyControlRulesHistoryResponse
	GetBody() *GetSqlConcurrencyControlRulesHistoryResponseBody
}

type GetSqlConcurrencyControlRulesHistoryResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSqlConcurrencyControlRulesHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSqlConcurrencyControlRulesHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlRulesHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlRulesHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSqlConcurrencyControlRulesHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSqlConcurrencyControlRulesHistoryResponse) GetBody() *GetSqlConcurrencyControlRulesHistoryResponseBody {
	return s.Body
}

func (s *GetSqlConcurrencyControlRulesHistoryResponse) SetHeaders(v map[string]*string) *GetSqlConcurrencyControlRulesHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponse) SetStatusCode(v int32) *GetSqlConcurrencyControlRulesHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponse) SetBody(v *GetSqlConcurrencyControlRulesHistoryResponseBody) *GetSqlConcurrencyControlRulesHistoryResponse {
	s.Body = v
	return s
}

func (s *GetSqlConcurrencyControlRulesHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
