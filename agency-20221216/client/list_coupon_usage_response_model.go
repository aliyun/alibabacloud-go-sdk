// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCouponUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCouponUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCouponUsageResponse
	GetStatusCode() *int32
	SetBody(v *ListCouponUsageResponseBody) *ListCouponUsageResponse
	GetBody() *ListCouponUsageResponseBody
}

type ListCouponUsageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCouponUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCouponUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCouponUsageResponse) GoString() string {
	return s.String()
}

func (s *ListCouponUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCouponUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCouponUsageResponse) GetBody() *ListCouponUsageResponseBody {
	return s.Body
}

func (s *ListCouponUsageResponse) SetHeaders(v map[string]*string) *ListCouponUsageResponse {
	s.Headers = v
	return s
}

func (s *ListCouponUsageResponse) SetStatusCode(v int32) *ListCouponUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCouponUsageResponse) SetBody(v *ListCouponUsageResponseBody) *ListCouponUsageResponse {
	s.Body = v
	return s
}

func (s *ListCouponUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
