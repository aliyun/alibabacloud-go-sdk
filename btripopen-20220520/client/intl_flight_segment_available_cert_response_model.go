// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightSegmentAvailableCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightSegmentAvailableCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightSegmentAvailableCertResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightSegmentAvailableCertResponseBody) *IntlFlightSegmentAvailableCertResponse
	GetBody() *IntlFlightSegmentAvailableCertResponseBody
}

type IntlFlightSegmentAvailableCertResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightSegmentAvailableCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightSegmentAvailableCertResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightSegmentAvailableCertResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightSegmentAvailableCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightSegmentAvailableCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightSegmentAvailableCertResponse) GetBody() *IntlFlightSegmentAvailableCertResponseBody {
	return s.Body
}

func (s *IntlFlightSegmentAvailableCertResponse) SetHeaders(v map[string]*string) *IntlFlightSegmentAvailableCertResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponse) SetStatusCode(v int32) *IntlFlightSegmentAvailableCertResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponse) SetBody(v *IntlFlightSegmentAvailableCertResponseBody) *IntlFlightSegmentAvailableCertResponse {
	s.Body = v
	return s
}

func (s *IntlFlightSegmentAvailableCertResponse) Validate() error {
	return dara.Validate(s)
}
