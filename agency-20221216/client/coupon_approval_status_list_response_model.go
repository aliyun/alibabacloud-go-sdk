// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCouponApprovalStatusListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CouponApprovalStatusListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CouponApprovalStatusListResponse
	GetStatusCode() *int32
	SetBody(v *CouponApprovalStatusListResponseBody) *CouponApprovalStatusListResponse
	GetBody() *CouponApprovalStatusListResponseBody
}

type CouponApprovalStatusListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CouponApprovalStatusListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CouponApprovalStatusListResponse) String() string {
	return dara.Prettify(s)
}

func (s CouponApprovalStatusListResponse) GoString() string {
	return s.String()
}

func (s *CouponApprovalStatusListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CouponApprovalStatusListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CouponApprovalStatusListResponse) GetBody() *CouponApprovalStatusListResponseBody {
	return s.Body
}

func (s *CouponApprovalStatusListResponse) SetHeaders(v map[string]*string) *CouponApprovalStatusListResponse {
	s.Headers = v
	return s
}

func (s *CouponApprovalStatusListResponse) SetStatusCode(v int32) *CouponApprovalStatusListResponse {
	s.StatusCode = &v
	return s
}

func (s *CouponApprovalStatusListResponse) SetBody(v *CouponApprovalStatusListResponseBody) *CouponApprovalStatusListResponse {
	s.Body = v
	return s
}

func (s *CouponApprovalStatusListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
