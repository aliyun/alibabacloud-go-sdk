// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIssueCouponForCustomerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IssueCouponForCustomerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IssueCouponForCustomerResponse
	GetStatusCode() *int32
	SetBody(v *IssueCouponForCustomerResponseBody) *IssueCouponForCustomerResponse
	GetBody() *IssueCouponForCustomerResponseBody
}

type IssueCouponForCustomerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IssueCouponForCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IssueCouponForCustomerResponse) String() string {
	return dara.Prettify(s)
}

func (s IssueCouponForCustomerResponse) GoString() string {
	return s.String()
}

func (s *IssueCouponForCustomerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IssueCouponForCustomerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IssueCouponForCustomerResponse) GetBody() *IssueCouponForCustomerResponseBody {
	return s.Body
}

func (s *IssueCouponForCustomerResponse) SetHeaders(v map[string]*string) *IssueCouponForCustomerResponse {
	s.Headers = v
	return s
}

func (s *IssueCouponForCustomerResponse) SetStatusCode(v int32) *IssueCouponForCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *IssueCouponForCustomerResponse) SetBody(v *IssueCouponForCustomerResponseBody) *IssueCouponForCustomerResponse {
	s.Body = v
	return s
}

func (s *IssueCouponForCustomerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
