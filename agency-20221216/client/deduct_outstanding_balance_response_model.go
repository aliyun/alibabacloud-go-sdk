// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeductOutstandingBalanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeductOutstandingBalanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeductOutstandingBalanceResponse
	GetStatusCode() *int32
	SetBody(v *DeductOutstandingBalanceResponseBody) *DeductOutstandingBalanceResponse
	GetBody() *DeductOutstandingBalanceResponseBody
}

type DeductOutstandingBalanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeductOutstandingBalanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeductOutstandingBalanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeductOutstandingBalanceResponse) GoString() string {
	return s.String()
}

func (s *DeductOutstandingBalanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeductOutstandingBalanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeductOutstandingBalanceResponse) GetBody() *DeductOutstandingBalanceResponseBody {
	return s.Body
}

func (s *DeductOutstandingBalanceResponse) SetHeaders(v map[string]*string) *DeductOutstandingBalanceResponse {
	s.Headers = v
	return s
}

func (s *DeductOutstandingBalanceResponse) SetStatusCode(v int32) *DeductOutstandingBalanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeductOutstandingBalanceResponse) SetBody(v *DeductOutstandingBalanceResponseBody) *DeductOutstandingBalanceResponse {
	s.Body = v
	return s
}

func (s *DeductOutstandingBalanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
