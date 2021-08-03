// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type SearchByPicRequest struct {
	PicContent *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	CategoryId *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Crop       *bool   `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Start      *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	Num        *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Fields     *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	RelationId *int64  `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	Pid        *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s SearchByPicRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchByPicRequest) GoString() string {
	return s.String()
}

func (s *SearchByPicRequest) SetPicContent(v string) *SearchByPicRequest {
	s.PicContent = &v
	return s
}

func (s *SearchByPicRequest) SetCategoryId(v int32) *SearchByPicRequest {
	s.CategoryId = &v
	return s
}

func (s *SearchByPicRequest) SetCrop(v bool) *SearchByPicRequest {
	s.Crop = &v
	return s
}

func (s *SearchByPicRequest) SetRegion(v string) *SearchByPicRequest {
	s.Region = &v
	return s
}

func (s *SearchByPicRequest) SetStart(v int32) *SearchByPicRequest {
	s.Start = &v
	return s
}

func (s *SearchByPicRequest) SetNum(v int32) *SearchByPicRequest {
	s.Num = &v
	return s
}

func (s *SearchByPicRequest) SetFields(v string) *SearchByPicRequest {
	s.Fields = &v
	return s
}

func (s *SearchByPicRequest) SetRelationId(v int64) *SearchByPicRequest {
	s.RelationId = &v
	return s
}

func (s *SearchByPicRequest) SetPid(v string) *SearchByPicRequest {
	s.Pid = &v
	return s
}

type SearchByPicAdvanceRequest struct {
	PicContentObject io.Reader `json:"PicContentObject,omitempty" xml:"PicContentObject,omitempty" require:"true"`
	CategoryId       *int32    `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Crop             *bool     `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Region           *string   `json:"Region,omitempty" xml:"Region,omitempty"`
	Start            *int32    `json:"Start,omitempty" xml:"Start,omitempty"`
	Num              *int32    `json:"Num,omitempty" xml:"Num,omitempty"`
	Fields           *string   `json:"Fields,omitempty" xml:"Fields,omitempty"`
	RelationId       *int64    `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	Pid              *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s SearchByPicAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchByPicAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SearchByPicAdvanceRequest) SetPicContentObject(v io.Reader) *SearchByPicAdvanceRequest {
	s.PicContentObject = v
	return s
}

func (s *SearchByPicAdvanceRequest) SetCategoryId(v int32) *SearchByPicAdvanceRequest {
	s.CategoryId = &v
	return s
}

func (s *SearchByPicAdvanceRequest) SetCrop(v bool) *SearchByPicAdvanceRequest {
	s.Crop = &v
	return s
}

func (s *SearchByPicAdvanceRequest) SetRegion(v string) *SearchByPicAdvanceRequest {
	s.Region = &v
	return s
}

func (s *SearchByPicAdvanceRequest) SetStart(v int32) *SearchByPicAdvanceRequest {
	s.Start = &v
	return s
}

func (s *SearchByPicAdvanceRequest) SetNum(v int32) *SearchByPicAdvanceRequest {
	s.Num = &v
	return s
}

func (s *SearchByPicAdvanceRequest) SetFields(v string) *SearchByPicAdvanceRequest {
	s.Fields = &v
	return s
}

func (s *SearchByPicAdvanceRequest) SetRelationId(v int64) *SearchByPicAdvanceRequest {
	s.RelationId = &v
	return s
}

func (s *SearchByPicAdvanceRequest) SetPid(v string) *SearchByPicAdvanceRequest {
	s.Pid = &v
	return s
}

type SearchByPicResponseBody struct {
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *SearchByPicResponseBodyData    `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PicInfo   *SearchByPicResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
}

func (s SearchByPicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchByPicResponseBody) GoString() string {
	return s.String()
}

func (s *SearchByPicResponseBody) SetSuccess(v bool) *SearchByPicResponseBody {
	s.Success = &v
	return s
}

func (s *SearchByPicResponseBody) SetCode(v int32) *SearchByPicResponseBody {
	s.Code = &v
	return s
}

func (s *SearchByPicResponseBody) SetMessage(v string) *SearchByPicResponseBody {
	s.Message = &v
	return s
}

func (s *SearchByPicResponseBody) SetData(v *SearchByPicResponseBodyData) *SearchByPicResponseBody {
	s.Data = v
	return s
}

func (s *SearchByPicResponseBody) SetRequestId(v string) *SearchByPicResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchByPicResponseBody) SetPicInfo(v *SearchByPicResponseBodyPicInfo) *SearchByPicResponseBody {
	s.PicInfo = v
	return s
}

type SearchByPicResponseBodyData struct {
	Auctions []*SearchByPicResponseBodyDataAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
}

func (s SearchByPicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchByPicResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchByPicResponseBodyData) SetAuctions(v []*SearchByPicResponseBodyDataAuctions) *SearchByPicResponseBodyData {
	s.Auctions = v
	return s
}

type SearchByPicResponseBodyDataAuctions struct {
	Result    *SearchByPicResponseBodyDataAuctionsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	RankScore *float32                                   `json:"RankScore,omitempty" xml:"RankScore,omitempty"`
}

func (s SearchByPicResponseBodyDataAuctions) String() string {
	return tea.Prettify(s)
}

func (s SearchByPicResponseBodyDataAuctions) GoString() string {
	return s.String()
}

func (s *SearchByPicResponseBodyDataAuctions) SetResult(v *SearchByPicResponseBodyDataAuctionsResult) *SearchByPicResponseBodyDataAuctions {
	s.Result = v
	return s
}

func (s *SearchByPicResponseBodyDataAuctions) SetRankScore(v float32) *SearchByPicResponseBodyDataAuctions {
	s.RankScore = &v
	return s
}

type SearchByPicResponseBodyDataAuctionsResult struct {
	ItemId               *string                                                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName             *string                                                 `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	Title                *string                                                 `json:"Title,omitempty" xml:"Title,omitempty"`
	Pic                  *string                                                 `json:"Pic,omitempty" xml:"Pic,omitempty"`
	PicUrl               *string                                                 `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	Price                *string                                                 `json:"Price,omitempty" xml:"Price,omitempty"`
	ReservePrice         *string                                                 `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	PromotionPrice       *string                                                 `json:"PromotionPrice,omitempty" xml:"PromotionPrice,omitempty"`
	ZkFinalPrice         *string                                                 `json:"ZkFinalPrice,omitempty" xml:"ZkFinalPrice,omitempty"`
	PriceAfterCoupon     *string                                                 `json:"PriceAfterCoupon,omitempty" xml:"PriceAfterCoupon,omitempty"`
	UserType             *int32                                                  `json:"UserType,omitempty" xml:"UserType,omitempty"`
	Provcity             *string                                                 `json:"Provcity,omitempty" xml:"Provcity,omitempty"`
	SellerNickName       *string                                                 `json:"SellerNickName,omitempty" xml:"SellerNickName,omitempty"`
	Nick                 *string                                                 `json:"Nick,omitempty" xml:"Nick,omitempty"`
	SellerId             *string                                                 `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	MonthSellCount       *int32                                                  `json:"MonthSellCount,omitempty" xml:"MonthSellCount,omitempty"`
	Volume               *int32                                                  `json:"Volume,omitempty" xml:"Volume,omitempty"`
	LevelOneCategoryName *string                                                 `json:"LevelOneCategoryName,omitempty" xml:"LevelOneCategoryName,omitempty"`
	CategoryName         *string                                                 `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	CouponActivityId     *string                                                 `json:"CouponActivityId,omitempty" xml:"CouponActivityId,omitempty"`
	CouponTotalCount     *string                                                 `json:"CouponTotalCount,omitempty" xml:"CouponTotalCount,omitempty"`
	CouponSendCount      *string                                                 `json:"CouponSendCount,omitempty" xml:"CouponSendCount,omitempty"`
	CouponRemainCount    *int32                                                  `json:"CouponRemainCount,omitempty" xml:"CouponRemainCount,omitempty"`
	CouponStartTime      *string                                                 `json:"CouponStartTime,omitempty" xml:"CouponStartTime,omitempty"`
	CouponEndTime        *string                                                 `json:"CouponEndTime,omitempty" xml:"CouponEndTime,omitempty"`
	CouponStartFee       *string                                                 `json:"CouponStartFee,omitempty" xml:"CouponStartFee,omitempty"`
	CouponAmount         *int32                                                  `json:"CouponAmount,omitempty" xml:"CouponAmount,omitempty"`
	CouponSaleTextInfo   *string                                                 `json:"CouponSaleTextInfo,omitempty" xml:"CouponSaleTextInfo,omitempty"`
	CouponInfo           *string                                                 `json:"CouponInfo,omitempty" xml:"CouponInfo,omitempty"`
	TkMktRate            *int32                                                  `json:"TkMktRate,omitempty" xml:"TkMktRate,omitempty"`
	TkRate               *int32                                                  `json:"TkRate,omitempty" xml:"TkRate,omitempty"`
	CommissionRate       *string                                                 `json:"CommissionRate,omitempty" xml:"CommissionRate,omitempty"`
	CouponShareUrl       *string                                                 `json:"CouponShareUrl,omitempty" xml:"CouponShareUrl,omitempty"`
	ClickUrl             *string                                                 `json:"ClickUrl,omitempty" xml:"ClickUrl,omitempty"`
	Url                  *string                                                 `json:"Url,omitempty" xml:"Url,omitempty"`
	ShortUrl             *string                                                 `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
	Key                  *string                                                 `json:"Key,omitempty" xml:"Key,omitempty"`
	ShopTitle            *string                                                 `json:"ShopTitle,omitempty" xml:"ShopTitle,omitempty"`
	MaxCommission        *SearchByPicResponseBodyDataAuctionsResultMaxCommission `json:"MaxCommission,omitempty" xml:"MaxCommission,omitempty" type:"Struct"`
}

func (s SearchByPicResponseBodyDataAuctionsResult) String() string {
	return tea.Prettify(s)
}

func (s SearchByPicResponseBodyDataAuctionsResult) GoString() string {
	return s.String()
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetItemId(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.ItemId = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetItemName(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.ItemName = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetTitle(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.Title = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetPic(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.Pic = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetPicUrl(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.PicUrl = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetPrice(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.Price = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetReservePrice(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.ReservePrice = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetPromotionPrice(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.PromotionPrice = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetZkFinalPrice(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.ZkFinalPrice = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetPriceAfterCoupon(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.PriceAfterCoupon = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetUserType(v int32) *SearchByPicResponseBodyDataAuctionsResult {
	s.UserType = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetProvcity(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.Provcity = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetSellerNickName(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.SellerNickName = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetNick(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.Nick = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetSellerId(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.SellerId = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetMonthSellCount(v int32) *SearchByPicResponseBodyDataAuctionsResult {
	s.MonthSellCount = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetVolume(v int32) *SearchByPicResponseBodyDataAuctionsResult {
	s.Volume = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetLevelOneCategoryName(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.LevelOneCategoryName = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCategoryName(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.CategoryName = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCouponActivityId(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.CouponActivityId = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCouponTotalCount(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.CouponTotalCount = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCouponSendCount(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.CouponSendCount = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCouponRemainCount(v int32) *SearchByPicResponseBodyDataAuctionsResult {
	s.CouponRemainCount = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCouponStartTime(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.CouponStartTime = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCouponEndTime(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.CouponEndTime = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCouponStartFee(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.CouponStartFee = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCouponAmount(v int32) *SearchByPicResponseBodyDataAuctionsResult {
	s.CouponAmount = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCouponSaleTextInfo(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.CouponSaleTextInfo = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCouponInfo(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.CouponInfo = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetTkMktRate(v int32) *SearchByPicResponseBodyDataAuctionsResult {
	s.TkMktRate = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetTkRate(v int32) *SearchByPicResponseBodyDataAuctionsResult {
	s.TkRate = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCommissionRate(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.CommissionRate = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetCouponShareUrl(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.CouponShareUrl = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetClickUrl(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.ClickUrl = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetUrl(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.Url = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetShortUrl(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.ShortUrl = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetKey(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.Key = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetShopTitle(v string) *SearchByPicResponseBodyDataAuctionsResult {
	s.ShopTitle = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResult) SetMaxCommission(v *SearchByPicResponseBodyDataAuctionsResultMaxCommission) *SearchByPicResponseBodyDataAuctionsResult {
	s.MaxCommission = v
	return s
}

type SearchByPicResponseBodyDataAuctionsResultMaxCommission struct {
	MaxCommissionRate           *string `json:"MaxCommissionRate,omitempty" xml:"MaxCommissionRate,omitempty"`
	MaxCommissionClickUrl       *string `json:"MaxCommissionClickUrl,omitempty" xml:"MaxCommissionClickUrl,omitempty"`
	MaxCommissionCouponShareUrl *string `json:"MaxCommissionCouponShareUrl,omitempty" xml:"MaxCommissionCouponShareUrl,omitempty"`
}

func (s SearchByPicResponseBodyDataAuctionsResultMaxCommission) String() string {
	return tea.Prettify(s)
}

func (s SearchByPicResponseBodyDataAuctionsResultMaxCommission) GoString() string {
	return s.String()
}

func (s *SearchByPicResponseBodyDataAuctionsResultMaxCommission) SetMaxCommissionRate(v string) *SearchByPicResponseBodyDataAuctionsResultMaxCommission {
	s.MaxCommissionRate = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResultMaxCommission) SetMaxCommissionClickUrl(v string) *SearchByPicResponseBodyDataAuctionsResultMaxCommission {
	s.MaxCommissionClickUrl = &v
	return s
}

func (s *SearchByPicResponseBodyDataAuctionsResultMaxCommission) SetMaxCommissionCouponShareUrl(v string) *SearchByPicResponseBodyDataAuctionsResultMaxCommission {
	s.MaxCommissionCouponShareUrl = &v
	return s
}

type SearchByPicResponseBodyPicInfo struct {
	MainRegion  *SearchByPicResponseBodyPicInfoMainRegion    `json:"MainRegion,omitempty" xml:"MainRegion,omitempty" type:"Struct"`
	MultiRegion []*SearchByPicResponseBodyPicInfoMultiRegion `json:"MultiRegion,omitempty" xml:"MultiRegion,omitempty" type:"Repeated"`
}

func (s SearchByPicResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchByPicResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchByPicResponseBodyPicInfo) SetMainRegion(v *SearchByPicResponseBodyPicInfoMainRegion) *SearchByPicResponseBodyPicInfo {
	s.MainRegion = v
	return s
}

func (s *SearchByPicResponseBodyPicInfo) SetMultiRegion(v []*SearchByPicResponseBodyPicInfoMultiRegion) *SearchByPicResponseBodyPicInfo {
	s.MultiRegion = v
	return s
}

type SearchByPicResponseBodyPicInfoMainRegion struct {
	Region          *string                                                    `json:"Region,omitempty" xml:"Region,omitempty"`
	MultiCategoryId []*SearchByPicResponseBodyPicInfoMainRegionMultiCategoryId `json:"MultiCategoryId,omitempty" xml:"MultiCategoryId,omitempty" type:"Repeated"`
}

func (s SearchByPicResponseBodyPicInfoMainRegion) String() string {
	return tea.Prettify(s)
}

func (s SearchByPicResponseBodyPicInfoMainRegion) GoString() string {
	return s.String()
}

func (s *SearchByPicResponseBodyPicInfoMainRegion) SetRegion(v string) *SearchByPicResponseBodyPicInfoMainRegion {
	s.Region = &v
	return s
}

func (s *SearchByPicResponseBodyPicInfoMainRegion) SetMultiCategoryId(v []*SearchByPicResponseBodyPicInfoMainRegionMultiCategoryId) *SearchByPicResponseBodyPicInfoMainRegion {
	s.MultiCategoryId = v
	return s
}

type SearchByPicResponseBodyPicInfoMainRegionMultiCategoryId struct {
	CategoryId *int32   `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Score      *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s SearchByPicResponseBodyPicInfoMainRegionMultiCategoryId) String() string {
	return tea.Prettify(s)
}

func (s SearchByPicResponseBodyPicInfoMainRegionMultiCategoryId) GoString() string {
	return s.String()
}

func (s *SearchByPicResponseBodyPicInfoMainRegionMultiCategoryId) SetCategoryId(v int32) *SearchByPicResponseBodyPicInfoMainRegionMultiCategoryId {
	s.CategoryId = &v
	return s
}

func (s *SearchByPicResponseBodyPicInfoMainRegionMultiCategoryId) SetScore(v float32) *SearchByPicResponseBodyPicInfoMainRegionMultiCategoryId {
	s.Score = &v
	return s
}

type SearchByPicResponseBodyPicInfoMultiRegion struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchByPicResponseBodyPicInfoMultiRegion) String() string {
	return tea.Prettify(s)
}

func (s SearchByPicResponseBodyPicInfoMultiRegion) GoString() string {
	return s.String()
}

func (s *SearchByPicResponseBodyPicInfoMultiRegion) SetRegion(v string) *SearchByPicResponseBodyPicInfoMultiRegion {
	s.Region = &v
	return s
}

type SearchByPicResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchByPicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchByPicResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchByPicResponse) GoString() string {
	return s.String()
}

func (s *SearchByPicResponse) SetHeaders(v map[string]*string) *SearchByPicResponse {
	s.Headers = v
	return s
}

func (s *SearchByPicResponse) SetBody(v *SearchByPicResponseBody) *SearchByPicResponse {
	s.Body = v
	return s
}

type SearchByUrlRequest struct {
	PicUrl     *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	CategoryId *int32  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Crop       *bool   `json:"Crop,omitempty" xml:"Crop,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Start      *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	Num        *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	Fields     *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	RelationId *int64  `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	Pid        *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s SearchByUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchByUrlRequest) GoString() string {
	return s.String()
}

func (s *SearchByUrlRequest) SetPicUrl(v string) *SearchByUrlRequest {
	s.PicUrl = &v
	return s
}

func (s *SearchByUrlRequest) SetCategoryId(v int32) *SearchByUrlRequest {
	s.CategoryId = &v
	return s
}

func (s *SearchByUrlRequest) SetCrop(v bool) *SearchByUrlRequest {
	s.Crop = &v
	return s
}

func (s *SearchByUrlRequest) SetRegion(v string) *SearchByUrlRequest {
	s.Region = &v
	return s
}

func (s *SearchByUrlRequest) SetStart(v int32) *SearchByUrlRequest {
	s.Start = &v
	return s
}

func (s *SearchByUrlRequest) SetNum(v int32) *SearchByUrlRequest {
	s.Num = &v
	return s
}

func (s *SearchByUrlRequest) SetFields(v string) *SearchByUrlRequest {
	s.Fields = &v
	return s
}

func (s *SearchByUrlRequest) SetRelationId(v int64) *SearchByUrlRequest {
	s.RelationId = &v
	return s
}

func (s *SearchByUrlRequest) SetPid(v string) *SearchByUrlRequest {
	s.Pid = &v
	return s
}

type SearchByUrlResponseBody struct {
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *SearchByUrlResponseBodyData    `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PicInfo   *SearchByUrlResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
}

func (s SearchByUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchByUrlResponseBody) GoString() string {
	return s.String()
}

func (s *SearchByUrlResponseBody) SetSuccess(v bool) *SearchByUrlResponseBody {
	s.Success = &v
	return s
}

func (s *SearchByUrlResponseBody) SetCode(v int32) *SearchByUrlResponseBody {
	s.Code = &v
	return s
}

func (s *SearchByUrlResponseBody) SetMessage(v string) *SearchByUrlResponseBody {
	s.Message = &v
	return s
}

func (s *SearchByUrlResponseBody) SetData(v *SearchByUrlResponseBodyData) *SearchByUrlResponseBody {
	s.Data = v
	return s
}

func (s *SearchByUrlResponseBody) SetRequestId(v string) *SearchByUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchByUrlResponseBody) SetPicInfo(v *SearchByUrlResponseBodyPicInfo) *SearchByUrlResponseBody {
	s.PicInfo = v
	return s
}

type SearchByUrlResponseBodyData struct {
	Auctions []*SearchByUrlResponseBodyDataAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
}

func (s SearchByUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchByUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchByUrlResponseBodyData) SetAuctions(v []*SearchByUrlResponseBodyDataAuctions) *SearchByUrlResponseBodyData {
	s.Auctions = v
	return s
}

type SearchByUrlResponseBodyDataAuctions struct {
	Result    *SearchByUrlResponseBodyDataAuctionsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	RankScore *float32                                   `json:"RankScore,omitempty" xml:"RankScore,omitempty"`
}

func (s SearchByUrlResponseBodyDataAuctions) String() string {
	return tea.Prettify(s)
}

func (s SearchByUrlResponseBodyDataAuctions) GoString() string {
	return s.String()
}

func (s *SearchByUrlResponseBodyDataAuctions) SetResult(v *SearchByUrlResponseBodyDataAuctionsResult) *SearchByUrlResponseBodyDataAuctions {
	s.Result = v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctions) SetRankScore(v float32) *SearchByUrlResponseBodyDataAuctions {
	s.RankScore = &v
	return s
}

type SearchByUrlResponseBodyDataAuctionsResult struct {
	ItemId               *string                                                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName             *string                                                 `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	Title                *string                                                 `json:"Title,omitempty" xml:"Title,omitempty"`
	Pic                  *string                                                 `json:"Pic,omitempty" xml:"Pic,omitempty"`
	PicUrl               *string                                                 `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	Price                *string                                                 `json:"Price,omitempty" xml:"Price,omitempty"`
	ReservePrice         *string                                                 `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	PromotionPrice       *string                                                 `json:"PromotionPrice,omitempty" xml:"PromotionPrice,omitempty"`
	ZkFinalPrice         *string                                                 `json:"ZkFinalPrice,omitempty" xml:"ZkFinalPrice,omitempty"`
	PriceAfterCoupon     *string                                                 `json:"PriceAfterCoupon,omitempty" xml:"PriceAfterCoupon,omitempty"`
	UserType             *int32                                                  `json:"UserType,omitempty" xml:"UserType,omitempty"`
	Provcity             *string                                                 `json:"Provcity,omitempty" xml:"Provcity,omitempty"`
	SellerNickName       *string                                                 `json:"SellerNickName,omitempty" xml:"SellerNickName,omitempty"`
	Nick                 *string                                                 `json:"Nick,omitempty" xml:"Nick,omitempty"`
	SellerId             *string                                                 `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	MonthSellCount       *int32                                                  `json:"MonthSellCount,omitempty" xml:"MonthSellCount,omitempty"`
	Volume               *int32                                                  `json:"Volume,omitempty" xml:"Volume,omitempty"`
	LevelOneCategoryName *string                                                 `json:"LevelOneCategoryName,omitempty" xml:"LevelOneCategoryName,omitempty"`
	CategoryName         *string                                                 `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	CouponActivityId     *string                                                 `json:"CouponActivityId,omitempty" xml:"CouponActivityId,omitempty"`
	CouponTotalCount     *string                                                 `json:"CouponTotalCount,omitempty" xml:"CouponTotalCount,omitempty"`
	CouponSendCount      *string                                                 `json:"CouponSendCount,omitempty" xml:"CouponSendCount,omitempty"`
	CouponRemainCount    *int32                                                  `json:"CouponRemainCount,omitempty" xml:"CouponRemainCount,omitempty"`
	CouponStartTime      *string                                                 `json:"CouponStartTime,omitempty" xml:"CouponStartTime,omitempty"`
	CouponEndTime        *string                                                 `json:"CouponEndTime,omitempty" xml:"CouponEndTime,omitempty"`
	CouponStartFee       *string                                                 `json:"CouponStartFee,omitempty" xml:"CouponStartFee,omitempty"`
	CouponAmount         *int32                                                  `json:"CouponAmount,omitempty" xml:"CouponAmount,omitempty"`
	CouponSaleTextInfo   *string                                                 `json:"CouponSaleTextInfo,omitempty" xml:"CouponSaleTextInfo,omitempty"`
	CouponInfo           *string                                                 `json:"CouponInfo,omitempty" xml:"CouponInfo,omitempty"`
	TkMktRate            *int32                                                  `json:"TkMktRate,omitempty" xml:"TkMktRate,omitempty"`
	TkRate               *int32                                                  `json:"TkRate,omitempty" xml:"TkRate,omitempty"`
	CommissionRate       *string                                                 `json:"CommissionRate,omitempty" xml:"CommissionRate,omitempty"`
	CouponShareUrl       *string                                                 `json:"CouponShareUrl,omitempty" xml:"CouponShareUrl,omitempty"`
	ClickUrl             *string                                                 `json:"ClickUrl,omitempty" xml:"ClickUrl,omitempty"`
	Url                  *string                                                 `json:"Url,omitempty" xml:"Url,omitempty"`
	ShortUrl             *string                                                 `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
	Key                  *string                                                 `json:"Key,omitempty" xml:"Key,omitempty"`
	ShopTitle            *string                                                 `json:"ShopTitle,omitempty" xml:"ShopTitle,omitempty"`
	MaxCommission        *SearchByUrlResponseBodyDataAuctionsResultMaxCommission `json:"MaxCommission,omitempty" xml:"MaxCommission,omitempty" type:"Struct"`
}

func (s SearchByUrlResponseBodyDataAuctionsResult) String() string {
	return tea.Prettify(s)
}

func (s SearchByUrlResponseBodyDataAuctionsResult) GoString() string {
	return s.String()
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetItemId(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.ItemId = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetItemName(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.ItemName = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetTitle(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.Title = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetPic(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.Pic = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetPicUrl(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.PicUrl = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetPrice(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.Price = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetReservePrice(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.ReservePrice = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetPromotionPrice(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.PromotionPrice = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetZkFinalPrice(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.ZkFinalPrice = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetPriceAfterCoupon(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.PriceAfterCoupon = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetUserType(v int32) *SearchByUrlResponseBodyDataAuctionsResult {
	s.UserType = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetProvcity(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.Provcity = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetSellerNickName(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.SellerNickName = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetNick(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.Nick = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetSellerId(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.SellerId = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetMonthSellCount(v int32) *SearchByUrlResponseBodyDataAuctionsResult {
	s.MonthSellCount = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetVolume(v int32) *SearchByUrlResponseBodyDataAuctionsResult {
	s.Volume = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetLevelOneCategoryName(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.LevelOneCategoryName = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCategoryName(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CategoryName = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCouponActivityId(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CouponActivityId = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCouponTotalCount(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CouponTotalCount = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCouponSendCount(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CouponSendCount = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCouponRemainCount(v int32) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CouponRemainCount = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCouponStartTime(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CouponStartTime = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCouponEndTime(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CouponEndTime = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCouponStartFee(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CouponStartFee = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCouponAmount(v int32) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CouponAmount = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCouponSaleTextInfo(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CouponSaleTextInfo = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCouponInfo(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CouponInfo = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetTkMktRate(v int32) *SearchByUrlResponseBodyDataAuctionsResult {
	s.TkMktRate = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetTkRate(v int32) *SearchByUrlResponseBodyDataAuctionsResult {
	s.TkRate = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCommissionRate(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CommissionRate = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetCouponShareUrl(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.CouponShareUrl = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetClickUrl(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.ClickUrl = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetUrl(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.Url = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetShortUrl(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.ShortUrl = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetKey(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.Key = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetShopTitle(v string) *SearchByUrlResponseBodyDataAuctionsResult {
	s.ShopTitle = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResult) SetMaxCommission(v *SearchByUrlResponseBodyDataAuctionsResultMaxCommission) *SearchByUrlResponseBodyDataAuctionsResult {
	s.MaxCommission = v
	return s
}

type SearchByUrlResponseBodyDataAuctionsResultMaxCommission struct {
	MaxCommissionRate           *string `json:"MaxCommissionRate,omitempty" xml:"MaxCommissionRate,omitempty"`
	MaxCommissionClickUrl       *string `json:"MaxCommissionClickUrl,omitempty" xml:"MaxCommissionClickUrl,omitempty"`
	MaxCommissionCouponShareUrl *string `json:"MaxCommissionCouponShareUrl,omitempty" xml:"MaxCommissionCouponShareUrl,omitempty"`
}

func (s SearchByUrlResponseBodyDataAuctionsResultMaxCommission) String() string {
	return tea.Prettify(s)
}

func (s SearchByUrlResponseBodyDataAuctionsResultMaxCommission) GoString() string {
	return s.String()
}

func (s *SearchByUrlResponseBodyDataAuctionsResultMaxCommission) SetMaxCommissionRate(v string) *SearchByUrlResponseBodyDataAuctionsResultMaxCommission {
	s.MaxCommissionRate = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResultMaxCommission) SetMaxCommissionClickUrl(v string) *SearchByUrlResponseBodyDataAuctionsResultMaxCommission {
	s.MaxCommissionClickUrl = &v
	return s
}

func (s *SearchByUrlResponseBodyDataAuctionsResultMaxCommission) SetMaxCommissionCouponShareUrl(v string) *SearchByUrlResponseBodyDataAuctionsResultMaxCommission {
	s.MaxCommissionCouponShareUrl = &v
	return s
}

type SearchByUrlResponseBodyPicInfo struct {
	MainRegion  *SearchByUrlResponseBodyPicInfoMainRegion    `json:"MainRegion,omitempty" xml:"MainRegion,omitempty" type:"Struct"`
	MultiRegion []*SearchByUrlResponseBodyPicInfoMultiRegion `json:"MultiRegion,omitempty" xml:"MultiRegion,omitempty" type:"Repeated"`
}

func (s SearchByUrlResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchByUrlResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchByUrlResponseBodyPicInfo) SetMainRegion(v *SearchByUrlResponseBodyPicInfoMainRegion) *SearchByUrlResponseBodyPicInfo {
	s.MainRegion = v
	return s
}

func (s *SearchByUrlResponseBodyPicInfo) SetMultiRegion(v []*SearchByUrlResponseBodyPicInfoMultiRegion) *SearchByUrlResponseBodyPicInfo {
	s.MultiRegion = v
	return s
}

type SearchByUrlResponseBodyPicInfoMainRegion struct {
	Region          *string                                                    `json:"Region,omitempty" xml:"Region,omitempty"`
	MultiCategoryId []*SearchByUrlResponseBodyPicInfoMainRegionMultiCategoryId `json:"MultiCategoryId,omitempty" xml:"MultiCategoryId,omitempty" type:"Repeated"`
}

func (s SearchByUrlResponseBodyPicInfoMainRegion) String() string {
	return tea.Prettify(s)
}

func (s SearchByUrlResponseBodyPicInfoMainRegion) GoString() string {
	return s.String()
}

func (s *SearchByUrlResponseBodyPicInfoMainRegion) SetRegion(v string) *SearchByUrlResponseBodyPicInfoMainRegion {
	s.Region = &v
	return s
}

func (s *SearchByUrlResponseBodyPicInfoMainRegion) SetMultiCategoryId(v []*SearchByUrlResponseBodyPicInfoMainRegionMultiCategoryId) *SearchByUrlResponseBodyPicInfoMainRegion {
	s.MultiCategoryId = v
	return s
}

type SearchByUrlResponseBodyPicInfoMainRegionMultiCategoryId struct {
	CategoryId *int32   `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Score      *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s SearchByUrlResponseBodyPicInfoMainRegionMultiCategoryId) String() string {
	return tea.Prettify(s)
}

func (s SearchByUrlResponseBodyPicInfoMainRegionMultiCategoryId) GoString() string {
	return s.String()
}

func (s *SearchByUrlResponseBodyPicInfoMainRegionMultiCategoryId) SetCategoryId(v int32) *SearchByUrlResponseBodyPicInfoMainRegionMultiCategoryId {
	s.CategoryId = &v
	return s
}

func (s *SearchByUrlResponseBodyPicInfoMainRegionMultiCategoryId) SetScore(v float32) *SearchByUrlResponseBodyPicInfoMainRegionMultiCategoryId {
	s.Score = &v
	return s
}

type SearchByUrlResponseBodyPicInfoMultiRegion struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchByUrlResponseBodyPicInfoMultiRegion) String() string {
	return tea.Prettify(s)
}

func (s SearchByUrlResponseBodyPicInfoMultiRegion) GoString() string {
	return s.String()
}

func (s *SearchByUrlResponseBodyPicInfoMultiRegion) SetRegion(v string) *SearchByUrlResponseBodyPicInfoMultiRegion {
	s.Region = &v
	return s
}

type SearchByUrlResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchByUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchByUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchByUrlResponse) GoString() string {
	return s.String()
}

func (s *SearchByUrlResponse) SetHeaders(v map[string]*string) *SearchByUrlResponse {
	s.Headers = v
	return s
}

func (s *SearchByUrlResponse) SetBody(v *SearchByUrlResponseBody) *SearchByUrlResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("imagesearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) SearchByPicWithOptions(request *SearchByPicRequest, runtime *util.RuntimeOptions) (_result *SearchByPicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchByPicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchByPic"), tea.String("2021-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchByPic(request *SearchByPicRequest) (_result *SearchByPicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchByPicResponse{}
	_body, _err := client.SearchByPicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchByPicAdvance(request *SearchByPicAdvanceRequest, runtime *util.RuntimeOptions) (_result *SearchByPicResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ImageSearch"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	searchByPicReq := &SearchByPicRequest{}
	openapiutil.Convert(request, searchByPicReq)
	if !tea.BoolValue(util.IsUnset(request.PicContentObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.ObjectKey,
			Content:     request.PicContentObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.AccessKeyId,
			Policy:              authResponse.EncodedPolicy,
			Signature:           authResponse.Signature,
			Key:                 authResponse.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		searchByPicReq.PicContent = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	searchByPicResp, _err := client.SearchByPicWithOptions(searchByPicReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = searchByPicResp
	return _result, _err
}

func (client *Client) SearchByUrlWithOptions(request *SearchByUrlRequest, runtime *util.RuntimeOptions) (_result *SearchByUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchByUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchByUrl"), tea.String("2021-05-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchByUrl(request *SearchByUrlRequest) (_result *SearchByUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchByUrlResponse{}
	_body, _err := client.SearchByUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
