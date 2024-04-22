// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApplyCreateDistributionOrderRequest struct {
	BuyerId                *string                                             `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	DeliveryAddress        *string                                             `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	DistributionOutTradeId *string                                             `json:"DistributionOutTradeId,omitempty" xml:"DistributionOutTradeId,omitempty"`
	DistributionSupplierId *string                                             `json:"DistributionSupplierId,omitempty" xml:"DistributionSupplierId,omitempty"`
	DistributorId          *string                                             `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	ExtInfo                *string                                             `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	ItemInfoLists          []*ApplyCreateDistributionOrderRequestItemInfoLists `json:"ItemInfoLists,omitempty" xml:"ItemInfoLists,omitempty" type:"Repeated"`
	TenantId               *string                                             `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ApplyCreateDistributionOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyCreateDistributionOrderRequest) GoString() string {
	return s.String()
}

func (s *ApplyCreateDistributionOrderRequest) SetBuyerId(v string) *ApplyCreateDistributionOrderRequest {
	s.BuyerId = &v
	return s
}

func (s *ApplyCreateDistributionOrderRequest) SetDeliveryAddress(v string) *ApplyCreateDistributionOrderRequest {
	s.DeliveryAddress = &v
	return s
}

func (s *ApplyCreateDistributionOrderRequest) SetDistributionOutTradeId(v string) *ApplyCreateDistributionOrderRequest {
	s.DistributionOutTradeId = &v
	return s
}

func (s *ApplyCreateDistributionOrderRequest) SetDistributionSupplierId(v string) *ApplyCreateDistributionOrderRequest {
	s.DistributionSupplierId = &v
	return s
}

func (s *ApplyCreateDistributionOrderRequest) SetDistributorId(v string) *ApplyCreateDistributionOrderRequest {
	s.DistributorId = &v
	return s
}

func (s *ApplyCreateDistributionOrderRequest) SetExtInfo(v string) *ApplyCreateDistributionOrderRequest {
	s.ExtInfo = &v
	return s
}

func (s *ApplyCreateDistributionOrderRequest) SetItemInfoLists(v []*ApplyCreateDistributionOrderRequestItemInfoLists) *ApplyCreateDistributionOrderRequest {
	s.ItemInfoLists = v
	return s
}

func (s *ApplyCreateDistributionOrderRequest) SetTenantId(v string) *ApplyCreateDistributionOrderRequest {
	s.TenantId = &v
	return s
}

type ApplyCreateDistributionOrderRequestItemInfoLists struct {
	DistributionMallId *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	LmItemId           *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	Price              *int64  `json:"Price,omitempty" xml:"Price,omitempty"`
	Quantity           *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// SKU
	SkuId *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
}

func (s ApplyCreateDistributionOrderRequestItemInfoLists) String() string {
	return tea.Prettify(s)
}

func (s ApplyCreateDistributionOrderRequestItemInfoLists) GoString() string {
	return s.String()
}

func (s *ApplyCreateDistributionOrderRequestItemInfoLists) SetDistributionMallId(v string) *ApplyCreateDistributionOrderRequestItemInfoLists {
	s.DistributionMallId = &v
	return s
}

func (s *ApplyCreateDistributionOrderRequestItemInfoLists) SetLmItemId(v string) *ApplyCreateDistributionOrderRequestItemInfoLists {
	s.LmItemId = &v
	return s
}

func (s *ApplyCreateDistributionOrderRequestItemInfoLists) SetPrice(v int64) *ApplyCreateDistributionOrderRequestItemInfoLists {
	s.Price = &v
	return s
}

func (s *ApplyCreateDistributionOrderRequestItemInfoLists) SetQuantity(v int32) *ApplyCreateDistributionOrderRequestItemInfoLists {
	s.Quantity = &v
	return s
}

func (s *ApplyCreateDistributionOrderRequestItemInfoLists) SetSkuId(v string) *ApplyCreateDistributionOrderRequestItemInfoLists {
	s.SkuId = &v
	return s
}

type ApplyCreateDistributionOrderShrinkRequest struct {
	BuyerId                *string `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	DeliveryAddress        *string `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	DistributionOutTradeId *string `json:"DistributionOutTradeId,omitempty" xml:"DistributionOutTradeId,omitempty"`
	DistributionSupplierId *string `json:"DistributionSupplierId,omitempty" xml:"DistributionSupplierId,omitempty"`
	DistributorId          *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	ExtInfo                *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	ItemInfoListsShrink    *string `json:"ItemInfoLists,omitempty" xml:"ItemInfoLists,omitempty"`
	TenantId               *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ApplyCreateDistributionOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyCreateDistributionOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyCreateDistributionOrderShrinkRequest) SetBuyerId(v string) *ApplyCreateDistributionOrderShrinkRequest {
	s.BuyerId = &v
	return s
}

func (s *ApplyCreateDistributionOrderShrinkRequest) SetDeliveryAddress(v string) *ApplyCreateDistributionOrderShrinkRequest {
	s.DeliveryAddress = &v
	return s
}

func (s *ApplyCreateDistributionOrderShrinkRequest) SetDistributionOutTradeId(v string) *ApplyCreateDistributionOrderShrinkRequest {
	s.DistributionOutTradeId = &v
	return s
}

func (s *ApplyCreateDistributionOrderShrinkRequest) SetDistributionSupplierId(v string) *ApplyCreateDistributionOrderShrinkRequest {
	s.DistributionSupplierId = &v
	return s
}

func (s *ApplyCreateDistributionOrderShrinkRequest) SetDistributorId(v string) *ApplyCreateDistributionOrderShrinkRequest {
	s.DistributorId = &v
	return s
}

func (s *ApplyCreateDistributionOrderShrinkRequest) SetExtInfo(v string) *ApplyCreateDistributionOrderShrinkRequest {
	s.ExtInfo = &v
	return s
}

func (s *ApplyCreateDistributionOrderShrinkRequest) SetItemInfoListsShrink(v string) *ApplyCreateDistributionOrderShrinkRequest {
	s.ItemInfoListsShrink = &v
	return s
}

func (s *ApplyCreateDistributionOrderShrinkRequest) SetTenantId(v string) *ApplyCreateDistributionOrderShrinkRequest {
	s.TenantId = &v
	return s
}

type ApplyCreateDistributionOrderResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *string `json:"Model,omitempty" xml:"Model,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ApplyCreateDistributionOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyCreateDistributionOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyCreateDistributionOrderResponseBody) SetCode(v string) *ApplyCreateDistributionOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyCreateDistributionOrderResponseBody) SetLogsId(v string) *ApplyCreateDistributionOrderResponseBody {
	s.LogsId = &v
	return s
}

func (s *ApplyCreateDistributionOrderResponseBody) SetMessage(v string) *ApplyCreateDistributionOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyCreateDistributionOrderResponseBody) SetModel(v string) *ApplyCreateDistributionOrderResponseBody {
	s.Model = &v
	return s
}

func (s *ApplyCreateDistributionOrderResponseBody) SetPageNumber(v int64) *ApplyCreateDistributionOrderResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ApplyCreateDistributionOrderResponseBody) SetPageSize(v int64) *ApplyCreateDistributionOrderResponseBody {
	s.PageSize = &v
	return s
}

func (s *ApplyCreateDistributionOrderResponseBody) SetRequestId(v string) *ApplyCreateDistributionOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyCreateDistributionOrderResponseBody) SetSubCode(v string) *ApplyCreateDistributionOrderResponseBody {
	s.SubCode = &v
	return s
}

func (s *ApplyCreateDistributionOrderResponseBody) SetSubMessage(v string) *ApplyCreateDistributionOrderResponseBody {
	s.SubMessage = &v
	return s
}

func (s *ApplyCreateDistributionOrderResponseBody) SetSuccess(v bool) *ApplyCreateDistributionOrderResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyCreateDistributionOrderResponseBody) SetTotalCount(v int64) *ApplyCreateDistributionOrderResponseBody {
	s.TotalCount = &v
	return s
}

type ApplyCreateDistributionOrderResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyCreateDistributionOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyCreateDistributionOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyCreateDistributionOrderResponse) GoString() string {
	return s.String()
}

func (s *ApplyCreateDistributionOrderResponse) SetHeaders(v map[string]*string) *ApplyCreateDistributionOrderResponse {
	s.Headers = v
	return s
}

func (s *ApplyCreateDistributionOrderResponse) SetStatusCode(v int32) *ApplyCreateDistributionOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyCreateDistributionOrderResponse) SetBody(v *ApplyCreateDistributionOrderResponseBody) *ApplyCreateDistributionOrderResponse {
	s.Body = v
	return s
}

type ApplyRefund4DistributionRequest struct {
	ApplyReasonTextId      *int64                                              `json:"ApplyReasonTextId,omitempty" xml:"ApplyReasonTextId,omitempty"`
	ApplyRefundCount       *int32                                              `json:"ApplyRefundCount,omitempty" xml:"ApplyRefundCount,omitempty"`
	ApplyRefundFee         *int64                                              `json:"ApplyRefundFee,omitempty" xml:"ApplyRefundFee,omitempty"`
	BizClaimType           *int32                                              `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	DistributorId          *string                                             `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	GoodsStatus            *int32                                              `json:"GoodsStatus,omitempty" xml:"GoodsStatus,omitempty"`
	LeaveMessage           *string                                             `json:"LeaveMessage,omitempty" xml:"LeaveMessage,omitempty"`
	LeavePictureLists      []*ApplyRefund4DistributionRequestLeavePictureLists `json:"LeavePictureLists,omitempty" xml:"LeavePictureLists,omitempty" type:"Repeated"`
	SubDistributionOrderId *string                                             `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
	TenantId               *string                                             `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ApplyRefund4DistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyRefund4DistributionRequest) GoString() string {
	return s.String()
}

func (s *ApplyRefund4DistributionRequest) SetApplyReasonTextId(v int64) *ApplyRefund4DistributionRequest {
	s.ApplyReasonTextId = &v
	return s
}

func (s *ApplyRefund4DistributionRequest) SetApplyRefundCount(v int32) *ApplyRefund4DistributionRequest {
	s.ApplyRefundCount = &v
	return s
}

func (s *ApplyRefund4DistributionRequest) SetApplyRefundFee(v int64) *ApplyRefund4DistributionRequest {
	s.ApplyRefundFee = &v
	return s
}

func (s *ApplyRefund4DistributionRequest) SetBizClaimType(v int32) *ApplyRefund4DistributionRequest {
	s.BizClaimType = &v
	return s
}

func (s *ApplyRefund4DistributionRequest) SetDistributorId(v string) *ApplyRefund4DistributionRequest {
	s.DistributorId = &v
	return s
}

func (s *ApplyRefund4DistributionRequest) SetGoodsStatus(v int32) *ApplyRefund4DistributionRequest {
	s.GoodsStatus = &v
	return s
}

func (s *ApplyRefund4DistributionRequest) SetLeaveMessage(v string) *ApplyRefund4DistributionRequest {
	s.LeaveMessage = &v
	return s
}

func (s *ApplyRefund4DistributionRequest) SetLeavePictureLists(v []*ApplyRefund4DistributionRequestLeavePictureLists) *ApplyRefund4DistributionRequest {
	s.LeavePictureLists = v
	return s
}

func (s *ApplyRefund4DistributionRequest) SetSubDistributionOrderId(v string) *ApplyRefund4DistributionRequest {
	s.SubDistributionOrderId = &v
	return s
}

func (s *ApplyRefund4DistributionRequest) SetTenantId(v string) *ApplyRefund4DistributionRequest {
	s.TenantId = &v
	return s
}

type ApplyRefund4DistributionRequestLeavePictureLists struct {
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Picture *string `json:"Picture,omitempty" xml:"Picture,omitempty"`
}

func (s ApplyRefund4DistributionRequestLeavePictureLists) String() string {
	return tea.Prettify(s)
}

func (s ApplyRefund4DistributionRequestLeavePictureLists) GoString() string {
	return s.String()
}

func (s *ApplyRefund4DistributionRequestLeavePictureLists) SetDesc(v string) *ApplyRefund4DistributionRequestLeavePictureLists {
	s.Desc = &v
	return s
}

func (s *ApplyRefund4DistributionRequestLeavePictureLists) SetPicture(v string) *ApplyRefund4DistributionRequestLeavePictureLists {
	s.Picture = &v
	return s
}

type ApplyRefund4DistributionShrinkRequest struct {
	ApplyReasonTextId       *int64  `json:"ApplyReasonTextId,omitempty" xml:"ApplyReasonTextId,omitempty"`
	ApplyRefundCount        *int32  `json:"ApplyRefundCount,omitempty" xml:"ApplyRefundCount,omitempty"`
	ApplyRefundFee          *int64  `json:"ApplyRefundFee,omitempty" xml:"ApplyRefundFee,omitempty"`
	BizClaimType            *int32  `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	DistributorId           *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	GoodsStatus             *int32  `json:"GoodsStatus,omitempty" xml:"GoodsStatus,omitempty"`
	LeaveMessage            *string `json:"LeaveMessage,omitempty" xml:"LeaveMessage,omitempty"`
	LeavePictureListsShrink *string `json:"LeavePictureLists,omitempty" xml:"LeavePictureLists,omitempty"`
	SubDistributionOrderId  *string `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
	TenantId                *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ApplyRefund4DistributionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyRefund4DistributionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyRefund4DistributionShrinkRequest) SetApplyReasonTextId(v int64) *ApplyRefund4DistributionShrinkRequest {
	s.ApplyReasonTextId = &v
	return s
}

func (s *ApplyRefund4DistributionShrinkRequest) SetApplyRefundCount(v int32) *ApplyRefund4DistributionShrinkRequest {
	s.ApplyRefundCount = &v
	return s
}

func (s *ApplyRefund4DistributionShrinkRequest) SetApplyRefundFee(v int64) *ApplyRefund4DistributionShrinkRequest {
	s.ApplyRefundFee = &v
	return s
}

func (s *ApplyRefund4DistributionShrinkRequest) SetBizClaimType(v int32) *ApplyRefund4DistributionShrinkRequest {
	s.BizClaimType = &v
	return s
}

func (s *ApplyRefund4DistributionShrinkRequest) SetDistributorId(v string) *ApplyRefund4DistributionShrinkRequest {
	s.DistributorId = &v
	return s
}

func (s *ApplyRefund4DistributionShrinkRequest) SetGoodsStatus(v int32) *ApplyRefund4DistributionShrinkRequest {
	s.GoodsStatus = &v
	return s
}

func (s *ApplyRefund4DistributionShrinkRequest) SetLeaveMessage(v string) *ApplyRefund4DistributionShrinkRequest {
	s.LeaveMessage = &v
	return s
}

func (s *ApplyRefund4DistributionShrinkRequest) SetLeavePictureListsShrink(v string) *ApplyRefund4DistributionShrinkRequest {
	s.LeavePictureListsShrink = &v
	return s
}

func (s *ApplyRefund4DistributionShrinkRequest) SetSubDistributionOrderId(v string) *ApplyRefund4DistributionShrinkRequest {
	s.SubDistributionOrderId = &v
	return s
}

func (s *ApplyRefund4DistributionShrinkRequest) SetTenantId(v string) *ApplyRefund4DistributionShrinkRequest {
	s.TenantId = &v
	return s
}

type ApplyRefund4DistributionResponseBody struct {
	Code       *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                    `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *ApplyRefund4DistributionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	PageNumber *int64                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ApplyRefund4DistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyRefund4DistributionResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyRefund4DistributionResponseBody) SetCode(v string) *ApplyRefund4DistributionResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyRefund4DistributionResponseBody) SetLogsId(v string) *ApplyRefund4DistributionResponseBody {
	s.LogsId = &v
	return s
}

func (s *ApplyRefund4DistributionResponseBody) SetMessage(v string) *ApplyRefund4DistributionResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyRefund4DistributionResponseBody) SetModel(v *ApplyRefund4DistributionResponseBodyModel) *ApplyRefund4DistributionResponseBody {
	s.Model = v
	return s
}

func (s *ApplyRefund4DistributionResponseBody) SetPageNumber(v int64) *ApplyRefund4DistributionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ApplyRefund4DistributionResponseBody) SetPageSize(v int64) *ApplyRefund4DistributionResponseBody {
	s.PageSize = &v
	return s
}

func (s *ApplyRefund4DistributionResponseBody) SetRequestId(v string) *ApplyRefund4DistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyRefund4DistributionResponseBody) SetSubCode(v string) *ApplyRefund4DistributionResponseBody {
	s.SubCode = &v
	return s
}

func (s *ApplyRefund4DistributionResponseBody) SetSubMessage(v string) *ApplyRefund4DistributionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *ApplyRefund4DistributionResponseBody) SetSuccess(v bool) *ApplyRefund4DistributionResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyRefund4DistributionResponseBody) SetTotalCount(v int64) *ApplyRefund4DistributionResponseBody {
	s.TotalCount = &v
	return s
}

type ApplyRefund4DistributionResponseBodyModel struct {
	DisputeId              *int64  `json:"DisputeId,omitempty" xml:"DisputeId,omitempty"`
	DisputeStatus          *int32  `json:"DisputeStatus,omitempty" xml:"DisputeStatus,omitempty"`
	DisputeType            *int32  `json:"DisputeType,omitempty" xml:"DisputeType,omitempty"`
	SubDistributionOrderId *string `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
}

func (s ApplyRefund4DistributionResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s ApplyRefund4DistributionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ApplyRefund4DistributionResponseBodyModel) SetDisputeId(v int64) *ApplyRefund4DistributionResponseBodyModel {
	s.DisputeId = &v
	return s
}

func (s *ApplyRefund4DistributionResponseBodyModel) SetDisputeStatus(v int32) *ApplyRefund4DistributionResponseBodyModel {
	s.DisputeStatus = &v
	return s
}

func (s *ApplyRefund4DistributionResponseBodyModel) SetDisputeType(v int32) *ApplyRefund4DistributionResponseBodyModel {
	s.DisputeType = &v
	return s
}

func (s *ApplyRefund4DistributionResponseBodyModel) SetSubDistributionOrderId(v string) *ApplyRefund4DistributionResponseBodyModel {
	s.SubDistributionOrderId = &v
	return s
}

type ApplyRefund4DistributionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyRefund4DistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyRefund4DistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyRefund4DistributionResponse) GoString() string {
	return s.String()
}

func (s *ApplyRefund4DistributionResponse) SetHeaders(v map[string]*string) *ApplyRefund4DistributionResponse {
	s.Headers = v
	return s
}

func (s *ApplyRefund4DistributionResponse) SetStatusCode(v int32) *ApplyRefund4DistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyRefund4DistributionResponse) SetBody(v *ApplyRefund4DistributionResponseBody) *ApplyRefund4DistributionResponse {
	s.Body = v
	return s
}

type CancelDistributionTradeRequest struct {
	DistributionTradeId *string `json:"DistributionTradeId,omitempty" xml:"DistributionTradeId,omitempty"`
	DistributorId       *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	TenantId            *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CancelDistributionTradeRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelDistributionTradeRequest) GoString() string {
	return s.String()
}

func (s *CancelDistributionTradeRequest) SetDistributionTradeId(v string) *CancelDistributionTradeRequest {
	s.DistributionTradeId = &v
	return s
}

func (s *CancelDistributionTradeRequest) SetDistributorId(v string) *CancelDistributionTradeRequest {
	s.DistributorId = &v
	return s
}

func (s *CancelDistributionTradeRequest) SetTenantId(v string) *CancelDistributionTradeRequest {
	s.TenantId = &v
	return s
}

type CancelDistributionTradeResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CancelDistributionTradeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelDistributionTradeResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDistributionTradeResponseBody) SetCode(v string) *CancelDistributionTradeResponseBody {
	s.Code = &v
	return s
}

func (s *CancelDistributionTradeResponseBody) SetLogsId(v string) *CancelDistributionTradeResponseBody {
	s.LogsId = &v
	return s
}

func (s *CancelDistributionTradeResponseBody) SetMessage(v string) *CancelDistributionTradeResponseBody {
	s.Message = &v
	return s
}

func (s *CancelDistributionTradeResponseBody) SetPageNumber(v int64) *CancelDistributionTradeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *CancelDistributionTradeResponseBody) SetPageSize(v int64) *CancelDistributionTradeResponseBody {
	s.PageSize = &v
	return s
}

func (s *CancelDistributionTradeResponseBody) SetRequestId(v string) *CancelDistributionTradeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDistributionTradeResponseBody) SetSubCode(v string) *CancelDistributionTradeResponseBody {
	s.SubCode = &v
	return s
}

func (s *CancelDistributionTradeResponseBody) SetSubMessage(v string) *CancelDistributionTradeResponseBody {
	s.SubMessage = &v
	return s
}

func (s *CancelDistributionTradeResponseBody) SetSuccess(v bool) *CancelDistributionTradeResponseBody {
	s.Success = &v
	return s
}

func (s *CancelDistributionTradeResponseBody) SetTotalCount(v int64) *CancelDistributionTradeResponseBody {
	s.TotalCount = &v
	return s
}

type CancelDistributionTradeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelDistributionTradeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelDistributionTradeResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelDistributionTradeResponse) GoString() string {
	return s.String()
}

func (s *CancelDistributionTradeResponse) SetHeaders(v map[string]*string) *CancelDistributionTradeResponse {
	s.Headers = v
	return s
}

func (s *CancelDistributionTradeResponse) SetStatusCode(v int32) *CancelDistributionTradeResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDistributionTradeResponse) SetBody(v *CancelDistributionTradeResponseBody) *CancelDistributionTradeResponse {
	s.Body = v
	return s
}

type CancelRefund4DistributionRequest struct {
	DisputeId              *int64  `json:"DisputeId,omitempty" xml:"DisputeId,omitempty"`
	DistributorId          *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	SubDistributionOrderId *string `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
	TenantId               *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CancelRefund4DistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelRefund4DistributionRequest) GoString() string {
	return s.String()
}

func (s *CancelRefund4DistributionRequest) SetDisputeId(v int64) *CancelRefund4DistributionRequest {
	s.DisputeId = &v
	return s
}

func (s *CancelRefund4DistributionRequest) SetDistributorId(v string) *CancelRefund4DistributionRequest {
	s.DistributorId = &v
	return s
}

func (s *CancelRefund4DistributionRequest) SetSubDistributionOrderId(v string) *CancelRefund4DistributionRequest {
	s.SubDistributionOrderId = &v
	return s
}

func (s *CancelRefund4DistributionRequest) SetTenantId(v string) *CancelRefund4DistributionRequest {
	s.TenantId = &v
	return s
}

type CancelRefund4DistributionResponseBody struct {
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                     `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *CancelRefund4DistributionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	PageNumber *int64                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CancelRefund4DistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelRefund4DistributionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRefund4DistributionResponseBody) SetCode(v string) *CancelRefund4DistributionResponseBody {
	s.Code = &v
	return s
}

func (s *CancelRefund4DistributionResponseBody) SetLogsId(v string) *CancelRefund4DistributionResponseBody {
	s.LogsId = &v
	return s
}

func (s *CancelRefund4DistributionResponseBody) SetMessage(v string) *CancelRefund4DistributionResponseBody {
	s.Message = &v
	return s
}

func (s *CancelRefund4DistributionResponseBody) SetModel(v *CancelRefund4DistributionResponseBodyModel) *CancelRefund4DistributionResponseBody {
	s.Model = v
	return s
}

func (s *CancelRefund4DistributionResponseBody) SetPageNumber(v int64) *CancelRefund4DistributionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *CancelRefund4DistributionResponseBody) SetPageSize(v int64) *CancelRefund4DistributionResponseBody {
	s.PageSize = &v
	return s
}

func (s *CancelRefund4DistributionResponseBody) SetRequestId(v string) *CancelRefund4DistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelRefund4DistributionResponseBody) SetSubCode(v string) *CancelRefund4DistributionResponseBody {
	s.SubCode = &v
	return s
}

func (s *CancelRefund4DistributionResponseBody) SetSubMessage(v string) *CancelRefund4DistributionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *CancelRefund4DistributionResponseBody) SetSuccess(v bool) *CancelRefund4DistributionResponseBody {
	s.Success = &v
	return s
}

func (s *CancelRefund4DistributionResponseBody) SetTotalCount(v int64) *CancelRefund4DistributionResponseBody {
	s.TotalCount = &v
	return s
}

type CancelRefund4DistributionResponseBodyModel struct {
	DisputeId              *int64  `json:"DisputeId,omitempty" xml:"DisputeId,omitempty"`
	DisputeStatus          *int32  `json:"DisputeStatus,omitempty" xml:"DisputeStatus,omitempty"`
	DisputeType            *int32  `json:"DisputeType,omitempty" xml:"DisputeType,omitempty"`
	SubDistributionOrderId *string `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
}

func (s CancelRefund4DistributionResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s CancelRefund4DistributionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CancelRefund4DistributionResponseBodyModel) SetDisputeId(v int64) *CancelRefund4DistributionResponseBodyModel {
	s.DisputeId = &v
	return s
}

func (s *CancelRefund4DistributionResponseBodyModel) SetDisputeStatus(v int32) *CancelRefund4DistributionResponseBodyModel {
	s.DisputeStatus = &v
	return s
}

func (s *CancelRefund4DistributionResponseBodyModel) SetDisputeType(v int32) *CancelRefund4DistributionResponseBodyModel {
	s.DisputeType = &v
	return s
}

func (s *CancelRefund4DistributionResponseBodyModel) SetSubDistributionOrderId(v string) *CancelRefund4DistributionResponseBodyModel {
	s.SubDistributionOrderId = &v
	return s
}

type CancelRefund4DistributionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRefund4DistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRefund4DistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelRefund4DistributionResponse) GoString() string {
	return s.String()
}

func (s *CancelRefund4DistributionResponse) SetHeaders(v map[string]*string) *CancelRefund4DistributionResponse {
	s.Headers = v
	return s
}

func (s *CancelRefund4DistributionResponse) SetStatusCode(v int32) *CancelRefund4DistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRefund4DistributionResponse) SetBody(v *CancelRefund4DistributionResponseBody) *CancelRefund4DistributionResponse {
	s.Body = v
	return s
}

type ConfirmDisburse4DistributionRequest struct {
	DistributionTradeId     *string `json:"DistributionTradeId,omitempty" xml:"DistributionTradeId,omitempty"`
	DistributorId           *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	MainDistributionOrderId *string `json:"MainDistributionOrderId,omitempty" xml:"MainDistributionOrderId,omitempty"`
	TenantId                *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ConfirmDisburse4DistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDisburse4DistributionRequest) GoString() string {
	return s.String()
}

func (s *ConfirmDisburse4DistributionRequest) SetDistributionTradeId(v string) *ConfirmDisburse4DistributionRequest {
	s.DistributionTradeId = &v
	return s
}

func (s *ConfirmDisburse4DistributionRequest) SetDistributorId(v string) *ConfirmDisburse4DistributionRequest {
	s.DistributorId = &v
	return s
}

func (s *ConfirmDisburse4DistributionRequest) SetMainDistributionOrderId(v string) *ConfirmDisburse4DistributionRequest {
	s.MainDistributionOrderId = &v
	return s
}

func (s *ConfirmDisburse4DistributionRequest) SetTenantId(v string) *ConfirmDisburse4DistributionRequest {
	s.TenantId = &v
	return s
}

type ConfirmDisburse4DistributionResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ConfirmDisburse4DistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDisburse4DistributionResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmDisburse4DistributionResponseBody) SetCode(v string) *ConfirmDisburse4DistributionResponseBody {
	s.Code = &v
	return s
}

func (s *ConfirmDisburse4DistributionResponseBody) SetLogsId(v string) *ConfirmDisburse4DistributionResponseBody {
	s.LogsId = &v
	return s
}

func (s *ConfirmDisburse4DistributionResponseBody) SetMessage(v string) *ConfirmDisburse4DistributionResponseBody {
	s.Message = &v
	return s
}

func (s *ConfirmDisburse4DistributionResponseBody) SetPageNumber(v int64) *ConfirmDisburse4DistributionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ConfirmDisburse4DistributionResponseBody) SetPageSize(v int64) *ConfirmDisburse4DistributionResponseBody {
	s.PageSize = &v
	return s
}

func (s *ConfirmDisburse4DistributionResponseBody) SetRequestId(v string) *ConfirmDisburse4DistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmDisburse4DistributionResponseBody) SetSubCode(v string) *ConfirmDisburse4DistributionResponseBody {
	s.SubCode = &v
	return s
}

func (s *ConfirmDisburse4DistributionResponseBody) SetSubMessage(v string) *ConfirmDisburse4DistributionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *ConfirmDisburse4DistributionResponseBody) SetSuccess(v bool) *ConfirmDisburse4DistributionResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmDisburse4DistributionResponseBody) SetTotalCount(v int64) *ConfirmDisburse4DistributionResponseBody {
	s.TotalCount = &v
	return s
}

type ConfirmDisburse4DistributionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmDisburse4DistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmDisburse4DistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmDisburse4DistributionResponse) GoString() string {
	return s.String()
}

func (s *ConfirmDisburse4DistributionResponse) SetHeaders(v map[string]*string) *ConfirmDisburse4DistributionResponse {
	s.Headers = v
	return s
}

func (s *ConfirmDisburse4DistributionResponse) SetStatusCode(v int32) *ConfirmDisburse4DistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmDisburse4DistributionResponse) SetBody(v *ConfirmDisburse4DistributionResponseBody) *ConfirmDisburse4DistributionResponse {
	s.Body = v
	return s
}

type InitApplyRefund4DistributionRequest struct {
	BizClaimType           *int32  `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	DistributorId          *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	GoodsStatus            *int32  `json:"GoodsStatus,omitempty" xml:"GoodsStatus,omitempty"`
	SubDistributionOrderId *string `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
	TenantId               *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s InitApplyRefund4DistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefund4DistributionRequest) GoString() string {
	return s.String()
}

func (s *InitApplyRefund4DistributionRequest) SetBizClaimType(v int32) *InitApplyRefund4DistributionRequest {
	s.BizClaimType = &v
	return s
}

func (s *InitApplyRefund4DistributionRequest) SetDistributorId(v string) *InitApplyRefund4DistributionRequest {
	s.DistributorId = &v
	return s
}

func (s *InitApplyRefund4DistributionRequest) SetGoodsStatus(v int32) *InitApplyRefund4DistributionRequest {
	s.GoodsStatus = &v
	return s
}

func (s *InitApplyRefund4DistributionRequest) SetSubDistributionOrderId(v string) *InitApplyRefund4DistributionRequest {
	s.SubDistributionOrderId = &v
	return s
}

func (s *InitApplyRefund4DistributionRequest) SetTenantId(v string) *InitApplyRefund4DistributionRequest {
	s.TenantId = &v
	return s
}

type InitApplyRefund4DistributionResponseBody struct {
	Code       *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                        `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *InitApplyRefund4DistributionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	PageNumber *int64                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s InitApplyRefund4DistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefund4DistributionResponseBody) GoString() string {
	return s.String()
}

func (s *InitApplyRefund4DistributionResponseBody) SetCode(v string) *InitApplyRefund4DistributionResponseBody {
	s.Code = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBody) SetLogsId(v string) *InitApplyRefund4DistributionResponseBody {
	s.LogsId = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBody) SetMessage(v string) *InitApplyRefund4DistributionResponseBody {
	s.Message = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBody) SetModel(v *InitApplyRefund4DistributionResponseBodyModel) *InitApplyRefund4DistributionResponseBody {
	s.Model = v
	return s
}

func (s *InitApplyRefund4DistributionResponseBody) SetPageNumber(v int64) *InitApplyRefund4DistributionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBody) SetPageSize(v int64) *InitApplyRefund4DistributionResponseBody {
	s.PageSize = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBody) SetRequestId(v string) *InitApplyRefund4DistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBody) SetSubCode(v string) *InitApplyRefund4DistributionResponseBody {
	s.SubCode = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBody) SetSubMessage(v string) *InitApplyRefund4DistributionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBody) SetSuccess(v bool) *InitApplyRefund4DistributionResponseBody {
	s.Success = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBody) SetTotalCount(v int64) *InitApplyRefund4DistributionResponseBody {
	s.TotalCount = &v
	return s
}

type InitApplyRefund4DistributionResponseBodyModel struct {
	BizClaimType           *int32                                                           `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	MainOrderRefund        *bool                                                            `json:"MainOrderRefund,omitempty" xml:"MainOrderRefund,omitempty"`
	MaxRefundFeeData       *InitApplyRefund4DistributionResponseBodyModelMaxRefundFeeData   `json:"MaxRefundFeeData,omitempty" xml:"MaxRefundFeeData,omitempty" type:"Struct"`
	RefundReasonList       []*InitApplyRefund4DistributionResponseBodyModelRefundReasonList `json:"RefundReasonList,omitempty" xml:"RefundReasonList,omitempty" type:"Repeated"`
	SubDistributionOrderId *string                                                          `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
}

func (s InitApplyRefund4DistributionResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefund4DistributionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *InitApplyRefund4DistributionResponseBodyModel) SetBizClaimType(v int32) *InitApplyRefund4DistributionResponseBodyModel {
	s.BizClaimType = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBodyModel) SetMainOrderRefund(v bool) *InitApplyRefund4DistributionResponseBodyModel {
	s.MainOrderRefund = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBodyModel) SetMaxRefundFeeData(v *InitApplyRefund4DistributionResponseBodyModelMaxRefundFeeData) *InitApplyRefund4DistributionResponseBodyModel {
	s.MaxRefundFeeData = v
	return s
}

func (s *InitApplyRefund4DistributionResponseBodyModel) SetRefundReasonList(v []*InitApplyRefund4DistributionResponseBodyModelRefundReasonList) *InitApplyRefund4DistributionResponseBodyModel {
	s.RefundReasonList = v
	return s
}

func (s *InitApplyRefund4DistributionResponseBodyModel) SetSubDistributionOrderId(v string) *InitApplyRefund4DistributionResponseBodyModel {
	s.SubDistributionOrderId = &v
	return s
}

type InitApplyRefund4DistributionResponseBodyModelMaxRefundFeeData struct {
	MaxRefundFee *int64 `json:"MaxRefundFee,omitempty" xml:"MaxRefundFee,omitempty"`
	MinRefundFee *int64 `json:"MinRefundFee,omitempty" xml:"MinRefundFee,omitempty"`
}

func (s InitApplyRefund4DistributionResponseBodyModelMaxRefundFeeData) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefund4DistributionResponseBodyModelMaxRefundFeeData) GoString() string {
	return s.String()
}

func (s *InitApplyRefund4DistributionResponseBodyModelMaxRefundFeeData) SetMaxRefundFee(v int64) *InitApplyRefund4DistributionResponseBodyModelMaxRefundFeeData {
	s.MaxRefundFee = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBodyModelMaxRefundFeeData) SetMinRefundFee(v int64) *InitApplyRefund4DistributionResponseBodyModelMaxRefundFeeData {
	s.MinRefundFee = &v
	return s
}

type InitApplyRefund4DistributionResponseBodyModelRefundReasonList struct {
	ProofRequired      *bool   `json:"ProofRequired,omitempty" xml:"ProofRequired,omitempty"`
	ReasonTextId       *string `json:"ReasonTextId,omitempty" xml:"ReasonTextId,omitempty"`
	ReasonTips         *string `json:"ReasonTips,omitempty" xml:"ReasonTips,omitempty"`
	RefundDescRequired *bool   `json:"RefundDescRequired,omitempty" xml:"RefundDescRequired,omitempty"`
}

func (s InitApplyRefund4DistributionResponseBodyModelRefundReasonList) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefund4DistributionResponseBodyModelRefundReasonList) GoString() string {
	return s.String()
}

func (s *InitApplyRefund4DistributionResponseBodyModelRefundReasonList) SetProofRequired(v bool) *InitApplyRefund4DistributionResponseBodyModelRefundReasonList {
	s.ProofRequired = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBodyModelRefundReasonList) SetReasonTextId(v string) *InitApplyRefund4DistributionResponseBodyModelRefundReasonList {
	s.ReasonTextId = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBodyModelRefundReasonList) SetReasonTips(v string) *InitApplyRefund4DistributionResponseBodyModelRefundReasonList {
	s.ReasonTips = &v
	return s
}

func (s *InitApplyRefund4DistributionResponseBodyModelRefundReasonList) SetRefundDescRequired(v bool) *InitApplyRefund4DistributionResponseBodyModelRefundReasonList {
	s.RefundDescRequired = &v
	return s
}

type InitApplyRefund4DistributionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitApplyRefund4DistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitApplyRefund4DistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s InitApplyRefund4DistributionResponse) GoString() string {
	return s.String()
}

func (s *InitApplyRefund4DistributionResponse) SetHeaders(v map[string]*string) *InitApplyRefund4DistributionResponse {
	s.Headers = v
	return s
}

func (s *InitApplyRefund4DistributionResponse) SetStatusCode(v int32) *InitApplyRefund4DistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *InitApplyRefund4DistributionResponse) SetBody(v *InitApplyRefund4DistributionResponseBody) *InitApplyRefund4DistributionResponse {
	s.Body = v
	return s
}

type InitModifyRefund4DistributionRequest struct {
	BizClaimType           *int32  `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	DisputeId              *int64  `json:"DisputeId,omitempty" xml:"DisputeId,omitempty"`
	DistributorId          *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	SubDistributionOrderId *string `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
	TenantId               *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s InitModifyRefund4DistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s InitModifyRefund4DistributionRequest) GoString() string {
	return s.String()
}

func (s *InitModifyRefund4DistributionRequest) SetBizClaimType(v int32) *InitModifyRefund4DistributionRequest {
	s.BizClaimType = &v
	return s
}

func (s *InitModifyRefund4DistributionRequest) SetDisputeId(v int64) *InitModifyRefund4DistributionRequest {
	s.DisputeId = &v
	return s
}

func (s *InitModifyRefund4DistributionRequest) SetDistributorId(v string) *InitModifyRefund4DistributionRequest {
	s.DistributorId = &v
	return s
}

func (s *InitModifyRefund4DistributionRequest) SetSubDistributionOrderId(v string) *InitModifyRefund4DistributionRequest {
	s.SubDistributionOrderId = &v
	return s
}

func (s *InitModifyRefund4DistributionRequest) SetTenantId(v string) *InitModifyRefund4DistributionRequest {
	s.TenantId = &v
	return s
}

type InitModifyRefund4DistributionResponseBody struct {
	Code       *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                         `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *InitModifyRefund4DistributionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	PageNumber *int64                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s InitModifyRefund4DistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitModifyRefund4DistributionResponseBody) GoString() string {
	return s.String()
}

func (s *InitModifyRefund4DistributionResponseBody) SetCode(v string) *InitModifyRefund4DistributionResponseBody {
	s.Code = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBody) SetLogsId(v string) *InitModifyRefund4DistributionResponseBody {
	s.LogsId = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBody) SetMessage(v string) *InitModifyRefund4DistributionResponseBody {
	s.Message = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBody) SetModel(v *InitModifyRefund4DistributionResponseBodyModel) *InitModifyRefund4DistributionResponseBody {
	s.Model = v
	return s
}

func (s *InitModifyRefund4DistributionResponseBody) SetPageNumber(v int64) *InitModifyRefund4DistributionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBody) SetPageSize(v int64) *InitModifyRefund4DistributionResponseBody {
	s.PageSize = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBody) SetRequestId(v string) *InitModifyRefund4DistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBody) SetSubCode(v string) *InitModifyRefund4DistributionResponseBody {
	s.SubCode = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBody) SetSubMessage(v string) *InitModifyRefund4DistributionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBody) SetSuccess(v bool) *InitModifyRefund4DistributionResponseBody {
	s.Success = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBody) SetTotalCount(v int64) *InitModifyRefund4DistributionResponseBody {
	s.TotalCount = &v
	return s
}

type InitModifyRefund4DistributionResponseBodyModel struct {
	BizClaimType           *int32                                                            `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	MainOrderRefund        *bool                                                             `json:"MainOrderRefund,omitempty" xml:"MainOrderRefund,omitempty"`
	MaxRefundFeeData       *InitModifyRefund4DistributionResponseBodyModelMaxRefundFeeData   `json:"MaxRefundFeeData,omitempty" xml:"MaxRefundFeeData,omitempty" type:"Struct"`
	RefundReasonList       []*InitModifyRefund4DistributionResponseBodyModelRefundReasonList `json:"RefundReasonList,omitempty" xml:"RefundReasonList,omitempty" type:"Repeated"`
	SubDistributionOrderId *string                                                           `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
}

func (s InitModifyRefund4DistributionResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s InitModifyRefund4DistributionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *InitModifyRefund4DistributionResponseBodyModel) SetBizClaimType(v int32) *InitModifyRefund4DistributionResponseBodyModel {
	s.BizClaimType = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBodyModel) SetMainOrderRefund(v bool) *InitModifyRefund4DistributionResponseBodyModel {
	s.MainOrderRefund = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBodyModel) SetMaxRefundFeeData(v *InitModifyRefund4DistributionResponseBodyModelMaxRefundFeeData) *InitModifyRefund4DistributionResponseBodyModel {
	s.MaxRefundFeeData = v
	return s
}

func (s *InitModifyRefund4DistributionResponseBodyModel) SetRefundReasonList(v []*InitModifyRefund4DistributionResponseBodyModelRefundReasonList) *InitModifyRefund4DistributionResponseBodyModel {
	s.RefundReasonList = v
	return s
}

func (s *InitModifyRefund4DistributionResponseBodyModel) SetSubDistributionOrderId(v string) *InitModifyRefund4DistributionResponseBodyModel {
	s.SubDistributionOrderId = &v
	return s
}

type InitModifyRefund4DistributionResponseBodyModelMaxRefundFeeData struct {
	MaxRefundFee *int64 `json:"MaxRefundFee,omitempty" xml:"MaxRefundFee,omitempty"`
	MinRefundFee *int64 `json:"MinRefundFee,omitempty" xml:"MinRefundFee,omitempty"`
}

func (s InitModifyRefund4DistributionResponseBodyModelMaxRefundFeeData) String() string {
	return tea.Prettify(s)
}

func (s InitModifyRefund4DistributionResponseBodyModelMaxRefundFeeData) GoString() string {
	return s.String()
}

func (s *InitModifyRefund4DistributionResponseBodyModelMaxRefundFeeData) SetMaxRefundFee(v int64) *InitModifyRefund4DistributionResponseBodyModelMaxRefundFeeData {
	s.MaxRefundFee = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBodyModelMaxRefundFeeData) SetMinRefundFee(v int64) *InitModifyRefund4DistributionResponseBodyModelMaxRefundFeeData {
	s.MinRefundFee = &v
	return s
}

type InitModifyRefund4DistributionResponseBodyModelRefundReasonList struct {
	ProofRequired      *bool   `json:"ProofRequired,omitempty" xml:"ProofRequired,omitempty"`
	ReasonTextId       *string `json:"ReasonTextId,omitempty" xml:"ReasonTextId,omitempty"`
	ReasonTips         *string `json:"ReasonTips,omitempty" xml:"ReasonTips,omitempty"`
	RefundDescRequired *bool   `json:"RefundDescRequired,omitempty" xml:"RefundDescRequired,omitempty"`
}

func (s InitModifyRefund4DistributionResponseBodyModelRefundReasonList) String() string {
	return tea.Prettify(s)
}

func (s InitModifyRefund4DistributionResponseBodyModelRefundReasonList) GoString() string {
	return s.String()
}

func (s *InitModifyRefund4DistributionResponseBodyModelRefundReasonList) SetProofRequired(v bool) *InitModifyRefund4DistributionResponseBodyModelRefundReasonList {
	s.ProofRequired = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBodyModelRefundReasonList) SetReasonTextId(v string) *InitModifyRefund4DistributionResponseBodyModelRefundReasonList {
	s.ReasonTextId = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBodyModelRefundReasonList) SetReasonTips(v string) *InitModifyRefund4DistributionResponseBodyModelRefundReasonList {
	s.ReasonTips = &v
	return s
}

func (s *InitModifyRefund4DistributionResponseBodyModelRefundReasonList) SetRefundDescRequired(v bool) *InitModifyRefund4DistributionResponseBodyModelRefundReasonList {
	s.RefundDescRequired = &v
	return s
}

type InitModifyRefund4DistributionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitModifyRefund4DistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitModifyRefund4DistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s InitModifyRefund4DistributionResponse) GoString() string {
	return s.String()
}

func (s *InitModifyRefund4DistributionResponse) SetHeaders(v map[string]*string) *InitModifyRefund4DistributionResponse {
	s.Headers = v
	return s
}

func (s *InitModifyRefund4DistributionResponse) SetStatusCode(v int32) *InitModifyRefund4DistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *InitModifyRefund4DistributionResponse) SetBody(v *InitModifyRefund4DistributionResponseBody) *InitModifyRefund4DistributionResponse {
	s.Body = v
	return s
}

type ListDistributionItemRequest struct {
	DistributionMallId *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributorId      *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	ItemStatus         *int32  `json:"ItemStatus,omitempty" xml:"ItemStatus,omitempty"`
	LmItemId           *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TenantId           *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDistributionItemRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemRequest) GoString() string {
	return s.String()
}

func (s *ListDistributionItemRequest) SetDistributionMallId(v string) *ListDistributionItemRequest {
	s.DistributionMallId = &v
	return s
}

func (s *ListDistributionItemRequest) SetDistributorId(v string) *ListDistributionItemRequest {
	s.DistributorId = &v
	return s
}

func (s *ListDistributionItemRequest) SetItemStatus(v int32) *ListDistributionItemRequest {
	s.ItemStatus = &v
	return s
}

func (s *ListDistributionItemRequest) SetLmItemId(v string) *ListDistributionItemRequest {
	s.LmItemId = &v
	return s
}

func (s *ListDistributionItemRequest) SetPageNumber(v int32) *ListDistributionItemRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDistributionItemRequest) SetPageSize(v int32) *ListDistributionItemRequest {
	s.PageSize = &v
	return s
}

func (s *ListDistributionItemRequest) SetTenantId(v string) *ListDistributionItemRequest {
	s.TenantId = &v
	return s
}

type ListDistributionItemResponseBody struct {
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                  `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      []*ListDistributionItemResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	PageNumber *int64                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDistributionItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemResponseBody) GoString() string {
	return s.String()
}

func (s *ListDistributionItemResponseBody) SetCode(v string) *ListDistributionItemResponseBody {
	s.Code = &v
	return s
}

func (s *ListDistributionItemResponseBody) SetLogsId(v string) *ListDistributionItemResponseBody {
	s.LogsId = &v
	return s
}

func (s *ListDistributionItemResponseBody) SetMessage(v string) *ListDistributionItemResponseBody {
	s.Message = &v
	return s
}

func (s *ListDistributionItemResponseBody) SetModel(v []*ListDistributionItemResponseBodyModel) *ListDistributionItemResponseBody {
	s.Model = v
	return s
}

func (s *ListDistributionItemResponseBody) SetPageNumber(v int64) *ListDistributionItemResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDistributionItemResponseBody) SetPageSize(v int64) *ListDistributionItemResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDistributionItemResponseBody) SetRequestId(v string) *ListDistributionItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDistributionItemResponseBody) SetSubCode(v string) *ListDistributionItemResponseBody {
	s.SubCode = &v
	return s
}

func (s *ListDistributionItemResponseBody) SetSubMessage(v string) *ListDistributionItemResponseBody {
	s.SubMessage = &v
	return s
}

func (s *ListDistributionItemResponseBody) SetSuccess(v bool) *ListDistributionItemResponseBody {
	s.Success = &v
	return s
}

func (s *ListDistributionItemResponseBody) SetTotalCount(v int64) *ListDistributionItemResponseBody {
	s.TotalCount = &v
	return s
}

type ListDistributionItemResponseBodyModel struct {
	Category                *string                                                   `json:"Category,omitempty" xml:"Category,omitempty"`
	CategoryChain           []*ListDistributionItemResponseBodyModelCategoryChain     `json:"CategoryChain,omitempty" xml:"CategoryChain,omitempty" type:"Repeated"`
	CategoryId              *int64                                                    `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	DescOption              *string                                                   `json:"DescOption,omitempty" xml:"DescOption,omitempty"`
	DistributionMallId      *string                                                   `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	GmtCreate               *string                                                   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified             *string                                                   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HasQuantity             *bool                                                     `json:"HasQuantity,omitempty" xml:"HasQuantity,omitempty"`
	IsCanSell               *bool                                                     `json:"IsCanSell,omitempty" xml:"IsCanSell,omitempty"`
	ItemDesc                *string                                                   `json:"ItemDesc,omitempty" xml:"ItemDesc,omitempty"`
	ItemId                  *int64                                                    `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemIdStr               *string                                                   `json:"ItemIdStr,omitempty" xml:"ItemIdStr,omitempty"`
	ItemImages              []*string                                                 `json:"ItemImages,omitempty" xml:"ItemImages,omitempty" type:"Repeated"`
	ItemName                *string                                                   `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	ItemTitle               *string                                                   `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	LmAttributeMap          map[string]*string                                        `json:"LmAttributeMap,omitempty" xml:"LmAttributeMap,omitempty"`
	LmAttributeModels       []*ListDistributionItemResponseBodyModelLmAttributeModels `json:"LmAttributeModels,omitempty" xml:"LmAttributeModels,omitempty" type:"Repeated"`
	LmItemId                *string                                                   `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	MainPicUrl              *string                                                   `json:"MainPicUrl,omitempty" xml:"MainPicUrl,omitempty"`
	PicUrl                  *string                                                   `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	PriceCentScope          *string                                                   `json:"PriceCentScope,omitempty" xml:"PriceCentScope,omitempty"`
	PropertiesJson          *string                                                   `json:"PropertiesJson,omitempty" xml:"PropertiesJson,omitempty"`
	Quantity                *int32                                                    `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ReservedPrice           *int64                                                    `json:"ReservedPrice,omitempty" xml:"ReservedPrice,omitempty"`
	ReservedPriceScope      *string                                                   `json:"ReservedPriceScope,omitempty" xml:"ReservedPriceScope,omitempty"`
	SimpleQuantity          *string                                                   `json:"SimpleQuantity,omitempty" xml:"SimpleQuantity,omitempty"`
	SimpleTotalSoldQuantity *string                                                   `json:"SimpleTotalSoldQuantity,omitempty" xml:"SimpleTotalSoldQuantity,omitempty"`
	SkuList                 []*ListDistributionItemResponseBodyModelSkuList           `json:"SkuList,omitempty" xml:"SkuList,omitempty" type:"Repeated"`
	Status                  *int32                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Tips                    *string                                                   `json:"Tips,omitempty" xml:"Tips,omitempty"`
	TotalSoldQuantity       *int32                                                    `json:"TotalSoldQuantity,omitempty" xml:"TotalSoldQuantity,omitempty"`
}

func (s ListDistributionItemResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ListDistributionItemResponseBodyModel) SetCategory(v string) *ListDistributionItemResponseBodyModel {
	s.Category = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetCategoryChain(v []*ListDistributionItemResponseBodyModelCategoryChain) *ListDistributionItemResponseBodyModel {
	s.CategoryChain = v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetCategoryId(v int64) *ListDistributionItemResponseBodyModel {
	s.CategoryId = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetDescOption(v string) *ListDistributionItemResponseBodyModel {
	s.DescOption = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetDistributionMallId(v string) *ListDistributionItemResponseBodyModel {
	s.DistributionMallId = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetGmtCreate(v string) *ListDistributionItemResponseBodyModel {
	s.GmtCreate = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetGmtModified(v string) *ListDistributionItemResponseBodyModel {
	s.GmtModified = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetHasQuantity(v bool) *ListDistributionItemResponseBodyModel {
	s.HasQuantity = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetIsCanSell(v bool) *ListDistributionItemResponseBodyModel {
	s.IsCanSell = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetItemDesc(v string) *ListDistributionItemResponseBodyModel {
	s.ItemDesc = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetItemId(v int64) *ListDistributionItemResponseBodyModel {
	s.ItemId = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetItemIdStr(v string) *ListDistributionItemResponseBodyModel {
	s.ItemIdStr = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetItemImages(v []*string) *ListDistributionItemResponseBodyModel {
	s.ItemImages = v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetItemName(v string) *ListDistributionItemResponseBodyModel {
	s.ItemName = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetItemTitle(v string) *ListDistributionItemResponseBodyModel {
	s.ItemTitle = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetLmAttributeMap(v map[string]*string) *ListDistributionItemResponseBodyModel {
	s.LmAttributeMap = v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetLmAttributeModels(v []*ListDistributionItemResponseBodyModelLmAttributeModels) *ListDistributionItemResponseBodyModel {
	s.LmAttributeModels = v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetLmItemId(v string) *ListDistributionItemResponseBodyModel {
	s.LmItemId = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetMainPicUrl(v string) *ListDistributionItemResponseBodyModel {
	s.MainPicUrl = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetPicUrl(v string) *ListDistributionItemResponseBodyModel {
	s.PicUrl = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetPriceCentScope(v string) *ListDistributionItemResponseBodyModel {
	s.PriceCentScope = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetPropertiesJson(v string) *ListDistributionItemResponseBodyModel {
	s.PropertiesJson = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetQuantity(v int32) *ListDistributionItemResponseBodyModel {
	s.Quantity = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetReservedPrice(v int64) *ListDistributionItemResponseBodyModel {
	s.ReservedPrice = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetReservedPriceScope(v string) *ListDistributionItemResponseBodyModel {
	s.ReservedPriceScope = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetSimpleQuantity(v string) *ListDistributionItemResponseBodyModel {
	s.SimpleQuantity = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetSimpleTotalSoldQuantity(v string) *ListDistributionItemResponseBodyModel {
	s.SimpleTotalSoldQuantity = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetSkuList(v []*ListDistributionItemResponseBodyModelSkuList) *ListDistributionItemResponseBodyModel {
	s.SkuList = v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetStatus(v int32) *ListDistributionItemResponseBodyModel {
	s.Status = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetTips(v string) *ListDistributionItemResponseBodyModel {
	s.Tips = &v
	return s
}

func (s *ListDistributionItemResponseBodyModel) SetTotalSoldQuantity(v int32) *ListDistributionItemResponseBodyModel {
	s.TotalSoldQuantity = &v
	return s
}

type ListDistributionItemResponseBodyModelCategoryChain struct {
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Leaf       *bool   `json:"Leaf,omitempty" xml:"Leaf,omitempty"`
	Level      *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId   *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s ListDistributionItemResponseBodyModelCategoryChain) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemResponseBodyModelCategoryChain) GoString() string {
	return s.String()
}

func (s *ListDistributionItemResponseBodyModelCategoryChain) SetCategoryId(v int64) *ListDistributionItemResponseBodyModelCategoryChain {
	s.CategoryId = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelCategoryChain) SetLeaf(v bool) *ListDistributionItemResponseBodyModelCategoryChain {
	s.Leaf = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelCategoryChain) SetLevel(v int32) *ListDistributionItemResponseBodyModelCategoryChain {
	s.Level = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelCategoryChain) SetName(v string) *ListDistributionItemResponseBodyModelCategoryChain {
	s.Name = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelCategoryChain) SetParentId(v int64) *ListDistributionItemResponseBodyModelCategoryChain {
	s.ParentId = &v
	return s
}

type ListDistributionItemResponseBodyModelLmAttributeModels struct {
	AttrId      *int64    `json:"AttrId,omitempty" xml:"AttrId,omitempty"`
	Category    *int32    `json:"Category,omitempty" xml:"Category,omitempty"`
	DataType    *string   `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Restriction *string   `json:"Restriction,omitempty" xml:"Restriction,omitempty"`
	ScopeList   []*string `json:"ScopeList,omitempty" xml:"ScopeList,omitempty" type:"Repeated"`
	Value       *string   `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDistributionItemResponseBodyModelLmAttributeModels) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemResponseBodyModelLmAttributeModels) GoString() string {
	return s.String()
}

func (s *ListDistributionItemResponseBodyModelLmAttributeModels) SetAttrId(v int64) *ListDistributionItemResponseBodyModelLmAttributeModels {
	s.AttrId = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelLmAttributeModels) SetCategory(v int32) *ListDistributionItemResponseBodyModelLmAttributeModels {
	s.Category = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelLmAttributeModels) SetDataType(v string) *ListDistributionItemResponseBodyModelLmAttributeModels {
	s.DataType = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelLmAttributeModels) SetDescription(v string) *ListDistributionItemResponseBodyModelLmAttributeModels {
	s.Description = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelLmAttributeModels) SetName(v string) *ListDistributionItemResponseBodyModelLmAttributeModels {
	s.Name = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelLmAttributeModels) SetRestriction(v string) *ListDistributionItemResponseBodyModelLmAttributeModels {
	s.Restriction = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelLmAttributeModels) SetScopeList(v []*string) *ListDistributionItemResponseBodyModelLmAttributeModels {
	s.ScopeList = v
	return s
}

func (s *ListDistributionItemResponseBodyModelLmAttributeModels) SetValue(v string) *ListDistributionItemResponseBodyModelLmAttributeModels {
	s.Value = &v
	return s
}

type ListDistributionItemResponseBodyModelSkuList struct {
	CanSell                *bool                                                            `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	CustomizedAttributeMap map[string]*string                                               `json:"CustomizedAttributeMap,omitempty" xml:"CustomizedAttributeMap,omitempty"`
	ExtInfo                *string                                                          `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	GmtModified            *string                                                          `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HasQuantity            *bool                                                            `json:"HasQuantity,omitempty" xml:"HasQuantity,omitempty"`
	ItemId                 *int64                                                           `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmAttributeModels      []*ListDistributionItemResponseBodyModelSkuListLmAttributeModels `json:"LmAttributeModels,omitempty" xml:"LmAttributeModels,omitempty" type:"Repeated"`
	LmItemId               *string                                                          `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	PriceCent              *int64                                                           `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	Quantity               *int64                                                           `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ReservedPrice          *int64                                                           `json:"ReservedPrice,omitempty" xml:"ReservedPrice,omitempty"`
	SimpleQuantity         *string                                                          `json:"SimpleQuantity,omitempty" xml:"SimpleQuantity,omitempty"`
	SkuDesc                *string                                                          `json:"SkuDesc,omitempty" xml:"SkuDesc,omitempty"`
	SkuId                  *int64                                                           `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuPicUrl              *string                                                          `json:"SkuPicUrl,omitempty" xml:"SkuPicUrl,omitempty"`
	SkuProperties          map[string]*string                                               `json:"SkuProperties,omitempty" xml:"SkuProperties,omitempty"`
	SkuPropertiesJson      *string                                                          `json:"SkuPropertiesJson,omitempty" xml:"SkuPropertiesJson,omitempty"`
	SkuTitle               *string                                                          `json:"SkuTitle,omitempty" xml:"SkuTitle,omitempty"`
	Status                 *int32                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Tips                   *string                                                          `json:"Tips,omitempty" xml:"Tips,omitempty"`
	LmSkuAttributeMap      map[string]*string                                               `json:"lmSkuAttributeMap,omitempty" xml:"lmSkuAttributeMap,omitempty"`
}

func (s ListDistributionItemResponseBodyModelSkuList) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemResponseBodyModelSkuList) GoString() string {
	return s.String()
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetCanSell(v bool) *ListDistributionItemResponseBodyModelSkuList {
	s.CanSell = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetCustomizedAttributeMap(v map[string]*string) *ListDistributionItemResponseBodyModelSkuList {
	s.CustomizedAttributeMap = v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetExtInfo(v string) *ListDistributionItemResponseBodyModelSkuList {
	s.ExtInfo = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetGmtModified(v string) *ListDistributionItemResponseBodyModelSkuList {
	s.GmtModified = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetHasQuantity(v bool) *ListDistributionItemResponseBodyModelSkuList {
	s.HasQuantity = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetItemId(v int64) *ListDistributionItemResponseBodyModelSkuList {
	s.ItemId = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetLmAttributeModels(v []*ListDistributionItemResponseBodyModelSkuListLmAttributeModels) *ListDistributionItemResponseBodyModelSkuList {
	s.LmAttributeModels = v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetLmItemId(v string) *ListDistributionItemResponseBodyModelSkuList {
	s.LmItemId = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetPriceCent(v int64) *ListDistributionItemResponseBodyModelSkuList {
	s.PriceCent = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetQuantity(v int64) *ListDistributionItemResponseBodyModelSkuList {
	s.Quantity = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetReservedPrice(v int64) *ListDistributionItemResponseBodyModelSkuList {
	s.ReservedPrice = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetSimpleQuantity(v string) *ListDistributionItemResponseBodyModelSkuList {
	s.SimpleQuantity = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetSkuDesc(v string) *ListDistributionItemResponseBodyModelSkuList {
	s.SkuDesc = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetSkuId(v int64) *ListDistributionItemResponseBodyModelSkuList {
	s.SkuId = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetSkuPicUrl(v string) *ListDistributionItemResponseBodyModelSkuList {
	s.SkuPicUrl = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetSkuProperties(v map[string]*string) *ListDistributionItemResponseBodyModelSkuList {
	s.SkuProperties = v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetSkuPropertiesJson(v string) *ListDistributionItemResponseBodyModelSkuList {
	s.SkuPropertiesJson = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetSkuTitle(v string) *ListDistributionItemResponseBodyModelSkuList {
	s.SkuTitle = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetStatus(v int32) *ListDistributionItemResponseBodyModelSkuList {
	s.Status = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetTips(v string) *ListDistributionItemResponseBodyModelSkuList {
	s.Tips = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuList) SetLmSkuAttributeMap(v map[string]*string) *ListDistributionItemResponseBodyModelSkuList {
	s.LmSkuAttributeMap = v
	return s
}

type ListDistributionItemResponseBodyModelSkuListLmAttributeModels struct {
	AttrId      *int64    `json:"AttrId,omitempty" xml:"AttrId,omitempty"`
	Category    *int32    `json:"Category,omitempty" xml:"Category,omitempty"`
	DataType    *string   `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Restriction *string   `json:"Restriction,omitempty" xml:"Restriction,omitempty"`
	ScopeList   []*string `json:"ScopeList,omitempty" xml:"ScopeList,omitempty" type:"Repeated"`
	Value       *string   `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDistributionItemResponseBodyModelSkuListLmAttributeModels) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemResponseBodyModelSkuListLmAttributeModels) GoString() string {
	return s.String()
}

func (s *ListDistributionItemResponseBodyModelSkuListLmAttributeModels) SetAttrId(v int64) *ListDistributionItemResponseBodyModelSkuListLmAttributeModels {
	s.AttrId = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuListLmAttributeModels) SetCategory(v int32) *ListDistributionItemResponseBodyModelSkuListLmAttributeModels {
	s.Category = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuListLmAttributeModels) SetDataType(v string) *ListDistributionItemResponseBodyModelSkuListLmAttributeModels {
	s.DataType = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuListLmAttributeModels) SetDescription(v string) *ListDistributionItemResponseBodyModelSkuListLmAttributeModels {
	s.Description = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuListLmAttributeModels) SetName(v string) *ListDistributionItemResponseBodyModelSkuListLmAttributeModels {
	s.Name = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuListLmAttributeModels) SetRestriction(v string) *ListDistributionItemResponseBodyModelSkuListLmAttributeModels {
	s.Restriction = &v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuListLmAttributeModels) SetScopeList(v []*string) *ListDistributionItemResponseBodyModelSkuListLmAttributeModels {
	s.ScopeList = v
	return s
}

func (s *ListDistributionItemResponseBodyModelSkuListLmAttributeModels) SetValue(v string) *ListDistributionItemResponseBodyModelSkuListLmAttributeModels {
	s.Value = &v
	return s
}

type ListDistributionItemResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDistributionItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDistributionItemResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemResponse) GoString() string {
	return s.String()
}

func (s *ListDistributionItemResponse) SetHeaders(v map[string]*string) *ListDistributionItemResponse {
	s.Headers = v
	return s
}

func (s *ListDistributionItemResponse) SetStatusCode(v int32) *ListDistributionItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDistributionItemResponse) SetBody(v *ListDistributionItemResponseBody) *ListDistributionItemResponse {
	s.Body = v
	return s
}

type ListDistributionItemWithoutCacheRequest struct {
	DistributionMallId *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributorId      *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	ItemStatus         *int32  `json:"ItemStatus,omitempty" xml:"ItemStatus,omitempty"`
	LmItemId           *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TenantId           *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDistributionItemWithoutCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemWithoutCacheRequest) GoString() string {
	return s.String()
}

func (s *ListDistributionItemWithoutCacheRequest) SetDistributionMallId(v string) *ListDistributionItemWithoutCacheRequest {
	s.DistributionMallId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheRequest) SetDistributorId(v string) *ListDistributionItemWithoutCacheRequest {
	s.DistributorId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheRequest) SetItemStatus(v int32) *ListDistributionItemWithoutCacheRequest {
	s.ItemStatus = &v
	return s
}

func (s *ListDistributionItemWithoutCacheRequest) SetLmItemId(v string) *ListDistributionItemWithoutCacheRequest {
	s.LmItemId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheRequest) SetPageNumber(v int32) *ListDistributionItemWithoutCacheRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDistributionItemWithoutCacheRequest) SetPageSize(v int32) *ListDistributionItemWithoutCacheRequest {
	s.PageSize = &v
	return s
}

func (s *ListDistributionItemWithoutCacheRequest) SetTenantId(v string) *ListDistributionItemWithoutCacheRequest {
	s.TenantId = &v
	return s
}

type ListDistributionItemWithoutCacheResponseBody struct {
	Code       *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                              `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      []*ListDistributionItemWithoutCacheResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	PageNumber *int64                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string                                              `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string                                              `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDistributionItemWithoutCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemWithoutCacheResponseBody) GoString() string {
	return s.String()
}

func (s *ListDistributionItemWithoutCacheResponseBody) SetCode(v string) *ListDistributionItemWithoutCacheResponseBody {
	s.Code = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBody) SetLogsId(v string) *ListDistributionItemWithoutCacheResponseBody {
	s.LogsId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBody) SetMessage(v string) *ListDistributionItemWithoutCacheResponseBody {
	s.Message = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBody) SetModel(v []*ListDistributionItemWithoutCacheResponseBodyModel) *ListDistributionItemWithoutCacheResponseBody {
	s.Model = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBody) SetPageNumber(v int64) *ListDistributionItemWithoutCacheResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBody) SetPageSize(v int64) *ListDistributionItemWithoutCacheResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBody) SetRequestId(v string) *ListDistributionItemWithoutCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBody) SetSubCode(v string) *ListDistributionItemWithoutCacheResponseBody {
	s.SubCode = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBody) SetSubMessage(v string) *ListDistributionItemWithoutCacheResponseBody {
	s.SubMessage = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBody) SetSuccess(v bool) *ListDistributionItemWithoutCacheResponseBody {
	s.Success = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBody) SetTotalCount(v int64) *ListDistributionItemWithoutCacheResponseBody {
	s.TotalCount = &v
	return s
}

type ListDistributionItemWithoutCacheResponseBodyModel struct {
	CategoryId             *int64                                                           `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryIds            []*int64                                                         `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	City                   *string                                                          `json:"City,omitempty" xml:"City,omitempty"`
	Current                *string                                                          `json:"Current,omitempty" xml:"Current,omitempty"`
	CustomizedAttributeMap map[string]*string                                               `json:"CustomizedAttributeMap,omitempty" xml:"CustomizedAttributeMap,omitempty"`
	DescOption             *string                                                          `json:"DescOption,omitempty" xml:"DescOption,omitempty"`
	DescPath               *string                                                          `json:"DescPath,omitempty" xml:"DescPath,omitempty"`
	DistributionMallId     *string                                                          `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	Features               map[string]*string                                               `json:"Features,omitempty" xml:"Features,omitempty"`
	FirstPicUrl            *string                                                          `json:"FirstPicUrl,omitempty" xml:"FirstPicUrl,omitempty"`
	HasQuantity            *bool                                                            `json:"HasQuantity,omitempty" xml:"HasQuantity,omitempty"`
	IforestProps           []map[string]*string                                             `json:"IforestProps,omitempty" xml:"IforestProps,omitempty" type:"Repeated"`
	InvoiceType            *int32                                                           `json:"InvoiceType,omitempty" xml:"InvoiceType,omitempty"`
	IsCanNotBeSoldCode     *string                                                          `json:"IsCanNotBeSoldCode,omitempty" xml:"IsCanNotBeSoldCode,omitempty"`
	IsCanNotBeSoldMessage  *string                                                          `json:"IsCanNotBeSoldMessage,omitempty" xml:"IsCanNotBeSoldMessage,omitempty"`
	IsCanSell              *bool                                                            `json:"IsCanSell,omitempty" xml:"IsCanSell,omitempty"`
	IsSellerPayPostfee     *bool                                                            `json:"IsSellerPayPostfee,omitempty" xml:"IsSellerPayPostfee,omitempty"`
	ItemId                 *int64                                                           `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemImages             []*string                                                        `json:"ItemImages,omitempty" xml:"ItemImages,omitempty" type:"Repeated"`
	ItemTitle              *string                                                          `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	ItemTotalSimpleValue   *string                                                          `json:"ItemTotalSimpleValue,omitempty" xml:"ItemTotalSimpleValue,omitempty"`
	ItemTotalValue         *int32                                                           `json:"ItemTotalValue,omitempty" xml:"ItemTotalValue,omitempty"`
	LmItemAttributeMap     map[string]*string                                               `json:"LmItemAttributeMap,omitempty" xml:"LmItemAttributeMap,omitempty"`
	LmItemCategory         *string                                                          `json:"LmItemCategory,omitempty" xml:"LmItemCategory,omitempty"`
	LmItemId               *string                                                          `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	MainPicUrl             *string                                                          `json:"MainPicUrl,omitempty" xml:"MainPicUrl,omitempty"`
	MinPrice               *int64                                                           `json:"MinPrice,omitempty" xml:"MinPrice,omitempty"`
	Properties             map[string][]*string                                             `json:"Properties,omitempty" xml:"Properties,omitempty"`
	Prov                   *string                                                          `json:"Prov,omitempty" xml:"Prov,omitempty"`
	Quantity               *int32                                                           `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ReservedPrice          *int64                                                           `json:"ReservedPrice,omitempty" xml:"ReservedPrice,omitempty"`
	SecuredTransactions    *int32                                                           `json:"SecuredTransactions,omitempty" xml:"SecuredTransactions,omitempty"`
	SimpleQuantity         *string                                                          `json:"SimpleQuantity,omitempty" xml:"SimpleQuantity,omitempty"`
	SkuModels              []*ListDistributionItemWithoutCacheResponseBodyModelSkuModels    `json:"SkuModels,omitempty" xml:"SkuModels,omitempty" type:"Repeated"`
	SkuPropertys           []*ListDistributionItemWithoutCacheResponseBodyModelSkuPropertys `json:"SkuPropertys,omitempty" xml:"SkuPropertys,omitempty" type:"Repeated"`
	ThirdPartyItemId       *string                                                          `json:"ThirdPartyItemId,omitempty" xml:"ThirdPartyItemId,omitempty"`
	ThirdPartyName         *string                                                          `json:"ThirdPartyName,omitempty" xml:"ThirdPartyName,omitempty"`
	UserType               *int32                                                           `json:"UserType,omitempty" xml:"UserType,omitempty"`
	VideoPicUrl            *string                                                          `json:"VideoPicUrl,omitempty" xml:"VideoPicUrl,omitempty"`
	VideoUrl               *string                                                          `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	VirtualItemType        *string                                                          `json:"VirtualItemType,omitempty" xml:"VirtualItemType,omitempty"`
}

func (s ListDistributionItemWithoutCacheResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemWithoutCacheResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetCategoryId(v int64) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.CategoryId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetCategoryIds(v []*int64) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.CategoryIds = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetCity(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.City = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetCurrent(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.Current = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetCustomizedAttributeMap(v map[string]*string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.CustomizedAttributeMap = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetDescOption(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.DescOption = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetDescPath(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.DescPath = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetDistributionMallId(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.DistributionMallId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetFeatures(v map[string]*string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.Features = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetFirstPicUrl(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.FirstPicUrl = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetHasQuantity(v bool) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.HasQuantity = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetIforestProps(v []map[string]*string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.IforestProps = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetInvoiceType(v int32) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.InvoiceType = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetIsCanNotBeSoldCode(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.IsCanNotBeSoldCode = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetIsCanNotBeSoldMessage(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.IsCanNotBeSoldMessage = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetIsCanSell(v bool) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.IsCanSell = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetIsSellerPayPostfee(v bool) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.IsSellerPayPostfee = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetItemId(v int64) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.ItemId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetItemImages(v []*string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.ItemImages = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetItemTitle(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.ItemTitle = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetItemTotalSimpleValue(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.ItemTotalSimpleValue = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetItemTotalValue(v int32) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.ItemTotalValue = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetLmItemAttributeMap(v map[string]*string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.LmItemAttributeMap = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetLmItemCategory(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.LmItemCategory = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetLmItemId(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.LmItemId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetMainPicUrl(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.MainPicUrl = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetMinPrice(v int64) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.MinPrice = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetProperties(v map[string][]*string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.Properties = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetProv(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.Prov = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetQuantity(v int32) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.Quantity = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetReservedPrice(v int64) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.ReservedPrice = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetSecuredTransactions(v int32) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.SecuredTransactions = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetSimpleQuantity(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.SimpleQuantity = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetSkuModels(v []*ListDistributionItemWithoutCacheResponseBodyModelSkuModels) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.SkuModels = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetSkuPropertys(v []*ListDistributionItemWithoutCacheResponseBodyModelSkuPropertys) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.SkuPropertys = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetThirdPartyItemId(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.ThirdPartyItemId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetThirdPartyName(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.ThirdPartyName = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetUserType(v int32) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.UserType = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetVideoPicUrl(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.VideoPicUrl = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetVideoUrl(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.VideoUrl = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModel) SetVirtualItemType(v string) *ListDistributionItemWithoutCacheResponseBodyModel {
	s.VirtualItemType = &v
	return s
}

type ListDistributionItemWithoutCacheResponseBodyModelSkuModels struct {
	CustomizedAttributeMap map[string]*string `json:"CustomizedAttributeMap,omitempty" xml:"CustomizedAttributeMap,omitempty"`
	DistributionMallId     *string            `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	ExtJson                *string            `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	HasQuantity            *bool              `json:"HasQuantity,omitempty" xml:"HasQuantity,omitempty"`
	InvoiceType            *int32             `json:"InvoiceType,omitempty" xml:"InvoiceType,omitempty"`
	IsCanNotBeSoldCode     *string            `json:"IsCanNotBeSoldCode,omitempty" xml:"IsCanNotBeSoldCode,omitempty"`
	IsCanNotBeSoldMessage  *string            `json:"IsCanNotBeSoldMessage,omitempty" xml:"IsCanNotBeSoldMessage,omitempty"`
	ItemId                 *int64             `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId               *string            `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	LmSkuAttributeMap      map[string]*string `json:"LmSkuAttributeMap,omitempty" xml:"LmSkuAttributeMap,omitempty"`
	PriceCent              *int64             `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	Quantity               *int32             `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ReservedPrice          *int64             `json:"ReservedPrice,omitempty" xml:"ReservedPrice,omitempty"`
	SimpleQuantity         *string            `json:"SimpleQuantity,omitempty" xml:"SimpleQuantity,omitempty"`
	SkuId                  *int64             `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuPicUrl              *string            `json:"SkuPicUrl,omitempty" xml:"SkuPicUrl,omitempty"`
	SkuPvs                 *string            `json:"SkuPvs,omitempty" xml:"SkuPvs,omitempty"`
	SkuTitle               *string            `json:"SkuTitle,omitempty" xml:"SkuTitle,omitempty"`
	Status                 *int32             `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierPrice          *int64             `json:"SupplierPrice,omitempty" xml:"SupplierPrice,omitempty"`
}

func (s ListDistributionItemWithoutCacheResponseBodyModelSkuModels) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemWithoutCacheResponseBodyModelSkuModels) GoString() string {
	return s.String()
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetCustomizedAttributeMap(v map[string]*string) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.CustomizedAttributeMap = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetDistributionMallId(v string) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.DistributionMallId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetExtJson(v string) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.ExtJson = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetHasQuantity(v bool) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.HasQuantity = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetInvoiceType(v int32) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.InvoiceType = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetIsCanNotBeSoldCode(v string) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.IsCanNotBeSoldCode = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetIsCanNotBeSoldMessage(v string) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.IsCanNotBeSoldMessage = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetItemId(v int64) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.ItemId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetLmItemId(v string) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.LmItemId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetLmSkuAttributeMap(v map[string]*string) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.LmSkuAttributeMap = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetPriceCent(v int64) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.PriceCent = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetQuantity(v int32) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.Quantity = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetReservedPrice(v int64) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.ReservedPrice = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetSimpleQuantity(v string) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.SimpleQuantity = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetSkuId(v int64) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.SkuId = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetSkuPicUrl(v string) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.SkuPicUrl = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetSkuPvs(v string) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.SkuPvs = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetSkuTitle(v string) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.SkuTitle = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetStatus(v int32) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.Status = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuModels) SetSupplierPrice(v int64) *ListDistributionItemWithoutCacheResponseBodyModelSkuModels {
	s.SupplierPrice = &v
	return s
}

type ListDistributionItemWithoutCacheResponseBodyModelSkuPropertys struct {
	Id     *int64                                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	Text   *string                                                                `json:"Text,omitempty" xml:"Text,omitempty"`
	Values []*ListDistributionItemWithoutCacheResponseBodyModelSkuPropertysValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListDistributionItemWithoutCacheResponseBodyModelSkuPropertys) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemWithoutCacheResponseBodyModelSkuPropertys) GoString() string {
	return s.String()
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuPropertys) SetId(v int64) *ListDistributionItemWithoutCacheResponseBodyModelSkuPropertys {
	s.Id = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuPropertys) SetText(v string) *ListDistributionItemWithoutCacheResponseBodyModelSkuPropertys {
	s.Text = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuPropertys) SetValues(v []*ListDistributionItemWithoutCacheResponseBodyModelSkuPropertysValues) *ListDistributionItemWithoutCacheResponseBodyModelSkuPropertys {
	s.Values = v
	return s
}

type ListDistributionItemWithoutCacheResponseBodyModelSkuPropertysValues struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ListDistributionItemWithoutCacheResponseBodyModelSkuPropertysValues) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemWithoutCacheResponseBodyModelSkuPropertysValues) GoString() string {
	return s.String()
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuPropertysValues) SetId(v int64) *ListDistributionItemWithoutCacheResponseBodyModelSkuPropertysValues {
	s.Id = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponseBodyModelSkuPropertysValues) SetText(v string) *ListDistributionItemWithoutCacheResponseBodyModelSkuPropertysValues {
	s.Text = &v
	return s
}

type ListDistributionItemWithoutCacheResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDistributionItemWithoutCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDistributionItemWithoutCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionItemWithoutCacheResponse) GoString() string {
	return s.String()
}

func (s *ListDistributionItemWithoutCacheResponse) SetHeaders(v map[string]*string) *ListDistributionItemWithoutCacheResponse {
	s.Headers = v
	return s
}

func (s *ListDistributionItemWithoutCacheResponse) SetStatusCode(v int32) *ListDistributionItemWithoutCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDistributionItemWithoutCacheResponse) SetBody(v *ListDistributionItemWithoutCacheResponseBody) *ListDistributionItemWithoutCacheResponse {
	s.Body = v
	return s
}

type ListDistributionMallRequest struct {
	ChannelSupplierId    *string `json:"ChannelSupplierId,omitempty" xml:"ChannelSupplierId,omitempty"`
	DistributionMallId   *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributionMallName *string `json:"DistributionMallName,omitempty" xml:"DistributionMallName,omitempty"`
	DistributorId        *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	EndDate              *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate            *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDistributionMallRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionMallRequest) GoString() string {
	return s.String()
}

func (s *ListDistributionMallRequest) SetChannelSupplierId(v string) *ListDistributionMallRequest {
	s.ChannelSupplierId = &v
	return s
}

func (s *ListDistributionMallRequest) SetDistributionMallId(v string) *ListDistributionMallRequest {
	s.DistributionMallId = &v
	return s
}

func (s *ListDistributionMallRequest) SetDistributionMallName(v string) *ListDistributionMallRequest {
	s.DistributionMallName = &v
	return s
}

func (s *ListDistributionMallRequest) SetDistributorId(v string) *ListDistributionMallRequest {
	s.DistributorId = &v
	return s
}

func (s *ListDistributionMallRequest) SetEndDate(v string) *ListDistributionMallRequest {
	s.EndDate = &v
	return s
}

func (s *ListDistributionMallRequest) SetPageNumber(v int32) *ListDistributionMallRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDistributionMallRequest) SetPageSize(v int32) *ListDistributionMallRequest {
	s.PageSize = &v
	return s
}

func (s *ListDistributionMallRequest) SetStartDate(v string) *ListDistributionMallRequest {
	s.StartDate = &v
	return s
}

func (s *ListDistributionMallRequest) SetTenantId(v string) *ListDistributionMallRequest {
	s.TenantId = &v
	return s
}

type ListDistributionMallResponseBody struct {
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                  `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      []*ListDistributionMallResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	PageNumber *int64                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDistributionMallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionMallResponseBody) GoString() string {
	return s.String()
}

func (s *ListDistributionMallResponseBody) SetCode(v string) *ListDistributionMallResponseBody {
	s.Code = &v
	return s
}

func (s *ListDistributionMallResponseBody) SetLogsId(v string) *ListDistributionMallResponseBody {
	s.LogsId = &v
	return s
}

func (s *ListDistributionMallResponseBody) SetMessage(v string) *ListDistributionMallResponseBody {
	s.Message = &v
	return s
}

func (s *ListDistributionMallResponseBody) SetModel(v []*ListDistributionMallResponseBodyModel) *ListDistributionMallResponseBody {
	s.Model = v
	return s
}

func (s *ListDistributionMallResponseBody) SetPageNumber(v int64) *ListDistributionMallResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDistributionMallResponseBody) SetPageSize(v int64) *ListDistributionMallResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDistributionMallResponseBody) SetRequestId(v string) *ListDistributionMallResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDistributionMallResponseBody) SetSubCode(v string) *ListDistributionMallResponseBody {
	s.SubCode = &v
	return s
}

func (s *ListDistributionMallResponseBody) SetSubMessage(v string) *ListDistributionMallResponseBody {
	s.SubMessage = &v
	return s
}

func (s *ListDistributionMallResponseBody) SetSuccess(v bool) *ListDistributionMallResponseBody {
	s.Success = &v
	return s
}

func (s *ListDistributionMallResponseBody) SetTotalCount(v int64) *ListDistributionMallResponseBody {
	s.TotalCount = &v
	return s
}

type ListDistributionMallResponseBodyModel struct {
	ChannelSupplierId    *string `json:"ChannelSupplierId,omitempty" xml:"ChannelSupplierId,omitempty"`
	DistributionMallId   *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributionMallName *string `json:"DistributionMallName,omitempty" xml:"DistributionMallName,omitempty"`
	DistributionMallType *string `json:"DistributionMallType,omitempty" xml:"DistributionMallType,omitempty"`
	EndDate              *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartDate            *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDistributionMallResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionMallResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ListDistributionMallResponseBodyModel) SetChannelSupplierId(v string) *ListDistributionMallResponseBodyModel {
	s.ChannelSupplierId = &v
	return s
}

func (s *ListDistributionMallResponseBodyModel) SetDistributionMallId(v string) *ListDistributionMallResponseBodyModel {
	s.DistributionMallId = &v
	return s
}

func (s *ListDistributionMallResponseBodyModel) SetDistributionMallName(v string) *ListDistributionMallResponseBodyModel {
	s.DistributionMallName = &v
	return s
}

func (s *ListDistributionMallResponseBodyModel) SetDistributionMallType(v string) *ListDistributionMallResponseBodyModel {
	s.DistributionMallType = &v
	return s
}

func (s *ListDistributionMallResponseBodyModel) SetEndDate(v string) *ListDistributionMallResponseBodyModel {
	s.EndDate = &v
	return s
}

func (s *ListDistributionMallResponseBodyModel) SetStartDate(v string) *ListDistributionMallResponseBodyModel {
	s.StartDate = &v
	return s
}

func (s *ListDistributionMallResponseBodyModel) SetStatus(v string) *ListDistributionMallResponseBodyModel {
	s.Status = &v
	return s
}

type ListDistributionMallResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDistributionMallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDistributionMallResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDistributionMallResponse) GoString() string {
	return s.String()
}

func (s *ListDistributionMallResponse) SetHeaders(v map[string]*string) *ListDistributionMallResponse {
	s.Headers = v
	return s
}

func (s *ListDistributionMallResponse) SetStatusCode(v int32) *ListDistributionMallResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDistributionMallResponse) SetBody(v *ListDistributionMallResponseBody) *ListDistributionMallResponse {
	s.Body = v
	return s
}

type ModifyRefund4DistributionRequest struct {
	ApplyReasonTextId      *int64                                               `json:"ApplyReasonTextId,omitempty" xml:"ApplyReasonTextId,omitempty"`
	ApplyRefundCount       *int32                                               `json:"ApplyRefundCount,omitempty" xml:"ApplyRefundCount,omitempty"`
	ApplyRefundFee         *int64                                               `json:"ApplyRefundFee,omitempty" xml:"ApplyRefundFee,omitempty"`
	BizClaimType           *int32                                               `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	DisputeId              *int64                                               `json:"DisputeId,omitempty" xml:"DisputeId,omitempty"`
	DistributorId          *string                                              `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	GoodsStatus            *int32                                               `json:"GoodsStatus,omitempty" xml:"GoodsStatus,omitempty"`
	LeaveMessage           *string                                              `json:"LeaveMessage,omitempty" xml:"LeaveMessage,omitempty"`
	LeavePictureLists      []*ModifyRefund4DistributionRequestLeavePictureLists `json:"LeavePictureLists,omitempty" xml:"LeavePictureLists,omitempty" type:"Repeated"`
	SubDistributionOrderId *string                                              `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
	TenantId               *string                                              `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ModifyRefund4DistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRefund4DistributionRequest) GoString() string {
	return s.String()
}

func (s *ModifyRefund4DistributionRequest) SetApplyReasonTextId(v int64) *ModifyRefund4DistributionRequest {
	s.ApplyReasonTextId = &v
	return s
}

func (s *ModifyRefund4DistributionRequest) SetApplyRefundCount(v int32) *ModifyRefund4DistributionRequest {
	s.ApplyRefundCount = &v
	return s
}

func (s *ModifyRefund4DistributionRequest) SetApplyRefundFee(v int64) *ModifyRefund4DistributionRequest {
	s.ApplyRefundFee = &v
	return s
}

func (s *ModifyRefund4DistributionRequest) SetBizClaimType(v int32) *ModifyRefund4DistributionRequest {
	s.BizClaimType = &v
	return s
}

func (s *ModifyRefund4DistributionRequest) SetDisputeId(v int64) *ModifyRefund4DistributionRequest {
	s.DisputeId = &v
	return s
}

func (s *ModifyRefund4DistributionRequest) SetDistributorId(v string) *ModifyRefund4DistributionRequest {
	s.DistributorId = &v
	return s
}

func (s *ModifyRefund4DistributionRequest) SetGoodsStatus(v int32) *ModifyRefund4DistributionRequest {
	s.GoodsStatus = &v
	return s
}

func (s *ModifyRefund4DistributionRequest) SetLeaveMessage(v string) *ModifyRefund4DistributionRequest {
	s.LeaveMessage = &v
	return s
}

func (s *ModifyRefund4DistributionRequest) SetLeavePictureLists(v []*ModifyRefund4DistributionRequestLeavePictureLists) *ModifyRefund4DistributionRequest {
	s.LeavePictureLists = v
	return s
}

func (s *ModifyRefund4DistributionRequest) SetSubDistributionOrderId(v string) *ModifyRefund4DistributionRequest {
	s.SubDistributionOrderId = &v
	return s
}

func (s *ModifyRefund4DistributionRequest) SetTenantId(v string) *ModifyRefund4DistributionRequest {
	s.TenantId = &v
	return s
}

type ModifyRefund4DistributionRequestLeavePictureLists struct {
	Desc    *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Picture *string `json:"Picture,omitempty" xml:"Picture,omitempty"`
}

func (s ModifyRefund4DistributionRequestLeavePictureLists) String() string {
	return tea.Prettify(s)
}

func (s ModifyRefund4DistributionRequestLeavePictureLists) GoString() string {
	return s.String()
}

func (s *ModifyRefund4DistributionRequestLeavePictureLists) SetDesc(v string) *ModifyRefund4DistributionRequestLeavePictureLists {
	s.Desc = &v
	return s
}

func (s *ModifyRefund4DistributionRequestLeavePictureLists) SetPicture(v string) *ModifyRefund4DistributionRequestLeavePictureLists {
	s.Picture = &v
	return s
}

type ModifyRefund4DistributionShrinkRequest struct {
	ApplyReasonTextId       *int64  `json:"ApplyReasonTextId,omitempty" xml:"ApplyReasonTextId,omitempty"`
	ApplyRefundCount        *int32  `json:"ApplyRefundCount,omitempty" xml:"ApplyRefundCount,omitempty"`
	ApplyRefundFee          *int64  `json:"ApplyRefundFee,omitempty" xml:"ApplyRefundFee,omitempty"`
	BizClaimType            *int32  `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	DisputeId               *int64  `json:"DisputeId,omitempty" xml:"DisputeId,omitempty"`
	DistributorId           *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	GoodsStatus             *int32  `json:"GoodsStatus,omitempty" xml:"GoodsStatus,omitempty"`
	LeaveMessage            *string `json:"LeaveMessage,omitempty" xml:"LeaveMessage,omitempty"`
	LeavePictureListsShrink *string `json:"LeavePictureLists,omitempty" xml:"LeavePictureLists,omitempty"`
	SubDistributionOrderId  *string `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
	TenantId                *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ModifyRefund4DistributionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyRefund4DistributionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyRefund4DistributionShrinkRequest) SetApplyReasonTextId(v int64) *ModifyRefund4DistributionShrinkRequest {
	s.ApplyReasonTextId = &v
	return s
}

func (s *ModifyRefund4DistributionShrinkRequest) SetApplyRefundCount(v int32) *ModifyRefund4DistributionShrinkRequest {
	s.ApplyRefundCount = &v
	return s
}

func (s *ModifyRefund4DistributionShrinkRequest) SetApplyRefundFee(v int64) *ModifyRefund4DistributionShrinkRequest {
	s.ApplyRefundFee = &v
	return s
}

func (s *ModifyRefund4DistributionShrinkRequest) SetBizClaimType(v int32) *ModifyRefund4DistributionShrinkRequest {
	s.BizClaimType = &v
	return s
}

func (s *ModifyRefund4DistributionShrinkRequest) SetDisputeId(v int64) *ModifyRefund4DistributionShrinkRequest {
	s.DisputeId = &v
	return s
}

func (s *ModifyRefund4DistributionShrinkRequest) SetDistributorId(v string) *ModifyRefund4DistributionShrinkRequest {
	s.DistributorId = &v
	return s
}

func (s *ModifyRefund4DistributionShrinkRequest) SetGoodsStatus(v int32) *ModifyRefund4DistributionShrinkRequest {
	s.GoodsStatus = &v
	return s
}

func (s *ModifyRefund4DistributionShrinkRequest) SetLeaveMessage(v string) *ModifyRefund4DistributionShrinkRequest {
	s.LeaveMessage = &v
	return s
}

func (s *ModifyRefund4DistributionShrinkRequest) SetLeavePictureListsShrink(v string) *ModifyRefund4DistributionShrinkRequest {
	s.LeavePictureListsShrink = &v
	return s
}

func (s *ModifyRefund4DistributionShrinkRequest) SetSubDistributionOrderId(v string) *ModifyRefund4DistributionShrinkRequest {
	s.SubDistributionOrderId = &v
	return s
}

func (s *ModifyRefund4DistributionShrinkRequest) SetTenantId(v string) *ModifyRefund4DistributionShrinkRequest {
	s.TenantId = &v
	return s
}

type ModifyRefund4DistributionResponseBody struct {
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                     `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *ModifyRefund4DistributionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	PageNumber *int64                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ModifyRefund4DistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyRefund4DistributionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRefund4DistributionResponseBody) SetCode(v string) *ModifyRefund4DistributionResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyRefund4DistributionResponseBody) SetLogsId(v string) *ModifyRefund4DistributionResponseBody {
	s.LogsId = &v
	return s
}

func (s *ModifyRefund4DistributionResponseBody) SetMessage(v string) *ModifyRefund4DistributionResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyRefund4DistributionResponseBody) SetModel(v *ModifyRefund4DistributionResponseBodyModel) *ModifyRefund4DistributionResponseBody {
	s.Model = v
	return s
}

func (s *ModifyRefund4DistributionResponseBody) SetPageNumber(v int64) *ModifyRefund4DistributionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ModifyRefund4DistributionResponseBody) SetPageSize(v int64) *ModifyRefund4DistributionResponseBody {
	s.PageSize = &v
	return s
}

func (s *ModifyRefund4DistributionResponseBody) SetRequestId(v string) *ModifyRefund4DistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRefund4DistributionResponseBody) SetSubCode(v string) *ModifyRefund4DistributionResponseBody {
	s.SubCode = &v
	return s
}

func (s *ModifyRefund4DistributionResponseBody) SetSubMessage(v string) *ModifyRefund4DistributionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *ModifyRefund4DistributionResponseBody) SetSuccess(v bool) *ModifyRefund4DistributionResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyRefund4DistributionResponseBody) SetTotalCount(v int64) *ModifyRefund4DistributionResponseBody {
	s.TotalCount = &v
	return s
}

type ModifyRefund4DistributionResponseBodyModel struct {
	DisputeId              *int64  `json:"DisputeId,omitempty" xml:"DisputeId,omitempty"`
	DisputeStatus          *int32  `json:"DisputeStatus,omitempty" xml:"DisputeStatus,omitempty"`
	DisputeType            *int32  `json:"DisputeType,omitempty" xml:"DisputeType,omitempty"`
	SubDistributionOrderId *string `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
}

func (s ModifyRefund4DistributionResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s ModifyRefund4DistributionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *ModifyRefund4DistributionResponseBodyModel) SetDisputeId(v int64) *ModifyRefund4DistributionResponseBodyModel {
	s.DisputeId = &v
	return s
}

func (s *ModifyRefund4DistributionResponseBodyModel) SetDisputeStatus(v int32) *ModifyRefund4DistributionResponseBodyModel {
	s.DisputeStatus = &v
	return s
}

func (s *ModifyRefund4DistributionResponseBodyModel) SetDisputeType(v int32) *ModifyRefund4DistributionResponseBodyModel {
	s.DisputeType = &v
	return s
}

func (s *ModifyRefund4DistributionResponseBodyModel) SetSubDistributionOrderId(v string) *ModifyRefund4DistributionResponseBodyModel {
	s.SubDistributionOrderId = &v
	return s
}

type ModifyRefund4DistributionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRefund4DistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRefund4DistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyRefund4DistributionResponse) GoString() string {
	return s.String()
}

func (s *ModifyRefund4DistributionResponse) SetHeaders(v map[string]*string) *ModifyRefund4DistributionResponse {
	s.Headers = v
	return s
}

func (s *ModifyRefund4DistributionResponse) SetStatusCode(v int32) *ModifyRefund4DistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRefund4DistributionResponse) SetBody(v *ModifyRefund4DistributionResponseBody) *ModifyRefund4DistributionResponse {
	s.Body = v
	return s
}

type QueryChildDivisionCodeByIdRequest struct {
	DistributorId *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	DivisionCode  *string `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryChildDivisionCodeByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryChildDivisionCodeByIdRequest) GoString() string {
	return s.String()
}

func (s *QueryChildDivisionCodeByIdRequest) SetDistributorId(v string) *QueryChildDivisionCodeByIdRequest {
	s.DistributorId = &v
	return s
}

func (s *QueryChildDivisionCodeByIdRequest) SetDivisionCode(v string) *QueryChildDivisionCodeByIdRequest {
	s.DivisionCode = &v
	return s
}

func (s *QueryChildDivisionCodeByIdRequest) SetTenantId(v string) *QueryChildDivisionCodeByIdRequest {
	s.TenantId = &v
	return s
}

type QueryChildDivisionCodeByIdResponseBody struct {
	Code       *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                      `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *QueryChildDivisionCodeByIdResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	PageNumber *int64                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string                                      `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string                                      `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryChildDivisionCodeByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryChildDivisionCodeByIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChildDivisionCodeByIdResponseBody) SetCode(v string) *QueryChildDivisionCodeByIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBody) SetLogsId(v string) *QueryChildDivisionCodeByIdResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBody) SetMessage(v string) *QueryChildDivisionCodeByIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBody) SetModel(v *QueryChildDivisionCodeByIdResponseBodyModel) *QueryChildDivisionCodeByIdResponseBody {
	s.Model = v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBody) SetPageNumber(v int64) *QueryChildDivisionCodeByIdResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBody) SetPageSize(v int64) *QueryChildDivisionCodeByIdResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBody) SetRequestId(v string) *QueryChildDivisionCodeByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBody) SetSubCode(v string) *QueryChildDivisionCodeByIdResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBody) SetSubMessage(v string) *QueryChildDivisionCodeByIdResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBody) SetSuccess(v bool) *QueryChildDivisionCodeByIdResponseBody {
	s.Success = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBody) SetTotalCount(v int64) *QueryChildDivisionCodeByIdResponseBody {
	s.TotalCount = &v
	return s
}

type QueryChildDivisionCodeByIdResponseBodyModel struct {
	DivisionList []*QueryChildDivisionCodeByIdResponseBodyModelDivisionList `json:"DivisionList,omitempty" xml:"DivisionList,omitempty" type:"Repeated"`
}

func (s QueryChildDivisionCodeByIdResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryChildDivisionCodeByIdResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryChildDivisionCodeByIdResponseBodyModel) SetDivisionList(v []*QueryChildDivisionCodeByIdResponseBodyModelDivisionList) *QueryChildDivisionCodeByIdResponseBodyModel {
	s.DivisionList = v
	return s
}

type QueryChildDivisionCodeByIdResponseBodyModelDivisionList struct {
	DivisionCode  *int64  `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
	DivisionLevel *int64  `json:"DivisionLevel,omitempty" xml:"DivisionLevel,omitempty"`
	DivisionName  *string `json:"DivisionName,omitempty" xml:"DivisionName,omitempty"`
	ParentId      *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Pinyin        *string `json:"Pinyin,omitempty" xml:"Pinyin,omitempty"`
}

func (s QueryChildDivisionCodeByIdResponseBodyModelDivisionList) String() string {
	return tea.Prettify(s)
}

func (s QueryChildDivisionCodeByIdResponseBodyModelDivisionList) GoString() string {
	return s.String()
}

func (s *QueryChildDivisionCodeByIdResponseBodyModelDivisionList) SetDivisionCode(v int64) *QueryChildDivisionCodeByIdResponseBodyModelDivisionList {
	s.DivisionCode = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBodyModelDivisionList) SetDivisionLevel(v int64) *QueryChildDivisionCodeByIdResponseBodyModelDivisionList {
	s.DivisionLevel = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBodyModelDivisionList) SetDivisionName(v string) *QueryChildDivisionCodeByIdResponseBodyModelDivisionList {
	s.DivisionName = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBodyModelDivisionList) SetParentId(v int64) *QueryChildDivisionCodeByIdResponseBodyModelDivisionList {
	s.ParentId = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponseBodyModelDivisionList) SetPinyin(v string) *QueryChildDivisionCodeByIdResponseBodyModelDivisionList {
	s.Pinyin = &v
	return s
}

type QueryChildDivisionCodeByIdResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryChildDivisionCodeByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryChildDivisionCodeByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryChildDivisionCodeByIdResponse) GoString() string {
	return s.String()
}

func (s *QueryChildDivisionCodeByIdResponse) SetHeaders(v map[string]*string) *QueryChildDivisionCodeByIdResponse {
	s.Headers = v
	return s
}

func (s *QueryChildDivisionCodeByIdResponse) SetStatusCode(v int32) *QueryChildDivisionCodeByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChildDivisionCodeByIdResponse) SetBody(v *QueryChildDivisionCodeByIdResponseBody) *QueryChildDivisionCodeByIdResponse {
	s.Body = v
	return s
}

type QueryDistributionBillDetailRequest struct {
	BillId               *string `json:"BillId,omitempty" xml:"BillId,omitempty"`
	BillPeriod           *string `json:"BillPeriod,omitempty" xml:"BillPeriod,omitempty"`
	BillStatus           *string `json:"BillStatus,omitempty" xml:"BillStatus,omitempty"`
	DistributionMallId   *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributionMallName *string `json:"DistributionMallName,omitempty" xml:"DistributionMallName,omitempty"`
	DistributorId        *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryDistributionBillDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDistributionBillDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryDistributionBillDetailRequest) SetBillId(v string) *QueryDistributionBillDetailRequest {
	s.BillId = &v
	return s
}

func (s *QueryDistributionBillDetailRequest) SetBillPeriod(v string) *QueryDistributionBillDetailRequest {
	s.BillPeriod = &v
	return s
}

func (s *QueryDistributionBillDetailRequest) SetBillStatus(v string) *QueryDistributionBillDetailRequest {
	s.BillStatus = &v
	return s
}

func (s *QueryDistributionBillDetailRequest) SetDistributionMallId(v string) *QueryDistributionBillDetailRequest {
	s.DistributionMallId = &v
	return s
}

func (s *QueryDistributionBillDetailRequest) SetDistributionMallName(v string) *QueryDistributionBillDetailRequest {
	s.DistributionMallName = &v
	return s
}

func (s *QueryDistributionBillDetailRequest) SetDistributorId(v string) *QueryDistributionBillDetailRequest {
	s.DistributorId = &v
	return s
}

func (s *QueryDistributionBillDetailRequest) SetPageNumber(v int32) *QueryDistributionBillDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryDistributionBillDetailRequest) SetPageSize(v int32) *QueryDistributionBillDetailRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDistributionBillDetailRequest) SetTenantId(v string) *QueryDistributionBillDetailRequest {
	s.TenantId = &v
	return s
}

type QueryDistributionBillDetailResponseBody struct {
	Code       *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                       `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *QueryDistributionBillDetailResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string                                       `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string                                       `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDistributionBillDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDistributionBillDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDistributionBillDetailResponseBody) SetCode(v string) *QueryDistributionBillDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDistributionBillDetailResponseBody) SetLogsId(v string) *QueryDistributionBillDetailResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryDistributionBillDetailResponseBody) SetMessage(v string) *QueryDistributionBillDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDistributionBillDetailResponseBody) SetModel(v *QueryDistributionBillDetailResponseBodyModel) *QueryDistributionBillDetailResponseBody {
	s.Model = v
	return s
}

func (s *QueryDistributionBillDetailResponseBody) SetRequestId(v string) *QueryDistributionBillDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDistributionBillDetailResponseBody) SetSubCode(v string) *QueryDistributionBillDetailResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryDistributionBillDetailResponseBody) SetSubMessage(v string) *QueryDistributionBillDetailResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryDistributionBillDetailResponseBody) SetSuccess(v bool) *QueryDistributionBillDetailResponseBody {
	s.Success = &v
	return s
}

type QueryDistributionBillDetailResponseBodyModel struct {
	Data       []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryDistributionBillDetailResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryDistributionBillDetailResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryDistributionBillDetailResponseBodyModel) SetData(v []*string) *QueryDistributionBillDetailResponseBodyModel {
	s.Data = v
	return s
}

func (s *QueryDistributionBillDetailResponseBodyModel) SetPageNumber(v int32) *QueryDistributionBillDetailResponseBodyModel {
	s.PageNumber = &v
	return s
}

func (s *QueryDistributionBillDetailResponseBodyModel) SetPageSize(v int32) *QueryDistributionBillDetailResponseBodyModel {
	s.PageSize = &v
	return s
}

func (s *QueryDistributionBillDetailResponseBodyModel) SetTotal(v int32) *QueryDistributionBillDetailResponseBodyModel {
	s.Total = &v
	return s
}

type QueryDistributionBillDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDistributionBillDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDistributionBillDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDistributionBillDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryDistributionBillDetailResponse) SetHeaders(v map[string]*string) *QueryDistributionBillDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryDistributionBillDetailResponse) SetStatusCode(v int32) *QueryDistributionBillDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDistributionBillDetailResponse) SetBody(v *QueryDistributionBillDetailResponseBody) *QueryDistributionBillDetailResponse {
	s.Body = v
	return s
}

type QueryDistributionMallRequest struct {
	DistributionMallId *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	TenantId           *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryDistributionMallRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDistributionMallRequest) GoString() string {
	return s.String()
}

func (s *QueryDistributionMallRequest) SetDistributionMallId(v string) *QueryDistributionMallRequest {
	s.DistributionMallId = &v
	return s
}

func (s *QueryDistributionMallRequest) SetTenantId(v string) *QueryDistributionMallRequest {
	s.TenantId = &v
	return s
}

type QueryDistributionMallResponseBody struct {
	BizViewData map[string]interface{}                  `json:"BizViewData,omitempty" xml:"BizViewData,omitempty"`
	Code        *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId      *string                                 `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message     *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Model       *QueryDistributionMallResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	PageNumber  *int64                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryDistributionMallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDistributionMallResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDistributionMallResponseBody) SetBizViewData(v map[string]interface{}) *QueryDistributionMallResponseBody {
	s.BizViewData = v
	return s
}

func (s *QueryDistributionMallResponseBody) SetCode(v string) *QueryDistributionMallResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDistributionMallResponseBody) SetLogsId(v string) *QueryDistributionMallResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryDistributionMallResponseBody) SetMessage(v string) *QueryDistributionMallResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDistributionMallResponseBody) SetModel(v *QueryDistributionMallResponseBodyModel) *QueryDistributionMallResponseBody {
	s.Model = v
	return s
}

func (s *QueryDistributionMallResponseBody) SetPageNumber(v int64) *QueryDistributionMallResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryDistributionMallResponseBody) SetPageSize(v int64) *QueryDistributionMallResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryDistributionMallResponseBody) SetRequestId(v string) *QueryDistributionMallResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDistributionMallResponseBody) SetSubCode(v string) *QueryDistributionMallResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryDistributionMallResponseBody) SetSubMessage(v string) *QueryDistributionMallResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryDistributionMallResponseBody) SetSuccess(v bool) *QueryDistributionMallResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDistributionMallResponseBody) SetTotalCount(v int64) *QueryDistributionMallResponseBody {
	s.TotalCount = &v
	return s
}

type QueryDistributionMallResponseBodyModel struct {
	ChannelSupplierId    *string `json:"ChannelSupplierId,omitempty" xml:"ChannelSupplierId,omitempty"`
	DistributionMallId   *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributionMallName *string `json:"DistributionMallName,omitempty" xml:"DistributionMallName,omitempty"`
	DistributionMallType *string `json:"DistributionMallType,omitempty" xml:"DistributionMallType,omitempty"`
	DistributorId        *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	EndDate              *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	StartDate            *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryDistributionMallResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryDistributionMallResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryDistributionMallResponseBodyModel) SetChannelSupplierId(v string) *QueryDistributionMallResponseBodyModel {
	s.ChannelSupplierId = &v
	return s
}

func (s *QueryDistributionMallResponseBodyModel) SetDistributionMallId(v string) *QueryDistributionMallResponseBodyModel {
	s.DistributionMallId = &v
	return s
}

func (s *QueryDistributionMallResponseBodyModel) SetDistributionMallName(v string) *QueryDistributionMallResponseBodyModel {
	s.DistributionMallName = &v
	return s
}

func (s *QueryDistributionMallResponseBodyModel) SetDistributionMallType(v string) *QueryDistributionMallResponseBodyModel {
	s.DistributionMallType = &v
	return s
}

func (s *QueryDistributionMallResponseBodyModel) SetDistributorId(v string) *QueryDistributionMallResponseBodyModel {
	s.DistributorId = &v
	return s
}

func (s *QueryDistributionMallResponseBodyModel) SetEndDate(v string) *QueryDistributionMallResponseBodyModel {
	s.EndDate = &v
	return s
}

func (s *QueryDistributionMallResponseBodyModel) SetStartDate(v string) *QueryDistributionMallResponseBodyModel {
	s.StartDate = &v
	return s
}

func (s *QueryDistributionMallResponseBodyModel) SetStatus(v string) *QueryDistributionMallResponseBodyModel {
	s.Status = &v
	return s
}

type QueryDistributionMallResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDistributionMallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDistributionMallResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDistributionMallResponse) GoString() string {
	return s.String()
}

func (s *QueryDistributionMallResponse) SetHeaders(v map[string]*string) *QueryDistributionMallResponse {
	s.Headers = v
	return s
}

func (s *QueryDistributionMallResponse) SetStatusCode(v int32) *QueryDistributionMallResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDistributionMallResponse) SetBody(v *QueryDistributionMallResponseBody) *QueryDistributionMallResponse {
	s.Body = v
	return s
}

type QueryDistributionTradeStatusRequest struct {
	DistributionSupplierId *string `json:"DistributionSupplierId,omitempty" xml:"DistributionSupplierId,omitempty"`
	DistributionTradeId    *string `json:"DistributionTradeId,omitempty" xml:"DistributionTradeId,omitempty"`
	DistributorId          *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	TenantId               *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryDistributionTradeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDistributionTradeStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDistributionTradeStatusRequest) SetDistributionSupplierId(v string) *QueryDistributionTradeStatusRequest {
	s.DistributionSupplierId = &v
	return s
}

func (s *QueryDistributionTradeStatusRequest) SetDistributionTradeId(v string) *QueryDistributionTradeStatusRequest {
	s.DistributionTradeId = &v
	return s
}

func (s *QueryDistributionTradeStatusRequest) SetDistributorId(v string) *QueryDistributionTradeStatusRequest {
	s.DistributorId = &v
	return s
}

func (s *QueryDistributionTradeStatusRequest) SetTenantId(v string) *QueryDistributionTradeStatusRequest {
	s.TenantId = &v
	return s
}

type QueryDistributionTradeStatusResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *string `json:"Model,omitempty" xml:"Model,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryDistributionTradeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDistributionTradeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDistributionTradeStatusResponseBody) SetCode(v string) *QueryDistributionTradeStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDistributionTradeStatusResponseBody) SetLogsId(v string) *QueryDistributionTradeStatusResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryDistributionTradeStatusResponseBody) SetMessage(v string) *QueryDistributionTradeStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDistributionTradeStatusResponseBody) SetModel(v string) *QueryDistributionTradeStatusResponseBody {
	s.Model = &v
	return s
}

func (s *QueryDistributionTradeStatusResponseBody) SetPageNumber(v int64) *QueryDistributionTradeStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryDistributionTradeStatusResponseBody) SetPageSize(v int64) *QueryDistributionTradeStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryDistributionTradeStatusResponseBody) SetRequestId(v string) *QueryDistributionTradeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDistributionTradeStatusResponseBody) SetSubCode(v string) *QueryDistributionTradeStatusResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryDistributionTradeStatusResponseBody) SetSubMessage(v string) *QueryDistributionTradeStatusResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryDistributionTradeStatusResponseBody) SetSuccess(v bool) *QueryDistributionTradeStatusResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDistributionTradeStatusResponseBody) SetTotalCount(v int64) *QueryDistributionTradeStatusResponseBody {
	s.TotalCount = &v
	return s
}

type QueryDistributionTradeStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDistributionTradeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDistributionTradeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDistributionTradeStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDistributionTradeStatusResponse) SetHeaders(v map[string]*string) *QueryDistributionTradeStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDistributionTradeStatusResponse) SetStatusCode(v int32) *QueryDistributionTradeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDistributionTradeStatusResponse) SetBody(v *QueryDistributionTradeStatusResponseBody) *QueryDistributionTradeStatusResponse {
	s.Body = v
	return s
}

type QueryItemDetailRequest struct {
	DistributionMallId *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributorId      *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	LmItemId           *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	TenantId           *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryItemDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryItemDetailRequest) SetDistributionMallId(v string) *QueryItemDetailRequest {
	s.DistributionMallId = &v
	return s
}

func (s *QueryItemDetailRequest) SetDistributorId(v string) *QueryItemDetailRequest {
	s.DistributorId = &v
	return s
}

func (s *QueryItemDetailRequest) SetLmItemId(v string) *QueryItemDetailRequest {
	s.LmItemId = &v
	return s
}

func (s *QueryItemDetailRequest) SetTenantId(v string) *QueryItemDetailRequest {
	s.TenantId = &v
	return s
}

type QueryItemDetailResponseBody struct {
	BizViewData map[string]interface{}            `json:"BizViewData,omitempty" xml:"BizViewData,omitempty"`
	Code        *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId      *string                           `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message     *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Model       *QueryItemDetailResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	PageNumber  *int64                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryItemDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryItemDetailResponseBody) SetBizViewData(v map[string]interface{}) *QueryItemDetailResponseBody {
	s.BizViewData = v
	return s
}

func (s *QueryItemDetailResponseBody) SetCode(v string) *QueryItemDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetLogsId(v string) *QueryItemDetailResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetMessage(v string) *QueryItemDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetModel(v *QueryItemDetailResponseBodyModel) *QueryItemDetailResponseBody {
	s.Model = v
	return s
}

func (s *QueryItemDetailResponseBody) SetPageNumber(v int64) *QueryItemDetailResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetPageSize(v int64) *QueryItemDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetRequestId(v string) *QueryItemDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetSubCode(v string) *QueryItemDetailResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetSubMessage(v string) *QueryItemDetailResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetSuccess(v bool) *QueryItemDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryItemDetailResponseBody) SetTotalCount(v int64) *QueryItemDetailResponseBody {
	s.TotalCount = &v
	return s
}

type QueryItemDetailResponseBodyModel struct {
	CanNotBeSoldCode       *string              `json:"CanNotBeSoldCode,omitempty" xml:"CanNotBeSoldCode,omitempty"`
	CanNotBeSoldMessage    *string              `json:"CanNotBeSoldMessage,omitempty" xml:"CanNotBeSoldMessage,omitempty"`
	CategoryId             *int64               `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryIds            []*int64             `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	City                   *string              `json:"City,omitempty" xml:"City,omitempty"`
	Current                *string              `json:"Current,omitempty" xml:"Current,omitempty"`
	CustomizedAttributeMap map[string]*string   `json:"CustomizedAttributeMap,omitempty" xml:"CustomizedAttributeMap,omitempty"`
	DescOption             *string              `json:"DescOption,omitempty" xml:"DescOption,omitempty"`
	DescPath               *string              `json:"DescPath,omitempty" xml:"DescPath,omitempty"`
	DistributionMallId     *string              `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	Features               map[string]*string   `json:"Features,omitempty" xml:"Features,omitempty"`
	FirstPicUrl            *string              `json:"FirstPicUrl,omitempty" xml:"FirstPicUrl,omitempty"`
	HasQuantity            *bool                `json:"HasQuantity,omitempty" xml:"HasQuantity,omitempty"`
	IforestProps           []map[string]*string `json:"IforestProps,omitempty" xml:"IforestProps,omitempty" type:"Repeated"`
	InvoiceType            *int32               `json:"InvoiceType,omitempty" xml:"InvoiceType,omitempty"`
	IsCanSell              *bool                `json:"IsCanSell,omitempty" xml:"IsCanSell,omitempty"`
	IsSellerPayPostfee     *bool                `json:"IsSellerPayPostfee,omitempty" xml:"IsSellerPayPostfee,omitempty"`
	ItemId                 *int64               `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemImages             []*string            `json:"ItemImages,omitempty" xml:"ItemImages,omitempty" type:"Repeated"`
	ItemTitle              *string              `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	ItemTotalSimpleValue   *string              `json:"ItemTotalSimpleValue,omitempty" xml:"ItemTotalSimpleValue,omitempty"`
	ItemTotalValue         *int32               `json:"ItemTotalValue,omitempty" xml:"ItemTotalValue,omitempty"`
	LmItemAttributeMap     map[string]*string   `json:"LmItemAttributeMap,omitempty" xml:"LmItemAttributeMap,omitempty"`
	LmItemCategory         *string              `json:"LmItemCategory,omitempty" xml:"LmItemCategory,omitempty"`
	LmItemId               *string              `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	MainPicUrl             *string              `json:"MainPicUrl,omitempty" xml:"MainPicUrl,omitempty"`
	MinPrice               *int64               `json:"MinPrice,omitempty" xml:"MinPrice,omitempty"`
	Properties             map[string][]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	Prov                   *string              `json:"Prov,omitempty" xml:"Prov,omitempty"`
	Quantity               *int32               `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ReservedPrice          *int64               `json:"ReservedPrice,omitempty" xml:"ReservedPrice,omitempty"`
	SimpleQuantity         *string              `json:"SimpleQuantity,omitempty" xml:"SimpleQuantity,omitempty"`
	// sku list
	SkuModels        []*QueryItemDetailResponseBodyModelSkuModels    `json:"SkuModels,omitempty" xml:"SkuModels,omitempty" type:"Repeated"`
	SkuPropertys     []*QueryItemDetailResponseBodyModelSkuPropertys `json:"SkuPropertys,omitempty" xml:"SkuPropertys,omitempty" type:"Repeated"`
	ThirdPartyItemId *string                                         `json:"ThirdPartyItemId,omitempty" xml:"ThirdPartyItemId,omitempty"`
	ThirdPartyName   *string                                         `json:"ThirdPartyName,omitempty" xml:"ThirdPartyName,omitempty"`
	VideoPicUrl      *string                                         `json:"VideoPicUrl,omitempty" xml:"VideoPicUrl,omitempty"`
	VideoUrl         *string                                         `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	VirtualItemType  *string                                         `json:"VirtualItemType,omitempty" xml:"VirtualItemType,omitempty"`
}

func (s QueryItemDetailResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryItemDetailResponseBodyModel) SetCanNotBeSoldCode(v string) *QueryItemDetailResponseBodyModel {
	s.CanNotBeSoldCode = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetCanNotBeSoldMessage(v string) *QueryItemDetailResponseBodyModel {
	s.CanNotBeSoldMessage = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetCategoryId(v int64) *QueryItemDetailResponseBodyModel {
	s.CategoryId = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetCategoryIds(v []*int64) *QueryItemDetailResponseBodyModel {
	s.CategoryIds = v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetCity(v string) *QueryItemDetailResponseBodyModel {
	s.City = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetCurrent(v string) *QueryItemDetailResponseBodyModel {
	s.Current = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetCustomizedAttributeMap(v map[string]*string) *QueryItemDetailResponseBodyModel {
	s.CustomizedAttributeMap = v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetDescOption(v string) *QueryItemDetailResponseBodyModel {
	s.DescOption = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetDescPath(v string) *QueryItemDetailResponseBodyModel {
	s.DescPath = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetDistributionMallId(v string) *QueryItemDetailResponseBodyModel {
	s.DistributionMallId = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetFeatures(v map[string]*string) *QueryItemDetailResponseBodyModel {
	s.Features = v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetFirstPicUrl(v string) *QueryItemDetailResponseBodyModel {
	s.FirstPicUrl = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetHasQuantity(v bool) *QueryItemDetailResponseBodyModel {
	s.HasQuantity = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetIforestProps(v []map[string]*string) *QueryItemDetailResponseBodyModel {
	s.IforestProps = v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetInvoiceType(v int32) *QueryItemDetailResponseBodyModel {
	s.InvoiceType = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetIsCanSell(v bool) *QueryItemDetailResponseBodyModel {
	s.IsCanSell = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetIsSellerPayPostfee(v bool) *QueryItemDetailResponseBodyModel {
	s.IsSellerPayPostfee = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetItemId(v int64) *QueryItemDetailResponseBodyModel {
	s.ItemId = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetItemImages(v []*string) *QueryItemDetailResponseBodyModel {
	s.ItemImages = v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetItemTitle(v string) *QueryItemDetailResponseBodyModel {
	s.ItemTitle = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetItemTotalSimpleValue(v string) *QueryItemDetailResponseBodyModel {
	s.ItemTotalSimpleValue = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetItemTotalValue(v int32) *QueryItemDetailResponseBodyModel {
	s.ItemTotalValue = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetLmItemAttributeMap(v map[string]*string) *QueryItemDetailResponseBodyModel {
	s.LmItemAttributeMap = v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetLmItemCategory(v string) *QueryItemDetailResponseBodyModel {
	s.LmItemCategory = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetLmItemId(v string) *QueryItemDetailResponseBodyModel {
	s.LmItemId = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetMainPicUrl(v string) *QueryItemDetailResponseBodyModel {
	s.MainPicUrl = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetMinPrice(v int64) *QueryItemDetailResponseBodyModel {
	s.MinPrice = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetProperties(v map[string][]*string) *QueryItemDetailResponseBodyModel {
	s.Properties = v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetProv(v string) *QueryItemDetailResponseBodyModel {
	s.Prov = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetQuantity(v int32) *QueryItemDetailResponseBodyModel {
	s.Quantity = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetReservedPrice(v int64) *QueryItemDetailResponseBodyModel {
	s.ReservedPrice = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetSimpleQuantity(v string) *QueryItemDetailResponseBodyModel {
	s.SimpleQuantity = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetSkuModels(v []*QueryItemDetailResponseBodyModelSkuModels) *QueryItemDetailResponseBodyModel {
	s.SkuModels = v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetSkuPropertys(v []*QueryItemDetailResponseBodyModelSkuPropertys) *QueryItemDetailResponseBodyModel {
	s.SkuPropertys = v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetThirdPartyItemId(v string) *QueryItemDetailResponseBodyModel {
	s.ThirdPartyItemId = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetThirdPartyName(v string) *QueryItemDetailResponseBodyModel {
	s.ThirdPartyName = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetVideoPicUrl(v string) *QueryItemDetailResponseBodyModel {
	s.VideoPicUrl = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetVideoUrl(v string) *QueryItemDetailResponseBodyModel {
	s.VideoUrl = &v
	return s
}

func (s *QueryItemDetailResponseBodyModel) SetVirtualItemType(v string) *QueryItemDetailResponseBodyModel {
	s.VirtualItemType = &v
	return s
}

type QueryItemDetailResponseBodyModelSkuModels struct {
	CanNotBeSoldCode       *string            `json:"CanNotBeSoldCode,omitempty" xml:"CanNotBeSoldCode,omitempty"`
	CanNotBeSoldMessage    *string            `json:"CanNotBeSoldMessage,omitempty" xml:"CanNotBeSoldMessage,omitempty"`
	CustomizedAttributeMap map[string]*string `json:"CustomizedAttributeMap,omitempty" xml:"CustomizedAttributeMap,omitempty"`
	DistributionMallId     *string            `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	ExtJson                *string            `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	HasQuantity            *bool              `json:"HasQuantity,omitempty" xml:"HasQuantity,omitempty"`
	InvoiceType            *int32             `json:"InvoiceType,omitempty" xml:"InvoiceType,omitempty"`
	ItemId                 *int64             `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId               *string            `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	LmSkuAttributeMap      map[string]*string `json:"LmSkuAttributeMap,omitempty" xml:"LmSkuAttributeMap,omitempty"`
	PriceCent              *int64             `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	Quantity               *int32             `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ReservedPrice          *int64             `json:"ReservedPrice,omitempty" xml:"ReservedPrice,omitempty"`
	SimpleQuantity         *string            `json:"SimpleQuantity,omitempty" xml:"SimpleQuantity,omitempty"`
	SkuId                  *int64             `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuPicUrl              *string            `json:"SkuPicUrl,omitempty" xml:"SkuPicUrl,omitempty"`
	SkuPvs                 *string            `json:"SkuPvs,omitempty" xml:"SkuPvs,omitempty"`
	SkuTitle               *string            `json:"SkuTitle,omitempty" xml:"SkuTitle,omitempty"`
	Status                 *int32             `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryItemDetailResponseBodyModelSkuModels) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailResponseBodyModelSkuModels) GoString() string {
	return s.String()
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetCanNotBeSoldCode(v string) *QueryItemDetailResponseBodyModelSkuModels {
	s.CanNotBeSoldCode = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetCanNotBeSoldMessage(v string) *QueryItemDetailResponseBodyModelSkuModels {
	s.CanNotBeSoldMessage = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetCustomizedAttributeMap(v map[string]*string) *QueryItemDetailResponseBodyModelSkuModels {
	s.CustomizedAttributeMap = v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetDistributionMallId(v string) *QueryItemDetailResponseBodyModelSkuModels {
	s.DistributionMallId = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetExtJson(v string) *QueryItemDetailResponseBodyModelSkuModels {
	s.ExtJson = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetHasQuantity(v bool) *QueryItemDetailResponseBodyModelSkuModels {
	s.HasQuantity = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetInvoiceType(v int32) *QueryItemDetailResponseBodyModelSkuModels {
	s.InvoiceType = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetItemId(v int64) *QueryItemDetailResponseBodyModelSkuModels {
	s.ItemId = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetLmItemId(v string) *QueryItemDetailResponseBodyModelSkuModels {
	s.LmItemId = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetLmSkuAttributeMap(v map[string]*string) *QueryItemDetailResponseBodyModelSkuModels {
	s.LmSkuAttributeMap = v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetPriceCent(v int64) *QueryItemDetailResponseBodyModelSkuModels {
	s.PriceCent = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetQuantity(v int32) *QueryItemDetailResponseBodyModelSkuModels {
	s.Quantity = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetReservedPrice(v int64) *QueryItemDetailResponseBodyModelSkuModels {
	s.ReservedPrice = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetSimpleQuantity(v string) *QueryItemDetailResponseBodyModelSkuModels {
	s.SimpleQuantity = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetSkuId(v int64) *QueryItemDetailResponseBodyModelSkuModels {
	s.SkuId = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetSkuPicUrl(v string) *QueryItemDetailResponseBodyModelSkuModels {
	s.SkuPicUrl = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetSkuPvs(v string) *QueryItemDetailResponseBodyModelSkuModels {
	s.SkuPvs = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetSkuTitle(v string) *QueryItemDetailResponseBodyModelSkuModels {
	s.SkuTitle = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuModels) SetStatus(v int32) *QueryItemDetailResponseBodyModelSkuModels {
	s.Status = &v
	return s
}

type QueryItemDetailResponseBodyModelSkuPropertys struct {
	Id     *int64                                                `json:"Id,omitempty" xml:"Id,omitempty"`
	Text   *string                                               `json:"Text,omitempty" xml:"Text,omitempty"`
	Values []*QueryItemDetailResponseBodyModelSkuPropertysValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryItemDetailResponseBodyModelSkuPropertys) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailResponseBodyModelSkuPropertys) GoString() string {
	return s.String()
}

func (s *QueryItemDetailResponseBodyModelSkuPropertys) SetId(v int64) *QueryItemDetailResponseBodyModelSkuPropertys {
	s.Id = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuPropertys) SetText(v string) *QueryItemDetailResponseBodyModelSkuPropertys {
	s.Text = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuPropertys) SetValues(v []*QueryItemDetailResponseBodyModelSkuPropertysValues) *QueryItemDetailResponseBodyModelSkuPropertys {
	s.Values = v
	return s
}

type QueryItemDetailResponseBodyModelSkuPropertysValues struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s QueryItemDetailResponseBodyModelSkuPropertysValues) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailResponseBodyModelSkuPropertysValues) GoString() string {
	return s.String()
}

func (s *QueryItemDetailResponseBodyModelSkuPropertysValues) SetId(v int64) *QueryItemDetailResponseBodyModelSkuPropertysValues {
	s.Id = &v
	return s
}

func (s *QueryItemDetailResponseBodyModelSkuPropertysValues) SetText(v string) *QueryItemDetailResponseBodyModelSkuPropertysValues {
	s.Text = &v
	return s
}

type QueryItemDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryItemDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryItemDetailResponse) SetStatusCode(v int32) *QueryItemDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryItemDetailResponse) SetBody(v *QueryItemDetailResponseBody) *QueryItemDetailResponse {
	s.Body = v
	return s
}

type QueryItemDetailWithDivisionRequest struct {
	DistributionMallId *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributorId      *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	DivisionCode       *string `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
	LmItemId           *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	TenantId           *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryItemDetailWithDivisionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailWithDivisionRequest) GoString() string {
	return s.String()
}

func (s *QueryItemDetailWithDivisionRequest) SetDistributionMallId(v string) *QueryItemDetailWithDivisionRequest {
	s.DistributionMallId = &v
	return s
}

func (s *QueryItemDetailWithDivisionRequest) SetDistributorId(v string) *QueryItemDetailWithDivisionRequest {
	s.DistributorId = &v
	return s
}

func (s *QueryItemDetailWithDivisionRequest) SetDivisionCode(v string) *QueryItemDetailWithDivisionRequest {
	s.DivisionCode = &v
	return s
}

func (s *QueryItemDetailWithDivisionRequest) SetLmItemId(v string) *QueryItemDetailWithDivisionRequest {
	s.LmItemId = &v
	return s
}

func (s *QueryItemDetailWithDivisionRequest) SetTenantId(v string) *QueryItemDetailWithDivisionRequest {
	s.TenantId = &v
	return s
}

type QueryItemDetailWithDivisionResponseBody struct {
	Code       *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                       `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *QueryItemDetailWithDivisionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	PageNumber *int64                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryItemDetailWithDivisionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailWithDivisionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryItemDetailWithDivisionResponseBody) SetCode(v string) *QueryItemDetailWithDivisionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBody) SetLogsId(v string) *QueryItemDetailWithDivisionResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBody) SetMessage(v string) *QueryItemDetailWithDivisionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBody) SetModel(v *QueryItemDetailWithDivisionResponseBodyModel) *QueryItemDetailWithDivisionResponseBody {
	s.Model = v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBody) SetPageNumber(v int64) *QueryItemDetailWithDivisionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBody) SetPageSize(v int64) *QueryItemDetailWithDivisionResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBody) SetRequestId(v string) *QueryItemDetailWithDivisionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBody) SetSubCode(v string) *QueryItemDetailWithDivisionResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBody) SetSubMessage(v string) *QueryItemDetailWithDivisionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBody) SetSuccess(v bool) *QueryItemDetailWithDivisionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBody) SetTotalCount(v int64) *QueryItemDetailWithDivisionResponseBody {
	s.TotalCount = &v
	return s
}

type QueryItemDetailWithDivisionResponseBodyModel struct {
	CanNotBeSoldCode       *string              `json:"CanNotBeSoldCode,omitempty" xml:"CanNotBeSoldCode,omitempty"`
	CanNotBeSoldMassage    *string              `json:"CanNotBeSoldMassage,omitempty" xml:"CanNotBeSoldMassage,omitempty"`
	CanSell                *bool                `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	CategoryId             *int64               `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryIds            []*int64             `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	City                   *string              `json:"City,omitempty" xml:"City,omitempty"`
	Current                *string              `json:"Current,omitempty" xml:"Current,omitempty"`
	CustomizedAttributeMap map[string]*string   `json:"CustomizedAttributeMap,omitempty" xml:"CustomizedAttributeMap,omitempty"`
	DescOption             *string              `json:"DescOption,omitempty" xml:"DescOption,omitempty"`
	DescPath               *string              `json:"DescPath,omitempty" xml:"DescPath,omitempty"`
	DistributionMallId     *string              `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	Features               map[string]*string   `json:"Features,omitempty" xml:"Features,omitempty"`
	FirstPicUrl            *string              `json:"FirstPicUrl,omitempty" xml:"FirstPicUrl,omitempty"`
	HasQuantity            *bool                `json:"HasQuantity,omitempty" xml:"HasQuantity,omitempty"`
	IforestProps           []map[string]*string `json:"IforestProps,omitempty" xml:"IforestProps,omitempty" type:"Repeated"`
	InvoiceType            *int32               `json:"InvoiceType,omitempty" xml:"InvoiceType,omitempty"`
	ItemId                 *int64               `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemImages             []*string            `json:"ItemImages,omitempty" xml:"ItemImages,omitempty" type:"Repeated"`
	ItemTitle              *string              `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	ItemTotalSimpleValue   *string              `json:"ItemTotalSimpleValue,omitempty" xml:"ItemTotalSimpleValue,omitempty"`
	ItemTotalValue         *int32               `json:"ItemTotalValue,omitempty" xml:"ItemTotalValue,omitempty"`
	LmItemAttributeMap     map[string]*string   `json:"LmItemAttributeMap,omitempty" xml:"LmItemAttributeMap,omitempty"`
	LmItemCategory         *string              `json:"LmItemCategory,omitempty" xml:"LmItemCategory,omitempty"`
	LmItemId               *string              `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	MainPicUrl             *string              `json:"MainPicUrl,omitempty" xml:"MainPicUrl,omitempty"`
	MinPrice               *int64               `json:"MinPrice,omitempty" xml:"MinPrice,omitempty"`
	Properties             map[string][]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	Prov                   *string              `json:"Prov,omitempty" xml:"Prov,omitempty"`
	Quantity               *int32               `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ReservePrice           *int64               `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	SecuredTransactions    *int32               `json:"SecuredTransactions,omitempty" xml:"SecuredTransactions,omitempty"`
	SellerPayPostfee       *bool                `json:"SellerPayPostfee,omitempty" xml:"SellerPayPostfee,omitempty"`
	SimpleQuantity         *string              `json:"SimpleQuantity,omitempty" xml:"SimpleQuantity,omitempty"`
	// sku list
	SkuModels        []*QueryItemDetailWithDivisionResponseBodyModelSkuModels    `json:"SkuModels,omitempty" xml:"SkuModels,omitempty" type:"Repeated"`
	SkuPropertys     []*QueryItemDetailWithDivisionResponseBodyModelSkuPropertys `json:"SkuPropertys,omitempty" xml:"SkuPropertys,omitempty" type:"Repeated"`
	ThirdPartyItemId *string                                                     `json:"ThirdPartyItemId,omitempty" xml:"ThirdPartyItemId,omitempty"`
	ThirdPartyName   *string                                                     `json:"ThirdPartyName,omitempty" xml:"ThirdPartyName,omitempty"`
	UserType         *int32                                                      `json:"UserType,omitempty" xml:"UserType,omitempty"`
	VideoPicUrl      *string                                                     `json:"VideoPicUrl,omitempty" xml:"VideoPicUrl,omitempty"`
	VideoUrl         *string                                                     `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	VirtualItemType  *string                                                     `json:"VirtualItemType,omitempty" xml:"VirtualItemType,omitempty"`
}

func (s QueryItemDetailWithDivisionResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailWithDivisionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetCanNotBeSoldCode(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.CanNotBeSoldCode = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetCanNotBeSoldMassage(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.CanNotBeSoldMassage = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetCanSell(v bool) *QueryItemDetailWithDivisionResponseBodyModel {
	s.CanSell = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetCategoryId(v int64) *QueryItemDetailWithDivisionResponseBodyModel {
	s.CategoryId = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetCategoryIds(v []*int64) *QueryItemDetailWithDivisionResponseBodyModel {
	s.CategoryIds = v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetCity(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.City = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetCurrent(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.Current = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetCustomizedAttributeMap(v map[string]*string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.CustomizedAttributeMap = v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetDescOption(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.DescOption = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetDescPath(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.DescPath = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetDistributionMallId(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.DistributionMallId = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetFeatures(v map[string]*string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.Features = v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetFirstPicUrl(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.FirstPicUrl = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetHasQuantity(v bool) *QueryItemDetailWithDivisionResponseBodyModel {
	s.HasQuantity = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetIforestProps(v []map[string]*string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.IforestProps = v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetInvoiceType(v int32) *QueryItemDetailWithDivisionResponseBodyModel {
	s.InvoiceType = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetItemId(v int64) *QueryItemDetailWithDivisionResponseBodyModel {
	s.ItemId = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetItemImages(v []*string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.ItemImages = v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetItemTitle(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.ItemTitle = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetItemTotalSimpleValue(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.ItemTotalSimpleValue = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetItemTotalValue(v int32) *QueryItemDetailWithDivisionResponseBodyModel {
	s.ItemTotalValue = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetLmItemAttributeMap(v map[string]*string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.LmItemAttributeMap = v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetLmItemCategory(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.LmItemCategory = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetLmItemId(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.LmItemId = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetMainPicUrl(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.MainPicUrl = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetMinPrice(v int64) *QueryItemDetailWithDivisionResponseBodyModel {
	s.MinPrice = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetProperties(v map[string][]*string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.Properties = v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetProv(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.Prov = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetQuantity(v int32) *QueryItemDetailWithDivisionResponseBodyModel {
	s.Quantity = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetReservePrice(v int64) *QueryItemDetailWithDivisionResponseBodyModel {
	s.ReservePrice = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetSecuredTransactions(v int32) *QueryItemDetailWithDivisionResponseBodyModel {
	s.SecuredTransactions = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetSellerPayPostfee(v bool) *QueryItemDetailWithDivisionResponseBodyModel {
	s.SellerPayPostfee = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetSimpleQuantity(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.SimpleQuantity = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetSkuModels(v []*QueryItemDetailWithDivisionResponseBodyModelSkuModels) *QueryItemDetailWithDivisionResponseBodyModel {
	s.SkuModels = v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetSkuPropertys(v []*QueryItemDetailWithDivisionResponseBodyModelSkuPropertys) *QueryItemDetailWithDivisionResponseBodyModel {
	s.SkuPropertys = v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetThirdPartyItemId(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.ThirdPartyItemId = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetThirdPartyName(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.ThirdPartyName = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetUserType(v int32) *QueryItemDetailWithDivisionResponseBodyModel {
	s.UserType = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetVideoPicUrl(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.VideoPicUrl = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetVideoUrl(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.VideoUrl = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModel) SetVirtualItemType(v string) *QueryItemDetailWithDivisionResponseBodyModel {
	s.VirtualItemType = &v
	return s
}

type QueryItemDetailWithDivisionResponseBodyModelSkuModels struct {
	CanNotBeSoldCode       *string            `json:"CanNotBeSoldCode,omitempty" xml:"CanNotBeSoldCode,omitempty"`
	CanNotBeSoldMassage    *string            `json:"CanNotBeSoldMassage,omitempty" xml:"CanNotBeSoldMassage,omitempty"`
	CustomizedAttributeMap map[string]*string `json:"CustomizedAttributeMap,omitempty" xml:"CustomizedAttributeMap,omitempty"`
	DistributionMallId     *string            `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	ExtJson                *string            `json:"ExtJson,omitempty" xml:"ExtJson,omitempty"`
	HasQuantity            *bool              `json:"HasQuantity,omitempty" xml:"HasQuantity,omitempty"`
	InvoiceType            *int32             `json:"InvoiceType,omitempty" xml:"InvoiceType,omitempty"`
	ItemId                 *int64             `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId               *string            `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	LmSkuAttributeMap      map[string]*string `json:"LmSkuAttributeMap,omitempty" xml:"LmSkuAttributeMap,omitempty"`
	PriceCent              *int64             `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	Quantity               *int32             `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ReservePrice           *int64             `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	SimpleQuantity         *string            `json:"SimpleQuantity,omitempty" xml:"SimpleQuantity,omitempty"`
	SkuId                  *int64             `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuPicUrl              *string            `json:"SkuPicUrl,omitempty" xml:"SkuPicUrl,omitempty"`
	SkuPvs                 *string            `json:"SkuPvs,omitempty" xml:"SkuPvs,omitempty"`
	SkuTitle               *string            `json:"SkuTitle,omitempty" xml:"SkuTitle,omitempty"`
	Status                 *int32             `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierPrice          *int64             `json:"SupplierPrice,omitempty" xml:"SupplierPrice,omitempty"`
}

func (s QueryItemDetailWithDivisionResponseBodyModelSkuModels) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailWithDivisionResponseBodyModelSkuModels) GoString() string {
	return s.String()
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetCanNotBeSoldCode(v string) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.CanNotBeSoldCode = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetCanNotBeSoldMassage(v string) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.CanNotBeSoldMassage = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetCustomizedAttributeMap(v map[string]*string) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.CustomizedAttributeMap = v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetDistributionMallId(v string) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.DistributionMallId = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetExtJson(v string) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.ExtJson = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetHasQuantity(v bool) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.HasQuantity = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetInvoiceType(v int32) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.InvoiceType = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetItemId(v int64) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.ItemId = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetLmItemId(v string) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.LmItemId = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetLmSkuAttributeMap(v map[string]*string) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.LmSkuAttributeMap = v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetPriceCent(v int64) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.PriceCent = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetQuantity(v int32) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.Quantity = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetReservePrice(v int64) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.ReservePrice = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetSimpleQuantity(v string) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.SimpleQuantity = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetSkuId(v int64) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.SkuId = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetSkuPicUrl(v string) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.SkuPicUrl = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetSkuPvs(v string) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.SkuPvs = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetSkuTitle(v string) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.SkuTitle = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetStatus(v int32) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.Status = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuModels) SetSupplierPrice(v int64) *QueryItemDetailWithDivisionResponseBodyModelSkuModels {
	s.SupplierPrice = &v
	return s
}

type QueryItemDetailWithDivisionResponseBodyModelSkuPropertys struct {
	Id     *int64                                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	Text   *string                                                           `json:"Text,omitempty" xml:"Text,omitempty"`
	Values []*QueryItemDetailWithDivisionResponseBodyModelSkuPropertysValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryItemDetailWithDivisionResponseBodyModelSkuPropertys) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailWithDivisionResponseBodyModelSkuPropertys) GoString() string {
	return s.String()
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuPropertys) SetId(v int64) *QueryItemDetailWithDivisionResponseBodyModelSkuPropertys {
	s.Id = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuPropertys) SetText(v string) *QueryItemDetailWithDivisionResponseBodyModelSkuPropertys {
	s.Text = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuPropertys) SetValues(v []*QueryItemDetailWithDivisionResponseBodyModelSkuPropertysValues) *QueryItemDetailWithDivisionResponseBodyModelSkuPropertys {
	s.Values = v
	return s
}

type QueryItemDetailWithDivisionResponseBodyModelSkuPropertysValues struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s QueryItemDetailWithDivisionResponseBodyModelSkuPropertysValues) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailWithDivisionResponseBodyModelSkuPropertysValues) GoString() string {
	return s.String()
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuPropertysValues) SetId(v int64) *QueryItemDetailWithDivisionResponseBodyModelSkuPropertysValues {
	s.Id = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponseBodyModelSkuPropertysValues) SetText(v string) *QueryItemDetailWithDivisionResponseBodyModelSkuPropertysValues {
	s.Text = &v
	return s
}

type QueryItemDetailWithDivisionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryItemDetailWithDivisionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryItemDetailWithDivisionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryItemDetailWithDivisionResponse) GoString() string {
	return s.String()
}

func (s *QueryItemDetailWithDivisionResponse) SetHeaders(v map[string]*string) *QueryItemDetailWithDivisionResponse {
	s.Headers = v
	return s
}

func (s *QueryItemDetailWithDivisionResponse) SetStatusCode(v int32) *QueryItemDetailWithDivisionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryItemDetailWithDivisionResponse) SetBody(v *QueryItemDetailWithDivisionResponseBody) *QueryItemDetailWithDivisionResponse {
	s.Body = v
	return s
}

type QueryItemGuideRetailPriceRequest struct {
	DistributionMallId *string   `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributorId      *string   `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	LmItemIds          []*string `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty" type:"Repeated"`
	TenantId           *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryItemGuideRetailPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemGuideRetailPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryItemGuideRetailPriceRequest) SetDistributionMallId(v string) *QueryItemGuideRetailPriceRequest {
	s.DistributionMallId = &v
	return s
}

func (s *QueryItemGuideRetailPriceRequest) SetDistributorId(v string) *QueryItemGuideRetailPriceRequest {
	s.DistributorId = &v
	return s
}

func (s *QueryItemGuideRetailPriceRequest) SetLmItemIds(v []*string) *QueryItemGuideRetailPriceRequest {
	s.LmItemIds = v
	return s
}

func (s *QueryItemGuideRetailPriceRequest) SetTenantId(v string) *QueryItemGuideRetailPriceRequest {
	s.TenantId = &v
	return s
}

type QueryItemGuideRetailPriceShrinkRequest struct {
	DistributionMallId *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributorId      *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	LmItemIdsShrink    *string `json:"LmItemIds,omitempty" xml:"LmItemIds,omitempty"`
	TenantId           *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryItemGuideRetailPriceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemGuideRetailPriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryItemGuideRetailPriceShrinkRequest) SetDistributionMallId(v string) *QueryItemGuideRetailPriceShrinkRequest {
	s.DistributionMallId = &v
	return s
}

func (s *QueryItemGuideRetailPriceShrinkRequest) SetDistributorId(v string) *QueryItemGuideRetailPriceShrinkRequest {
	s.DistributorId = &v
	return s
}

func (s *QueryItemGuideRetailPriceShrinkRequest) SetLmItemIdsShrink(v string) *QueryItemGuideRetailPriceShrinkRequest {
	s.LmItemIdsShrink = &v
	return s
}

func (s *QueryItemGuideRetailPriceShrinkRequest) SetTenantId(v string) *QueryItemGuideRetailPriceShrinkRequest {
	s.TenantId = &v
	return s
}

type QueryItemGuideRetailPriceResponseBody struct {
	Code       *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                       `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      []*QueryItemGuideRetailPriceResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string                                       `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string                                       `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryItemGuideRetailPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryItemGuideRetailPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryItemGuideRetailPriceResponseBody) SetCode(v string) *QueryItemGuideRetailPriceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBody) SetLogsId(v string) *QueryItemGuideRetailPriceResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBody) SetMessage(v string) *QueryItemGuideRetailPriceResponseBody {
	s.Message = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBody) SetModel(v []*QueryItemGuideRetailPriceResponseBodyModel) *QueryItemGuideRetailPriceResponseBody {
	s.Model = v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBody) SetRequestId(v string) *QueryItemGuideRetailPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBody) SetSubCode(v string) *QueryItemGuideRetailPriceResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBody) SetSubMessage(v string) *QueryItemGuideRetailPriceResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBody) SetSuccess(v bool) *QueryItemGuideRetailPriceResponseBody {
	s.Success = &v
	return s
}

type QueryItemGuideRetailPriceResponseBodyModel struct {
	DistributionMallId       *string                                                `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	GuideRetailPriceScope    *string                                                `json:"GuideRetailPriceScope,omitempty" xml:"GuideRetailPriceScope,omitempty"`
	ItemId                   *int64                                                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemTitle                *string                                                `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	LmItemId                 *string                                                `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	LowGuideRetailPriceScope *string                                                `json:"LowGuideRetailPriceScope,omitempty" xml:"LowGuideRetailPriceScope,omitempty"`
	ReservedPrice            *int64                                                 `json:"ReservedPrice,omitempty" xml:"ReservedPrice,omitempty"`
	ReservedPriceScope       *string                                                `json:"ReservedPriceScope,omitempty" xml:"ReservedPriceScope,omitempty"`
	SkuModels                []*QueryItemGuideRetailPriceResponseBodyModelSkuModels `json:"SkuModels,omitempty" xml:"SkuModels,omitempty" type:"Repeated"`
}

func (s QueryItemGuideRetailPriceResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryItemGuideRetailPriceResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryItemGuideRetailPriceResponseBodyModel) SetDistributionMallId(v string) *QueryItemGuideRetailPriceResponseBodyModel {
	s.DistributionMallId = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModel) SetGuideRetailPriceScope(v string) *QueryItemGuideRetailPriceResponseBodyModel {
	s.GuideRetailPriceScope = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModel) SetItemId(v int64) *QueryItemGuideRetailPriceResponseBodyModel {
	s.ItemId = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModel) SetItemTitle(v string) *QueryItemGuideRetailPriceResponseBodyModel {
	s.ItemTitle = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModel) SetLmItemId(v string) *QueryItemGuideRetailPriceResponseBodyModel {
	s.LmItemId = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModel) SetLowGuideRetailPriceScope(v string) *QueryItemGuideRetailPriceResponseBodyModel {
	s.LowGuideRetailPriceScope = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModel) SetReservedPrice(v int64) *QueryItemGuideRetailPriceResponseBodyModel {
	s.ReservedPrice = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModel) SetReservedPriceScope(v string) *QueryItemGuideRetailPriceResponseBodyModel {
	s.ReservedPriceScope = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModel) SetSkuModels(v []*QueryItemGuideRetailPriceResponseBodyModelSkuModels) *QueryItemGuideRetailPriceResponseBodyModel {
	s.SkuModels = v
	return s
}

type QueryItemGuideRetailPriceResponseBodyModelSkuModels struct {
	DistributionMallId  *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	GuideRetailPrice    *int64  `json:"GuideRetailPrice,omitempty" xml:"GuideRetailPrice,omitempty"`
	ItemId              *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId            *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	LowGuideRetailPrice *int64  `json:"LowGuideRetailPrice,omitempty" xml:"LowGuideRetailPrice,omitempty"`
	PriceCent           *int64  `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	ReservedPrice       *int64  `json:"ReservedPrice,omitempty" xml:"ReservedPrice,omitempty"`
	SkuId               *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuTitle            *string `json:"SkuTitle,omitempty" xml:"SkuTitle,omitempty"`
	Status              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryItemGuideRetailPriceResponseBodyModelSkuModels) String() string {
	return tea.Prettify(s)
}

func (s QueryItemGuideRetailPriceResponseBodyModelSkuModels) GoString() string {
	return s.String()
}

func (s *QueryItemGuideRetailPriceResponseBodyModelSkuModels) SetDistributionMallId(v string) *QueryItemGuideRetailPriceResponseBodyModelSkuModels {
	s.DistributionMallId = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModelSkuModels) SetGuideRetailPrice(v int64) *QueryItemGuideRetailPriceResponseBodyModelSkuModels {
	s.GuideRetailPrice = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModelSkuModels) SetItemId(v int64) *QueryItemGuideRetailPriceResponseBodyModelSkuModels {
	s.ItemId = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModelSkuModels) SetLmItemId(v string) *QueryItemGuideRetailPriceResponseBodyModelSkuModels {
	s.LmItemId = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModelSkuModels) SetLowGuideRetailPrice(v int64) *QueryItemGuideRetailPriceResponseBodyModelSkuModels {
	s.LowGuideRetailPrice = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModelSkuModels) SetPriceCent(v int64) *QueryItemGuideRetailPriceResponseBodyModelSkuModels {
	s.PriceCent = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModelSkuModels) SetReservedPrice(v int64) *QueryItemGuideRetailPriceResponseBodyModelSkuModels {
	s.ReservedPrice = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModelSkuModels) SetSkuId(v int64) *QueryItemGuideRetailPriceResponseBodyModelSkuModels {
	s.SkuId = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModelSkuModels) SetSkuTitle(v string) *QueryItemGuideRetailPriceResponseBodyModelSkuModels {
	s.SkuTitle = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponseBodyModelSkuModels) SetStatus(v int32) *QueryItemGuideRetailPriceResponseBodyModelSkuModels {
	s.Status = &v
	return s
}

type QueryItemGuideRetailPriceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryItemGuideRetailPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryItemGuideRetailPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryItemGuideRetailPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryItemGuideRetailPriceResponse) SetHeaders(v map[string]*string) *QueryItemGuideRetailPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryItemGuideRetailPriceResponse) SetStatusCode(v int32) *QueryItemGuideRetailPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryItemGuideRetailPriceResponse) SetBody(v *QueryItemGuideRetailPriceResponseBody) *QueryItemGuideRetailPriceResponse {
	s.Body = v
	return s
}

type QueryLogistics4DistributionRequest struct {
	DistributorId           *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	MainDistributionOrderId *string `json:"MainDistributionOrderId,omitempty" xml:"MainDistributionOrderId,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantId                *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryLogistics4DistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryLogistics4DistributionRequest) GoString() string {
	return s.String()
}

func (s *QueryLogistics4DistributionRequest) SetDistributorId(v string) *QueryLogistics4DistributionRequest {
	s.DistributorId = &v
	return s
}

func (s *QueryLogistics4DistributionRequest) SetMainDistributionOrderId(v string) *QueryLogistics4DistributionRequest {
	s.MainDistributionOrderId = &v
	return s
}

func (s *QueryLogistics4DistributionRequest) SetRequestId(v string) *QueryLogistics4DistributionRequest {
	s.RequestId = &v
	return s
}

func (s *QueryLogistics4DistributionRequest) SetTenantId(v string) *QueryLogistics4DistributionRequest {
	s.TenantId = &v
	return s
}

type QueryLogistics4DistributionResponseBody struct {
	Code       *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                         `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      []*QueryLogistics4DistributionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	PageNumber *int64                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryLogistics4DistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryLogistics4DistributionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLogistics4DistributionResponseBody) SetCode(v string) *QueryLogistics4DistributionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBody) SetLogsId(v string) *QueryLogistics4DistributionResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBody) SetMessage(v string) *QueryLogistics4DistributionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBody) SetModel(v []*QueryLogistics4DistributionResponseBodyModel) *QueryLogistics4DistributionResponseBody {
	s.Model = v
	return s
}

func (s *QueryLogistics4DistributionResponseBody) SetPageNumber(v int64) *QueryLogistics4DistributionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBody) SetPageSize(v int64) *QueryLogistics4DistributionResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBody) SetRequestId(v string) *QueryLogistics4DistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBody) SetSubCode(v string) *QueryLogistics4DistributionResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBody) SetSubMessage(v string) *QueryLogistics4DistributionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBody) SetSuccess(v bool) *QueryLogistics4DistributionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBody) SetTotalCount(v int64) *QueryLogistics4DistributionResponseBody {
	s.TotalCount = &v
	return s
}

type QueryLogistics4DistributionResponseBodyModel struct {
	DataProvider         *string                                                            `json:"DataProvider,omitempty" xml:"DataProvider,omitempty"`
	DataProviderTitle    *string                                                            `json:"DataProviderTitle,omitempty" xml:"DataProviderTitle,omitempty"`
	Goods                []*QueryLogistics4DistributionResponseBodyModelGoods               `json:"Goods,omitempty" xml:"Goods,omitempty" type:"Repeated"`
	LogisticsCompanyCode *string                                                            `json:"LogisticsCompanyCode,omitempty" xml:"LogisticsCompanyCode,omitempty"`
	LogisticsCompanyName *string                                                            `json:"LogisticsCompanyName,omitempty" xml:"LogisticsCompanyName,omitempty"`
	LogisticsDetailList  []*QueryLogistics4DistributionResponseBodyModelLogisticsDetailList `json:"LogisticsDetailList,omitempty" xml:"LogisticsDetailList,omitempty" type:"Repeated"`
	MailNo               *string                                                            `json:"MailNo,omitempty" xml:"MailNo,omitempty"`
}

func (s QueryLogistics4DistributionResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryLogistics4DistributionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryLogistics4DistributionResponseBodyModel) SetDataProvider(v string) *QueryLogistics4DistributionResponseBodyModel {
	s.DataProvider = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBodyModel) SetDataProviderTitle(v string) *QueryLogistics4DistributionResponseBodyModel {
	s.DataProviderTitle = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBodyModel) SetGoods(v []*QueryLogistics4DistributionResponseBodyModelGoods) *QueryLogistics4DistributionResponseBodyModel {
	s.Goods = v
	return s
}

func (s *QueryLogistics4DistributionResponseBodyModel) SetLogisticsCompanyCode(v string) *QueryLogistics4DistributionResponseBodyModel {
	s.LogisticsCompanyCode = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBodyModel) SetLogisticsCompanyName(v string) *QueryLogistics4DistributionResponseBodyModel {
	s.LogisticsCompanyName = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBodyModel) SetLogisticsDetailList(v []*QueryLogistics4DistributionResponseBodyModelLogisticsDetailList) *QueryLogistics4DistributionResponseBodyModel {
	s.LogisticsDetailList = v
	return s
}

func (s *QueryLogistics4DistributionResponseBodyModel) SetMailNo(v string) *QueryLogistics4DistributionResponseBodyModel {
	s.MailNo = &v
	return s
}

type QueryLogistics4DistributionResponseBodyModelGoods struct {
	GoodName *string `json:"GoodName,omitempty" xml:"GoodName,omitempty"`
	ItemId   *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	Quantity *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	SkuId    *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
}

func (s QueryLogistics4DistributionResponseBodyModelGoods) String() string {
	return tea.Prettify(s)
}

func (s QueryLogistics4DistributionResponseBodyModelGoods) GoString() string {
	return s.String()
}

func (s *QueryLogistics4DistributionResponseBodyModelGoods) SetGoodName(v string) *QueryLogistics4DistributionResponseBodyModelGoods {
	s.GoodName = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBodyModelGoods) SetItemId(v string) *QueryLogistics4DistributionResponseBodyModelGoods {
	s.ItemId = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBodyModelGoods) SetQuantity(v int32) *QueryLogistics4DistributionResponseBodyModelGoods {
	s.Quantity = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBodyModelGoods) SetSkuId(v string) *QueryLogistics4DistributionResponseBodyModelGoods {
	s.SkuId = &v
	return s
}

type QueryLogistics4DistributionResponseBodyModelLogisticsDetailList struct {
	OcurrTimeStr *string `json:"OcurrTimeStr,omitempty" xml:"OcurrTimeStr,omitempty"`
	StanderdDesc *string `json:"StanderdDesc,omitempty" xml:"StanderdDesc,omitempty"`
}

func (s QueryLogistics4DistributionResponseBodyModelLogisticsDetailList) String() string {
	return tea.Prettify(s)
}

func (s QueryLogistics4DistributionResponseBodyModelLogisticsDetailList) GoString() string {
	return s.String()
}

func (s *QueryLogistics4DistributionResponseBodyModelLogisticsDetailList) SetOcurrTimeStr(v string) *QueryLogistics4DistributionResponseBodyModelLogisticsDetailList {
	s.OcurrTimeStr = &v
	return s
}

func (s *QueryLogistics4DistributionResponseBodyModelLogisticsDetailList) SetStanderdDesc(v string) *QueryLogistics4DistributionResponseBodyModelLogisticsDetailList {
	s.StanderdDesc = &v
	return s
}

type QueryLogistics4DistributionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLogistics4DistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLogistics4DistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryLogistics4DistributionResponse) GoString() string {
	return s.String()
}

func (s *QueryLogistics4DistributionResponse) SetHeaders(v map[string]*string) *QueryLogistics4DistributionResponse {
	s.Headers = v
	return s
}

func (s *QueryLogistics4DistributionResponse) SetStatusCode(v int32) *QueryLogistics4DistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLogistics4DistributionResponse) SetBody(v *QueryLogistics4DistributionResponseBody) *QueryLogistics4DistributionResponse {
	s.Body = v
	return s
}

type QueryMallCategoryListRequest struct {
	CategoryId         *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	DistributionMallId *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributorId      *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	TenantId           *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryMallCategoryListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMallCategoryListRequest) GoString() string {
	return s.String()
}

func (s *QueryMallCategoryListRequest) SetCategoryId(v int64) *QueryMallCategoryListRequest {
	s.CategoryId = &v
	return s
}

func (s *QueryMallCategoryListRequest) SetDistributionMallId(v string) *QueryMallCategoryListRequest {
	s.DistributionMallId = &v
	return s
}

func (s *QueryMallCategoryListRequest) SetDistributorId(v string) *QueryMallCategoryListRequest {
	s.DistributorId = &v
	return s
}

func (s *QueryMallCategoryListRequest) SetTenantId(v string) *QueryMallCategoryListRequest {
	s.TenantId = &v
	return s
}

type QueryMallCategoryListResponseBody struct {
	Code       *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                   `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      []*QueryMallCategoryListResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string                                   `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string                                   `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMallCategoryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMallCategoryListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMallCategoryListResponseBody) SetCode(v string) *QueryMallCategoryListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMallCategoryListResponseBody) SetLogsId(v string) *QueryMallCategoryListResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryMallCategoryListResponseBody) SetMessage(v string) *QueryMallCategoryListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMallCategoryListResponseBody) SetModel(v []*QueryMallCategoryListResponseBodyModel) *QueryMallCategoryListResponseBody {
	s.Model = v
	return s
}

func (s *QueryMallCategoryListResponseBody) SetRequestId(v string) *QueryMallCategoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMallCategoryListResponseBody) SetSubCode(v string) *QueryMallCategoryListResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryMallCategoryListResponseBody) SetSubMessage(v string) *QueryMallCategoryListResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryMallCategoryListResponseBody) SetSuccess(v bool) *QueryMallCategoryListResponseBody {
	s.Success = &v
	return s
}

type QueryMallCategoryListResponseBodyModel struct {
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Leaf       *bool   `json:"Leaf,omitempty" xml:"Leaf,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId   *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s QueryMallCategoryListResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryMallCategoryListResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryMallCategoryListResponseBodyModel) SetCategoryId(v int64) *QueryMallCategoryListResponseBodyModel {
	s.CategoryId = &v
	return s
}

func (s *QueryMallCategoryListResponseBodyModel) SetLeaf(v bool) *QueryMallCategoryListResponseBodyModel {
	s.Leaf = &v
	return s
}

func (s *QueryMallCategoryListResponseBodyModel) SetName(v string) *QueryMallCategoryListResponseBodyModel {
	s.Name = &v
	return s
}

func (s *QueryMallCategoryListResponseBodyModel) SetParentId(v int64) *QueryMallCategoryListResponseBodyModel {
	s.ParentId = &v
	return s
}

type QueryMallCategoryListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMallCategoryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMallCategoryListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMallCategoryListResponse) GoString() string {
	return s.String()
}

func (s *QueryMallCategoryListResponse) SetHeaders(v map[string]*string) *QueryMallCategoryListResponse {
	s.Headers = v
	return s
}

func (s *QueryMallCategoryListResponse) SetStatusCode(v int32) *QueryMallCategoryListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMallCategoryListResponse) SetBody(v *QueryMallCategoryListResponseBody) *QueryMallCategoryListResponse {
	s.Body = v
	return s
}

type QueryOrderDetail4DistributionRequest struct {
	DistributorId           *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	MainDistributionOrderId *string `json:"MainDistributionOrderId,omitempty" xml:"MainDistributionOrderId,omitempty"`
	TenantId                *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryOrderDetail4DistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetail4DistributionRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderDetail4DistributionRequest) SetDistributorId(v string) *QueryOrderDetail4DistributionRequest {
	s.DistributorId = &v
	return s
}

func (s *QueryOrderDetail4DistributionRequest) SetMainDistributionOrderId(v string) *QueryOrderDetail4DistributionRequest {
	s.MainDistributionOrderId = &v
	return s
}

func (s *QueryOrderDetail4DistributionRequest) SetTenantId(v string) *QueryOrderDetail4DistributionRequest {
	s.TenantId = &v
	return s
}

type QueryOrderDetail4DistributionResponseBody struct {
	Code       *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                         `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *QueryOrderDetail4DistributionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	PageNumber *int64                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryOrderDetail4DistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetail4DistributionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderDetail4DistributionResponseBody) SetCode(v string) *QueryOrderDetail4DistributionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBody) SetLogsId(v string) *QueryOrderDetail4DistributionResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBody) SetMessage(v string) *QueryOrderDetail4DistributionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBody) SetModel(v *QueryOrderDetail4DistributionResponseBodyModel) *QueryOrderDetail4DistributionResponseBody {
	s.Model = v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBody) SetPageNumber(v int64) *QueryOrderDetail4DistributionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBody) SetPageSize(v int64) *QueryOrderDetail4DistributionResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBody) SetRequestId(v string) *QueryOrderDetail4DistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBody) SetSubCode(v string) *QueryOrderDetail4DistributionResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBody) SetSubMessage(v string) *QueryOrderDetail4DistributionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBody) SetSuccess(v bool) *QueryOrderDetail4DistributionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBody) SetTotalCount(v int64) *QueryOrderDetail4DistributionResponseBody {
	s.TotalCount = &v
	return s
}

type QueryOrderDetail4DistributionResponseBodyModel struct {
	CreateDate          *string                                                       `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DistributionOrderId *string                                                       `json:"DistributionOrderId,omitempty" xml:"DistributionOrderId,omitempty"`
	DistributorId       *string                                                       `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	LogisticsStatus     *string                                                       `json:"LogisticsStatus,omitempty" xml:"LogisticsStatus,omitempty"`
	OrderAmount         *string                                                       `json:"OrderAmount,omitempty" xml:"OrderAmount,omitempty"`
	OrderStatus         *string                                                       `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	SubOrderList        []*QueryOrderDetail4DistributionResponseBodyModelSubOrderList `json:"SubOrderList,omitempty" xml:"SubOrderList,omitempty" type:"Repeated"`
}

func (s QueryOrderDetail4DistributionResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetail4DistributionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryOrderDetail4DistributionResponseBodyModel) SetCreateDate(v string) *QueryOrderDetail4DistributionResponseBodyModel {
	s.CreateDate = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModel) SetDistributionOrderId(v string) *QueryOrderDetail4DistributionResponseBodyModel {
	s.DistributionOrderId = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModel) SetDistributorId(v string) *QueryOrderDetail4DistributionResponseBodyModel {
	s.DistributorId = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModel) SetLogisticsStatus(v string) *QueryOrderDetail4DistributionResponseBodyModel {
	s.LogisticsStatus = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModel) SetOrderAmount(v string) *QueryOrderDetail4DistributionResponseBodyModel {
	s.OrderAmount = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModel) SetOrderStatus(v string) *QueryOrderDetail4DistributionResponseBodyModel {
	s.OrderStatus = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModel) SetSubOrderList(v []*QueryOrderDetail4DistributionResponseBodyModelSubOrderList) *QueryOrderDetail4DistributionResponseBodyModel {
	s.SubOrderList = v
	return s
}

type QueryOrderDetail4DistributionResponseBodyModelSubOrderList struct {
	ItemId                  *string                                                                `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemPic                 *string                                                                `json:"ItemPic,omitempty" xml:"ItemPic,omitempty"`
	ItemPrice               []*QueryOrderDetail4DistributionResponseBodyModelSubOrderListItemPrice `json:"ItemPrice,omitempty" xml:"ItemPrice,omitempty" type:"Repeated"`
	ItemTitle               *string                                                                `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	LogisticsStatus         *string                                                                `json:"LogisticsStatus,omitempty" xml:"LogisticsStatus,omitempty"`
	MainDistributionOrderId *string                                                                `json:"MainDistributionOrderId,omitempty" xml:"MainDistributionOrderId,omitempty"`
	Number                  *string                                                                `json:"Number,omitempty" xml:"Number,omitempty"`
	OrderStatus             *string                                                                `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	SkuId                   *string                                                                `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuName                 *string                                                                `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
	SubDistributionOrderId  *string                                                                `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
}

func (s QueryOrderDetail4DistributionResponseBodyModelSubOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetail4DistributionResponseBodyModelSubOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderDetail4DistributionResponseBodyModelSubOrderList) SetItemId(v string) *QueryOrderDetail4DistributionResponseBodyModelSubOrderList {
	s.ItemId = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModelSubOrderList) SetItemPic(v string) *QueryOrderDetail4DistributionResponseBodyModelSubOrderList {
	s.ItemPic = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModelSubOrderList) SetItemPrice(v []*QueryOrderDetail4DistributionResponseBodyModelSubOrderListItemPrice) *QueryOrderDetail4DistributionResponseBodyModelSubOrderList {
	s.ItemPrice = v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModelSubOrderList) SetItemTitle(v string) *QueryOrderDetail4DistributionResponseBodyModelSubOrderList {
	s.ItemTitle = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModelSubOrderList) SetLogisticsStatus(v string) *QueryOrderDetail4DistributionResponseBodyModelSubOrderList {
	s.LogisticsStatus = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModelSubOrderList) SetMainDistributionOrderId(v string) *QueryOrderDetail4DistributionResponseBodyModelSubOrderList {
	s.MainDistributionOrderId = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModelSubOrderList) SetNumber(v string) *QueryOrderDetail4DistributionResponseBodyModelSubOrderList {
	s.Number = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModelSubOrderList) SetOrderStatus(v string) *QueryOrderDetail4DistributionResponseBodyModelSubOrderList {
	s.OrderStatus = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModelSubOrderList) SetSkuId(v string) *QueryOrderDetail4DistributionResponseBodyModelSubOrderList {
	s.SkuId = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModelSubOrderList) SetSkuName(v string) *QueryOrderDetail4DistributionResponseBodyModelSubOrderList {
	s.SkuName = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponseBodyModelSubOrderList) SetSubDistributionOrderId(v string) *QueryOrderDetail4DistributionResponseBodyModelSubOrderList {
	s.SubDistributionOrderId = &v
	return s
}

type QueryOrderDetail4DistributionResponseBodyModelSubOrderListItemPrice struct {
	FundAmountMoney *string `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
}

func (s QueryOrderDetail4DistributionResponseBodyModelSubOrderListItemPrice) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetail4DistributionResponseBodyModelSubOrderListItemPrice) GoString() string {
	return s.String()
}

func (s *QueryOrderDetail4DistributionResponseBodyModelSubOrderListItemPrice) SetFundAmountMoney(v string) *QueryOrderDetail4DistributionResponseBodyModelSubOrderListItemPrice {
	s.FundAmountMoney = &v
	return s
}

type QueryOrderDetail4DistributionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrderDetail4DistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrderDetail4DistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderDetail4DistributionResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderDetail4DistributionResponse) SetHeaders(v map[string]*string) *QueryOrderDetail4DistributionResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderDetail4DistributionResponse) SetStatusCode(v int32) *QueryOrderDetail4DistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrderDetail4DistributionResponse) SetBody(v *QueryOrderDetail4DistributionResponseBody) *QueryOrderDetail4DistributionResponse {
	s.Body = v
	return s
}

type QueryOrderList4DistributionRequest struct {
	DistributorId *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	FilterOption  *string `json:"FilterOption,omitempty" xml:"FilterOption,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryOrderList4DistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderList4DistributionRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderList4DistributionRequest) SetDistributorId(v string) *QueryOrderList4DistributionRequest {
	s.DistributorId = &v
	return s
}

func (s *QueryOrderList4DistributionRequest) SetFilterOption(v string) *QueryOrderList4DistributionRequest {
	s.FilterOption = &v
	return s
}

func (s *QueryOrderList4DistributionRequest) SetPageNumber(v int32) *QueryOrderList4DistributionRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryOrderList4DistributionRequest) SetPageSize(v int32) *QueryOrderList4DistributionRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOrderList4DistributionRequest) SetTenantId(v string) *QueryOrderList4DistributionRequest {
	s.TenantId = &v
	return s
}

type QueryOrderList4DistributionResponseBody struct {
	Code       *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                         `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      []*QueryOrderList4DistributionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	PageNumber *int64                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryOrderList4DistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderList4DistributionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderList4DistributionResponseBody) SetCode(v string) *QueryOrderList4DistributionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBody) SetLogsId(v string) *QueryOrderList4DistributionResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBody) SetMessage(v string) *QueryOrderList4DistributionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBody) SetModel(v []*QueryOrderList4DistributionResponseBodyModel) *QueryOrderList4DistributionResponseBody {
	s.Model = v
	return s
}

func (s *QueryOrderList4DistributionResponseBody) SetPageNumber(v int64) *QueryOrderList4DistributionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBody) SetPageSize(v int64) *QueryOrderList4DistributionResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBody) SetRequestId(v string) *QueryOrderList4DistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBody) SetSubCode(v string) *QueryOrderList4DistributionResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBody) SetSubMessage(v string) *QueryOrderList4DistributionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBody) SetSuccess(v bool) *QueryOrderList4DistributionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBody) SetTotalCount(v int64) *QueryOrderList4DistributionResponseBody {
	s.TotalCount = &v
	return s
}

type QueryOrderList4DistributionResponseBodyModel struct {
	CreateDate          *string                                                     `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DistributionOrderId *string                                                     `json:"DistributionOrderId,omitempty" xml:"DistributionOrderId,omitempty"`
	DistributorId       *string                                                     `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	LogisticsStatus     *string                                                     `json:"LogisticsStatus,omitempty" xml:"LogisticsStatus,omitempty"`
	OrderAmount         *string                                                     `json:"OrderAmount,omitempty" xml:"OrderAmount,omitempty"`
	OrderStatus         *string                                                     `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	SubOrderList        []*QueryOrderList4DistributionResponseBodyModelSubOrderList `json:"SubOrderList,omitempty" xml:"SubOrderList,omitempty" type:"Repeated"`
}

func (s QueryOrderList4DistributionResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderList4DistributionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryOrderList4DistributionResponseBodyModel) SetCreateDate(v string) *QueryOrderList4DistributionResponseBodyModel {
	s.CreateDate = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModel) SetDistributionOrderId(v string) *QueryOrderList4DistributionResponseBodyModel {
	s.DistributionOrderId = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModel) SetDistributorId(v string) *QueryOrderList4DistributionResponseBodyModel {
	s.DistributorId = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModel) SetLogisticsStatus(v string) *QueryOrderList4DistributionResponseBodyModel {
	s.LogisticsStatus = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModel) SetOrderAmount(v string) *QueryOrderList4DistributionResponseBodyModel {
	s.OrderAmount = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModel) SetOrderStatus(v string) *QueryOrderList4DistributionResponseBodyModel {
	s.OrderStatus = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModel) SetSubOrderList(v []*QueryOrderList4DistributionResponseBodyModelSubOrderList) *QueryOrderList4DistributionResponseBodyModel {
	s.SubOrderList = v
	return s
}

type QueryOrderList4DistributionResponseBodyModelSubOrderList struct {
	ItemId                  *string                                                              `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemPic                 *string                                                              `json:"ItemPic,omitempty" xml:"ItemPic,omitempty"`
	ItemPrice               []*QueryOrderList4DistributionResponseBodyModelSubOrderListItemPrice `json:"ItemPrice,omitempty" xml:"ItemPrice,omitempty" type:"Repeated"`
	ItemTitle               *string                                                              `json:"ItemTitle,omitempty" xml:"ItemTitle,omitempty"`
	LogisticsStatus         *string                                                              `json:"LogisticsStatus,omitempty" xml:"LogisticsStatus,omitempty"`
	MainDistributionOrderId *string                                                              `json:"MainDistributionOrderId,omitempty" xml:"MainDistributionOrderId,omitempty"`
	Number                  *string                                                              `json:"Number,omitempty" xml:"Number,omitempty"`
	OrderStatus             *string                                                              `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	SkuId                   *string                                                              `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuName                 *string                                                              `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
	SubDistributionOrderId  *string                                                              `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
}

func (s QueryOrderList4DistributionResponseBodyModelSubOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderList4DistributionResponseBodyModelSubOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrderList4DistributionResponseBodyModelSubOrderList) SetItemId(v string) *QueryOrderList4DistributionResponseBodyModelSubOrderList {
	s.ItemId = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModelSubOrderList) SetItemPic(v string) *QueryOrderList4DistributionResponseBodyModelSubOrderList {
	s.ItemPic = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModelSubOrderList) SetItemPrice(v []*QueryOrderList4DistributionResponseBodyModelSubOrderListItemPrice) *QueryOrderList4DistributionResponseBodyModelSubOrderList {
	s.ItemPrice = v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModelSubOrderList) SetItemTitle(v string) *QueryOrderList4DistributionResponseBodyModelSubOrderList {
	s.ItemTitle = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModelSubOrderList) SetLogisticsStatus(v string) *QueryOrderList4DistributionResponseBodyModelSubOrderList {
	s.LogisticsStatus = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModelSubOrderList) SetMainDistributionOrderId(v string) *QueryOrderList4DistributionResponseBodyModelSubOrderList {
	s.MainDistributionOrderId = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModelSubOrderList) SetNumber(v string) *QueryOrderList4DistributionResponseBodyModelSubOrderList {
	s.Number = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModelSubOrderList) SetOrderStatus(v string) *QueryOrderList4DistributionResponseBodyModelSubOrderList {
	s.OrderStatus = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModelSubOrderList) SetSkuId(v string) *QueryOrderList4DistributionResponseBodyModelSubOrderList {
	s.SkuId = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModelSubOrderList) SetSkuName(v string) *QueryOrderList4DistributionResponseBodyModelSubOrderList {
	s.SkuName = &v
	return s
}

func (s *QueryOrderList4DistributionResponseBodyModelSubOrderList) SetSubDistributionOrderId(v string) *QueryOrderList4DistributionResponseBodyModelSubOrderList {
	s.SubDistributionOrderId = &v
	return s
}

type QueryOrderList4DistributionResponseBodyModelSubOrderListItemPrice struct {
	FundAmountMoney *string `json:"FundAmountMoney,omitempty" xml:"FundAmountMoney,omitempty"`
}

func (s QueryOrderList4DistributionResponseBodyModelSubOrderListItemPrice) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderList4DistributionResponseBodyModelSubOrderListItemPrice) GoString() string {
	return s.String()
}

func (s *QueryOrderList4DistributionResponseBodyModelSubOrderListItemPrice) SetFundAmountMoney(v string) *QueryOrderList4DistributionResponseBodyModelSubOrderListItemPrice {
	s.FundAmountMoney = &v
	return s
}

type QueryOrderList4DistributionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrderList4DistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrderList4DistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderList4DistributionResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderList4DistributionResponse) SetHeaders(v map[string]*string) *QueryOrderList4DistributionResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderList4DistributionResponse) SetStatusCode(v int32) *QueryOrderList4DistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrderList4DistributionResponse) SetBody(v *QueryOrderList4DistributionResponseBody) *QueryOrderList4DistributionResponse {
	s.Body = v
	return s
}

type QueryRefundApplicationDetail4DistributionRequest struct {
	DistributorId          *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	SubDistributionOrderId *string `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
	TenantId               *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryRefundApplicationDetail4DistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRefundApplicationDetail4DistributionRequest) GoString() string {
	return s.String()
}

func (s *QueryRefundApplicationDetail4DistributionRequest) SetDistributorId(v string) *QueryRefundApplicationDetail4DistributionRequest {
	s.DistributorId = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionRequest) SetSubDistributionOrderId(v string) *QueryRefundApplicationDetail4DistributionRequest {
	s.SubDistributionOrderId = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionRequest) SetTenantId(v string) *QueryRefundApplicationDetail4DistributionRequest {
	s.TenantId = &v
	return s
}

type QueryRefundApplicationDetail4DistributionResponseBody struct {
	Code       *string                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                                     `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *QueryRefundApplicationDetail4DistributionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	PageNumber *int64                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryRefundApplicationDetail4DistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRefundApplicationDetail4DistributionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRefundApplicationDetail4DistributionResponseBody) SetCode(v string) *QueryRefundApplicationDetail4DistributionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBody) SetLogsId(v string) *QueryRefundApplicationDetail4DistributionResponseBody {
	s.LogsId = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBody) SetMessage(v string) *QueryRefundApplicationDetail4DistributionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBody) SetModel(v *QueryRefundApplicationDetail4DistributionResponseBodyModel) *QueryRefundApplicationDetail4DistributionResponseBody {
	s.Model = v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBody) SetPageNumber(v int64) *QueryRefundApplicationDetail4DistributionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBody) SetPageSize(v int64) *QueryRefundApplicationDetail4DistributionResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBody) SetRequestId(v string) *QueryRefundApplicationDetail4DistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBody) SetSubCode(v string) *QueryRefundApplicationDetail4DistributionResponseBody {
	s.SubCode = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBody) SetSubMessage(v string) *QueryRefundApplicationDetail4DistributionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBody) SetSuccess(v bool) *QueryRefundApplicationDetail4DistributionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBody) SetTotalCount(v int64) *QueryRefundApplicationDetail4DistributionResponseBody {
	s.TotalCount = &v
	return s
}

type QueryRefundApplicationDetail4DistributionResponseBodyModel struct {
	ApplyDisputeDesc             *string                                                                  `json:"ApplyDisputeDesc,omitempty" xml:"ApplyDisputeDesc,omitempty"`
	ApplyReason                  *QueryRefundApplicationDetail4DistributionResponseBodyModelApplyReason   `json:"ApplyReason,omitempty" xml:"ApplyReason,omitempty" type:"Struct"`
	BizClaimType                 *int32                                                                   `json:"BizClaimType,omitempty" xml:"BizClaimType,omitempty"`
	DisputeCreateTime            *string                                                                  `json:"DisputeCreateTime,omitempty" xml:"DisputeCreateTime,omitempty"`
	DisputeDesc                  *string                                                                  `json:"DisputeDesc,omitempty" xml:"DisputeDesc,omitempty"`
	DisputeEndTime               *string                                                                  `json:"DisputeEndTime,omitempty" xml:"DisputeEndTime,omitempty"`
	DisputeId                    *int64                                                                   `json:"DisputeId,omitempty" xml:"DisputeId,omitempty"`
	DisputeStatus                *int32                                                                   `json:"DisputeStatus,omitempty" xml:"DisputeStatus,omitempty"`
	DisputeType                  *int32                                                                   `json:"DisputeType,omitempty" xml:"DisputeType,omitempty"`
	DistributionOrderId          *string                                                                  `json:"DistributionOrderId,omitempty" xml:"DistributionOrderId,omitempty"`
	OrderLogisticsStatus         *int32                                                                   `json:"OrderLogisticsStatus,omitempty" xml:"OrderLogisticsStatus,omitempty"`
	RealRefundFee                *int64                                                                   `json:"RealRefundFee,omitempty" xml:"RealRefundFee,omitempty"`
	RefundFee                    *int64                                                                   `json:"RefundFee,omitempty" xml:"RefundFee,omitempty"`
	RefundFeeData                *QueryRefundApplicationDetail4DistributionResponseBodyModelRefundFeeData `json:"RefundFeeData,omitempty" xml:"RefundFeeData,omitempty" type:"Struct"`
	RefunderAddress              *string                                                                  `json:"RefunderAddress,omitempty" xml:"RefunderAddress,omitempty"`
	RefunderName                 *string                                                                  `json:"RefunderName,omitempty" xml:"RefunderName,omitempty"`
	RefunderTel                  *string                                                                  `json:"RefunderTel,omitempty" xml:"RefunderTel,omitempty"`
	RefunderZipCode              *string                                                                  `json:"RefunderZipCode,omitempty" xml:"RefunderZipCode,omitempty"`
	ReturnGoodCount              *int64                                                                   `json:"ReturnGoodCount,omitempty" xml:"ReturnGoodCount,omitempty"`
	ReturnGoodLogisticsStatus    *int32                                                                   `json:"ReturnGoodLogisticsStatus,omitempty" xml:"ReturnGoodLogisticsStatus,omitempty"`
	SellerAgreeMsg               *string                                                                  `json:"SellerAgreeMsg,omitempty" xml:"SellerAgreeMsg,omitempty"`
	SellerRefuseAgreementMessage *string                                                                  `json:"SellerRefuseAgreementMessage,omitempty" xml:"SellerRefuseAgreementMessage,omitempty"`
	SellerRefuseReason           *string                                                                  `json:"SellerRefuseReason,omitempty" xml:"SellerRefuseReason,omitempty"`
	SubDistributionOrderId       *string                                                                  `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
}

func (s QueryRefundApplicationDetail4DistributionResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QueryRefundApplicationDetail4DistributionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetApplyDisputeDesc(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.ApplyDisputeDesc = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetApplyReason(v *QueryRefundApplicationDetail4DistributionResponseBodyModelApplyReason) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.ApplyReason = v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetBizClaimType(v int32) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.BizClaimType = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetDisputeCreateTime(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.DisputeCreateTime = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetDisputeDesc(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.DisputeDesc = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetDisputeEndTime(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.DisputeEndTime = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetDisputeId(v int64) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.DisputeId = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetDisputeStatus(v int32) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.DisputeStatus = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetDisputeType(v int32) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.DisputeType = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetDistributionOrderId(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.DistributionOrderId = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetOrderLogisticsStatus(v int32) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.OrderLogisticsStatus = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetRealRefundFee(v int64) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.RealRefundFee = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetRefundFee(v int64) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.RefundFee = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetRefundFeeData(v *QueryRefundApplicationDetail4DistributionResponseBodyModelRefundFeeData) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.RefundFeeData = v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetRefunderAddress(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.RefunderAddress = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetRefunderName(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.RefunderName = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetRefunderTel(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.RefunderTel = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetRefunderZipCode(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.RefunderZipCode = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetReturnGoodCount(v int64) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.ReturnGoodCount = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetReturnGoodLogisticsStatus(v int32) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.ReturnGoodLogisticsStatus = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetSellerAgreeMsg(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.SellerAgreeMsg = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetSellerRefuseAgreementMessage(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.SellerRefuseAgreementMessage = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetSellerRefuseReason(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.SellerRefuseReason = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModel) SetSubDistributionOrderId(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModel {
	s.SubDistributionOrderId = &v
	return s
}

type QueryRefundApplicationDetail4DistributionResponseBodyModelApplyReason struct {
	ReasonTextId *int64  `json:"ReasonTextId,omitempty" xml:"ReasonTextId,omitempty"`
	ReasonTips   *string `json:"ReasonTips,omitempty" xml:"ReasonTips,omitempty"`
}

func (s QueryRefundApplicationDetail4DistributionResponseBodyModelApplyReason) String() string {
	return tea.Prettify(s)
}

func (s QueryRefundApplicationDetail4DistributionResponseBodyModelApplyReason) GoString() string {
	return s.String()
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModelApplyReason) SetReasonTextId(v int64) *QueryRefundApplicationDetail4DistributionResponseBodyModelApplyReason {
	s.ReasonTextId = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModelApplyReason) SetReasonTips(v string) *QueryRefundApplicationDetail4DistributionResponseBodyModelApplyReason {
	s.ReasonTips = &v
	return s
}

type QueryRefundApplicationDetail4DistributionResponseBodyModelRefundFeeData struct {
	MaxRefundFee *int64 `json:"MaxRefundFee,omitempty" xml:"MaxRefundFee,omitempty"`
	MinRefundFee *int64 `json:"MinRefundFee,omitempty" xml:"MinRefundFee,omitempty"`
}

func (s QueryRefundApplicationDetail4DistributionResponseBodyModelRefundFeeData) String() string {
	return tea.Prettify(s)
}

func (s QueryRefundApplicationDetail4DistributionResponseBodyModelRefundFeeData) GoString() string {
	return s.String()
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModelRefundFeeData) SetMaxRefundFee(v int64) *QueryRefundApplicationDetail4DistributionResponseBodyModelRefundFeeData {
	s.MaxRefundFee = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponseBodyModelRefundFeeData) SetMinRefundFee(v int64) *QueryRefundApplicationDetail4DistributionResponseBodyModelRefundFeeData {
	s.MinRefundFee = &v
	return s
}

type QueryRefundApplicationDetail4DistributionResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRefundApplicationDetail4DistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRefundApplicationDetail4DistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRefundApplicationDetail4DistributionResponse) GoString() string {
	return s.String()
}

func (s *QueryRefundApplicationDetail4DistributionResponse) SetHeaders(v map[string]*string) *QueryRefundApplicationDetail4DistributionResponse {
	s.Headers = v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponse) SetStatusCode(v int32) *QueryRefundApplicationDetail4DistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRefundApplicationDetail4DistributionResponse) SetBody(v *QueryRefundApplicationDetail4DistributionResponseBody) *QueryRefundApplicationDetail4DistributionResponse {
	s.Body = v
	return s
}

type RenderDistributionOrderRequest struct {
	BuyerId                *string                                        `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	DeliveryAddress        *string                                        `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	DistributionSupplierId *string                                        `json:"DistributionSupplierId,omitempty" xml:"DistributionSupplierId,omitempty"`
	DistributorId          *string                                        `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	ExtInfo                *string                                        `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	ItemInfoLists          []*RenderDistributionOrderRequestItemInfoLists `json:"ItemInfoLists,omitempty" xml:"ItemInfoLists,omitempty" type:"Repeated"`
	TenantId               *string                                        `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s RenderDistributionOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderRequest) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderRequest) SetBuyerId(v string) *RenderDistributionOrderRequest {
	s.BuyerId = &v
	return s
}

func (s *RenderDistributionOrderRequest) SetDeliveryAddress(v string) *RenderDistributionOrderRequest {
	s.DeliveryAddress = &v
	return s
}

func (s *RenderDistributionOrderRequest) SetDistributionSupplierId(v string) *RenderDistributionOrderRequest {
	s.DistributionSupplierId = &v
	return s
}

func (s *RenderDistributionOrderRequest) SetDistributorId(v string) *RenderDistributionOrderRequest {
	s.DistributorId = &v
	return s
}

func (s *RenderDistributionOrderRequest) SetExtInfo(v string) *RenderDistributionOrderRequest {
	s.ExtInfo = &v
	return s
}

func (s *RenderDistributionOrderRequest) SetItemInfoLists(v []*RenderDistributionOrderRequestItemInfoLists) *RenderDistributionOrderRequest {
	s.ItemInfoLists = v
	return s
}

func (s *RenderDistributionOrderRequest) SetTenantId(v string) *RenderDistributionOrderRequest {
	s.TenantId = &v
	return s
}

type RenderDistributionOrderRequestItemInfoLists struct {
	DistributionMallId *string `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	LmItemId           *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	Quantity           *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	SkuId              *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
}

func (s RenderDistributionOrderRequestItemInfoLists) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderRequestItemInfoLists) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderRequestItemInfoLists) SetDistributionMallId(v string) *RenderDistributionOrderRequestItemInfoLists {
	s.DistributionMallId = &v
	return s
}

func (s *RenderDistributionOrderRequestItemInfoLists) SetLmItemId(v string) *RenderDistributionOrderRequestItemInfoLists {
	s.LmItemId = &v
	return s
}

func (s *RenderDistributionOrderRequestItemInfoLists) SetQuantity(v int32) *RenderDistributionOrderRequestItemInfoLists {
	s.Quantity = &v
	return s
}

func (s *RenderDistributionOrderRequestItemInfoLists) SetSkuId(v string) *RenderDistributionOrderRequestItemInfoLists {
	s.SkuId = &v
	return s
}

type RenderDistributionOrderShrinkRequest struct {
	BuyerId                *string `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	DeliveryAddress        *string `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	DistributionSupplierId *string `json:"DistributionSupplierId,omitempty" xml:"DistributionSupplierId,omitempty"`
	DistributorId          *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	ExtInfo                *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	ItemInfoListsShrink    *string `json:"ItemInfoLists,omitempty" xml:"ItemInfoLists,omitempty"`
	TenantId               *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s RenderDistributionOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderShrinkRequest) SetBuyerId(v string) *RenderDistributionOrderShrinkRequest {
	s.BuyerId = &v
	return s
}

func (s *RenderDistributionOrderShrinkRequest) SetDeliveryAddress(v string) *RenderDistributionOrderShrinkRequest {
	s.DeliveryAddress = &v
	return s
}

func (s *RenderDistributionOrderShrinkRequest) SetDistributionSupplierId(v string) *RenderDistributionOrderShrinkRequest {
	s.DistributionSupplierId = &v
	return s
}

func (s *RenderDistributionOrderShrinkRequest) SetDistributorId(v string) *RenderDistributionOrderShrinkRequest {
	s.DistributorId = &v
	return s
}

func (s *RenderDistributionOrderShrinkRequest) SetExtInfo(v string) *RenderDistributionOrderShrinkRequest {
	s.ExtInfo = &v
	return s
}

func (s *RenderDistributionOrderShrinkRequest) SetItemInfoListsShrink(v string) *RenderDistributionOrderShrinkRequest {
	s.ItemInfoListsShrink = &v
	return s
}

func (s *RenderDistributionOrderShrinkRequest) SetTenantId(v string) *RenderDistributionOrderShrinkRequest {
	s.TenantId = &v
	return s
}

type RenderDistributionOrderResponseBody struct {
	Code       *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string                                   `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Model      *RenderDistributionOrderResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	PageNumber *int64                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s RenderDistributionOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBody) SetCode(v string) *RenderDistributionOrderResponseBody {
	s.Code = &v
	return s
}

func (s *RenderDistributionOrderResponseBody) SetLogsId(v string) *RenderDistributionOrderResponseBody {
	s.LogsId = &v
	return s
}

func (s *RenderDistributionOrderResponseBody) SetMessage(v string) *RenderDistributionOrderResponseBody {
	s.Message = &v
	return s
}

func (s *RenderDistributionOrderResponseBody) SetModel(v *RenderDistributionOrderResponseBodyModel) *RenderDistributionOrderResponseBody {
	s.Model = v
	return s
}

func (s *RenderDistributionOrderResponseBody) SetPageNumber(v int64) *RenderDistributionOrderResponseBody {
	s.PageNumber = &v
	return s
}

func (s *RenderDistributionOrderResponseBody) SetPageSize(v int64) *RenderDistributionOrderResponseBody {
	s.PageSize = &v
	return s
}

func (s *RenderDistributionOrderResponseBody) SetRequestId(v string) *RenderDistributionOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenderDistributionOrderResponseBody) SetSubCode(v string) *RenderDistributionOrderResponseBody {
	s.SubCode = &v
	return s
}

func (s *RenderDistributionOrderResponseBody) SetSubMessage(v string) *RenderDistributionOrderResponseBody {
	s.SubMessage = &v
	return s
}

func (s *RenderDistributionOrderResponseBody) SetSuccess(v bool) *RenderDistributionOrderResponseBody {
	s.Success = &v
	return s
}

func (s *RenderDistributionOrderResponseBody) SetTotalCount(v int64) *RenderDistributionOrderResponseBody {
	s.TotalCount = &v
	return s
}

type RenderDistributionOrderResponseBodyModel struct {
	AddressInfos               []*RenderDistributionOrderResponseBodyModelAddressInfos               `json:"AddressInfos,omitempty" xml:"AddressInfos,omitempty" type:"Repeated"`
	CanSell                    *bool                                                                 `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	ExtInfo                    map[string]*string                                                    `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	Message                    *string                                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RenderOrderInfos           []*RenderDistributionOrderResponseBodyModelRenderOrderInfos           `json:"RenderOrderInfos,omitempty" xml:"RenderOrderInfos,omitempty" type:"Repeated"`
	UnsellableRenderOrderInfos []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos `json:"UnsellableRenderOrderInfos,omitempty" xml:"UnsellableRenderOrderInfos,omitempty" type:"Repeated"`
}

func (s RenderDistributionOrderResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModel) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModel) SetAddressInfos(v []*RenderDistributionOrderResponseBodyModelAddressInfos) *RenderDistributionOrderResponseBodyModel {
	s.AddressInfos = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModel) SetCanSell(v bool) *RenderDistributionOrderResponseBodyModel {
	s.CanSell = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModel) SetExtInfo(v map[string]*string) *RenderDistributionOrderResponseBodyModel {
	s.ExtInfo = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModel) SetMessage(v string) *RenderDistributionOrderResponseBodyModel {
	s.Message = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModel) SetRenderOrderInfos(v []*RenderDistributionOrderResponseBodyModelRenderOrderInfos) *RenderDistributionOrderResponseBodyModel {
	s.RenderOrderInfos = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModel) SetUnsellableRenderOrderInfos(v []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos) *RenderDistributionOrderResponseBodyModel {
	s.UnsellableRenderOrderInfos = v
	return s
}

type RenderDistributionOrderResponseBodyModelAddressInfos struct {
	AddressDetail    *string `json:"AddressDetail,omitempty" xml:"AddressDetail,omitempty"`
	AddressId        *int64  `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	DivisionCode     *string `json:"DivisionCode,omitempty" xml:"DivisionCode,omitempty"`
	IsDefault        *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Receiver         *string `json:"Receiver,omitempty" xml:"Receiver,omitempty"`
	ReceiverPhone    *string `json:"ReceiverPhone,omitempty" xml:"ReceiverPhone,omitempty"`
	TownDivisionCode *string `json:"TownDivisionCode,omitempty" xml:"TownDivisionCode,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelAddressInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelAddressInfos) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelAddressInfos) SetAddressDetail(v string) *RenderDistributionOrderResponseBodyModelAddressInfos {
	s.AddressDetail = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelAddressInfos) SetAddressId(v int64) *RenderDistributionOrderResponseBodyModelAddressInfos {
	s.AddressId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelAddressInfos) SetDivisionCode(v string) *RenderDistributionOrderResponseBodyModelAddressInfos {
	s.DivisionCode = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelAddressInfos) SetIsDefault(v bool) *RenderDistributionOrderResponseBodyModelAddressInfos {
	s.IsDefault = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelAddressInfos) SetReceiver(v string) *RenderDistributionOrderResponseBodyModelAddressInfos {
	s.Receiver = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelAddressInfos) SetReceiverPhone(v string) *RenderDistributionOrderResponseBodyModelAddressInfos {
	s.ReceiverPhone = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelAddressInfos) SetTownDivisionCode(v string) *RenderDistributionOrderResponseBodyModelAddressInfos {
	s.TownDivisionCode = &v
	return s
}

type RenderDistributionOrderResponseBodyModelRenderOrderInfos struct {
	CanSell         *bool                                                                      `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	DeliveryInfos   []*RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos   `json:"DeliveryInfos,omitempty" xml:"DeliveryInfos,omitempty" type:"Repeated"`
	ExtInfo         map[string]*string                                                         `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	InvoiceInfo     *RenderDistributionOrderResponseBodyModelRenderOrderInfosInvoiceInfo       `json:"InvoiceInfo,omitempty" xml:"InvoiceInfo,omitempty" type:"Struct"`
	ItemInfos       []*RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos       `json:"ItemInfos,omitempty" xml:"ItemInfos,omitempty" type:"Repeated"`
	Message         *string                                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	ShopPromInstVOS []*RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS `json:"ShopPromInstVOS,omitempty" xml:"ShopPromInstVOS,omitempty" type:"Repeated"`
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfos) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfos) SetCanSell(v bool) *RenderDistributionOrderResponseBodyModelRenderOrderInfos {
	s.CanSell = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfos) SetDeliveryInfos(v []*RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos) *RenderDistributionOrderResponseBodyModelRenderOrderInfos {
	s.DeliveryInfos = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfos) SetExtInfo(v map[string]*string) *RenderDistributionOrderResponseBodyModelRenderOrderInfos {
	s.ExtInfo = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfos) SetInvoiceInfo(v *RenderDistributionOrderResponseBodyModelRenderOrderInfosInvoiceInfo) *RenderDistributionOrderResponseBodyModelRenderOrderInfos {
	s.InvoiceInfo = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfos) SetItemInfos(v []*RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) *RenderDistributionOrderResponseBodyModelRenderOrderInfos {
	s.ItemInfos = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfos) SetMessage(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfos {
	s.Message = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfos) SetShopPromInstVOS(v []*RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) *RenderDistributionOrderResponseBodyModelRenderOrderInfos {
	s.ShopPromInstVOS = v
	return s
}

type RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	PostFee     *int64  `json:"PostFee,omitempty" xml:"PostFee,omitempty"`
	ServiceType *int64  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos) SetDisplayName(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos {
	s.DisplayName = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos) SetId(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos {
	s.Id = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos) SetPostFee(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos {
	s.PostFee = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos) SetServiceType(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosDeliveryInfos {
	s.ServiceType = &v
	return s
}

type RenderDistributionOrderResponseBodyModelRenderOrderInfosInvoiceInfo struct {
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosInvoiceInfo) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosInvoiceInfo) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosInvoiceInfo) SetDesc(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosInvoiceInfo {
	s.Desc = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosInvoiceInfo) SetType(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosInvoiceInfo {
	s.Type = &v
	return s
}

type RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos struct {
	CanSell                *bool                                                                               `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	DistributionMallId     *string                                                                             `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributionSupplierId *string                                                                             `json:"DistributionSupplierId,omitempty" xml:"DistributionSupplierId,omitempty"`
	DistributorId          *string                                                                             `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	Features               map[string]*string                                                                  `json:"Features,omitempty" xml:"Features,omitempty"`
	ItemId                 *string                                                                             `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName               *string                                                                             `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	ItemPicUrl             *string                                                                             `json:"ItemPicUrl,omitempty" xml:"ItemPicUrl,omitempty"`
	ItemPromInstVOS        []*RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS `json:"ItemPromInstVOS,omitempty" xml:"ItemPromInstVOS,omitempty" type:"Repeated"`
	ItemUrl                *string                                                                             `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	Message                *string                                                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Price                  *int64                                                                              `json:"Price,omitempty" xml:"Price,omitempty"`
	PromotionFee           *int64                                                                              `json:"PromotionFee,omitempty" xml:"PromotionFee,omitempty"`
	Quantity               *int32                                                                              `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ReservePrice           *int64                                                                              `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	SkuId                  *int64                                                                              `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuName                *string                                                                             `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
	VirtualItemType        *string                                                                             `json:"VirtualItemType,omitempty" xml:"VirtualItemType,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetCanSell(v bool) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.CanSell = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetDistributionMallId(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.DistributionMallId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetDistributionSupplierId(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.DistributionSupplierId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetDistributorId(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.DistributorId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetFeatures(v map[string]*string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.Features = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetItemId(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.ItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetItemName(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.ItemName = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetItemPicUrl(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.ItemPicUrl = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetItemPromInstVOS(v []*RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.ItemPromInstVOS = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetItemUrl(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.ItemUrl = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetMessage(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.Message = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetPrice(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.Price = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetPromotionFee(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.PromotionFee = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetQuantity(v int32) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.Quantity = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetReservePrice(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.ReservePrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetSkuId(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.SkuId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetSkuName(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.SkuName = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos) SetVirtualItemType(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfos {
	s.VirtualItemType = &v
	return s
}

type RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS struct {
	AvailableItems []*RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems `json:"AvailableItems,omitempty" xml:"AvailableItems,omitempty" type:"Repeated"`
	CanUse         *bool                                                                                             `json:"CanUse,omitempty" xml:"CanUse,omitempty"`
	DiscountPrice  *int64                                                                                            `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	ExpireTime     *int64                                                                                            `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId     *string                                                                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Level          *string                                                                                           `json:"Level,omitempty" xml:"Level,omitempty"`
	LmItemId       *string                                                                                           `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	PromotionName  *string                                                                                           `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionType  *string                                                                                           `json:"PromotionType,omitempty" xml:"PromotionType,omitempty"`
	Reason         *string                                                                                           `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Selected       *bool                                                                                             `json:"Selected,omitempty" xml:"Selected,omitempty"`
	SkuIds         []*int64                                                                                          `json:"SkuIds,omitempty" xml:"SkuIds,omitempty" type:"Repeated"`
	SpecialPrice   *int64                                                                                            `json:"SpecialPrice,omitempty" xml:"SpecialPrice,omitempty"`
	SubBizCode     *string                                                                                           `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	TbSellerId     *int64                                                                                            `json:"TbSellerId,omitempty" xml:"TbSellerId,omitempty"`
	ThresholdPrice *int64                                                                                            `json:"ThresholdPrice,omitempty" xml:"ThresholdPrice,omitempty"`
	UseStartTime   *int64                                                                                            `json:"UseStartTime,omitempty" xml:"UseStartTime,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetAvailableItems(v []*RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.AvailableItems = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetCanUse(v bool) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.CanUse = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetDiscountPrice(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.DiscountPrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetExpireTime(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.ExpireTime = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetInstanceId(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.InstanceId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetLevel(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.Level = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetLmItemId(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.LmItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetPromotionName(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.PromotionName = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetPromotionType(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.PromotionType = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetReason(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.Reason = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetSelected(v bool) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.Selected = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetSkuIds(v []*int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.SkuIds = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetSpecialPrice(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.SpecialPrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetSubBizCode(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.SubBizCode = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetTbSellerId(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.TbSellerId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetThresholdPrice(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.ThresholdPrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS) SetUseStartTime(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOS {
	s.UseStartTime = &v
	return s
}

type RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems struct {
	ItemId       *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId     *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	LmShopId     *int64  `json:"LmShopId,omitempty" xml:"LmShopId,omitempty"`
	Number       *int32  `json:"Number,omitempty" xml:"Number,omitempty"`
	Points       *int64  `json:"Points,omitempty" xml:"Points,omitempty"`
	PointsAmount *int64  `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	PriceCent    *int64  `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	Removed      *bool   `json:"Removed,omitempty" xml:"Removed,omitempty"`
	SkuId        *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	TbSellerId   *int64  `json:"TbSellerId,omitempty" xml:"TbSellerId,omitempty"`
	UserPayFee   *int64  `json:"UserPayFee,omitempty" xml:"UserPayFee,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetItemId(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.ItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetLmItemId(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.LmItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetLmShopId(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.LmShopId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetNumber(v int32) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.Number = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetPoints(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.Points = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetPointsAmount(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.PointsAmount = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetPriceCent(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.PriceCent = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetRemoved(v bool) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.Removed = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetSkuId(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.SkuId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetTbSellerId(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.TbSellerId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetUserPayFee(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.UserPayFee = &v
	return s
}

type RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS struct {
	AvailableItems []*RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems `json:"AvailableItems,omitempty" xml:"AvailableItems,omitempty" type:"Repeated"`
	CanUse         *bool                                                                                    `json:"CanUse,omitempty" xml:"CanUse,omitempty"`
	DiscountPrice  *int64                                                                                   `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	ExpireTime     *int64                                                                                   `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId     *string                                                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Level          *string                                                                                  `json:"Level,omitempty" xml:"Level,omitempty"`
	LmItemId       *string                                                                                  `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	PromotionName  *string                                                                                  `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionType  *string                                                                                  `json:"PromotionType,omitempty" xml:"PromotionType,omitempty"`
	Reason         *string                                                                                  `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Selected       *bool                                                                                    `json:"Selected,omitempty" xml:"Selected,omitempty"`
	SkuIds         []*int64                                                                                 `json:"SkuIds,omitempty" xml:"SkuIds,omitempty" type:"Repeated"`
	SpecialPrice   *int64                                                                                   `json:"SpecialPrice,omitempty" xml:"SpecialPrice,omitempty"`
	SubBizCode     *string                                                                                  `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	TbSellerId     *int64                                                                                   `json:"TbSellerId,omitempty" xml:"TbSellerId,omitempty"`
	ThresholdPrice *int64                                                                                   `json:"ThresholdPrice,omitempty" xml:"ThresholdPrice,omitempty"`
	UseStartTime   *int64                                                                                   `json:"UseStartTime,omitempty" xml:"UseStartTime,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetAvailableItems(v []*RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.AvailableItems = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetCanUse(v bool) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.CanUse = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetDiscountPrice(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.DiscountPrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetExpireTime(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.ExpireTime = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetInstanceId(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.InstanceId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetLevel(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.Level = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetLmItemId(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.LmItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetPromotionName(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.PromotionName = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetPromotionType(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.PromotionType = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetReason(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.Reason = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetSelected(v bool) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.Selected = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetSkuIds(v []*int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.SkuIds = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetSpecialPrice(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.SpecialPrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetSubBizCode(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.SubBizCode = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetTbSellerId(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.TbSellerId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetThresholdPrice(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.ThresholdPrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS) SetUseStartTime(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOS {
	s.UseStartTime = &v
	return s
}

type RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems struct {
	ItemId       *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId     *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	LmShopId     *int64  `json:"LmShopId,omitempty" xml:"LmShopId,omitempty"`
	Number       *int32  `json:"Number,omitempty" xml:"Number,omitempty"`
	Points       *int64  `json:"Points,omitempty" xml:"Points,omitempty"`
	PointsAmount *int64  `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	PriceCent    *int64  `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	Removed      *bool   `json:"Removed,omitempty" xml:"Removed,omitempty"`
	SkuId        *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	TbSellerId   *int64  `json:"TbSellerId,omitempty" xml:"TbSellerId,omitempty"`
	UserPayFee   *int64  `json:"UserPayFee,omitempty" xml:"UserPayFee,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) SetItemId(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems {
	s.ItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) SetLmItemId(v string) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems {
	s.LmItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) SetLmShopId(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems {
	s.LmShopId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) SetNumber(v int32) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems {
	s.Number = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) SetPoints(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems {
	s.Points = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) SetPointsAmount(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems {
	s.PointsAmount = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) SetPriceCent(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems {
	s.PriceCent = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) SetRemoved(v bool) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems {
	s.Removed = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) SetSkuId(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems {
	s.SkuId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) SetTbSellerId(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems {
	s.TbSellerId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems) SetUserPayFee(v int64) *RenderDistributionOrderResponseBodyModelRenderOrderInfosShopPromInstVOSAvailableItems {
	s.UserPayFee = &v
	return s
}

type RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos struct {
	CanSell         *bool                                                                                `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	DeliveryInfos   []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos   `json:"DeliveryInfos,omitempty" xml:"DeliveryInfos,omitempty" type:"Repeated"`
	ExtInfo         map[string]*string                                                                   `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	InvoiceInfo     *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosInvoiceInfo       `json:"InvoiceInfo,omitempty" xml:"InvoiceInfo,omitempty" type:"Struct"`
	ItemInfos       []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos       `json:"ItemInfos,omitempty" xml:"ItemInfos,omitempty" type:"Repeated"`
	Message         *string                                                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	ShopPromInstVOS []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS `json:"ShopPromInstVOS,omitempty" xml:"ShopPromInstVOS,omitempty" type:"Repeated"`
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos) SetCanSell(v bool) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos {
	s.CanSell = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos) SetDeliveryInfos(v []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos {
	s.DeliveryInfos = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos) SetExtInfo(v map[string]*string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos {
	s.ExtInfo = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos) SetInvoiceInfo(v *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosInvoiceInfo) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos {
	s.InvoiceInfo = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos) SetItemInfos(v []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos {
	s.ItemInfos = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos) SetMessage(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos {
	s.Message = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos) SetShopPromInstVOS(v []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfos {
	s.ShopPromInstVOS = v
	return s
}

type RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	PostFee     *int64  `json:"PostFee,omitempty" xml:"PostFee,omitempty"`
	ServiceType *int64  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos) SetDisplayName(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos {
	s.DisplayName = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos) SetId(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos {
	s.Id = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos) SetPostFee(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos {
	s.PostFee = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos) SetServiceType(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosDeliveryInfos {
	s.ServiceType = &v
	return s
}

type RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosInvoiceInfo struct {
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosInvoiceInfo) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosInvoiceInfo) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosInvoiceInfo) SetDesc(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosInvoiceInfo {
	s.Desc = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosInvoiceInfo) SetType(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosInvoiceInfo {
	s.Type = &v
	return s
}

type RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos struct {
	CanSell                *bool                                                                                         `json:"CanSell,omitempty" xml:"CanSell,omitempty"`
	DistributionMallId     *string                                                                                       `json:"DistributionMallId,omitempty" xml:"DistributionMallId,omitempty"`
	DistributionSupplierId *string                                                                                       `json:"DistributionSupplierId,omitempty" xml:"DistributionSupplierId,omitempty"`
	DistributorId          *string                                                                                       `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	Features               map[string]*string                                                                            `json:"Features,omitempty" xml:"Features,omitempty"`
	ItemId                 *string                                                                                       `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName               *string                                                                                       `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	ItemPicUrl             *string                                                                                       `json:"ItemPicUrl,omitempty" xml:"ItemPicUrl,omitempty"`
	ItemPromInstVOS        []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS `json:"ItemPromInstVOS,omitempty" xml:"ItemPromInstVOS,omitempty" type:"Repeated"`
	ItemUrl                *string                                                                                       `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	Message                *string                                                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Price                  *int64                                                                                        `json:"Price,omitempty" xml:"Price,omitempty"`
	PromotionFee           *int64                                                                                        `json:"PromotionFee,omitempty" xml:"PromotionFee,omitempty"`
	Quantity               *int32                                                                                        `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ReservePrice           *int64                                                                                        `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	SkuId                  *int64                                                                                        `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	SkuName                *string                                                                                       `json:"SkuName,omitempty" xml:"SkuName,omitempty"`
	VirtualItemType        *string                                                                                       `json:"VirtualItemType,omitempty" xml:"VirtualItemType,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetCanSell(v bool) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.CanSell = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetDistributionMallId(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.DistributionMallId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetDistributionSupplierId(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.DistributionSupplierId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetDistributorId(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.DistributorId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetFeatures(v map[string]*string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.Features = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetItemId(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.ItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetItemName(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.ItemName = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetItemPicUrl(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.ItemPicUrl = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetItemPromInstVOS(v []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.ItemPromInstVOS = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetItemUrl(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.ItemUrl = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetMessage(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.Message = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetPrice(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.Price = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetPromotionFee(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.PromotionFee = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetQuantity(v int32) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.Quantity = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetReservePrice(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.ReservePrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetSkuId(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.SkuId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetSkuName(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.SkuName = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos) SetVirtualItemType(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfos {
	s.VirtualItemType = &v
	return s
}

type RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS struct {
	AvailableItems []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems `json:"AvailableItems,omitempty" xml:"AvailableItems,omitempty" type:"Repeated"`
	CanUse         *bool                                                                                                       `json:"CanUse,omitempty" xml:"CanUse,omitempty"`
	DiscountPrice  *int64                                                                                                      `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	ExpireTime     *int64                                                                                                      `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId     *string                                                                                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Level          *string                                                                                                     `json:"Level,omitempty" xml:"Level,omitempty"`
	LmItemId       *string                                                                                                     `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	PromotionName  *string                                                                                                     `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionType  *string                                                                                                     `json:"PromotionType,omitempty" xml:"PromotionType,omitempty"`
	Reason         *string                                                                                                     `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Selected       *bool                                                                                                       `json:"Selected,omitempty" xml:"Selected,omitempty"`
	SkuIds         []*int64                                                                                                    `json:"SkuIds,omitempty" xml:"SkuIds,omitempty" type:"Repeated"`
	SpecialPrice   *int64                                                                                                      `json:"SpecialPrice,omitempty" xml:"SpecialPrice,omitempty"`
	SubBizCode     *string                                                                                                     `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	TbSellerId     *int64                                                                                                      `json:"TbSellerId,omitempty" xml:"TbSellerId,omitempty"`
	ThresholdPrice *int64                                                                                                      `json:"ThresholdPrice,omitempty" xml:"ThresholdPrice,omitempty"`
	UseStartTime   *int64                                                                                                      `json:"UseStartTime,omitempty" xml:"UseStartTime,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetAvailableItems(v []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.AvailableItems = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetCanUse(v bool) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.CanUse = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetDiscountPrice(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.DiscountPrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetExpireTime(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.ExpireTime = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetInstanceId(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.InstanceId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetLevel(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.Level = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetLmItemId(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.LmItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetPromotionName(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.PromotionName = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetPromotionType(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.PromotionType = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetReason(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.Reason = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetSelected(v bool) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.Selected = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetSkuIds(v []*int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.SkuIds = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetSpecialPrice(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.SpecialPrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetSubBizCode(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.SubBizCode = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetTbSellerId(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.TbSellerId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetThresholdPrice(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.ThresholdPrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS) SetUseStartTime(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOS {
	s.UseStartTime = &v
	return s
}

type RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems struct {
	ItemId       *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId     *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	LmShopId     *int64  `json:"LmShopId,omitempty" xml:"LmShopId,omitempty"`
	Number       *int32  `json:"Number,omitempty" xml:"Number,omitempty"`
	Points       *int64  `json:"Points,omitempty" xml:"Points,omitempty"`
	PointsAmount *int64  `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	PriceCent    *int64  `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	Removed      *bool   `json:"Removed,omitempty" xml:"Removed,omitempty"`
	SkuId        *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	TbSellerId   *int64  `json:"TbSellerId,omitempty" xml:"TbSellerId,omitempty"`
	UserPayFee   *int64  `json:"UserPayFee,omitempty" xml:"UserPayFee,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetItemId(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.ItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetLmItemId(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.LmItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetLmShopId(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.LmShopId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetNumber(v int32) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.Number = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetPoints(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.Points = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetPointsAmount(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.PointsAmount = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetPriceCent(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.PriceCent = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetRemoved(v bool) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.Removed = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetSkuId(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.SkuId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetTbSellerId(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.TbSellerId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems) SetUserPayFee(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosItemInfosItemPromInstVOSAvailableItems {
	s.UserPayFee = &v
	return s
}

type RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS struct {
	AvailableItems []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems `json:"AvailableItems,omitempty" xml:"AvailableItems,omitempty" type:"Repeated"`
	CanUse         *bool                                                                                              `json:"CanUse,omitempty" xml:"CanUse,omitempty"`
	DiscountPrice  *int64                                                                                             `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	ExpireTime     *int64                                                                                             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId     *string                                                                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Level          *string                                                                                            `json:"Level,omitempty" xml:"Level,omitempty"`
	LmItemId       *string                                                                                            `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	PromotionName  *string                                                                                            `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionType  *string                                                                                            `json:"PromotionType,omitempty" xml:"PromotionType,omitempty"`
	Reason         *string                                                                                            `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Selected       *bool                                                                                              `json:"Selected,omitempty" xml:"Selected,omitempty"`
	SkuIds         []*int64                                                                                           `json:"SkuIds,omitempty" xml:"SkuIds,omitempty" type:"Repeated"`
	SpecialPrice   *int64                                                                                             `json:"SpecialPrice,omitempty" xml:"SpecialPrice,omitempty"`
	SubBizCode     *string                                                                                            `json:"SubBizCode,omitempty" xml:"SubBizCode,omitempty"`
	TbSellerId     *int64                                                                                             `json:"TbSellerId,omitempty" xml:"TbSellerId,omitempty"`
	ThresholdPrice *int64                                                                                             `json:"ThresholdPrice,omitempty" xml:"ThresholdPrice,omitempty"`
	UseStartTime   *int64                                                                                             `json:"UseStartTime,omitempty" xml:"UseStartTime,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetAvailableItems(v []*RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.AvailableItems = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetCanUse(v bool) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.CanUse = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetDiscountPrice(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.DiscountPrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetExpireTime(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.ExpireTime = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetInstanceId(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.InstanceId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetLevel(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.Level = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetLmItemId(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.LmItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetPromotionName(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.PromotionName = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetPromotionType(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.PromotionType = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetReason(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.Reason = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetSelected(v bool) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.Selected = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetSkuIds(v []*int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.SkuIds = v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetSpecialPrice(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.SpecialPrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetSubBizCode(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.SubBizCode = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetTbSellerId(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.TbSellerId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetThresholdPrice(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.ThresholdPrice = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS) SetUseStartTime(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOS {
	s.UseStartTime = &v
	return s
}

type RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems struct {
	ItemId       *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LmItemId     *string `json:"LmItemId,omitempty" xml:"LmItemId,omitempty"`
	LmShopId     *int64  `json:"LmShopId,omitempty" xml:"LmShopId,omitempty"`
	Number       *int32  `json:"Number,omitempty" xml:"Number,omitempty"`
	Points       *int64  `json:"Points,omitempty" xml:"Points,omitempty"`
	PointsAmount *int64  `json:"PointsAmount,omitempty" xml:"PointsAmount,omitempty"`
	PriceCent    *int64  `json:"PriceCent,omitempty" xml:"PriceCent,omitempty"`
	Removed      *bool   `json:"Removed,omitempty" xml:"Removed,omitempty"`
	SkuId        *int64  `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	TbSellerId   *int64  `json:"TbSellerId,omitempty" xml:"TbSellerId,omitempty"`
	UserPayFee   *int64  `json:"UserPayFee,omitempty" xml:"UserPayFee,omitempty"`
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) SetItemId(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems {
	s.ItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) SetLmItemId(v string) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems {
	s.LmItemId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) SetLmShopId(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems {
	s.LmShopId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) SetNumber(v int32) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems {
	s.Number = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) SetPoints(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems {
	s.Points = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) SetPointsAmount(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems {
	s.PointsAmount = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) SetPriceCent(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems {
	s.PriceCent = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) SetRemoved(v bool) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems {
	s.Removed = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) SetSkuId(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems {
	s.SkuId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) SetTbSellerId(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems {
	s.TbSellerId = &v
	return s
}

func (s *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems) SetUserPayFee(v int64) *RenderDistributionOrderResponseBodyModelUnsellableRenderOrderInfosShopPromInstVOSAvailableItems {
	s.UserPayFee = &v
	return s
}

type RenderDistributionOrderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenderDistributionOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenderDistributionOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s RenderDistributionOrderResponse) GoString() string {
	return s.String()
}

func (s *RenderDistributionOrderResponse) SetHeaders(v map[string]*string) *RenderDistributionOrderResponse {
	s.Headers = v
	return s
}

func (s *RenderDistributionOrderResponse) SetStatusCode(v int32) *RenderDistributionOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *RenderDistributionOrderResponse) SetBody(v *RenderDistributionOrderResponseBody) *RenderDistributionOrderResponse {
	s.Body = v
	return s
}

type SubmitReturnGoodLogistics4DistributionRequest struct {
	CpCode                 *string `json:"CpCode,omitempty" xml:"CpCode,omitempty"`
	DisputeId              *int64  `json:"DisputeId,omitempty" xml:"DisputeId,omitempty"`
	DistributorId          *string `json:"DistributorId,omitempty" xml:"DistributorId,omitempty"`
	LogisticsNo            *string `json:"LogisticsNo,omitempty" xml:"LogisticsNo,omitempty"`
	SubDistributionOrderId *string `json:"SubDistributionOrderId,omitempty" xml:"SubDistributionOrderId,omitempty"`
	TenantId               *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s SubmitReturnGoodLogistics4DistributionRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitReturnGoodLogistics4DistributionRequest) GoString() string {
	return s.String()
}

func (s *SubmitReturnGoodLogistics4DistributionRequest) SetCpCode(v string) *SubmitReturnGoodLogistics4DistributionRequest {
	s.CpCode = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionRequest) SetDisputeId(v int64) *SubmitReturnGoodLogistics4DistributionRequest {
	s.DisputeId = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionRequest) SetDistributorId(v string) *SubmitReturnGoodLogistics4DistributionRequest {
	s.DistributorId = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionRequest) SetLogisticsNo(v string) *SubmitReturnGoodLogistics4DistributionRequest {
	s.LogisticsNo = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionRequest) SetSubDistributionOrderId(v string) *SubmitReturnGoodLogistics4DistributionRequest {
	s.SubDistributionOrderId = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionRequest) SetTenantId(v string) *SubmitReturnGoodLogistics4DistributionRequest {
	s.TenantId = &v
	return s
}

type SubmitReturnGoodLogistics4DistributionResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	LogsId     *string `json:"LogsId,omitempty" xml:"LogsId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCode    *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
	SubMessage *string `json:"SubMessage,omitempty" xml:"SubMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SubmitReturnGoodLogistics4DistributionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitReturnGoodLogistics4DistributionResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitReturnGoodLogistics4DistributionResponseBody) SetCode(v string) *SubmitReturnGoodLogistics4DistributionResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionResponseBody) SetLogsId(v string) *SubmitReturnGoodLogistics4DistributionResponseBody {
	s.LogsId = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionResponseBody) SetMessage(v string) *SubmitReturnGoodLogistics4DistributionResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionResponseBody) SetPageNumber(v int64) *SubmitReturnGoodLogistics4DistributionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionResponseBody) SetPageSize(v int64) *SubmitReturnGoodLogistics4DistributionResponseBody {
	s.PageSize = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionResponseBody) SetRequestId(v string) *SubmitReturnGoodLogistics4DistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionResponseBody) SetSubCode(v string) *SubmitReturnGoodLogistics4DistributionResponseBody {
	s.SubCode = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionResponseBody) SetSubMessage(v string) *SubmitReturnGoodLogistics4DistributionResponseBody {
	s.SubMessage = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionResponseBody) SetSuccess(v bool) *SubmitReturnGoodLogistics4DistributionResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionResponseBody) SetTotalCount(v int64) *SubmitReturnGoodLogistics4DistributionResponseBody {
	s.TotalCount = &v
	return s
}

type SubmitReturnGoodLogistics4DistributionResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitReturnGoodLogistics4DistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitReturnGoodLogistics4DistributionResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitReturnGoodLogistics4DistributionResponse) GoString() string {
	return s.String()
}

func (s *SubmitReturnGoodLogistics4DistributionResponse) SetHeaders(v map[string]*string) *SubmitReturnGoodLogistics4DistributionResponse {
	s.Headers = v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionResponse) SetStatusCode(v int32) *SubmitReturnGoodLogistics4DistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitReturnGoodLogistics4DistributionResponse) SetBody(v *SubmitReturnGoodLogistics4DistributionResponseBody) *SubmitReturnGoodLogistics4DistributionResponse {
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

func (client *Client) ApplyCreateDistributionOrderWithOptions(tmpReq *ApplyCreateDistributionOrderRequest, runtime *util.RuntimeOptions) (_result *ApplyCreateDistributionOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ApplyCreateDistributionOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ItemInfoLists)) {
		request.ItemInfoListsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItemInfoLists, tea.String("ItemInfoLists"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuyerId)) {
		body["BuyerId"] = request.BuyerId
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryAddress)) {
		body["DeliveryAddress"] = request.DeliveryAddress
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionOutTradeId)) {
		body["DistributionOutTradeId"] = request.DistributionOutTradeId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionSupplierId)) {
		body["DistributionSupplierId"] = request.DistributionSupplierId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ItemInfoListsShrink)) {
		body["ItemInfoLists"] = request.ItemInfoListsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyCreateDistributionOrder"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyCreateDistributionOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyCreateDistributionOrder(request *ApplyCreateDistributionOrderRequest) (_result *ApplyCreateDistributionOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyCreateDistributionOrderResponse{}
	_body, _err := client.ApplyCreateDistributionOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyRefund4DistributionWithOptions(tmpReq *ApplyRefund4DistributionRequest, runtime *util.RuntimeOptions) (_result *ApplyRefund4DistributionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ApplyRefund4DistributionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LeavePictureLists)) {
		request.LeavePictureListsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LeavePictureLists, tea.String("LeavePictureLists"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyReasonTextId)) {
		body["ApplyReasonTextId"] = request.ApplyReasonTextId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyRefundCount)) {
		body["ApplyRefundCount"] = request.ApplyRefundCount
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyRefundFee)) {
		body["ApplyRefundFee"] = request.ApplyRefundFee
	}

	if !tea.BoolValue(util.IsUnset(request.BizClaimType)) {
		body["BizClaimType"] = request.BizClaimType
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.GoodsStatus)) {
		body["GoodsStatus"] = request.GoodsStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LeaveMessage)) {
		body["LeaveMessage"] = request.LeaveMessage
	}

	if !tea.BoolValue(util.IsUnset(request.LeavePictureListsShrink)) {
		body["LeavePictureLists"] = request.LeavePictureListsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SubDistributionOrderId)) {
		body["SubDistributionOrderId"] = request.SubDistributionOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyRefund4Distribution"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyRefund4DistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyRefund4Distribution(request *ApplyRefund4DistributionRequest) (_result *ApplyRefund4DistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyRefund4DistributionResponse{}
	_body, _err := client.ApplyRefund4DistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelDistributionTradeWithOptions(request *CancelDistributionTradeRequest, runtime *util.RuntimeOptions) (_result *CancelDistributionTradeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributionTradeId)) {
		body["DistributionTradeId"] = request.DistributionTradeId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelDistributionTrade"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelDistributionTradeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelDistributionTrade(request *CancelDistributionTradeRequest) (_result *CancelDistributionTradeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelDistributionTradeResponse{}
	_body, _err := client.CancelDistributionTradeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelRefund4DistributionWithOptions(request *CancelRefund4DistributionRequest, runtime *util.RuntimeOptions) (_result *CancelRefund4DistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisputeId)) {
		body["DisputeId"] = request.DisputeId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.SubDistributionOrderId)) {
		body["SubDistributionOrderId"] = request.SubDistributionOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelRefund4Distribution"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelRefund4DistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelRefund4Distribution(request *CancelRefund4DistributionRequest) (_result *CancelRefund4DistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelRefund4DistributionResponse{}
	_body, _err := client.CancelRefund4DistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfirmDisburse4DistributionWithOptions(request *ConfirmDisburse4DistributionRequest, runtime *util.RuntimeOptions) (_result *ConfirmDisburse4DistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributionTradeId)) {
		body["DistributionTradeId"] = request.DistributionTradeId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.MainDistributionOrderId)) {
		body["MainDistributionOrderId"] = request.MainDistributionOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfirmDisburse4Distribution"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfirmDisburse4DistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfirmDisburse4Distribution(request *ConfirmDisburse4DistributionRequest) (_result *ConfirmDisburse4DistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfirmDisburse4DistributionResponse{}
	_body, _err := client.ConfirmDisburse4DistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitApplyRefund4DistributionWithOptions(request *InitApplyRefund4DistributionRequest, runtime *util.RuntimeOptions) (_result *InitApplyRefund4DistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizClaimType)) {
		body["BizClaimType"] = request.BizClaimType
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.GoodsStatus)) {
		body["GoodsStatus"] = request.GoodsStatus
	}

	if !tea.BoolValue(util.IsUnset(request.SubDistributionOrderId)) {
		body["SubDistributionOrderId"] = request.SubDistributionOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InitApplyRefund4Distribution"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitApplyRefund4DistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitApplyRefund4Distribution(request *InitApplyRefund4DistributionRequest) (_result *InitApplyRefund4DistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitApplyRefund4DistributionResponse{}
	_body, _err := client.InitApplyRefund4DistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitModifyRefund4DistributionWithOptions(request *InitModifyRefund4DistributionRequest, runtime *util.RuntimeOptions) (_result *InitModifyRefund4DistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizClaimType)) {
		body["BizClaimType"] = request.BizClaimType
	}

	if !tea.BoolValue(util.IsUnset(request.DisputeId)) {
		body["DisputeId"] = request.DisputeId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.SubDistributionOrderId)) {
		body["SubDistributionOrderId"] = request.SubDistributionOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InitModifyRefund4Distribution"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitModifyRefund4DistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitModifyRefund4Distribution(request *InitModifyRefund4DistributionRequest) (_result *InitModifyRefund4DistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitModifyRefund4DistributionResponse{}
	_body, _err := client.InitModifyRefund4DistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDistributionItemWithOptions(request *ListDistributionItemRequest, runtime *util.RuntimeOptions) (_result *ListDistributionItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributionMallId)) {
		body["DistributionMallId"] = request.DistributionMallId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemStatus)) {
		body["ItemStatus"] = request.ItemStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LmItemId)) {
		body["LmItemId"] = request.LmItemId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDistributionItem"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDistributionItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDistributionItem(request *ListDistributionItemRequest) (_result *ListDistributionItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDistributionItemResponse{}
	_body, _err := client.ListDistributionItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDistributionItemWithoutCacheWithOptions(request *ListDistributionItemWithoutCacheRequest, runtime *util.RuntimeOptions) (_result *ListDistributionItemWithoutCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributionMallId)) {
		body["DistributionMallId"] = request.DistributionMallId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.ItemStatus)) {
		body["ItemStatus"] = request.ItemStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LmItemId)) {
		body["LmItemId"] = request.LmItemId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDistributionItemWithoutCache"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDistributionItemWithoutCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDistributionItemWithoutCache(request *ListDistributionItemWithoutCacheRequest) (_result *ListDistributionItemWithoutCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDistributionItemWithoutCacheResponse{}
	_body, _err := client.ListDistributionItemWithoutCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDistributionMallWithOptions(request *ListDistributionMallRequest, runtime *util.RuntimeOptions) (_result *ListDistributionMallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelSupplierId)) {
		body["ChannelSupplierId"] = request.ChannelSupplierId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionMallId)) {
		body["DistributionMallId"] = request.DistributionMallId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionMallName)) {
		body["DistributionMallName"] = request.DistributionMallName
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDistributionMall"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDistributionMallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDistributionMall(request *ListDistributionMallRequest) (_result *ListDistributionMallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDistributionMallResponse{}
	_body, _err := client.ListDistributionMallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyRefund4DistributionWithOptions(tmpReq *ModifyRefund4DistributionRequest, runtime *util.RuntimeOptions) (_result *ModifyRefund4DistributionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyRefund4DistributionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LeavePictureLists)) {
		request.LeavePictureListsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LeavePictureLists, tea.String("LeavePictureLists"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplyReasonTextId)) {
		body["ApplyReasonTextId"] = request.ApplyReasonTextId
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyRefundCount)) {
		body["ApplyRefundCount"] = request.ApplyRefundCount
	}

	if !tea.BoolValue(util.IsUnset(request.ApplyRefundFee)) {
		body["ApplyRefundFee"] = request.ApplyRefundFee
	}

	if !tea.BoolValue(util.IsUnset(request.BizClaimType)) {
		body["BizClaimType"] = request.BizClaimType
	}

	if !tea.BoolValue(util.IsUnset(request.DisputeId)) {
		body["DisputeId"] = request.DisputeId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.GoodsStatus)) {
		body["GoodsStatus"] = request.GoodsStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LeaveMessage)) {
		body["LeaveMessage"] = request.LeaveMessage
	}

	if !tea.BoolValue(util.IsUnset(request.LeavePictureListsShrink)) {
		body["LeavePictureLists"] = request.LeavePictureListsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SubDistributionOrderId)) {
		body["SubDistributionOrderId"] = request.SubDistributionOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyRefund4Distribution"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyRefund4DistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyRefund4Distribution(request *ModifyRefund4DistributionRequest) (_result *ModifyRefund4DistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyRefund4DistributionResponse{}
	_body, _err := client.ModifyRefund4DistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryChildDivisionCodeByIdWithOptions(request *QueryChildDivisionCodeByIdRequest, runtime *util.RuntimeOptions) (_result *QueryChildDivisionCodeByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.DivisionCode)) {
		body["DivisionCode"] = request.DivisionCode
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryChildDivisionCodeById"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryChildDivisionCodeByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryChildDivisionCodeById(request *QueryChildDivisionCodeByIdRequest) (_result *QueryChildDivisionCodeByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryChildDivisionCodeByIdResponse{}
	_body, _err := client.QueryChildDivisionCodeByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDistributionBillDetailWithOptions(request *QueryDistributionBillDetailRequest, runtime *util.RuntimeOptions) (_result *QueryDistributionBillDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillId)) {
		body["BillId"] = request.BillId
	}

	if !tea.BoolValue(util.IsUnset(request.BillPeriod)) {
		body["BillPeriod"] = request.BillPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.BillStatus)) {
		body["BillStatus"] = request.BillStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionMallId)) {
		body["DistributionMallId"] = request.DistributionMallId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionMallName)) {
		body["DistributionMallName"] = request.DistributionMallName
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDistributionBillDetail"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDistributionBillDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDistributionBillDetail(request *QueryDistributionBillDetailRequest) (_result *QueryDistributionBillDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDistributionBillDetailResponse{}
	_body, _err := client.QueryDistributionBillDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDistributionMallWithOptions(request *QueryDistributionMallRequest, runtime *util.RuntimeOptions) (_result *QueryDistributionMallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributionMallId)) {
		body["DistributionMallId"] = request.DistributionMallId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDistributionMall"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDistributionMallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDistributionMall(request *QueryDistributionMallRequest) (_result *QueryDistributionMallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDistributionMallResponse{}
	_body, _err := client.QueryDistributionMallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDistributionTradeStatusWithOptions(request *QueryDistributionTradeStatusRequest, runtime *util.RuntimeOptions) (_result *QueryDistributionTradeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributionSupplierId)) {
		body["DistributionSupplierId"] = request.DistributionSupplierId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionTradeId)) {
		body["DistributionTradeId"] = request.DistributionTradeId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDistributionTradeStatus"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDistributionTradeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDistributionTradeStatus(request *QueryDistributionTradeStatusRequest) (_result *QueryDistributionTradeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDistributionTradeStatusResponse{}
	_body, _err := client.QueryDistributionTradeStatusWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributionMallId)) {
		body["DistributionMallId"] = request.DistributionMallId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.LmItemId)) {
		body["LmItemId"] = request.LmItemId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryItemDetail"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryItemDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) QueryItemDetailWithDivisionWithOptions(request *QueryItemDetailWithDivisionRequest, runtime *util.RuntimeOptions) (_result *QueryItemDetailWithDivisionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributionMallId)) {
		body["DistributionMallId"] = request.DistributionMallId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.DivisionCode)) {
		body["DivisionCode"] = request.DivisionCode
	}

	if !tea.BoolValue(util.IsUnset(request.LmItemId)) {
		body["LmItemId"] = request.LmItemId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryItemDetailWithDivision"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryItemDetailWithDivisionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryItemDetailWithDivision(request *QueryItemDetailWithDivisionRequest) (_result *QueryItemDetailWithDivisionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryItemDetailWithDivisionResponse{}
	_body, _err := client.QueryItemDetailWithDivisionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryItemGuideRetailPriceWithOptions(tmpReq *QueryItemGuideRetailPriceRequest, runtime *util.RuntimeOptions) (_result *QueryItemGuideRetailPriceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryItemGuideRetailPriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.LmItemIds)) {
		request.LmItemIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LmItemIds, tea.String("LmItemIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributionMallId)) {
		body["DistributionMallId"] = request.DistributionMallId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.LmItemIdsShrink)) {
		body["LmItemIds"] = request.LmItemIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryItemGuideRetailPrice"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryItemGuideRetailPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryItemGuideRetailPrice(request *QueryItemGuideRetailPriceRequest) (_result *QueryItemGuideRetailPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryItemGuideRetailPriceResponse{}
	_body, _err := client.QueryItemGuideRetailPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryLogistics4DistributionWithOptions(request *QueryLogistics4DistributionRequest, runtime *util.RuntimeOptions) (_result *QueryLogistics4DistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.MainDistributionOrderId)) {
		body["MainDistributionOrderId"] = request.MainDistributionOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryLogistics4Distribution"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryLogistics4DistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryLogistics4Distribution(request *QueryLogistics4DistributionRequest) (_result *QueryLogistics4DistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryLogistics4DistributionResponse{}
	_body, _err := client.QueryLogistics4DistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMallCategoryListWithOptions(request *QueryMallCategoryListRequest, runtime *util.RuntimeOptions) (_result *QueryMallCategoryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionMallId)) {
		body["DistributionMallId"] = request.DistributionMallId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMallCategoryList"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMallCategoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMallCategoryList(request *QueryMallCategoryListRequest) (_result *QueryMallCategoryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMallCategoryListResponse{}
	_body, _err := client.QueryMallCategoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderDetail4DistributionWithOptions(request *QueryOrderDetail4DistributionRequest, runtime *util.RuntimeOptions) (_result *QueryOrderDetail4DistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.MainDistributionOrderId)) {
		body["MainDistributionOrderId"] = request.MainDistributionOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOrderDetail4Distribution"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrderDetail4DistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrderDetail4Distribution(request *QueryOrderDetail4DistributionRequest) (_result *QueryOrderDetail4DistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderDetail4DistributionResponse{}
	_body, _err := client.QueryOrderDetail4DistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderList4DistributionWithOptions(request *QueryOrderList4DistributionRequest, runtime *util.RuntimeOptions) (_result *QueryOrderList4DistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.FilterOption)) {
		body["FilterOption"] = request.FilterOption
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOrderList4Distribution"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrderList4DistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrderList4Distribution(request *QueryOrderList4DistributionRequest) (_result *QueryOrderList4DistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderList4DistributionResponse{}
	_body, _err := client.QueryOrderList4DistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRefundApplicationDetail4DistributionWithOptions(request *QueryRefundApplicationDetail4DistributionRequest, runtime *util.RuntimeOptions) (_result *QueryRefundApplicationDetail4DistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.SubDistributionOrderId)) {
		body["SubDistributionOrderId"] = request.SubDistributionOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRefundApplicationDetail4Distribution"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRefundApplicationDetail4DistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRefundApplicationDetail4Distribution(request *QueryRefundApplicationDetail4DistributionRequest) (_result *QueryRefundApplicationDetail4DistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRefundApplicationDetail4DistributionResponse{}
	_body, _err := client.QueryRefundApplicationDetail4DistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenderDistributionOrderWithOptions(tmpReq *RenderDistributionOrderRequest, runtime *util.RuntimeOptions) (_result *RenderDistributionOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RenderDistributionOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ItemInfoLists)) {
		request.ItemInfoListsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ItemInfoLists, tea.String("ItemInfoLists"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuyerId)) {
		body["BuyerId"] = request.BuyerId
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryAddress)) {
		body["DeliveryAddress"] = request.DeliveryAddress
	}

	if !tea.BoolValue(util.IsUnset(request.DistributionSupplierId)) {
		body["DistributionSupplierId"] = request.DistributionSupplierId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ItemInfoListsShrink)) {
		body["ItemInfoLists"] = request.ItemInfoListsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenderDistributionOrder"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenderDistributionOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenderDistributionOrder(request *RenderDistributionOrderRequest) (_result *RenderDistributionOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenderDistributionOrderResponse{}
	_body, _err := client.RenderDistributionOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitReturnGoodLogistics4DistributionWithOptions(request *SubmitReturnGoodLogistics4DistributionRequest, runtime *util.RuntimeOptions) (_result *SubmitReturnGoodLogistics4DistributionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CpCode)) {
		body["CpCode"] = request.CpCode
	}

	if !tea.BoolValue(util.IsUnset(request.DisputeId)) {
		body["DisputeId"] = request.DisputeId
	}

	if !tea.BoolValue(util.IsUnset(request.DistributorId)) {
		body["DistributorId"] = request.DistributorId
	}

	if !tea.BoolValue(util.IsUnset(request.LogisticsNo)) {
		body["LogisticsNo"] = request.LogisticsNo
	}

	if !tea.BoolValue(util.IsUnset(request.SubDistributionOrderId)) {
		body["SubDistributionOrderId"] = request.SubDistributionOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitReturnGoodLogistics4Distribution"),
		Version:     tea.String("2022-05-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitReturnGoodLogistics4DistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitReturnGoodLogistics4Distribution(request *SubmitReturnGoodLogistics4DistributionRequest) (_result *SubmitReturnGoodLogistics4DistributionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitReturnGoodLogistics4DistributionResponse{}
	_body, _err := client.SubmitReturnGoodLogistics4DistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
