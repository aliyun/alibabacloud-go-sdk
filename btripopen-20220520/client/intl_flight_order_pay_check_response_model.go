// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderPayCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightOrderPayCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightOrderPayCheckResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightOrderPayCheckResponseBody) *IntlFlightOrderPayCheckResponse
	GetBody() *IntlFlightOrderPayCheckResponseBody
}

type IntlFlightOrderPayCheckResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightOrderPayCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightOrderPayCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderPayCheckResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderPayCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightOrderPayCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightOrderPayCheckResponse) GetBody() *IntlFlightOrderPayCheckResponseBody {
	return s.Body
}

func (s *IntlFlightOrderPayCheckResponse) SetHeaders(v map[string]*string) *IntlFlightOrderPayCheckResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightOrderPayCheckResponse) SetStatusCode(v int32) *IntlFlightOrderPayCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightOrderPayCheckResponse) SetBody(v *IntlFlightOrderPayCheckResponseBody) *IntlFlightOrderPayCheckResponse {
	s.Body = v
	return s
}

func (s *IntlFlightOrderPayCheckResponse) Validate() error {
	return dara.Validate(s)
}
