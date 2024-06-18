// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AcceptDemandRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SE20183A0Q7C5556
	BizId   *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AcceptDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptDemandRequest) GoString() string {
	return s.String()
}

func (s *AcceptDemandRequest) SetBizId(v string) *AcceptDemandRequest {
	s.BizId = &v
	return s
}

func (s *AcceptDemandRequest) SetMessage(v string) *AcceptDemandRequest {
	s.Message = &v
	return s
}

type AcceptDemandResponseBody struct {
	BindUrl *string `json:"BindUrl,omitempty" xml:"BindUrl,omitempty"`
	// example:
	//
	// 497F7522-82B0-4BD4-84FE-AE8749E4C2F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptDemandResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptDemandResponseBody) SetBindUrl(v string) *AcceptDemandResponseBody {
	s.BindUrl = &v
	return s
}

func (s *AcceptDemandResponseBody) SetRequestId(v string) *AcceptDemandResponseBody {
	s.RequestId = &v
	return s
}

type AcceptDemandResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptDemandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptDemandResponse) GoString() string {
	return s.String()
}

func (s *AcceptDemandResponse) SetHeaders(v map[string]*string) *AcceptDemandResponse {
	s.Headers = v
	return s
}

func (s *AcceptDemandResponse) SetStatusCode(v int32) *AcceptDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptDemandResponse) SetBody(v *AcceptDemandResponseBody) *AcceptDemandResponse {
	s.Body = v
	return s
}

type BidDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	AuctionId *string `json:"AuctionId,omitempty" xml:"AuctionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RMB
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	MaxBid *float32 `json:"MaxBid,omitempty" xml:"MaxBid,omitempty"`
}

func (s BidDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BidDomainRequest) GoString() string {
	return s.String()
}

func (s *BidDomainRequest) SetAuctionId(v string) *BidDomainRequest {
	s.AuctionId = &v
	return s
}

func (s *BidDomainRequest) SetCurrency(v string) *BidDomainRequest {
	s.Currency = &v
	return s
}

func (s *BidDomainRequest) SetMaxBid(v float32) *BidDomainRequest {
	s.MaxBid = &v
	return s
}

type BidDomainResponseBody struct {
	// example:
	//
	// 12345678
	AuctionId *string `json:"AuctionId,omitempty" xml:"AuctionId,omitempty"`
	// example:
	//
	// CC615585-9D93-4179-BD16-09337E32A3A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BidDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BidDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BidDomainResponseBody) SetAuctionId(v string) *BidDomainResponseBody {
	s.AuctionId = &v
	return s
}

func (s *BidDomainResponseBody) SetRequestId(v string) *BidDomainResponseBody {
	s.RequestId = &v
	return s
}

type BidDomainResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BidDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BidDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BidDomainResponse) GoString() string {
	return s.String()
}

func (s *BidDomainResponse) SetHeaders(v map[string]*string) *BidDomainResponse {
	s.Headers = v
	return s
}

func (s *BidDomainResponse) SetStatusCode(v int32) *BidDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *BidDomainResponse) SetBody(v *BidDomainResponseBody) *BidDomainResponse {
	s.Body = v
	return s
}

type ChangeAuctionRequest struct {
	AuctionList []*ChangeAuctionRequestAuctionList `json:"AuctionList,omitempty" xml:"AuctionList,omitempty" type:"Repeated"`
}

func (s ChangeAuctionRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeAuctionRequest) GoString() string {
	return s.String()
}

func (s *ChangeAuctionRequest) SetAuctionList(v []*ChangeAuctionRequestAuctionList) *ChangeAuctionRequest {
	s.AuctionList = v
	return s
}

type ChangeAuctionRequestAuctionList struct {
	BidRecords []*ChangeAuctionRequestAuctionListBidRecords `json:"BidRecords,omitempty" xml:"BidRecords,omitempty" type:"Repeated"`
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	EndTime      *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsReserve    *int32   `json:"IsReserve,omitempty" xml:"IsReserve,omitempty"`
	ReservePrice *float32 `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	ReserveRange *string  `json:"ReserveRange,omitempty" xml:"ReserveRange,omitempty"`
	Status       *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	TimeLeft     *int64   `json:"TimeLeft,omitempty" xml:"TimeLeft,omitempty"`
	// This parameter is required.
	Winner *string `json:"Winner,omitempty" xml:"Winner,omitempty"`
	// This parameter is required.
	WinnerPrice *float32 `json:"WinnerPrice,omitempty" xml:"WinnerPrice,omitempty"`
}

func (s ChangeAuctionRequestAuctionList) String() string {
	return tea.Prettify(s)
}

func (s ChangeAuctionRequestAuctionList) GoString() string {
	return s.String()
}

func (s *ChangeAuctionRequestAuctionList) SetBidRecords(v []*ChangeAuctionRequestAuctionListBidRecords) *ChangeAuctionRequestAuctionList {
	s.BidRecords = v
	return s
}

func (s *ChangeAuctionRequestAuctionList) SetDomainName(v string) *ChangeAuctionRequestAuctionList {
	s.DomainName = &v
	return s
}

func (s *ChangeAuctionRequestAuctionList) SetEndTime(v string) *ChangeAuctionRequestAuctionList {
	s.EndTime = &v
	return s
}

func (s *ChangeAuctionRequestAuctionList) SetIsReserve(v int32) *ChangeAuctionRequestAuctionList {
	s.IsReserve = &v
	return s
}

func (s *ChangeAuctionRequestAuctionList) SetReservePrice(v float32) *ChangeAuctionRequestAuctionList {
	s.ReservePrice = &v
	return s
}

func (s *ChangeAuctionRequestAuctionList) SetReserveRange(v string) *ChangeAuctionRequestAuctionList {
	s.ReserveRange = &v
	return s
}

func (s *ChangeAuctionRequestAuctionList) SetStatus(v string) *ChangeAuctionRequestAuctionList {
	s.Status = &v
	return s
}

func (s *ChangeAuctionRequestAuctionList) SetTimeLeft(v int64) *ChangeAuctionRequestAuctionList {
	s.TimeLeft = &v
	return s
}

func (s *ChangeAuctionRequestAuctionList) SetWinner(v string) *ChangeAuctionRequestAuctionList {
	s.Winner = &v
	return s
}

func (s *ChangeAuctionRequestAuctionList) SetWinnerPrice(v float32) *ChangeAuctionRequestAuctionList {
	s.WinnerPrice = &v
	return s
}

type ChangeAuctionRequestAuctionListBidRecords struct {
	// This parameter is required.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// This parameter is required.
	Price *float32 `json:"Price,omitempty" xml:"Price,omitempty"`
	// This parameter is required.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ChangeAuctionRequestAuctionListBidRecords) String() string {
	return tea.Prettify(s)
}

func (s ChangeAuctionRequestAuctionListBidRecords) GoString() string {
	return s.String()
}

func (s *ChangeAuctionRequestAuctionListBidRecords) SetCreateTime(v string) *ChangeAuctionRequestAuctionListBidRecords {
	s.CreateTime = &v
	return s
}

func (s *ChangeAuctionRequestAuctionListBidRecords) SetPrice(v float32) *ChangeAuctionRequestAuctionListBidRecords {
	s.Price = &v
	return s
}

func (s *ChangeAuctionRequestAuctionListBidRecords) SetUserId(v string) *ChangeAuctionRequestAuctionListBidRecords {
	s.UserId = &v
	return s
}

type ChangeAuctionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeAuctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeAuctionResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeAuctionResponseBody) SetRequestId(v string) *ChangeAuctionResponseBody {
	s.RequestId = &v
	return s
}

type ChangeAuctionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeAuctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeAuctionResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeAuctionResponse) GoString() string {
	return s.String()
}

func (s *ChangeAuctionResponse) SetHeaders(v map[string]*string) *ChangeAuctionResponse {
	s.Headers = v
	return s
}

func (s *ChangeAuctionResponse) SetStatusCode(v int32) *ChangeAuctionResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeAuctionResponse) SetBody(v *ChangeAuctionResponseBody) *ChangeAuctionResponse {
	s.Body = v
	return s
}

type CheckDomainStatusRequest struct {
	// This parameter is required.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s CheckDomainStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckDomainStatusRequest) SetDomain(v string) *CheckDomainStatusRequest {
	s.Domain = &v
	return s
}

type CheckDomainStatusResponseBody struct {
	ErrorCode      *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Module         *CheckDomainStatusResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckDomainStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainStatusResponseBody) SetErrorCode(v string) *CheckDomainStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckDomainStatusResponseBody) SetHttpStatusCode(v int32) *CheckDomainStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckDomainStatusResponseBody) SetModule(v *CheckDomainStatusResponseBodyModule) *CheckDomainStatusResponseBody {
	s.Module = v
	return s
}

func (s *CheckDomainStatusResponseBody) SetRequestId(v string) *CheckDomainStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDomainStatusResponseBody) SetSuccess(v bool) *CheckDomainStatusResponseBody {
	s.Success = &v
	return s
}

type CheckDomainStatusResponseBodyModule struct {
	DeadDate *int64   `json:"DeadDate,omitempty" xml:"DeadDate,omitempty"`
	Domain   *string  `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EndTime  *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Price    *float32 `json:"Price,omitempty" xml:"Price,omitempty"`
	RegDate  *int64   `json:"RegDate,omitempty" xml:"RegDate,omitempty"`
}

func (s CheckDomainStatusResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainStatusResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CheckDomainStatusResponseBodyModule) SetDeadDate(v int64) *CheckDomainStatusResponseBodyModule {
	s.DeadDate = &v
	return s
}

func (s *CheckDomainStatusResponseBodyModule) SetDomain(v string) *CheckDomainStatusResponseBodyModule {
	s.Domain = &v
	return s
}

func (s *CheckDomainStatusResponseBodyModule) SetEndTime(v int64) *CheckDomainStatusResponseBodyModule {
	s.EndTime = &v
	return s
}

func (s *CheckDomainStatusResponseBodyModule) SetPrice(v float32) *CheckDomainStatusResponseBodyModule {
	s.Price = &v
	return s
}

func (s *CheckDomainStatusResponseBodyModule) SetRegDate(v int64) *CheckDomainStatusResponseBodyModule {
	s.RegDate = &v
	return s
}

type CheckDomainStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDomainStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDomainStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDomainStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckDomainStatusResponse) SetHeaders(v map[string]*string) *CheckDomainStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckDomainStatusResponse) SetStatusCode(v int32) *CheckDomainStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDomainStatusResponse) SetBody(v *CheckDomainStatusResponseBody) *CheckDomainStatusResponse {
	s.Body = v
	return s
}

type CheckSelectedDomainStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s CheckSelectedDomainStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckSelectedDomainStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckSelectedDomainStatusRequest) SetDomain(v string) *CheckSelectedDomainStatusRequest {
	s.Domain = &v
	return s
}

type CheckSelectedDomainStatusResponseBody struct {
	// example:
	//
	// OssFileNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Module         *CheckSelectedDomainStatusResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// E2598CAF-DBFE-494E-95EF-B42A33C178AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckSelectedDomainStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSelectedDomainStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSelectedDomainStatusResponseBody) SetErrorCode(v string) *CheckSelectedDomainStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckSelectedDomainStatusResponseBody) SetHttpStatusCode(v int32) *CheckSelectedDomainStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckSelectedDomainStatusResponseBody) SetModule(v *CheckSelectedDomainStatusResponseBodyModule) *CheckSelectedDomainStatusResponseBody {
	s.Module = v
	return s
}

func (s *CheckSelectedDomainStatusResponseBody) SetRequestId(v string) *CheckSelectedDomainStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSelectedDomainStatusResponseBody) SetSuccess(v bool) *CheckSelectedDomainStatusResponseBody {
	s.Success = &v
	return s
}

type CheckSelectedDomainStatusResponseBodyModule struct {
	// example:
	//
	// 1567353497
	DeadDate *int64 `json:"DeadDate,omitempty" xml:"DeadDate,omitempty"`
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1567353497
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Premium *bool  `json:"Premium,omitempty" xml:"Premium,omitempty"`
	// example:
	//
	// 20.00
	Price *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// example:
	//
	// 1566353497
	RegDate *int64 `json:"RegDate,omitempty" xml:"RegDate,omitempty"`
}

func (s CheckSelectedDomainStatusResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s CheckSelectedDomainStatusResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CheckSelectedDomainStatusResponseBodyModule) SetDeadDate(v int64) *CheckSelectedDomainStatusResponseBodyModule {
	s.DeadDate = &v
	return s
}

func (s *CheckSelectedDomainStatusResponseBodyModule) SetDomain(v string) *CheckSelectedDomainStatusResponseBodyModule {
	s.Domain = &v
	return s
}

func (s *CheckSelectedDomainStatusResponseBodyModule) SetEndTime(v int64) *CheckSelectedDomainStatusResponseBodyModule {
	s.EndTime = &v
	return s
}

func (s *CheckSelectedDomainStatusResponseBodyModule) SetPremium(v bool) *CheckSelectedDomainStatusResponseBodyModule {
	s.Premium = &v
	return s
}

func (s *CheckSelectedDomainStatusResponseBodyModule) SetPrice(v float64) *CheckSelectedDomainStatusResponseBodyModule {
	s.Price = &v
	return s
}

func (s *CheckSelectedDomainStatusResponseBodyModule) SetRegDate(v int64) *CheckSelectedDomainStatusResponseBodyModule {
	s.RegDate = &v
	return s
}

type CheckSelectedDomainStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSelectedDomainStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSelectedDomainStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckSelectedDomainStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckSelectedDomainStatusResponse) SetHeaders(v map[string]*string) *CheckSelectedDomainStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckSelectedDomainStatusResponse) SetStatusCode(v int32) *CheckSelectedDomainStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSelectedDomainStatusResponse) SetBody(v *CheckSelectedDomainStatusResponseBody) *CheckSelectedDomainStatusResponse {
	s.Body = v
	return s
}

type CreateFixedPriceDemandOrderRequest struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Source    *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateFixedPriceDemandOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFixedPriceDemandOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateFixedPriceDemandOrderRequest) SetCode(v string) *CreateFixedPriceDemandOrderRequest {
	s.Code = &v
	return s
}

func (s *CreateFixedPriceDemandOrderRequest) SetContactId(v string) *CreateFixedPriceDemandOrderRequest {
	s.ContactId = &v
	return s
}

func (s *CreateFixedPriceDemandOrderRequest) SetDomain(v string) *CreateFixedPriceDemandOrderRequest {
	s.Domain = &v
	return s
}

func (s *CreateFixedPriceDemandOrderRequest) SetSource(v string) *CreateFixedPriceDemandOrderRequest {
	s.Source = &v
	return s
}

type CreateFixedPriceDemandOrderResponseBody struct {
	ErrorCode      *string                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Module         *CreateFixedPriceDemandOrderResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFixedPriceDemandOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFixedPriceDemandOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFixedPriceDemandOrderResponseBody) SetErrorCode(v string) *CreateFixedPriceDemandOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFixedPriceDemandOrderResponseBody) SetHttpStatusCode(v int32) *CreateFixedPriceDemandOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateFixedPriceDemandOrderResponseBody) SetModule(v *CreateFixedPriceDemandOrderResponseBodyModule) *CreateFixedPriceDemandOrderResponseBody {
	s.Module = v
	return s
}

func (s *CreateFixedPriceDemandOrderResponseBody) SetRequestId(v string) *CreateFixedPriceDemandOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFixedPriceDemandOrderResponseBody) SetSuccess(v bool) *CreateFixedPriceDemandOrderResponseBody {
	s.Success = &v
	return s
}

type CreateFixedPriceDemandOrderResponseBodyModule struct {
	Domain  *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	Price   *int64  `json:"Price,omitempty" xml:"Price,omitempty"`
}

func (s CreateFixedPriceDemandOrderResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s CreateFixedPriceDemandOrderResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateFixedPriceDemandOrderResponseBodyModule) SetDomain(v string) *CreateFixedPriceDemandOrderResponseBodyModule {
	s.Domain = &v
	return s
}

func (s *CreateFixedPriceDemandOrderResponseBodyModule) SetOrderNo(v string) *CreateFixedPriceDemandOrderResponseBodyModule {
	s.OrderNo = &v
	return s
}

func (s *CreateFixedPriceDemandOrderResponseBodyModule) SetPrice(v int64) *CreateFixedPriceDemandOrderResponseBodyModule {
	s.Price = &v
	return s
}

type CreateFixedPriceDemandOrderResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFixedPriceDemandOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFixedPriceDemandOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFixedPriceDemandOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateFixedPriceDemandOrderResponse) SetHeaders(v map[string]*string) *CreateFixedPriceDemandOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateFixedPriceDemandOrderResponse) SetStatusCode(v int32) *CreateFixedPriceDemandOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFixedPriceDemandOrderResponse) SetBody(v *CreateFixedPriceDemandOrderResponseBody) *CreateFixedPriceDemandOrderResponse {
	s.Body = v
	return s
}

type CreateFixedPriceSelectedOrderRequest struct {
	// example:
	//
	// DX123456
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11935401
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20.00
	ExpectedPrice *float64 `json:"ExpectedPrice,omitempty" xml:"ExpectedPrice,omitempty"`
	// example:
	//
	// partnername
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateFixedPriceSelectedOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFixedPriceSelectedOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateFixedPriceSelectedOrderRequest) SetCode(v string) *CreateFixedPriceSelectedOrderRequest {
	s.Code = &v
	return s
}

func (s *CreateFixedPriceSelectedOrderRequest) SetContactId(v string) *CreateFixedPriceSelectedOrderRequest {
	s.ContactId = &v
	return s
}

func (s *CreateFixedPriceSelectedOrderRequest) SetDomainName(v string) *CreateFixedPriceSelectedOrderRequest {
	s.DomainName = &v
	return s
}

func (s *CreateFixedPriceSelectedOrderRequest) SetExpectedPrice(v float64) *CreateFixedPriceSelectedOrderRequest {
	s.ExpectedPrice = &v
	return s
}

func (s *CreateFixedPriceSelectedOrderRequest) SetSource(v string) *CreateFixedPriceSelectedOrderRequest {
	s.Source = &v
	return s
}

type CreateFixedPriceSelectedOrderResponseBody struct {
	// example:
	//
	// DomainNotOnSale
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Module         *CreateFixedPriceSelectedOrderResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// C50E41A0-09F1-4491-8DB8-AF55BD2D0CC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFixedPriceSelectedOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFixedPriceSelectedOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFixedPriceSelectedOrderResponseBody) SetErrorCode(v string) *CreateFixedPriceSelectedOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFixedPriceSelectedOrderResponseBody) SetHttpStatusCode(v int32) *CreateFixedPriceSelectedOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateFixedPriceSelectedOrderResponseBody) SetModule(v *CreateFixedPriceSelectedOrderResponseBodyModule) *CreateFixedPriceSelectedOrderResponseBody {
	s.Module = v
	return s
}

func (s *CreateFixedPriceSelectedOrderResponseBody) SetRequestId(v string) *CreateFixedPriceSelectedOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFixedPriceSelectedOrderResponseBody) SetSuccess(v bool) *CreateFixedPriceSelectedOrderResponseBody {
	s.Success = &v
	return s
}

type CreateFixedPriceSelectedOrderResponseBodyModule struct {
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 31199295f2074ce895645d386cb22c36
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	// example:
	//
	// 20.00
	Price *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
}

func (s CreateFixedPriceSelectedOrderResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s CreateFixedPriceSelectedOrderResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateFixedPriceSelectedOrderResponseBodyModule) SetDomain(v string) *CreateFixedPriceSelectedOrderResponseBodyModule {
	s.Domain = &v
	return s
}

func (s *CreateFixedPriceSelectedOrderResponseBodyModule) SetOrderNo(v string) *CreateFixedPriceSelectedOrderResponseBodyModule {
	s.OrderNo = &v
	return s
}

func (s *CreateFixedPriceSelectedOrderResponseBodyModule) SetPrice(v int64) *CreateFixedPriceSelectedOrderResponseBodyModule {
	s.Price = &v
	return s
}

type CreateFixedPriceSelectedOrderResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFixedPriceSelectedOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFixedPriceSelectedOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFixedPriceSelectedOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateFixedPriceSelectedOrderResponse) SetHeaders(v map[string]*string) *CreateFixedPriceSelectedOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateFixedPriceSelectedOrderResponse) SetStatusCode(v int32) *CreateFixedPriceSelectedOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFixedPriceSelectedOrderResponse) SetBody(v *CreateFixedPriceSelectedOrderResponseBody) *CreateFixedPriceSelectedOrderResponse {
	s.Body = v
	return s
}

type FailDemandRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SE20183A0Q7C5556
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// some message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s FailDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s FailDemandRequest) GoString() string {
	return s.String()
}

func (s *FailDemandRequest) SetBizId(v string) *FailDemandRequest {
	s.BizId = &v
	return s
}

func (s *FailDemandRequest) SetMessage(v string) *FailDemandRequest {
	s.Message = &v
	return s
}

type FailDemandResponseBody struct {
	// example:
	//
	// 497F7522-82B0-4BD4-84FE-AE8749E4C2F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FailDemandResponseBody) GoString() string {
	return s.String()
}

func (s *FailDemandResponseBody) SetRequestId(v string) *FailDemandResponseBody {
	s.RequestId = &v
	return s
}

type FailDemandResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FailDemandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FailDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s FailDemandResponse) GoString() string {
	return s.String()
}

func (s *FailDemandResponse) SetHeaders(v map[string]*string) *FailDemandResponse {
	s.Headers = v
	return s
}

func (s *FailDemandResponse) SetStatusCode(v int32) *FailDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *FailDemandResponse) SetBody(v *FailDemandResponseBody) *FailDemandResponse {
	s.Body = v
	return s
}

type FinishDemandRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SE20183A0Q7C5556
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// some message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s FinishDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishDemandRequest) GoString() string {
	return s.String()
}

func (s *FinishDemandRequest) SetBizId(v string) *FinishDemandRequest {
	s.BizId = &v
	return s
}

func (s *FinishDemandRequest) SetMessage(v string) *FinishDemandRequest {
	s.Message = &v
	return s
}

type FinishDemandResponseBody struct {
	// example:
	//
	// 497F7522-82B0-4BD4-84FE-AE8749E4C2F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FinishDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishDemandResponseBody) GoString() string {
	return s.String()
}

func (s *FinishDemandResponseBody) SetRequestId(v string) *FinishDemandResponseBody {
	s.RequestId = &v
	return s
}

type FinishDemandResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinishDemandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinishDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishDemandResponse) GoString() string {
	return s.String()
}

func (s *FinishDemandResponse) SetHeaders(v map[string]*string) *FinishDemandResponse {
	s.Headers = v
	return s
}

func (s *FinishDemandResponse) SetStatusCode(v int32) *FinishDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishDemandResponse) SetBody(v *FinishDemandResponseBody) *FinishDemandResponse {
	s.Body = v
	return s
}

type GetIntlDomainDownloadUrlResponseBody struct {
	AllowRetry     *bool         `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	AppName        *string       `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DynamicCode    *string       `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	ErrorCode      *string       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg       *string       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	HttpStatusCode *int32        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
	Url            *string       `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetIntlDomainDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIntlDomainDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntlDomainDownloadUrlResponseBody) SetAllowRetry(v bool) *GetIntlDomainDownloadUrlResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetIntlDomainDownloadUrlResponseBody) SetAppName(v string) *GetIntlDomainDownloadUrlResponseBody {
	s.AppName = &v
	return s
}

func (s *GetIntlDomainDownloadUrlResponseBody) SetDynamicCode(v string) *GetIntlDomainDownloadUrlResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetIntlDomainDownloadUrlResponseBody) SetDynamicMessage(v string) *GetIntlDomainDownloadUrlResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetIntlDomainDownloadUrlResponseBody) SetErrorArgs(v []interface{}) *GetIntlDomainDownloadUrlResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetIntlDomainDownloadUrlResponseBody) SetErrorCode(v string) *GetIntlDomainDownloadUrlResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetIntlDomainDownloadUrlResponseBody) SetErrorMsg(v string) *GetIntlDomainDownloadUrlResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetIntlDomainDownloadUrlResponseBody) SetHttpStatusCode(v int32) *GetIntlDomainDownloadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetIntlDomainDownloadUrlResponseBody) SetRequestId(v string) *GetIntlDomainDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntlDomainDownloadUrlResponseBody) SetSuccess(v bool) *GetIntlDomainDownloadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetIntlDomainDownloadUrlResponseBody) SetUrl(v string) *GetIntlDomainDownloadUrlResponseBody {
	s.Url = &v
	return s
}

type GetIntlDomainDownloadUrlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIntlDomainDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIntlDomainDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIntlDomainDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetIntlDomainDownloadUrlResponse) SetHeaders(v map[string]*string) *GetIntlDomainDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetIntlDomainDownloadUrlResponse) SetStatusCode(v int32) *GetIntlDomainDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntlDomainDownloadUrlResponse) SetBody(v *GetIntlDomainDownloadUrlResponseBody) *GetIntlDomainDownloadUrlResponse {
	s.Body = v
	return s
}

type GetReserveDomainUrlResponseBody struct {
	// example:
	//
	// D34B02AE-09AF-41C1-A6D3-951A2233EDB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// http://example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetReserveDomainUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetReserveDomainUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetReserveDomainUrlResponseBody) SetRequestId(v string) *GetReserveDomainUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetReserveDomainUrlResponseBody) SetUrl(v string) *GetReserveDomainUrlResponseBody {
	s.Url = &v
	return s
}

type GetReserveDomainUrlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReserveDomainUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReserveDomainUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetReserveDomainUrlResponse) GoString() string {
	return s.String()
}

func (s *GetReserveDomainUrlResponse) SetHeaders(v map[string]*string) *GetReserveDomainUrlResponse {
	s.Headers = v
	return s
}

func (s *GetReserveDomainUrlResponse) SetStatusCode(v int32) *GetReserveDomainUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReserveDomainUrlResponse) SetBody(v *GetReserveDomainUrlResponseBody) *GetReserveDomainUrlResponse {
	s.Body = v
	return s
}

type PurchaseIntlDomainRequest struct {
	// This parameter is required.
	AuctionId *string `json:"AuctionId,omitempty" xml:"AuctionId,omitempty"`
	// This parameter is required.
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// This parameter is required.
	Price *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
}

func (s PurchaseIntlDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s PurchaseIntlDomainRequest) GoString() string {
	return s.String()
}

func (s *PurchaseIntlDomainRequest) SetAuctionId(v string) *PurchaseIntlDomainRequest {
	s.AuctionId = &v
	return s
}

func (s *PurchaseIntlDomainRequest) SetCurrency(v string) *PurchaseIntlDomainRequest {
	s.Currency = &v
	return s
}

func (s *PurchaseIntlDomainRequest) SetPrice(v float64) *PurchaseIntlDomainRequest {
	s.Price = &v
	return s
}

type PurchaseIntlDomainResponseBody struct {
	AllowRetry     *bool         `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	AppName        *string       `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AuctionId      *string       `json:"AuctionId,omitempty" xml:"AuctionId,omitempty"`
	DynamicCode    *string       `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	ErrorCode      *string       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg       *string       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	HttpStatusCode *int32        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PurchaseIntlDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PurchaseIntlDomainResponseBody) GoString() string {
	return s.String()
}

func (s *PurchaseIntlDomainResponseBody) SetAllowRetry(v bool) *PurchaseIntlDomainResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *PurchaseIntlDomainResponseBody) SetAppName(v string) *PurchaseIntlDomainResponseBody {
	s.AppName = &v
	return s
}

func (s *PurchaseIntlDomainResponseBody) SetAuctionId(v string) *PurchaseIntlDomainResponseBody {
	s.AuctionId = &v
	return s
}

func (s *PurchaseIntlDomainResponseBody) SetDynamicCode(v string) *PurchaseIntlDomainResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PurchaseIntlDomainResponseBody) SetDynamicMessage(v string) *PurchaseIntlDomainResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PurchaseIntlDomainResponseBody) SetErrorArgs(v []interface{}) *PurchaseIntlDomainResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *PurchaseIntlDomainResponseBody) SetErrorCode(v string) *PurchaseIntlDomainResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PurchaseIntlDomainResponseBody) SetErrorMsg(v string) *PurchaseIntlDomainResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PurchaseIntlDomainResponseBody) SetHttpStatusCode(v int32) *PurchaseIntlDomainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PurchaseIntlDomainResponseBody) SetRequestId(v string) *PurchaseIntlDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurchaseIntlDomainResponseBody) SetSuccess(v bool) *PurchaseIntlDomainResponseBody {
	s.Success = &v
	return s
}

type PurchaseIntlDomainResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseIntlDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurchaseIntlDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s PurchaseIntlDomainResponse) GoString() string {
	return s.String()
}

func (s *PurchaseIntlDomainResponse) SetHeaders(v map[string]*string) *PurchaseIntlDomainResponse {
	s.Headers = v
	return s
}

func (s *PurchaseIntlDomainResponse) SetStatusCode(v int32) *PurchaseIntlDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *PurchaseIntlDomainResponse) SetBody(v *PurchaseIntlDomainResponseBody) *PurchaseIntlDomainResponse {
	s.Body = v
	return s
}

type QueryAuctionDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AuctionId *string `json:"AuctionId,omitempty" xml:"AuctionId,omitempty"`
}

func (s QueryAuctionDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAuctionDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryAuctionDetailRequest) SetAuctionId(v string) *QueryAuctionDetailRequest {
	s.AuctionId = &v
	return s
}

type QueryAuctionDetailResponseBody struct {
	// example:
	//
	// 1515961936000
	AuctionEndTime *int64 `json:"AuctionEndTime,omitempty" xml:"AuctionEndTime,omitempty"`
	// example:
	//
	// 123456
	AuctionId *string `json:"AuctionId,omitempty" xml:"AuctionId,omitempty"`
	// example:
	//
	// 1515961936000
	BookEndTime *int64 `json:"BookEndTime,omitempty" xml:"BookEndTime,omitempty"`
	// example:
	//
	// 4
	BookedPartner *string `json:"BookedPartner,omitempty" xml:"BookedPartner,omitempty"`
	// example:
	//
	// RMB
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 1515961936000
	DeliveryTime *int64 `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// example:
	//
	// test.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 0
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// example:
	//
	// 0
	FailCode *string `json:"FailCode,omitempty" xml:"FailCode,omitempty"`
	// example:
	//
	// 100
	HighBid *float32 `json:"HighBid,omitempty" xml:"HighBid,omitempty"`
	// example:
	//
	// abc
	HighBidder *string `json:"HighBidder,omitempty" xml:"HighBidder,omitempty"`
	// example:
	//
	// 110
	NextValidBid *float32 `json:"NextValidBid,omitempty" xml:"NextValidBid,omitempty"`
	// example:
	//
	// 4
	PartnerType *string `json:"PartnerType,omitempty" xml:"PartnerType,omitempty"`
	// example:
	//
	// 1515961936000
	PayEndTime *int64 `json:"PayEndTime,omitempty" xml:"PayEndTime,omitempty"`
	// example:
	//
	// 200
	PayPrice *float32 `json:"PayPrice,omitempty" xml:"PayPrice,omitempty"`
	// example:
	//
	// 1
	PayStatus *string `json:"PayStatus,omitempty" xml:"PayStatus,omitempty"`
	// example:
	//
	// 0
	ProduceStatus *string `json:"ProduceStatus,omitempty" xml:"ProduceStatus,omitempty"`
	// example:
	//
	// UUID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	ReserveMet   *bool    `json:"ReserveMet,omitempty" xml:"ReserveMet,omitempty"`
	ReservePrice *float32 `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 50
	TransferInPrice *float32 `json:"TransferInPrice,omitempty" xml:"TransferInPrice,omitempty"`
	// example:
	//
	// 100
	YourCurrentBid *float32 `json:"YourCurrentBid,omitempty" xml:"YourCurrentBid,omitempty"`
	// example:
	//
	// 120
	YourMaxBid *float32 `json:"YourMaxBid,omitempty" xml:"YourMaxBid,omitempty"`
}

func (s QueryAuctionDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAuctionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAuctionDetailResponseBody) SetAuctionEndTime(v int64) *QueryAuctionDetailResponseBody {
	s.AuctionEndTime = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetAuctionId(v string) *QueryAuctionDetailResponseBody {
	s.AuctionId = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetBookEndTime(v int64) *QueryAuctionDetailResponseBody {
	s.BookEndTime = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetBookedPartner(v string) *QueryAuctionDetailResponseBody {
	s.BookedPartner = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetCurrency(v string) *QueryAuctionDetailResponseBody {
	s.Currency = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetDeliveryTime(v int64) *QueryAuctionDetailResponseBody {
	s.DeliveryTime = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetDomainName(v string) *QueryAuctionDetailResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetDomainType(v string) *QueryAuctionDetailResponseBody {
	s.DomainType = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetFailCode(v string) *QueryAuctionDetailResponseBody {
	s.FailCode = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetHighBid(v float32) *QueryAuctionDetailResponseBody {
	s.HighBid = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetHighBidder(v string) *QueryAuctionDetailResponseBody {
	s.HighBidder = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetNextValidBid(v float32) *QueryAuctionDetailResponseBody {
	s.NextValidBid = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetPartnerType(v string) *QueryAuctionDetailResponseBody {
	s.PartnerType = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetPayEndTime(v int64) *QueryAuctionDetailResponseBody {
	s.PayEndTime = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetPayPrice(v float32) *QueryAuctionDetailResponseBody {
	s.PayPrice = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetPayStatus(v string) *QueryAuctionDetailResponseBody {
	s.PayStatus = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetProduceStatus(v string) *QueryAuctionDetailResponseBody {
	s.ProduceStatus = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetRequestId(v string) *QueryAuctionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetReserveMet(v bool) *QueryAuctionDetailResponseBody {
	s.ReserveMet = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetReservePrice(v float32) *QueryAuctionDetailResponseBody {
	s.ReservePrice = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetStatus(v string) *QueryAuctionDetailResponseBody {
	s.Status = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetTransferInPrice(v float32) *QueryAuctionDetailResponseBody {
	s.TransferInPrice = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetYourCurrentBid(v float32) *QueryAuctionDetailResponseBody {
	s.YourCurrentBid = &v
	return s
}

func (s *QueryAuctionDetailResponseBody) SetYourMaxBid(v float32) *QueryAuctionDetailResponseBody {
	s.YourMaxBid = &v
	return s
}

type QueryAuctionDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAuctionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAuctionDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAuctionDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryAuctionDetailResponse) SetHeaders(v map[string]*string) *QueryAuctionDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryAuctionDetailResponse) SetStatusCode(v int32) *QueryAuctionDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAuctionDetailResponse) SetBody(v *QueryAuctionDetailResponseBody) *QueryAuctionDetailResponse {
	s.Body = v
	return s
}

type QueryAuctionsRequest struct {
	AuctionEndTimeOrder *string `json:"AuctionEndTimeOrder,omitempty" xml:"AuctionEndTimeOrder,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Statuses *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
}

func (s QueryAuctionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAuctionsRequest) GoString() string {
	return s.String()
}

func (s *QueryAuctionsRequest) SetAuctionEndTimeOrder(v string) *QueryAuctionsRequest {
	s.AuctionEndTimeOrder = &v
	return s
}

func (s *QueryAuctionsRequest) SetCurrentPage(v int32) *QueryAuctionsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryAuctionsRequest) SetPageSize(v int32) *QueryAuctionsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAuctionsRequest) SetStatus(v string) *QueryAuctionsRequest {
	s.Status = &v
	return s
}

func (s *QueryAuctionsRequest) SetStatuses(v string) *QueryAuctionsRequest {
	s.Statuses = &v
	return s
}

type QueryAuctionsResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                           `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryAuctionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// UUID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryAuctionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAuctionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAuctionsResponseBody) SetCurrentPageNum(v int32) *QueryAuctionsResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryAuctionsResponseBody) SetData(v []*QueryAuctionsResponseBodyData) *QueryAuctionsResponseBody {
	s.Data = v
	return s
}

func (s *QueryAuctionsResponseBody) SetPageSize(v int32) *QueryAuctionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryAuctionsResponseBody) SetRequestId(v string) *QueryAuctionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAuctionsResponseBody) SetTotalItemNum(v int32) *QueryAuctionsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryAuctionsResponseBody) SetTotalPageNum(v int32) *QueryAuctionsResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryAuctionsResponseBodyData struct {
	// example:
	//
	// 1515961936000
	AuctionEndTime *int64 `json:"AuctionEndTime,omitempty" xml:"AuctionEndTime,omitempty"`
	// example:
	//
	// 123456
	AuctionId *string `json:"AuctionId,omitempty" xml:"AuctionId,omitempty"`
	// example:
	//
	// 1515961936000
	BookEndTime *int64 `json:"BookEndTime,omitempty" xml:"BookEndTime,omitempty"`
	// example:
	//
	// 4
	BookedPartner *string `json:"BookedPartner,omitempty" xml:"BookedPartner,omitempty"`
	// example:
	//
	// RMB
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 1515961936000
	DeliveryTime *int64 `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// example:
	//
	// test.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 0
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// example:
	//
	// 0
	FailCode *string `json:"FailCode,omitempty" xml:"FailCode,omitempty"`
	// example:
	//
	// 100
	HighBid *float32 `json:"HighBid,omitempty" xml:"HighBid,omitempty"`
	// example:
	//
	// abc
	HighBidder *string `json:"HighBidder,omitempty" xml:"HighBidder,omitempty"`
	// example:
	//
	// 110
	NextValidBid *float32 `json:"NextValidBid,omitempty" xml:"NextValidBid,omitempty"`
	// example:
	//
	// 4
	PartnerType *string `json:"PartnerType,omitempty" xml:"PartnerType,omitempty"`
	// example:
	//
	// 1515961936000
	PayEndTime *int64 `json:"PayEndTime,omitempty" xml:"PayEndTime,omitempty"`
	// example:
	//
	// 200
	PayPrice *float32 `json:"PayPrice,omitempty" xml:"PayPrice,omitempty"`
	// example:
	//
	// 1
	PayStatus *string `json:"PayStatus,omitempty" xml:"PayStatus,omitempty"`
	// example:
	//
	// 0
	ProduceStatus *string `json:"ProduceStatus,omitempty" xml:"ProduceStatus,omitempty"`
	ReserveMax    *int64  `json:"ReserveMax,omitempty" xml:"ReserveMax,omitempty"`
	// example:
	//
	// true
	ReserveMet   *bool  `json:"ReserveMet,omitempty" xml:"ReserveMet,omitempty"`
	ReserveMin   *int64 `json:"ReserveMin,omitempty" xml:"ReserveMin,omitempty"`
	ReservePrice *int64 `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 50
	TransferInPrice *float32 `json:"TransferInPrice,omitempty" xml:"TransferInPrice,omitempty"`
	// example:
	//
	// 100
	YourCurrentBid *float32 `json:"YourCurrentBid,omitempty" xml:"YourCurrentBid,omitempty"`
	// example:
	//
	// 120
	YourMaxBid *float32 `json:"YourMaxBid,omitempty" xml:"YourMaxBid,omitempty"`
}

func (s QueryAuctionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAuctionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAuctionsResponseBodyData) SetAuctionEndTime(v int64) *QueryAuctionsResponseBodyData {
	s.AuctionEndTime = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetAuctionId(v string) *QueryAuctionsResponseBodyData {
	s.AuctionId = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetBookEndTime(v int64) *QueryAuctionsResponseBodyData {
	s.BookEndTime = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetBookedPartner(v string) *QueryAuctionsResponseBodyData {
	s.BookedPartner = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetCurrency(v string) *QueryAuctionsResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetDeliveryTime(v int64) *QueryAuctionsResponseBodyData {
	s.DeliveryTime = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetDomainName(v string) *QueryAuctionsResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetDomainType(v string) *QueryAuctionsResponseBodyData {
	s.DomainType = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetFailCode(v string) *QueryAuctionsResponseBodyData {
	s.FailCode = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetHighBid(v float32) *QueryAuctionsResponseBodyData {
	s.HighBid = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetHighBidder(v string) *QueryAuctionsResponseBodyData {
	s.HighBidder = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetNextValidBid(v float32) *QueryAuctionsResponseBodyData {
	s.NextValidBid = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetPartnerType(v string) *QueryAuctionsResponseBodyData {
	s.PartnerType = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetPayEndTime(v int64) *QueryAuctionsResponseBodyData {
	s.PayEndTime = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetPayPrice(v float32) *QueryAuctionsResponseBodyData {
	s.PayPrice = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetPayStatus(v string) *QueryAuctionsResponseBodyData {
	s.PayStatus = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetProduceStatus(v string) *QueryAuctionsResponseBodyData {
	s.ProduceStatus = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetReserveMax(v int64) *QueryAuctionsResponseBodyData {
	s.ReserveMax = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetReserveMet(v bool) *QueryAuctionsResponseBodyData {
	s.ReserveMet = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetReserveMin(v int64) *QueryAuctionsResponseBodyData {
	s.ReserveMin = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetReservePrice(v int64) *QueryAuctionsResponseBodyData {
	s.ReservePrice = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetStatus(v string) *QueryAuctionsResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetTransferInPrice(v float32) *QueryAuctionsResponseBodyData {
	s.TransferInPrice = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetYourCurrentBid(v float32) *QueryAuctionsResponseBodyData {
	s.YourCurrentBid = &v
	return s
}

func (s *QueryAuctionsResponseBodyData) SetYourMaxBid(v float32) *QueryAuctionsResponseBodyData {
	s.YourMaxBid = &v
	return s
}

type QueryAuctionsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAuctionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAuctionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAuctionsResponse) GoString() string {
	return s.String()
}

func (s *QueryAuctionsResponse) SetHeaders(v map[string]*string) *QueryAuctionsResponse {
	s.Headers = v
	return s
}

func (s *QueryAuctionsResponse) SetStatusCode(v int32) *QueryAuctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAuctionsResponse) SetBody(v *QueryAuctionsResponseBody) *QueryAuctionsResponse {
	s.Body = v
	return s
}

type QueryBidRecordsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AuctionId *string `json:"AuctionId,omitempty" xml:"AuctionId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryBidRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBidRecordsRequest) GoString() string {
	return s.String()
}

func (s *QueryBidRecordsRequest) SetAuctionId(v string) *QueryBidRecordsRequest {
	s.AuctionId = &v
	return s
}

func (s *QueryBidRecordsRequest) SetCurrentPage(v int32) *QueryBidRecordsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryBidRecordsRequest) SetPageSize(v int32) *QueryBidRecordsRequest {
	s.PageSize = &v
	return s
}

type QueryBidRecordsResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                             `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryBidRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// UUID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryBidRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBidRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBidRecordsResponseBody) SetCurrentPageNum(v int32) *QueryBidRecordsResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryBidRecordsResponseBody) SetData(v []*QueryBidRecordsResponseBodyData) *QueryBidRecordsResponseBody {
	s.Data = v
	return s
}

func (s *QueryBidRecordsResponseBody) SetPageSize(v int32) *QueryBidRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryBidRecordsResponseBody) SetRequestId(v string) *QueryBidRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBidRecordsResponseBody) SetTotalItemNum(v int32) *QueryBidRecordsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryBidRecordsResponseBody) SetTotalPageNum(v int32) *QueryBidRecordsResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryBidRecordsResponseBodyData struct {
	// example:
	//
	// 50
	Bid *float32 `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// example:
	//
	// 1515961936000
	BidTime *int64 `json:"BidTime,omitempty" xml:"BidTime,omitempty"`
	// example:
	//
	// abc
	Bidder *string `json:"Bidder,omitempty" xml:"Bidder,omitempty"`
	// example:
	//
	// RMB
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// test.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s QueryBidRecordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryBidRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBidRecordsResponseBodyData) SetBid(v float32) *QueryBidRecordsResponseBodyData {
	s.Bid = &v
	return s
}

func (s *QueryBidRecordsResponseBodyData) SetBidTime(v int64) *QueryBidRecordsResponseBodyData {
	s.BidTime = &v
	return s
}

func (s *QueryBidRecordsResponseBodyData) SetBidder(v string) *QueryBidRecordsResponseBodyData {
	s.Bidder = &v
	return s
}

func (s *QueryBidRecordsResponseBodyData) SetCurrency(v string) *QueryBidRecordsResponseBodyData {
	s.Currency = &v
	return s
}

func (s *QueryBidRecordsResponseBodyData) SetDomainName(v string) *QueryBidRecordsResponseBodyData {
	s.DomainName = &v
	return s
}

type QueryBidRecordsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBidRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBidRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBidRecordsResponse) GoString() string {
	return s.String()
}

func (s *QueryBidRecordsResponse) SetHeaders(v map[string]*string) *QueryBidRecordsResponse {
	s.Headers = v
	return s
}

func (s *QueryBidRecordsResponse) SetStatusCode(v int32) *QueryBidRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBidRecordsResponse) SetBody(v *QueryBidRecordsResponseBody) *QueryBidRecordsResponse {
	s.Body = v
	return s
}

type QueryBookingDomainInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s QueryBookingDomainInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBookingDomainInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryBookingDomainInfoRequest) SetDomainName(v string) *QueryBookingDomainInfoRequest {
	s.DomainName = &v
	return s
}

type QueryBookingDomainInfoResponseBody struct {
	// example:
	//
	// 1234
	AuctionId *int32 `json:"AuctionId,omitempty" xml:"AuctionId,omitempty"`
	// example:
	//
	// 1517985730419
	BookEndTime *int64 `json:"BookEndTime,omitempty" xml:"BookEndTime,omitempty"`
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 15
	MaxBid *float32 `json:"MaxBid,omitempty" xml:"MaxBid,omitempty"`
	// example:
	//
	// 4
	PartnerType *string `json:"PartnerType,omitempty" xml:"PartnerType,omitempty"`
	// example:
	//
	// 234234njhjkhkj
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnatchNo  *string `json:"SnatchNo,omitempty" xml:"SnatchNo,omitempty"`
	// example:
	//
	// 17
	TransferInPrice *float32 `json:"TransferInPrice,omitempty" xml:"TransferInPrice,omitempty"`
}

func (s QueryBookingDomainInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBookingDomainInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBookingDomainInfoResponseBody) SetAuctionId(v int32) *QueryBookingDomainInfoResponseBody {
	s.AuctionId = &v
	return s
}

func (s *QueryBookingDomainInfoResponseBody) SetBookEndTime(v int64) *QueryBookingDomainInfoResponseBody {
	s.BookEndTime = &v
	return s
}

func (s *QueryBookingDomainInfoResponseBody) SetCurrency(v string) *QueryBookingDomainInfoResponseBody {
	s.Currency = &v
	return s
}

func (s *QueryBookingDomainInfoResponseBody) SetMaxBid(v float32) *QueryBookingDomainInfoResponseBody {
	s.MaxBid = &v
	return s
}

func (s *QueryBookingDomainInfoResponseBody) SetPartnerType(v string) *QueryBookingDomainInfoResponseBody {
	s.PartnerType = &v
	return s
}

func (s *QueryBookingDomainInfoResponseBody) SetRequestId(v string) *QueryBookingDomainInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBookingDomainInfoResponseBody) SetSnatchNo(v string) *QueryBookingDomainInfoResponseBody {
	s.SnatchNo = &v
	return s
}

func (s *QueryBookingDomainInfoResponseBody) SetTransferInPrice(v float32) *QueryBookingDomainInfoResponseBody {
	s.TransferInPrice = &v
	return s
}

type QueryBookingDomainInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBookingDomainInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBookingDomainInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBookingDomainInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryBookingDomainInfoResponse) SetHeaders(v map[string]*string) *QueryBookingDomainInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryBookingDomainInfoResponse) SetStatusCode(v int32) *QueryBookingDomainInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBookingDomainInfoResponse) SetBody(v *QueryBookingDomainInfoResponseBody) *QueryBookingDomainInfoResponse {
	s.Body = v
	return s
}

type QueryBrokerDemandRequest struct {
	// example:
	//
	// SE20183915FI0178
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryBrokerDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBrokerDemandRequest) GoString() string {
	return s.String()
}

func (s *QueryBrokerDemandRequest) SetBizId(v string) *QueryBrokerDemandRequest {
	s.BizId = &v
	return s
}

func (s *QueryBrokerDemandRequest) SetCurrentPage(v int32) *QueryBrokerDemandRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryBrokerDemandRequest) SetPageSize(v int32) *QueryBrokerDemandRequest {
	s.PageSize = &v
	return s
}

func (s *QueryBrokerDemandRequest) SetStatus(v string) *QueryBrokerDemandRequest {
	s.Status = &v
	return s
}

type QueryBrokerDemandResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                               `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryBrokerDemandResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 497F7522-82B0-4BD4-84FE-AE8749E4C2F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryBrokerDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBrokerDemandResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBrokerDemandResponseBody) SetCurrentPageNum(v int32) *QueryBrokerDemandResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryBrokerDemandResponseBody) SetData(v []*QueryBrokerDemandResponseBodyData) *QueryBrokerDemandResponseBody {
	s.Data = v
	return s
}

func (s *QueryBrokerDemandResponseBody) SetPageSize(v int32) *QueryBrokerDemandResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryBrokerDemandResponseBody) SetRequestId(v string) *QueryBrokerDemandResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBrokerDemandResponseBody) SetTotalItemNum(v int32) *QueryBrokerDemandResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryBrokerDemandResponseBody) SetTotalPageNum(v int32) *QueryBrokerDemandResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryBrokerDemandResponseBodyData struct {
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// example:
	//
	// 13300000001
	BargainSellerMobile *string `json:"BargainSellerMobile,omitempty" xml:"BargainSellerMobile,omitempty"`
	// example:
	//
	// 100
	BargainSellerPrice *float32 `json:"BargainSellerPrice,omitempty" xml:"BargainSellerPrice,omitempty"`
	// example:
	//
	// SE20183915FI0178
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// taobao.com
	DemandDomain *string `json:"DemandDomain,omitempty" xml:"DemandDomain,omitempty"`
	// example:
	//
	// 1
	DemandPrice *float32 `json:"DemandPrice,omitempty" xml:"DemandPrice,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Email       *string  `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 13300000000
	Mobile        *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	OrderType     *int32  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PartnerDomain *string `json:"PartnerDomain,omitempty" xml:"PartnerDomain,omitempty"`
	// example:
	//
	// test.com
	PayDomain *string `json:"PayDomain,omitempty" xml:"PayDomain,omitempty"`
	// example:
	//
	// 100
	PayPrice *float32 `json:"PayPrice,omitempty" xml:"PayPrice,omitempty"`
	// example:
	//
	// 1524800053000
	PayTime *int64 `json:"PayTime,omitempty" xml:"PayTime,omitempty"`
	// example:
	//
	// 1
	ProduceType *int32 `json:"ProduceType,omitempty" xml:"ProduceType,omitempty"`
	// example:
	//
	// 1524800053000
	PublishTime    *int64 `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	PurchaseStatus *int32 `json:"PurchaseStatus,omitempty" xml:"PurchaseStatus,omitempty"`
	// example:
	//
	// 18800
	ServicePayPrice *float32 `json:"ServicePayPrice,omitempty" xml:"ServicePayPrice,omitempty"`
	SettleBasePrice *float32 `json:"SettleBasePrice,omitempty" xml:"SettleBasePrice,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryBrokerDemandResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryBrokerDemandResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBrokerDemandResponseBodyData) SetAuditStatus(v int32) *QueryBrokerDemandResponseBodyData {
	s.AuditStatus = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetBargainSellerMobile(v string) *QueryBrokerDemandResponseBodyData {
	s.BargainSellerMobile = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetBargainSellerPrice(v float32) *QueryBrokerDemandResponseBodyData {
	s.BargainSellerPrice = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetBizId(v string) *QueryBrokerDemandResponseBodyData {
	s.BizId = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetDemandDomain(v string) *QueryBrokerDemandResponseBodyData {
	s.DemandDomain = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetDemandPrice(v float32) *QueryBrokerDemandResponseBodyData {
	s.DemandPrice = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetDescription(v string) *QueryBrokerDemandResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetEmail(v string) *QueryBrokerDemandResponseBodyData {
	s.Email = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetMobile(v string) *QueryBrokerDemandResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetOrderType(v int32) *QueryBrokerDemandResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetPartnerDomain(v string) *QueryBrokerDemandResponseBodyData {
	s.PartnerDomain = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetPayDomain(v string) *QueryBrokerDemandResponseBodyData {
	s.PayDomain = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetPayPrice(v float32) *QueryBrokerDemandResponseBodyData {
	s.PayPrice = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetPayTime(v int64) *QueryBrokerDemandResponseBodyData {
	s.PayTime = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetProduceType(v int32) *QueryBrokerDemandResponseBodyData {
	s.ProduceType = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetPublishTime(v int64) *QueryBrokerDemandResponseBodyData {
	s.PublishTime = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetPurchaseStatus(v int32) *QueryBrokerDemandResponseBodyData {
	s.PurchaseStatus = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetServicePayPrice(v float32) *QueryBrokerDemandResponseBodyData {
	s.ServicePayPrice = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetSettleBasePrice(v float32) *QueryBrokerDemandResponseBodyData {
	s.SettleBasePrice = &v
	return s
}

func (s *QueryBrokerDemandResponseBodyData) SetStatus(v string) *QueryBrokerDemandResponseBodyData {
	s.Status = &v
	return s
}

type QueryBrokerDemandResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBrokerDemandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBrokerDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBrokerDemandResponse) GoString() string {
	return s.String()
}

func (s *QueryBrokerDemandResponse) SetHeaders(v map[string]*string) *QueryBrokerDemandResponse {
	s.Headers = v
	return s
}

func (s *QueryBrokerDemandResponse) SetStatusCode(v int32) *QueryBrokerDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBrokerDemandResponse) SetBody(v *QueryBrokerDemandResponseBody) *QueryBrokerDemandResponse {
	s.Body = v
	return s
}

type QueryBrokerDemandRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SE20183A0Q7C5556
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryBrokerDemandRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBrokerDemandRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryBrokerDemandRecordRequest) SetBizId(v string) *QueryBrokerDemandRecordRequest {
	s.BizId = &v
	return s
}

func (s *QueryBrokerDemandRecordRequest) SetCurrentPage(v int32) *QueryBrokerDemandRecordRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryBrokerDemandRecordRequest) SetPageSize(v int32) *QueryBrokerDemandRecordRequest {
	s.PageSize = &v
	return s
}

type QueryBrokerDemandRecordResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                                   `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           *QueryBrokerDemandRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 497F7522-82B0-4BD4-84FE-AE8749E4C2F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryBrokerDemandRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryBrokerDemandRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryBrokerDemandRecordResponseBody) SetCurrentPageNum(v int32) *QueryBrokerDemandRecordResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryBrokerDemandRecordResponseBody) SetData(v *QueryBrokerDemandRecordResponseBodyData) *QueryBrokerDemandRecordResponseBody {
	s.Data = v
	return s
}

func (s *QueryBrokerDemandRecordResponseBody) SetPageSize(v int32) *QueryBrokerDemandRecordResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryBrokerDemandRecordResponseBody) SetRequestId(v string) *QueryBrokerDemandRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryBrokerDemandRecordResponseBody) SetTotalItemNum(v int32) *QueryBrokerDemandRecordResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryBrokerDemandRecordResponseBody) SetTotalPageNum(v int32) *QueryBrokerDemandRecordResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryBrokerDemandRecordResponseBodyData struct {
	BrokerDemandRecord []*QueryBrokerDemandRecordResponseBodyDataBrokerDemandRecord `json:"BrokerDemandRecord,omitempty" xml:"BrokerDemandRecord,omitempty" type:"Repeated"`
}

func (s QueryBrokerDemandRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryBrokerDemandRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryBrokerDemandRecordResponseBodyData) SetBrokerDemandRecord(v []*QueryBrokerDemandRecordResponseBodyDataBrokerDemandRecord) *QueryBrokerDemandRecordResponseBodyData {
	s.BrokerDemandRecord = v
	return s
}

type QueryBrokerDemandRecordResponseBodyDataBrokerDemandRecord struct {
	// example:
	//
	// SE20183A0Q7C5556
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 1525249317000
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s QueryBrokerDemandRecordResponseBodyDataBrokerDemandRecord) String() string {
	return tea.Prettify(s)
}

func (s QueryBrokerDemandRecordResponseBodyDataBrokerDemandRecord) GoString() string {
	return s.String()
}

func (s *QueryBrokerDemandRecordResponseBodyDataBrokerDemandRecord) SetBizId(v string) *QueryBrokerDemandRecordResponseBodyDataBrokerDemandRecord {
	s.BizId = &v
	return s
}

func (s *QueryBrokerDemandRecordResponseBodyDataBrokerDemandRecord) SetCreateTime(v int64) *QueryBrokerDemandRecordResponseBodyDataBrokerDemandRecord {
	s.CreateTime = &v
	return s
}

func (s *QueryBrokerDemandRecordResponseBodyDataBrokerDemandRecord) SetDescription(v string) *QueryBrokerDemandRecordResponseBodyDataBrokerDemandRecord {
	s.Description = &v
	return s
}

type QueryBrokerDemandRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBrokerDemandRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBrokerDemandRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBrokerDemandRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryBrokerDemandRecordResponse) SetHeaders(v map[string]*string) *QueryBrokerDemandRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryBrokerDemandRecordResponse) SetStatusCode(v int32) *QueryBrokerDemandRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBrokerDemandRecordResponse) SetBody(v *QueryBrokerDemandRecordResponseBody) *QueryBrokerDemandRecordResponse {
	s.Body = v
	return s
}

type QueryDomainTransferStatusRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s QueryDomainTransferStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainTransferStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainTransferStatusRequest) SetDomainName(v string) *QueryDomainTransferStatusRequest {
	s.DomainName = &v
	return s
}

type QueryDomainTransferStatusResponseBody struct {
	DomainTransferStatus []*QueryDomainTransferStatusResponseBodyDomainTransferStatus `json:"DomainTransferStatus,omitempty" xml:"DomainTransferStatus,omitempty" type:"Repeated"`
	RequestId            *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDomainTransferStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainTransferStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainTransferStatusResponseBody) SetDomainTransferStatus(v []*QueryDomainTransferStatusResponseBodyDomainTransferStatus) *QueryDomainTransferStatusResponseBody {
	s.DomainTransferStatus = v
	return s
}

func (s *QueryDomainTransferStatusResponseBody) SetRequestId(v string) *QueryDomainTransferStatusResponseBody {
	s.RequestId = &v
	return s
}

type QueryDomainTransferStatusResponseBodyDomainTransferStatus struct {
	DomainName              *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatusDescription *string `json:"DomainStatusDescription,omitempty" xml:"DomainStatusDescription,omitempty"`
}

func (s QueryDomainTransferStatusResponseBodyDomainTransferStatus) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainTransferStatusResponseBodyDomainTransferStatus) GoString() string {
	return s.String()
}

func (s *QueryDomainTransferStatusResponseBodyDomainTransferStatus) SetDomainName(v string) *QueryDomainTransferStatusResponseBodyDomainTransferStatus {
	s.DomainName = &v
	return s
}

func (s *QueryDomainTransferStatusResponseBodyDomainTransferStatus) SetDomainStatusDescription(v string) *QueryDomainTransferStatusResponseBodyDomainTransferStatus {
	s.DomainStatusDescription = &v
	return s
}

type QueryDomainTransferStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainTransferStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainTransferStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryDomainTransferStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainTransferStatusResponse) SetHeaders(v map[string]*string) *QueryDomainTransferStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainTransferStatusResponse) SetStatusCode(v int32) *QueryDomainTransferStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainTransferStatusResponse) SetBody(v *QueryDomainTransferStatusResponseBody) *QueryDomainTransferStatusResponse {
	s.Body = v
	return s
}

type QueryPurchasedDomainsRequest struct {
	CurrentPage        *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndOperationTime   *string `json:"EndOperationTime,omitempty" xml:"EndOperationTime,omitempty"`
	OpTimeOrder        *bool   `json:"OpTimeOrder,omitempty" xml:"OpTimeOrder,omitempty"`
	OperationStatus    *int32  `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType        *int32  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	StartOperationTime *string `json:"StartOperationTime,omitempty" xml:"StartOperationTime,omitempty"`
	UpdateTimeOrder    *bool   `json:"UpdateTimeOrder,omitempty" xml:"UpdateTimeOrder,omitempty"`
}

func (s QueryPurchasedDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchasedDomainsRequest) GoString() string {
	return s.String()
}

func (s *QueryPurchasedDomainsRequest) SetCurrentPage(v int32) *QueryPurchasedDomainsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryPurchasedDomainsRequest) SetDomainName(v string) *QueryPurchasedDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *QueryPurchasedDomainsRequest) SetEndOperationTime(v string) *QueryPurchasedDomainsRequest {
	s.EndOperationTime = &v
	return s
}

func (s *QueryPurchasedDomainsRequest) SetOpTimeOrder(v bool) *QueryPurchasedDomainsRequest {
	s.OpTimeOrder = &v
	return s
}

func (s *QueryPurchasedDomainsRequest) SetOperationStatus(v int32) *QueryPurchasedDomainsRequest {
	s.OperationStatus = &v
	return s
}

func (s *QueryPurchasedDomainsRequest) SetPageSize(v int32) *QueryPurchasedDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPurchasedDomainsRequest) SetProductType(v int32) *QueryPurchasedDomainsRequest {
	s.ProductType = &v
	return s
}

func (s *QueryPurchasedDomainsRequest) SetStartOperationTime(v string) *QueryPurchasedDomainsRequest {
	s.StartOperationTime = &v
	return s
}

func (s *QueryPurchasedDomainsRequest) SetUpdateTimeOrder(v bool) *QueryPurchasedDomainsRequest {
	s.UpdateTimeOrder = &v
	return s
}

type QueryPurchasedDomainsResponseBody struct {
	CurrentPageNum *int32                                   `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryPurchasedDomainsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageSize       *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalItemNum   *int32                                   `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                                   `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryPurchasedDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchasedDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPurchasedDomainsResponseBody) SetCurrentPageNum(v int32) *QueryPurchasedDomainsResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryPurchasedDomainsResponseBody) SetData(v []*QueryPurchasedDomainsResponseBodyData) *QueryPurchasedDomainsResponseBody {
	s.Data = v
	return s
}

func (s *QueryPurchasedDomainsResponseBody) SetPageSize(v int32) *QueryPurchasedDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryPurchasedDomainsResponseBody) SetRequestId(v string) *QueryPurchasedDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPurchasedDomainsResponseBody) SetTotalItemNum(v int32) *QueryPurchasedDomainsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryPurchasedDomainsResponseBody) SetTotalPageNum(v int32) *QueryPurchasedDomainsResponseBody {
	s.TotalPageNum = &v
	return s
}

type QueryPurchasedDomainsResponseBodyData struct {
	DeliveryTime    *string  `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	DomainName      *string  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OperationStatus *string  `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	OperationTime   *string  `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	ProductType     *string  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	TradeMoney      *float64 `json:"TradeMoney,omitempty" xml:"TradeMoney,omitempty"`
}

func (s QueryPurchasedDomainsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchasedDomainsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPurchasedDomainsResponseBodyData) SetDeliveryTime(v string) *QueryPurchasedDomainsResponseBodyData {
	s.DeliveryTime = &v
	return s
}

func (s *QueryPurchasedDomainsResponseBodyData) SetDomainName(v string) *QueryPurchasedDomainsResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *QueryPurchasedDomainsResponseBodyData) SetOperationStatus(v string) *QueryPurchasedDomainsResponseBodyData {
	s.OperationStatus = &v
	return s
}

func (s *QueryPurchasedDomainsResponseBodyData) SetOperationTime(v string) *QueryPurchasedDomainsResponseBodyData {
	s.OperationTime = &v
	return s
}

func (s *QueryPurchasedDomainsResponseBodyData) SetProductType(v string) *QueryPurchasedDomainsResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *QueryPurchasedDomainsResponseBodyData) SetTradeMoney(v float64) *QueryPurchasedDomainsResponseBodyData {
	s.TradeMoney = &v
	return s
}

type QueryPurchasedDomainsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPurchasedDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPurchasedDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPurchasedDomainsResponse) GoString() string {
	return s.String()
}

func (s *QueryPurchasedDomainsResponse) SetHeaders(v map[string]*string) *QueryPurchasedDomainsResponse {
	s.Headers = v
	return s
}

func (s *QueryPurchasedDomainsResponse) SetStatusCode(v int32) *QueryPurchasedDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPurchasedDomainsResponse) SetBody(v *QueryPurchasedDomainsResponseBody) *QueryPurchasedDomainsResponse {
	s.Body = v
	return s
}

type RecordDemandRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SE20183A0Q7C5556
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// some message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RecordDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s RecordDemandRequest) GoString() string {
	return s.String()
}

func (s *RecordDemandRequest) SetBizId(v string) *RecordDemandRequest {
	s.BizId = &v
	return s
}

func (s *RecordDemandRequest) SetMessage(v string) *RecordDemandRequest {
	s.Message = &v
	return s
}

type RecordDemandResponseBody struct {
	// example:
	//
	// 497F7522-82B0-4BD4-84FE-AE8749E4C2F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecordDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecordDemandResponseBody) GoString() string {
	return s.String()
}

func (s *RecordDemandResponseBody) SetRequestId(v string) *RecordDemandResponseBody {
	s.RequestId = &v
	return s
}

type RecordDemandResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecordDemandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecordDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s RecordDemandResponse) GoString() string {
	return s.String()
}

func (s *RecordDemandResponse) SetHeaders(v map[string]*string) *RecordDemandResponse {
	s.Headers = v
	return s
}

func (s *RecordDemandResponse) SetStatusCode(v int32) *RecordDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *RecordDemandResponse) SetBody(v *RecordDemandResponseBody) *RecordDemandResponse {
	s.Body = v
	return s
}

type RefuseDemandRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SE20183A0Q7C5556
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// some message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RefuseDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s RefuseDemandRequest) GoString() string {
	return s.String()
}

func (s *RefuseDemandRequest) SetBizId(v string) *RefuseDemandRequest {
	s.BizId = &v
	return s
}

func (s *RefuseDemandRequest) SetMessage(v string) *RefuseDemandRequest {
	s.Message = &v
	return s
}

type RefuseDemandResponseBody struct {
	// example:
	//
	// 497F7522-82B0-4BD4-84FE-AE8749E4C2F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefuseDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefuseDemandResponseBody) GoString() string {
	return s.String()
}

func (s *RefuseDemandResponseBody) SetRequestId(v string) *RefuseDemandResponseBody {
	s.RequestId = &v
	return s
}

type RefuseDemandResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefuseDemandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefuseDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s RefuseDemandResponse) GoString() string {
	return s.String()
}

func (s *RefuseDemandResponse) SetHeaders(v map[string]*string) *RefuseDemandResponse {
	s.Headers = v
	return s
}

func (s *RefuseDemandResponse) SetStatusCode(v int32) *RefuseDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *RefuseDemandResponse) SetBody(v *RefuseDemandResponseBody) *RefuseDemandResponse {
	s.Body = v
	return s
}

type RequestPayDemandRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SE20183A0Q7C5556
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// some message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	Price *float32 `json:"Price,omitempty" xml:"Price,omitempty"`
	// example:
	//
	// 1
	ProduceType *int32 `json:"ProduceType,omitempty" xml:"ProduceType,omitempty"`
}

func (s RequestPayDemandRequest) String() string {
	return tea.Prettify(s)
}

func (s RequestPayDemandRequest) GoString() string {
	return s.String()
}

func (s *RequestPayDemandRequest) SetBizId(v string) *RequestPayDemandRequest {
	s.BizId = &v
	return s
}

func (s *RequestPayDemandRequest) SetDomainName(v string) *RequestPayDemandRequest {
	s.DomainName = &v
	return s
}

func (s *RequestPayDemandRequest) SetMessage(v string) *RequestPayDemandRequest {
	s.Message = &v
	return s
}

func (s *RequestPayDemandRequest) SetPrice(v float32) *RequestPayDemandRequest {
	s.Price = &v
	return s
}

func (s *RequestPayDemandRequest) SetProduceType(v int32) *RequestPayDemandRequest {
	s.ProduceType = &v
	return s
}

type RequestPayDemandResponseBody struct {
	// example:
	//
	// 497F7522-82B0-4BD4-84FE-AE8749E4C2F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RequestPayDemandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RequestPayDemandResponseBody) GoString() string {
	return s.String()
}

func (s *RequestPayDemandResponseBody) SetRequestId(v string) *RequestPayDemandResponseBody {
	s.RequestId = &v
	return s
}

type RequestPayDemandResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RequestPayDemandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RequestPayDemandResponse) String() string {
	return tea.Prettify(s)
}

func (s RequestPayDemandResponse) GoString() string {
	return s.String()
}

func (s *RequestPayDemandResponse) SetHeaders(v map[string]*string) *RequestPayDemandResponse {
	s.Headers = v
	return s
}

func (s *RequestPayDemandResponse) SetStatusCode(v int32) *RequestPayDemandResponse {
	s.StatusCode = &v
	return s
}

func (s *RequestPayDemandResponse) SetBody(v *RequestPayDemandResponseBody) *RequestPayDemandResponse {
	s.Body = v
	return s
}

type ReserveDomainRequest struct {
	// example:
	//
	// 4
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// aliyun.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s ReserveDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s ReserveDomainRequest) GoString() string {
	return s.String()
}

func (s *ReserveDomainRequest) SetChannels(v []*string) *ReserveDomainRequest {
	s.Channels = v
	return s
}

func (s *ReserveDomainRequest) SetDomainName(v string) *ReserveDomainRequest {
	s.DomainName = &v
	return s
}

type ReserveDomainResponseBody struct {
	// example:
	//
	// 12080761
	AuctionId *string `json:"AuctionId,omitempty" xml:"AuctionId,omitempty"`
	// example:
	//
	// 64F63E07-3AF6-4D59-8616-55DF1A9E03ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReserveDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReserveDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ReserveDomainResponseBody) SetAuctionId(v string) *ReserveDomainResponseBody {
	s.AuctionId = &v
	return s
}

func (s *ReserveDomainResponseBody) SetRequestId(v string) *ReserveDomainResponseBody {
	s.RequestId = &v
	return s
}

type ReserveDomainResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReserveDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReserveDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s ReserveDomainResponse) GoString() string {
	return s.String()
}

func (s *ReserveDomainResponse) SetHeaders(v map[string]*string) *ReserveDomainResponse {
	s.Headers = v
	return s
}

func (s *ReserveDomainResponse) SetStatusCode(v int32) *ReserveDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ReserveDomainResponse) SetBody(v *ReserveDomainResponseBody) *ReserveDomainResponse {
	s.Body = v
	return s
}

type ReserveIntlDomainRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s ReserveIntlDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s ReserveIntlDomainRequest) GoString() string {
	return s.String()
}

func (s *ReserveIntlDomainRequest) SetDomainName(v string) *ReserveIntlDomainRequest {
	s.DomainName = &v
	return s
}

type ReserveIntlDomainResponseBody struct {
	AllowRetry     *bool         `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	AppName        *string       `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AuctionId      *string       `json:"AuctionId,omitempty" xml:"AuctionId,omitempty"`
	DynamicCode    *string       `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	ErrorCode      *string       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg       *string       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	HttpStatusCode *int32        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReserveIntlDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReserveIntlDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ReserveIntlDomainResponseBody) SetAllowRetry(v bool) *ReserveIntlDomainResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ReserveIntlDomainResponseBody) SetAppName(v string) *ReserveIntlDomainResponseBody {
	s.AppName = &v
	return s
}

func (s *ReserveIntlDomainResponseBody) SetAuctionId(v string) *ReserveIntlDomainResponseBody {
	s.AuctionId = &v
	return s
}

func (s *ReserveIntlDomainResponseBody) SetDynamicCode(v string) *ReserveIntlDomainResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ReserveIntlDomainResponseBody) SetDynamicMessage(v string) *ReserveIntlDomainResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ReserveIntlDomainResponseBody) SetErrorArgs(v []interface{}) *ReserveIntlDomainResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ReserveIntlDomainResponseBody) SetErrorCode(v string) *ReserveIntlDomainResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReserveIntlDomainResponseBody) SetErrorMsg(v string) *ReserveIntlDomainResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ReserveIntlDomainResponseBody) SetHttpStatusCode(v int32) *ReserveIntlDomainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReserveIntlDomainResponseBody) SetRequestId(v string) *ReserveIntlDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReserveIntlDomainResponseBody) SetSuccess(v bool) *ReserveIntlDomainResponseBody {
	s.Success = &v
	return s
}

type ReserveIntlDomainResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReserveIntlDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReserveIntlDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s ReserveIntlDomainResponse) GoString() string {
	return s.String()
}

func (s *ReserveIntlDomainResponse) SetHeaders(v map[string]*string) *ReserveIntlDomainResponse {
	s.Headers = v
	return s
}

func (s *ReserveIntlDomainResponse) SetStatusCode(v int32) *ReserveIntlDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ReserveIntlDomainResponse) SetBody(v *ReserveIntlDomainResponseBody) *ReserveIntlDomainResponse {
	s.Body = v
	return s
}

type SelectedDomainListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20231109
	ListDate *string `json:"ListDate,omitempty" xml:"ListDate,omitempty"`
}

func (s SelectedDomainListRequest) String() string {
	return tea.Prettify(s)
}

func (s SelectedDomainListRequest) GoString() string {
	return s.String()
}

func (s *SelectedDomainListRequest) SetListDate(v string) *SelectedDomainListRequest {
	s.ListDate = &v
	return s
}

type SelectedDomainListResponseBody struct {
	// example:
	//
	// OssFileNotFound
	ErrorCode *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Module    *SelectedDomainListResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 80011ABC-F573-4795-B0E8-377BFBBA3422
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SelectedDomainListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SelectedDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *SelectedDomainListResponseBody) SetErrorCode(v string) *SelectedDomainListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SelectedDomainListResponseBody) SetModule(v *SelectedDomainListResponseBodyModule) *SelectedDomainListResponseBody {
	s.Module = v
	return s
}

func (s *SelectedDomainListResponseBody) SetRequestId(v string) *SelectedDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SelectedDomainListResponseBody) SetSuccess(v bool) *SelectedDomainListResponseBody {
	s.Success = &v
	return s
}

type SelectedDomainListResponseBodyModule struct {
	// example:
	//
	// http://selected-domain.oss-cn-zhangjiakou.aliyuncs.com/aliyun_selected_domain_20231109.gz?Expires=1699524493&OSSAccessKeyId=LTAI5tPMAybR4gfSEjdfAk1F&Signature=2Tpo7Eaf%2BqIop8SuMtI91m%2FAFpY%3D
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
}

func (s SelectedDomainListResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s SelectedDomainListResponseBodyModule) GoString() string {
	return s.String()
}

func (s *SelectedDomainListResponseBodyModule) SetDownloadUrl(v string) *SelectedDomainListResponseBodyModule {
	s.DownloadUrl = &v
	return s
}

type SelectedDomainListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SelectedDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SelectedDomainListResponse) String() string {
	return tea.Prettify(s)
}

func (s SelectedDomainListResponse) GoString() string {
	return s.String()
}

func (s *SelectedDomainListResponse) SetHeaders(v map[string]*string) *SelectedDomainListResponse {
	s.Headers = v
	return s
}

func (s *SelectedDomainListResponse) SetStatusCode(v int32) *SelectedDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *SelectedDomainListResponse) SetBody(v *SelectedDomainListResponseBody) *SelectedDomainListResponse {
	s.Body = v
	return s
}

type SubmitPurchaseInfoRequest struct {
	BizId            *string   `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PurchaseCurrency *string   `json:"PurchaseCurrency,omitempty" xml:"PurchaseCurrency,omitempty"`
	PurchasePrice    *float64  `json:"PurchasePrice,omitempty" xml:"PurchasePrice,omitempty"`
	PurchaseProofs   []*string `json:"PurchaseProofs,omitempty" xml:"PurchaseProofs,omitempty" type:"Repeated"`
}

func (s SubmitPurchaseInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitPurchaseInfoRequest) GoString() string {
	return s.String()
}

func (s *SubmitPurchaseInfoRequest) SetBizId(v string) *SubmitPurchaseInfoRequest {
	s.BizId = &v
	return s
}

func (s *SubmitPurchaseInfoRequest) SetPurchaseCurrency(v string) *SubmitPurchaseInfoRequest {
	s.PurchaseCurrency = &v
	return s
}

func (s *SubmitPurchaseInfoRequest) SetPurchasePrice(v float64) *SubmitPurchaseInfoRequest {
	s.PurchasePrice = &v
	return s
}

func (s *SubmitPurchaseInfoRequest) SetPurchaseProofs(v []*string) *SubmitPurchaseInfoRequest {
	s.PurchaseProofs = v
	return s
}

type SubmitPurchaseInfoResponseBody struct {
	AllowRetry     *bool         `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	AppName        *string       `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DynamicCode    *string       `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	ErrorCode      *string       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg       *string       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	HttpStatusCode *int32        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Module         interface{}   `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId      *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitPurchaseInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitPurchaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPurchaseInfoResponseBody) SetAllowRetry(v bool) *SubmitPurchaseInfoResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *SubmitPurchaseInfoResponseBody) SetAppName(v string) *SubmitPurchaseInfoResponseBody {
	s.AppName = &v
	return s
}

func (s *SubmitPurchaseInfoResponseBody) SetDynamicCode(v string) *SubmitPurchaseInfoResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SubmitPurchaseInfoResponseBody) SetDynamicMessage(v string) *SubmitPurchaseInfoResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SubmitPurchaseInfoResponseBody) SetErrorArgs(v []interface{}) *SubmitPurchaseInfoResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *SubmitPurchaseInfoResponseBody) SetErrorCode(v string) *SubmitPurchaseInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitPurchaseInfoResponseBody) SetErrorMsg(v string) *SubmitPurchaseInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SubmitPurchaseInfoResponseBody) SetHttpStatusCode(v int32) *SubmitPurchaseInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitPurchaseInfoResponseBody) SetModule(v interface{}) *SubmitPurchaseInfoResponseBody {
	s.Module = v
	return s
}

func (s *SubmitPurchaseInfoResponseBody) SetRequestId(v string) *SubmitPurchaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitPurchaseInfoResponseBody) SetSuccess(v bool) *SubmitPurchaseInfoResponseBody {
	s.Success = &v
	return s
}

type SubmitPurchaseInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitPurchaseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitPurchaseInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitPurchaseInfoResponse) GoString() string {
	return s.String()
}

func (s *SubmitPurchaseInfoResponse) SetHeaders(v map[string]*string) *SubmitPurchaseInfoResponse {
	s.Headers = v
	return s
}

func (s *SubmitPurchaseInfoResponse) SetStatusCode(v int32) *SubmitPurchaseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitPurchaseInfoResponse) SetBody(v *SubmitPurchaseInfoResponseBody) *SubmitPurchaseInfoResponse {
	s.Body = v
	return s
}

type UpdatePartnerReservePriceRequest struct {
	// This parameter is required.
	BiddingId *int32 `json:"BiddingId,omitempty" xml:"BiddingId,omitempty"`
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	PartnerType *string `json:"PartnerType,omitempty" xml:"PartnerType,omitempty"`
	// This parameter is required.
	ReservePrice *float64 `json:"ReservePrice,omitempty" xml:"ReservePrice,omitempty"`
}

func (s UpdatePartnerReservePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePartnerReservePriceRequest) GoString() string {
	return s.String()
}

func (s *UpdatePartnerReservePriceRequest) SetBiddingId(v int32) *UpdatePartnerReservePriceRequest {
	s.BiddingId = &v
	return s
}

func (s *UpdatePartnerReservePriceRequest) SetDomainName(v string) *UpdatePartnerReservePriceRequest {
	s.DomainName = &v
	return s
}

func (s *UpdatePartnerReservePriceRequest) SetPartnerType(v string) *UpdatePartnerReservePriceRequest {
	s.PartnerType = &v
	return s
}

func (s *UpdatePartnerReservePriceRequest) SetReservePrice(v float64) *UpdatePartnerReservePriceRequest {
	s.ReservePrice = &v
	return s
}

type UpdatePartnerReservePriceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePartnerReservePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePartnerReservePriceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePartnerReservePriceResponseBody) SetRequestId(v string) *UpdatePartnerReservePriceResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePartnerReservePriceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePartnerReservePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePartnerReservePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePartnerReservePriceResponse) GoString() string {
	return s.String()
}

func (s *UpdatePartnerReservePriceResponse) SetHeaders(v map[string]*string) *UpdatePartnerReservePriceResponse {
	s.Headers = v
	return s
}

func (s *UpdatePartnerReservePriceResponse) SetStatusCode(v int32) *UpdatePartnerReservePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePartnerReservePriceResponse) SetBody(v *UpdatePartnerReservePriceResponseBody) *UpdatePartnerReservePriceResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("domain"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AcceptDemandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptDemandResponse
func (client *Client) AcceptDemandWithOptions(request *AcceptDemandRequest, runtime *util.RuntimeOptions) (_result *AcceptDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AcceptDemand"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AcceptDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AcceptDemandRequest
//
// @return AcceptDemandResponse
func (client *Client) AcceptDemand(request *AcceptDemandRequest) (_result *AcceptDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AcceptDemandResponse{}
	_body, _err := client.AcceptDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BidDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BidDomainResponse
func (client *Client) BidDomainWithOptions(request *BidDomainRequest, runtime *util.RuntimeOptions) (_result *BidDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuctionId)) {
		body["AuctionId"] = request.AuctionId
	}

	if !tea.BoolValue(util.IsUnset(request.Currency)) {
		body["Currency"] = request.Currency
	}

	if !tea.BoolValue(util.IsUnset(request.MaxBid)) {
		body["MaxBid"] = request.MaxBid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BidDomain"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BidDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - BidDomainRequest
//
// @return BidDomainResponse
func (client *Client) BidDomain(request *BidDomainRequest) (_result *BidDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BidDomainResponse{}
	_body, _err := client.BidDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ChangeAuctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeAuctionResponse
func (client *Client) ChangeAuctionWithOptions(request *ChangeAuctionRequest, runtime *util.RuntimeOptions) (_result *ChangeAuctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuctionList)) {
		body["AuctionList"] = request.AuctionList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeAuction"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeAuctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ChangeAuctionRequest
//
// @return ChangeAuctionResponse
func (client *Client) ChangeAuction(request *ChangeAuctionRequest) (_result *ChangeAuctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeAuctionResponse{}
	_body, _err := client.ChangeAuctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验域名在售状态
//
// @param request - CheckDomainStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDomainStatusResponse
func (client *Client) CheckDomainStatusWithOptions(request *CheckDomainStatusRequest, runtime *util.RuntimeOptions) (_result *CheckDomainStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckDomainStatus"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDomainStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验域名在售状态
//
// @param request - CheckDomainStatusRequest
//
// @return CheckDomainStatusResponse
func (client *Client) CheckDomainStatus(request *CheckDomainStatusRequest) (_result *CheckDomainStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDomainStatusResponse{}
	_body, _err := client.CheckDomainStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 一口价严选询价接口
//
// @param request - CheckSelectedDomainStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSelectedDomainStatusResponse
func (client *Client) CheckSelectedDomainStatusWithOptions(request *CheckSelectedDomainStatusRequest, runtime *util.RuntimeOptions) (_result *CheckSelectedDomainStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckSelectedDomainStatus"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckSelectedDomainStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 一口价严选询价接口
//
// @param request - CheckSelectedDomainStatusRequest
//
// @return CheckSelectedDomainStatusResponse
func (client *Client) CheckSelectedDomainStatus(request *CheckSelectedDomainStatusRequest) (_result *CheckSelectedDomainStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckSelectedDomainStatusResponse{}
	_body, _err := client.CheckSelectedDomainStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一口价需求单
//
// @param request - CreateFixedPriceDemandOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFixedPriceDemandOrderResponse
func (client *Client) CreateFixedPriceDemandOrderWithOptions(request *CreateFixedPriceDemandOrderRequest, runtime *util.RuntimeOptions) (_result *CreateFixedPriceDemandOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFixedPriceDemandOrder"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFixedPriceDemandOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一口价需求单
//
// @param request - CreateFixedPriceDemandOrderRequest
//
// @return CreateFixedPriceDemandOrderResponse
func (client *Client) CreateFixedPriceDemandOrder(request *CreateFixedPriceDemandOrderRequest) (_result *CreateFixedPriceDemandOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFixedPriceDemandOrderResponse{}
	_body, _err := client.CreateFixedPriceDemandOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 一口价严选下单购买接口，阿里云账户余额直接扣费
//
// @param request - CreateFixedPriceSelectedOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFixedPriceSelectedOrderResponse
func (client *Client) CreateFixedPriceSelectedOrderWithOptions(request *CreateFixedPriceSelectedOrderRequest, runtime *util.RuntimeOptions) (_result *CreateFixedPriceSelectedOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.ExpectedPrice)) {
		query["ExpectedPrice"] = request.ExpectedPrice
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFixedPriceSelectedOrder"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFixedPriceSelectedOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 一口价严选下单购买接口，阿里云账户余额直接扣费
//
// @param request - CreateFixedPriceSelectedOrderRequest
//
// @return CreateFixedPriceSelectedOrderResponse
func (client *Client) CreateFixedPriceSelectedOrder(request *CreateFixedPriceSelectedOrderRequest) (_result *CreateFixedPriceSelectedOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFixedPriceSelectedOrderResponse{}
	_body, _err := client.CreateFixedPriceSelectedOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - FailDemandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FailDemandResponse
func (client *Client) FailDemandWithOptions(request *FailDemandRequest, runtime *util.RuntimeOptions) (_result *FailDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FailDemand"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FailDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - FailDemandRequest
//
// @return FailDemandResponse
func (client *Client) FailDemand(request *FailDemandRequest) (_result *FailDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FailDemandResponse{}
	_body, _err := client.FailDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - FinishDemandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishDemandResponse
func (client *Client) FinishDemandWithOptions(request *FinishDemandRequest, runtime *util.RuntimeOptions) (_result *FinishDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FinishDemand"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FinishDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - FinishDemandRequest
//
// @return FinishDemandResponse
func (client *Client) FinishDemand(request *FinishDemandRequest) (_result *FinishDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FinishDemandResponse{}
	_body, _err := client.FinishDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetIntlDomainDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntlDomainDownloadUrlResponse
func (client *Client) GetIntlDomainDownloadUrlWithOptions(runtime *util.RuntimeOptions) (_result *GetIntlDomainDownloadUrlResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetIntlDomainDownloadUrl"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIntlDomainDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return GetIntlDomainDownloadUrlResponse
func (client *Client) GetIntlDomainDownloadUrl() (_result *GetIntlDomainDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIntlDomainDownloadUrlResponse{}
	_body, _err := client.GetIntlDomainDownloadUrlWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetReserveDomainUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReserveDomainUrlResponse
func (client *Client) GetReserveDomainUrlWithOptions(runtime *util.RuntimeOptions) (_result *GetReserveDomainUrlResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetReserveDomainUrl"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetReserveDomainUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return GetReserveDomainUrlResponse
func (client *Client) GetReserveDomainUrl() (_result *GetReserveDomainUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetReserveDomainUrlResponse{}
	_body, _err := client.GetReserveDomainUrlWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 购买国际站预释放域名
//
// @param request - PurchaseIntlDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PurchaseIntlDomainResponse
func (client *Client) PurchaseIntlDomainWithOptions(request *PurchaseIntlDomainRequest, runtime *util.RuntimeOptions) (_result *PurchaseIntlDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuctionId)) {
		body["AuctionId"] = request.AuctionId
	}

	if !tea.BoolValue(util.IsUnset(request.Currency)) {
		body["Currency"] = request.Currency
	}

	if !tea.BoolValue(util.IsUnset(request.Price)) {
		body["Price"] = request.Price
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PurchaseIntlDomain"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PurchaseIntlDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 购买国际站预释放域名
//
// @param request - PurchaseIntlDomainRequest
//
// @return PurchaseIntlDomainResponse
func (client *Client) PurchaseIntlDomain(request *PurchaseIntlDomainRequest) (_result *PurchaseIntlDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PurchaseIntlDomainResponse{}
	_body, _err := client.PurchaseIntlDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAuctionDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAuctionDetailResponse
func (client *Client) QueryAuctionDetailWithOptions(request *QueryAuctionDetailRequest, runtime *util.RuntimeOptions) (_result *QueryAuctionDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuctionId)) {
		body["AuctionId"] = request.AuctionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAuctionDetail"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAuctionDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAuctionDetailRequest
//
// @return QueryAuctionDetailResponse
func (client *Client) QueryAuctionDetail(request *QueryAuctionDetailRequest) (_result *QueryAuctionDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAuctionDetailResponse{}
	_body, _err := client.QueryAuctionDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAuctionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAuctionsResponse
func (client *Client) QueryAuctionsWithOptions(request *QueryAuctionsRequest, runtime *util.RuntimeOptions) (_result *QueryAuctionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuctionEndTimeOrder)) {
		body["AuctionEndTimeOrder"] = request.AuctionEndTimeOrder
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Statuses)) {
		body["Statuses"] = request.Statuses
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAuctions"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAuctionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAuctionsRequest
//
// @return QueryAuctionsResponse
func (client *Client) QueryAuctions(request *QueryAuctionsRequest) (_result *QueryAuctionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAuctionsResponse{}
	_body, _err := client.QueryAuctionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryBidRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBidRecordsResponse
func (client *Client) QueryBidRecordsWithOptions(request *QueryBidRecordsRequest, runtime *util.RuntimeOptions) (_result *QueryBidRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuctionId)) {
		body["AuctionId"] = request.AuctionId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBidRecords"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBidRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryBidRecordsRequest
//
// @return QueryBidRecordsResponse
func (client *Client) QueryBidRecords(request *QueryBidRecordsRequest) (_result *QueryBidRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBidRecordsResponse{}
	_body, _err := client.QueryBidRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryBookingDomainInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBookingDomainInfoResponse
func (client *Client) QueryBookingDomainInfoWithOptions(request *QueryBookingDomainInfoRequest, runtime *util.RuntimeOptions) (_result *QueryBookingDomainInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBookingDomainInfo"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBookingDomainInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryBookingDomainInfoRequest
//
// @return QueryBookingDomainInfoResponse
func (client *Client) QueryBookingDomainInfo(request *QueryBookingDomainInfoRequest) (_result *QueryBookingDomainInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBookingDomainInfoResponse{}
	_body, _err := client.QueryBookingDomainInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询回购订单列表
//
// @param request - QueryBrokerDemandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBrokerDemandResponse
func (client *Client) QueryBrokerDemandWithOptions(request *QueryBrokerDemandRequest, runtime *util.RuntimeOptions) (_result *QueryBrokerDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBrokerDemand"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBrokerDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询回购订单列表
//
// @param request - QueryBrokerDemandRequest
//
// @return QueryBrokerDemandResponse
func (client *Client) QueryBrokerDemand(request *QueryBrokerDemandRequest) (_result *QueryBrokerDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBrokerDemandResponse{}
	_body, _err := client.QueryBrokerDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryBrokerDemandRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBrokerDemandRecordResponse
func (client *Client) QueryBrokerDemandRecordWithOptions(request *QueryBrokerDemandRecordRequest, runtime *util.RuntimeOptions) (_result *QueryBrokerDemandRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryBrokerDemandRecord"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryBrokerDemandRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryBrokerDemandRecordRequest
//
// @return QueryBrokerDemandRecordResponse
func (client *Client) QueryBrokerDemandRecord(request *QueryBrokerDemandRecordRequest) (_result *QueryBrokerDemandRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryBrokerDemandRecordResponse{}
	_body, _err := client.QueryBrokerDemandRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryDomainTransferStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainTransferStatusResponse
func (client *Client) QueryDomainTransferStatusWithOptions(request *QueryDomainTransferStatusRequest, runtime *util.RuntimeOptions) (_result *QueryDomainTransferStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryDomainTransferStatus"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryDomainTransferStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryDomainTransferStatusRequest
//
// @return QueryDomainTransferStatusResponse
func (client *Client) QueryDomainTransferStatus(request *QueryDomainTransferStatusRequest) (_result *QueryDomainTransferStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryDomainTransferStatusResponse{}
	_body, _err := client.QueryDomainTransferStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryPurchasedDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPurchasedDomainsResponse
func (client *Client) QueryPurchasedDomainsWithOptions(request *QueryPurchasedDomainsRequest, runtime *util.RuntimeOptions) (_result *QueryPurchasedDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.EndOperationTime)) {
		body["EndOperationTime"] = request.EndOperationTime
	}

	if !tea.BoolValue(util.IsUnset(request.OpTimeOrder)) {
		body["OpTimeOrder"] = request.OpTimeOrder
	}

	if !tea.BoolValue(util.IsUnset(request.OperationStatus)) {
		body["OperationStatus"] = request.OperationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.StartOperationTime)) {
		body["StartOperationTime"] = request.StartOperationTime
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateTimeOrder)) {
		body["UpdateTimeOrder"] = request.UpdateTimeOrder
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPurchasedDomains"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPurchasedDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryPurchasedDomainsRequest
//
// @return QueryPurchasedDomainsResponse
func (client *Client) QueryPurchasedDomains(request *QueryPurchasedDomainsRequest) (_result *QueryPurchasedDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPurchasedDomainsResponse{}
	_body, _err := client.QueryPurchasedDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RecordDemandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecordDemandResponse
func (client *Client) RecordDemandWithOptions(request *RecordDemandRequest, runtime *util.RuntimeOptions) (_result *RecordDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecordDemand"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecordDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RecordDemandRequest
//
// @return RecordDemandResponse
func (client *Client) RecordDemand(request *RecordDemandRequest) (_result *RecordDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecordDemandResponse{}
	_body, _err := client.RecordDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RefuseDemandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefuseDemandResponse
func (client *Client) RefuseDemandWithOptions(request *RefuseDemandRequest, runtime *util.RuntimeOptions) (_result *RefuseDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefuseDemand"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefuseDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RefuseDemandRequest
//
// @return RefuseDemandResponse
func (client *Client) RefuseDemand(request *RefuseDemandRequest) (_result *RefuseDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefuseDemandResponse{}
	_body, _err := client.RefuseDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RequestPayDemandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RequestPayDemandResponse
func (client *Client) RequestPayDemandWithOptions(request *RequestPayDemandRequest, runtime *util.RuntimeOptions) (_result *RequestPayDemandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.Price)) {
		query["Price"] = request.Price
	}

	if !tea.BoolValue(util.IsUnset(request.ProduceType)) {
		query["ProduceType"] = request.ProduceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RequestPayDemand"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RequestPayDemandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RequestPayDemandRequest
//
// @return RequestPayDemandResponse
func (client *Client) RequestPayDemand(request *RequestPayDemandRequest) (_result *RequestPayDemandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RequestPayDemandResponse{}
	_body, _err := client.RequestPayDemandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReserveDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReserveDomainResponse
func (client *Client) ReserveDomainWithOptions(request *ReserveDomainRequest, runtime *util.RuntimeOptions) (_result *ReserveDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channels)) {
		body["Channels"] = request.Channels
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReserveDomain"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReserveDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReserveDomainRequest
//
// @return ReserveDomainResponse
func (client *Client) ReserveDomain(request *ReserveDomainRequest) (_result *ReserveDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReserveDomainResponse{}
	_body, _err := client.ReserveDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReserveIntlDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReserveIntlDomainResponse
func (client *Client) ReserveIntlDomainWithOptions(request *ReserveIntlDomainRequest, runtime *util.RuntimeOptions) (_result *ReserveIntlDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReserveIntlDomain"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReserveIntlDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReserveIntlDomainRequest
//
// @return ReserveIntlDomainResponse
func (client *Client) ReserveIntlDomain(request *ReserveIntlDomainRequest) (_result *ReserveIntlDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReserveIntlDomainResponse{}
	_body, _err := client.ReserveIntlDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 严选列表导出，明日凌晨2点前生成文件，导出凌晨1点前所有在售严选域名
//
// @param request - SelectedDomainListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectedDomainListResponse
func (client *Client) SelectedDomainListWithOptions(request *SelectedDomainListRequest, runtime *util.RuntimeOptions) (_result *SelectedDomainListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListDate)) {
		query["ListDate"] = request.ListDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SelectedDomainList"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SelectedDomainListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 严选列表导出，明日凌晨2点前生成文件，导出凌晨1点前所有在售严选域名
//
// @param request - SelectedDomainListRequest
//
// @return SelectedDomainListResponse
func (client *Client) SelectedDomainList(request *SelectedDomainListRequest) (_result *SelectedDomainListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SelectedDomainListResponse{}
	_body, _err := client.SelectedDomainListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交采购信息
//
// @param request - SubmitPurchaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitPurchaseInfoResponse
func (client *Client) SubmitPurchaseInfoWithOptions(request *SubmitPurchaseInfoRequest, runtime *util.RuntimeOptions) (_result *SubmitPurchaseInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaseCurrency)) {
		body["PurchaseCurrency"] = request.PurchaseCurrency
	}

	if !tea.BoolValue(util.IsUnset(request.PurchasePrice)) {
		body["PurchasePrice"] = request.PurchasePrice
	}

	if !tea.BoolValue(util.IsUnset(request.PurchaseProofs)) {
		body["PurchaseProofs"] = request.PurchaseProofs
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitPurchaseInfo"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitPurchaseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交采购信息
//
// @param request - SubmitPurchaseInfoRequest
//
// @return SubmitPurchaseInfoResponse
func (client *Client) SubmitPurchaseInfo(request *SubmitPurchaseInfoRequest) (_result *SubmitPurchaseInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitPurchaseInfoResponse{}
	_body, _err := client.SubmitPurchaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合作方同步报价
//
// @param request - UpdatePartnerReservePriceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePartnerReservePriceResponse
func (client *Client) UpdatePartnerReservePriceWithOptions(request *UpdatePartnerReservePriceRequest, runtime *util.RuntimeOptions) (_result *UpdatePartnerReservePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BiddingId)) {
		body["BiddingId"] = request.BiddingId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.PartnerType)) {
		body["PartnerType"] = request.PartnerType
	}

	if !tea.BoolValue(util.IsUnset(request.ReservePrice)) {
		body["ReservePrice"] = request.ReservePrice
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePartnerReservePrice"),
		Version:     tea.String("2018-02-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePartnerReservePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作方同步报价
//
// @param request - UpdatePartnerReservePriceRequest
//
// @return UpdatePartnerReservePriceResponse
func (client *Client) UpdatePartnerReservePrice(request *UpdatePartnerReservePriceRequest) (_result *UpdatePartnerReservePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePartnerReservePriceResponse{}
	_body, _err := client.UpdatePartnerReservePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
