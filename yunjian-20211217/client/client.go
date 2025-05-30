// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateDemandPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 262940
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s CreateDemandPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateDemandPlanHeaders) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanHeaders) SetCommonHeaders(v map[string]*string) *CreateDemandPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDemandPlanHeaders) SetYunUserId(v string) *CreateDemandPlanHeaders {
	s.YunUserId = &v
	return s
}

type CreateDemandPlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1065737167271819
	AccountId   *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FY2022
	Period    *string `json:"period,omitempty" xml:"period,omitempty"`
	Source    *string `json:"source,omitempty" xml:"source,omitempty"`
	TargetCid *int64  `json:"targetCid,omitempty" xml:"targetCid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// URGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 262940
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateDemandPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDemandPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanRequest) SetAccountId(v string) *CreateDemandPlanRequest {
	s.AccountId = &v
	return s
}

func (s *CreateDemandPlanRequest) SetDescription(v string) *CreateDemandPlanRequest {
	s.Description = &v
	return s
}

func (s *CreateDemandPlanRequest) SetName(v string) *CreateDemandPlanRequest {
	s.Name = &v
	return s
}

func (s *CreateDemandPlanRequest) SetPeriod(v string) *CreateDemandPlanRequest {
	s.Period = &v
	return s
}

func (s *CreateDemandPlanRequest) SetSource(v string) *CreateDemandPlanRequest {
	s.Source = &v
	return s
}

func (s *CreateDemandPlanRequest) SetTargetCid(v int64) *CreateDemandPlanRequest {
	s.TargetCid = &v
	return s
}

func (s *CreateDemandPlanRequest) SetType(v string) *CreateDemandPlanRequest {
	s.Type = &v
	return s
}

func (s *CreateDemandPlanRequest) SetUserId(v string) *CreateDemandPlanRequest {
	s.UserId = &v
	return s
}

type CreateDemandPlanResponseBody struct {
	// code
	//
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 111223
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 1e2b798516402440016572132e1459
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CreateDemandPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDemandPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanResponseBody) SetCode(v int64) *CreateDemandPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDemandPlanResponseBody) SetData(v int64) *CreateDemandPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDemandPlanResponseBody) SetMessage(v string) *CreateDemandPlanResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDemandPlanResponseBody) SetSuccess(v bool) *CreateDemandPlanResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDemandPlanResponseBody) SetTraceId(v string) *CreateDemandPlanResponseBody {
	s.TraceId = &v
	return s
}

type CreateDemandPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDemandPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDemandPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDemandPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanResponse) SetHeaders(v map[string]*string) *CreateDemandPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateDemandPlanResponse) SetStatusCode(v int32) *CreateDemandPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDemandPlanResponse) SetBody(v *CreateDemandPlanResponseBody) *CreateDemandPlanResponse {
	s.Body = v
	return s
}

type CreateDemandPlanV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s CreateDemandPlanV2Headers) String() string {
	return tea.Prettify(s)
}

func (s CreateDemandPlanV2Headers) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanV2Headers) SetCommonHeaders(v map[string]*string) *CreateDemandPlanV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *CreateDemandPlanV2Headers) SetYunUserId(v string) *CreateDemandPlanV2Headers {
	s.YunUserId = &v
	return s
}

type CreateDemandPlanV2Request struct {
	// This parameter is required.
	AccountId   *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	TargetCid   *int64  `json:"targetCid,omitempty" xml:"targetCid,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateDemandPlanV2Request) String() string {
	return tea.Prettify(s)
}

func (s CreateDemandPlanV2Request) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanV2Request) SetAccountId(v string) *CreateDemandPlanV2Request {
	s.AccountId = &v
	return s
}

func (s *CreateDemandPlanV2Request) SetDescription(v string) *CreateDemandPlanV2Request {
	s.Description = &v
	return s
}

func (s *CreateDemandPlanV2Request) SetName(v string) *CreateDemandPlanV2Request {
	s.Name = &v
	return s
}

func (s *CreateDemandPlanV2Request) SetProductType(v string) *CreateDemandPlanV2Request {
	s.ProductType = &v
	return s
}

func (s *CreateDemandPlanV2Request) SetTargetCid(v int64) *CreateDemandPlanV2Request {
	s.TargetCid = &v
	return s
}

func (s *CreateDemandPlanV2Request) SetType(v string) *CreateDemandPlanV2Request {
	s.Type = &v
	return s
}

func (s *CreateDemandPlanV2Request) SetUserId(v string) *CreateDemandPlanV2Request {
	s.UserId = &v
	return s
}

type CreateDemandPlanV2ResponseBody struct {
	Code    *int64  `json:"code,omitempty" xml:"code,omitempty"`
	Data    *int64  `json:"data,omitempty" xml:"data,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CreateDemandPlanV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDemandPlanV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanV2ResponseBody) SetCode(v int64) *CreateDemandPlanV2ResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDemandPlanV2ResponseBody) SetData(v int64) *CreateDemandPlanV2ResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDemandPlanV2ResponseBody) SetMessage(v string) *CreateDemandPlanV2ResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDemandPlanV2ResponseBody) SetSuccess(v bool) *CreateDemandPlanV2ResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDemandPlanV2ResponseBody) SetTraceId(v string) *CreateDemandPlanV2ResponseBody {
	s.TraceId = &v
	return s
}

type CreateDemandPlanV2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDemandPlanV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDemandPlanV2Response) String() string {
	return tea.Prettify(s)
}

func (s CreateDemandPlanV2Response) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanV2Response) SetHeaders(v map[string]*string) *CreateDemandPlanV2Response {
	s.Headers = v
	return s
}

func (s *CreateDemandPlanV2Response) SetStatusCode(v int32) *CreateDemandPlanV2Response {
	s.StatusCode = &v
	return s
}

func (s *CreateDemandPlanV2Response) SetBody(v *CreateDemandPlanV2ResponseBody) *CreateDemandPlanV2Response {
	s.Body = v
	return s
}

type DeleteUrgentDemandItemHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s DeleteUrgentDemandItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteUrgentDemandItemHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandItemHeaders) SetCommonHeaders(v map[string]*string) *DeleteUrgentDemandItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUrgentDemandItemHeaders) SetYunUserId(v string) *DeleteUrgentDemandItemHeaders {
	s.YunUserId = &v
	return s
}

type DeleteUrgentDemandItemRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111222
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
}

func (s DeleteUrgentDemandItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUrgentDemandItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandItemRequest) SetId(v int64) *DeleteUrgentDemandItemRequest {
	s.Id = &v
	return s
}

func (s *DeleteUrgentDemandItemRequest) SetModifier(v string) *DeleteUrgentDemandItemRequest {
	s.Modifier = &v
	return s
}

type DeleteUrgentDemandItemResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 1
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 212cf01016405759151137225e83cd
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DeleteUrgentDemandItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUrgentDemandItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandItemResponseBody) SetCode(v int64) *DeleteUrgentDemandItemResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUrgentDemandItemResponseBody) SetData(v int64) *DeleteUrgentDemandItemResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteUrgentDemandItemResponseBody) SetMessage(v string) *DeleteUrgentDemandItemResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUrgentDemandItemResponseBody) SetSuccess(v bool) *DeleteUrgentDemandItemResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUrgentDemandItemResponseBody) SetTraceId(v string) *DeleteUrgentDemandItemResponseBody {
	s.TraceId = &v
	return s
}

type DeleteUrgentDemandItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUrgentDemandItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUrgentDemandItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUrgentDemandItemResponse) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandItemResponse) SetHeaders(v map[string]*string) *DeleteUrgentDemandItemResponse {
	s.Headers = v
	return s
}

func (s *DeleteUrgentDemandItemResponse) SetStatusCode(v int32) *DeleteUrgentDemandItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUrgentDemandItemResponse) SetBody(v *DeleteUrgentDemandItemResponseBody) *DeleteUrgentDemandItemResponse {
	s.Body = v
	return s
}

type DeleteUrgentDemandPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s DeleteUrgentDemandPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeleteUrgentDemandPlanHeaders) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandPlanHeaders) SetCommonHeaders(v map[string]*string) *DeleteUrgentDemandPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteUrgentDemandPlanHeaders) SetYunUserId(v string) *DeleteUrgentDemandPlanHeaders {
	s.YunUserId = &v
	return s
}

type DeleteUrgentDemandPlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111111
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 222111
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
}

func (s DeleteUrgentDemandPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUrgentDemandPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandPlanRequest) SetId(v int64) *DeleteUrgentDemandPlanRequest {
	s.Id = &v
	return s
}

func (s *DeleteUrgentDemandPlanRequest) SetModifier(v string) *DeleteUrgentDemandPlanRequest {
	s.Modifier = &v
	return s
}

type DeleteUrgentDemandPlanResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 1
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 212cf01016405759151137225e83cd
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DeleteUrgentDemandPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUrgentDemandPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandPlanResponseBody) SetCode(v int64) *DeleteUrgentDemandPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUrgentDemandPlanResponseBody) SetData(v int64) *DeleteUrgentDemandPlanResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteUrgentDemandPlanResponseBody) SetMessage(v string) *DeleteUrgentDemandPlanResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUrgentDemandPlanResponseBody) SetSuccess(v bool) *DeleteUrgentDemandPlanResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUrgentDemandPlanResponseBody) SetTraceId(v string) *DeleteUrgentDemandPlanResponseBody {
	s.TraceId = &v
	return s
}

type DeleteUrgentDemandPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUrgentDemandPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUrgentDemandPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUrgentDemandPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandPlanResponse) SetHeaders(v map[string]*string) *DeleteUrgentDemandPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteUrgentDemandPlanResponse) SetStatusCode(v int32) *DeleteUrgentDemandPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUrgentDemandPlanResponse) SetBody(v *DeleteUrgentDemandPlanResponseBody) *DeleteUrgentDemandPlanResponse {
	s.Body = v
	return s
}

type DeliveryItemDetailSynHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s DeliveryItemDetailSynHeaders) String() string {
	return tea.Prettify(s)
}

func (s DeliveryItemDetailSynHeaders) GoString() string {
	return s.String()
}

func (s *DeliveryItemDetailSynHeaders) SetCommonHeaders(v map[string]*string) *DeliveryItemDetailSynHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeliveryItemDetailSynHeaders) SetYunUserId(v string) *DeliveryItemDetailSynHeaders {
	s.YunUserId = &v
	return s
}

type DeliveryItemDetailSynRequest struct {
	Channel                *string                                               `json:"channel,omitempty" xml:"channel,omitempty"`
	DeliveryItemDetailDTOS []*DeliveryItemDetailSynRequestDeliveryItemDetailDTOS `json:"deliveryItemDetailDTOS,omitempty" xml:"deliveryItemDetailDTOS,omitempty" type:"Repeated"`
	DeliveryItemId         *string                                               `json:"deliveryItemId,omitempty" xml:"deliveryItemId,omitempty"`
	DeliveryPlanId         *string                                               `json:"deliveryPlanId,omitempty" xml:"deliveryPlanId,omitempty"`
}

func (s DeliveryItemDetailSynRequest) String() string {
	return tea.Prettify(s)
}

func (s DeliveryItemDetailSynRequest) GoString() string {
	return s.String()
}

func (s *DeliveryItemDetailSynRequest) SetChannel(v string) *DeliveryItemDetailSynRequest {
	s.Channel = &v
	return s
}

func (s *DeliveryItemDetailSynRequest) SetDeliveryItemDetailDTOS(v []*DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) *DeliveryItemDetailSynRequest {
	s.DeliveryItemDetailDTOS = v
	return s
}

func (s *DeliveryItemDetailSynRequest) SetDeliveryItemId(v string) *DeliveryItemDetailSynRequest {
	s.DeliveryItemId = &v
	return s
}

func (s *DeliveryItemDetailSynRequest) SetDeliveryPlanId(v string) *DeliveryItemDetailSynRequest {
	s.DeliveryPlanId = &v
	return s
}

type DeliveryItemDetailSynRequestDeliveryItemDetailDTOS struct {
	ActualSupplyTime             *string `json:"actualSupplyTime,omitempty" xml:"actualSupplyTime,omitempty"`
	Amount                       *int64  `json:"amount,omitempty" xml:"amount,omitempty"`
	Comment                      *string `json:"comment,omitempty" xml:"comment,omitempty"`
	DeliveredAmount              *int64  `json:"deliveredAmount,omitempty" xml:"deliveredAmount,omitempty"`
	DeliveryItemId               *string `json:"deliveryItemId,omitempty" xml:"deliveryItemId,omitempty"`
	DeliveryPlanId               *string `json:"deliveryPlanId,omitempty" xml:"deliveryPlanId,omitempty"`
	LastSupplyTime               *string `json:"lastSupplyTime,omitempty" xml:"lastSupplyTime,omitempty"`
	Status                       *string `json:"status,omitempty" xml:"status,omitempty"`
	SubDemandSupplyPerformerName *string `json:"subDemandSupplyPerformerName,omitempty" xml:"subDemandSupplyPerformerName,omitempty"`
	SubDemandSupplyPerformerUid  *string `json:"subDemandSupplyPerformerUid,omitempty" xml:"subDemandSupplyPerformerUid,omitempty"`
	SubDemandSupplyPmName        *string `json:"subDemandSupplyPmName,omitempty" xml:"subDemandSupplyPmName,omitempty"`
	SubDemandSupplyPmUid         *string `json:"subDemandSupplyPmUid,omitempty" xml:"subDemandSupplyPmUid,omitempty"`
	SubOrderId                   *int64  `json:"subOrderId,omitempty" xml:"subOrderId,omitempty"`
	SupplyStatus                 *string `json:"supplyStatus,omitempty" xml:"supplyStatus,omitempty"`
	TotalOrderId                 *int64  `json:"totalOrderId,omitempty" xml:"totalOrderId,omitempty"`
}

func (s DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) String() string {
	return tea.Prettify(s)
}

func (s DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) GoString() string {
	return s.String()
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetActualSupplyTime(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.ActualSupplyTime = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetAmount(v int64) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.Amount = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetComment(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.Comment = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetDeliveredAmount(v int64) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.DeliveredAmount = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetDeliveryItemId(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.DeliveryItemId = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetDeliveryPlanId(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.DeliveryPlanId = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetLastSupplyTime(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.LastSupplyTime = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetStatus(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.Status = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetSubDemandSupplyPerformerName(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.SubDemandSupplyPerformerName = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetSubDemandSupplyPerformerUid(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.SubDemandSupplyPerformerUid = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetSubDemandSupplyPmName(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.SubDemandSupplyPmName = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetSubDemandSupplyPmUid(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.SubDemandSupplyPmUid = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetSubOrderId(v int64) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.SubOrderId = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetSupplyStatus(v string) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.SupplyStatus = &v
	return s
}

func (s *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS) SetTotalOrderId(v int64) *DeliveryItemDetailSynRequestDeliveryItemDetailDTOS {
	s.TotalOrderId = &v
	return s
}

type DeliveryItemDetailSynResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 212cf01016405759151137225e83cd
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DeliveryItemDetailSynResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeliveryItemDetailSynResponseBody) GoString() string {
	return s.String()
}

func (s *DeliveryItemDetailSynResponseBody) SetCode(v int64) *DeliveryItemDetailSynResponseBody {
	s.Code = &v
	return s
}

func (s *DeliveryItemDetailSynResponseBody) SetData(v bool) *DeliveryItemDetailSynResponseBody {
	s.Data = &v
	return s
}

func (s *DeliveryItemDetailSynResponseBody) SetMessage(v string) *DeliveryItemDetailSynResponseBody {
	s.Message = &v
	return s
}

func (s *DeliveryItemDetailSynResponseBody) SetSuccess(v bool) *DeliveryItemDetailSynResponseBody {
	s.Success = &v
	return s
}

func (s *DeliveryItemDetailSynResponseBody) SetTraceId(v string) *DeliveryItemDetailSynResponseBody {
	s.TraceId = &v
	return s
}

type DeliveryItemDetailSynResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeliveryItemDetailSynResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeliveryItemDetailSynResponse) String() string {
	return tea.Prettify(s)
}

func (s DeliveryItemDetailSynResponse) GoString() string {
	return s.String()
}

func (s *DeliveryItemDetailSynResponse) SetHeaders(v map[string]*string) *DeliveryItemDetailSynResponse {
	s.Headers = v
	return s
}

func (s *DeliveryItemDetailSynResponse) SetStatusCode(v int32) *DeliveryItemDetailSynResponse {
	s.StatusCode = &v
	return s
}

func (s *DeliveryItemDetailSynResponse) SetBody(v *DeliveryItemDetailSynResponseBody) *DeliveryItemDetailSynResponse {
	s.Body = v
	return s
}

type GetUrgentDemandItemListHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s GetUrgentDemandItemListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandItemListHeaders) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListHeaders) SetCommonHeaders(v map[string]*string) *GetUrgentDemandItemListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUrgentDemandItemListHeaders) SetYunUserId(v string) *GetUrgentDemandItemListHeaders {
	s.YunUserId = &v
	return s
}

type GetUrgentDemandItemListRequest struct {
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// example:
	//
	// ecs/yundisk
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	// example:
	//
	// 1
	Current *int64  `json:"current,omitempty" xml:"current,omitempty"`
	PlanId  *int64  `json:"planId,omitempty" xml:"planId,omitempty"`
	Region  *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 10
	Size *int64  `json:"size,omitempty" xml:"size,omitempty"`
	Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
}

func (s GetUrgentDemandItemListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandItemListRequest) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListRequest) SetCommodityCode(v string) *GetUrgentDemandItemListRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) SetCommodityTypeCode(v string) *GetUrgentDemandItemListRequest {
	s.CommodityTypeCode = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) SetCurrent(v int64) *GetUrgentDemandItemListRequest {
	s.Current = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) SetPlanId(v int64) *GetUrgentDemandItemListRequest {
	s.PlanId = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) SetRegion(v string) *GetUrgentDemandItemListRequest {
	s.Region = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) SetSize(v int64) *GetUrgentDemandItemListRequest {
	s.Size = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) SetZone(v string) *GetUrgentDemandItemListRequest {
	s.Zone = &v
	return s
}

type GetUrgentDemandItemListResponseBody struct {
	Code    *int64                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetUrgentDemandItemListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Success *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
	TraceId *string                                  `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetUrgentDemandItemListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandItemListResponseBody) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListResponseBody) SetCode(v int64) *GetUrgentDemandItemListResponseBody {
	s.Code = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBody) SetData(v *GetUrgentDemandItemListResponseBodyData) *GetUrgentDemandItemListResponseBody {
	s.Data = v
	return s
}

func (s *GetUrgentDemandItemListResponseBody) SetMessage(v string) *GetUrgentDemandItemListResponseBody {
	s.Message = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBody) SetSuccess(v bool) *GetUrgentDemandItemListResponseBody {
	s.Success = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBody) SetTraceId(v string) *GetUrgentDemandItemListResponseBody {
	s.TraceId = &v
	return s
}

type GetUrgentDemandItemListResponseBodyData struct {
	Current *int64                                            `json:"current,omitempty" xml:"current,omitempty"`
	Pages   *int64                                            `json:"pages,omitempty" xml:"pages,omitempty"`
	Records []*GetUrgentDemandItemListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	Size    *int64                                            `json:"size,omitempty" xml:"size,omitempty"`
	Total   *int64                                            `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetUrgentDemandItemListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandItemListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListResponseBodyData) SetCurrent(v int64) *GetUrgentDemandItemListResponseBodyData {
	s.Current = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyData) SetPages(v int64) *GetUrgentDemandItemListResponseBodyData {
	s.Pages = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyData) SetRecords(v []*GetUrgentDemandItemListResponseBodyDataRecords) *GetUrgentDemandItemListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyData) SetSize(v int64) *GetUrgentDemandItemListResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyData) SetTotal(v int64) *GetUrgentDemandItemListResponseBodyData {
	s.Total = &v
	return s
}

type GetUrgentDemandItemListResponseBodyDataRecords struct {
	Zone                   *string                                                               `json:"Zone,omitempty" xml:"Zone,omitempty"`
	AccountId              *string                                                               `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CommodityCode          *string                                                               `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	CommodityNum           *int64                                                                `json:"commodityNum,omitempty" xml:"commodityNum,omitempty"`
	CommodityTypeCode      *string                                                               `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	Creator                *string                                                               `json:"creator,omitempty" xml:"creator,omitempty"`
	CreatorName            *string                                                               `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	EffectTime             *string                                                               `json:"effectTime,omitempty" xml:"effectTime,omitempty"`
	GmtModified            *string                                                               `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                     *int64                                                                `json:"id,omitempty" xml:"id,omitempty"`
	Modifier               *string                                                               `json:"modifier,omitempty" xml:"modifier,omitempty"`
	ModifierName           *string                                                               `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	NetworkType            *string                                                               `json:"networkType,omitempty" xml:"networkType,omitempty"`
	PayDuration            *int64                                                                `json:"payDuration,omitempty" xml:"payDuration,omitempty"`
	PayDurationUnit        *string                                                               `json:"payDurationUnit,omitempty" xml:"payDurationUnit,omitempty"`
	PayType                *string                                                               `json:"payType,omitempty" xml:"payType,omitempty"`
	PlanId                 *int64                                                                `json:"planId,omitempty" xml:"planId,omitempty"`
	Region                 *string                                                               `json:"region,omitempty" xml:"region,omitempty"`
	UrgentDemandEbsRequest *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest `json:"urgentDemandEbsRequest,omitempty" xml:"urgentDemandEbsRequest,omitempty" type:"Struct"`
	UrgentDemandEcsRequest *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest `json:"urgentDemandEcsRequest,omitempty" xml:"urgentDemandEcsRequest,omitempty" type:"Struct"`
}

func (s GetUrgentDemandItemListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandItemListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetZone(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.Zone = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetAccountId(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.AccountId = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetCommodityCode(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.CommodityCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetCommodityNum(v int64) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.CommodityNum = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetCommodityTypeCode(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.CommodityTypeCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetCreator(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetCreatorName(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.CreatorName = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetEffectTime(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.EffectTime = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetGmtModified(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetId(v int64) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.Id = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetModifier(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.Modifier = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetModifierName(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.ModifierName = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetNetworkType(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.NetworkType = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetPayDuration(v int64) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.PayDuration = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetPayDurationUnit(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.PayDurationUnit = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetPayType(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.PayType = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetPlanId(v int64) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.PlanId = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetRegion(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.Region = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetUrgentDemandEbsRequest(v *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.UrgentDemandEbsRequest = v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetUrgentDemandEcsRequest(v *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.UrgentDemandEcsRequest = v
	return s
}

type GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest struct {
	// example:
	//
	// cloud_essd
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// example:
	//
	// 1
	CommodityNum *int64 `json:"commodityNum,omitempty" xml:"commodityNum,omitempty"`
	// example:
	//
	// yundisk
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	// example:
	//
	// 1
	DataDiskSize *int64 `json:"dataDiskSize,omitempty" xml:"dataDiskSize,omitempty"`
	// example:
	//
	// 111222
	ItemId *int64 `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// example:
	//
	// 1
	PerformanceLevel *int64 `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
}

func (s GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) SetCommodityCode(v string) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) SetCommodityNum(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	s.CommodityNum = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) SetCommodityTypeCode(v string) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	s.CommodityTypeCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) SetDataDiskSize(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	s.DataDiskSize = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) SetItemId(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	s.ItemId = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) SetPerformanceLevel(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	s.PerformanceLevel = &v
	return s
}

type GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest struct {
	CommodityCode     *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	CommodityNum      *int64  `json:"commodityNum,omitempty" xml:"commodityNum,omitempty"`
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	ItemId            *int64  `json:"itemId,omitempty" xml:"itemId,omitempty"`
	VcpuCount         *int64  `json:"vcpuCount,omitempty" xml:"vcpuCount,omitempty"`
}

func (s GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) SetCommodityCode(v string) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) SetCommodityNum(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest {
	s.CommodityNum = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) SetCommodityTypeCode(v string) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest {
	s.CommodityTypeCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) SetItemId(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest {
	s.ItemId = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) SetVcpuCount(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest {
	s.VcpuCount = &v
	return s
}

type GetUrgentDemandItemListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUrgentDemandItemListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUrgentDemandItemListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandItemListResponse) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListResponse) SetHeaders(v map[string]*string) *GetUrgentDemandItemListResponse {
	s.Headers = v
	return s
}

func (s *GetUrgentDemandItemListResponse) SetStatusCode(v int32) *GetUrgentDemandItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponse) SetBody(v *GetUrgentDemandItemListResponseBody) *GetUrgentDemandItemListResponse {
	s.Body = v
	return s
}

type GetUrgentDemandPlanDetailHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 262940
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s GetUrgentDemandPlanDetailHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandPlanDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanDetailHeaders) SetCommonHeaders(v map[string]*string) *GetUrgentDemandPlanDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUrgentDemandPlanDetailHeaders) SetYunUserId(v string) *GetUrgentDemandPlanDetailHeaders {
	s.YunUserId = &v
	return s
}

type GetUrgentDemandPlanDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111223
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
}

func (s GetUrgentDemandPlanDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandPlanDetailRequest) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanDetailRequest) SetPlanId(v string) *GetUrgentDemandPlanDetailRequest {
	s.PlanId = &v
	return s
}

type GetUrgentDemandPlanDetailResponseBody struct {
	// code
	//
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// body
	Data *GetUrgentDemandPlanDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 1e2b798516402440016572132e1459
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetUrgentDemandPlanDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandPlanDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanDetailResponseBody) SetCode(v int64) *GetUrgentDemandPlanDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBody) SetData(v *GetUrgentDemandPlanDetailResponseBodyData) *GetUrgentDemandPlanDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBody) SetMessage(v string) *GetUrgentDemandPlanDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBody) SetSuccess(v bool) *GetUrgentDemandPlanDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBody) SetTraceId(v string) *GetUrgentDemandPlanDetailResponseBody {
	s.TraceId = &v
	return s
}

type GetUrgentDemandPlanDetailResponseBodyData struct {
	AccountDept *string `json:"accountDept,omitempty" xml:"accountDept,omitempty"`
	// example:
	//
	// 1065737167271819
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// larus_prd
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// example:
	//
	// https://xxxxx
	ApprovalUrl *string `json:"approvalUrl,omitempty" xml:"approvalUrl,omitempty"`
	// example:
	//
	// {}
	BpmSubstate           map[string]interface{} `json:"bpmSubstate,omitempty" xml:"bpmSubstate,omitempty"`
	CommodityTypeCodeList []*string              `json:"commodityTypeCodeList,omitempty" xml:"commodityTypeCodeList,omitempty" type:"Repeated"`
	// example:
	//
	// 262940
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// xxx
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// xxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// URGENT
	DetailType *string `json:"detailType,omitempty" xml:"detailType,omitempty"`
	// example:
	//
	// 2021-12-17 16:53:21
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2021-12-17 16:53:21
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 262940
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// xx
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// example:
	//
	// 11223
	PlanId   *int64  `json:"planId,omitempty" xml:"planId,omitempty"`
	PlanName *string `json:"planName,omitempty" xml:"planName,omitempty"`
	// example:
	//
	// 220
	Status            *int64  `json:"status,omitempty" xml:"status,omitempty"`
	YunzhiProductName *string `json:"yunzhiProductName,omitempty" xml:"yunzhiProductName,omitempty"`
}

func (s GetUrgentDemandPlanDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandPlanDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetAccountDept(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.AccountDept = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetAccountId(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetAccountName(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetApprovalUrl(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.ApprovalUrl = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetBpmSubstate(v map[string]interface{}) *GetUrgentDemandPlanDetailResponseBodyData {
	s.BpmSubstate = v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetCommodityTypeCodeList(v []*string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.CommodityTypeCodeList = v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetCreator(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetCreatorName(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetDescription(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetDetailType(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.DetailType = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetGmtCreate(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetGmtModified(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetModifier(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetModifierName(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.ModifierName = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetPlanId(v int64) *GetUrgentDemandPlanDetailResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetPlanName(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.PlanName = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetStatus(v int64) *GetUrgentDemandPlanDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponseBodyData) SetYunzhiProductName(v string) *GetUrgentDemandPlanDetailResponseBodyData {
	s.YunzhiProductName = &v
	return s
}

type GetUrgentDemandPlanDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUrgentDemandPlanDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUrgentDemandPlanDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandPlanDetailResponse) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanDetailResponse) SetHeaders(v map[string]*string) *GetUrgentDemandPlanDetailResponse {
	s.Headers = v
	return s
}

func (s *GetUrgentDemandPlanDetailResponse) SetStatusCode(v int32) *GetUrgentDemandPlanDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUrgentDemandPlanDetailResponse) SetBody(v *GetUrgentDemandPlanDetailResponseBody) *GetUrgentDemandPlanDetailResponse {
	s.Body = v
	return s
}

type GetUrgentDemandPlanListHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s GetUrgentDemandPlanListHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandPlanListHeaders) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanListHeaders) SetCommonHeaders(v map[string]*string) *GetUrgentDemandPlanListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUrgentDemandPlanListHeaders) SetYunUserId(v string) *GetUrgentDemandPlanListHeaders {
	s.YunUserId = &v
	return s
}

type GetUrgentDemandPlanListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FY2022
	Period *string `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	PlanType *int64 `json:"planType,omitempty" xml:"planType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetUrgentDemandPlanListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandPlanListRequest) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanListRequest) SetCurrent(v int64) *GetUrgentDemandPlanListRequest {
	s.Current = &v
	return s
}

func (s *GetUrgentDemandPlanListRequest) SetPeriod(v string) *GetUrgentDemandPlanListRequest {
	s.Period = &v
	return s
}

func (s *GetUrgentDemandPlanListRequest) SetPlanType(v int64) *GetUrgentDemandPlanListRequest {
	s.PlanType = &v
	return s
}

func (s *GetUrgentDemandPlanListRequest) SetSize(v int64) *GetUrgentDemandPlanListRequest {
	s.Size = &v
	return s
}

func (s *GetUrgentDemandPlanListRequest) SetUserId(v string) *GetUrgentDemandPlanListRequest {
	s.UserId = &v
	return s
}

type GetUrgentDemandPlanListResponseBody struct {
	// example:
	//
	// 0
	Code *int64                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetUrgentDemandPlanListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2127968716405850615204514e9332
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetUrgentDemandPlanListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandPlanListResponseBody) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanListResponseBody) SetCode(v int64) *GetUrgentDemandPlanListResponseBody {
	s.Code = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBody) SetData(v *GetUrgentDemandPlanListResponseBodyData) *GetUrgentDemandPlanListResponseBody {
	s.Data = v
	return s
}

func (s *GetUrgentDemandPlanListResponseBody) SetMessage(v string) *GetUrgentDemandPlanListResponseBody {
	s.Message = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBody) SetSuccess(v bool) *GetUrgentDemandPlanListResponseBody {
	s.Success = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBody) SetTraceId(v string) *GetUrgentDemandPlanListResponseBody {
	s.TraceId = &v
	return s
}

type GetUrgentDemandPlanListResponseBodyData struct {
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// 2
	Pages   *int64                                            `json:"pages,omitempty" xml:"pages,omitempty"`
	Records []*GetUrgentDemandPlanListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 15
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetUrgentDemandPlanListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandPlanListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanListResponseBodyData) SetCurrent(v int64) *GetUrgentDemandPlanListResponseBodyData {
	s.Current = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyData) SetPages(v int64) *GetUrgentDemandPlanListResponseBodyData {
	s.Pages = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyData) SetRecords(v []*GetUrgentDemandPlanListResponseBodyDataRecords) *GetUrgentDemandPlanListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyData) SetSize(v int64) *GetUrgentDemandPlanListResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyData) SetTotal(v int64) *GetUrgentDemandPlanListResponseBodyData {
	s.Total = &v
	return s
}

type GetUrgentDemandPlanListResponseBodyDataRecords struct {
	// example:
	//
	// 1705524002740212
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// xxxx
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// example:
	//
	// ALIYUN
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// example:
	//
	// https://xxx
	ApprovalUrl *string `json:"approvalUrl,omitempty" xml:"approvalUrl,omitempty"`
	// example:
	//
	// 1111
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// xxxx
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// xxxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2021-12-20 10:29:50
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2021-12-20 10:29:50
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// xxxx
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// xxxx
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	PlanId       *int64  `json:"planId,omitempty" xml:"planId,omitempty"`
	PlanName     *string `json:"planName,omitempty" xml:"planName,omitempty"`
	Status       *int64  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetUrgentDemandPlanListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandPlanListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetAccountId(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.AccountId = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetAccountName(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.AccountName = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetAccountType(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.AccountType = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetApprovalUrl(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.ApprovalUrl = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetCreator(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetCreatorName(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.CreatorName = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetDescription(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetGmtCreate(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetGmtModified(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetModifier(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.Modifier = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetModifierName(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.ModifierName = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetPlanId(v int64) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.PlanId = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetPlanName(v string) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.PlanName = &v
	return s
}

func (s *GetUrgentDemandPlanListResponseBodyDataRecords) SetStatus(v int64) *GetUrgentDemandPlanListResponseBodyDataRecords {
	s.Status = &v
	return s
}

type GetUrgentDemandPlanListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUrgentDemandPlanListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUrgentDemandPlanListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUrgentDemandPlanListResponse) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandPlanListResponse) SetHeaders(v map[string]*string) *GetUrgentDemandPlanListResponse {
	s.Headers = v
	return s
}

func (s *GetUrgentDemandPlanListResponse) SetStatusCode(v int32) *GetUrgentDemandPlanListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUrgentDemandPlanListResponse) SetBody(v *GetUrgentDemandPlanListResponseBody) *GetUrgentDemandPlanListResponse {
	s.Body = v
	return s
}

type PushResourcePlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s PushResourcePlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s PushResourcePlanHeaders) GoString() string {
	return s.String()
}

func (s *PushResourcePlanHeaders) SetCommonHeaders(v map[string]*string) *PushResourcePlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushResourcePlanHeaders) SetYunUserId(v string) *PushResourcePlanHeaders {
	s.YunUserId = &v
	return s
}

type PushResourcePlanRequest struct {
	BufferCnt   *int64                               `json:"bufferCnt,omitempty" xml:"bufferCnt,omitempty"`
	DemandCount *int64                               `json:"demandCount,omitempty" xml:"demandCount,omitempty"`
	DemandId    *string                              `json:"demandId,omitempty" xml:"demandId,omitempty"`
	MethodList  []*PushResourcePlanRequestMethodList `json:"methodList,omitempty" xml:"methodList,omitempty" type:"Repeated"`
	RequestId   *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	RequireCnt  *int64                               `json:"requireCnt,omitempty" xml:"requireCnt,omitempty"`
	SubDemandId *string                              `json:"subDemandId,omitempty" xml:"subDemandId,omitempty"`
}

func (s PushResourcePlanRequest) String() string {
	return tea.Prettify(s)
}

func (s PushResourcePlanRequest) GoString() string {
	return s.String()
}

func (s *PushResourcePlanRequest) SetBufferCnt(v int64) *PushResourcePlanRequest {
	s.BufferCnt = &v
	return s
}

func (s *PushResourcePlanRequest) SetDemandCount(v int64) *PushResourcePlanRequest {
	s.DemandCount = &v
	return s
}

func (s *PushResourcePlanRequest) SetDemandId(v string) *PushResourcePlanRequest {
	s.DemandId = &v
	return s
}

func (s *PushResourcePlanRequest) SetMethodList(v []*PushResourcePlanRequestMethodList) *PushResourcePlanRequest {
	s.MethodList = v
	return s
}

func (s *PushResourcePlanRequest) SetRequestId(v string) *PushResourcePlanRequest {
	s.RequestId = &v
	return s
}

func (s *PushResourcePlanRequest) SetRequireCnt(v int64) *PushResourcePlanRequest {
	s.RequireCnt = &v
	return s
}

func (s *PushResourcePlanRequest) SetSubDemandId(v string) *PushResourcePlanRequest {
	s.SubDemandId = &v
	return s
}

type PushResourcePlanRequestMethodList struct {
	Azone            *string                                      `json:"azone,omitempty" xml:"azone,omitempty"`
	BufferCnt        *int64                                       `json:"bufferCnt,omitempty" xml:"bufferCnt,omitempty"`
	Cluster          *string                                      `json:"cluster,omitempty" xml:"cluster,omitempty"`
	Comment          *string                                      `json:"comment,omitempty" xml:"comment,omitempty"`
	ConvertHostCnt   *int64                                       `json:"convertHostCnt,omitempty" xml:"convertHostCnt,omitempty"`
	ConvertHostType  *string                                      `json:"convertHostType,omitempty" xml:"convertHostType,omitempty"`
	DataList         []*PushResourcePlanRequestMethodListDataList `json:"dataList,omitempty" xml:"dataList,omitempty" type:"Repeated"`
	DenamdCount      *int64                                       `json:"denamdCount,omitempty" xml:"denamdCount,omitempty"`
	GapCnt           *int64                                       `json:"gapCnt,omitempty" xml:"gapCnt,omitempty"`
	PromiseDate      *string                                      `json:"promiseDate,omitempty" xml:"promiseDate,omitempty"`
	Region           *string                                      `json:"region,omitempty" xml:"region,omitempty"`
	ResourceMethodId *int64                                       `json:"resourceMethodId,omitempty" xml:"resourceMethodId,omitempty"`
	RoomCode         *string                                      `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
}

func (s PushResourcePlanRequestMethodList) String() string {
	return tea.Prettify(s)
}

func (s PushResourcePlanRequestMethodList) GoString() string {
	return s.String()
}

func (s *PushResourcePlanRequestMethodList) SetAzone(v string) *PushResourcePlanRequestMethodList {
	s.Azone = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetBufferCnt(v int64) *PushResourcePlanRequestMethodList {
	s.BufferCnt = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetCluster(v string) *PushResourcePlanRequestMethodList {
	s.Cluster = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetComment(v string) *PushResourcePlanRequestMethodList {
	s.Comment = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetConvertHostCnt(v int64) *PushResourcePlanRequestMethodList {
	s.ConvertHostCnt = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetConvertHostType(v string) *PushResourcePlanRequestMethodList {
	s.ConvertHostType = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetDataList(v []*PushResourcePlanRequestMethodListDataList) *PushResourcePlanRequestMethodList {
	s.DataList = v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetDenamdCount(v int64) *PushResourcePlanRequestMethodList {
	s.DenamdCount = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetGapCnt(v int64) *PushResourcePlanRequestMethodList {
	s.GapCnt = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetPromiseDate(v string) *PushResourcePlanRequestMethodList {
	s.PromiseDate = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetRegion(v string) *PushResourcePlanRequestMethodList {
	s.Region = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetResourceMethodId(v int64) *PushResourcePlanRequestMethodList {
	s.ResourceMethodId = &v
	return s
}

func (s *PushResourcePlanRequestMethodList) SetRoomCode(v string) *PushResourcePlanRequestMethodList {
	s.RoomCode = &v
	return s
}

type PushResourcePlanRequestMethodListDataList struct {
	ClassZone       *string `json:"classZone,omitempty" xml:"classZone,omitempty"`
	ConvertHostType *string `json:"convertHostType,omitempty" xml:"convertHostType,omitempty"`
	LogicZone       *string `json:"logicZone,omitempty" xml:"logicZone,omitempty"`
	NetArch         *string `json:"netArch,omitempty" xml:"netArch,omitempty"`
	Nic             *string `json:"nic,omitempty" xml:"nic,omitempty"`
	Product3        *string `json:"product3,omitempty" xml:"product3,omitempty"`
	SafeZone        *string `json:"safeZone,omitempty" xml:"safeZone,omitempty"`
	Scenario        *string `json:"scenario,omitempty" xml:"scenario,omitempty"`
	SupplyAmount    *int64  `json:"supplyAmount,omitempty" xml:"supplyAmount,omitempty"`
	// This parameter is required.
	SupplyDate     *string `json:"supplyDate,omitempty" xml:"supplyDate,omitempty"`
	SupplyType     *int64  `json:"supplyType,omitempty" xml:"supplyType,omitempty"`
	SupplyVmAmount *int32  `json:"supplyVmAmount,omitempty" xml:"supplyVmAmount,omitempty"`
}

func (s PushResourcePlanRequestMethodListDataList) String() string {
	return tea.Prettify(s)
}

func (s PushResourcePlanRequestMethodListDataList) GoString() string {
	return s.String()
}

func (s *PushResourcePlanRequestMethodListDataList) SetClassZone(v string) *PushResourcePlanRequestMethodListDataList {
	s.ClassZone = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetConvertHostType(v string) *PushResourcePlanRequestMethodListDataList {
	s.ConvertHostType = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetLogicZone(v string) *PushResourcePlanRequestMethodListDataList {
	s.LogicZone = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetNetArch(v string) *PushResourcePlanRequestMethodListDataList {
	s.NetArch = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetNic(v string) *PushResourcePlanRequestMethodListDataList {
	s.Nic = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetProduct3(v string) *PushResourcePlanRequestMethodListDataList {
	s.Product3 = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetSafeZone(v string) *PushResourcePlanRequestMethodListDataList {
	s.SafeZone = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetScenario(v string) *PushResourcePlanRequestMethodListDataList {
	s.Scenario = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetSupplyAmount(v int64) *PushResourcePlanRequestMethodListDataList {
	s.SupplyAmount = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetSupplyDate(v string) *PushResourcePlanRequestMethodListDataList {
	s.SupplyDate = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetSupplyType(v int64) *PushResourcePlanRequestMethodListDataList {
	s.SupplyType = &v
	return s
}

func (s *PushResourcePlanRequestMethodListDataList) SetSupplyVmAmount(v int32) *PushResourcePlanRequestMethodListDataList {
	s.SupplyVmAmount = &v
	return s
}

type PushResourcePlanResponseBody struct {
	Code    *int64  `json:"code,omitempty" xml:"code,omitempty"`
	Data    *bool   `json:"data,omitempty" xml:"data,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s PushResourcePlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushResourcePlanResponseBody) GoString() string {
	return s.String()
}

func (s *PushResourcePlanResponseBody) SetCode(v int64) *PushResourcePlanResponseBody {
	s.Code = &v
	return s
}

func (s *PushResourcePlanResponseBody) SetData(v bool) *PushResourcePlanResponseBody {
	s.Data = &v
	return s
}

func (s *PushResourcePlanResponseBody) SetMessage(v string) *PushResourcePlanResponseBody {
	s.Message = &v
	return s
}

func (s *PushResourcePlanResponseBody) SetSuccess(v bool) *PushResourcePlanResponseBody {
	s.Success = &v
	return s
}

func (s *PushResourcePlanResponseBody) SetTraceId(v string) *PushResourcePlanResponseBody {
	s.TraceId = &v
	return s
}

type PushResourcePlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushResourcePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushResourcePlanResponse) String() string {
	return tea.Prettify(s)
}

func (s PushResourcePlanResponse) GoString() string {
	return s.String()
}

func (s *PushResourcePlanResponse) SetHeaders(v map[string]*string) *PushResourcePlanResponse {
	s.Headers = v
	return s
}

func (s *PushResourcePlanResponse) SetStatusCode(v int32) *PushResourcePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *PushResourcePlanResponse) SetBody(v *PushResourcePlanResponseBody) *PushResourcePlanResponse {
	s.Body = v
	return s
}

type QueryDeliveredSupplyItemsRequest struct {
	AccountId         *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
}

func (s QueryDeliveredSupplyItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDeliveredSupplyItemsRequest) GoString() string {
	return s.String()
}

func (s *QueryDeliveredSupplyItemsRequest) SetAccountId(v string) *QueryDeliveredSupplyItemsRequest {
	s.AccountId = &v
	return s
}

func (s *QueryDeliveredSupplyItemsRequest) SetCommodityTypeCode(v string) *QueryDeliveredSupplyItemsRequest {
	s.CommodityTypeCode = &v
	return s
}

type QueryDeliveredSupplyItemsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*QueryDeliveredSupplyItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s QueryDeliveredSupplyItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDeliveredSupplyItemsResponse) GoString() string {
	return s.String()
}

func (s *QueryDeliveredSupplyItemsResponse) SetHeaders(v map[string]*string) *QueryDeliveredSupplyItemsResponse {
	s.Headers = v
	return s
}

func (s *QueryDeliveredSupplyItemsResponse) SetStatusCode(v int32) *QueryDeliveredSupplyItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponse) SetBody(v []*QueryDeliveredSupplyItemsResponseBody) *QueryDeliveredSupplyItemsResponse {
	s.Body = v
	return s
}

type QueryDeliveredSupplyItemsResponseBody struct {
	AccountId         *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	DemandPlanId      *int64  `json:"demandPlanId,omitempty" xml:"demandPlanId,omitempty"`
	PlanTitle         *string `json:"planTitle,omitempty" xml:"planTitle,omitempty"`
	Region            *string `json:"region,omitempty" xml:"region,omitempty"`
	Zone              *string `json:"zone,omitempty" xml:"zone,omitempty"`
	CommodityCode     *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	DemandCount       *int32  `json:"demandCount,omitempty" xml:"demandCount,omitempty"`
	DeliveredAmount   *int32  `json:"deliveredAmount,omitempty" xml:"deliveredAmount,omitempty"`
	OpenCount         *int32  `json:"openCount,omitempty" xml:"openCount,omitempty"`
}

func (s QueryDeliveredSupplyItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDeliveredSupplyItemsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetAccountId(v string) *QueryDeliveredSupplyItemsResponseBody {
	s.AccountId = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetCommodityTypeCode(v string) *QueryDeliveredSupplyItemsResponseBody {
	s.CommodityTypeCode = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetDemandPlanId(v int64) *QueryDeliveredSupplyItemsResponseBody {
	s.DemandPlanId = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetPlanTitle(v string) *QueryDeliveredSupplyItemsResponseBody {
	s.PlanTitle = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetRegion(v string) *QueryDeliveredSupplyItemsResponseBody {
	s.Region = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetZone(v string) *QueryDeliveredSupplyItemsResponseBody {
	s.Zone = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetCommodityCode(v string) *QueryDeliveredSupplyItemsResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetDemandCount(v int32) *QueryDeliveredSupplyItemsResponseBody {
	s.DemandCount = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetDeliveredAmount(v int32) *QueryDeliveredSupplyItemsResponseBody {
	s.DeliveredAmount = &v
	return s
}

func (s *QueryDeliveredSupplyItemsResponseBody) SetOpenCount(v int32) *QueryDeliveredSupplyItemsResponseBody {
	s.OpenCount = &v
	return s
}

type QueryPeriodBudgetBillRequest struct {
	ObjectIds  *string `json:"objectIds,omitempty" xml:"objectIds,omitempty"`
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	Period     *string `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryPeriodBudgetBillRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPeriodBudgetBillRequest) GoString() string {
	return s.String()
}

func (s *QueryPeriodBudgetBillRequest) SetObjectIds(v string) *QueryPeriodBudgetBillRequest {
	s.ObjectIds = &v
	return s
}

func (s *QueryPeriodBudgetBillRequest) SetObjectType(v string) *QueryPeriodBudgetBillRequest {
	s.ObjectType = &v
	return s
}

func (s *QueryPeriodBudgetBillRequest) SetPeriod(v string) *QueryPeriodBudgetBillRequest {
	s.Period = &v
	return s
}

type QueryPeriodBudgetBillResponseBody struct {
	PeriodBudgetBillDTOS []*QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS `json:"periodBudgetBillDTOS,omitempty" xml:"periodBudgetBillDTOS,omitempty" type:"Repeated"`
}

func (s QueryPeriodBudgetBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPeriodBudgetBillResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPeriodBudgetBillResponseBody) SetPeriodBudgetBillDTOS(v []*QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) *QueryPeriodBudgetBillResponseBody {
	s.PeriodBudgetBillDTOS = v
	return s
}

type QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS struct {
	Bill         *float64 `json:"bill,omitempty" xml:"bill,omitempty"`
	Budget       *float64 `json:"budget,omitempty" xml:"budget,omitempty"`
	LastYearBill *float64 `json:"lastYearBill,omitempty" xml:"lastYearBill,omitempty"`
	Month        *string  `json:"month,omitempty" xml:"month,omitempty"`
}

func (s QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) String() string {
	return tea.Prettify(s)
}

func (s QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) GoString() string {
	return s.String()
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) SetBill(v float64) *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS {
	s.Bill = &v
	return s
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) SetBudget(v float64) *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS {
	s.Budget = &v
	return s
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) SetLastYearBill(v float64) *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS {
	s.LastYearBill = &v
	return s
}

func (s *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS) SetMonth(v string) *QueryPeriodBudgetBillResponseBodyPeriodBudgetBillDTOS {
	s.Month = &v
	return s
}

type QueryPeriodBudgetBillResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPeriodBudgetBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPeriodBudgetBillResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPeriodBudgetBillResponse) GoString() string {
	return s.String()
}

func (s *QueryPeriodBudgetBillResponse) SetHeaders(v map[string]*string) *QueryPeriodBudgetBillResponse {
	s.Headers = v
	return s
}

func (s *QueryPeriodBudgetBillResponse) SetStatusCode(v int32) *QueryPeriodBudgetBillResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPeriodBudgetBillResponse) SetBody(v *QueryPeriodBudgetBillResponseBody) *QueryPeriodBudgetBillResponse {
	s.Body = v
	return s
}

type SaveUrgentDemandItemHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111222
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s SaveUrgentDemandItemHeaders) String() string {
	return tea.Prettify(s)
}

func (s SaveUrgentDemandItemHeaders) GoString() string {
	return s.String()
}

func (s *SaveUrgentDemandItemHeaders) SetCommonHeaders(v map[string]*string) *SaveUrgentDemandItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveUrgentDemandItemHeaders) SetYunUserId(v string) *SaveUrgentDemandItemHeaders {
	s.YunUserId = &v
	return s
}

type SaveUrgentDemandItemRequest struct {
	// example:
	//
	// 12321312
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 111222
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// 2021-12-27 00:00:00
	EffectTime *string `json:"effectTime,omitempty" xml:"effectTime,omitempty"`
	// example:
	//
	// 111222
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// 网络类型 vpc（私有网络）/classic（经典网络）
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// example:
	//
	// 10
	PayDuration *string `json:"payDuration,omitempty" xml:"payDuration,omitempty"`
	// example:
	//
	// 购买时长单位(month(月)，week(周)，day(天))
	PayDurationUnit *string `json:"payDurationUnit,omitempty" xml:"payDurationUnit,omitempty"`
	// example:
	//
	// 付费类型 prepay(预付费)/postpay（后付费）
	PayType *string `json:"payType,omitempty" xml:"payType,omitempty"`
	// example:
	//
	// 111222
	PlanId *int64 `json:"planId,omitempty" xml:"planId,omitempty"`
	// example:
	//
	// cn-beijing
	Region                 *string                                            `json:"region,omitempty" xml:"region,omitempty"`
	UrgentDemandEbsRequest *SaveUrgentDemandItemRequestUrgentDemandEbsRequest `json:"urgentDemandEbsRequest,omitempty" xml:"urgentDemandEbsRequest,omitempty" type:"Struct"`
	UrgentDemandEcsRequest *SaveUrgentDemandItemRequestUrgentDemandEcsRequest `json:"urgentDemandEcsRequest,omitempty" xml:"urgentDemandEcsRequest,omitempty" type:"Struct"`
	// example:
	//
	// cn-beijing-a
	Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
}

func (s SaveUrgentDemandItemRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveUrgentDemandItemRequest) GoString() string {
	return s.String()
}

func (s *SaveUrgentDemandItemRequest) SetAccountId(v string) *SaveUrgentDemandItemRequest {
	s.AccountId = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetCreator(v string) *SaveUrgentDemandItemRequest {
	s.Creator = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetEffectTime(v string) *SaveUrgentDemandItemRequest {
	s.EffectTime = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetModifier(v string) *SaveUrgentDemandItemRequest {
	s.Modifier = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetNetworkType(v string) *SaveUrgentDemandItemRequest {
	s.NetworkType = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetPayDuration(v string) *SaveUrgentDemandItemRequest {
	s.PayDuration = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetPayDurationUnit(v string) *SaveUrgentDemandItemRequest {
	s.PayDurationUnit = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetPayType(v string) *SaveUrgentDemandItemRequest {
	s.PayType = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetPlanId(v int64) *SaveUrgentDemandItemRequest {
	s.PlanId = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetRegion(v string) *SaveUrgentDemandItemRequest {
	s.Region = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetUrgentDemandEbsRequest(v *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) *SaveUrgentDemandItemRequest {
	s.UrgentDemandEbsRequest = v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetUrgentDemandEcsRequest(v *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) *SaveUrgentDemandItemRequest {
	s.UrgentDemandEcsRequest = v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetZone(v string) *SaveUrgentDemandItemRequest {
	s.Zone = &v
	return s
}

type SaveUrgentDemandItemRequestUrgentDemandEbsRequest struct {
	// example:
	//
	// cloud_essd
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// example:
	//
	// 1
	CommodityNum *int64 `json:"commodityNum,omitempty" xml:"commodityNum,omitempty"`
	// example:
	//
	// yundisk
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	// example:
	//
	// 111222
	ItemId *int64 `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// example:
	//
	// 1
	PerformanceLevel *int64 `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
}

func (s SaveUrgentDemandItemRequestUrgentDemandEbsRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveUrgentDemandItemRequestUrgentDemandEbsRequest) GoString() string {
	return s.String()
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) SetCommodityCode(v string) *SaveUrgentDemandItemRequestUrgentDemandEbsRequest {
	s.CommodityCode = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) SetCommodityNum(v int64) *SaveUrgentDemandItemRequestUrgentDemandEbsRequest {
	s.CommodityNum = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) SetCommodityTypeCode(v string) *SaveUrgentDemandItemRequestUrgentDemandEbsRequest {
	s.CommodityTypeCode = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) SetItemId(v int64) *SaveUrgentDemandItemRequestUrgentDemandEbsRequest {
	s.ItemId = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) SetPerformanceLevel(v int64) *SaveUrgentDemandItemRequestUrgentDemandEbsRequest {
	s.PerformanceLevel = &v
	return s
}

type SaveUrgentDemandItemRequestUrgentDemandEcsRequest struct {
	// example:
	//
	// ecs.sn2ne.6xlarge
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// example:
	//
	// 2
	CommodityNum *int64 `json:"commodityNum,omitempty" xml:"commodityNum,omitempty"`
	// example:
	//
	// ecs
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	// example:
	//
	// 111222
	ItemId *int64 `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// example:
	//
	// 2
	VCpuCount *int64 `json:"vCpuCount,omitempty" xml:"vCpuCount,omitempty"`
}

func (s SaveUrgentDemandItemRequestUrgentDemandEcsRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveUrgentDemandItemRequestUrgentDemandEcsRequest) GoString() string {
	return s.String()
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) SetCommodityCode(v string) *SaveUrgentDemandItemRequestUrgentDemandEcsRequest {
	s.CommodityCode = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) SetCommodityNum(v int64) *SaveUrgentDemandItemRequestUrgentDemandEcsRequest {
	s.CommodityNum = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) SetCommodityTypeCode(v string) *SaveUrgentDemandItemRequestUrgentDemandEcsRequest {
	s.CommodityTypeCode = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) SetItemId(v int64) *SaveUrgentDemandItemRequestUrgentDemandEcsRequest {
	s.ItemId = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) SetVCpuCount(v int64) *SaveUrgentDemandItemRequestUrgentDemandEcsRequest {
	s.VCpuCount = &v
	return s
}

type SaveUrgentDemandItemResponseBody struct {
	// example:
	//
	// 0
	Code *int64   `json:"code,omitempty" xml:"code,omitempty"`
	Data []*int64 `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2107d95616405752026995105e83b0
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s SaveUrgentDemandItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveUrgentDemandItemResponseBody) GoString() string {
	return s.String()
}

func (s *SaveUrgentDemandItemResponseBody) SetCode(v int64) *SaveUrgentDemandItemResponseBody {
	s.Code = &v
	return s
}

func (s *SaveUrgentDemandItemResponseBody) SetData(v []*int64) *SaveUrgentDemandItemResponseBody {
	s.Data = v
	return s
}

func (s *SaveUrgentDemandItemResponseBody) SetMessage(v string) *SaveUrgentDemandItemResponseBody {
	s.Message = &v
	return s
}

func (s *SaveUrgentDemandItemResponseBody) SetSuccess(v bool) *SaveUrgentDemandItemResponseBody {
	s.Success = &v
	return s
}

func (s *SaveUrgentDemandItemResponseBody) SetTraceId(v string) *SaveUrgentDemandItemResponseBody {
	s.TraceId = &v
	return s
}

type SaveUrgentDemandItemResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveUrgentDemandItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveUrgentDemandItemResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveUrgentDemandItemResponse) GoString() string {
	return s.String()
}

func (s *SaveUrgentDemandItemResponse) SetHeaders(v map[string]*string) *SaveUrgentDemandItemResponse {
	s.Headers = v
	return s
}

func (s *SaveUrgentDemandItemResponse) SetStatusCode(v int32) *SaveUrgentDemandItemResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveUrgentDemandItemResponse) SetBody(v *SaveUrgentDemandItemResponseBody) *SaveUrgentDemandItemResponse {
	s.Body = v
	return s
}

type SubmitUrgentDemandPlanHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 262940
	YunUserId *string `json:"Yun-User-Id,omitempty" xml:"Yun-User-Id,omitempty"`
}

func (s SubmitUrgentDemandPlanHeaders) String() string {
	return tea.Prettify(s)
}

func (s SubmitUrgentDemandPlanHeaders) GoString() string {
	return s.String()
}

func (s *SubmitUrgentDemandPlanHeaders) SetCommonHeaders(v map[string]*string) *SubmitUrgentDemandPlanHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubmitUrgentDemandPlanHeaders) SetYunUserId(v string) *SubmitUrgentDemandPlanHeaders {
	s.YunUserId = &v
	return s
}

type SubmitUrgentDemandPlanRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 111223
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 262940
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SubmitUrgentDemandPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitUrgentDemandPlanRequest) GoString() string {
	return s.String()
}

func (s *SubmitUrgentDemandPlanRequest) SetPlanId(v string) *SubmitUrgentDemandPlanRequest {
	s.PlanId = &v
	return s
}

func (s *SubmitUrgentDemandPlanRequest) SetUserId(v string) *SubmitUrgentDemandPlanRequest {
	s.UserId = &v
	return s
}

type SubmitUrgentDemandPlanResponseBody struct {
	// code
	//
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 1e2b798516402440016572132e1459
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s SubmitUrgentDemandPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitUrgentDemandPlanResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitUrgentDemandPlanResponseBody) SetCode(v int64) *SubmitUrgentDemandPlanResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitUrgentDemandPlanResponseBody) SetData(v bool) *SubmitUrgentDemandPlanResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitUrgentDemandPlanResponseBody) SetMessage(v string) *SubmitUrgentDemandPlanResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitUrgentDemandPlanResponseBody) SetSuccess(v bool) *SubmitUrgentDemandPlanResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitUrgentDemandPlanResponseBody) SetTraceId(v string) *SubmitUrgentDemandPlanResponseBody {
	s.TraceId = &v
	return s
}

type SubmitUrgentDemandPlanResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitUrgentDemandPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitUrgentDemandPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitUrgentDemandPlanResponse) GoString() string {
	return s.String()
}

func (s *SubmitUrgentDemandPlanResponse) SetHeaders(v map[string]*string) *SubmitUrgentDemandPlanResponse {
	s.Headers = v
	return s
}

func (s *SubmitUrgentDemandPlanResponse) SetStatusCode(v int32) *SubmitUrgentDemandPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitUrgentDemandPlanResponse) SetBody(v *SubmitUrgentDemandPlanResponseBody) *SubmitUrgentDemandPlanResponse {
	s.Body = v
	return s
}

type AcceptFulfillmentDecisionRequest struct {
	DecisionConclusion *string `json:"DecisionConclusion,omitempty" xml:"DecisionConclusion,omitempty"`
	DecisionType       *string `json:"DecisionType,omitempty" xml:"DecisionType,omitempty"`
	OrderId            *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s AcceptFulfillmentDecisionRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptFulfillmentDecisionRequest) GoString() string {
	return s.String()
}

func (s *AcceptFulfillmentDecisionRequest) SetDecisionConclusion(v string) *AcceptFulfillmentDecisionRequest {
	s.DecisionConclusion = &v
	return s
}

func (s *AcceptFulfillmentDecisionRequest) SetDecisionType(v string) *AcceptFulfillmentDecisionRequest {
	s.DecisionType = &v
	return s
}

func (s *AcceptFulfillmentDecisionRequest) SetOrderId(v string) *AcceptFulfillmentDecisionRequest {
	s.OrderId = &v
	return s
}

type AcceptFulfillmentDecisionResponseBody struct {
	Data         *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AcceptFulfillmentDecisionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptFulfillmentDecisionResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptFulfillmentDecisionResponseBody) SetData(v bool) *AcceptFulfillmentDecisionResponseBody {
	s.Data = &v
	return s
}

func (s *AcceptFulfillmentDecisionResponseBody) SetErrorCode(v string) *AcceptFulfillmentDecisionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AcceptFulfillmentDecisionResponseBody) SetErrorMessage(v string) *AcceptFulfillmentDecisionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AcceptFulfillmentDecisionResponseBody) SetSuccess(v bool) *AcceptFulfillmentDecisionResponseBody {
	s.Success = &v
	return s
}

type AcceptFulfillmentDecisionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptFulfillmentDecisionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptFulfillmentDecisionResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptFulfillmentDecisionResponse) GoString() string {
	return s.String()
}

func (s *AcceptFulfillmentDecisionResponse) SetHeaders(v map[string]*string) *AcceptFulfillmentDecisionResponse {
	s.Headers = v
	return s
}

func (s *AcceptFulfillmentDecisionResponse) SetStatusCode(v int32) *AcceptFulfillmentDecisionResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptFulfillmentDecisionResponse) SetBody(v *AcceptFulfillmentDecisionResponseBody) *AcceptFulfillmentDecisionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("yunjian"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - CreateDemandPlanRequest
//
// @param headers - CreateDemandPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDemandPlanResponse
func (client *Client) CreateDemandPlanWithOptions(request *CreateDemandPlanRequest, headers *CreateDemandPlanHeaders, runtime *util.RuntimeOptions) (_result *CreateDemandPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCid)) {
		body["targetCid"] = request.TargetCid
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.YunUserId)) {
		realHeaders["Yun-User-Id"] = util.ToJSONString(headers.YunUserId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDemandPlan"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/demand/urgent/saveUrgentDemandPlanItem"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDemandPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateDemandPlanRequest
//
// @return CreateDemandPlanResponse
func (client *Client) CreateDemandPlan(request *CreateDemandPlanRequest) (_result *CreateDemandPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDemandPlanHeaders{}
	_result = &CreateDemandPlanResponse{}
	_body, _err := client.CreateDemandPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建plan2.0版本
//
// @param request - CreateDemandPlanV2Request
//
// @param headers - CreateDemandPlanV2Headers
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDemandPlanV2Response
func (client *Client) CreateDemandPlanV2WithOptions(request *CreateDemandPlanV2Request, headers *CreateDemandPlanV2Headers, runtime *util.RuntimeOptions) (_result *CreateDemandPlanV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["productType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetCid)) {
		body["targetCid"] = request.TargetCid
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.YunUserId)) {
		realHeaders["Yun-User-Id"] = util.ToJSONString(headers.YunUserId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDemandPlanV2"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/demand/urgent/saveDemandPlan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDemandPlanV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建plan2.0版本
//
// @param request - CreateDemandPlanV2Request
//
// @return CreateDemandPlanV2Response
func (client *Client) CreateDemandPlanV2(request *CreateDemandPlanV2Request) (_result *CreateDemandPlanV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateDemandPlanV2Headers{}
	_result = &CreateDemandPlanV2Response{}
	_body, _err := client.CreateDemandPlanV2WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 紧急需求单ite 删除
//
// @param request - DeleteUrgentDemandItemRequest
//
// @param headers - DeleteUrgentDemandItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUrgentDemandItemResponse
func (client *Client) DeleteUrgentDemandItemWithOptions(request *DeleteUrgentDemandItemRequest, headers *DeleteUrgentDemandItemHeaders, runtime *util.RuntimeOptions) (_result *DeleteUrgentDemandItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Modifier)) {
		query["modifier"] = request.Modifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.YunUserId)) {
		realHeaders["Yun-User-Id"] = util.ToJSONString(headers.YunUserId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUrgentDemandItem"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/demand/urgent/deleteUrgentDemandItem"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUrgentDemandItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 紧急需求单ite 删除
//
// @param request - DeleteUrgentDemandItemRequest
//
// @return DeleteUrgentDemandItemResponse
func (client *Client) DeleteUrgentDemandItem(request *DeleteUrgentDemandItemRequest) (_result *DeleteUrgentDemandItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteUrgentDemandItemHeaders{}
	_result = &DeleteUrgentDemandItemResponse{}
	_body, _err := client.DeleteUrgentDemandItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 紧急需求单plan删除
//
// @param request - DeleteUrgentDemandPlanRequest
//
// @param headers - DeleteUrgentDemandPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUrgentDemandPlanResponse
func (client *Client) DeleteUrgentDemandPlanWithOptions(request *DeleteUrgentDemandPlanRequest, headers *DeleteUrgentDemandPlanHeaders, runtime *util.RuntimeOptions) (_result *DeleteUrgentDemandPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Modifier)) {
		query["modifier"] = request.Modifier
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.YunUserId)) {
		realHeaders["Yun-User-Id"] = util.ToJSONString(headers.YunUserId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUrgentDemandPlan"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/demand/urgent/deleteUrgentDemandPlan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUrgentDemandPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 紧急需求单plan删除
//
// @param request - DeleteUrgentDemandPlanRequest
//
// @return DeleteUrgentDemandPlanResponse
func (client *Client) DeleteUrgentDemandPlan(request *DeleteUrgentDemandPlanRequest) (_result *DeleteUrgentDemandPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeleteUrgentDemandPlanHeaders{}
	_result = &DeleteUrgentDemandPlanResponse{}
	_body, _err := client.DeleteUrgentDemandPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 交付信息同步
//
// @param request - DeliveryItemDetailSynRequest
//
// @param headers - DeliveryItemDetailSynHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeliveryItemDetailSynResponse
func (client *Client) DeliveryItemDetailSynWithOptions(request *DeliveryItemDetailSynRequest, headers *DeliveryItemDetailSynHeaders, runtime *util.RuntimeOptions) (_result *DeliveryItemDetailSynResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryItemDetailDTOS)) {
		body["deliveryItemDetailDTOS"] = request.DeliveryItemDetailDTOS
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryItemId)) {
		body["deliveryItemId"] = request.DeliveryItemId
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryPlanId)) {
		body["deliveryPlanId"] = request.DeliveryPlanId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.YunUserId)) {
		realHeaders["Yun-User-Id"] = util.ToJSONString(headers.YunUserId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeliveryItemDetailSyn"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/supply/deliveryItemDataSync"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeliveryItemDetailSynResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 交付信息同步
//
// @param request - DeliveryItemDetailSynRequest
//
// @return DeliveryItemDetailSynResponse
func (client *Client) DeliveryItemDetailSyn(request *DeliveryItemDetailSynRequest) (_result *DeliveryItemDetailSynResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &DeliveryItemDetailSynHeaders{}
	_result = &DeliveryItemDetailSynResponse{}
	_body, _err := client.DeliveryItemDetailSynWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询报备单中报备项列表
//
// @param request - GetUrgentDemandItemListRequest
//
// @param headers - GetUrgentDemandItemListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUrgentDemandItemListResponse
func (client *Client) GetUrgentDemandItemListWithOptions(request *GetUrgentDemandItemListRequest, headers *GetUrgentDemandItemListHeaders, runtime *util.RuntimeOptions) (_result *GetUrgentDemandItemListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		body["commodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityTypeCode)) {
		body["commodityTypeCode"] = request.CommodityTypeCode
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		body["planId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Zone)) {
		body["zone"] = request.Zone
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.YunUserId)) {
		realHeaders["Yun-User-Id"] = util.ToJSONString(headers.YunUserId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUrgentDemandItemList"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/demand/urgent/getUrgentDemandItemList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUrgentDemandItemListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询报备单中报备项列表
//
// @param request - GetUrgentDemandItemListRequest
//
// @return GetUrgentDemandItemListResponse
func (client *Client) GetUrgentDemandItemList(request *GetUrgentDemandItemListRequest) (_result *GetUrgentDemandItemListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUrgentDemandItemListHeaders{}
	_result = &GetUrgentDemandItemListResponse{}
	_body, _err := client.GetUrgentDemandItemListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// getUrgentDemandPlanDetail
//
// @param request - GetUrgentDemandPlanDetailRequest
//
// @param headers - GetUrgentDemandPlanDetailHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUrgentDemandPlanDetailResponse
func (client *Client) GetUrgentDemandPlanDetailWithOptions(request *GetUrgentDemandPlanDetailRequest, headers *GetUrgentDemandPlanDetailHeaders, runtime *util.RuntimeOptions) (_result *GetUrgentDemandPlanDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		body["planId"] = request.PlanId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.YunUserId)) {
		realHeaders["Yun-User-Id"] = util.ToJSONString(headers.YunUserId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUrgentDemandPlanDetail"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/demand/urgent/getUrgentDemandPlanDetail"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUrgentDemandPlanDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// getUrgentDemandPlanDetail
//
// @param request - GetUrgentDemandPlanDetailRequest
//
// @return GetUrgentDemandPlanDetailResponse
func (client *Client) GetUrgentDemandPlanDetail(request *GetUrgentDemandPlanDetailRequest) (_result *GetUrgentDemandPlanDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUrgentDemandPlanDetailHeaders{}
	_result = &GetUrgentDemandPlanDetailResponse{}
	_body, _err := client.GetUrgentDemandPlanDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询报备单列表
//
// @param request - GetUrgentDemandPlanListRequest
//
// @param headers - GetUrgentDemandPlanListHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUrgentDemandPlanListResponse
func (client *Client) GetUrgentDemandPlanListWithOptions(request *GetUrgentDemandPlanListRequest, headers *GetUrgentDemandPlanListHeaders, runtime *util.RuntimeOptions) (_result *GetUrgentDemandPlanListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PlanType)) {
		body["planType"] = request.PlanType
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.YunUserId)) {
		realHeaders["Yun-User-Id"] = util.ToJSONString(headers.YunUserId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUrgentDemandPlanList"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/demand/urgent/getUrgentDemandPlanList"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUrgentDemandPlanListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询报备单列表
//
// @param request - GetUrgentDemandPlanListRequest
//
// @return GetUrgentDemandPlanListResponse
func (client *Client) GetUrgentDemandPlanList(request *GetUrgentDemandPlanListRequest) (_result *GetUrgentDemandPlanListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetUrgentDemandPlanListHeaders{}
	_result = &GetUrgentDemandPlanListResponse{}
	_body, _err := client.GetUrgentDemandPlanListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ecs资源方案
//
// @param request - PushResourcePlanRequest
//
// @param headers - PushResourcePlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushResourcePlanResponse
func (client *Client) PushResourcePlanWithOptions(request *PushResourcePlanRequest, headers *PushResourcePlanHeaders, runtime *util.RuntimeOptions) (_result *PushResourcePlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BufferCnt)) {
		body["bufferCnt"] = request.BufferCnt
	}

	if !tea.BoolValue(util.IsUnset(request.DemandCount)) {
		body["demandCount"] = request.DemandCount
	}

	if !tea.BoolValue(util.IsUnset(request.DemandId)) {
		body["demandId"] = request.DemandId
	}

	if !tea.BoolValue(util.IsUnset(request.MethodList)) {
		body["methodList"] = request.MethodList
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.RequireCnt)) {
		body["requireCnt"] = request.RequireCnt
	}

	if !tea.BoolValue(util.IsUnset(request.SubDemandId)) {
		body["subDemandId"] = request.SubDemandId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.YunUserId)) {
		realHeaders["Yun-User-Id"] = util.ToJSONString(headers.YunUserId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushResourcePlan"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/supply/resourcePlan/push"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PushResourcePlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ecs资源方案
//
// @param request - PushResourcePlanRequest
//
// @return PushResourcePlanResponse
func (client *Client) PushResourcePlan(request *PushResourcePlanRequest) (_result *PushResourcePlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PushResourcePlanHeaders{}
	_result = &PushResourcePlanResponse{}
	_body, _err := client.PushResourcePlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询accountId下所有存在交付状态（包括部分）的报备数据, 以及开通数据情况
//
// @param request - QueryDeliveredSupplyItemsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDeliveredSupplyItemsResponse
func (client *Client) QueryDeliveredSupplyItemsWithOptions(request *QueryDeliveredSupplyItemsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryDeliveredSupplyItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityTypeCode)) {
		query["commodityTypeCode"] = request.CommodityTypeCode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDeliveredSupplyItems"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/supply/queryDeliveredSupplyItems"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &QueryDeliveredSupplyItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询accountId下所有存在交付状态（包括部分）的报备数据, 以及开通数据情况
//
// @param request - QueryDeliveredSupplyItemsRequest
//
// @return QueryDeliveredSupplyItemsResponse
func (client *Client) QueryDeliveredSupplyItems(request *QueryDeliveredSupplyItemsRequest) (_result *QueryDeliveredSupplyItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryDeliveredSupplyItemsResponse{}
	_body, _err := client.QueryDeliveredSupplyItemsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询账单预算数据
//
// @param request - QueryPeriodBudgetBillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPeriodBudgetBillResponse
func (client *Client) QueryPeriodBudgetBillWithOptions(request *QueryPeriodBudgetBillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryPeriodBudgetBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectIds)) {
		query["objectIds"] = request.ObjectIds
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["objectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["period"] = request.Period
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPeriodBudgetBill"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/annual/budget/queryPeriodBudgetBill"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPeriodBudgetBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询账单预算数据
//
// @param request - QueryPeriodBudgetBillRequest
//
// @return QueryPeriodBudgetBillResponse
func (client *Client) QueryPeriodBudgetBill(request *QueryPeriodBudgetBillRequest) (_result *QueryPeriodBudgetBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryPeriodBudgetBillResponse{}
	_body, _err := client.QueryPeriodBudgetBillWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 紧急需求单item新增
//
// @param request - SaveUrgentDemandItemRequest
//
// @param headers - SaveUrgentDemandItemHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveUrgentDemandItemResponse
func (client *Client) SaveUrgentDemandItemWithOptions(request *SaveUrgentDemandItemRequest, headers *SaveUrgentDemandItemHeaders, runtime *util.RuntimeOptions) (_result *SaveUrgentDemandItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["accountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		body["creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.EffectTime)) {
		body["effectTime"] = request.EffectTime
	}

	if !tea.BoolValue(util.IsUnset(request.Modifier)) {
		body["modifier"] = request.Modifier
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		body["networkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.PayDuration)) {
		body["payDuration"] = request.PayDuration
	}

	if !tea.BoolValue(util.IsUnset(request.PayDurationUnit)) {
		body["payDurationUnit"] = request.PayDurationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		body["payType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		body["planId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.UrgentDemandEbsRequest)) {
		body["urgentDemandEbsRequest"] = request.UrgentDemandEbsRequest
	}

	if !tea.BoolValue(util.IsUnset(request.UrgentDemandEcsRequest)) {
		body["urgentDemandEcsRequest"] = request.UrgentDemandEcsRequest
	}

	if !tea.BoolValue(util.IsUnset(request.Zone)) {
		body["zone"] = request.Zone
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.YunUserId)) {
		realHeaders["Yun-User-Id"] = util.ToJSONString(headers.YunUserId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveUrgentDemandItem"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/demand/urgent/saveUrgentDemandItem"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveUrgentDemandItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 紧急需求单item新增
//
// @param request - SaveUrgentDemandItemRequest
//
// @return SaveUrgentDemandItemResponse
func (client *Client) SaveUrgentDemandItem(request *SaveUrgentDemandItemRequest) (_result *SaveUrgentDemandItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SaveUrgentDemandItemHeaders{}
	_result = &SaveUrgentDemandItemResponse{}
	_body, _err := client.SaveUrgentDemandItemWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// submitUrgentDemandPlan
//
// @param request - SubmitUrgentDemandPlanRequest
//
// @param headers - SubmitUrgentDemandPlanHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitUrgentDemandPlanResponse
func (client *Client) SubmitUrgentDemandPlanWithOptions(request *SubmitUrgentDemandPlanRequest, headers *SubmitUrgentDemandPlanHeaders, runtime *util.RuntimeOptions) (_result *SubmitUrgentDemandPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["planId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.YunUserId)) {
		realHeaders["Yun-User-Id"] = util.ToJSONString(headers.YunUserId)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitUrgentDemandPlan"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/demand/urgent/submitUrgentDemandPlan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitUrgentDemandPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// submitUrgentDemandPlan
//
// @param request - SubmitUrgentDemandPlanRequest
//
// @return SubmitUrgentDemandPlanResponse
func (client *Client) SubmitUrgentDemandPlan(request *SubmitUrgentDemandPlanRequest) (_result *SubmitUrgentDemandPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &SubmitUrgentDemandPlanHeaders{}
	_result = &SubmitUrgentDemandPlanResponse{}
	_body, _err := client.SubmitUrgentDemandPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 云产品交付决策反馈回调
//
// @param request - AcceptFulfillmentDecisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptFulfillmentDecisionResponse
func (client *Client) AcceptFulfillmentDecisionWithOptions(request *AcceptFulfillmentDecisionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AcceptFulfillmentDecisionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DecisionConclusion)) {
		body["DecisionConclusion"] = request.DecisionConclusion
	}

	if !tea.BoolValue(util.IsUnset(request.DecisionType)) {
		body["DecisionType"] = request.DecisionType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("acceptFulfillmentDecision"),
		Version:     tea.String("2021-12-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/acceptFulfillmentDecision"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AcceptFulfillmentDecisionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 云产品交付决策反馈回调
//
// @param request - AcceptFulfillmentDecisionRequest
//
// @return AcceptFulfillmentDecisionResponse
func (client *Client) AcceptFulfillmentDecision(request *AcceptFulfillmentDecisionRequest) (_result *AcceptFulfillmentDecisionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AcceptFulfillmentDecisionResponse{}
	_body, _err := client.AcceptFulfillmentDecisionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
