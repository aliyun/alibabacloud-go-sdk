// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTier2CouponApprovalDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTier2CouponApprovalDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTier2CouponApprovalDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetTier2CouponApprovalDetailResponseBody) *GetTier2CouponApprovalDetailResponse
	GetBody() *GetTier2CouponApprovalDetailResponseBody
}

type GetTier2CouponApprovalDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTier2CouponApprovalDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTier2CouponApprovalDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTier2CouponApprovalDetailResponse) GoString() string {
	return s.String()
}

func (s *GetTier2CouponApprovalDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTier2CouponApprovalDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTier2CouponApprovalDetailResponse) GetBody() *GetTier2CouponApprovalDetailResponseBody {
	return s.Body
}

func (s *GetTier2CouponApprovalDetailResponse) SetHeaders(v map[string]*string) *GetTier2CouponApprovalDetailResponse {
	s.Headers = v
	return s
}

func (s *GetTier2CouponApprovalDetailResponse) SetStatusCode(v int32) *GetTier2CouponApprovalDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponse) SetBody(v *GetTier2CouponApprovalDetailResponseBody) *GetTier2CouponApprovalDetailResponse {
	s.Body = v
	return s
}

func (s *GetTier2CouponApprovalDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
