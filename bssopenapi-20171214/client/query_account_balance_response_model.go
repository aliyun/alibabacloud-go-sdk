// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountBalanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccountBalanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccountBalanceResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccountBalanceResponseBody) *QueryAccountBalanceResponse
	GetBody() *QueryAccountBalanceResponseBody
}

type QueryAccountBalanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountBalanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountBalanceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountBalanceResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountBalanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccountBalanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccountBalanceResponse) GetBody() *QueryAccountBalanceResponseBody {
	return s.Body
}

func (s *QueryAccountBalanceResponse) SetHeaders(v map[string]*string) *QueryAccountBalanceResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountBalanceResponse) SetStatusCode(v int32) *QueryAccountBalanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountBalanceResponse) SetBody(v *QueryAccountBalanceResponseBody) *QueryAccountBalanceResponse {
	s.Body = v
	return s
}

func (s *QueryAccountBalanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
