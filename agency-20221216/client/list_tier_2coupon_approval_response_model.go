// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTier2CouponApprovalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTier2CouponApprovalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTier2CouponApprovalResponse
	GetStatusCode() *int32
	SetBody(v *ListTier2CouponApprovalResponseBody) *ListTier2CouponApprovalResponse
	GetBody() *ListTier2CouponApprovalResponseBody
}

type ListTier2CouponApprovalResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTier2CouponApprovalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTier2CouponApprovalResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTier2CouponApprovalResponse) GoString() string {
	return s.String()
}

func (s *ListTier2CouponApprovalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTier2CouponApprovalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTier2CouponApprovalResponse) GetBody() *ListTier2CouponApprovalResponseBody {
	return s.Body
}

func (s *ListTier2CouponApprovalResponse) SetHeaders(v map[string]*string) *ListTier2CouponApprovalResponse {
	s.Headers = v
	return s
}

func (s *ListTier2CouponApprovalResponse) SetStatusCode(v int32) *ListTier2CouponApprovalResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTier2CouponApprovalResponse) SetBody(v *ListTier2CouponApprovalResponseBody) *ListTier2CouponApprovalResponse {
	s.Body = v
	return s
}

func (s *ListTier2CouponApprovalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
