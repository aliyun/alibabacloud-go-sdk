// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCouponResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeCouponResponseBody
	GetCurrentPage() *int32
	SetData(v []*DescribeCouponResponseBodyData) *DescribeCouponResponseBody
	GetData() []*DescribeCouponResponseBodyData
	SetPageSize(v int32) *DescribeCouponResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCouponResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCouponResponseBody
	GetTotalCount() *int32
}

type DescribeCouponResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data list.
	Data []*DescribeCouponResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C880B065-A781-4F19-B6DD-3E0E3B715C64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCouponResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCouponResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCouponResponseBody) GetData() []*DescribeCouponResponseBodyData {
	return s.Data
}

func (s *DescribeCouponResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCouponResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCouponResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCouponResponseBody) SetCurrentPage(v int32) *DescribeCouponResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCouponResponseBody) SetData(v []*DescribeCouponResponseBodyData) *DescribeCouponResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCouponResponseBody) SetPageSize(v int32) *DescribeCouponResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCouponResponseBody) SetRequestId(v string) *DescribeCouponResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCouponResponseBody) SetTotalCount(v int32) *DescribeCouponResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCouponResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCouponResponseBodyData struct {
	// The face value.
	//
	// example:
	//
	// 9929.750000
	Amount        *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	CertainAmount *string `json:"CertainAmount,omitempty" xml:"CertainAmount,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// 59243658
	CouponId *int64 `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// 731074910070
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The coupon type.
	//
	// example:
	//
	// CERTAIN
	CouponType *string `json:"CouponType,omitempty" xml:"CouponType,omitempty"`
	// The coupon type name.
	//
	// example:
	//
	// 满减券
	CouponTypeName *string `json:"CouponTypeName,omitempty" xml:"CouponTypeName,omitempty"`
	// The currency.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The end time.
	//
	// example:
	//
	// 2021-03-06T15:12Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether there is a first purchase restriction.
	FirstBuy *bool `json:"FirstBuy,omitempty" xml:"FirstBuy,omitempty"`
	// The coupon issuance time.
	//
	// example:
	//
	// 2021-03-02T15:12Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The product code list.
	ItemNames []*string `json:"ItemNames,omitempty" xml:"ItemNames,omitempty" type:"Repeated"`
	// The amount limit.
	//
	// example:
	//
	// 无订单金额限制
	MoneyLimit *string `json:"MoneyLimit,omitempty" xml:"MoneyLimit,omitempty"`
	// The order duration limit rule.
	//
	// example:
	//
	// 预付费规则：购买订单时长大于3600s才能使用
	OrderTimeRule *string `json:"OrderTimeRule,omitempty" xml:"OrderTimeRule,omitempty"`
	// The remaining amount.
	//
	// example:
	//
	// 100.00
	RemainAmount *string `json:"RemainAmount,omitempty" xml:"RemainAmount,omitempty"`
	// The remarks.
	//
	// example:
	//
	// 新买28号
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The shared account list.
	ShareUidList []*DescribeCouponResponseBodyDataShareUidList `json:"ShareUidList,omitempty" xml:"ShareUidList,omitempty" type:"Repeated"`
	// Indicates whether to display the tag deduction button.
	//
	// example:
	//
	// true
	ShowSetDeductTagButton *bool `json:"ShowSetDeductTagButton,omitempty" xml:"ShowSetDeductTagButton,omitempty"`
	// The site.
	//
	// example:
	//
	// CHINA
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
	// The site name.
	//
	// example:
	//
	// 官网自营
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2021-03-02T15:12Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The applicable account.
	//
	// example:
	//
	// 1902671110151254
	SuitAccount *string `json:"SuitAccount,omitempty" xml:"SuitAccount,omitempty"`
	// The applicable product type. Valid values: all, which indicates that the coupon is applicable to all products. white, which indicates that the coupon is applicable to specified products. black, which indicates that the coupon is not applicable to specified products.
	//
	// example:
	//
	// all
	SuitItemType *string `json:"SuitItemType,omitempty" xml:"SuitItemType,omitempty"`
	// The coupon applicable scope.
	//
	// example:
	//
	// UNIVERSAL
	UniversalType *string `json:"UniversalType,omitempty" xml:"UniversalType,omitempty"`
	// The list of order types applicable to the coupon.
	YhOrderTypes []*string `json:"YhOrderTypes,omitempty" xml:"YhOrderTypes,omitempty" type:"Repeated"`
}

func (s DescribeCouponResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCouponResponseBodyData) GetAmount() *string {
	return s.Amount
}

func (s *DescribeCouponResponseBodyData) GetCertainAmount() *string {
	return s.CertainAmount
}

func (s *DescribeCouponResponseBodyData) GetCouponId() *int64 {
	return s.CouponId
}

func (s *DescribeCouponResponseBodyData) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribeCouponResponseBodyData) GetCouponType() *string {
	return s.CouponType
}

func (s *DescribeCouponResponseBodyData) GetCouponTypeName() *string {
	return s.CouponTypeName
}

func (s *DescribeCouponResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeCouponResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCouponResponseBodyData) GetFirstBuy() *bool {
	return s.FirstBuy
}

func (s *DescribeCouponResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeCouponResponseBodyData) GetItemNames() []*string {
	return s.ItemNames
}

func (s *DescribeCouponResponseBodyData) GetMoneyLimit() *string {
	return s.MoneyLimit
}

func (s *DescribeCouponResponseBodyData) GetOrderTimeRule() *string {
	return s.OrderTimeRule
}

func (s *DescribeCouponResponseBodyData) GetRemainAmount() *string {
	return s.RemainAmount
}

func (s *DescribeCouponResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCouponResponseBodyData) GetShareUidList() []*DescribeCouponResponseBodyDataShareUidList {
	return s.ShareUidList
}

func (s *DescribeCouponResponseBodyData) GetShowSetDeductTagButton() *bool {
	return s.ShowSetDeductTagButton
}

func (s *DescribeCouponResponseBodyData) GetSite() *string {
	return s.Site
}

func (s *DescribeCouponResponseBodyData) GetSiteName() *string {
	return s.SiteName
}

func (s *DescribeCouponResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCouponResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeCouponResponseBodyData) GetSuitAccount() *string {
	return s.SuitAccount
}

func (s *DescribeCouponResponseBodyData) GetSuitItemType() *string {
	return s.SuitItemType
}

func (s *DescribeCouponResponseBodyData) GetUniversalType() *string {
	return s.UniversalType
}

func (s *DescribeCouponResponseBodyData) GetYhOrderTypes() []*string {
	return s.YhOrderTypes
}

func (s *DescribeCouponResponseBodyData) SetAmount(v string) *DescribeCouponResponseBodyData {
	s.Amount = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetCertainAmount(v string) *DescribeCouponResponseBodyData {
	s.CertainAmount = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetCouponId(v int64) *DescribeCouponResponseBodyData {
	s.CouponId = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetCouponNo(v string) *DescribeCouponResponseBodyData {
	s.CouponNo = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetCouponType(v string) *DescribeCouponResponseBodyData {
	s.CouponType = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetCouponTypeName(v string) *DescribeCouponResponseBodyData {
	s.CouponTypeName = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetCurrency(v string) *DescribeCouponResponseBodyData {
	s.Currency = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetEndTime(v string) *DescribeCouponResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetFirstBuy(v bool) *DescribeCouponResponseBodyData {
	s.FirstBuy = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetGmtCreate(v string) *DescribeCouponResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetItemNames(v []*string) *DescribeCouponResponseBodyData {
	s.ItemNames = v
	return s
}

func (s *DescribeCouponResponseBodyData) SetMoneyLimit(v string) *DescribeCouponResponseBodyData {
	s.MoneyLimit = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetOrderTimeRule(v string) *DescribeCouponResponseBodyData {
	s.OrderTimeRule = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetRemainAmount(v string) *DescribeCouponResponseBodyData {
	s.RemainAmount = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetRemark(v string) *DescribeCouponResponseBodyData {
	s.Remark = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetShareUidList(v []*DescribeCouponResponseBodyDataShareUidList) *DescribeCouponResponseBodyData {
	s.ShareUidList = v
	return s
}

func (s *DescribeCouponResponseBodyData) SetShowSetDeductTagButton(v bool) *DescribeCouponResponseBodyData {
	s.ShowSetDeductTagButton = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetSite(v string) *DescribeCouponResponseBodyData {
	s.Site = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetSiteName(v string) *DescribeCouponResponseBodyData {
	s.SiteName = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetStartTime(v string) *DescribeCouponResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetStatus(v string) *DescribeCouponResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetSuitAccount(v string) *DescribeCouponResponseBodyData {
	s.SuitAccount = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetSuitItemType(v string) *DescribeCouponResponseBodyData {
	s.SuitItemType = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetUniversalType(v string) *DescribeCouponResponseBodyData {
	s.UniversalType = &v
	return s
}

func (s *DescribeCouponResponseBodyData) SetYhOrderTypes(v []*string) *DescribeCouponResponseBodyData {
	s.YhOrderTypes = v
	return s
}

func (s *DescribeCouponResponseBodyData) Validate() error {
	if s.ShareUidList != nil {
		for _, item := range s.ShareUidList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCouponResponseBodyDataShareUidList struct {
	// The Alibaba Cloud UID.
	//
	// example:
	//
	// 1902671110151254
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The Alibaba Cloud account.
	//
	// example:
	//
	// 阿里云计算有限公司
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s DescribeCouponResponseBodyDataShareUidList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponResponseBodyDataShareUidList) GoString() string {
	return s.String()
}

func (s *DescribeCouponResponseBodyDataShareUidList) GetUid() *string {
	return s.Uid
}

func (s *DescribeCouponResponseBodyDataShareUidList) GetUserNick() *string {
	return s.UserNick
}

func (s *DescribeCouponResponseBodyDataShareUidList) SetUid(v string) *DescribeCouponResponseBodyDataShareUidList {
	s.Uid = &v
	return s
}

func (s *DescribeCouponResponseBodyDataShareUidList) SetUserNick(v string) *DescribeCouponResponseBodyDataShareUidList {
	s.UserNick = &v
	return s
}

func (s *DescribeCouponResponseBodyDataShareUidList) Validate() error {
	return dara.Validate(s)
}
