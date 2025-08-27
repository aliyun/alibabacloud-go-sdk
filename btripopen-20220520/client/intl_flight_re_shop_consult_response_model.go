// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopConsultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IntlFlightReShopConsultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IntlFlightReShopConsultResponse
	GetStatusCode() *int32
	SetBody(v *IntlFlightReShopConsultResponseBody) *IntlFlightReShopConsultResponse
	GetBody() *IntlFlightReShopConsultResponseBody
}

type IntlFlightReShopConsultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IntlFlightReShopConsultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IntlFlightReShopConsultResponse) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopConsultResponse) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopConsultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IntlFlightReShopConsultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IntlFlightReShopConsultResponse) GetBody() *IntlFlightReShopConsultResponseBody {
	return s.Body
}

func (s *IntlFlightReShopConsultResponse) SetHeaders(v map[string]*string) *IntlFlightReShopConsultResponse {
	s.Headers = v
	return s
}

func (s *IntlFlightReShopConsultResponse) SetStatusCode(v int32) *IntlFlightReShopConsultResponse {
	s.StatusCode = &v
	return s
}

func (s *IntlFlightReShopConsultResponse) SetBody(v *IntlFlightReShopConsultResponseBody) *IntlFlightReShopConsultResponse {
	s.Body = v
	return s
}

func (s *IntlFlightReShopConsultResponse) Validate() error {
	return dara.Validate(s)
}
