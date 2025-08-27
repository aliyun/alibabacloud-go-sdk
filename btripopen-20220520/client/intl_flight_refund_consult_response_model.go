// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundConsultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightRefundConsultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightRefundConsultResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightRefundConsultResponseBody) *IntlFlightRefundConsultResponse
	GetBody() *IntlFlightRefundConsultResponseBody
}

type IntlFlightRefundConsultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightRefundConsultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightRefundConsultResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundConsultResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundConsultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightRefundConsultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightRefundConsultResponse) GetBody() *IntlFlightRefundConsultResponseBody {
	return s.Body
}

func (s *IntlFlightRefundConsultResponse) SetHeaders(v map[string]*string) *IntlFlightRefundConsultResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightRefundConsultResponse) SetStatusCode(v int32) *IntlFlightRefundConsultResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightRefundConsultResponse) SetBody(v *IntlFlightRefundConsultResponseBody) *IntlFlightRefundConsultResponse {
	s.Body = v
	return s
}

func (s *IntlFlightRefundConsultResponse) Validate() error {
	return dara.Validate(s)
}
