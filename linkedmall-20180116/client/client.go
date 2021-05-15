// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddAddressRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	AddressInfo           *string `json:"AddressInfo,omitempty" xml:"AddressInfo,omitempty"`
}

func (s AddAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAddressRequest) GoString() string {
	return s.String()
}

func (s *AddAddressRequest) SetBizId(v string) *AddAddressRequest {
	s.BizId = &v
	return s
}

func (s *AddAddressRequest) SetThirdPartyUserId(v string) *AddAddressRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *AddAddressRequest) SetUseAnonymousTbAccount(v bool) *AddAddressRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *AddAddressRequest) SetAddressInfo(v string) *AddAddressRequest {
	s.AddressInfo = &v
	return s
}

type AddAddressResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AddressId *int64  `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
}

func (s AddAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AddAddressResponseBody) SetCode(v string) *AddAddressResponseBody {
	s.Code = &v
	return s
}

func (s *AddAddressResponseBody) SetMessage(v string) *AddAddressResponseBody {
	s.Message = &v
	return s
}

func (s *AddAddressResponseBody) SetRequestId(v string) *AddAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAddressResponseBody) SetAddressId(v int64) *AddAddressResponseBody {
	s.AddressId = &v
	return s
}

type AddAddressResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAddressResponse) GoString() string {
	return s.String()
}

func (s *AddAddressResponse) SetHeaders(v map[string]*string) *AddAddressResponse {
	s.Headers = v
	return s
}

func (s *AddAddressResponse) SetBody(v *AddAddressResponseBody) *AddAddressResponse {
	s.Body = v
	return s
}

type AddItemLimitRuleRequest struct {
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SubBizCode   *string `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	LmActivityId *int64  `json:"LmActivityId,omitempty" xml:"LmActivityId,omitempty"`
	LmItemId     *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId       *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	UpperNum     *int32  `json:"UpperNum,omitempty" xml:"UpperNum,omitempty"`
	RuleType     *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	BeginTime    *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s AddItemLimitRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddItemLimitRuleRequest) GoString() string {
	return s.String()
}

func (s *AddItemLimitRuleRequest) SetBizId(v string) *AddItemLimitRuleRequest {
	s.BizId = &v
	return s
}

func (s *AddItemLimitRuleRequest) SetSubBizCode(v string) *AddItemLimitRuleRequest {
	s.SubBizCode = &v
	return s
}

func (s *AddItemLimitRuleRequest) SetLmActivityId(v int64) *AddItemLimitRuleRequest {
	s.LmActivityId = &v
	return s
}

func (s *AddItemLimitRuleRequest) SetLmItemId(v string) *AddItemLimitRuleRequest {
	s.LmItemId = &v
	return s
}

func (s *AddItemLimitRuleRequest) SetItemId(v int64) *AddItemLimitRuleRequest {
	s.ItemId = &v
	return s
}

func (s *AddItemLimitRuleRequest) SetUpperNum(v int32) *AddItemLimitRuleRequest {
	s.UpperNum = &v
	return s
}

func (s *AddItemLimitRuleRequest) SetRuleType(v int32) *AddItemLimitRuleRequest {
	s.RuleType = &v
	return s
}

func (s *AddItemLimitRuleRequest) SetBeginTime(v int64) *AddItemLimitRuleRequest {
	s.BeginTime = &v
	return s
}

func (s *AddItemLimitRuleRequest) SetEndTime(v int64) *AddItemLimitRuleRequest {
	s.EndTime = &v
	return s
}

type AddItemLimitRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Model     *int64  `json:"Model,omitempty" xml:"Model,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleId    *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s AddItemLimitRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddItemLimitRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddItemLimitRuleResponseBody) SetCode(v string) *AddItemLimitRuleResponseBody {
	s.Code = &v
	return s
}

func (s *AddItemLimitRuleResponseBody) SetModel(v int64) *AddItemLimitRuleResponseBody {
	s.Model = &v
	return s
}

func (s *AddItemLimitRuleResponseBody) SetMessage(v string) *AddItemLimitRuleResponseBody {
	s.Message = &v
	return s
}

func (s *AddItemLimitRuleResponseBody) SetRequestId(v string) *AddItemLimitRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddItemLimitRuleResponseBody) SetRuleId(v int64) *AddItemLimitRuleResponseBody {
	s.RuleId = &v
	return s
}

type AddItemLimitRuleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddItemLimitRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddItemLimitRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddItemLimitRuleResponse) GoString() string {
	return s.String()
}

func (s *AddItemLimitRuleResponse) SetHeaders(v map[string]*string) *AddItemLimitRuleResponse {
	s.Headers = v
	return s
}

func (s *AddItemLimitRuleResponse) SetBody(v *AddItemLimitRuleResponseBody) *AddItemLimitRuleResponse {
	s.Body = v
	return s
}

type AddItemToSubBizsRequest struct {
	BizId     *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ItemId    *int64                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId  *string                `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	SubBizIds map[string]interface{} `json:"SubBizIds,omitempty" xml:"SubBizIds,omitempty"`
}

func (s AddItemToSubBizsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddItemToSubBizsRequest) GoString() string {
	return s.String()
}

func (s *AddItemToSubBizsRequest) SetBizId(v string) *AddItemToSubBizsRequest {
	s.BizId = &v
	return s
}

func (s *AddItemToSubBizsRequest) SetItemId(v int64) *AddItemToSubBizsRequest {
	s.ItemId = &v
	return s
}

func (s *AddItemToSubBizsRequest) SetLmItemId(v string) *AddItemToSubBizsRequest {
	s.LmItemId = &v
	return s
}

func (s *AddItemToSubBizsRequest) SetSubBizIds(v map[string]interface{}) *AddItemToSubBizsRequest {
	s.SubBizIds = v
	return s
}

type AddItemToSubBizsShrinkRequest struct {
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ItemId          *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId        *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	SubBizIdsShrink *string `json:"SubBizIds,omitempty" xml:"SubBizIds,omitempty"`
}

func (s AddItemToSubBizsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddItemToSubBizsShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddItemToSubBizsShrinkRequest) SetBizId(v string) *AddItemToSubBizsShrinkRequest {
	s.BizId = &v
	return s
}

func (s *AddItemToSubBizsShrinkRequest) SetItemId(v int64) *AddItemToSubBizsShrinkRequest {
	s.ItemId = &v
	return s
}

func (s *AddItemToSubBizsShrinkRequest) SetLmItemId(v string) *AddItemToSubBizsShrinkRequest {
	s.LmItemId = &v
	return s
}

func (s *AddItemToSubBizsShrinkRequest) SetSubBizIdsShrink(v string) *AddItemToSubBizsShrinkRequest {
	s.SubBizIdsShrink = &v
	return s
}

type AddItemToSubBizsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddItemToSubBizsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddItemToSubBizsResponseBody) GoString() string {
	return s.String()
}

func (s *AddItemToSubBizsResponseBody) SetCode(v string) *AddItemToSubBizsResponseBody {
	s.Code = &v
	return s
}

func (s *AddItemToSubBizsResponseBody) SetMessage(v string) *AddItemToSubBizsResponseBody {
	s.Message = &v
	return s
}

func (s *AddItemToSubBizsResponseBody) SetRequestId(v string) *AddItemToSubBizsResponseBody {
	s.RequestId = &v
	return s
}

type AddItemToSubBizsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddItemToSubBizsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddItemToSubBizsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddItemToSubBizsResponse) GoString() string {
	return s.String()
}

func (s *AddItemToSubBizsResponse) SetHeaders(v map[string]*string) *AddItemToSubBizsResponse {
	s.Headers = v
	return s
}

func (s *AddItemToSubBizsResponse) SetBody(v *AddItemToSubBizsResponseBody) *AddItemToSubBizsResponse {
	s.Body = v
	return s
}

type AddSupplierNewItemsRequest struct {
	BizId    *string                               `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ItemList []*AddSupplierNewItemsRequestItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
}

func (s AddSupplierNewItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSupplierNewItemsRequest) GoString() string {
	return s.String()
}

func (s *AddSupplierNewItemsRequest) SetBizId(v string) *AddSupplierNewItemsRequest {
	s.BizId = &v
	return s
}

func (s *AddSupplierNewItemsRequest) SetItemList(v []*AddSupplierNewItemsRequestItemList) *AddSupplierNewItemsRequest {
	s.ItemList = v
	return s
}

type AddSupplierNewItemsRequestItemList struct {
	LmItemId *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId   *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	SkuList  []*int  `json:"SkuList,omitempty" xml:"SkuList,omitempty" type:"Repeated"`
}

func (s AddSupplierNewItemsRequestItemList) String() string {
	return tea.Prettify(s)
}

func (s AddSupplierNewItemsRequestItemList) GoString() string {
	return s.String()
}

func (s *AddSupplierNewItemsRequestItemList) SetLmItemId(v string) *AddSupplierNewItemsRequestItemList {
	s.LmItemId = &v
	return s
}

func (s *AddSupplierNewItemsRequestItemList) SetItemId(v int64) *AddSupplierNewItemsRequestItemList {
	s.ItemId = &v
	return s
}

func (s *AddSupplierNewItemsRequestItemList) SetSkuList(v []*int) *AddSupplierNewItemsRequestItemList {
	s.SkuList = v
	return s
}

type AddSupplierNewItemsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSupplierNewItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSupplierNewItemsResponseBody) GoString() string {
	return s.String()
}

func (s *AddSupplierNewItemsResponseBody) SetCode(v string) *AddSupplierNewItemsResponseBody {
	s.Code = &v
	return s
}

func (s *AddSupplierNewItemsResponseBody) SetMessage(v string) *AddSupplierNewItemsResponseBody {
	s.Message = &v
	return s
}

func (s *AddSupplierNewItemsResponseBody) SetRequestId(v string) *AddSupplierNewItemsResponseBody {
	s.RequestId = &v
	return s
}

type AddSupplierNewItemsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSupplierNewItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSupplierNewItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSupplierNewItemsResponse) GoString() string {
	return s.String()
}

func (s *AddSupplierNewItemsResponse) SetHeaders(v map[string]*string) *AddSupplierNewItemsResponse {
	s.Headers = v
	return s
}

func (s *AddSupplierNewItemsResponse) SetBody(v *AddSupplierNewItemsResponseBody) *AddSupplierNewItemsResponse {
	s.Body = v
	return s
}

type ApplyRefundRequest struct {
	BizId                 *string                               `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string                               `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	SubLmOrderId          *string                               `json:"SubLmOrderId,omitempty" xml:"SubLmOrderId,omitempty"`
	BizClaimType          *int32                                `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	ApplyRefundFee        *int64                                `json:"ApplyRefundFee,omitempty" xml:"ApplyRefundFee,omitempty"`
	ApplyRefundCount      *int32                                `json:"ApplyRefundCount,omitempty" xml:"ApplyRefundCount,omitempty"`
	ApplyReasonTextId     *int64                                `json:"ApplyReasonTextId,omitempty" xml:"ApplyReasonTextId,omitempty"`
	LeaveMessage          *string                               `json:"LeaveMessage,omitempty" xml:"LeaveMessage,omitempty"`
	GoodsStatus           *int32                                `json:"GoodsStatus,omitempty" xml:"GoodsStatus,omitempty"`
	UseAnonymousTbAccount *bool                                 `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string                               `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string                               `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	LeavePictureList      []*ApplyRefundRequestLeavePictureList `json:"LeavePictureList,omitempty" xml:"LeavePictureList,omitempty" type:"Repeated"`
}

func (s ApplyRefundRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyRefundRequest) GoString() string {
	return s.String()
}

func (s *ApplyRefundRequest) SetBizId(v string) *ApplyRefundRequest {
	s.BizId = &v
	return s
}

func (s *ApplyRefundRequest) SetBizUid(v string) *ApplyRefundRequest {
	s.BizUid = &v
	return s
}

func (s *ApplyRefundRequest) SetSubLmOrderId(v string) *ApplyRefundRequest {
	s.SubLmOrderId = &v
	return s
}

func (s *ApplyRefundRequest) SetBizClaimType(v int32) *ApplyRefundRequest {
	s.BizClaimType = &v
	return s
}

func (s *ApplyRefundRequest) SetApplyRefundFee(v int64) *ApplyRefundRequest {
	s.ApplyRefundFee = &v
	return s
}

func (s *ApplyRefundRequest) SetApplyRefundCount(v int32) *ApplyRefundRequest {
	s.ApplyRefundCount = &v
	return s
}

func (s *ApplyRefundRequest) SetApplyReasonTextId(v int64) *ApplyRefundRequest {
	s.ApplyReasonTextId = &v
	return s
}

func (s *ApplyRefundRequest) SetLeaveMessage(v string) *ApplyRefundRequest {
	s.LeaveMessage = &v
	return s
}

func (s *ApplyRefundRequest) SetGoodsStatus(v int32) *ApplyRefundRequest {
	s.GoodsStatus = &v
	return s
}

func (s *ApplyRefundRequest) SetUseAnonymousTbAccount(v bool) *ApplyRefundRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *ApplyRefundRequest) SetThirdPartyUserId(v string) *ApplyRefundRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *ApplyRefundRequest) SetAccountType(v string) *ApplyRefundRequest {
	s.AccountType = &v
	return s
}

func (s *ApplyRefundRequest) SetLeavePictureList(v []*ApplyRefundRequestLeavePictureList) *ApplyRefundRequest {
	s.LeavePictureList = v
	return s
}

type ApplyRefundRequestLeavePictureList struct {
	Picture *string `json:"Picture,omitempty" xml:"Picture,omitempty"`
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s ApplyRefundRequestLeavePictureList) String() string {
	return tea.Prettify(s)
}

func (s ApplyRefundRequestLeavePictureList) GoString() string {
	return s.String()
}

func (s *ApplyRefundRequestLeavePictureList) SetPicture(v string) *ApplyRefundRequestLeavePictureList {
	s.Picture = &v
	return s
}

func (s *ApplyRefundRequestLeavePictureList) SetDesc(v string) *ApplyRefundRequestLeavePictureList {
	s.Desc = &v
	return s
}

type ApplyRefundResponseBody struct {
	Code                  *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message               *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId             *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RefundApplicationData *ApplyRefundResponseBodyRefundApplicationData `json:"RefundApplicationData,omitempty" xml:"RefundApplicationData,omitempty" type:"Struct"`
}

func (s ApplyRefundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyRefundResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyRefundResponseBody) SetCode(v string) *ApplyRefundResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyRefundResponseBody) SetMessage(v string) *ApplyRefundResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyRefundResponseBody) SetRequestId(v string) *ApplyRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyRefundResponseBody) SetRefundApplicationData(v *ApplyRefundResponseBodyRefundApplicationData) *ApplyRefundResponseBody {
	s.RefundApplicationData = v
	return s
}

type ApplyRefundResponseBodyRefundApplicationData struct {
	DisputeType   *int32  `json:"DisputeType,omitempty" xml:"DisputeType,omitempty"`
	SubLmOrderId  *string `json:"SubLmOrderId,omitempty" xml:"SubLmOrderId,omitempty"`
	DisputeStatus *int32  `json:"DisputeStatus,omitempty" xml:"DisputeStatus,omitempty"`
}

func (s ApplyRefundResponseBodyRefundApplicationData) String() string {
	return tea.Prettify(s)
}

func (s ApplyRefundResponseBodyRefundApplicationData) GoString() string {
	return s.String()
}

func (s *ApplyRefundResponseBodyRefundApplicationData) SetDisputeType(v int32) *ApplyRefundResponseBodyRefundApplicationData {
	s.DisputeType = &v
	return s
}

func (s *ApplyRefundResponseBodyRefundApplicationData) SetSubLmOrderId(v string) *ApplyRefundResponseBodyRefundApplicationData {
	s.SubLmOrderId = &v
	return s
}

func (s *ApplyRefundResponseBodyRefundApplicationData) SetDisputeStatus(v int32) *ApplyRefundResponseBodyRefundApplicationData {
	s.DisputeStatus = &v
	return s
}

type ApplyRefundResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyRefundResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyRefundResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyRefundResponse) GoString() string {
	return s.String()
}

func (s *ApplyRefundResponse) SetHeaders(v map[string]*string) *ApplyRefundResponse {
	s.Headers = v
	return s
}

func (s *ApplyRefundResponse) SetBody(v *ApplyRefundResponseBody) *ApplyRefundResponse {
	s.Body = v
	return s
}

type BatchRegistAnonymousTbAccountRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	IdJsonList *string `json:"IdJsonList,omitempty" xml:"IdJsonList,omitempty"`
}

func (s BatchRegistAnonymousTbAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRegistAnonymousTbAccountRequest) GoString() string {
	return s.String()
}

func (s *BatchRegistAnonymousTbAccountRequest) SetBizId(v string) *BatchRegistAnonymousTbAccountRequest {
	s.BizId = &v
	return s
}

func (s *BatchRegistAnonymousTbAccountRequest) SetIdJsonList(v string) *BatchRegistAnonymousTbAccountRequest {
	s.IdJsonList = &v
	return s
}

type BatchRegistAnonymousTbAccountResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BatchId   *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
}

func (s BatchRegistAnonymousTbAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRegistAnonymousTbAccountResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRegistAnonymousTbAccountResponseBody) SetCode(v string) *BatchRegistAnonymousTbAccountResponseBody {
	s.Code = &v
	return s
}

func (s *BatchRegistAnonymousTbAccountResponseBody) SetMessage(v string) *BatchRegistAnonymousTbAccountResponseBody {
	s.Message = &v
	return s
}

func (s *BatchRegistAnonymousTbAccountResponseBody) SetRequestId(v string) *BatchRegistAnonymousTbAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchRegistAnonymousTbAccountResponseBody) SetBatchId(v string) *BatchRegistAnonymousTbAccountResponseBody {
	s.BatchId = &v
	return s
}

type BatchRegistAnonymousTbAccountResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchRegistAnonymousTbAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchRegistAnonymousTbAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRegistAnonymousTbAccountResponse) GoString() string {
	return s.String()
}

func (s *BatchRegistAnonymousTbAccountResponse) SetHeaders(v map[string]*string) *BatchRegistAnonymousTbAccountResponse {
	s.Headers = v
	return s
}

func (s *BatchRegistAnonymousTbAccountResponse) SetBody(v *BatchRegistAnonymousTbAccountResponseBody) *BatchRegistAnonymousTbAccountResponse {
	s.Body = v
	return s
}

type CancelOrderRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	LmOrderId             *string `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
}

func (s CancelOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelOrderRequest) SetBizId(v string) *CancelOrderRequest {
	s.BizId = &v
	return s
}

func (s *CancelOrderRequest) SetBizUid(v string) *CancelOrderRequest {
	s.BizUid = &v
	return s
}

func (s *CancelOrderRequest) SetLmOrderId(v string) *CancelOrderRequest {
	s.LmOrderId = &v
	return s
}

func (s *CancelOrderRequest) SetAccountType(v string) *CancelOrderRequest {
	s.AccountType = &v
	return s
}

func (s *CancelOrderRequest) SetThirdPartyUserId(v string) *CancelOrderRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *CancelOrderRequest) SetUseAnonymousTbAccount(v bool) *CancelOrderRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

type CancelOrderResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	LogsId     *string `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
}

func (s CancelOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrderResponseBody) SetRequestId(v string) *CancelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelOrderResponseBody) SetSuccess(v bool) *CancelOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CancelOrderResponseBody) SetSubMessage(v string) *CancelOrderResponseBody {
	s.SubMessage = &v
	return s
}

func (s *CancelOrderResponseBody) SetCode(v string) *CancelOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CancelOrderResponseBody) SetMessage(v string) *CancelOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CancelOrderResponseBody) SetSubCode(v string) *CancelOrderResponseBody {
	s.SubCode = &v
	return s
}

func (s *CancelOrderResponseBody) SetLogsId(v string) *CancelOrderResponseBody {
	s.LogsId = &v
	return s
}

type CancelOrderResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelOrderResponse) SetHeaders(v map[string]*string) *CancelOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelOrderResponse) SetBody(v *CancelOrderResponseBody) *CancelOrderResponse {
	s.Body = v
	return s
}

type CancelRefundRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	SubLmOrderId          *string `json:"SubLmOrderId,omitempty" xml:"SubLmOrderId,omitempty"`
	DisputeId             *int64  `json:"DisputeId,omitempty" xml:"DisputeId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s CancelRefundRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelRefundRequest) GoString() string {
	return s.String()
}

func (s *CancelRefundRequest) SetBizId(v string) *CancelRefundRequest {
	s.BizId = &v
	return s
}

func (s *CancelRefundRequest) SetBizUid(v string) *CancelRefundRequest {
	s.BizUid = &v
	return s
}

func (s *CancelRefundRequest) SetSubLmOrderId(v string) *CancelRefundRequest {
	s.SubLmOrderId = &v
	return s
}

func (s *CancelRefundRequest) SetDisputeId(v int64) *CancelRefundRequest {
	s.DisputeId = &v
	return s
}

func (s *CancelRefundRequest) SetUseAnonymousTbAccount(v bool) *CancelRefundRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *CancelRefundRequest) SetThirdPartyUserId(v string) *CancelRefundRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *CancelRefundRequest) SetAccountType(v string) *CancelRefundRequest {
	s.AccountType = &v
	return s
}

type CancelRefundResponseBody struct {
	Code                  *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message               *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId             *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RefundApplicationData *CancelRefundResponseBodyRefundApplicationData `json:"RefundApplicationData,omitempty" xml:"RefundApplicationData,omitempty" type:"Struct"`
}

func (s CancelRefundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelRefundResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRefundResponseBody) SetCode(v string) *CancelRefundResponseBody {
	s.Code = &v
	return s
}

func (s *CancelRefundResponseBody) SetMessage(v string) *CancelRefundResponseBody {
	s.Message = &v
	return s
}

func (s *CancelRefundResponseBody) SetRequestId(v string) *CancelRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelRefundResponseBody) SetRefundApplicationData(v *CancelRefundResponseBodyRefundApplicationData) *CancelRefundResponseBody {
	s.RefundApplicationData = v
	return s
}

type CancelRefundResponseBodyRefundApplicationData struct {
	DisputeType   *int32  `json:"DisputeType,omitempty" xml:"DisputeType,omitempty"`
	SubLmOrderId  *string `json:"SubLmOrderId,omitempty" xml:"SubLmOrderId,omitempty"`
	DisputeStatus *int32  `json:"DisputeStatus,omitempty" xml:"DisputeStatus,omitempty"`
}

func (s CancelRefundResponseBodyRefundApplicationData) String() string {
	return tea.Prettify(s)
}

func (s CancelRefundResponseBodyRefundApplicationData) GoString() string {
	return s.String()
}

func (s *CancelRefundResponseBodyRefundApplicationData) SetDisputeType(v int32) *CancelRefundResponseBodyRefundApplicationData {
	s.DisputeType = &v
	return s
}

func (s *CancelRefundResponseBodyRefundApplicationData) SetSubLmOrderId(v string) *CancelRefundResponseBodyRefundApplicationData {
	s.SubLmOrderId = &v
	return s
}

func (s *CancelRefundResponseBodyRefundApplicationData) SetDisputeStatus(v int32) *CancelRefundResponseBodyRefundApplicationData {
	s.DisputeStatus = &v
	return s
}

type CancelRefundResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelRefundResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelRefundResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelRefundResponse) GoString() string {
	return s.String()
}

func (s *CancelRefundResponse) SetHeaders(v map[string]*string) *CancelRefundResponse {
	s.Headers = v
	return s
}

func (s *CancelRefundResponse) SetBody(v *CancelRefundResponseBody) *CancelRefundResponse {
	s.Body = v
	return s
}

type ConfirmDisburseRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	LmOrderId             *string `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s ConfirmDisburseRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDisburseRequest) GoString() string {
	return s.String()
}

func (s *ConfirmDisburseRequest) SetBizId(v string) *ConfirmDisburseRequest {
	s.BizId = &v
	return s
}

func (s *ConfirmDisburseRequest) SetBizUid(v string) *ConfirmDisburseRequest {
	s.BizUid = &v
	return s
}

func (s *ConfirmDisburseRequest) SetLmOrderId(v string) *ConfirmDisburseRequest {
	s.LmOrderId = &v
	return s
}

func (s *ConfirmDisburseRequest) SetUseAnonymousTbAccount(v bool) *ConfirmDisburseRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *ConfirmDisburseRequest) SetThirdPartyUserId(v string) *ConfirmDisburseRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *ConfirmDisburseRequest) SetAccountType(v string) *ConfirmDisburseRequest {
	s.AccountType = &v
	return s
}

type ConfirmDisburseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfirmDisburseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDisburseResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmDisburseResponseBody) SetCode(v string) *ConfirmDisburseResponseBody {
	s.Code = &v
	return s
}

func (s *ConfirmDisburseResponseBody) SetMessage(v string) *ConfirmDisburseResponseBody {
	s.Message = &v
	return s
}

func (s *ConfirmDisburseResponseBody) SetRequestId(v string) *ConfirmDisburseResponseBody {
	s.RequestId = &v
	return s
}

type ConfirmDisburseResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfirmDisburseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConfirmDisburseResponse) SetBody(v *ConfirmDisburseResponseBody) *ConfirmDisburseResponse {
	s.Body = v
	return s
}

type CreateMovieTicketOrderRequest struct {
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid         *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	OutTradeId     *string `json:"OutTradeId,omitempty" xml:"OutTradeId,omitempty"`
	LockSeatAppKey *string `json:"LockSeatAppKey,omitempty" xml:"LockSeatAppKey,omitempty"`
	ExtJson        *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s CreateMovieTicketOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMovieTicketOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateMovieTicketOrderRequest) SetBizId(v string) *CreateMovieTicketOrderRequest {
	s.BizId = &v
	return s
}

func (s *CreateMovieTicketOrderRequest) SetBizUid(v string) *CreateMovieTicketOrderRequest {
	s.BizUid = &v
	return s
}

func (s *CreateMovieTicketOrderRequest) SetOutTradeId(v string) *CreateMovieTicketOrderRequest {
	s.OutTradeId = &v
	return s
}

func (s *CreateMovieTicketOrderRequest) SetLockSeatAppKey(v string) *CreateMovieTicketOrderRequest {
	s.LockSeatAppKey = &v
	return s
}

func (s *CreateMovieTicketOrderRequest) SetExtJson(v string) *CreateMovieTicketOrderRequest {
	s.ExtJson = &v
	return s
}

type CreateMovieTicketOrderResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                                  `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                                  `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                                  `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Model      *CreateMovieTicketOrderResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s CreateMovieTicketOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMovieTicketOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMovieTicketOrderResponseBody) SetRequestId(v string) *CreateMovieTicketOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMovieTicketOrderResponseBody) SetSuccess(v bool) *CreateMovieTicketOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMovieTicketOrderResponseBody) SetSubMessage(v string) *CreateMovieTicketOrderResponseBody {
	s.SubMessage = &v
	return s
}

func (s *CreateMovieTicketOrderResponseBody) SetCode(v string) *CreateMovieTicketOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMovieTicketOrderResponseBody) SetMessage(v string) *CreateMovieTicketOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMovieTicketOrderResponseBody) SetSubCode(v string) *CreateMovieTicketOrderResponseBody {
	s.SubCode = &v
	return s
}

func (s *CreateMovieTicketOrderResponseBody) SetTotalCount(v int64) *CreateMovieTicketOrderResponseBody {
	s.TotalCount = &v
	return s
}

func (s *CreateMovieTicketOrderResponseBody) SetLogsId(v string) *CreateMovieTicketOrderResponseBody {
	s.LogsId = &v
	return s
}

func (s *CreateMovieTicketOrderResponseBody) SetModel(v *CreateMovieTicketOrderResponseBodyModel) *CreateMovieTicketOrderResponseBody {
	s.Model = v
	return s
}

type CreateMovieTicketOrderResponseBodyModel struct {
	RedirectUrl *string                                             `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	PayTradeIds *CreateMovieTicketOrderResponseBodyModelPayTradeIds `json:"PayTradeIds,omitempty" xml:"PayTradeIds,omitempty" type:"Struct"`
	OrderIds    *CreateMovieTicketOrderResponseBodyModelOrderIds    `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Struct"`
}

func (s CreateMovieTicketOrderResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s CreateMovieTicketOrderResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CreateMovieTicketOrderResponseBodyModel) SetRedirectUrl(v string) *CreateMovieTicketOrderResponseBodyModel {
	s.RedirectUrl = &v
	return s
}

func (s *CreateMovieTicketOrderResponseBodyModel) SetPayTradeIds(v *CreateMovieTicketOrderResponseBodyModelPayTradeIds) *CreateMovieTicketOrderResponseBodyModel {
	s.PayTradeIds = v
	return s
}

func (s *CreateMovieTicketOrderResponseBodyModel) SetOrderIds(v *CreateMovieTicketOrderResponseBodyModelOrderIds) *CreateMovieTicketOrderResponseBodyModel {
	s.OrderIds = v
	return s
}

type CreateMovieTicketOrderResponseBodyModelPayTradeIds struct {
	PayTradeIds []*string `json:"PayTradeIds,omitempty" xml:"PayTradeIds,omitempty" type:"Repeated"`
}

func (s CreateMovieTicketOrderResponseBodyModelPayTradeIds) String() string {
	return tea.Prettify(s)
}

func (s CreateMovieTicketOrderResponseBodyModelPayTradeIds) GoString() string {
	return s.String()
}

func (s *CreateMovieTicketOrderResponseBodyModelPayTradeIds) SetPayTradeIds(v []*string) *CreateMovieTicketOrderResponseBodyModelPayTradeIds {
	s.PayTradeIds = v
	return s
}

type CreateMovieTicketOrderResponseBodyModelOrderIds struct {
	OrderIds []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s CreateMovieTicketOrderResponseBodyModelOrderIds) String() string {
	return tea.Prettify(s)
}

func (s CreateMovieTicketOrderResponseBodyModelOrderIds) GoString() string {
	return s.String()
}

func (s *CreateMovieTicketOrderResponseBodyModelOrderIds) SetOrderIds(v []*string) *CreateMovieTicketOrderResponseBodyModelOrderIds {
	s.OrderIds = v
	return s
}

type CreateMovieTicketOrderResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMovieTicketOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMovieTicketOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMovieTicketOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateMovieTicketOrderResponse) SetHeaders(v map[string]*string) *CreateMovieTicketOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateMovieTicketOrderResponse) SetBody(v *CreateMovieTicketOrderResponseBody) *CreateMovieTicketOrderResponse {
	s.Body = v
	return s
}

type CreateOrderRequest struct {
	BizId                 *string                       `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string                       `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	OutTradeId            *string                       `json:"OutTradeId,omitempty" xml:"OutTradeId,omitempty"`
	ItemId                *int64                        `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Quantity              *int64                        `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	TotalAmount           *int64                        `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	ExtJson               *string                       `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	DeliveryAddress       *string                       `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	OrderExpireTime       *int64                        `json:"OrderExpireTime,omitempty" xml:"OrderExpireTime,omitempty"`
	UseAnonymousTbAccount *bool                         `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string                       `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string                       `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	LmItemId              *string                       `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	BuyerMessageMap       *string                       `json:"BuyerMessageMap,omitempty" xml:"BuyerMessageMap,omitempty"`
	ItemList              []*CreateOrderRequestItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
}

func (s CreateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) SetBizId(v string) *CreateOrderRequest {
	s.BizId = &v
	return s
}

func (s *CreateOrderRequest) SetBizUid(v string) *CreateOrderRequest {
	s.BizUid = &v
	return s
}

func (s *CreateOrderRequest) SetOutTradeId(v string) *CreateOrderRequest {
	s.OutTradeId = &v
	return s
}

func (s *CreateOrderRequest) SetItemId(v int64) *CreateOrderRequest {
	s.ItemId = &v
	return s
}

func (s *CreateOrderRequest) SetQuantity(v int64) *CreateOrderRequest {
	s.Quantity = &v
	return s
}

func (s *CreateOrderRequest) SetTotalAmount(v int64) *CreateOrderRequest {
	s.TotalAmount = &v
	return s
}

func (s *CreateOrderRequest) SetExtJson(v string) *CreateOrderRequest {
	s.ExtJson = &v
	return s
}

func (s *CreateOrderRequest) SetDeliveryAddress(v string) *CreateOrderRequest {
	s.DeliveryAddress = &v
	return s
}

func (s *CreateOrderRequest) SetOrderExpireTime(v int64) *CreateOrderRequest {
	s.OrderExpireTime = &v
	return s
}

func (s *CreateOrderRequest) SetUseAnonymousTbAccount(v bool) *CreateOrderRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *CreateOrderRequest) SetThirdPartyUserId(v string) *CreateOrderRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *CreateOrderRequest) SetAccountType(v string) *CreateOrderRequest {
	s.AccountType = &v
	return s
}

func (s *CreateOrderRequest) SetLmItemId(v string) *CreateOrderRequest {
	s.LmItemId = &v
	return s
}

func (s *CreateOrderRequest) SetBuyerMessageMap(v string) *CreateOrderRequest {
	s.BuyerMessageMap = &v
	return s
}

func (s *CreateOrderRequest) SetItemList(v []*CreateOrderRequestItemList) *CreateOrderRequest {
	s.ItemList = v
	return s
}

type CreateOrderRequestItemList struct {
	SkuId    *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	LmItemId *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId   *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Quantity *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s CreateOrderRequestItemList) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderRequestItemList) GoString() string {
	return s.String()
}

func (s *CreateOrderRequestItemList) SetSkuId(v int64) *CreateOrderRequestItemList {
	s.SkuId = &v
	return s
}

func (s *CreateOrderRequestItemList) SetLmItemId(v string) *CreateOrderRequestItemList {
	s.LmItemId = &v
	return s
}

func (s *CreateOrderRequestItemList) SetItemId(v int64) *CreateOrderRequestItemList {
	s.ItemId = &v
	return s
}

func (s *CreateOrderRequestItemList) SetQuantity(v int32) *CreateOrderRequestItemList {
	s.Quantity = &v
	return s
}

type CreateOrderResponseBody struct {
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                       `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                       `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                       `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Model      *CreateOrderResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s CreateOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) SetRequestId(v string) *CreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponseBody) SetSuccess(v bool) *CreateOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOrderResponseBody) SetSubMessage(v string) *CreateOrderResponseBody {
	s.SubMessage = &v
	return s
}

func (s *CreateOrderResponseBody) SetCode(v string) *CreateOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOrderResponseBody) SetMessage(v string) *CreateOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOrderResponseBody) SetSubCode(v string) *CreateOrderResponseBody {
	s.SubCode = &v
	return s
}

func (s *CreateOrderResponseBody) SetTotalCount(v int64) *CreateOrderResponseBody {
	s.TotalCount = &v
	return s
}

func (s *CreateOrderResponseBody) SetLogsId(v string) *CreateOrderResponseBody {
	s.LogsId = &v
	return s
}

func (s *CreateOrderResponseBody) SetModel(v *CreateOrderResponseBodyModel) *CreateOrderResponseBody {
	s.Model = v
	return s
}

type CreateOrderResponseBodyModel struct {
	RedirectUrl *string                                  `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	LmOrderList *CreateOrderResponseBodyModelLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Struct"`
	PayTradeIds *CreateOrderResponseBodyModelPayTradeIds `json:"PayTradeIds,omitempty" xml:"PayTradeIds,omitempty" type:"Struct"`
	OrderIds    *CreateOrderResponseBodyModelOrderIds    `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Struct"`
}

func (s CreateOrderResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyModel) SetRedirectUrl(v string) *CreateOrderResponseBodyModel {
	s.RedirectUrl = &v
	return s
}

func (s *CreateOrderResponseBodyModel) SetLmOrderList(v *CreateOrderResponseBodyModelLmOrderList) *CreateOrderResponseBodyModel {
	s.LmOrderList = v
	return s
}

func (s *CreateOrderResponseBodyModel) SetPayTradeIds(v *CreateOrderResponseBodyModelPayTradeIds) *CreateOrderResponseBodyModel {
	s.PayTradeIds = v
	return s
}

func (s *CreateOrderResponseBodyModel) SetOrderIds(v *CreateOrderResponseBodyModelOrderIds) *CreateOrderResponseBodyModel {
	s.OrderIds = v
	return s
}

type CreateOrderResponseBodyModelLmOrderList struct {
	LmOrderList []*CreateOrderResponseBodyModelLmOrderListLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Repeated"`
}

func (s CreateOrderResponseBodyModelLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBodyModelLmOrderList) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyModelLmOrderList) SetLmOrderList(v []*CreateOrderResponseBodyModelLmOrderListLmOrderList) *CreateOrderResponseBodyModelLmOrderList {
	s.LmOrderList = v
	return s
}

type CreateOrderResponseBodyModelLmOrderListLmOrderList struct {
	LmOrderId *string `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
}

func (s CreateOrderResponseBodyModelLmOrderListLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBodyModelLmOrderListLmOrderList) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyModelLmOrderListLmOrderList) SetLmOrderId(v string) *CreateOrderResponseBodyModelLmOrderListLmOrderList {
	s.LmOrderId = &v
	return s
}

type CreateOrderResponseBodyModelPayTradeIds struct {
	PayTradeIds []*string `json:"PayTradeIds,omitempty" xml:"PayTradeIds,omitempty" type:"Repeated"`
}

func (s CreateOrderResponseBodyModelPayTradeIds) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBodyModelPayTradeIds) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyModelPayTradeIds) SetPayTradeIds(v []*string) *CreateOrderResponseBodyModelPayTradeIds {
	s.PayTradeIds = v
	return s
}

type CreateOrderResponseBodyModelOrderIds struct {
	OrderIds []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s CreateOrderResponseBodyModelOrderIds) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBodyModelOrderIds) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyModelOrderIds) SetOrderIds(v []*string) *CreateOrderResponseBodyModelOrderIds {
	s.OrderIds = v
	return s
}

type CreateOrderResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderResponse) SetHeaders(v map[string]*string) *CreateOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderResponse) SetBody(v *CreateOrderResponseBody) *CreateOrderResponse {
	s.Body = v
	return s
}

type CreateOrderV2Request struct {
	BizId                 *string                         `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string                         `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	OutTradeId            *string                         `json:"OutTradeId,omitempty" xml:"OutTradeId,omitempty"`
	ItemId                *int64                          `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Quantity              *int64                          `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	TotalAmount           *int64                          `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	ExtJson               *string                         `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	DeliveryAddress       *string                         `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	OrderExpireTime       *int64                          `json:"OrderExpireTime,omitempty" xml:"OrderExpireTime,omitempty"`
	UseAnonymousTbAccount *bool                           `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string                         `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string                         `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	LmItemId              *string                         `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	BuyerMessageMap       *string                         `json:"BuyerMessageMap,omitempty" xml:"BuyerMessageMap,omitempty"`
	ItemList              []*CreateOrderV2RequestItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
}

func (s CreateOrderV2Request) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderV2Request) GoString() string {
	return s.String()
}

func (s *CreateOrderV2Request) SetBizId(v string) *CreateOrderV2Request {
	s.BizId = &v
	return s
}

func (s *CreateOrderV2Request) SetBizUid(v string) *CreateOrderV2Request {
	s.BizUid = &v
	return s
}

func (s *CreateOrderV2Request) SetOutTradeId(v string) *CreateOrderV2Request {
	s.OutTradeId = &v
	return s
}

func (s *CreateOrderV2Request) SetItemId(v int64) *CreateOrderV2Request {
	s.ItemId = &v
	return s
}

func (s *CreateOrderV2Request) SetQuantity(v int64) *CreateOrderV2Request {
	s.Quantity = &v
	return s
}

func (s *CreateOrderV2Request) SetTotalAmount(v int64) *CreateOrderV2Request {
	s.TotalAmount = &v
	return s
}

func (s *CreateOrderV2Request) SetExtJson(v string) *CreateOrderV2Request {
	s.ExtJson = &v
	return s
}

func (s *CreateOrderV2Request) SetDeliveryAddress(v string) *CreateOrderV2Request {
	s.DeliveryAddress = &v
	return s
}

func (s *CreateOrderV2Request) SetOrderExpireTime(v int64) *CreateOrderV2Request {
	s.OrderExpireTime = &v
	return s
}

func (s *CreateOrderV2Request) SetUseAnonymousTbAccount(v bool) *CreateOrderV2Request {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *CreateOrderV2Request) SetThirdPartyUserId(v string) *CreateOrderV2Request {
	s.ThirdPartyUserId = &v
	return s
}

func (s *CreateOrderV2Request) SetAccountType(v string) *CreateOrderV2Request {
	s.AccountType = &v
	return s
}

func (s *CreateOrderV2Request) SetLmItemId(v string) *CreateOrderV2Request {
	s.LmItemId = &v
	return s
}

func (s *CreateOrderV2Request) SetBuyerMessageMap(v string) *CreateOrderV2Request {
	s.BuyerMessageMap = &v
	return s
}

func (s *CreateOrderV2Request) SetItemList(v []*CreateOrderV2RequestItemList) *CreateOrderV2Request {
	s.ItemList = v
	return s
}

type CreateOrderV2RequestItemList struct {
	SkuId    *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	LmItemId *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId   *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Quantity *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s CreateOrderV2RequestItemList) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderV2RequestItemList) GoString() string {
	return s.String()
}

func (s *CreateOrderV2RequestItemList) SetSkuId(v int64) *CreateOrderV2RequestItemList {
	s.SkuId = &v
	return s
}

func (s *CreateOrderV2RequestItemList) SetLmItemId(v string) *CreateOrderV2RequestItemList {
	s.LmItemId = &v
	return s
}

func (s *CreateOrderV2RequestItemList) SetItemId(v int64) *CreateOrderV2RequestItemList {
	s.ItemId = &v
	return s
}

func (s *CreateOrderV2RequestItemList) SetQuantity(v int32) *CreateOrderV2RequestItemList {
	s.Quantity = &v
	return s
}

type CreateOrderV2ResponseBody struct {
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                         `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                         `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                         `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Model      *CreateOrderV2ResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s CreateOrderV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderV2ResponseBody) SetRequestId(v string) *CreateOrderV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderV2ResponseBody) SetSuccess(v bool) *CreateOrderV2ResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOrderV2ResponseBody) SetSubMessage(v string) *CreateOrderV2ResponseBody {
	s.SubMessage = &v
	return s
}

func (s *CreateOrderV2ResponseBody) SetCode(v string) *CreateOrderV2ResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOrderV2ResponseBody) SetMessage(v string) *CreateOrderV2ResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOrderV2ResponseBody) SetSubCode(v string) *CreateOrderV2ResponseBody {
	s.SubCode = &v
	return s
}

func (s *CreateOrderV2ResponseBody) SetTotalCount(v int64) *CreateOrderV2ResponseBody {
	s.TotalCount = &v
	return s
}

func (s *CreateOrderV2ResponseBody) SetLogsId(v string) *CreateOrderV2ResponseBody {
	s.LogsId = &v
	return s
}

func (s *CreateOrderV2ResponseBody) SetModel(v *CreateOrderV2ResponseBodyModel) *CreateOrderV2ResponseBody {
	s.Model = v
	return s
}

type CreateOrderV2ResponseBodyModel struct {
	RedirectUrl *string                                    `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	LmOrderList *CreateOrderV2ResponseBodyModelLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Struct"`
	PayTradeIds *CreateOrderV2ResponseBodyModelPayTradeIds `json:"PayTradeIds,omitempty" xml:"PayTradeIds,omitempty" type:"Struct"`
	OrderIds    *CreateOrderV2ResponseBodyModelOrderIds    `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Struct"`
}

func (s CreateOrderV2ResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderV2ResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CreateOrderV2ResponseBodyModel) SetRedirectUrl(v string) *CreateOrderV2ResponseBodyModel {
	s.RedirectUrl = &v
	return s
}

func (s *CreateOrderV2ResponseBodyModel) SetLmOrderList(v *CreateOrderV2ResponseBodyModelLmOrderList) *CreateOrderV2ResponseBodyModel {
	s.LmOrderList = v
	return s
}

func (s *CreateOrderV2ResponseBodyModel) SetPayTradeIds(v *CreateOrderV2ResponseBodyModelPayTradeIds) *CreateOrderV2ResponseBodyModel {
	s.PayTradeIds = v
	return s
}

func (s *CreateOrderV2ResponseBodyModel) SetOrderIds(v *CreateOrderV2ResponseBodyModelOrderIds) *CreateOrderV2ResponseBodyModel {
	s.OrderIds = v
	return s
}

type CreateOrderV2ResponseBodyModelLmOrderList struct {
	LmOrderList []*CreateOrderV2ResponseBodyModelLmOrderListLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Repeated"`
}

func (s CreateOrderV2ResponseBodyModelLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderV2ResponseBodyModelLmOrderList) GoString() string {
	return s.String()
}

func (s *CreateOrderV2ResponseBodyModelLmOrderList) SetLmOrderList(v []*CreateOrderV2ResponseBodyModelLmOrderListLmOrderList) *CreateOrderV2ResponseBodyModelLmOrderList {
	s.LmOrderList = v
	return s
}

type CreateOrderV2ResponseBodyModelLmOrderListLmOrderList struct {
	LmOrderId *string `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
}

func (s CreateOrderV2ResponseBodyModelLmOrderListLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderV2ResponseBodyModelLmOrderListLmOrderList) GoString() string {
	return s.String()
}

func (s *CreateOrderV2ResponseBodyModelLmOrderListLmOrderList) SetLmOrderId(v string) *CreateOrderV2ResponseBodyModelLmOrderListLmOrderList {
	s.LmOrderId = &v
	return s
}

type CreateOrderV2ResponseBodyModelPayTradeIds struct {
	PayTradeIds []*string `json:"PayTradeIds,omitempty" xml:"PayTradeIds,omitempty" type:"Repeated"`
}

func (s CreateOrderV2ResponseBodyModelPayTradeIds) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderV2ResponseBodyModelPayTradeIds) GoString() string {
	return s.String()
}

func (s *CreateOrderV2ResponseBodyModelPayTradeIds) SetPayTradeIds(v []*string) *CreateOrderV2ResponseBodyModelPayTradeIds {
	s.PayTradeIds = v
	return s
}

type CreateOrderV2ResponseBodyModelOrderIds struct {
	OrderIds []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s CreateOrderV2ResponseBodyModelOrderIds) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderV2ResponseBodyModelOrderIds) GoString() string {
	return s.String()
}

func (s *CreateOrderV2ResponseBodyModelOrderIds) SetOrderIds(v []*string) *CreateOrderV2ResponseBodyModelOrderIds {
	s.OrderIds = v
	return s
}

type CreateOrderV2Response struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOrderV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrderV2Response) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderV2Response) GoString() string {
	return s.String()
}

func (s *CreateOrderV2Response) SetHeaders(v map[string]*string) *CreateOrderV2Response {
	s.Headers = v
	return s
}

func (s *CreateOrderV2Response) SetBody(v *CreateOrderV2ResponseBody) *CreateOrderV2Response {
	s.Body = v
	return s
}

type CreatePayUrlRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	BuyInfo               *string `json:"BuyInfo,omitempty" xml:"BuyInfo,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
}

func (s CreatePayUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePayUrlRequest) GoString() string {
	return s.String()
}

func (s *CreatePayUrlRequest) SetBizId(v string) *CreatePayUrlRequest {
	s.BizId = &v
	return s
}

func (s *CreatePayUrlRequest) SetBizUid(v string) *CreatePayUrlRequest {
	s.BizUid = &v
	return s
}

func (s *CreatePayUrlRequest) SetBuyInfo(v string) *CreatePayUrlRequest {
	s.BuyInfo = &v
	return s
}

func (s *CreatePayUrlRequest) SetUseAnonymousTbAccount(v bool) *CreatePayUrlRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *CreatePayUrlRequest) SetThirdPartyUserId(v string) *CreatePayUrlRequest {
	s.ThirdPartyUserId = &v
	return s
}

type CreatePayUrlResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Model     *CreatePayUrlResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s CreatePayUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePayUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePayUrlResponseBody) SetCode(v string) *CreatePayUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePayUrlResponseBody) SetMessage(v string) *CreatePayUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePayUrlResponseBody) SetRequestId(v string) *CreatePayUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePayUrlResponseBody) SetSuccess(v bool) *CreatePayUrlResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePayUrlResponseBody) SetModel(v *CreatePayUrlResponseBodyModel) *CreatePayUrlResponseBody {
	s.Model = v
	return s
}

type CreatePayUrlResponseBodyModel struct {
	RedirectUrl *string   `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	PayTradeIds []*string `json:"PayTradeIds,omitempty" xml:"PayTradeIds,omitempty" type:"Repeated"`
	LmOrderList []*string `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Repeated"`
	OrderIds    []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s CreatePayUrlResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s CreatePayUrlResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CreatePayUrlResponseBodyModel) SetRedirectUrl(v string) *CreatePayUrlResponseBodyModel {
	s.RedirectUrl = &v
	return s
}

func (s *CreatePayUrlResponseBodyModel) SetPayTradeIds(v []*string) *CreatePayUrlResponseBodyModel {
	s.PayTradeIds = v
	return s
}

func (s *CreatePayUrlResponseBodyModel) SetLmOrderList(v []*string) *CreatePayUrlResponseBodyModel {
	s.LmOrderList = v
	return s
}

func (s *CreatePayUrlResponseBodyModel) SetOrderIds(v []*string) *CreatePayUrlResponseBodyModel {
	s.OrderIds = v
	return s
}

type CreatePayUrlResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePayUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePayUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePayUrlResponse) GoString() string {
	return s.String()
}

func (s *CreatePayUrlResponse) SetHeaders(v map[string]*string) *CreatePayUrlResponse {
	s.Headers = v
	return s
}

func (s *CreatePayUrlResponse) SetBody(v *CreatePayUrlResponseBody) *CreatePayUrlResponse {
	s.Body = v
	return s
}

type CreateVirtualProductOrderRequest struct {
	BizId                 *string                                     `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string                                     `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	OutTradeId            *string                                     `json:"OutTradeId,omitempty" xml:"OutTradeId,omitempty"`
	ItemId                *int64                                      `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Quantity              *int64                                      `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	TotalAmount           *int64                                      `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	ExtJson               *string                                     `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	DeliveryAddress       *string                                     `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	OrderExpireTime       *int64                                      `json:"OrderExpireTime,omitempty" xml:"OrderExpireTime,omitempty"`
	UseAnonymousTbAccount *bool                                       `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string                                     `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string                                     `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	LmItemId              *string                                     `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemList              []*CreateVirtualProductOrderRequestItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
}

func (s CreateVirtualProductOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualProductOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualProductOrderRequest) SetBizId(v string) *CreateVirtualProductOrderRequest {
	s.BizId = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetBizUid(v string) *CreateVirtualProductOrderRequest {
	s.BizUid = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetOutTradeId(v string) *CreateVirtualProductOrderRequest {
	s.OutTradeId = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetItemId(v int64) *CreateVirtualProductOrderRequest {
	s.ItemId = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetQuantity(v int64) *CreateVirtualProductOrderRequest {
	s.Quantity = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetTotalAmount(v int64) *CreateVirtualProductOrderRequest {
	s.TotalAmount = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetExtJson(v string) *CreateVirtualProductOrderRequest {
	s.ExtJson = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetDeliveryAddress(v string) *CreateVirtualProductOrderRequest {
	s.DeliveryAddress = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetOrderExpireTime(v int64) *CreateVirtualProductOrderRequest {
	s.OrderExpireTime = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetUseAnonymousTbAccount(v bool) *CreateVirtualProductOrderRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetThirdPartyUserId(v string) *CreateVirtualProductOrderRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetAccountType(v string) *CreateVirtualProductOrderRequest {
	s.AccountType = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetLmItemId(v string) *CreateVirtualProductOrderRequest {
	s.LmItemId = &v
	return s
}

func (s *CreateVirtualProductOrderRequest) SetItemList(v []*CreateVirtualProductOrderRequestItemList) *CreateVirtualProductOrderRequest {
	s.ItemList = v
	return s
}

type CreateVirtualProductOrderRequestItemList struct {
	SkuId    *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	LmItemId *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId   *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Quantity *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s CreateVirtualProductOrderRequestItemList) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualProductOrderRequestItemList) GoString() string {
	return s.String()
}

func (s *CreateVirtualProductOrderRequestItemList) SetSkuId(v int64) *CreateVirtualProductOrderRequestItemList {
	s.SkuId = &v
	return s
}

func (s *CreateVirtualProductOrderRequestItemList) SetLmItemId(v string) *CreateVirtualProductOrderRequestItemList {
	s.LmItemId = &v
	return s
}

func (s *CreateVirtualProductOrderRequestItemList) SetItemId(v int64) *CreateVirtualProductOrderRequestItemList {
	s.ItemId = &v
	return s
}

func (s *CreateVirtualProductOrderRequestItemList) SetQuantity(v int32) *CreateVirtualProductOrderRequestItemList {
	s.Quantity = &v
	return s
}

type CreateVirtualProductOrderResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                                     `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                                     `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                                     `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Model      *CreateVirtualProductOrderResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s CreateVirtualProductOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualProductOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualProductOrderResponseBody) SetRequestId(v string) *CreateVirtualProductOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualProductOrderResponseBody) SetSuccess(v bool) *CreateVirtualProductOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVirtualProductOrderResponseBody) SetSubMessage(v string) *CreateVirtualProductOrderResponseBody {
	s.SubMessage = &v
	return s
}

func (s *CreateVirtualProductOrderResponseBody) SetCode(v string) *CreateVirtualProductOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVirtualProductOrderResponseBody) SetMessage(v string) *CreateVirtualProductOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVirtualProductOrderResponseBody) SetSubCode(v string) *CreateVirtualProductOrderResponseBody {
	s.SubCode = &v
	return s
}

func (s *CreateVirtualProductOrderResponseBody) SetTotalCount(v int64) *CreateVirtualProductOrderResponseBody {
	s.TotalCount = &v
	return s
}

func (s *CreateVirtualProductOrderResponseBody) SetLogsId(v string) *CreateVirtualProductOrderResponseBody {
	s.LogsId = &v
	return s
}

func (s *CreateVirtualProductOrderResponseBody) SetModel(v *CreateVirtualProductOrderResponseBodyModel) *CreateVirtualProductOrderResponseBody {
	s.Model = v
	return s
}

type CreateVirtualProductOrderResponseBodyModel struct {
	RedirectUrl *string                                                `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	LmOrderList *CreateVirtualProductOrderResponseBodyModelLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Struct"`
	PayTradeIds *CreateVirtualProductOrderResponseBodyModelPayTradeIds `json:"PayTradeIds,omitempty" xml:"PayTradeIds,omitempty" type:"Struct"`
	OrderIds    *CreateVirtualProductOrderResponseBodyModelOrderIds    `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Struct"`
}

func (s CreateVirtualProductOrderResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualProductOrderResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CreateVirtualProductOrderResponseBodyModel) SetRedirectUrl(v string) *CreateVirtualProductOrderResponseBodyModel {
	s.RedirectUrl = &v
	return s
}

func (s *CreateVirtualProductOrderResponseBodyModel) SetLmOrderList(v *CreateVirtualProductOrderResponseBodyModelLmOrderList) *CreateVirtualProductOrderResponseBodyModel {
	s.LmOrderList = v
	return s
}

func (s *CreateVirtualProductOrderResponseBodyModel) SetPayTradeIds(v *CreateVirtualProductOrderResponseBodyModelPayTradeIds) *CreateVirtualProductOrderResponseBodyModel {
	s.PayTradeIds = v
	return s
}

func (s *CreateVirtualProductOrderResponseBodyModel) SetOrderIds(v *CreateVirtualProductOrderResponseBodyModelOrderIds) *CreateVirtualProductOrderResponseBodyModel {
	s.OrderIds = v
	return s
}

type CreateVirtualProductOrderResponseBodyModelLmOrderList struct {
	LmOrderList []*CreateVirtualProductOrderResponseBodyModelLmOrderListLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Repeated"`
}

func (s CreateVirtualProductOrderResponseBodyModelLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualProductOrderResponseBodyModelLmOrderList) GoString() string {
	return s.String()
}

func (s *CreateVirtualProductOrderResponseBodyModelLmOrderList) SetLmOrderList(v []*CreateVirtualProductOrderResponseBodyModelLmOrderListLmOrderList) *CreateVirtualProductOrderResponseBodyModelLmOrderList {
	s.LmOrderList = v
	return s
}

type CreateVirtualProductOrderResponseBodyModelLmOrderListLmOrderList struct {
	LmOrderId *string `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
}

func (s CreateVirtualProductOrderResponseBodyModelLmOrderListLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualProductOrderResponseBodyModelLmOrderListLmOrderList) GoString() string {
	return s.String()
}

func (s *CreateVirtualProductOrderResponseBodyModelLmOrderListLmOrderList) SetLmOrderId(v string) *CreateVirtualProductOrderResponseBodyModelLmOrderListLmOrderList {
	s.LmOrderId = &v
	return s
}

type CreateVirtualProductOrderResponseBodyModelPayTradeIds struct {
	PayTradeIds []*string `json:"PayTradeIds,omitempty" xml:"PayTradeIds,omitempty" type:"Repeated"`
}

func (s CreateVirtualProductOrderResponseBodyModelPayTradeIds) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualProductOrderResponseBodyModelPayTradeIds) GoString() string {
	return s.String()
}

func (s *CreateVirtualProductOrderResponseBodyModelPayTradeIds) SetPayTradeIds(v []*string) *CreateVirtualProductOrderResponseBodyModelPayTradeIds {
	s.PayTradeIds = v
	return s
}

type CreateVirtualProductOrderResponseBodyModelOrderIds struct {
	OrderIds []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s CreateVirtualProductOrderResponseBodyModelOrderIds) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualProductOrderResponseBodyModelOrderIds) GoString() string {
	return s.String()
}

func (s *CreateVirtualProductOrderResponseBodyModelOrderIds) SetOrderIds(v []*string) *CreateVirtualProductOrderResponseBodyModelOrderIds {
	s.OrderIds = v
	return s
}

type CreateVirtualProductOrderResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVirtualProductOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVirtualProductOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualProductOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualProductOrderResponse) SetHeaders(v map[string]*string) *CreateVirtualProductOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualProductOrderResponse) SetBody(v *CreateVirtualProductOrderResponseBody) *CreateVirtualProductOrderResponse {
	s.Body = v
	return s
}

type CreateWithholdTradeRequest struct {
	OutRequestNo *string `json:"OutRequestNo,omitempty" xml:"OutRequestNo,omitempty"`
	OutTradeNo   *string `json:"OutTradeNo,omitempty" xml:"OutTradeNo,omitempty"`
	Subject      *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	BuyerId      *string `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	MerchantId   *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
	TotalAmount  *string `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	Body         *string `json:"Body,omitempty" xml:"Body,omitempty"`
	GoodsDetail  *string `json:"GoodsDetail,omitempty" xml:"GoodsDetail,omitempty"`
	AgreementNo  *string `json:"AgreementNo,omitempty" xml:"AgreementNo,omitempty"`
	SettleMode   *string `json:"SettleMode,omitempty" xml:"SettleMode,omitempty"`
	ExtInfo      *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s CreateWithholdTradeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWithholdTradeRequest) GoString() string {
	return s.String()
}

func (s *CreateWithholdTradeRequest) SetOutRequestNo(v string) *CreateWithholdTradeRequest {
	s.OutRequestNo = &v
	return s
}

func (s *CreateWithholdTradeRequest) SetOutTradeNo(v string) *CreateWithholdTradeRequest {
	s.OutTradeNo = &v
	return s
}

func (s *CreateWithholdTradeRequest) SetSubject(v string) *CreateWithholdTradeRequest {
	s.Subject = &v
	return s
}

func (s *CreateWithholdTradeRequest) SetBuyerId(v string) *CreateWithholdTradeRequest {
	s.BuyerId = &v
	return s
}

func (s *CreateWithholdTradeRequest) SetMerchantId(v string) *CreateWithholdTradeRequest {
	s.MerchantId = &v
	return s
}

func (s *CreateWithholdTradeRequest) SetTotalAmount(v string) *CreateWithholdTradeRequest {
	s.TotalAmount = &v
	return s
}

func (s *CreateWithholdTradeRequest) SetBody(v string) *CreateWithholdTradeRequest {
	s.Body = &v
	return s
}

func (s *CreateWithholdTradeRequest) SetGoodsDetail(v string) *CreateWithholdTradeRequest {
	s.GoodsDetail = &v
	return s
}

func (s *CreateWithholdTradeRequest) SetAgreementNo(v string) *CreateWithholdTradeRequest {
	s.AgreementNo = &v
	return s
}

func (s *CreateWithholdTradeRequest) SetSettleMode(v string) *CreateWithholdTradeRequest {
	s.SettleMode = &v
	return s
}

func (s *CreateWithholdTradeRequest) SetExtInfo(v string) *CreateWithholdTradeRequest {
	s.ExtInfo = &v
	return s
}

type CreateWithholdTradeResponseBody struct {
	Code                  *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message               *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId             *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WithholdTradeResponse *CreateWithholdTradeResponseBodyWithholdTradeResponse `json:"WithholdTradeResponse,omitempty" xml:"WithholdTradeResponse,omitempty" type:"Struct"`
}

func (s CreateWithholdTradeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWithholdTradeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWithholdTradeResponseBody) SetCode(v string) *CreateWithholdTradeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWithholdTradeResponseBody) SetMessage(v string) *CreateWithholdTradeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWithholdTradeResponseBody) SetRequestId(v string) *CreateWithholdTradeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWithholdTradeResponseBody) SetWithholdTradeResponse(v *CreateWithholdTradeResponseBodyWithholdTradeResponse) *CreateWithholdTradeResponseBody {
	s.WithholdTradeResponse = v
	return s
}

type CreateWithholdTradeResponseBodyWithholdTradeResponse struct {
	PaymentDate  *string `json:"PaymentDate,omitempty" xml:"PaymentDate,omitempty"`
	TradeNo      *string `json:"TradeNo,omitempty" xml:"TradeNo,omitempty"`
	OutRequestNo *string `json:"OutRequestNo,omitempty" xml:"OutRequestNo,omitempty"`
	OutTradeNo   *string `json:"OutTradeNo,omitempty" xml:"OutTradeNo,omitempty"`
}

func (s CreateWithholdTradeResponseBodyWithholdTradeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWithholdTradeResponseBodyWithholdTradeResponse) GoString() string {
	return s.String()
}

func (s *CreateWithholdTradeResponseBodyWithholdTradeResponse) SetPaymentDate(v string) *CreateWithholdTradeResponseBodyWithholdTradeResponse {
	s.PaymentDate = &v
	return s
}

func (s *CreateWithholdTradeResponseBodyWithholdTradeResponse) SetTradeNo(v string) *CreateWithholdTradeResponseBodyWithholdTradeResponse {
	s.TradeNo = &v
	return s
}

func (s *CreateWithholdTradeResponseBodyWithholdTradeResponse) SetOutRequestNo(v string) *CreateWithholdTradeResponseBodyWithholdTradeResponse {
	s.OutRequestNo = &v
	return s
}

func (s *CreateWithholdTradeResponseBodyWithholdTradeResponse) SetOutTradeNo(v string) *CreateWithholdTradeResponseBodyWithholdTradeResponse {
	s.OutTradeNo = &v
	return s
}

type CreateWithholdTradeResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWithholdTradeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWithholdTradeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWithholdTradeResponse) GoString() string {
	return s.String()
}

func (s *CreateWithholdTradeResponse) SetHeaders(v map[string]*string) *CreateWithholdTradeResponse {
	s.Headers = v
	return s
}

func (s *CreateWithholdTradeResponse) SetBody(v *CreateWithholdTradeResponseBody) *CreateWithholdTradeResponse {
	s.Body = v
	return s
}

type DeleteBizItemsRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SubBizId   *string `json:"SubBizId,omitempty" xml:"SubBizId,omitempty"`
	ItemIdList []*int  `json:"ItemIdList,omitempty" xml:"ItemIdList,omitempty" type:"Repeated"`
}

func (s DeleteBizItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizItemsRequest) GoString() string {
	return s.String()
}

func (s *DeleteBizItemsRequest) SetBizId(v string) *DeleteBizItemsRequest {
	s.BizId = &v
	return s
}

func (s *DeleteBizItemsRequest) SetSubBizId(v string) *DeleteBizItemsRequest {
	s.SubBizId = &v
	return s
}

func (s *DeleteBizItemsRequest) SetItemIdList(v []*int) *DeleteBizItemsRequest {
	s.ItemIdList = v
	return s
}

type DeleteBizItemsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBizItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBizItemsResponseBody) SetCode(v string) *DeleteBizItemsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBizItemsResponseBody) SetMessage(v string) *DeleteBizItemsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBizItemsResponseBody) SetRequestId(v string) *DeleteBizItemsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBizItemsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBizItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBizItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBizItemsResponse) GoString() string {
	return s.String()
}

func (s *DeleteBizItemsResponse) SetHeaders(v map[string]*string) *DeleteBizItemsResponse {
	s.Headers = v
	return s
}

func (s *DeleteBizItemsResponse) SetBody(v *DeleteBizItemsResponseBody) *DeleteBizItemsResponse {
	s.Body = v
	return s
}

type DeleteItemLimitRuleRequest struct {
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SubBizCode   *string `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	LmActivityId *int64  `json:"LmActivityId,omitempty" xml:"LmActivityId,omitempty"`
	LmItemId     *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId       *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteItemLimitRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemLimitRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteItemLimitRuleRequest) SetBizId(v string) *DeleteItemLimitRuleRequest {
	s.BizId = &v
	return s
}

func (s *DeleteItemLimitRuleRequest) SetSubBizCode(v string) *DeleteItemLimitRuleRequest {
	s.SubBizCode = &v
	return s
}

func (s *DeleteItemLimitRuleRequest) SetLmActivityId(v int64) *DeleteItemLimitRuleRequest {
	s.LmActivityId = &v
	return s
}

func (s *DeleteItemLimitRuleRequest) SetLmItemId(v string) *DeleteItemLimitRuleRequest {
	s.LmItemId = &v
	return s
}

func (s *DeleteItemLimitRuleRequest) SetItemId(v int64) *DeleteItemLimitRuleRequest {
	s.ItemId = &v
	return s
}

func (s *DeleteItemLimitRuleRequest) SetRuleId(v int64) *DeleteItemLimitRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteItemLimitRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Model     *bool   `json:"Model,omitempty" xml:"Model,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteItemLimitRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemLimitRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteItemLimitRuleResponseBody) SetCode(v string) *DeleteItemLimitRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteItemLimitRuleResponseBody) SetModel(v bool) *DeleteItemLimitRuleResponseBody {
	s.Model = &v
	return s
}

func (s *DeleteItemLimitRuleResponseBody) SetMessage(v string) *DeleteItemLimitRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteItemLimitRuleResponseBody) SetRequestId(v string) *DeleteItemLimitRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteItemLimitRuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteItemLimitRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteItemLimitRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteItemLimitRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteItemLimitRuleResponse) SetHeaders(v map[string]*string) *DeleteItemLimitRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteItemLimitRuleResponse) SetBody(v *DeleteItemLimitRuleResponseBody) *DeleteItemLimitRuleResponse {
	s.Body = v
	return s
}

type EnableOrderRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	OutTradeId            *string `json:"OutTradeId,omitempty" xml:"OutTradeId,omitempty"`
	LmOrderId             *string `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	ExtJson               *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s EnableOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableOrderRequest) GoString() string {
	return s.String()
}

func (s *EnableOrderRequest) SetBizId(v string) *EnableOrderRequest {
	s.BizId = &v
	return s
}

func (s *EnableOrderRequest) SetBizUid(v string) *EnableOrderRequest {
	s.BizUid = &v
	return s
}

func (s *EnableOrderRequest) SetOutTradeId(v string) *EnableOrderRequest {
	s.OutTradeId = &v
	return s
}

func (s *EnableOrderRequest) SetLmOrderId(v string) *EnableOrderRequest {
	s.LmOrderId = &v
	return s
}

func (s *EnableOrderRequest) SetExtJson(v string) *EnableOrderRequest {
	s.ExtJson = &v
	return s
}

func (s *EnableOrderRequest) SetUseAnonymousTbAccount(v bool) *EnableOrderRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *EnableOrderRequest) SetThirdPartyUserId(v string) *EnableOrderRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *EnableOrderRequest) SetAccountType(v string) *EnableOrderRequest {
	s.AccountType = &v
	return s
}

type EnableOrderResponseBody struct {
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                       `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                       `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                       `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Model      *EnableOrderResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s EnableOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableOrderResponseBody) GoString() string {
	return s.String()
}

func (s *EnableOrderResponseBody) SetRequestId(v string) *EnableOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableOrderResponseBody) SetSuccess(v bool) *EnableOrderResponseBody {
	s.Success = &v
	return s
}

func (s *EnableOrderResponseBody) SetSubMessage(v string) *EnableOrderResponseBody {
	s.SubMessage = &v
	return s
}

func (s *EnableOrderResponseBody) SetCode(v string) *EnableOrderResponseBody {
	s.Code = &v
	return s
}

func (s *EnableOrderResponseBody) SetMessage(v string) *EnableOrderResponseBody {
	s.Message = &v
	return s
}

func (s *EnableOrderResponseBody) SetSubCode(v string) *EnableOrderResponseBody {
	s.SubCode = &v
	return s
}

func (s *EnableOrderResponseBody) SetTotalCount(v int64) *EnableOrderResponseBody {
	s.TotalCount = &v
	return s
}

func (s *EnableOrderResponseBody) SetLogsId(v string) *EnableOrderResponseBody {
	s.LogsId = &v
	return s
}

func (s *EnableOrderResponseBody) SetModel(v *EnableOrderResponseBodyModel) *EnableOrderResponseBody {
	s.Model = v
	return s
}

type EnableOrderResponseBodyModel struct {
	RedirectUrl *string                                  `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	LmOrderList *EnableOrderResponseBodyModelLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Struct"`
	PayTradeIds *EnableOrderResponseBodyModelPayTradeIds `json:"PayTradeIds,omitempty" xml:"PayTradeIds,omitempty" type:"Struct"`
	OrderIds    *EnableOrderResponseBodyModelOrderIds    `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Struct"`
}

func (s EnableOrderResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s EnableOrderResponseBodyModel) GoString() string {
	return s.String()
}

func (s *EnableOrderResponseBodyModel) SetRedirectUrl(v string) *EnableOrderResponseBodyModel {
	s.RedirectUrl = &v
	return s
}

func (s *EnableOrderResponseBodyModel) SetLmOrderList(v *EnableOrderResponseBodyModelLmOrderList) *EnableOrderResponseBodyModel {
	s.LmOrderList = v
	return s
}

func (s *EnableOrderResponseBodyModel) SetPayTradeIds(v *EnableOrderResponseBodyModelPayTradeIds) *EnableOrderResponseBodyModel {
	s.PayTradeIds = v
	return s
}

func (s *EnableOrderResponseBodyModel) SetOrderIds(v *EnableOrderResponseBodyModelOrderIds) *EnableOrderResponseBodyModel {
	s.OrderIds = v
	return s
}

type EnableOrderResponseBodyModelLmOrderList struct {
	LmOrderList []*EnableOrderResponseBodyModelLmOrderListLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Repeated"`
}

func (s EnableOrderResponseBodyModelLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s EnableOrderResponseBodyModelLmOrderList) GoString() string {
	return s.String()
}

func (s *EnableOrderResponseBodyModelLmOrderList) SetLmOrderList(v []*EnableOrderResponseBodyModelLmOrderListLmOrderList) *EnableOrderResponseBodyModelLmOrderList {
	s.LmOrderList = v
	return s
}

type EnableOrderResponseBodyModelLmOrderListLmOrderList struct {
	LmOrderId *string `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
}

func (s EnableOrderResponseBodyModelLmOrderListLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s EnableOrderResponseBodyModelLmOrderListLmOrderList) GoString() string {
	return s.String()
}

func (s *EnableOrderResponseBodyModelLmOrderListLmOrderList) SetLmOrderId(v string) *EnableOrderResponseBodyModelLmOrderListLmOrderList {
	s.LmOrderId = &v
	return s
}

type EnableOrderResponseBodyModelPayTradeIds struct {
	PayTradeIds []*string `json:"PayTradeIds,omitempty" xml:"PayTradeIds,omitempty" type:"Repeated"`
}

func (s EnableOrderResponseBodyModelPayTradeIds) String() string {
	return tea.Prettify(s)
}

func (s EnableOrderResponseBodyModelPayTradeIds) GoString() string {
	return s.String()
}

func (s *EnableOrderResponseBodyModelPayTradeIds) SetPayTradeIds(v []*string) *EnableOrderResponseBodyModelPayTradeIds {
	s.PayTradeIds = v
	return s
}

type EnableOrderResponseBodyModelOrderIds struct {
	OrderIds []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s EnableOrderResponseBodyModelOrderIds) String() string {
	return tea.Prettify(s)
}

func (s EnableOrderResponseBodyModelOrderIds) GoString() string {
	return s.String()
}

func (s *EnableOrderResponseBodyModelOrderIds) SetOrderIds(v []*string) *EnableOrderResponseBodyModelOrderIds {
	s.OrderIds = v
	return s
}

type EnableOrderResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableOrderResponse) GoString() string {
	return s.String()
}

func (s *EnableOrderResponse) SetHeaders(v map[string]*string) *EnableOrderResponse {
	s.Headers = v
	return s
}

func (s *EnableOrderResponse) SetBody(v *EnableOrderResponseBody) *EnableOrderResponse {
	s.Body = v
	return s
}

type ExecuteNodeRequest struct {
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	ProcessId         *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	NodeInstanceId    *string `json:"NodeInstanceId,omitempty" xml:"NodeInstanceId,omitempty"`
	NodeId            *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	RequestData       *string `json:"RequestData,omitempty" xml:"RequestData,omitempty"`
}

func (s ExecuteNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteNodeRequest) GoString() string {
	return s.String()
}

func (s *ExecuteNodeRequest) SetProcessInstanceId(v string) *ExecuteNodeRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *ExecuteNodeRequest) SetProcessId(v string) *ExecuteNodeRequest {
	s.ProcessId = &v
	return s
}

func (s *ExecuteNodeRequest) SetNodeInstanceId(v string) *ExecuteNodeRequest {
	s.NodeInstanceId = &v
	return s
}

func (s *ExecuteNodeRequest) SetNodeId(v string) *ExecuteNodeRequest {
	s.NodeId = &v
	return s
}

func (s *ExecuteNodeRequest) SetRequestData(v string) *ExecuteNodeRequest {
	s.RequestData = &v
	return s
}

type ExecuteNodeResponseBody struct {
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                       `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                       `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                       `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Model      *ExecuteNodeResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s ExecuteNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteNodeResponseBody) SetRequestId(v string) *ExecuteNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteNodeResponseBody) SetSuccess(v bool) *ExecuteNodeResponseBody {
	s.Success = &v
	return s
}

func (s *ExecuteNodeResponseBody) SetSubMessage(v string) *ExecuteNodeResponseBody {
	s.SubMessage = &v
	return s
}

func (s *ExecuteNodeResponseBody) SetCode(v string) *ExecuteNodeResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteNodeResponseBody) SetMessage(v string) *ExecuteNodeResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteNodeResponseBody) SetSubCode(v string) *ExecuteNodeResponseBody {
	s.SubCode = &v
	return s
}

func (s *ExecuteNodeResponseBody) SetTotalCount(v int64) *ExecuteNodeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ExecuteNodeResponseBody) SetLogsId(v string) *ExecuteNodeResponseBody {
	s.LogsId = &v
	return s
}

func (s *ExecuteNodeResponseBody) SetModel(v *ExecuteNodeResponseBodyModel) *ExecuteNodeResponseBody {
	s.Model = v
	return s
}

type ExecuteNodeResponseBodyModel struct {
	ResponseData      map[string]interface{} `json:"ResponseData,omitempty" xml:"ResponseData,omitempty"`
	ProcessInstanceId *string                `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
}

func (s ExecuteNodeResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s ExecuteNodeResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ExecuteNodeResponseBodyModel) SetResponseData(v map[string]interface{}) *ExecuteNodeResponseBodyModel {
	s.ResponseData = v
	return s
}

func (s *ExecuteNodeResponseBodyModel) SetProcessInstanceId(v string) *ExecuteNodeResponseBodyModel {
	s.ProcessInstanceId = &v
	return s
}

type ExecuteNodeResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteNodeResponse) GoString() string {
	return s.String()
}

func (s *ExecuteNodeResponse) SetHeaders(v map[string]*string) *ExecuteNodeResponse {
	s.Headers = v
	return s
}

func (s *ExecuteNodeResponse) SetBody(v *ExecuteNodeResponseBody) *ExecuteNodeResponse {
	s.Body = v
	return s
}

type GetCategoryChainRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	ItemId     *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId   *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
}

func (s GetCategoryChainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryChainRequest) GoString() string {
	return s.String()
}

func (s *GetCategoryChainRequest) SetBizId(v string) *GetCategoryChainRequest {
	s.BizId = &v
	return s
}

func (s *GetCategoryChainRequest) SetCategoryId(v int64) *GetCategoryChainRequest {
	s.CategoryId = &v
	return s
}

func (s *GetCategoryChainRequest) SetItemId(v int64) *GetCategoryChainRequest {
	s.ItemId = &v
	return s
}

func (s *GetCategoryChainRequest) SetLmItemId(v string) *GetCategoryChainRequest {
	s.LmItemId = &v
	return s
}

type GetCategoryChainResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CategoryList []*GetCategoryChainResponseBodyCategoryList `json:"CategoryList,omitempty" xml:"CategoryList,omitempty" type:"Repeated"`
}

func (s GetCategoryChainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryChainResponseBody) GoString() string {
	return s.String()
}

func (s *GetCategoryChainResponseBody) SetCode(v string) *GetCategoryChainResponseBody {
	s.Code = &v
	return s
}

func (s *GetCategoryChainResponseBody) SetMessage(v string) *GetCategoryChainResponseBody {
	s.Message = &v
	return s
}

func (s *GetCategoryChainResponseBody) SetRequestId(v string) *GetCategoryChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCategoryChainResponseBody) SetCategoryList(v []*GetCategoryChainResponseBodyCategoryList) *GetCategoryChainResponseBody {
	s.CategoryList = v
	return s
}

type GetCategoryChainResponseBodyCategoryList struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s GetCategoryChainResponseBodyCategoryList) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryChainResponseBodyCategoryList) GoString() string {
	return s.String()
}

func (s *GetCategoryChainResponseBodyCategoryList) SetName(v string) *GetCategoryChainResponseBodyCategoryList {
	s.Name = &v
	return s
}

func (s *GetCategoryChainResponseBodyCategoryList) SetCategoryId(v int64) *GetCategoryChainResponseBodyCategoryList {
	s.CategoryId = &v
	return s
}

type GetCategoryChainResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCategoryChainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCategoryChainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryChainResponse) GoString() string {
	return s.String()
}

func (s *GetCategoryChainResponse) SetHeaders(v map[string]*string) *GetCategoryChainResponse {
	s.Headers = v
	return s
}

func (s *GetCategoryChainResponse) SetBody(v *GetCategoryChainResponseBody) *GetCategoryChainResponse {
	s.Body = v
	return s
}

type GetCategoryListRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s GetCategoryListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryListRequest) GoString() string {
	return s.String()
}

func (s *GetCategoryListRequest) SetBizId(v string) *GetCategoryListRequest {
	s.BizId = &v
	return s
}

func (s *GetCategoryListRequest) SetCategoryId(v int64) *GetCategoryListRequest {
	s.CategoryId = &v
	return s
}

type GetCategoryListResponseBody struct {
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CategoryList *GetCategoryListResponseBodyCategoryList `json:"CategoryList,omitempty" xml:"CategoryList,omitempty" type:"Struct"`
}

func (s GetCategoryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCategoryListResponseBody) SetCode(v string) *GetCategoryListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCategoryListResponseBody) SetMessage(v string) *GetCategoryListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCategoryListResponseBody) SetRequestId(v string) *GetCategoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCategoryListResponseBody) SetCategoryList(v *GetCategoryListResponseBodyCategoryList) *GetCategoryListResponseBody {
	s.CategoryList = v
	return s
}

type GetCategoryListResponseBodyCategoryList struct {
	Category []*GetCategoryListResponseBodyCategoryListCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Repeated"`
}

func (s GetCategoryListResponseBodyCategoryList) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryListResponseBodyCategoryList) GoString() string {
	return s.String()
}

func (s *GetCategoryListResponseBodyCategoryList) SetCategory(v []*GetCategoryListResponseBodyCategoryListCategory) *GetCategoryListResponseBodyCategoryList {
	s.Category = v
	return s
}

type GetCategoryListResponseBodyCategoryListCategory struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s GetCategoryListResponseBodyCategoryListCategory) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryListResponseBodyCategoryListCategory) GoString() string {
	return s.String()
}

func (s *GetCategoryListResponseBodyCategoryListCategory) SetName(v string) *GetCategoryListResponseBodyCategoryListCategory {
	s.Name = &v
	return s
}

func (s *GetCategoryListResponseBodyCategoryListCategory) SetCategoryId(v int64) *GetCategoryListResponseBodyCategoryListCategory {
	s.CategoryId = &v
	return s
}

type GetCategoryListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCategoryListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCategoryListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCategoryListResponse) GoString() string {
	return s.String()
}

func (s *GetCategoryListResponse) SetHeaders(v map[string]*string) *GetCategoryListResponse {
	s.Headers = v
	return s
}

func (s *GetCategoryListResponse) SetBody(v *GetCategoryListResponseBody) *GetCategoryListResponse {
	s.Body = v
	return s
}

type GetCustomServiceUrlRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	SellerId              *string `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	Cuid                  *string `json:"Cuid,omitempty" xml:"Cuid,omitempty"`
	Nick                  *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s GetCustomServiceUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomServiceUrlRequest) GoString() string {
	return s.String()
}

func (s *GetCustomServiceUrlRequest) SetBizId(v string) *GetCustomServiceUrlRequest {
	s.BizId = &v
	return s
}

func (s *GetCustomServiceUrlRequest) SetBizUid(v string) *GetCustomServiceUrlRequest {
	s.BizUid = &v
	return s
}

func (s *GetCustomServiceUrlRequest) SetSellerId(v string) *GetCustomServiceUrlRequest {
	s.SellerId = &v
	return s
}

func (s *GetCustomServiceUrlRequest) SetCuid(v string) *GetCustomServiceUrlRequest {
	s.Cuid = &v
	return s
}

func (s *GetCustomServiceUrlRequest) SetNick(v string) *GetCustomServiceUrlRequest {
	s.Nick = &v
	return s
}

func (s *GetCustomServiceUrlRequest) SetUseAnonymousTbAccount(v bool) *GetCustomServiceUrlRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *GetCustomServiceUrlRequest) SetThirdPartyUserId(v string) *GetCustomServiceUrlRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *GetCustomServiceUrlRequest) SetAccountType(v string) *GetCustomServiceUrlRequest {
	s.AccountType = &v
	return s
}

type GetCustomServiceUrlResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UrlData   *GetCustomServiceUrlResponseBodyUrlData `json:"UrlData,omitempty" xml:"UrlData,omitempty" type:"Struct"`
}

func (s GetCustomServiceUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomServiceUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomServiceUrlResponseBody) SetCode(v string) *GetCustomServiceUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomServiceUrlResponseBody) SetMessage(v string) *GetCustomServiceUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomServiceUrlResponseBody) SetRequestId(v string) *GetCustomServiceUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomServiceUrlResponseBody) SetUrlData(v *GetCustomServiceUrlResponseBodyUrlData) *GetCustomServiceUrlResponseBody {
	s.UrlData = v
	return s
}

type GetCustomServiceUrlResponseBodyUrlData struct {
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
}

func (s GetCustomServiceUrlResponseBodyUrlData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomServiceUrlResponseBodyUrlData) GoString() string {
	return s.String()
}

func (s *GetCustomServiceUrlResponseBodyUrlData) SetReturnUrl(v string) *GetCustomServiceUrlResponseBodyUrlData {
	s.ReturnUrl = &v
	return s
}

type GetCustomServiceUrlResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCustomServiceUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCustomServiceUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomServiceUrlResponse) GoString() string {
	return s.String()
}

func (s *GetCustomServiceUrlResponse) SetHeaders(v map[string]*string) *GetCustomServiceUrlResponse {
	s.Headers = v
	return s
}

func (s *GetCustomServiceUrlResponse) SetBody(v *GetCustomServiceUrlResponseBody) *GetCustomServiceUrlResponse {
	s.Body = v
	return s
}

type GetGuidePageRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s GetGuidePageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGuidePageRequest) GoString() string {
	return s.String()
}

func (s *GetGuidePageRequest) SetBizId(v string) *GetGuidePageRequest {
	s.BizId = &v
	return s
}

type GetGuidePageResponseBody struct {
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MiniShopInfo []*GetGuidePageResponseBodyMiniShopInfo `json:"MiniShopInfo,omitempty" xml:"MiniShopInfo,omitempty" type:"Repeated"`
}

func (s GetGuidePageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGuidePageResponseBody) GoString() string {
	return s.String()
}

func (s *GetGuidePageResponseBody) SetCode(v string) *GetGuidePageResponseBody {
	s.Code = &v
	return s
}

func (s *GetGuidePageResponseBody) SetMessage(v string) *GetGuidePageResponseBody {
	s.Message = &v
	return s
}

func (s *GetGuidePageResponseBody) SetRequestId(v string) *GetGuidePageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGuidePageResponseBody) SetMiniShopInfo(v []*GetGuidePageResponseBodyMiniShopInfo) *GetGuidePageResponseBody {
	s.MiniShopInfo = v
	return s
}

type GetGuidePageResponseBodyMiniShopInfo struct {
	LiteShopId *string `json:"LiteShopId,omitempty" xml:"LiteShopId,omitempty"`
	Src        *string `json:"Src,omitempty" xml:"Src,omitempty"`
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetGuidePageResponseBodyMiniShopInfo) String() string {
	return tea.Prettify(s)
}

func (s GetGuidePageResponseBodyMiniShopInfo) GoString() string {
	return s.String()
}

func (s *GetGuidePageResponseBodyMiniShopInfo) SetLiteShopId(v string) *GetGuidePageResponseBodyMiniShopInfo {
	s.LiteShopId = &v
	return s
}

func (s *GetGuidePageResponseBodyMiniShopInfo) SetSrc(v string) *GetGuidePageResponseBodyMiniShopInfo {
	s.Src = &v
	return s
}

func (s *GetGuidePageResponseBodyMiniShopInfo) SetBizId(v string) *GetGuidePageResponseBodyMiniShopInfo {
	s.BizId = &v
	return s
}

func (s *GetGuidePageResponseBodyMiniShopInfo) SetName(v string) *GetGuidePageResponseBodyMiniShopInfo {
	s.Name = &v
	return s
}

type GetGuidePageResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGuidePageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGuidePageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGuidePageResponse) GoString() string {
	return s.String()
}

func (s *GetGuidePageResponse) SetHeaders(v map[string]*string) *GetGuidePageResponse {
	s.Headers = v
	return s
}

func (s *GetGuidePageResponse) SetBody(v *GetGuidePageResponseBody) *GetGuidePageResponse {
	s.Body = v
	return s
}

type GetItemPromotionRequest struct {
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ItemId   *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
}

func (s GetItemPromotionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetItemPromotionRequest) GoString() string {
	return s.String()
}

func (s *GetItemPromotionRequest) SetBizId(v string) *GetItemPromotionRequest {
	s.BizId = &v
	return s
}

func (s *GetItemPromotionRequest) SetItemId(v int64) *GetItemPromotionRequest {
	s.ItemId = &v
	return s
}

func (s *GetItemPromotionRequest) SetLmItemId(v string) *GetItemPromotionRequest {
	s.LmItemId = &v
	return s
}

type GetItemPromotionResponseBody struct {
	RequestId          *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage         *string                                         `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code               *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode            *string                                         `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	ItemPromotionModel *GetItemPromotionResponseBodyItemPromotionModel `json:"ItemPromotionModel,omitempty" xml:"ItemPromotionModel,omitempty" type:"Struct"`
}

func (s GetItemPromotionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetItemPromotionResponseBody) GoString() string {
	return s.String()
}

func (s *GetItemPromotionResponseBody) SetRequestId(v string) *GetItemPromotionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetItemPromotionResponseBody) SetSuccess(v bool) *GetItemPromotionResponseBody {
	s.Success = &v
	return s
}

func (s *GetItemPromotionResponseBody) SetSubMessage(v string) *GetItemPromotionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *GetItemPromotionResponseBody) SetCode(v string) *GetItemPromotionResponseBody {
	s.Code = &v
	return s
}

func (s *GetItemPromotionResponseBody) SetMessage(v string) *GetItemPromotionResponseBody {
	s.Message = &v
	return s
}

func (s *GetItemPromotionResponseBody) SetSubCode(v string) *GetItemPromotionResponseBody {
	s.SubCode = &v
	return s
}

func (s *GetItemPromotionResponseBody) SetItemPromotionModel(v *GetItemPromotionResponseBodyItemPromotionModel) *GetItemPromotionResponseBody {
	s.ItemPromotionModel = v
	return s
}

type GetItemPromotionResponseBodyItemPromotionModel struct {
	EndTime       *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime     *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SkuPromotion  map[string]interface{} `json:"SkuPromotion,omitempty" xml:"SkuPromotion,omitempty"`
	PromotionFlag *bool                  `json:"PromotionFlag,omitempty" xml:"PromotionFlag,omitempty"`
	PromotionName *string                `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	LmItemId      *string                `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	PromotionDesc *string                `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	ItemId        *int64                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	PromotionId   *string                `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	ExtInfo       map[string]interface{} `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s GetItemPromotionResponseBodyItemPromotionModel) String() string {
	return tea.Prettify(s)
}

func (s GetItemPromotionResponseBodyItemPromotionModel) GoString() string {
	return s.String()
}

func (s *GetItemPromotionResponseBodyItemPromotionModel) SetEndTime(v string) *GetItemPromotionResponseBodyItemPromotionModel {
	s.EndTime = &v
	return s
}

func (s *GetItemPromotionResponseBodyItemPromotionModel) SetStartTime(v string) *GetItemPromotionResponseBodyItemPromotionModel {
	s.StartTime = &v
	return s
}

func (s *GetItemPromotionResponseBodyItemPromotionModel) SetSkuPromotion(v map[string]interface{}) *GetItemPromotionResponseBodyItemPromotionModel {
	s.SkuPromotion = v
	return s
}

func (s *GetItemPromotionResponseBodyItemPromotionModel) SetPromotionFlag(v bool) *GetItemPromotionResponseBodyItemPromotionModel {
	s.PromotionFlag = &v
	return s
}

func (s *GetItemPromotionResponseBodyItemPromotionModel) SetPromotionName(v string) *GetItemPromotionResponseBodyItemPromotionModel {
	s.PromotionName = &v
	return s
}

func (s *GetItemPromotionResponseBodyItemPromotionModel) SetLmItemId(v string) *GetItemPromotionResponseBodyItemPromotionModel {
	s.LmItemId = &v
	return s
}

func (s *GetItemPromotionResponseBodyItemPromotionModel) SetPromotionDesc(v string) *GetItemPromotionResponseBodyItemPromotionModel {
	s.PromotionDesc = &v
	return s
}

func (s *GetItemPromotionResponseBodyItemPromotionModel) SetItemId(v int64) *GetItemPromotionResponseBodyItemPromotionModel {
	s.ItemId = &v
	return s
}

func (s *GetItemPromotionResponseBodyItemPromotionModel) SetPromotionId(v string) *GetItemPromotionResponseBodyItemPromotionModel {
	s.PromotionId = &v
	return s
}

func (s *GetItemPromotionResponseBodyItemPromotionModel) SetExtInfo(v map[string]interface{}) *GetItemPromotionResponseBodyItemPromotionModel {
	s.ExtInfo = v
	return s
}

type GetItemPromotionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetItemPromotionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetItemPromotionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetItemPromotionResponse) GoString() string {
	return s.String()
}

func (s *GetItemPromotionResponse) SetHeaders(v map[string]*string) *GetItemPromotionResponse {
	s.Headers = v
	return s
}

func (s *GetItemPromotionResponse) SetBody(v *GetItemPromotionResponseBody) *GetItemPromotionResponse {
	s.Body = v
	return s
}

type GetLoginPageRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	TargetUrl             *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	FailUrl               *string `json:"FailUrl,omitempty" xml:"FailUrl,omitempty"`
}

func (s GetLoginPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginPageRequest) GoString() string {
	return s.String()
}

func (s *GetLoginPageRequest) SetBizId(v string) *GetLoginPageRequest {
	s.BizId = &v
	return s
}

func (s *GetLoginPageRequest) SetUseAnonymousTbAccount(v bool) *GetLoginPageRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *GetLoginPageRequest) SetThirdPartyUserId(v string) *GetLoginPageRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *GetLoginPageRequest) SetTargetUrl(v string) *GetLoginPageRequest {
	s.TargetUrl = &v
	return s
}

func (s *GetLoginPageRequest) SetFailUrl(v string) *GetLoginPageRequest {
	s.FailUrl = &v
	return s
}

type GetLoginPageResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UrlData   *GetLoginPageResponseBodyUrlData `json:"UrlData,omitempty" xml:"UrlData,omitempty" type:"Struct"`
}

func (s GetLoginPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoginPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginPageResponseBody) SetCode(v string) *GetLoginPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetLoginPageResponseBody) SetMessage(v string) *GetLoginPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetLoginPageResponseBody) SetRequestId(v string) *GetLoginPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginPageResponseBody) SetUrlData(v *GetLoginPageResponseBodyUrlData) *GetLoginPageResponseBody {
	s.UrlData = v
	return s
}

type GetLoginPageResponseBodyUrlData struct {
	ReturnUrl *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
}

func (s GetLoginPageResponseBodyUrlData) String() string {
	return tea.Prettify(s)
}

func (s GetLoginPageResponseBodyUrlData) GoString() string {
	return s.String()
}

func (s *GetLoginPageResponseBodyUrlData) SetReturnUrl(v string) *GetLoginPageResponseBodyUrlData {
	s.ReturnUrl = &v
	return s
}

type GetLoginPageResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLoginPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLoginPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoginPageResponse) GoString() string {
	return s.String()
}

func (s *GetLoginPageResponse) SetHeaders(v map[string]*string) *GetLoginPageResponse {
	s.Headers = v
	return s
}

func (s *GetLoginPageResponse) SetBody(v *GetLoginPageResponseBody) *GetLoginPageResponse {
	s.Body = v
	return s
}

type GetSwitchUrlRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	Url                   *string `json:"Url,omitempty" xml:"Url,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
}

func (s GetSwitchUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSwitchUrlRequest) GoString() string {
	return s.String()
}

func (s *GetSwitchUrlRequest) SetBizId(v string) *GetSwitchUrlRequest {
	s.BizId = &v
	return s
}

func (s *GetSwitchUrlRequest) SetBizUid(v string) *GetSwitchUrlRequest {
	s.BizUid = &v
	return s
}

func (s *GetSwitchUrlRequest) SetUrl(v string) *GetSwitchUrlRequest {
	s.Url = &v
	return s
}

func (s *GetSwitchUrlRequest) SetUseAnonymousTbAccount(v bool) *GetSwitchUrlRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *GetSwitchUrlRequest) SetThirdPartyUserId(v string) *GetSwitchUrlRequest {
	s.ThirdPartyUserId = &v
	return s
}

type GetSwitchUrlResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSwitchUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSwitchUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetSwitchUrlResponseBody) SetCode(v string) *GetSwitchUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetSwitchUrlResponseBody) SetUrl(v string) *GetSwitchUrlResponseBody {
	s.Url = &v
	return s
}

func (s *GetSwitchUrlResponseBody) SetMessage(v string) *GetSwitchUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetSwitchUrlResponseBody) SetRequestId(v string) *GetSwitchUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetSwitchUrlResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSwitchUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSwitchUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSwitchUrlResponse) GoString() string {
	return s.String()
}

func (s *GetSwitchUrlResponse) SetHeaders(v map[string]*string) *GetSwitchUrlResponse {
	s.Headers = v
	return s
}

func (s *GetSwitchUrlResponse) SetBody(v *GetSwitchUrlResponseBody) *GetSwitchUrlResponse {
	s.Body = v
	return s
}

type GetUserInfoRequest struct {
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	UserFlag  *string `json:"UserFlag,omitempty" xml:"UserFlag,omitempty"`
	AppName   *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	QueryJson *string `json:"QueryJson,omitempty" xml:"QueryJson,omitempty"`
}

func (s GetUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserInfoRequest) SetBizId(v string) *GetUserInfoRequest {
	s.BizId = &v
	return s
}

func (s *GetUserInfoRequest) SetUserFlag(v string) *GetUserInfoRequest {
	s.UserFlag = &v
	return s
}

func (s *GetUserInfoRequest) SetAppName(v string) *GetUserInfoRequest {
	s.AppName = &v
	return s
}

func (s *GetUserInfoRequest) SetQueryJson(v string) *GetUserInfoRequest {
	s.QueryJson = &v
	return s
}

type GetUserInfoResponseBody struct {
	Code        *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LoginResult *GetUserInfoResponseBodyLoginResult `json:"LoginResult,omitempty" xml:"LoginResult,omitempty" type:"Struct"`
}

func (s GetUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBody) SetCode(v string) *GetUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserInfoResponseBody) SetMessage(v string) *GetUserInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserInfoResponseBody) SetRequestId(v string) *GetUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserInfoResponseBody) SetLoginResult(v *GetUserInfoResponseBodyLoginResult) *GetUserInfoResponseBody {
	s.LoginResult = v
	return s
}

type GetUserInfoResponseBodyLoginResult struct {
	ReturnUrl   *string                `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	BizUserName *string                `json:"BizUserName,omitempty" xml:"BizUserName,omitempty"`
	BizUid      *string                `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	BizId       *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LmUserId    *int64                 `json:"LmUserId,omitempty" xml:"LmUserId,omitempty"`
	ExtInfo     map[string]interface{} `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	SubBizId    []*string              `json:"SubBizId,omitempty" xml:"SubBizId,omitempty" type:"Repeated"`
}

func (s GetUserInfoResponseBodyLoginResult) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponseBodyLoginResult) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBodyLoginResult) SetReturnUrl(v string) *GetUserInfoResponseBodyLoginResult {
	s.ReturnUrl = &v
	return s
}

func (s *GetUserInfoResponseBodyLoginResult) SetBizUserName(v string) *GetUserInfoResponseBodyLoginResult {
	s.BizUserName = &v
	return s
}

func (s *GetUserInfoResponseBodyLoginResult) SetBizUid(v string) *GetUserInfoResponseBodyLoginResult {
	s.BizUid = &v
	return s
}

func (s *GetUserInfoResponseBodyLoginResult) SetBizId(v string) *GetUserInfoResponseBodyLoginResult {
	s.BizId = &v
	return s
}

func (s *GetUserInfoResponseBodyLoginResult) SetLmUserId(v int64) *GetUserInfoResponseBodyLoginResult {
	s.LmUserId = &v
	return s
}

func (s *GetUserInfoResponseBodyLoginResult) SetExtInfo(v map[string]interface{}) *GetUserInfoResponseBodyLoginResult {
	s.ExtInfo = v
	return s
}

func (s *GetUserInfoResponseBodyLoginResult) SetSubBizId(v []*string) *GetUserInfoResponseBodyLoginResult {
	s.SubBizId = v
	return s
}

type GetUserInfoResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponse) SetHeaders(v map[string]*string) *GetUserInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserInfoResponse) SetBody(v *GetUserInfoResponseBody) *GetUserInfoResponse {
	s.Body = v
	return s
}

type GetWithholdSignPageUrlRequest struct {
	OutRequestNo               *string `json:"OutRequestNo,omitempty" xml:"OutRequestNo,omitempty"`
	ExternalAgreementNo        *string `json:"ExternalAgreementNo,omitempty" xml:"ExternalAgreementNo,omitempty"`
	MerchantId                 *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
	MerchantServiceName        *string `json:"MerchantServiceName,omitempty" xml:"MerchantServiceName,omitempty"`
	MerchantServiceDescription *string `json:"MerchantServiceDescription,omitempty" xml:"MerchantServiceDescription,omitempty"`
	IdentityParameters         *string `json:"IdentityParameters,omitempty" xml:"IdentityParameters,omitempty"`
	ReturnUrl                  *string `json:"ReturnUrl,omitempty" xml:"ReturnUrl,omitempty"`
	NotifyUrl                  *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	DeviceType                 *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	ExtInfo                    *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s GetWithholdSignPageUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWithholdSignPageUrlRequest) GoString() string {
	return s.String()
}

func (s *GetWithholdSignPageUrlRequest) SetOutRequestNo(v string) *GetWithholdSignPageUrlRequest {
	s.OutRequestNo = &v
	return s
}

func (s *GetWithholdSignPageUrlRequest) SetExternalAgreementNo(v string) *GetWithholdSignPageUrlRequest {
	s.ExternalAgreementNo = &v
	return s
}

func (s *GetWithholdSignPageUrlRequest) SetMerchantId(v string) *GetWithholdSignPageUrlRequest {
	s.MerchantId = &v
	return s
}

func (s *GetWithholdSignPageUrlRequest) SetMerchantServiceName(v string) *GetWithholdSignPageUrlRequest {
	s.MerchantServiceName = &v
	return s
}

func (s *GetWithholdSignPageUrlRequest) SetMerchantServiceDescription(v string) *GetWithholdSignPageUrlRequest {
	s.MerchantServiceDescription = &v
	return s
}

func (s *GetWithholdSignPageUrlRequest) SetIdentityParameters(v string) *GetWithholdSignPageUrlRequest {
	s.IdentityParameters = &v
	return s
}

func (s *GetWithholdSignPageUrlRequest) SetReturnUrl(v string) *GetWithholdSignPageUrlRequest {
	s.ReturnUrl = &v
	return s
}

func (s *GetWithholdSignPageUrlRequest) SetNotifyUrl(v string) *GetWithholdSignPageUrlRequest {
	s.NotifyUrl = &v
	return s
}

func (s *GetWithholdSignPageUrlRequest) SetDeviceType(v string) *GetWithholdSignPageUrlRequest {
	s.DeviceType = &v
	return s
}

func (s *GetWithholdSignPageUrlRequest) SetExtInfo(v string) *GetWithholdSignPageUrlRequest {
	s.ExtInfo = &v
	return s
}

type GetWithholdSignPageUrlResponseBody struct {
	Code                 *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message              *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WithholdSignResponse *GetWithholdSignPageUrlResponseBodyWithholdSignResponse `json:"WithholdSignResponse,omitempty" xml:"WithholdSignResponse,omitempty" type:"Struct"`
}

func (s GetWithholdSignPageUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWithholdSignPageUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetWithholdSignPageUrlResponseBody) SetCode(v string) *GetWithholdSignPageUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetWithholdSignPageUrlResponseBody) SetMessage(v string) *GetWithholdSignPageUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetWithholdSignPageUrlResponseBody) SetRequestId(v string) *GetWithholdSignPageUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWithholdSignPageUrlResponseBody) SetWithholdSignResponse(v *GetWithholdSignPageUrlResponseBodyWithholdSignResponse) *GetWithholdSignPageUrlResponseBody {
	s.WithholdSignResponse = v
	return s
}

type GetWithholdSignPageUrlResponseBodyWithholdSignResponse struct {
	PageUrl      *string `json:"PageUrl,omitempty" xml:"PageUrl,omitempty"`
	OutRequestNo *string `json:"OutRequestNo,omitempty" xml:"OutRequestNo,omitempty"`
}

func (s GetWithholdSignPageUrlResponseBodyWithholdSignResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWithholdSignPageUrlResponseBodyWithholdSignResponse) GoString() string {
	return s.String()
}

func (s *GetWithholdSignPageUrlResponseBodyWithholdSignResponse) SetPageUrl(v string) *GetWithholdSignPageUrlResponseBodyWithholdSignResponse {
	s.PageUrl = &v
	return s
}

func (s *GetWithholdSignPageUrlResponseBodyWithholdSignResponse) SetOutRequestNo(v string) *GetWithholdSignPageUrlResponseBodyWithholdSignResponse {
	s.OutRequestNo = &v
	return s
}

type GetWithholdSignPageUrlResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWithholdSignPageUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWithholdSignPageUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWithholdSignPageUrlResponse) GoString() string {
	return s.String()
}

func (s *GetWithholdSignPageUrlResponse) SetHeaders(v map[string]*string) *GetWithholdSignPageUrlResponse {
	s.Headers = v
	return s
}

func (s *GetWithholdSignPageUrlResponse) SetBody(v *GetWithholdSignPageUrlResponseBody) *GetWithholdSignPageUrlResponse {
	s.Body = v
	return s
}

type InitApplyRefundRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	SubLmOrderId          *string `json:"SubLmOrderId,omitempty" xml:"SubLmOrderId,omitempty"`
	BizClaimType          *int32  `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	GoodsStatus           *int32  `json:"GoodsStatus,omitempty" xml:"GoodsStatus,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s InitApplyRefundRequest) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefundRequest) GoString() string {
	return s.String()
}

func (s *InitApplyRefundRequest) SetBizId(v string) *InitApplyRefundRequest {
	s.BizId = &v
	return s
}

func (s *InitApplyRefundRequest) SetBizUid(v string) *InitApplyRefundRequest {
	s.BizUid = &v
	return s
}

func (s *InitApplyRefundRequest) SetSubLmOrderId(v string) *InitApplyRefundRequest {
	s.SubLmOrderId = &v
	return s
}

func (s *InitApplyRefundRequest) SetBizClaimType(v int32) *InitApplyRefundRequest {
	s.BizClaimType = &v
	return s
}

func (s *InitApplyRefundRequest) SetGoodsStatus(v int32) *InitApplyRefundRequest {
	s.GoodsStatus = &v
	return s
}

func (s *InitApplyRefundRequest) SetUseAnonymousTbAccount(v bool) *InitApplyRefundRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *InitApplyRefundRequest) SetThirdPartyUserId(v string) *InitApplyRefundRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *InitApplyRefundRequest) SetAccountType(v string) *InitApplyRefundRequest {
	s.AccountType = &v
	return s
}

type InitApplyRefundResponseBody struct {
	Code                *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message             *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	SubLmOrderId        *string                                         `json:"SubLmOrderId,omitempty" xml:"SubLmOrderId,omitempty"`
	RequestId           *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InitApplyRefundData *InitApplyRefundResponseBodyInitApplyRefundData `json:"InitApplyRefundData,omitempty" xml:"InitApplyRefundData,omitempty" type:"Struct"`
}

func (s InitApplyRefundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefundResponseBody) GoString() string {
	return s.String()
}

func (s *InitApplyRefundResponseBody) SetCode(v string) *InitApplyRefundResponseBody {
	s.Code = &v
	return s
}

func (s *InitApplyRefundResponseBody) SetMessage(v string) *InitApplyRefundResponseBody {
	s.Message = &v
	return s
}

func (s *InitApplyRefundResponseBody) SetSubLmOrderId(v string) *InitApplyRefundResponseBody {
	s.SubLmOrderId = &v
	return s
}

func (s *InitApplyRefundResponseBody) SetRequestId(v string) *InitApplyRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitApplyRefundResponseBody) SetInitApplyRefundData(v *InitApplyRefundResponseBodyInitApplyRefundData) *InitApplyRefundResponseBody {
	s.InitApplyRefundData = v
	return s
}

type InitApplyRefundResponseBodyInitApplyRefundData struct {
	MainOrderRefund  *bool                                                           `json:"MainOrderRefund,omitempty" xml:"MainOrderRefund,omitempty"`
	BizClaimType     *int32                                                          `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	RefundReasonList *InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonList `json:"RefundReasonList,omitempty" xml:"RefundReasonList,omitempty" type:"Struct"`
	MaxRefundFeeData *InitApplyRefundResponseBodyInitApplyRefundDataMaxRefundFeeData `json:"MaxRefundFeeData,omitempty" xml:"MaxRefundFeeData,omitempty" type:"Struct"`
}

func (s InitApplyRefundResponseBodyInitApplyRefundData) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefundResponseBodyInitApplyRefundData) GoString() string {
	return s.String()
}

func (s *InitApplyRefundResponseBodyInitApplyRefundData) SetMainOrderRefund(v bool) *InitApplyRefundResponseBodyInitApplyRefundData {
	s.MainOrderRefund = &v
	return s
}

func (s *InitApplyRefundResponseBodyInitApplyRefundData) SetBizClaimType(v int32) *InitApplyRefundResponseBodyInitApplyRefundData {
	s.BizClaimType = &v
	return s
}

func (s *InitApplyRefundResponseBodyInitApplyRefundData) SetRefundReasonList(v *InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonList) *InitApplyRefundResponseBodyInitApplyRefundData {
	s.RefundReasonList = v
	return s
}

func (s *InitApplyRefundResponseBodyInitApplyRefundData) SetMaxRefundFeeData(v *InitApplyRefundResponseBodyInitApplyRefundDataMaxRefundFeeData) *InitApplyRefundResponseBodyInitApplyRefundData {
	s.MaxRefundFeeData = v
	return s
}

type InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonList struct {
	RefundReasonList []*InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList `json:"RefundReasonList,omitempty" xml:"RefundReasonList,omitempty" type:"Repeated"`
}

func (s InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonList) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonList) GoString() string {
	return s.String()
}

func (s *InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonList) SetRefundReasonList(v []*InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList) *InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonList {
	s.RefundReasonList = v
	return s
}

type InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList struct {
	ReasonTextId       *int64  `json:"ReasonTextId,omitempty" xml:"ReasonTextId,omitempty"`
	ReasonTips         *string `json:"ReasonTips,omitempty" xml:"ReasonTips,omitempty"`
	ProofRequired      *bool   `json:"ProofRequired,omitempty" xml:"ProofRequired,omitempty"`
	RefundDescRequired *bool   `json:"RefundDescRequired,omitempty" xml:"RefundDescRequired,omitempty"`
}

func (s InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList) GoString() string {
	return s.String()
}

func (s *InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList) SetReasonTextId(v int64) *InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList {
	s.ReasonTextId = &v
	return s
}

func (s *InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList) SetReasonTips(v string) *InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList {
	s.ReasonTips = &v
	return s
}

func (s *InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList) SetProofRequired(v bool) *InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList {
	s.ProofRequired = &v
	return s
}

func (s *InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList) SetRefundDescRequired(v bool) *InitApplyRefundResponseBodyInitApplyRefundDataRefundReasonListRefundReasonList {
	s.RefundDescRequired = &v
	return s
}

type InitApplyRefundResponseBodyInitApplyRefundDataMaxRefundFeeData struct {
	MaxRefundFee *int32 `json:"MaxRefundFee,omitempty" xml:"MaxRefundFee,omitempty"`
	MinRefundFee *int32 `json:"MinRefundFee,omitempty" xml:"MinRefundFee,omitempty"`
}

func (s InitApplyRefundResponseBodyInitApplyRefundDataMaxRefundFeeData) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefundResponseBodyInitApplyRefundDataMaxRefundFeeData) GoString() string {
	return s.String()
}

func (s *InitApplyRefundResponseBodyInitApplyRefundDataMaxRefundFeeData) SetMaxRefundFee(v int32) *InitApplyRefundResponseBodyInitApplyRefundDataMaxRefundFeeData {
	s.MaxRefundFee = &v
	return s
}

func (s *InitApplyRefundResponseBodyInitApplyRefundDataMaxRefundFeeData) SetMinRefundFee(v int32) *InitApplyRefundResponseBodyInitApplyRefundDataMaxRefundFeeData {
	s.MinRefundFee = &v
	return s
}

type InitApplyRefundResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitApplyRefundResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitApplyRefundResponse) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefundResponse) GoString() string {
	return s.String()
}

func (s *InitApplyRefundResponse) SetHeaders(v map[string]*string) *InitApplyRefundResponse {
	s.Headers = v
	return s
}

func (s *InitApplyRefundResponse) SetBody(v *InitApplyRefundResponseBody) *InitApplyRefundResponse {
	s.Body = v
	return s
}

type ListItemActivitiesRequest struct {
	BizId     *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LmItemIds map[string]interface{} `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty"`
	ItemIds   map[string]interface{} `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
}

func (s ListItemActivitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListItemActivitiesRequest) GoString() string {
	return s.String()
}

func (s *ListItemActivitiesRequest) SetBizId(v string) *ListItemActivitiesRequest {
	s.BizId = &v
	return s
}

func (s *ListItemActivitiesRequest) SetLmItemIds(v map[string]interface{}) *ListItemActivitiesRequest {
	s.LmItemIds = v
	return s
}

func (s *ListItemActivitiesRequest) SetItemIds(v map[string]interface{}) *ListItemActivitiesRequest {
	s.ItemIds = v
	return s
}

type ListItemActivitiesShrinkRequest struct {
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LmItemIdsShrink *string `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty"`
	ItemIdsShrink   *string `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
}

func (s ListItemActivitiesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListItemActivitiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListItemActivitiesShrinkRequest) SetBizId(v string) *ListItemActivitiesShrinkRequest {
	s.BizId = &v
	return s
}

func (s *ListItemActivitiesShrinkRequest) SetLmItemIdsShrink(v string) *ListItemActivitiesShrinkRequest {
	s.LmItemIdsShrink = &v
	return s
}

func (s *ListItemActivitiesShrinkRequest) SetItemIdsShrink(v string) *ListItemActivitiesShrinkRequest {
	s.ItemIdsShrink = &v
	return s
}

type ListItemActivitiesResponseBody struct {
	Code                    *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                 *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId               *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LmItemActivityModelList *ListItemActivitiesResponseBodyLmItemActivityModelList `json:"LmItemActivityModelList,omitempty" xml:"LmItemActivityModelList,omitempty" type:"Struct"`
}

func (s ListItemActivitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListItemActivitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListItemActivitiesResponseBody) SetCode(v string) *ListItemActivitiesResponseBody {
	s.Code = &v
	return s
}

func (s *ListItemActivitiesResponseBody) SetMessage(v string) *ListItemActivitiesResponseBody {
	s.Message = &v
	return s
}

func (s *ListItemActivitiesResponseBody) SetRequestId(v string) *ListItemActivitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListItemActivitiesResponseBody) SetLmItemActivityModelList(v *ListItemActivitiesResponseBodyLmItemActivityModelList) *ListItemActivitiesResponseBody {
	s.LmItemActivityModelList = v
	return s
}

type ListItemActivitiesResponseBodyLmItemActivityModelList struct {
	LmItemActivityModel []*ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModel `json:"LmItemActivityModel,omitempty" xml:"LmItemActivityModel,omitempty" type:"Repeated"`
}

func (s ListItemActivitiesResponseBodyLmItemActivityModelList) String() string {
	return tea.Prettify(s)
}

func (s ListItemActivitiesResponseBodyLmItemActivityModelList) GoString() string {
	return s.String()
}

func (s *ListItemActivitiesResponseBodyLmItemActivityModelList) SetLmItemActivityModel(v []*ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModel) *ListItemActivitiesResponseBodyLmItemActivityModelList {
	s.LmItemActivityModel = v
	return s
}

type ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModel struct {
	LmItemId           *string                                                                                     `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId             *int64                                                                                      `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmActivityPopModel *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel `json:"LmActivityPopModel,omitempty" xml:"LmActivityPopModel,omitempty" type:"Struct"`
}

func (s ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModel) String() string {
	return tea.Prettify(s)
}

func (s ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModel) GoString() string {
	return s.String()
}

func (s *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModel) SetLmItemId(v string) *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModel {
	s.LmItemId = &v
	return s
}

func (s *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModel) SetItemId(v int64) *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModel {
	s.ItemId = &v
	return s
}

func (s *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModel) SetLmActivityPopModel(v *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel) *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModel {
	s.LmActivityPopModel = v
	return s
}

type ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel struct {
	SubBizCode  *string                `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	EndDate     *string                `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	DisplayDate *string                `json:"DisplayDate,omitempty" xml:"DisplayDate,omitempty"`
	BizId       *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	StartDate   *string                `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Title       *string                `json:"Title,omitempty" xml:"Title,omitempty"`
	ExtInfo     map[string]interface{} `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	Id          *int64                 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel) String() string {
	return tea.Prettify(s)
}

func (s ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel) GoString() string {
	return s.String()
}

func (s *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel) SetSubBizCode(v string) *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel {
	s.SubBizCode = &v
	return s
}

func (s *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel) SetEndDate(v string) *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel {
	s.EndDate = &v
	return s
}

func (s *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel) SetDisplayDate(v string) *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel {
	s.DisplayDate = &v
	return s
}

func (s *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel) SetBizId(v string) *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel {
	s.BizId = &v
	return s
}

func (s *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel) SetStartDate(v string) *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel {
	s.StartDate = &v
	return s
}

func (s *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel) SetTitle(v string) *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel {
	s.Title = &v
	return s
}

func (s *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel) SetExtInfo(v map[string]interface{}) *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel {
	s.ExtInfo = v
	return s
}

func (s *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel) SetId(v int64) *ListItemActivitiesResponseBodyLmItemActivityModelListLmItemActivityModelLmActivityPopModel {
	s.Id = &v
	return s
}

type ListItemActivitiesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListItemActivitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListItemActivitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListItemActivitiesResponse) GoString() string {
	return s.String()
}

func (s *ListItemActivitiesResponse) SetHeaders(v map[string]*string) *ListItemActivitiesResponse {
	s.Headers = v
	return s
}

func (s *ListItemActivitiesResponse) SetBody(v *ListItemActivitiesResponseBody) *ListItemActivitiesResponse {
	s.Body = v
	return s
}

type ModifyBasicAndBizItemsRequest struct {
	BizId    *string                                  `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SubBizId *string                                  `json:"SubBizId,omitempty" xml:"SubBizId,omitempty"`
	ItemList []*ModifyBasicAndBizItemsRequestItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
}

func (s ModifyBasicAndBizItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBasicAndBizItemsRequest) GoString() string {
	return s.String()
}

func (s *ModifyBasicAndBizItemsRequest) SetBizId(v string) *ModifyBasicAndBizItemsRequest {
	s.BizId = &v
	return s
}

func (s *ModifyBasicAndBizItemsRequest) SetSubBizId(v string) *ModifyBasicAndBizItemsRequest {
	s.SubBizId = &v
	return s
}

func (s *ModifyBasicAndBizItemsRequest) SetItemList(v []*ModifyBasicAndBizItemsRequestItemList) *ModifyBasicAndBizItemsRequest {
	s.ItemList = v
	return s
}

type ModifyBasicAndBizItemsRequestItemList struct {
	LmItemId *string                                         `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId   *int64                                          `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	SkuList  []*ModifyBasicAndBizItemsRequestItemListSkuList `json:"SkuList,omitempty" xml:"SkuList,omitempty" type:"Repeated"`
}

func (s ModifyBasicAndBizItemsRequestItemList) String() string {
	return tea.Prettify(s)
}

func (s ModifyBasicAndBizItemsRequestItemList) GoString() string {
	return s.String()
}

func (s *ModifyBasicAndBizItemsRequestItemList) SetLmItemId(v string) *ModifyBasicAndBizItemsRequestItemList {
	s.LmItemId = &v
	return s
}

func (s *ModifyBasicAndBizItemsRequestItemList) SetItemId(v int64) *ModifyBasicAndBizItemsRequestItemList {
	s.ItemId = &v
	return s
}

func (s *ModifyBasicAndBizItemsRequestItemList) SetSkuList(v []*ModifyBasicAndBizItemsRequestItemListSkuList) *ModifyBasicAndBizItemsRequestItemList {
	s.SkuList = v
	return s
}

type ModifyBasicAndBizItemsRequestItemListSkuList struct {
	StatusAction  *int64  `json:"StatusAction,omitempty" xml:"StatusAction,omitempty"`
	PriceCent     *int64  `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	SupplierPrice *int64  `json:"SupplierPrice,omitempty" xml:"SupplierPrice,omitempty"`
	SkuId         *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Points        *int64  `json:"Points,omitempty" xml:"Points,omitempty"`
	PointsAmount  *int64  `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	BenefitId     *string `json:"BenefitId,omitempty" xml:"BenefitId,omitempty"`
	Quantity      *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s ModifyBasicAndBizItemsRequestItemListSkuList) String() string {
	return tea.Prettify(s)
}

func (s ModifyBasicAndBizItemsRequestItemListSkuList) GoString() string {
	return s.String()
}

func (s *ModifyBasicAndBizItemsRequestItemListSkuList) SetStatusAction(v int64) *ModifyBasicAndBizItemsRequestItemListSkuList {
	s.StatusAction = &v
	return s
}

func (s *ModifyBasicAndBizItemsRequestItemListSkuList) SetPriceCent(v int64) *ModifyBasicAndBizItemsRequestItemListSkuList {
	s.PriceCent = &v
	return s
}

func (s *ModifyBasicAndBizItemsRequestItemListSkuList) SetSupplierPrice(v int64) *ModifyBasicAndBizItemsRequestItemListSkuList {
	s.SupplierPrice = &v
	return s
}

func (s *ModifyBasicAndBizItemsRequestItemListSkuList) SetSkuId(v int64) *ModifyBasicAndBizItemsRequestItemListSkuList {
	s.SkuId = &v
	return s
}

func (s *ModifyBasicAndBizItemsRequestItemListSkuList) SetPoints(v int64) *ModifyBasicAndBizItemsRequestItemListSkuList {
	s.Points = &v
	return s
}

func (s *ModifyBasicAndBizItemsRequestItemListSkuList) SetPointsAmount(v int64) *ModifyBasicAndBizItemsRequestItemListSkuList {
	s.PointsAmount = &v
	return s
}

func (s *ModifyBasicAndBizItemsRequestItemListSkuList) SetBenefitId(v string) *ModifyBasicAndBizItemsRequestItemListSkuList {
	s.BenefitId = &v
	return s
}

func (s *ModifyBasicAndBizItemsRequestItemListSkuList) SetQuantity(v int32) *ModifyBasicAndBizItemsRequestItemListSkuList {
	s.Quantity = &v
	return s
}

type ModifyBasicAndBizItemsResponseBody struct {
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FailedItemList *ModifyBasicAndBizItemsResponseBodyFailedItemList `json:"FailedItemList,omitempty" xml:"FailedItemList,omitempty" type:"Struct"`
}

func (s ModifyBasicAndBizItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBasicAndBizItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBasicAndBizItemsResponseBody) SetCode(v string) *ModifyBasicAndBizItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyBasicAndBizItemsResponseBody) SetMessage(v string) *ModifyBasicAndBizItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyBasicAndBizItemsResponseBody) SetRequestId(v string) *ModifyBasicAndBizItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBasicAndBizItemsResponseBody) SetFailedItemList(v *ModifyBasicAndBizItemsResponseBodyFailedItemList) *ModifyBasicAndBizItemsResponseBody {
	s.FailedItemList = v
	return s
}

type ModifyBasicAndBizItemsResponseBodyFailedItemList struct {
	Item []*ModifyBasicAndBizItemsResponseBodyFailedItemListItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s ModifyBasicAndBizItemsResponseBodyFailedItemList) String() string {
	return tea.Prettify(s)
}

func (s ModifyBasicAndBizItemsResponseBodyFailedItemList) GoString() string {
	return s.String()
}

func (s *ModifyBasicAndBizItemsResponseBodyFailedItemList) SetItem(v []*ModifyBasicAndBizItemsResponseBodyFailedItemListItem) *ModifyBasicAndBizItemsResponseBodyFailedItemList {
	s.Item = v
	return s
}

type ModifyBasicAndBizItemsResponseBodyFailedItemListItem struct {
	LmItemId  *string                                                        `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId    *int64                                                         `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	SkuIdList *ModifyBasicAndBizItemsResponseBodyFailedItemListItemSkuIdList `json:"SkuIdList,omitempty" xml:"SkuIdList,omitempty" type:"Struct"`
}

func (s ModifyBasicAndBizItemsResponseBodyFailedItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ModifyBasicAndBizItemsResponseBodyFailedItemListItem) GoString() string {
	return s.String()
}

func (s *ModifyBasicAndBizItemsResponseBodyFailedItemListItem) SetLmItemId(v string) *ModifyBasicAndBizItemsResponseBodyFailedItemListItem {
	s.LmItemId = &v
	return s
}

func (s *ModifyBasicAndBizItemsResponseBodyFailedItemListItem) SetItemId(v int64) *ModifyBasicAndBizItemsResponseBodyFailedItemListItem {
	s.ItemId = &v
	return s
}

func (s *ModifyBasicAndBizItemsResponseBodyFailedItemListItem) SetSkuIdList(v *ModifyBasicAndBizItemsResponseBodyFailedItemListItemSkuIdList) *ModifyBasicAndBizItemsResponseBodyFailedItemListItem {
	s.SkuIdList = v
	return s
}

type ModifyBasicAndBizItemsResponseBodyFailedItemListItemSkuIdList struct {
	Sku []*string `json:"Sku,omitempty" xml:"Sku,omitempty" type:"Repeated"`
}

func (s ModifyBasicAndBizItemsResponseBodyFailedItemListItemSkuIdList) String() string {
	return tea.Prettify(s)
}

func (s ModifyBasicAndBizItemsResponseBodyFailedItemListItemSkuIdList) GoString() string {
	return s.String()
}

func (s *ModifyBasicAndBizItemsResponseBodyFailedItemListItemSkuIdList) SetSku(v []*string) *ModifyBasicAndBizItemsResponseBodyFailedItemListItemSkuIdList {
	s.Sku = v
	return s
}

type ModifyBasicAndBizItemsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBasicAndBizItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBasicAndBizItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBasicAndBizItemsResponse) GoString() string {
	return s.String()
}

func (s *ModifyBasicAndBizItemsResponse) SetHeaders(v map[string]*string) *ModifyBasicAndBizItemsResponse {
	s.Headers = v
	return s
}

func (s *ModifyBasicAndBizItemsResponse) SetBody(v *ModifyBasicAndBizItemsResponseBody) *ModifyBasicAndBizItemsResponse {
	s.Body = v
	return s
}

type ModifyBizItemsRequest struct {
	BizId    *string                          `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SubBizId *string                          `json:"SubBizId,omitempty" xml:"SubBizId,omitempty"`
	ItemList []*ModifyBizItemsRequestItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
}

func (s ModifyBizItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBizItemsRequest) GoString() string {
	return s.String()
}

func (s *ModifyBizItemsRequest) SetBizId(v string) *ModifyBizItemsRequest {
	s.BizId = &v
	return s
}

func (s *ModifyBizItemsRequest) SetSubBizId(v string) *ModifyBizItemsRequest {
	s.SubBizId = &v
	return s
}

func (s *ModifyBizItemsRequest) SetItemList(v []*ModifyBizItemsRequestItemList) *ModifyBizItemsRequest {
	s.ItemList = v
	return s
}

type ModifyBizItemsRequestItemList struct {
	LmItemId *string                                 `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId   *int64                                  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	SkuList  []*ModifyBizItemsRequestItemListSkuList `json:"SkuList,omitempty" xml:"SkuList,omitempty" type:"Repeated"`
}

func (s ModifyBizItemsRequestItemList) String() string {
	return tea.Prettify(s)
}

func (s ModifyBizItemsRequestItemList) GoString() string {
	return s.String()
}

func (s *ModifyBizItemsRequestItemList) SetLmItemId(v string) *ModifyBizItemsRequestItemList {
	s.LmItemId = &v
	return s
}

func (s *ModifyBizItemsRequestItemList) SetItemId(v int64) *ModifyBizItemsRequestItemList {
	s.ItemId = &v
	return s
}

func (s *ModifyBizItemsRequestItemList) SetSkuList(v []*ModifyBizItemsRequestItemListSkuList) *ModifyBizItemsRequestItemList {
	s.SkuList = v
	return s
}

type ModifyBizItemsRequestItemListSkuList struct {
	StatusAction *int64  `json:"StatusAction,omitempty" xml:"StatusAction,omitempty"`
	PriceCent    *int64  `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	SkuId        *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Points       *int64  `json:"Points,omitempty" xml:"Points,omitempty"`
	PointsAmount *int64  `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	BenefitId    *string `json:"BenefitId,omitempty" xml:"BenefitId,omitempty"`
	Quantity     *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s ModifyBizItemsRequestItemListSkuList) String() string {
	return tea.Prettify(s)
}

func (s ModifyBizItemsRequestItemListSkuList) GoString() string {
	return s.String()
}

func (s *ModifyBizItemsRequestItemListSkuList) SetStatusAction(v int64) *ModifyBizItemsRequestItemListSkuList {
	s.StatusAction = &v
	return s
}

func (s *ModifyBizItemsRequestItemListSkuList) SetPriceCent(v int64) *ModifyBizItemsRequestItemListSkuList {
	s.PriceCent = &v
	return s
}

func (s *ModifyBizItemsRequestItemListSkuList) SetSkuId(v int64) *ModifyBizItemsRequestItemListSkuList {
	s.SkuId = &v
	return s
}

func (s *ModifyBizItemsRequestItemListSkuList) SetPoints(v int64) *ModifyBizItemsRequestItemListSkuList {
	s.Points = &v
	return s
}

func (s *ModifyBizItemsRequestItemListSkuList) SetPointsAmount(v int64) *ModifyBizItemsRequestItemListSkuList {
	s.PointsAmount = &v
	return s
}

func (s *ModifyBizItemsRequestItemListSkuList) SetBenefitId(v string) *ModifyBizItemsRequestItemListSkuList {
	s.BenefitId = &v
	return s
}

func (s *ModifyBizItemsRequestItemListSkuList) SetQuantity(v int32) *ModifyBizItemsRequestItemListSkuList {
	s.Quantity = &v
	return s
}

type ModifyBizItemsResponseBody struct {
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	FailedItemList *ModifyBizItemsResponseBodyFailedItemList `json:"FailedItemList,omitempty" xml:"FailedItemList,omitempty" type:"Struct"`
}

func (s ModifyBizItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBizItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBizItemsResponseBody) SetRequestId(v string) *ModifyBizItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBizItemsResponseBody) SetCode(v string) *ModifyBizItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyBizItemsResponseBody) SetMessage(v string) *ModifyBizItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyBizItemsResponseBody) SetPageNumber(v int32) *ModifyBizItemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ModifyBizItemsResponseBody) SetPageSize(v int32) *ModifyBizItemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ModifyBizItemsResponseBody) SetTotalCount(v int32) *ModifyBizItemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ModifyBizItemsResponseBody) SetFailedItemList(v *ModifyBizItemsResponseBodyFailedItemList) *ModifyBizItemsResponseBody {
	s.FailedItemList = v
	return s
}

type ModifyBizItemsResponseBodyFailedItemList struct {
	Item []*ModifyBizItemsResponseBodyFailedItemListItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s ModifyBizItemsResponseBodyFailedItemList) String() string {
	return tea.Prettify(s)
}

func (s ModifyBizItemsResponseBodyFailedItemList) GoString() string {
	return s.String()
}

func (s *ModifyBizItemsResponseBodyFailedItemList) SetItem(v []*ModifyBizItemsResponseBodyFailedItemListItem) *ModifyBizItemsResponseBodyFailedItemList {
	s.Item = v
	return s
}

type ModifyBizItemsResponseBodyFailedItemListItem struct {
	LmItemId  *string                                                `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId    *int64                                                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	SkuIdList *ModifyBizItemsResponseBodyFailedItemListItemSkuIdList `json:"SkuIdList,omitempty" xml:"SkuIdList,omitempty" type:"Struct"`
}

func (s ModifyBizItemsResponseBodyFailedItemListItem) String() string {
	return tea.Prettify(s)
}

func (s ModifyBizItemsResponseBodyFailedItemListItem) GoString() string {
	return s.String()
}

func (s *ModifyBizItemsResponseBodyFailedItemListItem) SetLmItemId(v string) *ModifyBizItemsResponseBodyFailedItemListItem {
	s.LmItemId = &v
	return s
}

func (s *ModifyBizItemsResponseBodyFailedItemListItem) SetItemId(v int64) *ModifyBizItemsResponseBodyFailedItemListItem {
	s.ItemId = &v
	return s
}

func (s *ModifyBizItemsResponseBodyFailedItemListItem) SetSkuIdList(v *ModifyBizItemsResponseBodyFailedItemListItemSkuIdList) *ModifyBizItemsResponseBodyFailedItemListItem {
	s.SkuIdList = v
	return s
}

type ModifyBizItemsResponseBodyFailedItemListItemSkuIdList struct {
	Sku []*string `json:"Sku,omitempty" xml:"Sku,omitempty" type:"Repeated"`
}

func (s ModifyBizItemsResponseBodyFailedItemListItemSkuIdList) String() string {
	return tea.Prettify(s)
}

func (s ModifyBizItemsResponseBodyFailedItemListItemSkuIdList) GoString() string {
	return s.String()
}

func (s *ModifyBizItemsResponseBodyFailedItemListItemSkuIdList) SetSku(v []*string) *ModifyBizItemsResponseBodyFailedItemListItemSkuIdList {
	s.Sku = v
	return s
}

type ModifyBizItemsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBizItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBizItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBizItemsResponse) GoString() string {
	return s.String()
}

func (s *ModifyBizItemsResponse) SetHeaders(v map[string]*string) *ModifyBizItemsResponse {
	s.Headers = v
	return s
}

func (s *ModifyBizItemsResponse) SetBody(v *ModifyBizItemsResponseBody) *ModifyBizItemsResponse {
	s.Body = v
	return s
}

type ModifyItemLimitRuleRequest struct {
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SubBizCode   *string `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	LmActivityId *int64  `json:"LmActivityId,omitempty" xml:"LmActivityId,omitempty"`
	LmItemId     *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId       *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	UpperNum     *int32  `json:"UpperNum,omitempty" xml:"UpperNum,omitempty"`
	RuleId       *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleType     *int32  `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	BeginTime    *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ModifyItemLimitRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyItemLimitRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyItemLimitRuleRequest) SetBizId(v string) *ModifyItemLimitRuleRequest {
	s.BizId = &v
	return s
}

func (s *ModifyItemLimitRuleRequest) SetSubBizCode(v string) *ModifyItemLimitRuleRequest {
	s.SubBizCode = &v
	return s
}

func (s *ModifyItemLimitRuleRequest) SetLmActivityId(v int64) *ModifyItemLimitRuleRequest {
	s.LmActivityId = &v
	return s
}

func (s *ModifyItemLimitRuleRequest) SetLmItemId(v string) *ModifyItemLimitRuleRequest {
	s.LmItemId = &v
	return s
}

func (s *ModifyItemLimitRuleRequest) SetItemId(v int64) *ModifyItemLimitRuleRequest {
	s.ItemId = &v
	return s
}

func (s *ModifyItemLimitRuleRequest) SetUpperNum(v int32) *ModifyItemLimitRuleRequest {
	s.UpperNum = &v
	return s
}

func (s *ModifyItemLimitRuleRequest) SetRuleId(v int64) *ModifyItemLimitRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyItemLimitRuleRequest) SetRuleType(v int32) *ModifyItemLimitRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyItemLimitRuleRequest) SetBeginTime(v int64) *ModifyItemLimitRuleRequest {
	s.BeginTime = &v
	return s
}

func (s *ModifyItemLimitRuleRequest) SetEndTime(v int64) *ModifyItemLimitRuleRequest {
	s.EndTime = &v
	return s
}

type ModifyItemLimitRuleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Model     *bool   `json:"Model,omitempty" xml:"Model,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyItemLimitRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyItemLimitRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyItemLimitRuleResponseBody) SetCode(v string) *ModifyItemLimitRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyItemLimitRuleResponseBody) SetModel(v bool) *ModifyItemLimitRuleResponseBody {
	s.Model = &v
	return s
}

func (s *ModifyItemLimitRuleResponseBody) SetMessage(v string) *ModifyItemLimitRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyItemLimitRuleResponseBody) SetRequestId(v string) *ModifyItemLimitRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyItemLimitRuleResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyItemLimitRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyItemLimitRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyItemLimitRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyItemLimitRuleResponse) SetHeaders(v map[string]*string) *ModifyItemLimitRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyItemLimitRuleResponse) SetBody(v *ModifyItemLimitRuleResponseBody) *ModifyItemLimitRuleResponse {
	s.Body = v
	return s
}

type NotifyPayOrderStatusRequest struct {
	ChannelId     *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OperationDate *string `json:"OperationDate,omitempty" xml:"OperationDate,omitempty"`
	PayTypes      *string `json:"PayTypes,omitempty" xml:"PayTypes,omitempty"`
	Amount        *int64  `json:"Amount,omitempty" xml:"Amount,omitempty"`
}

func (s NotifyPayOrderStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayOrderStatusRequest) GoString() string {
	return s.String()
}

func (s *NotifyPayOrderStatusRequest) SetChannelId(v string) *NotifyPayOrderStatusRequest {
	s.ChannelId = &v
	return s
}

func (s *NotifyPayOrderStatusRequest) SetRequestId(v string) *NotifyPayOrderStatusRequest {
	s.RequestId = &v
	return s
}

func (s *NotifyPayOrderStatusRequest) SetOperationDate(v string) *NotifyPayOrderStatusRequest {
	s.OperationDate = &v
	return s
}

func (s *NotifyPayOrderStatusRequest) SetPayTypes(v string) *NotifyPayOrderStatusRequest {
	s.PayTypes = &v
	return s
}

func (s *NotifyPayOrderStatusRequest) SetAmount(v int64) *NotifyPayOrderStatusRequest {
	s.Amount = &v
	return s
}

type NotifyPayOrderStatusResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s NotifyPayOrderStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayOrderStatusResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyPayOrderStatusResponseBody) SetCode(v string) *NotifyPayOrderStatusResponseBody {
	s.Code = &v
	return s
}

func (s *NotifyPayOrderStatusResponseBody) SetMessage(v string) *NotifyPayOrderStatusResponseBody {
	s.Message = &v
	return s
}

type NotifyPayOrderStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *NotifyPayOrderStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s NotifyPayOrderStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyPayOrderStatusResponse) GoString() string {
	return s.String()
}

func (s *NotifyPayOrderStatusResponse) SetHeaders(v map[string]*string) *NotifyPayOrderStatusResponse {
	s.Headers = v
	return s
}

func (s *NotifyPayOrderStatusResponse) SetBody(v *NotifyPayOrderStatusResponseBody) *NotifyPayOrderStatusResponse {
	s.Body = v
	return s
}

type NotifyWithholdFundRequest struct {
	ChannelId     *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OperationDate *string `json:"OperationDate,omitempty" xml:"OperationDate,omitempty"`
	PayTypes      *string `json:"PayTypes,omitempty" xml:"PayTypes,omitempty"`
	Amount        *int64  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	TenantOrderId *string `json:"TenantOrderId,omitempty" xml:"TenantOrderId,omitempty"`
}

func (s NotifyWithholdFundRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyWithholdFundRequest) GoString() string {
	return s.String()
}

func (s *NotifyWithholdFundRequest) SetChannelId(v string) *NotifyWithholdFundRequest {
	s.ChannelId = &v
	return s
}

func (s *NotifyWithholdFundRequest) SetRequestId(v string) *NotifyWithholdFundRequest {
	s.RequestId = &v
	return s
}

func (s *NotifyWithholdFundRequest) SetOperationDate(v string) *NotifyWithholdFundRequest {
	s.OperationDate = &v
	return s
}

func (s *NotifyWithholdFundRequest) SetPayTypes(v string) *NotifyWithholdFundRequest {
	s.PayTypes = &v
	return s
}

func (s *NotifyWithholdFundRequest) SetAmount(v int64) *NotifyWithholdFundRequest {
	s.Amount = &v
	return s
}

func (s *NotifyWithholdFundRequest) SetTenantOrderId(v string) *NotifyWithholdFundRequest {
	s.TenantOrderId = &v
	return s
}

type NotifyWithholdFundResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NotifyWithholdFundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyWithholdFundResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyWithholdFundResponseBody) SetCode(v string) *NotifyWithholdFundResponseBody {
	s.Code = &v
	return s
}

func (s *NotifyWithholdFundResponseBody) SetMessage(v string) *NotifyWithholdFundResponseBody {
	s.Message = &v
	return s
}

func (s *NotifyWithholdFundResponseBody) SetRequestId(v string) *NotifyWithholdFundResponseBody {
	s.RequestId = &v
	return s
}

type NotifyWithholdFundResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *NotifyWithholdFundResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s NotifyWithholdFundResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyWithholdFundResponse) GoString() string {
	return s.String()
}

func (s *NotifyWithholdFundResponse) SetHeaders(v map[string]*string) *NotifyWithholdFundResponse {
	s.Headers = v
	return s
}

func (s *NotifyWithholdFundResponse) SetBody(v *NotifyWithholdFundResponseBody) *NotifyWithholdFundResponse {
	s.Body = v
	return s
}

type QueryActivityItemsRequest struct {
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	LmActivityId *int64  `json:"LmActivityId,omitempty" xml:"LmActivityId,omitempty"`
}

func (s QueryActivityItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsRequest) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsRequest) SetBizId(v string) *QueryActivityItemsRequest {
	s.BizId = &v
	return s
}

func (s *QueryActivityItemsRequest) SetPageNumber(v int32) *QueryActivityItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryActivityItemsRequest) SetPageSize(v int32) *QueryActivityItemsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryActivityItemsRequest) SetLmActivityId(v int64) *QueryActivityItemsRequest {
	s.LmActivityId = &v
	return s
}

type QueryActivityItemsResponseBody struct {
	RequestId               *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code                    *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                 *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber              *int64                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                *int64                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount              *int64                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LmActivityItemModelList *QueryActivityItemsResponseBodyLmActivityItemModelList `json:"LmActivityItemModelList,omitempty" xml:"LmActivityItemModelList,omitempty" type:"Struct"`
}

func (s QueryActivityItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponseBody) SetRequestId(v string) *QueryActivityItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryActivityItemsResponseBody) SetCode(v string) *QueryActivityItemsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryActivityItemsResponseBody) SetMessage(v string) *QueryActivityItemsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryActivityItemsResponseBody) SetPageNumber(v int64) *QueryActivityItemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryActivityItemsResponseBody) SetPageSize(v int64) *QueryActivityItemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryActivityItemsResponseBody) SetTotalCount(v int64) *QueryActivityItemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryActivityItemsResponseBody) SetLmActivityItemModelList(v *QueryActivityItemsResponseBodyLmActivityItemModelList) *QueryActivityItemsResponseBody {
	s.LmActivityItemModelList = v
	return s
}

type QueryActivityItemsResponseBodyLmActivityItemModelList struct {
	LmActivityItemModel []*QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel `json:"LmActivityItemModel,omitempty" xml:"LmActivityItemModel,omitempty" type:"Repeated"`
}

func (s QueryActivityItemsResponseBodyLmActivityItemModelList) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponseBodyLmActivityItemModelList) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelList) SetLmActivityItemModel(v []*QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) *QueryActivityItemsResponseBodyLmActivityItemModelList {
	s.LmActivityItemModel = v
	return s
}

type QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel struct {
	ItemTitle            *string                                                                               `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	TbSellerId           *int64                                                                                `json:"TbSellerId,omitempty" xml:"TbSellerId,omitempty"`
	SellableQuantity     *int32                                                                                `json:"SellableQuantity,omitempty" xml:"SellableQuantity,omitempty"`
	TbShopId             *int64                                                                                `json:"TbShopId,omitempty" xml:"TbShopId,omitempty"`
	LmItemId             *string                                                                               `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	TbShopName           *string                                                                               `json:"TbShopName,omitempty" xml:"TbShopName,omitempty"`
	ItemActivityQuantity *int32                                                                                `json:"ItemActivityQuantity,omitempty" xml:"ItemActivityQuantity,omitempty"`
	LmActivityId         *int64                                                                                `json:"LmActivityId,omitempty" xml:"LmActivityId,omitempty"`
	CanSell              *bool                                                                                 `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	ItemId               *int64                                                                                `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	MainPicUrl           *string                                                                               `json:"MainPicUrl,omitempty" xml:"MainPicUrl,omitempty"`
	Tips                 *string                                                                               `json:"Tips,omitempty" xml:"Tips,omitempty"`
	LmShopId             *string                                                                               `json:"LmShopId,omitempty" xml:"LmShopId,omitempty"`
	ItemActivityStatus   *string                                                                               `json:"ItemActivityStatus,omitempty" xml:"ItemActivityStatus,omitempty"`
	SkuModelList         *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelList `json:"SkuModelList,omitempty" xml:"SkuModelList,omitempty" type:"Struct"`
}

func (s QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetItemTitle(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.ItemTitle = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetTbSellerId(v int64) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.TbSellerId = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetSellableQuantity(v int32) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.SellableQuantity = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetTbShopId(v int64) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.TbShopId = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetLmItemId(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.LmItemId = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetTbShopName(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.TbShopName = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetItemActivityQuantity(v int32) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.ItemActivityQuantity = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetLmActivityId(v int64) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.LmActivityId = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetCanSell(v bool) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.CanSell = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetItemId(v int64) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.ItemId = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetMainPicUrl(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.MainPicUrl = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetTips(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.Tips = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetLmShopId(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.LmShopId = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetItemActivityStatus(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.ItemActivityStatus = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel) SetSkuModelList(v *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelList) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModel {
	s.SkuModelList = v
	return s
}

type QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelList struct {
	LmActivityItemSkuModel []*QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel `json:"LmActivityItemSkuModel,omitempty" xml:"LmActivityItemSkuModel,omitempty" type:"Repeated"`
}

func (s QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelList) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelList) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelList) SetLmActivityItemSkuModel(v []*QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelList {
	s.LmActivityItemSkuModel = v
	return s
}

type QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel struct {
	ReservedPrice  *int64  `json:"ReservedPrice,omitempty" xml:"ReservedPrice,omitempty"`
	Tips           *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	LmItemId       *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	SkuId          *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	LmActivityId   *int64  `json:"LmActivityId,omitempty" xml:"LmActivityId,omitempty"`
	ActivityStatus *string `json:"ActivityStatus,omitempty" xml:"ActivityStatus,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ItemId         *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	SkuPicUrl      *string `json:"SkuPicUrl,omitempty" xml:"SkuPicUrl,omitempty"`
	ActivityPrice  *int64  `json:"ActivityPrice,omitempty" xml:"ActivityPrice,omitempty"`
	SkuTitle       *string `json:"SkuTitle,omitempty" xml:"SkuTitle,omitempty"`
}

func (s QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) SetReservedPrice(v int64) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel {
	s.ReservedPrice = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) SetTips(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel {
	s.Tips = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) SetLmItemId(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel {
	s.LmItemId = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) SetSkuId(v int64) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel {
	s.SkuId = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) SetLmActivityId(v int64) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel {
	s.LmActivityId = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) SetActivityStatus(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel {
	s.ActivityStatus = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) SetBizId(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel {
	s.BizId = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) SetItemId(v int64) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel {
	s.ItemId = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) SetSkuPicUrl(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel {
	s.SkuPicUrl = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) SetActivityPrice(v int64) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel {
	s.ActivityPrice = &v
	return s
}

func (s *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel) SetSkuTitle(v string) *QueryActivityItemsResponseBodyLmActivityItemModelListLmActivityItemModelSkuModelListLmActivityItemSkuModel {
	s.SkuTitle = &v
	return s
}

type QueryActivityItemsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryActivityItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryActivityItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryActivityItemsResponse) GoString() string {
	return s.String()
}

func (s *QueryActivityItemsResponse) SetHeaders(v map[string]*string) *QueryActivityItemsResponse {
	s.Headers = v
	return s
}

func (s *QueryActivityItemsResponse) SetBody(v *QueryActivityItemsResponseBody) *QueryActivityItemsResponse {
	s.Body = v
	return s
}

type QueryAddressRequest struct {
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Ip           *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	DivisionCode *string `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
}

func (s QueryAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressRequest) GoString() string {
	return s.String()
}

func (s *QueryAddressRequest) SetBizId(v string) *QueryAddressRequest {
	s.BizId = &v
	return s
}

func (s *QueryAddressRequest) SetIp(v string) *QueryAddressRequest {
	s.Ip = &v
	return s
}

func (s *QueryAddressRequest) SetDivisionCode(v string) *QueryAddressRequest {
	s.DivisionCode = &v
	return s
}

type QueryAddressResponseBody struct {
	Code            *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DivisionAddress *QueryAddressResponseBodyDivisionAddress `json:"DivisionAddress,omitempty" xml:"DivisionAddress,omitempty" type:"Struct"`
}

func (s QueryAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAddressResponseBody) SetCode(v string) *QueryAddressResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAddressResponseBody) SetMessage(v string) *QueryAddressResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAddressResponseBody) SetRequestId(v string) *QueryAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAddressResponseBody) SetDivisionAddress(v *QueryAddressResponseBodyDivisionAddress) *QueryAddressResponseBody {
	s.DivisionAddress = v
	return s
}

type QueryAddressResponseBodyDivisionAddress struct {
	DivisionAddress []*QueryAddressResponseBodyDivisionAddressDivisionAddress `json:"DivisionAddress,omitempty" xml:"DivisionAddress,omitempty" type:"Repeated"`
}

func (s QueryAddressResponseBodyDivisionAddress) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressResponseBodyDivisionAddress) GoString() string {
	return s.String()
}

func (s *QueryAddressResponseBodyDivisionAddress) SetDivisionAddress(v []*QueryAddressResponseBodyDivisionAddressDivisionAddress) *QueryAddressResponseBodyDivisionAddress {
	s.DivisionAddress = v
	return s
}

type QueryAddressResponseBodyDivisionAddressDivisionAddress struct {
	ParentId      *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	DivisionCode  *int64  `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
	DivisionLevel *int32  `json:"DivisionLevel,omitempty" xml:"DivisionLevel,omitempty"`
	DivisionName  *string `json:"DivisionName,omitempty" xml:"DivisionName,omitempty"`
}

func (s QueryAddressResponseBodyDivisionAddressDivisionAddress) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressResponseBodyDivisionAddressDivisionAddress) GoString() string {
	return s.String()
}

func (s *QueryAddressResponseBodyDivisionAddressDivisionAddress) SetParentId(v int64) *QueryAddressResponseBodyDivisionAddressDivisionAddress {
	s.ParentId = &v
	return s
}

func (s *QueryAddressResponseBodyDivisionAddressDivisionAddress) SetDivisionCode(v int64) *QueryAddressResponseBodyDivisionAddressDivisionAddress {
	s.DivisionCode = &v
	return s
}

func (s *QueryAddressResponseBodyDivisionAddressDivisionAddress) SetDivisionLevel(v int32) *QueryAddressResponseBodyDivisionAddressDivisionAddress {
	s.DivisionLevel = &v
	return s
}

func (s *QueryAddressResponseBodyDivisionAddressDivisionAddress) SetDivisionName(v string) *QueryAddressResponseBodyDivisionAddressDivisionAddress {
	s.DivisionName = &v
	return s
}

type QueryAddressResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressResponse) GoString() string {
	return s.String()
}

func (s *QueryAddressResponse) SetHeaders(v map[string]*string) *QueryAddressResponse {
	s.Headers = v
	return s
}

func (s *QueryAddressResponse) SetBody(v *QueryAddressResponseBody) *QueryAddressResponse {
	s.Body = v
	return s
}

type QueryAddressDetailRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	AddressInfo           *string `json:"AddressInfo,omitempty" xml:"AddressInfo,omitempty"`
}

func (s QueryAddressDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryAddressDetailRequest) SetBizId(v string) *QueryAddressDetailRequest {
	s.BizId = &v
	return s
}

func (s *QueryAddressDetailRequest) SetThirdPartyUserId(v string) *QueryAddressDetailRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *QueryAddressDetailRequest) SetUseAnonymousTbAccount(v bool) *QueryAddressDetailRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *QueryAddressDetailRequest) SetAddressInfo(v string) *QueryAddressDetailRequest {
	s.AddressInfo = &v
	return s
}

type QueryAddressDetailResponseBody struct {
	Code            *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeliveryAddress *QueryAddressDetailResponseBodyDeliveryAddress `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty" type:"Struct"`
}

func (s QueryAddressDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAddressDetailResponseBody) SetCode(v string) *QueryAddressDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAddressDetailResponseBody) SetMessage(v string) *QueryAddressDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAddressDetailResponseBody) SetRequestId(v string) *QueryAddressDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAddressDetailResponseBody) SetDeliveryAddress(v *QueryAddressDetailResponseBodyDeliveryAddress) *QueryAddressDetailResponseBody {
	s.DeliveryAddress = v
	return s
}

type QueryAddressDetailResponseBodyDeliveryAddress struct {
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	PostCode         *string `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	FullName         *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	AddressDetail    *string `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	City             *string `json:"City,omitempty" xml:"City,omitempty"`
	AddressId        *int64  `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	DivisionCode     *string `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
	Mobile           *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Country          *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Area             *string `json:"Area,omitempty" xml:"Area,omitempty"`
	TownDivisionCode *string `json:"TownDivisionCode,omitempty" xml:"TownDivisionCode,omitempty"`
	Town             *string `json:"Town,omitempty" xml:"Town,omitempty"`
	Province         *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s QueryAddressDetailResponseBodyDeliveryAddress) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressDetailResponseBodyDeliveryAddress) GoString() string {
	return s.String()
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetStatus(v int32) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.Status = &v
	return s
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetPostCode(v string) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.PostCode = &v
	return s
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetFullName(v string) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.FullName = &v
	return s
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetAddressDetail(v string) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.AddressDetail = &v
	return s
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetCity(v string) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.City = &v
	return s
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetAddressId(v int64) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.AddressId = &v
	return s
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetDivisionCode(v string) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.DivisionCode = &v
	return s
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetMobile(v string) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.Mobile = &v
	return s
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetCountry(v string) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.Country = &v
	return s
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetArea(v string) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.Area = &v
	return s
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetTownDivisionCode(v string) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.TownDivisionCode = &v
	return s
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetTown(v string) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.Town = &v
	return s
}

func (s *QueryAddressDetailResponseBodyDeliveryAddress) SetProvince(v string) *QueryAddressDetailResponseBodyDeliveryAddress {
	s.Province = &v
	return s
}

type QueryAddressDetailResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAddressDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAddressDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryAddressDetailResponse) SetHeaders(v map[string]*string) *QueryAddressDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryAddressDetailResponse) SetBody(v *QueryAddressDetailResponseBody) *QueryAddressDetailResponse {
	s.Body = v
	return s
}

type QueryAddressListRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
}

func (s QueryAddressListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressListRequest) GoString() string {
	return s.String()
}

func (s *QueryAddressListRequest) SetBizId(v string) *QueryAddressListRequest {
	s.BizId = &v
	return s
}

func (s *QueryAddressListRequest) SetThirdPartyUserId(v string) *QueryAddressListRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *QueryAddressListRequest) SetUseAnonymousTbAccount(v bool) *QueryAddressListRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

type QueryAddressListResponseBody struct {
	Code        *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AddressList []*QueryAddressListResponseBodyAddressList `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
}

func (s QueryAddressListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAddressListResponseBody) SetCode(v string) *QueryAddressListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAddressListResponseBody) SetMessage(v string) *QueryAddressListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAddressListResponseBody) SetRequestId(v string) *QueryAddressListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAddressListResponseBody) SetAddressList(v []*QueryAddressListResponseBodyAddressList) *QueryAddressListResponseBody {
	s.AddressList = v
	return s
}

type QueryAddressListResponseBodyAddressList struct {
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	PostCode         *string `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	FullName         *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	AddressDetail    *string `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	City             *string `json:"City,omitempty" xml:"City,omitempty"`
	AddressId        *int64  `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	DivisionCode     *string `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
	Mobile           *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Country          *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Area             *string `json:"Area,omitempty" xml:"Area,omitempty"`
	TownDivisionCode *string `json:"TownDivisionCode,omitempty" xml:"TownDivisionCode,omitempty"`
	Town             *string `json:"Town,omitempty" xml:"Town,omitempty"`
	Province         *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s QueryAddressListResponseBodyAddressList) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressListResponseBodyAddressList) GoString() string {
	return s.String()
}

func (s *QueryAddressListResponseBodyAddressList) SetStatus(v int32) *QueryAddressListResponseBodyAddressList {
	s.Status = &v
	return s
}

func (s *QueryAddressListResponseBodyAddressList) SetPostCode(v string) *QueryAddressListResponseBodyAddressList {
	s.PostCode = &v
	return s
}

func (s *QueryAddressListResponseBodyAddressList) SetFullName(v string) *QueryAddressListResponseBodyAddressList {
	s.FullName = &v
	return s
}

func (s *QueryAddressListResponseBodyAddressList) SetAddressDetail(v string) *QueryAddressListResponseBodyAddressList {
	s.AddressDetail = &v
	return s
}

func (s *QueryAddressListResponseBodyAddressList) SetCity(v string) *QueryAddressListResponseBodyAddressList {
	s.City = &v
	return s
}

func (s *QueryAddressListResponseBodyAddressList) SetAddressId(v int64) *QueryAddressListResponseBodyAddressList {
	s.AddressId = &v
	return s
}

func (s *QueryAddressListResponseBodyAddressList) SetDivisionCode(v string) *QueryAddressListResponseBodyAddressList {
	s.DivisionCode = &v
	return s
}

func (s *QueryAddressListResponseBodyAddressList) SetMobile(v string) *QueryAddressListResponseBodyAddressList {
	s.Mobile = &v
	return s
}

func (s *QueryAddressListResponseBodyAddressList) SetCountry(v string) *QueryAddressListResponseBodyAddressList {
	s.Country = &v
	return s
}

func (s *QueryAddressListResponseBodyAddressList) SetArea(v string) *QueryAddressListResponseBodyAddressList {
	s.Area = &v
	return s
}

func (s *QueryAddressListResponseBodyAddressList) SetTownDivisionCode(v string) *QueryAddressListResponseBodyAddressList {
	s.TownDivisionCode = &v
	return s
}

func (s *QueryAddressListResponseBodyAddressList) SetTown(v string) *QueryAddressListResponseBodyAddressList {
	s.Town = &v
	return s
}

func (s *QueryAddressListResponseBodyAddressList) SetProvince(v string) *QueryAddressListResponseBodyAddressList {
	s.Province = &v
	return s
}

type QueryAddressListResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAddressListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAddressListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAddressListResponse) GoString() string {
	return s.String()
}

func (s *QueryAddressListResponse) SetHeaders(v map[string]*string) *QueryAddressListResponse {
	s.Headers = v
	return s
}

func (s *QueryAddressListResponse) SetBody(v *QueryAddressListResponseBody) *QueryAddressListResponse {
	s.Body = v
	return s
}

type QueryAdvertisementSettleInfoRequest struct {
	MediaSettleDetailId *string `json:"MediaSettleDetailId,omitempty" xml:"MediaSettleDetailId,omitempty"`
	ChannelId           *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	SettleNo            *string `json:"SettleNo,omitempty" xml:"SettleNo,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime             *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber          *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ExtInfo             *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s QueryAdvertisementSettleInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvertisementSettleInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryAdvertisementSettleInfoRequest) SetMediaSettleDetailId(v string) *QueryAdvertisementSettleInfoRequest {
	s.MediaSettleDetailId = &v
	return s
}

func (s *QueryAdvertisementSettleInfoRequest) SetChannelId(v string) *QueryAdvertisementSettleInfoRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryAdvertisementSettleInfoRequest) SetSettleNo(v string) *QueryAdvertisementSettleInfoRequest {
	s.SettleNo = &v
	return s
}

func (s *QueryAdvertisementSettleInfoRequest) SetStartTime(v string) *QueryAdvertisementSettleInfoRequest {
	s.StartTime = &v
	return s
}

func (s *QueryAdvertisementSettleInfoRequest) SetEndTime(v string) *QueryAdvertisementSettleInfoRequest {
	s.EndTime = &v
	return s
}

func (s *QueryAdvertisementSettleInfoRequest) SetPageSize(v int32) *QueryAdvertisementSettleInfoRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAdvertisementSettleInfoRequest) SetPageNumber(v int32) *QueryAdvertisementSettleInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAdvertisementSettleInfoRequest) SetExtInfo(v string) *QueryAdvertisementSettleInfoRequest {
	s.ExtInfo = &v
	return s
}

type QueryAdvertisementSettleInfoResponseBody struct {
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                                        `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                                        `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                                        `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Model      *QueryAdvertisementSettleInfoResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s QueryAdvertisementSettleInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvertisementSettleInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAdvertisementSettleInfoResponseBody) SetRequestId(v string) *QueryAdvertisementSettleInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBody) SetSuccess(v bool) *QueryAdvertisementSettleInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBody) SetSubMessage(v string) *QueryAdvertisementSettleInfoResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBody) SetCode(v string) *QueryAdvertisementSettleInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBody) SetMessage(v string) *QueryAdvertisementSettleInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBody) SetSubCode(v string) *QueryAdvertisementSettleInfoResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBody) SetTotalCount(v int64) *QueryAdvertisementSettleInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBody) SetLogsId(v string) *QueryAdvertisementSettleInfoResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBody) SetModel(v *QueryAdvertisementSettleInfoResponseBodyModel) *QueryAdvertisementSettleInfoResponseBody {
	s.Model = v
	return s
}

type QueryAdvertisementSettleInfoResponseBodyModel struct {
	PageSize                *int32                                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber              *int32                                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount              *int32                                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	AdvertiseSettleInfoList []*QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList `json:"AdvertiseSettleInfoList,omitempty" xml:"AdvertiseSettleInfoList,omitempty" type:"Repeated"`
}

func (s QueryAdvertisementSettleInfoResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvertisementSettleInfoResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryAdvertisementSettleInfoResponseBodyModel) SetPageSize(v int32) *QueryAdvertisementSettleInfoResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModel) SetPageNumber(v int32) *QueryAdvertisementSettleInfoResponseBodyModel {
	s.PageNumber = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModel) SetTotalCount(v int32) *QueryAdvertisementSettleInfoResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModel) SetAdvertiseSettleInfoList(v []*QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) *QueryAdvertisementSettleInfoResponseBodyModel {
	s.AdvertiseSettleInfoList = v
	return s
}

type QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList struct {
	AdvertiseSettleDetailId *string `json:"AdvertiseSettleDetailId,omitempty" xml:"AdvertiseSettleDetailId,omitempty"`
	CreateDate              *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	MediaSettleDetailId     *string `json:"MediaSettleDetailId,omitempty" xml:"MediaSettleDetailId,omitempty"`
	ExtInfo                 *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	EndTime                 *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime               *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SettleNo                *string `json:"SettleNo,omitempty" xml:"SettleNo,omitempty"`
	AdvertiseSettleAmount   *string `json:"AdvertiseSettleAmount,omitempty" xml:"AdvertiseSettleAmount,omitempty"`
	SettleStatus            *string `json:"SettleStatus,omitempty" xml:"SettleStatus,omitempty"`
	AdvertiseName           *string `json:"AdvertiseName,omitempty" xml:"AdvertiseName,omitempty"`
	ChannelId               *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ModifiedDate            *string `json:"ModifiedDate,omitempty" xml:"ModifiedDate,omitempty"`
}

func (s QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) GoString() string {
	return s.String()
}

func (s *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) SetAdvertiseSettleDetailId(v string) *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList {
	s.AdvertiseSettleDetailId = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) SetCreateDate(v string) *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList {
	s.CreateDate = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) SetMediaSettleDetailId(v string) *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList {
	s.MediaSettleDetailId = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) SetExtInfo(v string) *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList {
	s.ExtInfo = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) SetEndTime(v string) *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList {
	s.EndTime = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) SetStartTime(v string) *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList {
	s.StartTime = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) SetSettleNo(v string) *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList {
	s.SettleNo = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) SetAdvertiseSettleAmount(v string) *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList {
	s.AdvertiseSettleAmount = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) SetSettleStatus(v string) *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList {
	s.SettleStatus = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) SetAdvertiseName(v string) *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList {
	s.AdvertiseName = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) SetChannelId(v string) *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList {
	s.ChannelId = &v
	return s
}

func (s *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList) SetModifiedDate(v string) *QueryAdvertisementSettleInfoResponseBodyModelAdvertiseSettleInfoList {
	s.ModifiedDate = &v
	return s
}

type QueryAdvertisementSettleInfoResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAdvertisementSettleInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAdvertisementSettleInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAdvertisementSettleInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryAdvertisementSettleInfoResponse) SetHeaders(v map[string]*string) *QueryAdvertisementSettleInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryAdvertisementSettleInfoResponse) SetBody(v *QueryAdvertisementSettleInfoResponseBody) *QueryAdvertisementSettleInfoResponse {
	s.Body = v
	return s
}

type QueryAgreementRequest struct {
	ExternalAgreementNo *string `json:"ExternalAgreementNo,omitempty" xml:"ExternalAgreementNo,omitempty"`
	MerchantId          *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
	AgreementNo         *string `json:"AgreementNo,omitempty" xml:"AgreementNo,omitempty"`
}

func (s QueryAgreementRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAgreementRequest) GoString() string {
	return s.String()
}

func (s *QueryAgreementRequest) SetExternalAgreementNo(v string) *QueryAgreementRequest {
	s.ExternalAgreementNo = &v
	return s
}

func (s *QueryAgreementRequest) SetMerchantId(v string) *QueryAgreementRequest {
	s.MerchantId = &v
	return s
}

func (s *QueryAgreementRequest) SetAgreementNo(v string) *QueryAgreementRequest {
	s.AgreementNo = &v
	return s
}

type QueryAgreementResponseBody struct {
	Code                   *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId              *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	QueryAgreementResponse *QueryAgreementResponseBodyQueryAgreementResponse `json:"QueryAgreementResponse,omitempty" xml:"QueryAgreementResponse,omitempty" type:"Struct"`
}

func (s QueryAgreementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAgreementResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAgreementResponseBody) SetCode(v string) *QueryAgreementResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAgreementResponseBody) SetMessage(v string) *QueryAgreementResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAgreementResponseBody) SetRequestId(v string) *QueryAgreementResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAgreementResponseBody) SetQueryAgreementResponse(v *QueryAgreementResponseBodyQueryAgreementResponse) *QueryAgreementResponseBody {
	s.QueryAgreementResponse = v
	return s
}

type QueryAgreementResponseBodyQueryAgreementResponse struct {
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ExternalAgreementNo *string `json:"ExternalAgreementNo,omitempty" xml:"ExternalAgreementNo,omitempty"`
	ValidTime           *string `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
	InvalidTime         *string `json:"InvalidTime,omitempty" xml:"InvalidTime,omitempty"`
	AgreementNo         *string `json:"AgreementNo,omitempty" xml:"AgreementNo,omitempty"`
	SignTime            *string `json:"SignTime,omitempty" xml:"SignTime,omitempty"`
}

func (s QueryAgreementResponseBodyQueryAgreementResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAgreementResponseBodyQueryAgreementResponse) GoString() string {
	return s.String()
}

func (s *QueryAgreementResponseBodyQueryAgreementResponse) SetStatus(v string) *QueryAgreementResponseBodyQueryAgreementResponse {
	s.Status = &v
	return s
}

func (s *QueryAgreementResponseBodyQueryAgreementResponse) SetExternalAgreementNo(v string) *QueryAgreementResponseBodyQueryAgreementResponse {
	s.ExternalAgreementNo = &v
	return s
}

func (s *QueryAgreementResponseBodyQueryAgreementResponse) SetValidTime(v string) *QueryAgreementResponseBodyQueryAgreementResponse {
	s.ValidTime = &v
	return s
}

func (s *QueryAgreementResponseBodyQueryAgreementResponse) SetInvalidTime(v string) *QueryAgreementResponseBodyQueryAgreementResponse {
	s.InvalidTime = &v
	return s
}

func (s *QueryAgreementResponseBodyQueryAgreementResponse) SetAgreementNo(v string) *QueryAgreementResponseBodyQueryAgreementResponse {
	s.AgreementNo = &v
	return s
}

func (s *QueryAgreementResponseBodyQueryAgreementResponse) SetSignTime(v string) *QueryAgreementResponseBodyQueryAgreementResponse {
	s.SignTime = &v
	return s
}

type QueryAgreementResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAgreementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAgreementResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAgreementResponse) GoString() string {
	return s.String()
}

func (s *QueryAgreementResponse) SetHeaders(v map[string]*string) *QueryAgreementResponse {
	s.Headers = v
	return s
}

func (s *QueryAgreementResponse) SetBody(v *QueryAgreementResponseBody) *QueryAgreementResponse {
	s.Body = v
	return s
}

type QueryAllCinemasRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CityCode   *int64  `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ExtJson    *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s QueryAllCinemasRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCinemasRequest) GoString() string {
	return s.String()
}

func (s *QueryAllCinemasRequest) SetBizId(v string) *QueryAllCinemasRequest {
	s.BizId = &v
	return s
}

func (s *QueryAllCinemasRequest) SetCityCode(v int64) *QueryAllCinemasRequest {
	s.CityCode = &v
	return s
}

func (s *QueryAllCinemasRequest) SetPageNumber(v int64) *QueryAllCinemasRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAllCinemasRequest) SetExtJson(v string) *QueryAllCinemasRequest {
	s.ExtJson = &v
	return s
}

type QueryAllCinemasResponseBody struct {
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                             `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                             `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                             `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Cinemas    *QueryAllCinemasResponseBodyCinemas `json:"Cinemas,omitempty" xml:"Cinemas,omitempty" type:"Struct"`
}

func (s QueryAllCinemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCinemasResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllCinemasResponseBody) SetRequestId(v string) *QueryAllCinemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAllCinemasResponseBody) SetSuccess(v bool) *QueryAllCinemasResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAllCinemasResponseBody) SetSubMessage(v string) *QueryAllCinemasResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryAllCinemasResponseBody) SetCode(v string) *QueryAllCinemasResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAllCinemasResponseBody) SetMessage(v string) *QueryAllCinemasResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAllCinemasResponseBody) SetSubCode(v string) *QueryAllCinemasResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryAllCinemasResponseBody) SetTotalCount(v int64) *QueryAllCinemasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryAllCinemasResponseBody) SetLogsId(v string) *QueryAllCinemasResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryAllCinemasResponseBody) SetCinemas(v *QueryAllCinemasResponseBodyCinemas) *QueryAllCinemasResponseBody {
	s.Cinemas = v
	return s
}

type QueryAllCinemasResponseBodyCinemas struct {
	Cinema []*QueryAllCinemasResponseBodyCinemasCinema `json:"Cinema,omitempty" xml:"Cinema,omitempty" type:"Repeated"`
}

func (s QueryAllCinemasResponseBodyCinemas) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCinemasResponseBodyCinemas) GoString() string {
	return s.String()
}

func (s *QueryAllCinemasResponseBodyCinemas) SetCinema(v []*QueryAllCinemasResponseBodyCinemasCinema) *QueryAllCinemasResponseBodyCinemas {
	s.Cinema = v
	return s
}

type QueryAllCinemasResponseBodyCinemasCinema struct {
	StandardId        *string `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	ScheduleCloseTime *int64  `json:"ScheduleCloseTime,omitempty" xml:"ScheduleCloseTime,omitempty"`
	CityId            *int64  `json:"CityId,omitempty" xml:"CityId,omitempty"`
	CinemaName        *string `json:"CinemaName,omitempty" xml:"CinemaName,omitempty"`
	CityName          *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	Address           *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Longitude         *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Latitude          *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Phone             *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QueryAllCinemasResponseBodyCinemasCinema) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCinemasResponseBodyCinemasCinema) GoString() string {
	return s.String()
}

func (s *QueryAllCinemasResponseBodyCinemasCinema) SetStandardId(v string) *QueryAllCinemasResponseBodyCinemasCinema {
	s.StandardId = &v
	return s
}

func (s *QueryAllCinemasResponseBodyCinemasCinema) SetScheduleCloseTime(v int64) *QueryAllCinemasResponseBodyCinemasCinema {
	s.ScheduleCloseTime = &v
	return s
}

func (s *QueryAllCinemasResponseBodyCinemasCinema) SetCityId(v int64) *QueryAllCinemasResponseBodyCinemasCinema {
	s.CityId = &v
	return s
}

func (s *QueryAllCinemasResponseBodyCinemasCinema) SetCinemaName(v string) *QueryAllCinemasResponseBodyCinemasCinema {
	s.CinemaName = &v
	return s
}

func (s *QueryAllCinemasResponseBodyCinemasCinema) SetCityName(v string) *QueryAllCinemasResponseBodyCinemasCinema {
	s.CityName = &v
	return s
}

func (s *QueryAllCinemasResponseBodyCinemasCinema) SetAddress(v string) *QueryAllCinemasResponseBodyCinemasCinema {
	s.Address = &v
	return s
}

func (s *QueryAllCinemasResponseBodyCinemasCinema) SetLongitude(v string) *QueryAllCinemasResponseBodyCinemasCinema {
	s.Longitude = &v
	return s
}

func (s *QueryAllCinemasResponseBodyCinemasCinema) SetLatitude(v string) *QueryAllCinemasResponseBodyCinemasCinema {
	s.Latitude = &v
	return s
}

func (s *QueryAllCinemasResponseBodyCinemasCinema) SetPhone(v string) *QueryAllCinemasResponseBodyCinemasCinema {
	s.Phone = &v
	return s
}

func (s *QueryAllCinemasResponseBodyCinemasCinema) SetId(v int64) *QueryAllCinemasResponseBodyCinemasCinema {
	s.Id = &v
	return s
}

type QueryAllCinemasResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllCinemasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllCinemasResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCinemasResponse) GoString() string {
	return s.String()
}

func (s *QueryAllCinemasResponse) SetHeaders(v map[string]*string) *QueryAllCinemasResponse {
	s.Headers = v
	return s
}

func (s *QueryAllCinemasResponse) SetBody(v *QueryAllCinemasResponseBody) *QueryAllCinemasResponse {
	s.Body = v
	return s
}

type QueryAllCitiesRequest struct {
	BizId   *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ExtJson map[string]interface{} `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s QueryAllCitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCitiesRequest) GoString() string {
	return s.String()
}

func (s *QueryAllCitiesRequest) SetBizId(v string) *QueryAllCitiesRequest {
	s.BizId = &v
	return s
}

func (s *QueryAllCitiesRequest) SetExtJson(v map[string]interface{}) *QueryAllCitiesRequest {
	s.ExtJson = v
	return s
}

type QueryAllCitiesShrinkRequest struct {
	BizId         *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ExtJsonShrink *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s QueryAllCitiesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCitiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryAllCitiesShrinkRequest) SetBizId(v string) *QueryAllCitiesShrinkRequest {
	s.BizId = &v
	return s
}

func (s *QueryAllCitiesShrinkRequest) SetExtJsonShrink(v string) *QueryAllCitiesShrinkRequest {
	s.ExtJsonShrink = &v
	return s
}

type QueryAllCitiesResponseBody struct {
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                           `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                           `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	LogsId     *string                           `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Cities     *QueryAllCitiesResponseBodyCities `json:"Cities,omitempty" xml:"Cities,omitempty" type:"Struct"`
}

func (s QueryAllCitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCitiesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAllCitiesResponseBody) SetRequestId(v string) *QueryAllCitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAllCitiesResponseBody) SetSuccess(v bool) *QueryAllCitiesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAllCitiesResponseBody) SetSubMessage(v string) *QueryAllCitiesResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryAllCitiesResponseBody) SetCode(v string) *QueryAllCitiesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAllCitiesResponseBody) SetMessage(v string) *QueryAllCitiesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAllCitiesResponseBody) SetSubCode(v string) *QueryAllCitiesResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryAllCitiesResponseBody) SetLogsId(v string) *QueryAllCitiesResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryAllCitiesResponseBody) SetCities(v *QueryAllCitiesResponseBodyCities) *QueryAllCitiesResponseBody {
	s.Cities = v
	return s
}

type QueryAllCitiesResponseBodyCities struct {
	City []*QueryAllCitiesResponseBodyCitiesCity `json:"City,omitempty" xml:"City,omitempty" type:"Repeated"`
}

func (s QueryAllCitiesResponseBodyCities) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCitiesResponseBodyCities) GoString() string {
	return s.String()
}

func (s *QueryAllCitiesResponseBodyCities) SetCity(v []*QueryAllCitiesResponseBodyCitiesCity) *QueryAllCitiesResponseBodyCities {
	s.City = v
	return s
}

type QueryAllCitiesResponseBodyCitiesCity struct {
	CityCode *int64  `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	ParentId *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PinYin   *string `json:"PinYin,omitempty" xml:"PinYin,omitempty"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QueryAllCitiesResponseBodyCitiesCity) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCitiesResponseBodyCitiesCity) GoString() string {
	return s.String()
}

func (s *QueryAllCitiesResponseBodyCitiesCity) SetCityCode(v int64) *QueryAllCitiesResponseBodyCitiesCity {
	s.CityCode = &v
	return s
}

func (s *QueryAllCitiesResponseBodyCitiesCity) SetParentId(v int64) *QueryAllCitiesResponseBodyCitiesCity {
	s.ParentId = &v
	return s
}

func (s *QueryAllCitiesResponseBodyCitiesCity) SetName(v string) *QueryAllCitiesResponseBodyCitiesCity {
	s.Name = &v
	return s
}

func (s *QueryAllCitiesResponseBodyCitiesCity) SetPinYin(v string) *QueryAllCitiesResponseBodyCitiesCity {
	s.PinYin = &v
	return s
}

func (s *QueryAllCitiesResponseBodyCitiesCity) SetId(v int64) *QueryAllCitiesResponseBodyCitiesCity {
	s.Id = &v
	return s
}

type QueryAllCitiesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAllCitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAllCitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAllCitiesResponse) GoString() string {
	return s.String()
}

func (s *QueryAllCitiesResponse) SetHeaders(v map[string]*string) *QueryAllCitiesResponse {
	s.Headers = v
	return s
}

func (s *QueryAllCitiesResponse) SetBody(v *QueryAllCitiesResponseBody) *QueryAllCitiesResponse {
	s.Body = v
	return s
}

type QueryBatchRegistAnonymousTbAccountResultRequest struct {
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
}

func (s QueryBatchRegistAnonymousTbAccountResultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchRegistAnonymousTbAccountResultRequest) GoString() string {
	return s.String()
}

func (s *QueryBatchRegistAnonymousTbAccountResultRequest) SetBizId(v string) *QueryBatchRegistAnonymousTbAccountResultRequest {
	s.BizId = &v
	return s
}

func (s *QueryBatchRegistAnonymousTbAccountResultRequest) SetBatchId(v string) *QueryBatchRegistAnonymousTbAccountResultRequest {
	s.BatchId = &v
	return s
}

type QueryBatchRegistAnonymousTbAccountResultResponseBody struct {
	Status    *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Code      *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BatchId   *string                                                      `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	FailIds   *QueryBatchRegistAnonymousTbAccountResultResponseBodyFailIds `json:"FailIds,omitempty" xml:"FailIds,omitempty" type:"Struct"`
}

func (s QueryBatchRegistAnonymousTbAccountResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchRegistAnonymousTbAccountResultResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBatchRegistAnonymousTbAccountResultResponseBody) SetStatus(v string) *QueryBatchRegistAnonymousTbAccountResultResponseBody {
	s.Status = &v
	return s
}

func (s *QueryBatchRegistAnonymousTbAccountResultResponseBody) SetCode(v string) *QueryBatchRegistAnonymousTbAccountResultResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBatchRegistAnonymousTbAccountResultResponseBody) SetMessage(v string) *QueryBatchRegistAnonymousTbAccountResultResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBatchRegistAnonymousTbAccountResultResponseBody) SetRequestId(v string) *QueryBatchRegistAnonymousTbAccountResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBatchRegistAnonymousTbAccountResultResponseBody) SetBatchId(v string) *QueryBatchRegistAnonymousTbAccountResultResponseBody {
	s.BatchId = &v
	return s
}

func (s *QueryBatchRegistAnonymousTbAccountResultResponseBody) SetFailIds(v *QueryBatchRegistAnonymousTbAccountResultResponseBodyFailIds) *QueryBatchRegistAnonymousTbAccountResultResponseBody {
	s.FailIds = v
	return s
}

type QueryBatchRegistAnonymousTbAccountResultResponseBodyFailIds struct {
	FailId []*string `json:"FailId,omitempty" xml:"FailId,omitempty" type:"Repeated"`
}

func (s QueryBatchRegistAnonymousTbAccountResultResponseBodyFailIds) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchRegistAnonymousTbAccountResultResponseBodyFailIds) GoString() string {
	return s.String()
}

func (s *QueryBatchRegistAnonymousTbAccountResultResponseBodyFailIds) SetFailId(v []*string) *QueryBatchRegistAnonymousTbAccountResultResponseBodyFailIds {
	s.FailId = v
	return s
}

type QueryBatchRegistAnonymousTbAccountResultResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBatchRegistAnonymousTbAccountResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBatchRegistAnonymousTbAccountResultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchRegistAnonymousTbAccountResultResponse) GoString() string {
	return s.String()
}

func (s *QueryBatchRegistAnonymousTbAccountResultResponse) SetHeaders(v map[string]*string) *QueryBatchRegistAnonymousTbAccountResultResponse {
	s.Headers = v
	return s
}

func (s *QueryBatchRegistAnonymousTbAccountResultResponse) SetBody(v *QueryBatchRegistAnonymousTbAccountResultResponseBody) *QueryBatchRegistAnonymousTbAccountResultResponse {
	s.Body = v
	return s
}

type QueryBestSession4ItemsRequest struct {
	BizId     *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LmItemIds map[string]interface{} `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty"`
	ItemIds   map[string]interface{} `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
}

func (s QueryBestSession4ItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBestSession4ItemsRequest) GoString() string {
	return s.String()
}

func (s *QueryBestSession4ItemsRequest) SetBizId(v string) *QueryBestSession4ItemsRequest {
	s.BizId = &v
	return s
}

func (s *QueryBestSession4ItemsRequest) SetLmItemIds(v map[string]interface{}) *QueryBestSession4ItemsRequest {
	s.LmItemIds = v
	return s
}

func (s *QueryBestSession4ItemsRequest) SetItemIds(v map[string]interface{}) *QueryBestSession4ItemsRequest {
	s.ItemIds = v
	return s
}

type QueryBestSession4ItemsShrinkRequest struct {
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LmItemIdsShrink *string `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty"`
	ItemIdsShrink   *string `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
}

func (s QueryBestSession4ItemsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBestSession4ItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryBestSession4ItemsShrinkRequest) SetBizId(v string) *QueryBestSession4ItemsShrinkRequest {
	s.BizId = &v
	return s
}

func (s *QueryBestSession4ItemsShrinkRequest) SetLmItemIdsShrink(v string) *QueryBestSession4ItemsShrinkRequest {
	s.LmItemIdsShrink = &v
	return s
}

func (s *QueryBestSession4ItemsShrinkRequest) SetItemIdsShrink(v string) *QueryBestSession4ItemsShrinkRequest {
	s.ItemIdsShrink = &v
	return s
}

type QueryBestSession4ItemsResponseBody struct {
	Code                           *string                                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                        *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId                      *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LmItemActivitySessionModelList *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelList `json:"LmItemActivitySessionModelList,omitempty" xml:"LmItemActivitySessionModelList,omitempty" type:"Struct"`
}

func (s QueryBestSession4ItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBestSession4ItemsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBestSession4ItemsResponseBody) SetCode(v string) *QueryBestSession4ItemsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBody) SetMessage(v string) *QueryBestSession4ItemsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBody) SetRequestId(v string) *QueryBestSession4ItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBody) SetLmItemActivitySessionModelList(v *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelList) *QueryBestSession4ItemsResponseBody {
	s.LmItemActivitySessionModelList = v
	return s
}

type QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelList struct {
	LmItemActivitySessionModel []*QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModel `json:"LmItemActivitySessionModel,omitempty" xml:"LmItemActivitySessionModel,omitempty" type:"Repeated"`
}

func (s QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelList) String() string {
	return tea.Prettify(s)
}

func (s QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelList) GoString() string {
	return s.String()
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelList) SetLmItemActivitySessionModel(v []*QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModel) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelList {
	s.LmItemActivitySessionModel = v
	return s
}

type QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModel struct {
	LmItemId               *string                                                                                                           `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId                 *int64                                                                                                            `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmActivitySessionModel *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel `json:"LmActivitySessionModel,omitempty" xml:"LmActivitySessionModel,omitempty" type:"Struct"`
}

func (s QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModel) String() string {
	return tea.Prettify(s)
}

func (s QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModel) GoString() string {
	return s.String()
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModel) SetLmItemId(v string) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModel {
	s.LmItemId = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModel) SetItemId(v int64) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModel {
	s.ItemId = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModel) SetLmActivitySessionModel(v *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModel {
	s.LmActivitySessionModel = v
	return s
}

type QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel struct {
	SubBizCode   *string                `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	EndDate      *string                `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	DisplayDate  *string                `json:"DisplayDate,omitempty" xml:"DisplayDate,omitempty"`
	LmSessionId  *int64                 `json:"LmSessionId,omitempty" xml:"LmSessionId,omitempty"`
	Description  *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	StartDate    *string                `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	LmActivityId *int64                 `json:"LmActivityId,omitempty" xml:"LmActivityId,omitempty"`
	BizId        *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Name         *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	ExtInfo      map[string]interface{} `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) String() string {
	return tea.Prettify(s)
}

func (s QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) GoString() string {
	return s.String()
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) SetSubBizCode(v string) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel {
	s.SubBizCode = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) SetEndDate(v string) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel {
	s.EndDate = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) SetDisplayDate(v string) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel {
	s.DisplayDate = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) SetLmSessionId(v int64) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel {
	s.LmSessionId = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) SetDescription(v string) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel {
	s.Description = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) SetStartDate(v string) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel {
	s.StartDate = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) SetLmActivityId(v int64) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel {
	s.LmActivityId = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) SetBizId(v string) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel {
	s.BizId = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) SetName(v string) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel {
	s.Name = &v
	return s
}

func (s *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel) SetExtInfo(v map[string]interface{}) *QueryBestSession4ItemsResponseBodyLmItemActivitySessionModelListLmItemActivitySessionModelLmActivitySessionModel {
	s.ExtInfo = v
	return s
}

type QueryBestSession4ItemsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBestSession4ItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBestSession4ItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBestSession4ItemsResponse) GoString() string {
	return s.String()
}

func (s *QueryBestSession4ItemsResponse) SetHeaders(v map[string]*string) *QueryBestSession4ItemsResponse {
	s.Headers = v
	return s
}

func (s *QueryBestSession4ItemsResponse) SetBody(v *QueryBestSession4ItemsResponseBody) *QueryBestSession4ItemsResponse {
	s.Body = v
	return s
}

type QueryBizItemListRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SubBizId   *string `json:"SubBizId,omitempty" xml:"SubBizId,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	LmItemIds  *string `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty"`
	ItemIds    *string `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s QueryBizItemListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemListRequest) GoString() string {
	return s.String()
}

func (s *QueryBizItemListRequest) SetBizId(v string) *QueryBizItemListRequest {
	s.BizId = &v
	return s
}

func (s *QueryBizItemListRequest) SetSubBizId(v string) *QueryBizItemListRequest {
	s.SubBizId = &v
	return s
}

func (s *QueryBizItemListRequest) SetPageSize(v int32) *QueryBizItemListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryBizItemListRequest) SetPageNumber(v int32) *QueryBizItemListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryBizItemListRequest) SetUserId(v string) *QueryBizItemListRequest {
	s.UserId = &v
	return s
}

func (s *QueryBizItemListRequest) SetLmItemIds(v string) *QueryBizItemListRequest {
	s.LmItemIds = &v
	return s
}

func (s *QueryBizItemListRequest) SetItemIds(v string) *QueryBizItemListRequest {
	s.ItemIds = &v
	return s
}

func (s *QueryBizItemListRequest) SetCategoryId(v int64) *QueryBizItemListRequest {
	s.CategoryId = &v
	return s
}

type QueryBizItemListResponseBody struct {
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ItemList   *QueryBizItemListResponseBodyItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Struct"`
}

func (s QueryBizItemListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBizItemListResponseBody) SetCode(v string) *QueryBizItemListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBizItemListResponseBody) SetMessage(v string) *QueryBizItemListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBizItemListResponseBody) SetTotalCount(v int32) *QueryBizItemListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryBizItemListResponseBody) SetPageNumber(v int32) *QueryBizItemListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryBizItemListResponseBody) SetPageSize(v int32) *QueryBizItemListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryBizItemListResponseBody) SetRequestId(v string) *QueryBizItemListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBizItemListResponseBody) SetItemList(v *QueryBizItemListResponseBodyItemList) *QueryBizItemListResponseBody {
	s.ItemList = v
	return s
}

type QueryBizItemListResponseBodyItemList struct {
	Item []*QueryBizItemListResponseBodyItemListItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryBizItemListResponseBodyItemList) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemListResponseBodyItemList) GoString() string {
	return s.String()
}

func (s *QueryBizItemListResponseBodyItemList) SetItem(v []*QueryBizItemListResponseBodyItemListItem) *QueryBizItemListResponseBodyItemList {
	s.Item = v
	return s
}

type QueryBizItemListResponseBodyItemListItem struct {
	ItemId             *int64                                           `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemTitle          *string                                          `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	CustomizedItemName *string                                          `json:"CustomizedItemName,omitempty" xml:"CustomizedItemName,omitempty"`
	MainPicUrl         *string                                          `json:"MainPicUrl,omitempty" xml:"MainPicUrl,omitempty"`
	ReservePrice       *int64                                           `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	CategoryId         *int64                                           `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CanSell            *bool                                            `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	TaobaoShopName     *string                                          `json:"TaobaoShopName,omitempty" xml:"TaobaoShopName,omitempty"`
	ExtJson            *string                                          `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	SellerId           *int64                                           `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	LmItemId           *string                                          `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	SkuList            *QueryBizItemListResponseBodyItemListItemSkuList `json:"SkuList,omitempty" xml:"SkuList,omitempty" type:"Struct"`
}

func (s QueryBizItemListResponseBodyItemListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemListResponseBodyItemListItem) GoString() string {
	return s.String()
}

func (s *QueryBizItemListResponseBodyItemListItem) SetItemId(v int64) *QueryBizItemListResponseBodyItemListItem {
	s.ItemId = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItem) SetItemTitle(v string) *QueryBizItemListResponseBodyItemListItem {
	s.ItemTitle = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItem) SetCustomizedItemName(v string) *QueryBizItemListResponseBodyItemListItem {
	s.CustomizedItemName = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItem) SetMainPicUrl(v string) *QueryBizItemListResponseBodyItemListItem {
	s.MainPicUrl = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItem) SetReservePrice(v int64) *QueryBizItemListResponseBodyItemListItem {
	s.ReservePrice = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItem) SetCategoryId(v int64) *QueryBizItemListResponseBodyItemListItem {
	s.CategoryId = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItem) SetCanSell(v bool) *QueryBizItemListResponseBodyItemListItem {
	s.CanSell = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItem) SetTaobaoShopName(v string) *QueryBizItemListResponseBodyItemListItem {
	s.TaobaoShopName = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItem) SetExtJson(v string) *QueryBizItemListResponseBodyItemListItem {
	s.ExtJson = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItem) SetSellerId(v int64) *QueryBizItemListResponseBodyItemListItem {
	s.SellerId = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItem) SetLmItemId(v string) *QueryBizItemListResponseBodyItemListItem {
	s.LmItemId = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItem) SetSkuList(v *QueryBizItemListResponseBodyItemListItemSkuList) *QueryBizItemListResponseBodyItemListItem {
	s.SkuList = v
	return s
}

type QueryBizItemListResponseBodyItemListItemSkuList struct {
	Sku []*QueryBizItemListResponseBodyItemListItemSkuListSku `json:"Sku,omitempty" xml:"Sku,omitempty" type:"Repeated"`
}

func (s QueryBizItemListResponseBodyItemListItemSkuList) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemListResponseBodyItemListItemSkuList) GoString() string {
	return s.String()
}

func (s *QueryBizItemListResponseBodyItemListItemSkuList) SetSku(v []*QueryBizItemListResponseBodyItemListItemSkuListSku) *QueryBizItemListResponseBodyItemListItemSkuList {
	s.Sku = v
	return s
}

type QueryBizItemListResponseBodyItemListItemSkuListSku struct {
	SkuId                  *int64                                                              `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuPicUrl              *string                                                             `json:"SkuPicUrl,omitempty" xml:"SkuPicUrl,omitempty"`
	SkuTitle               *string                                                             `json:"SkuTitle,omitempty" xml:"SkuTitle,omitempty"`
	PriceCent              *int64                                                              `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	Points                 *int64                                                              `json:"Points,omitempty" xml:"Points,omitempty"`
	PointsAmount           *int64                                                              `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	BenefitId              *string                                                             `json:"BenefitId,omitempty" xml:"BenefitId,omitempty"`
	CanSell                *bool                                                               `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	CustomizedAttributeMap map[string]*string                                                  `json:"CustomizedAttributeMap,omitempty" xml:"CustomizedAttributeMap,omitempty"`
	TaoBaoCurrentPrice     *int64                                                              `json:"TaoBaoCurrentPrice,omitempty" xml:"TaoBaoCurrentPrice,omitempty"`
	UserLabelList          *QueryBizItemListResponseBodyItemListItemSkuListSkuUserLabelList    `json:"UserLabelList,omitempty" xml:"UserLabelList,omitempty" type:"Struct"`
	GradePriceModels       *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModels `json:"GradePriceModels,omitempty" xml:"GradePriceModels,omitempty" type:"Struct"`
}

func (s QueryBizItemListResponseBodyItemListItemSkuListSku) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemListResponseBodyItemListItemSkuListSku) GoString() string {
	return s.String()
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSku) SetSkuId(v int64) *QueryBizItemListResponseBodyItemListItemSkuListSku {
	s.SkuId = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSku) SetSkuPicUrl(v string) *QueryBizItemListResponseBodyItemListItemSkuListSku {
	s.SkuPicUrl = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSku) SetSkuTitle(v string) *QueryBizItemListResponseBodyItemListItemSkuListSku {
	s.SkuTitle = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSku) SetPriceCent(v int64) *QueryBizItemListResponseBodyItemListItemSkuListSku {
	s.PriceCent = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSku) SetPoints(v int64) *QueryBizItemListResponseBodyItemListItemSkuListSku {
	s.Points = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSku) SetPointsAmount(v int64) *QueryBizItemListResponseBodyItemListItemSkuListSku {
	s.PointsAmount = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSku) SetBenefitId(v string) *QueryBizItemListResponseBodyItemListItemSkuListSku {
	s.BenefitId = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSku) SetCanSell(v bool) *QueryBizItemListResponseBodyItemListItemSkuListSku {
	s.CanSell = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSku) SetCustomizedAttributeMap(v map[string]*string) *QueryBizItemListResponseBodyItemListItemSkuListSku {
	s.CustomizedAttributeMap = v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSku) SetTaoBaoCurrentPrice(v int64) *QueryBizItemListResponseBodyItemListItemSkuListSku {
	s.TaoBaoCurrentPrice = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSku) SetUserLabelList(v *QueryBizItemListResponseBodyItemListItemSkuListSkuUserLabelList) *QueryBizItemListResponseBodyItemListItemSkuListSku {
	s.UserLabelList = v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSku) SetGradePriceModels(v *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModels) *QueryBizItemListResponseBodyItemListItemSkuListSku {
	s.GradePriceModels = v
	return s
}

type QueryBizItemListResponseBodyItemListItemSkuListSkuUserLabelList struct {
	UserLabelList []*string `json:"UserLabelList,omitempty" xml:"UserLabelList,omitempty" type:"Repeated"`
}

func (s QueryBizItemListResponseBodyItemListItemSkuListSkuUserLabelList) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemListResponseBodyItemListItemSkuListSkuUserLabelList) GoString() string {
	return s.String()
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuUserLabelList) SetUserLabelList(v []*string) *QueryBizItemListResponseBodyItemListItemSkuListSkuUserLabelList {
	s.UserLabelList = v
	return s
}

type QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModels struct {
	GradePriceModel []*QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel `json:"GradePriceModel,omitempty" xml:"GradePriceModel,omitempty" type:"Repeated"`
}

func (s QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModels) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModels) GoString() string {
	return s.String()
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModels) SetGradePriceModel(v []*QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModels {
	s.GradePriceModel = v
	return s
}

type QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel struct {
	CharacteristicCode *string                                                                                         `json:"CharacteristicCode,omitempty" xml:"CharacteristicCode,omitempty"`
	CharacteristicName *string                                                                                         `json:"CharacteristicName,omitempty" xml:"CharacteristicName,omitempty"`
	SubBizCode         *string                                                                                         `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	CanBuy             *bool                                                                                           `json:"CanBuy,omitempty" xml:"CanBuy,omitempty"`
	Exclusive          *bool                                                                                           `json:"Exclusive,omitempty" xml:"Exclusive,omitempty"`
	Recommend          *bool                                                                                           `json:"Recommend,omitempty" xml:"Recommend,omitempty"`
	PriceCent          *int64                                                                                          `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	PointsAmount       *int64                                                                                          `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	Points             *int64                                                                                          `json:"Points,omitempty" xml:"Points,omitempty"`
	PointPrice         *int64                                                                                          `json:"PointPrice,omitempty" xml:"PointPrice,omitempty"`
	ShowName           *string                                                                                         `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	AccessUrl          *string                                                                                         `json:"AccessUrl,omitempty" xml:"AccessUrl,omitempty"`
	Icon               *string                                                                                         `json:"Icon,omitempty" xml:"Icon,omitempty"`
	UserLabelList      *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModelUserLabelList `json:"UserLabelList,omitempty" xml:"UserLabelList,omitempty" type:"Struct"`
}

func (s QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) GoString() string {
	return s.String()
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetCharacteristicCode(v string) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.CharacteristicCode = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetCharacteristicName(v string) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.CharacteristicName = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetSubBizCode(v string) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.SubBizCode = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetCanBuy(v bool) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.CanBuy = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetExclusive(v bool) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.Exclusive = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetRecommend(v bool) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.Recommend = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetPriceCent(v int64) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.PriceCent = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetPointsAmount(v int64) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.PointsAmount = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetPoints(v int64) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.Points = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetPointPrice(v int64) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.PointPrice = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetShowName(v string) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.ShowName = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetAccessUrl(v string) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.AccessUrl = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetIcon(v string) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.Icon = &v
	return s
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel) SetUserLabelList(v *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModelUserLabelList) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModel {
	s.UserLabelList = v
	return s
}

type QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModelUserLabelList struct {
	UserLabelList []*string `json:"UserLabelList,omitempty" xml:"UserLabelList,omitempty" type:"Repeated"`
}

func (s QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModelUserLabelList) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModelUserLabelList) GoString() string {
	return s.String()
}

func (s *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModelUserLabelList) SetUserLabelList(v []*string) *QueryBizItemListResponseBodyItemListItemSkuListSkuGradePriceModelsGradePriceModelUserLabelList {
	s.UserLabelList = v
	return s
}

type QueryBizItemListResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBizItemListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBizItemListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemListResponse) GoString() string {
	return s.String()
}

func (s *QueryBizItemListResponse) SetHeaders(v map[string]*string) *QueryBizItemListResponse {
	s.Headers = v
	return s
}

func (s *QueryBizItemListResponse) SetBody(v *QueryBizItemListResponseBody) *QueryBizItemListResponse {
	s.Body = v
	return s
}

type QueryBizItemsRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SubBizId   *string `json:"SubBizId,omitempty" xml:"SubBizId,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryBizItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsRequest) GoString() string {
	return s.String()
}

func (s *QueryBizItemsRequest) SetBizId(v string) *QueryBizItemsRequest {
	s.BizId = &v
	return s
}

func (s *QueryBizItemsRequest) SetSubBizId(v string) *QueryBizItemsRequest {
	s.SubBizId = &v
	return s
}

func (s *QueryBizItemsRequest) SetPageSize(v int32) *QueryBizItemsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryBizItemsRequest) SetPageNumber(v int32) *QueryBizItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryBizItemsRequest) SetUserId(v string) *QueryBizItemsRequest {
	s.UserId = &v
	return s
}

type QueryBizItemsResponseBody struct {
	Model      *string `json:"Model,omitempty" xml:"Model,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryBizItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBizItemsResponseBody) SetModel(v string) *QueryBizItemsResponseBody {
	s.Model = &v
	return s
}

func (s *QueryBizItemsResponseBody) SetRequestId(v string) *QueryBizItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBizItemsResponseBody) SetCode(v string) *QueryBizItemsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBizItemsResponseBody) SetMessage(v string) *QueryBizItemsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBizItemsResponseBody) SetPageNumber(v int32) *QueryBizItemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryBizItemsResponseBody) SetPageSize(v int32) *QueryBizItemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryBizItemsResponseBody) SetTotalCount(v int32) *QueryBizItemsResponseBody {
	s.TotalCount = &v
	return s
}

type QueryBizItemsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBizItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBizItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsResponse) GoString() string {
	return s.String()
}

func (s *QueryBizItemsResponse) SetHeaders(v map[string]*string) *QueryBizItemsResponse {
	s.Headers = v
	return s
}

func (s *QueryBizItemsResponse) SetBody(v *QueryBizItemsResponseBody) *QueryBizItemsResponse {
	s.Body = v
	return s
}

type QueryBizItemsWithActivityRequest struct {
	BizId     *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LmItemIds map[string]interface{} `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty"`
	ItemIds   map[string]interface{} `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
}

func (s QueryBizItemsWithActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityRequest) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityRequest) SetBizId(v string) *QueryBizItemsWithActivityRequest {
	s.BizId = &v
	return s
}

func (s *QueryBizItemsWithActivityRequest) SetLmItemIds(v map[string]interface{}) *QueryBizItemsWithActivityRequest {
	s.LmItemIds = v
	return s
}

func (s *QueryBizItemsWithActivityRequest) SetItemIds(v map[string]interface{}) *QueryBizItemsWithActivityRequest {
	s.ItemIds = v
	return s
}

type QueryBizItemsWithActivityShrinkRequest struct {
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LmItemIdsShrink *string `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty"`
	ItemIdsShrink   *string `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
}

func (s QueryBizItemsWithActivityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityShrinkRequest) SetBizId(v string) *QueryBizItemsWithActivityShrinkRequest {
	s.BizId = &v
	return s
}

func (s *QueryBizItemsWithActivityShrinkRequest) SetLmItemIdsShrink(v string) *QueryBizItemsWithActivityShrinkRequest {
	s.LmItemIdsShrink = &v
	return s
}

func (s *QueryBizItemsWithActivityShrinkRequest) SetItemIdsShrink(v string) *QueryBizItemsWithActivityShrinkRequest {
	s.ItemIdsShrink = &v
	return s
}

type QueryBizItemsWithActivityResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ItemList  *QueryBizItemsWithActivityResponseBodyItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Struct"`
}

func (s QueryBizItemsWithActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBody) SetCode(v string) *QueryBizItemsWithActivityResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBody) SetMessage(v string) *QueryBizItemsWithActivityResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBody) SetRequestId(v string) *QueryBizItemsWithActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBody) SetItemList(v *QueryBizItemsWithActivityResponseBodyItemList) *QueryBizItemsWithActivityResponseBody {
	s.ItemList = v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemList struct {
	Item []*QueryBizItemsWithActivityResponseBodyItemListItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryBizItemsWithActivityResponseBodyItemList) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemList) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemList) SetItem(v []*QueryBizItemsWithActivityResponseBodyItemListItem) *QueryBizItemsWithActivityResponseBodyItemList {
	s.Item = v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemListItem struct {
	PicUrl            *string                                                      `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	TotalSoldQuantity *int32                                                       `json:"TotalSoldQuantity,omitempty" xml:"TotalSoldQuantity,omitempty"`
	ItemTitle         *string                                                      `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	MaxAllowedCount   *int32                                                       `json:"MaxAllowedCount,omitempty" xml:"MaxAllowedCount,omitempty"`
	SellerId          *int64                                                       `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	LmItemId          *string                                                      `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	CategoryId        *int64                                                       `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	ItemId            *int64                                                       `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ReservePrice      *int64                                                       `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	Quantity          *int64                                                       `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	TaobaoShopName    *string                                                      `json:"TaobaoShopName,omitempty" xml:"TaobaoShopName,omitempty"`
	SkuList           *QueryBizItemsWithActivityResponseBodyItemListItemSkuList    `json:"SkuList,omitempty" xml:"SkuList,omitempty" type:"Struct"`
	Activities        *QueryBizItemsWithActivityResponseBodyItemListItemActivities `json:"Activities,omitempty" xml:"Activities,omitempty" type:"Struct"`
}

func (s QueryBizItemsWithActivityResponseBodyItemListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemListItem) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetPicUrl(v string) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.PicUrl = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetTotalSoldQuantity(v int32) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.TotalSoldQuantity = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetItemTitle(v string) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.ItemTitle = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetMaxAllowedCount(v int32) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.MaxAllowedCount = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetSellerId(v int64) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.SellerId = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetLmItemId(v string) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.LmItemId = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetCategoryId(v int64) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.CategoryId = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetItemId(v int64) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.ItemId = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetReservePrice(v int64) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.ReservePrice = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetQuantity(v int64) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.Quantity = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetTaobaoShopName(v string) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.TaobaoShopName = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetSkuList(v *QueryBizItemsWithActivityResponseBodyItemListItemSkuList) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.SkuList = v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItem) SetActivities(v *QueryBizItemsWithActivityResponseBodyItemListItemActivities) *QueryBizItemsWithActivityResponseBodyItemListItem {
	s.Activities = v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemListItemSkuList struct {
	Sku []*QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku `json:"Sku,omitempty" xml:"Sku,omitempty" type:"Repeated"`
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemSkuList) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemSkuList) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemSkuList) SetSku(v []*QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku) *QueryBizItemsWithActivityResponseBodyItemListItemSkuList {
	s.Sku = v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku struct {
	Status          *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	PriceCent       *int64  `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	PointsInfo      *string `json:"PointsInfo,omitempty" xml:"PointsInfo,omitempty"`
	MaxAllowedCount *int32  `json:"MaxAllowedCount,omitempty" xml:"MaxAllowedCount,omitempty"`
	PointsKey       *string `json:"PointsKey,omitempty" xml:"PointsKey,omitempty"`
	SkuId           *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Points          *int64  `json:"Points,omitempty" xml:"Points,omitempty"`
	PointsAmount    *int64  `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	Quantity        *int64  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku) SetStatus(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku {
	s.Status = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku) SetPriceCent(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku {
	s.PriceCent = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku) SetPointsInfo(v string) *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku {
	s.PointsInfo = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku) SetMaxAllowedCount(v int32) *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku {
	s.MaxAllowedCount = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku) SetPointsKey(v string) *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku {
	s.PointsKey = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku) SetSkuId(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku {
	s.SkuId = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku) SetPoints(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku {
	s.Points = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku) SetPointsAmount(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku {
	s.PointsAmount = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku) SetQuantity(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemSkuListSku {
	s.Quantity = &v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemListItemActivities struct {
	Activity []*QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity `json:"Activity,omitempty" xml:"Activity,omitempty" type:"Repeated"`
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivities) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivities) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivities) SetActivity(v []*QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity) *QueryBizItemsWithActivityResponseBodyItemListItemActivities {
	s.Activity = v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity struct {
	StartDate        *int64                                                                               `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	LmActivityId     *int64                                                                               `json:"LmActivityId,omitempty" xml:"LmActivityId,omitempty"`
	EndDate          *int64                                                                               `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Title            *string                                                                              `json:"Title,omitempty" xml:"Title,omitempty"`
	Description      *string                                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	ActivitySessions *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessions `json:"ActivitySessions,omitempty" xml:"ActivitySessions,omitempty" type:"Struct"`
	ActivityItem     *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItem     `json:"ActivityItem,omitempty" xml:"ActivityItem,omitempty" type:"Struct"`
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity) SetStartDate(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity {
	s.StartDate = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity) SetLmActivityId(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity {
	s.LmActivityId = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity) SetEndDate(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity {
	s.EndDate = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity) SetTitle(v string) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity {
	s.Title = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity) SetDescription(v string) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity {
	s.Description = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity) SetActivitySessions(v *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessions) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity {
	s.ActivitySessions = v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity) SetActivityItem(v *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItem) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivity {
	s.ActivityItem = v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessions struct {
	ActivitySession []*QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession `json:"ActivitySession,omitempty" xml:"ActivitySession,omitempty" type:"Repeated"`
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessions) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessions) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessions) SetActivitySession(v []*QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessions {
	s.ActivitySession = v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession struct {
	SubBizCode          *string                                                                                                                `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	EndDate             *int64                                                                                                                 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	DisplayDate         *int64                                                                                                                 `json:"DisplayDate,omitempty" xml:"DisplayDate,omitempty"`
	LmSessionId         *int64                                                                                                                 `json:"LmSessionId,omitempty" xml:"LmSessionId,omitempty"`
	Description         *string                                                                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	StartDate           *int64                                                                                                                 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Title               *string                                                                                                                `json:"Title,omitempty" xml:"Title,omitempty"`
	ActivitySessionItem *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem `json:"ActivitySessionItem,omitempty" xml:"ActivitySessionItem,omitempty" type:"Struct"`
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession) SetSubBizCode(v string) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession {
	s.SubBizCode = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession) SetEndDate(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession {
	s.EndDate = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession) SetDisplayDate(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession {
	s.DisplayDate = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession) SetLmSessionId(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession {
	s.LmSessionId = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession) SetDescription(v string) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession {
	s.Description = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession) SetStartDate(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession {
	s.StartDate = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession) SetTitle(v string) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession {
	s.Title = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession) SetActivitySessionItem(v *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySession {
	s.ActivitySessionItem = v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem struct {
	SessionQuantity            *int64                                                                                                                                           `json:"SessionQuantity,omitempty" xml:"SessionQuantity,omitempty"`
	LimitQuantityForPerson     *int64                                                                                                                                           `json:"LimitQuantityForPerson,omitempty" xml:"LimitQuantityForPerson,omitempty"`
	SaleableQuantity           *int64                                                                                                                                           `json:"SaleableQuantity,omitempty" xml:"SaleableQuantity,omitempty"`
	ActivitySessionItemSkuList *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList `json:"ActivitySessionItemSkuList,omitempty" xml:"ActivitySessionItemSkuList,omitempty" type:"Struct"`
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem) SetSessionQuantity(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem {
	s.SessionQuantity = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem) SetLimitQuantityForPerson(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem {
	s.LimitQuantityForPerson = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem) SetSaleableQuantity(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem {
	s.SaleableQuantity = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem) SetActivitySessionItemSkuList(v *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItem {
	s.ActivitySessionItemSkuList = v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList struct {
	SkuId        *int64 `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	PriceCent    *int64 `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	Points       *int64 `json:"Points,omitempty" xml:"Points,omitempty"`
	PointsAmount *int64 `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList) SetSkuId(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList {
	s.SkuId = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList) SetPriceCent(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList {
	s.PriceCent = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList) SetPoints(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList {
	s.Points = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList) SetPointsAmount(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivitySessionsActivitySessionActivitySessionItemActivitySessionItemSkuList {
	s.PointsAmount = &v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItem struct {
	ActivityQuantity       *int64                                                                                              `json:"ActivityQuantity,omitempty" xml:"ActivityQuantity,omitempty"`
	LimitQuantityForPerson *int64                                                                                              `json:"LimitQuantityForPerson,omitempty" xml:"LimitQuantityForPerson,omitempty"`
	ActivityItemSkuList    *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuList `json:"ActivityItemSkuList,omitempty" xml:"ActivityItemSkuList,omitempty" type:"Struct"`
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItem) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItem) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItem) SetActivityQuantity(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItem {
	s.ActivityQuantity = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItem) SetLimitQuantityForPerson(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItem {
	s.LimitQuantityForPerson = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItem) SetActivityItemSkuList(v *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuList) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItem {
	s.ActivityItemSkuList = v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuList struct {
	ActivityItemSku []*QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuListActivityItemSku `json:"ActivityItemSku,omitempty" xml:"ActivityItemSku,omitempty" type:"Repeated"`
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuList) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuList) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuList) SetActivityItemSku(v []*QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuListActivityItemSku) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuList {
	s.ActivityItemSku = v
	return s
}

type QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuListActivityItemSku struct {
	ActivityPrice *int64 `json:"ActivityPrice,omitempty" xml:"ActivityPrice,omitempty"`
	SkuId         *int64 `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuListActivityItemSku) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuListActivityItemSku) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuListActivityItemSku) SetActivityPrice(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuListActivityItemSku {
	s.ActivityPrice = &v
	return s
}

func (s *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuListActivityItemSku) SetSkuId(v int64) *QueryBizItemsWithActivityResponseBodyItemListItemActivitiesActivityActivityItemActivityItemSkuListActivityItemSku {
	s.SkuId = &v
	return s
}

type QueryBizItemsWithActivityResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBizItemsWithActivityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBizItemsWithActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBizItemsWithActivityResponse) GoString() string {
	return s.String()
}

func (s *QueryBizItemsWithActivityResponse) SetHeaders(v map[string]*string) *QueryBizItemsWithActivityResponse {
	s.Headers = v
	return s
}

func (s *QueryBizItemsWithActivityResponse) SetBody(v *QueryBizItemsWithActivityResponseBody) *QueryBizItemsWithActivityResponse {
	s.Body = v
	return s
}

type QueryGuideItemGroupRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryGuideItemGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGuideItemGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryGuideItemGroupRequest) SetBizId(v string) *QueryGuideItemGroupRequest {
	s.BizId = &v
	return s
}

func (s *QueryGuideItemGroupRequest) SetGroupId(v string) *QueryGuideItemGroupRequest {
	s.GroupId = &v
	return s
}

func (s *QueryGuideItemGroupRequest) SetPageNumber(v int64) *QueryGuideItemGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryGuideItemGroupRequest) SetPageSize(v int64) *QueryGuideItemGroupRequest {
	s.PageSize = &v
	return s
}

type QueryGuideItemGroupResponseBody struct {
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage     *string                                        `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code           *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode        *string                                        `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	GuideItemGroup *QueryGuideItemGroupResponseBodyGuideItemGroup `json:"GuideItemGroup,omitempty" xml:"GuideItemGroup,omitempty" type:"Struct"`
}

func (s QueryGuideItemGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGuideItemGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGuideItemGroupResponseBody) SetRequestId(v string) *QueryGuideItemGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGuideItemGroupResponseBody) SetSuccess(v bool) *QueryGuideItemGroupResponseBody {
	s.Success = &v
	return s
}

func (s *QueryGuideItemGroupResponseBody) SetSubMessage(v string) *QueryGuideItemGroupResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryGuideItemGroupResponseBody) SetCode(v string) *QueryGuideItemGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGuideItemGroupResponseBody) SetMessage(v string) *QueryGuideItemGroupResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGuideItemGroupResponseBody) SetSubCode(v string) *QueryGuideItemGroupResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryGuideItemGroupResponseBody) SetGuideItemGroup(v *QueryGuideItemGroupResponseBodyGuideItemGroup) *QueryGuideItemGroupResponseBody {
	s.GuideItemGroup = v
	return s
}

type QueryGuideItemGroupResponseBodyGuideItemGroup struct {
	ItemInfo []*QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo `json:"ItemInfo,omitempty" xml:"ItemInfo,omitempty" type:"Repeated"`
}

func (s QueryGuideItemGroupResponseBodyGuideItemGroup) String() string {
	return tea.Prettify(s)
}

func (s QueryGuideItemGroupResponseBodyGuideItemGroup) GoString() string {
	return s.String()
}

func (s *QueryGuideItemGroupResponseBodyGuideItemGroup) SetItemInfo(v []*QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) *QueryGuideItemGroupResponseBodyGuideItemGroup {
	s.ItemInfo = v
	return s
}

type QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo struct {
	PriceCent    *int64  `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	MainPicUrl   *string `json:"MainPicUrl,omitempty" xml:"MainPicUrl,omitempty"`
	ItemTitle    *string `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	LmItemId     *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	PointPrice   *int64  `json:"PointPrice,omitempty" xml:"PointPrice,omitempty"`
	ItemId       *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Points       *int64  `json:"Points,omitempty" xml:"Points,omitempty"`
	PointsAmount *int64  `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	ReservePrice *string `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	ItemUrl      *string `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
}

func (s QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) GoString() string {
	return s.String()
}

func (s *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) SetPriceCent(v int64) *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo {
	s.PriceCent = &v
	return s
}

func (s *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) SetMainPicUrl(v string) *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo {
	s.MainPicUrl = &v
	return s
}

func (s *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) SetItemTitle(v string) *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo {
	s.ItemTitle = &v
	return s
}

func (s *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) SetLmItemId(v string) *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo {
	s.LmItemId = &v
	return s
}

func (s *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) SetPointPrice(v int64) *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo {
	s.PointPrice = &v
	return s
}

func (s *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) SetItemId(v int64) *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo {
	s.ItemId = &v
	return s
}

func (s *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) SetPoints(v int64) *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo {
	s.Points = &v
	return s
}

func (s *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) SetPointsAmount(v int64) *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo {
	s.PointsAmount = &v
	return s
}

func (s *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) SetReservePrice(v string) *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo {
	s.ReservePrice = &v
	return s
}

func (s *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo) SetItemUrl(v string) *QueryGuideItemGroupResponseBodyGuideItemGroupItemInfo {
	s.ItemUrl = &v
	return s
}

type QueryGuideItemGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGuideItemGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGuideItemGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGuideItemGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryGuideItemGroupResponse) SetHeaders(v map[string]*string) *QueryGuideItemGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryGuideItemGroupResponse) SetBody(v *QueryGuideItemGroupResponseBody) *QueryGuideItemGroupResponse {
	s.Body = v
	return s
}

type QueryHotMoviesRequest struct {
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CityCode *int64  `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	ExtJson  *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s QueryHotMoviesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotMoviesRequest) GoString() string {
	return s.String()
}

func (s *QueryHotMoviesRequest) SetBizId(v string) *QueryHotMoviesRequest {
	s.BizId = &v
	return s
}

func (s *QueryHotMoviesRequest) SetCityCode(v int64) *QueryHotMoviesRequest {
	s.CityCode = &v
	return s
}

func (s *QueryHotMoviesRequest) SetExtJson(v string) *QueryHotMoviesRequest {
	s.ExtJson = &v
	return s
}

type QueryHotMoviesResponseBody struct {
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                           `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                           `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	LogsId     *string                           `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Movies     *QueryHotMoviesResponseBodyMovies `json:"Movies,omitempty" xml:"Movies,omitempty" type:"Struct"`
}

func (s QueryHotMoviesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHotMoviesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotMoviesResponseBody) SetRequestId(v string) *QueryHotMoviesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHotMoviesResponseBody) SetSuccess(v bool) *QueryHotMoviesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryHotMoviesResponseBody) SetSubMessage(v string) *QueryHotMoviesResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryHotMoviesResponseBody) SetCode(v string) *QueryHotMoviesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryHotMoviesResponseBody) SetMessage(v string) *QueryHotMoviesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHotMoviesResponseBody) SetSubCode(v string) *QueryHotMoviesResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryHotMoviesResponseBody) SetLogsId(v string) *QueryHotMoviesResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryHotMoviesResponseBody) SetMovies(v *QueryHotMoviesResponseBodyMovies) *QueryHotMoviesResponseBody {
	s.Movies = v
	return s
}

type QueryHotMoviesResponseBodyMovies struct {
	Movie []*QueryHotMoviesResponseBodyMoviesMovie `json:"Movie,omitempty" xml:"Movie,omitempty" type:"Repeated"`
}

func (s QueryHotMoviesResponseBodyMovies) String() string {
	return tea.Prettify(s)
}

func (s QueryHotMoviesResponseBodyMovies) GoString() string {
	return s.String()
}

func (s *QueryHotMoviesResponseBodyMovies) SetMovie(v []*QueryHotMoviesResponseBodyMoviesMovie) *QueryHotMoviesResponseBodyMovies {
	s.Movie = v
	return s
}

type QueryHotMoviesResponseBodyMoviesMovie struct {
	Type              *string                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	MovieVersion      *string                                             `json:"MovieVersion,omitempty" xml:"MovieVersion,omitempty"`
	BackgroundPicture *string                                             `json:"BackgroundPicture,omitempty" xml:"BackgroundPicture,omitempty"`
	Remark            *string                                             `json:"Remark,omitempty" xml:"Remark,omitempty"`
	OpenDay           *string                                             `json:"OpenDay,omitempty" xml:"OpenDay,omitempty"`
	Highlight         *string                                             `json:"Highlight,omitempty" xml:"Highlight,omitempty"`
	Language          *string                                             `json:"Language,omitempty" xml:"Language,omitempty"`
	OpenTime          *string                                             `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	Director          *string                                             `json:"Director,omitempty" xml:"Director,omitempty"`
	Poster            *string                                             `json:"Poster,omitempty" xml:"Poster,omitempty"`
	MovieName         *string                                             `json:"MovieName,omitempty" xml:"MovieName,omitempty"`
	Description       *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	Country           *string                                             `json:"Country,omitempty" xml:"Country,omitempty"`
	Duration          *int64                                              `json:"Duration,omitempty" xml:"Duration,omitempty"`
	MovieNameEn       *string                                             `json:"MovieNameEn,omitempty" xml:"MovieNameEn,omitempty"`
	LeadingRole       *string                                             `json:"LeadingRole,omitempty" xml:"LeadingRole,omitempty"`
	Id                *int64                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	MovieTypeList     *QueryHotMoviesResponseBodyMoviesMovieMovieTypeList `json:"MovieTypeList,omitempty" xml:"MovieTypeList,omitempty" type:"Struct"`
	TrailerList       *QueryHotMoviesResponseBodyMoviesMovieTrailerList   `json:"TrailerList,omitempty" xml:"TrailerList,omitempty" type:"Struct"`
}

func (s QueryHotMoviesResponseBodyMoviesMovie) String() string {
	return tea.Prettify(s)
}

func (s QueryHotMoviesResponseBodyMoviesMovie) GoString() string {
	return s.String()
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetType(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.Type = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetMovieVersion(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.MovieVersion = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetBackgroundPicture(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.BackgroundPicture = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetRemark(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.Remark = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetOpenDay(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.OpenDay = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetHighlight(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.Highlight = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetLanguage(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.Language = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetOpenTime(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.OpenTime = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetDirector(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.Director = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetPoster(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.Poster = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetMovieName(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.MovieName = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetDescription(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.Description = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetCountry(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.Country = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetDuration(v int64) *QueryHotMoviesResponseBodyMoviesMovie {
	s.Duration = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetMovieNameEn(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.MovieNameEn = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetLeadingRole(v string) *QueryHotMoviesResponseBodyMoviesMovie {
	s.LeadingRole = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetId(v int64) *QueryHotMoviesResponseBodyMoviesMovie {
	s.Id = &v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetMovieTypeList(v *QueryHotMoviesResponseBodyMoviesMovieMovieTypeList) *QueryHotMoviesResponseBodyMoviesMovie {
	s.MovieTypeList = v
	return s
}

func (s *QueryHotMoviesResponseBodyMoviesMovie) SetTrailerList(v *QueryHotMoviesResponseBodyMoviesMovieTrailerList) *QueryHotMoviesResponseBodyMoviesMovie {
	s.TrailerList = v
	return s
}

type QueryHotMoviesResponseBodyMoviesMovieMovieTypeList struct {
	MovieTypeList []*string `json:"MovieTypeList,omitempty" xml:"MovieTypeList,omitempty" type:"Repeated"`
}

func (s QueryHotMoviesResponseBodyMoviesMovieMovieTypeList) String() string {
	return tea.Prettify(s)
}

func (s QueryHotMoviesResponseBodyMoviesMovieMovieTypeList) GoString() string {
	return s.String()
}

func (s *QueryHotMoviesResponseBodyMoviesMovieMovieTypeList) SetMovieTypeList(v []*string) *QueryHotMoviesResponseBodyMoviesMovieMovieTypeList {
	s.MovieTypeList = v
	return s
}

type QueryHotMoviesResponseBodyMoviesMovieTrailerList struct {
	TrailerList []*string `json:"TrailerList,omitempty" xml:"TrailerList,omitempty" type:"Repeated"`
}

func (s QueryHotMoviesResponseBodyMoviesMovieTrailerList) String() string {
	return tea.Prettify(s)
}

func (s QueryHotMoviesResponseBodyMoviesMovieTrailerList) GoString() string {
	return s.String()
}

func (s *QueryHotMoviesResponseBodyMoviesMovieTrailerList) SetTrailerList(v []*string) *QueryHotMoviesResponseBodyMoviesMovieTrailerList {
	s.TrailerList = v
	return s
}

type QueryHotMoviesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHotMoviesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHotMoviesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHotMoviesResponse) GoString() string {
	return s.String()
}

func (s *QueryHotMoviesResponse) SetHeaders(v map[string]*string) *QueryHotMoviesResponse {
	s.Headers = v
	return s
}

func (s *QueryHotMoviesResponse) SetBody(v *QueryHotMoviesResponseBody) *QueryHotMoviesResponse {
	s.Body = v
	return s
}

type QueryInventoryOfItemsInBizItemGroupRequest struct {
	BizId        *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	DivisionCode *string                `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
	ItemIds      map[string]interface{} `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
	LmItemIds    map[string]interface{} `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty"`
	SubBizId     *string                `json:"SubBizId,omitempty" xml:"SubBizId,omitempty"`
}

func (s QueryInventoryOfItemsInBizItemGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInventoryOfItemsInBizItemGroupRequest) GoString() string {
	return s.String()
}

func (s *QueryInventoryOfItemsInBizItemGroupRequest) SetBizId(v string) *QueryInventoryOfItemsInBizItemGroupRequest {
	s.BizId = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupRequest) SetDivisionCode(v string) *QueryInventoryOfItemsInBizItemGroupRequest {
	s.DivisionCode = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupRequest) SetItemIds(v map[string]interface{}) *QueryInventoryOfItemsInBizItemGroupRequest {
	s.ItemIds = v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupRequest) SetLmItemIds(v map[string]interface{}) *QueryInventoryOfItemsInBizItemGroupRequest {
	s.LmItemIds = v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupRequest) SetSubBizId(v string) *QueryInventoryOfItemsInBizItemGroupRequest {
	s.SubBizId = &v
	return s
}

type QueryInventoryOfItemsInBizItemGroupShrinkRequest struct {
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	DivisionCode    *string `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
	ItemIdsShrink   *string `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
	LmItemIdsShrink *string `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty"`
	SubBizId        *string `json:"SubBizId,omitempty" xml:"SubBizId,omitempty"`
}

func (s QueryInventoryOfItemsInBizItemGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInventoryOfItemsInBizItemGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryInventoryOfItemsInBizItemGroupShrinkRequest) SetBizId(v string) *QueryInventoryOfItemsInBizItemGroupShrinkRequest {
	s.BizId = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupShrinkRequest) SetDivisionCode(v string) *QueryInventoryOfItemsInBizItemGroupShrinkRequest {
	s.DivisionCode = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupShrinkRequest) SetItemIdsShrink(v string) *QueryInventoryOfItemsInBizItemGroupShrinkRequest {
	s.ItemIdsShrink = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupShrinkRequest) SetLmItemIdsShrink(v string) *QueryInventoryOfItemsInBizItemGroupShrinkRequest {
	s.LmItemIdsShrink = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupShrinkRequest) SetSubBizId(v string) *QueryInventoryOfItemsInBizItemGroupShrinkRequest {
	s.SubBizId = &v
	return s
}

type QueryInventoryOfItemsInBizItemGroupResponseBody struct {
	Code      *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ItemList  []*QueryInventoryOfItemsInBizItemGroupResponseBodyItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
}

func (s QueryInventoryOfItemsInBizItemGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInventoryOfItemsInBizItemGroupResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInventoryOfItemsInBizItemGroupResponseBody) SetCode(v string) *QueryInventoryOfItemsInBizItemGroupResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupResponseBody) SetMessage(v string) *QueryInventoryOfItemsInBizItemGroupResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupResponseBody) SetRequestId(v string) *QueryInventoryOfItemsInBizItemGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupResponseBody) SetItemList(v []*QueryInventoryOfItemsInBizItemGroupResponseBodyItemList) *QueryInventoryOfItemsInBizItemGroupResponseBody {
	s.ItemList = v
	return s
}

type QueryInventoryOfItemsInBizItemGroupResponseBodyItemList struct {
	LmItemId *string                                                           `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId   *int64                                                            `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Quantity *int32                                                            `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	SkuList  []*QueryInventoryOfItemsInBizItemGroupResponseBodyItemListSkuList `json:"SkuList,omitempty" xml:"SkuList,omitempty" type:"Repeated"`
}

func (s QueryInventoryOfItemsInBizItemGroupResponseBodyItemList) String() string {
	return tea.Prettify(s)
}

func (s QueryInventoryOfItemsInBizItemGroupResponseBodyItemList) GoString() string {
	return s.String()
}

func (s *QueryInventoryOfItemsInBizItemGroupResponseBodyItemList) SetLmItemId(v string) *QueryInventoryOfItemsInBizItemGroupResponseBodyItemList {
	s.LmItemId = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupResponseBodyItemList) SetItemId(v int64) *QueryInventoryOfItemsInBizItemGroupResponseBodyItemList {
	s.ItemId = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupResponseBodyItemList) SetQuantity(v int32) *QueryInventoryOfItemsInBizItemGroupResponseBodyItemList {
	s.Quantity = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupResponseBodyItemList) SetSkuList(v []*QueryInventoryOfItemsInBizItemGroupResponseBodyItemListSkuList) *QueryInventoryOfItemsInBizItemGroupResponseBodyItemList {
	s.SkuList = v
	return s
}

type QueryInventoryOfItemsInBizItemGroupResponseBodyItemListSkuList struct {
	SkuId    *int64 `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s QueryInventoryOfItemsInBizItemGroupResponseBodyItemListSkuList) String() string {
	return tea.Prettify(s)
}

func (s QueryInventoryOfItemsInBizItemGroupResponseBodyItemListSkuList) GoString() string {
	return s.String()
}

func (s *QueryInventoryOfItemsInBizItemGroupResponseBodyItemListSkuList) SetSkuId(v int64) *QueryInventoryOfItemsInBizItemGroupResponseBodyItemListSkuList {
	s.SkuId = &v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupResponseBodyItemListSkuList) SetQuantity(v int32) *QueryInventoryOfItemsInBizItemGroupResponseBodyItemListSkuList {
	s.Quantity = &v
	return s
}

type QueryInventoryOfItemsInBizItemGroupResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryInventoryOfItemsInBizItemGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryInventoryOfItemsInBizItemGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInventoryOfItemsInBizItemGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryInventoryOfItemsInBizItemGroupResponse) SetHeaders(v map[string]*string) *QueryInventoryOfItemsInBizItemGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryInventoryOfItemsInBizItemGroupResponse) SetBody(v *QueryInventoryOfItemsInBizItemGroupResponseBody) *QueryInventoryOfItemsInBizItemGroupResponse {
	s.Body = v
	return s
}

type QueryItemDetailRequest struct {
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ItemId   *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
}

func (s QueryItemDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryItemDetailRequest) SetBizId(v string) *QueryItemDetailRequest {
	s.BizId = &v
	return s
}

func (s *QueryItemDetailRequest) SetItemId(v int64) *QueryItemDetailRequest {
	s.ItemId = &v
	return s
}

func (s *QueryItemDetailRequest) SetLmItemId(v string) *QueryItemDetailRequest {
	s.LmItemId = &v
	return s
}

type QueryItemDetailResponseBody struct {
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                          `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                          `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	Item       *QueryItemDetailResponseBodyItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Struct"`
}

func (s QueryItemDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryItemDetailResponseBody) SetRequestId(v string) *QueryItemDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetSuccess(v bool) *QueryItemDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetSubMessage(v string) *QueryItemDetailResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetCode(v string) *QueryItemDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetMessage(v string) *QueryItemDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetSubCode(v string) *QueryItemDetailResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetItem(v *QueryItemDetailResponseBodyItem) *QueryItemDetailResponseBody {
	s.Item = v
	return s
}

type QueryItemDetailResponseBodyItem struct {
	ItemTitle              *string                                    `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	MinPoints              *int64                                     `json:"MinPoints,omitempty" xml:"MinPoints,omitempty"`
	DescOption             *string                                    `json:"DescOption,omitempty" xml:"DescOption,omitempty"`
	VideoPicUrl            *string                                    `json:"VideoPicUrl,omitempty" xml:"VideoPicUrl,omitempty"`
	ExtJson                *string                                    `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	IsSellerPayPostfee     *bool                                      `json:"IsSellerPayPostfee,omitempty" xml:"IsSellerPayPostfee,omitempty"`
	LmItemCategory         *string                                    `json:"LmItemCategory,omitempty" xml:"LmItemCategory,omitempty"`
	SellerPayPostfee       *bool                                      `json:"SellerPayPostfee,omitempty" xml:"SellerPayPostfee,omitempty"`
	ReservePrice           *int64                                     `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	Quantity               *int32                                     `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	VideoUrl               *string                                    `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	CustomizedAttributeMap map[string]interface{}                     `json:"CustomizedAttributeMap,omitempty" xml:"CustomizedAttributeMap,omitempty"`
	IforestPropsJson       *string                                    `json:"IforestPropsJson,omitempty" xml:"IforestPropsJson,omitempty"`
	PropertiesJson         *string                                    `json:"PropertiesJson,omitempty" xml:"PropertiesJson,omitempty"`
	IforestProps           *string                                    `json:"IforestProps,omitempty" xml:"IforestProps,omitempty"`
	LmItemId               *string                                    `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	SellerId               *int64                                     `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	TbShopName             *string                                    `json:"TbShopName,omitempty" xml:"TbShopName,omitempty"`
	ItemId                 *int64                                     `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	CanSell                *bool                                      `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	CenterInventory        *bool                                      `json:"CenterInventory,omitempty" xml:"CenterInventory,omitempty"`
	SellerType             *int32                                     `json:"SellerType,omitempty" xml:"SellerType,omitempty"`
	TotalSoldQuantity      *int32                                     `json:"TotalSoldQuantity,omitempty" xml:"TotalSoldQuantity,omitempty"`
	MainPicUrl             *string                                    `json:"MainPicUrl,omitempty" xml:"MainPicUrl,omitempty"`
	MinPrice               *int64                                     `json:"MinPrice,omitempty" xml:"MinPrice,omitempty"`
	IsCanSell              *bool                                      `json:"IsCanSell,omitempty" xml:"IsCanSell,omitempty"`
	CategoryId             *int64                                     `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	DescPath               *string                                    `json:"DescPath,omitempty" xml:"DescPath,omitempty"`
	Properties             *string                                    `json:"Properties,omitempty" xml:"Properties,omitempty"`
	Skus                   *QueryItemDetailResponseBodyItemSkus       `json:"Skus,omitempty" xml:"Skus,omitempty" type:"Struct"`
	ItemImages             *QueryItemDetailResponseBodyItemItemImages `json:"ItemImages,omitempty" xml:"ItemImages,omitempty" type:"Struct"`
}

func (s QueryItemDetailResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailResponseBodyItem) GoString() string {
	return s.String()
}

func (s *QueryItemDetailResponseBodyItem) SetItemTitle(v string) *QueryItemDetailResponseBodyItem {
	s.ItemTitle = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetMinPoints(v int64) *QueryItemDetailResponseBodyItem {
	s.MinPoints = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetDescOption(v string) *QueryItemDetailResponseBodyItem {
	s.DescOption = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetVideoPicUrl(v string) *QueryItemDetailResponseBodyItem {
	s.VideoPicUrl = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetExtJson(v string) *QueryItemDetailResponseBodyItem {
	s.ExtJson = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetIsSellerPayPostfee(v bool) *QueryItemDetailResponseBodyItem {
	s.IsSellerPayPostfee = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetLmItemCategory(v string) *QueryItemDetailResponseBodyItem {
	s.LmItemCategory = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetSellerPayPostfee(v bool) *QueryItemDetailResponseBodyItem {
	s.SellerPayPostfee = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetReservePrice(v int64) *QueryItemDetailResponseBodyItem {
	s.ReservePrice = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetQuantity(v int32) *QueryItemDetailResponseBodyItem {
	s.Quantity = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetVideoUrl(v string) *QueryItemDetailResponseBodyItem {
	s.VideoUrl = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetCustomizedAttributeMap(v map[string]interface{}) *QueryItemDetailResponseBodyItem {
	s.CustomizedAttributeMap = v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetIforestPropsJson(v string) *QueryItemDetailResponseBodyItem {
	s.IforestPropsJson = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetPropertiesJson(v string) *QueryItemDetailResponseBodyItem {
	s.PropertiesJson = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetIforestProps(v string) *QueryItemDetailResponseBodyItem {
	s.IforestProps = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetLmItemId(v string) *QueryItemDetailResponseBodyItem {
	s.LmItemId = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetSellerId(v int64) *QueryItemDetailResponseBodyItem {
	s.SellerId = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetTbShopName(v string) *QueryItemDetailResponseBodyItem {
	s.TbShopName = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetItemId(v int64) *QueryItemDetailResponseBodyItem {
	s.ItemId = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetCanSell(v bool) *QueryItemDetailResponseBodyItem {
	s.CanSell = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetCenterInventory(v bool) *QueryItemDetailResponseBodyItem {
	s.CenterInventory = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetSellerType(v int32) *QueryItemDetailResponseBodyItem {
	s.SellerType = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetTotalSoldQuantity(v int32) *QueryItemDetailResponseBodyItem {
	s.TotalSoldQuantity = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetMainPicUrl(v string) *QueryItemDetailResponseBodyItem {
	s.MainPicUrl = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetMinPrice(v int64) *QueryItemDetailResponseBodyItem {
	s.MinPrice = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetIsCanSell(v bool) *QueryItemDetailResponseBodyItem {
	s.IsCanSell = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetCategoryId(v int64) *QueryItemDetailResponseBodyItem {
	s.CategoryId = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetDescPath(v string) *QueryItemDetailResponseBodyItem {
	s.DescPath = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetProperties(v string) *QueryItemDetailResponseBodyItem {
	s.Properties = &v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetSkus(v *QueryItemDetailResponseBodyItemSkus) *QueryItemDetailResponseBodyItem {
	s.Skus = v
	return s
}

func (s *QueryItemDetailResponseBodyItem) SetItemImages(v *QueryItemDetailResponseBodyItemItemImages) *QueryItemDetailResponseBodyItem {
	s.ItemImages = v
	return s
}

type QueryItemDetailResponseBodyItemSkus struct {
	Sku []*QueryItemDetailResponseBodyItemSkusSku `json:"Sku,omitempty" xml:"Sku,omitempty" type:"Repeated"`
}

func (s QueryItemDetailResponseBodyItemSkus) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailResponseBodyItemSkus) GoString() string {
	return s.String()
}

func (s *QueryItemDetailResponseBodyItemSkus) SetSku(v []*QueryItemDetailResponseBodyItemSkusSku) *QueryItemDetailResponseBodyItemSkus {
	s.Sku = v
	return s
}

type QueryItemDetailResponseBodyItemSkusSku struct {
	PriceCent              *int64                 `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	LmItemId               *string                `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	PointPrice             *int64                 `json:"PointPrice,omitempty" xml:"PointPrice,omitempty"`
	CanSell                *bool                  `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	ItemId                 *int64                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	SkuTitle               *string                `json:"SkuTitle,omitempty" xml:"SkuTitle,omitempty"`
	SkuPropertiesJson      *string                `json:"SkuPropertiesJson,omitempty" xml:"SkuPropertiesJson,omitempty"`
	ExtJson                *string                `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	SkuProperties          *string                `json:"SkuProperties,omitempty" xml:"SkuProperties,omitempty"`
	SkuId                  *int64                 `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuPicUrl              *string                `json:"SkuPicUrl,omitempty" xml:"SkuPicUrl,omitempty"`
	Points                 *int64                 `json:"Points,omitempty" xml:"Points,omitempty"`
	ReservePrice           *int64                 `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	PointsAmount           *int64                 `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	Quantity               *int32                 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	CustomizedAttributeMap map[string]interface{} `json:"CustomizedAttributeMap,omitempty" xml:"CustomizedAttributeMap,omitempty"`
}

func (s QueryItemDetailResponseBodyItemSkusSku) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailResponseBodyItemSkusSku) GoString() string {
	return s.String()
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetPriceCent(v int64) *QueryItemDetailResponseBodyItemSkusSku {
	s.PriceCent = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetLmItemId(v string) *QueryItemDetailResponseBodyItemSkusSku {
	s.LmItemId = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetPointPrice(v int64) *QueryItemDetailResponseBodyItemSkusSku {
	s.PointPrice = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetCanSell(v bool) *QueryItemDetailResponseBodyItemSkusSku {
	s.CanSell = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetItemId(v int64) *QueryItemDetailResponseBodyItemSkusSku {
	s.ItemId = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetSkuTitle(v string) *QueryItemDetailResponseBodyItemSkusSku {
	s.SkuTitle = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetSkuPropertiesJson(v string) *QueryItemDetailResponseBodyItemSkusSku {
	s.SkuPropertiesJson = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetExtJson(v string) *QueryItemDetailResponseBodyItemSkusSku {
	s.ExtJson = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetSkuProperties(v string) *QueryItemDetailResponseBodyItemSkusSku {
	s.SkuProperties = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetSkuId(v int64) *QueryItemDetailResponseBodyItemSkusSku {
	s.SkuId = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetSkuPicUrl(v string) *QueryItemDetailResponseBodyItemSkusSku {
	s.SkuPicUrl = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetPoints(v int64) *QueryItemDetailResponseBodyItemSkusSku {
	s.Points = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetReservePrice(v int64) *QueryItemDetailResponseBodyItemSkusSku {
	s.ReservePrice = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetPointsAmount(v int64) *QueryItemDetailResponseBodyItemSkusSku {
	s.PointsAmount = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetQuantity(v int32) *QueryItemDetailResponseBodyItemSkusSku {
	s.Quantity = &v
	return s
}

func (s *QueryItemDetailResponseBodyItemSkusSku) SetCustomizedAttributeMap(v map[string]interface{}) *QueryItemDetailResponseBodyItemSkusSku {
	s.CustomizedAttributeMap = v
	return s
}

type QueryItemDetailResponseBodyItemItemImages struct {
	ItemImage []*string `json:"ItemImage,omitempty" xml:"ItemImage,omitempty" type:"Repeated"`
}

func (s QueryItemDetailResponseBodyItemItemImages) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailResponseBodyItemItemImages) GoString() string {
	return s.String()
}

func (s *QueryItemDetailResponseBodyItemItemImages) SetItemImage(v []*string) *QueryItemDetailResponseBodyItemItemImages {
	s.ItemImage = v
	return s
}

type QueryItemDetailResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryItemDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryItemDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryItemDetailResponse) SetHeaders(v map[string]*string) *QueryItemDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryItemDetailResponse) SetBody(v *QueryItemDetailResponseBody) *QueryItemDetailResponse {
	s.Body = v
	return s
}

type QueryItemDetailInnerRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ItemId                *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId              *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	DivisionCode          *string `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
	Ip                    *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s QueryItemDetailInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailInnerRequest) GoString() string {
	return s.String()
}

func (s *QueryItemDetailInnerRequest) SetBizId(v string) *QueryItemDetailInnerRequest {
	s.BizId = &v
	return s
}

func (s *QueryItemDetailInnerRequest) SetItemId(v int64) *QueryItemDetailInnerRequest {
	s.ItemId = &v
	return s
}

func (s *QueryItemDetailInnerRequest) SetLmItemId(v string) *QueryItemDetailInnerRequest {
	s.LmItemId = &v
	return s
}

func (s *QueryItemDetailInnerRequest) SetUseAnonymousTbAccount(v bool) *QueryItemDetailInnerRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *QueryItemDetailInnerRequest) SetThirdPartyUserId(v string) *QueryItemDetailInnerRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *QueryItemDetailInnerRequest) SetBizUid(v string) *QueryItemDetailInnerRequest {
	s.BizUid = &v
	return s
}

func (s *QueryItemDetailInnerRequest) SetDivisionCode(v string) *QueryItemDetailInnerRequest {
	s.DivisionCode = &v
	return s
}

func (s *QueryItemDetailInnerRequest) SetIp(v string) *QueryItemDetailInnerRequest {
	s.Ip = &v
	return s
}

type QueryItemDetailInnerResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Item      *QueryItemDetailInnerResponseBodyItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Struct"`
}

func (s QueryItemDetailInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailInnerResponseBody) GoString() string {
	return s.String()
}

func (s *QueryItemDetailInnerResponseBody) SetCode(v string) *QueryItemDetailInnerResponseBody {
	s.Code = &v
	return s
}

func (s *QueryItemDetailInnerResponseBody) SetMessage(v string) *QueryItemDetailInnerResponseBody {
	s.Message = &v
	return s
}

func (s *QueryItemDetailInnerResponseBody) SetRequestId(v string) *QueryItemDetailInnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryItemDetailInnerResponseBody) SetItem(v *QueryItemDetailInnerResponseBodyItem) *QueryItemDetailInnerResponseBody {
	s.Item = v
	return s
}

type QueryItemDetailInnerResponseBodyItem struct {
	ItemTitle          *string                                             `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	City               *string                                             `json:"City,omitempty" xml:"City,omitempty"`
	MinPoints          *int64                                              `json:"MinPoints,omitempty" xml:"MinPoints,omitempty"`
	DescOption         *string                                             `json:"DescOption,omitempty" xml:"DescOption,omitempty"`
	IsSellerPayPostfee *bool                                               `json:"IsSellerPayPostfee,omitempty" xml:"IsSellerPayPostfee,omitempty"`
	LmItemCategory     *string                                             `json:"LmItemCategory,omitempty" xml:"LmItemCategory,omitempty"`
	SellerPayPostfee   *bool                                               `json:"SellerPayPostfee,omitempty" xml:"SellerPayPostfee,omitempty"`
	ReservePrice       *int64                                              `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	Quantity           *int32                                              `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	SellerId           *int64                                              `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	TbShopName         *string                                             `json:"TbShopName,omitempty" xml:"TbShopName,omitempty"`
	ItemId             *int64                                              `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	CanSell            *bool                                               `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	CenterInventory    *bool                                               `json:"CenterInventory,omitempty" xml:"CenterInventory,omitempty"`
	TotalSoldQuantity  *int32                                              `json:"TotalSoldQuantity,omitempty" xml:"TotalSoldQuantity,omitempty"`
	MainPicUrl         *string                                             `json:"MainPicUrl,omitempty" xml:"MainPicUrl,omitempty"`
	MinPrice           *int64                                              `json:"MinPrice,omitempty" xml:"MinPrice,omitempty"`
	IsCanSell          *bool                                               `json:"IsCanSell,omitempty" xml:"IsCanSell,omitempty"`
	LmShopId           *int64                                              `json:"LmShopId,omitempty" xml:"LmShopId,omitempty"`
	CategoryId         *int64                                              `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	SellerNick         *string                                             `json:"SellerNick,omitempty" xml:"SellerNick,omitempty"`
	DescPath           *string                                             `json:"DescPath,omitempty" xml:"DescPath,omitempty"`
	Properties         map[string]interface{}                              `json:"Properties,omitempty" xml:"Properties,omitempty"`
	Province           *string                                             `json:"Province,omitempty" xml:"Province,omitempty"`
	Skus               []*QueryItemDetailInnerResponseBodyItemSkus         `json:"Skus,omitempty" xml:"Skus,omitempty" type:"Repeated"`
	SkuPropertys       []*QueryItemDetailInnerResponseBodyItemSkuPropertys `json:"SkuPropertys,omitempty" xml:"SkuPropertys,omitempty" type:"Repeated"`
	CategoryIds        []*int64                                            `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	IforestProps       []map[string]interface{}                            `json:"IforestProps,omitempty" xml:"IforestProps,omitempty" type:"Repeated"`
	ItemImages         []*string                                           `json:"ItemImages,omitempty" xml:"ItemImages,omitempty" type:"Repeated"`
}

func (s QueryItemDetailInnerResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailInnerResponseBodyItem) GoString() string {
	return s.String()
}

func (s *QueryItemDetailInnerResponseBodyItem) SetItemTitle(v string) *QueryItemDetailInnerResponseBodyItem {
	s.ItemTitle = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetCity(v string) *QueryItemDetailInnerResponseBodyItem {
	s.City = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetMinPoints(v int64) *QueryItemDetailInnerResponseBodyItem {
	s.MinPoints = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetDescOption(v string) *QueryItemDetailInnerResponseBodyItem {
	s.DescOption = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetIsSellerPayPostfee(v bool) *QueryItemDetailInnerResponseBodyItem {
	s.IsSellerPayPostfee = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetLmItemCategory(v string) *QueryItemDetailInnerResponseBodyItem {
	s.LmItemCategory = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetSellerPayPostfee(v bool) *QueryItemDetailInnerResponseBodyItem {
	s.SellerPayPostfee = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetReservePrice(v int64) *QueryItemDetailInnerResponseBodyItem {
	s.ReservePrice = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetQuantity(v int32) *QueryItemDetailInnerResponseBodyItem {
	s.Quantity = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetSellerId(v int64) *QueryItemDetailInnerResponseBodyItem {
	s.SellerId = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetTbShopName(v string) *QueryItemDetailInnerResponseBodyItem {
	s.TbShopName = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetItemId(v int64) *QueryItemDetailInnerResponseBodyItem {
	s.ItemId = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetCanSell(v bool) *QueryItemDetailInnerResponseBodyItem {
	s.CanSell = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetCenterInventory(v bool) *QueryItemDetailInnerResponseBodyItem {
	s.CenterInventory = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetTotalSoldQuantity(v int32) *QueryItemDetailInnerResponseBodyItem {
	s.TotalSoldQuantity = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetMainPicUrl(v string) *QueryItemDetailInnerResponseBodyItem {
	s.MainPicUrl = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetMinPrice(v int64) *QueryItemDetailInnerResponseBodyItem {
	s.MinPrice = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetIsCanSell(v bool) *QueryItemDetailInnerResponseBodyItem {
	s.IsCanSell = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetLmShopId(v int64) *QueryItemDetailInnerResponseBodyItem {
	s.LmShopId = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetCategoryId(v int64) *QueryItemDetailInnerResponseBodyItem {
	s.CategoryId = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetSellerNick(v string) *QueryItemDetailInnerResponseBodyItem {
	s.SellerNick = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetDescPath(v string) *QueryItemDetailInnerResponseBodyItem {
	s.DescPath = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetProperties(v map[string]interface{}) *QueryItemDetailInnerResponseBodyItem {
	s.Properties = v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetProvince(v string) *QueryItemDetailInnerResponseBodyItem {
	s.Province = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetSkus(v []*QueryItemDetailInnerResponseBodyItemSkus) *QueryItemDetailInnerResponseBodyItem {
	s.Skus = v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetSkuPropertys(v []*QueryItemDetailInnerResponseBodyItemSkuPropertys) *QueryItemDetailInnerResponseBodyItem {
	s.SkuPropertys = v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetCategoryIds(v []*int64) *QueryItemDetailInnerResponseBodyItem {
	s.CategoryIds = v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetIforestProps(v []map[string]interface{}) *QueryItemDetailInnerResponseBodyItem {
	s.IforestProps = v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItem) SetItemImages(v []*string) *QueryItemDetailInnerResponseBodyItem {
	s.ItemImages = v
	return s
}

type QueryItemDetailInnerResponseBodyItemSkus struct {
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	PriceCent    *int64  `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	SkuPvs       *string `json:"SkuPvs,omitempty" xml:"SkuPvs,omitempty"`
	LmItemId     *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	PointPrice   *int64  `json:"PointPrice,omitempty" xml:"PointPrice,omitempty"`
	ItemId       *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	SkuTitle     *string `json:"SkuTitle,omitempty" xml:"SkuTitle,omitempty"`
	ExtJson      *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	SkuId        *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuPicUrl    *string `json:"SkuPicUrl,omitempty" xml:"SkuPicUrl,omitempty"`
	Points       *int64  `json:"Points,omitempty" xml:"Points,omitempty"`
	PointsAmount *int64  `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	ReservePrice *int64  `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	SkuDesc      *string `json:"SkuDesc,omitempty" xml:"SkuDesc,omitempty"`
	Quantity     *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s QueryItemDetailInnerResponseBodyItemSkus) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailInnerResponseBodyItemSkus) GoString() string {
	return s.String()
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetStatus(v int32) *QueryItemDetailInnerResponseBodyItemSkus {
	s.Status = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetPriceCent(v int64) *QueryItemDetailInnerResponseBodyItemSkus {
	s.PriceCent = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetSkuPvs(v string) *QueryItemDetailInnerResponseBodyItemSkus {
	s.SkuPvs = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetLmItemId(v string) *QueryItemDetailInnerResponseBodyItemSkus {
	s.LmItemId = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetPointPrice(v int64) *QueryItemDetailInnerResponseBodyItemSkus {
	s.PointPrice = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetItemId(v int64) *QueryItemDetailInnerResponseBodyItemSkus {
	s.ItemId = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetSkuTitle(v string) *QueryItemDetailInnerResponseBodyItemSkus {
	s.SkuTitle = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetExtJson(v string) *QueryItemDetailInnerResponseBodyItemSkus {
	s.ExtJson = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetSkuId(v int64) *QueryItemDetailInnerResponseBodyItemSkus {
	s.SkuId = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetSkuPicUrl(v string) *QueryItemDetailInnerResponseBodyItemSkus {
	s.SkuPicUrl = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetPoints(v int64) *QueryItemDetailInnerResponseBodyItemSkus {
	s.Points = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetPointsAmount(v int64) *QueryItemDetailInnerResponseBodyItemSkus {
	s.PointsAmount = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetReservePrice(v int64) *QueryItemDetailInnerResponseBodyItemSkus {
	s.ReservePrice = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetSkuDesc(v string) *QueryItemDetailInnerResponseBodyItemSkus {
	s.SkuDesc = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkus) SetQuantity(v int32) *QueryItemDetailInnerResponseBodyItemSkus {
	s.Quantity = &v
	return s
}

type QueryItemDetailInnerResponseBodyItemSkuPropertys struct {
	Text   *string                                                   `json:"Text,omitempty" xml:"Text,omitempty"`
	Id     *int64                                                    `json:"Id,omitempty" xml:"Id,omitempty"`
	Values []*QueryItemDetailInnerResponseBodyItemSkuPropertysValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryItemDetailInnerResponseBodyItemSkuPropertys) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailInnerResponseBodyItemSkuPropertys) GoString() string {
	return s.String()
}

func (s *QueryItemDetailInnerResponseBodyItemSkuPropertys) SetText(v string) *QueryItemDetailInnerResponseBodyItemSkuPropertys {
	s.Text = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkuPropertys) SetId(v int64) *QueryItemDetailInnerResponseBodyItemSkuPropertys {
	s.Id = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkuPropertys) SetValues(v []*QueryItemDetailInnerResponseBodyItemSkuPropertysValues) *QueryItemDetailInnerResponseBodyItemSkuPropertys {
	s.Values = v
	return s
}

type QueryItemDetailInnerResponseBodyItemSkuPropertysValues struct {
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QueryItemDetailInnerResponseBodyItemSkuPropertysValues) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailInnerResponseBodyItemSkuPropertysValues) GoString() string {
	return s.String()
}

func (s *QueryItemDetailInnerResponseBodyItemSkuPropertysValues) SetText(v string) *QueryItemDetailInnerResponseBodyItemSkuPropertysValues {
	s.Text = &v
	return s
}

func (s *QueryItemDetailInnerResponseBodyItemSkuPropertysValues) SetId(v int64) *QueryItemDetailInnerResponseBodyItemSkuPropertysValues {
	s.Id = &v
	return s
}

type QueryItemDetailInnerResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryItemDetailInnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryItemDetailInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailInnerResponse) GoString() string {
	return s.String()
}

func (s *QueryItemDetailInnerResponse) SetHeaders(v map[string]*string) *QueryItemDetailInnerResponse {
	s.Headers = v
	return s
}

func (s *QueryItemDetailInnerResponse) SetBody(v *QueryItemDetailInnerResponseBody) *QueryItemDetailInnerResponse {
	s.Body = v
	return s
}

type QueryItemInSubBizsRequest struct {
	BizId     *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ItemId    *int64                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId  *string                `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	SubBizIds map[string]interface{} `json:"SubBizIds,omitempty" xml:"SubBizIds,omitempty"`
}

func (s QueryItemInSubBizsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInSubBizsRequest) GoString() string {
	return s.String()
}

func (s *QueryItemInSubBizsRequest) SetBizId(v string) *QueryItemInSubBizsRequest {
	s.BizId = &v
	return s
}

func (s *QueryItemInSubBizsRequest) SetItemId(v int64) *QueryItemInSubBizsRequest {
	s.ItemId = &v
	return s
}

func (s *QueryItemInSubBizsRequest) SetLmItemId(v string) *QueryItemInSubBizsRequest {
	s.LmItemId = &v
	return s
}

func (s *QueryItemInSubBizsRequest) SetSubBizIds(v map[string]interface{}) *QueryItemInSubBizsRequest {
	s.SubBizIds = v
	return s
}

type QueryItemInSubBizsShrinkRequest struct {
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ItemId          *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId        *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	SubBizIdsShrink *string `json:"SubBizIds,omitempty" xml:"SubBizIds,omitempty"`
}

func (s QueryItemInSubBizsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInSubBizsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryItemInSubBizsShrinkRequest) SetBizId(v string) *QueryItemInSubBizsShrinkRequest {
	s.BizId = &v
	return s
}

func (s *QueryItemInSubBizsShrinkRequest) SetItemId(v int64) *QueryItemInSubBizsShrinkRequest {
	s.ItemId = &v
	return s
}

func (s *QueryItemInSubBizsShrinkRequest) SetLmItemId(v string) *QueryItemInSubBizsShrinkRequest {
	s.LmItemId = &v
	return s
}

func (s *QueryItemInSubBizsShrinkRequest) SetSubBizIdsShrink(v string) *QueryItemInSubBizsShrinkRequest {
	s.SubBizIdsShrink = &v
	return s
}

type QueryItemInSubBizsResponseBody struct {
	Code        *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ItemBizList []*QueryItemInSubBizsResponseBodyItemBizList `json:"ItemBizList,omitempty" xml:"ItemBizList,omitempty" type:"Repeated"`
}

func (s QueryItemInSubBizsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInSubBizsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryItemInSubBizsResponseBody) SetCode(v string) *QueryItemInSubBizsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryItemInSubBizsResponseBody) SetMessage(v string) *QueryItemInSubBizsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryItemInSubBizsResponseBody) SetRequestId(v string) *QueryItemInSubBizsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryItemInSubBizsResponseBody) SetItemBizList(v []*QueryItemInSubBizsResponseBodyItemBizList) *QueryItemInSubBizsResponseBody {
	s.ItemBizList = v
	return s
}

type QueryItemInSubBizsResponseBodyItemBizList struct {
	ItemTitle      *string                                             `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	PropertiesJson *string                                             `json:"PropertiesJson,omitempty" xml:"PropertiesJson,omitempty"`
	LmItemId       *string                                             `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	SellerId       *int64                                              `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	TbShopName     *string                                             `json:"TbShopName,omitempty" xml:"TbShopName,omitempty"`
	CanSell        *bool                                               `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	ItemId         *int64                                              `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	DescOption     *string                                             `json:"DescOption,omitempty" xml:"DescOption,omitempty"`
	MainPicUrl     *string                                             `json:"MainPicUrl,omitempty" xml:"MainPicUrl,omitempty"`
	CategoryId     *int64                                              `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	SubBizId       *string                                             `json:"SubBizId,omitempty" xml:"SubBizId,omitempty"`
	ReservePrice   *int64                                              `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	Quantity       *int64                                              `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	SkuList        []*QueryItemInSubBizsResponseBodyItemBizListSkuList `json:"SkuList,omitempty" xml:"SkuList,omitempty" type:"Repeated"`
	ItemImages     []*string                                           `json:"ItemImages,omitempty" xml:"ItemImages,omitempty" type:"Repeated"`
}

func (s QueryItemInSubBizsResponseBodyItemBizList) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInSubBizsResponseBodyItemBizList) GoString() string {
	return s.String()
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetItemTitle(v string) *QueryItemInSubBizsResponseBodyItemBizList {
	s.ItemTitle = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetPropertiesJson(v string) *QueryItemInSubBizsResponseBodyItemBizList {
	s.PropertiesJson = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetLmItemId(v string) *QueryItemInSubBizsResponseBodyItemBizList {
	s.LmItemId = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetSellerId(v int64) *QueryItemInSubBizsResponseBodyItemBizList {
	s.SellerId = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetTbShopName(v string) *QueryItemInSubBizsResponseBodyItemBizList {
	s.TbShopName = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetCanSell(v bool) *QueryItemInSubBizsResponseBodyItemBizList {
	s.CanSell = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetItemId(v int64) *QueryItemInSubBizsResponseBodyItemBizList {
	s.ItemId = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetDescOption(v string) *QueryItemInSubBizsResponseBodyItemBizList {
	s.DescOption = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetMainPicUrl(v string) *QueryItemInSubBizsResponseBodyItemBizList {
	s.MainPicUrl = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetCategoryId(v int64) *QueryItemInSubBizsResponseBodyItemBizList {
	s.CategoryId = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetSubBizId(v string) *QueryItemInSubBizsResponseBodyItemBizList {
	s.SubBizId = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetReservePrice(v int64) *QueryItemInSubBizsResponseBodyItemBizList {
	s.ReservePrice = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetQuantity(v int64) *QueryItemInSubBizsResponseBodyItemBizList {
	s.Quantity = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetSkuList(v []*QueryItemInSubBizsResponseBodyItemBizListSkuList) *QueryItemInSubBizsResponseBodyItemBizList {
	s.SkuList = v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizList) SetItemImages(v []*string) *QueryItemInSubBizsResponseBodyItemBizList {
	s.ItemImages = v
	return s
}

type QueryItemInSubBizsResponseBodyItemBizListSkuList struct {
	PriceCent              *int64                 `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	PointPrice             *int64                 `json:"PointPrice,omitempty" xml:"PointPrice,omitempty"`
	CanSell                *bool                  `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	SkuTitle               *string                `json:"SkuTitle,omitempty" xml:"SkuTitle,omitempty"`
	SkuPropertiesJson      *string                `json:"SkuPropertiesJson,omitempty" xml:"SkuPropertiesJson,omitempty"`
	SkuProperties          *string                `json:"SkuProperties,omitempty" xml:"SkuProperties,omitempty"`
	SkuId                  *int64                 `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuPicUrl              *string                `json:"SkuPicUrl,omitempty" xml:"SkuPicUrl,omitempty"`
	Points                 *int64                 `json:"Points,omitempty" xml:"Points,omitempty"`
	PointsAmount           *int64                 `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	ReservePrice           *int64                 `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	BenefitId              *string                `json:"BenefitId,omitempty" xml:"BenefitId,omitempty"`
	CustomizedAttributeMap map[string]interface{} `json:"CustomizedAttributeMap,omitempty" xml:"CustomizedAttributeMap,omitempty"`
}

func (s QueryItemInSubBizsResponseBodyItemBizListSkuList) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInSubBizsResponseBodyItemBizListSkuList) GoString() string {
	return s.String()
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetPriceCent(v int64) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.PriceCent = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetPointPrice(v int64) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.PointPrice = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetCanSell(v bool) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.CanSell = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetSkuTitle(v string) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.SkuTitle = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetSkuPropertiesJson(v string) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.SkuPropertiesJson = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetSkuProperties(v string) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.SkuProperties = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetSkuId(v int64) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.SkuId = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetSkuPicUrl(v string) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.SkuPicUrl = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetPoints(v int64) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.Points = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetPointsAmount(v int64) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.PointsAmount = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetReservePrice(v int64) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.ReservePrice = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetBenefitId(v string) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.BenefitId = &v
	return s
}

func (s *QueryItemInSubBizsResponseBodyItemBizListSkuList) SetCustomizedAttributeMap(v map[string]interface{}) *QueryItemInSubBizsResponseBodyItemBizListSkuList {
	s.CustomizedAttributeMap = v
	return s
}

type QueryItemInSubBizsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryItemInSubBizsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryItemInSubBizsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInSubBizsResponse) GoString() string {
	return s.String()
}

func (s *QueryItemInSubBizsResponse) SetHeaders(v map[string]*string) *QueryItemInSubBizsResponse {
	s.Headers = v
	return s
}

func (s *QueryItemInSubBizsResponse) SetBody(v *QueryItemInSubBizsResponseBody) *QueryItemInSubBizsResponse {
	s.Body = v
	return s
}

type QueryItemInventoryRequest struct {
	BizId        *string                              `json:"BizId,omitempty" xml:"BizId,omitempty"`
	DivisionCode *string                              `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
	Ip           *string                              `json:"Ip,omitempty" xml:"Ip,omitempty"`
	ItemList     []*QueryItemInventoryRequestItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
}

func (s QueryItemInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInventoryRequest) GoString() string {
	return s.String()
}

func (s *QueryItemInventoryRequest) SetBizId(v string) *QueryItemInventoryRequest {
	s.BizId = &v
	return s
}

func (s *QueryItemInventoryRequest) SetDivisionCode(v string) *QueryItemInventoryRequest {
	s.DivisionCode = &v
	return s
}

func (s *QueryItemInventoryRequest) SetIp(v string) *QueryItemInventoryRequest {
	s.Ip = &v
	return s
}

func (s *QueryItemInventoryRequest) SetItemList(v []*QueryItemInventoryRequestItemList) *QueryItemInventoryRequest {
	s.ItemList = v
	return s
}

type QueryItemInventoryRequestItemList struct {
	LmItemId  *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId    *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	SkuIdList []*int  `json:"SkuIdList,omitempty" xml:"SkuIdList,omitempty" type:"Repeated"`
}

func (s QueryItemInventoryRequestItemList) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInventoryRequestItemList) GoString() string {
	return s.String()
}

func (s *QueryItemInventoryRequestItemList) SetLmItemId(v string) *QueryItemInventoryRequestItemList {
	s.LmItemId = &v
	return s
}

func (s *QueryItemInventoryRequestItemList) SetItemId(v int64) *QueryItemInventoryRequestItemList {
	s.ItemId = &v
	return s
}

func (s *QueryItemInventoryRequestItemList) SetSkuIdList(v []*int) *QueryItemInventoryRequestItemList {
	s.SkuIdList = v
	return s
}

type QueryItemInventoryResponseBody struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                                 `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                                 `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	ItemList   *QueryItemInventoryResponseBodyItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Struct"`
}

func (s QueryItemInventoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryItemInventoryResponseBody) SetRequestId(v string) *QueryItemInventoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryItemInventoryResponseBody) SetSuccess(v bool) *QueryItemInventoryResponseBody {
	s.Success = &v
	return s
}

func (s *QueryItemInventoryResponseBody) SetSubMessage(v string) *QueryItemInventoryResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryItemInventoryResponseBody) SetCode(v string) *QueryItemInventoryResponseBody {
	s.Code = &v
	return s
}

func (s *QueryItemInventoryResponseBody) SetMessage(v string) *QueryItemInventoryResponseBody {
	s.Message = &v
	return s
}

func (s *QueryItemInventoryResponseBody) SetSubCode(v string) *QueryItemInventoryResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryItemInventoryResponseBody) SetItemList(v *QueryItemInventoryResponseBodyItemList) *QueryItemInventoryResponseBody {
	s.ItemList = v
	return s
}

type QueryItemInventoryResponseBodyItemList struct {
	Item []*QueryItemInventoryResponseBodyItemListItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryItemInventoryResponseBodyItemList) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInventoryResponseBodyItemList) GoString() string {
	return s.String()
}

func (s *QueryItemInventoryResponseBodyItemList) SetItem(v []*QueryItemInventoryResponseBodyItemListItem) *QueryItemInventoryResponseBodyItemList {
	s.Item = v
	return s
}

type QueryItemInventoryResponseBodyItemListItem struct {
	LmItemId *string                                            `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId   *int64                                             `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	SkuList  *QueryItemInventoryResponseBodyItemListItemSkuList `json:"SkuList,omitempty" xml:"SkuList,omitempty" type:"Struct"`
}

func (s QueryItemInventoryResponseBodyItemListItem) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInventoryResponseBodyItemListItem) GoString() string {
	return s.String()
}

func (s *QueryItemInventoryResponseBodyItemListItem) SetLmItemId(v string) *QueryItemInventoryResponseBodyItemListItem {
	s.LmItemId = &v
	return s
}

func (s *QueryItemInventoryResponseBodyItemListItem) SetItemId(v int64) *QueryItemInventoryResponseBodyItemListItem {
	s.ItemId = &v
	return s
}

func (s *QueryItemInventoryResponseBodyItemListItem) SetSkuList(v *QueryItemInventoryResponseBodyItemListItemSkuList) *QueryItemInventoryResponseBodyItemListItem {
	s.SkuList = v
	return s
}

type QueryItemInventoryResponseBodyItemListItemSkuList struct {
	Sku []*QueryItemInventoryResponseBodyItemListItemSkuListSku `json:"Sku,omitempty" xml:"Sku,omitempty" type:"Repeated"`
}

func (s QueryItemInventoryResponseBodyItemListItemSkuList) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInventoryResponseBodyItemListItemSkuList) GoString() string {
	return s.String()
}

func (s *QueryItemInventoryResponseBodyItemListItemSkuList) SetSku(v []*QueryItemInventoryResponseBodyItemListItemSkuListSku) *QueryItemInventoryResponseBodyItemListItemSkuList {
	s.Sku = v
	return s
}

type QueryItemInventoryResponseBodyItemListItemSkuListSku struct {
	SkuId     *int64                                                         `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Inventory *QueryItemInventoryResponseBodyItemListItemSkuListSkuInventory `json:"Inventory,omitempty" xml:"Inventory,omitempty" type:"Struct"`
}

func (s QueryItemInventoryResponseBodyItemListItemSkuListSku) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInventoryResponseBodyItemListItemSkuListSku) GoString() string {
	return s.String()
}

func (s *QueryItemInventoryResponseBodyItemListItemSkuListSku) SetSkuId(v int64) *QueryItemInventoryResponseBodyItemListItemSkuListSku {
	s.SkuId = &v
	return s
}

func (s *QueryItemInventoryResponseBodyItemListItemSkuListSku) SetInventory(v *QueryItemInventoryResponseBodyItemListItemSkuListSkuInventory) *QueryItemInventoryResponseBodyItemListItemSkuListSku {
	s.Inventory = v
	return s
}

type QueryItemInventoryResponseBodyItemListItemSkuListSkuInventory struct {
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s QueryItemInventoryResponseBodyItemListItemSkuListSkuInventory) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInventoryResponseBodyItemListItemSkuListSkuInventory) GoString() string {
	return s.String()
}

func (s *QueryItemInventoryResponseBodyItemListItemSkuListSkuInventory) SetQuantity(v int64) *QueryItemInventoryResponseBodyItemListItemSkuListSkuInventory {
	s.Quantity = &v
	return s
}

type QueryItemInventoryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryItemInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryItemInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryItemInventoryResponse) GoString() string {
	return s.String()
}

func (s *QueryItemInventoryResponse) SetHeaders(v map[string]*string) *QueryItemInventoryResponse {
	s.Headers = v
	return s
}

func (s *QueryItemInventoryResponse) SetBody(v *QueryItemInventoryResponseBody) *QueryItemInventoryResponse {
	s.Body = v
	return s
}

type QueryLogisticsRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	LmOrderId             *int64  `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s QueryLogisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLogisticsRequest) GoString() string {
	return s.String()
}

func (s *QueryLogisticsRequest) SetBizId(v string) *QueryLogisticsRequest {
	s.BizId = &v
	return s
}

func (s *QueryLogisticsRequest) SetBizUid(v string) *QueryLogisticsRequest {
	s.BizUid = &v
	return s
}

func (s *QueryLogisticsRequest) SetLmOrderId(v int64) *QueryLogisticsRequest {
	s.LmOrderId = &v
	return s
}

func (s *QueryLogisticsRequest) SetUseAnonymousTbAccount(v bool) *QueryLogisticsRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *QueryLogisticsRequest) SetThirdPartyUserId(v string) *QueryLogisticsRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *QueryLogisticsRequest) SetAccountType(v string) *QueryLogisticsRequest {
	s.AccountType = &v
	return s
}

type QueryLogisticsResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *QueryLogisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryLogisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryLogisticsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLogisticsResponseBody) SetCode(v string) *QueryLogisticsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryLogisticsResponseBody) SetMessage(v string) *QueryLogisticsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryLogisticsResponseBody) SetRequestId(v string) *QueryLogisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLogisticsResponseBody) SetData(v *QueryLogisticsResponseBodyData) *QueryLogisticsResponseBody {
	s.Data = v
	return s
}

type QueryLogisticsResponseBodyData struct {
	Data []*QueryLogisticsResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s QueryLogisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryLogisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryLogisticsResponseBodyData) SetData(v []*QueryLogisticsResponseBodyDataData) *QueryLogisticsResponseBodyData {
	s.Data = v
	return s
}

type QueryLogisticsResponseBodyDataData struct {
	LogisticsCompanyCode *string                                                `json:"LogisticsCompanyCode,omitempty" xml:"LogisticsCompanyCode,omitempty"`
	MailNo               *string                                                `json:"MailNo,omitempty" xml:"MailNo,omitempty"`
	LogisticsCompanyName *string                                                `json:"LogisticsCompanyName,omitempty" xml:"LogisticsCompanyName,omitempty"`
	DataProvider         *string                                                `json:"DataProvider,omitempty" xml:"DataProvider,omitempty"`
	DataProviderTitle    *string                                                `json:"DataProviderTitle,omitempty" xml:"DataProviderTitle,omitempty"`
	LogisticsDetailList  *QueryLogisticsResponseBodyDataDataLogisticsDetailList `json:"LogisticsDetailList,omitempty" xml:"LogisticsDetailList,omitempty" type:"Struct"`
	Goods                *QueryLogisticsResponseBodyDataDataGoods               `json:"Goods,omitempty" xml:"Goods,omitempty" type:"Struct"`
}

func (s QueryLogisticsResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s QueryLogisticsResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *QueryLogisticsResponseBodyDataData) SetLogisticsCompanyCode(v string) *QueryLogisticsResponseBodyDataData {
	s.LogisticsCompanyCode = &v
	return s
}

func (s *QueryLogisticsResponseBodyDataData) SetMailNo(v string) *QueryLogisticsResponseBodyDataData {
	s.MailNo = &v
	return s
}

func (s *QueryLogisticsResponseBodyDataData) SetLogisticsCompanyName(v string) *QueryLogisticsResponseBodyDataData {
	s.LogisticsCompanyName = &v
	return s
}

func (s *QueryLogisticsResponseBodyDataData) SetDataProvider(v string) *QueryLogisticsResponseBodyDataData {
	s.DataProvider = &v
	return s
}

func (s *QueryLogisticsResponseBodyDataData) SetDataProviderTitle(v string) *QueryLogisticsResponseBodyDataData {
	s.DataProviderTitle = &v
	return s
}

func (s *QueryLogisticsResponseBodyDataData) SetLogisticsDetailList(v *QueryLogisticsResponseBodyDataDataLogisticsDetailList) *QueryLogisticsResponseBodyDataData {
	s.LogisticsDetailList = v
	return s
}

func (s *QueryLogisticsResponseBodyDataData) SetGoods(v *QueryLogisticsResponseBodyDataDataGoods) *QueryLogisticsResponseBodyDataData {
	s.Goods = v
	return s
}

type QueryLogisticsResponseBodyDataDataLogisticsDetailList struct {
	LogisticsDetailList []*QueryLogisticsResponseBodyDataDataLogisticsDetailListLogisticsDetailList `json:"LogisticsDetailList,omitempty" xml:"LogisticsDetailList,omitempty" type:"Repeated"`
}

func (s QueryLogisticsResponseBodyDataDataLogisticsDetailList) String() string {
	return tea.Prettify(s)
}

func (s QueryLogisticsResponseBodyDataDataLogisticsDetailList) GoString() string {
	return s.String()
}

func (s *QueryLogisticsResponseBodyDataDataLogisticsDetailList) SetLogisticsDetailList(v []*QueryLogisticsResponseBodyDataDataLogisticsDetailListLogisticsDetailList) *QueryLogisticsResponseBodyDataDataLogisticsDetailList {
	s.LogisticsDetailList = v
	return s
}

type QueryLogisticsResponseBodyDataDataLogisticsDetailListLogisticsDetailList struct {
	OcurrTimeStr *string `json:"OcurrTimeStr,omitempty" xml:"OcurrTimeStr,omitempty"`
	StanderdDesc *string `json:"StanderdDesc,omitempty" xml:"StanderdDesc,omitempty"`
}

func (s QueryLogisticsResponseBodyDataDataLogisticsDetailListLogisticsDetailList) String() string {
	return tea.Prettify(s)
}

func (s QueryLogisticsResponseBodyDataDataLogisticsDetailListLogisticsDetailList) GoString() string {
	return s.String()
}

func (s *QueryLogisticsResponseBodyDataDataLogisticsDetailListLogisticsDetailList) SetOcurrTimeStr(v string) *QueryLogisticsResponseBodyDataDataLogisticsDetailListLogisticsDetailList {
	s.OcurrTimeStr = &v
	return s
}

func (s *QueryLogisticsResponseBodyDataDataLogisticsDetailListLogisticsDetailList) SetStanderdDesc(v string) *QueryLogisticsResponseBodyDataDataLogisticsDetailListLogisticsDetailList {
	s.StanderdDesc = &v
	return s
}

type QueryLogisticsResponseBodyDataDataGoods struct {
	Goods []*QueryLogisticsResponseBodyDataDataGoodsGoods `json:"Goods,omitempty" xml:"Goods,omitempty" type:"Repeated"`
}

func (s QueryLogisticsResponseBodyDataDataGoods) String() string {
	return tea.Prettify(s)
}

func (s QueryLogisticsResponseBodyDataDataGoods) GoString() string {
	return s.String()
}

func (s *QueryLogisticsResponseBodyDataDataGoods) SetGoods(v []*QueryLogisticsResponseBodyDataDataGoodsGoods) *QueryLogisticsResponseBodyDataDataGoods {
	s.Goods = v
	return s
}

type QueryLogisticsResponseBodyDataDataGoodsGoods struct {
	ItemId   *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	GoodName *string `json:"GoodName,omitempty" xml:"GoodName,omitempty"`
	Quantity *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s QueryLogisticsResponseBodyDataDataGoodsGoods) String() string {
	return tea.Prettify(s)
}

func (s QueryLogisticsResponseBodyDataDataGoodsGoods) GoString() string {
	return s.String()
}

func (s *QueryLogisticsResponseBodyDataDataGoodsGoods) SetItemId(v int64) *QueryLogisticsResponseBodyDataDataGoodsGoods {
	s.ItemId = &v
	return s
}

func (s *QueryLogisticsResponseBodyDataDataGoodsGoods) SetGoodName(v string) *QueryLogisticsResponseBodyDataDataGoodsGoods {
	s.GoodName = &v
	return s
}

func (s *QueryLogisticsResponseBodyDataDataGoodsGoods) SetQuantity(v int32) *QueryLogisticsResponseBodyDataDataGoodsGoods {
	s.Quantity = &v
	return s
}

type QueryLogisticsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryLogisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryLogisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLogisticsResponse) GoString() string {
	return s.String()
}

func (s *QueryLogisticsResponse) SetHeaders(v map[string]*string) *QueryLogisticsResponse {
	s.Headers = v
	return s
}

func (s *QueryLogisticsResponse) SetBody(v *QueryLogisticsResponseBody) *QueryLogisticsResponse {
	s.Body = v
	return s
}

type QueryMediaSettleInfoRequest struct {
	MediaName  *string `json:"MediaName,omitempty" xml:"MediaName,omitempty"`
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	SettleNo   *string `json:"SettleNo,omitempty" xml:"SettleNo,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ExtInfo    *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s QueryMediaSettleInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMediaSettleInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryMediaSettleInfoRequest) SetMediaName(v string) *QueryMediaSettleInfoRequest {
	s.MediaName = &v
	return s
}

func (s *QueryMediaSettleInfoRequest) SetChannelId(v string) *QueryMediaSettleInfoRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryMediaSettleInfoRequest) SetSettleNo(v string) *QueryMediaSettleInfoRequest {
	s.SettleNo = &v
	return s
}

func (s *QueryMediaSettleInfoRequest) SetStartTime(v string) *QueryMediaSettleInfoRequest {
	s.StartTime = &v
	return s
}

func (s *QueryMediaSettleInfoRequest) SetEndTime(v string) *QueryMediaSettleInfoRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMediaSettleInfoRequest) SetPageSize(v int32) *QueryMediaSettleInfoRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMediaSettleInfoRequest) SetPageNumber(v int32) *QueryMediaSettleInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryMediaSettleInfoRequest) SetExtInfo(v string) *QueryMediaSettleInfoRequest {
	s.ExtInfo = &v
	return s
}

type QueryMediaSettleInfoResponseBody struct {
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                                `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                                `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                                `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Model      *QueryMediaSettleInfoResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s QueryMediaSettleInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMediaSettleInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMediaSettleInfoResponseBody) SetRequestId(v string) *QueryMediaSettleInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBody) SetSuccess(v bool) *QueryMediaSettleInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBody) SetSubMessage(v string) *QueryMediaSettleInfoResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBody) SetCode(v string) *QueryMediaSettleInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBody) SetMessage(v string) *QueryMediaSettleInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBody) SetSubCode(v string) *QueryMediaSettleInfoResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBody) SetTotalCount(v int64) *QueryMediaSettleInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBody) SetLogsId(v string) *QueryMediaSettleInfoResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBody) SetModel(v *QueryMediaSettleInfoResponseBodyModel) *QueryMediaSettleInfoResponseBody {
	s.Model = v
	return s
}

type QueryMediaSettleInfoResponseBodyModel struct {
	PageSize            *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber          *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount          *int32                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	MediaSettleInfoList []*QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList `json:"MediaSettleInfoList,omitempty" xml:"MediaSettleInfoList,omitempty" type:"Repeated"`
}

func (s QueryMediaSettleInfoResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryMediaSettleInfoResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryMediaSettleInfoResponseBodyModel) SetPageSize(v int32) *QueryMediaSettleInfoResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModel) SetPageNumber(v int32) *QueryMediaSettleInfoResponseBodyModel {
	s.PageNumber = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModel) SetTotalCount(v int32) *QueryMediaSettleInfoResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModel) SetMediaSettleInfoList(v []*QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) *QueryMediaSettleInfoResponseBodyModel {
	s.MediaSettleInfoList = v
	return s
}

type QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList struct {
	EndTime             *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	MediaSettleAmount   *string `json:"MediaSettleAmount,omitempty" xml:"MediaSettleAmount,omitempty"`
	SettleNo            *string `json:"SettleNo,omitempty" xml:"SettleNo,omitempty"`
	SettleStatus        *string `json:"SettleStatus,omitempty" xml:"SettleStatus,omitempty"`
	ChannelId           *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	MediaName           *string `json:"MediaName,omitempty" xml:"MediaName,omitempty"`
	MediaSettleDetailId *string `json:"MediaSettleDetailId,omitempty" xml:"MediaSettleDetailId,omitempty"`
	ModifiedDate        *string `json:"ModifiedDate,omitempty" xml:"ModifiedDate,omitempty"`
	CreateDate          *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	ExtInfo             *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) GoString() string {
	return s.String()
}

func (s *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) SetEndTime(v string) *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList {
	s.EndTime = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) SetStartTime(v string) *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList {
	s.StartTime = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) SetMediaSettleAmount(v string) *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList {
	s.MediaSettleAmount = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) SetSettleNo(v string) *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList {
	s.SettleNo = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) SetSettleStatus(v string) *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList {
	s.SettleStatus = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) SetChannelId(v string) *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList {
	s.ChannelId = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) SetMediaName(v string) *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList {
	s.MediaName = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) SetMediaSettleDetailId(v string) *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList {
	s.MediaSettleDetailId = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) SetModifiedDate(v string) *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList {
	s.ModifiedDate = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) SetCreateDate(v string) *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList {
	s.CreateDate = &v
	return s
}

func (s *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList) SetExtInfo(v string) *QueryMediaSettleInfoResponseBodyModelMediaSettleInfoList {
	s.ExtInfo = &v
	return s
}

type QueryMediaSettleInfoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMediaSettleInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMediaSettleInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMediaSettleInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryMediaSettleInfoResponse) SetHeaders(v map[string]*string) *QueryMediaSettleInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryMediaSettleInfoResponse) SetBody(v *QueryMediaSettleInfoResponseBody) *QueryMediaSettleInfoResponse {
	s.Body = v
	return s
}

type QueryMessagesRequest struct {
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Topic   *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	ExtJson *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s QueryMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMessagesRequest) GoString() string {
	return s.String()
}

func (s *QueryMessagesRequest) SetBizId(v string) *QueryMessagesRequest {
	s.BizId = &v
	return s
}

func (s *QueryMessagesRequest) SetTopic(v string) *QueryMessagesRequest {
	s.Topic = &v
	return s
}

func (s *QueryMessagesRequest) SetExtJson(v string) *QueryMessagesRequest {
	s.ExtJson = &v
	return s
}

type QueryMessagesResponseBody struct {
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage  *string                               `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code        *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	SubCode     *string                               `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	Message     *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	LogsId      *string                               `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	BizMessages *QueryMessagesResponseBodyBizMessages `json:"BizMessages,omitempty" xml:"BizMessages,omitempty" type:"Struct"`
}

func (s QueryMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMessagesResponseBody) SetRequestId(v string) *QueryMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMessagesResponseBody) SetSuccess(v bool) *QueryMessagesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMessagesResponseBody) SetSubMessage(v string) *QueryMessagesResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryMessagesResponseBody) SetCode(v string) *QueryMessagesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMessagesResponseBody) SetSubCode(v string) *QueryMessagesResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryMessagesResponseBody) SetMessage(v string) *QueryMessagesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMessagesResponseBody) SetLogsId(v string) *QueryMessagesResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryMessagesResponseBody) SetBizMessages(v *QueryMessagesResponseBodyBizMessages) *QueryMessagesResponseBody {
	s.BizMessages = v
	return s
}

type QueryMessagesResponseBodyBizMessages struct {
	BizMessage []*QueryMessagesResponseBodyBizMessagesBizMessage `json:"BizMessage,omitempty" xml:"BizMessage,omitempty" type:"Repeated"`
}

func (s QueryMessagesResponseBodyBizMessages) String() string {
	return tea.Prettify(s)
}

func (s QueryMessagesResponseBodyBizMessages) GoString() string {
	return s.String()
}

func (s *QueryMessagesResponseBodyBizMessages) SetBizMessage(v []*QueryMessagesResponseBodyBizMessagesBizMessage) *QueryMessagesResponseBodyBizMessages {
	s.BizMessage = v
	return s
}

type QueryMessagesResponseBodyBizMessagesBizMessage struct {
	ContentMapJson *string `json:"ContentMapJson,omitempty" xml:"ContentMapJson,omitempty"`
	PubTime        *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	DataId         *int64  `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Topic          *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMessagesResponseBodyBizMessagesBizMessage) String() string {
	return tea.Prettify(s)
}

func (s QueryMessagesResponseBodyBizMessagesBizMessage) GoString() string {
	return s.String()
}

func (s *QueryMessagesResponseBodyBizMessagesBizMessage) SetContentMapJson(v string) *QueryMessagesResponseBodyBizMessagesBizMessage {
	s.ContentMapJson = &v
	return s
}

func (s *QueryMessagesResponseBodyBizMessagesBizMessage) SetPubTime(v string) *QueryMessagesResponseBodyBizMessagesBizMessage {
	s.PubTime = &v
	return s
}

func (s *QueryMessagesResponseBodyBizMessagesBizMessage) SetDataId(v int64) *QueryMessagesResponseBodyBizMessagesBizMessage {
	s.DataId = &v
	return s
}

func (s *QueryMessagesResponseBodyBizMessagesBizMessage) SetTopic(v string) *QueryMessagesResponseBodyBizMessagesBizMessage {
	s.Topic = &v
	return s
}

type QueryMessagesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMessagesResponse) GoString() string {
	return s.String()
}

func (s *QueryMessagesResponse) SetHeaders(v map[string]*string) *QueryMessagesResponse {
	s.Headers = v
	return s
}

func (s *QueryMessagesResponse) SetBody(v *QueryMessagesResponseBody) *QueryMessagesResponse {
	s.Body = v
	return s
}

type QueryMovieCommentsRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	MovieId    *int64  `json:"MovieId,omitempty" xml:"MovieId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ExtJson    *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s QueryMovieCommentsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieCommentsRequest) GoString() string {
	return s.String()
}

func (s *QueryMovieCommentsRequest) SetBizId(v string) *QueryMovieCommentsRequest {
	s.BizId = &v
	return s
}

func (s *QueryMovieCommentsRequest) SetMovieId(v int64) *QueryMovieCommentsRequest {
	s.MovieId = &v
	return s
}

func (s *QueryMovieCommentsRequest) SetPageNumber(v int64) *QueryMovieCommentsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryMovieCommentsRequest) SetExtJson(v string) *QueryMovieCommentsRequest {
	s.ExtJson = &v
	return s
}

type QueryMovieCommentsResponseBody struct {
	RequestId   *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage  *string                                    `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code        *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode     *string                                    `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	LogsId      *string                                    `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Count       *int64                                     `json:"Count,omitempty" xml:"Count,omitempty"`
	CommentList *QueryMovieCommentsResponseBodyCommentList `json:"CommentList,omitempty" xml:"CommentList,omitempty" type:"Struct"`
}

func (s QueryMovieCommentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieCommentsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMovieCommentsResponseBody) SetRequestId(v string) *QueryMovieCommentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMovieCommentsResponseBody) SetSuccess(v bool) *QueryMovieCommentsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMovieCommentsResponseBody) SetSubMessage(v string) *QueryMovieCommentsResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryMovieCommentsResponseBody) SetCode(v string) *QueryMovieCommentsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMovieCommentsResponseBody) SetMessage(v string) *QueryMovieCommentsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMovieCommentsResponseBody) SetSubCode(v string) *QueryMovieCommentsResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryMovieCommentsResponseBody) SetLogsId(v string) *QueryMovieCommentsResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryMovieCommentsResponseBody) SetCount(v int64) *QueryMovieCommentsResponseBody {
	s.Count = &v
	return s
}

func (s *QueryMovieCommentsResponseBody) SetCommentList(v *QueryMovieCommentsResponseBodyCommentList) *QueryMovieCommentsResponseBody {
	s.CommentList = v
	return s
}

type QueryMovieCommentsResponseBodyCommentList struct {
	Comment []*QueryMovieCommentsResponseBodyCommentListComment `json:"Comment,omitempty" xml:"Comment,omitempty" type:"Repeated"`
}

func (s QueryMovieCommentsResponseBodyCommentList) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieCommentsResponseBodyCommentList) GoString() string {
	return s.String()
}

func (s *QueryMovieCommentsResponseBodyCommentList) SetComment(v []*QueryMovieCommentsResponseBodyCommentListComment) *QueryMovieCommentsResponseBodyCommentList {
	s.Comment = v
	return s
}

type QueryMovieCommentsResponseBodyCommentListComment struct {
	Remark      *int64  `json:"Remark,omitempty" xml:"Remark,omitempty"`
	FavorCount  *int64  `json:"FavorCount,omitempty" xml:"FavorCount,omitempty"`
	Subject     *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	MovieId     *int64  `json:"MovieId,omitempty" xml:"MovieId,omitempty"`
	NickName    *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	CommentTime *string `json:"CommentTime,omitempty" xml:"CommentTime,omitempty"`
}

func (s QueryMovieCommentsResponseBodyCommentListComment) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieCommentsResponseBodyCommentListComment) GoString() string {
	return s.String()
}

func (s *QueryMovieCommentsResponseBodyCommentListComment) SetRemark(v int64) *QueryMovieCommentsResponseBodyCommentListComment {
	s.Remark = &v
	return s
}

func (s *QueryMovieCommentsResponseBodyCommentListComment) SetFavorCount(v int64) *QueryMovieCommentsResponseBodyCommentListComment {
	s.FavorCount = &v
	return s
}

func (s *QueryMovieCommentsResponseBodyCommentListComment) SetSubject(v string) *QueryMovieCommentsResponseBodyCommentListComment {
	s.Subject = &v
	return s
}

func (s *QueryMovieCommentsResponseBodyCommentListComment) SetMovieId(v int64) *QueryMovieCommentsResponseBodyCommentListComment {
	s.MovieId = &v
	return s
}

func (s *QueryMovieCommentsResponseBodyCommentListComment) SetNickName(v string) *QueryMovieCommentsResponseBodyCommentListComment {
	s.NickName = &v
	return s
}

func (s *QueryMovieCommentsResponseBodyCommentListComment) SetContent(v string) *QueryMovieCommentsResponseBodyCommentListComment {
	s.Content = &v
	return s
}

func (s *QueryMovieCommentsResponseBodyCommentListComment) SetId(v int64) *QueryMovieCommentsResponseBodyCommentListComment {
	s.Id = &v
	return s
}

func (s *QueryMovieCommentsResponseBodyCommentListComment) SetCommentTime(v string) *QueryMovieCommentsResponseBodyCommentListComment {
	s.CommentTime = &v
	return s
}

type QueryMovieCommentsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMovieCommentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMovieCommentsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieCommentsResponse) GoString() string {
	return s.String()
}

func (s *QueryMovieCommentsResponse) SetHeaders(v map[string]*string) *QueryMovieCommentsResponse {
	s.Headers = v
	return s
}

func (s *QueryMovieCommentsResponse) SetBody(v *QueryMovieCommentsResponseBody) *QueryMovieCommentsResponse {
	s.Body = v
	return s
}

type QueryMovieSchedulesRequest struct {
	BizId    *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CinemaId *int64  `json:"CinemaId,omitempty" xml:"CinemaId,omitempty"`
	ExtJson  *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s QueryMovieSchedulesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieSchedulesRequest) GoString() string {
	return s.String()
}

func (s *QueryMovieSchedulesRequest) SetBizId(v string) *QueryMovieSchedulesRequest {
	s.BizId = &v
	return s
}

func (s *QueryMovieSchedulesRequest) SetCinemaId(v int64) *QueryMovieSchedulesRequest {
	s.CinemaId = &v
	return s
}

func (s *QueryMovieSchedulesRequest) SetExtJson(v string) *QueryMovieSchedulesRequest {
	s.ExtJson = &v
	return s
}

type QueryMovieSchedulesResponseBody struct {
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                                   `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                                   `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	LogsId     *string                                   `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Schedules  *QueryMovieSchedulesResponseBodySchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Struct"`
}

func (s QueryMovieSchedulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMovieSchedulesResponseBody) SetRequestId(v string) *QueryMovieSchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMovieSchedulesResponseBody) SetSuccess(v bool) *QueryMovieSchedulesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMovieSchedulesResponseBody) SetSubMessage(v string) *QueryMovieSchedulesResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryMovieSchedulesResponseBody) SetCode(v string) *QueryMovieSchedulesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMovieSchedulesResponseBody) SetMessage(v string) *QueryMovieSchedulesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMovieSchedulesResponseBody) SetSubCode(v string) *QueryMovieSchedulesResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryMovieSchedulesResponseBody) SetLogsId(v string) *QueryMovieSchedulesResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryMovieSchedulesResponseBody) SetSchedules(v *QueryMovieSchedulesResponseBodySchedules) *QueryMovieSchedulesResponseBody {
	s.Schedules = v
	return s
}

type QueryMovieSchedulesResponseBodySchedules struct {
	Schedule []*QueryMovieSchedulesResponseBodySchedulesSchedule `json:"Schedule,omitempty" xml:"Schedule,omitempty" type:"Repeated"`
}

func (s QueryMovieSchedulesResponseBodySchedules) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieSchedulesResponseBodySchedules) GoString() string {
	return s.String()
}

func (s *QueryMovieSchedulesResponseBodySchedules) SetSchedule(v []*QueryMovieSchedulesResponseBodySchedulesSchedule) *QueryMovieSchedulesResponseBodySchedules {
	s.Schedule = v
	return s
}

type QueryMovieSchedulesResponseBodySchedulesSchedule struct {
	MovieVersion        *string `json:"MovieVersion,omitempty" xml:"MovieVersion,omitempty"`
	SessionEndingTime   *string `json:"SessionEndingTime,omitempty" xml:"SessionEndingTime,omitempty"`
	MaxCanBuy           *int64  `json:"MaxCanBuy,omitempty" xml:"MaxCanBuy,omitempty"`
	MovieId             *int64  `json:"MovieId,omitempty" xml:"MovieId,omitempty"`
	ScheduleArea        *string `json:"ScheduleArea,omitempty" xml:"ScheduleArea,omitempty"`
	HallName            *string `json:"HallName,omitempty" xml:"HallName,omitempty"`
	IsExpired           *bool   `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	SessionStartingTime *string `json:"SessionStartingTime,omitempty" xml:"SessionStartingTime,omitempty"`
	Price               *int64  `json:"Price,omitempty" xml:"Price,omitempty"`
	SectionId           *string `json:"SectionId,omitempty" xml:"SectionId,omitempty"`
	ReleaseDate         *string `json:"ReleaseDate,omitempty" xml:"ReleaseDate,omitempty"`
	CinemaId            *int64  `json:"CinemaId,omitempty" xml:"CinemaId,omitempty"`
	ServiceFee          *int64  `json:"ServiceFee,omitempty" xml:"ServiceFee,omitempty"`
	Id                  *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QueryMovieSchedulesResponseBodySchedulesSchedule) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieSchedulesResponseBodySchedulesSchedule) GoString() string {
	return s.String()
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetMovieVersion(v string) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.MovieVersion = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetSessionEndingTime(v string) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.SessionEndingTime = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetMaxCanBuy(v int64) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.MaxCanBuy = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetMovieId(v int64) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.MovieId = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetScheduleArea(v string) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.ScheduleArea = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetHallName(v string) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.HallName = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetIsExpired(v bool) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.IsExpired = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetSessionStartingTime(v string) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.SessionStartingTime = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetPrice(v int64) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.Price = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetSectionId(v string) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.SectionId = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetReleaseDate(v string) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.ReleaseDate = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetCinemaId(v int64) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.CinemaId = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetServiceFee(v int64) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.ServiceFee = &v
	return s
}

func (s *QueryMovieSchedulesResponseBodySchedulesSchedule) SetId(v int64) *QueryMovieSchedulesResponseBodySchedulesSchedule {
	s.Id = &v
	return s
}

type QueryMovieSchedulesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMovieSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMovieSchedulesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieSchedulesResponse) GoString() string {
	return s.String()
}

func (s *QueryMovieSchedulesResponse) SetHeaders(v map[string]*string) *QueryMovieSchedulesResponse {
	s.Headers = v
	return s
}

func (s *QueryMovieSchedulesResponse) SetBody(v *QueryMovieSchedulesResponseBody) *QueryMovieSchedulesResponse {
	s.Body = v
	return s
}

type QueryMovieSeatsRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ScheduleId *int64  `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	ExtJson    *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s QueryMovieSeatsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieSeatsRequest) GoString() string {
	return s.String()
}

func (s *QueryMovieSeatsRequest) SetBizId(v string) *QueryMovieSeatsRequest {
	s.BizId = &v
	return s
}

func (s *QueryMovieSeatsRequest) SetScheduleId(v int64) *QueryMovieSeatsRequest {
	s.ScheduleId = &v
	return s
}

func (s *QueryMovieSeatsRequest) SetExtJson(v string) *QueryMovieSeatsRequest {
	s.ExtJson = &v
	return s
}

type QueryMovieSeatsResponseBody struct {
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                             `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                             `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	LogsId     *string                             `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	SeatMap    *QueryMovieSeatsResponseBodySeatMap `json:"SeatMap,omitempty" xml:"SeatMap,omitempty" type:"Struct"`
}

func (s QueryMovieSeatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMovieSeatsResponseBody) SetRequestId(v string) *QueryMovieSeatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMovieSeatsResponseBody) SetSuccess(v bool) *QueryMovieSeatsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMovieSeatsResponseBody) SetSubMessage(v string) *QueryMovieSeatsResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryMovieSeatsResponseBody) SetCode(v string) *QueryMovieSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMovieSeatsResponseBody) SetMessage(v string) *QueryMovieSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMovieSeatsResponseBody) SetSubCode(v string) *QueryMovieSeatsResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryMovieSeatsResponseBody) SetLogsId(v string) *QueryMovieSeatsResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryMovieSeatsResponseBody) SetSeatMap(v *QueryMovieSeatsResponseBodySeatMap) *QueryMovieSeatsResponseBody {
	s.SeatMap = v
	return s
}

type QueryMovieSeatsResponseBodySeatMap struct {
	MaxCanBuy  *int64                                   `json:"MaxCanBuy,omitempty" xml:"MaxCanBuy,omitempty"`
	TipMessage *string                                  `json:"TipMessage,omitempty" xml:"TipMessage,omitempty"`
	MaxRow     *int64                                   `json:"MaxRow,omitempty" xml:"MaxRow,omitempty"`
	MinColumn  *int64                                   `json:"MinColumn,omitempty" xml:"MinColumn,omitempty"`
	MinTopPx   *int64                                   `json:"MinTopPx,omitempty" xml:"MinTopPx,omitempty"`
	Notice     *string                                  `json:"Notice,omitempty" xml:"Notice,omitempty"`
	MaxColumn  *int64                                   `json:"MaxColumn,omitempty" xml:"MaxColumn,omitempty"`
	Regular    *bool                                    `json:"Regular,omitempty" xml:"Regular,omitempty"`
	MaxTopPx   *int64                                   `json:"MaxTopPx,omitempty" xml:"MaxTopPx,omitempty"`
	MaxLeftPx  *int64                                   `json:"MaxLeftPx,omitempty" xml:"MaxLeftPx,omitempty"`
	SoldCount  *int64                                   `json:"SoldCount,omitempty" xml:"SoldCount,omitempty"`
	MinRow     *int64                                   `json:"MinRow,omitempty" xml:"MinRow,omitempty"`
	SeatCount  *int64                                   `json:"SeatCount,omitempty" xml:"SeatCount,omitempty"`
	MinLeftPx  *int64                                   `json:"MinLeftPx,omitempty" xml:"MinLeftPx,omitempty"`
	Seats      *QueryMovieSeatsResponseBodySeatMapSeats `json:"Seats,omitempty" xml:"Seats,omitempty" type:"Struct"`
}

func (s QueryMovieSeatsResponseBodySeatMap) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieSeatsResponseBodySeatMap) GoString() string {
	return s.String()
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetMaxCanBuy(v int64) *QueryMovieSeatsResponseBodySeatMap {
	s.MaxCanBuy = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetTipMessage(v string) *QueryMovieSeatsResponseBodySeatMap {
	s.TipMessage = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetMaxRow(v int64) *QueryMovieSeatsResponseBodySeatMap {
	s.MaxRow = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetMinColumn(v int64) *QueryMovieSeatsResponseBodySeatMap {
	s.MinColumn = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetMinTopPx(v int64) *QueryMovieSeatsResponseBodySeatMap {
	s.MinTopPx = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetNotice(v string) *QueryMovieSeatsResponseBodySeatMap {
	s.Notice = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetMaxColumn(v int64) *QueryMovieSeatsResponseBodySeatMap {
	s.MaxColumn = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetRegular(v bool) *QueryMovieSeatsResponseBodySeatMap {
	s.Regular = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetMaxTopPx(v int64) *QueryMovieSeatsResponseBodySeatMap {
	s.MaxTopPx = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetMaxLeftPx(v int64) *QueryMovieSeatsResponseBodySeatMap {
	s.MaxLeftPx = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetSoldCount(v int64) *QueryMovieSeatsResponseBodySeatMap {
	s.SoldCount = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetMinRow(v int64) *QueryMovieSeatsResponseBodySeatMap {
	s.MinRow = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetSeatCount(v int64) *QueryMovieSeatsResponseBodySeatMap {
	s.SeatCount = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetMinLeftPx(v int64) *QueryMovieSeatsResponseBodySeatMap {
	s.MinLeftPx = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMap) SetSeats(v *QueryMovieSeatsResponseBodySeatMapSeats) *QueryMovieSeatsResponseBodySeatMap {
	s.Seats = v
	return s
}

type QueryMovieSeatsResponseBodySeatMapSeats struct {
	Seat []*QueryMovieSeatsResponseBodySeatMapSeatsSeat `json:"Seat,omitempty" xml:"Seat,omitempty" type:"Repeated"`
}

func (s QueryMovieSeatsResponseBodySeatMapSeats) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieSeatsResponseBodySeatMapSeats) GoString() string {
	return s.String()
}

func (s *QueryMovieSeatsResponseBodySeatMapSeats) SetSeat(v []*QueryMovieSeatsResponseBodySeatMapSeatsSeat) *QueryMovieSeatsResponseBodySeatMapSeats {
	s.Seat = v
	return s
}

type QueryMovieSeatsResponseBodySeatMapSeatsSeat struct {
	Status  *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	Flag    *int64  `json:"Flag,omitempty" xml:"Flag,omitempty"`
	ExtId   *string `json:"ExtId,omitempty" xml:"ExtId,omitempty"`
	RowName *string `json:"RowName,omitempty" xml:"RowName,omitempty"`
	TopPx   *int64  `json:"TopPx,omitempty" xml:"TopPx,omitempty"`
	Area    *string `json:"Area,omitempty" xml:"Area,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	LeftPx  *int64  `json:"LeftPx,omitempty" xml:"LeftPx,omitempty"`
	Column  *int64  `json:"Column,omitempty" xml:"Column,omitempty"`
	Row     *int64  `json:"Row,omitempty" xml:"Row,omitempty"`
}

func (s QueryMovieSeatsResponseBodySeatMapSeatsSeat) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieSeatsResponseBodySeatMapSeatsSeat) GoString() string {
	return s.String()
}

func (s *QueryMovieSeatsResponseBodySeatMapSeatsSeat) SetStatus(v int64) *QueryMovieSeatsResponseBodySeatMapSeatsSeat {
	s.Status = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMapSeatsSeat) SetFlag(v int64) *QueryMovieSeatsResponseBodySeatMapSeatsSeat {
	s.Flag = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMapSeatsSeat) SetExtId(v string) *QueryMovieSeatsResponseBodySeatMapSeatsSeat {
	s.ExtId = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMapSeatsSeat) SetRowName(v string) *QueryMovieSeatsResponseBodySeatMapSeatsSeat {
	s.RowName = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMapSeatsSeat) SetTopPx(v int64) *QueryMovieSeatsResponseBodySeatMapSeatsSeat {
	s.TopPx = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMapSeatsSeat) SetArea(v string) *QueryMovieSeatsResponseBodySeatMapSeatsSeat {
	s.Area = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMapSeatsSeat) SetName(v string) *QueryMovieSeatsResponseBodySeatMapSeatsSeat {
	s.Name = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMapSeatsSeat) SetLeftPx(v int64) *QueryMovieSeatsResponseBodySeatMapSeatsSeat {
	s.LeftPx = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMapSeatsSeat) SetColumn(v int64) *QueryMovieSeatsResponseBodySeatMapSeatsSeat {
	s.Column = &v
	return s
}

func (s *QueryMovieSeatsResponseBodySeatMapSeatsSeat) SetRow(v int64) *QueryMovieSeatsResponseBodySeatMapSeatsSeat {
	s.Row = &v
	return s
}

type QueryMovieSeatsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMovieSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMovieSeatsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieSeatsResponse) GoString() string {
	return s.String()
}

func (s *QueryMovieSeatsResponse) SetHeaders(v map[string]*string) *QueryMovieSeatsResponse {
	s.Headers = v
	return s
}

func (s *QueryMovieSeatsResponse) SetBody(v *QueryMovieSeatsResponseBody) *QueryMovieSeatsResponse {
	s.Body = v
	return s
}

type QueryMovieTicketsRequest struct {
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid  *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ExtJson *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s QueryMovieTicketsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieTicketsRequest) GoString() string {
	return s.String()
}

func (s *QueryMovieTicketsRequest) SetBizId(v string) *QueryMovieTicketsRequest {
	s.BizId = &v
	return s
}

func (s *QueryMovieTicketsRequest) SetBizUid(v string) *QueryMovieTicketsRequest {
	s.BizUid = &v
	return s
}

func (s *QueryMovieTicketsRequest) SetOrderId(v string) *QueryMovieTicketsRequest {
	s.OrderId = &v
	return s
}

func (s *QueryMovieTicketsRequest) SetExtJson(v string) *QueryMovieTicketsRequest {
	s.ExtJson = &v
	return s
}

type QueryMovieTicketsResponseBody struct {
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage  *string                                   `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code        *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	SubCode     *string                                   `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	Message     *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	LogsId      *string                                   `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	MovieTicket *QueryMovieTicketsResponseBodyMovieTicket `json:"MovieTicket,omitempty" xml:"MovieTicket,omitempty" type:"Struct"`
}

func (s QueryMovieTicketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMovieTicketsResponseBody) SetRequestId(v string) *QueryMovieTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMovieTicketsResponseBody) SetSuccess(v bool) *QueryMovieTicketsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMovieTicketsResponseBody) SetSubMessage(v string) *QueryMovieTicketsResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryMovieTicketsResponseBody) SetCode(v string) *QueryMovieTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMovieTicketsResponseBody) SetSubCode(v string) *QueryMovieTicketsResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryMovieTicketsResponseBody) SetMessage(v string) *QueryMovieTicketsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMovieTicketsResponseBody) SetLogsId(v string) *QueryMovieTicketsResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryMovieTicketsResponseBody) SetMovieTicket(v *QueryMovieTicketsResponseBodyMovieTicket) *QueryMovieTicketsResponseBody {
	s.MovieTicket = v
	return s
}

type QueryMovieTicketsResponseBodyMovieTicket struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TbOrderId      *string `json:"TbOrderId,omitempty" xml:"TbOrderId,omitempty"`
	TicketContents *string `json:"TicketContents,omitempty" xml:"TicketContents,omitempty"`
	ReturnMessage  *string `json:"ReturnMessage,omitempty" xml:"ReturnMessage,omitempty"`
}

func (s QueryMovieTicketsResponseBodyMovieTicket) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieTicketsResponseBodyMovieTicket) GoString() string {
	return s.String()
}

func (s *QueryMovieTicketsResponseBodyMovieTicket) SetStatus(v string) *QueryMovieTicketsResponseBodyMovieTicket {
	s.Status = &v
	return s
}

func (s *QueryMovieTicketsResponseBodyMovieTicket) SetTbOrderId(v string) *QueryMovieTicketsResponseBodyMovieTicket {
	s.TbOrderId = &v
	return s
}

func (s *QueryMovieTicketsResponseBodyMovieTicket) SetTicketContents(v string) *QueryMovieTicketsResponseBodyMovieTicket {
	s.TicketContents = &v
	return s
}

func (s *QueryMovieTicketsResponseBodyMovieTicket) SetReturnMessage(v string) *QueryMovieTicketsResponseBodyMovieTicket {
	s.ReturnMessage = &v
	return s
}

type QueryMovieTicketsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMovieTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMovieTicketsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMovieTicketsResponse) GoString() string {
	return s.String()
}

func (s *QueryMovieTicketsResponse) SetHeaders(v map[string]*string) *QueryMovieTicketsResponse {
	s.Headers = v
	return s
}

func (s *QueryMovieTicketsResponse) SetBody(v *QueryMovieTicketsResponseBody) *QueryMovieTicketsResponse {
	s.Body = v
	return s
}

type QueryOrderAndPaymentListRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	PageSize              *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber            *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	FilterOption          *string `json:"FilterOption,omitempty" xml:"FilterOption,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s QueryOrderAndPaymentListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListRequest) SetBizId(v string) *QueryOrderAndPaymentListRequest {
	s.BizId = &v
	return s
}

func (s *QueryOrderAndPaymentListRequest) SetBizUid(v string) *QueryOrderAndPaymentListRequest {
	s.BizUid = &v
	return s
}

func (s *QueryOrderAndPaymentListRequest) SetPageSize(v int64) *QueryOrderAndPaymentListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOrderAndPaymentListRequest) SetPageNumber(v int64) *QueryOrderAndPaymentListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryOrderAndPaymentListRequest) SetFilterOption(v string) *QueryOrderAndPaymentListRequest {
	s.FilterOption = &v
	return s
}

func (s *QueryOrderAndPaymentListRequest) SetUseAnonymousTbAccount(v bool) *QueryOrderAndPaymentListRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *QueryOrderAndPaymentListRequest) SetThirdPartyUserId(v string) *QueryOrderAndPaymentListRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *QueryOrderAndPaymentListRequest) SetAccountType(v string) *QueryOrderAndPaymentListRequest {
	s.AccountType = &v
	return s
}

type QueryOrderAndPaymentListResponseBody struct {
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code        *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize    *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount  *int64                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PostFee     *QueryOrderAndPaymentListResponseBodyPostFee     `json:"PostFee,omitempty" xml:"PostFee,omitempty" type:"Struct"`
	LmOrderList *QueryOrderAndPaymentListResponseBodyLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Struct"`
}

func (s QueryOrderAndPaymentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListResponseBody) SetRequestId(v string) *QueryOrderAndPaymentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBody) SetCode(v string) *QueryOrderAndPaymentListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBody) SetMessage(v string) *QueryOrderAndPaymentListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBody) SetPageSize(v int32) *QueryOrderAndPaymentListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBody) SetPageNumber(v int32) *QueryOrderAndPaymentListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBody) SetTotalCount(v int64) *QueryOrderAndPaymentListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBody) SetPostFee(v *QueryOrderAndPaymentListResponseBodyPostFee) *QueryOrderAndPaymentListResponseBody {
	s.PostFee = v
	return s
}

func (s *QueryOrderAndPaymentListResponseBody) SetLmOrderList(v *QueryOrderAndPaymentListResponseBodyLmOrderList) *QueryOrderAndPaymentListResponseBody {
	s.LmOrderList = v
	return s
}

type QueryOrderAndPaymentListResponseBodyPostFee struct {
	FundAmount      *int64  `json:"FundAmount,omitempty" xml:"FundAmount,omitempty"`
	FundType        *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	FundAmountMoney *int64  `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
}

func (s QueryOrderAndPaymentListResponseBodyPostFee) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListResponseBodyPostFee) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListResponseBodyPostFee) SetFundAmount(v int64) *QueryOrderAndPaymentListResponseBodyPostFee {
	s.FundAmount = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyPostFee) SetFundType(v string) *QueryOrderAndPaymentListResponseBodyPostFee {
	s.FundType = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyPostFee) SetFundAmountMoney(v int64) *QueryOrderAndPaymentListResponseBodyPostFee {
	s.FundAmountMoney = &v
	return s
}

type QueryOrderAndPaymentListResponseBodyLmOrderList struct {
	LmOrderList []*QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Repeated"`
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderList) SetLmOrderList(v []*QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) *QueryOrderAndPaymentListResponseBodyLmOrderList {
	s.LmOrderList = v
	return s
}

type QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList struct {
	LmOrderId           *int64                                                                         `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	EnableStatus        *int32                                                                         `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	ExtJson             *string                                                                        `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	OrderStatus         *int32                                                                         `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	LmPaymentId         *string                                                                        `json:"LmPaymentId,omitempty" xml:"LmPaymentId,omitempty"`
	ShopName            *string                                                                        `json:"ShopName,omitempty" xml:"ShopName,omitempty"`
	TbOrderId           *int64                                                                         `json:"TbOrderId,omitempty" xml:"TbOrderId,omitempty"`
	OrderAmount         *int64                                                                         `json:"OrderAmount,omitempty" xml:"OrderAmount,omitempty"`
	LogisticsStatus     *int32                                                                         `json:"LogisticsStatus,omitempty" xml:"LogisticsStatus,omitempty"`
	CreateDate          *string                                                                        `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	FundStructureModels *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModels `json:"FundStructureModels,omitempty" xml:"FundStructureModels,omitempty" type:"Struct"`
	SubOrderList        *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderList        `json:"SubOrderList,omitempty" xml:"SubOrderList,omitempty" type:"Struct"`
	PostFee             *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListPostFee             `json:"PostFee,omitempty" xml:"PostFee,omitempty" type:"Struct"`
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetLmOrderId(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetEnableStatus(v int32) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.EnableStatus = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetExtJson(v string) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.ExtJson = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetOrderStatus(v int32) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.OrderStatus = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetLmPaymentId(v string) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.LmPaymentId = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetShopName(v string) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.ShopName = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetTbOrderId(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.TbOrderId = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetOrderAmount(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.OrderAmount = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetLogisticsStatus(v int32) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.LogisticsStatus = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetCreateDate(v string) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.CreateDate = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetFundStructureModels(v *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModels) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.FundStructureModels = v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetSubOrderList(v *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderList) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.SubOrderList = v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList) SetPostFee(v *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListPostFee) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderList {
	s.PostFee = v
	return s
}

type QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModels struct {
	FundStructureModels []*QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels `json:"FundStructureModels,omitempty" xml:"FundStructureModels,omitempty" type:"Repeated"`
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModels) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModels) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModels) SetFundStructureModels(v []*QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModels {
	s.FundStructureModels = v
	return s
}

type QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels struct {
	FundAmount      *int64  `json:"FundAmount,omitempty" xml:"FundAmount,omitempty"`
	FundType        *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	FundAmountMoney *int64  `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels) SetFundAmount(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels {
	s.FundAmount = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels) SetFundType(v string) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels {
	s.FundType = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels) SetFundAmountMoney(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels {
	s.FundAmountMoney = &v
	return s
}

type QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderList struct {
	SubOrderList []*QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList `json:"SubOrderList,omitempty" xml:"SubOrderList,omitempty" type:"Repeated"`
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderList) SetSubOrderList(v []*QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderList {
	s.SubOrderList = v
	return s
}

type QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList struct {
	LmOrderId     *int64                                                                                           `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	EnableStatus  *int32                                                                                           `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	ItemTitle     *string                                                                                          `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	OrderStatus   *int32                                                                                           `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	SkuName       *string                                                                                          `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
	LmItemId      *string                                                                                          `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	SkuId         *int64                                                                                           `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Number        *int64                                                                                           `json:"Number,omitempty" xml:"Number,omitempty"`
	TbOrderId     *int64                                                                                           `json:"TbOrderId,omitempty" xml:"TbOrderId,omitempty"`
	ItemPic       *string                                                                                          `json:"ItemPic,omitempty" xml:"ItemPic,omitempty"`
	ItemId        *int64                                                                                           `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemPriceList *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList `json:"ItemPriceList,omitempty" xml:"ItemPriceList,omitempty" type:"Struct"`
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetLmOrderId(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetEnableStatus(v int32) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.EnableStatus = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetItemTitle(v string) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.ItemTitle = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetOrderStatus(v int32) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.OrderStatus = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetSkuName(v string) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.SkuName = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetLmItemId(v string) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.LmItemId = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetSkuId(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.SkuId = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetNumber(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.Number = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetTbOrderId(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.TbOrderId = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetItemPic(v string) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.ItemPic = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetItemId(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.ItemId = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetItemPriceList(v *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.ItemPriceList = v
	return s
}

type QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList struct {
	ItemPriceList []*QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList `json:"ItemPriceList,omitempty" xml:"ItemPriceList,omitempty" type:"Repeated"`
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList) SetItemPriceList(v []*QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList {
	s.ItemPriceList = v
	return s
}

type QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList struct {
	FundAmount      *int64  `json:"FundAmount,omitempty" xml:"FundAmount,omitempty"`
	FundType        *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	FundAmountMoney *int64  `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList) SetFundAmount(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList {
	s.FundAmount = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList) SetFundType(v string) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList {
	s.FundType = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList) SetFundAmountMoney(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList {
	s.FundAmountMoney = &v
	return s
}

type QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListPostFee struct {
	FundAmount      *int64  `json:"FundAmount,omitempty" xml:"FundAmount,omitempty"`
	FundType        *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	FundAmountMoney *int64  `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListPostFee) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListPostFee) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListPostFee) SetFundAmount(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListPostFee {
	s.FundAmount = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListPostFee) SetFundType(v string) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListPostFee {
	s.FundType = &v
	return s
}

func (s *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListPostFee) SetFundAmountMoney(v int64) *QueryOrderAndPaymentListResponseBodyLmOrderListLmOrderListPostFee {
	s.FundAmountMoney = &v
	return s
}

type QueryOrderAndPaymentListResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrderAndPaymentListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrderAndPaymentListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderAndPaymentListResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderAndPaymentListResponse) SetHeaders(v map[string]*string) *QueryOrderAndPaymentListResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderAndPaymentListResponse) SetBody(v *QueryOrderAndPaymentListResponseBody) *QueryOrderAndPaymentListResponse {
	s.Body = v
	return s
}

type QueryOrderCommissionRateRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	LmOrderId             *int64  `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s QueryOrderCommissionRateRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderCommissionRateRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderCommissionRateRequest) SetBizId(v string) *QueryOrderCommissionRateRequest {
	s.BizId = &v
	return s
}

func (s *QueryOrderCommissionRateRequest) SetBizUid(v string) *QueryOrderCommissionRateRequest {
	s.BizUid = &v
	return s
}

func (s *QueryOrderCommissionRateRequest) SetLmOrderId(v int64) *QueryOrderCommissionRateRequest {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderCommissionRateRequest) SetUseAnonymousTbAccount(v bool) *QueryOrderCommissionRateRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *QueryOrderCommissionRateRequest) SetThirdPartyUserId(v string) *QueryOrderCommissionRateRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *QueryOrderCommissionRateRequest) SetAccountType(v string) *QueryOrderCommissionRateRequest {
	s.AccountType = &v
	return s
}

type QueryOrderCommissionRateResponseBody struct {
	Code             *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CommissionModels []*QueryOrderCommissionRateResponseBodyCommissionModels `json:"CommissionModels,omitempty" xml:"CommissionModels,omitempty" type:"Repeated"`
}

func (s QueryOrderCommissionRateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderCommissionRateResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderCommissionRateResponseBody) SetCode(v string) *QueryOrderCommissionRateResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrderCommissionRateResponseBody) SetMessage(v string) *QueryOrderCommissionRateResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrderCommissionRateResponseBody) SetRequestId(v string) *QueryOrderCommissionRateResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderCommissionRateResponseBody) SetCommissionModels(v []*QueryOrderCommissionRateResponseBodyCommissionModels) *QueryOrderCommissionRateResponseBody {
	s.CommissionModels = v
	return s
}

type QueryOrderCommissionRateResponseBodyCommissionModels struct {
	LmOrderId      *int64                                                              `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	CommissionInfo *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfo `json:"CommissionInfo,omitempty" xml:"CommissionInfo,omitempty" type:"Struct"`
}

func (s QueryOrderCommissionRateResponseBodyCommissionModels) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderCommissionRateResponseBodyCommissionModels) GoString() string {
	return s.String()
}

func (s *QueryOrderCommissionRateResponseBodyCommissionModels) SetLmOrderId(v int64) *QueryOrderCommissionRateResponseBodyCommissionModels {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderCommissionRateResponseBodyCommissionModels) SetCommissionInfo(v *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfo) *QueryOrderCommissionRateResponseBodyCommissionModels {
	s.CommissionInfo = v
	return s
}

type QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfo struct {
	RateType   *string                                                                       `json:"RateType,omitempty" xml:"RateType,omitempty"`
	RateConfig *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfig `json:"RateConfig,omitempty" xml:"RateConfig,omitempty" type:"Struct"`
}

func (s QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfo) GoString() string {
	return s.String()
}

func (s *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfo) SetRateType(v string) *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfo {
	s.RateType = &v
	return s
}

func (s *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfo) SetRateConfig(v *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfig) *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfo {
	s.RateConfig = v
	return s
}

type QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfig struct {
	Configs []*QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfigConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
}

func (s QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfig) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfig) GoString() string {
	return s.String()
}

func (s *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfig) SetConfigs(v []*QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfigConfigs) *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfig {
	s.Configs = v
	return s
}

type QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfigConfigs struct {
	ValueUnit *string `json:"ValueUnit,omitempty" xml:"ValueUnit,omitempty"`
	Value     *int64  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfigConfigs) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfigConfigs) GoString() string {
	return s.String()
}

func (s *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfigConfigs) SetValueUnit(v string) *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfigConfigs {
	s.ValueUnit = &v
	return s
}

func (s *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfigConfigs) SetValue(v int64) *QueryOrderCommissionRateResponseBodyCommissionModelsCommissionInfoRateConfigConfigs {
	s.Value = &v
	return s
}

type QueryOrderCommissionRateResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrderCommissionRateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrderCommissionRateResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderCommissionRateResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderCommissionRateResponse) SetHeaders(v map[string]*string) *QueryOrderCommissionRateResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderCommissionRateResponse) SetBody(v *QueryOrderCommissionRateResponseBody) *QueryOrderCommissionRateResponse {
	s.Body = v
	return s
}

type QueryOrderDetailInnerRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	FilterOption          *string `json:"FilterOption,omitempty" xml:"FilterOption,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
}

func (s QueryOrderDetailInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetailInnerRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderDetailInnerRequest) SetBizId(v string) *QueryOrderDetailInnerRequest {
	s.BizId = &v
	return s
}

func (s *QueryOrderDetailInnerRequest) SetBizUid(v string) *QueryOrderDetailInnerRequest {
	s.BizUid = &v
	return s
}

func (s *QueryOrderDetailInnerRequest) SetFilterOption(v string) *QueryOrderDetailInnerRequest {
	s.FilterOption = &v
	return s
}

func (s *QueryOrderDetailInnerRequest) SetUseAnonymousTbAccount(v bool) *QueryOrderDetailInnerRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *QueryOrderDetailInnerRequest) SetThirdPartyUserId(v string) *QueryOrderDetailInnerRequest {
	s.ThirdPartyUserId = &v
	return s
}

type QueryOrderDetailInnerResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Model     *QueryOrderDetailInnerResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s QueryOrderDetailInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetailInnerResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderDetailInnerResponseBody) SetCode(v string) *QueryOrderDetailInnerResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBody) SetMessage(v string) *QueryOrderDetailInnerResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBody) SetRequestId(v string) *QueryOrderDetailInnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBody) SetModel(v *QueryOrderDetailInnerResponseBodyModel) *QueryOrderDetailInnerResponseBody {
	s.Model = v
	return s
}

type QueryOrderDetailInnerResponseBodyModel struct {
	Order *QueryOrderDetailInnerResponseBodyModelOrder `json:"order,omitempty" xml:"order,omitempty" type:"Struct"`
}

func (s QueryOrderDetailInnerResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetailInnerResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryOrderDetailInnerResponseBodyModel) SetOrder(v *QueryOrderDetailInnerResponseBodyModelOrder) *QueryOrderDetailInnerResponseBodyModel {
	s.Order = v
	return s
}

type QueryOrderDetailInnerResponseBodyModelOrder struct {
	Eticket              *bool                                                           `json:"Eticket,omitempty" xml:"Eticket,omitempty"`
	CreateDate           *string                                                         `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	LogisticsCompName    *string                                                         `json:"LogisticsCompName,omitempty" xml:"LogisticsCompName,omitempty"`
	ChannelCode          *string                                                         `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	LogisticsMobilePhone *string                                                         `json:"LogisticsMobilePhone,omitempty" xml:"LogisticsMobilePhone,omitempty"`
	ResExtInfo           *string                                                         `json:"ResExtInfo,omitempty" xml:"ResExtInfo,omitempty"`
	EnableStatus         *int32                                                          `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	ChannelBizType       *string                                                         `json:"ChannelBizType,omitempty" xml:"ChannelBizType,omitempty"`
	Shipping             *string                                                         `json:"Shipping,omitempty" xml:"Shipping,omitempty"`
	OrderPayInfo         *string                                                         `json:"OrderPayInfo,omitempty" xml:"OrderPayInfo,omitempty"`
	LogisticsStatusDesc  *string                                                         `json:"LogisticsStatusDesc,omitempty" xml:"LogisticsStatusDesc,omitempty"`
	TbOrderId            *string                                                         `json:"TbOrderId,omitempty" xml:"TbOrderId,omitempty"`
	LogisticsStatus      *int32                                                          `json:"LogisticsStatus,omitempty" xml:"LogisticsStatus,omitempty"`
	LmOrderId            *int64                                                          `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	SellerId             *int64                                                          `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	ShopName             *string                                                         `json:"ShopName,omitempty" xml:"ShopName,omitempty"`
	OrderAmount          *int64                                                          `json:"OrderAmount,omitempty" xml:"OrderAmount,omitempty"`
	ExtInfo              map[string]interface{}                                          `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	EndTime              *int64                                                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PayStatus            *int32                                                          `json:"PayStatus,omitempty" xml:"PayStatus,omitempty"`
	LogisticsNo          *string                                                         `json:"LogisticsNo,omitempty" xml:"LogisticsNo,omitempty"`
	LogisticsUserName    *string                                                         `json:"LogisticsUserName,omitempty" xml:"LogisticsUserName,omitempty"`
	LogisticsAddress     *string                                                         `json:"LogisticsAddress,omitempty" xml:"LogisticsAddress,omitempty"`
	PayWaterStatus       *int32                                                          `json:"PayWaterStatus,omitempty" xml:"PayWaterStatus,omitempty"`
	RefundStatus         *int32                                                          `json:"RefundStatus,omitempty" xml:"RefundStatus,omitempty"`
	SellerNick           *string                                                         `json:"SellerNick,omitempty" xml:"SellerNick,omitempty"`
	ChannelOrderId       *string                                                         `json:"ChannelOrderId,omitempty" xml:"ChannelOrderId,omitempty"`
	SubItemOrderList     *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderList    `json:"SubItemOrderList,omitempty" xml:"SubItemOrderList,omitempty" type:"Struct"`
	FundStructureModels  *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModels `json:"FundStructureModels,omitempty" xml:"FundStructureModels,omitempty" type:"Struct"`
	PostFee              *QueryOrderDetailInnerResponseBodyModelOrderPostFee             `json:"PostFee,omitempty" xml:"PostFee,omitempty" type:"Struct"`
}

func (s QueryOrderDetailInnerResponseBodyModelOrder) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetailInnerResponseBodyModelOrder) GoString() string {
	return s.String()
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetEticket(v bool) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.Eticket = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetCreateDate(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.CreateDate = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetLogisticsCompName(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.LogisticsCompName = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetChannelCode(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.ChannelCode = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetLogisticsMobilePhone(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.LogisticsMobilePhone = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetResExtInfo(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.ResExtInfo = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetEnableStatus(v int32) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.EnableStatus = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetChannelBizType(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.ChannelBizType = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetShipping(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.Shipping = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetOrderPayInfo(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.OrderPayInfo = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetLogisticsStatusDesc(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.LogisticsStatusDesc = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetTbOrderId(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.TbOrderId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetLogisticsStatus(v int32) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.LogisticsStatus = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetLmOrderId(v int64) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetSellerId(v int64) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.SellerId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetShopName(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.ShopName = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetOrderAmount(v int64) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.OrderAmount = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetExtInfo(v map[string]interface{}) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.ExtInfo = v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetEndTime(v int64) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.EndTime = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetPayStatus(v int32) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.PayStatus = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetLogisticsNo(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.LogisticsNo = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetLogisticsUserName(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.LogisticsUserName = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetLogisticsAddress(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.LogisticsAddress = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetPayWaterStatus(v int32) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.PayWaterStatus = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetRefundStatus(v int32) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.RefundStatus = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetSellerNick(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.SellerNick = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetChannelOrderId(v string) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.ChannelOrderId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetSubItemOrderList(v *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderList) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.SubItemOrderList = v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetFundStructureModels(v *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModels) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.FundStructureModels = v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrder) SetPostFee(v *QueryOrderDetailInnerResponseBodyModelOrderPostFee) *QueryOrderDetailInnerResponseBodyModelOrder {
	s.PostFee = v
	return s
}

type QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderList struct {
	SubItemOrder []*QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder `json:"SubItemOrder,omitempty" xml:"SubItemOrder,omitempty" type:"Repeated"`
}

func (s QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderList) SetSubItemOrder(v []*QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderList {
	s.SubItemOrder = v
	return s
}

type QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder struct {
	LmOrderId        *int64                                                                                `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	ItemPriceInfo    *string                                                                               `json:"ItemPriceInfo,omitempty" xml:"ItemPriceInfo,omitempty"`
	ItemTitle        *string                                                                               `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	TbSubOrderId     *int64                                                                                `json:"TbSubOrderId,omitempty" xml:"TbSubOrderId,omitempty"`
	LmItemId         *string                                                                               `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	TotalPaymentInfo *string                                                                               `json:"TotalPaymentInfo,omitempty" xml:"TotalPaymentInfo,omitempty"`
	ItemPic          *string                                                                               `json:"ItemPic,omitempty" xml:"ItemPic,omitempty"`
	ItemId           *int64                                                                                `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ChannelCode      *string                                                                               `json:"ChannelCode,omitempty" xml:"ChannelCode,omitempty"`
	PayStatus        *int32                                                                                `json:"PayStatus,omitempty" xml:"PayStatus,omitempty"`
	SkuName          *string                                                                               `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
	Number           *int32                                                                                `json:"Number,omitempty" xml:"Number,omitempty"`
	SkuId            *int64                                                                                `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	ItemPriceList    *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceList `json:"ItemPriceList,omitempty" xml:"ItemPriceList,omitempty" type:"Struct"`
}

func (s QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) GoString() string {
	return s.String()
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetLmOrderId(v int64) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetItemPriceInfo(v string) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.ItemPriceInfo = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetItemTitle(v string) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.ItemTitle = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetTbSubOrderId(v int64) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.TbSubOrderId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetLmItemId(v string) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.LmItemId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetTotalPaymentInfo(v string) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.TotalPaymentInfo = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetItemPic(v string) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.ItemPic = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetItemId(v int64) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.ItemId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetChannelCode(v string) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.ChannelCode = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetPayStatus(v int32) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.PayStatus = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetSkuName(v string) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.SkuName = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetNumber(v int32) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.Number = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetSkuId(v int64) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.SkuId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder) SetItemPriceList(v *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceList) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrder {
	s.ItemPriceList = v
	return s
}

type QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceList struct {
	ItemPrice []*QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice `json:"ItemPrice,omitempty" xml:"ItemPrice,omitempty" type:"Repeated"`
}

func (s QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceList) GoString() string {
	return s.String()
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceList) SetItemPrice(v []*QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceList {
	s.ItemPrice = v
	return s
}

type QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice struct {
	FundAmount      *int64  `json:"FundAmount,omitempty" xml:"FundAmount,omitempty"`
	TbOrderId       *int64  `json:"TbOrderId,omitempty" xml:"TbOrderId,omitempty"`
	FundType        *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	FundAmountMoney *int64  `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
	TbSubOrderId    *int64  `json:"TbSubOrderId,omitempty" xml:"TbSubOrderId,omitempty"`
}

func (s QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice) GoString() string {
	return s.String()
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice) SetFundAmount(v int64) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice {
	s.FundAmount = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice) SetTbOrderId(v int64) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice {
	s.TbOrderId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice) SetFundType(v string) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice {
	s.FundType = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice) SetFundAmountMoney(v int64) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice {
	s.FundAmountMoney = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice) SetTbSubOrderId(v int64) *QueryOrderDetailInnerResponseBodyModelOrderSubItemOrderListSubItemOrderItemPriceListItemPrice {
	s.TbSubOrderId = &v
	return s
}

type QueryOrderDetailInnerResponseBodyModelOrderFundStructureModels struct {
	FundStructure []*QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure `json:"FundStructure,omitempty" xml:"FundStructure,omitempty" type:"Repeated"`
}

func (s QueryOrderDetailInnerResponseBodyModelOrderFundStructureModels) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetailInnerResponseBodyModelOrderFundStructureModels) GoString() string {
	return s.String()
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModels) SetFundStructure(v []*QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure) *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModels {
	s.FundStructure = v
	return s
}

type QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure struct {
	FundAmount      *int64  `json:"FundAmount,omitempty" xml:"FundAmount,omitempty"`
	LmOrderId       *int64  `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	FundType        *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	FundAmountMoney *int64  `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
	TbSubOrderId    *int64  `json:"TbSubOrderId,omitempty" xml:"TbSubOrderId,omitempty"`
}

func (s QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure) GoString() string {
	return s.String()
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure) SetFundAmount(v int64) *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure {
	s.FundAmount = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure) SetLmOrderId(v int64) *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure) SetFundType(v string) *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure {
	s.FundType = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure) SetFundAmountMoney(v int64) *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure {
	s.FundAmountMoney = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure) SetTbSubOrderId(v int64) *QueryOrderDetailInnerResponseBodyModelOrderFundStructureModelsFundStructure {
	s.TbSubOrderId = &v
	return s
}

type QueryOrderDetailInnerResponseBodyModelOrderPostFee struct {
	FundAmount      *int64  `json:"FundAmount,omitempty" xml:"FundAmount,omitempty"`
	LmOrderId       *int64  `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	FundType        *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	FundAmountMoney *int64  `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
	TbSubOrderId    *int64  `json:"TbSubOrderId,omitempty" xml:"TbSubOrderId,omitempty"`
}

func (s QueryOrderDetailInnerResponseBodyModelOrderPostFee) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetailInnerResponseBodyModelOrderPostFee) GoString() string {
	return s.String()
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderPostFee) SetFundAmount(v int64) *QueryOrderDetailInnerResponseBodyModelOrderPostFee {
	s.FundAmount = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderPostFee) SetLmOrderId(v int64) *QueryOrderDetailInnerResponseBodyModelOrderPostFee {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderPostFee) SetFundType(v string) *QueryOrderDetailInnerResponseBodyModelOrderPostFee {
	s.FundType = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderPostFee) SetFundAmountMoney(v int64) *QueryOrderDetailInnerResponseBodyModelOrderPostFee {
	s.FundAmountMoney = &v
	return s
}

func (s *QueryOrderDetailInnerResponseBodyModelOrderPostFee) SetTbSubOrderId(v int64) *QueryOrderDetailInnerResponseBodyModelOrderPostFee {
	s.TbSubOrderId = &v
	return s
}

type QueryOrderDetailInnerResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrderDetailInnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrderDetailInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetailInnerResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderDetailInnerResponse) SetHeaders(v map[string]*string) *QueryOrderDetailInnerResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderDetailInnerResponse) SetBody(v *QueryOrderDetailInnerResponseBody) *QueryOrderDetailInnerResponse {
	s.Body = v
	return s
}

type QueryOrderIdByPayIdRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	PaymentId             *string `json:"PaymentId,omitempty" xml:"PaymentId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s QueryOrderIdByPayIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderIdByPayIdRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderIdByPayIdRequest) SetBizId(v string) *QueryOrderIdByPayIdRequest {
	s.BizId = &v
	return s
}

func (s *QueryOrderIdByPayIdRequest) SetBizUid(v string) *QueryOrderIdByPayIdRequest {
	s.BizUid = &v
	return s
}

func (s *QueryOrderIdByPayIdRequest) SetPaymentId(v string) *QueryOrderIdByPayIdRequest {
	s.PaymentId = &v
	return s
}

func (s *QueryOrderIdByPayIdRequest) SetUseAnonymousTbAccount(v bool) *QueryOrderIdByPayIdRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *QueryOrderIdByPayIdRequest) SetThirdPartyUserId(v string) *QueryOrderIdByPayIdRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *QueryOrderIdByPayIdRequest) SetAccountType(v string) *QueryOrderIdByPayIdRequest {
	s.AccountType = &v
	return s
}

type QueryOrderIdByPayIdResponseBody struct {
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LmOrderIds *QueryOrderIdByPayIdResponseBodyLmOrderIds `json:"LmOrderIds,omitempty" xml:"LmOrderIds,omitempty" type:"Struct"`
}

func (s QueryOrderIdByPayIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderIdByPayIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderIdByPayIdResponseBody) SetCode(v string) *QueryOrderIdByPayIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrderIdByPayIdResponseBody) SetMessage(v string) *QueryOrderIdByPayIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrderIdByPayIdResponseBody) SetRequestId(v string) *QueryOrderIdByPayIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderIdByPayIdResponseBody) SetLmOrderIds(v *QueryOrderIdByPayIdResponseBodyLmOrderIds) *QueryOrderIdByPayIdResponseBody {
	s.LmOrderIds = v
	return s
}

type QueryOrderIdByPayIdResponseBodyLmOrderIds struct {
	LmOrderIds []*QueryOrderIdByPayIdResponseBodyLmOrderIdsLmOrderIds `json:"LmOrderIds,omitempty" xml:"LmOrderIds,omitempty" type:"Repeated"`
}

func (s QueryOrderIdByPayIdResponseBodyLmOrderIds) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderIdByPayIdResponseBodyLmOrderIds) GoString() string {
	return s.String()
}

func (s *QueryOrderIdByPayIdResponseBodyLmOrderIds) SetLmOrderIds(v []*QueryOrderIdByPayIdResponseBodyLmOrderIdsLmOrderIds) *QueryOrderIdByPayIdResponseBodyLmOrderIds {
	s.LmOrderIds = v
	return s
}

type QueryOrderIdByPayIdResponseBodyLmOrderIdsLmOrderIds struct {
	LmOrderId *int64 `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
}

func (s QueryOrderIdByPayIdResponseBodyLmOrderIdsLmOrderIds) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderIdByPayIdResponseBodyLmOrderIdsLmOrderIds) GoString() string {
	return s.String()
}

func (s *QueryOrderIdByPayIdResponseBodyLmOrderIdsLmOrderIds) SetLmOrderId(v int64) *QueryOrderIdByPayIdResponseBodyLmOrderIdsLmOrderIds {
	s.LmOrderId = &v
	return s
}

type QueryOrderIdByPayIdResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrderIdByPayIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrderIdByPayIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderIdByPayIdResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderIdByPayIdResponse) SetHeaders(v map[string]*string) *QueryOrderIdByPayIdResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderIdByPayIdResponse) SetBody(v *QueryOrderIdByPayIdResponseBody) *QueryOrderIdByPayIdResponse {
	s.Body = v
	return s
}

type QueryOrderInfoAfterSaleRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ChannelUserId         *string `json:"ChannelUserId,omitempty" xml:"ChannelUserId,omitempty"`
	LmOrderId             *string `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s QueryOrderInfoAfterSaleRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderInfoAfterSaleRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderInfoAfterSaleRequest) SetBizId(v string) *QueryOrderInfoAfterSaleRequest {
	s.BizId = &v
	return s
}

func (s *QueryOrderInfoAfterSaleRequest) SetChannelUserId(v string) *QueryOrderInfoAfterSaleRequest {
	s.ChannelUserId = &v
	return s
}

func (s *QueryOrderInfoAfterSaleRequest) SetLmOrderId(v string) *QueryOrderInfoAfterSaleRequest {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderInfoAfterSaleRequest) SetUseAnonymousTbAccount(v bool) *QueryOrderInfoAfterSaleRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *QueryOrderInfoAfterSaleRequest) SetThirdPartyUserId(v string) *QueryOrderInfoAfterSaleRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *QueryOrderInfoAfterSaleRequest) SetAccountType(v string) *QueryOrderInfoAfterSaleRequest {
	s.AccountType = &v
	return s
}

type QueryOrderInfoAfterSaleResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Model     *QueryOrderInfoAfterSaleResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s QueryOrderInfoAfterSaleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderInfoAfterSaleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderInfoAfterSaleResponseBody) SetCode(v string) *QueryOrderInfoAfterSaleResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBody) SetMessage(v string) *QueryOrderInfoAfterSaleResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBody) SetRequestId(v string) *QueryOrderInfoAfterSaleResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBody) SetModel(v *QueryOrderInfoAfterSaleResponseBodyModel) *QueryOrderInfoAfterSaleResponseBody {
	s.Model = v
	return s
}

type QueryOrderInfoAfterSaleResponseBodyModel struct {
	LmOrderId            *int64                                                 `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	ShopServiceTelephone *string                                                `json:"ShopServiceTelephone,omitempty" xml:"ShopServiceTelephone,omitempty"`
	RefundAmount         *int64                                                 `json:"RefundAmount,omitempty" xml:"RefundAmount,omitempty"`
	XiaomiCode           *string                                                `json:"XiaomiCode,omitempty" xml:"XiaomiCode,omitempty"`
	ShopName             *string                                                `json:"ShopName,omitempty" xml:"ShopName,omitempty"`
	CreateDate           *string                                                `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	RefundRate           *string                                                `json:"RefundRate,omitempty" xml:"RefundRate,omitempty"`
	ExtJson              *string                                                `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	OrderStatus          *string                                                `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	RefundPoints         *int64                                                 `json:"RefundPoints,omitempty" xml:"RefundPoints,omitempty"`
	TbOrderId            *int64                                                 `json:"TbOrderId,omitempty" xml:"TbOrderId,omitempty"`
	Points               *int64                                                 `json:"Points,omitempty" xml:"Points,omitempty"`
	RefundStatus         *string                                                `json:"RefundStatus,omitempty" xml:"RefundStatus,omitempty"`
	PointsAmount         *int64                                                 `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	CashAmount           *string                                                `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	LogisticsList        *QueryOrderInfoAfterSaleResponseBodyModelLogisticsList `json:"LogisticsList,omitempty" xml:"LogisticsList,omitempty" type:"Struct"`
}

func (s QueryOrderInfoAfterSaleResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderInfoAfterSaleResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetLmOrderId(v int64) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetShopServiceTelephone(v string) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.ShopServiceTelephone = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetRefundAmount(v int64) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.RefundAmount = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetXiaomiCode(v string) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.XiaomiCode = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetShopName(v string) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.ShopName = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetCreateDate(v string) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.CreateDate = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetRefundRate(v string) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.RefundRate = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetExtJson(v string) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.ExtJson = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetOrderStatus(v string) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.OrderStatus = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetRefundPoints(v int64) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.RefundPoints = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetTbOrderId(v int64) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.TbOrderId = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetPoints(v int64) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.Points = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetRefundStatus(v string) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.RefundStatus = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetPointsAmount(v int64) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.PointsAmount = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetCashAmount(v string) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.CashAmount = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModel) SetLogisticsList(v *QueryOrderInfoAfterSaleResponseBodyModelLogisticsList) *QueryOrderInfoAfterSaleResponseBodyModel {
	s.LogisticsList = v
	return s
}

type QueryOrderInfoAfterSaleResponseBodyModelLogisticsList struct {
	Logistics []*QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics `json:"Logistics,omitempty" xml:"Logistics,omitempty" type:"Repeated"`
}

func (s QueryOrderInfoAfterSaleResponseBodyModelLogisticsList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderInfoAfterSaleResponseBodyModelLogisticsList) GoString() string {
	return s.String()
}

func (s *QueryOrderInfoAfterSaleResponseBodyModelLogisticsList) SetLogistics(v []*QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics) *QueryOrderInfoAfterSaleResponseBodyModelLogisticsList {
	s.Logistics = v
	return s
}

type QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics struct {
	LogisticsCompanyCode *string `json:"LogisticsCompanyCode,omitempty" xml:"LogisticsCompanyCode,omitempty"`
	LogisticsNo          *string `json:"LogisticsNo,omitempty" xml:"LogisticsNo,omitempty"`
	LogisticsStatus      *string `json:"LogisticsStatus,omitempty" xml:"LogisticsStatus,omitempty"`
	LogisticsCompanyName *string `json:"LogisticsCompanyName,omitempty" xml:"LogisticsCompanyName,omitempty"`
}

func (s QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics) GoString() string {
	return s.String()
}

func (s *QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics) SetLogisticsCompanyCode(v string) *QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics {
	s.LogisticsCompanyCode = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics) SetLogisticsNo(v string) *QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics {
	s.LogisticsNo = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics) SetLogisticsStatus(v string) *QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics {
	s.LogisticsStatus = &v
	return s
}

func (s *QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics) SetLogisticsCompanyName(v string) *QueryOrderInfoAfterSaleResponseBodyModelLogisticsListLogistics {
	s.LogisticsCompanyName = &v
	return s
}

type QueryOrderInfoAfterSaleResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrderInfoAfterSaleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrderInfoAfterSaleResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderInfoAfterSaleResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderInfoAfterSaleResponse) SetHeaders(v map[string]*string) *QueryOrderInfoAfterSaleResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderInfoAfterSaleResponse) SetBody(v *QueryOrderInfoAfterSaleResponseBody) *QueryOrderInfoAfterSaleResponse {
	s.Body = v
	return s
}

type QueryOrderItemInfoByPaymentIdForAiZhanYouRequest struct {
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid    *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	PaymentId *string `json:"PaymentId,omitempty" xml:"PaymentId,omitempty"`
}

func (s QueryOrderItemInfoByPaymentIdForAiZhanYouRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderItemInfoByPaymentIdForAiZhanYouRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouRequest) SetBizId(v string) *QueryOrderItemInfoByPaymentIdForAiZhanYouRequest {
	s.BizId = &v
	return s
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouRequest) SetBizUid(v string) *QueryOrderItemInfoByPaymentIdForAiZhanYouRequest {
	s.BizUid = &v
	return s
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouRequest) SetPaymentId(v string) *QueryOrderItemInfoByPaymentIdForAiZhanYouRequest {
	s.PaymentId = &v
	return s
}

type QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody struct {
	Code        *string                                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LmOrderList *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Struct"`
}

func (s QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody) SetCode(v string) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody) SetMessage(v string) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody) SetRequestId(v string) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody) SetLmOrderList(v *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderList) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody {
	s.LmOrderList = v
	return s
}

type QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderList struct {
	LmOrderList []*QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Repeated"`
}

func (s QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderList) SetLmOrderList(v []*QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderList {
	s.LmOrderList = v
	return s
}

type QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList struct {
	LmOrderId *int64  `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	ItemName  *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	SkuName   *string `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
	SkuId     *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	LmItemId  *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId    *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
}

func (s QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList) SetLmOrderId(v int64) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList) SetItemName(v string) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList {
	s.ItemName = &v
	return s
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList) SetSkuName(v string) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList {
	s.SkuName = &v
	return s
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList) SetSkuId(v int64) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList {
	s.SkuId = &v
	return s
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList) SetLmItemId(v string) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList {
	s.LmItemId = &v
	return s
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList) SetItemId(v int64) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBodyLmOrderListLmOrderList {
	s.ItemId = &v
	return s
}

type QueryOrderItemInfoByPaymentIdForAiZhanYouResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrderItemInfoByPaymentIdForAiZhanYouResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderItemInfoByPaymentIdForAiZhanYouResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponse) SetHeaders(v map[string]*string) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderItemInfoByPaymentIdForAiZhanYouResponse) SetBody(v *QueryOrderItemInfoByPaymentIdForAiZhanYouResponseBody) *QueryOrderItemInfoByPaymentIdForAiZhanYouResponse {
	s.Body = v
	return s
}

type QueryOrderListRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	PageSize              *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber            *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	FilterOption          *string `json:"FilterOption,omitempty" xml:"FilterOption,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s QueryOrderListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderListRequest) SetBizId(v string) *QueryOrderListRequest {
	s.BizId = &v
	return s
}

func (s *QueryOrderListRequest) SetBizUid(v string) *QueryOrderListRequest {
	s.BizUid = &v
	return s
}

func (s *QueryOrderListRequest) SetPageSize(v int64) *QueryOrderListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOrderListRequest) SetPageNumber(v int64) *QueryOrderListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryOrderListRequest) SetFilterOption(v string) *QueryOrderListRequest {
	s.FilterOption = &v
	return s
}

func (s *QueryOrderListRequest) SetUseAnonymousTbAccount(v bool) *QueryOrderListRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *QueryOrderListRequest) SetThirdPartyUserId(v string) *QueryOrderListRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *QueryOrderListRequest) SetAccountType(v string) *QueryOrderListRequest {
	s.AccountType = &v
	return s
}

type QueryOrderListResponseBody struct {
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code        *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize    *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount  *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PostFee     *QueryOrderListResponseBodyPostFee     `json:"PostFee,omitempty" xml:"PostFee,omitempty" type:"Struct"`
	LmOrderList *QueryOrderListResponseBodyLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Struct"`
}

func (s QueryOrderListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBody) SetRequestId(v string) *QueryOrderListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderListResponseBody) SetCode(v string) *QueryOrderListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrderListResponseBody) SetMessage(v string) *QueryOrderListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrderListResponseBody) SetPageSize(v int32) *QueryOrderListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryOrderListResponseBody) SetPageNumber(v int32) *QueryOrderListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryOrderListResponseBody) SetTotalCount(v int64) *QueryOrderListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryOrderListResponseBody) SetPostFee(v *QueryOrderListResponseBodyPostFee) *QueryOrderListResponseBody {
	s.PostFee = v
	return s
}

func (s *QueryOrderListResponseBody) SetLmOrderList(v *QueryOrderListResponseBodyLmOrderList) *QueryOrderListResponseBody {
	s.LmOrderList = v
	return s
}

type QueryOrderListResponseBodyPostFee struct {
	FundAmount      *int64  `json:"FundAmount,omitempty" xml:"FundAmount,omitempty"`
	FundType        *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	FundAmountMoney *int64  `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
}

func (s QueryOrderListResponseBodyPostFee) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyPostFee) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyPostFee) SetFundAmount(v int64) *QueryOrderListResponseBodyPostFee {
	s.FundAmount = &v
	return s
}

func (s *QueryOrderListResponseBodyPostFee) SetFundType(v string) *QueryOrderListResponseBodyPostFee {
	s.FundType = &v
	return s
}

func (s *QueryOrderListResponseBodyPostFee) SetFundAmountMoney(v int64) *QueryOrderListResponseBodyPostFee {
	s.FundAmountMoney = &v
	return s
}

type QueryOrderListResponseBodyLmOrderList struct {
	LmOrderList []*QueryOrderListResponseBodyLmOrderListLmOrderList `json:"LmOrderList,omitempty" xml:"LmOrderList,omitempty" type:"Repeated"`
}

func (s QueryOrderListResponseBodyLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyLmOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyLmOrderList) SetLmOrderList(v []*QueryOrderListResponseBodyLmOrderListLmOrderList) *QueryOrderListResponseBodyLmOrderList {
	s.LmOrderList = v
	return s
}

type QueryOrderListResponseBodyLmOrderListLmOrderList struct {
	LmOrderId           *int64                                                               `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	EnableStatus        *int32                                                               `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	ExtJson             *string                                                              `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	OrderStatus         *int32                                                               `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	ShopName            *string                                                              `json:"ShopName,omitempty" xml:"ShopName,omitempty"`
	TbOrderId           *int64                                                               `json:"TbOrderId,omitempty" xml:"TbOrderId,omitempty"`
	OrderAmount         *int64                                                               `json:"OrderAmount,omitempty" xml:"OrderAmount,omitempty"`
	LogisticsStatus     *int32                                                               `json:"LogisticsStatus,omitempty" xml:"LogisticsStatus,omitempty"`
	CreateDate          *string                                                              `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	FundStructureModels *QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModels `json:"FundStructureModels,omitempty" xml:"FundStructureModels,omitempty" type:"Struct"`
	SubOrderList        *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderList        `json:"SubOrderList,omitempty" xml:"SubOrderList,omitempty" type:"Struct"`
	PostFee             *QueryOrderListResponseBodyLmOrderListLmOrderListPostFee             `json:"PostFee,omitempty" xml:"PostFee,omitempty" type:"Struct"`
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderList) SetLmOrderId(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderList {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderList) SetEnableStatus(v int32) *QueryOrderListResponseBodyLmOrderListLmOrderList {
	s.EnableStatus = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderList) SetExtJson(v string) *QueryOrderListResponseBodyLmOrderListLmOrderList {
	s.ExtJson = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderList) SetOrderStatus(v int32) *QueryOrderListResponseBodyLmOrderListLmOrderList {
	s.OrderStatus = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderList) SetShopName(v string) *QueryOrderListResponseBodyLmOrderListLmOrderList {
	s.ShopName = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderList) SetTbOrderId(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderList {
	s.TbOrderId = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderList) SetOrderAmount(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderList {
	s.OrderAmount = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderList) SetLogisticsStatus(v int32) *QueryOrderListResponseBodyLmOrderListLmOrderList {
	s.LogisticsStatus = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderList) SetCreateDate(v string) *QueryOrderListResponseBodyLmOrderListLmOrderList {
	s.CreateDate = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderList) SetFundStructureModels(v *QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModels) *QueryOrderListResponseBodyLmOrderListLmOrderList {
	s.FundStructureModels = v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderList) SetSubOrderList(v *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderList) *QueryOrderListResponseBodyLmOrderListLmOrderList {
	s.SubOrderList = v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderList) SetPostFee(v *QueryOrderListResponseBodyLmOrderListLmOrderListPostFee) *QueryOrderListResponseBodyLmOrderListLmOrderList {
	s.PostFee = v
	return s
}

type QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModels struct {
	FundStructureModels []*QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels `json:"FundStructureModels,omitempty" xml:"FundStructureModels,omitempty" type:"Repeated"`
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModels) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModels) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModels) SetFundStructureModels(v []*QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels) *QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModels {
	s.FundStructureModels = v
	return s
}

type QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels struct {
	FundAmount      *int64  `json:"FundAmount,omitempty" xml:"FundAmount,omitempty"`
	FundType        *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	FundAmountMoney *int64  `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels) SetFundAmount(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels {
	s.FundAmount = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels) SetFundType(v string) *QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels {
	s.FundType = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels) SetFundAmountMoney(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderListFundStructureModelsFundStructureModels {
	s.FundAmountMoney = &v
	return s
}

type QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderList struct {
	SubOrderList []*QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList `json:"SubOrderList,omitempty" xml:"SubOrderList,omitempty" type:"Repeated"`
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderList) SetSubOrderList(v []*QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderList {
	s.SubOrderList = v
	return s
}

type QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList struct {
	LmOrderId     *int64                                                                                 `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	EnableStatus  *int32                                                                                 `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	ItemTitle     *string                                                                                `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	OrderStatus   *int32                                                                                 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	SkuName       *string                                                                                `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
	LmItemId      *string                                                                                `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	SkuId         *int64                                                                                 `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Number        *int64                                                                                 `json:"Number,omitempty" xml:"Number,omitempty"`
	TbOrderId     *int64                                                                                 `json:"TbOrderId,omitempty" xml:"TbOrderId,omitempty"`
	ItemPic       *string                                                                                `json:"ItemPic,omitempty" xml:"ItemPic,omitempty"`
	ItemId        *int64                                                                                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemPriceList *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList `json:"ItemPriceList,omitempty" xml:"ItemPriceList,omitempty" type:"Struct"`
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetLmOrderId(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetEnableStatus(v int32) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.EnableStatus = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetItemTitle(v string) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.ItemTitle = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetOrderStatus(v int32) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.OrderStatus = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetSkuName(v string) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.SkuName = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetLmItemId(v string) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.LmItemId = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetSkuId(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.SkuId = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetNumber(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.Number = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetTbOrderId(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.TbOrderId = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetItemPic(v string) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.ItemPic = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetItemId(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.ItemId = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList) SetItemPriceList(v *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderList {
	s.ItemPriceList = v
	return s
}

type QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList struct {
	ItemPriceList []*QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList `json:"ItemPriceList,omitempty" xml:"ItemPriceList,omitempty" type:"Repeated"`
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList) SetItemPriceList(v []*QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceList {
	s.ItemPriceList = v
	return s
}

type QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList struct {
	FundAmount      *int64  `json:"FundAmount,omitempty" xml:"FundAmount,omitempty"`
	FundType        *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	FundAmountMoney *int64  `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList) SetFundAmount(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList {
	s.FundAmount = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList) SetFundType(v string) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList {
	s.FundType = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList) SetFundAmountMoney(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderListSubOrderListSubOrderListItemPriceListItemPriceList {
	s.FundAmountMoney = &v
	return s
}

type QueryOrderListResponseBodyLmOrderListLmOrderListPostFee struct {
	FundAmount      *int64  `json:"FundAmount,omitempty" xml:"FundAmount,omitempty"`
	FundType        *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	FundAmountMoney *int64  `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListPostFee) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponseBodyLmOrderListLmOrderListPostFee) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListPostFee) SetFundAmount(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderListPostFee {
	s.FundAmount = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListPostFee) SetFundType(v string) *QueryOrderListResponseBodyLmOrderListLmOrderListPostFee {
	s.FundType = &v
	return s
}

func (s *QueryOrderListResponseBodyLmOrderListLmOrderListPostFee) SetFundAmountMoney(v int64) *QueryOrderListResponseBodyLmOrderListLmOrderListPostFee {
	s.FundAmountMoney = &v
	return s
}

type QueryOrderListResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrderListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrderListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderListResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderListResponse) SetHeaders(v map[string]*string) *QueryOrderListResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderListResponse) SetBody(v *QueryOrderListResponseBody) *QueryOrderListResponse {
	s.Body = v
	return s
}

type QueryOrderLogisticsRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	LmOrderId             *int64  `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s QueryOrderLogisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderLogisticsRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderLogisticsRequest) SetBizId(v string) *QueryOrderLogisticsRequest {
	s.BizId = &v
	return s
}

func (s *QueryOrderLogisticsRequest) SetBizUid(v string) *QueryOrderLogisticsRequest {
	s.BizUid = &v
	return s
}

func (s *QueryOrderLogisticsRequest) SetLmOrderId(v int64) *QueryOrderLogisticsRequest {
	s.LmOrderId = &v
	return s
}

func (s *QueryOrderLogisticsRequest) SetUseAnonymousTbAccount(v bool) *QueryOrderLogisticsRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *QueryOrderLogisticsRequest) SetThirdPartyUserId(v string) *QueryOrderLogisticsRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *QueryOrderLogisticsRequest) SetAccountType(v string) *QueryOrderLogisticsRequest {
	s.AccountType = &v
	return s
}

type QueryOrderLogisticsResponseBody struct {
	Code           *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderLogistics *QueryOrderLogisticsResponseBodyOrderLogistics `json:"OrderLogistics,omitempty" xml:"OrderLogistics,omitempty" type:"Struct"`
}

func (s QueryOrderLogisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderLogisticsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderLogisticsResponseBody) SetCode(v string) *QueryOrderLogisticsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrderLogisticsResponseBody) SetMessage(v string) *QueryOrderLogisticsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrderLogisticsResponseBody) SetRequestId(v string) *QueryOrderLogisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderLogisticsResponseBody) SetOrderLogistics(v *QueryOrderLogisticsResponseBodyOrderLogistics) *QueryOrderLogisticsResponseBody {
	s.OrderLogistics = v
	return s
}

type QueryOrderLogisticsResponseBodyOrderLogistics struct {
	LogisticsCompanyCode *string                                                           `json:"LogisticsCompanyCode,omitempty" xml:"LogisticsCompanyCode,omitempty"`
	LogisticsCompanyName *string                                                           `json:"LogisticsCompanyName,omitempty" xml:"LogisticsCompanyName,omitempty"`
	DataProvider         *string                                                           `json:"DataProvider,omitempty" xml:"DataProvider,omitempty"`
	DataProviderTitle    *string                                                           `json:"DataProviderTitle,omitempty" xml:"DataProviderTitle,omitempty"`
	LogisticsDetailList  *QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailList `json:"LogisticsDetailList,omitempty" xml:"LogisticsDetailList,omitempty" type:"Struct"`
	Receiver             *QueryOrderLogisticsResponseBodyOrderLogisticsReceiver            `json:"Receiver,omitempty" xml:"Receiver,omitempty" type:"Struct"`
	Fetcher              *QueryOrderLogisticsResponseBodyOrderLogisticsFetcher             `json:"Fetcher,omitempty" xml:"Fetcher,omitempty" type:"Struct"`
}

func (s QueryOrderLogisticsResponseBodyOrderLogistics) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderLogisticsResponseBodyOrderLogistics) GoString() string {
	return s.String()
}

func (s *QueryOrderLogisticsResponseBodyOrderLogistics) SetLogisticsCompanyCode(v string) *QueryOrderLogisticsResponseBodyOrderLogistics {
	s.LogisticsCompanyCode = &v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogistics) SetLogisticsCompanyName(v string) *QueryOrderLogisticsResponseBodyOrderLogistics {
	s.LogisticsCompanyName = &v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogistics) SetDataProvider(v string) *QueryOrderLogisticsResponseBodyOrderLogistics {
	s.DataProvider = &v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogistics) SetDataProviderTitle(v string) *QueryOrderLogisticsResponseBodyOrderLogistics {
	s.DataProviderTitle = &v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogistics) SetLogisticsDetailList(v *QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailList) *QueryOrderLogisticsResponseBodyOrderLogistics {
	s.LogisticsDetailList = v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogistics) SetReceiver(v *QueryOrderLogisticsResponseBodyOrderLogisticsReceiver) *QueryOrderLogisticsResponseBodyOrderLogistics {
	s.Receiver = v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogistics) SetFetcher(v *QueryOrderLogisticsResponseBodyOrderLogisticsFetcher) *QueryOrderLogisticsResponseBodyOrderLogistics {
	s.Fetcher = v
	return s
}

type QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailList struct {
	LogisticsDetailList []*QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailListLogisticsDetailList `json:"LogisticsDetailList,omitempty" xml:"LogisticsDetailList,omitempty" type:"Repeated"`
}

func (s QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailList) GoString() string {
	return s.String()
}

func (s *QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailList) SetLogisticsDetailList(v []*QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailListLogisticsDetailList) *QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailList {
	s.LogisticsDetailList = v
	return s
}

type QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailListLogisticsDetailList struct {
	OcurrTimeStr *string `json:"OcurrTimeStr,omitempty" xml:"OcurrTimeStr,omitempty"`
	StanderdDesc *string `json:"StanderdDesc,omitempty" xml:"StanderdDesc,omitempty"`
	StatusIcon   *string `json:"StatusIcon,omitempty" xml:"StatusIcon,omitempty"`
}

func (s QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailListLogisticsDetailList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailListLogisticsDetailList) GoString() string {
	return s.String()
}

func (s *QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailListLogisticsDetailList) SetOcurrTimeStr(v string) *QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailListLogisticsDetailList {
	s.OcurrTimeStr = &v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailListLogisticsDetailList) SetStanderdDesc(v string) *QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailListLogisticsDetailList {
	s.StanderdDesc = &v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailListLogisticsDetailList) SetStatusIcon(v string) *QueryOrderLogisticsResponseBodyOrderLogisticsLogisticsDetailListLogisticsDetailList {
	s.StatusIcon = &v
	return s
}

type QueryOrderLogisticsResponseBodyOrderLogisticsReceiver struct {
	Address     *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ZipCode     *string `json:"ZipCode,omitempty" xml:"ZipCode,omitempty"`
}

func (s QueryOrderLogisticsResponseBodyOrderLogisticsReceiver) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderLogisticsResponseBodyOrderLogisticsReceiver) GoString() string {
	return s.String()
}

func (s *QueryOrderLogisticsResponseBodyOrderLogisticsReceiver) SetAddress(v string) *QueryOrderLogisticsResponseBodyOrderLogisticsReceiver {
	s.Address = &v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogisticsReceiver) SetName(v string) *QueryOrderLogisticsResponseBodyOrderLogisticsReceiver {
	s.Name = &v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogisticsReceiver) SetPhoneNumber(v string) *QueryOrderLogisticsResponseBodyOrderLogisticsReceiver {
	s.PhoneNumber = &v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogisticsReceiver) SetZipCode(v string) *QueryOrderLogisticsResponseBodyOrderLogisticsReceiver {
	s.ZipCode = &v
	return s
}

type QueryOrderLogisticsResponseBodyOrderLogisticsFetcher struct {
	Address     *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ZipCode     *string `json:"ZipCode,omitempty" xml:"ZipCode,omitempty"`
}

func (s QueryOrderLogisticsResponseBodyOrderLogisticsFetcher) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderLogisticsResponseBodyOrderLogisticsFetcher) GoString() string {
	return s.String()
}

func (s *QueryOrderLogisticsResponseBodyOrderLogisticsFetcher) SetAddress(v string) *QueryOrderLogisticsResponseBodyOrderLogisticsFetcher {
	s.Address = &v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogisticsFetcher) SetName(v string) *QueryOrderLogisticsResponseBodyOrderLogisticsFetcher {
	s.Name = &v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogisticsFetcher) SetPhoneNumber(v string) *QueryOrderLogisticsResponseBodyOrderLogisticsFetcher {
	s.PhoneNumber = &v
	return s
}

func (s *QueryOrderLogisticsResponseBodyOrderLogisticsFetcher) SetZipCode(v string) *QueryOrderLogisticsResponseBodyOrderLogisticsFetcher {
	s.ZipCode = &v
	return s
}

type QueryOrderLogisticsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrderLogisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrderLogisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderLogisticsResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderLogisticsResponse) SetHeaders(v map[string]*string) *QueryOrderLogisticsResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderLogisticsResponse) SetBody(v *QueryOrderLogisticsResponseBody) *QueryOrderLogisticsResponse {
	s.Body = v
	return s
}

type QueryRefundApplicationDetailRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	SubLmOrderId          *string `json:"SubLmOrderId,omitempty" xml:"SubLmOrderId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s QueryRefundApplicationDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRefundApplicationDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRefundApplicationDetailRequest) SetBizId(v string) *QueryRefundApplicationDetailRequest {
	s.BizId = &v
	return s
}

func (s *QueryRefundApplicationDetailRequest) SetBizUid(v string) *QueryRefundApplicationDetailRequest {
	s.BizUid = &v
	return s
}

func (s *QueryRefundApplicationDetailRequest) SetSubLmOrderId(v string) *QueryRefundApplicationDetailRequest {
	s.SubLmOrderId = &v
	return s
}

func (s *QueryRefundApplicationDetailRequest) SetUseAnonymousTbAccount(v bool) *QueryRefundApplicationDetailRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *QueryRefundApplicationDetailRequest) SetThirdPartyUserId(v string) *QueryRefundApplicationDetailRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *QueryRefundApplicationDetailRequest) SetAccountType(v string) *QueryRefundApplicationDetailRequest {
	s.AccountType = &v
	return s
}

type QueryRefundApplicationDetailResponseBody struct {
	Code                    *string                                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                 *string                                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId               *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RefundApplicationDetail *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail `json:"RefundApplicationDetail,omitempty" xml:"RefundApplicationDetail,omitempty" type:"Struct"`
}

func (s QueryRefundApplicationDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRefundApplicationDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRefundApplicationDetailResponseBody) SetCode(v string) *QueryRefundApplicationDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBody) SetMessage(v string) *QueryRefundApplicationDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBody) SetRequestId(v string) *QueryRefundApplicationDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBody) SetRefundApplicationDetail(v *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) *QueryRefundApplicationDetailResponseBody {
	s.RefundApplicationDetail = v
	return s
}

type QueryRefundApplicationDetailResponseBodyRefundApplicationDetail struct {
	DisputeType                  *int32                                                                           `json:"DisputeType,omitempty" xml:"DisputeType,omitempty"`
	LmOrderId                    *string                                                                          `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	RefunderName                 *string                                                                          `json:"RefunderName,omitempty" xml:"RefunderName,omitempty"`
	SellerAgreeMsg               *string                                                                          `json:"SellerAgreeMsg,omitempty" xml:"SellerAgreeMsg,omitempty"`
	ReturnGoodLogisticsStatus    *int32                                                                           `json:"ReturnGoodLogisticsStatus,omitempty" xml:"ReturnGoodLogisticsStatus,omitempty"`
	DisputeDesc                  *string                                                                          `json:"DisputeDesc,omitempty" xml:"DisputeDesc,omitempty"`
	RefunderAddress              *string                                                                          `json:"RefunderAddress,omitempty" xml:"RefunderAddress,omitempty"`
	ReturnGoodCount              *int32                                                                           `json:"ReturnGoodCount,omitempty" xml:"ReturnGoodCount,omitempty"`
	DisputeId                    *int64                                                                           `json:"DisputeId,omitempty" xml:"DisputeId,omitempty"`
	DisputeEndTime               *string                                                                          `json:"DisputeEndTime,omitempty" xml:"DisputeEndTime,omitempty"`
	OrderLogisticsStatus         *int32                                                                           `json:"OrderLogisticsStatus,omitempty" xml:"OrderLogisticsStatus,omitempty"`
	BizClaimType                 *int32                                                                           `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	RealRefundFee                *int64                                                                           `json:"RealRefundFee,omitempty" xml:"RealRefundFee,omitempty"`
	RefundFee                    *int64                                                                           `json:"RefundFee,omitempty" xml:"RefundFee,omitempty"`
	SellerRefuseAgreementMessage *string                                                                          `json:"SellerRefuseAgreementMessage,omitempty" xml:"SellerRefuseAgreementMessage,omitempty"`
	DisputeCreateTime            *string                                                                          `json:"DisputeCreateTime,omitempty" xml:"DisputeCreateTime,omitempty"`
	RefunderTel                  *string                                                                          `json:"RefunderTel,omitempty" xml:"RefunderTel,omitempty"`
	SellerRefuseReason           *string                                                                          `json:"SellerRefuseReason,omitempty" xml:"SellerRefuseReason,omitempty"`
	SubLmOrderId                 *string                                                                          `json:"SubLmOrderId,omitempty" xml:"SubLmOrderId,omitempty"`
	ApplyDisputeDesc             *string                                                                          `json:"ApplyDisputeDesc,omitempty" xml:"ApplyDisputeDesc,omitempty"`
	DisputeStatus                *int32                                                                           `json:"DisputeStatus,omitempty" xml:"DisputeStatus,omitempty"`
	RefunderZipCode              *string                                                                          `json:"RefunderZipCode,omitempty" xml:"RefunderZipCode,omitempty"`
	MaxRefundFeeData             *QueryRefundApplicationDetailResponseBodyRefundApplicationDetailMaxRefundFeeData `json:"MaxRefundFeeData,omitempty" xml:"MaxRefundFeeData,omitempty" type:"Struct"`
	ApplyReasonText              *QueryRefundApplicationDetailResponseBodyRefundApplicationDetailApplyReasonText  `json:"ApplyReasonText,omitempty" xml:"ApplyReasonText,omitempty" type:"Struct"`
}

func (s QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) String() string {
	return tea.Prettify(s)
}

func (s QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) GoString() string {
	return s.String()
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetDisputeType(v int32) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.DisputeType = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetLmOrderId(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.LmOrderId = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetRefunderName(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.RefunderName = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetSellerAgreeMsg(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.SellerAgreeMsg = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetReturnGoodLogisticsStatus(v int32) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.ReturnGoodLogisticsStatus = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetDisputeDesc(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.DisputeDesc = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetRefunderAddress(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.RefunderAddress = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetReturnGoodCount(v int32) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.ReturnGoodCount = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetDisputeId(v int64) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.DisputeId = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetDisputeEndTime(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.DisputeEndTime = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetOrderLogisticsStatus(v int32) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.OrderLogisticsStatus = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetBizClaimType(v int32) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.BizClaimType = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetRealRefundFee(v int64) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.RealRefundFee = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetRefundFee(v int64) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.RefundFee = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetSellerRefuseAgreementMessage(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.SellerRefuseAgreementMessage = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetDisputeCreateTime(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.DisputeCreateTime = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetRefunderTel(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.RefunderTel = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetSellerRefuseReason(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.SellerRefuseReason = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetSubLmOrderId(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.SubLmOrderId = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetApplyDisputeDesc(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.ApplyDisputeDesc = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetDisputeStatus(v int32) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.DisputeStatus = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetRefunderZipCode(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.RefunderZipCode = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetMaxRefundFeeData(v *QueryRefundApplicationDetailResponseBodyRefundApplicationDetailMaxRefundFeeData) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.MaxRefundFeeData = v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail) SetApplyReasonText(v *QueryRefundApplicationDetailResponseBodyRefundApplicationDetailApplyReasonText) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetail {
	s.ApplyReasonText = v
	return s
}

type QueryRefundApplicationDetailResponseBodyRefundApplicationDetailMaxRefundFeeData struct {
	MaxRefundFee *int32 `json:"MaxRefundFee,omitempty" xml:"MaxRefundFee,omitempty"`
	MinRefundFee *int32 `json:"MinRefundFee,omitempty" xml:"MinRefundFee,omitempty"`
}

func (s QueryRefundApplicationDetailResponseBodyRefundApplicationDetailMaxRefundFeeData) String() string {
	return tea.Prettify(s)
}

func (s QueryRefundApplicationDetailResponseBodyRefundApplicationDetailMaxRefundFeeData) GoString() string {
	return s.String()
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetailMaxRefundFeeData) SetMaxRefundFee(v int32) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetailMaxRefundFeeData {
	s.MaxRefundFee = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetailMaxRefundFeeData) SetMinRefundFee(v int32) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetailMaxRefundFeeData {
	s.MinRefundFee = &v
	return s
}

type QueryRefundApplicationDetailResponseBodyRefundApplicationDetailApplyReasonText struct {
	ReasonTips   *string `json:"ReasonTips,omitempty" xml:"ReasonTips,omitempty"`
	ReasonTextId *int64  `json:"ReasonTextId,omitempty" xml:"ReasonTextId,omitempty"`
}

func (s QueryRefundApplicationDetailResponseBodyRefundApplicationDetailApplyReasonText) String() string {
	return tea.Prettify(s)
}

func (s QueryRefundApplicationDetailResponseBodyRefundApplicationDetailApplyReasonText) GoString() string {
	return s.String()
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetailApplyReasonText) SetReasonTips(v string) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetailApplyReasonText {
	s.ReasonTips = &v
	return s
}

func (s *QueryRefundApplicationDetailResponseBodyRefundApplicationDetailApplyReasonText) SetReasonTextId(v int64) *QueryRefundApplicationDetailResponseBodyRefundApplicationDetailApplyReasonText {
	s.ReasonTextId = &v
	return s
}

type QueryRefundApplicationDetailResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRefundApplicationDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRefundApplicationDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRefundApplicationDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryRefundApplicationDetailResponse) SetHeaders(v map[string]*string) *QueryRefundApplicationDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryRefundApplicationDetailResponse) SetBody(v *QueryRefundApplicationDetailResponseBody) *QueryRefundApplicationDetailResponse {
	s.Body = v
	return s
}

type QueryStatementsRequest struct {
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	PayeeIds     *string `json:"PayeeIds,omitempty" xml:"PayeeIds,omitempty"`
	SettleNoes   *string `json:"SettleNoes,omitempty" xml:"SettleNoes,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	SettleStatus *string `json:"SettleStatus,omitempty" xml:"SettleStatus,omitempty"`
	SettleType   *string `json:"SettleType,omitempty" xml:"SettleType,omitempty"`
	ExtInfo      *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s QueryStatementsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStatementsRequest) GoString() string {
	return s.String()
}

func (s *QueryStatementsRequest) SetTenantId(v string) *QueryStatementsRequest {
	s.TenantId = &v
	return s
}

func (s *QueryStatementsRequest) SetPayeeIds(v string) *QueryStatementsRequest {
	s.PayeeIds = &v
	return s
}

func (s *QueryStatementsRequest) SetSettleNoes(v string) *QueryStatementsRequest {
	s.SettleNoes = &v
	return s
}

func (s *QueryStatementsRequest) SetStartTime(v string) *QueryStatementsRequest {
	s.StartTime = &v
	return s
}

func (s *QueryStatementsRequest) SetEndTime(v string) *QueryStatementsRequest {
	s.EndTime = &v
	return s
}

func (s *QueryStatementsRequest) SetPageSize(v int32) *QueryStatementsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryStatementsRequest) SetPageNumber(v int32) *QueryStatementsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryStatementsRequest) SetSettleStatus(v string) *QueryStatementsRequest {
	s.SettleStatus = &v
	return s
}

func (s *QueryStatementsRequest) SetSettleType(v string) *QueryStatementsRequest {
	s.SettleType = &v
	return s
}

func (s *QueryStatementsRequest) SetExtInfo(v string) *QueryStatementsRequest {
	s.ExtInfo = &v
	return s
}

type QueryStatementsResponseBody struct {
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                           `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                           `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                           `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Model      *QueryStatementsResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s QueryStatementsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStatementsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStatementsResponseBody) SetRequestId(v string) *QueryStatementsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStatementsResponseBody) SetSuccess(v bool) *QueryStatementsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryStatementsResponseBody) SetSubMessage(v string) *QueryStatementsResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryStatementsResponseBody) SetCode(v string) *QueryStatementsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStatementsResponseBody) SetMessage(v string) *QueryStatementsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryStatementsResponseBody) SetSubCode(v string) *QueryStatementsResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryStatementsResponseBody) SetTotalCount(v int64) *QueryStatementsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryStatementsResponseBody) SetLogsId(v string) *QueryStatementsResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryStatementsResponseBody) SetModel(v *QueryStatementsResponseBodyModel) *QueryStatementsResponseBody {
	s.Model = v
	return s
}

type QueryStatementsResponseBodyModel struct {
	PageSize      *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber    *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount    *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	StatementList []*QueryStatementsResponseBodyModelStatementList `json:"StatementList,omitempty" xml:"StatementList,omitempty" type:"Repeated"`
}

func (s QueryStatementsResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryStatementsResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryStatementsResponseBodyModel) SetPageSize(v int32) *QueryStatementsResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *QueryStatementsResponseBodyModel) SetPageNumber(v int32) *QueryStatementsResponseBodyModel {
	s.PageNumber = &v
	return s
}

func (s *QueryStatementsResponseBodyModel) SetTotalCount(v int32) *QueryStatementsResponseBodyModel {
	s.TotalCount = &v
	return s
}

func (s *QueryStatementsResponseBodyModel) SetStatementList(v []*QueryStatementsResponseBodyModelStatementList) *QueryStatementsResponseBodyModel {
	s.StatementList = v
	return s
}

type QueryStatementsResponseBodyModelStatementList struct {
	PayeeAccountName *string `json:"PayeeAccountName,omitempty" xml:"PayeeAccountName,omitempty"`
	PayeeId          *string `json:"PayeeId,omitempty" xml:"PayeeId,omitempty"`
	PayeeAccountNo   *string `json:"PayeeAccountNo,omitempty" xml:"PayeeAccountNo,omitempty"`
	CreateDate       *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	ExtInfo          *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SettleNo         *string `json:"SettleNo,omitempty" xml:"SettleNo,omitempty"`
	Attributes       *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	SettleStatus     *string `json:"SettleStatus,omitempty" xml:"SettleStatus,omitempty"`
	PayeeName        *string `json:"PayeeName,omitempty" xml:"PayeeName,omitempty"`
	SettleAmount     *string `json:"SettleAmount,omitempty" xml:"SettleAmount,omitempty"`
	PayeeAccountId   *string `json:"PayeeAccountId,omitempty" xml:"PayeeAccountId,omitempty"`
	ModifiedDate     *string `json:"ModifiedDate,omitempty" xml:"ModifiedDate,omitempty"`
	StatusMessage    *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	TenantId         *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryStatementsResponseBodyModelStatementList) String() string {
	return tea.Prettify(s)
}

func (s QueryStatementsResponseBodyModelStatementList) GoString() string {
	return s.String()
}

func (s *QueryStatementsResponseBodyModelStatementList) SetPayeeAccountName(v string) *QueryStatementsResponseBodyModelStatementList {
	s.PayeeAccountName = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetPayeeId(v string) *QueryStatementsResponseBodyModelStatementList {
	s.PayeeId = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetPayeeAccountNo(v string) *QueryStatementsResponseBodyModelStatementList {
	s.PayeeAccountNo = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetCreateDate(v string) *QueryStatementsResponseBodyModelStatementList {
	s.CreateDate = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetExtInfo(v string) *QueryStatementsResponseBodyModelStatementList {
	s.ExtInfo = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetEndTime(v string) *QueryStatementsResponseBodyModelStatementList {
	s.EndTime = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetStartTime(v string) *QueryStatementsResponseBodyModelStatementList {
	s.StartTime = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetSettleNo(v string) *QueryStatementsResponseBodyModelStatementList {
	s.SettleNo = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetAttributes(v string) *QueryStatementsResponseBodyModelStatementList {
	s.Attributes = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetSettleStatus(v string) *QueryStatementsResponseBodyModelStatementList {
	s.SettleStatus = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetPayeeName(v string) *QueryStatementsResponseBodyModelStatementList {
	s.PayeeName = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetSettleAmount(v string) *QueryStatementsResponseBodyModelStatementList {
	s.SettleAmount = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetPayeeAccountId(v string) *QueryStatementsResponseBodyModelStatementList {
	s.PayeeAccountId = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetModifiedDate(v string) *QueryStatementsResponseBodyModelStatementList {
	s.ModifiedDate = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetStatusMessage(v string) *QueryStatementsResponseBodyModelStatementList {
	s.StatusMessage = &v
	return s
}

func (s *QueryStatementsResponseBodyModelStatementList) SetTenantId(v string) *QueryStatementsResponseBodyModelStatementList {
	s.TenantId = &v
	return s
}

type QueryStatementsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryStatementsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryStatementsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStatementsResponse) GoString() string {
	return s.String()
}

func (s *QueryStatementsResponse) SetHeaders(v map[string]*string) *QueryStatementsResponse {
	s.Headers = v
	return s
}

func (s *QueryStatementsResponse) SetBody(v *QueryStatementsResponseBody) *QueryStatementsResponse {
	s.Body = v
	return s
}

type QueryUnfinishedActivitiesRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryUnfinishedActivitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedActivitiesRequest) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedActivitiesRequest) SetBizId(v string) *QueryUnfinishedActivitiesRequest {
	s.BizId = &v
	return s
}

func (s *QueryUnfinishedActivitiesRequest) SetPageNumber(v int32) *QueryUnfinishedActivitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryUnfinishedActivitiesRequest) SetPageSize(v int32) *QueryUnfinishedActivitiesRequest {
	s.PageSize = &v
	return s
}

type QueryUnfinishedActivitiesResponseBody struct {
	RequestId              *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code                   *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize               *int64                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber             *int64                                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount             *int64                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LmActivityModelExtList *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtList `json:"LmActivityModelExtList,omitempty" xml:"LmActivityModelExtList,omitempty" type:"Struct"`
}

func (s QueryUnfinishedActivitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedActivitiesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedActivitiesResponseBody) SetRequestId(v string) *QueryUnfinishedActivitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBody) SetCode(v string) *QueryUnfinishedActivitiesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBody) SetMessage(v string) *QueryUnfinishedActivitiesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBody) SetPageSize(v int64) *QueryUnfinishedActivitiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBody) SetPageNumber(v int64) *QueryUnfinishedActivitiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBody) SetTotalCount(v int64) *QueryUnfinishedActivitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBody) SetLmActivityModelExtList(v *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtList) *QueryUnfinishedActivitiesResponseBody {
	s.LmActivityModelExtList = v
	return s
}

type QueryUnfinishedActivitiesResponseBodyLmActivityModelExtList struct {
	LmActivityModelV2Ext []*QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext `json:"LmActivityModelV2Ext,omitempty" xml:"LmActivityModelV2Ext,omitempty" type:"Repeated"`
}

func (s QueryUnfinishedActivitiesResponseBodyLmActivityModelExtList) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedActivitiesResponseBodyLmActivityModelExtList) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtList) SetLmActivityModelV2Ext(v []*QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtList {
	s.LmActivityModelV2Ext = v
	return s
}

type QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext struct {
	EndDate                 *string                                                                                                 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Description             *string                                                                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	StartDate               *string                                                                                                 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	LmActivityId            *int64                                                                                                  `json:"LmActivityId,omitempty" xml:"LmActivityId,omitempty"`
	BizId                   *string                                                                                                 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ActivityPicUrl          *string                                                                                                 `json:"ActivityPicUrl,omitempty" xml:"ActivityPicUrl,omitempty"`
	Name                    *string                                                                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	LmActivitySessionModels *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModels `json:"LmActivitySessionModels,omitempty" xml:"LmActivitySessionModels,omitempty" type:"Struct"`
}

func (s QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext) SetEndDate(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext {
	s.EndDate = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext) SetDescription(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext {
	s.Description = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext) SetStartDate(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext {
	s.StartDate = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext) SetLmActivityId(v int64) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext {
	s.LmActivityId = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext) SetBizId(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext {
	s.BizId = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext) SetActivityPicUrl(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext {
	s.ActivityPicUrl = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext) SetName(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext {
	s.Name = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext) SetLmActivitySessionModels(v *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModels) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2Ext {
	s.LmActivitySessionModels = v
	return s
}

type QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModels struct {
	LmActivitySessionModel []*QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel `json:"LmActivitySessionModel,omitempty" xml:"LmActivitySessionModel,omitempty" type:"Repeated"`
}

func (s QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModels) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModels) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModels) SetLmActivitySessionModel(v []*QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModels {
	s.LmActivitySessionModel = v
	return s
}

type QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel struct {
	SubBizCode   *string                `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	EndDate      *string                `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	DisplayDate  *string                `json:"DisplayDate,omitempty" xml:"DisplayDate,omitempty"`
	Description  *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	LmSessionId  *int64                 `json:"LmSessionId,omitempty" xml:"LmSessionId,omitempty"`
	StartDate    *string                `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	BizId        *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LmActivityId *int64                 `json:"LmActivityId,omitempty" xml:"LmActivityId,omitempty"`
	Name         *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	ExtInfo      map[string]interface{} `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) SetSubBizCode(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel {
	s.SubBizCode = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) SetEndDate(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel {
	s.EndDate = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) SetDisplayDate(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel {
	s.DisplayDate = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) SetDescription(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel {
	s.Description = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) SetLmSessionId(v int64) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel {
	s.LmSessionId = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) SetStartDate(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel {
	s.StartDate = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) SetBizId(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel {
	s.BizId = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) SetLmActivityId(v int64) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel {
	s.LmActivityId = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) SetName(v string) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel {
	s.Name = &v
	return s
}

func (s *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel) SetExtInfo(v map[string]interface{}) *QueryUnfinishedActivitiesResponseBodyLmActivityModelExtListLmActivityModelV2ExtLmActivitySessionModelsLmActivitySessionModel {
	s.ExtInfo = v
	return s
}

type QueryUnfinishedActivitiesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUnfinishedActivitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUnfinishedActivitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedActivitiesResponse) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedActivitiesResponse) SetHeaders(v map[string]*string) *QueryUnfinishedActivitiesResponse {
	s.Headers = v
	return s
}

func (s *QueryUnfinishedActivitiesResponse) SetBody(v *QueryUnfinishedActivitiesResponseBody) *QueryUnfinishedActivitiesResponse {
	s.Body = v
	return s
}

type QueryUnfinishedSessionsRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryTime  *int64  `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
}

func (s QueryUnfinishedSessionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessionsRequest) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessionsRequest) SetBizId(v string) *QueryUnfinishedSessionsRequest {
	s.BizId = &v
	return s
}

func (s *QueryUnfinishedSessionsRequest) SetPageNumber(v int32) *QueryUnfinishedSessionsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryUnfinishedSessionsRequest) SetPageSize(v int32) *QueryUnfinishedSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryUnfinishedSessionsRequest) SetQueryTime(v int64) *QueryUnfinishedSessionsRequest {
	s.QueryTime = &v
	return s
}

type QueryUnfinishedSessionsResponseBody struct {
	RequestId                  *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code                       *string                                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                    *string                                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber                 *int64                                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                   *int64                                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount                 *int64                                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LmActivitySessionModelList *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelList `json:"LmActivitySessionModelList,omitempty" xml:"LmActivitySessionModelList,omitempty" type:"Struct"`
}

func (s QueryUnfinishedSessionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessionsResponseBody) SetRequestId(v string) *QueryUnfinishedSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBody) SetCode(v string) *QueryUnfinishedSessionsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBody) SetMessage(v string) *QueryUnfinishedSessionsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBody) SetPageNumber(v int64) *QueryUnfinishedSessionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBody) SetPageSize(v int64) *QueryUnfinishedSessionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBody) SetTotalCount(v int64) *QueryUnfinishedSessionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBody) SetLmActivitySessionModelList(v *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelList) *QueryUnfinishedSessionsResponseBody {
	s.LmActivitySessionModelList = v
	return s
}

type QueryUnfinishedSessionsResponseBodyLmActivitySessionModelList struct {
	LmActivitySessionModel []*QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel `json:"LmActivitySessionModel,omitempty" xml:"LmActivitySessionModel,omitempty" type:"Repeated"`
}

func (s QueryUnfinishedSessionsResponseBodyLmActivitySessionModelList) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessionsResponseBodyLmActivitySessionModelList) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelList) SetLmActivitySessionModel(v []*QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel) *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelList {
	s.LmActivitySessionModel = v
	return s
}

type QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel struct {
	SubBizCode   *string `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	DisplayDate  *string `json:"DisplayDate,omitempty" xml:"DisplayDate,omitempty"`
	EndDate      *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LmSessionId  *int64  `json:"LmSessionId,omitempty" xml:"LmSessionId,omitempty"`
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LmActivityId *int64  `json:"LmActivityId,omitempty" xml:"LmActivityId,omitempty"`
	StartDate    *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel) SetSubBizCode(v string) *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel {
	s.SubBizCode = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel) SetDisplayDate(v string) *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel {
	s.DisplayDate = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel) SetEndDate(v string) *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel {
	s.EndDate = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel) SetDescription(v string) *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel {
	s.Description = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel) SetLmSessionId(v int64) *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel {
	s.LmSessionId = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel) SetBizId(v string) *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel {
	s.BizId = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel) SetLmActivityId(v int64) *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel {
	s.LmActivityId = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel) SetStartDate(v string) *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel {
	s.StartDate = &v
	return s
}

func (s *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel) SetName(v string) *QueryUnfinishedSessionsResponseBodyLmActivitySessionModelListLmActivitySessionModel {
	s.Name = &v
	return s
}

type QueryUnfinishedSessionsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUnfinishedSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUnfinishedSessionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessionsResponse) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessionsResponse) SetHeaders(v map[string]*string) *QueryUnfinishedSessionsResponse {
	s.Headers = v
	return s
}

func (s *QueryUnfinishedSessionsResponse) SetBody(v *QueryUnfinishedSessionsResponseBody) *QueryUnfinishedSessionsResponse {
	s.Body = v
	return s
}

type QueryUnfinishedSessions4ItemsRequest struct {
	BizId     *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LmItemIds map[string]interface{} `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty"`
	ItemIds   map[string]interface{} `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
	QueryTime *int64                 `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
}

func (s QueryUnfinishedSessions4ItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessions4ItemsRequest) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessions4ItemsRequest) SetBizId(v string) *QueryUnfinishedSessions4ItemsRequest {
	s.BizId = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsRequest) SetLmItemIds(v map[string]interface{}) *QueryUnfinishedSessions4ItemsRequest {
	s.LmItemIds = v
	return s
}

func (s *QueryUnfinishedSessions4ItemsRequest) SetItemIds(v map[string]interface{}) *QueryUnfinishedSessions4ItemsRequest {
	s.ItemIds = v
	return s
}

func (s *QueryUnfinishedSessions4ItemsRequest) SetQueryTime(v int64) *QueryUnfinishedSessions4ItemsRequest {
	s.QueryTime = &v
	return s
}

type QueryUnfinishedSessions4ItemsShrinkRequest struct {
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LmItemIdsShrink *string `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty"`
	ItemIdsShrink   *string `json:"ItemIds,omitempty" xml:"ItemIds,omitempty"`
	QueryTime       *int64  `json:"QueryTime,omitempty" xml:"QueryTime,omitempty"`
}

func (s QueryUnfinishedSessions4ItemsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessions4ItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessions4ItemsShrinkRequest) SetBizId(v string) *QueryUnfinishedSessions4ItemsShrinkRequest {
	s.BizId = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsShrinkRequest) SetLmItemIdsShrink(v string) *QueryUnfinishedSessions4ItemsShrinkRequest {
	s.LmItemIdsShrink = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsShrinkRequest) SetItemIdsShrink(v string) *QueryUnfinishedSessions4ItemsShrinkRequest {
	s.ItemIdsShrink = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsShrinkRequest) SetQueryTime(v int64) *QueryUnfinishedSessions4ItemsShrinkRequest {
	s.QueryTime = &v
	return s
}

type QueryUnfinishedSessions4ItemsResponseBody struct {
	Code                               *string                                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                            *string                                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId                          *string                                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LmItemActivitySessionModelListList *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListList `json:"LmItemActivitySessionModelListList,omitempty" xml:"LmItemActivitySessionModelListList,omitempty" type:"Struct"`
}

func (s QueryUnfinishedSessions4ItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessions4ItemsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessions4ItemsResponseBody) SetCode(v string) *QueryUnfinishedSessions4ItemsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBody) SetMessage(v string) *QueryUnfinishedSessions4ItemsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBody) SetRequestId(v string) *QueryUnfinishedSessions4ItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBody) SetLmItemActivitySessionModelListList(v *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListList) *QueryUnfinishedSessions4ItemsResponseBody {
	s.LmItemActivitySessionModelListList = v
	return s
}

type QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListList struct {
	LmItemActivitySessionModelList []*QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelList `json:"LmItemActivitySessionModelList,omitempty" xml:"LmItemActivitySessionModelList,omitempty" type:"Repeated"`
}

func (s QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListList) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListList) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListList) SetLmItemActivitySessionModelList(v []*QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelList) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListList {
	s.LmItemActivitySessionModelList = v
	return s
}

type QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelList struct {
	LmItemId                *string                                                                                                                           `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	ItemId                  *int64                                                                                                                            `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmActivitySessionModels *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModels `json:"LmActivitySessionModels,omitempty" xml:"LmActivitySessionModels,omitempty" type:"Struct"`
}

func (s QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelList) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelList) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelList) SetLmItemId(v string) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelList {
	s.LmItemId = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelList) SetItemId(v int64) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelList {
	s.ItemId = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelList) SetLmActivitySessionModels(v *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModels) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelList {
	s.LmActivitySessionModels = v
	return s
}

type QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModels struct {
	LmActivitySessionModel []*QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel `json:"LmActivitySessionModel,omitempty" xml:"LmActivitySessionModel,omitempty" type:"Repeated"`
}

func (s QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModels) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModels) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModels) SetLmActivitySessionModel(v []*QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModels {
	s.LmActivitySessionModel = v
	return s
}

type QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel struct {
	SubBizCode   *string                `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	EndDate      *string                `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	DisplayDate  *string                `json:"DisplayDate,omitempty" xml:"DisplayDate,omitempty"`
	LmSessionId  *int64                 `json:"LmSessionId,omitempty" xml:"LmSessionId,omitempty"`
	Description  *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	StartDate    *string                `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	LmActivityId *int64                 `json:"LmActivityId,omitempty" xml:"LmActivityId,omitempty"`
	BizId        *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Name         *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	ExtInfo      map[string]interface{} `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) SetSubBizCode(v string) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel {
	s.SubBizCode = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) SetEndDate(v string) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel {
	s.EndDate = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) SetDisplayDate(v string) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel {
	s.DisplayDate = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) SetLmSessionId(v int64) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel {
	s.LmSessionId = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) SetDescription(v string) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel {
	s.Description = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) SetStartDate(v string) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel {
	s.StartDate = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) SetLmActivityId(v int64) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel {
	s.LmActivityId = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) SetBizId(v string) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel {
	s.BizId = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) SetName(v string) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel {
	s.Name = &v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel) SetExtInfo(v map[string]interface{}) *QueryUnfinishedSessions4ItemsResponseBodyLmItemActivitySessionModelListListLmItemActivitySessionModelListLmActivitySessionModelsLmActivitySessionModel {
	s.ExtInfo = v
	return s
}

type QueryUnfinishedSessions4ItemsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUnfinishedSessions4ItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUnfinishedSessions4ItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUnfinishedSessions4ItemsResponse) GoString() string {
	return s.String()
}

func (s *QueryUnfinishedSessions4ItemsResponse) SetHeaders(v map[string]*string) *QueryUnfinishedSessions4ItemsResponse {
	s.Headers = v
	return s
}

func (s *QueryUnfinishedSessions4ItemsResponse) SetBody(v *QueryUnfinishedSessions4ItemsResponseBody) *QueryUnfinishedSessions4ItemsResponse {
	s.Body = v
	return s
}

type QueryUpcomingMoviesRequest struct {
	BizId    *string                `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CityCode *int64                 `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	ExtJson  map[string]interface{} `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s QueryUpcomingMoviesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUpcomingMoviesRequest) GoString() string {
	return s.String()
}

func (s *QueryUpcomingMoviesRequest) SetBizId(v string) *QueryUpcomingMoviesRequest {
	s.BizId = &v
	return s
}

func (s *QueryUpcomingMoviesRequest) SetCityCode(v int64) *QueryUpcomingMoviesRequest {
	s.CityCode = &v
	return s
}

func (s *QueryUpcomingMoviesRequest) SetExtJson(v map[string]interface{}) *QueryUpcomingMoviesRequest {
	s.ExtJson = v
	return s
}

type QueryUpcomingMoviesShrinkRequest struct {
	BizId         *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CityCode      *int64  `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	ExtJsonShrink *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s QueryUpcomingMoviesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUpcomingMoviesShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryUpcomingMoviesShrinkRequest) SetBizId(v string) *QueryUpcomingMoviesShrinkRequest {
	s.BizId = &v
	return s
}

func (s *QueryUpcomingMoviesShrinkRequest) SetCityCode(v int64) *QueryUpcomingMoviesShrinkRequest {
	s.CityCode = &v
	return s
}

func (s *QueryUpcomingMoviesShrinkRequest) SetExtJsonShrink(v string) *QueryUpcomingMoviesShrinkRequest {
	s.ExtJsonShrink = &v
	return s
}

type QueryUpcomingMoviesResponseBody struct {
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                                `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                                `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                                `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Movies     *QueryUpcomingMoviesResponseBodyMovies `json:"Movies,omitempty" xml:"Movies,omitempty" type:"Struct"`
}

func (s QueryUpcomingMoviesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUpcomingMoviesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUpcomingMoviesResponseBody) SetRequestId(v string) *QueryUpcomingMoviesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBody) SetSuccess(v bool) *QueryUpcomingMoviesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBody) SetSubMessage(v string) *QueryUpcomingMoviesResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBody) SetCode(v string) *QueryUpcomingMoviesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBody) SetMessage(v string) *QueryUpcomingMoviesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBody) SetSubCode(v string) *QueryUpcomingMoviesResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBody) SetTotalCount(v int64) *QueryUpcomingMoviesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBody) SetLogsId(v string) *QueryUpcomingMoviesResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBody) SetMovies(v *QueryUpcomingMoviesResponseBodyMovies) *QueryUpcomingMoviesResponseBody {
	s.Movies = v
	return s
}

type QueryUpcomingMoviesResponseBodyMovies struct {
	Movie []*QueryUpcomingMoviesResponseBodyMoviesMovie `json:"Movie,omitempty" xml:"Movie,omitempty" type:"Repeated"`
}

func (s QueryUpcomingMoviesResponseBodyMovies) String() string {
	return tea.Prettify(s)
}

func (s QueryUpcomingMoviesResponseBodyMovies) GoString() string {
	return s.String()
}

func (s *QueryUpcomingMoviesResponseBodyMovies) SetMovie(v []*QueryUpcomingMoviesResponseBodyMoviesMovie) *QueryUpcomingMoviesResponseBodyMovies {
	s.Movie = v
	return s
}

type QueryUpcomingMoviesResponseBodyMoviesMovie struct {
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	MovieVersion      *string `json:"MovieVersion,omitempty" xml:"MovieVersion,omitempty"`
	BackgroundPicture *string `json:"BackgroundPicture,omitempty" xml:"BackgroundPicture,omitempty"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	OpenDay           *string `json:"OpenDay,omitempty" xml:"OpenDay,omitempty"`
	Highlight         *string `json:"Highlight,omitempty" xml:"Highlight,omitempty"`
	MovieTypeList     *string `json:"MovieTypeList,omitempty" xml:"MovieTypeList,omitempty"`
	Language          *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Director          *string `json:"Director,omitempty" xml:"Director,omitempty"`
	OpenTime          *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	Poster            *string `json:"Poster,omitempty" xml:"Poster,omitempty"`
	MovieName         *string `json:"MovieName,omitempty" xml:"MovieName,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Country           *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Duration          *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	MovieNameEn       *string `json:"MovieNameEn,omitempty" xml:"MovieNameEn,omitempty"`
	LeadingRole       *string `json:"LeadingRole,omitempty" xml:"LeadingRole,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TrailerList       *string `json:"TrailerList,omitempty" xml:"TrailerList,omitempty"`
}

func (s QueryUpcomingMoviesResponseBodyMoviesMovie) String() string {
	return tea.Prettify(s)
}

func (s QueryUpcomingMoviesResponseBodyMoviesMovie) GoString() string {
	return s.String()
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetType(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.Type = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetMovieVersion(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.MovieVersion = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetBackgroundPicture(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.BackgroundPicture = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetRemark(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.Remark = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetOpenDay(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.OpenDay = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetHighlight(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.Highlight = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetMovieTypeList(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.MovieTypeList = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetLanguage(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.Language = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetDirector(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.Director = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetOpenTime(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.OpenTime = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetPoster(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.Poster = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetMovieName(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.MovieName = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetDescription(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.Description = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetCountry(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.Country = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetDuration(v int64) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.Duration = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetMovieNameEn(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.MovieNameEn = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetLeadingRole(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.LeadingRole = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetId(v int64) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.Id = &v
	return s
}

func (s *QueryUpcomingMoviesResponseBodyMoviesMovie) SetTrailerList(v string) *QueryUpcomingMoviesResponseBodyMoviesMovie {
	s.TrailerList = &v
	return s
}

type QueryUpcomingMoviesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUpcomingMoviesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUpcomingMoviesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUpcomingMoviesResponse) GoString() string {
	return s.String()
}

func (s *QueryUpcomingMoviesResponse) SetHeaders(v map[string]*string) *QueryUpcomingMoviesResponse {
	s.Headers = v
	return s
}

func (s *QueryUpcomingMoviesResponse) SetBody(v *QueryUpcomingMoviesResponseBody) *QueryUpcomingMoviesResponse {
	s.Body = v
	return s
}

type QueryWithholdTradeRequest struct {
	TradeNo    *string `json:"TradeNo,omitempty" xml:"TradeNo,omitempty"`
	OutTradeNo *string `json:"OutTradeNo,omitempty" xml:"OutTradeNo,omitempty"`
	MerchantId *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
}

func (s QueryWithholdTradeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWithholdTradeRequest) GoString() string {
	return s.String()
}

func (s *QueryWithholdTradeRequest) SetTradeNo(v string) *QueryWithholdTradeRequest {
	s.TradeNo = &v
	return s
}

func (s *QueryWithholdTradeRequest) SetOutTradeNo(v string) *QueryWithholdTradeRequest {
	s.OutTradeNo = &v
	return s
}

func (s *QueryWithholdTradeRequest) SetMerchantId(v string) *QueryWithholdTradeRequest {
	s.MerchantId = &v
	return s
}

type QueryWithholdTradeResponseBody struct {
	Code                       *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                    *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId                  *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	QueryWithholdTradeResponse *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse `json:"QueryWithholdTradeResponse,omitempty" xml:"QueryWithholdTradeResponse,omitempty" type:"Struct"`
}

func (s QueryWithholdTradeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryWithholdTradeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWithholdTradeResponseBody) SetCode(v string) *QueryWithholdTradeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryWithholdTradeResponseBody) SetMessage(v string) *QueryWithholdTradeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryWithholdTradeResponseBody) SetRequestId(v string) *QueryWithholdTradeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWithholdTradeResponseBody) SetQueryWithholdTradeResponse(v *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse) *QueryWithholdTradeResponseBody {
	s.QueryWithholdTradeResponse = v
	return s
}

type QueryWithholdTradeResponseBodyQueryWithholdTradeResponse struct {
	SettleStatus *string `json:"SettleStatus,omitempty" xml:"SettleStatus,omitempty"`
	TradeStatus  *string `json:"TradeStatus,omitempty" xml:"TradeStatus,omitempty"`
	TotalAmount  *string `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	TradeNo      *string `json:"TradeNo,omitempty" xml:"TradeNo,omitempty"`
	PaymentDate  *string `json:"PaymentDate,omitempty" xml:"PaymentDate,omitempty"`
	OutTradeNo   *string `json:"OutTradeNo,omitempty" xml:"OutTradeNo,omitempty"`
}

func (s QueryWithholdTradeResponseBodyQueryWithholdTradeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWithholdTradeResponseBodyQueryWithholdTradeResponse) GoString() string {
	return s.String()
}

func (s *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse) SetSettleStatus(v string) *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse {
	s.SettleStatus = &v
	return s
}

func (s *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse) SetTradeStatus(v string) *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse {
	s.TradeStatus = &v
	return s
}

func (s *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse) SetTotalAmount(v string) *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse {
	s.TotalAmount = &v
	return s
}

func (s *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse) SetTradeNo(v string) *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse {
	s.TradeNo = &v
	return s
}

func (s *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse) SetPaymentDate(v string) *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse {
	s.PaymentDate = &v
	return s
}

func (s *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse) SetOutTradeNo(v string) *QueryWithholdTradeResponseBodyQueryWithholdTradeResponse {
	s.OutTradeNo = &v
	return s
}

type QueryWithholdTradeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryWithholdTradeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryWithholdTradeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWithholdTradeResponse) GoString() string {
	return s.String()
}

func (s *QueryWithholdTradeResponse) SetHeaders(v map[string]*string) *QueryWithholdTradeResponse {
	s.Headers = v
	return s
}

func (s *QueryWithholdTradeResponse) SetBody(v *QueryWithholdTradeResponseBody) *QueryWithholdTradeResponse {
	s.Body = v
	return s
}

type RefundOrderRequest struct {
	OutRequestNo            *string `json:"OutRequestNo,omitempty" xml:"OutRequestNo,omitempty"`
	OutTradeNo              *string `json:"OutTradeNo,omitempty" xml:"OutTradeNo,omitempty"`
	TradeNo                 *string `json:"TradeNo,omitempty" xml:"TradeNo,omitempty"`
	RefundReason            *string `json:"RefundReason,omitempty" xml:"RefundReason,omitempty"`
	RefundAmount            *string `json:"RefundAmount,omitempty" xml:"RefundAmount,omitempty"`
	RefundRoyaltyParameters *string `json:"RefundRoyaltyParameters,omitempty" xml:"RefundRoyaltyParameters,omitempty"`
	ExtInfo                 *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	MerchantId              *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
}

func (s RefundOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundOrderRequest) GoString() string {
	return s.String()
}

func (s *RefundOrderRequest) SetOutRequestNo(v string) *RefundOrderRequest {
	s.OutRequestNo = &v
	return s
}

func (s *RefundOrderRequest) SetOutTradeNo(v string) *RefundOrderRequest {
	s.OutTradeNo = &v
	return s
}

func (s *RefundOrderRequest) SetTradeNo(v string) *RefundOrderRequest {
	s.TradeNo = &v
	return s
}

func (s *RefundOrderRequest) SetRefundReason(v string) *RefundOrderRequest {
	s.RefundReason = &v
	return s
}

func (s *RefundOrderRequest) SetRefundAmount(v string) *RefundOrderRequest {
	s.RefundAmount = &v
	return s
}

func (s *RefundOrderRequest) SetRefundRoyaltyParameters(v string) *RefundOrderRequest {
	s.RefundRoyaltyParameters = &v
	return s
}

func (s *RefundOrderRequest) SetExtInfo(v string) *RefundOrderRequest {
	s.ExtInfo = &v
	return s
}

func (s *RefundOrderRequest) SetMerchantId(v string) *RefundOrderRequest {
	s.MerchantId = &v
	return s
}

type RefundOrderResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RefundResponse *RefundOrderResponseBodyRefundResponse `json:"RefundResponse,omitempty" xml:"RefundResponse,omitempty" type:"Struct"`
}

func (s RefundOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RefundOrderResponseBody) SetCode(v string) *RefundOrderResponseBody {
	s.Code = &v
	return s
}

func (s *RefundOrderResponseBody) SetMessage(v string) *RefundOrderResponseBody {
	s.Message = &v
	return s
}

func (s *RefundOrderResponseBody) SetRequestId(v string) *RefundOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundOrderResponseBody) SetRefundResponse(v *RefundOrderResponseBodyRefundResponse) *RefundOrderResponseBody {
	s.RefundResponse = v
	return s
}

type RefundOrderResponseBodyRefundResponse struct {
	FundChange   *string `json:"FundChange,omitempty" xml:"FundChange,omitempty"`
	GmtRefundPay *string `json:"GmtRefundPay,omitempty" xml:"GmtRefundPay,omitempty"`
	TradeNo      *string `json:"TradeNo,omitempty" xml:"TradeNo,omitempty"`
	OutTradeNo   *string `json:"OutTradeNo,omitempty" xml:"OutTradeNo,omitempty"`
	OutRequestNo *string `json:"OutRequestNo,omitempty" xml:"OutRequestNo,omitempty"`
}

func (s RefundOrderResponseBodyRefundResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundOrderResponseBodyRefundResponse) GoString() string {
	return s.String()
}

func (s *RefundOrderResponseBodyRefundResponse) SetFundChange(v string) *RefundOrderResponseBodyRefundResponse {
	s.FundChange = &v
	return s
}

func (s *RefundOrderResponseBodyRefundResponse) SetGmtRefundPay(v string) *RefundOrderResponseBodyRefundResponse {
	s.GmtRefundPay = &v
	return s
}

func (s *RefundOrderResponseBodyRefundResponse) SetTradeNo(v string) *RefundOrderResponseBodyRefundResponse {
	s.TradeNo = &v
	return s
}

func (s *RefundOrderResponseBodyRefundResponse) SetOutTradeNo(v string) *RefundOrderResponseBodyRefundResponse {
	s.OutTradeNo = &v
	return s
}

func (s *RefundOrderResponseBodyRefundResponse) SetOutRequestNo(v string) *RefundOrderResponseBodyRefundResponse {
	s.OutRequestNo = &v
	return s
}

type RefundOrderResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefundOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefundOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundOrderResponse) GoString() string {
	return s.String()
}

func (s *RefundOrderResponse) SetHeaders(v map[string]*string) *RefundOrderResponse {
	s.Headers = v
	return s
}

func (s *RefundOrderResponse) SetBody(v *RefundOrderResponseBody) *RefundOrderResponse {
	s.Body = v
	return s
}

type RefundPointRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SellerId              *string `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	LmOrderId             *string `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	Reason                *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
}

func (s RefundPointRequest) String() string {
	return tea.Prettify(s)
}

func (s RefundPointRequest) GoString() string {
	return s.String()
}

func (s *RefundPointRequest) SetBizId(v string) *RefundPointRequest {
	s.BizId = &v
	return s
}

func (s *RefundPointRequest) SetSellerId(v string) *RefundPointRequest {
	s.SellerId = &v
	return s
}

func (s *RefundPointRequest) SetLmOrderId(v string) *RefundPointRequest {
	s.LmOrderId = &v
	return s
}

func (s *RefundPointRequest) SetReason(v string) *RefundPointRequest {
	s.Reason = &v
	return s
}

func (s *RefundPointRequest) SetUseAnonymousTbAccount(v bool) *RefundPointRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *RefundPointRequest) SetThirdPartyUserId(v string) *RefundPointRequest {
	s.ThirdPartyUserId = &v
	return s
}

type RefundPointResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefundPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefundPointResponseBody) GoString() string {
	return s.String()
}

func (s *RefundPointResponseBody) SetCode(v string) *RefundPointResponseBody {
	s.Code = &v
	return s
}

func (s *RefundPointResponseBody) SetMessage(v string) *RefundPointResponseBody {
	s.Message = &v
	return s
}

func (s *RefundPointResponseBody) SetRequestId(v string) *RefundPointResponseBody {
	s.RequestId = &v
	return s
}

type RefundPointResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefundPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefundPointResponse) String() string {
	return tea.Prettify(s)
}

func (s RefundPointResponse) GoString() string {
	return s.String()
}

func (s *RefundPointResponse) SetHeaders(v map[string]*string) *RefundPointResponse {
	s.Headers = v
	return s
}

func (s *RefundPointResponse) SetBody(v *RefundPointResponseBody) *RefundPointResponse {
	s.Body = v
	return s
}

type RefuseMerchantSyncTaskRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	SellerNick *string `json:"SellerNick,omitempty" xml:"SellerNick,omitempty"`
	TimeStamp  *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s RefuseMerchantSyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s RefuseMerchantSyncTaskRequest) GoString() string {
	return s.String()
}

func (s *RefuseMerchantSyncTaskRequest) SetBizId(v string) *RefuseMerchantSyncTaskRequest {
	s.BizId = &v
	return s
}

func (s *RefuseMerchantSyncTaskRequest) SetTaskId(v string) *RefuseMerchantSyncTaskRequest {
	s.TaskId = &v
	return s
}

func (s *RefuseMerchantSyncTaskRequest) SetSellerNick(v string) *RefuseMerchantSyncTaskRequest {
	s.SellerNick = &v
	return s
}

func (s *RefuseMerchantSyncTaskRequest) SetTimeStamp(v int64) *RefuseMerchantSyncTaskRequest {
	s.TimeStamp = &v
	return s
}

type RefuseMerchantSyncTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefuseMerchantSyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefuseMerchantSyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RefuseMerchantSyncTaskResponseBody) SetCode(v string) *RefuseMerchantSyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *RefuseMerchantSyncTaskResponseBody) SetMessage(v string) *RefuseMerchantSyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RefuseMerchantSyncTaskResponseBody) SetRequestId(v string) *RefuseMerchantSyncTaskResponseBody {
	s.RequestId = &v
	return s
}

type RefuseMerchantSyncTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefuseMerchantSyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefuseMerchantSyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s RefuseMerchantSyncTaskResponse) GoString() string {
	return s.String()
}

func (s *RefuseMerchantSyncTaskResponse) SetHeaders(v map[string]*string) *RefuseMerchantSyncTaskResponse {
	s.Headers = v
	return s
}

func (s *RefuseMerchantSyncTaskResponse) SetBody(v *RefuseMerchantSyncTaskResponseBody) *RefuseMerchantSyncTaskResponse {
	s.Body = v
	return s
}

type RegistAnonymousTbAccountRequest struct {
	BizId            *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ThirdPartyUserId *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
}

func (s RegistAnonymousTbAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s RegistAnonymousTbAccountRequest) GoString() string {
	return s.String()
}

func (s *RegistAnonymousTbAccountRequest) SetBizId(v string) *RegistAnonymousTbAccountRequest {
	s.BizId = &v
	return s
}

func (s *RegistAnonymousTbAccountRequest) SetThirdPartyUserId(v string) *RegistAnonymousTbAccountRequest {
	s.ThirdPartyUserId = &v
	return s
}

type RegistAnonymousTbAccountResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegistAnonymousTbAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegistAnonymousTbAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RegistAnonymousTbAccountResponseBody) SetCode(v string) *RegistAnonymousTbAccountResponseBody {
	s.Code = &v
	return s
}

func (s *RegistAnonymousTbAccountResponseBody) SetMessage(v string) *RegistAnonymousTbAccountResponseBody {
	s.Message = &v
	return s
}

func (s *RegistAnonymousTbAccountResponseBody) SetRequestId(v string) *RegistAnonymousTbAccountResponseBody {
	s.RequestId = &v
	return s
}

type RegistAnonymousTbAccountResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegistAnonymousTbAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegistAnonymousTbAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s RegistAnonymousTbAccountResponse) GoString() string {
	return s.String()
}

func (s *RegistAnonymousTbAccountResponse) SetHeaders(v map[string]*string) *RegistAnonymousTbAccountResponse {
	s.Headers = v
	return s
}

func (s *RegistAnonymousTbAccountResponse) SetBody(v *RegistAnonymousTbAccountResponseBody) *RegistAnonymousTbAccountResponse {
	s.Body = v
	return s
}

type ReleaseMovieSeatRequest struct {
	BizId            *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LockSeatApplyKey *string `json:"LockSeatApplyKey,omitempty" xml:"LockSeatApplyKey,omitempty"`
	BizUid           *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	ExtJson          *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s ReleaseMovieSeatRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseMovieSeatRequest) GoString() string {
	return s.String()
}

func (s *ReleaseMovieSeatRequest) SetBizId(v string) *ReleaseMovieSeatRequest {
	s.BizId = &v
	return s
}

func (s *ReleaseMovieSeatRequest) SetLockSeatApplyKey(v string) *ReleaseMovieSeatRequest {
	s.LockSeatApplyKey = &v
	return s
}

func (s *ReleaseMovieSeatRequest) SetBizUid(v string) *ReleaseMovieSeatRequest {
	s.BizUid = &v
	return s
}

func (s *ReleaseMovieSeatRequest) SetExtJson(v string) *ReleaseMovieSeatRequest {
	s.ExtJson = &v
	return s
}

type ReleaseMovieSeatResponseBody struct {
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage   *string                                   `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code         *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode      *string                                   `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	LogsId       *string                                   `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	ActionResult *ReleaseMovieSeatResponseBodyActionResult `json:"ActionResult,omitempty" xml:"ActionResult,omitempty" type:"Struct"`
}

func (s ReleaseMovieSeatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseMovieSeatResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseMovieSeatResponseBody) SetRequestId(v string) *ReleaseMovieSeatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseMovieSeatResponseBody) SetSuccess(v bool) *ReleaseMovieSeatResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseMovieSeatResponseBody) SetSubMessage(v string) *ReleaseMovieSeatResponseBody {
	s.SubMessage = &v
	return s
}

func (s *ReleaseMovieSeatResponseBody) SetCode(v string) *ReleaseMovieSeatResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseMovieSeatResponseBody) SetMessage(v string) *ReleaseMovieSeatResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseMovieSeatResponseBody) SetSubCode(v string) *ReleaseMovieSeatResponseBody {
	s.SubCode = &v
	return s
}

func (s *ReleaseMovieSeatResponseBody) SetLogsId(v string) *ReleaseMovieSeatResponseBody {
	s.LogsId = &v
	return s
}

func (s *ReleaseMovieSeatResponseBody) SetActionResult(v *ReleaseMovieSeatResponseBodyActionResult) *ReleaseMovieSeatResponseBody {
	s.ActionResult = v
	return s
}

type ReleaseMovieSeatResponseBodyActionResult struct {
	ReturnCode    *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	ReturnValue   *string `json:"ReturnValue,omitempty" xml:"ReturnValue,omitempty"`
	ReturnMessage *string `json:"ReturnMessage,omitempty" xml:"ReturnMessage,omitempty"`
}

func (s ReleaseMovieSeatResponseBodyActionResult) String() string {
	return tea.Prettify(s)
}

func (s ReleaseMovieSeatResponseBodyActionResult) GoString() string {
	return s.String()
}

func (s *ReleaseMovieSeatResponseBodyActionResult) SetReturnCode(v string) *ReleaseMovieSeatResponseBodyActionResult {
	s.ReturnCode = &v
	return s
}

func (s *ReleaseMovieSeatResponseBodyActionResult) SetReturnValue(v string) *ReleaseMovieSeatResponseBodyActionResult {
	s.ReturnValue = &v
	return s
}

func (s *ReleaseMovieSeatResponseBodyActionResult) SetReturnMessage(v string) *ReleaseMovieSeatResponseBodyActionResult {
	s.ReturnMessage = &v
	return s
}

type ReleaseMovieSeatResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseMovieSeatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseMovieSeatResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseMovieSeatResponse) GoString() string {
	return s.String()
}

func (s *ReleaseMovieSeatResponse) SetHeaders(v map[string]*string) *ReleaseMovieSeatResponse {
	s.Headers = v
	return s
}

func (s *ReleaseMovieSeatResponse) SetBody(v *ReleaseMovieSeatResponseBody) *ReleaseMovieSeatResponse {
	s.Body = v
	return s
}

type RemoveAddressRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	AddressInfo           *string `json:"AddressInfo,omitempty" xml:"AddressInfo,omitempty"`
}

func (s RemoveAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAddressRequest) GoString() string {
	return s.String()
}

func (s *RemoveAddressRequest) SetBizId(v string) *RemoveAddressRequest {
	s.BizId = &v
	return s
}

func (s *RemoveAddressRequest) SetThirdPartyUserId(v string) *RemoveAddressRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *RemoveAddressRequest) SetUseAnonymousTbAccount(v bool) *RemoveAddressRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *RemoveAddressRequest) SetAddressInfo(v string) *RemoveAddressRequest {
	s.AddressInfo = &v
	return s
}

type RemoveAddressResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAddressResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAddressResponseBody) SetCode(v string) *RemoveAddressResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveAddressResponseBody) SetMessage(v string) *RemoveAddressResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveAddressResponseBody) SetRequestId(v string) *RemoveAddressResponseBody {
	s.RequestId = &v
	return s
}

type RemoveAddressResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAddressResponse) GoString() string {
	return s.String()
}

func (s *RemoveAddressResponse) SetHeaders(v map[string]*string) *RemoveAddressResponse {
	s.Headers = v
	return s
}

func (s *RemoveAddressResponse) SetBody(v *RemoveAddressResponseBody) *RemoveAddressResponse {
	s.Body = v
	return s
}

type RemoveMessagesRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	MessageIds *string `json:"MessageIds,omitempty" xml:"MessageIds,omitempty"`
	ExtJson    *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s RemoveMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveMessagesRequest) GoString() string {
	return s.String()
}

func (s *RemoveMessagesRequest) SetBizId(v string) *RemoveMessagesRequest {
	s.BizId = &v
	return s
}

func (s *RemoveMessagesRequest) SetMessageIds(v string) *RemoveMessagesRequest {
	s.MessageIds = &v
	return s
}

func (s *RemoveMessagesRequest) SetExtJson(v string) *RemoveMessagesRequest {
	s.ExtJson = &v
	return s
}

type RemoveMessagesResponseBody struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage   *string                                 `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	SubCode      *string                                 `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	LogsId       *string                                 `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	ActionResult *RemoveMessagesResponseBodyActionResult `json:"ActionResult,omitempty" xml:"ActionResult,omitempty" type:"Struct"`
}

func (s RemoveMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMessagesResponseBody) SetRequestId(v string) *RemoveMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveMessagesResponseBody) SetSuccess(v bool) *RemoveMessagesResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveMessagesResponseBody) SetSubMessage(v string) *RemoveMessagesResponseBody {
	s.SubMessage = &v
	return s
}

func (s *RemoveMessagesResponseBody) SetCode(v string) *RemoveMessagesResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveMessagesResponseBody) SetSubCode(v string) *RemoveMessagesResponseBody {
	s.SubCode = &v
	return s
}

func (s *RemoveMessagesResponseBody) SetMessage(v string) *RemoveMessagesResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveMessagesResponseBody) SetLogsId(v string) *RemoveMessagesResponseBody {
	s.LogsId = &v
	return s
}

func (s *RemoveMessagesResponseBody) SetActionResult(v *RemoveMessagesResponseBodyActionResult) *RemoveMessagesResponseBody {
	s.ActionResult = v
	return s
}

type RemoveMessagesResponseBodyActionResult struct {
	ReturnCode    *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	ReturnValue   *string `json:"ReturnValue,omitempty" xml:"ReturnValue,omitempty"`
	ReturnMessage *string `json:"ReturnMessage,omitempty" xml:"ReturnMessage,omitempty"`
}

func (s RemoveMessagesResponseBodyActionResult) String() string {
	return tea.Prettify(s)
}

func (s RemoveMessagesResponseBodyActionResult) GoString() string {
	return s.String()
}

func (s *RemoveMessagesResponseBodyActionResult) SetReturnCode(v string) *RemoveMessagesResponseBodyActionResult {
	s.ReturnCode = &v
	return s
}

func (s *RemoveMessagesResponseBodyActionResult) SetReturnValue(v string) *RemoveMessagesResponseBodyActionResult {
	s.ReturnValue = &v
	return s
}

func (s *RemoveMessagesResponseBodyActionResult) SetReturnMessage(v string) *RemoveMessagesResponseBodyActionResult {
	s.ReturnMessage = &v
	return s
}

type RemoveMessagesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveMessagesResponse) GoString() string {
	return s.String()
}

func (s *RemoveMessagesResponse) SetHeaders(v map[string]*string) *RemoveMessagesResponse {
	s.Headers = v
	return s
}

func (s *RemoveMessagesResponse) SetBody(v *RemoveMessagesResponseBody) *RemoveMessagesResponse {
	s.Body = v
	return s
}

type RenderH5OrderRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	BuyOrderRequestModel  *string `json:"BuyOrderRequestModel,omitempty" xml:"BuyOrderRequestModel,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s RenderH5OrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RenderH5OrderRequest) GoString() string {
	return s.String()
}

func (s *RenderH5OrderRequest) SetBizId(v string) *RenderH5OrderRequest {
	s.BizId = &v
	return s
}

func (s *RenderH5OrderRequest) SetBizUid(v string) *RenderH5OrderRequest {
	s.BizUid = &v
	return s
}

func (s *RenderH5OrderRequest) SetUseAnonymousTbAccount(v bool) *RenderH5OrderRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *RenderH5OrderRequest) SetThirdPartyUserId(v string) *RenderH5OrderRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *RenderH5OrderRequest) SetBuyOrderRequestModel(v string) *RenderH5OrderRequest {
	s.BuyOrderRequestModel = &v
	return s
}

func (s *RenderH5OrderRequest) SetAccountType(v string) *RenderH5OrderRequest {
	s.AccountType = &v
	return s
}

type RenderH5OrderResponseBody struct {
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                         `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                         `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                         `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Model      *RenderH5OrderResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s RenderH5OrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenderH5OrderResponseBody) GoString() string {
	return s.String()
}

func (s *RenderH5OrderResponseBody) SetRequestId(v string) *RenderH5OrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenderH5OrderResponseBody) SetSuccess(v bool) *RenderH5OrderResponseBody {
	s.Success = &v
	return s
}

func (s *RenderH5OrderResponseBody) SetSubMessage(v string) *RenderH5OrderResponseBody {
	s.SubMessage = &v
	return s
}

func (s *RenderH5OrderResponseBody) SetCode(v string) *RenderH5OrderResponseBody {
	s.Code = &v
	return s
}

func (s *RenderH5OrderResponseBody) SetMessage(v string) *RenderH5OrderResponseBody {
	s.Message = &v
	return s
}

func (s *RenderH5OrderResponseBody) SetSubCode(v string) *RenderH5OrderResponseBody {
	s.SubCode = &v
	return s
}

func (s *RenderH5OrderResponseBody) SetTotalCount(v int64) *RenderH5OrderResponseBody {
	s.TotalCount = &v
	return s
}

func (s *RenderH5OrderResponseBody) SetLogsId(v string) *RenderH5OrderResponseBody {
	s.LogsId = &v
	return s
}

func (s *RenderH5OrderResponseBody) SetModel(v *RenderH5OrderResponseBodyModel) *RenderH5OrderResponseBody {
	s.Model = v
	return s
}

type RenderH5OrderResponseBodyModel struct {
	ExtInfo            map[string]interface{}                            `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	BuyerCurrentPoints *int64                                            `json:"BuyerCurrentPoints,omitempty" xml:"BuyerCurrentPoints,omitempty"`
	LmItemInfoList     []*RenderH5OrderResponseBodyModelLmItemInfoList   `json:"LmItemInfoList,omitempty" xml:"LmItemInfoList,omitempty" type:"Repeated"`
	DeliveryInfoList   []*RenderH5OrderResponseBodyModelDeliveryInfoList `json:"DeliveryInfoList,omitempty" xml:"DeliveryInfoList,omitempty" type:"Repeated"`
	AddressInfoList    []*RenderH5OrderResponseBodyModelAddressInfoList  `json:"AddressInfoList,omitempty" xml:"AddressInfoList,omitempty" type:"Repeated"`
	InvoiceInfo        *RenderH5OrderResponseBodyModelInvoiceInfo        `json:"InvoiceInfo,omitempty" xml:"InvoiceInfo,omitempty" type:"Struct"`
}

func (s RenderH5OrderResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s RenderH5OrderResponseBodyModel) GoString() string {
	return s.String()
}

func (s *RenderH5OrderResponseBodyModel) SetExtInfo(v map[string]interface{}) *RenderH5OrderResponseBodyModel {
	s.ExtInfo = v
	return s
}

func (s *RenderH5OrderResponseBodyModel) SetBuyerCurrentPoints(v int64) *RenderH5OrderResponseBodyModel {
	s.BuyerCurrentPoints = &v
	return s
}

func (s *RenderH5OrderResponseBodyModel) SetLmItemInfoList(v []*RenderH5OrderResponseBodyModelLmItemInfoList) *RenderH5OrderResponseBodyModel {
	s.LmItemInfoList = v
	return s
}

func (s *RenderH5OrderResponseBodyModel) SetDeliveryInfoList(v []*RenderH5OrderResponseBodyModelDeliveryInfoList) *RenderH5OrderResponseBodyModel {
	s.DeliveryInfoList = v
	return s
}

func (s *RenderH5OrderResponseBodyModel) SetAddressInfoList(v []*RenderH5OrderResponseBodyModelAddressInfoList) *RenderH5OrderResponseBodyModel {
	s.AddressInfoList = v
	return s
}

func (s *RenderH5OrderResponseBodyModel) SetInvoiceInfo(v *RenderH5OrderResponseBodyModelInvoiceInfo) *RenderH5OrderResponseBodyModel {
	s.InvoiceInfo = v
	return s
}

type RenderH5OrderResponseBodyModelLmItemInfoList struct {
	TbShopName      *string                `json:"TbShopName,omitempty" xml:"TbShopName,omitempty"`
	SellerId        *int64                 `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	LmItemId        *string                `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	Message         *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	CanSell         *bool                  `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	Cash            *int64                 `json:"Cash,omitempty" xml:"Cash,omitempty"`
	ItemId          *int64                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	VirtualItemType *string                `json:"VirtualItemType,omitempty" xml:"VirtualItemType,omitempty"`
	ItemName        *string                `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	ActualPrice     *int64                 `json:"ActualPrice,omitempty" xml:"ActualPrice,omitempty"`
	SkuName         *string                `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
	SkuId           *int64                 `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Points          *int64                 `json:"Points,omitempty" xml:"Points,omitempty"`
	ItemUrl         *string                `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	SellerNick      *string                `json:"SellerNick,omitempty" xml:"SellerNick,omitempty"`
	Quantity        *int32                 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	Features        map[string]interface{} `json:"Features,omitempty" xml:"Features,omitempty"`
	ItemPicUrl      *string                `json:"ItemPicUrl,omitempty" xml:"ItemPicUrl,omitempty"`
}

func (s RenderH5OrderResponseBodyModelLmItemInfoList) String() string {
	return tea.Prettify(s)
}

func (s RenderH5OrderResponseBodyModelLmItemInfoList) GoString() string {
	return s.String()
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetTbShopName(v string) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.TbShopName = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetSellerId(v int64) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.SellerId = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetLmItemId(v string) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.LmItemId = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetMessage(v string) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.Message = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetCanSell(v bool) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.CanSell = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetCash(v int64) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.Cash = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetItemId(v int64) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.ItemId = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetVirtualItemType(v string) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.VirtualItemType = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetItemName(v string) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.ItemName = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetActualPrice(v int64) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.ActualPrice = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetSkuName(v string) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.SkuName = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetSkuId(v int64) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.SkuId = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetPoints(v int64) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.Points = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetItemUrl(v string) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.ItemUrl = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetSellerNick(v string) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.SellerNick = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetQuantity(v int32) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.Quantity = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetFeatures(v map[string]interface{}) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.Features = v
	return s
}

func (s *RenderH5OrderResponseBodyModelLmItemInfoList) SetItemPicUrl(v string) *RenderH5OrderResponseBodyModelLmItemInfoList {
	s.ItemPicUrl = &v
	return s
}

type RenderH5OrderResponseBodyModelDeliveryInfoList struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	PostFee     *int64  `json:"PostFee,omitempty" xml:"PostFee,omitempty"`
	ServiceType *int64  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s RenderH5OrderResponseBodyModelDeliveryInfoList) String() string {
	return tea.Prettify(s)
}

func (s RenderH5OrderResponseBodyModelDeliveryInfoList) GoString() string {
	return s.String()
}

func (s *RenderH5OrderResponseBodyModelDeliveryInfoList) SetDisplayName(v string) *RenderH5OrderResponseBodyModelDeliveryInfoList {
	s.DisplayName = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelDeliveryInfoList) SetPostFee(v int64) *RenderH5OrderResponseBodyModelDeliveryInfoList {
	s.PostFee = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelDeliveryInfoList) SetServiceType(v int64) *RenderH5OrderResponseBodyModelDeliveryInfoList {
	s.ServiceType = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelDeliveryInfoList) SetId(v string) *RenderH5OrderResponseBodyModelDeliveryInfoList {
	s.Id = &v
	return s
}

type RenderH5OrderResponseBodyModelAddressInfoList struct {
	DivisionCode  *string `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
	Receiver      *string `json:"Receiver,omitempty" xml:"Receiver,omitempty"`
	AddressDetail *string `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	Default       *bool   `json:"Default,omitempty" xml:"Default,omitempty"`
	AddressId     *int64  `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	ReceiverPhone *string `json:"ReceiverPhone,omitempty" xml:"ReceiverPhone,omitempty"`
}

func (s RenderH5OrderResponseBodyModelAddressInfoList) String() string {
	return tea.Prettify(s)
}

func (s RenderH5OrderResponseBodyModelAddressInfoList) GoString() string {
	return s.String()
}

func (s *RenderH5OrderResponseBodyModelAddressInfoList) SetDivisionCode(v string) *RenderH5OrderResponseBodyModelAddressInfoList {
	s.DivisionCode = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelAddressInfoList) SetReceiver(v string) *RenderH5OrderResponseBodyModelAddressInfoList {
	s.Receiver = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelAddressInfoList) SetAddressDetail(v string) *RenderH5OrderResponseBodyModelAddressInfoList {
	s.AddressDetail = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelAddressInfoList) SetDefault(v bool) *RenderH5OrderResponseBodyModelAddressInfoList {
	s.Default = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelAddressInfoList) SetAddressId(v int64) *RenderH5OrderResponseBodyModelAddressInfoList {
	s.AddressId = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelAddressInfoList) SetReceiverPhone(v string) *RenderH5OrderResponseBodyModelAddressInfoList {
	s.ReceiverPhone = &v
	return s
}

type RenderH5OrderResponseBodyModelInvoiceInfo struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s RenderH5OrderResponseBodyModelInvoiceInfo) String() string {
	return tea.Prettify(s)
}

func (s RenderH5OrderResponseBodyModelInvoiceInfo) GoString() string {
	return s.String()
}

func (s *RenderH5OrderResponseBodyModelInvoiceInfo) SetType(v string) *RenderH5OrderResponseBodyModelInvoiceInfo {
	s.Type = &v
	return s
}

func (s *RenderH5OrderResponseBodyModelInvoiceInfo) SetDesc(v string) *RenderH5OrderResponseBodyModelInvoiceInfo {
	s.Desc = &v
	return s
}

type RenderH5OrderResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenderH5OrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenderH5OrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RenderH5OrderResponse) GoString() string {
	return s.String()
}

func (s *RenderH5OrderResponse) SetHeaders(v map[string]*string) *RenderH5OrderResponse {
	s.Headers = v
	return s
}

func (s *RenderH5OrderResponse) SetBody(v *RenderH5OrderResponseBody) *RenderH5OrderResponse {
	s.Body = v
	return s
}

type RenderOrderRequest struct {
	BizId                 *string                       `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string                       `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	ExtJson               *string                       `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	ItemList              []*RenderOrderRequestItemList `json:"ItemList,omitempty" xml:"ItemList,omitempty" type:"Repeated"`
	DeliveryAddress       *string                       `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	UseAnonymousTbAccount *bool                         `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string                       `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	LmItemId              *string                       `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	AccountType           *string                       `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s RenderOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RenderOrderRequest) GoString() string {
	return s.String()
}

func (s *RenderOrderRequest) SetBizId(v string) *RenderOrderRequest {
	s.BizId = &v
	return s
}

func (s *RenderOrderRequest) SetBizUid(v string) *RenderOrderRequest {
	s.BizUid = &v
	return s
}

func (s *RenderOrderRequest) SetExtJson(v string) *RenderOrderRequest {
	s.ExtJson = &v
	return s
}

func (s *RenderOrderRequest) SetItemList(v []*RenderOrderRequestItemList) *RenderOrderRequest {
	s.ItemList = v
	return s
}

func (s *RenderOrderRequest) SetDeliveryAddress(v string) *RenderOrderRequest {
	s.DeliveryAddress = &v
	return s
}

func (s *RenderOrderRequest) SetUseAnonymousTbAccount(v bool) *RenderOrderRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *RenderOrderRequest) SetThirdPartyUserId(v string) *RenderOrderRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *RenderOrderRequest) SetLmItemId(v string) *RenderOrderRequest {
	s.LmItemId = &v
	return s
}

func (s *RenderOrderRequest) SetAccountType(v string) *RenderOrderRequest {
	s.AccountType = &v
	return s
}

type RenderOrderRequestItemList struct {
	ItemId   *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	SkuId    *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Quantity *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	LmItemId *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
}

func (s RenderOrderRequestItemList) String() string {
	return tea.Prettify(s)
}

func (s RenderOrderRequestItemList) GoString() string {
	return s.String()
}

func (s *RenderOrderRequestItemList) SetItemId(v int64) *RenderOrderRequestItemList {
	s.ItemId = &v
	return s
}

func (s *RenderOrderRequestItemList) SetSkuId(v int64) *RenderOrderRequestItemList {
	s.SkuId = &v
	return s
}

func (s *RenderOrderRequestItemList) SetQuantity(v int32) *RenderOrderRequestItemList {
	s.Quantity = &v
	return s
}

func (s *RenderOrderRequestItemList) SetLmItemId(v string) *RenderOrderRequestItemList {
	s.LmItemId = &v
	return s
}

type RenderOrderResponseBody struct {
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code       *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                       `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string                       `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	LogsId     *string                       `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Success    *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Model      *RenderOrderResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s RenderOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenderOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RenderOrderResponseBody) SetRequestId(v string) *RenderOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenderOrderResponseBody) SetCode(v string) *RenderOrderResponseBody {
	s.Code = &v
	return s
}

func (s *RenderOrderResponseBody) SetMessage(v string) *RenderOrderResponseBody {
	s.Message = &v
	return s
}

func (s *RenderOrderResponseBody) SetSubCode(v string) *RenderOrderResponseBody {
	s.SubCode = &v
	return s
}

func (s *RenderOrderResponseBody) SetSubMessage(v string) *RenderOrderResponseBody {
	s.SubMessage = &v
	return s
}

func (s *RenderOrderResponseBody) SetLogsId(v string) *RenderOrderResponseBody {
	s.LogsId = &v
	return s
}

func (s *RenderOrderResponseBody) SetSuccess(v bool) *RenderOrderResponseBody {
	s.Success = &v
	return s
}

func (s *RenderOrderResponseBody) SetTotalCount(v int64) *RenderOrderResponseBody {
	s.TotalCount = &v
	return s
}

func (s *RenderOrderResponseBody) SetModel(v *RenderOrderResponseBodyModel) *RenderOrderResponseBody {
	s.Model = v
	return s
}

type RenderOrderResponseBodyModel struct {
	RenderOrderInfos *RenderOrderResponseBodyModelRenderOrderInfos `json:"RenderOrderInfos,omitempty" xml:"RenderOrderInfos,omitempty" type:"Struct"`
}

func (s RenderOrderResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s RenderOrderResponseBodyModel) GoString() string {
	return s.String()
}

func (s *RenderOrderResponseBodyModel) SetRenderOrderInfos(v *RenderOrderResponseBodyModelRenderOrderInfos) *RenderOrderResponseBodyModel {
	s.RenderOrderInfos = v
	return s
}

type RenderOrderResponseBodyModelRenderOrderInfos struct {
	RenderOrderInfos []*RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfos `json:"RenderOrderInfos,omitempty" xml:"RenderOrderInfos,omitempty" type:"Repeated"`
}

func (s RenderOrderResponseBodyModelRenderOrderInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderOrderResponseBodyModelRenderOrderInfos) GoString() string {
	return s.String()
}

func (s *RenderOrderResponseBodyModelRenderOrderInfos) SetRenderOrderInfos(v []*RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfos) *RenderOrderResponseBodyModelRenderOrderInfos {
	s.RenderOrderInfos = v
	return s
}

type RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfos struct {
	ExtInfo       map[string]*string                                                         `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	LmItemInfos   *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfos   `json:"LmItemInfos,omitempty" xml:"LmItemInfos,omitempty" type:"Struct"`
	DeliveryInfos *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfos `json:"DeliveryInfos,omitempty" xml:"DeliveryInfos,omitempty" type:"Struct"`
}

func (s RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfos) GoString() string {
	return s.String()
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfos) SetExtInfo(v map[string]*string) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfos {
	s.ExtInfo = v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfos) SetLmItemInfos(v *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfos) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfos {
	s.LmItemInfos = v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfos) SetDeliveryInfos(v *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfos) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfos {
	s.DeliveryInfos = v
	return s
}

type RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfos struct {
	LmItemInfos []*RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos `json:"LmItemInfos,omitempty" xml:"LmItemInfos,omitempty" type:"Repeated"`
}

func (s RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfos) GoString() string {
	return s.String()
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfos) SetLmItemInfos(v []*RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfos {
	s.LmItemInfos = v
	return s
}

type RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos struct {
	ItemId      *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	SkuId       *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Quantity    *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	CanSell     *bool   `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ItemName    *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	SkuName     *string `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
	LmItemId    *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	SellerId    *int64  `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	SellerNick  *string `json:"SellerNick,omitempty" xml:"SellerNick,omitempty"`
	Cash        *int64  `json:"Cash,omitempty" xml:"Cash,omitempty"`
	Points      *int64  `json:"Points,omitempty" xml:"Points,omitempty"`
	ActualPrice *int64  `json:"ActualPrice,omitempty" xml:"ActualPrice,omitempty"`
	ItemPicUrl  *string `json:"ItemPicUrl,omitempty" xml:"ItemPicUrl,omitempty"`
}

func (s RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) GoString() string {
	return s.String()
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetItemId(v int64) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.ItemId = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetSkuId(v int64) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.SkuId = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetQuantity(v int32) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.Quantity = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetCanSell(v bool) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.CanSell = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetMessage(v string) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.Message = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetItemName(v string) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.ItemName = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetSkuName(v string) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.SkuName = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetLmItemId(v string) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.LmItemId = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetSellerId(v int64) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.SellerId = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetSellerNick(v string) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.SellerNick = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetCash(v int64) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.Cash = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetPoints(v int64) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.Points = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetActualPrice(v int64) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.ActualPrice = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos) SetItemPicUrl(v string) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosLmItemInfosLmItemInfos {
	s.ItemPicUrl = &v
	return s
}

type RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfos struct {
	DeliveryInfos []*RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos `json:"DeliveryInfos,omitempty" xml:"DeliveryInfos,omitempty" type:"Repeated"`
}

func (s RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfos) GoString() string {
	return s.String()
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfos) SetDeliveryInfos(v []*RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfos {
	s.DeliveryInfos = v
	return s
}

type RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	PostFee     *int64  `json:"PostFee,omitempty" xml:"PostFee,omitempty"`
	ServiceType *int64  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos) GoString() string {
	return s.String()
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos) SetId(v string) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos {
	s.Id = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos) SetDisplayName(v string) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos {
	s.DisplayName = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos) SetPostFee(v int64) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos {
	s.PostFee = &v
	return s
}

func (s *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos) SetServiceType(v int64) *RenderOrderResponseBodyModelRenderOrderInfosRenderOrderInfosDeliveryInfosDeliveryInfos {
	s.ServiceType = &v
	return s
}

type RenderOrderResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenderOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenderOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RenderOrderResponse) GoString() string {
	return s.String()
}

func (s *RenderOrderResponse) SetHeaders(v map[string]*string) *RenderOrderResponse {
	s.Headers = v
	return s
}

func (s *RenderOrderResponse) SetBody(v *RenderOrderResponseBody) *RenderOrderResponse {
	s.Body = v
	return s
}

type RepayForPayUrlRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	LmOrderId             *int64  `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
}

func (s RepayForPayUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s RepayForPayUrlRequest) GoString() string {
	return s.String()
}

func (s *RepayForPayUrlRequest) SetBizId(v string) *RepayForPayUrlRequest {
	s.BizId = &v
	return s
}

func (s *RepayForPayUrlRequest) SetBizUid(v string) *RepayForPayUrlRequest {
	s.BizUid = &v
	return s
}

func (s *RepayForPayUrlRequest) SetLmOrderId(v int64) *RepayForPayUrlRequest {
	s.LmOrderId = &v
	return s
}

func (s *RepayForPayUrlRequest) SetUseAnonymousTbAccount(v bool) *RepayForPayUrlRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *RepayForPayUrlRequest) SetThirdPartyUserId(v string) *RepayForPayUrlRequest {
	s.ThirdPartyUserId = &v
	return s
}

type RepayForPayUrlResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Model     *RepayForPayUrlResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s RepayForPayUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RepayForPayUrlResponseBody) GoString() string {
	return s.String()
}

func (s *RepayForPayUrlResponseBody) SetCode(v string) *RepayForPayUrlResponseBody {
	s.Code = &v
	return s
}

func (s *RepayForPayUrlResponseBody) SetMessage(v string) *RepayForPayUrlResponseBody {
	s.Message = &v
	return s
}

func (s *RepayForPayUrlResponseBody) SetRequestId(v string) *RepayForPayUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *RepayForPayUrlResponseBody) SetModel(v *RepayForPayUrlResponseBodyModel) *RepayForPayUrlResponseBody {
	s.Model = v
	return s
}

type RepayForPayUrlResponseBodyModel struct {
	FrontUrl *string `json:"FrontUrl,omitempty" xml:"FrontUrl,omitempty"`
}

func (s RepayForPayUrlResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s RepayForPayUrlResponseBodyModel) GoString() string {
	return s.String()
}

func (s *RepayForPayUrlResponseBodyModel) SetFrontUrl(v string) *RepayForPayUrlResponseBodyModel {
	s.FrontUrl = &v
	return s
}

type RepayForPayUrlResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RepayForPayUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RepayForPayUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s RepayForPayUrlResponse) GoString() string {
	return s.String()
}

func (s *RepayForPayUrlResponse) SetHeaders(v map[string]*string) *RepayForPayUrlResponse {
	s.Headers = v
	return s
}

func (s *RepayForPayUrlResponse) SetBody(v *RepayForPayUrlResponseBody) *RepayForPayUrlResponse {
	s.Body = v
	return s
}

type RepayOrderRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	LmOrderId             *int64  `json:"LmOrderId,omitempty" xml:"LmOrderId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s RepayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RepayOrderRequest) GoString() string {
	return s.String()
}

func (s *RepayOrderRequest) SetBizId(v string) *RepayOrderRequest {
	s.BizId = &v
	return s
}

func (s *RepayOrderRequest) SetBizUid(v string) *RepayOrderRequest {
	s.BizUid = &v
	return s
}

func (s *RepayOrderRequest) SetLmOrderId(v int64) *RepayOrderRequest {
	s.LmOrderId = &v
	return s
}

func (s *RepayOrderRequest) SetUseAnonymousTbAccount(v bool) *RepayOrderRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *RepayOrderRequest) SetThirdPartyUserId(v string) *RepayOrderRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *RepayOrderRequest) SetAccountType(v string) *RepayOrderRequest {
	s.AccountType = &v
	return s
}

type RepayOrderResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RepayOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RepayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RepayOrderResponseBody) SetCode(v string) *RepayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *RepayOrderResponseBody) SetMessage(v string) *RepayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *RepayOrderResponseBody) SetRequestId(v string) *RepayOrderResponseBody {
	s.RequestId = &v
	return s
}

type RepayOrderResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RepayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RepayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RepayOrderResponse) GoString() string {
	return s.String()
}

func (s *RepayOrderResponse) SetHeaders(v map[string]*string) *RepayOrderResponse {
	s.Headers = v
	return s
}

func (s *RepayOrderResponse) SetBody(v *RepayOrderResponseBody) *RepayOrderResponse {
	s.Body = v
	return s
}

type ReserveMovieSeatRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ScheduleId *int64  `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	SeatIds    *string `json:"SeatIds,omitempty" xml:"SeatIds,omitempty"`
	SeatNames  *string `json:"SeatNames,omitempty" xml:"SeatNames,omitempty"`
	BizUid     *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	Mobile     *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	ExtJson    *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
}

func (s ReserveMovieSeatRequest) String() string {
	return tea.Prettify(s)
}

func (s ReserveMovieSeatRequest) GoString() string {
	return s.String()
}

func (s *ReserveMovieSeatRequest) SetBizId(v string) *ReserveMovieSeatRequest {
	s.BizId = &v
	return s
}

func (s *ReserveMovieSeatRequest) SetScheduleId(v int64) *ReserveMovieSeatRequest {
	s.ScheduleId = &v
	return s
}

func (s *ReserveMovieSeatRequest) SetSeatIds(v string) *ReserveMovieSeatRequest {
	s.SeatIds = &v
	return s
}

func (s *ReserveMovieSeatRequest) SetSeatNames(v string) *ReserveMovieSeatRequest {
	s.SeatNames = &v
	return s
}

func (s *ReserveMovieSeatRequest) SetBizUid(v string) *ReserveMovieSeatRequest {
	s.BizUid = &v
	return s
}

func (s *ReserveMovieSeatRequest) SetMobile(v string) *ReserveMovieSeatRequest {
	s.Mobile = &v
	return s
}

func (s *ReserveMovieSeatRequest) SetExtJson(v string) *ReserveMovieSeatRequest {
	s.ExtJson = &v
	return s
}

type ReserveMovieSeatResponseBody struct {
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage   *string                                   `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code         *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode      *string                                   `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	LogsId       *string                                   `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	ReservedSeat *ReserveMovieSeatResponseBodyReservedSeat `json:"ReservedSeat,omitempty" xml:"ReservedSeat,omitempty" type:"Struct"`
}

func (s ReserveMovieSeatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReserveMovieSeatResponseBody) GoString() string {
	return s.String()
}

func (s *ReserveMovieSeatResponseBody) SetRequestId(v string) *ReserveMovieSeatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReserveMovieSeatResponseBody) SetSuccess(v bool) *ReserveMovieSeatResponseBody {
	s.Success = &v
	return s
}

func (s *ReserveMovieSeatResponseBody) SetSubMessage(v string) *ReserveMovieSeatResponseBody {
	s.SubMessage = &v
	return s
}

func (s *ReserveMovieSeatResponseBody) SetCode(v string) *ReserveMovieSeatResponseBody {
	s.Code = &v
	return s
}

func (s *ReserveMovieSeatResponseBody) SetMessage(v string) *ReserveMovieSeatResponseBody {
	s.Message = &v
	return s
}

func (s *ReserveMovieSeatResponseBody) SetSubCode(v string) *ReserveMovieSeatResponseBody {
	s.SubCode = &v
	return s
}

func (s *ReserveMovieSeatResponseBody) SetLogsId(v string) *ReserveMovieSeatResponseBody {
	s.LogsId = &v
	return s
}

func (s *ReserveMovieSeatResponseBody) SetReservedSeat(v *ReserveMovieSeatResponseBodyReservedSeat) *ReserveMovieSeatResponseBody {
	s.ReservedSeat = v
	return s
}

type ReserveMovieSeatResponseBodyReservedSeat struct {
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ReservedTime      *int64  `json:"ReservedTime,omitempty" xml:"ReservedTime,omitempty"`
	ApplyKey          *string `json:"ApplyKey,omitempty" xml:"ApplyKey,omitempty"`
	DefaultLockSecond *int64  `json:"DefaultLockSecond,omitempty" xml:"DefaultLockSecond,omitempty"`
}

func (s ReserveMovieSeatResponseBodyReservedSeat) String() string {
	return tea.Prettify(s)
}

func (s ReserveMovieSeatResponseBodyReservedSeat) GoString() string {
	return s.String()
}

func (s *ReserveMovieSeatResponseBodyReservedSeat) SetStatus(v string) *ReserveMovieSeatResponseBodyReservedSeat {
	s.Status = &v
	return s
}

func (s *ReserveMovieSeatResponseBodyReservedSeat) SetReservedTime(v int64) *ReserveMovieSeatResponseBodyReservedSeat {
	s.ReservedTime = &v
	return s
}

func (s *ReserveMovieSeatResponseBodyReservedSeat) SetApplyKey(v string) *ReserveMovieSeatResponseBodyReservedSeat {
	s.ApplyKey = &v
	return s
}

func (s *ReserveMovieSeatResponseBodyReservedSeat) SetDefaultLockSecond(v int64) *ReserveMovieSeatResponseBodyReservedSeat {
	s.DefaultLockSecond = &v
	return s
}

type ReserveMovieSeatResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReserveMovieSeatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReserveMovieSeatResponse) String() string {
	return tea.Prettify(s)
}

func (s ReserveMovieSeatResponse) GoString() string {
	return s.String()
}

func (s *ReserveMovieSeatResponse) SetHeaders(v map[string]*string) *ReserveMovieSeatResponse {
	s.Headers = v
	return s
}

func (s *ReserveMovieSeatResponse) SetBody(v *ReserveMovieSeatResponseBody) *ReserveMovieSeatResponse {
	s.Body = v
	return s
}

type SettleOrderRequest struct {
	OutRequestNo      *string `json:"OutRequestNo,omitempty" xml:"OutRequestNo,omitempty"`
	TradeNo           *string `json:"TradeNo,omitempty" xml:"TradeNo,omitempty"`
	RoyaltyParameters *string `json:"RoyaltyParameters,omitempty" xml:"RoyaltyParameters,omitempty"`
	ExtInfo           *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	OutTradeNo        *string `json:"OutTradeNo,omitempty" xml:"OutTradeNo,omitempty"`
	MerchantId        *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
}

func (s SettleOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SettleOrderRequest) GoString() string {
	return s.String()
}

func (s *SettleOrderRequest) SetOutRequestNo(v string) *SettleOrderRequest {
	s.OutRequestNo = &v
	return s
}

func (s *SettleOrderRequest) SetTradeNo(v string) *SettleOrderRequest {
	s.TradeNo = &v
	return s
}

func (s *SettleOrderRequest) SetRoyaltyParameters(v string) *SettleOrderRequest {
	s.RoyaltyParameters = &v
	return s
}

func (s *SettleOrderRequest) SetExtInfo(v string) *SettleOrderRequest {
	s.ExtInfo = &v
	return s
}

func (s *SettleOrderRequest) SetOutTradeNo(v string) *SettleOrderRequest {
	s.OutTradeNo = &v
	return s
}

func (s *SettleOrderRequest) SetMerchantId(v string) *SettleOrderRequest {
	s.MerchantId = &v
	return s
}

type SettleOrderResponseBody struct {
	Code                     *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                  *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId                *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TradeOrderSettleResponse *SettleOrderResponseBodyTradeOrderSettleResponse `json:"TradeOrderSettleResponse,omitempty" xml:"TradeOrderSettleResponse,omitempty" type:"Struct"`
}

func (s SettleOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SettleOrderResponseBody) GoString() string {
	return s.String()
}

func (s *SettleOrderResponseBody) SetCode(v string) *SettleOrderResponseBody {
	s.Code = &v
	return s
}

func (s *SettleOrderResponseBody) SetMessage(v string) *SettleOrderResponseBody {
	s.Message = &v
	return s
}

func (s *SettleOrderResponseBody) SetRequestId(v string) *SettleOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SettleOrderResponseBody) SetTradeOrderSettleResponse(v *SettleOrderResponseBodyTradeOrderSettleResponse) *SettleOrderResponseBody {
	s.TradeOrderSettleResponse = v
	return s
}

type SettleOrderResponseBodyTradeOrderSettleResponse struct {
	TradeNo      *string `json:"TradeNo,omitempty" xml:"TradeNo,omitempty"`
	OutRequestNo *string `json:"OutRequestNo,omitempty" xml:"OutRequestNo,omitempty"`
}

func (s SettleOrderResponseBodyTradeOrderSettleResponse) String() string {
	return tea.Prettify(s)
}

func (s SettleOrderResponseBodyTradeOrderSettleResponse) GoString() string {
	return s.String()
}

func (s *SettleOrderResponseBodyTradeOrderSettleResponse) SetTradeNo(v string) *SettleOrderResponseBodyTradeOrderSettleResponse {
	s.TradeNo = &v
	return s
}

func (s *SettleOrderResponseBodyTradeOrderSettleResponse) SetOutRequestNo(v string) *SettleOrderResponseBodyTradeOrderSettleResponse {
	s.OutRequestNo = &v
	return s
}

type SettleOrderResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SettleOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SettleOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SettleOrderResponse) GoString() string {
	return s.String()
}

func (s *SettleOrderResponse) SetHeaders(v map[string]*string) *SettleOrderResponse {
	s.Headers = v
	return s
}

func (s *SettleOrderResponse) SetBody(v *SettleOrderResponseBody) *SettleOrderResponse {
	s.Body = v
	return s
}

type SubmitReturnGoodLogisticsRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid                *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	LogisticsNo           *string `json:"LogisticsNo,omitempty" xml:"LogisticsNo,omitempty"`
	CpCode                *string `json:"CpCode,omitempty" xml:"CpCode,omitempty"`
	SubLmOrderId          *string `json:"SubLmOrderId,omitempty" xml:"SubLmOrderId,omitempty"`
	DisputeId             *int64  `json:"DisputeId,omitempty" xml:"DisputeId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	AccountType           *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s SubmitReturnGoodLogisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitReturnGoodLogisticsRequest) GoString() string {
	return s.String()
}

func (s *SubmitReturnGoodLogisticsRequest) SetBizId(v string) *SubmitReturnGoodLogisticsRequest {
	s.BizId = &v
	return s
}

func (s *SubmitReturnGoodLogisticsRequest) SetBizUid(v string) *SubmitReturnGoodLogisticsRequest {
	s.BizUid = &v
	return s
}

func (s *SubmitReturnGoodLogisticsRequest) SetLogisticsNo(v string) *SubmitReturnGoodLogisticsRequest {
	s.LogisticsNo = &v
	return s
}

func (s *SubmitReturnGoodLogisticsRequest) SetCpCode(v string) *SubmitReturnGoodLogisticsRequest {
	s.CpCode = &v
	return s
}

func (s *SubmitReturnGoodLogisticsRequest) SetSubLmOrderId(v string) *SubmitReturnGoodLogisticsRequest {
	s.SubLmOrderId = &v
	return s
}

func (s *SubmitReturnGoodLogisticsRequest) SetDisputeId(v int64) *SubmitReturnGoodLogisticsRequest {
	s.DisputeId = &v
	return s
}

func (s *SubmitReturnGoodLogisticsRequest) SetUseAnonymousTbAccount(v bool) *SubmitReturnGoodLogisticsRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *SubmitReturnGoodLogisticsRequest) SetThirdPartyUserId(v string) *SubmitReturnGoodLogisticsRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *SubmitReturnGoodLogisticsRequest) SetAccountType(v string) *SubmitReturnGoodLogisticsRequest {
	s.AccountType = &v
	return s
}

type SubmitReturnGoodLogisticsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitReturnGoodLogisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitReturnGoodLogisticsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitReturnGoodLogisticsResponseBody) SetCode(v string) *SubmitReturnGoodLogisticsResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitReturnGoodLogisticsResponseBody) SetMessage(v string) *SubmitReturnGoodLogisticsResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitReturnGoodLogisticsResponseBody) SetRequestId(v string) *SubmitReturnGoodLogisticsResponseBody {
	s.RequestId = &v
	return s
}

type SubmitReturnGoodLogisticsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitReturnGoodLogisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitReturnGoodLogisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitReturnGoodLogisticsResponse) GoString() string {
	return s.String()
}

func (s *SubmitReturnGoodLogisticsResponse) SetHeaders(v map[string]*string) *SubmitReturnGoodLogisticsResponse {
	s.Headers = v
	return s
}

func (s *SubmitReturnGoodLogisticsResponse) SetBody(v *SubmitReturnGoodLogisticsResponseBody) *SubmitReturnGoodLogisticsResponse {
	s.Body = v
	return s
}

type SyncMerchantInfoRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	SellerNick *string `json:"SellerNick,omitempty" xml:"SellerNick,omitempty"`
	ItemList   *string `json:"ItemList,omitempty" xml:"ItemList,omitempty"`
	TimeStamp  *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s SyncMerchantInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncMerchantInfoRequest) GoString() string {
	return s.String()
}

func (s *SyncMerchantInfoRequest) SetBizId(v string) *SyncMerchantInfoRequest {
	s.BizId = &v
	return s
}

func (s *SyncMerchantInfoRequest) SetTaskId(v string) *SyncMerchantInfoRequest {
	s.TaskId = &v
	return s
}

func (s *SyncMerchantInfoRequest) SetSellerNick(v string) *SyncMerchantInfoRequest {
	s.SellerNick = &v
	return s
}

func (s *SyncMerchantInfoRequest) SetItemList(v string) *SyncMerchantInfoRequest {
	s.ItemList = &v
	return s
}

func (s *SyncMerchantInfoRequest) SetTimeStamp(v int64) *SyncMerchantInfoRequest {
	s.TimeStamp = &v
	return s
}

type SyncMerchantInfoResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Model     *SyncMerchantInfoResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s SyncMerchantInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncMerchantInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncMerchantInfoResponseBody) SetCode(v string) *SyncMerchantInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SyncMerchantInfoResponseBody) SetMessage(v string) *SyncMerchantInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SyncMerchantInfoResponseBody) SetRequestId(v string) *SyncMerchantInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncMerchantInfoResponseBody) SetModel(v *SyncMerchantInfoResponseBodyModel) *SyncMerchantInfoResponseBody {
	s.Model = v
	return s
}

type SyncMerchantInfoResponseBodyModel struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SyncMerchantInfoResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s SyncMerchantInfoResponseBodyModel) GoString() string {
	return s.String()
}

func (s *SyncMerchantInfoResponseBodyModel) SetStatus(v string) *SyncMerchantInfoResponseBodyModel {
	s.Status = &v
	return s
}

func (s *SyncMerchantInfoResponseBodyModel) SetUrl(v string) *SyncMerchantInfoResponseBodyModel {
	s.Url = &v
	return s
}

func (s *SyncMerchantInfoResponseBodyModel) SetTaskId(v string) *SyncMerchantInfoResponseBodyModel {
	s.TaskId = &v
	return s
}

type SyncMerchantInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncMerchantInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncMerchantInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncMerchantInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncMerchantInfoResponse) SetHeaders(v map[string]*string) *SyncMerchantInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncMerchantInfoResponse) SetBody(v *SyncMerchantInfoResponseBody) *SyncMerchantInfoResponse {
	s.Body = v
	return s
}

type UnsignWithholdAgreementRequest struct {
	OutRequestNo        *string `json:"OutRequestNo,omitempty" xml:"OutRequestNo,omitempty"`
	ExternalAgreementNo *string `json:"ExternalAgreementNo,omitempty" xml:"ExternalAgreementNo,omitempty"`
	MerchantId          *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
	AgreementNo         *string `json:"AgreementNo,omitempty" xml:"AgreementNo,omitempty"`
}

func (s UnsignWithholdAgreementRequest) String() string {
	return tea.Prettify(s)
}

func (s UnsignWithholdAgreementRequest) GoString() string {
	return s.String()
}

func (s *UnsignWithholdAgreementRequest) SetOutRequestNo(v string) *UnsignWithholdAgreementRequest {
	s.OutRequestNo = &v
	return s
}

func (s *UnsignWithholdAgreementRequest) SetExternalAgreementNo(v string) *UnsignWithholdAgreementRequest {
	s.ExternalAgreementNo = &v
	return s
}

func (s *UnsignWithholdAgreementRequest) SetMerchantId(v string) *UnsignWithholdAgreementRequest {
	s.MerchantId = &v
	return s
}

func (s *UnsignWithholdAgreementRequest) SetAgreementNo(v string) *UnsignWithholdAgreementRequest {
	s.AgreementNo = &v
	return s
}

type UnsignWithholdAgreementResponseBody struct {
	Code                 *string                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message              *string                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WithholdSignResponse *UnsignWithholdAgreementResponseBodyWithholdSignResponse `json:"WithholdSignResponse,omitempty" xml:"WithholdSignResponse,omitempty" type:"Struct"`
}

func (s UnsignWithholdAgreementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnsignWithholdAgreementResponseBody) GoString() string {
	return s.String()
}

func (s *UnsignWithholdAgreementResponseBody) SetCode(v string) *UnsignWithholdAgreementResponseBody {
	s.Code = &v
	return s
}

func (s *UnsignWithholdAgreementResponseBody) SetMessage(v string) *UnsignWithholdAgreementResponseBody {
	s.Message = &v
	return s
}

func (s *UnsignWithholdAgreementResponseBody) SetRequestId(v string) *UnsignWithholdAgreementResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnsignWithholdAgreementResponseBody) SetWithholdSignResponse(v *UnsignWithholdAgreementResponseBodyWithholdSignResponse) *UnsignWithholdAgreementResponseBody {
	s.WithholdSignResponse = v
	return s
}

type UnsignWithholdAgreementResponseBodyWithholdSignResponse struct {
	OutRequestNo *string `json:"OutRequestNo,omitempty" xml:"OutRequestNo,omitempty"`
}

func (s UnsignWithholdAgreementResponseBodyWithholdSignResponse) String() string {
	return tea.Prettify(s)
}

func (s UnsignWithholdAgreementResponseBodyWithholdSignResponse) GoString() string {
	return s.String()
}

func (s *UnsignWithholdAgreementResponseBodyWithholdSignResponse) SetOutRequestNo(v string) *UnsignWithholdAgreementResponseBodyWithholdSignResponse {
	s.OutRequestNo = &v
	return s
}

type UnsignWithholdAgreementResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnsignWithholdAgreementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnsignWithholdAgreementResponse) String() string {
	return tea.Prettify(s)
}

func (s UnsignWithholdAgreementResponse) GoString() string {
	return s.String()
}

func (s *UnsignWithholdAgreementResponse) SetHeaders(v map[string]*string) *UnsignWithholdAgreementResponse {
	s.Headers = v
	return s
}

func (s *UnsignWithholdAgreementResponse) SetBody(v *UnsignWithholdAgreementResponseBody) *UnsignWithholdAgreementResponse {
	s.Body = v
	return s
}

type UpdateAddressRequest struct {
	BizId                 *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ThirdPartyUserId      *string `json:"ThirdPartyUserId,omitempty" xml:"ThirdPartyUserId,omitempty"`
	UseAnonymousTbAccount *bool   `json:"UseAnonymousTbAccount,omitempty" xml:"UseAnonymousTbAccount,omitempty"`
	AddressInfo           *string `json:"AddressInfo,omitempty" xml:"AddressInfo,omitempty"`
}

func (s UpdateAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAddressRequest) GoString() string {
	return s.String()
}

func (s *UpdateAddressRequest) SetBizId(v string) *UpdateAddressRequest {
	s.BizId = &v
	return s
}

func (s *UpdateAddressRequest) SetThirdPartyUserId(v string) *UpdateAddressRequest {
	s.ThirdPartyUserId = &v
	return s
}

func (s *UpdateAddressRequest) SetUseAnonymousTbAccount(v bool) *UpdateAddressRequest {
	s.UseAnonymousTbAccount = &v
	return s
}

func (s *UpdateAddressRequest) SetAddressInfo(v string) *UpdateAddressRequest {
	s.AddressInfo = &v
	return s
}

type UpdateAddressResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAddressResponseBody) SetCode(v string) *UpdateAddressResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAddressResponseBody) SetMessage(v string) *UpdateAddressResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAddressResponseBody) SetRequestId(v string) *UpdateAddressResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAddressResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAddressResponse) GoString() string {
	return s.String()
}

func (s *UpdateAddressResponse) SetHeaders(v map[string]*string) *UpdateAddressResponse {
	s.Headers = v
	return s
}

func (s *UpdateAddressResponse) SetBody(v *UpdateAddressResponseBody) *UpdateAddressResponse {
	s.Body = v
	return s
}

type ValidateTaobaoAccountRequest struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizUid     *string `json:"BizUid,omitempty" xml:"BizUid,omitempty"`
	ExtJson    *string `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	TbUserNick *string `json:"TbUserNick,omitempty" xml:"TbUserNick,omitempty"`
	MobileNo   *string `json:"MobileNo,omitempty" xml:"MobileNo,omitempty"`
}

func (s ValidateTaobaoAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateTaobaoAccountRequest) GoString() string {
	return s.String()
}

func (s *ValidateTaobaoAccountRequest) SetBizId(v string) *ValidateTaobaoAccountRequest {
	s.BizId = &v
	return s
}

func (s *ValidateTaobaoAccountRequest) SetBizUid(v string) *ValidateTaobaoAccountRequest {
	s.BizUid = &v
	return s
}

func (s *ValidateTaobaoAccountRequest) SetExtJson(v string) *ValidateTaobaoAccountRequest {
	s.ExtJson = &v
	return s
}

func (s *ValidateTaobaoAccountRequest) SetTbUserNick(v string) *ValidateTaobaoAccountRequest {
	s.TbUserNick = &v
	return s
}

func (s *ValidateTaobaoAccountRequest) SetMobileNo(v string) *ValidateTaobaoAccountRequest {
	s.MobileNo = &v
	return s
}

type ValidateTaobaoAccountResponseBody struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	SubMessage *string                                 `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Code       *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	SubCode    *string                                 `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	LogsId     *string                                 `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Model      *ValidateTaobaoAccountResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
}

func (s ValidateTaobaoAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateTaobaoAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateTaobaoAccountResponseBody) SetRequestId(v string) *ValidateTaobaoAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateTaobaoAccountResponseBody) SetSuccess(v bool) *ValidateTaobaoAccountResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateTaobaoAccountResponseBody) SetSubMessage(v string) *ValidateTaobaoAccountResponseBody {
	s.SubMessage = &v
	return s
}

func (s *ValidateTaobaoAccountResponseBody) SetCode(v string) *ValidateTaobaoAccountResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateTaobaoAccountResponseBody) SetMessage(v string) *ValidateTaobaoAccountResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateTaobaoAccountResponseBody) SetSubCode(v string) *ValidateTaobaoAccountResponseBody {
	s.SubCode = &v
	return s
}

func (s *ValidateTaobaoAccountResponseBody) SetTotalCount(v int64) *ValidateTaobaoAccountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ValidateTaobaoAccountResponseBody) SetLogsId(v string) *ValidateTaobaoAccountResponseBody {
	s.LogsId = &v
	return s
}

func (s *ValidateTaobaoAccountResponseBody) SetModel(v *ValidateTaobaoAccountResponseBodyModel) *ValidateTaobaoAccountResponseBody {
	s.Model = v
	return s
}

type ValidateTaobaoAccountResponseBodyModel struct {
	Match *bool `json:"Match,omitempty" xml:"Match,omitempty"`
}

func (s ValidateTaobaoAccountResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s ValidateTaobaoAccountResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ValidateTaobaoAccountResponseBodyModel) SetMatch(v bool) *ValidateTaobaoAccountResponseBodyModel {
	s.Match = &v
	return s
}

type ValidateTaobaoAccountResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ValidateTaobaoAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateTaobaoAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateTaobaoAccountResponse) GoString() string {
	return s.String()
}

func (s *ValidateTaobaoAccountResponse) SetHeaders(v map[string]*string) *ValidateTaobaoAccountResponse {
	s.Headers = v
	return s
}

func (s *ValidateTaobaoAccountResponse) SetBody(v *ValidateTaobaoAccountResponseBody) *ValidateTaobaoAccountResponse {
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

func (client *Client) AddAddressWithOptions(request *AddAddressRequest, runtime *util.RuntimeOptions) (_result *AddAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddAddress"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAddress(request *AddAddressRequest) (_result *AddAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAddressResponse{}
	_body, _err := client.AddAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddItemLimitRuleWithOptions(request *AddItemLimitRuleRequest, runtime *util.RuntimeOptions) (_result *AddItemLimitRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddItemLimitRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddItemLimitRule"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddItemLimitRule(request *AddItemLimitRuleRequest) (_result *AddItemLimitRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddItemLimitRuleResponse{}
	_body, _err := client.AddItemLimitRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddItemToSubBizsWithOptions(tmpReq *AddItemToSubBizsRequest, runtime *util.RuntimeOptions) (_result *AddItemToSubBizsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AddItemToSubBizsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SubBizIds)) {
		request.SubBizIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubBizIds, tea.String("SubBizIds"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddItemToSubBizsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddItemToSubBizs"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddItemToSubBizs(request *AddItemToSubBizsRequest) (_result *AddItemToSubBizsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddItemToSubBizsResponse{}
	_body, _err := client.AddItemToSubBizsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSupplierNewItemsWithOptions(request *AddSupplierNewItemsRequest, runtime *util.RuntimeOptions) (_result *AddSupplierNewItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddSupplierNewItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddSupplierNewItems"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSupplierNewItems(request *AddSupplierNewItemsRequest) (_result *AddSupplierNewItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSupplierNewItemsResponse{}
	_body, _err := client.AddSupplierNewItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyRefundWithOptions(request *ApplyRefundRequest, runtime *util.RuntimeOptions) (_result *ApplyRefundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApplyRefundResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplyRefund"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyRefund(request *ApplyRefundRequest) (_result *ApplyRefundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyRefundResponse{}
	_body, _err := client.ApplyRefundWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchRegistAnonymousTbAccountWithOptions(request *BatchRegistAnonymousTbAccountRequest, runtime *util.RuntimeOptions) (_result *BatchRegistAnonymousTbAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchRegistAnonymousTbAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchRegistAnonymousTbAccount"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchRegistAnonymousTbAccount(request *BatchRegistAnonymousTbAccountRequest) (_result *BatchRegistAnonymousTbAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchRegistAnonymousTbAccountResponse{}
	_body, _err := client.BatchRegistAnonymousTbAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelOrderWithOptions(request *CancelOrderRequest, runtime *util.RuntimeOptions) (_result *CancelOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelOrder"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelOrder(request *CancelOrderRequest) (_result *CancelOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelOrderResponse{}
	_body, _err := client.CancelOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelRefundWithOptions(request *CancelRefundRequest, runtime *util.RuntimeOptions) (_result *CancelRefundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelRefundResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelRefund"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelRefund(request *CancelRefundRequest) (_result *CancelRefundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelRefundResponse{}
	_body, _err := client.CancelRefundWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmDisburseWithOptions(request *ConfirmDisburseRequest, runtime *util.RuntimeOptions) (_result *ConfirmDisburseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfirmDisburseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfirmDisburse"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmDisburse(request *ConfirmDisburseRequest) (_result *ConfirmDisburseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfirmDisburseResponse{}
	_body, _err := client.ConfirmDisburseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMovieTicketOrderWithOptions(request *CreateMovieTicketOrderRequest, runtime *util.RuntimeOptions) (_result *CreateMovieTicketOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateMovieTicketOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMovieTicketOrder"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMovieTicketOrder(request *CreateMovieTicketOrderRequest) (_result *CreateMovieTicketOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMovieTicketOrderResponse{}
	_body, _err := client.CreateMovieTicketOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrderWithOptions(request *CreateOrderRequest, runtime *util.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateOrder"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrder(request *CreateOrderRequest) (_result *CreateOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrderResponse{}
	_body, _err := client.CreateOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrderV2WithOptions(request *CreateOrderV2Request, runtime *util.RuntimeOptions) (_result *CreateOrderV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateOrderV2Response{}
	_body, _err := client.DoRPCRequest(tea.String("CreateOrderV2"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrderV2(request *CreateOrderV2Request) (_result *CreateOrderV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrderV2Response{}
	_body, _err := client.CreateOrderV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePayUrlWithOptions(request *CreatePayUrlRequest, runtime *util.RuntimeOptions) (_result *CreatePayUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePayUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePayUrl"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePayUrl(request *CreatePayUrlRequest) (_result *CreatePayUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePayUrlResponse{}
	_body, _err := client.CreatePayUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVirtualProductOrderWithOptions(request *CreateVirtualProductOrderRequest, runtime *util.RuntimeOptions) (_result *CreateVirtualProductOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVirtualProductOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVirtualProductOrder"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVirtualProductOrder(request *CreateVirtualProductOrderRequest) (_result *CreateVirtualProductOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVirtualProductOrderResponse{}
	_body, _err := client.CreateVirtualProductOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWithholdTradeWithOptions(request *CreateWithholdTradeRequest, runtime *util.RuntimeOptions) (_result *CreateWithholdTradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateWithholdTradeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateWithholdTrade"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWithholdTrade(request *CreateWithholdTradeRequest) (_result *CreateWithholdTradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWithholdTradeResponse{}
	_body, _err := client.CreateWithholdTradeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBizItemsWithOptions(request *DeleteBizItemsRequest, runtime *util.RuntimeOptions) (_result *DeleteBizItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteBizItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteBizItems"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBizItems(request *DeleteBizItemsRequest) (_result *DeleteBizItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBizItemsResponse{}
	_body, _err := client.DeleteBizItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteItemLimitRuleWithOptions(request *DeleteItemLimitRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteItemLimitRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteItemLimitRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteItemLimitRule"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteItemLimitRule(request *DeleteItemLimitRuleRequest) (_result *DeleteItemLimitRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteItemLimitRuleResponse{}
	_body, _err := client.DeleteItemLimitRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableOrderWithOptions(request *EnableOrderRequest, runtime *util.RuntimeOptions) (_result *EnableOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableOrder"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableOrder(request *EnableOrderRequest) (_result *EnableOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableOrderResponse{}
	_body, _err := client.EnableOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteNodeWithOptions(request *ExecuteNodeRequest, runtime *util.RuntimeOptions) (_result *ExecuteNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteNodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteNode"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteNode(request *ExecuteNodeRequest) (_result *ExecuteNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteNodeResponse{}
	_body, _err := client.ExecuteNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCategoryChainWithOptions(request *GetCategoryChainRequest, runtime *util.RuntimeOptions) (_result *GetCategoryChainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCategoryChainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCategoryChain"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCategoryChain(request *GetCategoryChainRequest) (_result *GetCategoryChainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCategoryChainResponse{}
	_body, _err := client.GetCategoryChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCategoryListWithOptions(request *GetCategoryListRequest, runtime *util.RuntimeOptions) (_result *GetCategoryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCategoryListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCategoryList"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCategoryList(request *GetCategoryListRequest) (_result *GetCategoryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCategoryListResponse{}
	_body, _err := client.GetCategoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCustomServiceUrlWithOptions(request *GetCustomServiceUrlRequest, runtime *util.RuntimeOptions) (_result *GetCustomServiceUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCustomServiceUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCustomServiceUrl"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCustomServiceUrl(request *GetCustomServiceUrlRequest) (_result *GetCustomServiceUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomServiceUrlResponse{}
	_body, _err := client.GetCustomServiceUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGuidePageWithOptions(request *GetGuidePageRequest, runtime *util.RuntimeOptions) (_result *GetGuidePageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetGuidePageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetGuidePage"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGuidePage(request *GetGuidePageRequest) (_result *GetGuidePageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGuidePageResponse{}
	_body, _err := client.GetGuidePageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetItemPromotionWithOptions(request *GetItemPromotionRequest, runtime *util.RuntimeOptions) (_result *GetItemPromotionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetItemPromotionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetItemPromotion"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetItemPromotion(request *GetItemPromotionRequest) (_result *GetItemPromotionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetItemPromotionResponse{}
	_body, _err := client.GetItemPromotionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLoginPageWithOptions(request *GetLoginPageRequest, runtime *util.RuntimeOptions) (_result *GetLoginPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLoginPageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLoginPage"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLoginPage(request *GetLoginPageRequest) (_result *GetLoginPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoginPageResponse{}
	_body, _err := client.GetLoginPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSwitchUrlWithOptions(request *GetSwitchUrlRequest, runtime *util.RuntimeOptions) (_result *GetSwitchUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSwitchUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSwitchUrl"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSwitchUrl(request *GetSwitchUrlRequest) (_result *GetSwitchUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSwitchUrlResponse{}
	_body, _err := client.GetSwitchUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserInfoWithOptions(request *GetUserInfoRequest, runtime *util.RuntimeOptions) (_result *GetUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserInfo"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserInfo(request *GetUserInfoRequest) (_result *GetUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserInfoResponse{}
	_body, _err := client.GetUserInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWithholdSignPageUrlWithOptions(request *GetWithholdSignPageUrlRequest, runtime *util.RuntimeOptions) (_result *GetWithholdSignPageUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWithholdSignPageUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWithholdSignPageUrl"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWithholdSignPageUrl(request *GetWithholdSignPageUrlRequest) (_result *GetWithholdSignPageUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWithholdSignPageUrlResponse{}
	_body, _err := client.GetWithholdSignPageUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitApplyRefundWithOptions(request *InitApplyRefundRequest, runtime *util.RuntimeOptions) (_result *InitApplyRefundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InitApplyRefundResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InitApplyRefund"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitApplyRefund(request *InitApplyRefundRequest) (_result *InitApplyRefundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitApplyRefundResponse{}
	_body, _err := client.InitApplyRefundWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListItemActivitiesWithOptions(tmpReq *ListItemActivitiesRequest, runtime *util.RuntimeOptions) (_result *ListItemActivitiesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListItemActivitiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LmItemIds)) {
		request.LmItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LmItemIds, tea.String("LmItemIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ItemIds)) {
		request.ItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItemIds, tea.String("ItemIds"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListItemActivitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListItemActivities"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListItemActivities(request *ListItemActivitiesRequest) (_result *ListItemActivitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListItemActivitiesResponse{}
	_body, _err := client.ListItemActivitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBasicAndBizItemsWithOptions(request *ModifyBasicAndBizItemsRequest, runtime *util.RuntimeOptions) (_result *ModifyBasicAndBizItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBasicAndBizItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBasicAndBizItems"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBasicAndBizItems(request *ModifyBasicAndBizItemsRequest) (_result *ModifyBasicAndBizItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBasicAndBizItemsResponse{}
	_body, _err := client.ModifyBasicAndBizItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBizItemsWithOptions(request *ModifyBizItemsRequest, runtime *util.RuntimeOptions) (_result *ModifyBizItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBizItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBizItems"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBizItems(request *ModifyBizItemsRequest) (_result *ModifyBizItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBizItemsResponse{}
	_body, _err := client.ModifyBizItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyItemLimitRuleWithOptions(request *ModifyItemLimitRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyItemLimitRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyItemLimitRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyItemLimitRule"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyItemLimitRule(request *ModifyItemLimitRuleRequest) (_result *ModifyItemLimitRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyItemLimitRuleResponse{}
	_body, _err := client.ModifyItemLimitRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NotifyPayOrderStatusWithOptions(request *NotifyPayOrderStatusRequest, runtime *util.RuntimeOptions) (_result *NotifyPayOrderStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &NotifyPayOrderStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("NotifyPayOrderStatus"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NotifyPayOrderStatus(request *NotifyPayOrderStatusRequest) (_result *NotifyPayOrderStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NotifyPayOrderStatusResponse{}
	_body, _err := client.NotifyPayOrderStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) NotifyWithholdFundWithOptions(request *NotifyWithholdFundRequest, runtime *util.RuntimeOptions) (_result *NotifyWithholdFundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &NotifyWithholdFundResponse{}
	_body, _err := client.DoRPCRequest(tea.String("NotifyWithholdFund"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) NotifyWithholdFund(request *NotifyWithholdFundRequest) (_result *NotifyWithholdFundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NotifyWithholdFundResponse{}
	_body, _err := client.NotifyWithholdFundWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryActivityItemsWithOptions(request *QueryActivityItemsRequest, runtime *util.RuntimeOptions) (_result *QueryActivityItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryActivityItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryActivityItems"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryActivityItems(request *QueryActivityItemsRequest) (_result *QueryActivityItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryActivityItemsResponse{}
	_body, _err := client.QueryActivityItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAddressWithOptions(request *QueryAddressRequest, runtime *util.RuntimeOptions) (_result *QueryAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAddress"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAddress(request *QueryAddressRequest) (_result *QueryAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAddressResponse{}
	_body, _err := client.QueryAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAddressDetailWithOptions(request *QueryAddressDetailRequest, runtime *util.RuntimeOptions) (_result *QueryAddressDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAddressDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAddressDetail"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAddressDetail(request *QueryAddressDetailRequest) (_result *QueryAddressDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAddressDetailResponse{}
	_body, _err := client.QueryAddressDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAddressListWithOptions(request *QueryAddressListRequest, runtime *util.RuntimeOptions) (_result *QueryAddressListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAddressListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAddressList"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAddressList(request *QueryAddressListRequest) (_result *QueryAddressListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAddressListResponse{}
	_body, _err := client.QueryAddressListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAdvertisementSettleInfoWithOptions(request *QueryAdvertisementSettleInfoRequest, runtime *util.RuntimeOptions) (_result *QueryAdvertisementSettleInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAdvertisementSettleInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAdvertisementSettleInfo"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAdvertisementSettleInfo(request *QueryAdvertisementSettleInfoRequest) (_result *QueryAdvertisementSettleInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAdvertisementSettleInfoResponse{}
	_body, _err := client.QueryAdvertisementSettleInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAgreementWithOptions(request *QueryAgreementRequest, runtime *util.RuntimeOptions) (_result *QueryAgreementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAgreementResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAgreement"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAgreement(request *QueryAgreementRequest) (_result *QueryAgreementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAgreementResponse{}
	_body, _err := client.QueryAgreementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllCinemasWithOptions(request *QueryAllCinemasRequest, runtime *util.RuntimeOptions) (_result *QueryAllCinemasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAllCinemasResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAllCinemas"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllCinemas(request *QueryAllCinemasRequest) (_result *QueryAllCinemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAllCinemasResponse{}
	_body, _err := client.QueryAllCinemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAllCitiesWithOptions(tmpReq *QueryAllCitiesRequest, runtime *util.RuntimeOptions) (_result *QueryAllCitiesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryAllCitiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExtJson)) {
		request.ExtJsonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtJson, tea.String("ExtJson"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAllCitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAllCities"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAllCities(request *QueryAllCitiesRequest) (_result *QueryAllCitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAllCitiesResponse{}
	_body, _err := client.QueryAllCitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBatchRegistAnonymousTbAccountResultWithOptions(request *QueryBatchRegistAnonymousTbAccountResultRequest, runtime *util.RuntimeOptions) (_result *QueryBatchRegistAnonymousTbAccountResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryBatchRegistAnonymousTbAccountResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryBatchRegistAnonymousTbAccountResult"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBatchRegistAnonymousTbAccountResult(request *QueryBatchRegistAnonymousTbAccountResultRequest) (_result *QueryBatchRegistAnonymousTbAccountResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBatchRegistAnonymousTbAccountResultResponse{}
	_body, _err := client.QueryBatchRegistAnonymousTbAccountResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBestSession4ItemsWithOptions(tmpReq *QueryBestSession4ItemsRequest, runtime *util.RuntimeOptions) (_result *QueryBestSession4ItemsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryBestSession4ItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LmItemIds)) {
		request.LmItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LmItemIds, tea.String("LmItemIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ItemIds)) {
		request.ItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItemIds, tea.String("ItemIds"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryBestSession4ItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryBestSession4Items"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBestSession4Items(request *QueryBestSession4ItemsRequest) (_result *QueryBestSession4ItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBestSession4ItemsResponse{}
	_body, _err := client.QueryBestSession4ItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBizItemListWithOptions(request *QueryBizItemListRequest, runtime *util.RuntimeOptions) (_result *QueryBizItemListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryBizItemListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryBizItemList"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBizItemList(request *QueryBizItemListRequest) (_result *QueryBizItemListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBizItemListResponse{}
	_body, _err := client.QueryBizItemListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBizItemsWithOptions(request *QueryBizItemsRequest, runtime *util.RuntimeOptions) (_result *QueryBizItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryBizItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryBizItems"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBizItems(request *QueryBizItemsRequest) (_result *QueryBizItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBizItemsResponse{}
	_body, _err := client.QueryBizItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBizItemsWithActivityWithOptions(tmpReq *QueryBizItemsWithActivityRequest, runtime *util.RuntimeOptions) (_result *QueryBizItemsWithActivityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryBizItemsWithActivityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LmItemIds)) {
		request.LmItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LmItemIds, tea.String("LmItemIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ItemIds)) {
		request.ItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItemIds, tea.String("ItemIds"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &QueryBizItemsWithActivityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryBizItemsWithActivity"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBizItemsWithActivity(request *QueryBizItemsWithActivityRequest) (_result *QueryBizItemsWithActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBizItemsWithActivityResponse{}
	_body, _err := client.QueryBizItemsWithActivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGuideItemGroupWithOptions(request *QueryGuideItemGroupRequest, runtime *util.RuntimeOptions) (_result *QueryGuideItemGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryGuideItemGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryGuideItemGroup"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGuideItemGroup(request *QueryGuideItemGroupRequest) (_result *QueryGuideItemGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGuideItemGroupResponse{}
	_body, _err := client.QueryGuideItemGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHotMoviesWithOptions(request *QueryHotMoviesRequest, runtime *util.RuntimeOptions) (_result *QueryHotMoviesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryHotMoviesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryHotMovies"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHotMovies(request *QueryHotMoviesRequest) (_result *QueryHotMoviesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryHotMoviesResponse{}
	_body, _err := client.QueryHotMoviesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryInventoryOfItemsInBizItemGroupWithOptions(tmpReq *QueryInventoryOfItemsInBizItemGroupRequest, runtime *util.RuntimeOptions) (_result *QueryInventoryOfItemsInBizItemGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryInventoryOfItemsInBizItemGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ItemIds)) {
		request.ItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItemIds, tea.String("ItemIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.LmItemIds)) {
		request.LmItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LmItemIds, tea.String("LmItemIds"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &QueryInventoryOfItemsInBizItemGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryInventoryOfItemsInBizItemGroup"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryInventoryOfItemsInBizItemGroup(request *QueryInventoryOfItemsInBizItemGroupRequest) (_result *QueryInventoryOfItemsInBizItemGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInventoryOfItemsInBizItemGroupResponse{}
	_body, _err := client.QueryInventoryOfItemsInBizItemGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryItemDetailWithOptions(request *QueryItemDetailRequest, runtime *util.RuntimeOptions) (_result *QueryItemDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryItemDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryItemDetail"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryItemDetail(request *QueryItemDetailRequest) (_result *QueryItemDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryItemDetailResponse{}
	_body, _err := client.QueryItemDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryItemDetailInnerWithOptions(request *QueryItemDetailInnerRequest, runtime *util.RuntimeOptions) (_result *QueryItemDetailInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryItemDetailInnerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryItemDetailInner"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryItemDetailInner(request *QueryItemDetailInnerRequest) (_result *QueryItemDetailInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryItemDetailInnerResponse{}
	_body, _err := client.QueryItemDetailInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryItemInSubBizsWithOptions(tmpReq *QueryItemInSubBizsRequest, runtime *util.RuntimeOptions) (_result *QueryItemInSubBizsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryItemInSubBizsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SubBizIds)) {
		request.SubBizIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubBizIds, tea.String("SubBizIds"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &QueryItemInSubBizsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryItemInSubBizs"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryItemInSubBizs(request *QueryItemInSubBizsRequest) (_result *QueryItemInSubBizsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryItemInSubBizsResponse{}
	_body, _err := client.QueryItemInSubBizsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryItemInventoryWithOptions(request *QueryItemInventoryRequest, runtime *util.RuntimeOptions) (_result *QueryItemInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryItemInventoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryItemInventory"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryItemInventory(request *QueryItemInventoryRequest) (_result *QueryItemInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryItemInventoryResponse{}
	_body, _err := client.QueryItemInventoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryLogisticsWithOptions(request *QueryLogisticsRequest, runtime *util.RuntimeOptions) (_result *QueryLogisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryLogisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryLogistics"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryLogistics(request *QueryLogisticsRequest) (_result *QueryLogisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryLogisticsResponse{}
	_body, _err := client.QueryLogisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMediaSettleInfoWithOptions(request *QueryMediaSettleInfoRequest, runtime *util.RuntimeOptions) (_result *QueryMediaSettleInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMediaSettleInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMediaSettleInfo"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMediaSettleInfo(request *QueryMediaSettleInfoRequest) (_result *QueryMediaSettleInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMediaSettleInfoResponse{}
	_body, _err := client.QueryMediaSettleInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMessagesWithOptions(request *QueryMessagesRequest, runtime *util.RuntimeOptions) (_result *QueryMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMessagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMessages"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMessages(request *QueryMessagesRequest) (_result *QueryMessagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMessagesResponse{}
	_body, _err := client.QueryMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMovieCommentsWithOptions(request *QueryMovieCommentsRequest, runtime *util.RuntimeOptions) (_result *QueryMovieCommentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMovieCommentsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMovieComments"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMovieComments(request *QueryMovieCommentsRequest) (_result *QueryMovieCommentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMovieCommentsResponse{}
	_body, _err := client.QueryMovieCommentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMovieSchedulesWithOptions(request *QueryMovieSchedulesRequest, runtime *util.RuntimeOptions) (_result *QueryMovieSchedulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMovieSchedulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMovieSchedules"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMovieSchedules(request *QueryMovieSchedulesRequest) (_result *QueryMovieSchedulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMovieSchedulesResponse{}
	_body, _err := client.QueryMovieSchedulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMovieSeatsWithOptions(request *QueryMovieSeatsRequest, runtime *util.RuntimeOptions) (_result *QueryMovieSeatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMovieSeatsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMovieSeats"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMovieSeats(request *QueryMovieSeatsRequest) (_result *QueryMovieSeatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMovieSeatsResponse{}
	_body, _err := client.QueryMovieSeatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMovieTicketsWithOptions(request *QueryMovieTicketsRequest, runtime *util.RuntimeOptions) (_result *QueryMovieTicketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &QueryMovieTicketsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMovieTickets"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMovieTickets(request *QueryMovieTicketsRequest) (_result *QueryMovieTicketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMovieTicketsResponse{}
	_body, _err := client.QueryMovieTicketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderAndPaymentListWithOptions(request *QueryOrderAndPaymentListRequest, runtime *util.RuntimeOptions) (_result *QueryOrderAndPaymentListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryOrderAndPaymentListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryOrderAndPaymentList"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrderAndPaymentList(request *QueryOrderAndPaymentListRequest) (_result *QueryOrderAndPaymentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderAndPaymentListResponse{}
	_body, _err := client.QueryOrderAndPaymentListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderCommissionRateWithOptions(request *QueryOrderCommissionRateRequest, runtime *util.RuntimeOptions) (_result *QueryOrderCommissionRateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryOrderCommissionRateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryOrderCommissionRate"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrderCommissionRate(request *QueryOrderCommissionRateRequest) (_result *QueryOrderCommissionRateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderCommissionRateResponse{}
	_body, _err := client.QueryOrderCommissionRateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderDetailInnerWithOptions(request *QueryOrderDetailInnerRequest, runtime *util.RuntimeOptions) (_result *QueryOrderDetailInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryOrderDetailInnerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryOrderDetailInner"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrderDetailInner(request *QueryOrderDetailInnerRequest) (_result *QueryOrderDetailInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderDetailInnerResponse{}
	_body, _err := client.QueryOrderDetailInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderIdByPayIdWithOptions(request *QueryOrderIdByPayIdRequest, runtime *util.RuntimeOptions) (_result *QueryOrderIdByPayIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryOrderIdByPayIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryOrderIdByPayId"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrderIdByPayId(request *QueryOrderIdByPayIdRequest) (_result *QueryOrderIdByPayIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderIdByPayIdResponse{}
	_body, _err := client.QueryOrderIdByPayIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderInfoAfterSaleWithOptions(request *QueryOrderInfoAfterSaleRequest, runtime *util.RuntimeOptions) (_result *QueryOrderInfoAfterSaleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &QueryOrderInfoAfterSaleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryOrderInfoAfterSale"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrderInfoAfterSale(request *QueryOrderInfoAfterSaleRequest) (_result *QueryOrderInfoAfterSaleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderInfoAfterSaleResponse{}
	_body, _err := client.QueryOrderInfoAfterSaleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderItemInfoByPaymentIdForAiZhanYouWithOptions(request *QueryOrderItemInfoByPaymentIdForAiZhanYouRequest, runtime *util.RuntimeOptions) (_result *QueryOrderItemInfoByPaymentIdForAiZhanYouResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryOrderItemInfoByPaymentIdForAiZhanYouResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryOrderItemInfoByPaymentIdForAiZhanYou"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrderItemInfoByPaymentIdForAiZhanYou(request *QueryOrderItemInfoByPaymentIdForAiZhanYouRequest) (_result *QueryOrderItemInfoByPaymentIdForAiZhanYouResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderItemInfoByPaymentIdForAiZhanYouResponse{}
	_body, _err := client.QueryOrderItemInfoByPaymentIdForAiZhanYouWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderListWithOptions(request *QueryOrderListRequest, runtime *util.RuntimeOptions) (_result *QueryOrderListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryOrderListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryOrderList"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrderList(request *QueryOrderListRequest) (_result *QueryOrderListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderListResponse{}
	_body, _err := client.QueryOrderListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderLogisticsWithOptions(request *QueryOrderLogisticsRequest, runtime *util.RuntimeOptions) (_result *QueryOrderLogisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryOrderLogisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryOrderLogistics"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrderLogistics(request *QueryOrderLogisticsRequest) (_result *QueryOrderLogisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderLogisticsResponse{}
	_body, _err := client.QueryOrderLogisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRefundApplicationDetailWithOptions(request *QueryRefundApplicationDetailRequest, runtime *util.RuntimeOptions) (_result *QueryRefundApplicationDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRefundApplicationDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRefundApplicationDetail"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRefundApplicationDetail(request *QueryRefundApplicationDetailRequest) (_result *QueryRefundApplicationDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRefundApplicationDetailResponse{}
	_body, _err := client.QueryRefundApplicationDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStatementsWithOptions(request *QueryStatementsRequest, runtime *util.RuntimeOptions) (_result *QueryStatementsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryStatementsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryStatements"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStatements(request *QueryStatementsRequest) (_result *QueryStatementsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryStatementsResponse{}
	_body, _err := client.QueryStatementsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUnfinishedActivitiesWithOptions(request *QueryUnfinishedActivitiesRequest, runtime *util.RuntimeOptions) (_result *QueryUnfinishedActivitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryUnfinishedActivitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryUnfinishedActivities"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUnfinishedActivities(request *QueryUnfinishedActivitiesRequest) (_result *QueryUnfinishedActivitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUnfinishedActivitiesResponse{}
	_body, _err := client.QueryUnfinishedActivitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUnfinishedSessionsWithOptions(request *QueryUnfinishedSessionsRequest, runtime *util.RuntimeOptions) (_result *QueryUnfinishedSessionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryUnfinishedSessionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryUnfinishedSessions"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUnfinishedSessions(request *QueryUnfinishedSessionsRequest) (_result *QueryUnfinishedSessionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUnfinishedSessionsResponse{}
	_body, _err := client.QueryUnfinishedSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUnfinishedSessions4ItemsWithOptions(tmpReq *QueryUnfinishedSessions4ItemsRequest, runtime *util.RuntimeOptions) (_result *QueryUnfinishedSessions4ItemsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryUnfinishedSessions4ItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LmItemIds)) {
		request.LmItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LmItemIds, tea.String("LmItemIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ItemIds)) {
		request.ItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItemIds, tea.String("ItemIds"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryUnfinishedSessions4ItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryUnfinishedSessions4Items"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUnfinishedSessions4Items(request *QueryUnfinishedSessions4ItemsRequest) (_result *QueryUnfinishedSessions4ItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUnfinishedSessions4ItemsResponse{}
	_body, _err := client.QueryUnfinishedSessions4ItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUpcomingMoviesWithOptions(tmpReq *QueryUpcomingMoviesRequest, runtime *util.RuntimeOptions) (_result *QueryUpcomingMoviesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryUpcomingMoviesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExtJson)) {
		request.ExtJsonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExtJson, tea.String("ExtJson"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryUpcomingMoviesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryUpcomingMovies"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUpcomingMovies(request *QueryUpcomingMoviesRequest) (_result *QueryUpcomingMoviesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUpcomingMoviesResponse{}
	_body, _err := client.QueryUpcomingMoviesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryWithholdTradeWithOptions(request *QueryWithholdTradeRequest, runtime *util.RuntimeOptions) (_result *QueryWithholdTradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryWithholdTradeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryWithholdTrade"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryWithholdTrade(request *QueryWithholdTradeRequest) (_result *QueryWithholdTradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryWithholdTradeResponse{}
	_body, _err := client.QueryWithholdTradeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefundOrderWithOptions(request *RefundOrderRequest, runtime *util.RuntimeOptions) (_result *RefundOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefundOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefundOrder"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefundOrder(request *RefundOrderRequest) (_result *RefundOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefundOrderResponse{}
	_body, _err := client.RefundOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefundPointWithOptions(request *RefundPointRequest, runtime *util.RuntimeOptions) (_result *RefundPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefundPointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefundPoint"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefundPoint(request *RefundPointRequest) (_result *RefundPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefundPointResponse{}
	_body, _err := client.RefundPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefuseMerchantSyncTaskWithOptions(request *RefuseMerchantSyncTaskRequest, runtime *util.RuntimeOptions) (_result *RefuseMerchantSyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefuseMerchantSyncTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefuseMerchantSyncTask"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefuseMerchantSyncTask(request *RefuseMerchantSyncTaskRequest) (_result *RefuseMerchantSyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefuseMerchantSyncTaskResponse{}
	_body, _err := client.RefuseMerchantSyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegistAnonymousTbAccountWithOptions(request *RegistAnonymousTbAccountRequest, runtime *util.RuntimeOptions) (_result *RegistAnonymousTbAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegistAnonymousTbAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegistAnonymousTbAccount"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegistAnonymousTbAccount(request *RegistAnonymousTbAccountRequest) (_result *RegistAnonymousTbAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegistAnonymousTbAccountResponse{}
	_body, _err := client.RegistAnonymousTbAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseMovieSeatWithOptions(request *ReleaseMovieSeatRequest, runtime *util.RuntimeOptions) (_result *ReleaseMovieSeatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseMovieSeatResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseMovieSeat"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseMovieSeat(request *ReleaseMovieSeatRequest) (_result *ReleaseMovieSeatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseMovieSeatResponse{}
	_body, _err := client.ReleaseMovieSeatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAddressWithOptions(request *RemoveAddressRequest, runtime *util.RuntimeOptions) (_result *RemoveAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveAddress"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAddress(request *RemoveAddressRequest) (_result *RemoveAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveAddressResponse{}
	_body, _err := client.RemoveAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveMessagesWithOptions(request *RemoveMessagesRequest, runtime *util.RuntimeOptions) (_result *RemoveMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveMessagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveMessages"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveMessages(request *RemoveMessagesRequest) (_result *RemoveMessagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveMessagesResponse{}
	_body, _err := client.RemoveMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenderH5OrderWithOptions(request *RenderH5OrderRequest, runtime *util.RuntimeOptions) (_result *RenderH5OrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenderH5OrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenderH5Order"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenderH5Order(request *RenderH5OrderRequest) (_result *RenderH5OrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenderH5OrderResponse{}
	_body, _err := client.RenderH5OrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenderOrderWithOptions(request *RenderOrderRequest, runtime *util.RuntimeOptions) (_result *RenderOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenderOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenderOrder"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenderOrder(request *RenderOrderRequest) (_result *RenderOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenderOrderResponse{}
	_body, _err := client.RenderOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RepayForPayUrlWithOptions(request *RepayForPayUrlRequest, runtime *util.RuntimeOptions) (_result *RepayForPayUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RepayForPayUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RepayForPayUrl"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RepayForPayUrl(request *RepayForPayUrlRequest) (_result *RepayForPayUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RepayForPayUrlResponse{}
	_body, _err := client.RepayForPayUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RepayOrderWithOptions(request *RepayOrderRequest, runtime *util.RuntimeOptions) (_result *RepayOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RepayOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RepayOrder"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RepayOrder(request *RepayOrderRequest) (_result *RepayOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RepayOrderResponse{}
	_body, _err := client.RepayOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReserveMovieSeatWithOptions(request *ReserveMovieSeatRequest, runtime *util.RuntimeOptions) (_result *ReserveMovieSeatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReserveMovieSeatResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReserveMovieSeat"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReserveMovieSeat(request *ReserveMovieSeatRequest) (_result *ReserveMovieSeatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReserveMovieSeatResponse{}
	_body, _err := client.ReserveMovieSeatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SettleOrderWithOptions(request *SettleOrderRequest, runtime *util.RuntimeOptions) (_result *SettleOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SettleOrderResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SettleOrder"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SettleOrder(request *SettleOrderRequest) (_result *SettleOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SettleOrderResponse{}
	_body, _err := client.SettleOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitReturnGoodLogisticsWithOptions(request *SubmitReturnGoodLogisticsRequest, runtime *util.RuntimeOptions) (_result *SubmitReturnGoodLogisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitReturnGoodLogisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitReturnGoodLogistics"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitReturnGoodLogistics(request *SubmitReturnGoodLogisticsRequest) (_result *SubmitReturnGoodLogisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitReturnGoodLogisticsResponse{}
	_body, _err := client.SubmitReturnGoodLogisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncMerchantInfoWithOptions(request *SyncMerchantInfoRequest, runtime *util.RuntimeOptions) (_result *SyncMerchantInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SyncMerchantInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SyncMerchantInfo"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncMerchantInfo(request *SyncMerchantInfoRequest) (_result *SyncMerchantInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncMerchantInfoResponse{}
	_body, _err := client.SyncMerchantInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnsignWithholdAgreementWithOptions(request *UnsignWithholdAgreementRequest, runtime *util.RuntimeOptions) (_result *UnsignWithholdAgreementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnsignWithholdAgreementResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnsignWithholdAgreement"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnsignWithholdAgreement(request *UnsignWithholdAgreementRequest) (_result *UnsignWithholdAgreementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnsignWithholdAgreementResponse{}
	_body, _err := client.UnsignWithholdAgreementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAddressWithOptions(request *UpdateAddressRequest, runtime *util.RuntimeOptions) (_result *UpdateAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAddress"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAddress(request *UpdateAddressRequest) (_result *UpdateAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAddressResponse{}
	_body, _err := client.UpdateAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateTaobaoAccountWithOptions(request *ValidateTaobaoAccountRequest, runtime *util.RuntimeOptions) (_result *ValidateTaobaoAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ValidateTaobaoAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ValidateTaobaoAccount"), tea.String("2018-01-16"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateTaobaoAccount(request *ValidateTaobaoAccountRequest) (_result *ValidateTaobaoAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateTaobaoAccountResponse{}
	_body, _err := client.ValidateTaobaoAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
