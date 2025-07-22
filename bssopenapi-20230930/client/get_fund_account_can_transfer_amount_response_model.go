// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountCanTransferAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFundAccountCanTransferAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFundAccountCanTransferAmountResponse
	GetStatusCode() *int32
	SetBody(v *GetFundAccountCanTransferAmountResponseBody) *GetFundAccountCanTransferAmountResponse
	GetBody() *GetFundAccountCanTransferAmountResponseBody
}

type GetFundAccountCanTransferAmountResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountCanTransferAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountCanTransferAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanTransferAmountResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanTransferAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFundAccountCanTransferAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFundAccountCanTransferAmountResponse) GetBody() *GetFundAccountCanTransferAmountResponseBody {
	return s.Body
}

func (s *GetFundAccountCanTransferAmountResponse) SetHeaders(v map[string]*string) *GetFundAccountCanTransferAmountResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountCanTransferAmountResponse) SetStatusCode(v int32) *GetFundAccountCanTransferAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountCanTransferAmountResponse) SetBody(v *GetFundAccountCanTransferAmountResponseBody) *GetFundAccountCanTransferAmountResponse {
	s.Body = v
	return s
}

func (s *GetFundAccountCanTransferAmountResponse) Validate() error {
	return dara.Validate(s)
}
