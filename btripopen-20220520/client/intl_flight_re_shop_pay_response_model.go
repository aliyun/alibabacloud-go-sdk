// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopPayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightReShopPayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightReShopPayResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightReShopPayResponseBody) *IntlFlightReShopPayResponse
	GetBody() *IntlFlightReShopPayResponseBody
}

type IntlFlightReShopPayResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightReShopPayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightReShopPayResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopPayResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopPayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightReShopPayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightReShopPayResponse) GetBody() *IntlFlightReShopPayResponseBody {
	return s.Body
}

func (s *IntlFlightReShopPayResponse) SetHeaders(v map[string]*string) *IntlFlightReShopPayResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightReShopPayResponse) SetStatusCode(v int32) *IntlFlightReShopPayResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightReShopPayResponse) SetBody(v *IntlFlightReShopPayResponseBody) *IntlFlightReShopPayResponse {
	s.Body = v
	return s
}

func (s *IntlFlightReShopPayResponse) Validate() error {
	return dara.Validate(s)
}
