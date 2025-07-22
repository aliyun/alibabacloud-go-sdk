// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConcurrencyControlKeywordsFromSqlTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSqlConcurrencyControlKeywordsFromSqlTextResponse
	GetStatusCode() *int32
	SetBody(v *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) *GetSqlConcurrencyControlKeywordsFromSqlTextResponse
	GetBody() *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody
}

type GetSqlConcurrencyControlKeywordsFromSqlTextResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSqlConcurrencyControlKeywordsFromSqlTextResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConcurrencyControlKeywordsFromSqlTextResponse) GoString() string {
	return s.String()
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponse) GetBody() *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody {
	return s.Body
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponse) SetHeaders(v map[string]*string) *GetSqlConcurrencyControlKeywordsFromSqlTextResponse {
	s.Headers = v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponse) SetStatusCode(v int32) *GetSqlConcurrencyControlKeywordsFromSqlTextResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponse) SetBody(v *GetSqlConcurrencyControlKeywordsFromSqlTextResponseBody) *GetSqlConcurrencyControlKeywordsFromSqlTextResponse {
	s.Body = v
	return s
}

func (s *GetSqlConcurrencyControlKeywordsFromSqlTextResponse) Validate() error {
	return dara.Validate(s)
}
