// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountCanAllocateCreditAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFundAccountCanAllocateCreditAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFundAccountCanAllocateCreditAmountResponse
	GetStatusCode() *int32
	SetBody(v *GetFundAccountCanAllocateCreditAmountResponseBody) *GetFundAccountCanAllocateCreditAmountResponse
	GetBody() *GetFundAccountCanAllocateCreditAmountResponseBody
}

type GetFundAccountCanAllocateCreditAmountResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountCanAllocateCreditAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountCanAllocateCreditAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanAllocateCreditAmountResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanAllocateCreditAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFundAccountCanAllocateCreditAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFundAccountCanAllocateCreditAmountResponse) GetBody() *GetFundAccountCanAllocateCreditAmountResponseBody {
	return s.Body
}

func (s *GetFundAccountCanAllocateCreditAmountResponse) SetHeaders(v map[string]*string) *GetFundAccountCanAllocateCreditAmountResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponse) SetStatusCode(v int32) *GetFundAccountCanAllocateCreditAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponse) SetBody(v *GetFundAccountCanAllocateCreditAmountResponseBody) *GetFundAccountCanAllocateCreditAmountResponse {
	s.Body = v
	return s
}

func (s *GetFundAccountCanAllocateCreditAmountResponse) Validate() error {
	return dara.Validate(s)
}
