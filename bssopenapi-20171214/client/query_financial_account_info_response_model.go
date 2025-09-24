// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFinancialAccountInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFinancialAccountInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFinancialAccountInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryFinancialAccountInfoResponseBody) *QueryFinancialAccountInfoResponse
	GetBody() *QueryFinancialAccountInfoResponseBody
}

type QueryFinancialAccountInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFinancialAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFinancialAccountInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFinancialAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryFinancialAccountInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFinancialAccountInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFinancialAccountInfoResponse) GetBody() *QueryFinancialAccountInfoResponseBody {
	return s.Body
}

func (s *QueryFinancialAccountInfoResponse) SetHeaders(v map[string]*string) *QueryFinancialAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryFinancialAccountInfoResponse) SetStatusCode(v int32) *QueryFinancialAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFinancialAccountInfoResponse) SetBody(v *QueryFinancialAccountInfoResponseBody) *QueryFinancialAccountInfoResponse {
	s.Body = v
	return s
}

func (s *QueryFinancialAccountInfoResponse) Validate() error {
	return dara.Validate(s)
}
