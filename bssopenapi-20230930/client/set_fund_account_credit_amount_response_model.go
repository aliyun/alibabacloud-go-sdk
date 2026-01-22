// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFundAccountCreditAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetFundAccountCreditAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetFundAccountCreditAmountResponse
	GetStatusCode() *int32
	SetBody(v *SetFundAccountCreditAmountResponseBody) *SetFundAccountCreditAmountResponse
	GetBody() *SetFundAccountCreditAmountResponseBody
}

type SetFundAccountCreditAmountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetFundAccountCreditAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetFundAccountCreditAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s SetFundAccountCreditAmountResponse) GoString() string {
	return s.String()
}

func (s *SetFundAccountCreditAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetFundAccountCreditAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetFundAccountCreditAmountResponse) GetBody() *SetFundAccountCreditAmountResponseBody {
	return s.Body
}

func (s *SetFundAccountCreditAmountResponse) SetHeaders(v map[string]*string) *SetFundAccountCreditAmountResponse {
	s.Headers = v
	return s
}

func (s *SetFundAccountCreditAmountResponse) SetStatusCode(v int32) *SetFundAccountCreditAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *SetFundAccountCreditAmountResponse) SetBody(v *SetFundAccountCreditAmountResponseBody) *SetFundAccountCreditAmountResponse {
	s.Body = v
	return s
}

func (s *SetFundAccountCreditAmountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
