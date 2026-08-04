// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateMemberBalanceTransactionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterCreateMemberBalanceTransactionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterCreateMemberBalanceTransactionResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterCreateMemberBalanceTransactionResponseBody) *ModelRouterCreateMemberBalanceTransactionResponse
	GetBody() *ModelRouterCreateMemberBalanceTransactionResponseBody
}

type ModelRouterCreateMemberBalanceTransactionResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterCreateMemberBalanceTransactionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterCreateMemberBalanceTransactionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateMemberBalanceTransactionResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateMemberBalanceTransactionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterCreateMemberBalanceTransactionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterCreateMemberBalanceTransactionResponse) GetBody() *ModelRouterCreateMemberBalanceTransactionResponseBody {
	return s.Body
}

func (s *ModelRouterCreateMemberBalanceTransactionResponse) SetHeaders(v map[string]*string) *ModelRouterCreateMemberBalanceTransactionResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionResponse) SetStatusCode(v int32) *ModelRouterCreateMemberBalanceTransactionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionResponse) SetBody(v *ModelRouterCreateMemberBalanceTransactionResponseBody) *ModelRouterCreateMemberBalanceTransactionResponse {
	s.Body = v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
