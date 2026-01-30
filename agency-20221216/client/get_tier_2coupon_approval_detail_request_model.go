// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTier2CouponApprovalDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationSheetId(v string) *GetTier2CouponApprovalDetailRequest
	GetApplicationSheetId() *string
}

type GetTier2CouponApprovalDetailRequest struct {
	// example:
	//
	// d54ca949-9b88-4514-add3-c6029c4027f4
	ApplicationSheetId *string `json:"ApplicationSheetId,omitempty" xml:"ApplicationSheetId,omitempty"`
}

func (s GetTier2CouponApprovalDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTier2CouponApprovalDetailRequest) GoString() string {
	return s.String()
}

func (s *GetTier2CouponApprovalDetailRequest) GetApplicationSheetId() *string {
	return s.ApplicationSheetId
}

func (s *GetTier2CouponApprovalDetailRequest) SetApplicationSheetId(v string) *GetTier2CouponApprovalDetailRequest {
	s.ApplicationSheetId = &v
	return s
}

func (s *GetTier2CouponApprovalDetailRequest) Validate() error {
	return dara.Validate(s)
}
