// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountTransactionDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccountTransactionDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccountTransactionDetailsResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccountTransactionDetailsResponseBody) *QueryAccountTransactionDetailsResponse
	GetBody() *QueryAccountTransactionDetailsResponseBody
}

type QueryAccountTransactionDetailsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountTransactionDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountTransactionDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTransactionDetailsResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccountTransactionDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccountTransactionDetailsResponse) GetBody() *QueryAccountTransactionDetailsResponseBody {
	return s.Body
}

func (s *QueryAccountTransactionDetailsResponse) SetHeaders(v map[string]*string) *QueryAccountTransactionDetailsResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountTransactionDetailsResponse) SetStatusCode(v int32) *QueryAccountTransactionDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponse) SetBody(v *QueryAccountTransactionDetailsResponseBody) *QueryAccountTransactionDetailsResponse {
	s.Body = v
	return s
}

func (s *QueryAccountTransactionDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
