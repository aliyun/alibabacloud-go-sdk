// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddMemberBasicInfoRequest struct {
	Body *AddMemberBasicInfoRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s AddMemberBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMemberBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *AddMemberBasicInfoRequest) SetBody(v *AddMemberBasicInfoRequestBody) *AddMemberBasicInfoRequest {
	s.Body = v
	return s
}

type AddMemberBasicInfoRequestBody struct {
	Area   *string `json:"Area,omitempty" xml:"Area,omitempty"`
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// example:
	//
	// 2022-09-08
	Birthday *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	// This parameter is required.
	Channels []*AddMemberBasicInfoRequestBodyChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	City     *string                                  `json:"City,omitempty" xml:"City,omitempty"`
	Country  *string                                  `json:"Country,omitempty" xml:"Country,omitempty"`
	Email    *string                                  `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// ""
	Extra          *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	MemberName     *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberNickName *string `json:"MemberNickName,omitempty" xml:"MemberNickName,omitempty"`
	MixMobile      *string `json:"MixMobile,omitempty" xml:"MixMobile,omitempty"`
	// example:
	//
	// 17716699087
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// This parameter is required.
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	// This parameter is required.
	PlatFormType *string `json:"PlatFormType,omitempty" xml:"PlatFormType,omitempty"`
	Province     *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Sex          *string `json:"Sex,omitempty" xml:"Sex,omitempty"`
}

func (s AddMemberBasicInfoRequestBody) String() string {
	return tea.Prettify(s)
}

func (s AddMemberBasicInfoRequestBody) GoString() string {
	return s.String()
}

func (s *AddMemberBasicInfoRequestBody) SetArea(v string) *AddMemberBasicInfoRequestBody {
	s.Area = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetAvatar(v string) *AddMemberBasicInfoRequestBody {
	s.Avatar = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetBirthday(v string) *AddMemberBasicInfoRequestBody {
	s.Birthday = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetChannels(v []*AddMemberBasicInfoRequestBodyChannels) *AddMemberBasicInfoRequestBody {
	s.Channels = v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetCity(v string) *AddMemberBasicInfoRequestBody {
	s.City = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetCountry(v string) *AddMemberBasicInfoRequestBody {
	s.Country = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetEmail(v string) *AddMemberBasicInfoRequestBody {
	s.Email = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetExtra(v string) *AddMemberBasicInfoRequestBody {
	s.Extra = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetGmtCreate(v string) *AddMemberBasicInfoRequestBody {
	s.GmtCreate = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetMemberName(v string) *AddMemberBasicInfoRequestBody {
	s.MemberName = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetMemberNickName(v string) *AddMemberBasicInfoRequestBody {
	s.MemberNickName = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetMixMobile(v string) *AddMemberBasicInfoRequestBody {
	s.MixMobile = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetMobile(v string) *AddMemberBasicInfoRequestBody {
	s.Mobile = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetOpenMerchantId(v string) *AddMemberBasicInfoRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetPlatFormType(v string) *AddMemberBasicInfoRequestBody {
	s.PlatFormType = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetProvince(v string) *AddMemberBasicInfoRequestBody {
	s.Province = &v
	return s
}

func (s *AddMemberBasicInfoRequestBody) SetSex(v string) *AddMemberBasicInfoRequestBody {
	s.Sex = &v
	return s
}

type AddMemberBasicInfoRequestBodyChannels struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	ChannelCode    *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	ChannelOpenId  *string `json:"ChannelOpenId,omitempty" xml:"ChannelOpenId,omitempty"`
	ChannelUnionId *string `json:"ChannelUnionId,omitempty" xml:"ChannelUnionId,omitempty"`
	Scene          *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s AddMemberBasicInfoRequestBodyChannels) String() string {
	return tea.Prettify(s)
}

func (s AddMemberBasicInfoRequestBodyChannels) GoString() string {
	return s.String()
}

func (s *AddMemberBasicInfoRequestBodyChannels) SetAppId(v string) *AddMemberBasicInfoRequestBodyChannels {
	s.AppId = &v
	return s
}

func (s *AddMemberBasicInfoRequestBodyChannels) SetChannelCode(v string) *AddMemberBasicInfoRequestBodyChannels {
	s.ChannelCode = &v
	return s
}

func (s *AddMemberBasicInfoRequestBodyChannels) SetChannelOpenId(v string) *AddMemberBasicInfoRequestBodyChannels {
	s.ChannelOpenId = &v
	return s
}

func (s *AddMemberBasicInfoRequestBodyChannels) SetChannelUnionId(v string) *AddMemberBasicInfoRequestBodyChannels {
	s.ChannelUnionId = &v
	return s
}

func (s *AddMemberBasicInfoRequestBodyChannels) SetScene(v string) *AddMemberBasicInfoRequestBodyChannels {
	s.Scene = &v
	return s
}

type AddMemberBasicInfoShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMemberBasicInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMemberBasicInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddMemberBasicInfoShrinkRequest) SetBodyShrink(v string) *AddMemberBasicInfoShrinkRequest {
	s.BodyShrink = &v
	return s
}

type AddMemberBasicInfoResponseBody struct {
	// example:
	//
	// Lydaas.QuickMember.SystemError
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// WMS_36606164948078_23218019
	OuterMemberId *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	// example:
	//
	// 1DEFC4F1-AF11-5A3C-93B9-2880768DA218
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddMemberBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMemberBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AddMemberBasicInfoResponseBody) SetErrorCode(v string) *AddMemberBasicInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddMemberBasicInfoResponseBody) SetErrorMessage(v string) *AddMemberBasicInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddMemberBasicInfoResponseBody) SetOuterMemberId(v string) *AddMemberBasicInfoResponseBody {
	s.OuterMemberId = &v
	return s
}

func (s *AddMemberBasicInfoResponseBody) SetRequestId(v string) *AddMemberBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMemberBasicInfoResponseBody) SetSuccess(v bool) *AddMemberBasicInfoResponseBody {
	s.Success = &v
	return s
}

type AddMemberBasicInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMemberBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMemberBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMemberBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *AddMemberBasicInfoResponse) SetHeaders(v map[string]*string) *AddMemberBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *AddMemberBasicInfoResponse) SetStatusCode(v int32) *AddMemberBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMemberBasicInfoResponse) SetBody(v *AddMemberBasicInfoResponseBody) *AddMemberBasicInfoResponse {
	s.Body = v
	return s
}

type BatchSaveOrderPopRequest struct {
	Orders []*BatchSaveOrderPopRequestOrders `json:"Orders,omitempty" xml:"Orders,omitempty" type:"Repeated"`
}

func (s BatchSaveOrderPopRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveOrderPopRequest) GoString() string {
	return s.String()
}

func (s *BatchSaveOrderPopRequest) SetOrders(v []*BatchSaveOrderPopRequestOrders) *BatchSaveOrderPopRequest {
	s.Orders = v
	return s
}

type BatchSaveOrderPopRequestOrders struct {
	// This parameter is required.
	//
	// example:
	//
	// doudian
	ChannelCode *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1699b2b974d444e3aa489f96457ef204
	ChannelOpenId *string `json:"ChannelOpenId,omitempty" xml:"ChannelOpenId,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Feature       *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testOpenMerchantId
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testOpenOrderId
	OpenOrderId *string `json:"OpenOrderId,omitempty" xml:"OpenOrderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-12-20 12:12:12
	OrderCreateTime *string `json:"OrderCreateTime,omitempty" xml:"OrderCreateTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	OrderPayment *string `json:"OrderPayment,omitempty" xml:"OrderPayment,omitempty"`
	PayTime      *string `json:"PayTime,omitempty" xml:"PayTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DOUDIAN
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testShopId
	ShopId *string `json:"ShopId,omitempty" xml:"ShopId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TRADE_FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	SubOrderModelList []*BatchSaveOrderPopRequestOrdersSubOrderModelList `json:"SubOrderModelList,omitempty" xml:"SubOrderModelList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	TotalFee     *string `json:"TotalFee,omitempty" xml:"TotalFee,omitempty"`
	BuyerOpenUid *string `json:"buyerOpenUid,omitempty" xml:"buyerOpenUid,omitempty"`
}

func (s BatchSaveOrderPopRequestOrders) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveOrderPopRequestOrders) GoString() string {
	return s.String()
}

func (s *BatchSaveOrderPopRequestOrders) SetChannelCode(v string) *BatchSaveOrderPopRequestOrders {
	s.ChannelCode = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetChannelOpenId(v string) *BatchSaveOrderPopRequestOrders {
	s.ChannelOpenId = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetEndTime(v string) *BatchSaveOrderPopRequestOrders {
	s.EndTime = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetFeature(v string) *BatchSaveOrderPopRequestOrders {
	s.Feature = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetOpenMerchantId(v string) *BatchSaveOrderPopRequestOrders {
	s.OpenMerchantId = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetOpenOrderId(v string) *BatchSaveOrderPopRequestOrders {
	s.OpenOrderId = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetOrderCreateTime(v string) *BatchSaveOrderPopRequestOrders {
	s.OrderCreateTime = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetOrderPayment(v string) *BatchSaveOrderPopRequestOrders {
	s.OrderPayment = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetPayTime(v string) *BatchSaveOrderPopRequestOrders {
	s.PayTime = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetPlatformType(v string) *BatchSaveOrderPopRequestOrders {
	s.PlatformType = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetShopId(v string) *BatchSaveOrderPopRequestOrders {
	s.ShopId = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetStatus(v string) *BatchSaveOrderPopRequestOrders {
	s.Status = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetSubOrderModelList(v []*BatchSaveOrderPopRequestOrdersSubOrderModelList) *BatchSaveOrderPopRequestOrders {
	s.SubOrderModelList = v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetTotalFee(v string) *BatchSaveOrderPopRequestOrders {
	s.TotalFee = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrders) SetBuyerOpenUid(v string) *BatchSaveOrderPopRequestOrders {
	s.BuyerOpenUid = &v
	return s
}

type BatchSaveOrderPopRequestOrdersSubOrderModelList struct {
	Feature *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testSubOrderId
	OpenSubOrderId *string `json:"OpenSubOrderId,omitempty" xml:"OpenSubOrderId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	OrderPayment *string `json:"OrderPayment,omitempty" xml:"OrderPayment,omitempty"`
	OutProductId *string `json:"OutProductId,omitempty" xml:"OutProductId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductAmount *string `json:"ProductAmount,omitempty" xml:"ProductAmount,omitempty"`
	ProductId     *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductName   *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	RefundStatus  *string `json:"RefundStatus,omitempty" xml:"RefundStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TRADE_FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	TotalFee *string `json:"TotalFee,omitempty" xml:"TotalFee,omitempty"`
}

func (s BatchSaveOrderPopRequestOrdersSubOrderModelList) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveOrderPopRequestOrdersSubOrderModelList) GoString() string {
	return s.String()
}

func (s *BatchSaveOrderPopRequestOrdersSubOrderModelList) SetFeature(v string) *BatchSaveOrderPopRequestOrdersSubOrderModelList {
	s.Feature = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrdersSubOrderModelList) SetOpenSubOrderId(v string) *BatchSaveOrderPopRequestOrdersSubOrderModelList {
	s.OpenSubOrderId = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrdersSubOrderModelList) SetOrderPayment(v string) *BatchSaveOrderPopRequestOrdersSubOrderModelList {
	s.OrderPayment = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrdersSubOrderModelList) SetOutProductId(v string) *BatchSaveOrderPopRequestOrdersSubOrderModelList {
	s.OutProductId = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrdersSubOrderModelList) SetProductAmount(v string) *BatchSaveOrderPopRequestOrdersSubOrderModelList {
	s.ProductAmount = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrdersSubOrderModelList) SetProductId(v string) *BatchSaveOrderPopRequestOrdersSubOrderModelList {
	s.ProductId = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrdersSubOrderModelList) SetProductName(v string) *BatchSaveOrderPopRequestOrdersSubOrderModelList {
	s.ProductName = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrdersSubOrderModelList) SetRefundStatus(v string) *BatchSaveOrderPopRequestOrdersSubOrderModelList {
	s.RefundStatus = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrdersSubOrderModelList) SetStatus(v string) *BatchSaveOrderPopRequestOrdersSubOrderModelList {
	s.Status = &v
	return s
}

func (s *BatchSaveOrderPopRequestOrdersSubOrderModelList) SetTotalFee(v string) *BatchSaveOrderPopRequestOrdersSubOrderModelList {
	s.TotalFee = &v
	return s
}

type BatchSaveOrderPopShrinkRequest struct {
	OrdersShrink *string `json:"Orders,omitempty" xml:"Orders,omitempty"`
}

func (s BatchSaveOrderPopShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveOrderPopShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchSaveOrderPopShrinkRequest) SetOrdersShrink(v string) *BatchSaveOrderPopShrinkRequest {
	s.OrdersShrink = &v
	return s
}

type BatchSaveOrderPopResponseBody struct {
	// example:
	//
	// testErrorCode
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// testErrorMessage
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200/400...
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchSaveOrderPopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveOrderPopResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSaveOrderPopResponseBody) SetErrorCode(v string) *BatchSaveOrderPopResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchSaveOrderPopResponseBody) SetErrorMessage(v string) *BatchSaveOrderPopResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchSaveOrderPopResponseBody) SetHttpStatusCode(v string) *BatchSaveOrderPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchSaveOrderPopResponseBody) SetRequestId(v string) *BatchSaveOrderPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSaveOrderPopResponseBody) SetSuccess(v bool) *BatchSaveOrderPopResponseBody {
	s.Success = &v
	return s
}

type BatchSaveOrderPopResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSaveOrderPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSaveOrderPopResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSaveOrderPopResponse) GoString() string {
	return s.String()
}

func (s *BatchSaveOrderPopResponse) SetHeaders(v map[string]*string) *BatchSaveOrderPopResponse {
	s.Headers = v
	return s
}

func (s *BatchSaveOrderPopResponse) SetStatusCode(v int32) *BatchSaveOrderPopResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSaveOrderPopResponse) SetBody(v *BatchSaveOrderPopResponseBody) *BatchSaveOrderPopResponse {
	s.Body = v
	return s
}

type CalculateMemberLevelRequest struct {
	Body *CalculateMemberLevelRequestBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
}

func (s CalculateMemberLevelRequest) String() string {
	return tea.Prettify(s)
}

func (s CalculateMemberLevelRequest) GoString() string {
	return s.String()
}

func (s *CalculateMemberLevelRequest) SetBody(v *CalculateMemberLevelRequestBody) *CalculateMemberLevelRequest {
	s.Body = v
	return s
}

type CalculateMemberLevelRequestBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentGrade *int32 `json:"CurrentGrade,omitempty" xml:"CurrentGrade,omitempty"`
	// This parameter is required.
	CurrentGradeName *string `json:"CurrentGradeName,omitempty" xml:"CurrentGradeName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4000006009418358
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1230094
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TAOBAO
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	// example:
	//
	// 0
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0235b7f20a11de9e2bf4c3494b6d998d
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s CalculateMemberLevelRequestBody) String() string {
	return tea.Prettify(s)
}

func (s CalculateMemberLevelRequestBody) GoString() string {
	return s.String()
}

func (s *CalculateMemberLevelRequestBody) SetCurrentGrade(v int32) *CalculateMemberLevelRequestBody {
	s.CurrentGrade = &v
	return s
}

func (s *CalculateMemberLevelRequestBody) SetCurrentGradeName(v string) *CalculateMemberLevelRequestBody {
	s.CurrentGradeName = &v
	return s
}

func (s *CalculateMemberLevelRequestBody) SetMemberId(v string) *CalculateMemberLevelRequestBody {
	s.MemberId = &v
	return s
}

func (s *CalculateMemberLevelRequestBody) SetOpenMerchantId(v string) *CalculateMemberLevelRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *CalculateMemberLevelRequestBody) SetPlatformType(v string) *CalculateMemberLevelRequestBody {
	s.PlatformType = &v
	return s
}

func (s *CalculateMemberLevelRequestBody) SetScore(v string) *CalculateMemberLevelRequestBody {
	s.Score = &v
	return s
}

func (s *CalculateMemberLevelRequestBody) SetSerialNo(v string) *CalculateMemberLevelRequestBody {
	s.SerialNo = &v
	return s
}

type CalculateMemberLevelShrinkRequest struct {
	BodyShrink *string `json:"Body,omitempty" xml:"Body,omitempty"`
}

func (s CalculateMemberLevelShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CalculateMemberLevelShrinkRequest) GoString() string {
	return s.String()
}

func (s *CalculateMemberLevelShrinkRequest) SetBodyShrink(v string) *CalculateMemberLevelShrinkRequest {
	s.BodyShrink = &v
	return s
}

type CalculateMemberLevelResponseBody struct {
	// example:
	//
	// Lydaas.QuickMember.SystemError
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 1
	Grade *string `json:"Grade,omitempty" xml:"Grade,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6062f9067f170700a2e7ef5a
	OuterMemberId *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	// example:
	//
	// 1DEFC4F1-AF11-5A3C-93B9-2880768DA218
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CalculateMemberLevelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CalculateMemberLevelResponseBody) GoString() string {
	return s.String()
}

func (s *CalculateMemberLevelResponseBody) SetErrorCode(v string) *CalculateMemberLevelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CalculateMemberLevelResponseBody) SetErrorMessage(v string) *CalculateMemberLevelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CalculateMemberLevelResponseBody) SetGrade(v string) *CalculateMemberLevelResponseBody {
	s.Grade = &v
	return s
}

func (s *CalculateMemberLevelResponseBody) SetHttpStatusCode(v string) *CalculateMemberLevelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CalculateMemberLevelResponseBody) SetMessage(v string) *CalculateMemberLevelResponseBody {
	s.Message = &v
	return s
}

func (s *CalculateMemberLevelResponseBody) SetOuterMemberId(v string) *CalculateMemberLevelResponseBody {
	s.OuterMemberId = &v
	return s
}

func (s *CalculateMemberLevelResponseBody) SetRequestId(v string) *CalculateMemberLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CalculateMemberLevelResponseBody) SetSuccess(v bool) *CalculateMemberLevelResponseBody {
	s.Success = &v
	return s
}

type CalculateMemberLevelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CalculateMemberLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CalculateMemberLevelResponse) String() string {
	return tea.Prettify(s)
}

func (s CalculateMemberLevelResponse) GoString() string {
	return s.String()
}

func (s *CalculateMemberLevelResponse) SetHeaders(v map[string]*string) *CalculateMemberLevelResponse {
	s.Headers = v
	return s
}

func (s *CalculateMemberLevelResponse) SetStatusCode(v int32) *CalculateMemberLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *CalculateMemberLevelResponse) SetBody(v *CalculateMemberLevelResponseBody) *CalculateMemberLevelResponse {
	s.Body = v
	return s
}

type EditMemberBasicInfoRequest struct {
	Body *EditMemberBasicInfoRequestBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
}

func (s EditMemberBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s EditMemberBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *EditMemberBasicInfoRequest) SetBody(v *EditMemberBasicInfoRequestBody) *EditMemberBasicInfoRequest {
	s.Body = v
	return s
}

type EditMemberBasicInfoRequestBody struct {
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// example:
	//
	// xxx.jpg
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// example:
	//
	// 2024-06-20
	Birthday *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	Country  *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// xxx.com
	Email      *string `json:"Email,omitempty" xml:"Email,omitempty"`
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// example:
	//
	// xxxx
	MemberNickName *string `json:"MemberNickName,omitempty" xml:"MemberNickName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ***********
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1230094
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TAOBAO
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	Province     *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Sex          *string `json:"Sex,omitempty" xml:"Sex,omitempty"`
}

func (s EditMemberBasicInfoRequestBody) String() string {
	return tea.Prettify(s)
}

func (s EditMemberBasicInfoRequestBody) GoString() string {
	return s.String()
}

func (s *EditMemberBasicInfoRequestBody) SetArea(v string) *EditMemberBasicInfoRequestBody {
	s.Area = &v
	return s
}

func (s *EditMemberBasicInfoRequestBody) SetAvatar(v string) *EditMemberBasicInfoRequestBody {
	s.Avatar = &v
	return s
}

func (s *EditMemberBasicInfoRequestBody) SetBirthday(v string) *EditMemberBasicInfoRequestBody {
	s.Birthday = &v
	return s
}

func (s *EditMemberBasicInfoRequestBody) SetCity(v string) *EditMemberBasicInfoRequestBody {
	s.City = &v
	return s
}

func (s *EditMemberBasicInfoRequestBody) SetCountry(v string) *EditMemberBasicInfoRequestBody {
	s.Country = &v
	return s
}

func (s *EditMemberBasicInfoRequestBody) SetEmail(v string) *EditMemberBasicInfoRequestBody {
	s.Email = &v
	return s
}

func (s *EditMemberBasicInfoRequestBody) SetMemberName(v string) *EditMemberBasicInfoRequestBody {
	s.MemberName = &v
	return s
}

func (s *EditMemberBasicInfoRequestBody) SetMemberNickName(v string) *EditMemberBasicInfoRequestBody {
	s.MemberNickName = &v
	return s
}

func (s *EditMemberBasicInfoRequestBody) SetMobile(v string) *EditMemberBasicInfoRequestBody {
	s.Mobile = &v
	return s
}

func (s *EditMemberBasicInfoRequestBody) SetOpenMerchantId(v string) *EditMemberBasicInfoRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *EditMemberBasicInfoRequestBody) SetPlatformType(v string) *EditMemberBasicInfoRequestBody {
	s.PlatformType = &v
	return s
}

func (s *EditMemberBasicInfoRequestBody) SetProvince(v string) *EditMemberBasicInfoRequestBody {
	s.Province = &v
	return s
}

func (s *EditMemberBasicInfoRequestBody) SetSex(v string) *EditMemberBasicInfoRequestBody {
	s.Sex = &v
	return s
}

type EditMemberBasicInfoShrinkRequest struct {
	BodyShrink *string `json:"Body,omitempty" xml:"Body,omitempty"`
}

func (s EditMemberBasicInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s EditMemberBasicInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *EditMemberBasicInfoShrinkRequest) SetBodyShrink(v string) *EditMemberBasicInfoShrinkRequest {
	s.BodyShrink = &v
	return s
}

type EditMemberBasicInfoResponseBody struct {
	// example:
	//
	// Lydaas.QuickMember.SystemError
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1DEFC4F1-AF11-5A3C-93B9-2880768DA218
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditMemberBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditMemberBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *EditMemberBasicInfoResponseBody) SetErrorCode(v string) *EditMemberBasicInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *EditMemberBasicInfoResponseBody) SetErrorMessage(v string) *EditMemberBasicInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *EditMemberBasicInfoResponseBody) SetHttpStatusCode(v string) *EditMemberBasicInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *EditMemberBasicInfoResponseBody) SetMessage(v string) *EditMemberBasicInfoResponseBody {
	s.Message = &v
	return s
}

func (s *EditMemberBasicInfoResponseBody) SetRequestId(v string) *EditMemberBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditMemberBasicInfoResponseBody) SetSuccess(v bool) *EditMemberBasicInfoResponseBody {
	s.Success = &v
	return s
}

type EditMemberBasicInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EditMemberBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditMemberBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s EditMemberBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *EditMemberBasicInfoResponse) SetHeaders(v map[string]*string) *EditMemberBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *EditMemberBasicInfoResponse) SetStatusCode(v int32) *EditMemberBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *EditMemberBasicInfoResponse) SetBody(v *EditMemberBasicInfoResponseBody) *EditMemberBasicInfoResponse {
	s.Body = v
	return s
}

type MemberAccountDetailPageQueryRequest struct {
	Body *MemberAccountDetailPageQueryRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s MemberAccountDetailPageQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountDetailPageQueryRequest) GoString() string {
	return s.String()
}

func (s *MemberAccountDetailPageQueryRequest) SetBody(v *MemberAccountDetailPageQueryRequestBody) *MemberAccountDetailPageQueryRequest {
	s.Body = v
	return s
}

type MemberAccountDetailPageQueryRequestBody struct {
	// This parameter is required.
	//
	// example:
	//
	// 20
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// example:
	//
	// 2011-09-02 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 750cff00c3e0996d220ac2861dafdfadsf
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WMS_36606164948078_23218019
	OuterMemberId *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	PlatFormType *string `json:"PlatFormType,omitempty" xml:"PlatFormType,omitempty"`
	// example:
	//
	// 2011-09-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s MemberAccountDetailPageQueryRequestBody) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountDetailPageQueryRequestBody) GoString() string {
	return s.String()
}

func (s *MemberAccountDetailPageQueryRequestBody) SetAccountType(v int32) *MemberAccountDetailPageQueryRequestBody {
	s.AccountType = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetEndTime(v string) *MemberAccountDetailPageQueryRequestBody {
	s.EndTime = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetOpenMerchantId(v string) *MemberAccountDetailPageQueryRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetOuterMemberId(v string) *MemberAccountDetailPageQueryRequestBody {
	s.OuterMemberId = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetPage(v int32) *MemberAccountDetailPageQueryRequestBody {
	s.Page = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetPageSize(v int32) *MemberAccountDetailPageQueryRequestBody {
	s.PageSize = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetPlatFormType(v string) *MemberAccountDetailPageQueryRequestBody {
	s.PlatFormType = &v
	return s
}

func (s *MemberAccountDetailPageQueryRequestBody) SetStartTime(v string) *MemberAccountDetailPageQueryRequestBody {
	s.StartTime = &v
	return s
}

type MemberAccountDetailPageQueryShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MemberAccountDetailPageQueryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountDetailPageQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *MemberAccountDetailPageQueryShrinkRequest) SetBodyShrink(v string) *MemberAccountDetailPageQueryShrinkRequest {
	s.BodyShrink = &v
	return s
}

type MemberAccountDetailPageQueryResponseBody struct {
	Data []*MemberAccountDetailPageQueryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Lydaas.QuickMember.SystemError
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1DEFC4F1-AF11-5A3C-93B9-2880768DA218
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 5000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s MemberAccountDetailPageQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountDetailPageQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MemberAccountDetailPageQueryResponseBody) SetData(v []*MemberAccountDetailPageQueryResponseBodyData) *MemberAccountDetailPageQueryResponseBody {
	s.Data = v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBody) SetErrorCode(v string) *MemberAccountDetailPageQueryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBody) SetErrorMessage(v string) *MemberAccountDetailPageQueryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBody) SetRequestId(v string) *MemberAccountDetailPageQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBody) SetSuccess(v string) *MemberAccountDetailPageQueryResponseBody {
	s.Success = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBody) SetTotalCount(v int32) *MemberAccountDetailPageQueryResponseBody {
	s.TotalCount = &v
	return s
}

type MemberAccountDetailPageQueryResponseBodyData struct {
	// example:
	//
	// 100
	AccountBalance *string `json:"AccountBalance,omitempty" xml:"AccountBalance,omitempty"`
	// example:
	//
	// 20
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// This parameter is required.
	ActivityType *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	ChannelCode  *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	// example:
	//
	// 20
	DetailValue *string `json:"DetailValue,omitempty" xml:"DetailValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ...
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// This parameter is required.
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 750cff00c3e0996d220ac2861dafdfadsf
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	// example:
	//
	// 1
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// MEMBER_fc498a12edd84dafd
	OuterMemberId *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	// example:
	//
	// ...
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s MemberAccountDetailPageQueryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountDetailPageQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetAccountBalance(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.AccountBalance = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetAccountType(v int32) *MemberAccountDetailPageQueryResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetActivityType(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.ActivityType = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetChannelCode(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.ChannelCode = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetDetailValue(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.DetailValue = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetExtra(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.Extra = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetGmtCreate(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetGmtModified(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetOpenMerchantId(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.OpenMerchantId = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetOperateType(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.OperateType = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetOuterMemberId(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.OuterMemberId = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponseBodyData) SetRemark(v string) *MemberAccountDetailPageQueryResponseBodyData {
	s.Remark = &v
	return s
}

type MemberAccountDetailPageQueryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MemberAccountDetailPageQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MemberAccountDetailPageQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s MemberAccountDetailPageQueryResponse) GoString() string {
	return s.String()
}

func (s *MemberAccountDetailPageQueryResponse) SetHeaders(v map[string]*string) *MemberAccountDetailPageQueryResponse {
	s.Headers = v
	return s
}

func (s *MemberAccountDetailPageQueryResponse) SetStatusCode(v int32) *MemberAccountDetailPageQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *MemberAccountDetailPageQueryResponse) SetBody(v *MemberAccountDetailPageQueryResponseBody) *MemberAccountDetailPageQueryResponse {
	s.Body = v
	return s
}

type MemberPointChangeRequest struct {
	Body *MemberPointChangeRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s MemberPointChangeRequest) String() string {
	return tea.Prettify(s)
}

func (s MemberPointChangeRequest) GoString() string {
	return s.String()
}

func (s *MemberPointChangeRequest) SetBody(v *MemberPointChangeRequestBody) *MemberPointChangeRequest {
	s.Body = v
	return s
}

type MemberPointChangeRequestBody struct {
	// This parameter is required.
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	ChannelCode *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	// example:
	//
	// ""
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 750cff00c3e0996d220ac2861dafdfadsf
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WMS_36606164948078_23218019
	OuterMemberId *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	// This parameter is required.
	PlatFormType *string `json:"PlatFormType,omitempty" xml:"PlatFormType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4LJDNJH9JUX48L41
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s MemberPointChangeRequestBody) String() string {
	return tea.Prettify(s)
}

func (s MemberPointChangeRequestBody) GoString() string {
	return s.String()
}

func (s *MemberPointChangeRequestBody) SetAccountType(v string) *MemberPointChangeRequestBody {
	s.AccountType = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetChannelCode(v string) *MemberPointChangeRequestBody {
	s.ChannelCode = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetExtra(v string) *MemberPointChangeRequestBody {
	s.Extra = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetOpenMerchantId(v string) *MemberPointChangeRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetOperateType(v string) *MemberPointChangeRequestBody {
	s.OperateType = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetOuterMemberId(v string) *MemberPointChangeRequestBody {
	s.OuterMemberId = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetPlatFormType(v string) *MemberPointChangeRequestBody {
	s.PlatFormType = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetQuantity(v string) *MemberPointChangeRequestBody {
	s.Quantity = &v
	return s
}

func (s *MemberPointChangeRequestBody) SetSerialNo(v string) *MemberPointChangeRequestBody {
	s.SerialNo = &v
	return s
}

type MemberPointChangeShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MemberPointChangeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s MemberPointChangeShrinkRequest) GoString() string {
	return s.String()
}

func (s *MemberPointChangeShrinkRequest) SetBodyShrink(v string) *MemberPointChangeShrinkRequest {
	s.BodyShrink = &v
	return s
}

type MemberPointChangeResponseBody struct {
	AccountBalance *string `json:"AccountBalance,omitempty" xml:"AccountBalance,omitempty"`
	// example:
	//
	// Lydaas.QuickMember.SystemError
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	LevelName *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	// example:
	//
	// B2CD5682-12C0-51A7-82FC-1D36091CADAD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MemberPointChangeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MemberPointChangeResponseBody) GoString() string {
	return s.String()
}

func (s *MemberPointChangeResponseBody) SetAccountBalance(v string) *MemberPointChangeResponseBody {
	s.AccountBalance = &v
	return s
}

func (s *MemberPointChangeResponseBody) SetErrorCode(v string) *MemberPointChangeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *MemberPointChangeResponseBody) SetErrorMessage(v string) *MemberPointChangeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *MemberPointChangeResponseBody) SetLevelName(v string) *MemberPointChangeResponseBody {
	s.LevelName = &v
	return s
}

func (s *MemberPointChangeResponseBody) SetRequestId(v string) *MemberPointChangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *MemberPointChangeResponseBody) SetSuccess(v string) *MemberPointChangeResponseBody {
	s.Success = &v
	return s
}

type MemberPointChangeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MemberPointChangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MemberPointChangeResponse) String() string {
	return tea.Prettify(s)
}

func (s MemberPointChangeResponse) GoString() string {
	return s.String()
}

func (s *MemberPointChangeResponse) SetHeaders(v map[string]*string) *MemberPointChangeResponse {
	s.Headers = v
	return s
}

func (s *MemberPointChangeResponse) SetStatusCode(v int32) *MemberPointChangeResponse {
	s.StatusCode = &v
	return s
}

func (s *MemberPointChangeResponse) SetBody(v *MemberPointChangeResponseBody) *MemberPointChangeResponse {
	s.Body = v
	return s
}

type QueryMemberBasicInfoRequest struct {
	Body *QueryMemberBasicInfoRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s QueryMemberBasicInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMemberBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryMemberBasicInfoRequest) SetBody(v *QueryMemberBasicInfoRequestBody) *QueryMemberBasicInfoRequest {
	s.Body = v
	return s
}

type QueryMemberBasicInfoRequestBody struct {
	ChannelCode   *string `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	ChannelOpenId *string `json:"ChannelOpenId,omitempty" xml:"ChannelOpenId,omitempty"`
	// example:
	//
	// 18888889999
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 750cff00c3e0996d220ac2861dafdfadsf
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	// example:
	//
	// WMS_36606164948078_23218019
	OuterMemberId *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	// This parameter is required.
	PlatFormType *string `json:"PlatFormType,omitempty" xml:"PlatFormType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryMemberBasicInfoRequestBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMemberBasicInfoRequestBody) GoString() string {
	return s.String()
}

func (s *QueryMemberBasicInfoRequestBody) SetChannelCode(v string) *QueryMemberBasicInfoRequestBody {
	s.ChannelCode = &v
	return s
}

func (s *QueryMemberBasicInfoRequestBody) SetChannelOpenId(v string) *QueryMemberBasicInfoRequestBody {
	s.ChannelOpenId = &v
	return s
}

func (s *QueryMemberBasicInfoRequestBody) SetMobile(v string) *QueryMemberBasicInfoRequestBody {
	s.Mobile = &v
	return s
}

func (s *QueryMemberBasicInfoRequestBody) SetOpenMerchantId(v string) *QueryMemberBasicInfoRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *QueryMemberBasicInfoRequestBody) SetOuterMemberId(v string) *QueryMemberBasicInfoRequestBody {
	s.OuterMemberId = &v
	return s
}

func (s *QueryMemberBasicInfoRequestBody) SetPlatFormType(v string) *QueryMemberBasicInfoRequestBody {
	s.PlatFormType = &v
	return s
}

func (s *QueryMemberBasicInfoRequestBody) SetType(v string) *QueryMemberBasicInfoRequestBody {
	s.Type = &v
	return s
}

type QueryMemberBasicInfoShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMemberBasicInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMemberBasicInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMemberBasicInfoShrinkRequest) SetBodyShrink(v string) *QueryMemberBasicInfoShrinkRequest {
	s.BodyShrink = &v
	return s
}

type QueryMemberBasicInfoResponseBody struct {
	Data *QueryMemberBasicInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Lydaas.QuickMember.SystemError
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 1DEFC4F1-AF11-5A3C-93B9-2880768DA218
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMemberBasicInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMemberBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMemberBasicInfoResponseBody) SetData(v *QueryMemberBasicInfoResponseBodyData) *QueryMemberBasicInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryMemberBasicInfoResponseBody) SetErrorCode(v string) *QueryMemberBasicInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBody) SetErrorMessage(v string) *QueryMemberBasicInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBody) SetRequestId(v string) *QueryMemberBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBody) SetSuccess(v bool) *QueryMemberBasicInfoResponseBody {
	s.Success = &v
	return s
}

type QueryMemberBasicInfoResponseBodyData struct {
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// This parameter is required.
	Avatar           *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Birthday         *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	City             *string `json:"City,omitempty" xml:"City,omitempty"`
	Country          *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Email            *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Extra            *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	LevelName        *string `json:"LevelName,omitempty" xml:"LevelName,omitempty"`
	LevelNum         *string `json:"LevelNum,omitempty" xml:"LevelNum,omitempty"`
	MemberName       *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	MemberNickName   *string `json:"MemberNickName,omitempty" xml:"MemberNickName,omitempty"`
	Mobile           *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	NearExpiredScore *string `json:"NearExpiredScore,omitempty" xml:"NearExpiredScore,omitempty"`
	OpenMerchantId   *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	OuterMemberId    *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	Points           *string `json:"Points,omitempty" xml:"Points,omitempty"`
	Province         *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Score            *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Sex              *string `json:"Sex,omitempty" xml:"Sex,omitempty"`
}

func (s QueryMemberBasicInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMemberBasicInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMemberBasicInfoResponseBodyData) SetArea(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Area = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetAvatar(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Avatar = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetBirthday(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Birthday = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetCity(v string) *QueryMemberBasicInfoResponseBodyData {
	s.City = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetCountry(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Country = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetEmail(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Email = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetExtra(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Extra = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetLevelName(v string) *QueryMemberBasicInfoResponseBodyData {
	s.LevelName = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetLevelNum(v string) *QueryMemberBasicInfoResponseBodyData {
	s.LevelNum = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetMemberName(v string) *QueryMemberBasicInfoResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetMemberNickName(v string) *QueryMemberBasicInfoResponseBodyData {
	s.MemberNickName = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetMobile(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetNearExpiredScore(v string) *QueryMemberBasicInfoResponseBodyData {
	s.NearExpiredScore = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetOpenMerchantId(v string) *QueryMemberBasicInfoResponseBodyData {
	s.OpenMerchantId = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetOuterMemberId(v string) *QueryMemberBasicInfoResponseBodyData {
	s.OuterMemberId = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetPoints(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Points = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetProvince(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Province = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetScore(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Score = &v
	return s
}

func (s *QueryMemberBasicInfoResponseBodyData) SetSex(v string) *QueryMemberBasicInfoResponseBodyData {
	s.Sex = &v
	return s
}

type QueryMemberBasicInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMemberBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMemberBasicInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMemberBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryMemberBasicInfoResponse) SetHeaders(v map[string]*string) *QueryMemberBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryMemberBasicInfoResponse) SetStatusCode(v int32) *QueryMemberBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMemberBasicInfoResponse) SetBody(v *QueryMemberBasicInfoResponseBody) *QueryMemberBasicInfoResponse {
	s.Body = v
	return s
}

type SyncCardInfoRequest struct {
	Body *SyncCardInfoRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s SyncCardInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoRequest) GoString() string {
	return s.String()
}

func (s *SyncCardInfoRequest) SetBody(v *SyncCardInfoRequestBody) *SyncCardInfoRequest {
	s.Body = v
	return s
}

type SyncCardInfoRequestBody struct {
	// example:
	//
	// 1236437142867408
	BuyerId *string `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	// example:
	//
	// ""
	Extra   *string                         `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Gifters *SyncCardInfoRequestBodyGifters `json:"Gifters,omitempty" xml:"Gifters,omitempty" type:"Struct"`
	// example:
	//
	// 2023-05-12 00:00:00
	OccurredAt *string `json:"OccurredAt,omitempty" xml:"OccurredAt,omitempty"`
	// example:
	//
	// 750cff00c3e0996d220ac2861dafdfadsf
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	// example:
	//
	// 223332140970453
	OrderId       *string                                 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PresentDetail []*SyncCardInfoRequestBodyPresentDetail `json:"PresentDetail,omitempty" xml:"PresentDetail,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-05-12 00:00:00
	ReceivedAt *string                           `json:"ReceivedAt,omitempty" xml:"ReceivedAt,omitempty"`
	Recipient  *SyncCardInfoRequestBodyRecipient `json:"Recipient,omitempty" xml:"Recipient,omitempty" type:"Struct"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Theme  *string `json:"Theme,omitempty" xml:"Theme,omitempty"`
	// example:
	//
	// 3eecd2b8a6f946ca
	TransferId *string `json:"TransferId,omitempty" xml:"TransferId,omitempty"`
	// example:
	//
	// 2023-05-12 00:00:00
	TransferredAt *string `json:"TransferredAt,omitempty" xml:"TransferredAt,omitempty"`
}

func (s SyncCardInfoRequestBody) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoRequestBody) GoString() string {
	return s.String()
}

func (s *SyncCardInfoRequestBody) SetBuyerId(v string) *SyncCardInfoRequestBody {
	s.BuyerId = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetExtra(v string) *SyncCardInfoRequestBody {
	s.Extra = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetGifters(v *SyncCardInfoRequestBodyGifters) *SyncCardInfoRequestBody {
	s.Gifters = v
	return s
}

func (s *SyncCardInfoRequestBody) SetOccurredAt(v string) *SyncCardInfoRequestBody {
	s.OccurredAt = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetOpenMerchantId(v string) *SyncCardInfoRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetOrderId(v string) *SyncCardInfoRequestBody {
	s.OrderId = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetPresentDetail(v []*SyncCardInfoRequestBodyPresentDetail) *SyncCardInfoRequestBody {
	s.PresentDetail = v
	return s
}

func (s *SyncCardInfoRequestBody) SetReceivedAt(v string) *SyncCardInfoRequestBody {
	s.ReceivedAt = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetRecipient(v *SyncCardInfoRequestBodyRecipient) *SyncCardInfoRequestBody {
	s.Recipient = v
	return s
}

func (s *SyncCardInfoRequestBody) SetStatus(v string) *SyncCardInfoRequestBody {
	s.Status = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetTheme(v string) *SyncCardInfoRequestBody {
	s.Theme = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetTransferId(v string) *SyncCardInfoRequestBody {
	s.TransferId = &v
	return s
}

func (s *SyncCardInfoRequestBody) SetTransferredAt(v string) *SyncCardInfoRequestBody {
	s.TransferredAt = &v
	return s
}

type SyncCardInfoRequestBodyGifters struct {
	// example:
	//
	// https://xxx/2.jpg
	HeaderImg *string `json:"HeaderImg,omitempty" xml:"HeaderImg,omitempty"`
	// example:
	//
	// 1000030820003
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// openId
	//
	// example:
	//
	// 82ace612cd377134d597e32e91562caf
	OpenId *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	// example:
	//
	// 13277778888
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s SyncCardInfoRequestBodyGifters) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoRequestBodyGifters) GoString() string {
	return s.String()
}

func (s *SyncCardInfoRequestBodyGifters) SetHeaderImg(v string) *SyncCardInfoRequestBodyGifters {
	s.HeaderImg = &v
	return s
}

func (s *SyncCardInfoRequestBodyGifters) SetId(v string) *SyncCardInfoRequestBodyGifters {
	s.Id = &v
	return s
}

func (s *SyncCardInfoRequestBodyGifters) SetNickname(v string) *SyncCardInfoRequestBodyGifters {
	s.Nickname = &v
	return s
}

func (s *SyncCardInfoRequestBodyGifters) SetOpenId(v string) *SyncCardInfoRequestBodyGifters {
	s.OpenId = &v
	return s
}

func (s *SyncCardInfoRequestBodyGifters) SetPhone(v string) *SyncCardInfoRequestBodyGifters {
	s.Phone = &v
	return s
}

type SyncCardInfoRequestBodyPresentDetail struct {
	Count  *int64   `json:"Count,omitempty" xml:"Count,omitempty"`
	ItemId *string  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Name   *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Price  *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	SkuId  *string  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
}

func (s SyncCardInfoRequestBodyPresentDetail) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoRequestBodyPresentDetail) GoString() string {
	return s.String()
}

func (s *SyncCardInfoRequestBodyPresentDetail) SetCount(v int64) *SyncCardInfoRequestBodyPresentDetail {
	s.Count = &v
	return s
}

func (s *SyncCardInfoRequestBodyPresentDetail) SetItemId(v string) *SyncCardInfoRequestBodyPresentDetail {
	s.ItemId = &v
	return s
}

func (s *SyncCardInfoRequestBodyPresentDetail) SetName(v string) *SyncCardInfoRequestBodyPresentDetail {
	s.Name = &v
	return s
}

func (s *SyncCardInfoRequestBodyPresentDetail) SetPrice(v float64) *SyncCardInfoRequestBodyPresentDetail {
	s.Price = &v
	return s
}

func (s *SyncCardInfoRequestBodyPresentDetail) SetSkuId(v string) *SyncCardInfoRequestBodyPresentDetail {
	s.SkuId = &v
	return s
}

type SyncCardInfoRequestBodyRecipient struct {
	// example:
	//
	// https://xxx/1.jpg
	HeaderImg *string `json:"HeaderImg,omitempty" xml:"HeaderImg,omitempty"`
	// example:
	//
	// 1000030855004
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// openId
	//
	// example:
	//
	// 1179ead68b3833fea61a802ddb1dd3ac
	OpenId *string `json:"OpenId,omitempty" xml:"OpenId,omitempty"`
	// example:
	//
	// 18899998888
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s SyncCardInfoRequestBodyRecipient) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoRequestBodyRecipient) GoString() string {
	return s.String()
}

func (s *SyncCardInfoRequestBodyRecipient) SetHeaderImg(v string) *SyncCardInfoRequestBodyRecipient {
	s.HeaderImg = &v
	return s
}

func (s *SyncCardInfoRequestBodyRecipient) SetId(v string) *SyncCardInfoRequestBodyRecipient {
	s.Id = &v
	return s
}

func (s *SyncCardInfoRequestBodyRecipient) SetNickname(v string) *SyncCardInfoRequestBodyRecipient {
	s.Nickname = &v
	return s
}

func (s *SyncCardInfoRequestBodyRecipient) SetOpenId(v string) *SyncCardInfoRequestBodyRecipient {
	s.OpenId = &v
	return s
}

func (s *SyncCardInfoRequestBodyRecipient) SetPhone(v string) *SyncCardInfoRequestBodyRecipient {
	s.Phone = &v
	return s
}

type SyncCardInfoShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncCardInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *SyncCardInfoShrinkRequest) SetBodyShrink(v string) *SyncCardInfoShrinkRequest {
	s.BodyShrink = &v
	return s
}

type SyncCardInfoResponseBody struct {
	// example:
	//
	// Lydaas.QuickMember.SystemError
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1DEFC4F1-AF11-5A3C-93B9-2880768DA218
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncCardInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncCardInfoResponseBody) SetErrorCode(v string) *SyncCardInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SyncCardInfoResponseBody) SetErrorMessage(v string) *SyncCardInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SyncCardInfoResponseBody) SetHttpStatusCode(v string) *SyncCardInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SyncCardInfoResponseBody) SetMessage(v string) *SyncCardInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SyncCardInfoResponseBody) SetRequestId(v string) *SyncCardInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncCardInfoResponseBody) SetSuccess(v string) *SyncCardInfoResponseBody {
	s.Success = &v
	return s
}

type SyncCardInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncCardInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncCardInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncCardInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncCardInfoResponse) SetHeaders(v map[string]*string) *SyncCardInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncCardInfoResponse) SetStatusCode(v int32) *SyncCardInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncCardInfoResponse) SetBody(v *SyncCardInfoResponseBody) *SyncCardInfoResponse {
	s.Body = v
	return s
}

type SyncMemberBehaviorInfoRequest struct {
	Body *SyncMemberBehaviorInfoRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s SyncMemberBehaviorInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncMemberBehaviorInfoRequest) GoString() string {
	return s.String()
}

func (s *SyncMemberBehaviorInfoRequest) SetBody(v *SyncMemberBehaviorInfoRequestBody) *SyncMemberBehaviorInfoRequest {
	s.Body = v
	return s
}

type SyncMemberBehaviorInfoRequestBody struct {
	ActionDuration  *string `json:"ActionDuration,omitempty" xml:"ActionDuration,omitempty"`
	ActionEndDate   *string `json:"ActionEndDate,omitempty" xml:"ActionEndDate,omitempty"`
	ActionResult    *bool   `json:"ActionResult,omitempty" xml:"ActionResult,omitempty"`
	ActionStartDate *string `json:"ActionStartDate,omitempty" xml:"ActionStartDate,omitempty"`
	ActionSubType   *string `json:"ActionSubType,omitempty" xml:"ActionSubType,omitempty"`
	// This parameter is required.
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	Extra      *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	OpenMerchantId *string `json:"OpenMerchantId,omitempty" xml:"OpenMerchantId,omitempty"`
	// This parameter is required.
	OuterMemberId *string `json:"OuterMemberId,omitempty" xml:"OuterMemberId,omitempty"`
	// This parameter is required.
	PlatFormType *string `json:"PlatFormType,omitempty" xml:"PlatFormType,omitempty"`
}

func (s SyncMemberBehaviorInfoRequestBody) String() string {
	return tea.Prettify(s)
}

func (s SyncMemberBehaviorInfoRequestBody) GoString() string {
	return s.String()
}

func (s *SyncMemberBehaviorInfoRequestBody) SetActionDuration(v string) *SyncMemberBehaviorInfoRequestBody {
	s.ActionDuration = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetActionEndDate(v string) *SyncMemberBehaviorInfoRequestBody {
	s.ActionEndDate = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetActionResult(v bool) *SyncMemberBehaviorInfoRequestBody {
	s.ActionResult = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetActionStartDate(v string) *SyncMemberBehaviorInfoRequestBody {
	s.ActionStartDate = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetActionSubType(v string) *SyncMemberBehaviorInfoRequestBody {
	s.ActionSubType = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetActionType(v string) *SyncMemberBehaviorInfoRequestBody {
	s.ActionType = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetExtra(v string) *SyncMemberBehaviorInfoRequestBody {
	s.Extra = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetOpenMerchantId(v string) *SyncMemberBehaviorInfoRequestBody {
	s.OpenMerchantId = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetOuterMemberId(v string) *SyncMemberBehaviorInfoRequestBody {
	s.OuterMemberId = &v
	return s
}

func (s *SyncMemberBehaviorInfoRequestBody) SetPlatFormType(v string) *SyncMemberBehaviorInfoRequestBody {
	s.PlatFormType = &v
	return s
}

type SyncMemberBehaviorInfoShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncMemberBehaviorInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncMemberBehaviorInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *SyncMemberBehaviorInfoShrinkRequest) SetBodyShrink(v string) *SyncMemberBehaviorInfoShrinkRequest {
	s.BodyShrink = &v
	return s
}

type SyncMemberBehaviorInfoResponseBody struct {
	// example:
	//
	// Lydaas.QuickMember.SystemError
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 1DEFC4F1-AF11-5A3C-93B9-2880768DA218
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncMemberBehaviorInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncMemberBehaviorInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncMemberBehaviorInfoResponseBody) SetErrorCode(v string) *SyncMemberBehaviorInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SyncMemberBehaviorInfoResponseBody) SetErrorMessage(v string) *SyncMemberBehaviorInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SyncMemberBehaviorInfoResponseBody) SetRequestId(v string) *SyncMemberBehaviorInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncMemberBehaviorInfoResponseBody) SetSuccess(v bool) *SyncMemberBehaviorInfoResponseBody {
	s.Success = &v
	return s
}

type SyncMemberBehaviorInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncMemberBehaviorInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncMemberBehaviorInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncMemberBehaviorInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncMemberBehaviorInfoResponse) SetHeaders(v map[string]*string) *SyncMemberBehaviorInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncMemberBehaviorInfoResponse) SetStatusCode(v int32) *SyncMemberBehaviorInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncMemberBehaviorInfoResponse) SetBody(v *SyncMemberBehaviorInfoResponseBody) *SyncMemberBehaviorInfoResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("retailadvqa"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 开放平台同步会员基础信息
//
// @param tmpReq - AddMemberBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMemberBasicInfoResponse
func (client *Client) AddMemberBasicInfoWithOptions(tmpReq *AddMemberBasicInfoRequest, runtime *util.RuntimeOptions) (_result *AddMemberBasicInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddMemberBasicInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddMemberBasicInfo"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMemberBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开放平台同步会员基础信息
//
// @param request - AddMemberBasicInfoRequest
//
// @return AddMemberBasicInfoResponse
func (client *Client) AddMemberBasicInfo(request *AddMemberBasicInfoRequest) (_result *AddMemberBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddMemberBasicInfoResponse{}
	_body, _err := client.AddMemberBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量保存订单信息
//
// @param tmpReq - BatchSaveOrderPopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSaveOrderPopResponse
func (client *Client) BatchSaveOrderPopWithOptions(tmpReq *BatchSaveOrderPopRequest, runtime *util.RuntimeOptions) (_result *BatchSaveOrderPopResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchSaveOrderPopShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Orders)) {
		request.OrdersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Orders, tea.String("Orders"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrdersShrink)) {
		query["Orders"] = request.OrdersShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchSaveOrderPop"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSaveOrderPopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量保存订单信息
//
// @param request - BatchSaveOrderPopRequest
//
// @return BatchSaveOrderPopResponse
func (client *Client) BatchSaveOrderPop(request *BatchSaveOrderPopRequest) (_result *BatchSaveOrderPopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSaveOrderPopResponse{}
	_body, _err := client.BatchSaveOrderPopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会员等级初始化
//
// @param tmpReq - CalculateMemberLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CalculateMemberLevelResponse
func (client *Client) CalculateMemberLevelWithOptions(tmpReq *CalculateMemberLevelRequest, runtime *util.RuntimeOptions) (_result *CalculateMemberLevelResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CalculateMemberLevelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("Body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["Body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CalculateMemberLevel"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CalculateMemberLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 会员等级初始化
//
// @param request - CalculateMemberLevelRequest
//
// @return CalculateMemberLevelResponse
func (client *Client) CalculateMemberLevel(request *CalculateMemberLevelRequest) (_result *CalculateMemberLevelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CalculateMemberLevelResponse{}
	_body, _err := client.CalculateMemberLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会员信息编辑API
//
// @param tmpReq - EditMemberBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditMemberBasicInfoResponse
func (client *Client) EditMemberBasicInfoWithOptions(tmpReq *EditMemberBasicInfoRequest, runtime *util.RuntimeOptions) (_result *EditMemberBasicInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &EditMemberBasicInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("Body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["Body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EditMemberBasicInfo"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditMemberBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 会员信息编辑API
//
// @param request - EditMemberBasicInfoRequest
//
// @return EditMemberBasicInfoResponse
func (client *Client) EditMemberBasicInfo(request *EditMemberBasicInfoRequest) (_result *EditMemberBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditMemberBasicInfoResponse{}
	_body, _err := client.EditMemberBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开放平台会员积分明细查询
//
// @param tmpReq - MemberAccountDetailPageQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MemberAccountDetailPageQueryResponse
func (client *Client) MemberAccountDetailPageQueryWithOptions(tmpReq *MemberAccountDetailPageQueryRequest, runtime *util.RuntimeOptions) (_result *MemberAccountDetailPageQueryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &MemberAccountDetailPageQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MemberAccountDetailPageQuery"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MemberAccountDetailPageQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开放平台会员积分明细查询
//
// @param request - MemberAccountDetailPageQueryRequest
//
// @return MemberAccountDetailPageQueryResponse
func (client *Client) MemberAccountDetailPageQuery(request *MemberAccountDetailPageQueryRequest) (_result *MemberAccountDetailPageQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MemberAccountDetailPageQueryResponse{}
	_body, _err := client.MemberAccountDetailPageQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 会员积分变更
//
// @param tmpReq - MemberPointChangeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MemberPointChangeResponse
func (client *Client) MemberPointChangeWithOptions(tmpReq *MemberPointChangeRequest, runtime *util.RuntimeOptions) (_result *MemberPointChangeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &MemberPointChangeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MemberPointChange"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MemberPointChangeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 会员积分变更
//
// @param request - MemberPointChangeRequest
//
// @return MemberPointChangeResponse
func (client *Client) MemberPointChange(request *MemberPointChangeRequest) (_result *MemberPointChangeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MemberPointChangeResponse{}
	_body, _err := client.MemberPointChangeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询会员基础信息
//
// @param tmpReq - QueryMemberBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMemberBasicInfoResponse
func (client *Client) QueryMemberBasicInfoWithOptions(tmpReq *QueryMemberBasicInfoRequest, runtime *util.RuntimeOptions) (_result *QueryMemberBasicInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryMemberBasicInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMemberBasicInfo"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMemberBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会员基础信息
//
// @param request - QueryMemberBasicInfoRequest
//
// @return QueryMemberBasicInfoResponse
func (client *Client) QueryMemberBasicInfo(request *QueryMemberBasicInfoRequest) (_result *QueryMemberBasicInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMemberBasicInfoResponse{}
	_body, _err := client.QueryMemberBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 卡券信息同步
//
// @param tmpReq - SyncCardInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncCardInfoResponse
func (client *Client) SyncCardInfoWithOptions(tmpReq *SyncCardInfoRequest, runtime *util.RuntimeOptions) (_result *SyncCardInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SyncCardInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncCardInfo"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncCardInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 卡券信息同步
//
// @param request - SyncCardInfoRequest
//
// @return SyncCardInfoResponse
func (client *Client) SyncCardInfo(request *SyncCardInfoRequest) (_result *SyncCardInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncCardInfoResponse{}
	_body, _err := client.SyncCardInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开放平台会员行为信息同步
//
// @param tmpReq - SyncMemberBehaviorInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncMemberBehaviorInfoResponse
func (client *Client) SyncMemberBehaviorInfoWithOptions(tmpReq *SyncMemberBehaviorInfoRequest, runtime *util.RuntimeOptions) (_result *SyncMemberBehaviorInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SyncMemberBehaviorInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncMemberBehaviorInfo"),
		Version:     tea.String("2023-04-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncMemberBehaviorInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开放平台会员行为信息同步
//
// @param request - SyncMemberBehaviorInfoRequest
//
// @return SyncMemberBehaviorInfoResponse
func (client *Client) SyncMemberBehaviorInfo(request *SyncMemberBehaviorInfoRequest) (_result *SyncMemberBehaviorInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncMemberBehaviorInfoResponse{}
	_body, _err := client.SyncMemberBehaviorInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
