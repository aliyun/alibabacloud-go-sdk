// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCouponApprovalStatusListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *CouponApprovalStatusListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *CouponApprovalStatusListRequest
	GetPageSize() *int32
	SetTemplateId(v string) *CouponApprovalStatusListRequest
	GetTemplateId() *string
	SetTemplateName(v string) *CouponApprovalStatusListRequest
	GetTemplateName() *string
	SetTemplateStatus(v string) *CouponApprovalStatusListRequest
	GetTemplateStatus() *string
}

type CouponApprovalStatusListRequest struct {
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the coupon template.
	//
	// example:
	//
	// 5093156
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Test coupon template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The status of the coupon template. Valid values:
	//
	// - **1**: application revoked
	//
	// - **2**: under review
	//
	// - **3**: review rejected
	//
	// - **4**: review approved
	//
	// - **5**: review approved but issuance failed
	//
	// - **6**: automatically approved
	//
	// - **7**: automatically approved but issuance failed
	//
	// example:
	//
	// 2
	TemplateStatus *string `json:"TemplateStatus,omitempty" xml:"TemplateStatus,omitempty"`
}

func (s CouponApprovalStatusListRequest) String() string {
	return dara.Prettify(s)
}

func (s CouponApprovalStatusListRequest) GoString() string {
	return s.String()
}

func (s *CouponApprovalStatusListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *CouponApprovalStatusListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CouponApprovalStatusListRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CouponApprovalStatusListRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CouponApprovalStatusListRequest) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *CouponApprovalStatusListRequest) SetPageNo(v int32) *CouponApprovalStatusListRequest {
	s.PageNo = &v
	return s
}

func (s *CouponApprovalStatusListRequest) SetPageSize(v int32) *CouponApprovalStatusListRequest {
	s.PageSize = &v
	return s
}

func (s *CouponApprovalStatusListRequest) SetTemplateId(v string) *CouponApprovalStatusListRequest {
	s.TemplateId = &v
	return s
}

func (s *CouponApprovalStatusListRequest) SetTemplateName(v string) *CouponApprovalStatusListRequest {
	s.TemplateName = &v
	return s
}

func (s *CouponApprovalStatusListRequest) SetTemplateStatus(v string) *CouponApprovalStatusListRequest {
	s.TemplateStatus = &v
	return s
}

func (s *CouponApprovalStatusListRequest) Validate() error {
	return dara.Validate(s)
}
