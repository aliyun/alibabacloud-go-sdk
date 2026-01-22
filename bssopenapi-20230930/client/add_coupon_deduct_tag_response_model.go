// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCouponDeductTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCouponDeductTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCouponDeductTagResponse
	GetStatusCode() *int32
	SetBody(v *AddCouponDeductTagResponseBody) *AddCouponDeductTagResponse
	GetBody() *AddCouponDeductTagResponseBody
}

type AddCouponDeductTagResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCouponDeductTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCouponDeductTagResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCouponDeductTagResponse) GoString() string {
	return s.String()
}

func (s *AddCouponDeductTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCouponDeductTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCouponDeductTagResponse) GetBody() *AddCouponDeductTagResponseBody {
	return s.Body
}

func (s *AddCouponDeductTagResponse) SetHeaders(v map[string]*string) *AddCouponDeductTagResponse {
	s.Headers = v
	return s
}

func (s *AddCouponDeductTagResponse) SetStatusCode(v int32) *AddCouponDeductTagResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCouponDeductTagResponse) SetBody(v *AddCouponDeductTagResponseBody) *AddCouponDeductTagResponse {
	s.Body = v
	return s
}

func (s *AddCouponDeductTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
