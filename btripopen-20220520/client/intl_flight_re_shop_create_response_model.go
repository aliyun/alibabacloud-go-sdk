// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightReShopCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightReShopCreateResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightReShopCreateResponseBody) *IntlFlightReShopCreateResponse
	GetBody() *IntlFlightReShopCreateResponseBody
}

type IntlFlightReShopCreateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightReShopCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightReShopCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopCreateResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightReShopCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightReShopCreateResponse) GetBody() *IntlFlightReShopCreateResponseBody {
	return s.Body
}

func (s *IntlFlightReShopCreateResponse) SetHeaders(v map[string]*string) *IntlFlightReShopCreateResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightReShopCreateResponse) SetStatusCode(v int32) *IntlFlightReShopCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightReShopCreateResponse) SetBody(v *IntlFlightReShopCreateResponseBody) *IntlFlightReShopCreateResponse {
	s.Body = v
	return s
}

func (s *IntlFlightReShopCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
