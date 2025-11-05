// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCoupondeductProductCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCoupondeductProductCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCoupondeductProductCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetCoupondeductProductCodeResponseBody) *GetCoupondeductProductCodeResponse
	GetBody() *GetCoupondeductProductCodeResponseBody
}

type GetCoupondeductProductCodeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCoupondeductProductCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCoupondeductProductCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCoupondeductProductCodeResponse) GoString() string {
	return s.String()
}

func (s *GetCoupondeductProductCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCoupondeductProductCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCoupondeductProductCodeResponse) GetBody() *GetCoupondeductProductCodeResponseBody {
	return s.Body
}

func (s *GetCoupondeductProductCodeResponse) SetHeaders(v map[string]*string) *GetCoupondeductProductCodeResponse {
	s.Headers = v
	return s
}

func (s *GetCoupondeductProductCodeResponse) SetStatusCode(v int32) *GetCoupondeductProductCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCoupondeductProductCodeResponse) SetBody(v *GetCoupondeductProductCodeResponseBody) *GetCoupondeductProductCodeResponse {
	s.Body = v
	return s
}

func (s *GetCoupondeductProductCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
