// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightCreateOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightCreateOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightCreateOrderResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightCreateOrderResponseBody) *IntlFlightCreateOrderResponse
	GetBody() *IntlFlightCreateOrderResponseBody
}

type IntlFlightCreateOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightCreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightCreateOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightCreateOrderResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightCreateOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightCreateOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightCreateOrderResponse) GetBody() *IntlFlightCreateOrderResponseBody {
	return s.Body
}

func (s *IntlFlightCreateOrderResponse) SetHeaders(v map[string]*string) *IntlFlightCreateOrderResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightCreateOrderResponse) SetStatusCode(v int32) *IntlFlightCreateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightCreateOrderResponse) SetBody(v *IntlFlightCreateOrderResponseBody) *IntlFlightCreateOrderResponse {
	s.Body = v
	return s
}

func (s *IntlFlightCreateOrderResponse) Validate() error {
	return dara.Validate(s)
}
