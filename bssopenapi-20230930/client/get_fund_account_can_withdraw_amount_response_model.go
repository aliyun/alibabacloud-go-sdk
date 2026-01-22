// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountCanWithdrawAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFundAccountCanWithdrawAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFundAccountCanWithdrawAmountResponse
	GetStatusCode() *int32
	SetBody(v *GetFundAccountCanWithdrawAmountResponseBody) *GetFundAccountCanWithdrawAmountResponse
	GetBody() *GetFundAccountCanWithdrawAmountResponseBody
}

type GetFundAccountCanWithdrawAmountResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountCanWithdrawAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountCanWithdrawAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanWithdrawAmountResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanWithdrawAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFundAccountCanWithdrawAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFundAccountCanWithdrawAmountResponse) GetBody() *GetFundAccountCanWithdrawAmountResponseBody {
	return s.Body
}

func (s *GetFundAccountCanWithdrawAmountResponse) SetHeaders(v map[string]*string) *GetFundAccountCanWithdrawAmountResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponse) SetStatusCode(v int32) *GetFundAccountCanWithdrawAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponse) SetBody(v *GetFundAccountCanWithdrawAmountResponseBody) *GetFundAccountCanWithdrawAmountResponse {
	s.Body = v
	return s
}

func (s *GetFundAccountCanWithdrawAmountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
