// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderPayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightOrderPayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightOrderPayResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightOrderPayResponseBody) *IntlFlightOrderPayResponse
	GetBody() *IntlFlightOrderPayResponseBody
}

type IntlFlightOrderPayResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightOrderPayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightOrderPayResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderPayResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderPayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightOrderPayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightOrderPayResponse) GetBody() *IntlFlightOrderPayResponseBody {
	return s.Body
}

func (s *IntlFlightOrderPayResponse) SetHeaders(v map[string]*string) *IntlFlightOrderPayResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightOrderPayResponse) SetStatusCode(v int32) *IntlFlightOrderPayResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightOrderPayResponse) SetBody(v *IntlFlightOrderPayResponseBody) *IntlFlightOrderPayResponse {
	s.Body = v
	return s
}

func (s *IntlFlightOrderPayResponse) Validate() error {
	return dara.Validate(s)
}
