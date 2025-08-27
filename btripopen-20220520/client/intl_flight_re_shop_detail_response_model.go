// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightReShopDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightReShopDetailResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightReShopDetailResponseBody) *IntlFlightReShopDetailResponse
	GetBody() *IntlFlightReShopDetailResponseBody
}

type IntlFlightReShopDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightReShopDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightReShopDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightReShopDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightReShopDetailResponse) GetBody() *IntlFlightReShopDetailResponseBody {
	return s.Body
}

func (s *IntlFlightReShopDetailResponse) SetHeaders(v map[string]*string) *IntlFlightReShopDetailResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightReShopDetailResponse) SetStatusCode(v int32) *IntlFlightReShopDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightReShopDetailResponse) SetBody(v *IntlFlightReShopDetailResponseBody) *IntlFlightReShopDetailResponse {
	s.Body = v
	return s
}

func (s *IntlFlightReShopDetailResponse) Validate() error {
	return dara.Validate(s)
}
