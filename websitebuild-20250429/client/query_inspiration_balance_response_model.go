// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInspirationBalanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInspirationBalanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInspirationBalanceResponse
	GetStatusCode() *int32
	SetBody(v *QueryInspirationBalanceResponseBody) *QueryInspirationBalanceResponse
	GetBody() *QueryInspirationBalanceResponseBody
}

type QueryInspirationBalanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInspirationBalanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInspirationBalanceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInspirationBalanceResponse) GoString() string {
	return s.String()
}

func (s *QueryInspirationBalanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInspirationBalanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInspirationBalanceResponse) GetBody() *QueryInspirationBalanceResponseBody {
	return s.Body
}

func (s *QueryInspirationBalanceResponse) SetHeaders(v map[string]*string) *QueryInspirationBalanceResponse {
	s.Headers = v
	return s
}

func (s *QueryInspirationBalanceResponse) SetStatusCode(v int32) *QueryInspirationBalanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInspirationBalanceResponse) SetBody(v *QueryInspirationBalanceResponseBody) *QueryInspirationBalanceResponse {
	s.Body = v
	return s
}

func (s *QueryInspirationBalanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
