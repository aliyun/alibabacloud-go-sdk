// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightRefundDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightRefundDetailResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightRefundDetailResponseBody) *IntlFlightRefundDetailResponse
	GetBody() *IntlFlightRefundDetailResponseBody
}

type IntlFlightRefundDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightRefundDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightRefundDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundDetailResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightRefundDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightRefundDetailResponse) GetBody() *IntlFlightRefundDetailResponseBody {
	return s.Body
}

func (s *IntlFlightRefundDetailResponse) SetHeaders(v map[string]*string) *IntlFlightRefundDetailResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightRefundDetailResponse) SetStatusCode(v int32) *IntlFlightRefundDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightRefundDetailResponse) SetBody(v *IntlFlightRefundDetailResponseBody) *IntlFlightRefundDetailResponse {
	s.Body = v
	return s
}

func (s *IntlFlightRefundDetailResponse) Validate() error {
	return dara.Validate(s)
}
