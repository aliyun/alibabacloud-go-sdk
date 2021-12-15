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

type AdaptGameVersionRequest struct {
	// 帧率
	FrameRate *string `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// 分辨率
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// 游戏版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s AdaptGameVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s AdaptGameVersionRequest) GoString() string {
	return s.String()
}

func (s *AdaptGameVersionRequest) SetFrameRate(v string) *AdaptGameVersionRequest {
	s.FrameRate = &v
	return s
}

func (s *AdaptGameVersionRequest) SetResolution(v string) *AdaptGameVersionRequest {
	s.Resolution = &v
	return s
}

func (s *AdaptGameVersionRequest) SetVersionId(v string) *AdaptGameVersionRequest {
	s.VersionId = &v
	return s
}

type AdaptGameVersionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Id of the task
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AdaptGameVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AdaptGameVersionResponseBody) GoString() string {
	return s.String()
}

func (s *AdaptGameVersionResponseBody) SetRequestId(v string) *AdaptGameVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AdaptGameVersionResponseBody) SetTaskId(v string) *AdaptGameVersionResponseBody {
	s.TaskId = &v
	return s
}

type AdaptGameVersionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AdaptGameVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AdaptGameVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s AdaptGameVersionResponse) GoString() string {
	return s.String()
}

func (s *AdaptGameVersionResponse) SetHeaders(v map[string]*string) *AdaptGameVersionResponse {
	s.Headers = v
	return s
}

func (s *AdaptGameVersionResponse) SetBody(v *AdaptGameVersionResponseBody) *AdaptGameVersionResponse {
	s.Body = v
	return s
}

type AddGameToProjectRequest struct {
	// 游戏iD
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AddGameToProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGameToProjectRequest) GoString() string {
	return s.String()
}

func (s *AddGameToProjectRequest) SetGameId(v string) *AddGameToProjectRequest {
	s.GameId = &v
	return s
}

func (s *AddGameToProjectRequest) SetProjectId(v string) *AddGameToProjectRequest {
	s.ProjectId = &v
	return s
}

type AddGameToProjectResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddGameToProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGameToProjectResponseBody) GoString() string {
	return s.String()
}

func (s *AddGameToProjectResponseBody) SetRequestId(v string) *AddGameToProjectResponseBody {
	s.RequestId = &v
	return s
}

type AddGameToProjectResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddGameToProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGameToProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGameToProjectResponse) GoString() string {
	return s.String()
}

func (s *AddGameToProjectResponse) SetHeaders(v map[string]*string) *AddGameToProjectResponse {
	s.Headers = v
	return s
}

func (s *AddGameToProjectResponse) SetBody(v *AddGameToProjectResponseBody) *AddGameToProjectResponse {
	s.Body = v
	return s
}

type BatchDispatchGameSlotRequest struct {
	QueueUserList *string `json:"QueueUserList,omitempty" xml:"QueueUserList,omitempty"`
}

func (s BatchDispatchGameSlotRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDispatchGameSlotRequest) GoString() string {
	return s.String()
}

func (s *BatchDispatchGameSlotRequest) SetQueueUserList(v string) *BatchDispatchGameSlotRequest {
	s.QueueUserList = &v
	return s
}

type BatchDispatchGameSlotResponseBody struct {
	QueueResultList []*BatchDispatchGameSlotResponseBodyQueueResultList `json:"QueueResultList,omitempty" xml:"QueueResultList,omitempty" type:"Repeated"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDispatchGameSlotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDispatchGameSlotResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDispatchGameSlotResponseBody) SetQueueResultList(v []*BatchDispatchGameSlotResponseBodyQueueResultList) *BatchDispatchGameSlotResponseBody {
	s.QueueResultList = v
	return s
}

func (s *BatchDispatchGameSlotResponseBody) SetRequestId(v string) *BatchDispatchGameSlotResponseBody {
	s.RequestId = &v
	return s
}

type BatchDispatchGameSlotResponseBodyQueueResultList struct {
	GameId      *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	GameSession *string `json:"GameSession,omitempty" xml:"GameSession,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	QueueCode   *int32  `json:"QueueCode,omitempty" xml:"QueueCode,omitempty"`
	QueueState  *int32  `json:"QueueState,omitempty" xml:"QueueState,omitempty"`
	RegionName  *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s BatchDispatchGameSlotResponseBodyQueueResultList) String() string {
	return tea.Prettify(s)
}

func (s BatchDispatchGameSlotResponseBodyQueueResultList) GoString() string {
	return s.String()
}

func (s *BatchDispatchGameSlotResponseBodyQueueResultList) SetGameId(v string) *BatchDispatchGameSlotResponseBodyQueueResultList {
	s.GameId = &v
	return s
}

func (s *BatchDispatchGameSlotResponseBodyQueueResultList) SetGameSession(v string) *BatchDispatchGameSlotResponseBodyQueueResultList {
	s.GameSession = &v
	return s
}

func (s *BatchDispatchGameSlotResponseBodyQueueResultList) SetMessage(v string) *BatchDispatchGameSlotResponseBodyQueueResultList {
	s.Message = &v
	return s
}

func (s *BatchDispatchGameSlotResponseBodyQueueResultList) SetQueueCode(v int32) *BatchDispatchGameSlotResponseBodyQueueResultList {
	s.QueueCode = &v
	return s
}

func (s *BatchDispatchGameSlotResponseBodyQueueResultList) SetQueueState(v int32) *BatchDispatchGameSlotResponseBodyQueueResultList {
	s.QueueState = &v
	return s
}

func (s *BatchDispatchGameSlotResponseBodyQueueResultList) SetRegionName(v string) *BatchDispatchGameSlotResponseBodyQueueResultList {
	s.RegionName = &v
	return s
}

func (s *BatchDispatchGameSlotResponseBodyQueueResultList) SetUserId(v string) *BatchDispatchGameSlotResponseBodyQueueResultList {
	s.UserId = &v
	return s
}

type BatchDispatchGameSlotResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchDispatchGameSlotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDispatchGameSlotResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDispatchGameSlotResponse) GoString() string {
	return s.String()
}

func (s *BatchDispatchGameSlotResponse) SetHeaders(v map[string]*string) *BatchDispatchGameSlotResponse {
	s.Headers = v
	return s
}

func (s *BatchDispatchGameSlotResponse) SetBody(v *BatchDispatchGameSlotResponseBody) *BatchDispatchGameSlotResponse {
	s.Body = v
	return s
}

type BatchStopGameSessionsRequest struct {
	GameId    *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Reason    *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Tags      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
	TrackInfo *string `json:"TrackInfo,omitempty" xml:"TrackInfo,omitempty"`
}

func (s BatchStopGameSessionsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStopGameSessionsRequest) GoString() string {
	return s.String()
}

func (s *BatchStopGameSessionsRequest) SetGameId(v string) *BatchStopGameSessionsRequest {
	s.GameId = &v
	return s
}

func (s *BatchStopGameSessionsRequest) SetProjectId(v string) *BatchStopGameSessionsRequest {
	s.ProjectId = &v
	return s
}

func (s *BatchStopGameSessionsRequest) SetReason(v string) *BatchStopGameSessionsRequest {
	s.Reason = &v
	return s
}

func (s *BatchStopGameSessionsRequest) SetTags(v string) *BatchStopGameSessionsRequest {
	s.Tags = &v
	return s
}

func (s *BatchStopGameSessionsRequest) SetToken(v string) *BatchStopGameSessionsRequest {
	s.Token = &v
	return s
}

func (s *BatchStopGameSessionsRequest) SetTrackInfo(v string) *BatchStopGameSessionsRequest {
	s.TrackInfo = &v
	return s
}

type BatchStopGameSessionsResponseBody struct {
	GameId     *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	QueueState *int32  `json:"QueueState,omitempty" xml:"QueueState,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TrackInfo  *string `json:"TrackInfo,omitempty" xml:"TrackInfo,omitempty"`
}

func (s BatchStopGameSessionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStopGameSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopGameSessionsResponseBody) SetGameId(v string) *BatchStopGameSessionsResponseBody {
	s.GameId = &v
	return s
}

func (s *BatchStopGameSessionsResponseBody) SetMessage(v string) *BatchStopGameSessionsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchStopGameSessionsResponseBody) SetProjectId(v string) *BatchStopGameSessionsResponseBody {
	s.ProjectId = &v
	return s
}

func (s *BatchStopGameSessionsResponseBody) SetQueueState(v int32) *BatchStopGameSessionsResponseBody {
	s.QueueState = &v
	return s
}

func (s *BatchStopGameSessionsResponseBody) SetRequestId(v string) *BatchStopGameSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStopGameSessionsResponseBody) SetSuccess(v bool) *BatchStopGameSessionsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchStopGameSessionsResponseBody) SetTrackInfo(v string) *BatchStopGameSessionsResponseBody {
	s.TrackInfo = &v
	return s
}

type BatchStopGameSessionsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchStopGameSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStopGameSessionsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStopGameSessionsResponse) GoString() string {
	return s.String()
}

func (s *BatchStopGameSessionsResponse) SetHeaders(v map[string]*string) *BatchStopGameSessionsResponse {
	s.Headers = v
	return s
}

func (s *BatchStopGameSessionsResponse) SetBody(v *BatchStopGameSessionsResponseBody) *BatchStopGameSessionsResponse {
	s.Body = v
	return s
}

type CloseOrderRequest struct {
	AccountDomain  *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
	BuyerAccountId *string `json:"BuyerAccountId,omitempty" xml:"BuyerAccountId,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CloseOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseOrderRequest) GoString() string {
	return s.String()
}

func (s *CloseOrderRequest) SetAccountDomain(v string) *CloseOrderRequest {
	s.AccountDomain = &v
	return s
}

func (s *CloseOrderRequest) SetBuyerAccountId(v string) *CloseOrderRequest {
	s.BuyerAccountId = &v
	return s
}

func (s *CloseOrderRequest) SetOrderId(v string) *CloseOrderRequest {
	s.OrderId = &v
	return s
}

type CloseOrderResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CloseOrderResponseBody) SetData(v bool) *CloseOrderResponseBody {
	s.Data = &v
	return s
}

func (s *CloseOrderResponseBody) SetRequestId(v string) *CloseOrderResponseBody {
	s.RequestId = &v
	return s
}

type CloseOrderResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseOrderResponse) GoString() string {
	return s.String()
}

func (s *CloseOrderResponse) SetHeaders(v map[string]*string) *CloseOrderResponse {
	s.Headers = v
	return s
}

func (s *CloseOrderResponse) SetBody(v *CloseOrderResponseBody) *CloseOrderResponse {
	s.Body = v
	return s
}

type CreateGameRequest struct {
	// 幂等参数，1-64位建议使用uuid
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 游戏名称
	GameName *string `json:"GameName,omitempty" xml:"GameName,omitempty"`
	// 平台类型
	PlatformType *int64 `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
}

func (s CreateGameRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGameRequest) GoString() string {
	return s.String()
}

func (s *CreateGameRequest) SetClientToken(v string) *CreateGameRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGameRequest) SetGameName(v string) *CreateGameRequest {
	s.GameName = &v
	return s
}

func (s *CreateGameRequest) SetPlatformType(v int64) *CreateGameRequest {
	s.PlatformType = &v
	return s
}

type CreateGameResponseBody struct {
	// 游戏ID
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGameResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGameResponseBody) SetGameId(v string) *CreateGameResponseBody {
	s.GameId = &v
	return s
}

func (s *CreateGameResponseBody) SetRequestId(v string) *CreateGameResponseBody {
	s.RequestId = &v
	return s
}

type CreateGameResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGameResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGameResponse) GoString() string {
	return s.String()
}

func (s *CreateGameResponse) SetHeaders(v map[string]*string) *CreateGameResponse {
	s.Headers = v
	return s
}

func (s *CreateGameResponse) SetBody(v *CreateGameResponseBody) *CreateGameResponse {
	s.Body = v
	return s
}

type CreateGameDeployWorkflowRequest struct {
	DownloadType *string `json:"DownloadType,omitempty" xml:"DownloadType,omitempty"`
	FileType     *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	FrameRate    *string `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	GameId       *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	GameVersion  *string `json:"GameVersion,omitempty" xml:"GameVersion,omitempty"`
	Hash         *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	Instance     *string `json:"Instance,omitempty" xml:"Instance,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Resolution   *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	VersionName  *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateGameDeployWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGameDeployWorkflowRequest) GoString() string {
	return s.String()
}

func (s *CreateGameDeployWorkflowRequest) SetDownloadType(v string) *CreateGameDeployWorkflowRequest {
	s.DownloadType = &v
	return s
}

func (s *CreateGameDeployWorkflowRequest) SetFileType(v string) *CreateGameDeployWorkflowRequest {
	s.FileType = &v
	return s
}

func (s *CreateGameDeployWorkflowRequest) SetFrameRate(v string) *CreateGameDeployWorkflowRequest {
	s.FrameRate = &v
	return s
}

func (s *CreateGameDeployWorkflowRequest) SetGameId(v string) *CreateGameDeployWorkflowRequest {
	s.GameId = &v
	return s
}

func (s *CreateGameDeployWorkflowRequest) SetGameVersion(v string) *CreateGameDeployWorkflowRequest {
	s.GameVersion = &v
	return s
}

func (s *CreateGameDeployWorkflowRequest) SetHash(v string) *CreateGameDeployWorkflowRequest {
	s.Hash = &v
	return s
}

func (s *CreateGameDeployWorkflowRequest) SetInstance(v string) *CreateGameDeployWorkflowRequest {
	s.Instance = &v
	return s
}

func (s *CreateGameDeployWorkflowRequest) SetProjectId(v string) *CreateGameDeployWorkflowRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateGameDeployWorkflowRequest) SetResolution(v string) *CreateGameDeployWorkflowRequest {
	s.Resolution = &v
	return s
}

func (s *CreateGameDeployWorkflowRequest) SetVersionName(v string) *CreateGameDeployWorkflowRequest {
	s.VersionName = &v
	return s
}

type CreateGameDeployWorkflowResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 任务id
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateGameDeployWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGameDeployWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGameDeployWorkflowResponseBody) SetRequestId(v string) *CreateGameDeployWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGameDeployWorkflowResponseBody) SetTaskId(v string) *CreateGameDeployWorkflowResponseBody {
	s.TaskId = &v
	return s
}

type CreateGameDeployWorkflowResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGameDeployWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGameDeployWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGameDeployWorkflowResponse) GoString() string {
	return s.String()
}

func (s *CreateGameDeployWorkflowResponse) SetHeaders(v map[string]*string) *CreateGameDeployWorkflowResponse {
	s.Headers = v
	return s
}

func (s *CreateGameDeployWorkflowResponse) SetBody(v *CreateGameDeployWorkflowResponseBody) *CreateGameDeployWorkflowResponse {
	s.Body = v
	return s
}

type CreateOrderRequest struct {
	AccountDomain   *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
	Amount          *int64  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	BuyerAccountId  *string `json:"BuyerAccountId,omitempty" xml:"BuyerAccountId,omitempty"`
	IdempotentCode  *string `json:"IdempotentCode,omitempty" xml:"IdempotentCode,omitempty"`
	ItemId          *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	OriginPrice     *int64  `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	SettlementPrice *int64  `json:"SettlementPrice,omitempty" xml:"SettlementPrice,omitempty"`
	SkuId           *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) SetAccountDomain(v string) *CreateOrderRequest {
	s.AccountDomain = &v
	return s
}

func (s *CreateOrderRequest) SetAmount(v int64) *CreateOrderRequest {
	s.Amount = &v
	return s
}

func (s *CreateOrderRequest) SetBuyerAccountId(v string) *CreateOrderRequest {
	s.BuyerAccountId = &v
	return s
}

func (s *CreateOrderRequest) SetIdempotentCode(v string) *CreateOrderRequest {
	s.IdempotentCode = &v
	return s
}

func (s *CreateOrderRequest) SetItemId(v string) *CreateOrderRequest {
	s.ItemId = &v
	return s
}

func (s *CreateOrderRequest) SetOriginPrice(v int64) *CreateOrderRequest {
	s.OriginPrice = &v
	return s
}

func (s *CreateOrderRequest) SetSettlementPrice(v int64) *CreateOrderRequest {
	s.SettlementPrice = &v
	return s
}

func (s *CreateOrderRequest) SetSkuId(v string) *CreateOrderRequest {
	s.SkuId = &v
	return s
}

type CreateOrderResponseBody struct {
	Data      *CreateOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) SetData(v *CreateOrderResponseBodyData) *CreateOrderResponseBody {
	s.Data = v
	return s
}

func (s *CreateOrderResponseBody) SetRequestId(v string) *CreateOrderResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrderResponseBodyData struct {
	AccountDomain     *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
	Amount            *int64  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	ApplyDeliveryTime *int64  `json:"ApplyDeliveryTime,omitempty" xml:"ApplyDeliveryTime,omitempty"`
	AutoUnlockTime    *int64  `json:"AutoUnlockTime,omitempty" xml:"AutoUnlockTime,omitempty"`
	BuyerAccountId    *string `json:"BuyerAccountId,omitempty" xml:"BuyerAccountId,omitempty"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FinishTime        *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	ItemId            *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	OrderId           *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OriginPrice       *int64  `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	SettlementPrice   *int64  `json:"SettlementPrice,omitempty" xml:"SettlementPrice,omitempty"`
	SkuId             *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyData) SetAccountDomain(v string) *CreateOrderResponseBodyData {
	s.AccountDomain = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetAmount(v int64) *CreateOrderResponseBodyData {
	s.Amount = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetApplyDeliveryTime(v int64) *CreateOrderResponseBodyData {
	s.ApplyDeliveryTime = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetAutoUnlockTime(v int64) *CreateOrderResponseBodyData {
	s.AutoUnlockTime = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetBuyerAccountId(v string) *CreateOrderResponseBodyData {
	s.BuyerAccountId = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetCreateTime(v int64) *CreateOrderResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetFinishTime(v int64) *CreateOrderResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetItemId(v string) *CreateOrderResponseBodyData {
	s.ItemId = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetOrderId(v string) *CreateOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetOriginPrice(v int64) *CreateOrderResponseBodyData {
	s.OriginPrice = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetSettlementPrice(v int64) *CreateOrderResponseBodyData {
	s.SettlementPrice = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetSkuId(v string) *CreateOrderResponseBodyData {
	s.SkuId = &v
	return s
}

func (s *CreateOrderResponseBodyData) SetStatus(v string) *CreateOrderResponseBodyData {
	s.Status = &v
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

type CreateProjectRequest struct {
	// 幂等参数，1-64位建议使用uuid
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetClientToken(v string) *CreateProjectRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

type CreateProjectResponseBody struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetProjectId(v string) *CreateProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateTokenRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CurrentToken *string `json:"CurrentToken,omitempty" xml:"CurrentToken,omitempty"`
	Session      *string `json:"Session,omitempty" xml:"Session,omitempty"`
}

func (s CreateTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateTokenRequest) SetClientToken(v string) *CreateTokenRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTokenRequest) SetCurrentToken(v string) *CreateTokenRequest {
	s.CurrentToken = &v
	return s
}

func (s *CreateTokenRequest) SetSession(v string) *CreateTokenRequest {
	s.Session = &v
	return s
}

type CreateTokenResponseBody struct {
	Data      *CreateTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBody) SetData(v *CreateTokenResponseBodyData) *CreateTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateTokenResponseBody) SetRequestId(v string) *CreateTokenResponseBody {
	s.RequestId = &v
	return s
}

type CreateTokenResponseBodyData struct {
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBodyData) SetToken(v string) *CreateTokenResponseBodyData {
	s.Token = &v
	return s
}

type CreateTokenResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateTokenResponse) SetHeaders(v map[string]*string) *CreateTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateTokenResponse) SetBody(v *CreateTokenResponseBody) *CreateTokenResponse {
	s.Body = v
	return s
}

type DeleteGameRequest struct {
	// 游戏ID
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
}

func (s DeleteGameRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGameRequest) GoString() string {
	return s.String()
}

func (s *DeleteGameRequest) SetGameId(v string) *DeleteGameRequest {
	s.GameId = &v
	return s
}

type DeleteGameResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGameResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGameResponseBody) SetRequestId(v string) *DeleteGameResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGameResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGameResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGameResponse) GoString() string {
	return s.String()
}

func (s *DeleteGameResponse) SetHeaders(v map[string]*string) *DeleteGameResponse {
	s.Headers = v
	return s
}

func (s *DeleteGameResponse) SetBody(v *DeleteGameResponseBody) *DeleteGameResponse {
	s.Body = v
	return s
}

type DeleteGameVersionRequest struct {
	// 游戏版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeleteGameVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGameVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteGameVersionRequest) SetVersionId(v string) *DeleteGameVersionRequest {
	s.VersionId = &v
	return s
}

type DeleteGameVersionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGameVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGameVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGameVersionResponseBody) SetRequestId(v string) *DeleteGameVersionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGameVersionResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGameVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGameVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGameVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteGameVersionResponse) SetHeaders(v map[string]*string) *DeleteGameVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteGameVersionResponse) SetBody(v *DeleteGameVersionResponseBody) *DeleteGameVersionResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetProjectId(v string) *DeleteProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteProjectResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DeliveryOrderRequest struct {
	AccountDomain  *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
	BuyerAccountId *string `json:"BuyerAccountId,omitempty" xml:"BuyerAccountId,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DeliveryOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeliveryOrderRequest) GoString() string {
	return s.String()
}

func (s *DeliveryOrderRequest) SetAccountDomain(v string) *DeliveryOrderRequest {
	s.AccountDomain = &v
	return s
}

func (s *DeliveryOrderRequest) SetBuyerAccountId(v string) *DeliveryOrderRequest {
	s.BuyerAccountId = &v
	return s
}

func (s *DeliveryOrderRequest) SetOrderId(v string) *DeliveryOrderRequest {
	s.OrderId = &v
	return s
}

type DeliveryOrderResponseBody struct {
	Data      *DeliveryOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeliveryOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeliveryOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DeliveryOrderResponseBody) SetData(v *DeliveryOrderResponseBodyData) *DeliveryOrderResponseBody {
	s.Data = v
	return s
}

func (s *DeliveryOrderResponseBody) SetRequestId(v string) *DeliveryOrderResponseBody {
	s.RequestId = &v
	return s
}

type DeliveryOrderResponseBodyData struct {
	DeliveryStatus *string `json:"DeliveryStatus,omitempty" xml:"DeliveryStatus,omitempty"`
}

func (s DeliveryOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeliveryOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeliveryOrderResponseBodyData) SetDeliveryStatus(v string) *DeliveryOrderResponseBodyData {
	s.DeliveryStatus = &v
	return s
}

type DeliveryOrderResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeliveryOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeliveryOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeliveryOrderResponse) GoString() string {
	return s.String()
}

func (s *DeliveryOrderResponse) SetHeaders(v map[string]*string) *DeliveryOrderResponse {
	s.Headers = v
	return s
}

func (s *DeliveryOrderResponse) SetBody(v *DeliveryOrderResponseBody) *DeliveryOrderResponse {
	s.Body = v
	return s
}

type DispatchGameSlotRequest struct {
	AccessKey      *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	BizParam       *string `json:"BizParam,omitempty" xml:"BizParam,omitempty"`
	Cancel         *bool   `json:"Cancel,omitempty" xml:"Cancel,omitempty"`
	ClientIp       *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	GameCommand    *string `json:"GameCommand,omitempty" xml:"GameCommand,omitempty"`
	GameId         *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	GameSession    *string `json:"GameSession,omitempty" xml:"GameSession,omitempty"`
	GameStartParam *string `json:"GameStartParam,omitempty" xml:"GameStartParam,omitempty"`
	Reconnect      *bool   `json:"Reconnect,omitempty" xml:"Reconnect,omitempty"`
	RegionName     *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	SystemInfo     *string `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty"`
	Tags           *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserLevel      *int32  `json:"UserLevel,omitempty" xml:"UserLevel,omitempty"`
}

func (s DispatchGameSlotRequest) String() string {
	return tea.Prettify(s)
}

func (s DispatchGameSlotRequest) GoString() string {
	return s.String()
}

func (s *DispatchGameSlotRequest) SetAccessKey(v string) *DispatchGameSlotRequest {
	s.AccessKey = &v
	return s
}

func (s *DispatchGameSlotRequest) SetBizParam(v string) *DispatchGameSlotRequest {
	s.BizParam = &v
	return s
}

func (s *DispatchGameSlotRequest) SetCancel(v bool) *DispatchGameSlotRequest {
	s.Cancel = &v
	return s
}

func (s *DispatchGameSlotRequest) SetClientIp(v string) *DispatchGameSlotRequest {
	s.ClientIp = &v
	return s
}

func (s *DispatchGameSlotRequest) SetGameCommand(v string) *DispatchGameSlotRequest {
	s.GameCommand = &v
	return s
}

func (s *DispatchGameSlotRequest) SetGameId(v string) *DispatchGameSlotRequest {
	s.GameId = &v
	return s
}

func (s *DispatchGameSlotRequest) SetGameSession(v string) *DispatchGameSlotRequest {
	s.GameSession = &v
	return s
}

func (s *DispatchGameSlotRequest) SetGameStartParam(v string) *DispatchGameSlotRequest {
	s.GameStartParam = &v
	return s
}

func (s *DispatchGameSlotRequest) SetReconnect(v bool) *DispatchGameSlotRequest {
	s.Reconnect = &v
	return s
}

func (s *DispatchGameSlotRequest) SetRegionName(v string) *DispatchGameSlotRequest {
	s.RegionName = &v
	return s
}

func (s *DispatchGameSlotRequest) SetSystemInfo(v string) *DispatchGameSlotRequest {
	s.SystemInfo = &v
	return s
}

func (s *DispatchGameSlotRequest) SetTags(v string) *DispatchGameSlotRequest {
	s.Tags = &v
	return s
}

func (s *DispatchGameSlotRequest) SetUserId(v string) *DispatchGameSlotRequest {
	s.UserId = &v
	return s
}

func (s *DispatchGameSlotRequest) SetUserLevel(v int32) *DispatchGameSlotRequest {
	s.UserLevel = &v
	return s
}

type DispatchGameSlotResponseBody struct {
	GameId      *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	GameSession *string `json:"GameSession,omitempty" xml:"GameSession,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	QueueCode   *int32  `json:"QueueCode,omitempty" xml:"QueueCode,omitempty"`
	QueueState  *int32  `json:"QueueState,omitempty" xml:"QueueState,omitempty"`
	RegionName  *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DispatchGameSlotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DispatchGameSlotResponseBody) GoString() string {
	return s.String()
}

func (s *DispatchGameSlotResponseBody) SetGameId(v string) *DispatchGameSlotResponseBody {
	s.GameId = &v
	return s
}

func (s *DispatchGameSlotResponseBody) SetGameSession(v string) *DispatchGameSlotResponseBody {
	s.GameSession = &v
	return s
}

func (s *DispatchGameSlotResponseBody) SetMessage(v string) *DispatchGameSlotResponseBody {
	s.Message = &v
	return s
}

func (s *DispatchGameSlotResponseBody) SetQueueCode(v int32) *DispatchGameSlotResponseBody {
	s.QueueCode = &v
	return s
}

func (s *DispatchGameSlotResponseBody) SetQueueState(v int32) *DispatchGameSlotResponseBody {
	s.QueueState = &v
	return s
}

func (s *DispatchGameSlotResponseBody) SetRegionName(v string) *DispatchGameSlotResponseBody {
	s.RegionName = &v
	return s
}

func (s *DispatchGameSlotResponseBody) SetRequestId(v string) *DispatchGameSlotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DispatchGameSlotResponseBody) SetUserId(v string) *DispatchGameSlotResponseBody {
	s.UserId = &v
	return s
}

type DispatchGameSlotResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DispatchGameSlotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DispatchGameSlotResponse) String() string {
	return tea.Prettify(s)
}

func (s DispatchGameSlotResponse) GoString() string {
	return s.String()
}

func (s *DispatchGameSlotResponse) SetHeaders(v map[string]*string) *DispatchGameSlotResponse {
	s.Headers = v
	return s
}

func (s *DispatchGameSlotResponse) SetBody(v *DispatchGameSlotResponseBody) *DispatchGameSlotResponse {
	s.Body = v
	return s
}

type GetGameCcuRequest struct {
	AccessKey  *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	GameId     *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s GetGameCcuRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGameCcuRequest) GoString() string {
	return s.String()
}

func (s *GetGameCcuRequest) SetAccessKey(v string) *GetGameCcuRequest {
	s.AccessKey = &v
	return s
}

func (s *GetGameCcuRequest) SetGameId(v string) *GetGameCcuRequest {
	s.GameId = &v
	return s
}

func (s *GetGameCcuRequest) SetRegionName(v string) *GetGameCcuRequest {
	s.RegionName = &v
	return s
}

type GetGameCcuResponseBody struct {
	DataList  []*GetGameCcuResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGameCcuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGameCcuResponseBody) GoString() string {
	return s.String()
}

func (s *GetGameCcuResponseBody) SetDataList(v []*GetGameCcuResponseBodyDataList) *GetGameCcuResponseBody {
	s.DataList = v
	return s
}

func (s *GetGameCcuResponseBody) SetRequestId(v string) *GetGameCcuResponseBody {
	s.RequestId = &v
	return s
}

type GetGameCcuResponseBodyDataList struct {
	Ccu      *int64  `json:"Ccu,omitempty" xml:"Ccu,omitempty"`
	GameId   *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetGameCcuResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetGameCcuResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetGameCcuResponseBodyDataList) SetCcu(v int64) *GetGameCcuResponseBodyDataList {
	s.Ccu = &v
	return s
}

func (s *GetGameCcuResponseBodyDataList) SetGameId(v string) *GetGameCcuResponseBodyDataList {
	s.GameId = &v
	return s
}

func (s *GetGameCcuResponseBodyDataList) SetRegionId(v string) *GetGameCcuResponseBodyDataList {
	s.RegionId = &v
	return s
}

type GetGameCcuResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGameCcuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGameCcuResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGameCcuResponse) GoString() string {
	return s.String()
}

func (s *GetGameCcuResponse) SetHeaders(v map[string]*string) *GetGameCcuResponse {
	s.Headers = v
	return s
}

func (s *GetGameCcuResponse) SetBody(v *GetGameCcuResponseBody) *GetGameCcuResponse {
	s.Body = v
	return s
}

type GetGameStockRequest struct {
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	GameId    *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	UserLevel *int64  `json:"UserLevel,omitempty" xml:"UserLevel,omitempty"`
}

func (s GetGameStockRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGameStockRequest) GoString() string {
	return s.String()
}

func (s *GetGameStockRequest) SetAccessKey(v string) *GetGameStockRequest {
	s.AccessKey = &v
	return s
}

func (s *GetGameStockRequest) SetGameId(v string) *GetGameStockRequest {
	s.GameId = &v
	return s
}

func (s *GetGameStockRequest) SetUserLevel(v int64) *GetGameStockRequest {
	s.UserLevel = &v
	return s
}

type GetGameStockResponseBody struct {
	GameId            *string                                      `json:"GameId,omitempty" xml:"GameId,omitempty"`
	InstanceStockList []*GetGameStockResponseBodyInstanceStockList `json:"InstanceStockList,omitempty" xml:"InstanceStockList,omitempty" type:"Repeated"`
	Message           *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId         *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGameStockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGameStockResponseBody) GoString() string {
	return s.String()
}

func (s *GetGameStockResponseBody) SetGameId(v string) *GetGameStockResponseBody {
	s.GameId = &v
	return s
}

func (s *GetGameStockResponseBody) SetInstanceStockList(v []*GetGameStockResponseBodyInstanceStockList) *GetGameStockResponseBody {
	s.InstanceStockList = v
	return s
}

func (s *GetGameStockResponseBody) SetMessage(v string) *GetGameStockResponseBody {
	s.Message = &v
	return s
}

func (s *GetGameStockResponseBody) SetRequestId(v string) *GetGameStockResponseBody {
	s.RequestId = &v
	return s
}

type GetGameStockResponseBodyInstanceStockList struct {
	AvailableSlots *int64  `json:"AvailableSlots,omitempty" xml:"AvailableSlots,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceSpec   *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	ReginName      *string `json:"ReginName,omitempty" xml:"ReginName,omitempty"`
	UserLevel      *int64  `json:"UserLevel,omitempty" xml:"UserLevel,omitempty"`
}

func (s GetGameStockResponseBodyInstanceStockList) String() string {
	return tea.Prettify(s)
}

func (s GetGameStockResponseBodyInstanceStockList) GoString() string {
	return s.String()
}

func (s *GetGameStockResponseBodyInstanceStockList) SetAvailableSlots(v int64) *GetGameStockResponseBodyInstanceStockList {
	s.AvailableSlots = &v
	return s
}

func (s *GetGameStockResponseBodyInstanceStockList) SetInstanceId(v string) *GetGameStockResponseBodyInstanceStockList {
	s.InstanceId = &v
	return s
}

func (s *GetGameStockResponseBodyInstanceStockList) SetInstanceSpec(v string) *GetGameStockResponseBodyInstanceStockList {
	s.InstanceSpec = &v
	return s
}

func (s *GetGameStockResponseBodyInstanceStockList) SetReginName(v string) *GetGameStockResponseBodyInstanceStockList {
	s.ReginName = &v
	return s
}

func (s *GetGameStockResponseBodyInstanceStockList) SetUserLevel(v int64) *GetGameStockResponseBodyInstanceStockList {
	s.UserLevel = &v
	return s
}

type GetGameStockResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGameStockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGameStockResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGameStockResponse) GoString() string {
	return s.String()
}

func (s *GetGameStockResponse) SetHeaders(v map[string]*string) *GetGameStockResponse {
	s.Headers = v
	return s
}

func (s *GetGameStockResponse) SetBody(v *GetGameStockResponseBody) *GetGameStockResponse {
	s.Body = v
	return s
}

type GetGameTrialSurplusDurationRequest struct {
	// 账号ID
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// 游戏ID
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetGameTrialSurplusDurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGameTrialSurplusDurationRequest) GoString() string {
	return s.String()
}

func (s *GetGameTrialSurplusDurationRequest) SetAccountId(v string) *GetGameTrialSurplusDurationRequest {
	s.AccountId = &v
	return s
}

func (s *GetGameTrialSurplusDurationRequest) SetGameId(v string) *GetGameTrialSurplusDurationRequest {
	s.GameId = &v
	return s
}

func (s *GetGameTrialSurplusDurationRequest) SetProjectId(v string) *GetGameTrialSurplusDurationRequest {
	s.ProjectId = &v
	return s
}

type GetGameTrialSurplusDurationResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 状态
	Status *float32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 剩余试玩时长
	SurplusDuration *float64 `json:"SurplusDuration,omitempty" xml:"SurplusDuration,omitempty"`
}

func (s GetGameTrialSurplusDurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGameTrialSurplusDurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetGameTrialSurplusDurationResponseBody) SetRequestId(v string) *GetGameTrialSurplusDurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGameTrialSurplusDurationResponseBody) SetStatus(v float32) *GetGameTrialSurplusDurationResponseBody {
	s.Status = &v
	return s
}

func (s *GetGameTrialSurplusDurationResponseBody) SetSurplusDuration(v float64) *GetGameTrialSurplusDurationResponseBody {
	s.SurplusDuration = &v
	return s
}

type GetGameTrialSurplusDurationResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGameTrialSurplusDurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGameTrialSurplusDurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGameTrialSurplusDurationResponse) GoString() string {
	return s.String()
}

func (s *GetGameTrialSurplusDurationResponse) SetHeaders(v map[string]*string) *GetGameTrialSurplusDurationResponse {
	s.Headers = v
	return s
}

func (s *GetGameTrialSurplusDurationResponse) SetBody(v *GetGameTrialSurplusDurationResponseBody) *GetGameTrialSurplusDurationResponse {
	s.Body = v
	return s
}

type GetGameVersionRequest struct {
	// 版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetGameVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGameVersionRequest) GoString() string {
	return s.String()
}

func (s *GetGameVersionRequest) SetVersionId(v string) *GetGameVersionRequest {
	s.VersionId = &v
	return s
}

type GetGameVersionResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// 版本名称
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// 版本号
	VersionNumber *string `json:"VersionNumber,omitempty" xml:"VersionNumber,omitempty"`
}

func (s GetGameVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGameVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetGameVersionResponseBody) SetRequestId(v string) *GetGameVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGameVersionResponseBody) SetVersionId(v string) *GetGameVersionResponseBody {
	s.VersionId = &v
	return s
}

func (s *GetGameVersionResponseBody) SetVersionName(v string) *GetGameVersionResponseBody {
	s.VersionName = &v
	return s
}

func (s *GetGameVersionResponseBody) SetVersionNumber(v string) *GetGameVersionResponseBody {
	s.VersionNumber = &v
	return s
}

type GetGameVersionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGameVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGameVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGameVersionResponse) GoString() string {
	return s.String()
}

func (s *GetGameVersionResponse) SetHeaders(v map[string]*string) *GetGameVersionResponse {
	s.Headers = v
	return s
}

func (s *GetGameVersionResponse) SetBody(v *GetGameVersionResponseBody) *GetGameVersionResponse {
	s.Body = v
	return s
}

type GetGameVersionProgressRequest struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetGameVersionProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGameVersionProgressRequest) GoString() string {
	return s.String()
}

func (s *GetGameVersionProgressRequest) SetTaskId(v string) *GetGameVersionProgressRequest {
	s.TaskId = &v
	return s
}

type GetGameVersionProgressResponseBody struct {
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Event       *string                `json:"Event,omitempty" xml:"Event,omitempty"`
	Extra       map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetGameVersionProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGameVersionProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetGameVersionProgressResponseBody) SetDescription(v string) *GetGameVersionProgressResponseBody {
	s.Description = &v
	return s
}

func (s *GetGameVersionProgressResponseBody) SetEvent(v string) *GetGameVersionProgressResponseBody {
	s.Event = &v
	return s
}

func (s *GetGameVersionProgressResponseBody) SetExtra(v map[string]interface{}) *GetGameVersionProgressResponseBody {
	s.Extra = v
	return s
}

func (s *GetGameVersionProgressResponseBody) SetRequestId(v string) *GetGameVersionProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGameVersionProgressResponseBody) SetStatus(v string) *GetGameVersionProgressResponseBody {
	s.Status = &v
	return s
}

type GetGameVersionProgressResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGameVersionProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGameVersionProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGameVersionProgressResponse) GoString() string {
	return s.String()
}

func (s *GetGameVersionProgressResponse) SetHeaders(v map[string]*string) *GetGameVersionProgressResponse {
	s.Headers = v
	return s
}

func (s *GetGameVersionProgressResponse) SetBody(v *GetGameVersionProgressResponseBody) *GetGameVersionProgressResponse {
	s.Body = v
	return s
}

type GetItemRequest struct {
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
}

func (s GetItemRequest) String() string {
	return tea.Prettify(s)
}

func (s GetItemRequest) GoString() string {
	return s.String()
}

func (s *GetItemRequest) SetItemId(v string) *GetItemRequest {
	s.ItemId = &v
	return s
}

type GetItemResponseBody struct {
	Data      *GetItemResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetItemResponseBody) SetData(v *GetItemResponseBodyData) *GetItemResponseBody {
	s.Data = v
	return s
}

func (s *GetItemResponseBody) SetRequestId(v string) *GetItemResponseBody {
	s.RequestId = &v
	return s
}

type GetItemResponseBodyData struct {
	CategoryId  *int64                          `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreateTime  *int64                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Games       []*GetItemResponseBodyDataGames `json:"Games,omitempty" xml:"Games,omitempty" type:"Repeated"`
	ItemId      *string                         `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ModifyTime  *int64                          `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	OriginPrice *int64                          `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	SalePrice   *int64                          `json:"SalePrice,omitempty" xml:"SalePrice,omitempty"`
	SellerId    *string                         `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	Skus        []*GetItemResponseBodyDataSkus  `json:"Skus,omitempty" xml:"Skus,omitempty" type:"Repeated"`
	Status      *int32                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Supplier    *string                         `json:"Supplier,omitempty" xml:"Supplier,omitempty"`
	Title       *string                         `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetItemResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetItemResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetItemResponseBodyData) SetCategoryId(v int64) *GetItemResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *GetItemResponseBodyData) SetCreateTime(v int64) *GetItemResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetItemResponseBodyData) SetDescription(v string) *GetItemResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetItemResponseBodyData) SetGames(v []*GetItemResponseBodyDataGames) *GetItemResponseBodyData {
	s.Games = v
	return s
}

func (s *GetItemResponseBodyData) SetItemId(v string) *GetItemResponseBodyData {
	s.ItemId = &v
	return s
}

func (s *GetItemResponseBodyData) SetModifyTime(v int64) *GetItemResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetItemResponseBodyData) SetOriginPrice(v int64) *GetItemResponseBodyData {
	s.OriginPrice = &v
	return s
}

func (s *GetItemResponseBodyData) SetSalePrice(v int64) *GetItemResponseBodyData {
	s.SalePrice = &v
	return s
}

func (s *GetItemResponseBodyData) SetSellerId(v string) *GetItemResponseBodyData {
	s.SellerId = &v
	return s
}

func (s *GetItemResponseBodyData) SetSkus(v []*GetItemResponseBodyDataSkus) *GetItemResponseBodyData {
	s.Skus = v
	return s
}

func (s *GetItemResponseBodyData) SetStatus(v int32) *GetItemResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetItemResponseBodyData) SetSupplier(v string) *GetItemResponseBodyData {
	s.Supplier = &v
	return s
}

func (s *GetItemResponseBodyData) SetTitle(v string) *GetItemResponseBodyData {
	s.Title = &v
	return s
}

type GetItemResponseBodyDataGames struct {
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetItemResponseBodyDataGames) String() string {
	return tea.Prettify(s)
}

func (s GetItemResponseBodyDataGames) GoString() string {
	return s.String()
}

func (s *GetItemResponseBodyDataGames) SetGameId(v string) *GetItemResponseBodyDataGames {
	s.GameId = &v
	return s
}

func (s *GetItemResponseBodyDataGames) SetName(v string) *GetItemResponseBodyDataGames {
	s.Name = &v
	return s
}

type GetItemResponseBodyDataSkus struct {
	CreateTime  *int64                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ItemId      *string                                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ModifyTime  *int64                                  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	OriginPrice *int64                                  `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	SalePrice   *int64                                  `json:"SalePrice,omitempty" xml:"SalePrice,omitempty"`
	SaleProps   []*GetItemResponseBodyDataSkusSaleProps `json:"SaleProps,omitempty" xml:"SaleProps,omitempty" type:"Repeated"`
	SkuId       *string                                 `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Status      *int32                                  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetItemResponseBodyDataSkus) String() string {
	return tea.Prettify(s)
}

func (s GetItemResponseBodyDataSkus) GoString() string {
	return s.String()
}

func (s *GetItemResponseBodyDataSkus) SetCreateTime(v int64) *GetItemResponseBodyDataSkus {
	s.CreateTime = &v
	return s
}

func (s *GetItemResponseBodyDataSkus) SetItemId(v string) *GetItemResponseBodyDataSkus {
	s.ItemId = &v
	return s
}

func (s *GetItemResponseBodyDataSkus) SetModifyTime(v int64) *GetItemResponseBodyDataSkus {
	s.ModifyTime = &v
	return s
}

func (s *GetItemResponseBodyDataSkus) SetOriginPrice(v int64) *GetItemResponseBodyDataSkus {
	s.OriginPrice = &v
	return s
}

func (s *GetItemResponseBodyDataSkus) SetSalePrice(v int64) *GetItemResponseBodyDataSkus {
	s.SalePrice = &v
	return s
}

func (s *GetItemResponseBodyDataSkus) SetSaleProps(v []*GetItemResponseBodyDataSkusSaleProps) *GetItemResponseBodyDataSkus {
	s.SaleProps = v
	return s
}

func (s *GetItemResponseBodyDataSkus) SetSkuId(v string) *GetItemResponseBodyDataSkus {
	s.SkuId = &v
	return s
}

func (s *GetItemResponseBodyDataSkus) SetStatus(v int32) *GetItemResponseBodyDataSkus {
	s.Status = &v
	return s
}

type GetItemResponseBodyDataSkusSaleProps struct {
	PropertyId   *int64  `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueId      *int64  `json:"ValueId,omitempty" xml:"ValueId,omitempty"`
}

func (s GetItemResponseBodyDataSkusSaleProps) String() string {
	return tea.Prettify(s)
}

func (s GetItemResponseBodyDataSkusSaleProps) GoString() string {
	return s.String()
}

func (s *GetItemResponseBodyDataSkusSaleProps) SetPropertyId(v int64) *GetItemResponseBodyDataSkusSaleProps {
	s.PropertyId = &v
	return s
}

func (s *GetItemResponseBodyDataSkusSaleProps) SetPropertyName(v string) *GetItemResponseBodyDataSkusSaleProps {
	s.PropertyName = &v
	return s
}

func (s *GetItemResponseBodyDataSkusSaleProps) SetValue(v string) *GetItemResponseBodyDataSkusSaleProps {
	s.Value = &v
	return s
}

func (s *GetItemResponseBodyDataSkusSaleProps) SetValueId(v int64) *GetItemResponseBodyDataSkusSaleProps {
	s.ValueId = &v
	return s
}

type GetItemResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetItemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetItemResponse) GoString() string {
	return s.String()
}

func (s *GetItemResponse) SetHeaders(v map[string]*string) *GetItemResponse {
	s.Headers = v
	return s
}

func (s *GetItemResponse) SetBody(v *GetItemResponseBody) *GetItemResponse {
	s.Body = v
	return s
}

type GetOutAccountBindDetailRequest struct {
	AccountDomain  *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	OutAccountType *string `json:"OutAccountType,omitempty" xml:"OutAccountType,omitempty"`
}

func (s GetOutAccountBindDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOutAccountBindDetailRequest) GoString() string {
	return s.String()
}

func (s *GetOutAccountBindDetailRequest) SetAccountDomain(v string) *GetOutAccountBindDetailRequest {
	s.AccountDomain = &v
	return s
}

func (s *GetOutAccountBindDetailRequest) SetAccountId(v string) *GetOutAccountBindDetailRequest {
	s.AccountId = &v
	return s
}

func (s *GetOutAccountBindDetailRequest) SetOutAccountType(v string) *GetOutAccountBindDetailRequest {
	s.OutAccountType = &v
	return s
}

type GetOutAccountBindDetailResponseBody struct {
	Data      *GetOutAccountBindDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOutAccountBindDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOutAccountBindDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetOutAccountBindDetailResponseBody) SetData(v *GetOutAccountBindDetailResponseBodyData) *GetOutAccountBindDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetOutAccountBindDetailResponseBody) SetRequestId(v string) *GetOutAccountBindDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetOutAccountBindDetailResponseBodyData struct {
	BindStatus      *int32  `json:"BindStatus,omitempty" xml:"BindStatus,omitempty"`
	OutAccountId    *string `json:"OutAccountId,omitempty" xml:"OutAccountId,omitempty"`
	OutAccountType  *string `json:"OutAccountType,omitempty" xml:"OutAccountType,omitempty"`
	Token           *string `json:"Token,omitempty" xml:"Token,omitempty"`
	TokenExpireTime *int64  `json:"TokenExpireTime,omitempty" xml:"TokenExpireTime,omitempty"`
}

func (s GetOutAccountBindDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOutAccountBindDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOutAccountBindDetailResponseBodyData) SetBindStatus(v int32) *GetOutAccountBindDetailResponseBodyData {
	s.BindStatus = &v
	return s
}

func (s *GetOutAccountBindDetailResponseBodyData) SetOutAccountId(v string) *GetOutAccountBindDetailResponseBodyData {
	s.OutAccountId = &v
	return s
}

func (s *GetOutAccountBindDetailResponseBodyData) SetOutAccountType(v string) *GetOutAccountBindDetailResponseBodyData {
	s.OutAccountType = &v
	return s
}

func (s *GetOutAccountBindDetailResponseBodyData) SetToken(v string) *GetOutAccountBindDetailResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetOutAccountBindDetailResponseBodyData) SetTokenExpireTime(v int64) *GetOutAccountBindDetailResponseBodyData {
	s.TokenExpireTime = &v
	return s
}

type GetOutAccountBindDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOutAccountBindDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOutAccountBindDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOutAccountBindDetailResponse) GoString() string {
	return s.String()
}

func (s *GetOutAccountBindDetailResponse) SetHeaders(v map[string]*string) *GetOutAccountBindDetailResponse {
	s.Headers = v
	return s
}

func (s *GetOutAccountBindDetailResponse) SetBody(v *GetOutAccountBindDetailResponseBody) *GetOutAccountBindDetailResponse {
	s.Body = v
	return s
}

type GetSessionRequest struct {
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSessionRequest) GoString() string {
	return s.String()
}

func (s *GetSessionRequest) SetToken(v string) *GetSessionRequest {
	s.Token = &v
	return s
}

type GetSessionResponseBody struct {
	Data      *GetSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionResponseBody) SetData(v *GetSessionResponseBodyData) *GetSessionResponseBody {
	s.Data = v
	return s
}

func (s *GetSessionResponseBody) SetRequestId(v string) *GetSessionResponseBody {
	s.RequestId = &v
	return s
}

type GetSessionResponseBodyData struct {
	Session *string `json:"Session,omitempty" xml:"Session,omitempty"`
}

func (s GetSessionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSessionResponseBodyData) SetSession(v string) *GetSessionResponseBodyData {
	s.Session = &v
	return s
}

type GetSessionResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSessionResponse) GoString() string {
	return s.String()
}

func (s *GetSessionResponse) SetHeaders(v map[string]*string) *GetSessionResponse {
	s.Headers = v
	return s
}

func (s *GetSessionResponse) SetBody(v *GetSessionResponseBody) *GetSessionResponse {
	s.Body = v
	return s
}

type GetStopGameTokenRequest struct {
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	GameId    *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
}

func (s GetStopGameTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStopGameTokenRequest) GoString() string {
	return s.String()
}

func (s *GetStopGameTokenRequest) SetAccessKey(v string) *GetStopGameTokenRequest {
	s.AccessKey = &v
	return s
}

func (s *GetStopGameTokenRequest) SetGameId(v string) *GetStopGameTokenRequest {
	s.GameId = &v
	return s
}

type GetStopGameTokenResponseBody struct {
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetStopGameTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStopGameTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetStopGameTokenResponseBody) SetExpireTime(v int64) *GetStopGameTokenResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetStopGameTokenResponseBody) SetRequestId(v string) *GetStopGameTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStopGameTokenResponseBody) SetToken(v string) *GetStopGameTokenResponseBody {
	s.Token = &v
	return s
}

type GetStopGameTokenResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStopGameTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStopGameTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStopGameTokenResponse) GoString() string {
	return s.String()
}

func (s *GetStopGameTokenResponse) SetHeaders(v map[string]*string) *GetStopGameTokenResponse {
	s.Headers = v
	return s
}

func (s *GetStopGameTokenResponse) SetBody(v *GetStopGameTokenResponseBody) *GetStopGameTokenResponse {
	s.Body = v
	return s
}

type ListBoughtGamesRequest struct {
	AccountDomain *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
	AccountId     *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListBoughtGamesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBoughtGamesRequest) GoString() string {
	return s.String()
}

func (s *ListBoughtGamesRequest) SetAccountDomain(v string) *ListBoughtGamesRequest {
	s.AccountDomain = &v
	return s
}

func (s *ListBoughtGamesRequest) SetAccountId(v string) *ListBoughtGamesRequest {
	s.AccountId = &v
	return s
}

func (s *ListBoughtGamesRequest) SetPageNumber(v int32) *ListBoughtGamesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBoughtGamesRequest) SetPageSize(v int32) *ListBoughtGamesRequest {
	s.PageSize = &v
	return s
}

type ListBoughtGamesResponseBody struct {
	Items      []*ListBoughtGamesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBoughtGamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBoughtGamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBoughtGamesResponseBody) SetItems(v []*ListBoughtGamesResponseBodyItems) *ListBoughtGamesResponseBody {
	s.Items = v
	return s
}

func (s *ListBoughtGamesResponseBody) SetPageNumber(v int32) *ListBoughtGamesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBoughtGamesResponseBody) SetPageSize(v int32) *ListBoughtGamesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBoughtGamesResponseBody) SetRequestId(v string) *ListBoughtGamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBoughtGamesResponseBody) SetTotalCount(v int32) *ListBoughtGamesResponseBody {
	s.TotalCount = &v
	return s
}

type ListBoughtGamesResponseBodyItems struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GameId    *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	GameName  *string `json:"GameName,omitempty" xml:"GameName,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListBoughtGamesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListBoughtGamesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListBoughtGamesResponseBodyItems) SetEndTime(v int64) *ListBoughtGamesResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *ListBoughtGamesResponseBodyItems) SetGameId(v string) *ListBoughtGamesResponseBodyItems {
	s.GameId = &v
	return s
}

func (s *ListBoughtGamesResponseBodyItems) SetGameName(v string) *ListBoughtGamesResponseBodyItems {
	s.GameName = &v
	return s
}

func (s *ListBoughtGamesResponseBodyItems) SetStartTime(v int64) *ListBoughtGamesResponseBodyItems {
	s.StartTime = &v
	return s
}

type ListBoughtGamesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBoughtGamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBoughtGamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBoughtGamesResponse) GoString() string {
	return s.String()
}

func (s *ListBoughtGamesResponse) SetHeaders(v map[string]*string) *ListBoughtGamesResponse {
	s.Headers = v
	return s
}

func (s *ListBoughtGamesResponse) SetBody(v *ListBoughtGamesResponseBody) *ListBoughtGamesResponse {
	s.Body = v
	return s
}

type ListContainerStatusRequest struct {
	GameSessionIdList []*ListContainerStatusRequestGameSessionIdList `json:"GameSessionIdList,omitempty" xml:"GameSessionIdList,omitempty" type:"Repeated"`
}

func (s ListContainerStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContainerStatusRequest) GoString() string {
	return s.String()
}

func (s *ListContainerStatusRequest) SetGameSessionIdList(v []*ListContainerStatusRequestGameSessionIdList) *ListContainerStatusRequest {
	s.GameSessionIdList = v
	return s
}

type ListContainerStatusRequestGameSessionIdList struct {
	GameSessionId *string `json:"GameSessionId,omitempty" xml:"GameSessionId,omitempty"`
}

func (s ListContainerStatusRequestGameSessionIdList) String() string {
	return tea.Prettify(s)
}

func (s ListContainerStatusRequestGameSessionIdList) GoString() string {
	return s.String()
}

func (s *ListContainerStatusRequestGameSessionIdList) SetGameSessionId(v string) *ListContainerStatusRequestGameSessionIdList {
	s.GameSessionId = &v
	return s
}

type ListContainerStatusResponseBody struct {
	DataList  []*ListContainerStatusResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListContainerStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContainerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListContainerStatusResponseBody) SetDataList(v []*ListContainerStatusResponseBodyDataList) *ListContainerStatusResponseBody {
	s.DataList = v
	return s
}

func (s *ListContainerStatusResponseBody) SetRequestId(v string) *ListContainerStatusResponseBody {
	s.RequestId = &v
	return s
}

type ListContainerStatusResponseBodyDataList struct {
	AccountId          *string                                                    `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	ContainerQuitTime  *int64                                                     `json:"ContainerQuitTime,omitempty" xml:"ContainerQuitTime,omitempty"`
	ContainerStartTime *int64                                                     `json:"ContainerStartTime,omitempty" xml:"ContainerStartTime,omitempty"`
	ContainerState     *string                                                    `json:"ContainerState,omitempty" xml:"ContainerState,omitempty"`
	GameId             *string                                                    `json:"GameId,omitempty" xml:"GameId,omitempty"`
	GameSessionId      *string                                                    `json:"GameSessionId,omitempty" xml:"GameSessionId,omitempty"`
	PlayerDetailList   []*ListContainerStatusResponseBodyDataListPlayerDetailList `json:"PlayerDetailList,omitempty" xml:"PlayerDetailList,omitempty" type:"Repeated"`
	ProjectId          *string                                                    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Tags               *string                                                    `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// 系统时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListContainerStatusResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListContainerStatusResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListContainerStatusResponseBodyDataList) SetAccountId(v string) *ListContainerStatusResponseBodyDataList {
	s.AccountId = &v
	return s
}

func (s *ListContainerStatusResponseBodyDataList) SetContainerQuitTime(v int64) *ListContainerStatusResponseBodyDataList {
	s.ContainerQuitTime = &v
	return s
}

func (s *ListContainerStatusResponseBodyDataList) SetContainerStartTime(v int64) *ListContainerStatusResponseBodyDataList {
	s.ContainerStartTime = &v
	return s
}

func (s *ListContainerStatusResponseBodyDataList) SetContainerState(v string) *ListContainerStatusResponseBodyDataList {
	s.ContainerState = &v
	return s
}

func (s *ListContainerStatusResponseBodyDataList) SetGameId(v string) *ListContainerStatusResponseBodyDataList {
	s.GameId = &v
	return s
}

func (s *ListContainerStatusResponseBodyDataList) SetGameSessionId(v string) *ListContainerStatusResponseBodyDataList {
	s.GameSessionId = &v
	return s
}

func (s *ListContainerStatusResponseBodyDataList) SetPlayerDetailList(v []*ListContainerStatusResponseBodyDataListPlayerDetailList) *ListContainerStatusResponseBodyDataList {
	s.PlayerDetailList = v
	return s
}

func (s *ListContainerStatusResponseBodyDataList) SetProjectId(v string) *ListContainerStatusResponseBodyDataList {
	s.ProjectId = &v
	return s
}

func (s *ListContainerStatusResponseBodyDataList) SetTags(v string) *ListContainerStatusResponseBodyDataList {
	s.Tags = &v
	return s
}

func (s *ListContainerStatusResponseBodyDataList) SetTimestamp(v int64) *ListContainerStatusResponseBodyDataList {
	s.Timestamp = &v
	return s
}

type ListContainerStatusResponseBodyDataListPlayerDetailList struct {
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	IsInitiator *bool   `json:"IsInitiator,omitempty" xml:"IsInitiator,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListContainerStatusResponseBodyDataListPlayerDetailList) String() string {
	return tea.Prettify(s)
}

func (s ListContainerStatusResponseBodyDataListPlayerDetailList) GoString() string {
	return s.String()
}

func (s *ListContainerStatusResponseBodyDataListPlayerDetailList) SetAccountId(v string) *ListContainerStatusResponseBodyDataListPlayerDetailList {
	s.AccountId = &v
	return s
}

func (s *ListContainerStatusResponseBodyDataListPlayerDetailList) SetIsInitiator(v bool) *ListContainerStatusResponseBodyDataListPlayerDetailList {
	s.IsInitiator = &v
	return s
}

func (s *ListContainerStatusResponseBodyDataListPlayerDetailList) SetStartTime(v int64) *ListContainerStatusResponseBodyDataListPlayerDetailList {
	s.StartTime = &v
	return s
}

type ListContainerStatusResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListContainerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContainerStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContainerStatusResponse) GoString() string {
	return s.String()
}

func (s *ListContainerStatusResponse) SetHeaders(v map[string]*string) *ListContainerStatusResponse {
	s.Headers = v
	return s
}

func (s *ListContainerStatusResponse) SetBody(v *ListContainerStatusResponseBody) *ListContainerStatusResponse {
	s.Body = v
	return s
}

type ListDeployableInstancesRequest struct {
	// 页码
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页大小
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 游戏版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListDeployableInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeployableInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListDeployableInstancesRequest) SetPageNumber(v int64) *ListDeployableInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeployableInstancesRequest) SetPageSize(v int64) *ListDeployableInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeployableInstancesRequest) SetProjectId(v string) *ListDeployableInstancesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeployableInstancesRequest) SetVersionId(v string) *ListDeployableInstancesRequest {
	s.VersionId = &v
	return s
}

type ListDeployableInstancesResponseBody struct {
	// 数据列表
	DataList []*ListDeployableInstancesResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// MaxResults本次请求所返回的最大记录条数
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TotalCount本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeployableInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeployableInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeployableInstancesResponseBody) SetDataList(v []*ListDeployableInstancesResponseBodyDataList) *ListDeployableInstancesResponseBody {
	s.DataList = v
	return s
}

func (s *ListDeployableInstancesResponseBody) SetPageNumber(v int32) *ListDeployableInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDeployableInstancesResponseBody) SetPageSize(v int32) *ListDeployableInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeployableInstancesResponseBody) SetRequestId(v string) *ListDeployableInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeployableInstancesResponseBody) SetTotalCount(v int64) *ListDeployableInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListDeployableInstancesResponseBodyDataList struct {
	// 实例ID
	CloudGameInstanceId *string `json:"CloudGameInstanceId,omitempty" xml:"CloudGameInstanceId,omitempty"`
	// 实例名称
	CloudGameInstanceName *string `json:"CloudGameInstanceName,omitempty" xml:"CloudGameInstanceName,omitempty"`
}

func (s ListDeployableInstancesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListDeployableInstancesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListDeployableInstancesResponseBodyDataList) SetCloudGameInstanceId(v string) *ListDeployableInstancesResponseBodyDataList {
	s.CloudGameInstanceId = &v
	return s
}

func (s *ListDeployableInstancesResponseBodyDataList) SetCloudGameInstanceName(v string) *ListDeployableInstancesResponseBodyDataList {
	s.CloudGameInstanceName = &v
	return s
}

type ListDeployableInstancesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDeployableInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDeployableInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeployableInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListDeployableInstancesResponse) SetHeaders(v map[string]*string) *ListDeployableInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListDeployableInstancesResponse) SetBody(v *ListDeployableInstancesResponseBody) *ListDeployableInstancesResponse {
	s.Body = v
	return s
}

type ListGameVersionsRequest struct {
	// 游戏ID
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListGameVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGameVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListGameVersionsRequest) SetGameId(v string) *ListGameVersionsRequest {
	s.GameId = &v
	return s
}

func (s *ListGameVersionsRequest) SetMaxResults(v int32) *ListGameVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGameVersionsRequest) SetNextToken(v string) *ListGameVersionsRequest {
	s.NextToken = &v
	return s
}

type ListGameVersionsResponseBody struct {
	// 总记录数
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 数据列表
	DataList []*ListGameVersionsResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// 本次请求所返回的最大记录条数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGameVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGameVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGameVersionsResponseBody) SetCount(v int64) *ListGameVersionsResponseBody {
	s.Count = &v
	return s
}

func (s *ListGameVersionsResponseBody) SetDataList(v []*ListGameVersionsResponseBodyDataList) *ListGameVersionsResponseBody {
	s.DataList = v
	return s
}

func (s *ListGameVersionsResponseBody) SetMaxResults(v int32) *ListGameVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGameVersionsResponseBody) SetNextToken(v string) *ListGameVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGameVersionsResponseBody) SetRequestId(v string) *ListGameVersionsResponseBody {
	s.RequestId = &v
	return s
}

type ListGameVersionsResponseBodyDataList struct {
	// 版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// 版本名称
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// 版本号
	VersionNumber *string `json:"VersionNumber,omitempty" xml:"VersionNumber,omitempty"`
}

func (s ListGameVersionsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListGameVersionsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListGameVersionsResponseBodyDataList) SetVersionId(v string) *ListGameVersionsResponseBodyDataList {
	s.VersionId = &v
	return s
}

func (s *ListGameVersionsResponseBodyDataList) SetVersionName(v string) *ListGameVersionsResponseBodyDataList {
	s.VersionName = &v
	return s
}

func (s *ListGameVersionsResponseBodyDataList) SetVersionNumber(v string) *ListGameVersionsResponseBodyDataList {
	s.VersionNumber = &v
	return s
}

type ListGameVersionsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGameVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGameVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGameVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListGameVersionsResponse) SetHeaders(v map[string]*string) *ListGameVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListGameVersionsResponse) SetBody(v *ListGameVersionsResponseBody) *ListGameVersionsResponse {
	s.Body = v
	return s
}

type ListGamesRequest struct {
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListGamesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGamesRequest) GoString() string {
	return s.String()
}

func (s *ListGamesRequest) SetMaxResults(v int32) *ListGamesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGamesRequest) SetNextToken(v string) *ListGamesRequest {
	s.NextToken = &v
	return s
}

type ListGamesResponseBody struct {
	// 总记录数
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 数据列表
	DataList []*ListGamesResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// 本次请求所返回的最大记录条数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGamesResponseBody) SetCount(v int64) *ListGamesResponseBody {
	s.Count = &v
	return s
}

func (s *ListGamesResponseBody) SetDataList(v []*ListGamesResponseBodyDataList) *ListGamesResponseBody {
	s.DataList = v
	return s
}

func (s *ListGamesResponseBody) SetMaxResults(v int32) *ListGamesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGamesResponseBody) SetNextToken(v string) *ListGamesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGamesResponseBody) SetRequestId(v string) *ListGamesResponseBody {
	s.RequestId = &v
	return s
}

type ListGamesResponseBodyDataList struct {
	// 游戏ID
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	// 游戏名称
	GameName *string `json:"GameName,omitempty" xml:"GameName,omitempty"`
	// 平台类型
	PlatformType *int64 `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
}

func (s ListGamesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListGamesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListGamesResponseBodyDataList) SetGameId(v string) *ListGamesResponseBodyDataList {
	s.GameId = &v
	return s
}

func (s *ListGamesResponseBodyDataList) SetGameName(v string) *ListGamesResponseBodyDataList {
	s.GameName = &v
	return s
}

func (s *ListGamesResponseBodyDataList) SetPlatformType(v int64) *ListGamesResponseBodyDataList {
	s.PlatformType = &v
	return s
}

type ListGamesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGamesResponse) GoString() string {
	return s.String()
}

func (s *ListGamesResponse) SetHeaders(v map[string]*string) *ListGamesResponse {
	s.Headers = v
	return s
}

func (s *ListGamesResponse) SetBody(v *ListGamesResponseBody) *ListGamesResponse {
	s.Body = v
	return s
}

type ListHistoryContainerStatusRequest struct {
	// 结束时间（Linux时间戳，单位毫秒）
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 上一个游戏会话ID
	LastGameSessionId *string `json:"LastGameSessionId,omitempty" xml:"LastGameSessionId,omitempty"`
	// 每页数量
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 开始时间（Linux时间戳，单位毫秒）
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListHistoryContainerStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHistoryContainerStatusRequest) GoString() string {
	return s.String()
}

func (s *ListHistoryContainerStatusRequest) SetEndTime(v int64) *ListHistoryContainerStatusRequest {
	s.EndTime = &v
	return s
}

func (s *ListHistoryContainerStatusRequest) SetLastGameSessionId(v string) *ListHistoryContainerStatusRequest {
	s.LastGameSessionId = &v
	return s
}

func (s *ListHistoryContainerStatusRequest) SetPageSize(v int64) *ListHistoryContainerStatusRequest {
	s.PageSize = &v
	return s
}

func (s *ListHistoryContainerStatusRequest) SetProjectId(v string) *ListHistoryContainerStatusRequest {
	s.ProjectId = &v
	return s
}

func (s *ListHistoryContainerStatusRequest) SetStartTime(v int64) *ListHistoryContainerStatusRequest {
	s.StartTime = &v
	return s
}

type ListHistoryContainerStatusResponseBody struct {
	// 容器状态信息集合
	DataList []*ListHistoryContainerStatusResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHistoryContainerStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHistoryContainerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListHistoryContainerStatusResponseBody) SetDataList(v []*ListHistoryContainerStatusResponseBodyDataList) *ListHistoryContainerStatusResponseBody {
	s.DataList = v
	return s
}

func (s *ListHistoryContainerStatusResponseBody) SetRequestId(v string) *ListHistoryContainerStatusResponseBody {
	s.RequestId = &v
	return s
}

type ListHistoryContainerStatusResponseBodyDataList struct {
	// 主机账号ID
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// 容器退出时间（Linux时间戳，单位毫秒）
	ContainerQuitTime *int64 `json:"ContainerQuitTime,omitempty" xml:"ContainerQuitTime,omitempty"`
	// 容器启动时间（Linux时间戳，单位毫秒）
	ContainerStartTime *int64 `json:"ContainerStartTime,omitempty" xml:"ContainerStartTime,omitempty"`
	// 容器状态
	ContainerState *string `json:"ContainerState,omitempty" xml:"ContainerState,omitempty"`
	// 游戏ID
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	// 游戏会话ID
	GameSessionId *string `json:"GameSessionId,omitempty" xml:"GameSessionId,omitempty"`
	// 玩家信息集合
	PlayerDetailList []*ListHistoryContainerStatusResponseBodyDataListPlayerDetailList `json:"PlayerDetailList,omitempty" xml:"PlayerDetailList,omitempty" type:"Repeated"`
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 自定义标识
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// 系统时间戳
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListHistoryContainerStatusResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListHistoryContainerStatusResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListHistoryContainerStatusResponseBodyDataList) SetAccountId(v string) *ListHistoryContainerStatusResponseBodyDataList {
	s.AccountId = &v
	return s
}

func (s *ListHistoryContainerStatusResponseBodyDataList) SetContainerQuitTime(v int64) *ListHistoryContainerStatusResponseBodyDataList {
	s.ContainerQuitTime = &v
	return s
}

func (s *ListHistoryContainerStatusResponseBodyDataList) SetContainerStartTime(v int64) *ListHistoryContainerStatusResponseBodyDataList {
	s.ContainerStartTime = &v
	return s
}

func (s *ListHistoryContainerStatusResponseBodyDataList) SetContainerState(v string) *ListHistoryContainerStatusResponseBodyDataList {
	s.ContainerState = &v
	return s
}

func (s *ListHistoryContainerStatusResponseBodyDataList) SetGameId(v string) *ListHistoryContainerStatusResponseBodyDataList {
	s.GameId = &v
	return s
}

func (s *ListHistoryContainerStatusResponseBodyDataList) SetGameSessionId(v string) *ListHistoryContainerStatusResponseBodyDataList {
	s.GameSessionId = &v
	return s
}

func (s *ListHistoryContainerStatusResponseBodyDataList) SetPlayerDetailList(v []*ListHistoryContainerStatusResponseBodyDataListPlayerDetailList) *ListHistoryContainerStatusResponseBodyDataList {
	s.PlayerDetailList = v
	return s
}

func (s *ListHistoryContainerStatusResponseBodyDataList) SetProjectId(v string) *ListHistoryContainerStatusResponseBodyDataList {
	s.ProjectId = &v
	return s
}

func (s *ListHistoryContainerStatusResponseBodyDataList) SetTags(v string) *ListHistoryContainerStatusResponseBodyDataList {
	s.Tags = &v
	return s
}

func (s *ListHistoryContainerStatusResponseBodyDataList) SetTimestamp(v int64) *ListHistoryContainerStatusResponseBodyDataList {
	s.Timestamp = &v
	return s
}

type ListHistoryContainerStatusResponseBodyDataListPlayerDetailList struct {
	// 账号ID
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// 是否主机
	IsInitiator *bool `json:"IsInitiator,omitempty" xml:"IsInitiator,omitempty"`
	// 玩家进入游戏时间
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListHistoryContainerStatusResponseBodyDataListPlayerDetailList) String() string {
	return tea.Prettify(s)
}

func (s ListHistoryContainerStatusResponseBodyDataListPlayerDetailList) GoString() string {
	return s.String()
}

func (s *ListHistoryContainerStatusResponseBodyDataListPlayerDetailList) SetAccountId(v string) *ListHistoryContainerStatusResponseBodyDataListPlayerDetailList {
	s.AccountId = &v
	return s
}

func (s *ListHistoryContainerStatusResponseBodyDataListPlayerDetailList) SetIsInitiator(v bool) *ListHistoryContainerStatusResponseBodyDataListPlayerDetailList {
	s.IsInitiator = &v
	return s
}

func (s *ListHistoryContainerStatusResponseBodyDataListPlayerDetailList) SetStartTime(v int64) *ListHistoryContainerStatusResponseBodyDataListPlayerDetailList {
	s.StartTime = &v
	return s
}

type ListHistoryContainerStatusResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHistoryContainerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHistoryContainerStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHistoryContainerStatusResponse) GoString() string {
	return s.String()
}

func (s *ListHistoryContainerStatusResponse) SetHeaders(v map[string]*string) *ListHistoryContainerStatusResponse {
	s.Headers = v
	return s
}

func (s *ListHistoryContainerStatusResponse) SetBody(v *ListHistoryContainerStatusResponseBody) *ListHistoryContainerStatusResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetMaxResults(v int32) *ListProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsRequest) SetNextToken(v string) *ListProjectsRequest {
	s.NextToken = &v
	return s
}

type ListProjectsResponseBody struct {
	// 总记录数
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 数据列表
	DataList []*ListProjectsResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// 本次请求所返回的最大记录条数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetCount(v int64) *ListProjectsResponseBody {
	s.Count = &v
	return s
}

func (s *ListProjectsResponseBody) SetDataList(v []*ListProjectsResponseBodyDataList) *ListProjectsResponseBody {
	s.DataList = v
	return s
}

func (s *ListProjectsResponseBody) SetMaxResults(v int32) *ListProjectsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsResponseBody) SetNextToken(v string) *ListProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectsResponseBodyDataList struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ListProjectsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataList) SetProjectId(v string) *ListProjectsResponseBodyDataList {
	s.ProjectId = &v
	return s
}

func (s *ListProjectsResponseBodyDataList) SetProjectName(v string) *ListProjectsResponseBodyDataList {
	s.ProjectName = &v
	return s
}

type ListProjectsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type QueryGameRequest struct {
	PageNo    *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TenantId  *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryGameRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGameRequest) GoString() string {
	return s.String()
}

func (s *QueryGameRequest) SetPageNo(v int32) *QueryGameRequest {
	s.PageNo = &v
	return s
}

func (s *QueryGameRequest) SetPageSize(v int32) *QueryGameRequest {
	s.PageSize = &v
	return s
}

func (s *QueryGameRequest) SetProjectId(v int64) *QueryGameRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryGameRequest) SetTenantId(v int64) *QueryGameRequest {
	s.TenantId = &v
	return s
}

type QueryGameResponseBody struct {
	Data       []*QueryGameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int32                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryGameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGameResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGameResponseBody) SetData(v []*QueryGameResponseBodyData) *QueryGameResponseBody {
	s.Data = v
	return s
}

func (s *QueryGameResponseBody) SetPageNumber(v int32) *QueryGameResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryGameResponseBody) SetPageSize(v int32) *QueryGameResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryGameResponseBody) SetRequestId(v string) *QueryGameResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGameResponseBody) SetTotalCount(v int32) *QueryGameResponseBody {
	s.TotalCount = &v
	return s
}

type QueryGameResponseBodyData struct {
	GameId    *int64  `json:"GameId,omitempty" xml:"GameId,omitempty"`
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TenantId  *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s QueryGameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryGameResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryGameResponseBodyData) SetGameId(v int64) *QueryGameResponseBodyData {
	s.GameId = &v
	return s
}

func (s *QueryGameResponseBodyData) SetGmtCreate(v string) *QueryGameResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryGameResponseBodyData) SetName(v string) *QueryGameResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryGameResponseBodyData) SetProjectId(v int64) *QueryGameResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *QueryGameResponseBodyData) SetTenantId(v int64) *QueryGameResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *QueryGameResponseBodyData) SetVersion(v string) *QueryGameResponseBodyData {
	s.Version = &v
	return s
}

type QueryGameResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryGameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryGameResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGameResponse) GoString() string {
	return s.String()
}

func (s *QueryGameResponse) SetHeaders(v map[string]*string) *QueryGameResponse {
	s.Headers = v
	return s
}

func (s *QueryGameResponse) SetBody(v *QueryGameResponseBody) *QueryGameResponse {
	s.Body = v
	return s
}

type QueryItemsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsRequest) GoString() string {
	return s.String()
}

func (s *QueryItemsRequest) SetPageNumber(v int32) *QueryItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryItemsRequest) SetPageSize(v int32) *QueryItemsRequest {
	s.PageSize = &v
	return s
}

type QueryItemsResponseBody struct {
	Data           *QueryItemsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int64                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryItemsResponseBody) SetData(v *QueryItemsResponseBodyData) *QueryItemsResponseBody {
	s.Data = v
	return s
}

func (s *QueryItemsResponseBody) SetHttpStatusCode(v int64) *QueryItemsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryItemsResponseBody) SetRequestId(v string) *QueryItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryItemsResponseBody) SetSuccess(v bool) *QueryItemsResponseBody {
	s.Success = &v
	return s
}

type QueryItemsResponseBodyData struct {
	Items      []*QueryItemsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryItemsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryItemsResponseBodyData) SetItems(v []*QueryItemsResponseBodyDataItems) *QueryItemsResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryItemsResponseBodyData) SetPageNumber(v int32) *QueryItemsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *QueryItemsResponseBodyData) SetPageSize(v int32) *QueryItemsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryItemsResponseBodyData) SetTotalCount(v int64) *QueryItemsResponseBodyData {
	s.TotalCount = &v
	return s
}

type QueryItemsResponseBodyDataItems struct {
	CategoryId  *int64                                  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreateTime  *int64                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	Games       []*QueryItemsResponseBodyDataItemsGames `json:"Games,omitempty" xml:"Games,omitempty" type:"Repeated"`
	ItemId      *string                                 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ModifyTime  *int64                                  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	OriginPrice *int64                                  `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	SalePrice   *int64                                  `json:"SalePrice,omitempty" xml:"SalePrice,omitempty"`
	SellerId    *string                                 `json:"SellerId,omitempty" xml:"SellerId,omitempty"`
	Skus        []*QueryItemsResponseBodyDataItemsSkus  `json:"Skus,omitempty" xml:"Skus,omitempty" type:"Repeated"`
	Status      *int32                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Supplier    *string                                 `json:"Supplier,omitempty" xml:"Supplier,omitempty"`
	Title       *string                                 `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s QueryItemsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryItemsResponseBodyDataItems) SetCategoryId(v int64) *QueryItemsResponseBodyDataItems {
	s.CategoryId = &v
	return s
}

func (s *QueryItemsResponseBodyDataItems) SetCreateTime(v int64) *QueryItemsResponseBodyDataItems {
	s.CreateTime = &v
	return s
}

func (s *QueryItemsResponseBodyDataItems) SetDescription(v string) *QueryItemsResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *QueryItemsResponseBodyDataItems) SetGames(v []*QueryItemsResponseBodyDataItemsGames) *QueryItemsResponseBodyDataItems {
	s.Games = v
	return s
}

func (s *QueryItemsResponseBodyDataItems) SetItemId(v string) *QueryItemsResponseBodyDataItems {
	s.ItemId = &v
	return s
}

func (s *QueryItemsResponseBodyDataItems) SetModifyTime(v int64) *QueryItemsResponseBodyDataItems {
	s.ModifyTime = &v
	return s
}

func (s *QueryItemsResponseBodyDataItems) SetOriginPrice(v int64) *QueryItemsResponseBodyDataItems {
	s.OriginPrice = &v
	return s
}

func (s *QueryItemsResponseBodyDataItems) SetSalePrice(v int64) *QueryItemsResponseBodyDataItems {
	s.SalePrice = &v
	return s
}

func (s *QueryItemsResponseBodyDataItems) SetSellerId(v string) *QueryItemsResponseBodyDataItems {
	s.SellerId = &v
	return s
}

func (s *QueryItemsResponseBodyDataItems) SetSkus(v []*QueryItemsResponseBodyDataItemsSkus) *QueryItemsResponseBodyDataItems {
	s.Skus = v
	return s
}

func (s *QueryItemsResponseBodyDataItems) SetStatus(v int32) *QueryItemsResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *QueryItemsResponseBodyDataItems) SetSupplier(v string) *QueryItemsResponseBodyDataItems {
	s.Supplier = &v
	return s
}

func (s *QueryItemsResponseBodyDataItems) SetTitle(v string) *QueryItemsResponseBodyDataItems {
	s.Title = &v
	return s
}

type QueryItemsResponseBodyDataItemsGames struct {
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryItemsResponseBodyDataItemsGames) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponseBodyDataItemsGames) GoString() string {
	return s.String()
}

func (s *QueryItemsResponseBodyDataItemsGames) SetGameId(v string) *QueryItemsResponseBodyDataItemsGames {
	s.GameId = &v
	return s
}

func (s *QueryItemsResponseBodyDataItemsGames) SetName(v string) *QueryItemsResponseBodyDataItemsGames {
	s.Name = &v
	return s
}

type QueryItemsResponseBodyDataItemsSkus struct {
	CreateTime  *int64                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ItemId      *string                                         `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ModifyTime  *int64                                          `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	OriginPrice *int64                                          `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	SalePrice   *int64                                          `json:"SalePrice,omitempty" xml:"SalePrice,omitempty"`
	SaleProps   []*QueryItemsResponseBodyDataItemsSkusSaleProps `json:"SaleProps,omitempty" xml:"SaleProps,omitempty" type:"Repeated"`
	SkuId       *string                                         `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Status      *int32                                          `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryItemsResponseBodyDataItemsSkus) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponseBodyDataItemsSkus) GoString() string {
	return s.String()
}

func (s *QueryItemsResponseBodyDataItemsSkus) SetCreateTime(v int64) *QueryItemsResponseBodyDataItemsSkus {
	s.CreateTime = &v
	return s
}

func (s *QueryItemsResponseBodyDataItemsSkus) SetItemId(v string) *QueryItemsResponseBodyDataItemsSkus {
	s.ItemId = &v
	return s
}

func (s *QueryItemsResponseBodyDataItemsSkus) SetModifyTime(v int64) *QueryItemsResponseBodyDataItemsSkus {
	s.ModifyTime = &v
	return s
}

func (s *QueryItemsResponseBodyDataItemsSkus) SetOriginPrice(v int64) *QueryItemsResponseBodyDataItemsSkus {
	s.OriginPrice = &v
	return s
}

func (s *QueryItemsResponseBodyDataItemsSkus) SetSalePrice(v int64) *QueryItemsResponseBodyDataItemsSkus {
	s.SalePrice = &v
	return s
}

func (s *QueryItemsResponseBodyDataItemsSkus) SetSaleProps(v []*QueryItemsResponseBodyDataItemsSkusSaleProps) *QueryItemsResponseBodyDataItemsSkus {
	s.SaleProps = v
	return s
}

func (s *QueryItemsResponseBodyDataItemsSkus) SetSkuId(v string) *QueryItemsResponseBodyDataItemsSkus {
	s.SkuId = &v
	return s
}

func (s *QueryItemsResponseBodyDataItemsSkus) SetStatus(v int32) *QueryItemsResponseBodyDataItemsSkus {
	s.Status = &v
	return s
}

type QueryItemsResponseBodyDataItemsSkusSaleProps struct {
	PropertyId   *int64  `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueId      *int64  `json:"ValueId,omitempty" xml:"ValueId,omitempty"`
}

func (s QueryItemsResponseBodyDataItemsSkusSaleProps) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponseBodyDataItemsSkusSaleProps) GoString() string {
	return s.String()
}

func (s *QueryItemsResponseBodyDataItemsSkusSaleProps) SetPropertyId(v int64) *QueryItemsResponseBodyDataItemsSkusSaleProps {
	s.PropertyId = &v
	return s
}

func (s *QueryItemsResponseBodyDataItemsSkusSaleProps) SetPropertyName(v string) *QueryItemsResponseBodyDataItemsSkusSaleProps {
	s.PropertyName = &v
	return s
}

func (s *QueryItemsResponseBodyDataItemsSkusSaleProps) SetValue(v string) *QueryItemsResponseBodyDataItemsSkusSaleProps {
	s.Value = &v
	return s
}

func (s *QueryItemsResponseBodyDataItemsSkusSaleProps) SetValueId(v int64) *QueryItemsResponseBodyDataItemsSkusSaleProps {
	s.ValueId = &v
	return s
}

type QueryItemsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponse) GoString() string {
	return s.String()
}

func (s *QueryItemsResponse) SetHeaders(v map[string]*string) *QueryItemsResponse {
	s.Headers = v
	return s
}

func (s *QueryItemsResponse) SetBody(v *QueryItemsResponseBody) *QueryItemsResponse {
	s.Body = v
	return s
}

type QueryOrderRequest struct {
	AccountDomain  *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
	BuyerAccountId *string `json:"BuyerAccountId,omitempty" xml:"BuyerAccountId,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s QueryOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderRequest) SetAccountDomain(v string) *QueryOrderRequest {
	s.AccountDomain = &v
	return s
}

func (s *QueryOrderRequest) SetBuyerAccountId(v string) *QueryOrderRequest {
	s.BuyerAccountId = &v
	return s
}

func (s *QueryOrderRequest) SetOrderId(v string) *QueryOrderRequest {
	s.OrderId = &v
	return s
}

type QueryOrderResponseBody struct {
	Data           *QueryOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DeliveryStatus *string                     `json:"DeliveryStatus,omitempty" xml:"DeliveryStatus,omitempty"`
	RefundStatus   *string                     `json:"RefundStatus,omitempty" xml:"RefundStatus,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBody) SetData(v *QueryOrderResponseBodyData) *QueryOrderResponseBody {
	s.Data = v
	return s
}

func (s *QueryOrderResponseBody) SetDeliveryStatus(v string) *QueryOrderResponseBody {
	s.DeliveryStatus = &v
	return s
}

func (s *QueryOrderResponseBody) SetRefundStatus(v string) *QueryOrderResponseBody {
	s.RefundStatus = &v
	return s
}

func (s *QueryOrderResponseBody) SetRequestId(v string) *QueryOrderResponseBody {
	s.RequestId = &v
	return s
}

type QueryOrderResponseBodyData struct {
	AccountDomain     *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
	Amount            *int64  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	ApplyDeliveryTime *int64  `json:"ApplyDeliveryTime,omitempty" xml:"ApplyDeliveryTime,omitempty"`
	AutoUnlockTime    *int64  `json:"AutoUnlockTime,omitempty" xml:"AutoUnlockTime,omitempty"`
	BuyerAccountId    *string `json:"BuyerAccountId,omitempty" xml:"BuyerAccountId,omitempty"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FinishTime        *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	ItemId            *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	OrderId           *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OriginPrice       *int64  `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	SettlementPrice   *int64  `json:"SettlementPrice,omitempty" xml:"SettlementPrice,omitempty"`
	SkuId             *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseBodyData) SetAccountDomain(v string) *QueryOrderResponseBodyData {
	s.AccountDomain = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetAmount(v int64) *QueryOrderResponseBodyData {
	s.Amount = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetApplyDeliveryTime(v int64) *QueryOrderResponseBodyData {
	s.ApplyDeliveryTime = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetAutoUnlockTime(v int64) *QueryOrderResponseBodyData {
	s.AutoUnlockTime = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetBuyerAccountId(v string) *QueryOrderResponseBodyData {
	s.BuyerAccountId = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetCreateTime(v int64) *QueryOrderResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetFinishTime(v int64) *QueryOrderResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetItemId(v string) *QueryOrderResponseBodyData {
	s.ItemId = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetOrderId(v string) *QueryOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetOriginPrice(v int64) *QueryOrderResponseBodyData {
	s.OriginPrice = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetSettlementPrice(v int64) *QueryOrderResponseBodyData {
	s.SettlementPrice = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetSkuId(v string) *QueryOrderResponseBodyData {
	s.SkuId = &v
	return s
}

func (s *QueryOrderResponseBodyData) SetStatus(v string) *QueryOrderResponseBodyData {
	s.Status = &v
	return s
}

type QueryOrderResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderResponse) SetHeaders(v map[string]*string) *QueryOrderResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderResponse) SetBody(v *QueryOrderResponseBody) *QueryOrderResponse {
	s.Body = v
	return s
}

type QueryOutAccountBindStatusRequest struct {
	AccountDomain *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
	AccountId     *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	GameId        *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
}

func (s QueryOutAccountBindStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOutAccountBindStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryOutAccountBindStatusRequest) SetAccountDomain(v string) *QueryOutAccountBindStatusRequest {
	s.AccountDomain = &v
	return s
}

func (s *QueryOutAccountBindStatusRequest) SetAccountId(v string) *QueryOutAccountBindStatusRequest {
	s.AccountId = &v
	return s
}

func (s *QueryOutAccountBindStatusRequest) SetGameId(v string) *QueryOutAccountBindStatusRequest {
	s.GameId = &v
	return s
}

type QueryOutAccountBindStatusResponseBody struct {
	Data      *QueryOutAccountBindStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryOutAccountBindStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOutAccountBindStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOutAccountBindStatusResponseBody) SetData(v *QueryOutAccountBindStatusResponseBodyData) *QueryOutAccountBindStatusResponseBody {
	s.Data = v
	return s
}

func (s *QueryOutAccountBindStatusResponseBody) SetRequestId(v string) *QueryOutAccountBindStatusResponseBody {
	s.RequestId = &v
	return s
}

type QueryOutAccountBindStatusResponseBodyData struct {
	BindStatus *int32 `json:"BindStatus,omitempty" xml:"BindStatus,omitempty"`
}

func (s QueryOutAccountBindStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryOutAccountBindStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOutAccountBindStatusResponseBodyData) SetBindStatus(v int32) *QueryOutAccountBindStatusResponseBodyData {
	s.BindStatus = &v
	return s
}

type QueryOutAccountBindStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryOutAccountBindStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOutAccountBindStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOutAccountBindStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryOutAccountBindStatusResponse) SetHeaders(v map[string]*string) *QueryOutAccountBindStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryOutAccountBindStatusResponse) SetBody(v *QueryOutAccountBindStatusResponseBody) *QueryOutAccountBindStatusResponse {
	s.Body = v
	return s
}

type QueryProjectRequest struct {
	PageNo    *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TenantId  *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectRequest) GoString() string {
	return s.String()
}

func (s *QueryProjectRequest) SetPageNo(v int32) *QueryProjectRequest {
	s.PageNo = &v
	return s
}

func (s *QueryProjectRequest) SetPageSize(v int32) *QueryProjectRequest {
	s.PageSize = &v
	return s
}

func (s *QueryProjectRequest) SetProjectId(v int64) *QueryProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryProjectRequest) SetTenantId(v int64) *QueryProjectRequest {
	s.TenantId = &v
	return s
}

type QueryProjectResponseBody struct {
	Data       []*QueryProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBody) SetData(v []*QueryProjectResponseBodyData) *QueryProjectResponseBody {
	s.Data = v
	return s
}

func (s *QueryProjectResponseBody) SetPageNumber(v int32) *QueryProjectResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryProjectResponseBody) SetPageSize(v int32) *QueryProjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryProjectResponseBody) SetRequestId(v string) *QueryProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProjectResponseBody) SetTotalCount(v int32) *QueryProjectResponseBody {
	s.TotalCount = &v
	return s
}

type QueryProjectResponseBodyData struct {
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantId *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBodyData) SetId(v int64) *QueryProjectResponseBodyData {
	s.Id = &v
	return s
}

func (s *QueryProjectResponseBodyData) SetName(v string) *QueryProjectResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryProjectResponseBodyData) SetTenantId(v int64) *QueryProjectResponseBodyData {
	s.TenantId = &v
	return s
}

type QueryProjectResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponse) GoString() string {
	return s.String()
}

func (s *QueryProjectResponse) SetHeaders(v map[string]*string) *QueryProjectResponse {
	s.Headers = v
	return s
}

func (s *QueryProjectResponse) SetBody(v *QueryProjectResponseBody) *QueryProjectResponse {
	s.Body = v
	return s
}

type QueryTenantRequest struct {
	PageNo   *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Param    *string `json:"Param,omitempty" xml:"Param,omitempty"`
}

func (s QueryTenantRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTenantRequest) GoString() string {
	return s.String()
}

func (s *QueryTenantRequest) SetPageNo(v int32) *QueryTenantRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTenantRequest) SetPageSize(v int32) *QueryTenantRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTenantRequest) SetParam(v string) *QueryTenantRequest {
	s.Param = &v
	return s
}

type QueryTenantResponseBody struct {
	Data       []*QueryTenantResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryTenantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTenantResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTenantResponseBody) SetData(v []*QueryTenantResponseBodyData) *QueryTenantResponseBody {
	s.Data = v
	return s
}

func (s *QueryTenantResponseBody) SetPageNumber(v int32) *QueryTenantResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryTenantResponseBody) SetPageSize(v int32) *QueryTenantResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTenantResponseBody) SetRequestId(v string) *QueryTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTenantResponseBody) SetTotalCount(v int32) *QueryTenantResponseBody {
	s.TotalCount = &v
	return s
}

type QueryTenantResponseBodyData struct {
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantId  *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryTenantResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTenantResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTenantResponseBodyData) SetGmtCreate(v string) *QueryTenantResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryTenantResponseBodyData) SetName(v string) *QueryTenantResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryTenantResponseBodyData) SetTenantId(v int64) *QueryTenantResponseBodyData {
	s.TenantId = &v
	return s
}

type QueryTenantResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTenantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTenantResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTenantResponse) GoString() string {
	return s.String()
}

func (s *QueryTenantResponse) SetHeaders(v map[string]*string) *QueryTenantResponse {
	s.Headers = v
	return s
}

func (s *QueryTenantResponse) SetBody(v *QueryTenantResponseBody) *QueryTenantResponse {
	s.Body = v
	return s
}

type RemoveGameFromProjectRequest struct {
	// 游戏iD
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RemoveGameFromProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGameFromProjectRequest) GoString() string {
	return s.String()
}

func (s *RemoveGameFromProjectRequest) SetGameId(v string) *RemoveGameFromProjectRequest {
	s.GameId = &v
	return s
}

func (s *RemoveGameFromProjectRequest) SetProjectId(v string) *RemoveGameFromProjectRequest {
	s.ProjectId = &v
	return s
}

type RemoveGameFromProjectResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveGameFromProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveGameFromProjectResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveGameFromProjectResponseBody) SetRequestId(v string) *RemoveGameFromProjectResponseBody {
	s.RequestId = &v
	return s
}

type RemoveGameFromProjectResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveGameFromProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveGameFromProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveGameFromProjectResponse) GoString() string {
	return s.String()
}

func (s *RemoveGameFromProjectResponse) SetHeaders(v map[string]*string) *RemoveGameFromProjectResponse {
	s.Headers = v
	return s
}

func (s *RemoveGameFromProjectResponse) SetBody(v *RemoveGameFromProjectResponseBody) *RemoveGameFromProjectResponse {
	s.Body = v
	return s
}

type SkipTrialPolicyRequest struct {
	GameSessionId *string `json:"GameSessionId,omitempty" xml:"GameSessionId,omitempty"`
}

func (s SkipTrialPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SkipTrialPolicyRequest) GoString() string {
	return s.String()
}

func (s *SkipTrialPolicyRequest) SetGameSessionId(v string) *SkipTrialPolicyRequest {
	s.GameSessionId = &v
	return s
}

type SkipTrialPolicyResponseBody struct {
	Data      *SkipTrialPolicyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SkipTrialPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SkipTrialPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SkipTrialPolicyResponseBody) SetData(v *SkipTrialPolicyResponseBodyData) *SkipTrialPolicyResponseBody {
	s.Data = v
	return s
}

func (s *SkipTrialPolicyResponseBody) SetRequestId(v string) *SkipTrialPolicyResponseBody {
	s.RequestId = &v
	return s
}

type SkipTrialPolicyResponseBodyData struct {
	SkipResult *int32 `json:"SkipResult,omitempty" xml:"SkipResult,omitempty"`
}

func (s SkipTrialPolicyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SkipTrialPolicyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SkipTrialPolicyResponseBodyData) SetSkipResult(v int32) *SkipTrialPolicyResponseBodyData {
	s.SkipResult = &v
	return s
}

type SkipTrialPolicyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SkipTrialPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SkipTrialPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SkipTrialPolicyResponse) GoString() string {
	return s.String()
}

func (s *SkipTrialPolicyResponse) SetHeaders(v map[string]*string) *SkipTrialPolicyResponse {
	s.Headers = v
	return s
}

func (s *SkipTrialPolicyResponse) SetBody(v *SkipTrialPolicyResponseBody) *SkipTrialPolicyResponse {
	s.Body = v
	return s
}

type StopGameSessionRequest struct {
	AccessKey   *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	BizParam    *string `json:"BizParam,omitempty" xml:"BizParam,omitempty"`
	GameId      *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	GameSession *string `json:"GameSession,omitempty" xml:"GameSession,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StopGameSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s StopGameSessionRequest) GoString() string {
	return s.String()
}

func (s *StopGameSessionRequest) SetAccessKey(v string) *StopGameSessionRequest {
	s.AccessKey = &v
	return s
}

func (s *StopGameSessionRequest) SetBizParam(v string) *StopGameSessionRequest {
	s.BizParam = &v
	return s
}

func (s *StopGameSessionRequest) SetGameId(v string) *StopGameSessionRequest {
	s.GameId = &v
	return s
}

func (s *StopGameSessionRequest) SetGameSession(v string) *StopGameSessionRequest {
	s.GameSession = &v
	return s
}

func (s *StopGameSessionRequest) SetReason(v string) *StopGameSessionRequest {
	s.Reason = &v
	return s
}

func (s *StopGameSessionRequest) SetUserId(v string) *StopGameSessionRequest {
	s.UserId = &v
	return s
}

type StopGameSessionResponseBody struct {
	GameId      *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	GameSession *string `json:"GameSession,omitempty" xml:"GameSession,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	QueueCode   *int32  `json:"QueueCode,omitempty" xml:"QueueCode,omitempty"`
	QueueState  *int32  `json:"QueueState,omitempty" xml:"QueueState,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopGameSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopGameSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StopGameSessionResponseBody) SetGameId(v string) *StopGameSessionResponseBody {
	s.GameId = &v
	return s
}

func (s *StopGameSessionResponseBody) SetGameSession(v string) *StopGameSessionResponseBody {
	s.GameSession = &v
	return s
}

func (s *StopGameSessionResponseBody) SetMessage(v string) *StopGameSessionResponseBody {
	s.Message = &v
	return s
}

func (s *StopGameSessionResponseBody) SetQueueCode(v int32) *StopGameSessionResponseBody {
	s.QueueCode = &v
	return s
}

func (s *StopGameSessionResponseBody) SetQueueState(v int32) *StopGameSessionResponseBody {
	s.QueueState = &v
	return s
}

func (s *StopGameSessionResponseBody) SetRequestId(v string) *StopGameSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopGameSessionResponseBody) SetSuccess(v bool) *StopGameSessionResponseBody {
	s.Success = &v
	return s
}

type StopGameSessionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopGameSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopGameSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s StopGameSessionResponse) GoString() string {
	return s.String()
}

func (s *StopGameSessionResponse) SetHeaders(v map[string]*string) *StopGameSessionResponse {
	s.Headers = v
	return s
}

func (s *StopGameSessionResponse) SetBody(v *StopGameSessionResponseBody) *StopGameSessionResponse {
	s.Body = v
	return s
}

type SubmitDeploymentRequest struct {
	// 实例ID列表
	CloudGameInstanceIds *string `json:"CloudGameInstanceIds,omitempty" xml:"CloudGameInstanceIds,omitempty"`
	// 游戏iD
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	// 操作类型
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 游戏版本ID
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s SubmitDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDeploymentRequest) GoString() string {
	return s.String()
}

func (s *SubmitDeploymentRequest) SetCloudGameInstanceIds(v string) *SubmitDeploymentRequest {
	s.CloudGameInstanceIds = &v
	return s
}

func (s *SubmitDeploymentRequest) SetGameId(v string) *SubmitDeploymentRequest {
	s.GameId = &v
	return s
}

func (s *SubmitDeploymentRequest) SetOperationType(v string) *SubmitDeploymentRequest {
	s.OperationType = &v
	return s
}

func (s *SubmitDeploymentRequest) SetProjectId(v string) *SubmitDeploymentRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitDeploymentRequest) SetVersionId(v string) *SubmitDeploymentRequest {
	s.VersionId = &v
	return s
}

type SubmitDeploymentResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDeploymentResponseBody) SetRequestId(v string) *SubmitDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDeploymentResponseBody) SetTaskId(v string) *SubmitDeploymentResponseBody {
	s.TaskId = &v
	return s
}

type SubmitDeploymentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDeploymentResponse) GoString() string {
	return s.String()
}

func (s *SubmitDeploymentResponse) SetHeaders(v map[string]*string) *SubmitDeploymentResponse {
	s.Headers = v
	return s
}

func (s *SubmitDeploymentResponse) SetBody(v *SubmitDeploymentResponseBody) *SubmitDeploymentResponse {
	s.Body = v
	return s
}

type SubmitInternalPurchaseChargeDataRequest struct {
	ActiveUserRetentionRateOneDay    *float32 `json:"ActiveUserRetentionRateOneDay,omitempty" xml:"ActiveUserRetentionRateOneDay,omitempty"`
	ActiveUserRetentionRateSevenDay  *float32 `json:"ActiveUserRetentionRateSevenDay,omitempty" xml:"ActiveUserRetentionRateSevenDay,omitempty"`
	ActiveUserRetentionRateThirtyDay *float32 `json:"ActiveUserRetentionRateThirtyDay,omitempty" xml:"ActiveUserRetentionRateThirtyDay,omitempty"`
	Arpu                             *float32 `json:"Arpu,omitempty" xml:"Arpu,omitempty"`
	ChargeDate                       *string  `json:"ChargeDate,omitempty" xml:"ChargeDate,omitempty"`
	Dau                              *int64   `json:"Dau,omitempty" xml:"Dau,omitempty"`
	GameId                           *string  `json:"GameId,omitempty" xml:"GameId,omitempty"`
	Mau                              *int64   `json:"Mau,omitempty" xml:"Mau,omitempty"`
	NewUserRetentionRateOneDay       *float32 `json:"NewUserRetentionRateOneDay,omitempty" xml:"NewUserRetentionRateOneDay,omitempty"`
	NewUserRetentionRateSevenDay     *float32 `json:"NewUserRetentionRateSevenDay,omitempty" xml:"NewUserRetentionRateSevenDay,omitempty"`
	NewUserRetentionRateThirtyDay    *float32 `json:"NewUserRetentionRateThirtyDay,omitempty" xml:"NewUserRetentionRateThirtyDay,omitempty"`
	PaymentConversionRate            *float32 `json:"PaymentConversionRate,omitempty" xml:"PaymentConversionRate,omitempty"`
	PlayTimeAverageOneDay            *float32 `json:"PlayTimeAverageOneDay,omitempty" xml:"PlayTimeAverageOneDay,omitempty"`
	PlayTimeAverageThirtyDay         *float32 `json:"PlayTimeAverageThirtyDay,omitempty" xml:"PlayTimeAverageThirtyDay,omitempty"`
	PlayTimeNinetyPointsOneDay       *float32 `json:"PlayTimeNinetyPointsOneDay,omitempty" xml:"PlayTimeNinetyPointsOneDay,omitempty"`
	PlayTimeNinetyPointsThirtyDay    *float32 `json:"PlayTimeNinetyPointsThirtyDay,omitempty" xml:"PlayTimeNinetyPointsThirtyDay,omitempty"`
	PlayTimeRangeOneDay              *string  `json:"PlayTimeRangeOneDay,omitempty" xml:"PlayTimeRangeOneDay,omitempty"`
	PlayTimeRangeThirtyDay           *string  `json:"PlayTimeRangeThirtyDay,omitempty" xml:"PlayTimeRangeThirtyDay,omitempty"`
	UserActivationRate               *float32 `json:"UserActivationRate,omitempty" xml:"UserActivationRate,omitempty"`
}

func (s SubmitInternalPurchaseChargeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseChargeDataRequest) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetActiveUserRetentionRateOneDay(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.ActiveUserRetentionRateOneDay = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetActiveUserRetentionRateSevenDay(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.ActiveUserRetentionRateSevenDay = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetActiveUserRetentionRateThirtyDay(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.ActiveUserRetentionRateThirtyDay = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetArpu(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.Arpu = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetChargeDate(v string) *SubmitInternalPurchaseChargeDataRequest {
	s.ChargeDate = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetDau(v int64) *SubmitInternalPurchaseChargeDataRequest {
	s.Dau = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetGameId(v string) *SubmitInternalPurchaseChargeDataRequest {
	s.GameId = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetMau(v int64) *SubmitInternalPurchaseChargeDataRequest {
	s.Mau = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetNewUserRetentionRateOneDay(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.NewUserRetentionRateOneDay = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetNewUserRetentionRateSevenDay(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.NewUserRetentionRateSevenDay = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetNewUserRetentionRateThirtyDay(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.NewUserRetentionRateThirtyDay = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetPaymentConversionRate(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.PaymentConversionRate = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetPlayTimeAverageOneDay(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.PlayTimeAverageOneDay = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetPlayTimeAverageThirtyDay(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.PlayTimeAverageThirtyDay = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetPlayTimeNinetyPointsOneDay(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.PlayTimeNinetyPointsOneDay = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetPlayTimeNinetyPointsThirtyDay(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.PlayTimeNinetyPointsThirtyDay = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetPlayTimeRangeOneDay(v string) *SubmitInternalPurchaseChargeDataRequest {
	s.PlayTimeRangeOneDay = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetPlayTimeRangeThirtyDay(v string) *SubmitInternalPurchaseChargeDataRequest {
	s.PlayTimeRangeThirtyDay = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataRequest) SetUserActivationRate(v float32) *SubmitInternalPurchaseChargeDataRequest {
	s.UserActivationRate = &v
	return s
}

type SubmitInternalPurchaseChargeDataResponseBody struct {
	Data      *SubmitInternalPurchaseChargeDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitInternalPurchaseChargeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseChargeDataResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseChargeDataResponseBody) SetData(v *SubmitInternalPurchaseChargeDataResponseBodyData) *SubmitInternalPurchaseChargeDataResponseBody {
	s.Data = v
	return s
}

func (s *SubmitInternalPurchaseChargeDataResponseBody) SetRequestId(v string) *SubmitInternalPurchaseChargeDataResponseBody {
	s.RequestId = &v
	return s
}

type SubmitInternalPurchaseChargeDataResponseBodyData struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status  *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitInternalPurchaseChargeDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseChargeDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseChargeDataResponseBodyData) SetMessage(v string) *SubmitInternalPurchaseChargeDataResponseBodyData {
	s.Message = &v
	return s
}

func (s *SubmitInternalPurchaseChargeDataResponseBodyData) SetStatus(v int32) *SubmitInternalPurchaseChargeDataResponseBodyData {
	s.Status = &v
	return s
}

type SubmitInternalPurchaseChargeDataResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitInternalPurchaseChargeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitInternalPurchaseChargeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseChargeDataResponse) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseChargeDataResponse) SetHeaders(v map[string]*string) *SubmitInternalPurchaseChargeDataResponse {
	s.Headers = v
	return s
}

func (s *SubmitInternalPurchaseChargeDataResponse) SetBody(v *SubmitInternalPurchaseChargeDataResponseBody) *SubmitInternalPurchaseChargeDataResponse {
	s.Body = v
	return s
}

type SubmitInternalPurchaseOrdersRequest struct {
	OrderList []*SubmitInternalPurchaseOrdersRequestOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Repeated"`
}

func (s SubmitInternalPurchaseOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseOrdersRequest) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseOrdersRequest) SetOrderList(v []*SubmitInternalPurchaseOrdersRequestOrderList) *SubmitInternalPurchaseOrdersRequest {
	s.OrderList = v
	return s
}

type SubmitInternalPurchaseOrdersRequestOrderList struct {
	BatchNumber *string `json:"BatchNumber,omitempty" xml:"BatchNumber,omitempty"`
	FinalPrice  *int64  `json:"FinalPrice,omitempty" xml:"FinalPrice,omitempty"`
	FinishTime  *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	GameId      *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	OrderId     *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RoleId      *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SubmitInternalPurchaseOrdersRequestOrderList) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseOrdersRequestOrderList) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseOrdersRequestOrderList) SetBatchNumber(v string) *SubmitInternalPurchaseOrdersRequestOrderList {
	s.BatchNumber = &v
	return s
}

func (s *SubmitInternalPurchaseOrdersRequestOrderList) SetFinalPrice(v int64) *SubmitInternalPurchaseOrdersRequestOrderList {
	s.FinalPrice = &v
	return s
}

func (s *SubmitInternalPurchaseOrdersRequestOrderList) SetFinishTime(v int64) *SubmitInternalPurchaseOrdersRequestOrderList {
	s.FinishTime = &v
	return s
}

func (s *SubmitInternalPurchaseOrdersRequestOrderList) SetGameId(v string) *SubmitInternalPurchaseOrdersRequestOrderList {
	s.GameId = &v
	return s
}

func (s *SubmitInternalPurchaseOrdersRequestOrderList) SetOrderId(v string) *SubmitInternalPurchaseOrdersRequestOrderList {
	s.OrderId = &v
	return s
}

func (s *SubmitInternalPurchaseOrdersRequestOrderList) SetRoleId(v string) *SubmitInternalPurchaseOrdersRequestOrderList {
	s.RoleId = &v
	return s
}

func (s *SubmitInternalPurchaseOrdersRequestOrderList) SetUserId(v string) *SubmitInternalPurchaseOrdersRequestOrderList {
	s.UserId = &v
	return s
}

type SubmitInternalPurchaseOrdersResponseBody struct {
	Data      *SubmitInternalPurchaseOrdersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitInternalPurchaseOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseOrdersResponseBody) SetData(v *SubmitInternalPurchaseOrdersResponseBodyData) *SubmitInternalPurchaseOrdersResponseBody {
	s.Data = v
	return s
}

func (s *SubmitInternalPurchaseOrdersResponseBody) SetRequestId(v string) *SubmitInternalPurchaseOrdersResponseBody {
	s.RequestId = &v
	return s
}

type SubmitInternalPurchaseOrdersResponseBodyData struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status  *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitInternalPurchaseOrdersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseOrdersResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseOrdersResponseBodyData) SetMessage(v string) *SubmitInternalPurchaseOrdersResponseBodyData {
	s.Message = &v
	return s
}

func (s *SubmitInternalPurchaseOrdersResponseBodyData) SetStatus(v int32) *SubmitInternalPurchaseOrdersResponseBodyData {
	s.Status = &v
	return s
}

type SubmitInternalPurchaseOrdersResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitInternalPurchaseOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitInternalPurchaseOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseOrdersResponse) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseOrdersResponse) SetHeaders(v map[string]*string) *SubmitInternalPurchaseOrdersResponse {
	s.Headers = v
	return s
}

func (s *SubmitInternalPurchaseOrdersResponse) SetBody(v *SubmitInternalPurchaseOrdersResponseBody) *SubmitInternalPurchaseOrdersResponse {
	s.Body = v
	return s
}

type SubmitInternalPurchaseReadyFlagRequest struct {
	BatchInfoList   []*SubmitInternalPurchaseReadyFlagRequestBatchInfoList `json:"BatchInfoList,omitempty" xml:"BatchInfoList,omitempty" type:"Repeated"`
	ChargeDate      *string                                                `json:"ChargeDate,omitempty" xml:"ChargeDate,omitempty"`
	GameId          *string                                                `json:"GameId,omitempty" xml:"GameId,omitempty"`
	OrderTotalCount *int32                                                 `json:"OrderTotalCount,omitempty" xml:"OrderTotalCount,omitempty"`
	Status          *int32                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitInternalPurchaseReadyFlagRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseReadyFlagRequest) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseReadyFlagRequest) SetBatchInfoList(v []*SubmitInternalPurchaseReadyFlagRequestBatchInfoList) *SubmitInternalPurchaseReadyFlagRequest {
	s.BatchInfoList = v
	return s
}

func (s *SubmitInternalPurchaseReadyFlagRequest) SetChargeDate(v string) *SubmitInternalPurchaseReadyFlagRequest {
	s.ChargeDate = &v
	return s
}

func (s *SubmitInternalPurchaseReadyFlagRequest) SetGameId(v string) *SubmitInternalPurchaseReadyFlagRequest {
	s.GameId = &v
	return s
}

func (s *SubmitInternalPurchaseReadyFlagRequest) SetOrderTotalCount(v int32) *SubmitInternalPurchaseReadyFlagRequest {
	s.OrderTotalCount = &v
	return s
}

func (s *SubmitInternalPurchaseReadyFlagRequest) SetStatus(v int32) *SubmitInternalPurchaseReadyFlagRequest {
	s.Status = &v
	return s
}

type SubmitInternalPurchaseReadyFlagRequestBatchInfoList struct {
	BatchNumbers *string `json:"BatchNumbers,omitempty" xml:"BatchNumbers,omitempty"`
	BatchSize    *int32  `json:"BatchSize,omitempty" xml:"BatchSize,omitempty"`
}

func (s SubmitInternalPurchaseReadyFlagRequestBatchInfoList) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseReadyFlagRequestBatchInfoList) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseReadyFlagRequestBatchInfoList) SetBatchNumbers(v string) *SubmitInternalPurchaseReadyFlagRequestBatchInfoList {
	s.BatchNumbers = &v
	return s
}

func (s *SubmitInternalPurchaseReadyFlagRequestBatchInfoList) SetBatchSize(v int32) *SubmitInternalPurchaseReadyFlagRequestBatchInfoList {
	s.BatchSize = &v
	return s
}

type SubmitInternalPurchaseReadyFlagResponseBody struct {
	Data      *SubmitInternalPurchaseReadyFlagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitInternalPurchaseReadyFlagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseReadyFlagResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseReadyFlagResponseBody) SetData(v *SubmitInternalPurchaseReadyFlagResponseBodyData) *SubmitInternalPurchaseReadyFlagResponseBody {
	s.Data = v
	return s
}

func (s *SubmitInternalPurchaseReadyFlagResponseBody) SetRequestId(v string) *SubmitInternalPurchaseReadyFlagResponseBody {
	s.RequestId = &v
	return s
}

type SubmitInternalPurchaseReadyFlagResponseBodyData struct {
	Message             *string `json:"Message,omitempty" xml:"Message,omitempty"`
	MissingBatchNumbers *string `json:"MissingBatchNumbers,omitempty" xml:"MissingBatchNumbers,omitempty"`
	Status              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SubmitInternalPurchaseReadyFlagResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseReadyFlagResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseReadyFlagResponseBodyData) SetMessage(v string) *SubmitInternalPurchaseReadyFlagResponseBodyData {
	s.Message = &v
	return s
}

func (s *SubmitInternalPurchaseReadyFlagResponseBodyData) SetMissingBatchNumbers(v string) *SubmitInternalPurchaseReadyFlagResponseBodyData {
	s.MissingBatchNumbers = &v
	return s
}

func (s *SubmitInternalPurchaseReadyFlagResponseBodyData) SetStatus(v int32) *SubmitInternalPurchaseReadyFlagResponseBodyData {
	s.Status = &v
	return s
}

type SubmitInternalPurchaseReadyFlagResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitInternalPurchaseReadyFlagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitInternalPurchaseReadyFlagResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitInternalPurchaseReadyFlagResponse) GoString() string {
	return s.String()
}

func (s *SubmitInternalPurchaseReadyFlagResponse) SetHeaders(v map[string]*string) *SubmitInternalPurchaseReadyFlagResponse {
	s.Headers = v
	return s
}

func (s *SubmitInternalPurchaseReadyFlagResponse) SetBody(v *SubmitInternalPurchaseReadyFlagResponseBody) *SubmitInternalPurchaseReadyFlagResponse {
	s.Body = v
	return s
}

type UploadGameVersionByDownloadRequest struct {
	DownloadType *string `json:"DownloadType,omitempty" xml:"DownloadType,omitempty"`
	FileType     *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	GameId       *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	GameVersion  *string `json:"GameVersion,omitempty" xml:"GameVersion,omitempty"`
	Hash         *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	VersionName  *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UploadGameVersionByDownloadRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadGameVersionByDownloadRequest) GoString() string {
	return s.String()
}

func (s *UploadGameVersionByDownloadRequest) SetDownloadType(v string) *UploadGameVersionByDownloadRequest {
	s.DownloadType = &v
	return s
}

func (s *UploadGameVersionByDownloadRequest) SetFileType(v string) *UploadGameVersionByDownloadRequest {
	s.FileType = &v
	return s
}

func (s *UploadGameVersionByDownloadRequest) SetGameId(v string) *UploadGameVersionByDownloadRequest {
	s.GameId = &v
	return s
}

func (s *UploadGameVersionByDownloadRequest) SetGameVersion(v string) *UploadGameVersionByDownloadRequest {
	s.GameVersion = &v
	return s
}

func (s *UploadGameVersionByDownloadRequest) SetHash(v string) *UploadGameVersionByDownloadRequest {
	s.Hash = &v
	return s
}

func (s *UploadGameVersionByDownloadRequest) SetVersionName(v string) *UploadGameVersionByDownloadRequest {
	s.VersionName = &v
	return s
}

type UploadGameVersionByDownloadResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 任务id
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UploadGameVersionByDownloadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadGameVersionByDownloadResponseBody) GoString() string {
	return s.String()
}

func (s *UploadGameVersionByDownloadResponseBody) SetRequestId(v string) *UploadGameVersionByDownloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadGameVersionByDownloadResponseBody) SetTaskId(v string) *UploadGameVersionByDownloadResponseBody {
	s.TaskId = &v
	return s
}

type UploadGameVersionByDownloadResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadGameVersionByDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadGameVersionByDownloadResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadGameVersionByDownloadResponse) GoString() string {
	return s.String()
}

func (s *UploadGameVersionByDownloadResponse) SetHeaders(v map[string]*string) *UploadGameVersionByDownloadResponse {
	s.Headers = v
	return s
}

func (s *UploadGameVersionByDownloadResponse) SetBody(v *UploadGameVersionByDownloadResponseBody) *UploadGameVersionByDownloadResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudgameapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AdaptGameVersionWithOptions(request *AdaptGameVersionRequest, runtime *util.RuntimeOptions) (_result *AdaptGameVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["FrameRate"] = request.FrameRate
	query["Resolution"] = request.Resolution
	query["VersionId"] = request.VersionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AdaptGameVersion"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AdaptGameVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AdaptGameVersion(request *AdaptGameVersionRequest) (_result *AdaptGameVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AdaptGameVersionResponse{}
	_body, _err := client.AdaptGameVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGameToProjectWithOptions(request *AddGameToProjectRequest, runtime *util.RuntimeOptions) (_result *AddGameToProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GameId"] = request.GameId
	query["ProjectId"] = request.ProjectId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddGameToProject"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddGameToProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddGameToProject(request *AddGameToProjectRequest) (_result *AddGameToProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddGameToProjectResponse{}
	_body, _err := client.AddGameToProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDispatchGameSlotWithOptions(request *BatchDispatchGameSlotRequest, runtime *util.RuntimeOptions) (_result *BatchDispatchGameSlotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueueUserList)) {
		body["QueueUserList"] = request.QueueUserList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDispatchGameSlot"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDispatchGameSlotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDispatchGameSlot(request *BatchDispatchGameSlotRequest) (_result *BatchDispatchGameSlotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDispatchGameSlotResponse{}
	_body, _err := client.BatchDispatchGameSlotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStopGameSessionsWithOptions(request *BatchStopGameSessionsRequest, runtime *util.RuntimeOptions) (_result *BatchStopGameSessionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GameId"] = request.GameId
	query["ProjectId"] = request.ProjectId
	query["Reason"] = request.Reason
	query["Tags"] = request.Tags
	query["Token"] = request.Token
	query["TrackInfo"] = request.TrackInfo
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchStopGameSessions"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchStopGameSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStopGameSessions(request *BatchStopGameSessionsRequest) (_result *BatchStopGameSessionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStopGameSessionsResponse{}
	_body, _err := client.BatchStopGameSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseOrderWithOptions(request *CloseOrderRequest, runtime *util.RuntimeOptions) (_result *CloseOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccountDomain"] = request.AccountDomain
	query["BuyerAccountId"] = request.BuyerAccountId
	query["OrderId"] = request.OrderId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseOrder"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseOrder(request *CloseOrderRequest) (_result *CloseOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseOrderResponse{}
	_body, _err := client.CloseOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGameWithOptions(request *CreateGameRequest, runtime *util.RuntimeOptions) (_result *CreateGameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["GameName"] = request.GameName
	query["PlatformType"] = request.PlatformType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGame"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGame(request *CreateGameRequest) (_result *CreateGameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGameResponse{}
	_body, _err := client.CreateGameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGameDeployWorkflowWithOptions(request *CreateGameDeployWorkflowRequest, runtime *util.RuntimeOptions) (_result *CreateGameDeployWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DownloadType"] = request.DownloadType
	query["FileType"] = request.FileType
	query["FrameRate"] = request.FrameRate
	query["GameId"] = request.GameId
	query["GameVersion"] = request.GameVersion
	query["Hash"] = request.Hash
	query["Instance"] = request.Instance
	query["ProjectId"] = request.ProjectId
	query["Resolution"] = request.Resolution
	query["VersionName"] = request.VersionName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGameDeployWorkflow"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGameDeployWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGameDeployWorkflow(request *CreateGameDeployWorkflowRequest) (_result *CreateGameDeployWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGameDeployWorkflowResponse{}
	_body, _err := client.CreateGameDeployWorkflowWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["AccountDomain"] = request.AccountDomain
	query["Amount"] = request.Amount
	query["BuyerAccountId"] = request.BuyerAccountId
	query["IdempotentCode"] = request.IdempotentCode
	query["ItemId"] = request.ItemId
	query["OriginPrice"] = request.OriginPrice
	query["SettlementPrice"] = request.SettlementPrice
	query["SkuId"] = request.SkuId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrder"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["ProjectName"] = request.ProjectName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTokenWithOptions(request *CreateTokenRequest, runtime *util.RuntimeOptions) (_result *CreateTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["CurrentToken"] = request.CurrentToken
	query["Session"] = request.Session
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateToken"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateToken(request *CreateTokenRequest) (_result *CreateTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTokenResponse{}
	_body, _err := client.CreateTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGameWithOptions(request *DeleteGameRequest, runtime *util.RuntimeOptions) (_result *DeleteGameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GameId"] = request.GameId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGame"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGame(request *DeleteGameRequest) (_result *DeleteGameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGameResponse{}
	_body, _err := client.DeleteGameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGameVersionWithOptions(request *DeleteGameVersionRequest, runtime *util.RuntimeOptions) (_result *DeleteGameVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["VersionId"] = request.VersionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGameVersion"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGameVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGameVersion(request *DeleteGameVersionRequest) (_result *DeleteGameVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGameVersionResponse{}
	_body, _err := client.DeleteGameVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ProjectId"] = request.ProjectId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeliveryOrderWithOptions(request *DeliveryOrderRequest, runtime *util.RuntimeOptions) (_result *DeliveryOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccountDomain"] = request.AccountDomain
	query["BuyerAccountId"] = request.BuyerAccountId
	query["OrderId"] = request.OrderId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeliveryOrder"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeliveryOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeliveryOrder(request *DeliveryOrderRequest) (_result *DeliveryOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeliveryOrderResponse{}
	_body, _err := client.DeliveryOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DispatchGameSlotWithOptions(request *DispatchGameSlotRequest, runtime *util.RuntimeOptions) (_result *DispatchGameSlotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.BizParam)) {
		body["BizParam"] = request.BizParam
	}

	if !tea.BoolValue(util.IsUnset(request.Cancel)) {
		body["Cancel"] = request.Cancel
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		body["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.GameCommand)) {
		body["GameCommand"] = request.GameCommand
	}

	if !tea.BoolValue(util.IsUnset(request.GameId)) {
		body["GameId"] = request.GameId
	}

	if !tea.BoolValue(util.IsUnset(request.GameSession)) {
		body["GameSession"] = request.GameSession
	}

	if !tea.BoolValue(util.IsUnset(request.GameStartParam)) {
		body["GameStartParam"] = request.GameStartParam
	}

	if !tea.BoolValue(util.IsUnset(request.Reconnect)) {
		body["Reconnect"] = request.Reconnect
	}

	if !tea.BoolValue(util.IsUnset(request.RegionName)) {
		body["RegionName"] = request.RegionName
	}

	if !tea.BoolValue(util.IsUnset(request.SystemInfo)) {
		body["SystemInfo"] = request.SystemInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserLevel)) {
		body["UserLevel"] = request.UserLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DispatchGameSlot"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DispatchGameSlotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DispatchGameSlot(request *DispatchGameSlotRequest) (_result *DispatchGameSlotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DispatchGameSlotResponse{}
	_body, _err := client.DispatchGameSlotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGameCcuWithOptions(request *GetGameCcuRequest, runtime *util.RuntimeOptions) (_result *GetGameCcuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessKey"] = request.AccessKey
	query["GameId"] = request.GameId
	query["RegionName"] = request.RegionName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGameCcu"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGameCcuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGameCcu(request *GetGameCcuRequest) (_result *GetGameCcuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGameCcuResponse{}
	_body, _err := client.GetGameCcuWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGameStockWithOptions(request *GetGameStockRequest, runtime *util.RuntimeOptions) (_result *GetGameStockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessKey"] = request.AccessKey
	query["GameId"] = request.GameId
	query["UserLevel"] = request.UserLevel
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGameStock"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGameStockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGameStock(request *GetGameStockRequest) (_result *GetGameStockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGameStockResponse{}
	_body, _err := client.GetGameStockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGameTrialSurplusDurationWithOptions(request *GetGameTrialSurplusDurationRequest, runtime *util.RuntimeOptions) (_result *GetGameTrialSurplusDurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccountId"] = request.AccountId
	query["GameId"] = request.GameId
	query["ProjectId"] = request.ProjectId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGameTrialSurplusDuration"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGameTrialSurplusDurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGameTrialSurplusDuration(request *GetGameTrialSurplusDurationRequest) (_result *GetGameTrialSurplusDurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGameTrialSurplusDurationResponse{}
	_body, _err := client.GetGameTrialSurplusDurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGameVersionWithOptions(request *GetGameVersionRequest, runtime *util.RuntimeOptions) (_result *GetGameVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["VersionId"] = request.VersionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGameVersion"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGameVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGameVersion(request *GetGameVersionRequest) (_result *GetGameVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGameVersionResponse{}
	_body, _err := client.GetGameVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGameVersionProgressWithOptions(request *GetGameVersionProgressRequest, runtime *util.RuntimeOptions) (_result *GetGameVersionProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGameVersionProgress"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGameVersionProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGameVersionProgress(request *GetGameVersionProgressRequest) (_result *GetGameVersionProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGameVersionProgressResponse{}
	_body, _err := client.GetGameVersionProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetItemWithOptions(request *GetItemRequest, runtime *util.RuntimeOptions) (_result *GetItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ItemId"] = request.ItemId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetItem"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetItem(request *GetItemRequest) (_result *GetItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetItemResponse{}
	_body, _err := client.GetItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOutAccountBindDetailWithOptions(request *GetOutAccountBindDetailRequest, runtime *util.RuntimeOptions) (_result *GetOutAccountBindDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccountDomain"] = request.AccountDomain
	query["AccountId"] = request.AccountId
	query["OutAccountType"] = request.OutAccountType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOutAccountBindDetail"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOutAccountBindDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOutAccountBindDetail(request *GetOutAccountBindDetailRequest) (_result *GetOutAccountBindDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOutAccountBindDetailResponse{}
	_body, _err := client.GetOutAccountBindDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSessionWithOptions(request *GetSessionRequest, runtime *util.RuntimeOptions) (_result *GetSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Token"] = request.Token
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSession"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSession(request *GetSessionRequest) (_result *GetSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSessionResponse{}
	_body, _err := client.GetSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStopGameTokenWithOptions(request *GetStopGameTokenRequest, runtime *util.RuntimeOptions) (_result *GetStopGameTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccessKey"] = request.AccessKey
	query["GameId"] = request.GameId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStopGameToken"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStopGameTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStopGameToken(request *GetStopGameTokenRequest) (_result *GetStopGameTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStopGameTokenResponse{}
	_body, _err := client.GetStopGameTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBoughtGamesWithOptions(request *ListBoughtGamesRequest, runtime *util.RuntimeOptions) (_result *ListBoughtGamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccountDomain"] = request.AccountDomain
	query["AccountId"] = request.AccountId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBoughtGames"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBoughtGamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBoughtGames(request *ListBoughtGamesRequest) (_result *ListBoughtGamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBoughtGamesResponse{}
	_body, _err := client.ListBoughtGamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContainerStatusWithOptions(request *ListContainerStatusRequest, runtime *util.RuntimeOptions) (_result *ListContainerStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GameSessionIdList"] = request.GameSessionIdList
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListContainerStatus"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListContainerStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContainerStatus(request *ListContainerStatusRequest) (_result *ListContainerStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContainerStatusResponse{}
	_body, _err := client.ListContainerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDeployableInstancesWithOptions(request *ListDeployableInstancesRequest, runtime *util.RuntimeOptions) (_result *ListDeployableInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ProjectId"] = request.ProjectId
	query["VersionId"] = request.VersionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeployableInstances"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeployableInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDeployableInstances(request *ListDeployableInstancesRequest) (_result *ListDeployableInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeployableInstancesResponse{}
	_body, _err := client.ListDeployableInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGameVersionsWithOptions(request *ListGameVersionsRequest, runtime *util.RuntimeOptions) (_result *ListGameVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GameId"] = request.GameId
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGameVersions"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGameVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGameVersions(request *ListGameVersionsRequest) (_result *ListGameVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGameVersionsResponse{}
	_body, _err := client.ListGameVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGamesWithOptions(request *ListGamesRequest, runtime *util.RuntimeOptions) (_result *ListGamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGames"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGames(request *ListGamesRequest) (_result *ListGamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGamesResponse{}
	_body, _err := client.ListGamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHistoryContainerStatusWithOptions(request *ListHistoryContainerStatusRequest, runtime *util.RuntimeOptions) (_result *ListHistoryContainerStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["EndTime"] = request.EndTime
	query["LastGameSessionId"] = request.LastGameSessionId
	query["PageSize"] = request.PageSize
	query["ProjectId"] = request.ProjectId
	query["StartTime"] = request.StartTime
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHistoryContainerStatus"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHistoryContainerStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHistoryContainerStatus(request *ListHistoryContainerStatusRequest) (_result *ListHistoryContainerStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHistoryContainerStatusResponse{}
	_body, _err := client.ListHistoryContainerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryGameWithOptions(request *QueryGameRequest, runtime *util.RuntimeOptions) (_result *QueryGameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PageNo"] = request.PageNo
	query["PageSize"] = request.PageSize
	query["ProjectId"] = request.ProjectId
	query["TenantId"] = request.TenantId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryGame"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryGame(request *QueryGameRequest) (_result *QueryGameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGameResponse{}
	_body, _err := client.QueryGameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryItemsWithOptions(request *QueryItemsRequest, runtime *util.RuntimeOptions) (_result *QueryItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryItems"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryItems(request *QueryItemsRequest) (_result *QueryItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryItemsResponse{}
	_body, _err := client.QueryItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOrderWithOptions(request *QueryOrderRequest, runtime *util.RuntimeOptions) (_result *QueryOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccountDomain"] = request.AccountDomain
	query["BuyerAccountId"] = request.BuyerAccountId
	query["OrderId"] = request.OrderId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOrder"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOrder(request *QueryOrderRequest) (_result *QueryOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOrderResponse{}
	_body, _err := client.QueryOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOutAccountBindStatusWithOptions(request *QueryOutAccountBindStatusRequest, runtime *util.RuntimeOptions) (_result *QueryOutAccountBindStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AccountDomain"] = request.AccountDomain
	query["AccountId"] = request.AccountId
	query["GameId"] = request.GameId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOutAccountBindStatus"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOutAccountBindStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOutAccountBindStatus(request *QueryOutAccountBindStatusRequest) (_result *QueryOutAccountBindStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOutAccountBindStatusResponse{}
	_body, _err := client.QueryOutAccountBindStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryProjectWithOptions(request *QueryProjectRequest, runtime *util.RuntimeOptions) (_result *QueryProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PageNo"] = request.PageNo
	query["PageSize"] = request.PageSize
	query["ProjectId"] = request.ProjectId
	query["TenantId"] = request.TenantId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryProject"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryProject(request *QueryProjectRequest) (_result *QueryProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryProjectResponse{}
	_body, _err := client.QueryProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTenantWithOptions(request *QueryTenantRequest, runtime *util.RuntimeOptions) (_result *QueryTenantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PageNo"] = request.PageNo
	query["PageSize"] = request.PageSize
	query["Param"] = request.Param
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTenant"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTenantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTenant(request *QueryTenantRequest) (_result *QueryTenantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTenantResponse{}
	_body, _err := client.QueryTenantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveGameFromProjectWithOptions(request *RemoveGameFromProjectRequest, runtime *util.RuntimeOptions) (_result *RemoveGameFromProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GameId"] = request.GameId
	query["ProjectId"] = request.ProjectId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveGameFromProject"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveGameFromProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveGameFromProject(request *RemoveGameFromProjectRequest) (_result *RemoveGameFromProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveGameFromProjectResponse{}
	_body, _err := client.RemoveGameFromProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SkipTrialPolicyWithOptions(request *SkipTrialPolicyRequest, runtime *util.RuntimeOptions) (_result *SkipTrialPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GameSessionId"] = request.GameSessionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SkipTrialPolicy"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SkipTrialPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SkipTrialPolicy(request *SkipTrialPolicyRequest) (_result *SkipTrialPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SkipTrialPolicyResponse{}
	_body, _err := client.SkipTrialPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopGameSessionWithOptions(request *StopGameSessionRequest, runtime *util.RuntimeOptions) (_result *StopGameSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.BizParam)) {
		body["BizParam"] = request.BizParam
	}

	if !tea.BoolValue(util.IsUnset(request.GameId)) {
		body["GameId"] = request.GameId
	}

	if !tea.BoolValue(util.IsUnset(request.GameSession)) {
		body["GameSession"] = request.GameSession
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		body["Reason"] = request.Reason
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopGameSession"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopGameSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopGameSession(request *StopGameSessionRequest) (_result *StopGameSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopGameSessionResponse{}
	_body, _err := client.StopGameSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDeploymentWithOptions(request *SubmitDeploymentRequest, runtime *util.RuntimeOptions) (_result *SubmitDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CloudGameInstanceIds"] = request.CloudGameInstanceIds
	query["GameId"] = request.GameId
	query["OperationType"] = request.OperationType
	query["ProjectId"] = request.ProjectId
	query["VersionId"] = request.VersionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitDeployment"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitDeployment(request *SubmitDeploymentRequest) (_result *SubmitDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitDeploymentResponse{}
	_body, _err := client.SubmitDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitInternalPurchaseChargeDataWithOptions(request *SubmitInternalPurchaseChargeDataRequest, runtime *util.RuntimeOptions) (_result *SubmitInternalPurchaseChargeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ActiveUserRetentionRateOneDay"] = request.ActiveUserRetentionRateOneDay
	query["ActiveUserRetentionRateSevenDay"] = request.ActiveUserRetentionRateSevenDay
	query["ActiveUserRetentionRateThirtyDay"] = request.ActiveUserRetentionRateThirtyDay
	query["Arpu"] = request.Arpu
	query["ChargeDate"] = request.ChargeDate
	query["Dau"] = request.Dau
	query["GameId"] = request.GameId
	query["Mau"] = request.Mau
	query["NewUserRetentionRateOneDay"] = request.NewUserRetentionRateOneDay
	query["NewUserRetentionRateSevenDay"] = request.NewUserRetentionRateSevenDay
	query["NewUserRetentionRateThirtyDay"] = request.NewUserRetentionRateThirtyDay
	query["PaymentConversionRate"] = request.PaymentConversionRate
	query["PlayTimeAverageOneDay"] = request.PlayTimeAverageOneDay
	query["PlayTimeAverageThirtyDay"] = request.PlayTimeAverageThirtyDay
	query["PlayTimeNinetyPointsOneDay"] = request.PlayTimeNinetyPointsOneDay
	query["PlayTimeNinetyPointsThirtyDay"] = request.PlayTimeNinetyPointsThirtyDay
	query["PlayTimeRangeOneDay"] = request.PlayTimeRangeOneDay
	query["PlayTimeRangeThirtyDay"] = request.PlayTimeRangeThirtyDay
	query["UserActivationRate"] = request.UserActivationRate
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitInternalPurchaseChargeData"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitInternalPurchaseChargeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitInternalPurchaseChargeData(request *SubmitInternalPurchaseChargeDataRequest) (_result *SubmitInternalPurchaseChargeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitInternalPurchaseChargeDataResponse{}
	_body, _err := client.SubmitInternalPurchaseChargeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitInternalPurchaseOrdersWithOptions(request *SubmitInternalPurchaseOrdersRequest, runtime *util.RuntimeOptions) (_result *SubmitInternalPurchaseOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OrderList"] = request.OrderList
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitInternalPurchaseOrders"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitInternalPurchaseOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitInternalPurchaseOrders(request *SubmitInternalPurchaseOrdersRequest) (_result *SubmitInternalPurchaseOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitInternalPurchaseOrdersResponse{}
	_body, _err := client.SubmitInternalPurchaseOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitInternalPurchaseReadyFlagWithOptions(request *SubmitInternalPurchaseReadyFlagRequest, runtime *util.RuntimeOptions) (_result *SubmitInternalPurchaseReadyFlagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["BatchInfoList"] = request.BatchInfoList
	query["ChargeDate"] = request.ChargeDate
	query["GameId"] = request.GameId
	query["OrderTotalCount"] = request.OrderTotalCount
	query["Status"] = request.Status
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitInternalPurchaseReadyFlag"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitInternalPurchaseReadyFlagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitInternalPurchaseReadyFlag(request *SubmitInternalPurchaseReadyFlagRequest) (_result *SubmitInternalPurchaseReadyFlagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitInternalPurchaseReadyFlagResponse{}
	_body, _err := client.SubmitInternalPurchaseReadyFlagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadGameVersionByDownloadWithOptions(request *UploadGameVersionByDownloadRequest, runtime *util.RuntimeOptions) (_result *UploadGameVersionByDownloadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DownloadType"] = request.DownloadType
	query["FileType"] = request.FileType
	query["GameId"] = request.GameId
	query["GameVersion"] = request.GameVersion
	query["Hash"] = request.Hash
	query["VersionName"] = request.VersionName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadGameVersionByDownload"),
		Version:     tea.String("2020-07-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadGameVersionByDownloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadGameVersionByDownload(request *UploadGameVersionByDownloadRequest) (_result *UploadGameVersionByDownloadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadGameVersionByDownloadResponse{}
	_body, _err := client.UploadGameVersionByDownloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
