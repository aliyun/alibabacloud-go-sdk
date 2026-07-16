// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopListSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightReShopListSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightReShopListSearchResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightReShopListSearchResponseBody) *IntlFlightReShopListSearchResponse
	GetBody() *IntlFlightReShopListSearchResponseBody
}

type IntlFlightReShopListSearchResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightReShopListSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightReShopListSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightReShopListSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightReShopListSearchResponse) GetBody() *IntlFlightReShopListSearchResponseBody {
	return s.Body
}

func (s *IntlFlightReShopListSearchResponse) SetHeaders(v map[string]*string) *IntlFlightReShopListSearchResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightReShopListSearchResponse) SetStatusCode(v int32) *IntlFlightReShopListSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponse) SetBody(v *IntlFlightReShopListSearchResponseBody) *IntlFlightReShopListSearchResponse {
	s.Body = v
	return s
}

func (s *IntlFlightReShopListSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
