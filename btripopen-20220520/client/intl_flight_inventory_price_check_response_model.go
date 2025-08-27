// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightInventoryPriceCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightInventoryPriceCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightInventoryPriceCheckResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightInventoryPriceCheckResponseBody) *IntlFlightInventoryPriceCheckResponse
	GetBody() *IntlFlightInventoryPriceCheckResponseBody
}

type IntlFlightInventoryPriceCheckResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightInventoryPriceCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightInventoryPriceCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightInventoryPriceCheckResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightInventoryPriceCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightInventoryPriceCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightInventoryPriceCheckResponse) GetBody() *IntlFlightInventoryPriceCheckResponseBody {
	return s.Body
}

func (s *IntlFlightInventoryPriceCheckResponse) SetHeaders(v map[string]*string) *IntlFlightInventoryPriceCheckResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponse) SetStatusCode(v int32) *IntlFlightInventoryPriceCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponse) SetBody(v *IntlFlightInventoryPriceCheckResponseBody) *IntlFlightInventoryPriceCheckResponse {
	s.Body = v
	return s
}

func (s *IntlFlightInventoryPriceCheckResponse) Validate() error {
	return dara.Validate(s)
}
