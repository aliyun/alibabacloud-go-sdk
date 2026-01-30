// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTier2CouponApprovalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationSheetId(v string) *ListTier2CouponApprovalRequest
	GetApplicationSheetId() *string
	SetApprovalStatus(v string) *ListTier2CouponApprovalRequest
	GetApprovalStatus() *string
	SetCurrentPage(v int32) *ListTier2CouponApprovalRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListTier2CouponApprovalRequest
	GetPageSize() *int32
	SetT2PartnerName(v string) *ListTier2CouponApprovalRequest
	GetT2PartnerName() *string
	SetT2PartnerUid(v int64) *ListTier2CouponApprovalRequest
	GetT2PartnerUid() *int64
}

type ListTier2CouponApprovalRequest struct {
	// example:
	//
	// d54ca949-9b88-4514-add3-c6029c4027f4
	ApplicationSheetId *string `json:"ApplicationSheetId,omitempty" xml:"ApplicationSheetId,omitempty"`
	// example:
	//
	// 1
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// test account
	T2PartnerName *string `json:"T2PartnerName,omitempty" xml:"T2PartnerName,omitempty"`
	// example:
	//
	// 5248516956402795
	T2PartnerUid *int64 `json:"T2PartnerUid,omitempty" xml:"T2PartnerUid,omitempty"`
}

func (s ListTier2CouponApprovalRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTier2CouponApprovalRequest) GoString() string {
	return s.String()
}

func (s *ListTier2CouponApprovalRequest) GetApplicationSheetId() *string {
	return s.ApplicationSheetId
}

func (s *ListTier2CouponApprovalRequest) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *ListTier2CouponApprovalRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTier2CouponApprovalRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTier2CouponApprovalRequest) GetT2PartnerName() *string {
	return s.T2PartnerName
}

func (s *ListTier2CouponApprovalRequest) GetT2PartnerUid() *int64 {
	return s.T2PartnerUid
}

func (s *ListTier2CouponApprovalRequest) SetApplicationSheetId(v string) *ListTier2CouponApprovalRequest {
	s.ApplicationSheetId = &v
	return s
}

func (s *ListTier2CouponApprovalRequest) SetApprovalStatus(v string) *ListTier2CouponApprovalRequest {
	s.ApprovalStatus = &v
	return s
}

func (s *ListTier2CouponApprovalRequest) SetCurrentPage(v int32) *ListTier2CouponApprovalRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTier2CouponApprovalRequest) SetPageSize(v int32) *ListTier2CouponApprovalRequest {
	s.PageSize = &v
	return s
}

func (s *ListTier2CouponApprovalRequest) SetT2PartnerName(v string) *ListTier2CouponApprovalRequest {
	s.T2PartnerName = &v
	return s
}

func (s *ListTier2CouponApprovalRequest) SetT2PartnerUid(v int64) *ListTier2CouponApprovalRequest {
	s.T2PartnerUid = &v
	return s
}

func (s *ListTier2CouponApprovalRequest) Validate() error {
	return dara.Validate(s)
}
