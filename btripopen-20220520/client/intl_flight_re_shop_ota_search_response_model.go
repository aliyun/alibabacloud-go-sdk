// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopOtaSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightReShopOtaSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightReShopOtaSearchResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightReShopOtaSearchResponseBody) *IntlFlightReShopOtaSearchResponse
	GetBody() *IntlFlightReShopOtaSearchResponseBody
}

type IntlFlightReShopOtaSearchResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightReShopOtaSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightReShopOtaSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightReShopOtaSearchResponse) GetBody() *IntlFlightReShopOtaSearchResponseBody {
	return s.Body
}

func (s *IntlFlightReShopOtaSearchResponse) SetHeaders(v map[string]*string) *IntlFlightReShopOtaSearchResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponse) SetStatusCode(v int32) *IntlFlightReShopOtaSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponse) SetBody(v *IntlFlightReShopOtaSearchResponseBody) *IntlFlightReShopOtaSearchResponse {
	s.Body = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
