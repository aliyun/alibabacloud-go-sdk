// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddSkillGroupsToUserRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SkillLevelList *string `json:"SkillLevelList,omitempty" xml:"SkillLevelList,omitempty"`
}

func (s AddSkillGroupsToUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSkillGroupsToUserRequest) GoString() string {
	return s.String()
}

func (s *AddSkillGroupsToUserRequest) SetInstanceId(v string) *AddSkillGroupsToUserRequest {
	s.InstanceId = &v
	return s
}

func (s *AddSkillGroupsToUserRequest) SetUserId(v string) *AddSkillGroupsToUserRequest {
	s.UserId = &v
	return s
}

func (s *AddSkillGroupsToUserRequest) SetSkillLevelList(v string) *AddSkillGroupsToUserRequest {
	s.SkillLevelList = &v
	return s
}

type AddSkillGroupsToUserResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s AddSkillGroupsToUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSkillGroupsToUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddSkillGroupsToUserResponseBody) SetCode(v string) *AddSkillGroupsToUserResponseBody {
	s.Code = &v
	return s
}

func (s *AddSkillGroupsToUserResponseBody) SetHttpStatusCode(v int32) *AddSkillGroupsToUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddSkillGroupsToUserResponseBody) SetMessage(v string) *AddSkillGroupsToUserResponseBody {
	s.Message = &v
	return s
}

func (s *AddSkillGroupsToUserResponseBody) SetRequestId(v string) *AddSkillGroupsToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSkillGroupsToUserResponseBody) SetParams(v []*string) *AddSkillGroupsToUserResponseBody {
	s.Params = v
	return s
}

type AddSkillGroupsToUserResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSkillGroupsToUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSkillGroupsToUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSkillGroupsToUserResponse) GoString() string {
	return s.String()
}

func (s *AddSkillGroupsToUserResponse) SetHeaders(v map[string]*string) *AddSkillGroupsToUserResponse {
	s.Headers = v
	return s
}

func (s *AddSkillGroupsToUserResponse) SetBody(v *AddSkillGroupsToUserResponseBody) *AddSkillGroupsToUserResponse {
	s.Body = v
	return s
}

type SaveWebRTCStatsRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CallId         *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	GeneralInfo    *string `json:"GeneralInfo,omitempty" xml:"GeneralInfo,omitempty"`
	SenderReport   *string `json:"SenderReport,omitempty" xml:"SenderReport,omitempty"`
	ReceiverReport *string `json:"ReceiverReport,omitempty" xml:"ReceiverReport,omitempty"`
	GoogAddress    *string `json:"GoogAddress,omitempty" xml:"GoogAddress,omitempty"`
}

func (s SaveWebRTCStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveWebRTCStatsRequest) GoString() string {
	return s.String()
}

func (s *SaveWebRTCStatsRequest) SetInstanceId(v string) *SaveWebRTCStatsRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveWebRTCStatsRequest) SetCallId(v string) *SaveWebRTCStatsRequest {
	s.CallId = &v
	return s
}

func (s *SaveWebRTCStatsRequest) SetGeneralInfo(v string) *SaveWebRTCStatsRequest {
	s.GeneralInfo = &v
	return s
}

func (s *SaveWebRTCStatsRequest) SetSenderReport(v string) *SaveWebRTCStatsRequest {
	s.SenderReport = &v
	return s
}

func (s *SaveWebRTCStatsRequest) SetReceiverReport(v string) *SaveWebRTCStatsRequest {
	s.ReceiverReport = &v
	return s
}

func (s *SaveWebRTCStatsRequest) SetGoogAddress(v string) *SaveWebRTCStatsRequest {
	s.GoogAddress = &v
	return s
}

type SaveWebRTCStatsResponseBody struct {
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TimeStamp      *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	RowCount       *int64  `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
}

func (s SaveWebRTCStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveWebRTCStatsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWebRTCStatsResponseBody) SetHttpStatusCode(v int64) *SaveWebRTCStatsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) SetRequestId(v string) *SaveWebRTCStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) SetSuccess(v bool) *SaveWebRTCStatsResponseBody {
	s.Success = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) SetCode(v string) *SaveWebRTCStatsResponseBody {
	s.Code = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) SetMessage(v string) *SaveWebRTCStatsResponseBody {
	s.Message = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) SetTimeStamp(v int64) *SaveWebRTCStatsResponseBody {
	s.TimeStamp = &v
	return s
}

func (s *SaveWebRTCStatsResponseBody) SetRowCount(v int64) *SaveWebRTCStatsResponseBody {
	s.RowCount = &v
	return s
}

type SaveWebRTCStatsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveWebRTCStatsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveWebRTCStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveWebRTCStatsResponse) GoString() string {
	return s.String()
}

func (s *SaveWebRTCStatsResponse) SetHeaders(v map[string]*string) *SaveWebRTCStatsResponse {
	s.Headers = v
	return s
}

func (s *SaveWebRTCStatsResponse) SetBody(v *SaveWebRTCStatsResponseBody) *SaveWebRTCStatsResponse {
	s.Body = v
	return s
}

type GetMonoRecordingRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ContactId  *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s GetMonoRecordingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMonoRecordingRequest) GoString() string {
	return s.String()
}

func (s *GetMonoRecordingRequest) SetInstanceId(v string) *GetMonoRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMonoRecordingRequest) SetContactId(v string) *GetMonoRecordingRequest {
	s.ContactId = &v
	return s
}

type GetMonoRecordingResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetMonoRecordingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetMonoRecordingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMonoRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *GetMonoRecordingResponseBody) SetCode(v string) *GetMonoRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *GetMonoRecordingResponseBody) SetHttpStatusCode(v int32) *GetMonoRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMonoRecordingResponseBody) SetMessage(v string) *GetMonoRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *GetMonoRecordingResponseBody) SetRequestId(v string) *GetMonoRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMonoRecordingResponseBody) SetData(v *GetMonoRecordingResponseBodyData) *GetMonoRecordingResponseBody {
	s.Data = v
	return s
}

type GetMonoRecordingResponseBodyData struct {
	FileUrl  *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetMonoRecordingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMonoRecordingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMonoRecordingResponseBodyData) SetFileUrl(v string) *GetMonoRecordingResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetMonoRecordingResponseBodyData) SetFileName(v string) *GetMonoRecordingResponseBodyData {
	s.FileName = &v
	return s
}

type GetMonoRecordingResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMonoRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMonoRecordingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMonoRecordingResponse) GoString() string {
	return s.String()
}

func (s *GetMonoRecordingResponse) SetHeaders(v map[string]*string) *GetMonoRecordingResponse {
	s.Headers = v
	return s
}

func (s *GetMonoRecordingResponse) SetBody(v *GetMonoRecordingResponseBody) *GetMonoRecordingResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetInstanceId(v string) *ListUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersRequest) SetSearchPattern(v string) *ListUsersRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

type ListUsersResponseBody struct {
	Code           *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                  `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *ListUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetCode(v string) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetHttpStatusCode(v int32) *ListUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetParams(v []*string) *ListUsersResponseBody {
	s.Params = v
	return s
}

func (s *ListUsersResponseBody) SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

type ListUsersResponseBodyData struct {
	PageNumber *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListUsersResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetPageNumber(v int32) *ListUsersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPageSize(v int32) *ListUsersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBodyData) SetTotalCount(v int32) *ListUsersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBodyData) SetList(v []*ListUsersResponseBodyDataList) *ListUsersResponseBodyData {
	s.List = v
	return s
}

type ListUsersResponseBodyDataList struct {
	DisplayName                *string                                                    `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	LoginName                  *string                                                    `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	Email                      *string                                                    `json:"Email,omitempty" xml:"Email,omitempty"`
	WorkMode                   *string                                                    `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	Mobile                     *string                                                    `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	UserId                     *string                                                    `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RoleName                   *string                                                    `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	RoleId                     *string                                                    `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	PrimaryAccount             *bool                                                      `json:"PrimaryAccount,omitempty" xml:"PrimaryAccount,omitempty"`
	PersonalOutboundNumberList []*ListUsersResponseBodyDataListPersonalOutboundNumberList `json:"PersonalOutboundNumberList,omitempty" xml:"PersonalOutboundNumberList,omitempty" type:"Repeated"`
	SkillLevelList             []*ListUsersResponseBodyDataListSkillLevelList             `json:"SkillLevelList,omitempty" xml:"SkillLevelList,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataList) SetDisplayName(v string) *ListUsersResponseBodyDataList {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetLoginName(v string) *ListUsersResponseBodyDataList {
	s.LoginName = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetEmail(v string) *ListUsersResponseBodyDataList {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetWorkMode(v string) *ListUsersResponseBodyDataList {
	s.WorkMode = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetMobile(v string) *ListUsersResponseBodyDataList {
	s.Mobile = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetUserId(v string) *ListUsersResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetRoleName(v string) *ListUsersResponseBodyDataList {
	s.RoleName = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetRoleId(v string) *ListUsersResponseBodyDataList {
	s.RoleId = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetPrimaryAccount(v bool) *ListUsersResponseBodyDataList {
	s.PrimaryAccount = &v
	return s
}

func (s *ListUsersResponseBodyDataList) SetPersonalOutboundNumberList(v []*ListUsersResponseBodyDataListPersonalOutboundNumberList) *ListUsersResponseBodyDataList {
	s.PersonalOutboundNumberList = v
	return s
}

func (s *ListUsersResponseBodyDataList) SetSkillLevelList(v []*ListUsersResponseBodyDataListSkillLevelList) *ListUsersResponseBodyDataList {
	s.SkillLevelList = v
	return s
}

type ListUsersResponseBodyDataListPersonalOutboundNumberList struct {
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Active   *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	Usage    *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListUsersResponseBodyDataListPersonalOutboundNumberList) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataListPersonalOutboundNumberList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) SetNumber(v string) *ListUsersResponseBodyDataListPersonalOutboundNumberList {
	s.Number = &v
	return s
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) SetActive(v bool) *ListUsersResponseBodyDataListPersonalOutboundNumberList {
	s.Active = &v
	return s
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) SetCity(v string) *ListUsersResponseBodyDataListPersonalOutboundNumberList {
	s.City = &v
	return s
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) SetUsage(v string) *ListUsersResponseBodyDataListPersonalOutboundNumberList {
	s.Usage = &v
	return s
}

func (s *ListUsersResponseBodyDataListPersonalOutboundNumberList) SetProvince(v string) *ListUsersResponseBodyDataListPersonalOutboundNumberList {
	s.Province = &v
	return s
}

type ListUsersResponseBodyDataListSkillLevelList struct {
	SkillLevel     *int32  `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
	SkillGroupId   *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s ListUsersResponseBodyDataListSkillLevelList) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataListSkillLevelList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataListSkillLevelList) SetSkillLevel(v int32) *ListUsersResponseBodyDataListSkillLevelList {
	s.SkillLevel = &v
	return s
}

func (s *ListUsersResponseBodyDataListSkillLevelList) SetSkillGroupId(v string) *ListUsersResponseBodyDataListSkillLevelList {
	s.SkillGroupId = &v
	return s
}

func (s *ListUsersResponseBodyDataListSkillLevelList) SetSkillGroupName(v string) *ListUsersResponseBodyDataListSkillLevelList {
	s.SkillGroupName = &v
	return s
}

type ListUsersResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListAgentStateLogsRequest struct {
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	AgentId    *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListAgentStateLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAgentStateLogsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentStateLogsRequest) SetStartTime(v int64) *ListAgentStateLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListAgentStateLogsRequest) SetEndTime(v int64) *ListAgentStateLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAgentStateLogsRequest) SetAgentId(v string) *ListAgentStateLogsRequest {
	s.AgentId = &v
	return s
}

func (s *ListAgentStateLogsRequest) SetInstanceId(v string) *ListAgentStateLogsRequest {
	s.InstanceId = &v
	return s
}

type ListAgentStateLogsResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*ListAgentStateLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListAgentStateLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAgentStateLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentStateLogsResponseBody) SetCode(v string) *ListAgentStateLogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentStateLogsResponseBody) SetHttpStatusCode(v int32) *ListAgentStateLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAgentStateLogsResponseBody) SetMessage(v string) *ListAgentStateLogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentStateLogsResponseBody) SetRequestId(v string) *ListAgentStateLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentStateLogsResponseBody) SetData(v []*ListAgentStateLogsResponseBodyData) *ListAgentStateLogsResponseBody {
	s.Data = v
	return s
}

type ListAgentStateLogsResponseBodyData struct {
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StateCode *string `json:"StateCode,omitempty" xml:"StateCode,omitempty"`
	Duration  *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListAgentStateLogsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAgentStateLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentStateLogsResponseBodyData) SetStartTime(v int64) *ListAgentStateLogsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListAgentStateLogsResponseBodyData) SetStateCode(v string) *ListAgentStateLogsResponseBodyData {
	s.StateCode = &v
	return s
}

func (s *ListAgentStateLogsResponseBodyData) SetDuration(v int64) *ListAgentStateLogsResponseBodyData {
	s.Duration = &v
	return s
}

func (s *ListAgentStateLogsResponseBodyData) SetState(v string) *ListAgentStateLogsResponseBodyData {
	s.State = &v
	return s
}

type ListAgentStateLogsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAgentStateLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAgentStateLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAgentStateLogsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentStateLogsResponse) SetHeaders(v map[string]*string) *ListAgentStateLogsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentStateLogsResponse) SetBody(v *ListAgentStateLogsResponseBody) *ListAgentStateLogsResponse {
	s.Body = v
	return s
}

type RemovePhoneNumberFromSkillGroupsRequest struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Number           *string `json:"Number,omitempty" xml:"Number,omitempty"`
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
}

func (s RemovePhoneNumberFromSkillGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemovePhoneNumberFromSkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumberFromSkillGroupsRequest) SetInstanceId(v string) *RemovePhoneNumberFromSkillGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *RemovePhoneNumberFromSkillGroupsRequest) SetNumber(v string) *RemovePhoneNumberFromSkillGroupsRequest {
	s.Number = &v
	return s
}

func (s *RemovePhoneNumberFromSkillGroupsRequest) SetSkillGroupIdList(v string) *RemovePhoneNumberFromSkillGroupsRequest {
	s.SkillGroupIdList = &v
	return s
}

type RemovePhoneNumberFromSkillGroupsResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePhoneNumberFromSkillGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemovePhoneNumberFromSkillGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumberFromSkillGroupsResponseBody) SetCode(v string) *RemovePhoneNumberFromSkillGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *RemovePhoneNumberFromSkillGroupsResponseBody) SetHttpStatusCode(v int32) *RemovePhoneNumberFromSkillGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemovePhoneNumberFromSkillGroupsResponseBody) SetMessage(v string) *RemovePhoneNumberFromSkillGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *RemovePhoneNumberFromSkillGroupsResponseBody) SetRequestId(v string) *RemovePhoneNumberFromSkillGroupsResponseBody {
	s.RequestId = &v
	return s
}

type RemovePhoneNumberFromSkillGroupsResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemovePhoneNumberFromSkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemovePhoneNumberFromSkillGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemovePhoneNumberFromSkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumberFromSkillGroupsResponse) SetHeaders(v map[string]*string) *RemovePhoneNumberFromSkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *RemovePhoneNumberFromSkillGroupsResponse) SetBody(v *RemovePhoneNumberFromSkillGroupsResponseBody) *RemovePhoneNumberFromSkillGroupsResponse {
	s.Body = v
	return s
}

type ListPhoneNumbersRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	Usage         *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	Active        *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPhoneNumbersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneNumbersRequest) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersRequest) SetInstanceId(v string) *ListPhoneNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListPhoneNumbersRequest) SetSearchPattern(v string) *ListPhoneNumbersRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListPhoneNumbersRequest) SetUsage(v string) *ListPhoneNumbersRequest {
	s.Usage = &v
	return s
}

func (s *ListPhoneNumbersRequest) SetActive(v bool) *ListPhoneNumbersRequest {
	s.Active = &v
	return s
}

func (s *ListPhoneNumbersRequest) SetPageNumber(v int32) *ListPhoneNumbersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPhoneNumbersRequest) SetPageSize(v int32) *ListPhoneNumbersRequest {
	s.PageSize = &v
	return s
}

type ListPhoneNumbersResponseBody struct {
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Data           *ListPhoneNumbersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListPhoneNumbersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersResponseBody) SetHttpStatusCode(v int32) *ListPhoneNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPhoneNumbersResponseBody) SetRequestId(v string) *ListPhoneNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPhoneNumbersResponseBody) SetCode(v string) *ListPhoneNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ListPhoneNumbersResponseBody) SetMessage(v string) *ListPhoneNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ListPhoneNumbersResponseBody) SetPageNumber(v int32) *ListPhoneNumbersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPhoneNumbersResponseBody) SetPageSize(v int32) *ListPhoneNumbersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPhoneNumbersResponseBody) SetData(v *ListPhoneNumbersResponseBodyData) *ListPhoneNumbersResponseBody {
	s.Data = v
	return s
}

type ListPhoneNumbersResponseBodyData struct {
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListPhoneNumbersResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListPhoneNumbersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneNumbersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersResponseBodyData) SetPageNumber(v int32) *ListPhoneNumbersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyData) SetPageSize(v int32) *ListPhoneNumbersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyData) SetTotalCount(v int32) *ListPhoneNumbersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyData) SetList(v []*ListPhoneNumbersResponseBodyDataList) *ListPhoneNumbersResponseBodyData {
	s.List = v
	return s
}

type ListPhoneNumbersResponseBodyDataList struct {
	Active          *bool                                              `json:"Active,omitempty" xml:"Active,omitempty"`
	CreateTime      *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId          *string                                            `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Tags            *string                                            `json:"Tags,omitempty" xml:"Tags,omitempty"`
	City            *string                                            `json:"City,omitempty" xml:"City,omitempty"`
	InstanceId      *string                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Usage           *string                                            `json:"Usage,omitempty" xml:"Usage,omitempty"`
	ContactFlowName *string                                            `json:"ContactFlowName,omitempty" xml:"ContactFlowName,omitempty"`
	Provider        *string                                            `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Number          *string                                            `json:"Number,omitempty" xml:"Number,omitempty"`
	ContactFlowId   *string                                            `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	Province        *string                                            `json:"Province,omitempty" xml:"Province,omitempty"`
	SkillGroups     []*ListPhoneNumbersResponseBodyDataListSkillGroups `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty" type:"Repeated"`
}

func (s ListPhoneNumbersResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneNumbersResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersResponseBodyDataList) SetActive(v bool) *ListPhoneNumbersResponseBodyDataList {
	s.Active = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetCreateTime(v string) *ListPhoneNumbersResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetUserId(v string) *ListPhoneNumbersResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetTags(v string) *ListPhoneNumbersResponseBodyDataList {
	s.Tags = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetCity(v string) *ListPhoneNumbersResponseBodyDataList {
	s.City = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetInstanceId(v string) *ListPhoneNumbersResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetUsage(v string) *ListPhoneNumbersResponseBodyDataList {
	s.Usage = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetContactFlowName(v string) *ListPhoneNumbersResponseBodyDataList {
	s.ContactFlowName = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetProvider(v string) *ListPhoneNumbersResponseBodyDataList {
	s.Provider = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetNumber(v string) *ListPhoneNumbersResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetContactFlowId(v string) *ListPhoneNumbersResponseBodyDataList {
	s.ContactFlowId = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetProvince(v string) *ListPhoneNumbersResponseBodyDataList {
	s.Province = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataList) SetSkillGroups(v []*ListPhoneNumbersResponseBodyDataListSkillGroups) *ListPhoneNumbersResponseBodyDataList {
	s.SkillGroups = v
	return s
}

type ListPhoneNumbersResponseBodyDataListSkillGroups struct {
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListPhoneNumbersResponseBodyDataListSkillGroups) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneNumbersResponseBodyDataListSkillGroups) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) SetDisplayName(v string) *ListPhoneNumbersResponseBodyDataListSkillGroups {
	s.DisplayName = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) SetInstanceId(v string) *ListPhoneNumbersResponseBodyDataListSkillGroups {
	s.InstanceId = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) SetName(v string) *ListPhoneNumbersResponseBodyDataListSkillGroups {
	s.Name = &v
	return s
}

func (s *ListPhoneNumbersResponseBodyDataListSkillGroups) SetSkillGroupId(v string) *ListPhoneNumbersResponseBodyDataListSkillGroups {
	s.SkillGroupId = &v
	return s
}

type ListPhoneNumbersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPhoneNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPhoneNumbersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneNumbersResponse) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersResponse) SetHeaders(v map[string]*string) *ListPhoneNumbersResponse {
	s.Headers = v
	return s
}

func (s *ListPhoneNumbersResponse) SetBody(v *ListPhoneNumbersResponseBody) *ListPhoneNumbersResponse {
	s.Body = v
	return s
}

type AddNumbersToSkillGroupRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	NumberList   *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
}

func (s AddNumbersToSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddNumbersToSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *AddNumbersToSkillGroupRequest) SetInstanceId(v string) *AddNumbersToSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AddNumbersToSkillGroupRequest) SetSkillGroupId(v string) *AddNumbersToSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *AddNumbersToSkillGroupRequest) SetNumberList(v string) *AddNumbersToSkillGroupRequest {
	s.NumberList = &v
	return s
}

type AddNumbersToSkillGroupResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddNumbersToSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddNumbersToSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddNumbersToSkillGroupResponseBody) SetCode(v string) *AddNumbersToSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddNumbersToSkillGroupResponseBody) SetHttpStatusCode(v int32) *AddNumbersToSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddNumbersToSkillGroupResponseBody) SetMessage(v string) *AddNumbersToSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AddNumbersToSkillGroupResponseBody) SetRequestId(v string) *AddNumbersToSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddNumbersToSkillGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddNumbersToSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddNumbersToSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddNumbersToSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *AddNumbersToSkillGroupResponse) SetHeaders(v map[string]*string) *AddNumbersToSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *AddNumbersToSkillGroupResponse) SetBody(v *AddNumbersToSkillGroupResponseBody) *AddNumbersToSkillGroupResponse {
	s.Body = v
	return s
}

type ResetAgentStateRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s ResetAgentStateRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAgentStateRequest) GoString() string {
	return s.String()
}

func (s *ResetAgentStateRequest) SetInstanceId(v string) *ResetAgentStateRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetAgentStateRequest) SetUserId(v string) *ResetAgentStateRequest {
	s.UserId = &v
	return s
}

func (s *ResetAgentStateRequest) SetDeviceId(v string) *ResetAgentStateRequest {
	s.DeviceId = &v
	return s
}

type ResetAgentStateResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                        `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *ResetAgentStateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ResetAgentStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAgentStateResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAgentStateResponseBody) SetCode(v string) *ResetAgentStateResponseBody {
	s.Code = &v
	return s
}

func (s *ResetAgentStateResponseBody) SetHttpStatusCode(v int32) *ResetAgentStateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetAgentStateResponseBody) SetMessage(v string) *ResetAgentStateResponseBody {
	s.Message = &v
	return s
}

func (s *ResetAgentStateResponseBody) SetRequestId(v string) *ResetAgentStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAgentStateResponseBody) SetParams(v []*string) *ResetAgentStateResponseBody {
	s.Params = v
	return s
}

func (s *ResetAgentStateResponseBody) SetData(v *ResetAgentStateResponseBodyData) *ResetAgentStateResponseBody {
	s.Data = v
	return s
}

type ResetAgentStateResponseBodyData struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s ResetAgentStateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ResetAgentStateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetAgentStateResponseBodyData) SetExtension(v string) *ResetAgentStateResponseBodyData {
	s.Extension = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetWorkMode(v string) *ResetAgentStateResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetDeviceId(v string) *ResetAgentStateResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetJobId(v string) *ResetAgentStateResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetUserId(v string) *ResetAgentStateResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetBreakCode(v string) *ResetAgentStateResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetInstanceId(v string) *ResetAgentStateResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetOutboundScenario(v bool) *ResetAgentStateResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetUserState(v string) *ResetAgentStateResponseBodyData {
	s.UserState = &v
	return s
}

func (s *ResetAgentStateResponseBodyData) SetSignedSkillGroupIdList(v []*string) *ResetAgentStateResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

type ResetAgentStateResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetAgentStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetAgentStateResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAgentStateResponse) GoString() string {
	return s.String()
}

func (s *ResetAgentStateResponse) SetHeaders(v map[string]*string) *ResetAgentStateResponse {
	s.Headers = v
	return s
}

func (s *ResetAgentStateResponse) SetBody(v *ResetAgentStateResponseBody) *ResetAgentStateResponse {
	s.Body = v
	return s
}

type ChangeWorkModeRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	WorkMode   *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s ChangeWorkModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeWorkModeRequest) GoString() string {
	return s.String()
}

func (s *ChangeWorkModeRequest) SetInstanceId(v string) *ChangeWorkModeRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeWorkModeRequest) SetUserId(v string) *ChangeWorkModeRequest {
	s.UserId = &v
	return s
}

func (s *ChangeWorkModeRequest) SetDeviceId(v string) *ChangeWorkModeRequest {
	s.DeviceId = &v
	return s
}

func (s *ChangeWorkModeRequest) SetWorkMode(v string) *ChangeWorkModeRequest {
	s.WorkMode = &v
	return s
}

type ChangeWorkModeResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                       `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *ChangeWorkModeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ChangeWorkModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeWorkModeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeWorkModeResponseBody) SetCode(v string) *ChangeWorkModeResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeWorkModeResponseBody) SetHttpStatusCode(v int32) *ChangeWorkModeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChangeWorkModeResponseBody) SetMessage(v string) *ChangeWorkModeResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeWorkModeResponseBody) SetRequestId(v string) *ChangeWorkModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeWorkModeResponseBody) SetParams(v []*string) *ChangeWorkModeResponseBody {
	s.Params = v
	return s
}

func (s *ChangeWorkModeResponseBody) SetData(v *ChangeWorkModeResponseBodyData) *ChangeWorkModeResponseBody {
	s.Data = v
	return s
}

type ChangeWorkModeResponseBodyData struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s ChangeWorkModeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ChangeWorkModeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeWorkModeResponseBodyData) SetExtension(v string) *ChangeWorkModeResponseBodyData {
	s.Extension = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetWorkMode(v string) *ChangeWorkModeResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetDeviceId(v string) *ChangeWorkModeResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetJobId(v string) *ChangeWorkModeResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetUserId(v string) *ChangeWorkModeResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetBreakCode(v string) *ChangeWorkModeResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetInstanceId(v string) *ChangeWorkModeResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetOutboundScenario(v bool) *ChangeWorkModeResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetUserState(v string) *ChangeWorkModeResponseBodyData {
	s.UserState = &v
	return s
}

func (s *ChangeWorkModeResponseBodyData) SetSignedSkillGroupIdList(v []*string) *ChangeWorkModeResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

type ChangeWorkModeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ChangeWorkModeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeWorkModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeWorkModeResponse) GoString() string {
	return s.String()
}

func (s *ChangeWorkModeResponse) SetHeaders(v map[string]*string) *ChangeWorkModeResponse {
	s.Headers = v
	return s
}

func (s *ChangeWorkModeResponse) SetBody(v *ChangeWorkModeResponseBody) *ChangeWorkModeResponse {
	s.Body = v
	return s
}

type GetTurnCredentialsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetTurnCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTurnCredentialsRequest) GoString() string {
	return s.String()
}

func (s *GetTurnCredentialsRequest) SetInstanceId(v string) *GetTurnCredentialsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTurnCredentialsRequest) SetUserId(v string) *GetTurnCredentialsRequest {
	s.UserId = &v
	return s
}

type GetTurnCredentialsResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                           `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *GetTurnCredentialsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetTurnCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTurnCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTurnCredentialsResponseBody) SetCode(v string) *GetTurnCredentialsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTurnCredentialsResponseBody) SetHttpStatusCode(v int32) *GetTurnCredentialsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTurnCredentialsResponseBody) SetMessage(v string) *GetTurnCredentialsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTurnCredentialsResponseBody) SetRequestId(v string) *GetTurnCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTurnCredentialsResponseBody) SetParams(v []*string) *GetTurnCredentialsResponseBody {
	s.Params = v
	return s
}

func (s *GetTurnCredentialsResponseBody) SetData(v *GetTurnCredentialsResponseBodyData) *GetTurnCredentialsResponseBody {
	s.Data = v
	return s
}

type GetTurnCredentialsResponseBodyData struct {
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s GetTurnCredentialsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTurnCredentialsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTurnCredentialsResponseBodyData) SetUserName(v string) *GetTurnCredentialsResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetTurnCredentialsResponseBodyData) SetPassword(v string) *GetTurnCredentialsResponseBodyData {
	s.Password = &v
	return s
}

type GetTurnCredentialsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTurnCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTurnCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTurnCredentialsResponse) GoString() string {
	return s.String()
}

func (s *GetTurnCredentialsResponse) SetHeaders(v map[string]*string) *GetTurnCredentialsResponse {
	s.Headers = v
	return s
}

func (s *GetTurnCredentialsResponse) SetBody(v *GetTurnCredentialsResponseBody) *GetTurnCredentialsResponse {
	s.Body = v
	return s
}

type AddPhoneNumbersRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	Usage         *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	NumberList    *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	NumberGroupId *string `json:"NumberGroupId,omitempty" xml:"NumberGroupId,omitempty"`
}

func (s AddPhoneNumbersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPhoneNumbersRequest) GoString() string {
	return s.String()
}

func (s *AddPhoneNumbersRequest) SetInstanceId(v string) *AddPhoneNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *AddPhoneNumbersRequest) SetContactFlowId(v string) *AddPhoneNumbersRequest {
	s.ContactFlowId = &v
	return s
}

func (s *AddPhoneNumbersRequest) SetUsage(v string) *AddPhoneNumbersRequest {
	s.Usage = &v
	return s
}

func (s *AddPhoneNumbersRequest) SetNumberList(v string) *AddPhoneNumbersRequest {
	s.NumberList = &v
	return s
}

func (s *AddPhoneNumbersRequest) SetNumberGroupId(v string) *AddPhoneNumbersRequest {
	s.NumberGroupId = &v
	return s
}

type AddPhoneNumbersResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s AddPhoneNumbersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddPhoneNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *AddPhoneNumbersResponseBody) SetCode(v string) *AddPhoneNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *AddPhoneNumbersResponseBody) SetHttpStatusCode(v int32) *AddPhoneNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddPhoneNumbersResponseBody) SetMessage(v string) *AddPhoneNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *AddPhoneNumbersResponseBody) SetRequestId(v string) *AddPhoneNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPhoneNumbersResponseBody) SetData(v []*string) *AddPhoneNumbersResponseBody {
	s.Data = v
	return s
}

type AddPhoneNumbersResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddPhoneNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddPhoneNumbersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPhoneNumbersResponse) GoString() string {
	return s.String()
}

func (s *AddPhoneNumbersResponse) SetHeaders(v map[string]*string) *AddPhoneNumbersResponse {
	s.Headers = v
	return s
}

func (s *AddPhoneNumbersResponse) SetBody(v *AddPhoneNumbersResponseBody) *AddPhoneNumbersResponse {
	s.Body = v
	return s
}

type SaveWebRtcInfoRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CallId      *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s SaveWebRtcInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveWebRtcInfoRequest) GoString() string {
	return s.String()
}

func (s *SaveWebRtcInfoRequest) SetInstanceId(v string) *SaveWebRtcInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveWebRtcInfoRequest) SetCallId(v string) *SaveWebRtcInfoRequest {
	s.CallId = &v
	return s
}

func (s *SaveWebRtcInfoRequest) SetJobId(v string) *SaveWebRtcInfoRequest {
	s.JobId = &v
	return s
}

func (s *SaveWebRtcInfoRequest) SetContentType(v string) *SaveWebRtcInfoRequest {
	s.ContentType = &v
	return s
}

func (s *SaveWebRtcInfoRequest) SetContent(v string) *SaveWebRtcInfoRequest {
	s.Content = &v
	return s
}

type SaveWebRtcInfoResponseBody struct {
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TimeStamp      *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	RowCount       *int64  `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
}

func (s SaveWebRtcInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveWebRtcInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWebRtcInfoResponseBody) SetHttpStatusCode(v int64) *SaveWebRtcInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetRequestId(v string) *SaveWebRtcInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetSuccess(v bool) *SaveWebRtcInfoResponseBody {
	s.Success = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetCode(v string) *SaveWebRtcInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetMessage(v string) *SaveWebRtcInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetTimeStamp(v int64) *SaveWebRtcInfoResponseBody {
	s.TimeStamp = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetRowCount(v int64) *SaveWebRtcInfoResponseBody {
	s.RowCount = &v
	return s
}

type SaveWebRtcInfoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveWebRtcInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveWebRtcInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveWebRtcInfoResponse) GoString() string {
	return s.String()
}

func (s *SaveWebRtcInfoResponse) SetHeaders(v map[string]*string) *SaveWebRtcInfoResponse {
	s.Headers = v
	return s
}

func (s *SaveWebRtcInfoResponse) SetBody(v *SaveWebRtcInfoResponseBody) *SaveWebRtcInfoResponse {
	s.Body = v
	return s
}

type ListIntervalSkillGroupReportRequest struct {
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval     *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s ListIntervalSkillGroupReportRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalSkillGroupReportRequest) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportRequest) SetSkillGroupId(v string) *ListIntervalSkillGroupReportRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListIntervalSkillGroupReportRequest) SetInstanceId(v string) *ListIntervalSkillGroupReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntervalSkillGroupReportRequest) SetStartTime(v int64) *ListIntervalSkillGroupReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportRequest) SetEndTime(v int64) *ListIntervalSkillGroupReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportRequest) SetInterval(v string) *ListIntervalSkillGroupReportRequest {
	s.Interval = &v
	return s
}

type ListIntervalSkillGroupReportResponseBody struct {
	Code           *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*ListIntervalSkillGroupReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListIntervalSkillGroupReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponseBody) SetCode(v string) *ListIntervalSkillGroupReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBody) SetHttpStatusCode(v int32) *ListIntervalSkillGroupReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBody) SetMessage(v string) *ListIntervalSkillGroupReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBody) SetRequestId(v string) *ListIntervalSkillGroupReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBody) SetData(v []*ListIntervalSkillGroupReportResponseBodyData) *ListIntervalSkillGroupReportResponseBody {
	s.Data = v
	return s
}

type ListIntervalSkillGroupReportResponseBodyData struct {
	StatsTime *int64                                                `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	Inbound   *ListIntervalSkillGroupReportResponseBodyDataInbound  `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Outbound  *ListIntervalSkillGroupReportResponseBodyDataOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall   *ListIntervalSkillGroupReportResponseBodyDataOverall  `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
}

func (s ListIntervalSkillGroupReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponseBodyData) SetStatsTime(v int64) *ListIntervalSkillGroupReportResponseBodyData {
	s.StatsTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyData) SetInbound(v *ListIntervalSkillGroupReportResponseBodyDataInbound) *ListIntervalSkillGroupReportResponseBodyData {
	s.Inbound = v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyData) SetOutbound(v *ListIntervalSkillGroupReportResponseBodyDataOutbound) *ListIntervalSkillGroupReportResponseBodyData {
	s.Outbound = v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyData) SetOverall(v *ListIntervalSkillGroupReportResponseBodyDataOverall) *ListIntervalSkillGroupReportResponseBodyData {
	s.Overall = v
	return s
}

type ListIntervalSkillGroupReportResponseBodyDataInbound struct {
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	CallsOverflow                *int64   `json:"CallsOverflow,omitempty" xml:"CallsOverflow,omitempty"`
	CallsAbandonedInRing         *int64   `json:"CallsAbandonedInRing,omitempty" xml:"CallsAbandonedInRing,omitempty"`
	CallsHandled                 *int64   `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	TotalAbandonedInRingTime     *int64   `json:"TotalAbandonedInRingTime,omitempty" xml:"TotalAbandonedInRingTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	CallsAttendedTransferOut     *int64   `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	AverageWaitTime              *float32 `json:"AverageWaitTime,omitempty" xml:"AverageWaitTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	MaxAbandonTime               *int64   `json:"MaxAbandonTime,omitempty" xml:"MaxAbandonTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsQueued                  *int64   `json:"CallsQueued,omitempty" xml:"CallsQueued,omitempty"`
	CallsBlindTransferIn         *int64   `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	AverageAbandonedInRingTime   *float32 `json:"AverageAbandonedInRingTime,omitempty" xml:"AverageAbandonedInRingTime,omitempty"`
	AverageAbandonTime           *float32 `json:"AverageAbandonTime,omitempty" xml:"AverageAbandonTime,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	CallsBlindTransferOut        *int64   `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	CallsAttendedTransferIn      *int64   `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	CallsAbandoned               *int64   `json:"CallsAbandoned,omitempty" xml:"CallsAbandoned,omitempty"`
	MaxAbandonedInQueueTime      *int64   `json:"MaxAbandonedInQueueTime,omitempty" xml:"MaxAbandonedInQueueTime,omitempty"`
	TotalWaitTime                *int64   `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	AbandonRate                  *float32 `json:"AbandonRate,omitempty" xml:"AbandonRate,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalAbandonTime             *int64   `json:"TotalAbandonTime,omitempty" xml:"TotalAbandonTime,omitempty"`
	CallsOffered                 *int64   `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	MaxAbandonedInRingTime       *int64   `json:"MaxAbandonedInRingTime,omitempty" xml:"MaxAbandonedInRingTime,omitempty"`
	MaxWaitTime                  *int64   `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	AverageAbandonedInQueueTime  *float32 `json:"AverageAbandonedInQueueTime,omitempty" xml:"AverageAbandonedInQueueTime,omitempty"`
	ServiceLevel20               *float32 `json:"ServiceLevel20,omitempty" xml:"ServiceLevel20,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	HandleRate                   *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	CallsTimeout                 *int64   `json:"CallsTimeout,omitempty" xml:"CallsTimeout,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	TotalAbandonedInQueueTime    *int64   `json:"TotalAbandonedInQueueTime,omitempty" xml:"TotalAbandonedInQueueTime,omitempty"`
	CallsAbandonedInQueue        *int64   `json:"CallsAbandonedInQueue,omitempty" xml:"CallsAbandonedInQueue,omitempty"`
}

func (s ListIntervalSkillGroupReportResponseBodyDataInbound) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponseBodyDataInbound) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageRingTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsOverflow(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsOverflow = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsAbandonedInRing(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsAbandonedInRing = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsHandled(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalWorkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalAbandonedInRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalAbandonedInRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxWorkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsAttendedTransferOut(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageWaitTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageWaitTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalHoldTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxAbandonTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxAbandonTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageWorkTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsQueued(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsQueued = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsBlindTransferIn(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetSatisfactionIndex(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageAbandonedInRingTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageAbandonedInRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageAbandonTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageAbandonTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsRinged(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsBlindTransferOut(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsAttendedTransferIn(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsAbandoned(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsAbandoned = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxAbandonedInQueueTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxAbandonedInQueueTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalWaitTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalWaitTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAbandonRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AbandonRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalAbandonTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalAbandonTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsOffered(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxAbandonedInRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxAbandonedInRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxWaitTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxWaitTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageAbandonedInQueueTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageAbandonedInQueueTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetServiceLevel20(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.ServiceLevel20 = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetMaxHoldTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetSatisfactionRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageTalkTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsHold(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetHandleRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.HandleRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsTimeout(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsTimeout = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetAverageHoldTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetTotalAbandonedInQueueTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.TotalAbandonedInQueueTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataInbound) SetCallsAbandonedInQueue(v int64) *ListIntervalSkillGroupReportResponseBodyDataInbound {
	s.CallsAbandonedInQueue = &v
	return s
}

type ListIntervalSkillGroupReportResponseBodyDataOutbound struct {
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	CallsDialed                  *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsAnswered                *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	CallsAttendedTransferOut     *int64   `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	TotalDialingTime             *int64   `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsBlindTransferIn         *int64   `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	CallsAttendedTransferIn      *int64   `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	CallsBlindTransferOut        *int64   `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	MaxDialingTime               *int64   `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	AnswerRate                   *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageDialingTime           *float32 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
}

func (s ListIntervalSkillGroupReportResponseBodyDataOutbound) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponseBodyDataOutbound) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetAverageRingTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsDialed(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsAnswered(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetTotalWorkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsAttendedTransferOut(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetMaxWorkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetTotalDialingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetTotalHoldTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetAverageWorkTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsBlindTransferIn(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetSatisfactionIndex(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsRinged(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsAttendedTransferIn(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsBlindTransferOut(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetTotalRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetMaxTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetMaxRingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetTotalTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetMaxDialingTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetAnswerRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetMaxHoldTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetAverageTalkTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetSatisfactionRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetCallsHold(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetAverageHoldTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOutbound) SetAverageDialingTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOutbound {
	s.AverageDialingTime = &v
	return s
}

type ListIntervalSkillGroupReportResponseBodyDataOverall struct {
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalLoggedInTime            *int64   `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	OccupancyRate                *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	AverageBreakTime             *float32 `json:"AverageBreakTime,omitempty" xml:"AverageBreakTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	MaxBreakTime                 *int64   `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	MaxReadyTime                 *int64   `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	AverageReadyTime             *float32 `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	TotalReadyTime               *int64   `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	TotalBreakTime               *int64   `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalCalls                   *int64   `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
}

func (s ListIntervalSkillGroupReportResponseBodyDataOverall) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponseBodyDataOverall) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalLoggedInTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetOccupancyRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalWorkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetMaxHoldTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetMaxWorkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetAverageBreakTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalHoldTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetSatisfactionRate(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetMaxBreakTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetAverageWorkTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetAverageTalkTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetSatisfactionIndex(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetSatisfactionSurveysOffered(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetSatisfactionSurveysResponded(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetMaxReadyTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetAverageReadyTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetAverageHoldTime(v float32) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalReadyTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalBreakTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetMaxTalkTime(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponseBodyDataOverall) SetTotalCalls(v int64) *ListIntervalSkillGroupReportResponseBodyDataOverall {
	s.TotalCalls = &v
	return s
}

type ListIntervalSkillGroupReportResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListIntervalSkillGroupReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIntervalSkillGroupReportResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponse) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponse) SetHeaders(v map[string]*string) *ListIntervalSkillGroupReportResponse {
	s.Headers = v
	return s
}

func (s *ListIntervalSkillGroupReportResponse) SetBody(v *ListIntervalSkillGroupReportResponseBody) *ListIntervalSkillGroupReportResponse {
	s.Body = v
	return s
}

type MonitorCallRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId        *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	MonitoredUserId *string `json:"MonitoredUserId,omitempty" xml:"MonitoredUserId,omitempty"`
	TimeoutSeconds  *int32  `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s MonitorCallRequest) String() string {
	return tea.Prettify(s)
}

func (s MonitorCallRequest) GoString() string {
	return s.String()
}

func (s *MonitorCallRequest) SetInstanceId(v string) *MonitorCallRequest {
	s.InstanceId = &v
	return s
}

func (s *MonitorCallRequest) SetUserId(v string) *MonitorCallRequest {
	s.UserId = &v
	return s
}

func (s *MonitorCallRequest) SetDeviceId(v string) *MonitorCallRequest {
	s.DeviceId = &v
	return s
}

func (s *MonitorCallRequest) SetMonitoredUserId(v string) *MonitorCallRequest {
	s.MonitoredUserId = &v
	return s
}

func (s *MonitorCallRequest) SetTimeoutSeconds(v int32) *MonitorCallRequest {
	s.TimeoutSeconds = &v
	return s
}

type MonitorCallResponseBody struct {
	Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                    `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *MonitorCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s MonitorCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MonitorCallResponseBody) GoString() string {
	return s.String()
}

func (s *MonitorCallResponseBody) SetCode(v string) *MonitorCallResponseBody {
	s.Code = &v
	return s
}

func (s *MonitorCallResponseBody) SetHttpStatusCode(v int32) *MonitorCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MonitorCallResponseBody) SetMessage(v string) *MonitorCallResponseBody {
	s.Message = &v
	return s
}

func (s *MonitorCallResponseBody) SetRequestId(v string) *MonitorCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *MonitorCallResponseBody) SetParams(v []*string) *MonitorCallResponseBody {
	s.Params = v
	return s
}

func (s *MonitorCallResponseBody) SetData(v *MonitorCallResponseBodyData) *MonitorCallResponseBody {
	s.Data = v
	return s
}

type MonitorCallResponseBodyData struct {
	CallContext *MonitorCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *MonitorCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s MonitorCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MonitorCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *MonitorCallResponseBodyData) SetCallContext(v *MonitorCallResponseBodyDataCallContext) *MonitorCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *MonitorCallResponseBodyData) SetUserContext(v *MonitorCallResponseBodyDataUserContext) *MonitorCallResponseBodyData {
	s.UserContext = v
	return s
}

type MonitorCallResponseBodyDataCallContext struct {
	CallType        *string                                                  `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*MonitorCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s MonitorCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s MonitorCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *MonitorCallResponseBodyDataCallContext) SetCallType(v string) *MonitorCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContext) SetInstanceId(v string) *MonitorCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContext) SetJobId(v string) *MonitorCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContext) SetChannelContexts(v []*MonitorCallResponseBodyDataCallContextChannelContexts) *MonitorCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type MonitorCallResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s MonitorCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s MonitorCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *MonitorCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *MonitorCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type MonitorCallResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s MonitorCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s MonitorCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *MonitorCallResponseBodyDataUserContext) SetExtension(v string) *MonitorCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetHeartbeat(v int64) *MonitorCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetWorkMode(v string) *MonitorCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetDeviceId(v string) *MonitorCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetUserId(v string) *MonitorCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetReserved(v int64) *MonitorCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetBreakCode(v string) *MonitorCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetInstanceId(v string) *MonitorCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *MonitorCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetMobile(v string) *MonitorCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetJobId(v string) *MonitorCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetUserState(v string) *MonitorCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *MonitorCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *MonitorCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type MonitorCallResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MonitorCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MonitorCallResponse) String() string {
	return tea.Prettify(s)
}

func (s MonitorCallResponse) GoString() string {
	return s.String()
}

func (s *MonitorCallResponse) SetHeaders(v map[string]*string) *MonitorCallResponse {
	s.Headers = v
	return s
}

func (s *MonitorCallResponse) SetBody(v *MonitorCallResponseBody) *MonitorCallResponse {
	s.Body = v
	return s
}

type RemoveUsersFromSkillGroupRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	UserIdList   *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s RemoveUsersFromSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromSkillGroupRequest) SetInstanceId(v string) *RemoveUsersFromSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveUsersFromSkillGroupRequest) SetSkillGroupId(v string) *RemoveUsersFromSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *RemoveUsersFromSkillGroupRequest) SetUserIdList(v string) *RemoveUsersFromSkillGroupRequest {
	s.UserIdList = &v
	return s
}

type RemoveUsersFromSkillGroupResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUsersFromSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromSkillGroupResponseBody) SetCode(v string) *RemoveUsersFromSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveUsersFromSkillGroupResponseBody) SetHttpStatusCode(v int32) *RemoveUsersFromSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveUsersFromSkillGroupResponseBody) SetMessage(v string) *RemoveUsersFromSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveUsersFromSkillGroupResponseBody) SetRequestId(v string) *RemoveUsersFromSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveUsersFromSkillGroupResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveUsersFromSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveUsersFromSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromSkillGroupResponse) SetHeaders(v map[string]*string) *RemoveUsersFromSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUsersFromSkillGroupResponse) SetBody(v *RemoveUsersFromSkillGroupResponseBody) *RemoveUsersFromSkillGroupResponse {
	s.Body = v
	return s
}

type DeleteSkillGroupRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Force        *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s DeleteSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupRequest) SetInstanceId(v string) *DeleteSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSkillGroupRequest) SetSkillGroupId(v string) *DeleteSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *DeleteSkillGroupRequest) SetForce(v bool) *DeleteSkillGroupRequest {
	s.Force = &v
	return s
}

type DeleteSkillGroupResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupResponseBody) SetCode(v string) *DeleteSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) SetHttpStatusCode(v int32) *DeleteSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) SetMessage(v string) *DeleteSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSkillGroupResponseBody) SetRequestId(v string) *DeleteSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSkillGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupResponse) SetHeaders(v map[string]*string) *DeleteSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSkillGroupResponse) SetBody(v *DeleteSkillGroupResponseBody) *DeleteSkillGroupResponse {
	s.Body = v
	return s
}

type BlindTransferRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId       *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Transferor     *string `json:"Transferor,omitempty" xml:"Transferor,omitempty"`
	Transferee     *string `json:"Transferee,omitempty" xml:"Transferee,omitempty"`
	TimeoutSeconds *int32  `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s BlindTransferRequest) String() string {
	return tea.Prettify(s)
}

func (s BlindTransferRequest) GoString() string {
	return s.String()
}

func (s *BlindTransferRequest) SetInstanceId(v string) *BlindTransferRequest {
	s.InstanceId = &v
	return s
}

func (s *BlindTransferRequest) SetUserId(v string) *BlindTransferRequest {
	s.UserId = &v
	return s
}

func (s *BlindTransferRequest) SetDeviceId(v string) *BlindTransferRequest {
	s.DeviceId = &v
	return s
}

func (s *BlindTransferRequest) SetTransferor(v string) *BlindTransferRequest {
	s.Transferor = &v
	return s
}

func (s *BlindTransferRequest) SetTransferee(v string) *BlindTransferRequest {
	s.Transferee = &v
	return s
}

func (s *BlindTransferRequest) SetTimeoutSeconds(v int32) *BlindTransferRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *BlindTransferRequest) SetJobId(v string) *BlindTransferRequest {
	s.JobId = &v
	return s
}

type BlindTransferResponseBody struct {
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                      `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *BlindTransferResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s BlindTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BlindTransferResponseBody) GoString() string {
	return s.String()
}

func (s *BlindTransferResponseBody) SetCode(v string) *BlindTransferResponseBody {
	s.Code = &v
	return s
}

func (s *BlindTransferResponseBody) SetHttpStatusCode(v int32) *BlindTransferResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BlindTransferResponseBody) SetMessage(v string) *BlindTransferResponseBody {
	s.Message = &v
	return s
}

func (s *BlindTransferResponseBody) SetRequestId(v string) *BlindTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *BlindTransferResponseBody) SetParams(v []*string) *BlindTransferResponseBody {
	s.Params = v
	return s
}

func (s *BlindTransferResponseBody) SetData(v *BlindTransferResponseBodyData) *BlindTransferResponseBody {
	s.Data = v
	return s
}

type BlindTransferResponseBodyData struct {
	CallContext *BlindTransferResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *BlindTransferResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s BlindTransferResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BlindTransferResponseBodyData) GoString() string {
	return s.String()
}

func (s *BlindTransferResponseBodyData) SetCallContext(v *BlindTransferResponseBodyDataCallContext) *BlindTransferResponseBodyData {
	s.CallContext = v
	return s
}

func (s *BlindTransferResponseBodyData) SetUserContext(v *BlindTransferResponseBodyDataUserContext) *BlindTransferResponseBodyData {
	s.UserContext = v
	return s
}

type BlindTransferResponseBodyDataCallContext struct {
	CallType        *string                                                    `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                    `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*BlindTransferResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s BlindTransferResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s BlindTransferResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *BlindTransferResponseBodyDataCallContext) SetCallType(v string) *BlindTransferResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContext) SetInstanceId(v string) *BlindTransferResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContext) SetJobId(v string) *BlindTransferResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContext) SetChannelContexts(v []*BlindTransferResponseBodyDataCallContextChannelContexts) *BlindTransferResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type BlindTransferResponseBodyDataCallContextChannelContexts struct {
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
}

func (s BlindTransferResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s BlindTransferResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetDestination(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetUserId(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetCallType(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetJobId(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *BlindTransferResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *BlindTransferResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

type BlindTransferResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s BlindTransferResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s BlindTransferResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *BlindTransferResponseBodyDataUserContext) SetExtension(v string) *BlindTransferResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetHeartbeat(v int64) *BlindTransferResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetWorkMode(v string) *BlindTransferResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetDeviceId(v string) *BlindTransferResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetUserId(v string) *BlindTransferResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetReserved(v int64) *BlindTransferResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetBreakCode(v string) *BlindTransferResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetInstanceId(v string) *BlindTransferResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetOutboundScenario(v bool) *BlindTransferResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetMobile(v string) *BlindTransferResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetJobId(v string) *BlindTransferResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetUserState(v string) *BlindTransferResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *BlindTransferResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *BlindTransferResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type BlindTransferResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BlindTransferResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BlindTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s BlindTransferResponse) GoString() string {
	return s.String()
}

func (s *BlindTransferResponse) SetHeaders(v map[string]*string) *BlindTransferResponse {
	s.Headers = v
	return s
}

func (s *BlindTransferResponse) SetBody(v *BlindTransferResponseBody) *BlindTransferResponse {
	s.Body = v
	return s
}

type ListSkillLevelsOfUserRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	IsMember      *bool   `json:"IsMember,omitempty" xml:"IsMember,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s ListSkillLevelsOfUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSkillLevelsOfUserRequest) GoString() string {
	return s.String()
}

func (s *ListSkillLevelsOfUserRequest) SetInstanceId(v string) *ListSkillLevelsOfUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSkillLevelsOfUserRequest) SetUserId(v string) *ListSkillLevelsOfUserRequest {
	s.UserId = &v
	return s
}

func (s *ListSkillLevelsOfUserRequest) SetIsMember(v bool) *ListSkillLevelsOfUserRequest {
	s.IsMember = &v
	return s
}

func (s *ListSkillLevelsOfUserRequest) SetPageNumber(v int32) *ListSkillLevelsOfUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSkillLevelsOfUserRequest) SetPageSize(v int32) *ListSkillLevelsOfUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillLevelsOfUserRequest) SetSearchPattern(v string) *ListSkillLevelsOfUserRequest {
	s.SearchPattern = &v
	return s
}

type ListSkillLevelsOfUserResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListSkillLevelsOfUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListSkillLevelsOfUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSkillLevelsOfUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillLevelsOfUserResponseBody) SetCode(v string) *ListSkillLevelsOfUserResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBody) SetHttpStatusCode(v int32) *ListSkillLevelsOfUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBody) SetMessage(v string) *ListSkillLevelsOfUserResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBody) SetRequestId(v string) *ListSkillLevelsOfUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBody) SetData(v *ListSkillLevelsOfUserResponseBodyData) *ListSkillLevelsOfUserResponseBody {
	s.Data = v
	return s
}

type ListSkillLevelsOfUserResponseBodyData struct {
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListSkillLevelsOfUserResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListSkillLevelsOfUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSkillLevelsOfUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillLevelsOfUserResponseBodyData) SetPageNumber(v int32) *ListSkillLevelsOfUserResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyData) SetPageSize(v int32) *ListSkillLevelsOfUserResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyData) SetTotalCount(v int32) *ListSkillLevelsOfUserResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyData) SetList(v []*ListSkillLevelsOfUserResponseBodyDataList) *ListSkillLevelsOfUserResponseBodyData {
	s.List = v
	return s
}

type ListSkillLevelsOfUserResponseBodyDataList struct {
	SkillLevel     *string `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
	SkillGroupId   *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s ListSkillLevelsOfUserResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListSkillLevelsOfUserResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListSkillLevelsOfUserResponseBodyDataList) SetSkillLevel(v string) *ListSkillLevelsOfUserResponseBodyDataList {
	s.SkillLevel = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyDataList) SetSkillGroupId(v string) *ListSkillLevelsOfUserResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListSkillLevelsOfUserResponseBodyDataList) SetSkillGroupName(v string) *ListSkillLevelsOfUserResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

type ListSkillLevelsOfUserResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSkillLevelsOfUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSkillLevelsOfUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSkillLevelsOfUserResponse) GoString() string {
	return s.String()
}

func (s *ListSkillLevelsOfUserResponse) SetHeaders(v map[string]*string) *ListSkillLevelsOfUserResponse {
	s.Headers = v
	return s
}

func (s *ListSkillLevelsOfUserResponse) SetBody(v *ListSkillLevelsOfUserResponseBody) *ListSkillLevelsOfUserResponse {
	s.Body = v
	return s
}

type ListUnassignedNumbersRequest struct {
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListUnassignedNumbersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUnassignedNumbersRequest) GoString() string {
	return s.String()
}

func (s *ListUnassignedNumbersRequest) SetSearchPattern(v string) *ListUnassignedNumbersRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListUnassignedNumbersRequest) SetPageNumber(v int32) *ListUnassignedNumbersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUnassignedNumbersRequest) SetPageSize(v int32) *ListUnassignedNumbersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUnassignedNumbersRequest) SetInstanceId(v string) *ListUnassignedNumbersRequest {
	s.InstanceId = &v
	return s
}

type ListUnassignedNumbersResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListUnassignedNumbersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListUnassignedNumbersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUnassignedNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUnassignedNumbersResponseBody) SetCode(v string) *ListUnassignedNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUnassignedNumbersResponseBody) SetHttpStatusCode(v int32) *ListUnassignedNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUnassignedNumbersResponseBody) SetMessage(v string) *ListUnassignedNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUnassignedNumbersResponseBody) SetRequestId(v string) *ListUnassignedNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUnassignedNumbersResponseBody) SetData(v *ListUnassignedNumbersResponseBodyData) *ListUnassignedNumbersResponseBody {
	s.Data = v
	return s
}

type ListUnassignedNumbersResponseBodyData struct {
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListUnassignedNumbersResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListUnassignedNumbersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUnassignedNumbersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUnassignedNumbersResponseBodyData) SetPageNumber(v int32) *ListUnassignedNumbersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyData) SetPageSize(v int32) *ListUnassignedNumbersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyData) SetTotalCount(v int32) *ListUnassignedNumbersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyData) SetList(v []*ListUnassignedNumbersResponseBodyDataList) *ListUnassignedNumbersResponseBodyData {
	s.List = v
	return s
}

type ListUnassignedNumbersResponseBodyDataList struct {
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Active   *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	Usage    *bool   `json:"Usage,omitempty" xml:"Usage,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListUnassignedNumbersResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListUnassignedNumbersResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListUnassignedNumbersResponseBodyDataList) SetNumber(v string) *ListUnassignedNumbersResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyDataList) SetActive(v bool) *ListUnassignedNumbersResponseBodyDataList {
	s.Active = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyDataList) SetCity(v string) *ListUnassignedNumbersResponseBodyDataList {
	s.City = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyDataList) SetUsage(v bool) *ListUnassignedNumbersResponseBodyDataList {
	s.Usage = &v
	return s
}

func (s *ListUnassignedNumbersResponseBodyDataList) SetProvince(v string) *ListUnassignedNumbersResponseBodyDataList {
	s.Province = &v
	return s
}

type ListUnassignedNumbersResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUnassignedNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUnassignedNumbersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUnassignedNumbersResponse) GoString() string {
	return s.String()
}

func (s *ListUnassignedNumbersResponse) SetHeaders(v map[string]*string) *ListUnassignedNumbersResponse {
	s.Headers = v
	return s
}

func (s *ListUnassignedNumbersResponse) SetBody(v *ListUnassignedNumbersResponseBody) *ListUnassignedNumbersResponse {
	s.Body = v
	return s
}

type GetInstanceTrendingReportRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s GetInstanceTrendingReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTrendingReportRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportRequest) SetInstanceId(v string) *GetInstanceTrendingReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) SetStartTime(v int64) *GetInstanceTrendingReportRequest {
	s.StartTime = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) SetEndTime(v int64) *GetInstanceTrendingReportRequest {
	s.EndTime = &v
	return s
}

type GetInstanceTrendingReportResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetInstanceTrendingReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetInstanceTrendingReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTrendingReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportResponseBody) SetCode(v string) *GetInstanceTrendingReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBody) SetHttpStatusCode(v int32) *GetInstanceTrendingReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBody) SetMessage(v string) *GetInstanceTrendingReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBody) SetRequestId(v string) *GetInstanceTrendingReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBody) SetData(v *GetInstanceTrendingReportResponseBodyData) *GetInstanceTrendingReportResponseBody {
	s.Data = v
	return s
}

type GetInstanceTrendingReportResponseBodyData struct {
	Inbound  []*GetInstanceTrendingReportResponseBodyDataInbound  `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Repeated"`
	Outbound []*GetInstanceTrendingReportResponseBodyDataOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Repeated"`
}

func (s GetInstanceTrendingReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTrendingReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportResponseBodyData) SetInbound(v []*GetInstanceTrendingReportResponseBodyDataInbound) *GetInstanceTrendingReportResponseBodyData {
	s.Inbound = v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyData) SetOutbound(v []*GetInstanceTrendingReportResponseBodyDataOutbound) *GetInstanceTrendingReportResponseBodyData {
	s.Outbound = v
	return s
}

type GetInstanceTrendingReportResponseBodyDataInbound struct {
	StatsTime             *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	CallsQueued           *int64 `json:"CallsQueued,omitempty" xml:"CallsQueued,omitempty"`
	CallsAbandonedInRing  *int64 `json:"CallsAbandonedInRing,omitempty" xml:"CallsAbandonedInRing,omitempty"`
	CallsHandled          *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	TotalCalls            *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	CallsAbandonedInIVR   *int64 `json:"CallsAbandonedInIVR,omitempty" xml:"CallsAbandonedInIVR,omitempty"`
	CallsAbandonedInQueue *int64 `json:"CallsAbandonedInQueue,omitempty" xml:"CallsAbandonedInQueue,omitempty"`
}

func (s GetInstanceTrendingReportResponseBodyDataInbound) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTrendingReportResponseBodyDataInbound) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetStatsTime(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.StatsTime = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetCallsQueued(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.CallsQueued = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetCallsAbandonedInRing(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.CallsAbandonedInRing = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetCallsHandled(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.CallsHandled = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetTotalCalls(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.TotalCalls = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetCallsAbandonedInIVR(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.CallsAbandonedInIVR = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataInbound) SetCallsAbandonedInQueue(v int64) *GetInstanceTrendingReportResponseBodyDataInbound {
	s.CallsAbandonedInQueue = &v
	return s
}

type GetInstanceTrendingReportResponseBodyDataOutbound struct {
	StatsTime     *int64 `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	CallsAnswered *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	TotalCalls    *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
}

func (s GetInstanceTrendingReportResponseBodyDataOutbound) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTrendingReportResponseBodyDataOutbound) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportResponseBodyDataOutbound) SetStatsTime(v int64) *GetInstanceTrendingReportResponseBodyDataOutbound {
	s.StatsTime = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataOutbound) SetCallsAnswered(v int64) *GetInstanceTrendingReportResponseBodyDataOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *GetInstanceTrendingReportResponseBodyDataOutbound) SetTotalCalls(v int64) *GetInstanceTrendingReportResponseBodyDataOutbound {
	s.TotalCalls = &v
	return s
}

type GetInstanceTrendingReportResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceTrendingReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceTrendingReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTrendingReportResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportResponse) SetHeaders(v map[string]*string) *GetInstanceTrendingReportResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceTrendingReportResponse) SetBody(v *GetInstanceTrendingReportResponseBody) *GetInstanceTrendingReportResponse {
	s.Body = v
	return s
}

type ListInstancesOfUserRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInstancesOfUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesOfUserRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserRequest) SetPageNumber(v int32) *ListInstancesOfUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesOfUserRequest) SetPageSize(v int32) *ListInstancesOfUserRequest {
	s.PageSize = &v
	return s
}

type ListInstancesOfUserResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                            `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *ListInstancesOfUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListInstancesOfUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesOfUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponseBody) SetCode(v string) *ListInstancesOfUserResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesOfUserResponseBody) SetHttpStatusCode(v int32) *ListInstancesOfUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesOfUserResponseBody) SetMessage(v string) *ListInstancesOfUserResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesOfUserResponseBody) SetRequestId(v string) *ListInstancesOfUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesOfUserResponseBody) SetParams(v []*string) *ListInstancesOfUserResponseBody {
	s.Params = v
	return s
}

func (s *ListInstancesOfUserResponseBody) SetData(v *ListInstancesOfUserResponseBodyData) *ListInstancesOfUserResponseBody {
	s.Data = v
	return s
}

type ListInstancesOfUserResponseBodyData struct {
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListInstancesOfUserResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListInstancesOfUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesOfUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponseBodyData) SetPageNumber(v int32) *ListInstancesOfUserResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyData) SetPageSize(v int32) *ListInstancesOfUserResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyData) SetTotalCount(v int32) *ListInstancesOfUserResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyData) SetList(v []*ListInstancesOfUserResponseBodyDataList) *ListInstancesOfUserResponseBodyData {
	s.List = v
	return s
}

type ListInstancesOfUserResponseBodyDataList struct {
	Status      *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	ConsoleUrl  *string                                              `json:"ConsoleUrl,omitempty" xml:"ConsoleUrl,omitempty"`
	Description *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	AliyunUid   *string                                              `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	Name        *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	DomainName  *string                                              `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Id          *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	AdminList   []*ListInstancesOfUserResponseBodyDataListAdminList  `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	NumberList  []*ListInstancesOfUserResponseBodyDataListNumberList `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
}

func (s ListInstancesOfUserResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesOfUserResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponseBodyDataList) SetStatus(v string) *ListInstancesOfUserResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetConsoleUrl(v string) *ListInstancesOfUserResponseBodyDataList {
	s.ConsoleUrl = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetDescription(v string) *ListInstancesOfUserResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetAliyunUid(v string) *ListInstancesOfUserResponseBodyDataList {
	s.AliyunUid = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetName(v string) *ListInstancesOfUserResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetDomainName(v string) *ListInstancesOfUserResponseBodyDataList {
	s.DomainName = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetId(v string) *ListInstancesOfUserResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetAdminList(v []*ListInstancesOfUserResponseBodyDataListAdminList) *ListInstancesOfUserResponseBodyDataList {
	s.AdminList = v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataList) SetNumberList(v []*ListInstancesOfUserResponseBodyDataListNumberList) *ListInstancesOfUserResponseBodyDataList {
	s.NumberList = v
	return s
}

type ListInstancesOfUserResponseBodyDataListAdminList struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Extension   *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	LoginName   *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	WorkMode    *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RoleName    *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleId      *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s ListInstancesOfUserResponseBodyDataListAdminList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesOfUserResponseBodyDataListAdminList) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetDisplayName(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.DisplayName = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetExtension(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.Extension = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetLoginName(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.LoginName = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetEmail(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.Email = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetWorkMode(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.WorkMode = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetMobile(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.Mobile = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetUserId(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.UserId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetRoleName(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.RoleName = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetInstanceId(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListAdminList) SetRoleId(v string) *ListInstancesOfUserResponseBodyDataListAdminList {
	s.RoleId = &v
	return s
}

type ListInstancesOfUserResponseBodyDataListNumberList struct {
	Active        *bool                                                           `json:"Active,omitempty" xml:"Active,omitempty"`
	UserId        *string                                                         `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Number        *string                                                         `json:"Number,omitempty" xml:"Number,omitempty"`
	City          *string                                                         `json:"City,omitempty" xml:"City,omitempty"`
	InstanceId    *string                                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Usage         *string                                                         `json:"Usage,omitempty" xml:"Usage,omitempty"`
	ContactFlowId *string                                                         `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	Province      *string                                                         `json:"Province,omitempty" xml:"Province,omitempty"`
	SkillGroups   []*ListInstancesOfUserResponseBodyDataListNumberListSkillGroups `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty" type:"Repeated"`
}

func (s ListInstancesOfUserResponseBodyDataListNumberList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesOfUserResponseBodyDataListNumberList) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetActive(v bool) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.Active = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetUserId(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.UserId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetNumber(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.Number = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetCity(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.City = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetInstanceId(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetUsage(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.Usage = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetContactFlowId(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.ContactFlowId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetProvince(v string) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.Province = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberList) SetSkillGroups(v []*ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) *ListInstancesOfUserResponseBodyDataListNumberList {
	s.SkillGroups = v
	return s
}

type ListInstancesOfUserResponseBodyDataListNumberListSkillGroups struct {
	DisplayName      *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PhoneNumberCount *int32  `json:"PhoneNumberCount,omitempty" xml:"PhoneNumberCount,omitempty"`
	SkillGroupId     *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	UserCount        *int32  `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetDisplayName(v string) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.DisplayName = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetDescription(v string) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.Description = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetPhoneNumberCount(v int32) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.PhoneNumberCount = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetSkillGroupId(v string) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.SkillGroupId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetUserCount(v int32) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.UserCount = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetInstanceId(v string) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups) SetName(v string) *ListInstancesOfUserResponseBodyDataListNumberListSkillGroups {
	s.Name = &v
	return s
}

type ListInstancesOfUserResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesOfUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancesOfUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesOfUserResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesOfUserResponse) SetHeaders(v map[string]*string) *ListInstancesOfUserResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesOfUserResponse) SetBody(v *ListInstancesOfUserResponseBody) *ListInstancesOfUserResponse {
	s.Body = v
	return s
}

type LaunchSurveyRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
}

func (s LaunchSurveyRequest) String() string {
	return tea.Prettify(s)
}

func (s LaunchSurveyRequest) GoString() string {
	return s.String()
}

func (s *LaunchSurveyRequest) SetInstanceId(v string) *LaunchSurveyRequest {
	s.InstanceId = &v
	return s
}

func (s *LaunchSurveyRequest) SetUserId(v string) *LaunchSurveyRequest {
	s.UserId = &v
	return s
}

func (s *LaunchSurveyRequest) SetDeviceId(v string) *LaunchSurveyRequest {
	s.DeviceId = &v
	return s
}

func (s *LaunchSurveyRequest) SetJobId(v string) *LaunchSurveyRequest {
	s.JobId = &v
	return s
}

func (s *LaunchSurveyRequest) SetContactFlowId(v string) *LaunchSurveyRequest {
	s.ContactFlowId = &v
	return s
}

type LaunchSurveyResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                     `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *LaunchSurveyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s LaunchSurveyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LaunchSurveyResponseBody) GoString() string {
	return s.String()
}

func (s *LaunchSurveyResponseBody) SetCode(v string) *LaunchSurveyResponseBody {
	s.Code = &v
	return s
}

func (s *LaunchSurveyResponseBody) SetHttpStatusCode(v int32) *LaunchSurveyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *LaunchSurveyResponseBody) SetMessage(v string) *LaunchSurveyResponseBody {
	s.Message = &v
	return s
}

func (s *LaunchSurveyResponseBody) SetRequestId(v string) *LaunchSurveyResponseBody {
	s.RequestId = &v
	return s
}

func (s *LaunchSurveyResponseBody) SetParams(v []*string) *LaunchSurveyResponseBody {
	s.Params = v
	return s
}

func (s *LaunchSurveyResponseBody) SetData(v *LaunchSurveyResponseBodyData) *LaunchSurveyResponseBody {
	s.Data = v
	return s
}

type LaunchSurveyResponseBodyData struct {
	CallContext *LaunchSurveyResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *LaunchSurveyResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s LaunchSurveyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LaunchSurveyResponseBodyData) GoString() string {
	return s.String()
}

func (s *LaunchSurveyResponseBodyData) SetCallContext(v *LaunchSurveyResponseBodyDataCallContext) *LaunchSurveyResponseBodyData {
	s.CallContext = v
	return s
}

func (s *LaunchSurveyResponseBodyData) SetUserContext(v *LaunchSurveyResponseBodyDataUserContext) *LaunchSurveyResponseBodyData {
	s.UserContext = v
	return s
}

type LaunchSurveyResponseBodyDataCallContext struct {
	CallType        *string                                                   `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*LaunchSurveyResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s LaunchSurveyResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s LaunchSurveyResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *LaunchSurveyResponseBodyDataCallContext) SetCallType(v string) *LaunchSurveyResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContext) SetInstanceId(v string) *LaunchSurveyResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContext) SetJobId(v string) *LaunchSurveyResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContext) SetChannelContexts(v []*LaunchSurveyResponseBodyDataCallContextChannelContexts) *LaunchSurveyResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type LaunchSurveyResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s LaunchSurveyResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s LaunchSurveyResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetDestination(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetUserId(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetCallType(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetJobId(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *LaunchSurveyResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type LaunchSurveyResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s LaunchSurveyResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s LaunchSurveyResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetExtension(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetHeartbeat(v int64) *LaunchSurveyResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetWorkMode(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetDeviceId(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetUserId(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetReserved(v int64) *LaunchSurveyResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetBreakCode(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetInstanceId(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetOutboundScenario(v bool) *LaunchSurveyResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetMobile(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetJobId(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetUserState(v string) *LaunchSurveyResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *LaunchSurveyResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *LaunchSurveyResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type LaunchSurveyResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LaunchSurveyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LaunchSurveyResponse) String() string {
	return tea.Prettify(s)
}

func (s LaunchSurveyResponse) GoString() string {
	return s.String()
}

func (s *LaunchSurveyResponse) SetHeaders(v map[string]*string) *LaunchSurveyResponse {
	s.Headers = v
	return s
}

func (s *LaunchSurveyResponse) SetBody(v *LaunchSurveyResponseBody) *LaunchSurveyResponse {
	s.Body = v
	return s
}

type ListIvrTrackingDetailsRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ContactId  *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListIvrTrackingDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIvrTrackingDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListIvrTrackingDetailsRequest) SetPageNumber(v int32) *ListIvrTrackingDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIvrTrackingDetailsRequest) SetPageSize(v int32) *ListIvrTrackingDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIvrTrackingDetailsRequest) SetContactId(v string) *ListIvrTrackingDetailsRequest {
	s.ContactId = &v
	return s
}

func (s *ListIvrTrackingDetailsRequest) SetInstanceId(v string) *ListIvrTrackingDetailsRequest {
	s.InstanceId = &v
	return s
}

type ListIvrTrackingDetailsResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListIvrTrackingDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListIvrTrackingDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIvrTrackingDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIvrTrackingDetailsResponseBody) SetCode(v string) *ListIvrTrackingDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBody) SetHttpStatusCode(v int32) *ListIvrTrackingDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBody) SetMessage(v string) *ListIvrTrackingDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBody) SetRequestId(v string) *ListIvrTrackingDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBody) SetData(v *ListIvrTrackingDetailsResponseBodyData) *ListIvrTrackingDetailsResponseBody {
	s.Data = v
	return s
}

type ListIvrTrackingDetailsResponseBodyData struct {
	PageNumber *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListIvrTrackingDetailsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListIvrTrackingDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIvrTrackingDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIvrTrackingDetailsResponseBodyData) SetPageNumber(v int32) *ListIvrTrackingDetailsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyData) SetPageSize(v int32) *ListIvrTrackingDetailsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyData) SetTotalCount(v int32) *ListIvrTrackingDetailsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyData) SetList(v []*ListIvrTrackingDetailsResponseBodyDataList) *ListIvrTrackingDetailsResponseBodyData {
	s.List = v
	return s
}

type ListIvrTrackingDetailsResponseBodyDataList struct {
	Instance         *string                `json:"Instance,omitempty" xml:"Instance,omitempty"`
	Callee           *string                `json:"Callee,omitempty" xml:"Callee,omitempty"`
	EnterTime        *int64                 `json:"EnterTime,omitempty" xml:"EnterTime,omitempty"`
	NodeVariables    map[string]interface{} `json:"NodeVariables,omitempty" xml:"NodeVariables,omitempty"`
	NodeExitCode     *string                `json:"NodeExitCode,omitempty" xml:"NodeExitCode,omitempty"`
	FlowName         *string                `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	FlowId           *string                `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	NodeProperties   map[string]interface{} `json:"NodeProperties,omitempty" xml:"NodeProperties,omitempty"`
	NodeType         *string                `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Caller           *string                `json:"Caller,omitempty" xml:"Caller,omitempty"`
	NodeName         *string                `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	ContactId        *string                `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	LeaveTime        *int64                 `json:"LeaveTime,omitempty" xml:"LeaveTime,omitempty"`
	ChannelVariables *string                `json:"ChannelVariables,omitempty" xml:"ChannelVariables,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	NodeId           *string                `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ListIvrTrackingDetailsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListIvrTrackingDetailsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetInstance(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.Instance = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetCallee(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.Callee = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetEnterTime(v int64) *ListIvrTrackingDetailsResponseBodyDataList {
	s.EnterTime = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetNodeVariables(v map[string]interface{}) *ListIvrTrackingDetailsResponseBodyDataList {
	s.NodeVariables = v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetNodeExitCode(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.NodeExitCode = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetFlowName(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.FlowName = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetFlowId(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.FlowId = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetNodeProperties(v map[string]interface{}) *ListIvrTrackingDetailsResponseBodyDataList {
	s.NodeProperties = v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetNodeType(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.NodeType = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetCaller(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.Caller = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetNodeName(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.NodeName = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetContactId(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.ContactId = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetLeaveTime(v int64) *ListIvrTrackingDetailsResponseBodyDataList {
	s.LeaveTime = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetChannelVariables(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.ChannelVariables = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetChannelId(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.ChannelId = &v
	return s
}

func (s *ListIvrTrackingDetailsResponseBodyDataList) SetNodeId(v string) *ListIvrTrackingDetailsResponseBodyDataList {
	s.NodeId = &v
	return s
}

type ListIvrTrackingDetailsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListIvrTrackingDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIvrTrackingDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIvrTrackingDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListIvrTrackingDetailsResponse) SetHeaders(v map[string]*string) *ListIvrTrackingDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListIvrTrackingDetailsResponse) SetBody(v *ListIvrTrackingDetailsResponseBody) *ListIvrTrackingDetailsResponse {
	s.Body = v
	return s
}

type ListBriefSkillGroupsRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListBriefSkillGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBriefSkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListBriefSkillGroupsRequest) SetInstanceId(v string) *ListBriefSkillGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBriefSkillGroupsRequest) SetSearchPattern(v string) *ListBriefSkillGroupsRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListBriefSkillGroupsRequest) SetPageNumber(v int32) *ListBriefSkillGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBriefSkillGroupsRequest) SetPageSize(v int32) *ListBriefSkillGroupsRequest {
	s.PageSize = &v
	return s
}

type ListBriefSkillGroupsResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListBriefSkillGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListBriefSkillGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBriefSkillGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBriefSkillGroupsResponseBody) SetCode(v string) *ListBriefSkillGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBody) SetHttpStatusCode(v int32) *ListBriefSkillGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBody) SetMessage(v string) *ListBriefSkillGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBody) SetRequestId(v string) *ListBriefSkillGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBody) SetData(v *ListBriefSkillGroupsResponseBodyData) *ListBriefSkillGroupsResponseBody {
	s.Data = v
	return s
}

type ListBriefSkillGroupsResponseBodyData struct {
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListBriefSkillGroupsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListBriefSkillGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListBriefSkillGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBriefSkillGroupsResponseBodyData) SetPageNumber(v int32) *ListBriefSkillGroupsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyData) SetPageSize(v int32) *ListBriefSkillGroupsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyData) SetTotalCount(v int32) *ListBriefSkillGroupsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyData) SetList(v []*ListBriefSkillGroupsResponseBodyDataList) *ListBriefSkillGroupsResponseBodyData {
	s.List = v
	return s
}

type ListBriefSkillGroupsResponseBodyDataList struct {
	DisplayName      *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PhoneNumberCount *int32  `json:"PhoneNumberCount,omitempty" xml:"PhoneNumberCount,omitempty"`
	SkillGroupId     *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName   *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	UserCount        *int32  `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListBriefSkillGroupsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListBriefSkillGroupsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetDisplayName(v string) *ListBriefSkillGroupsResponseBodyDataList {
	s.DisplayName = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetDescription(v string) *ListBriefSkillGroupsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetPhoneNumberCount(v int32) *ListBriefSkillGroupsResponseBodyDataList {
	s.PhoneNumberCount = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetSkillGroupId(v string) *ListBriefSkillGroupsResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetSkillGroupName(v string) *ListBriefSkillGroupsResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetUserCount(v int32) *ListBriefSkillGroupsResponseBodyDataList {
	s.UserCount = &v
	return s
}

func (s *ListBriefSkillGroupsResponseBodyDataList) SetInstanceId(v string) *ListBriefSkillGroupsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

type ListBriefSkillGroupsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListBriefSkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBriefSkillGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBriefSkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListBriefSkillGroupsResponse) SetHeaders(v map[string]*string) *ListBriefSkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListBriefSkillGroupsResponse) SetBody(v *ListBriefSkillGroupsResponseBody) *ListBriefSkillGroupsResponse {
	s.Body = v
	return s
}

type UnmuteCallRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s UnmuteCallRequest) String() string {
	return tea.Prettify(s)
}

func (s UnmuteCallRequest) GoString() string {
	return s.String()
}

func (s *UnmuteCallRequest) SetInstanceId(v string) *UnmuteCallRequest {
	s.InstanceId = &v
	return s
}

func (s *UnmuteCallRequest) SetUserId(v string) *UnmuteCallRequest {
	s.UserId = &v
	return s
}

func (s *UnmuteCallRequest) SetDeviceId(v string) *UnmuteCallRequest {
	s.DeviceId = &v
	return s
}

func (s *UnmuteCallRequest) SetJobId(v string) *UnmuteCallRequest {
	s.JobId = &v
	return s
}

func (s *UnmuteCallRequest) SetChannelId(v string) *UnmuteCallRequest {
	s.ChannelId = &v
	return s
}

type UnmuteCallResponseBody struct {
	Code           *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                   `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *UnmuteCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s UnmuteCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnmuteCallResponseBody) GoString() string {
	return s.String()
}

func (s *UnmuteCallResponseBody) SetCode(v string) *UnmuteCallResponseBody {
	s.Code = &v
	return s
}

func (s *UnmuteCallResponseBody) SetHttpStatusCode(v int32) *UnmuteCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UnmuteCallResponseBody) SetMessage(v string) *UnmuteCallResponseBody {
	s.Message = &v
	return s
}

func (s *UnmuteCallResponseBody) SetRequestId(v string) *UnmuteCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnmuteCallResponseBody) SetParams(v []*string) *UnmuteCallResponseBody {
	s.Params = v
	return s
}

func (s *UnmuteCallResponseBody) SetData(v *UnmuteCallResponseBodyData) *UnmuteCallResponseBody {
	s.Data = v
	return s
}

type UnmuteCallResponseBodyData struct {
	CallContext *UnmuteCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *UnmuteCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s UnmuteCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UnmuteCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnmuteCallResponseBodyData) SetCallContext(v *UnmuteCallResponseBodyDataCallContext) *UnmuteCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *UnmuteCallResponseBodyData) SetUserContext(v *UnmuteCallResponseBodyDataUserContext) *UnmuteCallResponseBodyData {
	s.UserContext = v
	return s
}

type UnmuteCallResponseBodyDataCallContext struct {
	CallType        *string                                                 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*UnmuteCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s UnmuteCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s UnmuteCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *UnmuteCallResponseBodyDataCallContext) SetCallType(v string) *UnmuteCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContext) SetInstanceId(v string) *UnmuteCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContext) SetJobId(v string) *UnmuteCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContext) SetChannelContexts(v []*UnmuteCallResponseBodyDataCallContextChannelContexts) *UnmuteCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type UnmuteCallResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s UnmuteCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s UnmuteCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *UnmuteCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *UnmuteCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type UnmuteCallResponseBodyDataUserContext struct {
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s UnmuteCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s UnmuteCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *UnmuteCallResponseBodyDataUserContext) SetHeartbeat(v int64) *UnmuteCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetExtension(v string) *UnmuteCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetWorkMode(v string) *UnmuteCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetMobile(v string) *UnmuteCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetDeviceId(v string) *UnmuteCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetJobId(v string) *UnmuteCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetUserId(v string) *UnmuteCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetBreakCode(v string) *UnmuteCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetInstanceId(v string) *UnmuteCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *UnmuteCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetUserState(v string) *UnmuteCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *UnmuteCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *UnmuteCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type UnmuteCallResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnmuteCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnmuteCallResponse) String() string {
	return tea.Prettify(s)
}

func (s UnmuteCallResponse) GoString() string {
	return s.String()
}

func (s *UnmuteCallResponse) SetHeaders(v map[string]*string) *UnmuteCallResponse {
	s.Headers = v
	return s
}

func (s *UnmuteCallResponse) SetBody(v *UnmuteCallResponseBody) *UnmuteCallResponse {
	s.Body = v
	return s
}

type ModifySkillLevelsOfUserRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SkillLevelList *string `json:"SkillLevelList,omitempty" xml:"SkillLevelList,omitempty"`
}

func (s ModifySkillLevelsOfUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySkillLevelsOfUserRequest) GoString() string {
	return s.String()
}

func (s *ModifySkillLevelsOfUserRequest) SetInstanceId(v string) *ModifySkillLevelsOfUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySkillLevelsOfUserRequest) SetUserId(v string) *ModifySkillLevelsOfUserRequest {
	s.UserId = &v
	return s
}

func (s *ModifySkillLevelsOfUserRequest) SetSkillLevelList(v string) *ModifySkillLevelsOfUserRequest {
	s.SkillLevelList = &v
	return s
}

type ModifySkillLevelsOfUserResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s ModifySkillLevelsOfUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySkillLevelsOfUserResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySkillLevelsOfUserResponseBody) SetCode(v string) *ModifySkillLevelsOfUserResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySkillLevelsOfUserResponseBody) SetHttpStatusCode(v int32) *ModifySkillLevelsOfUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifySkillLevelsOfUserResponseBody) SetMessage(v string) *ModifySkillLevelsOfUserResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySkillLevelsOfUserResponseBody) SetRequestId(v string) *ModifySkillLevelsOfUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySkillLevelsOfUserResponseBody) SetParams(v []*string) *ModifySkillLevelsOfUserResponseBody {
	s.Params = v
	return s
}

type ModifySkillLevelsOfUserResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySkillLevelsOfUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySkillLevelsOfUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySkillLevelsOfUserResponse) GoString() string {
	return s.String()
}

func (s *ModifySkillLevelsOfUserResponse) SetHeaders(v map[string]*string) *ModifySkillLevelsOfUserResponse {
	s.Headers = v
	return s
}

func (s *ModifySkillLevelsOfUserResponse) SetBody(v *ModifySkillLevelsOfUserResponseBody) *ModifySkillLevelsOfUserResponse {
	s.Body = v
	return s
}

type AssignUsersRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleId         *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	SkillLevelList *string `json:"SkillLevelList,omitempty" xml:"SkillLevelList,omitempty"`
	RamIdList      *string `json:"RamIdList,omitempty" xml:"RamIdList,omitempty"`
	WorkMode       *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s AssignUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignUsersRequest) GoString() string {
	return s.String()
}

func (s *AssignUsersRequest) SetInstanceId(v string) *AssignUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *AssignUsersRequest) SetRoleId(v string) *AssignUsersRequest {
	s.RoleId = &v
	return s
}

func (s *AssignUsersRequest) SetSkillLevelList(v string) *AssignUsersRequest {
	s.SkillLevelList = &v
	return s
}

func (s *AssignUsersRequest) SetRamIdList(v string) *AssignUsersRequest {
	s.RamIdList = &v
	return s
}

func (s *AssignUsersRequest) SetWorkMode(v string) *AssignUsersRequest {
	s.WorkMode = &v
	return s
}

type AssignUsersResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sync           *string `json:"Sync,omitempty" xml:"Sync,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	WorkflowId     *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AssignUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AssignUsersResponseBody) SetHttpStatusCode(v int32) *AssignUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AssignUsersResponseBody) SetData(v string) *AssignUsersResponseBody {
	s.Data = &v
	return s
}

func (s *AssignUsersResponseBody) SetRequestId(v string) *AssignUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignUsersResponseBody) SetSync(v string) *AssignUsersResponseBody {
	s.Sync = &v
	return s
}

func (s *AssignUsersResponseBody) SetCode(v string) *AssignUsersResponseBody {
	s.Code = &v
	return s
}

func (s *AssignUsersResponseBody) SetWorkflowId(v string) *AssignUsersResponseBody {
	s.WorkflowId = &v
	return s
}

func (s *AssignUsersResponseBody) SetMessage(v string) *AssignUsersResponseBody {
	s.Message = &v
	return s
}

type AssignUsersResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AssignUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssignUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignUsersResponse) GoString() string {
	return s.String()
}

func (s *AssignUsersResponse) SetHeaders(v map[string]*string) *AssignUsersResponse {
	s.Headers = v
	return s
}

func (s *AssignUsersResponse) SetBody(v *AssignUsersResponseBody) *AssignUsersResponse {
	s.Body = v
	return s
}

type ListUserLevelsOfSkillGroupRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId  *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	IsMember      *bool   `json:"IsMember,omitempty" xml:"IsMember,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s ListUserLevelsOfSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserLevelsOfSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *ListUserLevelsOfSkillGroupRequest) SetInstanceId(v string) *ListUserLevelsOfSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupRequest) SetSkillGroupId(v string) *ListUserLevelsOfSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupRequest) SetIsMember(v bool) *ListUserLevelsOfSkillGroupRequest {
	s.IsMember = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupRequest) SetPageNumber(v int32) *ListUserLevelsOfSkillGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupRequest) SetPageSize(v int32) *ListUserLevelsOfSkillGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupRequest) SetSearchPattern(v string) *ListUserLevelsOfSkillGroupRequest {
	s.SearchPattern = &v
	return s
}

type ListUserLevelsOfSkillGroupResponseBody struct {
	Code           *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListUserLevelsOfSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListUserLevelsOfSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserLevelsOfSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserLevelsOfSkillGroupResponseBody) SetCode(v string) *ListUserLevelsOfSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBody) SetHttpStatusCode(v int32) *ListUserLevelsOfSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBody) SetMessage(v string) *ListUserLevelsOfSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBody) SetRequestId(v string) *ListUserLevelsOfSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBody) SetData(v *ListUserLevelsOfSkillGroupResponseBodyData) *ListUserLevelsOfSkillGroupResponseBody {
	s.Data = v
	return s
}

type ListUserLevelsOfSkillGroupResponseBodyData struct {
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListUserLevelsOfSkillGroupResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListUserLevelsOfSkillGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserLevelsOfSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) SetPageNumber(v int32) *ListUserLevelsOfSkillGroupResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) SetPageSize(v int32) *ListUserLevelsOfSkillGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) SetTotalCount(v int32) *ListUserLevelsOfSkillGroupResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyData) SetList(v []*ListUserLevelsOfSkillGroupResponseBodyDataList) *ListUserLevelsOfSkillGroupResponseBodyData {
	s.List = v
	return s
}

type ListUserLevelsOfSkillGroupResponseBodyDataList struct {
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	SkillLevel     *int32  `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
	LoginName      *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SkillGroupId   *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	RoleName       *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	RoleId         *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s ListUserLevelsOfSkillGroupResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListUserLevelsOfSkillGroupResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetDisplayName(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.DisplayName = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetSkillLevel(v int32) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.SkillLevel = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetLoginName(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.LoginName = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetUserId(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetSkillGroupId(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetSkillGroupName(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetRoleName(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.RoleName = &v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponseBodyDataList) SetRoleId(v string) *ListUserLevelsOfSkillGroupResponseBodyDataList {
	s.RoleId = &v
	return s
}

type ListUserLevelsOfSkillGroupResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserLevelsOfSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserLevelsOfSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserLevelsOfSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *ListUserLevelsOfSkillGroupResponse) SetHeaders(v map[string]*string) *ListUserLevelsOfSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *ListUserLevelsOfSkillGroupResponse) SetBody(v *ListUserLevelsOfSkillGroupResponseBody) *ListUserLevelsOfSkillGroupResponse {
	s.Body = v
	return s
}

type ListRolesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) SetInstanceId(v string) *ListRolesRequest {
	s.InstanceId = &v
	return s
}

type ListRolesResponseBody struct {
	Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*ListRolesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) SetCode(v string) *ListRolesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRolesResponseBody) SetHttpStatusCode(v int32) *ListRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRolesResponseBody) SetMessage(v string) *ListRolesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) SetData(v []*ListRolesResponseBodyData) *ListRolesResponseBody {
	s.Data = v
	return s
}

type ListRolesResponseBodyData struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s ListRolesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyData) SetName(v string) *ListRolesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListRolesResponseBodyData) SetRoleId(v string) *ListRolesResponseBodyData {
	s.RoleId = &v
	return s
}

type ListRolesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRolesResponse) SetHeaders(v map[string]*string) *ListRolesResponse {
	s.Headers = v
	return s
}

func (s *ListRolesResponse) SetBody(v *ListRolesResponseBody) *ListRolesResponse {
	s.Body = v
	return s
}

type UpdateConfigItemsRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ObjectId    *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ObjectType  *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	ConfigItems *string `json:"ConfigItems,omitempty" xml:"ConfigItems,omitempty"`
}

func (s UpdateConfigItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigItemsRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigItemsRequest) SetInstanceId(v string) *UpdateConfigItemsRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConfigItemsRequest) SetObjectId(v string) *UpdateConfigItemsRequest {
	s.ObjectId = &v
	return s
}

func (s *UpdateConfigItemsRequest) SetObjectType(v string) *UpdateConfigItemsRequest {
	s.ObjectType = &v
	return s
}

func (s *UpdateConfigItemsRequest) SetConfigItems(v string) *UpdateConfigItemsRequest {
	s.ConfigItems = &v
	return s
}

type UpdateConfigItemsResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s UpdateConfigItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigItemsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigItemsResponseBody) SetCode(v string) *UpdateConfigItemsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConfigItemsResponseBody) SetHttpStatusCode(v int32) *UpdateConfigItemsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateConfigItemsResponseBody) SetMessage(v string) *UpdateConfigItemsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConfigItemsResponseBody) SetRequestId(v string) *UpdateConfigItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigItemsResponseBody) SetParams(v []*string) *UpdateConfigItemsResponseBody {
	s.Params = v
	return s
}

type UpdateConfigItemsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateConfigItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConfigItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigItemsResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigItemsResponse) SetHeaders(v map[string]*string) *UpdateConfigItemsResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigItemsResponse) SetBody(v *UpdateConfigItemsResponseBody) *UpdateConfigItemsResponse {
	s.Body = v
	return s
}

type GetCallDetailRecordRequest struct {
	ContactId  *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCallDetailRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCallDetailRecordRequest) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordRequest) SetContactId(v string) *GetCallDetailRecordRequest {
	s.ContactId = &v
	return s
}

func (s *GetCallDetailRecordRequest) SetInstanceId(v string) *GetCallDetailRecordRequest {
	s.InstanceId = &v
	return s
}

type GetCallDetailRecordResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetCallDetailRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetCallDetailRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCallDetailRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBody) SetCode(v string) *GetCallDetailRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetHttpStatusCode(v int32) *GetCallDetailRecordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetMessage(v string) *GetCallDetailRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetRequestId(v string) *GetCallDetailRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCallDetailRecordResponseBody) SetData(v *GetCallDetailRecordResponseBodyData) *GetCallDetailRecordResponseBody {
	s.Data = v
	return s
}

type GetCallDetailRecordResponseBodyData struct {
	ReleaseInitiator          *string                                           `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ContactDisposition        *string                                           `json:"ContactDisposition,omitempty" xml:"ContactDisposition,omitempty"`
	ContactType               *string                                           `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	AgentIds                  *string                                           `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	CallDuration              *int64                                            `json:"CallDuration,omitempty" xml:"CallDuration,omitempty"`
	RecordingReady            *bool                                             `json:"RecordingReady,omitempty" xml:"RecordingReady,omitempty"`
	EstablishedTime           *int64                                            `json:"EstablishedTime,omitempty" xml:"EstablishedTime,omitempty"`
	InstanceId                *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SatisfactionSurveyOffered *bool                                             `json:"SatisfactionSurveyOffered,omitempty" xml:"SatisfactionSurveyOffered,omitempty"`
	CalledNumber              *string                                           `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	AgentNames                *string                                           `json:"AgentNames,omitempty" xml:"AgentNames,omitempty"`
	Satisfaction              *int32                                            `json:"Satisfaction,omitempty" xml:"Satisfaction,omitempty"`
	StartTime                 *int64                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ContactId                 *string                                           `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	SatisfactionSurveyChannel *string                                           `json:"SatisfactionSurveyChannel,omitempty" xml:"SatisfactionSurveyChannel,omitempty"`
	ReleaseTime               *int64                                            `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	CallingNumber             *string                                           `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	SkillGroupNames           *string                                           `json:"SkillGroupNames,omitempty" xml:"SkillGroupNames,omitempty"`
	SkillGroupIds             *string                                           `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty"`
	AgentEvents               []*GetCallDetailRecordResponseBodyDataAgentEvents `json:"AgentEvents,omitempty" xml:"AgentEvents,omitempty" type:"Repeated"`
	IvrEvents                 []*GetCallDetailRecordResponseBodyDataIvrEvents   `json:"IvrEvents,omitempty" xml:"IvrEvents,omitempty" type:"Repeated"`
	QueueEvents               []*GetCallDetailRecordResponseBodyDataQueueEvents `json:"QueueEvents,omitempty" xml:"QueueEvents,omitempty" type:"Repeated"`
	CallerLocation            *string                                           `json:"CallerLocation,omitempty" xml:"CallerLocation,omitempty"`
	CalleeLocation            *string                                           `json:"CalleeLocation,omitempty" xml:"CalleeLocation,omitempty"`
}

func (s GetCallDetailRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyData) SetReleaseInitiator(v string) *GetCallDetailRecordResponseBodyData {
	s.ReleaseInitiator = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetContactDisposition(v string) *GetCallDetailRecordResponseBodyData {
	s.ContactDisposition = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetContactType(v string) *GetCallDetailRecordResponseBodyData {
	s.ContactType = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetAgentIds(v string) *GetCallDetailRecordResponseBodyData {
	s.AgentIds = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCallDuration(v int64) *GetCallDetailRecordResponseBodyData {
	s.CallDuration = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetRecordingReady(v bool) *GetCallDetailRecordResponseBodyData {
	s.RecordingReady = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetEstablishedTime(v int64) *GetCallDetailRecordResponseBodyData {
	s.EstablishedTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetInstanceId(v string) *GetCallDetailRecordResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetSatisfactionSurveyOffered(v bool) *GetCallDetailRecordResponseBodyData {
	s.SatisfactionSurveyOffered = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCalledNumber(v string) *GetCallDetailRecordResponseBodyData {
	s.CalledNumber = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetAgentNames(v string) *GetCallDetailRecordResponseBodyData {
	s.AgentNames = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetSatisfaction(v int32) *GetCallDetailRecordResponseBodyData {
	s.Satisfaction = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetStartTime(v int64) *GetCallDetailRecordResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetContactId(v string) *GetCallDetailRecordResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetSatisfactionSurveyChannel(v string) *GetCallDetailRecordResponseBodyData {
	s.SatisfactionSurveyChannel = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetReleaseTime(v int64) *GetCallDetailRecordResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCallingNumber(v string) *GetCallDetailRecordResponseBodyData {
	s.CallingNumber = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetSkillGroupNames(v string) *GetCallDetailRecordResponseBodyData {
	s.SkillGroupNames = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetSkillGroupIds(v string) *GetCallDetailRecordResponseBodyData {
	s.SkillGroupIds = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetAgentEvents(v []*GetCallDetailRecordResponseBodyDataAgentEvents) *GetCallDetailRecordResponseBodyData {
	s.AgentEvents = v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetIvrEvents(v []*GetCallDetailRecordResponseBodyDataIvrEvents) *GetCallDetailRecordResponseBodyData {
	s.IvrEvents = v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetQueueEvents(v []*GetCallDetailRecordResponseBodyDataQueueEvents) *GetCallDetailRecordResponseBodyData {
	s.QueueEvents = v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCallerLocation(v string) *GetCallDetailRecordResponseBodyData {
	s.CallerLocation = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyData) SetCalleeLocation(v string) *GetCallDetailRecordResponseBodyData {
	s.CalleeLocation = &v
	return s
}

type GetCallDetailRecordResponseBodyDataAgentEvents struct {
	AgentName     *string                                                        `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	AgentId       *string                                                        `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	SkillGroupId  *string                                                        `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	EventSequence []*GetCallDetailRecordResponseBodyDataAgentEventsEventSequence `json:"EventSequence,omitempty" xml:"EventSequence,omitempty" type:"Repeated"`
}

func (s GetCallDetailRecordResponseBodyDataAgentEvents) String() string {
	return tea.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataAgentEvents) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) SetAgentName(v string) *GetCallDetailRecordResponseBodyDataAgentEvents {
	s.AgentName = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) SetAgentId(v string) *GetCallDetailRecordResponseBodyDataAgentEvents {
	s.AgentId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) SetSkillGroupId(v string) *GetCallDetailRecordResponseBodyDataAgentEvents {
	s.SkillGroupId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAgentEvents) SetEventSequence(v []*GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) *GetCallDetailRecordResponseBodyDataAgentEvents {
	s.EventSequence = v
	return s
}

type GetCallDetailRecordResponseBodyDataAgentEventsEventSequence struct {
	Event     *string `json:"Event,omitempty" xml:"Event,omitempty"`
	EventTime *int64  `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) String() string {
	return tea.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) SetEvent(v string) *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence {
	s.Event = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence) SetEventTime(v int64) *GetCallDetailRecordResponseBodyDataAgentEventsEventSequence {
	s.EventTime = &v
	return s
}

type GetCallDetailRecordResponseBodyDataIvrEvents struct {
	FlowId        *string                                                      `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	EventSequence []*GetCallDetailRecordResponseBodyDataIvrEventsEventSequence `json:"EventSequence,omitempty" xml:"EventSequence,omitempty" type:"Repeated"`
}

func (s GetCallDetailRecordResponseBodyDataIvrEvents) String() string {
	return tea.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataIvrEvents) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataIvrEvents) SetFlowId(v string) *GetCallDetailRecordResponseBodyDataIvrEvents {
	s.FlowId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataIvrEvents) SetEventSequence(v []*GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) *GetCallDetailRecordResponseBodyDataIvrEvents {
	s.EventSequence = v
	return s
}

type GetCallDetailRecordResponseBodyDataIvrEventsEventSequence struct {
	Event     *string `json:"Event,omitempty" xml:"Event,omitempty"`
	EventTime *int64  `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) String() string {
	return tea.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) SetEvent(v string) *GetCallDetailRecordResponseBodyDataIvrEventsEventSequence {
	s.Event = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataIvrEventsEventSequence) SetEventTime(v int64) *GetCallDetailRecordResponseBodyDataIvrEventsEventSequence {
	s.EventTime = &v
	return s
}

type GetCallDetailRecordResponseBodyDataQueueEvents struct {
	QueueId       *string                                                        `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	QueueName     *string                                                        `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	FlowId        *string                                                        `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	QueueType     *int32                                                         `json:"QueueType,omitempty" xml:"QueueType,omitempty"`
	EventSequence []*GetCallDetailRecordResponseBodyDataQueueEventsEventSequence `json:"EventSequence,omitempty" xml:"EventSequence,omitempty" type:"Repeated"`
}

func (s GetCallDetailRecordResponseBodyDataQueueEvents) String() string {
	return tea.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataQueueEvents) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) SetQueueId(v string) *GetCallDetailRecordResponseBodyDataQueueEvents {
	s.QueueId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) SetQueueName(v string) *GetCallDetailRecordResponseBodyDataQueueEvents {
	s.QueueName = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) SetFlowId(v string) *GetCallDetailRecordResponseBodyDataQueueEvents {
	s.FlowId = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) SetQueueType(v int32) *GetCallDetailRecordResponseBodyDataQueueEvents {
	s.QueueType = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataQueueEvents) SetEventSequence(v []*GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) *GetCallDetailRecordResponseBodyDataQueueEvents {
	s.EventSequence = v
	return s
}

type GetCallDetailRecordResponseBodyDataQueueEventsEventSequence struct {
	Event     *string `json:"Event,omitempty" xml:"Event,omitempty"`
	EventTime *int64  `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
}

func (s GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) String() string {
	return tea.Prettify(s)
}

func (s GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) SetEvent(v string) *GetCallDetailRecordResponseBodyDataQueueEventsEventSequence {
	s.Event = &v
	return s
}

func (s *GetCallDetailRecordResponseBodyDataQueueEventsEventSequence) SetEventTime(v int64) *GetCallDetailRecordResponseBodyDataQueueEventsEventSequence {
	s.EventTime = &v
	return s
}

type GetCallDetailRecordResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCallDetailRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCallDetailRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCallDetailRecordResponse) GoString() string {
	return s.String()
}

func (s *GetCallDetailRecordResponse) SetHeaders(v map[string]*string) *GetCallDetailRecordResponse {
	s.Headers = v
	return s
}

func (s *GetCallDetailRecordResponse) SetBody(v *GetCallDetailRecordResponseBody) *GetCallDetailRecordResponse {
	s.Body = v
	return s
}

type ModifyPhoneNumberRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Usage         *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
}

func (s ModifyPhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *ModifyPhoneNumberRequest) SetInstanceId(v string) *ModifyPhoneNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPhoneNumberRequest) SetNumber(v string) *ModifyPhoneNumberRequest {
	s.Number = &v
	return s
}

func (s *ModifyPhoneNumberRequest) SetUsage(v string) *ModifyPhoneNumberRequest {
	s.Usage = &v
	return s
}

func (s *ModifyPhoneNumberRequest) SetContactFlowId(v string) *ModifyPhoneNumberRequest {
	s.ContactFlowId = &v
	return s
}

type ModifyPhoneNumberResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPhoneNumberResponseBody) SetCode(v string) *ModifyPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyPhoneNumberResponseBody) SetHttpStatusCode(v int32) *ModifyPhoneNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyPhoneNumberResponseBody) SetMessage(v string) *ModifyPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyPhoneNumberResponseBody) SetRequestId(v string) *ModifyPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPhoneNumberResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *ModifyPhoneNumberResponse) SetHeaders(v map[string]*string) *ModifyPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *ModifyPhoneNumberResponse) SetBody(v *ModifyPhoneNumberResponseBody) *ModifyPhoneNumberResponse {
	s.Body = v
	return s
}

type CoachCallRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId       *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	CoachedUserId  *string `json:"CoachedUserId,omitempty" xml:"CoachedUserId,omitempty"`
	TimeoutSeconds *int32  `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s CoachCallRequest) String() string {
	return tea.Prettify(s)
}

func (s CoachCallRequest) GoString() string {
	return s.String()
}

func (s *CoachCallRequest) SetInstanceId(v string) *CoachCallRequest {
	s.InstanceId = &v
	return s
}

func (s *CoachCallRequest) SetUserId(v string) *CoachCallRequest {
	s.UserId = &v
	return s
}

func (s *CoachCallRequest) SetDeviceId(v string) *CoachCallRequest {
	s.DeviceId = &v
	return s
}

func (s *CoachCallRequest) SetJobId(v string) *CoachCallRequest {
	s.JobId = &v
	return s
}

func (s *CoachCallRequest) SetCoachedUserId(v string) *CoachCallRequest {
	s.CoachedUserId = &v
	return s
}

func (s *CoachCallRequest) SetTimeoutSeconds(v int32) *CoachCallRequest {
	s.TimeoutSeconds = &v
	return s
}

type CoachCallResponseBody struct {
	Code           *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                  `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *CoachCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CoachCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CoachCallResponseBody) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBody) SetCode(v string) *CoachCallResponseBody {
	s.Code = &v
	return s
}

func (s *CoachCallResponseBody) SetHttpStatusCode(v int32) *CoachCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CoachCallResponseBody) SetMessage(v string) *CoachCallResponseBody {
	s.Message = &v
	return s
}

func (s *CoachCallResponseBody) SetRequestId(v string) *CoachCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *CoachCallResponseBody) SetParams(v []*string) *CoachCallResponseBody {
	s.Params = v
	return s
}

func (s *CoachCallResponseBody) SetData(v *CoachCallResponseBodyData) *CoachCallResponseBody {
	s.Data = v
	return s
}

type CoachCallResponseBodyData struct {
	CallContext *CoachCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *CoachCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s CoachCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CoachCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBodyData) SetCallContext(v *CoachCallResponseBodyDataCallContext) *CoachCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *CoachCallResponseBodyData) SetUserContext(v *CoachCallResponseBodyDataUserContext) *CoachCallResponseBodyData {
	s.UserContext = v
	return s
}

type CoachCallResponseBodyDataCallContext struct {
	CallType        *string                                                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*CoachCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s CoachCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s CoachCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBodyDataCallContext) SetCallType(v string) *CoachCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContext) SetInstanceId(v string) *CoachCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContext) SetJobId(v string) *CoachCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContext) SetChannelContexts(v []*CoachCallResponseBodyDataCallContextChannelContexts) *CoachCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type CoachCallResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s CoachCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s CoachCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *CoachCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *CoachCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type CoachCallResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Uri                    *string   `json:"Uri,omitempty" xml:"Uri,omitempty"`
	DeviceState            *string   `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s CoachCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s CoachCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *CoachCallResponseBodyDataUserContext) SetExtension(v string) *CoachCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetHeartbeat(v int64) *CoachCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetWorkMode(v string) *CoachCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetDeviceId(v string) *CoachCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetUserId(v string) *CoachCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetReserved(v int64) *CoachCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetBreakCode(v string) *CoachCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetInstanceId(v string) *CoachCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *CoachCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetUri(v string) *CoachCallResponseBodyDataUserContext {
	s.Uri = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetDeviceState(v string) *CoachCallResponseBodyDataUserContext {
	s.DeviceState = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetMobile(v string) *CoachCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetJobId(v string) *CoachCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetUserState(v string) *CoachCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *CoachCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *CoachCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type CoachCallResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CoachCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CoachCallResponse) String() string {
	return tea.Prettify(s)
}

func (s CoachCallResponse) GoString() string {
	return s.String()
}

func (s *CoachCallResponse) SetHeaders(v map[string]*string) *CoachCallResponse {
	s.Headers = v
	return s
}

func (s *CoachCallResponse) SetBody(v *CoachCallResponseBody) *CoachCallResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LoginName      *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Mobile         *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Email          *string `json:"Email,omitempty" xml:"Email,omitempty"`
	WorkMode       *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	RoleId         *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	SkillLevelList *string `json:"SkillLevelList,omitempty" xml:"SkillLevelList,omitempty"`
	ResetPassword  *bool   `json:"ResetPassword,omitempty" xml:"ResetPassword,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetInstanceId(v string) *CreateUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserRequest) SetLoginName(v string) *CreateUserRequest {
	s.LoginName = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetMobile(v string) *CreateUserRequest {
	s.Mobile = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetWorkMode(v string) *CreateUserRequest {
	s.WorkMode = &v
	return s
}

func (s *CreateUserRequest) SetRoleId(v string) *CreateUserRequest {
	s.RoleId = &v
	return s
}

func (s *CreateUserRequest) SetSkillLevelList(v string) *CreateUserRequest {
	s.SkillLevelList = &v
	return s
}

func (s *CreateUserRequest) SetResetPassword(v bool) *CreateUserRequest {
	s.ResetPassword = &v
	return s
}

type CreateUserResponseBody struct {
	Code           *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                   `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *CreateUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetCode(v string) *CreateUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserResponseBody) SetHttpStatusCode(v int32) *CreateUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateUserResponseBody) SetMessage(v string) *CreateUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetParams(v []*string) *CreateUserResponseBody {
	s.Params = v
	return s
}

func (s *CreateUserResponseBody) SetData(v *CreateUserResponseBodyData) *CreateUserResponseBody {
	s.Data = v
	return s
}

type CreateUserResponseBodyData struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Extension   *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	LoginName   *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	WorkMode    *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyData) SetDisplayName(v string) *CreateUserResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *CreateUserResponseBodyData) SetExtension(v string) *CreateUserResponseBodyData {
	s.Extension = &v
	return s
}

func (s *CreateUserResponseBodyData) SetEmail(v string) *CreateUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *CreateUserResponseBodyData) SetLoginName(v string) *CreateUserResponseBodyData {
	s.LoginName = &v
	return s
}

func (s *CreateUserResponseBodyData) SetWorkMode(v string) *CreateUserResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *CreateUserResponseBodyData) SetMobile(v string) *CreateUserResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *CreateUserResponseBodyData) SetUserId(v string) *CreateUserResponseBodyData {
	s.UserId = &v
	return s
}

type CreateUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type ListPrivilegesOfUserRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListPrivilegesOfUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrivilegesOfUserRequest) GoString() string {
	return s.String()
}

func (s *ListPrivilegesOfUserRequest) SetInstanceId(v string) *ListPrivilegesOfUserRequest {
	s.InstanceId = &v
	return s
}

type ListPrivilegesOfUserResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*ListPrivilegesOfUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListPrivilegesOfUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrivilegesOfUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivilegesOfUserResponseBody) SetCode(v string) *ListPrivilegesOfUserResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBody) SetHttpStatusCode(v int32) *ListPrivilegesOfUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBody) SetMessage(v string) *ListPrivilegesOfUserResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBody) SetRequestId(v string) *ListPrivilegesOfUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBody) SetData(v []*ListPrivilegesOfUserResponseBodyData) *ListPrivilegesOfUserResponseBody {
	s.Data = v
	return s
}

type ListPrivilegesOfUserResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Scope      *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListPrivilegesOfUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPrivilegesOfUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPrivilegesOfUserResponseBodyData) SetInstanceId(v string) *ListPrivilegesOfUserResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBodyData) SetScope(v string) *ListPrivilegesOfUserResponseBodyData {
	s.Scope = &v
	return s
}

func (s *ListPrivilegesOfUserResponseBodyData) SetName(v string) *ListPrivilegesOfUserResponseBodyData {
	s.Name = &v
	return s
}

type ListPrivilegesOfUserResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPrivilegesOfUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPrivilegesOfUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrivilegesOfUserResponse) GoString() string {
	return s.String()
}

func (s *ListPrivilegesOfUserResponse) SetHeaders(v map[string]*string) *ListPrivilegesOfUserResponse {
	s.Headers = v
	return s
}

func (s *ListPrivilegesOfUserResponse) SetBody(v *ListPrivilegesOfUserResponseBody) *ListPrivilegesOfUserResponse {
	s.Body = v
	return s
}

type AddPersonalNumbersToUserRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
}

func (s AddPersonalNumbersToUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPersonalNumbersToUserRequest) GoString() string {
	return s.String()
}

func (s *AddPersonalNumbersToUserRequest) SetInstanceId(v string) *AddPersonalNumbersToUserRequest {
	s.InstanceId = &v
	return s
}

func (s *AddPersonalNumbersToUserRequest) SetUserId(v string) *AddPersonalNumbersToUserRequest {
	s.UserId = &v
	return s
}

func (s *AddPersonalNumbersToUserRequest) SetNumberList(v string) *AddPersonalNumbersToUserRequest {
	s.NumberList = &v
	return s
}

type AddPersonalNumbersToUserResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s AddPersonalNumbersToUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddPersonalNumbersToUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddPersonalNumbersToUserResponseBody) SetCode(v string) *AddPersonalNumbersToUserResponseBody {
	s.Code = &v
	return s
}

func (s *AddPersonalNumbersToUserResponseBody) SetHttpStatusCode(v int32) *AddPersonalNumbersToUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddPersonalNumbersToUserResponseBody) SetMessage(v string) *AddPersonalNumbersToUserResponseBody {
	s.Message = &v
	return s
}

func (s *AddPersonalNumbersToUserResponseBody) SetRequestId(v string) *AddPersonalNumbersToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPersonalNumbersToUserResponseBody) SetData(v []*string) *AddPersonalNumbersToUserResponseBody {
	s.Data = v
	return s
}

type AddPersonalNumbersToUserResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddPersonalNumbersToUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddPersonalNumbersToUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPersonalNumbersToUserResponse) GoString() string {
	return s.String()
}

func (s *AddPersonalNumbersToUserResponse) SetHeaders(v map[string]*string) *AddPersonalNumbersToUserResponse {
	s.Headers = v
	return s
}

func (s *AddPersonalNumbersToUserResponse) SetBody(v *AddPersonalNumbersToUserResponseBody) *AddPersonalNumbersToUserResponse {
	s.Body = v
	return s
}

type ListHistoricalAgentReportRequest struct {
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	AgentIdList *string `json:"AgentIdList,omitempty" xml:"AgentIdList,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StopTime    *int64  `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListHistoricalAgentReportRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentReportRequest) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportRequest) SetPageNumber(v int32) *ListHistoricalAgentReportRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetPageSize(v int32) *ListHistoricalAgentReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetAgentIdList(v string) *ListHistoricalAgentReportRequest {
	s.AgentIdList = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetStartTime(v int64) *ListHistoricalAgentReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetStopTime(v int64) *ListHistoricalAgentReportRequest {
	s.StopTime = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetInstanceId(v string) *ListHistoricalAgentReportRequest {
	s.InstanceId = &v
	return s
}

type ListHistoricalAgentReportResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListHistoricalAgentReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListHistoricalAgentReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBody) SetCode(v string) *ListHistoricalAgentReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBody) SetHttpStatusCode(v int32) *ListHistoricalAgentReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBody) SetMessage(v string) *ListHistoricalAgentReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBody) SetRequestId(v string) *ListHistoricalAgentReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBody) SetData(v *ListHistoricalAgentReportResponseBodyData) *ListHistoricalAgentReportResponseBody {
	s.Data = v
	return s
}

type ListHistoricalAgentReportResponseBodyData struct {
	PageNumber *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListHistoricalAgentReportResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListHistoricalAgentReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyData) SetPageNumber(v int32) *ListHistoricalAgentReportResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyData) SetPageSize(v int32) *ListHistoricalAgentReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyData) SetTotalCount(v int32) *ListHistoricalAgentReportResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyData) SetList(v []*ListHistoricalAgentReportResponseBodyDataList) *ListHistoricalAgentReportResponseBodyData {
	s.List = v
	return s
}

type ListHistoricalAgentReportResponseBodyDataList struct {
	AgentName *string                                                `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	AgentId   *string                                                `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	Inbound   *ListHistoricalAgentReportResponseBodyDataListInbound  `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Outbound  *ListHistoricalAgentReportResponseBodyDataListOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall   *ListHistoricalAgentReportResponseBodyDataListOverall  `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
}

func (s ListHistoricalAgentReportResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetAgentName(v string) *ListHistoricalAgentReportResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetAgentId(v string) *ListHistoricalAgentReportResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetInbound(v *ListHistoricalAgentReportResponseBodyDataListInbound) *ListHistoricalAgentReportResponseBodyDataList {
	s.Inbound = v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetOutbound(v *ListHistoricalAgentReportResponseBodyDataListOutbound) *ListHistoricalAgentReportResponseBodyDataList {
	s.Outbound = v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataList) SetOverall(v *ListHistoricalAgentReportResponseBodyDataListOverall) *ListHistoricalAgentReportResponseBodyDataList {
	s.Overall = v
	return s
}

type ListHistoricalAgentReportResponseBodyDataListInbound struct {
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	CallsHandled                 *int64   `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	CallsAttendedTransferOut     *int64   `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsBlindTransferIn         *int64   `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	CallsAttendedTransferIn      *int64   `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	CallsBlindTransferOut        *int64   `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	CallsOffered                 *int64   `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	HandleRate                   *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
}

func (s ListHistoricalAgentReportResponseBodyDataListInbound) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyDataListInbound) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetAverageRingTime(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsHandled(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetTotalWorkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsAttendedTransferOut(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetMaxWorkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetTotalHoldTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetAverageWorkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsBlindTransferIn(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetSatisfactionIndex(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsRinged(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsAttendedTransferIn(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsBlindTransferOut(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetTotalRingTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetMaxTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetMaxRingTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetTotalTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsOffered(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetMaxHoldTime(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetAverageTalkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetSatisfactionRate(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetCallsHold(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.CallsHold = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetSatisfactionSurveysOffered(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetHandleRate(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.HandleRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetSatisfactionSurveysResponded(v int64) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListInbound) SetAverageHoldTime(v float32) *ListHistoricalAgentReportResponseBodyDataListInbound {
	s.AverageHoldTime = &v
	return s
}

type ListHistoricalAgentReportResponseBodyDataListOutbound struct {
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	CallsDialed                  *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsAnswered                *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	CallsAttendedTransferOut     *int64   `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	TotalDialingTime             *int64   `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsBlindTransferIn         *int64   `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	CallsAttendedTransferIn      *int64   `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	CallsBlindTransferOut        *int64   `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	MaxDialingTime               *int64   `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	AnswerRate                   *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageDialingTime           *float32 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
}

func (s ListHistoricalAgentReportResponseBodyDataListOutbound) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyDataListOutbound) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetAverageRingTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsDialed(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsAnswered(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetTotalWorkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsAttendedTransferOut(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetMaxWorkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetTotalDialingTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetTotalHoldTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetAverageWorkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsBlindTransferIn(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetSatisfactionIndex(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsRinged(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsAttendedTransferIn(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsBlindTransferOut(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetTotalRingTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetMaxTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetMaxRingTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetTotalTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetMaxDialingTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetAnswerRate(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetMaxHoldTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetAverageTalkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetSatisfactionRate(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetCallsHold(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetSatisfactionSurveysOffered(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetSatisfactionSurveysResponded(v int64) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetAverageHoldTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOutbound) SetAverageDialingTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOutbound {
	s.AverageDialingTime = &v
	return s
}

type ListHistoricalAgentReportResponseBodyDataListOverall struct {
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalLoggedInTime            *int64   `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	OccupancyRate                *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	AverageBreakTime             *float32 `json:"AverageBreakTime,omitempty" xml:"AverageBreakTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	MaxBreakTime                 *int64   `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	MaxReadyTime                 *int64   `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	AverageReadyTime             *float32 `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	TotalReadyTime               *int64   `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	TotalBreakTime               *int64   `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalCalls                   *int64   `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
}

func (s ListHistoricalAgentReportResponseBodyDataListOverall) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentReportResponseBodyDataListOverall) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalLoggedInTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetOccupancyRate(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalWorkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetMaxHoldTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetMaxWorkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetAverageBreakTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalHoldTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetSatisfactionRate(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetMaxBreakTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetAverageWorkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetAverageTalkTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetSatisfactionIndex(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetSatisfactionSurveysOffered(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetSatisfactionSurveysResponded(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetMaxReadyTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetAverageReadyTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetAverageHoldTime(v float32) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalReadyTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalBreakTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetMaxTalkTime(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentReportResponseBodyDataListOverall) SetTotalCalls(v int64) *ListHistoricalAgentReportResponseBodyDataListOverall {
	s.TotalCalls = &v
	return s
}

type ListHistoricalAgentReportResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHistoricalAgentReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHistoricalAgentReportResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentReportResponse) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponse) SetHeaders(v map[string]*string) *ListHistoricalAgentReportResponse {
	s.Headers = v
	return s
}

func (s *ListHistoricalAgentReportResponse) SetBody(v *ListHistoricalAgentReportResponseBody) *ListHistoricalAgentReportResponse {
	s.Body = v
	return s
}

type InterceptCallRequest struct {
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId          *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId             *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	InterceptedUserId *string `json:"InterceptedUserId,omitempty" xml:"InterceptedUserId,omitempty"`
	TimeoutSeconds    *int32  `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s InterceptCallRequest) String() string {
	return tea.Prettify(s)
}

func (s InterceptCallRequest) GoString() string {
	return s.String()
}

func (s *InterceptCallRequest) SetInstanceId(v string) *InterceptCallRequest {
	s.InstanceId = &v
	return s
}

func (s *InterceptCallRequest) SetUserId(v string) *InterceptCallRequest {
	s.UserId = &v
	return s
}

func (s *InterceptCallRequest) SetDeviceId(v string) *InterceptCallRequest {
	s.DeviceId = &v
	return s
}

func (s *InterceptCallRequest) SetJobId(v string) *InterceptCallRequest {
	s.JobId = &v
	return s
}

func (s *InterceptCallRequest) SetInterceptedUserId(v string) *InterceptCallRequest {
	s.InterceptedUserId = &v
	return s
}

func (s *InterceptCallRequest) SetTimeoutSeconds(v int32) *InterceptCallRequest {
	s.TimeoutSeconds = &v
	return s
}

type InterceptCallResponseBody struct {
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                      `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *InterceptCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s InterceptCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InterceptCallResponseBody) GoString() string {
	return s.String()
}

func (s *InterceptCallResponseBody) SetCode(v string) *InterceptCallResponseBody {
	s.Code = &v
	return s
}

func (s *InterceptCallResponseBody) SetHttpStatusCode(v int32) *InterceptCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InterceptCallResponseBody) SetMessage(v string) *InterceptCallResponseBody {
	s.Message = &v
	return s
}

func (s *InterceptCallResponseBody) SetRequestId(v string) *InterceptCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *InterceptCallResponseBody) SetParams(v []*string) *InterceptCallResponseBody {
	s.Params = v
	return s
}

func (s *InterceptCallResponseBody) SetData(v *InterceptCallResponseBodyData) *InterceptCallResponseBody {
	s.Data = v
	return s
}

type InterceptCallResponseBodyData struct {
	CallContext *InterceptCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *InterceptCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s InterceptCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InterceptCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *InterceptCallResponseBodyData) SetCallContext(v *InterceptCallResponseBodyDataCallContext) *InterceptCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *InterceptCallResponseBodyData) SetUserContext(v *InterceptCallResponseBodyDataUserContext) *InterceptCallResponseBodyData {
	s.UserContext = v
	return s
}

type InterceptCallResponseBodyDataCallContext struct {
	CallType        *string                                                    `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                    `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*InterceptCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s InterceptCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s InterceptCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *InterceptCallResponseBodyDataCallContext) SetCallType(v string) *InterceptCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContext) SetInstanceId(v string) *InterceptCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContext) SetJobId(v string) *InterceptCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContext) SetChannelContexts(v []*InterceptCallResponseBodyDataCallContextChannelContexts) *InterceptCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type InterceptCallResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s InterceptCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s InterceptCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *InterceptCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *InterceptCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type InterceptCallResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s InterceptCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s InterceptCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *InterceptCallResponseBodyDataUserContext) SetExtension(v string) *InterceptCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetHeartbeat(v int64) *InterceptCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetWorkMode(v string) *InterceptCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetDeviceId(v string) *InterceptCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetUserId(v string) *InterceptCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetReserved(v int64) *InterceptCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetBreakCode(v string) *InterceptCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetInstanceId(v string) *InterceptCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *InterceptCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetMobile(v string) *InterceptCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetJobId(v string) *InterceptCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetUserState(v string) *InterceptCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *InterceptCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *InterceptCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type InterceptCallResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InterceptCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InterceptCallResponse) String() string {
	return tea.Prettify(s)
}

func (s InterceptCallResponse) GoString() string {
	return s.String()
}

func (s *InterceptCallResponse) SetHeaders(v map[string]*string) *InterceptCallResponse {
	s.Headers = v
	return s
}

func (s *InterceptCallResponse) SetBody(v *InterceptCallResponseBody) *InterceptCallResponse {
	s.Body = v
	return s
}

type ListContactFlowsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListContactFlowsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContactFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListContactFlowsRequest) SetInstanceId(v string) *ListContactFlowsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListContactFlowsRequest) SetPageNumber(v int32) *ListContactFlowsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContactFlowsRequest) SetPageSize(v int32) *ListContactFlowsRequest {
	s.PageSize = &v
	return s
}

func (s *ListContactFlowsRequest) SetType(v string) *ListContactFlowsRequest {
	s.Type = &v
	return s
}

type ListContactFlowsResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListContactFlowsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListContactFlowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContactFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactFlowsResponseBody) SetCode(v string) *ListContactFlowsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContactFlowsResponseBody) SetHttpStatusCode(v int32) *ListContactFlowsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListContactFlowsResponseBody) SetMessage(v string) *ListContactFlowsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContactFlowsResponseBody) SetRequestId(v string) *ListContactFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContactFlowsResponseBody) SetData(v *ListContactFlowsResponseBodyData) *ListContactFlowsResponseBody {
	s.Data = v
	return s
}

type ListContactFlowsResponseBodyData struct {
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListContactFlowsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListContactFlowsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListContactFlowsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContactFlowsResponseBodyData) SetPageNumber(v int32) *ListContactFlowsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListContactFlowsResponseBodyData) SetPageSize(v int32) *ListContactFlowsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListContactFlowsResponseBodyData) SetTotalCount(v int32) *ListContactFlowsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListContactFlowsResponseBodyData) SetList(v []*ListContactFlowsResponseBodyDataList) *ListContactFlowsResponseBodyData {
	s.List = v
	return s
}

type ListContactFlowsResponseBodyDataList struct {
	Type          *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Definition    *string   `json:"Definition,omitempty" xml:"Definition,omitempty"`
	DraftId       *string   `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
	Description   *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	UpdatedTime   *string   `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	Editor        *string   `json:"Editor,omitempty" xml:"Editor,omitempty"`
	Published     *bool     `json:"Published,omitempty" xml:"Published,omitempty"`
	InstanceId    *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	ContactFlowId *string   `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	CreatedTime   *string   `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	NumberList    []*string `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
}

func (s ListContactFlowsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListContactFlowsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListContactFlowsResponseBodyDataList) SetType(v string) *ListContactFlowsResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetDefinition(v string) *ListContactFlowsResponseBodyDataList {
	s.Definition = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetDraftId(v string) *ListContactFlowsResponseBodyDataList {
	s.DraftId = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetDescription(v string) *ListContactFlowsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetUpdatedTime(v string) *ListContactFlowsResponseBodyDataList {
	s.UpdatedTime = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetEditor(v string) *ListContactFlowsResponseBodyDataList {
	s.Editor = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetPublished(v bool) *ListContactFlowsResponseBodyDataList {
	s.Published = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetInstanceId(v string) *ListContactFlowsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetName(v string) *ListContactFlowsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetContactFlowId(v string) *ListContactFlowsResponseBodyDataList {
	s.ContactFlowId = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetCreatedTime(v string) *ListContactFlowsResponseBodyDataList {
	s.CreatedTime = &v
	return s
}

func (s *ListContactFlowsResponseBodyDataList) SetNumberList(v []*string) *ListContactFlowsResponseBodyDataList {
	s.NumberList = v
	return s
}

type ListContactFlowsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListContactFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContactFlowsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContactFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListContactFlowsResponse) SetHeaders(v map[string]*string) *ListContactFlowsResponse {
	s.Headers = v
	return s
}

func (s *ListContactFlowsResponse) SetBody(v *ListContactFlowsResponseBody) *ListContactFlowsResponse {
	s.Body = v
	return s
}

type ListPersonalNumbersOfUserRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	IsMember      *bool   `json:"IsMember,omitempty" xml:"IsMember,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s ListPersonalNumbersOfUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPersonalNumbersOfUserRequest) GoString() string {
	return s.String()
}

func (s *ListPersonalNumbersOfUserRequest) SetInstanceId(v string) *ListPersonalNumbersOfUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListPersonalNumbersOfUserRequest) SetUserId(v string) *ListPersonalNumbersOfUserRequest {
	s.UserId = &v
	return s
}

func (s *ListPersonalNumbersOfUserRequest) SetPageNumber(v int32) *ListPersonalNumbersOfUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPersonalNumbersOfUserRequest) SetPageSize(v int32) *ListPersonalNumbersOfUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListPersonalNumbersOfUserRequest) SetIsMember(v bool) *ListPersonalNumbersOfUserRequest {
	s.IsMember = &v
	return s
}

func (s *ListPersonalNumbersOfUserRequest) SetSearchPattern(v string) *ListPersonalNumbersOfUserRequest {
	s.SearchPattern = &v
	return s
}

type ListPersonalNumbersOfUserResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListPersonalNumbersOfUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListPersonalNumbersOfUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPersonalNumbersOfUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonalNumbersOfUserResponseBody) SetCode(v string) *ListPersonalNumbersOfUserResponseBody {
	s.Code = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBody) SetHttpStatusCode(v int32) *ListPersonalNumbersOfUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBody) SetMessage(v string) *ListPersonalNumbersOfUserResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBody) SetRequestId(v string) *ListPersonalNumbersOfUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBody) SetData(v *ListPersonalNumbersOfUserResponseBodyData) *ListPersonalNumbersOfUserResponseBody {
	s.Data = v
	return s
}

type ListPersonalNumbersOfUserResponseBodyData struct {
	PageNumber *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListPersonalNumbersOfUserResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListPersonalNumbersOfUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPersonalNumbersOfUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPersonalNumbersOfUserResponseBodyData) SetPageNumber(v int32) *ListPersonalNumbersOfUserResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyData) SetPageSize(v int32) *ListPersonalNumbersOfUserResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyData) SetTotalCount(v int32) *ListPersonalNumbersOfUserResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyData) SetList(v []*ListPersonalNumbersOfUserResponseBodyDataList) *ListPersonalNumbersOfUserResponseBodyData {
	s.List = v
	return s
}

type ListPersonalNumbersOfUserResponseBodyDataList struct {
	Active        *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty"`
	City          *string `json:"City,omitempty" xml:"City,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	Province      *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListPersonalNumbersOfUserResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListPersonalNumbersOfUserResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) SetActive(v bool) *ListPersonalNumbersOfUserResponseBodyDataList {
	s.Active = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) SetNumber(v string) *ListPersonalNumbersOfUserResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) SetCity(v string) *ListPersonalNumbersOfUserResponseBodyDataList {
	s.City = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) SetInstanceId(v string) *ListPersonalNumbersOfUserResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) SetContactFlowId(v string) *ListPersonalNumbersOfUserResponseBodyDataList {
	s.ContactFlowId = &v
	return s
}

func (s *ListPersonalNumbersOfUserResponseBodyDataList) SetProvince(v string) *ListPersonalNumbersOfUserResponseBodyDataList {
	s.Province = &v
	return s
}

type ListPersonalNumbersOfUserResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPersonalNumbersOfUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPersonalNumbersOfUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPersonalNumbersOfUserResponse) GoString() string {
	return s.String()
}

func (s *ListPersonalNumbersOfUserResponse) SetHeaders(v map[string]*string) *ListPersonalNumbersOfUserResponse {
	s.Headers = v
	return s
}

func (s *ListPersonalNumbersOfUserResponse) SetBody(v *ListPersonalNumbersOfUserResponseBody) *ListPersonalNumbersOfUserResponse {
	s.Body = v
	return s
}

type StartPredictiveCallRequest struct {
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Caller               *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Callee               *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	TimeoutSeconds       *int32  `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	ContactFlowId        *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ContactFlowVariables *string `json:"ContactFlowVariables,omitempty" xml:"ContactFlowVariables,omitempty"`
}

func (s StartPredictiveCallRequest) String() string {
	return tea.Prettify(s)
}

func (s StartPredictiveCallRequest) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallRequest) SetInstanceId(v string) *StartPredictiveCallRequest {
	s.InstanceId = &v
	return s
}

func (s *StartPredictiveCallRequest) SetCaller(v string) *StartPredictiveCallRequest {
	s.Caller = &v
	return s
}

func (s *StartPredictiveCallRequest) SetCallee(v string) *StartPredictiveCallRequest {
	s.Callee = &v
	return s
}

func (s *StartPredictiveCallRequest) SetTimeoutSeconds(v int32) *StartPredictiveCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *StartPredictiveCallRequest) SetContactFlowId(v string) *StartPredictiveCallRequest {
	s.ContactFlowId = &v
	return s
}

func (s *StartPredictiveCallRequest) SetTags(v string) *StartPredictiveCallRequest {
	s.Tags = &v
	return s
}

func (s *StartPredictiveCallRequest) SetContactFlowVariables(v string) *StartPredictiveCallRequest {
	s.ContactFlowVariables = &v
	return s
}

type StartPredictiveCallResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                            `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *StartPredictiveCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s StartPredictiveCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartPredictiveCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBody) SetCode(v string) *StartPredictiveCallResponseBody {
	s.Code = &v
	return s
}

func (s *StartPredictiveCallResponseBody) SetHttpStatusCode(v int32) *StartPredictiveCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartPredictiveCallResponseBody) SetMessage(v string) *StartPredictiveCallResponseBody {
	s.Message = &v
	return s
}

func (s *StartPredictiveCallResponseBody) SetRequestId(v string) *StartPredictiveCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPredictiveCallResponseBody) SetParams(v []*string) *StartPredictiveCallResponseBody {
	s.Params = v
	return s
}

func (s *StartPredictiveCallResponseBody) SetData(v *StartPredictiveCallResponseBodyData) *StartPredictiveCallResponseBody {
	s.Data = v
	return s
}

type StartPredictiveCallResponseBodyData struct {
	CallContext *StartPredictiveCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *StartPredictiveCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s StartPredictiveCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StartPredictiveCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBodyData) SetCallContext(v *StartPredictiveCallResponseBodyDataCallContext) *StartPredictiveCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *StartPredictiveCallResponseBodyData) SetUserContext(v *StartPredictiveCallResponseBodyDataUserContext) *StartPredictiveCallResponseBodyData {
	s.UserContext = v
	return s
}

type StartPredictiveCallResponseBodyDataCallContext struct {
	CallType        *string                                                          `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                          `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*StartPredictiveCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s StartPredictiveCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s StartPredictiveCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBodyDataCallContext) SetCallType(v string) *StartPredictiveCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContext) SetInstanceId(v string) *StartPredictiveCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContext) SetJobId(v string) *StartPredictiveCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContext) SetChannelContexts(v []*StartPredictiveCallResponseBodyDataCallContextChannelContexts) *StartPredictiveCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type StartPredictiveCallResponseBodyDataCallContextChannelContexts struct {
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
}

func (s StartPredictiveCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s StartPredictiveCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *StartPredictiveCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

type StartPredictiveCallResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	DeviceState            *string   `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s StartPredictiveCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s StartPredictiveCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetExtension(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetHeartbeat(v int64) *StartPredictiveCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetWorkMode(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetDeviceId(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetUserId(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetReserved(v int64) *StartPredictiveCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetBreakCode(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetInstanceId(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *StartPredictiveCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetDeviceState(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.DeviceState = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetMobile(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetJobId(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetUserState(v string) *StartPredictiveCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *StartPredictiveCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *StartPredictiveCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type StartPredictiveCallResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartPredictiveCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartPredictiveCallResponse) String() string {
	return tea.Prettify(s)
}

func (s StartPredictiveCallResponse) GoString() string {
	return s.String()
}

func (s *StartPredictiveCallResponse) SetHeaders(v map[string]*string) *StartPredictiveCallResponse {
	s.Headers = v
	return s
}

func (s *StartPredictiveCallResponse) SetBody(v *StartPredictiveCallResponseBody) *StartPredictiveCallResponse {
	s.Body = v
	return s
}

type ListIntervalInstanceReportRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s ListIntervalInstanceReportRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalInstanceReportRequest) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportRequest) SetInstanceId(v string) *ListIntervalInstanceReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntervalInstanceReportRequest) SetStartTime(v int64) *ListIntervalInstanceReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListIntervalInstanceReportRequest) SetEndTime(v int64) *ListIntervalInstanceReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListIntervalInstanceReportRequest) SetInterval(v string) *ListIntervalInstanceReportRequest {
	s.Interval = &v
	return s
}

type ListIntervalInstanceReportResponseBody struct {
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*ListIntervalInstanceReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListIntervalInstanceReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalInstanceReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportResponseBody) SetCode(v string) *ListIntervalInstanceReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBody) SetHttpStatusCode(v int32) *ListIntervalInstanceReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBody) SetMessage(v string) *ListIntervalInstanceReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBody) SetRequestId(v string) *ListIntervalInstanceReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBody) SetData(v []*ListIntervalInstanceReportResponseBodyData) *ListIntervalInstanceReportResponseBody {
	s.Data = v
	return s
}

type ListIntervalInstanceReportResponseBodyData struct {
	StatsTime *int64                                              `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	Inbound   *ListIntervalInstanceReportResponseBodyDataInbound  `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Outbound  *ListIntervalInstanceReportResponseBodyDataOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall   *ListIntervalInstanceReportResponseBodyDataOverall  `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
}

func (s ListIntervalInstanceReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalInstanceReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportResponseBodyData) SetStatsTime(v int64) *ListIntervalInstanceReportResponseBodyData {
	s.StatsTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyData) SetInbound(v *ListIntervalInstanceReportResponseBodyDataInbound) *ListIntervalInstanceReportResponseBodyData {
	s.Inbound = v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyData) SetOutbound(v *ListIntervalInstanceReportResponseBodyDataOutbound) *ListIntervalInstanceReportResponseBodyData {
	s.Outbound = v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyData) SetOverall(v *ListIntervalInstanceReportResponseBodyDataOverall) *ListIntervalInstanceReportResponseBodyData {
	s.Overall = v
	return s
}

type ListIntervalInstanceReportResponseBodyDataInbound struct {
	AverageRingTime                *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	CallsVoicemail                 *int64   `json:"CallsVoicemail,omitempty" xml:"CallsVoicemail,omitempty"`
	MaxAbandonedInIVRTime          *int64   `json:"MaxAbandonedInIVRTime,omitempty" xml:"MaxAbandonedInIVRTime,omitempty"`
	CallsHandled                   *int64   `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	CallsIVRException              *int64   `json:"CallsIVRException,omitempty" xml:"CallsIVRException,omitempty"`
	CallsAbandonedInIVR            *int64   `json:"CallsAbandonedInIVR,omitempty" xml:"CallsAbandonedInIVR,omitempty"`
	MaxWorkTime                    *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	TotalHoldTime                  *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	MaxAbandonTime                 *int64   `json:"MaxAbandonTime,omitempty" xml:"MaxAbandonTime,omitempty"`
	AverageAbandonTime             *float32 `json:"AverageAbandonTime,omitempty" xml:"AverageAbandonTime,omitempty"`
	AbandonedRate                  *float32 `json:"AbandonedRate,omitempty" xml:"AbandonedRate,omitempty"`
	CallsRinged                    *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	CallsQueuingFailed             *int64   `json:"CallsQueuingFailed,omitempty" xml:"CallsQueuingFailed,omitempty"`
	TotalRingTime                  *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	TotalTalkTime                  *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	MaxAbandonedInRingTime         *int64   `json:"MaxAbandonedInRingTime,omitempty" xml:"MaxAbandonedInRingTime,omitempty"`
	CallsBlindTransferred          *int64   `json:"CallsBlindTransferred,omitempty" xml:"CallsBlindTransferred,omitempty"`
	AverageAbandonedInIVRTime      *float32 `json:"AverageAbandonedInIVRTime,omitempty" xml:"AverageAbandonedInIVRTime,omitempty"`
	AverageAbandonedInQueueTime    *float32 `json:"AverageAbandonedInQueueTime,omitempty" xml:"AverageAbandonedInQueueTime,omitempty"`
	MaxWaitTime                    *int64   `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	AverageTalkTime                *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsAttendedTransferred       *int64   `json:"CallsAttendedTransferred,omitempty" xml:"CallsAttendedTransferred,omitempty"`
	TotalAbandonedInIVRTime        *int64   `json:"TotalAbandonedInIVRTime,omitempty" xml:"TotalAbandonedInIVRTime,omitempty"`
	CallsQueuingOverflow           *int64   `json:"CallsQueuingOverflow,omitempty" xml:"CallsQueuingOverflow,omitempty"`
	CallsAbandonedInRing           *int64   `json:"CallsAbandonedInRing,omitempty" xml:"CallsAbandonedInRing,omitempty"`
	TotalAbandonedInRingTime       *int64   `json:"TotalAbandonedInRingTime,omitempty" xml:"TotalAbandonedInRingTime,omitempty"`
	TotalWorkTime                  *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	AverageWaitTime                *float32 `json:"AverageWaitTime,omitempty" xml:"AverageWaitTime,omitempty"`
	AverageWorkTime                *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsQueued                    *int64   `json:"CallsQueued,omitempty" xml:"CallsQueued,omitempty"`
	AverageAbandonedInRingTime     *float32 `json:"AverageAbandonedInRingTime,omitempty" xml:"AverageAbandonedInRingTime,omitempty"`
	SatisfactionIndex              *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	CallsAbandoned                 *int64   `json:"CallsAbandoned,omitempty" xml:"CallsAbandoned,omitempty"`
	MaxAbandonedInQueueTime        *int64   `json:"MaxAbandonedInQueueTime,omitempty" xml:"MaxAbandonedInQueueTime,omitempty"`
	CallsAbandonedInVoiceNavigator *int64   `json:"CallsAbandonedInVoiceNavigator,omitempty" xml:"CallsAbandonedInVoiceNavigator,omitempty"`
	TotalWaitTime                  *int64   `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	MaxTalkTime                    *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxRingTime                    *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	TotalAbandonTime               *int64   `json:"TotalAbandonTime,omitempty" xml:"TotalAbandonTime,omitempty"`
	CallsOffered                   *int64   `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	CallsQueuingTimeout            *int64   `json:"CallsQueuingTimeout,omitempty" xml:"CallsQueuingTimeout,omitempty"`
	ServiceLevel20                 *float32 `json:"ServiceLevel20,omitempty" xml:"ServiceLevel20,omitempty"`
	MaxHoldTime                    *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	CallsForwardToOutsideNumber    *int64   `json:"CallsForwardToOutsideNumber,omitempty" xml:"CallsForwardToOutsideNumber,omitempty"`
	SatisfactionRate               *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	CallsHold                      *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	SatisfactionSurveysOffered     *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	HandleRate                     *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	SatisfactionSurveysResponded   *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime                *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	CallsAbandonedInQueue          *int64   `json:"CallsAbandonedInQueue,omitempty" xml:"CallsAbandonedInQueue,omitempty"`
	TotalAbandonedInQueueTime      *int64   `json:"TotalAbandonedInQueueTime,omitempty" xml:"TotalAbandonedInQueueTime,omitempty"`
}

func (s ListIntervalInstanceReportResponseBodyDataInbound) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalInstanceReportResponseBodyDataInbound) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageRingTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsVoicemail(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsVoicemail = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxAbandonedInIVRTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonedInIVRTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsHandled(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsIVRException(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsIVRException = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInIVR(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInIVR = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxWorkTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalHoldTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxAbandonTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageAbandonTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAbandonedRate(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AbandonedRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsRinged(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsQueuingFailed(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsQueuingFailed = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalRingTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalTalkTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxAbandonedInRingTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonedInRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsBlindTransferred(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsBlindTransferred = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageAbandonedInIVRTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonedInIVRTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageAbandonedInQueueTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonedInQueueTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxWaitTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxWaitTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageTalkTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsAttendedTransferred(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsAttendedTransferred = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalAbandonedInIVRTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonedInIVRTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsQueuingOverflow(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsQueuingOverflow = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInRing(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInRing = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalAbandonedInRingTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonedInRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalWorkTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageWaitTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageWaitTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageWorkTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsQueued(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsQueued = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageAbandonedInRingTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonedInRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetSatisfactionIndex(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsAbandoned(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsAbandoned = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxAbandonedInQueueTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonedInQueueTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInVoiceNavigator(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInVoiceNavigator = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalWaitTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalWaitTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxTalkTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxRingTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalAbandonTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsOffered(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsQueuingTimeout(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsQueuingTimeout = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetServiceLevel20(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.ServiceLevel20 = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetMaxHoldTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsForwardToOutsideNumber(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsForwardToOutsideNumber = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetSatisfactionRate(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsHold(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetHandleRate(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.HandleRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetAverageHoldTime(v float32) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInQueue(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInQueue = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataInbound) SetTotalAbandonedInQueueTime(v int64) *ListIntervalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonedInQueueTime = &v
	return s
}

type ListIntervalInstanceReportResponseBodyDataOutbound struct {
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	CallsDialed                  *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsAnswered                *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	TotalDialingTime             *int64   `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	MaxDialingTime               *int64   `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	CallsBlindTransferred        *int64   `json:"CallsBlindTransferred,omitempty" xml:"CallsBlindTransferred,omitempty"`
	AnswerRate                   *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	CallsAttendedTransferred     *int64   `json:"CallsAttendedTransferred,omitempty" xml:"CallsAttendedTransferred,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageDialingTime           *float32 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
}

func (s ListIntervalInstanceReportResponseBodyDataOutbound) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalInstanceReportResponseBodyDataOutbound) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetAverageRingTime(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetCallsDialed(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetCallsAnswered(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetTotalWorkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetMaxWorkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetTotalDialingTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetTotalHoldTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetAverageWorkTime(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetSatisfactionIndex(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetCallsRinged(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetTotalRingTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetMaxTalkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetMaxRingTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetTotalTalkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetMaxDialingTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetCallsBlindTransferred(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.CallsBlindTransferred = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetAnswerRate(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetMaxHoldTime(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetAverageTalkTime(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetSatisfactionRate(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetCallsAttendedTransferred(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.CallsAttendedTransferred = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetCallsHold(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetAverageHoldTime(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOutbound) SetAverageDialingTime(v float32) *ListIntervalInstanceReportResponseBodyDataOutbound {
	s.AverageDialingTime = &v
	return s
}

type ListIntervalInstanceReportResponseBodyDataOverall struct {
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalLoggedInTime            *int64   `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	OccupancyRate                *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	AverageBreakTime             *float32 `json:"AverageBreakTime,omitempty" xml:"AverageBreakTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	MaxBreakTime                 *int64   `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	MaxReadyTime                 *int64   `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	AverageReadyTime             *float32 `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	TotalReadyTime               *int64   `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	TotalBreakTime               *int64   `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalCalls                   *int64   `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
}

func (s ListIntervalInstanceReportResponseBodyDataOverall) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalInstanceReportResponseBodyDataOverall) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalTalkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalLoggedInTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetOccupancyRate(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalWorkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetMaxHoldTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetMaxWorkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetAverageBreakTime(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalHoldTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetSatisfactionRate(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetMaxBreakTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetAverageWorkTime(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetAverageTalkTime(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetSatisfactionIndex(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetSatisfactionSurveysOffered(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetSatisfactionSurveysResponded(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetMaxReadyTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetAverageReadyTime(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetAverageHoldTime(v float32) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalReadyTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalBreakTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetMaxTalkTime(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalInstanceReportResponseBodyDataOverall) SetTotalCalls(v int64) *ListIntervalInstanceReportResponseBodyDataOverall {
	s.TotalCalls = &v
	return s
}

type ListIntervalInstanceReportResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListIntervalInstanceReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIntervalInstanceReportResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalInstanceReportResponse) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportResponse) SetHeaders(v map[string]*string) *ListIntervalInstanceReportResponse {
	s.Headers = v
	return s
}

func (s *ListIntervalInstanceReportResponse) SetBody(v *ListIntervalInstanceReportResponseBody) *ListIntervalInstanceReportResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	NumberList     *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	AdminRamIdList *string `json:"AdminRamIdList,omitempty" xml:"AdminRamIdList,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetName(v string) *CreateInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetDomainName(v string) *CreateInstanceRequest {
	s.DomainName = &v
	return s
}

func (s *CreateInstanceRequest) SetNumberList(v string) *CreateInstanceRequest {
	s.NumberList = &v
	return s
}

func (s *CreateInstanceRequest) SetAdminRamIdList(v string) *CreateInstanceRequest {
	s.AdminRamIdList = &v
	return s
}

type CreateInstanceResponseBody struct {
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Data           *string   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetHttpStatusCode(v int32) *CreateInstanceResponseBody {
	s.HttpStatusCode = &v
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

func (s *CreateInstanceResponseBody) SetData(v string) *CreateInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetParams(v []*string) *CreateInstanceResponseBody {
	s.Params = v
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

type RemoveSkillGroupsFromUserRequest struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
}

func (s RemoveSkillGroupsFromUserRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveSkillGroupsFromUserRequest) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupsFromUserRequest) SetInstanceId(v string) *RemoveSkillGroupsFromUserRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveSkillGroupsFromUserRequest) SetUserId(v string) *RemoveSkillGroupsFromUserRequest {
	s.UserId = &v
	return s
}

func (s *RemoveSkillGroupsFromUserRequest) SetSkillGroupIdList(v string) *RemoveSkillGroupsFromUserRequest {
	s.SkillGroupIdList = &v
	return s
}

type RemoveSkillGroupsFromUserResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveSkillGroupsFromUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSkillGroupsFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupsFromUserResponseBody) SetHttpStatusCode(v int32) *RemoveSkillGroupsFromUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveSkillGroupsFromUserResponseBody) SetCode(v string) *RemoveSkillGroupsFromUserResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveSkillGroupsFromUserResponseBody) SetMessage(v string) *RemoveSkillGroupsFromUserResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveSkillGroupsFromUserResponseBody) SetData(v string) *RemoveSkillGroupsFromUserResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveSkillGroupsFromUserResponseBody) SetRequestId(v string) *RemoveSkillGroupsFromUserResponseBody {
	s.RequestId = &v
	return s
}

type RemoveSkillGroupsFromUserResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveSkillGroupsFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSkillGroupsFromUserResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSkillGroupsFromUserResponse) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupsFromUserResponse) SetHeaders(v map[string]*string) *RemoveSkillGroupsFromUserResponse {
	s.Headers = v
	return s
}

func (s *RemoveSkillGroupsFromUserResponse) SetBody(v *RemoveSkillGroupsFromUserResponseBody) *RemoveSkillGroupsFromUserResponse {
	s.Body = v
	return s
}

type ListRealtimeAgentStatesRequest struct {
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AgentIdList  *string `json:"AgentIdList,omitempty" xml:"AgentIdList,omitempty"`
	StateList    *string `json:"StateList,omitempty" xml:"StateList,omitempty"`
	AgentName    *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s ListRealtimeAgentStatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeAgentStatesRequest) GoString() string {
	return s.String()
}

func (s *ListRealtimeAgentStatesRequest) SetSkillGroupId(v string) *ListRealtimeAgentStatesRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetPageNumber(v int32) *ListRealtimeAgentStatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetPageSize(v int32) *ListRealtimeAgentStatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetInstanceId(v string) *ListRealtimeAgentStatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetAgentIdList(v string) *ListRealtimeAgentStatesRequest {
	s.AgentIdList = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetStateList(v string) *ListRealtimeAgentStatesRequest {
	s.StateList = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetAgentName(v string) *ListRealtimeAgentStatesRequest {
	s.AgentName = &v
	return s
}

type ListRealtimeAgentStatesResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListRealtimeAgentStatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListRealtimeAgentStatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeAgentStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRealtimeAgentStatesResponseBody) SetCode(v string) *ListRealtimeAgentStatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBody) SetHttpStatusCode(v int32) *ListRealtimeAgentStatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBody) SetMessage(v string) *ListRealtimeAgentStatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBody) SetRequestId(v string) *ListRealtimeAgentStatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBody) SetData(v *ListRealtimeAgentStatesResponseBodyData) *ListRealtimeAgentStatesResponseBody {
	s.Data = v
	return s
}

type ListRealtimeAgentStatesResponseBodyData struct {
	PageNumber *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListRealtimeAgentStatesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListRealtimeAgentStatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeAgentStatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRealtimeAgentStatesResponseBodyData) SetPageNumber(v int32) *ListRealtimeAgentStatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyData) SetPageSize(v int32) *ListRealtimeAgentStatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyData) SetTotalCount(v int32) *ListRealtimeAgentStatesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyData) SetList(v []*ListRealtimeAgentStatesResponseBodyDataList) *ListRealtimeAgentStatesResponseBodyData {
	s.List = v
	return s
}

type ListRealtimeAgentStatesResponseBodyDataList struct {
	Extension        *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	StateCode        *string   `json:"StateCode,omitempty" xml:"StateCode,omitempty"`
	State            *string   `json:"State,omitempty" xml:"State,omitempty"`
	AgentId          *string   `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	StateTime        *int64    `json:"StateTime,omitempty" xml:"StateTime,omitempty"`
	AgentName        *string   `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	InstanceId       *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CounterParty     *string   `json:"CounterParty,omitempty" xml:"CounterParty,omitempty"`
	SkillGroupIdList []*string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty" type:"Repeated"`
}

func (s ListRealtimeAgentStatesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeAgentStatesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetExtension(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.Extension = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetStateCode(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.StateCode = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetState(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.State = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetAgentId(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetStateTime(v int64) *ListRealtimeAgentStatesResponseBodyDataList {
	s.StateTime = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetAgentName(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetInstanceId(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetCounterParty(v string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.CounterParty = &v
	return s
}

func (s *ListRealtimeAgentStatesResponseBodyDataList) SetSkillGroupIdList(v []*string) *ListRealtimeAgentStatesResponseBodyDataList {
	s.SkillGroupIdList = v
	return s
}

type ListRealtimeAgentStatesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRealtimeAgentStatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRealtimeAgentStatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeAgentStatesResponse) GoString() string {
	return s.String()
}

func (s *ListRealtimeAgentStatesResponse) SetHeaders(v map[string]*string) *ListRealtimeAgentStatesResponse {
	s.Headers = v
	return s
}

func (s *ListRealtimeAgentStatesResponse) SetBody(v *ListRealtimeAgentStatesResponseBody) *ListRealtimeAgentStatesResponse {
	s.Body = v
	return s
}

type LaunchAuthenticationRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId      *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
}

func (s LaunchAuthenticationRequest) String() string {
	return tea.Prettify(s)
}

func (s LaunchAuthenticationRequest) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationRequest) SetInstanceId(v string) *LaunchAuthenticationRequest {
	s.InstanceId = &v
	return s
}

func (s *LaunchAuthenticationRequest) SetUserId(v string) *LaunchAuthenticationRequest {
	s.UserId = &v
	return s
}

func (s *LaunchAuthenticationRequest) SetDeviceId(v string) *LaunchAuthenticationRequest {
	s.DeviceId = &v
	return s
}

func (s *LaunchAuthenticationRequest) SetJobId(v string) *LaunchAuthenticationRequest {
	s.JobId = &v
	return s
}

func (s *LaunchAuthenticationRequest) SetContactFlowId(v string) *LaunchAuthenticationRequest {
	s.ContactFlowId = &v
	return s
}

type LaunchAuthenticationResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                             `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *LaunchAuthenticationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s LaunchAuthenticationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LaunchAuthenticationResponseBody) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationResponseBody) SetCode(v string) *LaunchAuthenticationResponseBody {
	s.Code = &v
	return s
}

func (s *LaunchAuthenticationResponseBody) SetHttpStatusCode(v int32) *LaunchAuthenticationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *LaunchAuthenticationResponseBody) SetMessage(v string) *LaunchAuthenticationResponseBody {
	s.Message = &v
	return s
}

func (s *LaunchAuthenticationResponseBody) SetRequestId(v string) *LaunchAuthenticationResponseBody {
	s.RequestId = &v
	return s
}

func (s *LaunchAuthenticationResponseBody) SetParams(v []*string) *LaunchAuthenticationResponseBody {
	s.Params = v
	return s
}

func (s *LaunchAuthenticationResponseBody) SetData(v *LaunchAuthenticationResponseBodyData) *LaunchAuthenticationResponseBody {
	s.Data = v
	return s
}

type LaunchAuthenticationResponseBodyData struct {
	CallContext *LaunchAuthenticationResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *LaunchAuthenticationResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s LaunchAuthenticationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LaunchAuthenticationResponseBodyData) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationResponseBodyData) SetCallContext(v *LaunchAuthenticationResponseBodyDataCallContext) *LaunchAuthenticationResponseBodyData {
	s.CallContext = v
	return s
}

func (s *LaunchAuthenticationResponseBodyData) SetUserContext(v *LaunchAuthenticationResponseBodyDataUserContext) *LaunchAuthenticationResponseBodyData {
	s.UserContext = v
	return s
}

type LaunchAuthenticationResponseBodyDataCallContext struct {
	CallType        *string                                                           `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                           `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*LaunchAuthenticationResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s LaunchAuthenticationResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s LaunchAuthenticationResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) SetCallType(v string) *LaunchAuthenticationResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) SetInstanceId(v string) *LaunchAuthenticationResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) SetJobId(v string) *LaunchAuthenticationResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContext) SetChannelContexts(v []*LaunchAuthenticationResponseBodyDataCallContextChannelContexts) *LaunchAuthenticationResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type LaunchAuthenticationResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s LaunchAuthenticationResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s LaunchAuthenticationResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetDestination(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetUserId(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetCallType(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetJobId(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *LaunchAuthenticationResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type LaunchAuthenticationResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s LaunchAuthenticationResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s LaunchAuthenticationResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetExtension(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetHeartbeat(v int64) *LaunchAuthenticationResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetWorkMode(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetDeviceId(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetUserId(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetReserved(v int64) *LaunchAuthenticationResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetBreakCode(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetInstanceId(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetOutboundScenario(v bool) *LaunchAuthenticationResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetMobile(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetJobId(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetUserState(v string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *LaunchAuthenticationResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *LaunchAuthenticationResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type LaunchAuthenticationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LaunchAuthenticationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LaunchAuthenticationResponse) String() string {
	return tea.Prettify(s)
}

func (s LaunchAuthenticationResponse) GoString() string {
	return s.String()
}

func (s *LaunchAuthenticationResponse) SetHeaders(v map[string]*string) *LaunchAuthenticationResponse {
	s.Headers = v
	return s
}

func (s *LaunchAuthenticationResponse) SetBody(v *LaunchAuthenticationResponseBody) *LaunchAuthenticationResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

type ListInstancesResponseBody struct {
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                      `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *ListInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetParams(v []*string) *ListInstancesResponseBody {
	s.Params = v
	return s
}

func (s *ListInstancesResponseBody) SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

type ListInstancesResponseBodyData struct {
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListInstancesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) SetPageNumber(v int32) *ListInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageSize(v int32) *ListInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetTotalCount(v int32) *ListInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetList(v []*ListInstancesResponseBodyDataList) *ListInstancesResponseBodyData {
	s.List = v
	return s
}

type ListInstancesResponseBodyDataList struct {
	Status      *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	ConsoleUrl  *string                                        `json:"ConsoleUrl,omitempty" xml:"ConsoleUrl,omitempty"`
	Description *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	CreateTime  *int64                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AliyunUid   *string                                        `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	Name        *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	DomainName  *string                                        `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Id          *string                                        `json:"Id,omitempty" xml:"Id,omitempty"`
	AdminList   []*ListInstancesResponseBodyDataListAdminList  `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	NumberList  []*ListInstancesResponseBodyDataListNumberList `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataList) SetStatus(v string) *ListInstancesResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetConsoleUrl(v string) *ListInstancesResponseBodyDataList {
	s.ConsoleUrl = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetDescription(v string) *ListInstancesResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetCreateTime(v int64) *ListInstancesResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetAliyunUid(v string) *ListInstancesResponseBodyDataList {
	s.AliyunUid = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetName(v string) *ListInstancesResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetDomainName(v string) *ListInstancesResponseBodyDataList {
	s.DomainName = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetId(v string) *ListInstancesResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetAdminList(v []*ListInstancesResponseBodyDataListAdminList) *ListInstancesResponseBodyDataList {
	s.AdminList = v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetNumberList(v []*ListInstancesResponseBodyDataListNumberList) *ListInstancesResponseBodyDataList {
	s.NumberList = v
	return s
}

type ListInstancesResponseBodyDataListAdminList struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Extension   *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	LoginName   *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	WorkMode    *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RoleName    *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleId      *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s ListInstancesResponseBodyDataListAdminList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataListAdminList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataListAdminList) SetDisplayName(v string) *ListInstancesResponseBodyDataListAdminList {
	s.DisplayName = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetExtension(v string) *ListInstancesResponseBodyDataListAdminList {
	s.Extension = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetLoginName(v string) *ListInstancesResponseBodyDataListAdminList {
	s.LoginName = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetEmail(v string) *ListInstancesResponseBodyDataListAdminList {
	s.Email = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetWorkMode(v string) *ListInstancesResponseBodyDataListAdminList {
	s.WorkMode = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetMobile(v string) *ListInstancesResponseBodyDataListAdminList {
	s.Mobile = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetUserId(v string) *ListInstancesResponseBodyDataListAdminList {
	s.UserId = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetRoleName(v string) *ListInstancesResponseBodyDataListAdminList {
	s.RoleName = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetInstanceId(v string) *ListInstancesResponseBodyDataListAdminList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyDataListAdminList) SetRoleId(v string) *ListInstancesResponseBodyDataListAdminList {
	s.RoleId = &v
	return s
}

type ListInstancesResponseBodyDataListNumberList struct {
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s ListInstancesResponseBodyDataListNumberList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataListNumberList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataListNumberList) SetNumber(v string) *ListInstancesResponseBodyDataListNumberList {
	s.Number = &v
	return s
}

type ListInstancesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type GetHistoricalInstanceReportRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s GetHistoricalInstanceReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalInstanceReportRequest) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportRequest) SetInstanceId(v string) *GetHistoricalInstanceReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHistoricalInstanceReportRequest) SetStartTime(v int64) *GetHistoricalInstanceReportRequest {
	s.StartTime = &v
	return s
}

func (s *GetHistoricalInstanceReportRequest) SetEndTime(v int64) *GetHistoricalInstanceReportRequest {
	s.EndTime = &v
	return s
}

type GetHistoricalInstanceReportResponseBody struct {
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetHistoricalInstanceReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetHistoricalInstanceReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalInstanceReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponseBody) SetCode(v string) *GetHistoricalInstanceReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBody) SetHttpStatusCode(v int32) *GetHistoricalInstanceReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBody) SetMessage(v string) *GetHistoricalInstanceReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBody) SetRequestId(v string) *GetHistoricalInstanceReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBody) SetData(v *GetHistoricalInstanceReportResponseBodyData) *GetHistoricalInstanceReportResponseBody {
	s.Data = v
	return s
}

type GetHistoricalInstanceReportResponseBodyData struct {
	Inbound  *GetHistoricalInstanceReportResponseBodyDataInbound  `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Outbound *GetHistoricalInstanceReportResponseBodyDataOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall  *GetHistoricalInstanceReportResponseBodyDataOverall  `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
}

func (s GetHistoricalInstanceReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalInstanceReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponseBodyData) SetInbound(v *GetHistoricalInstanceReportResponseBodyDataInbound) *GetHistoricalInstanceReportResponseBodyData {
	s.Inbound = v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyData) SetOutbound(v *GetHistoricalInstanceReportResponseBodyDataOutbound) *GetHistoricalInstanceReportResponseBodyData {
	s.Outbound = v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyData) SetOverall(v *GetHistoricalInstanceReportResponseBodyDataOverall) *GetHistoricalInstanceReportResponseBodyData {
	s.Overall = v
	return s
}

type GetHistoricalInstanceReportResponseBodyDataInbound struct {
	AverageRingTime                *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	CallsVoicemail                 *int64   `json:"CallsVoicemail,omitempty" xml:"CallsVoicemail,omitempty"`
	MaxAbandonedInIVRTime          *int64   `json:"MaxAbandonedInIVRTime,omitempty" xml:"MaxAbandonedInIVRTime,omitempty"`
	CallsHandled                   *int64   `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	CallsIVRException              *int64   `json:"CallsIVRException,omitempty" xml:"CallsIVRException,omitempty"`
	CallsAbandonedInIVR            *int64   `json:"CallsAbandonedInIVR,omitempty" xml:"CallsAbandonedInIVR,omitempty"`
	MaxWorkTime                    *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	TotalHoldTime                  *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	MaxAbandonTime                 *int64   `json:"MaxAbandonTime,omitempty" xml:"MaxAbandonTime,omitempty"`
	AverageAbandonTime             *float32 `json:"AverageAbandonTime,omitempty" xml:"AverageAbandonTime,omitempty"`
	CallsRinged                    *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	CallsQueuingFailed             *int64   `json:"CallsQueuingFailed,omitempty" xml:"CallsQueuingFailed,omitempty"`
	TotalRingTime                  *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	AbandonRate                    *float32 `json:"AbandonRate,omitempty" xml:"AbandonRate,omitempty"`
	TotalTalkTime                  *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	MaxAbandonedInRingTime         *int64   `json:"MaxAbandonedInRingTime,omitempty" xml:"MaxAbandonedInRingTime,omitempty"`
	CallsBlindTransferred          *int64   `json:"CallsBlindTransferred,omitempty" xml:"CallsBlindTransferred,omitempty"`
	AverageAbandonedInIVRTime      *float32 `json:"AverageAbandonedInIVRTime,omitempty" xml:"AverageAbandonedInIVRTime,omitempty"`
	AverageAbandonedInQueueTime    *float32 `json:"AverageAbandonedInQueueTime,omitempty" xml:"AverageAbandonedInQueueTime,omitempty"`
	MaxWaitTime                    *int64   `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	AverageTalkTime                *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsAttendedTransferred       *int64   `json:"CallsAttendedTransferred,omitempty" xml:"CallsAttendedTransferred,omitempty"`
	TotalAbandonedInIVRTime        *int64   `json:"TotalAbandonedInIVRTime,omitempty" xml:"TotalAbandonedInIVRTime,omitempty"`
	CallsQueuingOverflow           *int64   `json:"CallsQueuingOverflow,omitempty" xml:"CallsQueuingOverflow,omitempty"`
	CallsAbandonedInRing           *int64   `json:"CallsAbandonedInRing,omitempty" xml:"CallsAbandonedInRing,omitempty"`
	TotalAbandonedInRingTime       *int64   `json:"TotalAbandonedInRingTime,omitempty" xml:"TotalAbandonedInRingTime,omitempty"`
	TotalWorkTime                  *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	AverageWaitTime                *float32 `json:"AverageWaitTime,omitempty" xml:"AverageWaitTime,omitempty"`
	AverageWorkTime                *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsQueued                    *int64   `json:"CallsQueued,omitempty" xml:"CallsQueued,omitempty"`
	AverageAbandonedInRingTime     *float32 `json:"AverageAbandonedInRingTime,omitempty" xml:"AverageAbandonedInRingTime,omitempty"`
	SatisfactionIndex              *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	CallsAbandoned                 *int64   `json:"CallsAbandoned,omitempty" xml:"CallsAbandoned,omitempty"`
	MaxAbandonedInQueueTime        *int64   `json:"MaxAbandonedInQueueTime,omitempty" xml:"MaxAbandonedInQueueTime,omitempty"`
	CallsAbandonedInVoiceNavigator *int64   `json:"CallsAbandonedInVoiceNavigator,omitempty" xml:"CallsAbandonedInVoiceNavigator,omitempty"`
	TotalWaitTime                  *int64   `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	MaxTalkTime                    *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxRingTime                    *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	TotalAbandonTime               *int64   `json:"TotalAbandonTime,omitempty" xml:"TotalAbandonTime,omitempty"`
	CallsOffered                   *int64   `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	CallsQueuingTimeout            *int64   `json:"CallsQueuingTimeout,omitempty" xml:"CallsQueuingTimeout,omitempty"`
	ServiceLevel20                 *float32 `json:"ServiceLevel20,omitempty" xml:"ServiceLevel20,omitempty"`
	MaxHoldTime                    *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	CallsForwardToOutsideNumber    *int64   `json:"CallsForwardToOutsideNumber,omitempty" xml:"CallsForwardToOutsideNumber,omitempty"`
	SatisfactionRate               *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	CallsHold                      *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	SatisfactionSurveysOffered     *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	HandleRate                     *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	SatisfactionSurveysResponded   *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime                *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	CallsAbandonedInQueue          *int64   `json:"CallsAbandonedInQueue,omitempty" xml:"CallsAbandonedInQueue,omitempty"`
	TotalAbandonedInQueueTime      *int64   `json:"TotalAbandonedInQueueTime,omitempty" xml:"TotalAbandonedInQueueTime,omitempty"`
}

func (s GetHistoricalInstanceReportResponseBodyDataInbound) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalInstanceReportResponseBodyDataInbound) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageRingTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsVoicemail(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsVoicemail = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxAbandonedInIVRTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonedInIVRTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsHandled(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsHandled = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsIVRException(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsIVRException = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInIVR(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInIVR = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxWorkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalHoldTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxAbandonTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageAbandonTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsRinged(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsRinged = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsQueuingFailed(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsQueuingFailed = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalRingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAbandonRate(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AbandonRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalTalkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxAbandonedInRingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonedInRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsBlindTransferred(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsBlindTransferred = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageAbandonedInIVRTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonedInIVRTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageAbandonedInQueueTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonedInQueueTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxWaitTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxWaitTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageTalkTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsAttendedTransferred(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsAttendedTransferred = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalAbandonedInIVRTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonedInIVRTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsQueuingOverflow(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsQueuingOverflow = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInRing(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInRing = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalAbandonedInRingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonedInRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalWorkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageWaitTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageWaitTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageWorkTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsQueued(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsQueued = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageAbandonedInRingTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageAbandonedInRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetSatisfactionIndex(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsAbandoned(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsAbandoned = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxAbandonedInQueueTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxAbandonedInQueueTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInVoiceNavigator(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInVoiceNavigator = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalWaitTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalWaitTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxTalkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxRingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalAbandonTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsOffered(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsOffered = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsQueuingTimeout(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsQueuingTimeout = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetServiceLevel20(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.ServiceLevel20 = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetMaxHoldTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsForwardToOutsideNumber(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsForwardToOutsideNumber = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetSatisfactionRate(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsHold(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsHold = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetSatisfactionSurveysOffered(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetHandleRate(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.HandleRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetSatisfactionSurveysResponded(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetAverageHoldTime(v float32) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetCallsAbandonedInQueue(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.CallsAbandonedInQueue = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataInbound) SetTotalAbandonedInQueueTime(v int64) *GetHistoricalInstanceReportResponseBodyDataInbound {
	s.TotalAbandonedInQueueTime = &v
	return s
}

type GetHistoricalInstanceReportResponseBodyDataOutbound struct {
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	CallsDialed                  *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsAnswered                *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	TotalDialingTime             *int64   `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	MaxDialingTime               *int64   `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	CallsBlindTransferred        *int64   `json:"CallsBlindTransferred,omitempty" xml:"CallsBlindTransferred,omitempty"`
	AnswerRate                   *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	CallsAttendedTransferred     *int64   `json:"CallsAttendedTransferred,omitempty" xml:"CallsAttendedTransferred,omitempty"`
	CallsHold                    *int32   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageDialingTime           *float32 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
}

func (s GetHistoricalInstanceReportResponseBodyDataOutbound) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalInstanceReportResponseBodyDataOutbound) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetAverageRingTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetCallsDialed(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.CallsDialed = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetCallsAnswered(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetTotalWorkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetMaxWorkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetTotalDialingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetTotalHoldTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetAverageWorkTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetSatisfactionIndex(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetCallsRinged(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.CallsRinged = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetTotalRingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetMaxTalkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetMaxRingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetTotalTalkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetMaxDialingTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetCallsBlindTransferred(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.CallsBlindTransferred = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetAnswerRate(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.AnswerRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetMaxHoldTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetAverageTalkTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetSatisfactionRate(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetCallsAttendedTransferred(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.CallsAttendedTransferred = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetCallsHold(v int32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.CallsHold = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetSatisfactionSurveysOffered(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetSatisfactionSurveysResponded(v int64) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetAverageHoldTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOutbound) SetAverageDialingTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOutbound {
	s.AverageDialingTime = &v
	return s
}

type GetHistoricalInstanceReportResponseBodyDataOverall struct {
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalLoggedInTime            *int64   `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	OccupancyRate                *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	AverageBreakTime             *float32 `json:"AverageBreakTime,omitempty" xml:"AverageBreakTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	MaxBreakTime                 *int64   `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	MaxReadyTime                 *int64   `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	AverageReadyTime             *float32 `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	TotalReadyTime               *int64   `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	TotalBreakTime               *int64   `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalCalls                   *int64   `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
}

func (s GetHistoricalInstanceReportResponseBodyDataOverall) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalInstanceReportResponseBodyDataOverall) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalTalkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalLoggedInTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetOccupancyRate(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.OccupancyRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalWorkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetMaxHoldTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetMaxWorkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetAverageBreakTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalHoldTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetSatisfactionRate(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetMaxBreakTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetAverageWorkTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetAverageTalkTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetSatisfactionIndex(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetSatisfactionSurveysOffered(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetSatisfactionSurveysResponded(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetMaxReadyTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetAverageReadyTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetAverageHoldTime(v float32) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalReadyTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalBreakTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetMaxTalkTime(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *GetHistoricalInstanceReportResponseBodyDataOverall) SetTotalCalls(v int64) *GetHistoricalInstanceReportResponseBodyDataOverall {
	s.TotalCalls = &v
	return s
}

type GetHistoricalInstanceReportResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHistoricalInstanceReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHistoricalInstanceReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalInstanceReportResponse) GoString() string {
	return s.String()
}

func (s *GetHistoricalInstanceReportResponse) SetHeaders(v map[string]*string) *GetHistoricalInstanceReportResponse {
	s.Headers = v
	return s
}

func (s *GetHistoricalInstanceReportResponse) SetBody(v *GetHistoricalInstanceReportResponseBody) *GetHistoricalInstanceReportResponse {
	s.Body = v
	return s
}

type RemoveUsersRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserIdList *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s RemoveUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersRequest) SetInstanceId(v string) *RemoveUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveUsersRequest) SetUserIdList(v string) *RemoveUsersRequest {
	s.UserIdList = &v
	return s
}

type RemoveUsersResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s RemoveUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBody) SetCode(v string) *RemoveUsersResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveUsersResponseBody) SetHttpStatusCode(v int32) *RemoveUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveUsersResponseBody) SetMessage(v string) *RemoveUsersResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveUsersResponseBody) SetRequestId(v string) *RemoveUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUsersResponseBody) SetParams(v []*string) *RemoveUsersResponseBody {
	s.Params = v
	return s
}

type RemoveUsersResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponse) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponse) SetHeaders(v map[string]*string) *RemoveUsersResponse {
	s.Headers = v
	return s
}

func (s *RemoveUsersResponse) SetBody(v *RemoveUsersResponseBody) *RemoveUsersResponse {
	s.Body = v
	return s
}

type StartBack2BackCallRequest struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Caller           *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Callee           *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	TimeoutSeconds   *int32  `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	Broker           *string `json:"Broker,omitempty" xml:"Broker,omitempty"`
	AdditionalBroker *string `json:"AdditionalBroker,omitempty" xml:"AdditionalBroker,omitempty"`
	Tags             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s StartBack2BackCallRequest) String() string {
	return tea.Prettify(s)
}

func (s StartBack2BackCallRequest) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallRequest) SetInstanceId(v string) *StartBack2BackCallRequest {
	s.InstanceId = &v
	return s
}

func (s *StartBack2BackCallRequest) SetCaller(v string) *StartBack2BackCallRequest {
	s.Caller = &v
	return s
}

func (s *StartBack2BackCallRequest) SetCallee(v string) *StartBack2BackCallRequest {
	s.Callee = &v
	return s
}

func (s *StartBack2BackCallRequest) SetTimeoutSeconds(v int32) *StartBack2BackCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *StartBack2BackCallRequest) SetBroker(v string) *StartBack2BackCallRequest {
	s.Broker = &v
	return s
}

func (s *StartBack2BackCallRequest) SetAdditionalBroker(v string) *StartBack2BackCallRequest {
	s.AdditionalBroker = &v
	return s
}

func (s *StartBack2BackCallRequest) SetTags(v string) *StartBack2BackCallRequest {
	s.Tags = &v
	return s
}

type StartBack2BackCallResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                           `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *StartBack2BackCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s StartBack2BackCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartBack2BackCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallResponseBody) SetCode(v string) *StartBack2BackCallResponseBody {
	s.Code = &v
	return s
}

func (s *StartBack2BackCallResponseBody) SetHttpStatusCode(v int32) *StartBack2BackCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartBack2BackCallResponseBody) SetMessage(v string) *StartBack2BackCallResponseBody {
	s.Message = &v
	return s
}

func (s *StartBack2BackCallResponseBody) SetRequestId(v string) *StartBack2BackCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartBack2BackCallResponseBody) SetParams(v []*string) *StartBack2BackCallResponseBody {
	s.Params = v
	return s
}

func (s *StartBack2BackCallResponseBody) SetData(v *StartBack2BackCallResponseBodyData) *StartBack2BackCallResponseBody {
	s.Data = v
	return s
}

type StartBack2BackCallResponseBodyData struct {
	CallContext *StartBack2BackCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *StartBack2BackCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s StartBack2BackCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StartBack2BackCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallResponseBodyData) SetCallContext(v *StartBack2BackCallResponseBodyDataCallContext) *StartBack2BackCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *StartBack2BackCallResponseBodyData) SetUserContext(v *StartBack2BackCallResponseBodyDataUserContext) *StartBack2BackCallResponseBodyData {
	s.UserContext = v
	return s
}

type StartBack2BackCallResponseBodyDataCallContext struct {
	CallType        *string                                                         `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                         `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*StartBack2BackCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s StartBack2BackCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s StartBack2BackCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallResponseBodyDataCallContext) SetCallType(v string) *StartBack2BackCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContext) SetInstanceId(v string) *StartBack2BackCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContext) SetJobId(v string) *StartBack2BackCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContext) SetChannelContexts(v []*StartBack2BackCallResponseBodyDataCallContextChannelContexts) *StartBack2BackCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type StartBack2BackCallResponseBodyDataCallContextChannelContexts struct {
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
}

func (s StartBack2BackCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s StartBack2BackCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *StartBack2BackCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

type StartBack2BackCallResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Uri                    *string   `json:"Uri,omitempty" xml:"Uri,omitempty"`
	DeviceState            *string   `json:"DeviceState,omitempty" xml:"DeviceState,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s StartBack2BackCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s StartBack2BackCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetExtension(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetHeartbeat(v int64) *StartBack2BackCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetWorkMode(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetDeviceId(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetUserId(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetReserved(v int64) *StartBack2BackCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetBreakCode(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetInstanceId(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *StartBack2BackCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetUri(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.Uri = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetDeviceState(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.DeviceState = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetMobile(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetJobId(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetUserState(v string) *StartBack2BackCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *StartBack2BackCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *StartBack2BackCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type StartBack2BackCallResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartBack2BackCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartBack2BackCallResponse) String() string {
	return tea.Prettify(s)
}

func (s StartBack2BackCallResponse) GoString() string {
	return s.String()
}

func (s *StartBack2BackCallResponse) SetHeaders(v map[string]*string) *StartBack2BackCallResponse {
	s.Headers = v
	return s
}

func (s *StartBack2BackCallResponse) SetBody(v *StartBack2BackCallResponseBody) *StartBack2BackCallResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Extension  *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetInstanceId(v string) *GetUserRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

func (s *GetUserRequest) SetExtension(v string) *GetUserRequest {
	s.Extension = &v
	return s
}

type GetUserResponseBody struct {
	Code           *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *GetUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetCode(v string) *GetUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResponseBody) SetHttpStatusCode(v int32) *GetUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetParams(v []*string) *GetUserResponseBody {
	s.Params = v
	return s
}

func (s *GetUserResponseBody) SetData(v *GetUserResponseBodyData) *GetUserResponseBody {
	s.Data = v
	return s
}

type GetUserResponseBodyData struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Extension   *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	LoginName   *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	WorkMode    *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RoleName    *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleId      *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s GetUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyData) SetDisplayName(v string) *GetUserResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBodyData) SetExtension(v string) *GetUserResponseBodyData {
	s.Extension = &v
	return s
}

func (s *GetUserResponseBodyData) SetLoginName(v string) *GetUserResponseBodyData {
	s.LoginName = &v
	return s
}

func (s *GetUserResponseBodyData) SetEmail(v string) *GetUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyData) SetWorkMode(v string) *GetUserResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *GetUserResponseBodyData) SetMobile(v string) *GetUserResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *GetUserResponseBodyData) SetUserId(v string) *GetUserResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyData) SetRoleName(v string) *GetUserResponseBodyData {
	s.RoleName = &v
	return s
}

func (s *GetUserResponseBodyData) SetInstanceId(v string) *GetUserResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetUserResponseBodyData) SetRoleId(v string) *GetUserResponseBodyData {
	s.RoleId = &v
	return s
}

type GetUserResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type RemovePhoneNumbersFromSkillGroupRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NumberList   *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s RemovePhoneNumbersFromSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemovePhoneNumbersFromSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumbersFromSkillGroupRequest) SetInstanceId(v string) *RemovePhoneNumbersFromSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupRequest) SetNumberList(v string) *RemovePhoneNumbersFromSkillGroupRequest {
	s.NumberList = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupRequest) SetSkillGroupId(v string) *RemovePhoneNumbersFromSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

type RemovePhoneNumbersFromSkillGroupResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePhoneNumbersFromSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemovePhoneNumbersFromSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) SetCode(v string) *RemovePhoneNumbersFromSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) SetHttpStatusCode(v int32) *RemovePhoneNumbersFromSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) SetMessage(v string) *RemovePhoneNumbersFromSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupResponseBody) SetRequestId(v string) *RemovePhoneNumbersFromSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemovePhoneNumbersFromSkillGroupResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemovePhoneNumbersFromSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemovePhoneNumbersFromSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemovePhoneNumbersFromSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumbersFromSkillGroupResponse) SetHeaders(v map[string]*string) *RemovePhoneNumbersFromSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *RemovePhoneNumbersFromSkillGroupResponse) SetBody(v *RemovePhoneNumbersFromSkillGroupResponseBody) *RemovePhoneNumbersFromSkillGroupResponse {
	s.Body = v
	return s
}

type CompleteAttendedTransferRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CompleteAttendedTransferRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteAttendedTransferRequest) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferRequest) SetInstanceId(v string) *CompleteAttendedTransferRequest {
	s.InstanceId = &v
	return s
}

func (s *CompleteAttendedTransferRequest) SetUserId(v string) *CompleteAttendedTransferRequest {
	s.UserId = &v
	return s
}

func (s *CompleteAttendedTransferRequest) SetDeviceId(v string) *CompleteAttendedTransferRequest {
	s.DeviceId = &v
	return s
}

func (s *CompleteAttendedTransferRequest) SetJobId(v string) *CompleteAttendedTransferRequest {
	s.JobId = &v
	return s
}

type CompleteAttendedTransferResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                                 `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *CompleteAttendedTransferResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CompleteAttendedTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompleteAttendedTransferResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBody) SetCode(v string) *CompleteAttendedTransferResponseBody {
	s.Code = &v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetHttpStatusCode(v int32) *CompleteAttendedTransferResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetMessage(v string) *CompleteAttendedTransferResponseBody {
	s.Message = &v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetRequestId(v string) *CompleteAttendedTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetParams(v []*string) *CompleteAttendedTransferResponseBody {
	s.Params = v
	return s
}

func (s *CompleteAttendedTransferResponseBody) SetData(v *CompleteAttendedTransferResponseBodyData) *CompleteAttendedTransferResponseBody {
	s.Data = v
	return s
}

type CompleteAttendedTransferResponseBodyData struct {
	ContextId   *int64                                               `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	CallContext *CompleteAttendedTransferResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *CompleteAttendedTransferResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s CompleteAttendedTransferResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CompleteAttendedTransferResponseBodyData) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBodyData) SetContextId(v int64) *CompleteAttendedTransferResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyData) SetCallContext(v *CompleteAttendedTransferResponseBodyDataCallContext) *CompleteAttendedTransferResponseBodyData {
	s.CallContext = v
	return s
}

func (s *CompleteAttendedTransferResponseBodyData) SetUserContext(v *CompleteAttendedTransferResponseBodyDataUserContext) *CompleteAttendedTransferResponseBodyData {
	s.UserContext = v
	return s
}

type CompleteAttendedTransferResponseBodyDataCallContext struct {
	CallType        *string                                                               `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*CompleteAttendedTransferResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s CompleteAttendedTransferResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s CompleteAttendedTransferResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) SetCallType(v string) *CompleteAttendedTransferResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) SetInstanceId(v string) *CompleteAttendedTransferResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) SetJobId(v string) *CompleteAttendedTransferResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContext) SetChannelContexts(v []*CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) *CompleteAttendedTransferResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type CompleteAttendedTransferResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetDestination(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserId(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetCallType(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetJobId(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *CompleteAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type CompleteAttendedTransferResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s CompleteAttendedTransferResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s CompleteAttendedTransferResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetExtension(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetHeartbeat(v int64) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetWorkMode(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetDeviceId(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetUserId(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetReserved(v int64) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetBreakCode(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetInstanceId(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetOutboundScenario(v bool) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetMobile(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetJobId(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetUserState(v string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *CompleteAttendedTransferResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *CompleteAttendedTransferResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type CompleteAttendedTransferResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CompleteAttendedTransferResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompleteAttendedTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteAttendedTransferResponse) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferResponse) SetHeaders(v map[string]*string) *CompleteAttendedTransferResponse {
	s.Headers = v
	return s
}

func (s *CompleteAttendedTransferResponse) SetBody(v *CompleteAttendedTransferResponseBody) *CompleteAttendedTransferResponse {
	s.Body = v
	return s
}

type ResetUserPasswordRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Password   *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ResetUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordRequest) SetInstanceId(v string) *ResetUserPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetUserPasswordRequest) SetUserId(v string) *ResetUserPasswordRequest {
	s.UserId = &v
	return s
}

func (s *ResetUserPasswordRequest) SetPassword(v string) *ResetUserPasswordRequest {
	s.Password = &v
	return s
}

type ResetUserPasswordResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBody) SetCode(v string) *ResetUserPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetHttpStatusCode(v int32) *ResetUserPasswordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetMessage(v string) *ResetUserPasswordResponseBody {
	s.Message = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetRequestId(v string) *ResetUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetUserPasswordResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponse) SetHeaders(v map[string]*string) *ResetUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetUserPasswordResponse) SetBody(v *ResetUserPasswordResponseBody) *ResetUserPasswordResponse {
	s.Body = v
	return s
}

type GetTurnServerListRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTurnServerListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTurnServerListRequest) GoString() string {
	return s.String()
}

func (s *GetTurnServerListRequest) SetInstanceId(v string) *GetTurnServerListRequest {
	s.InstanceId = &v
	return s
}

type GetTurnServerListResponseBody struct {
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Data           *string   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s GetTurnServerListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTurnServerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTurnServerListResponseBody) SetHttpStatusCode(v int32) *GetTurnServerListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetCode(v string) *GetTurnServerListResponseBody {
	s.Code = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetMessage(v string) *GetTurnServerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetData(v string) *GetTurnServerListResponseBody {
	s.Data = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetRequestId(v string) *GetTurnServerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTurnServerListResponseBody) SetParams(v []*string) *GetTurnServerListResponseBody {
	s.Params = v
	return s
}

type GetTurnServerListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTurnServerListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTurnServerListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTurnServerListResponse) GoString() string {
	return s.String()
}

func (s *GetTurnServerListResponse) SetHeaders(v map[string]*string) *GetTurnServerListResponse {
	s.Headers = v
	return s
}

func (s *GetTurnServerListResponse) SetBody(v *GetTurnServerListResponseBody) *GetTurnServerListResponse {
	s.Body = v
	return s
}

type GetNumberLocationRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Number     *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s GetNumberLocationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNumberLocationRequest) GoString() string {
	return s.String()
}

func (s *GetNumberLocationRequest) SetInstanceId(v string) *GetNumberLocationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNumberLocationRequest) SetNumber(v string) *GetNumberLocationRequest {
	s.Number = &v
	return s
}

type GetNumberLocationResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetNumberLocationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetNumberLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNumberLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetNumberLocationResponseBody) SetCode(v string) *GetNumberLocationResponseBody {
	s.Code = &v
	return s
}

func (s *GetNumberLocationResponseBody) SetHttpStatusCode(v int32) *GetNumberLocationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNumberLocationResponseBody) SetMessage(v string) *GetNumberLocationResponseBody {
	s.Message = &v
	return s
}

func (s *GetNumberLocationResponseBody) SetRequestId(v string) *GetNumberLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNumberLocationResponseBody) SetData(v *GetNumberLocationResponseBodyData) *GetNumberLocationResponseBody {
	s.Data = v
	return s
}

type GetNumberLocationResponseBodyData struct {
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s GetNumberLocationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetNumberLocationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNumberLocationResponseBodyData) SetNumber(v string) *GetNumberLocationResponseBodyData {
	s.Number = &v
	return s
}

func (s *GetNumberLocationResponseBodyData) SetCity(v string) *GetNumberLocationResponseBodyData {
	s.City = &v
	return s
}

func (s *GetNumberLocationResponseBodyData) SetProvince(v string) *GetNumberLocationResponseBodyData {
	s.Province = &v
	return s
}

type GetNumberLocationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetNumberLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNumberLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNumberLocationResponse) GoString() string {
	return s.String()
}

func (s *GetNumberLocationResponse) SetHeaders(v map[string]*string) *GetNumberLocationResponse {
	s.Headers = v
	return s
}

func (s *GetNumberLocationResponse) SetBody(v *GetNumberLocationResponseBody) *GetNumberLocationResponse {
	s.Body = v
	return s
}

type ListRamUsersRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRamUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRamUsersRequest) GoString() string {
	return s.String()
}

func (s *ListRamUsersRequest) SetInstanceId(v string) *ListRamUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRamUsersRequest) SetSearchPattern(v string) *ListRamUsersRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListRamUsersRequest) SetPageNumber(v int32) *ListRamUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRamUsersRequest) SetPageSize(v int32) *ListRamUsersRequest {
	s.PageSize = &v
	return s
}

type ListRamUsersResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                     `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *ListRamUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListRamUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRamUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRamUsersResponseBody) SetCode(v string) *ListRamUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListRamUsersResponseBody) SetHttpStatusCode(v int32) *ListRamUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRamUsersResponseBody) SetMessage(v string) *ListRamUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListRamUsersResponseBody) SetRequestId(v string) *ListRamUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRamUsersResponseBody) SetParams(v []*string) *ListRamUsersResponseBody {
	s.Params = v
	return s
}

func (s *ListRamUsersResponseBody) SetData(v *ListRamUsersResponseBodyData) *ListRamUsersResponseBody {
	s.Data = v
	return s
}

type ListRamUsersResponseBodyData struct {
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListRamUsersResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListRamUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRamUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRamUsersResponseBodyData) SetPageNumber(v int32) *ListRamUsersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListRamUsersResponseBodyData) SetPageSize(v int32) *ListRamUsersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListRamUsersResponseBodyData) SetTotalCount(v int32) *ListRamUsersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListRamUsersResponseBodyData) SetList(v []*ListRamUsersResponseBodyDataList) *ListRamUsersResponseBodyData {
	s.List = v
	return s
}

type ListRamUsersResponseBodyDataList struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	LoginName   *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	AliyunUid   *int64  `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	Primary     *bool   `json:"Primary,omitempty" xml:"Primary,omitempty"`
	RamId       *string `json:"RamId,omitempty" xml:"RamId,omitempty"`
}

func (s ListRamUsersResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListRamUsersResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListRamUsersResponseBodyDataList) SetDisplayName(v string) *ListRamUsersResponseBodyDataList {
	s.DisplayName = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) SetEmail(v string) *ListRamUsersResponseBodyDataList {
	s.Email = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) SetLoginName(v string) *ListRamUsersResponseBodyDataList {
	s.LoginName = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) SetMobile(v string) *ListRamUsersResponseBodyDataList {
	s.Mobile = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) SetAliyunUid(v int64) *ListRamUsersResponseBodyDataList {
	s.AliyunUid = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) SetPrimary(v bool) *ListRamUsersResponseBodyDataList {
	s.Primary = &v
	return s
}

func (s *ListRamUsersResponseBodyDataList) SetRamId(v string) *ListRamUsersResponseBodyDataList {
	s.RamId = &v
	return s
}

type ListRamUsersResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRamUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRamUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRamUsersResponse) GoString() string {
	return s.String()
}

func (s *ListRamUsersResponse) SetHeaders(v map[string]*string) *ListRamUsersResponse {
	s.Headers = v
	return s
}

func (s *ListRamUsersResponse) SetBody(v *ListRamUsersResponseBody) *ListRamUsersResponse {
	s.Body = v
	return s
}

type MuteCallRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s MuteCallRequest) String() string {
	return tea.Prettify(s)
}

func (s MuteCallRequest) GoString() string {
	return s.String()
}

func (s *MuteCallRequest) SetInstanceId(v string) *MuteCallRequest {
	s.InstanceId = &v
	return s
}

func (s *MuteCallRequest) SetUserId(v string) *MuteCallRequest {
	s.UserId = &v
	return s
}

func (s *MuteCallRequest) SetDeviceId(v string) *MuteCallRequest {
	s.DeviceId = &v
	return s
}

func (s *MuteCallRequest) SetJobId(v string) *MuteCallRequest {
	s.JobId = &v
	return s
}

func (s *MuteCallRequest) SetChannelId(v string) *MuteCallRequest {
	s.ChannelId = &v
	return s
}

type MuteCallResponseBody struct {
	Code           *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                 `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *MuteCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s MuteCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MuteCallResponseBody) GoString() string {
	return s.String()
}

func (s *MuteCallResponseBody) SetCode(v string) *MuteCallResponseBody {
	s.Code = &v
	return s
}

func (s *MuteCallResponseBody) SetHttpStatusCode(v int32) *MuteCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MuteCallResponseBody) SetMessage(v string) *MuteCallResponseBody {
	s.Message = &v
	return s
}

func (s *MuteCallResponseBody) SetRequestId(v string) *MuteCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *MuteCallResponseBody) SetParams(v []*string) *MuteCallResponseBody {
	s.Params = v
	return s
}

func (s *MuteCallResponseBody) SetData(v *MuteCallResponseBodyData) *MuteCallResponseBody {
	s.Data = v
	return s
}

type MuteCallResponseBodyData struct {
	CallContext *MuteCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *MuteCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s MuteCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MuteCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *MuteCallResponseBodyData) SetCallContext(v *MuteCallResponseBodyDataCallContext) *MuteCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *MuteCallResponseBodyData) SetUserContext(v *MuteCallResponseBodyDataUserContext) *MuteCallResponseBodyData {
	s.UserContext = v
	return s
}

type MuteCallResponseBodyDataCallContext struct {
	CallType        *string                                               `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*MuteCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s MuteCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s MuteCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *MuteCallResponseBodyDataCallContext) SetCallType(v string) *MuteCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContext) SetInstanceId(v string) *MuteCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContext) SetJobId(v string) *MuteCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContext) SetChannelContexts(v []*MuteCallResponseBodyDataCallContextChannelContexts) *MuteCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type MuteCallResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s MuteCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s MuteCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *MuteCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *MuteCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type MuteCallResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s MuteCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s MuteCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *MuteCallResponseBodyDataUserContext) SetExtension(v string) *MuteCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetHeartbeat(v int64) *MuteCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetWorkMode(v string) *MuteCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetDeviceId(v string) *MuteCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetUserId(v string) *MuteCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetReserved(v int64) *MuteCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetBreakCode(v string) *MuteCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetInstanceId(v string) *MuteCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *MuteCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetMobile(v string) *MuteCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetJobId(v string) *MuteCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetUserState(v string) *MuteCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *MuteCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *MuteCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type MuteCallResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MuteCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MuteCallResponse) String() string {
	return tea.Prettify(s)
}

func (s MuteCallResponse) GoString() string {
	return s.String()
}

func (s *MuteCallResponse) SetHeaders(v map[string]*string) *MuteCallResponse {
	s.Headers = v
	return s
}

func (s *MuteCallResponse) SetBody(v *MuteCallResponseBody) *MuteCallResponse {
	s.Body = v
	return s
}

type AnswerCallRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s AnswerCallRequest) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallRequest) GoString() string {
	return s.String()
}

func (s *AnswerCallRequest) SetInstanceId(v string) *AnswerCallRequest {
	s.InstanceId = &v
	return s
}

func (s *AnswerCallRequest) SetUserId(v string) *AnswerCallRequest {
	s.UserId = &v
	return s
}

func (s *AnswerCallRequest) SetDeviceId(v string) *AnswerCallRequest {
	s.DeviceId = &v
	return s
}

func (s *AnswerCallRequest) SetJobId(v string) *AnswerCallRequest {
	s.JobId = &v
	return s
}

type AnswerCallResponseBody struct {
	Code           *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                   `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *AnswerCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AnswerCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallResponseBody) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBody) SetCode(v string) *AnswerCallResponseBody {
	s.Code = &v
	return s
}

func (s *AnswerCallResponseBody) SetHttpStatusCode(v int32) *AnswerCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AnswerCallResponseBody) SetMessage(v string) *AnswerCallResponseBody {
	s.Message = &v
	return s
}

func (s *AnswerCallResponseBody) SetRequestId(v string) *AnswerCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnswerCallResponseBody) SetParams(v []*string) *AnswerCallResponseBody {
	s.Params = v
	return s
}

func (s *AnswerCallResponseBody) SetData(v *AnswerCallResponseBodyData) *AnswerCallResponseBody {
	s.Data = v
	return s
}

type AnswerCallResponseBodyData struct {
	ContextId   *int64                                 `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	CallContext *AnswerCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *AnswerCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s AnswerCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBodyData) SetContextId(v int64) *AnswerCallResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *AnswerCallResponseBodyData) SetCallContext(v *AnswerCallResponseBodyDataCallContext) *AnswerCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *AnswerCallResponseBodyData) SetUserContext(v *AnswerCallResponseBodyDataUserContext) *AnswerCallResponseBodyData {
	s.UserContext = v
	return s
}

type AnswerCallResponseBodyDataCallContext struct {
	JobId           *string                                                 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	InstanceId      *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ChannelContexts []*AnswerCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s AnswerCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBodyDataCallContext) SetJobId(v string) *AnswerCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContext) SetInstanceId(v string) *AnswerCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContext) SetChannelContexts(v []*AnswerCallResponseBodyDataCallContextChannelContexts) *AnswerCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type AnswerCallResponseBodyDataCallContextChannelContexts struct {
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
}

func (s AnswerCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *AnswerCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *AnswerCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

type AnswerCallResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s AnswerCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBodyDataUserContext) SetExtension(v string) *AnswerCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetHeartbeat(v int64) *AnswerCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetWorkMode(v string) *AnswerCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetDeviceId(v string) *AnswerCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetUserId(v string) *AnswerCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetReserved(v int64) *AnswerCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetBreakCode(v string) *AnswerCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetInstanceId(v string) *AnswerCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *AnswerCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetMobile(v string) *AnswerCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetJobId(v string) *AnswerCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetUserState(v string) *AnswerCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *AnswerCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *AnswerCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
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

type ListIntervalAgentReportRequest struct {
	AgentId    *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Interval   *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s ListIntervalAgentReportRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentReportRequest) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportRequest) SetAgentId(v string) *ListIntervalAgentReportRequest {
	s.AgentId = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetStartTime(v int64) *ListIntervalAgentReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetEndTime(v int64) *ListIntervalAgentReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetInstanceId(v string) *ListIntervalAgentReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetInterval(v string) *ListIntervalAgentReportRequest {
	s.Interval = &v
	return s
}

type ListIntervalAgentReportResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*ListIntervalAgentReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListIntervalAgentReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBody) SetCode(v string) *ListIntervalAgentReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntervalAgentReportResponseBody) SetHttpStatusCode(v int32) *ListIntervalAgentReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntervalAgentReportResponseBody) SetMessage(v string) *ListIntervalAgentReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntervalAgentReportResponseBody) SetRequestId(v string) *ListIntervalAgentReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntervalAgentReportResponseBody) SetData(v []*ListIntervalAgentReportResponseBodyData) *ListIntervalAgentReportResponseBody {
	s.Data = v
	return s
}

type ListIntervalAgentReportResponseBodyData struct {
	StatsTime *int64                                           `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
	Inbound   *ListIntervalAgentReportResponseBodyDataInbound  `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Outbound  *ListIntervalAgentReportResponseBodyDataOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall   *ListIntervalAgentReportResponseBodyDataOverall  `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
}

func (s ListIntervalAgentReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBodyData) SetStatsTime(v int64) *ListIntervalAgentReportResponseBodyData {
	s.StatsTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyData) SetInbound(v *ListIntervalAgentReportResponseBodyDataInbound) *ListIntervalAgentReportResponseBodyData {
	s.Inbound = v
	return s
}

func (s *ListIntervalAgentReportResponseBodyData) SetOutbound(v *ListIntervalAgentReportResponseBodyDataOutbound) *ListIntervalAgentReportResponseBodyData {
	s.Outbound = v
	return s
}

func (s *ListIntervalAgentReportResponseBodyData) SetOverall(v *ListIntervalAgentReportResponseBodyDataOverall) *ListIntervalAgentReportResponseBodyData {
	s.Overall = v
	return s
}

type ListIntervalAgentReportResponseBodyDataInbound struct {
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	CallsHandled                 *int64   `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	CallsAttendedTransferOut     *int64   `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsBlindTransferIn         *int64   `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	CallsAttendedTransferIn      *int64   `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	CallsBlindTransferOut        *int64   `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	CallsOffered                 *int64   `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	HandleRate                   *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
}

func (s ListIntervalAgentReportResponseBodyDataInbound) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentReportResponseBodyDataInbound) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetAverageRingTime(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsHandled(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetTotalWorkTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsAttendedTransferOut(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetMaxWorkTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetTotalHoldTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetAverageWorkTime(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsBlindTransferIn(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetSatisfactionIndex(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsRinged(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsAttendedTransferIn(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsBlindTransferOut(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetTotalRingTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetMaxTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetMaxRingTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetTotalTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsOffered(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetMaxHoldTime(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetAverageTalkTime(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetSatisfactionRate(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetCallsHold(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetHandleRate(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.HandleRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalAgentReportResponseBodyDataInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataInbound) SetAverageHoldTime(v float32) *ListIntervalAgentReportResponseBodyDataInbound {
	s.AverageHoldTime = &v
	return s
}

type ListIntervalAgentReportResponseBodyDataOutbound struct {
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	CallsDialed                  *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsAnswered                *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	CallsAttendedTransferOut     *int64   `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	TotalDialingTime             *int64   `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsBlindTransferIn         *int64   `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	CallsAttendedTransferIn      *int64   `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	CallsBlindTransferOut        *int64   `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	MaxDialingTime               *int64   `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	AnswerRate                   *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageDialingTime           *float32 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
}

func (s ListIntervalAgentReportResponseBodyDataOutbound) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentReportResponseBodyDataOutbound) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetAverageRingTime(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsDialed(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsAnswered(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetTotalWorkTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsAttendedTransferOut(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetMaxWorkTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetTotalDialingTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetTotalHoldTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetAverageWorkTime(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsBlindTransferIn(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetSatisfactionIndex(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsRinged(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsAttendedTransferIn(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsBlindTransferOut(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetTotalRingTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetMaxTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetMaxRingTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetTotalTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetMaxDialingTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetAnswerRate(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetMaxHoldTime(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetAverageTalkTime(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetSatisfactionRate(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetCallsHold(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetAverageHoldTime(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOutbound) SetAverageDialingTime(v float32) *ListIntervalAgentReportResponseBodyDataOutbound {
	s.AverageDialingTime = &v
	return s
}

type ListIntervalAgentReportResponseBodyDataOverall struct {
	OccupancyRate                *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	MaxBreakTime                 *int64   `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	MaxReadyTime                 *int64   `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalReadyTime               *int64   `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	LastCheckoutTime             *int64   `json:"LastCheckoutTime,omitempty" xml:"LastCheckoutTime,omitempty"`
	TotalCalls                   *int64   `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	TotalLoggedInTime            *int64   `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	AverageBreakTime             *float32 `json:"AverageBreakTime,omitempty" xml:"AverageBreakTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	FirstCheckInTime             *int64   `json:"FirstCheckInTime,omitempty" xml:"FirstCheckInTime,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageReadyTime             *float32 `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	TotalBreakTime               *int64   `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
}

func (s ListIntervalAgentReportResponseBodyDataOverall) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentReportResponseBodyDataOverall) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetOccupancyRate(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalWorkTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetMaxWorkTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalHoldTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetAverageWorkTime(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetMaxBreakTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetSatisfactionIndex(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetMaxReadyTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetMaxTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalReadyTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetLastCheckoutTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.LastCheckoutTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalCalls(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalCalls = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalLoggedInTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalTalkTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetMaxHoldTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetAverageBreakTime(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetAverageTalkTime(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetSatisfactionRate(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetSatisfactionSurveysOffered(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetFirstCheckInTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.FirstCheckInTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetSatisfactionSurveysResponded(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetAverageHoldTime(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetAverageReadyTime(v float32) *ListIntervalAgentReportResponseBodyDataOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListIntervalAgentReportResponseBodyDataOverall) SetTotalBreakTime(v int64) *ListIntervalAgentReportResponseBodyDataOverall {
	s.TotalBreakTime = &v
	return s
}

type ListIntervalAgentReportResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListIntervalAgentReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIntervalAgentReportResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentReportResponse) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponse) SetHeaders(v map[string]*string) *ListIntervalAgentReportResponse {
	s.Headers = v
	return s
}

func (s *ListIntervalAgentReportResponse) SetBody(v *ListIntervalAgentReportResponseBody) *ListIntervalAgentReportResponse {
	s.Body = v
	return s
}

type ListCallDetailRecordsRequest struct {
	PageNumber                  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime                   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime                     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ContactType                 *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	ContactDisposition          *string `json:"ContactDisposition,omitempty" xml:"ContactDisposition,omitempty"`
	ContactId                   *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	AgentId                     *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	SkillGroupId                *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SortOrder                   *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	InstanceId                  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderByField                *string `json:"OrderByField,omitempty" xml:"OrderByField,omitempty"`
	Criteria                    *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	CallingNumber               *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	CalledNumber                *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	SatisfactionList            *string `json:"SatisfactionList,omitempty" xml:"SatisfactionList,omitempty"`
	SatisfactionSurveyChannel   *string `json:"SatisfactionSurveyChannel,omitempty" xml:"SatisfactionSurveyChannel,omitempty"`
	SatisfactionDescriptionList *string `json:"SatisfactionDescriptionList,omitempty" xml:"SatisfactionDescriptionList,omitempty"`
}

func (s ListCallDetailRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCallDetailRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsRequest) SetPageNumber(v int32) *ListCallDetailRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetPageSize(v int32) *ListCallDetailRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetStartTime(v int64) *ListCallDetailRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetEndTime(v int64) *ListCallDetailRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetContactType(v string) *ListCallDetailRecordsRequest {
	s.ContactType = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetContactDisposition(v string) *ListCallDetailRecordsRequest {
	s.ContactDisposition = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetContactId(v string) *ListCallDetailRecordsRequest {
	s.ContactId = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetAgentId(v string) *ListCallDetailRecordsRequest {
	s.AgentId = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetSkillGroupId(v string) *ListCallDetailRecordsRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetSortOrder(v string) *ListCallDetailRecordsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetInstanceId(v string) *ListCallDetailRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetOrderByField(v string) *ListCallDetailRecordsRequest {
	s.OrderByField = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetCriteria(v string) *ListCallDetailRecordsRequest {
	s.Criteria = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetCallingNumber(v string) *ListCallDetailRecordsRequest {
	s.CallingNumber = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetCalledNumber(v string) *ListCallDetailRecordsRequest {
	s.CalledNumber = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetSatisfactionList(v string) *ListCallDetailRecordsRequest {
	s.SatisfactionList = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetSatisfactionSurveyChannel(v string) *ListCallDetailRecordsRequest {
	s.SatisfactionSurveyChannel = &v
	return s
}

func (s *ListCallDetailRecordsRequest) SetSatisfactionDescriptionList(v string) *ListCallDetailRecordsRequest {
	s.SatisfactionDescriptionList = &v
	return s
}

type ListCallDetailRecordsResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListCallDetailRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListCallDetailRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCallDetailRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsResponseBody) SetCode(v string) *ListCallDetailRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetHttpStatusCode(v int32) *ListCallDetailRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetMessage(v string) *ListCallDetailRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetRequestId(v string) *ListCallDetailRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallDetailRecordsResponseBody) SetData(v *ListCallDetailRecordsResponseBodyData) *ListCallDetailRecordsResponseBody {
	s.Data = v
	return s
}

type ListCallDetailRecordsResponseBodyData struct {
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListCallDetailRecordsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListCallDetailRecordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCallDetailRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsResponseBodyData) SetPageNumber(v int32) *ListCallDetailRecordsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyData) SetPageSize(v int32) *ListCallDetailRecordsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyData) SetTotalCount(v int32) *ListCallDetailRecordsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyData) SetList(v []*ListCallDetailRecordsResponseBodyDataList) *ListCallDetailRecordsResponseBodyData {
	s.List = v
	return s
}

type ListCallDetailRecordsResponseBodyDataList struct {
	ContactDisposition        *string `json:"ContactDisposition,omitempty" xml:"ContactDisposition,omitempty"`
	ContactType               *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	EstablishedTime           *int64  `json:"EstablishedTime,omitempty" xml:"EstablishedTime,omitempty"`
	CalledNumber              *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	AdditionalBroker          *string `json:"AdditionalBroker,omitempty" xml:"AdditionalBroker,omitempty"`
	SatisfactionIndex         *int32  `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	SatisfactionSurveyChannel *string `json:"SatisfactionSurveyChannel,omitempty" xml:"SatisfactionSurveyChannel,omitempty"`
	ReleaseTime               *int64  `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	WaitTime                  *int64  `json:"WaitTime,omitempty" xml:"WaitTime,omitempty"`
	SkillGroupNames           *string `json:"SkillGroupNames,omitempty" xml:"SkillGroupNames,omitempty"`
	IvrTime                   *int64  `json:"IvrTime,omitempty" xml:"IvrTime,omitempty"`
	SatisfactionDescription   *string `json:"SatisfactionDescription,omitempty" xml:"SatisfactionDescription,omitempty"`
	ReleaseInitiator          *string `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	AgentIds                  *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	CallDuration              *string `json:"CallDuration,omitempty" xml:"CallDuration,omitempty"`
	RecordingReady            *bool   `json:"RecordingReady,omitempty" xml:"RecordingReady,omitempty"`
	InstanceId                *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RingTime                  *int64  `json:"RingTime,omitempty" xml:"RingTime,omitempty"`
	SatisfactionSurveyOffered *bool   `json:"SatisfactionSurveyOffered,omitempty" xml:"SatisfactionSurveyOffered,omitempty"`
	AgentNames                *string `json:"AgentNames,omitempty" xml:"AgentNames,omitempty"`
	StartTime                 *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ContactId                 *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	RecordingDuration         *int64  `json:"RecordingDuration,omitempty" xml:"RecordingDuration,omitempty"`
	CallingNumber             *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	QueueTime                 *int64  `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	Broker                    *string `json:"Broker,omitempty" xml:"Broker,omitempty"`
	SkillGroupIds             *string `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty"`
	CallerLocation            *string `json:"CallerLocation,omitempty" xml:"CallerLocation,omitempty"`
	CalleeLocation            *string `json:"CalleeLocation,omitempty" xml:"CalleeLocation,omitempty"`
}

func (s ListCallDetailRecordsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListCallDetailRecordsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetContactDisposition(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.ContactDisposition = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetContactType(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.ContactType = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetEstablishedTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.EstablishedTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetCalledNumber(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.CalledNumber = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetAdditionalBroker(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.AdditionalBroker = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetSatisfactionIndex(v int32) *ListCallDetailRecordsResponseBodyDataList {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetSatisfactionSurveyChannel(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.SatisfactionSurveyChannel = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetReleaseTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.ReleaseTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetWaitTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.WaitTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetSkillGroupNames(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.SkillGroupNames = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetIvrTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.IvrTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetSatisfactionDescription(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.SatisfactionDescription = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetReleaseInitiator(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.ReleaseInitiator = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetAgentIds(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.AgentIds = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetCallDuration(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.CallDuration = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetRecordingReady(v bool) *ListCallDetailRecordsResponseBodyDataList {
	s.RecordingReady = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetInstanceId(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetRingTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.RingTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetSatisfactionSurveyOffered(v bool) *ListCallDetailRecordsResponseBodyDataList {
	s.SatisfactionSurveyOffered = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetAgentNames(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.AgentNames = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetStartTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetContactId(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.ContactId = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetRecordingDuration(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.RecordingDuration = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetCallingNumber(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.CallingNumber = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetQueueTime(v int64) *ListCallDetailRecordsResponseBodyDataList {
	s.QueueTime = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetBroker(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.Broker = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetSkillGroupIds(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.SkillGroupIds = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetCallerLocation(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.CallerLocation = &v
	return s
}

func (s *ListCallDetailRecordsResponseBodyDataList) SetCalleeLocation(v string) *ListCallDetailRecordsResponseBodyDataList {
	s.CalleeLocation = &v
	return s
}

type ListCallDetailRecordsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCallDetailRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCallDetailRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCallDetailRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsResponse) SetHeaders(v map[string]*string) *ListCallDetailRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListCallDetailRecordsResponse) SetBody(v *ListCallDetailRecordsResponseBody) *ListCallDetailRecordsResponse {
	s.Body = v
	return s
}

type RemovePhoneNumbersRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
}

func (s RemovePhoneNumbersRequest) String() string {
	return tea.Prettify(s)
}

func (s RemovePhoneNumbersRequest) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumbersRequest) SetInstanceId(v string) *RemovePhoneNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *RemovePhoneNumbersRequest) SetNumberList(v string) *RemovePhoneNumbersRequest {
	s.NumberList = &v
	return s
}

type RemovePhoneNumbersResponseBody struct {
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Data           *string   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FailureList    []*string `json:"FailureList,omitempty" xml:"FailureList,omitempty" type:"Repeated"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s RemovePhoneNumbersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemovePhoneNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumbersResponseBody) SetHttpStatusCode(v int32) *RemovePhoneNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemovePhoneNumbersResponseBody) SetCode(v string) *RemovePhoneNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *RemovePhoneNumbersResponseBody) SetMessage(v string) *RemovePhoneNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *RemovePhoneNumbersResponseBody) SetData(v string) *RemovePhoneNumbersResponseBody {
	s.Data = &v
	return s
}

func (s *RemovePhoneNumbersResponseBody) SetRequestId(v string) *RemovePhoneNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemovePhoneNumbersResponseBody) SetFailureList(v []*string) *RemovePhoneNumbersResponseBody {
	s.FailureList = v
	return s
}

func (s *RemovePhoneNumbersResponseBody) SetParams(v []*string) *RemovePhoneNumbersResponseBody {
	s.Params = v
	return s
}

type RemovePhoneNumbersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemovePhoneNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemovePhoneNumbersResponse) String() string {
	return tea.Prettify(s)
}

func (s RemovePhoneNumbersResponse) GoString() string {
	return s.String()
}

func (s *RemovePhoneNumbersResponse) SetHeaders(v map[string]*string) *RemovePhoneNumbersResponse {
	s.Headers = v
	return s
}

func (s *RemovePhoneNumbersResponse) SetBody(v *RemovePhoneNumbersResponseBody) *RemovePhoneNumbersResponse {
	s.Body = v
	return s
}

type CancelAttendedTransferRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CancelAttendedTransferRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelAttendedTransferRequest) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferRequest) SetInstanceId(v string) *CancelAttendedTransferRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelAttendedTransferRequest) SetUserId(v string) *CancelAttendedTransferRequest {
	s.UserId = &v
	return s
}

func (s *CancelAttendedTransferRequest) SetDeviceId(v string) *CancelAttendedTransferRequest {
	s.DeviceId = &v
	return s
}

func (s *CancelAttendedTransferRequest) SetJobId(v string) *CancelAttendedTransferRequest {
	s.JobId = &v
	return s
}

type CancelAttendedTransferResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                               `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *CancelAttendedTransferResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CancelAttendedTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelAttendedTransferResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferResponseBody) SetCode(v string) *CancelAttendedTransferResponseBody {
	s.Code = &v
	return s
}

func (s *CancelAttendedTransferResponseBody) SetHttpStatusCode(v int32) *CancelAttendedTransferResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelAttendedTransferResponseBody) SetMessage(v string) *CancelAttendedTransferResponseBody {
	s.Message = &v
	return s
}

func (s *CancelAttendedTransferResponseBody) SetRequestId(v string) *CancelAttendedTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAttendedTransferResponseBody) SetParams(v []*string) *CancelAttendedTransferResponseBody {
	s.Params = v
	return s
}

func (s *CancelAttendedTransferResponseBody) SetData(v *CancelAttendedTransferResponseBodyData) *CancelAttendedTransferResponseBody {
	s.Data = v
	return s
}

type CancelAttendedTransferResponseBodyData struct {
	ContextId   *int64                                             `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	CallContext *CancelAttendedTransferResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *CancelAttendedTransferResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s CancelAttendedTransferResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CancelAttendedTransferResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferResponseBodyData) SetContextId(v int64) *CancelAttendedTransferResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyData) SetCallContext(v *CancelAttendedTransferResponseBodyDataCallContext) *CancelAttendedTransferResponseBodyData {
	s.CallContext = v
	return s
}

func (s *CancelAttendedTransferResponseBodyData) SetUserContext(v *CancelAttendedTransferResponseBodyDataUserContext) *CancelAttendedTransferResponseBodyData {
	s.UserContext = v
	return s
}

type CancelAttendedTransferResponseBodyDataCallContext struct {
	CallType        *string                                                             `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                             `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*CancelAttendedTransferResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s CancelAttendedTransferResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s CancelAttendedTransferResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) SetCallType(v string) *CancelAttendedTransferResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) SetInstanceId(v string) *CancelAttendedTransferResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) SetJobId(v string) *CancelAttendedTransferResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContext) SetChannelContexts(v []*CancelAttendedTransferResponseBodyDataCallContextChannelContexts) *CancelAttendedTransferResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type CancelAttendedTransferResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s CancelAttendedTransferResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s CancelAttendedTransferResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetDestination(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserId(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetCallType(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetJobId(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *CancelAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type CancelAttendedTransferResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s CancelAttendedTransferResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s CancelAttendedTransferResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetExtension(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetHeartbeat(v int64) *CancelAttendedTransferResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetWorkMode(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetDeviceId(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetUserId(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetReserved(v int64) *CancelAttendedTransferResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetBreakCode(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetInstanceId(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetOutboundScenario(v bool) *CancelAttendedTransferResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetMobile(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetJobId(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetUserState(v string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *CancelAttendedTransferResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *CancelAttendedTransferResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type CancelAttendedTransferResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelAttendedTransferResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelAttendedTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelAttendedTransferResponse) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferResponse) SetHeaders(v map[string]*string) *CancelAttendedTransferResponse {
	s.Headers = v
	return s
}

func (s *CancelAttendedTransferResponse) SetBody(v *CancelAttendedTransferResponseBody) *CancelAttendedTransferResponse {
	s.Body = v
	return s
}

type TakeBreakRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s TakeBreakRequest) String() string {
	return tea.Prettify(s)
}

func (s TakeBreakRequest) GoString() string {
	return s.String()
}

func (s *TakeBreakRequest) SetInstanceId(v string) *TakeBreakRequest {
	s.InstanceId = &v
	return s
}

func (s *TakeBreakRequest) SetUserId(v string) *TakeBreakRequest {
	s.UserId = &v
	return s
}

func (s *TakeBreakRequest) SetDeviceId(v string) *TakeBreakRequest {
	s.DeviceId = &v
	return s
}

func (s *TakeBreakRequest) SetCode(v string) *TakeBreakRequest {
	s.Code = &v
	return s
}

type TakeBreakResponseBody struct {
	Code           *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                  `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *TakeBreakResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s TakeBreakResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TakeBreakResponseBody) GoString() string {
	return s.String()
}

func (s *TakeBreakResponseBody) SetCode(v string) *TakeBreakResponseBody {
	s.Code = &v
	return s
}

func (s *TakeBreakResponseBody) SetHttpStatusCode(v int32) *TakeBreakResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TakeBreakResponseBody) SetMessage(v string) *TakeBreakResponseBody {
	s.Message = &v
	return s
}

func (s *TakeBreakResponseBody) SetRequestId(v string) *TakeBreakResponseBody {
	s.RequestId = &v
	return s
}

func (s *TakeBreakResponseBody) SetParams(v []*string) *TakeBreakResponseBody {
	s.Params = v
	return s
}

func (s *TakeBreakResponseBody) SetData(v *TakeBreakResponseBodyData) *TakeBreakResponseBody {
	s.Data = v
	return s
}

type TakeBreakResponseBodyData struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s TakeBreakResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TakeBreakResponseBodyData) GoString() string {
	return s.String()
}

func (s *TakeBreakResponseBodyData) SetExtension(v string) *TakeBreakResponseBodyData {
	s.Extension = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetHeartbeat(v int64) *TakeBreakResponseBodyData {
	s.Heartbeat = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetWorkMode(v string) *TakeBreakResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetDeviceId(v string) *TakeBreakResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetUserId(v string) *TakeBreakResponseBodyData {
	s.UserId = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetReserved(v int64) *TakeBreakResponseBodyData {
	s.Reserved = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetBreakCode(v string) *TakeBreakResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetInstanceId(v string) *TakeBreakResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetOutboundScenario(v bool) *TakeBreakResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetMobile(v string) *TakeBreakResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetJobId(v string) *TakeBreakResponseBodyData {
	s.JobId = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetUserState(v string) *TakeBreakResponseBodyData {
	s.UserState = &v
	return s
}

func (s *TakeBreakResponseBodyData) SetSignedSkillGroupIdList(v []*string) *TakeBreakResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

type TakeBreakResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TakeBreakResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TakeBreakResponse) String() string {
	return tea.Prettify(s)
}

func (s TakeBreakResponse) GoString() string {
	return s.String()
}

func (s *TakeBreakResponse) SetHeaders(v map[string]*string) *TakeBreakResponse {
	s.Headers = v
	return s
}

func (s *TakeBreakResponse) SetBody(v *TakeBreakResponseBody) *TakeBreakResponse {
	s.Body = v
	return s
}

type ListHistoricalSkillGroupReportRequest struct {
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
	StartTime        *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime          *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListHistoricalSkillGroupReportRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalSkillGroupReportRequest) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportRequest) SetPageNumber(v int32) *ListHistoricalSkillGroupReportRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) SetPageSize(v int32) *ListHistoricalSkillGroupReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) SetSkillGroupIdList(v string) *ListHistoricalSkillGroupReportRequest {
	s.SkillGroupIdList = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) SetStartTime(v int64) *ListHistoricalSkillGroupReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) SetEndTime(v int64) *ListHistoricalSkillGroupReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportRequest) SetInstanceId(v string) *ListHistoricalSkillGroupReportRequest {
	s.InstanceId = &v
	return s
}

type ListHistoricalSkillGroupReportResponseBody struct {
	Code           *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListHistoricalSkillGroupReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListHistoricalSkillGroupReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBody) SetCode(v string) *ListHistoricalSkillGroupReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBody) SetHttpStatusCode(v int32) *ListHistoricalSkillGroupReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBody) SetMessage(v string) *ListHistoricalSkillGroupReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBody) SetRequestId(v string) *ListHistoricalSkillGroupReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBody) SetData(v *ListHistoricalSkillGroupReportResponseBodyData) *ListHistoricalSkillGroupReportResponseBody {
	s.Data = v
	return s
}

type ListHistoricalSkillGroupReportResponseBodyData struct {
	PageNumber *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListHistoricalSkillGroupReportResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListHistoricalSkillGroupReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) SetPageNumber(v int32) *ListHistoricalSkillGroupReportResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) SetPageSize(v int32) *ListHistoricalSkillGroupReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) SetTotalCount(v int32) *ListHistoricalSkillGroupReportResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyData) SetList(v []*ListHistoricalSkillGroupReportResponseBodyDataList) *ListHistoricalSkillGroupReportResponseBodyData {
	s.List = v
	return s
}

type ListHistoricalSkillGroupReportResponseBodyDataList struct {
	SkillGroupName *string                                                     `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	SkillGroupId   *string                                                     `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Inbound        *ListHistoricalSkillGroupReportResponseBodyDataListInbound  `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Outbound       *ListHistoricalSkillGroupReportResponseBodyDataListOutbound `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall        *ListHistoricalSkillGroupReportResponseBodyDataListOverall  `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
}

func (s ListHistoricalSkillGroupReportResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) SetSkillGroupName(v string) *ListHistoricalSkillGroupReportResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) SetSkillGroupId(v string) *ListHistoricalSkillGroupReportResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) SetInbound(v *ListHistoricalSkillGroupReportResponseBodyDataListInbound) *ListHistoricalSkillGroupReportResponseBodyDataList {
	s.Inbound = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) SetOutbound(v *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) *ListHistoricalSkillGroupReportResponseBodyDataList {
	s.Outbound = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataList) SetOverall(v *ListHistoricalSkillGroupReportResponseBodyDataListOverall) *ListHistoricalSkillGroupReportResponseBodyDataList {
	s.Overall = v
	return s
}

type ListHistoricalSkillGroupReportResponseBodyDataListInbound struct {
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	CallsOverflow                *int64   `json:"CallsOverflow,omitempty" xml:"CallsOverflow,omitempty"`
	CallsAbandonedInRing         *int64   `json:"CallsAbandonedInRing,omitempty" xml:"CallsAbandonedInRing,omitempty"`
	CallsHandled                 *int64   `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	TotalAbandonedInRingTime     *int64   `json:"TotalAbandonedInRingTime,omitempty" xml:"TotalAbandonedInRingTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	CallsAttendedTransferOut     *int64   `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	AverageWaitTime              *float32 `json:"AverageWaitTime,omitempty" xml:"AverageWaitTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	MaxAbandonTime               *int64   `json:"MaxAbandonTime,omitempty" xml:"MaxAbandonTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsQueued                  *int64   `json:"CallsQueued,omitempty" xml:"CallsQueued,omitempty"`
	CallsBlindTransferIn         *int64   `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	AverageAbandonedInRingTime   *float32 `json:"AverageAbandonedInRingTime,omitempty" xml:"AverageAbandonedInRingTime,omitempty"`
	AverageAbandonTime           *float32 `json:"AverageAbandonTime,omitempty" xml:"AverageAbandonTime,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	CallsBlindTransferOut        *int64   `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	CallsAttendedTransferIn      *int64   `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	CallsAbandoned               *int64   `json:"CallsAbandoned,omitempty" xml:"CallsAbandoned,omitempty"`
	MaxAbandonedInQueueTime      *int64   `json:"MaxAbandonedInQueueTime,omitempty" xml:"MaxAbandonedInQueueTime,omitempty"`
	TotalWaitTime                *int64   `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	AbandonRate                  *float32 `json:"AbandonRate,omitempty" xml:"AbandonRate,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalAbandonTime             *int64   `json:"TotalAbandonTime,omitempty" xml:"TotalAbandonTime,omitempty"`
	CallsOffered                 *int64   `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	MaxAbandonedInRingTime       *int64   `json:"MaxAbandonedInRingTime,omitempty" xml:"MaxAbandonedInRingTime,omitempty"`
	MaxWaitTime                  *int64   `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	AverageAbandonedInQueueTime  *float32 `json:"AverageAbandonedInQueueTime,omitempty" xml:"AverageAbandonedInQueueTime,omitempty"`
	ServiceLevel20               *float32 `json:"ServiceLevel20,omitempty" xml:"ServiceLevel20,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	HandleRate                   *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	CallsTimeout                 *int64   `json:"CallsTimeout,omitempty" xml:"CallsTimeout,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	TotalAbandonedInQueueTime    *int64   `json:"TotalAbandonedInQueueTime,omitempty" xml:"TotalAbandonedInQueueTime,omitempty"`
	CallsAbandonedInQueue        *int64   `json:"CallsAbandonedInQueue,omitempty" xml:"CallsAbandonedInQueue,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListInbound) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListInbound) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageRingTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsOverflow(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsOverflow = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsAbandonedInRing(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsAbandonedInRing = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsHandled(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalWorkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalAbandonedInRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalAbandonedInRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxWorkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsAttendedTransferOut(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageWaitTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageWaitTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalHoldTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxAbandonTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxAbandonTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageWorkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsQueued(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsQueued = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsBlindTransferIn(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetSatisfactionIndex(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageAbandonedInRingTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageAbandonedInRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageAbandonTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageAbandonTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsRinged(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsBlindTransferOut(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsAttendedTransferIn(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsAbandoned(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsAbandoned = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxAbandonedInQueueTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxAbandonedInQueueTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalWaitTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalWaitTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAbandonRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AbandonRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalAbandonTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalAbandonTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsOffered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxAbandonedInRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxAbandonedInRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxWaitTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxWaitTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageAbandonedInQueueTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageAbandonedInQueueTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetServiceLevel20(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.ServiceLevel20 = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetMaxHoldTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetSatisfactionRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageTalkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsHold(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsHold = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetSatisfactionSurveysOffered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetHandleRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.HandleRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsTimeout(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsTimeout = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetSatisfactionSurveysResponded(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetAverageHoldTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetTotalAbandonedInQueueTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.TotalAbandonedInQueueTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListInbound) SetCallsAbandonedInQueue(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListInbound {
	s.CallsAbandonedInQueue = &v
	return s
}

type ListHistoricalSkillGroupReportResponseBodyDataListOutbound struct {
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	CallsDialed                  *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsAnswered                *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	CallsAttendedTransferOut     *int64   `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	TotalDialingTime             *int64   `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsBlindTransferIn         *int64   `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	CallsAttendedTransferIn      *int64   `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	CallsBlindTransferOut        *int64   `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	MaxDialingTime               *int64   `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	AnswerRate                   *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageDialingTime           *float32 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListOutbound) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListOutbound) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetAverageRingTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsDialed(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsAnswered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetTotalWorkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsAttendedTransferOut(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetMaxWorkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetTotalDialingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetTotalHoldTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetAverageWorkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsBlindTransferIn(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionIndex(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsRinged(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsAttendedTransferIn(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsBlindTransferOut(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetTotalRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetMaxTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetMaxRingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetTotalTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetMaxDialingTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetAnswerRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetMaxHoldTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetAverageTalkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetCallsHold(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionSurveysOffered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionSurveysResponded(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetAverageHoldTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOutbound) SetAverageDialingTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOutbound {
	s.AverageDialingTime = &v
	return s
}

type ListHistoricalSkillGroupReportResponseBodyDataListOverall struct {
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalLoggedInTime            *int64   `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	OccupancyRate                *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	AverageBreakTime             *float32 `json:"AverageBreakTime,omitempty" xml:"AverageBreakTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	MaxBreakTime                 *int64   `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	MaxReadyTime                 *int64   `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	AverageReadyTime             *float32 `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	TotalReadyTime               *int64   `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	TotalBreakTime               *int64   `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalCalls                   *int64   `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListOverall) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponseBodyDataListOverall) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalLoggedInTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetOccupancyRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalWorkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetMaxHoldTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetMaxWorkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetAverageBreakTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalHoldTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetSatisfactionRate(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetMaxBreakTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetAverageWorkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetAverageTalkTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetSatisfactionIndex(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetSatisfactionSurveysOffered(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetSatisfactionSurveysResponded(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetMaxReadyTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetAverageReadyTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetAverageHoldTime(v float32) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalReadyTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalBreakTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetMaxTalkTime(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponseBodyDataListOverall) SetTotalCalls(v int64) *ListHistoricalSkillGroupReportResponseBodyDataListOverall {
	s.TotalCalls = &v
	return s
}

type ListHistoricalSkillGroupReportResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHistoricalSkillGroupReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHistoricalSkillGroupReportResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponse) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponse) SetHeaders(v map[string]*string) *ListHistoricalSkillGroupReportResponse {
	s.Headers = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponse) SetBody(v *ListHistoricalSkillGroupReportResponseBody) *ListHistoricalSkillGroupReportResponse {
	s.Body = v
	return s
}

type SendDtmfSignalingRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Dtmf       *string `json:"Dtmf,omitempty" xml:"Dtmf,omitempty"`
}

func (s SendDtmfSignalingRequest) String() string {
	return tea.Prettify(s)
}

func (s SendDtmfSignalingRequest) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingRequest) SetInstanceId(v string) *SendDtmfSignalingRequest {
	s.InstanceId = &v
	return s
}

func (s *SendDtmfSignalingRequest) SetUserId(v string) *SendDtmfSignalingRequest {
	s.UserId = &v
	return s
}

func (s *SendDtmfSignalingRequest) SetDeviceId(v string) *SendDtmfSignalingRequest {
	s.DeviceId = &v
	return s
}

func (s *SendDtmfSignalingRequest) SetJobId(v string) *SendDtmfSignalingRequest {
	s.JobId = &v
	return s
}

func (s *SendDtmfSignalingRequest) SetChannelId(v string) *SendDtmfSignalingRequest {
	s.ChannelId = &v
	return s
}

func (s *SendDtmfSignalingRequest) SetDtmf(v string) *SendDtmfSignalingRequest {
	s.Dtmf = &v
	return s
}

type SendDtmfSignalingResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                          `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *SendDtmfSignalingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s SendDtmfSignalingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendDtmfSignalingResponseBody) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingResponseBody) SetCode(v string) *SendDtmfSignalingResponseBody {
	s.Code = &v
	return s
}

func (s *SendDtmfSignalingResponseBody) SetHttpStatusCode(v int32) *SendDtmfSignalingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SendDtmfSignalingResponseBody) SetMessage(v string) *SendDtmfSignalingResponseBody {
	s.Message = &v
	return s
}

func (s *SendDtmfSignalingResponseBody) SetRequestId(v string) *SendDtmfSignalingResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendDtmfSignalingResponseBody) SetParams(v []*string) *SendDtmfSignalingResponseBody {
	s.Params = v
	return s
}

func (s *SendDtmfSignalingResponseBody) SetData(v *SendDtmfSignalingResponseBodyData) *SendDtmfSignalingResponseBody {
	s.Data = v
	return s
}

type SendDtmfSignalingResponseBodyData struct {
	CallContext *SendDtmfSignalingResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *SendDtmfSignalingResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s SendDtmfSignalingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendDtmfSignalingResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingResponseBodyData) SetCallContext(v *SendDtmfSignalingResponseBodyDataCallContext) *SendDtmfSignalingResponseBodyData {
	s.CallContext = v
	return s
}

func (s *SendDtmfSignalingResponseBodyData) SetUserContext(v *SendDtmfSignalingResponseBodyDataUserContext) *SendDtmfSignalingResponseBodyData {
	s.UserContext = v
	return s
}

type SendDtmfSignalingResponseBodyDataCallContext struct {
	CallType        *string                                                        `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                        `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*SendDtmfSignalingResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s SendDtmfSignalingResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s SendDtmfSignalingResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) SetCallType(v string) *SendDtmfSignalingResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) SetInstanceId(v string) *SendDtmfSignalingResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) SetJobId(v string) *SendDtmfSignalingResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContext) SetChannelContexts(v []*SendDtmfSignalingResponseBodyDataCallContextChannelContexts) *SendDtmfSignalingResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type SendDtmfSignalingResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s SendDtmfSignalingResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s SendDtmfSignalingResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetDestination(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetUserId(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetCallType(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetJobId(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *SendDtmfSignalingResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type SendDtmfSignalingResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s SendDtmfSignalingResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s SendDtmfSignalingResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetExtension(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetHeartbeat(v int64) *SendDtmfSignalingResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetWorkMode(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetDeviceId(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetUserId(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetReserved(v int64) *SendDtmfSignalingResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetBreakCode(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetInstanceId(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetOutboundScenario(v bool) *SendDtmfSignalingResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetMobile(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetJobId(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetUserState(v string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *SendDtmfSignalingResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *SendDtmfSignalingResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type SendDtmfSignalingResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendDtmfSignalingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendDtmfSignalingResponse) String() string {
	return tea.Prettify(s)
}

func (s SendDtmfSignalingResponse) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingResponse) SetHeaders(v map[string]*string) *SendDtmfSignalingResponse {
	s.Headers = v
	return s
}

func (s *SendDtmfSignalingResponse) SetBody(v *SendDtmfSignalingResponseBody) *SendDtmfSignalingResponse {
	s.Body = v
	return s
}

type ListRecentCallDetailRecordsRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Criteria   *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
}

func (s ListRecentCallDetailRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecentCallDetailRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListRecentCallDetailRecordsRequest) SetPageNumber(v int32) *ListRecentCallDetailRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecentCallDetailRecordsRequest) SetPageSize(v int32) *ListRecentCallDetailRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecentCallDetailRecordsRequest) SetStartTime(v int64) *ListRecentCallDetailRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *ListRecentCallDetailRecordsRequest) SetEndTime(v int64) *ListRecentCallDetailRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *ListRecentCallDetailRecordsRequest) SetInstanceId(v string) *ListRecentCallDetailRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRecentCallDetailRecordsRequest) SetCriteria(v string) *ListRecentCallDetailRecordsRequest {
	s.Criteria = &v
	return s
}

type ListRecentCallDetailRecordsResponseBody struct {
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListRecentCallDetailRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListRecentCallDetailRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecentCallDetailRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentCallDetailRecordsResponseBody) SetCode(v string) *ListRecentCallDetailRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBody) SetHttpStatusCode(v int32) *ListRecentCallDetailRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBody) SetMessage(v string) *ListRecentCallDetailRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBody) SetRequestId(v string) *ListRecentCallDetailRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBody) SetData(v *ListRecentCallDetailRecordsResponseBodyData) *ListRecentCallDetailRecordsResponseBody {
	s.Data = v
	return s
}

type ListRecentCallDetailRecordsResponseBodyData struct {
	PageNumber *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListRecentCallDetailRecordsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListRecentCallDetailRecordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRecentCallDetailRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRecentCallDetailRecordsResponseBodyData) SetPageNumber(v int32) *ListRecentCallDetailRecordsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyData) SetPageSize(v int32) *ListRecentCallDetailRecordsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyData) SetTotalCount(v int32) *ListRecentCallDetailRecordsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyData) SetList(v []*ListRecentCallDetailRecordsResponseBodyDataList) *ListRecentCallDetailRecordsResponseBodyData {
	s.List = v
	return s
}

type ListRecentCallDetailRecordsResponseBodyDataList struct {
	StartTime          *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ContactDisposition *string `json:"ContactDisposition,omitempty" xml:"ContactDisposition,omitempty"`
	ContactType        *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	AgentIds           *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	ContactId          *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	CallDuration       *string `json:"CallDuration,omitempty" xml:"CallDuration,omitempty"`
	CallingNumber      *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	Duration           *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CalledNumber       *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	SkillGroupIds      *string `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty"`
}

func (s ListRecentCallDetailRecordsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListRecentCallDetailRecordsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetStartTime(v int64) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetContactDisposition(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.ContactDisposition = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetContactType(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.ContactType = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetAgentIds(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.AgentIds = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetContactId(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.ContactId = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetCallDuration(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.CallDuration = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetCallingNumber(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.CallingNumber = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetDuration(v int64) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.Duration = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetInstanceId(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetCalledNumber(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.CalledNumber = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponseBodyDataList) SetSkillGroupIds(v string) *ListRecentCallDetailRecordsResponseBodyDataList {
	s.SkillGroupIds = &v
	return s
}

type ListRecentCallDetailRecordsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRecentCallDetailRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRecentCallDetailRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecentCallDetailRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListRecentCallDetailRecordsResponse) SetHeaders(v map[string]*string) *ListRecentCallDetailRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListRecentCallDetailRecordsResponse) SetBody(v *ListRecentCallDetailRecordsResponseBody) *ListRecentCallDetailRecordsResponse {
	s.Body = v
	return s
}

type InitiateAttendedTransferRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId       *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Transferor     *string `json:"Transferor,omitempty" xml:"Transferor,omitempty"`
	Transferee     *string `json:"Transferee,omitempty" xml:"Transferee,omitempty"`
	TimeoutSeconds *int32  `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s InitiateAttendedTransferRequest) String() string {
	return tea.Prettify(s)
}

func (s InitiateAttendedTransferRequest) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferRequest) SetInstanceId(v string) *InitiateAttendedTransferRequest {
	s.InstanceId = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetUserId(v string) *InitiateAttendedTransferRequest {
	s.UserId = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetDeviceId(v string) *InitiateAttendedTransferRequest {
	s.DeviceId = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTransferor(v string) *InitiateAttendedTransferRequest {
	s.Transferor = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTransferee(v string) *InitiateAttendedTransferRequest {
	s.Transferee = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetTimeoutSeconds(v int32) *InitiateAttendedTransferRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *InitiateAttendedTransferRequest) SetJobId(v string) *InitiateAttendedTransferRequest {
	s.JobId = &v
	return s
}

type InitiateAttendedTransferResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                                 `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *InitiateAttendedTransferResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s InitiateAttendedTransferResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitiateAttendedTransferResponseBody) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBody) SetCode(v string) *InitiateAttendedTransferResponseBody {
	s.Code = &v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetHttpStatusCode(v int32) *InitiateAttendedTransferResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetMessage(v string) *InitiateAttendedTransferResponseBody {
	s.Message = &v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetRequestId(v string) *InitiateAttendedTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetParams(v []*string) *InitiateAttendedTransferResponseBody {
	s.Params = v
	return s
}

func (s *InitiateAttendedTransferResponseBody) SetData(v *InitiateAttendedTransferResponseBodyData) *InitiateAttendedTransferResponseBody {
	s.Data = v
	return s
}

type InitiateAttendedTransferResponseBodyData struct {
	ContextId   *int64                                               `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	CallContext *InitiateAttendedTransferResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *InitiateAttendedTransferResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s InitiateAttendedTransferResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InitiateAttendedTransferResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBodyData) SetContextId(v int64) *InitiateAttendedTransferResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyData) SetCallContext(v *InitiateAttendedTransferResponseBodyDataCallContext) *InitiateAttendedTransferResponseBodyData {
	s.CallContext = v
	return s
}

func (s *InitiateAttendedTransferResponseBodyData) SetUserContext(v *InitiateAttendedTransferResponseBodyDataUserContext) *InitiateAttendedTransferResponseBodyData {
	s.UserContext = v
	return s
}

type InitiateAttendedTransferResponseBodyDataCallContext struct {
	CallType        *string                                                               `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*InitiateAttendedTransferResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s InitiateAttendedTransferResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s InitiateAttendedTransferResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) SetCallType(v string) *InitiateAttendedTransferResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) SetInstanceId(v string) *InitiateAttendedTransferResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) SetJobId(v string) *InitiateAttendedTransferResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContext) SetChannelContexts(v []*InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) *InitiateAttendedTransferResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type InitiateAttendedTransferResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetDestination(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserId(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetCallType(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetJobId(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *InitiateAttendedTransferResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type InitiateAttendedTransferResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s InitiateAttendedTransferResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s InitiateAttendedTransferResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetExtension(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetHeartbeat(v int64) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetWorkMode(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetDeviceId(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetUserId(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetReserved(v int64) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetBreakCode(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetInstanceId(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetOutboundScenario(v bool) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetMobile(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetJobId(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetUserState(v string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *InitiateAttendedTransferResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *InitiateAttendedTransferResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type InitiateAttendedTransferResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitiateAttendedTransferResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitiateAttendedTransferResponse) String() string {
	return tea.Prettify(s)
}

func (s InitiateAttendedTransferResponse) GoString() string {
	return s.String()
}

func (s *InitiateAttendedTransferResponse) SetHeaders(v map[string]*string) *InitiateAttendedTransferResponse {
	s.Headers = v
	return s
}

func (s *InitiateAttendedTransferResponse) SetBody(v *InitiateAttendedTransferResponseBody) *InitiateAttendedTransferResponse {
	s.Body = v
	return s
}

type MakeCallRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId       *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Caller         *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Callee         *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	TimeoutSeconds *int32  `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	Tags           *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s MakeCallRequest) String() string {
	return tea.Prettify(s)
}

func (s MakeCallRequest) GoString() string {
	return s.String()
}

func (s *MakeCallRequest) SetInstanceId(v string) *MakeCallRequest {
	s.InstanceId = &v
	return s
}

func (s *MakeCallRequest) SetUserId(v string) *MakeCallRequest {
	s.UserId = &v
	return s
}

func (s *MakeCallRequest) SetDeviceId(v string) *MakeCallRequest {
	s.DeviceId = &v
	return s
}

func (s *MakeCallRequest) SetCaller(v string) *MakeCallRequest {
	s.Caller = &v
	return s
}

func (s *MakeCallRequest) SetCallee(v string) *MakeCallRequest {
	s.Callee = &v
	return s
}

func (s *MakeCallRequest) SetTimeoutSeconds(v int32) *MakeCallRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *MakeCallRequest) SetTags(v string) *MakeCallRequest {
	s.Tags = &v
	return s
}

type MakeCallResponseBody struct {
	Code           *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                 `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *MakeCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s MakeCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MakeCallResponseBody) GoString() string {
	return s.String()
}

func (s *MakeCallResponseBody) SetCode(v string) *MakeCallResponseBody {
	s.Code = &v
	return s
}

func (s *MakeCallResponseBody) SetHttpStatusCode(v int32) *MakeCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MakeCallResponseBody) SetMessage(v string) *MakeCallResponseBody {
	s.Message = &v
	return s
}

func (s *MakeCallResponseBody) SetRequestId(v string) *MakeCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *MakeCallResponseBody) SetParams(v []*string) *MakeCallResponseBody {
	s.Params = v
	return s
}

func (s *MakeCallResponseBody) SetData(v *MakeCallResponseBodyData) *MakeCallResponseBody {
	s.Data = v
	return s
}

type MakeCallResponseBodyData struct {
	ContextId   *int64                               `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	CallContext *MakeCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *MakeCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s MakeCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s MakeCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *MakeCallResponseBodyData) SetContextId(v int64) *MakeCallResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *MakeCallResponseBodyData) SetCallContext(v *MakeCallResponseBodyDataCallContext) *MakeCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *MakeCallResponseBodyData) SetUserContext(v *MakeCallResponseBodyDataUserContext) *MakeCallResponseBodyData {
	s.UserContext = v
	return s
}

type MakeCallResponseBodyDataCallContext struct {
	CallType        *string                                               `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*MakeCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s MakeCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s MakeCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *MakeCallResponseBodyDataCallContext) SetCallType(v string) *MakeCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContext) SetInstanceId(v string) *MakeCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContext) SetJobId(v string) *MakeCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContext) SetChannelContexts(v []*MakeCallResponseBodyDataCallContextChannelContexts) *MakeCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type MakeCallResponseBodyDataCallContextChannelContexts struct {
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
}

func (s MakeCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s MakeCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *MakeCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *MakeCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

type MakeCallResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s MakeCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s MakeCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *MakeCallResponseBodyDataUserContext) SetExtension(v string) *MakeCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetWorkMode(v string) *MakeCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetDeviceId(v string) *MakeCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetJobId(v string) *MakeCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetUserId(v string) *MakeCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetBreakCode(v string) *MakeCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetInstanceId(v string) *MakeCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *MakeCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetUserState(v string) *MakeCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *MakeCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *MakeCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type MakeCallResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MakeCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MakeCallResponse) String() string {
	return tea.Prettify(s)
}

func (s MakeCallResponse) GoString() string {
	return s.String()
}

func (s *MakeCallResponse) SetHeaders(v map[string]*string) *MakeCallResponse {
	s.Headers = v
	return s
}

func (s *MakeCallResponse) SetBody(v *MakeCallResponseBody) *MakeCallResponse {
	s.Body = v
	return s
}

type GetInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) SetInstanceId(v string) *GetInstanceRequest {
	s.InstanceId = &v
	return s
}

type GetInstanceResponseBody struct {
	Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                    `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *GetInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetHttpStatusCode(v int32) *GetInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetParams(v []*string) *GetInstanceResponseBody {
	s.Params = v
	return s
}

func (s *GetInstanceResponseBody) SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody {
	s.Data = v
	return s
}

type GetInstanceResponseBodyData struct {
	Status      *string                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	ConsoleUrl  *string                                  `json:"ConsoleUrl,omitempty" xml:"ConsoleUrl,omitempty"`
	Description *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	AliyunUid   *string                                  `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	Name        *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	DomainName  *string                                  `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Id          *string                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	AdminList   []*GetInstanceResponseBodyDataAdminList  `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	NumberList  []*GetInstanceResponseBodyDataNumberList `json:"NumberList,omitempty" xml:"NumberList,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyData) SetStatus(v string) *GetInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetConsoleUrl(v string) *GetInstanceResponseBodyData {
	s.ConsoleUrl = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetDescription(v string) *GetInstanceResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetAliyunUid(v string) *GetInstanceResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetName(v string) *GetInstanceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetDomainName(v string) *GetInstanceResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetId(v string) *GetInstanceResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetAdminList(v []*GetInstanceResponseBodyDataAdminList) *GetInstanceResponseBodyData {
	s.AdminList = v
	return s
}

func (s *GetInstanceResponseBodyData) SetNumberList(v []*GetInstanceResponseBodyDataNumberList) *GetInstanceResponseBodyData {
	s.NumberList = v
	return s
}

type GetInstanceResponseBodyDataAdminList struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Extension   *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	LoginName   *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	WorkMode    *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RoleName    *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleId      *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s GetInstanceResponseBodyDataAdminList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataAdminList) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataAdminList) SetDisplayName(v string) *GetInstanceResponseBodyDataAdminList {
	s.DisplayName = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetExtension(v string) *GetInstanceResponseBodyDataAdminList {
	s.Extension = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetLoginName(v string) *GetInstanceResponseBodyDataAdminList {
	s.LoginName = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetEmail(v string) *GetInstanceResponseBodyDataAdminList {
	s.Email = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetWorkMode(v string) *GetInstanceResponseBodyDataAdminList {
	s.WorkMode = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetMobile(v string) *GetInstanceResponseBodyDataAdminList {
	s.Mobile = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetUserId(v string) *GetInstanceResponseBodyDataAdminList {
	s.UserId = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetRoleName(v string) *GetInstanceResponseBodyDataAdminList {
	s.RoleName = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetInstanceId(v string) *GetInstanceResponseBodyDataAdminList {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyDataAdminList) SetRoleId(v string) *GetInstanceResponseBodyDataAdminList {
	s.RoleId = &v
	return s
}

type GetInstanceResponseBodyDataNumberList struct {
	Active        *bool                                               `json:"Active,omitempty" xml:"Active,omitempty"`
	UserId        *string                                             `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Number        *string                                             `json:"Number,omitempty" xml:"Number,omitempty"`
	City          *string                                             `json:"City,omitempty" xml:"City,omitempty"`
	InstanceId    *string                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Usage         *string                                             `json:"Usage,omitempty" xml:"Usage,omitempty"`
	ContactFlowId *string                                             `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	Province      *string                                             `json:"Province,omitempty" xml:"Province,omitempty"`
	SkillGroups   []*GetInstanceResponseBodyDataNumberListSkillGroups `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyDataNumberList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataNumberList) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNumberList) SetActive(v bool) *GetInstanceResponseBodyDataNumberList {
	s.Active = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetUserId(v string) *GetInstanceResponseBodyDataNumberList {
	s.UserId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetNumber(v string) *GetInstanceResponseBodyDataNumberList {
	s.Number = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetCity(v string) *GetInstanceResponseBodyDataNumberList {
	s.City = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetInstanceId(v string) *GetInstanceResponseBodyDataNumberList {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetUsage(v string) *GetInstanceResponseBodyDataNumberList {
	s.Usage = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetContactFlowId(v string) *GetInstanceResponseBodyDataNumberList {
	s.ContactFlowId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetProvince(v string) *GetInstanceResponseBodyDataNumberList {
	s.Province = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberList) SetSkillGroups(v []*GetInstanceResponseBodyDataNumberListSkillGroups) *GetInstanceResponseBodyDataNumberList {
	s.SkillGroups = v
	return s
}

type GetInstanceResponseBodyDataNumberListSkillGroups struct {
	DisplayName      *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PhoneNumberCount *int32  `json:"PhoneNumberCount,omitempty" xml:"PhoneNumberCount,omitempty"`
	SkillGroupId     *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	UserCount        *int32  `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetInstanceResponseBodyDataNumberListSkillGroups) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataNumberListSkillGroups) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetDisplayName(v string) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.DisplayName = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetDescription(v string) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.Description = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetPhoneNumberCount(v int32) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.PhoneNumberCount = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetSkillGroupId(v string) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.SkillGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetUserCount(v int32) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.UserCount = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetInstanceId(v string) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNumberListSkillGroups) SetName(v string) *GetInstanceResponseBodyDataNumberListSkillGroups {
	s.Name = &v
	return s
}

type GetInstanceResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type AddUsersToSkillGroupRequest struct {
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId       *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	UserSkillLevelList *string `json:"UserSkillLevelList,omitempty" xml:"UserSkillLevelList,omitempty"`
}

func (s AddUsersToSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUsersToSkillGroupRequest) SetInstanceId(v string) *AddUsersToSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AddUsersToSkillGroupRequest) SetSkillGroupId(v string) *AddUsersToSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *AddUsersToSkillGroupRequest) SetUserSkillLevelList(v string) *AddUsersToSkillGroupRequest {
	s.UserSkillLevelList = &v
	return s
}

type AddUsersToSkillGroupResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUsersToSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddUsersToSkillGroupResponseBody) SetCode(v string) *AddUsersToSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddUsersToSkillGroupResponseBody) SetHttpStatusCode(v int32) *AddUsersToSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddUsersToSkillGroupResponseBody) SetMessage(v string) *AddUsersToSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AddUsersToSkillGroupResponseBody) SetRequestId(v string) *AddUsersToSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddUsersToSkillGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUsersToSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUsersToSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUsersToSkillGroupResponse) SetHeaders(v map[string]*string) *AddUsersToSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUsersToSkillGroupResponse) SetBody(v *AddUsersToSkillGroupResponseBody) *AddUsersToSkillGroupResponse {
	s.Body = v
	return s
}

type ListConfigItemsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ObjectId   *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s ListConfigItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConfigItemsRequest) GoString() string {
	return s.String()
}

func (s *ListConfigItemsRequest) SetInstanceId(v string) *ListConfigItemsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConfigItemsRequest) SetObjectId(v string) *ListConfigItemsRequest {
	s.ObjectId = &v
	return s
}

func (s *ListConfigItemsRequest) SetObjectType(v string) *ListConfigItemsRequest {
	s.ObjectType = &v
	return s
}

type ListConfigItemsResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                          `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           []*ListConfigItemsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListConfigItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConfigItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigItemsResponseBody) SetCode(v string) *ListConfigItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConfigItemsResponseBody) SetHttpStatusCode(v int32) *ListConfigItemsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListConfigItemsResponseBody) SetMessage(v string) *ListConfigItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConfigItemsResponseBody) SetRequestId(v string) *ListConfigItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigItemsResponseBody) SetParams(v []*string) *ListConfigItemsResponseBody {
	s.Params = v
	return s
}

func (s *ListConfigItemsResponseBody) SetData(v []*ListConfigItemsResponseBodyData) *ListConfigItemsResponseBody {
	s.Data = v
	return s
}

type ListConfigItemsResponseBodyData struct {
	ObjectId   *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListConfigItemsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListConfigItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConfigItemsResponseBodyData) SetObjectId(v string) *ListConfigItemsResponseBodyData {
	s.ObjectId = &v
	return s
}

func (s *ListConfigItemsResponseBodyData) SetValue(v string) *ListConfigItemsResponseBodyData {
	s.Value = &v
	return s
}

func (s *ListConfigItemsResponseBodyData) SetObjectType(v string) *ListConfigItemsResponseBodyData {
	s.ObjectType = &v
	return s
}

func (s *ListConfigItemsResponseBodyData) SetInstanceId(v string) *ListConfigItemsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListConfigItemsResponseBodyData) SetName(v string) *ListConfigItemsResponseBodyData {
	s.Name = &v
	return s
}

type ListConfigItemsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConfigItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConfigItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigItemsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigItemsResponse) SetHeaders(v map[string]*string) *ListConfigItemsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigItemsResponse) SetBody(v *ListConfigItemsResponseBody) *ListConfigItemsResponse {
	s.Body = v
	return s
}

type SignInGroupRequest struct {
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId                 *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId               *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	SignedSkillGroupIdList *string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty"`
}

func (s SignInGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SignInGroupRequest) GoString() string {
	return s.String()
}

func (s *SignInGroupRequest) SetInstanceId(v string) *SignInGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *SignInGroupRequest) SetUserId(v string) *SignInGroupRequest {
	s.UserId = &v
	return s
}

func (s *SignInGroupRequest) SetDeviceId(v string) *SignInGroupRequest {
	s.DeviceId = &v
	return s
}

func (s *SignInGroupRequest) SetSignedSkillGroupIdList(v string) *SignInGroupRequest {
	s.SignedSkillGroupIdList = &v
	return s
}

type SignInGroupResponseBody struct {
	Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                    `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *SignInGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s SignInGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SignInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SignInGroupResponseBody) SetCode(v string) *SignInGroupResponseBody {
	s.Code = &v
	return s
}

func (s *SignInGroupResponseBody) SetHttpStatusCode(v int32) *SignInGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SignInGroupResponseBody) SetMessage(v string) *SignInGroupResponseBody {
	s.Message = &v
	return s
}

func (s *SignInGroupResponseBody) SetRequestId(v string) *SignInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SignInGroupResponseBody) SetParams(v []*string) *SignInGroupResponseBody {
	s.Params = v
	return s
}

func (s *SignInGroupResponseBody) SetData(v *SignInGroupResponseBodyData) *SignInGroupResponseBody {
	s.Data = v
	return s
}

type SignInGroupResponseBodyData struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s SignInGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SignInGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *SignInGroupResponseBodyData) SetExtension(v string) *SignInGroupResponseBodyData {
	s.Extension = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetWorkMode(v string) *SignInGroupResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetDeviceId(v string) *SignInGroupResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetJobId(v string) *SignInGroupResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetUserId(v string) *SignInGroupResponseBodyData {
	s.UserId = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetBreakCode(v string) *SignInGroupResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetInstanceId(v string) *SignInGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetOutboundScenario(v bool) *SignInGroupResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetUserState(v string) *SignInGroupResponseBodyData {
	s.UserState = &v
	return s
}

func (s *SignInGroupResponseBodyData) SetSignedSkillGroupIdList(v []*string) *SignInGroupResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

type SignInGroupResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SignInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SignInGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SignInGroupResponse) GoString() string {
	return s.String()
}

func (s *SignInGroupResponse) SetHeaders(v map[string]*string) *SignInGroupResponse {
	s.Headers = v
	return s
}

func (s *SignInGroupResponse) SetBody(v *SignInGroupResponseBody) *SignInGroupResponse {
	s.Body = v
	return s
}

type GetRealtimeInstanceStatesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRealtimeInstanceStatesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeInstanceStatesRequest) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatesRequest) SetInstanceId(v string) *GetRealtimeInstanceStatesRequest {
	s.InstanceId = &v
	return s
}

type GetRealtimeInstanceStatesResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetRealtimeInstanceStatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetRealtimeInstanceStatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeInstanceStatesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatesResponseBody) SetCode(v string) *GetRealtimeInstanceStatesResponseBody {
	s.Code = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBody) SetHttpStatusCode(v int32) *GetRealtimeInstanceStatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBody) SetMessage(v string) *GetRealtimeInstanceStatesResponseBody {
	s.Message = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBody) SetRequestId(v string) *GetRealtimeInstanceStatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBody) SetData(v *GetRealtimeInstanceStatesResponseBodyData) *GetRealtimeInstanceStatesResponseBody {
	s.Data = v
	return s
}

type GetRealtimeInstanceStatesResponseBodyData struct {
	WorkingAgents      *int64  `json:"WorkingAgents,omitempty" xml:"WorkingAgents,omitempty"`
	LongestWaitingTime *int64  `json:"LongestWaitingTime,omitempty" xml:"LongestWaitingTime,omitempty"`
	LoggedInAgents     *int64  `json:"LoggedInAgents,omitempty" xml:"LoggedInAgents,omitempty"`
	TotalAgents        *int64  `json:"TotalAgents,omitempty" xml:"TotalAgents,omitempty"`
	WaitingCalls       *int64  `json:"WaitingCalls,omitempty" xml:"WaitingCalls,omitempty"`
	BreakingAgents     *int64  `json:"BreakingAgents,omitempty" xml:"BreakingAgents,omitempty"`
	TalkingAgents      *int64  `json:"TalkingAgents,omitempty" xml:"TalkingAgents,omitempty"`
	InteractiveCalls   *int64  `json:"InteractiveCalls,omitempty" xml:"InteractiveCalls,omitempty"`
	ReadyAgents        *int64  `json:"ReadyAgents,omitempty" xml:"ReadyAgents,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRealtimeInstanceStatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeInstanceStatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetWorkingAgents(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.WorkingAgents = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetLongestWaitingTime(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.LongestWaitingTime = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetLoggedInAgents(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.LoggedInAgents = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetTotalAgents(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.TotalAgents = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetWaitingCalls(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.WaitingCalls = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetBreakingAgents(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.BreakingAgents = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetTalkingAgents(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.TalkingAgents = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetInteractiveCalls(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.InteractiveCalls = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetReadyAgents(v int64) *GetRealtimeInstanceStatesResponseBodyData {
	s.ReadyAgents = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponseBodyData) SetInstanceId(v string) *GetRealtimeInstanceStatesResponseBodyData {
	s.InstanceId = &v
	return s
}

type GetRealtimeInstanceStatesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRealtimeInstanceStatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRealtimeInstanceStatesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeInstanceStatesResponse) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatesResponse) SetHeaders(v map[string]*string) *GetRealtimeInstanceStatesResponse {
	s.Headers = v
	return s
}

func (s *GetRealtimeInstanceStatesResponse) SetBody(v *GetRealtimeInstanceStatesResponseBody) *GetRealtimeInstanceStatesResponse {
	s.Body = v
	return s
}

type ModifySkillGroupRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifySkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySkillGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifySkillGroupRequest) SetInstanceId(v string) *ModifySkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySkillGroupRequest) SetSkillGroupId(v string) *ModifySkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ModifySkillGroupRequest) SetDisplayName(v string) *ModifySkillGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *ModifySkillGroupRequest) SetDescription(v string) *ModifySkillGroupRequest {
	s.Description = &v
	return s
}

type ModifySkillGroupResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySkillGroupResponseBody) SetCode(v string) *ModifySkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySkillGroupResponseBody) SetHttpStatusCode(v int32) *ModifySkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifySkillGroupResponseBody) SetMessage(v string) *ModifySkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySkillGroupResponseBody) SetRequestId(v string) *ModifySkillGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifySkillGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySkillGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifySkillGroupResponse) SetHeaders(v map[string]*string) *ModifySkillGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifySkillGroupResponse) SetBody(v *ModifySkillGroupResponseBody) *ModifySkillGroupResponse {
	s.Body = v
	return s
}

type ModifyUserRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Mobile     *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	WorkMode   *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	RoleId     *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s ModifyUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserRequest) SetInstanceId(v string) *ModifyUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserRequest) SetUserId(v string) *ModifyUserRequest {
	s.UserId = &v
	return s
}

func (s *ModifyUserRequest) SetMobile(v string) *ModifyUserRequest {
	s.Mobile = &v
	return s
}

func (s *ModifyUserRequest) SetWorkMode(v string) *ModifyUserRequest {
	s.WorkMode = &v
	return s
}

func (s *ModifyUserRequest) SetRoleId(v string) *ModifyUserRequest {
	s.RoleId = &v
	return s
}

type ModifyUserResponseBody struct {
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Data           *string   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s ModifyUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserResponseBody) SetHttpStatusCode(v int32) *ModifyUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyUserResponseBody) SetCode(v string) *ModifyUserResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyUserResponseBody) SetMessage(v string) *ModifyUserResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyUserResponseBody) SetData(v string) *ModifyUserResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyUserResponseBody) SetRequestId(v string) *ModifyUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserResponseBody) SetParams(v []*string) *ModifyUserResponseBody {
	s.Params = v
	return s
}

type ModifyUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserResponse) SetHeaders(v map[string]*string) *ModifyUserResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserResponse) SetBody(v *ModifyUserResponseBody) *ModifyUserResponse {
	s.Body = v
	return s
}

type AddPhoneNumberToSkillGroupsRequest struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Number           *string `json:"Number,omitempty" xml:"Number,omitempty"`
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
}

func (s AddPhoneNumberToSkillGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPhoneNumberToSkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *AddPhoneNumberToSkillGroupsRequest) SetInstanceId(v string) *AddPhoneNumberToSkillGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsRequest) SetNumber(v string) *AddPhoneNumberToSkillGroupsRequest {
	s.Number = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsRequest) SetSkillGroupIdList(v string) *AddPhoneNumberToSkillGroupsRequest {
	s.SkillGroupIdList = &v
	return s
}

type AddPhoneNumberToSkillGroupsResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPhoneNumberToSkillGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddPhoneNumberToSkillGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) SetCode(v string) *AddPhoneNumberToSkillGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) SetHttpStatusCode(v int32) *AddPhoneNumberToSkillGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) SetMessage(v string) *AddPhoneNumberToSkillGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *AddPhoneNumberToSkillGroupsResponseBody) SetRequestId(v string) *AddPhoneNumberToSkillGroupsResponseBody {
	s.RequestId = &v
	return s
}

type AddPhoneNumberToSkillGroupsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddPhoneNumberToSkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddPhoneNumberToSkillGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPhoneNumberToSkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *AddPhoneNumberToSkillGroupsResponse) SetHeaders(v map[string]*string) *AddPhoneNumberToSkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *AddPhoneNumberToSkillGroupsResponse) SetBody(v *AddPhoneNumberToSkillGroupsResponseBody) *AddPhoneNumberToSkillGroupsResponse {
	s.Body = v
	return s
}

type ListDevicesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListDevicesRequest) SetInstanceId(v string) *ListDevicesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDevicesRequest) SetUserId(v string) *ListDevicesRequest {
	s.UserId = &v
	return s
}

type ListDevicesResponseBody struct {
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                      `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           []*ListDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBody) SetCode(v string) *ListDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDevicesResponseBody) SetHttpStatusCode(v int32) *ListDevicesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDevicesResponseBody) SetMessage(v string) *ListDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDevicesResponseBody) SetRequestId(v string) *ListDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevicesResponseBody) SetParams(v []*string) *ListDevicesResponseBody {
	s.Params = v
	return s
}

func (s *ListDevicesResponseBody) SetData(v []*ListDevicesResponseBodyData) *ListDevicesResponseBody {
	s.Data = v
	return s
}

type ListDevicesResponseBodyData struct {
	Extension  *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Expires    *int64  `json:"Expires,omitempty" xml:"Expires,omitempty"`
	Contact    *string `json:"Contact,omitempty" xml:"Contact,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	CallId     *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListDevicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDevicesResponseBodyData) SetExtension(v string) *ListDevicesResponseBodyData {
	s.Extension = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetExpires(v int64) *ListDevicesResponseBodyData {
	s.Expires = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetContact(v string) *ListDevicesResponseBodyData {
	s.Contact = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetDeviceId(v string) *ListDevicesResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetUserId(v string) *ListDevicesResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetCallId(v string) *ListDevicesResponseBodyData {
	s.CallId = &v
	return s
}

func (s *ListDevicesResponseBodyData) SetInstanceId(v string) *ListDevicesResponseBodyData {
	s.InstanceId = &v
	return s
}

type ListDevicesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListDevicesResponse) SetHeaders(v map[string]*string) *ListDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListDevicesResponse) SetBody(v *ListDevicesResponseBody) *ListDevicesResponse {
	s.Body = v
	return s
}

type RetrieveCallRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s RetrieveCallRequest) String() string {
	return tea.Prettify(s)
}

func (s RetrieveCallRequest) GoString() string {
	return s.String()
}

func (s *RetrieveCallRequest) SetInstanceId(v string) *RetrieveCallRequest {
	s.InstanceId = &v
	return s
}

func (s *RetrieveCallRequest) SetUserId(v string) *RetrieveCallRequest {
	s.UserId = &v
	return s
}

func (s *RetrieveCallRequest) SetDeviceId(v string) *RetrieveCallRequest {
	s.DeviceId = &v
	return s
}

func (s *RetrieveCallRequest) SetJobId(v string) *RetrieveCallRequest {
	s.JobId = &v
	return s
}

func (s *RetrieveCallRequest) SetChannelId(v string) *RetrieveCallRequest {
	s.ChannelId = &v
	return s
}

type RetrieveCallResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                     `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *RetrieveCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s RetrieveCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetrieveCallResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveCallResponseBody) SetCode(v string) *RetrieveCallResponseBody {
	s.Code = &v
	return s
}

func (s *RetrieveCallResponseBody) SetHttpStatusCode(v int32) *RetrieveCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RetrieveCallResponseBody) SetMessage(v string) *RetrieveCallResponseBody {
	s.Message = &v
	return s
}

func (s *RetrieveCallResponseBody) SetRequestId(v string) *RetrieveCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrieveCallResponseBody) SetParams(v []*string) *RetrieveCallResponseBody {
	s.Params = v
	return s
}

func (s *RetrieveCallResponseBody) SetData(v *RetrieveCallResponseBodyData) *RetrieveCallResponseBody {
	s.Data = v
	return s
}

type RetrieveCallResponseBodyData struct {
	CallContext *RetrieveCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *RetrieveCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s RetrieveCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RetrieveCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetrieveCallResponseBodyData) SetCallContext(v *RetrieveCallResponseBodyDataCallContext) *RetrieveCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *RetrieveCallResponseBodyData) SetUserContext(v *RetrieveCallResponseBodyDataUserContext) *RetrieveCallResponseBodyData {
	s.UserContext = v
	return s
}

type RetrieveCallResponseBodyDataCallContext struct {
	CallType        *string                                                   `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*RetrieveCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s RetrieveCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s RetrieveCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *RetrieveCallResponseBodyDataCallContext) SetCallType(v string) *RetrieveCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContext) SetInstanceId(v string) *RetrieveCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContext) SetJobId(v string) *RetrieveCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContext) SetChannelContexts(v []*RetrieveCallResponseBodyDataCallContextChannelContexts) *RetrieveCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type RetrieveCallResponseBodyDataCallContextChannelContexts struct {
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
}

func (s RetrieveCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s RetrieveCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *RetrieveCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *RetrieveCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

type RetrieveCallResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s RetrieveCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s RetrieveCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *RetrieveCallResponseBodyDataUserContext) SetExtension(v string) *RetrieveCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetHeartbeat(v int64) *RetrieveCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetWorkMode(v string) *RetrieveCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetDeviceId(v string) *RetrieveCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetUserId(v string) *RetrieveCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetReserved(v int64) *RetrieveCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetBreakCode(v string) *RetrieveCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetInstanceId(v string) *RetrieveCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *RetrieveCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetMobile(v string) *RetrieveCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetJobId(v string) *RetrieveCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetUserState(v string) *RetrieveCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *RetrieveCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *RetrieveCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type RetrieveCallResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RetrieveCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RetrieveCallResponse) String() string {
	return tea.Prettify(s)
}

func (s RetrieveCallResponse) GoString() string {
	return s.String()
}

func (s *RetrieveCallResponse) SetHeaders(v map[string]*string) *RetrieveCallResponse {
	s.Headers = v
	return s
}

func (s *RetrieveCallResponse) SetBody(v *RetrieveCallResponseBody) *RetrieveCallResponse {
	s.Body = v
	return s
}

type ListSkillGroupsRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSkillGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListSkillGroupsRequest) SetInstanceId(v string) *ListSkillGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSkillGroupsRequest) SetSearchPattern(v string) *ListSkillGroupsRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListSkillGroupsRequest) SetPageNumber(v int32) *ListSkillGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSkillGroupsRequest) SetPageSize(v int32) *ListSkillGroupsRequest {
	s.PageSize = &v
	return s
}

type ListSkillGroupsResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListSkillGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListSkillGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillGroupsResponseBody) SetCode(v string) *ListSkillGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillGroupsResponseBody) SetHttpStatusCode(v int32) *ListSkillGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSkillGroupsResponseBody) SetMessage(v string) *ListSkillGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillGroupsResponseBody) SetRequestId(v string) *ListSkillGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillGroupsResponseBody) SetData(v *ListSkillGroupsResponseBodyData) *ListSkillGroupsResponseBody {
	s.Data = v
	return s
}

type ListSkillGroupsResponseBodyData struct {
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListSkillGroupsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListSkillGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillGroupsResponseBodyData) SetPageNumber(v int32) *ListSkillGroupsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSkillGroupsResponseBodyData) SetPageSize(v int32) *ListSkillGroupsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSkillGroupsResponseBodyData) SetTotalCount(v int32) *ListSkillGroupsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSkillGroupsResponseBodyData) SetList(v []*ListSkillGroupsResponseBodyDataList) *ListSkillGroupsResponseBodyData {
	s.List = v
	return s
}

type ListSkillGroupsResponseBodyDataList struct {
	DisplayName      *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PhoneNumberCount *int32  `json:"PhoneNumberCount,omitempty" xml:"PhoneNumberCount,omitempty"`
	SkillGroupId     *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName   *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	UserCount        *int32  `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListSkillGroupsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListSkillGroupsResponseBodyDataList) SetDisplayName(v string) *ListSkillGroupsResponseBodyDataList {
	s.DisplayName = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetDescription(v string) *ListSkillGroupsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetPhoneNumberCount(v int32) *ListSkillGroupsResponseBodyDataList {
	s.PhoneNumberCount = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetSkillGroupId(v string) *ListSkillGroupsResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetSkillGroupName(v string) *ListSkillGroupsResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetUserCount(v int32) *ListSkillGroupsResponseBodyDataList {
	s.UserCount = &v
	return s
}

func (s *ListSkillGroupsResponseBodyDataList) SetInstanceId(v string) *ListSkillGroupsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

type ListSkillGroupsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSkillGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListSkillGroupsResponse) SetHeaders(v map[string]*string) *ListSkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListSkillGroupsResponse) SetBody(v *ListSkillGroupsResponseBody) *ListSkillGroupsResponse {
	s.Body = v
	return s
}

type HoldCallRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Music      *string `json:"Music,omitempty" xml:"Music,omitempty"`
}

func (s HoldCallRequest) String() string {
	return tea.Prettify(s)
}

func (s HoldCallRequest) GoString() string {
	return s.String()
}

func (s *HoldCallRequest) SetInstanceId(v string) *HoldCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HoldCallRequest) SetUserId(v string) *HoldCallRequest {
	s.UserId = &v
	return s
}

func (s *HoldCallRequest) SetDeviceId(v string) *HoldCallRequest {
	s.DeviceId = &v
	return s
}

func (s *HoldCallRequest) SetJobId(v string) *HoldCallRequest {
	s.JobId = &v
	return s
}

func (s *HoldCallRequest) SetChannelId(v string) *HoldCallRequest {
	s.ChannelId = &v
	return s
}

func (s *HoldCallRequest) SetMusic(v string) *HoldCallRequest {
	s.Music = &v
	return s
}

type HoldCallResponseBody struct {
	Code           *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                 `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *HoldCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s HoldCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HoldCallResponseBody) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBody) SetCode(v string) *HoldCallResponseBody {
	s.Code = &v
	return s
}

func (s *HoldCallResponseBody) SetHttpStatusCode(v int32) *HoldCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *HoldCallResponseBody) SetMessage(v string) *HoldCallResponseBody {
	s.Message = &v
	return s
}

func (s *HoldCallResponseBody) SetRequestId(v string) *HoldCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HoldCallResponseBody) SetParams(v []*string) *HoldCallResponseBody {
	s.Params = v
	return s
}

func (s *HoldCallResponseBody) SetData(v *HoldCallResponseBodyData) *HoldCallResponseBody {
	s.Data = v
	return s
}

type HoldCallResponseBodyData struct {
	CallContext *HoldCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *HoldCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s HoldCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s HoldCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBodyData) SetCallContext(v *HoldCallResponseBodyDataCallContext) *HoldCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *HoldCallResponseBodyData) SetUserContext(v *HoldCallResponseBodyDataUserContext) *HoldCallResponseBodyData {
	s.UserContext = v
	return s
}

type HoldCallResponseBodyDataCallContext struct {
	CallType        *string                                               `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*HoldCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s HoldCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s HoldCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBodyDataCallContext) SetCallType(v string) *HoldCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContext) SetInstanceId(v string) *HoldCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContext) SetJobId(v string) *HoldCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContext) SetChannelContexts(v []*HoldCallResponseBodyDataCallContextChannelContexts) *HoldCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type HoldCallResponseBodyDataCallContextChannelContexts struct {
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
}

func (s HoldCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s HoldCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *HoldCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *HoldCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

type HoldCallResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s HoldCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s HoldCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBodyDataUserContext) SetExtension(v string) *HoldCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetHeartbeat(v int64) *HoldCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetWorkMode(v string) *HoldCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetDeviceId(v string) *HoldCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetUserId(v string) *HoldCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetReserved(v int64) *HoldCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetBreakCode(v string) *HoldCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetInstanceId(v string) *HoldCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *HoldCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetMobile(v string) *HoldCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetJobId(v string) *HoldCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetUserState(v string) *HoldCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *HoldCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *HoldCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
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

type RegisterDeviceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Password   *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s RegisterDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *RegisterDeviceRequest) SetInstanceId(v string) *RegisterDeviceRequest {
	s.InstanceId = &v
	return s
}

func (s *RegisterDeviceRequest) SetUserId(v string) *RegisterDeviceRequest {
	s.UserId = &v
	return s
}

func (s *RegisterDeviceRequest) SetDeviceId(v string) *RegisterDeviceRequest {
	s.DeviceId = &v
	return s
}

func (s *RegisterDeviceRequest) SetPassword(v string) *RegisterDeviceRequest {
	s.Password = &v
	return s
}

type RegisterDeviceResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s RegisterDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponseBody) SetCode(v string) *RegisterDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetHttpStatusCode(v int32) *RegisterDeviceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetMessage(v string) *RegisterDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetRequestId(v string) *RegisterDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterDeviceResponseBody) SetParams(v []*string) *RegisterDeviceResponseBody {
	s.Params = v
	return s
}

type RegisterDeviceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *RegisterDeviceResponse) SetHeaders(v map[string]*string) *RegisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *RegisterDeviceResponse) SetBody(v *RegisterDeviceResponseBody) *RegisterDeviceResponse {
	s.Body = v
	return s
}

type RemovePersonalNumbersFromUserRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
}

func (s RemovePersonalNumbersFromUserRequest) String() string {
	return tea.Prettify(s)
}

func (s RemovePersonalNumbersFromUserRequest) GoString() string {
	return s.String()
}

func (s *RemovePersonalNumbersFromUserRequest) SetInstanceId(v string) *RemovePersonalNumbersFromUserRequest {
	s.InstanceId = &v
	return s
}

func (s *RemovePersonalNumbersFromUserRequest) SetUserId(v string) *RemovePersonalNumbersFromUserRequest {
	s.UserId = &v
	return s
}

func (s *RemovePersonalNumbersFromUserRequest) SetNumberList(v string) *RemovePersonalNumbersFromUserRequest {
	s.NumberList = &v
	return s
}

type RemovePersonalNumbersFromUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePersonalNumbersFromUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemovePersonalNumbersFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePersonalNumbersFromUserResponseBody) SetCode(v string) *RemovePersonalNumbersFromUserResponseBody {
	s.Code = &v
	return s
}

func (s *RemovePersonalNumbersFromUserResponseBody) SetHttpStatusCode(v int32) *RemovePersonalNumbersFromUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemovePersonalNumbersFromUserResponseBody) SetMessage(v string) *RemovePersonalNumbersFromUserResponseBody {
	s.Message = &v
	return s
}

func (s *RemovePersonalNumbersFromUserResponseBody) SetRequestId(v string) *RemovePersonalNumbersFromUserResponseBody {
	s.RequestId = &v
	return s
}

type RemovePersonalNumbersFromUserResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemovePersonalNumbersFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemovePersonalNumbersFromUserResponse) String() string {
	return tea.Prettify(s)
}

func (s RemovePersonalNumbersFromUserResponse) GoString() string {
	return s.String()
}

func (s *RemovePersonalNumbersFromUserResponse) SetHeaders(v map[string]*string) *RemovePersonalNumbersFromUserResponse {
	s.Headers = v
	return s
}

func (s *RemovePersonalNumbersFromUserResponse) SetBody(v *RemovePersonalNumbersFromUserResponseBody) *RemovePersonalNumbersFromUserResponse {
	s.Body = v
	return s
}

type ModifyInstanceRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRequest) SetInstanceId(v string) *ModifyInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceRequest) SetDescription(v string) *ModifyInstanceRequest {
	s.Description = &v
	return s
}

type ModifyInstanceResponseBody struct {
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Data           *string   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s ModifyInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceResponseBody) SetHttpStatusCode(v int32) *ModifyInstanceResponseBody {
	s.HttpStatusCode = &v
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

func (s *ModifyInstanceResponseBody) SetData(v string) *ModifyInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetRequestId(v string) *ModifyInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceResponseBody) SetParams(v []*string) *ModifyInstanceResponseBody {
	s.Params = v
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

type ListOutboundNumbersOfUserRequest struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
}

func (s ListOutboundNumbersOfUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundNumbersOfUserRequest) GoString() string {
	return s.String()
}

func (s *ListOutboundNumbersOfUserRequest) SetInstanceId(v string) *ListOutboundNumbersOfUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOutboundNumbersOfUserRequest) SetUserId(v string) *ListOutboundNumbersOfUserRequest {
	s.UserId = &v
	return s
}

func (s *ListOutboundNumbersOfUserRequest) SetPageNumber(v int32) *ListOutboundNumbersOfUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOutboundNumbersOfUserRequest) SetPageSize(v int32) *ListOutboundNumbersOfUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListOutboundNumbersOfUserRequest) SetSkillGroupIdList(v string) *ListOutboundNumbersOfUserRequest {
	s.SkillGroupIdList = &v
	return s
}

type ListOutboundNumbersOfUserResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListOutboundNumbersOfUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListOutboundNumbersOfUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundNumbersOfUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListOutboundNumbersOfUserResponseBody) SetCode(v string) *ListOutboundNumbersOfUserResponseBody {
	s.Code = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBody) SetHttpStatusCode(v int32) *ListOutboundNumbersOfUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBody) SetMessage(v string) *ListOutboundNumbersOfUserResponseBody {
	s.Message = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBody) SetRequestId(v string) *ListOutboundNumbersOfUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBody) SetData(v *ListOutboundNumbersOfUserResponseBodyData) *ListOutboundNumbersOfUserResponseBody {
	s.Data = v
	return s
}

type ListOutboundNumbersOfUserResponseBodyData struct {
	PageNumber *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListOutboundNumbersOfUserResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListOutboundNumbersOfUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundNumbersOfUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOutboundNumbersOfUserResponseBodyData) SetPageNumber(v int32) *ListOutboundNumbersOfUserResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyData) SetPageSize(v int32) *ListOutboundNumbersOfUserResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyData) SetTotalCount(v int32) *ListOutboundNumbersOfUserResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyData) SetList(v []*ListOutboundNumbersOfUserResponseBodyDataList) *ListOutboundNumbersOfUserResponseBodyData {
	s.List = v
	return s
}

type ListOutboundNumbersOfUserResponseBodyDataList struct {
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListOutboundNumbersOfUserResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundNumbersOfUserResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) SetNumber(v string) *ListOutboundNumbersOfUserResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) SetCity(v string) *ListOutboundNumbersOfUserResponseBodyDataList {
	s.City = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) SetProvince(v string) *ListOutboundNumbersOfUserResponseBodyDataList {
	s.Province = &v
	return s
}

type ListOutboundNumbersOfUserResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOutboundNumbersOfUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOutboundNumbersOfUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundNumbersOfUserResponse) GoString() string {
	return s.String()
}

func (s *ListOutboundNumbersOfUserResponse) SetHeaders(v map[string]*string) *ListOutboundNumbersOfUserResponse {
	s.Headers = v
	return s
}

func (s *ListOutboundNumbersOfUserResponse) SetBody(v *ListOutboundNumbersOfUserResponseBody) *ListOutboundNumbersOfUserResponse {
	s.Body = v
	return s
}

type PollUserStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s PollUserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s PollUserStatusRequest) GoString() string {
	return s.String()
}

func (s *PollUserStatusRequest) SetInstanceId(v string) *PollUserStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *PollUserStatusRequest) SetUserId(v string) *PollUserStatusRequest {
	s.UserId = &v
	return s
}

func (s *PollUserStatusRequest) SetDeviceId(v string) *PollUserStatusRequest {
	s.DeviceId = &v
	return s
}

type PollUserStatusResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                       `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *PollUserStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s PollUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PollUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBody) SetCode(v string) *PollUserStatusResponseBody {
	s.Code = &v
	return s
}

func (s *PollUserStatusResponseBody) SetHttpStatusCode(v int32) *PollUserStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PollUserStatusResponseBody) SetMessage(v string) *PollUserStatusResponseBody {
	s.Message = &v
	return s
}

func (s *PollUserStatusResponseBody) SetRequestId(v string) *PollUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *PollUserStatusResponseBody) SetParams(v []*string) *PollUserStatusResponseBody {
	s.Params = v
	return s
}

func (s *PollUserStatusResponseBody) SetData(v *PollUserStatusResponseBodyData) *PollUserStatusResponseBody {
	s.Data = v
	return s
}

type PollUserStatusResponseBodyData struct {
	ContextId   *int64                                     `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	CallContext *PollUserStatusResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *PollUserStatusResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s PollUserStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PollUserStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBodyData) SetContextId(v int64) *PollUserStatusResponseBodyData {
	s.ContextId = &v
	return s
}

func (s *PollUserStatusResponseBodyData) SetCallContext(v *PollUserStatusResponseBodyDataCallContext) *PollUserStatusResponseBodyData {
	s.CallContext = v
	return s
}

func (s *PollUserStatusResponseBodyData) SetUserContext(v *PollUserStatusResponseBodyDataUserContext) *PollUserStatusResponseBodyData {
	s.UserContext = v
	return s
}

type PollUserStatusResponseBodyDataCallContext struct {
	CallType        *string                                                     `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                     `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*PollUserStatusResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s PollUserStatusResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s PollUserStatusResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBodyDataCallContext) SetCallType(v string) *PollUserStatusResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContext) SetInstanceId(v string) *PollUserStatusResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContext) SetJobId(v string) *PollUserStatusResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContext) SetChannelContexts(v []*PollUserStatusResponseBodyDataCallContextChannelContexts) *PollUserStatusResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type PollUserStatusResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	ChannelVariables *string                `json:"ChannelVariables,omitempty" xml:"ChannelVariables,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s PollUserStatusResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s PollUserStatusResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetDestination(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetUserId(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetCallType(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetChannelVariables(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.ChannelVariables = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetJobId(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *PollUserStatusResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *PollUserStatusResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type PollUserStatusResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s PollUserStatusResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s PollUserStatusResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponseBodyDataUserContext) SetExtension(v string) *PollUserStatusResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetHeartbeat(v int64) *PollUserStatusResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetWorkMode(v string) *PollUserStatusResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetDeviceId(v string) *PollUserStatusResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetUserId(v string) *PollUserStatusResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetReserved(v int64) *PollUserStatusResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetBreakCode(v string) *PollUserStatusResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetInstanceId(v string) *PollUserStatusResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetOutboundScenario(v bool) *PollUserStatusResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetMobile(v string) *PollUserStatusResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetJobId(v string) *PollUserStatusResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetUserState(v string) *PollUserStatusResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *PollUserStatusResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *PollUserStatusResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type PollUserStatusResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PollUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PollUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s PollUserStatusResponse) GoString() string {
	return s.String()
}

func (s *PollUserStatusResponse) SetHeaders(v map[string]*string) *PollUserStatusResponse {
	s.Headers = v
	return s
}

func (s *PollUserStatusResponse) SetBody(v *PollUserStatusResponseBody) *PollUserStatusResponse {
	s.Body = v
	return s
}

type ReadyForServiceRequest struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId         *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OutboundScenario *bool   `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
}

func (s ReadyForServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReadyForServiceRequest) GoString() string {
	return s.String()
}

func (s *ReadyForServiceRequest) SetInstanceId(v string) *ReadyForServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReadyForServiceRequest) SetUserId(v string) *ReadyForServiceRequest {
	s.UserId = &v
	return s
}

func (s *ReadyForServiceRequest) SetDeviceId(v string) *ReadyForServiceRequest {
	s.DeviceId = &v
	return s
}

func (s *ReadyForServiceRequest) SetOutboundScenario(v bool) *ReadyForServiceRequest {
	s.OutboundScenario = &v
	return s
}

type ReadyForServiceResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                        `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *ReadyForServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ReadyForServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReadyForServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ReadyForServiceResponseBody) SetCode(v string) *ReadyForServiceResponseBody {
	s.Code = &v
	return s
}

func (s *ReadyForServiceResponseBody) SetHttpStatusCode(v int32) *ReadyForServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReadyForServiceResponseBody) SetMessage(v string) *ReadyForServiceResponseBody {
	s.Message = &v
	return s
}

func (s *ReadyForServiceResponseBody) SetRequestId(v string) *ReadyForServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadyForServiceResponseBody) SetParams(v []*string) *ReadyForServiceResponseBody {
	s.Params = v
	return s
}

func (s *ReadyForServiceResponseBody) SetData(v *ReadyForServiceResponseBodyData) *ReadyForServiceResponseBody {
	s.Data = v
	return s
}

type ReadyForServiceResponseBodyData struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s ReadyForServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReadyForServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadyForServiceResponseBodyData) SetExtension(v string) *ReadyForServiceResponseBodyData {
	s.Extension = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetWorkMode(v string) *ReadyForServiceResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetDeviceId(v string) *ReadyForServiceResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetJobId(v string) *ReadyForServiceResponseBodyData {
	s.JobId = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetUserId(v string) *ReadyForServiceResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetBreakCode(v string) *ReadyForServiceResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetInstanceId(v string) *ReadyForServiceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetOutboundScenario(v bool) *ReadyForServiceResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetUserState(v string) *ReadyForServiceResponseBodyData {
	s.UserState = &v
	return s
}

func (s *ReadyForServiceResponseBodyData) SetSignedSkillGroupIdList(v []*string) *ReadyForServiceResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

type ReadyForServiceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReadyForServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReadyForServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReadyForServiceResponse) GoString() string {
	return s.String()
}

func (s *ReadyForServiceResponse) SetHeaders(v map[string]*string) *ReadyForServiceResponse {
	s.Headers = v
	return s
}

func (s *ReadyForServiceResponse) SetBody(v *ReadyForServiceResponseBody) *ReadyForServiceResponse {
	s.Body = v
	return s
}

type GetMultiChannelRecordingRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ContactId  *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s GetMultiChannelRecordingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMultiChannelRecordingRequest) GoString() string {
	return s.String()
}

func (s *GetMultiChannelRecordingRequest) SetInstanceId(v string) *GetMultiChannelRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMultiChannelRecordingRequest) SetContactId(v string) *GetMultiChannelRecordingRequest {
	s.ContactId = &v
	return s
}

type GetMultiChannelRecordingResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetMultiChannelRecordingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetMultiChannelRecordingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultiChannelRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiChannelRecordingResponseBody) SetCode(v string) *GetMultiChannelRecordingResponseBody {
	s.Code = &v
	return s
}

func (s *GetMultiChannelRecordingResponseBody) SetHttpStatusCode(v int32) *GetMultiChannelRecordingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMultiChannelRecordingResponseBody) SetMessage(v string) *GetMultiChannelRecordingResponseBody {
	s.Message = &v
	return s
}

func (s *GetMultiChannelRecordingResponseBody) SetRequestId(v string) *GetMultiChannelRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiChannelRecordingResponseBody) SetData(v *GetMultiChannelRecordingResponseBodyData) *GetMultiChannelRecordingResponseBody {
	s.Data = v
	return s
}

type GetMultiChannelRecordingResponseBodyData struct {
	FileUrl  *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetMultiChannelRecordingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMultiChannelRecordingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMultiChannelRecordingResponseBodyData) SetFileUrl(v string) *GetMultiChannelRecordingResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetMultiChannelRecordingResponseBodyData) SetFileName(v string) *GetMultiChannelRecordingResponseBodyData {
	s.FileName = &v
	return s
}

type GetMultiChannelRecordingResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMultiChannelRecordingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMultiChannelRecordingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultiChannelRecordingResponse) GoString() string {
	return s.String()
}

func (s *GetMultiChannelRecordingResponse) SetHeaders(v map[string]*string) *GetMultiChannelRecordingResponse {
	s.Headers = v
	return s
}

func (s *GetMultiChannelRecordingResponse) SetBody(v *GetMultiChannelRecordingResponseBody) *GetMultiChannelRecordingResponse {
	s.Body = v
	return s
}

type BargeInCallRequest struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId       *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	BargedUserId   *string `json:"BargedUserId,omitempty" xml:"BargedUserId,omitempty"`
	TimeoutSeconds *int32  `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s BargeInCallRequest) String() string {
	return tea.Prettify(s)
}

func (s BargeInCallRequest) GoString() string {
	return s.String()
}

func (s *BargeInCallRequest) SetInstanceId(v string) *BargeInCallRequest {
	s.InstanceId = &v
	return s
}

func (s *BargeInCallRequest) SetUserId(v string) *BargeInCallRequest {
	s.UserId = &v
	return s
}

func (s *BargeInCallRequest) SetDeviceId(v string) *BargeInCallRequest {
	s.DeviceId = &v
	return s
}

func (s *BargeInCallRequest) SetJobId(v string) *BargeInCallRequest {
	s.JobId = &v
	return s
}

func (s *BargeInCallRequest) SetBargedUserId(v string) *BargeInCallRequest {
	s.BargedUserId = &v
	return s
}

func (s *BargeInCallRequest) SetTimeoutSeconds(v int32) *BargeInCallRequest {
	s.TimeoutSeconds = &v
	return s
}

type BargeInCallResponseBody struct {
	Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                    `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *BargeInCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s BargeInCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BargeInCallResponseBody) GoString() string {
	return s.String()
}

func (s *BargeInCallResponseBody) SetCode(v string) *BargeInCallResponseBody {
	s.Code = &v
	return s
}

func (s *BargeInCallResponseBody) SetHttpStatusCode(v int32) *BargeInCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BargeInCallResponseBody) SetMessage(v string) *BargeInCallResponseBody {
	s.Message = &v
	return s
}

func (s *BargeInCallResponseBody) SetRequestId(v string) *BargeInCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *BargeInCallResponseBody) SetParams(v []*string) *BargeInCallResponseBody {
	s.Params = v
	return s
}

func (s *BargeInCallResponseBody) SetData(v *BargeInCallResponseBodyData) *BargeInCallResponseBody {
	s.Data = v
	return s
}

type BargeInCallResponseBodyData struct {
	CallContext *BargeInCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *BargeInCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s BargeInCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BargeInCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *BargeInCallResponseBodyData) SetCallContext(v *BargeInCallResponseBodyDataCallContext) *BargeInCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *BargeInCallResponseBodyData) SetUserContext(v *BargeInCallResponseBodyDataUserContext) *BargeInCallResponseBodyData {
	s.UserContext = v
	return s
}

type BargeInCallResponseBodyDataCallContext struct {
	CallType        *string                                                  `json:"CallType,omitempty" xml:"CallType,omitempty"`
	InstanceId      *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string                                                  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelContexts []*BargeInCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s BargeInCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s BargeInCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *BargeInCallResponseBodyDataCallContext) SetCallType(v string) *BargeInCallResponseBodyDataCallContext {
	s.CallType = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContext) SetInstanceId(v string) *BargeInCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContext) SetJobId(v string) *BargeInCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContext) SetChannelContexts(v []*BargeInCallResponseBodyDataCallContextChannelContexts) *BargeInCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type BargeInCallResponseBodyDataCallContextChannelContexts struct {
	Index            *int32                 `json:"Index,omitempty" xml:"Index,omitempty"`
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelFlags     *string                `json:"ChannelFlags,omitempty" xml:"ChannelFlags,omitempty"`
	SkillGroupId     *string                `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
}

func (s BargeInCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s BargeInCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetIndex(v int32) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.Index = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetChannelFlags(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.ChannelFlags = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetSkillGroupId(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.SkillGroupId = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

func (s *BargeInCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *BargeInCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

type BargeInCallResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s BargeInCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s BargeInCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *BargeInCallResponseBodyDataUserContext) SetExtension(v string) *BargeInCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetHeartbeat(v int64) *BargeInCallResponseBodyDataUserContext {
	s.Heartbeat = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetWorkMode(v string) *BargeInCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetDeviceId(v string) *BargeInCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetUserId(v string) *BargeInCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetReserved(v int64) *BargeInCallResponseBodyDataUserContext {
	s.Reserved = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetBreakCode(v string) *BargeInCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetInstanceId(v string) *BargeInCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *BargeInCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetMobile(v string) *BargeInCallResponseBodyDataUserContext {
	s.Mobile = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetJobId(v string) *BargeInCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetUserState(v string) *BargeInCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *BargeInCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *BargeInCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type BargeInCallResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BargeInCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BargeInCallResponse) String() string {
	return tea.Prettify(s)
}

func (s BargeInCallResponse) GoString() string {
	return s.String()
}

func (s *BargeInCallResponse) SetHeaders(v map[string]*string) *BargeInCallResponse {
	s.Headers = v
	return s
}

func (s *BargeInCallResponse) SetBody(v *BargeInCallResponseBody) *BargeInCallResponse {
	s.Body = v
	return s
}

type ListPhoneNumbersOfSkillGroupRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId  *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	IsMember      *bool   `json:"IsMember,omitempty" xml:"IsMember,omitempty"`
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	Active        *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
}

func (s ListPhoneNumbersOfSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneNumbersOfSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetInstanceId(v string) *ListPhoneNumbersOfSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetSkillGroupId(v string) *ListPhoneNumbersOfSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetPageNumber(v int32) *ListPhoneNumbersOfSkillGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetPageSize(v int32) *ListPhoneNumbersOfSkillGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetIsMember(v bool) *ListPhoneNumbersOfSkillGroupRequest {
	s.IsMember = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetSearchPattern(v string) *ListPhoneNumbersOfSkillGroupRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetActive(v bool) *ListPhoneNumbersOfSkillGroupRequest {
	s.Active = &v
	return s
}

type ListPhoneNumbersOfSkillGroupResponseBody struct {
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListPhoneNumbersOfSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListPhoneNumbersOfSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneNumbersOfSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) SetCode(v string) *ListPhoneNumbersOfSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) SetHttpStatusCode(v int32) *ListPhoneNumbersOfSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) SetMessage(v string) *ListPhoneNumbersOfSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) SetRequestId(v string) *ListPhoneNumbersOfSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBody) SetData(v *ListPhoneNumbersOfSkillGroupResponseBodyData) *ListPhoneNumbersOfSkillGroupResponseBody {
	s.Data = v
	return s
}

type ListPhoneNumbersOfSkillGroupResponseBodyData struct {
	PageNumber *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListPhoneNumbersOfSkillGroupResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListPhoneNumbersOfSkillGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneNumbersOfSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) SetPageNumber(v int32) *ListPhoneNumbersOfSkillGroupResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) SetPageSize(v int32) *ListPhoneNumbersOfSkillGroupResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) SetTotalCount(v int32) *ListPhoneNumbersOfSkillGroupResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyData) SetList(v []*ListPhoneNumbersOfSkillGroupResponseBodyDataList) *ListPhoneNumbersOfSkillGroupResponseBodyData {
	s.List = v
	return s
}

type ListPhoneNumbersOfSkillGroupResponseBodyDataList struct {
	Active        *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	Number        *string `json:"Number,omitempty" xml:"Number,omitempty"`
	City          *string `json:"City,omitempty" xml:"City,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Usage         *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	Province      *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListPhoneNumbersOfSkillGroupResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneNumbersOfSkillGroupResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetActive(v bool) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.Active = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetNumber(v string) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetCity(v string) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.City = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetInstanceId(v string) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetUsage(v string) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.Usage = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetContactFlowId(v string) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.ContactFlowId = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponseBodyDataList) SetProvince(v string) *ListPhoneNumbersOfSkillGroupResponseBodyDataList {
	s.Province = &v
	return s
}

type ListPhoneNumbersOfSkillGroupResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPhoneNumbersOfSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPhoneNumbersOfSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPhoneNumbersOfSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersOfSkillGroupResponse) SetHeaders(v map[string]*string) *ListPhoneNumbersOfSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupResponse) SetBody(v *ListPhoneNumbersOfSkillGroupResponseBody) *ListPhoneNumbersOfSkillGroupResponse {
	s.Body = v
	return s
}

type SignOutGroupRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s SignOutGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SignOutGroupRequest) GoString() string {
	return s.String()
}

func (s *SignOutGroupRequest) SetInstanceId(v string) *SignOutGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *SignOutGroupRequest) SetUserId(v string) *SignOutGroupRequest {
	s.UserId = &v
	return s
}

func (s *SignOutGroupRequest) SetDeviceId(v string) *SignOutGroupRequest {
	s.DeviceId = &v
	return s
}

type SignOutGroupResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                     `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *SignOutGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s SignOutGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SignOutGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SignOutGroupResponseBody) SetCode(v string) *SignOutGroupResponseBody {
	s.Code = &v
	return s
}

func (s *SignOutGroupResponseBody) SetHttpStatusCode(v int32) *SignOutGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SignOutGroupResponseBody) SetMessage(v string) *SignOutGroupResponseBody {
	s.Message = &v
	return s
}

func (s *SignOutGroupResponseBody) SetRequestId(v string) *SignOutGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SignOutGroupResponseBody) SetParams(v []*string) *SignOutGroupResponseBody {
	s.Params = v
	return s
}

func (s *SignOutGroupResponseBody) SetData(v *SignOutGroupResponseBodyData) *SignOutGroupResponseBody {
	s.Data = v
	return s
}

type SignOutGroupResponseBodyData struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Heartbeat              *int64    `json:"Heartbeat,omitempty" xml:"Heartbeat,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Reserved               *int64    `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	Mobile                 *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s SignOutGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SignOutGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *SignOutGroupResponseBodyData) SetExtension(v string) *SignOutGroupResponseBodyData {
	s.Extension = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetHeartbeat(v int64) *SignOutGroupResponseBodyData {
	s.Heartbeat = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetWorkMode(v string) *SignOutGroupResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetDeviceId(v string) *SignOutGroupResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetUserId(v string) *SignOutGroupResponseBodyData {
	s.UserId = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetReserved(v int64) *SignOutGroupResponseBodyData {
	s.Reserved = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetBreakCode(v string) *SignOutGroupResponseBodyData {
	s.BreakCode = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetInstanceId(v string) *SignOutGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetOutboundScenario(v bool) *SignOutGroupResponseBodyData {
	s.OutboundScenario = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetMobile(v string) *SignOutGroupResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetJobId(v string) *SignOutGroupResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetUserState(v string) *SignOutGroupResponseBodyData {
	s.UserState = &v
	return s
}

func (s *SignOutGroupResponseBodyData) SetSignedSkillGroupIdList(v []*string) *SignOutGroupResponseBodyData {
	s.SignedSkillGroupIdList = v
	return s
}

type SignOutGroupResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SignOutGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SignOutGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SignOutGroupResponse) GoString() string {
	return s.String()
}

func (s *SignOutGroupResponse) SetHeaders(v map[string]*string) *SignOutGroupResponse {
	s.Headers = v
	return s
}

func (s *SignOutGroupResponse) SetBody(v *SignOutGroupResponseBody) *SignOutGroupResponse {
	s.Body = v
	return s
}

type SaveRTCStatsV2Request struct {
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CallId         *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	GeneralInfo    *string `json:"GeneralInfo,omitempty" xml:"GeneralInfo,omitempty"`
	SenderReport   *string `json:"SenderReport,omitempty" xml:"SenderReport,omitempty"`
	ReceiverReport *string `json:"ReceiverReport,omitempty" xml:"ReceiverReport,omitempty"`
	GoogAddress    *string `json:"GoogAddress,omitempty" xml:"GoogAddress,omitempty"`
}

func (s SaveRTCStatsV2Request) String() string {
	return tea.Prettify(s)
}

func (s SaveRTCStatsV2Request) GoString() string {
	return s.String()
}

func (s *SaveRTCStatsV2Request) SetInstanceId(v string) *SaveRTCStatsV2Request {
	s.InstanceId = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetCallId(v string) *SaveRTCStatsV2Request {
	s.CallId = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetGeneralInfo(v string) *SaveRTCStatsV2Request {
	s.GeneralInfo = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetSenderReport(v string) *SaveRTCStatsV2Request {
	s.SenderReport = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetReceiverReport(v string) *SaveRTCStatsV2Request {
	s.ReceiverReport = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetGoogAddress(v string) *SaveRTCStatsV2Request {
	s.GoogAddress = &v
	return s
}

type SaveRTCStatsV2ResponseBody struct {
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TimeStamp      *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	RowCount       *int64  `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
}

func (s SaveRTCStatsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveRTCStatsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *SaveRTCStatsV2ResponseBody) SetHttpStatusCode(v int64) *SaveRTCStatsV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetRequestId(v string) *SaveRTCStatsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetSuccess(v bool) *SaveRTCStatsV2ResponseBody {
	s.Success = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetCode(v string) *SaveRTCStatsV2ResponseBody {
	s.Code = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetMessage(v string) *SaveRTCStatsV2ResponseBody {
	s.Message = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetTimeStamp(v int64) *SaveRTCStatsV2ResponseBody {
	s.TimeStamp = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetRowCount(v int64) *SaveRTCStatsV2ResponseBody {
	s.RowCount = &v
	return s
}

type SaveRTCStatsV2Response struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveRTCStatsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveRTCStatsV2Response) String() string {
	return tea.Prettify(s)
}

func (s SaveRTCStatsV2Response) GoString() string {
	return s.String()
}

func (s *SaveRTCStatsV2Response) SetHeaders(v map[string]*string) *SaveRTCStatsV2Response {
	s.Headers = v
	return s
}

func (s *SaveRTCStatsV2Response) SetBody(v *SaveRTCStatsV2ResponseBody) *SaveRTCStatsV2Response {
	s.Body = v
	return s
}

type GetHistoricalCallerReportRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	StopTime      *int64  `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetHistoricalCallerReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalCallerReportRequest) GoString() string {
	return s.String()
}

func (s *GetHistoricalCallerReportRequest) SetInstanceId(v string) *GetHistoricalCallerReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHistoricalCallerReportRequest) SetCallingNumber(v string) *GetHistoricalCallerReportRequest {
	s.CallingNumber = &v
	return s
}

func (s *GetHistoricalCallerReportRequest) SetStopTime(v int64) *GetHistoricalCallerReportRequest {
	s.StopTime = &v
	return s
}

func (s *GetHistoricalCallerReportRequest) SetStartTime(v int64) *GetHistoricalCallerReportRequest {
	s.StartTime = &v
	return s
}

type GetHistoricalCallerReportResponseBody struct {
	Code           *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *GetHistoricalCallerReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetHistoricalCallerReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalCallerReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoricalCallerReportResponseBody) SetCode(v string) *GetHistoricalCallerReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHistoricalCallerReportResponseBody) SetHttpStatusCode(v int32) *GetHistoricalCallerReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHistoricalCallerReportResponseBody) SetMessage(v string) *GetHistoricalCallerReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHistoricalCallerReportResponseBody) SetRequestId(v string) *GetHistoricalCallerReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHistoricalCallerReportResponseBody) SetData(v *GetHistoricalCallerReportResponseBodyData) *GetHistoricalCallerReportResponseBody {
	s.Data = v
	return s
}

type GetHistoricalCallerReportResponseBodyData struct {
	LastCallingTime *int64 `json:"LastCallingTime,omitempty" xml:"LastCallingTime,omitempty"`
	TotalCalls      *int64 `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
}

func (s GetHistoricalCallerReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalCallerReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistoricalCallerReportResponseBodyData) SetLastCallingTime(v int64) *GetHistoricalCallerReportResponseBodyData {
	s.LastCallingTime = &v
	return s
}

func (s *GetHistoricalCallerReportResponseBodyData) SetTotalCalls(v int64) *GetHistoricalCallerReportResponseBodyData {
	s.TotalCalls = &v
	return s
}

type GetHistoricalCallerReportResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHistoricalCallerReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHistoricalCallerReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalCallerReportResponse) GoString() string {
	return s.String()
}

func (s *GetHistoricalCallerReportResponse) SetHeaders(v map[string]*string) *GetHistoricalCallerReportResponse {
	s.Headers = v
	return s
}

func (s *GetHistoricalCallerReportResponse) SetBody(v *GetHistoricalCallerReportResponseBody) *GetHistoricalCallerReportResponse {
	s.Body = v
	return s
}

type ModifyUserLevelsOfSkillGroupRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId  *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	UserLevelList *string `json:"UserLevelList,omitempty" xml:"UserLevelList,omitempty"`
}

func (s ModifyUserLevelsOfSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserLevelsOfSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserLevelsOfSkillGroupRequest) SetInstanceId(v string) *ModifyUserLevelsOfSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupRequest) SetSkillGroupId(v string) *ModifyUserLevelsOfSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupRequest) SetUserLevelList(v string) *ModifyUserLevelsOfSkillGroupRequest {
	s.UserLevelList = &v
	return s
}

type ModifyUserLevelsOfSkillGroupResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserLevelsOfSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserLevelsOfSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) SetHttpStatusCode(v int32) *ModifyUserLevelsOfSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) SetCode(v string) *ModifyUserLevelsOfSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) SetMessage(v string) *ModifyUserLevelsOfSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) SetData(v string) *ModifyUserLevelsOfSkillGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponseBody) SetRequestId(v string) *ModifyUserLevelsOfSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserLevelsOfSkillGroupResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUserLevelsOfSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserLevelsOfSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserLevelsOfSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserLevelsOfSkillGroupResponse) SetHeaders(v map[string]*string) *ModifyUserLevelsOfSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserLevelsOfSkillGroupResponse) SetBody(v *ModifyUserLevelsOfSkillGroupResponseBody) *ModifyUserLevelsOfSkillGroupResponse {
	s.Body = v
	return s
}

type SaveTerminalLogRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CallId          *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	JobId           *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DataType        *int32  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UniqueRequestId *string `json:"UniqueRequestId,omitempty" xml:"UniqueRequestId,omitempty"`
	MethodName      *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
}

func (s SaveTerminalLogRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTerminalLogRequest) GoString() string {
	return s.String()
}

func (s *SaveTerminalLogRequest) SetInstanceId(v string) *SaveTerminalLogRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveTerminalLogRequest) SetCallId(v string) *SaveTerminalLogRequest {
	s.CallId = &v
	return s
}

func (s *SaveTerminalLogRequest) SetJobId(v string) *SaveTerminalLogRequest {
	s.JobId = &v
	return s
}

func (s *SaveTerminalLogRequest) SetAppName(v string) *SaveTerminalLogRequest {
	s.AppName = &v
	return s
}

func (s *SaveTerminalLogRequest) SetDataType(v int32) *SaveTerminalLogRequest {
	s.DataType = &v
	return s
}

func (s *SaveTerminalLogRequest) SetContent(v string) *SaveTerminalLogRequest {
	s.Content = &v
	return s
}

func (s *SaveTerminalLogRequest) SetStatus(v string) *SaveTerminalLogRequest {
	s.Status = &v
	return s
}

func (s *SaveTerminalLogRequest) SetUniqueRequestId(v string) *SaveTerminalLogRequest {
	s.UniqueRequestId = &v
	return s
}

func (s *SaveTerminalLogRequest) SetMethodName(v string) *SaveTerminalLogRequest {
	s.MethodName = &v
	return s
}

type SaveTerminalLogResponseBody struct {
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	TimeStamp      *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s SaveTerminalLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTerminalLogResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTerminalLogResponseBody) SetHttpStatusCode(v int64) *SaveTerminalLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetRequestId(v string) *SaveTerminalLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetSuccess(v bool) *SaveTerminalLogResponseBody {
	s.Success = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetCode(v string) *SaveTerminalLogResponseBody {
	s.Code = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetMessage(v string) *SaveTerminalLogResponseBody {
	s.Message = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetTimeStamp(v int64) *SaveTerminalLogResponseBody {
	s.TimeStamp = &v
	return s
}

type SaveTerminalLogResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveTerminalLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTerminalLogResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTerminalLogResponse) GoString() string {
	return s.String()
}

func (s *SaveTerminalLogResponse) SetHeaders(v map[string]*string) *SaveTerminalLogResponse {
	s.Headers = v
	return s
}

func (s *SaveTerminalLogResponse) SetBody(v *SaveTerminalLogResponseBody) *SaveTerminalLogResponse {
	s.Body = v
	return s
}

type ListRealtimeSkillGroupStatesRequest struct {
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListRealtimeSkillGroupStatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeSkillGroupStatesRequest) GoString() string {
	return s.String()
}

func (s *ListRealtimeSkillGroupStatesRequest) SetSkillGroupIdList(v string) *ListRealtimeSkillGroupStatesRequest {
	s.SkillGroupIdList = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesRequest) SetPageNumber(v int32) *ListRealtimeSkillGroupStatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesRequest) SetPageSize(v int32) *ListRealtimeSkillGroupStatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesRequest) SetInstanceId(v string) *ListRealtimeSkillGroupStatesRequest {
	s.InstanceId = &v
	return s
}

type ListRealtimeSkillGroupStatesResponseBody struct {
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *ListRealtimeSkillGroupStatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListRealtimeSkillGroupStatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeSkillGroupStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRealtimeSkillGroupStatesResponseBody) SetCode(v string) *ListRealtimeSkillGroupStatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBody) SetHttpStatusCode(v int32) *ListRealtimeSkillGroupStatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBody) SetMessage(v string) *ListRealtimeSkillGroupStatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBody) SetRequestId(v string) *ListRealtimeSkillGroupStatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBody) SetData(v *ListRealtimeSkillGroupStatesResponseBodyData) *ListRealtimeSkillGroupStatesResponseBody {
	s.Data = v
	return s
}

type ListRealtimeSkillGroupStatesResponseBodyData struct {
	PageNumber *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	List       []*ListRealtimeSkillGroupStatesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListRealtimeSkillGroupStatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeSkillGroupStatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) SetPageNumber(v int32) *ListRealtimeSkillGroupStatesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) SetPageSize(v int32) *ListRealtimeSkillGroupStatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) SetTotalCount(v int32) *ListRealtimeSkillGroupStatesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyData) SetList(v []*ListRealtimeSkillGroupStatesResponseBodyDataList) *ListRealtimeSkillGroupStatesResponseBodyData {
	s.List = v
	return s
}

type ListRealtimeSkillGroupStatesResponseBodyDataList struct {
	WorkingAgents      *int64  `json:"WorkingAgents,omitempty" xml:"WorkingAgents,omitempty"`
	LongestWaitingTime *int64  `json:"LongestWaitingTime,omitempty" xml:"LongestWaitingTime,omitempty"`
	LoggedInAgents     *int64  `json:"LoggedInAgents,omitempty" xml:"LoggedInAgents,omitempty"`
	WaitingCalls       *int64  `json:"WaitingCalls,omitempty" xml:"WaitingCalls,omitempty"`
	BreakingAgents     *int64  `json:"BreakingAgents,omitempty" xml:"BreakingAgents,omitempty"`
	TalkingAgents      *int64  `json:"TalkingAgents,omitempty" xml:"TalkingAgents,omitempty"`
	SkillGroupName     *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
	SkillGroupId       *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	ReadyAgents        *int64  `json:"ReadyAgents,omitempty" xml:"ReadyAgents,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListRealtimeSkillGroupStatesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeSkillGroupStatesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetWorkingAgents(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.WorkingAgents = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetLongestWaitingTime(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.LongestWaitingTime = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetLoggedInAgents(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.LoggedInAgents = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetWaitingCalls(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.WaitingCalls = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetBreakingAgents(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.BreakingAgents = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetTalkingAgents(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.TalkingAgents = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetSkillGroupName(v string) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetSkillGroupId(v string) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetReadyAgents(v int64) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.ReadyAgents = &v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponseBodyDataList) SetInstanceId(v string) *ListRealtimeSkillGroupStatesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

type ListRealtimeSkillGroupStatesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRealtimeSkillGroupStatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRealtimeSkillGroupStatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRealtimeSkillGroupStatesResponse) GoString() string {
	return s.String()
}

func (s *ListRealtimeSkillGroupStatesResponse) SetHeaders(v map[string]*string) *ListRealtimeSkillGroupStatesResponse {
	s.Headers = v
	return s
}

func (s *ListRealtimeSkillGroupStatesResponse) SetBody(v *ListRealtimeSkillGroupStatesResponseBody) *ListRealtimeSkillGroupStatesResponse {
	s.Body = v
	return s
}

type CreateSkillGroupRequest struct {
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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

func (s *CreateSkillGroupRequest) SetName(v string) *CreateSkillGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateSkillGroupRequest) SetDisplayName(v string) *CreateSkillGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateSkillGroupRequest) SetDescription(v string) *CreateSkillGroupRequest {
	s.Description = &v
	return s
}

type CreateSkillGroupResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           *CreateSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s CreateSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupResponseBody) SetCode(v string) *CreateSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetHttpStatusCode(v int32) *CreateSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetMessage(v string) *CreateSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetRequestId(v string) *CreateSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetData(v *CreateSkillGroupResponseBodyData) *CreateSkillGroupResponseBody {
	s.Data = v
	return s
}

type CreateSkillGroupResponseBodyData struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s CreateSkillGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupResponseBodyData) SetInstanceId(v string) *CreateSkillGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateSkillGroupResponseBodyData) SetDescription(v string) *CreateSkillGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateSkillGroupResponseBodyData) SetName(v string) *CreateSkillGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateSkillGroupResponseBodyData) SetSkillGroupId(v string) *CreateSkillGroupResponseBodyData {
	s.SkillGroupId = &v
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

type PickOutboundNumbersRequest struct {
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CalledNumber     *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	Count            *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
}

func (s PickOutboundNumbersRequest) String() string {
	return tea.Prettify(s)
}

func (s PickOutboundNumbersRequest) GoString() string {
	return s.String()
}

func (s *PickOutboundNumbersRequest) SetInstanceId(v string) *PickOutboundNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *PickOutboundNumbersRequest) SetCalledNumber(v string) *PickOutboundNumbersRequest {
	s.CalledNumber = &v
	return s
}

func (s *PickOutboundNumbersRequest) SetCount(v int32) *PickOutboundNumbersRequest {
	s.Count = &v
	return s
}

func (s *PickOutboundNumbersRequest) SetSkillGroupIdList(v string) *PickOutboundNumbersRequest {
	s.SkillGroupIdList = &v
	return s
}

type PickOutboundNumbersResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data           []*PickOutboundNumbersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s PickOutboundNumbersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PickOutboundNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *PickOutboundNumbersResponseBody) SetCode(v string) *PickOutboundNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *PickOutboundNumbersResponseBody) SetHttpStatusCode(v int32) *PickOutboundNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PickOutboundNumbersResponseBody) SetMessage(v string) *PickOutboundNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *PickOutboundNumbersResponseBody) SetRequestId(v string) *PickOutboundNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *PickOutboundNumbersResponseBody) SetData(v []*PickOutboundNumbersResponseBodyData) *PickOutboundNumbersResponseBody {
	s.Data = v
	return s
}

type PickOutboundNumbersResponseBodyData struct {
	Callee *PickOutboundNumbersResponseBodyDataCallee `json:"Callee,omitempty" xml:"Callee,omitempty" type:"Struct"`
	Caller *PickOutboundNumbersResponseBodyDataCaller `json:"Caller,omitempty" xml:"Caller,omitempty" type:"Struct"`
}

func (s PickOutboundNumbersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PickOutboundNumbersResponseBodyData) GoString() string {
	return s.String()
}

func (s *PickOutboundNumbersResponseBodyData) SetCallee(v *PickOutboundNumbersResponseBodyDataCallee) *PickOutboundNumbersResponseBodyData {
	s.Callee = v
	return s
}

func (s *PickOutboundNumbersResponseBodyData) SetCaller(v *PickOutboundNumbersResponseBodyDataCaller) *PickOutboundNumbersResponseBodyData {
	s.Caller = v
	return s
}

type PickOutboundNumbersResponseBodyDataCallee struct {
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s PickOutboundNumbersResponseBodyDataCallee) String() string {
	return tea.Prettify(s)
}

func (s PickOutboundNumbersResponseBodyDataCallee) GoString() string {
	return s.String()
}

func (s *PickOutboundNumbersResponseBodyDataCallee) SetNumber(v string) *PickOutboundNumbersResponseBodyDataCallee {
	s.Number = &v
	return s
}

func (s *PickOutboundNumbersResponseBodyDataCallee) SetCity(v string) *PickOutboundNumbersResponseBodyDataCallee {
	s.City = &v
	return s
}

func (s *PickOutboundNumbersResponseBodyDataCallee) SetProvince(v string) *PickOutboundNumbersResponseBodyDataCallee {
	s.Province = &v
	return s
}

type PickOutboundNumbersResponseBodyDataCaller struct {
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s PickOutboundNumbersResponseBodyDataCaller) String() string {
	return tea.Prettify(s)
}

func (s PickOutboundNumbersResponseBodyDataCaller) GoString() string {
	return s.String()
}

func (s *PickOutboundNumbersResponseBodyDataCaller) SetNumber(v string) *PickOutboundNumbersResponseBodyDataCaller {
	s.Number = &v
	return s
}

func (s *PickOutboundNumbersResponseBodyDataCaller) SetCity(v string) *PickOutboundNumbersResponseBodyDataCaller {
	s.City = &v
	return s
}

func (s *PickOutboundNumbersResponseBodyDataCaller) SetProvince(v string) *PickOutboundNumbersResponseBodyDataCaller {
	s.Province = &v
	return s
}

type PickOutboundNumbersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PickOutboundNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PickOutboundNumbersResponse) String() string {
	return tea.Prettify(s)
}

func (s PickOutboundNumbersResponse) GoString() string {
	return s.String()
}

func (s *PickOutboundNumbersResponse) SetHeaders(v map[string]*string) *PickOutboundNumbersResponse {
	s.Headers = v
	return s
}

func (s *PickOutboundNumbersResponse) SetBody(v *PickOutboundNumbersResponseBody) *PickOutboundNumbersResponse {
	s.Body = v
	return s
}

type ReleaseCallRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s ReleaseCallRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCallRequest) GoString() string {
	return s.String()
}

func (s *ReleaseCallRequest) SetInstanceId(v string) *ReleaseCallRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseCallRequest) SetUserId(v string) *ReleaseCallRequest {
	s.UserId = &v
	return s
}

func (s *ReleaseCallRequest) SetDeviceId(v string) *ReleaseCallRequest {
	s.DeviceId = &v
	return s
}

func (s *ReleaseCallRequest) SetJobId(v string) *ReleaseCallRequest {
	s.JobId = &v
	return s
}

func (s *ReleaseCallRequest) SetChannelId(v string) *ReleaseCallRequest {
	s.ChannelId = &v
	return s
}

type ReleaseCallResponseBody struct {
	Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                    `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *ReleaseCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ReleaseCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCallResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseCallResponseBody) SetCode(v string) *ReleaseCallResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseCallResponseBody) SetHttpStatusCode(v int32) *ReleaseCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleaseCallResponseBody) SetMessage(v string) *ReleaseCallResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseCallResponseBody) SetRequestId(v string) *ReleaseCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseCallResponseBody) SetParams(v []*string) *ReleaseCallResponseBody {
	s.Params = v
	return s
}

func (s *ReleaseCallResponseBody) SetData(v *ReleaseCallResponseBodyData) *ReleaseCallResponseBody {
	s.Data = v
	return s
}

type ReleaseCallResponseBodyData struct {
	CallContext *ReleaseCallResponseBodyDataCallContext `json:"CallContext,omitempty" xml:"CallContext,omitempty" type:"Struct"`
	UserContext *ReleaseCallResponseBodyDataUserContext `json:"UserContext,omitempty" xml:"UserContext,omitempty" type:"Struct"`
}

func (s ReleaseCallResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReleaseCallResponseBodyData) SetCallContext(v *ReleaseCallResponseBodyDataCallContext) *ReleaseCallResponseBodyData {
	s.CallContext = v
	return s
}

func (s *ReleaseCallResponseBodyData) SetUserContext(v *ReleaseCallResponseBodyDataUserContext) *ReleaseCallResponseBodyData {
	s.UserContext = v
	return s
}

type ReleaseCallResponseBodyDataCallContext struct {
	JobId           *string                                                  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	InstanceId      *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ChannelContexts []*ReleaseCallResponseBodyDataCallContextChannelContexts `json:"ChannelContexts,omitempty" xml:"ChannelContexts,omitempty" type:"Repeated"`
}

func (s ReleaseCallResponseBodyDataCallContext) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCallResponseBodyDataCallContext) GoString() string {
	return s.String()
}

func (s *ReleaseCallResponseBodyDataCallContext) SetJobId(v string) *ReleaseCallResponseBodyDataCallContext {
	s.JobId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContext) SetInstanceId(v string) *ReleaseCallResponseBodyDataCallContext {
	s.InstanceId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContext) SetChannelContexts(v []*ReleaseCallResponseBodyDataCallContextChannelContexts) *ReleaseCallResponseBodyDataCallContext {
	s.ChannelContexts = v
	return s
}

type ReleaseCallResponseBodyDataCallContextChannelContexts struct {
	ReleaseInitiator *string                `json:"ReleaseInitiator,omitempty" xml:"ReleaseInitiator,omitempty"`
	ChannelState     *string                `json:"ChannelState,omitempty" xml:"ChannelState,omitempty"`
	Destination      *string                `json:"Destination,omitempty" xml:"Destination,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Timestamp        *int64                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	AssociatedData   map[string]interface{} `json:"AssociatedData,omitempty" xml:"AssociatedData,omitempty"`
	ReleaseReason    *string                `json:"ReleaseReason,omitempty" xml:"ReleaseReason,omitempty"`
	CallType         *string                `json:"CallType,omitempty" xml:"CallType,omitempty"`
	JobId            *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ChannelId        *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Originator       *string                `json:"Originator,omitempty" xml:"Originator,omitempty"`
	UserExtension    *string                `json:"UserExtension,omitempty" xml:"UserExtension,omitempty"`
}

func (s ReleaseCallResponseBodyDataCallContextChannelContexts) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCallResponseBodyDataCallContextChannelContexts) GoString() string {
	return s.String()
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetReleaseInitiator(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseInitiator = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetChannelState(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.ChannelState = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetDestination(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.Destination = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetUserId(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.UserId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetTimestamp(v int64) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.Timestamp = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetAssociatedData(v map[string]interface{}) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.AssociatedData = v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetReleaseReason(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.ReleaseReason = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetCallType(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.CallType = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetJobId(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.JobId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetChannelId(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.ChannelId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetOriginator(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.Originator = &v
	return s
}

func (s *ReleaseCallResponseBodyDataCallContextChannelContexts) SetUserExtension(v string) *ReleaseCallResponseBodyDataCallContextChannelContexts {
	s.UserExtension = &v
	return s
}

type ReleaseCallResponseBodyDataUserContext struct {
	Extension              *string   `json:"Extension,omitempty" xml:"Extension,omitempty"`
	WorkMode               *string   `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	DeviceId               *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	JobId                  *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	UserId                 *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BreakCode              *string   `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OutboundScenario       *bool     `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	UserState              *string   `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SignedSkillGroupIdList []*string `json:"SignedSkillGroupIdList,omitempty" xml:"SignedSkillGroupIdList,omitempty" type:"Repeated"`
}

func (s ReleaseCallResponseBodyDataUserContext) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCallResponseBodyDataUserContext) GoString() string {
	return s.String()
}

func (s *ReleaseCallResponseBodyDataUserContext) SetExtension(v string) *ReleaseCallResponseBodyDataUserContext {
	s.Extension = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetWorkMode(v string) *ReleaseCallResponseBodyDataUserContext {
	s.WorkMode = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetDeviceId(v string) *ReleaseCallResponseBodyDataUserContext {
	s.DeviceId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetJobId(v string) *ReleaseCallResponseBodyDataUserContext {
	s.JobId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetUserId(v string) *ReleaseCallResponseBodyDataUserContext {
	s.UserId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetBreakCode(v string) *ReleaseCallResponseBodyDataUserContext {
	s.BreakCode = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetInstanceId(v string) *ReleaseCallResponseBodyDataUserContext {
	s.InstanceId = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetOutboundScenario(v bool) *ReleaseCallResponseBodyDataUserContext {
	s.OutboundScenario = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetUserState(v string) *ReleaseCallResponseBodyDataUserContext {
	s.UserState = &v
	return s
}

func (s *ReleaseCallResponseBodyDataUserContext) SetSignedSkillGroupIdList(v []*string) *ReleaseCallResponseBodyDataUserContext {
	s.SignedSkillGroupIdList = v
	return s
}

type ReleaseCallResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseCallResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCallResponse) GoString() string {
	return s.String()
}

func (s *ReleaseCallResponse) SetHeaders(v map[string]*string) *ReleaseCallResponse {
	s.Headers = v
	return s
}

func (s *ReleaseCallResponse) SetBody(v *ReleaseCallResponseBody) *ReleaseCallResponse {
	s.Body = v
	return s
}

type GetLoginDetailsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetLoginDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetLoginDetailsRequest) SetInstanceId(v string) *GetLoginDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLoginDetailsRequest) SetUserId(v string) *GetLoginDetailsRequest {
	s.UserId = &v
	return s
}

type GetLoginDetailsResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Params         []*string                        `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	Data           *GetLoginDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetLoginDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoginDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginDetailsResponseBody) SetCode(v string) *GetLoginDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *GetLoginDetailsResponseBody) SetHttpStatusCode(v int32) *GetLoginDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLoginDetailsResponseBody) SetMessage(v string) *GetLoginDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *GetLoginDetailsResponseBody) SetRequestId(v string) *GetLoginDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginDetailsResponseBody) SetParams(v []*string) *GetLoginDetailsResponseBody {
	s.Params = v
	return s
}

func (s *GetLoginDetailsResponseBody) SetData(v *GetLoginDetailsResponseBodyData) *GetLoginDetailsResponseBody {
	s.Data = v
	return s
}

type GetLoginDetailsResponseBodyData struct {
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Extension      *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	Signature      *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	SipServerUrl   *string `json:"SipServerUrl,omitempty" xml:"SipServerUrl,omitempty"`
	DeviceId       *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	AgentServerUrl *string `json:"AgentServerUrl,omitempty" xml:"AgentServerUrl,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserKey        *string `json:"UserKey,omitempty" xml:"UserKey,omitempty"`
}

func (s GetLoginDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLoginDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLoginDetailsResponseBodyData) SetDisplayName(v string) *GetLoginDetailsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetExtension(v string) *GetLoginDetailsResponseBodyData {
	s.Extension = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetSignature(v string) *GetLoginDetailsResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetSipServerUrl(v string) *GetLoginDetailsResponseBodyData {
	s.SipServerUrl = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetDeviceId(v string) *GetLoginDetailsResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetAgentServerUrl(v string) *GetLoginDetailsResponseBodyData {
	s.AgentServerUrl = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetUserId(v string) *GetLoginDetailsResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetLoginDetailsResponseBodyData) SetUserKey(v string) *GetLoginDetailsResponseBodyData {
	s.UserKey = &v
	return s
}

type GetLoginDetailsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLoginDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLoginDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoginDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetLoginDetailsResponse) SetHeaders(v map[string]*string) *GetLoginDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetLoginDetailsResponse) SetBody(v *GetLoginDetailsResponseBody) *GetLoginDetailsResponse {
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
		"ap-northeast-1":        tea.String("ccc.aliyuncs.com"),
		"ap-south-1":            tea.String("ccc.aliyuncs.com"),
		"ap-southeast-1":        tea.String("ccc.aliyuncs.com"),
		"ap-southeast-2":        tea.String("ccc.aliyuncs.com"),
		"ap-southeast-3":        tea.String("ccc.aliyuncs.com"),
		"ap-southeast-5":        tea.String("ccc.aliyuncs.com"),
		"cn-beijing":            tea.String("ccc.aliyuncs.com"),
		"cn-chengdu":            tea.String("ccc.aliyuncs.com"),
		"cn-hongkong":           tea.String("ccc.aliyuncs.com"),
		"cn-huhehaote":          tea.String("ccc.aliyuncs.com"),
		"cn-qingdao":            tea.String("ccc.aliyuncs.com"),
		"cn-shenzhen":           tea.String("ccc.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("ccc.aliyuncs.com"),
		"eu-central-1":          tea.String("ccc.aliyuncs.com"),
		"eu-west-1":             tea.String("ccc.aliyuncs.com"),
		"me-east-1":             tea.String("ccc.aliyuncs.com"),
		"us-east-1":             tea.String("ccc.aliyuncs.com"),
		"us-west-1":             tea.String("ccc.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("ccc.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("ccc.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("ccc.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("ccc.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ccc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddSkillGroupsToUserWithOptions(request *AddSkillGroupsToUserRequest, runtime *util.RuntimeOptions) (_result *AddSkillGroupsToUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddSkillGroupsToUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddSkillGroupsToUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSkillGroupsToUser(request *AddSkillGroupsToUserRequest) (_result *AddSkillGroupsToUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSkillGroupsToUserResponse{}
	_body, _err := client.AddSkillGroupsToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveWebRTCStatsWithOptions(request *SaveWebRTCStatsRequest, runtime *util.RuntimeOptions) (_result *SaveWebRTCStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveWebRTCStatsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveWebRTCStats"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveWebRTCStats(request *SaveWebRTCStatsRequest) (_result *SaveWebRTCStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveWebRTCStatsResponse{}
	_body, _err := client.SaveWebRTCStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMonoRecordingWithOptions(request *GetMonoRecordingRequest, runtime *util.RuntimeOptions) (_result *GetMonoRecordingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMonoRecordingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMonoRecording"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMonoRecording(request *GetMonoRecordingRequest) (_result *GetMonoRecordingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMonoRecordingResponse{}
	_body, _err := client.GetMonoRecordingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUsers"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAgentStateLogsWithOptions(request *ListAgentStateLogsRequest, runtime *util.RuntimeOptions) (_result *ListAgentStateLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAgentStateLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAgentStateLogs"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAgentStateLogs(request *ListAgentStateLogsRequest) (_result *ListAgentStateLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAgentStateLogsResponse{}
	_body, _err := client.ListAgentStateLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemovePhoneNumberFromSkillGroupsWithOptions(request *RemovePhoneNumberFromSkillGroupsRequest, runtime *util.RuntimeOptions) (_result *RemovePhoneNumberFromSkillGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemovePhoneNumberFromSkillGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemovePhoneNumberFromSkillGroups"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemovePhoneNumberFromSkillGroups(request *RemovePhoneNumberFromSkillGroupsRequest) (_result *RemovePhoneNumberFromSkillGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemovePhoneNumberFromSkillGroupsResponse{}
	_body, _err := client.RemovePhoneNumberFromSkillGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPhoneNumbersWithOptions(request *ListPhoneNumbersRequest, runtime *util.RuntimeOptions) (_result *ListPhoneNumbersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPhoneNumbersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPhoneNumbers"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPhoneNumbers(request *ListPhoneNumbersRequest) (_result *ListPhoneNumbersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPhoneNumbersResponse{}
	_body, _err := client.ListPhoneNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddNumbersToSkillGroupWithOptions(request *AddNumbersToSkillGroupRequest, runtime *util.RuntimeOptions) (_result *AddNumbersToSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddNumbersToSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddNumbersToSkillGroup"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddNumbersToSkillGroup(request *AddNumbersToSkillGroupRequest) (_result *AddNumbersToSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddNumbersToSkillGroupResponse{}
	_body, _err := client.AddNumbersToSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAgentStateWithOptions(request *ResetAgentStateRequest, runtime *util.RuntimeOptions) (_result *ResetAgentStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetAgentStateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetAgentState"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAgentState(request *ResetAgentStateRequest) (_result *ResetAgentStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAgentStateResponse{}
	_body, _err := client.ResetAgentStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeWorkModeWithOptions(request *ChangeWorkModeRequest, runtime *util.RuntimeOptions) (_result *ChangeWorkModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ChangeWorkModeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ChangeWorkMode"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeWorkMode(request *ChangeWorkModeRequest) (_result *ChangeWorkModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeWorkModeResponse{}
	_body, _err := client.ChangeWorkModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTurnCredentialsWithOptions(request *GetTurnCredentialsRequest, runtime *util.RuntimeOptions) (_result *GetTurnCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTurnCredentialsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTurnCredentials"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTurnCredentials(request *GetTurnCredentialsRequest) (_result *GetTurnCredentialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTurnCredentialsResponse{}
	_body, _err := client.GetTurnCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddPhoneNumbersWithOptions(request *AddPhoneNumbersRequest, runtime *util.RuntimeOptions) (_result *AddPhoneNumbersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddPhoneNumbersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddPhoneNumbers"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddPhoneNumbers(request *AddPhoneNumbersRequest) (_result *AddPhoneNumbersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddPhoneNumbersResponse{}
	_body, _err := client.AddPhoneNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveWebRtcInfoWithOptions(request *SaveWebRtcInfoRequest, runtime *util.RuntimeOptions) (_result *SaveWebRtcInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveWebRtcInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveWebRtcInfo"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveWebRtcInfo(request *SaveWebRtcInfoRequest) (_result *SaveWebRtcInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveWebRtcInfoResponse{}
	_body, _err := client.SaveWebRtcInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIntervalSkillGroupReportWithOptions(request *ListIntervalSkillGroupReportRequest, runtime *util.RuntimeOptions) (_result *ListIntervalSkillGroupReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListIntervalSkillGroupReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListIntervalSkillGroupReport"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIntervalSkillGroupReport(request *ListIntervalSkillGroupReportRequest) (_result *ListIntervalSkillGroupReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIntervalSkillGroupReportResponse{}
	_body, _err := client.ListIntervalSkillGroupReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MonitorCallWithOptions(request *MonitorCallRequest, runtime *util.RuntimeOptions) (_result *MonitorCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MonitorCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MonitorCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MonitorCall(request *MonitorCallRequest) (_result *MonitorCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MonitorCallResponse{}
	_body, _err := client.MonitorCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveUsersFromSkillGroupWithOptions(request *RemoveUsersFromSkillGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveUsersFromSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveUsersFromSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveUsersFromSkillGroup"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUsersFromSkillGroup(request *RemoveUsersFromSkillGroupRequest) (_result *RemoveUsersFromSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUsersFromSkillGroupResponse{}
	_body, _err := client.RemoveUsersFromSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSkillGroupWithOptions(request *DeleteSkillGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSkillGroup"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSkillGroup(request *DeleteSkillGroupRequest) (_result *DeleteSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSkillGroupResponse{}
	_body, _err := client.DeleteSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BlindTransferWithOptions(request *BlindTransferRequest, runtime *util.RuntimeOptions) (_result *BlindTransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BlindTransferResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BlindTransfer"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BlindTransfer(request *BlindTransferRequest) (_result *BlindTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BlindTransferResponse{}
	_body, _err := client.BlindTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSkillLevelsOfUserWithOptions(request *ListSkillLevelsOfUserRequest, runtime *util.RuntimeOptions) (_result *ListSkillLevelsOfUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSkillLevelsOfUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSkillLevelsOfUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSkillLevelsOfUser(request *ListSkillLevelsOfUserRequest) (_result *ListSkillLevelsOfUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSkillLevelsOfUserResponse{}
	_body, _err := client.ListSkillLevelsOfUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUnassignedNumbersWithOptions(request *ListUnassignedNumbersRequest, runtime *util.RuntimeOptions) (_result *ListUnassignedNumbersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUnassignedNumbersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUnassignedNumbers"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUnassignedNumbers(request *ListUnassignedNumbersRequest) (_result *ListUnassignedNumbersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUnassignedNumbersResponse{}
	_body, _err := client.ListUnassignedNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceTrendingReportWithOptions(request *GetInstanceTrendingReportRequest, runtime *util.RuntimeOptions) (_result *GetInstanceTrendingReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetInstanceTrendingReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetInstanceTrendingReport"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceTrendingReport(request *GetInstanceTrendingReportRequest) (_result *GetInstanceTrendingReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceTrendingReportResponse{}
	_body, _err := client.GetInstanceTrendingReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesOfUserWithOptions(request *ListInstancesOfUserRequest, runtime *util.RuntimeOptions) (_result *ListInstancesOfUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInstancesOfUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstancesOfUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstancesOfUser(request *ListInstancesOfUserRequest) (_result *ListInstancesOfUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesOfUserResponse{}
	_body, _err := client.ListInstancesOfUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LaunchSurveyWithOptions(request *LaunchSurveyRequest, runtime *util.RuntimeOptions) (_result *LaunchSurveyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LaunchSurveyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LaunchSurvey"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LaunchSurvey(request *LaunchSurveyRequest) (_result *LaunchSurveyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LaunchSurveyResponse{}
	_body, _err := client.LaunchSurveyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIvrTrackingDetailsWithOptions(request *ListIvrTrackingDetailsRequest, runtime *util.RuntimeOptions) (_result *ListIvrTrackingDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListIvrTrackingDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListIvrTrackingDetails"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIvrTrackingDetails(request *ListIvrTrackingDetailsRequest) (_result *ListIvrTrackingDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIvrTrackingDetailsResponse{}
	_body, _err := client.ListIvrTrackingDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBriefSkillGroupsWithOptions(request *ListBriefSkillGroupsRequest, runtime *util.RuntimeOptions) (_result *ListBriefSkillGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListBriefSkillGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListBriefSkillGroups"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBriefSkillGroups(request *ListBriefSkillGroupsRequest) (_result *ListBriefSkillGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBriefSkillGroupsResponse{}
	_body, _err := client.ListBriefSkillGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnmuteCallWithOptions(request *UnmuteCallRequest, runtime *util.RuntimeOptions) (_result *UnmuteCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnmuteCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnmuteCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnmuteCall(request *UnmuteCallRequest) (_result *UnmuteCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnmuteCallResponse{}
	_body, _err := client.UnmuteCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySkillLevelsOfUserWithOptions(request *ModifySkillLevelsOfUserRequest, runtime *util.RuntimeOptions) (_result *ModifySkillLevelsOfUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySkillLevelsOfUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySkillLevelsOfUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySkillLevelsOfUser(request *ModifySkillLevelsOfUserRequest) (_result *ModifySkillLevelsOfUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySkillLevelsOfUserResponse{}
	_body, _err := client.ModifySkillLevelsOfUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssignUsersWithOptions(request *AssignUsersRequest, runtime *util.RuntimeOptions) (_result *AssignUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AssignUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AssignUsers"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssignUsers(request *AssignUsersRequest) (_result *AssignUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssignUsersResponse{}
	_body, _err := client.AssignUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserLevelsOfSkillGroupWithOptions(request *ListUserLevelsOfSkillGroupRequest, runtime *util.RuntimeOptions) (_result *ListUserLevelsOfSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserLevelsOfSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserLevelsOfSkillGroup"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserLevelsOfSkillGroup(request *ListUserLevelsOfSkillGroupRequest) (_result *ListUserLevelsOfSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserLevelsOfSkillGroupResponse{}
	_body, _err := client.ListUserLevelsOfSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRolesWithOptions(request *ListRolesRequest, runtime *util.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRoles"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRoles(request *ListRolesRequest) (_result *ListRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConfigItemsWithOptions(request *UpdateConfigItemsRequest, runtime *util.RuntimeOptions) (_result *UpdateConfigItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateConfigItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateConfigItems"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConfigItems(request *UpdateConfigItemsRequest) (_result *UpdateConfigItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConfigItemsResponse{}
	_body, _err := client.UpdateConfigItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCallDetailRecordWithOptions(request *GetCallDetailRecordRequest, runtime *util.RuntimeOptions) (_result *GetCallDetailRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCallDetailRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCallDetailRecord"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCallDetailRecord(request *GetCallDetailRecordRequest) (_result *GetCallDetailRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCallDetailRecordResponse{}
	_body, _err := client.GetCallDetailRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPhoneNumberWithOptions(request *ModifyPhoneNumberRequest, runtime *util.RuntimeOptions) (_result *ModifyPhoneNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyPhoneNumberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyPhoneNumber"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPhoneNumber(request *ModifyPhoneNumberRequest) (_result *ModifyPhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPhoneNumberResponse{}
	_body, _err := client.ModifyPhoneNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CoachCallWithOptions(request *CoachCallRequest, runtime *util.RuntimeOptions) (_result *CoachCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CoachCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CoachCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CoachCall(request *CoachCallRequest) (_result *CoachCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CoachCallResponse{}
	_body, _err := client.CoachCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPrivilegesOfUserWithOptions(request *ListPrivilegesOfUserRequest, runtime *util.RuntimeOptions) (_result *ListPrivilegesOfUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPrivilegesOfUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPrivilegesOfUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPrivilegesOfUser(request *ListPrivilegesOfUserRequest) (_result *ListPrivilegesOfUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrivilegesOfUserResponse{}
	_body, _err := client.ListPrivilegesOfUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddPersonalNumbersToUserWithOptions(request *AddPersonalNumbersToUserRequest, runtime *util.RuntimeOptions) (_result *AddPersonalNumbersToUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddPersonalNumbersToUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddPersonalNumbersToUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddPersonalNumbersToUser(request *AddPersonalNumbersToUserRequest) (_result *AddPersonalNumbersToUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddPersonalNumbersToUserResponse{}
	_body, _err := client.AddPersonalNumbersToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHistoricalAgentReportWithOptions(request *ListHistoricalAgentReportRequest, runtime *util.RuntimeOptions) (_result *ListHistoricalAgentReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHistoricalAgentReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHistoricalAgentReport"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHistoricalAgentReport(request *ListHistoricalAgentReportRequest) (_result *ListHistoricalAgentReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHistoricalAgentReportResponse{}
	_body, _err := client.ListHistoricalAgentReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InterceptCallWithOptions(request *InterceptCallRequest, runtime *util.RuntimeOptions) (_result *InterceptCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InterceptCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InterceptCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InterceptCall(request *InterceptCallRequest) (_result *InterceptCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InterceptCallResponse{}
	_body, _err := client.InterceptCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContactFlowsWithOptions(request *ListContactFlowsRequest, runtime *util.RuntimeOptions) (_result *ListContactFlowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListContactFlowsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListContactFlows"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContactFlows(request *ListContactFlowsRequest) (_result *ListContactFlowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContactFlowsResponse{}
	_body, _err := client.ListContactFlowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPersonalNumbersOfUserWithOptions(request *ListPersonalNumbersOfUserRequest, runtime *util.RuntimeOptions) (_result *ListPersonalNumbersOfUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPersonalNumbersOfUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPersonalNumbersOfUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPersonalNumbersOfUser(request *ListPersonalNumbersOfUserRequest) (_result *ListPersonalNumbersOfUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPersonalNumbersOfUserResponse{}
	_body, _err := client.ListPersonalNumbersOfUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartPredictiveCallWithOptions(request *StartPredictiveCallRequest, runtime *util.RuntimeOptions) (_result *StartPredictiveCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartPredictiveCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartPredictiveCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartPredictiveCall(request *StartPredictiveCallRequest) (_result *StartPredictiveCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartPredictiveCallResponse{}
	_body, _err := client.StartPredictiveCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIntervalInstanceReportWithOptions(request *ListIntervalInstanceReportRequest, runtime *util.RuntimeOptions) (_result *ListIntervalInstanceReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListIntervalInstanceReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListIntervalInstanceReport"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIntervalInstanceReport(request *ListIntervalInstanceReportRequest) (_result *ListIntervalInstanceReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIntervalInstanceReportResponse{}
	_body, _err := client.ListIntervalInstanceReportWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CreateInstance"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RemoveSkillGroupsFromUserWithOptions(request *RemoveSkillGroupsFromUserRequest, runtime *util.RuntimeOptions) (_result *RemoveSkillGroupsFromUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveSkillGroupsFromUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveSkillGroupsFromUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSkillGroupsFromUser(request *RemoveSkillGroupsFromUserRequest) (_result *RemoveSkillGroupsFromUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveSkillGroupsFromUserResponse{}
	_body, _err := client.RemoveSkillGroupsFromUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRealtimeAgentStatesWithOptions(request *ListRealtimeAgentStatesRequest, runtime *util.RuntimeOptions) (_result *ListRealtimeAgentStatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRealtimeAgentStatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRealtimeAgentStates"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRealtimeAgentStates(request *ListRealtimeAgentStatesRequest) (_result *ListRealtimeAgentStatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRealtimeAgentStatesResponse{}
	_body, _err := client.ListRealtimeAgentStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LaunchAuthenticationWithOptions(request *LaunchAuthenticationRequest, runtime *util.RuntimeOptions) (_result *LaunchAuthenticationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LaunchAuthenticationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LaunchAuthentication"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LaunchAuthentication(request *LaunchAuthenticationRequest) (_result *LaunchAuthenticationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LaunchAuthenticationResponse{}
	_body, _err := client.LaunchAuthenticationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstances"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHistoricalInstanceReportWithOptions(request *GetHistoricalInstanceReportRequest, runtime *util.RuntimeOptions) (_result *GetHistoricalInstanceReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHistoricalInstanceReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHistoricalInstanceReport"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHistoricalInstanceReport(request *GetHistoricalInstanceReportRequest) (_result *GetHistoricalInstanceReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHistoricalInstanceReportResponse{}
	_body, _err := client.GetHistoricalInstanceReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveUsersWithOptions(request *RemoveUsersRequest, runtime *util.RuntimeOptions) (_result *RemoveUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveUsers"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUsers(request *RemoveUsersRequest) (_result *RemoveUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUsersResponse{}
	_body, _err := client.RemoveUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartBack2BackCallWithOptions(request *StartBack2BackCallRequest, runtime *util.RuntimeOptions) (_result *StartBack2BackCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartBack2BackCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartBack2BackCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartBack2BackCall(request *StartBack2BackCallRequest) (_result *StartBack2BackCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartBack2BackCallResponse{}
	_body, _err := client.StartBack2BackCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemovePhoneNumbersFromSkillGroupWithOptions(request *RemovePhoneNumbersFromSkillGroupRequest, runtime *util.RuntimeOptions) (_result *RemovePhoneNumbersFromSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemovePhoneNumbersFromSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemovePhoneNumbersFromSkillGroup"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemovePhoneNumbersFromSkillGroup(request *RemovePhoneNumbersFromSkillGroupRequest) (_result *RemovePhoneNumbersFromSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemovePhoneNumbersFromSkillGroupResponse{}
	_body, _err := client.RemovePhoneNumbersFromSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompleteAttendedTransferWithOptions(request *CompleteAttendedTransferRequest, runtime *util.RuntimeOptions) (_result *CompleteAttendedTransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CompleteAttendedTransferResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CompleteAttendedTransfer"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompleteAttendedTransfer(request *CompleteAttendedTransferRequest) (_result *CompleteAttendedTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompleteAttendedTransferResponse{}
	_body, _err := client.CompleteAttendedTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetUserPasswordWithOptions(request *ResetUserPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetUserPassword"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetUserPassword(request *ResetUserPasswordRequest) (_result *ResetUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.ResetUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTurnServerListWithOptions(request *GetTurnServerListRequest, runtime *util.RuntimeOptions) (_result *GetTurnServerListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTurnServerListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTurnServerList"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTurnServerList(request *GetTurnServerListRequest) (_result *GetTurnServerListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTurnServerListResponse{}
	_body, _err := client.GetTurnServerListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNumberLocationWithOptions(request *GetNumberLocationRequest, runtime *util.RuntimeOptions) (_result *GetNumberLocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetNumberLocationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetNumberLocation"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNumberLocation(request *GetNumberLocationRequest) (_result *GetNumberLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNumberLocationResponse{}
	_body, _err := client.GetNumberLocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRamUsersWithOptions(request *ListRamUsersRequest, runtime *util.RuntimeOptions) (_result *ListRamUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRamUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRamUsers"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRamUsers(request *ListRamUsersRequest) (_result *ListRamUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRamUsersResponse{}
	_body, _err := client.ListRamUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MuteCallWithOptions(request *MuteCallRequest, runtime *util.RuntimeOptions) (_result *MuteCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MuteCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MuteCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MuteCall(request *MuteCallRequest) (_result *MuteCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MuteCallResponse{}
	_body, _err := client.MuteCallWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("AnswerCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListIntervalAgentReportWithOptions(request *ListIntervalAgentReportRequest, runtime *util.RuntimeOptions) (_result *ListIntervalAgentReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListIntervalAgentReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListIntervalAgentReport"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIntervalAgentReport(request *ListIntervalAgentReportRequest) (_result *ListIntervalAgentReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIntervalAgentReportResponse{}
	_body, _err := client.ListIntervalAgentReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCallDetailRecordsWithOptions(request *ListCallDetailRecordsRequest, runtime *util.RuntimeOptions) (_result *ListCallDetailRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCallDetailRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCallDetailRecords"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCallDetailRecords(request *ListCallDetailRecordsRequest) (_result *ListCallDetailRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCallDetailRecordsResponse{}
	_body, _err := client.ListCallDetailRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemovePhoneNumbersWithOptions(request *RemovePhoneNumbersRequest, runtime *util.RuntimeOptions) (_result *RemovePhoneNumbersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemovePhoneNumbersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemovePhoneNumbers"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemovePhoneNumbers(request *RemovePhoneNumbersRequest) (_result *RemovePhoneNumbersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemovePhoneNumbersResponse{}
	_body, _err := client.RemovePhoneNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelAttendedTransferWithOptions(request *CancelAttendedTransferRequest, runtime *util.RuntimeOptions) (_result *CancelAttendedTransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelAttendedTransferResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelAttendedTransfer"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelAttendedTransfer(request *CancelAttendedTransferRequest) (_result *CancelAttendedTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelAttendedTransferResponse{}
	_body, _err := client.CancelAttendedTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TakeBreakWithOptions(request *TakeBreakRequest, runtime *util.RuntimeOptions) (_result *TakeBreakResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TakeBreakResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TakeBreak"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TakeBreak(request *TakeBreakRequest) (_result *TakeBreakResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TakeBreakResponse{}
	_body, _err := client.TakeBreakWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHistoricalSkillGroupReportWithOptions(request *ListHistoricalSkillGroupReportRequest, runtime *util.RuntimeOptions) (_result *ListHistoricalSkillGroupReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHistoricalSkillGroupReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHistoricalSkillGroupReport"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHistoricalSkillGroupReport(request *ListHistoricalSkillGroupReportRequest) (_result *ListHistoricalSkillGroupReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHistoricalSkillGroupReportResponse{}
	_body, _err := client.ListHistoricalSkillGroupReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendDtmfSignalingWithOptions(request *SendDtmfSignalingRequest, runtime *util.RuntimeOptions) (_result *SendDtmfSignalingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SendDtmfSignalingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SendDtmfSignaling"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendDtmfSignaling(request *SendDtmfSignalingRequest) (_result *SendDtmfSignalingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendDtmfSignalingResponse{}
	_body, _err := client.SendDtmfSignalingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRecentCallDetailRecordsWithOptions(request *ListRecentCallDetailRecordsRequest, runtime *util.RuntimeOptions) (_result *ListRecentCallDetailRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRecentCallDetailRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRecentCallDetailRecords"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRecentCallDetailRecords(request *ListRecentCallDetailRecordsRequest) (_result *ListRecentCallDetailRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRecentCallDetailRecordsResponse{}
	_body, _err := client.ListRecentCallDetailRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitiateAttendedTransferWithOptions(request *InitiateAttendedTransferRequest, runtime *util.RuntimeOptions) (_result *InitiateAttendedTransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InitiateAttendedTransferResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InitiateAttendedTransfer"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitiateAttendedTransfer(request *InitiateAttendedTransferRequest) (_result *InitiateAttendedTransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitiateAttendedTransferResponse{}
	_body, _err := client.InitiateAttendedTransferWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MakeCallWithOptions(request *MakeCallRequest, runtime *util.RuntimeOptions) (_result *MakeCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MakeCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MakeCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MakeCall(request *MakeCallRequest) (_result *MakeCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MakeCallResponse{}
	_body, _err := client.MakeCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetInstance"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUsersToSkillGroupWithOptions(request *AddUsersToSkillGroupRequest, runtime *util.RuntimeOptions) (_result *AddUsersToSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddUsersToSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddUsersToSkillGroup"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUsersToSkillGroup(request *AddUsersToSkillGroupRequest) (_result *AddUsersToSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUsersToSkillGroupResponse{}
	_body, _err := client.AddUsersToSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConfigItemsWithOptions(request *ListConfigItemsRequest, runtime *util.RuntimeOptions) (_result *ListConfigItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListConfigItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConfigItems"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConfigItems(request *ListConfigItemsRequest) (_result *ListConfigItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConfigItemsResponse{}
	_body, _err := client.ListConfigItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SignInGroupWithOptions(request *SignInGroupRequest, runtime *util.RuntimeOptions) (_result *SignInGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SignInGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SignInGroup"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SignInGroup(request *SignInGroupRequest) (_result *SignInGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SignInGroupResponse{}
	_body, _err := client.SignInGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRealtimeInstanceStatesWithOptions(request *GetRealtimeInstanceStatesRequest, runtime *util.RuntimeOptions) (_result *GetRealtimeInstanceStatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRealtimeInstanceStatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRealtimeInstanceStates"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRealtimeInstanceStates(request *GetRealtimeInstanceStatesRequest) (_result *GetRealtimeInstanceStatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRealtimeInstanceStatesResponse{}
	_body, _err := client.GetRealtimeInstanceStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySkillGroupWithOptions(request *ModifySkillGroupRequest, runtime *util.RuntimeOptions) (_result *ModifySkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySkillGroup"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySkillGroup(request *ModifySkillGroupRequest) (_result *ModifySkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySkillGroupResponse{}
	_body, _err := client.ModifySkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserWithOptions(request *ModifyUserRequest, runtime *util.RuntimeOptions) (_result *ModifyUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUser(request *ModifyUserRequest) (_result *ModifyUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserResponse{}
	_body, _err := client.ModifyUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddPhoneNumberToSkillGroupsWithOptions(request *AddPhoneNumberToSkillGroupsRequest, runtime *util.RuntimeOptions) (_result *AddPhoneNumberToSkillGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddPhoneNumberToSkillGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddPhoneNumberToSkillGroups"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddPhoneNumberToSkillGroups(request *AddPhoneNumberToSkillGroupsRequest) (_result *AddPhoneNumberToSkillGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddPhoneNumberToSkillGroupsResponse{}
	_body, _err := client.AddPhoneNumberToSkillGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDevicesWithOptions(request *ListDevicesRequest, runtime *util.RuntimeOptions) (_result *ListDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDevicesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDevices"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevices(request *ListDevicesRequest) (_result *ListDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevicesResponse{}
	_body, _err := client.ListDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RetrieveCallWithOptions(request *RetrieveCallRequest, runtime *util.RuntimeOptions) (_result *RetrieveCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RetrieveCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RetrieveCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RetrieveCall(request *RetrieveCallRequest) (_result *RetrieveCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RetrieveCallResponse{}
	_body, _err := client.RetrieveCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSkillGroupsWithOptions(request *ListSkillGroupsRequest, runtime *util.RuntimeOptions) (_result *ListSkillGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSkillGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSkillGroups"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSkillGroups(request *ListSkillGroupsRequest) (_result *ListSkillGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSkillGroupsResponse{}
	_body, _err := client.ListSkillGroupsWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("HoldCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RegisterDeviceWithOptions(request *RegisterDeviceRequest, runtime *util.RuntimeOptions) (_result *RegisterDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterDevice"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterDevice(request *RegisterDeviceRequest) (_result *RegisterDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.RegisterDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemovePersonalNumbersFromUserWithOptions(request *RemovePersonalNumbersFromUserRequest, runtime *util.RuntimeOptions) (_result *RemovePersonalNumbersFromUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemovePersonalNumbersFromUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemovePersonalNumbersFromUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemovePersonalNumbersFromUser(request *RemovePersonalNumbersFromUserRequest) (_result *RemovePersonalNumbersFromUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemovePersonalNumbersFromUserResponse{}
	_body, _err := client.RemovePersonalNumbersFromUserWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstance"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListOutboundNumbersOfUserWithOptions(request *ListOutboundNumbersOfUserRequest, runtime *util.RuntimeOptions) (_result *ListOutboundNumbersOfUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOutboundNumbersOfUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOutboundNumbersOfUser"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOutboundNumbersOfUser(request *ListOutboundNumbersOfUserRequest) (_result *ListOutboundNumbersOfUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOutboundNumbersOfUserResponse{}
	_body, _err := client.ListOutboundNumbersOfUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PollUserStatusWithOptions(request *PollUserStatusRequest, runtime *util.RuntimeOptions) (_result *PollUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PollUserStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PollUserStatus"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PollUserStatus(request *PollUserStatusRequest) (_result *PollUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PollUserStatusResponse{}
	_body, _err := client.PollUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReadyForServiceWithOptions(request *ReadyForServiceRequest, runtime *util.RuntimeOptions) (_result *ReadyForServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReadyForServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReadyForService"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReadyForService(request *ReadyForServiceRequest) (_result *ReadyForServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReadyForServiceResponse{}
	_body, _err := client.ReadyForServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMultiChannelRecordingWithOptions(request *GetMultiChannelRecordingRequest, runtime *util.RuntimeOptions) (_result *GetMultiChannelRecordingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMultiChannelRecordingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMultiChannelRecording"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMultiChannelRecording(request *GetMultiChannelRecordingRequest) (_result *GetMultiChannelRecordingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultiChannelRecordingResponse{}
	_body, _err := client.GetMultiChannelRecordingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BargeInCallWithOptions(request *BargeInCallRequest, runtime *util.RuntimeOptions) (_result *BargeInCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BargeInCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BargeInCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BargeInCall(request *BargeInCallRequest) (_result *BargeInCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BargeInCallResponse{}
	_body, _err := client.BargeInCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPhoneNumbersOfSkillGroupWithOptions(request *ListPhoneNumbersOfSkillGroupRequest, runtime *util.RuntimeOptions) (_result *ListPhoneNumbersOfSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPhoneNumbersOfSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPhoneNumbersOfSkillGroup"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPhoneNumbersOfSkillGroup(request *ListPhoneNumbersOfSkillGroupRequest) (_result *ListPhoneNumbersOfSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPhoneNumbersOfSkillGroupResponse{}
	_body, _err := client.ListPhoneNumbersOfSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SignOutGroupWithOptions(request *SignOutGroupRequest, runtime *util.RuntimeOptions) (_result *SignOutGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SignOutGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SignOutGroup"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SignOutGroup(request *SignOutGroupRequest) (_result *SignOutGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SignOutGroupResponse{}
	_body, _err := client.SignOutGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveRTCStatsV2WithOptions(request *SaveRTCStatsV2Request, runtime *util.RuntimeOptions) (_result *SaveRTCStatsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveRTCStatsV2Response{}
	_body, _err := client.DoRPCRequest(tea.String("SaveRTCStatsV2"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveRTCStatsV2(request *SaveRTCStatsV2Request) (_result *SaveRTCStatsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveRTCStatsV2Response{}
	_body, _err := client.SaveRTCStatsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHistoricalCallerReportWithOptions(request *GetHistoricalCallerReportRequest, runtime *util.RuntimeOptions) (_result *GetHistoricalCallerReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHistoricalCallerReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHistoricalCallerReport"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHistoricalCallerReport(request *GetHistoricalCallerReportRequest) (_result *GetHistoricalCallerReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHistoricalCallerReportResponse{}
	_body, _err := client.GetHistoricalCallerReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserLevelsOfSkillGroupWithOptions(request *ModifyUserLevelsOfSkillGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyUserLevelsOfSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyUserLevelsOfSkillGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUserLevelsOfSkillGroup"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserLevelsOfSkillGroup(request *ModifyUserLevelsOfSkillGroupRequest) (_result *ModifyUserLevelsOfSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserLevelsOfSkillGroupResponse{}
	_body, _err := client.ModifyUserLevelsOfSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTerminalLogWithOptions(request *SaveTerminalLogRequest, runtime *util.RuntimeOptions) (_result *SaveTerminalLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveTerminalLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveTerminalLog"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTerminalLog(request *SaveTerminalLogRequest) (_result *SaveTerminalLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTerminalLogResponse{}
	_body, _err := client.SaveTerminalLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRealtimeSkillGroupStatesWithOptions(request *ListRealtimeSkillGroupStatesRequest, runtime *util.RuntimeOptions) (_result *ListRealtimeSkillGroupStatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRealtimeSkillGroupStatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRealtimeSkillGroupStates"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRealtimeSkillGroupStates(request *ListRealtimeSkillGroupStatesRequest) (_result *ListRealtimeSkillGroupStatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRealtimeSkillGroupStatesResponse{}
	_body, _err := client.ListRealtimeSkillGroupStatesWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CreateSkillGroup"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) PickOutboundNumbersWithOptions(request *PickOutboundNumbersRequest, runtime *util.RuntimeOptions) (_result *PickOutboundNumbersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PickOutboundNumbersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PickOutboundNumbers"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PickOutboundNumbers(request *PickOutboundNumbersRequest) (_result *PickOutboundNumbersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PickOutboundNumbersResponse{}
	_body, _err := client.PickOutboundNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseCallWithOptions(request *ReleaseCallRequest, runtime *util.RuntimeOptions) (_result *ReleaseCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseCall"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseCall(request *ReleaseCallRequest) (_result *ReleaseCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseCallResponse{}
	_body, _err := client.ReleaseCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLoginDetailsWithOptions(request *GetLoginDetailsRequest, runtime *util.RuntimeOptions) (_result *GetLoginDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLoginDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLoginDetails"), tea.String("2020-07-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLoginDetails(request *GetLoginDetailsRequest) (_result *GetLoginDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoginDetailsResponse{}
	_body, _err := client.GetLoginDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
