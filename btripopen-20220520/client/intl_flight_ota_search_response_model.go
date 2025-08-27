// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOtaSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightOtaSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightOtaSearchResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightOtaSearchResponseBody) *IntlFlightOtaSearchResponse
	GetBody() *IntlFlightOtaSearchResponseBody
}

type IntlFlightOtaSearchResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightOtaSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightOtaSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightOtaSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightOtaSearchResponse) GetBody() *IntlFlightOtaSearchResponseBody {
	return s.Body
}

func (s *IntlFlightOtaSearchResponse) SetHeaders(v map[string]*string) *IntlFlightOtaSearchResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightOtaSearchResponse) SetStatusCode(v int32) *IntlFlightOtaSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightOtaSearchResponse) SetBody(v *IntlFlightOtaSearchResponseBody) *IntlFlightOtaSearchResponse {
	s.Body = v
	return s
}

func (s *IntlFlightOtaSearchResponse) Validate() error {
	return dara.Validate(s)
}
