// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightListingSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightListingSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightListingSearchResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightListingSearchResponseBody) *IntlFlightListingSearchResponse
	GetBody() *IntlFlightListingSearchResponseBody
}

type IntlFlightListingSearchResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightListingSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightListingSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightListingSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightListingSearchResponse) GetBody() *IntlFlightListingSearchResponseBody {
	return s.Body
}

func (s *IntlFlightListingSearchResponse) SetHeaders(v map[string]*string) *IntlFlightListingSearchResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightListingSearchResponse) SetStatusCode(v int32) *IntlFlightListingSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightListingSearchResponse) SetBody(v *IntlFlightListingSearchResponseBody) *IntlFlightListingSearchResponse {
	s.Body = v
	return s
}

func (s *IntlFlightListingSearchResponse) Validate() error {
	return dara.Validate(s)
}
