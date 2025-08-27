// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOtaItemDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightOtaItemDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightOtaItemDetailResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightOtaItemDetailResponseBody) *IntlFlightOtaItemDetailResponse
	GetBody() *IntlFlightOtaItemDetailResponseBody
}

type IntlFlightOtaItemDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightOtaItemDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightOtaItemDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightOtaItemDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightOtaItemDetailResponse) GetBody() *IntlFlightOtaItemDetailResponseBody {
	return s.Body
}

func (s *IntlFlightOtaItemDetailResponse) SetHeaders(v map[string]*string) *IntlFlightOtaItemDetailResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightOtaItemDetailResponse) SetStatusCode(v int32) *IntlFlightOtaItemDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponse) SetBody(v *IntlFlightOtaItemDetailResponseBody) *IntlFlightOtaItemDetailResponse {
	s.Body = v
	return s
}

func (s *IntlFlightOtaItemDetailResponse) Validate() error {
	return dara.Validate(s)
}
