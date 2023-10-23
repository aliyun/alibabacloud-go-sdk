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

type BindStorageOrderRequest struct {
	DeviceName            *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EnableDefaultPlan     *bool   `json:"EnableDefaultPlan,omitempty" xml:"EnableDefaultPlan,omitempty"`
	EventRecordDuration   *int32  `json:"EventRecordDuration,omitempty" xml:"EventRecordDuration,omitempty"`
	EventRecordProlong    *bool   `json:"EventRecordProlong,omitempty" xml:"EventRecordProlong,omitempty"`
	IotId                 *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	MaxRecordFileDuration *int32  `json:"MaxRecordFileDuration,omitempty" xml:"MaxRecordFileDuration,omitempty"`
	OrderId               *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	PreRecordDuration     *int32  `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	ProductKey            *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	UserId                *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName              *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s BindStorageOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s BindStorageOrderRequest) GoString() string {
	return s.String()
}

func (s *BindStorageOrderRequest) SetDeviceName(v string) *BindStorageOrderRequest {
	s.DeviceName = &v
	return s
}

func (s *BindStorageOrderRequest) SetEnableDefaultPlan(v bool) *BindStorageOrderRequest {
	s.EnableDefaultPlan = &v
	return s
}

func (s *BindStorageOrderRequest) SetEventRecordDuration(v int32) *BindStorageOrderRequest {
	s.EventRecordDuration = &v
	return s
}

func (s *BindStorageOrderRequest) SetEventRecordProlong(v bool) *BindStorageOrderRequest {
	s.EventRecordProlong = &v
	return s
}

func (s *BindStorageOrderRequest) SetIotId(v string) *BindStorageOrderRequest {
	s.IotId = &v
	return s
}

func (s *BindStorageOrderRequest) SetMaxRecordFileDuration(v int32) *BindStorageOrderRequest {
	s.MaxRecordFileDuration = &v
	return s
}

func (s *BindStorageOrderRequest) SetOrderId(v string) *BindStorageOrderRequest {
	s.OrderId = &v
	return s
}

func (s *BindStorageOrderRequest) SetPreRecordDuration(v int32) *BindStorageOrderRequest {
	s.PreRecordDuration = &v
	return s
}

func (s *BindStorageOrderRequest) SetProductKey(v string) *BindStorageOrderRequest {
	s.ProductKey = &v
	return s
}

func (s *BindStorageOrderRequest) SetUserId(v string) *BindStorageOrderRequest {
	s.UserId = &v
	return s
}

func (s *BindStorageOrderRequest) SetUserName(v string) *BindStorageOrderRequest {
	s.UserName = &v
	return s
}

type BindStorageOrderResponseBody struct {
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *BindStorageOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindStorageOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindStorageOrderResponseBody) GoString() string {
	return s.String()
}

func (s *BindStorageOrderResponseBody) SetCode(v string) *BindStorageOrderResponseBody {
	s.Code = &v
	return s
}

func (s *BindStorageOrderResponseBody) SetData(v *BindStorageOrderResponseBodyData) *BindStorageOrderResponseBody {
	s.Data = v
	return s
}

func (s *BindStorageOrderResponseBody) SetErrorMessage(v string) *BindStorageOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BindStorageOrderResponseBody) SetRequestId(v string) *BindStorageOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindStorageOrderResponseBody) SetSuccess(v bool) *BindStorageOrderResponseBody {
	s.Success = &v
	return s
}

type BindStorageOrderResponseBodyData struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Copies        *int32  `json:"Copies,omitempty" xml:"Copies,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC    *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	IdentityId    *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderType     *int32  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OutOrderNo    *string `json:"OutOrderNo,omitempty" xml:"OutOrderNo,omitempty"`
	PaymentStatus *int32  `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	PreConsume    *int32  `json:"PreConsume,omitempty" xml:"PreConsume,omitempty"`
	Price         *string `json:"Price,omitempty" xml:"Price,omitempty"`
	RecordType    *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC  *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s BindStorageOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BindStorageOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindStorageOrderResponseBodyData) SetCommodityCode(v string) *BindStorageOrderResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetCopies(v int32) *BindStorageOrderResponseBodyData {
	s.Copies = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetEndTime(v string) *BindStorageOrderResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetEndTimeUTC(v string) *BindStorageOrderResponseBodyData {
	s.EndTimeUTC = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetIdentityId(v string) *BindStorageOrderResponseBodyData {
	s.IdentityId = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetIotId(v string) *BindStorageOrderResponseBodyData {
	s.IotId = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetOrderId(v string) *BindStorageOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetOrderType(v int32) *BindStorageOrderResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetOutOrderNo(v string) *BindStorageOrderResponseBodyData {
	s.OutOrderNo = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetPaymentStatus(v int32) *BindStorageOrderResponseBodyData {
	s.PaymentStatus = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetPreConsume(v int32) *BindStorageOrderResponseBodyData {
	s.PreConsume = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetPrice(v string) *BindStorageOrderResponseBodyData {
	s.Price = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetRecordType(v int32) *BindStorageOrderResponseBodyData {
	s.RecordType = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetSpecification(v string) *BindStorageOrderResponseBodyData {
	s.Specification = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetStartTime(v string) *BindStorageOrderResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetStartTimeUTC(v string) *BindStorageOrderResponseBodyData {
	s.StartTimeUTC = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetStatus(v int32) *BindStorageOrderResponseBodyData {
	s.Status = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetUserId(v string) *BindStorageOrderResponseBodyData {
	s.UserId = &v
	return s
}

func (s *BindStorageOrderResponseBodyData) SetUserName(v string) *BindStorageOrderResponseBodyData {
	s.UserName = &v
	return s
}

type BindStorageOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindStorageOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindStorageOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s BindStorageOrderResponse) GoString() string {
	return s.String()
}

func (s *BindStorageOrderResponse) SetHeaders(v map[string]*string) *BindStorageOrderResponse {
	s.Headers = v
	return s
}

func (s *BindStorageOrderResponse) SetStatusCode(v int32) *BindStorageOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *BindStorageOrderResponse) SetBody(v *BindStorageOrderResponseBody) *BindStorageOrderResponse {
	s.Body = v
	return s
}

type CheckFreeStorageValidRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s CheckFreeStorageValidRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckFreeStorageValidRequest) GoString() string {
	return s.String()
}

func (s *CheckFreeStorageValidRequest) SetDeviceName(v string) *CheckFreeStorageValidRequest {
	s.DeviceName = &v
	return s
}

func (s *CheckFreeStorageValidRequest) SetIotId(v string) *CheckFreeStorageValidRequest {
	s.IotId = &v
	return s
}

func (s *CheckFreeStorageValidRequest) SetProductKey(v string) *CheckFreeStorageValidRequest {
	s.ProductKey = &v
	return s
}

type CheckFreeStorageValidResponseBody struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckFreeStorageValidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckFreeStorageValidResponseBody) GoString() string {
	return s.String()
}

func (s *CheckFreeStorageValidResponseBody) SetCode(v string) *CheckFreeStorageValidResponseBody {
	s.Code = &v
	return s
}

func (s *CheckFreeStorageValidResponseBody) SetErrorMessage(v string) *CheckFreeStorageValidResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CheckFreeStorageValidResponseBody) SetRequestId(v string) *CheckFreeStorageValidResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckFreeStorageValidResponseBody) SetSuccess(v bool) *CheckFreeStorageValidResponseBody {
	s.Success = &v
	return s
}

type CheckFreeStorageValidResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckFreeStorageValidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckFreeStorageValidResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckFreeStorageValidResponse) GoString() string {
	return s.String()
}

func (s *CheckFreeStorageValidResponse) SetHeaders(v map[string]*string) *CheckFreeStorageValidResponse {
	s.Headers = v
	return s
}

func (s *CheckFreeStorageValidResponse) SetStatusCode(v int32) *CheckFreeStorageValidResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckFreeStorageValidResponse) SetBody(v *CheckFreeStorageValidResponseBody) *CheckFreeStorageValidResponse {
	s.Body = v
	return s
}

type ConsumeFreeStorageRequest struct {
	DeviceName          *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	EnableDefaultPlan   *bool   `json:"EnableDefaultPlan,omitempty" xml:"EnableDefaultPlan,omitempty"`
	EventRecordDuration *int32  `json:"EventRecordDuration,omitempty" xml:"EventRecordDuration,omitempty"`
	EventRecordProlong  *bool   `json:"EventRecordProlong,omitempty" xml:"EventRecordProlong,omitempty"`
	ImmediateUse        *bool   `json:"ImmediateUse,omitempty" xml:"ImmediateUse,omitempty"`
	IotId               *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PreRecordDuration   *int32  `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	ProductKey          *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Quota               *int32  `json:"Quota,omitempty" xml:"Quota,omitempty"`
}

func (s ConsumeFreeStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s ConsumeFreeStorageRequest) GoString() string {
	return s.String()
}

func (s *ConsumeFreeStorageRequest) SetDeviceName(v string) *ConsumeFreeStorageRequest {
	s.DeviceName = &v
	return s
}

func (s *ConsumeFreeStorageRequest) SetEnableDefaultPlan(v bool) *ConsumeFreeStorageRequest {
	s.EnableDefaultPlan = &v
	return s
}

func (s *ConsumeFreeStorageRequest) SetEventRecordDuration(v int32) *ConsumeFreeStorageRequest {
	s.EventRecordDuration = &v
	return s
}

func (s *ConsumeFreeStorageRequest) SetEventRecordProlong(v bool) *ConsumeFreeStorageRequest {
	s.EventRecordProlong = &v
	return s
}

func (s *ConsumeFreeStorageRequest) SetImmediateUse(v bool) *ConsumeFreeStorageRequest {
	s.ImmediateUse = &v
	return s
}

func (s *ConsumeFreeStorageRequest) SetIotId(v string) *ConsumeFreeStorageRequest {
	s.IotId = &v
	return s
}

func (s *ConsumeFreeStorageRequest) SetPreRecordDuration(v int32) *ConsumeFreeStorageRequest {
	s.PreRecordDuration = &v
	return s
}

func (s *ConsumeFreeStorageRequest) SetProductKey(v string) *ConsumeFreeStorageRequest {
	s.ProductKey = &v
	return s
}

func (s *ConsumeFreeStorageRequest) SetQuota(v int32) *ConsumeFreeStorageRequest {
	s.Quota = &v
	return s
}

type ConsumeFreeStorageResponseBody struct {
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *ConsumeFreeStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConsumeFreeStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConsumeFreeStorageResponseBody) GoString() string {
	return s.String()
}

func (s *ConsumeFreeStorageResponseBody) SetCode(v string) *ConsumeFreeStorageResponseBody {
	s.Code = &v
	return s
}

func (s *ConsumeFreeStorageResponseBody) SetData(v *ConsumeFreeStorageResponseBodyData) *ConsumeFreeStorageResponseBody {
	s.Data = v
	return s
}

func (s *ConsumeFreeStorageResponseBody) SetErrorMessage(v string) *ConsumeFreeStorageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ConsumeFreeStorageResponseBody) SetRequestId(v string) *ConsumeFreeStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConsumeFreeStorageResponseBody) SetSuccess(v bool) *ConsumeFreeStorageResponseBody {
	s.Success = &v
	return s
}

type ConsumeFreeStorageResponseBodyData struct {
	Consumed     *int32  `json:"Consumed,omitempty" xml:"Consumed,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC   *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	Expired      *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Lifecycle    *int32  `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	Months       *int32  `json:"Months,omitempty" xml:"Months,omitempty"`
	RemainQuota  *int32  `json:"RemainQuota,omitempty" xml:"RemainQuota,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ConsumeFreeStorageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ConsumeFreeStorageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConsumeFreeStorageResponseBodyData) SetConsumed(v int32) *ConsumeFreeStorageResponseBodyData {
	s.Consumed = &v
	return s
}

func (s *ConsumeFreeStorageResponseBodyData) SetEndTime(v string) *ConsumeFreeStorageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ConsumeFreeStorageResponseBodyData) SetEndTimeUTC(v string) *ConsumeFreeStorageResponseBodyData {
	s.EndTimeUTC = &v
	return s
}

func (s *ConsumeFreeStorageResponseBodyData) SetExpired(v int32) *ConsumeFreeStorageResponseBodyData {
	s.Expired = &v
	return s
}

func (s *ConsumeFreeStorageResponseBodyData) SetLifecycle(v int32) *ConsumeFreeStorageResponseBodyData {
	s.Lifecycle = &v
	return s
}

func (s *ConsumeFreeStorageResponseBodyData) SetMonths(v int32) *ConsumeFreeStorageResponseBodyData {
	s.Months = &v
	return s
}

func (s *ConsumeFreeStorageResponseBodyData) SetRemainQuota(v int32) *ConsumeFreeStorageResponseBodyData {
	s.RemainQuota = &v
	return s
}

func (s *ConsumeFreeStorageResponseBodyData) SetStartTime(v string) *ConsumeFreeStorageResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ConsumeFreeStorageResponseBodyData) SetStartTimeUTC(v string) *ConsumeFreeStorageResponseBodyData {
	s.StartTimeUTC = &v
	return s
}

func (s *ConsumeFreeStorageResponseBodyData) SetType(v int32) *ConsumeFreeStorageResponseBodyData {
	s.Type = &v
	return s
}

type ConsumeFreeStorageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConsumeFreeStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConsumeFreeStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s ConsumeFreeStorageResponse) GoString() string {
	return s.String()
}

func (s *ConsumeFreeStorageResponse) SetHeaders(v map[string]*string) *ConsumeFreeStorageResponse {
	s.Headers = v
	return s
}

func (s *ConsumeFreeStorageResponse) SetStatusCode(v int32) *ConsumeFreeStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsumeFreeStorageResponse) SetBody(v *ConsumeFreeStorageResponseBody) *ConsumeFreeStorageResponse {
	s.Body = v
	return s
}

type CreateAndPayStorageOrderRequest struct {
	CommodityCode         *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Copies                *int32  `json:"Copies,omitempty" xml:"Copies,omitempty"`
	DeviceName            *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceNoOwner         *bool   `json:"DeviceNoOwner,omitempty" xml:"DeviceNoOwner,omitempty"`
	EnableDefaultPlan     *bool   `json:"EnableDefaultPlan,omitempty" xml:"EnableDefaultPlan,omitempty"`
	EventRecordDuration   *int32  `json:"EventRecordDuration,omitempty" xml:"EventRecordDuration,omitempty"`
	EventRecordProlong    *bool   `json:"EventRecordProlong,omitempty" xml:"EventRecordProlong,omitempty"`
	ImmediateUse          *bool   `json:"ImmediateUse,omitempty" xml:"ImmediateUse,omitempty"`
	IotId                 *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	MaxRecordFileDuration *int32  `json:"MaxRecordFileDuration,omitempty" xml:"MaxRecordFileDuration,omitempty"`
	PreRecordDuration     *int32  `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	ProductKey            *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	Specification         *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	UserId                *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName              *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateAndPayStorageOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPayStorageOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateAndPayStorageOrderRequest) SetCommodityCode(v string) *CreateAndPayStorageOrderRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetCopies(v int32) *CreateAndPayStorageOrderRequest {
	s.Copies = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetDeviceName(v string) *CreateAndPayStorageOrderRequest {
	s.DeviceName = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetDeviceNoOwner(v bool) *CreateAndPayStorageOrderRequest {
	s.DeviceNoOwner = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetEnableDefaultPlan(v bool) *CreateAndPayStorageOrderRequest {
	s.EnableDefaultPlan = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetEventRecordDuration(v int32) *CreateAndPayStorageOrderRequest {
	s.EventRecordDuration = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetEventRecordProlong(v bool) *CreateAndPayStorageOrderRequest {
	s.EventRecordProlong = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetImmediateUse(v bool) *CreateAndPayStorageOrderRequest {
	s.ImmediateUse = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetIotId(v string) *CreateAndPayStorageOrderRequest {
	s.IotId = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetMaxRecordFileDuration(v int32) *CreateAndPayStorageOrderRequest {
	s.MaxRecordFileDuration = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetPreRecordDuration(v int32) *CreateAndPayStorageOrderRequest {
	s.PreRecordDuration = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetProductKey(v string) *CreateAndPayStorageOrderRequest {
	s.ProductKey = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetSpecification(v string) *CreateAndPayStorageOrderRequest {
	s.Specification = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetUserId(v string) *CreateAndPayStorageOrderRequest {
	s.UserId = &v
	return s
}

func (s *CreateAndPayStorageOrderRequest) SetUserName(v string) *CreateAndPayStorageOrderRequest {
	s.UserName = &v
	return s
}

type CreateAndPayStorageOrderResponseBody struct {
	Code         *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *CreateAndPayStorageOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAndPayStorageOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPayStorageOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndPayStorageOrderResponseBody) SetCode(v string) *CreateAndPayStorageOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBody) SetData(v *CreateAndPayStorageOrderResponseBodyData) *CreateAndPayStorageOrderResponseBody {
	s.Data = v
	return s
}

func (s *CreateAndPayStorageOrderResponseBody) SetErrorMessage(v string) *CreateAndPayStorageOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBody) SetRequestId(v string) *CreateAndPayStorageOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBody) SetSuccess(v bool) *CreateAndPayStorageOrderResponseBody {
	s.Success = &v
	return s
}

type CreateAndPayStorageOrderResponseBodyData struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Copies        *int32  `json:"Copies,omitempty" xml:"Copies,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC    *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	IdentityId    *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderType     *int32  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OutOrderNo    *string `json:"OutOrderNo,omitempty" xml:"OutOrderNo,omitempty"`
	PaymentStatus *int32  `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	PreConsume    *int32  `json:"PreConsume,omitempty" xml:"PreConsume,omitempty"`
	Price         *string `json:"Price,omitempty" xml:"Price,omitempty"`
	RecordType    *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC  *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateAndPayStorageOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPayStorageOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetCommodityCode(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetCopies(v int32) *CreateAndPayStorageOrderResponseBodyData {
	s.Copies = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetEndTime(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetEndTimeUTC(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.EndTimeUTC = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetIdentityId(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.IdentityId = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetIotId(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.IotId = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetOrderId(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetOrderType(v int32) *CreateAndPayStorageOrderResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetOutOrderNo(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.OutOrderNo = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetPaymentStatus(v int32) *CreateAndPayStorageOrderResponseBodyData {
	s.PaymentStatus = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetPreConsume(v int32) *CreateAndPayStorageOrderResponseBodyData {
	s.PreConsume = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetPrice(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.Price = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetRecordType(v int32) *CreateAndPayStorageOrderResponseBodyData {
	s.RecordType = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetSpecification(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.Specification = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetStartTime(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetStartTimeUTC(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.StartTimeUTC = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetStatus(v int32) *CreateAndPayStorageOrderResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetUserId(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.UserId = &v
	return s
}

func (s *CreateAndPayStorageOrderResponseBodyData) SetUserName(v string) *CreateAndPayStorageOrderResponseBodyData {
	s.UserName = &v
	return s
}

type CreateAndPayStorageOrderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAndPayStorageOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAndPayStorageOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPayStorageOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateAndPayStorageOrderResponse) SetHeaders(v map[string]*string) *CreateAndPayStorageOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateAndPayStorageOrderResponse) SetStatusCode(v int32) *CreateAndPayStorageOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndPayStorageOrderResponse) SetBody(v *CreateAndPayStorageOrderResponseBody) *CreateAndPayStorageOrderResponse {
	s.Body = v
	return s
}

type EnableFreeStorageRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s EnableFreeStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableFreeStorageRequest) GoString() string {
	return s.String()
}

func (s *EnableFreeStorageRequest) SetDeviceName(v string) *EnableFreeStorageRequest {
	s.DeviceName = &v
	return s
}

func (s *EnableFreeStorageRequest) SetIotId(v string) *EnableFreeStorageRequest {
	s.IotId = &v
	return s
}

func (s *EnableFreeStorageRequest) SetProductKey(v string) *EnableFreeStorageRequest {
	s.ProductKey = &v
	return s
}

type EnableFreeStorageResponseBody struct {
	Code         *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *EnableFreeStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableFreeStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableFreeStorageResponseBody) GoString() string {
	return s.String()
}

func (s *EnableFreeStorageResponseBody) SetCode(v string) *EnableFreeStorageResponseBody {
	s.Code = &v
	return s
}

func (s *EnableFreeStorageResponseBody) SetData(v *EnableFreeStorageResponseBodyData) *EnableFreeStorageResponseBody {
	s.Data = v
	return s
}

func (s *EnableFreeStorageResponseBody) SetErrorMessage(v string) *EnableFreeStorageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *EnableFreeStorageResponseBody) SetRequestId(v string) *EnableFreeStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableFreeStorageResponseBody) SetSuccess(v bool) *EnableFreeStorageResponseBody {
	s.Success = &v
	return s
}

type EnableFreeStorageResponseBodyData struct {
	Consumed     *int32  `json:"Consumed,omitempty" xml:"Consumed,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC   *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	Expired      *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Lifecycle    *int32  `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	Months       *int32  `json:"Months,omitempty" xml:"Months,omitempty"`
	RemainQuota  *int32  `json:"RemainQuota,omitempty" xml:"RemainQuota,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s EnableFreeStorageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnableFreeStorageResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnableFreeStorageResponseBodyData) SetConsumed(v int32) *EnableFreeStorageResponseBodyData {
	s.Consumed = &v
	return s
}

func (s *EnableFreeStorageResponseBodyData) SetEndTime(v string) *EnableFreeStorageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *EnableFreeStorageResponseBodyData) SetEndTimeUTC(v string) *EnableFreeStorageResponseBodyData {
	s.EndTimeUTC = &v
	return s
}

func (s *EnableFreeStorageResponseBodyData) SetExpired(v int32) *EnableFreeStorageResponseBodyData {
	s.Expired = &v
	return s
}

func (s *EnableFreeStorageResponseBodyData) SetLifecycle(v int32) *EnableFreeStorageResponseBodyData {
	s.Lifecycle = &v
	return s
}

func (s *EnableFreeStorageResponseBodyData) SetMonths(v int32) *EnableFreeStorageResponseBodyData {
	s.Months = &v
	return s
}

func (s *EnableFreeStorageResponseBodyData) SetRemainQuota(v int32) *EnableFreeStorageResponseBodyData {
	s.RemainQuota = &v
	return s
}

func (s *EnableFreeStorageResponseBodyData) SetStartTime(v string) *EnableFreeStorageResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *EnableFreeStorageResponseBodyData) SetStartTimeUTC(v string) *EnableFreeStorageResponseBodyData {
	s.StartTimeUTC = &v
	return s
}

func (s *EnableFreeStorageResponseBodyData) SetType(v int32) *EnableFreeStorageResponseBodyData {
	s.Type = &v
	return s
}

type EnableFreeStorageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableFreeStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableFreeStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableFreeStorageResponse) GoString() string {
	return s.String()
}

func (s *EnableFreeStorageResponse) SetHeaders(v map[string]*string) *EnableFreeStorageResponse {
	s.Headers = v
	return s
}

func (s *EnableFreeStorageResponse) SetStatusCode(v int32) *EnableFreeStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableFreeStorageResponse) SetBody(v *EnableFreeStorageResponseBody) *EnableFreeStorageResponse {
	s.Body = v
	return s
}

type EnableStorageOrderRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s EnableStorageOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableStorageOrderRequest) GoString() string {
	return s.String()
}

func (s *EnableStorageOrderRequest) SetDeviceName(v string) *EnableStorageOrderRequest {
	s.DeviceName = &v
	return s
}

func (s *EnableStorageOrderRequest) SetIotId(v string) *EnableStorageOrderRequest {
	s.IotId = &v
	return s
}

func (s *EnableStorageOrderRequest) SetOrderId(v string) *EnableStorageOrderRequest {
	s.OrderId = &v
	return s
}

func (s *EnableStorageOrderRequest) SetProductKey(v string) *EnableStorageOrderRequest {
	s.ProductKey = &v
	return s
}

type EnableStorageOrderResponseBody struct {
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *EnableStorageOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableStorageOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableStorageOrderResponseBody) GoString() string {
	return s.String()
}

func (s *EnableStorageOrderResponseBody) SetCode(v string) *EnableStorageOrderResponseBody {
	s.Code = &v
	return s
}

func (s *EnableStorageOrderResponseBody) SetData(v *EnableStorageOrderResponseBodyData) *EnableStorageOrderResponseBody {
	s.Data = v
	return s
}

func (s *EnableStorageOrderResponseBody) SetErrorMessage(v string) *EnableStorageOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *EnableStorageOrderResponseBody) SetRequestId(v string) *EnableStorageOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableStorageOrderResponseBody) SetSuccess(v bool) *EnableStorageOrderResponseBody {
	s.Success = &v
	return s
}

type EnableStorageOrderResponseBodyData struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Copies        *int32  `json:"Copies,omitempty" xml:"Copies,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC    *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	IdentityId    *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderType     *int32  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OutOrderNo    *string `json:"OutOrderNo,omitempty" xml:"OutOrderNo,omitempty"`
	PaymentStatus *int32  `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	PreConsume    *int32  `json:"PreConsume,omitempty" xml:"PreConsume,omitempty"`
	Price         *string `json:"Price,omitempty" xml:"Price,omitempty"`
	RecordType    *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC  *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s EnableStorageOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EnableStorageOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *EnableStorageOrderResponseBodyData) SetCommodityCode(v string) *EnableStorageOrderResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetCopies(v int32) *EnableStorageOrderResponseBodyData {
	s.Copies = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetEndTime(v string) *EnableStorageOrderResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetEndTimeUTC(v string) *EnableStorageOrderResponseBodyData {
	s.EndTimeUTC = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetIdentityId(v string) *EnableStorageOrderResponseBodyData {
	s.IdentityId = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetIotId(v string) *EnableStorageOrderResponseBodyData {
	s.IotId = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetOrderId(v string) *EnableStorageOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetOrderType(v int32) *EnableStorageOrderResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetOutOrderNo(v string) *EnableStorageOrderResponseBodyData {
	s.OutOrderNo = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetPaymentStatus(v int32) *EnableStorageOrderResponseBodyData {
	s.PaymentStatus = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetPreConsume(v int32) *EnableStorageOrderResponseBodyData {
	s.PreConsume = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetPrice(v string) *EnableStorageOrderResponseBodyData {
	s.Price = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetRecordType(v int32) *EnableStorageOrderResponseBodyData {
	s.RecordType = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetSpecification(v string) *EnableStorageOrderResponseBodyData {
	s.Specification = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetStartTime(v string) *EnableStorageOrderResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetStartTimeUTC(v string) *EnableStorageOrderResponseBodyData {
	s.StartTimeUTC = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetStatus(v int32) *EnableStorageOrderResponseBodyData {
	s.Status = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetUserId(v string) *EnableStorageOrderResponseBodyData {
	s.UserId = &v
	return s
}

func (s *EnableStorageOrderResponseBodyData) SetUserName(v string) *EnableStorageOrderResponseBodyData {
	s.UserName = &v
	return s
}

type EnableStorageOrderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableStorageOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableStorageOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableStorageOrderResponse) GoString() string {
	return s.String()
}

func (s *EnableStorageOrderResponse) SetHeaders(v map[string]*string) *EnableStorageOrderResponse {
	s.Headers = v
	return s
}

func (s *EnableStorageOrderResponse) SetStatusCode(v int32) *EnableStorageOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableStorageOrderResponse) SetBody(v *EnableStorageOrderResponseBody) *EnableStorageOrderResponse {
	s.Body = v
	return s
}

type FreezeFreeStorageRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s FreezeFreeStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s FreezeFreeStorageRequest) GoString() string {
	return s.String()
}

func (s *FreezeFreeStorageRequest) SetDeviceName(v string) *FreezeFreeStorageRequest {
	s.DeviceName = &v
	return s
}

func (s *FreezeFreeStorageRequest) SetIotId(v string) *FreezeFreeStorageRequest {
	s.IotId = &v
	return s
}

func (s *FreezeFreeStorageRequest) SetProductKey(v string) *FreezeFreeStorageRequest {
	s.ProductKey = &v
	return s
}

type FreezeFreeStorageResponseBody struct {
	Code         *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *FreezeFreeStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FreezeFreeStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FreezeFreeStorageResponseBody) GoString() string {
	return s.String()
}

func (s *FreezeFreeStorageResponseBody) SetCode(v string) *FreezeFreeStorageResponseBody {
	s.Code = &v
	return s
}

func (s *FreezeFreeStorageResponseBody) SetData(v *FreezeFreeStorageResponseBodyData) *FreezeFreeStorageResponseBody {
	s.Data = v
	return s
}

func (s *FreezeFreeStorageResponseBody) SetErrorMessage(v string) *FreezeFreeStorageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *FreezeFreeStorageResponseBody) SetRequestId(v string) *FreezeFreeStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *FreezeFreeStorageResponseBody) SetSuccess(v bool) *FreezeFreeStorageResponseBody {
	s.Success = &v
	return s
}

type FreezeFreeStorageResponseBodyData struct {
	Consumed     *int32  `json:"Consumed,omitempty" xml:"Consumed,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC   *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	Expired      *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Lifecycle    *int32  `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	Months       *int32  `json:"Months,omitempty" xml:"Months,omitempty"`
	RemainQuota  *int32  `json:"RemainQuota,omitempty" xml:"RemainQuota,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s FreezeFreeStorageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FreezeFreeStorageResponseBodyData) GoString() string {
	return s.String()
}

func (s *FreezeFreeStorageResponseBodyData) SetConsumed(v int32) *FreezeFreeStorageResponseBodyData {
	s.Consumed = &v
	return s
}

func (s *FreezeFreeStorageResponseBodyData) SetEndTime(v string) *FreezeFreeStorageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *FreezeFreeStorageResponseBodyData) SetEndTimeUTC(v string) *FreezeFreeStorageResponseBodyData {
	s.EndTimeUTC = &v
	return s
}

func (s *FreezeFreeStorageResponseBodyData) SetExpired(v int32) *FreezeFreeStorageResponseBodyData {
	s.Expired = &v
	return s
}

func (s *FreezeFreeStorageResponseBodyData) SetLifecycle(v int32) *FreezeFreeStorageResponseBodyData {
	s.Lifecycle = &v
	return s
}

func (s *FreezeFreeStorageResponseBodyData) SetMonths(v int32) *FreezeFreeStorageResponseBodyData {
	s.Months = &v
	return s
}

func (s *FreezeFreeStorageResponseBodyData) SetRemainQuota(v int32) *FreezeFreeStorageResponseBodyData {
	s.RemainQuota = &v
	return s
}

func (s *FreezeFreeStorageResponseBodyData) SetStartTime(v string) *FreezeFreeStorageResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *FreezeFreeStorageResponseBodyData) SetStartTimeUTC(v string) *FreezeFreeStorageResponseBodyData {
	s.StartTimeUTC = &v
	return s
}

func (s *FreezeFreeStorageResponseBodyData) SetType(v int32) *FreezeFreeStorageResponseBodyData {
	s.Type = &v
	return s
}

type FreezeFreeStorageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FreezeFreeStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FreezeFreeStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s FreezeFreeStorageResponse) GoString() string {
	return s.String()
}

func (s *FreezeFreeStorageResponse) SetHeaders(v map[string]*string) *FreezeFreeStorageResponse {
	s.Headers = v
	return s
}

func (s *FreezeFreeStorageResponse) SetStatusCode(v int32) *FreezeFreeStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *FreezeFreeStorageResponse) SetBody(v *FreezeFreeStorageResponseBody) *FreezeFreeStorageResponse {
	s.Body = v
	return s
}

type FreezeStorageOrderRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceNoOwner *bool   `json:"DeviceNoOwner,omitempty" xml:"DeviceNoOwner,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s FreezeStorageOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s FreezeStorageOrderRequest) GoString() string {
	return s.String()
}

func (s *FreezeStorageOrderRequest) SetDeviceName(v string) *FreezeStorageOrderRequest {
	s.DeviceName = &v
	return s
}

func (s *FreezeStorageOrderRequest) SetDeviceNoOwner(v bool) *FreezeStorageOrderRequest {
	s.DeviceNoOwner = &v
	return s
}

func (s *FreezeStorageOrderRequest) SetIotId(v string) *FreezeStorageOrderRequest {
	s.IotId = &v
	return s
}

func (s *FreezeStorageOrderRequest) SetOrderId(v string) *FreezeStorageOrderRequest {
	s.OrderId = &v
	return s
}

func (s *FreezeStorageOrderRequest) SetProductKey(v string) *FreezeStorageOrderRequest {
	s.ProductKey = &v
	return s
}

type FreezeStorageOrderResponseBody struct {
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *FreezeStorageOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FreezeStorageOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FreezeStorageOrderResponseBody) GoString() string {
	return s.String()
}

func (s *FreezeStorageOrderResponseBody) SetCode(v string) *FreezeStorageOrderResponseBody {
	s.Code = &v
	return s
}

func (s *FreezeStorageOrderResponseBody) SetData(v *FreezeStorageOrderResponseBodyData) *FreezeStorageOrderResponseBody {
	s.Data = v
	return s
}

func (s *FreezeStorageOrderResponseBody) SetErrorMessage(v string) *FreezeStorageOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *FreezeStorageOrderResponseBody) SetRequestId(v string) *FreezeStorageOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *FreezeStorageOrderResponseBody) SetSuccess(v bool) *FreezeStorageOrderResponseBody {
	s.Success = &v
	return s
}

type FreezeStorageOrderResponseBodyData struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Copies        *int32  `json:"Copies,omitempty" xml:"Copies,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC    *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	IdentityId    *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderType     *int32  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OutOrderNo    *string `json:"OutOrderNo,omitempty" xml:"OutOrderNo,omitempty"`
	PaymentStatus *int32  `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	PreConsume    *int32  `json:"PreConsume,omitempty" xml:"PreConsume,omitempty"`
	Price         *string `json:"Price,omitempty" xml:"Price,omitempty"`
	RecordType    *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC  *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s FreezeStorageOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FreezeStorageOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *FreezeStorageOrderResponseBodyData) SetCommodityCode(v string) *FreezeStorageOrderResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetCopies(v int32) *FreezeStorageOrderResponseBodyData {
	s.Copies = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetEndTime(v string) *FreezeStorageOrderResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetEndTimeUTC(v string) *FreezeStorageOrderResponseBodyData {
	s.EndTimeUTC = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetIdentityId(v string) *FreezeStorageOrderResponseBodyData {
	s.IdentityId = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetIotId(v string) *FreezeStorageOrderResponseBodyData {
	s.IotId = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetOrderId(v string) *FreezeStorageOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetOrderType(v int32) *FreezeStorageOrderResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetOutOrderNo(v string) *FreezeStorageOrderResponseBodyData {
	s.OutOrderNo = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetPaymentStatus(v int32) *FreezeStorageOrderResponseBodyData {
	s.PaymentStatus = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetPreConsume(v int32) *FreezeStorageOrderResponseBodyData {
	s.PreConsume = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetPrice(v string) *FreezeStorageOrderResponseBodyData {
	s.Price = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetRecordType(v int32) *FreezeStorageOrderResponseBodyData {
	s.RecordType = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetSpecification(v string) *FreezeStorageOrderResponseBodyData {
	s.Specification = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetStartTime(v string) *FreezeStorageOrderResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetStartTimeUTC(v string) *FreezeStorageOrderResponseBodyData {
	s.StartTimeUTC = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetStatus(v int32) *FreezeStorageOrderResponseBodyData {
	s.Status = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetUserId(v string) *FreezeStorageOrderResponseBodyData {
	s.UserId = &v
	return s
}

func (s *FreezeStorageOrderResponseBodyData) SetUserName(v string) *FreezeStorageOrderResponseBodyData {
	s.UserName = &v
	return s
}

type FreezeStorageOrderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FreezeStorageOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FreezeStorageOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s FreezeStorageOrderResponse) GoString() string {
	return s.String()
}

func (s *FreezeStorageOrderResponse) SetHeaders(v map[string]*string) *FreezeStorageOrderResponse {
	s.Headers = v
	return s
}

func (s *FreezeStorageOrderResponse) SetStatusCode(v int32) *FreezeStorageOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *FreezeStorageOrderResponse) SetBody(v *FreezeStorageOrderResponseBody) *FreezeStorageOrderResponse {
	s.Body = v
	return s
}

type GenerateDeviceRequest struct {
	Amount     *int64  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GenerateDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceRequest) GoString() string {
	return s.String()
}

func (s *GenerateDeviceRequest) SetAmount(v int64) *GenerateDeviceRequest {
	s.Amount = &v
	return s
}

func (s *GenerateDeviceRequest) SetProductKey(v string) *GenerateDeviceRequest {
	s.ProductKey = &v
	return s
}

func (s *GenerateDeviceRequest) SetProjectId(v string) *GenerateDeviceRequest {
	s.ProjectId = &v
	return s
}

type GenerateDeviceResponseBody struct {
	Code         *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GenerateDeviceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDeviceResponseBody) SetCode(v string) *GenerateDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateDeviceResponseBody) SetData(v *GenerateDeviceResponseBodyData) *GenerateDeviceResponseBody {
	s.Data = v
	return s
}

func (s *GenerateDeviceResponseBody) SetErrorMessage(v string) *GenerateDeviceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GenerateDeviceResponseBody) SetRequestId(v string) *GenerateDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDeviceResponseBody) SetSuccess(v bool) *GenerateDeviceResponseBody {
	s.Success = &v
	return s
}

type GenerateDeviceResponseBodyData struct {
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
}

func (s GenerateDeviceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateDeviceResponseBodyData) SetBatchId(v string) *GenerateDeviceResponseBodyData {
	s.BatchId = &v
	return s
}

type GenerateDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceResponse) GoString() string {
	return s.String()
}

func (s *GenerateDeviceResponse) SetHeaders(v map[string]*string) *GenerateDeviceResponse {
	s.Headers = v
	return s
}

func (s *GenerateDeviceResponse) SetStatusCode(v int32) *GenerateDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDeviceResponse) SetBody(v *GenerateDeviceResponseBody) *GenerateDeviceResponse {
	s.Body = v
	return s
}

type GenerateDeviceByBatchIdRequest struct {
	BatchId    *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GenerateDeviceByBatchIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceByBatchIdRequest) GoString() string {
	return s.String()
}

func (s *GenerateDeviceByBatchIdRequest) SetBatchId(v string) *GenerateDeviceByBatchIdRequest {
	s.BatchId = &v
	return s
}

func (s *GenerateDeviceByBatchIdRequest) SetProductKey(v string) *GenerateDeviceByBatchIdRequest {
	s.ProductKey = &v
	return s
}

func (s *GenerateDeviceByBatchIdRequest) SetProjectId(v string) *GenerateDeviceByBatchIdRequest {
	s.ProjectId = &v
	return s
}

type GenerateDeviceByBatchIdResponseBody struct {
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *GenerateDeviceByBatchIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateDeviceByBatchIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceByBatchIdResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDeviceByBatchIdResponseBody) SetCode(v string) *GenerateDeviceByBatchIdResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateDeviceByBatchIdResponseBody) SetData(v *GenerateDeviceByBatchIdResponseBodyData) *GenerateDeviceByBatchIdResponseBody {
	s.Data = v
	return s
}

func (s *GenerateDeviceByBatchIdResponseBody) SetErrorMessage(v string) *GenerateDeviceByBatchIdResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GenerateDeviceByBatchIdResponseBody) SetRequestId(v string) *GenerateDeviceByBatchIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDeviceByBatchIdResponseBody) SetSuccess(v bool) *GenerateDeviceByBatchIdResponseBody {
	s.Success = &v
	return s
}

type GenerateDeviceByBatchIdResponseBodyData struct {
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
}

func (s GenerateDeviceByBatchIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceByBatchIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateDeviceByBatchIdResponseBodyData) SetBatchId(v string) *GenerateDeviceByBatchIdResponseBodyData {
	s.BatchId = &v
	return s
}

type GenerateDeviceByBatchIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateDeviceByBatchIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateDeviceByBatchIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDeviceByBatchIdResponse) GoString() string {
	return s.String()
}

func (s *GenerateDeviceByBatchIdResponse) SetHeaders(v map[string]*string) *GenerateDeviceByBatchIdResponse {
	s.Headers = v
	return s
}

func (s *GenerateDeviceByBatchIdResponse) SetStatusCode(v int32) *GenerateDeviceByBatchIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDeviceByBatchIdResponse) SetBody(v *GenerateDeviceByBatchIdResponseBody) *GenerateDeviceByBatchIdResponse {
	s.Body = v
	return s
}

type QueryBatchStatusRequest struct {
	BatchId    *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s QueryBatchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryBatchStatusRequest) SetBatchId(v string) *QueryBatchStatusRequest {
	s.BatchId = &v
	return s
}

func (s *QueryBatchStatusRequest) SetProductKey(v string) *QueryBatchStatusRequest {
	s.ProductKey = &v
	return s
}

func (s *QueryBatchStatusRequest) SetProjectId(v string) *QueryBatchStatusRequest {
	s.ProjectId = &v
	return s
}

type QueryBatchStatusResponseBody struct {
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryBatchStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryBatchStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBatchStatusResponseBody) SetCode(v string) *QueryBatchStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryBatchStatusResponseBody) SetData(v *QueryBatchStatusResponseBodyData) *QueryBatchStatusResponseBody {
	s.Data = v
	return s
}

func (s *QueryBatchStatusResponseBody) SetErrorMessage(v string) *QueryBatchStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryBatchStatusResponseBody) SetRequestId(v string) *QueryBatchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBatchStatusResponseBody) SetSuccess(v bool) *QueryBatchStatusResponseBody {
	s.Success = &v
	return s
}

type QueryBatchStatusResponseBodyData struct {
	InvalidDetailList []*QueryBatchStatusResponseBodyDataInvalidDetailList `json:"InvalidDetailList,omitempty" xml:"InvalidDetailList,omitempty" type:"Repeated"`
	InvalidList       []*string                                            `json:"InvalidList,omitempty" xml:"InvalidList,omitempty" type:"Repeated"`
	Status            *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	ValidList         []*string                                            `json:"ValidList,omitempty" xml:"ValidList,omitempty" type:"Repeated"`
}

func (s QueryBatchStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBatchStatusResponseBodyData) SetInvalidDetailList(v []*QueryBatchStatusResponseBodyDataInvalidDetailList) *QueryBatchStatusResponseBodyData {
	s.InvalidDetailList = v
	return s
}

func (s *QueryBatchStatusResponseBodyData) SetInvalidList(v []*string) *QueryBatchStatusResponseBodyData {
	s.InvalidList = v
	return s
}

func (s *QueryBatchStatusResponseBodyData) SetStatus(v string) *QueryBatchStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryBatchStatusResponseBodyData) SetValidList(v []*string) *QueryBatchStatusResponseBodyData {
	s.ValidList = v
	return s
}

type QueryBatchStatusResponseBodyDataInvalidDetailList struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
}

func (s QueryBatchStatusResponseBodyDataInvalidDetailList) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchStatusResponseBodyDataInvalidDetailList) GoString() string {
	return s.String()
}

func (s *QueryBatchStatusResponseBodyDataInvalidDetailList) SetDeviceName(v string) *QueryBatchStatusResponseBodyDataInvalidDetailList {
	s.DeviceName = &v
	return s
}

func (s *QueryBatchStatusResponseBodyDataInvalidDetailList) SetErrorMsg(v string) *QueryBatchStatusResponseBodyDataInvalidDetailList {
	s.ErrorMsg = &v
	return s
}

type QueryBatchStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryBatchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryBatchStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBatchStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryBatchStatusResponse) SetHeaders(v map[string]*string) *QueryBatchStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryBatchStatusResponse) SetStatusCode(v int32) *QueryBatchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBatchStatusResponse) SetBody(v *QueryBatchStatusResponseBody) *QueryBatchStatusResponse {
	s.Body = v
	return s
}

type QueryDevicesDownloadUrlRequest struct {
	BatchId    *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryDevicesDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicesDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryDevicesDownloadUrlRequest) SetBatchId(v string) *QueryDevicesDownloadUrlRequest {
	s.BatchId = &v
	return s
}

func (s *QueryDevicesDownloadUrlRequest) SetProductKey(v string) *QueryDevicesDownloadUrlRequest {
	s.ProductKey = &v
	return s
}

type QueryDevicesDownloadUrlResponseBody struct {
	Code         *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryDevicesDownloadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDevicesDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicesDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDevicesDownloadUrlResponseBody) SetCode(v string) *QueryDevicesDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDevicesDownloadUrlResponseBody) SetData(v *QueryDevicesDownloadUrlResponseBodyData) *QueryDevicesDownloadUrlResponseBody {
	s.Data = v
	return s
}

func (s *QueryDevicesDownloadUrlResponseBody) SetErrorMessage(v string) *QueryDevicesDownloadUrlResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryDevicesDownloadUrlResponseBody) SetRequestId(v string) *QueryDevicesDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDevicesDownloadUrlResponseBody) SetSuccess(v bool) *QueryDevicesDownloadUrlResponseBody {
	s.Success = &v
	return s
}

type QueryDevicesDownloadUrlResponseBodyData struct {
	OssDownloadUrl *string `json:"OssDownloadUrl,omitempty" xml:"OssDownloadUrl,omitempty"`
}

func (s QueryDevicesDownloadUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicesDownloadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDevicesDownloadUrlResponseBodyData) SetOssDownloadUrl(v string) *QueryDevicesDownloadUrlResponseBodyData {
	s.OssDownloadUrl = &v
	return s
}

type QueryDevicesDownloadUrlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryDevicesDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryDevicesDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDevicesDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryDevicesDownloadUrlResponse) SetHeaders(v map[string]*string) *QueryDevicesDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryDevicesDownloadUrlResponse) SetStatusCode(v int32) *QueryDevicesDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDevicesDownloadUrlResponse) SetBody(v *QueryDevicesDownloadUrlResponseBody) *QueryDevicesDownloadUrlResponse {
	s.Body = v
	return s
}

type QueryFreeStorageRequest struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	ProductKey *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryFreeStorageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryFreeStorageRequest) GoString() string {
	return s.String()
}

func (s *QueryFreeStorageRequest) SetDeviceName(v string) *QueryFreeStorageRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryFreeStorageRequest) SetIotId(v string) *QueryFreeStorageRequest {
	s.IotId = &v
	return s
}

func (s *QueryFreeStorageRequest) SetProductKey(v string) *QueryFreeStorageRequest {
	s.ProductKey = &v
	return s
}

type QueryFreeStorageResponseBody struct {
	Code         *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryFreeStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFreeStorageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryFreeStorageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFreeStorageResponseBody) SetCode(v string) *QueryFreeStorageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFreeStorageResponseBody) SetData(v *QueryFreeStorageResponseBodyData) *QueryFreeStorageResponseBody {
	s.Data = v
	return s
}

func (s *QueryFreeStorageResponseBody) SetErrorMessage(v string) *QueryFreeStorageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryFreeStorageResponseBody) SetRequestId(v string) *QueryFreeStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFreeStorageResponseBody) SetSuccess(v bool) *QueryFreeStorageResponseBody {
	s.Success = &v
	return s
}

type QueryFreeStorageResponseBodyData struct {
	Consumed     *int32  `json:"Consumed,omitempty" xml:"Consumed,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC   *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	Expired      *int32  `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Lifecycle    *int32  `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	Months       *int32  `json:"Months,omitempty" xml:"Months,omitempty"`
	RemainQuota  *int32  `json:"RemainQuota,omitempty" xml:"RemainQuota,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryFreeStorageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryFreeStorageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryFreeStorageResponseBodyData) SetConsumed(v int32) *QueryFreeStorageResponseBodyData {
	s.Consumed = &v
	return s
}

func (s *QueryFreeStorageResponseBodyData) SetEndTime(v string) *QueryFreeStorageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryFreeStorageResponseBodyData) SetEndTimeUTC(v string) *QueryFreeStorageResponseBodyData {
	s.EndTimeUTC = &v
	return s
}

func (s *QueryFreeStorageResponseBodyData) SetExpired(v int32) *QueryFreeStorageResponseBodyData {
	s.Expired = &v
	return s
}

func (s *QueryFreeStorageResponseBodyData) SetLifecycle(v int32) *QueryFreeStorageResponseBodyData {
	s.Lifecycle = &v
	return s
}

func (s *QueryFreeStorageResponseBodyData) SetMonths(v int32) *QueryFreeStorageResponseBodyData {
	s.Months = &v
	return s
}

func (s *QueryFreeStorageResponseBodyData) SetRemainQuota(v int32) *QueryFreeStorageResponseBodyData {
	s.RemainQuota = &v
	return s
}

func (s *QueryFreeStorageResponseBodyData) SetStartTime(v string) *QueryFreeStorageResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *QueryFreeStorageResponseBodyData) SetStartTimeUTC(v string) *QueryFreeStorageResponseBodyData {
	s.StartTimeUTC = &v
	return s
}

func (s *QueryFreeStorageResponseBodyData) SetType(v int32) *QueryFreeStorageResponseBodyData {
	s.Type = &v
	return s
}

type QueryFreeStorageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryFreeStorageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryFreeStorageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryFreeStorageResponse) GoString() string {
	return s.String()
}

func (s *QueryFreeStorageResponse) SetHeaders(v map[string]*string) *QueryFreeStorageResponse {
	s.Headers = v
	return s
}

func (s *QueryFreeStorageResponse) SetStatusCode(v int32) *QueryFreeStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFreeStorageResponse) SetBody(v *QueryFreeStorageResponseBody) *QueryFreeStorageResponse {
	s.Body = v
	return s
}

type QueryGenerateDevicesInfoListRequest struct {
	BatchId   *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s QueryGenerateDevicesInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGenerateDevicesInfoListRequest) GoString() string {
	return s.String()
}

func (s *QueryGenerateDevicesInfoListRequest) SetBatchId(v string) *QueryGenerateDevicesInfoListRequest {
	s.BatchId = &v
	return s
}

func (s *QueryGenerateDevicesInfoListRequest) SetPageNo(v int32) *QueryGenerateDevicesInfoListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryGenerateDevicesInfoListRequest) SetPageSize(v int32) *QueryGenerateDevicesInfoListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryGenerateDevicesInfoListRequest) SetProjectId(v string) *QueryGenerateDevicesInfoListRequest {
	s.ProjectId = &v
	return s
}

type QueryGenerateDevicesInfoListResponseBody struct {
	Code         *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryGenerateDevicesInfoListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryGenerateDevicesInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGenerateDevicesInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGenerateDevicesInfoListResponseBody) SetCode(v string) *QueryGenerateDevicesInfoListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGenerateDevicesInfoListResponseBody) SetData(v *QueryGenerateDevicesInfoListResponseBodyData) *QueryGenerateDevicesInfoListResponseBody {
	s.Data = v
	return s
}

func (s *QueryGenerateDevicesInfoListResponseBody) SetErrorMessage(v string) *QueryGenerateDevicesInfoListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryGenerateDevicesInfoListResponseBody) SetRequestId(v string) *QueryGenerateDevicesInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGenerateDevicesInfoListResponseBody) SetSuccess(v bool) *QueryGenerateDevicesInfoListResponseBody {
	s.Success = &v
	return s
}

type QueryGenerateDevicesInfoListResponseBodyData struct {
	ListData []*QueryGenerateDevicesInfoListResponseBodyDataListData `json:"ListData,omitempty" xml:"ListData,omitempty" type:"Repeated"`
	PageNo   *int32                                                  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryGenerateDevicesInfoListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryGenerateDevicesInfoListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryGenerateDevicesInfoListResponseBodyData) SetListData(v []*QueryGenerateDevicesInfoListResponseBodyDataListData) *QueryGenerateDevicesInfoListResponseBodyData {
	s.ListData = v
	return s
}

func (s *QueryGenerateDevicesInfoListResponseBodyData) SetPageNo(v int32) *QueryGenerateDevicesInfoListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryGenerateDevicesInfoListResponseBodyData) SetPageSize(v int32) *QueryGenerateDevicesInfoListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryGenerateDevicesInfoListResponseBodyData) SetTotal(v int32) *QueryGenerateDevicesInfoListResponseBodyData {
	s.Total = &v
	return s
}

type QueryGenerateDevicesInfoListResponseBodyDataListData struct {
	DeviceName   *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceSecret *string `json:"DeviceSecret,omitempty" xml:"DeviceSecret,omitempty"`
	IotId        *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s QueryGenerateDevicesInfoListResponseBodyDataListData) String() string {
	return tea.Prettify(s)
}

func (s QueryGenerateDevicesInfoListResponseBodyDataListData) GoString() string {
	return s.String()
}

func (s *QueryGenerateDevicesInfoListResponseBodyDataListData) SetDeviceName(v string) *QueryGenerateDevicesInfoListResponseBodyDataListData {
	s.DeviceName = &v
	return s
}

func (s *QueryGenerateDevicesInfoListResponseBodyDataListData) SetDeviceSecret(v string) *QueryGenerateDevicesInfoListResponseBodyDataListData {
	s.DeviceSecret = &v
	return s
}

func (s *QueryGenerateDevicesInfoListResponseBodyDataListData) SetIotId(v string) *QueryGenerateDevicesInfoListResponseBodyDataListData {
	s.IotId = &v
	return s
}

type QueryGenerateDevicesInfoListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryGenerateDevicesInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGenerateDevicesInfoListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGenerateDevicesInfoListResponse) GoString() string {
	return s.String()
}

func (s *QueryGenerateDevicesInfoListResponse) SetHeaders(v map[string]*string) *QueryGenerateDevicesInfoListResponse {
	s.Headers = v
	return s
}

func (s *QueryGenerateDevicesInfoListResponse) SetStatusCode(v int32) *QueryGenerateDevicesInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGenerateDevicesInfoListResponse) SetBody(v *QueryGenerateDevicesInfoListResponseBody) *QueryGenerateDevicesInfoListResponse {
	s.Body = v
	return s
}

type QueryGenerateDevicesRecordRequest struct {
	EndTime   *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNo    *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryGenerateDevicesRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGenerateDevicesRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryGenerateDevicesRecordRequest) SetEndTime(v int64) *QueryGenerateDevicesRecordRequest {
	s.EndTime = &v
	return s
}

func (s *QueryGenerateDevicesRecordRequest) SetPageNo(v int32) *QueryGenerateDevicesRecordRequest {
	s.PageNo = &v
	return s
}

func (s *QueryGenerateDevicesRecordRequest) SetPageSize(v int32) *QueryGenerateDevicesRecordRequest {
	s.PageSize = &v
	return s
}

func (s *QueryGenerateDevicesRecordRequest) SetStartTime(v int64) *QueryGenerateDevicesRecordRequest {
	s.StartTime = &v
	return s
}

type QueryGenerateDevicesRecordResponseBody struct {
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryGenerateDevicesRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                     `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryGenerateDevicesRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGenerateDevicesRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGenerateDevicesRecordResponseBody) SetCode(v string) *QueryGenerateDevicesRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBody) SetData(v *QueryGenerateDevicesRecordResponseBodyData) *QueryGenerateDevicesRecordResponseBody {
	s.Data = v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBody) SetErrorMessage(v string) *QueryGenerateDevicesRecordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBody) SetRequestId(v string) *QueryGenerateDevicesRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBody) SetSuccess(v bool) *QueryGenerateDevicesRecordResponseBody {
	s.Success = &v
	return s
}

type QueryGenerateDevicesRecordResponseBodyData struct {
	ListData []*QueryGenerateDevicesRecordResponseBodyDataListData `json:"ListData,omitempty" xml:"ListData,omitempty" type:"Repeated"`
	PageNo   *int32                                                `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryGenerateDevicesRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryGenerateDevicesRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryGenerateDevicesRecordResponseBodyData) SetListData(v []*QueryGenerateDevicesRecordResponseBodyDataListData) *QueryGenerateDevicesRecordResponseBodyData {
	s.ListData = v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBodyData) SetPageNo(v int32) *QueryGenerateDevicesRecordResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBodyData) SetPageSize(v int32) *QueryGenerateDevicesRecordResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBodyData) SetTotal(v int32) *QueryGenerateDevicesRecordResponseBodyData {
	s.Total = &v
	return s
}

type QueryGenerateDevicesRecordResponseBodyDataListData struct {
	ApplyDeviceCount *int64  `json:"ApplyDeviceCount,omitempty" xml:"ApplyDeviceCount,omitempty"`
	BatchId          *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	BatchStatus      *string `json:"BatchStatus,omitempty" xml:"BatchStatus,omitempty"`
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	NetworkType      *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OperateUid       *int64  `json:"OperateUid,omitempty" xml:"OperateUid,omitempty"`
	ProductKey       *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	ProductName      *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	SpecCode         *string `json:"SpecCode,omitempty" xml:"SpecCode,omitempty"`
	SuccessCount     *int64  `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s QueryGenerateDevicesRecordResponseBodyDataListData) String() string {
	return tea.Prettify(s)
}

func (s QueryGenerateDevicesRecordResponseBodyDataListData) GoString() string {
	return s.String()
}

func (s *QueryGenerateDevicesRecordResponseBodyDataListData) SetApplyDeviceCount(v int64) *QueryGenerateDevicesRecordResponseBodyDataListData {
	s.ApplyDeviceCount = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBodyDataListData) SetBatchId(v string) *QueryGenerateDevicesRecordResponseBodyDataListData {
	s.BatchId = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBodyDataListData) SetBatchStatus(v string) *QueryGenerateDevicesRecordResponseBodyDataListData {
	s.BatchStatus = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBodyDataListData) SetCreateTime(v int64) *QueryGenerateDevicesRecordResponseBodyDataListData {
	s.CreateTime = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBodyDataListData) SetNetworkType(v string) *QueryGenerateDevicesRecordResponseBodyDataListData {
	s.NetworkType = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBodyDataListData) SetOperateUid(v int64) *QueryGenerateDevicesRecordResponseBodyDataListData {
	s.OperateUid = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBodyDataListData) SetProductKey(v string) *QueryGenerateDevicesRecordResponseBodyDataListData {
	s.ProductKey = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBodyDataListData) SetProductName(v string) *QueryGenerateDevicesRecordResponseBodyDataListData {
	s.ProductName = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBodyDataListData) SetSpecCode(v string) *QueryGenerateDevicesRecordResponseBodyDataListData {
	s.SpecCode = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponseBodyDataListData) SetSuccessCount(v int64) *QueryGenerateDevicesRecordResponseBodyDataListData {
	s.SuccessCount = &v
	return s
}

type QueryGenerateDevicesRecordResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryGenerateDevicesRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGenerateDevicesRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGenerateDevicesRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryGenerateDevicesRecordResponse) SetHeaders(v map[string]*string) *QueryGenerateDevicesRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryGenerateDevicesRecordResponse) SetStatusCode(v int32) *QueryGenerateDevicesRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGenerateDevicesRecordResponse) SetBody(v *QueryGenerateDevicesRecordResponseBody) *QueryGenerateDevicesRecordResponse {
	s.Body = v
	return s
}

type QueryStorageCommodityListResponseBody struct {
	Code         *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         []*QueryStorageCommodityListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorMessage *string                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryStorageCommodityListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStorageCommodityListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStorageCommodityListResponseBody) SetCode(v string) *QueryStorageCommodityListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStorageCommodityListResponseBody) SetData(v []*QueryStorageCommodityListResponseBodyData) *QueryStorageCommodityListResponseBody {
	s.Data = v
	return s
}

func (s *QueryStorageCommodityListResponseBody) SetErrorMessage(v string) *QueryStorageCommodityListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryStorageCommodityListResponseBody) SetRequestId(v string) *QueryStorageCommodityListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStorageCommodityListResponseBody) SetSuccess(v bool) *QueryStorageCommodityListResponseBody {
	s.Success = &v
	return s
}

type QueryStorageCommodityListResponseBodyData struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	Lifecycle     *int32  `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	Months        *int32  `json:"Months,omitempty" xml:"Months,omitempty"`
	Price         *string `json:"Price,omitempty" xml:"Price,omitempty"`
	RecordType    *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s QueryStorageCommodityListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryStorageCommodityListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryStorageCommodityListResponseBodyData) SetCommodityCode(v string) *QueryStorageCommodityListResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *QueryStorageCommodityListResponseBodyData) SetCommodityName(v string) *QueryStorageCommodityListResponseBodyData {
	s.CommodityName = &v
	return s
}

func (s *QueryStorageCommodityListResponseBodyData) SetLifecycle(v int32) *QueryStorageCommodityListResponseBodyData {
	s.Lifecycle = &v
	return s
}

func (s *QueryStorageCommodityListResponseBodyData) SetMonths(v int32) *QueryStorageCommodityListResponseBodyData {
	s.Months = &v
	return s
}

func (s *QueryStorageCommodityListResponseBodyData) SetPrice(v string) *QueryStorageCommodityListResponseBodyData {
	s.Price = &v
	return s
}

func (s *QueryStorageCommodityListResponseBodyData) SetRecordType(v int32) *QueryStorageCommodityListResponseBodyData {
	s.RecordType = &v
	return s
}

func (s *QueryStorageCommodityListResponseBodyData) SetSpecification(v string) *QueryStorageCommodityListResponseBodyData {
	s.Specification = &v
	return s
}

type QueryStorageCommodityListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryStorageCommodityListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryStorageCommodityListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStorageCommodityListResponse) GoString() string {
	return s.String()
}

func (s *QueryStorageCommodityListResponse) SetHeaders(v map[string]*string) *QueryStorageCommodityListResponse {
	s.Headers = v
	return s
}

func (s *QueryStorageCommodityListResponse) SetStatusCode(v int32) *QueryStorageCommodityListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStorageCommodityListResponse) SetBody(v *QueryStorageCommodityListResponseBody) *QueryStorageCommodityListResponse {
	s.Body = v
	return s
}

type QueryStorageOrderRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceNoOwner *bool   `json:"DeviceNoOwner,omitempty" xml:"DeviceNoOwner,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryStorageOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStorageOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryStorageOrderRequest) SetDeviceName(v string) *QueryStorageOrderRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryStorageOrderRequest) SetDeviceNoOwner(v bool) *QueryStorageOrderRequest {
	s.DeviceNoOwner = &v
	return s
}

func (s *QueryStorageOrderRequest) SetIotId(v string) *QueryStorageOrderRequest {
	s.IotId = &v
	return s
}

func (s *QueryStorageOrderRequest) SetOrderId(v string) *QueryStorageOrderRequest {
	s.OrderId = &v
	return s
}

func (s *QueryStorageOrderRequest) SetProductKey(v string) *QueryStorageOrderRequest {
	s.ProductKey = &v
	return s
}

type QueryStorageOrderResponseBody struct {
	Code         *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryStorageOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryStorageOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStorageOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStorageOrderResponseBody) SetCode(v string) *QueryStorageOrderResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStorageOrderResponseBody) SetData(v *QueryStorageOrderResponseBodyData) *QueryStorageOrderResponseBody {
	s.Data = v
	return s
}

func (s *QueryStorageOrderResponseBody) SetErrorMessage(v string) *QueryStorageOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryStorageOrderResponseBody) SetRequestId(v string) *QueryStorageOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStorageOrderResponseBody) SetSuccess(v bool) *QueryStorageOrderResponseBody {
	s.Success = &v
	return s
}

type QueryStorageOrderResponseBodyData struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Copies        *int32  `json:"Copies,omitempty" xml:"Copies,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC    *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	IdentityId    *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderType     *int32  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OutOrderNo    *string `json:"OutOrderNo,omitempty" xml:"OutOrderNo,omitempty"`
	PaymentStatus *int32  `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	PreConsume    *int32  `json:"PreConsume,omitempty" xml:"PreConsume,omitempty"`
	Price         *string `json:"Price,omitempty" xml:"Price,omitempty"`
	RecordType    *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC  *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s QueryStorageOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryStorageOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryStorageOrderResponseBodyData) SetCommodityCode(v string) *QueryStorageOrderResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetCopies(v int32) *QueryStorageOrderResponseBodyData {
	s.Copies = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetEndTime(v string) *QueryStorageOrderResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetEndTimeUTC(v string) *QueryStorageOrderResponseBodyData {
	s.EndTimeUTC = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetIdentityId(v string) *QueryStorageOrderResponseBodyData {
	s.IdentityId = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetIotId(v string) *QueryStorageOrderResponseBodyData {
	s.IotId = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetOrderId(v string) *QueryStorageOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetOrderType(v int32) *QueryStorageOrderResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetOutOrderNo(v string) *QueryStorageOrderResponseBodyData {
	s.OutOrderNo = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetPaymentStatus(v int32) *QueryStorageOrderResponseBodyData {
	s.PaymentStatus = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetPreConsume(v int32) *QueryStorageOrderResponseBodyData {
	s.PreConsume = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetPrice(v string) *QueryStorageOrderResponseBodyData {
	s.Price = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetRecordType(v int32) *QueryStorageOrderResponseBodyData {
	s.RecordType = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetSpecification(v string) *QueryStorageOrderResponseBodyData {
	s.Specification = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetStartTime(v string) *QueryStorageOrderResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetStartTimeUTC(v string) *QueryStorageOrderResponseBodyData {
	s.StartTimeUTC = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetStatus(v int32) *QueryStorageOrderResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetUserId(v string) *QueryStorageOrderResponseBodyData {
	s.UserId = &v
	return s
}

func (s *QueryStorageOrderResponseBodyData) SetUserName(v string) *QueryStorageOrderResponseBodyData {
	s.UserName = &v
	return s
}

type QueryStorageOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryStorageOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryStorageOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStorageOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryStorageOrderResponse) SetHeaders(v map[string]*string) *QueryStorageOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryStorageOrderResponse) SetStatusCode(v int32) *QueryStorageOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStorageOrderResponse) SetBody(v *QueryStorageOrderResponseBody) *QueryStorageOrderResponse {
	s.Body = v
	return s
}

type QueryStorageOrderListRequest struct {
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceNoOwner *bool   `json:"DeviceNoOwner,omitempty" xml:"DeviceNoOwner,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	PageNo        *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductKey    *string `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
}

func (s QueryStorageOrderListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryStorageOrderListRequest) GoString() string {
	return s.String()
}

func (s *QueryStorageOrderListRequest) SetDeviceName(v string) *QueryStorageOrderListRequest {
	s.DeviceName = &v
	return s
}

func (s *QueryStorageOrderListRequest) SetDeviceNoOwner(v bool) *QueryStorageOrderListRequest {
	s.DeviceNoOwner = &v
	return s
}

func (s *QueryStorageOrderListRequest) SetIotId(v string) *QueryStorageOrderListRequest {
	s.IotId = &v
	return s
}

func (s *QueryStorageOrderListRequest) SetPageNo(v int32) *QueryStorageOrderListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryStorageOrderListRequest) SetPageSize(v int32) *QueryStorageOrderListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryStorageOrderListRequest) SetProductKey(v string) *QueryStorageOrderListRequest {
	s.ProductKey = &v
	return s
}

type QueryStorageOrderListResponseBody struct {
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *QueryStorageOrderListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryStorageOrderListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryStorageOrderListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStorageOrderListResponseBody) SetCode(v string) *QueryStorageOrderListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryStorageOrderListResponseBody) SetData(v *QueryStorageOrderListResponseBodyData) *QueryStorageOrderListResponseBody {
	s.Data = v
	return s
}

func (s *QueryStorageOrderListResponseBody) SetErrorMessage(v string) *QueryStorageOrderListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryStorageOrderListResponseBody) SetRequestId(v string) *QueryStorageOrderListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStorageOrderListResponseBody) SetSuccess(v bool) *QueryStorageOrderListResponseBody {
	s.Success = &v
	return s
}

type QueryStorageOrderListResponseBodyData struct {
	PageCount        *int32                                                   `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	PageNo           *int32                                                   `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize         *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StorageOrderList []*QueryStorageOrderListResponseBodyDataStorageOrderList `json:"StorageOrderList,omitempty" xml:"StorageOrderList,omitempty" type:"Repeated"`
	Total            *int32                                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryStorageOrderListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryStorageOrderListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryStorageOrderListResponseBodyData) SetPageCount(v int32) *QueryStorageOrderListResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyData) SetPageNo(v int32) *QueryStorageOrderListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyData) SetPageSize(v int32) *QueryStorageOrderListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyData) SetStorageOrderList(v []*QueryStorageOrderListResponseBodyDataStorageOrderList) *QueryStorageOrderListResponseBodyData {
	s.StorageOrderList = v
	return s
}

func (s *QueryStorageOrderListResponseBodyData) SetTotal(v int32) *QueryStorageOrderListResponseBodyData {
	s.Total = &v
	return s
}

type QueryStorageOrderListResponseBodyDataStorageOrderList struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Copies        *int32  `json:"Copies,omitempty" xml:"Copies,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC    *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	IdentityId    *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderType     *int32  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OutOrderNo    *string `json:"OutOrderNo,omitempty" xml:"OutOrderNo,omitempty"`
	PaymentStatus *int32  `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	PreConsume    *int32  `json:"PreConsume,omitempty" xml:"PreConsume,omitempty"`
	Price         *string `json:"Price,omitempty" xml:"Price,omitempty"`
	RecordType    *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC  *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s QueryStorageOrderListResponseBodyDataStorageOrderList) String() string {
	return tea.Prettify(s)
}

func (s QueryStorageOrderListResponseBodyDataStorageOrderList) GoString() string {
	return s.String()
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetCommodityCode(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.CommodityCode = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetCopies(v int32) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.Copies = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetEndTime(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.EndTime = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetEndTimeUTC(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.EndTimeUTC = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetIdentityId(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.IdentityId = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetIotId(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.IotId = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetOrderId(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.OrderId = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetOrderType(v int32) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.OrderType = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetOutOrderNo(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.OutOrderNo = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetPaymentStatus(v int32) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.PaymentStatus = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetPreConsume(v int32) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.PreConsume = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetPrice(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.Price = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetRecordType(v int32) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.RecordType = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetSpecification(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.Specification = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetStartTime(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.StartTime = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetStartTimeUTC(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.StartTimeUTC = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetStatus(v int32) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.Status = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetUserId(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.UserId = &v
	return s
}

func (s *QueryStorageOrderListResponseBodyDataStorageOrderList) SetUserName(v string) *QueryStorageOrderListResponseBodyDataStorageOrderList {
	s.UserName = &v
	return s
}

type QueryStorageOrderListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryStorageOrderListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryStorageOrderListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryStorageOrderListResponse) GoString() string {
	return s.String()
}

func (s *QueryStorageOrderListResponse) SetHeaders(v map[string]*string) *QueryStorageOrderListResponse {
	s.Headers = v
	return s
}

func (s *QueryStorageOrderListResponse) SetStatusCode(v int32) *QueryStorageOrderListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryStorageOrderListResponse) SetBody(v *QueryStorageOrderListResponseBody) *QueryStorageOrderListResponse {
	s.Body = v
	return s
}

type TransferStorageOrderRequest struct {
	DstIotId                     *string `json:"DstIotId,omitempty" xml:"DstIotId,omitempty"`
	EnableDefaultPlan            *bool   `json:"EnableDefaultPlan,omitempty" xml:"EnableDefaultPlan,omitempty"`
	EventRecordDuration          *int32  `json:"EventRecordDuration,omitempty" xml:"EventRecordDuration,omitempty"`
	EventRecordProlong           *bool   `json:"EventRecordProlong,omitempty" xml:"EventRecordProlong,omitempty"`
	ImmediateUse                 *bool   `json:"ImmediateUse,omitempty" xml:"ImmediateUse,omitempty"`
	PreRecordDuration            *int32  `json:"PreRecordDuration,omitempty" xml:"PreRecordDuration,omitempty"`
	SrcIotId                     *string `json:"SrcIotId,omitempty" xml:"SrcIotId,omitempty"`
	SrcOrderId                   *string `json:"SrcOrderId,omitempty" xml:"SrcOrderId,omitempty"`
	SupportCrossIdentityTransfer *bool   `json:"SupportCrossIdentityTransfer,omitempty" xml:"SupportCrossIdentityTransfer,omitempty"`
	UserId                       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName                     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s TransferStorageOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferStorageOrderRequest) GoString() string {
	return s.String()
}

func (s *TransferStorageOrderRequest) SetDstIotId(v string) *TransferStorageOrderRequest {
	s.DstIotId = &v
	return s
}

func (s *TransferStorageOrderRequest) SetEnableDefaultPlan(v bool) *TransferStorageOrderRequest {
	s.EnableDefaultPlan = &v
	return s
}

func (s *TransferStorageOrderRequest) SetEventRecordDuration(v int32) *TransferStorageOrderRequest {
	s.EventRecordDuration = &v
	return s
}

func (s *TransferStorageOrderRequest) SetEventRecordProlong(v bool) *TransferStorageOrderRequest {
	s.EventRecordProlong = &v
	return s
}

func (s *TransferStorageOrderRequest) SetImmediateUse(v bool) *TransferStorageOrderRequest {
	s.ImmediateUse = &v
	return s
}

func (s *TransferStorageOrderRequest) SetPreRecordDuration(v int32) *TransferStorageOrderRequest {
	s.PreRecordDuration = &v
	return s
}

func (s *TransferStorageOrderRequest) SetSrcIotId(v string) *TransferStorageOrderRequest {
	s.SrcIotId = &v
	return s
}

func (s *TransferStorageOrderRequest) SetSrcOrderId(v string) *TransferStorageOrderRequest {
	s.SrcOrderId = &v
	return s
}

func (s *TransferStorageOrderRequest) SetSupportCrossIdentityTransfer(v bool) *TransferStorageOrderRequest {
	s.SupportCrossIdentityTransfer = &v
	return s
}

func (s *TransferStorageOrderRequest) SetUserId(v string) *TransferStorageOrderRequest {
	s.UserId = &v
	return s
}

func (s *TransferStorageOrderRequest) SetUserName(v string) *TransferStorageOrderRequest {
	s.UserName = &v
	return s
}

type TransferStorageOrderResponseBody struct {
	Code         *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *TransferStorageOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferStorageOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferStorageOrderResponseBody) GoString() string {
	return s.String()
}

func (s *TransferStorageOrderResponseBody) SetCode(v string) *TransferStorageOrderResponseBody {
	s.Code = &v
	return s
}

func (s *TransferStorageOrderResponseBody) SetData(v *TransferStorageOrderResponseBodyData) *TransferStorageOrderResponseBody {
	s.Data = v
	return s
}

func (s *TransferStorageOrderResponseBody) SetErrorMessage(v string) *TransferStorageOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TransferStorageOrderResponseBody) SetRequestId(v string) *TransferStorageOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferStorageOrderResponseBody) SetSuccess(v bool) *TransferStorageOrderResponseBody {
	s.Success = &v
	return s
}

type TransferStorageOrderResponseBodyData struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Copies        *int32  `json:"Copies,omitempty" xml:"Copies,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTC    *string `json:"EndTimeUTC,omitempty" xml:"EndTimeUTC,omitempty"`
	IdentityId    *string `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	IotId         *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderType     *int32  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OutOrderNo    *string `json:"OutOrderNo,omitempty" xml:"OutOrderNo,omitempty"`
	PaymentStatus *int32  `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	PreConsume    *int32  `json:"PreConsume,omitempty" xml:"PreConsume,omitempty"`
	Price         *string `json:"Price,omitempty" xml:"Price,omitempty"`
	RecordType    *int32  `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimeUTC  *string `json:"StartTimeUTC,omitempty" xml:"StartTimeUTC,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s TransferStorageOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TransferStorageOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *TransferStorageOrderResponseBodyData) SetCommodityCode(v string) *TransferStorageOrderResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetCopies(v int32) *TransferStorageOrderResponseBodyData {
	s.Copies = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetEndTime(v string) *TransferStorageOrderResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetEndTimeUTC(v string) *TransferStorageOrderResponseBodyData {
	s.EndTimeUTC = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetIdentityId(v string) *TransferStorageOrderResponseBodyData {
	s.IdentityId = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetIotId(v string) *TransferStorageOrderResponseBodyData {
	s.IotId = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetOrderId(v string) *TransferStorageOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetOrderType(v int32) *TransferStorageOrderResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetOutOrderNo(v string) *TransferStorageOrderResponseBodyData {
	s.OutOrderNo = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetPaymentStatus(v int32) *TransferStorageOrderResponseBodyData {
	s.PaymentStatus = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetPreConsume(v int32) *TransferStorageOrderResponseBodyData {
	s.PreConsume = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetPrice(v string) *TransferStorageOrderResponseBodyData {
	s.Price = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetRecordType(v int32) *TransferStorageOrderResponseBodyData {
	s.RecordType = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetSpecification(v string) *TransferStorageOrderResponseBodyData {
	s.Specification = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetStartTime(v string) *TransferStorageOrderResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetStartTimeUTC(v string) *TransferStorageOrderResponseBodyData {
	s.StartTimeUTC = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetStatus(v int32) *TransferStorageOrderResponseBodyData {
	s.Status = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetUserId(v string) *TransferStorageOrderResponseBodyData {
	s.UserId = &v
	return s
}

func (s *TransferStorageOrderResponseBodyData) SetUserName(v string) *TransferStorageOrderResponseBodyData {
	s.UserName = &v
	return s
}

type TransferStorageOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferStorageOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferStorageOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferStorageOrderResponse) GoString() string {
	return s.String()
}

func (s *TransferStorageOrderResponse) SetHeaders(v map[string]*string) *TransferStorageOrderResponse {
	s.Headers = v
	return s
}

func (s *TransferStorageOrderResponse) SetStatusCode(v int32) *TransferStorageOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferStorageOrderResponse) SetBody(v *TransferStorageOrderResponseBody) *TransferStorageOrderResponse {
	s.Body = v
	return s
}

type UploadDeviceNameListRequest struct {
	DeviceNames []*string `json:"DeviceNames,omitempty" xml:"DeviceNames,omitempty" type:"Repeated"`
	ProductKey  *string   `json:"ProductKey,omitempty" xml:"ProductKey,omitempty"`
	ProjectId   *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UploadDeviceNameListRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadDeviceNameListRequest) GoString() string {
	return s.String()
}

func (s *UploadDeviceNameListRequest) SetDeviceNames(v []*string) *UploadDeviceNameListRequest {
	s.DeviceNames = v
	return s
}

func (s *UploadDeviceNameListRequest) SetProductKey(v string) *UploadDeviceNameListRequest {
	s.ProductKey = &v
	return s
}

func (s *UploadDeviceNameListRequest) SetProjectId(v string) *UploadDeviceNameListRequest {
	s.ProjectId = &v
	return s
}

type UploadDeviceNameListResponseBody struct {
	Code         *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         *UploadDeviceNameListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadDeviceNameListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadDeviceNameListResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDeviceNameListResponseBody) SetCode(v string) *UploadDeviceNameListResponseBody {
	s.Code = &v
	return s
}

func (s *UploadDeviceNameListResponseBody) SetData(v *UploadDeviceNameListResponseBodyData) *UploadDeviceNameListResponseBody {
	s.Data = v
	return s
}

func (s *UploadDeviceNameListResponseBody) SetErrorMessage(v string) *UploadDeviceNameListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UploadDeviceNameListResponseBody) SetRequestId(v string) *UploadDeviceNameListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDeviceNameListResponseBody) SetSuccess(v bool) *UploadDeviceNameListResponseBody {
	s.Success = &v
	return s
}

type UploadDeviceNameListResponseBodyData struct {
	BatchId                *string                                                  `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	InvalidDetailList      []*UploadDeviceNameListResponseBodyDataInvalidDetailList `json:"InvalidDetailList,omitempty" xml:"InvalidDetailList,omitempty" type:"Repeated"`
	InvalidDeviceNameList  []*string                                                `json:"InvalidDeviceNameList,omitempty" xml:"InvalidDeviceNameList,omitempty" type:"Repeated"`
	RepeatedDeviceNameList []*string                                                `json:"RepeatedDeviceNameList,omitempty" xml:"RepeatedDeviceNameList,omitempty" type:"Repeated"`
}

func (s UploadDeviceNameListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UploadDeviceNameListResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadDeviceNameListResponseBodyData) SetBatchId(v string) *UploadDeviceNameListResponseBodyData {
	s.BatchId = &v
	return s
}

func (s *UploadDeviceNameListResponseBodyData) SetInvalidDetailList(v []*UploadDeviceNameListResponseBodyDataInvalidDetailList) *UploadDeviceNameListResponseBodyData {
	s.InvalidDetailList = v
	return s
}

func (s *UploadDeviceNameListResponseBodyData) SetInvalidDeviceNameList(v []*string) *UploadDeviceNameListResponseBodyData {
	s.InvalidDeviceNameList = v
	return s
}

func (s *UploadDeviceNameListResponseBodyData) SetRepeatedDeviceNameList(v []*string) *UploadDeviceNameListResponseBodyData {
	s.RepeatedDeviceNameList = v
	return s
}

type UploadDeviceNameListResponseBodyDataInvalidDetailList struct {
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
}

func (s UploadDeviceNameListResponseBodyDataInvalidDetailList) String() string {
	return tea.Prettify(s)
}

func (s UploadDeviceNameListResponseBodyDataInvalidDetailList) GoString() string {
	return s.String()
}

func (s *UploadDeviceNameListResponseBodyDataInvalidDetailList) SetDeviceName(v string) *UploadDeviceNameListResponseBodyDataInvalidDetailList {
	s.DeviceName = &v
	return s
}

func (s *UploadDeviceNameListResponseBodyDataInvalidDetailList) SetErrorMsg(v string) *UploadDeviceNameListResponseBodyDataInvalidDetailList {
	s.ErrorMsg = &v
	return s
}

type UploadDeviceNameListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadDeviceNameListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadDeviceNameListResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadDeviceNameListResponse) GoString() string {
	return s.String()
}

func (s *UploadDeviceNameListResponse) SetHeaders(v map[string]*string) *UploadDeviceNameListResponse {
	s.Headers = v
	return s
}

func (s *UploadDeviceNameListResponse) SetStatusCode(v int32) *UploadDeviceNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDeviceNameListResponse) SetBody(v *UploadDeviceNameListResponseBody) *UploadDeviceNameListResponse {
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
		"ap-northeast-1":              tea.String("linkvisual.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("linkvisual.aliyuncs.com"),
		"ap-south-1":                  tea.String("linkvisual.aliyuncs.com"),
		"ap-southeast-1":              tea.String("linkvisual.aliyuncs.com"),
		"ap-southeast-2":              tea.String("linkvisual.aliyuncs.com"),
		"ap-southeast-3":              tea.String("linkvisual.aliyuncs.com"),
		"ap-southeast-5":              tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing":                  tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("linkvisual.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("linkvisual.aliyuncs.com"),
		"cn-chengdu":                  tea.String("linkvisual.aliyuncs.com"),
		"cn-edge-1":                   tea.String("linkvisual.aliyuncs.com"),
		"cn-fujian":                   tea.String("linkvisual.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("linkvisual.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("linkvisual.aliyuncs.com"),
		"cn-hongkong":                 tea.String("linkvisual.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("linkvisual.aliyuncs.com"),
		"cn-huhehaote":                tea.String("linkvisual.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("linkvisual.aliyuncs.com"),
		"cn-qingdao":                  tea.String("linkvisual.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("linkvisual.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("linkvisual.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("linkvisual.aliyuncs.com"),
		"cn-wuhan":                    tea.String("linkvisual.aliyuncs.com"),
		"cn-yushanfang":               tea.String("linkvisual.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("linkvisual.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("linkvisual.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("linkvisual.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("linkvisual.aliyuncs.com"),
		"eu-central-1":                tea.String("linkvisual.aliyuncs.com"),
		"eu-west-1":                   tea.String("linkvisual.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("linkvisual.aliyuncs.com"),
		"me-east-1":                   tea.String("linkvisual.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("linkvisual.aliyuncs.com"),
		"us-east-1":                   tea.String("linkvisual.aliyuncs.com"),
		"us-west-1":                   tea.String("linkvisual.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("linkvisual"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BindStorageOrderWithOptions(request *BindStorageOrderRequest, runtime *util.RuntimeOptions) (_result *BindStorageOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableDefaultPlan)) {
		query["EnableDefaultPlan"] = request.EnableDefaultPlan
	}

	if !tea.BoolValue(util.IsUnset(request.EventRecordDuration)) {
		query["EventRecordDuration"] = request.EventRecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.EventRecordProlong)) {
		query["EventRecordProlong"] = request.EventRecordProlong
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxRecordFileDuration)) {
		query["MaxRecordFileDuration"] = request.MaxRecordFileDuration
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.PreRecordDuration)) {
		query["PreRecordDuration"] = request.PreRecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindStorageOrder"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindStorageOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindStorageOrder(request *BindStorageOrderRequest) (_result *BindStorageOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindStorageOrderResponse{}
	_body, _err := client.BindStorageOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckFreeStorageValidWithOptions(request *CheckFreeStorageValidRequest, runtime *util.RuntimeOptions) (_result *CheckFreeStorageValidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckFreeStorageValid"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckFreeStorageValidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckFreeStorageValid(request *CheckFreeStorageValidRequest) (_result *CheckFreeStorageValidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckFreeStorageValidResponse{}
	_body, _err := client.CheckFreeStorageValidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConsumeFreeStorageWithOptions(request *ConsumeFreeStorageRequest, runtime *util.RuntimeOptions) (_result *ConsumeFreeStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.EnableDefaultPlan)) {
		query["EnableDefaultPlan"] = request.EnableDefaultPlan
	}

	if !tea.BoolValue(util.IsUnset(request.EventRecordDuration)) {
		query["EventRecordDuration"] = request.EventRecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.EventRecordProlong)) {
		query["EventRecordProlong"] = request.EventRecordProlong
	}

	if !tea.BoolValue(util.IsUnset(request.ImmediateUse)) {
		query["ImmediateUse"] = request.ImmediateUse
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.PreRecordDuration)) {
		query["PreRecordDuration"] = request.PreRecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Quota)) {
		query["Quota"] = request.Quota
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConsumeFreeStorage"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConsumeFreeStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConsumeFreeStorage(request *ConsumeFreeStorageRequest) (_result *ConsumeFreeStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConsumeFreeStorageResponse{}
	_body, _err := client.ConsumeFreeStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAndPayStorageOrderWithOptions(request *CreateAndPayStorageOrderRequest, runtime *util.RuntimeOptions) (_result *CreateAndPayStorageOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.Copies)) {
		query["Copies"] = request.Copies
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceNoOwner)) {
		query["DeviceNoOwner"] = request.DeviceNoOwner
	}

	if !tea.BoolValue(util.IsUnset(request.EnableDefaultPlan)) {
		query["EnableDefaultPlan"] = request.EnableDefaultPlan
	}

	if !tea.BoolValue(util.IsUnset(request.EventRecordDuration)) {
		query["EventRecordDuration"] = request.EventRecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.EventRecordProlong)) {
		query["EventRecordProlong"] = request.EventRecordProlong
	}

	if !tea.BoolValue(util.IsUnset(request.ImmediateUse)) {
		query["ImmediateUse"] = request.ImmediateUse
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxRecordFileDuration)) {
		query["MaxRecordFileDuration"] = request.MaxRecordFileDuration
	}

	if !tea.BoolValue(util.IsUnset(request.PreRecordDuration)) {
		query["PreRecordDuration"] = request.PreRecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.Specification)) {
		query["Specification"] = request.Specification
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAndPayStorageOrder"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAndPayStorageOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAndPayStorageOrder(request *CreateAndPayStorageOrderRequest) (_result *CreateAndPayStorageOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAndPayStorageOrderResponse{}
	_body, _err := client.CreateAndPayStorageOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableFreeStorageWithOptions(request *EnableFreeStorageRequest, runtime *util.RuntimeOptions) (_result *EnableFreeStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableFreeStorage"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableFreeStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableFreeStorage(request *EnableFreeStorageRequest) (_result *EnableFreeStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableFreeStorageResponse{}
	_body, _err := client.EnableFreeStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableStorageOrderWithOptions(request *EnableStorageOrderRequest, runtime *util.RuntimeOptions) (_result *EnableStorageOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableStorageOrder"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableStorageOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableStorageOrder(request *EnableStorageOrderRequest) (_result *EnableStorageOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableStorageOrderResponse{}
	_body, _err := client.EnableStorageOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FreezeFreeStorageWithOptions(request *FreezeFreeStorageRequest, runtime *util.RuntimeOptions) (_result *FreezeFreeStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FreezeFreeStorage"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FreezeFreeStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FreezeFreeStorage(request *FreezeFreeStorageRequest) (_result *FreezeFreeStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FreezeFreeStorageResponse{}
	_body, _err := client.FreezeFreeStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FreezeStorageOrderWithOptions(request *FreezeStorageOrderRequest, runtime *util.RuntimeOptions) (_result *FreezeStorageOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceNoOwner)) {
		query["DeviceNoOwner"] = request.DeviceNoOwner
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FreezeStorageOrder"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FreezeStorageOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FreezeStorageOrder(request *FreezeStorageOrderRequest) (_result *FreezeStorageOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FreezeStorageOrderResponse{}
	_body, _err := client.FreezeStorageOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateDeviceWithOptions(request *GenerateDeviceRequest, runtime *util.RuntimeOptions) (_result *GenerateDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateDevice"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateDevice(request *GenerateDeviceRequest) (_result *GenerateDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateDeviceResponse{}
	_body, _err := client.GenerateDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateDeviceByBatchIdWithOptions(request *GenerateDeviceByBatchIdRequest, runtime *util.RuntimeOptions) (_result *GenerateDeviceByBatchIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchId)) {
		query["BatchId"] = request.BatchId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateDeviceByBatchId"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateDeviceByBatchIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateDeviceByBatchId(request *GenerateDeviceByBatchIdRequest) (_result *GenerateDeviceByBatchIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateDeviceByBatchIdResponse{}
	_body, _err := client.GenerateDeviceByBatchIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryBatchStatusWithOptions(request *QueryBatchStatusRequest, runtime *util.RuntimeOptions) (_result *QueryBatchStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchId)) {
		query["BatchId"] = request.BatchId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBatchStatus"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBatchStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryBatchStatus(request *QueryBatchStatusRequest) (_result *QueryBatchStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBatchStatusResponse{}
	_body, _err := client.QueryBatchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryDevicesDownloadUrlWithOptions(request *QueryDevicesDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *QueryDevicesDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchId)) {
		query["BatchId"] = request.BatchId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDevicesDownloadUrl"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDevicesDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryDevicesDownloadUrl(request *QueryDevicesDownloadUrlRequest) (_result *QueryDevicesDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDevicesDownloadUrlResponse{}
	_body, _err := client.QueryDevicesDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryFreeStorageWithOptions(request *QueryFreeStorageRequest, runtime *util.RuntimeOptions) (_result *QueryFreeStorageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryFreeStorage"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryFreeStorageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryFreeStorage(request *QueryFreeStorageRequest) (_result *QueryFreeStorageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryFreeStorageResponse{}
	_body, _err := client.QueryFreeStorageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGenerateDevicesInfoListWithOptions(request *QueryGenerateDevicesInfoListRequest, runtime *util.RuntimeOptions) (_result *QueryGenerateDevicesInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BatchId)) {
		query["BatchId"] = request.BatchId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryGenerateDevicesInfoList"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGenerateDevicesInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGenerateDevicesInfoList(request *QueryGenerateDevicesInfoListRequest) (_result *QueryGenerateDevicesInfoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGenerateDevicesInfoListResponse{}
	_body, _err := client.QueryGenerateDevicesInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGenerateDevicesRecordWithOptions(request *QueryGenerateDevicesRecordRequest, runtime *util.RuntimeOptions) (_result *QueryGenerateDevicesRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryGenerateDevicesRecord"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGenerateDevicesRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGenerateDevicesRecord(request *QueryGenerateDevicesRecordRequest) (_result *QueryGenerateDevicesRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGenerateDevicesRecordResponse{}
	_body, _err := client.QueryGenerateDevicesRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStorageCommodityListWithOptions(runtime *util.RuntimeOptions) (_result *QueryStorageCommodityListResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QueryStorageCommodityList"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStorageCommodityListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStorageCommodityList() (_result *QueryStorageCommodityListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryStorageCommodityListResponse{}
	_body, _err := client.QueryStorageCommodityListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStorageOrderWithOptions(request *QueryStorageOrderRequest, runtime *util.RuntimeOptions) (_result *QueryStorageOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceNoOwner)) {
		query["DeviceNoOwner"] = request.DeviceNoOwner
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryStorageOrder"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStorageOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStorageOrder(request *QueryStorageOrderRequest) (_result *QueryStorageOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryStorageOrderResponse{}
	_body, _err := client.QueryStorageOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryStorageOrderListWithOptions(request *QueryStorageOrderListRequest, runtime *util.RuntimeOptions) (_result *QueryStorageOrderListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceName)) {
		query["DeviceName"] = request.DeviceName
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceNoOwner)) {
		query["DeviceNoOwner"] = request.DeviceNoOwner
	}

	if !tea.BoolValue(util.IsUnset(request.IotId)) {
		query["IotId"] = request.IotId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryStorageOrderList"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryStorageOrderListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryStorageOrderList(request *QueryStorageOrderListRequest) (_result *QueryStorageOrderListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryStorageOrderListResponse{}
	_body, _err := client.QueryStorageOrderListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferStorageOrderWithOptions(request *TransferStorageOrderRequest, runtime *util.RuntimeOptions) (_result *TransferStorageOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DstIotId)) {
		query["DstIotId"] = request.DstIotId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableDefaultPlan)) {
		query["EnableDefaultPlan"] = request.EnableDefaultPlan
	}

	if !tea.BoolValue(util.IsUnset(request.EventRecordDuration)) {
		query["EventRecordDuration"] = request.EventRecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.EventRecordProlong)) {
		query["EventRecordProlong"] = request.EventRecordProlong
	}

	if !tea.BoolValue(util.IsUnset(request.ImmediateUse)) {
		query["ImmediateUse"] = request.ImmediateUse
	}

	if !tea.BoolValue(util.IsUnset(request.PreRecordDuration)) {
		query["PreRecordDuration"] = request.PreRecordDuration
	}

	if !tea.BoolValue(util.IsUnset(request.SrcIotId)) {
		query["SrcIotId"] = request.SrcIotId
	}

	if !tea.BoolValue(util.IsUnset(request.SrcOrderId)) {
		query["SrcOrderId"] = request.SrcOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.SupportCrossIdentityTransfer)) {
		query["SupportCrossIdentityTransfer"] = request.SupportCrossIdentityTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferStorageOrder"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferStorageOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferStorageOrder(request *TransferStorageOrderRequest) (_result *TransferStorageOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferStorageOrderResponse{}
	_body, _err := client.TransferStorageOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadDeviceNameListWithOptions(request *UploadDeviceNameListRequest, runtime *util.RuntimeOptions) (_result *UploadDeviceNameListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductKey)) {
		query["ProductKey"] = request.ProductKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceNames)) {
		body["DeviceNames"] = request.DeviceNames
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadDeviceNameList"),
		Version:     tea.String("2023-06-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadDeviceNameListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadDeviceNameList(request *UploadDeviceNameListRequest) (_result *UploadDeviceNameListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadDeviceNameListResponse{}
	_body, _err := client.UploadDeviceNameListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
