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

type GetUserTokenRequest struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 用户id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 昵称
	Nick *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	// appKey
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s GetUserTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenRequest) GoString() string {
	return s.String()
}

func (s *GetUserTokenRequest) SetInstanceId(v string) *GetUserTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserTokenRequest) SetUserId(v string) *GetUserTokenRequest {
	s.UserId = &v
	return s
}

func (s *GetUserTokenRequest) SetNick(v string) *GetUserTokenRequest {
	s.Nick = &v
	return s
}

func (s *GetUserTokenRequest) SetAppKey(v string) *GetUserTokenRequest {
	s.AppKey = &v
	return s
}

type GetUserTokenResponseBody struct {
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 鹰眼id
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetUserTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 是否调用成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserTokenResponseBody) SetMessage(v string) *GetUserTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserTokenResponseBody) SetRequestId(v string) *GetUserTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserTokenResponseBody) SetData(v *GetUserTokenResponseBodyData) *GetUserTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetUserTokenResponseBody) SetCode(v string) *GetUserTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserTokenResponseBody) SetSuccess(v bool) *GetUserTokenResponseBody {
	s.Success = &v
	return s
}

type GetUserTokenResponseBodyData struct {
	Expires *int64  `json:"Expires,omitempty" xml:"Expires,omitempty"`
	Token   *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetUserTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserTokenResponseBodyData) SetExpires(v int64) *GetUserTokenResponseBodyData {
	s.Expires = &v
	return s
}

func (s *GetUserTokenResponseBodyData) SetToken(v string) *GetUserTokenResponseBodyData {
	s.Token = &v
	return s
}

type GetUserTokenResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenResponse) GoString() string {
	return s.String()
}

func (s *GetUserTokenResponse) SetHeaders(v map[string]*string) *GetUserTokenResponse {
	s.Headers = v
	return s
}

func (s *GetUserTokenResponse) SetBody(v *GetUserTokenResponseBody) *GetUserTokenResponse {
	s.Body = v
	return s
}

type SearchTicketListRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperatorId   *int64  `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	TicketStatus *string `json:"TicketStatus,omitempty" xml:"TicketStatus,omitempty"`
	PageNo       *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s SearchTicketListRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketListRequest) GoString() string {
	return s.String()
}

func (s *SearchTicketListRequest) SetClientToken(v string) *SearchTicketListRequest {
	s.ClientToken = &v
	return s
}

func (s *SearchTicketListRequest) SetInstanceId(v string) *SearchTicketListRequest {
	s.InstanceId = &v
	return s
}

func (s *SearchTicketListRequest) SetOperatorId(v int64) *SearchTicketListRequest {
	s.OperatorId = &v
	return s
}

func (s *SearchTicketListRequest) SetTicketStatus(v string) *SearchTicketListRequest {
	s.TicketStatus = &v
	return s
}

func (s *SearchTicketListRequest) SetPageNo(v int32) *SearchTicketListRequest {
	s.PageNo = &v
	return s
}

func (s *SearchTicketListRequest) SetPageSize(v int32) *SearchTicketListRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTicketListRequest) SetStartTime(v int64) *SearchTicketListRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTicketListRequest) SetEndTime(v int64) *SearchTicketListRequest {
	s.EndTime = &v
	return s
}

type SearchTicketListResponseBody struct {
	OnePageSize  *int32                              `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message      *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	TotalPage    *int32                              `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	TotalResults *int32                              `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
	PageNo       *int32                              `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	Data         []*SearchTicketListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchTicketListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketListResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTicketListResponseBody) SetOnePageSize(v int32) *SearchTicketListResponseBody {
	s.OnePageSize = &v
	return s
}

func (s *SearchTicketListResponseBody) SetRequestId(v string) *SearchTicketListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTicketListResponseBody) SetMessage(v string) *SearchTicketListResponseBody {
	s.Message = &v
	return s
}

func (s *SearchTicketListResponseBody) SetTotalPage(v int32) *SearchTicketListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *SearchTicketListResponseBody) SetTotalResults(v int32) *SearchTicketListResponseBody {
	s.TotalResults = &v
	return s
}

func (s *SearchTicketListResponseBody) SetPageNo(v int32) *SearchTicketListResponseBody {
	s.PageNo = &v
	return s
}

func (s *SearchTicketListResponseBody) SetData(v []*SearchTicketListResponseBodyData) *SearchTicketListResponseBody {
	s.Data = v
	return s
}

func (s *SearchTicketListResponseBody) SetCode(v string) *SearchTicketListResponseBody {
	s.Code = &v
	return s
}

func (s *SearchTicketListResponseBody) SetSuccess(v bool) *SearchTicketListResponseBody {
	s.Success = &v
	return s
}

type SearchTicketListResponseBodyData struct {
	MemberName   *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	CarbonCopy   *string `json:"CarbonCopy,omitempty" xml:"CarbonCopy,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ServiceId    *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	TicketId     *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	CreatorId    *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	FormData     *string `json:"FormData,omitempty" xml:"FormData,omitempty"`
	FromInfo     *string `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	ModifiedTime *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	TaskStatus   *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	CreatorName  *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	CategoryId   *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreatorType  *int32  `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	MemberId     *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	CaseStatus   *int32  `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SearchTicketListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketListResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchTicketListResponseBodyData) SetMemberName(v string) *SearchTicketListResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCarbonCopy(v string) *SearchTicketListResponseBodyData {
	s.CarbonCopy = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCreateTime(v int64) *SearchTicketListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetServiceId(v int64) *SearchTicketListResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetTicketId(v int64) *SearchTicketListResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetPriority(v int32) *SearchTicketListResponseBodyData {
	s.Priority = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCreatorId(v int64) *SearchTicketListResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetFormData(v string) *SearchTicketListResponseBodyData {
	s.FormData = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetFromInfo(v string) *SearchTicketListResponseBodyData {
	s.FromInfo = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetModifiedTime(v int64) *SearchTicketListResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetTaskStatus(v string) *SearchTicketListResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCreatorName(v string) *SearchTicketListResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCategoryId(v int64) *SearchTicketListResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCreatorType(v int32) *SearchTicketListResponseBodyData {
	s.CreatorType = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetMemberId(v int64) *SearchTicketListResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCaseStatus(v int32) *SearchTicketListResponseBodyData {
	s.CaseStatus = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetTemplateId(v int64) *SearchTicketListResponseBodyData {
	s.TemplateId = &v
	return s
}

type SearchTicketListResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchTicketListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTicketListResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketListResponse) GoString() string {
	return s.String()
}

func (s *SearchTicketListResponse) SetHeaders(v map[string]*string) *SearchTicketListResponse {
	s.Headers = v
	return s
}

func (s *SearchTicketListResponse) SetBody(v *SearchTicketListResponseBody) *SearchTicketListResponse {
	s.Body = v
	return s
}

type SendHotlineHeartBeatRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Token       *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s SendHotlineHeartBeatRequest) String() string {
	return tea.Prettify(s)
}

func (s SendHotlineHeartBeatRequest) GoString() string {
	return s.String()
}

func (s *SendHotlineHeartBeatRequest) SetClientToken(v string) *SendHotlineHeartBeatRequest {
	s.ClientToken = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) SetInstanceId(v string) *SendHotlineHeartBeatRequest {
	s.InstanceId = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) SetAccountName(v string) *SendHotlineHeartBeatRequest {
	s.AccountName = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) SetToken(v string) *SendHotlineHeartBeatRequest {
	s.Token = &v
	return s
}

type SendHotlineHeartBeatResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendHotlineHeartBeatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendHotlineHeartBeatResponseBody) GoString() string {
	return s.String()
}

func (s *SendHotlineHeartBeatResponseBody) SetMessage(v string) *SendHotlineHeartBeatResponseBody {
	s.Message = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) SetRequestId(v string) *SendHotlineHeartBeatResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) SetCode(v string) *SendHotlineHeartBeatResponseBody {
	s.Code = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) SetSuccess(v bool) *SendHotlineHeartBeatResponseBody {
	s.Success = &v
	return s
}

type SendHotlineHeartBeatResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendHotlineHeartBeatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendHotlineHeartBeatResponse) String() string {
	return tea.Prettify(s)
}

func (s SendHotlineHeartBeatResponse) GoString() string {
	return s.String()
}

func (s *SendHotlineHeartBeatResponse) SetHeaders(v map[string]*string) *SendHotlineHeartBeatResponse {
	s.Headers = v
	return s
}

func (s *SendHotlineHeartBeatResponse) SetBody(v *SendHotlineHeartBeatResponseBody) *SendHotlineHeartBeatResponse {
	s.Body = v
	return s
}

type TransferCallToSkillGroupRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	SkillGroupId     *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	IsSingleTransfer *bool   `json:"IsSingleTransfer,omitempty" xml:"IsSingleTransfer,omitempty"`
}

func (s TransferCallToSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *TransferCallToSkillGroupRequest) SetClientToken(v string) *TransferCallToSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetInstanceId(v string) *TransferCallToSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetAccountName(v string) *TransferCallToSkillGroupRequest {
	s.AccountName = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetSkillGroupId(v int64) *TransferCallToSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetCallId(v string) *TransferCallToSkillGroupRequest {
	s.CallId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetJobId(v string) *TransferCallToSkillGroupRequest {
	s.JobId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetConnectionId(v string) *TransferCallToSkillGroupRequest {
	s.ConnectionId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetHoldConnectionId(v string) *TransferCallToSkillGroupRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetType(v int32) *TransferCallToSkillGroupRequest {
	s.Type = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetIsSingleTransfer(v bool) *TransferCallToSkillGroupRequest {
	s.IsSingleTransfer = &v
	return s
}

type TransferCallToSkillGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferCallToSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *TransferCallToSkillGroupResponseBody) SetMessage(v string) *TransferCallToSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) SetRequestId(v string) *TransferCallToSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) SetCode(v string) *TransferCallToSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) SetSuccess(v bool) *TransferCallToSkillGroupResponseBody {
	s.Success = &v
	return s
}

type TransferCallToSkillGroupResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferCallToSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferCallToSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *TransferCallToSkillGroupResponse) SetHeaders(v map[string]*string) *TransferCallToSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *TransferCallToSkillGroupResponse) SetBody(v *TransferCallToSkillGroupResponseBody) *TransferCallToSkillGroupResponse {
	s.Body = v
	return s
}

type EditEntityRouteRequest struct {
	EntityId             *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName           *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityBizCode        *string `json:"EntityBizCode,omitempty" xml:"EntityBizCode,omitempty"`
	EntityBizCodeType    *string `json:"EntityBizCodeType,omitempty" xml:"EntityBizCodeType,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
	DepartmentId         *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	GroupId              *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ServiceId            *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ExtInfo              *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UniqueId             *int64  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s EditEntityRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s EditEntityRouteRequest) GoString() string {
	return s.String()
}

func (s *EditEntityRouteRequest) SetEntityId(v string) *EditEntityRouteRequest {
	s.EntityId = &v
	return s
}

func (s *EditEntityRouteRequest) SetEntityName(v string) *EditEntityRouteRequest {
	s.EntityName = &v
	return s
}

func (s *EditEntityRouteRequest) SetEntityBizCode(v string) *EditEntityRouteRequest {
	s.EntityBizCode = &v
	return s
}

func (s *EditEntityRouteRequest) SetEntityBizCodeType(v string) *EditEntityRouteRequest {
	s.EntityBizCodeType = &v
	return s
}

func (s *EditEntityRouteRequest) SetEntityRelationNumber(v string) *EditEntityRouteRequest {
	s.EntityRelationNumber = &v
	return s
}

func (s *EditEntityRouteRequest) SetDepartmentId(v string) *EditEntityRouteRequest {
	s.DepartmentId = &v
	return s
}

func (s *EditEntityRouteRequest) SetGroupId(v int64) *EditEntityRouteRequest {
	s.GroupId = &v
	return s
}

func (s *EditEntityRouteRequest) SetServiceId(v int64) *EditEntityRouteRequest {
	s.ServiceId = &v
	return s
}

func (s *EditEntityRouteRequest) SetExtInfo(v string) *EditEntityRouteRequest {
	s.ExtInfo = &v
	return s
}

func (s *EditEntityRouteRequest) SetInstanceId(v string) *EditEntityRouteRequest {
	s.InstanceId = &v
	return s
}

func (s *EditEntityRouteRequest) SetUniqueId(v int64) *EditEntityRouteRequest {
	s.UniqueId = &v
	return s
}

type EditEntityRouteResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditEntityRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditEntityRouteResponseBody) GoString() string {
	return s.String()
}

func (s *EditEntityRouteResponseBody) SetMessage(v string) *EditEntityRouteResponseBody {
	s.Message = &v
	return s
}

func (s *EditEntityRouteResponseBody) SetRequestId(v string) *EditEntityRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditEntityRouteResponseBody) SetCode(v string) *EditEntityRouteResponseBody {
	s.Code = &v
	return s
}

func (s *EditEntityRouteResponseBody) SetSuccess(v bool) *EditEntityRouteResponseBody {
	s.Success = &v
	return s
}

type EditEntityRouteResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditEntityRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditEntityRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s EditEntityRouteResponse) GoString() string {
	return s.String()
}

func (s *EditEntityRouteResponse) SetHeaders(v map[string]*string) *EditEntityRouteResponse {
	s.Headers = v
	return s
}

func (s *EditEntityRouteResponse) SetBody(v *EditEntityRouteResponseBody) *EditEntityRouteResponse {
	s.Body = v
	return s
}

type GetTagListRequest struct {
	EntityId   *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTagListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTagListRequest) GoString() string {
	return s.String()
}

func (s *GetTagListRequest) SetEntityId(v string) *GetTagListRequest {
	s.EntityId = &v
	return s
}

func (s *GetTagListRequest) SetEntityType(v string) *GetTagListRequest {
	s.EntityType = &v
	return s
}

func (s *GetTagListRequest) SetInstanceId(v string) *GetTagListRequest {
	s.InstanceId = &v
	return s
}

type GetTagListResponseBody struct {
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*GetTagListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTagListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTagListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagListResponseBody) SetMessage(v string) *GetTagListResponseBody {
	s.Message = &v
	return s
}

func (s *GetTagListResponseBody) SetRequestId(v string) *GetTagListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTagListResponseBody) SetData(v []*GetTagListResponseBodyData) *GetTagListResponseBody {
	s.Data = v
	return s
}

func (s *GetTagListResponseBody) SetCode(v string) *GetTagListResponseBody {
	s.Code = &v
	return s
}

func (s *GetTagListResponseBody) SetSuccess(v bool) *GetTagListResponseBody {
	s.Success = &v
	return s
}

type GetTagListResponseBodyData struct {
	ScenarioCode *string                                `json:"ScenarioCode,omitempty" xml:"ScenarioCode,omitempty"`
	TagGroupCode *string                                `json:"TagGroupCode,omitempty" xml:"TagGroupCode,omitempty"`
	TagValues    []*GetTagListResponseBodyDataTagValues `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
	TagGroupName *string                                `json:"TagGroupName,omitempty" xml:"TagGroupName,omitempty"`
}

func (s GetTagListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTagListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTagListResponseBodyData) SetScenarioCode(v string) *GetTagListResponseBodyData {
	s.ScenarioCode = &v
	return s
}

func (s *GetTagListResponseBodyData) SetTagGroupCode(v string) *GetTagListResponseBodyData {
	s.TagGroupCode = &v
	return s
}

func (s *GetTagListResponseBodyData) SetTagValues(v []*GetTagListResponseBodyDataTagValues) *GetTagListResponseBodyData {
	s.TagValues = v
	return s
}

func (s *GetTagListResponseBodyData) SetTagGroupName(v string) *GetTagListResponseBodyData {
	s.TagGroupName = &v
	return s
}

type GetTagListResponseBodyDataTagValues struct {
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TagName              *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagGroupCode         *string `json:"TagGroupCode,omitempty" xml:"TagGroupCode,omitempty"`
	TagCode              *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	TagGroupName         *string `json:"TagGroupName,omitempty" xml:"TagGroupName,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
}

func (s GetTagListResponseBodyDataTagValues) String() string {
	return tea.Prettify(s)
}

func (s GetTagListResponseBodyDataTagValues) GoString() string {
	return s.String()
}

func (s *GetTagListResponseBodyDataTagValues) SetStatus(v string) *GetTagListResponseBodyDataTagValues {
	s.Status = &v
	return s
}

func (s *GetTagListResponseBodyDataTagValues) SetDescription(v string) *GetTagListResponseBodyDataTagValues {
	s.Description = &v
	return s
}

func (s *GetTagListResponseBodyDataTagValues) SetTagName(v string) *GetTagListResponseBodyDataTagValues {
	s.TagName = &v
	return s
}

func (s *GetTagListResponseBodyDataTagValues) SetTagGroupCode(v string) *GetTagListResponseBodyDataTagValues {
	s.TagGroupCode = &v
	return s
}

func (s *GetTagListResponseBodyDataTagValues) SetTagCode(v string) *GetTagListResponseBodyDataTagValues {
	s.TagCode = &v
	return s
}

func (s *GetTagListResponseBodyDataTagValues) SetTagGroupName(v string) *GetTagListResponseBodyDataTagValues {
	s.TagGroupName = &v
	return s
}

func (s *GetTagListResponseBodyDataTagValues) SetEntityRelationNumber(v string) *GetTagListResponseBodyDataTagValues {
	s.EntityRelationNumber = &v
	return s
}

type GetTagListResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTagListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTagListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTagListResponse) GoString() string {
	return s.String()
}

func (s *GetTagListResponse) SetHeaders(v map[string]*string) *GetTagListResponse {
	s.Headers = v
	return s
}

func (s *GetTagListResponse) SetBody(v *GetTagListResponseBody) *GetTagListResponse {
	s.Body = v
	return s
}

type UpdateTicketRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TicketId    *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	OperatorId  *int64  `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	FormData    *string `json:"FormData,omitempty" xml:"FormData,omitempty"`
}

func (s UpdateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketRequest) GoString() string {
	return s.String()
}

func (s *UpdateTicketRequest) SetClientToken(v string) *UpdateTicketRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTicketRequest) SetInstanceId(v string) *UpdateTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTicketRequest) SetTicketId(v int64) *UpdateTicketRequest {
	s.TicketId = &v
	return s
}

func (s *UpdateTicketRequest) SetOperatorId(v int64) *UpdateTicketRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateTicketRequest) SetFormData(v string) *UpdateTicketRequest {
	s.FormData = &v
	return s
}

type UpdateTicketResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s UpdateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTicketResponseBody) SetMessage(v string) *UpdateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTicketResponseBody) SetRequestId(v string) *UpdateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTicketResponseBody) SetCode(v string) *UpdateTicketResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTicketResponseBody) SetSuccess(v bool) *UpdateTicketResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTicketResponseBody) SetHttpStatusCode(v int64) *UpdateTicketResponseBody {
	s.HttpStatusCode = &v
	return s
}

type UpdateTicketResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketResponse) GoString() string {
	return s.String()
}

func (s *UpdateTicketResponse) SetHeaders(v map[string]*string) *UpdateTicketResponse {
	s.Headers = v
	return s
}

func (s *UpdateTicketResponse) SetBody(v *UpdateTicketResponseBody) *UpdateTicketResponse {
	s.Body = v
	return s
}

type ListOutboundPhoneNumberRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s ListOutboundPhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberRequest) SetClientToken(v string) *ListOutboundPhoneNumberRequest {
	s.ClientToken = &v
	return s
}

func (s *ListOutboundPhoneNumberRequest) SetInstanceId(v string) *ListOutboundPhoneNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOutboundPhoneNumberRequest) SetAccountName(v string) *ListOutboundPhoneNumberRequest {
	s.AccountName = &v
	return s
}

type ListOutboundPhoneNumberResponseBody struct {
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s ListOutboundPhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberResponseBody) SetMessage(v string) *ListOutboundPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetRequestId(v string) *ListOutboundPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetData(v []*string) *ListOutboundPhoneNumberResponseBody {
	s.Data = v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetCode(v string) *ListOutboundPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetSuccess(v bool) *ListOutboundPhoneNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetHttpStatusCode(v int64) *ListOutboundPhoneNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

type ListOutboundPhoneNumberResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOutboundPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOutboundPhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberResponse) SetHeaders(v map[string]*string) *ListOutboundPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *ListOutboundPhoneNumberResponse) SetBody(v *ListOutboundPhoneNumberResponseBody) *ListOutboundPhoneNumberResponse {
	s.Body = v
	return s
}

type RemoveSkillGroupRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s RemoveSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupRequest) SetInstanceId(v string) *RemoveSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveSkillGroupRequest) SetSkillGroupId(v string) *RemoveSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *RemoveSkillGroupRequest) SetClientToken(v string) *RemoveSkillGroupRequest {
	s.ClientToken = &v
	return s
}

type RemoveSkillGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupResponseBody) SetMessage(v string) *RemoveSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) SetRequestId(v string) *RemoveSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) SetCode(v string) *RemoveSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) SetSuccess(v bool) *RemoveSkillGroupResponseBody {
	s.Success = &v
	return s
}

type RemoveSkillGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupResponse) SetHeaders(v map[string]*string) *RemoveSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveSkillGroupResponse) SetBody(v *RemoveSkillGroupResponseBody) *RemoveSkillGroupResponse {
	s.Body = v
	return s
}

type ListAgentBySkillGroupIdRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListAgentBySkillGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAgentBySkillGroupIdRequest) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdRequest) SetClientToken(v string) *ListAgentBySkillGroupIdRequest {
	s.ClientToken = &v
	return s
}

func (s *ListAgentBySkillGroupIdRequest) SetInstanceId(v string) *ListAgentBySkillGroupIdRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAgentBySkillGroupIdRequest) SetSkillGroupId(v int64) *ListAgentBySkillGroupIdRequest {
	s.SkillGroupId = &v
	return s
}

type ListAgentBySkillGroupIdResponseBody struct {
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*ListAgentBySkillGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAgentBySkillGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAgentBySkillGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdResponseBody) SetMessage(v string) *ListAgentBySkillGroupIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetRequestId(v string) *ListAgentBySkillGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetData(v []*ListAgentBySkillGroupIdResponseBodyData) *ListAgentBySkillGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetCode(v string) *ListAgentBySkillGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetSuccess(v bool) *ListAgentBySkillGroupIdResponseBody {
	s.Success = &v
	return s
}

type ListAgentBySkillGroupIdResponseBodyData struct {
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	AgentId     *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	TenantId    *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListAgentBySkillGroupIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAgentBySkillGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetStatus(v int32) *ListAgentBySkillGroupIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetDisplayName(v string) *ListAgentBySkillGroupIdResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetAgentId(v int64) *ListAgentBySkillGroupIdResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetAccountName(v string) *ListAgentBySkillGroupIdResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetTenantId(v int64) *ListAgentBySkillGroupIdResponseBodyData {
	s.TenantId = &v
	return s
}

type ListAgentBySkillGroupIdResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAgentBySkillGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAgentBySkillGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAgentBySkillGroupIdResponse) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdResponse) SetHeaders(v map[string]*string) *ListAgentBySkillGroupIdResponse {
	s.Headers = v
	return s
}

func (s *ListAgentBySkillGroupIdResponse) SetBody(v *ListAgentBySkillGroupIdResponseBody) *ListAgentBySkillGroupIdResponse {
	s.Body = v
	return s
}

type QueryHotlineSessionRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Acid           *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	CallType       *int32  `json:"CallType,omitempty" xml:"CallType,omitempty"`
	CalledNumber   *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber  *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	GroupId        *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	MemberId       *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName     *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	QueryEndTime   *int64  `json:"QueryEndTime,omitempty" xml:"QueryEndTime,omitempty"`
	QueryStartTime *int64  `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServicerName   *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	ServicerId     *string `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	Params         *string `json:"Params,omitempty" xml:"Params,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo         *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	CallResult     *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
}

func (s QueryHotlineSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineSessionRequest) GoString() string {
	return s.String()
}

func (s *QueryHotlineSessionRequest) SetInstanceId(v string) *QueryHotlineSessionRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetAcid(v string) *QueryHotlineSessionRequest {
	s.Acid = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetCallType(v int32) *QueryHotlineSessionRequest {
	s.CallType = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetCalledNumber(v string) *QueryHotlineSessionRequest {
	s.CalledNumber = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetCallingNumber(v string) *QueryHotlineSessionRequest {
	s.CallingNumber = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetGroupId(v int64) *QueryHotlineSessionRequest {
	s.GroupId = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetGroupName(v string) *QueryHotlineSessionRequest {
	s.GroupName = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetMemberId(v string) *QueryHotlineSessionRequest {
	s.MemberId = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetMemberName(v string) *QueryHotlineSessionRequest {
	s.MemberName = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetQueryEndTime(v int64) *QueryHotlineSessionRequest {
	s.QueryEndTime = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetQueryStartTime(v int64) *QueryHotlineSessionRequest {
	s.QueryStartTime = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetRequestId(v string) *QueryHotlineSessionRequest {
	s.RequestId = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetServicerName(v string) *QueryHotlineSessionRequest {
	s.ServicerName = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetServicerId(v string) *QueryHotlineSessionRequest {
	s.ServicerId = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetParams(v string) *QueryHotlineSessionRequest {
	s.Params = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetPageSize(v int32) *QueryHotlineSessionRequest {
	s.PageSize = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetPageNo(v int32) *QueryHotlineSessionRequest {
	s.PageNo = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetCallResult(v string) *QueryHotlineSessionRequest {
	s.CallResult = &v
	return s
}

type QueryHotlineSessionResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *QueryHotlineSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryHotlineSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineSessionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotlineSessionResponseBody) SetMessage(v string) *QueryHotlineSessionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHotlineSessionResponseBody) SetRequestId(v string) *QueryHotlineSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHotlineSessionResponseBody) SetData(v *QueryHotlineSessionResponseBodyData) *QueryHotlineSessionResponseBody {
	s.Data = v
	return s
}

func (s *QueryHotlineSessionResponseBody) SetCode(v string) *QueryHotlineSessionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryHotlineSessionResponseBody) SetSuccess(v bool) *QueryHotlineSessionResponseBody {
	s.Success = &v
	return s
}

type QueryHotlineSessionResponseBodyData struct {
	CallDetailRecord []*QueryHotlineSessionResponseBodyDataCallDetailRecord `json:"CallDetailRecord,omitempty" xml:"CallDetailRecord,omitempty" type:"Repeated"`
	PageSize         *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber       *int32                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount       *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryHotlineSessionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryHotlineSessionResponseBodyData) SetCallDetailRecord(v []*QueryHotlineSessionResponseBodyDataCallDetailRecord) *QueryHotlineSessionResponseBodyData {
	s.CallDetailRecord = v
	return s
}

func (s *QueryHotlineSessionResponseBodyData) SetPageSize(v int32) *QueryHotlineSessionResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyData) SetPageNumber(v int32) *QueryHotlineSessionResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyData) SetTotalCount(v int32) *QueryHotlineSessionResponseBodyData {
	s.TotalCount = &v
	return s
}

type QueryHotlineSessionResponseBodyDataCallDetailRecord struct {
	CallResult          *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	ServicerName        *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
	OutQueueTime        *string `json:"OutQueueTime,omitempty" xml:"OutQueueTime,omitempty"`
	CallContinueTime    *int32  `json:"CallContinueTime,omitempty" xml:"CallContinueTime,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PickUpTime          *string `json:"PickUpTime,omitempty" xml:"PickUpTime,omitempty"`
	RingContinueTime    *int32  `json:"RingContinueTime,omitempty" xml:"RingContinueTime,omitempty"`
	CalledNumber        *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	ServicerId          *string `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	HangUpTime          *string `json:"HangUpTime,omitempty" xml:"HangUpTime,omitempty"`
	EvaluationLevel     *int32  `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty"`
	HangUpRole          *string `json:"HangUpRole,omitempty" xml:"HangUpRole,omitempty"`
	MemberName          *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	EvaluationScore     *int32  `json:"EvaluationScore,omitempty" xml:"EvaluationScore,omitempty"`
	Acid                *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	RingStartTime       *string `json:"RingStartTime,omitempty" xml:"RingStartTime,omitempty"`
	CallType            *int32  `json:"CallType,omitempty" xml:"CallType,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupId             *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RingEndTime         *string `json:"RingEndTime,omitempty" xml:"RingEndTime,omitempty"`
	InQueueTime         *string `json:"InQueueTime,omitempty" xml:"InQueueTime,omitempty"`
	CallingNumber       *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	MemberId            *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	QueueUpContinueTime *int32  `json:"QueueUpContinueTime,omitempty" xml:"QueueUpContinueTime,omitempty"`
}

func (s QueryHotlineSessionResponseBodyDataCallDetailRecord) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineSessionResponseBodyDataCallDetailRecord) GoString() string {
	return s.String()
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetCallResult(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.CallResult = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetServicerName(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.ServicerName = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetOutQueueTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.OutQueueTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetCallContinueTime(v int32) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.CallContinueTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetCreateTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.CreateTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetPickUpTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.PickUpTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetRingContinueTime(v int32) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.RingContinueTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetCalledNumber(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.CalledNumber = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetServicerId(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.ServicerId = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetHangUpTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.HangUpTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetEvaluationLevel(v int32) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.EvaluationLevel = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetHangUpRole(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.HangUpRole = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetMemberName(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.MemberName = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetEvaluationScore(v int32) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.EvaluationScore = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetAcid(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.Acid = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetRingStartTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.RingStartTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetCallType(v int32) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.CallType = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetGroupName(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.GroupName = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetGroupId(v int64) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.GroupId = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetRingEndTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.RingEndTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetInQueueTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.InQueueTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetCallingNumber(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.CallingNumber = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetMemberId(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.MemberId = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetQueueUpContinueTime(v int32) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.QueueUpContinueTime = &v
	return s
}

type QueryHotlineSessionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryHotlineSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHotlineSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineSessionResponse) GoString() string {
	return s.String()
}

func (s *QueryHotlineSessionResponse) SetHeaders(v map[string]*string) *QueryHotlineSessionResponse {
	s.Headers = v
	return s
}

func (s *QueryHotlineSessionResponse) SetBody(v *QueryHotlineSessionResponseBody) *QueryHotlineSessionResponse {
	s.Body = v
	return s
}

type StartChatWorkRequest struct {
	// instanceId
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// accountName
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s StartChatWorkRequest) String() string {
	return tea.Prettify(s)
}

func (s StartChatWorkRequest) GoString() string {
	return s.String()
}

func (s *StartChatWorkRequest) SetInstanceId(v string) *StartChatWorkRequest {
	s.InstanceId = &v
	return s
}

func (s *StartChatWorkRequest) SetAccountName(v string) *StartChatWorkRequest {
	s.AccountName = &v
	return s
}

type StartChatWorkResponseBody struct {
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartChatWorkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartChatWorkResponseBody) GoString() string {
	return s.String()
}

func (s *StartChatWorkResponseBody) SetMessage(v string) *StartChatWorkResponseBody {
	s.Message = &v
	return s
}

func (s *StartChatWorkResponseBody) SetRequestId(v string) *StartChatWorkResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartChatWorkResponseBody) SetHttpStatusCode(v int32) *StartChatWorkResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartChatWorkResponseBody) SetData(v string) *StartChatWorkResponseBody {
	s.Data = &v
	return s
}

func (s *StartChatWorkResponseBody) SetCode(v string) *StartChatWorkResponseBody {
	s.Code = &v
	return s
}

func (s *StartChatWorkResponseBody) SetSuccess(v bool) *StartChatWorkResponseBody {
	s.Success = &v
	return s
}

type StartChatWorkResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartChatWorkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartChatWorkResponse) String() string {
	return tea.Prettify(s)
}

func (s StartChatWorkResponse) GoString() string {
	return s.String()
}

func (s *StartChatWorkResponse) SetHeaders(v map[string]*string) *StartChatWorkResponse {
	s.Headers = v
	return s
}

func (s *StartChatWorkResponse) SetBody(v *StartChatWorkResponseBody) *StartChatWorkResponse {
	s.Body = v
	return s
}

type HangupThirdCallRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
}

func (s HangupThirdCallRequest) String() string {
	return tea.Prettify(s)
}

func (s HangupThirdCallRequest) GoString() string {
	return s.String()
}

func (s *HangupThirdCallRequest) SetClientToken(v string) *HangupThirdCallRequest {
	s.ClientToken = &v
	return s
}

func (s *HangupThirdCallRequest) SetInstanceId(v string) *HangupThirdCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HangupThirdCallRequest) SetAccountName(v string) *HangupThirdCallRequest {
	s.AccountName = &v
	return s
}

func (s *HangupThirdCallRequest) SetCallId(v string) *HangupThirdCallRequest {
	s.CallId = &v
	return s
}

func (s *HangupThirdCallRequest) SetJobId(v string) *HangupThirdCallRequest {
	s.JobId = &v
	return s
}

func (s *HangupThirdCallRequest) SetConnectionId(v string) *HangupThirdCallRequest {
	s.ConnectionId = &v
	return s
}

type HangupThirdCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HangupThirdCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HangupThirdCallResponseBody) GoString() string {
	return s.String()
}

func (s *HangupThirdCallResponseBody) SetMessage(v string) *HangupThirdCallResponseBody {
	s.Message = &v
	return s
}

func (s *HangupThirdCallResponseBody) SetRequestId(v string) *HangupThirdCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HangupThirdCallResponseBody) SetCode(v string) *HangupThirdCallResponseBody {
	s.Code = &v
	return s
}

func (s *HangupThirdCallResponseBody) SetSuccess(v bool) *HangupThirdCallResponseBody {
	s.Success = &v
	return s
}

type HangupThirdCallResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HangupThirdCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HangupThirdCallResponse) String() string {
	return tea.Prettify(s)
}

func (s HangupThirdCallResponse) GoString() string {
	return s.String()
}

func (s *HangupThirdCallResponse) SetHeaders(v map[string]*string) *HangupThirdCallResponse {
	s.Headers = v
	return s
}

func (s *HangupThirdCallResponse) SetBody(v *HangupThirdCallResponseBody) *HangupThirdCallResponse {
	s.Body = v
	return s
}

type StartHotlineServiceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s StartHotlineServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineServiceRequest) GoString() string {
	return s.String()
}

func (s *StartHotlineServiceRequest) SetClientToken(v string) *StartHotlineServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartHotlineServiceRequest) SetInstanceId(v string) *StartHotlineServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartHotlineServiceRequest) SetAccountName(v string) *StartHotlineServiceRequest {
	s.AccountName = &v
	return s
}

type StartHotlineServiceResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s StartHotlineServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartHotlineServiceResponseBody) SetMessage(v string) *StartHotlineServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetRequestId(v string) *StartHotlineServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetData(v string) *StartHotlineServiceResponseBody {
	s.Data = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetCode(v string) *StartHotlineServiceResponseBody {
	s.Code = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetSuccess(v bool) *StartHotlineServiceResponseBody {
	s.Success = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetHttpStatusCode(v int64) *StartHotlineServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

type StartHotlineServiceResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartHotlineServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartHotlineServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineServiceResponse) GoString() string {
	return s.String()
}

func (s *StartHotlineServiceResponse) SetHeaders(v map[string]*string) *StartHotlineServiceResponse {
	s.Headers = v
	return s
}

func (s *StartHotlineServiceResponse) SetBody(v *StartHotlineServiceResponseBody) *StartHotlineServiceResponse {
	s.Body = v
	return s
}

type StartCallV2Request struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Caller      *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Callee      *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	JsonMsg     *string `json:"JsonMsg,omitempty" xml:"JsonMsg,omitempty"`
}

func (s StartCallV2Request) String() string {
	return tea.Prettify(s)
}

func (s StartCallV2Request) GoString() string {
	return s.String()
}

func (s *StartCallV2Request) SetClientToken(v string) *StartCallV2Request {
	s.ClientToken = &v
	return s
}

func (s *StartCallV2Request) SetInstanceId(v string) *StartCallV2Request {
	s.InstanceId = &v
	return s
}

func (s *StartCallV2Request) SetAccountName(v string) *StartCallV2Request {
	s.AccountName = &v
	return s
}

func (s *StartCallV2Request) SetCaller(v string) *StartCallV2Request {
	s.Caller = &v
	return s
}

func (s *StartCallV2Request) SetCallee(v string) *StartCallV2Request {
	s.Callee = &v
	return s
}

func (s *StartCallV2Request) SetJsonMsg(v string) *StartCallV2Request {
	s.JsonMsg = &v
	return s
}

type StartCallV2ResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartCallV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCallV2ResponseBody) GoString() string {
	return s.String()
}

func (s *StartCallV2ResponseBody) SetMessage(v string) *StartCallV2ResponseBody {
	s.Message = &v
	return s
}

func (s *StartCallV2ResponseBody) SetRequestId(v string) *StartCallV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCallV2ResponseBody) SetCode(v string) *StartCallV2ResponseBody {
	s.Code = &v
	return s
}

func (s *StartCallV2ResponseBody) SetSuccess(v bool) *StartCallV2ResponseBody {
	s.Success = &v
	return s
}

type StartCallV2Response struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartCallV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCallV2Response) String() string {
	return tea.Prettify(s)
}

func (s StartCallV2Response) GoString() string {
	return s.String()
}

func (s *StartCallV2Response) SetHeaders(v map[string]*string) *StartCallV2Response {
	s.Headers = v
	return s
}

func (s *StartCallV2Response) SetBody(v *StartCallV2ResponseBody) *StartCallV2Response {
	s.Body = v
	return s
}

type EnableRoleRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleId      *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s EnableRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableRoleRequest) GoString() string {
	return s.String()
}

func (s *EnableRoleRequest) SetClientToken(v string) *EnableRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableRoleRequest) SetInstanceId(v string) *EnableRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableRoleRequest) SetRoleId(v int64) *EnableRoleRequest {
	s.RoleId = &v
	return s
}

type EnableRoleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableRoleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableRoleResponseBody) SetMessage(v string) *EnableRoleResponseBody {
	s.Message = &v
	return s
}

func (s *EnableRoleResponseBody) SetRequestId(v string) *EnableRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableRoleResponseBody) SetCode(v string) *EnableRoleResponseBody {
	s.Code = &v
	return s
}

func (s *EnableRoleResponseBody) SetSuccess(v bool) *EnableRoleResponseBody {
	s.Success = &v
	return s
}

type EnableRoleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableRoleResponse) GoString() string {
	return s.String()
}

func (s *EnableRoleResponse) SetHeaders(v map[string]*string) *EnableRoleResponse {
	s.Headers = v
	return s
}

func (s *EnableRoleResponse) SetBody(v *EnableRoleResponseBody) *EnableRoleResponse {
	s.Body = v
	return s
}

type GetAgentRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) SetClientToken(v string) *GetAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAgentRequest) SetInstanceId(v string) *GetAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentRequest) SetAccountName(v string) *GetAgentRequest {
	s.AccountName = &v
	return s
}

type GetAgentResponseBody struct {
	RequestId      *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code           *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) SetRequestId(v string) *GetAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentResponseBody) SetData(v *GetAgentResponseBodyData) *GetAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentResponseBody) SetCode(v string) *GetAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentResponseBody) SetSuccess(v bool) *GetAgentResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentResponseBody) SetHttpStatusCode(v int64) *GetAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAgentResponseBody) SetMessage(v string) *GetAgentResponseBody {
	s.Message = &v
	return s
}

type GetAgentResponseBodyData struct {
	Status      *int32                               `json:"Status,omitempty" xml:"Status,omitempty"`
	DisplayName *string                              `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	AgentId     *int64                               `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	GroupList   []*GetAgentResponseBodyDataGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	AccountName *string                              `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	TenantId    *int64                               `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetAgentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyData) SetStatus(v int32) *GetAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAgentResponseBodyData) SetDisplayName(v string) *GetAgentResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetAgentResponseBodyData) SetAgentId(v int64) *GetAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetAgentResponseBodyData) SetGroupList(v []*GetAgentResponseBodyDataGroupList) *GetAgentResponseBodyData {
	s.GroupList = v
	return s
}

func (s *GetAgentResponseBodyData) SetAccountName(v string) *GetAgentResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *GetAgentResponseBodyData) SetTenantId(v int64) *GetAgentResponseBodyData {
	s.TenantId = &v
	return s
}

type GetAgentResponseBodyDataGroupList struct {
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ChannelType  *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	SkillGroupId *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAgentResponseBodyDataGroupList) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBodyDataGroupList) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyDataGroupList) SetDisplayName(v string) *GetAgentResponseBodyDataGroupList {
	s.DisplayName = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetDescription(v string) *GetAgentResponseBodyDataGroupList {
	s.Description = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetChannelType(v int32) *GetAgentResponseBodyDataGroupList {
	s.ChannelType = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetSkillGroupId(v int64) *GetAgentResponseBodyDataGroupList {
	s.SkillGroupId = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetName(v string) *GetAgentResponseBodyDataGroupList {
	s.Name = &v
	return s
}

type GetAgentResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponse) GoString() string {
	return s.String()
}

func (s *GetAgentResponse) SetHeaders(v map[string]*string) *GetAgentResponse {
	s.Headers = v
	return s
}

func (s *GetAgentResponse) SetBody(v *GetAgentResponseBody) *GetAgentResponse {
	s.Body = v
	return s
}

type GetCMSIdByForeignIdRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Nick       *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	ForeignId  *string `json:"ForeignId,omitempty" xml:"ForeignId,omitempty"`
	SourceId   *int64  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s GetCMSIdByForeignIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCMSIdByForeignIdRequest) GoString() string {
	return s.String()
}

func (s *GetCMSIdByForeignIdRequest) SetInstanceId(v string) *GetCMSIdByForeignIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCMSIdByForeignIdRequest) SetNick(v string) *GetCMSIdByForeignIdRequest {
	s.Nick = &v
	return s
}

func (s *GetCMSIdByForeignIdRequest) SetForeignId(v string) *GetCMSIdByForeignIdRequest {
	s.ForeignId = &v
	return s
}

func (s *GetCMSIdByForeignIdRequest) SetSourceId(v int64) *GetCMSIdByForeignIdRequest {
	s.SourceId = &v
	return s
}

type GetCMSIdByForeignIdResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetCMSIdByForeignIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCMSIdByForeignIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCMSIdByForeignIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetCMSIdByForeignIdResponseBody) SetMessage(v string) *GetCMSIdByForeignIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBody) SetRequestId(v string) *GetCMSIdByForeignIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBody) SetData(v *GetCMSIdByForeignIdResponseBodyData) *GetCMSIdByForeignIdResponseBody {
	s.Data = v
	return s
}

func (s *GetCMSIdByForeignIdResponseBody) SetCode(v string) *GetCMSIdByForeignIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBody) SetSuccess(v bool) *GetCMSIdByForeignIdResponseBody {
	s.Success = &v
	return s
}

type GetCMSIdByForeignIdResponseBodyData struct {
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	CustomerTypeId *int32  `json:"CustomerTypeId,omitempty" xml:"CustomerTypeId,omitempty"`
	Avatar         *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Gender         *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	ForeignId      *string `json:"ForeignId,omitempty" xml:"ForeignId,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Nick           *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	Anonymity      *bool   `json:"Anonymity,omitempty" xml:"Anonymity,omitempty"`
	Phone          *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s GetCMSIdByForeignIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCMSIdByForeignIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetStatus(v int32) *GetCMSIdByForeignIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetCustomerTypeId(v int32) *GetCMSIdByForeignIdResponseBodyData {
	s.CustomerTypeId = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetAvatar(v string) *GetCMSIdByForeignIdResponseBodyData {
	s.Avatar = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetGender(v string) *GetCMSIdByForeignIdResponseBodyData {
	s.Gender = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetForeignId(v string) *GetCMSIdByForeignIdResponseBodyData {
	s.ForeignId = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetUserId(v string) *GetCMSIdByForeignIdResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetNick(v string) *GetCMSIdByForeignIdResponseBodyData {
	s.Nick = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetAnonymity(v bool) *GetCMSIdByForeignIdResponseBodyData {
	s.Anonymity = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetPhone(v string) *GetCMSIdByForeignIdResponseBodyData {
	s.Phone = &v
	return s
}

type GetCMSIdByForeignIdResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCMSIdByForeignIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCMSIdByForeignIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCMSIdByForeignIdResponse) GoString() string {
	return s.String()
}

func (s *GetCMSIdByForeignIdResponse) SetHeaders(v map[string]*string) *GetCMSIdByForeignIdResponse {
	s.Headers = v
	return s
}

func (s *GetCMSIdByForeignIdResponse) SetBody(v *GetCMSIdByForeignIdResponseBody) *GetCMSIdByForeignIdResponse {
	s.Body = v
	return s
}

type TransferToThirdCallRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
}

func (s TransferToThirdCallRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferToThirdCallRequest) GoString() string {
	return s.String()
}

func (s *TransferToThirdCallRequest) SetClientToken(v string) *TransferToThirdCallRequest {
	s.ClientToken = &v
	return s
}

func (s *TransferToThirdCallRequest) SetInstanceId(v string) *TransferToThirdCallRequest {
	s.InstanceId = &v
	return s
}

func (s *TransferToThirdCallRequest) SetAccountName(v string) *TransferToThirdCallRequest {
	s.AccountName = &v
	return s
}

func (s *TransferToThirdCallRequest) SetCallId(v string) *TransferToThirdCallRequest {
	s.CallId = &v
	return s
}

func (s *TransferToThirdCallRequest) SetJobId(v string) *TransferToThirdCallRequest {
	s.JobId = &v
	return s
}

func (s *TransferToThirdCallRequest) SetConnectionId(v string) *TransferToThirdCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *TransferToThirdCallRequest) SetHoldConnectionId(v string) *TransferToThirdCallRequest {
	s.HoldConnectionId = &v
	return s
}

type TransferToThirdCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferToThirdCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferToThirdCallResponseBody) GoString() string {
	return s.String()
}

func (s *TransferToThirdCallResponseBody) SetMessage(v string) *TransferToThirdCallResponseBody {
	s.Message = &v
	return s
}

func (s *TransferToThirdCallResponseBody) SetRequestId(v string) *TransferToThirdCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferToThirdCallResponseBody) SetCode(v string) *TransferToThirdCallResponseBody {
	s.Code = &v
	return s
}

func (s *TransferToThirdCallResponseBody) SetSuccess(v bool) *TransferToThirdCallResponseBody {
	s.Success = &v
	return s
}

type TransferToThirdCallResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferToThirdCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferToThirdCallResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferToThirdCallResponse) GoString() string {
	return s.String()
}

func (s *TransferToThirdCallResponse) SetHeaders(v map[string]*string) *TransferToThirdCallResponse {
	s.Headers = v
	return s
}

func (s *TransferToThirdCallResponse) SetBody(v *TransferToThirdCallResponseBody) *TransferToThirdCallResponse {
	s.Body = v
	return s
}

type UpdateAgentRequest struct {
	ClientToken      *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName      *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName      *string  `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	SkillGroupId     []*int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty" type:"Repeated"`
	SkillGroupIdList []*int64 `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty" type:"Repeated"`
}

func (s UpdateAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequest) SetClientToken(v string) *UpdateAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAgentRequest) SetInstanceId(v string) *UpdateAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAgentRequest) SetAccountName(v string) *UpdateAgentRequest {
	s.AccountName = &v
	return s
}

func (s *UpdateAgentRequest) SetDisplayName(v string) *UpdateAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateAgentRequest) SetSkillGroupId(v []*int64) *UpdateAgentRequest {
	s.SkillGroupId = v
	return s
}

func (s *UpdateAgentRequest) SetSkillGroupIdList(v []*int64) *UpdateAgentRequest {
	s.SkillGroupIdList = v
	return s
}

type UpdateAgentResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s UpdateAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBody) SetMessage(v string) *UpdateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentResponseBody) SetRequestId(v string) *UpdateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentResponseBody) SetCode(v string) *UpdateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentResponseBody) SetSuccess(v bool) *UpdateAgentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgentResponseBody) SetHttpStatusCode(v int64) *UpdateAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

type UpdateAgentResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponse) SetHeaders(v map[string]*string) *UpdateAgentResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentResponse) SetBody(v *UpdateAgentResponseBody) *UpdateAgentResponse {
	s.Body = v
	return s
}

type ChangeChatAgentStatusRequest struct {
	// clientToken
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 示例id
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 账户名称
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// 修改到的状态类型
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s ChangeChatAgentStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeChatAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *ChangeChatAgentStatusRequest) SetClientToken(v string) *ChangeChatAgentStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) SetInstanceId(v string) *ChangeChatAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) SetAccountName(v string) *ChangeChatAgentStatusRequest {
	s.AccountName = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) SetMethod(v string) *ChangeChatAgentStatusRequest {
	s.Method = &v
	return s
}

type ChangeChatAgentStatusResponseBody struct {
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeChatAgentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeChatAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeChatAgentStatusResponseBody) SetMessage(v string) *ChangeChatAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetRequestId(v string) *ChangeChatAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetHttpStatusCode(v int32) *ChangeChatAgentStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetData(v string) *ChangeChatAgentStatusResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetCode(v string) *ChangeChatAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetSuccess(v bool) *ChangeChatAgentStatusResponseBody {
	s.Success = &v
	return s
}

type ChangeChatAgentStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeChatAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeChatAgentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeChatAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *ChangeChatAgentStatusResponse) SetHeaders(v map[string]*string) *ChangeChatAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *ChangeChatAgentStatusResponse) SetBody(v *ChangeChatAgentStatusResponseBody) *ChangeChatAgentStatusResponse {
	s.Body = v
	return s
}

type GenerateWebSocketSignRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GenerateWebSocketSignRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebSocketSignRequest) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignRequest) SetClientToken(v string) *GenerateWebSocketSignRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateWebSocketSignRequest) SetInstanceId(v string) *GenerateWebSocketSignRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateWebSocketSignRequest) SetAccountName(v string) *GenerateWebSocketSignRequest {
	s.AccountName = &v
	return s
}

type GenerateWebSocketSignResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GenerateWebSocketSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebSocketSignResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignResponseBody) SetMessage(v string) *GenerateWebSocketSignResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetRequestId(v string) *GenerateWebSocketSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetData(v string) *GenerateWebSocketSignResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetCode(v string) *GenerateWebSocketSignResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetSuccess(v bool) *GenerateWebSocketSignResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetHttpStatusCode(v int64) *GenerateWebSocketSignResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GenerateWebSocketSignResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateWebSocketSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateWebSocketSignResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebSocketSignResponse) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignResponse) SetHeaders(v map[string]*string) *GenerateWebSocketSignResponse {
	s.Headers = v
	return s
}

func (s *GenerateWebSocketSignResponse) SetBody(v *GenerateWebSocketSignResponseBody) *GenerateWebSocketSignResponse {
	s.Body = v
	return s
}

type UpdateRingStatusRequest struct {
	UniqueBizId   *string `json:"UniqueBizId,omitempty" xml:"UniqueBizId,omitempty"`
	CallOutStatus *string `json:"CallOutStatus,omitempty" xml:"CallOutStatus,omitempty"`
	Extra         *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateRingStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRingStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateRingStatusRequest) SetUniqueBizId(v string) *UpdateRingStatusRequest {
	s.UniqueBizId = &v
	return s
}

func (s *UpdateRingStatusRequest) SetCallOutStatus(v string) *UpdateRingStatusRequest {
	s.CallOutStatus = &v
	return s
}

func (s *UpdateRingStatusRequest) SetExtra(v string) *UpdateRingStatusRequest {
	s.Extra = &v
	return s
}

func (s *UpdateRingStatusRequest) SetInstanceId(v string) *UpdateRingStatusRequest {
	s.InstanceId = &v
	return s
}

type UpdateRingStatusResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s UpdateRingStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRingStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRingStatusResponseBody) SetMessage(v string) *UpdateRingStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRingStatusResponseBody) SetRequestId(v string) *UpdateRingStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRingStatusResponseBody) SetData(v string) *UpdateRingStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRingStatusResponseBody) SetCode(v string) *UpdateRingStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRingStatusResponseBody) SetSuccess(v bool) *UpdateRingStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRingStatusResponseBody) SetHttpStatusCode(v int64) *UpdateRingStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

type UpdateRingStatusResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRingStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRingStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRingStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateRingStatusResponse) SetHeaders(v map[string]*string) *UpdateRingStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateRingStatusResponse) SetBody(v *UpdateRingStatusResponseBody) *UpdateRingStatusResponse {
	s.Body = v
	return s
}

type CreateAgentRequest struct {
	// js sdk中自动生成的鉴权token
	ClientToken      *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName      *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DisplayName      *string  `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	SkillGroupId     []*int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty" type:"Repeated"`
	SkillGroupIdList []*int64 `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty" type:"Repeated"`
}

func (s CreateAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) SetClientToken(v string) *CreateAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgentRequest) SetInstanceId(v string) *CreateAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAgentRequest) SetAccountName(v string) *CreateAgentRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAgentRequest) SetDisplayName(v string) *CreateAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateAgentRequest) SetSkillGroupId(v []*int64) *CreateAgentRequest {
	s.SkillGroupId = v
	return s
}

func (s *CreateAgentRequest) SetSkillGroupIdList(v []*int64) *CreateAgentRequest {
	s.SkillGroupIdList = v
	return s
}

type CreateAgentResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s CreateAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBody) SetMessage(v string) *CreateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgentResponseBody) SetRequestId(v string) *CreateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentResponseBody) SetData(v int64) *CreateAgentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAgentResponseBody) SetCode(v string) *CreateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgentResponseBody) SetSuccess(v bool) *CreateAgentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgentResponseBody) SetHttpStatusCode(v int64) *CreateAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

type CreateAgentResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentResponse) SetHeaders(v map[string]*string) *CreateAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentResponse) SetBody(v *CreateAgentResponseBody) *CreateAgentResponse {
	s.Body = v
	return s
}

type DeleteEntityRouteRequest struct {
	UniqueId   *int64  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteEntityRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteEntityRouteRequest) SetUniqueId(v int64) *DeleteEntityRouteRequest {
	s.UniqueId = &v
	return s
}

func (s *DeleteEntityRouteRequest) SetInstanceId(v string) *DeleteEntityRouteRequest {
	s.InstanceId = &v
	return s
}

type DeleteEntityRouteResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEntityRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEntityRouteResponseBody) SetMessage(v string) *DeleteEntityRouteResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEntityRouteResponseBody) SetRequestId(v string) *DeleteEntityRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEntityRouteResponseBody) SetCode(v string) *DeleteEntityRouteResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEntityRouteResponseBody) SetSuccess(v bool) *DeleteEntityRouteResponseBody {
	s.Success = &v
	return s
}

type DeleteEntityRouteResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEntityRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEntityRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteEntityRouteResponse) SetHeaders(v map[string]*string) *DeleteEntityRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteEntityRouteResponse) SetBody(v *DeleteEntityRouteResponseBody) *DeleteEntityRouteResponse {
	s.Body = v
	return s
}

type GetHotlineGroupDetailReportRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate   *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate     *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DepIds      []*int  `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	GroupIds    []*int  `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
}

func (s GetHotlineGroupDetailReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportRequest) SetCurrentPage(v int32) *GetHotlineGroupDetailReportRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetPageSize(v int32) *GetHotlineGroupDetailReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetStartDate(v int64) *GetHotlineGroupDetailReportRequest {
	s.StartDate = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetEndDate(v int64) *GetHotlineGroupDetailReportRequest {
	s.EndDate = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetInstanceId(v string) *GetHotlineGroupDetailReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetDepIds(v []*int) *GetHotlineGroupDetailReportRequest {
	s.DepIds = v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetGroupIds(v []*int) *GetHotlineGroupDetailReportRequest {
	s.GroupIds = v
	return s
}

type GetHotlineGroupDetailReportResponseBody struct {
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetHotlineGroupDetailReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineGroupDetailReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponseBody) SetMessage(v string) *GetHotlineGroupDetailReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetRequestId(v string) *GetHotlineGroupDetailReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetData(v *GetHotlineGroupDetailReportResponseBodyData) *GetHotlineGroupDetailReportResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetCode(v string) *GetHotlineGroupDetailReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetSuccess(v string) *GetHotlineGroupDetailReportResponseBody {
	s.Success = &v
	return s
}

type GetHotlineGroupDetailReportResponseBodyData struct {
	Rows     []map[string]interface{}                              `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	PageSize *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                                `json:"Total,omitempty" xml:"Total,omitempty"`
	Columns  []*GetHotlineGroupDetailReportResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Page     *int32                                                `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s GetHotlineGroupDetailReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetRows(v []map[string]interface{}) *GetHotlineGroupDetailReportResponseBodyData {
	s.Rows = v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetPageSize(v int32) *GetHotlineGroupDetailReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetTotal(v int32) *GetHotlineGroupDetailReportResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetColumns(v []*GetHotlineGroupDetailReportResponseBodyDataColumns) *GetHotlineGroupDetailReportResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetPage(v int32) *GetHotlineGroupDetailReportResponseBodyData {
	s.Page = &v
	return s
}

type GetHotlineGroupDetailReportResponseBodyDataColumns struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetHotlineGroupDetailReportResponseBodyDataColumns) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponseBodyDataColumns) SetKey(v string) *GetHotlineGroupDetailReportResponseBodyDataColumns {
	s.Key = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyDataColumns) SetTitle(v string) *GetHotlineGroupDetailReportResponseBodyDataColumns {
	s.Title = &v
	return s
}

type GetHotlineGroupDetailReportResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineGroupDetailReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineGroupDetailReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponse) SetHeaders(v map[string]*string) *GetHotlineGroupDetailReportResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineGroupDetailReportResponse) SetBody(v *GetHotlineGroupDetailReportResponseBody) *GetHotlineGroupDetailReportResponse {
	s.Body = v
	return s
}

type CreateTicketRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TemplateId  *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	CategoryId  *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreatorId   *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorType *int32  `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	MemberId    *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName  *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	FromInfo    *string `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	Priority    *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	CarbonCopy  *string `json:"CarbonCopy,omitempty" xml:"CarbonCopy,omitempty"`
	FormData    *string `json:"FormData,omitempty" xml:"FormData,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) SetClientToken(v string) *CreateTicketRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTicketRequest) SetInstanceId(v string) *CreateTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTicketRequest) SetTemplateId(v int64) *CreateTicketRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateTicketRequest) SetCategoryId(v int64) *CreateTicketRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateTicketRequest) SetCreatorId(v int64) *CreateTicketRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTicketRequest) SetCreatorType(v int32) *CreateTicketRequest {
	s.CreatorType = &v
	return s
}

func (s *CreateTicketRequest) SetCreatorName(v string) *CreateTicketRequest {
	s.CreatorName = &v
	return s
}

func (s *CreateTicketRequest) SetMemberId(v int64) *CreateTicketRequest {
	s.MemberId = &v
	return s
}

func (s *CreateTicketRequest) SetMemberName(v string) *CreateTicketRequest {
	s.MemberName = &v
	return s
}

func (s *CreateTicketRequest) SetFromInfo(v string) *CreateTicketRequest {
	s.FromInfo = &v
	return s
}

func (s *CreateTicketRequest) SetPriority(v int32) *CreateTicketRequest {
	s.Priority = &v
	return s
}

func (s *CreateTicketRequest) SetCarbonCopy(v string) *CreateTicketRequest {
	s.CarbonCopy = &v
	return s
}

func (s *CreateTicketRequest) SetFormData(v string) *CreateTicketRequest {
	s.FormData = &v
	return s
}

type CreateTicketResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) SetMessage(v string) *CreateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTicketResponseBody) SetRequestId(v string) *CreateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicketResponseBody) SetData(v int64) *CreateTicketResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTicketResponseBody) SetCode(v string) *CreateTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTicketResponseBody) SetSuccess(v bool) *CreateTicketResponseBody {
	s.Success = &v
	return s
}

type CreateTicketResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateTicketResponse) SetHeaders(v map[string]*string) *CreateTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateTicketResponse) SetBody(v *CreateTicketResponseBody) *CreateTicketResponse {
	s.Body = v
	return s
}

type UpdateCustomerRequest struct {
	ProdLineId  *int64  `json:"ProdLineId,omitempty" xml:"ProdLineId,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TypeCode    *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ManagerName *string `json:"ManagerName,omitempty" xml:"ManagerName,omitempty"`
	Contacter   *string `json:"Contacter,omitempty" xml:"Contacter,omitempty"`
	Industry    *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	Position    *string `json:"Position,omitempty" xml:"Position,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Dingding    *string `json:"Dingding,omitempty" xml:"Dingding,omitempty"`
	OuterId     *string `json:"OuterId,omitempty" xml:"OuterId,omitempty"`
	OuterIdType *int32  `json:"OuterIdType,omitempty" xml:"OuterIdType,omitempty"`
	CustomerId  *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
}

func (s UpdateCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomerRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomerRequest) SetProdLineId(v int64) *UpdateCustomerRequest {
	s.ProdLineId = &v
	return s
}

func (s *UpdateCustomerRequest) SetBizType(v string) *UpdateCustomerRequest {
	s.BizType = &v
	return s
}

func (s *UpdateCustomerRequest) SetName(v string) *UpdateCustomerRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomerRequest) SetTypeCode(v string) *UpdateCustomerRequest {
	s.TypeCode = &v
	return s
}

func (s *UpdateCustomerRequest) SetPhone(v string) *UpdateCustomerRequest {
	s.Phone = &v
	return s
}

func (s *UpdateCustomerRequest) SetInstanceId(v string) *UpdateCustomerRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCustomerRequest) SetManagerName(v string) *UpdateCustomerRequest {
	s.ManagerName = &v
	return s
}

func (s *UpdateCustomerRequest) SetContacter(v string) *UpdateCustomerRequest {
	s.Contacter = &v
	return s
}

func (s *UpdateCustomerRequest) SetIndustry(v string) *UpdateCustomerRequest {
	s.Industry = &v
	return s
}

func (s *UpdateCustomerRequest) SetPosition(v string) *UpdateCustomerRequest {
	s.Position = &v
	return s
}

func (s *UpdateCustomerRequest) SetEmail(v string) *UpdateCustomerRequest {
	s.Email = &v
	return s
}

func (s *UpdateCustomerRequest) SetDingding(v string) *UpdateCustomerRequest {
	s.Dingding = &v
	return s
}

func (s *UpdateCustomerRequest) SetOuterId(v string) *UpdateCustomerRequest {
	s.OuterId = &v
	return s
}

func (s *UpdateCustomerRequest) SetOuterIdType(v int32) *UpdateCustomerRequest {
	s.OuterIdType = &v
	return s
}

func (s *UpdateCustomerRequest) SetCustomerId(v int64) *UpdateCustomerRequest {
	s.CustomerId = &v
	return s
}

type UpdateCustomerResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomerResponseBody) SetMessage(v string) *UpdateCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCustomerResponseBody) SetRequestId(v string) *UpdateCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomerResponseBody) SetData(v int64) *UpdateCustomerResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCustomerResponseBody) SetCode(v string) *UpdateCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCustomerResponseBody) SetSuccess(v bool) *UpdateCustomerResponseBody {
	s.Success = &v
	return s
}

type UpdateCustomerResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomerResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomerResponse) SetHeaders(v map[string]*string) *UpdateCustomerResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomerResponse) SetBody(v *UpdateCustomerResponseBody) *UpdateCustomerResponse {
	s.Body = v
	return s
}

type AnswerCallRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
}

func (s AnswerCallRequest) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallRequest) GoString() string {
	return s.String()
}

func (s *AnswerCallRequest) SetClientToken(v string) *AnswerCallRequest {
	s.ClientToken = &v
	return s
}

func (s *AnswerCallRequest) SetInstanceId(v string) *AnswerCallRequest {
	s.InstanceId = &v
	return s
}

func (s *AnswerCallRequest) SetAccountName(v string) *AnswerCallRequest {
	s.AccountName = &v
	return s
}

func (s *AnswerCallRequest) SetCallId(v string) *AnswerCallRequest {
	s.CallId = &v
	return s
}

func (s *AnswerCallRequest) SetJobId(v string) *AnswerCallRequest {
	s.JobId = &v
	return s
}

func (s *AnswerCallRequest) SetConnectionId(v string) *AnswerCallRequest {
	s.ConnectionId = &v
	return s
}

type AnswerCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AnswerCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallResponseBody) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBody) SetMessage(v string) *AnswerCallResponseBody {
	s.Message = &v
	return s
}

func (s *AnswerCallResponseBody) SetRequestId(v string) *AnswerCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnswerCallResponseBody) SetCode(v string) *AnswerCallResponseBody {
	s.Code = &v
	return s
}

func (s *AnswerCallResponseBody) SetSuccess(v bool) *AnswerCallResponseBody {
	s.Success = &v
	return s
}

type AnswerCallResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AnswerCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AnswerCallResponse) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallResponse) GoString() string {
	return s.String()
}

func (s *AnswerCallResponse) SetHeaders(v map[string]*string) *AnswerCallResponse {
	s.Headers = v
	return s
}

func (s *AnswerCallResponse) SetBody(v *AnswerCallResponseBody) *AnswerCallResponse {
	s.Body = v
	return s
}

type DeleteAgentRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DeleteAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentRequest) SetClientToken(v string) *DeleteAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAgentRequest) SetInstanceId(v string) *DeleteAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAgentRequest) SetAccountName(v string) *DeleteAgentRequest {
	s.AccountName = &v
	return s
}

type DeleteAgentResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s DeleteAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponseBody) SetMessage(v string) *DeleteAgentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAgentResponseBody) SetRequestId(v string) *DeleteAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentResponseBody) SetCode(v string) *DeleteAgentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAgentResponseBody) SetSuccess(v bool) *DeleteAgentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAgentResponseBody) SetHttpStatusCode(v int64) *DeleteAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

type DeleteAgentResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAgentResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponse) SetHeaders(v map[string]*string) *DeleteAgentResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentResponse) SetBody(v *DeleteAgentResponseBody) *DeleteAgentResponse {
	s.Body = v
	return s
}

type GetEntityTagRelationRequest struct {
	EntityId   *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetEntityTagRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEntityTagRelationRequest) GoString() string {
	return s.String()
}

func (s *GetEntityTagRelationRequest) SetEntityId(v string) *GetEntityTagRelationRequest {
	s.EntityId = &v
	return s
}

func (s *GetEntityTagRelationRequest) SetEntityType(v string) *GetEntityTagRelationRequest {
	s.EntityType = &v
	return s
}

func (s *GetEntityTagRelationRequest) SetInstanceId(v string) *GetEntityTagRelationRequest {
	s.InstanceId = &v
	return s
}

type GetEntityTagRelationResponseBody struct {
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*GetEntityTagRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEntityTagRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEntityTagRelationResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityTagRelationResponseBody) SetMessage(v string) *GetEntityTagRelationResponseBody {
	s.Message = &v
	return s
}

func (s *GetEntityTagRelationResponseBody) SetRequestId(v string) *GetEntityTagRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntityTagRelationResponseBody) SetData(v []*GetEntityTagRelationResponseBodyData) *GetEntityTagRelationResponseBody {
	s.Data = v
	return s
}

func (s *GetEntityTagRelationResponseBody) SetCode(v string) *GetEntityTagRelationResponseBody {
	s.Code = &v
	return s
}

func (s *GetEntityTagRelationResponseBody) SetSuccess(v bool) *GetEntityTagRelationResponseBody {
	s.Success = &v
	return s
}

type GetEntityTagRelationResponseBodyData struct {
	TagName      *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	TagGroupCode *string `json:"TagGroupCode,omitempty" xml:"TagGroupCode,omitempty"`
	EntityId     *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	TagCode      *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	EntityType   *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	TagGroupName *string `json:"TagGroupName,omitempty" xml:"TagGroupName,omitempty"`
}

func (s GetEntityTagRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEntityTagRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEntityTagRelationResponseBodyData) SetTagName(v string) *GetEntityTagRelationResponseBodyData {
	s.TagName = &v
	return s
}

func (s *GetEntityTagRelationResponseBodyData) SetTagGroupCode(v string) *GetEntityTagRelationResponseBodyData {
	s.TagGroupCode = &v
	return s
}

func (s *GetEntityTagRelationResponseBodyData) SetEntityId(v string) *GetEntityTagRelationResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *GetEntityTagRelationResponseBodyData) SetTagCode(v string) *GetEntityTagRelationResponseBodyData {
	s.TagCode = &v
	return s
}

func (s *GetEntityTagRelationResponseBodyData) SetEntityType(v string) *GetEntityTagRelationResponseBodyData {
	s.EntityType = &v
	return s
}

func (s *GetEntityTagRelationResponseBodyData) SetTagGroupName(v string) *GetEntityTagRelationResponseBodyData {
	s.TagGroupName = &v
	return s
}

type GetEntityTagRelationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEntityTagRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEntityTagRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEntityTagRelationResponse) GoString() string {
	return s.String()
}

func (s *GetEntityTagRelationResponse) SetHeaders(v map[string]*string) *GetEntityTagRelationResponse {
	s.Headers = v
	return s
}

func (s *GetEntityTagRelationResponse) SetBody(v *GetEntityTagRelationResponseBody) *GetEntityTagRelationResponse {
	s.Body = v
	return s
}

type GetHotlineAgentDetailRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetHotlineAgentDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailRequest) SetClientToken(v string) *GetHotlineAgentDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHotlineAgentDetailRequest) SetInstanceId(v string) *GetHotlineAgentDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineAgentDetailRequest) SetAccountName(v string) *GetHotlineAgentDetailRequest {
	s.AccountName = &v
	return s
}

type GetHotlineAgentDetailResponseBody struct {
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetHotlineAgentDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetHotlineAgentDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailResponseBody) SetMessage(v string) *GetHotlineAgentDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetRequestId(v string) *GetHotlineAgentDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetData(v *GetHotlineAgentDetailResponseBodyData) *GetHotlineAgentDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetCode(v string) *GetHotlineAgentDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetSuccess(v bool) *GetHotlineAgentDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetHttpStatusCode(v int64) *GetHotlineAgentDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetHotlineAgentDetailResponseBodyData struct {
	AgentStatusCode *string `json:"AgentStatusCode,omitempty" xml:"AgentStatusCode,omitempty"`
	Token           *string `json:"Token,omitempty" xml:"Token,omitempty"`
	AgentId         *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	Assigned        *bool   `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	RestType        *int32  `json:"RestType,omitempty" xml:"RestType,omitempty"`
	AgentStatus     *int32  `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	TenantId        *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetHotlineAgentDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAgentStatusCode(v string) *GetHotlineAgentDetailResponseBodyData {
	s.AgentStatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetToken(v string) *GetHotlineAgentDetailResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAgentId(v int64) *GetHotlineAgentDetailResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAssigned(v bool) *GetHotlineAgentDetailResponseBodyData {
	s.Assigned = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetRestType(v int32) *GetHotlineAgentDetailResponseBodyData {
	s.RestType = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAgentStatus(v int32) *GetHotlineAgentDetailResponseBodyData {
	s.AgentStatus = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetTenantId(v int64) *GetHotlineAgentDetailResponseBodyData {
	s.TenantId = &v
	return s
}

type GetHotlineAgentDetailResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineAgentDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineAgentDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailResponse) SetHeaders(v map[string]*string) *GetHotlineAgentDetailResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentDetailResponse) SetBody(v *GetHotlineAgentDetailResponseBody) *GetHotlineAgentDetailResponse {
	s.Body = v
	return s
}

type SuspendHotlineServiceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SuspendHotlineServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendHotlineServiceRequest) GoString() string {
	return s.String()
}

func (s *SuspendHotlineServiceRequest) SetClientToken(v string) *SuspendHotlineServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *SuspendHotlineServiceRequest) SetInstanceId(v string) *SuspendHotlineServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *SuspendHotlineServiceRequest) SetAccountName(v string) *SuspendHotlineServiceRequest {
	s.AccountName = &v
	return s
}

func (s *SuspendHotlineServiceRequest) SetType(v int32) *SuspendHotlineServiceRequest {
	s.Type = &v
	return s
}

type SuspendHotlineServiceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendHotlineServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendHotlineServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendHotlineServiceResponseBody) SetMessage(v string) *SuspendHotlineServiceResponseBody {
	s.Message = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) SetRequestId(v string) *SuspendHotlineServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) SetCode(v string) *SuspendHotlineServiceResponseBody {
	s.Code = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) SetSuccess(v bool) *SuspendHotlineServiceResponseBody {
	s.Success = &v
	return s
}

type SuspendHotlineServiceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SuspendHotlineServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendHotlineServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendHotlineServiceResponse) GoString() string {
	return s.String()
}

func (s *SuspendHotlineServiceResponse) SetHeaders(v map[string]*string) *SuspendHotlineServiceResponse {
	s.Headers = v
	return s
}

func (s *SuspendHotlineServiceResponse) SetBody(v *SuspendHotlineServiceResponseBody) *SuspendHotlineServiceResponse {
	s.Body = v
	return s
}

type GetCallsPerDayRequest struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DataIdStart      *string `json:"DataIdStart,omitempty" xml:"DataIdStart,omitempty"`
	DataIdEnd        *string `json:"DataIdEnd,omitempty" xml:"DataIdEnd,omitempty"`
	DataId           *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	HourId           *string `json:"HourId,omitempty" xml:"HourId,omitempty"`
	MinuteId         *string `json:"MinuteId,omitempty" xml:"MinuteId,omitempty"`
	PhoneNumbers     *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	HavePhoneNumbers *string `json:"HavePhoneNumbers,omitempty" xml:"HavePhoneNumbers,omitempty"`
	PageNo           *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize         *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetCallsPerDayRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCallsPerDayRequest) GoString() string {
	return s.String()
}

func (s *GetCallsPerDayRequest) SetInstanceId(v string) *GetCallsPerDayRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCallsPerDayRequest) SetDataIdStart(v string) *GetCallsPerDayRequest {
	s.DataIdStart = &v
	return s
}

func (s *GetCallsPerDayRequest) SetDataIdEnd(v string) *GetCallsPerDayRequest {
	s.DataIdEnd = &v
	return s
}

func (s *GetCallsPerDayRequest) SetDataId(v string) *GetCallsPerDayRequest {
	s.DataId = &v
	return s
}

func (s *GetCallsPerDayRequest) SetHourId(v string) *GetCallsPerDayRequest {
	s.HourId = &v
	return s
}

func (s *GetCallsPerDayRequest) SetMinuteId(v string) *GetCallsPerDayRequest {
	s.MinuteId = &v
	return s
}

func (s *GetCallsPerDayRequest) SetPhoneNumbers(v string) *GetCallsPerDayRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *GetCallsPerDayRequest) SetHavePhoneNumbers(v string) *GetCallsPerDayRequest {
	s.HavePhoneNumbers = &v
	return s
}

func (s *GetCallsPerDayRequest) SetPageNo(v int64) *GetCallsPerDayRequest {
	s.PageNo = &v
	return s
}

func (s *GetCallsPerDayRequest) SetPageSize(v int64) *GetCallsPerDayRequest {
	s.PageSize = &v
	return s
}

type GetCallsPerDayResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Data      *GetCallsPerDayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetCallsPerDayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCallsPerDayResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallsPerDayResponseBody) SetMessage(v string) *GetCallsPerDayResponseBody {
	s.Message = &v
	return s
}

func (s *GetCallsPerDayResponseBody) SetRequestId(v string) *GetCallsPerDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCallsPerDayResponseBody) SetCode(v string) *GetCallsPerDayResponseBody {
	s.Code = &v
	return s
}

func (s *GetCallsPerDayResponseBody) SetSuccess(v string) *GetCallsPerDayResponseBody {
	s.Success = &v
	return s
}

func (s *GetCallsPerDayResponseBody) SetData(v *GetCallsPerDayResponseBodyData) *GetCallsPerDayResponseBody {
	s.Data = v
	return s
}

type GetCallsPerDayResponseBodyData struct {
	TotalNum                *int64                                                   `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	PageSize                *int64                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo                  *string                                                  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	CallsPerdayResponseList []*GetCallsPerDayResponseBodyDataCallsPerdayResponseList `json:"CallsPerdayResponseList,omitempty" xml:"CallsPerdayResponseList,omitempty" type:"Repeated"`
}

func (s GetCallsPerDayResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCallsPerDayResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCallsPerDayResponseBodyData) SetTotalNum(v int64) *GetCallsPerDayResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *GetCallsPerDayResponseBodyData) SetPageSize(v int64) *GetCallsPerDayResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetCallsPerDayResponseBodyData) SetPageNo(v string) *GetCallsPerDayResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetCallsPerDayResponseBodyData) SetCallsPerdayResponseList(v []*GetCallsPerDayResponseBodyDataCallsPerdayResponseList) *GetCallsPerDayResponseBodyData {
	s.CallsPerdayResponseList = v
	return s
}

type GetCallsPerDayResponseBodyDataCallsPerdayResponseList struct {
	DataId     *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	HourId     *string `json:"HourId,omitempty" xml:"HourId,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	PhoneNum   *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	CallOutCnt *string `json:"CallOutCnt,omitempty" xml:"CallOutCnt,omitempty"`
	CallInCnt  *string `json:"CallInCnt,omitempty" xml:"CallInCnt,omitempty"`
	MinuteId   *string `json:"MinuteId,omitempty" xml:"MinuteId,omitempty"`
}

func (s GetCallsPerDayResponseBodyDataCallsPerdayResponseList) String() string {
	return tea.Prettify(s)
}

func (s GetCallsPerDayResponseBodyDataCallsPerdayResponseList) GoString() string {
	return s.String()
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetDataId(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.DataId = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetHourId(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.HourId = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetTenantId(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.TenantId = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetTenantName(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.TenantName = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetPhoneNum(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.PhoneNum = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetCallOutCnt(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.CallOutCnt = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetCallInCnt(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.CallInCnt = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetMinuteId(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.MinuteId = &v
	return s
}

type GetCallsPerDayResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCallsPerDayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCallsPerDayResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCallsPerDayResponse) GoString() string {
	return s.String()
}

func (s *GetCallsPerDayResponse) SetHeaders(v map[string]*string) *GetCallsPerDayResponse {
	s.Headers = v
	return s
}

func (s *GetCallsPerDayResponse) SetBody(v *GetCallsPerDayResponseBody) *GetCallsPerDayResponse {
	s.Body = v
	return s
}

type FetchCallRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
}

func (s FetchCallRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchCallRequest) GoString() string {
	return s.String()
}

func (s *FetchCallRequest) SetClientToken(v string) *FetchCallRequest {
	s.ClientToken = &v
	return s
}

func (s *FetchCallRequest) SetInstanceId(v string) *FetchCallRequest {
	s.InstanceId = &v
	return s
}

func (s *FetchCallRequest) SetAccountName(v string) *FetchCallRequest {
	s.AccountName = &v
	return s
}

func (s *FetchCallRequest) SetCallId(v string) *FetchCallRequest {
	s.CallId = &v
	return s
}

func (s *FetchCallRequest) SetJobId(v string) *FetchCallRequest {
	s.JobId = &v
	return s
}

func (s *FetchCallRequest) SetConnectionId(v string) *FetchCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *FetchCallRequest) SetHoldConnectionId(v string) *FetchCallRequest {
	s.HoldConnectionId = &v
	return s
}

type FetchCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FetchCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchCallResponseBody) GoString() string {
	return s.String()
}

func (s *FetchCallResponseBody) SetMessage(v string) *FetchCallResponseBody {
	s.Message = &v
	return s
}

func (s *FetchCallResponseBody) SetRequestId(v string) *FetchCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchCallResponseBody) SetCode(v string) *FetchCallResponseBody {
	s.Code = &v
	return s
}

func (s *FetchCallResponseBody) SetSuccess(v bool) *FetchCallResponseBody {
	s.Success = &v
	return s
}

type FetchCallResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FetchCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FetchCallResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchCallResponse) GoString() string {
	return s.String()
}

func (s *FetchCallResponse) SetHeaders(v map[string]*string) *FetchCallResponse {
	s.Headers = v
	return s
}

func (s *FetchCallResponse) SetBody(v *FetchCallResponseBody) *FetchCallResponse {
	s.Body = v
	return s
}

type GetHotlineAgentDetailReportRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate   *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate     *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	DepIds      []*int  `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	GroupIds    []*int  `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
}

func (s GetHotlineAgentDetailReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportRequest) SetCurrentPage(v int32) *GetHotlineAgentDetailReportRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetPageSize(v int32) *GetHotlineAgentDetailReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetStartDate(v int64) *GetHotlineAgentDetailReportRequest {
	s.StartDate = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetEndDate(v int64) *GetHotlineAgentDetailReportRequest {
	s.EndDate = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetInstanceId(v string) *GetHotlineAgentDetailReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetDepIds(v []*int) *GetHotlineAgentDetailReportRequest {
	s.DepIds = v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetGroupIds(v []*int) *GetHotlineAgentDetailReportRequest {
	s.GroupIds = v
	return s
}

type GetHotlineAgentDetailReportResponseBody struct {
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetHotlineAgentDetailReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineAgentDetailReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponseBody) SetMessage(v string) *GetHotlineAgentDetailReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetRequestId(v string) *GetHotlineAgentDetailReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetData(v *GetHotlineAgentDetailReportResponseBodyData) *GetHotlineAgentDetailReportResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetCode(v string) *GetHotlineAgentDetailReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetSuccess(v string) *GetHotlineAgentDetailReportResponseBody {
	s.Success = &v
	return s
}

type GetHotlineAgentDetailReportResponseBodyData struct {
	Rows     []map[string]interface{}                              `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	PageSize *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                                `json:"Total,omitempty" xml:"Total,omitempty"`
	Columns  []*GetHotlineAgentDetailReportResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Page     *int32                                                `json:"Page,omitempty" xml:"Page,omitempty"`
}

func (s GetHotlineAgentDetailReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetRows(v []map[string]interface{}) *GetHotlineAgentDetailReportResponseBodyData {
	s.Rows = v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetPageSize(v int32) *GetHotlineAgentDetailReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetTotal(v int32) *GetHotlineAgentDetailReportResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetColumns(v []*GetHotlineAgentDetailReportResponseBodyDataColumns) *GetHotlineAgentDetailReportResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetPage(v int32) *GetHotlineAgentDetailReportResponseBodyData {
	s.Page = &v
	return s
}

type GetHotlineAgentDetailReportResponseBodyDataColumns struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetHotlineAgentDetailReportResponseBodyDataColumns) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponseBodyDataColumns) SetKey(v string) *GetHotlineAgentDetailReportResponseBodyDataColumns {
	s.Key = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyDataColumns) SetTitle(v string) *GetHotlineAgentDetailReportResponseBodyDataColumns {
	s.Title = &v
	return s
}

type GetHotlineAgentDetailReportResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineAgentDetailReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineAgentDetailReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponse) SetHeaders(v map[string]*string) *GetHotlineAgentDetailReportResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentDetailReportResponse) SetBody(v *GetHotlineAgentDetailReportResponseBody) *GetHotlineAgentDetailReportResponse {
	s.Body = v
	return s
}

type QueryTicketCountRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperatorId  *int64  `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
}

func (s QueryTicketCountRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketCountRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketCountRequest) SetClientToken(v string) *QueryTicketCountRequest {
	s.ClientToken = &v
	return s
}

func (s *QueryTicketCountRequest) SetInstanceId(v string) *QueryTicketCountRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTicketCountRequest) SetOperatorId(v int64) *QueryTicketCountRequest {
	s.OperatorId = &v
	return s
}

type QueryTicketCountResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTicketCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketCountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTicketCountResponseBody) SetMessage(v string) *QueryTicketCountResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTicketCountResponseBody) SetRequestId(v string) *QueryTicketCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTicketCountResponseBody) SetData(v string) *QueryTicketCountResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTicketCountResponseBody) SetCode(v string) *QueryTicketCountResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTicketCountResponseBody) SetSuccess(v bool) *QueryTicketCountResponseBody {
	s.Success = &v
	return s
}

type QueryTicketCountResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTicketCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTicketCountResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketCountResponse) GoString() string {
	return s.String()
}

func (s *QueryTicketCountResponse) SetHeaders(v map[string]*string) *QueryTicketCountResponse {
	s.Headers = v
	return s
}

func (s *QueryTicketCountResponse) SetBody(v *QueryTicketCountResponseBody) *QueryTicketCountResponse {
	s.Body = v
	return s
}

type AppMessagePushRequest struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 用户编号
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// APP状态
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 过期时间
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
}

func (s AppMessagePushRequest) String() string {
	return tea.Prettify(s)
}

func (s AppMessagePushRequest) GoString() string {
	return s.String()
}

func (s *AppMessagePushRequest) SetInstanceId(v string) *AppMessagePushRequest {
	s.InstanceId = &v
	return s
}

func (s *AppMessagePushRequest) SetUserId(v string) *AppMessagePushRequest {
	s.UserId = &v
	return s
}

func (s *AppMessagePushRequest) SetStatus(v int32) *AppMessagePushRequest {
	s.Status = &v
	return s
}

func (s *AppMessagePushRequest) SetExpirationTime(v int64) *AppMessagePushRequest {
	s.ExpirationTime = &v
	return s
}

type AppMessagePushResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回数据
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 通信码
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AppMessagePushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppMessagePushResponseBody) GoString() string {
	return s.String()
}

func (s *AppMessagePushResponseBody) SetRequestId(v string) *AppMessagePushResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppMessagePushResponseBody) SetData(v string) *AppMessagePushResponseBody {
	s.Data = &v
	return s
}

func (s *AppMessagePushResponseBody) SetCode(v string) *AppMessagePushResponseBody {
	s.Code = &v
	return s
}

func (s *AppMessagePushResponseBody) SetMessage(v string) *AppMessagePushResponseBody {
	s.Message = &v
	return s
}

func (s *AppMessagePushResponseBody) SetSuccess(v bool) *AppMessagePushResponseBody {
	s.Success = &v
	return s
}

type AppMessagePushResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AppMessagePushResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppMessagePushResponse) String() string {
	return tea.Prettify(s)
}

func (s AppMessagePushResponse) GoString() string {
	return s.String()
}

func (s *AppMessagePushResponse) SetHeaders(v map[string]*string) *AppMessagePushResponse {
	s.Headers = v
	return s
}

func (s *AppMessagePushResponse) SetBody(v *AppMessagePushResponseBody) *AppMessagePushResponse {
	s.Body = v
	return s
}

type GetHotlineAgentStatusRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetHotlineAgentStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusRequest) SetInstanceId(v string) *GetHotlineAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineAgentStatusRequest) SetAccountName(v string) *GetHotlineAgentStatusRequest {
	s.AccountName = &v
	return s
}

type GetHotlineAgentStatusResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetHotlineAgentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusResponseBody) SetMessage(v string) *GetHotlineAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetRequestId(v string) *GetHotlineAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetData(v string) *GetHotlineAgentStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetCode(v string) *GetHotlineAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetSuccess(v bool) *GetHotlineAgentStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetHttpStatusCode(v int64) *GetHotlineAgentStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetHotlineAgentStatusResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineAgentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusResponse) SetHeaders(v map[string]*string) *GetHotlineAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentStatusResponse) SetBody(v *GetHotlineAgentStatusResponseBody) *GetHotlineAgentStatusResponse {
	s.Body = v
	return s
}

type GetHotlineWaitingNumberRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetHotlineWaitingNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineWaitingNumberRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineWaitingNumberRequest) SetClientToken(v string) *GetHotlineWaitingNumberRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHotlineWaitingNumberRequest) SetInstanceId(v string) *GetHotlineWaitingNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineWaitingNumberRequest) SetAccountName(v string) *GetHotlineWaitingNumberRequest {
	s.AccountName = &v
	return s
}

type GetHotlineWaitingNumberResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetHotlineWaitingNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineWaitingNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineWaitingNumberResponseBody) SetMessage(v string) *GetHotlineWaitingNumberResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetRequestId(v string) *GetHotlineWaitingNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetData(v int64) *GetHotlineWaitingNumberResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetCode(v string) *GetHotlineWaitingNumberResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetSuccess(v bool) *GetHotlineWaitingNumberResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetHttpStatusCode(v int64) *GetHotlineWaitingNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetHotlineWaitingNumberResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineWaitingNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineWaitingNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineWaitingNumberResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineWaitingNumberResponse) SetHeaders(v map[string]*string) *GetHotlineWaitingNumberResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineWaitingNumberResponse) SetBody(v *GetHotlineWaitingNumberResponseBody) *GetHotlineWaitingNumberResponse {
	s.Body = v
	return s
}

type StartCallRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Caller      *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Callee      *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
}

func (s StartCallRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCallRequest) GoString() string {
	return s.String()
}

func (s *StartCallRequest) SetClientToken(v string) *StartCallRequest {
	s.ClientToken = &v
	return s
}

func (s *StartCallRequest) SetInstanceId(v string) *StartCallRequest {
	s.InstanceId = &v
	return s
}

func (s *StartCallRequest) SetAccountName(v string) *StartCallRequest {
	s.AccountName = &v
	return s
}

func (s *StartCallRequest) SetCaller(v string) *StartCallRequest {
	s.Caller = &v
	return s
}

func (s *StartCallRequest) SetCallee(v string) *StartCallRequest {
	s.Callee = &v
	return s
}

type StartCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartCallResponseBody) SetMessage(v string) *StartCallResponseBody {
	s.Message = &v
	return s
}

func (s *StartCallResponseBody) SetRequestId(v string) *StartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCallResponseBody) SetCode(v string) *StartCallResponseBody {
	s.Code = &v
	return s
}

func (s *StartCallResponseBody) SetSuccess(v bool) *StartCallResponseBody {
	s.Success = &v
	return s
}

type StartCallResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCallResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCallResponse) GoString() string {
	return s.String()
}

func (s *StartCallResponse) SetHeaders(v map[string]*string) *StartCallResponse {
	s.Headers = v
	return s
}

func (s *StartCallResponse) SetBody(v *StartCallResponseBody) *StartCallResponse {
	s.Body = v
	return s
}

type AssignTicketRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TicketId    *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	OperatorId  *int64  `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	AcceptorId  *int64  `json:"AcceptorId,omitempty" xml:"AcceptorId,omitempty"`
}

func (s AssignTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketRequest) GoString() string {
	return s.String()
}

func (s *AssignTicketRequest) SetClientToken(v string) *AssignTicketRequest {
	s.ClientToken = &v
	return s
}

func (s *AssignTicketRequest) SetInstanceId(v string) *AssignTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *AssignTicketRequest) SetTicketId(v int64) *AssignTicketRequest {
	s.TicketId = &v
	return s
}

func (s *AssignTicketRequest) SetOperatorId(v int64) *AssignTicketRequest {
	s.OperatorId = &v
	return s
}

func (s *AssignTicketRequest) SetAcceptorId(v int64) *AssignTicketRequest {
	s.AcceptorId = &v
	return s
}

type AssignTicketResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssignTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketResponseBody) GoString() string {
	return s.String()
}

func (s *AssignTicketResponseBody) SetMessage(v string) *AssignTicketResponseBody {
	s.Message = &v
	return s
}

func (s *AssignTicketResponseBody) SetRequestId(v string) *AssignTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignTicketResponseBody) SetCode(v string) *AssignTicketResponseBody {
	s.Code = &v
	return s
}

func (s *AssignTicketResponseBody) SetSuccess(v bool) *AssignTicketResponseBody {
	s.Success = &v
	return s
}

type AssignTicketResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssignTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssignTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketResponse) GoString() string {
	return s.String()
}

func (s *AssignTicketResponse) SetHeaders(v map[string]*string) *AssignTicketResponse {
	s.Headers = v
	return s
}

func (s *AssignTicketResponse) SetBody(v *AssignTicketResponseBody) *AssignTicketResponse {
	s.Body = v
	return s
}

type HangupCallRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
}

func (s HangupCallRequest) String() string {
	return tea.Prettify(s)
}

func (s HangupCallRequest) GoString() string {
	return s.String()
}

func (s *HangupCallRequest) SetClientToken(v string) *HangupCallRequest {
	s.ClientToken = &v
	return s
}

func (s *HangupCallRequest) SetInstanceId(v string) *HangupCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HangupCallRequest) SetAccountName(v string) *HangupCallRequest {
	s.AccountName = &v
	return s
}

func (s *HangupCallRequest) SetCallId(v string) *HangupCallRequest {
	s.CallId = &v
	return s
}

func (s *HangupCallRequest) SetJobId(v string) *HangupCallRequest {
	s.JobId = &v
	return s
}

func (s *HangupCallRequest) SetConnectionId(v string) *HangupCallRequest {
	s.ConnectionId = &v
	return s
}

type HangupCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HangupCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HangupCallResponseBody) GoString() string {
	return s.String()
}

func (s *HangupCallResponseBody) SetMessage(v string) *HangupCallResponseBody {
	s.Message = &v
	return s
}

func (s *HangupCallResponseBody) SetRequestId(v string) *HangupCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HangupCallResponseBody) SetCode(v string) *HangupCallResponseBody {
	s.Code = &v
	return s
}

func (s *HangupCallResponseBody) SetSuccess(v bool) *HangupCallResponseBody {
	s.Success = &v
	return s
}

type HangupCallResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HangupCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HangupCallResponse) String() string {
	return tea.Prettify(s)
}

func (s HangupCallResponse) GoString() string {
	return s.String()
}

func (s *HangupCallResponse) SetHeaders(v map[string]*string) *HangupCallResponse {
	s.Headers = v
	return s
}

func (s *HangupCallResponse) SetBody(v *HangupCallResponseBody) *HangupCallResponse {
	s.Body = v
	return s
}

type GetOutbounNumListRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetOutbounNumListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListRequest) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListRequest) SetClientToken(v string) *GetOutbounNumListRequest {
	s.ClientToken = &v
	return s
}

func (s *GetOutbounNumListRequest) SetInstanceId(v string) *GetOutbounNumListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOutbounNumListRequest) SetAccountName(v string) *GetOutbounNumListRequest {
	s.AccountName = &v
	return s
}

type GetOutbounNumListResponseBody struct {
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetOutbounNumListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetOutbounNumListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBody) SetMessage(v string) *GetOutbounNumListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetRequestId(v string) *GetOutbounNumListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetData(v *GetOutbounNumListResponseBodyData) *GetOutbounNumListResponseBody {
	s.Data = v
	return s
}

func (s *GetOutbounNumListResponseBody) SetCode(v string) *GetOutbounNumListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetSuccess(v bool) *GetOutbounNumListResponseBody {
	s.Success = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetHttpStatusCode(v int64) *GetOutbounNumListResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetOutbounNumListResponseBodyData struct {
	NumGroup []*GetOutbounNumListResponseBodyDataNumGroup `json:"NumGroup,omitempty" xml:"NumGroup,omitempty" type:"Repeated"`
	Num      []*GetOutbounNumListResponseBodyDataNum      `json:"Num,omitempty" xml:"Num,omitempty" type:"Repeated"`
}

func (s GetOutbounNumListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBodyData) SetNumGroup(v []*GetOutbounNumListResponseBodyDataNumGroup) *GetOutbounNumListResponseBodyData {
	s.NumGroup = v
	return s
}

func (s *GetOutbounNumListResponseBodyData) SetNum(v []*GetOutbounNumListResponseBodyDataNum) *GetOutbounNumListResponseBodyData {
	s.Num = v
	return s
}

type GetOutbounNumListResponseBodyDataNumGroup struct {
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetOutbounNumListResponseBodyDataNumGroup) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponseBodyDataNumGroup) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) SetType(v int32) *GetOutbounNumListResponseBodyDataNumGroup {
	s.Type = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) SetValue(v string) *GetOutbounNumListResponseBodyDataNumGroup {
	s.Value = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) SetDescription(v string) *GetOutbounNumListResponseBodyDataNumGroup {
	s.Description = &v
	return s
}

type GetOutbounNumListResponseBodyDataNum struct {
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetOutbounNumListResponseBodyDataNum) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponseBodyDataNum) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBodyDataNum) SetType(v int32) *GetOutbounNumListResponseBodyDataNum {
	s.Type = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNum) SetValue(v string) *GetOutbounNumListResponseBodyDataNum {
	s.Value = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNum) SetDescription(v string) *GetOutbounNumListResponseBodyDataNum {
	s.Description = &v
	return s
}

type GetOutbounNumListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOutbounNumListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOutbounNumListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponse) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponse) SetHeaders(v map[string]*string) *GetOutbounNumListResponse {
	s.Headers = v
	return s
}

func (s *GetOutbounNumListResponse) SetBody(v *GetOutbounNumListResponseBody) *GetOutbounNumListResponse {
	s.Body = v
	return s
}

type CreateTicketWithBizDataRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TemplateId  *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	CategoryId  *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreatorId   *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorType *int32  `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	MemberId    *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName  *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	FromInfo    *string `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	Priority    *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	CarbonCopy  *string `json:"CarbonCopy,omitempty" xml:"CarbonCopy,omitempty"`
	FormData    *string `json:"FormData,omitempty" xml:"FormData,omitempty"`
	BizData     *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
}

func (s CreateTicketWithBizDataRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketWithBizDataRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketWithBizDataRequest) SetClientToken(v string) *CreateTicketWithBizDataRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetInstanceId(v string) *CreateTicketWithBizDataRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetTemplateId(v int64) *CreateTicketWithBizDataRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetCategoryId(v int64) *CreateTicketWithBizDataRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetCreatorId(v int64) *CreateTicketWithBizDataRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetCreatorType(v int32) *CreateTicketWithBizDataRequest {
	s.CreatorType = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetCreatorName(v string) *CreateTicketWithBizDataRequest {
	s.CreatorName = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetMemberId(v int64) *CreateTicketWithBizDataRequest {
	s.MemberId = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetMemberName(v string) *CreateTicketWithBizDataRequest {
	s.MemberName = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetFromInfo(v string) *CreateTicketWithBizDataRequest {
	s.FromInfo = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetPriority(v int32) *CreateTicketWithBizDataRequest {
	s.Priority = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetCarbonCopy(v string) *CreateTicketWithBizDataRequest {
	s.CarbonCopy = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetFormData(v string) *CreateTicketWithBizDataRequest {
	s.FormData = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetBizData(v string) *CreateTicketWithBizDataRequest {
	s.BizData = &v
	return s
}

type CreateTicketWithBizDataResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTicketWithBizDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketWithBizDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketWithBizDataResponseBody) SetMessage(v string) *CreateTicketWithBizDataResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTicketWithBizDataResponseBody) SetRequestId(v string) *CreateTicketWithBizDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicketWithBizDataResponseBody) SetData(v int64) *CreateTicketWithBizDataResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTicketWithBizDataResponseBody) SetCode(v string) *CreateTicketWithBizDataResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTicketWithBizDataResponseBody) SetSuccess(v bool) *CreateTicketWithBizDataResponseBody {
	s.Success = &v
	return s
}

type CreateTicketWithBizDataResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTicketWithBizDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTicketWithBizDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketWithBizDataResponse) GoString() string {
	return s.String()
}

func (s *CreateTicketWithBizDataResponse) SetHeaders(v map[string]*string) *CreateTicketWithBizDataResponse {
	s.Headers = v
	return s
}

func (s *CreateTicketWithBizDataResponse) SetBody(v *CreateTicketWithBizDataResponseBody) *CreateTicketWithBizDataResponse {
	s.Body = v
	return s
}

type SearchTicketByPhoneRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Phone        *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TicketStatus *string `json:"TicketStatus,omitempty" xml:"TicketStatus,omitempty"`
	PageNo       *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s SearchTicketByPhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByPhoneRequest) GoString() string {
	return s.String()
}

func (s *SearchTicketByPhoneRequest) SetClientToken(v string) *SearchTicketByPhoneRequest {
	s.ClientToken = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetInstanceId(v string) *SearchTicketByPhoneRequest {
	s.InstanceId = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetPhone(v string) *SearchTicketByPhoneRequest {
	s.Phone = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetTemplateId(v int64) *SearchTicketByPhoneRequest {
	s.TemplateId = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetTicketStatus(v string) *SearchTicketByPhoneRequest {
	s.TicketStatus = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetPageNo(v int32) *SearchTicketByPhoneRequest {
	s.PageNo = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetPageSize(v int32) *SearchTicketByPhoneRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetStartTime(v int64) *SearchTicketByPhoneRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetEndTime(v int64) *SearchTicketByPhoneRequest {
	s.EndTime = &v
	return s
}

type SearchTicketByPhoneResponseBody struct {
	OnePageSize  *int32                                 `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message      *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	TotalPage    *int32                                 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	TotalResults *int32                                 `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
	PageNo       *int32                                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	Data         []*SearchTicketByPhoneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchTicketByPhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTicketByPhoneResponseBody) SetOnePageSize(v int32) *SearchTicketByPhoneResponseBody {
	s.OnePageSize = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetRequestId(v string) *SearchTicketByPhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetMessage(v string) *SearchTicketByPhoneResponseBody {
	s.Message = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetTotalPage(v int32) *SearchTicketByPhoneResponseBody {
	s.TotalPage = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetTotalResults(v int32) *SearchTicketByPhoneResponseBody {
	s.TotalResults = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetPageNo(v int32) *SearchTicketByPhoneResponseBody {
	s.PageNo = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetData(v []*SearchTicketByPhoneResponseBodyData) *SearchTicketByPhoneResponseBody {
	s.Data = v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetCode(v string) *SearchTicketByPhoneResponseBody {
	s.Code = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetSuccess(v bool) *SearchTicketByPhoneResponseBody {
	s.Success = &v
	return s
}

type SearchTicketByPhoneResponseBodyData struct {
	MemberName   *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	CarbonCopy   *string `json:"CarbonCopy,omitempty" xml:"CarbonCopy,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ServiceId    *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	TicketId     *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	CreatorId    *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	FormData     *string `json:"FormData,omitempty" xml:"FormData,omitempty"`
	FromInfo     *string `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	ModifiedTime *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	TaskStatus   *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	CreatorName  *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	CategoryId   *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreatorType  *int32  `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	MemberId     *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	CaseStatus   *int32  `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SearchTicketByPhoneResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByPhoneResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchTicketByPhoneResponseBodyData) SetMemberName(v string) *SearchTicketByPhoneResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCarbonCopy(v string) *SearchTicketByPhoneResponseBodyData {
	s.CarbonCopy = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCreateTime(v int64) *SearchTicketByPhoneResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetServiceId(v int64) *SearchTicketByPhoneResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetTicketId(v int64) *SearchTicketByPhoneResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetPriority(v int32) *SearchTicketByPhoneResponseBodyData {
	s.Priority = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCreatorId(v int64) *SearchTicketByPhoneResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetFormData(v string) *SearchTicketByPhoneResponseBodyData {
	s.FormData = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetFromInfo(v string) *SearchTicketByPhoneResponseBodyData {
	s.FromInfo = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetModifiedTime(v int64) *SearchTicketByPhoneResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetTaskStatus(v string) *SearchTicketByPhoneResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCreatorName(v string) *SearchTicketByPhoneResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCategoryId(v int64) *SearchTicketByPhoneResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCreatorType(v int32) *SearchTicketByPhoneResponseBodyData {
	s.CreatorType = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetMemberId(v int64) *SearchTicketByPhoneResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCaseStatus(v int32) *SearchTicketByPhoneResponseBodyData {
	s.CaseStatus = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetTemplateId(v int64) *SearchTicketByPhoneResponseBodyData {
	s.TemplateId = &v
	return s
}

type SearchTicketByPhoneResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchTicketByPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTicketByPhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByPhoneResponse) GoString() string {
	return s.String()
}

func (s *SearchTicketByPhoneResponse) SetHeaders(v map[string]*string) *SearchTicketByPhoneResponse {
	s.Headers = v
	return s
}

func (s *SearchTicketByPhoneResponse) SetBody(v *SearchTicketByPhoneResponseBody) *SearchTicketByPhoneResponse {
	s.Body = v
	return s
}

type CreateThirdSsoAgentRequest struct {
	// clientToken
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// param1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// param2
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// param3
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// param4
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// param5
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// param6
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty" type:"Repeated"`
	// param7
	RoleIds []*int64 `json:"RoleIds,omitempty" xml:"RoleIds,omitempty" type:"Repeated"`
}

func (s CreateThirdSsoAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateThirdSsoAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateThirdSsoAgentRequest) SetClientToken(v string) *CreateThirdSsoAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetInstanceId(v string) *CreateThirdSsoAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetClientId(v string) *CreateThirdSsoAgentRequest {
	s.ClientId = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetAccountId(v string) *CreateThirdSsoAgentRequest {
	s.AccountId = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetAccountName(v string) *CreateThirdSsoAgentRequest {
	s.AccountName = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetDisplayName(v string) *CreateThirdSsoAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetSkillGroupIds(v []*int64) *CreateThirdSsoAgentRequest {
	s.SkillGroupIds = v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetRoleIds(v []*int64) *CreateThirdSsoAgentRequest {
	s.RoleIds = v
	return s
}

type CreateThirdSsoAgentResponseBody struct {
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// httpStatusCode
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 新创建的坐席id
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateThirdSsoAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateThirdSsoAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateThirdSsoAgentResponseBody) SetMessage(v string) *CreateThirdSsoAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetRequestId(v string) *CreateThirdSsoAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetHttpStatusCode(v int64) *CreateThirdSsoAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetData(v int64) *CreateThirdSsoAgentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetCode(v string) *CreateThirdSsoAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetSuccess(v bool) *CreateThirdSsoAgentResponseBody {
	s.Success = &v
	return s
}

type CreateThirdSsoAgentResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateThirdSsoAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateThirdSsoAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateThirdSsoAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateThirdSsoAgentResponse) SetHeaders(v map[string]*string) *CreateThirdSsoAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateThirdSsoAgentResponse) SetBody(v *CreateThirdSsoAgentResponseBody) *CreateThirdSsoAgentResponse {
	s.Body = v
	return s
}

type CreateEntityIvrRouteRequest struct {
	EntityId             *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName           *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityBizCode        *string `json:"EntityBizCode,omitempty" xml:"EntityBizCode,omitempty"`
	EntityBizCodeType    *string `json:"EntityBizCodeType,omitempty" xml:"EntityBizCodeType,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
	DepartmentId         *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	GroupId              *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ServiceId            *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ExtInfo              *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateEntityIvrRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityIvrRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateEntityIvrRouteRequest) SetEntityId(v string) *CreateEntityIvrRouteRequest {
	s.EntityId = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetEntityName(v string) *CreateEntityIvrRouteRequest {
	s.EntityName = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetEntityBizCode(v string) *CreateEntityIvrRouteRequest {
	s.EntityBizCode = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetEntityBizCodeType(v string) *CreateEntityIvrRouteRequest {
	s.EntityBizCodeType = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetEntityRelationNumber(v string) *CreateEntityIvrRouteRequest {
	s.EntityRelationNumber = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetDepartmentId(v string) *CreateEntityIvrRouteRequest {
	s.DepartmentId = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetGroupId(v int64) *CreateEntityIvrRouteRequest {
	s.GroupId = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetServiceId(v int64) *CreateEntityIvrRouteRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetExtInfo(v string) *CreateEntityIvrRouteRequest {
	s.ExtInfo = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetInstanceId(v string) *CreateEntityIvrRouteRequest {
	s.InstanceId = &v
	return s
}

type CreateEntityIvrRouteResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEntityIvrRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityIvrRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEntityIvrRouteResponseBody) SetMessage(v string) *CreateEntityIvrRouteResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEntityIvrRouteResponseBody) SetRequestId(v string) *CreateEntityIvrRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEntityIvrRouteResponseBody) SetCode(v string) *CreateEntityIvrRouteResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEntityIvrRouteResponseBody) SetSuccess(v bool) *CreateEntityIvrRouteResponseBody {
	s.Success = &v
	return s
}

type CreateEntityIvrRouteResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEntityIvrRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEntityIvrRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityIvrRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateEntityIvrRouteResponse) SetHeaders(v map[string]*string) *CreateEntityIvrRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateEntityIvrRouteResponse) SetBody(v *CreateEntityIvrRouteResponseBody) *CreateEntityIvrRouteResponse {
	s.Body = v
	return s
}

type CloseTicketRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TicketId    *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	ActionItems *string `json:"ActionItems,omitempty" xml:"ActionItems,omitempty"`
	OperatorId  *int64  `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
}

func (s CloseTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseTicketRequest) GoString() string {
	return s.String()
}

func (s *CloseTicketRequest) SetClientToken(v string) *CloseTicketRequest {
	s.ClientToken = &v
	return s
}

func (s *CloseTicketRequest) SetInstanceId(v string) *CloseTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *CloseTicketRequest) SetTicketId(v int64) *CloseTicketRequest {
	s.TicketId = &v
	return s
}

func (s *CloseTicketRequest) SetActionItems(v string) *CloseTicketRequest {
	s.ActionItems = &v
	return s
}

func (s *CloseTicketRequest) SetOperatorId(v int64) *CloseTicketRequest {
	s.OperatorId = &v
	return s
}

type CloseTicketResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s CloseTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CloseTicketResponseBody) SetMessage(v string) *CloseTicketResponseBody {
	s.Message = &v
	return s
}

func (s *CloseTicketResponseBody) SetRequestId(v string) *CloseTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseTicketResponseBody) SetCode(v string) *CloseTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CloseTicketResponseBody) SetSuccess(v bool) *CloseTicketResponseBody {
	s.Success = &v
	return s
}

func (s *CloseTicketResponseBody) SetHttpStatusCode(v int64) *CloseTicketResponseBody {
	s.HttpStatusCode = &v
	return s
}

type CloseTicketResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseTicketResponse) GoString() string {
	return s.String()
}

func (s *CloseTicketResponse) SetHeaders(v map[string]*string) *CloseTicketResponse {
	s.Headers = v
	return s
}

func (s *CloseTicketResponse) SetBody(v *CloseTicketResponseBody) *CloseTicketResponse {
	s.Body = v
	return s
}

type HoldCallRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
}

func (s HoldCallRequest) String() string {
	return tea.Prettify(s)
}

func (s HoldCallRequest) GoString() string {
	return s.String()
}

func (s *HoldCallRequest) SetClientToken(v string) *HoldCallRequest {
	s.ClientToken = &v
	return s
}

func (s *HoldCallRequest) SetInstanceId(v string) *HoldCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HoldCallRequest) SetAccountName(v string) *HoldCallRequest {
	s.AccountName = &v
	return s
}

func (s *HoldCallRequest) SetCallId(v string) *HoldCallRequest {
	s.CallId = &v
	return s
}

func (s *HoldCallRequest) SetJobId(v string) *HoldCallRequest {
	s.JobId = &v
	return s
}

func (s *HoldCallRequest) SetConnectionId(v string) *HoldCallRequest {
	s.ConnectionId = &v
	return s
}

type HoldCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HoldCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HoldCallResponseBody) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBody) SetMessage(v string) *HoldCallResponseBody {
	s.Message = &v
	return s
}

func (s *HoldCallResponseBody) SetRequestId(v string) *HoldCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HoldCallResponseBody) SetCode(v string) *HoldCallResponseBody {
	s.Code = &v
	return s
}

func (s *HoldCallResponseBody) SetSuccess(v bool) *HoldCallResponseBody {
	s.Success = &v
	return s
}

type HoldCallResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *HoldCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HoldCallResponse) String() string {
	return tea.Prettify(s)
}

func (s HoldCallResponse) GoString() string {
	return s.String()
}

func (s *HoldCallResponse) SetHeaders(v map[string]*string) *HoldCallResponse {
	s.Headers = v
	return s
}

func (s *HoldCallResponse) SetBody(v *HoldCallResponseBody) *HoldCallResponse {
	s.Body = v
	return s
}

type QueryRingDetailListRequest struct {
	PageSize         *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo           *int32                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	StartDate        *int64                 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate          *int64                 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	CallOutStatus    *string                `json:"CallOutStatus,omitempty" xml:"CallOutStatus,omitempty"`
	Extra            *string                `json:"Extra,omitempty" xml:"Extra,omitempty"`
	InstanceId       *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	FromPhoneNumList map[string]interface{} `json:"FromPhoneNumList,omitempty" xml:"FromPhoneNumList,omitempty"`
	ToPhoneNumList   map[string]interface{} `json:"ToPhoneNumList,omitempty" xml:"ToPhoneNumList,omitempty"`
}

func (s QueryRingDetailListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRingDetailListRequest) GoString() string {
	return s.String()
}

func (s *QueryRingDetailListRequest) SetPageSize(v int32) *QueryRingDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRingDetailListRequest) SetPageNo(v int32) *QueryRingDetailListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryRingDetailListRequest) SetStartDate(v int64) *QueryRingDetailListRequest {
	s.StartDate = &v
	return s
}

func (s *QueryRingDetailListRequest) SetEndDate(v int64) *QueryRingDetailListRequest {
	s.EndDate = &v
	return s
}

func (s *QueryRingDetailListRequest) SetCallOutStatus(v string) *QueryRingDetailListRequest {
	s.CallOutStatus = &v
	return s
}

func (s *QueryRingDetailListRequest) SetExtra(v string) *QueryRingDetailListRequest {
	s.Extra = &v
	return s
}

func (s *QueryRingDetailListRequest) SetInstanceId(v string) *QueryRingDetailListRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryRingDetailListRequest) SetFromPhoneNumList(v map[string]interface{}) *QueryRingDetailListRequest {
	s.FromPhoneNumList = v
	return s
}

func (s *QueryRingDetailListRequest) SetToPhoneNumList(v map[string]interface{}) *QueryRingDetailListRequest {
	s.ToPhoneNumList = v
	return s
}

type QueryRingDetailListShrinkRequest struct {
	PageSize               *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo                 *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	StartDate              *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	EndDate                *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	CallOutStatus          *string `json:"CallOutStatus,omitempty" xml:"CallOutStatus,omitempty"`
	Extra                  *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	FromPhoneNumListShrink *string `json:"FromPhoneNumList,omitempty" xml:"FromPhoneNumList,omitempty"`
	ToPhoneNumListShrink   *string `json:"ToPhoneNumList,omitempty" xml:"ToPhoneNumList,omitempty"`
}

func (s QueryRingDetailListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRingDetailListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryRingDetailListShrinkRequest) SetPageSize(v int32) *QueryRingDetailListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetPageNo(v int32) *QueryRingDetailListShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetStartDate(v int64) *QueryRingDetailListShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetEndDate(v int64) *QueryRingDetailListShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetCallOutStatus(v string) *QueryRingDetailListShrinkRequest {
	s.CallOutStatus = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetExtra(v string) *QueryRingDetailListShrinkRequest {
	s.Extra = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetInstanceId(v string) *QueryRingDetailListShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetFromPhoneNumListShrink(v string) *QueryRingDetailListShrinkRequest {
	s.FromPhoneNumListShrink = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetToPhoneNumListShrink(v string) *QueryRingDetailListShrinkRequest {
	s.ToPhoneNumListShrink = &v
	return s
}

type QueryRingDetailListResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s QueryRingDetailListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRingDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRingDetailListResponseBody) SetMessage(v string) *QueryRingDetailListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRingDetailListResponseBody) SetRequestId(v string) *QueryRingDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRingDetailListResponseBody) SetData(v string) *QueryRingDetailListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRingDetailListResponseBody) SetCode(v string) *QueryRingDetailListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRingDetailListResponseBody) SetSuccess(v bool) *QueryRingDetailListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRingDetailListResponseBody) SetHttpStatusCode(v int64) *QueryRingDetailListResponseBody {
	s.HttpStatusCode = &v
	return s
}

type QueryRingDetailListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRingDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRingDetailListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRingDetailListResponse) GoString() string {
	return s.String()
}

func (s *QueryRingDetailListResponse) SetHeaders(v map[string]*string) *QueryRingDetailListResponse {
	s.Headers = v
	return s
}

func (s *QueryRingDetailListResponse) SetBody(v *QueryRingDetailListResponseBody) *QueryRingDetailListResponse {
	s.Body = v
	return s
}

type SearchTicketByIdRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TicketId    *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	StatusCode  *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SearchTicketByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByIdRequest) GoString() string {
	return s.String()
}

func (s *SearchTicketByIdRequest) SetClientToken(v string) *SearchTicketByIdRequest {
	s.ClientToken = &v
	return s
}

func (s *SearchTicketByIdRequest) SetInstanceId(v string) *SearchTicketByIdRequest {
	s.InstanceId = &v
	return s
}

func (s *SearchTicketByIdRequest) SetTicketId(v int64) *SearchTicketByIdRequest {
	s.TicketId = &v
	return s
}

func (s *SearchTicketByIdRequest) SetStatusCode(v int32) *SearchTicketByIdRequest {
	s.StatusCode = &v
	return s
}

type SearchTicketByIdResponseBody struct {
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *SearchTicketByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s SearchTicketByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByIdResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTicketByIdResponseBody) SetMessage(v string) *SearchTicketByIdResponseBody {
	s.Message = &v
	return s
}

func (s *SearchTicketByIdResponseBody) SetRequestId(v string) *SearchTicketByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTicketByIdResponseBody) SetData(v *SearchTicketByIdResponseBodyData) *SearchTicketByIdResponseBody {
	s.Data = v
	return s
}

func (s *SearchTicketByIdResponseBody) SetCode(v string) *SearchTicketByIdResponseBody {
	s.Code = &v
	return s
}

func (s *SearchTicketByIdResponseBody) SetSuccess(v bool) *SearchTicketByIdResponseBody {
	s.Success = &v
	return s
}

func (s *SearchTicketByIdResponseBody) SetHttpStatusCode(v int64) *SearchTicketByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

type SearchTicketByIdResponseBodyData struct {
	CarbonCopy      *string                                            `json:"CarbonCopy,omitempty" xml:"CarbonCopy,omitempty"`
	MemberName      *string                                            `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	CreateTime      *int64                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ServiceId       *int64                                             `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	TicketId        *int64                                             `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	Priority        *int32                                             `json:"Priority,omitempty" xml:"Priority,omitempty"`
	CreatorId       *int64                                             `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	FormData        *string                                            `json:"FormData,omitempty" xml:"FormData,omitempty"`
	Activities      []*SearchTicketByIdResponseBodyDataActivities      `json:"Activities,omitempty" xml:"Activities,omitempty" type:"Repeated"`
	ActivityRecords []*SearchTicketByIdResponseBodyDataActivityRecords `json:"ActivityRecords,omitempty" xml:"ActivityRecords,omitempty" type:"Repeated"`
	FromInfo        *string                                            `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	ModifiedTime    *int64                                             `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	CreatorName     *string                                            `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	CategoryId      *int64                                             `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreatorType     *int32                                             `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	MemberId        *int64                                             `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	CaseStatus      *int32                                             `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	TemplateId      *int64                                             `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TicketName      *string                                            `json:"TicketName,omitempty" xml:"TicketName,omitempty"`
}

func (s SearchTicketByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchTicketByIdResponseBodyData) SetCarbonCopy(v string) *SearchTicketByIdResponseBodyData {
	s.CarbonCopy = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetMemberName(v string) *SearchTicketByIdResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCreateTime(v int64) *SearchTicketByIdResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetServiceId(v int64) *SearchTicketByIdResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetTicketId(v int64) *SearchTicketByIdResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetPriority(v int32) *SearchTicketByIdResponseBodyData {
	s.Priority = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCreatorId(v int64) *SearchTicketByIdResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetFormData(v string) *SearchTicketByIdResponseBodyData {
	s.FormData = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetActivities(v []*SearchTicketByIdResponseBodyDataActivities) *SearchTicketByIdResponseBodyData {
	s.Activities = v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetActivityRecords(v []*SearchTicketByIdResponseBodyDataActivityRecords) *SearchTicketByIdResponseBodyData {
	s.ActivityRecords = v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetFromInfo(v string) *SearchTicketByIdResponseBodyData {
	s.FromInfo = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetModifiedTime(v int64) *SearchTicketByIdResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCreatorName(v string) *SearchTicketByIdResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCategoryId(v int64) *SearchTicketByIdResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCreatorType(v int32) *SearchTicketByIdResponseBodyData {
	s.CreatorType = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetMemberId(v int64) *SearchTicketByIdResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCaseStatus(v int32) *SearchTicketByIdResponseBodyData {
	s.CaseStatus = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetTemplateId(v int64) *SearchTicketByIdResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetTicketName(v string) *SearchTicketByIdResponseBodyData {
	s.TicketName = &v
	return s
}

type SearchTicketByIdResponseBodyDataActivities struct {
	ActivityFormData *string `json:"ActivityFormData,omitempty" xml:"ActivityFormData,omitempty"`
	ActivityCode     *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
}

func (s SearchTicketByIdResponseBodyDataActivities) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByIdResponseBodyDataActivities) GoString() string {
	return s.String()
}

func (s *SearchTicketByIdResponseBodyDataActivities) SetActivityFormData(v string) *SearchTicketByIdResponseBodyDataActivities {
	s.ActivityFormData = &v
	return s
}

func (s *SearchTicketByIdResponseBodyDataActivities) SetActivityCode(v string) *SearchTicketByIdResponseBodyDataActivities {
	s.ActivityCode = &v
	return s
}

type SearchTicketByIdResponseBodyDataActivityRecords struct {
	ActionCode     *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	ActionCodeDesc *string `json:"ActionCodeDesc,omitempty" xml:"ActionCodeDesc,omitempty"`
	OperatorName   *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	Memo           *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
}

func (s SearchTicketByIdResponseBodyDataActivityRecords) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByIdResponseBodyDataActivityRecords) GoString() string {
	return s.String()
}

func (s *SearchTicketByIdResponseBodyDataActivityRecords) SetActionCode(v string) *SearchTicketByIdResponseBodyDataActivityRecords {
	s.ActionCode = &v
	return s
}

func (s *SearchTicketByIdResponseBodyDataActivityRecords) SetActionCodeDesc(v string) *SearchTicketByIdResponseBodyDataActivityRecords {
	s.ActionCodeDesc = &v
	return s
}

func (s *SearchTicketByIdResponseBodyDataActivityRecords) SetOperatorName(v string) *SearchTicketByIdResponseBodyDataActivityRecords {
	s.OperatorName = &v
	return s
}

func (s *SearchTicketByIdResponseBodyDataActivityRecords) SetMemo(v string) *SearchTicketByIdResponseBodyDataActivityRecords {
	s.Memo = &v
	return s
}

func (s *SearchTicketByIdResponseBodyDataActivityRecords) SetGmtCreate(v int64) *SearchTicketByIdResponseBodyDataActivityRecords {
	s.GmtCreate = &v
	return s
}

type SearchTicketByIdResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchTicketByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTicketByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByIdResponse) GoString() string {
	return s.String()
}

func (s *SearchTicketByIdResponse) SetHeaders(v map[string]*string) *SearchTicketByIdResponse {
	s.Headers = v
	return s
}

func (s *SearchTicketByIdResponse) SetBody(v *SearchTicketByIdResponseBody) *SearchTicketByIdResponse {
	s.Body = v
	return s
}

type UpdateSkillGroupRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId   *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 渠道类型
	ChannelType *int64 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
}

func (s UpdateSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupRequest) SetInstanceId(v string) *UpdateSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetSkillGroupId(v int64) *UpdateSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetSkillGroupName(v string) *UpdateSkillGroupRequest {
	s.SkillGroupName = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetDescription(v string) *UpdateSkillGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetDisplayName(v string) *UpdateSkillGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetClientToken(v string) *UpdateSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetChannelType(v int64) *UpdateSkillGroupRequest {
	s.ChannelType = &v
	return s
}

type UpdateSkillGroupResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s UpdateSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupResponseBody) SetMessage(v string) *UpdateSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetRequestId(v string) *UpdateSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetCode(v string) *UpdateSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetSuccess(v bool) *UpdateSkillGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetHttpStatusCode(v int64) *UpdateSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

type UpdateSkillGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupResponse) SetHeaders(v map[string]*string) *UpdateSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateSkillGroupResponse) SetBody(v *UpdateSkillGroupResponseBody) *UpdateSkillGroupResponse {
	s.Body = v
	return s
}

type QueryServiceConfigRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ViewCode   *string `json:"ViewCode,omitempty" xml:"ViewCode,omitempty"`
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s QueryServiceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryServiceConfigRequest) SetInstanceId(v string) *QueryServiceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryServiceConfigRequest) SetViewCode(v string) *QueryServiceConfigRequest {
	s.ViewCode = &v
	return s
}

func (s *QueryServiceConfigRequest) SetParameters(v string) *QueryServiceConfigRequest {
	s.Parameters = &v
	return s
}

type QueryServiceConfigResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryServiceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServiceConfigResponseBody) SetMessage(v string) *QueryServiceConfigResponseBody {
	s.Message = &v
	return s
}

func (s *QueryServiceConfigResponseBody) SetRequestId(v string) *QueryServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryServiceConfigResponseBody) SetData(v string) *QueryServiceConfigResponseBody {
	s.Data = &v
	return s
}

func (s *QueryServiceConfigResponseBody) SetCode(v string) *QueryServiceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *QueryServiceConfigResponseBody) SetSuccess(v bool) *QueryServiceConfigResponseBody {
	s.Success = &v
	return s
}

type QueryServiceConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServiceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryServiceConfigResponse) SetHeaders(v map[string]*string) *QueryServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryServiceConfigResponse) SetBody(v *QueryServiceConfigResponseBody) *QueryServiceConfigResponse {
	s.Body = v
	return s
}

type DisableRoleRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleId      *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s DisableRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableRoleRequest) GoString() string {
	return s.String()
}

func (s *DisableRoleRequest) SetClientToken(v string) *DisableRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableRoleRequest) SetInstanceId(v string) *DisableRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableRoleRequest) SetRoleId(v int64) *DisableRoleRequest {
	s.RoleId = &v
	return s
}

type DisableRoleResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableRoleResponseBody) SetMessage(v string) *DisableRoleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableRoleResponseBody) SetRequestId(v string) *DisableRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableRoleResponseBody) SetCode(v string) *DisableRoleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableRoleResponseBody) SetSuccess(v bool) *DisableRoleResponseBody {
	s.Success = &v
	return s
}

type DisableRoleResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableRoleResponse) GoString() string {
	return s.String()
}

func (s *DisableRoleResponse) SetHeaders(v map[string]*string) *DisableRoleResponse {
	s.Headers = v
	return s
}

func (s *DisableRoleResponse) SetBody(v *DisableRoleResponseBody) *DisableRoleResponse {
	s.Body = v
	return s
}

type GetEntityRouteListRequest struct {
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
	EntityName           *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
}

func (s GetEntityRouteListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteListRequest) GoString() string {
	return s.String()
}

func (s *GetEntityRouteListRequest) SetPageSize(v int32) *GetEntityRouteListRequest {
	s.PageSize = &v
	return s
}

func (s *GetEntityRouteListRequest) SetPageNo(v int32) *GetEntityRouteListRequest {
	s.PageNo = &v
	return s
}

func (s *GetEntityRouteListRequest) SetInstanceId(v string) *GetEntityRouteListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetEntityRouteListRequest) SetEntityRelationNumber(v string) *GetEntityRouteListRequest {
	s.EntityRelationNumber = &v
	return s
}

func (s *GetEntityRouteListRequest) SetEntityName(v string) *GetEntityRouteListRequest {
	s.EntityName = &v
	return s
}

type GetEntityRouteListResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetEntityRouteListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEntityRouteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteListResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityRouteListResponseBody) SetMessage(v string) *GetEntityRouteListResponseBody {
	s.Message = &v
	return s
}

func (s *GetEntityRouteListResponseBody) SetRequestId(v string) *GetEntityRouteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntityRouteListResponseBody) SetData(v *GetEntityRouteListResponseBodyData) *GetEntityRouteListResponseBody {
	s.Data = v
	return s
}

func (s *GetEntityRouteListResponseBody) SetCode(v string) *GetEntityRouteListResponseBody {
	s.Code = &v
	return s
}

func (s *GetEntityRouteListResponseBody) SetSuccess(v bool) *GetEntityRouteListResponseBody {
	s.Success = &v
	return s
}

type GetEntityRouteListResponseBodyData struct {
	EntityRouteList []*GetEntityRouteListResponseBodyDataEntityRouteList `json:"EntityRouteList,omitempty" xml:"EntityRouteList,omitempty" type:"Repeated"`
	PageNo          *int32                                               `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total           *int64                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetEntityRouteListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEntityRouteListResponseBodyData) SetEntityRouteList(v []*GetEntityRouteListResponseBodyDataEntityRouteList) *GetEntityRouteListResponseBodyData {
	s.EntityRouteList = v
	return s
}

func (s *GetEntityRouteListResponseBodyData) SetPageNo(v int32) *GetEntityRouteListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetEntityRouteListResponseBodyData) SetPageSize(v int32) *GetEntityRouteListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetEntityRouteListResponseBodyData) SetTotal(v int64) *GetEntityRouteListResponseBodyData {
	s.Total = &v
	return s
}

type GetEntityRouteListResponseBodyDataEntityRouteList struct {
	EntityBizCodeType    *string `json:"EntityBizCodeType,omitempty" xml:"EntityBizCodeType,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	EntityId             *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	ServiceId            *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	DepartmentId         *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	EntityBizCode        *string `json:"EntityBizCode,omitempty" xml:"EntityBizCode,omitempty"`
	UniqueId             *int64  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	EntityName           *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	ExtInfo              *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
}

func (s GetEntityRouteListResponseBodyDataEntityRouteList) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteListResponseBodyDataEntityRouteList) GoString() string {
	return s.String()
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetEntityBizCodeType(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.EntityBizCodeType = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetGroupId(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.GroupId = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetEntityId(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.EntityId = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetServiceId(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.ServiceId = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetDepartmentId(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.DepartmentId = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetEntityBizCode(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.EntityBizCode = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetUniqueId(v int64) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.UniqueId = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetEntityName(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.EntityName = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetExtInfo(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.ExtInfo = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetEntityRelationNumber(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.EntityRelationNumber = &v
	return s
}

type GetEntityRouteListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEntityRouteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEntityRouteListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteListResponse) GoString() string {
	return s.String()
}

func (s *GetEntityRouteListResponse) SetHeaders(v map[string]*string) *GetEntityRouteListResponse {
	s.Headers = v
	return s
}

func (s *GetEntityRouteListResponse) SetBody(v *GetEntityRouteListResponseBody) *GetEntityRouteListResponse {
	s.Body = v
	return s
}

type GetAuthInfoRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ForeignId  *string `json:"ForeignId,omitempty" xml:"ForeignId,omitempty"`
	AppKey     *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s GetAuthInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAuthInfoRequest) SetInstanceId(v string) *GetAuthInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAuthInfoRequest) SetForeignId(v string) *GetAuthInfoRequest {
	s.ForeignId = &v
	return s
}

func (s *GetAuthInfoRequest) SetAppKey(v string) *GetAuthInfoRequest {
	s.AppKey = &v
	return s
}

type GetAuthInfoResponseBody struct {
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetAuthInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAuthInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthInfoResponseBody) SetMessage(v string) *GetAuthInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuthInfoResponseBody) SetRequestId(v string) *GetAuthInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthInfoResponseBody) SetData(v *GetAuthInfoResponseBodyData) *GetAuthInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetAuthInfoResponseBody) SetCode(v string) *GetAuthInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthInfoResponseBody) SetSuccess(v bool) *GetAuthInfoResponseBody {
	s.Success = &v
	return s
}

type GetAuthInfoResponseBodyData struct {
	AppName   *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Time      *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	App       *string `json:"App,omitempty" xml:"App,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	TenantId  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetAuthInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAuthInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAuthInfoResponseBodyData) SetAppName(v string) *GetAuthInfoResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetTime(v int64) *GetAuthInfoResponseBodyData {
	s.Time = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetAppKey(v string) *GetAuthInfoResponseBodyData {
	s.AppKey = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetApp(v string) *GetAuthInfoResponseBodyData {
	s.App = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetUserId(v string) *GetAuthInfoResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetCode(v string) *GetAuthInfoResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetSessionId(v string) *GetAuthInfoResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetUserName(v string) *GetAuthInfoResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetTenantId(v string) *GetAuthInfoResponseBodyData {
	s.TenantId = &v
	return s
}

type GetAuthInfoResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAuthInfoResponse) SetHeaders(v map[string]*string) *GetAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAuthInfoResponse) SetBody(v *GetAuthInfoResponseBody) *GetAuthInfoResponse {
	s.Body = v
	return s
}

type UpdateRoleRequest struct {
	ClientToken  *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleId       *int64   `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName     *string  `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	Operator     *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	PermissionId []*int64 `json:"PermissionId,omitempty" xml:"PermissionId,omitempty" type:"Repeated"`
}

func (s UpdateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleRequest) SetClientToken(v string) *UpdateRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRoleRequest) SetInstanceId(v string) *UpdateRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRoleRequest) SetRoleId(v int64) *UpdateRoleRequest {
	s.RoleId = &v
	return s
}

func (s *UpdateRoleRequest) SetRoleName(v string) *UpdateRoleRequest {
	s.RoleName = &v
	return s
}

func (s *UpdateRoleRequest) SetOperator(v string) *UpdateRoleRequest {
	s.Operator = &v
	return s
}

func (s *UpdateRoleRequest) SetPermissionId(v []*int64) *UpdateRoleRequest {
	s.PermissionId = v
	return s
}

type UpdateRoleResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s UpdateRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponseBody) SetMessage(v string) *UpdateRoleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRoleResponseBody) SetRequestId(v string) *UpdateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRoleResponseBody) SetCode(v string) *UpdateRoleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRoleResponseBody) SetSuccess(v bool) *UpdateRoleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRoleResponseBody) SetHttpStatusCode(v int64) *UpdateRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

type UpdateRoleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponse) SetHeaders(v map[string]*string) *UpdateRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoleResponse) SetBody(v *UpdateRoleResponseBody) *UpdateRoleResponse {
	s.Body = v
	return s
}

type GetTicketTemplateSchemaRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TemplateId  *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTicketTemplateSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTicketTemplateSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateSchemaRequest) SetClientToken(v string) *GetTicketTemplateSchemaRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTicketTemplateSchemaRequest) SetInstanceId(v string) *GetTicketTemplateSchemaRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTicketTemplateSchemaRequest) SetTemplateId(v int64) *GetTicketTemplateSchemaRequest {
	s.TemplateId = &v
	return s
}

type GetTicketTemplateSchemaResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTicketTemplateSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTicketTemplateSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateSchemaResponseBody) SetMessage(v string) *GetTicketTemplateSchemaResponseBody {
	s.Message = &v
	return s
}

func (s *GetTicketTemplateSchemaResponseBody) SetRequestId(v string) *GetTicketTemplateSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTicketTemplateSchemaResponseBody) SetData(v string) *GetTicketTemplateSchemaResponseBody {
	s.Data = &v
	return s
}

func (s *GetTicketTemplateSchemaResponseBody) SetCode(v string) *GetTicketTemplateSchemaResponseBody {
	s.Code = &v
	return s
}

func (s *GetTicketTemplateSchemaResponseBody) SetSuccess(v bool) *GetTicketTemplateSchemaResponseBody {
	s.Success = &v
	return s
}

type GetTicketTemplateSchemaResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTicketTemplateSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTicketTemplateSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTicketTemplateSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateSchemaResponse) SetHeaders(v map[string]*string) *GetTicketTemplateSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetTicketTemplateSchemaResponse) SetBody(v *GetTicketTemplateSchemaResponseBody) *GetTicketTemplateSchemaResponse {
	s.Body = v
	return s
}

type TransferCallToPhoneRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Caller           *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Callee           *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	IsSingleTransfer *bool   `json:"IsSingleTransfer,omitempty" xml:"IsSingleTransfer,omitempty"`
	CallerPhone      *string `json:"callerPhone,omitempty" xml:"callerPhone,omitempty"`
	CalleePhone      *string `json:"calleePhone,omitempty" xml:"calleePhone,omitempty"`
}

func (s TransferCallToPhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToPhoneRequest) GoString() string {
	return s.String()
}

func (s *TransferCallToPhoneRequest) SetClientToken(v string) *TransferCallToPhoneRequest {
	s.ClientToken = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetInstanceId(v string) *TransferCallToPhoneRequest {
	s.InstanceId = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetAccountName(v string) *TransferCallToPhoneRequest {
	s.AccountName = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetCaller(v string) *TransferCallToPhoneRequest {
	s.Caller = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetCallee(v string) *TransferCallToPhoneRequest {
	s.Callee = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetCallId(v string) *TransferCallToPhoneRequest {
	s.CallId = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetJobId(v string) *TransferCallToPhoneRequest {
	s.JobId = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetConnectionId(v string) *TransferCallToPhoneRequest {
	s.ConnectionId = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetHoldConnectionId(v string) *TransferCallToPhoneRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetType(v int32) *TransferCallToPhoneRequest {
	s.Type = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetIsSingleTransfer(v bool) *TransferCallToPhoneRequest {
	s.IsSingleTransfer = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetCallerPhone(v string) *TransferCallToPhoneRequest {
	s.CallerPhone = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetCalleePhone(v string) *TransferCallToPhoneRequest {
	s.CalleePhone = &v
	return s
}

type TransferCallToPhoneResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s TransferCallToPhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *TransferCallToPhoneResponseBody) SetMessage(v string) *TransferCallToPhoneResponseBody {
	s.Message = &v
	return s
}

func (s *TransferCallToPhoneResponseBody) SetRequestId(v string) *TransferCallToPhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferCallToPhoneResponseBody) SetCode(v string) *TransferCallToPhoneResponseBody {
	s.Code = &v
	return s
}

func (s *TransferCallToPhoneResponseBody) SetSuccess(v bool) *TransferCallToPhoneResponseBody {
	s.Success = &v
	return s
}

func (s *TransferCallToPhoneResponseBody) SetHttpStatusCode(v int64) *TransferCallToPhoneResponseBody {
	s.HttpStatusCode = &v
	return s
}

type TransferCallToPhoneResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferCallToPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferCallToPhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToPhoneResponse) GoString() string {
	return s.String()
}

func (s *TransferCallToPhoneResponse) SetHeaders(v map[string]*string) *TransferCallToPhoneResponse {
	s.Headers = v
	return s
}

func (s *TransferCallToPhoneResponse) SetBody(v *TransferCallToPhoneResponseBody) *TransferCallToPhoneResponse {
	s.Body = v
	return s
}

type QuerySkillGroupsRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType   *int32  `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	GroupId     *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s QuerySkillGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsRequest) SetInstanceId(v string) *QuerySkillGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetPageNo(v int32) *QuerySkillGroupsRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetPageSize(v int32) *QuerySkillGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetClientToken(v string) *QuerySkillGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetGroupName(v string) *QuerySkillGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetGroupType(v int32) *QuerySkillGroupsRequest {
	s.GroupType = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetGroupId(v int64) *QuerySkillGroupsRequest {
	s.GroupId = &v
	return s
}

type QuerySkillGroupsResponseBody struct {
	OnePageSize  *int32                              `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	TotalPage    *int32                              `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage  *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	TotalResults *int32                              `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
	Data         []*QuerySkillGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s QuerySkillGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySkillGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsResponseBody) SetOnePageSize(v int32) *QuerySkillGroupsResponseBody {
	s.OnePageSize = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetTotalPage(v int32) *QuerySkillGroupsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetRequestId(v string) *QuerySkillGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetCurrentPage(v int32) *QuerySkillGroupsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetTotalResults(v int32) *QuerySkillGroupsResponseBody {
	s.TotalResults = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetData(v []*QuerySkillGroupsResponseBodyData) *QuerySkillGroupsResponseBody {
	s.Data = v
	return s
}

type QuerySkillGroupsResponseBodyData struct {
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ChannelType    *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	SkillGroupId   *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s QuerySkillGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySkillGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsResponseBodyData) SetDisplayName(v string) *QuerySkillGroupsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetDescription(v string) *QuerySkillGroupsResponseBodyData {
	s.Description = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetChannelType(v int32) *QuerySkillGroupsResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetSkillGroupName(v string) *QuerySkillGroupsResponseBodyData {
	s.SkillGroupName = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetSkillGroupId(v int64) *QuerySkillGroupsResponseBodyData {
	s.SkillGroupId = &v
	return s
}

type QuerySkillGroupsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QuerySkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySkillGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsResponse) SetHeaders(v map[string]*string) *QuerySkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *QuerySkillGroupsResponse) SetBody(v *QuerySkillGroupsResponseBody) *QuerySkillGroupsResponse {
	s.Body = v
	return s
}

type GetEntityRouteRequest struct {
	EntityId             *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	UniqueId             *int64  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	EntityName           *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
	EntityBizCode        *string `json:"EntityBizCode,omitempty" xml:"EntityBizCode,omitempty"`
}

func (s GetEntityRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteRequest) GoString() string {
	return s.String()
}

func (s *GetEntityRouteRequest) SetEntityId(v string) *GetEntityRouteRequest {
	s.EntityId = &v
	return s
}

func (s *GetEntityRouteRequest) SetUniqueId(v int64) *GetEntityRouteRequest {
	s.UniqueId = &v
	return s
}

func (s *GetEntityRouteRequest) SetInstanceId(v string) *GetEntityRouteRequest {
	s.InstanceId = &v
	return s
}

func (s *GetEntityRouteRequest) SetEntityName(v string) *GetEntityRouteRequest {
	s.EntityName = &v
	return s
}

func (s *GetEntityRouteRequest) SetEntityRelationNumber(v string) *GetEntityRouteRequest {
	s.EntityRelationNumber = &v
	return s
}

func (s *GetEntityRouteRequest) SetEntityBizCode(v string) *GetEntityRouteRequest {
	s.EntityBizCode = &v
	return s
}

type GetEntityRouteResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetEntityRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEntityRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityRouteResponseBody) SetMessage(v string) *GetEntityRouteResponseBody {
	s.Message = &v
	return s
}

func (s *GetEntityRouteResponseBody) SetRequestId(v string) *GetEntityRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntityRouteResponseBody) SetData(v *GetEntityRouteResponseBodyData) *GetEntityRouteResponseBody {
	s.Data = v
	return s
}

func (s *GetEntityRouteResponseBody) SetCode(v string) *GetEntityRouteResponseBody {
	s.Code = &v
	return s
}

func (s *GetEntityRouteResponseBody) SetSuccess(v bool) *GetEntityRouteResponseBody {
	s.Success = &v
	return s
}

type GetEntityRouteResponseBodyData struct {
	EntityBizCodeType    *string `json:"EntityBizCodeType,omitempty" xml:"EntityBizCodeType,omitempty"`
	GroupId              *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	EntityId             *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	ServiceId            *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	EntityBizCode        *string `json:"EntityBizCode,omitempty" xml:"EntityBizCode,omitempty"`
	DepartmentId         *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	UniqueId             *int64  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	EntityName           *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	ExtInfo              *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
}

func (s GetEntityRouteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEntityRouteResponseBodyData) SetEntityBizCodeType(v string) *GetEntityRouteResponseBodyData {
	s.EntityBizCodeType = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetGroupId(v int64) *GetEntityRouteResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetEntityId(v string) *GetEntityRouteResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetServiceId(v int64) *GetEntityRouteResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetEntityBizCode(v string) *GetEntityRouteResponseBodyData {
	s.EntityBizCode = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetDepartmentId(v string) *GetEntityRouteResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetUniqueId(v int64) *GetEntityRouteResponseBodyData {
	s.UniqueId = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetEntityName(v string) *GetEntityRouteResponseBodyData {
	s.EntityName = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetExtInfo(v string) *GetEntityRouteResponseBodyData {
	s.ExtInfo = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetEntityRelationNumber(v string) *GetEntityRouteResponseBodyData {
	s.EntityRelationNumber = &v
	return s
}

type GetEntityRouteResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEntityRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEntityRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteResponse) GoString() string {
	return s.String()
}

func (s *GetEntityRouteResponse) SetHeaders(v map[string]*string) *GetEntityRouteResponse {
	s.Headers = v
	return s
}

func (s *GetEntityRouteResponse) SetBody(v *GetEntityRouteResponseBody) *GetEntityRouteResponse {
	s.Body = v
	return s
}

type UpdateEntityTagRelationRequest struct {
	EntityTagParam *string `json:"EntityTagParam,omitempty" xml:"EntityTagParam,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateEntityTagRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityTagRelationRequest) GoString() string {
	return s.String()
}

func (s *UpdateEntityTagRelationRequest) SetEntityTagParam(v string) *UpdateEntityTagRelationRequest {
	s.EntityTagParam = &v
	return s
}

func (s *UpdateEntityTagRelationRequest) SetInstanceId(v string) *UpdateEntityTagRelationRequest {
	s.InstanceId = &v
	return s
}

type UpdateEntityTagRelationResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEntityTagRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityTagRelationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEntityTagRelationResponseBody) SetMessage(v string) *UpdateEntityTagRelationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEntityTagRelationResponseBody) SetRequestId(v string) *UpdateEntityTagRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEntityTagRelationResponseBody) SetData(v string) *UpdateEntityTagRelationResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateEntityTagRelationResponseBody) SetCode(v string) *UpdateEntityTagRelationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEntityTagRelationResponseBody) SetSuccess(v bool) *UpdateEntityTagRelationResponseBody {
	s.Success = &v
	return s
}

type UpdateEntityTagRelationResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEntityTagRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEntityTagRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityTagRelationResponse) GoString() string {
	return s.String()
}

func (s *UpdateEntityTagRelationResponse) SetHeaders(v map[string]*string) *UpdateEntityTagRelationResponse {
	s.Headers = v
	return s
}

func (s *UpdateEntityTagRelationResponse) SetBody(v *UpdateEntityTagRelationResponseBody) *UpdateEntityTagRelationResponse {
	s.Body = v
	return s
}

type CreateOuterCallCenterDataRequest struct {
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	InterveneTime *string `json:"InterveneTime,omitempty" xml:"InterveneTime,omitempty"`
	CallType      *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	FromPhoneNum  *string `json:"FromPhoneNum,omitempty" xml:"FromPhoneNum,omitempty"`
	ToPhoneNum    *string `json:"ToPhoneNum,omitempty" xml:"ToPhoneNum,omitempty"`
	EndReason     *string `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	UserInfo      *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
	RecordUrl     *string `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	BizId         *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ExtInfo       *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateOuterCallCenterDataRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOuterCallCenterDataRequest) GoString() string {
	return s.String()
}

func (s *CreateOuterCallCenterDataRequest) SetSessionId(v string) *CreateOuterCallCenterDataRequest {
	s.SessionId = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetInterveneTime(v string) *CreateOuterCallCenterDataRequest {
	s.InterveneTime = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetCallType(v string) *CreateOuterCallCenterDataRequest {
	s.CallType = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetFromPhoneNum(v string) *CreateOuterCallCenterDataRequest {
	s.FromPhoneNum = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetToPhoneNum(v string) *CreateOuterCallCenterDataRequest {
	s.ToPhoneNum = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetEndReason(v string) *CreateOuterCallCenterDataRequest {
	s.EndReason = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetUserInfo(v string) *CreateOuterCallCenterDataRequest {
	s.UserInfo = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetRecordUrl(v string) *CreateOuterCallCenterDataRequest {
	s.RecordUrl = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetBizType(v string) *CreateOuterCallCenterDataRequest {
	s.BizType = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetBizId(v string) *CreateOuterCallCenterDataRequest {
	s.BizId = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetTenantId(v string) *CreateOuterCallCenterDataRequest {
	s.TenantId = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetExtInfo(v string) *CreateOuterCallCenterDataRequest {
	s.ExtInfo = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetInstanceId(v string) *CreateOuterCallCenterDataRequest {
	s.InstanceId = &v
	return s
}

type CreateOuterCallCenterDataResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s CreateOuterCallCenterDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOuterCallCenterDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOuterCallCenterDataResponseBody) SetMessage(v string) *CreateOuterCallCenterDataResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOuterCallCenterDataResponseBody) SetRequestId(v string) *CreateOuterCallCenterDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOuterCallCenterDataResponseBody) SetCode(v string) *CreateOuterCallCenterDataResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOuterCallCenterDataResponseBody) SetSuccess(v bool) *CreateOuterCallCenterDataResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOuterCallCenterDataResponseBody) SetHttpStatusCode(v int64) *CreateOuterCallCenterDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

type CreateOuterCallCenterDataResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateOuterCallCenterDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOuterCallCenterDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOuterCallCenterDataResponse) GoString() string {
	return s.String()
}

func (s *CreateOuterCallCenterDataResponse) SetHeaders(v map[string]*string) *CreateOuterCallCenterDataResponse {
	s.Headers = v
	return s
}

func (s *CreateOuterCallCenterDataResponse) SetBody(v *CreateOuterCallCenterDataResponseBody) *CreateOuterCallCenterDataResponse {
	s.Body = v
	return s
}

type SendOutboundCommandRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName   *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	CalledNumber  *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CustomerInfo  *string `json:"CustomerInfo,omitempty" xml:"CustomerInfo,omitempty"`
}

func (s SendOutboundCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s SendOutboundCommandRequest) GoString() string {
	return s.String()
}

func (s *SendOutboundCommandRequest) SetInstanceId(v string) *SendOutboundCommandRequest {
	s.InstanceId = &v
	return s
}

func (s *SendOutboundCommandRequest) SetAccountName(v string) *SendOutboundCommandRequest {
	s.AccountName = &v
	return s
}

func (s *SendOutboundCommandRequest) SetCallingNumber(v string) *SendOutboundCommandRequest {
	s.CallingNumber = &v
	return s
}

func (s *SendOutboundCommandRequest) SetCalledNumber(v string) *SendOutboundCommandRequest {
	s.CalledNumber = &v
	return s
}

func (s *SendOutboundCommandRequest) SetCustomerInfo(v string) *SendOutboundCommandRequest {
	s.CustomerInfo = &v
	return s
}

type SendOutboundCommandResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendOutboundCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendOutboundCommandResponseBody) GoString() string {
	return s.String()
}

func (s *SendOutboundCommandResponseBody) SetMessage(v string) *SendOutboundCommandResponseBody {
	s.Message = &v
	return s
}

func (s *SendOutboundCommandResponseBody) SetRequestId(v string) *SendOutboundCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendOutboundCommandResponseBody) SetData(v string) *SendOutboundCommandResponseBody {
	s.Data = &v
	return s
}

func (s *SendOutboundCommandResponseBody) SetCode(v string) *SendOutboundCommandResponseBody {
	s.Code = &v
	return s
}

func (s *SendOutboundCommandResponseBody) SetSuccess(v bool) *SendOutboundCommandResponseBody {
	s.Success = &v
	return s
}

type SendOutboundCommandResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendOutboundCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendOutboundCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s SendOutboundCommandResponse) GoString() string {
	return s.String()
}

func (s *SendOutboundCommandResponse) SetHeaders(v map[string]*string) *SendOutboundCommandResponse {
	s.Headers = v
	return s
}

func (s *SendOutboundCommandResponse) SetBody(v *SendOutboundCommandResponseBody) *SendOutboundCommandResponse {
	s.Body = v
	return s
}

type CreateRoleRequest struct {
	ClientToken  *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleName     *string  `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	Operator     *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	PermissionId []*int64 `json:"PermissionId,omitempty" xml:"PermissionId,omitempty" type:"Repeated"`
}

func (s CreateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) SetClientToken(v string) *CreateRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRoleRequest) SetInstanceId(v string) *CreateRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRoleRequest) SetRoleName(v string) *CreateRoleRequest {
	s.RoleName = &v
	return s
}

func (s *CreateRoleRequest) SetOperator(v string) *CreateRoleRequest {
	s.Operator = &v
	return s
}

func (s *CreateRoleRequest) SetPermissionId(v []*int64) *CreateRoleRequest {
	s.PermissionId = v
	return s
}

type CreateRoleResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s CreateRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseBody) SetMessage(v string) *CreateRoleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRoleResponseBody) SetRequestId(v string) *CreateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoleResponseBody) SetData(v int64) *CreateRoleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRoleResponseBody) SetCode(v string) *CreateRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRoleResponseBody) SetSuccess(v bool) *CreateRoleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRoleResponseBody) SetHttpStatusCode(v int64) *CreateRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

type CreateRoleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateRoleResponse) SetHeaders(v map[string]*string) *CreateRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateRoleResponse) SetBody(v *CreateRoleResponseBody) *CreateRoleResponse {
	s.Body = v
	return s
}

type ListSkillGroupRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ChannelType *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
}

func (s ListSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *ListSkillGroupRequest) SetClientToken(v string) *ListSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ListSkillGroupRequest) SetInstanceId(v string) *ListSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupRequest) SetChannelType(v int32) *ListSkillGroupRequest {
	s.ChannelType = &v
	return s
}

type ListSkillGroupResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*ListSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillGroupResponseBody) SetMessage(v string) *ListSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetRequestId(v string) *ListSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetData(v []*ListSkillGroupResponseBodyData) *ListSkillGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillGroupResponseBody) SetCode(v string) *ListSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetSuccess(v bool) *ListSkillGroupResponseBody {
	s.Success = &v
	return s
}

type ListSkillGroupResponseBodyData struct {
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ChannelType  *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	SkillGroupId *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListSkillGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillGroupResponseBodyData) SetDisplayName(v string) *ListSkillGroupResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetDescription(v string) *ListSkillGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetChannelType(v int32) *ListSkillGroupResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetSkillGroupId(v int64) *ListSkillGroupResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetName(v string) *ListSkillGroupResponseBodyData {
	s.Name = &v
	return s
}

type ListSkillGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *ListSkillGroupResponse) SetHeaders(v map[string]*string) *ListSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *ListSkillGroupResponse) SetBody(v *ListSkillGroupResponseBody) *ListSkillGroupResponse {
	s.Body = v
	return s
}

type GrantRolesRequest struct {
	ClientToken *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Operator    *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	RoleId      []*int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty" type:"Repeated"`
}

func (s GrantRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantRolesRequest) GoString() string {
	return s.String()
}

func (s *GrantRolesRequest) SetClientToken(v string) *GrantRolesRequest {
	s.ClientToken = &v
	return s
}

func (s *GrantRolesRequest) SetInstanceId(v string) *GrantRolesRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantRolesRequest) SetAccountName(v string) *GrantRolesRequest {
	s.AccountName = &v
	return s
}

func (s *GrantRolesRequest) SetOperator(v string) *GrantRolesRequest {
	s.Operator = &v
	return s
}

func (s *GrantRolesRequest) SetRoleId(v []*int64) *GrantRolesRequest {
	s.RoleId = v
	return s
}

type GrantRolesResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GrantRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantRolesResponseBody) GoString() string {
	return s.String()
}

func (s *GrantRolesResponseBody) SetMessage(v string) *GrantRolesResponseBody {
	s.Message = &v
	return s
}

func (s *GrantRolesResponseBody) SetRequestId(v string) *GrantRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantRolesResponseBody) SetCode(v string) *GrantRolesResponseBody {
	s.Code = &v
	return s
}

func (s *GrantRolesResponseBody) SetSuccess(v bool) *GrantRolesResponseBody {
	s.Success = &v
	return s
}

func (s *GrantRolesResponseBody) SetHttpStatusCode(v int64) *GrantRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GrantRolesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GrantRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantRolesResponse) GoString() string {
	return s.String()
}

func (s *GrantRolesResponse) SetHeaders(v map[string]*string) *GrantRolesResponse {
	s.Headers = v
	return s
}

func (s *GrantRolesResponse) SetBody(v *GrantRolesResponseBody) *GrantRolesResponse {
	s.Body = v
	return s
}

type GetOuterCallCenterDataListRequest struct {
	SessionId      *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	FromPhoneNum   *string `json:"FromPhoneNum,omitempty" xml:"FromPhoneNum,omitempty"`
	ToPhoneNum     *string `json:"ToPhoneNum,omitempty" xml:"ToPhoneNum,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	QueryEndTime   *string `json:"QueryEndTime,omitempty" xml:"QueryEndTime,omitempty"`
}

func (s GetOuterCallCenterDataListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOuterCallCenterDataListRequest) GoString() string {
	return s.String()
}

func (s *GetOuterCallCenterDataListRequest) SetSessionId(v string) *GetOuterCallCenterDataListRequest {
	s.SessionId = &v
	return s
}

func (s *GetOuterCallCenterDataListRequest) SetFromPhoneNum(v string) *GetOuterCallCenterDataListRequest {
	s.FromPhoneNum = &v
	return s
}

func (s *GetOuterCallCenterDataListRequest) SetToPhoneNum(v string) *GetOuterCallCenterDataListRequest {
	s.ToPhoneNum = &v
	return s
}

func (s *GetOuterCallCenterDataListRequest) SetBizId(v string) *GetOuterCallCenterDataListRequest {
	s.BizId = &v
	return s
}

func (s *GetOuterCallCenterDataListRequest) SetInstanceId(v string) *GetOuterCallCenterDataListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOuterCallCenterDataListRequest) SetQueryStartTime(v string) *GetOuterCallCenterDataListRequest {
	s.QueryStartTime = &v
	return s
}

func (s *GetOuterCallCenterDataListRequest) SetQueryEndTime(v string) *GetOuterCallCenterDataListRequest {
	s.QueryEndTime = &v
	return s
}

type GetOuterCallCenterDataListResponseBody struct {
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*GetOuterCallCenterDataListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetOuterCallCenterDataListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOuterCallCenterDataListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOuterCallCenterDataListResponseBody) SetMessage(v string) *GetOuterCallCenterDataListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBody) SetRequestId(v string) *GetOuterCallCenterDataListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBody) SetData(v []*GetOuterCallCenterDataListResponseBodyData) *GetOuterCallCenterDataListResponseBody {
	s.Data = v
	return s
}

func (s *GetOuterCallCenterDataListResponseBody) SetCode(v string) *GetOuterCallCenterDataListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBody) SetSuccess(v bool) *GetOuterCallCenterDataListResponseBody {
	s.Success = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBody) SetHttpStatusCode(v int64) *GetOuterCallCenterDataListResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetOuterCallCenterDataListResponseBodyData struct {
	EndReason     *string `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	CallType      *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	Acid          *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	ToPhoneNum    *string `json:"ToPhoneNum,omitempty" xml:"ToPhoneNum,omitempty"`
	UserInfo      *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
	InterveneTime *string `json:"InterveneTime,omitempty" xml:"InterveneTime,omitempty"`
	BizId         *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	FromPhoneNum  *string `json:"FromPhoneNum,omitempty" xml:"FromPhoneNum,omitempty"`
	ExtInfo       *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
}

func (s GetOuterCallCenterDataListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOuterCallCenterDataListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetEndReason(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.EndReason = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetCallType(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.CallType = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetAcid(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.Acid = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetToPhoneNum(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.ToPhoneNum = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetUserInfo(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.UserInfo = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetInterveneTime(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.InterveneTime = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetBizId(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.BizId = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetSessionId(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetFromPhoneNum(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.FromPhoneNum = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetExtInfo(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.ExtInfo = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetBizType(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.BizType = &v
	return s
}

type GetOuterCallCenterDataListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOuterCallCenterDataListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOuterCallCenterDataListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOuterCallCenterDataListResponse) GoString() string {
	return s.String()
}

func (s *GetOuterCallCenterDataListResponse) SetHeaders(v map[string]*string) *GetOuterCallCenterDataListResponse {
	s.Headers = v
	return s
}

func (s *GetOuterCallCenterDataListResponse) SetBody(v *GetOuterCallCenterDataListResponseBody) *GetOuterCallCenterDataListResponse {
	s.Body = v
	return s
}

type QueryTicketsRequest struct {
	InstanceId  *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CaseId      *int64                 `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	CaseType    *int32                 `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	CaseStatus  *int32                 `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	SrType      *int64                 `json:"SrType,omitempty" xml:"SrType,omitempty"`
	TaskStatus  *int32                 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	ChannelId   *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChannelType *int32                 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	TouchId     *int64                 `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
	DealId      *int64                 `json:"DealId,omitempty" xml:"DealId,omitempty"`
	Extra       map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	AccountName *string                `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	PageSize    *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryTicketsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketsRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketsRequest) SetInstanceId(v string) *QueryTicketsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTicketsRequest) SetCaseId(v int64) *QueryTicketsRequest {
	s.CaseId = &v
	return s
}

func (s *QueryTicketsRequest) SetCaseType(v int32) *QueryTicketsRequest {
	s.CaseType = &v
	return s
}

func (s *QueryTicketsRequest) SetCaseStatus(v int32) *QueryTicketsRequest {
	s.CaseStatus = &v
	return s
}

func (s *QueryTicketsRequest) SetSrType(v int64) *QueryTicketsRequest {
	s.SrType = &v
	return s
}

func (s *QueryTicketsRequest) SetTaskStatus(v int32) *QueryTicketsRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryTicketsRequest) SetChannelId(v string) *QueryTicketsRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryTicketsRequest) SetChannelType(v int32) *QueryTicketsRequest {
	s.ChannelType = &v
	return s
}

func (s *QueryTicketsRequest) SetTouchId(v int64) *QueryTicketsRequest {
	s.TouchId = &v
	return s
}

func (s *QueryTicketsRequest) SetDealId(v int64) *QueryTicketsRequest {
	s.DealId = &v
	return s
}

func (s *QueryTicketsRequest) SetExtra(v map[string]interface{}) *QueryTicketsRequest {
	s.Extra = v
	return s
}

func (s *QueryTicketsRequest) SetAccountName(v string) *QueryTicketsRequest {
	s.AccountName = &v
	return s
}

func (s *QueryTicketsRequest) SetPageSize(v int32) *QueryTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTicketsRequest) SetCurrentPage(v int32) *QueryTicketsRequest {
	s.CurrentPage = &v
	return s
}

type QueryTicketsShrinkRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CaseId      *int64  `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	CaseType    *int32  `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	CaseStatus  *int32  `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	SrType      *int64  `json:"SrType,omitempty" xml:"SrType,omitempty"`
	TaskStatus  *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChannelType *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	TouchId     *int64  `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
	DealId      *int64  `json:"DealId,omitempty" xml:"DealId,omitempty"`
	ExtraShrink *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
}

func (s QueryTicketsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketsShrinkRequest) SetInstanceId(v string) *QueryTicketsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCaseId(v int64) *QueryTicketsShrinkRequest {
	s.CaseId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCaseType(v int32) *QueryTicketsShrinkRequest {
	s.CaseType = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCaseStatus(v int32) *QueryTicketsShrinkRequest {
	s.CaseStatus = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetSrType(v int64) *QueryTicketsShrinkRequest {
	s.SrType = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetTaskStatus(v int32) *QueryTicketsShrinkRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetChannelId(v string) *QueryTicketsShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetChannelType(v int32) *QueryTicketsShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetTouchId(v int64) *QueryTicketsShrinkRequest {
	s.TouchId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetDealId(v int64) *QueryTicketsShrinkRequest {
	s.DealId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetExtraShrink(v string) *QueryTicketsShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetAccountName(v string) *QueryTicketsShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetPageSize(v int32) *QueryTicketsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCurrentPage(v int32) *QueryTicketsShrinkRequest {
	s.CurrentPage = &v
	return s
}

type QueryTicketsResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTicketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTicketsResponseBody) SetMessage(v string) *QueryTicketsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTicketsResponseBody) SetRequestId(v string) *QueryTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTicketsResponseBody) SetData(v string) *QueryTicketsResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTicketsResponseBody) SetCode(v string) *QueryTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTicketsResponseBody) SetSuccess(v bool) *QueryTicketsResponseBody {
	s.Success = &v
	return s
}

type QueryTicketsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTicketsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketsResponse) GoString() string {
	return s.String()
}

func (s *QueryTicketsResponse) SetHeaders(v map[string]*string) *QueryTicketsResponse {
	s.Headers = v
	return s
}

func (s *QueryTicketsResponse) SetBody(v *QueryTicketsResponseBody) *QueryTicketsResponse {
	s.Body = v
	return s
}

type QueryTicketActionsRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TicketId       *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	ActionCodeList []*int  `json:"ActionCodeList,omitempty" xml:"ActionCodeList,omitempty" type:"Repeated"`
}

func (s QueryTicketActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketActionsRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketActionsRequest) SetInstanceId(v string) *QueryTicketActionsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTicketActionsRequest) SetTicketId(v string) *QueryTicketActionsRequest {
	s.TicketId = &v
	return s
}

func (s *QueryTicketActionsRequest) SetActionCodeList(v []*int) *QueryTicketActionsRequest {
	s.ActionCodeList = v
	return s
}

type QueryTicketActionsResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTicketActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketActionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTicketActionsResponseBody) SetMessage(v string) *QueryTicketActionsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTicketActionsResponseBody) SetRequestId(v string) *QueryTicketActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTicketActionsResponseBody) SetData(v string) *QueryTicketActionsResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTicketActionsResponseBody) SetCode(v string) *QueryTicketActionsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTicketActionsResponseBody) SetSuccess(v bool) *QueryTicketActionsResponseBody {
	s.Success = &v
	return s
}

type QueryTicketActionsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryTicketActionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTicketActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketActionsResponse) GoString() string {
	return s.String()
}

func (s *QueryTicketActionsResponse) SetHeaders(v map[string]*string) *QueryTicketActionsResponse {
	s.Headers = v
	return s
}

func (s *QueryTicketActionsResponse) SetBody(v *QueryTicketActionsResponseBody) *QueryTicketActionsResponse {
	s.Body = v
	return s
}

type TransferCallToAgentRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName       *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	TargetAccountName *string `json:"TargetAccountName,omitempty" xml:"TargetAccountName,omitempty"`
	CallId            *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId             *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId      *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId  *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
	Type              *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	IsSingleTransfer  *string `json:"IsSingleTransfer,omitempty" xml:"IsSingleTransfer,omitempty"`
}

func (s TransferCallToAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToAgentRequest) GoString() string {
	return s.String()
}

func (s *TransferCallToAgentRequest) SetClientToken(v string) *TransferCallToAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *TransferCallToAgentRequest) SetInstanceId(v string) *TransferCallToAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *TransferCallToAgentRequest) SetAccountName(v string) *TransferCallToAgentRequest {
	s.AccountName = &v
	return s
}

func (s *TransferCallToAgentRequest) SetTargetAccountName(v string) *TransferCallToAgentRequest {
	s.TargetAccountName = &v
	return s
}

func (s *TransferCallToAgentRequest) SetCallId(v string) *TransferCallToAgentRequest {
	s.CallId = &v
	return s
}

func (s *TransferCallToAgentRequest) SetJobId(v string) *TransferCallToAgentRequest {
	s.JobId = &v
	return s
}

func (s *TransferCallToAgentRequest) SetConnectionId(v string) *TransferCallToAgentRequest {
	s.ConnectionId = &v
	return s
}

func (s *TransferCallToAgentRequest) SetHoldConnectionId(v string) *TransferCallToAgentRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *TransferCallToAgentRequest) SetType(v int32) *TransferCallToAgentRequest {
	s.Type = &v
	return s
}

func (s *TransferCallToAgentRequest) SetIsSingleTransfer(v string) *TransferCallToAgentRequest {
	s.IsSingleTransfer = &v
	return s
}

type TransferCallToAgentResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferCallToAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToAgentResponseBody) GoString() string {
	return s.String()
}

func (s *TransferCallToAgentResponseBody) SetMessage(v string) *TransferCallToAgentResponseBody {
	s.Message = &v
	return s
}

func (s *TransferCallToAgentResponseBody) SetRequestId(v string) *TransferCallToAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferCallToAgentResponseBody) SetCode(v string) *TransferCallToAgentResponseBody {
	s.Code = &v
	return s
}

func (s *TransferCallToAgentResponseBody) SetSuccess(v bool) *TransferCallToAgentResponseBody {
	s.Success = &v
	return s
}

type TransferCallToAgentResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferCallToAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferCallToAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToAgentResponse) GoString() string {
	return s.String()
}

func (s *TransferCallToAgentResponse) SetHeaders(v map[string]*string) *TransferCallToAgentResponse {
	s.Headers = v
	return s
}

func (s *TransferCallToAgentResponse) SetBody(v *TransferCallToAgentResponseBody) *TransferCallToAgentResponse {
	s.Body = v
	return s
}

type FinishHotlineServiceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s FinishHotlineServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishHotlineServiceRequest) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceRequest) SetClientToken(v string) *FinishHotlineServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *FinishHotlineServiceRequest) SetInstanceId(v string) *FinishHotlineServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *FinishHotlineServiceRequest) SetAccountName(v string) *FinishHotlineServiceRequest {
	s.AccountName = &v
	return s
}

type FinishHotlineServiceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FinishHotlineServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishHotlineServiceResponseBody) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceResponseBody) SetMessage(v string) *FinishHotlineServiceResponseBody {
	s.Message = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetRequestId(v string) *FinishHotlineServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetCode(v string) *FinishHotlineServiceResponseBody {
	s.Code = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetSuccess(v bool) *FinishHotlineServiceResponseBody {
	s.Success = &v
	return s
}

type FinishHotlineServiceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FinishHotlineServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FinishHotlineServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishHotlineServiceResponse) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceResponse) SetHeaders(v map[string]*string) *FinishHotlineServiceResponse {
	s.Headers = v
	return s
}

func (s *FinishHotlineServiceResponse) SetBody(v *FinishHotlineServiceResponseBody) *FinishHotlineServiceResponse {
	s.Body = v
	return s
}

type JoinThirdCallRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
}

func (s JoinThirdCallRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinThirdCallRequest) GoString() string {
	return s.String()
}

func (s *JoinThirdCallRequest) SetClientToken(v string) *JoinThirdCallRequest {
	s.ClientToken = &v
	return s
}

func (s *JoinThirdCallRequest) SetInstanceId(v string) *JoinThirdCallRequest {
	s.InstanceId = &v
	return s
}

func (s *JoinThirdCallRequest) SetAccountName(v string) *JoinThirdCallRequest {
	s.AccountName = &v
	return s
}

func (s *JoinThirdCallRequest) SetCallId(v string) *JoinThirdCallRequest {
	s.CallId = &v
	return s
}

func (s *JoinThirdCallRequest) SetJobId(v string) *JoinThirdCallRequest {
	s.JobId = &v
	return s
}

func (s *JoinThirdCallRequest) SetConnectionId(v string) *JoinThirdCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *JoinThirdCallRequest) SetHoldConnectionId(v string) *JoinThirdCallRequest {
	s.HoldConnectionId = &v
	return s
}

type JoinThirdCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JoinThirdCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinThirdCallResponseBody) GoString() string {
	return s.String()
}

func (s *JoinThirdCallResponseBody) SetMessage(v string) *JoinThirdCallResponseBody {
	s.Message = &v
	return s
}

func (s *JoinThirdCallResponseBody) SetRequestId(v string) *JoinThirdCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinThirdCallResponseBody) SetCode(v string) *JoinThirdCallResponseBody {
	s.Code = &v
	return s
}

func (s *JoinThirdCallResponseBody) SetSuccess(v bool) *JoinThirdCallResponseBody {
	s.Success = &v
	return s
}

type JoinThirdCallResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *JoinThirdCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinThirdCallResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinThirdCallResponse) GoString() string {
	return s.String()
}

func (s *JoinThirdCallResponse) SetHeaders(v map[string]*string) *JoinThirdCallResponse {
	s.Headers = v
	return s
}

func (s *JoinThirdCallResponse) SetBody(v *JoinThirdCallResponseBody) *JoinThirdCallResponse {
	s.Body = v
	return s
}

type ExecuteActivityRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TicketId     *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	OperatorId   *int64  `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	ActivityCode *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	ActivityForm *string `json:"ActivityForm,omitempty" xml:"ActivityForm,omitempty"`
}

func (s ExecuteActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityRequest) GoString() string {
	return s.String()
}

func (s *ExecuteActivityRequest) SetClientToken(v string) *ExecuteActivityRequest {
	s.ClientToken = &v
	return s
}

func (s *ExecuteActivityRequest) SetInstanceId(v string) *ExecuteActivityRequest {
	s.InstanceId = &v
	return s
}

func (s *ExecuteActivityRequest) SetTicketId(v int64) *ExecuteActivityRequest {
	s.TicketId = &v
	return s
}

func (s *ExecuteActivityRequest) SetOperatorId(v int64) *ExecuteActivityRequest {
	s.OperatorId = &v
	return s
}

func (s *ExecuteActivityRequest) SetActivityCode(v string) *ExecuteActivityRequest {
	s.ActivityCode = &v
	return s
}

func (s *ExecuteActivityRequest) SetActivityForm(v string) *ExecuteActivityRequest {
	s.ActivityForm = &v
	return s
}

type ExecuteActivityResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteActivityResponseBody) SetMessage(v string) *ExecuteActivityResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteActivityResponseBody) SetRequestId(v string) *ExecuteActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteActivityResponseBody) SetCode(v string) *ExecuteActivityResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteActivityResponseBody) SetSuccess(v bool) *ExecuteActivityResponseBody {
	s.Success = &v
	return s
}

type ExecuteActivityResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteActivityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityResponse) GoString() string {
	return s.String()
}

func (s *ExecuteActivityResponse) SetHeaders(v map[string]*string) *ExecuteActivityResponse {
	s.Headers = v
	return s
}

func (s *ExecuteActivityResponse) SetBody(v *ExecuteActivityResponseBody) *ExecuteActivityResponse {
	s.Body = v
	return s
}

type GetGrantedRoleIdsRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s GetGrantedRoleIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGrantedRoleIdsRequest) GoString() string {
	return s.String()
}

func (s *GetGrantedRoleIdsRequest) SetClientToken(v string) *GetGrantedRoleIdsRequest {
	s.ClientToken = &v
	return s
}

func (s *GetGrantedRoleIdsRequest) SetInstanceId(v string) *GetGrantedRoleIdsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetGrantedRoleIdsRequest) SetAccountName(v string) *GetGrantedRoleIdsRequest {
	s.AccountName = &v
	return s
}

type GetGrantedRoleIdsResponseBody struct {
	Message        *string  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code           *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetGrantedRoleIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGrantedRoleIdsResponseBody) GoString() string {
	return s.String()
}

func (s *GetGrantedRoleIdsResponseBody) SetMessage(v string) *GetGrantedRoleIdsResponseBody {
	s.Message = &v
	return s
}

func (s *GetGrantedRoleIdsResponseBody) SetRequestId(v string) *GetGrantedRoleIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGrantedRoleIdsResponseBody) SetData(v []*int64) *GetGrantedRoleIdsResponseBody {
	s.Data = v
	return s
}

func (s *GetGrantedRoleIdsResponseBody) SetCode(v string) *GetGrantedRoleIdsResponseBody {
	s.Code = &v
	return s
}

func (s *GetGrantedRoleIdsResponseBody) SetSuccess(v bool) *GetGrantedRoleIdsResponseBody {
	s.Success = &v
	return s
}

func (s *GetGrantedRoleIdsResponseBody) SetHttpStatusCode(v int64) *GetGrantedRoleIdsResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetGrantedRoleIdsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGrantedRoleIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGrantedRoleIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGrantedRoleIdsResponse) GoString() string {
	return s.String()
}

func (s *GetGrantedRoleIdsResponse) SetHeaders(v map[string]*string) *GetGrantedRoleIdsResponse {
	s.Headers = v
	return s
}

func (s *GetGrantedRoleIdsResponse) SetBody(v *GetGrantedRoleIdsResponseBody) *GetGrantedRoleIdsResponse {
	s.Body = v
	return s
}

type ListHotlineRecordRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CallId      *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
}

func (s ListHotlineRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordRequest) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordRequest) SetClientToken(v string) *ListHotlineRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *ListHotlineRecordRequest) SetInstanceId(v string) *ListHotlineRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHotlineRecordRequest) SetCallId(v string) *ListHotlineRecordRequest {
	s.CallId = &v
	return s
}

type ListHotlineRecordResponseBody struct {
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*ListHotlineRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s ListHotlineRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordResponseBody) SetMessage(v string) *ListHotlineRecordResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetRequestId(v string) *ListHotlineRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetData(v []*ListHotlineRecordResponseBodyData) *ListHotlineRecordResponseBody {
	s.Data = v
	return s
}

func (s *ListHotlineRecordResponseBody) SetCode(v string) *ListHotlineRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetSuccess(v bool) *ListHotlineRecordResponseBody {
	s.Success = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetHttpStatusCode(v int64) *ListHotlineRecordResponseBody {
	s.HttpStatusCode = &v
	return s
}

type ListHotlineRecordResponseBodyData struct {
	EndTime      *bool   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *bool   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListHotlineRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordResponseBodyData) SetEndTime(v bool) *ListHotlineRecordResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetStartTime(v bool) *ListHotlineRecordResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetConnectionId(v string) *ListHotlineRecordResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetCallId(v string) *ListHotlineRecordResponseBodyData {
	s.CallId = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetUrl(v string) *ListHotlineRecordResponseBodyData {
	s.Url = &v
	return s
}

type ListHotlineRecordResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHotlineRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotlineRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordResponse) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordResponse) SetHeaders(v map[string]*string) *ListHotlineRecordResponse {
	s.Headers = v
	return s
}

func (s *ListHotlineRecordResponse) SetBody(v *ListHotlineRecordResponseBody) *ListHotlineRecordResponse {
	s.Body = v
	return s
}

type GetNumLocationRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PhoneNum    *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
}

func (s GetNumLocationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNumLocationRequest) GoString() string {
	return s.String()
}

func (s *GetNumLocationRequest) SetClientToken(v string) *GetNumLocationRequest {
	s.ClientToken = &v
	return s
}

func (s *GetNumLocationRequest) SetInstanceId(v string) *GetNumLocationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNumLocationRequest) SetPhoneNum(v string) *GetNumLocationRequest {
	s.PhoneNum = &v
	return s
}

type GetNumLocationResponseBody struct {
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
}

func (s GetNumLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNumLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetNumLocationResponseBody) SetMessage(v string) *GetNumLocationResponseBody {
	s.Message = &v
	return s
}

func (s *GetNumLocationResponseBody) SetRequestId(v string) *GetNumLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNumLocationResponseBody) SetData(v string) *GetNumLocationResponseBody {
	s.Data = &v
	return s
}

func (s *GetNumLocationResponseBody) SetCode(v string) *GetNumLocationResponseBody {
	s.Code = &v
	return s
}

func (s *GetNumLocationResponseBody) SetSuccess(v bool) *GetNumLocationResponseBody {
	s.Success = &v
	return s
}

func (s *GetNumLocationResponseBody) SetHttpStatusCode(v int64) *GetNumLocationResponseBody {
	s.HttpStatusCode = &v
	return s
}

type GetNumLocationResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetNumLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNumLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNumLocationResponse) GoString() string {
	return s.String()
}

func (s *GetNumLocationResponse) SetHeaders(v map[string]*string) *GetNumLocationResponse {
	s.Headers = v
	return s
}

func (s *GetNumLocationResponse) SetBody(v *GetNumLocationResponseBody) *GetNumLocationResponse {
	s.Body = v
	return s
}

type CreateSkillGroupRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ChannelType    *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupRequest) SetInstanceId(v string) *CreateSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSkillGroupRequest) SetSkillGroupName(v string) *CreateSkillGroupRequest {
	s.SkillGroupName = &v
	return s
}

func (s *CreateSkillGroupRequest) SetDescription(v string) *CreateSkillGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateSkillGroupRequest) SetDisplayName(v string) *CreateSkillGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateSkillGroupRequest) SetChannelType(v int32) *CreateSkillGroupRequest {
	s.ChannelType = &v
	return s
}

func (s *CreateSkillGroupRequest) SetClientToken(v string) *CreateSkillGroupRequest {
	s.ClientToken = &v
	return s
}

type CreateSkillGroupResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupResponseBody) SetMessage(v string) *CreateSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetRequestId(v string) *CreateSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetData(v int64) *CreateSkillGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetCode(v string) *CreateSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetSuccess(v bool) *CreateSkillGroupResponseBody {
	s.Success = &v
	return s
}

type CreateSkillGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupResponse) SetHeaders(v map[string]*string) *CreateSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillGroupResponse) SetBody(v *CreateSkillGroupResponseBody) *CreateSkillGroupResponse {
	s.Body = v
	return s
}

type CreateCustomerRequest struct {
	ProdLineId  *int64  `json:"ProdLineId,omitempty" xml:"ProdLineId,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TypeCode    *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ManagerName *string `json:"ManagerName,omitempty" xml:"ManagerName,omitempty"`
	Contacter   *string `json:"Contacter,omitempty" xml:"Contacter,omitempty"`
	Industry    *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	Position    *string `json:"Position,omitempty" xml:"Position,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Dingding    *string `json:"Dingding,omitempty" xml:"Dingding,omitempty"`
	OuterId     *string `json:"OuterId,omitempty" xml:"OuterId,omitempty"`
	OuterIdType *int32  `json:"OuterIdType,omitempty" xml:"OuterIdType,omitempty"`
}

func (s CreateCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequest) SetProdLineId(v int64) *CreateCustomerRequest {
	s.ProdLineId = &v
	return s
}

func (s *CreateCustomerRequest) SetBizType(v string) *CreateCustomerRequest {
	s.BizType = &v
	return s
}

func (s *CreateCustomerRequest) SetName(v string) *CreateCustomerRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomerRequest) SetTypeCode(v string) *CreateCustomerRequest {
	s.TypeCode = &v
	return s
}

func (s *CreateCustomerRequest) SetPhone(v string) *CreateCustomerRequest {
	s.Phone = &v
	return s
}

func (s *CreateCustomerRequest) SetInstanceId(v string) *CreateCustomerRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomerRequest) SetManagerName(v string) *CreateCustomerRequest {
	s.ManagerName = &v
	return s
}

func (s *CreateCustomerRequest) SetContacter(v string) *CreateCustomerRequest {
	s.Contacter = &v
	return s
}

func (s *CreateCustomerRequest) SetIndustry(v string) *CreateCustomerRequest {
	s.Industry = &v
	return s
}

func (s *CreateCustomerRequest) SetPosition(v string) *CreateCustomerRequest {
	s.Position = &v
	return s
}

func (s *CreateCustomerRequest) SetEmail(v string) *CreateCustomerRequest {
	s.Email = &v
	return s
}

func (s *CreateCustomerRequest) SetDingding(v string) *CreateCustomerRequest {
	s.Dingding = &v
	return s
}

func (s *CreateCustomerRequest) SetOuterId(v string) *CreateCustomerRequest {
	s.OuterId = &v
	return s
}

func (s *CreateCustomerRequest) SetOuterIdType(v int32) *CreateCustomerRequest {
	s.OuterIdType = &v
	return s
}

type CreateCustomerResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponseBody) SetMessage(v string) *CreateCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCustomerResponseBody) SetRequestId(v string) *CreateCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomerResponseBody) SetData(v int64) *CreateCustomerResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCustomerResponseBody) SetCode(v string) *CreateCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomerResponseBody) SetSuccess(v bool) *CreateCustomerResponseBody {
	s.Success = &v
	return s
}

type CreateCustomerResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponse) SetHeaders(v map[string]*string) *CreateCustomerResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomerResponse) SetBody(v *CreateCustomerResponseBody) *CreateCustomerResponse {
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
	client.EndpointMap = map[string]*string{
		"cn-shanghai": tea.String("scsp-vpc.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("scsp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetUserTokenWithOptions(request *GetUserTokenRequest, runtime *util.RuntimeOptions) (_result *GetUserTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserToken"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserToken(request *GetUserTokenRequest) (_result *GetUserTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserTokenResponse{}
	_body, _err := client.GetUserTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTicketListWithOptions(request *SearchTicketListRequest, runtime *util.RuntimeOptions) (_result *SearchTicketListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SearchTicketListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchTicketList"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTicketList(request *SearchTicketListRequest) (_result *SearchTicketListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTicketListResponse{}
	_body, _err := client.SearchTicketListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendHotlineHeartBeatWithOptions(request *SendHotlineHeartBeatRequest, runtime *util.RuntimeOptions) (_result *SendHotlineHeartBeatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendHotlineHeartBeatResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendHotlineHeartBeat"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendHotlineHeartBeat(request *SendHotlineHeartBeatRequest) (_result *SendHotlineHeartBeatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendHotlineHeartBeatResponse{}
	_body, _err := client.SendHotlineHeartBeatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferCallToSkillGroupWithOptions(request *TransferCallToSkillGroupRequest, runtime *util.RuntimeOptions) (_result *TransferCallToSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TransferCallToSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TransferCallToSkillGroup"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferCallToSkillGroup(request *TransferCallToSkillGroupRequest) (_result *TransferCallToSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferCallToSkillGroupResponse{}
	_body, _err := client.TransferCallToSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditEntityRouteWithOptions(request *EditEntityRouteRequest, runtime *util.RuntimeOptions) (_result *EditEntityRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EditEntityRouteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EditEntityRoute"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditEntityRoute(request *EditEntityRouteRequest) (_result *EditEntityRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditEntityRouteResponse{}
	_body, _err := client.EditEntityRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTagListWithOptions(request *GetTagListRequest, runtime *util.RuntimeOptions) (_result *GetTagListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTagListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTagList"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTagList(request *GetTagListRequest) (_result *GetTagListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTagListResponse{}
	_body, _err := client.GetTagListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTicketWithOptions(request *UpdateTicketRequest, runtime *util.RuntimeOptions) (_result *UpdateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTicket"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTicket(request *UpdateTicketRequest) (_result *UpdateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTicketResponse{}
	_body, _err := client.UpdateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOutboundPhoneNumberWithOptions(request *ListOutboundPhoneNumberRequest, runtime *util.RuntimeOptions) (_result *ListOutboundPhoneNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListOutboundPhoneNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOutboundPhoneNumber"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOutboundPhoneNumber(request *ListOutboundPhoneNumberRequest) (_result *ListOutboundPhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOutboundPhoneNumberResponse{}
	_body, _err := client.ListOutboundPhoneNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSkillGroupWithOptions(request *RemoveSkillGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveSkillGroup"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSkillGroup(request *RemoveSkillGroupRequest) (_result *RemoveSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveSkillGroupResponse{}
	_body, _err := client.RemoveSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAgentBySkillGroupIdWithOptions(request *ListAgentBySkillGroupIdRequest, runtime *util.RuntimeOptions) (_result *ListAgentBySkillGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListAgentBySkillGroupIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAgentBySkillGroupId"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAgentBySkillGroupId(request *ListAgentBySkillGroupIdRequest) (_result *ListAgentBySkillGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAgentBySkillGroupIdResponse{}
	_body, _err := client.ListAgentBySkillGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHotlineSessionWithOptions(request *QueryHotlineSessionRequest, runtime *util.RuntimeOptions) (_result *QueryHotlineSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryHotlineSessionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryHotlineSession"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHotlineSession(request *QueryHotlineSessionRequest) (_result *QueryHotlineSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryHotlineSessionResponse{}
	_body, _err := client.QueryHotlineSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartChatWorkWithOptions(request *StartChatWorkRequest, runtime *util.RuntimeOptions) (_result *StartChatWorkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartChatWorkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartChatWork"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartChatWork(request *StartChatWorkRequest) (_result *StartChatWorkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartChatWorkResponse{}
	_body, _err := client.StartChatWorkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HangupThirdCallWithOptions(request *HangupThirdCallRequest, runtime *util.RuntimeOptions) (_result *HangupThirdCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HangupThirdCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("HangupThirdCall"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HangupThirdCall(request *HangupThirdCallRequest) (_result *HangupThirdCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HangupThirdCallResponse{}
	_body, _err := client.HangupThirdCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartHotlineServiceWithOptions(request *StartHotlineServiceRequest, runtime *util.RuntimeOptions) (_result *StartHotlineServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartHotlineServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartHotlineService"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartHotlineService(request *StartHotlineServiceRequest) (_result *StartHotlineServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartHotlineServiceResponse{}
	_body, _err := client.StartHotlineServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCallV2WithOptions(request *StartCallV2Request, runtime *util.RuntimeOptions) (_result *StartCallV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartCallV2Response{}
	_body, _err := client.DoRPCRequest(tea.String("StartCallV2"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCallV2(request *StartCallV2Request) (_result *StartCallV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartCallV2Response{}
	_body, _err := client.StartCallV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableRoleWithOptions(request *EnableRoleRequest, runtime *util.RuntimeOptions) (_result *EnableRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableRole"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableRole(request *EnableRoleRequest) (_result *EnableRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableRoleResponse{}
	_body, _err := client.EnableRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentWithOptions(request *GetAgentRequest, runtime *util.RuntimeOptions) (_result *GetAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAgent"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgent(request *GetAgentRequest) (_result *GetAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentResponse{}
	_body, _err := client.GetAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCMSIdByForeignIdWithOptions(request *GetCMSIdByForeignIdRequest, runtime *util.RuntimeOptions) (_result *GetCMSIdByForeignIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCMSIdByForeignIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCMSIdByForeignId"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCMSIdByForeignId(request *GetCMSIdByForeignIdRequest) (_result *GetCMSIdByForeignIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCMSIdByForeignIdResponse{}
	_body, _err := client.GetCMSIdByForeignIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferToThirdCallWithOptions(request *TransferToThirdCallRequest, runtime *util.RuntimeOptions) (_result *TransferToThirdCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TransferToThirdCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TransferToThirdCall"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferToThirdCall(request *TransferToThirdCallRequest) (_result *TransferToThirdCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferToThirdCallResponse{}
	_body, _err := client.TransferToThirdCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAgentWithOptions(request *UpdateAgentRequest, runtime *util.RuntimeOptions) (_result *UpdateAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAgent"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAgent(request *UpdateAgentRequest) (_result *UpdateAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAgentResponse{}
	_body, _err := client.UpdateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeChatAgentStatusWithOptions(request *ChangeChatAgentStatusRequest, runtime *util.RuntimeOptions) (_result *ChangeChatAgentStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeChatAgentStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeChatAgentStatus"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeChatAgentStatus(request *ChangeChatAgentStatusRequest) (_result *ChangeChatAgentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeChatAgentStatusResponse{}
	_body, _err := client.ChangeChatAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateWebSocketSignWithOptions(request *GenerateWebSocketSignRequest, runtime *util.RuntimeOptions) (_result *GenerateWebSocketSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateWebSocketSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateWebSocketSign"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateWebSocketSign(request *GenerateWebSocketSignRequest) (_result *GenerateWebSocketSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateWebSocketSignResponse{}
	_body, _err := client.GenerateWebSocketSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRingStatusWithOptions(request *UpdateRingStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateRingStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateRingStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateRingStatus"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRingStatus(request *UpdateRingStatusRequest) (_result *UpdateRingStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRingStatusResponse{}
	_body, _err := client.UpdateRingStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAgentWithOptions(request *CreateAgentRequest, runtime *util.RuntimeOptions) (_result *CreateAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAgent"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAgent(request *CreateAgentRequest) (_result *CreateAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAgentResponse{}
	_body, _err := client.CreateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEntityRouteWithOptions(request *DeleteEntityRouteRequest, runtime *util.RuntimeOptions) (_result *DeleteEntityRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEntityRouteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEntityRoute"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEntityRoute(request *DeleteEntityRouteRequest) (_result *DeleteEntityRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEntityRouteResponse{}
	_body, _err := client.DeleteEntityRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineGroupDetailReportWithOptions(request *GetHotlineGroupDetailReportRequest, runtime *util.RuntimeOptions) (_result *GetHotlineGroupDetailReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHotlineGroupDetailReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineGroupDetailReport"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineGroupDetailReport(request *GetHotlineGroupDetailReportRequest) (_result *GetHotlineGroupDetailReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineGroupDetailReportResponse{}
	_body, _err := client.GetHotlineGroupDetailReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, runtime *util.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTicket"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCustomerWithOptions(request *UpdateCustomerRequest, runtime *util.RuntimeOptions) (_result *UpdateCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCustomerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCustomer"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCustomer(request *UpdateCustomerRequest) (_result *UpdateCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCustomerResponse{}
	_body, _err := client.UpdateCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AnswerCallWithOptions(request *AnswerCallRequest, runtime *util.RuntimeOptions) (_result *AnswerCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AnswerCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AnswerCall"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AnswerCall(request *AnswerCallRequest) (_result *AnswerCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AnswerCallResponse{}
	_body, _err := client.AnswerCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAgentWithOptions(request *DeleteAgentRequest, runtime *util.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAgent"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAgent(request *DeleteAgentRequest) (_result *DeleteAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAgentResponse{}
	_body, _err := client.DeleteAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEntityTagRelationWithOptions(request *GetEntityTagRelationRequest, runtime *util.RuntimeOptions) (_result *GetEntityTagRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetEntityTagRelationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEntityTagRelation"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEntityTagRelation(request *GetEntityTagRelationRequest) (_result *GetEntityTagRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEntityTagRelationResponse{}
	_body, _err := client.GetEntityTagRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineAgentDetailWithOptions(request *GetHotlineAgentDetailRequest, runtime *util.RuntimeOptions) (_result *GetHotlineAgentDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetHotlineAgentDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineAgentDetail"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineAgentDetail(request *GetHotlineAgentDetailRequest) (_result *GetHotlineAgentDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineAgentDetailResponse{}
	_body, _err := client.GetHotlineAgentDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuspendHotlineServiceWithOptions(request *SuspendHotlineServiceRequest, runtime *util.RuntimeOptions) (_result *SuspendHotlineServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SuspendHotlineServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SuspendHotlineService"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendHotlineService(request *SuspendHotlineServiceRequest) (_result *SuspendHotlineServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendHotlineServiceResponse{}
	_body, _err := client.SuspendHotlineServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCallsPerDayWithOptions(request *GetCallsPerDayRequest, runtime *util.RuntimeOptions) (_result *GetCallsPerDayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetCallsPerDayResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCallsPerDay"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCallsPerDay(request *GetCallsPerDayRequest) (_result *GetCallsPerDayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCallsPerDayResponse{}
	_body, _err := client.GetCallsPerDayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FetchCallWithOptions(request *FetchCallRequest, runtime *util.RuntimeOptions) (_result *FetchCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FetchCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FetchCall"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FetchCall(request *FetchCallRequest) (_result *FetchCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FetchCallResponse{}
	_body, _err := client.FetchCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineAgentDetailReportWithOptions(request *GetHotlineAgentDetailReportRequest, runtime *util.RuntimeOptions) (_result *GetHotlineAgentDetailReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHotlineAgentDetailReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineAgentDetailReport"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineAgentDetailReport(request *GetHotlineAgentDetailReportRequest) (_result *GetHotlineAgentDetailReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineAgentDetailReportResponse{}
	_body, _err := client.GetHotlineAgentDetailReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTicketCountWithOptions(request *QueryTicketCountRequest, runtime *util.RuntimeOptions) (_result *QueryTicketCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTicketCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTicketCount"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTicketCount(request *QueryTicketCountRequest) (_result *QueryTicketCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTicketCountResponse{}
	_body, _err := client.QueryTicketCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppMessagePushWithOptions(request *AppMessagePushRequest, runtime *util.RuntimeOptions) (_result *AppMessagePushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AppMessagePushResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AppMessagePush"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppMessagePush(request *AppMessagePushRequest) (_result *AppMessagePushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppMessagePushResponse{}
	_body, _err := client.AppMessagePushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineAgentStatusWithOptions(request *GetHotlineAgentStatusRequest, runtime *util.RuntimeOptions) (_result *GetHotlineAgentStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHotlineAgentStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineAgentStatus"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineAgentStatus(request *GetHotlineAgentStatusRequest) (_result *GetHotlineAgentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineAgentStatusResponse{}
	_body, _err := client.GetHotlineAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineWaitingNumberWithOptions(request *GetHotlineWaitingNumberRequest, runtime *util.RuntimeOptions) (_result *GetHotlineWaitingNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetHotlineWaitingNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHotlineWaitingNumber"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineWaitingNumber(request *GetHotlineWaitingNumberRequest) (_result *GetHotlineWaitingNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineWaitingNumberResponse{}
	_body, _err := client.GetHotlineWaitingNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCallWithOptions(request *StartCallRequest, runtime *util.RuntimeOptions) (_result *StartCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartCall"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCall(request *StartCallRequest) (_result *StartCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartCallResponse{}
	_body, _err := client.StartCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssignTicketWithOptions(request *AssignTicketRequest, runtime *util.RuntimeOptions) (_result *AssignTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssignTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssignTicket"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssignTicket(request *AssignTicketRequest) (_result *AssignTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssignTicketResponse{}
	_body, _err := client.AssignTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HangupCallWithOptions(request *HangupCallRequest, runtime *util.RuntimeOptions) (_result *HangupCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HangupCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("HangupCall"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HangupCall(request *HangupCallRequest) (_result *HangupCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HangupCallResponse{}
	_body, _err := client.HangupCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOutbounNumListWithOptions(request *GetOutbounNumListRequest, runtime *util.RuntimeOptions) (_result *GetOutbounNumListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOutbounNumListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOutbounNumList"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOutbounNumList(request *GetOutbounNumListRequest) (_result *GetOutbounNumListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOutbounNumListResponse{}
	_body, _err := client.GetOutbounNumListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTicketWithBizDataWithOptions(request *CreateTicketWithBizDataRequest, runtime *util.RuntimeOptions) (_result *CreateTicketWithBizDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTicketWithBizDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("createTicketWithBizData"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTicketWithBizData(request *CreateTicketWithBizDataRequest) (_result *CreateTicketWithBizDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTicketWithBizDataResponse{}
	_body, _err := client.CreateTicketWithBizDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTicketByPhoneWithOptions(request *SearchTicketByPhoneRequest, runtime *util.RuntimeOptions) (_result *SearchTicketByPhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SearchTicketByPhoneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchTicketByPhone"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTicketByPhone(request *SearchTicketByPhoneRequest) (_result *SearchTicketByPhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTicketByPhoneResponse{}
	_body, _err := client.SearchTicketByPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateThirdSsoAgentWithOptions(request *CreateThirdSsoAgentRequest, runtime *util.RuntimeOptions) (_result *CreateThirdSsoAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateThirdSsoAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateThirdSsoAgent"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateThirdSsoAgent(request *CreateThirdSsoAgentRequest) (_result *CreateThirdSsoAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateThirdSsoAgentResponse{}
	_body, _err := client.CreateThirdSsoAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEntityIvrRouteWithOptions(request *CreateEntityIvrRouteRequest, runtime *util.RuntimeOptions) (_result *CreateEntityIvrRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEntityIvrRouteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEntityIvrRoute"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEntityIvrRoute(request *CreateEntityIvrRouteRequest) (_result *CreateEntityIvrRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEntityIvrRouteResponse{}
	_body, _err := client.CreateEntityIvrRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseTicketWithOptions(request *CloseTicketRequest, runtime *util.RuntimeOptions) (_result *CloseTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CloseTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CloseTicket"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseTicket(request *CloseTicketRequest) (_result *CloseTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseTicketResponse{}
	_body, _err := client.CloseTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HoldCallWithOptions(request *HoldCallRequest, runtime *util.RuntimeOptions) (_result *HoldCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &HoldCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("HoldCall"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HoldCall(request *HoldCallRequest) (_result *HoldCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HoldCallResponse{}
	_body, _err := client.HoldCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRingDetailListWithOptions(tmpReq *QueryRingDetailListRequest, runtime *util.RuntimeOptions) (_result *QueryRingDetailListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryRingDetailListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FromPhoneNumList)) {
		request.FromPhoneNumListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FromPhoneNumList, tea.String("FromPhoneNumList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ToPhoneNumList)) {
		request.ToPhoneNumListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ToPhoneNumList, tea.String("ToPhoneNumList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRingDetailListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRingDetailList"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRingDetailList(request *QueryRingDetailListRequest) (_result *QueryRingDetailListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRingDetailListResponse{}
	_body, _err := client.QueryRingDetailListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTicketByIdWithOptions(request *SearchTicketByIdRequest, runtime *util.RuntimeOptions) (_result *SearchTicketByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SearchTicketByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchTicketById"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTicketById(request *SearchTicketByIdRequest) (_result *SearchTicketByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTicketByIdResponse{}
	_body, _err := client.SearchTicketByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSkillGroupWithOptions(request *UpdateSkillGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSkillGroup"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSkillGroup(request *UpdateSkillGroupRequest) (_result *UpdateSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSkillGroupResponse{}
	_body, _err := client.UpdateSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServiceConfigWithOptions(request *QueryServiceConfigRequest, runtime *util.RuntimeOptions) (_result *QueryServiceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &QueryServiceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryServiceConfig"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServiceConfig(request *QueryServiceConfigRequest) (_result *QueryServiceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryServiceConfigResponse{}
	_body, _err := client.QueryServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableRoleWithOptions(request *DisableRoleRequest, runtime *util.RuntimeOptions) (_result *DisableRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableRole"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableRole(request *DisableRoleRequest) (_result *DisableRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableRoleResponse{}
	_body, _err := client.DisableRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEntityRouteListWithOptions(request *GetEntityRouteListRequest, runtime *util.RuntimeOptions) (_result *GetEntityRouteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetEntityRouteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEntityRouteList"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEntityRouteList(request *GetEntityRouteListRequest) (_result *GetEntityRouteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEntityRouteListResponse{}
	_body, _err := client.GetEntityRouteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuthInfoWithOptions(request *GetAuthInfoRequest, runtime *util.RuntimeOptions) (_result *GetAuthInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAuthInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAuthInfo"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuthInfo(request *GetAuthInfoRequest) (_result *GetAuthInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthInfoResponse{}
	_body, _err := client.GetAuthInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRoleWithOptions(request *UpdateRoleRequest, runtime *util.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateRole"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRole(request *UpdateRoleRequest) (_result *UpdateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRoleResponse{}
	_body, _err := client.UpdateRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTicketTemplateSchemaWithOptions(request *GetTicketTemplateSchemaRequest, runtime *util.RuntimeOptions) (_result *GetTicketTemplateSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetTicketTemplateSchemaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTicketTemplateSchema"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTicketTemplateSchema(request *GetTicketTemplateSchemaRequest) (_result *GetTicketTemplateSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTicketTemplateSchemaResponse{}
	_body, _err := client.GetTicketTemplateSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferCallToPhoneWithOptions(request *TransferCallToPhoneRequest, runtime *util.RuntimeOptions) (_result *TransferCallToPhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TransferCallToPhoneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TransferCallToPhone"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferCallToPhone(request *TransferCallToPhoneRequest) (_result *TransferCallToPhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferCallToPhoneResponse{}
	_body, _err := client.TransferCallToPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySkillGroupsWithOptions(request *QuerySkillGroupsRequest, runtime *util.RuntimeOptions) (_result *QuerySkillGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QuerySkillGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QuerySkillGroups"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySkillGroups(request *QuerySkillGroupsRequest) (_result *QuerySkillGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySkillGroupsResponse{}
	_body, _err := client.QuerySkillGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEntityRouteWithOptions(request *GetEntityRouteRequest, runtime *util.RuntimeOptions) (_result *GetEntityRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetEntityRouteResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEntityRoute"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEntityRoute(request *GetEntityRouteRequest) (_result *GetEntityRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEntityRouteResponse{}
	_body, _err := client.GetEntityRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEntityTagRelationWithOptions(request *UpdateEntityTagRelationRequest, runtime *util.RuntimeOptions) (_result *UpdateEntityTagRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateEntityTagRelationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateEntityTagRelation"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEntityTagRelation(request *UpdateEntityTagRelationRequest) (_result *UpdateEntityTagRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEntityTagRelationResponse{}
	_body, _err := client.UpdateEntityTagRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOuterCallCenterDataWithOptions(request *CreateOuterCallCenterDataRequest, runtime *util.RuntimeOptions) (_result *CreateOuterCallCenterDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateOuterCallCenterDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateOuterCallCenterData"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOuterCallCenterData(request *CreateOuterCallCenterDataRequest) (_result *CreateOuterCallCenterDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOuterCallCenterDataResponse{}
	_body, _err := client.CreateOuterCallCenterDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendOutboundCommandWithOptions(request *SendOutboundCommandRequest, runtime *util.RuntimeOptions) (_result *SendOutboundCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendOutboundCommandResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendOutboundCommand"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendOutboundCommand(request *SendOutboundCommandRequest) (_result *SendOutboundCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendOutboundCommandResponse{}
	_body, _err := client.SendOutboundCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRoleWithOptions(request *CreateRoleRequest, runtime *util.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRole"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRole(request *CreateRoleRequest) (_result *CreateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRoleResponse{}
	_body, _err := client.CreateRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSkillGroupWithOptions(request *ListSkillGroupRequest, runtime *util.RuntimeOptions) (_result *ListSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSkillGroup"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSkillGroup(request *ListSkillGroupRequest) (_result *ListSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSkillGroupResponse{}
	_body, _err := client.ListSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantRolesWithOptions(request *GrantRolesRequest, runtime *util.RuntimeOptions) (_result *GrantRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GrantRolesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GrantRoles"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantRoles(request *GrantRolesRequest) (_result *GrantRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantRolesResponse{}
	_body, _err := client.GrantRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOuterCallCenterDataListWithOptions(request *GetOuterCallCenterDataListRequest, runtime *util.RuntimeOptions) (_result *GetOuterCallCenterDataListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOuterCallCenterDataListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOuterCallCenterDataList"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOuterCallCenterDataList(request *GetOuterCallCenterDataListRequest) (_result *GetOuterCallCenterDataListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOuterCallCenterDataListResponse{}
	_body, _err := client.GetOuterCallCenterDataListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTicketsWithOptions(tmpReq *QueryTicketsRequest, runtime *util.RuntimeOptions) (_result *QueryTicketsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryTicketsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extra)) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, tea.String("Extra"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTicketsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTickets"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTickets(request *QueryTicketsRequest) (_result *QueryTicketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTicketsResponse{}
	_body, _err := client.QueryTicketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTicketActionsWithOptions(request *QueryTicketActionsRequest, runtime *util.RuntimeOptions) (_result *QueryTicketActionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryTicketActionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryTicketActions"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTicketActions(request *QueryTicketActionsRequest) (_result *QueryTicketActionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTicketActionsResponse{}
	_body, _err := client.QueryTicketActionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferCallToAgentWithOptions(request *TransferCallToAgentRequest, runtime *util.RuntimeOptions) (_result *TransferCallToAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TransferCallToAgentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TransferCallToAgent"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferCallToAgent(request *TransferCallToAgentRequest) (_result *TransferCallToAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferCallToAgentResponse{}
	_body, _err := client.TransferCallToAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FinishHotlineServiceWithOptions(request *FinishHotlineServiceRequest, runtime *util.RuntimeOptions) (_result *FinishHotlineServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FinishHotlineServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FinishHotlineService"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FinishHotlineService(request *FinishHotlineServiceRequest) (_result *FinishHotlineServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FinishHotlineServiceResponse{}
	_body, _err := client.FinishHotlineServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinThirdCallWithOptions(request *JoinThirdCallRequest, runtime *util.RuntimeOptions) (_result *JoinThirdCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &JoinThirdCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("JoinThirdCall"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinThirdCall(request *JoinThirdCallRequest) (_result *JoinThirdCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinThirdCallResponse{}
	_body, _err := client.JoinThirdCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteActivityWithOptions(request *ExecuteActivityRequest, runtime *util.RuntimeOptions) (_result *ExecuteActivityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteActivityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteActivity"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteActivity(request *ExecuteActivityRequest) (_result *ExecuteActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteActivityResponse{}
	_body, _err := client.ExecuteActivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGrantedRoleIdsWithOptions(request *GetGrantedRoleIdsRequest, runtime *util.RuntimeOptions) (_result *GetGrantedRoleIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetGrantedRoleIdsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetGrantedRoleIds"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGrantedRoleIds(request *GetGrantedRoleIdsRequest) (_result *GetGrantedRoleIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGrantedRoleIdsResponse{}
	_body, _err := client.GetGrantedRoleIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotlineRecordWithOptions(request *ListHotlineRecordRequest, runtime *util.RuntimeOptions) (_result *ListHotlineRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListHotlineRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHotlineRecord"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotlineRecord(request *ListHotlineRecordRequest) (_result *ListHotlineRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotlineRecordResponse{}
	_body, _err := client.ListHotlineRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNumLocationWithOptions(request *GetNumLocationRequest, runtime *util.RuntimeOptions) (_result *GetNumLocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetNumLocationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetNumLocation"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNumLocation(request *GetNumLocationRequest) (_result *GetNumLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNumLocationResponse{}
	_body, _err := client.GetNumLocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSkillGroupWithOptions(request *CreateSkillGroupRequest, runtime *util.RuntimeOptions) (_result *CreateSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSkillGroup"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSkillGroup(request *CreateSkillGroupRequest) (_result *CreateSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSkillGroupResponse{}
	_body, _err := client.CreateSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomerWithOptions(request *CreateCustomerRequest, runtime *util.RuntimeOptions) (_result *CreateCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCustomerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCustomer"), tea.String("2020-07-02"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomer(request *CreateCustomerRequest) (_result *CreateCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomerResponse{}
	_body, _err := client.CreateCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
