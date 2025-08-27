// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundApplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightRefundApplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightRefundApplyResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightRefundApplyResponseBody) *IntlFlightRefundApplyResponse
	GetBody() *IntlFlightRefundApplyResponseBody
}

type IntlFlightRefundApplyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightRefundApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightRefundApplyResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundApplyResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundApplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightRefundApplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightRefundApplyResponse) GetBody() *IntlFlightRefundApplyResponseBody {
	return s.Body
}

func (s *IntlFlightRefundApplyResponse) SetHeaders(v map[string]*string) *IntlFlightRefundApplyResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightRefundApplyResponse) SetStatusCode(v int32) *IntlFlightRefundApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightRefundApplyResponse) SetBody(v *IntlFlightRefundApplyResponseBody) *IntlFlightRefundApplyResponse {
	s.Body = v
	return s
}

func (s *IntlFlightRefundApplyResponse) Validate() error {
	return dara.Validate(s)
}
