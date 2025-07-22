// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCouponDeductTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCouponDeductTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCouponDeductTagResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCouponDeductTagResponseBody) *DeleteCouponDeductTagResponse
	GetBody() *DeleteCouponDeductTagResponseBody
}

type DeleteCouponDeductTagResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCouponDeductTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCouponDeductTagResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCouponDeductTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteCouponDeductTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCouponDeductTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCouponDeductTagResponse) GetBody() *DeleteCouponDeductTagResponseBody {
	return s.Body
}

func (s *DeleteCouponDeductTagResponse) SetHeaders(v map[string]*string) *DeleteCouponDeductTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteCouponDeductTagResponse) SetStatusCode(v int32) *DeleteCouponDeductTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCouponDeductTagResponse) SetBody(v *DeleteCouponDeductTagResponseBody) *DeleteCouponDeductTagResponse {
	s.Body = v
	return s
}

func (s *DeleteCouponDeductTagResponse) Validate() error {
	return dara.Validate(s)
}
