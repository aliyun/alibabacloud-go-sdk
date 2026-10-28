// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCouponUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCouponUsageResponseBody
	GetCode() *string
	SetData(v []*ListCouponUsageResponseBodyData) *ListCouponUsageResponseBody
	GetData() []*ListCouponUsageResponseBodyData
	SetMessage(v string) *ListCouponUsageResponseBody
	GetMessage() *string
	SetPageInfo(v *ListCouponUsageResponseBodyPageInfo) *ListCouponUsageResponseBody
	GetPageInfo() *ListCouponUsageResponseBodyPageInfo
	SetRequestId(v string) *ListCouponUsageResponseBody
	GetRequestId() *string
}

type ListCouponUsageResponseBody struct {
	// The status code of the API operation.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListCouponUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The message returned by the API operation.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListCouponUsageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID. You can provide this ID for troubleshooting.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCouponUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCouponUsageResponseBody) GoString() string {
	return s.String()
}

func (s *ListCouponUsageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCouponUsageResponseBody) GetData() []*ListCouponUsageResponseBodyData {
	return s.Data
}

func (s *ListCouponUsageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCouponUsageResponseBody) GetPageInfo() *ListCouponUsageResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListCouponUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCouponUsageResponseBody) SetCode(v string) *ListCouponUsageResponseBody {
	s.Code = &v
	return s
}

func (s *ListCouponUsageResponseBody) SetData(v []*ListCouponUsageResponseBodyData) *ListCouponUsageResponseBody {
	s.Data = v
	return s
}

func (s *ListCouponUsageResponseBody) SetMessage(v string) *ListCouponUsageResponseBody {
	s.Message = &v
	return s
}

func (s *ListCouponUsageResponseBody) SetPageInfo(v *ListCouponUsageResponseBodyPageInfo) *ListCouponUsageResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListCouponUsageResponseBody) SetRequestId(v string) *ListCouponUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCouponUsageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCouponUsageResponseBodyData struct {
	// The Alibaba Cloud account.
	//
	// example:
	//
	// oqevfbveuadcrduzmf@ttirv.net
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The face value of the coupon.
	//
	// example:
	//
	// 200
	Amount *float64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The remaining balance.
	//
	// example:
	//
	// 0.01
	Balance *float64 `json:"Balance,omitempty" xml:"Balance,omitempty"`
	// The ID of the coupon.
	//
	// example:
	//
	// 59226280
	CouponId *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// The ID of the coupon template.
	//
	// example:
	//
	// 503802
	CouponTemplateId *int64 `json:"CouponTemplateId,omitempty" xml:"CouponTemplateId,omitempty"`
	// The validity period of the coupon.
	//
	// example:
	//
	// 2023-04-06 00:00:00 ~ 2023-04-07 00:00:00
	EffDate *string `json:"EffDate,omitempty" xml:"EffDate,omitempty"`
	// The time when the coupon was issued. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-04-06 19:32:10
	PublishDate *string `json:"PublishDate,omitempty" xml:"PublishDate,omitempty"`
	// The coupon status. Valid values:</br>
	//
	// AVAILABLE: Normal.</br>
	//
	// EXHAUSTED: Used up.</br>
	//
	// EXPIRED: Expired.</br>
	//
	// ABANDONED: Voided.</br>
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UID of the T2 partner.
	//
	// example:
	//
	// 5248516956402795
	T2PartnerUid *string `json:"T2PartnerUid,omitempty" xml:"T2PartnerUid,omitempty"`
	// The UID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1647668856741998
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListCouponUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCouponUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCouponUsageResponseBodyData) GetAccount() *string {
	return s.Account
}

func (s *ListCouponUsageResponseBodyData) GetAmount() *float64 {
	return s.Amount
}

func (s *ListCouponUsageResponseBodyData) GetBalance() *float64 {
	return s.Balance
}

func (s *ListCouponUsageResponseBodyData) GetCouponId() *string {
	return s.CouponId
}

func (s *ListCouponUsageResponseBodyData) GetCouponTemplateId() *int64 {
	return s.CouponTemplateId
}

func (s *ListCouponUsageResponseBodyData) GetEffDate() *string {
	return s.EffDate
}

func (s *ListCouponUsageResponseBodyData) GetPublishDate() *string {
	return s.PublishDate
}

func (s *ListCouponUsageResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListCouponUsageResponseBodyData) GetT2PartnerUid() *string {
	return s.T2PartnerUid
}

func (s *ListCouponUsageResponseBodyData) GetUid() *int64 {
	return s.Uid
}

func (s *ListCouponUsageResponseBodyData) SetAccount(v string) *ListCouponUsageResponseBodyData {
	s.Account = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetAmount(v float64) *ListCouponUsageResponseBodyData {
	s.Amount = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetBalance(v float64) *ListCouponUsageResponseBodyData {
	s.Balance = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetCouponId(v string) *ListCouponUsageResponseBodyData {
	s.CouponId = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetCouponTemplateId(v int64) *ListCouponUsageResponseBodyData {
	s.CouponTemplateId = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetEffDate(v string) *ListCouponUsageResponseBodyData {
	s.EffDate = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetPublishDate(v string) *ListCouponUsageResponseBodyData {
	s.PublishDate = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetStatus(v string) *ListCouponUsageResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetT2PartnerUid(v string) *ListCouponUsageResponseBodyData {
	s.T2PartnerUid = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) SetUid(v int64) *ListCouponUsageResponseBodyData {
	s.Uid = &v
	return s
}

func (s *ListCouponUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListCouponUsageResponseBodyPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 300
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCouponUsageResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCouponUsageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListCouponUsageResponseBodyPageInfo) GetPage() *int32 {
	return s.Page
}

func (s *ListCouponUsageResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCouponUsageResponseBodyPageInfo) GetTotal() *int32 {
	return s.Total
}

func (s *ListCouponUsageResponseBodyPageInfo) SetPage(v int32) *ListCouponUsageResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *ListCouponUsageResponseBodyPageInfo) SetPageSize(v int32) *ListCouponUsageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCouponUsageResponseBodyPageInfo) SetTotal(v int32) *ListCouponUsageResponseBodyPageInfo {
	s.Total = &v
	return s
}

func (s *ListCouponUsageResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
