// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderCancelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightOrderCancelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightOrderCancelResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightOrderCancelResponseBody) *IntlFlightOrderCancelResponse
	GetBody() *IntlFlightOrderCancelResponseBody
}

type IntlFlightOrderCancelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightOrderCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightOrderCancelResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderCancelResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderCancelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightOrderCancelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightOrderCancelResponse) GetBody() *IntlFlightOrderCancelResponseBody {
	return s.Body
}

func (s *IntlFlightOrderCancelResponse) SetHeaders(v map[string]*string) *IntlFlightOrderCancelResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightOrderCancelResponse) SetStatusCode(v int32) *IntlFlightOrderCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightOrderCancelResponse) SetBody(v *IntlFlightOrderCancelResponseBody) *IntlFlightOrderCancelResponse {
	s.Body = v
	return s
}

func (s *IntlFlightOrderCancelResponse) Validate() error {
	return dara.Validate(s)
}
