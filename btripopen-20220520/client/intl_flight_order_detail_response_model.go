// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightOrderDetailResponseBody) *IntlFlightOrderDetailResponse
	GetBody() *IntlFlightOrderDetailResponseBody
}

type IntlFlightOrderDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightOrderDetailResponse) GetBody() *IntlFlightOrderDetailResponseBody {
	return s.Body
}

func (s *IntlFlightOrderDetailResponse) SetHeaders(v map[string]*string) *IntlFlightOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightOrderDetailResponse) SetStatusCode(v int32) *IntlFlightOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightOrderDetailResponse) SetBody(v *IntlFlightOrderDetailResponseBody) *IntlFlightOrderDetailResponse {
	s.Body = v
	return s
}

func (s *IntlFlightOrderDetailResponse) Validate() error {
	return dara.Validate(s)
}
