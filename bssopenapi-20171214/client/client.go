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

type AllocateCostUnitResourceRequest struct {
	FromUnitUserId       *int64                                                 `json:"FromUnitUserId,omitempty" xml:"FromUnitUserId,omitempty"`
	FromUnitId           *int64                                                 `json:"FromUnitId,omitempty" xml:"FromUnitId,omitempty"`
	ToUnitUserId         *int64                                                 `json:"ToUnitUserId,omitempty" xml:"ToUnitUserId,omitempty"`
	ToUnitId             *int64                                                 `json:"ToUnitId,omitempty" xml:"ToUnitId,omitempty"`
	ResourceInstanceList []*AllocateCostUnitResourceRequestResourceInstanceList `json:"ResourceInstanceList,omitempty" xml:"ResourceInstanceList,omitempty" type:"Repeated"`
}

func (s AllocateCostUnitResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateCostUnitResourceRequest) GoString() string {
	return s.String()
}

func (s *AllocateCostUnitResourceRequest) SetFromUnitUserId(v int64) *AllocateCostUnitResourceRequest {
	s.FromUnitUserId = &v
	return s
}

func (s *AllocateCostUnitResourceRequest) SetFromUnitId(v int64) *AllocateCostUnitResourceRequest {
	s.FromUnitId = &v
	return s
}

func (s *AllocateCostUnitResourceRequest) SetToUnitUserId(v int64) *AllocateCostUnitResourceRequest {
	s.ToUnitUserId = &v
	return s
}

func (s *AllocateCostUnitResourceRequest) SetToUnitId(v int64) *AllocateCostUnitResourceRequest {
	s.ToUnitId = &v
	return s
}

func (s *AllocateCostUnitResourceRequest) SetResourceInstanceList(v []*AllocateCostUnitResourceRequestResourceInstanceList) *AllocateCostUnitResourceRequest {
	s.ResourceInstanceList = v
	return s
}

type AllocateCostUnitResourceRequestResourceInstanceList struct {
	ApportionCode  *string `json:"ApportionCode,omitempty" xml:"ApportionCode,omitempty"`
	CommodityCode  *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ResourceId     *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceUserId *int64  `json:"ResourceUserId,omitempty" xml:"ResourceUserId,omitempty"`
}

func (s AllocateCostUnitResourceRequestResourceInstanceList) String() string {
	return tea.Prettify(s)
}

func (s AllocateCostUnitResourceRequestResourceInstanceList) GoString() string {
	return s.String()
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) SetApportionCode(v string) *AllocateCostUnitResourceRequestResourceInstanceList {
	s.ApportionCode = &v
	return s
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) SetCommodityCode(v string) *AllocateCostUnitResourceRequestResourceInstanceList {
	s.CommodityCode = &v
	return s
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) SetResourceId(v string) *AllocateCostUnitResourceRequestResourceInstanceList {
	s.ResourceId = &v
	return s
}

func (s *AllocateCostUnitResourceRequestResourceInstanceList) SetResourceUserId(v int64) *AllocateCostUnitResourceRequestResourceInstanceList {
	s.ResourceUserId = &v
	return s
}

type AllocateCostUnitResourceResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *AllocateCostUnitResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AllocateCostUnitResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocateCostUnitResourceResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateCostUnitResourceResponseBody) SetRequestId(v string) *AllocateCostUnitResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBody) SetMessage(v string) *AllocateCostUnitResourceResponseBody {
	s.Message = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBody) SetCode(v string) *AllocateCostUnitResourceResponseBody {
	s.Code = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBody) SetSuccess(v bool) *AllocateCostUnitResourceResponseBody {
	s.Success = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBody) SetData(v *AllocateCostUnitResourceResponseBodyData) *AllocateCostUnitResourceResponseBody {
	s.Data = v
	return s
}

type AllocateCostUnitResourceResponseBodyData struct {
	IsSuccess    *bool  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	ToUnitUserId *int64 `json:"ToUnitUserId,omitempty" xml:"ToUnitUserId,omitempty"`
	ToUnitId     *int64 `json:"ToUnitId,omitempty" xml:"ToUnitId,omitempty"`
}

func (s AllocateCostUnitResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AllocateCostUnitResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AllocateCostUnitResourceResponseBodyData) SetIsSuccess(v bool) *AllocateCostUnitResourceResponseBodyData {
	s.IsSuccess = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBodyData) SetToUnitUserId(v int64) *AllocateCostUnitResourceResponseBodyData {
	s.ToUnitUserId = &v
	return s
}

func (s *AllocateCostUnitResourceResponseBodyData) SetToUnitId(v int64) *AllocateCostUnitResourceResponseBodyData {
	s.ToUnitId = &v
	return s
}

type AllocateCostUnitResourceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AllocateCostUnitResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AllocateCostUnitResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateCostUnitResourceResponse) GoString() string {
	return s.String()
}

func (s *AllocateCostUnitResourceResponse) SetHeaders(v map[string]*string) *AllocateCostUnitResourceResponse {
	s.Headers = v
	return s
}

func (s *AllocateCostUnitResourceResponse) SetBody(v *AllocateCostUnitResourceResponseBody) *AllocateCostUnitResourceResponse {
	s.Body = v
	return s
}

type ApplyInvoiceRequest struct {
	InvoiceAmount   *int64   `json:"InvoiceAmount,omitempty" xml:"InvoiceAmount,omitempty"`
	OwnerId         *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	CustomerId      *int64   `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	AddressId       *int64   `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	InvoicingType   *int32   `json:"InvoicingType,omitempty" xml:"InvoicingType,omitempty"`
	ProcessWay      *int32   `json:"ProcessWay,omitempty" xml:"ProcessWay,omitempty"`
	ApplyUserNick   *string  `json:"ApplyUserNick,omitempty" xml:"ApplyUserNick,omitempty"`
	SelectedIds     []*int64 `json:"SelectedIds,omitempty" xml:"SelectedIds,omitempty" type:"Repeated"`
	InvoiceByAmount *bool    `json:"InvoiceByAmount,omitempty" xml:"InvoiceByAmount,omitempty"`
}

func (s ApplyInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyInvoiceRequest) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceRequest) SetInvoiceAmount(v int64) *ApplyInvoiceRequest {
	s.InvoiceAmount = &v
	return s
}

func (s *ApplyInvoiceRequest) SetOwnerId(v int64) *ApplyInvoiceRequest {
	s.OwnerId = &v
	return s
}

func (s *ApplyInvoiceRequest) SetCustomerId(v int64) *ApplyInvoiceRequest {
	s.CustomerId = &v
	return s
}

func (s *ApplyInvoiceRequest) SetAddressId(v int64) *ApplyInvoiceRequest {
	s.AddressId = &v
	return s
}

func (s *ApplyInvoiceRequest) SetInvoicingType(v int32) *ApplyInvoiceRequest {
	s.InvoicingType = &v
	return s
}

func (s *ApplyInvoiceRequest) SetProcessWay(v int32) *ApplyInvoiceRequest {
	s.ProcessWay = &v
	return s
}

func (s *ApplyInvoiceRequest) SetApplyUserNick(v string) *ApplyInvoiceRequest {
	s.ApplyUserNick = &v
	return s
}

func (s *ApplyInvoiceRequest) SetSelectedIds(v []*int64) *ApplyInvoiceRequest {
	s.SelectedIds = v
	return s
}

func (s *ApplyInvoiceRequest) SetInvoiceByAmount(v bool) *ApplyInvoiceRequest {
	s.InvoiceByAmount = &v
	return s
}

type ApplyInvoiceResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *ApplyInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ApplyInvoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceResponseBody) SetRequestId(v string) *ApplyInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyInvoiceResponseBody) SetSuccess(v bool) *ApplyInvoiceResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyInvoiceResponseBody) SetCode(v string) *ApplyInvoiceResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyInvoiceResponseBody) SetMessage(v string) *ApplyInvoiceResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyInvoiceResponseBody) SetData(v *ApplyInvoiceResponseBodyData) *ApplyInvoiceResponseBody {
	s.Data = v
	return s
}

type ApplyInvoiceResponseBodyData struct {
	InvoiceApplyId *int64 `json:"InvoiceApplyId,omitempty" xml:"InvoiceApplyId,omitempty"`
}

func (s ApplyInvoiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ApplyInvoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceResponseBodyData) SetInvoiceApplyId(v int64) *ApplyInvoiceResponseBodyData {
	s.InvoiceApplyId = &v
	return s
}

type ApplyInvoiceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyInvoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyInvoiceResponse) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceResponse) SetHeaders(v map[string]*string) *ApplyInvoiceResponse {
	s.Headers = v
	return s
}

func (s *ApplyInvoiceResponse) SetBody(v *ApplyInvoiceResponseBody) *ApplyInvoiceResponse {
	s.Body = v
	return s
}

type CancelOrderRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CancelOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelOrderRequest) SetOrderId(v string) *CancelOrderRequest {
	s.OrderId = &v
	return s
}

func (s *CancelOrderRequest) SetOwnerId(v int64) *CancelOrderRequest {
	s.OwnerId = &v
	return s
}

type CancelOrderResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *CancelOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s *CancelOrderResponseBody) SetCode(v string) *CancelOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CancelOrderResponseBody) SetMessage(v string) *CancelOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CancelOrderResponseBody) SetData(v *CancelOrderResponseBodyData) *CancelOrderResponseBody {
	s.Data = v
	return s
}

type CancelOrderResponseBodyData struct {
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s CancelOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelOrderResponseBodyData) SetHostId(v string) *CancelOrderResponseBodyData {
	s.HostId = &v
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

type ChangeResellerConsumeAmountRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AdjustType   *string `json:"AdjustType,omitempty" xml:"AdjustType,omitempty"`
	Amount       *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Currency     *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
	OutBizId     *string `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
	ExtendMap    *string `json:"ExtendMap,omitempty" xml:"ExtendMap,omitempty"`
}

func (s ChangeResellerConsumeAmountRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResellerConsumeAmountRequest) GoString() string {
	return s.String()
}

func (s *ChangeResellerConsumeAmountRequest) SetOwnerId(v int64) *ChangeResellerConsumeAmountRequest {
	s.OwnerId = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetAdjustType(v string) *ChangeResellerConsumeAmountRequest {
	s.AdjustType = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetAmount(v string) *ChangeResellerConsumeAmountRequest {
	s.Amount = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetCurrency(v string) *ChangeResellerConsumeAmountRequest {
	s.Currency = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetBusinessType(v string) *ChangeResellerConsumeAmountRequest {
	s.BusinessType = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetSource(v string) *ChangeResellerConsumeAmountRequest {
	s.Source = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetOutBizId(v string) *ChangeResellerConsumeAmountRequest {
	s.OutBizId = &v
	return s
}

func (s *ChangeResellerConsumeAmountRequest) SetExtendMap(v string) *ChangeResellerConsumeAmountRequest {
	s.ExtendMap = &v
	return s
}

type ChangeResellerConsumeAmountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ChangeResellerConsumeAmountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResellerConsumeAmountResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResellerConsumeAmountResponseBody) SetRequestId(v string) *ChangeResellerConsumeAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResellerConsumeAmountResponseBody) SetCode(v string) *ChangeResellerConsumeAmountResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeResellerConsumeAmountResponseBody) SetMessage(v string) *ChangeResellerConsumeAmountResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeResellerConsumeAmountResponseBody) SetSuccess(v bool) *ChangeResellerConsumeAmountResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeResellerConsumeAmountResponseBody) SetData(v string) *ChangeResellerConsumeAmountResponseBody {
	s.Data = &v
	return s
}

type ChangeResellerConsumeAmountResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeResellerConsumeAmountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeResellerConsumeAmountResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResellerConsumeAmountResponse) GoString() string {
	return s.String()
}

func (s *ChangeResellerConsumeAmountResponse) SetHeaders(v map[string]*string) *ChangeResellerConsumeAmountResponse {
	s.Headers = v
	return s
}

func (s *ChangeResellerConsumeAmountResponse) SetBody(v *ChangeResellerConsumeAmountResponseBody) *ChangeResellerConsumeAmountResponse {
	s.Body = v
	return s
}

type ConvertChargeTypeRequest struct {
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	Period           *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ConvertChargeTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ConvertChargeTypeRequest) SetOwnerId(v int64) *ConvertChargeTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ConvertChargeTypeRequest) SetProductType(v string) *ConvertChargeTypeRequest {
	s.ProductType = &v
	return s
}

func (s *ConvertChargeTypeRequest) SetSubscriptionType(v string) *ConvertChargeTypeRequest {
	s.SubscriptionType = &v
	return s
}

func (s *ConvertChargeTypeRequest) SetPeriod(v int32) *ConvertChargeTypeRequest {
	s.Period = &v
	return s
}

func (s *ConvertChargeTypeRequest) SetProductCode(v string) *ConvertChargeTypeRequest {
	s.ProductCode = &v
	return s
}

func (s *ConvertChargeTypeRequest) SetInstanceId(v string) *ConvertChargeTypeRequest {
	s.InstanceId = &v
	return s
}

type ConvertChargeTypeResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *ConvertChargeTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ConvertChargeTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertChargeTypeResponseBody) SetRequestId(v string) *ConvertChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertChargeTypeResponseBody) SetSuccess(v bool) *ConvertChargeTypeResponseBody {
	s.Success = &v
	return s
}

func (s *ConvertChargeTypeResponseBody) SetCode(v string) *ConvertChargeTypeResponseBody {
	s.Code = &v
	return s
}

func (s *ConvertChargeTypeResponseBody) SetMessage(v string) *ConvertChargeTypeResponseBody {
	s.Message = &v
	return s
}

func (s *ConvertChargeTypeResponseBody) SetData(v *ConvertChargeTypeResponseBodyData) *ConvertChargeTypeResponseBody {
	s.Data = v
	return s
}

type ConvertChargeTypeResponseBodyData struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ConvertChargeTypeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ConvertChargeTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConvertChargeTypeResponseBodyData) SetOrderId(v string) *ConvertChargeTypeResponseBodyData {
	s.OrderId = &v
	return s
}

type ConvertChargeTypeResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConvertChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConvertChargeTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ConvertChargeTypeResponse) SetHeaders(v map[string]*string) *ConvertChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ConvertChargeTypeResponse) SetBody(v *ConvertChargeTypeResponseBody) *ConvertChargeTypeResponse {
	s.Body = v
	return s
}

type CreateAgAccountRequest struct {
	LoginEmail     *string `json:"LoginEmail,omitempty" xml:"LoginEmail,omitempty"`
	AccountAttr    *string `json:"AccountAttr,omitempty" xml:"AccountAttr,omitempty"`
	EnterpriseName *string `json:"EnterpriseName,omitempty" xml:"EnterpriseName,omitempty"`
	FirstName      *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	LastName       *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	NationCode     *string `json:"NationCode,omitempty" xml:"NationCode,omitempty"`
	ProvinceName   *string `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	CityName       *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	Postcode       *string `json:"Postcode,omitempty" xml:"Postcode,omitempty"`
}

func (s CreateAgAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAgAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAgAccountRequest) SetLoginEmail(v string) *CreateAgAccountRequest {
	s.LoginEmail = &v
	return s
}

func (s *CreateAgAccountRequest) SetAccountAttr(v string) *CreateAgAccountRequest {
	s.AccountAttr = &v
	return s
}

func (s *CreateAgAccountRequest) SetEnterpriseName(v string) *CreateAgAccountRequest {
	s.EnterpriseName = &v
	return s
}

func (s *CreateAgAccountRequest) SetFirstName(v string) *CreateAgAccountRequest {
	s.FirstName = &v
	return s
}

func (s *CreateAgAccountRequest) SetLastName(v string) *CreateAgAccountRequest {
	s.LastName = &v
	return s
}

func (s *CreateAgAccountRequest) SetNationCode(v string) *CreateAgAccountRequest {
	s.NationCode = &v
	return s
}

func (s *CreateAgAccountRequest) SetProvinceName(v string) *CreateAgAccountRequest {
	s.ProvinceName = &v
	return s
}

func (s *CreateAgAccountRequest) SetCityName(v string) *CreateAgAccountRequest {
	s.CityName = &v
	return s
}

func (s *CreateAgAccountRequest) SetPostcode(v string) *CreateAgAccountRequest {
	s.Postcode = &v
	return s
}

type CreateAgAccountResponseBody struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code          *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Success       *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	AgRelationDto *CreateAgAccountResponseBodyAgRelationDto `json:"AgRelationDto,omitempty" xml:"AgRelationDto,omitempty" type:"Struct"`
}

func (s CreateAgAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAgAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgAccountResponseBody) SetRequestId(v string) *CreateAgAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgAccountResponseBody) SetCode(v string) *CreateAgAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgAccountResponseBody) SetMessage(v string) *CreateAgAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgAccountResponseBody) SetSuccess(v bool) *CreateAgAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgAccountResponseBody) SetAgRelationDto(v *CreateAgAccountResponseBodyAgRelationDto) *CreateAgAccountResponseBody {
	s.AgRelationDto = v
	return s
}

type CreateAgAccountResponseBodyAgRelationDto struct {
	Pk               *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Mpk              *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	RamAdminRoleName *string `json:"RamAdminRoleName,omitempty" xml:"RamAdminRoleName,omitempty"`
}

func (s CreateAgAccountResponseBodyAgRelationDto) String() string {
	return tea.Prettify(s)
}

func (s CreateAgAccountResponseBodyAgRelationDto) GoString() string {
	return s.String()
}

func (s *CreateAgAccountResponseBodyAgRelationDto) SetPk(v string) *CreateAgAccountResponseBodyAgRelationDto {
	s.Pk = &v
	return s
}

func (s *CreateAgAccountResponseBodyAgRelationDto) SetType(v string) *CreateAgAccountResponseBodyAgRelationDto {
	s.Type = &v
	return s
}

func (s *CreateAgAccountResponseBodyAgRelationDto) SetMpk(v string) *CreateAgAccountResponseBodyAgRelationDto {
	s.Mpk = &v
	return s
}

func (s *CreateAgAccountResponseBodyAgRelationDto) SetRamAdminRoleName(v string) *CreateAgAccountResponseBodyAgRelationDto {
	s.RamAdminRoleName = &v
	return s
}

type CreateAgAccountResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAgAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAgAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAgAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAgAccountResponse) SetHeaders(v map[string]*string) *CreateAgAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAgAccountResponse) SetBody(v *CreateAgAccountResponseBody) *CreateAgAccountResponse {
	s.Body = v
	return s
}

type CreateCostUnitRequest struct {
	UnitEntityList []*CreateCostUnitRequestUnitEntityList `json:"UnitEntityList,omitempty" xml:"UnitEntityList,omitempty" type:"Repeated"`
}

func (s CreateCostUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCostUnitRequest) GoString() string {
	return s.String()
}

func (s *CreateCostUnitRequest) SetUnitEntityList(v []*CreateCostUnitRequestUnitEntityList) *CreateCostUnitRequest {
	s.UnitEntityList = v
	return s
}

type CreateCostUnitRequestUnitEntityList struct {
	OwnerUid     *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	ParentUnitId *int64  `json:"ParentUnitId,omitempty" xml:"ParentUnitId,omitempty"`
	UnitName     *string `json:"UnitName,omitempty" xml:"UnitName,omitempty"`
}

func (s CreateCostUnitRequestUnitEntityList) String() string {
	return tea.Prettify(s)
}

func (s CreateCostUnitRequestUnitEntityList) GoString() string {
	return s.String()
}

func (s *CreateCostUnitRequestUnitEntityList) SetOwnerUid(v int64) *CreateCostUnitRequestUnitEntityList {
	s.OwnerUid = &v
	return s
}

func (s *CreateCostUnitRequestUnitEntityList) SetParentUnitId(v int64) *CreateCostUnitRequestUnitEntityList {
	s.ParentUnitId = &v
	return s
}

func (s *CreateCostUnitRequestUnitEntityList) SetUnitName(v string) *CreateCostUnitRequestUnitEntityList {
	s.UnitName = &v
	return s
}

type CreateCostUnitResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *CreateCostUnitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CreateCostUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCostUnitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCostUnitResponseBody) SetRequestId(v string) *CreateCostUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCostUnitResponseBody) SetMessage(v string) *CreateCostUnitResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCostUnitResponseBody) SetCode(v string) *CreateCostUnitResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCostUnitResponseBody) SetSuccess(v bool) *CreateCostUnitResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCostUnitResponseBody) SetData(v *CreateCostUnitResponseBodyData) *CreateCostUnitResponseBody {
	s.Data = v
	return s
}

type CreateCostUnitResponseBodyData struct {
	CostUnitDtoList []*CreateCostUnitResponseBodyDataCostUnitDtoList `json:"CostUnitDtoList,omitempty" xml:"CostUnitDtoList,omitempty" type:"Repeated"`
}

func (s CreateCostUnitResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateCostUnitResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCostUnitResponseBodyData) SetCostUnitDtoList(v []*CreateCostUnitResponseBodyDataCostUnitDtoList) *CreateCostUnitResponseBodyData {
	s.CostUnitDtoList = v
	return s
}

type CreateCostUnitResponseBodyDataCostUnitDtoList struct {
	UnitId       *int64  `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
	ParentUnitId *int64  `json:"ParentUnitId,omitempty" xml:"ParentUnitId,omitempty"`
	OwnerUid     *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	UnitName     *string `json:"UnitName,omitempty" xml:"UnitName,omitempty"`
}

func (s CreateCostUnitResponseBodyDataCostUnitDtoList) String() string {
	return tea.Prettify(s)
}

func (s CreateCostUnitResponseBodyDataCostUnitDtoList) GoString() string {
	return s.String()
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) SetUnitId(v int64) *CreateCostUnitResponseBodyDataCostUnitDtoList {
	s.UnitId = &v
	return s
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) SetParentUnitId(v int64) *CreateCostUnitResponseBodyDataCostUnitDtoList {
	s.ParentUnitId = &v
	return s
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) SetOwnerUid(v int64) *CreateCostUnitResponseBodyDataCostUnitDtoList {
	s.OwnerUid = &v
	return s
}

func (s *CreateCostUnitResponseBodyDataCostUnitDtoList) SetUnitName(v string) *CreateCostUnitResponseBodyDataCostUnitDtoList {
	s.UnitName = &v
	return s
}

type CreateCostUnitResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCostUnitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCostUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCostUnitResponse) GoString() string {
	return s.String()
}

func (s *CreateCostUnitResponse) SetHeaders(v map[string]*string) *CreateCostUnitResponse {
	s.Headers = v
	return s
}

func (s *CreateCostUnitResponse) SetBody(v *CreateCostUnitResponseBody) *CreateCostUnitResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	ProductCode      *string                           `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Parameter        []*CreateInstanceRequestParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
	OwnerId          *int64                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductType      *string                           `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string                           `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	Period           *int32                            `json:"Period,omitempty" xml:"Period,omitempty"`
	RenewalStatus    *string                           `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	RenewPeriod      *int32                            `json:"RenewPeriod,omitempty" xml:"RenewPeriod,omitempty"`
	ClientToken      *string                           `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Logistics        *string                           `json:"Logistics,omitempty" xml:"Logistics,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetProductCode(v string) *CreateInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateInstanceRequest) SetParameter(v []*CreateInstanceRequestParameter) *CreateInstanceRequest {
	s.Parameter = v
	return s
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetProductType(v string) *CreateInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *CreateInstanceRequest) SetSubscriptionType(v string) *CreateInstanceRequest {
	s.SubscriptionType = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int32) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetRenewalStatus(v string) *CreateInstanceRequest {
	s.RenewalStatus = &v
	return s
}

func (s *CreateInstanceRequest) SetRenewPeriod(v int32) *CreateInstanceRequest {
	s.RenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) SetLogistics(v string) *CreateInstanceRequest {
	s.Logistics = &v
	return s
}

type CreateInstanceRequestParameter struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestParameter) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestParameter) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestParameter) SetCode(v string) *CreateInstanceRequestParameter {
	s.Code = &v
	return s
}

func (s *CreateInstanceRequestParameter) SetValue(v string) *CreateInstanceRequestParameter {
	s.Value = &v
	return s
}

type CreateInstanceResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *CreateInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetSuccess(v bool) *CreateInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceResponseBody) SetCode(v string) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetMessage(v string) *CreateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBody) SetData(v *CreateInstanceResponseBodyData) *CreateInstanceResponseBody {
	s.Data = v
	return s
}

type CreateInstanceResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyData) SetInstanceId(v string) *CreateInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetOrderId(v string) *CreateInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateResellerUserQuotaRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Amount   *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	OutBizId *string `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
}

func (s CreateResellerUserQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResellerUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *CreateResellerUserQuotaRequest) SetOwnerId(v int64) *CreateResellerUserQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateResellerUserQuotaRequest) SetAmount(v string) *CreateResellerUserQuotaRequest {
	s.Amount = &v
	return s
}

func (s *CreateResellerUserQuotaRequest) SetCurrency(v string) *CreateResellerUserQuotaRequest {
	s.Currency = &v
	return s
}

func (s *CreateResellerUserQuotaRequest) SetOutBizId(v string) *CreateResellerUserQuotaRequest {
	s.OutBizId = &v
	return s
}

type CreateResellerUserQuotaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s CreateResellerUserQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResellerUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResellerUserQuotaResponseBody) SetRequestId(v string) *CreateResellerUserQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResellerUserQuotaResponseBody) SetCode(v string) *CreateResellerUserQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *CreateResellerUserQuotaResponseBody) SetMessage(v string) *CreateResellerUserQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *CreateResellerUserQuotaResponseBody) SetSuccess(v bool) *CreateResellerUserQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *CreateResellerUserQuotaResponseBody) SetData(v bool) *CreateResellerUserQuotaResponseBody {
	s.Data = &v
	return s
}

type CreateResellerUserQuotaResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateResellerUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateResellerUserQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResellerUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *CreateResellerUserQuotaResponse) SetHeaders(v map[string]*string) *CreateResellerUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *CreateResellerUserQuotaResponse) SetBody(v *CreateResellerUserQuotaResponseBody) *CreateResellerUserQuotaResponse {
	s.Body = v
	return s
}

type CreateResourcePackageRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductCode   *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PackageType   *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	EffectiveDate *string `json:"EffectiveDate,omitempty" xml:"EffectiveDate,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Duration      *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PricingCycle  *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s CreateResourcePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *CreateResourcePackageRequest) SetOwnerId(v int64) *CreateResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateResourcePackageRequest) SetProductCode(v string) *CreateResourcePackageRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateResourcePackageRequest) SetPackageType(v string) *CreateResourcePackageRequest {
	s.PackageType = &v
	return s
}

func (s *CreateResourcePackageRequest) SetEffectiveDate(v string) *CreateResourcePackageRequest {
	s.EffectiveDate = &v
	return s
}

func (s *CreateResourcePackageRequest) SetSpecification(v string) *CreateResourcePackageRequest {
	s.Specification = &v
	return s
}

func (s *CreateResourcePackageRequest) SetDuration(v int32) *CreateResourcePackageRequest {
	s.Duration = &v
	return s
}

func (s *CreateResourcePackageRequest) SetPricingCycle(v string) *CreateResourcePackageRequest {
	s.PricingCycle = &v
	return s
}

type CreateResourcePackageResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *int64                                 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *CreateResourcePackageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CreateResourcePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourcePackageResponseBody) SetRequestId(v string) *CreateResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourcePackageResponseBody) SetOrderId(v int64) *CreateResourcePackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateResourcePackageResponseBody) SetSuccess(v bool) *CreateResourcePackageResponseBody {
	s.Success = &v
	return s
}

func (s *CreateResourcePackageResponseBody) SetCode(v string) *CreateResourcePackageResponseBody {
	s.Code = &v
	return s
}

func (s *CreateResourcePackageResponseBody) SetMessage(v string) *CreateResourcePackageResponseBody {
	s.Message = &v
	return s
}

func (s *CreateResourcePackageResponseBody) SetData(v *CreateResourcePackageResponseBodyData) *CreateResourcePackageResponseBody {
	s.Data = v
	return s
}

type CreateResourcePackageResponseBodyData struct {
	OrderId    *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateResourcePackageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateResourcePackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateResourcePackageResponseBodyData) SetOrderId(v int64) *CreateResourcePackageResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateResourcePackageResponseBodyData) SetInstanceId(v string) *CreateResourcePackageResponseBodyData {
	s.InstanceId = &v
	return s
}

type CreateResourcePackageResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateResourcePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *CreateResourcePackageResponse) SetHeaders(v map[string]*string) *CreateResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *CreateResourcePackageResponse) SetBody(v *CreateResourcePackageResponseBody) *CreateResourcePackageResponse {
	s.Body = v
	return s
}

type DeleteCostUnitRequest struct {
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	UnitId   *int64 `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
}

func (s DeleteCostUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCostUnitRequest) GoString() string {
	return s.String()
}

func (s *DeleteCostUnitRequest) SetOwnerUid(v int64) *DeleteCostUnitRequest {
	s.OwnerUid = &v
	return s
}

func (s *DeleteCostUnitRequest) SetUnitId(v int64) *DeleteCostUnitRequest {
	s.UnitId = &v
	return s
}

type DeleteCostUnitResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *DeleteCostUnitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DeleteCostUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCostUnitResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCostUnitResponseBody) SetRequestId(v string) *DeleteCostUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCostUnitResponseBody) SetMessage(v string) *DeleteCostUnitResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCostUnitResponseBody) SetCode(v string) *DeleteCostUnitResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCostUnitResponseBody) SetSuccess(v bool) *DeleteCostUnitResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCostUnitResponseBody) SetData(v *DeleteCostUnitResponseBodyData) *DeleteCostUnitResponseBody {
	s.Data = v
	return s
}

type DeleteCostUnitResponseBodyData struct {
	IsSuccess *bool  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	UnitId    *int64 `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
	OwnerUid  *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
}

func (s DeleteCostUnitResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteCostUnitResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteCostUnitResponseBodyData) SetIsSuccess(v bool) *DeleteCostUnitResponseBodyData {
	s.IsSuccess = &v
	return s
}

func (s *DeleteCostUnitResponseBodyData) SetUnitId(v int64) *DeleteCostUnitResponseBodyData {
	s.UnitId = &v
	return s
}

func (s *DeleteCostUnitResponseBodyData) SetOwnerUid(v int64) *DeleteCostUnitResponseBodyData {
	s.OwnerUid = &v
	return s
}

type DeleteCostUnitResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCostUnitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCostUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCostUnitResponse) GoString() string {
	return s.String()
}

func (s *DeleteCostUnitResponse) SetHeaders(v map[string]*string) *DeleteCostUnitResponse {
	s.Headers = v
	return s
}

func (s *DeleteCostUnitResponse) SetBody(v *DeleteCostUnitResponseBody) *DeleteCostUnitResponse {
	s.Body = v
	return s
}

type DescribePricingModuleRequest struct {
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s DescribePricingModuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePricingModuleRequest) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleRequest) SetOwnerId(v int64) *DescribePricingModuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePricingModuleRequest) SetProductCode(v string) *DescribePricingModuleRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribePricingModuleRequest) SetProductType(v string) *DescribePricingModuleRequest {
	s.ProductType = &v
	return s
}

func (s *DescribePricingModuleRequest) SetSubscriptionType(v string) *DescribePricingModuleRequest {
	s.SubscriptionType = &v
	return s
}

type DescribePricingModuleResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *DescribePricingModuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribePricingModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePricingModuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBody) SetRequestId(v string) *DescribePricingModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePricingModuleResponseBody) SetSuccess(v bool) *DescribePricingModuleResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePricingModuleResponseBody) SetCode(v string) *DescribePricingModuleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePricingModuleResponseBody) SetMessage(v string) *DescribePricingModuleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePricingModuleResponseBody) SetData(v *DescribePricingModuleResponseBodyData) *DescribePricingModuleResponseBody {
	s.Data = v
	return s
}

type DescribePricingModuleResponseBodyData struct {
	ModuleList    *DescribePricingModuleResponseBodyDataModuleList    `json:"ModuleList,omitempty" xml:"ModuleList,omitempty" type:"Struct"`
	AttributeList *DescribePricingModuleResponseBodyDataAttributeList `json:"AttributeList,omitempty" xml:"AttributeList,omitempty" type:"Struct"`
}

func (s DescribePricingModuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribePricingModuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyData) SetModuleList(v *DescribePricingModuleResponseBodyDataModuleList) *DescribePricingModuleResponseBodyData {
	s.ModuleList = v
	return s
}

func (s *DescribePricingModuleResponseBodyData) SetAttributeList(v *DescribePricingModuleResponseBodyDataAttributeList) *DescribePricingModuleResponseBodyData {
	s.AttributeList = v
	return s
}

type DescribePricingModuleResponseBodyDataModuleList struct {
	Module []*DescribePricingModuleResponseBodyDataModuleListModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
}

func (s DescribePricingModuleResponseBodyDataModuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataModuleList) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataModuleList) SetModule(v []*DescribePricingModuleResponseBodyDataModuleListModule) *DescribePricingModuleResponseBodyDataModuleList {
	s.Module = v
	return s
}

type DescribePricingModuleResponseBodyDataModuleListModule struct {
	ModuleCode *string                                                          `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleName *string                                                          `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	PriceType  *string                                                          `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
	Currency   *string                                                          `json:"Currency,omitempty" xml:"Currency,omitempty"`
	ConfigList *DescribePricingModuleResponseBodyDataModuleListModuleConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Struct"`
}

func (s DescribePricingModuleResponseBodyDataModuleListModule) String() string {
	return tea.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataModuleListModule) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) SetModuleCode(v string) *DescribePricingModuleResponseBodyDataModuleListModule {
	s.ModuleCode = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) SetModuleName(v string) *DescribePricingModuleResponseBodyDataModuleListModule {
	s.ModuleName = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) SetPriceType(v string) *DescribePricingModuleResponseBodyDataModuleListModule {
	s.PriceType = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) SetCurrency(v string) *DescribePricingModuleResponseBodyDataModuleListModule {
	s.Currency = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataModuleListModule) SetConfigList(v *DescribePricingModuleResponseBodyDataModuleListModuleConfigList) *DescribePricingModuleResponseBodyDataModuleListModule {
	s.ConfigList = v
	return s
}

type DescribePricingModuleResponseBodyDataModuleListModuleConfigList struct {
	ConfigList []*string `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
}

func (s DescribePricingModuleResponseBodyDataModuleListModuleConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataModuleListModuleConfigList) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataModuleListModuleConfigList) SetConfigList(v []*string) *DescribePricingModuleResponseBodyDataModuleListModuleConfigList {
	s.ConfigList = v
	return s
}

type DescribePricingModuleResponseBodyDataAttributeList struct {
	Attribute []*DescribePricingModuleResponseBodyDataAttributeListAttribute `json:"Attribute,omitempty" xml:"Attribute,omitempty" type:"Repeated"`
}

func (s DescribePricingModuleResponseBodyDataAttributeList) String() string {
	return tea.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataAttributeList) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataAttributeList) SetAttribute(v []*DescribePricingModuleResponseBodyDataAttributeListAttribute) *DescribePricingModuleResponseBodyDataAttributeList {
	s.Attribute = v
	return s
}

type DescribePricingModuleResponseBodyDataAttributeListAttribute struct {
	Code   *string                                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Name   *string                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Unit   *string                                                            `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Values *DescribePricingModuleResponseBodyDataAttributeListAttributeValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
}

func (s DescribePricingModuleResponseBodyDataAttributeListAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataAttributeListAttribute) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) SetCode(v string) *DescribePricingModuleResponseBodyDataAttributeListAttribute {
	s.Code = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) SetName(v string) *DescribePricingModuleResponseBodyDataAttributeListAttribute {
	s.Name = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) SetUnit(v string) *DescribePricingModuleResponseBodyDataAttributeListAttribute {
	s.Unit = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttribute) SetValues(v *DescribePricingModuleResponseBodyDataAttributeListAttributeValues) *DescribePricingModuleResponseBodyDataAttributeListAttribute {
	s.Values = v
	return s
}

type DescribePricingModuleResponseBodyDataAttributeListAttributeValues struct {
	AttributeValue []*DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty" type:"Repeated"`
}

func (s DescribePricingModuleResponseBodyDataAttributeListAttributeValues) String() string {
	return tea.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataAttributeListAttributeValues) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValues) SetAttributeValue(v []*DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) *DescribePricingModuleResponseBodyDataAttributeListAttributeValues {
	s.AttributeValue = v
	return s
}

type DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue struct {
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) String() string {
	return tea.Prettify(s)
}

func (s DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) SetType(v string) *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue {
	s.Type = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) SetName(v string) *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue {
	s.Name = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) SetValue(v string) *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue {
	s.Value = &v
	return s
}

func (s *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue) SetRemark(v string) *DescribePricingModuleResponseBodyDataAttributeListAttributeValuesAttributeValue {
	s.Remark = &v
	return s
}

type DescribePricingModuleResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePricingModuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePricingModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePricingModuleResponse) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleResponse) SetHeaders(v map[string]*string) *DescribePricingModuleResponse {
	s.Headers = v
	return s
}

func (s *DescribePricingModuleResponse) SetBody(v *DescribePricingModuleResponseBody) *DescribePricingModuleResponse {
	s.Body = v
	return s
}

type DescribeResourcePackageProductRequest struct {
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribeResourcePackageProductRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductRequest) SetProductCode(v string) *DescribeResourcePackageProductRequest {
	s.ProductCode = &v
	return s
}

type DescribeResourcePackageProductResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *int64                                          `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *DescribeResourcePackageProductResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeResourcePackageProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBody) SetRequestId(v string) *DescribeResourcePackageProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBody) SetOrderId(v int64) *DescribeResourcePackageProductResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBody) SetSuccess(v bool) *DescribeResourcePackageProductResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBody) SetCode(v string) *DescribeResourcePackageProductResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBody) SetMessage(v string) *DescribeResourcePackageProductResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBody) SetData(v *DescribeResourcePackageProductResponseBodyData) *DescribeResourcePackageProductResponseBody {
	s.Data = v
	return s
}

type DescribeResourcePackageProductResponseBodyData struct {
	ResourcePackages *DescribeResourcePackageProductResponseBodyDataResourcePackages `json:"ResourcePackages,omitempty" xml:"ResourcePackages,omitempty" type:"Struct"`
}

func (s DescribeResourcePackageProductResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyData) SetResourcePackages(v *DescribeResourcePackageProductResponseBodyDataResourcePackages) *DescribeResourcePackageProductResponseBodyData {
	s.ResourcePackages = v
	return s
}

type DescribeResourcePackageProductResponseBodyDataResourcePackages struct {
	ResourcePackage []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage `json:"ResourcePackage,omitempty" xml:"ResourcePackage,omitempty" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackages) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackages) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackages) SetResourcePackage(v []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) *DescribeResourcePackageProductResponseBodyDataResourcePackages {
	s.ResourcePackage = v
	return s
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage struct {
	ProductCode  *string                                                                                    `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType  *string                                                                                    `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Name         *string                                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	PackageTypes *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes `json:"PackageTypes,omitempty" xml:"PackageTypes,omitempty" type:"Struct"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) SetProductCode(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage {
	s.ProductCode = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) SetProductType(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage {
	s.ProductType = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) SetName(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage {
	s.Name = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage) SetPackageTypes(v *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackage {
	s.PackageTypes = v
	return s
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes struct {
	PackageType []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType `json:"PackageType,omitempty" xml:"PackageType,omitempty" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes) SetPackageType(v []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypes {
	s.PackageType = v
	return s
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType struct {
	Name           *string                                                                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Code           *string                                                                                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Properties     *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties     `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
	Specifications *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications `json:"Specifications,omitempty" xml:"Specifications,omitempty" type:"Struct"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) SetName(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType {
	s.Name = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) SetCode(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType {
	s.Code = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) SetProperties(v *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType {
	s.Properties = v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType) SetSpecifications(v *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageType {
	s.Specifications = v
	return s
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties struct {
	Property []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties) SetProperty(v []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeProperties {
	s.Property = v
	return s
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) SetName(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty {
	s.Name = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty) SetValue(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypePropertiesProperty {
	s.Value = &v
	return s
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications struct {
	Specification []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification `json:"Specification,omitempty" xml:"Specification,omitempty" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications) SetSpecification(v []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecifications {
	s.Specification = v
	return s
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification struct {
	Name               *string                                                                                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Value              *string                                                                                                                                            `json:"Value,omitempty" xml:"Value,omitempty"`
	AvailableDurations *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations `json:"AvailableDurations,omitempty" xml:"AvailableDurations,omitempty" type:"Struct"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) SetName(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification {
	s.Name = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) SetValue(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification {
	s.Value = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification) SetAvailableDurations(v *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecification {
	s.AvailableDurations = v
	return s
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations struct {
	AvailableDuration []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration `json:"AvailableDuration,omitempty" xml:"AvailableDuration,omitempty" type:"Repeated"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations) SetAvailableDuration(v []*DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurations {
	s.AvailableDuration = v
	return s
}

type DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *int32  `json:"Value,omitempty" xml:"Value,omitempty"`
	Unit  *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) SetName(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration {
	s.Name = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) SetValue(v int32) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration {
	s.Value = &v
	return s
}

func (s *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration) SetUnit(v string) *DescribeResourcePackageProductResponseBodyDataResourcePackagesResourcePackagePackageTypesPackageTypeSpecificationsSpecificationAvailableDurationsAvailableDuration {
	s.Unit = &v
	return s
}

type DescribeResourcePackageProductResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeResourcePackageProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourcePackageProductResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackageProductResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackageProductResponse) SetHeaders(v map[string]*string) *DescribeResourcePackageProductResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcePackageProductResponse) SetBody(v *DescribeResourcePackageProductResponseBody) *DescribeResourcePackageProductResponse {
	s.Body = v
	return s
}

type DescribeSplitItemBillRequest struct {
	BillingCycle     *string                                  `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	ProductCode      *string                                  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string                                  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string                                  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	OwnerId          *int64                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	NextToken        *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults       *int32                                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	BillOwnerId      *int64                                   `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	TagFilter        []*DescribeSplitItemBillRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Repeated"`
	InstanceID       *string                                  `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	SplitItemID      *string                                  `json:"SplitItemID,omitempty" xml:"SplitItemID,omitempty"`
}

func (s DescribeSplitItemBillRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSplitItemBillRequest) GoString() string {
	return s.String()
}

func (s *DescribeSplitItemBillRequest) SetBillingCycle(v string) *DescribeSplitItemBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetProductCode(v string) *DescribeSplitItemBillRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetProductType(v string) *DescribeSplitItemBillRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetSubscriptionType(v string) *DescribeSplitItemBillRequest {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetOwnerId(v int64) *DescribeSplitItemBillRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetNextToken(v string) *DescribeSplitItemBillRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetMaxResults(v int32) *DescribeSplitItemBillRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetBillOwnerId(v int64) *DescribeSplitItemBillRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetTagFilter(v []*DescribeSplitItemBillRequestTagFilter) *DescribeSplitItemBillRequest {
	s.TagFilter = v
	return s
}

func (s *DescribeSplitItemBillRequest) SetInstanceID(v string) *DescribeSplitItemBillRequest {
	s.InstanceID = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetSplitItemID(v string) *DescribeSplitItemBillRequest {
	s.SplitItemID = &v
	return s
}

type DescribeSplitItemBillRequestTagFilter struct {
	TagKey    *string   `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s DescribeSplitItemBillRequestTagFilter) String() string {
	return tea.Prettify(s)
}

func (s DescribeSplitItemBillRequestTagFilter) GoString() string {
	return s.String()
}

func (s *DescribeSplitItemBillRequestTagFilter) SetTagKey(v string) *DescribeSplitItemBillRequestTagFilter {
	s.TagKey = &v
	return s
}

func (s *DescribeSplitItemBillRequestTagFilter) SetTagValues(v []*string) *DescribeSplitItemBillRequestTagFilter {
	s.TagValues = v
	return s
}

type DescribeSplitItemBillResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *DescribeSplitItemBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeSplitItemBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSplitItemBillResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSplitItemBillResponseBody) SetRequestId(v string) *DescribeSplitItemBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSplitItemBillResponseBody) SetSuccess(v bool) *DescribeSplitItemBillResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSplitItemBillResponseBody) SetCode(v string) *DescribeSplitItemBillResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSplitItemBillResponseBody) SetMessage(v string) *DescribeSplitItemBillResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSplitItemBillResponseBody) SetData(v *DescribeSplitItemBillResponseBodyData) *DescribeSplitItemBillResponseBody {
	s.Data = v
	return s
}

type DescribeSplitItemBillResponseBodyData struct {
	BillingCycle *string                                       `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	AccountID    *string                                       `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	AccountName  *string                                       `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken    *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults   *int32                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Items        []*DescribeSplitItemBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeSplitItemBillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeSplitItemBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSplitItemBillResponseBodyData) SetBillingCycle(v string) *DescribeSplitItemBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) SetAccountID(v string) *DescribeSplitItemBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) SetAccountName(v string) *DescribeSplitItemBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) SetTotalCount(v int32) *DescribeSplitItemBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) SetNextToken(v string) *DescribeSplitItemBillResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) SetMaxResults(v int32) *DescribeSplitItemBillResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyData) SetItems(v []*DescribeSplitItemBillResponseBodyDataItems) *DescribeSplitItemBillResponseBodyData {
	s.Items = v
	return s
}

type DescribeSplitItemBillResponseBodyDataItems struct {
	InstanceID                *string  `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	BillingType               *string  `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CostUnit                  *string  `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
	ProductCode               *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType               *string  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType          *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	ProductName               *string  `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductDetail             *string  `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	OwnerID                   *string  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	BillingItem               *string  `json:"BillingItem,omitempty" xml:"BillingItem,omitempty"`
	ListPrice                 *string  `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	ListPriceUnit             *string  `json:"ListPriceUnit,omitempty" xml:"ListPriceUnit,omitempty"`
	Usage                     *string  `json:"Usage,omitempty" xml:"Usage,omitempty"`
	UsageUnit                 *string  `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty"`
	DeductedByResourcePackage *string  `json:"DeductedByResourcePackage,omitempty" xml:"DeductedByResourcePackage,omitempty"`
	PretaxGrossAmount         *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	InvoiceDiscount           *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	DeductedByCoupons         *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	PretaxAmount              *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	DeductedByCashCoupons     *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByPrepaidCard     *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	PaymentAmount             *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	OutstandingAmount         *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	Currency                  *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	NickName                  *string  `json:"NickName,omitempty" xml:"NickName,omitempty"`
	ResourceGroup             *string  `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	Tag                       *string  `json:"Tag,omitempty" xml:"Tag,omitempty"`
	InstanceConfig            *string  `json:"InstanceConfig,omitempty" xml:"InstanceConfig,omitempty"`
	InstanceSpec              *string  `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	InternetIP                *string  `json:"InternetIP,omitempty" xml:"InternetIP,omitempty"`
	IntranetIP                *string  `json:"IntranetIP,omitempty" xml:"IntranetIP,omitempty"`
	Region                    *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	Zone                      *string  `json:"Zone,omitempty" xml:"Zone,omitempty"`
	Item                      *string  `json:"Item,omitempty" xml:"Item,omitempty"`
	ServicePeriod             *string  `json:"ServicePeriod,omitempty" xml:"ServicePeriod,omitempty"`
	BillingDate               *string  `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	SplitItemID               *string  `json:"SplitItemID,omitempty" xml:"SplitItemID,omitempty"`
	SplitItemName             *string  `json:"SplitItemName,omitempty" xml:"SplitItemName,omitempty"`
	PipCode                   *string  `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	CommodityCode             *string  `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ServicePeriodUnit         *string  `json:"ServicePeriodUnit,omitempty" xml:"ServicePeriodUnit,omitempty"`
	SplitCommodityCode        *string  `json:"SplitCommodityCode,omitempty" xml:"SplitCommodityCode,omitempty"`
	SplitProductDetail        *string  `json:"SplitProductDetail,omitempty" xml:"SplitProductDetail,omitempty"`
	SplitAccountID            *string  `json:"SplitAccountID,omitempty" xml:"SplitAccountID,omitempty"`
	SplitAccountName          *string  `json:"SplitAccountName,omitempty" xml:"SplitAccountName,omitempty"`
	SplitBillingCycle         *string  `json:"SplitBillingCycle,omitempty" xml:"SplitBillingCycle,omitempty"`
}

func (s DescribeSplitItemBillResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSplitItemBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetInstanceID(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.InstanceID = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetBillingType(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.BillingType = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetCostUnit(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.CostUnit = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetProductCode(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetProductType(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ProductType = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSubscriptionType(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetProductName(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ProductName = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetProductDetail(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ProductDetail = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetOwnerID(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.OwnerID = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetBillingItem(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.BillingItem = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetListPrice(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ListPrice = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetListPriceUnit(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ListPriceUnit = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetUsage(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.Usage = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetUsageUnit(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.UsageUnit = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetDeductedByResourcePackage(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.DeductedByResourcePackage = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetPretaxGrossAmount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.PretaxGrossAmount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetInvoiceDiscount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.InvoiceDiscount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetDeductedByCoupons(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.DeductedByCoupons = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetPretaxAmount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.PretaxAmount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetDeductedByCashCoupons(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetDeductedByPrepaidCard(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetPaymentAmount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.PaymentAmount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetOutstandingAmount(v float32) *DescribeSplitItemBillResponseBodyDataItems {
	s.OutstandingAmount = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetCurrency(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.Currency = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetNickName(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.NickName = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetResourceGroup(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetTag(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.Tag = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetInstanceConfig(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.InstanceConfig = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetInstanceSpec(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetInternetIP(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.InternetIP = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetIntranetIP(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.IntranetIP = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetRegion(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetZone(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.Zone = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetItem(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.Item = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetServicePeriod(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ServicePeriod = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetBillingDate(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.BillingDate = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitItemID(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitItemID = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitItemName(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitItemName = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetPipCode(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.PipCode = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetCommodityCode(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.CommodityCode = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetServicePeriodUnit(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.ServicePeriodUnit = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitCommodityCode(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitCommodityCode = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitProductDetail(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitProductDetail = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitAccountID(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitAccountID = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitAccountName(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitAccountName = &v
	return s
}

func (s *DescribeSplitItemBillResponseBodyDataItems) SetSplitBillingCycle(v string) *DescribeSplitItemBillResponseBodyDataItems {
	s.SplitBillingCycle = &v
	return s
}

type DescribeSplitItemBillResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSplitItemBillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSplitItemBillResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSplitItemBillResponse) GoString() string {
	return s.String()
}

func (s *DescribeSplitItemBillResponse) SetHeaders(v map[string]*string) *DescribeSplitItemBillResponse {
	s.Headers = v
	return s
}

func (s *DescribeSplitItemBillResponse) SetBody(v *DescribeSplitItemBillResponseBody) *DescribeSplitItemBillResponse {
	s.Body = v
	return s
}

type EnableBillGenerationRequest struct {
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s EnableBillGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableBillGenerationRequest) GoString() string {
	return s.String()
}

func (s *EnableBillGenerationRequest) SetProductCode(v string) *EnableBillGenerationRequest {
	s.ProductCode = &v
	return s
}

func (s *EnableBillGenerationRequest) SetOwnerId(v int64) *EnableBillGenerationRequest {
	s.OwnerId = &v
	return s
}

type EnableBillGenerationResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *EnableBillGenerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s EnableBillGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableBillGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *EnableBillGenerationResponseBody) SetRequestId(v string) *EnableBillGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableBillGenerationResponseBody) SetSuccess(v bool) *EnableBillGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *EnableBillGenerationResponseBody) SetCode(v string) *EnableBillGenerationResponseBody {
	s.Code = &v
	return s
}

func (s *EnableBillGenerationResponseBody) SetMessage(v string) *EnableBillGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *EnableBillGenerationResponseBody) SetData(v *EnableBillGenerationResponseBodyData) *EnableBillGenerationResponseBody {
	s.Data = v
	return s
}

type EnableBillGenerationResponseBodyData struct {
	Boolean *bool `json:"Boolean,omitempty" xml:"Boolean,omitempty"`
}

func (s EnableBillGenerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnableBillGenerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnableBillGenerationResponseBodyData) SetBoolean(v bool) *EnableBillGenerationResponseBodyData {
	s.Boolean = &v
	return s
}

type EnableBillGenerationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableBillGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableBillGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableBillGenerationResponse) GoString() string {
	return s.String()
}

func (s *EnableBillGenerationResponse) SetHeaders(v map[string]*string) *EnableBillGenerationResponse {
	s.Headers = v
	return s
}

func (s *EnableBillGenerationResponse) SetBody(v *EnableBillGenerationResponseBody) *EnableBillGenerationResponse {
	s.Body = v
	return s
}

type GetCustomerAccountInfoRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s GetCustomerAccountInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerAccountInfoRequest) SetOwnerId(v int64) *GetCustomerAccountInfoRequest {
	s.OwnerId = &v
	return s
}

type GetCustomerAccountInfoResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *GetCustomerAccountInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetCustomerAccountInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerAccountInfoResponseBody) SetRequestId(v string) *GetCustomerAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBody) SetSuccess(v bool) *GetCustomerAccountInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBody) SetCode(v string) *GetCustomerAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBody) SetMessage(v string) *GetCustomerAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBody) SetData(v *GetCustomerAccountInfoResponseBodyData) *GetCustomerAccountInfoResponseBody {
	s.Data = v
	return s
}

type GetCustomerAccountInfoResponseBodyData struct {
	LoginEmail        *string `json:"LoginEmail,omitempty" xml:"LoginEmail,omitempty"`
	AccountType       *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Mpk               *int64  `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	HostingStatus     *string `json:"HostingStatus,omitempty" xml:"HostingStatus,omitempty"`
	CreditLimitStatus *string `json:"CreditLimitStatus,omitempty" xml:"CreditLimitStatus,omitempty"`
	IsCertified       *bool   `json:"IsCertified,omitempty" xml:"IsCertified,omitempty"`
}

func (s GetCustomerAccountInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerAccountInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerAccountInfoResponseBodyData) SetLoginEmail(v string) *GetCustomerAccountInfoResponseBodyData {
	s.LoginEmail = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBodyData) SetAccountType(v string) *GetCustomerAccountInfoResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBodyData) SetMpk(v int64) *GetCustomerAccountInfoResponseBodyData {
	s.Mpk = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBodyData) SetHostingStatus(v string) *GetCustomerAccountInfoResponseBodyData {
	s.HostingStatus = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBodyData) SetCreditLimitStatus(v string) *GetCustomerAccountInfoResponseBodyData {
	s.CreditLimitStatus = &v
	return s
}

func (s *GetCustomerAccountInfoResponseBodyData) SetIsCertified(v bool) *GetCustomerAccountInfoResponseBodyData {
	s.IsCertified = &v
	return s
}

type GetCustomerAccountInfoResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCustomerAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCustomerAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerAccountInfoResponse) SetHeaders(v map[string]*string) *GetCustomerAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerAccountInfoResponse) SetBody(v *GetCustomerAccountInfoResponseBody) *GetCustomerAccountInfoResponse {
	s.Body = v
	return s
}

type GetCustomerListResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *GetCustomerListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetCustomerListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomerListResponseBody) SetRequestId(v string) *GetCustomerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomerListResponseBody) SetSuccess(v bool) *GetCustomerListResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomerListResponseBody) SetCode(v string) *GetCustomerListResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomerListResponseBody) SetMessage(v string) *GetCustomerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomerListResponseBody) SetData(v *GetCustomerListResponseBodyData) *GetCustomerListResponseBody {
	s.Data = v
	return s
}

type GetCustomerListResponseBodyData struct {
	UidList []*string `json:"UidList,omitempty" xml:"UidList,omitempty" type:"Repeated"`
}

func (s GetCustomerListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomerListResponseBodyData) SetUidList(v []*string) *GetCustomerListResponseBodyData {
	s.UidList = v
	return s
}

type GetCustomerListResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCustomerListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCustomerListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomerListResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerListResponse) SetHeaders(v map[string]*string) *GetCustomerListResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerListResponse) SetBody(v *GetCustomerListResponseBody) *GetCustomerListResponse {
	s.Body = v
	return s
}

type GetOrderDetailRequest struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s GetOrderDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrderDetailRequest) GoString() string {
	return s.String()
}

func (s *GetOrderDetailRequest) SetOrderId(v string) *GetOrderDetailRequest {
	s.OrderId = &v
	return s
}

func (s *GetOrderDetailRequest) SetOwnerId(v int64) *GetOrderDetailRequest {
	s.OwnerId = &v
	return s
}

type GetOrderDetailResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *GetOrderDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetOrderDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBody) SetRequestId(v string) *GetOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrderDetailResponseBody) SetMessage(v string) *GetOrderDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetOrderDetailResponseBody) SetCode(v string) *GetOrderDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrderDetailResponseBody) SetSuccess(v bool) *GetOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetOrderDetailResponseBody) SetData(v *GetOrderDetailResponseBodyData) *GetOrderDetailResponseBody {
	s.Data = v
	return s
}

type GetOrderDetailResponseBodyData struct {
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum    *int32                                   `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	HostName   *string                                  `json:"HostName,omitempty" xml:"HostName,omitempty"`
	OrderList  *GetOrderDetailResponseBodyDataOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Struct"`
}

func (s GetOrderDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOrderDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyData) SetTotalCount(v int32) *GetOrderDetailResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetOrderDetailResponseBodyData) SetPageSize(v int32) *GetOrderDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetOrderDetailResponseBodyData) SetPageNum(v int32) *GetOrderDetailResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetOrderDetailResponseBodyData) SetHostName(v string) *GetOrderDetailResponseBodyData {
	s.HostName = &v
	return s
}

func (s *GetOrderDetailResponseBodyData) SetOrderList(v *GetOrderDetailResponseBodyDataOrderList) *GetOrderDetailResponseBodyData {
	s.OrderList = v
	return s
}

type GetOrderDetailResponseBodyDataOrderList struct {
	Order []*GetOrderDetailResponseBodyDataOrderListOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Repeated"`
}

func (s GetOrderDetailResponseBodyDataOrderList) String() string {
	return tea.Prettify(s)
}

func (s GetOrderDetailResponseBodyDataOrderList) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyDataOrderList) SetOrder(v []*GetOrderDetailResponseBodyDataOrderListOrder) *GetOrderDetailResponseBodyDataOrderList {
	s.Order = v
	return s
}

type GetOrderDetailResponseBodyDataOrderListOrder struct {
	Operator          *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	AfterTaxAmount    *string `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	SubOrderId        *string `json:"SubOrderId,omitempty" xml:"SubOrderId,omitempty"`
	Config            *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Tax               *string `json:"Tax,omitempty" xml:"Tax,omitempty"`
	PaymentTime       *string `json:"PaymentTime,omitempty" xml:"PaymentTime,omitempty"`
	PaymentCurrency   *string `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	UsageEndTime      *string `json:"UsageEndTime,omitempty" xml:"UsageEndTime,omitempty"`
	SubscriptionType  *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	PretaxGrossAmount *string `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	OrderType         *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	Currency          *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	UsageStartTime    *string `json:"UsageStartTime,omitempty" xml:"UsageStartTime,omitempty"`
	OriginalConfig    *string `json:"OriginalConfig,omitempty" xml:"OriginalConfig,omitempty"`
	PaymentStatus     *string `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProductType       *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RelatedOrderId    *string `json:"RelatedOrderId,omitempty" xml:"RelatedOrderId,omitempty"`
	Quantity          *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	OrderId           *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PretaxAmount      *string `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	OrderSubType      *string `json:"OrderSubType,omitempty" xml:"OrderSubType,omitempty"`
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceIDs       *string `json:"InstanceIDs,omitempty" xml:"InstanceIDs,omitempty"`
	PretaxAmountLocal *string `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	CommodityCode     *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s GetOrderDetailResponseBodyDataOrderListOrder) String() string {
	return tea.Prettify(s)
}

func (s GetOrderDetailResponseBodyDataOrderListOrder) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetOperator(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.Operator = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetAfterTaxAmount(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.AfterTaxAmount = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetSubOrderId(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.SubOrderId = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetConfig(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.Config = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetTax(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.Tax = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetPaymentTime(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.PaymentTime = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetPaymentCurrency(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.PaymentCurrency = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetUsageEndTime(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.UsageEndTime = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetSubscriptionType(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.SubscriptionType = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetPretaxGrossAmount(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.PretaxGrossAmount = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetOrderType(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.OrderType = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetCurrency(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.Currency = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetUsageStartTime(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.UsageStartTime = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetOriginalConfig(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.OriginalConfig = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetPaymentStatus(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.PaymentStatus = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetProductCode(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.ProductCode = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetCreateTime(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.CreateTime = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetProductType(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.ProductType = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetRelatedOrderId(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.RelatedOrderId = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetQuantity(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.Quantity = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetOrderId(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.OrderId = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetPretaxAmount(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.PretaxAmount = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetOrderSubType(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.OrderSubType = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetRegion(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.Region = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetInstanceIDs(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.InstanceIDs = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetPretaxAmountLocal(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.PretaxAmountLocal = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetCommodityCode(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.CommodityCode = &v
	return s
}

type GetOrderDetailResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrderDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponse) SetHeaders(v map[string]*string) *GetOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetOrderDetailResponse) SetBody(v *GetOrderDetailResponseBody) *GetOrderDetailResponse {
	s.Body = v
	return s
}

type GetPayAsYouGoPriceRequest struct {
	OwnerId          *int64                                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductCode      *string                                `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string                                `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string                                `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	Region           *string                                `json:"Region,omitempty" xml:"Region,omitempty"`
	ModuleList       []*GetPayAsYouGoPriceRequestModuleList `json:"ModuleList,omitempty" xml:"ModuleList,omitempty" type:"Repeated"`
}

func (s GetPayAsYouGoPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceRequest) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceRequest) SetOwnerId(v int64) *GetPayAsYouGoPriceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPayAsYouGoPriceRequest) SetProductCode(v string) *GetPayAsYouGoPriceRequest {
	s.ProductCode = &v
	return s
}

func (s *GetPayAsYouGoPriceRequest) SetProductType(v string) *GetPayAsYouGoPriceRequest {
	s.ProductType = &v
	return s
}

func (s *GetPayAsYouGoPriceRequest) SetSubscriptionType(v string) *GetPayAsYouGoPriceRequest {
	s.SubscriptionType = &v
	return s
}

func (s *GetPayAsYouGoPriceRequest) SetRegion(v string) *GetPayAsYouGoPriceRequest {
	s.Region = &v
	return s
}

func (s *GetPayAsYouGoPriceRequest) SetModuleList(v []*GetPayAsYouGoPriceRequestModuleList) *GetPayAsYouGoPriceRequest {
	s.ModuleList = v
	return s
}

type GetPayAsYouGoPriceRequestModuleList struct {
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	Config     *string `json:"Config,omitempty" xml:"Config,omitempty"`
	PriceType  *string `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
}

func (s GetPayAsYouGoPriceRequestModuleList) String() string {
	return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceRequestModuleList) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceRequestModuleList) SetModuleCode(v string) *GetPayAsYouGoPriceRequestModuleList {
	s.ModuleCode = &v
	return s
}

func (s *GetPayAsYouGoPriceRequestModuleList) SetConfig(v string) *GetPayAsYouGoPriceRequestModuleList {
	s.Config = &v
	return s
}

func (s *GetPayAsYouGoPriceRequestModuleList) SetPriceType(v string) *GetPayAsYouGoPriceRequestModuleList {
	s.PriceType = &v
	return s
}

type GetPayAsYouGoPriceResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *GetPayAsYouGoPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetPayAsYouGoPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponseBody) SetRequestId(v string) *GetPayAsYouGoPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBody) SetSuccess(v bool) *GetPayAsYouGoPriceResponseBody {
	s.Success = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBody) SetCode(v string) *GetPayAsYouGoPriceResponseBody {
	s.Code = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBody) SetMessage(v string) *GetPayAsYouGoPriceResponseBody {
	s.Message = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBody) SetData(v *GetPayAsYouGoPriceResponseBodyData) *GetPayAsYouGoPriceResponseBody {
	s.Data = v
	return s
}

type GetPayAsYouGoPriceResponseBodyData struct {
	Currency         *string                                             `json:"Currency,omitempty" xml:"Currency,omitempty"`
	ModuleDetails    *GetPayAsYouGoPriceResponseBodyDataModuleDetails    `json:"ModuleDetails,omitempty" xml:"ModuleDetails,omitempty" type:"Struct"`
	PromotionDetails *GetPayAsYouGoPriceResponseBodyDataPromotionDetails `json:"PromotionDetails,omitempty" xml:"PromotionDetails,omitempty" type:"Struct"`
}

func (s GetPayAsYouGoPriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponseBodyData) SetCurrency(v string) *GetPayAsYouGoPriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyData) SetModuleDetails(v *GetPayAsYouGoPriceResponseBodyDataModuleDetails) *GetPayAsYouGoPriceResponseBodyData {
	s.ModuleDetails = v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyData) SetPromotionDetails(v *GetPayAsYouGoPriceResponseBodyDataPromotionDetails) *GetPayAsYouGoPriceResponseBodyData {
	s.PromotionDetails = v
	return s
}

type GetPayAsYouGoPriceResponseBodyDataModuleDetails struct {
	ModuleDetail []*GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail `json:"ModuleDetail,omitempty" xml:"ModuleDetail,omitempty" type:"Repeated"`
}

func (s GetPayAsYouGoPriceResponseBodyDataModuleDetails) String() string {
	return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseBodyDataModuleDetails) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetails) SetModuleDetail(v []*GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) *GetPayAsYouGoPriceResponseBodyDataModuleDetails {
	s.ModuleDetail = v
	return s
}

type GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail struct {
	ModuleCode        *string  `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	OriginalCost      *float32 `json:"OriginalCost,omitempty" xml:"OriginalCost,omitempty"`
	InvoiceDiscount   *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	CostAfterDiscount *float32 `json:"CostAfterDiscount,omitempty" xml:"CostAfterDiscount,omitempty"`
	UnitPrice         *float32 `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`
}

func (s GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) String() string {
	return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) SetModuleCode(v string) *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail {
	s.ModuleCode = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) SetOriginalCost(v float32) *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail {
	s.OriginalCost = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) SetInvoiceDiscount(v float32) *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail {
	s.InvoiceDiscount = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) SetCostAfterDiscount(v float32) *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail {
	s.CostAfterDiscount = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) SetUnitPrice(v float32) *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail {
	s.UnitPrice = &v
	return s
}

type GetPayAsYouGoPriceResponseBodyDataPromotionDetails struct {
	PromotionDetail []*GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail `json:"PromotionDetail,omitempty" xml:"PromotionDetail,omitempty" type:"Repeated"`
}

func (s GetPayAsYouGoPriceResponseBodyDataPromotionDetails) String() string {
	return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseBodyDataPromotionDetails) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetails) SetPromotionDetail(v []*GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) *GetPayAsYouGoPriceResponseBodyDataPromotionDetails {
	s.PromotionDetail = v
	return s
}

type GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail struct {
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionId   *int64  `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) String() string {
	return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) SetPromotionName(v string) *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail {
	s.PromotionName = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) SetPromotionDesc(v string) *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail {
	s.PromotionDesc = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) SetPromotionId(v int64) *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail {
	s.PromotionId = &v
	return s
}

type GetPayAsYouGoPriceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPayAsYouGoPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPayAsYouGoPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPayAsYouGoPriceResponse) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponse) SetHeaders(v map[string]*string) *GetPayAsYouGoPriceResponse {
	s.Headers = v
	return s
}

func (s *GetPayAsYouGoPriceResponse) SetBody(v *GetPayAsYouGoPriceResponseBody) *GetPayAsYouGoPriceResponse {
	s.Body = v
	return s
}

type GetResourcePackagePriceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductCode   *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PackageType   *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	EffectiveDate *string `json:"EffectiveDate,omitempty" xml:"EffectiveDate,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Duration      *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PricingCycle  *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetResourcePackagePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePackagePriceRequest) GoString() string {
	return s.String()
}

func (s *GetResourcePackagePriceRequest) SetOwnerId(v int64) *GetResourcePackagePriceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetProductCode(v string) *GetResourcePackagePriceRequest {
	s.ProductCode = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetPackageType(v string) *GetResourcePackagePriceRequest {
	s.PackageType = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetEffectiveDate(v string) *GetResourcePackagePriceRequest {
	s.EffectiveDate = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetSpecification(v string) *GetResourcePackagePriceRequest {
	s.Specification = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetDuration(v int32) *GetResourcePackagePriceRequest {
	s.Duration = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetPricingCycle(v string) *GetResourcePackagePriceRequest {
	s.PricingCycle = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetOrderType(v string) *GetResourcePackagePriceRequest {
	s.OrderType = &v
	return s
}

func (s *GetResourcePackagePriceRequest) SetInstanceId(v string) *GetResourcePackagePriceRequest {
	s.InstanceId = &v
	return s
}

type GetResourcePackagePriceResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *GetResourcePackagePriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetResourcePackagePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePackagePriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcePackagePriceResponseBody) SetRequestId(v string) *GetResourcePackagePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourcePackagePriceResponseBody) SetSuccess(v bool) *GetResourcePackagePriceResponseBody {
	s.Success = &v
	return s
}

func (s *GetResourcePackagePriceResponseBody) SetCode(v string) *GetResourcePackagePriceResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourcePackagePriceResponseBody) SetMessage(v string) *GetResourcePackagePriceResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourcePackagePriceResponseBody) SetData(v *GetResourcePackagePriceResponseBodyData) *GetResourcePackagePriceResponseBody {
	s.Data = v
	return s
}

type GetResourcePackagePriceResponseBodyData struct {
	Currency      *string                                            `json:"Currency,omitempty" xml:"Currency,omitempty"`
	OriginalPrice *float32                                           `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float32                                           `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	DiscountPrice *float32                                           `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	Promotions    *GetResourcePackagePriceResponseBodyDataPromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Struct"`
}

func (s GetResourcePackagePriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePackagePriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourcePackagePriceResponseBodyData) SetCurrency(v string) *GetResourcePackagePriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetResourcePackagePriceResponseBodyData) SetOriginalPrice(v float32) *GetResourcePackagePriceResponseBodyData {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourcePackagePriceResponseBodyData) SetTradePrice(v float32) *GetResourcePackagePriceResponseBodyData {
	s.TradePrice = &v
	return s
}

func (s *GetResourcePackagePriceResponseBodyData) SetDiscountPrice(v float32) *GetResourcePackagePriceResponseBodyData {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourcePackagePriceResponseBodyData) SetPromotions(v *GetResourcePackagePriceResponseBodyDataPromotions) *GetResourcePackagePriceResponseBodyData {
	s.Promotions = v
	return s
}

type GetResourcePackagePriceResponseBodyDataPromotions struct {
	Promotion []*GetResourcePackagePriceResponseBodyDataPromotionsPromotion `json:"Promotion,omitempty" xml:"Promotion,omitempty" type:"Repeated"`
}

func (s GetResourcePackagePriceResponseBodyDataPromotions) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePackagePriceResponseBodyDataPromotions) GoString() string {
	return s.String()
}

func (s *GetResourcePackagePriceResponseBodyDataPromotions) SetPromotion(v []*GetResourcePackagePriceResponseBodyDataPromotionsPromotion) *GetResourcePackagePriceResponseBodyDataPromotions {
	s.Promotion = v
	return s
}

type GetResourcePackagePriceResponseBodyDataPromotionsPromotion struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetResourcePackagePriceResponseBodyDataPromotionsPromotion) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePackagePriceResponseBodyDataPromotionsPromotion) GoString() string {
	return s.String()
}

func (s *GetResourcePackagePriceResponseBodyDataPromotionsPromotion) SetId(v int64) *GetResourcePackagePriceResponseBodyDataPromotionsPromotion {
	s.Id = &v
	return s
}

func (s *GetResourcePackagePriceResponseBodyDataPromotionsPromotion) SetName(v string) *GetResourcePackagePriceResponseBodyDataPromotionsPromotion {
	s.Name = &v
	return s
}

type GetResourcePackagePriceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetResourcePackagePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourcePackagePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePackagePriceResponse) GoString() string {
	return s.String()
}

func (s *GetResourcePackagePriceResponse) SetHeaders(v map[string]*string) *GetResourcePackagePriceResponse {
	s.Headers = v
	return s
}

func (s *GetResourcePackagePriceResponse) SetBody(v *GetResourcePackagePriceResponseBody) *GetResourcePackagePriceResponse {
	s.Body = v
	return s
}

type GetSubscriptionPriceRequest struct {
	ServicePeriodUnit     *string                                  `json:"ServicePeriodUnit,omitempty" xml:"ServicePeriodUnit,omitempty"`
	SubscriptionType      *string                                  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	OwnerId               *int64                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductCode           *string                                  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	OrderType             *string                                  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	ServicePeriodQuantity *int32                                   `json:"ServicePeriodQuantity,omitempty" xml:"ServicePeriodQuantity,omitempty"`
	ProductType           *string                                  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Region                *string                                  `json:"Region,omitempty" xml:"Region,omitempty"`
	InstanceId            *string                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModuleList            []*GetSubscriptionPriceRequestModuleList `json:"ModuleList,omitempty" xml:"ModuleList,omitempty" type:"Repeated"`
	Quantity              *int32                                   `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s GetSubscriptionPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionPriceRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceRequest) SetServicePeriodUnit(v string) *GetSubscriptionPriceRequest {
	s.ServicePeriodUnit = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetSubscriptionType(v string) *GetSubscriptionPriceRequest {
	s.SubscriptionType = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetOwnerId(v int64) *GetSubscriptionPriceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetProductCode(v string) *GetSubscriptionPriceRequest {
	s.ProductCode = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetOrderType(v string) *GetSubscriptionPriceRequest {
	s.OrderType = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetServicePeriodQuantity(v int32) *GetSubscriptionPriceRequest {
	s.ServicePeriodQuantity = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetProductType(v string) *GetSubscriptionPriceRequest {
	s.ProductType = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetRegion(v string) *GetSubscriptionPriceRequest {
	s.Region = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetInstanceId(v string) *GetSubscriptionPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSubscriptionPriceRequest) SetModuleList(v []*GetSubscriptionPriceRequestModuleList) *GetSubscriptionPriceRequest {
	s.ModuleList = v
	return s
}

func (s *GetSubscriptionPriceRequest) SetQuantity(v int32) *GetSubscriptionPriceRequest {
	s.Quantity = &v
	return s
}

type GetSubscriptionPriceRequestModuleList struct {
	ModuleCode   *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	Config       *string `json:"Config,omitempty" xml:"Config,omitempty"`
	ModuleStatus *int32  `json:"ModuleStatus,omitempty" xml:"ModuleStatus,omitempty"`
	Tag          *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetSubscriptionPriceRequestModuleList) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionPriceRequestModuleList) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceRequestModuleList) SetModuleCode(v string) *GetSubscriptionPriceRequestModuleList {
	s.ModuleCode = &v
	return s
}

func (s *GetSubscriptionPriceRequestModuleList) SetConfig(v string) *GetSubscriptionPriceRequestModuleList {
	s.Config = &v
	return s
}

func (s *GetSubscriptionPriceRequestModuleList) SetModuleStatus(v int32) *GetSubscriptionPriceRequestModuleList {
	s.ModuleStatus = &v
	return s
}

func (s *GetSubscriptionPriceRequestModuleList) SetTag(v string) *GetSubscriptionPriceRequestModuleList {
	s.Tag = &v
	return s
}

type GetSubscriptionPriceResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *GetSubscriptionPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetSubscriptionPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBody) SetRequestId(v string) *GetSubscriptionPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscriptionPriceResponseBody) SetSuccess(v bool) *GetSubscriptionPriceResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubscriptionPriceResponseBody) SetCode(v string) *GetSubscriptionPriceResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubscriptionPriceResponseBody) SetMessage(v string) *GetSubscriptionPriceResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubscriptionPriceResponseBody) SetData(v *GetSubscriptionPriceResponseBodyData) *GetSubscriptionPriceResponseBody {
	s.Data = v
	return s
}

type GetSubscriptionPriceResponseBodyData struct {
	OriginalPrice    *float32                                              `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	DiscountPrice    *float32                                              `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	TradePrice       *float32                                              `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	Currency         *string                                               `json:"Currency,omitempty" xml:"Currency,omitempty"`
	Quantity         *int32                                                `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ModuleDetails    *GetSubscriptionPriceResponseBodyDataModuleDetails    `json:"ModuleDetails,omitempty" xml:"ModuleDetails,omitempty" type:"Struct"`
	PromotionDetails *GetSubscriptionPriceResponseBodyDataPromotionDetails `json:"PromotionDetails,omitempty" xml:"PromotionDetails,omitempty" type:"Struct"`
}

func (s GetSubscriptionPriceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBodyData) SetOriginalPrice(v float32) *GetSubscriptionPriceResponseBodyData {
	s.OriginalPrice = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetDiscountPrice(v float32) *GetSubscriptionPriceResponseBodyData {
	s.DiscountPrice = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetTradePrice(v float32) *GetSubscriptionPriceResponseBodyData {
	s.TradePrice = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetCurrency(v string) *GetSubscriptionPriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetQuantity(v int32) *GetSubscriptionPriceResponseBodyData {
	s.Quantity = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetModuleDetails(v *GetSubscriptionPriceResponseBodyDataModuleDetails) *GetSubscriptionPriceResponseBodyData {
	s.ModuleDetails = v
	return s
}

func (s *GetSubscriptionPriceResponseBodyData) SetPromotionDetails(v *GetSubscriptionPriceResponseBodyDataPromotionDetails) *GetSubscriptionPriceResponseBodyData {
	s.PromotionDetails = v
	return s
}

type GetSubscriptionPriceResponseBodyDataModuleDetails struct {
	ModuleDetail []*GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail `json:"ModuleDetail,omitempty" xml:"ModuleDetail,omitempty" type:"Repeated"`
}

func (s GetSubscriptionPriceResponseBodyDataModuleDetails) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponseBodyDataModuleDetails) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetails) SetModuleDetail(v []*GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) *GetSubscriptionPriceResponseBodyDataModuleDetails {
	s.ModuleDetail = v
	return s
}

type GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail struct {
	ModuleCode        *string  `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	OriginalCost      *float32 `json:"OriginalCost,omitempty" xml:"OriginalCost,omitempty"`
	InvoiceDiscount   *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	CostAfterDiscount *float32 `json:"CostAfterDiscount,omitempty" xml:"CostAfterDiscount,omitempty"`
	UnitPrice         *float32 `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`
}

func (s GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) SetModuleCode(v string) *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail {
	s.ModuleCode = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) SetOriginalCost(v float32) *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail {
	s.OriginalCost = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) SetInvoiceDiscount(v float32) *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail {
	s.InvoiceDiscount = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) SetCostAfterDiscount(v float32) *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail {
	s.CostAfterDiscount = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail) SetUnitPrice(v float32) *GetSubscriptionPriceResponseBodyDataModuleDetailsModuleDetail {
	s.UnitPrice = &v
	return s
}

type GetSubscriptionPriceResponseBodyDataPromotionDetails struct {
	PromotionDetail []*GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail `json:"PromotionDetail,omitempty" xml:"PromotionDetail,omitempty" type:"Repeated"`
}

func (s GetSubscriptionPriceResponseBodyDataPromotionDetails) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponseBodyDataPromotionDetails) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetails) SetPromotionDetail(v []*GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) *GetSubscriptionPriceResponseBodyDataPromotionDetails {
	s.PromotionDetail = v
	return s
}

type GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail struct {
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionId   *int64  `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) SetPromotionName(v string) *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail {
	s.PromotionName = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) SetPromotionDesc(v string) *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail {
	s.PromotionDesc = &v
	return s
}

func (s *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail) SetPromotionId(v int64) *GetSubscriptionPriceResponseBodyDataPromotionDetailsPromotionDetail {
	s.PromotionId = &v
	return s
}

type GetSubscriptionPriceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSubscriptionPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSubscriptionPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionPriceResponse) GoString() string {
	return s.String()
}

func (s *GetSubscriptionPriceResponse) SetHeaders(v map[string]*string) *GetSubscriptionPriceResponse {
	s.Headers = v
	return s
}

func (s *GetSubscriptionPriceResponse) SetBody(v *GetSubscriptionPriceResponseBody) *GetSubscriptionPriceResponse {
	s.Body = v
	return s
}

type ModifyCostUnitRequest struct {
	UnitEntityList []*ModifyCostUnitRequestUnitEntityList `json:"UnitEntityList,omitempty" xml:"UnitEntityList,omitempty" type:"Repeated"`
}

func (s ModifyCostUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyCostUnitRequest) GoString() string {
	return s.String()
}

func (s *ModifyCostUnitRequest) SetUnitEntityList(v []*ModifyCostUnitRequestUnitEntityList) *ModifyCostUnitRequest {
	s.UnitEntityList = v
	return s
}

type ModifyCostUnitRequestUnitEntityList struct {
	NewUnitName *string `json:"NewUnitName,omitempty" xml:"NewUnitName,omitempty"`
	OwnerUid    *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	UnitId      *int64  `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
}

func (s ModifyCostUnitRequestUnitEntityList) String() string {
	return tea.Prettify(s)
}

func (s ModifyCostUnitRequestUnitEntityList) GoString() string {
	return s.String()
}

func (s *ModifyCostUnitRequestUnitEntityList) SetNewUnitName(v string) *ModifyCostUnitRequestUnitEntityList {
	s.NewUnitName = &v
	return s
}

func (s *ModifyCostUnitRequestUnitEntityList) SetOwnerUid(v int64) *ModifyCostUnitRequestUnitEntityList {
	s.OwnerUid = &v
	return s
}

func (s *ModifyCostUnitRequestUnitEntityList) SetUnitId(v int64) *ModifyCostUnitRequestUnitEntityList {
	s.UnitId = &v
	return s
}

type ModifyCostUnitResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      []*ModifyCostUnitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ModifyCostUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyCostUnitResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCostUnitResponseBody) SetRequestId(v string) *ModifyCostUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCostUnitResponseBody) SetMessage(v string) *ModifyCostUnitResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyCostUnitResponseBody) SetCode(v string) *ModifyCostUnitResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyCostUnitResponseBody) SetSuccess(v bool) *ModifyCostUnitResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyCostUnitResponseBody) SetData(v []*ModifyCostUnitResponseBodyData) *ModifyCostUnitResponseBody {
	s.Data = v
	return s
}

type ModifyCostUnitResponseBodyData struct {
	IsSuccess *bool  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	UnitId    *int64 `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
	OwnerUid  *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
}

func (s ModifyCostUnitResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyCostUnitResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyCostUnitResponseBodyData) SetIsSuccess(v bool) *ModifyCostUnitResponseBodyData {
	s.IsSuccess = &v
	return s
}

func (s *ModifyCostUnitResponseBodyData) SetUnitId(v int64) *ModifyCostUnitResponseBodyData {
	s.UnitId = &v
	return s
}

func (s *ModifyCostUnitResponseBodyData) SetOwnerUid(v int64) *ModifyCostUnitResponseBodyData {
	s.OwnerUid = &v
	return s
}

type ModifyCostUnitResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyCostUnitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyCostUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyCostUnitResponse) GoString() string {
	return s.String()
}

func (s *ModifyCostUnitResponse) SetHeaders(v map[string]*string) *ModifyCostUnitResponse {
	s.Headers = v
	return s
}

func (s *ModifyCostUnitResponse) SetBody(v *ModifyCostUnitResponseBody) *ModifyCostUnitResponse {
	s.Body = v
	return s
}

type ModifyInstanceRequest struct {
	ProductCode      *string                           `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	OwnerId          *int64                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductType      *string                           `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string                           `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	ModifyType       *string                           `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	InstanceId       *string                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Parameter        []*ModifyInstanceRequestParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
	ClientToken      *string                           `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s ModifyInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRequest) SetProductCode(v string) *ModifyInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *ModifyInstanceRequest) SetOwnerId(v int64) *ModifyInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceRequest) SetProductType(v string) *ModifyInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyInstanceRequest) SetSubscriptionType(v string) *ModifyInstanceRequest {
	s.SubscriptionType = &v
	return s
}

func (s *ModifyInstanceRequest) SetModifyType(v string) *ModifyInstanceRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyInstanceRequest) SetInstanceId(v string) *ModifyInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRequest) SetParameter(v []*ModifyInstanceRequestParameter) *ModifyInstanceRequest {
	s.Parameter = v
	return s
}

func (s *ModifyInstanceRequest) SetClientToken(v string) *ModifyInstanceRequest {
	s.ClientToken = &v
	return s
}

type ModifyInstanceRequestParameter struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyInstanceRequestParameter) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRequestParameter) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRequestParameter) SetCode(v string) *ModifyInstanceRequestParameter {
	s.Code = &v
	return s
}

func (s *ModifyInstanceRequestParameter) SetValue(v string) *ModifyInstanceRequestParameter {
	s.Value = &v
	return s
}

type ModifyInstanceResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *ModifyInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ModifyInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBody) SetRequestId(v string) *ModifyInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetSuccess(v bool) *ModifyInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetCode(v string) *ModifyInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetMessage(v string) *ModifyInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetData(v *ModifyInstanceResponseBodyData) *ModifyInstanceResponseBody {
	s.Data = v
	return s
}

type ModifyInstanceResponseBodyData struct {
	HostId  *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBodyData) SetHostId(v string) *ModifyInstanceResponseBodyData {
	s.HostId = &v
	return s
}

func (s *ModifyInstanceResponseBodyData) SetOrderId(v string) *ModifyInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

type ModifyInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponse) SetHeaders(v map[string]*string) *ModifyInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceResponse) SetBody(v *ModifyInstanceResponseBody) *ModifyInstanceResponse {
	s.Body = v
	return s
}

type QueryAccountBalanceResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryAccountBalanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryAccountBalanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountBalanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountBalanceResponseBody) SetRequestId(v string) *QueryAccountBalanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountBalanceResponseBody) SetSuccess(v bool) *QueryAccountBalanceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountBalanceResponseBody) SetCode(v string) *QueryAccountBalanceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountBalanceResponseBody) SetMessage(v string) *QueryAccountBalanceResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccountBalanceResponseBody) SetData(v *QueryAccountBalanceResponseBodyData) *QueryAccountBalanceResponseBody {
	s.Data = v
	return s
}

type QueryAccountBalanceResponseBodyData struct {
	AvailableAmount     *string `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	AvailableCashAmount *string `json:"AvailableCashAmount,omitempty" xml:"AvailableCashAmount,omitempty"`
	CreditAmount        *string `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	MybankCreditAmount  *string `json:"MybankCreditAmount,omitempty" xml:"MybankCreditAmount,omitempty"`
	Currency            *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
}

func (s QueryAccountBalanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountBalanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccountBalanceResponseBodyData) SetAvailableAmount(v string) *QueryAccountBalanceResponseBodyData {
	s.AvailableAmount = &v
	return s
}

func (s *QueryAccountBalanceResponseBodyData) SetAvailableCashAmount(v string) *QueryAccountBalanceResponseBodyData {
	s.AvailableCashAmount = &v
	return s
}

func (s *QueryAccountBalanceResponseBodyData) SetCreditAmount(v string) *QueryAccountBalanceResponseBodyData {
	s.CreditAmount = &v
	return s
}

func (s *QueryAccountBalanceResponseBodyData) SetMybankCreditAmount(v string) *QueryAccountBalanceResponseBodyData {
	s.MybankCreditAmount = &v
	return s
}

func (s *QueryAccountBalanceResponseBodyData) SetCurrency(v string) *QueryAccountBalanceResponseBodyData {
	s.Currency = &v
	return s
}

type QueryAccountBalanceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAccountBalanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAccountBalanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountBalanceResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountBalanceResponse) SetHeaders(v map[string]*string) *QueryAccountBalanceResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountBalanceResponse) SetBody(v *QueryAccountBalanceResponseBody) *QueryAccountBalanceResponse {
	s.Body = v
	return s
}

type QueryAccountBillRequest struct {
	BillingCycle     *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OwnerID          *int64  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	IsGroupByProduct *bool   `json:"IsGroupByProduct,omitempty" xml:"IsGroupByProduct,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	BillOwnerId      *int64  `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	Granularity      *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	BillingDate      *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
}

func (s QueryAccountBillRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountBillRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountBillRequest) SetBillingCycle(v string) *QueryAccountBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *QueryAccountBillRequest) SetPageNum(v int32) *QueryAccountBillRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAccountBillRequest) SetPageSize(v int32) *QueryAccountBillRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAccountBillRequest) SetOwnerID(v int64) *QueryAccountBillRequest {
	s.OwnerID = &v
	return s
}

func (s *QueryAccountBillRequest) SetIsGroupByProduct(v bool) *QueryAccountBillRequest {
	s.IsGroupByProduct = &v
	return s
}

func (s *QueryAccountBillRequest) SetProductCode(v string) *QueryAccountBillRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryAccountBillRequest) SetBillOwnerId(v int64) *QueryAccountBillRequest {
	s.BillOwnerId = &v
	return s
}

func (s *QueryAccountBillRequest) SetGranularity(v string) *QueryAccountBillRequest {
	s.Granularity = &v
	return s
}

func (s *QueryAccountBillRequest) SetBillingDate(v string) *QueryAccountBillRequest {
	s.BillingDate = &v
	return s
}

type QueryAccountBillResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryAccountBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryAccountBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountBillResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountBillResponseBody) SetRequestId(v string) *QueryAccountBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountBillResponseBody) SetSuccess(v bool) *QueryAccountBillResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountBillResponseBody) SetCode(v string) *QueryAccountBillResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountBillResponseBody) SetMessage(v string) *QueryAccountBillResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccountBillResponseBody) SetData(v *QueryAccountBillResponseBodyData) *QueryAccountBillResponseBody {
	s.Data = v
	return s
}

type QueryAccountBillResponseBodyData struct {
	BillingCycle *string                                `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	AccountID    *string                                `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	AccountName  *string                                `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	TotalCount   *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageNum      *int32                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Items        *QueryAccountBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s QueryAccountBillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccountBillResponseBodyData) SetBillingCycle(v string) *QueryAccountBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QueryAccountBillResponseBodyData) SetAccountID(v string) *QueryAccountBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QueryAccountBillResponseBodyData) SetAccountName(v string) *QueryAccountBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryAccountBillResponseBodyData) SetTotalCount(v int32) *QueryAccountBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryAccountBillResponseBodyData) SetPageNum(v int32) *QueryAccountBillResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryAccountBillResponseBodyData) SetPageSize(v int32) *QueryAccountBillResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAccountBillResponseBodyData) SetItems(v *QueryAccountBillResponseBodyDataItems) *QueryAccountBillResponseBodyData {
	s.Items = v
	return s
}

type QueryAccountBillResponseBodyDataItems struct {
	Item []*QueryAccountBillResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryAccountBillResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryAccountBillResponseBodyDataItems) SetItem(v []*QueryAccountBillResponseBodyDataItemsItem) *QueryAccountBillResponseBodyDataItems {
	s.Item = v
	return s
}

type QueryAccountBillResponseBodyDataItemsItem struct {
	CostUnit              *string  `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
	OwnerID               *string  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	PretaxGrossAmount     *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	InvoiceDiscount       *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	DeductedByCoupons     *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	PretaxAmount          *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	PaymentAmount         *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	OutstandingAmount     *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	Currency              *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	OwnerName             *string  `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	ProductCode           *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName           *string  `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	SubscriptionType      *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	PipCode               *string  `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	BillingDate           *string  `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
}

func (s QueryAccountBillResponseBodyDataItemsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountBillResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetCostUnit(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.CostUnit = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetOwnerID(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QueryAccountBillResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetCurrency(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetOwnerName(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.OwnerName = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetProductCode(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetProductName(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.ProductName = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetSubscriptionType(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetPipCode(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.PipCode = &v
	return s
}

func (s *QueryAccountBillResponseBodyDataItemsItem) SetBillingDate(v string) *QueryAccountBillResponseBodyDataItemsItem {
	s.BillingDate = &v
	return s
}

type QueryAccountBillResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAccountBillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAccountBillResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountBillResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountBillResponse) SetHeaders(v map[string]*string) *QueryAccountBillResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountBillResponse) SetBody(v *QueryAccountBillResponseBody) *QueryAccountBillResponse {
	s.Body = v
	return s
}

type QueryAccountTransactionDetailsRequest struct {
	TransactionNumber    *string `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	RecordID             *string `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	TransactionChannelSN *string `json:"TransactionChannelSN,omitempty" xml:"TransactionChannelSN,omitempty"`
	CreateTimeStart      *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	CreateTimeEnd        *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	TransactionType      *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
	TransactionChannel   *string `json:"TransactionChannel,omitempty" xml:"TransactionChannel,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults           *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s QueryAccountTransactionDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTransactionDetailsRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionDetailsRequest) SetTransactionNumber(v string) *QueryAccountTransactionDetailsRequest {
	s.TransactionNumber = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetRecordID(v string) *QueryAccountTransactionDetailsRequest {
	s.RecordID = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetTransactionChannelSN(v string) *QueryAccountTransactionDetailsRequest {
	s.TransactionChannelSN = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetCreateTimeStart(v string) *QueryAccountTransactionDetailsRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetCreateTimeEnd(v string) *QueryAccountTransactionDetailsRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetTransactionType(v string) *QueryAccountTransactionDetailsRequest {
	s.TransactionType = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetTransactionChannel(v string) *QueryAccountTransactionDetailsRequest {
	s.TransactionChannel = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetNextToken(v string) *QueryAccountTransactionDetailsRequest {
	s.NextToken = &v
	return s
}

func (s *QueryAccountTransactionDetailsRequest) SetMaxResults(v int32) *QueryAccountTransactionDetailsRequest {
	s.MaxResults = &v
	return s
}

type QueryAccountTransactionDetailsResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryAccountTransactionDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryAccountTransactionDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTransactionDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionDetailsResponseBody) SetRequestId(v string) *QueryAccountTransactionDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBody) SetSuccess(v bool) *QueryAccountTransactionDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBody) SetCode(v string) *QueryAccountTransactionDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBody) SetMessage(v string) *QueryAccountTransactionDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBody) SetData(v *QueryAccountTransactionDetailsResponseBodyData) *QueryAccountTransactionDetailsResponseBody {
	s.Data = v
	return s
}

type QueryAccountTransactionDetailsResponseBodyData struct {
	AccountName             *string                                                                `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	TotalCount              *int32                                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken               *string                                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults              *int32                                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	AccountTransactionsList *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList `json:"AccountTransactionsList,omitempty" xml:"AccountTransactionsList,omitempty" type:"Struct"`
}

func (s QueryAccountTransactionDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTransactionDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionDetailsResponseBodyData) SetAccountName(v string) *QueryAccountTransactionDetailsResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyData) SetTotalCount(v int32) *QueryAccountTransactionDetailsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyData) SetNextToken(v string) *QueryAccountTransactionDetailsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyData) SetMaxResults(v int32) *QueryAccountTransactionDetailsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyData) SetAccountTransactionsList(v *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList) *QueryAccountTransactionDetailsResponseBodyData {
	s.AccountTransactionsList = v
	return s
}

type QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList struct {
	AccountTransactionsList []*QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList `json:"AccountTransactionsList,omitempty" xml:"AccountTransactionsList,omitempty" type:"Repeated"`
}

func (s QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList) SetAccountTransactionsList(v []*QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsList {
	s.AccountTransactionsList = v
	return s
}

type QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList struct {
	TransactionNumber    *string `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	TransactionTime      *string `json:"TransactionTime,omitempty" xml:"TransactionTime,omitempty"`
	TransactionFlow      *string `json:"TransactionFlow,omitempty" xml:"TransactionFlow,omitempty"`
	TransactionType      *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
	TransactionChannel   *string `json:"TransactionChannel,omitempty" xml:"TransactionChannel,omitempty"`
	TransactionChannelSN *string `json:"TransactionChannelSN,omitempty" xml:"TransactionChannelSN,omitempty"`
	FundType             *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	RecordID             *string `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	Remarks              *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	BillingCycle         *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Amount               *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Balance              *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	TransactionAccount   *string `json:"TransactionAccount,omitempty" xml:"TransactionAccount,omitempty"`
}

func (s QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionNumber(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionNumber = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionTime(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionTime = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionFlow(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionFlow = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionType(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionType = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionChannel(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionChannel = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionChannelSN(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionChannelSN = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetFundType(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.FundType = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetRecordID(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.RecordID = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetRemarks(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.Remarks = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetBillingCycle(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.BillingCycle = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetAmount(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.Amount = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetBalance(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.Balance = &v
	return s
}

func (s *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionAccount(v string) *QueryAccountTransactionDetailsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionAccount = &v
	return s
}

type QueryAccountTransactionDetailsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAccountTransactionDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAccountTransactionDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTransactionDetailsResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionDetailsResponse) SetHeaders(v map[string]*string) *QueryAccountTransactionDetailsResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountTransactionDetailsResponse) SetBody(v *QueryAccountTransactionDetailsResponseBody) *QueryAccountTransactionDetailsResponse {
	s.Body = v
	return s
}

type QueryAccountTransactionsRequest struct {
	TransactionNumber    *string `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	RecordID             *string `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	TransactionChannelSN *string `json:"TransactionChannelSN,omitempty" xml:"TransactionChannelSN,omitempty"`
	CreateTimeStart      *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	CreateTimeEnd        *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	PageNum              *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryAccountTransactionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTransactionsRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionsRequest) SetTransactionNumber(v string) *QueryAccountTransactionsRequest {
	s.TransactionNumber = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetRecordID(v string) *QueryAccountTransactionsRequest {
	s.RecordID = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetTransactionChannelSN(v string) *QueryAccountTransactionsRequest {
	s.TransactionChannelSN = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetCreateTimeStart(v string) *QueryAccountTransactionsRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetCreateTimeEnd(v string) *QueryAccountTransactionsRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetPageNum(v int32) *QueryAccountTransactionsRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAccountTransactionsRequest) SetPageSize(v int32) *QueryAccountTransactionsRequest {
	s.PageSize = &v
	return s
}

type QueryAccountTransactionsResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryAccountTransactionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryAccountTransactionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTransactionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionsResponseBody) SetRequestId(v string) *QueryAccountTransactionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAccountTransactionsResponseBody) SetSuccess(v bool) *QueryAccountTransactionsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAccountTransactionsResponseBody) SetCode(v string) *QueryAccountTransactionsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAccountTransactionsResponseBody) SetMessage(v string) *QueryAccountTransactionsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAccountTransactionsResponseBody) SetData(v *QueryAccountTransactionsResponseBodyData) *QueryAccountTransactionsResponseBody {
	s.Data = v
	return s
}

type QueryAccountTransactionsResponseBodyData struct {
	AccountName             *string                                                          `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	TotalCount              *int32                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageNum                 *int32                                                           `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize                *int32                                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	AccountTransactionsList *QueryAccountTransactionsResponseBodyDataAccountTransactionsList `json:"AccountTransactionsList,omitempty" xml:"AccountTransactionsList,omitempty" type:"Struct"`
}

func (s QueryAccountTransactionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTransactionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionsResponseBodyData) SetAccountName(v string) *QueryAccountTransactionsResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyData) SetTotalCount(v int32) *QueryAccountTransactionsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyData) SetPageNum(v int32) *QueryAccountTransactionsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyData) SetPageSize(v int32) *QueryAccountTransactionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyData) SetAccountTransactionsList(v *QueryAccountTransactionsResponseBodyDataAccountTransactionsList) *QueryAccountTransactionsResponseBodyData {
	s.AccountTransactionsList = v
	return s
}

type QueryAccountTransactionsResponseBodyDataAccountTransactionsList struct {
	AccountTransactionsList []*QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList `json:"AccountTransactionsList,omitempty" xml:"AccountTransactionsList,omitempty" type:"Repeated"`
}

func (s QueryAccountTransactionsResponseBodyDataAccountTransactionsList) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTransactionsResponseBodyDataAccountTransactionsList) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsList) SetAccountTransactionsList(v []*QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) *QueryAccountTransactionsResponseBodyDataAccountTransactionsList {
	s.AccountTransactionsList = v
	return s
}

type QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList struct {
	TransactionNumber    *string `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	TransactionTime      *string `json:"TransactionTime,omitempty" xml:"TransactionTime,omitempty"`
	TransactionFlow      *string `json:"TransactionFlow,omitempty" xml:"TransactionFlow,omitempty"`
	TransactionType      *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
	TransactionChannel   *string `json:"TransactionChannel,omitempty" xml:"TransactionChannel,omitempty"`
	TransactionChannelSN *string `json:"TransactionChannelSN,omitempty" xml:"TransactionChannelSN,omitempty"`
	FundType             *string `json:"FundType,omitempty" xml:"FundType,omitempty"`
	RecordID             *string `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	Remarks              *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	BillingCycle         *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Amount               *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Balance              *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	TransactionAccount   *string `json:"TransactionAccount,omitempty" xml:"TransactionAccount,omitempty"`
}

func (s QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionNumber(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionNumber = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionTime(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionTime = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionFlow(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionFlow = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionType(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionType = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionChannel(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionChannel = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionChannelSN(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionChannelSN = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetFundType(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.FundType = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetRecordID(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.RecordID = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetRemarks(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.Remarks = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetBillingCycle(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.BillingCycle = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetAmount(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.Amount = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetBalance(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.Balance = &v
	return s
}

func (s *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList) SetTransactionAccount(v string) *QueryAccountTransactionsResponseBodyDataAccountTransactionsListAccountTransactionsList {
	s.TransactionAccount = &v
	return s
}

type QueryAccountTransactionsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAccountTransactionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAccountTransactionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAccountTransactionsResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountTransactionsResponse) SetHeaders(v map[string]*string) *QueryAccountTransactionsResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountTransactionsResponse) SetBody(v *QueryAccountTransactionsResponseBody) *QueryAccountTransactionsResponse {
	s.Body = v
	return s
}

type QueryAvailableInstancesRequest struct {
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	InstanceIDs      *string `json:"InstanceIDs,omitempty" xml:"InstanceIDs,omitempty"`
	EndTimeStart     *string `json:"EndTimeStart,omitempty" xml:"EndTimeStart,omitempty"`
	EndTimeEnd       *string `json:"EndTimeEnd,omitempty" xml:"EndTimeEnd,omitempty"`
	CreateTimeStart  *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	CreateTimeEnd    *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	RenewStatus      *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
}

func (s QueryAvailableInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAvailableInstancesRequest) GoString() string {
	return s.String()
}

func (s *QueryAvailableInstancesRequest) SetRegion(v string) *QueryAvailableInstancesRequest {
	s.Region = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetOwnerId(v int64) *QueryAvailableInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetPageNum(v int32) *QueryAvailableInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetPageSize(v int32) *QueryAvailableInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetProductCode(v string) *QueryAvailableInstancesRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetProductType(v string) *QueryAvailableInstancesRequest {
	s.ProductType = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetSubscriptionType(v string) *QueryAvailableInstancesRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetInstanceIDs(v string) *QueryAvailableInstancesRequest {
	s.InstanceIDs = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetEndTimeStart(v string) *QueryAvailableInstancesRequest {
	s.EndTimeStart = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetEndTimeEnd(v string) *QueryAvailableInstancesRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetCreateTimeStart(v string) *QueryAvailableInstancesRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetCreateTimeEnd(v string) *QueryAvailableInstancesRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryAvailableInstancesRequest) SetRenewStatus(v string) *QueryAvailableInstancesRequest {
	s.RenewStatus = &v
	return s
}

type QueryAvailableInstancesResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryAvailableInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryAvailableInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAvailableInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvailableInstancesResponseBody) SetRequestId(v string) *QueryAvailableInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvailableInstancesResponseBody) SetSuccess(v bool) *QueryAvailableInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAvailableInstancesResponseBody) SetCode(v string) *QueryAvailableInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAvailableInstancesResponseBody) SetMessage(v string) *QueryAvailableInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAvailableInstancesResponseBody) SetData(v *QueryAvailableInstancesResponseBodyData) *QueryAvailableInstancesResponseBody {
	s.Data = v
	return s
}

type QueryAvailableInstancesResponseBodyData struct {
	PageNum      *int32                                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount   *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	InstanceList []*QueryAvailableInstancesResponseBodyDataInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
}

func (s QueryAvailableInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAvailableInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAvailableInstancesResponseBodyData) SetPageNum(v int32) *QueryAvailableInstancesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyData) SetPageSize(v int32) *QueryAvailableInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyData) SetTotalCount(v int32) *QueryAvailableInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyData) SetInstanceList(v []*QueryAvailableInstancesResponseBodyDataInstanceList) *QueryAvailableInstancesResponseBodyData {
	s.InstanceList = v
	return s
}

type QueryAvailableInstancesResponseBodyDataInstanceList struct {
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SellerId            *int64  `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType         *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType    *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	InstanceID          *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	Region              *string `json:"Region,omitempty" xml:"Region,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime             *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StopTime            *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	ReleaseTime         *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	ExpectedReleaseTime *string `json:"ExpectedReleaseTime,omitempty" xml:"ExpectedReleaseTime,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubStatus           *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	RenewStatus         *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	RenewalDuration     *int32  `json:"RenewalDuration,omitempty" xml:"RenewalDuration,omitempty"`
	RenewalDurationUnit *string `json:"RenewalDurationUnit,omitempty" xml:"RenewalDurationUnit,omitempty"`
	Seller              *string `json:"Seller,omitempty" xml:"Seller,omitempty"`
}

func (s QueryAvailableInstancesResponseBodyDataInstanceList) String() string {
	return tea.Prettify(s)
}

func (s QueryAvailableInstancesResponseBodyDataInstanceList) GoString() string {
	return s.String()
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetOwnerId(v int64) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.OwnerId = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetSellerId(v int64) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.SellerId = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetProductCode(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.ProductCode = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetProductType(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.ProductType = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetSubscriptionType(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.SubscriptionType = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetInstanceID(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.InstanceID = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetRegion(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.Region = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetCreateTime(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.CreateTime = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetEndTime(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.EndTime = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetStopTime(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.StopTime = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetReleaseTime(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.ReleaseTime = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetExpectedReleaseTime(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.ExpectedReleaseTime = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetStatus(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.Status = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetSubStatus(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.SubStatus = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetRenewStatus(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.RenewStatus = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetRenewalDuration(v int32) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.RenewalDuration = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetRenewalDurationUnit(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.RenewalDurationUnit = &v
	return s
}

func (s *QueryAvailableInstancesResponseBodyDataInstanceList) SetSeller(v string) *QueryAvailableInstancesResponseBodyDataInstanceList {
	s.Seller = &v
	return s
}

type QueryAvailableInstancesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryAvailableInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAvailableInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAvailableInstancesResponse) GoString() string {
	return s.String()
}

func (s *QueryAvailableInstancesResponse) SetHeaders(v map[string]*string) *QueryAvailableInstancesResponse {
	s.Headers = v
	return s
}

func (s *QueryAvailableInstancesResponse) SetBody(v *QueryAvailableInstancesResponseBody) *QueryAvailableInstancesResponse {
	s.Body = v
	return s
}

type QueryBillRequest struct {
	BillingCycle           *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Type                   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ProductCode            *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType            *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType       *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	IsHideZeroCharge       *bool   `json:"IsHideZeroCharge,omitempty" xml:"IsHideZeroCharge,omitempty"`
	IsDisplayLocalCurrency *bool   `json:"IsDisplayLocalCurrency,omitempty" xml:"IsDisplayLocalCurrency,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum                *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize               *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	BillOwnerId            *int64  `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
}

func (s QueryBillRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBillRequest) GoString() string {
	return s.String()
}

func (s *QueryBillRequest) SetBillingCycle(v string) *QueryBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *QueryBillRequest) SetType(v string) *QueryBillRequest {
	s.Type = &v
	return s
}

func (s *QueryBillRequest) SetProductCode(v string) *QueryBillRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryBillRequest) SetProductType(v string) *QueryBillRequest {
	s.ProductType = &v
	return s
}

func (s *QueryBillRequest) SetSubscriptionType(v string) *QueryBillRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QueryBillRequest) SetIsHideZeroCharge(v bool) *QueryBillRequest {
	s.IsHideZeroCharge = &v
	return s
}

func (s *QueryBillRequest) SetIsDisplayLocalCurrency(v bool) *QueryBillRequest {
	s.IsDisplayLocalCurrency = &v
	return s
}

func (s *QueryBillRequest) SetOwnerId(v int64) *QueryBillRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryBillRequest) SetPageNum(v int32) *QueryBillRequest {
	s.PageNum = &v
	return s
}

func (s *QueryBillRequest) SetPageSize(v int32) *QueryBillRequest {
	s.PageSize = &v
	return s
}

func (s *QueryBillRequest) SetBillOwnerId(v int64) *QueryBillRequest {
	s.BillOwnerId = &v
	return s
}

type QueryBillResponseBody struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBillResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBillResponseBody) SetRequestId(v string) *QueryBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBillResponseBody) SetSuccess(v bool) *QueryBillResponseBody {
	s.Success = &v
	return s
}

func (s *QueryBillResponseBody) SetCode(v string) *QueryBillResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBillResponseBody) SetMessage(v string) *QueryBillResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBillResponseBody) SetData(v *QueryBillResponseBodyData) *QueryBillResponseBody {
	s.Data = v
	return s
}

type QueryBillResponseBodyData struct {
	BillingCycle *string                         `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	AccountID    *string                         `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	AccountName  *string                         `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	PageNum      *int32                          `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount   *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items        *QueryBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s QueryBillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBillResponseBodyData) SetBillingCycle(v string) *QueryBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QueryBillResponseBodyData) SetAccountID(v string) *QueryBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QueryBillResponseBodyData) SetAccountName(v string) *QueryBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryBillResponseBodyData) SetPageNum(v int32) *QueryBillResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryBillResponseBodyData) SetPageSize(v int32) *QueryBillResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryBillResponseBodyData) SetTotalCount(v int32) *QueryBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryBillResponseBodyData) SetItems(v *QueryBillResponseBodyDataItems) *QueryBillResponseBodyData {
	s.Items = v
	return s
}

type QueryBillResponseBodyDataItems struct {
	Item []*QueryBillResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryBillResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QueryBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryBillResponseBodyDataItems) SetItem(v []*QueryBillResponseBodyDataItemsItem) *QueryBillResponseBodyDataItems {
	s.Item = v
	return s
}

type QueryBillResponseBodyDataItemsItem struct {
	RecordID              *string  `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	Item                  *string  `json:"Item,omitempty" xml:"Item,omitempty"`
	OwnerID               *string  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	UsageStartTime        *string  `json:"UsageStartTime,omitempty" xml:"UsageStartTime,omitempty"`
	UsageEndTime          *string  `json:"UsageEndTime,omitempty" xml:"UsageEndTime,omitempty"`
	PaymentTime           *string  `json:"PaymentTime,omitempty" xml:"PaymentTime,omitempty"`
	ProductCode           *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType           *string  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType      *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	ProductName           *string  `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductDetail         *string  `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	PretaxGrossAmount     *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	DeductedByCoupons     *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	InvoiceDiscount       *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	PretaxAmount          *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	Currency              *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	PretaxAmountLocal     *float32 `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	Tax                   *float32 `json:"Tax,omitempty" xml:"Tax,omitempty"`
	PaymentAmount         *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	OutstandingAmount     *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	AfterTaxAmount        *float32 `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	Status                *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	PaymentCurrency       *string  `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	PaymentTransactionID  *string  `json:"PaymentTransactionID,omitempty" xml:"PaymentTransactionID,omitempty"`
	RoundDownDiscount     *string  `json:"RoundDownDiscount,omitempty" xml:"RoundDownDiscount,omitempty"`
	SubOrderId            *string  `json:"SubOrderId,omitempty" xml:"SubOrderId,omitempty"`
	PipCode               *string  `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	CommodityCode         *string  `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s QueryBillResponseBodyDataItemsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryBillResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QueryBillResponseBodyDataItemsItem) SetRecordID(v string) *QueryBillResponseBodyDataItemsItem {
	s.RecordID = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetItem(v string) *QueryBillResponseBodyDataItemsItem {
	s.Item = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetOwnerID(v string) *QueryBillResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetUsageStartTime(v string) *QueryBillResponseBodyDataItemsItem {
	s.UsageStartTime = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetUsageEndTime(v string) *QueryBillResponseBodyDataItemsItem {
	s.UsageEndTime = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPaymentTime(v string) *QueryBillResponseBodyDataItemsItem {
	s.PaymentTime = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetProductCode(v string) *QueryBillResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetProductType(v string) *QueryBillResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetSubscriptionType(v string) *QueryBillResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetProductName(v string) *QueryBillResponseBodyDataItemsItem {
	s.ProductName = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetProductDetail(v string) *QueryBillResponseBodyDataItemsItem {
	s.ProductDetail = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QueryBillResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetCurrency(v string) *QueryBillResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPretaxAmountLocal(v float32) *QueryBillResponseBodyDataItemsItem {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetTax(v float32) *QueryBillResponseBodyDataItemsItem {
	s.Tax = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryBillResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryBillResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetAfterTaxAmount(v float32) *QueryBillResponseBodyDataItemsItem {
	s.AfterTaxAmount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetStatus(v string) *QueryBillResponseBodyDataItemsItem {
	s.Status = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPaymentCurrency(v string) *QueryBillResponseBodyDataItemsItem {
	s.PaymentCurrency = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPaymentTransactionID(v string) *QueryBillResponseBodyDataItemsItem {
	s.PaymentTransactionID = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetRoundDownDiscount(v string) *QueryBillResponseBodyDataItemsItem {
	s.RoundDownDiscount = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetSubOrderId(v string) *QueryBillResponseBodyDataItemsItem {
	s.SubOrderId = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetPipCode(v string) *QueryBillResponseBodyDataItemsItem {
	s.PipCode = &v
	return s
}

func (s *QueryBillResponseBodyDataItemsItem) SetCommodityCode(v string) *QueryBillResponseBodyDataItemsItem {
	s.CommodityCode = &v
	return s
}

type QueryBillResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBillResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBillResponse) GoString() string {
	return s.String()
}

func (s *QueryBillResponse) SetHeaders(v map[string]*string) *QueryBillResponse {
	s.Headers = v
	return s
}

func (s *QueryBillResponse) SetBody(v *QueryBillResponseBody) *QueryBillResponse {
	s.Body = v
	return s
}

type QueryBillOverviewRequest struct {
	BillingCycle     *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	BillOwnerId      *int64  `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
}

func (s QueryBillOverviewRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBillOverviewRequest) GoString() string {
	return s.String()
}

func (s *QueryBillOverviewRequest) SetBillingCycle(v string) *QueryBillOverviewRequest {
	s.BillingCycle = &v
	return s
}

func (s *QueryBillOverviewRequest) SetProductCode(v string) *QueryBillOverviewRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryBillOverviewRequest) SetProductType(v string) *QueryBillOverviewRequest {
	s.ProductType = &v
	return s
}

func (s *QueryBillOverviewRequest) SetSubscriptionType(v string) *QueryBillOverviewRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QueryBillOverviewRequest) SetBillOwnerId(v int64) *QueryBillOverviewRequest {
	s.BillOwnerId = &v
	return s
}

type QueryBillOverviewResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryBillOverviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryBillOverviewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBillOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBillOverviewResponseBody) SetRequestId(v string) *QueryBillOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBillOverviewResponseBody) SetSuccess(v bool) *QueryBillOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *QueryBillOverviewResponseBody) SetCode(v string) *QueryBillOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBillOverviewResponseBody) SetMessage(v string) *QueryBillOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBillOverviewResponseBody) SetData(v *QueryBillOverviewResponseBodyData) *QueryBillOverviewResponseBody {
	s.Data = v
	return s
}

type QueryBillOverviewResponseBodyData struct {
	BillingCycle *string                                 `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	AccountID    *string                                 `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	AccountName  *string                                 `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Items        *QueryBillOverviewResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s QueryBillOverviewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryBillOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBillOverviewResponseBodyData) SetBillingCycle(v string) *QueryBillOverviewResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QueryBillOverviewResponseBodyData) SetAccountID(v string) *QueryBillOverviewResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QueryBillOverviewResponseBodyData) SetAccountName(v string) *QueryBillOverviewResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryBillOverviewResponseBodyData) SetItems(v *QueryBillOverviewResponseBodyDataItems) *QueryBillOverviewResponseBodyData {
	s.Items = v
	return s
}

type QueryBillOverviewResponseBodyDataItems struct {
	Item []*QueryBillOverviewResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryBillOverviewResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QueryBillOverviewResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryBillOverviewResponseBodyDataItems) SetItem(v []*QueryBillOverviewResponseBodyDataItemsItem) *QueryBillOverviewResponseBodyDataItems {
	s.Item = v
	return s
}

type QueryBillOverviewResponseBodyDataItemsItem struct {
	Item                  *string  `json:"Item,omitempty" xml:"Item,omitempty"`
	ProductCode           *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType           *string  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType      *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	ProductName           *string  `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductDetail         *string  `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	PretaxGrossAmount     *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	InvoiceDiscount       *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	DeductedByCoupons     *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	PretaxAmount          *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	Currency              *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	PaymentAmount         *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	OutstandingAmount     *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	PretaxAmountLocal     *float32 `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	Tax                   *float32 `json:"Tax,omitempty" xml:"Tax,omitempty"`
	AfterTaxAmount        *float32 `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	PaymentCurrency       *string  `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	RoundDownDiscount     *string  `json:"RoundDownDiscount,omitempty" xml:"RoundDownDiscount,omitempty"`
	PipCode               *string  `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	CommodityCode         *string  `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s QueryBillOverviewResponseBodyDataItemsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryBillOverviewResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetItem(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.Item = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetProductCode(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetProductType(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetSubscriptionType(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetProductName(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.ProductName = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetProductDetail(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.ProductDetail = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetCurrency(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetPretaxAmountLocal(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetTax(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.Tax = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetAfterTaxAmount(v float32) *QueryBillOverviewResponseBodyDataItemsItem {
	s.AfterTaxAmount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetPaymentCurrency(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.PaymentCurrency = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetRoundDownDiscount(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.RoundDownDiscount = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetPipCode(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.PipCode = &v
	return s
}

func (s *QueryBillOverviewResponseBodyDataItemsItem) SetCommodityCode(v string) *QueryBillOverviewResponseBodyDataItemsItem {
	s.CommodityCode = &v
	return s
}

type QueryBillOverviewResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBillOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBillOverviewResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBillOverviewResponse) GoString() string {
	return s.String()
}

func (s *QueryBillOverviewResponse) SetHeaders(v map[string]*string) *QueryBillOverviewResponse {
	s.Headers = v
	return s
}

func (s *QueryBillOverviewResponse) SetBody(v *QueryBillOverviewResponseBody) *QueryBillOverviewResponse {
	s.Body = v
	return s
}

type QueryBillToOSSSubscriptionResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryBillToOSSSubscriptionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryBillToOSSSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBillToOSSSubscriptionResponseBody) SetRequestId(v string) *QueryBillToOSSSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBody) SetSuccess(v bool) *QueryBillToOSSSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBody) SetCode(v string) *QueryBillToOSSSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBody) SetMessage(v string) *QueryBillToOSSSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBody) SetData(v *QueryBillToOSSSubscriptionResponseBodyData) *QueryBillToOSSSubscriptionResponseBody {
	s.Data = v
	return s
}

type QueryBillToOSSSubscriptionResponseBodyData struct {
	AccountID   *string                                          `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	AccountName *string                                          `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Items       *QueryBillToOSSSubscriptionResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s QueryBillToOSSSubscriptionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBillToOSSSubscriptionResponseBodyData) SetAccountID(v string) *QueryBillToOSSSubscriptionResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyData) SetAccountName(v string) *QueryBillToOSSSubscriptionResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyData) SetItems(v *QueryBillToOSSSubscriptionResponseBodyDataItems) *QueryBillToOSSSubscriptionResponseBodyData {
	s.Items = v
	return s
}

type QueryBillToOSSSubscriptionResponseBodyDataItems struct {
	Item []*QueryBillToOSSSubscriptionResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryBillToOSSSubscriptionResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItems) SetItem(v []*QueryBillToOSSSubscriptionResponseBodyDataItemsItem) *QueryBillToOSSSubscriptionResponseBodyDataItems {
	s.Item = v
	return s
}

type QueryBillToOSSSubscriptionResponseBodyDataItemsItem struct {
	SubscribeType     *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
	SubscribeBucket   *string `json:"SubscribeBucket,omitempty" xml:"SubscribeBucket,omitempty"`
	BucketOwnerId     *int64  `json:"BucketOwnerId,omitempty" xml:"BucketOwnerId,omitempty"`
	SubscribeTime     *string `json:"SubscribeTime,omitempty" xml:"SubscribeTime,omitempty"`
	SubscribeLanguage *string `json:"SubscribeLanguage,omitempty" xml:"SubscribeLanguage,omitempty"`
}

func (s QueryBillToOSSSubscriptionResponseBodyDataItemsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) SetSubscribeType(v string) *QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	s.SubscribeType = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) SetSubscribeBucket(v string) *QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	s.SubscribeBucket = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) SetBucketOwnerId(v int64) *QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	s.BucketOwnerId = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) SetSubscribeTime(v string) *QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	s.SubscribeTime = &v
	return s
}

func (s *QueryBillToOSSSubscriptionResponseBodyDataItemsItem) SetSubscribeLanguage(v string) *QueryBillToOSSSubscriptionResponseBodyDataItemsItem {
	s.SubscribeLanguage = &v
	return s
}

type QueryBillToOSSSubscriptionResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryBillToOSSSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBillToOSSSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBillToOSSSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *QueryBillToOSSSubscriptionResponse) SetHeaders(v map[string]*string) *QueryBillToOSSSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *QueryBillToOSSSubscriptionResponse) SetBody(v *QueryBillToOSSSubscriptionResponseBody) *QueryBillToOSSSubscriptionResponse {
	s.Body = v
	return s
}

type QueryCashCouponsRequest struct {
	ExpiryTimeEnd   *string `json:"ExpiryTimeEnd,omitempty" xml:"ExpiryTimeEnd,omitempty"`
	ExpiryTimeStart *string `json:"ExpiryTimeStart,omitempty" xml:"ExpiryTimeStart,omitempty"`
	EffectiveOrNot  *bool   `json:"EffectiveOrNot,omitempty" xml:"EffectiveOrNot,omitempty"`
}

func (s QueryCashCouponsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCashCouponsRequest) GoString() string {
	return s.String()
}

func (s *QueryCashCouponsRequest) SetExpiryTimeEnd(v string) *QueryCashCouponsRequest {
	s.ExpiryTimeEnd = &v
	return s
}

func (s *QueryCashCouponsRequest) SetExpiryTimeStart(v string) *QueryCashCouponsRequest {
	s.ExpiryTimeStart = &v
	return s
}

func (s *QueryCashCouponsRequest) SetEffectiveOrNot(v bool) *QueryCashCouponsRequest {
	s.EffectiveOrNot = &v
	return s
}

type QueryCashCouponsResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryCashCouponsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryCashCouponsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCashCouponsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCashCouponsResponseBody) SetRequestId(v string) *QueryCashCouponsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCashCouponsResponseBody) SetSuccess(v bool) *QueryCashCouponsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCashCouponsResponseBody) SetCode(v string) *QueryCashCouponsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCashCouponsResponseBody) SetMessage(v string) *QueryCashCouponsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCashCouponsResponseBody) SetData(v *QueryCashCouponsResponseBodyData) *QueryCashCouponsResponseBody {
	s.Data = v
	return s
}

type QueryCashCouponsResponseBodyData struct {
	CashCoupon []*QueryCashCouponsResponseBodyDataCashCoupon `json:"CashCoupon,omitempty" xml:"CashCoupon,omitempty" type:"Repeated"`
}

func (s QueryCashCouponsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCashCouponsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCashCouponsResponseBodyData) SetCashCoupon(v []*QueryCashCouponsResponseBodyDataCashCoupon) *QueryCashCouponsResponseBodyData {
	s.CashCoupon = v
	return s
}

type QueryCashCouponsResponseBodyDataCashCoupon struct {
	CashCouponId        *int64  `json:"CashCouponId,omitempty" xml:"CashCouponId,omitempty"`
	CashCouponNo        *string `json:"CashCouponNo,omitempty" xml:"CashCouponNo,omitempty"`
	GrantedTime         *string `json:"GrantedTime,omitempty" xml:"GrantedTime,omitempty"`
	EffectiveTime       *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ExpiryTime          *string `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	ApplicableProducts  *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	ApplicableScenarios *string `json:"ApplicableScenarios,omitempty" xml:"ApplicableScenarios,omitempty"`
	NominalValue        *string `json:"NominalValue,omitempty" xml:"NominalValue,omitempty"`
	Balance             *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryCashCouponsResponseBodyDataCashCoupon) String() string {
	return tea.Prettify(s)
}

func (s QueryCashCouponsResponseBodyDataCashCoupon) GoString() string {
	return s.String()
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetCashCouponId(v int64) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.CashCouponId = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetCashCouponNo(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.CashCouponNo = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetGrantedTime(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.GrantedTime = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetEffectiveTime(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.EffectiveTime = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetExpiryTime(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.ExpiryTime = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetApplicableProducts(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.ApplicableProducts = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetApplicableScenarios(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.ApplicableScenarios = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetNominalValue(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.NominalValue = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetBalance(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.Balance = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetStatus(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.Status = &v
	return s
}

type QueryCashCouponsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCashCouponsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCashCouponsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCashCouponsResponse) GoString() string {
	return s.String()
}

func (s *QueryCashCouponsResponse) SetHeaders(v map[string]*string) *QueryCashCouponsResponse {
	s.Headers = v
	return s
}

func (s *QueryCashCouponsResponse) SetBody(v *QueryCashCouponsResponseBody) *QueryCashCouponsResponse {
	s.Body = v
	return s
}

type QueryCostUnitRequest struct {
	OwnerUid     *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	ParentUnitId *int64 `json:"ParentUnitId,omitempty" xml:"ParentUnitId,omitempty"`
	PageNum      *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryCostUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCostUnitRequest) GoString() string {
	return s.String()
}

func (s *QueryCostUnitRequest) SetOwnerUid(v int64) *QueryCostUnitRequest {
	s.OwnerUid = &v
	return s
}

func (s *QueryCostUnitRequest) SetParentUnitId(v int64) *QueryCostUnitRequest {
	s.ParentUnitId = &v
	return s
}

func (s *QueryCostUnitRequest) SetPageNum(v int32) *QueryCostUnitRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCostUnitRequest) SetPageSize(v int32) *QueryCostUnitRequest {
	s.PageSize = &v
	return s
}

type QueryCostUnitResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *QueryCostUnitResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryCostUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCostUnitResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResponseBody) SetRequestId(v string) *QueryCostUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCostUnitResponseBody) SetMessage(v string) *QueryCostUnitResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCostUnitResponseBody) SetCode(v string) *QueryCostUnitResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCostUnitResponseBody) SetSuccess(v bool) *QueryCostUnitResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCostUnitResponseBody) SetData(v *QueryCostUnitResponseBodyData) *QueryCostUnitResponseBody {
	s.Data = v
	return s
}

type QueryCostUnitResponseBodyData struct {
	TotalCount      *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize        *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum         *int32                                          `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	CostUnitDtoList []*QueryCostUnitResponseBodyDataCostUnitDtoList `json:"CostUnitDtoList,omitempty" xml:"CostUnitDtoList,omitempty" type:"Repeated"`
}

func (s QueryCostUnitResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCostUnitResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResponseBodyData) SetTotalCount(v int32) *QueryCostUnitResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryCostUnitResponseBodyData) SetPageSize(v int32) *QueryCostUnitResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryCostUnitResponseBodyData) SetPageNum(v int32) *QueryCostUnitResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryCostUnitResponseBodyData) SetCostUnitDtoList(v []*QueryCostUnitResponseBodyDataCostUnitDtoList) *QueryCostUnitResponseBodyData {
	s.CostUnitDtoList = v
	return s
}

type QueryCostUnitResponseBodyDataCostUnitDtoList struct {
	UnitId       *int64  `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
	ParentUnitId *int64  `json:"ParentUnitId,omitempty" xml:"ParentUnitId,omitempty"`
	OwnerUid     *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	UnitName     *string `json:"UnitName,omitempty" xml:"UnitName,omitempty"`
}

func (s QueryCostUnitResponseBodyDataCostUnitDtoList) String() string {
	return tea.Prettify(s)
}

func (s QueryCostUnitResponseBodyDataCostUnitDtoList) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) SetUnitId(v int64) *QueryCostUnitResponseBodyDataCostUnitDtoList {
	s.UnitId = &v
	return s
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) SetParentUnitId(v int64) *QueryCostUnitResponseBodyDataCostUnitDtoList {
	s.ParentUnitId = &v
	return s
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) SetOwnerUid(v int64) *QueryCostUnitResponseBodyDataCostUnitDtoList {
	s.OwnerUid = &v
	return s
}

func (s *QueryCostUnitResponseBodyDataCostUnitDtoList) SetUnitName(v string) *QueryCostUnitResponseBodyDataCostUnitDtoList {
	s.UnitName = &v
	return s
}

type QueryCostUnitResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCostUnitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCostUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCostUnitResponse) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResponse) SetHeaders(v map[string]*string) *QueryCostUnitResponse {
	s.Headers = v
	return s
}

func (s *QueryCostUnitResponse) SetBody(v *QueryCostUnitResponseBody) *QueryCostUnitResponse {
	s.Body = v
	return s
}

type QueryCostUnitResourceRequest struct {
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	UnitId   *int64 `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
	PageNum  *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryCostUnitResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCostUnitResourceRequest) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceRequest) SetOwnerUid(v int64) *QueryCostUnitResourceRequest {
	s.OwnerUid = &v
	return s
}

func (s *QueryCostUnitResourceRequest) SetUnitId(v int64) *QueryCostUnitResourceRequest {
	s.UnitId = &v
	return s
}

func (s *QueryCostUnitResourceRequest) SetPageNum(v int32) *QueryCostUnitResourceRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCostUnitResourceRequest) SetPageSize(v int32) *QueryCostUnitResourceRequest {
	s.PageSize = &v
	return s
}

type QueryCostUnitResourceResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *QueryCostUnitResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryCostUnitResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCostUnitResourceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceResponseBody) SetRequestId(v string) *QueryCostUnitResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCostUnitResourceResponseBody) SetMessage(v string) *QueryCostUnitResourceResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCostUnitResourceResponseBody) SetCode(v string) *QueryCostUnitResourceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCostUnitResourceResponseBody) SetSuccess(v bool) *QueryCostUnitResourceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCostUnitResourceResponseBody) SetData(v *QueryCostUnitResourceResponseBodyData) *QueryCostUnitResourceResponseBody {
	s.Data = v
	return s
}

type QueryCostUnitResourceResponseBodyData struct {
	TotalCount              *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize                *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum                 *int32                                                          `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	ResourceInstanceDtoList []*QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList `json:"ResourceInstanceDtoList,omitempty" xml:"ResourceInstanceDtoList,omitempty" type:"Repeated"`
	CostUnit                *QueryCostUnitResourceResponseBodyDataCostUnit                  `json:"CostUnit,omitempty" xml:"CostUnit,omitempty" type:"Struct"`
	CostUnitStatisInfo      *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo        `json:"CostUnitStatisInfo,omitempty" xml:"CostUnitStatisInfo,omitempty" type:"Struct"`
}

func (s QueryCostUnitResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCostUnitResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceResponseBodyData) SetTotalCount(v int32) *QueryCostUnitResourceResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyData) SetPageSize(v int32) *QueryCostUnitResourceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyData) SetPageNum(v int32) *QueryCostUnitResourceResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyData) SetResourceInstanceDtoList(v []*QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) *QueryCostUnitResourceResponseBodyData {
	s.ResourceInstanceDtoList = v
	return s
}

func (s *QueryCostUnitResourceResponseBodyData) SetCostUnit(v *QueryCostUnitResourceResponseBodyDataCostUnit) *QueryCostUnitResourceResponseBodyData {
	s.CostUnit = v
	return s
}

func (s *QueryCostUnitResourceResponseBodyData) SetCostUnitStatisInfo(v *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) *QueryCostUnitResourceResponseBodyData {
	s.CostUnitStatisInfo = v
	return s
}

type QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList struct {
	ResourceUserId   *int64  `json:"ResourceUserId,omitempty" xml:"ResourceUserId,omitempty"`
	ResourceTag      *string `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty"`
	RelatedResources *string `json:"RelatedResources,omitempty" xml:"RelatedResources,omitempty"`
	ApportionName    *string `json:"ApportionName,omitempty" xml:"ApportionName,omitempty"`
	ResourceId       *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	CommodityCode    *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ResourceStatus   *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	ResourceType     *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceUserName *string `json:"ResourceUserName,omitempty" xml:"ResourceUserName,omitempty"`
	ResourceNick     *string `json:"ResourceNick,omitempty" xml:"ResourceNick,omitempty"`
	ResourceGroup    *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	CommodityName    *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	ApportionCode    *string `json:"ApportionCode,omitempty" xml:"ApportionCode,omitempty"`
}

func (s QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) String() string {
	return tea.Prettify(s)
}

func (s QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceUserId(v int64) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceUserId = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceTag(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceTag = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetRelatedResources(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.RelatedResources = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetApportionName(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ApportionName = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceId(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceId = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetCommodityCode(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.CommodityCode = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceStatus(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceStatus = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceType(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceType = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceUserName(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceUserName = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceNick(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceNick = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceGroup(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceGroup = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetCommodityName(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.CommodityName = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetApportionCode(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ApportionCode = &v
	return s
}

type QueryCostUnitResourceResponseBodyDataCostUnit struct {
	UnitId       *int64  `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
	ParentUnitId *int64  `json:"ParentUnitId,omitempty" xml:"ParentUnitId,omitempty"`
	OwnerUid     *int64  `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	UnitName     *string `json:"UnitName,omitempty" xml:"UnitName,omitempty"`
}

func (s QueryCostUnitResourceResponseBodyDataCostUnit) String() string {
	return tea.Prettify(s)
}

func (s QueryCostUnitResourceResponseBodyDataCostUnit) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) SetUnitId(v int64) *QueryCostUnitResourceResponseBodyDataCostUnit {
	s.UnitId = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) SetParentUnitId(v int64) *QueryCostUnitResourceResponseBodyDataCostUnit {
	s.ParentUnitId = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) SetOwnerUid(v int64) *QueryCostUnitResourceResponseBodyDataCostUnit {
	s.OwnerUid = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) SetUnitName(v string) *QueryCostUnitResourceResponseBodyDataCostUnit {
	s.UnitName = &v
	return s
}

type QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo struct {
	SubUnitCount            *int64 `json:"SubUnitCount,omitempty" xml:"SubUnitCount,omitempty"`
	TotalResourceGroupCount *int64 `json:"TotalResourceGroupCount,omitempty" xml:"TotalResourceGroupCount,omitempty"`
	TotalResourceCount      *int64 `json:"TotalResourceCount,omitempty" xml:"TotalResourceCount,omitempty"`
	UserCount               *int64 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	ResourceCount           *int64 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	TotalUserCount          *int64 `json:"TotalUserCount,omitempty" xml:"TotalUserCount,omitempty"`
	ResourceGroupCount      *int64 `json:"ResourceGroupCount,omitempty" xml:"ResourceGroupCount,omitempty"`
}

func (s QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetSubUnitCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.SubUnitCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetTotalResourceGroupCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.TotalResourceGroupCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetTotalResourceCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.TotalResourceCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetUserCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.UserCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetResourceCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.ResourceCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetTotalUserCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.TotalUserCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetResourceGroupCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.ResourceGroupCount = &v
	return s
}

type QueryCostUnitResourceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCostUnitResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCostUnitResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCostUnitResourceResponse) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceResponse) SetHeaders(v map[string]*string) *QueryCostUnitResourceResponse {
	s.Headers = v
	return s
}

func (s *QueryCostUnitResourceResponse) SetBody(v *QueryCostUnitResourceResponseBody) *QueryCostUnitResourceResponse {
	s.Body = v
	return s
}

type QueryCustomerAddressListRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s QueryCustomerAddressListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerAddressListRequest) GoString() string {
	return s.String()
}

func (s *QueryCustomerAddressListRequest) SetOwnerId(v int64) *QueryCustomerAddressListRequest {
	s.OwnerId = &v
	return s
}

type QueryCustomerAddressListResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryCustomerAddressListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryCustomerAddressListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerAddressListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomerAddressListResponseBody) SetRequestId(v string) *QueryCustomerAddressListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCustomerAddressListResponseBody) SetSuccess(v bool) *QueryCustomerAddressListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCustomerAddressListResponseBody) SetCode(v string) *QueryCustomerAddressListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCustomerAddressListResponseBody) SetMessage(v string) *QueryCustomerAddressListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCustomerAddressListResponseBody) SetData(v *QueryCustomerAddressListResponseBodyData) *QueryCustomerAddressListResponseBody {
	s.Data = v
	return s
}

type QueryCustomerAddressListResponseBodyData struct {
	CustomerInvoiceAddressList *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList `json:"CustomerInvoiceAddressList,omitempty" xml:"CustomerInvoiceAddressList,omitempty" type:"Struct"`
}

func (s QueryCustomerAddressListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerAddressListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCustomerAddressListResponseBodyData) SetCustomerInvoiceAddressList(v *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList) *QueryCustomerAddressListResponseBodyData {
	s.CustomerInvoiceAddressList = v
	return s
}

type QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList struct {
	CustomerInvoiceAddress []*QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress `json:"CustomerInvoiceAddress,omitempty" xml:"CustomerInvoiceAddress,omitempty" type:"Repeated"`
}

func (s QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList) GoString() string {
	return s.String()
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList) SetCustomerInvoiceAddress(v []*QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressList {
	s.CustomerInvoiceAddress = v
	return s
}

type QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress struct {
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	UserId          *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserNick        *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
	Addressee       *string `json:"Addressee,omitempty" xml:"Addressee,omitempty"`
	Province        *string `json:"Province,omitempty" xml:"Province,omitempty"`
	City            *string `json:"City,omitempty" xml:"City,omitempty"`
	County          *string `json:"County,omitempty" xml:"County,omitempty"`
	Street          *string `json:"Street,omitempty" xml:"Street,omitempty"`
	PostalCode      *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	Phone           *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	BizType         *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	DeliveryAddress *string `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
}

func (s QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) GoString() string {
	return s.String()
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetId(v int64) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.Id = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetUserId(v int64) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.UserId = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetUserNick(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.UserNick = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetAddressee(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.Addressee = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetProvince(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.Province = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetCity(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.City = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetCounty(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.County = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetStreet(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.Street = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetPostalCode(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.PostalCode = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetPhone(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.Phone = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetBizType(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.BizType = &v
	return s
}

func (s *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress) SetDeliveryAddress(v string) *QueryCustomerAddressListResponseBodyDataCustomerInvoiceAddressListCustomerInvoiceAddress {
	s.DeliveryAddress = &v
	return s
}

type QueryCustomerAddressListResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCustomerAddressListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCustomerAddressListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCustomerAddressListResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomerAddressListResponse) SetHeaders(v map[string]*string) *QueryCustomerAddressListResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomerAddressListResponse) SetBody(v *QueryCustomerAddressListResponseBody) *QueryCustomerAddressListResponse {
	s.Body = v
	return s
}

type QueryEvaluateListRequest struct {
	Type            *int32    `json:"Type,omitempty" xml:"Type,omitempty"`
	OutBizId        *string   `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
	OwnerId         *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum         *int32    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartAmount     *int64    `json:"StartAmount,omitempty" xml:"StartAmount,omitempty"`
	EndAmount       *int64    `json:"EndAmount,omitempty" xml:"EndAmount,omitempty"`
	StartBizTime    *string   `json:"StartBizTime,omitempty" xml:"StartBizTime,omitempty"`
	EndBizTime      *string   `json:"EndBizTime,omitempty" xml:"EndBizTime,omitempty"`
	SortType        *int32    `json:"SortType,omitempty" xml:"SortType,omitempty"`
	StartSearchTime *string   `json:"StartSearchTime,omitempty" xml:"StartSearchTime,omitempty"`
	EndSearchTime   *string   `json:"EndSearchTime,omitempty" xml:"EndSearchTime,omitempty"`
	BillCycle       *string   `json:"BillCycle,omitempty" xml:"BillCycle,omitempty"`
	BizTypeList     []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
}

func (s QueryEvaluateListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEvaluateListRequest) GoString() string {
	return s.String()
}

func (s *QueryEvaluateListRequest) SetType(v int32) *QueryEvaluateListRequest {
	s.Type = &v
	return s
}

func (s *QueryEvaluateListRequest) SetOutBizId(v string) *QueryEvaluateListRequest {
	s.OutBizId = &v
	return s
}

func (s *QueryEvaluateListRequest) SetOwnerId(v int64) *QueryEvaluateListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryEvaluateListRequest) SetPageNum(v int32) *QueryEvaluateListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryEvaluateListRequest) SetPageSize(v int32) *QueryEvaluateListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryEvaluateListRequest) SetStartAmount(v int64) *QueryEvaluateListRequest {
	s.StartAmount = &v
	return s
}

func (s *QueryEvaluateListRequest) SetEndAmount(v int64) *QueryEvaluateListRequest {
	s.EndAmount = &v
	return s
}

func (s *QueryEvaluateListRequest) SetStartBizTime(v string) *QueryEvaluateListRequest {
	s.StartBizTime = &v
	return s
}

func (s *QueryEvaluateListRequest) SetEndBizTime(v string) *QueryEvaluateListRequest {
	s.EndBizTime = &v
	return s
}

func (s *QueryEvaluateListRequest) SetSortType(v int32) *QueryEvaluateListRequest {
	s.SortType = &v
	return s
}

func (s *QueryEvaluateListRequest) SetStartSearchTime(v string) *QueryEvaluateListRequest {
	s.StartSearchTime = &v
	return s
}

func (s *QueryEvaluateListRequest) SetEndSearchTime(v string) *QueryEvaluateListRequest {
	s.EndSearchTime = &v
	return s
}

func (s *QueryEvaluateListRequest) SetBillCycle(v string) *QueryEvaluateListRequest {
	s.BillCycle = &v
	return s
}

func (s *QueryEvaluateListRequest) SetBizTypeList(v []*string) *QueryEvaluateListRequest {
	s.BizTypeList = v
	return s
}

type QueryEvaluateListResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryEvaluateListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryEvaluateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEvaluateListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEvaluateListResponseBody) SetRequestId(v string) *QueryEvaluateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEvaluateListResponseBody) SetSuccess(v bool) *QueryEvaluateListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryEvaluateListResponseBody) SetCode(v string) *QueryEvaluateListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEvaluateListResponseBody) SetMessage(v string) *QueryEvaluateListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEvaluateListResponseBody) SetData(v *QueryEvaluateListResponseBodyData) *QueryEvaluateListResponseBody {
	s.Data = v
	return s
}

type QueryEvaluateListResponseBodyData struct {
	HostId                      *string                                        `json:"HostId,omitempty" xml:"HostId,omitempty"`
	PageNum                     *int32                                         `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize                    *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount                  *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalInvoiceAmount          *int64                                         `json:"TotalInvoiceAmount,omitempty" xml:"TotalInvoiceAmount,omitempty"`
	TotalUnAppliedInvoiceAmount *int64                                         `json:"TotalUnAppliedInvoiceAmount,omitempty" xml:"TotalUnAppliedInvoiceAmount,omitempty"`
	EvaluateList                *QueryEvaluateListResponseBodyDataEvaluateList `json:"EvaluateList,omitempty" xml:"EvaluateList,omitempty" type:"Struct"`
}

func (s QueryEvaluateListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEvaluateListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEvaluateListResponseBodyData) SetHostId(v string) *QueryEvaluateListResponseBodyData {
	s.HostId = &v
	return s
}

func (s *QueryEvaluateListResponseBodyData) SetPageNum(v int32) *QueryEvaluateListResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryEvaluateListResponseBodyData) SetPageSize(v int32) *QueryEvaluateListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryEvaluateListResponseBodyData) SetTotalCount(v int32) *QueryEvaluateListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyData) SetTotalInvoiceAmount(v int64) *QueryEvaluateListResponseBodyData {
	s.TotalInvoiceAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyData) SetTotalUnAppliedInvoiceAmount(v int64) *QueryEvaluateListResponseBodyData {
	s.TotalUnAppliedInvoiceAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyData) SetEvaluateList(v *QueryEvaluateListResponseBodyDataEvaluateList) *QueryEvaluateListResponseBodyData {
	s.EvaluateList = v
	return s
}

type QueryEvaluateListResponseBodyDataEvaluateList struct {
	Evaluate []*QueryEvaluateListResponseBodyDataEvaluateListEvaluate `json:"Evaluate,omitempty" xml:"Evaluate,omitempty" type:"Repeated"`
}

func (s QueryEvaluateListResponseBodyDataEvaluateList) String() string {
	return tea.Prettify(s)
}

func (s QueryEvaluateListResponseBodyDataEvaluateList) GoString() string {
	return s.String()
}

func (s *QueryEvaluateListResponseBodyDataEvaluateList) SetEvaluate(v []*QueryEvaluateListResponseBodyDataEvaluateListEvaluate) *QueryEvaluateListResponseBodyDataEvaluateList {
	s.Evaluate = v
	return s
}

type QueryEvaluateListResponseBodyDataEvaluateListEvaluate struct {
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	GmtCreate          *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified        *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	UserId             *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserNick           *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
	OutBizId           *string `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
	BillId             *int64  `json:"BillId,omitempty" xml:"BillId,omitempty"`
	ItemId             *int64  `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	BillCycle          *string `json:"BillCycle,omitempty" xml:"BillCycle,omitempty"`
	BizType            *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OriginalAmount     *int64  `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	PresentAmount      *int64  `json:"PresentAmount,omitempty" xml:"PresentAmount,omitempty"`
	CanInvoiceAmount   *int64  `json:"CanInvoiceAmount,omitempty" xml:"CanInvoiceAmount,omitempty"`
	InvoicedAmount     *int64  `json:"InvoicedAmount,omitempty" xml:"InvoicedAmount,omitempty"`
	OffsetCostAmount   *int64  `json:"OffsetCostAmount,omitempty" xml:"OffsetCostAmount,omitempty"`
	OffsetAcceptAmount *int64  `json:"OffsetAcceptAmount,omitempty" xml:"OffsetAcceptAmount,omitempty"`
	Status             *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	OpId               *string `json:"OpId,omitempty" xml:"OpId,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	BizTime            *string `json:"BizTime,omitempty" xml:"BizTime,omitempty"`
	Type               *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryEvaluateListResponseBodyDataEvaluateListEvaluate) String() string {
	return tea.Prettify(s)
}

func (s QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GoString() string {
	return s.String()
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetId(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.Id = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetGmtCreate(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.GmtCreate = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetGmtModified(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.GmtModified = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetUserId(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.UserId = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetUserNick(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.UserNick = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetOutBizId(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.OutBizId = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetBillId(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.BillId = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetItemId(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.ItemId = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetBillCycle(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.BillCycle = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetBizType(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.BizType = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetOriginalAmount(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.OriginalAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetPresentAmount(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.PresentAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetCanInvoiceAmount(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.CanInvoiceAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetInvoicedAmount(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.InvoicedAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetOffsetCostAmount(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.OffsetCostAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetOffsetAcceptAmount(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.OffsetAcceptAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetStatus(v int32) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.Status = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetOpId(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.OpId = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetName(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.Name = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetBizTime(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.BizTime = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetType(v int32) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.Type = &v
	return s
}

type QueryEvaluateListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryEvaluateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryEvaluateListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEvaluateListResponse) GoString() string {
	return s.String()
}

func (s *QueryEvaluateListResponse) SetHeaders(v map[string]*string) *QueryEvaluateListResponse {
	s.Headers = v
	return s
}

func (s *QueryEvaluateListResponse) SetBody(v *QueryEvaluateListResponseBody) *QueryEvaluateListResponse {
	s.Body = v
	return s
}

type QueryFinancialAccountInfoRequest struct {
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryFinancialAccountInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFinancialAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryFinancialAccountInfoRequest) SetUserId(v int64) *QueryFinancialAccountInfoRequest {
	s.UserId = &v
	return s
}

type QueryFinancialAccountInfoResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryFinancialAccountInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryFinancialAccountInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFinancialAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFinancialAccountInfoResponseBody) SetCode(v string) *QueryFinancialAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBody) SetRequestId(v string) *QueryFinancialAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBody) SetSuccess(v bool) *QueryFinancialAccountInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBody) SetMessage(v string) *QueryFinancialAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBody) SetData(v *QueryFinancialAccountInfoResponseBodyData) *QueryFinancialAccountInfoResponseBody {
	s.Data = v
	return s
}

type QueryFinancialAccountInfoResponseBodyData struct {
	UserName           *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	IsFinancialAccount *bool   `json:"IsFinancialAccount,omitempty" xml:"IsFinancialAccount,omitempty"`
	AccountType        *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	MemberNickName     *string `json:"MemberNickName,omitempty" xml:"MemberNickName,omitempty"`
	MemberGroupId      *int64  `json:"MemberGroupId,omitempty" xml:"MemberGroupId,omitempty"`
	MemberGroupName    *string `json:"MemberGroupName,omitempty" xml:"MemberGroupName,omitempty"`
}

func (s QueryFinancialAccountInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFinancialAccountInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFinancialAccountInfoResponseBodyData) SetUserName(v string) *QueryFinancialAccountInfoResponseBodyData {
	s.UserName = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBodyData) SetIsFinancialAccount(v bool) *QueryFinancialAccountInfoResponseBodyData {
	s.IsFinancialAccount = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBodyData) SetAccountType(v string) *QueryFinancialAccountInfoResponseBodyData {
	s.AccountType = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBodyData) SetMemberNickName(v string) *QueryFinancialAccountInfoResponseBodyData {
	s.MemberNickName = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBodyData) SetMemberGroupId(v int64) *QueryFinancialAccountInfoResponseBodyData {
	s.MemberGroupId = &v
	return s
}

func (s *QueryFinancialAccountInfoResponseBodyData) SetMemberGroupName(v string) *QueryFinancialAccountInfoResponseBodyData {
	s.MemberGroupName = &v
	return s
}

type QueryFinancialAccountInfoResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryFinancialAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFinancialAccountInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFinancialAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryFinancialAccountInfoResponse) SetHeaders(v map[string]*string) *QueryFinancialAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryFinancialAccountInfoResponse) SetBody(v *QueryFinancialAccountInfoResponseBody) *QueryFinancialAccountInfoResponse {
	s.Body = v
	return s
}

type QueryInstanceBillRequest struct {
	BillingCycle     *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	IsBillingItem    *bool   `json:"IsBillingItem,omitempty" xml:"IsBillingItem,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	IsHideZeroCharge *bool   `json:"IsHideZeroCharge,omitempty" xml:"IsHideZeroCharge,omitempty"`
	BillingDate      *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	Granularity      *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	BillOwnerId      *int64  `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
}

func (s QueryInstanceBillRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceBillRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceBillRequest) SetBillingCycle(v string) *QueryInstanceBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *QueryInstanceBillRequest) SetProductCode(v string) *QueryInstanceBillRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryInstanceBillRequest) SetProductType(v string) *QueryInstanceBillRequest {
	s.ProductType = &v
	return s
}

func (s *QueryInstanceBillRequest) SetSubscriptionType(v string) *QueryInstanceBillRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QueryInstanceBillRequest) SetOwnerId(v int64) *QueryInstanceBillRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryInstanceBillRequest) SetIsBillingItem(v bool) *QueryInstanceBillRequest {
	s.IsBillingItem = &v
	return s
}

func (s *QueryInstanceBillRequest) SetPageNum(v int32) *QueryInstanceBillRequest {
	s.PageNum = &v
	return s
}

func (s *QueryInstanceBillRequest) SetPageSize(v int32) *QueryInstanceBillRequest {
	s.PageSize = &v
	return s
}

func (s *QueryInstanceBillRequest) SetIsHideZeroCharge(v bool) *QueryInstanceBillRequest {
	s.IsHideZeroCharge = &v
	return s
}

func (s *QueryInstanceBillRequest) SetBillingDate(v string) *QueryInstanceBillRequest {
	s.BillingDate = &v
	return s
}

func (s *QueryInstanceBillRequest) SetGranularity(v string) *QueryInstanceBillRequest {
	s.Granularity = &v
	return s
}

func (s *QueryInstanceBillRequest) SetBillOwnerId(v int64) *QueryInstanceBillRequest {
	s.BillOwnerId = &v
	return s
}

type QueryInstanceBillResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryInstanceBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryInstanceBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceBillResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceBillResponseBody) SetRequestId(v string) *QueryInstanceBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceBillResponseBody) SetSuccess(v bool) *QueryInstanceBillResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInstanceBillResponseBody) SetCode(v string) *QueryInstanceBillResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstanceBillResponseBody) SetMessage(v string) *QueryInstanceBillResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstanceBillResponseBody) SetData(v *QueryInstanceBillResponseBodyData) *QueryInstanceBillResponseBody {
	s.Data = v
	return s
}

type QueryInstanceBillResponseBodyData struct {
	BillingCycle *string                                 `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	AccountID    *string                                 `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	AccountName  *string                                 `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	TotalCount   *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageNum      *int32                                  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Items        *QueryInstanceBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s QueryInstanceBillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInstanceBillResponseBodyData) SetBillingCycle(v string) *QueryInstanceBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QueryInstanceBillResponseBodyData) SetAccountID(v string) *QueryInstanceBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QueryInstanceBillResponseBodyData) SetAccountName(v string) *QueryInstanceBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryInstanceBillResponseBodyData) SetTotalCount(v int32) *QueryInstanceBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyData) SetPageNum(v int32) *QueryInstanceBillResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryInstanceBillResponseBodyData) SetPageSize(v int32) *QueryInstanceBillResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryInstanceBillResponseBodyData) SetItems(v *QueryInstanceBillResponseBodyDataItems) *QueryInstanceBillResponseBodyData {
	s.Items = v
	return s
}

type QueryInstanceBillResponseBodyDataItems struct {
	Item []*QueryInstanceBillResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryInstanceBillResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryInstanceBillResponseBodyDataItems) SetItem(v []*QueryInstanceBillResponseBodyDataItemsItem) *QueryInstanceBillResponseBodyDataItems {
	s.Item = v
	return s
}

type QueryInstanceBillResponseBodyDataItemsItem struct {
	InstanceID                *string  `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	BillingType               *string  `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CostUnit                  *string  `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
	ProductCode               *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType               *string  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType          *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	ProductName               *string  `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductDetail             *string  `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	OwnerID                   *string  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	BillingItem               *string  `json:"BillingItem,omitempty" xml:"BillingItem,omitempty"`
	ListPrice                 *string  `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	ListPriceUnit             *string  `json:"ListPriceUnit,omitempty" xml:"ListPriceUnit,omitempty"`
	Usage                     *string  `json:"Usage,omitempty" xml:"Usage,omitempty"`
	UsageUnit                 *string  `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty"`
	DeductedByResourcePackage *string  `json:"DeductedByResourcePackage,omitempty" xml:"DeductedByResourcePackage,omitempty"`
	PretaxGrossAmount         *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	InvoiceDiscount           *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	DeductedByCoupons         *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	PretaxAmount              *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	DeductedByCashCoupons     *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByPrepaidCard     *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	PaymentAmount             *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	OutstandingAmount         *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	Currency                  *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	NickName                  *string  `json:"NickName,omitempty" xml:"NickName,omitempty"`
	ResourceGroup             *string  `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	Tag                       *string  `json:"Tag,omitempty" xml:"Tag,omitempty"`
	InstanceConfig            *string  `json:"InstanceConfig,omitempty" xml:"InstanceConfig,omitempty"`
	InstanceSpec              *string  `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	InternetIP                *string  `json:"InternetIP,omitempty" xml:"InternetIP,omitempty"`
	IntranetIP                *string  `json:"IntranetIP,omitempty" xml:"IntranetIP,omitempty"`
	Region                    *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	Zone                      *string  `json:"Zone,omitempty" xml:"Zone,omitempty"`
	Item                      *string  `json:"Item,omitempty" xml:"Item,omitempty"`
	ServicePeriod             *string  `json:"ServicePeriod,omitempty" xml:"ServicePeriod,omitempty"`
	BillingDate               *string  `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	ServicePeriodUnit         *string  `json:"ServicePeriodUnit,omitempty" xml:"ServicePeriodUnit,omitempty"`
	PipCode                   *string  `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	CommodityCode             *string  `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s QueryInstanceBillResponseBodyDataItemsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceBillResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetInstanceID(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.InstanceID = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetBillingType(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.BillingType = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetCostUnit(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.CostUnit = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetProductCode(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetProductType(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetSubscriptionType(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetProductName(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ProductName = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetProductDetail(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ProductDetail = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetOwnerID(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetBillingItem(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.BillingItem = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetListPrice(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ListPrice = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetListPriceUnit(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ListPriceUnit = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetUsage(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.Usage = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetUsageUnit(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.UsageUnit = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetDeductedByResourcePackage(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.DeductedByResourcePackage = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QueryInstanceBillResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetCurrency(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetNickName(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.NickName = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetResourceGroup(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ResourceGroup = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetTag(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.Tag = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetInstanceConfig(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.InstanceConfig = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetInstanceSpec(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.InstanceSpec = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetInternetIP(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.InternetIP = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetIntranetIP(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.IntranetIP = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetRegion(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.Region = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetZone(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.Zone = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetItem(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.Item = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetServicePeriod(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ServicePeriod = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetBillingDate(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.BillingDate = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetServicePeriodUnit(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.ServicePeriodUnit = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetPipCode(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.PipCode = &v
	return s
}

func (s *QueryInstanceBillResponseBodyDataItemsItem) SetCommodityCode(v string) *QueryInstanceBillResponseBodyDataItemsItem {
	s.CommodityCode = &v
	return s
}

type QueryInstanceBillResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryInstanceBillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryInstanceBillResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceBillResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceBillResponse) SetHeaders(v map[string]*string) *QueryInstanceBillResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceBillResponse) SetBody(v *QueryInstanceBillResponseBody) *QueryInstanceBillResponse {
	s.Body = v
	return s
}

type QueryInstanceByTagRequest struct {
	ResourceType *string                         `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   []*string                       `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*QueryInstanceByTagRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s QueryInstanceByTagRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceByTagRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceByTagRequest) SetResourceType(v string) *QueryInstanceByTagRequest {
	s.ResourceType = &v
	return s
}

func (s *QueryInstanceByTagRequest) SetResourceId(v []*string) *QueryInstanceByTagRequest {
	s.ResourceId = v
	return s
}

func (s *QueryInstanceByTagRequest) SetTag(v []*QueryInstanceByTagRequestTag) *QueryInstanceByTagRequest {
	s.Tag = v
	return s
}

type QueryInstanceByTagRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryInstanceByTagRequestTag) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceByTagRequestTag) GoString() string {
	return s.String()
}

func (s *QueryInstanceByTagRequestTag) SetKey(v string) *QueryInstanceByTagRequestTag {
	s.Key = &v
	return s
}

func (s *QueryInstanceByTagRequestTag) SetValue(v string) *QueryInstanceByTagRequestTag {
	s.Value = &v
	return s
}

type QueryInstanceByTagResponseBody struct {
	Code        *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Message     *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken   *string                                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TagResource []*QueryInstanceByTagResponseBodyTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s QueryInstanceByTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceByTagResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceByTagResponseBody) SetCode(v string) *QueryInstanceByTagResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstanceByTagResponseBody) SetRequestId(v string) *QueryInstanceByTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceByTagResponseBody) SetSuccess(v bool) *QueryInstanceByTagResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInstanceByTagResponseBody) SetMessage(v string) *QueryInstanceByTagResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstanceByTagResponseBody) SetNextToken(v string) *QueryInstanceByTagResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryInstanceByTagResponseBody) SetTagResource(v []*QueryInstanceByTagResponseBodyTagResource) *QueryInstanceByTagResponseBody {
	s.TagResource = v
	return s
}

type QueryInstanceByTagResponseBodyTagResource struct {
	ResourceId   *string                                         `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string                                         `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*QueryInstanceByTagResponseBodyTagResourceTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s QueryInstanceByTagResponseBodyTagResource) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceByTagResponseBodyTagResource) GoString() string {
	return s.String()
}

func (s *QueryInstanceByTagResponseBodyTagResource) SetResourceId(v string) *QueryInstanceByTagResponseBodyTagResource {
	s.ResourceId = &v
	return s
}

func (s *QueryInstanceByTagResponseBodyTagResource) SetResourceType(v string) *QueryInstanceByTagResponseBodyTagResource {
	s.ResourceType = &v
	return s
}

func (s *QueryInstanceByTagResponseBodyTagResource) SetTag(v []*QueryInstanceByTagResponseBodyTagResourceTag) *QueryInstanceByTagResponseBodyTagResource {
	s.Tag = v
	return s
}

type QueryInstanceByTagResponseBodyTagResourceTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryInstanceByTagResponseBodyTagResourceTag) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceByTagResponseBodyTagResourceTag) GoString() string {
	return s.String()
}

func (s *QueryInstanceByTagResponseBodyTagResourceTag) SetKey(v string) *QueryInstanceByTagResponseBodyTagResourceTag {
	s.Key = &v
	return s
}

func (s *QueryInstanceByTagResponseBodyTagResourceTag) SetValue(v string) *QueryInstanceByTagResponseBodyTagResourceTag {
	s.Value = &v
	return s
}

type QueryInstanceByTagResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryInstanceByTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryInstanceByTagResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceByTagResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceByTagResponse) SetHeaders(v map[string]*string) *QueryInstanceByTagResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceByTagResponse) SetBody(v *QueryInstanceByTagResponseBody) *QueryInstanceByTagResponse {
	s.Body = v
	return s
}

type QueryInstanceGaapCostRequest struct {
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	BillingCycle     *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s QueryInstanceGaapCostRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceGaapCostRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceGaapCostRequest) SetPageNum(v int32) *QueryInstanceGaapCostRequest {
	s.PageNum = &v
	return s
}

func (s *QueryInstanceGaapCostRequest) SetPageSize(v int32) *QueryInstanceGaapCostRequest {
	s.PageSize = &v
	return s
}

func (s *QueryInstanceGaapCostRequest) SetBillingCycle(v string) *QueryInstanceGaapCostRequest {
	s.BillingCycle = &v
	return s
}

func (s *QueryInstanceGaapCostRequest) SetProductCode(v string) *QueryInstanceGaapCostRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryInstanceGaapCostRequest) SetProductType(v string) *QueryInstanceGaapCostRequest {
	s.ProductType = &v
	return s
}

func (s *QueryInstanceGaapCostRequest) SetSubscriptionType(v string) *QueryInstanceGaapCostRequest {
	s.SubscriptionType = &v
	return s
}

type QueryInstanceGaapCostResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryInstanceGaapCostResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryInstanceGaapCostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceGaapCostResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceGaapCostResponseBody) SetRequestId(v string) *QueryInstanceGaapCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBody) SetSuccess(v bool) *QueryInstanceGaapCostResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBody) SetCode(v string) *QueryInstanceGaapCostResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBody) SetMessage(v string) *QueryInstanceGaapCostResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBody) SetData(v *QueryInstanceGaapCostResponseBodyData) *QueryInstanceGaapCostResponseBody {
	s.Data = v
	return s
}

type QueryInstanceGaapCostResponseBodyData struct {
	HostId     *string                                       `json:"HostId,omitempty" xml:"HostId,omitempty"`
	PageNum    *int32                                        `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Modules    *QueryInstanceGaapCostResponseBodyDataModules `json:"Modules,omitempty" xml:"Modules,omitempty" type:"Struct"`
}

func (s QueryInstanceGaapCostResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceGaapCostResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInstanceGaapCostResponseBodyData) SetHostId(v string) *QueryInstanceGaapCostResponseBodyData {
	s.HostId = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyData) SetPageNum(v int32) *QueryInstanceGaapCostResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyData) SetPageSize(v int32) *QueryInstanceGaapCostResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyData) SetTotalCount(v int32) *QueryInstanceGaapCostResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyData) SetModules(v *QueryInstanceGaapCostResponseBodyDataModules) *QueryInstanceGaapCostResponseBodyData {
	s.Modules = v
	return s
}

type QueryInstanceGaapCostResponseBodyDataModules struct {
	Module []*QueryInstanceGaapCostResponseBodyDataModulesModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
}

func (s QueryInstanceGaapCostResponseBodyDataModules) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceGaapCostResponseBodyDataModules) GoString() string {
	return s.String()
}

func (s *QueryInstanceGaapCostResponseBodyDataModules) SetModule(v []*QueryInstanceGaapCostResponseBodyDataModulesModule) *QueryInstanceGaapCostResponseBodyDataModules {
	s.Module = v
	return s
}

type QueryInstanceGaapCostResponseBodyDataModulesModule struct {
	BillingCycle                     *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	InstanceID                       *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	ProductCode                      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType                      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType                 *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	Tag                              *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	ResourceGroup                    *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	AccountingUnit                   *string `json:"AccountingUnit,omitempty" xml:"AccountingUnit,omitempty"`
	PayerAccount                     *string `json:"PayerAccount,omitempty" xml:"PayerAccount,omitempty"`
	OwnerID                          *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	Region                           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Currency                         *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	PaymentCurrency                  *string `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	OrderType                        *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PayTime                          *string `json:"PayTime,omitempty" xml:"PayTime,omitempty"`
	PretaxGrossAmount                *string `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	PricingDiscount                  *string `json:"PricingDiscount,omitempty" xml:"PricingDiscount,omitempty"`
	DeductedByCoupons                *string `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	PretaxAmount                     *string `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	PretaxAmountLocal                *string `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	DeductedByCashCoupons            *string `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByPrepaidCard            *string `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	PaymentAmount                    *string `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	GaapPretaxGrossAmount            *string `json:"GaapPretaxGrossAmount,omitempty" xml:"GaapPretaxGrossAmount,omitempty"`
	GaapPricingDiscount              *string `json:"GaapPricingDiscount,omitempty" xml:"GaapPricingDiscount,omitempty"`
	GaapDeductedByCoupons            *string `json:"GaapDeductedByCoupons,omitempty" xml:"GaapDeductedByCoupons,omitempty"`
	GaapPretaxAmount                 *string `json:"GaapPretaxAmount,omitempty" xml:"GaapPretaxAmount,omitempty"`
	GaapPretaxAmountLocal            *string `json:"GaapPretaxAmountLocal,omitempty" xml:"GaapPretaxAmountLocal,omitempty"`
	GaapDeductedByCashCoupons        *string `json:"GaapDeductedByCashCoupons,omitempty" xml:"GaapDeductedByCashCoupons,omitempty"`
	GaapDeductedByPrepaidCard        *string `json:"GaapDeductedByPrepaidCard,omitempty" xml:"GaapDeductedByPrepaidCard,omitempty"`
	GaapPaymentAmount                *string `json:"GaapPaymentAmount,omitempty" xml:"GaapPaymentAmount,omitempty"`
	MonthGaapPretaxGrossAmount       *string `json:"MonthGaapPretaxGrossAmount,omitempty" xml:"MonthGaapPretaxGrossAmount,omitempty"`
	MonthGaapPricingDiscount         *string `json:"MonthGaapPricingDiscount,omitempty" xml:"MonthGaapPricingDiscount,omitempty"`
	MonthGaapDeductedByCoupons       *string `json:"MonthGaapDeductedByCoupons,omitempty" xml:"MonthGaapDeductedByCoupons,omitempty"`
	MonthGaapPretaxAmount            *string `json:"MonthGaapPretaxAmount,omitempty" xml:"MonthGaapPretaxAmount,omitempty"`
	MonthGaapPretaxAmountLocal       *string `json:"MonthGaapPretaxAmountLocal,omitempty" xml:"MonthGaapPretaxAmountLocal,omitempty"`
	MonthGaapDeductedByCashCoupons   *string `json:"MonthGaapDeductedByCashCoupons,omitempty" xml:"MonthGaapDeductedByCashCoupons,omitempty"`
	MonthGaapDeductedByPrepaidCard   *string `json:"MonthGaapDeductedByPrepaidCard,omitempty" xml:"MonthGaapDeductedByPrepaidCard,omitempty"`
	MonthGaapPaymentAmount           *string `json:"MonthGaapPaymentAmount,omitempty" xml:"MonthGaapPaymentAmount,omitempty"`
	UnallocatedPaymentAmount         *string `json:"UnallocatedPaymentAmount,omitempty" xml:"UnallocatedPaymentAmount,omitempty"`
	UsageStartDate                   *string `json:"UsageStartDate,omitempty" xml:"UsageStartDate,omitempty"`
	UsageEndDate                     *string `json:"UsageEndDate,omitempty" xml:"UsageEndDate,omitempty"`
	BillType                         *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	OrderId                          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	SubOrderId                       *string `json:"SubOrderId,omitempty" xml:"SubOrderId,omitempty"`
	UnallocatedPretaxGrossAmount     *string `json:"UnallocatedPretaxGrossAmount,omitempty" xml:"UnallocatedPretaxGrossAmount,omitempty"`
	UnallocatedPricingDiscount       *string `json:"UnallocatedPricingDiscount,omitempty" xml:"UnallocatedPricingDiscount,omitempty"`
	UnallocatedDeductedByCoupons     *string `json:"UnallocatedDeductedByCoupons,omitempty" xml:"UnallocatedDeductedByCoupons,omitempty"`
	UnallocatedPretaxAmount          *string `json:"UnallocatedPretaxAmount,omitempty" xml:"UnallocatedPretaxAmount,omitempty"`
	UnallocatedPretaxAmountLocal     *string `json:"UnallocatedPretaxAmountLocal,omitempty" xml:"UnallocatedPretaxAmountLocal,omitempty"`
	UnallocatedDeductedByCashCoupons *string `json:"UnallocatedDeductedByCashCoupons,omitempty" xml:"UnallocatedDeductedByCashCoupons,omitempty"`
	UnallocatedDeductedByPrepaidCard *string `json:"UnallocatedDeductedByPrepaidCard,omitempty" xml:"UnallocatedDeductedByPrepaidCard,omitempty"`
}

func (s QueryInstanceGaapCostResponseBodyDataModulesModule) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceGaapCostResponseBodyDataModulesModule) GoString() string {
	return s.String()
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetBillingCycle(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.BillingCycle = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetInstanceID(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.InstanceID = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetProductCode(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.ProductCode = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetProductType(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.ProductType = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetSubscriptionType(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.SubscriptionType = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetTag(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.Tag = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetResourceGroup(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.ResourceGroup = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetAccountingUnit(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.AccountingUnit = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPayerAccount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PayerAccount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetOwnerID(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.OwnerID = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetRegion(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.Region = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetCurrency(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.Currency = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPaymentCurrency(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PaymentCurrency = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetOrderType(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.OrderType = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPayTime(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PayTime = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPretaxGrossAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPricingDiscount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PricingDiscount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetDeductedByCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.DeductedByCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPretaxAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PretaxAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPretaxAmountLocal(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetDeductedByCashCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetDeductedByPrepaidCard(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPaymentAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PaymentAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapPretaxGrossAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapPretaxGrossAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapPricingDiscount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapPricingDiscount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapDeductedByCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapDeductedByCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapPretaxAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapPretaxAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapPretaxAmountLocal(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapPretaxAmountLocal = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapDeductedByCashCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapDeductedByCashCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapDeductedByPrepaidCard(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapDeductedByPrepaidCard = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapPaymentAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapPaymentAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapPretaxGrossAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapPretaxGrossAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapPricingDiscount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapPricingDiscount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapDeductedByCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapDeductedByCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapPretaxAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapPretaxAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapPretaxAmountLocal(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapPretaxAmountLocal = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapDeductedByCashCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapDeductedByCashCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapDeductedByPrepaidCard(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapDeductedByPrepaidCard = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapPaymentAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapPaymentAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedPaymentAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedPaymentAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUsageStartDate(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UsageStartDate = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUsageEndDate(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UsageEndDate = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetBillType(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.BillType = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetOrderId(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.OrderId = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetSubOrderId(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.SubOrderId = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedPretaxGrossAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedPretaxGrossAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedPricingDiscount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedPricingDiscount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedDeductedByCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedDeductedByCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedPretaxAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedPretaxAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedPretaxAmountLocal(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedPretaxAmountLocal = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedDeductedByCashCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedDeductedByCashCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedDeductedByPrepaidCard(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedDeductedByPrepaidCard = &v
	return s
}

type QueryInstanceGaapCostResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryInstanceGaapCostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryInstanceGaapCostResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceGaapCostResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceGaapCostResponse) SetHeaders(v map[string]*string) *QueryInstanceGaapCostResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceGaapCostResponse) SetBody(v *QueryInstanceGaapCostResponseBody) *QueryInstanceGaapCostResponse {
	s.Body = v
	return s
}

type QueryInvoicingCustomerListRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s QueryInvoicingCustomerListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoicingCustomerListRequest) GoString() string {
	return s.String()
}

func (s *QueryInvoicingCustomerListRequest) SetOwnerId(v int64) *QueryInvoicingCustomerListRequest {
	s.OwnerId = &v
	return s
}

type QueryInvoicingCustomerListResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryInvoicingCustomerListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryInvoicingCustomerListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoicingCustomerListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInvoicingCustomerListResponseBody) SetRequestId(v string) *QueryInvoicingCustomerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBody) SetSuccess(v bool) *QueryInvoicingCustomerListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBody) SetCode(v string) *QueryInvoicingCustomerListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBody) SetMessage(v string) *QueryInvoicingCustomerListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBody) SetData(v *QueryInvoicingCustomerListResponseBodyData) *QueryInvoicingCustomerListResponseBody {
	s.Data = v
	return s
}

type QueryInvoicingCustomerListResponseBodyData struct {
	CustomerInvoiceList *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList `json:"CustomerInvoiceList,omitempty" xml:"CustomerInvoiceList,omitempty" type:"Struct"`
}

func (s QueryInvoicingCustomerListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoicingCustomerListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInvoicingCustomerListResponseBodyData) SetCustomerInvoiceList(v *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList) *QueryInvoicingCustomerListResponseBodyData {
	s.CustomerInvoiceList = v
	return s
}

type QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList struct {
	CustomerInvoice []*QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice `json:"CustomerInvoice,omitempty" xml:"CustomerInvoice,omitempty" type:"Repeated"`
}

func (s QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList) GoString() string {
	return s.String()
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList) SetCustomerInvoice(v []*QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList {
	s.CustomerInvoice = v
	return s
}

type QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice struct {
	Id                      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	UserId                  *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserNick                *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
	InvoiceTitle            *string `json:"InvoiceTitle,omitempty" xml:"InvoiceTitle,omitempty"`
	CustomerType            *int64  `json:"CustomerType,omitempty" xml:"CustomerType,omitempty"`
	TaxpayerType            *int64  `json:"TaxpayerType,omitempty" xml:"TaxpayerType,omitempty"`
	Bank                    *string `json:"Bank,omitempty" xml:"Bank,omitempty"`
	BankNo                  *string `json:"BankNo,omitempty" xml:"BankNo,omitempty"`
	OperatingLicenseAddress *string `json:"OperatingLicenseAddress,omitempty" xml:"OperatingLicenseAddress,omitempty"`
	OperatingLicensePhone   *string `json:"OperatingLicensePhone,omitempty" xml:"OperatingLicensePhone,omitempty"`
	RegisterNo              *string `json:"RegisterNo,omitempty" xml:"RegisterNo,omitempty"`
	StartCycle              *int64  `json:"StartCycle,omitempty" xml:"StartCycle,omitempty"`
	Status                  *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	GmtCreate               *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	TaxationLicense         *string `json:"TaxationLicense,omitempty" xml:"TaxationLicense,omitempty"`
	AdjustType              *int64  `json:"AdjustType,omitempty" xml:"AdjustType,omitempty"`
	EndCycle                *int64  `json:"EndCycle,omitempty" xml:"EndCycle,omitempty"`
	TitleChangeInstructions *string `json:"TitleChangeInstructions,omitempty" xml:"TitleChangeInstructions,omitempty"`
	IssueType               *int64  `json:"IssueType,omitempty" xml:"IssueType,omitempty"`
	Type                    *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
	DefaultRemark           *string `json:"DefaultRemark,omitempty" xml:"DefaultRemark,omitempty"`
}

func (s QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GoString() string {
	return s.String()
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetId(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.Id = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetUserId(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.UserId = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetUserNick(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.UserNick = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetInvoiceTitle(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.InvoiceTitle = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetCustomerType(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.CustomerType = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetTaxpayerType(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.TaxpayerType = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetBank(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.Bank = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetBankNo(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.BankNo = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetOperatingLicenseAddress(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.OperatingLicenseAddress = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetOperatingLicensePhone(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.OperatingLicensePhone = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetRegisterNo(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.RegisterNo = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetStartCycle(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.StartCycle = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetStatus(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.Status = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetGmtCreate(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.GmtCreate = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetTaxationLicense(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.TaxationLicense = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetAdjustType(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.AdjustType = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetEndCycle(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.EndCycle = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetTitleChangeInstructions(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.TitleChangeInstructions = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetIssueType(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.IssueType = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetType(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.Type = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetDefaultRemark(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.DefaultRemark = &v
	return s
}

type QueryInvoicingCustomerListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryInvoicingCustomerListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryInvoicingCustomerListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInvoicingCustomerListResponse) GoString() string {
	return s.String()
}

func (s *QueryInvoicingCustomerListResponse) SetHeaders(v map[string]*string) *QueryInvoicingCustomerListResponse {
	s.Headers = v
	return s
}

func (s *QueryInvoicingCustomerListResponse) SetBody(v *QueryInvoicingCustomerListResponseBody) *QueryInvoicingCustomerListResponse {
	s.Body = v
	return s
}

type QueryMonthlyBillRequest struct {
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
}

func (s QueryMonthlyBillRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyBillRequest) GoString() string {
	return s.String()
}

func (s *QueryMonthlyBillRequest) SetBillingCycle(v string) *QueryMonthlyBillRequest {
	s.BillingCycle = &v
	return s
}

type QueryMonthlyBillResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryMonthlyBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryMonthlyBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyBillResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMonthlyBillResponseBody) SetRequestId(v string) *QueryMonthlyBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMonthlyBillResponseBody) SetSuccess(v bool) *QueryMonthlyBillResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMonthlyBillResponseBody) SetCode(v string) *QueryMonthlyBillResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMonthlyBillResponseBody) SetMessage(v string) *QueryMonthlyBillResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMonthlyBillResponseBody) SetData(v *QueryMonthlyBillResponseBodyData) *QueryMonthlyBillResponseBody {
	s.Data = v
	return s
}

type QueryMonthlyBillResponseBodyData struct {
	OutstandingAmount      *float32                               `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	TotalOutstandingAmount *float32                               `json:"TotalOutstandingAmount,omitempty" xml:"TotalOutstandingAmount,omitempty"`
	NewInvoiceAmount       *float32                               `json:"NewInvoiceAmount,omitempty" xml:"NewInvoiceAmount,omitempty"`
	BillingCycle           *string                                `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Items                  *QueryMonthlyBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s QueryMonthlyBillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMonthlyBillResponseBodyData) SetOutstandingAmount(v float32) *QueryMonthlyBillResponseBodyData {
	s.OutstandingAmount = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyData) SetTotalOutstandingAmount(v float32) *QueryMonthlyBillResponseBodyData {
	s.TotalOutstandingAmount = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyData) SetNewInvoiceAmount(v float32) *QueryMonthlyBillResponseBodyData {
	s.NewInvoiceAmount = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyData) SetBillingCycle(v string) *QueryMonthlyBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyData) SetItems(v *QueryMonthlyBillResponseBodyDataItems) *QueryMonthlyBillResponseBodyData {
	s.Items = v
	return s
}

type QueryMonthlyBillResponseBodyDataItems struct {
	Item []*QueryMonthlyBillResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryMonthlyBillResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryMonthlyBillResponseBodyDataItems) SetItem(v []*QueryMonthlyBillResponseBodyDataItemsItem) *QueryMonthlyBillResponseBodyDataItems {
	s.Item = v
	return s
}

type QueryMonthlyBillResponseBodyDataItemsItem struct {
	Item                  *string  `json:"Item,omitempty" xml:"Item,omitempty"`
	ProductCode           *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType           *string  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType      *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	SolutionCode          *string  `json:"SolutionCode,omitempty" xml:"SolutionCode,omitempty"`
	SolutionName          *string  `json:"SolutionName,omitempty" xml:"SolutionName,omitempty"`
	PretaxGrossAmount     *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	InvoiceDiscount       *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	DeductedByCoupons     *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	PretaxAmount          *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	Currency              *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	PretaxAmountLocal     *float32 `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	Tax                   *float32 `json:"Tax,omitempty" xml:"Tax,omitempty"`
	AfterTaxAmount        *float32 `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	OutstandingAmount     *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	PaymentAmount         *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	PaymentCurrency       *string  `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
}

func (s QueryMonthlyBillResponseBodyDataItemsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyBillResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetItem(v string) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.Item = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetProductCode(v string) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetProductType(v string) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetSubscriptionType(v string) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetSolutionCode(v string) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.SolutionCode = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetSolutionName(v string) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.SolutionName = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetCurrency(v string) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetPretaxAmountLocal(v float32) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetTax(v float32) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.Tax = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetAfterTaxAmount(v float32) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.AfterTaxAmount = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QueryMonthlyBillResponseBodyDataItemsItem) SetPaymentCurrency(v string) *QueryMonthlyBillResponseBodyDataItemsItem {
	s.PaymentCurrency = &v
	return s
}

type QueryMonthlyBillResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMonthlyBillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMonthlyBillResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyBillResponse) GoString() string {
	return s.String()
}

func (s *QueryMonthlyBillResponse) SetHeaders(v map[string]*string) *QueryMonthlyBillResponse {
	s.Headers = v
	return s
}

func (s *QueryMonthlyBillResponse) SetBody(v *QueryMonthlyBillResponseBody) *QueryMonthlyBillResponse {
	s.Body = v
	return s
}

type QueryMonthlyInstanceConsumptionRequest struct {
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	BillingCycle     *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s QueryMonthlyInstanceConsumptionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyInstanceConsumptionRequest) GoString() string {
	return s.String()
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetProductCode(v string) *QueryMonthlyInstanceConsumptionRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetOwnerId(v int64) *QueryMonthlyInstanceConsumptionRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetPageNum(v int32) *QueryMonthlyInstanceConsumptionRequest {
	s.PageNum = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetPageSize(v int32) *QueryMonthlyInstanceConsumptionRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetBillingCycle(v string) *QueryMonthlyInstanceConsumptionRequest {
	s.BillingCycle = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetProductType(v string) *QueryMonthlyInstanceConsumptionRequest {
	s.ProductType = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionRequest) SetSubscriptionType(v string) *QueryMonthlyInstanceConsumptionRequest {
	s.SubscriptionType = &v
	return s
}

type QueryMonthlyInstanceConsumptionResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryMonthlyInstanceConsumptionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryMonthlyInstanceConsumptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyInstanceConsumptionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMonthlyInstanceConsumptionResponseBody) SetRequestId(v string) *QueryMonthlyInstanceConsumptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBody) SetSuccess(v bool) *QueryMonthlyInstanceConsumptionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBody) SetCode(v string) *QueryMonthlyInstanceConsumptionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBody) SetMessage(v string) *QueryMonthlyInstanceConsumptionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBody) SetData(v *QueryMonthlyInstanceConsumptionResponseBodyData) *QueryMonthlyInstanceConsumptionResponseBody {
	s.Data = v
	return s
}

type QueryMonthlyInstanceConsumptionResponseBodyData struct {
	PageNum      *int32                                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount   *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	BillingCycle *string                                               `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Items        *QueryMonthlyInstanceConsumptionResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s QueryMonthlyInstanceConsumptionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyInstanceConsumptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyData) SetPageNum(v int32) *QueryMonthlyInstanceConsumptionResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyData) SetPageSize(v int32) *QueryMonthlyInstanceConsumptionResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyData) SetTotalCount(v int32) *QueryMonthlyInstanceConsumptionResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyData) SetBillingCycle(v string) *QueryMonthlyInstanceConsumptionResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyData) SetItems(v *QueryMonthlyInstanceConsumptionResponseBodyDataItems) *QueryMonthlyInstanceConsumptionResponseBodyData {
	s.Items = v
	return s
}

type QueryMonthlyInstanceConsumptionResponseBodyDataItems struct {
	Item []*QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QueryMonthlyInstanceConsumptionResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyInstanceConsumptionResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItems) SetItem(v []*QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) *QueryMonthlyInstanceConsumptionResponseBodyDataItems {
	s.Item = v
	return s
}

type QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem struct {
	InstanceID        *string  `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	ProductCode       *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType       *string  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType  *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	Tag               *string  `json:"Tag,omitempty" xml:"Tag,omitempty"`
	ResourceGroup     *string  `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	PayerAccount      *string  `json:"PayerAccount,omitempty" xml:"PayerAccount,omitempty"`
	OwnerID           *string  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	Region            *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	PretaxGrossAmount *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	DiscountAmount    *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	PretaxAmount      *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	Currency          *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	PretaxAmountLocal *float32 `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	Tax               *float32 `json:"Tax,omitempty" xml:"Tax,omitempty"`
	AfterTaxAmount    *float32 `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	PaymentCurrency   *string  `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
}

func (s QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetInstanceID(v string) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.InstanceID = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetProductCode(v string) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetProductType(v string) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetSubscriptionType(v string) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetTag(v string) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.Tag = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetResourceGroup(v string) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.ResourceGroup = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetPayerAccount(v string) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.PayerAccount = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetOwnerID(v string) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetRegion(v string) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.Region = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetDiscountAmount(v float32) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.DiscountAmount = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetCurrency(v string) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetPretaxAmountLocal(v float32) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetTax(v float32) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.Tax = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetAfterTaxAmount(v float32) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.AfterTaxAmount = &v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem) SetPaymentCurrency(v string) *QueryMonthlyInstanceConsumptionResponseBodyDataItemsItem {
	s.PaymentCurrency = &v
	return s
}

type QueryMonthlyInstanceConsumptionResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryMonthlyInstanceConsumptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMonthlyInstanceConsumptionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMonthlyInstanceConsumptionResponse) GoString() string {
	return s.String()
}

func (s *QueryMonthlyInstanceConsumptionResponse) SetHeaders(v map[string]*string) *QueryMonthlyInstanceConsumptionResponse {
	s.Headers = v
	return s
}

func (s *QueryMonthlyInstanceConsumptionResponse) SetBody(v *QueryMonthlyInstanceConsumptionResponseBody) *QueryMonthlyInstanceConsumptionResponse {
	s.Body = v
	return s
}

type QueryOrdersRequest struct {
	CreateTimeEnd    *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	OrderType        *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PaymentStatus    *string `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	CreateTimeStart  *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s QueryOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrdersRequest) GoString() string {
	return s.String()
}

func (s *QueryOrdersRequest) SetCreateTimeEnd(v string) *QueryOrdersRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *QueryOrdersRequest) SetPageNum(v int32) *QueryOrdersRequest {
	s.PageNum = &v
	return s
}

func (s *QueryOrdersRequest) SetPageSize(v int32) *QueryOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *QueryOrdersRequest) SetProductCode(v string) *QueryOrdersRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryOrdersRequest) SetProductType(v string) *QueryOrdersRequest {
	s.ProductType = &v
	return s
}

func (s *QueryOrdersRequest) SetSubscriptionType(v string) *QueryOrdersRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QueryOrdersRequest) SetOrderType(v string) *QueryOrdersRequest {
	s.OrderType = &v
	return s
}

func (s *QueryOrdersRequest) SetPaymentStatus(v string) *QueryOrdersRequest {
	s.PaymentStatus = &v
	return s
}

func (s *QueryOrdersRequest) SetCreateTimeStart(v string) *QueryOrdersRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *QueryOrdersRequest) SetOwnerId(v int64) *QueryOrdersRequest {
	s.OwnerId = &v
	return s
}

type QueryOrdersResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrdersResponseBody) SetRequestId(v string) *QueryOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrdersResponseBody) SetSuccess(v bool) *QueryOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *QueryOrdersResponseBody) SetCode(v string) *QueryOrdersResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOrdersResponseBody) SetMessage(v string) *QueryOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOrdersResponseBody) SetData(v *QueryOrdersResponseBodyData) *QueryOrdersResponseBody {
	s.Data = v
	return s
}

type QueryOrdersResponseBodyData struct {
	HostName   *string                               `json:"HostName,omitempty" xml:"HostName,omitempty"`
	PageNum    *int32                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	OrderList  *QueryOrdersResponseBodyDataOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Struct"`
}

func (s QueryOrdersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOrdersResponseBodyData) SetHostName(v string) *QueryOrdersResponseBodyData {
	s.HostName = &v
	return s
}

func (s *QueryOrdersResponseBodyData) SetPageNum(v int32) *QueryOrdersResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryOrdersResponseBodyData) SetPageSize(v int32) *QueryOrdersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryOrdersResponseBodyData) SetTotalCount(v int32) *QueryOrdersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryOrdersResponseBodyData) SetOrderList(v *QueryOrdersResponseBodyDataOrderList) *QueryOrdersResponseBodyData {
	s.OrderList = v
	return s
}

type QueryOrdersResponseBodyDataOrderList struct {
	Order []*QueryOrdersResponseBodyDataOrderListOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Repeated"`
}

func (s QueryOrdersResponseBodyDataOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryOrdersResponseBodyDataOrderList) GoString() string {
	return s.String()
}

func (s *QueryOrdersResponseBodyDataOrderList) SetOrder(v []*QueryOrdersResponseBodyDataOrderListOrder) *QueryOrdersResponseBodyDataOrderList {
	s.Order = v
	return s
}

type QueryOrdersResponseBodyDataOrderListOrder struct {
	OrderId           *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType       *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType  *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	OrderType         *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	CreateTime        *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PaymentTime       *string `json:"PaymentTime,omitempty" xml:"PaymentTime,omitempty"`
	PaymentStatus     *string `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	PretaxGrossAmount *string `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	PretaxAmount      *string `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	Currency          *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	PretaxAmountLocal *string `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	Tax               *string `json:"Tax,omitempty" xml:"Tax,omitempty"`
	AfterTaxAmount    *string `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	PaymentCurrency   *string `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	RelatedOrderId    *string `json:"RelatedOrderId,omitempty" xml:"RelatedOrderId,omitempty"`
	CommodityCode     *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s QueryOrdersResponseBodyDataOrderListOrder) String() string {
	return tea.Prettify(s)
}

func (s QueryOrdersResponseBodyDataOrderListOrder) GoString() string {
	return s.String()
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetOrderId(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.OrderId = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetProductCode(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.ProductCode = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetProductType(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.ProductType = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetSubscriptionType(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.SubscriptionType = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetOrderType(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.OrderType = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetCreateTime(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.CreateTime = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetPaymentTime(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.PaymentTime = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetPaymentStatus(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.PaymentStatus = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetPretaxGrossAmount(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetPretaxAmount(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.PretaxAmount = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetCurrency(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.Currency = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetPretaxAmountLocal(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetTax(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.Tax = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetAfterTaxAmount(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.AfterTaxAmount = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetPaymentCurrency(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.PaymentCurrency = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetRelatedOrderId(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.RelatedOrderId = &v
	return s
}

func (s *QueryOrdersResponseBodyDataOrderListOrder) SetCommodityCode(v string) *QueryOrdersResponseBodyDataOrderListOrder {
	s.CommodityCode = &v
	return s
}

type QueryOrdersResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *QueryOrdersResponse) SetBody(v *QueryOrdersResponseBody) *QueryOrdersResponse {
	s.Body = v
	return s
}

type QueryPermissionListRequest struct {
	RelationId *int64 `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
}

func (s QueryPermissionListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionListRequest) GoString() string {
	return s.String()
}

func (s *QueryPermissionListRequest) SetRelationId(v int64) *QueryPermissionListRequest {
	s.RelationId = &v
	return s
}

type QueryPermissionListResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryPermissionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryPermissionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPermissionListResponseBody) SetCode(v string) *QueryPermissionListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPermissionListResponseBody) SetRequestId(v string) *QueryPermissionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPermissionListResponseBody) SetSuccess(v bool) *QueryPermissionListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPermissionListResponseBody) SetMessage(v string) *QueryPermissionListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPermissionListResponseBody) SetData(v *QueryPermissionListResponseBodyData) *QueryPermissionListResponseBody {
	s.Data = v
	return s
}

type QueryPermissionListResponseBodyData struct {
	MasterId       *int64                                               `json:"MasterId,omitempty" xml:"MasterId,omitempty"`
	MemberId       *int64                                               `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	RelationType   *string                                              `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	State          *string                                              `json:"State,omitempty" xml:"State,omitempty"`
	SetupTime      *string                                              `json:"SetupTime,omitempty" xml:"SetupTime,omitempty"`
	StartTime      *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PermissionList []*QueryPermissionListResponseBodyDataPermissionList `json:"PermissionList,omitempty" xml:"PermissionList,omitempty" type:"Repeated"`
}

func (s QueryPermissionListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPermissionListResponseBodyData) SetMasterId(v int64) *QueryPermissionListResponseBodyData {
	s.MasterId = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetMemberId(v int64) *QueryPermissionListResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetRelationType(v string) *QueryPermissionListResponseBodyData {
	s.RelationType = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetState(v string) *QueryPermissionListResponseBodyData {
	s.State = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetSetupTime(v string) *QueryPermissionListResponseBodyData {
	s.SetupTime = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetStartTime(v string) *QueryPermissionListResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetEndTime(v string) *QueryPermissionListResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryPermissionListResponseBodyData) SetPermissionList(v []*QueryPermissionListResponseBodyDataPermissionList) *QueryPermissionListResponseBodyData {
	s.PermissionList = v
	return s
}

type QueryPermissionListResponseBodyDataPermissionList struct {
	PermissionCode *string `json:"PermissionCode,omitempty" xml:"PermissionCode,omitempty"`
	PermissionName *string `json:"PermissionName,omitempty" xml:"PermissionName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s QueryPermissionListResponseBodyDataPermissionList) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionListResponseBodyDataPermissionList) GoString() string {
	return s.String()
}

func (s *QueryPermissionListResponseBodyDataPermissionList) SetPermissionCode(v string) *QueryPermissionListResponseBodyDataPermissionList {
	s.PermissionCode = &v
	return s
}

func (s *QueryPermissionListResponseBodyDataPermissionList) SetPermissionName(v string) *QueryPermissionListResponseBodyDataPermissionList {
	s.PermissionName = &v
	return s
}

func (s *QueryPermissionListResponseBodyDataPermissionList) SetStartTime(v string) *QueryPermissionListResponseBodyDataPermissionList {
	s.StartTime = &v
	return s
}

func (s *QueryPermissionListResponseBodyDataPermissionList) SetEndTime(v string) *QueryPermissionListResponseBodyDataPermissionList {
	s.EndTime = &v
	return s
}

type QueryPermissionListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPermissionListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPermissionListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPermissionListResponse) GoString() string {
	return s.String()
}

func (s *QueryPermissionListResponse) SetHeaders(v map[string]*string) *QueryPermissionListResponse {
	s.Headers = v
	return s
}

func (s *QueryPermissionListResponse) SetBody(v *QueryPermissionListResponseBody) *QueryPermissionListResponse {
	s.Body = v
	return s
}

type QueryPrepaidCardsRequest struct {
	ExpiryTimeEnd   *string `json:"ExpiryTimeEnd,omitempty" xml:"ExpiryTimeEnd,omitempty"`
	ExpiryTimeStart *string `json:"ExpiryTimeStart,omitempty" xml:"ExpiryTimeStart,omitempty"`
	EffectiveOrNot  *bool   `json:"EffectiveOrNot,omitempty" xml:"EffectiveOrNot,omitempty"`
}

func (s QueryPrepaidCardsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPrepaidCardsRequest) GoString() string {
	return s.String()
}

func (s *QueryPrepaidCardsRequest) SetExpiryTimeEnd(v string) *QueryPrepaidCardsRequest {
	s.ExpiryTimeEnd = &v
	return s
}

func (s *QueryPrepaidCardsRequest) SetExpiryTimeStart(v string) *QueryPrepaidCardsRequest {
	s.ExpiryTimeStart = &v
	return s
}

func (s *QueryPrepaidCardsRequest) SetEffectiveOrNot(v bool) *QueryPrepaidCardsRequest {
	s.EffectiveOrNot = &v
	return s
}

type QueryPrepaidCardsResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryPrepaidCardsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryPrepaidCardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPrepaidCardsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPrepaidCardsResponseBody) SetRequestId(v string) *QueryPrepaidCardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPrepaidCardsResponseBody) SetSuccess(v bool) *QueryPrepaidCardsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPrepaidCardsResponseBody) SetCode(v string) *QueryPrepaidCardsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPrepaidCardsResponseBody) SetMessage(v string) *QueryPrepaidCardsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPrepaidCardsResponseBody) SetData(v *QueryPrepaidCardsResponseBodyData) *QueryPrepaidCardsResponseBody {
	s.Data = v
	return s
}

type QueryPrepaidCardsResponseBodyData struct {
	PrepaidCard []*QueryPrepaidCardsResponseBodyDataPrepaidCard `json:"PrepaidCard,omitempty" xml:"PrepaidCard,omitempty" type:"Repeated"`
}

func (s QueryPrepaidCardsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPrepaidCardsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPrepaidCardsResponseBodyData) SetPrepaidCard(v []*QueryPrepaidCardsResponseBodyDataPrepaidCard) *QueryPrepaidCardsResponseBodyData {
	s.PrepaidCard = v
	return s
}

type QueryPrepaidCardsResponseBodyDataPrepaidCard struct {
	PrepaidCardId       *int64  `json:"PrepaidCardId,omitempty" xml:"PrepaidCardId,omitempty"`
	PrepaidCardNo       *string `json:"PrepaidCardNo,omitempty" xml:"PrepaidCardNo,omitempty"`
	GrantedTime         *string `json:"GrantedTime,omitempty" xml:"GrantedTime,omitempty"`
	EffectiveTime       *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ExpiryTime          *string `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	ApplicableProducts  *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	ApplicableScenarios *string `json:"ApplicableScenarios,omitempty" xml:"ApplicableScenarios,omitempty"`
	NominalValue        *string `json:"NominalValue,omitempty" xml:"NominalValue,omitempty"`
	Balance             *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryPrepaidCardsResponseBodyDataPrepaidCard) String() string {
	return tea.Prettify(s)
}

func (s QueryPrepaidCardsResponseBodyDataPrepaidCard) GoString() string {
	return s.String()
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetPrepaidCardId(v int64) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.PrepaidCardId = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetPrepaidCardNo(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.PrepaidCardNo = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetGrantedTime(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.GrantedTime = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetEffectiveTime(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.EffectiveTime = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetExpiryTime(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.ExpiryTime = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetApplicableProducts(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.ApplicableProducts = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetApplicableScenarios(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.ApplicableScenarios = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetNominalValue(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.NominalValue = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetBalance(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.Balance = &v
	return s
}

func (s *QueryPrepaidCardsResponseBodyDataPrepaidCard) SetStatus(v string) *QueryPrepaidCardsResponseBodyDataPrepaidCard {
	s.Status = &v
	return s
}

type QueryPrepaidCardsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryPrepaidCardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPrepaidCardsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPrepaidCardsResponse) GoString() string {
	return s.String()
}

func (s *QueryPrepaidCardsResponse) SetHeaders(v map[string]*string) *QueryPrepaidCardsResponse {
	s.Headers = v
	return s
}

func (s *QueryPrepaidCardsResponse) SetBody(v *QueryPrepaidCardsResponseBody) *QueryPrepaidCardsResponse {
	s.Body = v
	return s
}

type QueryProductListRequest struct {
	QueryTotalCount *bool  `json:"QueryTotalCount,omitempty" xml:"QueryTotalCount,omitempty"`
	PageNum         *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize        *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryProductListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProductListRequest) GoString() string {
	return s.String()
}

func (s *QueryProductListRequest) SetQueryTotalCount(v bool) *QueryProductListRequest {
	s.QueryTotalCount = &v
	return s
}

func (s *QueryProductListRequest) SetPageNum(v int32) *QueryProductListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryProductListRequest) SetPageSize(v int32) *QueryProductListRequest {
	s.PageSize = &v
	return s
}

type QueryProductListResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryProductListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryProductListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProductListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProductListResponseBody) SetRequestId(v string) *QueryProductListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProductListResponseBody) SetSuccess(v bool) *QueryProductListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryProductListResponseBody) SetCode(v string) *QueryProductListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryProductListResponseBody) SetMessage(v string) *QueryProductListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryProductListResponseBody) SetData(v *QueryProductListResponseBodyData) *QueryProductListResponseBody {
	s.Data = v
	return s
}

type QueryProductListResponseBodyData struct {
	TotalCount  *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageNum     *int32                                       `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductList *QueryProductListResponseBodyDataProductList `json:"ProductList,omitempty" xml:"ProductList,omitempty" type:"Struct"`
}

func (s QueryProductListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryProductListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProductListResponseBodyData) SetTotalCount(v int32) *QueryProductListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryProductListResponseBodyData) SetPageNum(v int32) *QueryProductListResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryProductListResponseBodyData) SetPageSize(v int32) *QueryProductListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryProductListResponseBodyData) SetProductList(v *QueryProductListResponseBodyDataProductList) *QueryProductListResponseBodyData {
	s.ProductList = v
	return s
}

type QueryProductListResponseBodyDataProductList struct {
	Product []*QueryProductListResponseBodyDataProductListProduct `json:"Product,omitempty" xml:"Product,omitempty" type:"Repeated"`
}

func (s QueryProductListResponseBodyDataProductList) String() string {
	return tea.Prettify(s)
}

func (s QueryProductListResponseBodyDataProductList) GoString() string {
	return s.String()
}

func (s *QueryProductListResponseBodyDataProductList) SetProduct(v []*QueryProductListResponseBodyDataProductListProduct) *QueryProductListResponseBodyDataProductList {
	s.Product = v
	return s
}

type QueryProductListResponseBodyDataProductListProduct struct {
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName      *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s QueryProductListResponseBodyDataProductListProduct) String() string {
	return tea.Prettify(s)
}

func (s QueryProductListResponseBodyDataProductListProduct) GoString() string {
	return s.String()
}

func (s *QueryProductListResponseBodyDataProductListProduct) SetProductCode(v string) *QueryProductListResponseBodyDataProductListProduct {
	s.ProductCode = &v
	return s
}

func (s *QueryProductListResponseBodyDataProductListProduct) SetProductName(v string) *QueryProductListResponseBodyDataProductListProduct {
	s.ProductName = &v
	return s
}

func (s *QueryProductListResponseBodyDataProductListProduct) SetProductType(v string) *QueryProductListResponseBodyDataProductListProduct {
	s.ProductType = &v
	return s
}

func (s *QueryProductListResponseBodyDataProductListProduct) SetSubscriptionType(v string) *QueryProductListResponseBodyDataProductListProduct {
	s.SubscriptionType = &v
	return s
}

type QueryProductListResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryProductListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryProductListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProductListResponse) GoString() string {
	return s.String()
}

func (s *QueryProductListResponse) SetHeaders(v map[string]*string) *QueryProductListResponse {
	s.Headers = v
	return s
}

func (s *QueryProductListResponse) SetBody(v *QueryProductListResponseBody) *QueryProductListResponse {
	s.Body = v
	return s
}

type QueryRedeemRequest struct {
	ExpiryTimeStart *string `json:"ExpiryTimeStart,omitempty" xml:"ExpiryTimeStart,omitempty"`
	ExpiryTimeEnd   *string `json:"ExpiryTimeEnd,omitempty" xml:"ExpiryTimeEnd,omitempty"`
	EffectiveOrNot  *bool   `json:"EffectiveOrNot,omitempty" xml:"EffectiveOrNot,omitempty"`
	PageNum         *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryRedeemRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRedeemRequest) GoString() string {
	return s.String()
}

func (s *QueryRedeemRequest) SetExpiryTimeStart(v string) *QueryRedeemRequest {
	s.ExpiryTimeStart = &v
	return s
}

func (s *QueryRedeemRequest) SetExpiryTimeEnd(v string) *QueryRedeemRequest {
	s.ExpiryTimeEnd = &v
	return s
}

func (s *QueryRedeemRequest) SetEffectiveOrNot(v bool) *QueryRedeemRequest {
	s.EffectiveOrNot = &v
	return s
}

func (s *QueryRedeemRequest) SetPageNum(v int32) *QueryRedeemRequest {
	s.PageNum = &v
	return s
}

func (s *QueryRedeemRequest) SetPageSize(v int32) *QueryRedeemRequest {
	s.PageSize = &v
	return s
}

type QueryRedeemResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryRedeemResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryRedeemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRedeemResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRedeemResponseBody) SetRequestId(v string) *QueryRedeemResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRedeemResponseBody) SetSuccess(v bool) *QueryRedeemResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRedeemResponseBody) SetCode(v string) *QueryRedeemResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRedeemResponseBody) SetMessage(v string) *QueryRedeemResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRedeemResponseBody) SetData(v *QueryRedeemResponseBodyData) *QueryRedeemResponseBody {
	s.Data = v
	return s
}

type QueryRedeemResponseBodyData struct {
	PageNum    *int64                             `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Redeem     *QueryRedeemResponseBodyDataRedeem `json:"Redeem,omitempty" xml:"Redeem,omitempty" type:"Struct"`
}

func (s QueryRedeemResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRedeemResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRedeemResponseBodyData) SetPageNum(v int64) *QueryRedeemResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryRedeemResponseBodyData) SetPageSize(v int64) *QueryRedeemResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRedeemResponseBodyData) SetTotalCount(v int64) *QueryRedeemResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryRedeemResponseBodyData) SetRedeem(v *QueryRedeemResponseBodyDataRedeem) *QueryRedeemResponseBodyData {
	s.Redeem = v
	return s
}

type QueryRedeemResponseBodyDataRedeem struct {
	Redeem []*QueryRedeemResponseBodyDataRedeemRedeem `json:"Redeem,omitempty" xml:"Redeem,omitempty" type:"Repeated"`
}

func (s QueryRedeemResponseBodyDataRedeem) String() string {
	return tea.Prettify(s)
}

func (s QueryRedeemResponseBodyDataRedeem) GoString() string {
	return s.String()
}

func (s *QueryRedeemResponseBodyDataRedeem) SetRedeem(v []*QueryRedeemResponseBodyDataRedeemRedeem) *QueryRedeemResponseBodyDataRedeem {
	s.Redeem = v
	return s
}

type QueryRedeemResponseBodyDataRedeemRedeem struct {
	RedeemId           *string `json:"RedeemId,omitempty" xml:"RedeemId,omitempty"`
	RedeemNo           *string `json:"RedeemNo,omitempty" xml:"RedeemNo,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	GrantedTime        *string `json:"GrantedTime,omitempty" xml:"GrantedTime,omitempty"`
	EffectiveTime      *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ExpiryTime         *string `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	NominalValue       *string `json:"NominalValue,omitempty" xml:"NominalValue,omitempty"`
	Balance            *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	ApplicableProducts *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	Specification      *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s QueryRedeemResponseBodyDataRedeemRedeem) String() string {
	return tea.Prettify(s)
}

func (s QueryRedeemResponseBodyDataRedeemRedeem) GoString() string {
	return s.String()
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetRedeemId(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.RedeemId = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetRedeemNo(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.RedeemNo = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetStatus(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.Status = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetGrantedTime(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.GrantedTime = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetEffectiveTime(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.EffectiveTime = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetExpiryTime(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.ExpiryTime = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetNominalValue(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.NominalValue = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetBalance(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.Balance = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetApplicableProducts(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.ApplicableProducts = &v
	return s
}

func (s *QueryRedeemResponseBodyDataRedeemRedeem) SetSpecification(v string) *QueryRedeemResponseBodyDataRedeemRedeem {
	s.Specification = &v
	return s
}

type QueryRedeemResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRedeemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRedeemResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRedeemResponse) GoString() string {
	return s.String()
}

func (s *QueryRedeemResponse) SetHeaders(v map[string]*string) *QueryRedeemResponse {
	s.Headers = v
	return s
}

func (s *QueryRedeemResponse) SetBody(v *QueryRedeemResponseBody) *QueryRedeemResponse {
	s.Body = v
	return s
}

type QueryRelationListRequest struct {
	UserId     *int64    `json:"UserId,omitempty" xml:"UserId,omitempty"`
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	PageNum    *int32    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryRelationListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationListRequest) GoString() string {
	return s.String()
}

func (s *QueryRelationListRequest) SetUserId(v int64) *QueryRelationListRequest {
	s.UserId = &v
	return s
}

func (s *QueryRelationListRequest) SetStatusList(v []*string) *QueryRelationListRequest {
	s.StatusList = v
	return s
}

func (s *QueryRelationListRequest) SetPageNum(v int32) *QueryRelationListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryRelationListRequest) SetPageSize(v int32) *QueryRelationListRequest {
	s.PageSize = &v
	return s
}

type QueryRelationListResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryRelationListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryRelationListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRelationListResponseBody) SetCode(v string) *QueryRelationListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRelationListResponseBody) SetRequestId(v string) *QueryRelationListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRelationListResponseBody) SetSuccess(v bool) *QueryRelationListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRelationListResponseBody) SetMessage(v string) *QueryRelationListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRelationListResponseBody) SetData(v *QueryRelationListResponseBodyData) *QueryRelationListResponseBody {
	s.Data = v
	return s
}

type QueryRelationListResponseBodyData struct {
	PageNum                   *int32                                                        `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize                  *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount                *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	FinancialRelationInfoList []*QueryRelationListResponseBodyDataFinancialRelationInfoList `json:"FinancialRelationInfoList,omitempty" xml:"FinancialRelationInfoList,omitempty" type:"Repeated"`
}

func (s QueryRelationListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRelationListResponseBodyData) SetPageNum(v int32) *QueryRelationListResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryRelationListResponseBodyData) SetPageSize(v int32) *QueryRelationListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRelationListResponseBodyData) SetTotalCount(v int32) *QueryRelationListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryRelationListResponseBodyData) SetFinancialRelationInfoList(v []*QueryRelationListResponseBodyDataFinancialRelationInfoList) *QueryRelationListResponseBodyData {
	s.FinancialRelationInfoList = v
	return s
}

type QueryRelationListResponseBodyDataFinancialRelationInfoList struct {
	RelationId      *int64  `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	AccountType     *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountId       *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName     *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountNickName *string `json:"AccountNickName,omitempty" xml:"AccountNickName,omitempty"`
	RelationType    *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	SetupTime       *string `json:"SetupTime,omitempty" xml:"SetupTime,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s QueryRelationListResponseBodyDataFinancialRelationInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationListResponseBodyDataFinancialRelationInfoList) GoString() string {
	return s.String()
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetRelationId(v int64) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.RelationId = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetAccountType(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.AccountType = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetAccountId(v int64) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.AccountId = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetAccountName(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.AccountName = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetAccountNickName(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.AccountNickName = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetRelationType(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.RelationType = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetState(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.State = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetSetupTime(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.SetupTime = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetStartTime(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.StartTime = &v
	return s
}

func (s *QueryRelationListResponseBodyDataFinancialRelationInfoList) SetEndTime(v string) *QueryRelationListResponseBodyDataFinancialRelationInfoList {
	s.EndTime = &v
	return s
}

type QueryRelationListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRelationListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRelationListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationListResponse) GoString() string {
	return s.String()
}

func (s *QueryRelationListResponse) SetHeaders(v map[string]*string) *QueryRelationListResponse {
	s.Headers = v
	return s
}

func (s *QueryRelationListResponse) SetBody(v *QueryRelationListResponseBody) *QueryRelationListResponse {
	s.Body = v
	return s
}

type QueryResellerAvailableQuotaRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ItemCodes *string `json:"ItemCodes,omitempty" xml:"ItemCodes,omitempty"`
}

func (s QueryResellerAvailableQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryResellerAvailableQuotaRequest) GoString() string {
	return s.String()
}

func (s *QueryResellerAvailableQuotaRequest) SetOwnerId(v int64) *QueryResellerAvailableQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryResellerAvailableQuotaRequest) SetItemCodes(v string) *QueryResellerAvailableQuotaRequest {
	s.ItemCodes = &v
	return s
}

type QueryResellerAvailableQuotaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s QueryResellerAvailableQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryResellerAvailableQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResellerAvailableQuotaResponseBody) SetRequestId(v string) *QueryResellerAvailableQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryResellerAvailableQuotaResponseBody) SetCode(v string) *QueryResellerAvailableQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *QueryResellerAvailableQuotaResponseBody) SetMessage(v string) *QueryResellerAvailableQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *QueryResellerAvailableQuotaResponseBody) SetSuccess(v bool) *QueryResellerAvailableQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *QueryResellerAvailableQuotaResponseBody) SetData(v string) *QueryResellerAvailableQuotaResponseBody {
	s.Data = &v
	return s
}

type QueryResellerAvailableQuotaResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryResellerAvailableQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryResellerAvailableQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryResellerAvailableQuotaResponse) GoString() string {
	return s.String()
}

func (s *QueryResellerAvailableQuotaResponse) SetHeaders(v map[string]*string) *QueryResellerAvailableQuotaResponse {
	s.Headers = v
	return s
}

func (s *QueryResellerAvailableQuotaResponse) SetBody(v *QueryResellerAvailableQuotaResponseBody) *QueryResellerAvailableQuotaResponse {
	s.Body = v
	return s
}

type QueryResourcePackageInstancesRequest struct {
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductCode     *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ExpiryTimeStart *string `json:"ExpiryTimeStart,omitempty" xml:"ExpiryTimeStart,omitempty"`
	ExpiryTimeEnd   *string `json:"ExpiryTimeEnd,omitempty" xml:"ExpiryTimeEnd,omitempty"`
	PageNum         *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryResourcePackageInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesRequest) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesRequest) SetOwnerId(v int64) *QueryResourcePackageInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryResourcePackageInstancesRequest) SetProductCode(v string) *QueryResourcePackageInstancesRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryResourcePackageInstancesRequest) SetExpiryTimeStart(v string) *QueryResourcePackageInstancesRequest {
	s.ExpiryTimeStart = &v
	return s
}

func (s *QueryResourcePackageInstancesRequest) SetExpiryTimeEnd(v string) *QueryResourcePackageInstancesRequest {
	s.ExpiryTimeEnd = &v
	return s
}

func (s *QueryResourcePackageInstancesRequest) SetPageNum(v int32) *QueryResourcePackageInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *QueryResourcePackageInstancesRequest) SetPageSize(v int32) *QueryResourcePackageInstancesRequest {
	s.PageSize = &v
	return s
}

type QueryResourcePackageInstancesResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Page      *int32                                         `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize  *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total     *int32                                         `json:"Total,omitempty" xml:"Total,omitempty"`
	Data      *QueryResourcePackageInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryResourcePackageInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesResponseBody) SetRequestId(v string) *QueryResourcePackageInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetSuccess(v bool) *QueryResourcePackageInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetCode(v string) *QueryResourcePackageInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetMessage(v string) *QueryResourcePackageInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetPage(v int32) *QueryResourcePackageInstancesResponseBody {
	s.Page = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetPageSize(v int32) *QueryResourcePackageInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetTotal(v int32) *QueryResourcePackageInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBody) SetData(v *QueryResourcePackageInstancesResponseBodyData) *QueryResourcePackageInstancesResponseBody {
	s.Data = v
	return s
}

type QueryResourcePackageInstancesResponseBodyData struct {
	HostId     *string                                                 `json:"HostId,omitempty" xml:"HostId,omitempty"`
	PageNum    *string                                                 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *string                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *string                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Instances  *QueryResourcePackageInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
}

func (s QueryResourcePackageInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesResponseBodyData) SetHostId(v string) *QueryResourcePackageInstancesResponseBodyData {
	s.HostId = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyData) SetPageNum(v string) *QueryResourcePackageInstancesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyData) SetPageSize(v string) *QueryResourcePackageInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyData) SetTotalCount(v string) *QueryResourcePackageInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyData) SetInstances(v *QueryResourcePackageInstancesResponseBodyDataInstances) *QueryResourcePackageInstancesResponseBodyData {
	s.Instances = v
	return s
}

type QueryResourcePackageInstancesResponseBodyDataInstances struct {
	Instance []*QueryResourcePackageInstancesResponseBodyDataInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s QueryResourcePackageInstancesResponseBodyDataInstances) String() string {
	return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstances) SetInstance(v []*QueryResourcePackageInstancesResponseBodyDataInstancesInstance) *QueryResourcePackageInstancesResponseBodyDataInstances {
	s.Instance = v
	return s
}

type QueryResourcePackageInstancesResponseBodyDataInstancesInstance struct {
	InstanceId          *string                                                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Region              *string                                                                           `json:"Region,omitempty" xml:"Region,omitempty"`
	TotalAmount         *string                                                                           `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	TotalAmountUnit     *string                                                                           `json:"TotalAmountUnit,omitempty" xml:"TotalAmountUnit,omitempty"`
	RemainingAmount     *string                                                                           `json:"RemainingAmount,omitempty" xml:"RemainingAmount,omitempty"`
	RemainingAmountUnit *string                                                                           `json:"RemainingAmountUnit,omitempty" xml:"RemainingAmountUnit,omitempty"`
	EffectiveTime       *string                                                                           `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	ExpiryTime          *string                                                                           `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	Remark              *string                                                                           `json:"Remark,omitempty" xml:"Remark,omitempty"`
	PackageType         *string                                                                           `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Status              *string                                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	DeductType          *string                                                                           `json:"DeductType,omitempty" xml:"DeductType,omitempty"`
	ApplicableProducts  *QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty" type:"Struct"`
}

func (s QueryResourcePackageInstancesResponseBodyDataInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseBodyDataInstancesInstance) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetInstanceId(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetRegion(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.Region = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetTotalAmount(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.TotalAmount = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetTotalAmountUnit(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.TotalAmountUnit = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetRemainingAmount(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.RemainingAmount = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetRemainingAmountUnit(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.RemainingAmountUnit = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetEffectiveTime(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.EffectiveTime = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetExpiryTime(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.ExpiryTime = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetRemark(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.Remark = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetPackageType(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.PackageType = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetStatus(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.Status = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetDeductType(v string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.DeductType = &v
	return s
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstance) SetApplicableProducts(v *QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts) *QueryResourcePackageInstancesResponseBodyDataInstancesInstance {
	s.ApplicableProducts = v
	return s
}

type QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts struct {
	Product []*string `json:"Product,omitempty" xml:"Product,omitempty" type:"Repeated"`
}

func (s QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts) String() string {
	return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts) SetProduct(v []*string) *QueryResourcePackageInstancesResponseBodyDataInstancesInstanceApplicableProducts {
	s.Product = v
	return s
}

type QueryResourcePackageInstancesResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryResourcePackageInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryResourcePackageInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryResourcePackageInstancesResponse) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesResponse) SetHeaders(v map[string]*string) *QueryResourcePackageInstancesResponse {
	s.Headers = v
	return s
}

func (s *QueryResourcePackageInstancesResponse) SetBody(v *QueryResourcePackageInstancesResponseBody) *QueryResourcePackageInstancesResponse {
	s.Body = v
	return s
}

type QueryRIUtilizationDetailRequest struct {
	RIInstanceId       *string `json:"RIInstanceId,omitempty" xml:"RIInstanceId,omitempty"`
	InstanceSpec       *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	RICommodityCode    *string `json:"RICommodityCode,omitempty" xml:"RICommodityCode,omitempty"`
	DeductedInstanceId *string `json:"DeductedInstanceId,omitempty" xml:"DeductedInstanceId,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNum            *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryRIUtilizationDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRIUtilizationDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRIUtilizationDetailRequest) SetRIInstanceId(v string) *QueryRIUtilizationDetailRequest {
	s.RIInstanceId = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetInstanceSpec(v string) *QueryRIUtilizationDetailRequest {
	s.InstanceSpec = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetRICommodityCode(v string) *QueryRIUtilizationDetailRequest {
	s.RICommodityCode = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetDeductedInstanceId(v string) *QueryRIUtilizationDetailRequest {
	s.DeductedInstanceId = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetStartTime(v string) *QueryRIUtilizationDetailRequest {
	s.StartTime = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetEndTime(v string) *QueryRIUtilizationDetailRequest {
	s.EndTime = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetPageNum(v int32) *QueryRIUtilizationDetailRequest {
	s.PageNum = &v
	return s
}

func (s *QueryRIUtilizationDetailRequest) SetPageSize(v int32) *QueryRIUtilizationDetailRequest {
	s.PageSize = &v
	return s
}

type QueryRIUtilizationDetailResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryRIUtilizationDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryRIUtilizationDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRIUtilizationDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRIUtilizationDetailResponseBody) SetRequestId(v string) *QueryRIUtilizationDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBody) SetSuccess(v bool) *QueryRIUtilizationDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBody) SetCode(v string) *QueryRIUtilizationDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBody) SetMessage(v string) *QueryRIUtilizationDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBody) SetData(v *QueryRIUtilizationDetailResponseBodyData) *QueryRIUtilizationDetailResponseBody {
	s.Data = v
	return s
}

type QueryRIUtilizationDetailResponseBodyData struct {
	PageNum    *int64                                              `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int64                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DetailList *QueryRIUtilizationDetailResponseBodyDataDetailList `json:"DetailList,omitempty" xml:"DetailList,omitempty" type:"Struct"`
}

func (s QueryRIUtilizationDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRIUtilizationDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRIUtilizationDetailResponseBodyData) SetPageNum(v int64) *QueryRIUtilizationDetailResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyData) SetPageSize(v int64) *QueryRIUtilizationDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyData) SetTotalCount(v int64) *QueryRIUtilizationDetailResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyData) SetDetailList(v *QueryRIUtilizationDetailResponseBodyDataDetailList) *QueryRIUtilizationDetailResponseBodyData {
	s.DetailList = v
	return s
}

type QueryRIUtilizationDetailResponseBodyDataDetailList struct {
	DetailList []*QueryRIUtilizationDetailResponseBodyDataDetailListDetailList `json:"DetailList,omitempty" xml:"DetailList,omitempty" type:"Repeated"`
}

func (s QueryRIUtilizationDetailResponseBodyDataDetailList) String() string {
	return tea.Prettify(s)
}

func (s QueryRIUtilizationDetailResponseBodyDataDetailList) GoString() string {
	return s.String()
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailList) SetDetailList(v []*QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) *QueryRIUtilizationDetailResponseBodyDataDetailList {
	s.DetailList = v
	return s
}

type QueryRIUtilizationDetailResponseBodyDataDetailListDetailList struct {
	RIInstanceId          *string  `json:"RIInstanceId,omitempty" xml:"RIInstanceId,omitempty"`
	InstanceSpec          *string  `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	DeductedInstanceId    *string  `json:"DeductedInstanceId,omitempty" xml:"DeductedInstanceId,omitempty"`
	DeductedCommodityCode *string  `json:"DeductedCommodityCode,omitempty" xml:"DeductedCommodityCode,omitempty"`
	DeductDate            *string  `json:"DeductDate,omitempty" xml:"DeductDate,omitempty"`
	DeductHours           *string  `json:"DeductHours,omitempty" xml:"DeductHours,omitempty"`
	DeductedProductDetail *string  `json:"DeductedProductDetail,omitempty" xml:"DeductedProductDetail,omitempty"`
	DeductQuantity        *float32 `json:"DeductQuantity,omitempty" xml:"DeductQuantity,omitempty"`
	DeductFactorTotal     *float32 `json:"DeductFactorTotal,omitempty" xml:"DeductFactorTotal,omitempty"`
}

func (s QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) String() string {
	return tea.Prettify(s)
}

func (s QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) GoString() string {
	return s.String()
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetRIInstanceId(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.RIInstanceId = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetInstanceSpec(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.InstanceSpec = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductedInstanceId(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductedInstanceId = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductedCommodityCode(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductedCommodityCode = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductDate(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductDate = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductHours(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductHours = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductedProductDetail(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductedProductDetail = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductQuantity(v float32) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductQuantity = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductFactorTotal(v float32) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductFactorTotal = &v
	return s
}

type QueryRIUtilizationDetailResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRIUtilizationDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRIUtilizationDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRIUtilizationDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryRIUtilizationDetailResponse) SetHeaders(v map[string]*string) *QueryRIUtilizationDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryRIUtilizationDetailResponse) SetBody(v *QueryRIUtilizationDetailResponseBody) *QueryRIUtilizationDetailResponse {
	s.Body = v
	return s
}

type QuerySavingsPlansDeductLogRequest struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Locale       *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	PageNum      *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s QuerySavingsPlansDeductLogRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySavingsPlansDeductLogRequest) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDeductLogRequest) SetInstanceType(v string) *QuerySavingsPlansDeductLogRequest {
	s.InstanceType = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) SetPageSize(v int32) *QuerySavingsPlansDeductLogRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) SetLocale(v string) *QuerySavingsPlansDeductLogRequest {
	s.Locale = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) SetPageNum(v int32) *QuerySavingsPlansDeductLogRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) SetInstanceId(v string) *QuerySavingsPlansDeductLogRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) SetStartTime(v string) *QuerySavingsPlansDeductLogRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySavingsPlansDeductLogRequest) SetEndTime(v string) *QuerySavingsPlansDeductLogRequest {
	s.EndTime = &v
	return s
}

type QuerySavingsPlansDeductLogResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
	Message   *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QuerySavingsPlansDeductLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QuerySavingsPlansDeductLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySavingsPlansDeductLogResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDeductLogResponseBody) SetCode(v string) *QuerySavingsPlansDeductLogResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBody) SetRequestId(v string) *QuerySavingsPlansDeductLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBody) SetSuccess(v bool) *QuerySavingsPlansDeductLogResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBody) SetMessage(v string) *QuerySavingsPlansDeductLogResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBody) SetData(v *QuerySavingsPlansDeductLogResponseBodyData) *QuerySavingsPlansDeductLogResponseBody {
	s.Data = v
	return s
}

type QuerySavingsPlansDeductLogResponseBodyData struct {
	PageNum    *int32                                             `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items      []*QuerySavingsPlansDeductLogResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s QuerySavingsPlansDeductLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySavingsPlansDeductLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) SetPageNum(v int32) *QuerySavingsPlansDeductLogResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) SetPageSize(v int32) *QuerySavingsPlansDeductLogResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) SetTotalCount(v int32) *QuerySavingsPlansDeductLogResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) SetItems(v []*QuerySavingsPlansDeductLogResponseBodyDataItems) *QuerySavingsPlansDeductLogResponseBodyData {
	s.Items = v
	return s
}

type QuerySavingsPlansDeductLogResponseBodyDataItems struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	SavingsType      *string `json:"SavingsType,omitempty" xml:"SavingsType,omitempty"`
	BillModule       *string `json:"BillModule,omitempty" xml:"BillModule,omitempty"`
	DeductFee        *string `json:"DeductFee,omitempty" xml:"DeductFee,omitempty"`
	DeductRate       *string `json:"DeductRate,omitempty" xml:"DeductRate,omitempty"`
	UserId           *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeductCommodity  *string `json:"DeductCommodity,omitempty" xml:"DeductCommodity,omitempty"`
	DeductInstanceId *string `json:"DeductInstanceId,omitempty" xml:"DeductInstanceId,omitempty"`
	DiscountRate     *string `json:"DiscountRate,omitempty" xml:"DiscountRate,omitempty"`
}

func (s QuerySavingsPlansDeductLogResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QuerySavingsPlansDeductLogResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetInstanceId(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.InstanceId = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetStartTime(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetEndTime(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.EndTime = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetSavingsType(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.SavingsType = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetBillModule(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.BillModule = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetDeductFee(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.DeductFee = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetDeductRate(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.DeductRate = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetUserId(v int64) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.UserId = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetDeductCommodity(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.DeductCommodity = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetDeductInstanceId(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.DeductInstanceId = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetDiscountRate(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.DiscountRate = &v
	return s
}

type QuerySavingsPlansDeductLogResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySavingsPlansDeductLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySavingsPlansDeductLogResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySavingsPlansDeductLogResponse) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDeductLogResponse) SetHeaders(v map[string]*string) *QuerySavingsPlansDeductLogResponse {
	s.Headers = v
	return s
}

func (s *QuerySavingsPlansDeductLogResponse) SetBody(v *QuerySavingsPlansDeductLogResponseBody) *QuerySavingsPlansDeductLogResponse {
	s.Body = v
	return s
}

type QuerySavingsPlansInstanceRequest struct {
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Locale     *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	PageNum    *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s QuerySavingsPlansInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySavingsPlansInstanceRequest) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansInstanceRequest) SetPageSize(v int32) *QuerySavingsPlansInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetLocale(v string) *QuerySavingsPlansInstanceRequest {
	s.Locale = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetPageNum(v int32) *QuerySavingsPlansInstanceRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetInstanceId(v string) *QuerySavingsPlansInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetStartTime(v string) *QuerySavingsPlansInstanceRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySavingsPlansInstanceRequest) SetEndTime(v string) *QuerySavingsPlansInstanceRequest {
	s.EndTime = &v
	return s
}

type QuerySavingsPlansInstanceResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QuerySavingsPlansInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QuerySavingsPlansInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySavingsPlansInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansInstanceResponseBody) SetCode(v string) *QuerySavingsPlansInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBody) SetRequestId(v string) *QuerySavingsPlansInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBody) SetSuccess(v bool) *QuerySavingsPlansInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBody) SetMessage(v string) *QuerySavingsPlansInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBody) SetData(v *QuerySavingsPlansInstanceResponseBodyData) *QuerySavingsPlansInstanceResponseBody {
	s.Data = v
	return s
}

type QuerySavingsPlansInstanceResponseBodyData struct {
	PageNum    *int32                                            `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items      []*QuerySavingsPlansInstanceResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s QuerySavingsPlansInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySavingsPlansInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansInstanceResponseBodyData) SetPageNum(v int32) *QuerySavingsPlansInstanceResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyData) SetPageSize(v int32) *QuerySavingsPlansInstanceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyData) SetTotalCount(v int32) *QuerySavingsPlansInstanceResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyData) SetItems(v []*QuerySavingsPlansInstanceResponseBodyDataItems) *QuerySavingsPlansInstanceResponseBodyData {
	s.Items = v
	return s
}

type QuerySavingsPlansInstanceResponseBodyDataItems struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SavingsType    *string `json:"SavingsType,omitempty" xml:"SavingsType,omitempty"`
	InstanceFamily *string `json:"InstanceFamily,omitempty" xml:"InstanceFamily,omitempty"`
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	PoolValue      *string `json:"PoolValue,omitempty" xml:"PoolValue,omitempty"`
	Currency       *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PayMode        *string `json:"PayMode,omitempty" xml:"PayMode,omitempty"`
	PrepayFee      *string `json:"PrepayFee,omitempty" xml:"PrepayFee,omitempty"`
	TotalSave      *string `json:"TotalSave,omitempty" xml:"TotalSave,omitempty"`
	Utilization    *string `json:"Utilization,omitempty" xml:"Utilization,omitempty"`
	Share          *bool   `json:"Share,omitempty" xml:"Share,omitempty"`
}

func (s QuerySavingsPlansInstanceResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QuerySavingsPlansInstanceResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetInstanceId(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.InstanceId = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetSavingsType(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.SavingsType = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetInstanceFamily(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.InstanceFamily = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetRegion(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetPoolValue(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.PoolValue = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetCurrency(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.Currency = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetStatus(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetStartTime(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetEndTime(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.EndTime = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetPayMode(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.PayMode = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetPrepayFee(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.PrepayFee = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetTotalSave(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.TotalSave = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetUtilization(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.Utilization = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetShare(v bool) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.Share = &v
	return s
}

type QuerySavingsPlansInstanceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySavingsPlansInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySavingsPlansInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySavingsPlansInstanceResponse) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansInstanceResponse) SetHeaders(v map[string]*string) *QuerySavingsPlansInstanceResponse {
	s.Headers = v
	return s
}

func (s *QuerySavingsPlansInstanceResponse) SetBody(v *QuerySavingsPlansInstanceResponseBody) *QuerySavingsPlansInstanceResponse {
	s.Body = v
	return s
}

type QuerySettleBillRequest struct {
	BillingCycle           *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Type                   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ProductCode            *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType            *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType       *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	IsHideZeroCharge       *bool   `json:"IsHideZeroCharge,omitempty" xml:"IsHideZeroCharge,omitempty"`
	IsDisplayLocalCurrency *bool   `json:"IsDisplayLocalCurrency,omitempty" xml:"IsDisplayLocalCurrency,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	NextToken              *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults             *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	BillOwnerId            *int64  `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
}

func (s QuerySettleBillRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySettleBillRequest) GoString() string {
	return s.String()
}

func (s *QuerySettleBillRequest) SetBillingCycle(v string) *QuerySettleBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *QuerySettleBillRequest) SetType(v string) *QuerySettleBillRequest {
	s.Type = &v
	return s
}

func (s *QuerySettleBillRequest) SetProductCode(v string) *QuerySettleBillRequest {
	s.ProductCode = &v
	return s
}

func (s *QuerySettleBillRequest) SetProductType(v string) *QuerySettleBillRequest {
	s.ProductType = &v
	return s
}

func (s *QuerySettleBillRequest) SetSubscriptionType(v string) *QuerySettleBillRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QuerySettleBillRequest) SetIsHideZeroCharge(v bool) *QuerySettleBillRequest {
	s.IsHideZeroCharge = &v
	return s
}

func (s *QuerySettleBillRequest) SetIsDisplayLocalCurrency(v bool) *QuerySettleBillRequest {
	s.IsDisplayLocalCurrency = &v
	return s
}

func (s *QuerySettleBillRequest) SetOwnerId(v int64) *QuerySettleBillRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySettleBillRequest) SetNextToken(v string) *QuerySettleBillRequest {
	s.NextToken = &v
	return s
}

func (s *QuerySettleBillRequest) SetMaxResults(v int32) *QuerySettleBillRequest {
	s.MaxResults = &v
	return s
}

func (s *QuerySettleBillRequest) SetBillOwnerId(v int64) *QuerySettleBillRequest {
	s.BillOwnerId = &v
	return s
}

type QuerySettleBillResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QuerySettleBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QuerySettleBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySettleBillResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySettleBillResponseBody) SetRequestId(v string) *QuerySettleBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySettleBillResponseBody) SetSuccess(v bool) *QuerySettleBillResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySettleBillResponseBody) SetCode(v string) *QuerySettleBillResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySettleBillResponseBody) SetMessage(v string) *QuerySettleBillResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySettleBillResponseBody) SetData(v *QuerySettleBillResponseBodyData) *QuerySettleBillResponseBody {
	s.Data = v
	return s
}

type QuerySettleBillResponseBodyData struct {
	BillingCycle *string                               `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	AccountID    *string                               `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	AccountName  *string                               `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	NextToken    *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults   *int32                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	TotalCount   *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items        *QuerySettleBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s QuerySettleBillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySettleBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySettleBillResponseBodyData) SetBillingCycle(v string) *QuerySettleBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QuerySettleBillResponseBodyData) SetAccountID(v string) *QuerySettleBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QuerySettleBillResponseBodyData) SetAccountName(v string) *QuerySettleBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QuerySettleBillResponseBodyData) SetNextToken(v string) *QuerySettleBillResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *QuerySettleBillResponseBodyData) SetMaxResults(v int32) *QuerySettleBillResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *QuerySettleBillResponseBodyData) SetTotalCount(v int32) *QuerySettleBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QuerySettleBillResponseBodyData) SetItems(v *QuerySettleBillResponseBodyDataItems) *QuerySettleBillResponseBodyData {
	s.Items = v
	return s
}

type QuerySettleBillResponseBodyDataItems struct {
	Item []*QuerySettleBillResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QuerySettleBillResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QuerySettleBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QuerySettleBillResponseBodyDataItems) SetItem(v []*QuerySettleBillResponseBodyDataItemsItem) *QuerySettleBillResponseBodyDataItems {
	s.Item = v
	return s
}

type QuerySettleBillResponseBodyDataItemsItem struct {
	RecordID              *string  `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	Item                  *string  `json:"Item,omitempty" xml:"Item,omitempty"`
	OwnerID               *string  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	UsageStartTime        *string  `json:"UsageStartTime,omitempty" xml:"UsageStartTime,omitempty"`
	UsageEndTime          *string  `json:"UsageEndTime,omitempty" xml:"UsageEndTime,omitempty"`
	PaymentTime           *string  `json:"PaymentTime,omitempty" xml:"PaymentTime,omitempty"`
	ProductCode           *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType           *string  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType      *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	ProductName           *string  `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductDetail         *string  `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	PretaxGrossAmount     *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	DeductedByCoupons     *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	InvoiceDiscount       *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	PretaxAmount          *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	Currency              *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	PretaxAmountLocal     *float32 `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	Tax                   *float32 `json:"Tax,omitempty" xml:"Tax,omitempty"`
	PaymentAmount         *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	DeductedByCashCoupons *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByPrepaidCard *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	OutstandingAmount     *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	AfterTaxAmount        *float32 `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	Status                *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	PaymentCurrency       *string  `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	PaymentTransactionID  *string  `json:"PaymentTransactionID,omitempty" xml:"PaymentTransactionID,omitempty"`
	RoundDownDiscount     *string  `json:"RoundDownDiscount,omitempty" xml:"RoundDownDiscount,omitempty"`
	SubOrderId            *string  `json:"SubOrderId,omitempty" xml:"SubOrderId,omitempty"`
	PipCode               *string  `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	CommodityCode         *string  `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s QuerySettleBillResponseBodyDataItemsItem) String() string {
	return tea.Prettify(s)
}

func (s QuerySettleBillResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetRecordID(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.RecordID = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetItem(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.Item = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetOwnerID(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetUsageStartTime(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.UsageStartTime = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetUsageEndTime(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.UsageEndTime = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPaymentTime(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.PaymentTime = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetProductCode(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetProductType(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetSubscriptionType(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetProductName(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.ProductName = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetProductDetail(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.ProductDetail = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetCurrency(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPretaxAmountLocal(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetTax(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.Tax = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetAfterTaxAmount(v float32) *QuerySettleBillResponseBodyDataItemsItem {
	s.AfterTaxAmount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetStatus(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.Status = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPaymentCurrency(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.PaymentCurrency = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPaymentTransactionID(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.PaymentTransactionID = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetRoundDownDiscount(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.RoundDownDiscount = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetSubOrderId(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.SubOrderId = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetPipCode(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.PipCode = &v
	return s
}

func (s *QuerySettleBillResponseBodyDataItemsItem) SetCommodityCode(v string) *QuerySettleBillResponseBodyDataItemsItem {
	s.CommodityCode = &v
	return s
}

type QuerySettleBillResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySettleBillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySettleBillResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySettleBillResponse) GoString() string {
	return s.String()
}

func (s *QuerySettleBillResponse) SetHeaders(v map[string]*string) *QuerySettleBillResponse {
	s.Headers = v
	return s
}

func (s *QuerySettleBillResponse) SetBody(v *QuerySettleBillResponseBody) *QuerySettleBillResponse {
	s.Body = v
	return s
}

type QuerySettlementBillRequest struct {
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	BillingCycle     *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime          *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	IsHideZeroCharge *bool   `json:"IsHideZeroCharge,omitempty" xml:"IsHideZeroCharge,omitempty"`
}

func (s QuerySettlementBillRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySettlementBillRequest) GoString() string {
	return s.String()
}

func (s *QuerySettlementBillRequest) SetPageSize(v int32) *QuerySettlementBillRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySettlementBillRequest) SetOwnerId(v int64) *QuerySettlementBillRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySettlementBillRequest) SetPageNum(v int32) *QuerySettlementBillRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySettlementBillRequest) SetBillingCycle(v string) *QuerySettlementBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *QuerySettlementBillRequest) SetStartTime(v string) *QuerySettlementBillRequest {
	s.StartTime = &v
	return s
}

func (s *QuerySettlementBillRequest) SetEndTime(v string) *QuerySettlementBillRequest {
	s.EndTime = &v
	return s
}

func (s *QuerySettlementBillRequest) SetType(v string) *QuerySettlementBillRequest {
	s.Type = &v
	return s
}

func (s *QuerySettlementBillRequest) SetProductCode(v string) *QuerySettlementBillRequest {
	s.ProductCode = &v
	return s
}

func (s *QuerySettlementBillRequest) SetProductType(v string) *QuerySettlementBillRequest {
	s.ProductType = &v
	return s
}

func (s *QuerySettlementBillRequest) SetSubscriptionType(v string) *QuerySettlementBillRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QuerySettlementBillRequest) SetIsHideZeroCharge(v bool) *QuerySettlementBillRequest {
	s.IsHideZeroCharge = &v
	return s
}

type QuerySettlementBillResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QuerySettlementBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QuerySettlementBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySettlementBillResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySettlementBillResponseBody) SetRequestId(v string) *QuerySettlementBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySettlementBillResponseBody) SetSuccess(v bool) *QuerySettlementBillResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySettlementBillResponseBody) SetCode(v string) *QuerySettlementBillResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySettlementBillResponseBody) SetMessage(v string) *QuerySettlementBillResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySettlementBillResponseBody) SetData(v *QuerySettlementBillResponseBodyData) *QuerySettlementBillResponseBody {
	s.Data = v
	return s
}

type QuerySettlementBillResponseBodyData struct {
	PageNum      *int32                                    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount   *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	BillingCycle *string                                   `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Items        *QuerySettlementBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s QuerySettlementBillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySettlementBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySettlementBillResponseBodyData) SetPageNum(v int32) *QuerySettlementBillResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QuerySettlementBillResponseBodyData) SetPageSize(v int32) *QuerySettlementBillResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QuerySettlementBillResponseBodyData) SetTotalCount(v int32) *QuerySettlementBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QuerySettlementBillResponseBodyData) SetBillingCycle(v string) *QuerySettlementBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QuerySettlementBillResponseBodyData) SetItems(v *QuerySettlementBillResponseBodyDataItems) *QuerySettlementBillResponseBodyData {
	s.Items = v
	return s
}

type QuerySettlementBillResponseBodyDataItems struct {
	Item []*QuerySettlementBillResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QuerySettlementBillResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QuerySettlementBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QuerySettlementBillResponseBodyDataItems) SetItem(v []*QuerySettlementBillResponseBodyDataItemsItem) *QuerySettlementBillResponseBodyDataItems {
	s.Item = v
	return s
}

type QuerySettlementBillResponseBodyDataItemsItem struct {
	RecordID                    *string  `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	Item                        *string  `json:"Item,omitempty" xml:"Item,omitempty"`
	PayerAccount                *string  `json:"PayerAccount,omitempty" xml:"PayerAccount,omitempty"`
	OwnerID                     *string  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	CreateTime                  *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UsageStartTime              *string  `json:"UsageStartTime,omitempty" xml:"UsageStartTime,omitempty"`
	UsageEndTime                *string  `json:"UsageEndTime,omitempty" xml:"UsageEndTime,omitempty"`
	SuborderID                  *string  `json:"SuborderID,omitempty" xml:"SuborderID,omitempty"`
	OrderID                     *string  `json:"OrderID,omitempty" xml:"OrderID,omitempty"`
	OrderType                   *string  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	LinkedCustomerOrderID       *string  `json:"LinkedCustomerOrderID,omitempty" xml:"LinkedCustomerOrderID,omitempty"`
	OriginalOrderID             *string  `json:"OriginalOrderID,omitempty" xml:"OriginalOrderID,omitempty"`
	PaymentTime                 *string  `json:"PaymentTime,omitempty" xml:"PaymentTime,omitempty"`
	SolutionID                  *string  `json:"SolutionID,omitempty" xml:"SolutionID,omitempty"`
	SolutionName                *string  `json:"SolutionName,omitempty" xml:"SolutionName,omitempty"`
	BillID                      *string  `json:"BillID,omitempty" xml:"BillID,omitempty"`
	ProductCode                 *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType                 *string  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType            *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	Region                      *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	Config                      *string  `json:"Config,omitempty" xml:"Config,omitempty"`
	Quantity                    *string  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	PretaxGrossAmount           *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	ChargeDiscount              *float32 `json:"ChargeDiscount,omitempty" xml:"ChargeDiscount,omitempty"`
	DeductedByCoupons           *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	AccountDiscount             *float32 `json:"AccountDiscount,omitempty" xml:"AccountDiscount,omitempty"`
	Promotion                   *string  `json:"Promotion,omitempty" xml:"Promotion,omitempty"`
	PretaxAmount                *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	Currency                    *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	PretaxAmountLocal           *float32 `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	PreviousBillingCycleBalance *float32 `json:"PreviousBillingCycleBalance,omitempty" xml:"PreviousBillingCycleBalance,omitempty"`
	Tax                         *float32 `json:"Tax,omitempty" xml:"Tax,omitempty"`
	AfterTaxAmount              *float32 `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	Status                      *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	ClearedTime                 *string  `json:"ClearedTime,omitempty" xml:"ClearedTime,omitempty"`
	OutstandingAmount           *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	DeductedByCashCoupons       *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByPrepaidCard       *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	MybankPaymentAmount         *float32 `json:"MybankPaymentAmount,omitempty" xml:"MybankPaymentAmount,omitempty"`
	PaymentAmount               *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	PaymentCurrency             *string  `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	Seller                      *string  `json:"Seller,omitempty" xml:"Seller,omitempty"`
	InvoiceNo                   *string  `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
}

func (s QuerySettlementBillResponseBodyDataItemsItem) String() string {
	return tea.Prettify(s)
}

func (s QuerySettlementBillResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetRecordID(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.RecordID = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetItem(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.Item = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetPayerAccount(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.PayerAccount = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetOwnerID(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetCreateTime(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.CreateTime = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetUsageStartTime(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.UsageStartTime = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetUsageEndTime(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.UsageEndTime = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetSuborderID(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.SuborderID = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetOrderID(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.OrderID = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetOrderType(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.OrderType = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetLinkedCustomerOrderID(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.LinkedCustomerOrderID = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetOriginalOrderID(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.OriginalOrderID = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetPaymentTime(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.PaymentTime = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetSolutionID(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.SolutionID = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetSolutionName(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.SolutionName = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetBillID(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.BillID = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetProductCode(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetProductType(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetSubscriptionType(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetRegion(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.Region = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetConfig(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.Config = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetQuantity(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.Quantity = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetChargeDiscount(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.ChargeDiscount = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetAccountDiscount(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.AccountDiscount = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetPromotion(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.Promotion = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetCurrency(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetPretaxAmountLocal(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetPreviousBillingCycleBalance(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.PreviousBillingCycleBalance = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetTax(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.Tax = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetAfterTaxAmount(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.AfterTaxAmount = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetStatus(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.Status = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetClearedTime(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.ClearedTime = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetMybankPaymentAmount(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.MybankPaymentAmount = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QuerySettlementBillResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetPaymentCurrency(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.PaymentCurrency = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetSeller(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.Seller = &v
	return s
}

func (s *QuerySettlementBillResponseBodyDataItemsItem) SetInvoiceNo(v string) *QuerySettlementBillResponseBodyDataItemsItem {
	s.InvoiceNo = &v
	return s
}

type QuerySettlementBillResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySettlementBillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySettlementBillResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySettlementBillResponse) GoString() string {
	return s.String()
}

func (s *QuerySettlementBillResponse) SetHeaders(v map[string]*string) *QuerySettlementBillResponse {
	s.Headers = v
	return s
}

func (s *QuerySettlementBillResponse) SetBody(v *QuerySettlementBillResponseBody) *QuerySettlementBillResponse {
	s.Body = v
	return s
}

type QuerySplitItemBillRequest struct {
	BillingCycle     *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	BillOwnerId      *int64  `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
}

func (s QuerySplitItemBillRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySplitItemBillRequest) GoString() string {
	return s.String()
}

func (s *QuerySplitItemBillRequest) SetBillingCycle(v string) *QuerySplitItemBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetProductCode(v string) *QuerySplitItemBillRequest {
	s.ProductCode = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetProductType(v string) *QuerySplitItemBillRequest {
	s.ProductType = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetSubscriptionType(v string) *QuerySplitItemBillRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetOwnerId(v int64) *QuerySplitItemBillRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetPageNum(v int32) *QuerySplitItemBillRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetPageSize(v int32) *QuerySplitItemBillRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetBillOwnerId(v int64) *QuerySplitItemBillRequest {
	s.BillOwnerId = &v
	return s
}

type QuerySplitItemBillResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QuerySplitItemBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QuerySplitItemBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySplitItemBillResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySplitItemBillResponseBody) SetRequestId(v string) *QuerySplitItemBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySplitItemBillResponseBody) SetSuccess(v bool) *QuerySplitItemBillResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySplitItemBillResponseBody) SetCode(v string) *QuerySplitItemBillResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySplitItemBillResponseBody) SetMessage(v string) *QuerySplitItemBillResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySplitItemBillResponseBody) SetData(v *QuerySplitItemBillResponseBodyData) *QuerySplitItemBillResponseBody {
	s.Data = v
	return s
}

type QuerySplitItemBillResponseBodyData struct {
	BillingCycle *string                                  `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	AccountID    *string                                  `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	AccountName  *string                                  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	TotalCount   *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageNum      *int32                                   `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Items        *QuerySplitItemBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s QuerySplitItemBillResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySplitItemBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySplitItemBillResponseBodyData) SetBillingCycle(v string) *QuerySplitItemBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) SetAccountID(v string) *QuerySplitItemBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) SetAccountName(v string) *QuerySplitItemBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) SetTotalCount(v int32) *QuerySplitItemBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) SetPageNum(v int32) *QuerySplitItemBillResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) SetPageSize(v int32) *QuerySplitItemBillResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) SetItems(v *QuerySplitItemBillResponseBodyDataItems) *QuerySplitItemBillResponseBodyData {
	s.Items = v
	return s
}

type QuerySplitItemBillResponseBodyDataItems struct {
	Item []*QuerySplitItemBillResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QuerySplitItemBillResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QuerySplitItemBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QuerySplitItemBillResponseBodyDataItems) SetItem(v []*QuerySplitItemBillResponseBodyDataItemsItem) *QuerySplitItemBillResponseBodyDataItems {
	s.Item = v
	return s
}

type QuerySplitItemBillResponseBodyDataItemsItem struct {
	InstanceID                *string  `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	BillingType               *string  `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CostUnit                  *string  `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
	ProductCode               *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType               *string  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType          *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	ProductName               *string  `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductDetail             *string  `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	OwnerID                   *string  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	BillingItem               *string  `json:"BillingItem,omitempty" xml:"BillingItem,omitempty"`
	ListPrice                 *string  `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	ListPriceUnit             *string  `json:"ListPriceUnit,omitempty" xml:"ListPriceUnit,omitempty"`
	Usage                     *string  `json:"Usage,omitempty" xml:"Usage,omitempty"`
	UsageUnit                 *string  `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty"`
	DeductedByResourcePackage *string  `json:"DeductedByResourcePackage,omitempty" xml:"DeductedByResourcePackage,omitempty"`
	PretaxGrossAmount         *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	InvoiceDiscount           *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	DeductedByCoupons         *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	PretaxAmount              *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	DeductedByCashCoupons     *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByPrepaidCard     *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	PaymentAmount             *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	OutstandingAmount         *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	Currency                  *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	NickName                  *string  `json:"NickName,omitempty" xml:"NickName,omitempty"`
	ResourceGroup             *string  `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	Tag                       *string  `json:"Tag,omitempty" xml:"Tag,omitempty"`
	InstanceConfig            *string  `json:"InstanceConfig,omitempty" xml:"InstanceConfig,omitempty"`
	InstanceSpec              *string  `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	InternetIP                *string  `json:"InternetIP,omitempty" xml:"InternetIP,omitempty"`
	IntranetIP                *string  `json:"IntranetIP,omitempty" xml:"IntranetIP,omitempty"`
	Region                    *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	Zone                      *string  `json:"Zone,omitempty" xml:"Zone,omitempty"`
	Item                      *string  `json:"Item,omitempty" xml:"Item,omitempty"`
	ServicePeriod             *string  `json:"ServicePeriod,omitempty" xml:"ServicePeriod,omitempty"`
	BillingDate               *string  `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	SplitItemID               *string  `json:"SplitItemID,omitempty" xml:"SplitItemID,omitempty"`
	SplitItemName             *string  `json:"SplitItemName,omitempty" xml:"SplitItemName,omitempty"`
	PipCode                   *string  `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	CommodityCode             *string  `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ServicePeriodUnit         *string  `json:"ServicePeriodUnit,omitempty" xml:"ServicePeriodUnit,omitempty"`
	SplitCommodityCode        *string  `json:"SplitCommodityCode,omitempty" xml:"SplitCommodityCode,omitempty"`
	SplitProductDetail        *string  `json:"SplitProductDetail,omitempty" xml:"SplitProductDetail,omitempty"`
	SplitAccountID            *string  `json:"SplitAccountID,omitempty" xml:"SplitAccountID,omitempty"`
	SplitAccountName          *string  `json:"SplitAccountName,omitempty" xml:"SplitAccountName,omitempty"`
	SplitBillingCycle         *string  `json:"SplitBillingCycle,omitempty" xml:"SplitBillingCycle,omitempty"`
}

func (s QuerySplitItemBillResponseBodyDataItemsItem) String() string {
	return tea.Prettify(s)
}

func (s QuerySplitItemBillResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetInstanceID(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.InstanceID = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetBillingType(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.BillingType = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetCostUnit(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.CostUnit = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetProductCode(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetProductType(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSubscriptionType(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetProductName(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ProductName = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetProductDetail(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ProductDetail = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetOwnerID(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetBillingItem(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.BillingItem = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetListPrice(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ListPrice = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetListPriceUnit(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ListPriceUnit = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetUsage(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.Usage = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetUsageUnit(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.UsageUnit = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetDeductedByResourcePackage(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.DeductedByResourcePackage = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetCurrency(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetNickName(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.NickName = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetResourceGroup(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ResourceGroup = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetTag(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.Tag = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetInstanceConfig(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.InstanceConfig = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetInstanceSpec(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.InstanceSpec = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetInternetIP(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.InternetIP = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetIntranetIP(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.IntranetIP = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetRegion(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.Region = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetZone(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.Zone = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetItem(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.Item = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetServicePeriod(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ServicePeriod = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetBillingDate(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.BillingDate = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitItemID(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitItemID = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitItemName(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitItemName = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetPipCode(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.PipCode = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetCommodityCode(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.CommodityCode = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetServicePeriodUnit(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ServicePeriodUnit = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitCommodityCode(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitCommodityCode = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitProductDetail(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitProductDetail = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitAccountID(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitAccountID = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitAccountName(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitAccountName = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitBillingCycle(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitBillingCycle = &v
	return s
}

type QuerySplitItemBillResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySplitItemBillResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySplitItemBillResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySplitItemBillResponse) GoString() string {
	return s.String()
}

func (s *QuerySplitItemBillResponse) SetHeaders(v map[string]*string) *QuerySplitItemBillResponse {
	s.Headers = v
	return s
}

func (s *QuerySplitItemBillResponse) SetBody(v *QuerySplitItemBillResponseBody) *QuerySplitItemBillResponse {
	s.Body = v
	return s
}

type QueryUserOmsDataRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Table     *string `json:"Table,omitempty" xml:"Table,omitempty"`
	DataType  *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Marker    *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryUserOmsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryUserOmsDataRequest) GoString() string {
	return s.String()
}

func (s *QueryUserOmsDataRequest) SetOwnerId(v int64) *QueryUserOmsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryUserOmsDataRequest) SetTable(v string) *QueryUserOmsDataRequest {
	s.Table = &v
	return s
}

func (s *QueryUserOmsDataRequest) SetDataType(v string) *QueryUserOmsDataRequest {
	s.DataType = &v
	return s
}

func (s *QueryUserOmsDataRequest) SetStartTime(v string) *QueryUserOmsDataRequest {
	s.StartTime = &v
	return s
}

func (s *QueryUserOmsDataRequest) SetEndTime(v string) *QueryUserOmsDataRequest {
	s.EndTime = &v
	return s
}

func (s *QueryUserOmsDataRequest) SetMarker(v string) *QueryUserOmsDataRequest {
	s.Marker = &v
	return s
}

func (s *QueryUserOmsDataRequest) SetPageSize(v int32) *QueryUserOmsDataRequest {
	s.PageSize = &v
	return s
}

type QueryUserOmsDataResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *QueryUserOmsDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryUserOmsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryUserOmsDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryUserOmsDataResponseBody) SetRequestId(v string) *QueryUserOmsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryUserOmsDataResponseBody) SetSuccess(v bool) *QueryUserOmsDataResponseBody {
	s.Success = &v
	return s
}

func (s *QueryUserOmsDataResponseBody) SetCode(v string) *QueryUserOmsDataResponseBody {
	s.Code = &v
	return s
}

func (s *QueryUserOmsDataResponseBody) SetMessage(v string) *QueryUserOmsDataResponseBody {
	s.Message = &v
	return s
}

func (s *QueryUserOmsDataResponseBody) SetData(v *QueryUserOmsDataResponseBodyData) *QueryUserOmsDataResponseBody {
	s.Data = v
	return s
}

type QueryUserOmsDataResponseBodyData struct {
	Marker  *string              `json:"Marker,omitempty" xml:"Marker,omitempty"`
	HostId  *string              `json:"HostId,omitempty" xml:"HostId,omitempty"`
	OmsData []map[string]*string `json:"OmsData,omitempty" xml:"OmsData,omitempty" type:"Repeated"`
}

func (s QueryUserOmsDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryUserOmsDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryUserOmsDataResponseBodyData) SetMarker(v string) *QueryUserOmsDataResponseBodyData {
	s.Marker = &v
	return s
}

func (s *QueryUserOmsDataResponseBodyData) SetHostId(v string) *QueryUserOmsDataResponseBodyData {
	s.HostId = &v
	return s
}

func (s *QueryUserOmsDataResponseBodyData) SetOmsData(v []map[string]*string) *QueryUserOmsDataResponseBodyData {
	s.OmsData = v
	return s
}

type QueryUserOmsDataResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryUserOmsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryUserOmsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryUserOmsDataResponse) GoString() string {
	return s.String()
}

func (s *QueryUserOmsDataResponse) SetHeaders(v map[string]*string) *QueryUserOmsDataResponse {
	s.Headers = v
	return s
}

func (s *QueryUserOmsDataResponse) SetBody(v *QueryUserOmsDataResponseBody) *QueryUserOmsDataResponse {
	s.Body = v
	return s
}

type RenewInstanceRequest struct {
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RenewPeriod *int32  `json:"RenewPeriod,omitempty" xml:"RenewPeriod,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetProductCode(v string) *RenewInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetRenewPeriod(v int32) *RenewInstanceRequest {
	s.RenewPeriod = &v
	return s
}

func (s *RenewInstanceRequest) SetClientToken(v string) *RenewInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewInstanceRequest) SetProductType(v string) *RenewInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *RenewInstanceRequest) SetOwnerId(v int64) *RenewInstanceRequest {
	s.OwnerId = &v
	return s
}

type RenewInstanceResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *RenewInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetSuccess(v bool) *RenewInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RenewInstanceResponseBody) SetCode(v string) *RenewInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RenewInstanceResponseBody) SetMessage(v string) *RenewInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *RenewInstanceResponseBody) SetData(v *RenewInstanceResponseBodyData) *RenewInstanceResponseBody {
	s.Data = v
	return s
}

type RenewInstanceResponseBodyData struct {
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s RenewInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBodyData) SetOrderId(v string) *RenewInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

type RenewInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type RenewResourcePackageRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	EffectiveDate *string `json:"EffectiveDate,omitempty" xml:"EffectiveDate,omitempty"`
	Duration      *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PricingCycle  *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s RenewResourcePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *RenewResourcePackageRequest) SetOwnerId(v int64) *RenewResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewResourcePackageRequest) SetInstanceId(v string) *RenewResourcePackageRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewResourcePackageRequest) SetEffectiveDate(v string) *RenewResourcePackageRequest {
	s.EffectiveDate = &v
	return s
}

func (s *RenewResourcePackageRequest) SetDuration(v int32) *RenewResourcePackageRequest {
	s.Duration = &v
	return s
}

func (s *RenewResourcePackageRequest) SetPricingCycle(v string) *RenewResourcePackageRequest {
	s.PricingCycle = &v
	return s
}

type RenewResourcePackageResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *int64                                `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *RenewResourcePackageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RenewResourcePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *RenewResourcePackageResponseBody) SetRequestId(v string) *RenewResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewResourcePackageResponseBody) SetOrderId(v int64) *RenewResourcePackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewResourcePackageResponseBody) SetSuccess(v bool) *RenewResourcePackageResponseBody {
	s.Success = &v
	return s
}

func (s *RenewResourcePackageResponseBody) SetCode(v string) *RenewResourcePackageResponseBody {
	s.Code = &v
	return s
}

func (s *RenewResourcePackageResponseBody) SetMessage(v string) *RenewResourcePackageResponseBody {
	s.Message = &v
	return s
}

func (s *RenewResourcePackageResponseBody) SetData(v *RenewResourcePackageResponseBodyData) *RenewResourcePackageResponseBody {
	s.Data = v
	return s
}

type RenewResourcePackageResponseBodyData struct {
	OrderId    *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RenewResourcePackageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RenewResourcePackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewResourcePackageResponseBodyData) SetOrderId(v int64) *RenewResourcePackageResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *RenewResourcePackageResponseBodyData) SetInstanceId(v string) *RenewResourcePackageResponseBodyData {
	s.InstanceId = &v
	return s
}

type RenewResourcePackageResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewResourcePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *RenewResourcePackageResponse) SetHeaders(v map[string]*string) *RenewResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *RenewResourcePackageResponse) SetBody(v *RenewResourcePackageResponseBody) *RenewResourcePackageResponse {
	s.Body = v
	return s
}

type SaveUserCreditRequest struct {
	AvoidExpiration          *bool   `json:"AvoidExpiration,omitempty" xml:"AvoidExpiration,omitempty"`
	AvoidPrepaidNotification *bool   `json:"AvoidPrepaidNotification,omitempty" xml:"AvoidPrepaidNotification,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AvoidNotification        *bool   `json:"AvoidNotification,omitempty" xml:"AvoidNotification,omitempty"`
	CreditValue              *string `json:"CreditValue,omitempty" xml:"CreditValue,omitempty"`
	AvoidPrepaidExpiration   *bool   `json:"AvoidPrepaidExpiration,omitempty" xml:"AvoidPrepaidExpiration,omitempty"`
	Operator                 *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	CreditType               *string `json:"CreditType,omitempty" xml:"CreditType,omitempty"`
}

func (s SaveUserCreditRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveUserCreditRequest) GoString() string {
	return s.String()
}

func (s *SaveUserCreditRequest) SetAvoidExpiration(v bool) *SaveUserCreditRequest {
	s.AvoidExpiration = &v
	return s
}

func (s *SaveUserCreditRequest) SetAvoidPrepaidNotification(v bool) *SaveUserCreditRequest {
	s.AvoidPrepaidNotification = &v
	return s
}

func (s *SaveUserCreditRequest) SetDescription(v string) *SaveUserCreditRequest {
	s.Description = &v
	return s
}

func (s *SaveUserCreditRequest) SetAvoidNotification(v bool) *SaveUserCreditRequest {
	s.AvoidNotification = &v
	return s
}

func (s *SaveUserCreditRequest) SetCreditValue(v string) *SaveUserCreditRequest {
	s.CreditValue = &v
	return s
}

func (s *SaveUserCreditRequest) SetAvoidPrepaidExpiration(v bool) *SaveUserCreditRequest {
	s.AvoidPrepaidExpiration = &v
	return s
}

func (s *SaveUserCreditRequest) SetOperator(v string) *SaveUserCreditRequest {
	s.Operator = &v
	return s
}

func (s *SaveUserCreditRequest) SetCreditType(v string) *SaveUserCreditRequest {
	s.CreditType = &v
	return s
}

type SaveUserCreditResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SaveUserCreditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveUserCreditResponseBody) GoString() string {
	return s.String()
}

func (s *SaveUserCreditResponseBody) SetCode(v string) *SaveUserCreditResponseBody {
	s.Code = &v
	return s
}

func (s *SaveUserCreditResponseBody) SetSuccess(v bool) *SaveUserCreditResponseBody {
	s.Success = &v
	return s
}

func (s *SaveUserCreditResponseBody) SetRequestId(v string) *SaveUserCreditResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveUserCreditResponseBody) SetMessage(v string) *SaveUserCreditResponseBody {
	s.Message = &v
	return s
}

type SaveUserCreditResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveUserCreditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveUserCreditResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveUserCreditResponse) GoString() string {
	return s.String()
}

func (s *SaveUserCreditResponse) SetHeaders(v map[string]*string) *SaveUserCreditResponse {
	s.Headers = v
	return s
}

func (s *SaveUserCreditResponse) SetBody(v *SaveUserCreditResponseBody) *SaveUserCreditResponse {
	s.Body = v
	return s
}

type SetCreditLabelActionRequest struct {
	ActionType                *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	Uid                       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	ClearCycle                *string `json:"ClearCycle,omitempty" xml:"ClearCycle,omitempty"`
	CreditAmount              *string `json:"CreditAmount,omitempty" xml:"CreditAmount,omitempty"`
	CurrencyCode              *string `json:"CurrencyCode,omitempty" xml:"CurrencyCode,omitempty"`
	DailyCycle                *string `json:"DailyCycle,omitempty" xml:"DailyCycle,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IsNeedAddSettleLabel      *string `json:"IsNeedAddSettleLabel,omitempty" xml:"IsNeedAddSettleLabel,omitempty"`
	IsNeedAdjustCreditAccount *string `json:"IsNeedAdjustCreditAccount,omitempty" xml:"IsNeedAdjustCreditAccount,omitempty"`
	IsNeedSaveNotifyRule      *string `json:"IsNeedSaveNotifyRule,omitempty" xml:"IsNeedSaveNotifyRule,omitempty"`
	IsNeedSetCreditAmount     *string `json:"IsNeedSetCreditAmount,omitempty" xml:"IsNeedSetCreditAmount,omitempty"`
	NeedNotice                *bool   `json:"NeedNotice,omitempty" xml:"NeedNotice,omitempty"`
	NewCreateMode             *bool   `json:"NewCreateMode,omitempty" xml:"NewCreateMode,omitempty"`
	Operator                  *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteCode                  *string `json:"SiteCode,omitempty" xml:"SiteCode,omitempty"`
	Source                    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s SetCreditLabelActionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCreditLabelActionRequest) GoString() string {
	return s.String()
}

func (s *SetCreditLabelActionRequest) SetActionType(v string) *SetCreditLabelActionRequest {
	s.ActionType = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetUid(v string) *SetCreditLabelActionRequest {
	s.Uid = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetClearCycle(v string) *SetCreditLabelActionRequest {
	s.ClearCycle = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetCreditAmount(v string) *SetCreditLabelActionRequest {
	s.CreditAmount = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetCurrencyCode(v string) *SetCreditLabelActionRequest {
	s.CurrencyCode = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetDailyCycle(v string) *SetCreditLabelActionRequest {
	s.DailyCycle = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetDescription(v string) *SetCreditLabelActionRequest {
	s.Description = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetIsNeedAddSettleLabel(v string) *SetCreditLabelActionRequest {
	s.IsNeedAddSettleLabel = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetIsNeedAdjustCreditAccount(v string) *SetCreditLabelActionRequest {
	s.IsNeedAdjustCreditAccount = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetIsNeedSaveNotifyRule(v string) *SetCreditLabelActionRequest {
	s.IsNeedSaveNotifyRule = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetIsNeedSetCreditAmount(v string) *SetCreditLabelActionRequest {
	s.IsNeedSetCreditAmount = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetNeedNotice(v bool) *SetCreditLabelActionRequest {
	s.NeedNotice = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetNewCreateMode(v bool) *SetCreditLabelActionRequest {
	s.NewCreateMode = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetOperator(v string) *SetCreditLabelActionRequest {
	s.Operator = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetRequestId(v string) *SetCreditLabelActionRequest {
	s.RequestId = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetSiteCode(v string) *SetCreditLabelActionRequest {
	s.SiteCode = &v
	return s
}

func (s *SetCreditLabelActionRequest) SetSource(v string) *SetCreditLabelActionRequest {
	s.Source = &v
	return s
}

type SetCreditLabelActionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCreditLabelActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetCreditLabelActionResponseBody) GoString() string {
	return s.String()
}

func (s *SetCreditLabelActionResponseBody) SetCode(v string) *SetCreditLabelActionResponseBody {
	s.Code = &v
	return s
}

func (s *SetCreditLabelActionResponseBody) SetData(v bool) *SetCreditLabelActionResponseBody {
	s.Data = &v
	return s
}

func (s *SetCreditLabelActionResponseBody) SetMessage(v string) *SetCreditLabelActionResponseBody {
	s.Message = &v
	return s
}

func (s *SetCreditLabelActionResponseBody) SetSuccess(v bool) *SetCreditLabelActionResponseBody {
	s.Success = &v
	return s
}

func (s *SetCreditLabelActionResponseBody) SetRequestId(v string) *SetCreditLabelActionResponseBody {
	s.RequestId = &v
	return s
}

type SetCreditLabelActionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetCreditLabelActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetCreditLabelActionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCreditLabelActionResponse) GoString() string {
	return s.String()
}

func (s *SetCreditLabelActionResponse) SetHeaders(v map[string]*string) *SetCreditLabelActionResponse {
	s.Headers = v
	return s
}

func (s *SetCreditLabelActionResponse) SetBody(v *SetCreditLabelActionResponseBody) *SetCreditLabelActionResponse {
	s.Body = v
	return s
}

type SetRenewalRequest struct {
	RenewalPeriod     *int32  `json:"RenewalPeriod,omitempty" xml:"RenewalPeriod,omitempty"`
	InstanceIDs       *string `json:"InstanceIDs,omitempty" xml:"InstanceIDs,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType       *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SubscriptionType  *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	RenewalPeriodUnit *string `json:"RenewalPeriodUnit,omitempty" xml:"RenewalPeriodUnit,omitempty"`
	RenewalStatus     *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
}

func (s SetRenewalRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRenewalRequest) GoString() string {
	return s.String()
}

func (s *SetRenewalRequest) SetRenewalPeriod(v int32) *SetRenewalRequest {
	s.RenewalPeriod = &v
	return s
}

func (s *SetRenewalRequest) SetInstanceIDs(v string) *SetRenewalRequest {
	s.InstanceIDs = &v
	return s
}

func (s *SetRenewalRequest) SetOwnerId(v int64) *SetRenewalRequest {
	s.OwnerId = &v
	return s
}

func (s *SetRenewalRequest) SetProductCode(v string) *SetRenewalRequest {
	s.ProductCode = &v
	return s
}

func (s *SetRenewalRequest) SetProductType(v string) *SetRenewalRequest {
	s.ProductType = &v
	return s
}

func (s *SetRenewalRequest) SetSubscriptionType(v string) *SetRenewalRequest {
	s.SubscriptionType = &v
	return s
}

func (s *SetRenewalRequest) SetRenewalPeriodUnit(v string) *SetRenewalRequest {
	s.RenewalPeriodUnit = &v
	return s
}

func (s *SetRenewalRequest) SetRenewalStatus(v string) *SetRenewalRequest {
	s.RenewalStatus = &v
	return s
}

type SetRenewalResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SetRenewalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRenewalResponseBody) GoString() string {
	return s.String()
}

func (s *SetRenewalResponseBody) SetRequestId(v string) *SetRenewalResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetRenewalResponseBody) SetSuccess(v bool) *SetRenewalResponseBody {
	s.Success = &v
	return s
}

func (s *SetRenewalResponseBody) SetCode(v string) *SetRenewalResponseBody {
	s.Code = &v
	return s
}

func (s *SetRenewalResponseBody) SetMessage(v string) *SetRenewalResponseBody {
	s.Message = &v
	return s
}

type SetRenewalResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetRenewalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetRenewalResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRenewalResponse) GoString() string {
	return s.String()
}

func (s *SetRenewalResponse) SetHeaders(v map[string]*string) *SetRenewalResponse {
	s.Headers = v
	return s
}

func (s *SetRenewalResponse) SetBody(v *SetRenewalResponseBody) *SetRenewalResponse {
	s.Body = v
	return s
}

type SetResellerUserAlarmThresholdRequest struct {
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AlarmType       *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	AlarmThresholds *string `json:"AlarmThresholds,omitempty" xml:"AlarmThresholds,omitempty"`
}

func (s SetResellerUserAlarmThresholdRequest) String() string {
	return tea.Prettify(s)
}

func (s SetResellerUserAlarmThresholdRequest) GoString() string {
	return s.String()
}

func (s *SetResellerUserAlarmThresholdRequest) SetOwnerId(v int64) *SetResellerUserAlarmThresholdRequest {
	s.OwnerId = &v
	return s
}

func (s *SetResellerUserAlarmThresholdRequest) SetAlarmType(v string) *SetResellerUserAlarmThresholdRequest {
	s.AlarmType = &v
	return s
}

func (s *SetResellerUserAlarmThresholdRequest) SetAlarmThresholds(v string) *SetResellerUserAlarmThresholdRequest {
	s.AlarmThresholds = &v
	return s
}

type SetResellerUserAlarmThresholdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s SetResellerUserAlarmThresholdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetResellerUserAlarmThresholdResponseBody) GoString() string {
	return s.String()
}

func (s *SetResellerUserAlarmThresholdResponseBody) SetRequestId(v string) *SetResellerUserAlarmThresholdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetResellerUserAlarmThresholdResponseBody) SetCode(v string) *SetResellerUserAlarmThresholdResponseBody {
	s.Code = &v
	return s
}

func (s *SetResellerUserAlarmThresholdResponseBody) SetMessage(v string) *SetResellerUserAlarmThresholdResponseBody {
	s.Message = &v
	return s
}

func (s *SetResellerUserAlarmThresholdResponseBody) SetSuccess(v bool) *SetResellerUserAlarmThresholdResponseBody {
	s.Success = &v
	return s
}

func (s *SetResellerUserAlarmThresholdResponseBody) SetData(v bool) *SetResellerUserAlarmThresholdResponseBody {
	s.Data = &v
	return s
}

type SetResellerUserAlarmThresholdResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetResellerUserAlarmThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetResellerUserAlarmThresholdResponse) String() string {
	return tea.Prettify(s)
}

func (s SetResellerUserAlarmThresholdResponse) GoString() string {
	return s.String()
}

func (s *SetResellerUserAlarmThresholdResponse) SetHeaders(v map[string]*string) *SetResellerUserAlarmThresholdResponse {
	s.Headers = v
	return s
}

func (s *SetResellerUserAlarmThresholdResponse) SetBody(v *SetResellerUserAlarmThresholdResponseBody) *SetResellerUserAlarmThresholdResponse {
	s.Body = v
	return s
}

type SetResellerUserQuotaRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Amount   *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	OutBizId *string `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
}

func (s SetResellerUserQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s SetResellerUserQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetResellerUserQuotaRequest) SetOwnerId(v int64) *SetResellerUserQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *SetResellerUserQuotaRequest) SetAmount(v string) *SetResellerUserQuotaRequest {
	s.Amount = &v
	return s
}

func (s *SetResellerUserQuotaRequest) SetCurrency(v string) *SetResellerUserQuotaRequest {
	s.Currency = &v
	return s
}

func (s *SetResellerUserQuotaRequest) SetOutBizId(v string) *SetResellerUserQuotaRequest {
	s.OutBizId = &v
	return s
}

type SetResellerUserQuotaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s SetResellerUserQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetResellerUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetResellerUserQuotaResponseBody) SetRequestId(v string) *SetResellerUserQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetResellerUserQuotaResponseBody) SetCode(v string) *SetResellerUserQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *SetResellerUserQuotaResponseBody) SetMessage(v string) *SetResellerUserQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *SetResellerUserQuotaResponseBody) SetSuccess(v bool) *SetResellerUserQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *SetResellerUserQuotaResponseBody) SetData(v bool) *SetResellerUserQuotaResponseBody {
	s.Data = &v
	return s
}

type SetResellerUserQuotaResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetResellerUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetResellerUserQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s SetResellerUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *SetResellerUserQuotaResponse) SetHeaders(v map[string]*string) *SetResellerUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *SetResellerUserQuotaResponse) SetBody(v *SetResellerUserQuotaResponseBody) *SetResellerUserQuotaResponse {
	s.Body = v
	return s
}

type SetResellerUserStatusRequest struct {
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
}

func (s SetResellerUserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetResellerUserStatusRequest) GoString() string {
	return s.String()
}

func (s *SetResellerUserStatusRequest) SetOwnerId(v string) *SetResellerUserStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SetResellerUserStatusRequest) SetStatus(v string) *SetResellerUserStatusRequest {
	s.Status = &v
	return s
}

func (s *SetResellerUserStatusRequest) SetBusinessType(v string) *SetResellerUserStatusRequest {
	s.BusinessType = &v
	return s
}

type SetResellerUserStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s SetResellerUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetResellerUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetResellerUserStatusResponseBody) SetRequestId(v string) *SetResellerUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetResellerUserStatusResponseBody) SetCode(v string) *SetResellerUserStatusResponseBody {
	s.Code = &v
	return s
}

func (s *SetResellerUserStatusResponseBody) SetMessage(v string) *SetResellerUserStatusResponseBody {
	s.Message = &v
	return s
}

func (s *SetResellerUserStatusResponseBody) SetSuccess(v bool) *SetResellerUserStatusResponseBody {
	s.Success = &v
	return s
}

func (s *SetResellerUserStatusResponseBody) SetData(v bool) *SetResellerUserStatusResponseBody {
	s.Data = &v
	return s
}

type SetResellerUserStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetResellerUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetResellerUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetResellerUserStatusResponse) GoString() string {
	return s.String()
}

func (s *SetResellerUserStatusResponse) SetHeaders(v map[string]*string) *SetResellerUserStatusResponse {
	s.Headers = v
	return s
}

func (s *SetResellerUserStatusResponse) SetBody(v *SetResellerUserStatusResponseBody) *SetResellerUserStatusResponse {
	s.Body = v
	return s
}

type SubscribeBillToOSSRequest struct {
	SubscribeBucket         *string `json:"SubscribeBucket,omitempty" xml:"SubscribeBucket,omitempty"`
	SubscribeType           *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
	MultAccountRelSubscribe *string `json:"MultAccountRelSubscribe,omitempty" xml:"MultAccountRelSubscribe,omitempty"`
	BucketOwnerId           *int64  `json:"BucketOwnerId,omitempty" xml:"BucketOwnerId,omitempty"`
}

func (s SubscribeBillToOSSRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscribeBillToOSSRequest) GoString() string {
	return s.String()
}

func (s *SubscribeBillToOSSRequest) SetSubscribeBucket(v string) *SubscribeBillToOSSRequest {
	s.SubscribeBucket = &v
	return s
}

func (s *SubscribeBillToOSSRequest) SetSubscribeType(v string) *SubscribeBillToOSSRequest {
	s.SubscribeType = &v
	return s
}

func (s *SubscribeBillToOSSRequest) SetMultAccountRelSubscribe(v string) *SubscribeBillToOSSRequest {
	s.MultAccountRelSubscribe = &v
	return s
}

func (s *SubscribeBillToOSSRequest) SetBucketOwnerId(v int64) *SubscribeBillToOSSRequest {
	s.BucketOwnerId = &v
	return s
}

type SubscribeBillToOSSResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SubscribeBillToOSSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubscribeBillToOSSResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeBillToOSSResponseBody) SetRequestId(v string) *SubscribeBillToOSSResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscribeBillToOSSResponseBody) SetSuccess(v bool) *SubscribeBillToOSSResponseBody {
	s.Success = &v
	return s
}

func (s *SubscribeBillToOSSResponseBody) SetCode(v string) *SubscribeBillToOSSResponseBody {
	s.Code = &v
	return s
}

func (s *SubscribeBillToOSSResponseBody) SetMessage(v string) *SubscribeBillToOSSResponseBody {
	s.Message = &v
	return s
}

type SubscribeBillToOSSResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubscribeBillToOSSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubscribeBillToOSSResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscribeBillToOSSResponse) GoString() string {
	return s.String()
}

func (s *SubscribeBillToOSSResponse) SetHeaders(v map[string]*string) *SubscribeBillToOSSResponse {
	s.Headers = v
	return s
}

func (s *SubscribeBillToOSSResponse) SetBody(v *SubscribeBillToOSSResponseBody) *SubscribeBillToOSSResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetCode(v string) *TagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *TagResourcesResponseBody) SetMessage(v string) *TagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *TagResourcesResponseBody) SetData(v bool) *TagResourcesResponseBody {
	s.Data = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnsubscribeBillToOSSRequest struct {
	SubscribeType           *string `json:"SubscribeType,omitempty" xml:"SubscribeType,omitempty"`
	MultAccountRelSubscribe *string `json:"MultAccountRelSubscribe,omitempty" xml:"MultAccountRelSubscribe,omitempty"`
}

func (s UnsubscribeBillToOSSRequest) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeBillToOSSRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeBillToOSSRequest) SetSubscribeType(v string) *UnsubscribeBillToOSSRequest {
	s.SubscribeType = &v
	return s
}

func (s *UnsubscribeBillToOSSRequest) SetMultAccountRelSubscribe(v string) *UnsubscribeBillToOSSRequest {
	s.MultAccountRelSubscribe = &v
	return s
}

type UnsubscribeBillToOSSResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UnsubscribeBillToOSSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeBillToOSSResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeBillToOSSResponseBody) SetRequestId(v string) *UnsubscribeBillToOSSResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnsubscribeBillToOSSResponseBody) SetSuccess(v bool) *UnsubscribeBillToOSSResponseBody {
	s.Success = &v
	return s
}

func (s *UnsubscribeBillToOSSResponseBody) SetCode(v string) *UnsubscribeBillToOSSResponseBody {
	s.Code = &v
	return s
}

func (s *UnsubscribeBillToOSSResponseBody) SetMessage(v string) *UnsubscribeBillToOSSResponseBody {
	s.Message = &v
	return s
}

type UnsubscribeBillToOSSResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnsubscribeBillToOSSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnsubscribeBillToOSSResponse) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeBillToOSSResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeBillToOSSResponse) SetHeaders(v map[string]*string) *UnsubscribeBillToOSSResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeBillToOSSResponse) SetBody(v *UnsubscribeBillToOSSResponseBody) *UnsubscribeBillToOSSResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

type UntagResourcesResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetCode(v string) *UntagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *UntagResourcesResponseBody) SetMessage(v string) *UntagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *UntagResourcesResponseBody) SetData(v bool) *UntagResourcesResponseBody {
	s.Data = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpgradeResourcePackageRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	EffectiveDate *string `json:"EffectiveDate,omitempty" xml:"EffectiveDate,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s UpgradeResourcePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *UpgradeResourcePackageRequest) SetOwnerId(v int64) *UpgradeResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeResourcePackageRequest) SetInstanceId(v string) *UpgradeResourcePackageRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeResourcePackageRequest) SetEffectiveDate(v string) *UpgradeResourcePackageRequest {
	s.EffectiveDate = &v
	return s
}

func (s *UpgradeResourcePackageRequest) SetSpecification(v string) *UpgradeResourcePackageRequest {
	s.Specification = &v
	return s
}

type UpgradeResourcePackageResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *int64                                  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *UpgradeResourcePackageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s UpgradeResourcePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeResourcePackageResponseBody) SetRequestId(v string) *UpgradeResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeResourcePackageResponseBody) SetOrderId(v int64) *UpgradeResourcePackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeResourcePackageResponseBody) SetSuccess(v bool) *UpgradeResourcePackageResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeResourcePackageResponseBody) SetCode(v string) *UpgradeResourcePackageResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeResourcePackageResponseBody) SetMessage(v string) *UpgradeResourcePackageResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeResourcePackageResponseBody) SetData(v *UpgradeResourcePackageResponseBodyData) *UpgradeResourcePackageResponseBody {
	s.Data = v
	return s
}

type UpgradeResourcePackageResponseBodyData struct {
	OrderId    *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpgradeResourcePackageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpgradeResourcePackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeResourcePackageResponseBodyData) SetOrderId(v int64) *UpgradeResourcePackageResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *UpgradeResourcePackageResponseBodyData) SetInstanceId(v string) *UpgradeResourcePackageResponseBodyData {
	s.InstanceId = &v
	return s
}

type UpgradeResourcePackageResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeResourcePackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeResourcePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeResourcePackageResponse) GoString() string {
	return s.String()
}

func (s *UpgradeResourcePackageResponse) SetHeaders(v map[string]*string) *UpgradeResourcePackageResponse {
	s.Headers = v
	return s
}

func (s *UpgradeResourcePackageResponse) SetBody(v *UpgradeResourcePackageResponseBody) *UpgradeResourcePackageResponse {
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
		"ap-northeast-1":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":                  tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-1":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"cn-beijing":                  tea.String("business.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("business.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("business.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("business.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("business.aliyuncs.com"),
		"cn-chengdu":                  tea.String("business.aliyuncs.com"),
		"cn-edge-1":                   tea.String("business.aliyuncs.com"),
		"cn-fujian":                   tea.String("business.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("business.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("business.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("business.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("business.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("business.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("business.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("business.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("business.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("business.aliyuncs.com"),
		"cn-hongkong":                 tea.String("business.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("business.aliyuncs.com"),
		"cn-huhehaote":                tea.String("business.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("business.aliyuncs.com"),
		"cn-qingdao":                  tea.String("business.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("business.aliyuncs.com"),
		"cn-shanghai":                 tea.String("business.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("business.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("business.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("business.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("business.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("business.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("business.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("business.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("business.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("business.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("business.aliyuncs.com"),
		"cn-wuhan":                    tea.String("business.aliyuncs.com"),
		"cn-yushanfang":               tea.String("business.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("business.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("business.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("business.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("business.aliyuncs.com"),
		"eu-central-1":                tea.String("business.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":                   tea.String("business.ap-southeast-1.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("business.ap-southeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("business.ap-southeast-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("business.ap-southeast-1.aliyuncs.com"),
		"us-east-1":                   tea.String("business.ap-southeast-1.aliyuncs.com"),
		"us-west-1":                   tea.String("business.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("bssopenapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AllocateCostUnitResourceWithOptions(request *AllocateCostUnitResourceRequest, runtime *util.RuntimeOptions) (_result *AllocateCostUnitResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AllocateCostUnitResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AllocateCostUnitResource"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateCostUnitResource(request *AllocateCostUnitResourceRequest) (_result *AllocateCostUnitResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateCostUnitResourceResponse{}
	_body, _err := client.AllocateCostUnitResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyInvoiceWithOptions(request *ApplyInvoiceRequest, runtime *util.RuntimeOptions) (_result *ApplyInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApplyInvoiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplyInvoice"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyInvoice(request *ApplyInvoiceRequest) (_result *ApplyInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyInvoiceResponse{}
	_body, _err := client.ApplyInvoiceWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CancelOrder"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ChangeResellerConsumeAmountWithOptions(request *ChangeResellerConsumeAmountRequest, runtime *util.RuntimeOptions) (_result *ChangeResellerConsumeAmountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeResellerConsumeAmountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeResellerConsumeAmount"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeResellerConsumeAmount(request *ChangeResellerConsumeAmountRequest) (_result *ChangeResellerConsumeAmountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResellerConsumeAmountResponse{}
	_body, _err := client.ChangeResellerConsumeAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConvertChargeTypeWithOptions(request *ConvertChargeTypeRequest, runtime *util.RuntimeOptions) (_result *ConvertChargeTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConvertChargeTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConvertChargeType"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConvertChargeType(request *ConvertChargeTypeRequest) (_result *ConvertChargeTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertChargeTypeResponse{}
	_body, _err := client.ConvertChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAgAccountWithOptions(request *CreateAgAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAgAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAgAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAgAccount"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAgAccount(request *CreateAgAccountRequest) (_result *CreateAgAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAgAccountResponse{}
	_body, _err := client.CreateAgAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCostUnitWithOptions(request *CreateCostUnitRequest, runtime *util.RuntimeOptions) (_result *CreateCostUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCostUnitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCostUnit"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCostUnit(request *CreateCostUnitRequest) (_result *CreateCostUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCostUnitResponse{}
	_body, _err := client.CreateCostUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateInstance"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResellerUserQuotaWithOptions(request *CreateResellerUserQuotaRequest, runtime *util.RuntimeOptions) (_result *CreateResellerUserQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateResellerUserQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateResellerUserQuota"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateResellerUserQuota(request *CreateResellerUserQuotaRequest) (_result *CreateResellerUserQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResellerUserQuotaResponse{}
	_body, _err := client.CreateResellerUserQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourcePackageWithOptions(request *CreateResourcePackageRequest, runtime *util.RuntimeOptions) (_result *CreateResourcePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateResourcePackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateResourcePackage"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateResourcePackage(request *CreateResourcePackageRequest) (_result *CreateResourcePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourcePackageResponse{}
	_body, _err := client.CreateResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCostUnitWithOptions(request *DeleteCostUnitRequest, runtime *util.RuntimeOptions) (_result *DeleteCostUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCostUnitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCostUnit"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCostUnit(request *DeleteCostUnitRequest) (_result *DeleteCostUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCostUnitResponse{}
	_body, _err := client.DeleteCostUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePricingModuleWithOptions(request *DescribePricingModuleRequest, runtime *util.RuntimeOptions) (_result *DescribePricingModuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePricingModuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePricingModule"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePricingModule(request *DescribePricingModuleRequest) (_result *DescribePricingModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePricingModuleResponse{}
	_body, _err := client.DescribePricingModuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourcePackageProductWithOptions(request *DescribeResourcePackageProductRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackageProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeResourcePackageProductResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeResourcePackageProduct"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourcePackageProduct(request *DescribeResourcePackageProductRequest) (_result *DescribeResourcePackageProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourcePackageProductResponse{}
	_body, _err := client.DescribeResourcePackageProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSplitItemBillWithOptions(request *DescribeSplitItemBillRequest, runtime *util.RuntimeOptions) (_result *DescribeSplitItemBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSplitItemBillResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSplitItemBill"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSplitItemBill(request *DescribeSplitItemBillRequest) (_result *DescribeSplitItemBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSplitItemBillResponse{}
	_body, _err := client.DescribeSplitItemBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableBillGenerationWithOptions(request *EnableBillGenerationRequest, runtime *util.RuntimeOptions) (_result *EnableBillGenerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableBillGenerationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableBillGeneration"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableBillGeneration(request *EnableBillGenerationRequest) (_result *EnableBillGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableBillGenerationResponse{}
	_body, _err := client.EnableBillGenerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCustomerAccountInfoWithOptions(request *GetCustomerAccountInfoRequest, runtime *util.RuntimeOptions) (_result *GetCustomerAccountInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCustomerAccountInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCustomerAccountInfo"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCustomerAccountInfo(request *GetCustomerAccountInfoRequest) (_result *GetCustomerAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomerAccountInfoResponse{}
	_body, _err := client.GetCustomerAccountInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCustomerListWithOptions(runtime *util.RuntimeOptions) (_result *GetCustomerListResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &GetCustomerListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCustomerList"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCustomerList() (_result *GetCustomerListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomerListResponse{}
	_body, _err := client.GetCustomerListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrderDetailWithOptions(request *GetOrderDetailRequest, runtime *util.RuntimeOptions) (_result *GetOrderDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOrderDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOrderDetail"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrderDetail(request *GetOrderDetailRequest) (_result *GetOrderDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOrderDetailResponse{}
	_body, _err := client.GetOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPayAsYouGoPriceWithOptions(request *GetPayAsYouGoPriceRequest, runtime *util.RuntimeOptions) (_result *GetPayAsYouGoPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPayAsYouGoPriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPayAsYouGoPrice"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPayAsYouGoPrice(request *GetPayAsYouGoPriceRequest) (_result *GetPayAsYouGoPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPayAsYouGoPriceResponse{}
	_body, _err := client.GetPayAsYouGoPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourcePackagePriceWithOptions(request *GetResourcePackagePriceRequest, runtime *util.RuntimeOptions) (_result *GetResourcePackagePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetResourcePackagePriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetResourcePackagePrice"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourcePackagePrice(request *GetResourcePackagePriceRequest) (_result *GetResourcePackagePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourcePackagePriceResponse{}
	_body, _err := client.GetResourcePackagePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSubscriptionPriceWithOptions(request *GetSubscriptionPriceRequest, runtime *util.RuntimeOptions) (_result *GetSubscriptionPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSubscriptionPriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSubscriptionPrice"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSubscriptionPrice(request *GetSubscriptionPriceRequest) (_result *GetSubscriptionPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubscriptionPriceResponse{}
	_body, _err := client.GetSubscriptionPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyCostUnitWithOptions(request *ModifyCostUnitRequest, runtime *util.RuntimeOptions) (_result *ModifyCostUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyCostUnitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyCostUnit"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCostUnit(request *ModifyCostUnitRequest) (_result *ModifyCostUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyCostUnitResponse{}
	_body, _err := client.ModifyCostUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceWithOptions(request *ModifyInstanceRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstance"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstance(request *ModifyInstanceRequest) (_result *ModifyInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.ModifyInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAccountBalanceWithOptions(runtime *util.RuntimeOptions) (_result *QueryAccountBalanceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &QueryAccountBalanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAccountBalance"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAccountBalance() (_result *QueryAccountBalanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAccountBalanceResponse{}
	_body, _err := client.QueryAccountBalanceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAccountBillWithOptions(request *QueryAccountBillRequest, runtime *util.RuntimeOptions) (_result *QueryAccountBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAccountBillResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAccountBill"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAccountBill(request *QueryAccountBillRequest) (_result *QueryAccountBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAccountBillResponse{}
	_body, _err := client.QueryAccountBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAccountTransactionDetailsWithOptions(request *QueryAccountTransactionDetailsRequest, runtime *util.RuntimeOptions) (_result *QueryAccountTransactionDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAccountTransactionDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAccountTransactionDetails"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAccountTransactionDetails(request *QueryAccountTransactionDetailsRequest) (_result *QueryAccountTransactionDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAccountTransactionDetailsResponse{}
	_body, _err := client.QueryAccountTransactionDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAccountTransactionsWithOptions(request *QueryAccountTransactionsRequest, runtime *util.RuntimeOptions) (_result *QueryAccountTransactionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAccountTransactionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAccountTransactions"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAccountTransactions(request *QueryAccountTransactionsRequest) (_result *QueryAccountTransactionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAccountTransactionsResponse{}
	_body, _err := client.QueryAccountTransactionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAvailableInstancesWithOptions(request *QueryAvailableInstancesRequest, runtime *util.RuntimeOptions) (_result *QueryAvailableInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryAvailableInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryAvailableInstances"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAvailableInstances(request *QueryAvailableInstancesRequest) (_result *QueryAvailableInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAvailableInstancesResponse{}
	_body, _err := client.QueryAvailableInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBillWithOptions(request *QueryBillRequest, runtime *util.RuntimeOptions) (_result *QueryBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryBillResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryBill"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBill(request *QueryBillRequest) (_result *QueryBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBillResponse{}
	_body, _err := client.QueryBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBillOverviewWithOptions(request *QueryBillOverviewRequest, runtime *util.RuntimeOptions) (_result *QueryBillOverviewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryBillOverviewResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryBillOverview"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBillOverview(request *QueryBillOverviewRequest) (_result *QueryBillOverviewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBillOverviewResponse{}
	_body, _err := client.QueryBillOverviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBillToOSSSubscriptionWithOptions(runtime *util.RuntimeOptions) (_result *QueryBillToOSSSubscriptionResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &QueryBillToOSSSubscriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryBillToOSSSubscription"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBillToOSSSubscription() (_result *QueryBillToOSSSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBillToOSSSubscriptionResponse{}
	_body, _err := client.QueryBillToOSSSubscriptionWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCashCouponsWithOptions(request *QueryCashCouponsRequest, runtime *util.RuntimeOptions) (_result *QueryCashCouponsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCashCouponsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCashCoupons"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCashCoupons(request *QueryCashCouponsRequest) (_result *QueryCashCouponsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCashCouponsResponse{}
	_body, _err := client.QueryCashCouponsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCostUnitWithOptions(request *QueryCostUnitRequest, runtime *util.RuntimeOptions) (_result *QueryCostUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCostUnitResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCostUnit"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCostUnit(request *QueryCostUnitRequest) (_result *QueryCostUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCostUnitResponse{}
	_body, _err := client.QueryCostUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCostUnitResourceWithOptions(request *QueryCostUnitResourceRequest, runtime *util.RuntimeOptions) (_result *QueryCostUnitResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCostUnitResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCostUnitResource"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCostUnitResource(request *QueryCostUnitResourceRequest) (_result *QueryCostUnitResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCostUnitResourceResponse{}
	_body, _err := client.QueryCostUnitResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCustomerAddressListWithOptions(request *QueryCustomerAddressListRequest, runtime *util.RuntimeOptions) (_result *QueryCustomerAddressListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCustomerAddressListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCustomerAddressList"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCustomerAddressList(request *QueryCustomerAddressListRequest) (_result *QueryCustomerAddressListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCustomerAddressListResponse{}
	_body, _err := client.QueryCustomerAddressListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEvaluateListWithOptions(request *QueryEvaluateListRequest, runtime *util.RuntimeOptions) (_result *QueryEvaluateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryEvaluateListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryEvaluateList"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEvaluateList(request *QueryEvaluateListRequest) (_result *QueryEvaluateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEvaluateListResponse{}
	_body, _err := client.QueryEvaluateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFinancialAccountInfoWithOptions(request *QueryFinancialAccountInfoRequest, runtime *util.RuntimeOptions) (_result *QueryFinancialAccountInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryFinancialAccountInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryFinancialAccountInfo"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFinancialAccountInfo(request *QueryFinancialAccountInfoRequest) (_result *QueryFinancialAccountInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFinancialAccountInfoResponse{}
	_body, _err := client.QueryFinancialAccountInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryInstanceBillWithOptions(request *QueryInstanceBillRequest, runtime *util.RuntimeOptions) (_result *QueryInstanceBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryInstanceBillResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryInstanceBill"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryInstanceBill(request *QueryInstanceBillRequest) (_result *QueryInstanceBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInstanceBillResponse{}
	_body, _err := client.QueryInstanceBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryInstanceByTagWithOptions(request *QueryInstanceByTagRequest, runtime *util.RuntimeOptions) (_result *QueryInstanceByTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryInstanceByTagResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryInstanceByTag"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryInstanceByTag(request *QueryInstanceByTagRequest) (_result *QueryInstanceByTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInstanceByTagResponse{}
	_body, _err := client.QueryInstanceByTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryInstanceGaapCostWithOptions(request *QueryInstanceGaapCostRequest, runtime *util.RuntimeOptions) (_result *QueryInstanceGaapCostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryInstanceGaapCostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryInstanceGaapCost"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryInstanceGaapCost(request *QueryInstanceGaapCostRequest) (_result *QueryInstanceGaapCostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInstanceGaapCostResponse{}
	_body, _err := client.QueryInstanceGaapCostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryInvoicingCustomerListWithOptions(request *QueryInvoicingCustomerListRequest, runtime *util.RuntimeOptions) (_result *QueryInvoicingCustomerListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryInvoicingCustomerListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryInvoicingCustomerList"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryInvoicingCustomerList(request *QueryInvoicingCustomerListRequest) (_result *QueryInvoicingCustomerListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInvoicingCustomerListResponse{}
	_body, _err := client.QueryInvoicingCustomerListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMonthlyBillWithOptions(request *QueryMonthlyBillRequest, runtime *util.RuntimeOptions) (_result *QueryMonthlyBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMonthlyBillResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMonthlyBill"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMonthlyBill(request *QueryMonthlyBillRequest) (_result *QueryMonthlyBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMonthlyBillResponse{}
	_body, _err := client.QueryMonthlyBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMonthlyInstanceConsumptionWithOptions(request *QueryMonthlyInstanceConsumptionRequest, runtime *util.RuntimeOptions) (_result *QueryMonthlyInstanceConsumptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryMonthlyInstanceConsumptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryMonthlyInstanceConsumption"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMonthlyInstanceConsumption(request *QueryMonthlyInstanceConsumptionRequest) (_result *QueryMonthlyInstanceConsumptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMonthlyInstanceConsumptionResponse{}
	_body, _err := client.QueryMonthlyInstanceConsumptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrdersWithOptions(request *QueryOrdersRequest, runtime *util.RuntimeOptions) (_result *QueryOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryOrdersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryOrders"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrders(request *QueryOrdersRequest) (_result *QueryOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrdersResponse{}
	_body, _err := client.QueryOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPermissionListWithOptions(request *QueryPermissionListRequest, runtime *util.RuntimeOptions) (_result *QueryPermissionListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryPermissionListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryPermissionList"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPermissionList(request *QueryPermissionListRequest) (_result *QueryPermissionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPermissionListResponse{}
	_body, _err := client.QueryPermissionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPrepaidCardsWithOptions(request *QueryPrepaidCardsRequest, runtime *util.RuntimeOptions) (_result *QueryPrepaidCardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryPrepaidCardsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryPrepaidCards"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPrepaidCards(request *QueryPrepaidCardsRequest) (_result *QueryPrepaidCardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPrepaidCardsResponse{}
	_body, _err := client.QueryPrepaidCardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryProductListWithOptions(request *QueryProductListRequest, runtime *util.RuntimeOptions) (_result *QueryProductListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryProductListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryProductList"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryProductList(request *QueryProductListRequest) (_result *QueryProductListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryProductListResponse{}
	_body, _err := client.QueryProductListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRedeemWithOptions(request *QueryRedeemRequest, runtime *util.RuntimeOptions) (_result *QueryRedeemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &QueryRedeemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRedeem"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRedeem(request *QueryRedeemRequest) (_result *QueryRedeemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRedeemResponse{}
	_body, _err := client.QueryRedeemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRelationListWithOptions(request *QueryRelationListRequest, runtime *util.RuntimeOptions) (_result *QueryRelationListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRelationListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRelationList"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRelationList(request *QueryRelationListRequest) (_result *QueryRelationListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRelationListResponse{}
	_body, _err := client.QueryRelationListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryResellerAvailableQuotaWithOptions(request *QueryResellerAvailableQuotaRequest, runtime *util.RuntimeOptions) (_result *QueryResellerAvailableQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryResellerAvailableQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryResellerAvailableQuota"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryResellerAvailableQuota(request *QueryResellerAvailableQuotaRequest) (_result *QueryResellerAvailableQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryResellerAvailableQuotaResponse{}
	_body, _err := client.QueryResellerAvailableQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryResourcePackageInstancesWithOptions(request *QueryResourcePackageInstancesRequest, runtime *util.RuntimeOptions) (_result *QueryResourcePackageInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryResourcePackageInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryResourcePackageInstances"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryResourcePackageInstances(request *QueryResourcePackageInstancesRequest) (_result *QueryResourcePackageInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryResourcePackageInstancesResponse{}
	_body, _err := client.QueryResourcePackageInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRIUtilizationDetailWithOptions(request *QueryRIUtilizationDetailRequest, runtime *util.RuntimeOptions) (_result *QueryRIUtilizationDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRIUtilizationDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRIUtilizationDetail"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRIUtilizationDetail(request *QueryRIUtilizationDetailRequest) (_result *QueryRIUtilizationDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRIUtilizationDetailResponse{}
	_body, _err := client.QueryRIUtilizationDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySavingsPlansDeductLogWithOptions(request *QuerySavingsPlansDeductLogRequest, runtime *util.RuntimeOptions) (_result *QuerySavingsPlansDeductLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySavingsPlansDeductLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySavingsPlansDeductLog"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySavingsPlansDeductLog(request *QuerySavingsPlansDeductLogRequest) (_result *QuerySavingsPlansDeductLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySavingsPlansDeductLogResponse{}
	_body, _err := client.QuerySavingsPlansDeductLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySavingsPlansInstanceWithOptions(request *QuerySavingsPlansInstanceRequest, runtime *util.RuntimeOptions) (_result *QuerySavingsPlansInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySavingsPlansInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySavingsPlansInstance"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySavingsPlansInstance(request *QuerySavingsPlansInstanceRequest) (_result *QuerySavingsPlansInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySavingsPlansInstanceResponse{}
	_body, _err := client.QuerySavingsPlansInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySettleBillWithOptions(request *QuerySettleBillRequest, runtime *util.RuntimeOptions) (_result *QuerySettleBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySettleBillResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySettleBill"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySettleBill(request *QuerySettleBillRequest) (_result *QuerySettleBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySettleBillResponse{}
	_body, _err := client.QuerySettleBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySettlementBillWithOptions(request *QuerySettlementBillRequest, runtime *util.RuntimeOptions) (_result *QuerySettlementBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySettlementBillResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySettlementBill"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySettlementBill(request *QuerySettlementBillRequest) (_result *QuerySettlementBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySettlementBillResponse{}
	_body, _err := client.QuerySettlementBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySplitItemBillWithOptions(request *QuerySplitItemBillRequest, runtime *util.RuntimeOptions) (_result *QuerySplitItemBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySplitItemBillResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySplitItemBill"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySplitItemBill(request *QuerySplitItemBillRequest) (_result *QuerySplitItemBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySplitItemBillResponse{}
	_body, _err := client.QuerySplitItemBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryUserOmsDataWithOptions(request *QueryUserOmsDataRequest, runtime *util.RuntimeOptions) (_result *QueryUserOmsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryUserOmsDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryUserOmsData"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryUserOmsData(request *QueryUserOmsDataRequest) (_result *QueryUserOmsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryUserOmsDataResponse{}
	_body, _err := client.QueryUserOmsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenewInstance"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewResourcePackageWithOptions(request *RenewResourcePackageRequest, runtime *util.RuntimeOptions) (_result *RenewResourcePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenewResourcePackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenewResourcePackage"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewResourcePackage(request *RenewResourcePackageRequest) (_result *RenewResourcePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewResourcePackageResponse{}
	_body, _err := client.RenewResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveUserCreditWithOptions(request *SaveUserCreditRequest, runtime *util.RuntimeOptions) (_result *SaveUserCreditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveUserCreditResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveUserCredit"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveUserCredit(request *SaveUserCreditRequest) (_result *SaveUserCreditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveUserCreditResponse{}
	_body, _err := client.SaveUserCreditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCreditLabelActionWithOptions(request *SetCreditLabelActionRequest, runtime *util.RuntimeOptions) (_result *SetCreditLabelActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetCreditLabelActionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetCreditLabelAction"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCreditLabelAction(request *SetCreditLabelActionRequest) (_result *SetCreditLabelActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCreditLabelActionResponse{}
	_body, _err := client.SetCreditLabelActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetRenewalWithOptions(request *SetRenewalRequest, runtime *util.RuntimeOptions) (_result *SetRenewalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetRenewalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetRenewal"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetRenewal(request *SetRenewalRequest) (_result *SetRenewalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetRenewalResponse{}
	_body, _err := client.SetRenewalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetResellerUserAlarmThresholdWithOptions(request *SetResellerUserAlarmThresholdRequest, runtime *util.RuntimeOptions) (_result *SetResellerUserAlarmThresholdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetResellerUserAlarmThresholdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetResellerUserAlarmThreshold"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetResellerUserAlarmThreshold(request *SetResellerUserAlarmThresholdRequest) (_result *SetResellerUserAlarmThresholdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetResellerUserAlarmThresholdResponse{}
	_body, _err := client.SetResellerUserAlarmThresholdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetResellerUserQuotaWithOptions(request *SetResellerUserQuotaRequest, runtime *util.RuntimeOptions) (_result *SetResellerUserQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetResellerUserQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetResellerUserQuota"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetResellerUserQuota(request *SetResellerUserQuotaRequest) (_result *SetResellerUserQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetResellerUserQuotaResponse{}
	_body, _err := client.SetResellerUserQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetResellerUserStatusWithOptions(request *SetResellerUserStatusRequest, runtime *util.RuntimeOptions) (_result *SetResellerUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetResellerUserStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetResellerUserStatus"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetResellerUserStatus(request *SetResellerUserStatusRequest) (_result *SetResellerUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetResellerUserStatusResponse{}
	_body, _err := client.SetResellerUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubscribeBillToOSSWithOptions(request *SubscribeBillToOSSRequest, runtime *util.RuntimeOptions) (_result *SubscribeBillToOSSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubscribeBillToOSSResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubscribeBillToOSS"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubscribeBillToOSS(request *SubscribeBillToOSSRequest) (_result *SubscribeBillToOSSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubscribeBillToOSSResponse{}
	_body, _err := client.SubscribeBillToOSSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnsubscribeBillToOSSWithOptions(request *UnsubscribeBillToOSSRequest, runtime *util.RuntimeOptions) (_result *UnsubscribeBillToOSSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnsubscribeBillToOSSResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnsubscribeBillToOSS"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnsubscribeBillToOSS(request *UnsubscribeBillToOSSRequest) (_result *UnsubscribeBillToOSSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnsubscribeBillToOSSResponse{}
	_body, _err := client.UnsubscribeBillToOSSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeResourcePackageWithOptions(request *UpgradeResourcePackageRequest, runtime *util.RuntimeOptions) (_result *UpgradeResourcePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeResourcePackageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeResourcePackage"), tea.String("2017-12-14"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeResourcePackage(request *UpgradeResourcePackageRequest) (_result *UpgradeResourcePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeResourcePackageResponse{}
	_body, _err := client.UpgradeResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
