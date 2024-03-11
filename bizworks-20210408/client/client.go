// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryUsageResponseBody struct {
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBody) SetCode(v string) *QueryUsageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUsageResponseBody) SetData(v *QueryUsageResponseBodyData) *QueryUsageResponseBody {
	s.Data = v
	return s
}

func (s *QueryUsageResponseBody) SetMessage(v string) *QueryUsageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUsageResponseBody) SetRequestId(v string) *QueryUsageResponseBody {
	s.RequestId = &v
	return s
}

type QueryUsageResponseBodyData struct {
	ArticleInfos            []*QueryUsageResponseBodyDataArticleInfos            `json:"articleInfos,omitempty" xml:"articleInfos,omitempty" type:"Repeated"`
	CommoditySpecUsageInfos []*QueryUsageResponseBodyDataCommoditySpecUsageInfos `json:"commoditySpecUsageInfos,omitempty" xml:"commoditySpecUsageInfos,omitempty" type:"Repeated"`
	SubProductInfos         []*QueryUsageResponseBodyDataSubProductInfos         `json:"subProductInfos,omitempty" xml:"subProductInfos,omitempty" type:"Repeated"`
}

func (s QueryUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBodyData) SetArticleInfos(v []*QueryUsageResponseBodyDataArticleInfos) *QueryUsageResponseBodyData {
	s.ArticleInfos = v
	return s
}

func (s *QueryUsageResponseBodyData) SetCommoditySpecUsageInfos(v []*QueryUsageResponseBodyDataCommoditySpecUsageInfos) *QueryUsageResponseBodyData {
	s.CommoditySpecUsageInfos = v
	return s
}

func (s *QueryUsageResponseBodyData) SetSubProductInfos(v []*QueryUsageResponseBodyDataSubProductInfos) *QueryUsageResponseBodyData {
	s.SubProductInfos = v
	return s
}

type QueryUsageResponseBodyDataArticleInfos struct {
	ArticleType *string `json:"articleType,omitempty" xml:"articleType,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	Url         *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryUsageResponseBodyDataArticleInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBodyDataArticleInfos) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBodyDataArticleInfos) SetArticleType(v string) *QueryUsageResponseBodyDataArticleInfos {
	s.ArticleType = &v
	return s
}

func (s *QueryUsageResponseBodyDataArticleInfos) SetTitle(v string) *QueryUsageResponseBodyDataArticleInfos {
	s.Title = &v
	return s
}

func (s *QueryUsageResponseBodyDataArticleInfos) SetUrl(v string) *QueryUsageResponseBodyDataArticleInfos {
	s.Url = &v
	return s
}

type QueryUsageResponseBodyDataCommoditySpecUsageInfos struct {
	CommodityConfig              *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig                `json:"commodityConfig,omitempty" xml:"commodityConfig,omitempty" type:"Struct"`
	PurchasedCommoditySpecUsages []*QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsages `json:"purchasedCommoditySpecUsages,omitempty" xml:"purchasedCommoditySpecUsages,omitempty" type:"Repeated"`
}

func (s QueryUsageResponseBodyDataCommoditySpecUsageInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBodyDataCommoditySpecUsageInfos) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfos) SetCommodityConfig(v *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) *QueryUsageResponseBodyDataCommoditySpecUsageInfos {
	s.CommodityConfig = v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfos) SetPurchasedCommoditySpecUsages(v []*QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsages) *QueryUsageResponseBodyDataCommoditySpecUsageInfos {
	s.PurchasedCommoditySpecUsages = v
	return s
}

type QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig struct {
	BetaDeadlineMs        *int64    `json:"betaDeadlineMs,omitempty" xml:"betaDeadlineMs,omitempty"`
	CommodityCode         *string   `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	CommodityName         *string   `json:"commodityName,omitempty" xml:"commodityName,omitempty"`
	ExpiredToReleasedHour *int64    `json:"expiredToReleasedHour,omitempty" xml:"expiredToReleasedHour,omitempty"`
	InBeta                *bool     `json:"inBeta,omitempty" xml:"inBeta,omitempty"`
	Independent           *bool     `json:"independent,omitempty" xml:"independent,omitempty"`
	PaidType              *string   `json:"paidType,omitempty" xml:"paidType,omitempty"`
	ProductCode           *string   `json:"productCode,omitempty" xml:"productCode,omitempty"`
	QuotaItems            []*string `json:"quotaItems,omitempty" xml:"quotaItems,omitempty" type:"Repeated"`
	SpecCodes             []*string `json:"specCodes,omitempty" xml:"specCodes,omitempty" type:"Repeated"`
}

func (s QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) SetBetaDeadlineMs(v int64) *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig {
	s.BetaDeadlineMs = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) SetCommodityCode(v string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig {
	s.CommodityCode = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) SetCommodityName(v string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig {
	s.CommodityName = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) SetExpiredToReleasedHour(v int64) *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig {
	s.ExpiredToReleasedHour = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) SetInBeta(v bool) *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig {
	s.InBeta = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) SetIndependent(v bool) *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig {
	s.Independent = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) SetPaidType(v string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig {
	s.PaidType = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) SetProductCode(v string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig {
	s.ProductCode = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) SetQuotaItems(v []*string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig {
	s.QuotaItems = v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig) SetSpecCodes(v []*string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosCommodityConfig {
	s.SpecCodes = v
	return s
}

type QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsages struct {
	CommoditySpec  *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec  `json:"commoditySpec,omitempty" xml:"commoditySpec,omitempty" type:"Struct"`
	CommodityUsage *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommodityUsage `json:"commodityUsage,omitempty" xml:"commodityUsage,omitempty" type:"Struct"`
}

func (s QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsages) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsages) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsages) SetCommoditySpec(v *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsages {
	s.CommoditySpec = v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsages) SetCommodityUsage(v *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommodityUsage) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsages {
	s.CommodityUsage = v
	return s
}

type QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec struct {
	ApplicationNum         *string   `json:"applicationNum,omitempty" xml:"applicationNum,omitempty"`
	CommercializeStatus    *string   `json:"commercializeStatus,omitempty" xml:"commercializeStatus,omitempty"`
	CommodityCode          *string   `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	CommodityName          *string   `json:"commodityName,omitempty" xml:"commodityName,omitempty"`
	CpuNum                 *int64    `json:"cpuNum,omitempty" xml:"cpuNum,omitempty"`
	ExpireDate             *int64    `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	MemNum                 *int64    `json:"memNum,omitempty" xml:"memNum,omitempty"`
	NextBuyActions         []*string `json:"nextBuyActions,omitempty" xml:"nextBuyActions,omitempty" type:"Repeated"`
	NtmCommodityInstanceId *string   `json:"ntmCommodityInstanceId,omitempty" xml:"ntmCommodityInstanceId,omitempty"`
	OpenDate               *int64    `json:"openDate,omitempty" xml:"openDate,omitempty"`
	RelatedSubProducts     []*string `json:"relatedSubProducts,omitempty" xml:"relatedSubProducts,omitempty" type:"Repeated"`
	RemainingTime          *string   `json:"remainingTime,omitempty" xml:"remainingTime,omitempty"`
	SpecCode               *string   `json:"specCode,omitempty" xml:"specCode,omitempty"`
}

func (s QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetApplicationNum(v string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.ApplicationNum = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetCommercializeStatus(v string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.CommercializeStatus = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetCommodityCode(v string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.CommodityCode = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetCommodityName(v string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.CommodityName = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetCpuNum(v int64) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.CpuNum = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetExpireDate(v int64) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.ExpireDate = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetMemNum(v int64) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.MemNum = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetNextBuyActions(v []*string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.NextBuyActions = v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetNtmCommodityInstanceId(v string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.NtmCommodityInstanceId = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetOpenDate(v int64) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.OpenDate = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetRelatedSubProducts(v []*string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.RelatedSubProducts = v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetRemainingTime(v string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.RemainingTime = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec) SetSpecCode(v string) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommoditySpec {
	s.SpecCode = &v
	return s
}

type QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommodityUsage struct {
	ApplicationNum *int64 `json:"applicationNum,omitempty" xml:"applicationNum,omitempty"`
	CpuNum         *int64 `json:"cpuNum,omitempty" xml:"cpuNum,omitempty"`
	MemNum         *int64 `json:"memNum,omitempty" xml:"memNum,omitempty"`
}

func (s QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommodityUsage) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommodityUsage) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommodityUsage) SetApplicationNum(v int64) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommodityUsage {
	s.ApplicationNum = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommodityUsage) SetCpuNum(v int64) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommodityUsage {
	s.CpuNum = &v
	return s
}

func (s *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommodityUsage) SetMemNum(v int64) *QueryUsageResponseBodyDataCommoditySpecUsageInfosPurchasedCommoditySpecUsagesCommodityUsage {
	s.MemNum = &v
	return s
}

type QueryUsageResponseBodyDataSubProductInfos struct {
	SubProduct                            *QueryUsageResponseBodyDataSubProductInfosSubProduct                              `json:"subProduct,omitempty" xml:"subProduct,omitempty" type:"Struct"`
	SubProductCommoditySpecsForGuideToBuy []*QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuy `json:"subProductCommoditySpecsForGuideToBuy,omitempty" xml:"subProductCommoditySpecsForGuideToBuy,omitempty" type:"Repeated"`
}

func (s QueryUsageResponseBodyDataSubProductInfos) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBodyDataSubProductInfos) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBodyDataSubProductInfos) SetSubProduct(v *QueryUsageResponseBodyDataSubProductInfosSubProduct) *QueryUsageResponseBodyDataSubProductInfos {
	s.SubProduct = v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfos) SetSubProductCommoditySpecsForGuideToBuy(v []*QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuy) *QueryUsageResponseBodyDataSubProductInfos {
	s.SubProductCommoditySpecsForGuideToBuy = v
	return s
}

type QueryUsageResponseBodyDataSubProductInfosSubProduct struct {
	Display          *bool                  `json:"display,omitempty" xml:"display,omitempty"`
	ExtendProperties map[string]interface{} `json:"extendProperties,omitempty" xml:"extendProperties,omitempty"`
	SubProductCode   *string                `json:"subProductCode,omitempty" xml:"subProductCode,omitempty"`
	SubProductName   *string                `json:"subProductName,omitempty" xml:"subProductName,omitempty"`
	SubProductState  *string                `json:"subProductState,omitempty" xml:"subProductState,omitempty"`
}

func (s QueryUsageResponseBodyDataSubProductInfosSubProduct) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBodyDataSubProductInfosSubProduct) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProduct) SetDisplay(v bool) *QueryUsageResponseBodyDataSubProductInfosSubProduct {
	s.Display = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProduct) SetExtendProperties(v map[string]interface{}) *QueryUsageResponseBodyDataSubProductInfosSubProduct {
	s.ExtendProperties = v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProduct) SetSubProductCode(v string) *QueryUsageResponseBodyDataSubProductInfosSubProduct {
	s.SubProductCode = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProduct) SetSubProductName(v string) *QueryUsageResponseBodyDataSubProductInfosSubProduct {
	s.SubProductName = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProduct) SetSubProductState(v string) *QueryUsageResponseBodyDataSubProductInfosSubProduct {
	s.SubProductState = &v
	return s
}

type QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuy struct {
	CommodityConfig *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig `json:"commodityConfig,omitempty" xml:"commodityConfig,omitempty" type:"Struct"`
	CommoditySpec   *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec   `json:"commoditySpec,omitempty" xml:"commoditySpec,omitempty" type:"Struct"`
	ReleaseTime     *int64                                                                                         `json:"releaseTime,omitempty" xml:"releaseTime,omitempty"`
}

func (s QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuy) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuy) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuy) SetCommodityConfig(v *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuy {
	s.CommodityConfig = v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuy) SetCommoditySpec(v *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuy {
	s.CommoditySpec = v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuy) SetReleaseTime(v int64) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuy {
	s.ReleaseTime = &v
	return s
}

type QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig struct {
	BetaDeadlineMs        *int64    `json:"betaDeadlineMs,omitempty" xml:"betaDeadlineMs,omitempty"`
	CommodityCode         *string   `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	CommodityName         *string   `json:"commodityName,omitempty" xml:"commodityName,omitempty"`
	ExpiredToReleasedHour *int64    `json:"expiredToReleasedHour,omitempty" xml:"expiredToReleasedHour,omitempty"`
	InBeta                *bool     `json:"inBeta,omitempty" xml:"inBeta,omitempty"`
	Independent           *bool     `json:"independent,omitempty" xml:"independent,omitempty"`
	PaidType              *string   `json:"paidType,omitempty" xml:"paidType,omitempty"`
	ProductCode           *string   `json:"productCode,omitempty" xml:"productCode,omitempty"`
	QuotaItems            []*string `json:"quotaItems,omitempty" xml:"quotaItems,omitempty" type:"Repeated"`
	SpecCodes             []*string `json:"specCodes,omitempty" xml:"specCodes,omitempty" type:"Repeated"`
}

func (s QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) SetBetaDeadlineMs(v int64) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig {
	s.BetaDeadlineMs = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) SetCommodityCode(v string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig {
	s.CommodityCode = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) SetCommodityName(v string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig {
	s.CommodityName = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) SetExpiredToReleasedHour(v int64) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig {
	s.ExpiredToReleasedHour = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) SetInBeta(v bool) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig {
	s.InBeta = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) SetIndependent(v bool) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig {
	s.Independent = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) SetPaidType(v string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig {
	s.PaidType = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) SetProductCode(v string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig {
	s.ProductCode = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) SetQuotaItems(v []*string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig {
	s.QuotaItems = v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig) SetSpecCodes(v []*string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommodityConfig {
	s.SpecCodes = v
	return s
}

type QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec struct {
	ApplicationNum         *int64    `json:"applicationNum,omitempty" xml:"applicationNum,omitempty"`
	CommercializeStatus    *string   `json:"commercializeStatus,omitempty" xml:"commercializeStatus,omitempty"`
	CommodityCode          *string   `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	CommodityName          *string   `json:"commodityName,omitempty" xml:"commodityName,omitempty"`
	CpuNum                 *int64    `json:"cpuNum,omitempty" xml:"cpuNum,omitempty"`
	ExpireDate             *int64    `json:"expireDate,omitempty" xml:"expireDate,omitempty"`
	MemNum                 *int64    `json:"memNum,omitempty" xml:"memNum,omitempty"`
	NextBuyActions         *string   `json:"nextBuyActions,omitempty" xml:"nextBuyActions,omitempty"`
	NtmCommodityInstanceId *string   `json:"ntmCommodityInstanceId,omitempty" xml:"ntmCommodityInstanceId,omitempty"`
	OpenDate               *int64    `json:"openDate,omitempty" xml:"openDate,omitempty"`
	RelatedSubProducts     []*string `json:"relatedSubProducts,omitempty" xml:"relatedSubProducts,omitempty" type:"Repeated"`
	SpecCode               *string   `json:"specCode,omitempty" xml:"specCode,omitempty"`
	SpecName               *string   `json:"specName,omitempty" xml:"specName,omitempty"`
}

func (s QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) GoString() string {
	return s.String()
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetApplicationNum(v int64) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.ApplicationNum = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetCommercializeStatus(v string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.CommercializeStatus = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetCommodityCode(v string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.CommodityCode = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetCommodityName(v string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.CommodityName = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetCpuNum(v int64) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.CpuNum = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetExpireDate(v int64) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.ExpireDate = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetMemNum(v int64) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.MemNum = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetNextBuyActions(v string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.NextBuyActions = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetNtmCommodityInstanceId(v string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.NtmCommodityInstanceId = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetOpenDate(v int64) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.OpenDate = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetRelatedSubProducts(v []*string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.RelatedSubProducts = v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetSpecCode(v string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.SpecCode = &v
	return s
}

func (s *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec) SetSpecName(v string) *QueryUsageResponseBodyDataSubProductInfosSubProductCommoditySpecsForGuideToBuyCommoditySpec {
	s.SpecName = &v
	return s
}

type QueryUsageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUsageResponse) GoString() string {
	return s.String()
}

func (s *QueryUsageResponse) SetHeaders(v map[string]*string) *QueryUsageResponse {
	s.Headers = v
	return s
}

func (s *QueryUsageResponse) SetStatusCode(v int32) *QueryUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUsageResponse) SetBody(v *QueryUsageResponseBody) *QueryUsageResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("bizworks"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) QueryUsageWithOptions(runtime *util.RuntimeOptions) (_result *QueryUsageResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryUsage"),
		Version:     tea.String("2021-04-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUsage() (_result *QueryUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUsageResponse{}
	_body, _err := client.QueryUsageWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
