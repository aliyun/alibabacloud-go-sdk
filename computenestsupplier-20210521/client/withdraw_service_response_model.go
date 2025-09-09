// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WithdrawServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WithdrawServiceResponse
	GetStatusCode() *int32
	SetBody(v *WithdrawServiceResponseBody) *WithdrawServiceResponse
	GetBody() *WithdrawServiceResponseBody
}

type WithdrawServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WithdrawServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WithdrawServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s WithdrawServiceResponse) GoString() string {
	return s.String()
}

func (s *WithdrawServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WithdrawServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WithdrawServiceResponse) GetBody() *WithdrawServiceResponseBody {
	return s.Body
}

func (s *WithdrawServiceResponse) SetHeaders(v map[string]*string) *WithdrawServiceResponse {
	s.Headers = v
	return s
}

func (s *WithdrawServiceResponse) SetStatusCode(v int32) *WithdrawServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *WithdrawServiceResponse) SetBody(v *WithdrawServiceResponseBody) *WithdrawServiceResponse {
	s.Body = v
	return s
}

func (s *WithdrawServiceResponse) Validate() error {
	return dara.Validate(s)
}
