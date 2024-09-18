// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddressInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 陕西省西安市新城区xxx大厦xx室
	AddressDetail *string `json:"addressDetail,omitempty" xml:"addressDetail,omitempty"`
	// example:
	//
	// 0
	AddressId *int64 `json:"addressId,omitempty" xml:"addressId,omitempty"`
	// example:
	//
	// 610102
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 任先生
	Receiver *string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 182***5674
	ReceiverPhone *string `json:"receiverPhone,omitempty" xml:"receiverPhone,omitempty"`
	// example:
	//
	// 61010212
	TownDivisionCode *string `json:"townDivisionCode,omitempty" xml:"townDivisionCode,omitempty"`
}

func (s AddressInfo) String() string {
	return tea.Prettify(s)
}

func (s AddressInfo) GoString() string {
	return s.String()
}

func (s *AddressInfo) SetAddressDetail(v string) *AddressInfo {
	s.AddressDetail = &v
	return s
}

func (s *AddressInfo) SetAddressId(v int64) *AddressInfo {
	s.AddressId = &v
	return s
}

func (s *AddressInfo) SetDivisionCode(v string) *AddressInfo {
	s.DivisionCode = &v
	return s
}

func (s *AddressInfo) SetReceiver(v string) *AddressInfo {
	s.Receiver = &v
	return s
}

func (s *AddressInfo) SetReceiverPhone(v string) *AddressInfo {
	s.ReceiverPhone = &v
	return s
}

func (s *AddressInfo) SetTownDivisionCode(v string) *AddressInfo {
	s.TownDivisionCode = &v
	return s
}

type ApplyReason struct {
	// example:
	//
	// 403769
	ReasonTextId *int64 `json:"reasonTextId,omitempty" xml:"reasonTextId,omitempty"`
	// example:
	//
	// 不想要了
	ReasonTips *string `json:"reasonTips,omitempty" xml:"reasonTips,omitempty"`
}

func (s ApplyReason) String() string {
	return tea.Prettify(s)
}

func (s ApplyReason) GoString() string {
	return s.String()
}

func (s *ApplyReason) SetReasonTextId(v int64) *ApplyReason {
	s.ReasonTextId = &v
	return s
}

func (s *ApplyReason) SetReasonTips(v string) *ApplyReason {
	s.ReasonTips = &v
	return s
}

type Category struct {
	// example:
	//
	// 201792301
	CategoryId *int64 `json:"categoryId,omitempty" xml:"categoryId,omitempty"`
	// example:
	//
	// false
	IsLeaf *bool `json:"isLeaf,omitempty" xml:"isLeaf,omitempty"`
	// example:
	//
	// 2
	Level *int32 `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// 方便面/拉面/挂面/轻食面速食
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 50016422
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
}

func (s Category) String() string {
	return tea.Prettify(s)
}

func (s Category) GoString() string {
	return s.String()
}

func (s *Category) SetCategoryId(v int64) *Category {
	s.CategoryId = &v
	return s
}

func (s *Category) SetIsLeaf(v bool) *Category {
	s.IsLeaf = &v
	return s
}

func (s *Category) SetLevel(v int32) *Category {
	s.Level = &v
	return s
}

func (s *Category) SetName(v string) *Category {
	s.Name = &v
	return s
}

func (s *Category) SetParentId(v int64) *Category {
	s.ParentId = &v
	return s
}

type CategoryListQuery struct {
	CategoryIds []*int64 `json:"categoryIds,omitempty" xml:"categoryIds,omitempty" type:"Repeated"`
	// example:
	//
	// 5200001
	ParentCategoryId *int64 `json:"parentCategoryId,omitempty" xml:"parentCategoryId,omitempty"`
}

func (s CategoryListQuery) String() string {
	return tea.Prettify(s)
}

func (s CategoryListQuery) GoString() string {
	return s.String()
}

func (s *CategoryListQuery) SetCategoryIds(v []*int64) *CategoryListQuery {
	s.CategoryIds = v
	return s
}

func (s *CategoryListQuery) SetParentCategoryId(v int64) *CategoryListQuery {
	s.ParentCategoryId = &v
	return s
}

type CategoryListResult struct {
	Categories []*Category `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	RequestId  *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CategoryListResult) String() string {
	return tea.Prettify(s)
}

func (s CategoryListResult) GoString() string {
	return s.String()
}

func (s *CategoryListResult) SetCategories(v []*Category) *CategoryListResult {
	s.Categories = v
	return s
}

func (s *CategoryListResult) SetRequestId(v string) *CategoryListResult {
	s.RequestId = &v
	return s
}

type ConfirmDisburseCmd struct {
	// example:
	//
	// 6692****5457
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// 6692****5696
	PurchaseOrderId *string `json:"purchaseOrderId,omitempty" xml:"purchaseOrderId,omitempty"`
}

func (s ConfirmDisburseCmd) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDisburseCmd) GoString() string {
	return s.String()
}

func (s *ConfirmDisburseCmd) SetOrderId(v string) *ConfirmDisburseCmd {
	s.OrderId = &v
	return s
}

func (s *ConfirmDisburseCmd) SetPurchaseOrderId(v string) *ConfirmDisburseCmd {
	s.PurchaseOrderId = &v
	return s
}

type ConfirmDisburseResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ConfirmDisburseResult) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDisburseResult) GoString() string {
	return s.String()
}

func (s *ConfirmDisburseResult) SetRequestId(v string) *ConfirmDisburseResult {
	s.RequestId = &v
	return s
}

func (s *ConfirmDisburseResult) SetResult(v string) *ConfirmDisburseResult {
	s.Result = &v
	return s
}

type CooperationShop struct {
	CooperationCompanyId *string `json:"cooperationCompanyId,omitempty" xml:"cooperationCompanyId,omitempty"`
	CooperationShopId    *string `json:"cooperationShopId,omitempty" xml:"cooperationShopId,omitempty"`
	ShopId               *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
}

func (s CooperationShop) String() string {
	return tea.Prettify(s)
}

func (s CooperationShop) GoString() string {
	return s.String()
}

func (s *CooperationShop) SetCooperationCompanyId(v string) *CooperationShop {
	s.CooperationCompanyId = &v
	return s
}

func (s *CooperationShop) SetCooperationShopId(v string) *CooperationShop {
	s.CooperationShopId = &v
	return s
}

func (s *CooperationShop) SetShopId(v string) *CooperationShop {
	s.ShopId = &v
	return s
}

type CreateAliPayUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
}

func (s CreateAliPayUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAliPayUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateAliPayUrlRequest) SetShopId(v string) *CreateAliPayUrlRequest {
	s.ShopId = &v
	return s
}

type CreateAliPayUrlResult struct {
	AccountId          *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	MemberId           *string `json:"memberId,omitempty" xml:"memberId,omitempty"`
	ZftWithholdSignUrl *string `json:"zftWithholdSignUrl,omitempty" xml:"zftWithholdSignUrl,omitempty"`
}

func (s CreateAliPayUrlResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAliPayUrlResult) GoString() string {
	return s.String()
}

func (s *CreateAliPayUrlResult) SetAccountId(v string) *CreateAliPayUrlResult {
	s.AccountId = &v
	return s
}

func (s *CreateAliPayUrlResult) SetMemberId(v string) *CreateAliPayUrlResult {
	s.MemberId = &v
	return s
}

func (s *CreateAliPayUrlResult) SetZftWithholdSignUrl(v string) *CreateAliPayUrlResult {
	s.ZftWithholdSignUrl = &v
	return s
}

type DeliveryInfo struct {
	// example:
	//
	// 快递 免邮
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 20
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0
	PostFee *int64 `json:"postFee,omitempty" xml:"postFee,omitempty"`
	// example:
	//
	// -4
	ServiceType *int64 `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
}

func (s DeliveryInfo) String() string {
	return tea.Prettify(s)
}

func (s DeliveryInfo) GoString() string {
	return s.String()
}

func (s *DeliveryInfo) SetDisplayName(v string) *DeliveryInfo {
	s.DisplayName = &v
	return s
}

func (s *DeliveryInfo) SetId(v string) *DeliveryInfo {
	s.Id = &v
	return s
}

func (s *DeliveryInfo) SetPostFee(v int64) *DeliveryInfo {
	s.PostFee = &v
	return s
}

func (s *DeliveryInfo) SetServiceType(v int64) *DeliveryInfo {
	s.ServiceType = &v
	return s
}

type DistributionMaxRefundFee struct {
	// example:
	//
	// 100
	MaxRefundFee *int64 `json:"maxRefundFee,omitempty" xml:"maxRefundFee,omitempty"`
	// example:
	//
	// 1
	MinRefundFee *int64 `json:"minRefundFee,omitempty" xml:"minRefundFee,omitempty"`
}

func (s DistributionMaxRefundFee) String() string {
	return tea.Prettify(s)
}

func (s DistributionMaxRefundFee) GoString() string {
	return s.String()
}

func (s *DistributionMaxRefundFee) SetMaxRefundFee(v int64) *DistributionMaxRefundFee {
	s.MaxRefundFee = &v
	return s
}

func (s *DistributionMaxRefundFee) SetMinRefundFee(v int64) *DistributionMaxRefundFee {
	s.MinRefundFee = &v
	return s
}

type Division struct {
	// example:
	//
	// 310000
	DivisionCode *int64 `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// example:
	//
	// 2
	DivisionLevel *int64 `json:"divisionLevel,omitempty" xml:"divisionLevel,omitempty"`
	// example:
	//
	// 上海
	DivisionName *string `json:"divisionName,omitempty" xml:"divisionName,omitempty"`
	// example:
	//
	// 1
	ParentId *int64 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// example:
	//
	// shang hai
	Pinyin *string `json:"pinyin,omitempty" xml:"pinyin,omitempty"`
}

func (s Division) String() string {
	return tea.Prettify(s)
}

func (s Division) GoString() string {
	return s.String()
}

func (s *Division) SetDivisionCode(v int64) *Division {
	s.DivisionCode = &v
	return s
}

func (s *Division) SetDivisionLevel(v int64) *Division {
	s.DivisionLevel = &v
	return s
}

func (s *Division) SetDivisionName(v string) *Division {
	s.DivisionName = &v
	return s
}

func (s *Division) SetParentId(v int64) *Division {
	s.ParentId = &v
	return s
}

func (s *Division) SetPinyin(v string) *Division {
	s.Pinyin = &v
	return s
}

type DivisionPageResult struct {
	DivisionList []*Division `json:"divisionList,omitempty" xml:"divisionList,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DivisionPageResult) String() string {
	return tea.Prettify(s)
}

func (s DivisionPageResult) GoString() string {
	return s.String()
}

func (s *DivisionPageResult) SetDivisionList(v []*Division) *DivisionPageResult {
	s.DivisionList = v
	return s
}

func (s *DivisionPageResult) SetRequestId(v string) *DivisionPageResult {
	s.RequestId = &v
	return s
}

type DivisionQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
}

func (s DivisionQuery) String() string {
	return tea.Prettify(s)
}

func (s DivisionQuery) GoString() string {
	return s.String()
}

func (s *DivisionQuery) SetDivisionCode(v string) *DivisionQuery {
	s.DivisionCode = &v
	return s
}

type GeneralBill struct {
	BillId      *string   `json:"billId,omitempty" xml:"billId,omitempty"`
	BillPeriod  *string   `json:"billPeriod,omitempty" xml:"billPeriod,omitempty"`
	DownloadUrl []*string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty" type:"Repeated"`
	EndTime     *string   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	GmtCreate   *string   `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified *string   `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	ShopId      *string   `json:"shopId,omitempty" xml:"shopId,omitempty"`
	ShopName    *string   `json:"shopName,omitempty" xml:"shopName,omitempty"`
	StartTime   *string   `json:"startTime,omitempty" xml:"startTime,omitempty"`
	TotalAmount *Money    `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
}

func (s GeneralBill) String() string {
	return tea.Prettify(s)
}

func (s GeneralBill) GoString() string {
	return s.String()
}

func (s *GeneralBill) SetBillId(v string) *GeneralBill {
	s.BillId = &v
	return s
}

func (s *GeneralBill) SetBillPeriod(v string) *GeneralBill {
	s.BillPeriod = &v
	return s
}

func (s *GeneralBill) SetDownloadUrl(v []*string) *GeneralBill {
	s.DownloadUrl = v
	return s
}

func (s *GeneralBill) SetEndTime(v string) *GeneralBill {
	s.EndTime = &v
	return s
}

func (s *GeneralBill) SetGmtCreate(v string) *GeneralBill {
	s.GmtCreate = &v
	return s
}

func (s *GeneralBill) SetGmtModified(v string) *GeneralBill {
	s.GmtModified = &v
	return s
}

func (s *GeneralBill) SetShopId(v string) *GeneralBill {
	s.ShopId = &v
	return s
}

func (s *GeneralBill) SetShopName(v string) *GeneralBill {
	s.ShopName = &v
	return s
}

func (s *GeneralBill) SetStartTime(v string) *GeneralBill {
	s.StartTime = &v
	return s
}

func (s *GeneralBill) SetTotalAmount(v *Money) *GeneralBill {
	s.TotalAmount = v
	return s
}

type GeneralBillPageQuery struct {
	Asc            *bool   `json:"asc,omitempty" xml:"asc,omitempty"`
	BillId         *string `json:"billId,omitempty" xml:"billId,omitempty"`
	BillPeriod     *string `json:"billPeriod,omitempty" xml:"billPeriod,omitempty"`
	Limit          *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ShopId         *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	Start          *int32  `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GeneralBillPageQuery) String() string {
	return tea.Prettify(s)
}

func (s GeneralBillPageQuery) GoString() string {
	return s.String()
}

func (s *GeneralBillPageQuery) SetAsc(v bool) *GeneralBillPageQuery {
	s.Asc = &v
	return s
}

func (s *GeneralBillPageQuery) SetBillId(v string) *GeneralBillPageQuery {
	s.BillId = &v
	return s
}

func (s *GeneralBillPageQuery) SetBillPeriod(v string) *GeneralBillPageQuery {
	s.BillPeriod = &v
	return s
}

func (s *GeneralBillPageQuery) SetLimit(v int32) *GeneralBillPageQuery {
	s.Limit = &v
	return s
}

func (s *GeneralBillPageQuery) SetOrderBy(v string) *GeneralBillPageQuery {
	s.OrderBy = &v
	return s
}

func (s *GeneralBillPageQuery) SetOrderDirection(v string) *GeneralBillPageQuery {
	s.OrderDirection = &v
	return s
}

func (s *GeneralBillPageQuery) SetPageNumber(v int32) *GeneralBillPageQuery {
	s.PageNumber = &v
	return s
}

func (s *GeneralBillPageQuery) SetPageSize(v int32) *GeneralBillPageQuery {
	s.PageSize = &v
	return s
}

func (s *GeneralBillPageQuery) SetShopId(v string) *GeneralBillPageQuery {
	s.ShopId = &v
	return s
}

func (s *GeneralBillPageQuery) SetStart(v int32) *GeneralBillPageQuery {
	s.Start = &v
	return s
}

type GeneralBillPageResult struct {
	GeneralBills []*GeneralBill `json:"generalBills,omitempty" xml:"generalBills,omitempty" type:"Repeated"`
	PageNumber   *int32         `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize     *int32         `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GeneralBillPageResult) String() string {
	return tea.Prettify(s)
}

func (s GeneralBillPageResult) GoString() string {
	return s.String()
}

func (s *GeneralBillPageResult) SetGeneralBills(v []*GeneralBill) *GeneralBillPageResult {
	s.GeneralBills = v
	return s
}

func (s *GeneralBillPageResult) SetPageNumber(v int32) *GeneralBillPageResult {
	s.PageNumber = &v
	return s
}

func (s *GeneralBillPageResult) SetPageSize(v int32) *GeneralBillPageResult {
	s.PageSize = &v
	return s
}

func (s *GeneralBillPageResult) SetRequestId(v string) *GeneralBillPageResult {
	s.RequestId = &v
	return s
}

func (s *GeneralBillPageResult) SetTotal(v int32) *GeneralBillPageResult {
	s.Total = &v
	return s
}

type Good struct {
	// example:
	//
	// 儿童学习桌
	GoodName *string `json:"goodName,omitempty" xml:"goodName,omitempty"`
	// example:
	//
	// 6600****6736
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// 1
	Quantity *int32  `json:"quantity,omitempty" xml:"quantity,omitempty"`
	SkuId    *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
	SkuTitle *string `json:"skuTitle,omitempty" xml:"skuTitle,omitempty"`
}

func (s Good) String() string {
	return tea.Prettify(s)
}

func (s Good) GoString() string {
	return s.String()
}

func (s *Good) SetGoodName(v string) *Good {
	s.GoodName = &v
	return s
}

func (s *Good) SetProductId(v string) *Good {
	s.ProductId = &v
	return s
}

func (s *Good) SetQuantity(v int32) *Good {
	s.Quantity = &v
	return s
}

func (s *Good) SetSkuId(v string) *Good {
	s.SkuId = &v
	return s
}

func (s *Good) SetSkuTitle(v string) *Good {
	s.SkuTitle = &v
	return s
}

type GoodsShippingNoticeCreateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// OTHER-其他 POST-中国邮政 EMS-EMS EYB-EMS快递包裹 POSTB-邮政快递包裹 STO-申通快递 YTO-圆通快递 YUNDA-韵达快递 ZJS-宅急送 FEDEX-联邦快递 DBKD-德邦物流 SHQ-华强物流 TN-特能 TAOBAO-淘宝物流 ZTO-中通快递 HTKY-百世快递 TTKDEX-天天快递 SF-顺丰速运 ZTKY-中铁物流 QFKD-全峰快递 JT-极兔物流
	CpCode *string `json:"cpCode,omitempty" xml:"cpCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6693****4352
	DisputeId *string `json:"disputeId,omitempty" xml:"disputeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SF145****4351
	LogisticsNo *string `json:"logisticsNo,omitempty" xml:"logisticsNo,omitempty"`
}

func (s GoodsShippingNoticeCreateCmd) String() string {
	return tea.Prettify(s)
}

func (s GoodsShippingNoticeCreateCmd) GoString() string {
	return s.String()
}

func (s *GoodsShippingNoticeCreateCmd) SetCpCode(v string) *GoodsShippingNoticeCreateCmd {
	s.CpCode = &v
	return s
}

func (s *GoodsShippingNoticeCreateCmd) SetDisputeId(v string) *GoodsShippingNoticeCreateCmd {
	s.DisputeId = &v
	return s
}

func (s *GoodsShippingNoticeCreateCmd) SetLogisticsNo(v string) *GoodsShippingNoticeCreateCmd {
	s.LogisticsNo = &v
	return s
}

type GoodsShippingNoticeCreateResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GoodsShippingNoticeCreateResult) String() string {
	return tea.Prettify(s)
}

func (s GoodsShippingNoticeCreateResult) GoString() string {
	return s.String()
}

func (s *GoodsShippingNoticeCreateResult) SetRequestId(v string) *GoodsShippingNoticeCreateResult {
	s.RequestId = &v
	return s
}

func (s *GoodsShippingNoticeCreateResult) SetResult(v string) *GoodsShippingNoticeCreateResult {
	s.Result = &v
	return s
}

type LeavePictureList struct {
	// example:
	//
	// 退款凭证
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// https://
	Picture *string `json:"picture,omitempty" xml:"picture,omitempty"`
}

func (s LeavePictureList) String() string {
	return tea.Prettify(s)
}

func (s LeavePictureList) GoString() string {
	return s.String()
}

func (s *LeavePictureList) SetDesc(v string) *LeavePictureList {
	s.Desc = &v
	return s
}

func (s *LeavePictureList) SetPicture(v string) *LeavePictureList {
	s.Picture = &v
	return s
}

type LimitRule struct {
	BeginTime *int64  `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	Condcase  *string `json:"condcase,omitempty" xml:"condcase,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	LimitNum  *int32  `json:"limitNum,omitempty" xml:"limitNum,omitempty"`
	RuleType  *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
}

func (s LimitRule) String() string {
	return tea.Prettify(s)
}

func (s LimitRule) GoString() string {
	return s.String()
}

func (s *LimitRule) SetBeginTime(v int64) *LimitRule {
	s.BeginTime = &v
	return s
}

func (s *LimitRule) SetCondcase(v string) *LimitRule {
	s.Condcase = &v
	return s
}

func (s *LimitRule) SetEndTime(v int64) *LimitRule {
	s.EndTime = &v
	return s
}

func (s *LimitRule) SetLimitNum(v int32) *LimitRule {
	s.LimitNum = &v
	return s
}

func (s *LimitRule) SetRuleType(v string) *LimitRule {
	s.RuleType = &v
	return s
}

type LogisticsDetail struct {
	// example:
	//
	// 2023-09-01 00:00:00
	OcurrTimeStr *string `json:"ocurrTimeStr,omitempty" xml:"ocurrTimeStr,omitempty"`
	// example:
	//
	// 已签收
	StanderdDesc *string `json:"standerdDesc,omitempty" xml:"standerdDesc,omitempty"`
}

func (s LogisticsDetail) String() string {
	return tea.Prettify(s)
}

func (s LogisticsDetail) GoString() string {
	return s.String()
}

func (s *LogisticsDetail) SetOcurrTimeStr(v string) *LogisticsDetail {
	s.OcurrTimeStr = &v
	return s
}

func (s *LogisticsDetail) SetStanderdDesc(v string) *LogisticsDetail {
	s.StanderdDesc = &v
	return s
}

type LogisticsInformationData struct {
	LogisticsStatus      *string `json:"logisticsStatus,omitempty" xml:"logisticsStatus,omitempty"`
	ModifiedTime         *string `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	OrderId              *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	OrderLineId          *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
	OuterPurchaseOrderId *string `json:"outerPurchaseOrderId,omitempty" xml:"outerPurchaseOrderId,omitempty"`
	PurchaserId          *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	TrackingCompanyCode  *string `json:"trackingCompanyCode,omitempty" xml:"trackingCompanyCode,omitempty"`
	TrackingCompanyName  *string `json:"trackingCompanyName,omitempty" xml:"trackingCompanyName,omitempty"`
	TrackingNumber       *string `json:"trackingNumber,omitempty" xml:"trackingNumber,omitempty"`
}

func (s LogisticsInformationData) String() string {
	return tea.Prettify(s)
}

func (s LogisticsInformationData) GoString() string {
	return s.String()
}

func (s *LogisticsInformationData) SetLogisticsStatus(v string) *LogisticsInformationData {
	s.LogisticsStatus = &v
	return s
}

func (s *LogisticsInformationData) SetModifiedTime(v string) *LogisticsInformationData {
	s.ModifiedTime = &v
	return s
}

func (s *LogisticsInformationData) SetOrderId(v string) *LogisticsInformationData {
	s.OrderId = &v
	return s
}

func (s *LogisticsInformationData) SetOrderLineId(v string) *LogisticsInformationData {
	s.OrderLineId = &v
	return s
}

func (s *LogisticsInformationData) SetOuterPurchaseOrderId(v string) *LogisticsInformationData {
	s.OuterPurchaseOrderId = &v
	return s
}

func (s *LogisticsInformationData) SetPurchaserId(v string) *LogisticsInformationData {
	s.PurchaserId = &v
	return s
}

func (s *LogisticsInformationData) SetTrackingCompanyCode(v string) *LogisticsInformationData {
	s.TrackingCompanyCode = &v
	return s
}

func (s *LogisticsInformationData) SetTrackingCompanyName(v string) *LogisticsInformationData {
	s.TrackingCompanyName = &v
	return s
}

func (s *LogisticsInformationData) SetTrackingNumber(v string) *LogisticsInformationData {
	s.TrackingNumber = &v
	return s
}

type LogisticsOrderListResult struct {
	LogisticsOrderList []*LogisticsOrderResult `json:"logisticsOrderList,omitempty" xml:"logisticsOrderList,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s LogisticsOrderListResult) String() string {
	return tea.Prettify(s)
}

func (s LogisticsOrderListResult) GoString() string {
	return s.String()
}

func (s *LogisticsOrderListResult) SetLogisticsOrderList(v []*LogisticsOrderResult) *LogisticsOrderListResult {
	s.LogisticsOrderList = v
	return s
}

func (s *LogisticsOrderListResult) SetRequestId(v string) *LogisticsOrderListResult {
	s.RequestId = &v
	return s
}

type LogisticsOrderResult struct {
	// example:
	//
	// 菜鸟
	DataProvider *string `json:"dataProvider,omitempty" xml:"dataProvider,omitempty"`
	// example:
	//
	// 菜鸟裹裹
	DataProviderTitle *string `json:"dataProviderTitle,omitempty" xml:"dataProviderTitle,omitempty"`
	Goods             []*Good `json:"goods,omitempty" xml:"goods,omitempty" type:"Repeated"`
	// example:
	//
	// SF （OTHER-其他 POST-中国邮政 EMS-EMS EYB-EMS快递包裹 POSTB-邮政快递包裹 STO-申通快递 YTO-圆通快递 YUNDA-韵达快递 ZJS-宅急送 FEDEX-联邦快递 DBKD-德邦物流 SHQ-华强物流 TN-特能 TAOBAO-淘宝物流 ZTO-中通快递 HTKY-百世快递 TTKDEX-天天快递 SF-顺丰速运 ZTKY-中铁物流 QFKD-全峰快递 JT-极兔物流）
	LogisticsCompanyCode *string `json:"logisticsCompanyCode,omitempty" xml:"logisticsCompanyCode,omitempty"`
	// example:
	//
	// 顺丰
	LogisticsCompanyName *string            `json:"logisticsCompanyName,omitempty" xml:"logisticsCompanyName,omitempty"`
	LogisticsDetailList  []*LogisticsDetail `json:"logisticsDetailList,omitempty" xml:"logisticsDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// SF234***2345
	MailNo *string `json:"mailNo,omitempty" xml:"mailNo,omitempty"`
}

func (s LogisticsOrderResult) String() string {
	return tea.Prettify(s)
}

func (s LogisticsOrderResult) GoString() string {
	return s.String()
}

func (s *LogisticsOrderResult) SetDataProvider(v string) *LogisticsOrderResult {
	s.DataProvider = &v
	return s
}

func (s *LogisticsOrderResult) SetDataProviderTitle(v string) *LogisticsOrderResult {
	s.DataProviderTitle = &v
	return s
}

func (s *LogisticsOrderResult) SetGoods(v []*Good) *LogisticsOrderResult {
	s.Goods = v
	return s
}

func (s *LogisticsOrderResult) SetLogisticsCompanyCode(v string) *LogisticsOrderResult {
	s.LogisticsCompanyCode = &v
	return s
}

func (s *LogisticsOrderResult) SetLogisticsCompanyName(v string) *LogisticsOrderResult {
	s.LogisticsCompanyName = &v
	return s
}

func (s *LogisticsOrderResult) SetLogisticsDetailList(v []*LogisticsDetail) *LogisticsOrderResult {
	s.LogisticsDetailList = v
	return s
}

func (s *LogisticsOrderResult) SetMailNo(v string) *LogisticsOrderResult {
	s.MailNo = &v
	return s
}

type MemberAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
}

func (s MemberAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountRequest) GoString() string {
	return s.String()
}

func (s *MemberAccountRequest) SetShopId(v string) *MemberAccountRequest {
	s.ShopId = &v
	return s
}

type MemberAccountResult struct {
	// example:
	//
	// "yue***@newburn.cn"
	AccountNo []*string `json:"accountNo,omitempty" xml:"accountNo,omitempty" type:"Repeated"`
	// example:
	//
	// 123
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
}

func (s MemberAccountResult) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountResult) GoString() string {
	return s.String()
}

func (s *MemberAccountResult) SetAccountNo(v []*string) *MemberAccountResult {
	s.AccountNo = v
	return s
}

func (s *MemberAccountResult) SetShopId(v string) *MemberAccountResult {
	s.ShopId = &v
	return s
}

type Money struct {
	Amount         *int64         `json:"amount,omitempty" xml:"amount,omitempty"`
	AmountAsString *string        `json:"amountAsString,omitempty" xml:"amountAsString,omitempty"`
	AmountString   *string        `json:"amountString,omitempty" xml:"amountString,omitempty"`
	Cent           *int64         `json:"cent,omitempty" xml:"cent,omitempty"`
	Currency       *MoneyCurrency `json:"currency,omitempty" xml:"currency,omitempty" type:"Struct"`
	CurrencyCode   *string        `json:"currencyCode,omitempty" xml:"currencyCode,omitempty"`
	Positive       *bool          `json:"positive,omitempty" xml:"positive,omitempty"`
}

func (s Money) String() string {
	return tea.Prettify(s)
}

func (s Money) GoString() string {
	return s.String()
}

func (s *Money) SetAmount(v int64) *Money {
	s.Amount = &v
	return s
}

func (s *Money) SetAmountAsString(v string) *Money {
	s.AmountAsString = &v
	return s
}

func (s *Money) SetAmountString(v string) *Money {
	s.AmountString = &v
	return s
}

func (s *Money) SetCent(v int64) *Money {
	s.Cent = &v
	return s
}

func (s *Money) SetCurrency(v *MoneyCurrency) *Money {
	s.Currency = v
	return s
}

func (s *Money) SetCurrencyCode(v string) *Money {
	s.CurrencyCode = &v
	return s
}

func (s *Money) SetPositive(v bool) *Money {
	s.Positive = &v
	return s
}

type MoneyCurrency struct {
	CurrencyCode          *string `json:"currencyCode,omitempty" xml:"currencyCode,omitempty"`
	DefaultFractionDigits *int32  `json:"defaultFractionDigits,omitempty" xml:"defaultFractionDigits,omitempty"`
	DisplayName           *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	NumericCode           *int32  `json:"numericCode,omitempty" xml:"numericCode,omitempty"`
	Symbol                *string `json:"symbol,omitempty" xml:"symbol,omitempty"`
}

func (s MoneyCurrency) String() string {
	return tea.Prettify(s)
}

func (s MoneyCurrency) GoString() string {
	return s.String()
}

func (s *MoneyCurrency) SetCurrencyCode(v string) *MoneyCurrency {
	s.CurrencyCode = &v
	return s
}

func (s *MoneyCurrency) SetDefaultFractionDigits(v int32) *MoneyCurrency {
	s.DefaultFractionDigits = &v
	return s
}

func (s *MoneyCurrency) SetDisplayName(v string) *MoneyCurrency {
	s.DisplayName = &v
	return s
}

func (s *MoneyCurrency) SetNumericCode(v int32) *MoneyCurrency {
	s.NumericCode = &v
	return s
}

func (s *MoneyCurrency) SetSymbol(v string) *MoneyCurrency {
	s.Symbol = &v
	return s
}

type OrderLineResult struct {
	// example:
	//
	// 1
	LogisticsStatus *string `json:"logisticsStatus,omitempty" xml:"logisticsStatus,omitempty"`
	// example:
	//
	// 1
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
	// example:
	//
	// 6692****5457
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// 6692****5458
	OrderLineId *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
	// example:
	//
	// 1
	OrderLineStatus *string `json:"orderLineStatus,omitempty" xml:"orderLineStatus,omitempty"`
	// example:
	//
	// 100
	PayFee *int64 `json:"payFee,omitempty" xml:"payFee,omitempty"`
	// example:
	//
	// 6600****6736
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// //img.alicdn.com/imgextra/i4/2216003305543/O1CN01bip3Un1qokG0
	ProductPic *string `json:"productPic,omitempty" xml:"productPic,omitempty"`
	// example:
	//
	// 儿童学习桌
	ProductTitle *string `json:"productTitle,omitempty" xml:"productTitle,omitempty"`
	// example:
	//
	// 6600****6737
	SkuId *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
	// example:
	//
	// 浅绿色
	SkuTitle *string `json:"skuTitle,omitempty" xml:"skuTitle,omitempty"`
}

func (s OrderLineResult) String() string {
	return tea.Prettify(s)
}

func (s OrderLineResult) GoString() string {
	return s.String()
}

func (s *OrderLineResult) SetLogisticsStatus(v string) *OrderLineResult {
	s.LogisticsStatus = &v
	return s
}

func (s *OrderLineResult) SetNumber(v string) *OrderLineResult {
	s.Number = &v
	return s
}

func (s *OrderLineResult) SetOrderId(v string) *OrderLineResult {
	s.OrderId = &v
	return s
}

func (s *OrderLineResult) SetOrderLineId(v string) *OrderLineResult {
	s.OrderLineId = &v
	return s
}

func (s *OrderLineResult) SetOrderLineStatus(v string) *OrderLineResult {
	s.OrderLineStatus = &v
	return s
}

func (s *OrderLineResult) SetPayFee(v int64) *OrderLineResult {
	s.PayFee = &v
	return s
}

func (s *OrderLineResult) SetProductId(v string) *OrderLineResult {
	s.ProductId = &v
	return s
}

func (s *OrderLineResult) SetProductPic(v string) *OrderLineResult {
	s.ProductPic = &v
	return s
}

func (s *OrderLineResult) SetProductTitle(v string) *OrderLineResult {
	s.ProductTitle = &v
	return s
}

func (s *OrderLineResult) SetSkuId(v string) *OrderLineResult {
	s.SkuId = &v
	return s
}

func (s *OrderLineResult) SetSkuTitle(v string) *OrderLineResult {
	s.SkuTitle = &v
	return s
}

type OrderListResult struct {
	OrderList []*OrderResult `json:"orderList,omitempty" xml:"orderList,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s OrderListResult) String() string {
	return tea.Prettify(s)
}

func (s OrderListResult) GoString() string {
	return s.String()
}

func (s *OrderListResult) SetOrderList(v []*OrderResult) *OrderListResult {
	s.OrderList = v
	return s
}

func (s *OrderListResult) SetRequestId(v string) *OrderListResult {
	s.RequestId = &v
	return s
}

func (s *OrderListResult) SetTotal(v int32) *OrderListResult {
	s.Total = &v
	return s
}

type OrderPageQuery struct {
	OrderIdList []*string `json:"orderIdList,omitempty" xml:"orderIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 6692****5696
	PurchaseOrderId *string `json:"purchaseOrderId,omitempty" xml:"purchaseOrderId,omitempty"`
}

func (s OrderPageQuery) String() string {
	return tea.Prettify(s)
}

func (s OrderPageQuery) GoString() string {
	return s.String()
}

func (s *OrderPageQuery) SetOrderIdList(v []*string) *OrderPageQuery {
	s.OrderIdList = v
	return s
}

func (s *OrderPageQuery) SetPageNumber(v int32) *OrderPageQuery {
	s.PageNumber = &v
	return s
}

func (s *OrderPageQuery) SetPageSize(v int32) *OrderPageQuery {
	s.PageSize = &v
	return s
}

func (s *OrderPageQuery) SetPurchaseOrderId(v string) *OrderPageQuery {
	s.PurchaseOrderId = &v
	return s
}

type OrderProductResult struct {
	// example:
	//
	// true
	CanSell  *bool                  `json:"canSell,omitempty" xml:"canSell,omitempty"`
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// example:
	//
	// 库存为0
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 100
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 6600****6736
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// //img.alicdn.com/imgextra/i4/2216003305543/O1CN01bip3Un1qokG0
	ProductPicUrl *string `json:"productPicUrl,omitempty" xml:"productPicUrl,omitempty"`
	// example:
	//
	// 儿童学习桌
	ProductTitle *string `json:"productTitle,omitempty" xml:"productTitle,omitempty"`
	// example:
	//
	// 56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// example:
	//
	// 1
	Quantity *int32 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// example:
	//
	// 6600****6737
	SkuId *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
	// example:
	//
	// 浅绿色
	SkuTitle *string `json:"skuTitle,omitempty" xml:"skuTitle,omitempty"`
}

func (s OrderProductResult) String() string {
	return tea.Prettify(s)
}

func (s OrderProductResult) GoString() string {
	return s.String()
}

func (s *OrderProductResult) SetCanSell(v bool) *OrderProductResult {
	s.CanSell = &v
	return s
}

func (s *OrderProductResult) SetFeatures(v map[string]interface{}) *OrderProductResult {
	s.Features = v
	return s
}

func (s *OrderProductResult) SetMessage(v string) *OrderProductResult {
	s.Message = &v
	return s
}

func (s *OrderProductResult) SetPrice(v int64) *OrderProductResult {
	s.Price = &v
	return s
}

func (s *OrderProductResult) SetProductId(v string) *OrderProductResult {
	s.ProductId = &v
	return s
}

func (s *OrderProductResult) SetProductPicUrl(v string) *OrderProductResult {
	s.ProductPicUrl = &v
	return s
}

func (s *OrderProductResult) SetProductTitle(v string) *OrderProductResult {
	s.ProductTitle = &v
	return s
}

func (s *OrderProductResult) SetPurchaserId(v string) *OrderProductResult {
	s.PurchaserId = &v
	return s
}

func (s *OrderProductResult) SetQuantity(v int32) *OrderProductResult {
	s.Quantity = &v
	return s
}

func (s *OrderProductResult) SetSkuId(v string) *OrderProductResult {
	s.SkuId = &v
	return s
}

func (s *OrderProductResult) SetSkuTitle(v string) *OrderProductResult {
	s.SkuTitle = &v
	return s
}

type OrderRenderProductDTO struct {
	// This parameter is required.
	//
	// example:
	//
	// 6600****6736
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6600****6737
	SkuId *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
}

func (s OrderRenderProductDTO) String() string {
	return tea.Prettify(s)
}

func (s OrderRenderProductDTO) GoString() string {
	return s.String()
}

func (s *OrderRenderProductDTO) SetProductId(v string) *OrderRenderProductDTO {
	s.ProductId = &v
	return s
}

func (s *OrderRenderProductDTO) SetPurchaserId(v string) *OrderRenderProductDTO {
	s.PurchaserId = &v
	return s
}

func (s *OrderRenderProductDTO) SetQuantity(v int32) *OrderRenderProductDTO {
	s.Quantity = &v
	return s
}

func (s *OrderRenderProductDTO) SetSkuId(v string) *OrderRenderProductDTO {
	s.SkuId = &v
	return s
}

type OrderRenderResult struct {
	// example:
	//
	// true
	CanSell          *bool                  `json:"canSell,omitempty" xml:"canSell,omitempty"`
	DeliveryInfoList []*DeliveryInfo        `json:"deliveryInfoList,omitempty" xml:"deliveryInfoList,omitempty" type:"Repeated"`
	ExtInfo          map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// example:
	//
	// 库存为0
	Message     *string               `json:"message,omitempty" xml:"message,omitempty"`
	ProductList []*OrderProductResult `json:"productList,omitempty" xml:"productList,omitempty" type:"Repeated"`
}

func (s OrderRenderResult) String() string {
	return tea.Prettify(s)
}

func (s OrderRenderResult) GoString() string {
	return s.String()
}

func (s *OrderRenderResult) SetCanSell(v bool) *OrderRenderResult {
	s.CanSell = &v
	return s
}

func (s *OrderRenderResult) SetDeliveryInfoList(v []*DeliveryInfo) *OrderRenderResult {
	s.DeliveryInfoList = v
	return s
}

func (s *OrderRenderResult) SetExtInfo(v map[string]interface{}) *OrderRenderResult {
	s.ExtInfo = v
	return s
}

func (s *OrderRenderResult) SetMessage(v string) *OrderRenderResult {
	s.Message = &v
	return s
}

func (s *OrderRenderResult) SetProductList(v []*OrderProductResult) *OrderRenderResult {
	s.ProductList = v
	return s
}

type OrderResult struct {
	// example:
	//
	// 2023-09-11T12:22:24.000+08:00
	CreateDate *string `json:"createDate,omitempty" xml:"createDate,omitempty"`
	// example:
	//
	// 12****01
	DistributorId *string `json:"distributorId,omitempty" xml:"distributorId,omitempty"`
	// example:
	//
	// 1
	LogisticsStatus *string `json:"logisticsStatus,omitempty" xml:"logisticsStatus,omitempty"`
	// example:
	//
	// 100
	OrderAmount *int64 `json:"orderAmount,omitempty" xml:"orderAmount,omitempty"`
	// example:
	//
	// 系统关单
	OrderClosedReason *string `json:"orderClosedReason,omitempty" xml:"orderClosedReason,omitempty"`
	// example:
	//
	// 6692****5457
	OrderId       *string            `json:"orderId,omitempty" xml:"orderId,omitempty"`
	OrderLineList []*OrderLineResult `json:"orderLineList,omitempty" xml:"orderLineList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	OrderStatus *string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty"`
	// example:
	//
	// 841471F6-5D61-1331-8C38-2****B55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OrderResult) String() string {
	return tea.Prettify(s)
}

func (s OrderResult) GoString() string {
	return s.String()
}

func (s *OrderResult) SetCreateDate(v string) *OrderResult {
	s.CreateDate = &v
	return s
}

func (s *OrderResult) SetDistributorId(v string) *OrderResult {
	s.DistributorId = &v
	return s
}

func (s *OrderResult) SetLogisticsStatus(v string) *OrderResult {
	s.LogisticsStatus = &v
	return s
}

func (s *OrderResult) SetOrderAmount(v int64) *OrderResult {
	s.OrderAmount = &v
	return s
}

func (s *OrderResult) SetOrderClosedReason(v string) *OrderResult {
	s.OrderClosedReason = &v
	return s
}

func (s *OrderResult) SetOrderId(v string) *OrderResult {
	s.OrderId = &v
	return s
}

func (s *OrderResult) SetOrderLineList(v []*OrderLineResult) *OrderResult {
	s.OrderLineList = v
	return s
}

func (s *OrderResult) SetOrderStatus(v string) *OrderResult {
	s.OrderStatus = &v
	return s
}

func (s *OrderResult) SetRequestId(v string) *OrderResult {
	s.RequestId = &v
	return s
}

type Product struct {
	// example:
	//
	// Apple/苹果
	BrandName *string `json:"brandName,omitempty" xml:"brandName,omitempty"`
	// example:
	//
	// true
	CanSell       *bool       `json:"canSell,omitempty" xml:"canSell,omitempty"`
	CategoryChain []*Category `json:"categoryChain,omitempty" xml:"categoryChain,omitempty" type:"Repeated"`
	// example:
	//
	// 201856501
	CategoryLeafId *int64 `json:"categoryLeafId,omitempty" xml:"categoryLeafId,omitempty"`
	// example:
	//
	// https://img.alicdn.com/descpath/O1CN01wciRDp22AEU14435dsf34
	DescPath *string `json:"descPath,omitempty" xml:"descPath,omitempty"`
	// example:
	//
	// 110000
	DivisionCode     *string                  `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	ExtendProperties []*ProductExtendProperty `json:"extendProperties,omitempty" xml:"extendProperties,omitempty" type:"Repeated"`
	// example:
	//
	// 100+
	FuzzyQuantity *string      `json:"fuzzyQuantity,omitempty" xml:"fuzzyQuantity,omitempty"`
	Images        []*string    `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	LimitRules    []*LimitRule `json:"limitRules,omitempty" xml:"limitRules,omitempty" type:"Repeated"`
	// example:
	//
	// 21000017-4580902812
	LmItemId *string `json:"lmItemId,omitempty" xml:"lmItemId,omitempty"`
	// example:
	//
	// https://img.alicdn.com/imgextra/i3/2214281521988/O1CN01w4vomR1QYYEx6nyr5_!!2214281521988.jpg
	PicUrl *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	// example:
	//
	// 660460842235822080
	ProductId    *string        `json:"productId,omitempty" xml:"productId,omitempty"`
	ProductSpecs []*ProductSpec `json:"productSpecs,omitempty" xml:"productSpecs,omitempty" type:"Repeated"`
	// example:
	//
	// Online
	ProductStatus *string `json:"productStatus,omitempty" xml:"productStatus,omitempty"`
	// example:
	//
	// Normal
	ProductType *string            `json:"productType,omitempty" xml:"productType,omitempty"`
	Properties  []*ProductProperty `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	// example:
	//
	// -1
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 21000017
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	Skus   []*Sku  `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	// example:
	//
	// 100+
	SoldQuantity *string `json:"soldQuantity,omitempty" xml:"soldQuantity,omitempty"`
	// example:
	//
	// 3040203000000000000
	TaxCode *string `json:"taxCode,omitempty" xml:"taxCode,omitempty"`
	// example:
	//
	// 600
	TaxRate *int32 `json:"taxRate,omitempty" xml:"taxRate,omitempty"`
	// example:
	//
	// 发财树
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s Product) String() string {
	return tea.Prettify(s)
}

func (s Product) GoString() string {
	return s.String()
}

func (s *Product) SetBrandName(v string) *Product {
	s.BrandName = &v
	return s
}

func (s *Product) SetCanSell(v bool) *Product {
	s.CanSell = &v
	return s
}

func (s *Product) SetCategoryChain(v []*Category) *Product {
	s.CategoryChain = v
	return s
}

func (s *Product) SetCategoryLeafId(v int64) *Product {
	s.CategoryLeafId = &v
	return s
}

func (s *Product) SetDescPath(v string) *Product {
	s.DescPath = &v
	return s
}

func (s *Product) SetDivisionCode(v string) *Product {
	s.DivisionCode = &v
	return s
}

func (s *Product) SetExtendProperties(v []*ProductExtendProperty) *Product {
	s.ExtendProperties = v
	return s
}

func (s *Product) SetFuzzyQuantity(v string) *Product {
	s.FuzzyQuantity = &v
	return s
}

func (s *Product) SetImages(v []*string) *Product {
	s.Images = v
	return s
}

func (s *Product) SetLimitRules(v []*LimitRule) *Product {
	s.LimitRules = v
	return s
}

func (s *Product) SetLmItemId(v string) *Product {
	s.LmItemId = &v
	return s
}

func (s *Product) SetPicUrl(v string) *Product {
	s.PicUrl = &v
	return s
}

func (s *Product) SetProductId(v string) *Product {
	s.ProductId = &v
	return s
}

func (s *Product) SetProductSpecs(v []*ProductSpec) *Product {
	s.ProductSpecs = v
	return s
}

func (s *Product) SetProductStatus(v string) *Product {
	s.ProductStatus = &v
	return s
}

func (s *Product) SetProductType(v string) *Product {
	s.ProductType = &v
	return s
}

func (s *Product) SetProperties(v []*ProductProperty) *Product {
	s.Properties = v
	return s
}

func (s *Product) SetQuantity(v int64) *Product {
	s.Quantity = &v
	return s
}

func (s *Product) SetRequestId(v string) *Product {
	s.RequestId = &v
	return s
}

func (s *Product) SetShopId(v string) *Product {
	s.ShopId = &v
	return s
}

func (s *Product) SetSkus(v []*Sku) *Product {
	s.Skus = v
	return s
}

func (s *Product) SetSoldQuantity(v string) *Product {
	s.SoldQuantity = &v
	return s
}

func (s *Product) SetTaxCode(v string) *Product {
	s.TaxCode = &v
	return s
}

func (s *Product) SetTaxRate(v int32) *Product {
	s.TaxRate = &v
	return s
}

func (s *Product) SetTitle(v string) *Product {
	s.Title = &v
	return s
}

type ProductDTO struct {
	// example:
	//
	// 100
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6600****6736
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6600****6737
	SkuId *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
}

func (s ProductDTO) String() string {
	return tea.Prettify(s)
}

func (s ProductDTO) GoString() string {
	return s.String()
}

func (s *ProductDTO) SetPrice(v int64) *ProductDTO {
	s.Price = &v
	return s
}

func (s *ProductDTO) SetProductId(v string) *ProductDTO {
	s.ProductId = &v
	return s
}

func (s *ProductDTO) SetPurchaserId(v string) *ProductDTO {
	s.PurchaserId = &v
	return s
}

func (s *ProductDTO) SetQuantity(v int32) *ProductDTO {
	s.Quantity = &v
	return s
}

func (s *ProductDTO) SetSkuId(v string) *ProductDTO {
	s.SkuId = &v
	return s
}

type ProductExtendProperty struct {
	// example:
	//
	// ss_picture_scene
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 场景图
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ProductExtendProperty) String() string {
	return tea.Prettify(s)
}

func (s ProductExtendProperty) GoString() string {
	return s.String()
}

func (s *ProductExtendProperty) SetKey(v string) *ProductExtendProperty {
	s.Key = &v
	return s
}

func (s *ProductExtendProperty) SetValue(v string) *ProductExtendProperty {
	s.Value = &v
	return s
}

type ProductPageResult struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32     `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Products []*Product `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ProductPageResult) String() string {
	return tea.Prettify(s)
}

func (s ProductPageResult) GoString() string {
	return s.String()
}

func (s *ProductPageResult) SetPageNumber(v int32) *ProductPageResult {
	s.PageNumber = &v
	return s
}

func (s *ProductPageResult) SetPageSize(v int32) *ProductPageResult {
	s.PageSize = &v
	return s
}

func (s *ProductPageResult) SetProducts(v []*Product) *ProductPageResult {
	s.Products = v
	return s
}

func (s *ProductPageResult) SetRequestId(v string) *ProductPageResult {
	s.RequestId = &v
	return s
}

func (s *ProductPageResult) SetTotal(v int32) *ProductPageResult {
	s.Total = &v
	return s
}

type ProductPrice struct {
	// example:
	//
	// 120
	FundAmountMoney *string `json:"fundAmountMoney,omitempty" xml:"fundAmountMoney,omitempty"`
}

func (s ProductPrice) String() string {
	return tea.Prettify(s)
}

func (s ProductPrice) GoString() string {
	return s.String()
}

func (s *ProductPrice) SetFundAmountMoney(v string) *ProductPrice {
	s.FundAmountMoney = &v
	return s
}

type ProductProperty struct {
	// example:
	//
	// 颜色
	Text   *string   `json:"text,omitempty" xml:"text,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ProductProperty) String() string {
	return tea.Prettify(s)
}

func (s ProductProperty) GoString() string {
	return s.String()
}

func (s *ProductProperty) SetText(v string) *ProductProperty {
	s.Text = &v
	return s
}

func (s *ProductProperty) SetValues(v []*string) *ProductProperty {
	s.Values = v
	return s
}

type ProductQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 22000009
	DistributorShopId *string `json:"distributorShopId,omitempty" xml:"distributorShopId,omitempty"`
	// example:
	//
	// 110000
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
}

func (s ProductQuery) String() string {
	return tea.Prettify(s)
}

func (s ProductQuery) GoString() string {
	return s.String()
}

func (s *ProductQuery) SetDistributorShopId(v string) *ProductQuery {
	s.DistributorShopId = &v
	return s
}

func (s *ProductQuery) SetDivisionCode(v string) *ProductQuery {
	s.DivisionCode = &v
	return s
}

type ProductSaleInfo struct {
	// example:
	//
	// true
	CanSell *bool `json:"canSell,omitempty" xml:"canSell,omitempty"`
	// example:
	//
	// 110000
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// example:
	//
	// 有货
	FuzzyQuantity *string      `json:"fuzzyQuantity,omitempty" xml:"fuzzyQuantity,omitempty"`
	LimitRules    []*LimitRule `json:"limitRules,omitempty" xml:"limitRules,omitempty" type:"Repeated"`
	// example:
	//
	// 21000017-4580902812
	LmItemId *string `json:"lmItemId,omitempty" xml:"lmItemId,omitempty"`
	// example:
	//
	// 660460842235822080
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// Online
	ProductStatus *string `json:"productStatus,omitempty" xml:"productStatus,omitempty"`
	// example:
	//
	// 10
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 21000017
	ShopId *string        `json:"shopId,omitempty" xml:"shopId,omitempty"`
	Skus   []*SkuSaleInfo `json:"skus,omitempty" xml:"skus,omitempty" type:"Repeated"`
	// example:
	//
	// 发财树
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ProductSaleInfo) String() string {
	return tea.Prettify(s)
}

func (s ProductSaleInfo) GoString() string {
	return s.String()
}

func (s *ProductSaleInfo) SetCanSell(v bool) *ProductSaleInfo {
	s.CanSell = &v
	return s
}

func (s *ProductSaleInfo) SetDivisionCode(v string) *ProductSaleInfo {
	s.DivisionCode = &v
	return s
}

func (s *ProductSaleInfo) SetFuzzyQuantity(v string) *ProductSaleInfo {
	s.FuzzyQuantity = &v
	return s
}

func (s *ProductSaleInfo) SetLimitRules(v []*LimitRule) *ProductSaleInfo {
	s.LimitRules = v
	return s
}

func (s *ProductSaleInfo) SetLmItemId(v string) *ProductSaleInfo {
	s.LmItemId = &v
	return s
}

func (s *ProductSaleInfo) SetProductId(v string) *ProductSaleInfo {
	s.ProductId = &v
	return s
}

func (s *ProductSaleInfo) SetProductStatus(v string) *ProductSaleInfo {
	s.ProductStatus = &v
	return s
}

func (s *ProductSaleInfo) SetQuantity(v int64) *ProductSaleInfo {
	s.Quantity = &v
	return s
}

func (s *ProductSaleInfo) SetRequestId(v string) *ProductSaleInfo {
	s.RequestId = &v
	return s
}

func (s *ProductSaleInfo) SetShopId(v string) *ProductSaleInfo {
	s.ShopId = &v
	return s
}

func (s *ProductSaleInfo) SetSkus(v []*SkuSaleInfo) *ProductSaleInfo {
	s.Skus = v
	return s
}

func (s *ProductSaleInfo) SetTitle(v string) *ProductSaleInfo {
	s.Title = &v
	return s
}

type ProductSaleInfoListQuery struct {
	// example:
	//
	// 110000
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// This parameter is required.
	ProductIds []*string `json:"productIds,omitempty" xml:"productIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 22000009
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
}

func (s ProductSaleInfoListQuery) String() string {
	return tea.Prettify(s)
}

func (s ProductSaleInfoListQuery) GoString() string {
	return s.String()
}

func (s *ProductSaleInfoListQuery) SetDivisionCode(v string) *ProductSaleInfoListQuery {
	s.DivisionCode = &v
	return s
}

func (s *ProductSaleInfoListQuery) SetProductIds(v []*string) *ProductSaleInfoListQuery {
	s.ProductIds = v
	return s
}

func (s *ProductSaleInfoListQuery) SetPurchaserId(v string) *ProductSaleInfoListQuery {
	s.PurchaserId = &v
	return s
}

type ProductSaleInfoListResult struct {
	ProductSaleInfos []*ProductSaleInfo `json:"productSaleInfos,omitempty" xml:"productSaleInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ProductSaleInfoListResult) String() string {
	return tea.Prettify(s)
}

func (s ProductSaleInfoListResult) GoString() string {
	return s.String()
}

func (s *ProductSaleInfoListResult) SetProductSaleInfos(v []*ProductSaleInfo) *ProductSaleInfoListResult {
	s.ProductSaleInfos = v
	return s
}

func (s *ProductSaleInfoListResult) SetRequestId(v string) *ProductSaleInfoListResult {
	s.RequestId = &v
	return s
}

type ProductSaleInfoQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 22000009
	DistributorShopId *string `json:"distributorShopId,omitempty" xml:"distributorShopId,omitempty"`
	// example:
	//
	// 110000
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
}

func (s ProductSaleInfoQuery) String() string {
	return tea.Prettify(s)
}

func (s ProductSaleInfoQuery) GoString() string {
	return s.String()
}

func (s *ProductSaleInfoQuery) SetDistributorShopId(v string) *ProductSaleInfoQuery {
	s.DistributorShopId = &v
	return s
}

func (s *ProductSaleInfoQuery) SetDivisionCode(v string) *ProductSaleInfoQuery {
	s.DivisionCode = &v
	return s
}

type ProductSpec struct {
	// example:
	//
	// 颜色分类
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 1000
	KeyId  *int64              `json:"keyId,omitempty" xml:"keyId,omitempty"`
	Values []*ProductSpecValue `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ProductSpec) String() string {
	return tea.Prettify(s)
}

func (s ProductSpec) GoString() string {
	return s.String()
}

func (s *ProductSpec) SetKey(v string) *ProductSpec {
	s.Key = &v
	return s
}

func (s *ProductSpec) SetKeyId(v int64) *ProductSpec {
	s.KeyId = &v
	return s
}

func (s *ProductSpec) SetValues(v []*ProductSpecValue) *ProductSpec {
	s.Values = v
	return s
}

type ProductSpecValue struct {
	// example:
	//
	// 白色
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// 秘色
	ValueAlias *string `json:"valueAlias,omitempty" xml:"valueAlias,omitempty"`
	// example:
	//
	// 1000
	ValueId *int64 `json:"valueId,omitempty" xml:"valueId,omitempty"`
}

func (s ProductSpecValue) String() string {
	return tea.Prettify(s)
}

func (s ProductSpecValue) GoString() string {
	return s.String()
}

func (s *ProductSpecValue) SetValue(v string) *ProductSpecValue {
	s.Value = &v
	return s
}

func (s *ProductSpecValue) SetValueAlias(v string) *ProductSpecValue {
	s.ValueAlias = &v
	return s
}

func (s *ProductSpecValue) SetValueId(v int64) *ProductSpecValue {
	s.ValueId = &v
	return s
}

type PurchaseOrderCreateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// buyer2138237
	BuyerId *string `json:"buyerId,omitempty" xml:"buyerId,omitempty"`
	// This parameter is required.
	DeliveryAddress *AddressInfo           `json:"deliveryAddress,omitempty" xml:"deliveryAddress,omitempty"`
	ExtInfo         map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// outer123456
	OuterPurchaseOrderId *string `json:"outerPurchaseOrderId,omitempty" xml:"outerPurchaseOrderId,omitempty"`
	// This parameter is required.
	ProductList []*ProductDTO `json:"productList,omitempty" xml:"productList,omitempty" type:"Repeated"`
}

func (s PurchaseOrderCreateCmd) String() string {
	return tea.Prettify(s)
}

func (s PurchaseOrderCreateCmd) GoString() string {
	return s.String()
}

func (s *PurchaseOrderCreateCmd) SetBuyerId(v string) *PurchaseOrderCreateCmd {
	s.BuyerId = &v
	return s
}

func (s *PurchaseOrderCreateCmd) SetDeliveryAddress(v *AddressInfo) *PurchaseOrderCreateCmd {
	s.DeliveryAddress = v
	return s
}

func (s *PurchaseOrderCreateCmd) SetExtInfo(v map[string]interface{}) *PurchaseOrderCreateCmd {
	s.ExtInfo = v
	return s
}

func (s *PurchaseOrderCreateCmd) SetOuterPurchaseOrderId(v string) *PurchaseOrderCreateCmd {
	s.OuterPurchaseOrderId = &v
	return s
}

func (s *PurchaseOrderCreateCmd) SetProductList(v []*ProductDTO) *PurchaseOrderCreateCmd {
	s.ProductList = v
	return s
}

type PurchaseOrderCreateResult struct {
	// example:
	//
	// 6692****5696
	PurchaseOrderId *string `json:"purchaseOrderId,omitempty" xml:"purchaseOrderId,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PurchaseOrderCreateResult) String() string {
	return tea.Prettify(s)
}

func (s PurchaseOrderCreateResult) GoString() string {
	return s.String()
}

func (s *PurchaseOrderCreateResult) SetPurchaseOrderId(v string) *PurchaseOrderCreateResult {
	s.PurchaseOrderId = &v
	return s
}

func (s *PurchaseOrderCreateResult) SetRequestId(v string) *PurchaseOrderCreateResult {
	s.RequestId = &v
	return s
}

type PurchaseOrderRenderQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// test1234567
	BuyerId *string `json:"buyerId,omitempty" xml:"buyerId,omitempty"`
	// This parameter is required.
	DeliveryAddress *AddressInfo `json:"deliveryAddress,omitempty" xml:"deliveryAddress,omitempty"`
	// example:
	//
	// {}
	ExtInfo map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// This parameter is required.
	ProductList []*OrderRenderProductDTO `json:"productList,omitempty" xml:"productList,omitempty" type:"Repeated"`
}

func (s PurchaseOrderRenderQuery) String() string {
	return tea.Prettify(s)
}

func (s PurchaseOrderRenderQuery) GoString() string {
	return s.String()
}

func (s *PurchaseOrderRenderQuery) SetBuyerId(v string) *PurchaseOrderRenderQuery {
	s.BuyerId = &v
	return s
}

func (s *PurchaseOrderRenderQuery) SetDeliveryAddress(v *AddressInfo) *PurchaseOrderRenderQuery {
	s.DeliveryAddress = v
	return s
}

func (s *PurchaseOrderRenderQuery) SetExtInfo(v map[string]interface{}) *PurchaseOrderRenderQuery {
	s.ExtInfo = v
	return s
}

func (s *PurchaseOrderRenderQuery) SetProductList(v []*OrderRenderProductDTO) *PurchaseOrderRenderQuery {
	s.ProductList = v
	return s
}

type PurchaseOrderRenderResult struct {
	AddressList []*AddressInfo `json:"addressList,omitempty" xml:"addressList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	CanSell *bool                  `json:"canSell,omitempty" xml:"canSell,omitempty"`
	ExtInfo map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// example:
	//
	// 库存为0
	Message   *string              `json:"message,omitempty" xml:"message,omitempty"`
	OrderList []*OrderRenderResult `json:"orderList,omitempty" xml:"orderList,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId           *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	UnsellableOrderList []*OrderRenderResult `json:"unsellableOrderList,omitempty" xml:"unsellableOrderList,omitempty" type:"Repeated"`
}

func (s PurchaseOrderRenderResult) String() string {
	return tea.Prettify(s)
}

func (s PurchaseOrderRenderResult) GoString() string {
	return s.String()
}

func (s *PurchaseOrderRenderResult) SetAddressList(v []*AddressInfo) *PurchaseOrderRenderResult {
	s.AddressList = v
	return s
}

func (s *PurchaseOrderRenderResult) SetCanSell(v bool) *PurchaseOrderRenderResult {
	s.CanSell = &v
	return s
}

func (s *PurchaseOrderRenderResult) SetExtInfo(v map[string]interface{}) *PurchaseOrderRenderResult {
	s.ExtInfo = v
	return s
}

func (s *PurchaseOrderRenderResult) SetMessage(v string) *PurchaseOrderRenderResult {
	s.Message = &v
	return s
}

func (s *PurchaseOrderRenderResult) SetOrderList(v []*OrderRenderResult) *PurchaseOrderRenderResult {
	s.OrderList = v
	return s
}

func (s *PurchaseOrderRenderResult) SetRequestId(v string) *PurchaseOrderRenderResult {
	s.RequestId = &v
	return s
}

func (s *PurchaseOrderRenderResult) SetUnsellableOrderList(v []*OrderRenderResult) *PurchaseOrderRenderResult {
	s.UnsellableOrderList = v
	return s
}

type PurchaseOrderStatusResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PurchaseOrderStatusResult) String() string {
	return tea.Prettify(s)
}

func (s PurchaseOrderStatusResult) GoString() string {
	return s.String()
}

func (s *PurchaseOrderStatusResult) SetRequestId(v string) *PurchaseOrderStatusResult {
	s.RequestId = &v
	return s
}

func (s *PurchaseOrderStatusResult) SetStatus(v string) *PurchaseOrderStatusResult {
	s.Status = &v
	return s
}

type RefundFeeData struct {
	// example:
	//
	// 100
	MaxRefundFee *int64 `json:"maxRefundFee,omitempty" xml:"maxRefundFee,omitempty"`
	// example:
	//
	// 1
	MinRefundFee *int64 `json:"minRefundFee,omitempty" xml:"minRefundFee,omitempty"`
}

func (s RefundFeeData) String() string {
	return tea.Prettify(s)
}

func (s RefundFeeData) GoString() string {
	return s.String()
}

func (s *RefundFeeData) SetMaxRefundFee(v int64) *RefundFeeData {
	s.MaxRefundFee = &v
	return s
}

func (s *RefundFeeData) SetMinRefundFee(v int64) *RefundFeeData {
	s.MinRefundFee = &v
	return s
}

type RefundOrderCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 47821
	ApplyReasonTextId *int64 `json:"applyReasonTextId,omitempty" xml:"applyReasonTextId,omitempty"`
	// example:
	//
	// 不想要了
	ApplyReasonTips *string `json:"applyReasonTips,omitempty" xml:"applyReasonTips,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ApplyRefundCount *int32 `json:"applyRefundCount,omitempty" xml:"applyRefundCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	ApplyRefundFee *int64 `json:"applyRefundFee,omitempty" xml:"applyRefundFee,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizClaimType *int32 `json:"bizClaimType,omitempty" xml:"bizClaimType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	GoodsStatus *int32 `json:"goodsStatus,omitempty" xml:"goodsStatus,omitempty"`
	// example:
	//
	// 不想要了
	LeaveMessage      *string             `json:"leaveMessage,omitempty" xml:"leaveMessage,omitempty"`
	LeavePictureLists []*LeavePictureList `json:"leavePictureLists,omitempty" xml:"leavePictureLists,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 6692****5458
	OrderLineId *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
}

func (s RefundOrderCmd) String() string {
	return tea.Prettify(s)
}

func (s RefundOrderCmd) GoString() string {
	return s.String()
}

func (s *RefundOrderCmd) SetApplyReasonTextId(v int64) *RefundOrderCmd {
	s.ApplyReasonTextId = &v
	return s
}

func (s *RefundOrderCmd) SetApplyReasonTips(v string) *RefundOrderCmd {
	s.ApplyReasonTips = &v
	return s
}

func (s *RefundOrderCmd) SetApplyRefundCount(v int32) *RefundOrderCmd {
	s.ApplyRefundCount = &v
	return s
}

func (s *RefundOrderCmd) SetApplyRefundFee(v int64) *RefundOrderCmd {
	s.ApplyRefundFee = &v
	return s
}

func (s *RefundOrderCmd) SetBizClaimType(v int32) *RefundOrderCmd {
	s.BizClaimType = &v
	return s
}

func (s *RefundOrderCmd) SetGoodsStatus(v int32) *RefundOrderCmd {
	s.GoodsStatus = &v
	return s
}

func (s *RefundOrderCmd) SetLeaveMessage(v string) *RefundOrderCmd {
	s.LeaveMessage = &v
	return s
}

func (s *RefundOrderCmd) SetLeavePictureLists(v []*LeavePictureList) *RefundOrderCmd {
	s.LeavePictureLists = v
	return s
}

func (s *RefundOrderCmd) SetOrderLineId(v string) *RefundOrderCmd {
	s.OrderLineId = &v
	return s
}

type RefundOrderResult struct {
	// example:
	//
	// 6693****4352
	DisputeId *string `json:"disputeId,omitempty" xml:"disputeId,omitempty"`
	// example:
	//
	// 1
	DisputeStatus *int32 `json:"disputeStatus,omitempty" xml:"disputeStatus,omitempty"`
	// example:
	//
	// 6692****5458
	OrderLineId *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RefundOrderResult) String() string {
	return tea.Prettify(s)
}

func (s RefundOrderResult) GoString() string {
	return s.String()
}

func (s *RefundOrderResult) SetDisputeId(v string) *RefundOrderResult {
	s.DisputeId = &v
	return s
}

func (s *RefundOrderResult) SetDisputeStatus(v int32) *RefundOrderResult {
	s.DisputeStatus = &v
	return s
}

func (s *RefundOrderResult) SetOrderLineId(v string) *RefundOrderResult {
	s.OrderLineId = &v
	return s
}

func (s *RefundOrderResult) SetRequestId(v string) *RefundOrderResult {
	s.RequestId = &v
	return s
}

type RefundReason struct {
	// example:
	//
	// true
	ProofRequired *bool `json:"proofRequired,omitempty" xml:"proofRequired,omitempty"`
	// example:
	//
	// 47683
	ReasonTextId *string `json:"reasonTextId,omitempty" xml:"reasonTextId,omitempty"`
	// example:
	//
	// 不想要了
	ReasonTips *string `json:"reasonTips,omitempty" xml:"reasonTips,omitempty"`
	// example:
	//
	// true
	RefundDescRequired *bool `json:"refundDescRequired,omitempty" xml:"refundDescRequired,omitempty"`
}

func (s RefundReason) String() string {
	return tea.Prettify(s)
}

func (s RefundReason) GoString() string {
	return s.String()
}

func (s *RefundReason) SetProofRequired(v bool) *RefundReason {
	s.ProofRequired = &v
	return s
}

func (s *RefundReason) SetReasonTextId(v string) *RefundReason {
	s.ReasonTextId = &v
	return s
}

func (s *RefundReason) SetReasonTips(v string) *RefundReason {
	s.ReasonTips = &v
	return s
}

func (s *RefundReason) SetRefundDescRequired(v bool) *RefundReason {
	s.RefundDescRequired = &v
	return s
}

type RefundRenderCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizClaimType *int32 `json:"bizClaimType,omitempty" xml:"bizClaimType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	GoodsStatus *int32 `json:"goodsStatus,omitempty" xml:"goodsStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6692****5458
	OrderLineId *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
}

func (s RefundRenderCmd) String() string {
	return tea.Prettify(s)
}

func (s RefundRenderCmd) GoString() string {
	return s.String()
}

func (s *RefundRenderCmd) SetBizClaimType(v int32) *RefundRenderCmd {
	s.BizClaimType = &v
	return s
}

func (s *RefundRenderCmd) SetGoodsStatus(v int32) *RefundRenderCmd {
	s.GoodsStatus = &v
	return s
}

func (s *RefundRenderCmd) SetOrderLineId(v string) *RefundRenderCmd {
	s.OrderLineId = &v
	return s
}

type RefundRenderResult struct {
	// example:
	//
	// 1
	BizClaimType     *int32                    `json:"bizClaimType,omitempty" xml:"bizClaimType,omitempty"`
	MaxRefundFeeData *DistributionMaxRefundFee `json:"maxRefundFeeData,omitempty" xml:"maxRefundFeeData,omitempty"`
	// example:
	//
	// 6692****5458
	OrderLineId      *string         `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
	RefundReasonList []*RefundReason `json:"refundReasonList,omitempty" xml:"refundReasonList,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RefundRenderResult) String() string {
	return tea.Prettify(s)
}

func (s RefundRenderResult) GoString() string {
	return s.String()
}

func (s *RefundRenderResult) SetBizClaimType(v int32) *RefundRenderResult {
	s.BizClaimType = &v
	return s
}

func (s *RefundRenderResult) SetMaxRefundFeeData(v *DistributionMaxRefundFee) *RefundRenderResult {
	s.MaxRefundFeeData = v
	return s
}

func (s *RefundRenderResult) SetOrderLineId(v string) *RefundRenderResult {
	s.OrderLineId = &v
	return s
}

func (s *RefundRenderResult) SetRefundReasonList(v []*RefundReason) *RefundRenderResult {
	s.RefundReasonList = v
	return s
}

func (s *RefundRenderResult) SetRequestId(v string) *RefundRenderResult {
	s.RequestId = &v
	return s
}

type RefundResult struct {
	// example:
	//
	// 多拍不想要
	ApplyDisputeDesc *string      `json:"applyDisputeDesc,omitempty" xml:"applyDisputeDesc,omitempty"`
	ApplyReason      *ApplyReason `json:"applyReason,omitempty" xml:"applyReason,omitempty"`
	// example:
	//
	// 1
	BizClaimType *int32 `json:"bizClaimType,omitempty" xml:"bizClaimType,omitempty"`
	// example:
	//
	// 2023-09-02T00:00:00.000Z
	DisputeCreateTime *string `json:"disputeCreateTime,omitempty" xml:"disputeCreateTime,omitempty"`
	// example:
	//
	// 多拍不想要
	DisputeDesc *string `json:"disputeDesc,omitempty" xml:"disputeDesc,omitempty"`
	// example:
	//
	// 2023-09-02T12:00:00.000Z
	DisputeEndTime *string `json:"disputeEndTime,omitempty" xml:"disputeEndTime,omitempty"`
	// example:
	//
	// 6693****4352
	DisputeId *string `json:"disputeId,omitempty" xml:"disputeId,omitempty"`
	// example:
	//
	// 1
	DisputeStatus *int32 `json:"disputeStatus,omitempty" xml:"disputeStatus,omitempty"`
	// example:
	//
	// 6692****5457
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// 6692****5458
	OrderLineId *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
	// example:
	//
	// 1
	OrderLogisticsStatus *int32 `json:"orderLogisticsStatus,omitempty" xml:"orderLogisticsStatus,omitempty"`
	// example:
	//
	// 1
	RefundFee     *int64         `json:"refundFee,omitempty" xml:"refundFee,omitempty"`
	RefundFeeData *RefundFeeData `json:"refundFeeData,omitempty" xml:"refundFeeData,omitempty"`
	// example:
	//
	// 阿里云云谷
	RefunderAddress *string `json:"refunderAddress,omitempty" xml:"refunderAddress,omitempty"`
	// example:
	//
	// 赵先生
	RefunderName *string `json:"refunderName,omitempty" xml:"refunderName,omitempty"`
	// example:
	//
	// 182****1334
	RefunderTel *string `json:"refunderTel,omitempty" xml:"refunderTel,omitempty"`
	// example:
	//
	// 331001
	RefunderZipCode *string `json:"refunderZipCode,omitempty" xml:"refunderZipCode,omitempty"`
	// example:
	//
	// 841471F6-5D61-1331-8C38-2****B55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0
	ReturnGoodLogisticsStatus *int32 `json:"returnGoodLogisticsStatus,omitempty" xml:"returnGoodLogisticsStatus,omitempty"`
	// example:
	//
	// 同意退款
	SellerAgreeMsg *string `json:"sellerAgreeMsg,omitempty" xml:"sellerAgreeMsg,omitempty"`
	// example:
	//
	// 不同意退款
	SellerRefuseAgreementMessage *string `json:"sellerRefuseAgreementMessage,omitempty" xml:"sellerRefuseAgreementMessage,omitempty"`
	// example:
	//
	// 商品没问题，买家举证无效
	SellerRefuseReason *string `json:"sellerRefuseReason,omitempty" xml:"sellerRefuseReason,omitempty"`
}

func (s RefundResult) String() string {
	return tea.Prettify(s)
}

func (s RefundResult) GoString() string {
	return s.String()
}

func (s *RefundResult) SetApplyDisputeDesc(v string) *RefundResult {
	s.ApplyDisputeDesc = &v
	return s
}

func (s *RefundResult) SetApplyReason(v *ApplyReason) *RefundResult {
	s.ApplyReason = v
	return s
}

func (s *RefundResult) SetBizClaimType(v int32) *RefundResult {
	s.BizClaimType = &v
	return s
}

func (s *RefundResult) SetDisputeCreateTime(v string) *RefundResult {
	s.DisputeCreateTime = &v
	return s
}

func (s *RefundResult) SetDisputeDesc(v string) *RefundResult {
	s.DisputeDesc = &v
	return s
}

func (s *RefundResult) SetDisputeEndTime(v string) *RefundResult {
	s.DisputeEndTime = &v
	return s
}

func (s *RefundResult) SetDisputeId(v string) *RefundResult {
	s.DisputeId = &v
	return s
}

func (s *RefundResult) SetDisputeStatus(v int32) *RefundResult {
	s.DisputeStatus = &v
	return s
}

func (s *RefundResult) SetOrderId(v string) *RefundResult {
	s.OrderId = &v
	return s
}

func (s *RefundResult) SetOrderLineId(v string) *RefundResult {
	s.OrderLineId = &v
	return s
}

func (s *RefundResult) SetOrderLogisticsStatus(v int32) *RefundResult {
	s.OrderLogisticsStatus = &v
	return s
}

func (s *RefundResult) SetRefundFee(v int64) *RefundResult {
	s.RefundFee = &v
	return s
}

func (s *RefundResult) SetRefundFeeData(v *RefundFeeData) *RefundResult {
	s.RefundFeeData = v
	return s
}

func (s *RefundResult) SetRefunderAddress(v string) *RefundResult {
	s.RefunderAddress = &v
	return s
}

func (s *RefundResult) SetRefunderName(v string) *RefundResult {
	s.RefunderName = &v
	return s
}

func (s *RefundResult) SetRefunderTel(v string) *RefundResult {
	s.RefunderTel = &v
	return s
}

func (s *RefundResult) SetRefunderZipCode(v string) *RefundResult {
	s.RefunderZipCode = &v
	return s
}

func (s *RefundResult) SetRequestId(v string) *RefundResult {
	s.RequestId = &v
	return s
}

func (s *RefundResult) SetReturnGoodLogisticsStatus(v int32) *RefundResult {
	s.ReturnGoodLogisticsStatus = &v
	return s
}

func (s *RefundResult) SetSellerAgreeMsg(v string) *RefundResult {
	s.SellerAgreeMsg = &v
	return s
}

func (s *RefundResult) SetSellerRefuseAgreementMessage(v string) *RefundResult {
	s.SellerRefuseAgreementMessage = &v
	return s
}

func (s *RefundResult) SetSellerRefuseReason(v string) *RefundResult {
	s.SellerRefuseReason = &v
	return s
}

type Shop struct {
	// example:
	//
	// 12***01
	CooperationShops []*CooperationShop `json:"cooperationShops,omitempty" xml:"cooperationShops,omitempty" type:"Repeated"`
	// example:
	//
	// 12****09
	DistributorId *string `json:"distributorId,omitempty" xml:"distributorId,omitempty"`
	// example:
	//
	// 2024-09-09 00:00:00
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// PID56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// example:
	//
	// 48A34399-A72C-1E23-8388-7E63****E927
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 22****09
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// example:
	//
	// 儿童分销店铺
	ShopName *string `json:"shopName,omitempty" xml:"shopName,omitempty"`
	// example:
	//
	// DistributorQYG
	ShopType *string `json:"shopType,omitempty" xml:"shopType,omitempty"`
	// example:
	//
	// 2023-09-09 00:00:00
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// example:
	//
	// Working
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s Shop) String() string {
	return tea.Prettify(s)
}

func (s Shop) GoString() string {
	return s.String()
}

func (s *Shop) SetCooperationShops(v []*CooperationShop) *Shop {
	s.CooperationShops = v
	return s
}

func (s *Shop) SetDistributorId(v string) *Shop {
	s.DistributorId = &v
	return s
}

func (s *Shop) SetEndDate(v string) *Shop {
	s.EndDate = &v
	return s
}

func (s *Shop) SetPurchaserId(v string) *Shop {
	s.PurchaserId = &v
	return s
}

func (s *Shop) SetRequestId(v string) *Shop {
	s.RequestId = &v
	return s
}

func (s *Shop) SetShopId(v string) *Shop {
	s.ShopId = &v
	return s
}

func (s *Shop) SetShopName(v string) *Shop {
	s.ShopName = &v
	return s
}

func (s *Shop) SetShopType(v string) *Shop {
	s.ShopType = &v
	return s
}

func (s *Shop) SetStartDate(v string) *Shop {
	s.StartDate = &v
	return s
}

func (s *Shop) SetStatus(v string) *Shop {
	s.Status = &v
	return s
}

type ShopCreateRequest struct {
	// example:
	//
	// 12344335
	AfterSaleDingTalkId *string `json:"afterSaleDingTalkId,omitempty" xml:"afterSaleDingTalkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 店铺描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12344335
	OperatorDingTalkId *string `json:"operatorDingTalkId,omitempty" xml:"operatorDingTalkId,omitempty"`
	// example:
	//
	// 12344335
	PreSaleDingTalkId *string `json:"preSaleDingTalkId,omitempty" xml:"preSaleDingTalkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 店铺名称
	ShopName *string `json:"shopName,omitempty" xml:"shopName,omitempty"`
}

func (s ShopCreateRequest) String() string {
	return tea.Prettify(s)
}

func (s ShopCreateRequest) GoString() string {
	return s.String()
}

func (s *ShopCreateRequest) SetAfterSaleDingTalkId(v string) *ShopCreateRequest {
	s.AfterSaleDingTalkId = &v
	return s
}

func (s *ShopCreateRequest) SetDescription(v string) *ShopCreateRequest {
	s.Description = &v
	return s
}

func (s *ShopCreateRequest) SetOperatorDingTalkId(v string) *ShopCreateRequest {
	s.OperatorDingTalkId = &v
	return s
}

func (s *ShopCreateRequest) SetPreSaleDingTalkId(v string) *ShopCreateRequest {
	s.PreSaleDingTalkId = &v
	return s
}

func (s *ShopCreateRequest) SetShopName(v string) *ShopCreateRequest {
	s.ShopName = &v
	return s
}

type ShopCreateResult struct {
	// example:
	//
	// 123
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// example:
	//
	// WORKING
	ShopStatus *string `json:"shopStatus,omitempty" xml:"shopStatus,omitempty"`
}

func (s ShopCreateResult) String() string {
	return tea.Prettify(s)
}

func (s ShopCreateResult) GoString() string {
	return s.String()
}

func (s *ShopCreateResult) SetShopId(v string) *ShopCreateResult {
	s.ShopId = &v
	return s
}

func (s *ShopCreateResult) SetShopStatus(v string) *ShopCreateResult {
	s.ShopStatus = &v
	return s
}

type ShopPageDataResult struct {
	// example:
	//
	// 12****01
	CooperationShops []*CooperationShop `json:"cooperationShops,omitempty" xml:"cooperationShops,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-09-01 00:00:00
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// PID56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// example:
	//
	// 22****09
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// example:
	//
	// 儿童座椅分销店铺
	ShopName *string `json:"shopName,omitempty" xml:"shopName,omitempty"`
	// example:
	//
	// DistributorQYG
	ShopType *string `json:"shopType,omitempty" xml:"shopType,omitempty"`
	// example:
	//
	// 2023-09-01 00:00:00
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// example:
	//
	// Working
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ShopPageDataResult) String() string {
	return tea.Prettify(s)
}

func (s ShopPageDataResult) GoString() string {
	return s.String()
}

func (s *ShopPageDataResult) SetCooperationShops(v []*CooperationShop) *ShopPageDataResult {
	s.CooperationShops = v
	return s
}

func (s *ShopPageDataResult) SetEndDate(v string) *ShopPageDataResult {
	s.EndDate = &v
	return s
}

func (s *ShopPageDataResult) SetPurchaserId(v string) *ShopPageDataResult {
	s.PurchaserId = &v
	return s
}

func (s *ShopPageDataResult) SetShopId(v string) *ShopPageDataResult {
	s.ShopId = &v
	return s
}

func (s *ShopPageDataResult) SetShopName(v string) *ShopPageDataResult {
	s.ShopName = &v
	return s
}

func (s *ShopPageDataResult) SetShopType(v string) *ShopPageDataResult {
	s.ShopType = &v
	return s
}

func (s *ShopPageDataResult) SetStartDate(v string) *ShopPageDataResult {
	s.StartDate = &v
	return s
}

func (s *ShopPageDataResult) SetStatus(v string) *ShopPageDataResult {
	s.Status = &v
	return s
}

type ShopPageResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ShopList  []*ShopPageDataResult `json:"shopList,omitempty" xml:"shopList,omitempty" type:"Repeated"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ShopPageResult) String() string {
	return tea.Prettify(s)
}

func (s ShopPageResult) GoString() string {
	return s.String()
}

func (s *ShopPageResult) SetRequestId(v string) *ShopPageResult {
	s.RequestId = &v
	return s
}

func (s *ShopPageResult) SetShopList(v []*ShopPageDataResult) *ShopPageResult {
	s.ShopList = v
	return s
}

func (s *ShopPageResult) SetTotal(v int32) *ShopPageResult {
	s.Total = &v
	return s
}

type ShopStatusChangeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WORKING
	ShopStatus *string `json:"shopStatus,omitempty" xml:"shopStatus,omitempty"`
}

func (s ShopStatusChangeRequest) String() string {
	return tea.Prettify(s)
}

func (s ShopStatusChangeRequest) GoString() string {
	return s.String()
}

func (s *ShopStatusChangeRequest) SetShopId(v string) *ShopStatusChangeRequest {
	s.ShopId = &v
	return s
}

func (s *ShopStatusChangeRequest) SetShopStatus(v string) *ShopStatusChangeRequest {
	s.ShopStatus = &v
	return s
}

type ShopStatusChangeResult struct {
	// example:
	//
	// true
	Operate *bool `json:"operate,omitempty" xml:"operate,omitempty"`
}

func (s ShopStatusChangeResult) String() string {
	return tea.Prettify(s)
}

func (s ShopStatusChangeResult) GoString() string {
	return s.String()
}

func (s *ShopStatusChangeResult) SetOperate(v bool) *ShopStatusChangeResult {
	s.Operate = &v
	return s
}

type Sku struct {
	// example:
	//
	// 6922454329176
	Barcode *string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// example:
	//
	// true
	CanSell *bool `json:"canSell,omitempty" xml:"canSell,omitempty"`
	// example:
	//
	// 110000
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// example:
	//
	// 有货
	FuzzyQuantity *string `json:"fuzzyQuantity,omitempty" xml:"fuzzyQuantity,omitempty"`
	// example:
	//
	// 999900
	MarkPrice *int64 `json:"markPrice,omitempty" xml:"markPrice,omitempty"`
	// example:
	//
	// https:////img.alicdn.com/imgextra///img.alicdn.com/imgextra/i2/2216003305543/O1CN010DEQCX1qokFYGRfPE_!!2216003305543.png
	PicUrl *string `json:"picUrl,omitempty" xml:"picUrl,omitempty"`
	// example:
	//
	// 999900
	PlatformPrice *int64 `json:"platformPrice,omitempty" xml:"platformPrice,omitempty"`
	// example:
	//
	// 19800
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 660460842235822080
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// -1
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// example:
	//
	// 3
	RankValue *int64 `json:"rankValue,omitempty" xml:"rankValue,omitempty"`
	// example:
	//
	// 21000017
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// example:
	//
	// 660460842235822081
	SkuId    *string    `json:"skuId,omitempty" xml:"skuId,omitempty"`
	SkuSpecs []*SkuSpec `json:"skuSpecs,omitempty" xml:"skuSpecs,omitempty" type:"Repeated"`
	// example:
	//
	// 颜色分类:天蓝色
	SkuSpecsCode *string `json:"skuSpecsCode,omitempty" xml:"skuSpecsCode,omitempty"`
	// example:
	//
	// Online
	SkuStatus            *string `json:"skuStatus,omitempty" xml:"skuStatus,omitempty"`
	SuggestedRetailPrice *int64  `json:"suggestedRetailPrice,omitempty" xml:"suggestedRetailPrice,omitempty"`
	// example:
	//
	// 天蓝色
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s Sku) String() string {
	return tea.Prettify(s)
}

func (s Sku) GoString() string {
	return s.String()
}

func (s *Sku) SetBarcode(v string) *Sku {
	s.Barcode = &v
	return s
}

func (s *Sku) SetCanSell(v bool) *Sku {
	s.CanSell = &v
	return s
}

func (s *Sku) SetDivisionCode(v string) *Sku {
	s.DivisionCode = &v
	return s
}

func (s *Sku) SetFuzzyQuantity(v string) *Sku {
	s.FuzzyQuantity = &v
	return s
}

func (s *Sku) SetMarkPrice(v int64) *Sku {
	s.MarkPrice = &v
	return s
}

func (s *Sku) SetPicUrl(v string) *Sku {
	s.PicUrl = &v
	return s
}

func (s *Sku) SetPlatformPrice(v int64) *Sku {
	s.PlatformPrice = &v
	return s
}

func (s *Sku) SetPrice(v int64) *Sku {
	s.Price = &v
	return s
}

func (s *Sku) SetProductId(v string) *Sku {
	s.ProductId = &v
	return s
}

func (s *Sku) SetQuantity(v int64) *Sku {
	s.Quantity = &v
	return s
}

func (s *Sku) SetRankValue(v int64) *Sku {
	s.RankValue = &v
	return s
}

func (s *Sku) SetShopId(v string) *Sku {
	s.ShopId = &v
	return s
}

func (s *Sku) SetSkuId(v string) *Sku {
	s.SkuId = &v
	return s
}

func (s *Sku) SetSkuSpecs(v []*SkuSpec) *Sku {
	s.SkuSpecs = v
	return s
}

func (s *Sku) SetSkuSpecsCode(v string) *Sku {
	s.SkuSpecsCode = &v
	return s
}

func (s *Sku) SetSkuStatus(v string) *Sku {
	s.SkuStatus = &v
	return s
}

func (s *Sku) SetSuggestedRetailPrice(v int64) *Sku {
	s.SuggestedRetailPrice = &v
	return s
}

func (s *Sku) SetTitle(v string) *Sku {
	s.Title = &v
	return s
}

type SkuQueryParam struct {
	// This parameter is required.
	//
	// example:
	//
	// 660460842235822080
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 660460842235822081
	SkuId *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
}

func (s SkuQueryParam) String() string {
	return tea.Prettify(s)
}

func (s SkuQueryParam) GoString() string {
	return s.String()
}

func (s *SkuQueryParam) SetProductId(v string) *SkuQueryParam {
	s.ProductId = &v
	return s
}

func (s *SkuQueryParam) SetSkuId(v string) *SkuQueryParam {
	s.SkuId = &v
	return s
}

type SkuSaleInfo struct {
	// example:
	//
	// true
	CanSell *bool `json:"canSell,omitempty" xml:"canSell,omitempty"`
	// example:
	//
	// 110000
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// example:
	//
	// 有货
	FuzzyQuantity *string `json:"fuzzyQuantity,omitempty" xml:"fuzzyQuantity,omitempty"`
	// example:
	//
	// 999900
	MarkPrice *int64 `json:"markPrice,omitempty" xml:"markPrice,omitempty"`
	// example:
	//
	// 19800
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 660460842235822080
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// -1
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// example:
	//
	// 21000017
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// example:
	//
	// 660460842235822081
	SkuId *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
	// example:
	//
	// Online
	SkuStatus *string `json:"skuStatus,omitempty" xml:"skuStatus,omitempty"`
	// example:
	//
	// 天蓝色
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s SkuSaleInfo) String() string {
	return tea.Prettify(s)
}

func (s SkuSaleInfo) GoString() string {
	return s.String()
}

func (s *SkuSaleInfo) SetCanSell(v bool) *SkuSaleInfo {
	s.CanSell = &v
	return s
}

func (s *SkuSaleInfo) SetDivisionCode(v string) *SkuSaleInfo {
	s.DivisionCode = &v
	return s
}

func (s *SkuSaleInfo) SetFuzzyQuantity(v string) *SkuSaleInfo {
	s.FuzzyQuantity = &v
	return s
}

func (s *SkuSaleInfo) SetMarkPrice(v int64) *SkuSaleInfo {
	s.MarkPrice = &v
	return s
}

func (s *SkuSaleInfo) SetPrice(v int64) *SkuSaleInfo {
	s.Price = &v
	return s
}

func (s *SkuSaleInfo) SetProductId(v string) *SkuSaleInfo {
	s.ProductId = &v
	return s
}

func (s *SkuSaleInfo) SetQuantity(v int64) *SkuSaleInfo {
	s.Quantity = &v
	return s
}

func (s *SkuSaleInfo) SetShopId(v string) *SkuSaleInfo {
	s.ShopId = &v
	return s
}

func (s *SkuSaleInfo) SetSkuId(v string) *SkuSaleInfo {
	s.SkuId = &v
	return s
}

func (s *SkuSaleInfo) SetSkuStatus(v string) *SkuSaleInfo {
	s.SkuStatus = &v
	return s
}

func (s *SkuSaleInfo) SetTitle(v string) *SkuSaleInfo {
	s.Title = &v
	return s
}

type SkuSaleInfoListQuery struct {
	// example:
	//
	// 110000
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 21000017
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// This parameter is required.
	SkuQueryParams []*SkuQueryParam `json:"skuQueryParams,omitempty" xml:"skuQueryParams,omitempty" type:"Repeated"`
}

func (s SkuSaleInfoListQuery) String() string {
	return tea.Prettify(s)
}

func (s SkuSaleInfoListQuery) GoString() string {
	return s.String()
}

func (s *SkuSaleInfoListQuery) SetDivisionCode(v string) *SkuSaleInfoListQuery {
	s.DivisionCode = &v
	return s
}

func (s *SkuSaleInfoListQuery) SetPurchaserId(v string) *SkuSaleInfoListQuery {
	s.PurchaserId = &v
	return s
}

func (s *SkuSaleInfoListQuery) SetSkuQueryParams(v []*SkuQueryParam) *SkuSaleInfoListQuery {
	s.SkuQueryParams = v
	return s
}

type SkuSaleInfoListResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId    *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SkuSaleInfos []*SkuSaleInfo `json:"skuSaleInfos,omitempty" xml:"skuSaleInfos,omitempty" type:"Repeated"`
}

func (s SkuSaleInfoListResult) String() string {
	return tea.Prettify(s)
}

func (s SkuSaleInfoListResult) GoString() string {
	return s.String()
}

func (s *SkuSaleInfoListResult) SetRequestId(v string) *SkuSaleInfoListResult {
	s.RequestId = &v
	return s
}

func (s *SkuSaleInfoListResult) SetSkuSaleInfos(v []*SkuSaleInfo) *SkuSaleInfoListResult {
	s.SkuSaleInfos = v
	return s
}

type SkuSpec struct {
	// example:
	//
	// 颜色分类
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 1000
	KeyId *int64 `json:"keyId,omitempty" xml:"keyId,omitempty"`
	// example:
	//
	// 天蓝色
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// 秘色
	ValueAlias *string `json:"valueAlias,omitempty" xml:"valueAlias,omitempty"`
	// example:
	//
	// 1000
	ValueId *int64 `json:"valueId,omitempty" xml:"valueId,omitempty"`
}

func (s SkuSpec) String() string {
	return tea.Prettify(s)
}

func (s SkuSpec) GoString() string {
	return s.String()
}

func (s *SkuSpec) SetKey(v string) *SkuSpec {
	s.Key = &v
	return s
}

func (s *SkuSpec) SetKeyId(v int64) *SkuSpec {
	s.KeyId = &v
	return s
}

func (s *SkuSpec) SetValue(v string) *SkuSpec {
	s.Value = &v
	return s
}

func (s *SkuSpec) SetValueAlias(v string) *SkuSpec {
	s.ValueAlias = &v
	return s
}

func (s *SkuSpec) SetValueId(v int64) *SkuSpec {
	s.ValueId = &v
	return s
}

type CancelRefundOrderResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundOrderResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRefundOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelRefundOrderResponse) SetHeaders(v map[string]*string) *CancelRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelRefundOrderResponse) SetStatusCode(v int32) *CancelRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRefundOrderResponse) SetBody(v *RefundOrderResult) *CancelRefundOrderResponse {
	s.Body = v
	return s
}

type ConfirmDisburseRequest struct {
	// This parameter is required.
	Body *ConfirmDisburseCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmDisburseRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDisburseRequest) GoString() string {
	return s.String()
}

func (s *ConfirmDisburseRequest) SetBody(v *ConfirmDisburseCmd) *ConfirmDisburseRequest {
	s.Body = v
	return s
}

type ConfirmDisburseResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmDisburseResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmDisburseResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDisburseResponse) GoString() string {
	return s.String()
}

func (s *ConfirmDisburseResponse) SetHeaders(v map[string]*string) *ConfirmDisburseResponse {
	s.Headers = v
	return s
}

func (s *ConfirmDisburseResponse) SetStatusCode(v int32) *ConfirmDisburseResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmDisburseResponse) SetBody(v *ConfirmDisburseResult) *ConfirmDisburseResponse {
	s.Body = v
	return s
}

type CreateGoodsShippingNoticeRequest struct {
	// This parameter is required.
	Body *GoodsShippingNoticeCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGoodsShippingNoticeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGoodsShippingNoticeRequest) GoString() string {
	return s.String()
}

func (s *CreateGoodsShippingNoticeRequest) SetBody(v *GoodsShippingNoticeCreateCmd) *CreateGoodsShippingNoticeRequest {
	s.Body = v
	return s
}

type CreateGoodsShippingNoticeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GoodsShippingNoticeCreateResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGoodsShippingNoticeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGoodsShippingNoticeResponse) GoString() string {
	return s.String()
}

func (s *CreateGoodsShippingNoticeResponse) SetHeaders(v map[string]*string) *CreateGoodsShippingNoticeResponse {
	s.Headers = v
	return s
}

func (s *CreateGoodsShippingNoticeResponse) SetStatusCode(v int32) *CreateGoodsShippingNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGoodsShippingNoticeResponse) SetBody(v *GoodsShippingNoticeCreateResult) *CreateGoodsShippingNoticeResponse {
	s.Body = v
	return s
}

type CreatePurchaseOrderRequest struct {
	// This parameter is required.
	Body *PurchaseOrderCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePurchaseOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePurchaseOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePurchaseOrderRequest) SetBody(v *PurchaseOrderCreateCmd) *CreatePurchaseOrderRequest {
	s.Body = v
	return s
}

type CreatePurchaseOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseOrderCreateResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePurchaseOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePurchaseOrderResponse) GoString() string {
	return s.String()
}

func (s *CreatePurchaseOrderResponse) SetHeaders(v map[string]*string) *CreatePurchaseOrderResponse {
	s.Headers = v
	return s
}

func (s *CreatePurchaseOrderResponse) SetStatusCode(v int32) *CreatePurchaseOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePurchaseOrderResponse) SetBody(v *PurchaseOrderCreateResult) *CreatePurchaseOrderResponse {
	s.Body = v
	return s
}

type CreateRefundOrderRequest struct {
	// This parameter is required.
	Body *RefundOrderCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRefundOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRefundOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateRefundOrderRequest) SetBody(v *RefundOrderCmd) *CreateRefundOrderRequest {
	s.Body = v
	return s
}

type CreateRefundOrderResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundOrderResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRefundOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateRefundOrderResponse) SetHeaders(v map[string]*string) *CreateRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateRefundOrderResponse) SetStatusCode(v int32) *CreateRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRefundOrderResponse) SetBody(v *RefundOrderResult) *CreateRefundOrderResponse {
	s.Body = v
	return s
}

type GetOrderResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OrderResult       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrderResponse) GoString() string {
	return s.String()
}

func (s *GetOrderResponse) SetHeaders(v map[string]*string) *GetOrderResponse {
	s.Headers = v
	return s
}

func (s *GetOrderResponse) SetStatusCode(v int32) *GetOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrderResponse) SetBody(v *OrderResult) *GetOrderResponse {
	s.Body = v
	return s
}

type GetPurchaseOrderStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseOrderStatusResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPurchaseOrderStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPurchaseOrderStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPurchaseOrderStatusResponse) SetHeaders(v map[string]*string) *GetPurchaseOrderStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPurchaseOrderStatusResponse) SetStatusCode(v int32) *GetPurchaseOrderStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPurchaseOrderStatusResponse) SetBody(v *PurchaseOrderStatusResult) *GetPurchaseOrderStatusResponse {
	s.Body = v
	return s
}

type GetPurchaserShopResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Shop              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPurchaserShopResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPurchaserShopResponse) GoString() string {
	return s.String()
}

func (s *GetPurchaserShopResponse) SetHeaders(v map[string]*string) *GetPurchaserShopResponse {
	s.Headers = v
	return s
}

func (s *GetPurchaserShopResponse) SetStatusCode(v int32) *GetPurchaserShopResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPurchaserShopResponse) SetBody(v *Shop) *GetPurchaserShopResponse {
	s.Body = v
	return s
}

type GetRefundOrderResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundResult      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRefundOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *GetRefundOrderResponse) SetHeaders(v map[string]*string) *GetRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *GetRefundOrderResponse) SetStatusCode(v int32) *GetRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRefundOrderResponse) SetBody(v *RefundResult) *GetRefundOrderResponse {
	s.Body = v
	return s
}

type GetSelectionProductRequest struct {
	// example:
	//
	// 110000
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
}

func (s GetSelectionProductRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSelectionProductRequest) GoString() string {
	return s.String()
}

func (s *GetSelectionProductRequest) SetDivisionCode(v string) *GetSelectionProductRequest {
	s.DivisionCode = &v
	return s
}

func (s *GetSelectionProductRequest) SetPurchaserId(v string) *GetSelectionProductRequest {
	s.PurchaserId = &v
	return s
}

type GetSelectionProductResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Product           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSelectionProductResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSelectionProductResponse) GoString() string {
	return s.String()
}

func (s *GetSelectionProductResponse) SetHeaders(v map[string]*string) *GetSelectionProductResponse {
	s.Headers = v
	return s
}

func (s *GetSelectionProductResponse) SetStatusCode(v int32) *GetSelectionProductResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSelectionProductResponse) SetBody(v *Product) *GetSelectionProductResponse {
	s.Body = v
	return s
}

type GetSelectionProductSaleInfoRequest struct {
	// example:
	//
	// 110000
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
}

func (s GetSelectionProductSaleInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSelectionProductSaleInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSelectionProductSaleInfoRequest) SetDivisionCode(v string) *GetSelectionProductSaleInfoRequest {
	s.DivisionCode = &v
	return s
}

func (s *GetSelectionProductSaleInfoRequest) SetPurchaserId(v string) *GetSelectionProductSaleInfoRequest {
	s.PurchaserId = &v
	return s
}

type GetSelectionProductSaleInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProductSaleInfo   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSelectionProductSaleInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSelectionProductSaleInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSelectionProductSaleInfoResponse) SetHeaders(v map[string]*string) *GetSelectionProductSaleInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSelectionProductSaleInfoResponse) SetStatusCode(v int32) *GetSelectionProductSaleInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSelectionProductSaleInfoResponse) SetBody(v *ProductSaleInfo) *GetSelectionProductSaleInfoResponse {
	s.Body = v
	return s
}

type ListCategoriesRequest struct {
	Body *CategoryListQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCategoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCategoriesRequest) GoString() string {
	return s.String()
}

func (s *ListCategoriesRequest) SetBody(v *CategoryListQuery) *ListCategoriesRequest {
	s.Body = v
	return s
}

type ListCategoriesResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CategoryListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListCategoriesResponse) SetHeaders(v map[string]*string) *ListCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListCategoriesResponse) SetStatusCode(v int32) *ListCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCategoriesResponse) SetBody(v *CategoryListResult) *ListCategoriesResponse {
	s.Body = v
	return s
}

type ListLogisticsOrdersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogisticsOrderListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogisticsOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogisticsOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListLogisticsOrdersResponse) SetHeaders(v map[string]*string) *ListLogisticsOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListLogisticsOrdersResponse) SetStatusCode(v int32) *ListLogisticsOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogisticsOrdersResponse) SetBody(v *LogisticsOrderListResult) *ListLogisticsOrdersResponse {
	s.Body = v
	return s
}

type ListPurchaserShopsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPurchaserShopsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPurchaserShopsRequest) GoString() string {
	return s.String()
}

func (s *ListPurchaserShopsRequest) SetPageNumber(v int32) *ListPurchaserShopsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPurchaserShopsRequest) SetPageSize(v int32) *ListPurchaserShopsRequest {
	s.PageSize = &v
	return s
}

type ListPurchaserShopsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShopPageResult    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPurchaserShopsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPurchaserShopsResponse) GoString() string {
	return s.String()
}

func (s *ListPurchaserShopsResponse) SetHeaders(v map[string]*string) *ListPurchaserShopsResponse {
	s.Headers = v
	return s
}

func (s *ListPurchaserShopsResponse) SetStatusCode(v int32) *ListPurchaserShopsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPurchaserShopsResponse) SetBody(v *ShopPageResult) *ListPurchaserShopsResponse {
	s.Body = v
	return s
}

type ListSelectionProductSaleInfosRequest struct {
	// This parameter is required.
	Body *ProductSaleInfoListQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSelectionProductSaleInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSelectionProductSaleInfosRequest) GoString() string {
	return s.String()
}

func (s *ListSelectionProductSaleInfosRequest) SetBody(v *ProductSaleInfoListQuery) *ListSelectionProductSaleInfosRequest {
	s.Body = v
	return s
}

type ListSelectionProductSaleInfosResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProductSaleInfoListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSelectionProductSaleInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSelectionProductSaleInfosResponse) GoString() string {
	return s.String()
}

func (s *ListSelectionProductSaleInfosResponse) SetHeaders(v map[string]*string) *ListSelectionProductSaleInfosResponse {
	s.Headers = v
	return s
}

func (s *ListSelectionProductSaleInfosResponse) SetStatusCode(v int32) *ListSelectionProductSaleInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSelectionProductSaleInfosResponse) SetBody(v *ProductSaleInfoListResult) *ListSelectionProductSaleInfosResponse {
	s.Body = v
	return s
}

type ListSelectionProductsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
}

func (s ListSelectionProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSelectionProductsRequest) GoString() string {
	return s.String()
}

func (s *ListSelectionProductsRequest) SetPageNumber(v int32) *ListSelectionProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSelectionProductsRequest) SetPageSize(v int32) *ListSelectionProductsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSelectionProductsRequest) SetPurchaserId(v string) *ListSelectionProductsRequest {
	s.PurchaserId = &v
	return s
}

type ListSelectionProductsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProductPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSelectionProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSelectionProductsResponse) GoString() string {
	return s.String()
}

func (s *ListSelectionProductsResponse) SetHeaders(v map[string]*string) *ListSelectionProductsResponse {
	s.Headers = v
	return s
}

func (s *ListSelectionProductsResponse) SetStatusCode(v int32) *ListSelectionProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSelectionProductsResponse) SetBody(v *ProductPageResult) *ListSelectionProductsResponse {
	s.Body = v
	return s
}

type ListSelectionSkuSaleInfosRequest struct {
	// This parameter is required.
	Body *SkuSaleInfoListQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSelectionSkuSaleInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSelectionSkuSaleInfosRequest) GoString() string {
	return s.String()
}

func (s *ListSelectionSkuSaleInfosRequest) SetBody(v *SkuSaleInfoListQuery) *ListSelectionSkuSaleInfosRequest {
	s.Body = v
	return s
}

type ListSelectionSkuSaleInfosResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SkuSaleInfoListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSelectionSkuSaleInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSelectionSkuSaleInfosResponse) GoString() string {
	return s.String()
}

func (s *ListSelectionSkuSaleInfosResponse) SetHeaders(v map[string]*string) *ListSelectionSkuSaleInfosResponse {
	s.Headers = v
	return s
}

func (s *ListSelectionSkuSaleInfosResponse) SetStatusCode(v int32) *ListSelectionSkuSaleInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSelectionSkuSaleInfosResponse) SetBody(v *SkuSaleInfoListResult) *ListSelectionSkuSaleInfosResponse {
	s.Body = v
	return s
}

type QueryChildDivisionCodeRequest struct {
	// This parameter is required.
	Body *DivisionQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryChildDivisionCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryChildDivisionCodeRequest) GoString() string {
	return s.String()
}

func (s *QueryChildDivisionCodeRequest) SetBody(v *DivisionQuery) *QueryChildDivisionCodeRequest {
	s.Body = v
	return s
}

type QueryChildDivisionCodeResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DivisionPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryChildDivisionCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryChildDivisionCodeResponse) GoString() string {
	return s.String()
}

func (s *QueryChildDivisionCodeResponse) SetHeaders(v map[string]*string) *QueryChildDivisionCodeResponse {
	s.Headers = v
	return s
}

func (s *QueryChildDivisionCodeResponse) SetStatusCode(v int32) *QueryChildDivisionCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChildDivisionCodeResponse) SetBody(v *DivisionPageResult) *QueryChildDivisionCodeResponse {
	s.Body = v
	return s
}

type QueryOrdersRequest struct {
	// This parameter is required.
	Body *OrderPageQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrdersRequest) GoString() string {
	return s.String()
}

func (s *QueryOrdersRequest) SetBody(v *OrderPageQuery) *QueryOrdersRequest {
	s.Body = v
	return s
}

type QueryOrdersResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OrderListResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrdersResponse) GoString() string {
	return s.String()
}

func (s *QueryOrdersResponse) SetHeaders(v map[string]*string) *QueryOrdersResponse {
	s.Headers = v
	return s
}

func (s *QueryOrdersResponse) SetStatusCode(v int32) *QueryOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrdersResponse) SetBody(v *OrderListResult) *QueryOrdersResponse {
	s.Body = v
	return s
}

type RenderPurchaseOrderRequest struct {
	// This parameter is required.
	Body *PurchaseOrderRenderQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenderPurchaseOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RenderPurchaseOrderRequest) GoString() string {
	return s.String()
}

func (s *RenderPurchaseOrderRequest) SetBody(v *PurchaseOrderRenderQuery) *RenderPurchaseOrderRequest {
	s.Body = v
	return s
}

type RenderPurchaseOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseOrderRenderResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenderPurchaseOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RenderPurchaseOrderResponse) GoString() string {
	return s.String()
}

func (s *RenderPurchaseOrderResponse) SetHeaders(v map[string]*string) *RenderPurchaseOrderResponse {
	s.Headers = v
	return s
}

func (s *RenderPurchaseOrderResponse) SetStatusCode(v int32) *RenderPurchaseOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *RenderPurchaseOrderResponse) SetBody(v *PurchaseOrderRenderResult) *RenderPurchaseOrderResponse {
	s.Body = v
	return s
}

type RenderRefundOrderRequest struct {
	// This parameter is required.
	Body *RefundRenderCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenderRefundOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RenderRefundOrderRequest) GoString() string {
	return s.String()
}

func (s *RenderRefundOrderRequest) SetBody(v *RefundRenderCmd) *RenderRefundOrderRequest {
	s.Body = v
	return s
}

type RenderRefundOrderResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundRenderResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenderRefundOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RenderRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *RenderRefundOrderResponse) SetHeaders(v map[string]*string) *RenderRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *RenderRefundOrderResponse) SetStatusCode(v int32) *RenderRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *RenderRefundOrderResponse) SetBody(v *RefundRenderResult) *RenderRefundOrderResponse {
	s.Body = v
	return s
}

type SplitPurchaseOrderRequest struct {
	Body *PurchaseOrderRenderQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SplitPurchaseOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SplitPurchaseOrderRequest) GoString() string {
	return s.String()
}

func (s *SplitPurchaseOrderRequest) SetBody(v *PurchaseOrderRenderQuery) *SplitPurchaseOrderRequest {
	s.Body = v
	return s
}

type SplitPurchaseOrderResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseOrderRenderResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SplitPurchaseOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SplitPurchaseOrderResponse) GoString() string {
	return s.String()
}

func (s *SplitPurchaseOrderResponse) SetHeaders(v map[string]*string) *SplitPurchaseOrderResponse {
	s.Headers = v
	return s
}

func (s *SplitPurchaseOrderResponse) SetStatusCode(v int32) *SplitPurchaseOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *SplitPurchaseOrderResponse) SetBody(v *PurchaseOrderRenderResult) *SplitPurchaseOrderResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":                 tea.String("linkedmall.aliyuncs.com"),
		"cn-shanghai":                 tea.String("linkedmall.aliyuncs.com"),
		"ap-northeast-1":              tea.String("linkedmall.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("linkedmall.aliyuncs.com"),
		"ap-south-1":                  tea.String("linkedmall.aliyuncs.com"),
		"ap-southeast-1":              tea.String("linkedmall.aliyuncs.com"),
		"ap-southeast-2":              tea.String("linkedmall.aliyuncs.com"),
		"ap-southeast-3":              tea.String("linkedmall.aliyuncs.com"),
		"ap-southeast-5":              tea.String("linkedmall.aliyuncs.com"),
		"cn-beijing":                  tea.String("linkedmall.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("linkedmall.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("linkedmall.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("linkedmall.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("linkedmall.aliyuncs.com"),
		"cn-chengdu":                  tea.String("linkedmall.aliyuncs.com"),
		"cn-edge-1":                   tea.String("linkedmall.aliyuncs.com"),
		"cn-fujian":                   tea.String("linkedmall.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("linkedmall.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("linkedmall.aliyuncs.com"),
		"cn-hongkong":                 tea.String("linkedmall.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("linkedmall.aliyuncs.com"),
		"cn-huhehaote":                tea.String("linkedmall.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("linkedmall.aliyuncs.com"),
		"cn-qingdao":                  tea.String("linkedmall.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("linkedmall.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("linkedmall.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("linkedmall.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("linkedmall.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("linkedmall.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("linkedmall.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("linkedmall.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("linkedmall.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("linkedmall.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("linkedmall.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("linkedmall.aliyuncs.com"),
		"cn-wuhan":                    tea.String("linkedmall.aliyuncs.com"),
		"cn-yushanfang":               tea.String("linkedmall.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("linkedmall.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("linkedmall.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("linkedmall.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("linkedmall.aliyuncs.com"),
		"eu-central-1":                tea.String("linkedmall.aliyuncs.com"),
		"eu-west-1":                   tea.String("linkedmall.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("linkedmall.aliyuncs.com"),
		"me-east-1":                   tea.String("linkedmall.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("linkedmall.aliyuncs.com"),
		"us-east-1":                   tea.String("linkedmall.aliyuncs.com"),
		"us-west-1":                   tea.String("linkedmall.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("linkedmall"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消逆向单
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRefundOrderResponse
func (client *Client) CancelRefundOrderWithOptions(disputeId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelRefundOrderResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CancelRefundOrder"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/refunds/" + tea.StringValue(openapiutil.GetEncodeParam(disputeId)) + "/commands/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelRefundOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消逆向单
//
// @return CancelRefundOrderResponse
func (client *Client) CancelRefundOrder(disputeId *string) (_result *CancelRefundOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelRefundOrderResponse{}
	_body, _err := client.CancelRefundOrderWithOptions(disputeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 确认收货（订单）
//
// @param request - ConfirmDisburseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfirmDisburseResponse
func (client *Client) ConfirmDisburseWithOptions(request *ConfirmDisburseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ConfirmDisburseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfirmDisburse"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/orders/commands/confirmDisburse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmDisburseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 确认收货（订单）
//
// @param request - ConfirmDisburseRequest
//
// @return ConfirmDisburseResponse
func (client *Client) ConfirmDisburse(request *ConfirmDisburseRequest) (_result *ConfirmDisburseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConfirmDisburseResponse{}
	_body, _err := client.ConfirmDisburseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交运单信息
//
// @param request - CreateGoodsShippingNoticeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGoodsShippingNoticeResponse
func (client *Client) CreateGoodsShippingNoticeWithOptions(request *CreateGoodsShippingNoticeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGoodsShippingNoticeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGoodsShippingNotice"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/refunds/command/createGoodsShippingNotice"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGoodsShippingNoticeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交运单信息
//
// @param request - CreateGoodsShippingNoticeRequest
//
// @return CreateGoodsShippingNoticeResponse
func (client *Client) CreateGoodsShippingNotice(request *CreateGoodsShippingNoticeRequest) (_result *CreateGoodsShippingNoticeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGoodsShippingNoticeResponse{}
	_body, _err := client.CreateGoodsShippingNoticeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建采购单
//
// @param request - CreatePurchaseOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePurchaseOrderResponse
func (client *Client) CreatePurchaseOrderWithOptions(request *CreatePurchaseOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePurchaseOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePurchaseOrder"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/purchaseOrders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePurchaseOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建采购单
//
// @param request - CreatePurchaseOrderRequest
//
// @return CreatePurchaseOrderResponse
func (client *Client) CreatePurchaseOrder(request *CreatePurchaseOrderRequest) (_result *CreatePurchaseOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePurchaseOrderResponse{}
	_body, _err := client.CreatePurchaseOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建逆向单
//
// @param request - CreateRefundOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRefundOrderResponse
func (client *Client) CreateRefundOrderWithOptions(request *CreateRefundOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRefundOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRefundOrder"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/refunds"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRefundOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建逆向单
//
// @param request - CreateRefundOrderRequest
//
// @return CreateRefundOrderResponse
func (client *Client) CreateRefundOrder(request *CreateRefundOrderRequest) (_result *CreateRefundOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRefundOrderResponse{}
	_body, _err := client.CreateRefundOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询主单详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrderResponse
func (client *Client) GetOrderWithOptions(orderId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOrderResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrder"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/orders/" + tea.StringValue(openapiutil.GetEncodeParam(orderId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询主单详情
//
// @return GetOrderResponse
func (client *Client) GetOrder(orderId *string) (_result *GetOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOrderResponse{}
	_body, _err := client.GetOrderWithOptions(orderId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询采购单状态
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPurchaseOrderStatusResponse
func (client *Client) GetPurchaseOrderStatusWithOptions(purchaseOrderId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPurchaseOrderStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPurchaseOrderStatus"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/purchaseOrders/" + tea.StringValue(openapiutil.GetEncodeParam(purchaseOrderId)) + "/status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPurchaseOrderStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询采购单状态
//
// @return GetPurchaseOrderStatusResponse
func (client *Client) GetPurchaseOrderStatus(purchaseOrderId *string) (_result *GetPurchaseOrderStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPurchaseOrderStatusResponse{}
	_body, _err := client.GetPurchaseOrderStatusWithOptions(purchaseOrderId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询分销商店铺
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPurchaserShopResponse
func (client *Client) GetPurchaserShopWithOptions(purchaserId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPurchaserShopResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPurchaserShop"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/purchaserShops/" + tea.StringValue(openapiutil.GetEncodeParam(purchaserId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPurchaserShopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询分销商店铺
//
// @return GetPurchaserShopResponse
func (client *Client) GetPurchaserShop(purchaserId *string) (_result *GetPurchaserShopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPurchaserShopResponse{}
	_body, _err := client.GetPurchaserShopWithOptions(purchaserId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询逆向单详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRefundOrderResponse
func (client *Client) GetRefundOrderWithOptions(disputeId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRefundOrderResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRefundOrder"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/refunds/" + tea.StringValue(openapiutil.GetEncodeParam(disputeId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRefundOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询逆向单详情
//
// @return GetRefundOrderResponse
func (client *Client) GetRefundOrder(disputeId *string) (_result *GetRefundOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRefundOrderResponse{}
	_body, _err := client.GetRefundOrderWithOptions(disputeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询选品池商品详情
//
// @param request - GetSelectionProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSelectionProductResponse
func (client *Client) GetSelectionProductWithOptions(productId *string, request *GetSelectionProductRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSelectionProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DivisionCode)) {
		query["divisionCode"] = request.DivisionCode
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserId)) {
		query["purchaserId"] = request.PurchaserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSelectionProduct"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/products/" + tea.StringValue(openapiutil.GetEncodeParam(productId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSelectionProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询选品池商品详情
//
// @param request - GetSelectionProductRequest
//
// @return GetSelectionProductResponse
func (client *Client) GetSelectionProduct(productId *string, request *GetSelectionProductRequest) (_result *GetSelectionProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSelectionProductResponse{}
	_body, _err := client.GetSelectionProductWithOptions(productId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询选品池商品库存
//
// @param request - GetSelectionProductSaleInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSelectionProductSaleInfoResponse
func (client *Client) GetSelectionProductSaleInfoWithOptions(productId *string, request *GetSelectionProductSaleInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSelectionProductSaleInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DivisionCode)) {
		query["divisionCode"] = request.DivisionCode
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserId)) {
		query["purchaserId"] = request.PurchaserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSelectionProductSaleInfo"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/products/" + tea.StringValue(openapiutil.GetEncodeParam(productId)) + "/saleInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSelectionProductSaleInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询选品池商品库存
//
// @param request - GetSelectionProductSaleInfoRequest
//
// @return GetSelectionProductSaleInfoResponse
func (client *Client) GetSelectionProductSaleInfo(productId *string, request *GetSelectionProductSaleInfoRequest) (_result *GetSelectionProductSaleInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSelectionProductSaleInfoResponse{}
	_body, _err := client.GetSelectionProductSaleInfoWithOptions(productId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询类目
//
// @param request - ListCategoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCategoriesResponse
func (client *Client) ListCategoriesWithOptions(request *ListCategoriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCategoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCategories"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/categories/commands/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询类目
//
// @param request - ListCategoriesRequest
//
// @return ListCategoriesResponse
func (client *Client) ListCategories(request *ListCategoriesRequest) (_result *ListCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCategoriesResponse{}
	_body, _err := client.ListCategoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询物流信息（订单）
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogisticsOrdersResponse
func (client *Client) ListLogisticsOrdersWithOptions(orderId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLogisticsOrdersResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListLogisticsOrders"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/orders/" + tea.StringValue(openapiutil.GetEncodeParam(orderId)) + "/logisticsOrders"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLogisticsOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询物流信息（订单）
//
// @return ListLogisticsOrdersResponse
func (client *Client) ListLogisticsOrders(orderId *string) (_result *ListLogisticsOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogisticsOrdersResponse{}
	_body, _err := client.ListLogisticsOrdersWithOptions(orderId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 采购方店铺列表查询
//
// @param request - ListPurchaserShopsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPurchaserShopsResponse
func (client *Client) ListPurchaserShopsWithOptions(request *ListPurchaserShopsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPurchaserShopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPurchaserShops"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/purchaserShops"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPurchaserShopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 采购方店铺列表查询
//
// @param request - ListPurchaserShopsRequest
//
// @return ListPurchaserShopsResponse
func (client *Client) ListPurchaserShops(request *ListPurchaserShopsRequest) (_result *ListPurchaserShopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPurchaserShopsResponse{}
	_body, _err := client.ListPurchaserShopsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询选品池商品库存
//
// @param request - ListSelectionProductSaleInfosRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSelectionProductSaleInfosResponse
func (client *Client) ListSelectionProductSaleInfosWithOptions(request *ListSelectionProductSaleInfosRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSelectionProductSaleInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSelectionProductSaleInfos"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/products/saleInfo/commands/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSelectionProductSaleInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询选品池商品库存
//
// @param request - ListSelectionProductSaleInfosRequest
//
// @return ListSelectionProductSaleInfosResponse
func (client *Client) ListSelectionProductSaleInfos(request *ListSelectionProductSaleInfosRequest) (_result *ListSelectionProductSaleInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSelectionProductSaleInfosResponse{}
	_body, _err := client.ListSelectionProductSaleInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询商品列表
//
// @param request - ListSelectionProductsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSelectionProductsResponse
func (client *Client) ListSelectionProductsWithOptions(request *ListSelectionProductsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSelectionProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaserId)) {
		query["purchaserId"] = request.PurchaserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSelectionProducts"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/products"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSelectionProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询商品列表
//
// @param request - ListSelectionProductsRequest
//
// @return ListSelectionProductsResponse
func (client *Client) ListSelectionProducts(request *ListSelectionProductsRequest) (_result *ListSelectionProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSelectionProductsResponse{}
	_body, _err := client.ListSelectionProductsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询选品池商品SKU库存
//
// @param request - ListSelectionSkuSaleInfosRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSelectionSkuSaleInfosResponse
func (client *Client) ListSelectionSkuSaleInfosWithOptions(request *ListSelectionSkuSaleInfosRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSelectionSkuSaleInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSelectionSkuSaleInfos"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/selectionPool/skus/saleInfo/commands/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSelectionSkuSaleInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询选品池商品SKU库存
//
// @param request - ListSelectionSkuSaleInfosRequest
//
// @return ListSelectionSkuSaleInfosResponse
func (client *Client) ListSelectionSkuSaleInfos(request *ListSelectionSkuSaleInfosRequest) (_result *ListSelectionSkuSaleInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSelectionSkuSaleInfosResponse{}
	_body, _err := client.ListSelectionSkuSaleInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询地址divisionCode
//
// @param request - QueryChildDivisionCodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChildDivisionCodeResponse
func (client *Client) QueryChildDivisionCodeWithOptions(request *QueryChildDivisionCodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryChildDivisionCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryChildDivisionCode"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/division/commands/queryChildDivisionCode"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryChildDivisionCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询地址divisionCode
//
// @param request - QueryChildDivisionCodeRequest
//
// @return QueryChildDivisionCodeResponse
func (client *Client) QueryChildDivisionCode(request *QueryChildDivisionCodeRequest) (_result *QueryChildDivisionCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryChildDivisionCodeResponse{}
	_body, _err := client.QueryChildDivisionCodeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询主单列表
//
// @param request - QueryOrdersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrdersResponse
func (client *Client) QueryOrdersWithOptions(request *QueryOrdersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOrders"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/orders/commands/query"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询主单列表
//
// @param request - QueryOrdersRequest
//
// @return QueryOrdersResponse
func (client *Client) QueryOrders(request *QueryOrdersRequest) (_result *QueryOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryOrdersResponse{}
	_body, _err := client.QueryOrdersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 渲染采购单
//
// @param request - RenderPurchaseOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenderPurchaseOrderResponse
func (client *Client) RenderPurchaseOrderWithOptions(request *RenderPurchaseOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RenderPurchaseOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenderPurchaseOrder"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/purchaseOrders/commands/render"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RenderPurchaseOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 渲染采购单
//
// @param request - RenderPurchaseOrderRequest
//
// @return RenderPurchaseOrderResponse
func (client *Client) RenderPurchaseOrder(request *RenderPurchaseOrderRequest) (_result *RenderPurchaseOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenderPurchaseOrderResponse{}
	_body, _err := client.RenderPurchaseOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 逆向单渲染
//
// @param request - RenderRefundOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenderRefundOrderResponse
func (client *Client) RenderRefundOrderWithOptions(request *RenderRefundOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RenderRefundOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenderRefundOrder"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/refunds/commands/render"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RenderRefundOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 逆向单渲染
//
// @param request - RenderRefundOrderRequest
//
// @return RenderRefundOrderResponse
func (client *Client) RenderRefundOrder(request *RenderRefundOrderRequest) (_result *RenderRefundOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenderRefundOrderResponse{}
	_body, _err := client.RenderRefundOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 渲染拆分采购单
//
// @param request - SplitPurchaseOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SplitPurchaseOrderResponse
func (client *Client) SplitPurchaseOrderWithOptions(request *SplitPurchaseOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SplitPurchaseOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("SplitPurchaseOrder"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/opensaas-s2b/opensaas-s2b-biz-trade/v2/purchaseOrders/commands/split"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SplitPurchaseOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 渲染拆分采购单
//
// @param request - SplitPurchaseOrderRequest
//
// @return SplitPurchaseOrderResponse
func (client *Client) SplitPurchaseOrder(request *SplitPurchaseOrderRequest) (_result *SplitPurchaseOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SplitPurchaseOrderResponse{}
	_body, _err := client.SplitPurchaseOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
