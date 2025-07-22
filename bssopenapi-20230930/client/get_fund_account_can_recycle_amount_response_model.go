// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountCanRecycleAmountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFundAccountCanRecycleAmountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFundAccountCanRecycleAmountResponse
	GetStatusCode() *int32
	SetBody(v *GetFundAccountCanRecycleAmountResponseBody) *GetFundAccountCanRecycleAmountResponse
	GetBody() *GetFundAccountCanRecycleAmountResponseBody
}

type GetFundAccountCanRecycleAmountResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFundAccountCanRecycleAmountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFundAccountCanRecycleAmountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanRecycleAmountResponse) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanRecycleAmountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFundAccountCanRecycleAmountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFundAccountCanRecycleAmountResponse) GetBody() *GetFundAccountCanRecycleAmountResponseBody {
	return s.Body
}

func (s *GetFundAccountCanRecycleAmountResponse) SetHeaders(v map[string]*string) *GetFundAccountCanRecycleAmountResponse {
	s.Headers = v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponse) SetStatusCode(v int32) *GetFundAccountCanRecycleAmountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponse) SetBody(v *GetFundAccountCanRecycleAmountResponseBody) *GetFundAccountCanRecycleAmountResponse {
	s.Body = v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponse) Validate() error {
	return dara.Validate(s)
}
