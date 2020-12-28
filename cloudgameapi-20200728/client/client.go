// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type QueryOutAccountBindStatusRequest struct {
	AccountId     *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	GameId        *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	AccountDomain *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
}

func (s QueryOutAccountBindStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOutAccountBindStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryOutAccountBindStatusRequest) SetAccountId(v string) *QueryOutAccountBindStatusRequest {
	s.AccountId = &v
	return s
}

func (s *QueryOutAccountBindStatusRequest) SetGameId(v string) *QueryOutAccountBindStatusRequest {
	s.GameId = &v
	return s
}

func (s *QueryOutAccountBindStatusRequest) SetAccountDomain(v string) *QueryOutAccountBindStatusRequest {
	s.AccountDomain = &v
	return s
}

type QueryOutAccountBindStatusResponse struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *QueryOutAccountBindStatusResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s QueryOutAccountBindStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOutAccountBindStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryOutAccountBindStatusResponse) SetRequestId(v string) *QueryOutAccountBindStatusResponse {
	s.RequestId = &v
	return s
}

func (s *QueryOutAccountBindStatusResponse) SetData(v *QueryOutAccountBindStatusResponseData) *QueryOutAccountBindStatusResponse {
	s.Data = v
	return s
}

type QueryOutAccountBindStatusResponseData struct {
	BindStatus *int `json:"BindStatus,omitempty" xml:"BindStatus,omitempty" require:"true"`
}

func (s QueryOutAccountBindStatusResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryOutAccountBindStatusResponseData) GoString() string {
	return s.String()
}

func (s *QueryOutAccountBindStatusResponseData) SetBindStatus(v int) *QueryOutAccountBindStatusResponseData {
	s.BindStatus = &v
	return s
}

type CreateTokenRequest struct {
	Session      *string `json:"Session,omitempty" xml:"Session,omitempty" require:"true"`
	CurrentToken *string `json:"CurrentToken,omitempty" xml:"CurrentToken,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateTokenRequest) SetSession(v string) *CreateTokenRequest {
	s.Session = &v
	return s
}

func (s *CreateTokenRequest) SetCurrentToken(v string) *CreateTokenRequest {
	s.CurrentToken = &v
	return s
}

func (s *CreateTokenRequest) SetClientToken(v string) *CreateTokenRequest {
	s.ClientToken = &v
	return s
}

type CreateTokenResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *CreateTokenResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s CreateTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateTokenResponse) SetRequestId(v string) *CreateTokenResponse {
	s.RequestId = &v
	return s
}

func (s *CreateTokenResponse) SetData(v *CreateTokenResponseData) *CreateTokenResponse {
	s.Data = v
	return s
}

type CreateTokenResponseData struct {
	Token *string `json:"Token,omitempty" xml:"Token,omitempty" require:"true"`
}

func (s CreateTokenResponseData) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponseData) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseData) SetToken(v string) *CreateTokenResponseData {
	s.Token = &v
	return s
}

type GetSessionRequest struct {
	Token *string `json:"Token,omitempty" xml:"Token,omitempty" require:"true"`
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

type GetSessionResponse struct {
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetSessionResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSessionResponse) GoString() string {
	return s.String()
}

func (s *GetSessionResponse) SetRequestId(v string) *GetSessionResponse {
	s.RequestId = &v
	return s
}

func (s *GetSessionResponse) SetData(v *GetSessionResponseData) *GetSessionResponse {
	s.Data = v
	return s
}

type GetSessionResponseData struct {
	Session *string `json:"Session,omitempty" xml:"Session,omitempty" require:"true"`
}

func (s GetSessionResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetSessionResponseData) GoString() string {
	return s.String()
}

func (s *GetSessionResponseData) SetSession(v string) *GetSessionResponseData {
	s.Session = &v
	return s
}

type SkipTrialPolicyRequest struct {
	GameSessionId *string `json:"GameSessionId,omitempty" xml:"GameSessionId,omitempty" require:"true"`
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

type SkipTrialPolicyResponse struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *SkipTrialPolicyResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s SkipTrialPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SkipTrialPolicyResponse) GoString() string {
	return s.String()
}

func (s *SkipTrialPolicyResponse) SetRequestId(v string) *SkipTrialPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *SkipTrialPolicyResponse) SetData(v *SkipTrialPolicyResponseData) *SkipTrialPolicyResponse {
	s.Data = v
	return s
}

type SkipTrialPolicyResponseData struct {
	SkipResult *int `json:"SkipResult,omitempty" xml:"SkipResult,omitempty" require:"true"`
}

func (s SkipTrialPolicyResponseData) String() string {
	return tea.Prettify(s)
}

func (s SkipTrialPolicyResponseData) GoString() string {
	return s.String()
}

func (s *SkipTrialPolicyResponseData) SetSkipResult(v int) *SkipTrialPolicyResponseData {
	s.SkipResult = &v
	return s
}

type QueryTenantRequest struct {
	Param    *string `json:"Param,omitempty" xml:"Param,omitempty"`
	PageNo   *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s QueryTenantRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTenantRequest) GoString() string {
	return s.String()
}

func (s *QueryTenantRequest) SetParam(v string) *QueryTenantRequest {
	s.Param = &v
	return s
}

func (s *QueryTenantRequest) SetPageNo(v int) *QueryTenantRequest {
	s.PageNo = &v
	return s
}

func (s *QueryTenantRequest) SetPageSize(v int) *QueryTenantRequest {
	s.PageSize = &v
	return s
}

type QueryTenantResponse struct {
	RequestId  *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                       `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Data       []*QueryTenantResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryTenantResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTenantResponse) GoString() string {
	return s.String()
}

func (s *QueryTenantResponse) SetRequestId(v string) *QueryTenantResponse {
	s.RequestId = &v
	return s
}

func (s *QueryTenantResponse) SetPageNumber(v int) *QueryTenantResponse {
	s.PageNumber = &v
	return s
}

func (s *QueryTenantResponse) SetPageSize(v int) *QueryTenantResponse {
	s.PageSize = &v
	return s
}

func (s *QueryTenantResponse) SetTotalCount(v int) *QueryTenantResponse {
	s.TotalCount = &v
	return s
}

func (s *QueryTenantResponse) SetData(v []*QueryTenantResponseData) *QueryTenantResponse {
	s.Data = v
	return s
}

type QueryTenantResponseData struct {
	TenantId  *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty" require:"true"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
}

func (s QueryTenantResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryTenantResponseData) GoString() string {
	return s.String()
}

func (s *QueryTenantResponseData) SetTenantId(v int64) *QueryTenantResponseData {
	s.TenantId = &v
	return s
}

func (s *QueryTenantResponseData) SetName(v string) *QueryTenantResponseData {
	s.Name = &v
	return s
}

func (s *QueryTenantResponseData) SetGmtCreate(v string) *QueryTenantResponseData {
	s.GmtCreate = &v
	return s
}

type QueryProjectRequest struct {
	PageNo    *int   `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize  *int   `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TenantId  *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectRequest) GoString() string {
	return s.String()
}

func (s *QueryProjectRequest) SetPageNo(v int) *QueryProjectRequest {
	s.PageNo = &v
	return s
}

func (s *QueryProjectRequest) SetPageSize(v int) *QueryProjectRequest {
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

type QueryProjectResponse struct {
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                        `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Data       []*QueryProjectResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponse) GoString() string {
	return s.String()
}

func (s *QueryProjectResponse) SetRequestId(v string) *QueryProjectResponse {
	s.RequestId = &v
	return s
}

func (s *QueryProjectResponse) SetPageNumber(v int) *QueryProjectResponse {
	s.PageNumber = &v
	return s
}

func (s *QueryProjectResponse) SetPageSize(v int) *QueryProjectResponse {
	s.PageSize = &v
	return s
}

func (s *QueryProjectResponse) SetTotalCount(v int) *QueryProjectResponse {
	s.TotalCount = &v
	return s
}

func (s *QueryProjectResponse) SetData(v []*QueryProjectResponseData) *QueryProjectResponse {
	s.Data = v
	return s
}

type QueryProjectResponseData struct {
	TenantId *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty" require:"true"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty" require:"true"`
}

func (s QueryProjectResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryProjectResponseData) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseData) SetTenantId(v int64) *QueryProjectResponseData {
	s.TenantId = &v
	return s
}

func (s *QueryProjectResponseData) SetName(v string) *QueryProjectResponseData {
	s.Name = &v
	return s
}

func (s *QueryProjectResponseData) SetId(v int64) *QueryProjectResponseData {
	s.Id = &v
	return s
}

type QueryGameRequest struct {
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageNo    *int   `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize  *int   `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TenantId  *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryGameRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGameRequest) GoString() string {
	return s.String()
}

func (s *QueryGameRequest) SetProjectId(v int64) *QueryGameRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryGameRequest) SetPageNo(v int) *QueryGameRequest {
	s.PageNo = &v
	return s
}

func (s *QueryGameRequest) SetPageSize(v int) *QueryGameRequest {
	s.PageSize = &v
	return s
}

func (s *QueryGameRequest) SetTenantId(v int64) *QueryGameRequest {
	s.TenantId = &v
	return s
}

type QueryGameResponse struct {
	RequestId  *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                     `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Data       []*QueryGameResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Repeated"`
}

func (s QueryGameResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGameResponse) GoString() string {
	return s.String()
}

func (s *QueryGameResponse) SetRequestId(v string) *QueryGameResponse {
	s.RequestId = &v
	return s
}

func (s *QueryGameResponse) SetPageNumber(v int) *QueryGameResponse {
	s.PageNumber = &v
	return s
}

func (s *QueryGameResponse) SetPageSize(v int) *QueryGameResponse {
	s.PageSize = &v
	return s
}

func (s *QueryGameResponse) SetTotalCount(v int) *QueryGameResponse {
	s.TotalCount = &v
	return s
}

func (s *QueryGameResponse) SetData(v []*QueryGameResponseData) *QueryGameResponse {
	s.Data = v
	return s
}

type QueryGameResponseData struct {
	TenantId  *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty" require:"true"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
	ProjectId *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty" require:"true"`
	GameId    *int64  `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	Version   *string `json:"Version,omitempty" xml:"Version,omitempty" require:"true"`
}

func (s QueryGameResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryGameResponseData) GoString() string {
	return s.String()
}

func (s *QueryGameResponseData) SetTenantId(v int64) *QueryGameResponseData {
	s.TenantId = &v
	return s
}

func (s *QueryGameResponseData) SetName(v string) *QueryGameResponseData {
	s.Name = &v
	return s
}

func (s *QueryGameResponseData) SetGmtCreate(v string) *QueryGameResponseData {
	s.GmtCreate = &v
	return s
}

func (s *QueryGameResponseData) SetProjectId(v int64) *QueryGameResponseData {
	s.ProjectId = &v
	return s
}

func (s *QueryGameResponseData) SetGameId(v int64) *QueryGameResponseData {
	s.GameId = &v
	return s
}

func (s *QueryGameResponseData) SetVersion(v string) *QueryGameResponseData {
	s.Version = &v
	return s
}

type BatchStopGameSessionsRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty" require:"true"`
	GameId    *string `json:"GameId,omitempty" xml:"GameId,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Reason    *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	TrackInfo *string `json:"TrackInfo,omitempty" xml:"TrackInfo,omitempty"`
}

func (s BatchStopGameSessionsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStopGameSessionsRequest) GoString() string {
	return s.String()
}

func (s *BatchStopGameSessionsRequest) SetProjectId(v string) *BatchStopGameSessionsRequest {
	s.ProjectId = &v
	return s
}

func (s *BatchStopGameSessionsRequest) SetGameId(v string) *BatchStopGameSessionsRequest {
	s.GameId = &v
	return s
}

func (s *BatchStopGameSessionsRequest) SetToken(v string) *BatchStopGameSessionsRequest {
	s.Token = &v
	return s
}

func (s *BatchStopGameSessionsRequest) SetReason(v string) *BatchStopGameSessionsRequest {
	s.Reason = &v
	return s
}

func (s *BatchStopGameSessionsRequest) SetTrackInfo(v string) *BatchStopGameSessionsRequest {
	s.TrackInfo = &v
	return s
}

type BatchStopGameSessionsResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	QueueState *int    `json:"QueueState,omitempty" xml:"QueueState,omitempty" require:"true"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	GameId     *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty" require:"true"`
	TrackInfo  *string `json:"TrackInfo,omitempty" xml:"TrackInfo,omitempty" require:"true"`
}

func (s BatchStopGameSessionsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStopGameSessionsResponse) GoString() string {
	return s.String()
}

func (s *BatchStopGameSessionsResponse) SetRequestId(v string) *BatchStopGameSessionsResponse {
	s.RequestId = &v
	return s
}

func (s *BatchStopGameSessionsResponse) SetSuccess(v bool) *BatchStopGameSessionsResponse {
	s.Success = &v
	return s
}

func (s *BatchStopGameSessionsResponse) SetQueueState(v int) *BatchStopGameSessionsResponse {
	s.QueueState = &v
	return s
}

func (s *BatchStopGameSessionsResponse) SetMessage(v string) *BatchStopGameSessionsResponse {
	s.Message = &v
	return s
}

func (s *BatchStopGameSessionsResponse) SetGameId(v string) *BatchStopGameSessionsResponse {
	s.GameId = &v
	return s
}

func (s *BatchStopGameSessionsResponse) SetProjectId(v string) *BatchStopGameSessionsResponse {
	s.ProjectId = &v
	return s
}

func (s *BatchStopGameSessionsResponse) SetTrackInfo(v string) *BatchStopGameSessionsResponse {
	s.TrackInfo = &v
	return s
}

type GetStopGameTokenRequest struct {
	GameId    *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" require:"true"`
}

func (s GetStopGameTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStopGameTokenRequest) GoString() string {
	return s.String()
}

func (s *GetStopGameTokenRequest) SetGameId(v string) *GetStopGameTokenRequest {
	s.GameId = &v
	return s
}

func (s *GetStopGameTokenRequest) SetAccessKey(v string) *GetStopGameTokenRequest {
	s.AccessKey = &v
	return s
}

type GetStopGameTokenResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty" require:"true"`
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
}

func (s GetStopGameTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStopGameTokenResponse) GoString() string {
	return s.String()
}

func (s *GetStopGameTokenResponse) SetRequestId(v string) *GetStopGameTokenResponse {
	s.RequestId = &v
	return s
}

func (s *GetStopGameTokenResponse) SetToken(v string) *GetStopGameTokenResponse {
	s.Token = &v
	return s
}

func (s *GetStopGameTokenResponse) SetExpireTime(v int64) *GetStopGameTokenResponse {
	s.ExpireTime = &v
	return s
}

type QueryItemsRequest struct {
	PageNumber *int `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s QueryItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsRequest) GoString() string {
	return s.String()
}

func (s *QueryItemsRequest) SetPageNumber(v int) *QueryItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryItemsRequest) SetPageSize(v int) *QueryItemsRequest {
	s.PageSize = &v
	return s
}

type QueryItemsResponse struct {
	Success        *bool                   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	RequestId      *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	HttpStatusCode *int64                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty" require:"true"`
	Data           *QueryItemsResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s QueryItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponse) GoString() string {
	return s.String()
}

func (s *QueryItemsResponse) SetSuccess(v bool) *QueryItemsResponse {
	s.Success = &v
	return s
}

func (s *QueryItemsResponse) SetRequestId(v string) *QueryItemsResponse {
	s.RequestId = &v
	return s
}

func (s *QueryItemsResponse) SetHttpStatusCode(v int64) *QueryItemsResponse {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryItemsResponse) SetData(v *QueryItemsResponseData) *QueryItemsResponse {
	s.Data = v
	return s
}

type QueryItemsResponseData struct {
	PageNumber *int                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                           `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int64                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Items      []*QueryItemsResponseDataItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Repeated"`
}

func (s QueryItemsResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponseData) GoString() string {
	return s.String()
}

func (s *QueryItemsResponseData) SetPageNumber(v int) *QueryItemsResponseData {
	s.PageNumber = &v
	return s
}

func (s *QueryItemsResponseData) SetPageSize(v int) *QueryItemsResponseData {
	s.PageSize = &v
	return s
}

func (s *QueryItemsResponseData) SetTotalCount(v int64) *QueryItemsResponseData {
	s.TotalCount = &v
	return s
}

func (s *QueryItemsResponseData) SetItems(v []*QueryItemsResponseDataItems) *QueryItemsResponseData {
	s.Items = v
	return s
}

type QueryItemsResponseDataItems struct {
	ModifyTime  *int64                              `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	SalePrice   *int64                              `json:"SalePrice,omitempty" xml:"SalePrice,omitempty" require:"true"`
	Description *string                             `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	OriginPrice *int64                              `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty" require:"true"`
	CreateTime  *int64                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	Title       *string                             `json:"Title,omitempty" xml:"Title,omitempty" require:"true"`
	ItemId      *string                             `json:"ItemId,omitempty" xml:"ItemId,omitempty" require:"true"`
	SellerId    *string                             `json:"SellerId,omitempty" xml:"SellerId,omitempty" require:"true"`
	Supplier    *string                             `json:"Supplier,omitempty" xml:"Supplier,omitempty" require:"true"`
	CategoryId  *int64                              `json:"CategoryId,omitempty" xml:"CategoryId,omitempty" require:"true"`
	Status      *int                                `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Skus        []*QueryItemsResponseDataItemsSkus  `json:"Skus,omitempty" xml:"Skus,omitempty" require:"true" type:"Repeated"`
	Games       []*QueryItemsResponseDataItemsGames `json:"Games,omitempty" xml:"Games,omitempty" require:"true" type:"Repeated"`
}

func (s QueryItemsResponseDataItems) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponseDataItems) GoString() string {
	return s.String()
}

func (s *QueryItemsResponseDataItems) SetModifyTime(v int64) *QueryItemsResponseDataItems {
	s.ModifyTime = &v
	return s
}

func (s *QueryItemsResponseDataItems) SetSalePrice(v int64) *QueryItemsResponseDataItems {
	s.SalePrice = &v
	return s
}

func (s *QueryItemsResponseDataItems) SetDescription(v string) *QueryItemsResponseDataItems {
	s.Description = &v
	return s
}

func (s *QueryItemsResponseDataItems) SetOriginPrice(v int64) *QueryItemsResponseDataItems {
	s.OriginPrice = &v
	return s
}

func (s *QueryItemsResponseDataItems) SetCreateTime(v int64) *QueryItemsResponseDataItems {
	s.CreateTime = &v
	return s
}

func (s *QueryItemsResponseDataItems) SetTitle(v string) *QueryItemsResponseDataItems {
	s.Title = &v
	return s
}

func (s *QueryItemsResponseDataItems) SetItemId(v string) *QueryItemsResponseDataItems {
	s.ItemId = &v
	return s
}

func (s *QueryItemsResponseDataItems) SetSellerId(v string) *QueryItemsResponseDataItems {
	s.SellerId = &v
	return s
}

func (s *QueryItemsResponseDataItems) SetSupplier(v string) *QueryItemsResponseDataItems {
	s.Supplier = &v
	return s
}

func (s *QueryItemsResponseDataItems) SetCategoryId(v int64) *QueryItemsResponseDataItems {
	s.CategoryId = &v
	return s
}

func (s *QueryItemsResponseDataItems) SetStatus(v int) *QueryItemsResponseDataItems {
	s.Status = &v
	return s
}

func (s *QueryItemsResponseDataItems) SetSkus(v []*QueryItemsResponseDataItemsSkus) *QueryItemsResponseDataItems {
	s.Skus = v
	return s
}

func (s *QueryItemsResponseDataItems) SetGames(v []*QueryItemsResponseDataItemsGames) *QueryItemsResponseDataItems {
	s.Games = v
	return s
}

type QueryItemsResponseDataItemsSkus struct {
	ItemId      *string                                     `json:"ItemId,omitempty" xml:"ItemId,omitempty" require:"true"`
	ModifyTime  *int64                                      `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	SalePrice   *int64                                      `json:"SalePrice,omitempty" xml:"SalePrice,omitempty" require:"true"`
	OriginPrice *int64                                      `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty" require:"true"`
	CreateTime  *int64                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	SkuId       *string                                     `json:"SkuId,omitempty" xml:"SkuId,omitempty" require:"true"`
	Status      *int                                        `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	SaleProps   []*QueryItemsResponseDataItemsSkusSaleProps `json:"SaleProps,omitempty" xml:"SaleProps,omitempty" require:"true" type:"Repeated"`
}

func (s QueryItemsResponseDataItemsSkus) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponseDataItemsSkus) GoString() string {
	return s.String()
}

func (s *QueryItemsResponseDataItemsSkus) SetItemId(v string) *QueryItemsResponseDataItemsSkus {
	s.ItemId = &v
	return s
}

func (s *QueryItemsResponseDataItemsSkus) SetModifyTime(v int64) *QueryItemsResponseDataItemsSkus {
	s.ModifyTime = &v
	return s
}

func (s *QueryItemsResponseDataItemsSkus) SetSalePrice(v int64) *QueryItemsResponseDataItemsSkus {
	s.SalePrice = &v
	return s
}

func (s *QueryItemsResponseDataItemsSkus) SetOriginPrice(v int64) *QueryItemsResponseDataItemsSkus {
	s.OriginPrice = &v
	return s
}

func (s *QueryItemsResponseDataItemsSkus) SetCreateTime(v int64) *QueryItemsResponseDataItemsSkus {
	s.CreateTime = &v
	return s
}

func (s *QueryItemsResponseDataItemsSkus) SetSkuId(v string) *QueryItemsResponseDataItemsSkus {
	s.SkuId = &v
	return s
}

func (s *QueryItemsResponseDataItemsSkus) SetStatus(v int) *QueryItemsResponseDataItemsSkus {
	s.Status = &v
	return s
}

func (s *QueryItemsResponseDataItemsSkus) SetSaleProps(v []*QueryItemsResponseDataItemsSkusSaleProps) *QueryItemsResponseDataItemsSkus {
	s.SaleProps = v
	return s
}

type QueryItemsResponseDataItemsSkusSaleProps struct {
	ValueId      *int64  `json:"ValueId,omitempty" xml:"ValueId,omitempty" require:"true"`
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty" require:"true"`
	PropertyId   *int64  `json:"PropertyId,omitempty" xml:"PropertyId,omitempty" require:"true"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s QueryItemsResponseDataItemsSkusSaleProps) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponseDataItemsSkusSaleProps) GoString() string {
	return s.String()
}

func (s *QueryItemsResponseDataItemsSkusSaleProps) SetValueId(v int64) *QueryItemsResponseDataItemsSkusSaleProps {
	s.ValueId = &v
	return s
}

func (s *QueryItemsResponseDataItemsSkusSaleProps) SetPropertyName(v string) *QueryItemsResponseDataItemsSkusSaleProps {
	s.PropertyName = &v
	return s
}

func (s *QueryItemsResponseDataItemsSkusSaleProps) SetPropertyId(v int64) *QueryItemsResponseDataItemsSkusSaleProps {
	s.PropertyId = &v
	return s
}

func (s *QueryItemsResponseDataItemsSkusSaleProps) SetValue(v string) *QueryItemsResponseDataItemsSkusSaleProps {
	s.Value = &v
	return s
}

type QueryItemsResponseDataItemsGames struct {
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s QueryItemsResponseDataItemsGames) String() string {
	return tea.Prettify(s)
}

func (s QueryItemsResponseDataItemsGames) GoString() string {
	return s.String()
}

func (s *QueryItemsResponseDataItemsGames) SetGameId(v string) *QueryItemsResponseDataItemsGames {
	s.GameId = &v
	return s
}

func (s *QueryItemsResponseDataItemsGames) SetName(v string) *QueryItemsResponseDataItemsGames {
	s.Name = &v
	return s
}

type GetItemRequest struct {
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty" require:"true"`
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

type GetItemResponse struct {
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *GetItemResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s GetItemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetItemResponse) GoString() string {
	return s.String()
}

func (s *GetItemResponse) SetRequestId(v string) *GetItemResponse {
	s.RequestId = &v
	return s
}

func (s *GetItemResponse) SetData(v *GetItemResponseData) *GetItemResponse {
	s.Data = v
	return s
}

type GetItemResponseData struct {
	ItemId      *string                     `json:"ItemId,omitempty" xml:"ItemId,omitempty" require:"true"`
	CreateTime  *int64                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	ModifyTime  *int64                      `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	SellerId    *string                     `json:"SellerId,omitempty" xml:"SellerId,omitempty" require:"true"`
	CategoryId  *int64                      `json:"CategoryId,omitempty" xml:"CategoryId,omitempty" require:"true"`
	Title       *string                     `json:"Title,omitempty" xml:"Title,omitempty" require:"true"`
	OriginPrice *int64                      `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty" require:"true"`
	SalePrice   *int64                      `json:"SalePrice,omitempty" xml:"SalePrice,omitempty" require:"true"`
	Status      *int                        `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	Description *string                     `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Supplier    *string                     `json:"Supplier,omitempty" xml:"Supplier,omitempty" require:"true"`
	Games       []*GetItemResponseDataGames `json:"Games,omitempty" xml:"Games,omitempty" require:"true" type:"Repeated"`
	Skus        []*GetItemResponseDataSkus  `json:"Skus,omitempty" xml:"Skus,omitempty" require:"true" type:"Repeated"`
}

func (s GetItemResponseData) String() string {
	return tea.Prettify(s)
}

func (s GetItemResponseData) GoString() string {
	return s.String()
}

func (s *GetItemResponseData) SetItemId(v string) *GetItemResponseData {
	s.ItemId = &v
	return s
}

func (s *GetItemResponseData) SetCreateTime(v int64) *GetItemResponseData {
	s.CreateTime = &v
	return s
}

func (s *GetItemResponseData) SetModifyTime(v int64) *GetItemResponseData {
	s.ModifyTime = &v
	return s
}

func (s *GetItemResponseData) SetSellerId(v string) *GetItemResponseData {
	s.SellerId = &v
	return s
}

func (s *GetItemResponseData) SetCategoryId(v int64) *GetItemResponseData {
	s.CategoryId = &v
	return s
}

func (s *GetItemResponseData) SetTitle(v string) *GetItemResponseData {
	s.Title = &v
	return s
}

func (s *GetItemResponseData) SetOriginPrice(v int64) *GetItemResponseData {
	s.OriginPrice = &v
	return s
}

func (s *GetItemResponseData) SetSalePrice(v int64) *GetItemResponseData {
	s.SalePrice = &v
	return s
}

func (s *GetItemResponseData) SetStatus(v int) *GetItemResponseData {
	s.Status = &v
	return s
}

func (s *GetItemResponseData) SetDescription(v string) *GetItemResponseData {
	s.Description = &v
	return s
}

func (s *GetItemResponseData) SetSupplier(v string) *GetItemResponseData {
	s.Supplier = &v
	return s
}

func (s *GetItemResponseData) SetGames(v []*GetItemResponseDataGames) *GetItemResponseData {
	s.Games = v
	return s
}

func (s *GetItemResponseData) SetSkus(v []*GetItemResponseDataSkus) *GetItemResponseData {
	s.Skus = v
	return s
}

type GetItemResponseDataGames struct {
	GameId *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetItemResponseDataGames) String() string {
	return tea.Prettify(s)
}

func (s GetItemResponseDataGames) GoString() string {
	return s.String()
}

func (s *GetItemResponseDataGames) SetGameId(v string) *GetItemResponseDataGames {
	s.GameId = &v
	return s
}

func (s *GetItemResponseDataGames) SetName(v string) *GetItemResponseDataGames {
	s.Name = &v
	return s
}

type GetItemResponseDataSkus struct {
	SkuId       *string                             `json:"SkuId,omitempty" xml:"SkuId,omitempty" require:"true"`
	CreateTime  *int64                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	ModifyTime  *int64                              `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty" require:"true"`
	ItemId      *string                             `json:"ItemId,omitempty" xml:"ItemId,omitempty" require:"true"`
	OriginPrice *int64                              `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty" require:"true"`
	SalePrice   *int64                              `json:"SalePrice,omitempty" xml:"SalePrice,omitempty" require:"true"`
	Status      *int                                `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	SaleProps   []*GetItemResponseDataSkusSaleProps `json:"SaleProps,omitempty" xml:"SaleProps,omitempty" require:"true" type:"Repeated"`
}

func (s GetItemResponseDataSkus) String() string {
	return tea.Prettify(s)
}

func (s GetItemResponseDataSkus) GoString() string {
	return s.String()
}

func (s *GetItemResponseDataSkus) SetSkuId(v string) *GetItemResponseDataSkus {
	s.SkuId = &v
	return s
}

func (s *GetItemResponseDataSkus) SetCreateTime(v int64) *GetItemResponseDataSkus {
	s.CreateTime = &v
	return s
}

func (s *GetItemResponseDataSkus) SetModifyTime(v int64) *GetItemResponseDataSkus {
	s.ModifyTime = &v
	return s
}

func (s *GetItemResponseDataSkus) SetItemId(v string) *GetItemResponseDataSkus {
	s.ItemId = &v
	return s
}

func (s *GetItemResponseDataSkus) SetOriginPrice(v int64) *GetItemResponseDataSkus {
	s.OriginPrice = &v
	return s
}

func (s *GetItemResponseDataSkus) SetSalePrice(v int64) *GetItemResponseDataSkus {
	s.SalePrice = &v
	return s
}

func (s *GetItemResponseDataSkus) SetStatus(v int) *GetItemResponseDataSkus {
	s.Status = &v
	return s
}

func (s *GetItemResponseDataSkus) SetSaleProps(v []*GetItemResponseDataSkusSaleProps) *GetItemResponseDataSkus {
	s.SaleProps = v
	return s
}

type GetItemResponseDataSkusSaleProps struct {
	PropertyId   *int64  `json:"PropertyId,omitempty" xml:"PropertyId,omitempty" require:"true"`
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty" require:"true"`
	ValueId      *int64  `json:"ValueId,omitempty" xml:"ValueId,omitempty" require:"true"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s GetItemResponseDataSkusSaleProps) String() string {
	return tea.Prettify(s)
}

func (s GetItemResponseDataSkusSaleProps) GoString() string {
	return s.String()
}

func (s *GetItemResponseDataSkusSaleProps) SetPropertyId(v int64) *GetItemResponseDataSkusSaleProps {
	s.PropertyId = &v
	return s
}

func (s *GetItemResponseDataSkusSaleProps) SetPropertyName(v string) *GetItemResponseDataSkusSaleProps {
	s.PropertyName = &v
	return s
}

func (s *GetItemResponseDataSkusSaleProps) SetValueId(v int64) *GetItemResponseDataSkusSaleProps {
	s.ValueId = &v
	return s
}

func (s *GetItemResponseDataSkusSaleProps) SetValue(v string) *GetItemResponseDataSkusSaleProps {
	s.Value = &v
	return s
}

type DeliveryOrderRequest struct {
	BuyerAccountId *string `json:"BuyerAccountId,omitempty" xml:"BuyerAccountId,omitempty" require:"true"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
	AccountDomain  *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
}

func (s DeliveryOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeliveryOrderRequest) GoString() string {
	return s.String()
}

func (s *DeliveryOrderRequest) SetBuyerAccountId(v string) *DeliveryOrderRequest {
	s.BuyerAccountId = &v
	return s
}

func (s *DeliveryOrderRequest) SetOrderId(v string) *DeliveryOrderRequest {
	s.OrderId = &v
	return s
}

func (s *DeliveryOrderRequest) SetAccountDomain(v string) *DeliveryOrderRequest {
	s.AccountDomain = &v
	return s
}

type DeliveryOrderResponse struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *DeliveryOrderResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s DeliveryOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeliveryOrderResponse) GoString() string {
	return s.String()
}

func (s *DeliveryOrderResponse) SetRequestId(v string) *DeliveryOrderResponse {
	s.RequestId = &v
	return s
}

func (s *DeliveryOrderResponse) SetData(v *DeliveryOrderResponseData) *DeliveryOrderResponse {
	s.Data = v
	return s
}

type DeliveryOrderResponseData struct {
	DeliveryStatus *string `json:"DeliveryStatus,omitempty" xml:"DeliveryStatus,omitempty" require:"true"`
}

func (s DeliveryOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s DeliveryOrderResponseData) GoString() string {
	return s.String()
}

func (s *DeliveryOrderResponseData) SetDeliveryStatus(v string) *DeliveryOrderResponseData {
	s.DeliveryStatus = &v
	return s
}

type CreateOrderRequest struct {
	BuyerAccountId  *string `json:"BuyerAccountId,omitempty" xml:"BuyerAccountId,omitempty" require:"true"`
	ItemId          *string `json:"ItemId,omitempty" xml:"ItemId,omitempty" require:"true"`
	SkuId           *string `json:"SkuId,omitempty" xml:"SkuId,omitempty" require:"true"`
	OriginPrice     *int64  `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty" require:"true"`
	SettlementPrice *int64  `json:"SettlementPrice,omitempty" xml:"SettlementPrice,omitempty" require:"true"`
	Amount          *int64  `json:"Amount,omitempty" xml:"Amount,omitempty" require:"true"`
	GameIdList      *string `json:"GameIdList,omitempty" xml:"GameIdList,omitempty" require:"true"`
	IdempotentCode  *string `json:"IdempotentCode,omitempty" xml:"IdempotentCode,omitempty" require:"true"`
	AccountDomain   *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) SetBuyerAccountId(v string) *CreateOrderRequest {
	s.BuyerAccountId = &v
	return s
}

func (s *CreateOrderRequest) SetItemId(v string) *CreateOrderRequest {
	s.ItemId = &v
	return s
}

func (s *CreateOrderRequest) SetSkuId(v string) *CreateOrderRequest {
	s.SkuId = &v
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

func (s *CreateOrderRequest) SetAmount(v int64) *CreateOrderRequest {
	s.Amount = &v
	return s
}

func (s *CreateOrderRequest) SetGameIdList(v string) *CreateOrderRequest {
	s.GameIdList = &v
	return s
}

func (s *CreateOrderRequest) SetIdempotentCode(v string) *CreateOrderRequest {
	s.IdempotentCode = &v
	return s
}

func (s *CreateOrderRequest) SetAccountDomain(v string) *CreateOrderRequest {
	s.AccountDomain = &v
	return s
}

type CreateOrderResponse struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *CreateOrderResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s CreateOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderResponse) SetRequestId(v string) *CreateOrderResponse {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponse) SetData(v *CreateOrderResponseData) *CreateOrderResponse {
	s.Data = v
	return s
}

type CreateOrderResponseData struct {
	OrderId           *string `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
	ItemId            *string `json:"ItemId,omitempty" xml:"ItemId,omitempty" require:"true"`
	SkuId             *string `json:"SkuId,omitempty" xml:"SkuId,omitempty" require:"true"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	ApplyDeliveryTime *int64  `json:"ApplyDeliveryTime,omitempty" xml:"ApplyDeliveryTime,omitempty" require:"true"`
	FinishTime        *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty" require:"true"`
	OriginPrice       *int64  `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty" require:"true"`
	SettlementPrice   *int64  `json:"SettlementPrice,omitempty" xml:"SettlementPrice,omitempty" require:"true"`
	Amount            *int64  `json:"Amount,omitempty" xml:"Amount,omitempty" require:"true"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	BuyerAccountId    *string `json:"BuyerAccountId,omitempty" xml:"BuyerAccountId,omitempty" require:"true"`
	AutoUnlockTime    *int64  `json:"AutoUnlockTime,omitempty" xml:"AutoUnlockTime,omitempty" require:"true"`
	AccountDomain     *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty" require:"true"`
}

func (s CreateOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseData) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseData) SetOrderId(v string) *CreateOrderResponseData {
	s.OrderId = &v
	return s
}

func (s *CreateOrderResponseData) SetItemId(v string) *CreateOrderResponseData {
	s.ItemId = &v
	return s
}

func (s *CreateOrderResponseData) SetSkuId(v string) *CreateOrderResponseData {
	s.SkuId = &v
	return s
}

func (s *CreateOrderResponseData) SetCreateTime(v int64) *CreateOrderResponseData {
	s.CreateTime = &v
	return s
}

func (s *CreateOrderResponseData) SetApplyDeliveryTime(v int64) *CreateOrderResponseData {
	s.ApplyDeliveryTime = &v
	return s
}

func (s *CreateOrderResponseData) SetFinishTime(v int64) *CreateOrderResponseData {
	s.FinishTime = &v
	return s
}

func (s *CreateOrderResponseData) SetOriginPrice(v int64) *CreateOrderResponseData {
	s.OriginPrice = &v
	return s
}

func (s *CreateOrderResponseData) SetSettlementPrice(v int64) *CreateOrderResponseData {
	s.SettlementPrice = &v
	return s
}

func (s *CreateOrderResponseData) SetAmount(v int64) *CreateOrderResponseData {
	s.Amount = &v
	return s
}

func (s *CreateOrderResponseData) SetStatus(v string) *CreateOrderResponseData {
	s.Status = &v
	return s
}

func (s *CreateOrderResponseData) SetBuyerAccountId(v string) *CreateOrderResponseData {
	s.BuyerAccountId = &v
	return s
}

func (s *CreateOrderResponseData) SetAutoUnlockTime(v int64) *CreateOrderResponseData {
	s.AutoUnlockTime = &v
	return s
}

func (s *CreateOrderResponseData) SetAccountDomain(v string) *CreateOrderResponseData {
	s.AccountDomain = &v
	return s
}

type QueryOrderRequest struct {
	BuyerAccountId *string `json:"BuyerAccountId,omitempty" xml:"BuyerAccountId,omitempty" require:"true"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
	AccountDomain  *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
}

func (s QueryOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderRequest) SetBuyerAccountId(v string) *QueryOrderRequest {
	s.BuyerAccountId = &v
	return s
}

func (s *QueryOrderRequest) SetOrderId(v string) *QueryOrderRequest {
	s.OrderId = &v
	return s
}

func (s *QueryOrderRequest) SetAccountDomain(v string) *QueryOrderRequest {
	s.AccountDomain = &v
	return s
}

type QueryOrderResponse struct {
	RequestId      *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DeliveryStatus *string                 `json:"DeliveryStatus,omitempty" xml:"DeliveryStatus,omitempty" require:"true"`
	RefundStatus   *string                 `json:"RefundStatus,omitempty" xml:"RefundStatus,omitempty" require:"true"`
	Data           *QueryOrderResponseData `json:"Data,omitempty" xml:"Data,omitempty" require:"true" type:"Struct"`
}

func (s QueryOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderResponse) SetRequestId(v string) *QueryOrderResponse {
	s.RequestId = &v
	return s
}

func (s *QueryOrderResponse) SetDeliveryStatus(v string) *QueryOrderResponse {
	s.DeliveryStatus = &v
	return s
}

func (s *QueryOrderResponse) SetRefundStatus(v string) *QueryOrderResponse {
	s.RefundStatus = &v
	return s
}

func (s *QueryOrderResponse) SetData(v *QueryOrderResponseData) *QueryOrderResponse {
	s.Data = v
	return s
}

type QueryOrderResponseData struct {
	OrderId           *string `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
	ItemId            *string `json:"ItemId,omitempty" xml:"ItemId,omitempty" require:"true"`
	SkuId             *string `json:"SkuId,omitempty" xml:"SkuId,omitempty" require:"true"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	ApplyDeliveryTime *int64  `json:"ApplyDeliveryTime,omitempty" xml:"ApplyDeliveryTime,omitempty" require:"true"`
	FinishTime        *int64  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty" require:"true"`
	OriginPrice       *int64  `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty" require:"true"`
	SettlementPrice   *int64  `json:"SettlementPrice,omitempty" xml:"SettlementPrice,omitempty" require:"true"`
	Amount            *int64  `json:"Amount,omitempty" xml:"Amount,omitempty" require:"true"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	BuyerAccountId    *string `json:"BuyerAccountId,omitempty" xml:"BuyerAccountId,omitempty" require:"true"`
	AutoUnlockTime    *int64  `json:"AutoUnlockTime,omitempty" xml:"AutoUnlockTime,omitempty" require:"true"`
	AccountDomain     *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty" require:"true"`
}

func (s QueryOrderResponseData) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderResponseData) GoString() string {
	return s.String()
}

func (s *QueryOrderResponseData) SetOrderId(v string) *QueryOrderResponseData {
	s.OrderId = &v
	return s
}

func (s *QueryOrderResponseData) SetItemId(v string) *QueryOrderResponseData {
	s.ItemId = &v
	return s
}

func (s *QueryOrderResponseData) SetSkuId(v string) *QueryOrderResponseData {
	s.SkuId = &v
	return s
}

func (s *QueryOrderResponseData) SetCreateTime(v int64) *QueryOrderResponseData {
	s.CreateTime = &v
	return s
}

func (s *QueryOrderResponseData) SetApplyDeliveryTime(v int64) *QueryOrderResponseData {
	s.ApplyDeliveryTime = &v
	return s
}

func (s *QueryOrderResponseData) SetFinishTime(v int64) *QueryOrderResponseData {
	s.FinishTime = &v
	return s
}

func (s *QueryOrderResponseData) SetOriginPrice(v int64) *QueryOrderResponseData {
	s.OriginPrice = &v
	return s
}

func (s *QueryOrderResponseData) SetSettlementPrice(v int64) *QueryOrderResponseData {
	s.SettlementPrice = &v
	return s
}

func (s *QueryOrderResponseData) SetAmount(v int64) *QueryOrderResponseData {
	s.Amount = &v
	return s
}

func (s *QueryOrderResponseData) SetStatus(v string) *QueryOrderResponseData {
	s.Status = &v
	return s
}

func (s *QueryOrderResponseData) SetBuyerAccountId(v string) *QueryOrderResponseData {
	s.BuyerAccountId = &v
	return s
}

func (s *QueryOrderResponseData) SetAutoUnlockTime(v int64) *QueryOrderResponseData {
	s.AutoUnlockTime = &v
	return s
}

func (s *QueryOrderResponseData) SetAccountDomain(v string) *QueryOrderResponseData {
	s.AccountDomain = &v
	return s
}

type ListBoughtGamesRequest struct {
	AccountId     *string `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	AccountDomain *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
	PageNumber    *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListBoughtGamesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBoughtGamesRequest) GoString() string {
	return s.String()
}

func (s *ListBoughtGamesRequest) SetAccountId(v string) *ListBoughtGamesRequest {
	s.AccountId = &v
	return s
}

func (s *ListBoughtGamesRequest) SetAccountDomain(v string) *ListBoughtGamesRequest {
	s.AccountDomain = &v
	return s
}

func (s *ListBoughtGamesRequest) SetPageNumber(v int) *ListBoughtGamesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBoughtGamesRequest) SetPageSize(v int) *ListBoughtGamesRequest {
	s.PageSize = &v
	return s
}

type ListBoughtGamesResponse struct {
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                            `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	Items      []*ListBoughtGamesResponseItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Repeated"`
}

func (s ListBoughtGamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBoughtGamesResponse) GoString() string {
	return s.String()
}

func (s *ListBoughtGamesResponse) SetRequestId(v string) *ListBoughtGamesResponse {
	s.RequestId = &v
	return s
}

func (s *ListBoughtGamesResponse) SetPageNumber(v int) *ListBoughtGamesResponse {
	s.PageNumber = &v
	return s
}

func (s *ListBoughtGamesResponse) SetPageSize(v int) *ListBoughtGamesResponse {
	s.PageSize = &v
	return s
}

func (s *ListBoughtGamesResponse) SetTotalCount(v int) *ListBoughtGamesResponse {
	s.TotalCount = &v
	return s
}

func (s *ListBoughtGamesResponse) SetItems(v []*ListBoughtGamesResponseItems) *ListBoughtGamesResponse {
	s.Items = v
	return s
}

type ListBoughtGamesResponseItems struct {
	GameId    *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	GameName  *string `json:"GameName,omitempty" xml:"GameName,omitempty" require:"true"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s ListBoughtGamesResponseItems) String() string {
	return tea.Prettify(s)
}

func (s ListBoughtGamesResponseItems) GoString() string {
	return s.String()
}

func (s *ListBoughtGamesResponseItems) SetGameId(v string) *ListBoughtGamesResponseItems {
	s.GameId = &v
	return s
}

func (s *ListBoughtGamesResponseItems) SetGameName(v string) *ListBoughtGamesResponseItems {
	s.GameName = &v
	return s
}

func (s *ListBoughtGamesResponseItems) SetStartTime(v int64) *ListBoughtGamesResponseItems {
	s.StartTime = &v
	return s
}

func (s *ListBoughtGamesResponseItems) SetEndTime(v int64) *ListBoughtGamesResponseItems {
	s.EndTime = &v
	return s
}

type CloseOrderRequest struct {
	BuyerAccountId *string `json:"BuyerAccountId,omitempty" xml:"BuyerAccountId,omitempty" require:"true"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
	AccountDomain  *string `json:"AccountDomain,omitempty" xml:"AccountDomain,omitempty"`
}

func (s CloseOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseOrderRequest) GoString() string {
	return s.String()
}

func (s *CloseOrderRequest) SetBuyerAccountId(v string) *CloseOrderRequest {
	s.BuyerAccountId = &v
	return s
}

func (s *CloseOrderRequest) SetOrderId(v string) *CloseOrderRequest {
	s.OrderId = &v
	return s
}

func (s *CloseOrderRequest) SetAccountDomain(v string) *CloseOrderRequest {
	s.AccountDomain = &v
	return s
}

type CloseOrderResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s CloseOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseOrderResponse) GoString() string {
	return s.String()
}

func (s *CloseOrderResponse) SetRequestId(v string) *CloseOrderResponse {
	s.RequestId = &v
	return s
}

func (s *CloseOrderResponse) SetData(v bool) *CloseOrderResponse {
	s.Data = &v
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

type BatchDispatchGameSlotResponse struct {
	RequestId       *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	QueueResultList []*BatchDispatchGameSlotResponseQueueResultList `json:"QueueResultList,omitempty" xml:"QueueResultList,omitempty" require:"true" type:"Repeated"`
}

func (s BatchDispatchGameSlotResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDispatchGameSlotResponse) GoString() string {
	return s.String()
}

func (s *BatchDispatchGameSlotResponse) SetRequestId(v string) *BatchDispatchGameSlotResponse {
	s.RequestId = &v
	return s
}

func (s *BatchDispatchGameSlotResponse) SetQueueResultList(v []*BatchDispatchGameSlotResponseQueueResultList) *BatchDispatchGameSlotResponse {
	s.QueueResultList = v
	return s
}

type BatchDispatchGameSlotResponseQueueResultList struct {
	GameId      *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	GameSession *string `json:"GameSession,omitempty" xml:"GameSession,omitempty" require:"true"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	QueueCode   *int    `json:"QueueCode,omitempty" xml:"QueueCode,omitempty" require:"true"`
	QueueState  *int    `json:"QueueState,omitempty" xml:"QueueState,omitempty" require:"true"`
	RegionName  *string `json:"RegionName,omitempty" xml:"RegionName,omitempty" require:"true"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
}

func (s BatchDispatchGameSlotResponseQueueResultList) String() string {
	return tea.Prettify(s)
}

func (s BatchDispatchGameSlotResponseQueueResultList) GoString() string {
	return s.String()
}

func (s *BatchDispatchGameSlotResponseQueueResultList) SetGameId(v string) *BatchDispatchGameSlotResponseQueueResultList {
	s.GameId = &v
	return s
}

func (s *BatchDispatchGameSlotResponseQueueResultList) SetGameSession(v string) *BatchDispatchGameSlotResponseQueueResultList {
	s.GameSession = &v
	return s
}

func (s *BatchDispatchGameSlotResponseQueueResultList) SetMessage(v string) *BatchDispatchGameSlotResponseQueueResultList {
	s.Message = &v
	return s
}

func (s *BatchDispatchGameSlotResponseQueueResultList) SetQueueCode(v int) *BatchDispatchGameSlotResponseQueueResultList {
	s.QueueCode = &v
	return s
}

func (s *BatchDispatchGameSlotResponseQueueResultList) SetQueueState(v int) *BatchDispatchGameSlotResponseQueueResultList {
	s.QueueState = &v
	return s
}

func (s *BatchDispatchGameSlotResponseQueueResultList) SetRegionName(v string) *BatchDispatchGameSlotResponseQueueResultList {
	s.RegionName = &v
	return s
}

func (s *BatchDispatchGameSlotResponseQueueResultList) SetUserId(v string) *BatchDispatchGameSlotResponseQueueResultList {
	s.UserId = &v
	return s
}

type StopGameSessionRequest struct {
	GameId      *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	AccessKey   *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" require:"true"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	BizParam    *string `json:"BizParam,omitempty" xml:"BizParam,omitempty"`
	GameSession *string `json:"GameSession,omitempty" xml:"GameSession,omitempty" require:"true"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s StopGameSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s StopGameSessionRequest) GoString() string {
	return s.String()
}

func (s *StopGameSessionRequest) SetGameId(v string) *StopGameSessionRequest {
	s.GameId = &v
	return s
}

func (s *StopGameSessionRequest) SetAccessKey(v string) *StopGameSessionRequest {
	s.AccessKey = &v
	return s
}

func (s *StopGameSessionRequest) SetUserId(v string) *StopGameSessionRequest {
	s.UserId = &v
	return s
}

func (s *StopGameSessionRequest) SetBizParam(v string) *StopGameSessionRequest {
	s.BizParam = &v
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

type StopGameSessionResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	GameId      *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	GameSession *string `json:"GameSession,omitempty" xml:"GameSession,omitempty" require:"true"`
	QueueState  *int    `json:"QueueState,omitempty" xml:"QueueState,omitempty" require:"true"`
	QueueCode   *int    `json:"QueueCode,omitempty" xml:"QueueCode,omitempty" require:"true"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Success     *bool   `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
}

func (s StopGameSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s StopGameSessionResponse) GoString() string {
	return s.String()
}

func (s *StopGameSessionResponse) SetRequestId(v string) *StopGameSessionResponse {
	s.RequestId = &v
	return s
}

func (s *StopGameSessionResponse) SetGameId(v string) *StopGameSessionResponse {
	s.GameId = &v
	return s
}

func (s *StopGameSessionResponse) SetGameSession(v string) *StopGameSessionResponse {
	s.GameSession = &v
	return s
}

func (s *StopGameSessionResponse) SetQueueState(v int) *StopGameSessionResponse {
	s.QueueState = &v
	return s
}

func (s *StopGameSessionResponse) SetQueueCode(v int) *StopGameSessionResponse {
	s.QueueCode = &v
	return s
}

func (s *StopGameSessionResponse) SetMessage(v string) *StopGameSessionResponse {
	s.Message = &v
	return s
}

func (s *StopGameSessionResponse) SetSuccess(v bool) *StopGameSessionResponse {
	s.Success = &v
	return s
}

type DispatchGameSlotRequest struct {
	GameId         *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	AccessKey      *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" require:"true"`
	RegionName     *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	BizParam       *string `json:"BizParam,omitempty" xml:"BizParam,omitempty"`
	Cancel         *bool   `json:"Cancel,omitempty" xml:"Cancel,omitempty"`
	GameSession    *string `json:"GameSession,omitempty" xml:"GameSession,omitempty"`
	GameStartParam *string `json:"GameStartParam,omitempty" xml:"GameStartParam,omitempty"`
	GameCommand    *string `json:"GameCommand,omitempty" xml:"GameCommand,omitempty"`
	SystemInfo     *string `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty"`
	ClientIp       *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	Reconnect      *bool   `json:"Reconnect,omitempty" xml:"Reconnect,omitempty"`
}

func (s DispatchGameSlotRequest) String() string {
	return tea.Prettify(s)
}

func (s DispatchGameSlotRequest) GoString() string {
	return s.String()
}

func (s *DispatchGameSlotRequest) SetGameId(v string) *DispatchGameSlotRequest {
	s.GameId = &v
	return s
}

func (s *DispatchGameSlotRequest) SetAccessKey(v string) *DispatchGameSlotRequest {
	s.AccessKey = &v
	return s
}

func (s *DispatchGameSlotRequest) SetRegionName(v string) *DispatchGameSlotRequest {
	s.RegionName = &v
	return s
}

func (s *DispatchGameSlotRequest) SetUserId(v string) *DispatchGameSlotRequest {
	s.UserId = &v
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

func (s *DispatchGameSlotRequest) SetGameSession(v string) *DispatchGameSlotRequest {
	s.GameSession = &v
	return s
}

func (s *DispatchGameSlotRequest) SetGameStartParam(v string) *DispatchGameSlotRequest {
	s.GameStartParam = &v
	return s
}

func (s *DispatchGameSlotRequest) SetGameCommand(v string) *DispatchGameSlotRequest {
	s.GameCommand = &v
	return s
}

func (s *DispatchGameSlotRequest) SetSystemInfo(v string) *DispatchGameSlotRequest {
	s.SystemInfo = &v
	return s
}

func (s *DispatchGameSlotRequest) SetClientIp(v string) *DispatchGameSlotRequest {
	s.ClientIp = &v
	return s
}

func (s *DispatchGameSlotRequest) SetReconnect(v bool) *DispatchGameSlotRequest {
	s.Reconnect = &v
	return s
}

type DispatchGameSlotResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	GameId      *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	GameSession *string `json:"GameSession,omitempty" xml:"GameSession,omitempty" require:"true"`
	QueueState  *int    `json:"QueueState,omitempty" xml:"QueueState,omitempty" require:"true"`
	QueueCode   *int    `json:"QueueCode,omitempty" xml:"QueueCode,omitempty" require:"true"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	RegionName  *string `json:"RegionName,omitempty" xml:"RegionName,omitempty" require:"true"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
}

func (s DispatchGameSlotResponse) String() string {
	return tea.Prettify(s)
}

func (s DispatchGameSlotResponse) GoString() string {
	return s.String()
}

func (s *DispatchGameSlotResponse) SetRequestId(v string) *DispatchGameSlotResponse {
	s.RequestId = &v
	return s
}

func (s *DispatchGameSlotResponse) SetGameId(v string) *DispatchGameSlotResponse {
	s.GameId = &v
	return s
}

func (s *DispatchGameSlotResponse) SetGameSession(v string) *DispatchGameSlotResponse {
	s.GameSession = &v
	return s
}

func (s *DispatchGameSlotResponse) SetQueueState(v int) *DispatchGameSlotResponse {
	s.QueueState = &v
	return s
}

func (s *DispatchGameSlotResponse) SetQueueCode(v int) *DispatchGameSlotResponse {
	s.QueueCode = &v
	return s
}

func (s *DispatchGameSlotResponse) SetMessage(v string) *DispatchGameSlotResponse {
	s.Message = &v
	return s
}

func (s *DispatchGameSlotResponse) SetRegionName(v string) *DispatchGameSlotResponse {
	s.RegionName = &v
	return s
}

func (s *DispatchGameSlotResponse) SetUserId(v string) *DispatchGameSlotResponse {
	s.UserId = &v
	return s
}

type GetGameCcuRequest struct {
	GameId     *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	AccessKey  *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" require:"true"`
}

func (s GetGameCcuRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGameCcuRequest) GoString() string {
	return s.String()
}

func (s *GetGameCcuRequest) SetGameId(v string) *GetGameCcuRequest {
	s.GameId = &v
	return s
}

func (s *GetGameCcuRequest) SetRegionName(v string) *GetGameCcuRequest {
	s.RegionName = &v
	return s
}

func (s *GetGameCcuRequest) SetAccessKey(v string) *GetGameCcuRequest {
	s.AccessKey = &v
	return s
}

type GetGameCcuResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DataList  []*GetGameCcuResponseDataList `json:"DataList,omitempty" xml:"DataList,omitempty" require:"true" type:"Repeated"`
}

func (s GetGameCcuResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGameCcuResponse) GoString() string {
	return s.String()
}

func (s *GetGameCcuResponse) SetRequestId(v string) *GetGameCcuResponse {
	s.RequestId = &v
	return s
}

func (s *GetGameCcuResponse) SetDataList(v []*GetGameCcuResponseDataList) *GetGameCcuResponse {
	s.DataList = v
	return s
}

type GetGameCcuResponseDataList struct {
	Ccu      *int64  `json:"Ccu,omitempty" xml:"Ccu,omitempty" require:"true"`
	GameId   *string `json:"GameId,omitempty" xml:"GameId,omitempty" require:"true"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s GetGameCcuResponseDataList) String() string {
	return tea.Prettify(s)
}

func (s GetGameCcuResponseDataList) GoString() string {
	return s.String()
}

func (s *GetGameCcuResponseDataList) SetCcu(v int64) *GetGameCcuResponseDataList {
	s.Ccu = &v
	return s
}

func (s *GetGameCcuResponseDataList) SetGameId(v string) *GetGameCcuResponseDataList {
	s.GameId = &v
	return s
}

func (s *GetGameCcuResponseDataList) SetRegionId(v string) *GetGameCcuResponseDataList {
	s.RegionId = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
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

func (client *Client) QueryOutAccountBindStatusWithOptions(request *QueryOutAccountBindStatusRequest, runtime *util.RuntimeOptions) (_result *QueryOutAccountBindStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryOutAccountBindStatusResponse{}
	_body, _err := client.DoRequest(tea.String("QueryOutAccountBindStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateTokenWithOptions(request *CreateTokenRequest, runtime *util.RuntimeOptions) (_result *CreateTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateTokenResponse{}
	_body, _err := client.DoRequest(tea.String("CreateToken"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetSessionWithOptions(request *GetSessionRequest, runtime *util.RuntimeOptions) (_result *GetSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSessionResponse{}
	_body, _err := client.DoRequest(tea.String("GetSession"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) SkipTrialPolicyWithOptions(request *SkipTrialPolicyRequest, runtime *util.RuntimeOptions) (_result *SkipTrialPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SkipTrialPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("SkipTrialPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) QueryTenantWithOptions(request *QueryTenantRequest, runtime *util.RuntimeOptions) (_result *QueryTenantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryTenantResponse{}
	_body, _err := client.DoRequest(tea.String("QueryTenant"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) QueryProjectWithOptions(request *QueryProjectRequest, runtime *util.RuntimeOptions) (_result *QueryProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryProjectResponse{}
	_body, _err := client.DoRequest(tea.String("QueryProject"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) QueryGameWithOptions(request *QueryGameRequest, runtime *util.RuntimeOptions) (_result *QueryGameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryGameResponse{}
	_body, _err := client.DoRequest(tea.String("QueryGame"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) BatchStopGameSessionsWithOptions(request *BatchStopGameSessionsRequest, runtime *util.RuntimeOptions) (_result *BatchStopGameSessionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BatchStopGameSessionsResponse{}
	_body, _err := client.DoRequest(tea.String("BatchStopGameSessions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetStopGameTokenWithOptions(request *GetStopGameTokenRequest, runtime *util.RuntimeOptions) (_result *GetStopGameTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetStopGameTokenResponse{}
	_body, _err := client.DoRequest(tea.String("GetStopGameToken"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) QueryItemsWithOptions(request *QueryItemsRequest, runtime *util.RuntimeOptions) (_result *QueryItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryItemsResponse{}
	_body, _err := client.DoRequest(tea.String("QueryItems"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetItemWithOptions(request *GetItemRequest, runtime *util.RuntimeOptions) (_result *GetItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetItemResponse{}
	_body, _err := client.DoRequest(tea.String("GetItem"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeliveryOrderWithOptions(request *DeliveryOrderRequest, runtime *util.RuntimeOptions) (_result *DeliveryOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeliveryOrderResponse{}
	_body, _err := client.DoRequest(tea.String("DeliveryOrder"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CreateOrderWithOptions(request *CreateOrderRequest, runtime *util.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.DoRequest(tea.String("CreateOrder"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) QueryOrderWithOptions(request *QueryOrderRequest, runtime *util.RuntimeOptions) (_result *QueryOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryOrderResponse{}
	_body, _err := client.DoRequest(tea.String("QueryOrder"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) ListBoughtGamesWithOptions(request *ListBoughtGamesRequest, runtime *util.RuntimeOptions) (_result *ListBoughtGamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListBoughtGamesResponse{}
	_body, _err := client.DoRequest(tea.String("ListBoughtGames"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) CloseOrderWithOptions(request *CloseOrderRequest, runtime *util.RuntimeOptions) (_result *CloseOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CloseOrderResponse{}
	_body, _err := client.DoRequest(tea.String("CloseOrder"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) BatchDispatchGameSlotWithOptions(request *BatchDispatchGameSlotRequest, runtime *util.RuntimeOptions) (_result *BatchDispatchGameSlotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BatchDispatchGameSlotResponse{}
	_body, _err := client.DoRequest(tea.String("BatchDispatchGameSlot"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) StopGameSessionWithOptions(request *StopGameSessionRequest, runtime *util.RuntimeOptions) (_result *StopGameSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StopGameSessionResponse{}
	_body, _err := client.DoRequest(tea.String("StopGameSession"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DispatchGameSlotWithOptions(request *DispatchGameSlotRequest, runtime *util.RuntimeOptions) (_result *DispatchGameSlotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DispatchGameSlotResponse{}
	_body, _err := client.DoRequest(tea.String("DispatchGameSlot"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
	_result = &GetGameCcuResponse{}
	_body, _err := client.DoRequest(tea.String("GetGameCcu"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-07-28"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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
