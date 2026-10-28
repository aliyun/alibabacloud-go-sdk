// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTier2CouponApprovalDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTier2CouponApprovalDetailResponseBody
	GetCode() *string
	SetData(v *GetTier2CouponApprovalDetailResponseBodyData) *GetTier2CouponApprovalDetailResponseBody
	GetData() *GetTier2CouponApprovalDetailResponseBodyData
	SetMessage(v string) *GetTier2CouponApprovalDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTier2CouponApprovalDetailResponseBody
	GetRequestId() *string
}

type GetTier2CouponApprovalDetailResponseBody struct {
	// The return code. A value of 200 indicates success. Other values indicate error codes. For more information, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *GetTier2CouponApprovalDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0b9a079e17691387754512757e6a1b
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTier2CouponApprovalDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTier2CouponApprovalDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetTier2CouponApprovalDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTier2CouponApprovalDetailResponseBody) GetData() *GetTier2CouponApprovalDetailResponseBodyData {
	return s.Data
}

func (s *GetTier2CouponApprovalDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTier2CouponApprovalDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTier2CouponApprovalDetailResponseBody) SetCode(v string) *GetTier2CouponApprovalDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBody) SetData(v *GetTier2CouponApprovalDetailResponseBodyData) *GetTier2CouponApprovalDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBody) SetMessage(v string) *GetTier2CouponApprovalDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBody) SetRequestId(v string) *GetTier2CouponApprovalDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTier2CouponApprovalDetailResponseBodyData struct {
	// The basic information of the application.
	ApplicantInfo *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo `json:"ApplicantInfo,omitempty" xml:"ApplicantInfo,omitempty" type:"Struct"`
	// The reason for the application.
	//
	// example:
	//
	// Reason for application
	ApplicationReason *string `json:"ApplicationReason,omitempty" xml:"ApplicationReason,omitempty"`
	// The ID of the application form.
	//
	// example:
	//
	// 0b9a079e17691387754512757e6a1b
	ApplicationSheetId *string `json:"ApplicationSheetId,omitempty" xml:"ApplicationSheetId,omitempty"`
	// The approval status. Valid values:
	//
	// - 1: Revoked.
	//
	// - 2: Under review.
	//
	// - 3: Rejected.
	//
	// - 4: Approved.
	//
	// - 5: Approved but issuance failed.
	//
	// - 8: Pending.
	//
	// example:
	//
	// 2
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// The list of customer UIDs that receive the coupons.
	CouponReceiptUidList []*GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList `json:"CouponReceiptUidList,omitempty" xml:"CouponReceiptUidList,omitempty" type:"Repeated"`
	// The estimated remaining quota of PoC coupon benefits for new partner customers after the coupons are issued.
	//
	// example:
	//
	// 20.00
	RemainingQuota *string `json:"RemainingQuota,omitempty" xml:"RemainingQuota,omitempty"`
	// The requested amount.
	//
	// example:
	//
	// 100.00
	TotalAmount *string `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
}

func (s GetTier2CouponApprovalDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTier2CouponApprovalDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) GetApplicantInfo() *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo {
	return s.ApplicantInfo
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) GetApplicationReason() *string {
	return s.ApplicationReason
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) GetApplicationSheetId() *string {
	return s.ApplicationSheetId
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) GetCouponReceiptUidList() []*GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList {
	return s.CouponReceiptUidList
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) GetRemainingQuota() *string {
	return s.RemainingQuota
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) GetTotalAmount() *string {
	return s.TotalAmount
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) SetApplicantInfo(v *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) *GetTier2CouponApprovalDetailResponseBodyData {
	s.ApplicantInfo = v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) SetApplicationReason(v string) *GetTier2CouponApprovalDetailResponseBodyData {
	s.ApplicationReason = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) SetApplicationSheetId(v string) *GetTier2CouponApprovalDetailResponseBodyData {
	s.ApplicationSheetId = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) SetApprovalStatus(v string) *GetTier2CouponApprovalDetailResponseBodyData {
	s.ApprovalStatus = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) SetCouponReceiptUidList(v []*GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList) *GetTier2CouponApprovalDetailResponseBodyData {
	s.CouponReceiptUidList = v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) SetRemainingQuota(v string) *GetTier2CouponApprovalDetailResponseBodyData {
	s.RemainingQuota = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) SetTotalAmount(v string) *GetTier2CouponApprovalDetailResponseBodyData {
	s.TotalAmount = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyData) Validate() error {
	if s.ApplicantInfo != nil {
		if err := s.ApplicantInfo.Validate(); err != nil {
			return err
		}
	}
	if s.CouponReceiptUidList != nil {
		for _, item := range s.CouponReceiptUidList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo struct {
	// The applicable products. Valid values:
	//
	// - UNIVERSAL: All products.
	//
	// - A specific commodity code.
	//
	// example:
	//
	// UNIVERSAL
	ApplicableProducts *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	// The timestamp when the application was submitted.
	//
	// example:
	//
	// 2026-01-21 11:24
	ApplicationTime *string `json:"ApplicationTime,omitempty" xml:"ApplicationTime,omitempty"`
	// The order type for which the coupon is used. Valid values:
	//
	// - ALL: Alibaba Cloud subscription.
	//
	// - BILLING: Alibaba Cloud pay-as-you-go bills.
	//
	// - ALL,BILLING: Alibaba Cloud subscription and pay-as-you-go bills.
	//
	// example:
	//
	// ALL,BILLING
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The name of the tier-2 distributor.
	//
	// example:
	//
	// test account
	T2PartnerName *string `json:"T2PartnerName,omitempty" xml:"T2PartnerName,omitempty"`
	// The UID of the tier-2 distributor.
	//
	// example:
	//
	// 5248516846402795
	T2PartnerUid *int64 `json:"T2PartnerUid,omitempty" xml:"T2PartnerUid,omitempty"`
	// The validity period, which can be a number of valid days or a valid time range.
	//
	// example:
	//
	// 2026-01-21 ~ 2199-12-31
	ValidUntil *string `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
}

func (s GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) String() string {
	return dara.Prettify(s)
}

func (s GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) GoString() string {
	return s.String()
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) GetApplicableProducts() *string {
	return s.ApplicableProducts
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) GetApplicationTime() *string {
	return s.ApplicationTime
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) GetOrderType() *string {
	return s.OrderType
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) GetT2PartnerName() *string {
	return s.T2PartnerName
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) GetT2PartnerUid() *int64 {
	return s.T2PartnerUid
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) GetValidUntil() *string {
	return s.ValidUntil
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) SetApplicableProducts(v string) *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo {
	s.ApplicableProducts = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) SetApplicationTime(v string) *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo {
	s.ApplicationTime = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) SetOrderType(v string) *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo {
	s.OrderType = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) SetT2PartnerName(v string) *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo {
	s.T2PartnerName = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) SetT2PartnerUid(v int64) *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo {
	s.T2PartnerUid = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) SetValidUntil(v string) *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo {
	s.ValidUntil = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataApplicantInfo) Validate() error {
	return dara.Validate(s)
}

type GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList struct {
	// The denomination.
	//
	// example:
	//
	// 100.00
	NominalValue *string `json:"NominalValue,omitempty" xml:"NominalValue,omitempty"`
	// The UID of the customer.
	//
	// example:
	//
	// 1703016242044705
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList) String() string {
	return dara.Prettify(s)
}

func (s GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList) GoString() string {
	return s.String()
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList) GetNominalValue() *string {
	return s.NominalValue
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList) GetUid() *int64 {
	return s.Uid
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList) SetNominalValue(v string) *GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList {
	s.NominalValue = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList) SetUid(v int64) *GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList {
	s.Uid = &v
	return s
}

func (s *GetTier2CouponApprovalDetailResponseBodyDataCouponReceiptUidList) Validate() error {
	return dara.Validate(s)
}
