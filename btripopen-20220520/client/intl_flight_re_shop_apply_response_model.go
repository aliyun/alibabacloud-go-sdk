// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopApplyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightReShopApplyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightReShopApplyResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightReShopApplyResponseBody) *IntlFlightReShopApplyResponse
	GetBody() *IntlFlightReShopApplyResponseBody
}

type IntlFlightReShopApplyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightReShopApplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightReShopApplyResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopApplyResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopApplyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightReShopApplyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightReShopApplyResponse) GetBody() *IntlFlightReShopApplyResponseBody {
	return s.Body
}

func (s *IntlFlightReShopApplyResponse) SetHeaders(v map[string]*string) *IntlFlightReShopApplyResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightReShopApplyResponse) SetStatusCode(v int32) *IntlFlightReShopApplyResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightReShopApplyResponse) SetBody(v *IntlFlightReShopApplyResponseBody) *IntlFlightReShopApplyResponse {
	s.Body = v
	return s
}

func (s *IntlFlightReShopApplyResponse) Validate() error {
	return dara.Validate(s)
}
