// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountAvailableAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFundAccountAvailableAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFundAccountAvailableAmountResponse
	GetStatusCode() *int32
	SetBody(v *GetFundAccountAvailableAmountResponseBody) *GetFundAccountAvailableAmountResponse
	GetBody() *GetFundAccountAvailableAmountResponseBody
}

type GetFundAccountAvailableAmountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountAvailableAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountAvailableAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountAvailableAmountResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountAvailableAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFundAccountAvailableAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFundAccountAvailableAmountResponse) GetBody() *GetFundAccountAvailableAmountResponseBody {
	return s.Body
}

func (s *GetFundAccountAvailableAmountResponse) SetHeaders(v map[string]*string) *GetFundAccountAvailableAmountResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountAvailableAmountResponse) SetStatusCode(v int32) *GetFundAccountAvailableAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountAvailableAmountResponse) SetBody(v *GetFundAccountAvailableAmountResponseBody) *GetFundAccountAvailableAmountResponse {
	s.Body = v
	return s
}

func (s *GetFundAccountAvailableAmountResponse) Validate() error {
	return dara.Validate(s)
}
