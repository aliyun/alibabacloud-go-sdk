// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountTransactionDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFundAccountTransactionDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFundAccountTransactionDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetFundAccountTransactionDetailsResponseBody) *GetFundAccountTransactionDetailsResponse
	GetBody() *GetFundAccountTransactionDetailsResponseBody
}

type GetFundAccountTransactionDetailsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountTransactionDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountTransactionDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountTransactionDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountTransactionDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFundAccountTransactionDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFundAccountTransactionDetailsResponse) GetBody() *GetFundAccountTransactionDetailsResponseBody {
	return s.Body
}

func (s *GetFundAccountTransactionDetailsResponse) SetHeaders(v map[string]*string) *GetFundAccountTransactionDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountTransactionDetailsResponse) SetStatusCode(v int32) *GetFundAccountTransactionDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountTransactionDetailsResponse) SetBody(v *GetFundAccountTransactionDetailsResponseBody) *GetFundAccountTransactionDetailsResponse {
	s.Body = v
	return s
}

func (s *GetFundAccountTransactionDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
