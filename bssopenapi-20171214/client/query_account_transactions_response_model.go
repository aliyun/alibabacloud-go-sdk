// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountTransactionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccountTransactionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccountTransactionsResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccountTransactionsResponseBody) *QueryAccountTransactionsResponse
	GetBody() *QueryAccountTransactionsResponseBody
}

type QueryAccountTransactionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountTransactionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountTransactionsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTransactionsResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccountTransactionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccountTransactionsResponse) GetBody() *QueryAccountTransactionsResponseBody {
	return s.Body
}

func (s *QueryAccountTransactionsResponse) SetHeaders(v map[string]*string) *QueryAccountTransactionsResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountTransactionsResponse) SetStatusCode(v int32) *QueryAccountTransactionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountTransactionsResponse) SetBody(v *QueryAccountTransactionsResponseBody) *QueryAccountTransactionsResponse {
	s.Body = v
	return s
}

func (s *QueryAccountTransactionsResponse) Validate() error {
	return dara.Validate(s)
}
