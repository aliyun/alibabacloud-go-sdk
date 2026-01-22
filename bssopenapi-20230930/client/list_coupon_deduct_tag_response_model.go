// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCouponDeductTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCouponDeductTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCouponDeductTagResponse
	GetStatusCode() *int32
	SetBody(v *ListCouponDeductTagResponseBody) *ListCouponDeductTagResponse
	GetBody() *ListCouponDeductTagResponseBody
}

type ListCouponDeductTagResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCouponDeductTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCouponDeductTagResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCouponDeductTagResponse) GoString() string {
	return s.String()
}

func (s *ListCouponDeductTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCouponDeductTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCouponDeductTagResponse) GetBody() *ListCouponDeductTagResponseBody {
	return s.Body
}

func (s *ListCouponDeductTagResponse) SetHeaders(v map[string]*string) *ListCouponDeductTagResponse {
	s.Headers = v
	return s
}

func (s *ListCouponDeductTagResponse) SetStatusCode(v int32) *ListCouponDeductTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCouponDeductTagResponse) SetBody(v *ListCouponDeductTagResponseBody) *ListCouponDeductTagResponse {
	s.Body = v
	return s
}

func (s *ListCouponDeductTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
