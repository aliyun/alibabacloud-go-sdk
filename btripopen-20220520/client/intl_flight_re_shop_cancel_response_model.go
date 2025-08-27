// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopCancelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightReShopCancelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightReShopCancelResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightReShopCancelResponseBody) *IntlFlightReShopCancelResponse
	GetBody() *IntlFlightReShopCancelResponseBody
}

type IntlFlightReShopCancelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightReShopCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightReShopCancelResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopCancelResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopCancelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightReShopCancelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightReShopCancelResponse) GetBody() *IntlFlightReShopCancelResponseBody {
	return s.Body
}

func (s *IntlFlightReShopCancelResponse) SetHeaders(v map[string]*string) *IntlFlightReShopCancelResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightReShopCancelResponse) SetStatusCode(v int32) *IntlFlightReShopCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightReShopCancelResponse) SetBody(v *IntlFlightReShopCancelResponseBody) *IntlFlightReShopCancelResponse {
	s.Body = v
	return s
}

func (s *IntlFlightReShopCancelResponse) Validate() error {
	return dara.Validate(s)
}
