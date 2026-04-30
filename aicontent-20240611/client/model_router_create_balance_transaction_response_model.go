// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateBalanceTransactionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCreateBalanceTransactionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCreateBalanceTransactionResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCreateBalanceTransactionResponseBody) *ModelRouterCreateBalanceTransactionResponse
	GetBody() *ModelRouterCreateBalanceTransactionResponseBody
}

type ModelRouterCreateBalanceTransactionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCreateBalanceTransactionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCreateBalanceTransactionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateBalanceTransactionResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateBalanceTransactionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCreateBalanceTransactionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCreateBalanceTransactionResponse) GetBody() *ModelRouterCreateBalanceTransactionResponseBody {
	return s.Body
}

func (s *ModelRouterCreateBalanceTransactionResponse) SetHeaders(v map[string]*string) *ModelRouterCreateBalanceTransactionResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCreateBalanceTransactionResponse) SetStatusCode(v int32) *ModelRouterCreateBalanceTransactionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCreateBalanceTransactionResponse) SetBody(v *ModelRouterCreateBalanceTransactionResponseBody) *ModelRouterCreateBalanceTransactionResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCreateBalanceTransactionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
